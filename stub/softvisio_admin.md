## Usage

Sources:

- [Upstream source](https://github.com/softvisio/postgresql-softvisio-admin/blob/v1.8.0/softvisio_admin.control)
- [Upstream source](https://github.com/softvisio/postgresql-softvisio-admin/blob/v1.8.0/softvisio_admin--1.7.1.sql)
- [Upstream source](https://github.com/softvisio/postgresql-softvisio-admin/blob/v1.8.0/softvisio_admin--1.7.0--1.7.1.sql)
- [Upstream source](https://github.com/softvisio/postgresql-softvisio-admin/blob/v1.8.0/Makefile)

- [Control](https://github.com/softvisio/postgresql-softvisio-admin/blob/v1.7.0/softvisio_admin.control)
- [Version 1.7.0 SQL](https://github.com/softvisio/postgresql-softvisio-admin/blob/v1.7.0/softvisio_admin--1.7.0.sql)
- [Installed SQL files](https://github.com/softvisio/postgresql-softvisio-admin/blob/v1.7.0/Makefile)
- [Changelog](https://github.com/softvisio/postgresql-softvisio-admin/blob/v1.7.0/CHANGELOG.md)

`softvisio_admin` 1.7.0 provides SQL/PL/pgSQL helpers for database and role provisioning and cluster-wide extension inventory or updates. Installation requires a superuser, `dblink` and `pgcrypto`; no native library or preload is required.

### Provision and inspect

The actual 1.7.0 SQL exposes `create_database(name text)` with one argument. It creates a missing login role with a generated password, creates a UTF8 database from `template0` using ICU locale `und` and `C.UTF-8`, then assigns ownership and creates its schema. The server must support this locale configuration. The older README's optional collation argument does not match this SQL version.

```sql
CREATE EXTENSION softvisio_admin CASCADE;
CALL create_database('appdb');
SELECT * FROM outdated_extensions();
```

### Cross-database operations

`outdated_extensions()` visits connectable, non-template databases through dblink and compares installed and default versions. `update_extensions()` performs updates across those databases, using the PostGIS upgrade helper where appropriate. It can report remote errors as notices and continue; inspect returned installed versions and notices. Remote statements commit independently, so failure does not roll back earlier databases or provisioning steps.

### Upgrade and access boundaries

The 1.7.0 source installs only the new base SQL file and removes historical upgrade scripts. An existing older installation therefore has no ordinary update path supplied by this distribution alone. Check `pg_extension_update_paths('softvisio_admin')` and preserve a reviewed migration plan before changing server files; do not assume `ALTER EXTENSION UPDATE` will succeed. Keep these routines restricted to trusted administrators. They issue powerful remote SQL, and a newly generated password is printed in a NOTICE; keep client and server logging access controlled. Review and test every affected extension's upgrade before invoking the cluster-wide updater.

### Distribution 1.8.0 and SQL 1.7.1

The 1.8.0 release retains default SQL version 1.7.1 and the same one-argument create_database interface. Its Makefile now supplies the 1.7.0-to-1.7.1 upgrade script, which is empty; that specific update path is available after installing matching files. The README still advertises an optional collation parameter absent from this SQL. This does not supply historical paths from versions older than 1.7.0 or make cross-database operations atomic.
