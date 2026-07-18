## Usage

Sources:

- [Official extension control file](https://github.com/postgrespro/tsvector2/blob/f3fd908c4b09729524052ca99790e30058aa112b/tsvector2.control)
- [Official upstream documentation](https://pgxn.org/dist/tsvector2/1.0.0/)
- [Official PGXN distribution page](https://pgxn.org/dist/tsvector2/)

`tsvector2` — Compressed alternative to PostgreSQL tsvector without the original 1 MB limit, with compatible search functions and GIN, GiST, and optional RUM indexing.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "tsvector2";
SELECT extversion
FROM pg_extension
WHERE extname = 'tsvector2';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
