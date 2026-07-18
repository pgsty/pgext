## Usage

Sources:

- [Official extension control file](https://github.com/theory/pg-hostname/blob/37d7fea1d393b3bb2093e419b181c7d0826b3c6e/hostname.control)
- [Official upstream documentation](https://github.com/theory/pg-hostname/blob/37d7fea1d393b3bb2093e419b181c7d0826b3c6e/README.md)
- [Official PGXN distribution page](https://pgxn.org/dist/hostname/)

`hostname` — Get the server host name

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "hostname";
SELECT extversion
FROM pg_extension
WHERE extname = 'hostname';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
