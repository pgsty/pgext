## Usage

Sources:

- [README](https://github.com/Tecnisys-OSS/pgembed/blob/1aed2a0f7e7c7be729cb9558c031133453aeeee3/README.md)
- [Control file](https://github.com/Tecnisys-OSS/pgembed/blob/1aed2a0f7e7c7be729cb9558c031133453aeeee3/pgembed.control)
- [sql/pgembed--0.1.0.sql](https://github.com/Tecnisys-OSS/pgembed/blob/1aed2a0f7e7c7be729cb9558c031133453aeeee3/sql/pgembed--0.1.0.sql)
- [LICENSE](https://github.com/Tecnisys-OSS/pgembed/blob/1aed2a0f7e7c7be729cb9558c031133453aeeee3/LICENSE)

`pgembed` generates pgvector embeddings through Ollama, OpenAI, or a custom HTTP endpoint. It is implemented with untrusted PL/Python; inference runs in an external service, not inside the database.

### Core Workflow

A superuser must make `plpython3u` and `vector` available. Start a compatible Ollama server with the requested model before running the query.

```sql
CREATE EXTENSION plpython3u;
CREATE EXTENSION vector;
CREATE EXTENSION pgembed;
SELECT pgembed.embed_ollama('PostgreSQL extensions', 'embeddinggemma');
SELECT * FROM pgembed.embed_batch_ollama(
  ARRAY['PostgreSQL extensions', 'Vector similarity'], 'embeddinggemma');
```

### API and Configuration

The `embed_ollama`, `embed_openai`, and `embed_custom` functions in the `pgembed` schema return vectors; their batch forms handle text arrays. Match the destination vector dimension to the selected model. `pgembed.url_allowlist` controls accepted hosts. `pgembed.max_retries`, `pgembed.initial_backoff_ms`, and `pgembed.max_backoff_ms` govern retries; `pgembed.circuit_breaker_threshold` and `pgembed.circuit_breaker_reset_timeout_s` govern a backend-local circuit breaker.

The README specifies PostgreSQL 13 or newer. No preload or server restart is needed for the extension itself. HTTP calls hold a database backend until completion; remote text and credentials require deliberate access control. A configurable hostname allowlist is not a substitute for database privileges and network policy. Do not put actual API keys in logged SQL. The Tecnisys Community License 1.0 restricts competing hosted-service use.
