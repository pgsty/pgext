## Usage

Sources:

- [Official upstream README](https://github.com/Percona-Lab/clickhousedb_fdw/blob/c6b386f93d97430ea4da1fd680f005a84a2cf3ec/README.md)

`clickhousedb_fdw` — ClickHouse foreign data wrapper with SELECT/INSERT support plus aggregate and join pushdown.

The reviewed catalog snapshot records version `1`, kind `standard`, and implementation language `C`.
The curated compatibility set is `11,12,13`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "clickhousedb_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'clickhousedb_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
