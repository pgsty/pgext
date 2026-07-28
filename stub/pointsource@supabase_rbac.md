## Usage

Sources:

- [Official database.dev package page](https://database.dev/pointsource/supabase_rbac)

`pointsource@supabase_rbac` — Role-Based Access Control for your Supabase project. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "pointsource@supabase_rbac";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `_build_user_claims` is an extension function.
- `_check_permission_escalation` is an extension function.
- `_check_role_escalation` is an extension function.
- `_get_user_groups` is an extension function.
- `_jwt_is_expired` is an extension function.
- `_on_group_created` is an extension function.
- `_on_role_definition_change` is an extension function.
- `_set_updated_at` is an extension function.
- `_sync_member_metadata` is an extension function.
- `_sync_member_permission` is an extension function.
- `_validate_grantable_roles` is an extension function.
- `_validate_permissions` is an extension function.
- `_validate_roles` is an extension function.
- `accept_invite` is an extension function.

### Requirements and Caveats

- The catalog records version `5.2.1`.
- This is a database.dev/pg_tle package; register or generate its package migration before issuing the quoted `CREATE EXTENSION` identity.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
