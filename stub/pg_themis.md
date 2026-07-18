## Usage

Sources:

- [Pinned upstream README](https://github.com/cossacklabs/pg_themis/blob/609525bdd564d90c0067398fa0debe99882b12e8/README.md)
- [Version 1.0 installation SQL](https://github.com/cossacklabs/pg_themis/blob/609525bdd564d90c0067398fa0debe99882b12e8/pg_themis--1.0.sql)
- [Pinned C implementation](https://github.com/cossacklabs/pg_themis/blob/609525bdd564d90c0067398fa0debe99882b12e8/pg_themis.c)
- [Pinned extension control file](https://github.com/cossacklabs/pg_themis/blob/609525bdd564d90c0067398fa0debe99882b12e8/pg_themis.control)
- [Current Themis product documentation](https://www.cossacklabs.com/themis/)

pg_themis version 1.0 is a 2016 PostgreSQL binding for the external Themis cryptographic library. It exposes Secure Cell Seal encryption and decryption with a symmetric key, plus Secure Message encryption with a recipient public key and decryption with the corresponding private key.

### Secure Cell round trip

Install a compatible system Themis library first. This isolated example uses pgcrypto only to create a temporary key:

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_themis;

WITH key_material AS MATERIALIZED (
    SELECT gen_random_bytes(32) AS key
)
SELECT convert_from(
    pg_themis_scell_decrypt_seal(
        pg_themis_scell_encrypt_seal(
            convert_to('payload', 'UTF8'),
            key
        ),
        key
    ),
    'UTF8'
)
FROM key_material;
```

The extension does not generate, store, rotate, authorize, or audit application keys. Those controls remain the caller's responsibility.

### Unsafe legacy Secure Message path

The Secure Message helper narrows plaintext or ciphertext length to an unsigned one-byte value before calling Themis. Inputs longer than 255 bytes are therefore truncated modulo 256. Its decrypt path also reads an embedded 32-bit public-key length before first proving the ciphertext is at least four bytes, then trusts that length for pointer arithmetic. Malformed ciphertext can reach out-of-bounds reads. Do not use the Secure Message SQL functions from this snapshot.

All four functions receive plaintext or keys as SQL bytea arguments. Query text, parameters, activity views, logs, tracing, errors, backups, and memory can expose secrets. The installation SQL leaves default PUBLIC execution in place. Revoke it, keep key material outside SQL literals, and use a reviewed key-management design.

The upstream PostgreSQL binding has not changed since 2016, and its broad PostgreSQL 9.1+ statement is not a current compatibility guarantee. Current Themis itself may evolve independently of this binding. Pin and test the exact library ABI, use authenticated test vectors, fuzz malformed and boundary-length inputs, and prefer a maintained integration for production cryptography.
