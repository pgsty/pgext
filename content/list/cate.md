---
title: "By Category"
weight: 100
---

PostgreSQL Extensions (444 ext in 375 pkg) categorized into 16 categories.



| {{< category "time" >}} | {{< category "gis" >}}  | {{< category "rag" >}}   | {{< category "fts" >}}  | {{< category "olap" >}} | {{< category "feat" >}} | {{< category "lang" >}} | {{< category "type" >}} |
|------------------------|-------------------------|--------------------------|-------------------------|-------------------------|-------------------------|-------------------------|-------------------------|
| {{< category "util" >}} | {{< category "func" >}} | {{< category "admin" >}} | {{< category "stat" >}} | {{< category "sec" >}}  | {{< category "fdw" >}}  | {{< category "sim" >}}  | {{< category "etl" >}}  |


## TIME

TimescaleDB, Versioning & Temporal Table, Crontab, Async & Background Job Scheduler, ...

| ID | Extension / Package | Version | Description |
|:---:|:---|:---|:---|
| 1000 | {{< alias "timescaledb" >}} | 2.24.0 | Enables scalable inserts and complex queries for time-series data |
| 1010 | {{< alias "timescaledb_toolkit" >}} | 1.22.0 | Library of analytical hyperfunctions, time-series pipelining, and other SQL utilities |
| 1020 | {{< alias "timeseries" "pg_timeseries" >}} | 0.2.0 | Convenience API for time series stack |
| 1030 | {{< alias "periods" >}} | 1.2.3 | Provide Standard SQL functionality for PERIODs and SYSTEM VERSIONING |
| 1040 | {{< alias "temporal_tables" >}} | 1.2.2 | temporal tables |
| 1050 | {{< alias "emaj" >}} | 4.7.1 | Enables fine-grained write logging and time travel on subsets of the database. |
| 1060 | {{< alias "table_version" >}} | 1.11.1 | PostgreSQL table versioning extension |
| 1070 | {{< alias "pg_cron" >}} | 1.6.7 | Job scheduler for PostgreSQL |
| 1080 | {{< alias "pg_task" >}} | 1.0.0 | execute any sql command at any specific time at background |
| 1090 | {{< alias "pg_later" >}} | 0.4.0 | Run queries now and get results later |
| 1100 | {{< alias "pg_background" >}} | 1.5 | Run SQL queries in the background |

## GIS

GeoSpatial Data Types, Operators, and Indexes, Hexagonal Indexing, OGR Data FDW, GeoIP & MobilityDB, etc...

| ID | Extension / Package | Version | Description |
|:---:|:---|:---|:---|
| 1500 | {{< alias "postgis" >}} | 3.6.1 | PostGIS geometry and geography spatial types and functions |
| 1501 | {{< alias "postgis_topology" "postgis" >}} | 3.6.1 | PostGIS topology spatial types and functions |
| 1502 | {{< alias "postgis_raster" "postgis" >}} | 3.6.1 | PostGIS raster types and functions |
| 1503 | {{< alias "postgis_sfcgal" "postgis" >}} | 3.6.1 | PostGIS SFCGAL functions |
| 1504 | {{< alias "postgis_tiger_geocoder" "postgis" >}} | 3.6.1 | PostGIS tiger geocoder and reverse geocoder |
| 1505 | {{< alias "address_standardizer" "postgis" >}} | 3.6.1 | Used to parse an address into constituent elements. Generally used to support geocoding address normalization step. |
| 1506 | {{< alias "address_standardizer_data_us" "postgis" >}} | 3.6.1 | Address Standardizer US dataset example |
| 1510 | {{< alias "pgrouting" >}} | 3.8.0 | pgRouting Extension |
| 1520 | {{< alias "pointcloud" >}} | 1.2.5 | data type for lidar point clouds |
| 1521 | {{< alias "pointcloud_postgis" "pointcloud" >}} | 1.2.5 | integration for pointcloud LIDAR data and PostGIS geometry data |
| 1530 | {{< alias "h3" "pg_h3" >}} | 4.2.3 | H3 bindings for PostgreSQL |
| 1531 | {{< alias "h3_postgis" "pg_h3" >}} | 4.2.3 | H3 PostGIS integration |
| 1540 | {{< alias "q3c" >}} | 2.0.1 | q3c sky indexing plugin |
| 1550 | {{< alias "ogr_fdw" >}} | 1.1.7 | foreign-data wrapper for GIS data access |
| 1560 | {{< alias "geoip" >}} | 0.3.0 | IP-based geolocation query |
| 1570 | {{< alias "pg_polyline" >}} | 0.0.1 | Fast Google Encoded Polyline encoding & decoding for postgres |
| 1590 | {{< alias "pg_geohash" >}} | 1.0 | Handle geohash based functionality for spatial coordinates |
| 1650 | {{< alias "mobilitydb" >}} | 1.3.0 | MobilityDB geospatial trajectory data management & analysis platform |
| 1651 | {{< alias "mobilitydb_datagen" "mobilitydb" >}} | 1.3.0 | MobilityDB random data generator functions |
| 1680 | {{< alias "tzf" "pg_tzf" >}} | 0.2.3 | Fast lookup timezone name by GPS coordinates |
| 1690 | {{< alias "earthdistance" >}} | 1.2 | calculate great-circle distances on the surface of the Earth |

## RAG

Vector Database with Ivfflat, HNSW, DiskANN Indexes, AI & ML in SQL interface, Similarity Funcs, etc...

| ID | Extension / Package | Version | Description |
|:---:|:---|:---|:---|
| 1800 | {{< alias "vector" "pgvector" >}} | 0.8.1 | vector data type and ivfflat and hnsw access methods |
| 1810 | {{< alias "vchord" >}} | 1.0.0 | Vector database plugin for Postgres, written in Rust |
| 1820 | {{< alias "vectorscale" "pgvectorscale" >}} | 0.9.0 | Advanced indexing for vector data with DiskANN |
| 1830 | {{< alias "vectorize" "pg_vectorize" >}} | 0.26.0 | The simplest way to do vector search on Postgres |
| 1840 | {{< alias "pg_similarity" >}} | 1.0 | support similarity queries |
| 1850 | {{< alias "smlar" >}} | 1.0 | Effective similarity search |
| 1860 | {{< alias "pg_summarize" >}} | 0.0.1 | Text Summarization using LLMs. Built using pgrx |
| 1870 | {{< alias "pg_tiktoken" >}} | 0.0.1 | tiktoken tokenizer for use with OpenAI models in postgres |
| 1880 | {{< alias "pg4ml" >}} | 2.0 | Machine learning framework for PostgreSQL |
| 1890 | {{< alias "pgml" >}} | 2.10.0 | Run AL/ML workloads with SQL interface |

## FTS

ElasticSearch Alternative with BM25, 2-gram/3-gram Fuzzy Search, Zhparser & Hunspell Segregation Dicts, etc...

