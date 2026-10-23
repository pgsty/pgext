## Usage

Sources:

- [Upstream README](https://github.com/bitner/pg_table_range/blob/v0.1.2/README.md)
- [Version and PostgreSQL feature matrix](https://github.com/bitner/pg_table_range/blob/v0.1.2/Cargo.toml)
- [Planner-hook and GUC registration](https://github.com/bitner/pg_table_range/blob/v0.1.2/src/lib.rs)
- [0.1.2 upgrade and correctness fixes](https://github.com/bitner/pg_table_range/blob/v0.1.2/CHANGELOG.md)

pg_table_range is an experimental PostgreSQL 16–18 extension for planning-time partition pruning on columns that are not partition keys. Its table_range index access method stores each partition's actual scalar minimum and maximum, range extent, or PostGIS geometry extent. A planner hook removes a partition only when its summary proves that the partition cannot match.

### Basic use

Build the summary index on a partitioned table and compare plans with representative predicates:

```sql
CREATE EXTENSION pg_table_range;

CREATE INDEX events_tr
ON events USING table_range (val, created_at);

EXPLAIN (COSTS OFF)
SELECT *
FROM events
WHERE val >= 250;

SET table_range.enable_pruning = off;
EXPLAIN (COSTS OFF)
SELECT *
FROM events
WHERE val >= 250;
RESET table_range.enable_pruning;
```

The index is never selected to scan rows; it owns compact summaries used by the planner. Inserts widen summaries automatically. Deletes remain safe but can leave them wider than live data until REINDEX tightens them; VACUUM does not rebuild summaries.

Supported pruning predicates include scalar comparisons, BETWEEN, null tests, constant IN or ANY lists, and overlap for range and PostGIS geometry types. Unsupported expressions conservatively keep the partition.

### Upgrade from 0.1.1

Version 0.1.2 fixes missing rows and backend crashes in 0.1.1. Install the new binaries and restart PostgreSQL so no backend still runs the old code. Then update the extension and rebuild every table_range index, using the parent index for partitioned tables:

```sql
ALTER EXTENSION pg_table_range UPDATE TO '0.1.2';
REINDEX INDEX events_tr;
```

Old summaries are ignored until rebuilt, keeping all partitions rather than risking incorrect pruning. Partial indexes are rejected. REINDEX is also the way to tighten summaries after deletes.

### Caveats

- Upstream explicitly labels the project experimental and heavily AI-driven. Qualify correctness, crash recovery, upgrades, and planner behavior on production-shaped data before adoption.
- Planning cost is linear in the partition count. Native partition-key pruning is far cheaper; use pg_table_range only for selective non-key predicates where the execution savings justify planning work.
- Index creation and non-key planning lock every partition. With thousands of partitions, the default max_locks_per_transaction can fail with out-of-shared-memory errors; increasing it consumes shared memory and requires a restart.
- Deletes can reduce selectivity until maintenance rebuilds or tightens summaries. An invalid table_range index silently disables its pruning.
- The control file requires superuser installation. Limit table and index DDL privileges and test the Rust/pgrx build for the exact PostgreSQL major version.
