## 用法

来源：

- [EDB Wait States 官方文档](https://www.enterprisedb.com/docs/pg_extensions/wait_states/)
- [官方安装指南](https://www.enterprisedb.com/docs/pg_extensions/wait_states/installing/)
- [官方使用指南](https://www.enterprisedb.com/docs/pg_extensions/wait_states/using/)

`edb_wait_states` 随时间采样 backend wait event，并持久保存 query、session、system 与 wait-event 数据用于性能分析。

### 启用

安装 EDB Wait States 1.6.0，预加载其 worker，重启并创建扩展：

```ini
shared_preload_libraries = 'edb_wait_states'
edb_wait_states.sampling_interval = '1s'
edb_wait_states.retention_period = 604800
```

```sql
CREATE EXTENSION edb_wait_states;
```

若 EDB 诊断工具需要读取数据，应向 `pg_monitor` 授予安装 schema 的 `USAGE`。

### 分析样本

```sql
SELECT query, session_id, wait_event_type, wait_event
FROM edb_wait_states_data(
  now() - interval '15 minutes',
  now()
);

SELECT * FROM edb_wait_states_wait_events();
SELECT edb_wait_states_directory_size();
```

`edb_wait_states_queries`、`edb_wait_states_sessions` 与 `edb_wait_states_sql_statements` 提供采样文件上的高层视图。

### 保留与隐私

样本保存在 `edb_wait_states.directory` 下；修改路径需要重启。采样会增加磁盘使用量，也可能保留 query text、role、database 与 session detail。应设置有限 retention period、保护目录与 SQL 函数、监控大小，并在策略要求提前删除时使用 `edb_wait_states_purge`。

