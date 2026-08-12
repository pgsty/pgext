---
title: TODO / Roadmap
description: Current extension packaging backlog, candidates, maintenance risks, and retired projects.
weight: 700
---

This page reflects the local `pgext` catalog and package matrix as of 2026-08-12. A version shown in the catalog is not considered delivered until the corresponding RPM/DEB packages are built and indexed.

## Current Package Work

### Pigsty-maintained upgrades

| Package                 | Current | Target | Scope          | Note                                                                                   |
|:------------------------|:--------|:-------|:---------------|:---------------------------------------------------------------------------------------|
| `pg_profile`            | 4.11    | 4.15   | DEB            | PGDG RPM is already 4.15; Pigsty DEB and source are still 4.11.                        |
| `pg_readonly`           | 1.0.5   | 1.0.6  | DEB            | PGDG RPM is already 1.0.6.                                                             |
| `pg_statement_rollback` | 1.5     | 1.6    | DEB            | PGDG RPM is already 1.6.                                                               |
| `pgsodium`              | 3.1.9   | 3.1.11 | Pigsty RPM/DEB | PGDG has 3.1.11 on EL10, while Pigsty and older PGDG targets remain on 3.1.9.          |
| `topn`                  | 2.7.0   | 2.7.1  | DEB            | Pigsty DEB/source remain on 2.7.0; PGDG RPM coverage is mixed between 2.7.0 and 2.7.1. |

Do not update package metadata merely to close these rows: first build, test, index, and rescan the actual artifacts.

### Package-matrix gap

