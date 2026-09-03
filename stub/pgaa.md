## Usage

Sources:

- [Official PGAA documentation](https://www.enterprisedb.com/docs/pgaa/1.11/)
- [Official installation guide](https://www.enterprisedb.com/docs/pgaa/1.11/installing/)
- [Official quickstart](https://www.enterprisedb.com/docs/pgaa/1.11/overview/quick_start/)

`pgaa` is EDB Postgres Analytics Accelerator, a table access method that queries Parquet, Delta Lake, and Iceberg data through a vectorized Seafowl engine while remaining accessible from PostgreSQL SQL.

### Enablement

PGAA 1.11.0 supports PostgreSQL, PGE, and EPAS 16–18 on documented platforms. Install the matching package, preload it, start its managed engine, restart, and create it with dependencies:

```ini
shared_preload_libraries = 'pgaa,pgfs'
pgaa.autostart_seafowl = on
```

```sql
CREATE EXTENSION pgaa CASCADE;
```

`CASCADE` creates EDB PGFS. Verify that the installed `pgfs` is the EDB provider module, not an unrelated same-name extension.

### Register Storage and Query Data

```sql
SELECT pgfs.create_storage_location(
  'analytics',
  's3://warehouse-bucket',
  '{"region":"us-east-1"}'
);

CREATE TABLE lineitem ()
USING pgaa
WITH (
  pgaa.storage_location = 'analytics',
  pgaa.path = 'tpch/lineitem',
  pgaa.format = 'delta'
);

SELECT count(*) FROM lineitem;
```

PGAA can also use Iceberg catalogs, write supported formats with CTAS, and offload selected workloads to Spark.

### Operational Boundaries

Queries depend on object-store availability, credentials, metadata consistency, and the Seafowl/Spark execution process. Protect storage-location secrets, control network egress, monitor worker lifecycle and fallback plans, and test format/version compatibility. PostgreSQL backup alone does not capture external lake data; coordinate catalog, object storage, PGFS metadata, and database recovery as one system.
