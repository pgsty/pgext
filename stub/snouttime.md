## Usage

Sources:

- [README.md](https://github.com/snoutdata/snouttime/blob/f92202af9f8e5d6caeae52fa312209cf8dc42059/README.md)
- [snouttime.control](https://github.com/snoutdata/snouttime/blob/f92202af9f8e5d6caeae52fa312209cf8dc42059/snouttime.control)
- [Cargo.toml](https://github.com/snoutdata/snouttime/blob/f92202af9f8e5d6caeae52fa312209cf8dc42059/Cargo.toml)
- [src/series.sql](https://github.com/snoutdata/snouttime/blob/f92202af9f8e5d6caeae52fa312209cf8dc42059/src/series.sql)
- [src/worker.rs](https://github.com/snoutdata/snouttime/blob/f92202af9f8e5d6caeae52fa312209cf8dc42059/src/worker.rs)
- [src/jobs.sql](https://github.com/snoutdata/snouttime/blob/f92202af9f8e5d6caeae52fa312209cf8dc42059/src/jobs.sql)
- [src/columnar/mod.rs](https://github.com/snoutdata/snouttime/blob/f92202af9f8e5d6caeae52fa312209cf8dc42059/src/columnar/mod.rs)
- [sql/snouttime--0.1.5--0.1.6.sql](https://github.com/snoutdata/snouttime/blob/f92202af9f8e5d6caeae52fa312209cf8dc42059/sql/snouttime--0.1.5--0.1.6.sql)

`snouttime` 0.1.6 provides time-series partitions, columnar storage, rollups and temporal joins on PostgreSQL 17 and 18. It is an early source release; evaluate its storage and upgrade behavior before placing existing data under its management.

### Core Workflow

Install the SQL extension as a superuser. Series operations require the relevant table ownership privileges. This example starts with a new table; conversion of an existing table is a structural operation, not merely a metadata label. Basic functions and manually run jobs do not require preload.

```sql
CREATE EXTENSION snouttime;
CREATE TABLE metrics (ts timestamptz NOT NULL, value double precision);
SELECT snouttime.create_series('metrics', 'ts',
  partition_interval => interval '1 day');
INSERT INTO metrics VALUES (now(), 42);
SELECT snouttime.bucket('1 hour', ts) AS bucket, avg(value)
FROM metrics GROUP BY 1 ORDER BY 1;
SELECT snouttime.run_due_job();
```

### Partitions and Query Objects

`snouttime.create_series` creates a native partitioned parent and retains existing rows in its default partition. Use `partition_interval` for time columns or `partition_width` for integer time. Existing constraints and dependencies can restrict conversion; inspect the upstream conversion requirements before applying it to an application table.

`snouttime.bucket` groups time values; `snouttime.gapfill`, `snouttime.locf` and `snouttime.interpolate` generate or fill missing buckets. `snouttime.first` and `snouttime.last` select values by time, with unspecified tie order. `snouttime.counter_delta` and `snouttime.counter_rate` require input ordered by time. Percentile and distinct-count sketches are approximate. `snouttime.asof_join` selects the latest matching right-side observation; `snouttime.window_join` aggregates a bounded time window. These queries run with caller privileges.

### Jobs, Retention and Rollups

For jobs to start at server boot, merge `snouttime` into `shared_preload_libraries`, set `snouttime.databases` to the comma-separated database list and restart. `snouttime.interval` defaults to 10 seconds. A superuser can instead call `snouttime.start_worker()` for the current database without preload; that dynamic worker does not survive a server restart. Manual runners call `snouttime.run_due_job()` until it returns false. Jobs act under their table owners' privileges.

`snouttime.create_rollup` creates a view that combines materialized buckets with raw changes; `snouttime.refresh_rollup` refreshes manually. Writes made directly to a child partition bypass its change tracking. Retention removes old partitions; review retention and backup policies before enabling it.

### Columnar Storage and Object Storage

`snouttime.seal` rewrites a partition to the `snouttime_columnar` access method; `snouttime.unseal` converts it back to heap. Late writes use an internal delta store. Non-unique indexes normally cover only those later writes; use the documented `keep_indexes` option when full coverage is needed. Unique indexes and exclusion constraints still cover all rows. BRIN indexes, `CREATE INDEX CONCURRENTLY` and `TABLESAMPLE` are rejected on sealed partitions. Concurrent updates may require a client retry.

Tiering moves sealed storage to S3-compatible object storage. Credential GUCs require preload and superuser management; server environment credentials are another documented path. `snouttime.tier_gc` defaults to off because restored databases can still reference old objects. Do not enable garbage collection without considering those copies.

### Schemas and Upgrades

The extension is not trusted or relocatable. Public objects live in `snouttime`, while columnar side tables use the protected `snouttime_internal` schema. Version 0.1.6 repairs reading sealed partitions as non-superusers. Upgrade through `ALTER EXTENSION snouttime UPDATE` after installing the matching library and scripts. No downgrade scripts are provided: reverting requires a backup from before the upgrade.
