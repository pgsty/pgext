## Usage

Sources:

- [pgContext 0.2.0 README](https://github.com/evokoa/pgcontext/blob/v0.2.0/README.md)
- [pgContext 0.2.0 release notes](https://github.com/evokoa/pgcontext/blob/v0.2.0/docs/user_guide/release_notes.md)
- [pgContext collection quickstart](https://github.com/evokoa/pgcontext/blob/v0.2.0/docs/user_guide/quickstart.md)
- [pgContext index guide](https://github.com/evokoa/pgcontext/blob/v0.2.0/docs/user_guide/indexes.md)
- [pgContext known limitations](https://github.com/evokoa/pgcontext/blob/v0.2.0/docs/user_guide/limitations.md)
- [pgContext control file](https://github.com/evokoa/pgcontext/blob/v0.2.0/pgcontext.control)
- [pgvector coexistence guide](https://github.com/evokoa/pgcontext/blob/v0.2.0/docs/user_guide/pgvector_coexist.md)

`pgcontext` keeps vector and hybrid retrieval inside PostgreSQL. It provides pgContext-owned vector types, collection metadata over application tables, registered-field filters, exact search, persisted HNSW, and dense plus full-text fusion. Application rows remain authoritative for MVCC, ACL/RLS, backup, and replication; indexes and generated artifacts are rebuildable acceleration state.

Version 0.2.0 targets PostgreSQL 17 and 18, with the release's controlled-pilot qualification centered on PostgreSQL 17. Advanced HNSW, non-dense, quantized, mapped, and late-interaction paths retain explicit experimental boundaries.

### Core Workflow

```sql
CREATE EXTENSION pgcontext;

CREATE TABLE public.docs (
    id text PRIMARY KEY,
    embedding pgcontext.vector(2) NOT NULL,
    status text NOT NULL,
    body text NOT NULL,
    metadata jsonb NOT NULL
);

INSERT INTO public.docs (id, embedding, status, body, metadata) VALUES
    ('doc-1', '[1,0]'::pgcontext.vector, 'published', 'postgres vector search', '{"topic":"postgres"}'),
    ('doc-2', '[0,1]'::pgcontext.vector, 'published', 'rust extension guide', '{"topic":"rust"}');

SELECT * FROM pgcontext.create_collection('docs', 'public.docs');
SELECT pgcontext.register_vector('docs', 'embedding', 'embedding', 2, 'l2');
SELECT pgcontext.register_filter_column('docs', 'status', 'status');
SELECT pgcontext.register_jsonb_path('docs', 'topic', 'metadata', ARRAY['topic']);
SELECT pgcontext.upsert_points('docs', ARRAY['doc-1', 'doc-2']);

SELECT source_key, score
FROM pgcontext.search(
    'docs',
    '[1,0]'::pgcontext.vector,
    '{"must":[{"key":"status","match":"published"}]}'::jsonb,
    10
);
```

Collections describe application-owned tables; they do not copy those rows into another authoritative store. Search, count, facets, grouping, scrolling, recommendation, and discovery share registered vector and filter definitions.

### HNSW and Hybrid Retrieval

```sql
SET maintenance_work_mem = '2GB';
CREATE INDEX docs_embedding_hnsw ON public.docs
    USING pgcontext_hnsw
    (embedding pgcontext.vector_hnsw_cosine_ops);
RESET maintenance_work_mem;

SELECT source_key, score
FROM pgcontext.query(
    'docs',
    '[1,0]'::pgcontext.vector,
    'postgres search',
    'body',
    10
);
```

Dense HNSW operator classes cover L2, inner product, cosine, and L1. Index builds enforce `maintenance_work_mem`; size the build budget first, then compare approximate results with exact search and `pgcontext.recall_check`. `pgcontext.query` combines dense and PostgreSQL full-text branches with reciprocal-rank fusion.

### Important Objects

- `pgcontext.vector`, `pgcontext.halfvec`, `pgcontext.sparsevec`, and `pgcontext.bitvec` are extension-owned types; the non-dense variants remain experimental.
- `pgcontext.create_collection`, registration functions, and point-mapping functions define the retrieval contract over source tables.
- `pgcontext.search`, `count`, `facet`, `scroll`, `grouped_search`, `recommend`, and `discover` provide table-backed retrieval.
- `pgcontext.query` and `explain` expose composable and hybrid retrieval.
- `pgcontext_hnsw` plus metric-specific operator classes provide ANN indexing.
- Index status, diagnostics, vacuum advice, recall checks, optimization status, and bounded telemetry support operational review.

### Upgrade and pgvector Boundary

Version 0.2.0 moves pgContext-owned types into the fixed `pgcontext` schema. Existing standalone 0.1.0 installations can use the packaged update as a superuser:

```sql
ALTER EXTENSION pgcontext UPDATE TO '0.2.0';
```

Afterward, qualify types such as `pgcontext.vector(1536)` or deliberately add the schema to `search_path`. A 0.1.0 database where public vector types are owned by pgvector is rejected before mutation; inventory dependencies, install 0.2.0 and the separate `pgcontext_pgvector` bridge, recreate registrations and dependents, and rebuild pgContext indexes over the unchanged pgvector columns.

The main extension has no dependency on pgvector. Its types are distinct from `public.vector`, `public.halfvec`, and `public.sparsevec`; do not assume install order aliases one extension's types to the other.

### Operational Boundaries

- `CREATE EXTENSION` and updates require a PostgreSQL superuser because pgContext installs an access method. Granted application APIs do not require superuser.
- No `shared_preload_libraries`, `LOAD`, or restart is required by the main extension.
- The early-release HNSW on-disk format is not stable; plan and validate index rebuilds across releases rather than treating index files as portable data.
- Exact reranking, MVCC, ACL, and RLS checks remain correctness boundaries, but they do not replace workload-specific recall, latency, restart, VACUUM, replication, and failure testing.
- Remove collection registrations and inspect dependent application objects before dropping the extension; avoid unreviewed `CASCADE`.
