## Usage

Sources:

- [Official README](https://github.com/pgexporter/pgexporter_ext/blob/c73074eb5b7aa2768e1a41045905c01ebcdab197/README.md)
- [Official getting-started guide](https://github.com/pgexporter/pgexporter_ext/blob/c73074eb5b7aa2768e1a41045905c01ebcdab197/doc/GETTING_STARTED.md)
- [Extension control file and SQL scripts](https://github.com/pgexporter/pgexporter_ext/tree/c73074eb5b7aa2768e1a41045905c01ebcdab197/sql)
- [Metric implementation](https://github.com/pgexporter/pgexporter_ext/blob/c73074eb5b7aa2768e1a41045905c01ebcdab197/src/pgexporter_ext/lib.c)
- [Filesystem and log implementation](https://github.com/pgexporter/pgexporter_ext/blob/c73074eb5b7aa2768e1a41045905c01ebcdab197/src/pgexporter_ext/utils.c)

`pgexporter_ext` exposes Linux host, filesystem, PostgreSQL log, and FIPS information through SQL for collection by pgexporter. Its functions execute inside PostgreSQL with the server operating-system account's access, so treat the granted monitoring role as host-observation privilege rather than ordinary database read access.

### Setup and Core Workflow

Follow upstream's supported setup: preload the module, restart PostgreSQL, install it in the `postgres` database, and grant the built-in monitoring role to the exporter login.

```conf
shared_preload_libraries = 'pgexporter_ext'
```

```sql
CREATE EXTENSION pgexporter_ext;
GRANT pg_monitor TO pgexporter;

SET ROLE pgexporter;
SELECT pgexporter_ext_version();
SELECT * FROM pgexporter_ext_get_functions();
SELECT * FROM pgexporter_ext_os_info();
SELECT * FROM pgexporter_ext_load_avg();
```

The installation scripts revoke public execution and grant the functions to `pg_monitor`. Review the consequences before granting that predefined role; it already includes broad PostgreSQL monitoring capabilities.

### Metric Families

- `pgexporter_ext_version()`, `pgexporter_ext_is_supported(text)`, and `pgexporter_ext_get_functions()` provide discovery and capability metadata.
- `pgexporter_ext_os_info()`, `pgexporter_ext_cpu_info()`, `pgexporter_ext_memory_info()`, `pgexporter_ext_network_info()`, and `pgexporter_ext_load_avg()` expose host information.
- `pgexporter_ext_used_space(text)`, `pgexporter_ext_free_space(text)`, and `pgexporter_ext_total_space(text)` inspect caller-supplied filesystem paths.
- `pgexporter_ext_fips()` reports whether the linked OpenSSL context is in FIPS mode.
- The `pgexporter_ext_log_*()` family counts occurrences of PostgreSQL severity strings from `DEBUG5` through `PANIC` in the configured log directory.

### Cost and Security Boundaries

The path functions traverse directories or call filesystem-stat APIs as the PostgreSQL operating-system user. A `pg_monitor` member can probe arbitrary paths accessible to that account and infer their existence or size. Restrict database login and network access accordingly.

Every log-count function synchronously scans every regular file in `log_directory`, including gzip, bzip2, LZ4, and Zstandard files, and counts severity substrings rather than parsing structured PostgreSQL log records. Frequent scrapes or a large retained log set can impose severe CPU, I/O, decompression, and query-latency costs. The current source defines `pgexporter.log_cache_refresh_interval`, but the log functions still call the full scan directly; do not assume the GUC provides effective caching.

### Version and Packaging Caveat

The reviewed control file declares version `0.3.1`, but the same source tree contains only a base `pgexporter_ext--0.1.0.sql` plus incremental upgrade scripts; it has no base `pgexporter_ext--0.3.1.sql`. A direct `CREATE EXTENSION pgexporter_ext` from an unmodified current source installation can therefore fail. Verify that the distributed package supplies a valid base script, or install version `0.1.0` and test the complete `ALTER EXTENSION ... UPDATE` chain in a disposable database before rollout.

Upstream reports PostgreSQL 13+ on Linux and the current build declares C17 plus OpenSSL and compression-library dependencies. Use a package built for the exact PostgreSQL major and operating system, then validate function availability and scrape cost on the actual host.