| ID | Extension / Package | Version | Description |
|:---:|:---|:---|:---|
| 2100 | {{< alias "pg_search" >}} | 0.21.2 | Full text search for PostgreSQL using BM25 |
| 2110 | {{< alias "pgroonga" >}} | 4.0.4 | Use Groonga as index, fast full text search platform for all languages! |
| 2111 | {{< alias "pgroonga_database" "pgroonga" >}} | 4.0.4 | PGroonga database management module |
| 2120 | {{< alias "pg_bigm" >}} | 1.2 | create 2-gram (bigram) index for faster full text search. |
| 2130 | {{< alias "zhparser" >}} | 2.3 | a parser for full-text search of Chinese |
| 2140 | {{< alias "pg_bestmatch" >}} | 0.0.2 | Generate BM25 sparse vector inside PostgreSQL |
| 2150 | {{< alias "vchord_bm25" >}} | 0.3.0 | A postgresql extension for bm25 ranking algorithm |
| 2160 | {{< alias "pg_tokenizer" >}} | 0.1.1 | Tokenizers for full-text search |
| 2170 | {{< alias "biscuit" "pg_biscuit" >}} | 2.2.2 | IAM-LIKE pattern matching with bitmap indexing |
| 2180 | {{< alias "pg_textsearch" >}} | 0.4.0 | Full-text search with BM25 ranking |
| 2270 | {{< alias "hunspell_cs_cz" >}} | 1.0 | Czech Hunspell Dictionary |
| 2271 | {{< alias "hunspell_de_de" >}} | 1.0 | German Hunspell Dictionary |
| 2272 | {{< alias "hunspell_en_us" >}} | 1.0 | en_US Hunspell Dictionary |
| 2273 | {{< alias "hunspell_fr" >}} | 1.0 | French Hunspell Dictionary |
| 2274 | {{< alias "hunspell_ne_np" >}} | 1.0 | Nepali Hunspell Dictionary |
| 2275 | {{< alias "hunspell_nl_nl" >}} | 1.0 | Dutch Hunspell Dictionary |
| 2276 | {{< alias "hunspell_nn_no" >}} | 1.0 | Norwegian (norsk) Hunspell Dictionary |
| 2277 | {{< alias "hunspell_pt_pt" >}} | 1.0 | Portuguese Hunspell Dictionary |
| 2278 | {{< alias "hunspell_ru_ru" >}} | 1.0 | Russian Hunspell Dictionary |
| 2279 | {{< alias "hunspell_ru_ru_aot" >}} | 1.0 | Russian Hunspell Dictionary (from AOT.ru group) |
| 2380 | {{< alias "fuzzystrmatch" >}} | 1.2 | determine similarities and distance between strings |
| 2390 | {{< alias "pg_trgm" >}} | 1.6 | text similarity measurement and index searching based on trigrams |

## OLAP

DuckDB Integration with FDW & PG Lakehouse, Access Parquet from File/S3, Sharding with Citus/Partman/PlProxy, ...

| ID | Extension / Package | Version | Description |
|:---:|:---|:---|:---|
| 2400 | {{< alias "citus" >}} | 14.0.0 | Distributed PostgreSQL as an extension |
| 2401 | {{< alias "citus_columnar" "citus" >}} | 14.0.0 | Citus columnar storage engine |
| 2410 | {{< alias "columnar" "hydra" >}} | 1.1.2 | Hydra Columnar extension |
| 2420 | {{< alias "pg_analytics" >}} | 0.3.7 | Postgres for analytics, powered by DuckDB |
| 2430 | {{< alias "pg_duckdb" >}} | 1.1.1 | DuckDB Embedded in Postgres |
| 2440 | {{< alias "pg_mooncake" >}} | 0.2.0 | Columnstore Table in Postgres |
| 2460 | {{< alias "pg_clickhouse" >}} | 0.1.2 | Interfaces to query ClickHouse databases from PostgreSQL |
| 2470 | {{< alias "duckdb_fdw" >}} | 1.1.2 | DuckDB Foreign Data Wrapper |
| 2480 | {{< alias "pg_parquet" >}} | 0.5.1 | copy data between Postgres and Parquet |
| 2500 | {{< alias "pg_fkpart" >}} | 1.7.0 | Table partitioning by foreign key utility |
| 2510 | {{< alias "pg_partman" >}} | 5.4.0 | Extension to manage partitioned tables by time or ID |
| 2520 | {{< alias "plproxy" >}} | 2.11.0 | Database partitioning implemented as procedural language |
| 2530 | {{< alias "pg_strom" >}} | 6.0 | PG-Strom - big-data processing acceleration using GPU and NVME |
| 2590 | {{< alias "tablefunc" >}} | 1.0 | functions that manipulate whole tables, including crosstab |

## FEAT

OpenCypher with AGE, GraphQL, JsonSchema, Hints & Hypo Index, HLL, Rum, IVM, ChemRDKit, and Message Queues,...

