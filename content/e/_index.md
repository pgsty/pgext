---
title: "Extensions"
breadcrumbs: false
excludeSearch: true
weight: 1
---

Available PostgreSQL Extensions (424 ext in 356 pkg) categorized into 16 categories:

{{< category "time" >}}
{{< category "gis" >}}
{{< category "rag" >}}
{{< category "fts" >}}
{{< category "olap" >}}
{{< category "feat" >}}
{{< category "lang" >}}
{{< category "type" >}}
{{< category "util" >}}
{{< category "func" >}}
{{< category "admin" >}}
{{< category "stat" >}}
{{< category "sec" >}}
{{< category "fdw" >}}
{{< category "sim" >}}
{{< category "etl" >}}


## TIME

TimescaleDB, Versioning & Temporal Table, Crontab, Async & Background Job Scheduler, ...

| ID | Extension | Package | Version | Description |
|:---:|:---|:---|:---|:---|
| 1000 | {{< ext "timescaledb" >}} | {{< ext "timescaledb" "timescaledb" >}} | 2.22.0 | Enables scalable inserts and complex queries for time-series data |
| 1010 | {{< ext "timescaledb_toolkit" >}} | {{< ext "timescaledb_toolkit" "timescaledb_toolkit" >}} | 1.21.0 | Library of analytical hyperfunctions, time-series pipelining, and other SQL utilities |
| 1020 | {{< ext "timeseries" >}} | {{< ext "timeseries" "pg_timeseries" >}} | 0.1.6 | Convenience API for time series stack |
| 1030 | {{< ext "periods" >}} | {{< ext "periods" "periods" >}} | 1.2.3 | Provide Standard SQL functionality for PERIODs and SYSTEM VERSIONING |
| 1040 | {{< ext "temporal_tables" >}} | {{< ext "temporal_tables" "temporal_tables" >}} | 1.2.2 | temporal tables |
| 1050 | {{< ext "emaj" >}} | {{< ext "emaj" "emaj" >}} | 4.7.0 | Enables fine-grained write logging and time travel on subsets of the database. |
| 1060 | {{< ext "table_version" >}} | {{< ext "table_version" "table_version" >}} | 1.11.1 | PostgreSQL table versioning extension |
| 1070 | {{< ext "pg_cron" >}} | {{< ext "pg_cron" "pg_cron" >}} | 1.6.7 | Job scheduler for PostgreSQL |
| 1080 | {{< ext "pg_task" >}} | {{< ext "pg_task" "pg_task" >}} | 1.0.0 | execute any sql command at any specific time at background |
| 1090 | {{< ext "pg_later" >}} | {{< ext "pg_later" "pg_later" >}} | 0.3.0 | Run queries now and get results later |
| 1100 | {{< ext "pg_background" >}} | {{< ext "pg_background" "pg_background" >}} | 1.3 | Run SQL queries in the background |

## GIS

GeoSpatial Data Types, Operators, and Indexes, Hexagonal Indexing, OGR Data FDW, GeoIP & MobilityDB, etc...

| ID | Extension | Package | Version | Description |
|:---:|:---|:---|:---|:---|
| 1500 | {{< ext "postgis" >}} | {{< ext "postgis" "postgis" >}} | 3.6.0 | PostGIS geometry and geography spatial types and functions |
| 1501 | {{< ext "postgis_topology" >}} | {{< ext "postgis" "postgis" >}} | 3.6.0 | PostGIS topology spatial types and functions |
| 1502 | {{< ext "postgis_raster" >}} | {{< ext "postgis" "postgis" >}} | 3.6.0 | PostGIS raster types and functions |
| 1503 | {{< ext "postgis_sfcgal" >}} | {{< ext "postgis" "postgis" >}} | 3.6.0 | PostGIS SFCGAL functions |
| 1504 | {{< ext "postgis_tiger_geocoder" >}} | {{< ext "postgis" "postgis" >}} | 3.6.0 | PostGIS tiger geocoder and reverse geocoder |
| 1505 | {{< ext "address_standardizer" >}} | {{< ext "postgis" "postgis" >}} | 3.6.0 | Used to parse an address into constituent elements. Generally used to support geocoding address normalization step. |
| 1506 | {{< ext "address_standardizer_data_us" >}} | {{< ext "postgis" "postgis" >}} | 3.6.0 | Address Standardizer US dataset example |
| 1510 | {{< ext "pgrouting" >}} | {{< ext "pgrouting" "pgrouting" >}} | 3.8.0 | pgRouting Extension |
| 1520 | {{< ext "pointcloud" >}} | {{< ext "pointcloud" "pointcloud" >}} | 1.2.5 | data type for lidar point clouds |
| 1521 | {{< ext "pointcloud_postgis" >}} | {{< ext "pointcloud" "pointcloud" >}} | 1.2.5 | integration for pointcloud LIDAR data and PostGIS geometry data |
| 1530 | {{< ext "h3" >}} | {{< ext "h3" "pg_h3" >}} | 4.2.3 | H3 bindings for PostgreSQL |
| 1531 | {{< ext "h3_postgis" >}} | {{< ext "h3" "pg_h3" >}} | 4.2.3 | H3 PostGIS integration |
| 1540 | {{< ext "q3c" >}} | {{< ext "q3c" "q3c" >}} | 2.0.1 | q3c sky indexing plugin |
| 1550 | {{< ext "ogr_fdw" >}} | {{< ext "ogr_fdw" "ogr_fdw" >}} | 1.1.7 | foreign-data wrapper for GIS data access |
| 1560 | {{< ext "geoip" >}} | {{< ext "geoip" "geoip" >}} | 0.3.0 | IP-based geolocation query |
| 1570 | {{< ext "pg_polyline" >}} | {{< ext "pg_polyline" "pg_polyline" >}} | 0.0.1 | Fast Google Encoded Polyline encoding & decoding for postgres |
| 1590 | {{< ext "pg_geohash" >}} | {{< ext "pg_geohash" "pg_geohash" >}} | 1.0 | Handle geohash based functionality for spatial coordinates |
| 1650 | {{< ext "mobilitydb" >}} | {{< ext "mobilitydb" "mobilitydb" >}} | 1.2.0 | MobilityDB geospatial trajectory data management & analysis platform |
| 1680 | {{< ext "tzf" >}} | {{< ext "tzf" "pg_tzf" >}} | 0.2.2 | Fast lookup timezone name by GPS coordinates |
| 1690 | {{< ext "earthdistance" >}} | {{< ext "earthdistance" "earthdistance" >}} | 1.2 | calculate great-circle distances on the surface of the Earth |

## RAG

Vector Database with Ivfflat, HNSW, DiskANN Indexes, AI & ML in SQL interface, Similarity Funcs, etc...

| ID | Extension | Package | Version | Description |
|:---:|:---|:---|:---|:---|
| 1800 | {{< ext "vector" >}} | {{< ext "vector" "pgvector" >}} | 0.8.1 | vector data type and ivfflat and hnsw access methods |
| 1810 | {{< ext "vchord" >}} | {{< ext "vchord" "vchord" >}} | 0.5.1 | Vector database plugin for Postgres, written in Rust |
| 1820 | {{< ext "vectorscale" >}} | {{< ext "vectorscale" "pgvectorscale" >}} | 0.8.0 | Advanced indexing for vector data with DiskANN |
| 1830 | {{< ext "vectorize" >}} | {{< ext "vectorize" "pg_vectorize" >}} | 0.22.2 | The simplest way to do vector search on Postgres |
| 1840 | {{< ext "pg_similarity" >}} | {{< ext "pg_similarity" "pg_similarity" >}} | 1.0 | support similarity queries |
| 1850 | {{< ext "smlar" >}} | {{< ext "smlar" "smlar" >}} | 1.0 | Effective similarity search |
| 1860 | {{< ext "pg_summarize" >}} | {{< ext "pg_summarize" "pg_summarize" >}} | 0.0.1 | Text Summarization using LLMs. Built using pgrx |
| 1870 | {{< ext "pg_tiktoken" >}} | {{< ext "pg_tiktoken" "pg_tiktoken" >}} | 0.0.1 | tiktoken tokenizer for use with OpenAI models in postgres |
| 1880 | {{< ext "pg4ml" >}} | {{< ext "pg4ml" "pg4ml" >}} | 2.0 | Machine learning framework for PostgreSQL |
| 1890 | {{< ext "pgml" >}} | {{< ext "pgml" "pgml" >}} | 2.10.0 | Run AL/ML workloads with SQL interface |

## FTS

ElasticSearch Alternative with BM25, 2-gram/3-gram Fuzzy Search, Zhparser & Hunspell Segregation Dicts, etc...

