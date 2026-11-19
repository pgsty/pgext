## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_inference/pg_inference.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_inference/pg_inference--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_inference/meson.build)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_inference/pg_inference.c)

`pg_inference` `1.0` provides onnx inference using alohadb core vector types. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_inference;
SELECT * FROM pg_list_models();
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

pg_load_model accepts model bytes; pg_infer and pg_batch_infer consume AlohaDB vector values. The library must be compiled with ONNX Runtime or model operations fail. pg_inference.device selects CPU/CUDA capability and pg_inference.max_models bounds the backend-local cache. Models must match expected tensor shapes. Treat model bytes as trusted inputs and budget backend memory; loading a model is not a durable deployment across sessions.

This module corresponds to the former upstream name `alohadb_inference`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
