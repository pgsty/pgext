## Usage

Sources:

- [lagodb-iceberg/README.md](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/lagodb-iceberg/README.md)
- [README.md](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/README.md)
- [lagodb-iceberg/src/lib.rs](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/lagodb-iceberg/src/lib.rs)
- [Cargo.toml](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/Cargo.toml)
- [lagodb-iceberg/lagodb_iceberg.control](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/lagodb-iceberg/lagodb_iceberg.control)

`lagodb_iceberg` 0.1.0 provides the iceberg table access method for managed tables and a foreign-data wrapper for existing REST catalogs. Managed tables support SQL DML, transactions, selected schema changes and maintenance. Foreign catalog tables remain owned by their external catalog. Iceberg v1 rejects row-level update/delete; repeatable-read isolation is unsupported and serializable mode does not provide full PostgreSQL SSI semantics.

### Core Workflow

```ini
shared_preload_libraries = 'lagodb_base'
lagodb.provider_libraries = 'lagodb_iceberg,lagodb_connectors'
```

```sql
CREATE EXTENSION lagodb_base;
CREATE EXTENSION lagodb_iceberg;
CREATE TABLE local_events (event_time timestamptz, device_id bigint) USING iceberg;
INSERT INTO local_events VALUES (now(), 101);
SELECT * FROM local_events;
```

### Operational Boundaries

This is development software, not recommended upstream for production. PostgreSQL 17 is the documented target. A superuser installs the extensions; preload the base runtime and configure provider libraries before restarting. Changing provider libraries requires another restart. S3 has upstream end-to-end coverage; GCS and Azure remain experimental. Delta Lake support is only a separate skeleton and is not supplied by these extensions.
