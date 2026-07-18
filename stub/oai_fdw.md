## Usage

Sources:

- [Official extension control file](https://api.pgxn.org/src/oai_fdw/oai_fdw-1.13.0/oai_fdw.control)
- [Official upstream documentation](https://pgxn.org/dist/oai_fdw/1.13.0/README.html)
- [Official PGXN distribution page](https://pgxn.org/dist/oai_fdw/)

`oai_fdw` — OAI-PMH foreign data wrapper for PostgreSQL

The reviewed catalog snapshot records version `1.13`, kind `standard`, and implementation language `C`.
The curated compatibility set is `11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "oai_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'oai_fdw';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
