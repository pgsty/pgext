## Usage

Sources:

- [extension/README.md](https://github.com/cpuchip/pg-ai-stewards/blob/c78212cf4dbca4a711cbc6416e06dad25d576b1d/extension/README.md)
- [extension/Cargo.toml](https://github.com/cpuchip/pg-ai-stewards/blob/c78212cf4dbca4a711cbc6416e06dad25d576b1d/extension/Cargo.toml)
- [extension/pg_ai_stewards.control](https://github.com/cpuchip/pg-ai-stewards/blob/c78212cf4dbca4a711cbc6416e06dad25d576b1d/extension/pg_ai_stewards.control)
- [extension/src/lib.rs](https://github.com/cpuchip/pg-ai-stewards/blob/c78212cf4dbca4a711cbc6416e06dad25d576b1d/extension/src/lib.rs)
- [extension/src/bgworker.rs](https://github.com/cpuchip/pg-ai-stewards/blob/c78212cf4dbca4a711cbc6416e06dad25d576b1d/extension/src/bgworker.rs)
- [extension/init/00-extensions.sql](https://github.com/cpuchip/pg-ai-stewards/blob/c78212cf4dbca4a711cbc6416e06dad25d576b1d/extension/init/00-extensions.sql)

`pg_ai_stewards` 0.3.0 stores agent state, work queues, model configuration and dispatch results in PostgreSQL 18. The native Rust core depends on `vector`; its generated installation SQL embeds the migration chain at the pinned source revision.

### Enable and inspect

Preserve other entries when adding the preload library and restart PostgreSQL. The worker connects to the database named by the server environment variable `POSTGRES_DB`, defaulting to `stewards`; create the extension in that database.

```conf
shared_preload_libraries = 'pg_ai_stewards'
```

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_ai_stewards;
SELECT stewards.version();
SELECT * FROM stewards.providers_loaded();
SELECT stewards.enqueue('echo', 'echo', '{"hello":"world"}'::jsonb);
SELECT * FROM stewards.work_queue ORDER BY id DESC LIMIT 5;
```

### Runtime boundary

Creation requires a superuser. Preloading registers dispatch workers; `STEWARDS_DISPATCHER_WORKERS` defaults to 4 and is bounded to 1–16. Provider configuration determines credentials, models and destinations. An echo task checks the queue path without requiring a real model; actual inference needs configured providers.

The `stewards` schema contains queues, configuration, pipelines and history. Restrict access to credentials and dispatch/tool functions, and review external requests and workload costs. The upstream deployment also applies a runtime migration manifest and optional overlays: these are separate from the core extension bundle. Follow manifest order rather than replaying SQL files alphabetically. Installing the companion pack or UI is not implied by creating the core.
