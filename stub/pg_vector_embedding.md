## Usage

Sources:

- [Official documentation](https://github.com/hank-cp/pg_vector_embedding/blob/c53bf7f93ad9f4cc27fa657b40749488a1000eae/README.md)
- [Extension control file](https://github.com/hank-cp/pg_vector_embedding/blob/c53bf7f93ad9f4cc27fa657b40749488a1000eae/pg_vector_embedding.control)
- [Official repository](https://github.com/hank-cp/pg_vector_embedding)

`pg_vector_embedding` Trigger-driven vector embedding through external HTTP services and a background queue.

### Enablement

Install the files for the intended server, then create `pg_vector_embedding` in the target database:

```sql
CREATE EXTENSION pg_vector_embedding CASCADE;
```

The reviewed control or official workflow requires `http`, `pg_background_queue`, `vector`. `CASCADE` only succeeds when those extension files are already installed on the server.

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- 1. Setup
CREATE EXTENSION pg_vector_embedding CASCADE;

ALTER DATABASE mydb SET pg_vector_embedding.embedding_url = 'https://api.example.com/v1/embeddings';
ALTER DATABASE mydb SET pg_vector_embedding.embedding_api_key = 'sk-xxxxx';

\c  -- Reconnect to apply settings

-- 2. Create and register table
CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    title TEXT,
    content TEXT,
    embedding VECTOR(1024)
);

SELECT ve_enable('public', 'articles', ARRAY['title', 'content'], 'embedding');

-- 3. Insert data (embeddings computed automatically in background)
INSERT INTO articles (title, content) VALUES
    ('PostgreSQL Extensions', 'Learn how to build powerful PostgreSQL extensions'),
    ('Vector Search', 'Implementing semantic search with pgvector');

-- 4. Wait for background processing (or check if embeddings are ready)
SELECT COUNT(*) FROM articles WHERE embedding IS NOT NULL;

-- 5. Perform similarity search
WITH search_query AS (
    SELECT ve_compute_embedding(
        '{"title":"PostgreSQL","content":"tutorial"}'::text
    ) AS query_embedding
)
SELECT id, title, embedding <-> query_embedding AS distance
FROM articles, search_query
WHERE embedding IS NOT NULL
ORDER BY distance
LIMIT 5;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `ve_compute_embedding` | FUNCTION | Callable function from the reviewed install surface. |
| `ve_enable` | FUNCTION | Callable function from the reviewed install surface. |
| `ve_trigger` | FUNCTION | Callable function from the reviewed install surface. |
| `ve_compact_row_data` | FUNCTION | Callable function from the reviewed install surface. |
| `ve_disable` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Treat extension upgrades as database changes: review the upstream upgrade path, privileges, locks, and backup/restore behavior first.
