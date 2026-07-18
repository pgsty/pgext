## Usage

Sources:

- [parquet_fdw README at the reviewed commit](https://github.com/adjust/parquet_fdw/blob/d15664ebdcb1cbb759a7bb39fd26fb2fa2fff3ea/README.md)
- [parquet_fdw.control at the reviewed commit](https://github.com/adjust/parquet_fdw/blob/d15664ebdcb1cbb759a7bb39fd26fb2fa2fff3ea/parquet_fdw.control)

`parquet_fdw` is an experimental, read-only foreign data wrapper for Apache Parquet files. It uses Apache Arrow and Parquet libraries to expose one file or a set of files as a PostgreSQL foreign table.

### Read a Parquet File

```sql
CREATE EXTENSION parquet_fdw;
CREATE SERVER parquet_srv FOREIGN DATA WRAPPER parquet_fdw;
CREATE USER MAPPING FOR CURRENT_USER
  SERVER parquet_srv
  OPTIONS (user 'postgres');

CREATE FOREIGN TABLE userdata (
  id integer,
  first_name text,
  last_name text
)
SERVER parquet_srv
OPTIONS (filename '/mnt/data/userdata.parquet');

SELECT * FROM userdata LIMIT 20;
```

The `filename` option can contain space-separated paths. Advanced options include `sorted`, `files_func`, `max_open_files`, `use_mmap`, and `use_threads`; upstream also documents `IMPORT FOREIGN SCHEMA` for discovering files in a server-side directory.

### Caveats

- The wrapper is read-only and upstream labels it experimental. Paths refer to the PostgreSQL server's filesystem, not the SQL client's filesystem.
- Building the reviewed version requires system Apache Arrow/Parquet libraries; the README states version 0.15 or newer.
- Supported mappings include integer, floating-point, timestamp, date, string, binary, list, and map values. Structs and nested lists are not supported.
- The control file identifies extension version 0.2, matching the catalog version. Validate the external library and PostgreSQL combinations used by the deployment.
