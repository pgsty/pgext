


## Usage

Sources:

- [PGXN pg_roaringbitmap 1.2.0](https://pgxn.org/dist/pg_roaringbitmap/1.2.0/)
- [pg_roaringbitmap README](https://github.com/ChenHuajun/pg_roaringbitmap)
- [pg_roaringbitmap CHANGELOG](https://github.com/ChenHuajun/pg_roaringbitmap/blob/master/CHANGELOG.md)
- [Local package metadata](../db/extension.csv)

`pg_roaringbitmap` installs the PostgreSQL extension `roaringbitmap`, which provides compressed bitmap types and set-operation functions backed by Roaring Bitmaps. Use it for compact integer-set storage, fast unions/intersections, cohort filters, faceting, and bitmap aggregation.

v1.2.0 adds `rb_runoptimize()` / `rb64_runoptimize()` to shrink bitmap binary size, preserves the legacy `rb_exsit` spelling for backward compatibility, adds PostgreSQL 19 compatibility, and validates untrusted bitmap input in receive functions.

### Create the Extension

```sql
CREATE EXTENSION IF NOT EXISTS roaringbitmap;

SET roaringbitmap.output_format = 'array';
SELECT rb_build(ARRAY[1, 2, 3, 4, 5]);
```

`roaringbitmap.output_format` can be `bytea` or `array`. The default output format is `bytea`, which is better for large bitmaps; `array` is easier to read interactively.

### Data Types

- `roaringbitmap` stores unsigned 32-bit integer sets over the logical range `[0, 4294967296)`.
- `roaringbitmap64` stores unsigned 64-bit integer sets and uses the `rb64_` function family.

```sql
CREATE TABLE cohorts (
  segment text PRIMARY KEY,
  users32 roaringbitmap,
  users64 roaringbitmap64
);
```

### Build and Convert

```sql
INSERT INTO cohorts(segment, users32)
VALUES ('trial', rb_build(ARRAY[1, 2, 3, 100, 200]));

INSERT INTO cohorts(segment, users32)
SELECT 'active', rb_build_agg(id)
FROM generate_series(1, 100000) AS id;

SELECT rb_cardinality(users32) FROM cohorts WHERE segment = 'active';
SELECT rb_to_array(users32) FROM cohorts WHERE segment = 'trial';
SELECT rb_iterate(users32) FROM cohorts WHERE segment = 'trial';
```

For 64-bit values, use `rb64_build()`, `rb64_build_agg()`, `rb64_to_array()`, and `rb64_iterate()`.

### Set Operations

```sql
SELECT rb_build(ARRAY[1,2,3]) | rb_build(ARRAY[3,4,5]);  -- union
SELECT rb_build(ARRAY[1,2,3]) & rb_build(ARRAY[3,4,5]);  -- intersection
SELECT rb_build(ARRAY[1,2,3]) # rb_build(ARRAY[3,4,5]);  -- xor
SELECT rb_build(ARRAY[1,2,3]) - rb_build(ARRAY[3,4,5]);  -- difference

SELECT rb_build(ARRAY[1,2,3]) | 9;                       -- add element
SELECT rb_build(ARRAY[1,2,3]) - 2;                       -- remove element
```

Containment and overlap operators are available:

```sql
SELECT rb_build(ARRAY[1,2,3]) @> rb_build(ARRAY[2,3]);
SELECT rb_build(ARRAY[2,3]) <@ rb_build(ARRAY[1,2,3]);
SELECT rb_build(ARRAY[1,2,3]) && rb_build(ARRAY[3,4,5]);
```

### Cardinality and Range Helpers

```sql
SELECT rb_and_cardinality(a.users32, b.users32);
SELECT rb_or_cardinality(a.users32, b.users32);
SELECT rb_xor_cardinality(a.users32, b.users32);
SELECT rb_andnot_cardinality(a.users32, b.users32);

SELECT rb_range(users32, 100, 1000);
SELECT rb_range_cardinality(users32, 100, 1000);
SELECT rb_fill(users32, 100, 200);
SELECT rb_clear(users32, 100, 200);
SELECT rb_flip(users32, 100, 200);

SELECT rb_min(users32), rb_max(users32), rb_rank(users32, 500), rb_index(users32, 500);
SELECT rb_jaccard_dist(a.users32, b.users32);
```

The 64-bit range helpers use the `rb64_` prefix. Since v1.1.0, `range_end = 0` means unlimited for several `rb64_` range/select functions.

### Aggregate Functions

```sql
SELECT rb_build_agg(user_id) FROM events;
SELECT rb_or_agg(users32) FROM cohorts;
SELECT rb_and_agg(users32) FROM cohorts;
SELECT rb_xor_agg(users32) FROM cohorts;

SELECT rb64_build_agg(user_id::bigint) FROM events;
SELECT rb64_or_agg(users64) FROM cohorts;
```

### Optimize Serialized Size

```sql
UPDATE cohorts
SET users32 = rb_runoptimize(users32);

UPDATE cohorts
SET users64 = rb64_runoptimize(users64);
```

`rb_runoptimize()` and `rb64_runoptimize()` can reduce serialized bitmap size for suitable data distributions. Measure before using them in hot write paths.

### Caveats

- Pigsty uses extension file name `roaringbitmap.md`; the upstream package name is `pg_roaringbitmap`.
- Source builds require PostgreSQL headers and CRoaring dependencies. The README notes regression testing before release covers PostgreSQL 13 and above.
- `Makefile_native` can compile with SIMD instructions and may be faster on matching CPUs, but binaries built that way can crash with `SIGILL` on machines without the required CPU features.
- Use `bytea` output for large bitmaps and `array` output for human inspection.
