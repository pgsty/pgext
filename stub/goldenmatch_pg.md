## Usage

Sources:

- [PostgreSQL extension README](https://github.com/benseverndev-oss/goldenmatch/blob/ec6ea0e4834fdd478961ff5808be2facc43740e2/packages/rust/extensions/postgres/README.md)
- [Extension control file](https://github.com/benseverndev-oss/goldenmatch/blob/ec6ea0e4834fdd478961ff5808be2facc43740e2/packages/rust/extensions/postgres/goldenmatch_pg.control)
- [Cargo package metadata](https://github.com/benseverndev-oss/goldenmatch/blob/ec6ea0e4834fdd478961ff5808be2facc43740e2/packages/rust/extensions/postgres/Cargo.toml)
- [PostgreSQL bridge source](https://github.com/benseverndev-oss/goldenmatch/tree/ec6ea0e4834fdd478961ff5808be2facc43740e2/packages/rust/extensions/postgres/src)

`goldenmatch_pg` version `0.13.0` exposes GoldenMatch entity-resolution, deduplication, record matching, profiling, validation, correction, and pipeline helpers in the fixed `goldenmatch` schema. Some scalar and graph kernels are native Rust; broader operations embed CPython and call the server-installed `goldenmatch` package.

### Safe scalar start

```sql
CREATE EXTENSION goldenmatch_pg;
SELECT goldenmatch.goldenmatch_score(
  'John Smith', 'Jon Smyth', 'jaro_winkler'
);
```

The documented runtime requires PostgreSQL 15–17, Python 3.11+, and a compatible `goldenmatch` Python package in the interpreter linked into PostgreSQL. Table/job functions read named relations and create or drop result state; identity/correction APIs can accept server-side database or memory paths. Matching scores are heuristic, not identity proof. Restrict operational and file-backed functions, audit dynamic identifiers and generated objects, pin both Rust and Python dependencies, and validate thresholds against labeled data with human review.
