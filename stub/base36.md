

## Usage

> [base36: base36 encoding/decoding data types for PostgreSQL](https://github.com/adjust/pg-base36)

Provides data types for encoding and decoding values using the base36 scheme, with conversion between base36 and integer types.

```sql
CREATE EXTENSION base36;
```

### Types

| Type | Storage | Max String Length | Max Numeric Value |
|---|---|---|---|
| `base36` | 4 bytes (int) | 6 characters | 2,147,483,647 |
| `bigbase36` | 8 bytes (bigint) | 13 characters | 9,223,372,036,854,775,807 |

### Examples

```sql
-- Encode integer to base36
SELECT 1234567::base36;          -- 'qglj'

-- Decode base36 to integer
SELECT 'qglj'::base36::int;     -- 1234567

-- Bigbase36 for larger values
SELECT 9223372036854775807::bigbase36;           -- 'i1y004og0svr'
SELECT 'i1y004og0svr'::bigbase36::bigint;        -- 9223372036854775807
```
