## Usage

Sources:

- [README.md](https://github.com/ankus-net/worf/blob/7299f69272cbde9ee01b0cd03e46fc7ddeb54597/README.md)
- [LICENSE](https://github.com/ankus-net/worf/blob/7299f69272cbde9ee01b0cd03e46fc7ddeb54597/LICENSE)
- [src/Worf/Worf.csproj](https://github.com/ankus-net/worf/blob/7299f69272cbde9ee01b0cd03e46fc7ddeb54597/src/Worf/Worf.csproj)
- [src/Worf/worf.control](https://github.com/ankus-net/worf/blob/7299f69272cbde9ee01b0cd03e46fc7ddeb54597/src/Worf/worf.control)
- [src/Worf/Worf.cs](https://github.com/ankus-net/worf/blob/7299f69272cbde9ee01b0cd03e46fc7ddeb54597/src/Worf/Worf.cs)
- [src/Worf/Node.cs](https://github.com/ankus-net/worf/blob/7299f69272cbde9ee01b0cd03e46fc7ddeb54597/src/Worf/Node.cs)
- [.github/workflows/ci.yml](https://github.com/ankus-net/worf/blob/7299f69272cbde9ee01b0cd03e46fc7ddeb54597/.github/workflows/ci.yml)

`worf` provides per-database logical replication with replication sets, multiple providers and configurable conflict handling. This C#/.NET Native AOT source declares version 0.1.0 in both its Ankus project and runtime; Ankus generates the matching control and versioned SQL from the control file and SQL attributes. No tagged release was available at review. Upstream supports PostgreSQL 15–19.

### Configure Both Nodes

Install matching files on provider and subscriber. Preserve existing preload libraries, configure logical WAL and enough workers, slots and senders, then restart PostgreSQL:

```ini
wal_level = logical
shared_preload_libraries = 'worf'
max_worker_processes = 10
max_replication_slots = 10
max_wal_senders = 10
worf.conflict_resolution = 'error'
```

These capacities are example values; size them for the actual nodes and other consumers. On PostgreSQL 15.19, 16.15, 17.11, 18.6, 19 or newer, the provider must also allow worf in the output-plugin allow-list while preserving other required plugins:

```ini
output_plugin_libraries = 'pgoutput, test_decoding, worf'
```

Superuser administration is required. Install in each participating database, and create matching schemas, tables, column types and primary keys on both sides. Subscriber constraints must be no stricter than the provider's.

### Provider and Subscriber

On the provider:

```sql
CREATE EXTENSION worf;
SELECT worf.create_node(
  node_name := 'provider1',
  dsn := 'host=providerhost port=5432 dbname=app'
);
SELECT worf.replication_set_add_all_tables('default', ARRAY['public']);
```

On the subscriber:

```sql
CREATE EXTENSION worf;
SELECT worf.create_node(
  node_name := 'subscriber1',
  dsn := 'host=subscriberhost port=5432 dbname=app'
);
SELECT worf.create_subscription(
  subscription_name := 'subscription1',
  provider_dsn := 'host=providerhost port=5432 dbname=app'
);
SELECT worf.wait_for_subscription_sync_complete('subscription1');
SELECT * FROM worf.show_subscription_status();
```

Configure provider authentication for normal database replication connections, and protect connection credentials. The fixed `worf` schema contains node, subscription and replication-set state. `worf.replicate_ddl_command` can queue an explicitly chosen DDL operation; ordinary DDL is not automatically replicated. `worf_create_subscriber` clones with a full base backup and does not limit the copy to selected replication sets.

### Conflicts and Maintenance Boundaries

The default conflict policy is `apply_remote`, which can overwrite local rows; the example deliberately chooses `error` to stop for review. Timestamp-based policies and `keep_local` require `track_commit_timestamp`. Multi-provider conflict handling does not establish complete multi-master semantics. Monitor replication slots and WAL retention, synchronization state and logged conflicts.

UPDATE/DELETE require a usable primary key or replica identity; `REPLICA IDENTITY FULL` is unsupported. Extra unique constraints can stop conflict application, and deferrable unique indexes are unsupported. Foreign-key constraints are not enforced during apply. UNLOGGED/TEMPORARY tables and large objects are not replicated. Sequence state propagates periodically rather than per increment. Table resynchronization truncates immediately and leaves the table empty while data is copied. Plan schema changes and synchronization as maintenance operations, not automatic failover guarantees.
