---
title: "Extensions"
weight: 10
excludeSearch: true
comments: false
---

## Statistics

| **Category** | **All** | **PGDG** | **PIGSTY** | **CONTRIB** | **MISS** | **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:---------|--------:|--------:|----------:|-----------:|--------:|--------:|--------:|--------:|--------:|--------:|--------:|
| **ALL** | 448 | 156 | 274 | 71 | 0 | 429 | 437 | 436 | 436 | 420 | 390 |
| **EL** | 442 | 150 | 274 | 71 | 6 | 419 | 430 | 430 | 430 | 414 | 384 |
| **Debian** | 431 | 108 | 252 | 71 | 17 | 412 | 421 | 419 | 419 | 403 | 374 |

## Categories

| **Category** | **Count** | **Extensions** |
|:------------:|:---------:|:---------------|
| {{< category time >}} | 11 | {{< ext "timescaledb" >}} {{< ext "timescaledb_toolkit" >}} {{< ext "timeseries" >}} {{< ext "periods" >}} {{< ext "temporal_tables" >}} {{< ext "emaj" >}} {{< ext "table_version" >}} {{< ext "pg_cron" >}} {{< ext "pg_task" >}} {{< ext "pg_later" >}} {{< ext "pg_background" >}} |
| {{< category gis >}} | 21 | {{< ext "postgis" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} {{< ext "address_standardizer_data_us" >}} {{< ext "pgrouting" >}} {{< ext "pointcloud" >}} {{< ext "pointcloud_postgis" >}} {{< ext "h3" >}} {{< ext "h3_postgis" >}} {{< ext "q3c" >}} {{< ext "ogr_fdw" >}} {{< ext "geoip" >}} {{< ext "pg_polyline" >}} {{< ext "pg_geohash" >}} {{< ext "mobilitydb" >}} {{< ext "mobilitydb_datagen" >}} {{< ext "tzf" >}} {{< ext "earthdistance" >}} |
| {{< category rag >}} | 10 | {{< ext "vector" >}} {{< ext "vchord" >}} {{< ext "vectorscale" >}} {{< ext "vectorize" >}} {{< ext "pg_similarity" >}} {{< ext "smlar" >}} {{< ext "pg_summarize" >}} {{< ext "pg_tiktoken" >}} {{< ext "pg4ml" >}} {{< ext "pgml" >}} |
| {{< category fts >}} | 22 | {{< ext "pg_search" >}} {{< ext "pgroonga" >}} {{< ext "pgroonga_database" >}} {{< ext "pg_bigm" >}} {{< ext "zhparser" >}} {{< ext "pg_bestmatch" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_tokenizer" >}} {{< ext "biscuit" >}} {{< ext "pg_textsearch" >}} {{< ext "hunspell_cs_cz" >}} {{< ext "hunspell_de_de" >}} {{< ext "hunspell_en_us" >}} {{< ext "hunspell_fr" >}} {{< ext "hunspell_ne_np" >}} {{< ext "hunspell_nl_nl" >}} {{< ext "hunspell_nn_no" >}} {{< ext "hunspell_pt_pt" >}} {{< ext "hunspell_ru_ru" >}} {{< ext "hunspell_ru_ru_aot" >}} {{< ext "fuzzystrmatch" >}} {{< ext "pg_trgm" >}} |
| {{< category olap >}} | 14 | {{< ext "citus" >}} {{< ext "citus_columnar" >}} {{< ext "columnar" >}} {{< ext "pg_analytics" >}} {{< ext "pg_duckdb" >}} {{< ext "pg_mooncake" >}} {{< ext "pg_clickhouse" >}} {{< ext "duckdb_fdw" >}} {{< ext "pg_parquet" >}} {{< ext "pg_fkpart" >}} {{< ext "pg_partman" >}} {{< ext "plproxy" >}} {{< ext "pg_strom" >}} {{< ext "tablefunc" >}} |
| {{< category feat >}} | 61 | {{< ext "age" >}} {{< ext "hll" >}} {{< ext "rum" >}} {{< ext "pg_ai_query" >}} {{< ext "pg_ttl_index" >}} {{< ext "pg_graphql" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_hint_plan" >}} {{< ext "hypopg" >}} {{< ext "index_advisor" >}} {{< ext "plan_filter" >}} {{< ext "imgsmlr" >}} {{< ext "pg_ivm" >}} {{< ext "pg_incremental" >}} {{< ext "pgmq" >}} {{< ext "pgq" >}} {{< ext "orioledb" >}} {{< ext "pg_cardano" >}} {{< ext "rdkit" >}} {{< ext "omni" >}} {{< ext "omni_auth" >}} {{< ext "omni_aws" >}} {{< ext "omni_cloudevents" >}} {{< ext "omni_containers" >}} {{< ext "omni_credentials" >}} {{< ext "omni_csv" >}} {{< ext "omni_datasets" >}} {{< ext "omni_email" >}} {{< ext "omni_http" >}} {{< ext "omni_httpc" >}} {{< ext "omni_httpd" >}} {{< ext "omni_id" >}} {{< ext "omni_json" >}} {{< ext "omni_kube" >}} {{< ext "omni_ledger" >}} {{< ext "omni_manifest" >}} {{< ext "omni_mimetypes" >}} {{< ext "omni_os" >}} {{< ext "omni_polyfill" >}} {{< ext "omni_python" >}} {{< ext "omni_regex" >}} {{< ext "omni_rest" >}} {{< ext "omni_schema" >}} {{< ext "omni_seq" >}} {{< ext "omni_service" >}} {{< ext "omni_session" >}} {{< ext "omni_shmem" >}} {{< ext "omni_sql" >}} {{< ext "omni_sqlite" >}} {{< ext "omni_test" >}} {{< ext "omni_txn" >}} {{< ext "omni_types" >}} {{< ext "omni_var" >}} {{< ext "omni_vfs" >}} {{< ext "omni_vfs_types_v1" >}} {{< ext "omni_web" >}} {{< ext "omni_worker" >}} {{< ext "omni_xml" >}} {{< ext "omni_yaml" >}} {{< ext "bloom" >}} |
| {{< category lang >}} | 33 | {{< ext "pg_tle" >}} {{< ext "plv8" >}} {{< ext "pljs" >}} {{< ext "pllua" >}} {{< ext "hstore_pllua" >}} {{< ext "plluau" >}} {{< ext "hstore_plluau" >}} {{< ext "plprql" >}} {{< ext "pldbgapi" >}} {{< ext "plpgsql_check" >}} {{< ext "plprofiler" >}} {{< ext "plsh" >}} {{< ext "pljava" >}} {{< ext "plr" >}} {{< ext "plxslt" >}} {{< ext "pgtap" >}} {{< ext "faker" >}} {{< ext "dbt2" >}} {{< ext "pltcl" >}} {{< ext "pltclu" >}} {{< ext "plperl" >}} {{< ext "bool_plperl" >}} {{< ext "hstore_plperl" >}} {{< ext "jsonb_plperl" >}} {{< ext "plperlu" >}} {{< ext "bool_plperlu" >}} {{< ext "jsonb_plperlu" >}} {{< ext "hstore_plperlu" >}} {{< ext "plpgsql" >}} {{< ext "plpython3u" >}} {{< ext "jsonb_plpython3u" >}} {{< ext "ltree_plpython3u" >}} {{< ext "hstore_plpython3u" >}} |
| {{< category type >}} | 37 | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} {{< ext "pgfaceting" >}} {{< ext "pg_sphere" >}} {{< ext "country" >}} {{< ext "pg_xenophile" >}} {{< ext "l10n_table_dependent_extension" >}} {{< ext "currency" >}} {{< ext "collection" >}} {{< ext "pgmp" >}} {{< ext "numeral" >}} {{< ext "pg_rational" >}} {{< ext "uint" >}} {{< ext "uint128" >}} {{< ext "hashtypes" >}} {{< ext "ip4r" >}} {{< ext "pg_duration" >}} {{< ext "uri" >}} {{< ext "emailaddr" >}} {{< ext "acl" >}} {{< ext "debversion" >}} {{< ext "pg_rrule" >}} {{< ext "timestamp9" >}} {{< ext "chkpass" >}} {{< ext "isn" >}} {{< ext "seg" >}} {{< ext "cube" >}} {{< ext "ltree" >}} {{< ext "hstore" >}} {{< ext "citext" >}} {{< ext "xml2" >}} |
| {{< category util >}} | 32 | {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "pg_curl" >}} {{< ext "pg_retry" >}} {{< ext "pgjq" >}} {{< ext "pgjwt" >}} {{< ext "pg_smtp_client" >}} {{< ext "pg_html5_email_address" >}} {{< ext "url_encode" >}} {{< ext "pgsql_tweaks" >}} {{< ext "pg_extra_time" >}} {{< ext "pgpcre" >}} {{< ext "icu_ext" >}} {{< ext "pgqr" >}} {{< ext "pg_protobuf" >}} {{< ext "envvar" >}} {{< ext "floatfile" >}} {{< ext "pg_render" >}} {{< ext "pg_readme" >}} {{< ext "pg_readme_test_extension" >}} {{< ext "ddl_historization" >}} {{< ext "data_historization" >}} {{< ext "schedoc" >}} {{< ext "hashlib" >}} {{< ext "xxhash" >}} {{< ext "shacrypt" >}} {{< ext "cryptint" >}} {{< ext "pguecc" >}} {{< ext "sparql" >}} |
| {{< category func >}} | 46 | {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "permuteseq" >}} {{< ext "pg_hashids" >}} {{< ext "sequential_uuids" >}} {{< ext "typeid" >}} {{< ext "topn" >}} {{< ext "quantile" >}} {{< ext "lower_quantile" >}} {{< ext "count_distinct" >}} {{< ext "omnisketch" >}} {{< ext "ddsketch" >}} {{< ext "vasco" >}} {{< ext "xicor" >}} {{< ext "weighted_statistics" >}} {{< ext "tdigest" >}} {{< ext "first_last_agg" >}} {{< ext "extra_window_functions" >}} {{< ext "floatvec" >}} {{< ext "aggs_for_vecs" >}} {{< ext "aggs_for_arrays" >}} {{< ext "pg_csv" >}} {{< ext "arraymath" >}} {{< ext "pg_math" >}} {{< ext "random" >}} {{< ext "base36" >}} {{< ext "base62" >}} {{< ext "pg_base58" >}} {{< ext "financial" >}} {{< ext "convert" >}} {{< ext "refint" >}} {{< ext "autoinc" >}} {{< ext "insert_username" >}} {{< ext "moddatetime" >}} {{< ext "tsm_system_time" >}} {{< ext "dict_xsyn" >}} {{< ext "tsm_system_rows" >}} {{< ext "tcn" >}} {{< ext "uuid-ossp" >}} {{< ext "btree_gist" >}} {{< ext "btree_gin" >}} {{< ext "intarray" >}} {{< ext "intagg" >}} {{< ext "dict_int" >}} {{< ext "unaccent" >}} |
| {{< category admin >}} | 38 | {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} {{< ext "pg_squeeze" >}} {{< ext "pg_dirtyread" >}} {{< ext "pgfincore" >}} {{< ext "pg_cooldown" >}} {{< ext "ddlx" >}} {{< ext "pglinter" >}} {{< ext "prioritize" >}} {{< ext "pg_checksums" >}} {{< ext "pg_readonly" >}} {{< ext "pgdd" >}} {{< ext "pg_permissions" >}} {{< ext "pgautofailover" >}} {{< ext "pg_catcheck" >}} {{< ext "pre_prepare" >}} {{< ext "pg_upless" >}} {{< ext "pgcozy" >}} {{< ext "pg_orphaned" >}} {{< ext "pg_crash" >}} {{< ext "pg_cheat_funcs" >}} {{< ext "fio" >}} {{< ext "pg_savior" >}} {{< ext "safeupdate" >}} {{< ext "pg_drop_events" >}} {{< ext "table_log" >}} {{< ext "pgagent" >}} {{< ext "pg_prewarm" >}} {{< ext "pgpool_adm" >}} {{< ext "pgpool_recovery" >}} {{< ext "pgpool_regclass" >}} {{< ext "lo" >}} {{< ext "basic_archive" >}} {{< ext "basebackup_to_shell" >}} {{< ext "old_snapshot" >}} {{< ext "adminpack" >}} {{< ext "amcheck" >}} {{< ext "pg_surgery" >}} |
| {{< category stat >}} | 34 | {{< ext "pg_profile" >}} {{< ext "pg_tracing" >}} {{< ext "pg_show_plans" >}} {{< ext "pg_stat_kcache" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_store_plans" >}} {{< ext "pg_track_settings" >}} {{< ext "pg_wait_sampling" >}} {{< ext "pgsentinel" >}} {{< ext "system_stats" >}} {{< ext "meta" >}} {{< ext "pgnodemx" >}} {{< ext "pg_proctab" >}} {{< ext "pg_sqlog" >}} {{< ext "bgw_replstatus" >}} {{< ext "pgmeminfo" >}} {{< ext "toastinfo" >}} {{< ext "explain_ui" >}} {{< ext "pg_relusage" >}} {{< ext "pagevis" >}} {{< ext "powa" >}} {{< ext "pg_overexplain" >}} {{< ext "pg_logicalinspect" >}} {{< ext "pageinspect" >}} {{< ext "pgrowlocks" >}} {{< ext "sslinfo" >}} {{< ext "pg_buffercache" >}} {{< ext "pg_walinspect" >}} {{< ext "pg_freespacemap" >}} {{< ext "pg_visibility" >}} {{< ext "pgstattuple" >}} {{< ext "auto_explain" >}} {{< ext "pg_stat_statements" >}} |
| {{< category sec >}} | 28 | {{< ext "passwordcheck_cracklib" >}} {{< ext "supautils" >}} {{< ext "pgsodium" >}} {{< ext "supabase_vault" >}} {{< ext "pg_session_jwt" >}} {{< ext "anon" >}} {{< ext "pgsmcrypto" >}} {{< ext "pg_enigma" >}} {{< ext "pgaudit" >}} {{< ext "pgauditlogtofile" >}} {{< ext "pg_auditor" >}} {{< ext "logerrors" >}} {{< ext "pg_auth_mon" >}} {{< ext "pg_jobmon" >}} {{< ext "credcheck" >}} {{< ext "pgcryptokey" >}} {{< ext "pg_pwhash" >}} {{< ext "login_hook" >}} {{< ext "set_user" >}} {{< ext "pg_snakeoil" >}} {{< ext "pgextwlist" >}} {{< ext "sslutils" >}} {{< ext "noset" >}} {{< ext "pg_tde" >}} {{< ext "sepgsql" >}} {{< ext "auth_delay" >}} {{< ext "pgcrypto" >}} {{< ext "passwordcheck" >}} |
| {{< category fdw >}} | 25 | {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "odbc_fdw" >}} {{< ext "jdbc_fdw" >}} {{< ext "pgspider_ext" >}} {{< ext "mysql_fdw" >}} {{< ext "oracle_fdw" >}} {{< ext "tds_fdw" >}} {{< ext "db2_fdw" >}} {{< ext "sqlite_fdw" >}} {{< ext "pgbouncer_fdw" >}} {{< ext "etcd_fdw" >}} {{< ext "informix_fdw" >}} {{< ext "nominatim_fdw" >}} {{< ext "mongo_fdw" >}} {{< ext "redis_fdw" >}} {{< ext "redis" >}} {{< ext "kafka_fdw" >}} {{< ext "hdfs_fdw" >}} {{< ext "firebird_fdw" >}} {{< ext "aws_s3" >}} {{< ext "log_fdw" >}} {{< ext "dblink" >}} {{< ext "file_fdw" >}} {{< ext "postgres_fdw" >}} |
| {{< category sim >}} | 19 | {{< ext "documentdb" >}} {{< ext "documentdb_core" >}} {{< ext "documentdb_distributed" >}} {{< ext "documentdb_extended_rum" >}} {{< ext "orafce" >}} {{< ext "pgtt" >}} {{< ext "session_variable" >}} {{< ext "pg_statement_rollback" >}} {{< ext "pg_dbms_metadata" >}} {{< ext "pg_dbms_lock" >}} {{< ext "pg_dbms_job" >}} {{< ext "pg_dbms_errlog" >}} {{< ext "pg_utl_smtp" >}} {{< ext "babelfishpg_common" >}} {{< ext "babelfishpg_tsql" >}} {{< ext "babelfishpg_tds" >}} {{< ext "babelfishpg_money" >}} {{< ext "spat" >}} {{< ext "pgmemcache" >}} |
| {{< category etl >}} | 17 | {{< ext "pglogical" >}} {{< ext "pglogical_origin" >}} {{< ext "pglogical_ticker" >}} {{< ext "pgl_ddl_deploy" >}} {{< ext "pg_failover_slots" >}} {{< ext "db_migrator" >}} {{< ext "pgactive" >}} {{< ext "wal2json" >}} {{< ext "wal2mongo" >}} {{< ext "decoderbufs" >}} {{< ext "decoder_raw" >}} {{< ext "mimeo" >}} {{< ext "repmgr" >}} {{< ext "pg_fact_loader" >}} {{< ext "pg_bulkload" >}} {{< ext "test_decoding" >}} {{< ext "pgoutput" >}} |

## Extensions

There are 448 available PostgreSQL extensions:

| Extension | PG Versions | Attribute | Description |
|:----------|:------------|:---------:|:--------------|
| {{< ext "timescaledb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | Enables scalable inserts and complex queries for time-series data |
| {{< ext "timescaledb_toolkit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | Library of analytical hyperfunctions, time-series pipelining, and other SQL utilities |
| {{< ext "timeseries" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Convenience API for time series stack |
| {{< ext "periods" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Provide Standard SQL functionality for PERIODs and SYSTEM VERSIONING |
| {{< ext "temporal_tables" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | temporal tables |
| {{< ext "emaj" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Enables fine-grained write logging and time travel on subsets of the database. |
| {{< ext "table_version" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostgreSQL table versioning extension |
| {{< ext "pg_cron" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | Job scheduler for PostgreSQL |
| {{< ext "pg_task" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | execute any sql command at any specific time at background |
| {{< ext "pg_later" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Run queries now and get results later |
| {{< ext "pg_background" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Run SQL queries in the background |
| {{< ext "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostGIS geometry and geography spatial types and functions |
| {{< ext "postgis_topology" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostGIS topology spatial types and functions |
| {{< ext "postgis_raster" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostGIS raster types and functions |
| {{< ext "postgis_sfcgal" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | PostGIS SFCGAL functions |
| {{< ext "postgis_tiger_geocoder" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | PostGIS tiger geocoder and reverse geocoder |
| {{< ext "address_standardizer" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Used to parse an address into constituent elements. Generally used to support geocoding address normalization step. |
| {{< ext "address_standardizer_data_us" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Address Standardizer US dataset example |
| {{< ext "pgrouting" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | pgRouting Extension |
| {{< ext "pointcloud" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | data type for lidar point clouds |
| {{< ext "pointcloud_postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | integration for pointcloud LIDAR data and PostGIS geometry data |
| {{< ext "h3" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | H3 bindings for PostgreSQL |
| {{< ext "h3_postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | H3 PostGIS integration |
| {{< ext "q3c" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | q3c sky indexing plugin |
| {{< ext "ogr_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | foreign-data wrapper for GIS data access |
| {{< ext "geoip" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | IP-based geolocation query |
| {{< ext "pg_polyline" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | Fast Google Encoded Polyline encoding & decoding for postgres |
| {{< ext "pg_geohash" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Handle geohash based functionality for spatial coordinates |
| {{< ext "mobilitydb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | MobilityDB geospatial trajectory data management & analysis platform |
| {{< ext "mobilitydb_datagen" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | MobilityDB random data generator functions |
| {{< ext "tzf" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | Fast lookup timezone name by GPS coordinates |
| {{< ext "earthdistance" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | calculate great-circle distances on the surface of the Earth |
| {{< ext "vector" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | vector data type and ivfflat and hnsw access methods |
| {{< ext "vchord" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | Vector database plugin for Postgres, written in Rust |
| {{< ext "vectorscale" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | Advanced indexing for vector data with DiskANN |
| {{< ext "vectorize" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | The simplest way to do vector search on Postgres |
| {{< ext "pg_similarity" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | support similarity queries |
| {{< ext "smlar" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Effective similarity search |
| {{< ext "pg_summarize" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Text Summarization using LLMs. Built using pgrx |
| {{< ext "pg_tiktoken" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | tiktoken tokenizer for use with OpenAI models in postgres |
| {{< ext "pg4ml" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | Machine learning framework for PostgreSQL |
| {{< ext "pgml" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | Run AL/ML workloads with SQL interface |
| {{< ext "pg_search" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | Full text search for PostgreSQL using BM25 |
| {{< ext "pgroonga" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | Use Groonga as index, fast full text search platform for all languages! |
| {{< ext "pgroonga_database" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | PGroonga database management module |
| {{< ext "pg_bigm" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | create 2-gram (bigram) index for faster full text search. |
| {{< ext "zhparser" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | a parser for full-text search of Chinese |
| {{< ext "pg_bestmatch" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | Generate BM25 sparse vector inside PostgreSQL |
| {{< ext "vchord_bm25" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | A postgresql extension for bm25 ranking algorithm |
| {{< ext "pg_tokenizer" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | Tokenizers for full-text search |
| {{< ext "biscuit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,r,r,r" >}} | {{< badge content="--s-d--" color="blue" >}} | IAM-LIKE pattern matching with bitmap indexing |
| {{< ext "pg_textsearch" >}} | {{< pgver "18,17,16,15,14,13" "g,g,r,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | Full-text search with BM25 ranking |
| {{< ext "hunspell_cs_cz" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Czech Hunspell Dictionary |
| {{< ext "hunspell_de_de" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | German Hunspell Dictionary |
| {{< ext "hunspell_en_us" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | en_US Hunspell Dictionary |
| {{< ext "hunspell_fr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | French Hunspell Dictionary |
| {{< ext "hunspell_ne_np" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Nepali Hunspell Dictionary |
| {{< ext "hunspell_nl_nl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Dutch Hunspell Dictionary |
| {{< ext "hunspell_nn_no" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Norwegian (norsk) Hunspell Dictionary |
| {{< ext "hunspell_pt_pt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Portuguese Hunspell Dictionary |
| {{< ext "hunspell_ru_ru" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Russian Hunspell Dictionary |
| {{< ext "hunspell_ru_ru_aot" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Russian Hunspell Dictionary (from AOT.ru group) |
| {{< ext "fuzzystrmatch" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | determine similarities and distance between strings |
| {{< ext "pg_trgm" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | text similarity measurement and index searching based on trigrams |
| {{< ext "citus" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,r,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | Distributed PostgreSQL as an extension |
| {{< ext "citus_columnar" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,r,r,r" >}} | {{< badge content="--s-d--" color="blue" >}} | Citus columnar storage engine |
| {{< ext "columnar" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Hydra Columnar extension |
| {{< ext "pg_analytics" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | Postgres for analytics, powered by DuckDB |
| {{< ext "pg_duckdb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | DuckDB Embedded in Postgres |
| {{< ext "pg_mooncake" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="---Ld--" color="blue" >}} | Columnstore Table in Postgres |
| {{< ext "pg_clickhouse" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Interfaces to query ClickHouse databases from PostgreSQL |
| {{< ext "duckdb_fdw" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | DuckDB Foreign Data Wrapper |
| {{< ext "pg_parquet" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLdt-" color="blue" >}} | copy data between Postgres and Parquet |
| {{< ext "pg_fkpart" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Table partitioning by foreign key utility |
| {{< ext "pg_partman" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Extension to manage partitioned tables by time or ID |
| {{< ext "plproxy" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Database partitioning implemented as procedural language |
| {{< ext "pg_strom" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PG-Strom - big-data processing acceleration using GPU and NVME |
| {{< ext "tablefunc" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | functions that manipulate whole tables, including crosstab |
| {{< ext "age" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | AGE graph database extension |
| {{< ext "hll" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | type for storing hyperloglog data |
| {{< ext "rum" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | RUM index access method |
| {{< ext "pg_ai_query" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | AI-powered SQL query generation for PostgreSQL |
| {{< ext "pg_ttl_index" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | Automatic data expiration with TTL indexes |
| {{< ext "pg_graphql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | Add in-database GraphQL support |
| {{< ext "pg_jsonschema" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | PostgreSQL extension providing JSON Schema validation |
| {{< ext "jsquery" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | data type for jsonb inspection |
| {{< ext "pg_hint_plan" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Give PostgreSQL ability to manually force some decisions in execution plans. |
| {{< ext "hypopg" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Hypothetical indexes for PostgreSQL |
| {{< ext "index_advisor" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Query index advisor |
| {{< ext "plan_filter" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | filter statements by their execution plans. |
| {{< ext "imgsmlr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Image similarity with haar |
| {{< ext "pg_ivm" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | incremental view maintenance on PostgreSQL |
| {{< ext "pg_incremental" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,r,r,r" >}} | {{< badge content="--s-d--" color="blue" >}} | Incremental Processing by Crunchy Data |
| {{< ext "pgmq" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | A lightweight message queue. Like AWS SQS and RSMQ but on Postgres. |
| {{< ext "pgq" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Generic queue for PostgreSQL |
| {{< ext "orioledb" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | OrioleDB, the next generation transactional engine |
| {{< ext "pg_cardano" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | A suite of Cardano-related tools |
| {{< ext "rdkit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Cheminformatics functionality for PostgreSQL. |
| {{< ext "omni" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | Advanced adapter for Postgres extensions |
| {{< ext "omni_auth" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Basic session management |
| {{< ext "omni_aws" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | Amazon Web Services APIs (S3) |
| {{< ext "omni_cloudevents" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | CloudEvents support |
| {{< ext "omni_containers" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Docker container management |
| {{< ext "omni_credentials" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Application credential management |
| {{< ext "omni_csv" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | CSV toolkit |
| {{< ext "omni_datasets" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Dataset provisioning |
| {{< ext "omni_email" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | E-mail framework |
| {{< ext "omni_http" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Basic HTTP types |
| {{< ext "omni_httpc" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | HTTP client |
| {{< ext "omni_httpd" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | HTTP server |
| {{< ext "omni_id" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Identity types |
| {{< ext "omni_json" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | JSON toolkit |
| {{< ext "omni_kube" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Kubernetes (k8s) integration |
| {{< ext "omni_ledger" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Financial ledger |
| {{< ext "omni_manifest" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Package installation manifests |
| {{< ext "omni_mimetypes" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | MIME types |
| {{< ext "omni_os" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Operating system integration |
| {{< ext "omni_polyfill" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Postgres API polyfills |
| {{< ext "omni_python" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | First-class Python support |
| {{< ext "omni_regex" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | PCRE-compatible regular expressions |
| {{< ext "omni_rest" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | REST API toolkit (with PostgREST support) |
| {{< ext "omni_schema" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Advanced schema management tooling |
| {{< ext "omni_seq" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Distributed integer sequences |
| {{< ext "omni_service" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Service management |
| {{< ext "omni_session" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Session management |
| {{< ext "omni_shmem" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Shared Memory Management |
| {{< ext "omni_sql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Programmatic SQL manipulation |
| {{< ext "omni_sqlite" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Embedded SQLite |
| {{< ext "omni_test" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Testing framework |
| {{< ext "omni_txn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Transaction management |
| {{< ext "omni_types" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Advanced types |
| {{< ext "omni_var" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Scoped variables |
| {{< ext "omni_vfs" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Virtual File System |
| {{< ext "omni_vfs_types_v1" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Virtual File System types (v1) |
| {{< ext "omni_web" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Common web stack primitives |
| {{< ext "omni_worker" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Generalized worker pool |
| {{< ext "omni_xml" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | XML toolkit |
| {{< ext "omni_yaml" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | YAML toolkit |
| {{< ext "bloom" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | bloom access method - signature file based index |
| {{< ext "pg_tle" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | Trusted Language Extensions for PostgreSQL |
| {{< ext "plv8" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PL/JavaScript (v8) trusted procedural language |
| {{< ext "pljs" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | PL/JS trusted procedural language |
| {{< ext "pllua" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Lua as a procedural language |
| {{< ext "hstore_pllua" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Hstore transform for Lua |
| {{< ext "plluau" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Lua as an untrusted procedural language |
| {{< ext "hstore_plluau" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Hstore transform for untrusted Lua |
| {{< ext "plprql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Use PRQL in PostgreSQL - Pipelined Relational Query Language |
| {{< ext "pldbgapi" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | server-side support for debugging PL/pgSQL functions |
| {{< ext "plpgsql_check" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | extended check for plpgsql functions |
| {{< ext "plprofiler" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | server-side support for profiling PL/pgSQL functions |
| {{< ext "plsh" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | PL/sh procedural language |
| {{< ext "pljava" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PL/Java procedural language |
| {{< ext "plr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | load R interpreter and execute R script from within a database |
| {{< ext "plxslt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | XSLT procedural language for PostgreSQL |
| {{< ext "pgtap" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Unit testing for PostgreSQL |
| {{< ext "faker" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Wrapper for the Faker Python library |
| {{< ext "dbt2" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | OSDL-DBT-2 test kit |
| {{< ext "pltcl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | PL/Tcl procedural language |
| {{< ext "pltclu" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | PL/TclU untrusted procedural language |
| {{< ext "plperl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | PL/Perl procedural language |
| {{< ext "bool_plperl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | transform between bool and plperl |
| {{< ext "hstore_plperl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | transform between hstore and plperl |
| {{< ext "jsonb_plperl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | transform between jsonb and plperl |
| {{< ext "plperlu" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | PL/PerlU untrusted procedural language |
| {{< ext "bool_plperlu" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | transform between bool and plperlu |
| {{< ext "jsonb_plperlu" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | transform between jsonb and plperlu |
| {{< ext "hstore_plperlu" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | transform between hstore and plperlu |
| {{< ext "plpgsql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | PL/pgSQL procedural language |
| {{< ext "plpython3u" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | PL/Python3U untrusted procedural language |
| {{< ext "jsonb_plpython3u" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d-r" color="blue" >}} | transform between jsonb and plpython3u |
| {{< ext "ltree_plpython3u" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d-r" color="blue" >}} | transform between ltree and plpython3u |
| {{< ext "hstore_plpython3u" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d-r" color="blue" >}} | transform between hstore and plpython3u |
| {{< ext "prefix" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Prefix Range module for PostgreSQL |
| {{< ext "semver" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Semantic version data type |
| {{< ext "unit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | SI units extension |
| {{< ext "pgpdf" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLdtr" color="blue" >}} | PDF type with meta admin & Full-Text Search |
| {{< ext "pglite_fusion" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | Embed an SQLite database in your PostgreSQL table |
| {{< ext "md5hash" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | type for storing 128-bit binary data inline |
| {{< ext "asn1oid" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | asn1oid extension |
| {{< ext "roaringbitmap" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | support for Roaring Bitmaps |
| {{< ext "pgfaceting" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | fast faceting queries using an inverted index |
| {{< ext "pg_sphere" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | spherical objects with useful functions, operators and index support |
| {{< ext "country" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Country data type, ISO 3166-1 |
| {{< ext "pg_xenophile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | More than the bare necessities for PostgreSQL i18n and l10n. |
| {{< ext "l10n_table_dependent_extension" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | PostgreSQL l10n toolbox |
| {{< ext "currency" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Custom PostgreSQL currency type in 1Byte |
| {{< ext "collection" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | Memory optimized data type to be used inside of plpglsql func |
| {{< ext "pgmp" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Multiple Precision Arithmetic extension |
| {{< ext "numeral" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | numeral datatypes extension |
| {{< ext "pg_rational" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | bigint fractions |
| {{< ext "uint" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | unsigned integer types |
| {{< ext "uint128" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Native uint128 type |
| {{< ext "hashtypes" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | sha1, md5 and other data types for PostgreSQL |
| {{< ext "ip4r" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | IPv4/v6 and IPv4/v6 range index type for PostgreSQL |
| {{< ext "pg_duration" >}} | {{< pgver "18,17,16,15,14,13" "g,g,r,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | data type for representing durations |
| {{< ext "uri" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | URI Data type for PostgreSQL |
| {{< ext "emailaddr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Email address type for PostgreSQL |
| {{< ext "acl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | ACL Data type |
| {{< ext "debversion" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Debian version number data type |
| {{< ext "pg_rrule" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | RRULE field type for PostgreSQL |
| {{< ext "timestamp9" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | timestamp nanosecond resolution |
| {{< ext "chkpass" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | data type for auto-encrypted passwords |
| {{< ext "isn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | data types for international product numbering standards |
| {{< ext "seg" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | data type for representing line segments or floating-point intervals |
| {{< ext "cube" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | data type for multidimensional cubes |
| {{< ext "ltree" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | data type for hierarchical tree-like structures |
| {{< ext "hstore" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | data type for storing sets of (key, value) pairs |
| {{< ext "citext" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | data type for case-insensitive character strings |
| {{< ext "xml2" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | XPath querying and XSLT |
| {{< ext "gzip" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | gzip and gunzip functions. |
| {{< ext "bzip" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Bzip compression and decompression |
| {{< ext "zstd" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Zstandard compression algorithm implementation in PostgreSQL |
| {{< ext "http" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | HTTP client for PostgreSQL, allows web page retrieval inside the database. |
| {{< ext "pg_net" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | Async HTTP Requests |
| {{< ext "pg_curl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Run curl actions for data transfer in URL syntax |
| {{< ext "pg_retry" >}} | {{< pgver "18,17,16,15,14,13" "g,g,r,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | Retry SQL statements on transient errors with exponential backoff |
| {{< ext "pgjq" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dtr" color="blue" >}} | Use jq in Postgres |
| {{< ext "pgjwt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | JSON Web Token API for Postgresql |
| {{< ext "pg_smtp_client" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | PostgreSQL extension to send email using SMTP |
| {{< ext "pg_html5_email_address" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | PostgreSQL email validation that is consistent with the HTML5 spec |
| {{< ext "url_encode" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | url_encode, url_decode functions |
| {{< ext "pgsql_tweaks" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Some functions and views for daily usage |
| {{< ext "pg_extra_time" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | Some date time functions and operators that, |
| {{< ext "pgpcre" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Perl Compatible Regular Expression functions |
| {{< ext "icu_ext" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Access ICU functions |
| {{< ext "pgqr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | QR Code generator from PostgreSQL |
| {{< ext "pg_protobuf" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Protobuf support for PostgreSQL |
| {{< ext "envvar" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Fetch the value of an environment variable |
| {{< ext "floatfile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Simple file storage for arrays of floats |
| {{< ext "pg_render" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | Render HTML in SQL |
| {{< ext "pg_readme" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | Generate a README.md document for a database extension or schema |
| {{< ext "pg_readme_test_extension" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | Test generating a README.md document for extension or schema |
| {{< ext "ddl_historization" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Historize the ddl changes inside PostgreSQL database |
| {{< ext "data_historization" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | PLPGSQL Script to historize data in partitionned table |
| {{< ext "schedoc" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Cross documentation between Django and DBT projects |
| {{< ext "hashlib" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | Stable hash functions for Postgres |
| {{< ext "xxhash" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | xxhash functions for PostgreSQL |
| {{< ext "shacrypt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Implements SHA256-CRYPT and SHA512-CRYPT password encryption schemes |
| {{< ext "cryptint" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Encryption functions for int and bigint values |
| {{< ext "pguecc" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | uECC bindings for Postgres |
| {{< ext "sparql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Query SPARQL datasource with SQL |
| {{< ext "pg_idkit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | multi-tool for generating new/niche universally unique identifiers (ex. UUIDv6, ULID, KSUID) |
| {{< ext "pgx_ulid" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | ulid type and methods |
| {{< ext "pg_uuidv7" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Create UUIDv7 values in postgres |
| {{< ext "permuteseq" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Pseudo-randomly permute sequences with a format-preserving encryption on elements |
| {{< ext "pg_hashids" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Short unique id generator for PostgreSQL, using hashids |
| {{< ext "sequential_uuids" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | generator of sequential UUIDs |
| {{< ext "typeid" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Allows to use TypeIDs in Postgres natively |
| {{< ext "topn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | type for top-n JSONB |
| {{< ext "quantile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Quantile aggregation function |
| {{< ext "lower_quantile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Lower quantile aggregate function |
| {{< ext "count_distinct" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | An alternative to COUNT(DISTINCT …) aggregate, usable with HashAggregate |
| {{< ext "omnisketch" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | data structure for on-line agg of data into approximate sketch |
| {{< ext "ddsketch" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Provides ddsketch aggregate function |
| {{< ext "vasco" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | discover hidden correlations in your data with MIC |
| {{< ext "xicor" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | XI Correlation Coefficient in Postgres |
| {{< ext "weighted_statistics" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | High-performance weighted statistics functions for sparse data |
| {{< ext "tdigest" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Provides tdigest aggregate function. |
| {{< ext "first_last_agg" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | first() and last() aggregate functions |
| {{< ext "extra_window_functions" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Extra Window Functions for PostgreSQL |
| {{< ext "floatvec" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Math for vectors (arrays) of numbers |
| {{< ext "aggs_for_vecs" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Aggregate functions for array inputs |
| {{< ext "aggs_for_arrays" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Various functions for computing statistics on arrays of numbers |
| {{< ext "pg_csv" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dtr" color="blue" >}} | Flexible CSV processing for Postgres |
| {{< ext "arraymath" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Array math and operators that work element by element on the contents of arrays |
| {{< ext "pg_math" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | GSL statistical functions for postgresql |
| {{< ext "random" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | random data generator |
| {{< ext "base36" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Integer Base36 types |
| {{< ext "base62" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Base62 extension for PostgreSQL |
| {{< ext "pg_base58" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Base58 Encoder/Decoder Extension for PostgreSQL |
| {{< ext "financial" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Financial aggregate functions |
| {{< ext "convert" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | conversion functions for spatial, routing and other specialized uses |
| {{< ext "refint" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | functions for implementing referential integrity (obsolete) |
| {{< ext "autoinc" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | functions for autoincrementing fields |
| {{< ext "insert_username" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | functions for tracking who changed a table |
| {{< ext "moddatetime" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | functions for tracking last modification time |
| {{< ext "tsm_system_time" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | TABLESAMPLE method which accepts time in milliseconds as a limit |
| {{< ext "dict_xsyn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | text search dictionary template for extended synonym processing |
| {{< ext "tsm_system_rows" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | TABLESAMPLE method which accepts number of rows as a limit |
| {{< ext "tcn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | Triggered change notifications |
| {{< ext "uuid-ossp" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | generate universally unique identifiers (UUIDs) |
| {{< ext "btree_gist" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | support for indexing common datatypes in GiST |
| {{< ext "btree_gin" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | support for indexing common datatypes in GIN |
| {{< ext "intarray" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---dt-" color="blue" >}} | functions, operators, and index support for 1-D arrays of integers |
| {{< ext "intagg" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | integer aggregator and enumerator (obsolete) |
| {{< ext "dict_int" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | text search dictionary template for integers |
| {{< ext "unaccent" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | text search dictionary that removes accents |
| {{< ext "pg_repack" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="-bs-d--" color="blue" >}} | Reorganize tables in PostgreSQL databases with minimal locks |
| {{< ext "pg_rewrite" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | Tool allows read write to the table during the rewriting |
| {{< ext "pg_squeeze" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | A tool to remove unused space from a relation. |
| {{< ext "pg_dirtyread" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Read dead but unvacuumed rows from table |
| {{< ext "pgfincore" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | examine and manage the os buffer cache |
| {{< ext "pg_cooldown" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | remove buffered pages for specific relations |
| {{< ext "ddlx" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | DDL eXtractor functions |
| {{< ext "pglinter" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostgreSQL Linting and Analysis Extension |
| {{< ext "prioritize" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | get and set the priority of PostgreSQL backends |
| {{< ext "pg_checksums" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s---r" color="blue" >}} | Activate/deactivate/verify checksums in offline Postgres clusters |
| {{< ext "pg_readonly" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | cluster database read only |
| {{< ext "pgdd" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | Introspect pg data dictionary via standard SQL |
| {{< ext "pg_permissions" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | view object permissions and compare them with the desired state |
| {{< ext "pgautofailover" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | pg_auto_failover |
| {{< ext "pg_catcheck" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Diagnosing system catalog corruption |
| {{< ext "pre_prepare" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Pre Prepare your Statement server side |
| {{< ext "pg_upless" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | Detect Useless UPDATE |
| {{< ext "pgcozy" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Pre-warming shared buffers according to previous pg_buffercache snapshots for PostgreSQL. |
| {{< ext "pg_orphaned" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Deal with orphaned files |
| {{< ext "pg_crash" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | Send random signals to random processes |
| {{< ext "pg_cheat_funcs" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Provides cheat (but useful) functions |
| {{< ext "fio" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | PostgreSQL File I/O Functions |
| {{< ext "pg_savior" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Postgres extension to save OOPS mistakes |
| {{< ext "safeupdate" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sL---" color="blue" >}} | Require criteria for UPDATE and DELETE |
| {{< ext "pg_drop_events" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | logs transaction ids of drop table, drop column, drop materialized view statements |
| {{< ext "table_log" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | record table modification logs and PITR for table/row |
| {{< ext "pgagent" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | A PostgreSQL job scheduler |
| {{< ext "pg_prewarm" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sLd--" color="blue" >}} | prewarm relation data |
| {{< ext "pgpool_adm" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Administrative functions for pgPool |
| {{< ext "pgpool_recovery" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | recovery functions for pgpool-II for V4.3 |
| {{< ext "pgpool_regclass" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | replacement for regclass |
| {{< ext "lo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | Large Object maintenance |
| {{< ext "basic_archive" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="c-s----" color="blue" >}} | an example of an archive module |
| {{< ext "basebackup_to_shell" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="c-s----" color="blue" >}} | adds a custom basebackup target called shell |
| {{< ext "old_snapshot" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,r" >}} | {{< badge content="c-s-d--" color="blue" >}} | utilities in support of old_snapshot_threshold |
| {{< ext "adminpack" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | administrative functions for PostgreSQL |
| {{< ext "amcheck" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | functions for verifying relation integrity |
| {{< ext "pg_surgery" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="c-s-d--" color="blue" >}} | extension to perform surgery on a damaged relation |
| {{< ext "pg_profile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostgreSQL load profile repository and report builder |
| {{< ext "pg_tracing" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | Distributed Tracing for PostgreSQL |
| {{< ext "pg_show_plans" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | show query plans of all currently running SQL statements |
| {{< ext "pg_stat_kcache" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | Kernel statistics gathering |
| {{< ext "pg_stat_monitor" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | The pg_stat_monitor is a PostgreSQL Query Performance Monitoring tool, based on PostgreSQL contrib module pg_stat_statements. pg_stat_monitor provides aggregated statistics, client information, plan details including plan, and histogram information. |
| {{< ext "pg_qualstats" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | An extension collecting statistics about quals |
| {{< ext "pg_store_plans" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | track plan statistics of all SQL statements executed |
| {{< ext "pg_track_settings" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | Track settings changes |
| {{< ext "pg_wait_sampling" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | sampling based statistics of wait events |
| {{< ext "pgsentinel" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | active session history |
| {{< ext "system_stats" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | EnterpriseDB system statistics for PostgreSQL |
| {{< ext "meta" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Normalized, friendlier system catalog for PostgreSQL |
| {{< ext "pgnodemx" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Capture node OS metrics via SQL queries |
| {{< ext "pg_proctab" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | PostgreSQL extension to access the OS process table |
| {{< ext "pg_sqlog" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Provide SQL interface to logs |
| {{< ext "bgw_replstatus" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | Small PostgreSQL background worker to report whether a node is a replication master or standby |
| {{< ext "pgmeminfo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | show memory usage |
| {{< ext "toastinfo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | show details on toasted datums |
| {{< ext "explain_ui" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | easily jump into a visual plan UI for any SQL query |
| {{< ext "pg_relusage" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | Log all the queries that reference a particular column |
| {{< ext "pagevis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Visualise database pages in ascii code |
| {{< ext "powa" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostgreSQL Workload Analyser-core |
| {{< ext "pg_overexplain" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="c-sL---" color="blue" >}} | Allow EXPLAIN to dump even more details |
| {{< ext "pg_logicalinspect" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="c-s-d--" color="blue" >}} | Logical decoding components inspection |
| {{< ext "pageinspect" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | inspect the contents of database pages at a low level |
| {{< ext "pgrowlocks" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | show row-level locking information |
| {{< ext "sslinfo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | information about SSL certificates |
| {{< ext "pg_buffercache" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | examine the shared buffer cache |
| {{< ext "pg_walinspect" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="c-s-d--" color="blue" >}} | functions to inspect contents of PostgreSQL Write-Ahead Log |
| {{< ext "pg_freespacemap" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | examine the free space map (FSM) |
| {{< ext "pg_visibility" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | examine the visibility map (VM) and page-level visibility info |
| {{< ext "pgstattuple" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | show tuple-level statistics |
| {{< ext "auto_explain" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | Provides a means for logging execution plans of slow statements automatically |
| {{< ext "pg_stat_statements" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sLd--" color="blue" >}} | track planning and execution statistics of all SQL statements executed |
| {{< ext "passwordcheck_cracklib" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | Strengthen PostgreSQL user password checks with cracklib |
| {{< ext "supautils" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | Extension that secures a cluster on a cloud environment |
| {{< ext "pgsodium" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | Postgres extension for libsodium functions |
| {{< ext "supabase_vault" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Supabase Vault Extension |
| {{< ext "pg_session_jwt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | Manage authentication sessions using JWTs |
| {{< ext "anon" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | PostgreSQL Anonymizer (anon) extension |
| {{< ext "pgsmcrypto" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | PostgreSQL SM Algorithm Extension |
| {{< ext "pg_enigma" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | Encrypted postgres data type |
| {{< ext "pgaudit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | provides auditing functionality |
| {{< ext "pgauditlogtofile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | pgAudit addon to redirect audit log to an independent file |
| {{< ext "pg_auditor" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | Audit data changes and provide flashback ability |
| {{< ext "logerrors" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Function for collecting statistics about messages in logfile |
| {{< ext "pg_auth_mon" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | monitor connection attempts per user |
| {{< ext "pg_jobmon" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Extension for logging and monitoring functions in PostgreSQL |
| {{< ext "credcheck" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | credcheck - postgresql plain text credential checker |
| {{< ext "pgcryptokey" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | cryptographic key management |
| {{< ext "pg_pwhash" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Advanced password hashing methods for PostgreSQL |
| {{< ext "login_hook" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | login_hook - hook to execute login_hook.login() at login time |
| {{< ext "set_user" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | similar to SET ROLE but with added logging |
| {{< ext "pg_snakeoil" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL--r" color="blue" >}} | The PostgreSQL Antivirus |
| {{< ext "pgextwlist" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | PostgreSQL Extension Whitelisting |
| {{< ext "sslutils" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | A Postgres extension for managing SSL certificates through SQL |
| {{< ext "noset" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | Module for blocking SET variables for non-super users. |
| {{< ext "pg_tde" >}} | {{< pgver "18,17,16,15,14,13" "g,g,r,r,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | Percona pg_tde access method |
| {{< ext "sepgsql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | label-based mandatory access control (MAC) based on SELinux security policy. |
| {{< ext "auth_delay" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | pause briefly before reporting authentication failure |
| {{< ext "pgcrypto" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | cryptographic functions |
| {{< ext "passwordcheck" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | checks user passwords and reject weak password |
| {{< ext "wrappers" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | Foreign data wrappers developed by Supabase |
| {{< ext "multicorn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Fetch foreign data in Python in your PostgreSQL server. |
| {{< ext "odbc_fdw" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Foreign data wrapper for accessing remote databases using ODBC |
| {{< ext "jdbc_fdw" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | foreign-data wrapper for remote servers available over JDBC |
| {{< ext "pgspider_ext" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | foreign-data wrapper for remote PGSpider servers |
| {{< ext "mysql_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Foreign data wrapper for querying a MySQL server |
| {{< ext "oracle_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | foreign data wrapper for Oracle access |
| {{< ext "tds_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Foreign data wrapper for querying a TDS database (Sybase or Microsoft SQL Server) |
| {{< ext "db2_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | foreign data wrapper for DB2 access |
| {{< ext "sqlite_fdw" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | SQLite Foreign Data Wrapper |
| {{< ext "pgbouncer_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Extension for querying PgBouncer stats from normal SQL views & running pgbouncer commands from normal SQL functions |
| {{< ext "etcd_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Foreign data wrapper for etcd |
| {{< ext "informix_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Foreign data wrapper for Informix access |
| {{< ext "nominatim_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | Nominatim Foreign Data Wrapper for PostgreSQL |
| {{< ext "mongo_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | foreign data wrapper for MongoDB access |
| {{< ext "redis_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Foreign data wrapper for querying a Redis server |
| {{< ext "redis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Send redis pub/sub messages to Redis from PostgreSQL Directly |
| {{< ext "kafka_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | kafka Foreign Data Wrapper for CSV formatted messages |
| {{< ext "hdfs_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | foreign-data wrapper for remote hdfs servers |
| {{< ext "firebird_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Foreign data wrapper for Firebird |
| {{< ext "aws_s3" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | aws_s3 postgres extension to import/export data from/to s3 |
| {{< ext "log_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | foreign-data wrapper for Postgres log file access |
| {{< ext "dblink" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | connect to other PostgreSQL databases from within a database |
| {{< ext "file_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | foreign-data wrapper for flat file access |
| {{< ext "postgres_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | foreign-data wrapper for remote PostgreSQL servers |
| {{< ext "documentdb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | API surface for DocumentDB for PostgreSQL |
| {{< ext "documentdb_core" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | Core API surface for DocumentDB for PostgreSQL |
| {{< ext "documentdb_distributed" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | Multi-Node API surface for DocumentDB |
| {{< ext "documentdb_extended_rum" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | DocumentDB Extended RUM index access method |
| {{< ext "orafce" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Functions and operators that emulate a subset of functions and packages from the Oracle RDBMS |
| {{< ext "pgtt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Extension to add Global Temporary Tables feature to PostgreSQL |
| {{< ext "session_variable" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | Registration and manipulation of session variables and constants |
| {{< ext "pg_statement_rollback" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | Server side rollback at statement level for PostgreSQL like Oracle or DB2 |
| {{< ext "pg_dbms_metadata" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Extension to add Oracle DBMS_METADATA compatibility to PostgreSQL |
| {{< ext "pg_dbms_lock" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Extension to add Oracle DBMS_LOCK full compatibility to PostgreSQL |
| {{< ext "pg_dbms_job" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Extension to add Oracle DBMS_JOB full compatibility to PostgreSQL |
| {{< ext "pg_dbms_errlog" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Emulate DBMS_ERRLOG Oracle module to log DML errors in a dedicated table. |
| {{< ext "pg_utl_smtp" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="----d--" color="blue" >}} | Oracle UTL_SMTP compatibility extension for PostgreSQL |
| {{< ext "babelfishpg_common" >}} | {{< pgver "18,17,16,15,14,13" "r,r,r,g,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | SQL Server Transact SQL Datatype Support |
| {{< ext "babelfishpg_tsql" >}} | {{< pgver "18,17,16,15,14,13" "r,r,r,g,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | SQL Server Transact SQL compatibility |
| {{< ext "babelfishpg_tds" >}} | {{< pgver "18,17,16,15,14,13" "r,r,r,g,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | SQL Server TDS protocol extension |
| {{< ext "babelfishpg_money" >}} | {{< pgver "18,17,16,15,14,13" "r,r,r,g,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | SQL Server Money Data Type |
| {{< ext "spat" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | Redis-like In-Memory DB Embedded in Postgres |
| {{< ext "pgmemcache" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | memcached interface |
| {{< ext "pglogical" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | PostgreSQL Logical Replication |
| {{< ext "pglogical_origin" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Dummy extension for compatibility when upgrading from Postgres 9.4 |
| {{< ext "pglogical_ticker" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | Have an accurate view on pglogical replication delay |
| {{< ext "pgl_ddl_deploy" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | automated ddl deployment using pglogical |
| {{< ext "pg_failover_slots" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL--r" color="blue" >}} | PG Failover Slots extension |
| {{< ext "db_migrator" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | Tools to migrate other databases to PostgreSQL |
| {{< ext "pgactive" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="-bsLd--" color="blue" >}} | Active-Active Replication Extension for PostgreSQL |
| {{< ext "wal2json" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | Changing data capture in JSON format |
| {{< ext "wal2mongo" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | PostgreSQL logical decoding output plugin for MongoDB |
| {{< ext "decoderbufs" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | Logical decoding plugin that delivers WAL stream changes using a Protocol Buffer format |
| {{< ext "decoder_raw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | Output plugin for logical replication in Raw SQL format |
| {{< ext "mimeo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | Extension for specialized, per-table replication between PostgreSQL instances |
| {{< ext "repmgr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | Replication manager for PostgreSQL |
| {{< ext "pg_fact_loader" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | build fact tables with Postgres |
| {{< ext "pg_bulkload" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="-bs-d--" color="blue" >}} | pg_bulkload is a high speed data loading utility for PostgreSQL |
| {{< ext "test_decoding" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s----" color="blue" >}} | SQL-based test/example module for WAL logical decoding |
| {{< ext "pgoutput" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s----" color="blue" >}} | Logical Replication output plugin |
