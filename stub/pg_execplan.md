## Usage

Sources:

- [Extension control file](https://github.com/danolivo/execplan/blob/8591179dc1c95f93b2a4987bef648240dcf9b899/pg_execplan.control)
- [Installation SQL](https://github.com/danolivo/execplan/blob/8591179dc1c95f93b2a4987bef648240dcf9b899/init.sql)
- [SQL file generation](https://github.com/danolivo/execplan/blob/8591179dc1c95f93b2a4987bef648240dcf9b899/Makefile)
- [Two-node integration test](https://github.com/danolivo/execplan/blob/8591179dc1c95f93b2a4987bef648240dcf9b899/tests/rpl.sh)
- [Project and PostgreSQL patch](https://github.com/danolivo/execplan/blob/8591179dc1c95f93b2a4987bef648240dcf9b899/README.md)

`pg_execplan` is a research extension for exchanging serialized query plans between compatible, patched PostgreSQL 12 development instances. The `pg_repeater` prototype uses its SQL interface on the receiving server. The upstream project has several revisions labeled `0.1` with different function names. The sources above identify the revision that exposes `pg_exec_plan`.

### Enable the receiving database

Use a superuser in a database running the matching patched server and library. The upstream integration test preloads `postgres_fdw` and `pg_execplan` on both nodes. The control file declares no SQL extension dependencies and fixes the extension in the schema chosen at installation.

```sql
CREATE EXTENSION pg_execplan;

SELECT p.oid::regprocedure
FROM pg_extension AS e
JOIN pg_depend AS d ON d.refclassid = 'pg_extension'::regclass
                   AND d.refobjid = e.oid AND d.deptype = 'e'
JOIN pg_proc AS p ON d.classid = 'pg_proc'::regclass AND p.oid = d.objid
WHERE e.extname = 'pg_execplan'
ORDER BY p.proname;
```

At this revision, `pg_store_query_plan` writes a planned statement to a server-side file, `pg_exec_stored_plan` reads a saved plan, and `pg_exec_plan` accepts query text and its serialized plan. Use trusted input and restrict access to these functions. Plans contain internal server structures. Use matching source revisions and catalog definitions on both nodes.
