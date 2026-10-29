## Usage

Sources:

- [Current reviewed research limitations](https://github.com/gburd/ra/blob/8166ea89e91bfef98e1c31fe140b9525d37a1d33/README.md)
- [Ra v0.6.0 README](https://github.com/gburd/ra/blob/v0.6.0/README.md)
- [Extension control file](https://github.com/gburd/ra/blob/v0.6.0/crates/ra-pg-extension/pg_ra_planner.control)
- [Extension Cargo manifest](https://github.com/gburd/ra/blob/v0.6.0/crates/ra-pg-extension/Cargo.toml)
- [Extension SQL](https://github.com/gburd/ra/blob/v0.6.0/crates/ra-pg-extension/sql/pg_ra_planner--0.1.0.sql)

`pg_ra_planner` is the experimental PostgreSQL planner-hook integration for the Ra relational-algebra optimizer. The repository tag is v0.6.0, while the extension control and SQL version are 0.1.0. It is a research preview and the upstream README records known wrong-answer defects; do not use it for production correctness.

### Enable the preview

```sql
CREATE EXTENSION pg_ra_planner;
```

The library must be loaded through `shared_preload_libraries` before the server starts, so changing it requires a restart. The versioned bootstrap SQL checks that planner GUCs are registered; the Rust crate also exposes diagnostic functions. The source manifest uses developer-local pgrx paths, so a reproducible portable build has not been established here. The `parser_hook` feature additionally needs a PostgreSQL build carrying the upstream hook patch.

### Planner controls

Cargo feature flags cover PostgreSQL 13–19; this is a source feature list, not an independently verified support matrix.

The SQL script checks that the library registered its settings. The documented controls include `ra_planner.enabled`, `ra_planner.min_confidence`, `ra_planner.log_decisions` and `ra_planner.max_relations`.

```sql
SET ra_planner.enabled = on;
SET ra_planner.min_confidence = 0.9;
SET ra_planner.log_decisions = on;
```

Run the upstream verification corpus against a disposable cluster before any experiment. Keep a native PostgreSQL planner fallback available, review plans and results, and treat the optimizer's learned model and plan output as experimental.
