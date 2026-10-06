## Usage

Sources:

- [README.md](https://github.com/matteuccimarco/pg-slim/blob/1ec4df3a35c395d9d0d7ba25c963775fbfa80088/README.md)
- [sql/pg_slim--0.1.0.sql](https://github.com/matteuccimarco/pg-slim/blob/1ec4df3a35c395d9d0d7ba25c963775fbfa80088/sql/pg_slim--0.1.0.sql)
- [pg_slim.control](https://github.com/matteuccimarco/pg-slim/blob/1ec4df3a35c395d9d0d7ba25c963775fbfa80088/pg_slim.control)

`pg_slim` 0.1.0 adds the `slim` type, JSONB conversions, field access, containment and comparison operators. It is an early storage-format implementation, not a guaranteed space-saving replacement for JSONB.

### Core Workflow

```sql
CREATE EXTENSION pg_slim;
SELECT slim_decode(slim_encode('{"name":"Alice","age":30}'::jsonb));
SELECT slim_encode('{"name":"Alice"}'::jsonb)->>'name';
```

### Operational Boundaries

`slim_encode` and `slim_decode` convert JSONB and slim values. `slim_typeof`, `slim_array_length`, `slim_object_keys` and `slim_pretty` inspect values. B-tree and hash operator classes are supplied; a GIN operator class is not implemented. Some operations decode to JSONB internally and updates replace the entire value.

Install as a superuser. The source declares no preload requirement and no complete major-version support matrix; the installation example uses PostgreSQL 15. Measure actual storage and query costs and validate round trips before moving application data.
