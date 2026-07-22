## Usage

Sources:

- [Official upstream documentation](https://github.com/umitanuki/tinyint-postgresql/blob/76495ba27565df2e6758e8f26778e16e963004fb/doc/tinyint.md)
- [Official extension SQL definition](https://github.com/umitanuki/tinyint-postgresql/blob/76495ba27565df2e6758e8f26778e16e963004fb/tinyint.sql.in.c)
- [Official regression tests](https://github.com/umitanuki/tinyint-postgresql/blob/76495ba27565df2e6758e8f26778e16e963004fb/sql/tinyint.sql)

`tinyint` adds a one-byte signed integer type for compact values whose documented range is -127 through 127. It supports integer-style casts, comparisons, arithmetic, aggregates, and indexes, but row alignment can erase the expected storage saving and the old upstream code needs compatibility testing on the target PostgreSQL release.

### Core Workflow

Create the extension and use the type for values that are guaranteed to stay in its narrow range:

```sql
CREATE EXTENSION tinyint;

CREATE TABLE sensor_flags (
    sensor_id bigint PRIMARY KEY,
    state tinyint NOT NULL CHECK (state BETWEEN 0 AND 3)
);

INSERT INTO sensor_flags VALUES (1, 2::tinyint);

SELECT state, count(*)
FROM sensor_flags
GROUP BY state
ORDER BY state;
```

Conversions from `smallint`, `integer`, and `bigint` are assignment casts and reject out-of-range values. Conversions from `tinyint` to the wider integer types are implicit.

### Operators, Aggregates, and Indexes

- Arithmetic includes unary signs and `+`, `-`, `*`, and `/`; mixed operations with wider integer types return the wider type.
- Comparisons work between `tinyint` and `smallint`, `integer`, or `bigint` in either argument order.
- `min`, `max`, `sum`, and `avg` are defined; `sum` uses a `bigint` transition/result and `avg` uses PostgreSQL's integer average finalizer.
- `btree_tinyint_ops` and `hash_tinyint_ops` support scalar B-tree and hash indexing. `_tinyint_ops` supports GIN indexing of `tinyint[]` arrays.

### Caveats

Arithmetic and casts raise an error on overflow, so validate inputs before bulk loads and migrations. This is a signed type and is not a drop-in replacement for an unsigned `TINYINT` from another database.

Although the datum itself occupies one byte, PostgreSQL can add alignment padding depending on column order and surrounding types. Measure the actual tuple and index sizes before adopting it purely for space savings. The reviewed upstream revision dates from 2017 and its published compatibility material is old; build, install, upgrade, dump/restore, operator, aggregate, and index tests are required on each intended PostgreSQL major version.
