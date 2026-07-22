## Usage

Sources:

- [Official README](https://github.com/CrunchyData/crunchy_check_access/blob/c8f01b40d7507715c4f0fda18a651916b010653a/README.md)
- [Extension SQL: effective access](https://github.com/CrunchyData/crunchy_check_access/blob/c8f01b40d7507715c4f0fda18a651916b010653a/sql/check_access.sql)
- [Extension SQL: raw grants](https://github.com/CrunchyData/crunchy_check_access/blob/c8f01b40d7507715c4f0fda18a651916b010653a/sql/check_grants.sql)

`check_access` 0.1 is a database-local auditing extension from Crunchy Data. It expands PostgreSQL role membership and reports either privileges a role can actually use or the underlying grants. Install it as a database superuser; it inspects the current database only and does not change privileges.

### Core Workflow

Create the extension, then start with the caller-safe views:

```sql
CREATE EXTENSION check_access;

SELECT *
FROM my_privs
WHERE object_schema = 'app';

SELECT *
FROM my_privs_sys
WHERE object_schema = 'pg_catalog';
```

`my_privs` omits system objects, while `my_privs_sys` includes them. Both are backed by security-definer functions and are the only reporting interfaces left executable by `PUBLIC`. Access to the broader functions is revoked by the installation script, so an administrator must grant it deliberately if non-superusers should audit other roles.

For a privileged audit, compare usable access with entitlement data:

```sql
SELECT * FROM all_access() WHERE base_role = 'app_user';
SELECT * FROM all_access(true) WHERE base_role = 'app_user';

SELECT * FROM check_access('app_user', false);
SELECT * FROM check_grants('app_user', false);
```

`check_access` and `all_access` require the role to have `USAGE` on the containing schema, so they describe effective access. `check_grants` and `all_grants` retain raw object grants even when schema access makes them unusable. Their Boolean argument controls inclusion of system objects.

### Report Columns and Role Expansion

Important columns include `base_role` (the role being audited), `as_role` (the role that supplies the privilege), `role_path` (membership ancestry), object identity, privilege name, and whether the privilege has grant option. Membership is followed even across `NOINHERIT` links; `role_path` marks each ancestor with `(true)` or `(false)` to show inheritability.

The SQL covers database, tablespace, foreign-data wrapper, foreign server, language, schema, function, table, view, sequence, and column privileges. This is an old catalog-level implementation: it predates several newer PostgreSQL object kinds and does not claim comprehensive coverage of procedures, materialized views, partitioned relations, or future catalog changes. Validate its output against the PostgreSQL version in use before treating it as a complete security audit.
