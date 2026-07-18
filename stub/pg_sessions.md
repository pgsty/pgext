## Usage

Sources:

- [Pinned upstream README](https://github.com/menardorama/pg_sessions/blob/6471e4f563f14f77f5354081d46f40e011f0ca9/README.md)
- [Pinned version 1.3 installation SQL](https://github.com/menardorama/pg_sessions/blob/6471e4f563f14f77f5354081d46f40e011f0ca9/sql/pg_sessions--1.3.sql)
- [Pinned C implementation](https://github.com/menardorama/pg_sessions/blob/6471e4f563f14f77f5354081d46f40e011f0ca9/src/pg_sessions.c)
- [Pinned distribution metadata](https://github.com/menardorama/pg_sessions/blob/6471e4f563f14f77f5354081d46f40e011f0ca9/META.json)
- [Formal PGXN distribution page](https://pgxn.org/dist/pg_sessions/)

pg_sessions is a legacy session and normalized-statement statistics collector derived from old pg_stat_statements code. It keys entries by user, database, process ID, and query ID, and combines execution counters with Linux process CPU, memory, and I/O measurements.

### Preload and query

The module needs shared memory and executor hooks, so configure it before server start:

```conf
shared_preload_libraries = 'pg_sessions'
pg_sessions.max = 5000
pg_sessions.track = 'all'
pg_sessions.track_utility = on
pg_sessions.track_all_steps = on
pg_sessions.track_system_metrics = on
```

After restarting PostgreSQL:

```sql
CREATE EXTENSION pg_sessions;

SELECT datname, pid, queryid, status, calls, total_time,
       resident_memory_size, bytes_reads, bytes_writes
FROM pg_sessions
ORDER BY last_executed_timestamp DESC;
```

The view is granted to PUBLIC, but the C implementation masks query IDs and query text belonging to other users unless the caller is a superuser. The reset function is revoked from PUBLIC.

### Version and operational limits

The extension control and install SQL use object version 1.3, matching this catalog, while the repository META file and formal PGXN package identify distribution version 0.0.5. Upstream requires PostgreSQL 9.5 and says older versions do not work; there is no current compatibility statement for maintained PostgreSQL releases. The source uses obsolete server internals and must not be assumed to compile or run on a modern major version.

System metrics read Linux procfs and are not portable to other operating systems. Tracking every executor step adds hooks and write activity to all statements, while the shared hash and external query-text file consume memory and storage. The C defaults enable utility tracking, save-on-shutdown, all-step tracking, and system metrics, despite the README showing save disabled in its sample. Benchmark overhead, restrict view access, decide whether query text may be retained, and test crash/restart persistence before any operational use.
