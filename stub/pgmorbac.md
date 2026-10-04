## Usage

Sources:

- [pgmorbac.control](https://github.com/crudylabs/pgmorbac/blob/3e8380d4f7707d391b4a6d59221d804832df49f3/pgmorbac.control)
- [README.md](https://github.com/crudylabs/pgmorbac/blob/3e8380d4f7707d391b4a6d59221d804832df49f3/README.md)
- [SECURITY.md](https://github.com/crudylabs/pgmorbac/blob/3e8380d4f7707d391b4a6d59221d804832df49f3/SECURITY.md)
- [Makefile](https://github.com/crudylabs/pgmorbac/blob/3e8380d4f7707d391b4a6d59221d804832df49f3/Makefile)
- [src/authorization.sql](https://github.com/crudylabs/pgmorbac/blob/3e8380d4f7707d391b4a6d59221d804832df49f3/src/authorization.sql)
- [src/rls.sql](https://github.com/crudylabs/pgmorbac/blob/3e8380d4f7707d391b4a6d59221d804832df49f3/src/rls.sql)
- [src/system_rls.sql](https://github.com/crudylabs/pgmorbac/blob/3e8380d4f7707d391b4a6d59221d804832df49f3/src/system_rls.sql)

`pgmorbac` implements organization-based authorization in the fixed `morbac` schema. Roles, activities, views, permissions, prohibitions and delegation form a policy model that application RLS policies can consult.

### Core Workflow

```sql
CREATE EXTENSION pgmorbac;
SELECT morbac.refresh_hierarchy_cache();
```

After policies and trusted session identity have been configured:

```sql
SELECT morbac.is_allowed(
  '00000000-0000-0000-0000-000000000001'::uuid,
  NULL::uuid, 'read', 'documents');
```

### Operational Boundaries

Install as an administrator on PostgreSQL 13 or later; this SQL extension needs no preload or restart. Create organizations and roles, assign application user UUIDs, compile policy, then attach `morbac.rls_check` to the intended application tables. Installing the extension alone does not enable RLS on application tables.

Use `morbac.is_allowed` for an object's authorization decision, `morbac.is_allowed_nocache` to bypass its result cache, and `morbac.refresh_hierarchy_cache` to refresh hierarchy data. `morbac.has_permission` is a capability probe for UI gating, not an object-level authorization check. A NULL row organization denotes an unattributed object, not every organization.

Normally, rules with higher priority win and a prohibition wins an equal-priority tie. Users listed in `morbac.system_principals` bypass prohibition checks, but still need an applicable permission. Protect this table along with policy tables, context functions and delegation management. The application must supply authenticated `morbac.user_id` and organization context; a user-controlled session setting is not proof of identity. Apply explicit grants and verify access as the actual application role. This document describes source version 1.0.0 at the cited revision.