| ID | Extension | Package | Version | Description |
|:---:|:---|:---|:---|:---|
| 2100 | {{< ext "pg_search" >}} | {{< ext "pg_search" "pg_search" >}} | 0.18.1 | Full text search for PostgreSQL using BM25 |
| 2110 | {{< ext "pgroonga" >}} | {{< ext "pgroonga" "pgroonga" >}} | 4.0.0 | Use Groonga as index, fast full text search platform for all languages! |
| 2111 | {{< ext "pgroonga_database" >}} | {{< ext "pgroonga" "pgroonga" >}} | 4.0.0 | PGroonga database management module |
| 2120 | {{< ext "pg_bigm" >}} | {{< ext "pg_bigm" "pg_bigm" >}} | 1.2 | create 2-gram (bigram) index for faster full text search. |
| 2130 | {{< ext "zhparser" >}} | {{< ext "zhparser" "zhparser" >}} | 2.3 | a parser for full-text search of Chinese |
| 2140 | {{< ext "pg_bestmatch" >}} | {{< ext "pg_bestmatch" "pg_bestmatch" >}} | 0.0.1 | Generate BM25 sparse vector inside PostgreSQL |
| 2150 | {{< ext "vchord_bm25" >}} | {{< ext "vchord_bm25" "vchord_bm25" >}} | 0.2.1 | A postgresql extension for bm25 ranking algorithm |
| 2160 | {{< ext "pg_tokenizer" >}} | {{< ext "pg_tokenizer" "pg_tokenizer" >}} | 0.1.0 | Tokenizers for full-text search |
| 2170 | {{< ext "hunspell_cs_cz" >}} | {{< ext "hunspell_cs_cz" "hunspell_cs_cz" >}} | 1.0 | Czech Hunspell Dictionary |
| 2171 | {{< ext "hunspell_de_de" >}} | {{< ext "hunspell_de_de" "hunspell_de_de" >}} | 1.0 | German Hunspell Dictionary |
| 2172 | {{< ext "hunspell_en_us" >}} | {{< ext "hunspell_en_us" "hunspell_en_us" >}} | 1.0 | en_US Hunspell Dictionary |
| 2173 | {{< ext "hunspell_fr" >}} | {{< ext "hunspell_fr" "hunspell_fr" >}} | 1.0 | French Hunspell Dictionary |
| 2174 | {{< ext "hunspell_ne_np" >}} | {{< ext "hunspell_ne_np" "hunspell_ne_np" >}} | 1.0 | Nepali Hunspell Dictionary |
| 2175 | {{< ext "hunspell_nl_nl" >}} | {{< ext "hunspell_nl_nl" "hunspell_nl_nl" >}} | 1.0 | Dutch Hunspell Dictionary |
| 2176 | {{< ext "hunspell_nn_no" >}} | {{< ext "hunspell_nn_no" "hunspell_nn_no" >}} | 1.0 | Norwegian (norsk) Hunspell Dictionary |
| 2177 | {{< ext "hunspell_pt_pt" >}} | {{< ext "hunspell_pt_pt" "hunspell_pt_pt" >}} | 1.0 | Portuguese Hunspell Dictionary |
| 2178 | {{< ext "hunspell_ru_ru" >}} | {{< ext "hunspell_ru_ru" "hunspell_ru_ru" >}} | 1.0 | Russian Hunspell Dictionary |
| 2179 | {{< ext "hunspell_ru_ru_aot" >}} | {{< ext "hunspell_ru_ru_aot" "hunspell_ru_ru_aot" >}} | 1.0 | Russian Hunspell Dictionary (from AOT.ru group) |
| 2180 | {{< ext "fuzzystrmatch" >}} | {{< ext "fuzzystrmatch" "fuzzystrmatch" >}} | 1.2 | determine similarities and distance between strings |
| 2190 | {{< ext "pg_trgm" >}} | {{< ext "pg_trgm" "pg_trgm" >}} | 1.6 | text similarity measurement and index searching based on trigrams |

## OLAP

DuckDB Integration with FDW & PG Lakehouse, Access Parquet from File/S3, Sharding with Citus/Partman/PlProxy, ...

| ID | Extension | Package | Version | Description |
|:---:|:---|:---|:---|:---|
| 2400 | {{< ext "citus" >}} | {{< ext "citus" "citus" >}} | 13.2.0 | Distributed PostgreSQL as an extension |
| 2401 | {{< ext "citus_columnar" >}} | {{< ext "citus" "citus" >}} | 13.2.0 | Citus columnar storage engine |
| 2410 | {{< ext "columnar" >}} | {{< ext "columnar" "hydra" >}} | 1.1.2 | Hydra Columnar extension |
| 2420 | {{< ext "pg_analytics" >}} | {{< ext "pg_analytics" "pg_analytics" >}} | 0.3.7 | Postgres for analytics, powered by DuckDB |
| 2430 | {{< ext "pg_duckdb" >}} | {{< ext "pg_duckdb" "pg_duckdb" >}} | 0.3.1 | DuckDB Embedded in Postgres |
| 2440 | {{< ext "pg_mooncake" >}} | {{< ext "pg_mooncake" "pg_mooncake" >}} | 0.1.2 | Columnstore Table in Postgres |
| 2450 | {{< ext "duckdb_fdw" >}} | {{< ext "duckdb_fdw" "duckdb_fdw" >}} | 1.1.2 | DuckDB Foreign Data Wrapper |
| 2460 | {{< ext "pg_parquet" >}} | {{< ext "pg_parquet" "pg_parquet" >}} | 0.4.3 | copy data between Postgres and Parquet |
| 2500 | {{< ext "pg_fkpart" >}} | {{< ext "pg_fkpart" "pg_fkpart" >}} | 1.7.0 | Table partitioning by foreign key utility |
| 2510 | {{< ext "pg_partman" >}} | {{< ext "pg_partman" "pg_partman" >}} | 5.2.4 | Extension to manage partitioned tables by time or ID |
| 2520 | {{< ext "plproxy" >}} | {{< ext "plproxy" "plproxy" >}} | 2.11.0 | Database partitioning implemented as procedural language |
| 2530 | {{< ext "pg_strom" >}} | {{< ext "pg_strom" "pg_strom" >}} | 6.0 | PG-Strom - big-data processing acceleration using GPU and NVME |
| 2590 | {{< ext "tablefunc" >}} | {{< ext "tablefunc" "tablefunc" >}} | 1.0 | functions that manipulate whole tables, including crosstab |

## FEAT

OpenCypher with AGE, GraphQL, JsonSchema, Hints & Hypo Index, HLL, Rum, IVM, ChemRDKit, and Message Queues,...

