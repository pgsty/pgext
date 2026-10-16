## 用法

来源：

- [README.md](https://github.com/fabriziomello/pg_stat_log/blob/20b57e2e2eff884bfcf3ede3c095762e614c51d1/README.md)
- [pg_stat_log.control](https://github.com/fabriziomello/pg_stat_log/blob/20b57e2e2eff884bfcf3ede3c095762e614c51d1/pg_stat_log.control)
- [pg_stat_log--0.1.sql](https://github.com/fabriziomello/pg_stat_log/blob/20b57e2e2eff884bfcf3ede3c095762e614c51d1/pg_stat_log--0.1.sql)
- [pg_stat_log--0.1--0.3.sql](https://github.com/fabriziomello/pg_stat_log/blob/20b57e2e2eff884bfcf3ede3c095762e614c51d1/pg_stat_log--0.1--0.3.sql)

`pg_stat_log` 0.3 按后端进程、数据库、角色、严重程度和 SQLSTATE 汇总 PostgreSQL 18+ 日志消息，并可按查询标识符分组。将共享库加入既有预加载列表，重启后再创建 SQL 扩展。默认 SQL 授权允许 `pg_read_all_stats` 查看计数。

### 核心用法

```ini
shared_preload_libraries = 'pg_stat_log'
```

```sql
CREATE EXTENSION pg_stat_log;
SELECT database_name, user_name, query_id, elevel, sqlerrcode_name, count
FROM pg_stat_log
ORDER BY count DESC;
SELECT * FROM pg_stat_log_info();
```

### 运行边界

视图 `pg_stat_log` 在 `pg_stat_log_data()` 的基础上解析名称。`pg_stat_log_info()` 返回 `max_entries`、`num_entries`、`n_dropped` 和 `stats_reset`；`pg_stat_log_reset()` 清空计数并释放容量，其执行权限默认受限。`pg_stat_log.enabled` 和 `pg_stat_log.min_error_level` 控制收集，默认值分别为 true 和 warning。只有进入服务器日志的消息才会计数：`log_min_messages` 可以提前过滤消息，被捕获的异常也不会计入。并行工作进程的消息可能让同一次逻辑故障被多次统计。

`pg_stat_log.compute_query_id` 默认关闭，修改需要重启；PostgreSQL 的 `compute_query_id` 不能设为 off。没有查询标识符时显示 NULL。`pg_stat_log.max_entries` 默认 1024，修改需要重启，容量用尽后会丢弃新的组合，可通过 `n_dropped` 监测。修改容量会在下次启动时重置 PostgreSQL 的全部累积统计。计数可跨正常重启保留，丢弃遥测信息不会保留；崩溃恢复会丢弃累积统计。

0.3 将连接认证、授权和断开分别映射到 `PSL01`、`PSL02`、`PSL03`。收集需要开启对应的服务器日志，并设置 `pg_stat_log.min_error_level = 'log'` 或更低阈值；角色归属于登录角色。

上游 0.1 和 0.2 发布均安装 SQL 版本 0.1。升级到 0.3 需要替换匹配的共享库并重启，首次启动会重置 PostgreSQL 的全部累积统计。随后在每个已安装的数据库中执行 `ALTER EXTENSION pg_stat_log UPDATE TO '0.3'`。升级会重建统计视图和函数：需先删除再重建依赖视图，并重新授予自定义权限；默认的 `pg_read_all_stats` 授权会恢复。
