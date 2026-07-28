## Usage

Sources:

- [Official extension control file (rsam.control)](https://github.com/nekit2-002/rsam/blob/5052ad151b6442f5b066c6b47f9e0448fe9506e0/postgres-rust-table-am/rsam.control)
- [Official implementation source](https://github.com/nekit2-002/rsam/blob/5052ad151b6442f5b066c6b47f9e0448fe9506e0/postgres-rust-table-am/src/lib.rs)
- [Official Rust package manifest](https://github.com/nekit2-002/rsam/blob/5052ad151b6442f5b066c6b47f9e0448fe9506e0/postgres-rust-table-am/Cargo.toml)

`rsam` — A repository for new access method for Postgresql, written in Rust. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION rsam;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
