## Usage

Sources:

- [Official extension control file](https://github.com/umitanuki/tinyint-postgresql/blob/76495ba27565df2e6758e8f26778e16e963004fb/tinyint.control)
- [Official upstream documentation](https://pgxn.org/dist/tinyint/doc/tinyint.html)
- [Official PGXN distribution page](https://pgxn.org/dist/tinyint/)

`tinyint` — One-byte signed integer type with arithmetic, casts, aggregates, and B-tree, hash, and GIN operator classes.

The reviewed catalog snapshot records version `0.1.1`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "tinyint";
SELECT extversion
FROM pg_extension
WHERE extname = 'tinyint';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
