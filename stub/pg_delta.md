## Usage

Sources:

- [extensions/pg_delta/README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_delta/README.md)
- [extensions/pg_delta/pg_delta.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_delta/pg_delta.control)
- [extensions/pg_delta/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_delta/src/lib.rs)
- [extensions/pg_delta/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_delta/pgbrew.toml)
- [docs/pg_delta.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/docs/pg_delta.md)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)

`pg_delta` 0.3.0 integrates PostgreSQL with Delta Lake through reads, exports and managed streams. The new index mode catalogs the Delta transaction log and prunes files for in-place FDW queries.

### Core Workflow

```conf
shared_preload_libraries = 'pg_delta'
```

```sql
CREATE EXTENSION pg_delta;
SELECT * FROM delta.list_tables();
SELECT delta.status();
```

### Operational Boundaries

Superuser installation is required. Preload `pg_delta` and restart for its stream manager. Configure the worker database and storage credentials deliberately. The `delta` schema exposes table/stream creation, refresh, status, history and export interfaces; index mode uses `pg_delta_server`. Access to cloud paths and SQL definitions is privileged. The manual marks logical-replication CDC export as not implemented; polling and snapshot modes have their own update/delete and recovery behavior. This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.
