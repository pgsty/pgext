---
title: TODO / Roadmap
description: PG Extensions in the radar! 
weight: 700
---



## Version Bump

- pg_search 0.19.1
- pg_uuidv7 1.7.0
- orafce 4.16.1
- plpgsql_check 2.8.3
- pgTAP 1.3.4
- pglinter 0.0.17
- vchord 0.5.3
- pg_partman 5.3.0
- icu_ext 1.10.0
- pgsql_tweaks 1.0.2
- pg_roaringbitmap 0.5.5
- pgmq 1.7.0
- uint128 1.1.1
- pg_task 2.1.27
- emaj 4.7.1
- ddlx 0.30.0
- hdfs_fdw_13 2.3.3
- logerrors 2.1.5
- mysql_fdw 2.9.3
- pg_snakeoil 1.4
- pg_squeeze 1.9.1
- pg_stat_kcache 2.3.1
- pljava 1.6.10
- tds_fdw 2.0.5
- postgres-decoderbufs 3.3.0
- pg_partman 5.3.0
- pgauditlogtofile 1.7.3

## To Be Added

- pgelog 1.0.2
- oai_fdw 1.11.0
- pgcalendar 1.0.1
- weighted_statistics 1.0.0
- rdf_fdw 2.1.0
- pg_ttl_index 1.0.2

## Icebox

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

## Lack of Maintenance

- hydra: https://github.com/hydradatabase/columnar , no longer maintained since 17
- age: https://github.com/apache/age


## Not Planned

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


## Retired

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

## Missing License

- [jsonb_apply](https://github.com/Florents-Tselai/jsonb_apply) 0.1.0


## EL Only

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


## Debian Only

- mobilitydb
- rdkit
- hstore_pllua
- hstore_plluau
- debversion
- pg_rrule


## Resource

- [**PGXN**](https://pgxn.org/): https://pgxn.org/recent/
- [**PGRPMS**](https://git.postgresql.org/gitweb/?p=pgrpms.git;a=summary) : [https://git.postgresql.org/gitweb](https://git.postgresql.org/gitweb/?p=pgrpms.git)
- [**PGDEBS**](https://salsa.debian.org/postgresql): https://salsa.debian.org/postgresql
- [**Gist** of 1000+ PG Extension](https://gist.github.com/joelonsql/e5aa27f8cc9bd22b8999b7de8aee9d47)
- https://www.pgextensions.org/