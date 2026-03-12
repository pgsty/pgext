

## Usage

> [pg_buffercache: inspect the shared buffer cache](https://www.postgresql.org/docs/current/pgbuffercache.html)

pg_buffercache provides views and functions to examine what is stored in the PostgreSQL shared buffer cache in real time.

### Views

**`pg_buffercache`** -- detailed per-buffer information:

```sql
-- Top 10 relations by buffer usage
SELECT n.nspname, c.relname, count(*) AS buffers
FROM pg_buffercache b
JOIN pg_class c ON b.relfilenode = pg_relation_filenode(c.oid)
  AND b.reldatabase IN (0, (SELECT oid FROM pg_database WHERE datname = current_database()))
JOIN pg_namespace n ON n.oid = c.relnamespace
GROUP BY n.nspname, c.relname
ORDER BY 3 DESC
LIMIT 10;
```

Columns: `bufferid`, `relfilenode`, `reltablespace`, `reldatabase`, `relforknumber`, `relblocknumber`, `isdirty`, `usagecount`, `pinning_backends`.

### Summary Functions

```sql
-- Quick buffer cache summary (cheaper than the view)
SELECT * FROM pg_buffercache_summary();
--  buffers_used | buffers_unused | buffers_dirty | buffers_pinned | usagecount_avg

-- Buffer distribution by usage count
SELECT * FROM pg_buffercache_usage_counts();
--  usage_count | buffers | dirty | pinned
```

### Eviction Functions (Superuser, Developer Testing Only)

```sql
-- Evict a single buffer by ID
SELECT * FROM pg_buffercache_evict(42);

-- Evict all buffers for a relation
SELECT * FROM pg_buffercache_evict_relation('my_table'::regclass);

-- Evict all unpinned buffers
SELECT * FROM pg_buffercache_evict_all();
```

### NUMA View

```sql
-- NUMA node mapping for shared buffers
SELECT * FROM pg_buffercache_numa;
-- Returns: bufferid, os_page_num, numa_node
```

### Access

Restricted to superusers and roles with `pg_monitor` privileges.
