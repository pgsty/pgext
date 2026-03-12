


## Usage

> [pgsodium: libsodium-based cryptographic functions for PostgreSQL](https://github.com/michelp/pgsodium)

`pgsodium` is an encryption library extension for PostgreSQL using the [libsodium](https://download.libsodium.org/doc/) library. It provides a direct SQL interface to libsodium, server-managed key derivation, and Transparent Column Encryption (TCE).

```sql
CREATE EXTENSION pgsodium;
```

### Generating Random Data

```sql
SELECT pgsodium.randombytes_random();
SELECT pgsodium.randombytes_buf(16);         -- 16 random bytes
SELECT pgsodium.randombytes_uniform(100);    -- random int 0-99
```

### Secret Key Encryption (Authenticated)

```sql
SELECT * FROM pgsodium.crypto_secretbox_keygen();
SELECT pgsodium.crypto_secretbox('message', nonce, key);
SELECT pgsodium.crypto_secretbox_open(ciphertext, nonce, key);
```

### Public Key Encryption

```sql
SELECT * FROM pgsodium.crypto_box_new_keypair();
SELECT pgsodium.crypto_box('message', nonce, public_key, secret_key);
SELECT pgsodium.crypto_box_open(ciphertext, nonce, public_key, secret_key);
```

### Public Key Signatures

```sql
SELECT * FROM pgsodium.crypto_sign_new_keypair();
SELECT pgsodium.crypto_sign('message', secret_key);
SELECT pgsodium.crypto_sign_open(signed_message, public_key);
```

### Password Hashing

```sql
SELECT pgsodium.crypto_pwhash_str('my_password');
SELECT pgsodium.crypto_pwhash_str_verify(hash, 'my_password');
```

### Hashing

```sql
SELECT pgsodium.crypto_generichash('data');
SELECT pgsodium.crypto_shorthash('data', key);
```

### Server Key Management

pgsodium can load an external root key into memory that is never accessible to SQL. Sub-keys are derived by key id:

```sql
SELECT * FROM pgsodium.create_key();
-- Returns a UUID key id for use with TCE or encryption functions
```

### Transparent Column Encryption (TCE)

```sql
CREATE TABLE private.users (
    id bigserial PRIMARY KEY,
    secret text
);

SECURITY LABEL FOR pgsodium ON COLUMN private.users.secret
  IS 'ENCRYPT WITH KEY ID dfc44293-fa78-4a1a-9ef9-7e600e63e101';
```

Encrypted data is stored on disk and automatically decrypted via a generated view.

### Security Roles

- `pgsodium_keyiduser` -- less privileged, can only access keys by UUID
- `pgsodium_keymaker` -- more privileged, can work with raw keys
