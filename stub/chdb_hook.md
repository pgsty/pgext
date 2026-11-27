## Usage

Sources:

- [doc/chdb_hook.md](https://api.pgxn.org/src/chdb/chdb-0.1.4/doc/chdb_hook.md)
- [chdb_hook.control](https://api.pgxn.org/src/chdb/chdb-0.1.4/chdb_hook.control)
- [sql/chdb_hook.sql](https://api.pgxn.org/src/chdb/chdb-0.1.4/sql/chdb_hook.sql)
- [Makefile](https://api.pgxn.org/src/chdb/chdb-0.1.4/Makefile)
- [CHANGELOG.md](https://api.pgxn.org/src/chdb/chdb-0.1.4/CHANGELOG.md)
- [LICENSE.md](https://api.pgxn.org/src/chdb/chdb-0.1.4/LICENSE.md)

`chdb_hook` intercepts URL-based `COPY` and extended `CREATE TABLE` commands to move data through chDB. Distribution 0.1.4 adds a control file and SQL version 0.1 with `chdb_hook_version()`. Direct loading remains supported, so creating the extension is optional for the hook workflow.

### Load and Use

A superuser can enable it for one session:

```sql
LOAD 'chdb_hook';
CREATE TABLE times (id integer, months integer, days integer);
COPY times
FROM 's3://datasets-documentation/my-test-bucket-768/some_prefix/some_file_1.csv';
```

The new optional SQL wrapper is installed with:

`CREATE EXTENSION chdb_hook` / `SELECT chdb_hook_version()`

`CREATE EXTENSION` loads the module only in its current session. Subsequent sessions must use `LOAD` or call `chdb_hook_version()`, unless configured for preload. `session_preload_libraries` can target selected databases or roles. `shared_preload_libraries` enables it cluster-wide and requires a server restart. The wrapper is untrusted and relocatable.

### Sources and Permissions

Supported URL schemes include `file`, `http`, `https`, `s3`, `gs`, `gcs`, `oss`, `az`, `azure`, `abfss`, `abfs` and `hdfs`. Normal relation permissions apply: `COPY TO` requires `SELECT`, while `COPY FROM` requires `INSERT` and a read-write transaction. Local files also require `pg_read_server_files` or `pg_write_server_files` plus operating-system permissions. A build can disable local URLs with `NO_FILE_SCHEME`.

`structure_from` and `copy_from` table options support schema inference and data loading. Format, structure, compression, timeout and credentials are explicit COPY options. `encoding_check` accepts `fail`, `replace`, `remove` or `truncate`; URL-based `COPY FROM` with `WHERE` is rejected.

### Host Credentials and Resource Limits

`chdb_hook.host_credentials_roles` is a superuser-controlled list of existing roles and defaults empty. Only their members can use host AWS credentials when no explicit keys are supplied. The S3 `role_arn` option can assume an IAM role using those credentials; it requires the same membership and applies only without `access_key`.

`chdb_hook.max_memory` now defaults to 4 GB; `0` permits unlimited usage. `chdb_hook.max_threads` and `chdb_hook.max_parsing_threads` control concurrency. These settings need superuser privileges. Use explicit limits for large imports.

### Maintenance Boundaries

The module supports PostgreSQL 15 and later on Linux or macOS and needs the matching helper/libchdb setup. Declared timestamp precision and session `TimeZone` are preserved in supported conversions. Verify complex-type and encoding round trips; keep credentials out of SQL logs, especially debugging output. Existing direct-load installations can continue without SQL objects; the version function is available only after installing the new wrapper.
