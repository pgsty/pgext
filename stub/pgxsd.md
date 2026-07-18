## Usage

Sources:

- [Official extension control file](https://github.com/johto/pgxsd/blob/ff89bb220cf47ab287b900b47f9eb645088b4422/pgxsd.control)

`pgxsd` — Validate PostgreSQL XML values against XSD schemas stored inside the database.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pgxsd";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgxsd';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
