## Usage

Sources:

- [Official extension control file (object_reference.control)](https://api.pgxn.org/src/object_reference/object_reference-0.1.0/object_reference.control)
- [Official extension SQL (object_reference.sql)](https://api.pgxn.org/src/object_reference/object_reference-0.1.0/sql/object_reference.sql)

`object_reference` — Provides immutable references to Postgres objects. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION object_reference;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `__object_reference.create_function(function_name text , args text , options text , body text , comment text , grants text DEFAULT NULL)` is an extension function and returns `void`.
- `__object_reference.exec(sql text)` is an extension function and returns `void`.
- `__object_reference.safe_dump(relation regclass , filter text DEFAULT '')` is an extension function and returns `void`.
- `snitch()` is an extension function and returns `event_trigger`.
- `_object_reference.object` is a table installed or managed by the extension.
- `_object_reference.object_group` is a table installed or managed by the extension.
- `_object_reference.object_group__object` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- Install the confirmed extension dependencies first: `cat_tools`, `count_nulls`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
