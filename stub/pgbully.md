## Usage

Sources:

- [README](https://github.com/pgElephant/pgbully/blob/259558c49216ea011b7444fe763bebf3bf0ca2ab/README.md)
- [Control file](https://github.com/pgElephant/pgbully/blob/259558c49216ea011b7444fe763bebf3bf0ca2ab/pgbully.control)
- [sql/pgbully--1.0.sql](https://github.com/pgElephant/pgbully/blob/259558c49216ea011b7444fe763bebf3bf0ca2ab/sql/pgbully--1.0.sql)

`pgbully` elects the highest reachable node ID as leader using the Bully algorithm. It coordinates independent PostgreSQL nodes; it does not replicate application tables or implement failover and fencing.

### Configure the Cluster

On every PostgreSQL 15–18 node, configure a unique node ID and an identical peer list, with working libpq authentication between nodes. Append the library to the preload list and restart. Example for node 1:

```conf
shared_preload_libraries = 'pgbully'
pgbully.node_id = 1
pgbully.nodes = '1: host=10.0.0.1 port=5432 dbname=postgres, 2: host=10.0.0.2 port=5432 dbname=postgres, 3: host=10.0.0.3 port=5432 dbname=postgres'
pgbully.heartbeat_interval = '1s'
pgbully.election_timeout = '5s'
```

### SQL Workflow

```sql
CREATE EXTENSION pgbully;
SELECT * FROM pgbully.status();
SELECT node_id, is_leader, reachable, last_seen FROM pgbully.cluster;
SELECT pgbully.is_leader(), pgbully.get_leader(), pgbully.get_term();
```

### Failure Boundaries

Installation requires a superuser. Status and membership objects live in the `pgbully` schema. `pgbully.kv_put` is leader-only and `pgbully.kv_get` reads local state; the optional store is best-effort rather than quorum-durable.

A network partition can elect one leader in each partition. Applications requiring a globally unique writer must provide an independent quorum/fencing mechanism; a successful leader check is not sufficient authorization to promote a database or perform an irreversible external action.
