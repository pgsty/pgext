## Usage

Sources:

- [Official PGD documentation](https://www.enterprisedb.com/docs/pgd/6.5/)
- [Official node-creation guide](https://www.enterprisedb.com/docs/pgd/6.5/node_management/creating_nodes/)
- [Official node management API](https://www.enterprisedb.com/docs/pgd/6.5/reference/nodes-management-interfaces/)

`bdr` is the provider extension at the core of EDB Postgres Distributed (PGD), providing active-active logical replication, node groups, conflict management, consensus, and distributed commit policies.

### Install and Preload

Install the matching PGD 6.5 package on every node, enable commit timestamps, preload `bdr`, and restart all nodes before creating the extension:

```ini
shared_preload_libraries = 'bdr'
track_commit_timestamp = on
```

```sql
CREATE EXTENSION bdr;
SELECT bdr.bdr_version();
```

Only a superuser creates the extension. PGD administration should otherwise use the predefined `bdr_` roles.

### Initialize and Join a Group

```sql
SELECT bdr.create_node(
  node_name := 'node_a',
  dsn       := 'host=node-a dbname=app'
);

SELECT bdr.create_node_group(
  node_group_name := 'app_group'
);
```

On another prepared node, create `bdr`, call `bdr.create_node`, and join with `bdr.join_node_group`. Use `bdr.wait_for_join_completion` before directing traffic.

### Distributed-System Boundaries

Group creation, join, parting, and consensus operations are asynchronous and are not rolled back with the surrounding SQL transaction. Every node must have compatible PGD/PostgreSQL binaries and extension packages before DDL replication reaches it. Plan replication slots/origins, conflict policy, commit scopes, sequence behavior, DDL, role propagation, backup/restore, node loss, and network partitions; a successful `CREATE EXTENSION` proves only local installation, not cluster health.
