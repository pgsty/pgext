


## Usage

> [old_snapshot: utilities in support of old_snapshot_threshold](https://www.postgresql.org/docs/16/oldsnapshot.html)

The `old_snapshot` extension provides inspection functions for the server state related to the `old_snapshot_threshold` configuration parameter.

Note: this chapter is removed from PostgreSQL current docs in newer releases; use versioned docs when needed.

### Function

```sql
-- View the timestamp-to-XID mapping table
SELECT * FROM pg_old_snapshot_time_mapping();
```

### Function Signature

```sql
pg_old_snapshot_time_mapping(
    array_offset OUT int4,
    end_timestamp OUT timestamptz,
    newest_xmin OUT xid
) RETURNS SETOF record
```

### Output Columns

| Column | Type | Description |
|--------|------|-------------|
| `array_offset` | int4 | Index position in the mapping array |
| `end_timestamp` | timestamptz | End of the corresponding one-minute interval |
| `newest_xmin` | xid | Newest xmin of any snapshot taken during that minute |

### Context

PostgreSQL's `old_snapshot_threshold` parameter controls how long a snapshot can remain valid. The server maintains an internal mapping of timestamps to transaction IDs to implement this feature. This extension exposes that mapping for inspection and debugging.

```sql
-- Check the old_snapshot_threshold setting
SHOW old_snapshot_threshold;

-- Inspect the current mapping entries
SELECT array_offset, end_timestamp, newest_xmin
FROM pg_old_snapshot_time_mapping()
ORDER BY array_offset;
```
