---
title: DEB Changelog
description: DEB Extension Changelog
weight: 300
---

Check [PGSQL Repo](/repo/pgsql) to learn how to use the PGSQL APT repo.

## 2026-07-21

This batch consolidates extension changes across the RPM and DEB repositories since 2026-07-07: 26 new extensions, 50 Rust/pgrx upgrades or rebuilds, 33 other version upgrades, and 18 package-side or support-matrix changes. Old and New refer to the logical extension version; package-version differences are noted explicitly.

| Type | Extension | Old | New | Comment |
|:-----|:----------|:----|:----|:--------|
| New | argm | - | 1.1.1 | New; PG 14-18; adds the PostgreSQL 16+ varlena header compatibility fix. |
| New | cron_utils | - | 0.1.0 | New SQL-only extension; PG 14-18; PGXN marks 0.1.0 as unstable. |
| New | fbsql | - | 0.1.0 | New; PG 16-18; requires PL/R >= 8.4.0. |
| New | oidc_validator | - | 0.1.0 | New Rust OIDC module; PG 18; upstream pgrx 0.16.1 patched to 0.19.1; 16 platforms; upstream provides no license grant. |
| New | online_advisor | - | 1.0 | New; PG 14-18; backports the PG18 hook fix; PG14-16 require preload. |
| New | pg_cjk_parser | - | 0.1.0 | New; PG 14-18; adds PG_CONFIG build selection and dependency fixes. |
| New | pg_extension_base | - | 3.4 | New pg_lake 3.4.0 component; PG 16-18; requires preload. |
| New | pg_extension_updater | - | 3.4 | New optional pg_lake 3.4.0 component; PG 16-18. |
| New | pg_fts | - | 0.2.0 | New; PG 17-18; trusted and relocatable; RPM also ships an llvmjit subpackage. |
| New | pg_jieba | - | 1.1.0 | New; package 2.0.1, extension 1.1.0; PG 14-18; fixes LexDescr allocation and CMake 4 builds. |
| New | pg_kpart | - | 1.0 | New; PG 14-18; planner hook requires shared/session preload; 16 platforms. |
| New | pg_lake | - | 3.4 | New; package 3.4.0, extension 3.4; PG 16-18; RPM supports EL9/10 only, DEB covers all Debian/Ubuntu targets. |
| New | pg_lake_copy | - | 3.4 | New pg_lake 3.4.0 component; PG 16-18; RPM supports EL9/10 only. |
| New | pg_lake_engine | - | 3.4 | New pg_lake query-engine component; PG 16-18; delegated execution requires pgduck_server. |
| New | pg_lake_iceberg | - | 3.4 | New pg_lake component; PG 16-18; requires pg_lake_engine. |
| New | pg_lake_table | - | 3.4 | New pg_lake component; PG 16-18; loaded in order by pg_extension_base. |
| New | pg_map | - | 3.4 | New pg_lake component; PG 16-18; not the unrelated array-mapping project. |
| New | pg_oidc_validator | - | 0.2 | New Percona C++ OIDC module; PG 18; GCC 11/12 patch; DEB fully covered, RPM EL10 only after EL8/9 libstdc++ ABI failures. |
| New | pg_roast | - | 1.0 | New; PG 14-18; pins main commit ccbf012; periodic worker requires preload. |
| New | pg_tiktoken_c | - | 1.1 | New; PG 14-18; pins commit fa2957b; DESTDIR/correctness patches; pinned snapshot lacks LICENSE. |
| New | pgfr_analyze | - | 2.29.2 | New secondary pg_flight_recorder 2.29.2 extension; PG 15-18; requires pgfr_record. |
| New | pgfr_record | - | 2.29.2 | New primary pg_flight_recorder extension; PG 15-18, no PG14; normalizes the control version and defers scheduling until commit. |
| New | pgmemento | - | 0.7.4 | New SQL-only extension; PG 14-18; includes upgrade scripts from 0.7-0.7.3 to 0.7.4. |
| New | pgmonitor | - | 2.2.0 | New; PG 14-18; metrics need no preload, optional worker requires pgmonitor_bgw preload. |
| New | pgsqlmock | - | 1.0.1 | New; PG 14-18; fixes the control dependency name from pgTap to pgtap. |
| New | plx | - | 1.3.1 | New; PG 14-18; validated on 16 platforms; uses the built-in PL/pgSQL handler. |
| Rust | anon | 3.1.1 | 3.1.3 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | block_copy_command | 0.1.5 | 0.1.5 | requires preload; pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | convert | 0.1.0 | 0.1.0 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | etcd_fdw | 0.0.1 | 0.0.1 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | explain_ui | 0.0.2 | 0.0.2 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | graph | 0.1.7 | 0.1.8 | package name pggraph; pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | jsonschema | 0.1.9 | 0.1.9 | distinct from Supabase pg_jsonschema; pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_base58 | 0.0.1 | 0.0.1 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_bestmatch | 0.0.2 | 0.0.2 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_cardano | 1.2.0 | 1.2.0 | pgrx 0.18.1 -> 0.19.1; PG 15-18. |
| Rust | pg_command_fw | 0.1.0 | 0.1.0 | hook activation requires preload; pgrx 0.18.1 -> 0.19.1; PG 15-18. |
| Rust | pg_durable | 0.2.2 | 0.2.3 | requires preload and a superuser worker role; pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_enigma | 0.5.0 | 0.5.0 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_eviltransform | 0.0.2 | 0.0.4 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_graphql | 1.6.1 | 1.6.1 | not an official upstream release; pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_idkit | 0.4.0 | 0.4.0 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_jsonschema | 0.3.4 | 0.3.4 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_kazsearch | 2.2.0 | 2.3.0 | upstream uses pgrx 0.17.0, Pigsty patches to 0.19.1; pgrx 0.18.1 -> 0.19.1; PG 16-18. |
| Rust | pg_later | 0.4.0 | 0.4.0 | also includes a runtime compatibility fix; pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_mooncake | 0.2.0 | 0.2.0 | 0.2.0 is not an official release; pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_parquet | 0.5.1 | 0.5.1 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_pinyin | 0.0.4 | 0.0.5 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_polyline | 0.0.1 | 0.0.1 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_render | 0.1.3 | 0.1.3 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_rrf | 0.0.3 | 0.0.3 | adds PG18; pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_search | 0.24.0 | 0.24.3 | pins the builder Rust toolchain; pgrx 0.18.1 -> 0.19.1; PG 15-18. |
| Rust | pg_session_jwt | 0.5.0 | 0.5.0 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_smtp_client | 0.2.1 | 0.2.1 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_strict | 1.0.5 | 1.0.5 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_summarize | 0.0.1 | 0.0.1 | includes a PG18 compatibility fix; pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_tiktoken | 0.0.1 | 0.0.1 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_tokenizer | 0.1.1 | 0.1.1 | includes a PG18 compatibility fix; pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pg_trickle | 0.81.0 | 0.81.0 | PG18 only; preserves the pgrx schema linker section; pgrx 0.18.1 -> 0.19.1; PG 18. |
| Rust | pg_when | 0.1.9 | 0.1.9 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pgdd | 0.6.1 | 0.6.1 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pglinter | 2.0.0 | 2.0.0 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pglite_fusion | 0.0.6 | 0.0.6 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pgmqtt | 0.3.0 | 0.4.1 | CDC requires wal_level=logical; pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pgrdf | 0.6.4 | 0.6.20 | adds PG18; preload recommended; pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pgsmcrypto | 0.1.1 | 0.1.1 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | pgx_ulid | 0.2.3 | 0.2.3 | only gen_monotonic_ulid() requires preload; pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | plprql | 18.0.1 | 18.0.1 | also includes a runtime compatibility fix; pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | timescaledb_toolkit | 1.23.0 | 1.23.0 | pgrx 0.18.1 -> 0.19.1; PG 15-18. |
| Rust | typeid | 0.3.0 | 0.3.0 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | tzf | 0.3.0 | 0.3.0 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | vchord | 1.1.1 | 1.1.1 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | vchord_bm25 | 0.3.0 | 0.3.0 | bm25 AM conflicts with pg_search/pg_textsearch; pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | vectorize | 0.26.2 | 0.26.2 | version was incorrectly changed to 0.23.0 and reverted; pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | vectorscale | 0.9.0 | 0.9.0 | pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Rust | wrappers | 0.6.1 | 0.6.2 | intermediate DuckDB FDW feature change was reverted; pgrx 0.18.1 -> 0.19.1; PG 14-18. |
| Upgrade | babelfishpg_tsql | 5.5.0 | 5.4.0 | Downward correction; expands from PG17 to PG17-18; WiltonDB kernel only. |
| Upgrade | biscuit | 2.4.1 | 2.4.3 | PG 16-18; package 2.4.3, but control/SQL default remains 2.4.1. |
| Upgrade | decoderbufs | 3.5.0 | 3.6.0 | PG 14-18; PGDG RPM/DEB. |
| Upgrade | documentdb | 0.113 | 0.114 | Bundled with three subextensions; PG 15-18; 16 platforms; final flow drops temporary export/cache patches. |
| Upgrade | documentdb_core | 0.113 | 0.114 | Bundled documentdb extension; PG 15-18; 16 platforms. |
| Upgrade | documentdb_distributed | 0.113 | 0.114 | Bundled documentdb extension; PG 15-18; 16 platforms. |
| Upgrade | documentdb_extended_rum | 0.113 | 0.114 | Bundled documentdb extension; PG 15-18; 16 platforms. |
| Upgrade | http | 1.7.1 | 1.7.2 | PG 14-18; PGDG RPM/DEB. |
| Upgrade | jdbc_fdw | 0.4.0 | 0.5.0 | Source/package 0.5.0, SQL extension 1.2; RPM PG14-16 with no EL aarch64; DEB PG14-18 with PG18 patch. |
| Upgrade | nominatim_fdw | 1.3 | 2.0.0 | PG 14-18; Pigsty RPM/DEB aligned. |
| Upgrade | odbc_fdw | 0.5.1 | 0.6.1 | Source/package 0.6.1, SQL extension 0.5.2; PG 14-18; PGDG RPM, Pigsty DEB. |
| Upgrade | ogr_fdw | 1.1.8 | 1.1.9 | PG 14-18; PGDG RPM/DEB. |
| Upgrade | pg_dbms_errlog | 2.2 | 2.4 | PG 14-18; requires pg_statement_rollback and preload/restart. |
| Upgrade | pg_ivm | 1.14 | 1.15 | PG 14-18. |
| Upgrade | pg_statement_rollback | 1.5 | 1.6 | PG 14-18; PGDG RPM, Pigsty DEB. |
| Upgrade | pg_tde | 2.1 | 2.2 | PG 17-18 only; Percona TDE kernel. |
| Upgrade | pgauditlogtofile | 1.8.4 | 1.8.5 | PG 14-18; PGDG RPM/DEB. |
| Upgrade | pgbson | 2.0.2 | 2.0.4 | Package 2.0.4 still installs SQL extension 2.0; RPM family postgresbson; requires libbson. |
| Upgrade | pgclone | 4.3.2 | 4.4.2 | PG 14-18; retains RPM LLVM_BINPATH fix; async/progress features require preload. |
| Upgrade | pgextwlist | 1.19 | 1.20 | PG 14-18; PGDG RPM/DEB. |
| Upgrade | pgmnemo | 0.12.1 | 0.13.0 | Support narrowed to PG 17-18; SQL-only; requires pgvector >= 0.7.0. |
| Upgrade | pgmq | 1.11.1 | 1.12.0 | PG 14-18; SQL-only. |
| Upgrade | pgsentinel | 1.4.1 | 1.4.2 | PG 14-18; this batch has RPM 1.4.2 only, while DEB/control remains 1.4.1; explicit RPM-only exception. |
| Upgrade | plpgsql_check | 2.9.2 | 2.10.1 | PG 14-18; Pigsty RPM/DEB aligned. |
| Upgrade | plproxy | 2.11.0 | 2.12.0 | PG 14-18; Pigsty RPM and PGDG DEB aligned. |
| Upgrade | powa | 5.1.2 | 5.2.0 | PG 14-18; PGDG DEB is 5.2.0 while RPM remains 5.1.0. |
| Upgrade | provsql | 1.10.0 | 1.11.0 | PG 14-18. |
| Upgrade | re2 | 0.3.0 | 0.4.1 | PG 16-18; Pigsty RPM/DEB. |
| Upgrade | snowflake | 2.4 | 2.5.0 | Expands from PG18 to PG15-18; pgEdge kernel only. |
| Upgrade | spock | 5.0.6 | 5.0.10 | Expands from PG18 to PG15-18; pgEdge kernel only. |
| Upgrade | tdigest | 1.4.3 | 1.4.4 | PG 14-18; PGDG RPM/DEB. |
| Upgrade | timescaledb | 2.28.2 | 2.28.3 | PG 15-18; timescaledb-tsl; 16 platforms. |
| Upgrade | vector | 0.8.4 | 0.8.5 | PG 14-18; PGDG RPM and Pigsty DEB aligned. |
| Package | babelfishpg_money | 1.1.0 | 1.1.0 | Version unchanged; expands from PG17 to PG17-18; WiltonDB kernel only. |
| Package | babelfishpg_tds | 1.0.0 | 1.0.0 | Version unchanged; expands from PG17 to PG17-18; WiltonDB kernel only. |
| Package | dbt2 | 0.61.7 | 0.61.7 | Adds Pigsty DEB; PG 14-18; SQL extension version remains 0.45.0. |
| Package | faker | 0.5.3 | 0.5.3 | Adds Pigsty DEB; PG 14-18; includes PG17+ parser patch. |
| Package | gb18030_2022 | 1.0 | 1.0 | IvorySQL contrib package 5.1 -> 5.4; IvorySQL 18 only; 16 platforms. |
| Package | hdfs_fdw | 2.3.3 | 2.3.3 | Adds Pigsty DEB; PG 14-18; SQL extension version 2.0.5. |
| Package | ivorysql_ora | 1.0 | 1.0 | IvorySQL contrib package 5.1 -> 5.4; IvorySQL 18 only; 16 platforms. |
| Package | mobilitydb | 1.3.0 | 1.3.0 | Adds Pigsty RPM 1.3.0; PG 14-18 on EL8/9/10, both architectures; DEB already at 1.3.0. |
| Package | mobilitydb_datagen | 1.3.0 | 1.3.0 | Adds RPM with the mobilitydb package; PG 14-18. |
| Package | ora_btree_gin | 1.0 | 1.0 | IvorySQL contrib package 5.1 -> 5.4; IvorySQL 18 only; 16 platforms. |
| Package | ora_btree_gist | 1.0 | 1.0 | IvorySQL contrib package 5.1 -> 5.4; IvorySQL 18 only; 16 platforms. |
| Package | pg_dbms_job | 2.0 | 2.0 | Adds Pigsty DEB; PG 14-18; worker requires preload/restart. |
| Package | pg_dbms_lock | 2.0 | 2.0 | Adds Pigsty DEB; PG 14-18; includes pg_background v2 API compatibility. |
| Package | pg_dbms_metadata | 1.0.0 | 1.0.0 | Adds Pigsty DEB; PG 14-18; PGDG RPM lacks PG15 on EL8 aarch64. |
| Package | pg_fact_loader | 2.0.1 | 2.0.1 | Version unchanged; DEB adds PG18 and fills PG14-18 gaps for Ubuntu 26.04. |
| Package | pg_get_functiondef | 1.0 | 1.0 | IvorySQL contrib package 5.1 -> 5.4; IvorySQL 18 only; 16 platforms. |
| Package | pgbouncer_fdw | 1.4.0 | 1.4.0 | Adds Pigsty DEB; PG 14-18; runtime requires a PgBouncer admin console. |
| Package | plisql | 1.0 | 1.0 | IvorySQL contrib package 5.1 -> 5.4; IvorySQL 18 only; 16 platforms. |

