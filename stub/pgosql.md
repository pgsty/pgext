## Usage

Sources:

- [Official extension control file](https://github.com/pgoracle/pgora-osql/blob/abb37e20238dc80d29a5440b62d453625e7e766b/pgosql.control)
- [Official upstream documentation](https://github.com/pgoracle/pgora-osql/blob/abb37e20238dc80d29a5440b62d453625e7e766b/README.md)

`pgosql` — C procedural-language handler providing PL/SQL-style compatibility and Oracle-like sys catalog views.

The reviewed catalog snapshot records version `9.4`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pgosql";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgosql';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
