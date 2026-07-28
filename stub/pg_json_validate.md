## Usage

Sources:

- [Official extension control file (pg_json_validate.control)](https://github.com/jefbarn/pg_json_validate/blob/9804a7931c8f9b0a0dcb15c99d6a7d61488d13f5/pg_json_validate.control)
- [Official implementation source](https://github.com/jefbarn/pg_json_validate/blob/9804a7931c8f9b0a0dcb15c99d6a7d61488d13f5/src/lib.rs)
- [Official Rust package manifest](https://github.com/jefbarn/pg_json_validate/blob/9804a7931c8f9b0a0dcb15c99d6a7d61488d13f5/Cargo.toml)

`pg_json_validate` — JSON Schema validation functions for jsonb values. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_json_validate;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `json_schema_is_valid` is an extension function.
- `json_schema_validate` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
