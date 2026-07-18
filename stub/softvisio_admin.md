## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/softvisio/postgresql-softvisio-admin/blob/67e923d3761b0ede5adb16a3ec9cf47779d8099e/README.md)
- [Version 1.2.9 SQL objects](https://github.com/softvisio/postgresql-softvisio-admin/blob/67e923d3761b0ede5adb16a3ec9cf47779d8099e/softvisio_admin--1.2.9.sql)
- [Extension control file](https://github.com/softvisio/postgresql-softvisio-admin/blob/67e923d3761b0ede5adb16a3ec9cf47779d8099e/softvisio_admin.control)

`softvisio_admin` is a pure-SQL administration extension for creating an application database and role, reporting outdated extensions across a cluster, and updating installed extensions across databases.

```sql
CREATE EXTENSION softvisio_admin CASCADE;
CALL create_database('appdb', 'C.UTF-8');
SELECT * FROM outdated_extensions();
```

`create_database` creates a role with a random password when needed, creates the database, and grants privileges. `update_extensions()` performs cluster-wide extension updates rather than a read-only check.

These operations require broad administrative authority and can affect every database. Install in a dedicated administration database, revoke execution from `PUBLIC`, grant only to a controlled operator role, and capture generated credentials through an approved secret workflow. Review the SQL for dynamic identifiers and remote database connections. Take backups and test all extension update paths before calling `update_extensions()`; extension upgrades can run arbitrary migration SQL, take locks, change objects, or fail partway across the cluster. Do not expose these routines to application roles.
