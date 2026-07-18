## Usage

Sources:

- [Official extension control file](https://github.com/ossc-db/dblink_plus/blob/cb0efffe7f9098e352ccf5da65738630cf40dc5a/dblink_plus.control)
- [Official upstream documentation](https://ossc-db.github.io/dblink_plus/index.html)
- [Official project or provider page](https://ossc-db.github.io/dblink_plus/)

`dblink_plus` — Run SQL against external PostgreSQL, Oracle, MySQL, and SQLite databases from PostgreSQL.

The reviewed catalog snapshot records version `1.0.11`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "dblink_plus";
SELECT extversion
FROM pg_extension
WHERE extname = 'dblink_plus';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
