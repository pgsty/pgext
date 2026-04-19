## 用法

- 来源：[README](https://github.com/CrunchyData/pg_incremental/blob/main/README.md)，[v1.5.0 release](https://github.com/CrunchyData/pg_incremental/releases/tag/v1.5.0)

`pg_incremental` 为 append-only table 和 file feed 定义 exactly-once 增量流水线。上游文档记录了三类 pipeline：sequence、time-interval 和 file-list。

### 安装与调度模型

上游 README 仍然使用基于 `pg_cron` 的调度模型，并通过下面的方式安装：

```sql
CREATE EXTENSION pg_incremental CASCADE;
```

除非显式指定 `execute_immediately := false`，否则 pipeline 会在创建时立刻运行一次，之后继续按 `pg_cron` 调度执行。README 还说明，即使没有新数据，每次调度执行也会出现在 `cron.job_run_details` 中。

### Sequence Pipelines

sequence pipeline 用于处理可安全消费的序列值范围：

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
```

README 记录了 `max_batch_size`，可用于限制每次运行处理的 sequence ID 数量。

### Time-Interval Pipelines

当命令希望把 `$1` 和 `$2` 作为时间区间边界接收时，可以使用时间窗口：

```sql
SELECT incremental.create_time_interval_pipeline('event-aggregation', '1 day', $$
  INSERT INTO events_agg
  SELECT event_time::date, count(DISTINCT event_id)
  FROM events
  WHERE event_time >= $1 AND event_time < $2
  GROUP BY 1
$$);
```

对于导出类任务，README 记录了 `batched := false`，这样每个时间区间都会单独执行。

### File-List Pipelines

file-list pipeline 用于处理新发现的文件：

```sql
SELECT incremental.create_file_list_pipeline('event-import', 's3://mybucket/events/*.csv', $$
  SELECT import_events($1)
$$);
```

v1.5.0 release 为 file-list pipeline 增加了 `max_batches_per_run`。README 还记录了 `incremental.skip_file()`，可将坏文件永久标记为已处理。

### 运维与监控

README 记录了以下接口：

- `CALL incremental.execute_pipeline(name)`：若存在新工作则执行一次。
- `SELECT incremental.reset_pipeline(name)`：重置进度。
- `SELECT incremental.drop_pipeline(name)`：删除 pipeline。
- `incremental.sequence_pipelines`、`incremental.time_interval_pipelines`、`incremental.file_list_pipelines` 与 `incremental.processed_files` 等视图和表。

v1.5.0 release note 还提到修复了在未安装 `pg_cron` 环境下的 `DROP EXTENSION` 问题。
