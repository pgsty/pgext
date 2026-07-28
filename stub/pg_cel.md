## Usage

Sources:

- [Official upstream README](https://github.com/thomasdarimont/pgauthz/blob/9d4351743ecce44a1a76cafe796d7c85fa03cf31/extensions/pg-cel/README.md)
- [Official extension control file (pg_cel.control)](https://github.com/thomasdarimont/pgauthz/blob/9d4351743ecce44a1a76cafe796d7c85fa03cf31/extensions/pg-cel/pg_cel.control)
- [Official implementation source](https://github.com/thomasdarimont/pgauthz/blob/9d4351743ecce44a1a76cafe796d7c85fa03cf31/extensions/pg-cel/src/lib.rs)

`pg_cel` — A small pgrx PostgreSQL extension that evaluates CEL (Common Expression Language) expressions, so pgauthz conditions can be written in CEL instead of raw SQL. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_cel;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `cel_compile_check` is an extension function.
- `cel_eval_bool` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