## 2026-07-07

| Package           | Old        | New        | Comment       |
|:------------------|:-----------|:-----------|:--------------|
| polardb-17        | 17.9.1.0   | 17.10.1.0  | PG 17         |
| agensgraph-17     | 2.16.0     | 2.17.0     | PG 17.10      |
| openhalodb-14     | 1.0-beta   | 1.0-2      | OpenHaloDB    |
| babelfish-17      | 5.4.0      | 5.4.0      | PG 17.7       |
| babelfish-18      | -          | 6.0.0      | PG 18.3       |
| pgedge-15         | -          | 15.18      | Spock 5.0.10  |
| pgedge-16         | -          | 16.14      | Spock 5.0.10  |
| pgedge-17         | 17.9       | 17.10      | Spock 5.0.10  |
| pgedge-18         | 18.3       | 18.4       | Spock 5.0.10  |
| orioledb-16       | 1.8-beta16 | 1.8-beta16 | PG 16         |
| orioledb-17       | 1.8-beta16 | 1.8-beta16 | PG 17         |
| orioledb-18       | 1.8-beta16 | 1.8-beta16 | PG 18         |
| ivorysql-18       | 5.0        | 5.4        | PG 18         |
| cloudberry        | 2.1.0-1    | 2.1.0-2    | rebuild       |
| cloudberry-backup | 2.1.0-1    | 2.1.0-2    | backup        |
| cloudberry-pxf    | 2.1.0-1    | 2.1.0-2    | PXF           |

