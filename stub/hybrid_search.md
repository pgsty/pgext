## Usage

Sources:

- [Official hybrid_search package page](https://database.dev/langchain/hybrid_search)
- [Version 1.1.1 package SQL and documentation](https://github.com/supabase/dbdev/blob/8b3b966d97f87027a36c052e679ebb45f7e25736/supabase/migrations/20231207112129_langchain%40hybrid_search-1.1.1.sql)
- [Official dbdev project README](https://github.com/supabase/dbdev/blob/8b3b966d97f87027a36c052e679ebb45f7e25736/README.md)

`hybrid_search` is the catalog name for the database.dev package `langchain@hybrid_search`. It installs a small LangChain/Supabase schema that combines pgvector cosine-distance search with PostgreSQL full-text ranking over a shared `documents` table.

### Core Workflow

Install the package through `dbdev`, create its required `vector` dependency, and enable the registry extension at version `1.1.0`:

```sql
SELECT dbdev.install('langchain@hybrid_search');
CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION "langchain@hybrid_search"
    SCHEMA public
    VERSION '1.1.0';

INSERT INTO documents (content, metadata)
VALUES ('PostgreSQL supports full-text search', '{"topic":"postgres"}');

SELECT *
FROM kw_match_documents('full text', 10);
```

The installed objects are `documents(id, content, metadata, embedding)`, `match_documents(query_embedding vector(1536), match_count integer, filter jsonb)`, and `kw_match_documents(query_text text, match_count integer)`. `match_documents` orders by cosine distance and supports a JSONB containment filter; `kw_match_documents` ranks `to_tsvector(content)` matches.

### Operational Notes

The package requires `vector`, and the official version `1.1.0` instructions state that dependency resolution is manual. Its table and functions use fixed, unqualified names, and the embedding column is `vector(1536)`; inspect conflicts and adapt your embedding model accordingly. The two functions return separate result sets—merging, deduplicating, and reranking them is performed by the application-side hybrid retriever, not by one SQL function in this package.
