

## Usage

> [uint128: unsigned 128-bit integer type for PostgreSQL](https://github.com/pg-uint/pg-uint128)

The `uint128` extension provides comprehensive unsigned and 128-bit integer types with full operator and index support.

```sql
CREATE EXTENSION uint128;
```

### Data Types

| Type | Size | Range |
|------|------|-------|
| `uint1` | 8-bit | 0 to 255 |
| `uint2` | 16-bit | 0 to 65535 |
| `uint4` | 32-bit | 0 to 4294967295 |
| `uint8` | 64-bit | 0 to 18446744073709551615 |
| `uint16` | 128-bit | 0 to 340282366920938463463374607431768211455 |
| `int1` | 8-bit | -128 to 127 |
| `int16` | 128-bit | -170141183460469231731687303715884105728 to 170141183460469231731687303715884105727 |

### Operators

- **Arithmetic**: `+`, `-`, `*`, `/`, `%`
- **Bitwise**: `#` (XOR), `&` (AND), `|` (OR), `~` (NOT), `<<` (left shift), `>>` (right shift)
- **Comparison**: `=`, `<>`, `>`, `<`, `>=`, `<=`

Mixed-type arithmetic between signed and unsigned types is supported.

### Features

- Range types for all integer types (`uint1range`, `uint16range`, etc.) with GiST indexing
- Casts to/from `numeric`, `real`, `double`, `uuid` (uint16 only), `json`, `jsonb`
- Aggregate functions: `SUM`, `AVG`, `MIN`, `MAX`
- `generate_series()` support for all types
- Btree and hash index support
- Binary send/receive protocol support
