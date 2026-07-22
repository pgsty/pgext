## Usage

Sources:

- [Official README](https://github.com/2ndQuadrant/pglog/blob/5d57809e4e73913a06b4aaa6553f53f43d010a39/README.asciidoc)
- [Official control file](https://github.com/2ndQuadrant/pglog/blob/5d57809e4e73913a06b4aaa6553f53f43d010a39/pglog.control)
- [Official installation SQL](https://github.com/2ndQuadrant/pglog/blob/5d57809e4e73913a06b4aaa6553f53f43d010a39/pglog--1.0.sql)
- [Official logger implementation](https://github.com/2ndQuadrant/pglog/blob/5d57809e4e73913a06b4aaa6553f53f43d010a39/pglog.c)
- [Official spool reader](https://github.com/2ndQuadrant/pglog/blob/5d57809e4e73913a06b4aaa6553f53f43d010a39/pglog_spool.c)

`pglog` is a PostgreSQL 9.3-era proof of concept that captures messages through an error hook, appends them to its own spool files, and exposes those files through a read-only foreign table. It is separate from PostgreSQL's normal logging collector and should not replace supported log shipping or observability pipelines.

### Core Workflow

Preload the library, choose a server-local spool directory, and restart PostgreSQL before installing the SQL objects.

```conf
shared_preload_libraries = 'pglog'
pglog.directory = 'pglog_spool'
pglog.min_messages = 'WARNING'
pglog.rotation_age = '1d'
```

```sql
CREATE EXTENSION pglog;

SELECT log_time, database_name, error_severity, message
FROM pglog
ORDER BY log_time DESC
LIMIT 100;
```

The `pglog` foreign table uses a CSV-log-shaped schema containing connection, statement, error, and application fields. Queries scan spool files from the configured directory; the wrapper does not push filters down or collect foreign-table statistics.

### Configuration and Objects

- `pglog.directory` is relative to `PGDATA` unless an absolute path is supplied; its default is `pglog_spool`.
- `pglog.min_messages` sets the minimum captured severity and defaults to `WARNING`.
- `pglog.rotation_age` controls time-based file rotation in minutes; `0` disables rotation.
- `pglog_server` and the `pglog` foreign table are created by the extension.
- `pglog_severity` is the enum used for severity values.

### Operational Boundaries

The reviewed `1.0` source does not serialize append operations across PostgreSQL backends, so concurrent messages can interleave or corrupt records. Rotation creates new files but does not remove old ones. The reader opens at most 16 directory entries in filesystem order and does not isolate the current database.

Log rows can contain statements, details, user names, host data, and application text from multiple databases. Grant access narrowly and protect the spool directory as sensitive data. Because the code targets obsolete PostgreSQL APIs and has no current compatibility statement, validate it only on an isolated test build; prefer maintained PostgreSQL logging and collection mechanisms for operational use.
