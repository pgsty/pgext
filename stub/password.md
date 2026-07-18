## Usage

Sources:

- [Official extension control file](https://github.com/asio/password/blob/cf7efe3a92cde526072c0cf31158f20b99408e70/password.control)
- [Official upstream documentation](https://github.com/asio/password/blob/cf7efe3a92cde526072c0cf31158f20b99408e70/README)

`password` — PostgreSQL data type for passwords

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "password";
SELECT extversion
FROM pg_extension
WHERE extname = 'password';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
