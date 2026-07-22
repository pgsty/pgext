## Usage

Sources:

- [Official README](https://github.com/adjust/istore/blob/46c1cfeceeea193b75fd3fa11bc6959d8ac4d26f/README.md)
- [Official extension SQL](https://github.com/adjust/istore/blob/46c1cfeceeea193b75fd3fa11bc6959d8ac4d26f/istore--0.1.12.sql)
- [Official extension control file](https://github.com/adjust/istore/blob/46c1cfeceeea193b75fd3fa11bc6959d8ac4d26f/istore.control)

`istore` provides compact integer-key maps for PostgreSQL. The `istore` type stores integer values, while `bigistore` stores bigint values; both are useful for sparse counters, distributions, and analytical aggregates keyed by integer identifiers.

### Core Workflow

Text input resembles hstore, and array constructors can aggregate repeated keys:

```sql
CREATE EXTENSION istore;

SELECT istore(ARRAY[1, 2, 1, 3, 2, 2]);
SELECT '1=>4,2=>5'::istore -> 1;
SELECT '1=>4,2=>5'::istore + '1=>4,3=>6'::istore;

CREATE INDEX metrics_keys_idx
  ON metrics USING gin (values istore_key_ops);
```

Key lookup uses `->`; `?`, `?&`, and `?|` test key presence. `||` concatenates stores, while arithmetic operators combine matching keys. The `istore_key_ops` GIN operator class accelerates key-existence predicates.

### Important Objects and Caveats

Constructors and transforms include `istore(...)`, `bigistore(...)`, `akeys`, `avals`, `each`, `slice`, `compact`, `fill_gaps`, `accumulate`, `clamp`, and delete helpers. `sum_up` summarizes values, and aggregate SUM returns `bigistore` so that totals have a wider value type. Per-key minimum and maximum aggregates are also provided.

The cast from `istore` to `bigistore` is implicit; narrowing assignments can overflow. Arithmetic and division still require application-level checks for overflow and zero divisors. Large values can be TOASTed, and GIN indexes add write and maintenance cost, so benchmark the actual distribution and update workload instead of treating upstream examples as performance guarantees.
