

## Usage

> [pg_hashids: generate short unique ids from integers](https://github.com/iCyberon/pg_hashids)

Converts numbers like 347 into strings like "yr8". Useful for obfuscating database primary keys.

```sql
CREATE EXTENSION pg_hashids;
```

### Functions

| Function | Description |
|---|---|
| `id_encode(number [, salt [, min_length [, alphabet]]])` | Encode an integer to a hashid string |
| `id_decode(hash, salt, min_length [, alphabet])` | Decode a hashid string back to integer array |
| `id_decode_once(hash [, salt [, min_length [, alphabet]]])` | Decode a hashid string back to a single integer |

### Examples

```sql
-- Basic encoding
SELECT id_encode(1001);                                    -- 'jNl'
SELECT id_encode(1234567, 'This is my salt');              -- 'Pdzxp'
SELECT id_encode(1234567, 'This is my salt', 10);          -- 'PlRPdzxpR7'

-- Custom alphabet
SELECT id_encode(1234567, 'This is my salt', 10, 'abcdefghijABCDxFGHIJ1234567890');
-- '3GJ956J9B9'

-- Decoding (use the same salt!)
SELECT id_decode('PlRPdzxpR7', 'This is my salt', 10);     -- 1234567
SELECT id_decode_once('jNl');                               -- 1001
SELECT id_decode_once('Pdzxp', 'This is my salt');          -- 1234567
```
