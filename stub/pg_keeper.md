## Usage

Sources:

- [Official README](https://github.com/MasahikoSawada/pg_keeper/blob/4bfc5773fae35f55daf6c6015a947456209c51a8/README.md)
- [Main background-worker source](https://github.com/MasahikoSawada/pg_keeper/blob/4bfc5773fae35f55daf6c6015a947456209c51a8/pg_keeper.c)
- [Standby promotion source](https://github.com/MasahikoSawada/pg_keeper/blob/4bfc5773fae35f55daf6c6015a947456209c51a8/standby.c)

`pg_keeper` is a legacy two-node PostgreSQL clustering module. A preloaded background worker on each node polls its partner: the primary switches from synchronous to asynchronous replication after repeated standby failures, while the standby promotes itself after repeated primary failures. It is not a consensus system and does not by itself prevent split brain.

### Configuration Boundary

The upstream repository builds a shared module only; it contains no extension control or versioned SQL file. Configure and preload it on both an already-working primary/standby pair rather than issuing `CREATE EXTENSION`:

```conf
shared_preload_libraries = 'pg_keeper'
max_worker_processes = 8

pg_keeper.keepalives_time = 5
pg_keeper.keepalives_count = 3
pg_keeper.my_conninfo = 'host=127.0.0.1 port=5432 dbname=postgres'
pg_keeper.partner_conninfo = 'host=partner port=5432 dbname=postgres'
pg_keeper.after_command = '/usr/local/sbin/verified-fencing-hook'
```

The source spells the interval and count settings as `keepalives_time` and `keepalives_count`; preserve those exact names. `partner_conninfo` is used for heartbeat queries and `my_conninfo` for local `ALTER SYSTEM` operations. Both connection strings are mandatory.

### Failure Behavior

On a synchronous primary, a heartbeat failure count reaching the configured threshold causes `pg_keeper` to clear `synchronous_standby_names` through `ALTER SYSTEM`, allowing transactions to proceed asynchronously. Restoring synchronous replication after the standby recovers is an operator action.

On a standby, reaching the same threshold creates the historical promotion trigger and promotes the local server. The optional `after_command` runs after promotion and was intended for fencing or STONITH integration. A reachable-database check cannot distinguish primary failure from a network partition, so reliable external fencing is essential.

Verify worker states and log transitions on both nodes after every restart. The approximate detection interval is the polling interval multiplied by the failure count, plus connection timeout and scheduling overhead.

### Compatibility and Safety

The README identifies its own release as 1.0, while the catalog entry reports 2.0; the upstream repository has no control file that substantiates an extension version. Record that discrepancy rather than assuming catalog 2.0 represents a newer upstream API.

The project was tested on CentOS 6.5 with PostgreSQL 9.5 and 9.6 and last changed in 2019. Its documentation uses removed recovery configuration and its promotion code targets old server behavior. Do not deploy it on modern PostgreSQL without a full source port and destructive failover tests. Prefer a maintained HA system with distributed consensus and proven fencing. Connection credentials, `ALTER SYSTEM` permissions, promotion, synchronous-replication loss, failback, timeline divergence, and double-primary recovery all require explicit operational design.
