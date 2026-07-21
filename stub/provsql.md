## Usage

Sources:

- [ProvSQL 1.11.0 documentation](https://github.com/PierreSenellart/provsql/blob/v1.11.0/doc/provsql.md)
- [ProvSQL 1.11.0 release](https://github.com/PierreSenellart/provsql/releases/tag/v1.11.0)
- [ProvSQL 1.10.0 release](https://github.com/PierreSenellart/provsql/releases/tag/v1.10.0)
- [ProvSQL 1.11.0 control file](https://github.com/PierreSenellart/provsql/blob/v1.11.0/provsql.common.control)
- [ProvSQL user documentation](https://provsql.org/docs/user/introduction.html)

`provsql` adds semiring provenance and uncertainty management to PostgreSQL. Upstream documents provenance tracking, semiring evaluation, probabilities, Shapley and Banzhaf values, where-provenance, update provenance, and temporal features.

### Load and Track Provenance

```ini
shared_preload_libraries = 'provsql'
```

```sql
CREATE EXTENSION provsql CASCADE;
```

The `CASCADE` form installs `uuid-ossp` automatically if needed. The getting-started guide says the preload step is mandatory because ProvSQL installs a planner hook.

```sql
SELECT provsql.add_provenance('mytable');

SELECT name, provenance()
FROM mytable;

SELECT provsql.remove_provenance('mytable');
```

The user docs also describe provenance mappings:

```sql
SELECT create_provenance_mapping('my_mapping', 'mytable', 'column_name');
SELECT create_provenance_mapping_view('my_mapping_view', 'mytable', 'column_name');
```

### Probability and Influence

Assign probabilities to tuple tokens:

```sql
SELECT set_prob(provenance(), 0.8)
FROM mytable
WHERE id = 1;

SELECT name, probability_evaluate(provenance()) AS prob
FROM mytable;
```

Compute influence scores:

```sql
SELECT shapley(provenance(), m.token)
FROM mytable, my_mapping AS m;

SELECT banzhaf(provenance(), m.token)
FROM mytable, my_mapping AS m;
```

The docs also describe `shapley_all_vars` and `banzhaf_all_vars` for computing scores for all input variables at once.

### Built-in Semirings

Built-in semiring functions use a provenance token and a provenance mapping table:

```sql
SELECT name, sr_boolean(provenance(), 'my_mapping')
FROM mytable;

SELECT name, sr_formula(provenance(), 'my_mapping')
FROM mytable;

SELECT name, sr_how(provenance(), 'my_mapping')
FROM mytable;
```

Current docs include compiled wrappers for `sr_how`, `sr_which`, `sr_tropical`, `sr_viterbi`, `sr_lukasiewicz`, `sr_minmax`, and `sr_maxmin`. For PostgreSQL 14 and later they also include `sr_temporal`, `sr_interval_num`, and `sr_interval_int` over multirange values.

```sql
SELECT city,
       sr_minmax(provenance(), 'personnel_level',
                 'unclassified'::classification_level) AS clearance
FROM (SELECT DISTINCT city FROM personnel) AS t;

SELECT entity_id, sr_temporal(provenance(), 'validity_mapping')
FROM mytable;
```

Advanced users can still define custom semirings and evaluate them with `provenance_evaluate` or `aggregation_evaluate`; upstream recommends the compiled semirings when one matches the needed algebra.

### Extra Modes and Helpers

Session GUCs documented upstream include:

```sql
SET provsql.active = on;
SET provsql.where_provenance = on;
SET provsql.update_provenance = on;
SET provsql.last_eval_method = on;
SET provsql.tool_search_path = '/opt/d4:/home/postgres/bin';
SET provsql.aggtoken_text_as_uuid = on;
```

`provsql.tool_search_path` is used for external probability and visualization tools such as `d4`, `c2d`, `dsharp`, `minic2d`, `weightmc`, and `graph-easy`. `provsql.last_eval_method` stores the last chosen probability-evaluation method. `provsql.aggtoken_text_as_uuid` makes aggregate-token cells render as their provenance UUIDs; `agg_token_value_text(token)` can recover the display text for those aggregate tokens.

The user guide separately documents where-provenance helpers, update provenance, temporal helpers such as `get_valid_time`, `timetravel`, `timeslice`, `history`, and `undo`, circuit-inspection helpers `circuit_subgraph(root, max_depth)` and `resolve_input(uuid)`, and `setup_search_path()` for preparing the helper search path.

### Current Probability and Inference Surface

The 1.9 through 1.11 releases materially expand SQL coverage and probability evaluation:

- subqueries outside `FROM`, including `EXISTS`, `NOT EXISTS`, `IN`, `NOT IN`, `ANY`, `ALL`, row-valued `IN`, scalar subqueries, and `ARRAY(SELECT ...)`;
- `LEFT`, `RIGHT`, and `FULL` outer joins, plus corrected `EXCEPT` and `EXCEPT ALL` provenance;
- SQL-faithful `NULL` handling for aggregates and exact `HAVING` aggregate probabilities for `COUNT`, `SUM`, `MIN`, `MAX`, and `AVG`;
- probability-method selection through the method catalog and cost chooser, with `karp-luby`, `stopping-rule`, `sieve`, `d-tree`, and `probability_bounds`;
- exact bounded-treewidth recursive reachability, unsafe-UCQ joint-width compilation, Möbius inversion for safe UCQs, and absorptive provenance for cyclic recursion;
- conditional events and distributions through the `target | evidence` operator and the whole-tuple `given()`/prefix form;
- continuous and discrete `random_variable` families, including normal, gamma, log-normal, beta, Weibull, Pareto, inverse-gamma, inverse-Gaussian, logistic, Poisson, binomial, geometric, hypergeometric, and negative-binomial distributions;
- hierarchical Bayesian models where distribution parameters are themselves random variables, with conjugate posterior updates when a closed form is available;
- maintained provenance mappings that remain correct as source data changes, plus SQL-conformant `NULL` behavior for `NOT IN`, `EXCEPT`, and nullable random variables.

For example, condition a continuous value on observed evidence and read the posterior expectation:

```sql
WITH model AS (
  SELECT normal(20, 5) AS reading
)
SELECT expected(reading | (reading > 25))
FROM model;
```

The `agg_token` type supports arithmetic, unary minus, and comparisons for probabilistic aggregate expressions. Use the official probability and continuous-distribution chapters to choose between exact, compiled, and sampling-based evaluation methods.

### Notes

- The 1.11.0 control file sets `default_version = '1.11.0'`, requires `uuid-ossp`, marks the extension trusted, and is not relocatable.
- Upstream documentation says ProvSQL has been tested on PostgreSQL 10 through 18.
- `provsql.update_provenance` and the multirange semirings require PostgreSQL 14 or later.
- Update-provenance tracking remains experimental; validate its storage and performance costs before enabling it broadly.
