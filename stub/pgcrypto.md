


## Usage

> [pgcrypto: Cryptographic functions for PostgreSQL](https://www.postgresql.org/docs/current/pgcrypto.html)

`pgcrypto` provides cryptographic functions including hashing, password hashing, PGP encryption, and raw encryption.

```sql
CREATE EXTENSION pgcrypto;
```

### General Hashing

```sql
SELECT digest('data', 'sha256');                    -- Binary hash
SELECT encode(digest('data', 'sha256'), 'hex');     -- Hex output
SELECT hmac('data', 'secret_key', 'sha256');        -- HMAC
```

Supported algorithms: `md5`, `sha1`, `sha224`, `sha256`, `sha384`, `sha512`.

### Password Hashing

```sql
-- Set a new password
UPDATE users SET pswhash = crypt('new password', gen_salt('bf'));

-- Verify a password
SELECT (pswhash = crypt('entered password', pswhash)) AS valid FROM users;
```

`gen_salt()` types: `bf` (Blowfish), `md5`, `xdes`, `des`, `sha256crypt`, `sha512crypt`.

### PGP Symmetric Encryption

```sql
-- Encrypt
SELECT pgp_sym_encrypt('secret data', 'password');
SELECT pgp_sym_encrypt('secret data', 'password', 'cipher-algo=aes256, compress-algo=1');

-- Decrypt
SELECT pgp_sym_decrypt(encrypted_data, 'password');
```

### PGP Public Key Encryption

```sql
-- Encrypt with public key
SELECT pgp_pub_encrypt('secret data', dearmor(pubkey));

-- Decrypt with private key
SELECT pgp_pub_decrypt(encrypted_data, dearmor(seckey));
SELECT pgp_pub_decrypt(encrypted_data, dearmor(seckey), 'key_password');
```

### Key Utilities

```sql
SELECT pgp_key_id(dearmor(key_text));       -- Extract key ID
SELECT armor(binary_data);                   -- ASCII armor
SELECT dearmor(armored_text);                -- Remove armor
SELECT pgp_armor_headers(armored_text);      -- Extract armor headers
```

### Raw Encryption

```sql
SELECT encrypt('data'::bytea, 'key'::bytea, 'aes');
SELECT decrypt(encrypted, 'key'::bytea, 'aes');

-- With IV
SELECT encrypt_iv('data'::bytea, 'key'::bytea, 'iv'::bytea, 'aes-cbc/pad:pkcs');
SELECT decrypt_iv(encrypted, 'key'::bytea, 'iv'::bytea, 'aes-cbc/pad:pkcs');
```

Algorithms: `bf` (Blowfish), `aes`. Modes: `cbc` (default), `cfb`, `ecb`. Padding: `pkcs` (default), `none`.

### Random Data

```sql
SELECT gen_random_bytes(16);        -- 16 cryptographic random bytes
SELECT gen_random_uuid();           -- Random UUID v4
```

### PGP Encryption Options

| Option | Values | Default |
|--------|--------|---------|
| `cipher-algo` | `bf`, `aes128`, `aes192`, `aes256`, `3des`, `cast5` | `aes128` |
| `compress-algo` | `0` (none), `1` (ZIP), `2` (ZLIB) | `0` |
| `compress-level` | `0-9` | `6` |
| `s2k-mode` | `0`, `1`, `3` | `3` |
