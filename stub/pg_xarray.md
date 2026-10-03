## Usage

Sources:

- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_xarray/pg_xarray.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_xarray/pg_xarray.control)
- [extensions/pg_xarray/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_xarray/src/lib.rs)
- [extensions/pg_xarray/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_xarray/pgbrew.toml)
- [extensions/pg_xarray/demo/01_register_local.sql](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_xarray/demo/01_register_local.sql)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)

`pg_xarray` 0.3.0 catalogs scientific datasets, variables, chunks and meshes in `pgx`, with scientific-array queries and optional WMS serving.

### Core Workflow

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pg_xarray;
SELECT pgx.register_dataset('weather', 'zarr');
SELECT * FROM pgx.list_datasets();
```

### Operational Boundaries

Requires `postgis` and superuser installation. Dataset/file registration points at server-readable local or object-storage resources; restrict those paths and credentials. The SQL surface includes `register_dataset`, `register_file`, `register_variable`, `register_chunk` and `list_datasets`. WMS requires preload/restart and explicit `pg_xarray.wms_enabled`; it defaults off. Enabling the SQL catalog does not start a public tile service. Preserve the external array data and catalog together when backing up. This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.
