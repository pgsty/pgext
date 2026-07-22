## Usage

Sources:

- [Upstream documentation](https://github.com/blueskylxb/utinyint/blob/8c5d9d832af6d148668d86769fa2b662b6582a19/doc/utinyint.md)
- [SQL type and casts](https://github.com/blueskylxb/utinyint/blob/8c5d9d832af6d148668d86769fa2b662b6582a19/utinyint.sql.in.c)
- [C type implementation](https://github.com/blueskylxb/utinyint/blob/8c5d9d832af6d148668d86769fa2b662b6582a19/utinyint.c)

`utinyint` version `0.1.1` defines a one-byte integer type intended to represent unsigned values. Its main benefit is compact scalar storage; PostgreSQL row alignment can erase that saving depending on neighboring columns.

### Core Workflow

```sql
CREATE EXTENSION utinyint;

CREATE TABLE sensor_value (
  sensor_id bigint,
  channel utinyint
);

INSERT INTO sensor_value VALUES (1, 0), (1, 128), (1, 255);
SELECT channel, channel::integer FROM sensor_value;
```

The reviewed install SQL defines binary/text I/O and casts between `utinyint` and `smallint`, `integer`, and `boolean`. Conversion to the wider integer types is implicit; conversion back is assignment-only. Boolean false maps to zero, true maps to one, and any nonzero value maps to true.

### Compatibility Caveats

The C header uses an unsigned byte and conversion functions accept values through `UCHAR_MAX`, but the upstream prose says the range is `-127` through `127`. The same prose claims arithmetic, aggregates, B-tree, and GIN support that does not appear in the reviewed installation SQL. Treat the source/install script as the API boundary and verify `0`, `128`, and `255` on the exact packaged build before use.

This custom base type has on-disk and binary-send formats tied to its C implementation. Test range errors, casts, dump/restore, replication, upgrades, and row-size effects. Prefer `smallint` when predictable built-in operators, indexing, portability, and maintenance matter more than one byte per scalar.
