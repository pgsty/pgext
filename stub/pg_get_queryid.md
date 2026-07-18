## Usage

Sources:

- [Official upstream documentation](https://github.com/pgsentinel/pg_get_queryid/blob/826bc9f71282d0047330e9d6cb4de71471299623/README.md)

`pg_get_queryid` — get last queryid generated/used by pg_stat_statements for a given backend pid

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `pg_stat_statements`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_get_queryid";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_get_queryid';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
