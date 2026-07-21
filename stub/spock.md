


## Usage

Sources:

- [Spock v5.0.10 README](https://github.com/pgEdge/spock/blob/v5.0.10/README.md)
- [Getting started](https://github.com/pgEdge/spock/blob/v5.0.10/docs/getting_started.md)
- [Configuration reference](https://github.com/pgEdge/spock/blob/v5.0.10/docs/configuring.md)
- [Limitations](https://github.com/pgEdge/spock/blob/v5.0.10/docs/limitations.md)
- [Release notes](https://github.com/pgEdge/spock/blob/v5.0.10/docs/spock_release_notes.md)

`spock` provides active-active logical replication for PostgreSQL 15 through 18. Each participating database is a Spock node; a multi-master topology is formed by creating directed subscriptions between nodes.

### Configuration

In `postgresql.conf`:

```ini
wal_level = 'logical'
max_worker_processes = 10
max_replication_slots = 10
max_wal_senders = 10
shared_preload_libraries = 'spock'
track_commit_timestamp = on
spock.enable_ddl_replication = on
spock.include_ddl_repset = on
```

### Enabling

```sql
CREATE EXTENSION spock;
```

### Creating Nodes

On each node, create a node identity:

```sql
-- Node 1
SELECT spock.node_create(
    node_name := 'n1',
    dsn := 'host=10.0.0.5 port=5432 dbname=mydb'
);

-- Node 2
SELECT spock.node_create(
    node_name := 'n2',
    dsn := 'host=10.0.0.7 port=5432 dbname=mydb'
);
```

### Creating Subscriptions

For multi-master, each node subscribes to every other node:

```sql
-- On n1: subscribe to n2
SELECT spock.sub_create(
    subscription_name := 'sub_n1n2',
    provider_dsn := 'host=10.0.0.7 port=5432 dbname=mydb'
);

-- On n2: subscribe to n1
SELECT spock.sub_create(
    subscription_name := 'sub_n2n1',
    provider_dsn := 'host=10.0.0.5 port=5432 dbname=mydb'
);
```

### Replication Set Management

```sql
-- Add table to replication
SELECT spock.repset_add_table('default', 'my_table');

-- Remove table from replication
SELECT spock.repset_remove_table('default', 'my_table');

-- Add all tables in a schema
SELECT spock.repset_add_all_tables('default', '{public}');
```

### Key Features

- Multi-master (active-active) replication
- Automatic DDL replication
- Conflict detection and resolution using commit timestamps
- Row and column filtering
- Supports PostgreSQL 15, 16, 17, and 18
- Tables must have primary keys and matching schemas across nodes

### Operations and Caveats

- Install `spock` and add it to `shared_preload_libraries` on every participating server before creating nodes or subscriptions.
- Keep table definitions, data types, primary keys, and relevant unique indexes identical across nodes. Coordinate DDL even when DDL replication is enabled.
- Replicated tables need a primary key or another usable replica identity. Temporary and unlogged tables are not replication targets.
- Spock operates per database. Repeat extension and topology setup for each database that participates.
- Active-active conflict handling depends on commit timestamps and policy. Test simultaneous inserts and updates, especially nullable unique keys, before production use.
- Upstream documents platform/build requirements in the README; verify that the PostgreSQL build and Spock package used on every node are compatible.

### Version 5.0.10

`5.0.10` is a patch release in the 5.0 line. Its release notes include fixes for unique indexes containing `NULL`, `NULLS NOT DISTINCT` conflict handling, refreshing cached index metadata after an index is dropped, exception-path memory handling, and numerical version checks used during rolling patch upgrades. Upgrade every node to a compatible patch level and validate subscriptions after the rolling change.
