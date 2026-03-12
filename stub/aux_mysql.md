

## Usage

> [aux_mysql: MySQL Supplementary Extension](https://github.com/HaloTech-Co-Ltd/openHalo)

The `aux_mysql` extension is part of the openHalo project, providing MySQL compatibility functions and features for PostgreSQL. It enables PostgreSQL to understand MySQL SQL dialect and communication protocol.

### Enabling

```sql
CREATE EXTENSION aux_mysql CASCADE;
```

### Overview

When used with openHalo's MySQL compatibility mode, this extension allows:

- MySQL client connections via the MySQL wire protocol (port 3306)
- MySQL-compatible SQL syntax parsing
- MySQL-compatible functions and operators

### MySQL Compatibility Mode

Configure in `postgresql.conf`:

```ini
database_compat_mode = 'mysql'      -- enable MySQL mode
mysql.listener_on = true            -- enable MySQL protocol listener
mysql.port = 3306                   -- MySQL protocol port
```

After enabling, MySQL clients can connect directly:

```bash
mysql -P 3306 -h 127.0.0.1
```

### Key Features

- MySQL-compatible SQL dialect support
- MySQL wire protocol compatibility (TDS)
- MySQL-style authentication (`mysql_native_password`)
- Common MySQL functions and operators available in PostgreSQL

### Notes

- This extension is designed to work as part of the openHalo distribution
- Standard PostgreSQL connections continue to work alongside MySQL protocol connections
- Not all MySQL features are supported; focuses on commonly used functionality
