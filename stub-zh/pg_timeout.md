## 用法

来源：

- [官方 README](https://github.com/pierreforstmann/pg_timeout/blob/cabf38c88b48da64e84d226793da762e5c6e14a1/README.md)
- [官方扩展 SQL](https://github.com/pierreforstmann/pg_timeout/blob/cabf38c88b48da64e84d226793da762e5c6e14a1/pg_timeout--1.0.sql)
- [官方工作进程实现](https://github.com/pierreforstmann/pg_timeout/blob/cabf38c88b48da64e84d226793da762e5c6e14a1/pg_timeout.c)

`pg_timeout` 是一个已归档的后台工作进程，会终止 `pg_stat_activity.state` 持续超过全局阈值的 `idle` 集群会话。当前 PostgreSQL 已提供内置 `idle_session_timeout`；除非维护旧服务器，否则应优先使用受支持的内置设置。

### 启用方式

实际行为来自预加载动态库，而不是调用其 SQL 函数。配置后重启 PostgreSQL：

```ini
shared_preload_libraries = 'pg_timeout'
pg_timeout.naptime = 30
pg_timeout.idle_session_timeout = 300
```

重启后可以创建扩展，以登记 `pg_timeout_main()` SQL 对象，但普通用户绝不能直接调用这个工作进程入口：

```sql
CREATE EXTENSION pg_timeout;
SHOW pg_timeout.naptime;
SHOW pg_timeout.idle_session_timeout;
```

`pg_timeout.naptime` 是扫描间隔秒数；`pg_timeout.idle_session_timeout` 是空闲阈值秒数。两者均可重载。

### 终止语义

工作进程连接本地 `postgres` 数据库，扫描集群级 `pg_stat_activity` 视图，记录匹配会话的元数据，并对状态严格等于 `idle` 且 `state_change` 更早的每个会话调用 `pg_terminate_backend()`。它特意不会终止 `idle in transaction` 会话。

### 运维边界

阈值作用于整个集群：没有按角色、数据库、应用或白名单设置的例外，因此连接池、管理员、复制工具和维护客户端都可能被断开。终止按休眠间隔采样，不会精确到秒。工作进程使用较高的终止权限，并记录用户名、数据库、应用名和主机名。启用前应测试重连行为与例外需求，监控重启循环，并另用 `idle_in_transaction_session_timeout` 处理遗弃事务。在受支持的 PostgreSQL 版本上，应改用角色/数据库范围的内置超时设置。
