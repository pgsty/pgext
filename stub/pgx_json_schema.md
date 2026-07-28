## Usage

Sources:

- [Official upstream README](https://github.com/jefbarn/pgx_json_schema/blob/00a573b165139eae9ff78149b22a3b4dcc2a9c69/README.md)
- [Official extension control file (pgx_json_schema.control)](https://github.com/jefbarn/pgx_json_schema/blob/00a573b165139eae9ff78149b22a3b4dcc2a9c69/pgx_json_schema.control)
- [Official implementation source](https://github.com/jefbarn/pgx_json_schema/blob/00a573b165139eae9ff78149b22a3b4dcc2a9c69/src/lib.rs)

`pgx_json_schema` — A JSON Schema validator for Postgres implemented in Rust. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgx_json_schema;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
