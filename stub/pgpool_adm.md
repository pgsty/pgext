


## Usage

> [pgpool_adm: Administrative functions for pgPool](https://pgpool.net/)

The `pgpool_adm` extension provides SQL-callable wrapper functions for Pgpool-II PCP (Pgpool Control Protocol) commands, enabling administration of Pgpool-II from within PostgreSQL.

### Available Functions

| Function | Description |
|----------|-------------|
| `pgpool_adm_pcp_node_info` | Display information on a given backend node |
| `pgpool_adm_pcp_health_check_stats` | Display health check statistics for a node |
| `pgpool_adm_pcp_pool_status` | Retrieve parameters from pgpool.conf |
| `pgpool_adm_pcp_node_count` | Get the number of backend nodes |
| `pgpool_adm_pcp_attach_node` | Attach a backend node |
| `pgpool_adm_pcp_detach_node` | Detach a backend node |
| `pgpool_adm_pcp_proc_info` | Display Pgpool-II child process information |

### Call Methods

Functions support two calling conventions:

**Direct parameters** (hostname, port, username, password, plus function-specific args):

```sql
SELECT * FROM pgpool_adm_pcp_node_info('localhost', 9898, 'admin', 'password', 0);
SELECT * FROM pgpool_adm_pcp_node_count('localhost', 9898, 'admin', 'password');
SELECT * FROM pgpool_adm_pcp_pool_status('localhost', 9898, 'admin', 'password');
```

**Foreign server reference** (uses port 9898 and credentials from `~/.pcppass`):

```sql
SELECT * FROM pgpool_adm_pcp_node_info(server_name := 'pgpool_server', node_id := 0);
SELECT * FROM pgpool_adm_pcp_node_count(server_name := 'pgpool_server');
```

### Node Management

```sql
-- Detach a backend node
SELECT pgpool_adm_pcp_detach_node('localhost', 9898, 'admin', 'password', 1);

-- Re-attach a backend node
SELECT pgpool_adm_pcp_attach_node('localhost', 9898, 'admin', 'password', 1);
```

The default PCP communication port is 9898. Credentials can be managed via the `.pcppass` file in the user's home directory.
