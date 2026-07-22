## Usage

Sources:

- [Official README](https://github.com/04ina/buffercache_tools/blob/90f4edbc7ba3fb84dd43853e0cf030beb46f226a/README.md)
- [Official extension SQL](https://github.com/04ina/buffercache_tools/blob/90f4edbc7ba3fb84dd43853e0cf030beb46f226a/buffercache_tools--1.0.sql)
- [Official C implementation](https://github.com/04ina/buffercache_tools/blob/90f4edbc7ba3fb84dd43853e0cf030beb46f226a/buffercache_tools.c)

`buffercache_tools` exposes low-level PostgreSQL shared-buffer inspection and mutation helpers. It is intended for controlled debugging and fault-injection work; functions that rewrite tags, dirty, flush, or invalidate buffers can break consistency and must not be used as routine administration tools.

### Core Workflow

Install the extension as a superuser, then start with read-only inspection:

```sql
CREATE EXTENSION buffercache_tools;

SELECT * FROM pg_show_buffer(42);
SELECT * FROM pg_show_relation_buffers('public.orders');
```

`pg_show_buffer(integer)` describes one shared buffer. `pg_show_relation_buffers(text)` lists buffers for a relation and can also report the current session's local temporary buffers. `pg_read_page_into_buffer(text, text, bigint)` reads a selected relation fork and block into the buffer cache.

### Mutation Surface and Safety

The mutation family includes `pg_change_buffer`, `pg_change_relation_fork_buffers`, `pg_change_relation_buffers`, `pg_change_database_buffers`, `pg_change_tablespace_buffers`, `pg_change_all_valid_buffers`, and `pg_change_buffer_by_page`. Supported modes include `mark_dirty`, `flush`, `invalidate`, `change_spcoid`, `change_dboid`, `change_relnumber`, `change_forknum`, and `change_blocknum`.

The implementation requires superuser access for inspection and mutation helpers. It rejects buffers belonging to another session's temporary relation, and temporary buffers cannot be mutated. Even in a test cluster, take a recoverable copy first, isolate concurrent traffic, and expect that tag changes or invalidation can cause lost writes, wrong-page reads, crashes, or restart recovery. Prefer a disposable cluster when exercising the mutation API.