## 2026-07-01

| Package       | Old      | New     | Comment                               |
|:--------------|:---------|:--------|:--------------------------------------|
| timescaledb   | 2.28.0   | 2.28.2  | PG 15-18                              |
| documentdb    | 0.112    | 0.113   | PG 15-18                              |
| citus         | 14.0.0-4 | 14.1.0  | PG 16-18                              |
| pgvector      | 0.8.3    | 0.8.4   | PG 14-18                              |
| plpgsql_check | 2.9.1    | 2.9.2   | PG 14-18                              |
| roaringbitmap | 1.1.0    | 1.2.0-2 | PG 14-18, llvm-lto packaging fix      |
| biscuit       | 2.3.0    | 2.4.1   | PG 16-18                              |
| pgmnemo       | 0.8.3    | 0.12.1  | PG 14-18                              |
| pg_ivm        | 1.14     | 1.15    | PG 14-18                              |
| rdf_fdw       | 2.5.0    | 2.6.0   | PG 14-18, libcurl compatibility patch |

## 2026-06-18

Updated Rust extension pgrx version to 0.18.1.

| Package             | Old        | New        | Comment                  |
|:--------------------|:-----------|:-----------|:-------------------------|
| timescaledb         | 2.27.2     | 2.28.0     | PG 15-18                 |
| timeseries          | 0.2.0      | 0.2.1      | PG 14-18                 |
| documentdb          | 0.110      | 0.112      | PG 15-18                 |
| pg_clickhouse       | 0.3.1      | 0.3.2      | PG 14-18                 |
| biscuit             | 2.2.2      | 2.3.0      | PG 16-18                 |
| pg_http             | 1.7.0      | 1.7.1      | PG 14-18                 |
| pg_gzip             | 1.0.0      | 1.1.0      | PG 14-18                 |
| pgvector            | 0.8.2      | 0.8.3      | PG 14-18                 |
| age                 | 1.7.0      | 1.7.0      | PG 17-18                 |
| mongo_fdw           | -          | 5.5.3      | new, PG 14-18            |
| orioledb            | 1.7-beta15 | 1.8-beta16 | Build for PG 16, 17, 18  |
| etcd_fdw            | 0.0.0      | 0.0.1      | PG 14-18, pgrx 0.18.1    |
| pg_anon             | 3.0.13     | 3.1.1      | PG 14-18, pgrx 0.18.1    |
| pg_graphql          | 1.5.12     | 1.6.1      | PG 14-18, pgrx 0.18.1    |
| pg_kazsearch        | 2.0.0      | 2.2.0      | PG 16-18, pgrx 0.18.1    |
| pg_session_jwt      | 0.4.0      | 0.5.0      | PG 14-18, pgrx 0.18.1    |
| pg_tzf              | 0.2.4      | 0.3.0      | PG 14-18, pgrx 0.18.1    |
| pg_vectorize        | 0.26.1     | 0.26.2     | PG 14-18, pgrx 0.18.1    |
| pglinter            | 1.1.2      | 2.0.0      | PG 14-18, pgrx 0.18.1    |
| pgmqtt              | 0.1.0      | 0.3.0      | PG 14-18, pgrx 0.18.1    |
| timescaledb_toolkit | 1.22.0     | 1.23.0     | PG 15-18, pgrx 0.18.1    |
| wrappers            | 0.6.0      | 0.6.1      | PG 14-18, pgrx 0.18.1    |
| pgrdf               | 0.5.0      | 0.6.4      | PG 14-17, pgrx 0.18.1    |
| pg_ducklake         | -          | 1.0.0      | new, PG 14-18            |
| pg_mockable         | -          | 1.1.0      | new, PG 14-18            |
| pg_stat_backtrace   | -          | 1.0.0      | new, PG 14-18, libunwind |
| multicorn           | -          | 3.2        | new, PG 14-18            |
| passwordpolicy      | -          | 2.0.5      | new, PG 14-18            |
| pgdisablelogerror   | -          | 1.0        | new, PG 14-18            |
| plpgsql_wrap        | -          | 1.0        | new, PG 14-18            |

## 2026-06-12

Added most Ubuntu 26.04 Resolute core extension packages.

| Package        | Old    | New    | Comment                          |
|:---------------|:-------|:-------|:---------------------------------|
| firebird_fdw   | 1.4.1  | 1.4.2  | PG 14-18                         |
| pg_background  | 1.9.2  | 2.0    | PG 14-18                         |
| pg_clickhouse  | 0.3.0  | 0.3.1  | PG 14-18                         |
| pg_dirtyread   | 2.7    | 2.8    | PG 14-18, RPM only               |
| pg_net         | 0.20.2 | 0.20.3 | PG 14-18, EL10 only              |
| pg_search      | 0.23.1 | 0.24.0 | PG 15-18                         |
| pg_stat_ch     | 0.3.6  | 0.3.6  | PG 16-18, EL9/EL10 only, rebuild |
| pg_trickle     | 0.40.0 | 0.81.0 | PG 18 only                       |
| plpgsql_check  | 2.9.0  | 2.9.1  | PG 14-18                         |
| provsql        | 1.8.0  | 1.9.0  | PG 14-18                         |
| re2            | 0.1.1  | 0.3.0  | PG 16-18                         |
| storage_engine | 2.3.0  | 2.4.0  | PG 15-18                         |
| timescaledb    | 2.27.0 | 2.27.2 | PG 15-18                         |
| pggraph        | 0.1.5  | 0.1.7  | PG 14-18                         |
| pgmnemo        | 0.7.2  | 0.8.3  | PG 14-18                         |
| pgsql_tweaks   | 1.0.2  | 1.0.3  | PG 14-18                         |
| pg_orca        | -      | 1.0.0  | new, PG 18 only                  |
| pg_projection  | -      | 1.0.0  | new, PG 14-18                    |
| pg_sorted_heap | -      | 0.14.0 | new, PG 16-18                    |
| pg_stl         | -      | 1.0.0  | new, PG 16-18                    |
| pg_uuid_v8     | -      | 1.0.0  | new, PG 14-18                    |
| pg_pinyin      | -      | 0.0.4  | new, PG 14-18                    |
| pg_task        | -      | 2.1.29 | new, PG 14-18, pcre2grep fix     |
| pg_extra_time  | -      | 2.1.0  | new, PG 14-18                    |
| fsm_core       | -      | 1.1.0  | new, PG 15-18                    |

