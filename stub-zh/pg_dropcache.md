## 用法

来源：

- [官方 README 与警告](https://github.com/pierreforstmann/pg_dropcache/blob/e13cb92f52ce671e1e3cf2b29aa0f9908d8162ff/README.md)
- [扩展 SQL 与版本门禁](https://github.com/pierreforstmann/pg_dropcache/blob/e13cb92f52ce671e1e3cf2b29aa0f9908d8162ff/pg_dropcache--1.0.0.sql)
- [缓冲区驱逐实现](https://github.com/pierreforstmann/pg_dropcache/blob/e13cb92f52ce671e1e3cf2b29aa0f9908d8162ff/pg_dropcache.c)

`pg_dropcache` 1.0.0 强制失效当前数据库或单个关系的 PostgreSQL 共享缓冲区。它用于 PostgreSQL 11–16 上的破坏性缓存实验。它可能在不刷盘时丢弃脏页，从而导致数据丢失或损坏；绝不能用于生产或不可替代的集群。

### 隔离测试流程

只有在取得可恢复副本并停止全部并发写入之后才可执行：

```sql
CREATE EXTENSION pg_dropcache;

CHECKPOINT;

SELECT pg_drop_rel_cache('bench.measurements'::regclass);

SELECT pg_drop_cache();
```

`pg_drop_rel_cache(regclass)` 驱逐单个关系所有 fork 的缓冲区。`pg_drop_cache()` 调用 PostgreSQL 内部数据库缓冲区驱逐例程，作用于当前数据库。`CHECKPOINT` 可以降低脏页风险，但不能让并发使用变得安全；之后可能立即出现新的脏缓冲区。

### 安全与版本边界

这些函数使用 PostgreSQL 私有缓冲区管理 API，并绕过正常持久性预期。应只向专用测试管理员开放，断开应用连接，并准备从已知良好副本重建集群。缓存实验还必须区分 PostgreSQL `shared_buffers` 与操作系统页缓存；此扩展不会清除后者。

安装脚本拒绝低于 11 的 PostgreSQL，也拒绝 17 或更高版本。在 PostgreSQL 17+ 上，上游要求改用由 `pg_buffercache` 提供并维护的 `pg_buffercache_evict`、`pg_buffercache_evict_relation` 和 `pg_buffercache_evict_all`。不要移除版本门禁，也不要跨服务器大版本复制旧二进制文件。
