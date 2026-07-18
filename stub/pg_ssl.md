## Usage

Sources:

- [Version 1.0 SQL declaration](https://github.com/RekGRpth/pg_ssl/blob/e0d6397ad7f034085e2bf3972c901c6f73e858ee/pg_ssl--1.0.sql)
- [OpenSSL C implementation](https://github.com/RekGRpth/pg_ssl/blob/e0d6397ad7f034085e2bf3972c901c6f73e858ee/pg_ssl.c)
- [Extension control file](https://github.com/RekGRpth/pg_ssl/blob/e0d6397ad7f034085e2bf3972c901c6f73e858ee/pg_ssl.control)
- [MIT license](https://github.com/RekGRpth/pg_ssl/blob/e0d6397ad7f034085e2bf3972c901c6f73e858ee/LICENSE)

`pg_ssl` exposes one C function, `sign(cert text, data text)`, which uses OpenSSL to create a PKCS#7 signature and returns base64 text. The first argument is parsed for both an X.509 certificate and its private key.

```sql
CREATE EXTENSION pg_ssl;
SELECT oid::regprocedure
FROM pg_proc
WHERE oid = 'sign(text,text)'::regprocedure;
```

Do not paste a production private key into SQL, query logs, monitoring samples, client history, or application parameters. PostgreSQL text values are not a hardware-backed key store, and the function has no key identifier, password callback, certificate-chain policy, or protected signing service boundary.

The upstream README is empty and there are no tests, releases, usage guide, or compatibility matrix. Treat version 1.0 as unreviewed native cryptographic glue. Audit memory handling and OpenSSL compatibility, revoke `EXECUTE` from `PUBLIC`, and prefer an external key-management signer for real credentials.
