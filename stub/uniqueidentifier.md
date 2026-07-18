## Usage

Sources:

- [Official extension control file](https://github.com/autodatadirect/uniqueidentifier/blob/85976b564e5cd3362c30ba388b95c3df19bc4b34/uniqueidentifier.control)

`uniqueidentifier` — Legacy 16-byte uniqueidentifier data type, comparison operators, casts, and newid() generation for PostgreSQL.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "uniqueidentifier";
SELECT extversion
FROM pg_extension
WHERE extname = 'uniqueidentifier';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
