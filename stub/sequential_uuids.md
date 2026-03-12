

## Usage

> [sequential_uuids: sequential UUID generators for better index locality](https://github.com/tvondra/sequential-uuids)

Generates UUIDs with sequential patterns to reduce random I/O in indexes while maintaining sufficient randomness to avoid collisions.

```sql
CREATE EXTENSION sequential_uuids;
```

### Functions

| Function | Description |
|---|---|
| `uuid_sequence_nextval(sequence regclass, block_size int DEFAULT 65536, block_count int DEFAULT 65536)` | Generate a sequential UUID based on a sequence |
| `uuid_time_nextval(interval_length int DEFAULT 60, interval_count int DEFAULT 65536)` | Generate a sequential UUID based on current timestamp |

### Examples

```sql
CREATE SEQUENCE my_seq;

-- Sequence-based UUID generation
SELECT uuid_sequence_nextval('my_seq'::regclass);

-- Time-based UUID generation (wraps around every ~45 days with defaults)
SELECT uuid_time_nextval();

-- Use as default for a column
CREATE TABLE orders (
  id uuid DEFAULT uuid_time_nextval() PRIMARY KEY,
  data text
);

-- Custom block size and count
SELECT uuid_sequence_nextval('my_seq', 256, 65536);
SELECT uuid_time_nextval(120, 32768);
```
