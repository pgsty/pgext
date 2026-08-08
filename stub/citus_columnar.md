## Usage

Sources:

- [Citus v14.2.0 columnar control file](https://github.com/citusdata/citus/blob/v14.2.0/src/backend/columnar/citus_columnar.control)
- [Citus v14.2.0 columnar option helper](https://github.com/citusdata/citus/blob/v14.2.0/src/backend/columnar/sql/udfs/alter_columnar_table_set/latest.sql)
- [Citus columnar-storage documentation](https://docs.citusdata.com/en/stable/admin_guide/table_management.html#columnar-storage)
- [Citus v14.2.0 release](https://github.com/citusdata/citus/releases/tag/v14.2.0)

`citus_columnar` provides an append-oriented columnar table access method for PostgreSQL. It is shipped by the Citus 14.2 package but is a separate extension: the package release is `14.2.0`, while the extension control version is `14.2-1`. Use it for scan-heavy archival or analytical tables whose workload fits its write and feature restrictions.

### Create a Columnar Table

```sql
CREATE EXTENSION citus_columnar;

CREATE TABLE events_archive (
  event_at timestamptz NOT NULL,
  tenant_id bigint NOT NULL,
  kind text,
  payload jsonb
) USING columnar;
```

`citus_columnar` itself does not require `shared_preload_libraries`. Preloading `citus` is still required when the database also uses the distributed `citus` extension.

### Load and Query Data

Columnar storage groups rows into stripes and compresses columns in chunks. Bulk inserts in reasonably sized transactions produce better stripes than a stream of tiny transactions.

```sql
INSERT INTO events_archive
SELECT event_at, tenant_id, kind, payload
FROM events
WHERE event_at < now() - interval '90 days';

SELECT tenant_id, count(*), min(event_at), max(event_at)
FROM events_archive
GROUP BY tenant_id;
```

### Convert with the Citus Extension

When the main `citus` extension is also preloaded and installed, use its helper to convert a local or distributed table:

```sql
SELECT alter_table_set_access_method('events_archive', 'columnar');
SELECT alter_table_set_access_method('events_archive', 'heap');
```

Conversion rewrites the table. Converting to columnar drops existing indexes, so inventory dependent indexes and constraints before running it and schedule enough disk and lock time for the rewrite.

`alter_table_set_access_method()` belongs to `citus`, not to standalone `citus_columnar`. Without the main extension, create a new `USING columnar` table and copy data into it instead of assuming this helper exists.

### Tune Compression

Inspect and change table-level columnar options with the documented helpers:

```sql
SELECT alter_columnar_table_set(
  'events_archive',
  compression => 'zstd',
  compression_level => 3,
  stripe_row_limit => 150000,
  chunk_group_row_limit => 10000
);
```

New settings affect newly written stripes. Rewrite existing data when old stripes also need the new layout.

### Operational Boundaries

- Columnar tables are intended for append-heavy use. `UPDATE` and `DELETE` are not supported, and space left by rolled-back writes is not reclaimed through ordinary heap-style maintenance.
- TOAST is not available; large values remain inline and can hit PostgreSQL's row-size limits.
- Row locks, `AFTER ... FOR EACH ROW` triggers, serializable isolation, logical decoding, foreign keys, unlogged tables, and several scan types are unsupported. Check the current upstream limitation list before adopting the access method.
- Ordinary heap assumptions about indexes, vacuum, replication, triggers, and constraints do not automatically apply. Validate every required database feature against a representative columnar table.
- The extension installs in `pg_catalog`, is not relocatable, and has SQL version `14.2-1`; use that version when checking or updating `pg_extension`, not the package version `14.2.0`.