| ID | Extension / Package | Version | Description |
|:---:|:---|:---|:---|
| 2730 | {{< alias "age" >}} | 1.6.0 | AGE graph database extension |
| 2740 | {{< alias "hll" >}} | 2.19 | type for storing hyperloglog data |
| 2750 | {{< alias "rum" >}} | 1.3.15 | RUM index access method |
| 2760 | {{< alias "pg_ai_query" >}} | 0.1.1 | AI-powered SQL query generation for PostgreSQL |
| 2780 | {{< alias "pg_ttl_index" >}} | 2.0.0 | Automatic data expiration with TTL indexes |
| 2790 | {{< alias "pg_graphql" >}} | 1.5.12 | Add in-database GraphQL support |
| 2800 | {{< alias "pg_jsonschema" >}} | 0.3.3 | PostgreSQL extension providing JSON Schema validation |
| 2810 | {{< alias "jsquery" >}} | 1.2 | data type for jsonb inspection |
| 2820 | {{< alias "pg_hint_plan" >}} | 1.8.0 | Give PostgreSQL ability to manually force some decisions in execution plans. |
| 2830 | {{< alias "hypopg" >}} | 1.4.2 | Hypothetical indexes for PostgreSQL |
| 2840 | {{< alias "index_advisor" >}} | 0.2.0 | Query index advisor |
| 2850 | {{< alias "plan_filter" "pg_plan_filter" >}} | 0.0.1 | filter statements by their execution plans. |
| 2860 | {{< alias "imgsmlr" >}} | 1.0 | Image similarity with haar |
| 2870 | {{< alias "pg_ivm" >}} | 1.13 | incremental view maintenance on PostgreSQL |
| 2880 | {{< alias "pg_incremental" >}} | 1.2.0 | Incremental Processing by Crunchy Data |
| 2890 | {{< alias "pgmq" >}} | 1.8.1 | A lightweight message queue. Like AWS SQS and RSMQ but on Postgres. |
| 2900 | {{< alias "pgq" >}} | 3.5.1 | Generic queue for PostgreSQL |
| 2910 | {{< alias "orioledb" >}} | 1.5 | OrioleDB, the next generation transactional engine |
| 2920 | {{< alias "pg_cardano" >}} | 1.1.1 | A suite of Cardano-related tools |
| 2930 | {{< alias "rdkit" >}} | 202503.1 | Cheminformatics functionality for PostgreSQL. |
| 2940 | {{< alias "omni" "omnigres" >}} | 0.2.14 | Advanced adapter for Postgres extensions |
| 2941 | {{< alias "omni_auth" "omnigres" >}} | 0.1.3 | Basic session management |
| 2942 | {{< alias "omni_aws" "omnigres" >}} | 0.1.2 | Amazon Web Services APIs (S3) |
| 2943 | {{< alias "omni_cloudevents" "omnigres" >}} | 0.1.0 | CloudEvents support |
| 2944 | {{< alias "omni_containers" "omnigres" >}} | 0.2.0 | Docker container management |
| 2945 | {{< alias "omni_credentials" "omnigres" >}} | 0.2.0 | Application credential management |
| 2946 | {{< alias "omni_csv" >}} | 0.1.1 | CSV toolkit |
| 2947 | {{< alias "omni_datasets" >}} | 0.1.0 | Dataset provisioning |
| 2948 | {{< alias "omni_email" "omnigres" >}} | 0.1.0 | E-mail framework |
| 2949 | {{< alias "omni_http" "omnigres" >}} | 0.1.0 | Basic HTTP types |
| 2950 | {{< alias "omni_httpc" "omnigres" >}} | 0.1.10 | HTTP client |
| 2951 | {{< alias "omni_httpd" "omnigres" >}} | 0.4.11 | HTTP server |
| 2952 | {{< alias "omni_id" "omnigres" >}} | 0.4.3 | Identity types |
| 2953 | {{< alias "omni_json" "omnigres" >}} | 0.1.1 | JSON toolkit |
| 2954 | {{< alias "omni_kube" "omnigres" >}} | 0.4.2 | Kubernetes (k8s) integration |
| 2955 | {{< alias "omni_ledger" "omnigres" >}} | 0.1.3 | Financial ledger |
| 2956 | {{< alias "omni_manifest" "omnigres" >}} | 0.1.2 | Package installation manifests |
| 2957 | {{< alias "omni_mimetypes" "omnigres" >}} | 0.1.0 | MIME types |
| 2958 | {{< alias "omni_os" "omnigres" >}} | 0.1.1 | Operating system integration |
| 2959 | {{< alias "omni_polyfill" "omnigres" >}} | 0.2.2 | Postgres API polyfills |
| 2960 | {{< alias "omni_python" "omnigres" >}} | 0.1.1 | First-class Python support |
| 2961 | {{< alias "omni_regex" "omnigres" >}} | 0.1.0 | PCRE-compatible regular expressions |
| 2962 | {{< alias "omni_rest" "omnigres" >}} | 0.1.1 | REST API toolkit (with PostgREST support) |
| 2963 | {{< alias "omni_schema" "omnigres" >}} | 0.3.0 | Advanced schema management tooling |
| 2964 | {{< alias "omni_seq" "omnigres" >}} | 0.1.1 | Distributed integer sequences |
| 2965 | {{< alias "omni_service" "omnigres" >}} | 0.1.0 | Service management |
| 2966 | {{< alias "omni_session" "omnigres" >}} | 0.2.0 | Session management |
| 2967 | {{< alias "omni_shmem" >}} | 0.1.0 | Shared Memory Management |
| 2968 | {{< alias "omni_sql" "omnigres" >}} | 0.5.3 | Programmatic SQL manipulation |
| 2969 | {{< alias "omni_sqlite" "omnigres" >}} | 0.2.2 | Embedded SQLite |
| 2970 | {{< alias "omni_test" "omnigres" >}} | 0.4.0 | Testing framework |
| 2971 | {{< alias "omni_txn" "omnigres" >}} | 0.5.0 | Transaction management |
| 2972 | {{< alias "omni_types" "omnigres" >}} | 0.3.6 | Advanced types |
| 2973 | {{< alias "omni_var" "omnigres" >}} | 0.3.0 | Scoped variables |
| 2974 | {{< alias "omni_vfs" "omnigres" >}} | 0.2.2 | Virtual File System |
| 2975 | {{< alias "omni_vfs_types_v1" "omnigres" >}} | 0.1.0 | Virtual File System types (v1) |
| 2976 | {{< alias "omni_web" "omnigres" >}} | 0.3.0 | Common web stack primitives |
| 2977 | {{< alias "omni_worker" "omnigres" >}} | 0.2.1 | Generalized worker pool |
| 2978 | {{< alias "omni_xml" "omnigres" >}} | 0.1.2 | XML toolkit |
| 2979 | {{< alias "omni_yaml" "omnigres" >}} | 0.1.0 | YAML toolkit |
| 2990 | {{< alias "bloom" >}} | 1.0 | bloom access method - signature file based index |

## LANG

Develop, Test, Package, and Deliver Stored Procedures written in various PL/Languages: Java, Js, Lua, R, Sh, PRQL, ...

| ID | Extension / Package | Version | Description |
|:---:|:---|:---|:---|
| 3000 | {{< alias "pg_tle" >}} | 1.5.2 | Trusted Language Extensions for PostgreSQL |
| 3010 | {{< alias "plv8" >}} | 3.2.4 | PL/JavaScript (v8) trusted procedural language |
| 3011 | {{< alias "pljs" >}} | 1.0.4 | PL/JS trusted procedural language |
| 3020 | {{< alias "pllua" >}} | 2.0.12 | Lua as a procedural language |
| 3021 | {{< alias "hstore_pllua" "pllua" >}} | 2.0.12 | Hstore transform for Lua |
| 3030 | {{< alias "plluau" "pllua" >}} | 2.0.12 | Lua as an untrusted procedural language |
| 3031 | {{< alias "hstore_plluau" "pllua" >}} | 2.0.12 | Hstore transform for untrusted Lua |
| 3040 | {{< alias "plprql" >}} | 18.0.0 | Use PRQL in PostgreSQL - Pipelined Relational Query Language |
| 3050 | {{< alias "pldbgapi" "pldebugger" >}} | 1.9 | server-side support for debugging PL/pgSQL functions |
| 3060 | {{< alias "plpgsql_check" >}} | 2.8.5 | extended check for plpgsql functions |
| 3070 | {{< alias "plprofiler" >}} | 4.2.5 | server-side support for profiling PL/pgSQL functions |
| 3080 | {{< alias "plsh" >}} | 1.20220917 | PL/sh procedural language |
| 3090 | {{< alias "pljava" >}} | 1.6.10 | PL/Java procedural language |
| 3100 | {{< alias "plr" >}} | 8.4.8 | load R interpreter and execute R script from within a database |
| 3110 | {{< alias "plxslt" >}} | 0.20140221 | XSLT procedural language for PostgreSQL |
| 3200 | {{< alias "pgtap" >}} | 1.3.4 | Unit testing for PostgreSQL |
| 3210 | {{< alias "faker" >}} | 0.5.3 | Wrapper for the Faker Python library |
| 3220 | {{< alias "dbt2" >}} | 0.61.7 | OSDL-DBT-2 test kit |
| 3240 | {{< alias "pltcl" >}} | 1.0 | PL/Tcl procedural language |
| 3250 | {{< alias "pltclu" "pltcl" >}} | 1.0 | PL/TclU untrusted procedural language |
| 3260 | {{< alias "plperl" >}} | 1.0 | PL/Perl procedural language |
| 3261 | {{< alias "bool_plperl" "plperl" >}} | 1.0 | transform between bool and plperl |
| 3262 | {{< alias "hstore_plperl" "plperl" >}} | 1.0 | transform between hstore and plperl |
| 3263 | {{< alias "jsonb_plperl" "plperl" >}} | 1.0 | transform between jsonb and plperl |
| 3270 | {{< alias "plperlu" >}} | 1.0 | PL/PerlU untrusted procedural language |
| 3271 | {{< alias "bool_plperlu" "plperlu" >}} | 1.0 | transform between bool and plperlu |
| 3272 | {{< alias "jsonb_plperlu" "plperlu" >}} | 1.0 | transform between jsonb and plperlu |
| 3273 | {{< alias "hstore_plperlu" "plperlu" >}} | 1.0 | transform between hstore and plperlu |
| 3280 | {{< alias "plpgsql" >}} | 1.0 | PL/pgSQL procedural language |
| 3290 | {{< alias "plpython3u" >}} | 1.0 | PL/Python3U untrusted procedural language |
| 3291 | {{< alias "jsonb_plpython3u" "plpython3u" >}} | 1.0 | transform between jsonb and plpython3u |
| 3292 | {{< alias "ltree_plpython3u" "plpython3u" >}} | 1.0 | transform between ltree and plpython3u |
| 3293 | {{< alias "hstore_plpython3u" "plpython3u" >}} | 1.0 | transform between hstore and plpython3u |

