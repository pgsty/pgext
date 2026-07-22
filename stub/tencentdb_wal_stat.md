## Usage

Sources:

- [Official TencentDB WAL analysis documentation](https://cloud.tencent.com/document/product/409/130512)
- [Official TencentDB for PostgreSQL product page](https://cloud.tencent.com/product/postgres)

`tencentdb_wal_stat` is a TencentDB for PostgreSQL diagnostic extension that reads recent WAL files and attributes record counts, data bytes, full-page images, resource-manager types, and relation identifiers. Use it to identify databases, schemas, tables, or indexes associated with a WAL surge; it is not a long-term WAL history store.

### Requirements and Core Workflow

Tencent documents version `1.0` for TencentDB PostgreSQL 15 or later, executed on the primary by a member of `pg_tencentdb_superuser`:

```sql
CREATE EXTENSION IF NOT EXISTS tencentdb_wal_stat;

SELECT extname, extversion
FROM pg_extension
WHERE extname = 'tencentdb_wal_stat';
```

Pass the number of recent WAL files to inspect, from `1` through `500`. This query finds the largest table/index contributors in the last 50 available files:

```sql
SELECT database_name,
       schema_name,
       relation_name,
       relation_kind,
       rmgr_name,
       wal_records,
       pg_size_pretty((wal_bytes + wal_fpi_bytes)::bigint) AS total_size
FROM tencentdb_wal_stat(50)
WHERE relation_kind IN ('table', 'index')
ORDER BY (wal_bytes + wal_fpi_bytes) DESC
LIMIT 20;
```

Aggregate `wal_bytes + wal_fpi_bytes` for total attributed size. Group by `database_name`, `schema_name`, or `rmgr_name` to move from instance-level triage to a database, schema, or write type.

### Output and Interpretation

- `database_oid`, `schema_oid`, and `relation_oid` identify objects; companion name columns are resolved where catalogs are available.
- `relation_kind` distinguishes tables, indexes, TOAST, sequences, materialized views, shared/system objects, and unknown objects.
- `rmgr_name` classifies WAL by resource manager, such as `Heap`, `Btree`, or `Transaction`.
- `wal_records` and `wal_bytes` count records and non-FPI bytes; `wal_fpi` and `wal_fpi_bytes` count full-page images and their bytes.

### Operational Boundaries

- The function is available only on the primary and requires the provider role. It is a TencentDB component; do not expect the package or role on community PostgreSQL.
- Larger `wal_num` values read more WAL and take longer. Tencent recommends using large ranges off peak; start small for incident triage and monitor the query itself.
- Only WAL files that have not been recycled can be analyzed. The argument is a file count, not a time interval, so traffic volume and WAL segment size affect the covered period.
- Names for objects in other databases may appear as `<oid:N>` or `<filenode:N>` because their catalogs cannot be resolved from the current database. Preserve OIDs and join names from the correct database rather than treating placeholders as missing WAL.
- Attribution is diagnostic evidence, not a causal proof. Correlate it with workload, checkpoints, replication slots/lag, archiving, disk use, and query statistics before changing indexes, throttling tenants, or tuning `max_wal_size`.
