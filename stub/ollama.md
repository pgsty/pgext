## Usage

Sources:

- [pg_ollama README at the reviewed commit](https://github.com/ryicoh/pg_ollama/blob/c4b5dc429f0b9558cf6b9c4f8da5aaf59e3a9df1/README.md)
- [Version 1.0 SQL interface](https://github.com/ryicoh/pg_ollama/blob/c4b5dc429f0b9558cf6b9c4f8da5aaf59e3a9df1/ollama--1.0.sql)
- [ollama.control at the reviewed commit](https://github.com/ryicoh/pg_ollama/blob/c4b5dc429f0b9558cf6b9c4f8da5aaf59e3a9df1/ollama.control)

`ollama` is a PL/pgSQL client for an external Ollama server. It stores model, endpoint, timeout, and system-message settings in the `ollama` schema, sends synchronous chat requests through the `http` extension, and records every request and response. It is useful for small SQL-driven model calls where blocking database backends and prompt persistence are acceptable; it is not an in-database model runtime.

### Core Workflow

Install the undeclared `http` dependency first. `ollama_insert_default_settings()` creates a one-time default using `http://localhost:11434`, model `gemma:2b`, a 60-second timeout, and a system message.

```sql
CREATE EXTENSION IF NOT EXISTS http;
CREATE EXTENSION ollama;

SELECT ollama_insert_default_settings();
SELECT ollama('Explain PostgreSQL MVCC in one sentence.');
```

The Ollama process must already be reachable from the database server and the configured model must already exist there. For a custom deployment, populate `ollama.models`, `ollama.endpoints`, `ollama.settings`, and ordered `ollama.setting_messages` rows instead of calling the default helper.

### SQL Interface

- `ollama(text)` uses the single setting whose `is_default` flag is true.
- `ollama(integer, text)` uses an explicit setting ID.
- `ollama_insert_default_settings()` inserts the sample endpoint, model, setting, and system message. It is not idempotent because repeated calls collide with the model primary key.
- `ollama.logs` stores the input, JSON request, HTTP status, JSON response, extracted output, and timestamps. Failed requests can leave an unfinished log row because the row is inserted before the HTTP call.

The returned content is extracted with JSON operators and coerced to the declared text result, so callers should confirm whether their client receives JSON-style quoting before composing it into other text.

### Operational Notes

- The control file does not declare dependencies, but the install SQL references `http_response`, `http_post`, and `http_set_curlopt`; install `http` before `ollama`.
- Each call occupies a PostgreSQL backend while waiting for `/api/chat`. Endpoint failures, model latency, and network policy directly affect query latency and transaction lifetime.
- Prompts, system messages, complete request JSON, responses, and outputs are persisted in ordinary tables. Restrict schema access, define retention, and do not send secrets or regulated data without an approved data boundary.
- Endpoint URLs are database data and control outbound traffic from the PostgreSQL host. Limit who can modify settings and apply network egress controls to reduce server-side request forgery risk.
- Catalog and control both identify version `1.0`; upstream publishes no broad PostgreSQL compatibility matrix. Test the SQL dependency and HTTP behavior on the exact server major used in production.
