## Usage

Sources:

- [Official extension control file](https://database.dev/langchain/hybrid_search)
- [Official upstream README](https://github.com/supabase/dbdev/blob/8b3b966d97f87027a36c052e679ebb45f7e25736/README.md)

`hybrid_search` — Hybrid pgvector and full-text search helper functions for LangChain and Supabase.

The reviewed catalog snapshot records version `1.1.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`, `vector`.

```sql
CREATE EXTENSION "hybrid_search";
SELECT extversion
FROM pg_extension
WHERE extname = 'hybrid_search';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
