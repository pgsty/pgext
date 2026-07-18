## Usage

Sources:

- [Official extension control file](https://api.pgxn.org/src/otp/otp-1.0.0/otp.control)
- [Official upstream documentation](https://pgxn.org/dist/otp/1.0.0/README.html)
- [Official PGXN distribution page](https://pgxn.org/dist/otp/)

`otp` — Implementation of TOTP compatible with Google Authenticator

The reviewed catalog snapshot records version `1.0.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plperlu`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "otp";
SELECT extversion
FROM pg_extension
WHERE extname = 'otp';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
