## 用法

来源：

- [官方 README](https://github.com/CrunchyData/pg_incremental/blob/v1.6.1/README.md)
- [v1.6.1 发行版](https://github.com/CrunchyData/pg_incremental/releases/tag/v1.6.1)
- [控制文件](https://github.com/CrunchyData/pg_incremental/blob/v1.6.1/pg_incremental.control)

`pg_incremental` 为只追加表和文件数据源定义恰好一次语义的增量流水线。它依赖 `pg_cron`，上游核验矩阵覆盖 PostgreSQL 17–18，支持序列、时间区间和文件列表三类流水线。

### 安装与调度

```sql
CREATE EXTENSION pg_incremental CASCADE;
GRANT USAGE ON SCHEMA cron TO app_user;
```

流水线默认创建时立即执行，随后由 `pg_cron` 调度；需要推迟首次运行时设置 `execute_immediately := false`。进度更新与流水线命令在同一事务中提交，命令失败不会推进已处理范围。

### 流水线类型

```sql
SELECT incremental.create_sequence_pipeline('event-aggregation', 'events', $$
  INSERT INTO events_agg
  SELECT date_trunc('day', event_time), count(*)
  FROM events
  WHERE event_id BETWEEN $1 AND $2
  GROUP BY 1
  ON CONFLICT (day) DO UPDATE
  SET event_count = events_agg.event_count + excluded.event_count
$$);

SELECT incremental.create_time_interval_pipeline('hourly-export', '1 hour', $$
  SELECT export_events($1, $2)
$$);

SELECT incremental.create_file_list_pipeline('event-import', 's3://bucket/events/*.csv', $$
  SELECT import_events($1)
$$);
```

序列流水线传入安全的 `$1`/`$2` 范围，时间流水线传入区间边界，文件列表流水线传入新发现的文件。`max_batch_size`、`max_batches_per_run` 和 `batched := false` 可调整批次行为。`incremental.execute_pipeline`、`reset_pipeline`、`drop_pipeline` 以及状态视图用于运维控制。

### 1.6.1 安全修复

1.6.1 修复 `_drop_trigger` 的 `search_path` 注入问题，要求创建和执行序列流水线的角色拥有序列 `SELECT` 或 `USAGE` 权限，并让文件列表函数以调用者身份执行。升级前应复核对象所有权和授权；本地包事实与该上游发行版分开维护。
