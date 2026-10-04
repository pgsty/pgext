## Usage

Sources:

- [Version 0.2.5 README](https://github.com/pgexporter/pgexporter_ext/blob/0.2.5/README.md)
- [Version 0.2.5 setup guide](https://github.com/pgexporter/pgexporter_ext/blob/0.2.5/doc/GETTING_STARTED.md)
- [Version 0.2.5 SQL definitions](https://github.com/pgexporter/pgexporter_ext/tree/0.2.5/sql)
- [Version 0.2.3 SQL definitions](https://github.com/pgexporter/pgexporter_ext/tree/0.2.3/sql)
- [Version 0.2.4 SQL definitions](https://github.com/pgexporter/pgexporter_ext/tree/0.2.4/sql)
- [Version 0.2.5 filesystem implementation](https://github.com/pgexporter/pgexporter_ext/blob/0.2.5/src/pgexporter_ext/utils.c)

`pgexporter_ext` exposes Linux host and filesystem metrics through SQL for collection by pgexporter. The examples below apply to the packaged 0.2.3–0.2.5 APIs. Functions run with the PostgreSQL operating-system account's access, so restrict the monitoring role to trusted users.

### Setup and Core Workflow

The upstream setup guide configures the module in `shared_preload_libraries`, followed by a PostgreSQL restart. Preserve any other libraries already configured.

```ini
shared_preload_libraries = 'pgexporter_ext'
```

In the `postgres` database, a privileged administrator installs the extension and grants monitoring access to the existing exporter login:

```sql
CREATE EXTENSION pgexporter_ext;
GRANT pg_monitor TO pgexporter;

SET ROLE pgexporter;
SELECT pgexporter_version_ext();
SELECT * FROM pgexporter_get_functions();
SELECT * FROM pgexporter_os_info();
SELECT * FROM pgexporter_load_avg();
RESET ROLE;
```

The SQL scripts revoke public execution and grant it to `pg_monitor`. That predefined role also grants broad PostgreSQL monitoring access. The release includes a base installation script and an upgrade chain; PostgreSQL can use that chain when installing the default version with `CREATE EXTENSION`.

### Metric Functions

- `pgexporter_information_ext()` and `pgexporter_version_ext()` return extension information and version text.
- `pgexporter_get_functions()` lists metric names, whether they take input, descriptions, and types. `pgexporter_is_supported(text)` checks a metric name.
- `pgexporter_os_info()`, `pgexporter_cpu_info()`, `pgexporter_memory_info()`, `pgexporter_network_info()`, and `pgexporter_load_avg()` return host metrics.
- `pgexporter_used_space(text)`, `pgexporter_free_space(text)`, and `pgexporter_total_space(text)` return byte counts for a filesystem path.

```sql
SELECT pgexporter_is_supported('pgexporter_load_avg');
SELECT pgexporter_free_space('/var/lib/postgresql');
```

Choose a path that exists on the server and is accessible to its operating-system account. These releases do not provide the later FIPS or log-count APIs.

### Operational Boundaries

Filesystem calls use the server's filesystem view, which may be a container rather than the host. The used-space function walks directories, so large trees can make scrapes expensive. A `pg_monitor` member can inspect accessible paths; restrict database access and choose collection paths deliberately.

Upstream lists Linux and PostgreSQL 13+; this catalog supplies packages for PostgreSQL 14–18. Use the exact PostgreSQL-major package. The principal RPM supplier is PGDG at 0.2.4 (EL8 PostgreSQL 14–16 use 0.2.3), while Pigsty supplies DEB 0.2.5; enable the Pigsty repository as well when following the cross-platform installation instructions.
