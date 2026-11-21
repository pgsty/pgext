## Usage

Sources:

- [pg_gaps.control](https://github.com/gburd/pg_gaps/blob/d35680cc3303e9791465256bfa9f5d8feb5615c2/pg_gaps.control)
- [pg_gaps--0.1.0.sql](https://github.com/gburd/pg_gaps/blob/d35680cc3303e9791465256bfa9f5d8feb5615c2/pg_gaps--0.1.0.sql)
- [ROADMAP.md](https://github.com/gburd/pg_gaps/blob/d35680cc3303e9791465256bfa9f5d8feb5615c2/ROADMAP.md)
- [doc/DESIGN.md](https://github.com/gburd/pg_gaps/blob/d35680cc3303e9791465256bfa9f5d8feb5615c2/doc/DESIGN.md)
- [LICENSE](https://github.com/gburd/pg_gaps/blob/d35680cc3303e9791465256bfa9f5d8feb5615c2/LICENSE)

`pg_gaps` is an experimental Bayesian-model extension for incomplete relational data. The reviewed 0.1.0 revision contains model catalogs, exact inference, optional sampling and conditional-probability learning. The roadmap records implemented phases and takes precedence over the outdated README statement that the extension does nothing. It targets PostgreSQL 17–19 and installs in the fixed `gaps` schema.

### Define a Small Model

Create the extension as a superuser. Model functions use the caller's table permissions and the catalogs apply role-based row security.

```sql
CREATE EXTENSION pg_gaps;
CREATE TABLE observations (id integer PRIMARY KEY, state text);
INSERT INTO observations VALUES (1, 'yes'), (2, NULL);
SELECT gaps.create_model('example');
SELECT gaps.declare_attribute(
  'example', 'observations', 'state', ARRAY['yes', 'no']);
SELECT gaps.add_stripe('example', 'state', 'observations', 'state');
SELECT gaps.add_factor(
  'example', 'prior', 'state', ARRAY[]::text[], ARRAY[0.6,0.4]::float8[]);
SELECT * FROM gaps.validate('example');
SELECT * FROM gaps.marginals('example');
SELECT * FROM gaps.impute('example');
```

A declared attribute has a finite domain; entity keys default to the table's primary key. Stripes identify groups of random variables, factors attach conditional-probability tables, and parent mappings connect related entities. `gaps.set_cpt()` replaces a factor's probabilities; `gaps.learn_cpts()` estimates them from complete observations with optional additive smoothing.

### Interpret Results

`gaps.marginals()` reports a distribution for each missing value. Variables in one component can be correlated, so multiplying their marginal probabilities does not produce the joint distribution. `gaps.impute()` returns the joint most likely assignment for exact components. With sampling it returns per-variable marginal choices, not a joint optimum, and the joint-probability field is NULL.

Inference defaults to exact computation or an error when resource bounds are exceeded. Positive `gaps.samples` enables sampling for components beyond those bounds; agreement between chains is not proof of correctness. `gaps.max_factor_cells` and `gaps.max_total_cells` bound inference allocations.

### Boundaries

The catalog tables and their sequences are registered for dump/restore. Data is read under caller privileges; learned probabilities become visible to roles allowed to read the model. Treat the learned model as data that can disclose information about its training observations.

Schema changes can invalidate model assumptions, and constants whose input depends on session settings remain an open issue. Validation is conservative and a failed structural proof does not always establish an invalid grounded model. This early source revision has no documented stable upgrade contract; keep model backups and test changes in an isolated database.
