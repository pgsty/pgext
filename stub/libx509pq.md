## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/crtsh/libx509pq/blob/bed05c6e87d79a98d79e42a1fac16386eead0c13/README.md)
- [Version 1.3 SQL objects](https://github.com/crtsh/libx509pq/blob/bed05c6e87d79a98d79e42a1fac16386eead0c13/libx509pq--1.3.sql)
- [OpenSSL-backed implementation](https://github.com/crtsh/libx509pq/blob/bed05c6e87d79a98d79e42a1fac16386eead0c13/libx509pq.c)

`libx509pq` exposes OpenSSL X.509 certificate parsing to SQL. It accepts DER certificates as `bytea` and extracts subject, issuer, validity, serial number, key and signature information, extensions, alternate names, fingerprints, and related fields.

```sql
CREATE EXTENSION libx509pq;
SELECT x509_commonName(der),
       x509_issuerName(der),
       x509_notBefore(der),
       x509_notAfter(der),
       x509_keyAlgorithm(der),
       x509_keySize(der)
FROM certificate_store;
```

For multiple basic fields, `x509_basic_info(der)` avoids parsing the same certificate repeatedly. Other APIs enumerate extensions and names, expose public keys, identify selected weak-key patterns, and verify one certificate signature against a supplied key.

Parsing or signature verification is not full PKIX path validation: it does not by itself establish trust, hostname validity, revocation status, policy, or current acceptability. Malformed DER reaches native OpenSSL code; cap input size, fuzz and regression-test untrusted corpora, and keep PostgreSQL and OpenSSL patched. Pin the OpenSSL ABI used at build and runtime, review sentinel-versus-null error behavior, restrict expensive functions, and benchmark bulk parsing.
