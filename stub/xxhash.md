


## Usage

> [xxhash: xxHash hash functions for PostgreSQL](https://github.com/hatarist/pg_xxhash)

Provides xxHash non-cryptographic hash functions that return hex strings or bytea.

### Text Output Functions

```sql
SELECT xxh32('https://example.com');     -- 'ba15a4a8'
SELECT xxh64('https://example.com');     -- 'b131752760b48654'
SELECT xxh3_64('https://example.com');   -- '9398cc7c078760e6'
SELECT xxh128('https://example.com');    -- '4879d6aa9d88e9c7a169c008892d4829'
```

### Binary Output Functions

```sql
SELECT xxh32b('https://example.com');    -- \xba15a4a8
SELECT xxh64b('https://example.com');    -- \xb131752760b48654
SELECT xxh3_64b('https://example.com');  -- \x9398cc7c078760e6
SELECT xxh128b('https://example.com');   -- \x4879d6aa9d88e9c7a169c008892d4829
```

### Available Functions

| Function | Bits | Output |
|----------|------|--------|
| `xxh32(text)` | 32 | hex text |
| `xxh64(text)` | 64 | hex text |
| `xxh3_64(text)` | 64 | hex text |
| `xxh128(text)` | 128 | hex text |
| `xxh32b(text)` | 32 | bytea |
| `xxh64b(text)` | 64 | bytea |
| `xxh3_64b(text)` | 64 | bytea |
| `xxh128b(text)` | 128 | bytea |
