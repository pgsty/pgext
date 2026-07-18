## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/pierreforstmann/pg_dropcache/blob/e13cb92f52ce671e1e3cf2b29aa0f9908d8162ff/README.md)
- [Version 1.0.0 SQL definitions](https://github.com/pierreforstmann/pg_dropcache/blob/e13cb92f52ce671e1e3cf2b29aa0f9908d8162ff/pg_dropcache--1.0.0.sql)
- [Extension control file](https://github.com/pierreforstmann/pg_dropcache/blob/e13cb92f52ce671e1e3cf2b29aa0f9908d8162ff/pg_dropcache.control)

`pg_dropcache` invalidates pages in PostgreSQL `shared_buffers`, either for the current database or for one relation. It is intended for controlled cache-sensitive tests and diagnostics.

```sql
CREATE EXTENSION pg_dropcache;

-- Evict buffers belonging to one relation.
SELECT pg_drop_rel_cache('bench_data'::regclass);

-- Evict buffers for the current database.
SELECT pg_drop_cache();
```

### Data-Loss Warning

The implementation discards dirty buffers without flushing them to disk. Calling either function can lose committed data and corrupt a test database; never run it on production or on data that must survive. Restrict execution privileges and use a disposable instance.

Upstream supports PostgreSQL 11 through 16. On PostgreSQL 17 or later, use `pg_buffercache_evict()`, `pg_buffercache_evict_relation()`, or `pg_buffercache_evict_all()` from the bundled `pg_buffercache` extension instead; the installation SQL rejects those server versions.
