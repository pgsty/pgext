---
title: RPM 变更日志
description: RPM 扩展包变更日志
weight: 200
---

参阅 [PGSQL 仓库](/zh/repo/pgsql)，了解如何使用 PGSQL YUM/DNF 仓库。

## 2025-11-15

| 名称                                                            |   旧版本   |   新版本   | 备注      |
|:--------------------------------------------------------------|:-------:|:-------:|:--------|
| [pg_textsearch](https://github.com/timescale/pg_textsearch)   |    -    |  0.1.0  | 新扩展     |
| [pg_clickhouse](https://github.com/clickhouse/pg_clickhouse/) |    -    |  0.1.0  | 新扩展     |
| [pg_ai_query](https://github.com/benodiwal/pg_ai_query)       |    -    |  0.1.1  | 新扩展     |
| timescaledb                                                   | 2.23.1  | 2.24.0  |         |
| pg_search                                                     | 0.20.0  | 0.20.4  |         |
| pg_duckdb                                                     | 1.1.0-1 | 1.1.0-2 | 官方版本    |
| pg_biscuit                                                    |   1.0   |  2.0.1  | 新仓库     |
| pg_convert                                                    |  0.0.4  |  0.0.5  | 移除 pg13 |
| pgdd                                                          |  0.6.0  |  0.6.1  | 移除 pg13 |
| pglinter                                                      |  1.0.0  |  1.0.1  |         |
| pg_session_jwt                                                |  0.3.3  |  0.4.0  |         |
| pg_anon                                                       |  2.4.1  |  2.5.1  |         |
| pg_enigma                                                     |  0.4.0  |  0.5.0  |         |
| wrappers                                                      |  0.5.6  |  0.5.7  |         |
| pg_vectorize                                                  | 0.25.0  | 0.26.0  |         |

## 2025-12-04

| 名称      | 旧版本 | 新版本 | 备注       |
|:--------|:---:|:---:|:---------|
| synchdb |  -  | 1.3 | EL9 only |


## 2025-11-20

| 名称                     |  旧版本   |  新版本   | 备注                         |
|:-----------------------|:------:|:------:|:---------------------------|
| vchord                 | 0.5.3  | 1.0.0  |                            |
| pg_later               | 0.3.1  | 0.4.0  |                            |
| pgvectorscale          | 0.8.0  | 0.9.0  | -pg13, +pg18               |
| pglite_fusion          | 0.0.5  | 0.0.6  |                            |
| pgx_ulid               | 0.2.1  | 0.2.2  |                            |
| pg_search              | 0.19.5 | 0.19.7 | resume PIGSTY building     |
| citus                  | 13.2.0 | 13.2.0 | official tag               |
| timescaledb            | 2.23.0 | 2.23.1 |                            |
| pg_profile             |  4.10  |  4.11  |                            |
| pglinter               |        | 1.0.0  | new                        |
| pg_typeid              |        | 0.3.0  | head with pg18 support     |
| pg_enigma              |        | 0.4.0  | vonng patched pgrx version |
| pg_retry               |        | 1.0.0  | new, pg17-18               |
| pg_biscuit             |        |  1.0   | new, pg16-18               |
| pg_weighted_statistics |        | 1.0.0  | new, pg13-18               |
| pg_stat_monitor        | 2.2.0  | 2.3.1  | fix PGDG pg18 missing      |
| pgmq                   | 1.7.0  | 1.8.0  |                            |
| documentdb             | 0.106  | 0.107  | ferretdb fork              |
| PolarDB                |        | 15.15  | 15.15.5.0-38948055         |

## 2025-11-10

为几乎所有扩展添加 PostgreSQL 18 支持

|          名称           |    旧版本     |    新版本     |  备注   |
|:---------------------:|:----------:|:----------:|:-----:|
|       omni_csv        |     -      |   0.1.1    |  new  |
|     omni_datasets     |     -      |   0.1.0    |  new  |
|      omni_shmem       |     -      |   0.1.0    |  new  |
|        pg_csv         |     -      |   1.0.1    |  new  |
|    pg_dbms_errlog     |     -      |    2.2     |  new  |
|       pg_rrule        |     -      |   0.2.0    |  new  |
|        plxslt         |     -      | 0.20140221 |  new  |
|         anon          |   2.3.0    |   2.4.1    | +pg18 |
|      collection       |   1.0.0    |   1.1.0    | +pg18 |
|       credcheck       |    3.0     |    4.2     | +pg18 |
|         emaj          |   4.7.0    |   4.7.1    | +pg18 |
|      explain_ui       |   0.0.1    |   0.0.2    | +pg18 |
|     firebird_fdw      |   1.4.0    |   1.4.1    | +pg18 |
|       logerrors       |   2.1.3    |   2.1.5    | +pg18 |
|       multicorn       |    3.0     |    3.2     | +pg18 |
|         omni          |   0.2.9    |   0.2.14   | +pg18 |
|      omni_email       |     0      |   0.1.0    | +pg18 |
|      omni_httpc       |   0.1.5    |   0.1.10   | +pg18 |
|      omni_httpd       |   0.4.6    |   0.4.11   | +pg18 |
|        omni_id        |   0.4.2    |   0.4.3    | +pg18 |
|       omni_kube       |   0.1.1    |   0.4.2    | +pg18 |
|      omni_ledger      |   0.1.2    |   0.1.3    | +pg18 |
|       omni_sql        |   0.5.1    |   0.5.3    | +pg18 |
|      omni_sqlite      |   0.1.2    |   0.2.2    | +pg18 |
|      omni_types       |   0.3.4    |   0.3.6    | +pg18 |
|       omni_vfs        |   0.2.1    |   0.2.2    | +pg18 |
|      omni_worker      |   0.1.0    |   0.2.1    | +pg18 |
|        periods        |   1.2.2    |   1.2.3    | +pg18 |
|     pg_bestmatch      |   0.0.1    |   0.0.2    | +pg18 |
|      pg_cardano       |   1.0.5    |   1.1.1    | +pg18 |
|     pg_checksums      |    1.1     |    1.3     | +pg18 |
|       pg_duckdb       |   0.3.1    |   1.1.0    | +pg18 |
|   pg_failover_slots   |   1.1.0    |   1.2.0    | +pg18 |
|      pg_graphql       |   1.5.11   |   1.5.12   | +pg18 |
|       pg_idkit        |   0.3.1    |   0.4.0    | +pg18 |
|       pg_later        |   0.3.0    |   0.3.1    | +pg18 |
|      pg_mooncake      |   0.1.2    |   0.2.0    | +pg18 |
|        pg_net         |   0.9.2    |   0.20.0   | +pg18 |
|      pg_parquet       |   0.4.3    |   0.5.1    | +pg18 |
|       pg_render       |   0.1.2    |   0.1.3    | +pg18 |
|    pg_session_jwt     |   0.3.1    |   0.3.3    | +pg18 |
|    pg_smtp_client     |   0.2.0    |   0.2.1    | +pg18 |
|       pg_sphere       |   1.5.1    |   1.5.2    | +pg18 |
| pg_statement_rollback |    1.4     |    1.5     | +pg18 |
|    pg_store_plans     |    1.8     |    1.9     | +pg18 |
|        pg_tle         |   1.5.1    |   1.5.2    | +pg18 |
|     pg_tokenizer      |   0.1.0    |   0.1.1    | +pg18 |
|       pg_uuidv7       |   1.6.0    |   1.7.0    | +pg18 |
|       pgactive        |   2.1.6    |   2.1.7    | +pg18 |
|       pglogical       |   2.4.5    |   2.4.6    | +pg18 |
|   pglogical_origin    |   2.4.5    |   2.4.6    | +pg18 |
|         pgmq          |   1.5.1    |   1.7.0    | +pg18 |
|      pgsmcrypto       |   0.1.0    |   0.1.1    | +pg18 |
|       pgx_ulid        |   0.2.0    |   0.2.1    | +pg18 |
|       pldbgapi        |    1.8     |    1.9     | +pg18 |
|        pljava         |   1.6.8    |   1.6.10   | +pg18 |
|        plprql         |   1.0.0    |   18.0.0   | +pg18 |
|     roaringbitmap     |   0.5.4    |   0.5.5    | +pg18 |
|        semver         |   0.32.1   |   0.40.0   | +pg18 |
|       supautils       |   2.10.0   |   3.0.2    | +pg18 |
|        tds_fdw        |   2.0.4    |   2.0.5    | +pg18 |
|      timescaledb      |   2.22.0   |   2.23.0   | +pg18 |
|  timescaledb_toolkit  |   1.21.0   |   1.22.0   | +pg18 |
|      timeseries       |   0.1.6    |   0.1.7    | +pg18 |
|          tzf          |   0.2.2    |   0.2.3    | +pg18 |
|        vchord         |   0.5.1    |   0.5.3    | +pg18 |
|      vchord_bm25      |   0.2.1    |   0.2.2    | +pg18 |
|       vectorize       |   0.22.2   |   0.25.0   | +pg18 |
|       wrappers        |   0.5.4    |   0.5.6    | +pg18 |
|         gzip          |   1.0.1    |   1.0.0    | +pg18 |
|        hypopg         |   1.4.1    |   1.4.2    | +pg18 |
|      mobilitydb       |   1.2.0    |   1.3.0    | +pg18 |
|       mongo_fdw       |   5.5.1    |   5.5.3    | +pg18 |
|        orafce         |   4.14.4   |   4.14.6   | +pg18 |
|     pg_hint_plan      |   1.7.1    |   1.8.0    | +pg18 |
|        pg_ivm         |    1.11    |    1.13    | +pg18 |
|      pg_partman       |   5.2.4    |   5.3.1    | +pg18 |
|       pg_search       |   0.18.1   |   0.19.2   | +pg18 |
|     pg_show_plans     |   2.1.6    |   2.1.7    | +pg18 |
|        pgpcre         |     1      | 0.20190509 | +pg18 |
|       pgroonga        |   4.0.0    |   4.0.4    | +pg18 |
|   pgroonga_database   |   4.0.0    |   4.0.4    | +pg18 |
|     plpgsql_check     |   2.8.2    |   2.8.3    | +pg18 |
|         uint          | 1.20231206 | 1.20250815 | +pg18 |
|        uint128        |   1.1.0    |   1.1.1    | +pg18 |
|        omni_*         |  20250525  |  20251108  | +pg18 |
|          acl          |   1.0.4    |            | +pg18 |
|    aggs_for_arrays    |   1.3.3    |            | +pg18 |
|     aggs_for_vecs     |   1.4.0    |            | +pg18 |
|       arraymath       |    1.1     |            | +pg18 |
|        asn1oid        |    1.6     |            | +pg18 |
|        aws_s3         |   0.0.1    |            | +pg18 |
|        base36         |   1.0.0    |            | +pg18 |
|        base62         |   0.0.1    |            | +pg18 |
|         bzip          |   1.0.0    |            | +pg18 |
|        chkpass        |    1.0     |            | +pg18 |
|        convert        |   0.0.4    |            | +pg18 |
|    count_distinct     |   3.0.2    |            | +pg18 |
|        country        |   0.0.3    |            | +pg18 |
|       cryptint        |   1.0.0    |            | +pg18 |
|       currency        |   0.0.3    |            | +pg18 |
|  data_historization   |   1.1.0    |            | +pg18 |
|      db_migrator      |   1.0.0    |            | +pg18 |
|         dbt2          |   0.61.7   |            | +pg18 |
|   ddl_historization   |   0.0.7    |            | +pg18 |
|       ddsketch        |   1.0.1    |            | +pg18 |
|      decoder_raw      |    1.0     |            | +pg18 |
|      decoderbufs      |   3.2.0    |            | +pg18 |
|       emailaddr       |     0      |            | +pg18 |
|        envvar         |   1.0.1    |            | +pg18 |
|         faker         |   0.5.3    |            | +pg18 |
|       financial       |   1.0.1    |            | +pg18 |
|          fio          |    1.0     |            | +pg18 |
|    first_last_agg     |   0.1.4    |            | +pg18 |
|       floatfile       |   1.3.1    |            | +pg18 |
|       floatvec        |   1.1.1    |            | +pg18 |
|         geoip         |   0.3.0    |            | +pg18 |
|        hashlib        |    1.1     |            | +pg18 |
|       hashtypes       |   0.1.5    |            | +pg18 |
|          hll          |    2.18    |            | +pg18 |
|      hunspell_*       |    1.0     |            | +pg18 |
|        imgsmlr        |    1.0     |            | +pg18 |
|     index_advisor     |   0.2.0    |            | +pg18 |
|       kafka_fdw       |   0.0.3    |            | +pg18 |
|      login_hook       |    1.7     |            | +pg18 |
|      oracle_fdw       |   2.8.0    |            | +pg18 |
|      pg_auth_mon      |    3.0     |            | +pg18 |
|     pg_background     |    1.3     |            | +pg18 |
|        pg_bigm        |    1.2     |            | +pg18 |
|        pg_cron        |   1.6.7    |            | +pg18 |
|      pg_profile       |    4.10    |            | +pg18 |
|    pg_stat_kcache     |   2.3.0    |            | +pg18 |
|         pgdd          |   0.6.0    |            | +pg18 |
|         pgjwt         |   0.2.0    |            | +pg18 |
|       pgnodemx        |    1.7     |            | +pg18 |
|       pgsodium        |   3.1.9    |            | +pg18 |
|         pgtap         |   1.3.3    |            | +pg18 |
|      plprofiler       |   4.2.5    |            | +pg18 |
|        plproxy        |   2.11.0   |            | +pg18 |
|          plr          |   8.4.8    |            | +pg18 |
|         plv8          |   3.2.4    |            | +pg18 |
|      pointcloud       |   1.2.5    |            | +pg18 |
|         powa          |   5.0.1    |            | +pg18 |
|        prefix         |   1.2.10   |            | +pg18 |
|          q3c          |   2.0.1    |            | +pg18 |
|       redis_fdw       |    1.0     |            | +pg18 |
|   session_variable    |    3.4     |            | +pg18 |
|       set_user        |   4.1.0    |            | +pg18 |
|     system_stats      |    3.2     |            | +pg18 |
|    temporal_tables    |   1.2.2    |            | +pg18 |
|         topn          |   2.7.0    |            | +pg18 |
|         unit          |    7.10    |            | +pg18 |
|       zhparser        |    2.3     |            | +pg18 |
|         zstd          |   1.1.2    |            | +pg18 |




## 2025-09-04

|     名称      |   旧版本   |   新版本   |           备注           |
|:-----------:|:-------:|:-------:|:----------------------:|
| timesacledb | 2.21.1  | 2.22.0  |                        |
|    citus    | 13.1.0  | 13.2.0  |                        |
| documentdb  | 0.105.0 | 0.106.0 | work with ferretdb 2.5 |
|    ddlx     |  0.29   |  0.30   |         + pg18         |
|   icu_ext   |  1.9.0  | 1.10.0  |         + pg18         |
|   asn1oid   |   1.5   |   1.6   |         + pg18         |
|   uint128   |  1.0.0  |  1.1.0  |         + pg18         |
|  toastinfo  |   1.5   |   1.6   |         + pg18         |
|   vchord    |  0.4.3  |  0.5.1  |      pgrx 0.16.0       |
|  pg_idkit   |  0.3.0  |  0.3.1  |      pgrx 0.15.0       |
|  pg_search  | 0.17.3  | 0.18.0  |      pgrx 0.15.0       |
| pg_parquet  |  0.4.0  |  0.4.3  |      pgrx 0.16.0       |
|  wrappers   |  0.5.3  |  0.5.4  |      pgrx 0.14.3       |
| pg_rewrite  |    -    |  2.0.0  | + Debian/Ubuntu (PGDG) |
| pg_tracing  |    -    | 0.1.3-2 |       + pg 14/18       |
|   pg_curl   |   2.4   |  2.4.5  |   new version epoch    |
| pg_rewrite  |    -    |  2.0.0  |    Import from PGDG    |
| pg_tracing  |    -    |  1.3.0  |     + pg14 / pg18      |
|  pgactive   |  2.1.5  |  2.1.6  |         + pg18         |
|  sentinel   |   1.1   |   1.2   |          1.2           |
|   pg_tle    | 1.5.1-1 | 1.5.1-2 |         + pg18         |
|  redis_fdw  |         |         |         + pg18         |
| pgextwlist  |  1.17   |  1.19   |         + pg18         |
|  wal2json   |         |   1.6   |         + pg18         |
|  pgvector   |         |  0.8.1  |         + pg18         |


## 2025-07-24

|          名称           |    旧版本     |         新版本         |              备注               |
|:---------------------:|:----------:|:-------------------:|:-----------------------------:|
|       OrioleDB        | beta11 1.4 |     beta12 1.5      |   work with oriolepg 17.11    |
|       OriolePG        |    17.9    |        17.11        | work with orioledb 1.5 beta12 |
|      documentdb       |  0.104.0   |       0.105.0       |    work with ferretdb 2.4     |
|      timescaledb      |   2.20.0   |       2.21.1        |                               |
|       supautils       |   2.9.2    |       2.10.0        |     `.so` location change     |
|         plv8          |   3.2.3    |        3.2.4        |                               |
| postgresql_anonymizer |   3.1.1    | 2.3.0 (pgrx 0.14.3) |                               |
|       wrappers        |   0.5.0    | 0.5.3 (pgrx 0.14.3) |      pgrx version change      |
|     pgvectorscale     |   0.7.1    | 0.8.0 (pgrx 0.12.9) |                               |
|       pg_search       |   0.15.8   |  0.17.0 (download)  |  fix el icu dependency issue  |



## 2025-06-24

|     名称      |   旧版本   |                                  新版本                                   |        备注         |
|:-----------:|:-------:|:----------------------------------------------------------------------:|:-----------------:|
|    citus    | 13.0.3  |                                 13.1.0                                 |                   |
| timescaledb | 2.20.0  |                                 2.21.0                                 |                   |
|   vchord    |  0.3.0  | [0.4.3](https://github.com/tensorchord/VectorChord/releases/tag/0.4.3) |                   |
|  pgactive   |    -    |                                 2.1.5                                  | require pgfeutils |
| documentdb  | 0.103.0 |                                0.104.0                                 |  add arm support  |


## 2025-05-26

|       名称        |   旧版本    |                           新版本                           | 备注 |
|:---------------:|:--------:|:-------------------------------------------------------:|:--:|
|      pgdd       |  0.5.0   |     [0.6.0](https://github.com/rustprooflabs/pgdd)      |    |
|     convert     |    -     |    [0.0.4](https://github.com/rustprooflabs/convert)    |    |
|    pg_idkit     |  0.2.0   |     [0.3.0](https://github.com/VADOSWARE/pg_idkit)      |    |
| pg_tokenizer.rs |    -     | [0.1.0](https://github.com/tensorchord/pg_tokenizer.rs) |    |
|    pg_render    |    -     |      [0.1.2](https://github.com/mkaski/pg_render)       |    |
|    pgx_ulid     |    -     |     [0.2.0](https://github.com/pksunkara/pgx_ulid)      |    |
|    orioledb     | 1.4.0b10 |    [1.4.0b11](https://github.com/orioledb/orioledb)     |    |


## 2025-05-22

|      名称      | 旧版本 |                                          新版本                                           | 备注 |
|:------------:|:---:|:--------------------------------------------------------------------------------------:|:--:|
|  openhalodb  |  -  |                       [14.10](https://github.com/pgsty/openHalo)                       |    |
|     spat     |  -  |                   [0.1.0a4](https://github.com/Florents-Tselai/spat)                   |    |
|  pgsentinel  |  -  |         [1.1.0](https://github.com/pgsentinel/pgsentinel/releases/tag/v1.1.0)          |    |
| timescaledb  |  -  |         [2.20.0](https://github.com/timescale/timescaledb/releases/tag/2.20.0)         |    |
|  sqlite_fdw  |  -  |          [2.5.0](https://github.com/pgspider/sqlite_fdw/releases/tag/v2.5.0)           |    |
|  documentdb  |  -  | [0.103.0](https://github.com/FerretDB/documentdb/releases/tag/v0.103.0-ferretdb-2.2.0) |    |
|     tzf      |  -  |           [0.2.2](https://github.com/ringsaturn/pg-tzf/releases/tag/v0.2.2)            |    |
| pg_vectorize |  -  |        [0.22.2](https://github.com/ChuckHend/pg_vectorize/releases/tag/v0.22.2)        |    |
|   wrappers   |  -  |           [0.5.0](https://github.com/supabase/wrappers/releases/tag/v0.5.0)            |    |


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

|        名称         |  旧版本  |  新版本   | 备注 |
|:-----------------:|:-----:|:------:|:--:|
|    timescaledb    |   -   | 2.19.0 |    |
|       citus       |   -   | 13.0.2 |    |
|    documentdb     |   -   | 1.102  |    |
|   pg_analytics    |   -   | 0.3.7  |    |
|     pg_search     |   -   | 0.15.8 |    |
|       emaj        |   -   | 4.6.0  |    |
|   pgsql_tweaks    |   -   | 0.11.0 |    |
|   pgvectorscale   |   -   | 0.6.0  |    |
|  pg_session_jwt   |   -   | 0.2.0  |    |
|     wrappers      |   -   | 0.4.5  |    |
|    pg_parquet     |   -   | 0.3.1  |    |
|      vchord       |   -   | 0.2.2  |    |
|      pg_tle       | 1.2.0 | 1.5.0  |    |
|     supautils     | 2.5.0 | 2.6.0  |    |
|     sslutils      |  1.3  |  1.4   |    |
|    pg_profile     |  4.7  |  4.8   |    |
|   pg_jsonschema   | 0.3.2 | 0.3.3  |    |
|  pg_incremental   | 1.1.1 | 1.2.0  |    |
| ddl_historization |  0.7  | 0.0.7  |    |
|     pg_sqlog      | 3.1.7 |  1.6   |    |
|     pg_random     |   -   |   -    |    |
|  pg_stat_monitor  | 2.1.0 | 2.1.1  |    |
|    pg_profile     |  4.7  |  4.8   |    |


## 2024-10-16

|       名称        | 旧版本 |     新版本      | 备注 |
|:---------------:|:---:|:------------:|:--:|
|  pg_timeseries  |  -  |    0.1.6     |    |
|      pgmq       |  -  |    1.4.4     |    |
|   pg_protobuf   |  -  |    16 17     |    |
|    pg_uuidv7    |  -  |     1.6      |    |
|   pg_readonly   |  -  |    latest    |    |
|      pgddl      |  -  |     0.28     |    |
|  pg_safeupdate  |  -  |    latest    |    |
| pg_stat_monitor |  -  |     2.1      |    |
|   pg_profile    |  -  |     4.7      |    |
|  system_stats   |  -  |     3.2      |    |
|   pg_auth_mon   |  -  |     3.0      |    |
|   login_hook    |  -  |     1.6      |    |
|    logerrors    |  -  |    2.1.3     |    |
|   pg-orphaned   |  -  |    latest    |    |
|    pgnodemx     |  -  |     1.7      |    |
|    sslutils     |  -  | 1.4 (+16,17) |    |


