## Usage

Sources:

- [Official extension control file](https://api.pgxn.org/src/passwd-fdw/passwd-fdw-0.4.0/passwd-fdw.control)
- [Official upstream documentation](https://pgxn.org/dist/passwd-fdw/0.4.0/README.html)
- [Official PGXN distribution page](https://pgxn.org/dist/passwd-fdw/)

`passwd-fdw` — Foreign data wrapper for Unix/Linux passwd and group databases.

The reviewed catalog snapshot records version `0.4.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "passwd-fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'passwd-fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
