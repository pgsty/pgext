## Usage

Sources:

- [Official database.dev package page](https://database.dev/pmnzt/custom_roles_patch)

`pmnzt@custom_roles_patch` — custom_roles. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "pmnzt@custom_roles_patch";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `custom_roles_update_to_app_metadata()` is an extension function and returns `trigger`.
- `get_user_roles()` is an extension function and returns `text[]`.
- `user_has_role(_role text)` is an extension function and returns `boolean`.
- `user_role_in(_roles text[])` is an extension function and returns `boolean`.
- `user_roles_match(_roles text[])` is an extension function and returns `boolean`.
- `custom_role_names` is a table installed or managed by the extension.
- `custom_user_roles` is a table installed or managed by the extension.
- `on_custom_role_change` is an extension-defined trigger.

### Requirements and Caveats

- The catalog records version `0.0.1`.
- This is a database.dev/pg_tle package; register or generate its package migration before issuing the quoted `CREATE EXTENSION` identity.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
