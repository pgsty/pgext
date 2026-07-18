## Usage

Sources:

- [Upstream README](https://github.com/johto/pgocrypto/blob/87fa6cc60064ad0c90cd023b5727433e6eeaf41a/README.md)
- [Version 1.0 install SQL](https://github.com/johto/pgocrypto/blob/87fa6cc60064ad0c90cd023b5727433e6eeaf41a/pgocrypto--1.0.sql)
- [Go interoperability implementation](https://github.com/johto/pgocrypto/blob/87fa6cc60064ad0c90cd023b5727433e6eeaf41a/pgocrypto.go)

pgocrypto provides matching symmetric-encryption functions for PostgreSQL and Go. PostgreSQL functions encrypt bytea or text with AES-CBC and prefix a random 16-byte initialization vector; string ciphertext is Base64 encoded. The paired Go code uses the same wire format.

### Basic use

The extension calls pgcrypto functions without schema qualification and does not declare the dependency. Install both extensions in the same controlled schema and set a safe search path:

```sql
CREATE SCHEMA crypto;
CREATE EXTENSION pgcrypto WITH SCHEMA crypto;
CREATE EXTENSION pgocrypto WITH SCHEMA crypto;

SET search_path = crypto, pg_catalog;

WITH k AS (
    SELECT decode(repeat('01', 32), 'hex') AS key
)
SELECT pgo_decrypt_string(
           pgo_encrypt_string('hello', key),
           key
       )
FROM k;
```

The example key is deterministic test data. Production keys must come from a dedicated key-management design and must not be embedded in SQL, logs, schema objects, or application source.

### Caveats

- AES-CBC encryption here is not authenticated. An attacker can modify ciphertext without a built-in integrity failure, so the format is malleable and is not an AEAD envelope.
- The API accepts raw key bytes and supplies no password derivation, key identifier, rotation metadata, versioning, access control, or secure storage.
- Decryption errors and application behavior can create padding-oracle risks. Do not expose distinct cryptographic failures to untrusted callers.
- Prefer a maintained authenticated-encryption format for new security-sensitive systems. If interoperability requires this legacy format, add an independently verified authentication layer and a documented migration plan.
- Control the function search path because unqualified pgcrypto calls otherwise depend on session configuration. Restrict execute privileges and avoid logging plaintext or key parameters.
