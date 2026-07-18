## Usage

Sources:

- [Official upstream documentation](https://database.dev/langchain/embedding_search)

`embedding_search` — LangChain/Supabase helper package for storing and searching document embeddings with pgvector.

The reviewed catalog snapshot records version `1.1.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `vector`.

```sql
CREATE EXTENSION "embedding_search";
SELECT extversion
FROM pg_extension
WHERE extname = 'embedding_search';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
