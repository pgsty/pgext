## Usage

Sources:

- [Official upstream README](https://github.com/paulharter/letter/blob/0ca81e61bb444dc30b46e0ee6b3c415f6e696a39/README.md)
- [Official extension control file (letter.control)](https://github.com/paulharter/letter/blob/0ca81e61bb444dc30b46e0ee6b3c415f6e696a39/letter.control)
- [Official extension SQL (letter--0.1.sql)](https://github.com/paulharter/letter/blob/0ca81e61bb444dc30b46e0ee6b3c415f6e696a39/sql/letter--0.1.sql)

`letter` — letter: role-based access control for PostgreSQL. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION letter;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `letter.assign(source_table text, user_column text, scope_table text DEFAULT NULL, role_name text DEFAULT NULL, role_column text DEFAULT NULL, if_fn text DEFAULT NULL)` is an extension function and returns `boolean`.
- `letter.cache_inval()` is an extension function and returns `trigger`.
- `letter.enforce_delete()` is an extension function and returns `trigger`.
- `letter.enforce_insert()` is an extension function and returns `trigger`.
- `letter.enforce_update()` is an extension function and returns `trigger`.
- `letter.grant(privilege text, on_table text, role text, columns text[], scope text, using_path text[] DEFAULT NULL, check_fn text DEFAULT NULL)` is an extension function and returns `boolean`.
- `letter.list_grants(filter_role text DEFAULT NULL)` is an extension function and returns `TABLE`.
- `letter.read(table_name text, condition text DEFAULT NULL)` is an extension function and returns `SETOF`.
- `letter.revoke(privilege text, on_table text, role text, columns text[], scope text)` is an extension function and returns `boolean`.
- `letter.role_cleanup()` is an extension function and returns `trigger`.
- `letter.unassign(source_table text, user_column text, scope_table text DEFAULT NULL, role_name text DEFAULT NULL, role_column text DEFAULT NULL)` is an extension function and returns `boolean`.
- `letter.user_permissions(p_user_id text)` is an extension function and returns `TABLE`.
- `letter.assignments` is a table installed or managed by the extension.
- `letter.grants` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
