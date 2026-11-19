## Usage

Sources:

- [Official source](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.1/README.md)
- [Official source](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.1/vectorfast.control)
- [Official source](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.1/LICENSE)
- [Official source](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.1/sql/vectorfast--0.2.1.sql)
- [Official source](https://api.pgxn.org/dist/pgvectorfast/0.2.1/META.json)

vectorfast 0.2.1 connects the pgvectorfast distribution's vector extension with Calico. It is an experimental cluster-wide buffer-manager replacement and optimistic HNSW prefetch path, not a per-table switch.

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

calico.optimistic_reads and hnsw.enable_prefetch default to on. The documented optimistic path targets vector_l2_ops. Release 0.2.1 fixes private buffer-pin reservation in the VACUUM fallback. Review cluster-wide correctness and recovery with the exact server build; do not infer support for LTO, other CPU architectures or operating systems.
