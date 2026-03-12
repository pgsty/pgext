

## Usage

> [decoderbufs: Logical decoding plugin that delivers WAL stream changes using a Protocol Buffer format](https://github.com/debezium/postgres-decoderbufs)

A PostgreSQL logical decoding output plugin that serializes WAL changes into Protocol Buffers format, primarily used by the Debezium PostgreSQL connector for change data capture.

### Configuration

In `postgresql.conf`:

```ini
shared_preload_libraries = 'decoderbufs'
wal_level = logical
max_wal_senders = 8
max_replication_slots = 4
```

### Using with SQL (Debug Mode)

```sql
-- Create a logical replication slot
SELECT * FROM pg_create_logical_replication_slot('decoderbufs_demo', 'decoderbufs');

-- Perform table modifications
INSERT INTO my_table VALUES (1, 'test');
UPDATE my_table SET col = 'updated' WHERE id = 1;

-- Peek at changes in debug text mode
SELECT data FROM pg_logical_slot_peek_changes(
    'decoderbufs_demo', NULL, NULL, 'debug-mode', '1');

-- Consume changes
SELECT data FROM pg_logical_slot_get_changes(
    'decoderbufs_demo', NULL, NULL, 'debug-mode', '1');

-- Check slot status
SELECT * FROM pg_replication_slots WHERE slot_type = 'logical';
```

### Type Mappings

| PostgreSQL Type    | Protobuf Field   |
|--------------------|------------------|
| BOOL               | datum_boolean    |
| INT2, INT4         | datum_int32      |
| INT8, OID          | datum_int64      |
| FLOAT4             | datum_float      |
| FLOAT8, NUMERIC    | datum_double     |
| CHAR, VARCHAR, TEXT | datum_string    |
| JSON, XML, UUID    | datum_string     |
| TIMESTAMP(TZ)      | datum_string     |
| BYTEA              | datum_bytes      |
| POINT, PostGIS     | datum_point      |

### Notes

- For UPDATE/DELETE, set [REPLICA IDENTITY](https://www.postgresql.org/docs/current/sql-altertable.html#SQL-CREATETABLE-REPLICA-IDENTITY) appropriately
- Binary Protocol Buffer output is consumed by the Debezium Postgres Connector
- `debug-mode` option provides human-readable text output for SQL console testing
- Requires `protobuf-c` library and PostGIS development packages for compilation
