## Usage

Sources:

- [Official database.dev package page](https://database.dev/mansueli/function_vc)

`mansueli-function_vc` — In database version control for postgres functions. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "mansueli-function_vc";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `archive.save_function_history(function_name text, args text, return_type text, source_code text, schema_name text default 'public', lang_settings text default 'plpgsql')` is an extension function and returns `void`.
- `archive.setup_function_history(schema_name text default 'public')` is an extension function and returns `VOID`.
- `calculate_version()` is an extension function and returns `TRIGGER`.
- `public.create_function_from_source(function_text text, schema_name text default 'public')` is an extension function and returns `text`.
- `rollback_function(func_name text, version_no integer default 0, schema_n text default 'public')` is an extension function and returns `text`.
- `archive.function_history` is a table installed or managed by the extension.
- `archive` is a schema created by the extension.
- `before_insert_function_history` is an extension-defined trigger.

### Requirements and Caveats

- The catalog records version `1.0.1`.
- This is a database.dev/pg_tle package; register or generate its package migration before issuing the quoted `CREATE EXTENSION` identity.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
