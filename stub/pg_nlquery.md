## Usage

Sources:

- [README.md](https://github.com/schetan/pg_nlquery/blob/d15c30aec342ce1eab707391152b9595978e140a/README.md)
- [pg_nlquery--1.0.sql](https://github.com/schetan/pg_nlquery/blob/d15c30aec342ce1eab707391152b9595978e140a/pg_nlquery--1.0.sql)
- [pg_nlquery.c](https://github.com/schetan/pg_nlquery/blob/d15c30aec342ce1eab707391152b9595978e140a/pg_nlquery.c)
- [schema_introspect.c](https://github.com/schetan/pg_nlquery/blob/d15c30aec342ce1eab707391152b9595978e140a/schema_introspect.c)
- [sql_validate.c](https://github.com/schetan/pg_nlquery/blob/d15c30aec342ce1eab707391152b9595978e140a/sql_validate.c)
- [spi_exec.c](https://github.com/schetan/pg_nlquery/blob/d15c30aec342ce1eab707391152b9595978e140a/spi_exec.c)
- [pg_nlquery.control](https://github.com/schetan/pg_nlquery/blob/d15c30aec342ce1eab707391152b9595978e140a/pg_nlquery.control)

`pg_nlquery` sends a question and selected-schema metadata to an LLM endpoint, validates the generated statement, and optionally executes it. `nlp_explain` returns SQL without executing it; `nlp_query` returns one JSON value per result row. Install as a superuser. The library uses libcurl and does not require preload.

### Core Workflow

```sql
CREATE EXTENSION pg_nlquery;
SET pg_nlquery.api_format = 'openai';
SET pg_nlquery.base_url = 'http://localhost:11434/v1';
SET pg_nlquery.model = 'qwen2.5-coder:7b';
SET pg_nlquery.api_key = 'ollama';
SET pg_nlquery.schema_filter = 'public';
SELECT nlp_explain('show the five most recent orders');
SELECT * FROM nlp_query('show the five most recent orders');
```

### Operational Boundaries

The example uses a local OpenAI-compatible endpoint. A nonempty API key is required even for this local configuration. Set pg_nlquery.schema_filter to limit which schemas are described; column_hints adds explanatory text sent to the model. timeout_ms defaults to 30000 and max_tokens to 1024. Schema metadata, hints and the question leave the database for the configured endpoint. SQL runs with caller privileges. The validator accepts one top-level SELECT, and SPI execution uses read-only mode; this is not a complete SQL-function or external-side-effect sandbox. Review generated SQL and use a limited role. Results are materialized without an extension row limit. The API-key GUC has a superuser-only visibility flag but USERSET context; do not treat it as a credential vault. The README claims PostgreSQL 13+, without a tested-major matrix in this revision.
