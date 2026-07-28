## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_handlebars/pg_handlebars-1.0.7/README.md)
- [Official extension control file (pg_handlebars.control)](https://api.pgxn.org/src/pg_handlebars/pg_handlebars-1.0.7/pg_handlebars.control)
- [Official extension SQL (pg_handlebars--1.0.sql)](https://api.pgxn.org/src/pg_handlebars/pg_handlebars-1.0.7/pg_handlebars--1.0.sql)

`pg_handlebars` — PostgreSQL implementation of handlebars templating. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_handlebars;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `handlebars(json JSON, template TEXT)` is an extension function and returns `TEXT`.
- `handlebars(json JSON, template TEXT, file TEXT)` is an extension function and returns `BOOL`.
- `handlebars_compiler_flag_all()` is an extension function and returns `void`.
- `handlebars_compiler_flag_alternate_decorators()` is an extension function and returns `void`.
- `handlebars_compiler_flag_assume_objects()` is an extension function and returns `void`.
- `handlebars_compiler_flag_compat()` is an extension function and returns `void`.
- `handlebars_compiler_flag_explicit_partial_context()` is an extension function and returns `void`.
- `handlebars_compiler_flag_ignore_standalone()` is an extension function and returns `void`.
- `handlebars_compiler_flag_known_helpers_only()` is an extension function and returns `void`.
- `handlebars_compiler_flag_mustache_style_lambdas()` is an extension function and returns `void`.
- `handlebars_compiler_flag_no_escape()` is an extension function and returns `void`.
- `handlebars_compiler_flag_none()` is an extension function and returns `void`.
- `handlebars_compiler_flag_prevent_indent()` is an extension function and returns `void`.
- `handlebars_compiler_flag_strict()` is an extension function and returns `void`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
