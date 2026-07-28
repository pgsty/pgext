## Usage

Sources:

- [Official upstream README](https://github.com/kaznak/pgx_basex_codec/blob/8a272c5ac73ad3965e0a37f15e415cc56df33e89/README.md)
- [Official extension control file (pgx_basex_codec.control)](https://github.com/kaznak/pgx_basex_codec/blob/8a272c5ac73ad3965e0a37f15e415cc56df33e89/pgx_basex_codec.control)
- [Official implementation source](https://github.com/kaznak/pgx_basex_codec/blob/8a272c5ac73ad3965e0a37f15e415cc56df33e89/src/lib.rs)

`pgx_basex_codec` — A PostgreSQL extension for BaseX encoding and decoding using pgrx. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgx_basex_codec;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
