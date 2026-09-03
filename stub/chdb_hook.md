## Usage

Sources:

- [PGXN 0.1.0 README](https://pgxn.org/dist/chdb/0.1.0/README.html)
- [chdb_hook 0.1.0 documentation](https://pgxn.org/dist/chdb/0.1.0/doc/chdb_hook.html)
- [chdb 0.1.0 distribution metadata](https://api.pgxn.org/src/chdb/chdb-0.1.0/META.json)
- [Apache 2.0 license](https://api.pgxn.org/src/chdb/chdb-0.1.0/LICENSE.md)

`chdb_hook` is the headless module shipped by the `chdb` distribution. It intercepts URL-based `COPY` and selected `CREATE TABLE` statements so chDB can read or write local files, HTTP resources, S3, Google Cloud Storage, Azure Blob/ABFS, and HDFS formats. It has no control file and creates no SQL objects.

### Session-Scoped Workflow

```sql
LOAD 'chdb_hook';

CREATE TABLE times (
    id integer,
    months integer,
    days integer
);

COPY times
FROM 's3://datasets-documentation/my-test-bucket-768/some_prefix/some_file_1.csv';
```

Load the module explicitly for one session, through `session_preload_libraries` for selected roles or databases, or through `shared_preload_libraries` cluster-wide. Only the last form requires a PostgreSQL restart.

### Supported Paths and Privileges

The hook recognizes `file`, `http`, `https`, `s3`, `gs`, `gcs`, `oss`, `az`, `azure`, `abfss`, `abfs`, and `hdfs` URL schemes. It preserves normal relation privileges: `COPY TO` needs `SELECT`, while `COPY FROM` needs `INSERT` and a read-write transaction. Local `file` access additionally needs `pg_read_server_files` or `pg_write_server_files` plus operating-system permissions.

The `structure_from` and `copy_from` table options can infer columns or populate a new table from a URL. Credentials, format, compression, timeouts, wildcards, and explicit structures are passed into the chDB path; keep secrets out of statement logs and catalog-visible table options whenever possible.

### Boundaries

Loading `chdb_hook` changes parsing and execution of ordinary PostgreSQL commands for the session. Granting server-file roles also gives access to cloud and HTTP paths through the hook, so combine it only with trusted SQL callers and constrained network egress. The format bridge has documented NULL, array, JSON, geometry, time, Protobuf, Parquet, Arrow, and encoding edge cases; test representative round trips before migration or backup use.

Version 0.1.0 depends on the same helper and libchdb 26.7.0 or newer as `chdb`, and supports PostgreSQL 16 or newer on Linux or macOS. A helper failure aborts the initiating operation. Large imports and exports can consume substantial memory, CPU, disk, and network bandwidth, so set chDB resource limits and monitor the host.
