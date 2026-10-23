## Usage

Sources:

- [Extension control](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb_gw_host/pg_documentdb_gw_host.control)
- [Cargo version and supported majors](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb_gw_host/Cargo.toml)
- [Worker initialization](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb_gw_host/src/bgworker.rs)
- [Required settings](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb_gw_host/src/gucs.rs)
- [Official README](https://github.com/documentdb/documentdb/blob/v1.0-0/README.md)

`pg_documentdb_gw_host` hosts the DocumentDB wire-protocol gateway in a PostgreSQL background worker. Its extension version remains 0.1.0 in DocumentDB 1.0-0. The source declares PostgreSQL 15–18 support and uses pgrx 0.16.0; confirm that the chosen distribution actually includes this optional module.

### Enable the Worker

Install matching DocumentDB dependencies and SQL objects as a superuser:

```sql
CREATE EXTENSION documentdb CASCADE;
CREATE EXTENSION pg_documentdb_gw_host;
```

SQL installation alone does not enable a persistent gateway. Add `pg_documentdb_gw_host` to the deployment's existing `shared_preload_libraries`, configure `documentdb_gateway.database` to the database containing the extension, and set `documentdb_gateway.setup_configuration_file` to a valid gateway configuration file on the database host. Both settings are superuser-only and have postmaster context: apply them before restarting PostgreSQL.

### Operational Boundaries

The worker connects to the configured database and starts the gateway using its setup and TLS configuration. Missing settings or invalid configuration prevent startup. Database connection failures cause the worker to shut down rather than stay idle; the registered worker has a one-second restart interval, so inspect repeated startup errors.

This module is not a separate collection API. Use DocumentDB client protocols through the configured gateway endpoint; do not expose its listener until authentication, TLS, network access and the connection budget are configured. Keep its matching base extension and gateway libraries together during upgrades.
