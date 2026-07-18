## Usage

Sources:

- [Official extension control file](https://github.com/CrunchyData/crunchy_check_access/blob/c8f01b40d7507715c4f0fda18a651916b010653a/check_access.control)
- [Official upstream documentation](https://github.com/CrunchyData/crunchy_check_access/blob/c8f01b40d7507715c4f0fda18a651916b010653a/README.md)

`check_access` — Functions and views to facilitate PostgreSQL object access inspection

The reviewed catalog snapshot records version `0.1`, kind `puresql`, and implementation language `SQL`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "check_access";
SELECT extversion
FROM pg_extension
WHERE extname = 'check_access';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
