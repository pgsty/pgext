## Usage

Sources:

- [pgContext 0.2.0 pgvector coexistence guide](https://github.com/evokoa/pgcontext/blob/v0.2.0/docs/user_guide/pgvector_coexist.md)
- [pgContext 0.2.0 pgvector migration guide](https://github.com/evokoa/pgcontext/blob/v0.2.0/docs/user_guide/pgvector_migration.md)
- [pgcontext_pgvector control file](https://github.com/evokoa/pgcontext/blob/v0.2.0/pgcontext_pgvector.control)
- [pgcontext_pgvector extension SQL](https://github.com/evokoa/pgcontext/blob/v0.2.0/sql/pgcontext_pgvector--0.2.0.sql)
- [pgContext 0.2.0 release notes](https://github.com/evokoa/pgcontext/blob/v0.2.0/docs/user_guide/release_notes.md)

`pgcontext_pgvector` is the optional pgContext companion bridge for serving pgContext HNSW indexes over columns owned by the pgvector extension. It does not merge the two type systems or copy application data; it adds certified casts, support functions, and operator classes while exact distance semantics remain bound to pgvector operators.

### Certified Profile and Installation

Version 0.2.0 fails closed unless the database uses PostgreSQL 17, pgContext 0.2.0, and pgvector 0.8.x installed in `public`. Install the prerequisites and bridge explicitly:

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pgcontext;
CREATE EXTENSION pgcontext_pgvector;
```

The reverse order of the two prerequisite extensions is also valid, but `pgcontext_pgvector` must come after both. Installation requires superuser privileges.

### Index an Existing pgvector Column

```sql
CREATE INDEX items_embedding_pgc
    ON items USING pgcontext_hnsw
       (embedding pgcontext.vector_hnsw_pgvector_cosine_ops);

SELECT id
FROM items
ORDER BY embedding <=> $1::public.vector
LIMIT 10;
```

Existing pgvector-spelled SQL can use the pgContext access method. ANN candidates are resolved to live heap rows and reranked with the pgvector operator, preserving its `double precision` distance result semantics.

### Important Objects

- `pgcontext.vector_hnsw_pgvector_l2_ops`, `pgcontext.vector_hnsw_pgvector_ip_ops`, `pgcontext.vector_hnsw_pgvector_cosine_ops`, and `pgcontext.vector_hnsw_pgvector_l1_ops` serve existing `public.vector` columns.
- `pgcontext.sparsevec_hnsw_pgvector_cosine_ops` serves certified `public.sparsevec` columns, subject to the documented 16,000-dimension and page-envelope limits.
- `pgcontext.migration_report()` inventories pgvector columns, dependencies, HNSW, and IVFFlat without requiring the bridge.
- Ownership-conversion functions provide reviewed fast or restricted-online workflows; IVFFlat is rebuilt as HNSW rather than converted in place.

### Dependency and Removal Boundaries

The main `pgcontext` extension remains independent of pgvector. Bridge indexes depend on `pgcontext_pgvector`, and the bridge depends on both parent extensions, so PostgreSQL blocks removal under `RESTRICT` until those indexes are removed or converted.

Do not use `DROP EXTENSION vector CASCADE` as a migration method. Inventory arrays, views, functions, prepared sessions, expression indexes, and other application dependencies first. The bridge does not provide every pgvector helper, IVFFlat, iterative-scan GUC, parallel-build, subvector, or progress-reporting behavior.

No preload or restart is required. The bridge is a privileged compatibility surface, not a promise that all future pgContext, pgvector, PostgreSQL-major, or on-disk index combinations are compatible; rerun the certified preflight and rebuild validation when any component changes.
