## Usage

Sources:

- [Official README](https://github.com/maciekgajewski/postgres-uints/blob/4bdc1d98729f77ab6dd32fa563a1c7b7c7d3559d/README.md)
- [Version 0.9 extension SQL](https://github.com/maciekgajewski/postgres-uints/blob/4bdc1d98729f77ab6dd32fa563a1c7b7c7d3559d/uints--0.9.sql)
- [Unsigned integer implementation](https://github.com/maciekgajewski/postgres-uints/blob/4bdc1d98729f77ab6dd32fa563a1c7b7c7d3559d/uint.c)

`uints` adds unsigned 16-bit and 32-bit integer types named `uint2` and `uint4`. Use them when the database must reject negative input while retaining fixed-width storage, arithmetic, comparisons, bit operations, and ordinary B-tree or hash indexes.

### Core Workflow

```sql
CREATE EXTENSION uints;

CREATE TABLE device_counters (
    device_id bigint PRIMARY KEY,
    small_counter uint2 NOT NULL,
    large_counter uint4 NOT NULL
);

INSERT INTO device_counters VALUES
    (1, 65535, 4294967295);

SELECT small_counter + 1::uint2,
       large_counter & 255::uint4,
       large_counter >> 8
FROM device_counters;

CREATE INDEX device_counters_large_idx
ON device_counters(large_counter);
```

`uint2` covers `0` through `65535`; `uint4` covers `0` through `4294967295`. Input outside the unsigned range and arithmetic overflow or underflow raise errors.

### Operators, Casts, and Indexes

Both types support comparison operators, `+`, `-`, `*`, `/`, `%`, bitwise `&`, `|`, `#`, `~`, and shifts `<<` and `>>`. Mixed `uint2` and `uint4` comparisons and arithmetic return the wider type where defined. `uint2` converts implicitly to `uint4`, while narrowing conversion is explicit. Default B-tree and hash operator classes make equality, ordering, grouping, and indexed lookup available.

### Caveats

The extension SQL defines only `uint2` and `uint4`; it does not install a `uint8` SQL type. Same-width casts between signed PostgreSQL types and these unsigned types are declared without conversion functions, so values with the high bit set can be interpreted by their raw bit pattern rather than as a mathematically equivalent signed value. Avoid relying on `uint4::integer` or negative `integer::uint4` conversions; convert through validated text or numeric logic when the full unsigned range is involved.

Division by zero and range overflow are errors. Version `0.9` is an old C implementation, so verify ABI and PostgreSQL-version compatibility, dump/restore behavior, client-driver decoding, and all boundary values before production use. No preload or restart is required.
