## Usage

Sources:

- [Official upstream documentation](https://github.com/pgEdge/pg_semantic_cache/blob/38e4a910f1ae7e4050abbb8bc1f9cf5d44abf4cc/docs/index.md)

`pg_semantic_cache` — Caches query results and retrieves semantically similar results using pgvector embeddings and configurable eviction policies.

The reviewed catalog snapshot records version `0.1.0-beta4`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `plpgsql`, `vector`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_semantic_cache";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_semantic_cache';
```

The curated lifecycle is `preview`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
