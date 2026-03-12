


## Usage

> [omni_seq: Distributed integer sequences](https://docs.omnigres.org/omni_seq/id/)

The `omni_seq` extension provides distributed (prefixed) identifier types to work around PostgreSQL's limitation where sequence data is not replicated in logical replication.

### Types

Types follow the pattern `omni_seq.id_<PREFIX_TYPE>_<ID_TYPE>`. Supported base types: `int16`, `int32`, `int64`. When prefix and identifier types match, the shorthand `omni_seq.id_int64` is available.

### Create a Table with Distributed IDs

```sql
CREATE SEQUENCE seq;
CREATE TABLE t (
    id omni_seq.id_int64 PRIMARY KEY NOT NULL
       DEFAULT omni_seq.id_int64_nextval(omni_seq.system_identifier(), 'seq')
);
```

IDs display as `PREFIX:IDENTIFIER` (e.g., `7222168279780171472:1`).

### Functions

- **`omni_seq.id_int64_nextval(node_id, sequence_name)`** -- Generates the next distributed ID
- **`omni_seq.system_identifier()`** -- Returns the unique 64-bit system identifier from pg_control
- **`omni_seq.id_int64_int32_make(prefix, identifier)`** -- Constructs a distributed ID from components (useful during migrations)

### Migration Example

```sql
BEGIN;
LOCK TABLE my_table;
ALTER TABLE my_table ALTER COLUMN id DROP IDENTITY IF EXISTS;
CREATE SEQUENCE my_table_id_seq;
ALTER TABLE my_table
    ALTER COLUMN id TYPE omni_seq.id_int64_int32
        USING omni_seq.id_int64_int32_make(0, id),
    ALTER COLUMN id SET DEFAULT omni_seq.id_int64_int32_nextval(
        omni_seq.system_identifier(), 'my_table_id_seq');
COMMIT;
```
