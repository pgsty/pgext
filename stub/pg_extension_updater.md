## Usage

Sources:

- [Official pg_extension_updater README](https://github.com/Snowflake-Labs/pg_lake/blob/v3.5.3/pg_extension_updater/README.md)
- [Version 3.5 control file](https://github.com/Snowflake-Labs/pg_lake/blob/v3.5.3/pg_extension_updater/pg_extension_updater.control)
- [Worker registration SQL](https://github.com/Snowflake-Labs/pg_lake/blob/v3.5.3/pg_extension_updater/pg_extension_updater--1.1.sql)
- [Updater implementation and configuration](https://github.com/Snowflake-Labs/pg_lake/blob/v3.5.3/pg_extension_updater/src/pg_extension_updater.c)
- [pg_extension_base preload documentation](https://github.com/Snowflake-Labs/pg_lake/blob/v3.5.3/pg_extension_base/README.md)

`pg_extension_updater` runs `ALTER EXTENSION ... UPDATE` for each extension whose installed version differs from its available default version when the database lifecycle worker starts. It is intended to reduce SQL/binary version mismatches after new extension files are deployed; it does not install missing extensions or replace release testing.

The pg_lake release and package version is `3.5.3`; the SQL extension version is `3.5`. Keep the libraries and query server from the same release.

### Enable Automatic Updates

The updater depends on `pg_extension_base`, which must be preloaded cluster-wide:

```conf
shared_preload_libraries = 'pg_extension_base'
```

After restarting PostgreSQL, create the updater in every database where automatic updates are desired:

```sql
CREATE EXTENSION pg_extension_updater CASCADE;
```

Creating it in `template1` causes new databases cloned from that template to contain the updater. The worker does not update extensions inside the template database itself, but starts in a newly cloned database.

### Runtime Behavior

- The `#!shared_preload_libraries` directive in the updater's control file lets `pg_extension_base` load its library.
- Installation registers the internal `extension_updater.main(internal)` lifecycle worker.
- At worker startup, it reads `pg_available_extensions` and updates entries whose `installed_version` differs from `default_version`.
- A failed `ALTER EXTENSION` is logged as a warning and is attempted only once during that worker run.
- `pg_extension_updater.enable` is a postmaster setting, defaults to `on`, and disables the worker's updates when set to `off`; changing it requires a restart.

### Caveats

- Automatic migration can execute arbitrary upgrade SQL supplied by every installed extension. Validate packages and upgrade paths before enabling it on production databases.
- Review extension dependency changes and take application-specific backups independently; a warning does not roll back unrelated successful extension updates.
- There is no user-facing force-update function or per-extension allowlist in version `3.5`.
- The `3.4` to `3.5` updater migration is empty; automatic updates still depend on reviewed upgrade scripts for each installed extension.
