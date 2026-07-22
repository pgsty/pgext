## Usage

Sources:

- [Official README and warning](https://github.com/pierreforstmann/pg_dropcache/blob/e13cb92f52ce671e1e3cf2b29aa0f9908d8162ff/README.md)
- [Extension SQL and version gate](https://github.com/pierreforstmann/pg_dropcache/blob/e13cb92f52ce671e1e3cf2b29aa0f9908d8162ff/pg_dropcache--1.0.0.sql)
- [Buffer eviction implementation](https://github.com/pierreforstmann/pg_dropcache/blob/e13cb92f52ce671e1e3cf2b29aa0f9908d8162ff/pg_dropcache.c)

`pg_dropcache` 1.0.0 forcibly invalidates PostgreSQL shared buffers for the current database or one relation. It exists for destructive cache experiments on PostgreSQL 11–16. It can discard dirty pages without flushing them and therefore can cause data loss or corruption; never use it on a production or irreplaceable cluster.

### Isolated Test Workflow

Only after taking a recoverable copy and stopping all concurrent writers:

```sql
CREATE EXTENSION pg_dropcache;

CHECKPOINT;

SELECT pg_drop_rel_cache('bench.measurements'::regclass);

SELECT pg_drop_cache();
```

`pg_drop_rel_cache(regclass)` evicts buffers for all forks of one relation. `pg_drop_cache()` calls PostgreSQL's internal database-buffer eviction routine for the current database. `CHECKPOINT` reduces the dirty-page hazard but does not make concurrent use safe; new dirty buffers can appear immediately afterward.

### Safety and Version Boundaries

The functions use PostgreSQL private buffer-manager APIs and bypass normal durability expectations. Restrict execution to a dedicated test administrator, disconnect applications, and be prepared to recreate the cluster from a known-good copy. Cache experiments must also distinguish PostgreSQL `shared_buffers` from the operating-system page cache, which this extension does not clear.

The installation script rejects PostgreSQL versions below 11 and version 17 or later. On PostgreSQL 17+, upstream directs users to the maintained `pg_buffercache_evict`, `pg_buffercache_evict_relation`, and `pg_buffercache_evict_all` functions supplied by `pg_buffercache`. Do not remove the version gate or copy the old binary between server majors.
