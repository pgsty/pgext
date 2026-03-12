


## Usage

> [pgsmcrypto: SM national cryptographic algorithm extension for PostgreSQL](https://github.com/zhuobie/pgsmcrypto)

`pgsmcrypto` provides Chinese national cryptographic (SM series) algorithms for PostgreSQL, including SM3 hashing, SM2 asymmetric encryption/signing, and SM4 symmetric encryption.

```sql
CREATE EXTENSION pgsmcrypto;
```

### SM3 Message Digest

```sql
SELECT sm3_hash_string('abc');              -- Returns 64-char hex string (32 bytes)
SELECT sm3_hash('abc'::bytea);              -- Hash bytea input
SELECT sm3_hash(E'\\x616263');              -- Hash raw hex input
```

### SM2 Asymmetric Encryption

#### Key Generation

```sql
SELECT sm2_gen_keypair();                   -- Returns {private_key, public_key} array
SELECT sm2_privkey_valid('f774...');        -- Validate private key (1=valid)
SELECT sm2_pubkey_valid('8093...');         -- Validate public key (1=valid)
SELECT sm2_pk_from_sk('f774...');           -- Derive public key from private key
```

#### Key Export/Import (PEM)

```sql
SELECT sm2_keypair_to_pem_bytes('f774...');       -- Private key to PEM
SELECT sm2_pubkey_to_pem_bytes('8093...');        -- Public key to PEM
SELECT sm2_keypair_from_pem_bytes(pem_bytes);     -- Import from PEM
SELECT sm2_pubkey_from_pem_bytes(pem_bytes);      -- Import public key from PEM
```

#### Sign and Verify

```sql
-- Raw sign/verify (signs message directly)
WITH s AS (
    SELECT sm2_sign_raw('abc'::bytea, 'f774...') AS sig
)
SELECT sm2_verify_raw('abc'::bytea, sig, '8093...') FROM s;

-- Standard sign/verify (SM2 specification with id + SM3 digest)
WITH s AS (
    SELECT sm2_sign('myid'::bytea, 'abc'::bytea, 'f774...') AS sig
)
SELECT sm2_verify('myid'::bytea, 'abc'::bytea, sig, '8093...') FROM s;
```

#### Encrypt and Decrypt

```sql
-- Standard encrypt/decrypt
WITH c AS (
    SELECT sm2_encrypt('abc'::bytea, '8093...') AS enc
)
SELECT sm2_decrypt(enc, 'f774...') FROM c;

-- Also available: sm2_encrypt_c1c2c3, sm2_encrypt_asna1, sm2_encrypt_hex, sm2_encrypt_base64
-- with corresponding decrypt variants
```

### SM4 Symmetric Encryption

```sql
-- ECB mode (key must be 16 bytes)
SELECT sm4_encrypt_ecb('abc'::bytea, '1234567812345678'::bytea);
SELECT sm4_decrypt_ecb(encrypted, '1234567812345678'::bytea);

-- CBC mode (key and IV must be 16 bytes)
SELECT sm4_encrypt_cbc('abc'::bytea, '1234567812345678'::bytea, '0000000000000000'::bytea);
SELECT sm4_decrypt_cbc(encrypted, '1234567812345678'::bytea, '0000000000000000'::bytea);
```
