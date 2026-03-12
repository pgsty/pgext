


## Usage

> [pgauditlogtofile: Redirect pgAudit logs to an independent file](https://github.com/fmbiete/pgauditlogtofile)

`pgauditlogtofile` is an addon to pgAudit that redirects audit log lines to a separate file instead of the PostgreSQL server log, with automatic rotation support.

```sql
CREATE EXTENSION pgauditlogtofile;
```

### Configuration Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pgaudit.log_format` | `csv` | Output format: `csv` or `json` |
| `pgaudit.log_directory` | `log` | Directory for audit files (empty disables) |
| `pgaudit.log_filename` | `audit-%Y%m%d_%H%M.log` | Filename pattern (supports time patterns) |
| `pgaudit.log_file_mode` | `0600` | File permissions for audit logs |
| `pgaudit.log_rotation_age` | `1440` | Rotation interval in minutes (1 day) |
| `pgaudit.log_connections` | `off` | Log connection events (requires `log_connections = on`) |
| `pgaudit.log_disconnections` | `off` | Log disconnection events (requires `log_disconnections = on`) |
| `pgaudit.log_autoclose_minutes` | `0` | Auto-close file handler after N minutes of inactivity |
| `pgaudit.log_execution_time` | `off` | Measure statement execution time |
| `pgaudit.log_execution_memory` | `off` | Measure memory footprint of statements |

### Setup

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'pgaudit, pgauditlogtofile'
pgaudit.log_directory = 'log'
pgaudit.log_filename = 'audit-%Y%m%d_%H%M.log'
pgaudit.log_rotation_age = 1440
```

Audit entries are written to the separate file while server logs remain clean.
