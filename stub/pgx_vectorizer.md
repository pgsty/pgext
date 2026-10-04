## Usage

Sources:

- [FEP 18 quick-start guide](https://www.postgresql.fastware.com/knowledge-base/quick-start-guides)
- [FEP 18 knowledge data management manual](https://www.postgresql.fastware.com/hubfs/_Global/Manuals/FEP-v18forx86-KnowledgeDataManagementUserGuide.pdf)

`pgx_vectorizer` 3.0 is bundled with Fujitsu Enterprise Postgres 18 for automatic embeddings and semantic search. It depends on `ai`, `vector` and `plpython3u`.

### Core Workflow

Configure the vendor Python runtime, worker capacity and embedding provider first. Add `pgx_vectorizer` to `shared_preload_libraries` and restart, then enable it in the target database:

```sql
CREATE EXTENSION pgx_vectorizer CASCADE;
SELECT pgx_vectorizer.start_vectorize_scheduler();
```

As the extension owner, register a worker login with `pgx_vectorizer.set_worker_setting`, and configure its connection permissions and password file. With an accessible Ollama service providing `all-minilm`, define vectorization for a table with a primary key and a text column:

```sql
CREATE TABLE sample_table (id bigint PRIMARY KEY, contents text);
INSERT INTO sample_table VALUES (1, 'PostgreSQL supports streaming replication.');
SELECT pgx_vectorizer.pgx_create_vectorizer(
  'sample_table'::regclass,
  destination => 'sample_embeddings',
  embedding => ai.embedding_ollama('all-minilm', 384),
  chunking => ai.chunking_recursive_character_text_splitter('contents'),
  scheduling => pgx_vectorizer.schedule_vectorizer(interval '1 hour')
);
```

After vectorization finishes, search the resulting view:

```sql
SELECT * FROM pgx_vectorizer.pgx_similarity_search(
  'sample_embeddings'::regclass, 'database replication', 5, '<=>');
```

### Operational Boundaries

Configure provider access for both background workers and foreground searches. Search users need execution privileges on functions in the `ai` schema. Protect credentials in `pgx_vectorizer.worker_setting_table`. Renaming the source table or changing its primary key or text-column definition requires redefining vectorization. Pause scheduling for relevant maintenance. This feature requires the Fujitsu runtime.
