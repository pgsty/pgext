

## Usage

> [pg_show_plans: show query plans of running SQL statements](https://github.com/cybertec-postgresql/pg_show_plans)

pg_show_plans displays the execution plans of all currently running SQL statements in real time. Plans are stored in a shared memory hash table.

### Viewing Running Plans

```sql
-- Show plans of all currently running queries
SELECT * FROM pg_show_plans;

-- Show plans along with their query text
SELECT * FROM pg_show_plans_q;
```

The views return:

| Column | Type | Description |
|--------|------|-------------|
| `pid` | integer | Server process ID |
| `level` | integer | Query nest level (0 = top level) |
| `userid` | oid | User OID |
| `dbid` | oid | Database OID |
| `plan` | text | Query plan text |
| `query` | text | Query string (only in `pg_show_plans_q`) |

### Nested Queries

When a function invokes SQL statements, they appear at deeper nesting levels:

```sql
SELECT * FROM pg_show_plans;
  pid  | level | userid | dbid  |                       plan
-------+-------+--------+-------+-----------------------------------------------
 11504 |     0 |     10 | 16384 | Function Scan on print_item  (cost=...)
 11504 |     1 |     10 | 16384 | Result  (cost=0.00..0.01 rows=1 width=4)
```

### GUC Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pg_show_plans.plan_format` | `text` | Output format: `text`, `json`, `yaml`, `xml` |
| `pg_show_plans.max_plan_length` | 16384 | Max plan length in bytes (affects shared memory) |
| `pg_show_plans.is_enabled` | `true` | Enable or disable the extension at runtime |
