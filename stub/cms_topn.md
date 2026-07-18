## Usage

Sources:

- [Official upstream documentation](https://github.com/ozturkosu/cms_topn/blob/78ce0d1e0437c0b35419d963685d5de57a87078e/README.md)

`cms_topn` — Count-min sketch data type and aggregates for approximate frequency and top-N queries

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "cms_topn";
SELECT extversion
FROM pg_extension
WHERE extname = 'cms_topn';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
