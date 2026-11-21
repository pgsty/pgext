## Usage

Sources:

- [README.md](https://api.pgxn.org/src/uint128/uint128-1.3.0/README.md)
- [uint128.control](https://api.pgxn.org/src/uint128/uint128-1.3.0/uint128.control)
- [uint128--1.2.0--1.3.0.sql](https://api.pgxn.org/src/uint128/uint128-1.3.0/uint128--1.2.0--1.3.0.sql)
- [1.3.0](https://github.com/pg-uint/pg-uint128/releases/tag/1.3.0)


`uint128` adds unsigned integers and signed 8-bit and 128-bit integer types. Version 1.3.0 supports PostgreSQL 12–18 and fixes operator metadata that could make indexed comparisons return incorrect results.

### Store and Query Integers

Installation requires a superuser; the control file does not mark the extension trusted. It is relocatable and requires no preload.

```sql
CREATE EXTENSION uint128;
CREATE TABLE ids (id uint16 PRIMARY KEY);
INSERT INTO ids VALUES ('340282366920938463463374607431768211455'::uint16);
SELECT * FROM ids WHERE id > '18446744073709551615'::uint16;
SELECT * FROM generate_series(1::uint4, 3::uint4);
```

The unsigned types `uint1`, `uint2`, `uint4`, `uint8` and `uint16` occupy 1, 2, 4, 8 and 16 bytes. Their range is zero through 2 raised to the bit width minus one. `int1` and `int16` are signed 1-byte and 16-byte types.

### Operators and Indexes

The types provide arithmetic, comparisons, bitwise operations, B-tree and hash indexes, binary I/O, casts and `generate_series`. Corresponding range types, including `uint1range` and `uint16range`, support GiST indexes; multiranges depend on server support. `SUM` and `AVG` use `numeric` internally; `MIN` and `MAX` are also available.

Mixed signed/unsigned arithmetic is sensitive to operand order and result type. Cast explicitly when a negative result is possible: an unsigned result cannot represent it. Conversions and arithmetic may raise range errors; do not assume wraparound.

### Upgrade to 1.3.0

Install the matching files, then run as a superuser:

```sql
ALTER EXTENSION uint128 UPDATE TO '1.3.0';
```

This corrects comparison commutators and removes invalid mixed-type addition/subtraction commutators. It also fixes installation with the stricter operator checks in PostgreSQL 17.12. The supplied SQL updates operator metadata in place while preserving data, indexes, views and materialized views. Check the installed version and representative indexed comparisons after upgrading.
