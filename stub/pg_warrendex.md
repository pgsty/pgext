## Usage

Sources:

- [Official README.md](https://github.com/soulstompp/pg-warren/blob/4b4cfd4792e9f3c509d7c9187e2a4efe18509780/pg-warrendex/README.md)
- [Official pg_warrendex.control](https://github.com/soulstompp/pg-warren/blob/4b4cfd4792e9f3c509d7c9187e2a4efe18509780/pg-warrendex/pg_warrendex.control)
- [Official Cargo.toml](https://github.com/soulstompp/pg-warren/blob/4b4cfd4792e9f3c509d7c9187e2a4efe18509780/pg-warrendex/Cargo.toml)
- [Official lib.rs](https://github.com/soulstompp/pg-warren/blob/4b4cfd4792e9f3c509d7c9187e2a4efe18509780/pg-warrendex/src/lib.rs)
- [Official pg_warrendex--0.0.0--0.1.0.sql](https://github.com/soulstompp/pg-warren/blob/4b4cfd4792e9f3c509d7c9187e2a4efe18509780/pg-warrendex/sql/pg_warrendex--0.0.0--0.1.0.sql)

`pg_warrendex` 0.1.0 provides the `warrendex` covering index access method. Sorted runs retain keys and included columns so qualifying queries can avoid heap reads once visibility permits. Upstream documents testing on PostgreSQL 18 only; a Cargo feature alone is not evidence of PostgreSQL 19 support.

### Core Workflow

```sql
CREATE EXTENSION pg_warrendex;
CREATE TABLE index_demo (id bigint, group_id integer, label text);
CREATE INDEX index_demo_cover ON index_demo
    USING warrendex (group_id, id) INCLUDE (label);
VACUUM (ANALYZE) index_demo;
EXPLAIN (ANALYZE, BUFFERS)
SELECT id, label FROM index_demo
WHERE group_id = 7 AND id >= 100
ORDER BY id;
```

Installation requires a superuser; the extension is untrusted and not relocatable. The source uses pgrx 0.19.2 and the sibling Rust support crate warren-pg-speller. This crate is a build dependency, not another CREATE EXTENSION dependency.

### Planning and Maintenance

The library installs planner hooks when loaded. To apply them before a session plans its statements, use `LOAD 'pg_warrendex'`; optional `shared_preload_libraries` applies them across sessions and requires a restart. Ordinary index operations and planner-hook optimizations are separate capabilities.

Keys use ascending order with NULLs last; `DESC` and `NULLS FIRST` are rejected. Operator classes cover common integer, text, numeric, temporal, boolean, UUID and binary types. No storage parameters are accepted. Inspect plans and heap fetch counts rather than assuming every query becomes index-only.

Inserts can seal and merge runs inside the writing transaction. VACUUM marks dead entries; `REINDEX` consolidates runs and reclaims their obsolete pages. An index written by another extension version is refused by scans, writes and VACUUM, so rebuild it before use after a version change. Evaluate this early release on representative disposable data.
