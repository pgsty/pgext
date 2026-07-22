## Usage

Sources:

- [Official README](https://github.com/citusdata/pg_paxos/blob/7d4f656b1cfa0943626b658fbc23a375eab41c3a/README.md)
- [Extension control file](https://github.com/citusdata/pg_paxos/blob/7d4f656b1cfa0943626b658fbc23a375eab41c3a/pg_paxos.control)
- [SQL implementation](https://github.com/citusdata/pg_paxos/blob/7d4f656b1cfa0943626b658fbc23a375eab41c3a/sql/pg_paxos.sql)
- [Planner and executor hooks](https://github.com/citusdata/pg_paxos/blob/7d4f656b1cfa0943626b658fbc23a375eab41c3a/src/pg_paxos.c)

`pg_paxos` is an abandoned, experimental implementation of Multi-Paxos table replication for PostgreSQL 9.4–9.6-era servers. It logs SQL text and re-executes it on group members; use it only for historical study or isolated experiments, not as a supported high-availability or production replication system.

### Setup and Core Workflow

Install the declared `dblink` dependency on every node, preload `pg_paxos`, assign a unique node ID, and restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_paxos'
pg_paxos.node_id = 'node-a'
```

```sql
CREATE EXTENSION dblink;
CREATE EXTENSION pg_paxos;

CREATE TABLE coordinates (x integer, y integer);
SELECT paxos_create_group('geo', 'host=10.0.0.1 dbname=postgres');
SELECT paxos_replicate_table('geo', 'coordinates'::regclass);
```

On a new node with the same table definition and extensions, join the group with both the existing node's and the new node's connection strings:

```sql
SELECT paxos_join_group(
    'geo',
    'host=10.0.0.1 dbname=postgres',
    'host=10.0.0.2 dbname=postgres'
);
```

If another table is added after group formation, call `paxos_replicate_table(text, regclass)` on every node as directed by upstream.

### Configuration and Consistency

- `pg_paxos.node_id` must uniquely identify a node.
- `pg_paxos.enabled` enables interception for managed tables.
- `pg_paxos.consistency_model` is `strong` by default; `optimistic` reduces read coordination but can return stale data after failures.
- `pg_paxos.allow_mutable_functions` defaults to off because nondeterministic functions can make replicas diverge.

The planner and executor hooks require every relation referenced by a managed query to belong to one Paxos group. A write is transformed into SQL text, appended to the group log, and replayed. Writes on replicated tables are forbidden inside an explicit transaction block, and only single-statement execution is supported in the replay path.

### Safety and Compatibility

Connection strings are stored in `pgp_metadata` tables and used by `dblink`; protect the schema and avoid embedding long-lived passwords where possible. The SQL-log design requires deterministic behavior. Sequences, triggers, mutable functions, time and random values, external side effects, and schema drift can all cause divergence or duplicate effects. Enabling `pg_paxos.allow_mutable_functions` removes only one guard and does not make such operations safe.

Version `0.2` dates from 2016 and contains server-internal compatibility code only for PostgreSQL 9.4, 9.5, and 9.6. It lacks the operational maturity, upgrade path, monitoring, recovery procedures, and modern PostgreSQL support expected of production consensus systems. Treat `strong` as the algorithm's intended mode in this experimental implementation, not a current service guarantee; independently test failures, partitions, membership changes, crash recovery, and data reconciliation in disposable environments.