## TYPE

Dedicate New Data Types Like: prefix, semver, uint, SIUnit, RoaringBitmap, Rational, Sphere, Hash, RRule, and more...

| ID | Extension / Package | Version | Description |
|:---:|:---|:---|:---|
| 3500 | {{< alias "prefix" "pg_prefix" >}} | 1.2.10 | Prefix Range module for PostgreSQL |
| 3510 | {{< alias "semver" "pg_semver" >}} | 0.41.0 | Semantic version data type |
| 3520 | {{< alias "unit" "pgunit" >}} | 7.10 | SI units extension |
| 3530 | {{< alias "pgpdf" >}} | 0.1.0 | PDF type with meta admin & Full-Text Search |
| 3540 | {{< alias "pglite_fusion" >}} | 0.0.6 | Embed an SQLite database in your PostgreSQL table |
| 3550 | {{< alias "md5hash" >}} | 1.0.1 | type for storing 128-bit binary data inline |
| 3560 | {{< alias "asn1oid" >}} | 1.6 | asn1oid extension |
| 3570 | {{< alias "roaringbitmap" "pg_roaringbitmap" >}} | 1.1.0 | support for Roaring Bitmaps |
| 3580 | {{< alias "pgfaceting" >}} | 0.2.0 | fast faceting queries using an inverted index |
| 3590 | {{< alias "pg_sphere" "pgsphere" >}} | 1.5.2 | spherical objects with useful functions, operators and index support |
| 3600 | {{< alias "country" "pg_country" >}} | 0.0.3 | Country data type, ISO 3166-1 |
| 3610 | {{< alias "pg_xenophile" >}} | 0.8.3 | More than the bare necessities for PostgreSQL i18n and l10n. |
| 3611 | {{< alias "l10n_table_dependent_extension" "pg_xenophile" >}} | 0.8.3 | PostgreSQL l10n toolbox |
| 3620 | {{< alias "currency" "pg_currency" >}} | 0.0.3 | Custom PostgreSQL currency type in 1Byte |
| 3630 | {{< alias "collection" "pgcollection" >}} | 1.1.0 | Memory optimized data type to be used inside of plpglsql func |
| 3700 | {{< alias "pgmp" >}} | 1.0.5 | Multiple Precision Arithmetic extension |
| 3710 | {{< alias "numeral" >}} | 1.3 | numeral datatypes extension |
| 3720 | {{< alias "pg_rational" >}} | 0.0.2 | bigint fractions |
| 3730 | {{< alias "uint" "pguint" >}} | 1.20250815 | unsigned integer types |
| 3740 | {{< alias "uint128" "pg_uint128" >}} | 1.1.1 | Native uint128 type |
| 3750 | {{< alias "hashtypes" >}} | 0.1.5 | sha1, md5 and other data types for PostgreSQL |
| 3820 | {{< alias "ip4r" >}} | 2.4.2 | IPv4/v6 and IPv4/v6 range index type for PostgreSQL |
| 3830 | {{< alias "pg_duration" >}} | 1.0.2 | data type for representing durations |
| 3840 | {{< alias "uri" "pg_uri" >}} | 1.20151224 | URI Data type for PostgreSQL |
| 3850 | {{< alias "emailaddr" "pg_emailaddr" >}} | 0 | Email address type for PostgreSQL |
| 3860 | {{< alias "acl" "pg_acl" >}} | 1.0.4 | ACL Data type |
| 3870 | {{< alias "debversion" >}} | 1.2.0 | Debian version number data type |
| 3880 | {{< alias "pg_rrule" >}} | 0.3.0 | RRULE field type for PostgreSQL |
| 3890 | {{< alias "timestamp9" >}} | 1.4.0 | timestamp nanosecond resolution |
| 3920 | {{< alias "chkpass" >}} | 1.0 | data type for auto-encrypted passwords |
| 3930 | {{< alias "isn" >}} | 1.2 | data types for international product numbering standards |
| 3940 | {{< alias "seg" >}} | 1.4 | data type for representing line segments or floating-point intervals |
| 3950 | {{< alias "cube" >}} | 1.5 | data type for multidimensional cubes |
| 3960 | {{< alias "ltree" >}} | 1.3 | data type for hierarchical tree-like structures |
| 3970 | {{< alias "hstore" >}} | 1.8 | data type for storing sets of (key, value) pairs |
| 3980 | {{< alias "citext" >}} | 1.6 | data type for case-insensitive character strings |
| 3990 | {{< alias "xml2" >}} | 1.1 | XPath querying and XSLT |

## UTIL

Utilities such as send http request, perform gzip/zstd compress, send mails, Regex, ICU, encoding, docs, Encryption,...

