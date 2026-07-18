## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/pierreforstmann/pg_dropcache/blob/e13cb92f52ce671e1e3cf2b29aa0f9908d8162ff/README.md)
- [1.0.0 版本 SQL 定义](https://github.com/pierreforstmann/pg_dropcache/blob/e13cb92f52ce671e1e3cf2b29aa0f9908d8162ff/pg_dropcache--1.0.0.sql)
- [扩展 control 文件](https://github.com/pierreforstmann/pg_dropcache/blob/e13cb92f52ce671e1e3cf2b29aa0f9908d8162ff/pg_dropcache.control)

`pg_dropcache` 使 PostgreSQL `shared_buffers` 中属于当前数据库或单个关系的页面失效。它仅适合在受控环境中进行与缓存有关的测试和诊断。

```sql
CREATE EXTENSION pg_dropcache;

-- Evict buffers belonging to one relation.
SELECT pg_drop_rel_cache('bench_data'::regclass);

-- Evict buffers for the current database.
SELECT pg_drop_cache();
```

### 数据丢失警告

该实现会直接丢弃脏缓冲区，不先写回磁盘。调用任一函数都可能丢失已提交数据并破坏测试数据库；绝不能用于生产环境或必须保留的数据。应限制执行权限，并且只在可丢弃实例中使用。

上游支持 PostgreSQL 11 至 16。PostgreSQL 17 或更高版本应改用内置 `pg_buffercache` 扩展提供的 `pg_buffercache_evict()`、`pg_buffercache_evict_relation()` 或 `pg_buffercache_evict_all()`；该扩展的安装 SQL 会拒绝这些服务端版本。
