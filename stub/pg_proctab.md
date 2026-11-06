## Usage

Sources:

- [pgnodemx 2.0.1 documentation](https://github.com/pgnodemx/pgnodemx/blob/v2.0.1/README.md)
- [Compatibility control file](https://github.com/pgnodemx/pgnodemx/blob/v2.0.1/pg_proctab.control)
- [Compatibility SQL](https://github.com/pgnodemx/pgnodemx/blob/v2.0.1/pg_proctab--0.0.10-compat.sql)

The `pg_proctab` shipped in `pgnodemx` provides five SQL wrappers for the legacy process-monitoring interface. It depends on `pgnodemx`; the package version is 2.0.1, while its SQL installation version is `0.0.10-compat`.

### Enable

Add `pgnodemx` to `shared_preload_libraries`, preserving other entries, then restart PostgreSQL:

```ini
shared_preload_libraries = 'pgnodemx'
```

Install both extensions in the same schema as a superuser:

```sql
CREATE EXTENSION pgnodemx;
CREATE EXTENSION pg_proctab VERSION '0.0.10-compat';
```

The explicit version is required because the main control file has no default version. The audited PGDG RPM packages omit that main control file; use a package containing it before running this command. Querying the underlying OS metrics requires superuser privileges or membership in `pg_monitor`.

### Query

```sql
SELECT a.pid, a.query, p.utime, p.stime, p.vsize, p.rss
FROM pg_stat_activity a
JOIN pg_proctab() p ON a.pid = p.pid;

SELECT "user", nice, system, idle, iowait FROM pg_cputime();
SELECT load1, load5, load15, last_pid FROM pg_loadavg();
SELECT memused, memfree, memshared, membuffers, memcached,
       swapused, swapfree, swapcached
FROM pg_memusage();
SELECT devname, reads_completed, writes_completed,
       sectors_read, sectors_written
FROM pg_diskusage();
```

These functions read Linux process and system statistics visible to the PostgreSQL server. Container namespaces and OS permissions affect what is visible. In this compatibility layer, `rss` is reported in KiB, memory totals in MiB, and `memshared` comes from `Shmem`.
