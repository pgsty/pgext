## Usage

Sources:

- [Official extension control file](https://github.com/zdk/pg-search-thai/blob/e943164a3640a65c5714028447189f9a3a7b4f82/thai_parser/thai_parser.control)
- [Official upstream documentation](https://github.com/zdk/pg-search-thai/blob/e943164a3640a65c5714028447189f9a3a7b4f82/README.md)

`thai_parser` — Thai-language parser and text search configuration for PostgreSQL full-text search.

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "thai_parser";
SELECT extversion
FROM pg_extension
WHERE extname = 'thai_parser';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
