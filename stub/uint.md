

## Usage

> [uint: unsigned integer types for PostgreSQL](https://github.com/petere/pguint)

The `uint` extension adds unsigned and small integer types to PostgreSQL.

```sql
CREATE EXTENSION uint;
```

### Data Types

| Type | Description | Range |
|------|-------------|-------|
| `int1` | Signed 8-bit integer | -128 to 127 |
| `uint1` | Unsigned 8-bit integer | 0 to 255 |
| `uint2` | Unsigned 16-bit integer | 0 to 65535 |
| `uint4` | Unsigned 32-bit integer | 0 to 4294967295 |
| `uint8` | Unsigned 64-bit integer | 0 to 18446744073709551615 |

### Usage

```sql
CREATE TABLE foo (
    a uint4,
    b text
);

SELECT * FROM foo WHERE a > 4;
SELECT avg(a) FROM foo;
```

### Operators and Functions

All types include a full set of arithmetic operators (`+`, `-`, `*`, `/`, `%`), comparison operators (`=`, `<>`, `<`, `>`, `<=`, `>=`), and operators for each combination of types. Standard aggregate functions and index support (btree, hash) are also provided.
