## Usage

Sources:

- [README](https://github.com/univec-ai/postvec/blob/fe3d109d11c2a567ef0e1689b75773787238e08c/postvec/README.md)
- [Control file](https://github.com/univec-ai/postvec/blob/fe3d109d11c2a567ef0e1689b75773787238e08c/postvec/postvec.control)
- [postvec/Cargo.toml](https://github.com/univec-ai/postvec/blob/fe3d109d11c2a567ef0e1689b75773787238e08c/postvec/Cargo.toml)
- [postvec/src/lib.rs](https://github.com/univec-ai/postvec/blob/fe3d109d11c2a567ef0e1689b75773787238e08c/postvec/src/lib.rs)
- [postvec/LICENSE](https://github.com/univec-ai/postvec/blob/fe3d109d11c2a567ef0e1689b75773787238e08c/postvec/LICENSE)

`postvec` maintains embedding columns asynchronously and combines full-text and vector results. These notes describe the pinned 0.2.0 development source; the published 0.1.0-1 artifacts are a separate version boundary.

### Enablement

PostgreSQL 16–18 and `vector` 0.8 or newer are required. Configure an embedded CPU inference backend with downloaded models, or a reachable remote UniVec backend. Add `postvec` to `shared_preload_libraries`, configure `postvec.database`, and restart; remote mode also uses `postvec.grpc_endpoints` and `postvec.http_endpoints`.

```sql
CREATE EXTENSION postvec CASCADE;
SELECT postvec.refresh_models();
```

### Core Workflow

```sql
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

`postvec.enable` adds a shadow vector column, triggers, and backfill work; wait for the queue before expecting complete search results. `postvec.create_vector_index` builds the vector index; its automatic/immediate paths use blocking index creation. Use a separate concurrent index build for busy tables. `postvec.adopt` adopts existing vectors, and `postvec.disable` removes management. Model conversion requires a compatible converter and is not universally lossless.

Management requires ownership or superuser rights. The worker runs as the bootstrap superuser and bypasses RLS, so review the documented row-security restrictions and grants before enabling a table. Remote inference sends source text off the database host. The extension uses the PostgreSQL license; the separate server and downloaded models have their own terms.