| ID | Extension | Package | Version | Description |
|:---:|:---|:---|:---|:---|
| 2760 | {{< ext "age" >}} | {{< ext "age" "age" >}} | 1.5.0 | AGE graph database extension |
| 2770 | {{< ext "hll" >}} | {{< ext "hll" "hll" >}} | 2.18 | type for storing hyperloglog data |
| 2780 | {{< ext "rum" >}} | {{< ext "rum" "rum" >}} | 1.3.14 | RUM index access method |
| 2790 | {{< ext "pg_graphql" >}} | {{< ext "pg_graphql" "pg_graphql" >}} | 1.5.11 | Add in-database GraphQL support |
| 2800 | {{< ext "pg_jsonschema" >}} | {{< ext "pg_jsonschema" "pg_jsonschema" >}} | 0.3.3 | PostgreSQL extension providing JSON Schema validation |
| 2810 | {{< ext "jsquery" >}} | {{< ext "jsquery" "jsquery" >}} | 1.2 | data type for jsonb inspection |
| 2820 | {{< ext "pg_hint_plan" >}} | {{< ext "pg_hint_plan" "pg_hint_plan" >}} | 1.7.1 | Give PostgreSQL ability to manually force some decisions in execution plans. |
| 2830 | {{< ext "hypopg" >}} | {{< ext "hypopg" "hypopg" >}} | 1.4.2 | Hypothetical indexes for PostgreSQL |
| 2840 | {{< ext "index_advisor" >}} | {{< ext "index_advisor" "index_advisor" >}} | 0.2.0 | Query index advisor |
| 2850 | {{< ext "plan_filter" >}} | {{< ext "plan_filter" "pg_plan_filter" >}} | 0.0.1 | filter statements by their execution plans. |
| 2860 | {{< ext "imgsmlr" >}} | {{< ext "imgsmlr" "imgsmlr" >}} | 1.0 | Image similarity with haar |
| 2870 | {{< ext "pg_ivm" >}} | {{< ext "pg_ivm" "pg_ivm" >}} | 1.12 | incremental view maintenance on PostgreSQL |
| 2880 | {{< ext "pg_incremental" >}} | {{< ext "pg_incremental" "pg_incremental" >}} | 1.2.0 | Incremental Processing by Crunchy Data |
| 2900 | {{< ext "pgmq" >}} | {{< ext "pgmq" "pgmq" >}} | 1.5.1 | A lightweight message queue. Like AWS SQS and RSMQ but on Postgres. |
| 2910 | {{< ext "pgq" >}} | {{< ext "pgq" "pgq" >}} | 3.5.1 | Generic queue for PostgreSQL |
| 2920 | {{< ext "orioledb" >}} | {{< ext "orioledb" "orioledb" >}} | 1.5 | OrioleDB, the next generation transactional engine |
| 2930 | {{< ext "pg_cardano" >}} | {{< ext "pg_cardano" "pg_cardano" >}} | 1.0.5 | A suite of Cardano-related tools |
| 2940 | {{< ext "rdkit" >}} | {{< ext "rdkit" "rdkit" >}} | 202503.1 | Cheminformatics functionality for PostgreSQL. |
| 2951 | {{< ext "omni" >}} | {{< ext "omni" "omnigres" >}} | 0.2.9 | Advanced adapter for Postgres extensions |
| 2952 | {{< ext "omni_auth" >}} | {{< ext "omni" "omnigres" >}} | 0.1.3 | Basic session management |
| 2953 | {{< ext "omni_aws" >}} | {{< ext "omni" "omnigres" >}} | 0.1.2 | Amazon Web Services APIs (S3) |
| 2954 | {{< ext "omni_cloudevents" >}} | {{< ext "omni" "omnigres" >}} | 0.1.0 | CloudEvents support |
| 2955 | {{< ext "omni_containers" >}} | {{< ext "omni" "omnigres" >}} | 0.2.0 | Docker container management |
| 2956 | {{< ext "omni_credentials" >}} | {{< ext "omni" "omnigres" >}} | 0.2.0 | Application credential management |
| 2958 | {{< ext "omni_email" >}} | {{< ext "omni" "omnigres" >}} | 0 | E-mail framework |
| 2959 | {{< ext "omni_http" >}} | {{< ext "omni" "omnigres" >}} | 0.1.0 | Basic HTTP types |
| 2960 | {{< ext "omni_httpc" >}} | {{< ext "omni" "omnigres" >}} | 0.1.5 | HTTP client |
| 2961 | {{< ext "omni_httpd" >}} | {{< ext "omni" "omnigres" >}} | 0.4.6 | HTTP server |
| 2962 | {{< ext "omni_id" >}} | {{< ext "omni" "omnigres" >}} | 0.4.2 | Identity types |
| 2963 | {{< ext "omni_json" >}} | {{< ext "omni" "omnigres" >}} | 0.1.1 | JSON toolkit |
| 2964 | {{< ext "omni_kube" >}} | {{< ext "omni" "omnigres" >}} | 0.1.1 | Kubernetes (k8s) integration |
| 2965 | {{< ext "omni_ledger" >}} | {{< ext "omni" "omnigres" >}} | 0.1.2 | Financial ledger |
| 2966 | {{< ext "omni_manifest" >}} | {{< ext "omni" "omnigres" >}} | 0.1.2 | Package installation manifests |
| 2967 | {{< ext "omni_mimetypes" >}} | {{< ext "omni" "omnigres" >}} | 0.1.0 | MIME types |
| 2968 | {{< ext "omni_os" >}} | {{< ext "omni" "omnigres" >}} | 0.1.1 | Operating system integration |
| 2969 | {{< ext "omni_polyfill" >}} | {{< ext "omni" "omnigres" >}} | 0.2.2 | Postgres API polyfills |
| 2970 | {{< ext "omni_python" >}} | {{< ext "omni" "omnigres" >}} | 0.1.1 | First-class Python support |
| 2971 | {{< ext "omni_regex" >}} | {{< ext "omni" "omnigres" >}} | 0.1.0 | PCRE-compatible regular expressions |
| 2972 | {{< ext "omni_rest" >}} | {{< ext "omni" "omnigres" >}} | 0.1.1 | REST API toolkit (with PostgREST support) |
| 2973 | {{< ext "omni_schema" >}} | {{< ext "omni" "omnigres" >}} | 0.3.0 | Advanced schema management tooling |
| 2974 | {{< ext "omni_seq" >}} | {{< ext "omni" "omnigres" >}} | 0.1.1 | Distributed integer sequences |
| 2975 | {{< ext "omni_service" >}} | {{< ext "omni" "omnigres" >}} | 0.1.0 | Service management |
| 2976 | {{< ext "omni_session" >}} | {{< ext "omni" "omnigres" >}} | 0.2.0 | Session management |
| 2977 | {{< ext "omni_sql" >}} | {{< ext "omni" "omnigres" >}} | 0.5.1 | Programmatic SQL manipulation |
| 2979 | {{< ext "omni_sqlite" >}} | {{< ext "omni" "omnigres" >}} | 0.1.2 | Embedded SQLite |
| 2980 | {{< ext "omni_test" >}} | {{< ext "omni" "omnigres" >}} | 0.4.0 | Testing framework |
| 2981 | {{< ext "omni_txn" >}} | {{< ext "omni" "omnigres" >}} | 0.5.0 | Transaction management |
| 2982 | {{< ext "omni_types" >}} | {{< ext "omni" "omnigres" >}} | 0.3.4 | Advanced types |
| 2983 | {{< ext "omni_var" >}} | {{< ext "omni" "omnigres" >}} | 0.3.0 | Scoped variables |
| 2984 | {{< ext "omni_vfs" >}} | {{< ext "omni" "omnigres" >}} | 0.2.1 | Virtual File System |
| 2985 | {{< ext "omni_vfs_types_v1" >}} | {{< ext "omni" "omnigres" >}} | 0.1.0 | Virtual File System types (v1) |
| 2986 | {{< ext "omni_web" >}} | {{< ext "omni" "omnigres" >}} | 0.3.0 | Common web stack primitives |
| 2987 | {{< ext "omni_worker" >}} | {{< ext "omni" "omnigres" >}} | 0.1.0 | Generalized worker pool |
| 2988 | {{< ext "omni_xml" >}} | {{< ext "omni" "omnigres" >}} | 0.1.2 | XML toolkit |
| 2989 | {{< ext "omni_yaml" >}} | {{< ext "omni" "omnigres" >}} | 0.1.0 | YAML toolkit |
| 2990 | {{< ext "bloom" >}} | {{< ext "bloom" "bloom" >}} | 1.0 | bloom access method - signature file based index |

## LANG

Develop, Test, Package, and Deliver Stored Procedures written in various PL/Languages: Java, Js, Lua, R, Sh, PRQL, ...

| ID | Extension | Package | Version | Description |
|:---:|:---|:---|:---|:---|
| 3000 | {{< ext "pg_tle" >}} | {{< ext "pg_tle" "pg_tle" >}} | 1.5.1 | Trusted Language Extensions for PostgreSQL |
| 3010 | {{< ext "plv8" >}} | {{< ext "plv8" "plv8" >}} | 3.2.4 | PL/JavaScript (v8) trusted procedural language |
| 3020 | {{< ext "pllua" >}} | {{< ext "pllua" "pllua" >}} | 2.0.12 | Lua as a procedural language |
| 3021 | {{< ext "hstore_pllua" >}} | {{< ext "pllua" "pllua" >}} | 2.0.12 | Hstore transform for Lua |
| 3030 | {{< ext "plluau" >}} | {{< ext "pllua" "pllua" >}} | 2.0.12 | Lua as an untrusted procedural language |
| 3031 | {{< ext "hstore_plluau" >}} | {{< ext "pllua" "pllua" >}} | 2.0.12 | Hstore transform for untrusted Lua |
| 3040 | {{< ext "plprql" >}} | {{< ext "plprql" "plprql" >}} | 1.0.0 | Use PRQL in PostgreSQL - Pipelined Relational Query Language |
| 3050 | {{< ext "pldbgapi" >}} | {{< ext "pldbgapi" "pldebugger" >}} | 1.8 | server-side support for debugging PL/pgSQL functions |
| 3060 | {{< ext "plpgsql_check" >}} | {{< ext "plpgsql_check" "plpgsql_check" >}} | 2.8.2 | extended check for plpgsql functions |
| 3070 | {{< ext "plprofiler" >}} | {{< ext "plprofiler" "plprofiler" >}} | 4.2.5 | server-side support for profiling PL/pgSQL functions |
| 3080 | {{< ext "plsh" >}} | {{< ext "plsh" "plsh" >}} | 1.20220917 | PL/sh procedural language |
| 3090 | {{< ext "pljava" >}} | {{< ext "pljava" "pljava" >}} | 1.6.9 | PL/Java procedural language |
| 3100 | {{< ext "plr" >}} | {{< ext "plr" "plr" >}} | 8.4.8 | load R interpreter and execute R script from within a database |
| 3200 | {{< ext "pgtap" >}} | {{< ext "pgtap" "pgtap" >}} | 1.3.3 | Unit testing for PostgreSQL |
| 3210 | {{< ext "faker" >}} | {{< ext "faker" "faker" >}} | 0.5.3 | Wrapper for the Faker Python library |
| 3220 | {{< ext "dbt2" >}} | {{< ext "dbt2" "dbt2" >}} | 0.45.0 | OSDL-DBT-2 test kit |
| 3240 | {{< ext "pltcl" >}} | {{< ext "pltcl" "pltcl" >}} | 1.0 | PL/Tcl procedural language |
| 3250 | {{< ext "pltclu" >}} | {{< ext "pltcl" "pltcl" >}} | 1.0 | PL/TclU untrusted procedural language |
| 3260 | {{< ext "plperl" >}} | {{< ext "plperl" "plperl" >}} | 1.0 | PL/Perl procedural language |
| 3261 | {{< ext "bool_plperl" >}} | {{< ext "plperl" "plperl" >}} | 1.0 | transform between bool and plperl |
| 3262 | {{< ext "hstore_plperl" >}} | {{< ext "plperl" "plperl" >}} | 1.0 | transform between hstore and plperl |
| 3263 | {{< ext "jsonb_plperl" >}} | {{< ext "plperl" "plperl" >}} | 1.0 | transform between jsonb and plperl |
| 3270 | {{< ext "plperlu" >}} | {{< ext "plperlu" "plperlu" >}} | 1.0 | PL/PerlU untrusted procedural language |
| 3271 | {{< ext "bool_plperlu" >}} | {{< ext "plperlu" "plperlu" >}} | 1.0 | transform between bool and plperlu |
| 3272 | {{< ext "jsonb_plperlu" >}} | {{< ext "plperlu" "plperlu" >}} | 1.0 | transform between jsonb and plperlu |
| 3273 | {{< ext "hstore_plperlu" >}} | {{< ext "plperlu" "plperlu" >}} | 1.0 | transform between hstore and plperlu |
| 3280 | {{< ext "plpgsql" >}} | {{< ext "plpgsql" "plpgsql" >}} | 1.0 | PL/pgSQL procedural language |
| 3290 | {{< ext "plpython3u" >}} | {{< ext "plpython3u" "plpython3u" >}} | 1.0 | PL/Python3U untrusted procedural language |
| 3291 | {{< ext "jsonb_plpython3u" >}} | {{< ext "plpython3u" "plpython3u" >}} | 1.0 | transform between jsonb and plpython3u |
| 3292 | {{< ext "ltree_plpython3u" >}} | {{< ext "plpython3u" "plpython3u" >}} | 1.0 | transform between ltree and plpython3u |
| 3293 | {{< ext "hstore_plpython3u" >}} | {{< ext "plpython3u" "plpython3u" >}} | 1.0 | transform between hstore and plpython3u |

