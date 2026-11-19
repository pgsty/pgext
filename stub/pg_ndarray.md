## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ndarray/pg_ndarray.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ndarray/pg_ndarray--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ndarray/meson.build)

`pg_ndarray` `1.0` provides chunked compressed n-dimensional arrays. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_ndarray;
SELECT ndarray_create('demo', 'float32', '{4,4}', '{2,2}', fill_value => '0');
SELECT ndarray_read('demo', '{0,0}', '{2,2}', '{1,1}');
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

Array metadata and zstd-compressed chunks live in ordinary tables. ndarray_write, ndarray_reduce, ndarray_map and ndarray_storage support updates, computation and inspection; core tensor helpers are AlohaDB-specific. Slices use zero-based start/stop/step bounds and unwritten chunks return the configured fill value. Review array dimensions, chunk sizes, write locks and role access before large workloads.

This module corresponds to the former upstream name `alohadb_ndarray`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
