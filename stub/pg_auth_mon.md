


## Usage

> [pg_auth_mon: Monitor authentication attempts](https://github.com/RafiaSabih/pg_auth_mon)

`pg_auth_mon` stores authentication attempt statistics per user, tracking successful and failed logins with timestamps.

```sql
CREATE EXTENSION pg_auth_mon;
```

### Configuration

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'pg_auth_mon'
```

Optional GUC:

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pg_auth_mon.log_period` | `0` | Dump stats to Postgres log every N minutes (0=off) |

### Querying Authentication Stats

```sql
SELECT * FROM pg_auth_mon;
```

The `pg_auth_mon` view provides per-user information:

| Column | Description |
|--------|-------------|
| `uid` | User OID (0 for invalid/non-existing usernames) |
| `successful_auth_count` | Total number of successful logins |
| `last_successful_auth` | Timestamp of last successful login |
| `failed_auth_count` | Total number of failed authentications |
| `last_failed_auth` | Timestamp of last failed login |
| `hba_conflicts_count` | Count of HBA file conflicts |

All login attempts with invalid usernames are aggregated into a single row with OID 0.
