## Usage

Sources:

- [Official extension control file](https://github.com/mrcsparker/postllm/blob/0f7d67b574b9d676efaf1024dc201fa563b2c111/postllm.control)
- [Official upstream documentation](https://github.com/mrcsparker/postllm/blob/0f7d67b574b9d676efaf1024dc201fa563b2c111/README.md)
- [Official API reference](https://github.com/mrcsparker/postllm/blob/0f7d67b574b9d676efaf1024dc201fa563b2c111/docs/reference.md)
- [Official configuration guide](https://github.com/mrcsparker/postllm/blob/0f7d67b574b9d676efaf1024dc201fa563b2c111/docs/configuration.md)

`postllm` 0.1.0 makes hosted and local LLM workflows callable from PostgreSQL SQL. It covers chat and completion, embeddings, chunking, retrieval and reranking, plus profiles, secrets, role permissions, model aliases, and outbound-host policy. Inference runs synchronously inside database backends.

### Core Workflow

```sql
CREATE EXTENSION postllm;

SELECT postllm.configure(
  runtime => 'openai',
  model => 'gpt-4o-mini',
  base_url => 'http://127.0.0.1:11434/v1/chat/completions'
);

SELECT postllm.chat_text(ARRAY[
  postllm.system('You are concise.'),
  postllm.user('Explain MVCC in one sentence.')
]);
```

Use `postllm.runtime_discover()` and `postllm.runtime_ready()` before serving traffic. `postllm.settings()` and `postllm.capabilities()` describe the active environment. The API reference groups message builders, generation functions, structured outputs and tools, embeddings, `postllm.ingest_document`, reranking, hybrid ranking, and RAG helpers.

### Governance and Operations

Use the `postllm.profile*`, `postllm.secret*`, `postllm.permission*`, and `postllm.model_alias*` families instead of scattering configuration through application SQL. Restrict `http_allowed_hosts` and `http_allowed_providers`. Model calls can send prompts, rows, documents, and images outside PostgreSQL and can hold a backend while waiting for network or local inference; account for privacy, authorization, timeouts, retries, cost, provider retention, and transaction duration before production use.
