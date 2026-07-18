## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/cert/)
- [Official upstream README](https://github.com/beargiles/pg-cert/blob/555b021861e0c8a52b8fcfb08e732945e56f5d82/README.md)

`cert` — PostgreSQL data types and functions for X.509 certificates, certificate requests, and related OpenSSL objects.

The reviewed catalog snapshot records version `0.1.1`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `bignum`.

```sql
CREATE EXTENSION "cert";
SELECT extversion
FROM pg_extension
WHERE extname = 'cert';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
