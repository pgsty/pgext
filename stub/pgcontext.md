## Usage

Sources:

- [PGXN pgContext 0.1.0 release](https://pgxn.org/dist/pgcontext/0.1.0/)
- [Official pgContext 0.1.0 README](https://github.com/Evokoa/pgContext/blob/v0.1.0/README.md)
- [Official pgcontext control file](https://github.com/Evokoa/pgContext/blob/v0.1.0/pgcontext.control)
- [Official PostgreSQL 17 quickstart](https://github.com/Evokoa/pgContext/blob/v0.1.0/docs/user_guide/quickstart.md)
- [Official configuration guide](https://github.com/Evokoa/pgContext/blob/v0.1.0/docs/user_guide/configuration.md)
- [Official index guide](https://github.com/Evokoa/pgContext/blob/v0.1.0/docs/user_guide/indexes.md)
- [Official 0.1.0 limitations](https://github.com/Evokoa/pgContext/blob/v0.1.0/docs/user_guide/limitations.md)
- [Official pgvector coexistence guide](https://github.com/Evokoa/pgContext/blob/v0.1.0/docs/user_guide/pgvector_coexist.md)

`pgcontext` 0.1.0 is a PostgreSQL 17 extension for exact vector search, registered-field filtering, persisted HNSW indexing, and dense plus full-text hybrid retrieval over ordinary PostgreSQL tables. Application tables remain authoritative for data, MVCC visibility, ACL/RLS, backup, and replication; HNSW state is a derived, rebuildable index. The stable V1 surface centers on exact table-backed retrieval, while `pgcontext_hnsw`, non-dense vector variants, and several advanced serving paths retain explicit experimental boundaries.

### Core Workflow

Once the extension files are installed into PostgreSQL 17, create a table-backed collection and register only the fields that retrieval is allowed to use:

```sql
CREATE EXTENSION pgcontext;

CREATE TABLE public.docs (
    id text PRIMARY KEY,
    embedding vector(3) NOT NULL,
    status text NOT NULL,
    body text NOT NULL,
    metadata jsonb NOT NULL
);

INSERT INTO public.docs VALUES
    ('postgres', '[1,0,0]', 'published', 'postgres vector search', '{"topic":"database"}'),
    ('rust',     '[0,1,0]', 'published', 'rust extension guide',  '{"topic":"systems"}'),
    ('draft',    '[1,1,0]', 'draft',     'internal notes',        '{"topic":"database"}');

SELECT * FROM pgcontext.create_collection('docs', 'public.docs');
SELECT pgcontext.register_vector('docs', 'embedding', 'embedding', 3, 'cosine');
SELECT pgcontext.register_filter_column('docs', 'status', 'status');
SELECT pgcontext.register_jsonb_path('docs', 'topic', 'metadata', ARRAY['topic']);
SELECT pgcontext.upsert_points('docs', ARRAY['postgres', 'rust', 'draft']);

SELECT source_key, score
FROM pgcontext.search(
    'docs',
    '[1,0,0]'::vector,
    '{"must":[{"key":"status","match":"published"}]}'::jsonb,
    10
);
```

Filters accept only registered columns and JSONB paths. The same filter grammar is shared by search, count, and facets:

```sql
SELECT pgcontext.count(
    'docs',
    '{"must":[{"key":"topic","match":"database"}]}'::jsonb
);

SELECT * FROM pgcontext.facet('docs', 'topic', NULL, 10);
```

### HNSW and Hybrid Retrieval

Add the experimental HNSW access method only after exact search supplies a correctness oracle for the workload:

```sql
CREATE INDEX docs_embedding_hnsw
ON public.docs USING pgcontext_hnsw (
    embedding pgcontext.vector_hnsw_cosine_ops
);

SELECT source_key, score
FROM pgcontext.search('docs', '[1,0,0]'::vector, 10);

SELECT source_key, score
FROM pgcontext.query(
    'docs',
    '[1,0,0]'::vector,
    'postgres search',
    'body',
    10
);
```

Dense HNSW operator classes cover L2, cosine, negative inner product, and L1. Use `pgcontext.index_status`, `pgcontext.index_diagnostics`, and `pgcontext.recall_check` before routing a workload to an approximate path. `pgcontext.query` performs reciprocal-rank fusion for dense plus PostgreSQL full-text retrieval; keep `pgcontext.search` for a single-vector request.

### Important Objects

- `vector` is the stable dense type. `halfvec`, `sparsevec`, and `bitvec` are SQL-visible but experimental in V1.
- `pgcontext.create_collection`, `collection_info`, `drop_collection`, and collection aliases manage registrations over application-owned tables.
- `register_vector`, `register_filter_column`, `register_jsonb_path`, and `upsert_points` establish the explicit retrieval contract.
- `search`, `count`, `facet`, `scroll`, `grouped_search`, `recommend`, and `discover` provide table-backed retrieval and navigation.
- `query` and `explain` expose hybrid retrieval and its SQL-visible execution stages.
- `pgcontext_hnsw` plus the `vector_hnsw_ops`, `vector_hnsw_cosine_ops`, `vector_hnsw_ip_ops`, and `vector_hnsw_l1_ops` operator classes provide dense ANN indexing.
- `optimization_status`, `index_status`, `index_diagnostics`, `vacuum_advice`, `recall_check`, and `telemetry` support operational inspection.

### PostgreSQL and pgvector Boundaries

- V1 supports PostgreSQL 17 only. PostgreSQL 15, 16, and 18 are roadmap targets, not supported 0.1.0 majors.
- The control file sets `superuser = false` and `relocatable = false`; the extension installs the fixed `pgcontext` schema and `$libdir/pgcontext` library. It does not require `shared_preload_libraries`, `LOAD`, or a server restart.
- Installing files on the host still requires filesystem access appropriate to the PostgreSQL 17 installation. A source build is pinned to Rust 1.96.0 and `cargo-pgrx` 0.19.1.
- `pgcontext` can supply its own vector types. On a database that already uses pgvector, install `CREATE EXTENSION vector` first and `CREATE EXTENSION pgcontext` second so pgContext binds to pgvector's existing types.
- pgContext has no required external service. Ordinary PostgreSQL memory, WAL, authentication, backup, and ACL/RLS settings remain the operational foundation.

### V1 Caveats

- `pgcontext_hnsw` is experimental and does not promise long-term on-disk compatibility. Rebuild affected HNSW indexes when an upgrade changes their page format, and qualify recall, latency, MVCC, RLS, restart, VACUUM, and replica behavior on production-shaped data.
- A full HNSW delta segment can trigger synchronous compaction on an insert. The stall grows with index size; tune `pgcontext.hnsw_delta_segment_limit` and `pgcontext.hnsw_compact_on_threshold_max_mb`, or disable automatic threshold compaction and schedule `pgcontext.compact()` or `REINDEX` deliberately.
- Sparse and other non-dense ANN coverage is incomplete. Quantization helpers do not imply quantized HNSW serving, and internally maintained late-interaction or memory-mapped serving remains outside the stable V1 path.
- Remove collection registrations and review dependent application tables before dropping the extension. `DROP EXTENSION pgcontext` intentionally does not treat user tables as disposable acceleration state.
