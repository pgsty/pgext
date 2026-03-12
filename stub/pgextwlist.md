


## Usage

> [pgextwlist: PostgreSQL extension whitelisting](https://github.com/dimitri/pgextwlist)

`pgextwlist` implements extension whitelisting: only explicitly allowed extensions can be installed, and whitelisted extensions are installed with superuser privileges even when requested by non-superusers.

### Configuration

Add to `postgresql.conf`:

```ini
local_preload_libraries = 'pgextwlist'
extwlist.extensions = 'hstore,cube,pg_stat_statements'
```

Or per-role:

```sql
ALTER ROLE adminuser SET extwlist.extensions = 'pg_stat_statements, postgis';
```

| Parameter | Description |
|-----------|-------------|
| `extwlist.extensions` | Comma-separated list of whitelisted extensions |
| `extwlist.custom_path` | Filesystem path for custom pre/post scripts |

### Behavior

Non-superusers can install whitelisted extensions:

```sql
-- Allowed (hstore is whitelisted)
CREATE EXTENSION hstore;

-- Blocked (not whitelisted)
CREATE EXTENSION earthdistance;
-- ERROR: extension "earthdistance" is not whitelisted
```

Operations `CREATE EXTENSION`, `DROP EXTENSION`, `ALTER EXTENSION ... UPDATE`, and `COMMENT ON EXTENSION` are run as superuser for whitelisted extensions.

### Custom Scripts

Place scripts in `${extwlist.custom_path}/extname/`:

| Script | When |
|--------|------|
| `before--1.0.sql` | Before installing version 1.0 |
| `before-create.sql` | Before CREATE (fallback) |
| `after--1.0.sql` | After installing version 1.0 |
| `after-create.sql` | After CREATE (fallback) |
| `before-update.sql` / `after-update.sql` | Around ALTER EXTENSION UPDATE |
| `before-drop.sql` / `after-drop.sql` | Around DROP EXTENSION |

Custom scripts support template variables: `@extschema@`, `@current_user@`, `@database_owner@`.
