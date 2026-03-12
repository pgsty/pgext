


## Usage

> [pg_savior: Postgres extension to save OOPS mistakes](https://github.com/viggy28/pg_savior)

The `pg_savior` extension intercepts query execution to prevent accidental data deletion. It hooks into the executor to detect dangerous DELETE operations and block them.

### How It Works

When a DELETE statement is processed, `pg_savior` checks for:

- **Missing WHERE clauses** on DELETE commands
- **Index scan operations** in DELETE query plans
- **Complex queries** involving CTEs and subqueries in DELETE operations

When a risky DELETE is detected, the extension prevents execution and returns an informational message with zero rows affected.

### Example

```sql
CREATE EXTENSION pg_savior;

-- Attempting a DELETE without WHERE clause
DELETE FROM emp;
-- INFO:  pg_savior: DELETE statement detected
-- INFO:  pg_savior: WHERE clause is mandatory for a DELETE statement
-- DELETE 0  (no rows affected, data preserved)

-- Normal DELETE with WHERE clause works as expected
DELETE FROM emp WHERE id = 42;
-- DELETE 1
```

### Notes

- The extension operates through PostgreSQL executor hooks, requiring no changes to application code
- Only DELETE statements are currently intercepted; UPDATE operations are not affected
- Planned features include preventing `CREATE INDEX` without `CONCURRENTLY` and blocking `DROP DATABASE`
