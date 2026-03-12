


## Usage

> [pg_enigma: Encrypted data type for PostgreSQL using PGP and RSA keys](https://github.com/SoftwareLibreMx/pg_enigma)

`pg_enigma` provides an `Enigma` encrypted data type for PostgreSQL that encrypts data at rest using PGP or OpenSSL RSA keys. Data is stored encrypted and only decrypted when the private key is loaded into memory.

```sql
CREATE EXTENSION IF NOT EXISTS pg_enigma;
```

### PGP Key Encryption

```sql
-- Create a table with an encrypted column (key slot 2)
CREATE TABLE test_pgp (
    id SERIAL,
    val Enigma(2)
);

-- Load the public key for encryption
SELECT set_public_key_from_file(2, '/path/to/public-key.asc');

-- Insert data (automatically encrypted with the public key)
INSERT INTO test_pgp (val) VALUES ('A secret value'::Text);

-- Without private key, SELECT returns encrypted PGP message
SELECT * FROM test_pgp;

-- Load private key to enable decryption
SELECT set_private_key_from_file(2, '/path/to/private-key.asc', 'passphrase');

-- Now SELECT returns decrypted plaintext
SELECT * FROM test_pgp;
-- id |      val
-- ----+----------------
--   1 | A secret value

-- Remove private key from memory
SELECT forget_private_key(2);
-- Subsequent SELECTs return encrypted data again
```

### RSA Key Encryption

```sql
CREATE TABLE test_rsa (
    id SERIAL,
    val Enigma(3)
);

SELECT set_public_key_from_file(3, '/path/to/alice_public.pem');
INSERT INTO test_rsa (val) VALUES ('Another secret value'::Text);

SELECT set_private_key_from_file(3, '/path/to/alice_private.pem', 'passphrase');
SELECT * FROM test_rsa;

SELECT forget_private_key(3);
```

### Functions

| Function | Description |
|----------|-------------|
| `set_public_key_from_file(slot, path)` | Load a public key for encryption |
| `set_private_key_from_file(slot, path, passphrase)` | Load a private key for decryption |
| `forget_private_key(slot)` | Remove private key from memory |