| ID | Extension / Package | Version | Description |
|:---:|:---|:---|:---|
| 4010 | {{< alias "gzip" "pg_gzip" >}} | 1.0.0 | gzip and gunzip functions. |
| 4020 | {{< alias "bzip" "pg_bzip" >}} | 1.0.0 | Bzip compression and decompression |
| 4030 | {{< alias "zstd" "pg_zstd" >}} | 1.1.2 | Zstandard compression algorithm implementation in PostgreSQL |
| 4070 | {{< alias "http" "pg_http" >}} | 1.7.0 | HTTP client for PostgreSQL, allows web page retrieval inside the database. |
| 4080 | {{< alias "pg_net" >}} | 0.20.0 | Async HTTP Requests |
| 4090 | {{< alias "pg_curl" >}} | 2.4.5 | Run curl actions for data transfer in URL syntax |
| 4100 | {{< alias "pg_retry" >}} | 1.0.0 | Retry SQL statements on transient errors with exponential backoff |
| 4150 | {{< alias "pgjq" >}} | 0.1.0 | Use jq in Postgres |
| 4160 | {{< alias "pgjwt" >}} | 0.2.0 | JSON Web Token API for Postgresql |
| 4170 | {{< alias "pg_smtp_client" >}} | 0.2.1 | PostgreSQL extension to send email using SMTP |
| 4180 | {{< alias "pg_html5_email_address" >}} | 1.2.3 | PostgreSQL email validation that is consistent with the HTML5 spec |
| 4190 | {{< alias "url_encode" >}} | 1.2.5 | url_encode, url_decode functions |
| 4200 | {{< alias "pgsql_tweaks" >}} | 1.0.2 | Some functions and views for daily usage |
| 4220 | {{< alias "pg_extra_time" >}} | 2.0.0 | Some date time functions and operators that, |
| 4230 | {{< alias "pgpcre" >}} | 0.20190509 | Perl Compatible Regular Expression functions |
| 4240 | {{< alias "icu_ext" >}} | 1.10.0 | Access ICU functions |
| 4250 | {{< alias "pgqr" >}} | 1.0 | QR Code generator from PostgreSQL |
| 4260 | {{< alias "pg_protobuf" >}} | 1.0 | Protobuf support for PostgreSQL |
| 4270 | {{< alias "envvar" "pg_envvar" >}} | 1.0.1 | Fetch the value of an environment variable |
| 4280 | {{< alias "floatfile" >}} | 1.3.1 | Simple file storage for arrays of floats |
| 4290 | {{< alias "pg_render" >}} | 0.1.3 | Render HTML in SQL |
| 4300 | {{< alias "pg_readme" >}} | 0.7.0 | Generate a README.md document for a database extension or schema |
| 4301 | {{< alias "pg_readme_test_extension" "pg_readme" >}} | 0.7.0 | Test generating a README.md document for extension or schema |
| 4310 | {{< alias "ddl_historization" >}} | 0.0.7 | Historize the ddl changes inside PostgreSQL database |
| 4320 | {{< alias "data_historization" >}} | 1.1.0 | PLPGSQL Script to historize data in partitionned table |
| 4330 | {{< alias "schedoc" "pg_schedoc" >}} | 0.0.1 | Cross documentation between Django and DBT projects |
| 4400 | {{< alias "hashlib" "pg_hashlib" >}} | 1.1 | Stable hash functions for Postgres |
| 4430 | {{< alias "xxhash" "pg_xxhash" >}} | 0.0.1 | xxhash functions for PostgreSQL |
| 4440 | {{< alias "shacrypt" >}} | 1.1 | Implements SHA256-CRYPT and SHA512-CRYPT password encryption schemes |
| 4450 | {{< alias "cryptint" >}} | 1.0.0 | Encryption functions for int and bigint values |
| 4460 | {{< alias "pguecc" "pg_ecdsa" >}} | 1.0 | uECC bindings for Postgres |
| 4470 | {{< alias "sparql" "pgsparql" >}} | 1.0 | Query SPARQL datasource with SQL |

## FUNC

Function such as id generator, aggregations, sketches, vector functions, mathematical functions and digest functions...

| ID | Extension / Package | Version | Description |
|:---:|:---|:---|:---|
| 4500 | {{< alias "pg_idkit" >}} | 0.4.0 | multi-tool for generating new/niche universally unique identifiers (ex. UUIDv6, ULID, KSUID) |
| 4510 | {{< alias "pgx_ulid" >}} | 0.2.2 | ulid type and methods |
| 4540 | {{< alias "pg_uuidv7" >}} | 1.7.0 | Create UUIDv7 values in postgres |
| 4550 | {{< alias "permuteseq" >}} | 1.2.2 | Pseudo-randomly permute sequences with a format-preserving encryption on elements |
| 4560 | {{< alias "pg_hashids" >}} | 1.3 | Short unique id generator for PostgreSQL, using hashids |
| 4570 | {{< alias "sequential_uuids" >}} | 1.0.3 | generator of sequential UUIDs |
| 4580 | {{< alias "typeid" "pg_typeid" >}} | 0.3.0 | Allows to use TypeIDs in Postgres natively |
| 4600 | {{< alias "topn" >}} | 2.7.0 | type for top-n JSONB |
| 4610 | {{< alias "quantile" >}} | 1.1.8 | Quantile aggregation function |
| 4620 | {{< alias "lower_quantile" >}} | 1.0.3 | Lower quantile aggregate function |
| 4630 | {{< alias "count_distinct" >}} | 3.0.2 | An alternative to COUNT(DISTINCT …) aggregate, usable with HashAggregate |
| 4640 | {{< alias "omnisketch" >}} | 1.0.2 | data structure for on-line agg of data into approximate sketch |
| 4650 | {{< alias "ddsketch" >}} | 1.0.1 | Provides ddsketch aggregate function |
| 4660 | {{< alias "vasco" >}} | 0.1.0 | discover hidden correlations in your data with MIC |
| 4670 | {{< alias "xicor" "pgxicor" >}} | 0.1.0 | XI Correlation Coefficient in Postgres |
| 4680 | {{< alias "weighted_statistics" "pg_weighted_statistics" >}} | 1.0.0 | High-performance weighted statistics functions for sparse data |
| 4700 | {{< alias "tdigest" >}} | 1.4.3 | Provides tdigest aggregate function. |
| 4710 | {{< alias "first_last_agg" >}} | 0.1.4 | first() and last() aggregate functions |
| 4720 | {{< alias "extra_window_functions" >}} | 1.0 | Extra Window Functions for PostgreSQL |
| 4730 | {{< alias "floatvec" >}} | 1.1.1 | Math for vectors (arrays) of numbers |
| 4740 | {{< alias "aggs_for_vecs" >}} | 1.4.0 | Aggregate functions for array inputs |
| 4750 | {{< alias "aggs_for_arrays" >}} | 1.3.3 | Various functions for computing statistics on arrays of numbers |
| 4760 | {{< alias "pg_csv" >}} | 1.0.1 | Flexible CSV processing for Postgres |
| 4770 | {{< alias "arraymath" "pg_arraymath" >}} | 1.1 | Array math and operators that work element by element on the contents of arrays |
| 4780 | {{< alias "pg_math" >}} | 1.0 | GSL statistical functions for postgresql |
| 4790 | {{< alias "random" "pg_random" >}} | 2.0.0 | random data generator |
| 4800 | {{< alias "base36" "pg_base36" >}} | 1.0.0 | Integer Base36 types |
| 4810 | {{< alias "base62" "pg_base62" >}} | 0.0.1 | Base62 extension for PostgreSQL |
| 4830 | {{< alias "pg_base58" >}} | 0.0.1 | Base58 Encoder/Decoder Extension for PostgreSQL |
| 4840 | {{< alias "financial" "pg_financial" >}} | 1.0.1 | Financial aggregate functions |
| 4850 | {{< alias "convert" "pg_convert" >}} | 0.1.0 | conversion functions for spatial, routing and other specialized uses |
| 4880 | {{< alias "refint" >}} | 1.0 | functions for implementing referential integrity (obsolete) |
| 4881 | {{< alias "autoinc" >}} | 1.0 | functions for autoincrementing fields |
| 4882 | {{< alias "insert_username" >}} | 1.0 | functions for tracking who changed a table |
| 4883 | {{< alias "moddatetime" >}} | 1.0 | functions for tracking last modification time |
| 4890 | {{< alias "tsm_system_time" >}} | 1.0 | TABLESAMPLE method which accepts time in milliseconds as a limit |
| 4900 | {{< alias "dict_xsyn" >}} | 1.0 | text search dictionary template for extended synonym processing |
| 4910 | {{< alias "tsm_system_rows" >}} | 1.0 | TABLESAMPLE method which accepts number of rows as a limit |
| 4920 | {{< alias "tcn" >}} | 1.0 | Triggered change notifications |
| 4930 | {{< alias "uuid-ossp" >}} | 1.1 | generate universally unique identifiers (UUIDs) |
| 4940 | {{< alias "btree_gist" >}} | 1.7 | support for indexing common datatypes in GiST |
| 4950 | {{< alias "btree_gin" >}} | 1.3 | support for indexing common datatypes in GIN |
| 4960 | {{< alias "intarray" >}} | 1.5 | functions, operators, and index support for 1-D arrays of integers |
| 4970 | {{< alias "intagg" >}} | 1.1 | integer aggregator and enumerator (obsolete) |
| 4980 | {{< alias "dict_int" >}} | 1.0 | text search dictionary template for integers |
| 4990 | {{< alias "unaccent" >}} | 1.1 | text search dictionary that removes accents |

