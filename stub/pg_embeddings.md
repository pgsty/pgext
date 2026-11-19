## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_embeddings/pg_embeddings.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_embeddings/pg_embeddings--1.1.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_embeddings/meson.build)

`pg_embeddings` `1.1` provides embedding lifecycle and hybrid retrieval over alohadb vectors. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_embeddings;
SELECT embedding_binary_quantize('[1,-1,0.5]'::vector);
```

The control file does not require a superuser, but the installer still needs database/schema creation privileges.

### Operational boundaries

Register a real SQL text-to-vector function with embedding_model_register, then use embedding_track, embedding_process and embedding_hybrid_search. Queue processing must be scheduled explicitly; it is not automatically enabled by CREATE EXTENSION. Model dimensions and caller table/RLS rights apply. The module uses AlohaDB core vector helpers, not the separate pgvector extension. Before upgrading an older 1.0 installation, follow the documented regprocedure-to-text migration for embedding_model.embed_fn.

This module corresponds to the former upstream name `alohadb_embeddings`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
