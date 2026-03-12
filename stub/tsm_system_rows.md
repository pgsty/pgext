

## Usage

> [tsm_system_rows: row-count-based TABLESAMPLE method](https://www.postgresql.org/docs/current/tsm-system-rows.html)

Provides the `SYSTEM_ROWS` table sampling method that returns exactly the specified number of rows.

```sql
CREATE EXTENSION tsm_system_rows;
```

### TABLESAMPLE Method

`SYSTEM_ROWS(count int)` -- maximum number of rows to return.

### Examples

```sql
-- Sample exactly 100 rows
SELECT * FROM my_table TABLESAMPLE SYSTEM_ROWS(100);

-- Quick peek at 10 rows from a large table
SELECT * FROM large_table TABLESAMPLE SYSTEM_ROWS(10);
```

Performs block-level sampling (may exhibit clustering effects with small samples). Returns all rows if the table has fewer rows than requested. Does not support `REPEATABLE`.