## ADMIN

Utilities for Bloat Control, DirtyRead, BufferInspect, DDL Generate, ChecksumVerify, Permission, Priority, Catalog,...

| ID | Extension / Package | Version | Description |
|:---:|:---|:---|:---|
| 5010 | {{< alias "pg_repack" >}} | 1.5.3 | Reorganize tables in PostgreSQL databases with minimal locks |
| 5020 | {{< alias "pg_rewrite" >}} | 2.0.0 | Tool allows read write to the table during the rewriting |
| 5040 | {{< alias "pg_squeeze" >}} | 1.9.1 | A tool to remove unused space from a relation. |
| 5050 | {{< alias "pg_dirtyread" >}} | 2.7 | Read dead but unvacuumed rows from table |
| 5060 | {{< alias "pgfincore" >}} | 1.3.1 | examine and manage the os buffer cache |
| 5070 | {{< alias "pg_cooldown" >}} | 0.1 | remove buffered pages for specific relations |
| 5080 | {{< alias "ddlx" "pg_ddlx" >}} | 0.30 | DDL eXtractor functions |
| 5090 | {{< alias "pglinter" >}} | 1.0.1 | PostgreSQL Linting and Analysis Extension |
| 5100 | {{< alias "prioritize" "pg_prioritize" >}} | 1.0.4 | get and set the priority of PostgreSQL backends |
| 5110 | {{< alias "pg_checksums" >}} | 1.3 | Activate/deactivate/verify checksums in offline Postgres clusters |
| 5120 | {{< alias "pg_readonly" >}} | 1.0.3 | cluster database read only |
| 5130 | {{< alias "pgdd" >}} | 0.6.1 | Introspect pg data dictionary via standard SQL |
| 5140 | {{< alias "pg_permissions" >}} | 1.4 | view object permissions and compare them with the desired state |
| 5150 | {{< alias "pgautofailover" >}} | 2.2 | pg_auto_failover |
| 5160 | {{< alias "pg_catcheck" >}} | 1.6.0 | Diagnosing system catalog corruption |
| 5170 | {{< alias "pre_prepare" "preprepare" >}} | 0.9 | Pre Prepare your Statement server side |
| 5180 | {{< alias "pg_upless" >}} | 0.0.3 | Detect Useless UPDATE |
| 5190 | {{< alias "pgcozy" >}} | 1.0 | Pre-warming shared buffers according to previous pg_buffercache snapshots for PostgreSQL. |
| 5200 | {{< alias "pg_orphaned" >}} | 1.0 | Deal with orphaned files |
| 5210 | {{< alias "pg_crash" >}} | 1.0 | Send random signals to random processes |
| 5220 | {{< alias "pg_cheat_funcs" >}} | 1.0 | Provides cheat (but useful) functions |
| 5230 | {{< alias "fio" "pg_fio" >}} | 1.0 | PostgreSQL File I/O Functions |
| 5810 | {{< alias "pg_savior" >}} | 0.0.1 | Postgres extension to save OOPS mistakes |
| 5820 | {{< alias "safeupdate" >}} | 1.5 | Require criteria for UPDATE and DELETE |
| 5830 | {{< alias "pg_drop_events" >}} | 0.1.0 | logs transaction ids of drop table, drop column, drop materialized view statements |
| 5840 | {{< alias "table_log" >}} | 0.6.4 | record table modification logs and PITR for table/row |
| 5880 | {{< alias "pgagent" >}} | 4.2.3 | A PostgreSQL job scheduler |
| 5890 | {{< alias "pg_prewarm" >}} | 1.2 | prewarm relation data |
| 5900 | {{< alias "pgpool_adm" "pgpool" >}} | 4.7.0 | Administrative functions for pgPool |
| 5910 | {{< alias "pgpool_recovery" "pgpool" >}} | 4.7.0 | recovery functions for pgpool-II for V4.3 |
| 5920 | {{< alias "pgpool_regclass" "pgpool" >}} | 4.7.0 | replacement for regclass |
| 5930 | {{< alias "lo" >}} | 1.1 | Large Object maintenance |
| 5940 | {{< alias "basic_archive" >}} | - | an example of an archive module |
| 5950 | {{< alias "basebackup_to_shell" >}} | - | adds a custom basebackup target called shell |
| 5960 | {{< alias "old_snapshot" >}} | 1.0 | utilities in support of old_snapshot_threshold |
| 5970 | {{< alias "adminpack" >}} | 2.1 | administrative functions for PostgreSQL |
| 5980 | {{< alias "amcheck" >}} | 1.4 | functions for verifying relation integrity |
| 5990 | {{< alias "pg_surgery" >}} | 1.0 | extension to perform surgery on a damaged relation |

## STAT

Observability Catalogs, Monitoring Metrics & Views, Statistics, Query Plans, WaitSampling, SlowLogs, and etc...

