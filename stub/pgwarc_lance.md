## Usage

Sources:

- [README.en.md](https://github.com/bjm1244/pgwarc_lance/blob/v0.1.27/README.en.md)
- [pgwarc_lance.control](https://github.com/bjm1244/pgwarc_lance/blob/v0.1.27/pgwarc_lance.control)
- [Cargo.toml](https://github.com/bjm1244/pgwarc_lance/blob/v0.1.27/Cargo.toml)
- [lib.rs](https://github.com/bjm1244/pgwarc_lance/blob/v0.1.27/src/lib.rs)
- [pgwarc_lance--0.1.26--0.1.27.sql](https://github.com/bjm1244/pgwarc_lance/blob/v0.1.27/sql/pgwarc_lance--0.1.26--0.1.27.sql)
- [SECURITY.md](https://github.com/bjm1244/pgwarc_lance/blob/v0.1.27/SECURITY.md)
- [sql-cli-contract.md](https://github.com/bjm1244/pgwarc_lance/blob/v0.1.27/docs/sql-cli-contract.md)
- [LICENSE](https://github.com/bjm1244/pgwarc_lance/blob/v0.1.27/LICENSE)


`pgwarc_lance` 0.1.27 combines table-backed BM25, local Lance vector search and reciprocal-rank fusion for WARC/web archives. It is a research prototype. The current release was validated on Linux x86-64 with PostgreSQL 16.15; other build features do not establish supported release targets.

### Index and Search Text

Install as a superuser. The extension is nonrelocatable and does not require preload. Fresh installation revokes PUBLIC execution on its functions; grant only necessary functions and table access to trusted application roles.

```sql
CREATE EXTENSION pgwarc_lance;
SELECT bm25_index_document(1, 'PostgreSQL archive search');
SELECT bm25_index_document(2, 'Database backup and recovery');
SELECT * FROM bm25_search('archive', 10);
```

BM25 data is stored in the `pgwarc_lance` schema. `tokenize_korean` uses CJK bigrams and ASCII tokenization. `bm25_doc_count` and `bm25_avg_doc_len` expose corpus statistics. `hybrid_search` combines text and vector rankings; `hybrid_warc_search_filtered` applies the same eligibility filter to both candidate sets.

### Lance and the Outbox

`lance_create_table`, `lance_insert_many` and `lance_vector_search` manage and query local datasets. Set the administrator-controlled `pgwarc_lance.allowed_uri_prefix` to an existing local data directory and grant file-access functions only to trusted roles. Remote/object-store Lance and multiple independent writers are unsupported.

Direct Lance writes change external files immediately and are not rolled back by PostgreSQL. For coordinated ingestion, `lance_outbox_enqueue_revision` records vector payload and document revisions in the caller's SQL transaction. After commit, the external worker invokes `lance_outbox_process_next` in standalone autocommit calls. This provides eventual consistency, without an atomic commit across PostgreSQL and Lance.

`lance_outbox_status` exposes pending and failed work. Require completion before relying on vector read-after-write behavior. `lance_document_delete` checks expected revisions and records tombstones plus ordered delete events; it also requires destructive-operation permission and the corresponding administrator setting. Do not modify internal revision or queue tables directly.

### Backup and Upgrade

Lance files are outside PostgreSQL WAL and `pg_dump`. Stop importer and worker activity, reconcile the queue, and preserve matching SQL data and Lance directories. Restore both before resuming work. A Lance-only version rollback does not remove SQL tombstones or restore logical documents.

The supplied 0.1.26-to-0.1.27 migration advances the version without SQL object or grant changes:

```sql
ALTER EXTENSION pgwarc_lance UPDATE TO '0.1.27';
```

Use matching files and an available upgrade chain. Independent relevance evaluation, sustained workload capacity and production support remain open upstream; validate those boundaries before use.