## TYPE

Dedicate New Data Types Like: prefix, semver, uint, SIUnit, RoaringBitmap, Rational, Sphere, Hash, RRule, and more...

| ID | Extension | Package | Version | Description |
|:---:|:---|:---|:---|:---|
| 3500 | {{< ext "prefix" >}} | {{< ext "prefix" "pg_prefix" >}} | 1.2.10 | Prefix Range module for PostgreSQL |
| 3510 | {{< ext "semver" >}} | {{< ext "semver" "pg_semver" >}} | 0.40.0 | Semantic version data type |
| 3520 | {{< ext "unit" >}} | {{< ext "unit" "pgunit" >}} | 7.10 | SI units extension |
| 3530 | {{< ext "pgpdf" >}} | {{< ext "pgpdf" "pgpdf" >}} | 0.1.0 | PDF type with meta admin & Full-Text Search |
| 3540 | {{< ext "pglite_fusion" >}} | {{< ext "pglite_fusion" "pglite_fusion" >}} | 0.0.5 | Embed an SQLite database in your PostgreSQL table |
| 3550 | {{< ext "md5hash" >}} | {{< ext "md5hash" "md5hash" >}} | 1.0.1 | type for storing 128-bit binary data inline |
| 3560 | {{< ext "asn1oid" >}} | {{< ext "asn1oid" "asn1oid" >}} | 1.6 | asn1oid extension |
| 3570 | {{< ext "roaringbitmap" >}} | {{< ext "roaringbitmap" "roaringbitmap" >}} | 0.5.4 | support for Roaring Bitmaps |
| 3580 | {{< ext "pgfaceting" >}} | {{< ext "pgfaceting" "pgfaceting" >}} | 0.2.0 | fast faceting queries using an inverted index |
| 3590 | {{< ext "pg_sphere" >}} | {{< ext "pg_sphere" "pgsphere" >}} | 1.5.1 | spherical objects with useful functions, operators and index support |
| 3600 | {{< ext "country" >}} | {{< ext "country" "pg_country" >}} | 0.0.3 | Country data type, ISO 3166-1 |
| 3610 | {{< ext "pg_xenophile" >}} | {{< ext "pg_xenophile" "pg_xenophile" >}} | 0.8.3 | More than the bare necessities for PostgreSQL i18n and l10n. |
| 3611 | {{< ext "l10n_table_dependent_extension" >}} | {{< ext "pg_xenophile" "pg_xenophile" >}} | 0.8.3 | PostgreSQL l10n toolbox |
| 3620 | {{< ext "currency" >}} | {{< ext "currency" "pg_currency" >}} | 0.0.3 | Custom PostgreSQL currency type in 1Byte |
| 3630 | {{< ext "collection" >}} | {{< ext "collection" "pg_collection" >}} | 1.0.0 | Memory optimized data type to be used inside of plpglsql func |
| 3700 | {{< ext "pgmp" >}} | {{< ext "pgmp" "pgmp" >}} | 1.0.5 | Multiple Precision Arithmetic extension |
| 3710 | {{< ext "numeral" >}} | {{< ext "numeral" "numeral" >}} | 1.3 | numeral datatypes extension |
| 3720 | {{< ext "pg_rational" >}} | {{< ext "pg_rational" "pg_rational" >}} | 0.0.2 | bigint fractions |
| 3730 | {{< ext "uint" >}} | {{< ext "uint" "pguint" >}} | 1.20231206 | unsigned integer types |
| 3740 | {{< ext "uint128" >}} | {{< ext "uint128" "pg_uint128" >}} | 1.1.0 | Native uint128 type |
| 3750 | {{< ext "hashtypes" >}} | {{< ext "hashtypes" "hashtypes" >}} | 0.1.5 | sha1, md5 and other data types for PostgreSQL |
| 3820 | {{< ext "ip4r" >}} | {{< ext "ip4r" "ip4r" >}} | 2.4.2 | IPv4/v6 and IPv4/v6 range index type for PostgreSQL |
| 3830 | {{< ext "pg_duration" >}} | {{< ext "pg_duration" "pg_duration" >}} | 1.0.2 | data type for representing durations |
| 3840 | {{< ext "uri" >}} | {{< ext "uri" "pg_uri" >}} | 1.20151224 | URI Data type for PostgreSQL |
| 3850 | {{< ext "emailaddr" >}} | {{< ext "emailaddr" "pgemailaddr" >}} | 0 | Email address type for PostgreSQL |
| 3860 | {{< ext "acl" >}} | {{< ext "acl" "pg_acl" >}} | 1.0.4 | ACL Data type |
| 3870 | {{< ext "debversion" >}} | {{< ext "debversion" "debversion" >}} | 1.2.0 | Debian version number data type |
| 3880 | {{< ext "pg_rrule" >}} | {{< ext "pg_rrule" "pg_rrule" >}} | 0.2.0 | RRULE field type for PostgreSQL |
| 3890 | {{< ext "timestamp9" >}} | {{< ext "timestamp9" "timestamp9" >}} | 1.4.0 | timestamp nanosecond resolution |
| 3920 | {{< ext "chkpass" >}} | {{< ext "chkpass" "chkpass" >}} | 1.0 | data type for auto-encrypted passwords |
| 3930 | {{< ext "isn" >}} | {{< ext "isn" "isn" >}} | 1.2 | data types for international product numbering standards |
| 3940 | {{< ext "seg" >}} | {{< ext "seg" "seg" >}} | 1.4 | data type for representing line segments or floating-point intervals |
| 3950 | {{< ext "cube" >}} | {{< ext "cube" "cube" >}} | 1.5 | data type for multidimensional cubes |
| 3960 | {{< ext "ltree" >}} | {{< ext "ltree" "ltree" >}} | 1.3 | data type for hierarchical tree-like structures |
| 3970 | {{< ext "hstore" >}} | {{< ext "hstore" "hstore" >}} | 1.8 | data type for storing sets of (key, value) pairs |
| 3980 | {{< ext "citext" >}} | {{< ext "citext" "citext" >}} | 1.6 | data type for case-insensitive character strings |
| 3990 | {{< ext "xml2" >}} | {{< ext "xml2" "xml2" >}} | 1.1 | XPath querying and XSLT |

## UTIL

Utilities such as send http request, perform gzip/zstd compress, send mails, Regex, ICU, encoding, docs, Encryption,...

| ID | Extension | Package | Version | Description |
|:---:|:---|:---|:---|:---|
| 4010 | {{< ext "gzip" >}} | {{< ext "gzip" "pg_gzip" >}} | 1.0.1 | gzip and gunzip functions. |
| 4020 | {{< ext "bzip" >}} | {{< ext "bzip" "pg_bzip" >}} | 1.0.0 | Bzip compression and decompression |
| 4030 | {{< ext "zstd" >}} | {{< ext "zstd" "pg_zstd" >}} | 1.1.2 | Zstandard compression algorithm implementation in PostgreSQL |
| 4070 | {{< ext "http" >}} | {{< ext "http" "pg_http" >}} | 1.7.0 | HTTP client for PostgreSQL, allows web page retrieval inside the database. |
| 4080 | {{< ext "pg_net" >}} | {{< ext "pg_net" "pg_net" >}} | 0.9.2 | Async HTTP Requests |
| 4090 | {{< ext "pg_curl" >}} | {{< ext "pg_curl" "pg_curl" >}} | 2.4.5 | Run curl actions for data transfer in URL syntax |
| 4150 | {{< ext "pgjq" >}} | {{< ext "pgjq" "pgjq" >}} | 0.1.0 | Use jq in Postgres |
| 4160 | {{< ext "pgjwt" >}} | {{< ext "pgjwt" "pgjwt" >}} | 0.2.0 | JSON Web Token API for Postgresql |
| 4170 | {{< ext "pg_smtp_client" >}} | {{< ext "pg_smtp_client" "pg_smtp_client" >}} | 0.2.0 | PostgreSQL extension to send email using SMTP |
| 4180 | {{< ext "pg_html5_email_address" >}} | {{< ext "pg_html5_email_address" "pg_html5_email_address" >}} | 1.2.3 | PostgreSQL email validation that is consistent with the HTML5 spec |
| 4190 | {{< ext "url_encode" >}} | {{< ext "url_encode" "url_encode" >}} | 1.2.5 | url_encode, url_decode functions |
| 4200 | {{< ext "pgsql_tweaks" >}} | {{< ext "pgsql_tweaks" "pgsql_tweaks" >}} | 0.11.3 | Some functions and views for daily usage |
| 4220 | {{< ext "pg_extra_time" >}} | {{< ext "pg_extra_time" "pg_extra_time" >}} | 2.0.0 | Some date time functions and operators that, |
| 4230 | {{< ext "pgpcre" >}} | {{< ext "pgpcre" "pgpcre" >}} | 1 | Perl Compatible Regular Expression functions |
| 4240 | {{< ext "icu_ext" >}} | {{< ext "icu_ext" "icu_ext" >}} | 1.10.0 | Access ICU functions |
| 4250 | {{< ext "pgqr" >}} | {{< ext "pgqr" "pgqr" >}} | 1.0 | QR Code generator from PostgreSQL |
| 4260 | {{< ext "pg_protobuf" >}} | {{< ext "pg_protobuf" "pg_protobuf" >}} | 1.0 | Protobuf support for PostgreSQL |
| 4270 | {{< ext "envvar" >}} | {{< ext "envvar" "envvar" >}} | 1.0.1 | Fetch the value of an environment variable |
| 4280 | {{< ext "floatfile" >}} | {{< ext "floatfile" "floatfile" >}} | 1.3.1 | Simple file storage for arrays of floats |
| 4290 | {{< ext "pg_render" >}} | {{< ext "pg_render" "pg_render" >}} | 0.1.2 | Render HTML in SQL |
| 4300 | {{< ext "pg_readme" >}} | {{< ext "pg_readme" "pg_readme" >}} | 0.7.0 | Generate a README.md document for a database extension or schema |
| 4301 | {{< ext "pg_readme_test_extension" >}} | {{< ext "pg_readme" "pg_readme" >}} | 0.7.0 | Test generating a README.md document for extension or schema |
| 4310 | {{< ext "ddl_historization" >}} | {{< ext "ddl_historization" "ddl_historization" >}} | 0.0.7 | Historize the ddl changes inside PostgreSQL database |
| 4320 | {{< ext "data_historization" >}} | {{< ext "data_historization" "data_historization" >}} | 1.1.0 | PLPGSQL Script to historize data in partitionned table |
| 4330 | {{< ext "schedoc" >}} | {{< ext "schedoc" "pg_schedoc" >}} | 0.0.1 | Cross documentation between Django and DBT projects |
| 4400 | {{< ext "hashlib" >}} | {{< ext "hashlib" "pg_hashlib" >}} | 1.1 | Stable hash functions for Postgres |
| 4430 | {{< ext "xxhash" >}} | {{< ext "xxhash" "pg_xxhash" >}} | 0.0.1 | xxhash functions for PostgreSQL |
| 4440 | {{< ext "shacrypt" >}} | {{< ext "shacrypt" "shacrypt" >}} | 1.1 | Implements SHA256-CRYPT and SHA512-CRYPT password encryption schemes |
| 4450 | {{< ext "cryptint" >}} | {{< ext "cryptint" "cryptint" >}} | 1.0.0 | Encryption functions for int and bigint values |
| 4460 | {{< ext "pguecc" >}} | {{< ext "pguecc" "pg_ecdsa" >}} | 1.0 | uECC bindings for Postgres |
| 4470 | {{< ext "sparql" >}} | {{< ext "sparql" "pgsparql" >}} | 1.0 | Query SPARQL datasource with SQL |

