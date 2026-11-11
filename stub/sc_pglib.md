## Usage

Sources:

- [Usage and operational boundaries](https://github.com/roelofrossouw/simply-cpp-pglib/blob/5b437c79bf5cda145e85c336a5c156f9b6b1939a/README.md)
- [Extension control](https://github.com/roelofrossouw/simply-cpp-pglib/blob/5b437c79bf5cda145e85c336a5c156f9b6b1939a/sql/sc_pglib.control)
- [Initial SQL](https://github.com/roelofrossouw/simply-cpp-pglib/blob/5b437c79bf5cda145e85c336a5c156f9b6b1939a/sql/sc_pglib--1.0.sql)
- [Upgrade to SQL version 1.18](https://github.com/roelofrossouw/simply-cpp-pglib/blob/5b437c79bf5cda145e85c336a5c156f9b6b1939a/sql/sc_pglib--1.17--1.18.sql)

`sc_pglib` exposes the simply-cpp libraries to PostgreSQL. SQL/control version `1.18` is separate from the project's package version. Its API includes Base64 conversion, spreadsheet and CSV operations, server files, Redis, Firebase and application-specific integrations. The documented Ubuntu package targets are PostgreSQL 16 and 18; each library must match its server major.

### Enable and use

After installing the matching library and SQL files, enable the extension as a superuser:

```sql
CREATE EXTENSION sc_pglib;

SELECT sc_base64_encode('Hello World!');
SELECT convert_from(sc_base64_decode('SGVsbG8gV29ybGQh'), 'UTF8');
```

These Base64 functions are immutable, strict and parallel safe. Invalid encoded input raises an error. Existing databases use `ALTER EXTENSION sc_pglib UPDATE` after installing the corresponding files; the source supplies an initial script and a sequence of upgrade scripts. The extension is relocatable and does not require shared preload.

### Files and external services

`sc_xlsx` writes query results into a workbook on the database server:

```sql
SELECT sc_xlsx('/var/www/data/output/result.xlsx',
               'SELECT 1 AS sample', 'Results');
```

File operations use the database server's OS identity. Ordinary callers are limited to upstream's configured or compiled-in directories; `pg_read_server_files` and `pg_write_server_files` permit broader access. Review replacement behavior and grant only the functions an application needs. File and external-service side effects are not undone by a database transaction rollback.

Redis uses a foreign server with the `sc_redis` wrapper and optional user mappings. Firebase uses `sc_firebase` servers and service-account user mappings. The APIs perform external operations and are not a substitute for transactional PostgreSQL tables. Keep credentials in protected user mappings or the documented server configuration, not in SQL examples or public logs.

### Sensitive operations

`sc_read_text`, `sc_write_text`, `sc_read_dsn` and `sc_http_get` have `PUBLIC` execution revoked. HTTP access is limited by the superuser-configured `sc_pglib.http_urls` allowlist and does not follow redirects. The `sc_oneweb` integration uses a server-side configuration file, selected by `sc_pglib.oneweb_config`; its privileged operations also have public execution revoked. Inspect the version-specific upstream documentation before exposing these functions to application roles.
