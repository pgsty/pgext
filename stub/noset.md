


## Usage

> [noset: Prevent users from changing session parameters via SET/RESET](https://gitlab.com/ongresinc/extensions/noset)

`noset` is a loadable module that prevents specific users from using `SET` or `RESET` commands to change session parameters.

```sql
CREATE EXTENSION noset;
```

### Configuration

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'noset'
```

### GUC Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `noset.enabled` | `false` | Enable SET/RESET blocking for the role |
| `noset.parameters` | `*` | Parameters to block (comma-separated, `*` = all) |

### Setting Up Per-User Restrictions

```sql
-- Block ALL SET/RESET for a user
ALTER USER appuser SET noset.enabled = true;

-- Block only specific parameters
ALTER USER appuser SET noset.enabled = true;
ALTER USER appuser SET noset.parameters = 'work_mem,jit';
```

### Example

```sql
-- As appuser:
SET work_mem = '1GB';
-- ERROR: permission denied to set/reset parameter 'set work_mem = '1GB';'

SET maintenance_work_mem = '1GB';
-- SET (allowed, not in blocked list)
```

### Finding Restricted Users

```sql
SELECT usename, useconfig FROM pg_user
WHERE useconfig IS NOT NULL
  AND array['noset.enabled=on'] <@ useconfig;
```

### Notes

- Does not apply to superusers
- The extension revokes access to the `set_config` function from PUBLIC
