## Usage

Sources:

- [Official upstream README](https://github.com/cardano-community/pg_blake2b/blob/c1144ef6ff938fff9aa49574aa345805b6ca656f/README.md)
- [Official extension control file (blake2b.control)](https://github.com/cardano-community/pg_blake2b/blob/c1144ef6ff938fff9aa49574aa345805b6ca656f/blake2b.control)
- [Official extension SQL (blake2b--1.0.sql)](https://github.com/cardano-community/pg_blake2b/blob/c1144ef6ff938fff9aa49574aa345805b6ca656f/blake2b--1.0.sql)

`blake2b` — PostgreSQL extension for fast secure hashing. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION blake2b;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `blake2b(data bytea, digest_size integer DEFAULT NULL, key bytea DEFAULT NULL)` is an extension function and returns `bytea`.
- `blake2b(data text, digest_size integer DEFAULT NULL, key bytea DEFAULT NULL)` is an extension function and returns `bytea`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
