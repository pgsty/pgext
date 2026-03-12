


## Usage

> [pg_drop_events: logs transaction ids of drop table, drop column, drop materialized view statements](https://github.com/bolajiwahab/pg_drop_events)

The `pg_drop_events` extension uses event triggers to automatically log details about DROP operations on tables, columns, and materialized views. The logged information can be used for point-in-time recovery (PITR) after accidental drops.

### Tracked Operations

- `DROP TABLE`
- `DROP COLUMN` (via `ALTER TABLE`)
- `DROP MATERIALIZED VIEW`

### Logged Information

| Column | Description |
|--------|-------------|
| `pid` | Process identifier |
| `usename` | Database user who executed the command |
| `query` | The SQL statement |
| `xact_id` | Transaction identifier |
| `wal_position` | Write-ahead log position |
| `objid` | Object identifier |
| `object_name` | Fully qualified name of dropped object |
| `object_type` | Object classification (table, table column, etc.) |
| `xact_time` | Timestamp of the transaction |

### Example

```sql
CREATE EXTENSION pg_drop_events;

-- Drop a table
DROP TABLE t.t3;
-- NOTICE: table t.t3 dropped by transaction 1085.

-- Query the event log
SELECT * FROM pg_drop_events;
```

### Point-in-Time Recovery

The logged data maps directly to PostgreSQL recovery parameters:

| pg_drop_events column | Recovery parameter |
|-----------------------|-------------------|
| `xact_id` | `recovery_target_xid` |
| `xact_time` | `recovery_target_time` |
| `wal_position` | `recovery_target_lsn` |

Use these values in `postgresql.conf` or `recovery.conf` to restore the database to a point just before the accidental drop.
