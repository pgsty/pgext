## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_learned_stats/pg_learned_stats.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_learned_stats/pg_learned_stats--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_learned_stats/meson.build)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_learned_stats/learned_stats.c)

`pg_learned_stats` `1.0` provides experimental learned cardinality estimates for the planner. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_learned_stats;
SELECT * FROM pg_learned_stats_status();
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

Add pg_learned_stats to shared_preload_libraries and restart before using its functions. pg_learned_stats.enabled controls the hooks and pg_learned_stats.train_interval controls training; pg_learned_stats_train and pg_learned_stats_reset manage learned state. This module uses AlohaDB planner integration and shared memory. Rehearse plan changes on representative workloads; learned estimates do not guarantee better plans.

This module corresponds to the former upstream name `alohadb_learned_stats`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
