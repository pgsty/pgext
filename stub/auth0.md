## Usage

Sources:

- [Official extension control file](https://github.com/brunoenten/pg_auth0/blob/c6831e8eb821384fd5ef9950e31bfe35450913e7/auth0.control)
- [Official upstream documentation](https://github.com/brunoenten/pg_auth0/blob/c6831e8eb821384fd5ef9950e31bfe35450913e7/README.md)

`auth0` — Auth0 Management API client for PostgreSQL.

The reviewed catalog snapshot records version `0.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `http`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "auth0";
SELECT extversion
FROM pg_extension
WHERE extname = 'auth0';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
