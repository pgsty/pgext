## 用法

来源：

- [官方上游 README](https://github.com/mbpcore/pg_pulse/blob/f5c1cf08ac37fcdc080ec502ebf97149e7f8c499/README.md)
- [官方扩展控制文件 (pg_pulse.control)](https://github.com/mbpcore/pg_pulse/blob/f5c1cf08ac37fcdc080ec502ebf97149e7f8c499/pg_pulse.control)
- [官方扩展 SQL (pg_pulse--1.0.sql)](https://github.com/mbpcore/pg_pulse/blob/f5c1cf08ac37fcdc080ec502ebf97149e7f8c499/pg_pulse--1.0.sql)

`pg_pulse` — 一个无锁的、内存中的 PostgreSQL 14–18 活动会话历史 (ASH) 抽样器，灵感来源于 Oracle 的 ASH，旨在实现低开销的生产环境中的持续会话级监控。在收集或解释相应的 PostgreSQL 统计信息时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_pulse;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pg_pulse_conn_ring_sample(OUT sample_time timestamptz, OUT total_connections int, OUT active_connections int, OUT idle_connections int, OUT idle_in_transaction_connections int, OUT idle_in_transaction_aborted_connections int, OUT other_connections int)` 是一个扩展函数，返回 `SETOF`。
- `pg_pulse_purge_history(retention_days int DEFAULT NULL)` 是一个扩展函数，返回 `void`。
- `pg_pulse_connections_live` 是一个由扩展定义的视图。
- `pg_pulse_live` 是一个由扩展定义的视图。
- `pg_pulse_connection_history` 是一个由扩展安装或管理的表。
- `pg_pulse_history` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
