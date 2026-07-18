## Usage

Sources:

- [Official upstream documentation](https://docs.pgedge.com/pgedge-vectorizer/)

`pgedge_vectorizer` — Asynchronous text chunking and embedding generation for PostgreSQL using background workers and pgvector.

The reviewed catalog snapshot records version `1.1`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `vector`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgedge_vectorizer";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgedge_vectorizer';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
