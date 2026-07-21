## Usage

Sources:

- [Official README for version 1.0](https://github.com/knizhnik/online_advisor/blob/1.0/README.md)
- [Extension control file](https://github.com/knizhnik/online_advisor/blob/1.0/online_advisor.control)
- [Version 1.0 SQL objects](https://github.com/knizhnik/online_advisor/blob/1.0/online_advisor--1.0.sql)
- [Sample preload configuration](https://github.com/knizhnik/online_advisor/blob/1.0/online_advisor.conf)

`online_advisor` observes PostgreSQL execution plans and workload timing, then recommends indexes, extended statistics, or prepared statements. It reports candidates only; it never creates an index or statistics object automatically.

### Core Workflow

Preload the library and restart PostgreSQL:

```conf
shared_preload_libraries = 'online_advisor'
```

Create and activate the extension in each database whose workload should be observed:

```sql
CREATE EXTENSION online_advisor;

-- Calling an extension function activates collection in this database.
SELECT get_executor_stats();
```

After representative workload has run, inspect the recommendations:

```sql
SELECT * FROM proposed_indexes;
SELECT * FROM proposed_statistics;
SELECT * FROM get_executor_stats();

-- Keep separate index candidates instead of combining compatible clauses.
SELECT * FROM propose_indexes(combine => false);
```

Review each generated `create_index` or `create_statistics` statement before applying it. Run `ANALYZE` after creating an index or statistics object so the planner can use current statistics.

### Objects and Settings

- `proposed_indexes`: view over `propose_indexes(combine, reset)` with filtering volume, call count, elapsed time, and a candidate `CREATE INDEX` statement.
- `proposed_statistics`: view over `propose_statistics(combine, reset)` with misestimation, call count, elapsed time, and a candidate `CREATE STATISTICS` statement.
- `get_executor_stats(reset)`: returns aggregate planning and execution time, query count, and planning-overhead ratios.
- `online_advisor.filtered_threshold`: minimum filtered-row count considered for an index proposal; default `1000`.
- `online_advisor.misestimation_threshold`: actual-to-estimated row ratio considered for statistics; default `10`.
- `online_advisor.min_rows`: minimum returned rows for misestimation analysis; default `1000`.
- `online_advisor.max_index_proposals` and `online_advisor.max_stat_proposals`: proposal capacities; set them before the extension is activated.
- `online_advisor.do_instrumentation`, `online_advisor.log_duration`, and `online_advisor.prepare_threshold`: control collection and prepared-statement advice.

### Caveats

- Instrumentation adds workload overhead; measure it on the target system and disable collection when it is not needed.
- The index heuristic does not reason about operator ordering in compound indexes, join indexes, or indexes used only to avoid sorting.
- The extension does not estimate the benefit of a proposed index. Use plan review or a hypothetical-index tool before building expensive indexes.
- Advice is database-local and depends on the workload observed since activation or reset.