| ID | Extension / Package | Version | Description |
|:---:|:---|:---|:---|
| 6000 | {{< alias "pg_profile" >}} | 4.11 | PostgreSQL load profile repository and report builder |
| 6010 | {{< alias "pg_tracing" >}} | 0.1.3 | Distributed Tracing for PostgreSQL |
| 6210 | {{< alias "pg_show_plans" >}} | 2.1.7 | show query plans of all currently running SQL statements |
| 6220 | {{< alias "pg_stat_kcache" >}} | 2.3.1 | Kernel statistics gathering |
| 6230 | {{< alias "pg_stat_monitor" >}} | 2.3.1 | The pg_stat_monitor is a PostgreSQL Query Performance Monitoring tool, based on PostgreSQL contrib module pg_stat_statements. pg_stat_monitor provides aggregated statistics, client information, plan details including plan, and histogram information. |
| 6240 | {{< alias "pg_qualstats" >}} | 2.1.3 | An extension collecting statistics about quals |
| 6250 | {{< alias "pg_store_plans" >}} | 1.9 | track plan statistics of all SQL statements executed |
| 6260 | {{< alias "pg_track_settings" >}} | 2.1.2 | Track settings changes |
| 6270 | {{< alias "pg_wait_sampling" >}} | 1.1.9 | sampling based statistics of wait events |
| 6280 | {{< alias "pgsentinel" >}} | 1.3.1 | active session history |
| 6290 | {{< alias "system_stats" >}} | 3.2 | EnterpriseDB system statistics for PostgreSQL |
| 6300 | {{< alias "meta" "pg_meta" >}} | 0.4.0 | Normalized, friendlier system catalog for PostgreSQL |
| 6310 | {{< alias "pgnodemx" >}} | 1.7 | Capture node OS metrics via SQL queries |
| 6320 | {{< alias "pg_proctab" "pgnodemx" >}} | 1.7 | PostgreSQL extension to access the OS process table |
| 6330 | {{< alias "pg_sqlog" >}} | 1.6 | Provide SQL interface to logs |
| 6340 | {{< alias "bgw_replstatus" >}} | 1.0.8 | Small PostgreSQL background worker to report whether a node is a replication master or standby |
| 6350 | {{< alias "pgmeminfo" >}} | 1.0.0 | show memory usage |
| 6360 | {{< alias "toastinfo" >}} | 1.5 | show details on toasted datums |
| 6370 | {{< alias "explain_ui" "pg_explain_ui" >}} | 0.0.2 | easily jump into a visual plan UI for any SQL query |
| 6380 | {{< alias "pg_relusage" >}} | 0.0.1 | Log all the queries that reference a particular column |
| 6800 | {{< alias "pagevis" >}} | 0.1 | Visualise database pages in ascii code |
| 6810 | {{< alias "powa" >}} | 5.1.1 | PostgreSQL Workload Analyser-core |
| 6880 | {{< alias "pg_overexplain" >}} | 1.0 | Allow EXPLAIN to dump even more details |
| 6890 | {{< alias "pg_logicalinspect" >}} | 1.0 | Logical decoding components inspection |
| 6900 | {{< alias "pageinspect" >}} | 1.12 | inspect the contents of database pages at a low level |
| 6910 | {{< alias "pgrowlocks" >}} | 1.2 | show row-level locking information |
| 6920 | {{< alias "sslinfo" >}} | 1.2 | information about SSL certificates |
| 6930 | {{< alias "pg_buffercache" >}} | 1.5 | examine the shared buffer cache |
| 6940 | {{< alias "pg_walinspect" >}} | 1.1 | functions to inspect contents of PostgreSQL Write-Ahead Log |
| 6950 | {{< alias "pg_freespacemap" >}} | 1.2 | examine the free space map (FSM) |
| 6960 | {{< alias "pg_visibility" >}} | 1.2 | examine the visibility map (VM) and page-level visibility info |
| 6970 | {{< alias "pgstattuple" >}} | 1.5 | show tuple-level statistics |
| 6980 | {{< alias "auto_explain" >}} | - | Provides a means for logging execution plans of slow statements automatically |
| 6990 | {{< alias "pg_stat_statements" >}} | 1.11 | track planning and execution statistics of all SQL statements executed |

## SEC

Auditing Logs, Enforce Passwords, Keep Secrets, TDE, SM Algorithm, Login Hooks, Log Errors, Extension White List, ...

| ID | Extension / Package | Version | Description |
|:---:|:---|:---|:---|
| 7000 | {{< alias "passwordcheck_cracklib" >}} | 3.1.0 | Strengthen PostgreSQL user password checks with cracklib |
| 7010 | {{< alias "supautils" >}} | 3.0.2 | Extension that secures a cluster on a cloud environment |
| 7020 | {{< alias "pgsodium" >}} | 3.1.9 | Postgres extension for libsodium functions |
| 7030 | {{< alias "supabase_vault" "pg_vault" >}} | 0.3.1 | Supabase Vault Extension |
| 7040 | {{< alias "pg_session_jwt" >}} | 0.4.0 | Manage authentication sessions using JWTs |
| 7050 | {{< alias "anon" "pg_anon" >}} | 2.5.1 | PostgreSQL Anonymizer (anon) extension |
| 7060 | {{< alias "pgsmcrypto" >}} | 0.1.1 | PostgreSQL SM Algorithm Extension |
| 7070 | {{< alias "pg_enigma" >}} | 0.5.0 | Encrypted postgres data type |
| 7100 | {{< alias "pgaudit" >}} | 18.0 | provides auditing functionality |
| 7120 | {{< alias "pgauditlogtofile" >}} | 1.7.6 | pgAudit addon to redirect audit log to an independent file |
| 7130 | {{< alias "pg_auditor" >}} | 0.2 | Audit data changes and provide flashback ability |
| 7140 | {{< alias "logerrors" >}} | 2.1.5 | Function for collecting statistics about messages in logfile |
| 7150 | {{< alias "pg_auth_mon" >}} | 3.0 | monitor connection attempts per user |
| 7160 | {{< alias "pg_jobmon" >}} | 1.4.1 | Extension for logging and monitoring functions in PostgreSQL |
| 7310 | {{< alias "credcheck" >}} | 4.4 | credcheck - postgresql plain text credential checker |
| 7320 | {{< alias "pgcryptokey" >}} | 0.85 | cryptographic key management |
| 7360 | {{< alias "login_hook" >}} | 1.7 | login_hook - hook to execute login_hook.login() at login time |
| 7370 | {{< alias "set_user" >}} | 4.2.0 | similar to SET ROLE but with added logging |
| 7380 | {{< alias "pg_snakeoil" >}} | 1.4 | The PostgreSQL Antivirus |
| 7390 | {{< alias "pgextwlist" >}} | 1.19 | PostgreSQL Extension Whitelisting |
| 7410 | {{< alias "sslutils" >}} | 1.4 | A Postgres extension for managing SSL certificates through SQL |
| 7420 | {{< alias "noset" "pg_noset" >}} | 0.3.0 | Module for blocking SET variables for non-super users. |
| 7500 | {{< alias "pg_tde" >}} | 1.0 | Percona pg_tde access method |
| 7960 | {{< alias "sepgsql" >}} | - | label-based mandatory access control (MAC) based on SELinux security policy. |
| 7970 | {{< alias "auth_delay" >}} | - | pause briefly before reporting authentication failure |
| 7980 | {{< alias "pgcrypto" >}} | 1.3 | cryptographic functions |
| 7990 | {{< alias "passwordcheck" >}} | - | checks user passwords and reject weak password |

## FDW

Wrappers & Multicorn for FDW Development, Access other DBMS: MySQL, Mongo, SQLite, MSSQL, Oracle, HDFS, DB2,...

