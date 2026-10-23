## Usage

Sources:

- [Official source](https://github.com/MasahikoSawada/pg_raft/blob/320d2cb89a5a971ed2b3f9430b16a5ed20344da1/README.md)
- [Official source](https://github.com/MasahikoSawada/pg_raft/blob/320d2cb89a5a971ed2b3f9430b16a5ed20344da1/pg_raft.control)
- [Official source](https://github.com/MasahikoSawada/pg_raft/blob/320d2cb89a5a971ed2b3f9430b16a5ed20344da1/pg_raft--1.0.sql)

`pg_raft` is an unfinished Raft experiment with SQL version 1.0, reviewed at source revision 320d2cb8. It provides node-management and leader-election machinery; the upstream checklist still leaves log replication, joining/leaving groups and the launcher incomplete.

### Experimental Setup

The upstream setup preloads the library, selects the worker database and restarts the server:

```ini
shared_preload_libraries = 'pg_raft'
pg_raft.database = 'postgres'
```

In that database, a superuser can follow the prototype's node-creation example:

```sql
CREATE EXTENSION pg_raft;
SELECT pgraft.create_node('node-1', 'port=5432 dbname=postgres');
```

The library must be included alongside any other preloaded libraries. This is an experiment, not a complete recipe for a replicated production cluster.

### Objects and Boundaries

The `pgraft` schema contains node, group, state and log objects, along with internal voting and append-entry RPC functions. These interfaces are implementation details of the prototype; the existence of an RPC function does not establish complete consensus or durable replicated-log behavior.

The reviewed repository does not establish a supported PostgreSQL-major range or an explicit license. Confirm both with upstream before adoption. Use isolated test clusters, protect connection information, and do not use the prototype to make availability or data-durability guarantees.
