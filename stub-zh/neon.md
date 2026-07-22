## 用法

来源：

- [Neon 官方扩展文档](https://neon.com/docs/extensions/neon)
- [官方扩展 README](https://github.com/neondatabase/neon/blob/842a5091d5db4c23aeb29aea070c37ad06b12d63/pgxn/neon/README.md)
- [官方公开控制文件](https://github.com/neondatabase/neon/blob/842a5091d5db4c23aeb29aea070c37ad06b12d63/pgxn/neon/neon.control)

`neon` 暴露 Neon 计算节点统计信息和服务内部函数，其中最重要的是本地文件缓存指标。它属于 Neon 修改版 PostgreSQL 的计算栈；在普通 PostgreSQL 上只安装其 SQL 对象，并不会得到 Neon 存储、WAL safekeeper 或本地文件缓存。

### 核心流程

在 Neon 数据库中，应在需要查看视图的数据库里安装扩展，运行有代表性的工作负载，然后检查缓存行为：

```sql
CREATE EXTENSION neon;

SELECT file_cache_misses,
       file_cache_hits,
       file_cache_used,
       file_cache_writes,
       file_cache_hit_ratio
FROM neon_stat_file_cache;
```

`file_cache_hits` 统计未命中 PostgreSQL shared buffers 后由本地文件缓存提供的请求。`file_cache_misses` 统计必须访问 Neon 存储的请求。`file_cache_hit_ratio` 是本地文件缓存命中数在命中数与未命中数之和中的百分比，并不是 PostgreSQL shared-buffer 命中率。

### 重要对象

- `neon_stat_file_cache` 是有正式文档的用户视图，包含缓存未命中、命中、访问、写入与命中率。
- `pg_cluster_size` 按 Neon 存储语义报告集群大小。
- `approximate_working_set_size` 和 `approximate_working_set_size_seconds` 用于估算工作集大小。
- `get_local_cache_state`、`get_prewarm_info` 和 `prewarm_local_cache` 支持本地文件缓存状态捕获与预热。
- `backpressure_lsns`、`backpressure_throttling_time`、`neon_backend_perf_counters` 和 `neon_perf_counters` 暴露主要供 Neon 监控使用的服务内部信息。

服务支持时，应通过 `pg_monitor` 授予监控权限。不能因为预热或内部控制函数出现在扩展中，就把它们广泛开放。

### 生命周期与部署边界

这些统计覆盖整个计算节点，而不是单个数据库或表，并且只在该计算节点生命周期内存在。计算节点停止、重启或调整规格后统计会重置，因此应在运行代表性负载后再测量。较小工作集可能完全由 shared buffers 提供，因而几乎没有本地文件缓存观测。

核心 `neon` 共享库在服务器启动时加载，实现 Neon 的存储管理器、pageserver 通信、WAL proposer、控制面钩子和文件缓存。其预加载与服务配置由 Neon 管理；客户不应像普通可移植扩展一样自行添加或调优。

版本来源存在差异：目录记录为 1.9，而所引公开上游提交中的控制文件声明为 1.6。托管服务可能部署私有或更新的打包版本，因此应在目标 Neon 计算节点查询 `pg_extension.extversion`，不能仅根据公开控制文件推断服务构建。