## FUNC

Function such as id generator, aggregations, sketches, vector functions, mathematical functions and digest functions...

| ID | Extension | Package | Version | Description |
|:---:|:---|:---|:---|:---|
| 4500 | {{< ext "pg_idkit" >}} | {{< ext "pg_idkit" "pg_idkit" >}} | 0.3.1 | multi-tool for generating new/niche universally unique identifiers (ex. UUIDv6, ULID, KSUID) |
| 4510 | {{< ext "pgx_ulid" >}} | {{< ext "pgx_ulid" "pgx_ulid" >}} | 0.2.0 | ulid type and methods |
| 4540 | {{< ext "pg_uuidv7" >}} | {{< ext "pg_uuidv7" "pg_uuidv7" >}} | 1.6.0 | Create UUIDv7 values in postgres |
| 4550 | {{< ext "permuteseq" >}} | {{< ext "permuteseq" "permuteseq" >}} | 1.2.2 | Pseudo-randomly permute sequences with a format-preserving encryption on elements |
| 4560 | {{< ext "pg_hashids" >}} | {{< ext "pg_hashids" "pg_hashids" >}} | 1.3 | Short unique id generator for PostgreSQL, using hashids |
| 4570 | {{< ext "sequential_uuids" >}} | {{< ext "sequential_uuids" "sequential_uuids" >}} | 1.0.3 | generator of sequential UUIDs |
| 4600 | {{< ext "topn" >}} | {{< ext "topn" "topn" >}} | 2.7.0 | type for top-n JSONB |
| 4610 | {{< ext "quantile" >}} | {{< ext "quantile" "quantile" >}} | 1.1.8 | Quantile aggregation function |
| 4620 | {{< ext "lower_quantile" >}} | {{< ext "lower_quantile" "lower_quantile" >}} | 1.0.3 | Lower quantile aggregate function |
| 4630 | {{< ext "count_distinct" >}} | {{< ext "count_distinct" "count_distinct" >}} | 3.0.2 | An alternative to COUNT(DISTINCT …) aggregate, usable with HashAggregate |
| 4640 | {{< ext "omnisketch" >}} | {{< ext "omnisketch" "omnisketch" >}} | 1.0.2 | data structure for on-line agg of data into approximate sketch |
| 4650 | {{< ext "ddsketch" >}} | {{< ext "ddsketch" "ddsketch" >}} | 1.0.1 | Provides ddsketch aggregate function |
| 4660 | {{< ext "vasco" >}} | {{< ext "vasco" "vasco" >}} | 0.1.0 | discover hidden correlations in your data with MIC |
| 4670 | {{< ext "xicor" >}} | {{< ext "xicor" "pgxicor" >}} | 0.1.0 | XI Correlation Coefficient in Postgres |
| 4700 | {{< ext "tdigest" >}} | {{< ext "tdigest" "tdigest" >}} | 1.4.3 | Provides tdigest aggregate function. |
| 4710 | {{< ext "first_last_agg" >}} | {{< ext "first_last_agg" "first_last_agg" >}} | 0.1.4 | first() and last() aggregate functions |
| 4720 | {{< ext "extra_window_functions" >}} | {{< ext "extra_window_functions" "extra_window_functions" >}} | 1.0 | Extra Window Functions for PostgreSQL |
| 4730 | {{< ext "floatvec" >}} | {{< ext "floatvec" "floatvec" >}} | 1.1.1 | Math for vectors (arrays) of numbers |
| 4740 | {{< ext "aggs_for_vecs" >}} | {{< ext "aggs_for_vecs" "aggs_for_vecs" >}} | 1.4.0 | Aggregate functions for array inputs |
| 4750 | {{< ext "aggs_for_arrays" >}} | {{< ext "aggs_for_arrays" "aggs_for_arrays" >}} | 1.3.3 | Various functions for computing statistics on arrays of numbers |
| 4760 | {{< ext "arraymath" >}} | {{< ext "arraymath" "pg_arraymath" >}} | 1.1 | Array math and operators that work element by element on the contents of arrays |
| 4770 | {{< ext "pg_math" >}} | {{< ext "pg_math" "pg_math" >}} | 1.0 | GSL statistical functions for postgresql |
| 4780 | {{< ext "random" >}} | {{< ext "random" "pg_random" >}} | 2.0.0 | random data generator |
| 4800 | {{< ext "base36" >}} | {{< ext "base36" "pg_base36" >}} | 1.0.0 | Integer Base36 types |
| 4810 | {{< ext "base62" >}} | {{< ext "base62" "pg_base62" >}} | 0.0.1 | Base62 extension for PostgreSQL |
| 4830 | {{< ext "pg_base58" >}} | {{< ext "pg_base58" "pg_base58" >}} | 0.0.1 | Base58 Encoder/Decoder Extension for PostgreSQL |
| 4840 | {{< ext "financial" >}} | {{< ext "financial" "pg_financial" >}} | 1.0.1 | Financial aggregate functions |
| 4850 | {{< ext "convert" >}} | {{< ext "convert" "pg_convert" >}} | 0.0.4 | conversion functions for spatial, routing and other specialized uses |
| 4880 | {{< ext "refint" >}} | {{< ext "refint" "refint" >}} | 1.0 | functions for implementing referential integrity (obsolete) |
| 4881 | {{< ext "autoinc" >}} | {{< ext "autoinc" "autoinc" >}} | 1.0 | functions for autoincrementing fields |
| 4882 | {{< ext "insert_username" >}} | {{< ext "insert_username" "insert_username" >}} | 1.0 | functions for tracking who changed a table |
| 4883 | {{< ext "moddatetime" >}} | {{< ext "moddatetime" "moddatetime" >}} | 1.0 | functions for tracking last modification time |
| 4890 | {{< ext "tsm_system_time" >}} | {{< ext "tsm_system_time" "tsm_system_time" >}} | 1.0 | TABLESAMPLE method which accepts time in milliseconds as a limit |
| 4900 | {{< ext "dict_xsyn" >}} | {{< ext "dict_xsyn" "dict_xsyn" >}} | 1.0 | text search dictionary template for extended synonym processing |
| 4910 | {{< ext "tsm_system_rows" >}} | {{< ext "tsm_system_rows" "tsm_system_rows" >}} | 1.0 | TABLESAMPLE method which accepts number of rows as a limit |
| 4920 | {{< ext "tcn" >}} | {{< ext "tcn" "tcn" >}} | 1.0 | Triggered change notifications |
| 4930 | {{< ext "uuid-ossp" >}} | {{< ext "uuid-ossp" "uuid-ossp" >}} | 1.1 | generate universally unique identifiers (UUIDs) |
| 4940 | {{< ext "btree_gist" >}} | {{< ext "btree_gist" "btree_gist" >}} | 1.7 | support for indexing common datatypes in GiST |
| 4950 | {{< ext "btree_gin" >}} | {{< ext "btree_gin" "btree_gin" >}} | 1.3 | support for indexing common datatypes in GIN |
| 4960 | {{< ext "intarray" >}} | {{< ext "intarray" "intarray" >}} | 1.5 | functions, operators, and index support for 1-D arrays of integers |
| 4970 | {{< ext "intagg" >}} | {{< ext "intagg" "intagg" >}} | 1.1 | integer aggregator and enumerator (obsolete) |
| 4980 | {{< ext "dict_int" >}} | {{< ext "dict_int" "dict_int" >}} | 1.0 | text search dictionary template for integers |
| 4990 | {{< ext "unaccent" >}} | {{< ext "unaccent" "unaccent" >}} | 1.1 | text search dictionary that removes accents |

## ADMIN

