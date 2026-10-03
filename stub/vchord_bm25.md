## Usage

Sources:

- [0.3.0 README](https://github.com/supervc-stack/VectorChord-bm25/blob/0.3.0/README.md)
- [Control file](https://github.com/supervc-stack/VectorChord-bm25/blob/0.3.0/vchord_bm25.control)
- [0.3.0 SQL objects](https://github.com/supervc-stack/VectorChord-bm25/blob/0.3.0/sql/install/vchord_bm25--0.3.0.sql)
- [Query settings](https://github.com/supervc-stack/VectorChord-bm25/blob/0.3.0/src/guc.rs)
- [0.3.0 migration](https://github.com/supervc-stack/VectorChord-bm25/blob/0.3.0/sql/vchord_bm25--0.2.2--0.3.0.sql)
- [0.3.0 release](https://github.com/supervc-stack/VectorChord-bm25/releases/tag/0.3.0)
- [Tokenizer installation](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/01-installation.md)
- [Tokenizer models](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/06-model.md)

`vchord_bm25` provides BM25 ranking with a sparse token-frequency type and the `bm25` index access method. Tokenization is supplied separately, commonly by pg_tokenizer. Extension objects live in the fixed `bm25_catalog` schema, and creation requires superuser privileges.

### Core Workflow

The example uses pg_tokenizer, which requires preloading and a restart. Preserve existing entries in the preload list:

```conf
shared_preload_libraries = 'pg_tokenizer'
```

```sql
CREATE EXTENSION pg_tokenizer;
CREATE EXTENSION vchord_bm25;
SET search_path = public, tokenizer_catalog, bm25_catalog;

SELECT create_tokenizer('english', $$
model = "bert_base_uncased"
$$);
CREATE TABLE documents (
    id bigserial PRIMARY KEY,
    passage text,
    embedding bm25vector
);
INSERT INTO documents(passage) VALUES ('PostgreSQL full text search');
UPDATE documents SET embedding = tokenize(passage, 'english')::bm25vector;
CREATE INDEX documents_bm25 ON documents USING bm25 (embedding bm25_ops);

SELECT id, passage,
       embedding <&> to_bm25query('documents_bm25',
           tokenize('PostgreSQL', 'english')::bm25vector) AS score
FROM documents
ORDER BY score
LIMIT 10;
```

The index supplies corpus statistics to `to_bm25query`; the score from `<&>` is negative, so ascending order returns greater relevance first. Use the same tokenizer/model for documents and queries. Update stored token vectors when source text changes, or use the tokenizer's maintenance-trigger helper. Changing the vocabulary requires retokenizing stored documents before rebuilding their index.

### Types, Functions, and Search Limits

- `bm25vector` stores token IDs and frequencies; the integer-array cast aggregates duplicate IDs and discards token order.
- `bm25query` binds the query vector to an index. `to_bm25query(regclass, bm25vector)` constructs it; `bm25_ops` is the index operator class.
- `bm25_catalog.bm25_limit` defaults to 100 and limits candidates returned by the index. Increase it for larger SQL limits or restrictive filters; changing SQL LIMIT alone does not increase this candidate budget.
- `bm25_catalog.enable_index` controls use of the index; `bm25_catalog.enable_prefilter` controls prefiltering. Both default to true.
- `bm25_catalog.segment_growing_max_page_size` defaults to 4096 pages before sealing a growing segment.

```sql
SET bm25_catalog.bm25_limit = 1000;
```

The access-method name is global: this extension cannot coexist in a database with another extension that creates the same bm25 access method, including pg_textsearch and the compatibility alias in pg_search. Sparse frequencies do not preserve positions for phrase matching. Chinese text can use a custom corpus model with a Jieba pre-tokenizer; Japanese Lindera support depends on the tokenizer build and dictionary configuration.

### Upgrade to 0.3.0

```sql
ALTER EXTENSION vchord_bm25 UPDATE TO '0.3.0';
```

The 0.2.2-to-0.3.0 migration adds `bm25_page_inspect(regclass, integer)`, returning diagnostic page text. The release changes sealed-segment page allocation for small tokens and does not document a mandatory index rebuild. Install matching extension files before updating database objects; replacing the preloaded tokenizer library also requires a restart. Keep tokenization and ranking upgrades compatible and check representative query results.
