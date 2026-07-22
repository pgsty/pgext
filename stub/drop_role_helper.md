## Usage

Sources:

- [Official README](https://github.com/cybertec-postgresql/drop_role_helper/blob/master/README.md)
- [Extension control file](https://github.com/cybertec-postgresql/drop_role_helper/blob/master/drop_role_helper.control)
- [Version 1.0 extension SQL](https://github.com/cybertec-postgresql/drop_role_helper/blob/master/drop_role_helper--1.0.sql)
- [PostgreSQL role-removal documentation](https://www.postgresql.org/docs/current/role-removal.html)

drop_role_helper provides a set-returning function that generates the revocation statements needed before removing a PostgreSQL role. It does not execute those statements, transfer ownership, remove owned objects, visit other databases, or drop the role itself.

### Core Workflow

Run the ownership steps and generated revocations in every database where the target role has dependencies. In psql, the meta-command shown below executes each returned row as SQL.

```sql
CREATE EXTENSION drop_role_helper SCHEMA public;

REASSIGN OWNED BY dummy TO otheruser;
DROP OWNED BY dummy;

SELECT * FROM public.drop_role_helper('dummy') \gexec
```

After repeating that sequence in every database in the cluster, remove the cluster-wide role:

```sql
DROP ROLE dummy;
```

Inspect the generated text before execution when using a client without psql's execution facility:

```sql
SELECT * FROM public.drop_role_helper('dummy');
```

### Generated Statements

The function accepts a role name through PostgreSQL's role OID type and returns one statement per affected object group. Version 1.0 covers explicit privileges on relations and sequences, columns, functions and procedures, schemas, large objects, types, languages, databases, tablespaces, foreign servers, and foreign-data wrappers. It also emits statements for matching default privileges.

### Privileges and Caveats

Calling the generator itself requires no special privilege, but executing its output requires enough authority to revoke every affected grant—normally superuser rights or ownership of the objects. For default privileges, the executor must also be a member of the role whose defaults are being changed.

Ownership handling remains separate. Reassignment does not cover every object class, so the documented sequence still includes the owned-object cleanup step. Role dependencies are database-local while the role is cluster-wide; a successful pass in one database does not prove the role is removable from the cluster.
