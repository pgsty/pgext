## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_utility_trigger_functions/pg_utility_trigger_functions-1.9.3/README.md)
- [Official extension control file (pg_utility_trigger_functions.control)](https://api.pgxn.org/src/pg_utility_trigger_functions/pg_utility_trigger_functions-1.9.3/pg_utility_trigger_functions.control)
- [Official extension SQL (pg_utility_trigger_functions--1.0.0.sql)](https://api.pgxn.org/src/pg_utility_trigger_functions/pg_utility_trigger_functions-1.9.3/sql/pg_utility_trigger_functions--1.0.0.sql)

`pg_utility_trigger_functions` — The pg_utility_trigger_functions PostgreSQL extensions bundles together some pet trigger functions that the extension author—BigSmoke—likes to walk through various PostgreSQL projects. Use it for the corresponding SQL or database utility workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pg_utility_trigger_functions;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `coalesce_sibling_fields()` is an extension function and returns `trigger`.
- `copy_fields_from_foreign_table()` is an extension function and returns `trigger`.
- `fallback_to_fields_from_foreign_table()` is an extension function and returns `trigger`.
- `no_delete()` is an extension function and returns `trigger`.
- `nullify_columns()` is an extension function and returns `trigger`.
- `overwrite_composite_field_in_referencing_table()` is an extension function and returns `trigger`.
- `overwrite_fields_in_referencing_table()` is an extension function and returns `trigger`.
- `pg_utility_trigger_functions_meta_pgxn()` is an extension function and returns `jsonb`.
- `pg_utility_trigger_functions_readme()` is an extension function and returns `text`.
- `set_installed_extension_version_from_name()` is an extension function and returns `trigger`.
- `test__mock.now()` is an extension function and returns `timestamptz`.
- `update_updated_at()` is an extension function and returns `trigger`.
- `test__coalesce_sibling_fields` is an extension procedure.
- `test__copy_fields_from_foreign_table` is an extension procedure.

### Requirements and Caveats

- The reviewed control file declares default version `1.9.3`.
- Install the confirmed extension dependencies first: `hstore`.
- The control file marks the extension as relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
