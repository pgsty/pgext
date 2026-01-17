---
title: 扩展待办列表
description: 需要更新，已经退役的 PG 扩展列表。 
weight: 700
---


# 待办事项

- citus 14.0 with pg18 support
- pg_ttl_index 0.1.0
- etcd_fdw 0.0.0
- convert 0.0.5 -> 0.1.0
- pg_textsearch 0.1.0 -> 0.4.0
- pgmq 1.8.0 -> 1.8.1
- documentdb_extened_rum
- mobilitydb_datagen


## 候选列表

- pgmq 1.8.1
- [spock](https://github.com/pgEdge/spock) : Logical multi-master PostgreSQL replication
- [pg_lake](https://github.com/Snowflake-Labs/pg_lake/) : Postgres with Iceberg and data lake access
- [pgai](https://github.com/timescale/pgai) : A suite of tools to develop RAG, semantic search, and other AI applications more easily with PostgreSQL
- [river](https://github.com/riverqueue/river)
- omnigres on legacy systems
- plv8 on el10

## 最近更新

- age 1.6.0 with pg18 support (age)
- pgsentinel 1.3.0 -> 1.3.1 rpm
- pg_timeseries 0.1.8 -> 0.2.0
- pg_clickhouse 0.1.0 -> 0.1.2
- pg_bulkload 3.1.22 -> 3.1.23
- pg_biscuit 2.0.2 -> 2.2.2
- documentdb 0.107 -> 0.109 (switch to upstream)
- pg_partman 5.4.0 new
- pljs new 1.0.4
- pg_timeseries 0.1.8
- pg_duckdb 1.1.1
- pg_summarize rebuild for pg18 on deb
- semver 0.41.0 (rpm only)
- pg_search 0.20.5
- supautils rebuild for pg18 on deb
- [pg_textsearch](https://github.com/timescale/pg_textsearch) 0.1.0 new
- [pg_clickhouse](https://github.com/clickhouse/pg_clickhouse/) 0.1.0 new
- [pg_ai_query](https://github.com/benodiwal/pg_ai_query) 0.1.1 new
- timescaledb 2.23.1 -> 2.24.0
- pg_search 0.20.0 -> 0.20.3
- convert 0.0.4 -> 0.0.5
- pglinter 1.0.0 -> 1.0.1
- pgdd 0.6.0 -> 0.6.1
- pg_session_jwt 0.3.3 -> 0.4.0
- pg_anon 2.4.1 -> 2.5.1
- pg_enigma 0.4.0 -> 0.5.0
- wrappers 0.5.6 -> 0.5.7
- pg_vectorize 0.25.0 -> 0.26.0


## 暂时雪藏

- pgelog 1.0.2
- oai_fdw 1.11.0
- rdf_fdw 2.1.0
- pg_ttl_index 1.0.2
- [pgcalendar](https://github.com/h4kbas/pgcalendar) bad makefile
- [dsef](https://github.com/ardentperf/dsef/)
- [pg_mustach](https://github.com/RekGRpth/pg_mustach)
- [is_jsonb_valid](https://github.com/furstenheim/is_jsonb_valid)
- [pg_kafka](https://github.com/xstevens/pg_kafka)
- [pg_jieba](https://github.com/jaiminpan/pg_jieba)
- [OneSparse](https://github.com/OneSparse/OneSparse)
- [PipelineDB](https://github.com/pipelinedb/pipelinedb)
- [SQL Firewall](https://github.com/uptimejp/sql_firewall)
- [zcurve](https://github.com/bmuratshin/zcurve)
- [PG.NET](https://github.com/Brick-Abode/pldotnet/releases)
- [pg_scws](https://github.com/jaiminpan/pg_scws)
- [themsis](https://github.com/cossacklabs/pg_themis)
- [pgspeck](https://github.com/johto/pgspeck)
- [lsm3](https://github.com/postgrespro/lsm3)
- [monq](https://github.com/postgrespro/monq)
- [pg_badplan](https://github.com/trustly/pg_badplan)
- [pg_recall](https://github.com/mreithub/pg_recall)
- [pgfsm](https://github.com/michelp/pgfsm)
- [pg_trgm pro](https://github.com/postgrespro/pg_trgm_pro)
- [weighted_mean](https://github.com/Kozea/weighted_mean)
- [kmeans](https://github.com/umitanuki/kmeans-postgresql)
- [pgjwt_rs](https://github.com/vishvish/pgjwt_rs)
- [plrust](https://github.com/tcdi/plrust)

## 缺少维护

- hydra: https://github.com/hydradatabase/columnar , no longer maintained since 17
- age: https://github.com/apache/age


## 尚未规划

- pg_base62: https://github.com/bkircher/pg_base62/blob/main/Cargo.toml
- pg_conda: PostgreSQL extension that adds types and functions for the conda ecosystem
- pgfdb: https://github.com/fabianlindfors/pgfdb
- postgres-ical: https://github.com/edgarogh/postgres-ical
- pgsloth: https://github.com/jamessewell/pgsloth
- pgfaker: https://github.com/rustprooflabs/pgfaker
- pglance - PostgreSQL Lance Table Extension
- pg_oidc_validator_rust https://github.com/UnAfraid/pg_oidc_validator_rust
- pg_top: not ready due to cmake error
- pg_quack, we already have a pg_lakehouse
- pg_telemetry, we already have better observability
- pgx_ulid, https://github.com/pksunkara/pgx_ulid, already covered by pg_idkit (MIT, but RUST)
- embedding: obsolete
- zson https://github.com/postgrespro/zson MIT C (too old)
- pghydro https://github.com/pghydro/pghydro C GPL-2.0 6.6 (no makefile)
- pg_natural_sort_order https://github.com/Zeleo/pg_natural_sort_order (too old)
- pg_query_state https://github.com/postgrespro/pg_query_state
- pgsampler https://github.com/no0p/pgsampler
- pg_lz4 https://github.com/zilder/pg_lz4
- pg_amqp https://github.com/omniti-labs/pg_amqp
- tinyint https://github.com/umitanuki/tinyint-postgresql
- pg_blkchain https://github.com/blkchain/pg_blkchain
- hashtypes https://github.com/pandrewhk/hashtypes
- foreign_table_exposer https://github.com/komamitsu/foreign_table_exposer
- ldap_fdw https://github.com/guedes/ldap_fdw
- pg_backtrace https://github.com/postgrespro/pg_backtrace (only works on PG12)
- connection_limits https://github.com/tvondra/connection_limits
- fixeddecimal https://github.com/2ndQuadrant/fixeddecimal
- fuzzywuzzy https://github.com/hooopo/pg-fuzzywuzzy
- pg_paxos https://github.com/microsoft/pg_paxos


## 退役淘汰

- parquet_s3_fdw: retired due to too much duckdb better alternatives
- pg_tier: retired due to parquet_s3_fdw deps
- pg_mon: retired due to pg17 in-compatibility
- pg_search: retired due to moving to official release procedure
- pg_bm25: retired due to renaming to pg_search
- pg_analytics: retired due to moving to official release procedure, and once renaming to pg_lakehouse
- pg_lakehouse: retired due to renaming back to pg_analytics
- pg_sparse: retired due to merge into pgvector, and no longer maintained
- mysqlcompat: retire due to conflict func with higher version of PG
- pg_comparator: retired due to removing from PGDG repo
- pg_proctab: retired due to covered by pgnodemx
- pg_statviz: broken deps and replaceable functionality
- [pg_net](https://github.com/supabase/pg_net) : retired due to moving into PGDG repo
- [pg_tle](https://github.com/aws/pg_tle) : retired due to moving into PGDG repo
- [pg_bigm](https://github.com/pgbigm/pg_bigm) : retired due to moving into PGDG repo
- [pgsql-http](https://github.com/pramsey/pgsql-http) : retired due to moving into PGDG repo
- [pgsql-gzip](https://github.com/pramsey/pgsql-gzip) : retired due to moving into PGDG repo
- [pg_dirtyread](https://github.com/df7cb/pg_dirtyread) : retired due to moving into PGDG repo
- [pointcloud](https://github.com/pgpointcloud/pointcloud) : retired due to moving into PGDG repo
- pg_top: retired due to too much trouble
- pg_timeit: retired due incompatible on arm64
- vacuumlo & oid2name: binary command, actually they are not extensions
- pgdd: remove due to not actively maintained and legacy pgrx version

## 缺少许可证

- [jsonb_apply](https://github.com/Florents-Tselai/jsonb_apply) 0.1.0


## EL 独有

- pg_strom
- faker
- dbt2
- pg_top
- multicorn
- odbc_fdw
- jdbc_fdw
- tds_fdw
- db2_fdw
- sqlite_fdw
- pgbouncer_fdw
- mongo_fdw
- hdfs_fdw
- pg_dbms_metadata
- pg_dbms_lock
- pg_dbms_job


## Debian 独有

- mobilitydb
- rdkit
- hstore_pllua
- hstore_plluau
- debversion
- pg_rrule


## 相关资源

- [**PGXN**](https://pgxn.org/): https://pgxn.org/recent/
- [**PGRPMS**](https://git.postgresql.org/gitweb/?p=pgrpms.git;a=summary) : [https://git.postgresql.org/gitweb](https://git.postgresql.org/gitweb/?p=pgrpms.git)
- [**PGDEBS**](https://salsa.debian.org/postgresql): https://salsa.debian.org/postgresql
- [**Gist** of 1000+ PG Extension](https://gist.github.com/joelonsql/e5aa27f8cc9bd22b8999b7de8aee9d47)
- https://www.pgextensions.org/