## Usage

Sources:

- [TencentDB for PostgreSQL supported extension versions](https://cloud.tencent.com/document/product/409/75121)
- [TencentDB for PostgreSQL 9.5 extension matrix](https://cloud.tencent.com/document/product/409/75123)

`pg_stat_error` 1.0 is a TencentDB for PostgreSQL managed extension. Tencent Cloud publishes its availability matrix but does not publish an authoritative SQL object reference or behavioral contract for the extension. Do not infer views, functions, retention, collection scope, or privilege requirements from its name.

### Availability and Installation

Check the actual TencentDB instance before creating it:

```sql
SELECT name, version, installed
FROM pg_available_extension_versions
WHERE name = 'pg_stat_error';

CREATE EXTENSION pg_stat_error;

SELECT extname, extversion
FROM pg_extension
WHERE extname = 'pg_stat_error';
```

The current provider overview lists version 1.0 for PostgreSQL 10–13 and no support for PostgreSQL 14–18. Older provider pages also list version 1 for TencentDB PostgreSQL 9.3 and 9.5 kernel releases. Availability can additionally depend on the managed kernel revision, so the instance catalog is the final check.

### Safe Discovery Boundary

After installation, use `\dx+ pg_stat_error` in `psql` or inspect `pg_depend`/`pg_proc`/`pg_class` as an administrator to identify the objects that the exact managed build installed. Obtain Tencent Cloud support documentation for their columns, reset behavior, retention, visibility, and overhead before granting access or building monitoring around them.

This extension is not a community package and cannot be reproduced by copying a library to stock PostgreSQL. The cited provider material proves availability and version only; it does not establish that the extension captures every error, replaces PostgreSQL logs, persists across failover, or exposes data to ordinary users.
