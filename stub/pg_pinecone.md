## Usage

Sources:

- [Official README](https://github.com/pinecone-io/pg_pinecone/blob/v0.12.1/README.md)
- [Control file](https://github.com/pinecone-io/pg_pinecone/blob/v0.12.1/pg_pinecone/pg_pinecone.control)
- [Changelog](https://github.com/pinecone-io/pg_pinecone/blob/v0.12.1/CHANGELOG.md)
- [Upgrade runbook](https://github.com/pinecone-io/pg_pinecone/blob/v0.12.1/docs/runbook.md)
- [Hosted embedding tutorial](https://github.com/pinecone-io/pg_pinecone/blob/v0.12.1/docs/tutorial-gemini-document-search.md)
- [License](https://github.com/pinecone-io/pg_pinecone/blob/v0.12.1/LICENSE)

`pg_pinecone` 0.12.1 brings Pinecone-compatible vector and full-text search engines into PostgreSQL. It requires `vector` and superuser installation, supports PostgreSQL 17 and 18, and offers local `pgvector` and `native` engines plus a `remote` Pinecone engine.

### Create a Native Index

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_pinecone;
CREATE TABLE items (id bigint PRIMARY KEY, embedding vector(3));
INSERT INTO items VALUES (1, '[1,2,3]'), (2, '[4,5,6]');
SELECT pinecone.create_index('items', 'embedding');
SELECT id FROM items ORDER BY embedding <=> '[1,2,4]' LIMIT 2;
```

The native engine remains the local default and runs without network access; `pgvector` keeps an existing HNSW or IVFFlat path. The remote engine synchronizes a table to a Pinecone serverless index while PostgreSQL retains the rows. Functions and settings live in schema `pinecone`, including `create_index`, `set_engine`, `topk`, `status` and `pinecone.api_key`.

### Installation and Upgrade

Prebuilt archives, Docker images and pgrx builds are described in the upstream README. The source build uses cargo-pgrx 0.19.2. Version 0.12.1 re-releases 0.12.0 under a fresh tag to repair the release workflow; extension code and the empty 0.11.0-to-0.12.0 upgrade script are unchanged.

For an in-place upgrade from 0.10.0 or later, install matching files, restart PostgreSQL to load the new library, then run `ALTER EXTENSION pg_pinecone UPDATE` in each database. Check `pinecone.worker_version()` for matching compiled and installed versions. Native inner-product indexes created before 0.11.0 still need `REINDEX`; updating the extension does not rebuild them. The older `pinegres` extension has no in-place upgrade path.

### Preload and Security

Native and `pgvector` engines do not need preload. The remote worker needs `shared_preload_libraries = 'pg_pinecone'` and a restart; preserve existing libraries. Remote mode sends vector and metadata data to Pinecone, so protect `pinecone.api_key`, review network and retention policy, and treat the superuser-only extension as a trusted component. The extension is not relocatable.

### Hosted Document Embedding in 0.12

Track a text column with an `embed` model option. Pinecone then embeds the column server-side on writes and queries; no pgvector column or local `pinecone.embed` call is needed. This sends the configured text to Pinecone and requires the remote worker and a stored API key configured through `pinecone.set_api_key`. The local engine refuses `mode => 'embed'` because PostgreSQL full-text search cannot run the hosted model. `gemini-embedding-2` requires a Builder or higher Pinecone plan and is unavailable through the standalone `pinecone.embed` API.

```sql
CREATE TABLE docs (id bigint PRIMARY KEY, body text);
INSERT INTO docs VALUES (1, 'PostgreSQL stores relational data.');
SELECT pinecone.track_docs(
    'docs',
    text_columns => '{"body": {"embed": {"model": "gemini-embedding-2"}}}'
);
SELECT pinecone.flush_docs('docs');
SELECT pinecone.search_docs('docs', 'query text', mode => 'embed');
```

`pinecone.flush_docs` waits for delivery, but remote indexing is eventually consistent, so an immediate search may not see the new rows. Only one text column per index may use `embed`, and it cannot be combined with a supplied vector column.
