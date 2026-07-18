## Usage

Sources:

- [Upstream README](https://github.com/zhangh43/vectorize_engine/blob/e886c5d2b3e61d1c1e8b34cfa46664a454a7ef05/README.md)
- [Extension control file](https://github.com/zhangh43/vectorize_engine/blob/e886c5d2b3e61d1c1e8b34cfa46664a454a7ef05/vectorize_engine.control)
- [SQL installation script](https://github.com/zhangh43/vectorize_engine/blob/e886c5d2b3e61d1c1e8b34cfa46664a454a7ef05/vectorize_engine--1.0.sql)
- [Planner-hook implementation](https://github.com/zhangh43/vectorize_engine/blob/e886c5d2b3e61d1c1e8b34cfa46664a454a7ef05/plan.c)

`vectorize_engine` version `1.0` is a prototype vectorized executor for PostgreSQL 9.6. A post-planner hook replaces supported sequential-scan and aggregate nodes with custom scan nodes; a GUC leaves the feature disabled by default.

### Example

Build PostgreSQL with suitable SIMD flags, put `vectorize_engine` in `shared_preload_libraries`, restart, and then run:

```sql
CREATE EXTENSION vectorize_engine;
SET enable_vectorize_engine = on;
EXPLAIN (ANALYZE, BUFFERS) SELECT category, sum(amount)
FROM benchmark_rows GROUP BY category;
```

The source copies and modifies PostgreSQL executor code and is explicitly based on PostgreSQL 9.6; the repository has not changed since 2020. Only a limited node/type/expression subset is vectorized, with fallback intended for unsupported plans. It is not compatible with modern server internals without substantial porting. Use only on a matching disposable benchmark instance and compare every result and plan with the feature disabled.
