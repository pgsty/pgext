## Usage

Sources:

- [Official native component control file](https://github.com/veridit/sql_saga/blob/9e39fcc7fc862fcb793f59d996eb01268cf9e822/native/sql_saga_native.control)
- [Official native SQL bindings](https://github.com/veridit/sql_saga/blob/9e39fcc7fc862fcb793f59d996eb01268cf9e822/src/2n_temporal_merge_native.sql)
- [Official architecture decision](https://github.com/veridit/sql_saga/blob/9e39fcc7fc862fcb793f59d996eb01268cf9e822/doc/decision-log.md#decision-companion-rust-extension-for-performance-critical-paths-2026-02-22)
- [Official Rust implementation](https://github.com/veridit/sql_saga/blob/9e39fcc7fc862fcb793f59d996eb01268cf9e822/native/src/lib.rs)

`sql_saga_native` is the Rust/pgrx companion library for `sql_saga`'s performance-critical `temporal_merge` planner and executor-introspection cache. It is a headless component with no independently created extension object. The main `sql_saga` extension registers SQL functions backed by the matching `$libdir/sql_saga_native` library.

### Core Workflow

Install a package that provides both matching `sql_saga` and `sql_saga_native` artifacts. After the main extension is installed, inspect the current connection's native cache through its registered functions:

```sql
SELECT stat_name, stat_value
FROM sql_saga.temporal_merge_native_cache_stats();

SELECT sql_saga.temporal_merge_native_cache_reset();
```

Normal applications call `sql_saga.temporal_merge(...)`; they do not call the native planner directly. The main procedure uses the native sweep-line planner by default. For diagnosis, a session can select the PL/pgSQL planner instead:

```sql
SET sql_saga.temporal_merge.use_plpgsql_planner = true;
```

### Native Interfaces

- `temporal_merge_plan_native(...)`: internal planner that fills `pg_temp.temporal_merge_plan` and returns its row count.
- `temporal_merge_executor_introspect(...)`: internal per-connection cache of target/source metadata and generated DML fragments.
- `temporal_merge_native_cache_stats()`: reports planner/executor cache entries, prepared read statements, hits, and misses for the current connection.
- `temporal_merge_native_cache_reset()`: clears current-connection native caches, prepared statements, and counters.

The planner reads source and target rows in bulk, performs temporal segmentation/coalescing in Rust with a sweep-line algorithm, and emits the same plan shape consumed by the PL/pgSQL executor. Cache state is thread-local, which maps to one PostgreSQL backend connection.

### Caveats

There is no independent DDL lifecycle for `sql_saga_native`; version and ABI must match the `sql_saga` SQL bindings that name its C symbols. Do not create a separate extension or expose the internal planner as an application API. Upgrade the main extension and companion library atomically, then exercise temporal merges and cache reset/stats on the target PostgreSQL 18 build.

Caches are per connection, not shared across sessions. Source column names and types participate in cache validation, but operational DDL and role changes should still be tested with fresh sessions. The upstream performance ratios are measurements from its own benchmark, not guarantees for other schemas or workloads.
