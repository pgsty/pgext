


## Usage

> [pg_ttl_index: Automatic data expiration with TTL indexes](https://github.com/ibrahimkarimeddin/postgres-extensions-pg_ttl)

`pg_ttl_index` provides automatic data expiration by associating a TTL (time-to-live) with a timestamp column. A background worker periodically deletes rows whose timestamp exceeds the configured expiration interval.

### Quick Start

```sql
-- Start the background worker
SELECT ttl_start_worker();

-- Create a table with a timestamp column
CREATE TABLE user_sessions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    session_data JSONB,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- Expire rows after 1 hour (3600 seconds)
SELECT ttl_create_index('user_sessions', 'created_at', 3600);
```

### Functions

| Function | Description |
|----------|-------------|
| `ttl_start_worker()` | Start the background worker for automatic cleanup |
| `ttl_worker_status()` | Check if the worker is running |
| `ttl_runner()` | Manually trigger cleanup |
| `ttl_create_index(table, column, expire_seconds [, batch_size])` | Configure TTL expiration |
| `ttl_drop_index(table, column)` | Remove TTL configuration |
| `ttl_summary()` | List all TTL indexes with stats |

### Examples

Session management with 24-hour expiry:

```sql
SELECT ttl_create_index('sessions', 'created_at', 86400, 5000);
```

Log retention for 7 days:

```sql
SELECT ttl_create_index('app_logs', 'logged_at', 604800);
```

Cache entries with custom expiry column (0 means the column itself holds the absolute expiry time):

```sql
SELECT ttl_create_index('cache_entries', 'expires_at', 0);
```

### Monitoring

```sql
SELECT * FROM ttl_summary();
```

Pause cleanup for a specific table:

```sql
UPDATE ttl_index_table SET active = false WHERE table_name = 'user_sessions';
```

### Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `pg_ttl_index.naptime` | Cleanup interval in seconds | 60 |
| `pg_ttl_index.enabled` | Enable/disable worker globally | on |

```sql
ALTER SYSTEM SET pg_ttl_index.naptime = 30;
SELECT pg_reload_conf();
```
