## Usage

Sources:

- [pg_volvec.control](https://github.com/yuuch/pg_volvec/blob/a2fc6dfcc4fc09595f2894de18ff5447c80af941/pg_volvec.control)
- [README.md](https://github.com/yuuch/pg_volvec/blob/a2fc6dfcc4fc09595f2894de18ff5447c80af941/README.md)
- [pg_volvec--1.0.sql](https://github.com/yuuch/pg_volvec/blob/a2fc6dfcc4fc09595f2894de18ff5447c80af941/pg_volvec--1.0.sql)
- [src/bridge/pg_volvec.c](https://github.com/yuuch/pg_volvec/blob/a2fc6dfcc4fc09595f2894de18ff5447c80af941/src/bridge/pg_volvec.c)

`pg_volvec` hooks the PostgreSQL executor and runs supported plan subtrees through a columnar engine, while keeping PostgreSQL planning. It is an experimental executor with LLVM-based expression and tuple-deforming support.

### Core Workflow

```sql
CREATE EXTENSION pg_volvec;
LOAD 'pg_volvec';
SET pg_volvec.enabled = on;
EXPLAIN SELECT sum(v) FROM (VALUES (1), (2), (3)) AS sample(v);
```

### Operational Boundaries

The README requires PostgreSQL 17+ at the specified upstream revision with LLVM JIT enabled; this is not a guarantee for every major or minor build. Install/load as a superuser. The installation SQL loads the library for the current backend; other sessions need their own load or an appropriate preload configuration. `pg_volvec.enabled`, `pg_volvec.trace_hooks` and `pg_volvec.jit_deform` control the implementation. Unsupported plans fall back to PostgreSQL; verify correctness and performance on representative queries. Historical scripts using a different extension name are not the current enablement path.
