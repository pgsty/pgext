---
title: DEB Changelog
description: DEB Extension Changelog
weight: 300
---



## 2025-09-06

|     Name      |   Old   |   New   | Comment                 |
|:-------------:|:-------:|:-------:|:------------------------|
|  timesacledb  | 2.21.1  | 2.22.0  |                         |
|     citus     | 13.1.0  | 13.2.0  |                         |
|  documentdb   | 0.105.0 | 0.106.0 | work with ferretdb 2.5  |
|     ddlx      |  0.29   |  0.30   | + pg18                  |
|    uint128    |  1.0.0  |  1.1.0  | + pg18                  |
|    vchord     |  0.4.3  |  0.5.1  | pgrx 0.16.0             |
|   pg_idkit    |  0.3.0  |  0.3.1  | pgrx 0.15.0             |
|   pg_search   | 0.17.3  | 0.18.0  | pgrx 0.15.0             |
|  pg_parquet   |  0.4.0  |  0.4.3  | pgrx 0.16.0             |
|   wrappers    |  0.5.3  |  0.5.4  | pgrx 0.14.3             |
|  pg_rewrite   |    -    |  2.0.0  | + Debian/Ubuntu         |
|  pg_tracing   |    -    | 0.1.3-2 | + pg 14/18              |
|    pg_curl    |   2.4   |  2.4.5  |                         |
|    pg_ivm     |  1.11   |  1.12   | + pg18                  |
|  pg_rewrite   |    -    |  2.0.0  | new extension           |
|  pg_tracing   |    -    |  1.3.0  | + pg14 / pg18           |
|   pgactive    |  2.1.5  |  2.1.6  | + pg18                  |
|  pgsentinel   |   1.1   |   1.2   | 1.2                     |
|    pg_tle     | 1.5.1-1 | 1.5.1-2 | + pg18                  |
|   redis_fdw   |         |         | + pg18                  |
|     emaj      |   4.6   |   4.7   |                         |
| table_version | 1.11.0  | 1.11.1  |                         |


## 2025-07-24

|         Name          |    Old     |         New         | Comment                        |
|:---------------------:|:----------:|:-------------------:|:-------------------------------|
|       OrioleDB        | beta11 1.4 |     beta12 1.5      | work with oriolepg 17.11       |
|       OriolePG        |    17.9    |        17.11        | work with orioledb 1.5 beta12  |
|      documentdb       |  0.104.0   |       0.105.0       | work with ferretdb 2.4         |
|      timescaledb      |   2.20.0   |       2.21.1        |                                |
|       supautils       |   2.9.2    |       2.10.0        | `.so` location change          |
|         plv8          |   3.2.3    |        3.2.4        |                                |
| postgresql_anonymizer |   3.1.1    | 2.3.0 (pgrx 0.14.3) |                                |
|       wrappers        |   0.5.0    | 0.5.3 (pgrx 0.14.3) | pgrx version change            |
|     pgvectorscale     |   0.7.1    | 0.8.0 (pgrx 0.12.9) |                                |
|       pg_search       |   0.15.8   |  0.17.0 (download)  | fix el icu dependency issue    |
|      pg_profile       |   4.8.0    |       4.10.0        |                                |


## 2025-07-04

|      Name       |  Old   |    New     | Comment             |
|:---------------:|:------:|:----------:|:--------------------|
|    orioledb     |        | 1.4 beta11 | rebuild             |
|  pgvectorscale  | 0.7.1  |   0.7.1    | rebuild to fix bug  |
| pg_stat_monitor | 2.1.1  |   2.2.0    |                     |
|  pgsql-tweaks   | 0.11.1 |   0.11.3   |                     |
|     pg_tle      | 1.5.0  |   1.5.1    |                     |
|     pg_curl     |  2.4   |   2.4.5    |                     |


## 2025-06-24

|    Name     |   Old   |                                  New                                   | Comment           |
|:-----------:|:-------:|:----------------------------------------------------------------------:|:------------------|
|    citus    | 13.0.3  |                                 13.1.0                                 |                   |
| timescaledb | 2.20.0  |                                 2.21.0                                 |                   |
|   vchord    |  0.3.0  | [0.4.3](https://github.com/tensorchord/VectorChord/releases/tag/0.4.3) |                   |
|  pgactive   |    -    |                                 2.1.5                                  | require pgfeutils |
| documentdb  | 0.103.0 |                                0.104.0                                 | add arm support   |


## 2025-05-26

|      Name       |   Old    |                           New                           | Comment  |
|:---------------:|:--------:|:-------------------------------------------------------:|:---------|
|      pgdd       |  0.5.0   |     [0.6.0](https://github.com/rustprooflabs/pgdd)      |          |
|     convert     |    -     |    [0.0.4](https://github.com/rustprooflabs/convert)    |          |
|    pg_idkit     |  0.2.0   |     [0.3.0](https://github.com/VADOSWARE/pg_idkit)      |          |
| pg_tokenizer.rs |    -     | [0.1.0](https://github.com/tensorchord/pg_tokenizer.rs) |          |
|    pg_render    |    -     |      [0.1.2](https://github.com/mkaski/pg_render)       |          |
|    pgx_ulid     |    -     |     [0.2.0](https://github.com/pksunkara/pgx_ulid)      |          |
|     pg_ivm      |  1.10.0  |       [1.11.0](https://github.com/sraoss/pg_ivm)        |          |
|    orioledb     | 1.4.0b10 |    [1.4.0b11](https://github.com/orioledb/orioledb)     |          |


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


