## Usage

Sources:

- [Official configuration, schema, and warning](https://github.com/asotolongo/stat_record_old/blob/4c9d43aba5168f4dcefdd843932b9c8ecf4112e7/README.md)
- [Version 0.4.0 installation SQL](https://github.com/asotolongo/stat_record_old/blob/4c9d43aba5168f4dcefdd843932b9c8ecf4112e7/stat_record--0.4.0.sql)
- [Extension control file](https://github.com/asotolongo/stat_record_old/blob/4c9d43aba5168f4dcefdd843932b9c8ecf4112e7/stat_record.control)

`stat_record` runs a background worker that periodically snapshots cluster, database, object, query, WAL, and configuration statistics into the fixed `_stat_record` schema. It provides historical reports without an external scheduler, but the upstream repository is abandoned and explicitly warns that bugs exist.

### Core Workflow

Install the extension and its `pg_stat_statements` dependency, configure both libraries for shared preload, then restart PostgreSQL:

```sql
CREATE EXTENSION stat_record CASCADE;
```

```conf
shared_preload_libraries = 'pg_stat_statements,stat_record'
stat_record.database_name = 'postgres'
stat_record.interval = 3600
stat_record.retention = 7
```

After restart, the `stat_record worker` takes snapshots at the configured interval. A manual snapshot is also available:

```sql
SELECT _stat_record.take_record();
SELECT * FROM _stat_record.lastest_records(10);
```

### Stored Data and Reports

- `_stat_record._record_number`, `_global_stat`, `_db_stat`, and `_query_stat` store snapshot identifiers and statistics.
- `_stat_record.detail_record(bigint)`, `global_report_record(bigint, bigint)`, and `total_report_record(bigint, bigint)` expose report views over snapshots.
- `_stat_record.delete_record(bigint)` removes one snapshot; `truncate_record(boolean)` removes all history and can reset numbering.
- `_stat_record.export_total_report_record(bigint, bigint, text)` writes a server-side CSV file at the supplied path.

### Safety and Compatibility

Upstream states PostgreSQL 10 or later but does not document testing against modern statistics-catalog changes. Version `0.4.0` has no declared license and carries an explicit bug warning. The worker uses a configured database, stores query text and broad server metadata, and can create server-side files. Review worker privileges, query-text sensitivity, output paths, retention and table growth, restart behavior, and PostgreSQL-major compatibility. Prefer a maintained monitoring stack for production history, and test removal or upgrade with the worker stopped.
