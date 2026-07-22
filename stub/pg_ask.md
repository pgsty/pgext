## Usage

Sources:

- [Official README](https://github.com/Abiji-2020/pg_ask/blob/7c58aeb31042b75c4f1f675b621329c689003e2c/README.md)
- [Extension SQL](https://github.com/Abiji-2020/pg_ask/blob/7c58aeb31042b75c4f1f675b621329c689003e2c/pg/pg_ask--1.0.sql)
- [SQL-generation entry point](https://github.com/Abiji-2020/pg_ask/blob/7c58aeb31042b75c4f1f675b621329c689003e2c/src/pg_ask.cpp)
- [Schema explorer](https://github.com/Abiji-2020/pg_ask/blob/7c58aeb31042b75c4f1f675b621329c689003e2c/src/explorer.cpp)
- [AI client implementation](https://github.com/Abiji-2020/pg_ask/blob/7c58aeb31042b75c4f1f675b621329c689003e2c/src/ai_engine.cpp)

`pg_ask` sends a natural-language prompt and database-schema metadata to an OpenAI-compatible endpoint and returns generated SQL. Use `pg_gen_query(text)` only as a review aid under a tightly restricted role; `pg_gen_execute(text)` executes untrusted model output with the caller's database privileges and is unsafe to expose directly.

### Core Workflow

Set `PG_ASK_AI_KEY` in the PostgreSQL server process environment before starting the server, install the extension, and configure the endpoint and model for a dedicated role. Revoke the default public function access before any use:

```sql
CREATE EXTENSION pg_ask;

REVOKE EXECUTE ON FUNCTION pg_gen_query(text) FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION pg_gen_execute(text) FROM PUBLIC;
GRANT EXECUTE ON FUNCTION pg_gen_query(text) TO sql_preview_role;

ALTER ROLE sql_preview_role SET pg_ask.model = 'gpt-4o';
ALTER ROLE sql_preview_role SET pg_ask.endpoint = 'https://api.openai.com/v1';

SET ROLE sql_preview_role;
SELECT pg_gen_query('count orders created today');
```

Always inspect and validate the returned text before running it through a separately controlled execution path. The extension does not parse, allow-list, parameterize, or force read-only SQL.

### API and Configuration

- `pg_gen_query(text) → text` collects non-system table and column metadata, performs a synchronous provider request, and returns model-generated SQL.
- `pg_gen_execute(text) → refcursor` calls the generator and opens a fixed cursor named `ai_query_result` using dynamic SQL; cursor use requires an explicit transaction.
- `pg_ask.model` is a user-settable model name.
- `pg_ask.endpoint` is a user-settable provider base URL.
- `PG_ASK_AI_KEY` is inherited from the PostgreSQL server environment.

`pg_gen_execute(text)` would be consumed with `BEGIN`, `FETCH ALL FROM ai_query_result`, and `COMMIT`, but granting it lets the model issue any statement allowed to the caller. Prefer not to grant it at all.

### Security and Operational Boundaries

The schema explorer sends schema, table, column, and type names from non-system schemas together with the user's prompt. Treat this as metadata disclosure and review provider retention, residency, and confidentiality requirements. Prompts and schema details leave the transaction through a synchronous network call that cannot be rolled back.

Because `pg_ask.endpoint` is user-settable while the server supplies `PG_ASK_AI_KEY`, a grantee may redirect requests and exfiltrate the key or reach internal HTTP services. Restrict function and configuration privileges, constrain outbound network access, and do not reuse a broadly privileged provider credential.

The SQL file declares `pg_gen_query(text)` as `PARALLEL SAFE` even though the implementation reads catalogs and performs network I/O; do not rely on that declaration for safe parallel execution. The README targets PostgreSQL 16+, while other repository material has mentioned older versions, so verify the exact packaged build and test failure, timeout, and provider-response behavior before adoption.
