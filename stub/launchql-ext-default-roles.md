## Usage

Sources:

- [packages/default-roles/launchql-ext-default-roles.control](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/default-roles/launchql-ext-default-roles.control)
- [packages/default-roles/sql/launchql-ext-default-roles--0.4.5.sql](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/default-roles/sql/launchql-ext-default-roles--0.4.5.sql)

`launchql-ext-default-roles` 0.4.5 creates the shared roles expected by LaunchQL extensions.

### Core Workflow

```sql
CREATE EXTENSION "launchql-ext-default-roles" CASCADE;
```

### Roles and Privileges

The SQL creates `anonymous`, `authenticated` and `administrator` only when they do not already exist. A newly created administrator inherits the other two groups. Review existing role definitions before installation; existing roles are reused rather than normalized.

The control requires `plpgsql` and permits a non-superuser installer, but role creation still requires the relevant cluster privileges. There is no shared library or preload. Roles are cluster-wide, and `DROP EXTENSION` does not remove them. PostgreSQL-major compatibility is not declared in the reviewed component.

Newly creating `administrator` also grants it `BYPASSRLS`, which requires superuser privileges. Review membership in this role carefully.
