## Usage

Sources:

- [Upstream README](https://github.com/cybertec-postgresql/drop_role_helper/blob/b56500c0b4167ed8cb744c390c0b7cf24c0fd9d9/README.md)
- [Extension SQL](https://github.com/cybertec-postgresql/drop_role_helper/blob/b56500c0b4167ed8cb744c390c0b7cf24c0fd9d9/drop_role_helper--1.0.sql)
- [Extension control file](https://github.com/cybertec-postgresql/drop_role_helper/blob/b56500c0b4167ed8cb744c390c0b7cf24c0fd9d9/drop_role_helper.control)

`drop_role_helper` provides the table function `drop_role_helper(regrole)`, which generates SQL statements that revoke a role's privileges in the current database. In `psql`, feed the returned statements to `\gexec` after dealing with owned objects:

```sql
REASSIGN OWNED BY doomed_role TO replacement_role;
DROP OWNED BY doomed_role;

CREATE EXTENSION drop_role_helper SCHEMA public;
SELECT * FROM public.drop_role_helper('doomed_role') \gexec
```

Repeat that workflow in every database in the cluster, then run `DROP ROLE doomed_role` once no dependencies remain.

### Permissions and safety

Calling `drop_role_helper(regrole)` itself requires no special permission. Executing its generated SQL requires superuser rights or ownership of every affected object; default privileges can additionally require membership in their defining role. Review the generated rows before using `\gexec`, especially on shared or production clusters.
