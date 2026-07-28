## Usage

Sources:

- [Official upstream README](https://github.com/allenvox/pg_seal/blob/3658d3608ba4ea3867cc0758149fed75730b88e3/README.md)
- [Official extension control file (pg_seal.control)](https://github.com/allenvox/pg_seal/blob/3658d3608ba4ea3867cc0758149fed75730b88e3/pg_seal.control)
- [Official implementation source](https://github.com/allenvox/pg_seal/blob/3658d3608ba4ea3867cc0758149fed75730b88e3/src/lib.rs)

`pg_seal` — A PostgreSQL extension written in Rust (pgrx) for a cryptographic audit log of data changes. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_seal;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_seal_append` is an extension function.
- `pg_seal_hash` is an extension function.
- `pg_seal_verify()` is an extension function.
- `pg_seal_verify_detail()` is an extension function.

### Requirements and Caveats

- The catalog records version `1.0.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
