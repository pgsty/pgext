## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/pasha132/pg_idm/blob/9968b613451322e0c4cc9cb2ecccf0693883d7f3/README.md)
- [Version 1.0 SQL objects](https://github.com/pasha132/pg_idm/blob/9968b613451322e0c4cc9cb2ecccf0693883d7f3/pg_idm--1.0.sql)
- [Extension control file](https://github.com/pasha132/pg_idm/blob/9968b613451322e0c4cc9cb2ecccf0693883d7f3/pg_idm.control)

`pg_idm` exposes PostgreSQL roles and role memberships as writable relational views. `INSERT`, `UPDATE`, and `DELETE` operations are translated by `INSTEAD OF` triggers into role DDL, which can simplify tightly controlled identity-management integrations. The rows represent cluster-wide security principals, not data local to one database.

Install it in a dedicated caller-chosen schema because the extension is non-relocatable after creation and uses generic object names.

```sql
CREATE SCHEMA idm;
CREATE EXTENSION pg_idm WITH SCHEMA idm;

SELECT rolname, rolcanlogin, rolconnlimit
FROM idm.roles
ORDER BY rolname;
```

### Manage Roles

```sql
INSERT INTO idm.roles (rolname, rolcanlogin, rolconnlimit)
VALUES ('app_reader', true, 10)
RETURNING rolname, rolcanlogin, rolconnlimit;

UPDATE idm.roles
SET rolcanlogin = false
WHERE rolname = 'app_reader';

DELETE FROM idm.roles
WHERE rolname = 'app_reader';
```

`idm.roles` projects the attributes in `pg_roles`, including `rolsuper`, `rolinherit`, `rolcreaterole`, `rolcreatedb`, `rolcanlogin`, `rolreplication`, `rolconnlimit`, `rolpassword`, `rolvaliduntil`, and `rolbypassrls`. Omitted insert attributes use the extension's role defaults; updates generate only changed role options.

### Manage Memberships

```sql
INSERT INTO idm.auth_members (roleid, member, admin_option)
VALUES ('reporting'::regrole, 'app_reader'::regrole, false);

DELETE FROM idm.auth_members
WHERE roleid = 'reporting'::regrole
  AND member = 'app_reader'::regrole;
```

`idm.auth_members` mirrors `pg_auth_members`; inserts issue `GRANT` and deletes issue `REVOKE`. The view does not provide an update workflow for memberships.

### Security and Operational Notes

- Trigger functions are explicitly `SECURITY INVOKER`. A caller needs both DML privileges on the views and the PostgreSQL authority required for the generated role command; the extension does not elevate the caller.
- Role changes affect every database in the cluster. Restrict write grants to a dedicated administrative role, audit them like direct role DDL, and test ownership and dependency failures before automated deletion.
- Password values supplied through DML can appear in client history, SQL logs, audit records, or monitoring. Prefer approved secret-handling and parameterized interfaces; do not use the view as a password-distribution channel without reviewing those paths.
- The control file says installation itself does not require a superuser, but actual role attributes such as `SUPERUSER`, `REPLICATION`, and `BYPASSRLS` remain subject to core PostgreSQL privilege rules.
- Upstream version `1.0` claims PostgreSQL 9.6 and later but publishes no maintained compatibility matrix. Verify generated DDL and catalog columns against the exact server major in use.
