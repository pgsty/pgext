## Usage

Sources:

- [Official README](https://pgxn.org/dist/bloompg/0.1.2/README.html)
- [Extension control file](https://api.pgxn.org/src/bloompg/bloompg-0.1.2/bloompg.control)
- [Versioned installation SQL](https://api.pgxn.org/src/bloompg/bloompg-0.1.2/sql/bloompg--0.1.2.sql)
- [Configuration reference](https://pgxn.org/dist/bloompg/0.1.2/docs/CONFIGURATION.html)
- [Security policy](https://pgxn.org/dist/bloompg/0.1.2/SECURITY.html)

`bloompg` 0.1.2 is a PostgreSQL 18 planner-hook extension for controlled analytical workloads. It executes selected relation inputs while planning, propagates exact bitmap or Bloom filters across strict hashjoinable equality joins, materializes reduced inputs, and gives the smaller join problem to PostgreSQL for a second planning pass. Existing SQL does not need to change. It is most useful when a selective table can remove a large share of rows from joined tables.

### Core Workflow

Install the library, add `bloompg` to `shared_preload_libraries`, and restart PostgreSQL so every new backend receives its planner hook:

```conf
shared_preload_libraries = 'bloompg'
```

After the restart, create the extension as a superuser in each database that will use it, then run analytical SQL normally:

```sql
CREATE EXTENSION bloompg;
SELECT bloompg_version();

SET bloompg.enable = on;
SET bloompg.transfer_progress_metric = 'ndv';
SET bloompg.transfer_workers = 8;

SELECT count(*)
FROM fact
JOIN dimension USING (dimension_id)
WHERE dimension.region = 'APAC';
```

`bloompg.enable` defaults to `on`. The recommended `bloompg.transfer_progress_metric` is `ndv`, which uses exact distinct-key cardinality to avoid redundant transfers where supported; `rows` retains row-cardinality and filter-lineage propagation. `bloompg.transfer_workers` limits parallel workers for one eligible transfer scan, and zero makes those scans serial.

### Functions and Settings

- `bloompg_version()` returns the installed extension version as `text`.
- `bloompg_last_trace()` returns the latest backend-local transfer trace as `text`.
- `bloompg_last_profile()` returns the latest structured backend-local profile as `jsonb`.
- `bloompg.materialization_memory` limits optimizer-time materializations and parallel DSM copies for one query; its detected-memory default is capped at `2GB`.
- `bloompg.sample_mode` selects persistent `prepared` samples or query-local `instant` samples. `bloompg.sample_cache_dir` chooses the prepared-sample cache directory and is superuser-only.
- `bloompg.profile` collects structured diagnostics, while `bloompg.profile_log` controls whether completed profiles are written to the PostgreSQL log.
- `bloompg.index_transfer` is an experimental, opt-in exact-key B-tree materialization path and is `off` by default.

Enable diagnostics only for the session being investigated, run the query, and inspect the latest backend-local results:

```sql
SET bloompg.profile = on;
SET bloompg.profile_log = off;

SELECT count(*)
FROM fact
JOIN dimension USING (dimension_id)
WHERE dimension.region = 'APAC';

SELECT jsonb_pretty(bloompg_last_profile());
SELECT bloompg_last_trace();
```

### Scope, Fallback, and Safety

Version 0.1.2 is PostgreSQL 18-only and was developed and tested on PostgreSQL 18.4 on Linux. The control file sets `superuser = true` and `trusted = false`, so only a superuser can create it; the extension is relocatable. Changing `shared_preload_libraries` requires a server restart.

The supported core is read-only queries with strict, hashjoinable equality joins. Modifying statements, row locking, row-level security, volatile expressions, and unsupported query scopes bypass BloomPG and retain PostgreSQL planning. Prepared samples influence scheduling and estimates only; filters that eliminate rows are built from exact scans under the active statement snapshot.

BloomPG performs real scans and in-memory materialization during planning. Its query-wide budget defaults to one eighth of detected host or cgroup memory, capped at 2 GB. If the next allocation does not fit, it abandons the optimizer work and executes the original PostgreSQL plan rather than spilling to disk. Size `bloompg.materialization_memory` for peak analytical concurrency, validate native fallback under representative failures, and leave `bloompg.index_transfer` disabled unless its experimental path has been evaluated.

Upstream describes BloomPG as an analytical research extension and does not recommend it as an unreviewed default for multi-tenant or security-sensitive clusters. Test result equivalence, planning latency, memory pressure, parallel-worker demand, hook interaction, and log handling before production use; use `bloompg.enable` as the session-level rollback switch.
