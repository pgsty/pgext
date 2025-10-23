---
title: "Extensions"
breadcrumbs: false
excludeSearch: true
comments: false
weight: 900
---

PostgreSQL extensions enhance the database with additional functionality.


## Packages

There are 357 available PostgreSQL packages:

| Package | Version | Repo | Category | RPM | DEB |
|:--------|:--------|:-----|:---------|:-----|:-----|
| {{< ext "timescaledb" >}} | `2.22.0` | {{< badge content="Link" link="https://github.com/timescale/timescaledb" >}} | {{< category "TIME" >}} | `timescaledb-tsl_$v*` | `postgresql-$v-timescaledb-tsl` |
| {{< ext "timescaledb_toolkit" >}} | `1.21.0` | {{< badge content="Link" link="https://github.com/timescale/timescaledb-toolkit" >}} | {{< category "TIME" >}} | `timescaledb-toolkit_$v` | `postgresql-$v-timescaledb-toolkit` |
| {{< ext "timeseries" "pg_timeseries" >}} | `0.1.6` | {{< badge content="Link" link="https://github.com/ChuckHend/pg_timeseries" >}} | {{< category "TIME" >}} | `pg_timeseries_$v` | `postgresql-$v-pg-timeseries` |
| {{< ext "periods" >}} | `1.2.3` | {{< badge content="Link" link="https://github.com/xocolatl/periods" >}} | {{< category "TIME" >}} | `periods_$v*` | `postgresql-$v-periods` |
| {{< ext "temporal_tables" >}} | `1.2.2` | {{< badge content="Link" link="https://pgxn.org/dist/temporal_tables/" >}} | {{< category "TIME" >}} | `temporal_tables_$v*` | `postgresql-$v-temporal-tables` |
| {{< ext "emaj" >}} | `4.7.0` | {{< badge content="Link" link="https://github.com/dalibo/emaj" >}} | {{< category "TIME" >}} | `e-maj_$v` | `postgresql-$v-emaj` |
| {{< ext "table_version" >}} | `1.11.1` | {{< badge content="Link" link="https://github.com/linz/postgresql-tableversion" >}} | {{< category "TIME" >}} | `table_version_$v` | `postgresql-$v-table-version` |
| {{< ext "pg_cron" >}} | `1.6.7` | {{< badge content="Link" link="https://github.com/citusdata/pg_cron" >}} | {{< category "TIME" >}} | `pg_cron_$v*` | `postgresql-$v-cron` |
| {{< ext "pg_task" >}} | `1.0.0` | {{< badge content="Link" link="https://github.com/RekGRpth/pg_task" >}} | {{< category "TIME" >}} | `pg_task_$v*` | `postgresql-$v-pg-task` |
| {{< ext "pg_later" >}} | `0.3.0` | {{< badge content="Link" link="https://github.com/ChuckHend/pg_later" >}} | {{< category "TIME" >}} | `pg_later_$v` | `postgresql-$v-pg-later` |
| {{< ext "pg_background" >}} | `1.3` | {{< badge content="Link" link="https://github.com/vibhorkum/pg_background" >}} | {{< category "TIME" >}} | `pg_background_$v*` | `postgresql-$v-pg-background` |
| {{< ext "postgis" >}} | `3.6.0` | {{< badge content="Link" link="https://git.osgeo.org/gitea/postgis/postgis" >}} | {{< category "GIS" >}} | `postgis36_$v*` | `postgresql-$v-postgis-3 postgresql-$v-postgis-3-scripts` |
| {{< ext "pgrouting" >}} | `3.8.0` | {{< badge content="Link" link="https://github.com/pgRouting/pgrouting" >}} | {{< category "GIS" >}} | `pgrouting_$v*` | `postgresql-$v-pgrouting postgresql-$v-pgrouting-scripts` |
| {{< ext "pointcloud" >}} | `1.2.5` | {{< badge content="Link" link="https://github.com/pgpointcloud/pointcloud" >}} | {{< category "GIS" >}} | `pointcloud_$v*` | `postgresql-$v-pointcloud` |
| {{< ext "h3" "pg_h3" >}} | `4.2.3` | {{< badge content="Link" link="https://github.com/zachasme/h3-pg" >}} | {{< category "GIS" >}} | `h3-pg_$v*` | `postgresql-$v-h3` |
| {{< ext "q3c" >}} | `2.0.1` | {{< badge content="Link" link="https://github.com/segasai/q3c" >}} | {{< category "GIS" >}} | `q3c_$v*` | `postgresql-$v-q3c` |
| {{< ext "ogr_fdw" >}} | `1.1.7` | {{< badge content="Link" link="https://github.com/pramsey/pgsql-ogr-fdw" >}} | {{< category "GIS" >}} | `ogr_fdw_$v*` | `postgresql-$v-ogr-fdw` |
| {{< ext "geoip" >}} | `0.3.0` | {{< badge content="Link" link="https://github.com/tvondra/geoip" >}} | {{< category "GIS" >}} | `geoip_$v` | `postgresql-$v-geoip` |
| {{< ext "pg_polyline" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/yihong0618/pg_polyline" >}} | {{< category "GIS" >}} | `pg_polyline_$v` | `postgresql-$v-pg-polyline` |
| {{< ext "pg_geohash" >}} | `1.0` | {{< badge content="Link" link="https://github.com/jistok/pg_geohash" >}} | {{< category "GIS" >}} | `pg_geohash_$v*` | `postgresql-$v-pg-geohash` |
| {{< ext "mobilitydb" >}} | `1.2.0` | {{< badge content="Link" link="https://github.com/MobilityDB/MobilityDB" >}} | {{< category "GIS" >}} | - | `postgresql-$v-mobilitydb` |
| {{< ext "tzf" "pg_tzf" >}} | `0.2.2` | {{< badge content="Link" link="https://github.com/ringsaturn/pg-tzf" >}} | {{< category "GIS" >}} | `pg_tzf_$v` | `postgresql-$v-tzf` |
| {{< ext "earthdistance" >}} | `1.2` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/earthdistance.html" >}} | {{< category "GIS" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "vector" "pgvector" >}} | `0.8.1` | {{< badge content="Link" link="https://github.com/pgvector/pgvector" >}} | {{< category "RAG" >}} | `pgvector_$v*` | `postgresql-$v-pgvector` |
| {{< ext "vchord" >}} | `0.5.1` | {{< badge content="Link" link="https://github.com/tensorchord/VectorChord" >}} | {{< category "RAG" >}} | `vchord_$v` | `postgresql-$v-vchord` |
| {{< ext "vectorscale" "pgvectorscale" >}} | `0.8.0` | {{< badge content="Link" link="https://github.com/timescale/pgvectorscale" >}} | {{< category "RAG" >}} | `pgvectorscale_$v` | `postgresql-$v-pgvectorscale` |
| {{< ext "vectorize" "pg_vectorize" >}} | `0.22.2` | {{< badge content="Link" link="https://github.com/ChuckHend/pg_vectorize" >}} | {{< category "RAG" >}} | `pg_vectorize_$v` | `postgresql-$v-pg-vectorize` |
| {{< ext "pg_similarity" >}} | `1.0` | {{< badge content="Link" link="https://github.com/eulerto/pg_similarity" >}} | {{< category "RAG" >}} | `pg_similarity_$v*` | `postgresql-$v-similarity` |
| {{< ext "smlar" >}} | `1.0` | {{< badge content="Link" link="https://github.com/jirutka/smlar" >}} | {{< category "RAG" >}} | `smlar_$v*` | `postgresql-$v-smlar` |
| {{< ext "pg_summarize" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/HexaCluster/pg_summarize" >}} | {{< category "RAG" >}} | `pg_summarize_$v` | `postgresql-$v-pg-summarize` |
| {{< ext "pg_tiktoken" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/kelvich/pg_tiktoken" >}} | {{< category "RAG" >}} | `pg_tiktoken_$v` | `postgresql-$v-pg-tiktoken` |
| {{< ext "pg4ml" >}} | `2.0` | {{< badge content="Link" link="https://gitee.com/guotiecheng/plpgsql_pg4ml" >}} | {{< category "RAG" >}} | `pg4ml_$v` | `postgresql-$v-pg4ml` |
| {{< ext "pgml" >}} | `2.10.0` | {{< badge content="Link" link="https://github.com/postgresml/postgresml" >}} | {{< category "RAG" >}} | `pgml_$v` | `postgresql-$v-pgml` |
| {{< ext "pg_search" >}} | `0.18.1` | {{< badge content="Link" link="https://github.com/paradedb/paradedb/tree/dev/pg_search" >}} | {{< category "FTS" >}} | `pg_search_$v` | `postgresql-$v-pg-search` |
| {{< ext "pgroonga" >}} | `4.0.0` | {{< badge content="Link" link="https://github.com/pgroonga/pgroonga" >}} | {{< category "FTS" >}} | `pgroonga_$v*` | `postgresql-$v-pgroonga` |
| {{< ext "pg_bigm" >}} | `1.2` | {{< badge content="Link" link="https://github.com/pgbigm/pg_bigm" >}} | {{< category "FTS" >}} | `pg_bigm_$v*` | `postgresql-$v-pg-bigm` |
| {{< ext "zhparser" >}} | `2.3` | {{< badge content="Link" link="https://github.com/amutu/zhparser" >}} | {{< category "FTS" >}} | `zhparser_$v*` | `postgresql-$v-zhparser` |
| {{< ext "pg_bestmatch" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/tensorchord/pg_bestmatch.rs" >}} | {{< category "FTS" >}} | `pg_bestmatch_$v` | `postgresql-$v-pg-bestmatch` |
| {{< ext "vchord_bm25" >}} | `0.2.1` | {{< badge content="Link" link="https://github.com/tensorchord/VectorChord-bm25" >}} | {{< category "FTS" >}} | `vchord_bm25_$v` | `postgresql-$v-vchord-bm25` |
| {{< ext "pg_tokenizer" >}} | `0.1.0` | {{< badge content="Link" link="https://github.com/tensorchord/pg_tokenizer.rs" >}} | {{< category "FTS" >}} | `pg_tokenizer_$v` | `postgresql-$v-pg-tokenizer` |
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
| {{< ext "pg_duckdb" >}} | `0.3.1` | {{< badge content="Link" link="https://github.com/duckdb/pg_duckdb" >}} | {{< category "OLAP" >}} | `pg_duckdb_$v*` | `postgresql-$v-pg-duckdb` |
| {{< ext "pg_mooncake" >}} | `0.1.2` | {{< badge content="Link" link="https://github.com/Mooncake-Labs/pg_mooncake" >}} | {{< category "OLAP" >}} | `pg_mooncake_$v*` | `postgresql-$v-pg-mooncake` |
| {{< ext "duckdb_fdw" >}} | `1.1.2` | {{< badge content="Link" link="https://github.com/alitrack/duckdb_fdw" >}} | {{< category "OLAP" >}} | `duckdb_fdw_$v*` | `postgresql-$v-duckdb-fdw` |
| {{< ext "pg_parquet" >}} | `0.4.3` | {{< badge content="Link" link="https://github.com/CrunchyData/pg_parquet/" >}} | {{< category "OLAP" >}} | `pg_parquet_$v` | `postgresql-$v-pg-parquet` |
| {{< ext "pg_fkpart" >}} | `1.7.0` | {{< badge content="Link" link="https://github.com/lemoineat/pg_fkpart" >}} | {{< category "OLAP" >}} | `pg_fkpart_$v` | `postgresql-$v-pg-fkpart` |
| {{< ext "pg_partman" >}} | `5.2.4` | {{< badge content="Link" link="https://github.com/pgpartman/pg_partman" >}} | {{< category "OLAP" >}} | `pg_partman_$v*` | `postgresql-$v-partman` |
| {{< ext "plproxy" >}} | `2.11.0` | {{< badge content="Link" link="https://github.com/plproxy/plproxy" >}} | {{< category "OLAP" >}} | `plproxy_$v*` | `postgresql-$v-plproxy` |
| {{< ext "pg_strom" >}} | `6.0` | {{< badge content="Link" link="https://github.com/heterodb/pg-strom" >}} | {{< category "OLAP" >}} | `pg_strom_$v*` | - |
| {{< ext "tablefunc" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/tablefunc.html" >}} | {{< category "OLAP" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "age" >}} | `1.5.0` | {{< badge content="Link" link="https://github.com/apache/age" >}} | {{< category "FEAT" >}} | `apache-age_$v*` | `postgresql-$v-age` |
| {{< ext "hll" >}} | `2.18` | {{< badge content="Link" link="https://github.com/citusdata/postgresql-hll" >}} | {{< category "FEAT" >}} | `hll_$v*` | `postgresql-$v-hll` |
| {{< ext "rum" >}} | `1.3.14` | {{< badge content="Link" link="https://github.com/postgrespro/rum" >}} | {{< category "FEAT" >}} | `rum_$v` | `postgresql-$v-rum` |
| {{< ext "pg_graphql" >}} | `1.5.11` | {{< badge content="Link" link="https://github.com/supabase/pg_graphql" >}} | {{< category "FEAT" >}} | `pg_graphql_$v` | `postgresql-$v-pg-graphql` |
| {{< ext "pg_jsonschema" >}} | `0.3.3` | {{< badge content="Link" link="https://github.com/supabase/pg_jsonschema" >}} | {{< category "FEAT" >}} | `pg_jsonschema_$v` | `postgresql-$v-pg-jsonschema` |
| {{< ext "jsquery" >}} | `1.2` | {{< badge content="Link" link="https://github.com/postgrespro/jsquery" >}} | {{< category "FEAT" >}} | `jsquery_$v*` | `postgresql-$v-jsquery` |
| {{< ext "pg_hint_plan" >}} | `1.7.1` | {{< badge content="Link" link="https://github.com/ossc-db/pg_hint_plan" >}} | {{< category "FEAT" >}} | `pg_hint_plan_$v*` | `postgresql-$v-pg-hint-plan` |
| {{< ext "hypopg" >}} | `1.4.2` | {{< badge content="Link" link="https://github.com/HypoPG/hypopg" >}} | {{< category "FEAT" >}} | `hypopg_$v*` | `postgresql-$v-hypopg` |
| {{< ext "index_advisor" >}} | `0.2.0` | {{< badge content="Link" link="https://github.com/supabase/index_advisor" >}} | {{< category "FEAT" >}} | `index_advisor_$v` | `postgresql-$v-index-advisor` |
| {{< ext "plan_filter" "pg_plan_filter" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/pgexperts/pg_plan_filter" >}} | {{< category "FEAT" >}} | `pg_plan_filter_$v*` | `postgresql-$v-pg-plan-filter` |
| {{< ext "imgsmlr" >}} | `1.0` | {{< badge content="Link" link="https://github.com/postgrespro/imgsmlr" >}} | {{< category "FEAT" >}} | `imgsmlr_$v*` | `postgresql-$v-imgsmlr` |
| {{< ext "pg_ivm" >}} | `1.12` | {{< badge content="Link" link="https://github.com/sraoss/pg_ivm" >}} | {{< category "FEAT" >}} | `pg_ivm_$v*` | `postgresql-$v-pg-ivm` |
| {{< ext "pg_incremental" >}} | `1.2.0` | {{< badge content="Link" link="https://github.com/CrunchyData/pg_incremental" >}} | {{< category "FEAT" >}} | `pg_incremental_$v*` | `postgresql-$v-pg-incremental` |
| {{< ext "pgmq" >}} | `1.5.1` | {{< badge content="Link" link="https://github.com/pgmq/pgmq" >}} | {{< category "FEAT" >}} | `pgmq_$v` | `postgresql-$v-pgmq` |
| {{< ext "pgq" >}} | `3.5.1` | {{< badge content="Link" link="https://github.com/pgq/pgq" >}} | {{< category "FEAT" >}} | `pgq_$v*` | `postgresql-$v-pgq3` |
| {{< ext "orioledb" >}} | `1.5` | {{< badge content="Link" link="https://github.com/orioledb/orioledb" >}} | {{< category "FEAT" >}} | `orioledb_$v*` | `oriolepg-$v-orioledb` |
| {{< ext "pg_cardano" >}} | `1.0.5` | {{< badge content="Link" link="https://github.com/Fell-x27/pg_cardano" >}} | {{< category "FEAT" >}} | `pg_cardano_$v` | `postgresql-$v-pg-cardano` |
| {{< ext "rdkit" >}} | `202503.1` | {{< badge content="Link" link="https://github.com/rdkit/rdkit" >}} | {{< category "FEAT" >}} | - | `postgresql-$v-rdkit` |
| {{< ext "omni" "omnigres" >}} | `0.2.9` | {{< badge content="Link" link="https://github.com/omnigres/omnigres" >}} | {{< category "FEAT" >}} | `omnigres_$v` | `postgresql-$v-omnigres` |
| {{< ext "bloom" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/bloom.html" >}} | {{< category "FEAT" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "pg_tle" >}} | `1.5.1` | {{< badge content="Link" link="https://github.com/aws/pg_tle" >}} | {{< category "LANG" >}} | `pg_tle_$v*` | `postgresql-$v-pg-tle` |
| {{< ext "plv8" >}} | `3.2.4` | {{< badge content="Link" link="https://github.com/plv8/plv8" >}} | {{< category "LANG" >}} | `plv8_$v*` | `postgresql-$v-plv8` |
| {{< ext "pllua" >}} | `2.0.12` | {{< badge content="Link" link="https://github.com/pllua/pllua" >}} | {{< category "LANG" >}} | `pllua_$v*` | `postgresql-$v-pllua` |
| {{< ext "plprql" >}} | `1.0.0` | {{< badge content="Link" link="https://github.com/kaspermarstal/plprql" >}} | {{< category "LANG" >}} | `plprql_$v` | `postgresql-$v-plprql` |
| {{< ext "pldbgapi" "pldebugger" >}} | `1.8` | {{< badge content="Link" link="https://github.com/EnterpriseDB/pldebugger" >}} | {{< category "LANG" >}} | `pldebugger_$v*` | `postgresql-$v-pldebugger` |
| {{< ext "plpgsql_check" >}} | `2.8.2` | {{< badge content="Link" link="https://github.com/okbob/plpgsql_check" >}} | {{< category "LANG" >}} | `plpgsql_check_$v*` | `postgresql-$v-plpgsql-check` |
| {{< ext "plprofiler" >}} | `4.2.5` | {{< badge content="Link" link="https://github.com/bigsql/plprofiler" >}} | {{< category "LANG" >}} | `plprofiler_$v*` | `postgresql-$v-plprofiler` |
| {{< ext "plsh" >}} | `1.20220917` | {{< badge content="Link" link="https://github.com/petere/plsh" >}} | {{< category "LANG" >}} | `plsh_$v*` | `postgresql-$v-plsh` |
| {{< ext "pljava" >}} | `1.6.9` | {{< badge content="Link" link="https://github.com/tada/pljava" >}} | {{< category "LANG" >}} | `pljava_$v*` | `postgresql-$v-pljava` |
| {{< ext "plr" >}} | `8.4.8` | {{< badge content="Link" link="https://github.com/postgres-plr/plr" >}} | {{< category "LANG" >}} | `plr_$v*` | `postgresql-$v-plr` |
| {{< ext "pgtap" >}} | `1.3.3` | {{< badge content="Link" link="https://github.com/theory/pgtap" >}} | {{< category "LANG" >}} | `pgtap_$v*` | `postgresql-$v-pgtap` |
| {{< ext "faker" >}} | `0.5.3` | {{< badge content="Link" link="https://github.com/anpandu/postgresql_faker" >}} | {{< category "LANG" >}} | `postgresql_faker_$v*` | - |
| {{< ext "dbt2" >}} | `0.45.0` | {{< badge content="Link" link="https://github.com/osdldbt/dbt2" >}} | {{< category "LANG" >}} | `dbt2-pg$v-extensions*` | - |
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
| {{< ext "roaringbitmap" >}} | `0.5.4` | {{< badge content="Link" link="https://github.com/ChenHuajun/pg_roaringbitmap" >}} | {{< category "TYPE" >}} | `pg_roaringbitmap_$v*` | `postgresql-$v-roaringbitmap` |
| {{< ext "pgfaceting" >}} | `0.2.0` | {{< badge content="Link" link="https://github.com/cybertec-postgresql/pgfaceting" >}} | {{< category "TYPE" >}} | `pgfaceting_$v` | `postgresql-$v-pgfaceting` |
| {{< ext "pg_sphere" "pgsphere" >}} | `1.5.1` | {{< badge content="Link" link="https://github.com/postgrespro/pgsphere" >}} | {{< category "TYPE" >}} | `pgsphere_$v*` | `postgresql-$v-pgsphere` |
| {{< ext "country" "pg_country" >}} | `0.0.3` | {{< badge content="Link" link="https://github.com/adjust/pg-country" >}} | {{< category "TYPE" >}} | `pg_country_$v*` | `postgresql-$v-pg-country` |
| {{< ext "pg_xenophile" >}} | `0.8.3` | {{< badge content="Link" link="https://github.com/bigsmoke/pg_xenophile" >}} | {{< category "TYPE" >}} | `pg_xenophile_$v` | `postgresql-$v-pg-xenophile` |
| {{< ext "currency" "pg_currency" >}} | `0.0.3` | {{< badge content="Link" link="https://github.com/adjust/pg-currency" >}} | {{< category "TYPE" >}} | `pg_currency_$v*` | `postgresql-$v-pg-currency` |
| {{< ext "collection" "pg_collection" >}} | `1.0.0` | {{< badge content="Link" link="https://github.com/aws/pgcollection" >}} | {{< category "TYPE" >}} | `pgcollection_$v*` | `postgresql-$v-collection` |
| {{< ext "pgmp" >}} | `1.0.5` | {{< badge content="Link" link="https://github.com/dvarrazzo/pgmp/" >}} | {{< category "TYPE" >}} | `pgmp_$v*` | `postgresql-$v-pgmp` |
| {{< ext "numeral" >}} | `1.3` | {{< badge content="Link" link="https://github.com/df7cb/postgresql-numeral" >}} | {{< category "TYPE" >}} | `numeral_$v*` | `postgresql-$v-numeral` |
| {{< ext "pg_rational" >}} | `0.0.2` | {{< badge content="Link" link="https://github.com/begriffs/pg_rational" >}} | {{< category "TYPE" >}} | `pg_rational_$v*` | `postgresql-$v-rational` |
| {{< ext "uint" "pguint" >}} | `1.20231206` | {{< badge content="Link" link="https://github.com/petere/pguint" >}} | {{< category "TYPE" >}} | `pguint_$v*` | `postgresql-$v-pguint` |
| {{< ext "uint128" "pg_uint128" >}} | `1.1.0` | {{< badge content="Link" link="https://github.com/pg-uint/pg-uint128" >}} | {{< category "TYPE" >}} | `pg_uint128_$v*` | `postgresql-$v-pg-uint128` |
| {{< ext "hashtypes" >}} | `0.1.5` | {{< badge content="Link" link="https://github.com/adjust/hashtypes/" >}} | {{< category "TYPE" >}} | `hashtypes_$v*` | `postgresql-$v-hashtypes` |
| {{< ext "ip4r" >}} | `2.4.2` | {{< badge content="Link" link="https://github.com/RhodiumToad/ip4r" >}} | {{< category "TYPE" >}} | `ip4r_$v*` | `postgresql-$v-ip4r` |
| {{< ext "pg_duration" >}} | `1.0.2` | {{< badge content="Link" link="https://github.com/jkosh44/pg_duration" >}} | {{< category "TYPE" >}} | `pg_duration_$v*` | `postgresql-$v-pg-duration` |
| {{< ext "uri" "pg_uri" >}} | `1.20151224` | {{< badge content="Link" link="https://github.com/petere/pguri" >}} | {{< category "TYPE" >}} | `pg_uri_$v*` | `postgresql-$v-pg-uri` |
| {{< ext "emailaddr" "pgemailaddr" >}} | `0` | {{< badge content="Link" link="https://github.com/petere/pgemailaddr" >}} | {{< category "TYPE" >}} | `pg_emailaddr_$v*` | `postgresql-$v-pg-emailaddr` |
| {{< ext "acl" "pg_acl" >}} | `1.0.4` | {{< badge content="Link" link="https://github.com/arkhipov/acl" >}} | {{< category "TYPE" >}} | `acl_$v*` | `postgresql-$v-acl` |
| {{< ext "debversion" >}} | `1.2.0` | {{< badge content="Link" link="https://github.com/ATIX-AG/postgresql-debversion-evr" >}} | {{< category "TYPE" >}} | - | `postgresql-$v-debversion` |
| {{< ext "pg_rrule" >}} | `0.2.0` | {{< badge content="Link" link="https://github.com/petropavel13/pg_rrule" >}} | {{< category "TYPE" >}} | - | `postgresql-$v-pg-rrule` |
| {{< ext "timestamp9" >}} | `1.4.0` | {{< badge content="Link" link="https://github.com/optiver/timestamp9" >}} | {{< category "TYPE" >}} | `timestamp9_$v*` | `postgresql-$v-timestamp9` |
| {{< ext "chkpass" >}} | `1.0` | {{< badge content="Link" link="https://github.com/lacanoid/chkpass" >}} | {{< category "TYPE" >}} | `chkpass_$v*` | `postgresql-$v-chkpass` |
| {{< ext "isn" >}} | `1.2` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/isn.html" >}} | {{< category "TYPE" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "seg" >}} | `1.4` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/seg.html" >}} | {{< category "TYPE" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "cube" >}} | `1.5` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/cube.html" >}} | {{< category "TYPE" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "ltree" >}} | `1.3` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/ltree.html" >}} | {{< category "TYPE" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "hstore" >}} | `1.8` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/hstore.html" >}} | {{< category "TYPE" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "citext" >}} | `1.6` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/citext.html" >}} | {{< category "TYPE" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "xml2" >}} | `1.1` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/xml2.html" >}} | {{< category "TYPE" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "gzip" "pg_gzip" >}} | `1.0.1` | {{< badge content="Link" link="https://github.com/pramsey/pgsql-gzip" >}} | {{< category "UTIL" >}} | `pgsql_gzip_$v*` | `postgresql-$v-gzip` |
| {{< ext "bzip" "pg_bzip" >}} | `1.0.0` | {{< badge content="Link" link="https://github.com/steve-chavez/pg_bzip" >}} | {{< category "UTIL" >}} | `pg_bzip_$v*` | `postgresql-$v-bzip` |
| {{< ext "zstd" "pg_zstd" >}} | `1.1.2` | {{< badge content="Link" link="https://github.com/grahamedgecombe/pgzstd" >}} | {{< category "UTIL" >}} | `pg_zstd_$v*` | `postgresql-$v-zstd` |
| {{< ext "http" "pg_http" >}} | `1.7.0` | {{< badge content="Link" link="https://github.com/pramsey/pgsql-http" >}} | {{< category "UTIL" >}} | `pgsql_http_$v*` | `postgresql-$v-http` |
| {{< ext "pg_net" >}} | `0.9.2` | {{< badge content="Link" link="https://github.com/supabase/pg_net" >}} | {{< category "UTIL" >}} | `pg_net_$v*` | `postgresql-$v-pg-net` |
| {{< ext "pg_curl" >}} | `2.4.5` | {{< badge content="Link" link="https://github.com/RekGRpth/pg_curl" >}} | {{< category "UTIL" >}} | `pg_curl_$v*` | `postgresql-$v-pg-curl` |
| {{< ext "pgjq" >}} | `0.1.0` | {{< badge content="Link" link="https://github.com/Florents-Tselai/pgJQ" >}} | {{< category "UTIL" >}} | `pgjq_$v*` | `postgresql-$v-pgjq` |
| {{< ext "pgjwt" >}} | `0.2.0` | {{< badge content="Link" link="https://github.com/michelp/pgjwt" >}} | {{< category "UTIL" >}} | `pgjwt_$v` | `postgresql-$v-pgjwt` |
| {{< ext "pg_smtp_client" >}} | `0.2.0` | {{< badge content="Link" link="https://github.com/brianpursley/pg_smtp_client" >}} | {{< category "UTIL" >}} | `pg_smtp_client_$v` | `postgresql-$v-pg-smtp-client` |
| {{< ext "pg_html5_email_address" >}} | `1.2.3` | {{< badge content="Link" link="https://github.com/bigsmoke/pg_html5_email_address" >}} | {{< category "UTIL" >}} | `pg_html5_email_address_$v` | `postgresql-$v-pg-html5-email-address` |
| {{< ext "url_encode" >}} | `1.2.5` | {{< badge content="Link" link="https://github.com/okbob/url_encode" >}} | {{< category "UTIL" >}} | `url_encode_$v*` | `postgresql-$v-url-encode` |
| {{< ext "pgsql_tweaks" >}} | `0.11.3` | {{< badge content="Link" link="https://codeberg.org/pgsql_tweaks/pgsql_tweaks" >}} | {{< category "UTIL" >}} | `pgsql_tweaks_$v` | `postgresql-$v-pgsql-tweaks` |
| {{< ext "pg_extra_time" >}} | `2.0.0` | {{< badge content="Link" link="https://github.com/bigsmoke/pg_extra_time" >}} | {{< category "UTIL" >}} | `pg_extra_time_$v` | `postgresql-$v-pg-extra-time` |
| {{< ext "pgpcre" >}} | `1` | {{< badge content="Link" link="https://github.com/petere/pgpcre" >}} | {{< category "UTIL" >}} | `pgpcre_$v` | `postgresql-$v-pgpcre` |
| {{< ext "icu_ext" >}} | `1.10.0` | {{< badge content="Link" link="https://github.com/dverite/icu_ext" >}} | {{< category "UTIL" >}} | `icu_ext_$v*` | `postgresql-$v-icu-ext` |
| {{< ext "pgqr" >}} | `1.0` | {{< badge content="Link" link="https://github.com/AbdulYadi/pgqr" >}} | {{< category "UTIL" >}} | `pgqr_$v*` | `postgresql-$v-pgqr` |
| {{< ext "pg_protobuf" >}} | `1.0` | {{< badge content="Link" link="https://github.com/afiskon/pg_protobuf" >}} | {{< category "UTIL" >}} | `pg_protobuf_$v` | `postgresql-$v-pg-protobuf` |
| {{< ext "envvar" >}} | `1.0.1` | {{< badge content="Link" link="https://github.com/theory/pg-envvar" >}} | {{< category "UTIL" >}} | `pg_envvar_$v*` | `postgresql-$v-pg-envvar` |
| {{< ext "floatfile" >}} | `1.3.1` | {{< badge content="Link" link="https://github.com/pjungwir/floatfile" >}} | {{< category "UTIL" >}} | `floatfile_$v*` | `postgresql-$v-floatfile` |
| {{< ext "pg_render" >}} | `0.1.2` | {{< badge content="Link" link="https://github.com/mkaski/pg_render" >}} | {{< category "UTIL" >}} | `pg_render_$v` | `postgresql-$v-pg-render` |
| {{< ext "pg_readme" >}} | `0.7.0` | {{< badge content="Link" link="https://github.com/bigsmoke/pg_readme" >}} | {{< category "UTIL" >}} | `pg_readme_$v` | `postgresql-$v-pg-readme` |
| {{< ext "ddl_historization" >}} | `0.0.7` | {{< badge content="Link" link="https://github.com/rodo/pg_ddl_historization" >}} | {{< category "UTIL" >}} | `ddl_historization_$v` | `postgresql-$v-ddl-historization` |
| {{< ext "data_historization" >}} | `1.1.0` | {{< badge content="Link" link="https://github.com/rodo/postgresql-data-historization" >}} | {{< category "UTIL" >}} | `data_historization_$v` | `postgresql-$v-data-historization` |
| {{< ext "schedoc" "pg_schedoc" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/ZeroGachis/pg_schedoc" >}} | {{< category "UTIL" >}} | `pg_schedoc_$v` | `postgresql-$v-pg-schedoc` |
| {{< ext "hashlib" "pg_hashlib" >}} | `1.1` | {{< badge content="Link" link="https://github.com/markokr/pghashlib" >}} | {{< category "UTIL" >}} | `pg_hashlib_$v` | `postgresql-$v-pg-hashlib` |
| {{< ext "xxhash" "pg_xxhash" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/hatarist/pg_xxhash" >}} | {{< category "UTIL" >}} | `pg_xxhash_$v*` | `postgresql-$v-pg-xxhash` |
| {{< ext "shacrypt" >}} | `1.1` | {{< badge content="Link" link="https://github.com/dverite/postgres-shacrypt" >}} | {{< category "UTIL" >}} | `postgres_shacrypt_$v*` | `postgresql-$v-shacrypt` |
| {{< ext "cryptint" >}} | `1.0.0` | {{< badge content="Link" link="https://github.com/dverite/cryptint" >}} | {{< category "UTIL" >}} | `cryptint_$v*` | `postgresql-$v-cryptint` |
| {{< ext "pguecc" "pg_ecdsa" >}} | `1.0` | {{< badge content="Link" link="https://github.com/ameensol/pg-ecdsa" >}} | {{< category "UTIL" >}} | `pg_ecdsa_$v*` | `postgresql-$v-pg-ecdsa` |
| {{< ext "sparql" "pgsparql" >}} | `1.0` | {{< badge content="Link" link="https://github.com/lacanoid/pgsparql" >}} | {{< category "UTIL" >}} | `pgsparql_$v` | `postgresql-$v-pgsparql` |
| {{< ext "pg_idkit" >}} | `0.3.1` | {{< badge content="Link" link="https://github.com/VADOSWARE/pg_idkit" >}} | {{< category "FUNC" >}} | `pg_idkit_$v` | `postgresql-$v-pg-idkit` |
| {{< ext "pgx_ulid" >}} | `0.2.0` | {{< badge content="Link" link="https://github.com/pksunkara/pgx_ulid" >}} | {{< category "FUNC" >}} | `pgx_ulid_$v` | `postgresql-$v-pgx-ulid` |
| {{< ext "pg_uuidv7" >}} | `1.6.0` | {{< badge content="Link" link="https://github.com/fboulnois/pg_uuidv7" >}} | {{< category "FUNC" >}} | `pg_uuidv7_$v*` | `postgresql-$v-pg-uuidv7` |
| {{< ext "permuteseq" >}} | `1.2.2` | {{< badge content="Link" link="https://github.com/dverite/permuteseq" >}} | {{< category "FUNC" >}} | `permuteseq_$v*` | `postgresql-$v-permuteseq` |
| {{< ext "pg_hashids" >}} | `1.3` | {{< badge content="Link" link="https://github.com/iCyberon/pg_hashids" >}} | {{< category "FUNC" >}} | `pg_hashids_$v*` | `postgresql-$v-pg-hashids` |
| {{< ext "sequential_uuids" >}} | `1.0.3` | {{< badge content="Link" link="https://github.com/tvondra/sequential-uuids" >}} | {{< category "FUNC" >}} | `sequential_uuids_$v` | `postgresql-$v-sequential-uuids` |
| {{< ext "topn" >}} | `2.7.0` | {{< badge content="Link" link="https://github.com/citusdata/postgresql-topn" >}} | {{< category "FUNC" >}} | `topn_$v*` | `postgresql-$v-topn` |
| {{< ext "quantile" >}} | `1.1.8` | {{< badge content="Link" link="https://github.com/tvondra/quantile" >}} | {{< category "FUNC" >}} | `quantile_$v*` | `postgresql-$v-quantile` |
| {{< ext "lower_quantile" >}} | `1.0.3` | {{< badge content="Link" link="https://github.com/tvondra/lower_quantile" >}} | {{< category "FUNC" >}} | `lower_quantile_$v*` | `postgresql-$v-lower-quantile` |
| {{< ext "count_distinct" >}} | `3.0.2` | {{< badge content="Link" link="https://github.com/tvondra/count_distinct" >}} | {{< category "FUNC" >}} | `count_distinct_$v*` | `postgresql-$v-count-distinct` |
| {{< ext "omnisketch" >}} | `1.0.2` | {{< badge content="Link" link="https://github.com/tvondra/omnisketch" >}} | {{< category "FUNC" >}} | `omnisketch_$v*` | `postgresql-$v-omnisketch` |
| {{< ext "ddsketch" >}} | `1.0.1` | {{< badge content="Link" link="https://github.com/tvondra/ddsketch" >}} | {{< category "FUNC" >}} | `ddsketch_$v*` | `postgresql-$v-ddsketch` |
| {{< ext "vasco" >}} | `0.1.0` | {{< badge content="Link" link="https://github.com/Florents-Tselai/vasco" >}} | {{< category "FUNC" >}} | `vasco_$v*` | `postgresql-$v-vasco` |
| {{< ext "xicor" "pgxicor" >}} | `0.1.0` | {{< badge content="Link" link="https://github.com/Florents-Tselai/pgxicor" >}} | {{< category "FUNC" >}} | `pgxicor_$v*` | `postgresql-$v-pgxicor` |
| {{< ext "tdigest" >}} | `1.4.3` | {{< badge content="Link" link="https://github.com/tvondra/tdigest" >}} | {{< category "FUNC" >}} | `tdigest_$v*` | `postgresql-$v-tdigest` |
| {{< ext "first_last_agg" >}} | `0.1.4` | {{< badge content="Link" link="https://github.com/wulczer/first_last_agg" >}} | {{< category "FUNC" >}} | `first_last_agg_$v` | `postgresql-$v-first-last-agg` |
| {{< ext "extra_window_functions" >}} | `1.0` | {{< badge content="Link" link="https://github.com/xocolatl/extra_window_functions" >}} | {{< category "FUNC" >}} | `extra_window_functions_$v*` | `postgresql-$v-extra-window-functions` |
| {{< ext "floatvec" >}} | `1.1.1` | {{< badge content="Link" link="https://github.com/pjungwir/floatvec" >}} | {{< category "FUNC" >}} | `floatvec_$v*` | `postgresql-$v-floatvec` |
| {{< ext "aggs_for_vecs" >}} | `1.4.0` | {{< badge content="Link" link="https://github.com/pjungwir/aggs_for_vecs" >}} | {{< category "FUNC" >}} | `aggs_for_vecs_$v*` | `postgresql-$v-aggs-for-vecs` |
| {{< ext "aggs_for_arrays" >}} | `1.3.3` | {{< badge content="Link" link="https://github.com/pjungwir/aggs_for_arrays" >}} | {{< category "FUNC" >}} | `aggs_for_arrays_$v*` | `postgresql-$v-aggs-for-arrays` |
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
| {{< ext "pg_checksums" >}} | `1.2` | {{< badge content="Link" link="https://github.com/credativ/pg_checksums" >}} | {{< category "ADMIN" >}} | `pg_checksums_$v*` | `postgresql-$v-pg-checksums` |
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
| {{< ext "pg_show_plans" >}} | `2.1.6` | {{< badge content="Link" link="https://github.com/cybertec-postgresql/pg_show_plans" >}} | {{< category "STAT" >}} | `pg_show_plans_$v*` | `postgresql-$v-show-plans` |
| {{< ext "pg_stat_kcache" >}} | `2.3.0` | {{< badge content="Link" link="https://github.com/powa-team/pg_stat_kcache" >}} | {{< category "STAT" >}} | `pg_stat_kcache_$v*` | `postgresql-$v-pg-stat-kcache` |
| {{< ext "pg_stat_monitor" >}} | `2.2.0` | {{< badge content="Link" link="https://github.com/percona/pg_stat_monitor" >}} | {{< category "STAT" >}} | `pg_stat_monitor_$v*` | `postgresql-$v-pg-stat-monitor` |
| {{< ext "pg_qualstats" >}} | `2.1.2` | {{< badge content="Link" link="https://github.com/powa-team/pg_qualstats" >}} | {{< category "STAT" >}} | `pg_qualstats_$v*` | `postgresql-$v-pg-qualstats` |
| {{< ext "pg_store_plans" >}} | `1.8` | {{< badge content="Link" link="https://github.com/ossc-db/pg_store_plans" >}} | {{< category "STAT" >}} | `pg_store_plans_$v*` | `postgresql-$v-pg-store-plan` |
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
| {{< ext "explain_ui" "pg_explain_ui" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/davidgomes/pg-explain-ui" >}} | {{< category "STAT" >}} | `pg_explain_ui_$v` | `postgresql-$v-pg-explain-ui` |
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
| {{< ext "supautils" >}} | `2.10.0` | {{< badge content="Link" link="https://github.com/supabase/supautils" >}} | {{< category "SEC" >}} | `supautils_$v` | `postgresql-$v-supautils` |
| {{< ext "pgsodium" >}} | `3.1.9` | {{< badge content="Link" link="https://github.com/michelp/pgsodium" >}} | {{< category "SEC" >}} | `pgsodium_$v*` | `postgresql-$v-pgsodium` |
| {{< ext "supabase_vault" "pg_vault" >}} | `0.3.1` | {{< badge content="Link" link="https://github.com/supabase/vault" >}} | {{< category "SEC" >}} | `vault_$v*` | `postgresql-$v-vault` |
| {{< ext "pg_session_jwt" >}} | `0.3.1` | {{< badge content="Link" link="https://github.com/neondatabase/pg_session_jwt" >}} | {{< category "SEC" >}} | `pg_session_jwt_$v` | `postgresql-$v-pg-session-jwt` |
| {{< ext "anon" "pg_anon" >}} | `2.3.0` | {{< badge content="Link" link="https://gitlab.com/dalibo/postgresql_anonymizer/" >}} | {{< category "SEC" >}} | `pg_anon_$v` | `postgresql-$v-pg-anon` |
| {{< ext "pg_tde" >}} | `1.0` | {{< badge content="Link" link="https://github.com/Percona-Lab/pg_tde" >}} | {{< category "SEC" >}} | `percona-postgresql$v` | `percona-postgresql-$v` |
| {{< ext "pgsmcrypto" >}} | `0.1.0` | {{< badge content="Link" link="https://github.com/zhuobie/pgsmcrypto" >}} | {{< category "SEC" >}} | `pgsmcrypto_$v` | `postgresql-$v-pgsmcrypto` |
| {{< ext "pgaudit" >}} | `17.1` | {{< badge content="Link" link="https://github.com/pgaudit/pgaudit" >}} | {{< category "SEC" >}} | `pgaudit_$v*` | `postgresql-$v-pgaudit` |
| {{< ext "pgauditlogtofile" >}} | `1.7.1` | {{< badge content="Link" link="https://github.com/fmbiete/pgauditlogtofile" >}} | {{< category "SEC" >}} | `pgauditlogtofile_$v*` | `postgresql-$v-pgauditlogtofile` |
| {{< ext "pg_auth_mon" >}} | `3.0` | {{< badge content="Link" link="https://github.com/RafiaSabih/pg_auth_mon" >}} | {{< category "SEC" >}} | `pg_auth_mon_$v*` | `postgresql-$v-pg-auth-mon` |
| {{< ext "credcheck" >}} | `3.0` | {{< badge content="Link" link="https://github.com/MigOpsRepos/credcheck" >}} | {{< category "SEC" >}} | `credcheck_$v*` | `postgresql-$v-credcheck` |
| {{< ext "pgcryptokey" >}} | `0.85` | {{< badge content="Link" link="https://momjian.us/download/pgcryptokey/" >}} | {{< category "SEC" >}} | `pgcryptokey_$v` | `postgresql-$v-pgcryptokey` |
| {{< ext "pg_jobmon" >}} | `1.4.1` | {{< badge content="Link" link="https://github.com/omniti-labs/pg_jobmon" >}} | {{< category "SEC" >}} | `pg_jobmon_$v` | `postgresql-$v-pg-jobmon` |
| {{< ext "logerrors" >}} | `2.1.3` | {{< badge content="Link" link="https://github.com/munakoiso/logerrors" >}} | {{< category "SEC" >}} | `logerrors_$v*` | `postgresql-$v-logerrors` |
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
| {{< ext "wrappers" >}} | `0.5.4` | {{< badge content="Link" link="https://github.com/supabase/wrappers" >}} | {{< category "FDW" >}} | `wrappers_$v` | `postgresql-$v-wrappers` |
| {{< ext "multicorn" >}} | `3.0` | {{< badge content="Link" link="https://github.com/pgsql-io/multicorn2" >}} | {{< category "FDW" >}} | `multicorn2_$v*` | - |
| {{< ext "odbc_fdw" >}} | `0.5.1` | {{< badge content="Link" link="https://github.com/CartoDB/odbc_fdw" >}} | {{< category "FDW" >}} | `odbc_fdw_$v*` | - |
| {{< ext "jdbc_fdw" >}} | `1.2` | {{< badge content="Link" link="https://github.com/pgspider/jdbc_fdw" >}} | {{< category "FDW" >}} | `jdbc_fdw_$v*` | - |
| {{< ext "pgspider_ext" >}} | `1.3.0` | {{< badge content="Link" link="https://github.com/pgspider/pgspider_ext" >}} | {{< category "FDW" >}} | `pgspider_ext_$v*` | `postgresql-$v-pgspider-ext` |
| {{< ext "mysql_fdw" >}} | `2.9.2` | {{< badge content="Link" link="https://github.com/EnterpriseDB/mysql_fdw" >}} | {{< category "FDW" >}} | `mysql_fdw_$v*` | `postgresql-$v-mysql-fdw` |
| {{< ext "oracle_fdw" >}} | `2.8.0` | {{< badge content="Link" link="https://github.com/laurenz/oracle_fdw" >}} | {{< category "FDW" >}} | `oracle_fdw_$v*` | `postgresql-$v-oracle-fdw` |
| {{< ext "tds_fdw" >}} | `2.0.4` | {{< badge content="Link" link="https://github.com/tds-fdw/tds_fdw" >}} | {{< category "FDW" >}} | `tds_fdw_$v*` | `postgresql-$v-tds-fdw` |
| {{< ext "db2_fdw" >}} | `7.0.0` | {{< badge content="Link" link="https://github.com/wolfgangbrandl/db2_fdw" >}} | {{< category "FDW" >}} | `db2_fdw_$v*` | - |
| {{< ext "sqlite_fdw" >}} | `2.5.0` | {{< badge content="Link" link="https://github.com/pgspider/sqlite_fdw" >}} | {{< category "FDW" >}} | `sqlite_fdw_$v*` | `postgresql-$v-sqlite-fdw` |
| {{< ext "pgbouncer_fdw" >}} | `1.4.0` | {{< badge content="Link" link="https://github.com/CrunchyData/pgbouncer_fdw" >}} | {{< category "FDW" >}} | `pgbouncer_fdw_$v` | - |
| {{< ext "mongo_fdw" >}} | `1.1` | {{< badge content="Link" link="https://github.com/EnterpriseDB/mongo_fdw" >}} | {{< category "FDW" >}} | `mongo_fdw_$v*` | - |
| {{< ext "redis_fdw" >}} | `1.0` | {{< badge content="Link" link="https://github.com/pg-redis-fdw/redis_fdw" >}} | {{< category "FDW" >}} | `redis_fdw_$v*` | `postgresql-$v-redis-fdw` |
| {{< ext "redis" "pg_redis_pubsub" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/brettlaforge/pg_redis_pubsub" >}} | {{< category "FDW" >}} | `pg_redis_pubsub_$v*` | `postgresql-$v-pg-redis-pubsub` |
| {{< ext "kafka_fdw" >}} | `0.0.3` | {{< badge content="Link" link="https://github.com/adjust/kafka_fdw" >}} | {{< category "FDW" >}} | `kafka_fdw_$v` | `postgresql-$v-kafka-fdw` |
| {{< ext "hdfs_fdw" >}} | `2.3.2` | {{< badge content="Link" link="https://github.com/EnterpriseDB/hdfs_fdw" >}} | {{< category "FDW" >}} | `hdfs_fdw_$v*` | - |
| {{< ext "firebird_fdw" >}} | `1.4.0` | {{< badge content="Link" link="https://github.com/ibarwick/firebird_fdw" >}} | {{< category "FDW" >}} | `firebird_fdw_$v` | `postgresql-$v-firebird-fdw` |
| {{< ext "aws_s3" >}} | `0.0.1` | {{< badge content="Link" link="https://github.com/chimpler/postgres-aws-s3" >}} | {{< category "FDW" >}} | `aws_s3_$v` | `postgresql-$v-aws-s3` |
| {{< ext "log_fdw" >}} | `1.4` | {{< badge content="Link" link="https://github.com/aws/postgresql-logfdw" >}} | {{< category "FDW" >}} | `log_fdw_$v*` | `postgresql-$v-log-fdw` |
| {{< ext "dblink" >}} | `1.2` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/dblink.html" >}} | {{< category "FDW" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "file_fdw" >}} | `1.0` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/file-fdw.html" >}} | {{< category "FDW" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "postgres_fdw" >}} | `1.1` | {{< badge content="Link" link="https://www.postgresql.org/docs/current/postgres-fdw.html" >}} | {{< category "FDW" >}} | `postgresql$v-contrib` | `postgresql-$v` |
| {{< ext "documentdb" >}} | `0.106` | {{< badge content="Link" link="https://github.com/microsoft/documentdb" >}} | {{< category "SIM" >}} | `documentdb_$v*` | `postgresql-$v-documentdb` |
| {{< ext "orafce" >}} | `4.14.4` | {{< badge content="Link" link="https://github.com/orafce/orafce" >}} | {{< category "SIM" >}} | `orafce_$v` | `postgresql-$v-orafce` |
| {{< ext "pgtt" >}} | `4.4` | {{< badge content="Link" link="https://github.com/darold/pgtt" >}} | {{< category "SIM" >}} | `pgtt_$v*` | `postgresql-$v-pgtt` |
| {{< ext "session_variable" >}} | `3.4` | {{< badge content="Link" link="https://github.com/splendiddata/session_variable" >}} | {{< category "SIM" >}} | `session_variable_$v*` | `postgresql-$v-session-variable` |
| {{< ext "pg_statement_rollback" >}} | `1.4` | {{< badge content="Link" link="https://github.com/lzlabs/pg_statement_rollback" >}} | {{< category "SIM" >}} | `pg_statement_rollback_$v*` | `postgresql-$v-pg-statement-rollback` |
| {{< ext "pg_dbms_metadata" >}} | `1.0.0` | {{< badge content="Link" link="https://github.com/HexaCluster/pg_dbms_metadata" >}} | {{< category "SIM" >}} | `pg_dbms_metadata_$v` | - |
| {{< ext "pg_dbms_lock" >}} | `1.0` | {{< badge content="Link" link="https://github.com/HexaCluster/pg_dbms_lock" >}} | {{< category "SIM" >}} | `pg_dbms_lock_$v` | - |
| {{< ext "pg_dbms_job" >}} | `1.5` | {{< badge content="Link" link="https://github.com/MigOpsRepos/pg_dbms_job" >}} | {{< category "SIM" >}} | `pg_dbms_job_$v` | - |
| {{< ext "babelfishpg_common" >}} | `3.3.3` | {{< badge content="Link" link="https://babelfishpg.org/" >}} | {{< category "SIM" >}} | `babelfishpg-common*` | `babelfishpg-common` |
| {{< ext "babelfishpg_tsql" >}} | `3.3.1` | {{< badge content="Link" link="https://babelfishpg.org/" >}} | {{< category "SIM" >}} | `babelfishpg-tsql*` | `babelfishpg-tsql` |
| {{< ext "babelfishpg_tds" >}} | `1.0.0` | {{< badge content="Link" link="https://babelfishpg.org/" >}} | {{< category "SIM" >}} | `babelfishpg-tds*` | `babelfishpg-tds` |
| {{< ext "babelfishpg_money" >}} | `1.1.0` | {{< badge content="Link" link="https://babelfishpg.org/" >}} | {{< category "SIM" >}} | `babelfishpg-money*` | `babelfishpg-money` |
| {{< ext "spat" >}} | `0.1.0a4` | {{< badge content="Link" link="https://github.com/Florents-Tselai/spat" >}} | {{< category "SIM" >}} | `spat_$v*` | `postgresql-$v-spat` |
| {{< ext "pgmemcache" >}} | `2.3.0` | {{< badge content="Link" link="https://github.com/ohmu/pgmemcache" >}} | {{< category "SIM" >}} | `pgmemcache_$v*` | `postgresql-$v-pgmemcache` |
| {{< ext "pglogical" >}} | `2.4.5` | {{< badge content="Link" link="https://github.com/2ndQuadrant/pglogical" >}} | {{< category "ETL" >}} | `pglogical_$v*` | `postgresql-$v-pglogical` |
| {{< ext "pglogical_ticker" >}} | `1.4.1` | {{< badge content="Link" link="https://github.com/enova/pglogical_ticker" >}} | {{< category "ETL" >}} | `pglogical_ticker_$v*` | `postgresql-$v-pglogical-ticker` |
| {{< ext "pgl_ddl_deploy" >}} | `2.2.1` | {{< badge content="Link" link="https://github.com/enova/pgl_ddl_deploy" >}} | {{< category "ETL" >}} | `pgl_ddl_deploy_$v*` | `postgresql-$v-pgl-ddl-deploy` |
| {{< ext "pg_failover_slots" >}} | `1.1.0` | {{< badge content="Link" link="https://github.com/EnterpriseDB/pg_failover_slots" >}} | {{< category "ETL" >}} | `pg_failover_slots_$v*` | `postgresql-$v-pg-failover-slots` |
| {{< ext "db_migrator" >}} | `1.0.0` | {{< badge content="Link" link="https://github.com/cybertec-postgresql/db_migrator" >}} | {{< category "ETL" >}} | `db_migrator_$v` | `postgresql-$v-db-migrator` |
| {{< ext "pgactive" >}} | `2.1.6` | {{< badge content="Link" link="https://github.com/aws/pgactive" >}} | {{< category "ETL" >}} | `pgactive_$v*` | `postgresql-$v-pgactive` |
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

## Extensions

There are 424 available PostgreSQL extensions:

| Extension | PG Versions | Attribute | Description |
|:----------|:------------|:---------:|:------------|
| {{< ext "timescaledb" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | Enables scalable inserts and complex queries for time-series data |
| {{< ext "timescaledb_toolkit" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | Library of analytical hyperfunctions, time-series pipelining, and other SQL utilities |
| {{< ext "timeseries" "pg_timeseries" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Convenience API for time series stack |
| {{< ext "periods" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Provide Standard SQL functionality for PERIODs and SYSTEM VERSIONING |
| {{< ext "temporal_tables" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | temporal tables |
| {{< ext "emaj" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Enables fine-grained write logging and time travel on subsets of the database. |
| {{< ext "table_version" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostgreSQL table versioning extension |
| {{< ext "pg_cron" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | Job scheduler for PostgreSQL |
| {{< ext "pg_task" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | execute any sql command at any specific time at background |
| {{< ext "pg_later" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Run queries now and get results later |
| {{< ext "pg_background" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Run SQL queries in the background |
| {{< ext "postgis" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostGIS geometry and geography spatial types and functions |
| {{< ext "postgis_topology" "postgis" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostGIS topology spatial types and functions |
| {{< ext "postgis_raster" "postgis" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostGIS raster types and functions |
| {{< ext "postgis_sfcgal" "postgis" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | PostGIS SFCGAL functions |
| {{< ext "postgis_tiger_geocoder" "postgis" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | PostGIS tiger geocoder and reverse geocoder |
| {{< ext "address_standardizer" "postgis" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Used to parse an address into constituent elements. Generally used to support geocoding address norm |
| {{< ext "address_standardizer_data_us" "postgis" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Address Standardizer US dataset example |
| {{< ext "pgrouting" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | pgRouting Extension |
| {{< ext "pointcloud" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | data type for lidar point clouds |
| {{< ext "pointcloud_postgis" "pointcloud" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | integration for pointcloud LIDAR data and PostGIS geometry data |
| {{< ext "h3" "pg_h3" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | H3 bindings for PostgreSQL |
| {{< ext "h3_postgis" "pg_h3" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | H3 PostGIS integration |
| {{< ext "q3c" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | q3c sky indexing plugin |
| {{< ext "ogr_fdw" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | foreign-data wrapper for GIS data access |
| {{< ext "geoip" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | IP-based geolocation query |
| {{< ext "pg_polyline" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | Fast Google Encoded Polyline encoding & decoding for postgres |
| {{< ext "pg_geohash" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Handle geohash based functionality for spatial coordinates |
| {{< ext "mobilitydb" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | MobilityDB geospatial trajectory data management & analysis platform |
| {{< ext "tzf" "pg_tzf" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Fast lookup timezone name by GPS coordinates |
| {{< ext "earthdistance" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | calculate great-circle distances on the surface of the Earth |
| {{< ext "vector" "pgvector" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | vector data type and ivfflat and hnsw access methods |
| {{< ext "vchord" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | Vector database plugin for Postgres, written in Rust |
| {{< ext "vectorscale" "pgvectorscale" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Advanced indexing for vector data with DiskANN |
| {{< ext "vectorize" "pg_vectorize" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | The simplest way to do vector search on Postgres |
| {{< ext "pg_similarity" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | support similarity queries |
| {{< ext "smlar" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Effective similarity search |
| {{< ext "pg_summarize" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Text Summarization using LLMs. Built using pgrx |
| {{< ext "pg_tiktoken" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | tiktoken tokenizer for use with OpenAI models in postgres |
| {{< ext "pg4ml" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | Machine learning framework for PostgreSQL |
| {{< ext "pgml" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | Run AL/ML workloads with SQL interface |
| {{< ext "pg_search" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | Full text search for PostgreSQL using BM25 |
| {{< ext "pgroonga" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | Use Groonga as index, fast full text search platform for all languages! |
| {{< ext "pgroonga_database" "pgroonga" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | PGroonga database management module |
| {{< ext "pg_bigm" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | create 2-gram (bigram) index for faster full text search. |
| {{< ext "zhparser" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | a parser for full-text search of Chinese |
| {{< ext "pg_bestmatch" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | Generate BM25 sparse vector inside PostgreSQL |
| {{< ext "vchord_bm25" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | A postgresql extension for bm25 ranking algorithm |
| {{< ext "pg_tokenizer" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | Tokenizers for full-text search |
| {{< ext "hunspell_cs_cz" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Czech Hunspell Dictionary |
| {{< ext "hunspell_de_de" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | German Hunspell Dictionary |
| {{< ext "hunspell_en_us" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | en_US Hunspell Dictionary |
| {{< ext "hunspell_fr" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | French Hunspell Dictionary |
| {{< ext "hunspell_ne_np" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Nepali Hunspell Dictionary |
| {{< ext "hunspell_nl_nl" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Dutch Hunspell Dictionary |
| {{< ext "hunspell_nn_no" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Norwegian (norsk) Hunspell Dictionary |
| {{< ext "hunspell_pt_pt" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Portuguese Hunspell Dictionary |
| {{< ext "hunspell_ru_ru" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Russian Hunspell Dictionary |
| {{< ext "hunspell_ru_ru_aot" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Russian Hunspell Dictionary (from AOT.ru group) |
| {{< ext "fuzzystrmatch" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | determine similarities and distance between strings |
| {{< ext "pg_trgm" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | text similarity measurement and index searching based on trigrams |
| {{< ext "citus" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | Distributed PostgreSQL as an extension |
| {{< ext "citus_columnar" "citus" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | Citus columnar storage engine |
| {{< ext "columnar" "hydra" >}} | {{< pgver "18,17,16,15,14" "r,r,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Hydra Columnar extension |
| {{< ext "pg_analytics" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | Postgres for analytics, powered by DuckDB |
| {{< ext "pg_duckdb" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | DuckDB Embedded in Postgres |
| {{< ext "pg_mooncake" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Columnstore Table in Postgres |
| {{< ext "duckdb_fdw" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | DuckDB Foreign Data Wrapper |
| {{< ext "pg_parquet" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sLdt-" color="blue" >}} | copy data between Postgres and Parquet |
| {{< ext "pg_fkpart" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Table partitioning by foreign key utility |
| {{< ext "pg_partman" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Extension to manage partitioned tables by time or ID |
| {{< ext "plproxy" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Database partitioning implemented as procedural language |
| {{< ext "pg_strom" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PG-Strom - big-data processing acceleration using GPU and NVME |
| {{< ext "tablefunc" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | functions that manipulate whole tables, including crosstab |
| {{< ext "age" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | AGE graph database extension |
| {{< ext "hll" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | type for storing hyperloglog data |
| {{< ext "rum" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | RUM index access method |
| {{< ext "pg_graphql" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Add in-database GraphQL support |
| {{< ext "pg_jsonschema" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | PostgreSQL extension providing JSON Schema validation |
| {{< ext "jsquery" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | data type for jsonb inspection |
| {{< ext "pg_hint_plan" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Give PostgreSQL ability to manually force some decisions in execution plans. |
| {{< ext "hypopg" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Hypothetical indexes for PostgreSQL |
| {{< ext "index_advisor" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Query index advisor |
| {{< ext "plan_filter" "pg_plan_filter" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | filter statements by their execution plans. |
| {{< ext "imgsmlr" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Image similarity with haar |
| {{< ext "pg_ivm" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | incremental view maintenance on PostgreSQL |
| {{< ext "pg_incremental" >}} | {{< pgver "18,17,16,15,14" "g,g,g,r,r" >}} | {{< badge content="--s-d--" color="blue" >}} | Incremental Processing by Crunchy Data |
| {{< ext "pgmq" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | A lightweight message queue. Like AWS SQS and RSMQ but on Postgres. |
| {{< ext "pgq" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Generic queue for PostgreSQL |
| {{< ext "orioledb" >}} | {{< pgver "18,17,16,15,14" "r,g,r,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | OrioleDB, the next generation transactional engine |
| {{< ext "pg_cardano" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | A suite of Cardano-related tools |
| {{< ext "rdkit" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Cheminformatics functionality for PostgreSQL. |
| {{< ext "omni" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | Advanced adapter for Postgres extensions |
| {{< ext "omni_auth" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Basic session management |
| {{< ext "omni_aws" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | Amazon Web Services APIs (S3) |
| {{< ext "omni_cloudevents" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | CloudEvents support |
| {{< ext "omni_containers" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Docker container management |
| {{< ext "omni_credentials" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Application credential management |
| {{< ext "omni_email" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | E-mail framework |
| {{< ext "omni_http" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Basic HTTP types |
| {{< ext "omni_httpc" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | HTTP client |
| {{< ext "omni_httpd" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | HTTP server |
| {{< ext "omni_id" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Identity types |
| {{< ext "omni_json" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | JSON toolkit |
| {{< ext "omni_kube" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Kubernetes (k8s) integration |
| {{< ext "omni_ledger" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Financial ledger |
| {{< ext "omni_manifest" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Package installation manifests |
| {{< ext "omni_mimetypes" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | MIME types |
| {{< ext "omni_os" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Operating system integration |
| {{< ext "omni_polyfill" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Postgres API polyfills |
| {{< ext "omni_python" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | First-class Python support |
| {{< ext "omni_regex" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | PCRE-compatible regular expressions |
| {{< ext "omni_rest" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | REST API toolkit (with PostgREST support) |
| {{< ext "omni_schema" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Advanced schema management tooling |
| {{< ext "omni_seq" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Distributed integer sequences |
| {{< ext "omni_service" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Service management |
| {{< ext "omni_session" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Session management |
| {{< ext "omni_sql" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Programmatic SQL manipulation |
| {{< ext "omni_sqlite" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Embedded SQLite |
| {{< ext "omni_test" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Testing framework |
| {{< ext "omni_txn" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Transaction management |
| {{< ext "omni_types" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Advanced types |
| {{< ext "omni_var" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Scoped variables |
| {{< ext "omni_vfs" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Virtual File System |
| {{< ext "omni_vfs_types_v1" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Virtual File System types (v1) |
| {{< ext "omni_web" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Common web stack primitives |
| {{< ext "omni_worker" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Generalized worker pool |
| {{< ext "omni_xml" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | XML toolkit |
| {{< ext "omni_yaml" "omnigres" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | YAML toolkit |
| {{< ext "bloom" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | bloom access method - signature file based index |
| {{< ext "pg_tle" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | Trusted Language Extensions for PostgreSQL |
| {{< ext "plv8" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PL/JavaScript (v8) trusted procedural language |
| {{< ext "pllua" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Lua as a procedural language |
| {{< ext "hstore_pllua" "pllua" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Hstore transform for Lua |
| {{< ext "plluau" "pllua" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Lua as an untrusted procedural language |
| {{< ext "hstore_plluau" "pllua" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Hstore transform for untrusted Lua |
| {{< ext "plprql" >}} | {{< pgver "18,17,16,15,14" "r,r,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Use PRQL in PostgreSQL - Pipelined Relational Query Language |
| {{< ext "pldbgapi" "pldebugger" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | server-side support for debugging PL/pgSQL functions |
| {{< ext "plpgsql_check" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | extended check for plpgsql functions |
| {{< ext "plprofiler" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | server-side support for profiling PL/pgSQL functions |
| {{< ext "plsh" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | PL/sh procedural language |
| {{< ext "pljava" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PL/Java procedural language |
| {{< ext "plr" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | load R interpreter and execute R script from within a database |
| {{< ext "pgtap" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Unit testing for PostgreSQL |
| {{< ext "faker" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Wrapper for the Faker Python library |
| {{< ext "dbt2" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | OSDL-DBT-2 test kit |
| {{< ext "pltcl" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | PL/Tcl procedural language |
| {{< ext "pltclu" "pltcl" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | PL/TclU untrusted procedural language |
| {{< ext "plperl" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | PL/Perl procedural language |
| {{< ext "bool_plperl" "plperl" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | transform between bool and plperl |
| {{< ext "hstore_plperl" "plperl" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | transform between hstore and plperl |
| {{< ext "jsonb_plperl" "plperl" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | transform between jsonb and plperl |
| {{< ext "plperlu" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | PL/PerlU untrusted procedural language |
| {{< ext "bool_plperlu" "plperlu" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | transform between bool and plperlu |
| {{< ext "jsonb_plperlu" "plperlu" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | transform between jsonb and plperlu |
| {{< ext "hstore_plperlu" "plperlu" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | transform between hstore and plperlu |
| {{< ext "plpgsql" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | PL/pgSQL procedural language |
| {{< ext "plpython3u" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | PL/Python3U untrusted procedural language |
| {{< ext "jsonb_plpython3u" "plpython3u" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c---d-r" color="blue" >}} | transform between jsonb and plpython3u |
| {{< ext "ltree_plpython3u" "plpython3u" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d-r" color="blue" >}} | transform between ltree and plpython3u |
| {{< ext "hstore_plpython3u" "plpython3u" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c---d-r" color="blue" >}} | transform between hstore and plpython3u |
| {{< ext "prefix" "pg_prefix" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Prefix Range module for PostgreSQL |
| {{< ext "semver" "pg_semver" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Semantic version data type |
| {{< ext "unit" "pgunit" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | SI units extension |
| {{< ext "pgpdf" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--sLdtr" color="blue" >}} | PDF type with meta admin & Full-Text Search |
| {{< ext "pglite_fusion" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | Embed an SQLite database in your PostgreSQL table |
| {{< ext "md5hash" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | type for storing 128-bit binary data inline |
| {{< ext "asn1oid" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | asn1oid extension |
| {{< ext "roaringbitmap" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | support for Roaring Bitmaps |
| {{< ext "pgfaceting" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | fast faceting queries using an inverted index |
| {{< ext "pg_sphere" "pgsphere" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | spherical objects with useful functions, operators and index support |
| {{< ext "country" "pg_country" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Country data type, ISO 3166-1 |
| {{< ext "pg_xenophile" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | More than the bare necessities for PostgreSQL i18n and l10n. |
| {{< ext "l10n_table_dependent_extension" "pg_xenophile" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | PostgreSQL l10n toolbox |
| {{< ext "currency" "pg_currency" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Custom PostgreSQL currency type in 1Byte |
| {{< ext "collection" "pg_collection" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Memory optimized data type to be used inside of plpglsql func |
| {{< ext "pgmp" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Multiple Precision Arithmetic extension |
| {{< ext "numeral" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | numeral datatypes extension |
| {{< ext "pg_rational" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | bigint fractions |
| {{< ext "uint" "pguint" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | unsigned integer types |
| {{< ext "uint128" "pg_uint128" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Native uint128 type |
| {{< ext "hashtypes" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | sha1, md5 and other data types for PostgreSQL |
| {{< ext "ip4r" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | IPv4/v6 and IPv4/v6 range index type for PostgreSQL |
| {{< ext "pg_duration" >}} | {{< pgver "18,17,16,15,14" "g,g,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | data type for representing durations |
| {{< ext "uri" "pg_uri" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | URI Data type for PostgreSQL |
| {{< ext "emailaddr" "pgemailaddr" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Email address type for PostgreSQL |
| {{< ext "acl" "pg_acl" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | ACL Data type |
| {{< ext "debversion" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Debian version number data type |
| {{< ext "pg_rrule" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | RRULE field type for PostgreSQL |
| {{< ext "timestamp9" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | timestamp nanosecond resolution |
| {{< ext "chkpass" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | data type for auto-encrypted passwords |
| {{< ext "isn" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | data types for international product numbering standards |
| {{< ext "seg" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | data type for representing line segments or floating-point intervals |
| {{< ext "cube" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | data type for multidimensional cubes |
| {{< ext "ltree" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | data type for hierarchical tree-like structures |
| {{< ext "hstore" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | data type for storing sets of (key, value) pairs |
| {{< ext "citext" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | data type for case-insensitive character strings |
| {{< ext "xml2" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | XPath querying and XSLT |
| {{< ext "gzip" "pg_gzip" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | gzip and gunzip functions. |
| {{< ext "bzip" "pg_bzip" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Bzip compression and decompression |
| {{< ext "zstd" "pg_zstd" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Zstandard compression algorithm implementation in PostgreSQL |
| {{< ext "http" "pg_http" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | HTTP client for PostgreSQL, allows web page retrieval inside the database. |
| {{< ext "pg_net" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | Async HTTP Requests |
| {{< ext "pg_curl" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Run curl actions for data transfer in URL syntax |
| {{< ext "pgjq" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | Use jq in Postgres |
| {{< ext "pgjwt" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | JSON Web Token API for Postgresql |
| {{< ext "pg_smtp_client" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | PostgreSQL extension to send email using SMTP |
| {{< ext "pg_html5_email_address" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | PostgreSQL email validation that is consistent with the HTML5 spec |
| {{< ext "url_encode" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | url_encode, url_decode functions |
| {{< ext "pgsql_tweaks" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Some functions and views for daily usage |
| {{< ext "pg_extra_time" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | Some date time functions and operators that, |
| {{< ext "pgpcre" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Perl Compatible Regular Expression functions |
| {{< ext "icu_ext" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Access ICU functions |
| {{< ext "pgqr" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | QR Code generator from PostgreSQL |
| {{< ext "pg_protobuf" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Protobuf support for PostgreSQL |
| {{< ext "envvar" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Fetch the value of an environment variable |
| {{< ext "floatfile" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Simple file storage for arrays of floats |
| {{< ext "pg_render" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Render HTML in SQL |
| {{< ext "pg_readme" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | Generate a README.md document for a database extension or schema |
| {{< ext "pg_readme_test_extension" "pg_readme" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | Test generating a README.md document for extension or schema |
| {{< ext "ddl_historization" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Historize the ddl changes inside PostgreSQL database |
| {{< ext "data_historization" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | PLPGSQL Script to historize data in partitionned table |
| {{< ext "schedoc" "pg_schedoc" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Cross documentation between Django and DBT projects |
| {{< ext "hashlib" "pg_hashlib" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | Stable hash functions for Postgres |
| {{< ext "xxhash" "pg_xxhash" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | xxhash functions for PostgreSQL |
| {{< ext "shacrypt" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Implements SHA256-CRYPT and SHA512-CRYPT password encryption schemes |
| {{< ext "cryptint" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Encryption functions for int and bigint values |
| {{< ext "pguecc" "pg_ecdsa" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | uECC bindings for Postgres |
| {{< ext "sparql" "pgsparql" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Query SPARQL datasource with SQL |
| {{< ext "pg_idkit" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | multi-tool for generating new/niche universally unique identifiers (ex. UUIDv6, ULID, KSUID) |
| {{< ext "pgx_ulid" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | ulid type and methods |
| {{< ext "pg_uuidv7" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Create UUIDv7 values in postgres |
| {{< ext "permuteseq" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Pseudo-randomly permute sequences with a format-preserving encryption on elements |
| {{< ext "pg_hashids" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Short unique id generator for PostgreSQL, using hashids |
| {{< ext "sequential_uuids" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | generator of sequential UUIDs |
| {{< ext "topn" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | type for top-n JSONB |
| {{< ext "quantile" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Quantile aggregation function |
| {{< ext "lower_quantile" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Lower quantile aggregate function |
| {{< ext "count_distinct" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | An alternative to COUNT(DISTINCT …) aggregate, usable with HashAggregate |
| {{< ext "omnisketch" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | data structure for on-line agg of data into approximate sketch |
| {{< ext "ddsketch" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Provides ddsketch aggregate function |
| {{< ext "vasco" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | discover hidden correlations in your data with MIC |
| {{< ext "xicor" "pgxicor" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | XI Correlation Coefficient in Postgres |
| {{< ext "tdigest" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Provides tdigest aggregate function. |
| {{< ext "first_last_agg" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | first() and last() aggregate functions |
| {{< ext "extra_window_functions" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Extra Window Functions for PostgreSQL |
| {{< ext "floatvec" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Math for vectors (arrays) of numbers |
| {{< ext "aggs_for_vecs" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Aggregate functions for array inputs |
| {{< ext "aggs_for_arrays" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Various functions for computing statistics on arrays of numbers |
| {{< ext "arraymath" "pg_arraymath" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Array math and operators that work element by element on the contents of arrays |
| {{< ext "pg_math" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | GSL statistical functions for postgresql |
| {{< ext "random" "pg_random" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | random data generator |
| {{< ext "base36" "pg_base36" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Integer Base36 types |
| {{< ext "base62" "pg_base62" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Base62 extension for PostgreSQL |
| {{< ext "pg_base58" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Base58 Encoder/Decoder Extension for PostgreSQL |
| {{< ext "financial" "pg_financial" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Financial aggregate functions |
| {{< ext "convert" "pg_convert" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | conversion functions for spatial, routing and other specialized uses |
| {{< ext "refint" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | functions for implementing referential integrity (obsolete) |
| {{< ext "autoinc" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | functions for autoincrementing fields |
| {{< ext "insert_username" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | functions for tracking who changed a table |
| {{< ext "moddatetime" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | functions for tracking last modification time |
| {{< ext "tsm_system_time" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | TABLESAMPLE method which accepts time in milliseconds as a limit |
| {{< ext "dict_xsyn" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | text search dictionary template for extended synonym processing |
| {{< ext "tsm_system_rows" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | TABLESAMPLE method which accepts number of rows as a limit |
| {{< ext "tcn" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | Triggered change notifications |
| {{< ext "uuid-ossp" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | generate universally unique identifiers (UUIDs) |
| {{< ext "btree_gist" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | support for indexing common datatypes in GiST |
| {{< ext "btree_gin" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | support for indexing common datatypes in GIN |
| {{< ext "intarray" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c---dt-" color="blue" >}} | functions, operators, and index support for 1-D arrays of integers |
| {{< ext "intagg" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | integer aggregator and enumerator (obsolete) |
| {{< ext "dict_int" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | text search dictionary template for integers |
| {{< ext "unaccent" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | text search dictionary that removes accents |
| {{< ext "pg_repack" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="-bs-d--" color="blue" >}} | Reorganize tables in PostgreSQL databases with minimal locks |
| {{< ext "pg_rewrite" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | Tool allows read write to the table during the rewriting |
| {{< ext "pg_squeeze" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | A tool to remove unused space from a relation. |
| {{< ext "pg_dirtyread" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Read dead but unvacuumed rows from table |
| {{< ext "pgfincore" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | examine and manage the os buffer cache |
| {{< ext "pg_cooldown" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | remove buffered pages for specific relations |
| {{< ext "ddlx" "pg_ddlx" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | DDL eXtractor functions |
| {{< ext "prioritize" "pg_prioritize" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | get and set the priority of PostgreSQL backends |
| {{< ext "pg_checksums" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s---r" color="blue" >}} | Activate/deactivate/verify checksums in offline Postgres clusters |
| {{< ext "pg_readonly" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | cluster database read only |
| {{< ext "pgdd" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | Introspect pg data dictionary via standard SQL |
| {{< ext "pg_permissions" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | view object permissions and compare them with the desired state |
| {{< ext "pgautofailover" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | pg_auto_failover |
| {{< ext "pg_catcheck" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Diagnosing system catalog corruption |
| {{< ext "pre_prepare" "preprepare" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Pre Prepare your Statement server side |
| {{< ext "pg_upless" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | Detect Useless UPDATE |
| {{< ext "pgcozy" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Pre-warming shared buffers according to previous pg_buffercache snapshots for PostgreSQL. |
| {{< ext "pg_orphaned" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Deal with orphaned files |
| {{< ext "pg_crash" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | Send random signals to random processes |
| {{< ext "pg_cheat_funcs" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Provides cheat (but useful) functions |
| {{< ext "fio" "pg_fio" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | PostgreSQL File I/O Functions |
| {{< ext "pg_savior" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Postgres extension to save OOPS mistakes |
| {{< ext "safeupdate" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | Require criteria for UPDATE and DELETE |
| {{< ext "pg_drop_events" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | logs transaction ids of drop table, drop column, drop materialized view statements |
| {{< ext "table_log" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | record table modification logs and PITR for table/row |
| {{< ext "pgagent" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | A PostgreSQL job scheduler |
| {{< ext "pg_prewarm" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-sLd--" color="blue" >}} | prewarm relation data |
| {{< ext "pgpool_adm" "pgpool" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Administrative functions for pgPool |
| {{< ext "pgpool_recovery" "pgpool" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | recovery functions for pgpool-II for V4.3 |
| {{< ext "pgpool_regclass" "pgpool" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | replacement for regclass |
| {{< ext "lo" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | Large Object maintenance |
| {{< ext "basic_archive" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,r" >}} | {{< badge content="c-s----" color="blue" >}} | an example of an archive module |
| {{< ext "basebackup_to_shell" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,r" >}} | {{< badge content="c-s----" color="blue" >}} | adds a custom basebackup target called shell |
| {{< ext "old_snapshot" >}} | {{< pgver "18,17,16,15,14" "r,r,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | utilities in support of old_snapshot_threshold |
| {{< ext "adminpack" >}} | {{< pgver "18,17,16,15,14" "r,r,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | administrative functions for PostgreSQL |
| {{< ext "amcheck" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | functions for verifying relation integrity |
| {{< ext "pg_surgery" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | extension to perform surgery on a damaged relation |
| {{< ext "pg_profile" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostgreSQL load profile repository and report builder |
| {{< ext "pg_tracing" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | Distributed Tracing for PostgreSQL |
| {{< ext "pg_show_plans" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | show query plans of all currently running SQL statements |
| {{< ext "pg_stat_kcache" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | Kernel statistics gathering |
| {{< ext "pg_stat_monitor" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | The pg_stat_monitor is a PostgreSQL Query Performance Monitoring tool, based on PostgreSQL contrib m |
| {{< ext "pg_qualstats" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | An extension collecting statistics about quals |
| {{< ext "pg_store_plans" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | track plan statistics of all SQL statements executed |
| {{< ext "pg_track_settings" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | Track settings changes |
| {{< ext "pg_wait_sampling" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | sampling based statistics of wait events |
| {{< ext "pgsentinel" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | active session history |
| {{< ext "system_stats" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | EnterpriseDB system statistics for PostgreSQL |
| {{< ext "meta" "pg_meta" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Normalized, friendlier system catalog for PostgreSQL |
| {{< ext "pgnodemx" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Capture node OS metrics via SQL queries |
| {{< ext "pg_proctab" "pgnodemx" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | PostgreSQL extension to access the OS process table |
| {{< ext "pg_sqlog" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Provide SQL interface to logs |
| {{< ext "bgw_replstatus" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | Small PostgreSQL background worker to report whether a node is a replication master or standby |
| {{< ext "pgmeminfo" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | show memory usage |
| {{< ext "toastinfo" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | show details on toasted datums |
| {{< ext "explain_ui" "pg_explain_ui" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | easily jump into a visual plan UI for any SQL query |
| {{< ext "pg_relusage" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | Log all the queries that reference a particular column |
| {{< ext "pagevis" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Visualise database pages in ascii code |
| {{< ext "powa" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostgreSQL Workload Analyser-core |
| {{< ext "pg_overexplain" >}} | {{< pgver "18,17,16,15,14" "g,r,r,r,r" >}} | {{< badge content="c-sL---" color="blue" >}} | Allow EXPLAIN to dump even more details |
| {{< ext "pg_logicalinspect" >}} | {{< pgver "18,17,16,15,14" "g,r,r,r,r" >}} | {{< badge content="c-s-d--" color="blue" >}} | Logical decoding components inspection |
| {{< ext "pageinspect" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | inspect the contents of database pages at a low level |
| {{< ext "pgrowlocks" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | show row-level locking information |
| {{< ext "sslinfo" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | information about SSL certificates |
| {{< ext "pg_buffercache" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | examine the shared buffer cache |
| {{< ext "pg_walinspect" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,r" >}} | {{< badge content="c-s-d--" color="blue" >}} | functions to inspect contents of PostgreSQL Write-Ahead Log |
| {{< ext "pg_freespacemap" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | examine the free space map (FSM) |
| {{< ext "pg_visibility" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | examine the visibility map (VM) and page-level visibility info |
| {{< ext "pgstattuple" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | show tuple-level statistics |
| {{< ext "auto_explain" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | Provides a means for logging execution plans of slow statements automatically |
| {{< ext "pg_stat_statements" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-sLd--" color="blue" >}} | track planning and execution statistics of all SQL statements executed |
| {{< ext "passwordcheck_cracklib" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | Strengthen PostgreSQL user password checks with cracklib |
| {{< ext "supautils" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | Extension that secures a cluster on a cloud environment |
| {{< ext "pgsodium" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | Postgres extension for libsodium functions |
| {{< ext "supabase_vault" "pg_vault" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Supabase Vault Extension |
| {{< ext "pg_session_jwt" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | Manage authentication sessions using JWTs |
| {{< ext "anon" "pg_anon" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | PostgreSQL Anonymizer (anon) extension |
| {{< ext "pg_tde" >}} | {{< pgver "18,17,16,15,14" "r,g,r,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | Percona pg_tde access method |
| {{< ext "pgsmcrypto" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostgreSQL SM Algorithm Extension |
| {{< ext "pgaudit" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | provides auditing functionality |
| {{< ext "pgauditlogtofile" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | pgAudit addon to redirect audit log to an independent file |
| {{< ext "pg_auth_mon" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | monitor connection attempts per user |
| {{< ext "credcheck" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | credcheck - postgresql plain text credential checker |
| {{< ext "pgcryptokey" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | cryptographic key management |
| {{< ext "pg_jobmon" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Extension for logging and monitoring functions in PostgreSQL |
| {{< ext "logerrors" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Function for collecting statistics about messages in logfile |
| {{< ext "login_hook" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | login_hook - hook to execute login_hook.login() at login time |
| {{< ext "set_user" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | similar to SET ROLE but with added logging |
| {{< ext "pg_snakeoil" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--sL--r" color="blue" >}} | The PostgreSQL Antivirus |
| {{< ext "pgextwlist" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | PostgreSQL Extension Whitelisting |
| {{< ext "pg_auditor" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Audit data changes and provide flashback ability |
| {{< ext "sslutils" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | A Postgres extension for managing SSL certificates through SQL |
| {{< ext "noset" "pg_noset" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | Module for blocking SET variables for non-super users. |
| {{< ext "sepgsql" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | label-based mandatory access control (MAC) based on SELinux security policy. |
| {{< ext "auth_delay" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | pause briefly before reporting authentication failure |
| {{< ext "pgcrypto" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | cryptographic functions |
| {{< ext "passwordcheck" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | checks user passwords and reject weak password |
| {{< ext "wrappers" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | Foreign data wrappers developed by Supabase |
| {{< ext "multicorn" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Fetch foreign data in Python in your PostgreSQL server. |
| {{< ext "odbc_fdw" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Foreign data wrapper for accessing remote databases using ODBC |
| {{< ext "jdbc_fdw" >}} | {{< pgver "18,17,16,15,14" "r,r,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | foreign-data wrapper for remote servers available over JDBC |
| {{< ext "pgspider_ext" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | foreign-data wrapper for remote PGSpider servers |
| {{< ext "mysql_fdw" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Foreign data wrapper for querying a MySQL server |
| {{< ext "oracle_fdw" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | foreign data wrapper for Oracle access |
| {{< ext "tds_fdw" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Foreign data wrapper for querying a TDS database (Sybase or Microsoft SQL Server) |
| {{< ext "db2_fdw" >}} | {{< pgver "18,17,16,15,14" "r,r,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | foreign data wrapper for DB2 access |
| {{< ext "sqlite_fdw" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | SQLite Foreign Data Wrapper |
| {{< ext "pgbouncer_fdw" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Extension for querying PgBouncer stats from normal SQL views & running pgbouncer commands from norma |
| {{< ext "mongo_fdw" >}} | {{< pgver "18,17,16,15,14" "r,r,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | foreign data wrapper for MongoDB access |
| {{< ext "redis_fdw" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Foreign data wrapper for querying a Redis server |
| {{< ext "redis" "pg_redis_pubsub" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Send redis pub/sub messages to Redis from PostgreSQL Directly |
| {{< ext "kafka_fdw" >}} | {{< pgver "18,17,16,15,14" "r,r,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | kafka Foreign Data Wrapper for CSV formatted messages |
| {{< ext "hdfs_fdw" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | foreign-data wrapper for remote hdfs servers |
| {{< ext "firebird_fdw" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Foreign data wrapper for Firebird |
| {{< ext "aws_s3" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | aws_s3 postgres extension to import/export data from/to s3 |
| {{< ext "log_fdw" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | foreign-data wrapper for Postgres log file access |
| {{< ext "dblink" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | connect to other PostgreSQL databases from within a database |
| {{< ext "file_fdw" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | foreign-data wrapper for flat file access |
| {{< ext "postgres_fdw" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | foreign-data wrapper for remote PostgreSQL servers |
| {{< ext "documentdb" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | API surface for DocumentDB for PostgreSQL |
| {{< ext "documentdb_core" "documentdb" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | Core API surface for DocumentDB for PostgreSQL |
| {{< ext "documentdb_distributed" "documentdb" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | Multi-Node API surface for DocumentDB |
| {{< ext "orafce" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Functions and operators that emulate a subset of functions and packages from the Oracle RDBMS |
| {{< ext "pgtt" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Extension to add Global Temporary Tables feature to PostgreSQL |
| {{< ext "session_variable" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Registration and manipulation of session variables and constants |
| {{< ext "pg_statement_rollback" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | Server side rollback at statement level for PostgreSQL like Oracle or DB2 |
| {{< ext "pg_dbms_metadata" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Extension to add Oracle DBMS_METADATA compatibility to PostgreSQL |
| {{< ext "pg_dbms_lock" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Extension to add Oracle DBMS_LOCK full compatibility to PostgreSQL |
| {{< ext "pg_dbms_job" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Extension to add Oracle DBMS_JOB full compatibility to PostgreSQL |
| {{< ext "babelfishpg_common" >}} | {{< pgver "18,17,16,15,14" "r,r,r,g,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | SQL Server Transact SQL Datatype Support |
| {{< ext "babelfishpg_tsql" >}} | {{< pgver "18,17,16,15,14" "r,r,r,g,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | SQL Server Transact SQL compatibility |
| {{< ext "babelfishpg_tds" >}} | {{< pgver "18,17,16,15,14" "r,r,r,g,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | SQL Server TDS protocol extension |
| {{< ext "babelfishpg_money" >}} | {{< pgver "18,17,16,15,14" "r,r,r,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | SQL Server Money Data Type |
| {{< ext "spat" >}} | {{< pgver "18,17,16,15,14" "r,g,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | Redis-like In-Memory DB Embedded in Postgres |
| {{< ext "pgmemcache" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | memcached interface |
| {{< ext "pglogical" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostgreSQL Logical Replication |
| {{< ext "pglogical_origin" "pglogical" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Dummy extension for compatibility when upgrading from Postgres 9.4 |
| {{< ext "pglogical_ticker" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | Have an accurate view on pglogical replication delay |
| {{< ext "pgl_ddl_deploy" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | automated ddl deployment using pglogical |
| {{< ext "pg_failover_slots" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--sL--r" color="blue" >}} | PG Failover Slots extension |
| {{< ext "db_migrator" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | Tools to migrate other databases to PostgreSQL |
| {{< ext "pgactive" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="-bsLd--" color="blue" >}} | Active-Active Replication Extension for PostgreSQL |
| {{< ext "wal2json" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | Changing data capture in JSON format |
| {{< ext "wal2mongo" >}} | {{< pgver "18,17,16,15,14" "r,r,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | PostgreSQL logical decoding output plugin for MongoDB |
| {{< ext "decoderbufs" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | Logical decoding plugin that delivers WAL stream changes using a Protocol Buffer format |
| {{< ext "decoder_raw" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | Output plugin for logical replication in Raw SQL format |
| {{< ext "mimeo" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Extension for specialized, per-table replication between PostgreSQL instances |
| {{< ext "repmgr" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Replication manager for PostgreSQL |
| {{< ext "pg_fact_loader" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | build fact tables with Postgres |
| {{< ext "pg_bulkload" >}} | {{< pgver "18,17,16,15,14" "r,g,g,g,g" >}} | {{< badge content="-bs-d--" color="blue" >}} | pg_bulkload is a high speed data loading utility for PostgreSQL |
| {{< ext "test_decoding" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s----" color="blue" >}} | SQL-based test/example module for WAL logical decoding |
| {{< ext "pgoutput" >}} | {{< pgver "18,17,16,15,14" "g,g,g,g,g" >}} | {{< badge content="c-s----" color="blue" >}} | Logical Replication output plugin |

**Attribute Legend**: `c`:contrib `b`:bin `s`:lib `l`:load `d`:ddl `t`:trusted `r`:relocatable