## Usage

Sources:

- [Official README](https://github.com/citusdata/pg_intpair/blob/master/README.md)
- [Extension control file](https://github.com/citusdata/pg_intpair/blob/master/intpair.control)
- [Version 0.2 extension SQL](https://github.com/citusdata/pg_intpair/blob/master/intpair--0.2.sql)
- [Regression SQL examples](https://github.com/citusdata/pg_intpair/blob/master/sql/intpair.sql)

intpair installs a compact fixed-length base type for a pair of signed 64-bit integers. The SQL extension name is intpair, while the upstream project and package are named pg_intpair; use it when dense storage, comparison, and indexing of exactly two integers are preferable to a composite type or array.

### Core Workflow

Construct values with the two-argument function or the textual pair representation. Components are addressed with zero-based subscripts.

```sql
CREATE EXTENSION intpair;

CREATE TABLE edges (
    endpoints int64pair NOT NULL
);

INSERT INTO edges VALUES
    (int64pair(10, 20)),
    ('(10,30)'::int64pair),
    ('(-1,5)'::int64pair);

SELECT endpoints,
       endpoints[0] AS first_value,
       endpoints[1] AS second_value
FROM edges
ORDER BY endpoints;
```

### Operators and Indexes

Version 0.2 defines equality and lexicographic ordering over the first component and then the second. It supplies default B-tree and hash operator classes, so ordinary indexes support equality, range predicates, and ordering.

```sql
CREATE INDEX edges_endpoints_idx ON edges (endpoints);

SELECT *
FROM edges
WHERE endpoints >= int64pair(10, 0)
  AND endpoints <  int64pair(11, 0);
```

### Migration and Caveats

The upstream README demonstrates migration from a two-bigint composite through textual I/O. Test that conversion on a copy first, because the input must match the extension's pair syntax and an implicit cast changes resolution for later queries.

```sql
ALTER TABLE legacy_edges
ALTER COLUMN endpoints TYPE int64pair
USING endpoints::text::int64pair;
```

The type is always a 16-byte pair and does not represent a variable-length list. It is relocatable and declares no preload or restart requirement. Applications should use the exact extension name intpair and type name shown in the examples rather than deriving them from the pg_intpair package name.
