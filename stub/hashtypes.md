

## Usage

> [hashtypes: hash and checksum data types (SHA, CRC32)](https://github.com/adjust/hashtypes/)

The `hashtypes` extension provides native data types for storing hash values and checksums in their binary representation, saving storage compared to text.

```sql
CREATE EXTENSION hashtypes;
```

### Data Types

| Type | Storage | Description |
|------|---------|-------------|
| `sha1` | 20 bytes | SHA-1 hash (160-bit) |
| `sha224` | 28 bytes | SHA-224 hash (224-bit) |
| `sha256` | 32 bytes | SHA-256 hash (256-bit) |
| `sha384` | 48 bytes | SHA-384 hash (384-bit) |
| `sha512` | 64 bytes | SHA-512 hash (512-bit) |
| `crc32` | 4 bytes | CRC-32 checksum |

### Usage

```sql
CREATE TABLE objects (
    hash sha256 PRIMARY KEY,
    data bytea
);

INSERT INTO objects VALUES (
    'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855',
    '\x'
);

SELECT * FROM objects WHERE hash = 'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855';
```

### Operators

All types support comparison operators: `=`, `<>`, `<`, `>`, `<=`, `>=`.

### Index Support

Btree and hash index operator classes are provided for all types.

### Casts

All types support bidirectional casts to/from `text` and `bytea`.
