## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_model_registry/pg_model_registry.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_model_registry/pg_model_registry--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_model_registry/meson.build)

`pg_model_registry` `1.0` provides model versions, deployment metadata and experiments. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_model_registry;
SELECT model_register('demo', 'onnx', 'Example model');
SELECT * FROM model_list();
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

model_version_add stores artifacts and metrics; model_deploy, model_rollback and model_get_active manage active version metadata. experiment_start/end/list track runs. Registry deployment changes metadata; an inference consumer must use the selected artifact. Protect artifact and deployment tables with explicit grants, and include stored binaries and history in backups.

This module corresponds to the former upstream name `alohadb_model_registry`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
