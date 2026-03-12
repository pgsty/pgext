

## Usage

> [wal2json: Changing data capture in JSON format](https://github.com/eulerto/wal2json)

A logical decoding output plugin that produces JSON-formatted change data capture from the PostgreSQL WAL.

### Configuration

In `postgresql.conf`:

```ini
wal_level = logical
max_replication_slots = 10
max_wal_senders = 10
```

### Using with Streaming Protocol (pg_recvlogical)

```bash
# Create a replication slot
pg_recvlogical -d postgres --slot test_slot --create-slot -P wal2json

# Start consuming changes
pg_recvlogical -d postgres --slot test_slot --start -o pretty-print=1 -f -

# Drop the slot when done
pg_recvlogical -d postgres --slot test_slot --drop-slot
```

### Using with SQL Functions

```sql
-- Create a logical replication slot
SELECT * FROM pg_create_logical_replication_slot('test_slot', 'wal2json');

-- Peek at changes (does not consume)
SELECT data FROM pg_logical_slot_peek_changes('test_slot', NULL, NULL);

-- Get and consume changes
SELECT data FROM pg_logical_slot_get_changes('test_slot', NULL, NULL,
    'pretty-print', '1');

-- Drop the slot
SELECT pg_drop_replication_slot('test_slot');
```

### Output Format v1 (JSON per transaction)

```json
{
  "change": [
    {
      "kind": "insert",
      "schema": "public",
      "table": "my_table",
      "columnnames": ["a", "b"],
      "columntypes": ["integer", "text"],
      "columnvalues": [1, "hello"]
    },
    {
      "kind": "delete",
      "schema": "public",
      "table": "my_table",
      "oldkeys": {
        "keynames": ["a"],
        "keytypes": ["integer"],
        "keyvalues": [1]
      }
    }
  ]
}
```

### Output Format v2 (JSON per tuple)

Enable with: `'format-version', '2'`

### Key Parameters

- `include-xids` - add transaction ID (default: false)
- `include-timestamp` - add timestamp (default: false)
- `include-schemas` - add schema name (default: true)
- `include-types` - add column types (default: true)
- `include-pk` - add primary key info (default: false)
- `include-lsn` - add WAL LSN (default: false)
- `include-not-null` - add NOT NULL info (default: false)
- `include-default` - add default expressions (default: false)
- `pretty-print` - format JSON output (default: false)
- `filter-tables` - comma-separated list of tables to include
- `add-tables` - same as filter-tables
- `filter-msg-prefixes` - filter logical messages by prefix
- `format-version` - 1 (per-transaction) or 2 (per-tuple)
- `actions` - filter by action type: insert, update, delete, truncate
