## Usage

Sources:

- [Official upstream documentation](https://github.com/jaiminpan/pg_jieba/blob/d0ffac8028328b2566a889ff4db3d74ba63d1b42/README.md)

`pg_jieba` — Chinese full-text search parser extension using Jieba tokenization.

The reviewed catalog snapshot records version `1.1.1`, kind `standard`, and implementation language `C++`.
The curated compatibility set is `10,11`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_jieba";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_jieba';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
