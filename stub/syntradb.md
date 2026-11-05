## Usage

Sources:

- [Upstream source 1](https://github.com/fikrifirmanf/syntradb/blob/1c8e61c1e0a1573df8dace235946d5a887d34527/README.md)
- [Upstream source 2](https://github.com/fikrifirmanf/syntradb/blob/1c8e61c1e0a1573df8dace235946d5a887d34527/crates/syntra-pg/syntradb.control)

`syntradb` 0.12.0 offers explicit analytical functions and persisted column segments alongside PostgreSQL heap storage.

### Core Workflow

`syntra_heap_int4_summary` reads a supported heap column with the caller's snapshot. Segment functions encode and scan fixed data; captured sources and checkpoints require explicit lifecycle management. `syntra_apply_prune_capture` applies and prunes a captured source atomically. These interfaces do not automatically replace ordinary query plans.

```sql
CREATE EXTENSION syntradb;
SELECT syntra_version();
SELECT rows, non_nulls, sum, min, max
FROM syntra_int4_summary(ARRAY[1, NULL, -2, 8]::int4[]);
```

### Operational Boundaries

The tested baseline is PostgreSQL 18.6; installation requires superuser. This is a prototype, and automatic analytical planner paths remain unimplemented. Optional background checkpoint maintenance requires its documented preload, restart and database list. Respect ownership, supported column types and resource limits; stored segments should not be mistaken for automatically fresh copies.
