## Usage

Sources:

- [Official extension control file](https://github.com/huangjimmy/pg_cjk_parser/blob/master/pg_cjk_parser.control)

`pg_cjk_parser` — Full-text search parser that splits CJK text into 2-gram tokens

The reviewed catalog snapshot records version `0.2.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_cjk_parser";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_cjk_parser';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
