

## Usage

> [pg_base58: base58 encoding/decoding for PostgreSQL](https://github.com/Fell-x27/pg_base58)

Provides functions to encode and decode data using Base58 encoding.

```sql
CREATE EXTENSION pg_base58;
```

### Functions

| Function | Description |
|---|---|
| `base58_encode(bytea)` | Encode bytea data to Base58 text |
| `base58_decode(text)` | Decode Base58 text back to bytea |

### Examples

```sql
-- Encode a string to Base58
SELECT base58_encode('hello'::bytea);
-- 'Cn8eVZg'

-- Decode a Base58 string back
SELECT base58_decode('Cn8eVZg');
-- '\x68656c6c6f'  (i.e., 'hello')

-- Round-trip
SELECT convert_from(base58_decode(base58_encode('hello'::bytea)), 'UTF8');
-- 'hello'
```
