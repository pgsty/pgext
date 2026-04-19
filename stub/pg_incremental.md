## Usage

- Sources: [README](https://github.com/CrunchyData/pg_incremental/blob/main/README.md), [v1.5.0 release](https://github.com/CrunchyData/pg_incremental/releases/tag/v1.5.0)

`pg_incremental` defines exactly-once incremental pipelines for append-only tables and file feeds. Upstream documents three pipeline types: sequence, time-interval, and file-list.

### Install And Scheduling Model

The upstream README still documents `pg_cron`-backed scheduling and installs with:

```sql
CREATE EXTENSION pg_incremental CASCADE;
```

Pipelines run immediately when created unless `execute_immediately := false`, then continue on a `pg_cron` schedule. The README notes that each scheduled execution appears in `cron.job_run_details` even when no new data is available.

### Sequence Pipelines

Use sequence pipelines to process safe ranges of sequence values:

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

The README documents `max_batch_size` for limiting how many sequence IDs are processed per run.

### Time-Interval Pipelines

Use time windows when the command should receive `$1` and `$2` as a passed interval:

```sql
SELECT incremental.create_time_interval_pipeline('event-aggregation', '1 day', $$
  INSERT INTO events_agg
  SELECT event_time::date, count(DISTINCT event_id)
  FROM events
  WHERE event_time >= $1 AND event_time < $2
  GROUP BY 1
$$);
```

For export-style jobs, the README documents `batched := false` so each interval runs separately.

### File-List Pipelines

Use file-list pipelines to process newly discovered files:

```sql
SELECT incremental.create_file_list_pipeline('event-import', 's3://mybucket/events/*.csv', $$
  SELECT import_events($1)
$$);
```

The v1.5.0 release adds `max_batches_per_run` to file-list pipelines. The README documents `incremental.skip_file()` for permanently marking a bad file as processed.

### Operations And Monitoring

The README documents:

- `CALL incremental.execute_pipeline(name)`: run once if new work exists.
- `SELECT incremental.reset_pipeline(name)`: reset progress.
- `SELECT incremental.drop_pipeline(name)`: remove a pipeline.
- Views and tables such as `incremental.sequence_pipelines`, `incremental.time_interval_pipelines`, `incremental.file_list_pipelines`, and `incremental.processed_files`.

The v1.5.0 release note also calls out a `DROP EXTENSION` fix for environments where `pg_cron` is not present.
