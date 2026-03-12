


## Usage

> [pg_incremental: Incremental Data Processing in PostgreSQL](https://github.com/CrunchyData/pg_incremental)

The `pg_incremental` extension provides fast, reliable incremental batch processing pipelines in PostgreSQL. It defines parameterized queries that execute periodically for new data, ensuring exactly-once delivery.

```sql
CREATE EXTENSION pg_incremental CASCADE;  -- depends on pg_cron
```

### Pipeline Types

There are three types of pipelines:

- **Sequence pipelines** -- Process ranges of sequence values from a table
- **Time interval pipelines** -- Process time ranges after intervals pass
- **File list pipelines** -- Process new files from a file listing function

### Sequence Pipeline

Create a pipeline that incrementally aggregates new rows using a sequence:

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

`$1` and `$2` are set to the lowest and highest sequence values that can be safely processed.

With batch size limiting:

```sql
SELECT incremental.create_sequence_pipeline(
  'event-aggregation', 'events',
  $$ ... $$,
  schedule := '* * * * *',
  max_batch_size := 10000
);
```

### Time Interval Pipeline

Process data in fixed time intervals:

```sql
SELECT incremental.create_time_interval_pipeline('event-aggregation', '1 day', $$
  INSERT INTO events_agg
  SELECT event_time::date, count(distinct event_id)
  FROM events
  WHERE event_time >= $1 AND event_time < $2
  GROUP BY 1
$$);
```

`$1` and `$2` are set to the start and end (exclusive) of the time range.

For per-interval execution (e.g., daily exports):

```sql
SELECT incremental.create_time_interval_pipeline('event-export',
  time_interval := '1 day',
  batched := false,
  start_time := '2024-11-01',
  command := $$ SELECT export_events($1, $2) $$
);
```

### File List Pipeline

Process new files as they appear:

```sql
SELECT incremental.create_file_list_pipeline('event-import', 's3://mybucket/events/*.csv', $$
  SELECT import_events($1)
$$);
```

### Management Functions

| Function | Description |
|----------|-------------|
| `incremental.execute_pipeline(name)` | Manually execute a pipeline (only if new data exists) |
| `incremental.reset_pipeline(name)` | Reset pipeline to reprocess from the beginning |
| `incremental.drop_pipeline(name)` | Remove a pipeline |
| `incremental.skip_file(pipeline, path)` | Skip a faulty file in a file list pipeline |

### Monitoring

```sql
SELECT * FROM incremental.sequence_pipelines;
SELECT * FROM incremental.time_interval_pipelines;
SELECT * FROM incremental.processed_files;
```

Check job outcomes via pg_cron:

```sql
SELECT jobname, start_time, status, return_message
FROM cron.job_run_details JOIN cron.job USING (jobid)
WHERE jobname LIKE 'pipeline:%' ORDER BY 1 DESC LIMIT 5;
```
