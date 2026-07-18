## Usage

Sources:

- [pg_ollama README at the reviewed commit](https://github.com/ryicoh/pg_ollama/blob/c4b5dc429f0b9558cf6b9c4f8da5aaf59e3a9df1/README.md)
- [ollama.control at the reviewed commit](https://github.com/ryicoh/pg_ollama/blob/c4b5dc429f0b9558cf6b9c4f8da5aaf59e3a9df1/ollama.control)
- [Version 1.0 install SQL](https://github.com/ryicoh/pg_ollama/blob/c4b5dc429f0b9558cf6b9c4f8da5aaf59e3a9df1/ollama--1.0.sql)

`ollama` is a PL/pgSQL client for an Ollama server. It uses the `http` extension to send chat requests from the PostgreSQL backend, stores endpoint, model, prompt-template, and timeout settings, and records requests and responses in `ollama.logs`.

The SQL interface provides `ollama(text)` for the default setting and `ollama(integer, text)` for an explicit setting ID. `ollama_insert_default_settings` creates an endpoint at `http://localhost:11434`, selects `gemma:2b`, and adds a basic system message.

### Basic Setup

```sql
CREATE EXTENSION IF NOT EXISTS http;
CREATE EXTENSION ollama;

SELECT ollama_insert_default_settings();
SELECT ollama('Explain PostgreSQL MVCC in one sentence.');
```

The Ollama process and selected model must already be available at the configured endpoint.

### Caveats

- The control file does not declare dependencies, but the install SQL needs `http` and `plpgsql`; install the required extensions before `ollama`.
- Prompts, request JSON, responses, and generated output are persisted in `ollama.logs`. Apply appropriate access control and retention policy.
- Calls make synchronous outbound HTTP requests from a database backend. Endpoint failure, model latency, and network policy directly affect the SQL call.
- Catalog and control both identify version `1.0`; upstream provides no broad PostgreSQL compatibility matrix.
