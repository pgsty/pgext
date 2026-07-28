## Usage

Sources:

- [Official upstream README](https://github.com/veer66/chamkho-pg/blob/3f946dc4280dc6dc80f189d3f6f7a35f60ebc9bf/README.org)
- [Official extension control file (chamkho_parser.control)](https://github.com/veer66/chamkho-pg/blob/3f946dc4280dc6dc80f189d3f6f7a35f60ebc9bf/chamkho_parser.control)
- [Official implementation source](https://github.com/veer66/chamkho-pg/blob/3f946dc4280dc6dc80f189d3f6f7a35f60ebc9bf/src/lib.rs)

`chamkho_parser` — ~cargo pgrx run~ — builds and starts a Postgres instance with the extension installed ~cargo pgrx test~ — runs the test suite. Use it for the corresponding text-search, parsing, or linguistic workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION chamkho_parser;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `chamkho_parser_end` is an extension function.
- `chamkho_parser_get_token` is an extension function.
- `chamkho_parser_start` is an extension function.

### Requirements and Caveats

- The catalog records version `0.6.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
