## Usage

Sources:

- [Official extension SQL](https://api.pgxn.org/src/openssl/openssl-0.0.1/sql/openssl.sql)
- [Official C implementation](https://api.pgxn.org/src/openssl/openssl-0.0.1/src/pg_openssl.c)
- [Official extension control file](https://api.pgxn.org/src/openssl/openssl-0.0.1/openssl.control)

`openssl` version `0.0.1` is an early PGXN experiment that attempts to list algorithms exposed by the server's linked OpenSSL library. Despite its name, it provides no SQL encryption, decryption, hashing, signing, or key-management operation.

### Exposed Objects

The extension declares `cipher_info` and `digest_info` composite types and two set-returning functions:

```sql
CREATE EXTENSION openssl;

SELECT * FROM list_ciphers();
SELECT * FROM list_digests();
```

`cipher_info` is intended to report fields such as algorithm name, block size, key length, IV length, mode, and OpenSSL flags. `digest_info` is intended to report name, block size, digest length, public-key type, flags, and numeric identifier.

### Compatibility and Safety Boundary

The reviewed C implementation uses legacy OpenSSL enumeration APIs and fixed storage for at most 100 algorithms. More seriously, `list_digests()` forces a single output row while the code that formats that row is disabled, leaving its result construction incomplete. Treat digest output as unreliable and do not depend on this extension for inventory or compliance decisions.

The source also contains an unexposed key-type listing function and tests whose function names do not match the installed SQL interface. Use `openssl` only in a disposable compatibility lab after verifying the exact PostgreSQL and OpenSSL build. For production cryptography, use maintained PostgreSQL facilities or application libraries with documented algorithms and key handling.