## 2026-05-01

| Package          | Old    | New    | Comment                                                                             |
|:-----------------|:-------|:-------|:------------------------------------------------------------------------------------|
| pg_pathcheck     | -      | 0.9.1  | new planner Path diagnostics module, PG 17-18, shared_preload_libraries required    |
| pg_bikram_sambat | -      | 0.1.0  | new Bikram Sambat date type and AD/BS conversions                                   |
| timescaledb      | 2.26.3 | 2.26.4 | TSL minor update                                                                    |
| storage_engine   | 1.2.3  | 1.3.4  | PGXN update, columnar and row-compressed table access methods                       |
| pgproto          | 0.3.3  | 0.5.0  | PGXN update, native Protobuf support, C/PGXS build                                  |
| pg_savior        | 0.0.1  | 0.1.0  | PGXN update, high-risk DDL/DML hook, preload or LOAD required                       |
| pg_textsearch    | 1.0.0  | 1.1.0  | Timescale BM25 full-text search update, PG 17-18, shared_preload_libraries required |
| pg_trickle       | 0.31.0 | 0.40.0 | Rust/pgrx update, PG 18 only                                                        |
| pgedge           | 17.9   | 18.3   | rebuilt for PG 18                                                                   |
| spock            | 5.0.5  | 5.0.6  | rebuilt for PG 18                                                                   |
| lolor            | 1.2.2  | 1.2.2  | rebuilt for PG 18                                                                   |
| snowflake        | 2.4    | 2.4    | rebuilt for PG 18                                                                   |

Built for Debian 12/13 and Ubuntu 22.04/24.04/26.04 on amd64 and arm64.

## 2026-04-25

| Package        | Old    | New    | Comment                       |
|:---------------|:-------|:-------|:------------------------------|
| documentdb     | 0.109  | 0.110  | bump to upstream, PG 15-18    |
| pg_trickle     | 0.20.0 | 0.31.0 | bump, PG 18 only, pgrx 0.18.0 |
| pg_search      | 0.23.0 | 0.23.1 | bump, PG 15-18, pgrx 0.18.0   |
| pg_ivm         | 1.13   | 1.14   | bump, PG 14-18                |
| prefix         | 1.2.10 | 1.2.11 | bump, PG 14-18, PGDG          |
| credcheck      | 4.6    | 4.7    | bump, PG 14-18, PGDG          |
| pg_dbms_job    | 1.5    | 2.0    | bump, PG 14-18, PGDG          |
| storage_engine | 1.0.7  | 1.2.3  | bump, PG 14-18                |
| pgmq           | 1.11.0 | 1.11.1 | bump, PG 14-18                |
| parray_gin     | 1.4.0  | 1.5.0  | bump, PG 14-18                |
| rdf_fdw        | 2.4.0  | 2.5.0  | bump, PG 14-18                |
| pg_accumulator | -      | 1.1.3  | new, PG 14-18                 |

## 2026-04-19

| Package           | Old    | New    | Comment                                 |
|:------------------|:-------|:-------|:----------------------------------------|
| cloudberry        | 2.0.0  | 2.1.0  | bump, split backup and PXF subpackages  |
| cloudberry-backup | -      | 2.1.0  | new companion package                   |
| cloudberry-pxf    | -      | 2.1.0  | new companion package                   |
| oriolepg          | 17.16  | 17.18  | kernel update for orioledb beta15 / 1.7 |
| orioledb          | 1.6    | 1.7    | bump, paired with oriolepg 17.18        |
| timescaledb       | 2.26.2 | 2.26.3 | bump                                    |
| pg_search         | 0.22.6 | 0.23.0 | bump                                    |
| pg_trickle        | 0.17.0 | 0.20.0 | bump, PG 18 only                        |
| pg_clickhouse     | 0.1.10 | 0.2.0  | bump                                    |
| pg_stat_ch        | 0.3.4  | 0.3.6  | bump, PG 16-18                          |
| pgclone           | 3.6.0  | 4.0.0  | bump                                    |
| pgproto           | 0.2.4  | 0.3.3  | bump                                    |
| pgxicor           | 0.1.0  | 0.1.1  | bump                                    |
| storage_engine    | -      | 1.0.7  | new, PG 14-18                           |
| re2               | -      | 0.1.1  | new, PG 16-18                           |
| ulak              | -      | 0.0.2  | new, PG 14-18                           |

## 2026-04-14

| Package            | Old    | New      | Comment                              |
|:-------------------|:-------|:---------|:-------------------------------------|
| block_copy_command | -      | 0.1.5    | new, PG 14-18, pgrx 0.17.0           |
| pg_kazsearch       | -      | 2.0.0    | new, PG 16-18, pgrx 0.17.0           |
| pg_rrf             | -      | 0.0.3    | new, PG 14-17, pgrx 0.16.1 -> 0.17.0 |
| pgmqtt             | -      | 0.1.0    | new, PG 14-18, pgrx 0.16.1 -> 0.17.0 |
| pg_when            | -      | 0.1.9    | new, PG 14-18, pgrx 0.17.0           |
| provsql            | -      | 1.2.3    | new, PG 14-18                        |
| pg_isok            | -      | 1.4.1    | new, PG 14-18                        |
| pg_byteamagic      | -      | 0.2.4    | new, PG 14-18                        |
| logical_ddl        | -      | 0.1.0    | new, PG 14-18                        |
| datasketches       | -      | 1.7.0    | new, PG 14-18                        |
| pg_text_semver     | -      | 1.2.1    | new, PG 14-18                        |
| external_file      | -      | 1.2      | new, PG 14-18                        |
| pg_query_rewrite   | -      | 0.0.5    | new, PG 14-18                        |
| pghydro            | -      | 6.6      | new, PG 14-18                        |
| pg_datasentinel    | -      | 1.0      | new, PG 15-18                        |
| onesparse          | -      | 1.0.0    | new, PG 18 only                      |
| rdkit              | -      | 202503.6 | new, PG 14-18                        |
| pg_dispatch        | -      | 0.1.5    | new, PG 14-18                        |
| pg_fsql            | -      | 1.1.0    | new, PG 14-18                        |
| pg_liquid          | -      | 0.1.7    | new, PG 14-18                        |
| pg_regresql        | -      | 2.0.0    | new, PG 14-18                        |
| pg_slug_gen        | -      | 1.0.0    | new, PG 15-18                        |
| pg_stat_ch         | -      | 0.3.4    | new, PG 16-18                        |
| pg_variables       | -      | 1.2.5    | new, PG 14-18                        |
| pgcalendar         | -      | 1.1.0    | new, PG 14-18                        |
| pgclone            | -      | 3.6.0    | new, PG 14-18                        |
| pgelog             | -      | 1.0.2    | new, PG 14-18                        |
| pglock             | -      | 1.0.0    | new, PG 14-18                        |
| pgproto            | -      | 0.2.4    | new, PG 14-18                        |
| postgresbson       | -      | 2.0.2    | new, PG 14-18                        |
| rdf_fdw            | -      | 2.4.0    | new, PG 14-18                        |
| parray_gin         | -      | 1.4.0    | new, PG 14-18                        |
| timescaledb        | 2.25.2 | 2.26.2   | bump                                 |
| pg_background      | 1.8    | 1.9.2    | bump                                 |
| pg_ivm             | 1.13   | 1.14     | bump                                 |
| system_stats       | 3.2    | 4.0      | bump                                 |
| nominatim_fdw      | 1.1.0  | 1.2      | bump                                 |
| pg_textsearch      | 0.5.0  | 1.0.0    | bump                                 |
| pg_clickhouse      | 0.1.5  | 0.1.10   | bump                                 |
| pg_search          | 0.22.2 | 0.22.6   | bump                                 |
| pg_store_plans     | 1.9    | 1.10     | bump                                 |
| pg_tzf             | 0.2.3  | 0.2.4    | bump, pgrx 0.17.0                    |
| pg_anon            | 3.0.1  | 3.0.13   | bump, pgrx 0.16.1 -> 0.17.0          |
| pg_cardano         | 1.1.1  | 1.2.0    | bump, pgrx 0.17.0                    |
| pg_strict          | 1.0.3  | 1.0.5    | bump, pgrx 0.16.1 -> 0.17.0          |
| pg_vectorize       | 0.26.0 | 0.26.1   | bump, pgrx 0.16.1 -> 0.17.0          |
| pglinter           | 1.1.1  | 1.1.2    | bump, pgrx 0.16.1 -> 0.17.0          |
| pgx_ulid           | 0.2.2  | 0.2.3    | bump, pgrx 0.17.0                    |
| wrappers           | 0.5.7  | 0.6.0    | bump, pgrx 0.16.1 -> 0.17.0          |
| pg_trickle         | 0.16.0 | 0.17.0   | bump, pgrx 0.17.0                    |
| supautils          | 3.1.0  | 3.2.1    | bump                                 |
| ddl_historization  | 0.0.7  | 0.2      | bump                                 |
| pg_incremental     | 1.4.1  | 1.5.0    | bump                                 |
| pg_failover_slots  | 1.2.0  | 1.2.1    | bump                                 |
| PolarDB            | 15.15  | 17.9.1.0 | PG 15 -> 17                          |

