


## 用法

来源：[README](https://github.com/fabriziomello/pg_stat_log/blob/main/README.md)，[SQL objects](https://github.com/fabriziomello/pg_stat_log/blob/main/pg_stat_log--0.1.sql)，[control file](https://github.com/fabriziomello/pg_stat_log/blob/main/pg_stat_log.control)

`pg_stat_log` 收集 PostgreSQL 日志消息的累积统计。它挂接 `emit_log_hook`，按后台进程类型、数据库、用户、严重级别、SQLSTATE 和 SQLSTATE 条件名统计日志消息。

### 启用

`pg_stat_log` 需要 PostgreSQL 18 或更新版本，并且必须预加载：

```conf
shared_preload_libraries = 'pg_stat_log'
```

重启 PostgreSQL 后创建扩展：

```sql
CREATE EXTENSION pg_stat_log;
```

### 查看统计

```sql
SELECT *
FROM pg_stat_log
ORDER BY count DESC;
```

视图包含 `backend_type`、`database_oid`、`database_name`、`user_oid`、`user_name`、`elevel`、`sqlerrcode`、`sqlerrcode_name`、`count`。

### 常用查询

```sql
SELECT elevel, sqlerrcode, sqlerrcode_name, sum(count) AS total
FROM pg_stat_log
GROUP BY elevel, sqlerrcode, sqlerrcode_name
ORDER BY total DESC
LIMIT 10;

SELECT backend_type, elevel, sqlerrcode_name, count
FROM pg_stat_log
WHERE backend_type <> 'client backend'
ORDER BY count DESC;
```

### 重置与容量

```sql
SELECT pg_stat_log_reset();
SELECT * FROM pg_stat_log_info();
```

`pg_stat_log_info()` 返回 `max_entries`、`num_entries`、`n_dropped`、`stats_reset`。如果 `n_dropped` 持续增长，应增大 `pg_stat_log.max_entries`。

### 配置

设置项包括 `pg_stat_log.enabled`、`pg_stat_log.min_error_level`、`pg_stat_log.max_entries`。

`emit_log_hook` 只能看到真正进入服务器日志的消息，因此 `log_min_messages` 是可统计消息的下限。
