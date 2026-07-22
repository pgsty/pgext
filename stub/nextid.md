## Usage

Sources:

- [Official README at the catalog revision](https://github.com/wlltmrt/pg-nextid/blob/d0a32e4df366aa118a4f2cfc55df3bfaf570af59/README.md)
- [Implementation at the catalog revision](https://github.com/wlltmrt/pg-nextid/blob/d0a32e4df366aa118a4f2cfc55df3bfaf570af59/nextid.c)

`nextid` 0.1 generates a 64-bit time-and-shard identifier from a sequence value, a shard number, and the current wall clock. It can provide locally generated sortable-looking IDs, but uniqueness depends on caller-managed rate, shard, and clock invariants.

### Core Workflow

```sql
CREATE EXTENSION nextid;
CREATE SEQUENCE order_id_seq;

-- Use a globally assigned shard number in the range 0..8191.
SELECT c_next_id(nextval('order_id_seq'), 42);

CREATE TABLE orders (
    id bigint PRIMARY KEY DEFAULT c_next_id(nextval('order_id_seq'), 42),
    payload jsonb NOT NULL
);
```

### Function and Layout

- `c_next_id(bigint, integer)` combines milliseconds since the implementation's fixed epoch with the shard shifted by 10 bits and the sequence value modulo 1024.
- The low 10 bits repeat every 1024 sequence values; the remaining low-bit space is intended for a 13-bit shard identifier.
- The function is strict and returns bigint. It performs no validation or masking of the shard argument.

### Uniqueness Boundaries

- Allocate shard IDs centrally and keep them between 0 and 8191. Negative or wider values can overlap the timestamp portion and corrupt the layout.
- Do not generate more than 1024 identifiers in the same millisecond on one shard. The sequence remainder then repeats.
- A wall-clock rollback can produce duplicates or IDs that do not sort by creation order. Use clock monitoring and database uniqueness constraints; do not treat the bit layout alone as a guarantee.
- Changing the epoch, shard allocation, or algorithm requires a compatibility plan for previously stored IDs.
