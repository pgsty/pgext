

## Usage

> [pg_track_optimizer: detect suboptimal query plans via cardinality estimation errors](https://github.com/danolivo/pg_track_optimizer)

pg_track_optimizer automatically detects queries with poor cardinality estimates by comparing planner predictions to actual execution results. It calculates multiple error metrics using logarithmic scale.

### Enable Tracking

```sql
-- Track only problematic queries in production
SET pg_track_optimizer.mode = 'normal';

-- Track all queries during debugging
SET pg_track_optimizer.mode = 'forced';

-- Log EXPLAIN for queries exceeding error threshold
SET pg_track_optimizer.log_min_error = 2.0;
```

### Viewing Tracked Queries

```sql
SELECT queryid, query,
       avg_avg, avg_min, avg_max,
       rms_avg, rms_min, rms_max,
       time_avg, blks_avg, nexecs
FROM pg_track_optimizer
ORDER BY avg_avg DESC
LIMIT 10;

-- Using the RStats type directly
SELECT queryid, query,
       wca_error -> 'mean' AS avg_wca_error,
       blks_accessed -> 'mean' AS avg_blocks
FROM pg_track_optimizer()
WHERE blks_accessed -> 'mean' > 1000
ORDER BY wca_error -> 'mean' DESC;
```

### Error Metrics

| Metric | Description |
|--------|-------------|
| `avg_error` | Simple average of log-scale errors across plan nodes |
| `rms_error` | Root Mean Square, emphasizes large errors |
| `twa_error` | Time-Weighted Average, highlights slow nodes |
| `wca_error` | Cost-Weighted Average, highlights high-cost nodes |
| `f_join_filter` | JOIN filtering overhead factor |
| `f_scan_filter` | Scan filtering overhead factor |

### Managing Statistics

```sql
-- Save statistics to disk
SELECT pg_track_optimizer_flush();

-- Clear all tracked statistics
SELECT pg_track_optimizer_reset();

-- Check extension status
SELECT * FROM pg_track_optimizer_status;
```

### Configuration

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pg_track_optimizer.mode` | `disabled` | `disabled`, `normal`, `forced` |
| `pg_track_optimizer.log_min_error` | (none) | Error threshold for logging EXPLAIN |
| `pg_track_optimizer.hash_mem` | (default) | Shared memory limit in KB |
| `pg_track_optimizer.auto_flush` | `on` | Auto-save stats on backend shutdown |
