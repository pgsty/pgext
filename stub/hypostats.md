## Usage

Sources:

- [Pinned control file](https://github.com/lmwnshn/hypostats/blob/0181b25d84430d7d7cfc1971b1c94c270cfebbb4/hypostats.control)
- [Pinned implementation](https://github.com/lmwnshn/hypostats/blob/0181b25d84430d7d7cfc1971b1c94c270cfebbb4/src/lib.rs)

`hypostats` is an experimental pgrx extension that serializes a column's internal `pg_statistic` row to JSON text and can write that row back. Its narrow use case is testing planner behavior with copied or synthetic statistics; it directly changes a system catalog and is not a general statistics backup format.

### Core Workflow

Install it as a superuser, identify the relation OID and attribute number, and dump the current statistics:

```sql
CREATE EXTENSION hypostats;

SELECT attnum, attname
FROM pg_attribute
WHERE attrelid = 'public.events'::regclass
  AND attnum > 0
  AND NOT attisdropped;

SELECT pg_statistic_dump(
  'public.events'::regclass::oid::integer,
  2::smallint
);
```

`pg_statistic_dump(integer, smallint)` returns JSON text or `NULL` when no matching statistics tuple exists. Pass an unchanged dump to `pg_statistic_load(text)` to insert or replace the row named by the OIDs embedded in that JSON:

```sql
SELECT pg_statistic_load('<json-returned-by-pg_statistic_dump>');
```

The loader returns `true` after its catalog write. `anyarray_elemtype(anyarray)` is also exposed as a low-level helper returning an element type OID.

### Safety Boundaries

The JSON contains relation, operator, collation, and element-type OIDs. Those identifiers are cluster-local and can change after dump/restore, extension recreation, or object replacement; loading a document from another cluster or against a different relation can corrupt planner statistics or fail.

The implementation supports statistic value arrays whose element types are `float4`, `float8`, `int4`, or `int8`; other element types cause an error. It writes `pg_statistic` directly under a row-exclusive lock, bypassing `ANALYZE`'s normal computation and validation. A later `ANALYZE` may replace the loaded row.

Use only on disposable test systems, retain a recovery path, and verify plans with `EXPLAIN`. The extension is version `0.0.0`, is non-relocatable, and requires superuser installation. It does not require `shared_preload_libraries`.