| ID | Extension / Package | Version | Description |
|:---:|:---|:---|:---|
| 8500 | {{< alias "wrappers" >}} | 0.5.7 | Foreign data wrappers developed by Supabase |
| 8510 | {{< alias "multicorn" >}} | 3.2 | Fetch foreign data in Python in your PostgreSQL server. |
| 8520 | {{< alias "odbc_fdw" >}} | 0.5.1 | Foreign data wrapper for accessing remote databases using ODBC |
| 8530 | {{< alias "jdbc_fdw" >}} | 0.4.0 | foreign-data wrapper for remote servers available over JDBC |
| 8540 | {{< alias "pgspider_ext" >}} | 1.3.0 | foreign-data wrapper for remote PGSpider servers |
| 8600 | {{< alias "mysql_fdw" >}} | 2.9.3 | Foreign data wrapper for querying a MySQL server |
| 8610 | {{< alias "oracle_fdw" >}} | 2.8.0 | foreign data wrapper for Oracle access |
| 8620 | {{< alias "tds_fdw" >}} | 2.0.5 | Foreign data wrapper for querying a TDS database (Sybase or Microsoft SQL Server) |
| 8630 | {{< alias "db2_fdw" >}} | 18.0.1 | foreign data wrapper for DB2 access |
| 8640 | {{< alias "sqlite_fdw" >}} | 2.5.0 | SQLite Foreign Data Wrapper |
| 8650 | {{< alias "pgbouncer_fdw" >}} | 1.4.0 | Extension for querying PgBouncer stats from normal SQL views & running pgbouncer commands from normal SQL functions |
| 8660 | {{< alias "etcd_fdw" >}} | 0.0.0 | Foreign data wrapper for etcd |
| 8700 | {{< alias "mongo_fdw" >}} | 5.5.3 | foreign data wrapper for MongoDB access |
| 8710 | {{< alias "redis_fdw" >}} | 1.0 | Foreign data wrapper for querying a Redis server |
| 8720 | {{< alias "redis" "pg_redis_pubsub" >}} | 0.0.1 | Send redis pub/sub messages to Redis from PostgreSQL Directly |
| 8730 | {{< alias "kafka_fdw" >}} | 0.0.3 | kafka Foreign Data Wrapper for CSV formatted messages |
| 8740 | {{< alias "hdfs_fdw" >}} | 2.3.3 | foreign-data wrapper for remote hdfs servers |
| 8750 | {{< alias "firebird_fdw" >}} | 1.4.1 | Foreign data wrapper for Firebird |
| 8800 | {{< alias "aws_s3" >}} | 0.0.1 | aws_s3 postgres extension to import/export data from/to s3 |
| 8810 | {{< alias "log_fdw" >}} | 1.4 | foreign-data wrapper for Postgres log file access |
| 8970 | {{< alias "dblink" >}} | 1.2 | connect to other PostgreSQL databases from within a database |
| 8980 | {{< alias "file_fdw" >}} | 1.0 | foreign-data wrapper for flat file access |
| 8990 | {{< alias "postgres_fdw" >}} | 1.1 | foreign-data wrapper for remote PostgreSQL servers |

## SIM

Protocol Simulation & heterogeneous DBMS Compatibility: Oracle, MSSQL, DB2, MySQL, Memcached, and Babelfish!

| ID | Extension / Package | Version | Description |
|:---:|:---|:---|:---|
| 9000 | {{< alias "documentdb" >}} | 0.109 | API surface for DocumentDB for PostgreSQL |
| 9010 | {{< alias "documentdb_core" "documentdb" >}} | 0.109 | Core API surface for DocumentDB for PostgreSQL |
| 9020 | {{< alias "documentdb_distributed" "documentdb" >}} | 0.109 | Multi-Node API surface for DocumentDB |
| 9030 | {{< alias "documentdb_extended_rum" "documentdb" >}} | 0.109 | DocumentDB Extended RUM index access method |
| 9100 | {{< alias "orafce" >}} | 4.16.3 | Functions and operators that emulate a subset of functions and packages from the Oracle RDBMS |
| 9110 | {{< alias "pgtt" >}} | 4.4 | Extension to add Global Temporary Tables feature to PostgreSQL |
| 9120 | {{< alias "session_variable" >}} | 3.4 | Registration and manipulation of session variables and constants |
| 9130 | {{< alias "pg_statement_rollback" >}} | 1.5 | Server side rollback at statement level for PostgreSQL like Oracle or DB2 |
| 9240 | {{< alias "pg_dbms_metadata" >}} | 1.0.0 | Extension to add Oracle DBMS_METADATA compatibility to PostgreSQL |
| 9250 | {{< alias "pg_dbms_lock" >}} | 1.0 | Extension to add Oracle DBMS_LOCK full compatibility to PostgreSQL |
| 9260 | {{< alias "pg_dbms_job" >}} | 1.5 | Extension to add Oracle DBMS_JOB full compatibility to PostgreSQL |
| 9270 | {{< alias "pg_dbms_errlog" >}} | 2.2 | Emulate DBMS_ERRLOG Oracle module to log DML errors in a dedicated table. |
| 9300 | {{< alias "babelfishpg_common" >}} | 3.3.3 | SQL Server Transact SQL Datatype Support |
| 9310 | {{< alias "babelfishpg_tsql" >}} | 3.3.1 | SQL Server Transact SQL compatibility |
| 9320 | {{< alias "babelfishpg_tds" >}} | 1.0.0 | SQL Server TDS protocol extension |
| 9330 | {{< alias "babelfishpg_money" >}} | 1.1.0 | SQL Server Money Data Type |
| 9400 | {{< alias "spat" >}} | 0.1.0a4 | Redis-like In-Memory DB Embedded in Postgres |
| 9410 | {{< alias "pgmemcache" >}} | 2.3.0 | memcached interface |

## ETL

Logical Replication, Decoding, CDC in protobuf/JSON/Mongo format, Copy & Load & Compare Postgres Databases,...

| ID | Extension / Package | Version | Description |
|:---:|:---|:---|:---|
| 9500 | {{< alias "pglogical" >}} | 2.4.6 | PostgreSQL Logical Replication |
| 9501 | {{< alias "pglogical_origin" "pglogical" >}} | 2.4.6 | Dummy extension for compatibility when upgrading from Postgres 9.4 |
| 9510 | {{< alias "pglogical_ticker" >}} | 1.4.1 | Have an accurate view on pglogical replication delay |
| 9520 | {{< alias "pgl_ddl_deploy" >}} | 2.2.1 | automated ddl deployment using pglogical |
| 9530 | {{< alias "pg_failover_slots" >}} | 1.2.0 | PG Failover Slots extension |
| 9540 | {{< alias "db_migrator" >}} | 1.0.0 | Tools to migrate other databases to PostgreSQL |
| 9550 | {{< alias "pgactive" >}} | 2.1.7 | Active-Active Replication Extension for PostgreSQL |
| 9630 | {{< alias "wal2json" >}} | 2.6 | Changing data capture in JSON format |
| 9640 | {{< alias "wal2mongo" >}} | 1.0.7 | PostgreSQL logical decoding output plugin for MongoDB |
| 9650 | {{< alias "decoderbufs" >}} | 3.4.0 | Logical decoding plugin that delivers WAL stream changes using a Protocol Buffer format |
| 9660 | {{< alias "decoder_raw" >}} | 1.0 | Output plugin for logical replication in Raw SQL format |
| 9700 | {{< alias "mimeo" >}} | 1.5.1 | Extension for specialized, per-table replication between PostgreSQL instances |
| 9710 | {{< alias "repmgr" >}} | 5.5.0 | Replication manager for PostgreSQL |
| 9820 | {{< alias "pg_fact_loader" >}} | 2.0.1 | build fact tables with Postgres |
| 9830 | {{< alias "pg_bulkload" >}} | 3.1.23 | pg_bulkload is a high speed data loading utility for PostgreSQL |
| 9970 | {{< alias "test_decoding" >}} | - | SQL-based test/example module for WAL logical decoding |
| 9980 | {{< alias "pgoutput" >}} | - | Logical Replication output plugin |

