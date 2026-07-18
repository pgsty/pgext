## Usage

Sources:

- [Official extension control file](https://github.com/umitanuki/twitter_fdw/blob/9433a19b1ee7abb1ac08c4471ece2238785d677b/twitter_fdw.control)
- [Official PGXN distribution page](https://pgxn.org/dist/twitter_fdw/)

`twitter_fdw` — Legacy foreign data wrapper for the retired unauthenticated Twitter Search API.

The reviewed catalog snapshot records version `1.1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "twitter_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'twitter_fdw';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
