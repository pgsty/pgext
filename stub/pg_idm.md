## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/pasha132/pg_idm/blob/9968b613451322e0c4cc9cb2ecccf0693883d7f3/README.md)
- [Version 1.0 SQL objects](https://github.com/pasha132/pg_idm/blob/9968b613451322e0c4cc9cb2ecccf0693883d7f3/pg_idm--1.0.sql)
- [Extension control file](https://github.com/pasha132/pg_idm/blob/9968b613451322e0c4cc9cb2ecccf0693883d7f3/pg_idm.control)
- [GPL-3.0 license](https://github.com/pasha132/pg_idm/blob/9968b613451322e0c4cc9cb2ecccf0693883d7f3/LICENSE)

`pg_idm` presents PostgreSQL roles and memberships as writable relational views. Install it in a dedicated caller-chosen schema to avoid object-name collisions.

```sql
CREATE SCHEMA idm;
CREATE EXTENSION pg_idm WITH SCHEMA idm;
SELECT rolname, rolcanlogin, rolconnlimit
FROM idm.roles
ORDER BY rolname;
```

DML on `idm.roles` maps inserts, updates, and deletes to role creation, alteration, and removal. DML on `idm.auth_members` maps inserts and deletes to membership grants and revocations. These operations change cluster-wide security principals, not merely rows in the current database.

The trigger functions execute with the caller's privileges, so callers still need the corresponding PostgreSQL role-management authority. Grant write access to these views only to a tightly controlled administrative role, and audit changes just as you would direct `CREATE ROLE`, `ALTER ROLE`, `GRANT`, and `DROP ROLE` commands.
