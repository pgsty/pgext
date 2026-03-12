

## Usage

> [pg_track_settings: track PostgreSQL configuration changes](https://github.com/rjuju/pg_track_settings)

pg_track_settings records changes to PostgreSQL settings over time, tracking both global settings (`pg_settings`) and per-role/database overrides (`pg_db_role_setting`).

### Taking Snapshots

Call periodically (via cron or PoWA) to capture current settings:

```sql
SELECT pg_track_settings_snapshot();
```

### Viewing Settings at a Point in Time

```sql
-- All settings at a specific time
SELECT * FROM pg_track_settings('2024-01-15 10:00:00');

-- All overloaded (per-role/database) settings at a specific time
SELECT * FROM pg_track_db_role_settings('2024-01-15 10:00:00');
```

### Comparing Settings Between Two Times

```sql
-- Find what changed in the last hour
SELECT * FROM pg_track_settings_diff(now() - interval '1 hour', now());

-- Compare overloaded settings
SELECT * FROM pg_track_db_role_settings_diff(now() - interval '1 hour', now());
```

### Viewing Change History

```sql
-- History of a specific setting
SELECT * FROM pg_track_settings_log('work_mem');

-- History of an overloaded setting
SELECT * FROM pg_track_db_role_settings_log('work_mem');

-- PostgreSQL restart history
SELECT * FROM pg_reboot;
```

### Resetting History

```sql
SELECT pg_track_settings_reset();
```

### Functions Summary

| Function | Description |
|----------|-------------|
| `pg_track_settings_snapshot()` | Capture current settings |
| `pg_track_settings(timestamptz)` | All settings at a given time |
| `pg_track_settings_diff(timestamptz, timestamptz)` | Settings that changed between two times |
| `pg_track_settings_log(text)` | History of a specific setting |
| `pg_track_db_role_settings(timestamptz)` | Overloaded settings at a given time |
| `pg_track_db_role_settings_diff(timestamptz, timestamptz)` | Overloaded settings changes |
| `pg_track_db_role_settings_log(text)` | History of a specific overloaded setting |
| `pg_track_settings_reset()` | Clear all history |
