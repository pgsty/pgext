## Usage

Sources:

- [pgduck_rs GOAL at the reviewed commit](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/GOAL.md)
- [pgduck_rs.control at the reviewed commit](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/pgduck_rs.control)
- [Cargo package metadata](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/Cargo.toml)
- [read_parquet implementation](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/src/read_parquet.rs)
- [Extension module exports](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/src/lib.rs)

`pgduck_rs` is an exploratory pgrx extension that embeds DuckDB in a PostgreSQL backend. At the reviewed commit, its only data-reading operation is `read_parquet(path text)`, which returns each Parquet row as `SETOF jsonb`; `hello_pgduck_rs()` is only a demonstration function.

### Core Workflow

Install the extension as a superuser, then pass it a file path visible on the database server:

```sql
CREATE EXTENSION pgduck_rs;

SELECT row
FROM read_parquet('/srv/postgres/import/events.parquet') AS t(row);
```

The path is resolved by the server process, not by the SQL client. The PostgreSQL operating-system account must be able to read the file.

### Object Index

- `read_parquet(path text) RETURNS SETOF jsonb`: opens the local Parquet file through an in-memory DuckDB connection and emits one JSON object per row.
- `hello_pgduck_rs() RETURNS text`: returns a fixed greeting and is not part of the Parquet workflow.

### Operational Boundaries

- The upstream GOAL explicitly describes “learn, don't ship” prototype work. Repository availability does not make it production-ready.
- Each `read_parquet` call creates a new in-memory DuckDB connection and first collects the complete result into a Rust vector. Large results therefore consume backend memory before any row is returned.
- The reviewed implementation uses DuckDB's local `read_parquet(?)` call. It does not load `httpfs`, create secrets, or provide a remote-object-store workflow. Treat its path argument as server-file access.
- DuckDB executes inside the PostgreSQL backend. Parser, file-decoder, allocation, or native-code failures therefore share the backend process boundary.
- The control file requires superuser installation and marks the extension non-relocatable. Cargo and the control file report version `0.0.0`; the declared pgrx build features for PostgreSQL 13 through 18 are not a compatibility or support guarantee.
