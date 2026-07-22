## Usage

Sources:

- [Project status and deprecation notice](https://github.com/timescale/pgai/blob/47d74affd8a16d0b7f11c8f5bd6e35f7cd43abc1/README.md)
- [PostgreSQL extension README](https://github.com/timescale/pgai/blob/47d74affd8a16d0b7f11c8f5bd6e35f7cd43abc1/projects/extension/README.md)
- [Extension control file](https://github.com/timescale/pgai/blob/47d74affd8a16d0b7f11c8f5bd6e35f7cd43abc1/projects/extension/sql/output/ai.control)
- [API-key and privilege guidance](https://github.com/timescale/pgai/tree/47d74affd8a16d0b7f11c8f5bd6e35f7cd43abc1/projects/extension/docs/security)

ai is the archived PostgreSQL-extension component of pgai. It calls model providers, loads datasets, chunks text, and supports database-side AI/RAG experiments. Upstream states that pgai is no longer maintained or supported as of February 2026, so use this extension only to operate or migrate an existing deployment.

### Core Workflow

The extension is installed in the fixed `ai` schema and requires `vector` and the untrusted `plpython3u` language. Provider functions run inside the database backend. A local Ollama call avoids putting a hosted-provider secret in this minimal example:

```sql
CREATE EXTENSION IF NOT EXISTS ai CASCADE;

SELECT answer->>'response'
FROM ai.ollama_generate(
  'tinyllama',
  'Summarize why database transactions are useful.'
);
```

Use the provider-specific functions for OpenAI, Anthropic, Cohere, Voyage AI, LiteLLM, or Ollama, then persist returned embeddings in a `vector` column when building a retrieval workflow. Network calls are synchronous and are not rolled back when the surrounding SQL transaction aborts.

### Important Objects

- `ai.ollama_generate` sends a prompt to an Ollama endpoint and returns the provider response.
- `ai.openai_embed`, `ai.cohere_embed`, and the other provider families expose model calls from SQL.
- `ai.load_dataset` imports a Hugging Face dataset from the PostgreSQL server process.
- `ai.grant_ai_usage` is a security-definer helper that grants broad schema, table, sequence, and function access; its admin mode also grants grant options.

### Security and Lifecycle Notes

The extension executes untrusted Python in PostgreSQL and can make outbound requests, so restrict installation and function execution to trusted roles. Prefer a provider key store or server-process environment variables; upstream warns that passing keys directly can expose them through logs, SQL text, or monitoring. Review the broad grants made by `ai.grant_ai_usage` before using it. The catalog development version and the archived upstream source do not imply continued compatibility or security fixes; plan to move model calls to a maintained external worker or library.
