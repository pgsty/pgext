## Usage

Sources:

- [README](https://github.com/RustedBytes/pg-crypto/blob/a07210ac1f98e5a722843c7389afdd0cd26ebe3c/README.md)
- [Control file](https://github.com/RustedBytes/pg-crypto/blob/a07210ac1f98e5a722843c7389afdd0cd26ebe3c/pg_crypto.control)
- [Cargo.toml](https://github.com/RustedBytes/pg-crypto/blob/a07210ac1f98e5a722843c7389afdd0cd26ebe3c/Cargo.toml)
- [src/lib.rs](https://github.com/RustedBytes/pg-crypto/blob/a07210ac1f98e5a722843c7389afdd0cd26ebe3c/src/lib.rs)
- [docs/SECURITY.md](https://github.com/RustedBytes/pg-crypto/blob/a07210ac1f98e5a722843c7389afdd0cd26ebe3c/docs/SECURITY.md)

`pg_crypto` implements familiar pgcrypto-style functions and additional cryptographic primitives in Rust on PostgreSQL 14–18. It is a separate project from the built-in `pgcrypto` extension.

### Core Workflow

```sql
CREATE EXTENSION pg_crypto;
SELECT encode(digest('hello', 'sha256'), 'hex');
WITH key AS (SELECT gen_random_bytes(32) AS value)
SELECT xchacha20poly1305_decrypt(
  xchacha20poly1305_encrypt('example'::bytea, value, 'context'::bytea),
  value, 'context'::bytea
) FROM key;
```

### API

`digest` and `hmac` compute hashes and message authentication codes. `crypt` and `gen_salt` support password-hash compatibility. Symmetric and public-key OpenPGP functions include `pgp_sym_encrypt`, `pgp_sym_decrypt`, and armor helpers. Modern functions include `xchacha20poly1305_encrypt`, `xchacha20poly1305_decrypt`, `argon2id_hash`, and `secretbox`.

Authenticated encryption verifies its authentication tag and associated data. Legacy raw encryption interfaces exist for compatibility, but require the caller to supply correct key sizes, IVs, and authentication design.

### Security and Installation

The extension is relocatable and needs no preload or OpenSSL runtime. Overlapping SQL names mean `pg_crypto` and `pgcrypto` cannot coexist in the same schema. Review the pinned compatibility and security documentation before migrating existing ciphertext or hashes.

Keys, plaintext, and passwords reach the database server; transport security, SQL logging, role access, and backups remain part of the design. Do not interpret API compatibility as evidence of an independent cryptographic audit.
