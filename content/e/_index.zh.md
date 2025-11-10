---
title: "扩展列表"
breadcrumbs: false
excludeSearch: true
comments: false
weight: 900
---

在 PIGSTY 与 PGDG 仓库中，总共有 431 个可用的 PostgreSQL 扩展，与 361 类软件包：

## 扩展包

共有 361 个可用的 PostgreSQL 扩展软件包：

| 包 | 版本 | 仓库 | 分类 | RPM | DEB |
|:---|:-----|:-----|:-----|:-----|:-----|
| {{< ext "timescaledb" >}} | `2.23.0` | {{< badge content="Link" link="https://github.com/timescale/timescaledb" >}} | {{< category "TIME" >}} | `timescaledb-tsl_$v*` | `postgresql-$v-timescaledb-tsl` |
| {{< ext "timescaledb_toolkit" >}} | `1.22.0` | {{< badge content="Link" link="https://github.com/timescale/timescaledb-toolkit" >}} | {{< category "TIME" >}} | `timescaledb-toolkit_$v` | `postgresql-$v-timescaledb-toolkit` |
| {{< ext "timeseries" "pg_timeseries" >}} | `0.1.7` | {{< badge content="Link" link="https://github.com/ChuckHend/pg_timeseries" >}} | {{< category "TIME" >}} | `pg_timeseries_$v` | `postgresql-$v-pg-timeseries` |
| {{< ext "periods" >}} | `1.2.3` | {{< badge content="Link" link="https://github.com/xocolatl/periods" >}} | {{< category "TIME" >}} | `periods_$v*` | `postgresql-$v-periods` |
| {{< ext "temporal_tables" >}} | `1.2.2` | {{< badge content="Link" link="https://pgxn.org/dist/temporal_tables/" >}} | {{< category "TIME" >}} | `temporal_tables_$v*` | `postgresql-$v-temporal-tables` |
| {{< ext "emaj" >}} | `4.7.1` | {{< badge content="Link" link="https://github.com/dalibo/emaj" >}} | {{< category "TIME" >}} | `e-maj_$v` | `postgresql-$v-emaj` |
| {{< ext "table_version" >}} | `1.11.1` | {{< badge content="Link" link="https://github.com/linz/postgresql-tableversion" >}} | {{< category "TIME" >}} | `table_version_$v` | `postgresql-$v-table-version` |
| {{< ext "pg_cron" >}} | `1.6.7` | {{< badge content="Link" link="https://github.com/citusdata/pg_cron" >}} | {{< category "TIME" >}} | `pg_cron_$v*` | `postgresql-$v-cron` |
| {{< ext "pg_task" >}} | `1.0.0` | {{< badge content="Link" link="https://github.com/RekGRpth/pg_task" >}} | {{< category "TIME" >}} | `pg_task_$v*` | `postgresql-$v-pg-task` |
| {{< ext "pg_later" >}} | `0.3.1` | {{< badge content="Link" link="https://github.com/ChuckHend/pg_later" >}} | {{< category "TIME" >}} | `pg_later_$v` | `postgresql-$v-pg-later` |
| {{< ext "pg_background" >}} | `1.5` | {{< badge content="Link" link="https://github.com/vibhorkum/pg_background" >}} | {{< category "TIME" >}} | `pg_background_$v*` | `postgresql-$v-pg-background` |
| {{< ext "postgis" >}} | `3.6.0` | {{< badge content="Link" link="https://git.osgeo.org/gitea/postgis/postgis" >}} | {{< category "GIS" >}} | `postgis36_$v*` | `postgresql-$v-postgis-3 postgresql-$v-postgis-3-scripts` |
| {{< ext "pgrouting" >}} | `3.8.0` | {{< badge content="Link" link="https://github.com/pgRouting/pgrouting" >}} | {{< category "GIS" >}} | `pgrouting_$v*` | `postgresql-$v-pgrouting postgresql-$v-pgrouting-scripts` |
| {{< ext "pointcloud" >}} | `1.2.5` | {{< badge content="Link" link="https://github.com/pgpointcloud/pointcloud" >}} | {{< category "GIS" >}} | `pointcloud_$v*` | `postgresql-$v-pointcloud` |
| {{< ext "h3" "pg_h3" >}} | `4.2.3` | {{< badge content="Link" link="https://github.com/zachasme/h3-pg" >}} | {{< category "GIS" >}} | `h3-pg_$v*` | `postgresql-$v-h3` |
| {{< ext "q3c" >}} | `2.0.1` | {{< badge content="Link" link="https://github.com/segasai/q3c" >}} | {{< category "GIS" >}} | `q3c_$v*` | `postgresql-$v-q3c` |
| {{< ext "ogr_fdw" >}} | `1.1.7` | {{< badge content="Link" link="https://github.com/pramsey/pgsql-ogr-fdw" >}} | {{< category "GIS" >}} | `ogr_fdw_$v*` | `postgresql-$v-ogr-fdw` |
| {{< ext "geoip" >}} | `0.3.0` | {{< badge content="Link" link="https://github.com/tvondra/geoip" >}} | {{< category "GIS" >}} | `geoip_$v` | `postgresql-$v-geoip` |
| {{< ext "pg_polyline" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/yihong0618/pg_polyline" >}} | {{< category "GIS" >}} | `pg_polyline_$v` | `postgresql-$v-pg-polyline` |
| {{< ext "pg_geohash" >}} | `1.0` | {{< badge content="Link" link="https://github.com/jistok/pg_geohash" >}} | {{< category "GIS" >}} | `pg_geohash_$v*` | `postgresql-$v-pg-geohash` |
| {{< ext "mobilitydb" >}} | `1.3.0` | {{< badge content="Link" link="https://github.com/MobilityDB/MobilityDB" >}} | {{< category "GIS" >}} | - | `postgresql-$v-mobilitydb` |
| {{< ext "tzf" "pg_tzf" >}} | `0.2.3` | {{< badge content="Link" link="https://github.com/ringsaturn/pg-tzf" >}} | {{< category "GIS" >}} | `pg_tzf_$v` | `postgresql-$v-tzf` |
| {{< ext "earthdistance" >}} | `1.2` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/earthdistance.html" >}} | {{< category "GIS" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "vector" "pgvector" >}} | `0.8.1` | {{< badge content="Link" link="https://github.com/pgvector/pgvector" >}} | {{< category "RAG" >}} | `pgvector_$v*` | `postgresql-$v-pgvector` |
| {{< ext "vchord" >}} | `0.5.3` | {{< badge content="Link" link="https://github.com/tensorchord/VectorChord" >}} | {{< category "RAG" >}} | `vchord_$v` | `postgresql-$v-vchord` |
| {{< ext "vectorscale" "pgvectorscale" >}} | `0.8.0` | {{< badge content="Link" link="https://github.com/timescale/pgvectorscale" >}} | {{< category "RAG" >}} | `pgvectorscale_$v` | `postgresql-$v-pgvectorscale` |
| {{< ext "vectorize" "pg_vectorize" >}} | `0.25.0` | {{< badge content="Link" link="https://github.com/ChuckHend/pg_vectorize" >}} | {{< category "RAG" >}} | `pg_vectorize_$v` | `postgresql-$v-pg-vectorize` |
| {{< ext "pg_similarity" >}} | `1.0` | {{< badge content="Link" link="https://github.com/eulerto/pg_similarity" >}} | {{< category "RAG" >}} | `pg_similarity_$v*` | `postgresql-$v-similarity` |
| {{< ext "smlar" >}} | `1.0` | {{< badge content="Link" link="https://github.com/jirutka/smlar" >}} | {{< category "RAG" >}} | `smlar_$v*` | `postgresql-$v-smlar` |
| {{< ext "pg_summarize" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/HexaCluster/pg_summarize" >}} | {{< category "RAG" >}} | `pg_summarize_$v` | `postgresql-$v-pg-summarize` |
| {{< ext "pg_tiktoken" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/kelvich/pg_tiktoken" >}} | {{< category "RAG" >}} | `pg_tiktoken_$v` | `postgresql-$v-pg-tiktoken` |
| {{< ext "pg4ml" >}} | `2.0` | {{< badge content="Link" link="https://gitee.com/guotiecheng/plpgsql_pg4ml" >}} | {{< category "RAG" >}} | `pg4ml_$v` | `postgresql-$v-pg4ml` |
| {{< ext "pgml" >}} | `2.10.0` | {{< badge content="Link" link="https://github.com/postgresml/postgresml" >}} | {{< category "RAG" >}} | `pgml_$v` | `postgresql-$v-pgml` |
| {{< ext "pg_search" >}} | `0.19.3` | {{< badge content="Link" link="https://github.com/paradedb/paradedb/tree/dev/pg_search" >}} | {{< category "FTS" >}} | `pg_search_$v` | `postgresql-$v-pg-search` |
| {{< ext "pgroonga" >}} | `4.0.4` | {{< badge content="Link" link="https://github.com/pgroonga/pgroonga" >}} | {{< category "FTS" >}} | `pgroonga_$v*` | `postgresql-$v-pgroonga` |
| {{< ext "pg_bigm" >}} | `1.2` | {{< badge content="Link" link="https://github.com/pgbigm/pg_bigm" >}} | {{< category "FTS" >}} | `pg_bigm_$v*` | `postgresql-$v-pg-bigm` |
| {{< ext "zhparser" >}} | `2.3` | {{< badge content="Link" link="https://github.com/amutu/zhparser" >}} | {{< category "FTS" >}} | `zhparser_$v*` | `postgresql-$v-zhparser` |
| {{< ext "pg_bestmatch" >}} | `0.0.2` | {{< badge content="Link" link="https://github.com/tensorchord/pg_bestmatch.rs" >}} | {{< category "FTS" >}} | `pg_bestmatch_$v` | `postgresql-$v-pg-bestmatch` |
| {{< ext "vchord_bm25" >}} | `0.2.2` | {{< badge content="Link" link="https://github.com/tensorchord/VectorChord-bm25" >}} | {{< category "FTS" >}} | `vchord_bm25_$v` | `postgresql-$v-vchord-bm25` |
| {{< ext "pg_tokenizer" >}} | `0.1.1` | {{< badge content="Link" link="https://github.com/tensorchord/pg_tokenizer.rs" >}} | {{< category "FTS" >}} | `pg_tokenizer_$v` | `postgresql-$v-pg-tokenizer` |
| {{< ext "hunspell_cs_cz" >}} | `1.0` | {{< badge content="Link" link="https://github.com/postgrespro/hunspell_dicts" >}} | {{< category "FTS" >}} | `hunspell_cs_cz_$v` | `postgresql-$v-hunspell-cs-cz` |
| {{< ext "hunspell_de_de" >}} | `1.0` | {{< badge content="Link" link="https://github.com/postgrespro/hunspell_dicts" >}} | {{< category "FTS" >}} | `hunspell_de_de_$v` | `postgresql-$v-hunspell-de-de` |
| {{< ext "hunspell_en_us" >}} | `1.0` | {{< badge content="Link" link="https://github.com/postgrespro/hunspell_dicts" >}} | {{< category "FTS" >}} | `hunspell_en_us_$v` | `postgresql-$v-hunspell-en-us` |
| {{< ext "hunspell_fr" >}} | `1.0` | {{< badge content="Link" link="https://github.com/postgrespro/hunspell_dicts" >}} | {{< category "FTS" >}} | `hunspell_fr_$v` | `postgresql-$v-hunspell-fr` |
| {{< ext "hunspell_ne_np" >}} | `1.0` | {{< badge content="Link" link="https://github.com/postgrespro/hunspell_dicts" >}} | {{< category "FTS" >}} | `hunspell_ne_np_$v` | `postgresql-$v-hunspell-ne-np` |
| {{< ext "hunspell_nl_nl" >}} | `1.0` | {{< badge content="Link" link="https://github.com/postgrespro/hunspell_dicts" >}} | {{< category "FTS" >}} | `hunspell_nl_nl_$v` | `postgresql-$v-hunspell-nl-nl` |
| {{< ext "hunspell_nn_no" >}} | `1.0` | {{< badge content="Link" link="https://github.com/postgrespro/hunspell_dicts" >}} | {{< category "FTS" >}} | `hunspell_nn_no_$v` | `postgresql-$v-hunspell-nn-no` |
| {{< ext "hunspell_pt_pt" >}} | `1.0` | {{< badge content="Link" link="https://github.com/postgrespro/hunspell_dicts" >}} | {{< category "FTS" >}} | `hunspell_pt_pt_$v` | `postgresql-$v-hunspell-pt-pt` |
| {{< ext "hunspell_ru_ru" >}} | `1.0` | {{< badge content="Link" link="https://github.com/postgrespro/hunspell_dicts" >}} | {{< category "FTS" >}} | `hunspell_ru_ru_$v` | `postgresql-$v-hunspell-ru-ru` |
| {{< ext "hunspell_ru_ru_aot" >}} | `1.0` | {{< badge content="Link" link="https://github.com/postgrespro/hunspell_dicts" >}} | {{< category "FTS" >}} | `hunspell_ru_ru_aot_$v` | `postgresql-$v-hunspell-ru-ru-aot` |
| {{< ext "fuzzystrmatch" >}} | `1.2` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/fuzzystrmatch.html" >}} | {{< category "FTS" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "pg_trgm" >}} | `1.6` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/pgtrgm.html" >}} | {{< category "FTS" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "citus" >}} | `13.2.0` | {{< badge content="Link" link="https://github.com/citusdata/citus" >}} | {{< category "OLAP" >}} | `citus_$v*` | `postgresql-$v-citus` |
| {{< ext "columnar" "hydra" >}} | `1.1.2` | {{< badge content="Link" link="https://github.com/hydradatabase/hydra" >}} | {{< category "OLAP" >}} | `hydra_$v*` | `postgresql-$v-hydra` |
| {{< ext "pg_analytics" >}} | `0.3.7` | {{< badge content="Link" link="https://github.com/paradedb/pg_analytics" >}} | {{< category "OLAP" >}} | `pg_analytics_$v` | `postgresql-$v-pg-analytics` |
| {{< ext "pg_duckdb" >}} | `1.1.0` | {{< badge content="Link" link="https://github.com/duckdb/pg_duckdb" >}} | {{< category "OLAP" >}} | `pg_duckdb_$v*` | `postgresql-$v-pg-duckdb` |
| {{< ext "pg_mooncake" >}} | `0.2.0` | {{< badge content="Link" link="https://github.com/Mooncake-Labs/pg_mooncake" >}} | {{< category "OLAP" >}} | `pg_mooncake_$v*` | `postgresql-$v-pg-mooncake` |
| {{< ext "duckdb_fdw" >}} | `1.1.2` | {{< badge content="Link" link="https://github.com/alitrack/duckdb_fdw" >}} | {{< category "OLAP" >}} | `duckdb_fdw_$v*` | `postgresql-$v-duckdb-fdw` |
| {{< ext "pg_parquet" >}} | `0.5.1` | {{< badge content="Link" link="https://github.com/CrunchyData/pg_parquet/" >}} | {{< category "OLAP" >}} | `pg_parquet_$v` | `postgresql-$v-pg-parquet` |
| {{< ext "pg_fkpart" >}} | `1.7.0` | {{< badge content="Link" link="https://github.com/lemoineat/pg_fkpart" >}} | {{< category "OLAP" >}} | `pg_fkpart_$v` | `postgresql-$v-pg-fkpart` |
| {{< ext "pg_partman" >}} | `5.3.1` | {{< badge content="Link" link="https://github.com/pgpartman/pg_partman" >}} | {{< category "OLAP" >}} | `pg_partman_$v*` | `postgresql-$v-partman` |
| {{< ext "plproxy" >}} | `2.11.0` | {{< badge content="Link" link="https://github.com/plproxy/plproxy" >}} | {{< category "OLAP" >}} | `plproxy_$v*` | `postgresql-$v-plproxy` |
| {{< ext "pg_strom" >}} | `6.0` | {{< badge content="Link" link="https://github.com/heterodb/pg-strom" >}} | {{< category "OLAP" >}} | `pg_strom_$v*` | - |
| {{< ext "tablefunc" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/tablefunc.html" >}} | {{< category "OLAP" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "age" >}} | `1.5.0` | {{< badge content="Link" link="https://github.com/apache/age" >}} | {{< category "FEAT" >}} | `apache-age_$v*` | `postgresql-$v-age` |
| {{< ext "hll" >}} | `2.18` | {{< badge content="Link" link="https://github.com/citusdata/postgresql-hll" >}} | {{< category "FEAT" >}} | `hll_$v*` | `postgresql-$v-hll` |
| {{< ext "rum" >}} | `1.3.14` | {{< badge content="Link" link="https://github.com/postgrespro/rum" >}} | {{< category "FEAT" >}} | `rum_$v` | `postgresql-$v-rum` |
| {{< ext "pg_graphql" >}} | `1.5.12` | {{< badge content="Link" link="https://github.com/supabase/pg_graphql" >}} | {{< category "FEAT" >}} | `pg_graphql_$v` | `postgresql-$v-pg-graphql` |
| {{< ext "pg_jsonschema" >}} | `0.3.3` | {{< badge content="Link" link="https://github.com/supabase/pg_jsonschema" >}} | {{< category "FEAT" >}} | `pg_jsonschema_$v` | `postgresql-$v-pg-jsonschema` |
| {{< ext "jsquery" >}} | `1.2` | {{< badge content="Link" link="https://github.com/postgrespro/jsquery" >}} | {{< category "FEAT" >}} | `jsquery_$v*` | `postgresql-$v-jsquery` |
| {{< ext "pg_hint_plan" >}} | `1.8.0` | {{< badge content="Link" link="https://github.com/ossc-db/pg_hint_plan" >}} | {{< category "FEAT" >}} | `pg_hint_plan_$v*` | `postgresql-$v-pg-hint-plan` |
| {{< ext "hypopg" >}} | `1.4.2` | {{< badge content="Link" link="https://github.com/HypoPG/hypopg" >}} | {{< category "FEAT" >}} | `hypopg_$v*` | `postgresql-$v-hypopg` |
| {{< ext "index_advisor" >}} | `0.2.0` | {{< badge content="Link" link="https://github.com/supabase/index_advisor" >}} | {{< category "FEAT" >}} | `index_advisor_$v` | `postgresql-$v-index-advisor` |
| {{< ext "plan_filter" "pg_plan_filter" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/pgexperts/pg_plan_filter" >}} | {{< category "FEAT" >}} | `pg_plan_filter_$v*` | `postgresql-$v-pg-plan-filter` |
| {{< ext "imgsmlr" >}} | `1.0` | {{< badge content="Link" link="https://github.com/postgrespro/imgsmlr" >}} | {{< category "FEAT" >}} | `imgsmlr_$v*` | `postgresql-$v-imgsmlr` |
| {{< ext "pg_ivm" >}} | `1.13` | {{< badge content="Link" link="https://github.com/sraoss/pg_ivm" >}} | {{< category "FEAT" >}} | `pg_ivm_$v*` | `postgresql-$v-pg-ivm` |
| {{< ext "pg_incremental" >}} | `1.2.0` | {{< badge content="Link" link="https://github.com/CrunchyData/pg_incremental" >}} | {{< category "FEAT" >}} | `pg_incremental_$v*` | `postgresql-$v-pg-incremental` |
| {{< ext "pgmq" >}} | `1.7.0` | {{< badge content="Link" link="https://github.com/pgmq/pgmq" >}} | {{< category "FEAT" >}} | `pgmq_$v` | `postgresql-$v-pgmq` |
| {{< ext "pgq" >}} | `3.5.1` | {{< badge content="Link" link="https://github.com/pgq/pgq" >}} | {{< category "FEAT" >}} | `pgq_$v*` | `postgresql-$v-pgq3` |
| {{< ext "orioledb" >}} | `1.5` | {{< badge content="Link" link="https://github.com/orioledb/orioledb" >}} | {{< category "FEAT" >}} | `orioledb_$v*` | `oriolepg-$v-orioledb` |
| {{< ext "pg_cardano" >}} | `1.1.1` | {{< badge content="Link" link="https://github.com/Fell-x27/pg_cardano" >}} | {{< category "FEAT" >}} | `pg_cardano_$v` | `postgresql-$v-pg-cardano` |
| {{< ext "rdkit" >}} | `202503.1` | {{< badge content="Link" link="https://github.com/rdkit/rdkit" >}} | {{< category "FEAT" >}} | - | `postgresql-$v-rdkit` |
| {{< ext "omni" "omnigres" >}} | `0.2.14` | {{< badge content="Link" link="https://github.com/omnigres/omnigres" >}} | {{< category "FEAT" >}} | `omnigres_$v` | `postgresql-$v-omnigres` |
| {{< ext "omnisketch" >}} | `1.0.2` | {{< badge content="Link" link="https://github.com/tvondra/omnisketch" >}} | {{< category "FUNC" >}} | `omnisketch_$v*` | `postgresql-$v-omnisketch` |
| {{< ext "bloom" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/bloom.html" >}} | {{< category "FEAT" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "pg_tle" >}} | `1.5.2` | {{< badge content="Link" link="https://github.com/aws/pg_tle" >}} | {{< category "LANG" >}} | `pg_tle_$v*` | `postgresql-$v-pg-tle` |
| {{< ext "plv8" >}} | `3.2.4` | {{< badge content="Link" link="https://github.com/plv8/plv8" >}} | {{< category "LANG" >}} | `plv8_$v*` | `postgresql-$v-plv8` |
| {{< ext "pljs" >}} | `1.0.3` | {{< badge content="Link" link="https://github.com/plv8/pljs" >}} | {{< category "LANG" >}} | - | `postgresql-$v-pljs` |
| {{< ext "pllua" >}} | `2.0.12` | {{< badge content="Link" link="https://github.com/pllua/pllua" >}} | {{< category "LANG" >}} | `pllua_$v*` | `postgresql-$v-pllua` |
| {{< ext "plprql" >}} | `18.0.0` | {{< badge content="Link" link="https://github.com/kaspermarstal/plprql" >}} | {{< category "LANG" >}} | `plprql_$v` | `postgresql-$v-plprql` |
| {{< ext "pldbgapi" "pldebugger" >}} | `1.9` | {{< badge content="Link" link="https://github.com/EnterpriseDB/pldebugger" >}} | {{< category "LANG" >}} | `pldebugger_$v*` | `postgresql-$v-pldebugger` |
| {{< ext "plpgsql_check" >}} | `2.8.3` | {{< badge content="Link" link="https://github.com/okbob/plpgsql_check" >}} | {{< category "LANG" >}} | `plpgsql_check_$v*` | `postgresql-$v-plpgsql-check` |
| {{< ext "plprofiler" >}} | `4.2.5` | {{< badge content="Link" link="https://github.com/bigsql/plprofiler" >}} | {{< category "LANG" >}} | `plprofiler_$v*` | `postgresql-$v-plprofiler` |
| {{< ext "plsh" >}} | `1.20220917` | {{< badge content="Link" link="https://github.com/petere/plsh" >}} | {{< category "LANG" >}} | `plsh_$v*` | `postgresql-$v-plsh` |
| {{< ext "pljava" >}} | `1.6.10` | {{< badge content="Link" link="https://github.com/tada/pljava" >}} | {{< category "LANG" >}} | `pljava_$v*` | `postgresql-$v-pljava` |
| {{< ext "plr" >}} | `8.4.8` | {{< badge content="Link" link="https://github.com/postgres-plr/plr" >}} | {{< category "LANG" >}} | `plr_$v*` | `postgresql-$v-plr` |
| {{< ext "plxslt" >}} | `0.20140221` | {{< badge content="Link" link="https://github.com/petere/plxslt" >}} | {{< category "LANG" >}} | `plxslt_$v*` | `postgresql-$v-plxslt` |
| {{< ext "pgtap" >}} | `1.3.3` | {{< badge content="Link" link="https://github.com/theory/pgtap" >}} | {{< category "LANG" >}} | `pgtap_$v*` | `postgresql-$v-pgtap` |
| {{< ext "faker" >}} | `0.5.3` | {{< badge content="Link" link="https://github.com/anpandu/postgresql_faker" >}} | {{< category "LANG" >}} | `postgresql_faker_$v*` | - |
| {{< ext "dbt2" >}} | `0.61.7` | {{< badge content="Link" link="https://github.com/osdldbt/dbt2" >}} | {{< category "LANG" >}} | `dbt2-pg$v-extensions*` | - |
| {{< ext "pltcl" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/pltcl.html" >}} | {{< category "LANG" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "plperl" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/plperl.html" >}} | {{< category "LANG" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "plperlu" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/plperl.html" >}} | {{< category "LANG" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "plpgsql" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/plpgsql.html" >}} | {{< category "LANG" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "plpython3u" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/plpython.html" >}} | {{< category "LANG" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "prefix" "pg_prefix" >}} | `1.2.10` | {{< badge content="Link" link="https://github.com/dimitri/prefix" >}} | {{< category "TYPE" >}} | `prefix_$v*` | `postgresql-$v-prefix` |
| {{< ext "semver" "pg_semver" >}} | `0.40.0` | {{< badge content="Link" link="https://github.com/theory/pg-semver" >}} | {{< category "TYPE" >}} | `semver_$v*` | `postgresql-$v-semver` |
| {{< ext "unit" "pgunit" >}} | `7.10` | {{< badge content="Link" link="https://github.com/df7cb/postgresql-unit" >}} | {{< category "TYPE" >}} | `postgresql-unit_$v*` | `postgresql-$v-unit` |
| {{< ext "pgpdf" >}} | `0.1.0` | {{< badge content="Link" link="https://github.com/Florents-Tselai/pgpdf" >}} | {{< category "TYPE" >}} | `pgpdf_$v*` | `postgresql-$v-pgpdf` |
| {{< ext "pglite_fusion" >}} | `0.0.5` | {{< badge content="Link" link="https://github.com/frectonz/pglite-fusion" >}} | {{< category "TYPE" >}} | `pglite_fusion_$v` | `postgresql-$v-pglite-fusion` |
| {{< ext "md5hash" >}} | `1.0.1` | {{< badge content="Link" link="https://github.com/tvondra/md5hash" >}} | {{< category "TYPE" >}} | `md5hash_$v*` | `postgresql-$v-md5hash` |
| {{< ext "asn1oid" >}} | `1.6` | {{< badge content="Link" link="https://github.com/df7cb/pgsql-asn1oid" >}} | {{< category "TYPE" >}} | `asn1oid_$v*` | `postgresql-$v-asn1oid` |
| {{< ext "roaringbitmap" "pg_roaringbitmap" >}} | `0.5.5` | {{< badge content="Link" link="https://github.com/ChenHuajun/pg_roaringbitmap" >}} | {{< category "TYPE" >}} | `pg_roaringbitmap_$v*` | `postgresql-$v-roaringbitmap` |
| {{< ext "pgfaceting" >}} | `0.2.0` | {{< badge content="Link" link="https://github.com/cybertec-postgresql/pgfaceting" >}} | {{< category "TYPE" >}} | `pgfaceting_$v` | `postgresql-$v-pgfaceting` |
| {{< ext "pg_sphere" "pgsphere" >}} | `1.5.2` | {{< badge content="Link" link="https://github.com/postgrespro/pgsphere" >}} | {{< category "TYPE" >}} | `pgsphere_$v*` | `postgresql-$v-pgsphere` |
| {{< ext "country" "pg_country" >}} | `0.0.3` | {{< badge content="Link" link="https://github.com/adjust/pg-country" >}} | {{< category "TYPE" >}} | `pg_country_$v*` | `postgresql-$v-pg-country` |
| {{< ext "pg_xenophile" >}} | `0.8.3` | {{< badge content="Link" link="https://github.com/bigsmoke/pg_xenophile" >}} | {{< category "TYPE" >}} | `pg_xenophile_$v` | `postgresql-$v-pg-xenophile` |
| {{< ext "currency" "pg_currency" >}} | `0.0.3` | {{< badge content="Link" link="https://github.com/adjust/pg-currency" >}} | {{< category "TYPE" >}} | `pg_currency_$v*` | `postgresql-$v-pg-currency` |
| {{< ext "collection" "pgcollection" >}} | `1.1.0` | {{< badge content="Link" link="https://github.com/aws/pgcollection" >}} | {{< category "TYPE" >}} | `pgcollection_$v*` | `postgresql-$v-collection` |
| {{< ext "pgmp" >}} | `1.0.5` | {{< badge content="Link" link="https://github.com/dvarrazzo/pgmp/" >}} | {{< category "TYPE" >}} | `pgmp_$v*` | `postgresql-$v-pgmp` |
| {{< ext "numeral" >}} | `1.3` | {{< badge content="Link" link="https://github.com/df7cb/postgresql-numeral" >}} | {{< category "TYPE" >}} | `numeral_$v*` | `postgresql-$v-numeral` |
| {{< ext "pg_rational" >}} | `0.0.2` | {{< badge content="Link" link="https://github.com/begriffs/pg_rational" >}} | {{< category "TYPE" >}} | `pg_rational_$v*` | `postgresql-$v-rational` |
| {{< ext "uint" "pguint" >}} | `1.20250815` | {{< badge content="Link" link="https://github.com/petere/pguint" >}} | {{< category "TYPE" >}} | `pguint_$v*` | `postgresql-$v-pguint` |
| {{< ext "uint128" "pg_uint128" >}} | `1.1.1` | {{< badge content="Link" link="https://github.com/pg-uint/pg-uint128" >}} | {{< category "TYPE" >}} | `pg_uint128_$v*` | `postgresql-$v-pg-uint128` |
| {{< ext "hashtypes" >}} | `0.1.5` | {{< badge content="Link" link="https://github.com/adjust/hashtypes/" >}} | {{< category "TYPE" >}} | `hashtypes_$v*` | `postgresql-$v-hashtypes` |
| {{< ext "ip4r" >}} | `2.4.2` | {{< badge content="Link" link="https://github.com/RhodiumToad/ip4r" >}} | {{< category "TYPE" >}} | `ip4r_$v*` | `postgresql-$v-ip4r` |
| {{< ext "pg_duration" >}} | `1.0.2` | {{< badge content="Link" link="https://github.com/jkosh44/pg_duration" >}} | {{< category "TYPE" >}} | `pg_duration_$v*` | `postgresql-$v-pg-duration` |
| {{< ext "uri" "pg_uri" >}} | `1.20151224` | {{< badge content="Link" link="https://github.com/petere/pguri" >}} | {{< category "TYPE" >}} | `pg_uri_$v*` | `postgresql-$v-pg-uri` |
| {{< ext "emailaddr" "pg_emailaddr" >}} | `0` | {{< badge content="Link" link="https://github.com/petere/pgemailaddr" >}} | {{< category "TYPE" >}} | `pg_emailaddr_$v*` | `postgresql-$v-pg-emailaddr` |
| {{< ext "acl" "pg_acl" >}} | `1.0.4` | {{< badge content="Link" link="https://github.com/arkhipov/acl" >}} | {{< category "TYPE" >}} | `acl_$v*` | `postgresql-$v-acl` |
| {{< ext "debversion" >}} | `1.2.0` | {{< badge content="Link" link="https://github.com/ATIX-AG/postgresql-debversion-evr" >}} | {{< category "TYPE" >}} | - | `postgresql-$v-debversion` |
| {{< ext "pg_rrule" >}} | `0.2.0` | {{< badge content="Link" link="https://github.com/petropavel13/pg_rrule" >}} | {{< category "TYPE" >}} | `pg_rrule_$v` | `postgresql-$v-pg-rrule` |
| {{< ext "timestamp9" >}} | `1.4.0` | {{< badge content="Link" link="https://github.com/optiver/timestamp9" >}} | {{< category "TYPE" >}} | `timestamp9_$v*` | `postgresql-$v-timestamp9` |
| {{< ext "chkpass" >}} | `1.0` | {{< badge content="Link" link="https://github.com/lacanoid/chkpass" >}} | {{< category "TYPE" >}} | `chkpass_$v*` | `postgresql-$v-chkpass` |
| {{< ext "isn" >}} | `1.2` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/isn.html" >}} | {{< category "TYPE" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "seg" >}} | `1.4` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/seg.html" >}} | {{< category "TYPE" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "cube" >}} | `1.5` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/cube.html" >}} | {{< category "TYPE" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "ltree" >}} | `1.3` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/ltree.html" >}} | {{< category "TYPE" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "hstore" >}} | `1.8` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/hstore.html" >}} | {{< category "TYPE" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "citext" >}} | `1.6` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/citext.html" >}} | {{< category "TYPE" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "xml2" >}} | `1.1` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/xml2.html" >}} | {{< category "TYPE" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "gzip" "pg_gzip" >}} | `1.0.0` | {{< badge content="Link" link="https://github.com/pramsey/pgsql-gzip" >}} | {{< category "UTIL" >}} | `pg_gzip_$v*` | `postgresql-$v-gzip` |
| {{< ext "bzip" "pg_bzip" >}} | `1.0.0` | {{< badge content="Link" link="https://github.com/steve-chavez/pg_bzip" >}} | {{< category "UTIL" >}} | `pg_bzip_$v*` | `postgresql-$v-bzip` |
| {{< ext "zstd" "pg_zstd" >}} | `1.1.2` | {{< badge content="Link" link="https://github.com/grahamedgecombe/pgzstd" >}} | {{< category "UTIL" >}} | `pg_zstd_$v*` | `postgresql-$v-zstd` |
| {{< ext "http" "pg_http" >}} | `1.7.0` | {{< badge content="Link" link="https://github.com/pramsey/pgsql-http" >}} | {{< category "UTIL" >}} | `pg_http_$v*` | `postgresql-$v-http` |
| {{< ext "pg_net" >}} | `0.20.0` | {{< badge content="Link" link="https://github.com/supabase/pg_net" >}} | {{< category "UTIL" >}} | `pg_net_$v*` | `postgresql-$v-pg-net` |
| {{< ext "pg_curl" >}} | `2.4.5` | {{< badge content="Link" link="https://github.com/RekGRpth/pg_curl" >}} | {{< category "UTIL" >}} | `pg_curl_$v*` | `postgresql-$v-pg-curl` |
| {{< ext "pgjq" >}} | `0.1.0` | {{< badge content="Link" link="https://github.com/Florents-Tselai/pgJQ" >}} | {{< category "UTIL" >}} | `pgjq_$v*` | `postgresql-$v-pgjq` |
| {{< ext "pgjwt" >}} | `0.2.0` | {{< badge content="Link" link="https://github.com/michelp/pgjwt" >}} | {{< category "UTIL" >}} | `pgjwt_$v` | `postgresql-$v-pgjwt` |
| {{< ext "pg_smtp_client" >}} | `0.2.1` | {{< badge content="Link" link="https://github.com/brianpursley/pg_smtp_client" >}} | {{< category "UTIL" >}} | `pg_smtp_client_$v` | `postgresql-$v-pg-smtp-client` |
| {{< ext "pg_html5_email_address" >}} | `1.2.3` | {{< badge content="Link" link="https://github.com/bigsmoke/pg_html5_email_address" >}} | {{< category "UTIL" >}} | `pg_html5_email_address_$v` | `postgresql-$v-pg-html5-email-address` |
| {{< ext "url_encode" >}} | `1.2.5` | {{< badge content="Link" link="https://github.com/okbob/url_encode" >}} | {{< category "UTIL" >}} | `url_encode_$v*` | `postgresql-$v-url-encode` |
| {{< ext "pgsql_tweaks" >}} | `1.0.2` | {{< badge content="Link" link="https://codeberg.org/pgsql_tweaks/pgsql_tweaks" >}} | {{< category "UTIL" >}} | `pgsql_tweaks_$v` | `postgresql-$v-pgsql-tweaks` |
| {{< ext "pg_extra_time" >}} | `2.0.0` | {{< badge content="Link" link="https://github.com/bigsmoke/pg_extra_time" >}} | {{< category "UTIL" >}} | `pg_extra_time_$v` | `postgresql-$v-pg-extra-time` |
| {{< ext "pgpcre" >}} | `0.20190509` | {{< badge content="Link" link="https://github.com/petere/pgpcre" >}} | {{< category "UTIL" >}} | `pgpcre_$v` | `postgresql-$v-pgpcre` |
| {{< ext "icu_ext" >}} | `1.10.0` | {{< badge content="Link" link="https://github.com/dverite/icu_ext" >}} | {{< category "UTIL" >}} | `icu_ext_$v*` | `postgresql-$v-icu-ext` |
| {{< ext "pgqr" >}} | `1.0` | {{< badge content="Link" link="https://github.com/AbdulYadi/pgqr" >}} | {{< category "UTIL" >}} | `pgqr_$v*` | `postgresql-$v-pgqr` |
| {{< ext "pg_protobuf" >}} | `1.0` | {{< badge content="Link" link="https://github.com/afiskon/pg_protobuf" >}} | {{< category "UTIL" >}} | `pg_protobuf_$v` | `postgresql-$v-pg-protobuf` |
| {{< ext "envvar" "pg_envvar" >}} | `1.0.1` | {{< badge content="Link" link="https://github.com/theory/pg-envvar" >}} | {{< category "UTIL" >}} | `pg_envvar_$v*` | `postgresql-$v-pg-envvar` |
| {{< ext "floatfile" >}} | `1.3.1` | {{< badge content="Link" link="https://github.com/pjungwir/floatfile" >}} | {{< category "UTIL" >}} | `floatfile_$v*` | `postgresql-$v-floatfile` |
| {{< ext "pg_render" >}} | `0.1.3` | {{< badge content="Link" link="https://github.com/mkaski/pg_render" >}} | {{< category "UTIL" >}} | `pg_render_$v` | `postgresql-$v-pg-render` |
| {{< ext "pg_readme" >}} | `0.7.0` | {{< badge content="Link" link="https://github.com/bigsmoke/pg_readme" >}} | {{< category "UTIL" >}} | `pg_readme_$v` | `postgresql-$v-pg-readme` |
| {{< ext "ddl_historization" >}} | `0.0.7` | {{< badge content="Link" link="https://github.com/rodo/pg_ddl_historization" >}} | {{< category "UTIL" >}} | `ddl_historization_$v` | `postgresql-$v-ddl-historization` |
| {{< ext "data_historization" >}} | `1.1.0` | {{< badge content="Link" link="https://github.com/rodo/postgresql-data-historization" >}} | {{< category "UTIL" >}} | `data_historization_$v` | `postgresql-$v-data-historization` |
| {{< ext "schedoc" "pg_schedoc" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/ZeroGachis/pg_schedoc" >}} | {{< category "UTIL" >}} | `pg_schedoc_$v` | `postgresql-$v-pg-schedoc` |
| {{< ext "hashlib" "pg_hashlib" >}} | `1.1` | {{< badge content="Link" link="https://github.com/markokr/pghashlib" >}} | {{< category "UTIL" >}} | `pg_hashlib_$v` | `postgresql-$v-pg-hashlib` |
| {{< ext "xxhash" "pg_xxhash" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/hatarist/pg_xxhash" >}} | {{< category "UTIL" >}} | `pg_xxhash_$v*` | `postgresql-$v-pg-xxhash` |
| {{< ext "shacrypt" >}} | `1.1` | {{< badge content="Link" link="https://github.com/dverite/postgres-shacrypt" >}} | {{< category "UTIL" >}} | `shacrypt_$v*` | `postgresql-$v-shacrypt` |
| {{< ext "cryptint" >}} | `1.0.0` | {{< badge content="Link" link="https://github.com/dverite/cryptint" >}} | {{< category "UTIL" >}} | `cryptint_$v*` | `postgresql-$v-cryptint` |
| {{< ext "pguecc" "pg_ecdsa" >}} | `1.0` | {{< badge content="Link" link="https://github.com/ameensol/pg-ecdsa" >}} | {{< category "UTIL" >}} | `pg_ecdsa_$v*` | `postgresql-$v-pg-ecdsa` |
| {{< ext "sparql" "pgsparql" >}} | `1.0` | {{< badge content="Link" link="https://github.com/lacanoid/pgsparql" >}} | {{< category "UTIL" >}} | `pgsparql_$v` | `postgresql-$v-pgsparql` |
| {{< ext "pg_idkit" >}} | `0.4.0` | {{< badge content="Link" link="https://github.com/VADOSWARE/pg_idkit" >}} | {{< category "FUNC" >}} | `pg_idkit_$v` | `postgresql-$v-pg-idkit` |
| {{< ext "pgx_ulid" >}} | `0.2.1` | {{< badge content="Link" link="https://github.com/pksunkara/pgx_ulid" >}} | {{< category "FUNC" >}} | `pgx_ulid_$v` | `postgresql-$v-pgx-ulid` |
| {{< ext "pg_uuidv7" >}} | `1.6.0` | {{< badge content="Link" link="https://github.com/fboulnois/pg_uuidv7" >}} | {{< category "FUNC" >}} | `pg_uuidv7_$v*` | `postgresql-$v-pg-uuidv7` |
| {{< ext "permuteseq" >}} | `1.2.2` | {{< badge content="Link" link="https://github.com/dverite/permuteseq" >}} | {{< category "FUNC" >}} | `permuteseq_$v*` | `postgresql-$v-permuteseq` |
| {{< ext "pg_hashids" >}} | `1.3` | {{< badge content="Link" link="https://github.com/iCyberon/pg_hashids" >}} | {{< category "FUNC" >}} | `pg_hashids_$v*` | `postgresql-$v-pg-hashids` |
| {{< ext "sequential_uuids" >}} | `1.0.3` | {{< badge content="Link" link="https://github.com/tvondra/sequential-uuids" >}} | {{< category "FUNC" >}} | `sequential_uuids_$v` | `postgresql-$v-sequential-uuids` |
| {{< ext "topn" >}} | `2.7.0` | {{< badge content="Link" link="https://github.com/citusdata/postgresql-topn" >}} | {{< category "FUNC" >}} | `topn_$v*` | `postgresql-$v-topn` |
| {{< ext "quantile" >}} | `1.1.8` | {{< badge content="Link" link="https://github.com/tvondra/quantile" >}} | {{< category "FUNC" >}} | `quantile_$v*` | `postgresql-$v-quantile` |
| {{< ext "lower_quantile" >}} | `1.0.3` | {{< badge content="Link" link="https://github.com/tvondra/lower_quantile" >}} | {{< category "FUNC" >}} | `lower_quantile_$v*` | `postgresql-$v-lower-quantile` |
| {{< ext "count_distinct" >}} | `3.0.2` | {{< badge content="Link" link="https://github.com/tvondra/count_distinct" >}} | {{< category "FUNC" >}} | `count_distinct_$v*` | `postgresql-$v-count-distinct` |
| {{< ext "ddsketch" >}} | `1.0.1` | {{< badge content="Link" link="https://github.com/tvondra/ddsketch" >}} | {{< category "FUNC" >}} | `ddsketch_$v*` | `postgresql-$v-ddsketch` |
| {{< ext "vasco" >}} | `0.1.0` | {{< badge content="Link" link="https://github.com/Florents-Tselai/vasco" >}} | {{< category "FUNC" >}} | `vasco_$v*` | `postgresql-$v-vasco` |
| {{< ext "xicor" "pgxicor" >}} | `0.1.0` | {{< badge content="Link" link="https://github.com/Florents-Tselai/pgxicor" >}} | {{< category "FUNC" >}} | `pgxicor_$v*` | `postgresql-$v-pgxicor` |
| {{< ext "tdigest" >}} | `1.4.3` | {{< badge content="Link" link="https://github.com/tvondra/tdigest" >}} | {{< category "FUNC" >}} | `tdigest_$v*` | `postgresql-$v-tdigest` |
| {{< ext "first_last_agg" >}} | `0.1.4` | {{< badge content="Link" link="https://github.com/wulczer/first_last_agg" >}} | {{< category "FUNC" >}} | `first_last_agg_$v` | `postgresql-$v-first-last-agg` |
| {{< ext "extra_window_functions" >}} | `1.0` | {{< badge content="Link" link="https://github.com/xocolatl/extra_window_functions" >}} | {{< category "FUNC" >}} | `extra_window_functions_$v*` | `postgresql-$v-extra-window-functions` |
| {{< ext "floatvec" >}} | `1.1.1` | {{< badge content="Link" link="https://github.com/pjungwir/floatvec" >}} | {{< category "FUNC" >}} | `floatvec_$v*` | `postgresql-$v-floatvec` |
| {{< ext "aggs_for_vecs" >}} | `1.4.0` | {{< badge content="Link" link="https://github.com/pjungwir/aggs_for_vecs" >}} | {{< category "FUNC" >}} | `aggs_for_vecs_$v*` | `postgresql-$v-aggs-for-vecs` |
| {{< ext "aggs_for_arrays" >}} | `1.3.3` | {{< badge content="Link" link="https://github.com/pjungwir/aggs_for_arrays" >}} | {{< category "FUNC" >}} | `aggs_for_arrays_$v*` | `postgresql-$v-aggs-for-arrays` |
| {{< ext "pg_csv" >}} | `1.0.1` | {{< badge content="Link" link="https://github.com/PostgREST/pg_csv" >}} | {{< category "FUNC" >}} | `pg_csv_$v*` | `postgresql-$v-pg-csv` |
| {{< ext "arraymath" "pg_arraymath" >}} | `1.1` | {{< badge content="Link" link="https://github.com/pramsey/pgsql-arraymath" >}} | {{< category "FUNC" >}} | `pg_arraymath_$v*` | `postgresql-$v-pg-arraymath` |
| {{< ext "pg_math" >}} | `1.0` | {{< badge content="Link" link="https://github.com/chanukyasds/pg_math" >}} | {{< category "FUNC" >}} | `pg_math_$v*` | `postgresql-$v-pg-math` |
| {{< ext "random" "pg_random" >}} | `2.0.0` | {{< badge content="Link" link="https://github.com/tvondra/random" >}} | {{< category "FUNC" >}} | `pg_random_$v*` | `postgresql-$v-random` |
| {{< ext "base36" "pg_base36" >}} | `1.0.0` | {{< badge content="Link" link="https://github.com/adjust/pg-base36" >}} | {{< category "FUNC" >}} | `pg_base36_$v*` | `postgresql-$v-base36` |
| {{< ext "base62" "pg_base62" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/adjust/pg-base62" >}} | {{< category "FUNC" >}} | `pg_base62_$v*` | `postgresql-$v-base62` |
| {{< ext "pg_base58" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/Fell-x27/pg_base58" >}} | {{< category "FUNC" >}} | `pg_base58_$v` | `postgresql-$v-pg-base58` |
| {{< ext "financial" "pg_financial" >}} | `1.0.1` | {{< badge content="Link" link="https://github.com/intgr/pg_financial" >}} | {{< category "FUNC" >}} | `pg_financial_$v*` | `postgresql-$v-pg-financial` |
| {{< ext "convert" "pg_convert" >}} | `0.0.4` | {{< badge content="Link" link="https://github.com/rustprooflabs/convert" >}} | {{< category "FUNC" >}} | `pg_convert_$v` | `postgresql-$v-convert` |
| {{< ext "refint" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/contrib-spi.html#CONTRIB-SPI-REFINT" >}} | {{< category "FUNC" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "autoinc" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/contrib-spi.html#CONTRIB-SPI-AUTOINC" >}} | {{< category "FUNC" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "insert_username" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/contrib-spi.html#CONTRIB-SPI-INSERT-USERNAME" >}} | {{< category "FUNC" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "moddatetime" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/contrib-spi.html#CONTRIB-SPI-MODDATETIME" >}} | {{< category "FUNC" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "tsm_system_time" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/tsm-system-time.html" >}} | {{< category "FUNC" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "dict_xsyn" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/dict-xsyn.html" >}} | {{< category "FUNC" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "tsm_system_rows" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/tsm-system-rows.html" >}} | {{< category "FUNC" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "tcn" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/tcn.html" >}} | {{< category "FUNC" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "uuid-ossp" >}} | `1.1` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/uuid-ossp.html" >}} | {{< category "FUNC" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "btree_gist" >}} | `1.7` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/btree-gist.html" >}} | {{< category "FUNC" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "btree_gin" >}} | `1.3` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/btree-gin.html" >}} | {{< category "FUNC" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "intarray" >}} | `1.5` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/intarray.html" >}} | {{< category "FUNC" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "intagg" >}} | `1.1` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/intagg.html" >}} | {{< category "FUNC" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "dict_int" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/dict-int.html" >}} | {{< category "FUNC" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "unaccent" >}} | `1.1` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/unaccent.html" >}} | {{< category "FUNC" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "pg_repack" >}} | `1.5.2` | {{< badge content="Link" link="https://github.com/reorg/pg_repack" >}} | {{< category "ADMIN" >}} | `pg_repack_$v*` | `postgresql-$v-repack` |
| {{< ext "pg_rewrite" >}} | `2.0.0` | {{< badge content="Link" link="https://github.com/cybertec-postgresql/pg_rewrite" >}} | {{< category "ADMIN" >}} | `pg_rewrite_$v*` | `postgresql-$v-pg-rewrite` |
| {{< ext "pg_squeeze" >}} | `1.9.0` | {{< badge content="Link" link="https://github.com/cybertec-postgresql/pg_squeeze" >}} | {{< category "ADMIN" >}} | `pg_squeeze_$v*` | `postgresql-$v-squeeze` |
| {{< ext "pg_dirtyread" >}} | `2.7` | {{< badge content="Link" link="https://github.com/df7cb/pg_dirtyread" >}} | {{< category "ADMIN" >}} | `pg_dirtyread_$v*` | `postgresql-$v-dirtyread` |
| {{< ext "pgfincore" >}} | `1.3.1` | {{< badge content="Link" link="https://github.com/klando/pgfincore" >}} | {{< category "ADMIN" >}} | `pgfincore_$v*` | `postgresql-$v-pgfincore` |
| {{< ext "pg_cooldown" >}} | `0.1` | {{< badge content="Link" link="https://github.com/rbergm/pg_cooldown" >}} | {{< category "ADMIN" >}} | `pg_cooldown_$v*` | `postgresql-$v-pg-cooldown` |
| {{< ext "ddlx" "pg_ddlx" >}} | `0.30` | {{< badge content="Link" link="https://github.com/lacanoid/pgddl" >}} | {{< category "ADMIN" >}} | `ddlx_$v` | `postgresql-$v-ddlx` |
| {{< ext "prioritize" "pg_prioritize" >}} | `1.0.4` | {{< badge content="Link" link="https://github.com/schmiddy/pg_prioritize" >}} | {{< category "ADMIN" >}} | `pg_prioritize_$v*` | `postgresql-$v-prioritize` |
| {{< ext "pg_checksums" >}} | `1.3` | {{< badge content="Link" link="https://github.com/credativ/pg_checksums" >}} | {{< category "ADMIN" >}} | `pg_checksums_$v*` | `postgresql-$v-pg-checksums` |
| {{< ext "pg_readonly" >}} | `1.0.3` | {{< badge content="Link" link="https://github.com/pierreforstmann/pg_readonly" >}} | {{< category "ADMIN" >}} | `pg_readonly_$v*` | `postgresql-$v-pg-readonly` |
| {{< ext "pgdd" >}} | `0.6.0` | {{< badge content="Link" link="https://github.com/rustprooflabs/pgdd" >}} | {{< category "ADMIN" >}} | `pgdd_$v` | `postgresql-$v-pgdd` |
| {{< ext "pg_permissions" >}} | `1.4` | {{< badge content="Link" link="https://github.com/cybertec-postgresql/pg_permissions" >}} | {{< category "ADMIN" >}} | `pg_permissions_$v` | `postgresql-$v-pg-permissions` |
| {{< ext "pgautofailover" >}} | `2.2` | {{< badge content="Link" link="https://github.com/hapostgres/pg_auto_failover" >}} | {{< category "ADMIN" >}} | `pg_auto_failover_$v*` | `postgresql-$v-auto-failover` |
| {{< ext "pg_catcheck" >}} | `1.6.0` | {{< badge content="Link" link="https://github.com/EnterpriseDB/pg_catcheck" >}} | {{< category "ADMIN" >}} | `pg_catcheck_$v*` | `postgresql-$v-pg-catcheck` |
| {{< ext "pre_prepare" "preprepare" >}} | `0.9` | {{< badge content="Link" link="https://github.com/dimitri/preprepare" >}} | {{< category "ADMIN" >}} | `preprepare_$v*` | `postgresql-$v-preprepare` |
| {{< ext "pg_upless" >}} | `0.0.3` | {{< badge content="Link" link="https://github.com/rodo/pg_upless" >}} | {{< category "ADMIN" >}} | `pg_upless_$v` | `postgresql-$v-pg-upless` |
| {{< ext "pgcozy" >}} | `1.0` | {{< badge content="Link" link="https://github.com/vventirozos/pgcozy" >}} | {{< category "ADMIN" >}} | `pgcozy_$v` | `postgresql-$v-pgcozy` |
| {{< ext "pg_orphaned" >}} | `1.0` | {{< badge content="Link" link="https://github.com/bdrouvot/pg_orphaned" >}} | {{< category "ADMIN" >}} | `pg_orphaned_$v*` | `postgresql-$v-pg-orphaned` |
| {{< ext "pg_crash" >}} | `1.0` | {{< badge content="Link" link="https://github.com/cybertec-postgresql/pg_crash" >}} | {{< category "ADMIN" >}} | `pg_crash_$v*` | `postgresql-$v-pg-crash` |
| {{< ext "pg_cheat_funcs" >}} | `1.0` | {{< badge content="Link" link="https://github.com/MasaoFujii/pg_cheat_funcs" >}} | {{< category "ADMIN" >}} | `pg_cheat_funcs_$v*` | `postgresql-$v-pg-cheat-funcs` |
| {{< ext "fio" "pg_fio" >}} | `1.0` | {{< badge content="Link" link="https://github.com/csimsek/pgsql-fio" >}} | {{< category "ADMIN" >}} | `pg_fio_$v` | `postgresql-$v-pg-fio` |
| {{< ext "pg_savior" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/viggy28/pg_savior" >}} | {{< category "ADMIN" >}} | `pg_savior_$v*` | `postgresql-$v-pg-savior` |
| {{< ext "safeupdate" >}} | `1.5` | {{< badge content="Link" link="https://github.com/eradman/pg-safeupdate" >}} | {{< category "ADMIN" >}} | `safeupdate_$v*` | `postgresql-$v-pg-safeupdate` |
| {{< ext "pg_drop_events" >}} | `0.1.0` | {{< badge content="Link" link="https://github.com/bolajiwahab/pg_drop_events" >}} | {{< category "ADMIN" >}} | `pg_drop_events_$v` | `postgresql-$v-pg-drop-events` |
| {{< ext "table_log" >}} | `0.6.4` | {{< badge content="Link" link="https://github.com/df7cb/table_log" >}} | {{< category "ADMIN" >}} | `table_log_$v` | `postgresql-$v-tablelog` |
| {{< ext "pgagent" >}} | `4.2.3` | {{< badge content="Link" link="https://www.pgadmin.org/docs/pgadmin4/development/pgagent.html" >}} | {{< category "ADMIN" >}} | `pgagent_$v*` | `pgagent` |
| {{< ext "pg_prewarm" >}} | `1.2` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/pgprewarm.html" >}} | {{< category "ADMIN" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "pgpool_adm" "pgpool" >}} | `4.6.3` | {{< badge content="Link" link="https://pgpool.net/" >}} | {{< category "ADMIN" >}} | `pgpool-II-pg$v-extensions` | `postgresql-$v-pgpool2` |
| {{< ext "lo" >}} | `1.1` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/lo.html" >}} | {{< category "ADMIN" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "basic_archive" >}} | `-` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/basic-archive.html" >}} | {{< category "ADMIN" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "basebackup_to_shell" >}} | `-` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/basebackup-to-shell.html" >}} | {{< category "ADMIN" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "old_snapshot" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/oldsnapshot.html" >}} | {{< category "ADMIN" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "adminpack" >}} | `2.1` | {{< badge content="Link" link="https://www.postgresql.org/docs/16/adminpack.html" >}} | {{< category "ADMIN" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "amcheck" >}} | `1.4` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/amcheck.html" >}} | {{< category "ADMIN" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "pg_surgery" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/pgsurgery.html" >}} | {{< category "ADMIN" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "pg_profile" >}} | `4.10` | {{< badge content="Link" link="https://github.com/zubkov-andrei/pg_profile" >}} | {{< category "STAT" >}} | `pg_profile_$v*` | `postgresql-$v-pg-profile` |
| {{< ext "pg_tracing" >}} | `0.1.3` | {{< badge content="Link" link="https://github.com/DataDog/pg_tracing" >}} | {{< category "STAT" >}} | `pg_tracing_$v*` | `postgresql-$v-pg-tracing` |
| {{< ext "pg_show_plans" >}} | `2.1.7` | {{< badge content="Link" link="https://github.com/cybertec-postgresql/pg_show_plans" >}} | {{< category "STAT" >}} | `pg_show_plans_$v*` | `postgresql-$v-show-plans` |
| {{< ext "pg_stat_kcache" >}} | `2.3.0` | {{< badge content="Link" link="https://github.com/powa-team/pg_stat_kcache" >}} | {{< category "STAT" >}} | `pg_stat_kcache_$v*` | `postgresql-$v-pg-stat-kcache` |
| {{< ext "pg_stat_monitor" >}} | `2.2.0` | {{< badge content="Link" link="https://github.com/percona/pg_stat_monitor" >}} | {{< category "STAT" >}} | `pg_stat_monitor_$v*` | `postgresql-$v-pg-stat-monitor` |
| {{< ext "pg_qualstats" >}} | `2.1.2` | {{< badge content="Link" link="https://github.com/powa-team/pg_qualstats" >}} | {{< category "STAT" >}} | `pg_qualstats_$v*` | `postgresql-$v-pg-qualstats` |
| {{< ext "pg_store_plans" >}} | `1.9` | {{< badge content="Link" link="https://github.com/ossc-db/pg_store_plans" >}} | {{< category "STAT" >}} | `pg_store_plans_$v*` | `postgresql-$v-pg-store-plan` |
| {{< ext "pg_track_settings" >}} | `2.1.2` | {{< badge content="Link" link="https://github.com/rjuju/pg_track_settings" >}} | {{< category "STAT" >}} | `pg_track_settings_$v` | `postgresql-$v-pg-track-settings` |
| {{< ext "pg_wait_sampling" >}} | `1.1.9` | {{< badge content="Link" link="https://github.com/postgrespro/pg_wait_sampling" >}} | {{< category "STAT" >}} | `pg_wait_sampling_$v*` | `postgresql-$v-pg-wait-sampling` |
| {{< ext "pgsentinel" >}} | `1.2.0` | {{< badge content="Link" link="https://github.com/pgsentinel/pgsentinel" >}} | {{< category "STAT" >}} | `pgsentinel_$v*` | `postgresql-$v-pgsentinel` |
| {{< ext "system_stats" >}} | `3.2` | {{< badge content="Link" link="https://github.com/EnterpriseDB/system_stats" >}} | {{< category "STAT" >}} | `system_stats_$v*` | `postgresql-$v-system-stats` |
| {{< ext "meta" "pg_meta" >}} | `0.4.0` | {{< badge content="Link" link="https://github.com/aquameta/meta" >}} | {{< category "STAT" >}} | `pg_meta_$v` | `postgresql-$v-pg-meta` |
| {{< ext "pgnodemx" >}} | `1.7` | {{< badge content="Link" link="https://github.com/CrunchyData/pgnodemx" >}} | {{< category "STAT" >}} | `pgnodemx_$v` | `postgresql-$v-pgnodemx` |
| {{< ext "pg_sqlog" >}} | `1.6` | {{< badge content="Link" link="https://github.com/kouber/pg_sqlog" >}} | {{< category "STAT" >}} | `pg_sqlog_$v` | `postgresql-$v-pg-sqlog` |
| {{< ext "bgw_replstatus" >}} | `1.0.8` | {{< badge content="Link" link="https://github.com/mhagander/bgw_replstatus" >}} | {{< category "STAT" >}} | `bgw_replstatus_$v*` | `postgresql-$v-bgw-replstatus` |
| {{< ext "pgmeminfo" >}} | `1.0.0` | {{< badge content="Link" link="https://github.com/okbob/pgmeminfo" >}} | {{< category "STAT" >}} | `pgmeminfo_$v*` | `postgresql-$v-pgmeminfo` |
| {{< ext "toastinfo" >}} | `1.5` | {{< badge content="Link" link="https://github.com/credativ/toastinfo" >}} | {{< category "STAT" >}} | `toastinfo_$v*` | `postgresql-$v-toastinfo` |
| {{< ext "explain_ui" "pg_explain_ui" >}} | `0.0.2` | {{< badge content="Link" link="https://github.com/davidgomes/pg-explain-ui" >}} | {{< category "STAT" >}} | `pg_explain_ui_$v` | `postgresql-$v-pg-explain-ui` |
| {{< ext "pg_relusage" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/adept/pg_relusage" >}} | {{< category "STAT" >}} | `pg_relusage_$v` | `postgresql-$v-pg-relusage` |
| {{< ext "pagevis" >}} | `0.1` | {{< badge content="Link" link="https://github.com/hollobon/pagevis" >}} | {{< category "STAT" >}} | `pagevis_$v` | `postgresql-$v-pagevis` |
| {{< ext "powa" >}} | `5.0.1` | {{< badge content="Link" link="https://github.com/powa-team/powa" >}} | {{< category "STAT" >}} | `powa_$v*` | `postgresql-$v-powa` |
| {{< ext "pg_overexplain" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/devel/pgoverexplain.html" >}} | {{< category "STAT" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "pg_logicalinspect" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/devel/pglogicalinspect.html" >}} | {{< category "STAT" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "pageinspect" >}} | `1.12` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/pageinspect.html" >}} | {{< category "STAT" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "pgrowlocks" >}} | `1.2` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/pgrowlocks.html" >}} | {{< category "STAT" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "sslinfo" >}} | `1.2` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/sslinfo.html" >}} | {{< category "STAT" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "pg_buffercache" >}} | `1.5` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/pgbuffercache.html" >}} | {{< category "STAT" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "pg_walinspect" >}} | `1.1` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/pgwalinspect.html" >}} | {{< category "STAT" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "pg_freespacemap" >}} | `1.2` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/pgfreespacemap.html" >}} | {{< category "STAT" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "pg_visibility" >}} | `1.2` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/pgvisibility.html" >}} | {{< category "STAT" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "pgstattuple" >}} | `1.5` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/pgstattuple.html" >}} | {{< category "STAT" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "auto_explain" >}} | `-` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/auto-explain.html" >}} | {{< category "STAT" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "pg_stat_statements" >}} | `1.11` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/pgstatstatements.html" >}} | {{< category "STAT" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "passwordcheck_cracklib" >}} | `3.1.0` | {{< badge content="Link" link="https://github.com/devrimgunduz/passwordcheck_cracklib" >}} | {{< category "SEC" >}} | `passwordcheck_cracklib_$v*` | `postgresql-$v-passwordcheck-cracklib` |
| {{< ext "supautils" >}} | `3.0.2` | {{< badge content="Link" link="https://github.com/supabase/supautils" >}} | {{< category "SEC" >}} | `supautils_$v` | `postgresql-$v-supautils` |
| {{< ext "pgsodium" >}} | `3.1.9` | {{< badge content="Link" link="https://github.com/michelp/pgsodium" >}} | {{< category "SEC" >}} | `pgsodium_$v*` | `postgresql-$v-pgsodium` |
| {{< ext "supabase_vault" "pg_vault" >}} | `0.3.1` | {{< badge content="Link" link="https://github.com/supabase/vault" >}} | {{< category "SEC" >}} | `vault_$v*` | `postgresql-$v-vault` |
| {{< ext "pg_session_jwt" >}} | `0.3.3` | {{< badge content="Link" link="https://github.com/neondatabase/pg_session_jwt" >}} | {{< category "SEC" >}} | `pg_session_jwt_$v` | `postgresql-$v-pg-session-jwt` |
| {{< ext "anon" "pg_anon" >}} | `2.4.1` | {{< badge content="Link" link="https://gitlab.com/dalibo/postgresql_anonymizer/" >}} | {{< category "SEC" >}} | `pg_anon_$v` | `postgresql-$v-pg-anon` |
| {{< ext "pg_tde" >}} | `1.0` | {{< badge content="Link" link="https://github.com/Percona-Lab/pg_tde" >}} | {{< category "SEC" >}} | `percona-postgresql$v` | `percona-postgresql-$v` |
| {{< ext "pgsmcrypto" >}} | `0.1.1` | {{< badge content="Link" link="https://github.com/zhuobie/pgsmcrypto" >}} | {{< category "SEC" >}} | `pgsmcrypto_$v` | `postgresql-$v-pgsmcrypto` |
| {{< ext "pgaudit" >}} | `17.1` | {{< badge content="Link" link="https://github.com/pgaudit/pgaudit" >}} | {{< category "SEC" >}} | `pgaudit_$v*` | `postgresql-$v-pgaudit` |
| {{< ext "pgauditlogtofile" >}} | `1.7.1` | {{< badge content="Link" link="https://github.com/fmbiete/pgauditlogtofile" >}} | {{< category "SEC" >}} | `pgauditlogtofile_$v*` | `postgresql-$v-pgauditlogtofile` |
| {{< ext "pg_auth_mon" >}} | `3.0` | {{< badge content="Link" link="https://github.com/RafiaSabih/pg_auth_mon" >}} | {{< category "SEC" >}} | `pg_auth_mon_$v*` | `postgresql-$v-pg-auth-mon` |
| {{< ext "credcheck" >}} | `4.2` | {{< badge content="Link" link="https://github.com/MigOpsRepos/credcheck" >}} | {{< category "SEC" >}} | `credcheck_$v*` | `postgresql-$v-credcheck` |
| {{< ext "pgcryptokey" >}} | `0.85` | {{< badge content="Link" link="https://momjian.us/download/pgcryptokey/" >}} | {{< category "SEC" >}} | `pgcryptokey_$v` | `postgresql-$v-pgcryptokey` |
| {{< ext "pg_jobmon" >}} | `1.4.1` | {{< badge content="Link" link="https://github.com/omniti-labs/pg_jobmon" >}} | {{< category "SEC" >}} | `pg_jobmon_$v` | `postgresql-$v-pg-jobmon` |
| {{< ext "logerrors" >}} | `2.1.5` | {{< badge content="Link" link="https://github.com/munakoiso/logerrors" >}} | {{< category "SEC" >}} | `logerrors_$v*` | `postgresql-$v-logerrors` |
| {{< ext "login_hook" >}} | `1.7` | {{< badge content="Link" link="https://github.com/splendiddata/login_hook" >}} | {{< category "SEC" >}} | `login_hook_$v*` | `postgresql-$v-login-hook` |
| {{< ext "set_user" >}} | `4.1.0` | {{< badge content="Link" link="https://github.com/pgaudit/set_user" >}} | {{< category "SEC" >}} | `set_user_$v*` | `postgresql-$v-set-user` |
| {{< ext "pg_snakeoil" >}} | `1.4` | {{< badge content="Link" link="https://github.com/credativ/pg_snakeoil" >}} | {{< category "SEC" >}} | `pg_snakeoil_$v*` | `postgresql-$v-snakeoil` |
| {{< ext "pgextwlist" >}} | `1.19` | {{< badge content="Link" link="https://github.com/dimitri/pgextwlist" >}} | {{< category "SEC" >}} | `pgextwlist_$v*` | `postgresql-$v-pgextwlist` |
| {{< ext "pg_auditor" >}} | `0.2` | {{< badge content="Link" link="https://github.com/kouber/pg_auditor" >}} | {{< category "SEC" >}} | `pg_auditor_$v` | `postgresql-$v-pg-auditor` |
| {{< ext "sslutils" >}} | `1.4` | {{< badge content="Link" link="https://github.com/EnterpriseDB/sslutils" >}} | {{< category "SEC" >}} | `sslutils_$v*` | `postgresql-$v-sslutils` |
| {{< ext "noset" "pg_noset" >}} | `0.3.0` | {{< badge content="Link" link="https://gitlab.com/ongresinc/extensions/noset" >}} | {{< category "SEC" >}} | `noset_$v*` | `postgresql-$v-noset` |
| {{< ext "sepgsql" >}} | `-` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/sepgsql.html" >}} | {{< category "SEC" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "auth_delay" >}} | `-` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/auth-delay.html" >}} | {{< category "SEC" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "pgcrypto" >}} | `1.3` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/pgcrypto.html" >}} | {{< category "SEC" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "passwordcheck" >}} | `-` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/passwordcheck.html" >}} | {{< category "SEC" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "wrappers" >}} | `0.5.6` | {{< badge content="Link" link="https://github.com/supabase/wrappers" >}} | {{< category "FDW" >}} | `wrappers_$v` | `postgresql-$v-wrappers` |
| {{< ext "multicorn" >}} | `3.2` | {{< badge content="Link" link="https://github.com/pgsql-io/multicorn2" >}} | {{< category "FDW" >}} | `multicorn2_$v*` | - |
| {{< ext "odbc_fdw" >}} | `0.5.1` | {{< badge content="Link" link="https://github.com/CartoDB/odbc_fdw" >}} | {{< category "FDW" >}} | `odbc_fdw_$v*` | - |
| {{< ext "jdbc_fdw" >}} | `1.2` | {{< badge content="Link" link="https://github.com/pgspider/jdbc_fdw" >}} | {{< category "FDW" >}} | `jdbc_fdw_$v*` | - |
| {{< ext "pgspider_ext" >}} | `1.3.0` | {{< badge content="Link" link="https://github.com/pgspider/pgspider_ext" >}} | {{< category "FDW" >}} | `pgspider_ext_$v*` | `postgresql-$v-pgspider-ext` |
| {{< ext "mysql_fdw" >}} | `2.9.3` | {{< badge content="Link" link="https://github.com/EnterpriseDB/mysql_fdw" >}} | {{< category "FDW" >}} | `mysql_fdw_$v*` | `postgresql-$v-mysql-fdw` |
| {{< ext "oracle_fdw" >}} | `2.8.0` | {{< badge content="Link" link="https://github.com/laurenz/oracle_fdw" >}} | {{< category "FDW" >}} | `oracle_fdw_$v*` | `postgresql-$v-oracle-fdw` |
| {{< ext "tds_fdw" >}} | `2.0.5` | {{< badge content="Link" link="https://github.com/tds-fdw/tds_fdw" >}} | {{< category "FDW" >}} | `tds_fdw_$v*` | `postgresql-$v-tds-fdw` |
| {{< ext "db2_fdw" >}} | `7.0.0` | {{< badge content="Link" link="https://github.com/wolfgangbrandl/db2_fdw" >}} | {{< category "FDW" >}} | `db2_fdw_$v*` | - |
| {{< ext "sqlite_fdw" >}} | `2.5.0` | {{< badge content="Link" link="https://github.com/pgspider/sqlite_fdw" >}} | {{< category "FDW" >}} | `sqlite_fdw_$v*` | `postgresql-$v-sqlite-fdw` |
| {{< ext "pgbouncer_fdw" >}} | `1.4.0` | {{< badge content="Link" link="https://github.com/CrunchyData/pgbouncer_fdw" >}} | {{< category "FDW" >}} | `pgbouncer_fdw_$v` | - |
| {{< ext "mongo_fdw" >}} | `5.5.3` | {{< badge content="Link" link="https://github.com/EnterpriseDB/mongo_fdw" >}} | {{< category "FDW" >}} | `mongo_fdw_$v*` | - |
| {{< ext "redis_fdw" >}} | `1.0` | {{< badge content="Link" link="https://github.com/pg-redis-fdw/redis_fdw" >}} | {{< category "FDW" >}} | `redis_fdw_$v*` | `postgresql-$v-redis-fdw` |
| {{< ext "redis" "pg_redis_pubsub" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/brettlaforge/pg_redis_pubsub" >}} | {{< category "FDW" >}} | `pg_redis_pubsub_$v*` | `postgresql-$v-pg-redis-pubsub` |
| {{< ext "kafka_fdw" >}} | `0.0.3` | {{< badge content="Link" link="https://github.com/adjust/kafka_fdw" >}} | {{< category "FDW" >}} | `kafka_fdw_$v` | `postgresql-$v-kafka-fdw` |
| {{< ext "hdfs_fdw" >}} | `2.3.2` | {{< badge content="Link" link="https://github.com/EnterpriseDB/hdfs_fdw" >}} | {{< category "FDW" >}} | `hdfs_fdw_$v*` | - |
| {{< ext "firebird_fdw" >}} | `1.4.1` | {{< badge content="Link" link="https://github.com/ibarwick/firebird_fdw" >}} | {{< category "FDW" >}} | `firebird_fdw_$v` | `postgresql-$v-firebird-fdw` |
| {{< ext "aws_s3" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/chimpler/postgres-aws-s3" >}} | {{< category "FDW" >}} | `aws_s3_$v` | `postgresql-$v-aws-s3` |
| {{< ext "log_fdw" >}} | `1.4` | {{< badge content="Link" link="https://github.com/aws/postgresql-logfdw" >}} | {{< category "FDW" >}} | `log_fdw_$v*` | `postgresql-$v-log-fdw` |
| {{< ext "dblink" >}} | `1.2` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/dblink.html" >}} | {{< category "FDW" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "file_fdw" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/file-fdw.html" >}} | {{< category "FDW" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "postgres_fdw" >}} | `1.1` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/postgres-fdw.html" >}} | {{< category "FDW" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "documentdb" >}} | `0.106` | {{< badge content="Link" link="https://github.com/microsoft/documentdb" >}} | {{< category "SIM" >}} | `documentdb_$v*` | `postgresql-$v-documentdb` |
| {{< ext "orafce" >}} | `4.14.6` | {{< badge content="Link" link="https://github.com/orafce/orafce" >}} | {{< category "SIM" >}} | `orafce_$v` | `postgresql-$v-orafce` |
| {{< ext "pgtt" >}} | `4.4` | {{< badge content="Link" link="https://github.com/darold/pgtt" >}} | {{< category "SIM" >}} | `pgtt_$v*` | `postgresql-$v-pgtt` |
| {{< ext "session_variable" >}} | `3.4` | {{< badge content="Link" link="https://github.com/splendiddata/session_variable" >}} | {{< category "SIM" >}} | `session_variable_$v*` | `postgresql-$v-session-variable` |
| {{< ext "pg_statement_rollback" >}} | `1.4` | {{< badge content="Link" link="https://github.com/lzlabs/pg_statement_rollback" >}} | {{< category "SIM" >}} | `pg_statement_rollback_$v*` | `postgresql-$v-pg-statement-rollback` |
| {{< ext "pg_dbms_metadata" >}} | `1.0.0` | {{< badge content="Link" link="https://github.com/HexaCluster/pg_dbms_metadata" >}} | {{< category "SIM" >}} | `pg_dbms_metadata_$v` | - |
| {{< ext "pg_dbms_lock" >}} | `1.0` | {{< badge content="Link" link="https://github.com/HexaCluster/pg_dbms_lock" >}} | {{< category "SIM" >}} | `pg_dbms_lock_$v` | - |
| {{< ext "pg_dbms_job" >}} | `1.5` | {{< badge content="Link" link="https://github.com/MigOpsRepos/pg_dbms_job" >}} | {{< category "SIM" >}} | `pg_dbms_job_$v` | - |
| {{< ext "pg_dbms_errlog" >}} | `2.2` | {{< badge content="Link" link="https://github.com/HexaCluster/pg_dbms_errlog" >}} | {{< category "SIM" >}} | `pg_dbms_errlog_$v*` | - |
| {{< ext "babelfishpg_common" >}} | `3.3.3` | {{< badge content="Link" link="https://babelfishpg.org/" >}} | {{< category "SIM" >}} | `babelfishpg-common*` | `babelfishpg-common` |
| {{< ext "babelfishpg_tsql" >}} | `3.3.1` | {{< badge content="Link" link="https://babelfishpg.org/" >}} | {{< category "SIM" >}} | `babelfishpg-tsql*` | `babelfishpg-tsql` |
| {{< ext "babelfishpg_tds" >}} | `1.0.0` | {{< badge content="Link" link="https://babelfishpg.org/" >}} | {{< category "SIM" >}} | `babelfishpg-tds*` | `babelfishpg-tds` |
| {{< ext "babelfishpg_money" >}} | `1.1.0` | {{< badge content="Link" link="https://babelfishpg.org/" >}} | {{< category "SIM" >}} | `babelfishpg-money*` | `babelfishpg-money` |
| {{< ext "spat" >}} | `0.1.0a4` | {{< badge content="Link" link="https://github.com/Florents-Tselai/spat" >}} | {{< category "SIM" >}} | `spat_$v*` | `postgresql-$v-spat` |
| {{< ext "pgmemcache" >}} | `2.3.0` | {{< badge content="Link" link="https://github.com/ohmu/pgmemcache" >}} | {{< category "SIM" >}} | `pgmemcache_$v*` | `postgresql-$v-pgmemcache` |
| {{< ext "pglogical" >}} | `2.4.6` | {{< badge content="Link" link="https://github.com/2ndQuadrant/pglogical" >}} | {{< category "ETL" >}} | `pglogical_$v*` | `postgresql-$v-pglogical` |
| {{< ext "pglogical_ticker" >}} | `1.4.1` | {{< badge content="Link" link="https://github.com/enova/pglogical_ticker" >}} | {{< category "ETL" >}} | `pglogical_ticker_$v*` | `postgresql-$v-pglogical-ticker` |
| {{< ext "pgl_ddl_deploy" >}} | `2.2.1` | {{< badge content="Link" link="https://github.com/enova/pgl_ddl_deploy" >}} | {{< category "ETL" >}} | `pgl_ddl_deploy_$v*` | `postgresql-$v-pgl-ddl-deploy` |
| {{< ext "pg_failover_slots" >}} | `1.2.0` | {{< badge content="Link" link="https://github.com/EnterpriseDB/pg_failover_slots" >}} | {{< category "ETL" >}} | `pg_failover_slots_$v*` | `postgresql-$v-pg-failover-slots` |
| {{< ext "db_migrator" >}} | `1.0.0` | {{< badge content="Link" link="https://github.com/cybertec-postgresql/db_migrator" >}} | {{< category "ETL" >}} | `db_migrator_$v` | `postgresql-$v-db-migrator` |
| {{< ext "pgactive" >}} | `2.1.7` | {{< badge content="Link" link="https://github.com/aws/pgactive" >}} | {{< category "ETL" >}} | `pgactive_$v*` | `postgresql-$v-pgactive` |
| {{< ext "wal2json" >}} | `2.6` | {{< badge content="Link" link="https://github.com/eulerto/wal2json" >}} | {{< category "ETL" >}} | `wal2json_$v*` | `postgresql-$v-wal2json` |
| {{< ext "wal2mongo" >}} | `1.0.7` | {{< badge content="Link" link="https://github.com/HighgoSoftware/wal2mongo" >}} | {{< category "ETL" >}} | `wal2mongo_$v*` | `postgresql-$v-wal2mongo` |
| {{< ext "decoderbufs" >}} | `3.2.0` | {{< badge content="Link" link="https://github.com/debezium/postgres-decoderbufs" >}} | {{< category "ETL" >}} | `postgres-decoderbufs_$v*` | `postgresql-$v-decoderbufs` |
| {{< ext "decoder_raw" >}} | `1.0` | {{< badge content="Link" link="https://github.com/michaelpq/pg_plugins/blob/main/decoder_raw/" >}} | {{< category "ETL" >}} | `decoder_raw_$v*` | `postgresql-$v-decoder-raw` |
| {{< ext "mimeo" >}} | `1.5.1` | {{< badge content="Link" link="https://github.com/omniti-labs/mimeo" >}} | {{< category "ETL" >}} | `mimeo_$v` | `postgresql-$v-mimeo` |
| {{< ext "repmgr" >}} | `5.5.0` | {{< badge content="Link" link="https://github.com/EnterpriseDB/repmgr" >}} | {{< category "ETL" >}} | `repmgr_$v*` | `postgresql-$v-repmgr` |
| {{< ext "pg_fact_loader" >}} | `2.0.1` | {{< badge content="Link" link="https://github.com/enova/pg_fact_loader" >}} | {{< category "ETL" >}} | `pg_fact_loader_$v*` | `postgresql-$v-pg-fact-loader` |
| {{< ext "pg_bulkload" >}} | `3.1.22` | {{< badge content="Link" link="https://github.com/ossc-db/pg_bulkload" >}} | {{< category "ETL" >}} | `pg_bulkload_$v*` | `postgresql-$v-pg-bulkload` |
| {{< ext "test_decoding" >}} | `-` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/test-decoding.html" >}} | {{< category "ETL" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "pgoutput" >}} | `-` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/protocol-logical-replication.html" >}} | {{< category "ETL" >}} | `postgresql$v-contrib` | `postgresql-$v` |

