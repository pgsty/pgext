## Usage

Sources:

- [Official extension control file](https://api.pgxn.org/src/openssl/openssl-0.0.1/openssl.control)
- [Official upstream documentation](https://pgxn.org/dist/openssl/0.0.1/README.html)
- [Official PGXN distribution page](https://pgxn.org/dist/openssl/)

`openssl` — PostgreSQL extension for cryptographic functions and types provided by OpenSSL

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "openssl";
SELECT extversion
FROM pg_extension
WHERE extname = 'openssl';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
