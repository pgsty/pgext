## Usage

Sources:

- [Version 0.0.1 installation SQL](https://github.com/abris-platform/pg_abris/blob/cca112ac51e23ee130be17a25cb0724a085aa58e/pg_abris--0.0.1.sql)
- [Extension control file](https://github.com/abris-platform/pg_abris/blob/cca112ac51e23ee130be17a25cb0724a085aa58e/pg_abris.control)
- [Upstream build notes](https://github.com/abris-platform/pg_abris/blob/cca112ac51e23ee130be17a25cb0724a085aa58e/README.md)

`pg_abris` is a pure PL/pgSQL metadata component for the Abris platform. It creates a fixed `meta` schema containing metadata tables, views, and helper functions, and requires `uuid-ossp`.

```sql
CREATE EXTENSION "uuid-ossp";
CREATE EXTENSION pg_abris;
SELECT nspname FROM pg_namespace WHERE nspname = 'meta';
```

Although the control file says relocatable, the SQL hard-codes `meta`, so moving the extension is unsafe. Upstream provides almost no user documentation, no release history, no PostgreSQL support statement, and no clear license file. Review every installed object and privilege in the SQL script, isolate it to an Abris-specific database, and do not treat version `0.0.1` as a general metadata framework.
