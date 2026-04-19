## Usage

Sources: [README](https://github.com/PierreSenellart/provsql/blob/master/doc/provsql.md), [getting started](https://provsql.org/docs/user/getting-provsql.html), [user docs](https://provsql.org/docs/), [SQL API index](https://provsql.org/docs/sql/)

`provsql` adds semiring provenance and uncertainty management to PostgreSQL. Upstream documents provenance tracking, semiring evaluation, probabilities, Shapley and Banzhaf values, where-provenance, update provenance, and temporal features.

### Load the extension

```ini
shared_preload_libraries = 'provsql'
```

```sql
CREATE EXTENSION provsql CASCADE;
```

The `CASCADE` form installs `uuid-ossp` automatically if needed. The getting-started guide says the preload step is mandatory because ProvSQL installs a planner hook.

### Enable provenance on tables

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

### Probability and semiring workflows

Assign probabilities to tuple tokens:

```sql
SELECT set_prob(provenance(), 0.8)
FROM mytable
WHERE id = 1;
```

Evaluate provenance in a semiring:

```sql
SELECT city,
       provenance_evaluate(
         provenance(),
         'personnel_level',
         'unclassified'::classification_level,
         'security_plus',
         'security_times'
       )
FROM (SELECT DISTINCT city FROM personnel) AS t;
```

Compute influence scores:

```sql
SELECT shapley(provenance(), m.token)
FROM my_mapping AS m;
```

The docs also describe `shapley_all_vars`, `banzhaf`, and `banzhaf_all_vars`.

### Extra modes

Session GUCs documented upstream include:

```sql
SET provsql.where_provenance = on;
SET provsql.update_provenance = on;
```

The user guide separately documents temporal helpers such as `get_valid_time`, `timetravel`, `timeslice`, `history`, and `undo`.

### Notes

- Upstream tests ProvSQL on PostgreSQL 10 through 18.
- Git tags show `v1.2.3` as the current packaged release in the repository.
