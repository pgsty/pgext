


## Usage

> [table_log: record table modification logs and PITR for table/row](https://github.com/df7cb/table_log)

The `table_log` extension records INSERT, UPDATE, and DELETE operations on a table into a separate log table, enabling point-in-time recovery at the table or row level.

### Initialize Logging

```sql
CREATE EXTENSION table_log;

-- Basic setup: creates a log table and trigger for 'my_table'
-- Level 5 = log trigger_id + trigger_user + trigger columns
SELECT table_log_init(5, 'my_table');

-- With explicit log schema
SELECT table_log_init(5, 'my_table', 'log_schema');

-- Full form with all options
SELECT table_log_init(
    5,                -- level: 3=minimal, 4=+user, 5=+id+user
    'public',         -- source schema
    'my_table',       -- source table
    'log_schema',     -- log table schema
    'my_table_log',   -- log table name (default: {table}_log)
    'SINGLE',         -- partition mode: 'SINGLE' or 'PARTITION'
    false,            -- basic_mode (simpler trigger)
    '{INSERT, UPDATE, DELETE}'::text[]  -- actions to log
);
```

### Log Table Structure

The log table mirrors the original table columns plus metadata:

| Column | Description |
|--------|-------------|
| `trigger_mode` | Operation type: INSERT, UPDATE, DELETE |
| `trigger_tuple` | Tuple version: 'old' or 'new' |
| `trigger_changed` | Timestamp of the change |
| `trigger_id` | Sequential ID (level 4+) |
| `trigger_user` | User who made the change (level 5) |

### Point-in-Time Restore

```sql
-- Restore table to a specific point in time
SELECT table_log_restore_table(
    'my_table',           -- original table name
    'my_table_log',       -- log table name
    'id',                 -- primary key column
    'trigger_changed',    -- timestamp column in log
    'trigger_tuple',      -- tuple type column in log
    '2024-01-15 10:30:00' -- restore to this timestamp
);
```

### Trigger Functions

| Function | Description |
|----------|-------------|
| `table_log()` | Full trigger function logging all columns |
| `table_log_basic()` | Basic trigger function with simpler logging |
| `table_log_restore_table(...)` | Restore table state to a given timestamp |
