## Usage

Sources:

- [Upstream configuration, schema, and warning](https://github.com/asotolongo/stat_record_old/blob/4c9d43aba5168f4dcefdd843932b9c8ecf4112e7/README.md)
- [Extension control file](https://github.com/asotolongo/stat_record_old/blob/4c9d43aba5168f4dcefdd843932b9c8ecf4112e7/stat_record.control)

`stat_record` runs a background worker that periodically snapshots cluster, database, object, query, WAL, and configuration statistics into the fixed `_stat_record` schema. It depends on `pg_stat_statements` and can generate historical reports.

```conf
shared_preload_libraries = 'pg_stat_statements,stat_record'
stat_record.database_name = 'postgres'
stat_record.interval = 3600
stat_record.retention = 7
```

After restart, install with `CREATE EXTENSION stat_record CASCADE`. The repository is abandoned, has no declared license, and upstream explicitly warns that bugs exist. The worker connects with powerful defaults, stores query text and broad metadata, and exposes server-side CSV export. Review privileges and sensitive-data retention, constrain paths, monitor table growth, and prefer a maintained monitoring stack.