Utilities for Bloat Control, DirtyRead, BufferInspect, DDL Generate, ChecksumVerify, Permission, Priority, Catalog,...

| ID | Extension | Package | Version | Description |
|:---:|:---|:---|:---|:---|
| 5010 | {{< ext "pg_repack" >}} | {{< ext "pg_repack" "pg_repack" >}} | 1.5.2 | Reorganize tables in PostgreSQL databases with minimal locks |
| 5020 | {{< ext "pg_rewrite" >}} | {{< ext "pg_rewrite" "pg_rewrite" >}} | 2.0.0 | Tool allows read write to the table during the rewriting |
| 5040 | {{< ext "pg_squeeze" >}} | {{< ext "pg_squeeze" "pg_squeeze" >}} | 1.9.0 | A tool to remove unused space from a relation. |
| 5050 | {{< ext "pg_dirtyread" >}} | {{< ext "pg_dirtyread" "pg_dirtyread" >}} | 2.7 | Read dead but unvacuumed rows from table |
| 5060 | {{< ext "pgfincore" >}} | {{< ext "pgfincore" "pgfincore" >}} | 1.3.1 | examine and manage the os buffer cache |
| 5070 | {{< ext "pg_cooldown" >}} | {{< ext "pg_cooldown" "pg_cooldown" >}} | 0.1 | remove buffered pages for specific relations |
| 5080 | {{< ext "ddlx" >}} | {{< ext "ddlx" "pg_ddlx" >}} | 0.30 | DDL eXtractor functions |
| 5090 | {{< ext "prioritize" >}} | {{< ext "prioritize" "pg_prioritize" >}} | 1.0.4 | get and set the priority of PostgreSQL backends |
| 5110 | {{< ext "pg_checksums" >}} | {{< ext "pg_checksums" "pg_checksums" >}} | 1.2 | Activate/deactivate/verify checksums in offline Postgres clusters |
| 5120 | {{< ext "pg_readonly" >}} | {{< ext "pg_readonly" "pg_readonly" >}} | 1.0.3 | cluster database read only |
| 5130 | {{< ext "pgdd" >}} | {{< ext "pgdd" "pgdd" >}} | 0.6.0 | Introspect pg data dictionary via standard SQL |
| 5140 | {{< ext "pg_permissions" >}} | {{< ext "pg_permissions" "pg_permissions" >}} | 1.4 | view object permissions and compare them with the desired state |
| 5150 | {{< ext "pgautofailover" >}} | {{< ext "pgautofailover" "pgautofailover" >}} | 2.2 | pg_auto_failover |
| 5160 | {{< ext "pg_catcheck" >}} | {{< ext "pg_catcheck" "pg_catcheck" >}} | 1.6.0 | Diagnosing system catalog corruption |
| 5170 | {{< ext "pre_prepare" >}} | {{< ext "pre_prepare" "preprepare" >}} | 0.9 | Pre Prepare your Statement server side |
| 5180 | {{< ext "pg_upless" >}} | {{< ext "pg_upless" "pg_upless" >}} | 0.0.3 | Detect Useless UPDATE |
| 5190 | {{< ext "pgcozy" >}} | {{< ext "pgcozy" "pgcozy" >}} | 1.0 | Pre-warming shared buffers according to previous pg_buffercache snapshots for PostgreSQL. |
| 5200 | {{< ext "pg_orphaned" >}} | {{< ext "pg_orphaned" "pg_orphaned" >}} | 1.0 | Deal with orphaned files |
| 5210 | {{< ext "pg_crash" >}} | {{< ext "pg_crash" "pg_crash" >}} | 1.0 | Send random signals to random processes |
| 5220 | {{< ext "pg_cheat_funcs" >}} | {{< ext "pg_cheat_funcs" "pg_cheat_funcs" >}} | 1.0 | Provides cheat (but useful) functions |
| 5230 | {{< ext "fio" >}} | {{< ext "fio" "pg_fio" >}} | 1.0 | PostgreSQL File I/O Functions |
| 5810 | {{< ext "pg_savior" >}} | {{< ext "pg_savior" "pg_savior" >}} | 0.0.1 | Postgres extension to save OOPS mistakes |
| 5820 | {{< ext "safeupdate" >}} | {{< ext "safeupdate" "safeupdate" >}} | 1.5 | Require criteria for UPDATE and DELETE |
| 5830 | {{< ext "pg_drop_events" >}} | {{< ext "pg_drop_events" "pg_drop_events" >}} | 0.1.0 | logs transaction ids of drop table, drop column, drop materialized view statements |
| 5840 | {{< ext "table_log" >}} | {{< ext "table_log" "table_log" >}} | 0.6.4 | record table modification logs and PITR for table/row |
| 5880 | {{< ext "pgagent" >}} | {{< ext "pgagent" "pgagent" >}} | 4.2.3 | A PostgreSQL job scheduler |
| 5890 | {{< ext "pg_prewarm" >}} | {{< ext "pg_prewarm" "pg_prewarm" >}} | 1.2 | prewarm relation data |
| 5900 | {{< ext "pgpool_adm" >}} | {{< ext "pgpool_adm" "pgpool" >}} | 4.6.3 | Administrative functions for pgPool |
| 5910 | {{< ext "pgpool_recovery" >}} | {{< ext "pgpool_adm" "pgpool" >}} | 4.6.3 | recovery functions for pgpool-II for V4.3 |
| 5920 | {{< ext "pgpool_regclass" >}} | {{< ext "pgpool_adm" "pgpool" >}} | 4.6.3 | replacement for regclass |
| 5930 | {{< ext "lo" >}} | {{< ext "lo" "lo" >}} | 1.1 | Large Object maintenance |
| 5940 | {{< ext "basic_archive" >}} | {{< ext "basic_archive" "basic_archive" >}} | - | an example of an archive module |
| 5950 | {{< ext "basebackup_to_shell" >}} | {{< ext "basebackup_to_shell" "basebackup_to_shell" >}} | - | adds a custom basebackup target called shell |
| 5960 | {{< ext "old_snapshot" >}} | {{< ext "old_snapshot" "old_snapshot" >}} | 1.0 | utilities in support of old_snapshot_threshold |
| 5970 | {{< ext "adminpack" >}} | {{< ext "adminpack" "adminpack" >}} | 2.1 | administrative functions for PostgreSQL |
| 5980 | {{< ext "amcheck" >}} | {{< ext "amcheck" "amcheck" >}} | 1.4 | functions for verifying relation integrity |
| 5990 | {{< ext "pg_surgery" >}} | {{< ext "pg_surgery" "pg_surgery" >}} | 1.0 | extension to perform surgery on a damaged relation |

## STAT

Observability Catalogs, Monitoring Metrics & Views, Statistics, Query Plans, WaitSampling, SlowLogs, and etc...

