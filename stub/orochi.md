## Usage

Sources:

- [Upstream README](https://github.com/howlerops/orochi-db/blob/38f0228648b344829ce62b0a27286fe1067d66f2/README.md)
- [Extension control file](https://github.com/howlerops/orochi-db/blob/38f0228648b344829ce62b0a27286fe1067d66f2/extensions/postgres/orochi.control)
- [Version 1.0 SQL surface](https://github.com/howlerops/orochi-db/blob/38f0228648b344829ce62b0a27286fe1067d66f2/extensions/postgres/sql/orochi--1.0.sql)
- [Beta testing plan](https://github.com/howlerops/orochi-db/blob/38f0228648b344829ce62b0a27286fe1067d66f2/docs/BETA_TESTING_PLAN.md)

`orochi` version `1.0` is a broad PostgreSQL 16+ HTAP project covering custom vectors, distributed tables, hypertables, columnar/compressed storage, tiering, continuous aggregates, approximate functions, pipelines/CDC, workload controls, JSON processing, authentication, and Raft coordination. It installs a large SQL/C surface in the fixed `orochi` schema and integrates through hooks and background components.

### Start with inventory

```sql
CREATE EXTENSION orochi;
SELECT orochi.orochi_version();
SELECT extversion FROM pg_extension WHERE extname = 'orochi';
SELECT proname FROM pg_proc
WHERE pronamespace = 'orochi'::regnamespace ORDER BY proname;
```

Do not enable sharding, tiering, pipelines, authentication hooks, or distributed DDL based only on the README. The repository contains ambitious production-readiness claims alongside a beta plan with core extension checks still listed for validation, and examples/names have drifted between documents and SQL. Installation is superuser-only and some features contact PostgreSQL nodes, S3, Kafka, or web services. Evaluate one subsystem at a time on disposable data, audit every external credential and generated object, and prove crash, upgrade, transaction, failover, and rollback behavior before production use.
