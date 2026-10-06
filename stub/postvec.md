## Usage

Sources:

- [0.4.0 migration](https://github.com/univec-ai/postvec/blob/d972a6e7892d90363af677d916b91bef63e33a11/postvec/sql/postvec--0.3.0--0.4.0.sql)
- [postvec/README.md](https://github.com/univec-ai/postvec/blob/d972a6e7892d90363af677d916b91bef63e33a11/postvec/README.md)
- [postvec/postvec.control](https://github.com/univec-ai/postvec/blob/d972a6e7892d90363af677d916b91bef63e33a11/postvec/postvec.control)
- [postvec/Cargo.toml](https://github.com/univec-ai/postvec/blob/d972a6e7892d90363af677d916b91bef63e33a11/postvec/Cargo.toml)
- [README.md](https://github.com/univec-ai/postvec/blob/d972a6e7892d90363af677d916b91bef63e33a11/README.md)
- [postvec/src/api/search.rs](https://github.com/univec-ai/postvec/blob/d972a6e7892d90363af677d916b91bef63e33a11/postvec/src/api/search.rs)
- [postvec/src/api/registry.rs](https://github.com/univec-ai/postvec/blob/d972a6e7892d90363af677d916b91bef63e33a11/postvec/src/api/registry.rs)
- [postvec/src/api/status.rs](https://github.com/univec-ai/postvec/blob/d972a6e7892d90363af677d916b91bef63e33a11/postvec/src/api/status.rs)
- [postvec/sql/postvec--0.2.0--0.3.0.sql](https://github.com/univec-ai/postvec/blob/d972a6e7892d90363af677d916b91bef63e33a11/postvec/sql/postvec--0.2.0--0.3.0.sql)

`postvec` 0.4.0 asynchronously maintains embedding columns and combines lexical and semantic search. The reviewed release is postvec-v0.4.0-1; the package suffix is separate from the extension version.

### Enablement and Core Workflow

Requires PostgreSQL 16–18 and `vector` 0.8 or newer. Add `postvec` to `shared_preload_libraries`, configure the target database and inference backend, then restart. Published packages include embedded and remote modes; a source build without the embedded feature must use remote inference. Make the chosen model available before registering a table.

```sql
CREATE EXTENSION postvec CASCADE;
SELECT postvec.refresh_models();
CREATE TABLE docs (id bigserial PRIMARY KEY, body text);
SELECT postvec.enable('public.docs', 'body',
  model => 'sentence-transformers-all-minilm-l6-v2', create_fts_index => true);
INSERT INTO docs(body) VALUES ('PostgreSQL extension development');
SELECT * FROM postvec.status();
```

Commit the inserted row and wait until `postvec.status()` reports `pending_jobs = 0` before searching:

```sql
SELECT d.body, s.rrf_score
FROM postvec.search('public.docs', 'body', 'database extensions') s
JOIN docs d ON d.id = s.pk_value::bigint ORDER BY s.rrf_score DESC;
```

### Objects and Boundaries

`postvec.enable` installs triggers, a managed vector column and backfill work. Wait for processing before expecting complete search results. `postvec.adopt` registers existing vectors; `postvec.disable` stops management. `postvec.create_vector_index` creates the vector index; automatic index creation can block writes. `postvec.retry_dead` retries eligible dead-letter jobs.

Management requires table ownership or superuser rights. Review row-security restrictions because the worker uses elevated privileges. Remote inference transmits source text outside the database host. The extension uses the PostgreSQL license, while the separate server and downloaded models have their own terms.

### Version 0.4.0 Upgrade

Install matching library and extension files, restart to load the preloaded library, then run `ALTER EXTENSION postvec UPDATE TO '0.4.0'`. The 0.3.0-to-0.4.0 migration takes the extension schema advisory lock and declares no additional schema changes. Check worker health, pending work and search results after upgrading.

Use `postvec.mode = 'embedded'` for local inference or `postvec.mode = 'grpc'` for the separate server. Optional hosted providers transmit input to that provider and require credentials on the inference host. When crossing older releases, apply the intervening migrations: 0.3.0 hardens trigger `search_path` and canonicalizes primary keys, which can requeue backfill. Ambiguous historical keys remain in `jobs_dead` until resolved.
