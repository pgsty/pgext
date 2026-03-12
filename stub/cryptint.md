


## Usage

> [cryptint: Encrypt and decrypt integers with SKIP32 and XTEA](https://github.com/dverite/cryptint)

### SKIP32 (32-bit integer encryption)

Block cipher with 24 rounds based on Skipjack. Encrypts 32-bit values with an 80-bit (10-byte) key.

```sql
SELECT skip32_encrypt(42, '\x00010203040506070809'::bytea);
-- encrypted int4

SELECT skip32_decrypt(175586429, '\x00010203040506070809'::bytea);
-- 0
```

Roundtrip example:

```sql
SELECT i, skip32_encrypt(i, '\x00010203040506070809'::bytea) AS encrypted,
       skip32_decrypt(skip32_encrypt(i, '\x00010203040506070809'::bytea), '\x00010203040506070809'::bytea) AS decrypted
FROM generate_series(-3, 3) AS i;
```

### XTEA (64-bit integer encryption)

Block cipher with 64 rounds. Encrypts 64-bit values with a 128-bit (16-byte) key.

```sql
SELECT xtea_encrypt(42::bigint, '\x000102030405060708090a0b0c0d0e0f'::bytea);
-- encrypted bigint

SELECT xtea_decrypt(103200416458222088::bigint, '\x000102030405060708090a0b0c0d0e0f'::bytea);
-- 1
```

### Functions Reference

| Function | Description |
|----------|-------------|
| `skip32_encrypt(int, bytea)` | Encrypt int4 with 10-byte key |
| `skip32_decrypt(int, bytea)` | Decrypt int4 with 10-byte key |
| `xtea_encrypt(bigint, bytea)` | Encrypt bigint with 16-byte key |
| `xtea_decrypt(bigint, bytea)` | Decrypt bigint with 16-byte key |
