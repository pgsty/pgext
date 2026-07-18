## Usage

Sources:

- [Official extension control file](https://github.com/MagicFun1241/pg_bleve/blob/2757cbd1e3db950a723b785611a89d42b4f8c548/pg_bleve.control)
- [Official upstream documentation](https://github.com/MagicFun1241/pg_bleve/blob/2757cbd1e3db950a723b785611a89d42b4f8c548/README.md)

`pg_bleve` — PostgreSQL extension for Bleve search integration

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `Go`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_bleve";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_bleve';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
