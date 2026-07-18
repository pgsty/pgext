## Usage

Sources:

- [Official extension control file](https://github.com/beargiles/pg-bignum/blob/79950064ca96595dd0a26c81915ea2d7396e6986/bignum.control)
- [Official upstream documentation](https://pgxn.org/dist/bignum/doc/bignum.html)
- [Official PGXN distribution page](https://pgxn.org/dist/bignum/)

`bignum` — Unlimited precision integers for PostgreSQL

The reviewed catalog snapshot records version `0.8.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "bignum";
SELECT extversion
FROM pg_extension
WHERE extname = 'bignum';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
