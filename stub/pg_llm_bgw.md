## Usage

Sources:

- [Official upstream README](https://github.com/cicorias/postgres-llm-extension-bw/blob/c3b7c81fea02bdea8b8ab559c0b56e656af4bb9f/README.md)
- [Official extension control file (pg_llm_bgw.control)](https://github.com/cicorias/postgres-llm-extension-bw/blob/c3b7c81fea02bdea8b8ab559c0b56e656af4bb9f/pg_llm_bgw/pg_llm_bgw.control)
- [Official implementation source](https://github.com/cicorias/postgres-llm-extension-bw/blob/c3b7c81fea02bdea8b8ab559c0b56e656af4bb9f/pg_llm_bgw/src/lib.rs)

`pg_llm_bgw` — example of Rust implementation of background workder to call LLM from within PostGres client DML SQL. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_llm_bgw;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `llm_ask` is an extension function.
- `llm_provider()` is an extension function.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
