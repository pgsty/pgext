## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_role_fkey_trigger_functions/pg_role_fkey_trigger_functions-1.0.4/README.md)
- [Official extension control file (pg_role_fkey_trigger_functions.control)](https://api.pgxn.org/src/pg_role_fkey_trigger_functions/pg_role_fkey_trigger_functions-1.0.4/pg_role_fkey_trigger_functions.control)
- [Official extension SQL (pg_role_fkey_trigger_functions--0.11.7.sql)](https://api.pgxn.org/src/pg_role_fkey_trigger_functions/pg_role_fkey_trigger_functions-1.0.4/sql/pg_role_fkey_trigger_functions--0.11.7.sql)

`pg_role_fkey_trigger_functions` — The pg_role_fkey_trigger_functions PostgreSQL extension offers a bunch of trigger functions to help establish and/or maintain referential integrity for columns that reference PostgreSQL ROLE NAMEs. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_role_fkey_trigger_functions;

-- Using the `SET` command:
set pg_role_fkey_trigger_functions.trusted_tables TO '{pg_temp.evil_temp_tbl}';

-- Using the `set_config()` function:
select set_config(
    'pg_role_fkey_trigger_functions.trusted_tables',
    '{pg_temp.evil_temp_tbl}',
    false
);

-- Or, appending to instead of replacing the list of trusted tables:
select set_config(
    'pg_role_fkey_trigger_functions.trusted_tables',
    coalesce(
        current_setting('pg_role_fkey_trigger_functions.trusted_tables', true),
        '{}'
    )::text[] || 'pg_temp.evil_temp_tbl',
    false
);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `enforce_fkey_to_db_role()` is an extension function and returns `trigger`.
- `grant_role_in_column1_to_role_in_column2()` is an extension function and returns `trigger`.
- `maintain_referenced_role()` is an extension function and returns `trigger`.
- `pg_role_fkey_trigger_functions_meta_pgxn()` is an extension function and returns `jsonb`.
- `pg_role_fkey_trigger_functions_readme()` is an extension function and returns `text`.
- `revoke_role_in_column1_from_role_in_column2()` is an extension function and returns `trigger`.
- `test__pg_role_fkey_trigger_functions` is an extension procedure.
- `test_dump_restore__maintain_referenced_role` is an extension procedure.
- `test__customer` is a table installed or managed by the extension.
- `test__tbl` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.4`.
- The control file marks the extension as relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
