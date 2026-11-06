## Usage

Sources:

- [README v5.5.1](https://github.com/EnterpriseDB/repmgr/blob/v5.5.1/README.md)
- [Control](https://github.com/EnterpriseDB/repmgr/blob/v5.5.1/repmgr.control)
- [SQL 5.5](https://github.com/EnterpriseDB/repmgr/blob/v5.5.1/repmgr--5.5.sql)
- [Quickstart](https://github.com/EnterpriseDB/repmgr/blob/v5.5.1/doc/quickstart.xml)
- [repmgrd configuration](https://github.com/EnterpriseDB/repmgr/blob/v5.5.1/doc/repmgrd-configuration.xml)
- [Upgrade guide](https://github.com/EnterpriseDB/repmgr/blob/v5.5.1/doc/upgrading-repmgr.xml)
- [5.5.1 release](https://github.com/EnterpriseDB/repmgr/releases/tag/v5.5.1)

`repmgr` manages PostgreSQL physical replication, monitoring and failover. Release 5.5.1 supports PostgreSQL 13–18 and retains SQL extension version 5.5. It fixes escaping of values substituted into external notification commands. Use the same repmgr release on all cluster nodes.

### Enabling

Use a dedicated metadata database and management role, with replication connections and the required SSH access configured. The upstream quickstart uses a superuser role; a normal role needs `--superuser` for privileged operations. Run the following SQL in the metadata database with sufficient privileges:

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

Before using `repmgrd`, configure this on each node and restart PostgreSQL:

```ini
shared_preload_libraries = 'repmgr'
```

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

### 5.5.1 Upgrade

Pause or stop automatic failover as directed by the upgrade guide, update the programs on each node, then restart a running `repmgrd` and check cluster status. Both 5.5.0 and 5.5.1 use SQL version `5.5`; do not request a nonexistent SQL version 5.5.1. Follow the appropriate extension upgrade chain when coming from an older SQL version.
