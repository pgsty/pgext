

## Usage

> [base62: base62 encoding/decoding data types for PostgreSQL](https://github.com/adjust/pg-base62)

Provides data types for encoding and decoding values using the base62 scheme (0-9, A-Z, a-z).

```sql
CREATE EXTENSION base62;
```

### Types

| Type | Storage | Max String Length | Max Numeric Value |
|---|---|---|---|
| `base62` | 4 bytes (int) | 6 characters | 2,147,483,647 |
| `bigbase62` | 8 bytes (bigint) | 11 characters | 9,223,372,036,854,775,807 |
| `hugebase62` | 16 bytes | 20 characters | (bytea conversion) |

### Examples

```sql
-- Encode/decode base62
SELECT 2147483647::base62;          -- '2LKcb1'
SELECT '2LKcb1'::base62::int;      -- 2147483647

-- Bigbase62 for larger values
SELECT 9223372036854775807::bigbase62;           -- 'AzL8n0Y58m7'
SELECT 'AzL8n0Y58m7'::bigbase62::bigint;        -- 9223372036854775807

-- Hugebase62 with bytea conversion
SELECT 'AzL8n0Y58m7AzL8n0Y58'::hugebase62;
SELECT 'AzL8n0Y58m7AzL8n0Y58'::hugebase62::bytea;
SELECT '\x960c06065a6ed8ffff1e7149f40b1800'::bytea::hugebase62;

-- Note: base62 is case-sensitive
SELECT '2lkcb'::base62::int;   -- 40933305
SELECT '2LKCB'::base62::int;   -- 34635195
```
