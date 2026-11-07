## Usage

Sources:

- [1.0.1 upgrade](https://github.com/crudylabs/pgmorbac/blob/v1.0.2/src/upgrades/pgmorbac--1.0.0--1.0.1.sql)
- [1.0.1 changes](https://github.com/crudylabs/pgmorbac/blob/v1.0.2/CHANGELOG.md)
- [pgmorbac.control](https://github.com/crudylabs/pgmorbac/blob/v1.0.2/pgmorbac.control)
- [README.md](https://github.com/crudylabs/pgmorbac/blob/v1.0.2/README.md)
- [SECURITY.md](https://github.com/crudylabs/pgmorbac/blob/v1.0.2/SECURITY.md)
- [Makefile](https://github.com/crudylabs/pgmorbac/blob/v1.0.2/Makefile)
- [src/authorization.sql](https://github.com/crudylabs/pgmorbac/blob/v1.0.2/src/authorization.sql)
- [src/rls.sql](https://github.com/crudylabs/pgmorbac/blob/v1.0.2/src/rls.sql)
- [src/system_rls.sql](https://github.com/crudylabs/pgmorbac/blob/v1.0.2/src/system_rls.sql)
- [1.0.2 upgrade](https://github.com/crudylabs/pgmorbac/blob/v1.0.2/src/upgrades/pgmorbac--1.0.1--1.0.2.sql)

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

Normally, rules with higher priority win and a prohibition wins an equal-priority tie. Users listed in `morbac.system_principals` bypass prohibition checks, but still need an applicable permission. Protect this table along with policy tables, context functions and delegation management. The application must supply authenticated `morbac.user_id` and organization context; a user-controlled session setting is not proof of identity. Apply explicit grants and verify access as the actual application role. This document describes source version 1.0.2 at the cited revision.

### 1.0.1 Authorization Cost

Uncached authorization resolves relevant organization-role pairs once instead of repeating that work for unrelated tenants' rules. Permission decisions retain the 1.0.0 semantics. After installing matching SQL files, run `ALTER EXTENSION pgmorbac UPDATE TO '1.0.1'` and verify the application's RLS behavior with its actual roles.

### 1.0.2 NULL Identity Handling

`morbac.is_allowed` and `morbac.is_allowed_nocache` now return false for a NULL user ID, matching `morbac.rls_check`, instead of failing on the authorization cache's non-null constraint. Decisions for actual users are unchanged. After installing the release SQL files, advance an existing 1.0.1 installation with:

```sql
ALTER EXTENSION pgmorbac UPDATE TO '1.0.2';
```
