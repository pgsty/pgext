## Usage

Sources:

- [Official README at the catalog revision](https://github.com/hollobon/postgresql_track_renames/blob/dcc35b6031c81f304e11f5c533a1a1a773e5bc2b/README.md)
- [Extension SQL at the catalog revision](https://github.com/hollobon/postgresql_track_renames/blob/dcc35b6031c81f304e11f5c533a1a1a773e5bc2b/track_renames--0.1.sql)
- [Event-trigger implementation at the catalog revision](https://github.com/hollobon/postgresql_track_renames/blob/dcc35b6031c81f304e11f5c533a1a1a773e5bc2b/track_renames.c)

`track_renames` 0.1 supplies an event-trigger function that calls an application callback after PostgreSQL rename commands. The implementation requires a five-text-argument callback; the four-argument callback shown in the old README does not match the pinned C source.

### Core Workflow

Create a callback visible by its unqualified name, set the callback name, and register an event trigger. Creating event triggers requires superuser privileges.

```sql
CREATE EXTENSION track_renames;

CREATE TABLE rename_log (
    object_type text,
    schema_name text,
    object_name text,
    subobject_name text,
    new_name text,
    changed_at timestamptz DEFAULT clock_timestamp()
);

CREATE FUNCTION log_rename(text, text, text, text, text)
RETURNS void
LANGUAGE sql
AS $$
  INSERT INTO rename_log
  VALUES ($1, $2, $3, $4, $5, clock_timestamp());
$$;

SET track_renames.function = 'log_rename';

CREATE EVENT TRIGGER track_object_renames
ON ddl_command_end
WHEN TAG IN ('ALTER TABLE', 'ALTER FUNCTION', 'ALTER TYPE', 'ALTER VIEW')
EXECUTE FUNCTION track_renames();
```

Persist the setting at database or role level only after confirming the callback exists in every affected session's search path.

### Callback Contract

- The intended callback returns void and must accept five text arguments: object type, schema name, object name, subobject name, and new name.
- Some arguments are NULL for rename forms whose parse tree does not expose a relation or subobject.
- The event trigger ignores intercepted commands whose parse tree is not a rename statement.

### Operational Caveats

- Callback lookup uses an unqualified function name and the session search path. Use a locked-down search path and prevent untrusted users from creating shadow functions in visible schemas.
- Callback errors run in the DDL transaction and can make the rename fail. Keep logging code short, deterministic, and locally transactional.
- The code targets old PostgreSQL parse-tree APIs and enumerations. Compile and regression-test against each target server major and every rename class you depend on.
