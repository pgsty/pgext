## Usage

Sources:

- [Official README](https://github.com/onelazyteam/pg_llm/blob/master/README.md)
- [Official version 1.1 SQL](https://github.com/onelazyteam/pg_llm/blob/master/sql/pg_llm--1.1.sql)
- [Official control file](https://github.com/onelazyteam/pg_llm/blob/master/pg_llm.control)

`pg_llm` integrates remote and local language models with PostgreSQL. Version `1.1` provides model registration, chat and streaming calls, multi-turn sessions, parallel model selection, text-to-SQL, SQL analysis/reporting, vector-backed knowledge search, feedback, audit logs, and request traces. It depends on `vector` and sends model prompts through backend processes.

### Core Workflow

Preload the library and restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_llm'
```

Install `vector`, create the extension, register a model, and make a call. The example uses a psql variable so the key is not written literally into command history:

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_llm;

SELECT pg_llm_add_model(
  'qwen',
  'qianwen-chat',
  :'api_key',
  '{"model_name":"qwen-turbo","api_endpoint":"https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation"}'
);

SELECT pg_llm_chat('qwen', 'Explain PostgreSQL MVCC briefly.');
```

For multi-turn use, create a session with `pg_llm_create_session(max_messages)` and pass its returned ID to `pg_llm_multi_turn_chat`.

### Main API Groups

- Models/chat: `pg_llm_add_model`, `pg_llm_remove_model`, `pg_llm_chat`, `pg_llm_chat_stream`, `pg_llm_chat_json`, and parallel-chat variants.
- Sessions: `pg_llm_create_session`, `pg_llm_multi_turn_chat`, session lookup/state/update/delete/cleanup functions.
- Data work: `pg_llm_text2sql[_json]`, `pg_llm_execute_sql_with_analysis`, and `pg_llm_generate_report`.
- Retrieval: `pg_llm_get_embedding`, vector storage/search, `pg_llm_add_knowledge`, and `pg_llm_search_knowledge`.
- Observability: `pg_llm_record_feedback`, `pg_llm_get_audit_log`, and `pg_llm_get_trace`.

Model, session, vector, audit, trace, report, knowledge, and feedback state is stored in `_pg_llm_catalog`; version `1.1` creates 64-dimensional pgvector columns and IVFFlat indexes for its internal retrieval tables.

### Caveats

The extension requires PostgreSQL 14+, a C++20 build, libcurl, OpenSSL, JsonCpp, and pgvector. Preloading and outbound HTTP occur inside database server processes. Provider latency, rate limits, TLS/DNS failures, and large/streaming responses can therefore occupy backends and transactions.

Prompts, SQL text, query results, API keys, audit records, and knowledge documents may be sensitive. Restrict `EXECUTE` on network, text-to-SQL, and SQL-execution functions; protect `_pg_llm_catalog`; use a managed secret path; and review retention/logging. Treat generated SQL and model advice as untrusted input—validate permissions, use read-only roles where possible, and never execute it automatically in a privileged session.
