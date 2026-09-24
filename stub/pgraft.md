## Usage

Sources:

- [README.md](https://github.com/pgElephant/pgraft/blob/bc816afee21018a6836356fcc8d2cbf44753be54/README.md)
- [pgraft.control](https://github.com/pgElephant/pgraft/blob/bc816afee21018a6836356fcc8d2cbf44753be54/pgraft.control)
- [pgraft--2.0.0.sql](https://github.com/pgElephant/pgraft/blob/bc816afee21018a6836356fcc8d2cbf44753be54/pgraft--2.0.0.sql)
- [docs/user-guide/configuration.md](https://github.com/pgElephant/pgraft/blob/bc816afee21018a6836356fcc8d2cbf44753be54/docs/user-guide/configuration.md)

`pgraft` provides Raft consensus, leader election, and replicated key-value state through C and Go libraries. The pinned 2.0.0 source succeeds the older RAM subproject and targets PostgreSQL 14–18; it does not make arbitrary application tables replicate automatically.

### Configure and Enable

On every node, preload `pgraft` and configure a unique `pgraft.name`, the same ordered `pgraft.initial_cluster` list and `pgraft.initial_cluster_token`, local `pgraft.listen_peer_urls`, and durable `pgraft.data_dir`. Configure peer connectivity and restart the server. Then enable the SQL API in the database.

```sql
CREATE EXTENSION pgraft;
SELECT * FROM pgraft.get_cluster_status();
SELECT * FROM pgraft.get_nodes();
SELECT pgraft.is_leader(), pgraft.get_leader();
```

### SQL Operations

`pgraft.kv_put` and `pgraft.kv_delete` mutate key-value state on the leader. `pgraft.kv_get` reads local state. `pgraft.add_node` and `pgraft.remove_node` change membership and must also be issued on the leader. `pgraft.log_get_replication_status` and `pgraft.log_get_stats` inspect replication. Preserve logs, snapshots, and HardState in the configured durable directory; membership and recovery require quorum-aware operation.

### Upgrade Boundary

Version 2.0.0 moves the old unqualified function names into the `pgraft` schema and removes their prefix. For example, `pgraft_get_cluster_status` becomes `pgraft.get_cluster_status`. Update application calls with the extension upgrade; dependent objects can block the upgrade because it does not cascade.

```sql
ALTER EXTENSION pgraft UPDATE TO '2.0.0';
```

Restart after installing and upgrading so both native libraries reload together. Mutating functions are no longer executable by `PUBLIC`; grant only needed operations. These notes describe pinned 2.0.0 source, while the repository’s published release list still shows 1.0 releases.