## 扩展

共有 431 个可用的 PostgreSQL 扩展：

| 扩展 | PG 版本列表          | 属性 | 描述 |
|:-----|:-------------------|:----:|:-----|
| {{< ext "timescaledb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | 时序数据库扩展插件 |
| {{< ext "timescaledb_toolkit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | 超表分析查询，时间序列流式处理，以及其他SQL工具 |
| {{< ext "timeseries" "pg_timeseries" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | 时序数据API封装 |
| {{< ext "periods" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 为 PERIODs 和 SYSTEM VERSIONING 提供标准 SQL 功能 |
| {{< ext "temporal_tables" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 时态表功能支持 |
| {{< ext "emaj" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 让数据库的子集具有细粒度日志和时间旅行功能 |
| {{< ext "table_version" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostgreSQL 版本控制表扩展 |
| {{< ext "pg_cron" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | 定时任务调度器 |
| {{< ext "pg_task" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | 在特定时间点在后台执行SQL命令 |
| {{< ext "pg_later" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 执行查询，并在稍后异步获取查询结果 |
| {{< ext "pg_background" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 在后台运行 SQL 查询 |
| {{< ext "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostGIS 几何和地理空间扩展 |
| {{< ext "postgis_topology" "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostGIS 拓扑空间类型和函数 |
| {{< ext "postgis_raster" "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostGIS 光栅类型和函数 |
| {{< ext "postgis_sfcgal" "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | PostGIS SFCGAL 函数 |
| {{< ext "postgis_tiger_geocoder" "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | PostGIS tiger 地理编码器和反向地理编码器 |
| {{< ext "address_standardizer" "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 地址标准化函数。 |
| {{< ext "address_standardizer_data_us" "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 地址标准化函数：美国数据集示例 |
| {{< ext "pgrouting" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 提供寻路能力 |
| {{< ext "pointcloud" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 提供激光雷达点云数据类型支持 |
| {{< ext "pointcloud_postgis" "pointcloud" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | 将激光雷达点云与PostGIS几何类型相集成 |
| {{< ext "h3" "pg_h3" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | H3六边形层级索引支持 |
| {{< ext "h3_postgis" "pg_h3" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | H3与PostGIS集成的扩展插件 |
| {{< ext "q3c" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Q3C天空索引插件 |
| {{< ext "ogr_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | GIS 数据外部数据源包装器 |
| {{< ext "geoip" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | IP 地理位置扩展（围绕 MaxMind GeoLite 数据集的包装器） |
| {{< ext "pg_polyline" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | Google快速Polyline编码解码扩展 |
| {{< ext "pg_geohash" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 使用GeoHash处理空间坐标的函数包 |
| {{< ext "mobilitydb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | MobilityDB地理空间投影数据管理分析平台 |
| {{< ext "tzf" "pg_tzf" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | 快速根据GPS经纬度坐标查找时区 |
| {{< ext "earthdistance" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 计算地球表面上的大圆距离 |
| {{< ext "vector" "pgvector" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 向量数据类型和 ivfflat / hnsw 访问方法 |
| {{< ext "vchord" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | 使用Rust重写的高性能向量扩展 |
| {{< ext "vectorscale" "pgvectorscale" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 使用DiskANN算法对向量进行高效索引 |
| {{< ext "vectorize" "pg_vectorize" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | 在PostgreSQL中封装RAG向量检索服务 |
| {{< ext "pg_similarity" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 提供17种距离度量函数 |
| {{< ext "smlar" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 高效的相似度搜索函数 |
| {{< ext "pg_summarize" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 使用LLM对文本字段进行总结 |
| {{< ext "pg_tiktoken" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 在PostgreSQL中计算OpenAI使用的Token数 |
| {{< ext "pg4ml" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | PG4ML是一个机器学习框架 |
| {{< ext "pgml" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | PostgresML：用SQL运行机器学习算法并训练模型 |
| {{< ext "pg_search" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | ParadeDB BM25算法全文检索插件，ES全文检索 |
| {{< ext "pgroonga" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | 使用Groonga，面向所有语言的高速全文检索平台 |
| {{< ext "pgroonga_database" "pgroonga" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | PGGroonga 数据库管理模块 |
| {{< ext "pg_bigm" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 基于二字组的多语言全文检索扩展 |
| {{< ext "zhparser" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 中文分词，全文搜索解析器 |
| {{< ext "pg_bestmatch" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | 在数据库内生成BM25稀疏向量 |
| {{< ext "vchord_bm25" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | BM25排序算法 |
| {{< ext "pg_tokenizer" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | 用于全文检索的分词器 |
| {{< ext "hunspell_cs_cz" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Hunspell捷克语全文检索词典 |
| {{< ext "hunspell_de_de" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Hunspell德语全文检索词典 |
| {{< ext "hunspell_en_us" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Hunspell英语全文检索词典 |
| {{< ext "hunspell_fr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Hunspell法语全文检索词典 |
| {{< ext "hunspell_ne_np" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Hunspell尼泊尔语全文检索词典 |
| {{< ext "hunspell_nl_nl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Hunspell荷兰语全文检索词典 |
| {{< ext "hunspell_nn_no" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Hunspell挪威语全文检索词典 |
| {{< ext "hunspell_pt_pt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Hunspell葡萄牙语全文检索词典 |
| {{< ext "hunspell_ru_ru" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Hunspell俄语全文检索词典 |
| {{< ext "hunspell_ru_ru_aot" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Hunspell俄语全文检索词典（来自AOT.ru小组） |
| {{< ext "fuzzystrmatch" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | 确定字符串之间的相似性和距离 |
| {{< ext "pg_trgm" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 文本相似度测量函数与模糊检索 |
| {{< ext "citus" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | Citus 分布式数据库 |
| {{< ext "citus_columnar" "citus" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,r,r" >}} | {{< badge content="--s-d--" color="blue" >}} | Citus 列式存储引擎 |
| {{< ext "columnar" "hydra" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 开源列式存储扩展 |
| {{< ext "pg_analytics" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | 由 DuckDB 驱动的数据分析引擎 |
| {{< ext "pg_duckdb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | 在PostgreSQL中的嵌入式DuckDB扩展 |
| {{< ext "pg_mooncake" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="---Ld--" color="blue" >}} | PostgreSQL列式存储表 |
| {{< ext "duckdb_fdw" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | DuckDB 外部数据源包装器 |
| {{< ext "pg_parquet" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLdt-" color="blue" >}} | 在PostgreSQL与本地/S3中的Parquet文件复制数据 |
| {{< ext "pg_fkpart" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | 按外键实用程序进行表分区的扩展 |
| {{< ext "pg_partman" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 用于按时间或 ID 管理分区表的扩展 |
| {{< ext "plproxy" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 作为过程语言实现的数据库分区 |
| {{< ext "pg_strom" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 使用GPU与NVMe加速大数据处理 |
| {{< ext "tablefunc" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | 交叉表函数 |
| {{< ext "age" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Apache AGE，图数据库扩展 （Deb可用） |
| {{< ext "hll" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | hyperloglog 数据类型 |
| {{< ext "rum" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | RUM 索引访问方法 |
| {{< ext "pg_graphql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | PG内的GraphQL支持 |
| {{< ext "pg_jsonschema" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | 提供JSON Schema校验能力 |
| {{< ext "jsquery" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 用于内省 JSONB 数据类型的查询类型 |
| {{< ext "pg_hint_plan" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 添加强制指定执行计划的能力 |
| {{< ext "hypopg" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 假设索引，用于创建一个虚拟索引检验执行计划 |
| {{< ext "index_advisor" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | 查询索引建议器 |
| {{< ext "plan_filter" "pg_plan_filter" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | 使用执行计划代价过滤阻止特定查询语句 |
| {{< ext "imgsmlr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 使用Haar小波分析计算图片相似度 |
| {{< ext "pg_ivm" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 增量维护的物化视图 |
| {{< ext "pg_incremental" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,r,r,r" >}} | {{< badge content="--s-d--" color="blue" >}} | 增量处理流式事件 |
| {{< ext "pgmq" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | 基于Postgres实现类似AWS SQS/RSMQ的消息队列 |
| {{< ext "pgq" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 通用队列的PG实现 |
| {{< ext "orioledb" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | OrioleDB，下一代事务处理引擎 |
| {{< ext "pg_cardano" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Cardano相关工具包：加密函数，地址编解码，区块链处理 |
| {{< ext "rdkit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 在PostgreSQL化学领域数据管理功能 |
| {{< ext "omni" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | PostgreSQL即平台，Omnigres主扩展与加载器 |
| {{< ext "omni_auth" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Omnigres 基础会话认证管理模块 |
| {{< ext "omni_aws" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | Omnigres AWS S3 API封装 |
| {{< ext "omni_cloudevents" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | Omnigres CloudEvents 支持 |
| {{< ext "omni_containers" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres Docker容器管理模块 |
| {{< ext "omni_credentials" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Omnigres 应用密钥管理模块 |
| {{< ext "omni_csv" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres CSV 工具箱 |
| {{< ext "omni_datasets" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Omnigres 数据库置备工具 |
| {{< ext "omni_email" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Omnigres Email 框架 |
| {{< ext "omni_http" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Omnigres 基本HTTP类型 |
| {{< ext "omni_httpc" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres HTTP客户端 |
| {{< ext "omni_httpd" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres HTTP服务器 |
| {{< ext "omni_id" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Omnigres ID身份数据类型 |
| {{< ext "omni_json" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | Omnigres JSON工具箱 |
| {{< ext "omni_kube" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres Kubernetes集成模块 |
| {{< ext "omni_ledger" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres 金融账本模块 |
| {{< ext "omni_manifest" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Omnigres 包管理清单模块 |
| {{< ext "omni_mimetypes" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Omnigres MIME数据类型 |
| {{< ext "omni_os" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres 操作系统集成模块 |
| {{< ext "omni_polyfill" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres Postgres多态API |
| {{< ext "omni_python" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres 第一类Python支持模块 |
| {{< ext "omni_regex" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Omnigres PCRE兼容正则表达式模块 |
| {{< ext "omni_rest" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Omnigres REST API 工具包 |
| {{< ext "omni_schema" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Omnigres 高级模式管理组件 |
| {{< ext "omni_seq" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres 分布式整型序列号 |
| {{< ext "omni_service" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Omnigres 服务管理器 |
| {{< ext "omni_session" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Omnigres 会话管理器 |
| {{< ext "omni_shmem" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres 共享内存管理 |
| {{< ext "omni_sql" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres SQL编程组件 |
| {{< ext "omni_sqlite" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres 嵌入的SQLite支持 |
| {{< ext "omni_test" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Omnigres 测试框架 |
| {{< ext "omni_txn" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres 事务管理器模块 |
| {{< ext "omni_types" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres 高级数据类型模块 |
| {{< ext "omni_var" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres 局部变量模块 |
| {{< ext "omni_vfs" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres 虚拟文件系统 |
| {{< ext "omni_vfs_types_v1" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Omnigres 虚拟文件系统（v1） |
| {{< ext "omni_web" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres Web工具箱 |
| {{< ext "omni_worker" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres 通用Worker池 |
| {{< ext "omni_xml" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres XML工具包 |
| {{< ext "omni_yaml" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Omnigres YAML工具包 |
| {{< ext "omnisketch" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 实现OmniSketch数据结构，实现近似摘要聚合 |
| {{< ext "bloom" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | bloom 索引-基于指纹的索引 |
| {{< ext "pg_tle" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | AWS 可信语言扩展 |
| {{< ext "plv8" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PL/JavaScript (v8) 可信过程程序语言 |
| {{< ext "pljs" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | PL/JS 可信过程程序语言 |
| {{< ext "pllua" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Lua 程序语言 |
| {{< ext "hstore_pllua" "pllua" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Lua 程序语言的Hstore适配扩展 |
| {{< ext "plluau" "pllua" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Lua 程序语言（不受信任的） |
| {{< ext "hstore_plluau" "pllua" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Lua 程序语言的Hstore适配扩展（不受信任的） |
| {{< ext "plprql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 在PostgreSQL使用PRQL——管线式关系查询语言 |
| {{< ext "pldbgapi" "pldebugger" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 用于调试 PL/pgSQL 函数的服务器端支持 |
| {{< ext "plpgsql_check" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | 对 plpgsql 函数进行扩展检查 |
| {{< ext "plprofiler" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 剖析 PL/pgSQL 函数 |
| {{< ext "plsh" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | PL/sh 程序语言 |
| {{< ext "pljava" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Java 程序语言 |
| {{< ext "plr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 从数据库中加载R语言解释器并执行R脚本 |
| {{< ext "plxslt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | XSLT 存储过程语言 |
| {{< ext "pgtap" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostgreSQL单元测试框架 |
| {{< ext "faker" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 插入生成的测试伪造数据，Python库的包装 |
| {{< ext "dbt2" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | OSDL-DBT-2 测试组件 |
| {{< ext "pltcl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | PL/TCL 存储过程语言 |
| {{< ext "pltclu" "pltcl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | PL/TCL 存储过程语言（未受信/高权限） |
| {{< ext "plperl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | PL/Perl 存储过程语言 |
| {{< ext "bool_plperl" "plperl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 在 bool 和 plperl 之间转换 |
| {{< ext "hstore_plperl" "plperl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 在 hstore 和 plperl 之间转换适配类型 |
| {{< ext "jsonb_plperl" "plperl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | 在 jsonb 和 plperl 之间转换 |
| {{< ext "plperlu" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | PL/PerlU 存储过程语言（未受信/高权限） |
| {{< ext "bool_plperlu" "plperlu" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | 在 bool 和 plperlu 之间转换 |
| {{< ext "jsonb_plperlu" "plperlu" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | 在 jsonb 和 plperlu 之间转换 |
| {{< ext "hstore_plperlu" "plperlu" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | 在 hstore 和 plperlu 之间转换适配类型 |
| {{< ext "plpgsql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | PL/pgSQL 程序设计语言 |
| {{< ext "plpython3u" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | PL/Python3 存储过程语言（未受信/高权限） |
| {{< ext "jsonb_plpython3u" "plpython3u" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d-r" color="blue" >}} | 在 jsonb 和 plpython3u 之间转换 |
| {{< ext "ltree_plpython3u" "plpython3u" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d-r" color="blue" >}} | 在 ltree 和 plpython3u 之间转换 |
| {{< ext "hstore_plpython3u" "plpython3u" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d-r" color="blue" >}} | 在 hstore 和 plpython3u 之间转换 |
| {{< ext "prefix" "pg_prefix" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 前缀树数据类型 |
| {{< ext "semver" "pg_semver" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 语义版本号数据类型 |
| {{< ext "unit" "pgunit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | SI 国标单位扩展 |
| {{< ext "pgpdf" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLdtr" color="blue" >}} | PDF数据类型，管理函数与全文检索 |
| {{< ext "pglite_fusion" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | 在PG表中嵌入SQLite数据库作为数据类型 |
| {{< ext "md5hash" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 提供128位MD5的原生数据类型 |
| {{< ext "asn1oid" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | ASN1OID数据类型支持 |
| {{< ext "roaringbitmap" "pg_roaringbitmap" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 支持RoaringBitmap数据类型 |
| {{< ext "pgfaceting" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | 使用倒排索引的高速切面查询 |
| {{< ext "pg_sphere" "pgsphere" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 球面对象函数、运算符与索引支持 |
| {{< ext "country" "pg_country" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 国家代码数据类型，遵循ISO 3166-1标准 |
| {{< ext "pg_xenophile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | PostgreSQL i8n与l10n工具包 |
| {{< ext "l10n_table_dependent_extension" "pg_xenophile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | PostgreSQL l10n 工具包 |
| {{< ext "currency" "pg_currency" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 使用1字节表示的货币数据类型 |
| {{< ext "collection" "pgcollection" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | 在PlPGSQL中使用的内存优化高性能集合数据结构 |
| {{< ext "pgmp" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 多精度算术扩展 |
| {{< ext "numeral" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 数值类型扩展 |
| {{< ext "pg_rational" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 使用BIGINT表示的有理数数据类型 |
| {{< ext "uint" "pguint" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 无符号整型数据类型 |
| {{< ext "uint128" "pg_uint128" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 原生128位无符号整型数据类型 |
| {{< ext "hashtypes" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 包括SHA1，MD5在内的多种哈希数据类型 |
| {{< ext "ip4r" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | PostgreSQL 的 IPv4/v6 和 IPv4/v6 范围索引类型 |
| {{< ext "pg_duration" >}} | {{< pgver "18,17,16,15,14,13" "g,g,r,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | 用于表示时间段的强化数据类型 |
| {{< ext "uri" "pg_uri" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | URI数据类型 |
| {{< ext "emailaddr" "pg_emailaddr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Email地址数据类型 |
| {{< ext "acl" "pg_acl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | ACL数据类型 |
| {{< ext "debversion" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Debian版本号数据类型 |
| {{< ext "pg_rrule" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 日历重复规则RRULE数据类型 |
| {{< ext "timestamp9" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 纳秒分辨率时间戳 |
| {{< ext "chkpass" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 数据类型：自动加密的密码 |
| {{< ext "isn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | 用于国际产品编号标准的数据类型 |
| {{< ext "seg" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | 表示线段或浮点间隔的数据类型 |
| {{< ext "cube" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 用于存储多维立方体的数据类型 |
| {{< ext "ltree" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | 用于表示分层树状结构的数据类型 |
| {{< ext "hstore" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | 用于存储（键，值）对集合的数据类型 |
| {{< ext "citext" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | 提供大小写不敏感的字符串类型 |
| {{< ext "xml2" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | XPath 查询和 XSLT |
| {{< ext "gzip" "pg_gzip" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 使用SQL执行Gzip压缩与解压缩 |
| {{< ext "bzip" "pg_bzip" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | BZIP压缩解压缩函数包 |
| {{< ext "zstd" "pg_zstd" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | ZSTD压缩解压缩函数包 |
| {{< ext "http" "pg_http" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | HTTP客户端，允许在数据库内收发HTTP请求 (supabase) |
| {{< ext "pg_net" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | 用 SQL 进行异步非阻塞HTTP/HTTPS 请求的扩展 (supabase) |
| {{< ext "pg_curl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 封装CURL，执行各种用URL传输数据的操作 |
| {{< ext "pgjq" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dtr" color="blue" >}} | 在Postgres中使用jq查询JSON |
| {{< ext "pgjwt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | JSON Web Token API 的PG实现 (supabase) |
| {{< ext "pg_smtp_client" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | 使用SMTP从PostgreSQL内发送邮件的客户端扩展 |
| {{< ext "pg_html5_email_address" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | 验证Email是否符合HTML5规范的扩展 |
| {{< ext "url_encode" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 提供URL编码解码函数 |
| {{< ext "pgsql_tweaks" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 一些日常会用到的便利函数与视图 |
| {{< ext "pg_extra_time" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | 一些关于日期与时间的扩展函数 |
| {{< ext "pgpcre" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | PCRE/Perl风格的正则表达式支持 |
| {{< ext "icu_ext" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 访问ICU库提供的函数 |
| {{< ext "pgqr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 从数据库中直接生成QR二维码 |
| {{< ext "pg_protobuf" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 提供Protobuf函数支持 |
| {{< ext "envvar" "pg_envvar" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 获取环境变量的函数 |
| {{< ext "floatfile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 将浮点数组存储到文件中而不是堆表中 |
| {{< ext "pg_render" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | 使用SQL渲染HTML页面 |
| {{< ext "pg_readme" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | 为模式与扩展生成Markdown文档 |
| {{< ext "pg_readme_test_extension" "pg_readme" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | 为模式与扩展生成Markdown文档 |
| {{< ext "ddl_historization" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | 用SQL将所有DDL变更写入到数据库表中 |
| {{< ext "data_historization" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | 用SQL将数据变更历史保存到分区表中 |
| {{< ext "schedoc" "pg_schedoc" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | 在Django与DBT之间通过注释文档交换元数据 |
| {{< ext "hashlib" "pg_hashlib" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | 稳定哈希函数包 |
| {{< ext "xxhash" "pg_xxhash" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | xxhash哈希函数包 |
| {{< ext "shacrypt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 实现SHA256-CRYPT与SHA512-CRYPT密钥加密算法 |
| {{< ext "cryptint" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 加密INT与BIGINT类型 |
| {{< ext "pguecc" "pg_ecdsa" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | PostgreSQL的uECC绑定，椭圆曲线加解密函数包 |
| {{< ext "sparql" "pgsparql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | 使用SQL查询SPARQL数据源 |
| {{< ext "pg_idkit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 生成各式各样的唯一标识符：UUIDv6, ULID, KSUID |
| {{< ext "pgx_ulid" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | ULID数据类型与函数 |
| {{< ext "pg_uuidv7" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | UUIDv7 支持 |
| {{< ext "permuteseq" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 伪随机数ID置换生成器 |
| {{< ext "pg_hashids" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 加盐将整型ID转为短字符串ID |
| {{< ext "sequential_uuids" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 生成连续生成的UUID |
| {{< ext "topn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | top-n JSONB 的类型 |
| {{< ext "quantile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Quantile聚合函数 |
| {{< ext "lower_quantile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Lower Quantile 聚合函数 |
| {{< ext "count_distinct" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | COUNT(DISTINCT …) 聚合的替代方案 |
| {{< ext "ddsketch" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 实现DDSketch数据结构，实现在线的Quantile聚合 |
| {{< ext "vasco" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 使用MIC发现数据中隐含的关联 |
| {{< ext "xicor" "pgxicor" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | 在PG中计算XI相关系数 |
| {{< ext "tdigest" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | tdigest 聚合函数 |
| {{< ext "first_last_agg" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | first() 与 last() 聚合函数 |
| {{< ext "extra_window_functions" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 额外的窗口函数 |
| {{< ext "floatvec" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 数组类型数学运算扩展 |
| {{< ext "aggs_for_vecs" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 针对数组类型的聚合函数集合扩展 |
| {{< ext "aggs_for_arrays" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 计算数组聚合统计值的函数包 |
| {{< ext "pg_csv" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dtr" color="blue" >}} | 灵活的CSV聚合处理函数 |
| {{< ext "arraymath" "pg_arraymath" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 数组逐元素数学运算符包 |
| {{< ext "pg_math" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 使用GSL库的数学统计函数 |
| {{< ext "random" "pg_random" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 随机数生成器 |
| {{< ext "base36" "pg_base36" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Base36编码解码扩展 |
| {{< ext "base62" "pg_base62" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Base62编码解码扩展 |
| {{< ext "pg_base58" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Base58 编码/解码函数 |
| {{< ext "financial" "pg_financial" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 金融领域聚合函数 |
| {{< ext "convert" "pg_convert" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | 用于空间里程等的公英制转换函数 |
| {{< ext "refint" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 实现引用完整性的函数 |
| {{< ext "autoinc" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 用于自动递增字段的函数 |
| {{< ext "insert_username" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 用于跟踪谁更改了表的函数 |
| {{< ext "moddatetime" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 跟踪最后修改时间 |
| {{< ext "tsm_system_time" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | 接受毫秒数限制的 TABLESAMPLE 方法 |
| {{< ext "dict_xsyn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 用于扩展同义词处理的文本搜索字典模板 |
| {{< ext "tsm_system_rows" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | 接受行数限制的 TABLESAMPLE 方法 |
| {{< ext "tcn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | 用触发器通知变更 |
| {{< ext "uuid-ossp" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | 生成通用唯一标识符（UUIDs） |
| {{< ext "btree_gist" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | 用GiST索引常见数据类型 |
| {{< ext "btree_gin" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | 用GIN索引常见数据类型 |
| {{< ext "intarray" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---dt-" color="blue" >}} | 1维整数数组的额外函数、运算符和索引支持 |
| {{< ext "intagg" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | 整数聚合器和枚举器（过时） |
| {{< ext "dict_int" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | 用于整数的文本搜索字典模板 |
| {{< ext "unaccent" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | 删除重音的文本搜索字典 |
| {{< ext "pg_repack" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="-bs-d--" color="blue" >}} | 在线垃圾清理与表膨胀治理 |
| {{< ext "pg_rewrite" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | 在线重写整表，不阻塞读写 |
| {{< ext "pg_squeeze" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | 从关系中删除未使用空间 |
| {{< ext "pg_dirtyread" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 从表中读取尚未垃圾回收的行 |
| {{< ext "pgfincore" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 检查和管理操作系统缓冲区缓存 |
| {{< ext "pg_cooldown" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | 从缓冲区中移除特定关系的页面 |
| {{< ext "ddlx" "pg_ddlx" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 提取数据库对象的DDL |
| {{< ext "prioritize" "pg_prioritize" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 获取和设置 PostgreSQL 后端的优先级 |
| {{< ext "pg_checksums" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s---r" color="blue" >}} | 在离线模式下激活/启用/禁用数据库集群的校验和功能 |
| {{< ext "pg_readonly" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 将集群设置为只读 |
| {{< ext "pgdd" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | 提供通过标准SQL查询数据库目录集簇的能力 |
| {{< ext "pg_permissions" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | 查看对象权限并将其与期望状态进行比较 |
| {{< ext "pgautofailover" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | PG 自动故障迁移 |
| {{< ext "pg_catcheck" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 用于诊断系统目录是否损坏的工具 |
| {{< ext "pre_prepare" "preprepare" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 在服务端预先准备好PreparedStatement备用 |
| {{< ext "pg_upless" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | 检测表上的无用UPDATE |
| {{< ext "pgcozy" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | 根据先前的pg_buffercache快照预热内存缓冲区 |
| {{< ext "pg_orphaned" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 处理孤儿文件的扩展插件 |
| {{< ext "pg_crash" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | 向数据库进程随机发送信号模拟故障 |
| {{< ext "pg_cheat_funcs" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 一些超级实用的作弊函数 |
| {{< ext "fio" "pg_fio" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | PostgreSQL文件IO函数包 |
| {{< ext "pg_savior" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 阻止不带条件的全表更新以避免意外事故 |
| {{< ext "safeupdate" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sL---" color="blue" >}} | 强制在 UPDATE 和 DELETE 时提供 Where 条件 |
| {{< ext "pg_drop_events" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | 记录删表删列删视图的事务号，辅助PITR确定时间点 |
| {{< ext "table_log" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 记录某张表的修改日志并做表/行级时间点恢复 |
| {{< ext "pgagent" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostgreSQL任务调度工具，与PGADMIN配合使用 |
| {{< ext "pg_prewarm" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sLd--" color="blue" >}} | 预热关系数据 |
| {{< ext "pgpool_adm" "pgpool" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | PGPool 管理函数 |
| {{< ext "pgpool_recovery" "pgpool" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | PGPool辅助扩展，从v4.3提供的恢复函数 |
| {{< ext "pgpool_regclass" "pgpool" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | PGPool辅助扩展，RegClass替代 |
| {{< ext "lo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | 大对象维护 |
| {{< ext "basic_archive" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="c-s----" color="blue" >}} | 归档模块样例 |
| {{< ext "basebackup_to_shell" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="c-s----" color="blue" >}} | 添加一种备份到Shell终端到基础备份方式 |
| {{< ext "old_snapshot" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,r" >}} | {{< badge content="c-s-d--" color="blue" >}} | 支持 old_snapshot_threshold 的实用程序 |
| {{< ext "adminpack" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | PostgreSQL 管理函数集合 |
| {{< ext "amcheck" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 校验关系完整性 |
| {{< ext "pg_surgery" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="c-s-d--" color="blue" >}} | 对损坏的关系进行手术 |
| {{< ext "pg_profile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostgreSQL 数据库负载记录与AWR报表工具 |
| {{< ext "pg_tracing" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | PostgreSQL分布式Tracing |
| {{< ext "pg_show_plans" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 打印所有当前正在运行查询的执行计划 |
| {{< ext "pg_stat_kcache" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | 内核统计信息收集 |
| {{< ext "pg_stat_monitor" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | 提供查询聚合统计、客户端信息、执行计划详细信息和直方图 |
| {{< ext "pg_qualstats" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | 收集有关 quals 的统计信息的扩展 |
| {{< ext "pg_store_plans" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | 跟踪所有执行的 SQL 语句的计划统计信息 |
| {{< ext "pg_track_settings" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | 跟踪设置更改 |
| {{< ext "pg_wait_sampling" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | 基于采样的等待事件统计 |
| {{< ext "pgsentinel" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | 活跃会话历史 |
| {{< ext "system_stats" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | PostgreSQL 的系统统计函数 |
| {{< ext "meta" "pg_meta" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | 标准化，更友好的PostgreSQL系统目录视图 |
| {{< ext "pgnodemx" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 使用SQL查询获取操作系统指标 |
| {{< ext "pg_proctab" "pgnodemx" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 通过SQL接口访问操作系统进程表 |
| {{< ext "pg_sqlog" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 提供访问PostgreSQL日志的SQL接口 |
| {{< ext "bgw_replstatus" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | 用于汇报本机主从状态的后台工作进程 |
| {{< ext "pgmeminfo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 显示内存使用情况 |
| {{< ext "toastinfo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 显示TOAST字段的详细信息 |
| {{< ext "explain_ui" "pg_explain_ui" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | 快速跳转至PEV查阅可视化执行计划 |
| {{< ext "pg_relusage" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | 打印查询引用的表与列 |
| {{< ext "pagevis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | 使用ASCII字符可视化数据库物理页面布局 |
| {{< ext "powa" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostgreSQL 工作负载分析器-核心 |
| {{< ext "pg_overexplain" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="c-sL---" color="blue" >}} | 允许 EXPLAIN 转储更多详细 |
| {{< ext "pg_logicalinspect" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="c-s-d--" color="blue" >}} | 检视逻辑解码组件详情 |
| {{< ext "pageinspect" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 检查数据库页面二进制内容 |
| {{< ext "pgrowlocks" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 显示行级锁信息 |
| {{< ext "sslinfo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 关于 SSL 证书的信息 |
| {{< ext "pg_buffercache" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 检查共享缓冲区缓存 |
| {{< ext "pg_walinspect" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="c-s-d--" color="blue" >}} | 用于检查 PostgreSQL WAL 日志内容的函数 |
| {{< ext "pg_freespacemap" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 检查自由空间映射的内容（FSM） |
| {{< ext "pg_visibility" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 检查可见性图（VM）和页面级可见性信息 |
| {{< ext "pgstattuple" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 显示元组级统计信息 |
| {{< ext "auto_explain" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | 提供一种自动记录执行计划的手段 |
| {{< ext "pg_stat_statements" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sLd--" color="blue" >}} | 跟踪所有执行的 SQL 语句的计划和执行统计信息 |
| {{< ext "passwordcheck_cracklib" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | 使用cracklib加固PG用户密码 |
| {{< ext "supautils" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | 用于在云环境中确保数据库集群的安全 |
| {{< ext "pgsodium" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | 表数据加密存储 TDE |
| {{< ext "supabase_vault" "pg_vault" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 在 Vault 中存储加密凭证的扩展 (supabase) |
| {{< ext "pg_session_jwt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | 使用JWT进行会话认证 |
| {{< ext "anon" "pg_anon" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | 数据匿名化处理工具 |
| {{< ext "pg_tde" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | Percona加密存储引擎 |
| {{< ext "pgsmcrypto" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 为PostgreSQL提供商密算法支持：SM2,SM3,SM4 |
| {{< ext "pgaudit" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | 提供审计功能 |
| {{< ext "pgauditlogtofile" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | pgAudit 子扩展，将审计日志写入单独的文件中 |
| {{< ext "pg_auth_mon" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 监控每个用户的连接尝试 |
| {{< ext "credcheck" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | 明文凭证检查器 |
| {{< ext "pgcryptokey" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | PG密钥管理 |
| {{< ext "pg_jobmon" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 记录和监控函数 |
| {{< ext "logerrors" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 用于收集日志文件中消息统计信息的函数 |
| {{< ext "login_hook" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 在用户登陆时执行login_hook.login()函数 |
| {{< ext "set_user" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 增加了日志记录的 SET ROLE |
| {{< ext "pg_snakeoil" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL--r" color="blue" >}} | PostgreSQL动态链接库反病毒功能 |
| {{< ext "pgextwlist" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | PostgreSQL扩展白名单功能 |
| {{< ext "pg_auditor" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | 审计数据变更并提供闪回能力 |
| {{< ext "sslutils" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 使用SQL管理SSL证书 |
| {{< ext "noset" "pg_noset" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | 阻止非超级用户使用SET/RESET设置变量 |
| {{< ext "sepgsql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | 基于SELinux标签的强制访问控制 |
| {{< ext "auth_delay" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | 在返回认证失败前暂停一会，避免爆破 |
| {{< ext "pgcrypto" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | 实用加解密函数 |
| {{< ext "passwordcheck" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | 用于强制拒绝修改弱密码的扩展 |
| {{< ext "wrappers" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | Supabase提供的外部数据源包装器捆绑包 |
| {{< ext "multicorn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 用Python编写自定义的外部数据源包装器 |
| {{< ext "odbc_fdw" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 访问ODBC可访问的任何外部数据源 |
| {{< ext "jdbc_fdw" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 访问JDBC可访问的任何外部数据源 |
| {{< ext "pgspider_ext" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | 使用多种FDW访问远程数据库服务器 |
| {{< ext "mysql_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | MySQL外部数据包装器 |
| {{< ext "oracle_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 提供对Oracle的外部数据源包装器 |
| {{< ext "tds_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | TDS 数据库（Sybase/SQL Server）外部数据包装器 |
| {{< ext "db2_fdw" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 提供对DB2的外部数据源包装器 |
| {{< ext "sqlite_fdw" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | SQLite 外部数据包装器 |
| {{< ext "pgbouncer_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 用SQL查询pgbouncer统计信息，并执行pgbouncer命令 |
| {{< ext "mongo_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | MongoDB 外部数据包装器 |
| {{< ext "redis_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 查询外部Redis数据源 |
| {{< ext "redis" "pg_redis_pubsub" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 从PG向Redis发送Pub/Sub消息 |
| {{< ext "kafka_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Kafka外部数据源包装器 |
| {{< ext "hdfs_fdw" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | hdfs 外部数据包装器 |
| {{< ext "firebird_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Firebird外部数据源包装器 |
| {{< ext "aws_s3" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | 从S3导入导出数据的外部数据源包装器 |
| {{< ext "log_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | 访问PostgreSQL日志文件的FDW |
| {{< ext "dblink" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 从数据库内连接到其他 PostgreSQL 数据库 |
| {{< ext "file_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 访问外部文件的外部数据包装器 |
| {{< ext "postgres_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | 用于远程 PostgreSQL 服务器的外部数据包装器 |
| {{< ext "documentdb" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | 微软DocumentDB的API层 |
| {{< ext "documentdb_core" "documentdb" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | 微软DocumentDB的核心API层实现 |
| {{< ext "documentdb_distributed" "documentdb" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | DocumentDB多节点模式的API层 |
| {{< ext "orafce" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 模拟 Oracle RDBMS 的一部分函数和包的函数和运算符 |
| {{< ext "pgtt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 类似Oracle的全局临时表功能 |
| {{< ext "session_variable" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Oracle兼容的会话变量/常量操作函数 |
| {{< ext "pg_statement_rollback" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | 在服务端提供类似Oracle/DB2的语句级回滚能力 |
| {{< ext "pg_dbms_metadata" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 添加 Oracle DBMS_METADATA 兼容性支持的扩展 |
| {{< ext "pg_dbms_lock" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 为PG添加对 Oracle DBMS_LOCK 的完整兼容性支持 |
| {{< ext "pg_dbms_job" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 添加 Oracle DBMS_JOB 兼容性支持的扩展 |
| {{< ext "pg_dbms_errlog" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 模仿 Oracle DBMS_ERRLOG 模块来记录特定表的DML错误 |
| {{< ext "babelfishpg_common" >}} | {{< pgver "18,17,16,15,14,13" "r,r,r,g,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | SQL Server 数据类型兼容扩展 |
| {{< ext "babelfishpg_tsql" >}} | {{< pgver "18,17,16,15,14,13" "r,r,r,g,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | SQL Server SQL语法兼容性扩展 |
| {{< ext "babelfishpg_tds" >}} | {{< pgver "18,17,16,15,14,13" "r,r,r,g,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | SQL Server TDS线缆协议兼容扩展 |
| {{< ext "babelfishpg_money" >}} | {{< pgver "18,17,16,15,14,13" "r,r,r,g,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | SQL Server 货币数据类型兼容扩展 |
| {{< ext "spat" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | 在PG中嵌入Redis风格的内存数据库 |
| {{< ext "pgmemcache" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | 为PG提供memcached兼容接口 |
| {{< ext "pglogical" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | PostgreSQL逻辑复制：三方扩展实现 |
| {{< ext "pglogical_origin" "pglogical" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 用于从 Postgres 9.4 升级时的兼容性虚拟扩展 |
| {{< ext "pglogical_ticker" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | pglogical复制延迟以秒计的精确视图 |
| {{< ext "pgl_ddl_deploy" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 使用 pglogical 执行自动 DDL 部署 |
| {{< ext "pg_failover_slots" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL--r" color="blue" >}} | 在Failover过程中保留复制槽 |
| {{< ext "db_migrator" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | 使用FDW从其他DBMS迁移到PostgreSQL |
| {{< ext "pgactive" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="-bsLd--" color="blue" >}} | PostgreSQL多主逻辑复制 |
| {{< ext "wal2json" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | 用逻辑解码捕获 JSON 格式的 CDC 变更 |
| {{< ext "wal2mongo" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | 使用逻辑解码捕获MongoDB JSON格式的CDC变更 |
| {{< ext "decoderbufs" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | 将WAL逻辑解码为ProtocolBuffer协议的消息 |
| {{< ext "decoder_raw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | 逻辑复制解码输出插件：RAW SQL格式 |
| {{< ext "mimeo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | 在PostgreSQL实例间进行表级复制 |
| {{< ext "repmgr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostgreSQL复制管理组件 |
| {{< ext "pg_fact_loader" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | 在 Postgres 中构建事实表 |
| {{< ext "pg_bulkload" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="-bs-d--" color="blue" >}} | 向 PostgreSQL 中高速加载数据 |
| {{< ext "test_decoding" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s----" color="blue" >}} | 基于SQL的WAL逻辑解码样例 |
| {{< ext "pgoutput" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s----" color="blue" >}} | PG内置的逻辑解码输出插件 |

**属性说明**: `c`:contrib `b`:二进制 `s`:动态库 `l`:需加载 `d`:需DDL `t`:无需特权 `r`:可重定位