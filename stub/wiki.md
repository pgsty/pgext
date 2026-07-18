## Usage

Sources:

- [Official extension control file](https://github.com/lacanoid/pgwiki/blob/dbcc433acd6fb2dacc78d0621103a8e2d5c187ec/wiki.control)
- [Official project or provider page](https://github.com/lacanoid/pgwiki/wiki)

`wiki` — In-database wiki data model and SQL/PLPerl API with optional MediaWiki integration.

The reviewed catalog snapshot records version `0.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plperl`, `plperlu`.
The curated compatibility set is `10,11,12`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "wiki";
SELECT extversion
FROM pg_extension
WHERE extname = 'wiki';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
