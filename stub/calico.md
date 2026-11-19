## Usage

Sources:

- [Upstream source](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.1/README.md)
- [Upstream source](https://api.pgxn.org/dist/pgvectorfast/0.2.1/META.json)

- [Official source](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.0/README.md)
- [Official source](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.0/calico.control)
- [Official source](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.0/sql/calico--0.2.sql)
- [Official source](https://api.pgxn.org/src/pgvectorfast/pgvectorfast-0.2.0/LICENSE)
- [Official source](https://api.pgxn.org/dist/pgvectorfast/0.2.0/META.json)

`calico` is the buffer-manager extension in the pgvectorfast 0.2.1 distribution; its SQL version is 0.2. Activation replaces the shared buffer manager for the entire PostgreSQL cluster, including ordinary relations and every database. It is an experimental operational change, with no workload-independent performance guarantee.

### Enable and inspect

Use Linux x86-64 and PostgreSQL 18 or 19 with matching non-LTO server symbols or debug symbols. Capstone must remain available at runtime. Add the library to the existing preload list and restart, then run the following as a superuser. The activation call must be outside a transaction block.

```ini
shared_preload_libraries = 'vectorfast'
```

```sql
CREATE EXTENSION calico;
SELECT calico.status();
CALL calico.activate();
SELECT * FROM calico.pool;
```

### Interface and scope

`calico.status()` reports the current mode. `calico.activate(timeout_ms)` defaults to 30000 milliseconds; public execution is revoked. `calico.pool_statistics()` and `calico.pool` expose pool mappings and statistics. Activation lasts until the server restarts. Creating SQL objects alone does not activate the replacement.

The distribution also supplies `vector`, an existing extension identity with compatible pgvector types and indexes; it is not a control-file dependency of `calico`. Its library is `vectorfast`. `calico.optimistic_reads` enables optimistic HNSW reads and `hnsw.enable_prefetch` enables grouped prefetch; both default to on. The documented optimistic path is for `vector_l2_ops`. Do not assume arbitrary PostgreSQL builds, LTO binaries, other operating systems or architectures support interception.

### Distribution 0.2.1

The 0.2.1 distribution fixes private buffer-pin reservation when VACUUM falls back from optimistic reads. The calico SQL version remains 0.2; deploy the matching vectorfast library and follow its preload/restart and explicit activation procedure. The separate vectorfast extension now has its own identity and depends on vector and calico.
