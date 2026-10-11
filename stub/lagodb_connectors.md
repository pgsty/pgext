## Usage

Sources:

- [lagodb-connectors/README.md](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/lagodb-connectors/README.md)
- [README.md](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/README.md)
- [lagodb-connectors/src/lib.rs](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/lagodb-connectors/src/lib.rs)
- [Cargo.toml](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/Cargo.toml)
- [lagodb-connectors/lagodb_connectors.control](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/lagodb-connectors/lagodb_connectors.control)

`lagodb_connectors` 0.1.0 exposes text, CSV, JSON, Avro and Parquet files through foreign tables and COPY. Foreign servers select storage providers and user mappings hold credentials. Exact-object tables are read-only; prefix tables support append-only inserts, not row updates or deletes. Local native COPY uses PostgreSQL server-file privileges.

### Core Workflow

```ini
shared_preload_libraries = 'lagodb_base'
lagodb.provider_libraries = 'lagodb_iceberg,lagodb_connectors'
```

```sql
CREATE EXTENSION lagodb_base;
CREATE EXTENSION lagodb_connectors;
CREATE SERVER local_files FOREIGN DATA WRAPPER lagodb_connectors;
CREATE FOREIGN TABLE local_events ()
SERVER local_files OPTIONS (path '/tmp/events.parquet');
SELECT * FROM local_events;
```

### Operational Boundaries

This is development software, not recommended upstream for production. PostgreSQL 17 is the documented target. A superuser installs the extensions; preload the base runtime and configure provider libraries before restarting. Changing provider libraries requires another restart. S3 has upstream end-to-end coverage; GCS and Azure remain experimental. Delta Lake support is only a separate skeleton and is not supplied by these extensions.
