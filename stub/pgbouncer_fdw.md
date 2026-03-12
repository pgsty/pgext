


## Usage

> [pgbouncer_fdw: Extension for querying PgBouncer stats from normal SQL views and running PgBouncer commands from normal SQL functions](https://github.com/CrunchyData/pgbouncer_fdw)

### Create Server

```sql
CREATE EXTENSION pgbouncer_fdw;

CREATE SERVER pgbouncer FOREIGN DATA WRAPPER dblink_fdw
  OPTIONS (host 'localhost', port '6432', dbname 'pgbouncer');
```

For multiple PgBouncer instances:

```sql
CREATE SERVER pgbouncer1 FOREIGN DATA WRAPPER dblink_fdw
  OPTIONS (host '192.168.1.10', port '6432', dbname 'pgbouncer');
CREATE SERVER pgbouncer2 FOREIGN DATA WRAPPER dblink_fdw
  OPTIONS (host '192.168.1.11', port '6432', dbname 'pgbouncer');

INSERT INTO pgbouncer_fdw_targets (target_host) VALUES ('pgbouncer1'), ('pgbouncer2');
UPDATE pgbouncer_fdw_targets SET active = false WHERE target_host = 'pgbouncer';
```

### Create User Mapping

```sql
CREATE USER MAPPING FOR PUBLIC SERVER pgbouncer
  OPTIONS (user 'ccp_monitoring', password 'mypassword');
```

### Available Views

| View | Description |
|------|-------------|
| `pgbouncer_clients` | Client connection details |
| `pgbouncer_pools` | Connection pool statistics |
| `pgbouncer_servers` | Backend server status |
| `pgbouncer_stats` | Statistics summary |
| `pgbouncer_databases` | Database definitions |
| `pgbouncer_config` | Configuration parameters |
| `pgbouncer_lists` | Internal lists |
| `pgbouncer_dns_hosts` | DNS host cache |
| `pgbouncer_dns_zones` | DNS zone cache |
| `pgbouncer_sockets` | Socket information |
| `pgbouncer_users` | User configuration |

```sql
SELECT * FROM pgbouncer_pools;
SELECT * FROM pgbouncer_stats;
SELECT database, cl_active, cl_waiting, sv_active FROM pgbouncer_pools;
```

When monitoring multiple instances, each row includes a `pgbouncer_target_host` column identifying the source.

### Command Functions

Administrative functions (require explicit `GRANT EXECUTE`):

```sql
SELECT pgbouncer_command_reload();              -- Reload configuration
SELECT pgbouncer_command_pause('mydb');          -- Pause a database
SELECT pgbouncer_command_resume('mydb');         -- Resume a database
SELECT pgbouncer_command_kill('mydb');           -- Kill connections
SELECT pgbouncer_command_disable('mydb');        -- Disable a database
SELECT pgbouncer_command_enable('mydb');         -- Enable a database
SELECT pgbouncer_command_reconnect('mydb');      -- Reconnect to backend
SELECT pgbouncer_command_set('key', 'value');    -- Set a parameter
SELECT pgbouncer_command_shutdown();             -- Shutdown PgBouncer
SELECT pgbouncer_command_suspend();              -- Suspend operations
SELECT pgbouncer_command_wait_close('mydb');     -- Wait for connections to close
```

### Permissions

```sql
GRANT USAGE ON FOREIGN SERVER pgbouncer TO monitoring_user;
GRANT SELECT ON pgbouncer_pools TO monitoring_user;
GRANT SELECT ON pgbouncer_stats TO monitoring_user;
GRANT EXECUTE ON FUNCTION pgbouncer_command_reload() TO admin_user;
```
