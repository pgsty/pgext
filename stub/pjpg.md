## Usage

Sources:

- [Official extension control file](https://github.com/pjungwir/pjpg/blob/9437cb6b5efaa7bb7e67e625226c69962486a9a0/pjpg.control)
- [Official upstream documentation](https://github.com/pjungwir/pjpg/blob/9437cb6b5efaa7bb7e67e625226c69962486a9a0/README.md)

`pjpg` — provide general-purpose SQL utilities for time zones and ranges

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "pjpg";
SELECT extversion
FROM pg_extension
WHERE extname = 'pjpg';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
