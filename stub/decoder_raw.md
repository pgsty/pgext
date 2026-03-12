

## Usage

> [decoder_raw: Output plugin for logical replication in Raw SQL format](https://github.com/michaelpq/pg_plugins/blob/main/decoder_raw/)

A logical decoding output plugin that converts WAL changes into raw SQL statements. Part of the pg_plugins collection by Michael Paquier.

### Configuration

In `postgresql.conf`:

```ini
wal_level = logical
max_replication_slots = 10
max_wal_senders = 10
```

### Using with SQL Functions

```sql
-- Create a logical replication slot
SELECT * FROM pg_create_logical_replication_slot('raw_slot', 'decoder_raw');

-- Perform DML operations
INSERT INTO my_table VALUES (1, 'hello');
UPDATE my_table SET val = 'world' WHERE id = 1;
DELETE FROM my_table WHERE id = 1;

-- Get changes as raw SQL
SELECT data FROM pg_logical_slot_get_changes('raw_slot', NULL, NULL);
-- Output:
-- INSERT INTO public.my_table (id, val) VALUES (1, 'hello');
-- UPDATE public.my_table SET val = 'world' WHERE id = 1;
-- DELETE FROM public.my_table WHERE id = 1;

-- Drop the slot
SELECT pg_drop_replication_slot('raw_slot');
```

### Using with pg_recvlogical

```bash
# Create slot
pg_recvlogical -d postgres --slot raw_slot --create-slot -P decoder_raw

# Stream changes as SQL statements
pg_recvlogical -d postgres --slot raw_slot --start -f -

# Drop slot
pg_recvlogical -d postgres --slot raw_slot --drop-slot
```

### Key Features

- Outputs changes as executable SQL statements (INSERT, UPDATE, DELETE)
- Useful for debugging logical decoding or replaying changes on another database
- Tables should have REPLICA IDENTITY set for proper UPDATE/DELETE output
- Lightweight plugin designed as a template for custom logical decoding plugins
