## Usage

Sources:

- [Official documentation](https://github.com/munakoiso/pg_comment_stats/blob/69463b103e40682f6d50198240fa0c0cab045826/README.md)
- [Control file](https://github.com/munakoiso/pg_comment_stats/blob/69463b103e40682f6d50198240fa0c0cab045826/pg_comment_stats.control)
- [Version 1.0 SQL](https://github.com/munakoiso/pg_comment_stats/blob/69463b103e40682f6d50198240fa0c0cab045826/pg_comment_stats--1.0.sql)
- [Statistics readers](https://github.com/munakoiso/pg_comment_stats/blob/69463b103e40682f6d50198240fa0c0cab045826/pg_comment_stats.c)

`pg_comment_stats` 1.0 groups query resource measurements using key-value labels in SQL comments. It combines counters from `pg_stat_kcache` with a bounded time buffer; labels can identify an application operation across several SQL statements.

### Prerequisites and Core Workflow

The source build also needs the upstream pg_time_buffer library sources and the pg_stat_kcache header. Preserve existing preload entries, add the three modules in dependency order, and restart:

```conf
shared_preload_libraries = 'pg_stat_statements,pg_stat_kcache,pg_comment_stats'
```

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION pg_stat_kcache;
CREATE EXTENSION pg_comment_stats;
/* service: orders operation: lookup */ SELECT 1;
SELECT * FROM pgcs_get_stats();
SELECT * FROM pgcs_get_stats_time_interval(now() - interval '30 seconds', now());
```

### Objects and Configuration

`pgcs_get_stats()` and `pgcs_get_stats_time_interval()` return labels, query count, user/database IDs, I/O bytes, CPU times, page faults, and context-switch counters. `pgcs_get_buffer_stats()` reports saved and available buffer entries. `pgcs_exclude_key()`, `pgcs_get_excluded_keys()`, and `pgcs_reset_excluded_keys()` manage labels excluded from aggregation.

`pg_comment_stats.buffer_size` sets shared-memory capacity in MB; `pg_comment_stats.stat_time_interval` sets retention in seconds; `pg_comment_stats.excluded_keys` supplies an initial exclusion list. Expired entries leave the buffer, so this is not durable historical storage.

### Privileges and Limitations

Setup requires administrative access. The installation SQL does not revoke public execution from its readers or exclusion controls, and the pinned C reader does not filter returned counters by the calling user. Review function grants before use in a shared database, and keep sensitive data out of comment labels. Some resource counters depend on operating-system support. No complete supported PostgreSQL-major matrix is declared by this source revision.
