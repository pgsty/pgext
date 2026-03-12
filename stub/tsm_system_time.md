

## Usage

> [tsm_system_time: time-based TABLESAMPLE method](https://www.postgresql.org/docs/current/tsm-system-time.html)

Provides the `SYSTEM_TIME` table sampling method that returns as many rows as can be read within a specified time limit.

```sql
CREATE EXTENSION tsm_system_time;
```

### TABLESAMPLE Method

`SYSTEM_TIME(milliseconds float)` -- maximum time to spend reading the table.

### Examples

```sql
-- Sample rows readable within 1 second
SELECT * FROM my_table TABLESAMPLE SYSTEM_TIME(1000);

-- Sample from a large table with a 500ms budget
SELECT count(*) FROM large_table TABLESAMPLE SYSTEM_TIME(500);
```

Performs block-level sampling (not row-level). If the entire table can be read within the time limit, all rows are returned. Does not support `REPEATABLE`.
