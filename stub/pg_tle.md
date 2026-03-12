


## Usage

> [pg_tle: Trusted Language Extensions for PostgreSQL](https://github.com/aws/pg_tle)

`pg_tle` lets you create and manage PostgreSQL extensions using trusted languages (SQL, PL/pgSQL, PL/v8, PL/Perl) without requiring filesystem access or server restarts.

Add `pg_tle` to `shared_preload_libraries` in `postgresql.conf`:

```
shared_preload_libraries = 'pg_tle'
```

### Install a TLE Extension

```sql
SELECT pgtle.install_extension(
  'my_extension',      -- extension name
  '1.0',               -- version
  'My custom extension', -- description
  $_pgtle_$
    CREATE FUNCTION my_func() RETURNS text AS $$
      SELECT 'hello from my_extension';
    $$ LANGUAGE SQL;
  $_pgtle_$
);
```

### Manage Extension Versions

```sql
-- Add an upgrade path
SELECT pgtle.install_update_path(
  'my_extension',  -- extension name
  '1.0',           -- from version
  '1.1',           -- to version
  $_pgtle_$
    CREATE OR REPLACE FUNCTION my_func() RETURNS text AS $$
      SELECT 'hello from my_extension v1.1';
    $$ LANGUAGE SQL;
  $_pgtle_$
);

-- Set default version
SELECT pgtle.set_default_version('my_extension', '1.1');
```

### Use the TLE Extension

```sql
CREATE EXTENSION my_extension;
SELECT my_func();  -- 'hello from my_extension'
ALTER EXTENSION my_extension UPDATE TO '1.1';
```

### Remove a TLE Extension

```sql
DROP EXTENSION my_extension;
SELECT pgtle.uninstall_extension('my_extension');
```

### Hooks and Features

Register custom hooks (e.g., password check hooks):

```sql
SELECT pgtle.register_feature('my_password_check', 'passcheck');
SELECT pgtle.unregister_feature('my_password_check', 'passcheck');
```

### Key Functions

| Function | Description |
|----------|-------------|
| `pgtle.install_extension()` | Install a new TLE extension |
| `pgtle.install_update_path()` | Add an upgrade path between versions |
| `pgtle.set_default_version()` | Set the default version for an extension |
| `pgtle.uninstall_extension()` | Remove a TLE extension |
| `pgtle.register_feature()` | Register a feature hook |
| `pgtle.unregister_feature()` | Unregister a feature hook |
| `pgtle.available_extensions()` | List available TLE extensions |
| `pgtle.available_extension_versions()` | List available versions |
