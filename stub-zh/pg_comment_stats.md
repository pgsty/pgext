## 用法

来源：

- [Official documentation](https://github.com/munakoiso/pg_comment_stats/blob/69463b103e40682f6d50198240fa0c0cab045826/README.md)
- [Control file](https://github.com/munakoiso/pg_comment_stats/blob/69463b103e40682f6d50198240fa0c0cab045826/pg_comment_stats.control)
- [Version 1.0 SQL](https://github.com/munakoiso/pg_comment_stats/blob/69463b103e40682f6d50198240fa0c0cab045826/pg_comment_stats--1.0.sql)
- [Statistics readers](https://github.com/munakoiso/pg_comment_stats/blob/69463b103e40682f6d50198240fa0c0cab045826/pg_comment_stats.c)

`pg_comment_stats` 1.0 使用 SQL 注释中的键值标签聚合查询资源指标，将 `pg_stat_kcache` 计数与有界时间缓冲区结合，可跨多条 SQL 标识同一类应用操作。

### 前提与核心流程

源码构建还需要上游 pg_time_buffer 库源码和 pg_stat_kcache 头文件。保留已有预加载项，按依赖顺序加入三个模块后重启：

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

### 对象与配置

`pgcs_get_stats()` 和 `pgcs_get_stats_time_interval()` 返回标签、查询次数、用户及数据库 ID、I/O 字节数、CPU 时间、缺页和上下文切换计数。`pgcs_get_buffer_stats()` 返回已保存和可用的缓冲区记录数。`pgcs_exclude_key()`、`pgcs_get_excluded_keys()` 与 `pgcs_reset_excluded_keys()` 管理不参与聚合的标签。

`pg_comment_stats.buffer_size` 以 MB 设置共享内存容量；`pg_comment_stats.stat_time_interval` 以秒设置保留时间；`pg_comment_stats.excluded_keys` 提供初始排除列表。过期记录会离开缓冲区，因此它不是持久历史存储。

### 权限与限制

配置需要管理权限。安装 SQL 没有撤销读取函数和排除控制函数的公共执行权限，固定版本的 C 读取函数也不会按调用用户过滤返回的计数。在共享数据库中使用前应审查函数授权，且不要在注释标签中放入敏感数据。部分资源计数依赖操作系统支持。该源码版本没有声明完整的 PostgreSQL 主版本支持矩阵。
