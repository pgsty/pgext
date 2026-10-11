## Usage

Sources:

- [lagodb-base/README.md](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/lagodb-base/README.md)
- [README.md](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/README.md)
- [lagodb-base/src/lib.rs](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/lagodb-base/src/lib.rs)
- [Cargo.toml](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/Cargo.toml)
- [lagodb-base/lagodb_base.control](https://github.com/lagodb/lagodb/blob/3a7c7ccffe3b9d3b04f458830d767aeaa703a4a5/lagodb-base/lagodb_base.control)

`lagodb_base` 0.1.0 hosts LagoDB worker supervision, storage volumes, provider loading and query offload. It is the shared dependency for the connector and Iceberg extensions. It routes supported object-URI COPY operations and keeps ordinary PostgreSQL query paths available.

### Core Workflow

```ini
shared_preload_libraries = 'lagodb_base'
lagodb.provider_libraries = 'lagodb_iceberg,lagodb_connectors'
```

```sql
CREATE EXTENSION lagodb_base;
```

### Operational Boundaries

This is development software, not recommended upstream for production. PostgreSQL 17 is the documented target. A superuser installs the extensions; preload the base runtime and configure provider libraries before restarting. Changing provider libraries requires another restart. S3 has upstream end-to-end coverage; GCS and Azure remain experimental. Delta Lake support is only a separate skeleton and is not supplied by these extensions.