| ID | Extension | Package | Version | Description |
|:---:|:---|:---|:---|:---|
| 6000 | {{< ext "pg_profile" >}} | {{< ext "pg_profile" "pg_profile" >}} | 4.10 | PostgreSQL load profile repository and report builder |
| 6010 | {{< ext "pg_tracing" >}} | {{< ext "pg_tracing" "pg_tracing" >}} | 0.1.3 | Distributed Tracing for PostgreSQL |
| 6210 | {{< ext "pg_show_plans" >}} | {{< ext "pg_show_plans" "pg_show_plans" >}} | 2.1.6 | show query plans of all currently running SQL statements |
| 6220 | {{< ext "pg_stat_kcache" >}} | {{< ext "pg_stat_kcache" "pg_stat_kcache" >}} | 2.3.0 | Kernel statistics gathering |
| 6230 | {{< ext "pg_stat_monitor" >}} | {{< ext "pg_stat_monitor" "pg_stat_monitor" >}} | 2.2.0 | The pg_stat_monitor is a PostgreSQL Query Performance Monitoring tool, based on PostgreSQL contrib module pg_stat_statements. pg_stat_monitor provides aggregated statistics, client information, plan details including plan, and histogram information. |
| 6240 | {{< ext "pg_qualstats" >}} | {{< ext "pg_qualstats" "pg_qualstats" >}} | 2.1.2 | An extension collecting statistics about quals |
| 6250 | {{< ext "pg_store_plans" >}} | {{< ext "pg_store_plans" "pg_store_plans" >}} | 1.8 | track plan statistics of all SQL statements executed |
| 6260 | {{< ext "pg_track_settings" >}} | {{< ext "pg_track_settings" "pg_track_settings" >}} | 2.1.2 | Track settings changes |
| 6270 | {{< ext "pg_wait_sampling" >}} | {{< ext "pg_wait_sampling" "pg_wait_sampling" >}} | 1.1.9 | sampling based statistics of wait events |
| 6280 | {{< ext "pgsentinel" >}} | {{< ext "pgsentinel" "pgsentinel" >}} | 1.2.0 | active session history |
| 6290 | {{< ext "system_stats" >}} | {{< ext "system_stats" "system_stats" >}} | 3.2 | EnterpriseDB system statistics for PostgreSQL |
| 6300 | {{< ext "meta" >}} | {{< ext "meta" "pg_meta" >}} | 0.4.0 | Normalized, friendlier system catalog for PostgreSQL |
| 6310 | {{< ext "pgnodemx" >}} | {{< ext "pgnodemx" "pgnodemx" >}} | 1.7 | Capture node OS metrics via SQL queries |
| 6320 | {{< ext "pg_proctab" >}} | {{< ext "pgnodemx" "pgnodemx" >}} | 1.7 | PostgreSQL extension to access the OS process table |
| 6330 | {{< ext "pg_sqlog" >}} | {{< ext "pg_sqlog" "pg_sqlog" >}} | 1.6 | Provide SQL interface to logs |
| 6340 | {{< ext "bgw_replstatus" >}} | {{< ext "bgw_replstatus" "bgw_replstatus" >}} | 1.0.8 | Small PostgreSQL background worker to report whether a node is a replication master or standby |
| 6350 | {{< ext "pgmeminfo" >}} | {{< ext "pgmeminfo" "pgmeminfo" >}} | 1.0.0 | show memory usage |
| 6360 | {{< ext "toastinfo" >}} | {{< ext "toastinfo" "toastinfo" >}} | 1.5 | show details on toasted datums |
| 6370 | {{< ext "explain_ui" >}} | {{< ext "explain_ui" "pg_explain_ui" >}} | 0.0.1 | easily jump into a visual plan UI for any SQL query |
| 6380 | {{< ext "pg_relusage" >}} | {{< ext "pg_relusage" "pg_relusage" >}} | 0.0.1 | Log all the queries that reference a particular column |
| 6800 | {{< ext "pagevis" >}} | {{< ext "pagevis" "pagevis" >}} | 0.1 | Visualise database pages in ascii code |
| 6810 | {{< ext "powa" >}} | {{< ext "powa" "powa" >}} | 5.0.1 | PostgreSQL Workload Analyser-core |
| 6880 | {{< ext "pg_overexplain" >}} | {{< ext "pg_overexplain" "pg_overexplain" >}} | 1.0 | Allow EXPLAIN to dump even more details |
| 6890 | {{< ext "pg_logicalinspect" >}} | {{< ext "pg_logicalinspect" "pg_logicalinspect" >}} | 1.0 | Logical decoding components inspection |
| 6900 | {{< ext "pageinspect" >}} | {{< ext "pageinspect" "pageinspect" >}} | 1.12 | inspect the contents of database pages at a low level |
| 6910 | {{< ext "pgrowlocks" >}} | {{< ext "pgrowlocks" "pgrowlocks" >}} | 1.2 | show row-level locking information |
| 6920 | {{< ext "sslinfo" >}} | {{< ext "sslinfo" "sslinfo" >}} | 1.2 | information about SSL certificates |
| 6930 | {{< ext "pg_buffercache" >}} | {{< ext "pg_buffercache" "pg_buffercache" >}} | 1.5 | examine the shared buffer cache |
| 6940 | {{< ext "pg_walinspect" >}} | {{< ext "pg_walinspect" "pg_walinspect" >}} | 1.1 | functions to inspect contents of PostgreSQL Write-Ahead Log |
| 6950 | {{< ext "pg_freespacemap" >}} | {{< ext "pg_freespacemap" "pg_freespacemap" >}} | 1.2 | examine the free space map (FSM) |
| 6960 | {{< ext "pg_visibility" >}} | {{< ext "pg_visibility" "pg_visibility" >}} | 1.2 | examine the visibility map (VM) and page-level visibility info |
| 6970 | {{< ext "pgstattuple" >}} | {{< ext "pgstattuple" "pgstattuple" >}} | 1.5 | show tuple-level statistics |
| 6980 | {{< ext "auto_explain" >}} | {{< ext "auto_explain" "auto_explain" >}} | - | Provides a means for logging execution plans of slow statements automatically |
| 6990 | {{< ext "pg_stat_statements" >}} | {{< ext "pg_stat_statements" "pg_stat_statements" >}} | 1.11 | track planning and execution statistics of all SQL statements executed |

## SEC

Auditing Logs, Enforce Passwords, Keep Secrets, TDE, SM Algorithm, Login Hooks, Log Errors, Extension White List, ...

| ID | Extension | Package | Version | Description |
|:---:|:---|:---|:---|:---|
| 7000 | {{< ext "passwordcheck_cracklib" >}} | {{< ext "passwordcheck_cracklib" "passwordcheck" >}} | 3.1.0 | Strengthen PostgreSQL user password checks with cracklib |
| 7010 | {{< ext "supautils" >}} | {{< ext "supautils" "supautils" >}} | 2.10.0 | Extension that secures a cluster on a cloud environment |
| 7020 | {{< ext "pgsodium" >}} | {{< ext "pgsodium" "pgsodium" >}} | 3.1.9 | Postgres extension for libsodium functions |
| 7030 | {{< ext "supabase_vault" >}} | {{< ext "supabase_vault" "pg_vault" >}} | 0.3.1 | Supabase Vault Extension |
| 7040 | {{< ext "pg_session_jwt" >}} | {{< ext "pg_session_jwt" "pg_session_jwt" >}} | 0.3.1 | Manage authentication sessions using JWTs |
| 7050 | {{< ext "anon" >}} | {{< ext "anon" "pg_anon" >}} | 2.3.0 | PostgreSQL Anonymizer (anon) extension |
| 7060 | {{< ext "pg_tde" >}} | {{< ext "pg_tde" "pg_tde" >}} | 1.0 | Percona pg_tde access method |
| 7070 | {{< ext "pgsmcrypto" >}} | {{< ext "pgsmcrypto" "pgsmcrypto" >}} | 0.1.0 | PostgreSQL SM Algorithm Extension |
| 7080 | {{< ext "pgaudit" >}} | {{< ext "pgaudit" "pgaudit" >}} | 17.1 | provides auditing functionality |
| 7090 | {{< ext "pgauditlogtofile" >}} | {{< ext "pgauditlogtofile" "pgauditlogtofile" >}} | 1.7.1 | pgAudit addon to redirect audit log to an independent file |
| 7100 | {{< ext "pg_auth_mon" >}} | {{< ext "pg_auth_mon" "pg_auth_mon" >}} | 3.0 | monitor connection attempts per user |
| 7110 | {{< ext "credcheck" >}} | {{< ext "credcheck" "credcheck" >}} | 3.0 | credcheck - postgresql plain text credential checker |
| 7120 | {{< ext "pgcryptokey" >}} | {{< ext "pgcryptokey" "pgcryptokey" >}} | 0.85 | cryptographic key management |
| 7130 | {{< ext "pg_jobmon" >}} | {{< ext "pg_jobmon" "pg_jobmon" >}} | 1.4.1 | Extension for logging and monitoring functions in PostgreSQL |
| 7140 | {{< ext "logerrors" >}} | {{< ext "logerrors" "logerrors" >}} | 2.1.3 | Function for collecting statistics about messages in logfile |
| 7150 | {{< ext "login_hook" >}} | {{< ext "login_hook" "login_hook" >}} | 1.7 | login_hook - hook to execute login_hook.login() at login time |
| 7160 | {{< ext "set_user" >}} | {{< ext "set_user" "set_user" >}} | 4.1.0 | similar to SET ROLE but with added logging |
| 7170 | {{< ext "pg_snakeoil" >}} | {{< ext "pg_snakeoil" "pg_snakeoil" >}} | 1.4 | The PostgreSQL Antivirus |
| 7180 | {{< ext "pgextwlist" >}} | {{< ext "pgextwlist" "pgextwlist" >}} | 1.19 | PostgreSQL Extension Whitelisting |
| 7190 | {{< ext "pg_auditor" >}} | {{< ext "pg_auditor" "pg_auditor" >}} | 0.2 | Audit data changes and provide flashback ability |
| 7200 | {{< ext "sslutils" >}} | {{< ext "sslutils" "sslutils" >}} | 1.4 | A Postgres extension for managing SSL certificates through SQL |
| 7210 | {{< ext "noset" >}} | {{< ext "noset" "pg_noset" >}} | 0.3.0 | Module for blocking SET variables for non-super users. |
| 7960 | {{< ext "sepgsql" >}} | {{< ext "sepgsql" "sepgsql" >}} | - | label-based mandatory access control (MAC) based on SELinux security policy. |
| 7970 | {{< ext "auth_delay" >}} | {{< ext "auth_delay" "auth_delay" >}} | - | pause briefly before reporting authentication failure |
| 7980 | {{< ext "pgcrypto" >}} | {{< ext "pgcrypto" "pgcrypto" >}} | 1.3 | cryptographic functions |
| 7990 | {{< ext "passwordcheck" >}} | {{< ext "passwordcheck_cracklib" "passwordcheck" >}} | - | checks user passwords and reject weak password |

## FDW

Wrappers & Multicorn for FDW Development, Access other DBMS: MySQL, Mongo, SQLite, MSSQL, Oracle, HDFS, DB2,...

