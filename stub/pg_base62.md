## Usage

Sources:

- [Official upstream README](https://github.com/bkircher/pg_base62/blob/19f8807844d59f44a776f43752b1c23635ccc5e2/README.md)
- [Official extension control file (pg_base62.control)](https://github.com/bkircher/pg_base62/blob/19f8807844d59f44a776f43752b1c23635ccc5e2/pg_base62.control)
- [Official implementation source](https://github.com/bkircher/pg_base62/blob/19f8807844d59f44a776f43752b1c23635ccc5e2/src/lib.rs)

`pg_base62` — A PostgreSQL extension for encoding UUIDs to Base62 alphabet and decoding Base62 back to UUID. It is built in Rust with pgrx. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_base62;

SELECT base62_encode('f81d4fae-7dec-11d0-a765-00a0c91e6bf6'::uuid);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `base62_decode` is an extension function.
- `base62_encode` is an extension function.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
