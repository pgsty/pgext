## 用法

来源：

- [项目 README](https://github.com/postgresml/pg_stat_sysinfo/blob/b92c5161c22272aa955ad18e3d2e091d3ca6948c/README.md)
- [扩展 control 文件](https://github.com/postgresml/pg_stat_sysinfo/blob/b92c5161c22272aa955ad18e3d2e091d3ca6948c/pg_stat_sysinfo.control)
- [包元数据](https://github.com/postgresml/pg_stat_sysinfo/blob/b92c5161c22272aa955ad18e3d2e091d3ca6948c/Cargo.toml)
- [SQL 函数与视图](https://github.com/postgresml/pg_stat_sysinfo/blob/b92c5161c22272aa955ad18e3d2e091d3ca6948c/src/lib.rs)
- [采集器实现](https://github.com/postgresml/pg_stat_sysinfo/blob/b92c5161c22272aa955ad18e3d2e091d3ca6948c/src/collector.rs)
- [后台缓存 worker](https://github.com/postgresml/pg_stat_sysinfo/blob/b92c5161c22272aa955ad18e3d2e091d3ca6948c/src/cache_worker.rs)

`pg_stat_sysinfo` 0.0.1 是一个 pgrx 系统指标原型。它报告系统负载、CPU 用量、内存与 swap 容量和用量，以及已挂载磁盘容量。`pg_stat_sysinfo_collect()` 同步采样；可选后台 worker 把近期样本存入共享内存环形缓冲区，并通过 `pg_stat_sysinfo` 视图公开。

### 直接与缓存采集

直接采集不需要预加载：

```sql
CREATE EXTENSION pg_stat_sysinfo;

SELECT * FROM pg_stat_sysinfo_collect();
```

缓存采集需要配置库和间隔，然后重启 PostgreSQL：

```conf
shared_preload_libraries = 'pg_stat_sysinfo.so'
pg_stat_sysinfo.interval = '1s'
```

```sql
SELECT metric, dimensions, at, value
FROM pg_stat_sysinfo
ORDER BY at DESC;

SELECT * FROM pg_stat_sysinfo_cache_summary();
```

固定缓存约为 1280 KiB，因此保留时长随文件系统数量和采集间隔变化。配置 reload 可以修改间隔，并定期刷新磁盘列表。

### 可见性、开销与兼容性

这些指标会暴露主机负载、内存、swap、文件系统挂载名称和容量。应检查函数默认执行权限，并只向监控角色授权。直接调用会刷新操作系统计数器，并可能为系统库的最小 CPU 采样间隔短暂等待；过度轮询会增加后端和文件系统枚举工作。

审查到的 Cargo feature 使用 pgrx 0.8.3，覆盖 PostgreSQL 11 至 15，最后一次代码变更发生在 2023 年。没有证据证明它支持当前大版本或具有特殊挂载命名空间的容器。数值是某一时刻的主机观测，不是 PostgreSQL 归因，也不是持久监控。应验证单位、cgroup 可见性、磁盘发现、故障转移行为和采样开销，并把重要指标导出到受维护的监控系统。
