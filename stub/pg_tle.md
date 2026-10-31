## Usage

Sources:

- [v1.5.3 README](https://github.com/aws/pg_tle/blob/v1.5.3/README.md)
- [v1.5.3 release](https://github.com/aws/pg_tle/releases/tag/v1.5.3)
- [1.5.2 to 1.5.3 SQL](https://github.com/aws/pg_tle/blob/v1.5.3/pg_tle--1.5.2--1.5.3.sql)
- [v1.5.3 control template](https://github.com/aws/pg_tle/blob/v1.5.3/pg_tle.control.in)

`pg_tle` 1.5.3 manages extensions written in trusted PostgreSQL languages. Native installation needs shared preload and a server restart; subsequent TLE definitions are stored in the database. Upstream documents PostgreSQL 12 through 19 Beta 4.

### Core Workflow

```ini
shared_preload_libraries = 'pg_tle'
```

```sql
CREATE EXTENSION pg_tle;
SELECT pgtle.available_extensions();
```

### Operational Boundaries

`pgtle.install_extension()` stores an extension definition; `CREATE EXTENSION` enables it in a database. `pgtle.install_update_path()` and `pgtle.set_default_version()` manage versions, while `pgtle.available_extensions()` lists definitions. `pgtle.set_extension_schema()` sets schema metadata for a managed extension in 1.5.3, which also tightens identifier validation. TLE administration is restricted through `pgtle_admin`; executing an installed extension still follows PostgreSQL language and object privileges. Dropping an installed extension and uninstalling its TLE definition are separate operations. Only use trusted extension code and review dependencies and upgrade SQL.
