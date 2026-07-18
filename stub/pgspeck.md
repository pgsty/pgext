## Usage

Sources:

- [Upstream README](https://github.com/johto/pgspeck/blob/6120dffe2d293ed1a3e201ca299fba348a7b381d/README.md)
- [Extension control file](https://github.com/johto/pgspeck/blob/6120dffe2d293ed1a3e201ca299fba348a7b381d/pgspeck.control)
- [Version 1.0 SQL interface](https://github.com/johto/pgspeck/blob/6120dffe2d293ed1a3e201ca299fba348a7b381d/pgspeck--1.0.sql)
- [PGXN package metadata](https://github.com/johto/pgspeck/blob/6120dffe2d293ed1a3e201ca299fba348a7b381d/META.json)

`pgspeck` version `1.0` provides deterministic Speck block-cipher permutations for turning sequential integer values into less obvious integer sequences. The 32-bit functions use the low 32 bits of the plaintext and a 64-bit key; the 48-bit functions use the low 48 bits of the plaintext and two nonzero 48-bit key components.

### Verify a reversible mapping

```sql
CREATE EXTENSION pgspeck;

SELECT pgspeck_encrypt32(42, 123456789);

SELECT pgspeck_decrypt32(
    pgspeck_encrypt32(42, 123456789),
    123456789
);
```

Use this only as reversible identifier obfuscation. Small deterministic blocks reveal repeated inputs, provide no nonce or authentication, and do not protect integrity; wrapping or truncation also requires careful domain analysis. Keys passed in SQL can leak through function definitions, logs, monitoring, query text, or role access. Do not use `pgspeck` for passwords, personal data, secrets, access tokens, or modern application encryption. Keep authorization independent of the obscured identifier and use a reviewed authenticated-encryption design whenever confidentiality matters.
