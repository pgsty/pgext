## Usage

Sources:

- [Official README](https://github.com/instructure/jsoncdc/blob/46a04f733b554d792c2609f50831a5edd5a256af/README.md)
- [Official PGXN documentation](https://github.com/instructure/jsoncdc/blob/46a04f733b554d792c2609f50831a5edd5a256af/doc/jsoncdc.md)
- [PostgreSQL logical decoding documentation](https://www.postgresql.org/docs/current/logicaldecoding-explanation.html)

`jsoncdc` version `0.1.0` is a logical-decoding output plug-in that translates PostgreSQL WAL changes into newline-delimited JSON. It is loaded by a logical replication slot; it does not require `CREATE EXTENSION` and installs no SQL API.

### Core Workflow

Configure logical decoding first: `wal_level` must be `logical`, and `max_replication_slots` must leave capacity for the slot. Then connect to the database with privileges to create and consume a logical slot.

```sql
SELECT *
FROM pg_create_logical_replication_slot('jsoncdc_slot', 'jsoncdc');

-- After transactions modify user tables:
SELECT *
FROM pg_logical_slot_get_changes('jsoncdc_slot', NULL, NULL);
```

The stream emits a begin object, table schema/name records, insert/update/delete records, optional logical-message records, and a commit object with transaction ID and timestamp. Each JSON object occupies one output line.

### Replication-Slot Operations

```sql
SELECT slot_name, active, restart_lsn, confirmed_flush_lsn
FROM pg_replication_slots
WHERE slot_name = 'jsoncdc_slot';

SELECT pg_drop_replication_slot('jsoncdc_slot');
```

Logical slots persist independently of consumers and retain required WAL and catalog rows. Monitor lag, make consumers tolerant of replay after a crash, and drop a slot when it is no longer needed.

### Compatibility Boundary

The upstream code is an old Rust output plug-in, last documented for version `0.1.0`. Verify that it builds against the exact PostgreSQL server before adoption. DDL is not a row-change stream, and clients must parse the plug-in's documented record sequence rather than assuming one self-contained JSON document per transaction.
