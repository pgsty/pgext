## Usage

Sources:

- [extensions/pg_s3/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_s3/Cargo.toml)
- [extensions/pg_s3/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_s3/src/lib.rs)
- [crates/pg_bgworker/src/supervision.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/crates/pg_bgworker/src/supervision.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [extensions/pg_s3/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_s3/pgbrew.toml)
- [extensions/pg_s3/pg_s3.control](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_s3/pg_s3.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/README.md)
- [extensions/pg_s3/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_s3/Cargo.toml)
- [extensions/pg_s3/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_s3/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/LICENSE)
- [extensions/pg_s3/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_s3/pgbrew.toml)

`pg_s3` 0.3.2 serves an experimental S3-compatible HTTP API while storing bucket/object metadata in `pgs3` and binary payloads on local disk.

### Core Workflow

```ini
shared_preload_libraries = 'pg_s3'
pg_s3.database = 'postgres'
pg_s3.host = '127.0.0.1'
pg_s3.port = 9100
```

```sql
CREATE EXTENSION pg_s3;
SELECT pgs3.create_bucket('sample-bucket');
SELECT * FROM pgs3.list_buckets();
```

### Operational Boundaries

The control requires superuser installation. Preload and restart. Configure `pg_s3.database`, `pg_s3.host`, `pg_s3.port` and `pg_s3.data_directory`; defaults include all interfaces, port 9100 and the postgres database. The SQL API manages buckets, object metadata and listing. Backups must cover both database metadata and payload files consistently. Restrict filesystem paths and network access; source availability does not establish complete S3 API, authentication or multi-node durability compatibility. This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.

### Current Release and Upgrade

Extension version 0.3.2 uses `pg_s3.database` for background workers. Create the extension in that database; a worker now waits instead of repeatedly exiting when it is absent. Changing restart-sensitive worker configuration requires a restart. From extension 0.3.0 onward, install matching files and use ALTER EXTENSION UPDATE. Earlier versions still require the migration described above. Withdrawn repository 0.4.0 bottles must be replaced by 0.5.0; upstream bottles do not imply Pigsty package availability.

### Worker Supervision

Version 0.3.2 adds `pgs3.worker_status()` and the superuser-only `pgs3.reset_workers()`. Status reports worker name, state, failures, restarts, PID, time of entry into the current state and last failure. `pg_s3.max_worker_failures` defaults to 10 consecutive failures; 0 retries indefinitely. Restart delays back off from 5 to 60 seconds, and reaching the limit leaves a failed worker idle until reset.

Install matching libraries, restart PostgreSQL to initialize the new shared memory, and update the extension SQL in the configured worker database. Inspect the failure before resetting; a reset does not repair its cause.

```sql
ALTER EXTENSION pg_s3 UPDATE;
SELECT * FROM pgs3.worker_status();
```
