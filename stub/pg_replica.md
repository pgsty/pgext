## Usage

Sources:

- [README.md](https://api.pgxn.org/src/pg_replica/pg_replica-0.7.4/README.md)
- [pg_replica.control](https://api.pgxn.org/src/pg_replica/pg_replica-0.7.4/pg_replica.control)
- [Cargo.toml](https://api.pgxn.org/src/pg_replica/pg_replica-0.7.4/Cargo.toml)

`pg_replica` 0.7.4 adds embedded Raft leadership and failover to physical PostgreSQL replication. It targets PostgreSQL 18 and requires a configured cluster, preload and restart; extension creation alone does not establish replication.

### Core Workflow

```ini
shared_preload_libraries = 'pg_replica'
```

```sql
CREATE EXTENSION pg_replica;
SELECT replica.status();
```

### Operational Boundaries

`replica.status()` reports leadership and replication state. `replica.rotate_credential()` changes a credential on the current primary and is a no-op on standbys. A supervisor still manages process restarts, and clients need reconnect/routing logic. Quorum-synchronous durability is opt-in; do not describe asynchronous operation as lossless. Rejoining can use rewind or a complete base backup, and adopting a restored primary can discard data from the former primary. Review node membership, credentials, fencing and recovery procedures before enabling the background workers. Installation requires a superuser.
