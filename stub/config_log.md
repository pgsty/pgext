## Usage

Sources:

- [Official README](https://github.com/ibarwick/config_log/blob/db66842eec9398701bd045fad7e6a09f6b77eeb5/README.md)
- [Official extension SQL](https://github.com/ibarwick/config_log/blob/db66842eec9398701bd045fad7e6a09f6b77eeb5/config_log--0.1.7.sql)
- [Official C implementation](https://github.com/ibarwick/config_log/blob/db66842eec9398701bd045fad7e6a09f6b77eeb5/config_log.c)
- [Official extension control file](https://github.com/ibarwick/config_log/blob/db66842eec9398701bd045fad7e6a09f6b77eeb5/config_log.control)

`config_log` is an experimental background worker that records PostgreSQL configuration-file settings at server start and after a configuration reload. Version `0.1.7` tracks changes in one database per cluster; it does not continuously watch arbitrary files.

### Core Workflow

Choose the database and schema before starting the worker. The defaults are `postgres` and `public`:

```conf
config_log.database = 'postgres'
config_log.schema = 'public'
shared_preload_libraries = 'config_log'
```

Restart PostgreSQL, connect to that database, and install the extension in the configured schema:

```sql
CREATE EXTENSION config_log;
SELECT * FROM pg_settings_log_current ORDER BY name;
```

The schema must be reachable through the superuser's search path. A later `pg_reload_conf()` sends the worker a reload signal; the worker invokes `pg_settings_logger()` and appends configuration differences.

### Objects and Caveats

`pg_settings_log` stores INSERT, UPDATE, and DELETE history with timestamps. `pg_settings_log_current` presents the latest state, while `pg_settings_logger()` performs a comparison. The initial snapshot is limited to rows in `pg_settings` whose source is `configuration file`.

The extension revokes public access from its table, view, and logger function, but recorded rows include `sourcefile` and `sourceline`. Grant access deliberately because these paths and settings may disclose deployment details. Only the configured database is monitored, and changing `config_log.database` or `config_log.schema` requires coordinating the extension objects, preload configuration, and a server restart.
