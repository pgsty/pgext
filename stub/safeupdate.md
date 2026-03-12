


## Usage

> [safeupdate: Require criteria for UPDATE and DELETE](https://github.com/eradman/pg-safeupdate)

The `safeupdate` extension prevents accidental mass data changes by raising an error whenever `UPDATE` or `DELETE` statements are executed without a `WHERE` clause.

### Activation

```sql
-- Per-session
LOAD 'safeupdate';

-- Per-database (persistent)
ALTER DATABASE mydb SET session_preload_libraries = 'safeupdate';

-- Global (all databases, requires restart)
-- shared_preload_libraries = 'safeupdate'   -- in postgresql.conf
```

### Behavior

```sql
-- Blocked: UPDATE without WHERE
UPDATE rack SET fan_speed = 70;
-- ERROR: UPDATE requires a WHERE clause

-- Blocked: DELETE without WHERE
DELETE FROM rack;
-- ERROR: DELETE requires a WHERE clause

-- Allowed: with WHERE clause
UPDATE rack SET fan_speed = 90 WHERE fan_speed = 70;

-- Workaround: explicit always-true condition
UPDATE rack SET fan_speed = 90 WHERE 1 = 1;
```

### Administrative Override

```sql
-- Temporarily disable protection in current session
SET safeupdate.enabled = 0;
```

CTE-based modifications without WHERE conditions are also blocked. The extension is particularly useful with PostgREST or other systems that provide direct write access to the database.
