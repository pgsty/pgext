---
title: "Extensions"
breadcrumbs: false
excludeSearch: true
comments: false
weight: 900
---

There are 461 available PostgreSQL extensions:

| Extension | PG Versions | Attribute | Category | Description |
|:----------|:------------|:---------:|:--------:|:--------------|
| {{< ext "timescaledb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "TIME" >}} | Enables scalable inserts and complex queries for time-series data |
| {{< ext "timescaledb_toolkit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "TIME" >}} | Library of analytical hyperfunctions, time-series pipelining, and other SQL utilities |
| {{< ext "timeseries" "pg_timeseries" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "TIME" >}} | Convenience API for time series stack |
| {{< ext "periods" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TIME" >}} | Provide Standard SQL functionality for PERIODs and SYSTEM VERSIONING |
| {{< ext "temporal_tables" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TIME" >}} | temporal tables |
| {{< ext "emaj" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TIME" >}} | Enables fine-grained write logging and time travel on subsets of the database. |
| {{< ext "table_version" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TIME" >}} | PostgreSQL table versioning extension |
| {{< ext "pg_cron" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "TIME" >}} | Job scheduler for PostgreSQL |
| {{< ext "pg_task" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "TIME" >}} | execute any sql command at any specific time at background |
| {{< ext "pg_later" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TIME" >}} | Run queries now and get results later |
| {{< ext "pg_background" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TIME" >}} | Run SQL queries in the background |
| {{< ext "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "GIS" >}} | PostGIS geometry and geography spatial types and functions |
| {{< ext "postgis_topology" "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "GIS" >}} | PostGIS topology spatial types and functions |
| {{< ext "postgis_raster" "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "GIS" >}} | PostGIS raster types and functions |
| {{< ext "postgis_sfcgal" "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | PostGIS SFCGAL functions |
| {{< ext "postgis_tiger_geocoder" "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "GIS" >}} | PostGIS tiger geocoder and reverse geocoder |
| {{< ext "address_standardizer" "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | Used to parse an address into constituent elements. Generally used to support geocoding address norm |
| {{< ext "address_standardizer_data_us" "postgis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | Address Standardizer US dataset example |
| {{< ext "pgrouting" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | pgRouting Extension |
| {{< ext "pointcloud" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "GIS" >}} | data type for lidar point clouds |
| {{< ext "pointcloud_postgis" "pointcloud" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "GIS" >}} | integration for pointcloud LIDAR data and PostGIS geometry data |
| {{< ext "h3" "pg_h3" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | H3 bindings for PostgreSQL |
| {{< ext "h3_postgis" "pg_h3" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | H3 PostGIS integration |
| {{< ext "q3c" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | q3c sky indexing plugin |
| {{< ext "ogr_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | foreign-data wrapper for GIS data access |
| {{< ext "geoip" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "GIS" >}} | IP-based geolocation query |
| {{< ext "pg_polyline" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "GIS" >}} | Fast Google Encoded Polyline encoding & decoding for postgres |
| {{< ext "pg_geohash" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | Handle geohash based functionality for spatial coordinates |
| {{< ext "mobilitydb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "GIS" >}} | MobilityDB geospatial trajectory data management & analysis platform |
| {{< ext "mobilitydb_datagen" "mobilitydb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "GIS" >}} | MobilityDB random data generator functions |
| {{< ext "tzf" "pg_tzf" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "GIS" >}} | Fast lookup timezone name by GPS coordinates |
| {{< ext "earthdistance" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "GIS" >}} | calculate great-circle distances on the surface of the Earth |
| {{< ext "vector" "pgvector" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "RAG" >}} | vector data type and ivfflat and hnsw access methods |
| {{< ext "vchord" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "RAG" >}} | Vector database plugin for Postgres, written in Rust |
| {{< ext "vectorscale" "pgvectorscale" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "RAG" >}} | Advanced indexing for vector data with DiskANN |
| {{< ext "vectorize" "pg_vectorize" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "RAG" >}} | The simplest way to do vector search on Postgres |
| {{< ext "pg_similarity" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "RAG" >}} | support similarity queries |
| {{< ext "smlar" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "RAG" >}} | Effective similarity search |
| {{< ext "pg_summarize" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "RAG" >}} | Text Summarization using LLMs. Built using pgrx |
| {{< ext "pg_tiktoken" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "RAG" >}} | tiktoken tokenizer for use with OpenAI models in postgres |
| {{< ext "pg4ml" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | {{< category "RAG" >}} | Machine learning framework for PostgreSQL |
| {{< ext "pgml" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "RAG" >}} | Run AL/ML workloads with SQL interface |
| {{< ext "pg_search" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "FTS" >}} | Full text search for PostgreSQL using BM25 |
| {{< ext "pgroonga" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | {{< category "FTS" >}} | Use Groonga as index, fast full text search platform for all languages! |
| {{< ext "pgroonga_database" "pgroonga" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | {{< category "FTS" >}} | PGroonga database management module |
| {{< ext "pg_bigm" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FTS" >}} | create 2-gram (bigram) index for faster full text search. |
| {{< ext "zhparser" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FTS" >}} | a parser for full-text search of Chinese |
| {{< ext "pg_bestmatch" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "FTS" >}} | Generate BM25 sparse vector inside PostgreSQL |
| {{< ext "vchord_bm25" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "FTS" >}} | A postgresql extension for bm25 ranking algorithm |
| {{< ext "pg_tokenizer" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "FTS" >}} | Tokenizers for full-text search |
| {{< ext "biscuit" "pg_biscuit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,r,r,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FTS" >}} | IAM-LIKE pattern matching with bitmap indexing |
| {{< ext "pg_textsearch" >}} | {{< pgver "18,17,16,15,14,13" "g,g,r,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FTS" >}} | Full-text search with BM25 ranking |
| {{< ext "hunspell_cs_cz" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | Czech Hunspell Dictionary |
| {{< ext "hunspell_de_de" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | German Hunspell Dictionary |
| {{< ext "hunspell_en_us" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | en_US Hunspell Dictionary |
| {{< ext "hunspell_fr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | French Hunspell Dictionary |
| {{< ext "hunspell_ne_np" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | Nepali Hunspell Dictionary |
| {{< ext "hunspell_nl_nl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | Dutch Hunspell Dictionary |
| {{< ext "hunspell_nn_no" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | Norwegian (norsk) Hunspell Dictionary |
| {{< ext "hunspell_pt_pt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | Portuguese Hunspell Dictionary |
| {{< ext "hunspell_ru_ru" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | Russian Hunspell Dictionary |
| {{< ext "hunspell_ru_ru_aot" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FTS" >}} | Russian Hunspell Dictionary (from AOT.ru group) |
| {{< ext "fuzzystrmatch" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "FTS" >}} | determine similarities and distance between strings |
| {{< ext "pg_trgm" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FTS" >}} | text similarity measurement and index searching based on trigrams |
| {{< ext "citus" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,r,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "OLAP" >}} | Distributed PostgreSQL as an extension |
| {{< ext "citus_columnar" "citus" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,r,r,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "OLAP" >}} | Citus columnar storage engine |
| {{< ext "columnar" "hydra" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "OLAP" >}} | Hydra Columnar extension |
| {{< ext "pg_analytics" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "OLAP" >}} | Postgres for analytics, powered by DuckDB |
| {{< ext "pg_duckdb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "OLAP" >}} | DuckDB Embedded in Postgres |
| {{< ext "pg_mooncake" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="---Ld--" color="blue" >}} | {{< category "OLAP" >}} | Columnstore Table in Postgres |
| {{< ext "pg_clickhouse" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "OLAP" >}} | Interfaces to query ClickHouse databases from PostgreSQL |
| {{< ext "duckdb_fdw" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "OLAP" >}} | DuckDB Foreign Data Wrapper |
| {{< ext "pg_parquet" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLdt-" color="blue" >}} | {{< category "OLAP" >}} | copy data between Postgres and Parquet |
| {{< ext "pg_fkpart" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "OLAP" >}} | Table partitioning by foreign key utility |
| {{< ext "pg_partman" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "OLAP" >}} | Extension to manage partitioned tables by time or ID |
| {{< ext "plproxy" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "OLAP" >}} | Database partitioning implemented as procedural language |
| {{< ext "pg_strom" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "OLAP" >}} | PG-Strom - big-data processing acceleration using GPU and NVME |
| {{< ext "tablefunc" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "OLAP" >}} | functions that manipulate whole tables, including crosstab |
| {{< ext "age" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | AGE graph database extension |
| {{< ext "hll" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | type for storing hyperloglog data |
| {{< ext "rum" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FEAT" >}} | RUM index access method |
| {{< ext "pg_ai_query" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | AI-powered SQL query generation for PostgreSQL |
| {{< ext "pg_ttl_index" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "FEAT" >}} | Automatic data expiration with TTL indexes |
| {{< ext "pg_graphql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Add in-database GraphQL support |
| {{< ext "pg_jsonschema" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | PostgreSQL extension providing JSON Schema validation |
| {{< ext "jsquery" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FEAT" >}} | data type for jsonb inspection |
| {{< ext "pg_hint_plan" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Give PostgreSQL ability to manually force some decisions in execution plans. |
| {{< ext "hypopg" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FEAT" >}} | Hypothetical indexes for PostgreSQL |
| {{< ext "index_advisor" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FEAT" >}} | Query index advisor |
| {{< ext "plan_filter" "pg_plan_filter" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "FEAT" >}} | filter statements by their execution plans. |
| {{< ext "imgsmlr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FEAT" >}} | Image similarity with haar |
| {{< ext "pg_ivm" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | incremental view maintenance on PostgreSQL |
| {{< ext "pg_incremental" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,r,r,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Incremental Processing by Crunchy Data |
| {{< ext "pgmb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | A simple PostgreSQL Message Broker system |
| {{< ext "pgmq" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "FEAT" >}} | A lightweight message queue. Like AWS SQS and RSMQ but on Postgres. |
| {{< ext "pgq" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Generic queue for PostgreSQL |
| {{< ext "orioledb" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "FEAT" >}} | OrioleDB, the next generation transactional engine |
| {{< ext "pg_cardano" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | A suite of Cardano-related tools |
| {{< ext "rdkit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FEAT" >}} | Cheminformatics functionality for PostgreSQL. |
| {{< ext "omni" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "FEAT" >}} | Advanced adapter for Postgres extensions |
| {{< ext "omni_auth" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Basic session management |
| {{< ext "omni_aws" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | {{< category "FEAT" >}} | Amazon Web Services APIs (S3) |
| {{< ext "omni_cloudevents" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | {{< category "FEAT" >}} | CloudEvents support |
| {{< ext "omni_containers" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Docker container management |
| {{< ext "omni_credentials" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Application credential management |
| {{< ext "omni_csv" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | CSV toolkit |
| {{< ext "omni_datasets" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Dataset provisioning |
| {{< ext "omni_email" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | E-mail framework |
| {{< ext "omni_http" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Basic HTTP types |
| {{< ext "omni_httpc" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | HTTP client |
| {{< ext "omni_httpd" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | HTTP server |
| {{< ext "omni_id" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FEAT" >}} | Identity types |
| {{< ext "omni_json" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | {{< category "FEAT" >}} | JSON toolkit |
| {{< ext "omni_kube" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Kubernetes (k8s) integration |
| {{< ext "omni_ledger" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Financial ledger |
| {{< ext "omni_manifest" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Package installation manifests |
| {{< ext "omni_mimetypes" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | MIME types |
| {{< ext "omni_os" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Operating system integration |
| {{< ext "omni_polyfill" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Postgres API polyfills |
| {{< ext "omni_python" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | First-class Python support |
| {{< ext "omni_regex" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FEAT" >}} | PCRE-compatible regular expressions |
| {{< ext "omni_rest" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | REST API toolkit (with PostgREST support) |
| {{< ext "omni_schema" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Advanced schema management tooling |
| {{< ext "omni_seq" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Distributed integer sequences |
| {{< ext "omni_service" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Service management |
| {{< ext "omni_session" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Session management |
| {{< ext "omni_shmem" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Shared Memory Management |
| {{< ext "omni_sql" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Programmatic SQL manipulation |
| {{< ext "omni_sqlite" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Embedded SQLite |
| {{< ext "omni_test" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Testing framework |
| {{< ext "omni_txn" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Transaction management |
| {{< ext "omni_types" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Advanced types |
| {{< ext "omni_var" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Scoped variables |
| {{< ext "omni_vfs" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Virtual File System |
| {{< ext "omni_vfs_types_v1" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "FEAT" >}} | Virtual File System types (v1) |
| {{< ext "omni_web" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Common web stack primitives |
| {{< ext "omni_worker" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | Generalized worker pool |
| {{< ext "omni_xml" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | XML toolkit |
| {{< ext "omni_yaml" "omnigres" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FEAT" >}} | YAML toolkit |
| {{< ext "bloom" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FEAT" >}} | bloom access method - signature file based index |
| {{< ext "pg_tle" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "LANG" >}} | Trusted Language Extensions for PostgreSQL |
| {{< ext "plv8" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | PL/JavaScript (v8) trusted procedural language |
| {{< ext "pljs" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | PL/JS trusted procedural language |
| {{< ext "pllua" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | Lua as a procedural language |
| {{< ext "hstore_pllua" "pllua" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "LANG" >}} | Hstore transform for Lua |
| {{< ext "plluau" "pllua" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | Lua as an untrusted procedural language |
| {{< ext "hstore_plluau" "pllua" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "LANG" >}} | Hstore transform for untrusted Lua |
| {{< ext "plprql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | Use PRQL in PostgreSQL - Pipelined Relational Query Language |
| {{< ext "pldbgapi" "pldebugger" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "LANG" >}} | server-side support for debugging PL/pgSQL functions |
| {{< ext "plpgsql_check" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "LANG" >}} | extended check for plpgsql functions |
| {{< ext "plprofiler" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "LANG" >}} | server-side support for profiling PL/pgSQL functions |
| {{< ext "plsh" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "LANG" >}} | PL/sh procedural language |
| {{< ext "pljava" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | PL/Java procedural language |
| {{< ext "plr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | load R interpreter and execute R script from within a database |
| {{< ext "plxslt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "LANG" >}} | XSLT procedural language for PostgreSQL |
| {{< ext "pgtap" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | Unit testing for PostgreSQL |
| {{< ext "faker" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | Wrapper for the Faker Python library |
| {{< ext "dbt2" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "LANG" >}} | OSDL-DBT-2 test kit |
| {{< ext "pltcl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "LANG" >}} | PL/Tcl procedural language |
| {{< ext "pltclu" "pltcl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | {{< category "LANG" >}} | PL/TclU untrusted procedural language |
| {{< ext "plperl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "LANG" >}} | PL/Perl procedural language |
| {{< ext "bool_plperl" "plperl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "LANG" >}} | transform between bool and plperl |
| {{< ext "hstore_plperl" "plperl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "LANG" >}} | transform between hstore and plperl |
| {{< ext "jsonb_plperl" "plperl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | {{< category "LANG" >}} | transform between jsonb and plperl |
| {{< ext "plperlu" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "LANG" >}} | PL/PerlU untrusted procedural language |
| {{< ext "bool_plperlu" "plperlu" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | {{< category "LANG" >}} | transform between bool and plperlu |
| {{< ext "jsonb_plperlu" "plperlu" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | {{< category "LANG" >}} | transform between jsonb and plperlu |
| {{< ext "hstore_plperlu" "plperlu" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | {{< category "LANG" >}} | transform between hstore and plperlu |
| {{< ext "plpgsql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "LANG" >}} | PL/pgSQL procedural language |
| {{< ext "plpython3u" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "LANG" >}} | PL/Python3U untrusted procedural language |
| {{< ext "jsonb_plpython3u" "plpython3u" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d-r" color="blue" >}} | {{< category "LANG" >}} | transform between jsonb and plpython3u |
| {{< ext "ltree_plpython3u" "plpython3u" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d-r" color="blue" >}} | {{< category "LANG" >}} | transform between ltree and plpython3u |
| {{< ext "hstore_plpython3u" "plpython3u" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d-r" color="blue" >}} | {{< category "LANG" >}} | transform between hstore and plpython3u |
| {{< ext "prefix" "pg_prefix" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | Prefix Range module for PostgreSQL |
| {{< ext "semver" "pg_semver" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | Semantic version data type |
| {{< ext "unit" "pgunit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TYPE" >}} | SI units extension |
| {{< ext "pgpdf" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLdtr" color="blue" >}} | {{< category "TYPE" >}} | PDF type with meta admin & Full-Text Search |
| {{< ext "pglite_fusion" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "TYPE" >}} | Embed an SQLite database in your PostgreSQL table |
| {{< ext "md5hash" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | type for storing 128-bit binary data inline |
| {{< ext "asn1oid" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | asn1oid extension |
| {{< ext "roaringbitmap" "pg_roaringbitmap" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | support for Roaring Bitmaps |
| {{< ext "pgfaceting" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | {{< category "TYPE" >}} | fast faceting queries using an inverted index |
| {{< ext "pg_sphere" "pgsphere" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | spherical objects with useful functions, operators and index support |
| {{< ext "country" "pg_country" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | Country data type, ISO 3166-1 |
| {{< ext "pg_xenophile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "TYPE" >}} | More than the bare necessities for PostgreSQL i18n and l10n. |
| {{< ext "l10n_table_dependent_extension" "pg_xenophile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | {{< category "TYPE" >}} | PostgreSQL l10n toolbox |
| {{< ext "currency" "pg_currency" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | Custom PostgreSQL currency type in 1Byte |
| {{< ext "collection" "pgcollection" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | Memory optimized data type to be used inside of plpglsql func |
| {{< ext "pgmp" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | Multiple Precision Arithmetic extension |
| {{< ext "numeral" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | numeral datatypes extension |
| {{< ext "pg_rational" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TYPE" >}} | bigint fractions |
| {{< ext "uint" "pguint" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | unsigned integer types |
| {{< ext "uint128" "pg_uint128" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | Native uint128 type |
| {{< ext "hashtypes" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TYPE" >}} | sha1, md5 and other data types for PostgreSQL |
| {{< ext "ip4r" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | IPv4/v6 and IPv4/v6 range index type for PostgreSQL |
| {{< ext "pg_duration" >}} | {{< pgver "18,17,16,15,14,13" "g,g,r,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | data type for representing durations |
| {{< ext "uri" "pg_uri" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | URI Data type for PostgreSQL |
| {{< ext "emailaddr" "pg_emailaddr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TYPE" >}} | Email address type for PostgreSQL |
| {{< ext "acl" "pg_acl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | ACL Data type |
| {{< ext "debversion" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | Debian version number data type |
| {{< ext "pg_rrule" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "TYPE" >}} | RRULE field type for PostgreSQL |
| {{< ext "timestamp9" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TYPE" >}} | timestamp nanosecond resolution |
| {{< ext "chkpass" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "TYPE" >}} | data type for auto-encrypted passwords |
| {{< ext "isn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "TYPE" >}} | data types for international product numbering standards |
| {{< ext "seg" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "TYPE" >}} | data type for representing line segments or floating-point intervals |
| {{< ext "cube" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "TYPE" >}} | data type for multidimensional cubes |
| {{< ext "ltree" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "TYPE" >}} | data type for hierarchical tree-like structures |
| {{< ext "hstore" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "TYPE" >}} | data type for storing sets of (key, value) pairs |
| {{< ext "citext" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "TYPE" >}} | data type for case-insensitive character strings |
| {{< ext "xml2" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | {{< category "TYPE" >}} | XPath querying and XSLT |
| {{< ext "gzip" "pg_gzip" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | gzip and gunzip functions. |
| {{< ext "bzip" "pg_bzip" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "UTIL" >}} | Bzip compression and decompression |
| {{< ext "zstd" "pg_zstd" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | Zstandard compression algorithm implementation in PostgreSQL |
| {{< ext "http" "pg_http" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "UTIL" >}} | HTTP client for PostgreSQL, allows web page retrieval inside the database. |
| {{< ext "pg_net" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "UTIL" >}} | Async HTTP Requests |
| {{< ext "pg_curl" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | Run curl actions for data transfer in URL syntax |
| {{< ext "pg_retry" >}} | {{< pgver "18,17,16,15,14,13" "g,g,r,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | Retry SQL statements on transient errors with exponential backoff |
| {{< ext "pgjq" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dtr" color="blue" >}} | {{< category "UTIL" >}} | Use jq in Postgres |
| {{< ext "pgjwt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "UTIL" >}} | JSON Web Token API for Postgresql |
| {{< ext "pg_smtp_client" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "UTIL" >}} | PostgreSQL extension to send email using SMTP |
| {{< ext "pg_html5_email_address" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "UTIL" >}} | PostgreSQL email validation that is consistent with the HTML5 spec |
| {{< ext "url_encode" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | url_encode, url_decode functions |
| {{< ext "pgsql_tweaks" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | Some functions and views for daily usage |
| {{< ext "pg_extra_time" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | {{< category "UTIL" >}} | Some date time functions and operators that, |
| {{< ext "pgpcre" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | Perl Compatible Regular Expression functions |
| {{< ext "icu_ext" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | Access ICU functions |
| {{< ext "pgqr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | QR Code generator from PostgreSQL |
| {{< ext "pg_protobuf" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | Protobuf support for PostgreSQL |
| {{< ext "envvar" "pg_envvar" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "UTIL" >}} | Fetch the value of an environment variable |
| {{< ext "floatfile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | Simple file storage for arrays of floats |
| {{< ext "pg_render" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "UTIL" >}} | Render HTML in SQL |
| {{< ext "pg_readme" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | {{< category "UTIL" >}} | Generate a README.md document for a database extension or schema |
| {{< ext "pg_readme_test_extension" "pg_readme" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dtr" color="blue" >}} | {{< category "UTIL" >}} | Test generating a README.md document for extension or schema |
| {{< ext "ddl_historization" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "UTIL" >}} | Historize the ddl changes inside PostgreSQL database |
| {{< ext "data_historization" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | {{< category "UTIL" >}} | PLPGSQL Script to historize data in partitionned table |
| {{< ext "schedoc" "pg_schedoc" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "UTIL" >}} | Cross documentation between Django and DBT projects |
| {{< ext "hashlib" "pg_hashlib" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | {{< category "UTIL" >}} | Stable hash functions for Postgres |
| {{< ext "xxhash" "pg_xxhash" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | {{< category "UTIL" >}} | xxhash functions for PostgreSQL |
| {{< ext "shacrypt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | Implements SHA256-CRYPT and SHA512-CRYPT password encryption schemes |
| {{< ext "cryptint" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | Encryption functions for int and bigint values |
| {{< ext "pguecc" "pg_ecdsa" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "UTIL" >}} | uECC bindings for Postgres |
| {{< ext "sparql" "pgsparql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "UTIL" >}} | Query SPARQL datasource with SQL |
| {{< ext "pg_idkit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FUNC" >}} | multi-tool for generating new/niche universally unique identifiers (ex. UUIDv6, ULID, KSUID) |
| {{< ext "pgx_ulid" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "FUNC" >}} | ulid type and methods |
| {{< ext "pg_uuidv7" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | Create UUIDv7 values in postgres |
| {{< ext "permuteseq" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | Pseudo-randomly permute sequences with a format-preserving encryption on elements |
| {{< ext "pg_hashids" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | Short unique id generator for PostgreSQL, using hashids |
| {{< ext "sequential_uuids" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | generator of sequential UUIDs |
| {{< ext "typeid" "pg_typeid" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FUNC" >}} | Allows to use TypeIDs in Postgres natively |
| {{< ext "snowflake" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FUNC" >}} | Snowflake-style 64-bit ID generator and sequence utilities for PostgreSQL |
| {{< ext "topn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FUNC" >}} | type for top-n JSONB |
| {{< ext "quantile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FUNC" >}} | Quantile aggregation function |
| {{< ext "lower_quantile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FUNC" >}} | Lower quantile aggregate function |
| {{< ext "count_distinct" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | An alternative to COUNT(DISTINCT …) aggregate, usable with HashAggregate |
| {{< ext "omnisketch" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | data structure for on-line agg of data into approximate sketch |
| {{< ext "ddsketch" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | Provides ddsketch aggregate function |
| {{< ext "vasco" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | discover hidden correlations in your data with MIC |
| {{< ext "xicor" "pgxicor" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | {{< category "FUNC" >}} | XI Correlation Coefficient in Postgres |
| {{< ext "weighted_statistics" "pg_weighted_statistics" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dtr" color="blue" >}} | {{< category "FUNC" >}} | High-performance weighted statistics functions for sparse data |
| {{< ext "tdigest" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | Provides tdigest aggregate function. |
| {{< ext "first_last_agg" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | first() and last() aggregate functions |
| {{< ext "extra_window_functions" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | Extra Window Functions for PostgreSQL |
| {{< ext "floatvec" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | Math for vectors (arrays) of numbers |
| {{< ext "aggs_for_vecs" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | Aggregate functions for array inputs |
| {{< ext "aggs_for_arrays" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | Various functions for computing statistics on arrays of numbers |
| {{< ext "pg_csv" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dtr" color="blue" >}} | {{< category "FUNC" >}} | Flexible CSV processing for Postgres |
| {{< ext "arraymath" "pg_arraymath" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | Array math and operators that work element by element on the contents of arrays |
| {{< ext "pg_math" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | GSL statistical functions for postgresql |
| {{< ext "random" "pg_random" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | random data generator |
| {{< ext "base36" "pg_base36" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FUNC" >}} | Integer Base36 types |
| {{< ext "base62" "pg_base62" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FUNC" >}} | Base62 extension for PostgreSQL |
| {{< ext "pg_base58" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FUNC" >}} | Base58 Encoder/Decoder Extension for PostgreSQL |
| {{< ext "financial" "pg_financial" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FUNC" >}} | Financial aggregate functions |
| {{< ext "convert" "pg_convert" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "FUNC" >}} | conversion functions for spatial, routing and other specialized uses |
| {{< ext "refint" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FUNC" >}} | functions for implementing referential integrity (obsolete) |
| {{< ext "autoinc" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FUNC" >}} | functions for autoincrementing fields |
| {{< ext "insert_username" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FUNC" >}} | functions for tracking who changed a table |
| {{< ext "moddatetime" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FUNC" >}} | functions for tracking last modification time |
| {{< ext "tsm_system_time" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "FUNC" >}} | TABLESAMPLE method which accepts time in milliseconds as a limit |
| {{< ext "dict_xsyn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FUNC" >}} | text search dictionary template for extended synonym processing |
| {{< ext "tsm_system_rows" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "FUNC" >}} | TABLESAMPLE method which accepts number of rows as a limit |
| {{< ext "tcn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "FUNC" >}} | Triggered change notifications |
| {{< ext "uuid-ossp" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "FUNC" >}} | generate universally unique identifiers (UUIDs) |
| {{< ext "btree_gist" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "FUNC" >}} | support for indexing common datatypes in GiST |
| {{< ext "btree_gin" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "FUNC" >}} | support for indexing common datatypes in GIN |
| {{< ext "intarray" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---dt-" color="blue" >}} | {{< category "FUNC" >}} | functions, operators, and index support for 1-D arrays of integers |
| {{< ext "intagg" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c---d--" color="blue" >}} | {{< category "FUNC" >}} | integer aggregator and enumerator (obsolete) |
| {{< ext "dict_int" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "FUNC" >}} | text search dictionary template for integers |
| {{< ext "unaccent" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "FUNC" >}} | text search dictionary that removes accents |
| {{< ext "pg_repack" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="-bs-d--" color="blue" >}} | {{< category "ADMIN" >}} | Reorganize tables in PostgreSQL databases with minimal locks |
| {{< ext "pg_rewrite" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "ADMIN" >}} | Tool allows read write to the table during the rewriting |
| {{< ext "pg_squeeze" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "ADMIN" >}} | A tool to remove unused space from a relation. |
| {{< ext "pg_dirtyread" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "ADMIN" >}} | Read dead but unvacuumed rows from table |
| {{< ext "pgfincore" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "ADMIN" >}} | examine and manage the os buffer cache |
| {{< ext "pg_cooldown" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "ADMIN" >}} | remove buffered pages for specific relations |
| {{< ext "ddlx" "pg_ddlx" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ADMIN" >}} | DDL eXtractor functions |
| {{< ext "pglinter" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ADMIN" >}} | PostgreSQL Linting and Analysis Extension |
| {{< ext "prioritize" "pg_prioritize" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "ADMIN" >}} | get and set the priority of PostgreSQL backends |
| {{< ext "pg_checksums" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s---r" color="blue" >}} | {{< category "ADMIN" >}} | Activate/deactivate/verify checksums in offline Postgres clusters |
| {{< ext "pg_readonly" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ADMIN" >}} | cluster database read only |
| {{< ext "pgdd" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "ADMIN" >}} | Introspect pg data dictionary via standard SQL |
| {{< ext "pg_permissions" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "ADMIN" >}} | view object permissions and compare them with the desired state |
| {{< ext "pgautofailover" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "ADMIN" >}} | pg_auto_failover |
| {{< ext "pg_catcheck" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ADMIN" >}} | Diagnosing system catalog corruption |
| {{< ext "pre_prepare" "preprepare" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "ADMIN" >}} | Pre Prepare your Statement server side |
| {{< ext "pg_upless" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | {{< category "ADMIN" >}} | Detect Useless UPDATE |
| {{< ext "pgcozy" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "ADMIN" >}} | Pre-warming shared buffers according to previous pg_buffercache snapshots for PostgreSQL. |
| {{< ext "pg_orphaned" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ADMIN" >}} | Deal with orphaned files |
| {{< ext "pg_crash" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "ADMIN" >}} | Send random signals to random processes |
| {{< ext "pg_cheat_funcs" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "ADMIN" >}} | Provides cheat (but useful) functions |
| {{< ext "fio" "pg_fio" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "ADMIN" >}} | PostgreSQL File I/O Functions |
| {{< ext "pg_savior" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "ADMIN" >}} | Postgres extension to save OOPS mistakes |
| {{< ext "safeupdate" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "ADMIN" >}} | Require criteria for UPDATE and DELETE |
| {{< ext "pg_strict" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "ADMIN" >}} | Prevent dangerous UPDATE and DELETE without WHERE clause |
| {{< ext "pg_drop_events" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "ADMIN" >}} | logs transaction ids of drop table, drop column, drop materialized view statements |
| {{< ext "table_log" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ADMIN" >}} | record table modification logs and PITR for table/row |
| {{< ext "pgagent" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ADMIN" >}} | A PostgreSQL job scheduler |
| {{< ext "pg_prewarm" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sLd--" color="blue" >}} | {{< category "ADMIN" >}} | prewarm relation data |
| {{< ext "pgpool_adm" "pgpool" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "ADMIN" >}} | Administrative functions for pgPool |
| {{< ext "pgpool_recovery" "pgpool" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "ADMIN" >}} | recovery functions for pgpool-II for V4.3 |
| {{< ext "pgpool_regclass" "pgpool" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "ADMIN" >}} | replacement for regclass |
| {{< ext "lo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "ADMIN" >}} | Large Object maintenance |
| {{< ext "basic_archive" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="c-s----" color="blue" >}} | {{< category "ADMIN" >}} | an example of an archive module |
| {{< ext "basebackup_to_shell" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="c-s----" color="blue" >}} | {{< category "ADMIN" >}} | adds a custom basebackup target called shell |
| {{< ext "old_snapshot" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,r" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "ADMIN" >}} | utilities in support of old_snapshot_threshold |
| {{< ext "adminpack" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "ADMIN" >}} | administrative functions for PostgreSQL |
| {{< ext "amcheck" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "ADMIN" >}} | functions for verifying relation integrity |
| {{< ext "pg_surgery" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "ADMIN" >}} | extension to perform surgery on a damaged relation |
| {{< ext "pg_profile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "STAT" >}} | PostgreSQL load profile repository and report builder |
| {{< ext "pg_tracing" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "STAT" >}} | Distributed Tracing for PostgreSQL |
| {{< ext "pg_show_plans" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "STAT" >}} | show query plans of all currently running SQL statements |
| {{< ext "pg_stat_kcache" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "STAT" >}} | Kernel statistics gathering |
| {{< ext "pg_stat_monitor" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "STAT" >}} | The pg_stat_monitor is a PostgreSQL Query Performance Monitoring tool, based on PostgreSQL contrib m |
| {{< ext "pg_qualstats" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "STAT" >}} | An extension collecting statistics about quals |
| {{< ext "pg_store_plans" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "STAT" >}} | track plan statistics of all SQL statements executed |
| {{< ext "pg_track_settings" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "STAT" >}} | Track settings changes |
| {{< ext "pg_track_optimizer" >}} | {{< pgver "18,17,16,15,14,13" "g,g,r,r,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "STAT" >}} | Track planning decisions in comparison with execution reality |
| {{< ext "pg_wait_sampling" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "STAT" >}} | sampling based statistics of wait events |
| {{< ext "pgsentinel" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "STAT" >}} | active session history |
| {{< ext "system_stats" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "STAT" >}} | EnterpriseDB system statistics for PostgreSQL |
| {{< ext "meta" "pg_meta" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "STAT" >}} | Normalized, friendlier system catalog for PostgreSQL |
| {{< ext "pgnodemx" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "STAT" >}} | Capture node OS metrics via SQL queries |
| {{< ext "pg_proctab" "pgnodemx" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "STAT" >}} | PostgreSQL extension to access the OS process table |
| {{< ext "pg_sqlog" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "STAT" >}} | Provide SQL interface to logs |
| {{< ext "bgw_replstatus" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "STAT" >}} | Small PostgreSQL background worker to report whether a node is a replication master or standby |
| {{< ext "pgmeminfo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "STAT" >}} | show memory usage |
| {{< ext "toastinfo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "STAT" >}} | show details on toasted datums |
| {{< ext "explain_ui" "pg_explain_ui" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "STAT" >}} | easily jump into a visual plan UI for any SQL query |
| {{< ext "pg_relusage" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "STAT" >}} | Log all the queries that reference a particular column |
| {{< ext "pagevis" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "STAT" >}} | Visualise database pages in ascii code |
| {{< ext "powa" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "STAT" >}} | PostgreSQL Workload Analyser-core |
| {{< ext "pg_overexplain" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="c-sL---" color="blue" >}} | {{< category "STAT" >}} | Allow EXPLAIN to dump even more details |
| {{< ext "pg_logicalinspect" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "STAT" >}} | Logical decoding components inspection |
| {{< ext "pageinspect" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "STAT" >}} | inspect the contents of database pages at a low level |
| {{< ext "pgrowlocks" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "STAT" >}} | show row-level locking information |
| {{< ext "sslinfo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "STAT" >}} | information about SSL certificates |
| {{< ext "pg_buffercache" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "STAT" >}} | examine the shared buffer cache |
| {{< ext "pg_walinspect" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "STAT" >}} | functions to inspect contents of PostgreSQL Write-Ahead Log |
| {{< ext "pg_freespacemap" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "STAT" >}} | examine the free space map (FSM) |
| {{< ext "pg_visibility" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "STAT" >}} | examine the visibility map (VM) and page-level visibility info |
| {{< ext "pgstattuple" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "STAT" >}} | show tuple-level statistics |
| {{< ext "auto_explain" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | {{< category "STAT" >}} | Provides a means for logging execution plans of slow statements automatically |
| {{< ext "pg_stat_statements" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sLd--" color="blue" >}} | {{< category "STAT" >}} | track planning and execution statistics of all SQL statements executed |
| {{< ext "passwordcheck_cracklib" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "SEC" >}} | Strengthen PostgreSQL user password checks with cracklib |
| {{< ext "supautils" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "SEC" >}} | Extension that secures a cluster on a cloud environment |
| {{< ext "pgsodium" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "SEC" >}} | Postgres extension for libsodium functions |
| {{< ext "supabase_vault" "pg_vault" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SEC" >}} | Supabase Vault Extension |
| {{< ext "pg_session_jwt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "SEC" >}} | Manage authentication sessions using JWTs |
| {{< ext "anon" "pg_anon" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "SEC" >}} | PostgreSQL Anonymizer (anon) extension |
| {{< ext "pgsmcrypto" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SEC" >}} | PostgreSQL SM Algorithm Extension |
| {{< ext "pg_enigma" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SEC" >}} | Encrypted postgres data type |
| {{< ext "pgaudit" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "SEC" >}} | provides auditing functionality |
| {{< ext "pgauditlogtofile" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SEC" >}} | pgAudit addon to redirect audit log to an independent file |
| {{< ext "pg_auditor" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "SEC" >}} | Audit data changes and provide flashback ability |
| {{< ext "logerrors" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SEC" >}} | Function for collecting statistics about messages in logfile |
| {{< ext "pg_auth_mon" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SEC" >}} | monitor connection attempts per user |
| {{< ext "pg_jobmon" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SEC" >}} | Extension for logging and monitoring functions in PostgreSQL |
| {{< ext "credcheck" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "SEC" >}} | credcheck - postgresql plain text credential checker |
| {{< ext "pgcryptokey" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SEC" >}} | cryptographic key management |
| {{< ext "pg_pwhash" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SEC" >}} | Advanced password hashing methods for PostgreSQL |
| {{< ext "login_hook" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SEC" >}} | login_hook - hook to execute login_hook.login() at login time |
| {{< ext "set_user" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SEC" >}} | similar to SET ROLE but with added logging |
| {{< ext "pg_snakeoil" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL--r" color="blue" >}} | {{< category "SEC" >}} | The PostgreSQL Antivirus |
| {{< ext "pgextwlist" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "SEC" >}} | PostgreSQL Extension Whitelisting |
| {{< ext "sslutils" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SEC" >}} | A Postgres extension for managing SSL certificates through SQL |
| {{< ext "noset" "pg_noset" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "SEC" >}} | Module for blocking SET variables for non-super users. |
| {{< ext "pg_tde" >}} | {{< pgver "18,17,16,15,14,13" "g,g,r,r,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "SEC" >}} | Percona pg_tde access method |
| {{< ext "sepgsql" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | {{< category "SEC" >}} | label-based mandatory access control (MAC) based on SELinux security policy. |
| {{< ext "auth_delay" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | {{< category "SEC" >}} | pause briefly before reporting authentication failure |
| {{< ext "pgcrypto" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-dt-" color="blue" >}} | {{< category "SEC" >}} | cryptographic functions |
| {{< ext "passwordcheck" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-sL---" color="blue" >}} | {{< category "SEC" >}} | checks user passwords and reject weak password |
| {{< ext "wrappers" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "FDW" >}} | Foreign data wrappers developed by Supabase |
| {{< ext "multicorn" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FDW" >}} | Fetch foreign data in Python in your PostgreSQL server. |
| {{< ext "odbc_fdw" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | Foreign data wrapper for accessing remote databases using ODBC |
| {{< ext "jdbc_fdw" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | foreign-data wrapper for remote servers available over JDBC |
| {{< ext "pgspider_ext" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | foreign-data wrapper for remote PGSpider servers |
| {{< ext "mysql_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | Foreign data wrapper for querying a MySQL server |
| {{< ext "oracle_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FDW" >}} | foreign data wrapper for Oracle access |
| {{< ext "tds_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | Foreign data wrapper for querying a TDS database (Sybase or Microsoft SQL Server) |
| {{< ext "db2_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FDW" >}} | foreign data wrapper for DB2 access |
| {{< ext "sqlite_fdw" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | SQLite Foreign Data Wrapper |
| {{< ext "pgbouncer_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FDW" >}} | Extension for querying PgBouncer stats from normal SQL views & running pgbouncer commands from norma |
| {{< ext "etcd_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FDW" >}} | Foreign data wrapper for etcd |
| {{< ext "informix_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | Foreign data wrapper for Informix access |
| {{< ext "nominatim_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | Nominatim Foreign Data Wrapper for PostgreSQL |
| {{< ext "mongo_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FDW" >}} | foreign data wrapper for MongoDB access |
| {{< ext "redis_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | Foreign data wrapper for querying a Redis server |
| {{< ext "redis" "pg_redis_pubsub" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | Send redis pub/sub messages to Redis from PostgreSQL Directly |
| {{< ext "kafka_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | kafka Foreign Data Wrapper for CSV formatted messages |
| {{< ext "hdfs_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FDW" >}} | foreign-data wrapper for remote hdfs servers |
| {{< ext "firebird_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "FDW" >}} | Foreign data wrapper for Firebird |
| {{< ext "aws_s3" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d-r" color="blue" >}} | {{< category "FDW" >}} | aws_s3 postgres extension to import/export data from/to s3 |
| {{< ext "log_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "FDW" >}} | foreign-data wrapper for Postgres log file access |
| {{< ext "dblink" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FDW" >}} | connect to other PostgreSQL databases from within a database |
| {{< ext "file_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FDW" >}} | foreign-data wrapper for flat file access |
| {{< ext "postgres_fdw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s-d--" color="blue" >}} | {{< category "FDW" >}} | foreign-data wrapper for remote PostgreSQL servers |
| {{< ext "documentdb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "SIM" >}} | API surface for DocumentDB for PostgreSQL |
| {{< ext "documentdb_core" "documentdb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "SIM" >}} | Core API surface for DocumentDB for PostgreSQL |
| {{< ext "documentdb_distributed" "documentdb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "SIM" >}} | Multi-Node API surface for DocumentDB |
| {{< ext "documentdb_extended_rum" "documentdb" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "SIM" >}} | DocumentDB Extended RUM index access method |
| {{< ext "orafce" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SIM" >}} | Functions and operators that emulate a subset of functions and packages from the Oracle RDBMS |
| {{< ext "pgtt" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SIM" >}} | Extension to add Global Temporary Tables feature to PostgreSQL |
| {{< ext "session_variable" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SIM" >}} | Registration and manipulation of session variables and constants |
| {{< ext "pg_statement_rollback" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL---" color="blue" >}} | {{< category "SIM" >}} | Server side rollback at statement level for PostgreSQL like Oracle or DB2 |
| {{< ext "ivorysql_ora" "ivorysql" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SIM" >}} | Oracle Compatible extension on Postgres Database |
| {{< ext "ora_btree_gin" "ivorysql" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "SIM" >}} | Support for indexing oracle datatypes in GIN |
| {{< ext "ora_btree_gist" "ivorysql" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "SIM" >}} | Support for oracle indexing common datatypes in GiST |
| {{< ext "pg_get_functiondef" "ivorysql" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "SIM" >}} | Get function's definition |
| {{< ext "plisql" "ivorysql" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "SIM" >}} | PL/iSQL procedural language |
| {{< ext "gb18030_2022" "ivorysql" >}} | {{< pgver "18,17,16,15,14,13" "g,r,r,r,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "SIM" >}} | Support GB18030-2022 and UTF-8 conversion |
| {{< ext "pg_dbms_metadata" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SIM" >}} | Extension to add Oracle DBMS_METADATA compatibility to PostgreSQL |
| {{< ext "pg_dbms_lock" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SIM" >}} | Extension to add Oracle DBMS_LOCK full compatibility to PostgreSQL |
| {{< ext "pg_dbms_job" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SIM" >}} | Extension to add Oracle DBMS_JOB full compatibility to PostgreSQL |
| {{< ext "pg_dbms_errlog" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SIM" >}} | Emulate DBMS_ERRLOG Oracle module to log DML errors in a dedicated table. |
| {{< ext "pg_utl_smtp" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,r" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "SIM" >}} | Oracle UTL_SMTP compatibility extension for PostgreSQL |
| {{< ext "babelfishpg_common" "babelfish" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SIM" >}} | SQL Server Transact SQL Datatype Support |
| {{< ext "babelfishpg_tsql" "babelfish" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SIM" >}} | SQL Server Transact SQL compatibility |
| {{< ext "babelfishpg_tds" "babelfish" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--sLd-r" color="blue" >}} | {{< category "SIM" >}} | SQL Server TDS protocol extension |
| {{< ext "babelfishpg_money" "babelfish" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "SIM" >}} | SQL Server Money Data Type |
| {{< ext "spat" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SIM" >}} | Redis-like In-Memory DB Embedded in Postgres |
| {{< ext "pgmemcache" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d-r" color="blue" >}} | {{< category "SIM" >}} | memcached interface |
| {{< ext "aux_mysql" "openhalo" >}} | {{< pgver "18,17,16,15,14,13" "r,r,r,r,g,r" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "SIM" >}} | MySQL Supplementary Extension |
| {{< ext "pglogical" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "ETL" >}} | PostgreSQL Logical Replication |
| {{< ext "pglogical_origin" "pglogical" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ETL" >}} | Dummy extension for compatibility when upgrading from Postgres 9.4 |
| {{< ext "pglogical_ticker" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sLd--" color="blue" >}} | {{< category "ETL" >}} | Have an accurate view on pglogical replication delay |
| {{< ext "pgl_ddl_deploy" >}} | {{< pgver "18,17,16,15,14,13" "r,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ETL" >}} | automated ddl deployment using pglogical |
| {{< ext "pg_failover_slots" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--sL--r" color="blue" >}} | {{< category "ETL" >}} | PG Failover Slots extension |
| {{< ext "db_migrator" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----dt-" color="blue" >}} | {{< category "ETL" >}} | Tools to migrate other databases to PostgreSQL |
| {{< ext "pgactive" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="-bsLd--" color="blue" >}} | {{< category "ETL" >}} | Active-Active Replication Extension for PostgreSQL |
| {{< ext "spock" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="-bsLd--" color="blue" >}} | {{< category "ETL" >}} | Multi-master logical replication extension for PostgreSQL |
| {{< ext "lolor" >}} | {{< pgver "18,17,16,15,14,13" "r,g,r,r,r,r" >}} | {{< badge content="--s-dt-" color="blue" >}} | {{< category "ETL" >}} | Logical-replication-friendly replacement for PostgreSQL large objects |
| {{< ext "wal2json" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | {{< category "ETL" >}} | Changing data capture in JSON format |
| {{< ext "wal2mongo" >}} | {{< pgver "18,17,16,15,14,13" "r,r,g,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | {{< category "ETL" >}} | PostgreSQL logical decoding output plugin for MongoDB |
| {{< ext "decoderbufs" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | {{< category "ETL" >}} | Logical decoding plugin that delivers WAL stream changes using a Protocol Buffer format |
| {{< ext "decoder_raw" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s----" color="blue" >}} | {{< category "ETL" >}} | Output plugin for logical replication in Raw SQL format |
| {{< ext "mimeo" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="----d--" color="blue" >}} | {{< category "ETL" >}} | Extension for specialized, per-table replication between PostgreSQL instances |
| {{< ext "repmgr" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ETL" >}} | Replication manager for PostgreSQL |
| {{< ext "pg_fact_loader" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="--s-d--" color="blue" >}} | {{< category "ETL" >}} | build fact tables with Postgres |
| {{< ext "pg_bulkload" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="-bs-d--" color="blue" >}} | {{< category "ETL" >}} | pg_bulkload is a high speed data loading utility for PostgreSQL |
| {{< ext "test_decoding" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s----" color="blue" >}} | {{< category "ETL" >}} | SQL-based test/example module for WAL logical decoding |
| {{< ext "pgoutput" >}} | {{< pgver "18,17,16,15,14,13" "g,g,g,g,g,g" >}} | {{< badge content="c-s----" color="blue" >}} | {{< category "ETL" >}} | Logical Replication output plugin |

**Attribute Legend**: `c`:contrib `b`:bin `s`:lib `l`:load `d`:ddl `t`:trusted `r`:relocatable