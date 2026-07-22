## Usage

Sources:

- [Official README](https://github.com/JoelDiaz222/pg_gembed/blob/c7f4d005a2462a1af9464424a5418cdca40bd0a4/README.md)
- [Official extension SQL](https://github.com/JoelDiaz222/pg_gembed/blob/c7f4d005a2462a1af9464424a5418cdca40bd0a4/sql/pg_gembed--1.0.0.sql)
- [Official background-worker implementation](https://github.com/JoelDiaz222/pg_gembed/blob/c7f4d005a2462a1af9464424a5418cdca40bd0a4/src/embedding_worker.c)

`pg_gembed` 1.0.0 generates text, image, and multimodal vector embeddings from SQL by calling the local Rust `libgembed` core. It supports local or remote pluggable backends and returns the `vector` type, so the `vector` extension is a required dependency.

### Core Workflow

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_gembed;

SELECT embed_text(
  'embed_anything',
  'Qdrant/all-MiniLM-L6-v2-onnx',
  'Hello world'
);

CREATE TABLE articles (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  content text NOT NULL,
  embedding vector(384)
);

INSERT INTO articles (content, embedding)
VALUES (
  'Transformers use attention mechanisms.',
  embed_text(
    'embed_anything',
    'Qdrant/all-MiniLM-L6-v2-onnx',
    'Transformers use attention mechanisms.'
  )
);

SELECT id, content
FROM articles
ORDER BY embedding <=> embed_text(
  'embed_anything',
  'Qdrant/all-MiniLM-L6-v2-onnx',
  'machine learning'
)
LIMIT 10;
```

The backend and model strings select an implementation and model known to `libgembed`. Model files, network access, credentials, and acceleration requirements depend on that backend.

### Important Objects

- `embed_text(text, text, text)` and `embed_texts(text, text, text[])` embed one or many texts.
- `embed_image(text, text, bytea)`, `embed_images(text, text, bytea[])`, and their with-ID variants embed image data.
- `embed_multimodal(text, text, bytea[], text[])` combines image and text inputs.
- `embed_image_directory(text, text, text)` and `embed_image_directories(text, text, text[])` read image paths on the database server.
- `input_type` and the scalar/array `embed(...)` overloads dispatch text, image, or image-directory inputs.
- `gembed.embedding_jobs` and `gembed.job_status` store and report asynchronous embedding jobs.

### Preload and Worker Boundary

Synchronous embedding functions can load the library in a session without a restart. The background job worker is different: `_PG_init` registers it only while PostgreSQL processes `shared_preload_libraries`. To use `gembed.embedding_jobs`, preload `pg_gembed` and restart PostgreSQL. `gembed.embedding_worker_naptime` defaults to 10 seconds and `gembed.embedding_worker_batch_size` defaults to 256; both are reloadable GUCs.

### Operational Notes

The control file requires superuser installation and `vector`; the README lists PostgreSQL 17 or later. The worker in the pinned 1.0.0 source opens a fixed database named `joeldiaz`, so review or patch that code before relying on asynchronous jobs in another database. Directory functions read server-local paths, and embedding backends may load large models or contact configured gRPC/HTTP services. Size model memory, CPU/GPU use, network exposure, and parallel query behavior on the target system before production use.
