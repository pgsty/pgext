## Usage

Sources:

- [Official source](https://github.com/alitrack/pgtantivy/blob/v0.5.0/README.md)
- [Official source](https://github.com/alitrack/pgtantivy/blob/v0.5.0/Cargo.toml)
- [Official source](https://github.com/alitrack/pgtantivy/blob/v0.5.0/pgtantivy.control)
- [Official source](https://github.com/alitrack/pgtantivy/blob/v0.5.0/src/lib.rs)
- [Official source](https://github.com/alitrack/pgtantivy/blob/v0.5.0/src/am.rs)

`pgtantivy` 0.5.0 provides Tantivy BM25 search with Chinese jieba and English tokenization on PostgreSQL 17–19. It exposes explicit snapshot-index functions, a PostgreSQL index access method, and optional hybrid search through pgvector. Installation requires a superuser; the extension is not trusted or relocatable.

### Snapshot Search

```sql
CREATE EXTENSION pgtantivy;
CREATE TABLE docs (doc_id bigint PRIMARY KEY, body text NOT NULL);
INSERT INTO docs VALUES (1, 'PostgreSQL 数据库全文检索');
SELECT pgtantivy_create(
    'docs_idx',
    (SELECT json_agg(json_build_object('id', doc_id, 'text', body))::text FROM docs),
    'chinese');
SELECT id, score FROM pgtantivy_search('docs_idx', '数据库 检索', 10);
```

Search uses AND semantics across words. In a new backend, call `pgtantivy_load('docs_idx')` before using the named index. `pgtantivy_add` upserts JSON documents by ID, and `pgtantivy_remove` removes supplied IDs. `pgtantivy_snippet` produces highlighted fragments.

### Storage and Synchronization

Index files live under the configurable `pgtantivy.index_dir`, defaulting to a directory inside the PostgreSQL data directory. The loaded registry is per backend. A manually built index is a snapshot: ordinary table writes need explicit synchronization or the documented `pgtantivy_sync_trigger`. Trigger changes are buffered and flushed at PRE_COMMIT; abort clears pending operations. Do not assume these separate Tantivy files have PostgreSQL WAL/MVCC or crash-atomicity guarantees.

`pgtantivy_drift_check` compares a recorded LSN with cluster-wide WAL progress. A positive result is a staleness signal even if only unrelated tables changed. `pgtantivy_touch` advances the anchor; it does not synchronize table contents.

### Other Interfaces and Boundaries

The `pgtantivy` index access method and `@@@` operator support bitmap index scans. The optional `pgtantivy_hybrid` function combines BM25 with pgvector results; the core extension does not declare pgvector as a mandatory dependency.

Local and S3 bundles support export, import and version selection. Bundle FNV checksums detect accidental corruption but do not authenticate untrusted content. SQL functions accept server paths and object-store credentials; restrict execution and filesystem access to trusted roles, and avoid placing credentials in logged SQL. Back up and restore-test sidecar indexes together with the source data, or plan to rebuild them.
