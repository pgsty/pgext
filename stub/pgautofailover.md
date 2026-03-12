


## Usage

> [pgautofailover: pg_auto_failover](https://github.com/hapostgres/pg_auto_failover)

pg_auto_failover is an extension and service for automated PostgreSQL failover. It consists of a monitor node (running the `pgautofailover` extension), and keeper processes on each data node managed by the `pg_autoctl` CLI.

### Architecture

- **Monitor**: a PostgreSQL instance with the `pgautofailover` extension that implements a state machine for failover decisions
- **Keeper** (`pg_autoctl run`): runs on each data node, reports health to the monitor and executes state transitions
- Supports 2+ node setups with synchronous replication by default
- Supports Citus HA (since v2.0)

### Key CLI Commands

```bash
# Create the monitor
pg_autoctl create monitor --pgdata /path/to/monitor --pgport 5000

# Create a data node (primary or secondary auto-assigned)
pg_autoctl create postgres --pgdata /path/to/data --pgport 5001 --monitor postgres://monitor:5000/pg_auto_failover

# Run the keeper (foreground)
pg_autoctl run --pgdata /path/to/data

# Check cluster state
pg_autoctl show state --pgdata /path/to/monitor

# Perform a manual switchover
pg_autoctl perform switchover --pgdata /path/to/monitor

# Perform a manual failover
pg_autoctl perform failover --pgdata /path/to/monitor
```

### Failover Behavior

The monitor tracks node health. When a secondary becomes unavailable or its lag is too high, it is removed from `synchronous_standby_names` on the primary. Failover/switchover operations are blocked until the secondary is healthy again, preventing data loss.

### Documentation

Full documentation: [pg-auto-failover.readthedocs.io](https://pg-auto-failover.readthedocs.io/en/main/)
