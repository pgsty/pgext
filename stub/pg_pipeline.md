## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_pipeline/pg_pipeline.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_pipeline/pg_pipeline--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_pipeline/meson.build)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_pipeline/pg_pipeline.c)

`pg_pipeline` `1.0` provides dag task scheduling and execution history. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_pipeline;
SELECT pipeline_create('demo');
SELECT pipeline_task_add('demo', 'step1', 'sql', 'SELECT 1');
SELECT pipeline_run('demo');
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

The SQL functions enqueue work and manage dependencies/history. To execute queued tasks, preload pg_pipeline, restart and set pg_pipeline.database; pg_pipeline.check_interval controls polling. Task SQL runs inside a privileged background-worker context, so only trusted operators should edit definitions or start runs. Do not treat max_parallel or a queued run identifier as proof that tasks completed; inspect task/run history and cancellation state.

This module corresponds to the former upstream name `alohadb_pipeline`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
