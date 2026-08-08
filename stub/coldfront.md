## Usage

Sources:

- [ColdFront v1.0.0-beta1 README](https://github.com/pgEdge/coldfront/blob/v1.0.0-beta1/README.md)
- [ColdFront v1.0.0-beta1 release](https://github.com/pgEdge/coldfront/releases/tag/v1.0.0-beta1)
- [ColdFront 1.0 control file](https://github.com/pgEdge/coldfront/blob/v1.0.0-beta1/extension/coldfront/coldfront.control)
- [ColdFront 1.0 extension SQL](https://github.com/pgEdge/coldfront/blob/v1.0.0-beta1/extension/coldfront/coldfront--1.0.sql)
- [ColdFront usage guide](https://github.com/pgEdge/coldfront/blob/v1.0.0-beta1/docs/usage.md)
- [ColdFront architecture](https://github.com/pgEdge/coldfront/blob/v1.0.0-beta1/docs/architecture.md)
- [ColdFront tiered-mode architecture](https://github.com/pgEdge/coldfront/blob/v1.0.0-beta1/docs/architecture_tiered.md)

`coldfront` transparently presents PostgreSQL hot partitions and writable Apache Iceberg cold data as one SQL relation. Use it to evaluate time-based hot/cold tiering, or an Iceberg-only table behind a PostgreSQL view. Release v1.0.0-beta1 is a public beta: upstream explicitly says not to use it in production because interfaces, on-disk formats, behavior, and data safety can still change.

### Configure the Runtime

ColdFront depends on `pg_duckdb` for in-process Iceberg I/O. Both libraries must be preloaded, in the documented order, and changing this setting requires a PostgreSQL restart:

```conf
shared_preload_libraries = 'pg_duckdb,coldfront'
coldfront.warehouse = 'wh'
coldfront.lakekeeper_endpoint = 'http://lakekeeper:8181/catalog'
```

The v1.0.0-beta1 stack is not an arbitrary stock DuckDB combination: its release documentation pins `pg_duckdb` with DuckDB 1.5.3 from PR 1025 plus a patched `duckdb-iceberg`. Use the release's matching build or image so Iceberg commits and strict-reader interoperability have the expected patches.

After restarting, create the extensions in each database in the same dependency order, then store the object-storage credential:

```sql
CREATE EXTENSION pg_duckdb;
CREATE EXTENSION coldfront;

SELECT coldfront.set_storage_secret(
  'access-key',
  'secret-key',
  'minio.example.com:9000',
  'us-east-1',
  'path',
  true
);
```

Lakekeeper and the object store are external services. Bootstrap the Iceberg REST catalog and warehouse before using a managed relation. For cloud AWS S3, pass `NULL` as the endpoint and the real bucket region so DuckDB uses virtual-hosted HTTPS addressing. Do not embed production secrets in migration files or examples.

### Run a Tiered Table

Tiered mode starts from a native range-partitioned PostgreSQL table. Its primary key must cover the partition key so the archiver can capture concurrent changes safely:

```sql
CREATE TABLE events (
  id bigint GENERATED ALWAYS AS IDENTITY,
  ts timestamptz NOT NULL,
  status text,
  payload jsonb,
  PRIMARY KEY (id, ts)
) PARTITION BY RANGE (ts);
```

Register and reconcile the lifecycle with the separate Go `archiver` executable:

```bash
archiver register --config /etc/coldfront/config.yaml \
  --table events --period monthly \
  --hot-period "1 month" --retention "5 years"
archiver --config /etc/coldfront/config.yaml
```

These are operating-system commands, not extension SQL APIs. Run `archiver` from cron or a service timer and alert on failures. It pre-creates partitions, converts the original table into a hot heap such as `_events` plus a transparent `events` view, moves partitions older than `hot_period` to Iceberg, advances `coldfront.archive_watermark`, and permanently expires data older than `retention_period`. `retention_period` must exceed `hot_period`.

Applications continue to use the public relation:

```sql
SELECT id, ts, status FROM events
WHERE ts >= '2026-07-01'
ORDER BY ts DESC;

SET coldfront.allow_mixed_writes = off;
UPDATE events
SET status = 'fixed'
WHERE ts >= '2026-07-01' AND id = 42;
```

Setting `coldfront.allow_mixed_writes = off` makes an ambiguous `UPDATE` or `DELETE` fail instead of writing both tiers. This is safer when the application can provide a partition-key predicate. The default is permissive and a dual-tier commit is not crash-safe.

### Create an Iceberg-Only Table

Decoupled mode stores every row in Iceberg and does not use `archiver` or a PostgreSQL hot heap. Pre-create the Iceberg namespace in Lakekeeper, then let the extension create the external table, wrapper view, and registry row:

```sql
SELECT coldfront.create_iceberg_table(
  p_schema  => 'public',
  p_table   => 'events_archive',
  p_columns => '[
    {"name":"id", "type":"bigint"},
    {"name":"ts", "type":"timestamptz"},
    {"name":"status", "type":"text"},
    {"name":"payload", "type":"jsonb"}
  ]'::jsonb
);

INSERT INTO events_archive VALUES
  (1, now(), 'new', '{"source":"demo"}');
SELECT * FROM events_archive;
```

`SELECT`, `INSERT`, `UPDATE`, and `DELETE` target the wrapper relation; the hook rewrites data-modifying statements to Iceberg. PostgreSQL is only the SQL and compute front end for this table, so the rows do not become part of its heap storage or ordinary PostgreSQL backup.

### Important Objects

- `coldfront.set_storage_secret(...)` records an S3 or S3-compatible credential and materializes a persistent DuckDB secret. `coldfront.set_storage_secret_azure(...)` is the Azure ADLS Gen2 counterpart.
- `coldfront.create_iceberg_table(...)` provisions an Iceberg-only table and its PostgreSQL wrapper. It is not the tiered-table registration path.
- `coldfront.grant_app_access(regrole)` grants a non-superuser the registry-derived runtime privileges for managed views; it deliberately does not grant administrative functions or server-file roles.
- `coldfront.tiered_views` registers transparent relations, `coldfront.archive_watermark` records the hot/cold cutoff, and `coldfront.partition_config` stores per-table lifecycle policy.
- `coldfront.warehouse` and `coldfront.lakekeeper_endpoint` select the external catalog. `coldfront.allow_mixed_writes` controls ambiguous cross-tier DML, and `coldfront.local_pg_dsn` enables the documented PostgreSQL-to-Iceberg streaming path.
- `archiver`, `partitioner`, and `compactor` are separate Go programs. `archiver` is required for tiered movement; `partitioner` can manage PostgreSQL partitions without a cold tier; `compactor` performs Iceberg maintenance. Installing `coldfront` does not schedule any of them.

### DDL and DML Boundaries

- A tiered application's relation is a view and its native hot table is typically `_events`. Do not bypass the view for ordinary writes or assume direct hot-table changes also affect Iceberg.
- The DDL hook mirrors a documented subset issued against the registered hot table: `ALTER TABLE _events ADD/DROP COLUMN`, safe `ALTER COLUMN ... TYPE` promotions, column rename, and table/view rename. Unsupported types or conversions are rejected; test every schema migration against a copy of both tiers.
- `DROP TABLE _events`, dropping the managed view, and `TRUNCATE _events` are blocked because a one-sided operation would orphan or expose cold rows. First unregister the lifecycle, then dismantle the PostgreSQL and Iceberg sides deliberately.
- DML that touches the cold tier does not support `RETURNING`. Tiered self-joins, `DELETE ... USING` the same managed view, and subqueries that reference that view a second time are rejected.
- A normal `ROLLBACK` coordinates PostgreSQL and DuckDB transactions, but a backend crash between the Iceberg snapshot operation and PostgreSQL commit can leave orphaned object-store files. Strict single-tier predicates reduce the dual-tier risk; Iceberg snapshot expiry and orphan-file maintenance remain operational requirements.
- Repeatable-read isolation does not extend across multiple Iceberg scans in a long PostgreSQL transaction. Reads can observe concurrent external commits between scans even though read-your-own-write works.

### Backup and Recovery

Logical `pg_dump` includes `coldfront.tiered_views`, `coldfront.archive_watermark`, and `coldfront.partition_config`, but it does not copy Iceberg data, the Lakekeeper catalog database, object-storage credentials in `coldfront.storage_secret`, or transient bakery claim tables. After restoring PostgreSQL, re-run `coldfront.set_storage_secret(...)`; the restored metadata then points to the same external Iceberg tables.

Back up and recover the PostgreSQL hot tier, Lakekeeper catalog database, and object-store bucket as separate but coordinated systems. Preserve compatible Iceberg metadata and objects at the same recovery point; restoring only PostgreSQL is not a cold-data restore, and deleting or recreating the external warehouse can make the restored registry unusable.

The tiered retention path eventually destroys data after exporting it. Back up before shortening `retention_period`, rehearse recovery while the external services are unavailable, and keep the object-store lifecycle policy from deleting files still referenced by Iceberg snapshots.

### Version and Compatibility Boundaries

- Release v1.0.0-beta1 ships SQL extension version `1.0`; the control file fixes schema `coldfront`, is not relocatable, and requires preloading plus a server restart.
- The tagged release documents stock PostgreSQL 16, 17, and 18. Use the exact dependency matrix validated by the release rather than assuming another `pg_duckdb`, DuckDB, or `duckdb-iceberg` build is compatible.
- Supported Iceberg-backed columns include common integer, floating-point, Boolean, temporal, UUID, text, bounded numeric, `bytea`, `json`, `jsonb`, and `interval` types. Arrays, enums, composites, ranges, unbounded numeric, `inet`, and `cidr` are rejected.
- `jsonb` and `json` are stored as strings in Iceberg and surface through the managed view as `json`; JSONB-only operators such as `?` and `@>` require an explicit cast back to `jsonb`.
- Public beta status is the dominant boundary: do not treat successful evaluation, an upstream benchmark, or an available source build as production-readiness evidence.
