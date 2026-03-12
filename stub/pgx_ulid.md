

## Usage

> [pgx_ulid: ULID type and methods for PostgreSQL](https://github.com/pksunkara/pgx_ulid)

```sql
CREATE EXTENSION pgx_ulid;
```

### ULID Type

The extension provides a native `ulid` type -- a 26-character, lexicographically sortable identifier stored in binary.

### Functions

| Function | Description |
|---|---|
| `gen_ulid()` | Generate a new ULID |
| `gen_monotonic_ulid()` | Generate monotonically increasing ULIDs (requires `shared_preload_libraries`) |

### Casting

- `ulid::timestamp` -- extract creation time from a ULID
- `timestamp::ulid` -- produce a ULID from a timestamp (zeroed random part)
- `ulid::uuid` / `uuid::ulid` -- convert between ULID and UUID

### Examples

```sql
-- Use ULID as a primary key
CREATE TABLE users (
  id ulid NOT NULL DEFAULT gen_ulid() PRIMARY KEY,
  name text NOT NULL
);

-- Query by text representation
SELECT * FROM users WHERE id = '01ARZ3NDEKTSV4RRFFQ69G5FAV';

-- Extract timestamp from ULID
ALTER TABLE users
ADD COLUMN created_at timestamp GENERATED ALWAYS AS (id::timestamp) STORED;

-- Range query by date
SELECT * FROM users
WHERE id BETWEEN '2023-09-15'::timestamp::ulid AND '2023-09-16'::timestamp::ulid;
```
