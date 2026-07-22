## Usage

Sources:

- [Official README](https://github.com/adjust/pg-numhstore/blob/754e2482998b83710b49298141f79e9d4fef3161/README.md)
- [Extension SQL for version 0.1.7](https://github.com/adjust/pg-numhstore/blob/754e2482998b83710b49298141f79e9d4fef3161/numhstore--0.1.7.sql)
- [Extension control file](https://github.com/adjust/pg-numhstore/blob/754e2482998b83710b49298141f79e9d4fef3161/numhstore.control)

`numhstore` adds numeric key/value stores named `inthstore` and `floathstore`. They reuse the textual shape and binary representation of `hstore` while adding arithmetic, summaries, and aggregates for sparse integer or floating-point maps.

### Core Workflow

The extension requires PostgreSQL's `hstore` extension.

```sql
CREATE EXTENSION hstore;
CREATE EXTENSION numhstore;
```

Cast an `hstore` value to the numeric type that matches the stored values, then perform per-key arithmetic.

```sql
SELECT 'a=>1,b=>2'::hstore::inthstore
     + 'a=>4,c=>8'::hstore::inthstore;

SELECT 'cpu=>1.5,mem=>2.25'::hstore::floathstore * 2::numeric;
```

Aggregate sparse maps across rows. Missing keys follow the operator's own semantics, so test the exact operation rather than assuming ordinary scalar NULL behavior.

```sql
CREATE TEMP TABLE counters(v inthstore);
INSERT INTO counters VALUES
  ('ok=>2,error=>1'),
  ('ok=>3,retry=>4');

SELECT sum(v), avg(v)
FROM counters;
```

### Important Objects

- `inthstore`: integer-valued store with implicit casts to and from `hstore` and to `floathstore`.
- `floathstore`: floating-point/numeric store with implicit casts to and from `hstore`.
- `+`, `-`, `*`, `/`: per-key arithmetic between stores or between a store and a scalar.
- `sum(inthstore)`, `sum(floathstore)`, `avg(inthstore)`, `avg(floathstore)`: aggregate sparse maps.
- `array_count(integer[])`, `array_count(text[])`, `array_add(text[], integer[])`: build or combine count maps.
- `hstore_sum_up(hstore)`, `hstore_array(hstore)`, `hstore_length(hstore)`: sum values, return values as an array, or count keys.

### Semantics and Caveats

The custom types are binary-compatible with `hstore`, but their numeric values are still represented through hstore-style text. Invalid numeric strings fail when arithmetic or casts parse them. Operator behavior for a key present on only one side is not uniform: addition and subtraction retain union-like keys, while multiplication and division can produce NULL or zero-like results according to the implementation. Test missing keys, NULLs, division by zero, overflow, and floating-point rounding with your data.

Implicit casts can make overloaded expressions surprising; use explicit casts at API and schema boundaries. The pinned upstream source does not state a current PostgreSQL compatibility matrix, so validate version `0.1.7` on the target server before production use.
