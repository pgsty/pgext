## Usage

Sources:

- [Official extension control file](https://github.com/MasaoFujii/postgres_security/blob/8a5d11461b5381b4cd3a7cae70d1d97bb73f2711/postgres_security.control)

`postgres_security` — Define a schema-local text type that duplicates PostgreSQL text storage and I/O behavior

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "postgres_security";
SELECT extversion
FROM pg_extension
WHERE extname = 'postgres_security';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
