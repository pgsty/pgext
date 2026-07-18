## Usage

Sources:

- [pgduck_rs GOAL at the reviewed commit](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/GOAL.md)
- [pgduck_rs.control at the reviewed commit](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/pgduck_rs.control)
- [Cargo package metadata](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/Cargo.toml)
- [read_parquet implementation](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/src/read_parquet.rs)
- [Extension module exports](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/src/lib.rs)

`pgduck_rs` is an exploratory pgrx extension for calling embedded DuckDB from PostgreSQL. At the reviewed commit, its implemented data operation is `read_parquet(path text)`, which returns each Parquet row as `SETOF jsonb`; the module also exports the demonstration function `hello_pgduck_rs`.

### Read a Local Parquet File

```sql
CREATE EXTENSION pgduck_rs;
SELECT row
FROM read_parquet('/var/lib/postgresql/data/events.parquet') AS t(row);
```

The path is resolved on the database server and must be readable by the PostgreSQL operating-system account.

### Caveats

- The upstream GOAL explicitly describes “learn, don't ship” prototype work. Catalog lifecycle `active` means the repository is reachable; it does not override that upstream readiness statement.
- Every `read_parquet` call creates a fresh in-memory DuckDB connection and materializes the full result before returning rows. Large files can consume substantial backend memory.
- The reviewed code supports local files only and does not load `httpfs`, credentials, secrets, or remote-object-store support. Apply normal server-file access controls.
- Cargo and the control file both report version `0.0.0`. The pgrx feature set names PostgreSQL 13 through 18, but this is not a production compatibility guarantee.
