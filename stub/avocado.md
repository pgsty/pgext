## Usage

Sources:

- [Pinned upstream PostgreSQL extension guide](https://github.com/avocadodb/avocadodb/blob/18d30fd558a4c7f538f9667f788291e2ad6d43f9/README.md#postgresql-extension-new-in-v22)
- [Pinned extension control file](https://github.com/avocadodb/avocadodb/blob/18d30fd558a4c7f538f9667f788291e2ad6d43f9/avocado-pgext/avocado.control)
- [Pinned extension schema](https://github.com/avocadodb/avocadodb/blob/18d30fd558a4c7f538f9667f788291e2ad6d43f9/avocado-pgext/sql/schema.sql)
- [Pinned SQL function source](https://github.com/avocadodb/avocadodb/blob/18d30fd558a4c7f538f9667f788291e2ad6d43f9/avocado-pgext/src/lib.rs)

`avocado` version `2.2.0` is a pgrx PostgreSQL extension for ingesting text artifacts, generating embeddings, retrieving relevant spans, and compiling citation-bearing context for AI agents. It stores data in ordinary PostgreSQL tables under schema `avocado` and uses `vector` for its 1,024-dimensional HNSW index.

### Core Workflow

The control file requires `vector`, marks installation as superuser-only, and makes the extension non-relocatable. After installation, initialize the tables before ingesting content:

```sql
CREATE EXTENSION vector;
CREATE EXTENSION avocado;
SELECT avocado_init();

SELECT avocado_ingest_artifact(
  'docs/auth.md',
  'Authentication uses short-lived access tokens and rotating refresh tokens.',
  '{"type":"markdown"}'::jsonb
);

SELECT avocado_search_spans('How are access tokens refreshed?', 5);
SELECT avocado_compile(
  'How does authentication work?',
  '{"token_budget":2000,"enable_mmr":true}'::jsonb
);
```

FastEmbed with `all-MiniLM-L6-v2` is the default embedding provider. An Ollama endpoint can be selected for the current backend before ingestion or search:

```sql
SELECT avocado_set_embedding_provider('ollama');
SELECT avocado_set_ollama_config('http://localhost:11434', 'bge-m3');
SELECT avocado_embedding_config();
SELECT avocado_test_embedding('configuration check');
```

### Important Objects

- `avocado_ingest_artifact(path, content, metadata)`: chunks text by lines, embeds each span, and stores the artifact and spans.
- `avocado_search_spans(query, limit)`: returns matching spans and cosine-similarity scores as `jsonb`.
- `avocado_compile(query, config)`: applies vector search, optional MMR diversification, token-budget packing, deterministic path/line ordering, and returns context plus citations as `jsonb`.
- `avocado_create_session()`, `avocado_add_message()`, and `avocado_get_conversation_history()`: maintain conversation state.
- `avocado_register_agent()` and relation functions: maintain agent metadata and agreement/disagreement records.
- `avocado_stats()` and `avocado_version()`: report stored-object counts, embedding configuration, and extension version.

### Operational Notes

The reviewed build exposes only a PostgreSQL 16 pgrx feature and its container is based on PostgreSQL 16; verify any other server version independently. The default FastEmbed model can require a model download and substantial backend memory, while Ollama performs synchronous HTTP calls from a PostgreSQL backend. Restrict execution privileges, outbound access, and provider configuration accordingly, and set statement timeouts appropriate for embedding latency.

Embeddings are padded or truncated to `vector(1024)`, so changing providers can mix models and dimensions in one store; re-ingest data rather than comparing incompatible embeddings. The reviewed `2.2.0` ingestion path creates a new artifact identifier before upserting by unique path, so test repeated ingestion of the same path and span replacement before relying on updates. Back up the `avocado` tables as normal database data and validate HNSW build cost, model reproducibility, and retrieval quality on the actual corpus.
