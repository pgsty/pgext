

## Usage

> [pg_uuidv7: Create valid version 7 UUIDs in PostgreSQL](https://github.com/fboulnois/pg_uuidv7)

```sql
CREATE EXTENSION pg_uuidv7;
```

### Functions

| Function | Description |
|---|---|
| `uuid_generate_v7()` | Generate a new UUIDv7 |
| `uuid_v7_to_timestamptz(uuid)` | Extract the timestamp from a UUIDv7 |
| `uuid_timestamptz_to_v7(timestamptz [, bool])` | Convert a timestamp to a UUIDv7 (set second arg to `true` to zero the random bits) |

### Examples

```sql
-- Generate a UUIDv7
SELECT uuid_generate_v7();
-- 018570bb-4a7d-7c7e-8df4-6d47afd8c8fc

-- Extract timestamp from UUIDv7
SELECT uuid_v7_to_timestamptz('018570bb-4a7d-7c7e-8df4-6d47afd8c8fc');
-- 2023-01-02 04:26:40.637+00

-- Convert timestamp to UUIDv7
SELECT uuid_timestamptz_to_v7('2023-01-02 04:26:40.637+00');
-- 018570bb-4a7d-7630-a5c4-89b795024c5d

-- For date range queries, zero the random bits
SELECT uuid_timestamptz_to_v7('2023-01-02 04:26:40.637+00', true);
-- 018570bb-4a7d-7000-8000-000000000000

-- Use as primary key
CREATE TABLE events (
  id uuid NOT NULL DEFAULT uuid_generate_v7() PRIMARY KEY,
  data text
);
```
