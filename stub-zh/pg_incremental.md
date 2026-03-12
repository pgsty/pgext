


## 用法

> [pg_incremental: PostgreSQL 中的增量数据处理](https://github.com/CrunchyData/pg_incremental)

`pg_incremental` 扩展在 PostgreSQL 中提供快速、可靠的增量批处理流水线。它定义参数化查询，为新数据周期性执行，确保恰好一次投递。

```sql
CREATE EXTENSION pg_incremental CASCADE;  -- 依赖 pg_cron
```

### 流水线类型

共有三种类型的流水线：

- **序列流水线** -- 处理来自表的序列值范围
- **时间间隔流水线** -- 在时间间隔过后处理时间范围
- **文件列表流水线** -- 处理文件列表函数返回的新文件

### 序列流水线

创建一个使用序列增量聚合新行的流水线：

```sql
SELECT incremental.create_sequence_pipeline('event-aggregation', 'events', $$
  INSERT INTO events_agg
  SELECT date_trunc('day', event_time), count(*)
  FROM events
  WHERE event_id BETWEEN $1 AND $2
  GROUP BY 1
  ON CONFLICT (day) DO UPDATE SET event_count = events_agg.event_count + excluded.event_count
$$);
```

`$1` 和 `$2` 被设置为可以安全处理的最小和最大序列值。

带批次大小限制：

```sql
SELECT incremental.create_sequence_pipeline(
  'event-aggregation', 'events',
  $$ ... $$,
  schedule := '* * * * *',
  max_batch_size := 10000
);
```

### 时间间隔流水线

按固定时间间隔处理数据：

```sql
SELECT incremental.create_time_interval_pipeline('event-aggregation', '1 day', $$
  INSERT INTO events_agg
  SELECT event_time::date, count(distinct event_id)
  FROM events
  WHERE event_time >= $1 AND event_time < $2
  GROUP BY 1
$$);
```

`$1` 和 `$2` 被设置为时间范围的起始和结束（不包含）。

按间隔执行（例如每日导出）：

```sql
SELECT incremental.create_time_interval_pipeline('event-export',
  time_interval := '1 day',
  batched := false,
  start_time := '2024-11-01',
  command := $$ SELECT export_events($1, $2) $$
);
```

### 文件列表流水线

在新文件出现时处理它们：

```sql
SELECT incremental.create_file_list_pipeline('event-import', 's3://mybucket/events/*.csv', $$
  SELECT import_events($1)
$$);
```

### 管理函数

| 函数 | 描述 |
|------|------|
| `incremental.execute_pipeline(name)` | 手动执行流水线（仅在有新数据时） |
| `incremental.reset_pipeline(name)` | 重置流水线，从头开始重新处理 |
| `incremental.drop_pipeline(name)` | 删除流水线 |
| `incremental.skip_file(pipeline, path)` | 在文件列表流水线中跳过有问题的文件 |

### 监控

```sql
SELECT * FROM incremental.sequence_pipelines;
SELECT * FROM incremental.time_interval_pipelines;
SELECT * FROM incremental.processed_files;
```

通过 pg_cron 检查作业结果：

```sql
SELECT jobname, start_time, status, return_message
FROM cron.job_run_details JOIN cron.job USING (jobid)
WHERE jobname LIKE 'pipeline:%' ORDER BY 1 DESC LIMIT 5;
```