## 2026-03-21

| Package            | Old     | New       | Comment |
|:-------------------|:--------|:----------|:--------|
| pg_search          | 0.21.12 | 0.22.2    |         |
| pg_track_optimizer | 0.9.1   | 0.9.2     |         |
| pgcollection       | 1.0.0   | 2.0.0     |         |
| pg_ttl_index       | 2.0.0   | 3.0.0     |         |
| pg_clickhouse      | 0.1.4   | 0.1.5     |         |
| pdu                |         | 3.0.25.12 | new     |
| pgdog              |         | 0.1.32    | new     |

## 2026-03-05

Drop PG 13 support for all extensions

| Package          | Old     | New     | Comment                      |
|:-----------------|:--------|:--------|:-----------------------------|
| aggs_for_vecs    | 1.4.0   | 1.4.1   | upstream bump, PG 14-18      |
| timescaledb      | 2.25.1  | 2.25.2  | upstream bump, PG 15-18      |
| vchord           | 1.1.0   | 1.1.1   | upstream bump, PG 14-18      |
| vchord_bm25      | 0.3.0-1 | 0.3.0-2 | packaging fix, PG 14-18      |
| pg_pinyin        | -       | 0.0.2   | new package, PG 14-18        |
| pg_eviltransform | -       | 0.0.2   | new package, version aligned |
| qos              | -       | 1.0.0   | new package, PG 15-18        |

## 2026-02-27

| Package           | Old      | New      | Comment                              |
|:------------------|:---------|:---------|:-------------------------------------|
| timescaledb       | 2.25.0   | 2.25.1   |                                      |
| citus             | 14.0.0-3 | 14.0.0-4 | rebuilt with latest official release |
| age               | 1.7.0    | 1.7.0    | added PG 17 build for 1.7.0          |
| pg_background     | -        | 1.8      | DEB only, no RPM package             |
| pgmq              | 1.10.0   | 1.10.1   | package unavailable for now          |
| pg_search         | 0.21.6   | 0.21.9   | direct download usage                |
| oriolepg          | 17.11    | 17.16    | OriolePG kernel update               |
| orioledb          | beta12   | beta14   | paired with OriolePG 17.16           |
| openhalo          | 14.10    | 1.0      | updated and renamed, 14.18           |
| pgedge            | -        | 17.9     | new multi-master edge kernel         |
| spock             | -        | 5.0.5    | new, pgEdge core extension           |
| lolor             | -        | 1.2.2    | new, pgEdge core extension           |
| snowflake         | -        | 2.4      | new, pgEdge core extension           |
| babelfishpg       | -        | 5.5.0    | new BabelfishPG package group        |
| babelfish         | -        | 5.5.0    | new Babelfish compatibility package  |
| antlr4-runtime413 | -        | 4.13     | new runtime dependency for Babelfish |

## 2026-02-12

| Package            | OLD        | NEW          | Comment                      |
|--------------------|------------|--------------|------------------------------|
| timescaledb        | 2.24.0     | 2.25.0       |                              |
| pg_incremental     | 1.2.0      | 1.4.1        |                              |
| pg_bigm            | 1.2        | 1.2-20250903 |                              |
| pg_net             | 0.20.0     | 0.20.2       | ubuntu22 libcurl too low     |
| pgmq               | 1.9.0      | 1.10.0       |                              |
| pg_textsearch      | 0.4.0      | 0.5.0        |                              |
| pljs               | 1.0.4      | 1.0.5        |                              |
| sslutils           | 1.4-1      | 1.4-2        |                              |
| supautils          | 3.0.2      | 3.1.0        |                              |
| pg_math            | 1.0        | 1.1.0        |                              |
| pgsentinel         | 1.3.1      | 1.4.0        |                              |
| pg_uri             | 1.20151224 | 1.20251029   |                              |
| pgcollection       | 1.1.0      | 1.1.1        |                              |
| pg_readonly        | 1.0.3      | 1.0.4        |                              |
| timestamp9         | 1.4.0-1    | 1.4.0-2      | rebuild to fix deps          |
| plprql             | 18.0.0     | 18.0.1       |                              |
| pglinter           | 1.0.1      | 1.1.0        |                              |
| pg_jsonschema      | 0.3.3      | 0.3.4        |                              |
| pg_anon            | 2.5.1      | 3.0.1        |                              |
| pg_search          | 0.21.4     | 0.21.6       |                              |
| pg_graphql         | 1.5.12-1   | 1.5.12-2     | switch to official release   |
| pg_summarize       | 0.0.1-2    | 0.0.1-3      | rebuild to fix pg18 issue    |
| nominatim_fdw      |            | 1.1.0        | new, sync with pgdg yum repo |
| pg_utl_smtp        |            | 1.0.0        | new, sync with pgdg yum repo |
| pg_strict          | -          | 1.0.2        | new Rust extension           |
| pg_track_optimizer | -          | 0.9.1        | new extension                |
| pgmb               | -          | 1.0.0        | new extension                |

## 2026-01-25

| Name          | Old            | New            | Comment                 |
|:--------------|:---------------|:---------------|:------------------------|
| age           | 1.6.0          | 1.7.0          | PG 18 only              |
| citus         | 14.0.0-1PIGSTY | 14.0.0-2PIGSTY | official branch release |
| pg_clickhouse | 0.1.2          | 0.1.3          |                         |
| pgmq          | 1.8.1          | 1.9.0          |                         |
| pg_search     | 0.21.2         | 0.21.4         |                         |

## 2026-01-16

| Name                    | Old    | New    | Comment               |
|:------------------------|:-------|:-------|:----------------------|
| etcd_fdw                |        | 0.0.0  | new                   |
| pg_ttl_index            |        | 0.1.0  | new                   |
| citus                   | 13.2.0 | 14.0.0 | +pg18, pre-release    |
| pg_search               | 0.20.5 | 0.21.2 | +pg18                 |
| pg_clickhouse           | 0.1.0  | 0.1.2  |                       |
| pg_textsearch           | 0.1.0  | 0.4.0  |                       |
| pg_convert              | 0.0.5  | 0.1.0  |                       |
| pg_timeseries           | 0.1.8  | 0.2.0  |                       |
| biscuit                 | 2.0.1  | 2.2.2  |                       |
| pgmq                    | 1.8.0  | 1.8.1  |                       |
| documentdb              | 0.107  | 0.109  | +pg18, use ms version |
| pg_bulkload             | 3.1.22 | 3.1.23 | +pg18                 |
| age                     | -      | 1.6.0  | +pg18 use PGDG        |
| pgsentinel              | 1.2.0  | 1.3.1  | use PGDG              |
| pljs                    | -      | 1.0.4  | use PGDG              |
| pg_partman              | 5.3.0  | 5.4.0  | use PGDG              |
| pgfincore               | -      | 1.3.1  | use PGDG              |
| documentdb_extended_rum |        | 0.109  | new                   |
| mobilitydb_datagen      |        | 1.3.0  | new                   |

