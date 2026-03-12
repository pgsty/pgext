


## Usage

> [pg_upless: Detect Useless UPDATE](https://github.com/rodo/pg_upless)

pg_upless detects UPDATE statements that change no actual values (common with ORMs). It works by installing triggers on monitored tables and collecting statistics. It is intended as a diagnostic tool, not for permanent use.

### Start Monitoring

```sql
-- Monitor a specific table
SELECT pg_upless_start('public', 'boats');

-- Monitor all tables in a schema
SELECT pg_upless_start('public');
```

### Stop Monitoring

```sql
-- Stop monitoring a specific table
SELECT pg_upless_stop('public', 'boats');

-- Stop monitoring all tables in a schema
SELECT pg_upless_stop('public');
```

### View Statistics

The extension creates two tables in the `pg_upless` schema:

- **`pg_upless_stats`** -- stores counts of total vs useless updates per table
- **`pg_upless_start_time`** -- records when monitoring started (for rate calculations)
