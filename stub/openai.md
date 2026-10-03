## Usage

Sources:

- [openai.control](https://github.com/pramsey/pgsql-openai/blob/83eedc5982d5e525e1cb7a0b31315fc75bd05382/openai.control)
- [README.md](https://github.com/pramsey/pgsql-openai/blob/83eedc5982d5e525e1cb7a0b31315fc75bd05382/README.md)
- [openai--1.0.sql](https://github.com/pramsey/pgsql-openai/blob/83eedc5982d5e525e1cb7a0b31315fc75bd05382/openai--1.0.sql)
- [Makefile](https://github.com/pramsey/pgsql-openai/blob/83eedc5982d5e525e1cb7a0b31315fc75bd05382/Makefile)

`openai` provides SQL functions for model listing, prompts, embeddings and image analysis over an OpenAI-compatible endpoint. It uses the `http` extension; the reviewed project is SQL-only despite a library-path field in its control file.

### Core Workflow

```sql
CREATE EXTENSION http;
CREATE EXTENSION openai;
SET openai.api_uri = 'http://127.0.0.1:11434/v1/';
SET openai.api_key = 'none';
SELECT * FROM openai.models();
```

### Operational Boundaries

Installation is privileged and no preload is specified. Configure endpoint, API key and model settings in the session before remote calls. The example uses an already-running local endpoint; it does not install a model server. `openai.models`, `openai.prompt`, `openai.vector` and `openai.image` issue external requests, so payloads leave the database and latency/errors affect the calling query. Protect credentials, limit HTTP timeouts, and validate returned model output. No explicit project license was found.
