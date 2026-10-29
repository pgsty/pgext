## Usage

Sources:

- [Official README](https://github.com/CrunchyData/pg_incremental/blob/v1.6.1/README.md)
- [v1.6.1 release](https://github.com/CrunchyData/pg_incremental/releases/tag/v1.6.1)
- [Control file](https://github.com/CrunchyData/pg_incremental/blob/v1.6.1/pg_incremental.control)

`pg_incremental` defines exactly-once incremental pipelines for append-only tables and file feeds. It depends on `pg_cron` and supports sequence, time-interval and file-list pipelines on PostgreSQL 17–18 in the reviewed upstream test matrix.

### Install and schedule

```sql
CREATE EXTENSION pg_incremental CASCADE;
GRANT USAGE ON SCHEMA cron TO app_user;
```

Pipelines execute immediately by default and then run from `pg_cron`; set `execute_immediately := false` when the initial run should be deferred. The progress update and the pipeline command commit together, so a failed command does not advance the processed range.

### Pipeline types

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

Sequence pipelines pass a safe `$1`/`$2` range; time pipelines pass interval boundaries; file-list pipelines pass a newly discovered file. `max_batch_size`, `max_batches_per_run` and `batched := false` tune workload shape. `incremental.execute_pipeline`, `reset_pipeline`, `drop_pipeline` and the pipeline status views provide operational control.

### Security fixes in 1.6.1

Release 1.6.1 fixes a `_drop_trigger` `search_path` injection issue, requires `SELECT` or `USAGE` on a sequence when a sequence pipeline is created and run, and executes file-list functions as the caller. Review ownership and grants before upgrading; existing package facts are separate from this upstream release.
