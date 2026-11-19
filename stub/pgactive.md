## Usage

Sources:

- [Official source](https://github.com/aws/pgactive/blob/v2.1.10/README.md)
- [Official source](https://github.com/aws/pgactive/blob/v2.1.10/docs/quickstart.md)
- [Official source](https://github.com/aws/pgactive/blob/v2.1.10/docs/README.md)
- [Official source](https://github.com/aws/pgactive/blob/v2.1.10/docs/referenceguide.md)
- [Official source](https://github.com/aws/pgactive/blob/v2.1.10/pgactive.control.in)
- [Official source](https://github.com/aws/pgactive/blob/v2.1.10/include/pgactive_version.h.in)
- [Official source](https://github.com/aws/pgactive/blob/v2.1.10/pgactive--2.1.10.sql)
- [Official source](https://github.com/aws/pgactive/releases/tag/v2.1.10)

`pgactive` 2.1.10 provides active-active logical replication between databases in a full mesh. Writes commit locally and replicate to peers; applications must handle lag, conflicts and node-local locking rather than assume global serializability.

### Configure a group

The release requires PostgreSQL 18.7, 17.12, 16.16, 15.20 or 14.25 and later minors in those branches; it also names upcoming PostgreSQL 19. Preload the library, configure logical WAL and sufficient replication slots/workers, then restart. Installation and management functions require superuser access.

```ini
shared_preload_libraries = 'pgactive'
wal_level = logical
```

```sql
CREATE EXTENSION pgactive;
SELECT pgactive.pgactive_create_group(
  node_name := 'node1-app',
  node_dsn := 'user_mapping=postgres pgactive_foreign_server=node1'
);
SELECT pgactive.pgactive_wait_for_node_ready();
```

The example assumes an existing node1 foreign server using pgactive_fdw, a protected user mapping and working connections to the database itself and every peer. Configure actual credentials through protected mappings or password files. On a fresh joining database, create the extension and use pgactive.pgactive_join_group with its own node_dsn and the peer's join_using_dsn, then wait for readiness. The built-in join performs coordinated initialization; an arbitrary independent dump does not substitute for it.

### Operations and upgrades

Every peer must be reachable when a node joins. Plan unique keys, sequence allocation, conflict handling and replication-slot retention. DDL replication is disabled by default; enabling it requires consistent policy on senders and receivers. Apply workers default to applying DML as the table owner, while replication connections and worker management remain privileged.

Version 2.1.10 hardens replication-name and tuple handling and fixes conflict-handler ABI and shared-memory lock offsets. A 2.1.10 node may join an older group, but a node on 2.1.9 or earlier cannot join a 2.1.10 node until upgraded. Install matching files and execute the supplied SQL upgrade in each database; coordinate restart/reconnection and join order across the group.
