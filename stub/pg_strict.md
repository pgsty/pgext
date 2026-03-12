


## Usage

> [pg_strict: Prevent dangerous UPDATE and DELETE without WHERE clause](https://github.com/spa5k/pg_strict)

The `pg_strict` extension blocks `UPDATE` and `DELETE` statements that lack a `WHERE` clause. It operates at the parse/analyze stage via `post_parse_analyze_hook`, providing three enforcement modes per statement type.

### Configuration Parameters

| Parameter | Modes | Description |
|-----------|-------|-------------|
| `pg_strict.require_where_on_update` | `on`/`warn`/`off` | Enforce WHERE on UPDATE |
| `pg_strict.require_where_on_delete` | `on`/`warn`/`off` | Enforce WHERE on DELETE |

- **`on`**: Reject statements without WHERE (raises error)
- **`warn`**: Allow but emit a warning log
- **`off`**: Standard PostgreSQL behavior

### Session-Level Configuration

```sql
SET pg_strict.require_where_on_update = 'on';
SET pg_strict.require_where_on_delete = 'warn';
```

### Persistent Configuration

```sql
ALTER DATABASE postgres SET pg_strict.require_where_on_update = 'on';
ALTER ROLE app_service SET pg_strict.require_where_on_delete = 'on';
ALTER ROLE dba_admin SET pg_strict.require_where_on_update = 'off';
```

### Transactional Override

```sql
BEGIN;
SET LOCAL pg_strict.require_where_on_delete = 'off';
DELETE FROM temp_table;  -- allowed within this transaction
COMMIT;
```

### API Functions

```sql
SELECT pg_strict_version();           -- extension version
SELECT pg_strict_config();            -- all settings with values and descriptions

-- Validate queries programmatically
SELECT pg_strict_check_where_clause('DELETE FROM t', 'DELETE');  -- returns boolean
SELECT pg_strict_validate_update('UPDATE t SET x=1');
SELECT pg_strict_validate_delete('DELETE FROM t');

-- Quick mode toggles
SELECT pg_strict_enable_update();     -- set update enforcement to 'on'
SELECT pg_strict_warn_delete();       -- set delete enforcement to 'warn'
SELECT pg_strict_disable_update();    -- set update enforcement to 'off'
```

Any non-null WHERE condition is accepted (including `WHERE false`). CTE statements are supported.
