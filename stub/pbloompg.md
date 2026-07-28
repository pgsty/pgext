## Usage

Sources:

- [Official upstream README](https://github.com/godwhoa/pbloom/blob/fc189b436337ae0812800c06903162d0b6ee1fe9/README.md)
- [Official extension control file (pbloompg.control)](https://github.com/godwhoa/pbloom/blob/fc189b436337ae0812800c06903162d0b6ee1fe9/pg/pbloompg.control)
- [Official implementation source](https://github.com/godwhoa/pbloom/blob/fc189b436337ae0812800c06903162d0b6ee1fe9/pg/src/lib.rs)

`pbloompg` — portable bloom filter, create, serialize, query across Go, Rust, Postgres. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pbloompg;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pbloom_add` is an extension function.
- `pbloom_contains` is an extension function.
- `pbloom_create` is an extension function.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