- [`pg_statviz`](https://github.com/vyruss/pg_statviz) is the only package family currently containing `MISS` cells: 20 of 80 active PG/OS targets. The gaps are RPM PG17 on EL8/9/10, RPM PG18 on EL8/9, and DEB PG14-18 on Ubuntu 22.04, all on both architectures. Upstream/control is 1.1, PGDG DEB is 1.1, and PGDG RPM is still 0.9. Decide whether Pigsty should fill these cells or explicitly classify unsupported targets as `N/A`.

### External repository skew

These differences are visible in the current matrix but are maintained by PGDG rather than being automatic Pigsty rebuild tasks:

| Package            | RPM   | DEB   |
|:-------------------|:------|:------|
| `credcheck`        | 4.7   | 5.0   |
| `decoderbufs`      | 3.5.0 | 3.6.0 |
| `pg_permissions`   | 1.4.1 | 1.4   |
| `pg_stat_kcache`   | 2.3.1 | 2.3.2 |
| `pgauditlogtofile` | 1.8.4 | 1.8.5 |
| `postgis`          | 3.6.3 | 3.6.4 |
| `powa`             | 5.1.0 | 5.2.0 |

## Packaging Candidates

All candidates below already exist in `pgext.universe` but are not in the packaged baseline.

### Worth evaluating

- [`pgedge_vectorizer` 1.1](https://github.com/pgEdge/pgedge-vectorizer): asynchronous text chunking and embedding generation; C background worker, requires preload and pgvector. The current 1.1 tag is still named `v1.1-test1`.
- [`synchdb`](https://github.com/Hornetlabs/synchdb): direct CDC from MySQL, SQL Server, and Oracle; upstream release is 1.4 while the control version remains 1.0, with a large C/Java runtime boundary.
- [`pg_onnx`](https://github.com/kibae/pg_onnx): ONNX inference in PostgreSQL; extension version is 1.2.1 while the project release is 1.28.0, with a large C++ runtime dependency.
- [`pg_deltax` 0.2.1](https://github.com/xataio/deltax): active Rust time-series extension; confirm PostgreSQL-major and pgrx support before packaging.
- [`steampipe_postgres_fdw` 1.0](https://github.com/turbot/steampipe-postgres-fdw): zero-ETL access to cloud services and APIs; review its Go runtime and plugin distribution boundary.
- [`pg_mustach`](https://github.com/RekGRpth/pg_mustach): small C Mustache implementation; its latest tag is `v1.0.0` while the control default is already 3.0, so the release boundary must be clarified first.
- [`is_jsonb_valid` 0.1.4](https://github.com/furstenheim/is_jsonb_valid): actively maintained C implementation of JSON Schema draft 4/7 validation; evaluate overlap with existing JSON Schema extensions.
- [`oai_fdw` 1.13](https://github.com/jimjonesbr/oai_fdw): active OAI-PMH FDW with a narrow academic-metadata use case.
- [`pgjwt_rs` 0.1.2](https://github.com/vishvish/pgjwt): Rust JWT verification for RS256 and Ed25519; evaluate overlap with the existing JWT/security set.

### Icebox / needs review

- [`coldfront`](https://github.com/pgEdge/coldfront): public beta for PG16-18; requires preload, `pg_duckdb`, patched DuckDB/Iceberg components, and auxiliary services. Not production-ready.
- [`ruvector`](https://github.com/ruvnet/RuVector): broad and fast-moving Rust/vector monorepo; the catalog says 0.3.0, the current PostgreSQL crate says 2.0.6, and repository tags are already in the 2.2 series.
- [`pg_deeplake`](https://github.com/activeloopai/deeplake) and [`vexdb_lite`](https://github.com/VexDB-THU/VexDB-Lite): interesting vector projects, but the PostgreSQL packaging boundary and runtime dependencies need a separate review.
- [`plrust`](https://github.com/pgcentralfoundation/plrust): substantial compiler/sandbox toolchain; project release is 1.2.8 while the control default remains 1.1, and upstream only declares PG13-16 features.
- [`pg_query_state`](https://github.com/postgrespro/pg_query_state): requires two matching PostgreSQL core patches, so it is not a normal extension-package candidate.
- `pg_conda`, `pgfdb`, `postgres_ical`, `pgfaker`, `pgsloth`, `pg_kafka`, `pgspeck`, `dsef`, `pg_fsql`, `pg_liquid`, and `pg_regresql`: retained in Universe for later review; low priority, experimental, narrow, or lacking a clean current release boundary.

## Status Corrections

### Completed since the previous list

- `re2` 0.4.1, `spock` 5.0.10, `pg_lake` 3.4, and the `omnigres` package family are indexed with no `MISS` cells in their supported matrices.
- `age` 1.8.0, `pg_jieba` 2.0.1, `onesparse` 1.0.0, `pgelog` 1.0.2, `rdf_fdw` 2.7.0, `pg_ttl_index` 3.0.0, and `pgcalendar` 1.1.0 are packaged.
- `pg_statviz` is now cataloged from PGDG; only the explicit matrix gaps above remain.

### Still packaged, not retired

`pg_search`, `pg_net`, `pg_tle`, `pg_bigm`, `http`, `gzip`, `pg_dirtyread`, `pointcloud`, `pg_proctab`, `pgdd`, `pgx_ulid`, `hashtypes`, and `pghydro` remain in the active packaged catalog. Their previous placement under “Retired” or “Not Planned” was obsolete.

### Maintenance watch

- [`columnar` 1.1.2](https://github.com/hydradatabase/columnar) remains packaged for PG14-16, but upstream has had no commit since 2025-02-10 and there is no PG17/18 support in the current catalog.
- Apache AGE is active and packaged; it is no longer classified as lacking maintenance.

## Not Planned or Retired

- [`timescale/pgai`](https://github.com/timescale/pgai) is archived. [`river`](https://github.com/riverqueue/river) is a Go job library, not a PostgreSQL extension.
- `pg_bm25` was superseded by `pg_search`; `pg_analytics` is archived; `pg_lakehouse` and `embedding` are deprecated; `pg_sparse` was folded into pgvector.
- PipelineDB is deprecated. `sql_firewall`, `zcurve`, and `pg_comparator` are abandoned; `weighted_mean` and `pg_paxos` are archived.
- `pg_lz4` and `pg_query_state` require patched PostgreSQL cores. `vacuumlo`, `oid2name`, and `pg_top` are command-line programs rather than extension packages.
- Legacy projects with no current packaging plan include `zson`, `pg_natural_sort_order`, `pgsampler`, `pg_amqp`, `tinyint`, `pg_blkchain`, `foreign_table_exposer`, `ldap_fdw`, `pg_backtrace`, `connection_limits`, `fixeddecimal`, `fuzzywuzzy`, `pg_scws`, `pg_themis`, `lsm3`, `monq`, `pg_recall`, and `kmeans`.
- [`jsonb_apply` 0.1.0](https://github.com/Florents-Tselai/jsonb_apply) remains blocked because the upstream repository does not declare a license.

## One-sided Package Coverage

The old EL-only and Debian-only lists were stale. The current catalog has only these one-sided package families:

### RPM only

- `db2_fdw` 18.2.0
- `informix_fdw` 0.6.3 — PGDG non-free; requires the IBM Informix Client SDK
- `pg_strom` 6.1 — GPU/NVMe extension; no DEB package

### DEB only

- `debversion` 1.2.0

## Resources

- [PGXN recent releases](https://pgxn.org/recent/)
- [PGDG RPM packaging](https://git.postgresql.org/gitweb/?p=pgrpms.git;a=summary)
- [PGDG Debian packaging](https://salsa.debian.org/postgresql)
- [1000+ PostgreSQL extensions](https://gist.github.com/joelonsql/e5aa27f8cc9bd22b8999b7de8aee9d47)
- [PostgreSQL Extension Network](https://www.pgextensions.org/)
