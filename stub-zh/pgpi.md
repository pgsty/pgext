## 用法

来源：

- [官方源码 1](https://github.com/s-hironobu/pg_plan_inspector/blob/7d929b64421f229800fe84838c36a414e7159de4/README.md)
- [官方源码 2](https://github.com/s-hironobu/pg_plan_inspector/blob/7d929b64421f229800fe84838c36a414e7159de4/extension/pgpi.control)
- [官方源码 3](https://github.com/s-hironobu/pg_plan_inspector/blob/7d929b64421f229800fe84838c36a414e7159de4/extension/pgpi--0.4.sql)
- [官方源码 4](https://github.com/s-hironobu/pg_plan_inspector/blob/7d929b64421f229800fe84838c36a414e7159de4/extension/Makefile)
- [官方源码 5](https://github.com/s-hironobu/pg_plan_inspector/blob/7d929b64421f229800fe84838c36a414e7159de4/extension/pgpi.c)

`pgpi` 是 `pg_plan_inspector` 研究框架中的 PostgreSQL 扩展。控制文件和安装 SQL 使用 0.4 版本，框架将这一系列称为 2.0。它记录执行遥测，并可用静态指令或外部估计器的结果替换规划器基数估计。

### 启用与检查

当前源码要求 PostgreSQL 19，支持 x86_64/ARM64 的 Linux 与 macOS；运行时函数跳转替换仅在 Linux 可用，macOS 仍可使用官方钩子路径。启用预加载并指定日志数据库后，重启 PostgreSQL：

```ini
shared_preload_libraries = 'pgpi'
pgpi.log_database = 'postgres'
```

已有预加载项时应追加，而非覆盖。在需要使用扩展的数据库中，以超级用户执行：

```sql
CREATE EXTENSION pgpi;
SELECT * FROM pgpi_status();
SELECT pgpi_log_status();
```

后台进程在日志数据库中创建 `pgpi.plans` 和 `pgpi.stats`。SQL 函数支持迁移模式，日志表则保留在自己的模式内。

### 采集与测试估计值

干预默认关闭，计划采集单独控制。具备权限的测试会话可以采集负载，并在收集前刷出缓冲区：

```sql
SET pgpi.log_plans = on;
SET pgpi.label = 'baseline';
SELECT count(*) FROM pg_class;
SELECT pgpi_log_flush(10000);
SET pgpi.log_plans = off;
```

`pgpi.log_plans` 需要超级用户权限或显式的参数授权。Python 客户端需要 Python 3.10+ 及所选估计器的依赖；它将日志收集到 Parquet 数据集，再通过 UNIX 套接字提供估计值。将 `pgpi.callback` 指向该套接字，用 `pgpi.mode = 'shadow'` 记录回答而不应用，`enforce` 则应用回答。也可通过 `pgpi.directives` 提供静态结果。

### 对象与边界

- `pgpi_requests` 与 `pgpi_clear` 用于查看或清空当前后端的请求日志。`pgpi_status` 报告已安装的跳转替换，`pgpi_log_status` 展示共享缓冲区及后台进程状态。
- `pgpi_log_flush` 超时或未预加载时返回 false。`pgpi_stats_reset_seen` 重置已记住的统计信息版本，默认不允许 `PUBLIC` 执行。
- 回调超时默认 2000 ms。失败后暂停请求五秒，并采用 PostgreSQL 自身的估计值；估计器也可以主动不作答。
- 日志默认无限保留；可通过 `pgpi.log_retention` 或 `pgpi.log_max_rows` 限制计划存储。更改 `pgpi.log_database`、`pgpi.log_buffer_size` 和 `pgpi.stats_max` 需要重启。
- 日志可能包含 SQL、绑定参数及统计信息，应限制访问。采集会增加执行监测开销；估计干预和跳转替换属于实验功能，可能改变计划或性能。项目并未保证更准确的基数估计一定加快查询。
