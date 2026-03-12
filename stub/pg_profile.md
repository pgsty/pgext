

## Usage

> [pg_profile: historical performance profiling tool for PostgreSQL](https://github.com/zubkov-andrei/pg_profile)

pg_profile collects periodic samples of PostgreSQL statistics and generates detailed historical performance reports. It depends on `pg_stat_statements` and optionally uses `pg_stat_kcache` and `pg_wait_sampling` for additional metrics.

### Taking Samples

Samples must be taken periodically (e.g., via cron). Each sample captures the current state of statistics:

```sql
SELECT profile.take_sample();
```

### Generating Reports

Build a report between two sample IDs to analyze performance during that interval:

```sql
-- Regular report between samples 1 and 2
SELECT profile.get_report(1, 2);

-- Differential report comparing two intervals
SELECT profile.get_diffreport(1, 2, 3, 4);
```

### Managing Servers

pg_profile can collect statistics from remote clusters:

```sql
-- Define a remote server
SELECT profile.create_server('remote', 'host=remote_host dbname=postgres');

-- List defined servers
SELECT * FROM profile.show_servers();

-- Enable/disable a server
SELECT profile.enable_server('remote');
SELECT profile.disable_server('remote');
```

### Baselines

Baselines protect sample ranges from automatic cleanup:

```sql
-- Create a baseline preserving samples 10 through 20
SELECT profile.create_baseline('incident_2024', 10, 20);

-- List baselines
SELECT * FROM profile.show_baselines();

-- Drop a baseline
SELECT profile.drop_baseline('incident_2024');
```

### Retention

Control how long samples are kept:

```sql
-- Set retention to 7 days for the local server
SELECT profile.set_server_max_sample_age('local', 7);
```

### Sample Information

```sql
-- Show available samples
SELECT * FROM profile.show_samples();

-- Show time spent taking samples (requires pg_profile.track_sample_timings = on)
SELECT * FROM v_sample_timings;
```

### Recommended Settings

```
track_activities = on
track_counts = on
track_io_timing = on
track_wal_io_timing = on      # PG 14+
track_functions = all
```
