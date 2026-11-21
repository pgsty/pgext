## Usage

Sources:

- [Official source](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.2/CHANGELOG.md)

- [Official source](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.2/README.md)
- [Official source](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.2/vectorfast.control)
- [Official source](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.2/LICENSE)
- [Official source](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.2/sql/vectorfast--0.2.1.sql)
- [Official source](https://api.pgxn.org/dist/pgvectorfast/0.2.2/META.json)

vectorfast 0.2.2 connects the pgvectorfast distribution's vector extension with Calico. It is an experimental cluster-wide buffer-manager replacement and optimistic HNSW prefetch path, not a per-table switch.

### Enablement

Use Linux x86-64 and PostgreSQL 18 or 19 with matching non-LTO server/debug symbols and runtime Capstone. Preserve other preload entries and restart the server.

```ini
shared_preload_libraries = 'vectorfast'
```

```sql
CREATE EXTENSION vector;
CREATE EXTENSION calico;
CREATE EXTENSION vectorfast;
SELECT calico.status();
SELECT * FROM calico.pool;
```

Run creation as a superuser outside an explicit transaction block: vectorfast activates Calico and depends on both vector and calico. Activation affects ordinary relations and every database. After a later server restart, call calico.activate() again outside a transaction block; installed SQL objects alone do not reactivate it.

```sql
CALL calico.activate();
```

### Operation

calico.optimistic_reads and hnsw.enable_prefetch default to on. Version 0.2.2 covers L2 and built-in inner-product optimistic reads, including normalized cosine search. Release 0.2.1 fixes private buffer-pin reservation in the VACUUM fallback. Review cluster-wide correctness and recovery with the exact server build; do not infer support for LTO, other CPU architectures or operating systems.

### Distribution 0.2.2

Version 0.2.2 extends optimistic reads to built-in inner-product distance, including HNSW cosine search over normalized vectors; the earlier L2-only description no longer covers the supported path. It fixes visited-cache ownership for interleaved scans and honors `hnsw.enable_prefetch` when a cursor resumes. Iterative searches use compact candidate queues and batched I/O. The `calico` SQL version remains 0.2; the `vectorfast` SQL version is 0.2.2. Preserve the cluster-wide preload, restart, explicit activation and non-LTO platform requirements above.
