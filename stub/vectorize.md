


## Usage

Sources: [repo README](https://github.com/ChuckHend/pg_vectorize/blob/v0.26.2/README.md), [extension README](https://github.com/ChuckHend/pg_vectorize/blob/v0.26.2/extension/README.md), [v0.26.2 release](https://github.com/ChuckHend/pg_vectorize/releases/tag/v0.26.2)

`vectorize` is the PostgreSQL extension from `pg_vectorize`. Upstream documents two modes: a standalone HTTP service and the in-database SQL extension. For the packaged extension here, the SQL workflow is the relevant one.

### Enable The Extension

```sql
ALTER SYSTEM SET shared_preload_libraries = 'vectorize,pg_cron';
ALTER SYSTEM SET cron.database_name = 'postgres';

CREATE EXTENSION vectorize CASCADE;
```

The extension README lists `pg_cron`, `pgmq`, and `pgvector` as dependencies, plus `vectorize.embedding_service_url` for the embedding service.

### Create A Search Job

The high-level SQL API starts with `vectorize.table()`:

```sql
SELECT vectorize.table(
  job_name    => 'product_search_hf',
  relation    => 'products',
  primary_key => 'product_id',
  columns     => ARRAY['product_name', 'description'],
  transformer => 'sentence-transformers/all-MiniLM-L6-v2',
  schedule    => 'realtime'
);
```

The extension README says this creates and maintains an embeddings column for the source table.

### Search, RAG, And Direct Model Calls

Search with:

```sql
SELECT * FROM vectorize.search(
  job_name       => 'product_search_hf',
  query          => 'accessories for mobile devices',
  return_columns => ARRAY['product_id', 'product_name'],
  num_results    => 3
);
```

Upstream also documents:

- `vectorize.rag()` for retrieval-augmented answers.
- `vectorize.generate()` for text generation.
- `vectorize.encode()` for direct embedding generation.
- `vectorize.import_embeddings()` for loading precomputed vectors.

### Update Behavior And v0.26.2 Note

The extension README says `schedule => '* * * * *'` checks for updates every minute, while `schedule => 'realtime'` creates triggers for immediate refresh on inserts and updates.

`pg_vectorize` 0.26.2 updates vector-serve and security-related dependencies; there is no material SQL/API delta beyond the existing README surface.
