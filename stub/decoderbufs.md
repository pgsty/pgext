## Usage

Sources:

- [Debezium decoderbufs 3.6.0.Final README](https://github.com/debezium/postgres-decoderbufs/blob/v3.6.0.Final/README.md)
- [Output-plugin control file](https://github.com/debezium/postgres-decoderbufs/blob/v3.6.0.Final/decoderbufs.control)
- [Protocol Buffers schema](https://github.com/debezium/postgres-decoderbufs/blob/v3.6.0.Final/proto/pg_logicaldec.proto)

`decoderbufs` is a headless PostgreSQL logical-decoding output plugin used by the Debezium PostgreSQL connector. It turns WAL changes into Protocol Buffers messages; it does not create a user SQL schema and does not require `CREATE EXTENSION`.

### Configure PostgreSQL

Enable the plugin and logical replication in `postgresql.conf`, size the sender and slot limits for the expected consumers, then restart PostgreSQL:

```conf
shared_preload_libraries = 'decoderbufs'
wal_level = logical
max_wal_senders = 8
max_replication_slots = 4
```

The replication login also needs the `REPLICATION` attribute and a matching `pg_hba.conf` rule. Use authentication appropriate for the network rather than the README's local demonstration settings.

### Core Workflow

Create a logical slot whose output plugin is `decoderbufs`:

```sql
SELECT *
FROM pg_create_logical_replication_slot('decoderbufs_demo', 'decoderbufs');
```

For inspection in `psql`, ask the plugin for debug text:

```sql
SELECT data
FROM pg_logical_slot_peek_changes(
  'decoderbufs_demo', NULL, NULL, 'debug-mode', '1'
);

SELECT data
FROM pg_logical_slot_get_changes(
  'decoderbufs_demo', NULL, NULL, 'debug-mode', '1'
);
```

`peek` leaves the confirmed position unchanged; `get` advances it. Normal Debezium operation consumes the binary messages defined by `pg_logicaldec.proto` rather than enabling debug mode.

### Important Objects and Boundaries

- `decoderbufs` is the logical-decoding output-plugin name passed when a slot is created.
- `debug-mode = 1` provides human-readable output for troubleshooting only.
- The Protobuf message carries transaction metadata, relation and column information, operation kind, old keys, and typed values.
- Tables that must emit sufficient data for `UPDATE` and `DELETE` require an appropriate `REPLICA IDENTITY`.

Logical slots retain WAL until a consumer confirms progress. Monitor `pg_replication_slots` and remove abandoned slots deliberately to prevent disk exhaustion. Schema changes, replica identity, unsupported data-type mappings, and large transactions should be tested with the matching Debezium connector version.

The upstream build requires PostgreSQL 9.6 or newer and protobuf-c; PostGIS support is compiled when available. The package release moves with Debezium to 3.6.0.Final, while the plugin control metadata remains SQL version 0.1.0 because this is an output plugin rather than a migration-based SQL extension.