| ID | Extension | Package | Version | Description |
|:---:|:---|:---|:---|:---|
| 8500 | {{< ext "wrappers" >}} | {{< ext "wrappers" "wrappers" >}} | 0.5.4 | Foreign data wrappers developed by Supabase |
| 8510 | {{< ext "multicorn" >}} | {{< ext "multicorn" "multicorn" >}} | 3.0 | Fetch foreign data in Python in your PostgreSQL server. |
| 8520 | {{< ext "odbc_fdw" >}} | {{< ext "odbc_fdw" "odbc_fdw" >}} | 0.5.1 | Foreign data wrapper for accessing remote databases using ODBC |
| 8530 | {{< ext "jdbc_fdw" >}} | {{< ext "jdbc_fdw" "jdbc_fdw" >}} | 1.2 | foreign-data wrapper for remote servers available over JDBC |
| 8540 | {{< ext "pgspider_ext" >}} | {{< ext "pgspider_ext" "pgspider_ext" >}} | 1.3.0 | foreign-data wrapper for remote PGSpider servers |
| 8600 | {{< ext "mysql_fdw" >}} | {{< ext "mysql_fdw" "mysql_fdw" >}} | 2.9.2 | Foreign data wrapper for querying a MySQL server |
| 8610 | {{< ext "oracle_fdw" >}} | {{< ext "oracle_fdw" "oracle_fdw" >}} | 2.8.0 | foreign data wrapper for Oracle access |
| 8620 | {{< ext "tds_fdw" >}} | {{< ext "tds_fdw" "tds_fdw" >}} | 2.0.4 | Foreign data wrapper for querying a TDS database (Sybase or Microsoft SQL Server) |
| 8630 | {{< ext "db2_fdw" >}} | {{< ext "db2_fdw" "db2_fdw" >}} | 7.0.0 | foreign data wrapper for DB2 access |
| 8640 | {{< ext "sqlite_fdw" >}} | {{< ext "sqlite_fdw" "sqlite_fdw" >}} | 2.5.0 | SQLite Foreign Data Wrapper |
| 8650 | {{< ext "pgbouncer_fdw" >}} | {{< ext "pgbouncer_fdw" "pgbouncer_fdw" >}} | 1.4.0 | Extension for querying PgBouncer stats from normal SQL views & running pgbouncer commands from normal SQL functions |
| 8700 | {{< ext "mongo_fdw" >}} | {{< ext "mongo_fdw" "mongo_fdw" >}} | 1.1 | foreign data wrapper for MongoDB access |
| 8710 | {{< ext "redis_fdw" >}} | {{< ext "redis_fdw" "redis_fdw" >}} | 1.0 | Foreign data wrapper for querying a Redis server |
| 8720 | {{< ext "redis" >}} | {{< ext "redis" "pg_redis_pubsub" >}} | 0.0.1 | Send redis pub/sub messages to Redis from PostgreSQL Directly |
| 8730 | {{< ext "kafka_fdw" >}} | {{< ext "kafka_fdw" "kafka_fdw" >}} | 0.0.3 | kafka Foreign Data Wrapper for CSV formatted messages |
| 8740 | {{< ext "hdfs_fdw" >}} | {{< ext "hdfs_fdw" "hdfs_fdw" >}} | 2.3.2 | foreign-data wrapper for remote hdfs servers |
| 8750 | {{< ext "firebird_fdw" >}} | {{< ext "firebird_fdw" "firebird_fdw" >}} | 1.4.0 | Foreign data wrapper for Firebird |
| 8800 | {{< ext "aws_s3" >}} | {{< ext "aws_s3" "aws_s3" >}} | 0.0.1 | aws_s3 postgres extension to import/export data from/to s3 |
| 8810 | {{< ext "log_fdw" >}} | {{< ext "log_fdw" "log_fdw" >}} | 1.4 | foreign-data wrapper for Postgres log file access |
| 8970 | {{< ext "dblink" >}} | {{< ext "dblink" "dblink" >}} | 1.2 | connect to other PostgreSQL databases from within a database |
| 8980 | {{< ext "file_fdw" >}} | {{< ext "file_fdw" "file_fdw" >}} | 1.0 | foreign-data wrapper for flat file access |
| 8990 | {{< ext "postgres_fdw" >}} | {{< ext "postgres_fdw" "postgres_fdw" >}} | 1.1 | foreign-data wrapper for remote PostgreSQL servers |

## SIM

Protocol Simulation & heterogeneous DBMS Compatibility: Oracle, MSSQL, DB2, MySQL, Memcached, and Babelfish!

| ID | Extension | Package | Version | Description |
|:---:|:---|:---|:---|:---|
| 9000 | {{< ext "documentdb" >}} | {{< ext "documentdb" "documentdb" >}} | 0.106 | API surface for DocumentDB for PostgreSQL |
| 9010 | {{< ext "documentdb_core" >}} | {{< ext "documentdb" "documentdb" >}} | 0.106 | Core API surface for DocumentDB for PostgreSQL |
| 9020 | {{< ext "documentdb_distributed" >}} | {{< ext "documentdb" "documentdb" >}} | 0.106 | Multi-Node API surface for DocumentDB |
| 9100 | {{< ext "orafce" >}} | {{< ext "orafce" "orafce" >}} | 4.14.4 | Functions and operators that emulate a subset of functions and packages from the Oracle RDBMS |
| 9110 | {{< ext "pgtt" >}} | {{< ext "pgtt" "pgtt" >}} | 4.4 | Extension to add Global Temporary Tables feature to PostgreSQL |
| 9120 | {{< ext "session_variable" >}} | {{< ext "session_variable" "session_variable" >}} | 3.4 | Registration and manipulation of session variables and constants |
| 9130 | {{< ext "pg_statement_rollback" >}} | {{< ext "pg_statement_rollback" "pg_statement_rollback" >}} | 1.4 | Server side rollback at statement level for PostgreSQL like Oracle or DB2 |
| 9240 | {{< ext "pg_dbms_metadata" >}} | {{< ext "pg_dbms_metadata" "pg_dbms_metadata" >}} | 1.0.0 | Extension to add Oracle DBMS_METADATA compatibility to PostgreSQL |
| 9250 | {{< ext "pg_dbms_lock" >}} | {{< ext "pg_dbms_lock" "pg_dbms_lock" >}} | 1.0 | Extension to add Oracle DBMS_LOCK full compatibility to PostgreSQL |
| 9260 | {{< ext "pg_dbms_job" >}} | {{< ext "pg_dbms_job" "pg_dbms_job" >}} | 1.5 | Extension to add Oracle DBMS_JOB full compatibility to PostgreSQL |
| 9300 | {{< ext "babelfishpg_common" >}} | {{< ext "babelfishpg_common" "babelfishpg_common" >}} | 3.3.3 | SQL Server Transact SQL Datatype Support |
| 9310 | {{< ext "babelfishpg_tsql" >}} | {{< ext "babelfishpg_tsql" "babelfishpg_tsql" >}} | 3.3.1 | SQL Server Transact SQL compatibility |
| 9320 | {{< ext "babelfishpg_tds" >}} | {{< ext "babelfishpg_tds" "babelfishpg_tds" >}} | 1.0.0 | SQL Server TDS protocol extension |
| 9330 | {{< ext "babelfishpg_money" >}} | {{< ext "babelfishpg_money" "babelfishpg_money" >}} | 1.1.0 | SQL Server Money Data Type |
| 9400 | {{< ext "spat" >}} | {{< ext "spat" "spat" >}} | 0.1.0a4 | Redis-like In-Memory DB Embedded in Postgres |
| 9410 | {{< ext "pgmemcache" >}} | {{< ext "pgmemcache" "pgmemcache" >}} | 2.3.0 | memcached interface |

## ETL

Logical Replication, Decoding, CDC in protobuf/JSON/Mongo format, Copy & Load & Compare Postgres Databases,...

| ID | Extension | Package | Version | Description |
|:---:|:---|:---|:---|:---|
| 9500 | {{< ext "pglogical" >}} | {{< ext "pglogical" "pglogical" >}} | 2.4.5 | PostgreSQL Logical Replication |
| 9501 | {{< ext "pglogical_origin" >}} | {{< ext "pglogical" "pglogical" >}} | 2.4.5 | Dummy extension for compatibility when upgrading from Postgres 9.4 |
| 9510 | {{< ext "pglogical_ticker" >}} | {{< ext "pglogical_ticker" "pglogical_ticker" >}} | 1.4.1 | Have an accurate view on pglogical replication delay |
| 9520 | {{< ext "pgl_ddl_deploy" >}} | {{< ext "pgl_ddl_deploy" "pgl_ddl_deploy" >}} | 2.2.1 | automated ddl deployment using pglogical |
| 9530 | {{< ext "pg_failover_slots" >}} | {{< ext "pg_failover_slots" "pg_failover_slots" >}} | 1.1.0 | PG Failover Slots extension |
| 9540 | {{< ext "db_migrator" >}} | {{< ext "db_migrator" "db_migrator" >}} | 1.0.0 | Tools to migrate other databases to PostgreSQL |
| 9550 | {{< ext "pgactive" >}} | {{< ext "pgactive" "pgactive" >}} | 2.1.6 | Active-Active Replication Extension for PostgreSQL |
| 9630 | {{< ext "wal2json" >}} | {{< ext "wal2json" "wal2json" >}} | 2.6 | Changing data capture in JSON format |
| 9640 | {{< ext "wal2mongo" >}} | {{< ext "wal2mongo" "wal2mongo" >}} | 1.0.7 | PostgreSQL logical decoding output plugin for MongoDB |
| 9650 | {{< ext "decoderbufs" >}} | {{< ext "decoderbufs" "decoderbufs" >}} | 3.2.0 | Logical decoding plugin that delivers WAL stream changes using a Protocol Buffer format |
| 9660 | {{< ext "decoder_raw" >}} | {{< ext "decoder_raw" "decoder_raw" >}} | 1.0 | Output plugin for logical replication in Raw SQL format |
| 9700 | {{< ext "mimeo" >}} | {{< ext "mimeo" "mimeo" >}} | 1.5.1 | Extension for specialized, per-table replication between PostgreSQL instances |
| 9710 | {{< ext "repmgr" >}} | {{< ext "repmgr" "repmgr" >}} | 5.5.0 | Replication manager for PostgreSQL |
| 9820 | {{< ext "pg_fact_loader" >}} | {{< ext "pg_fact_loader" "pg_fact_loader" >}} | 2.0.1 | build fact tables with Postgres |
| 9830 | {{< ext "pg_bulkload" >}} | {{< ext "pg_bulkload" "pg_bulkload" >}} | 3.1.22 | pg_bulkload is a high speed data loading utility for PostgreSQL |
| 9970 | {{< ext "test_decoding" >}} | {{< ext "test_decoding" "test_decoding" >}} | - | SQL-based test/example module for WAL logical decoding |
| 9980 | {{< ext "pgoutput" >}} | {{< ext "pgoutput" "pgoutput" >}} | - | Logical Replication output plugin |
