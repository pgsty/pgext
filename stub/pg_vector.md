## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_vector/pg_vector.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_vector/pg_vector--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_vector/meson.build)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_vector/vector_ivfflat.c)

`pg_vector` `1.0` provides named ann index helpers over alohadb core vectors. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_vector;
SELECT vector_l2_distance('[1,2]'::vector, '[2,3]'::vector);
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

The vector type and core HNSW/IVFFlat access methods already belong to the AlohaDB kernel. This module adds vector_index_create/insert/rebuild/search and its own index metadata; it is distinct from pgvector. Rebuild trains the helper index, and applications must maintain helper entries explicitly. Choose dimensions, metric and probe/list counts deliberately; approximate search can omit neighbors and helper state must be backed up with source data.

This module corresponds to the former upstream name `alohadb_vector`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
