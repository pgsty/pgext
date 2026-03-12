


## 用法

> [pg_buffercache: 检查共享缓冲区缓存](https://www.postgresql.org/docs/current/pgbuffercache.html)

pg_buffercache 提供视图和函数来实时检查 PostgreSQL 共享缓冲区缓存中存储的内容。

### 视图

**`pg_buffercache`** -- 详细的每个缓冲区信息：

```sql
-- 按缓冲区使用量排名前 10 的关系
SELECT n.nspname, c.relname, count(*) AS buffers
FROM pg_buffercache b
JOIN pg_class c ON b.relfilenode = pg_relation_filenode(c.oid)
  AND b.reldatabase IN (0, (SELECT oid FROM pg_database WHERE datname = current_database()))
JOIN pg_namespace n ON n.oid = c.relnamespace
GROUP BY n.nspname, c.relname
ORDER BY 3 DESC
LIMIT 10;
```

列：`bufferid`、`relfilenode`、`reltablespace`、`reldatabase`、`relforknumber`、`relblocknumber`、`isdirty`、`usagecount`、`pinning_backends`。

### 汇总函数

```sql
-- 快速缓冲区缓存摘要（比视图开销更小）
SELECT * FROM pg_buffercache_summary();
--  buffers_used | buffers_unused | buffers_dirty | buffers_pinned | usagecount_avg

-- 按使用计数的缓冲区分布
SELECT * FROM pg_buffercache_usage_counts();
--  usage_count | buffers | dirty | pinned
```

### 驱逐函数（仅限超级用户，仅用于开发测试）

```sql
-- 按 ID 驱逐单个缓冲区
SELECT * FROM pg_buffercache_evict(42);

-- 驱逐某个关系的所有缓冲区
SELECT * FROM pg_buffercache_evict_relation('my_table'::regclass);

-- 驱逐所有未固定的缓冲区
SELECT * FROM pg_buffercache_evict_all();
```

### NUMA 视图

```sql
-- 共享缓冲区的 NUMA 节点映射
SELECT * FROM pg_buffercache_numa;
-- 返回：bufferid、os_page_num、numa_node
```

### 访问权限

限制为超级用户和拥有 `pg_monitor` 权限的角色。