## 2025-12-25

| Name          | Old     | New     | Comment  |
|:--------------|:--------|:--------|:---------|
| pg_duckdb     | 1.1.0   | 1.1.1   |          |
| pg_search     | 0.20.4  | 0.20.5  |          |
| vchord_bm25   | 0.2.2   | 0.3.0   |          |
| pg_semver     | 0.40.0  | 0.41.0  |          |
| pg_timeseries | 0.1.7   | 0.1.8   |          |
| supautils     | 3.0.2-1 | 3.0.2-2 | fix pg18 |
| pg_summarize  | 0.0.1-1 | 0.0.1-2 | fix pg18 |

## 2025-12-16

| Name                                                          | Old     | New     | Comment          |
|:--------------------------------------------------------------|:--------|:--------|:-----------------|
| [pg_textsearch](https://github.com/timescale/pg_textsearch)   | -       | 0.1.0   | new              |
| [pg_clickhouse](https://github.com/clickhouse/pg_clickhouse/) | -       | 0.1.0   | new              |
| [pg_ai_query](https://github.com/benodiwal/pg_ai_query)       | -       | 0.1.1   | new              |
| timescaledb                                                   | 2.23.1  | 2.24.0  |                  |
| pg_search                                                     | 0.20.0  | 0.20.4  |                  |
| pg_duckdb                                                     | 1.1.0-1 | 1.1.0-2 | official release |
| pg_biscuit                                                    | 1.0     | 2.0.1   | new repo         |
| pg_convert                                                    | 0.0.4   | 0.0.5   | rm legacy pg     |
| pgdd                                                          | 0.6.0   | 0.6.1   | rm legacy pg     |
| pglinter                                                      | 1.0.0   | 1.0.1   |                  |
| pg_session_jwt                                                | 0.3.3   | 0.4.0   |                  |
| pg_anon                                                       | 2.4.1   | 2.5.1   |                  |
| pg_enigma                                                     | 0.4.0   | 0.5.0   |                  |
| wrappers                                                      | 0.5.6   | 0.5.7   |                  |
| pg_vectorize                                                  | 0.25.0  | 0.26.0  | fix pg18         |
| pg_tiktoken                                                   | -       | -       | fix pg18         |
| pg_tzf                                                        | -       | -       | fix pg18         |
| pglite_fusion                                                 | -       | -       | fix pg18         |
| pgsmcrypto                                                    | -       | -       | fix pg18         |
| pgx_ulid                                                      | -       | -       | fix pg18         |
| plprql                                                        | -       | -       | fix pg18         |

## 2025-12-04

| Name    | Old | New | Comment           |
|:--------|:----|:----|:------------------|
| synchdb | -   | 1.3 | Ubuntu 22/24 only |

## 2025-11-20

| Name                   | Old    | New    | Comment                    |
|:-----------------------|:-------|:-------|:---------------------------|
| vchord                 | 0.5.3  | 1.0.0  |                            |
| pg_later               | 0.3.1  | 0.4.0  |                            |
| pgvectorscale          | 0.8.0  | 0.9.0  | -legacy, +pg18             |
| pglite_fusion          | 0.0.5  | 0.0.6  |                            |
| pgx_ulid               | 0.2.1  | 0.2.2  |                            |
| pg_search              | 0.19.5 | 0.20.0 | resume PIGSTY building     |
| citus                  | 13.2.0 | 13.2.0 | official tag               |
| timescaledb            | 2.23.0 | 2.23.1 |                            |
| pg_profile             | 4.10   | 4.11   |                            |
| pglinter               |        | 1.0.0  | new                        |
| pg_typeid              |        | 0.3.0  | head with pg18 support     |
| pg_enigma              |        | 0.4.0  | vonng patched pgrx version |
| pg_retry               |        | 1.0.0  | new, pg17-18               |
| pg_biscuit             |        | 1.0    | new, pg16-18               |
| pg_weighted_statistics |        | 1.0.0  | new, pg14-18               |
| pg_stat_monitor        | 2.3.0  | 2.3.1  |                            |
| pgmq                   | 1.7.0  | 1.8.0  |                            |
| documentdb             | 0.106  | 0.107  | ferretdb fork              |
| PolarDB                |        | 15.15  | 15.15.5.0-38948055         |

## 2025-11-10

Add PostgreSQL 18 support for almost all extensions

| Name                  |    Old     |    New     | Comment |
|:----------------------|:----------:|:----------:|:--------|
| omni_csv              |     -      |   0.1.1    | new     |
| omni_datasets         |     -      |   0.1.0    | new     |
| omni_shmem            |     -      |   0.1.0    | new     |
| pg_csv                |     -      |   1.0.1    | new     |
| pljs                  |     -      |   1.0.3    | new     |
| plxslt                |     -      | 0.20140221 | new     |
| credcheck             |    3.0     |    4.2     | +pg18   |
| dbt2                  |   0.45.0   |   0.61.7   | +pg18   |
| h3                    |   4.1.3    |   4.2.3    | +pg18   |
| h3_postgis            |   4.1.3    |   4.2.3    | +pg18   |
| mongo_fdw             |    1.1     |   5.5.3    | +pg18   |
| multicorn             |    3.0     |    3.2     | +pg18   |
| orafce                |   4.14.4   |   4.14.6   | +pg18   |
| pg_hint_plan          |   1.7.0    |   1.8.0    | +pg18   |
| pg_search             |   0.18.1   |   0.19.2   | +pg18   |
| pg_show_plans         |   2.1.6    |   2.1.7    | +pg18   |
| pgactive              |   2.1.6    |   2.1.7    | +pg18   |
| pgpcre                |     1      | 0.20190509 | +pg18   |
| plpgsql_check         |   2.8.2    |   2.8.3    | +pg18   |
| roaringbitmap         |   0.5.4    |   0.5.5    | +pg18   |
| uint                  | 1.20231206 | 1.20250815 | +pg18   |
| uint128               |   1.1.0    |   1.1.1    | +pg18   |
| anon                  |   2.3.0    |   2.4.1    | +pg18   |
| collection            |   1.0.0    |   1.1.0    | +pg18   |
| emaj                  |   4.7.0    |   4.7.1    | +pg18   |
| explain_ui            |   0.0.1    |   0.0.2    | +pg18   |
| firebird_fdw          |   1.4.0    |   1.4.1    | +pg18   |
| login_hook            |    1.6     |    1.7     | +pg18   |
| logerrors             |   2.1.3    |   2.1.5    | +pg18   |
| mobilitydb            |   1.2.0    |   1.3.0    | +pg18   |
| omni                  |   0.2.9    |   0.2.14   | +pg18   |
| omni_httpc            |   0.1.5    |   0.1.10   | +pg18   |
| omni_httpd            |   0.4.6    |   0.4.11   | +pg18   |
| omni_kube             |   0.1.1    |   0.4.2    | +pg18   |
| omni_sql              |   0.5.1    |   0.5.3    | +pg18   |
| omni_sqlite           |   0.1.2    |   0.2.2    | +pg18   |
| omni_worker           |   0.1.0    |   0.2.1    | +pg18   |
| pg_cardano            |   1.0.5    |   1.1.1    | +pg18   |
| pg_checksums          |    1.2     |    1.3     | +pg18   |
| pg_cron               |   1.6.5    |   1.6.7    | +pg18   |
| pg_duckdb             |   0.3.1    |   1.1.0    | +pg18   |
| pg_failover_slots     |   1.1.0    |   1.2.0    | +pg18   |
| pg_graphql            |   1.5.11   |   1.5.12   | +pg18   |
| pg_idkit              |   0.3.1    |   0.4.0    | +pg18   |
| pg_mooncake           |   0.1.2    |   0.2.0    | +pg18   |
| pg_net                |   0.9.2    |   0.20.0   | +pg18   |
| pg_parquet            |   0.4.3    |   0.5.1    | +pg18   |
| pg_partman            |   5.2.4    |   5.3.0    | +pg18   |
| pg_session_jwt        |   0.3.1    |   0.3.3    | +pg18   |
| pg_sphere             |   1.5.1    |   1.5.2    | +pg18   |
| pg_stat_monitor       |   2.2.0    |   2.3.0    | +pg18   |
| pg_statement_rollback |    1.4     |    1.5     | +pg18   |
| pg_store_plans        |    1.8     |    1.9     | +pg18   |
| pg_task               |   1.0.0    |   2.1.12   | +pg18   |
| pg_tle                |   1.5.1    |   1.5.2    | +pg18   |
| pg_uuidv7             |   1.6.0    |   1.7.0    | +pg18   |
| pglogical             |   2.4.5    |   2.4.6    | +pg18   |
| pgmq                  |   1.5.1    |   1.7.0    | +pg18   |
| pgroonga              |   4.0.0    |   4.0.4    | +pg18   |
| pgsql_tweaks          |   0.11.3   |   1.0.2    | +pg18   |
| pldbgapi              |    1.8     |    1.9     | +pg18   |
| plprql                |   1.0.0    |   18.0.0   | +pg18   |
| supautils             |   2.10.0   |   3.0.2    | +pg18   |
| timescaledb           |   2.22.0   |   2.23.0   | +pg18   |
| timescaledb_toolkit   |   1.21.0   |   1.22.0   | +pg18   |
| vchord                |   0.5.1    |   0.5.3    | +pg18   |
| vectorize             |   0.22.2   |   0.25.0   | +pg18   |
| wrappers              |   0.5.4    |   0.5.6    | +pg18   |
| acl                   |   1.0.4    |     -      | +pg18   |
| aggs_for_arrays       |   1.3.3    |     -      | +pg18   |
| aggs_for_vecs         |   1.4.0    |     -      | +pg18   |
| base36                |   1.0.0    |     -      | +pg18   |
| hashlib               |    1.1     |     -      | +pg18   |
| hll                   |    2.18    |     -      | +pg18   |
| imgsmlr               |    1.0     |     -      | +pg18   |
| index_advisor         |   0.2.0    |     -      | +pg18   |
| kafka_fdw             |   0.0.3    |     -      | +pg18   |
| pg_auth_mon           |    3.0     |     -      | +pg18   |
| pg_background         |    1.3     |     -      | +pg18   |
| pg_bigm               |    1.2     |     -      | +pg18   |
| pg_profile            |    4.10    |     -      | +pg18   |
| pg_stat_kcache        |   2.3.0    |     -      | +pg18   |
| pgdd                  |   0.6.0    |     -      | +pg18   |
| pgjwt                 |   0.2.0    |     -      | +pg18   |
| pgmp                  |   1.0.5    |     -      | +pg18   |
| plprofiler            |   4.2.5    |     -      | +pg18   |
| plv8                  |   3.2.4    |     -      | +pg18   |
| redis_fdw             |    1.0     |     -      | +pg18   |
| repmgr                |   5.5.0    |     -      | +pg18   |
| system_stats          |    3.2     |     -      | +pg18   |
| topn                  |   2.7.0    |     -      | +pg18   |
| zhparser              |    2.3     |     -      | +pg18   |

## 2025-09-06

|     Name      |   Old   |   New   | Comment                |
|:-------------:|:-------:|:-------:|:-----------------------|
|  timesacledb  | 2.21.1  | 2.22.0  |                        |
|     citus     | 13.1.0  | 13.2.0  |                        |
|  documentdb   | 0.105.0 | 0.106.0 | work with ferretdb 2.5 |
|     ddlx      |  0.29   |  0.30   | + pg18                 |
|    uint128    |  1.0.0  |  1.1.0  | + pg18                 |
|    vchord     |  0.4.3  |  0.5.1  | pgrx 0.16.0            |
|   pg_idkit    |  0.3.0  |  0.3.1  | pgrx 0.15.0            |
|   pg_search   | 0.17.3  | 0.18.0  | pgrx 0.15.0            |
|  pg_parquet   |  0.4.0  |  0.4.3  | pgrx 0.16.0            |
|   wrappers    |  0.5.3  |  0.5.4  | pgrx 0.14.3            |
|  pg_rewrite   |    -    |  2.0.0  | + Debian/Ubuntu        |
|  pg_tracing   |    -    | 0.1.3-2 | + pg 14/18             |
|    pg_curl    |   2.4   |  2.4.5  |                        |
|    pg_ivm     |  1.11   |  1.12   | + pg18                 |
|  pg_rewrite   |    -    |  2.0.0  | new extension          |
|  pg_tracing   |    -    |  1.3.0  | + pg14 / pg18          |
|   pgactive    |  2.1.5  |  2.1.6  | + pg18                 |
|  pgsentinel   |   1.1   |   1.2   | 1.2                    |
|    pg_tle     | 1.5.1-1 | 1.5.1-2 | + pg18                 |
|   redis_fdw   |         |         | + pg18                 |
|     emaj      |   4.6   |   4.7   |                        |
| table_version | 1.11.0  | 1.11.1  |                        |

## 2025-07-24

|         Name          |    Old     |         New         | Comment                       |
|:---------------------:|:----------:|:-------------------:|:------------------------------|
|       OrioleDB        | beta11 1.4 |     beta12 1.5      | work with oriolepg 17.11      |
|       OriolePG        |    17.9    |        17.11        | work with orioledb 1.5 beta12 |
|      documentdb       |  0.104.0   |       0.105.0       | work with ferretdb 2.4        |
|      timescaledb      |   2.20.0   |       2.21.1        |                               |
|       supautils       |   2.9.2    |       2.10.0        | `.so` location change         |
|         plv8          |   3.2.3    |        3.2.4        |                               |
| postgresql_anonymizer |   3.1.1    | 2.3.0 (pgrx 0.14.3) |                               |
|       wrappers        |   0.5.0    | 0.5.3 (pgrx 0.14.3) | pgrx version change           |
|     pgvectorscale     |   0.7.1    | 0.8.0 (pgrx 0.12.9) |                               |
|       pg_search       |   0.15.8   |  0.17.0 (download)  | fix el icu dependency issue   |
|      pg_profile       |   4.8.0    |       4.10.0        |                               |

## 2025-07-04

|      Name       |  Old   |    New     | Comment            |
|:---------------:|:------:|:----------:|:-------------------|
|    orioledb     |        | 1.4 beta11 | rebuild            |
|  pgvectorscale  | 0.7.1  |   0.7.1    | rebuild to fix bug |
| pg_stat_monitor | 2.1.1  |   2.2.0    |                    |
|  pgsql-tweaks   | 0.11.1 |   0.11.3   |                    |
|     pg_tle      | 1.5.0  |   1.5.1    |                    |
|     pg_curl     |  2.4   |   2.4.5    |                    |

## 2025-06-24

|    Name     |   Old   |                                  New                                   | Comment           |
|:-----------:|:-------:|:----------------------------------------------------------------------:|:------------------|
|    citus    | 13.0.3  |                                 13.1.0                                 |                   |
| timescaledb | 2.20.0  |                                 2.21.0                                 |                   |
|   vchord    |  0.3.0  | [0.4.3](https://github.com/tensorchord/VectorChord/releases/tag/0.4.3) |                   |
|  pgactive   |    -    |                                 2.1.5                                  | require pgfeutils |
| documentdb  | 0.103.0 |                                0.104.0                                 | add arm support   |

## 2025-05-26

|      Name       |   Old    |                           New                           | Comment |
|:---------------:|:--------:|:-------------------------------------------------------:|:--------|
|      pgdd       |  0.5.0   |     [0.6.0](https://github.com/rustprooflabs/pgdd)      |         |
|     convert     |    -     |    [0.0.4](https://github.com/rustprooflabs/convert)    |         |
|    pg_idkit     |  0.2.0   |     [0.3.0](https://github.com/VADOSWARE/pg_idkit)      |         |
| pg_tokenizer.rs |    -     | [0.1.0](https://github.com/tensorchord/pg_tokenizer.rs) |         |
|    pg_render    |    -     |      [0.1.2](https://github.com/mkaski/pg_render)       |         |
|    pgx_ulid     |    -     |     [0.2.0](https://github.com/pksunkara/pgx_ulid)      |         |
|     pg_ivm      |  1.10.0  |       [1.11.0](https://github.com/sraoss/pg_ivm)        |         |
|    orioledb     | 1.4.0b10 |    [1.4.0b11](https://github.com/orioledb/orioledb)     |         |

## 2025-05-22

|     Name     | Old |                                          New                                           | Comment |
|:------------:|:---:|:--------------------------------------------------------------------------------------:|:-------:|
|  openhalodb  |  -  |                       [14.10](https://github.com/pgsty/openHalo)                       |         |
|     spat     |  -  |                   [0.1.0a4](https://github.com/Florents-Tselai/spat)                   |         |
|  pgsentinel  |  -  |         [1.1.0](https://github.com/pgsentinel/pgsentinel/releases/tag/v1.1.0)          |         |
| timescaledb  |  -  |         [2.20.0](https://github.com/timescale/timescaledb/releases/tag/2.20.0)         |         |
|  sqlite_fdw  |  -  |          [2.5.0](https://github.com/pgspider/sqlite_fdw/releases/tag/v2.5.0)           |         |
|  documentdb  |  -  | [0.103.0](https://github.com/FerretDB/documentdb/releases/tag/v0.103.0-ferretdb-2.2.0) |         |
|     tzf      |  -  |           [0.2.2](https://github.com/ringsaturn/pg-tzf/releases/tag/v0.2.2)            |         |
| pg_vectorize |  -  |        [0.22.2](https://github.com/ChuckHend/pg_vectorize/releases/tag/v0.22.2)        |         |
|   wrappers   |  -  |           [0.5.0](https://github.com/supabase/wrappers/releases/tag/v0.5.0)            |         |

## 2025-05-07

|        Name         | Old |                                               New                                                | Comment |
|:-------------------:|:---:|:------------------------------------------------------------------------------------------------:|:-------:|
|      omnigres       |  -  | [20250507](https://github.com/omnigres/omnigres/commit/413feff21f9f7310023d8cfd92b83f2a251b1aa4) |         |
|        citus        |  -  |                [12.0.3](https://github.com/citusdata/citus/releases/tag/v13.0.3)                 |         |
|     timescaledb     |  -  |              [2.19.3](https://github.com/timescale/timescaledb/releases/tag/2.19.3)              |         |
|      supautils      |  -  |                [2.9.1](https://github.com/supabase/supautils/releases/tag/v2.9.1)                |         |
|      pg_envvar      |  -  |                 [1.0.1](https://github.com/theory/pg-envvar/releases/tag/v1.0.1)                 |         |
|    pgcollection     |  -  |                 [1.0.0](https://github.com/aws/pgcollection/releases/tag/v1.0.0)                 |         |
|    aggs_for_vecs    |  -  |              [1.4.0](https://github.com/pjungwir/aggs_for_vecs/releases/tag/1.4.0)               |         |
|     pg_tracing      |  -  |                [0.1.3](https://github.com/DataDog/pg_tracing/releases/tag/v0.1.3)                |         |
|        pgmq         |  -  |                    [1.5.1](https://github.com/pgmq/pgmq/releases/tag/v1.5.1)                     |         |
|       tzf-pg        |  -  |                [0.2.0](https://github.com/ringsaturn/tzf-pg/releases/tag/v0.2.0)                 |         |
|      pg_search      |  -  |              [0.15.18](https://github.com/paradedb/paradedb/releases/tag/v0.15.18)               |         |
|        anon         |  -  |   [2.1.1](https://gitlab.com/dalibo/postgresql_anonymizer/-/tree/latest/debian?ref_type=heads)   |         |
|     pg_parquet      |  -  |              [0.4.0](https://github.com/CrunchyData/pg_parquet/releases/tag/v0.3.2)              |         |
|     pg_cardano      |  -  |                 [1.0.5](https://github.com/Fell-x27/pg_cardano/commits/master/)                  |         |
|    pglite_fusion    |  -  |              [0.0.5](https://github.com/frectonz/pglite-fusion/releases/tag/v0.0.5)              |         |
|     vchord_bm25     |  -  |           [0.2.1](https://github.com/tensorchord/VectorChord-bm25/releases/tag/0.2.1)            |         |
|       vchord        |  -  |              [0.3.0](https://github.com/tensorchord/VectorChord/releases/tag/0.3.0)              |         |
| timescaledb-toolkit |  -  |          [1.21.0](https://github.com/timescale/timescaledb-toolkit/releases/tag/1.21.0)          |         |
|    pgvectorscale    |  -  |              [0.7.1](https://github.com/timescale/pgvectorscale/releases/tag/0.7.1)              |         |
|   pg_session_jwt    |  -  |           [0.3.1](https://github.com/neondatabase/pg_session_jwt/releases/tag/v0.3.1)            |         |

## 2025-03-20

|       Name        |  Old  |  New   | Comment |
|:-----------------:|:-----:|:------:|:-------:|
|    timescaledb    |   -   | 2.19.0 |         |
|       citus       |   -   | 13.0.2 |         |
|    documentdb     |   -   | 1.102  |         |
|   pg_analytics    |   -   | 0.3.7  |         |
|     pg_search     |   -   | 0.15.8 |         |
|      pg_ivm       |   -   |  1.10  |         |
|       emaj        |   -   | 4.6.0  |         |
|   pgsql_tweaks    |   -   | 0.11.0 |         |
|   pgvectorscale   |   -   | 0.6.0  |         |
|  pg_session_jwt   |   -   | 0.2.0  |         |
|     wrappers      |   -   | 0.4.5  |         |
|    pg_parquet     |   -   | 0.3.1  |         |
|      vchord       |   -   | 0.2.2  |         |
|      pg_tle       | 1.2.0 | 1.5.0  |         |
|     supautils     | 2.5.0 | 2.6.0  |         |
|     sslutils      |  1.3  |  1.4   |         |
|    pg_profile     |  4.7  |  4.8   |         |
|   pg_jsonschema   | 0.3.2 | 0.3.3  |         |
|  pg_incremental   | 1.1.1 | 1.2.0  |         |
| ddl_historization |  0.7  | 0.0.7  |         |
|     pg_sqlog      | 3.1.7 |  1.6   |         |
|     pg_random     |   -   |   -    |         |
|  pg_stat_monitor  | 2.1.0 | 2.1.1  |         |
|    pg_profile     |  4.7  |  4.8   |         |

## 2024-10-16

|      Name       | Old |     New      | Comment |
|:---------------:|:---:|:------------:|:-------:|
|     pg_ivm      |  -  |     1.9      |         |
|  pg_timeseries  |  -  |    0.1.6     |         |
|      pgmq       |  -  |    1.4.4     |         |
|   pg_protobuf   |  -  |    16 17     |         |
|    pg_uuidv7    |  -  |     1.6      |         |
|   pg_readonly   |  -  |    latest    |         |
|      pgddl      |  -  |     0.28     |         |
|  pg_safeupdate  |  -  |    latest    |         |
| pg_stat_monitor |  -  |     2.1      |         |
|   pg_profile    |  -  |     4.7      |         |
|  system_stats   |  -  |     3.2      |         |
|   pg_auth_mon   |  -  |     3.0      |         |
|   login_hook    |  -  |     1.6      |         |
|    logerrors    |  -  |    2.1.3     |         |
|   pg-orphaned   |  -  |    latest    |         |
|    pgnodemx     |  -  |     1.7      |         |
|    sslutils     |  -  | 1.4 (+16,17) |         |
