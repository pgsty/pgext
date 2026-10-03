## Usage

Sources:

- [postvec/README.md](https://github.com/univec-ai/postvec/blob/6ca448ca9b90cb8c45bef9cce8019a7053bf0bd9/postvec/README.md)
- [postvec/postvec.control](https://github.com/univec-ai/postvec/blob/6ca448ca9b90cb8c45bef9cce8019a7053bf0bd9/postvec/postvec.control)
- [postvec/Cargo.toml](https://github.com/univec-ai/postvec/blob/6ca448ca9b90cb8c45bef9cce8019a7053bf0bd9/postvec/Cargo.toml)
- [postvec/sql/managed/functions.sql](https://github.com/univec-ai/postvec/blob/6ca448ca9b90cb8c45bef9cce8019a7053bf0bd9/postvec/sql/managed/functions.sql)
- [postvec/sql/managed/lifecycle.sql](https://github.com/univec-ai/postvec/blob/6ca448ca9b90cb8c45bef9cce8019a7053bf0bd9/postvec/sql/managed/lifecycle.sql)
- [postvec/sql/postvec--0.2.0--0.3.0.sql](https://github.com/univec-ai/postvec/blob/6ca448ca9b90cb8c45bef9cce8019a7053bf0bd9/postvec/sql/postvec--0.2.0--0.3.0.sql)

`postvec` 0.3.0 asynchronously maintains embedding columns and combines lexical and semantic search. The reviewed release is postvec-v0.3.0-1; the package suffix is separate from the extension version.

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
SELECT d.body, s.rrf_score
FROM postvec.search('public.docs', 'body', 'database extensions') s
JOIN docs d ON d.id = s.pk_value::bigint ORDER BY s.rrf_score DESC;
```

### Objects and Boundaries

`postvec.enable` installs triggers, a managed vector column and backfill work. Wait for processing before expecting complete search results. `postvec.adopt` registers existing vectors; `postvec.disable` stops management. `postvec.create_vector_index` creates the vector index; automatic index creation can block writes. `postvec.retry_dead` retries eligible dead-letter jobs.

Management requires table ownership or superuser rights. Review row-security restrictions because the worker uses elevated privileges. Remote inference transmits source text outside the database host. The extension uses the PostgreSQL license, while the separate server and downloaded models have their own terms.

### Version 0.3.0 Upgrade

After updating the library, `ALTER EXTENSION postvec UPDATE TO '0.3.0'` applies the SQL migration. It hardens trigger `search_path` and canonicalizes date/time-sensitive primary keys. The migration can reset backfill checkpoints and queue work again. Ambiguous old keys are retained as evidence in `jobs_dead`; automatic retry refuses them until their identity is resolved. Budget for the additional backfill and inspect migration results before assuming all vectors are current.
