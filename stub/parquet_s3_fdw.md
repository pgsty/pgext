## Usage

Sources:

- [Official README](https://github.com/pgspider/parquet_s3_fdw/blob/cf4384dba7d1f1b6252d5357311df17f5eeb1775/README.md)
- [Official FDW callbacks](https://github.com/pgspider/parquet_s3_fdw/blob/cf4384dba7d1f1b6252d5357311df17f5eeb1775/src/parquet_fdw.c)
- [Official extension control file](https://github.com/pgspider/parquet_s3_fdw/blob/cf4384dba7d1f1b6252d5357311df17f5eeb1775/parquet_s3_fdw.control)

`parquet_s3_fdw` maps Parquet files on the PostgreSQL server, Amazon S3, or a MinIO-compatible endpoint as foreign tables. It supports typed and schemaless scans, row-group pruning, parallel/multifile readers, import helpers, and file-rewriting `INSERT`, `UPDATE`, and `DELETE` operations.

### Core Workflow

Create a server, store S3 credentials in a user mapping, and point a foreign table at one or more files:

```sql
CREATE EXTENSION parquet_s3_fdw;

CREATE SERVER parquet_s3_srv
  FOREIGN DATA WRAPPER parquet_s3_fdw;

CREATE USER MAPPING FOR app_reader
  SERVER parquet_s3_srv
  OPTIONS (user '<access-key>', password '<secret-key>');

CREATE FOREIGN TABLE user_events (
  user_id bigint OPTIONS (key 'true'),
  event_name text,
  observed_at timestamp
)
SERVER parquet_s3_srv
OPTIONS (
  filename 's3://analytics/events.parquet',
  region 'ap-northeast-1'
);

SELECT * FROM user_events WHERE user_id = 42;
```

For MinIO, set server option `use_minio` and configure the documented `endpoint`. Local paths are resolved on the database server under the PostgreSQL operating-system account.

### Important Options and Features

- `filename` selects a space-separated file list; `dirname` selects a directory. A single foreign table cannot mix local and S3 paths.
- `sorted`, `files_in_order`, `max_open_files`, `use_mmap`, and `use_threads` select ordered, cached, memory-mapped, or parallel decoding strategies.
- `files_func` dynamically returns file paths; `insert_file_selector` chooses the target file for new rows. Treat both functions as privileged routing logic.
- `schemaless = 'true'` represents each row in one `jsonb` column. Typed mode maps documented Arrow primitives plus arrays and maps; structs and nested lists are unsupported.
- `IMPORT FOREIGN SCHEMA`, `import_parquet_s3`, and `import_parquet_s3_explicit` can derive foreign tables from files.

### Modification and Transaction Boundaries

- Upstream explicitly states that transactions are not supported. A PostgreSQL rollback cannot be assumed to undo a Parquet/S3 modification; never combine these writes with transactional business updates expecting atomicity.
- Updates and deletes rebuild cached data and overwrite the old Parquet file. Large files are slow, and concurrent modification of the same file can produce inconsistent results.
- Marked `key` columns identify rows for update/delete but the FDW does not enforce their uniqueness. Validate uniqueness externally and avoid concurrent writers.
- `WITH CHECK OPTION`, `ON CONFLICT`, and `RETURNING` are not supported. Compression is limited to `zstd`, `snappy`, or uncompressed output with the documented restrictions.
- Credentials stored in user mappings are catalog secrets; restrict server and mapping visibility, use narrow object-store policies and TLS, rotate keys, and audit endpoint/region settings.
- Pin the required Arrow/Parquet and AWS SDK versions and compiler ABI. Test schema evolution, time zones, nulls, numeric precision, map/list conversion, network failure, partial overwrite, connection caching, and restore procedures before production use.
