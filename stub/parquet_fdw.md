## Usage

Sources:

- [Official upstream documentation](https://github.com/adjust/parquet_fdw/blob/d15664ebdcb1cbb759a7bb39fd26fb2fa2fff3ea/README.md)
- [Official extension control file](https://github.com/adjust/parquet_fdw/blob/d15664ebdcb1cbb759a7bb39fd26fb2fa2fff3ea/parquet_fdw.control)
- [Official option and import implementation](https://github.com/adjust/parquet_fdw/blob/d15664ebdcb1cbb759a7bb39fd26fb2fa2fff3ea/src/parquet_impl.cpp)

`parquet_fdw` 0.2 is an experimental, read-only foreign data wrapper for local Apache Parquet files. It uses Apache Arrow and Parquet C++ libraries, supports multiple files, parallel scans, sorted-file merge strategies, and schema import. The paths are resolved on the PostgreSQL server, not the client.

### Core Workflow

Create a server and map a Parquet file to a foreign table:

```sql
CREATE EXTENSION parquet_fdw;
CREATE SERVER parquet_srv FOREIGN DATA WRAPPER parquet_fdw;

CREATE FOREIGN TABLE public.userdata (
  id integer,
  first_name text,
  last_name text
)
SERVER parquet_srv
OPTIONS (filename '/srv/parquet/userdata.parquet');

SELECT * FROM public.userdata WHERE id >= 100;
```

The PostgreSQL operating-system account must be able to read every selected path. Writes through the foreign table are not supported.

### Files, Types, and Import

`filename` accepts a space-separated file list. Alternatively, `files_func` names a function taking one JSONB argument and returning `text[]`; `files_func_arg` supplies that JSONB value. Other options include `sorted`, `files_in_order`, `use_mmap`, `use_threads`, and `max_open_files`. Global switches include `parquet_fdw.use_threads`, `parquet_fdw.enable_multifile`, and `parquet_fdw.enable_multifile_merge`.

Supported mappings include Arrow integers, floats, timestamp, date, string, binary, list, and map to PostgreSQL numeric types, timestamp, date, text, bytea, arrays, and JSONB. Structs and nested lists are not supported. `IMPORT FOREIGN SCHEMA` treats the remote schema name as a quoted local directory path and creates foreign tables from files in that directory.

### Security and Correctness

The reviewed validator checks path existence and options but does not enforce a `file_fdw`-style superuser restriction. A role able to create or alter these foreign tables can attempt to read files accessible to the PostgreSQL server account or invoke an approved `files_func`. Restrict foreign-server usage, schema creation, table ownership, and function ownership to trusted administrators; use absolute paths below a dedicated read-only directory and schema-qualify immutable file-list functions.

The `sorted` and `files_in_order` options are promises to the planner. Incorrect declarations can produce incorrectly ordered results, so verify file ranges before enabling them. Pin matching PostgreSQL, Arrow, and Parquet library versions and test type conversion, timestamp units, NULLs, schema drift, memory mapping, file replacement, open-file limits, parallel scans, and malformed input on representative data.
