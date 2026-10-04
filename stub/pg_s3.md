## Usage

Sources:

- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/CHANGELOG.md)
- [extensions/pg_s3/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_s3/pgbrew.toml)
- [extensions/pg_s3/pg_s3.control](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_s3/pg_s3.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/README.md)
- [extensions/pg_s3/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_s3/Cargo.toml)
- [extensions/pg_s3/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_s3/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/LICENSE)
- [extensions/pg_s3/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_s3/pgbrew.toml)

`pg_s3` 0.3.1 serves an experimental S3-compatible HTTP API while storing bucket/object metadata in `pgs3` and binary payloads on local disk.

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

Extension version 0.3.1 uses `pg_s3.database` for background workers. Create the extension in that database; a worker now waits instead of repeatedly exiting when it is absent. Changing restart-sensitive worker configuration requires a restart. From extension 0.3.0 onward, install matching files and use ALTER EXTENSION UPDATE. Earlier versions still require the migration described above. Withdrawn repository 0.4.0 bottles must be replaced by 0.4.1; upstream bottles do not imply Pigsty package availability.
