## Usage

Sources:

- [extensions/pg_delta/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_delta/Cargo.toml)
- [extensions/pg_delta/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_delta/src/lib.rs)
- [crates/pg_bgworker/src/supervision.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/crates/pg_bgworker/src/supervision.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [extensions/pg_delta/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_delta/pgbrew.toml)
- [extensions/pg_delta/README.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_delta/README.md)
- [extensions/pg_delta/pg_delta.control](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_delta/pg_delta.control)
- [extensions/pg_delta/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_delta/src/lib.rs)
- [extensions/pg_delta/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_delta/pgbrew.toml)
- [docs/pg_delta.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/docs/pg_delta.md)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/LICENSE)

`pg_delta` 0.3.3 integrates PostgreSQL with Delta Lake through reads, exports and managed streams. The new index mode catalogs the Delta transaction log and prunes files for in-place FDW queries.

### Core Workflow

```ini
shared_preload_libraries = 'pg_delta'
```

```sql
CREATE EXTENSION pg_delta;
SELECT * FROM delta.list_tables();
SELECT delta.status();
```

### Operational Boundaries

Superuser installation is required. Preload `pg_delta` and restart for its stream manager. Configure the worker database and storage credentials deliberately. The `delta` schema exposes table/stream creation, refresh, status, history and export interfaces; index mode uses `pg_delta_server`. Access to cloud paths and SQL definitions is privileged. The manual marks logical-replication CDC export as not implemented; polling and snapshot modes have their own update/delete and recovery behavior. This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.

### Current Release and Upgrade

Extension version 0.3.3 uses `delta.database` for background workers. Create the extension in that database; a worker now waits instead of repeatedly exiting when it is absent. Changing restart-sensitive worker configuration requires a restart. From extension 0.3.0 onward, install matching files and use ALTER EXTENSION UPDATE. Earlier versions still require the migration described above. Withdrawn repository 0.4.0 bottles must be replaced by 0.5.0; upstream bottles do not imply Pigsty package availability.

### Worker Supervision

Version 0.3.3 adds `delta.worker_status()` and the superuser-only `delta.reset_workers()`. Status reports worker name, state, failures, restarts, PID, time of entry into the current state and last failure. `delta.max_worker_failures` defaults to 10 consecutive failures; 0 retries indefinitely. Restart delays back off from 5 to 60 seconds, and reaching the limit leaves a failed worker idle until reset.

Install matching libraries, restart PostgreSQL to initialize the new shared memory, and update the extension SQL in the configured worker database. Inspect the failure before resetting; a reset does not repair its cause.

```sql
ALTER EXTENSION pg_delta UPDATE;
SELECT * FROM delta.worker_status();
```
