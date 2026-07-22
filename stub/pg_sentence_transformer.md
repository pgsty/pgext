## Usage

Sources:

- [pg_sentence_transformer README at the reviewed commit](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/README.md)
- [Version 0.1 install SQL](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/pg_sentence_transformer--0.1.sql)
- [Background-worker implementation](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/pg_sentence_transformer.c)
- [Python model implementation](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/pg_sentence_transformer/model.py)
- [Extension control file](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/pg_sentence_transformer.control)

`pg_sentence_transformer` is a prototype background worker that runs a Hugging Face sentence-transformer model inside a PostgreSQL process. It watches one configured text table, queues inserted or changed rows, and writes `vector(384)` embeddings into a sibling table. This avoids an external inference service but moves Python, model memory, downloads, and inference failures into the database server boundary.

It requires `plpython3u`, `vector`, the extension's Python package and dependencies, and loading through `shared_preload_libraries`.

### Configure and Enable

All extension settings are superuser-only `PGC_POSTMASTER` parameters, so changing them requires a server restart.

```conf
shared_preload_libraries = 'pg_sentence_transformer'
sentence_transformer.venv_path = '/opt/pg_sentence_transformer/venv'
sentence_transformer.database_name = 'appdb'
sentence_transformer.schema_name = 'public'
sentence_transformer.src_table_name = 'posts'
sentence_transformer.src_column = 'body'
sentence_transformer.id_column = 'id'
```

Install the Python environment where the PostgreSQL operating-system user can read it, restart the server, and create the extension in the configured database.

```sql
CREATE EXTENSION pg_sentence_transformer CASCADE;
```

At worker startup, `sentence_transformer.prepare` creates `public.posts_embeddings`-style storage, an `ivfflat` index, enqueue and delete triggers, and initial queue rows. `sentence_transformer.process` claims one queue row with `FOR UPDATE SKIP LOCKED`, computes embeddings, upserts them, and deletes the queue item.

### Fixed Model and Prototype Semantics

The Python module hard-codes `all-MiniLM-L6-v2`, whose output is stored as `vector(384)`. `sentence_transformer.prepare` downloads NLTK `punkt` and `punkt_tab` data, while model initialization can download model files if they are not cached. Provision and verify all artifacts before enabling the worker on a network-restricted or production server.

The processor splits one source value into sentences and upserts every sentence under the same `ref_id`. Consequently, later sentences overwrite earlier embeddings; version `0.1` does not implement document pooling or one row per sentence.

### Operational Notes

- Preloading starts one worker for one configured database, schema, table, source column, and integral ID column. A single server configuration does not manage an arbitrary set of tables.
- The control file says relocatable, but the SQL hard-codes the `sentence_transformer` schema. Treat that schema as fixed.
- Schema, table, and column settings are interpolated into dynamic SQL rather than quoted as identifiers. Keep them to simple trusted identifiers and validate the generated objects before exposing source-table writes.
- Model inference and untrusted Python run inside the PostgreSQL worker. Measure resident memory, CPU contention, restart loops, queue growth, Python ABI compatibility, model-cache ownership, and failure recovery.
- The worker can start before extension SQL exists and is configured to restart after failure. Sequence preload, restart, dependency creation, and extension creation in a controlled maintenance window and confirm the worker becomes healthy afterward.
- Updates to the text column enqueue work and deletes remove embeddings, but changing only the ID column is outside the trigger's update condition. Test all application mutation paths before relying on referential freshness.
