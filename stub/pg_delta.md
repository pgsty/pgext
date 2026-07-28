## Usage

Sources:

- [Official upstream README](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_delta/README.md)
- [Official extension control file (pg_delta.control)](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_delta/pg_delta.control)
- [Official implementation source](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_delta/src/lib.rs)

`pg_delta` — Delta Lake streaming integration for PostgreSQL. Stream data bidirectionally between Postgres and Delta Lake tables. Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_delta;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `create_table` is an extension function.
- `drop_export` is an extension function.
- `drop_table` is an extension function.
- `export` is an extension function.
- `export_table` is an extension function.
- `extension_docs()` is an extension function.
- `history` is an extension function.
- `info` is an extension function.
- `list_exports()` is an extension function.
- `list_tables()` is an extension function.
- `read` is an extension function.
- `refresh` is an extension function.
- `schema` is an extension function.
- `status()` is an extension function.

### Requirements and Caveats

- The catalog records version `0.2.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
