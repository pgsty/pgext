## Usage

Sources:

- [Official upstream README](https://github.com/dpxcc/pg_mooncake_rust/blob/d8e41dac48b45b9c0eaf06a7cd117fff6a65f532/README.md)
- [Official extension control file (pg_mooncake_rust.control)](https://github.com/dpxcc/pg_mooncake_rust/blob/d8e41dac48b45b9c0eaf06a7cd117fff6a65f532/pg_mooncake_rust.control)
- [Official implementation source](https://github.com/dpxcc/pg_mooncake_rust/blob/d8e41dac48b45b9c0eaf06a7cd117fff6a65f532/src/lib.rs)

`pg_mooncake_rust` — pg mooncake rust: Columnstore table in Postgres. Use it for the corresponding analytical or storage workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_mooncake_rust;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.0.1`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
