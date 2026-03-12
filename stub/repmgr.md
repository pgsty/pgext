

## Usage

> [repmgr: Replication manager for PostgreSQL](https://github.com/EnterpriseDB/repmgr)

A suite of tools to manage replication and failover within a PostgreSQL cluster. Supports setting up standby servers, monitoring replication, and performing failover/switchover.

### Enabling

```sql
CREATE EXTENSION repmgr;
```

### Configuration

Create `repmgr.conf` on each node:

```ini
node_id=1
node_name='node1'
conninfo='host=node1 dbname=repmgr user=repmgr'
data_directory='/var/lib/postgresql/data'
```

### Register Primary

```bash
repmgr -f /etc/repmgr.conf primary register
```

### Clone and Register Standby

```bash
# Clone from primary
repmgr -h primary-host -U repmgr -d repmgr -f /etc/repmgr.conf standby clone

# Start the standby
pg_ctl -D /var/lib/postgresql/data start

# Register the standby
repmgr -f /etc/repmgr.conf standby register
```

### Monitoring

```bash
# Show cluster status
repmgr -f /etc/repmgr.conf cluster show

# Show replication status
repmgr -f /etc/repmgr.conf node status
```

### Manual Switchover

Promote a standby to primary (planned switchover):

```bash
repmgr -f /etc/repmgr.conf standby switchover
```

### Automatic Failover with repmgrd

Start the repmgr daemon on each node:

```bash
repmgrd -f /etc/repmgr.conf
```

Configure failover in `repmgr.conf`:

```ini
failover='automatic'
promote_command='repmgr standby promote -f /etc/repmgr.conf'
follow_command='repmgr standby follow -f /etc/repmgr.conf --upstream-node-id=%n'
```

### Key Commands

- `repmgr primary register` - register a primary node
- `repmgr standby clone` - clone a standby from primary
- `repmgr standby register` - register a standby node
- `repmgr standby promote` - promote standby to primary
- `repmgr standby follow` - follow a new primary
- `repmgr standby switchover` - planned switchover
- `repmgr cluster show` - display cluster status
- `repmgr cluster event` - show cluster events
- `repmgr node check` - health check a node
