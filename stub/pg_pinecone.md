## Usage

Sources:

- [Official README](https://github.com/pinecone-io/pg_pinecone/blob/ca5fc2261429aee036b899e372c5772794c9b963/README.md)
- [Control file](https://github.com/pinecone-io/pg_pinecone/blob/ca5fc2261429aee036b899e372c5772794c9b963/pg_pinecone/pg_pinecone.control)
- [Changelog](https://github.com/pinecone-io/pg_pinecone/blob/ca5fc2261429aee036b899e372c5772794c9b963/CHANGELOG.md)
- [License](https://github.com/pinecone-io/pg_pinecone/blob/ca5fc2261429aee036b899e372c5772794c9b963/LICENSE)

`pg_pinecone` 0.11.0 brings Pinecone-compatible vector and full-text search engines into PostgreSQL. It requires `vector` and superuser installation, supports PostgreSQL 17 and 18, and offers local `pgvector` and `native` engines plus a `remote` Pinecone engine.

### Create a Native Index

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_pinecone;
CREATE TABLE items (id bigint PRIMARY KEY, embedding vector(3));
SELECT pinecone.create_index('items', 'embedding');
SELECT id FROM items ORDER BY embedding <=> '[1,2,4]' LIMIT 2;
```

The native engine is the 0.11.0 default and runs without network access; `pgvector` keeps an existing HNSW or IVFFlat path. The remote engine synchronizes a table to a Pinecone serverless index while PostgreSQL retains the rows. Functions and settings live in schema `pinecone`, including `create_index`, `set_engine`, `topk`, `status` and `pinecone.api_key`.

### Installation and Upgrade

Prebuilt archives, Docker images and pgrx builds are described in the upstream README. The source build uses cargo-pgrx 0.19.2. `ALTER EXTENSION pg_pinecone UPDATE` upgrades 0.10.0 to 0.11.0 after matching files are installed; the release changes the native default and repairs cleanup of dropped hnsw indexes. Reindex native inner-product indexes created before 0.11.0 when the changelog calls for it.

### Preload and Security

Native and `pgvector` engines do not need preload. The remote worker needs `shared_preload_libraries = 'pg_pinecone'` and a restart; preserve existing libraries. Remote mode sends vector and metadata data to Pinecone, so protect `pinecone.api_key`, review network and retention policy, and treat the superuser-only extension as a trusted component. The extension is not relocatable.
