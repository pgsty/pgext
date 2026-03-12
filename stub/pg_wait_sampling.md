

## Usage

> [pg_wait_sampling: sampling based statistics of wait events](https://github.com/postgrespro/pg_wait_sampling)

pg_wait_sampling collects wait event statistics by periodically sampling what each backend is waiting on. It provides both a recent history ring buffer and a cumulative profile.

### Views

**Current wait events:**

```sql
SELECT * FROM pg_wait_sampling_current;
```

| Column | Type | Description |
|--------|------|-------------|
| `pid` | int4 | Process ID |
| `event_type` | text | Wait event type |
| `event` | text | Wait event name |
| `queryid` | int8 | Query ID |

**Wait event history** (ring buffer of recent samples):

```sql
SELECT * FROM pg_wait_sampling_history;
```

| Column | Type | Description |
|--------|------|-------------|
| `pid` | int4 | Process ID |
| `ts` | timestamptz | Sample timestamp |
| `event_type` | text | Wait event type |
| `event` | text | Wait event name |
| `queryid` | int8 | Query ID |

**Wait event profile** (cumulative counts):

```sql
SELECT * FROM pg_wait_sampling_profile;
```

| Column | Type | Description |
|--------|------|-------------|
| `pid` | int4 | Process ID |
| `event_type` | text | Wait event type |
| `event` | text | Wait event name |
| `queryid` | int8 | Query ID |
| `count` | text | Count of samples |

**Get current wait event for a specific process:**

```sql
SELECT * FROM pg_wait_sampling_get_current(12345);
```

### Reset Profile

```sql
SELECT pg_wait_sampling_reset_profile();
```

### Configuration

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pg_wait_sampling.history_size` | 5000 | Ring buffer size |
| `pg_wait_sampling.history_period` | 10 | History sampling period (ms) |
| `pg_wait_sampling.profile_period` | 10 | Profile sampling period (ms) |
| `pg_wait_sampling.profile_pid` | `true` | Collect profile per-process |
| `pg_wait_sampling.profile_queries` | `top` | Per-query profile: `none`, `top`, `all` |
| `pg_wait_sampling.sample_cpu` | `true` | Sample on-CPU backends (event columns NULL) |

### Example: Top Wait Events

```sql
SELECT event_type, event, count
FROM pg_wait_sampling_profile
ORDER BY count DESC
LIMIT 10;
```
