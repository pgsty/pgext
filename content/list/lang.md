---
title: "By Language"
description: "PostgreSQL extensions organized by implementation language"
---

{{< language "C" >}} {{< language "C++" >}} {{< language "Rust" >}} {{< language "Java" >}} {{< language "Python" >}} {{< language "SQL" >}} {{< language "Data" >}}

## Summary

| Language | Count | Description |
|:-------:|:-----:|:------------|
| {{< language "C" >}} | 335 | The traditional PostgreSQL extension language |
| {{< language "SQL" >}} | 37 | Pure SQL extensions and functions |
| {{< language "Rust" >}} | 33 | Extensions written in Rust with the pgrx framework |
| {{< language "Data" >}} | 10 | Data-only extensions |
| {{< language "C++" >}} | 6 | Extensions leveraging C++ features and libraries |
| {{< language "Python" >}} | 2 | Extensions written in Python |
| {{< language "Java" >}} | 1 | Extensions running on JVM |



## C




{{< language "C" >}}
{{< badge content="335 Extensions" color="gray" icon="truck" >}}

The traditional PostgreSQL extension language

| ID | Extension | Package | Description |
|:---:|:---|:---|:---|
| 1000 | {{< ext "timescaledb" >}} | {{< ext "timescaledb" >}} | Enables scalable inserts and complex queries for time-series data |
| 1030 | {{< ext "periods" >}} | {{< ext "periods" >}} | Provide Standard SQL functionality for PERIODs and SYSTEM VERSIONING |
| 1040 | {{< ext "temporal_tables" >}} | {{< ext "temporal_tables" >}} | temporal tables |
| 1070 | {{< ext "pg_cron" >}} | {{< ext "pg_cron" >}} | Job scheduler for PostgreSQL |
| 1080 | {{< ext "pg_task" >}} | {{< ext "pg_task" >}} | execute any sql command at any specific time at background |
| 1100 | {{< ext "pg_background" >}} | {{< ext "pg_background" >}} | Run SQL queries in the background |
| 1500 | {{< ext "postgis" >}} | {{< ext "postgis" >}} | PostGIS geometry and geography spatial types and functions |
| 1501 | {{< ext "postgis_topology" >}} | {{< ext "postgis" >}} | PostGIS topology spatial types and functions |
| 1502 | {{< ext "postgis_raster" >}} | {{< ext "postgis" >}} | PostGIS raster types and functions |
| 1503 | {{< ext "postgis_sfcgal" >}} | {{< ext "postgis" >}} | PostGIS SFCGAL functions |
| 1504 | {{< ext "postgis_tiger_geocoder" >}} | {{< ext "postgis" >}} | PostGIS tiger geocoder and reverse geocoder |
| 1505 | {{< ext "address_standardizer" >}} | {{< ext "postgis" >}} | Used to parse an address into constituent elements. Generally used to support geocoding address normalization step. |
| 1506 | {{< ext "address_standardizer_data_us" >}} | {{< ext "postgis" >}} | Address Standardizer US dataset example |
| 1520 | {{< ext "pointcloud" >}} | {{< ext "pointcloud" >}} | data type for lidar point clouds |
| 1521 | {{< ext "pointcloud_postgis" >}} | {{< ext "pointcloud" >}} | integration for pointcloud LIDAR data and PostGIS geometry data |
| 1530 | {{< ext "h3" >}} | {{< ext "h3" "pg_h3" >}} | H3 bindings for PostgreSQL |
| 1531 | {{< ext "h3_postgis" >}} | {{< ext "h3" "pg_h3" >}} | H3 PostGIS integration |
| 1540 | {{< ext "q3c" >}} | {{< ext "q3c" >}} | q3c sky indexing plugin |
| 1550 | {{< ext "ogr_fdw" >}} | {{< ext "ogr_fdw" >}} | foreign-data wrapper for GIS data access |
| 1590 | {{< ext "pg_geohash" >}} | {{< ext "pg_geohash" >}} | Handle geohash based functionality for spatial coordinates |
| 1650 | {{< ext "mobilitydb" >}} | {{< ext "mobilitydb" >}} | MobilityDB geospatial trajectory data management & analysis platform |
| 1690 | {{< ext "earthdistance" >}} | {{< ext "earthdistance" >}} | calculate great-circle distances on the surface of the Earth |
| 1800 | {{< ext "vector" >}} | {{< ext "vector" "pgvector" >}} | vector data type and ivfflat and hnsw access methods |
| 1840 | {{< ext "pg_similarity" >}} | {{< ext "pg_similarity" >}} | support similarity queries |
| 1850 | {{< ext "smlar" >}} | {{< ext "smlar" >}} | Effective similarity search |
| 1880 | {{< ext "pg4ml" >}} | {{< ext "pg4ml" >}} | Machine learning framework for PostgreSQL |
| 2110 | {{< ext "pgroonga" >}} | {{< ext "pgroonga" >}} | Use Groonga as index, fast full text search platform for all languages! |
| 2111 | {{< ext "pgroonga_database" >}} | {{< ext "pgroonga" >}} | PGroonga database management module |
| 2120 | {{< ext "pg_bigm" >}} | {{< ext "pg_bigm" >}} | create 2-gram (bigram) index for faster full text search. |
| 2130 | {{< ext "zhparser" >}} | {{< ext "zhparser" >}} | a parser for full-text search of Chinese |
| 2180 | {{< ext "fuzzystrmatch" >}} | {{< ext "fuzzystrmatch" >}} | determine similarities and distance between strings |
| 2190 | {{< ext "pg_trgm" >}} | {{< ext "pg_trgm" >}} | text similarity measurement and index searching based on trigrams |
| 2400 | {{< ext "citus" >}} | {{< ext "citus" >}} | Distributed PostgreSQL as an extension |
| 2401 | {{< ext "citus_columnar" >}} | {{< ext "citus" >}} | Citus columnar storage engine |
| 2410 | {{< ext "columnar" >}} | {{< ext "columnar" "hydra" >}} | Hydra Columnar extension |
| 2450 | {{< ext "duckdb_fdw" >}} | {{< ext "duckdb_fdw" >}} | DuckDB Foreign Data Wrapper |
| 2510 | {{< ext "pg_partman" >}} | {{< ext "pg_partman" >}} | Extension to manage partitioned tables by time or ID |
| 2520 | {{< ext "plproxy" >}} | {{< ext "plproxy" >}} | Database partitioning implemented as procedural language |
| 2530 | {{< ext "pg_strom" >}} | {{< ext "pg_strom" >}} | PG-Strom - big-data processing acceleration using GPU and NVME |
| 2590 | {{< ext "tablefunc" >}} | {{< ext "tablefunc" >}} | functions that manipulate whole tables, including crosstab |
| 2760 | {{< ext "age" >}} | {{< ext "age" >}} | AGE graph database extension |
| 2780 | {{< ext "rum" >}} | {{< ext "rum" >}} | RUM index access method |
| 2810 | {{< ext "jsquery" >}} | {{< ext "jsquery" >}} | data type for jsonb inspection |
| 2820 | {{< ext "pg_hint_plan" >}} | {{< ext "pg_hint_plan" >}} | Give PostgreSQL ability to manually force some decisions in execution plans. |
| 2830 | {{< ext "hypopg" >}} | {{< ext "hypopg" >}} | Hypothetical indexes for PostgreSQL |
| 2850 | {{< ext "plan_filter" >}} | {{< ext "plan_filter" "pg_plan_filter" >}} | filter statements by their execution plans. |
| 2860 | {{< ext "imgsmlr" >}} | {{< ext "imgsmlr" >}} | Image similarity with haar |
| 2870 | {{< ext "pg_ivm" >}} | {{< ext "pg_ivm" >}} | incremental view maintenance on PostgreSQL |
| 2880 | {{< ext "pg_incremental" >}} | {{< ext "pg_incremental" >}} | Incremental Processing by Crunchy Data |
| 2910 | {{< ext "pgq" >}} | {{< ext "pgq" >}} | Generic queue for PostgreSQL |
| 2920 | {{< ext "orioledb" >}} | {{< ext "orioledb" >}} | OrioleDB, the next generation transactional engine |
| 2951 | {{< ext "omni" >}} | {{< ext "omni" "omnigres" >}} | Advanced adapter for Postgres extensions |
| 2952 | {{< ext "omni_auth" >}} | {{< ext "omni" "omnigres" >}} | Basic session management |
| 2953 | {{< ext "omni_aws" >}} | {{< ext "omni" "omnigres" >}} | Amazon Web Services APIs (S3) |
| 2954 | {{< ext "omni_cloudevents" >}} | {{< ext "omni" "omnigres" >}} | CloudEvents support |
| 2955 | {{< ext "omni_containers" >}} | {{< ext "omni" "omnigres" >}} | Docker container management |
| 2956 | {{< ext "omni_credentials" >}} | {{< ext "omni" "omnigres" >}} | Application credential management |
| 2958 | {{< ext "omni_email" >}} | {{< ext "omni" "omnigres" >}} | E-mail framework |
| 2959 | {{< ext "omni_http" >}} | {{< ext "omni" "omnigres" >}} | Basic HTTP types |
| 2960 | {{< ext "omni_httpc" >}} | {{< ext "omni" "omnigres" >}} | HTTP client |
| 2961 | {{< ext "omni_httpd" >}} | {{< ext "omni" "omnigres" >}} | HTTP server |
| 2962 | {{< ext "omni_id" >}} | {{< ext "omni" "omnigres" >}} | Identity types |
| 2963 | {{< ext "omni_json" >}} | {{< ext "omni" "omnigres" >}} | JSON toolkit |
| 2964 | {{< ext "omni_kube" >}} | {{< ext "omni" "omnigres" >}} | Kubernetes (k8s) integration |
| 2965 | {{< ext "omni_ledger" >}} | {{< ext "omni" "omnigres" >}} | Financial ledger |
| 2966 | {{< ext "omni_manifest" >}} | {{< ext "omni" "omnigres" >}} | Package installation manifests |
| 2967 | {{< ext "omni_mimetypes" >}} | {{< ext "omni" "omnigres" >}} | MIME types |
| 2968 | {{< ext "omni_os" >}} | {{< ext "omni" "omnigres" >}} | Operating system integration |
| 2969 | {{< ext "omni_polyfill" >}} | {{< ext "omni" "omnigres" >}} | Postgres API polyfills |
| 2970 | {{< ext "omni_python" >}} | {{< ext "omni" "omnigres" >}} | First-class Python support |
| 2971 | {{< ext "omni_regex" >}} | {{< ext "omni" "omnigres" >}} | PCRE-compatible regular expressions |
| 2972 | {{< ext "omni_rest" >}} | {{< ext "omni" "omnigres" >}} | REST API toolkit (with PostgREST support) |
| 2973 | {{< ext "omni_schema" >}} | {{< ext "omni" "omnigres" >}} | Advanced schema management tooling |
| 2974 | {{< ext "omni_seq" >}} | {{< ext "omni" "omnigres" >}} | Distributed integer sequences |
| 2975 | {{< ext "omni_service" >}} | {{< ext "omni" "omnigres" >}} | Service management |
| 2976 | {{< ext "omni_session" >}} | {{< ext "omni" "omnigres" >}} | Session management |
| 2977 | {{< ext "omni_sql" >}} | {{< ext "omni" "omnigres" >}} | Programmatic SQL manipulation |
| 2979 | {{< ext "omni_sqlite" >}} | {{< ext "omni" "omnigres" >}} | Embedded SQLite |
| 2980 | {{< ext "omni_test" >}} | {{< ext "omni" "omnigres" >}} | Testing framework |
| 2981 | {{< ext "omni_txn" >}} | {{< ext "omni" "omnigres" >}} | Transaction management |
| 2982 | {{< ext "omni_types" >}} | {{< ext "omni" "omnigres" >}} | Advanced types |
| 2983 | {{< ext "omni_var" >}} | {{< ext "omni" "omnigres" >}} | Scoped variables |
| 2984 | {{< ext "omni_vfs" >}} | {{< ext "omni" "omnigres" >}} | Virtual File System |
| 2985 | {{< ext "omni_vfs_types_v1" >}} | {{< ext "omni" "omnigres" >}} | Virtual File System types (v1) |
| 2986 | {{< ext "omni_web" >}} | {{< ext "omni" "omnigres" >}} | Common web stack primitives |
| 2987 | {{< ext "omni_worker" >}} | {{< ext "omni" "omnigres" >}} | Generalized worker pool |
| 2988 | {{< ext "omni_xml" >}} | {{< ext "omni" "omnigres" >}} | XML toolkit |
| 2989 | {{< ext "omni_yaml" >}} | {{< ext "omni" "omnigres" >}} | YAML toolkit |
| 2990 | {{< ext "bloom" >}} | {{< ext "bloom" >}} | bloom access method - signature file based index |
| 3000 | {{< ext "pg_tle" >}} | {{< ext "pg_tle" >}} | Trusted Language Extensions for PostgreSQL |
| 3020 | {{< ext "pllua" >}} | {{< ext "pllua" >}} | Lua as a procedural language |
| 3021 | {{< ext "hstore_pllua" >}} | {{< ext "pllua" >}} | Hstore transform for Lua |
| 3030 | {{< ext "plluau" >}} | {{< ext "pllua" >}} | Lua as an untrusted procedural language |
| 3031 | {{< ext "hstore_plluau" >}} | {{< ext "pllua" >}} | Hstore transform for untrusted Lua |
| 3050 | {{< ext "pldbgapi" >}} | {{< ext "pldbgapi" "pldebugger" >}} | server-side support for debugging PL/pgSQL functions |
| 3060 | {{< ext "plpgsql_check" >}} | {{< ext "plpgsql_check" >}} | extended check for plpgsql functions |
| 3070 | {{< ext "plprofiler" >}} | {{< ext "plprofiler" >}} | server-side support for profiling PL/pgSQL functions |
| 3080 | {{< ext "plsh" >}} | {{< ext "plsh" >}} | PL/sh procedural language |
| 3100 | {{< ext "plr" >}} | {{< ext "plr" >}} | load R interpreter and execute R script from within a database |
| 3200 | {{< ext "pgtap" >}} | {{< ext "pgtap" >}} | Unit testing for PostgreSQL |
| 3220 | {{< ext "dbt2" >}} | {{< ext "dbt2" >}} | OSDL-DBT-2 test kit |
| 3240 | {{< ext "pltcl" >}} | {{< ext "pltcl" >}} | PL/Tcl procedural language |
| 3250 | {{< ext "pltclu" >}} | {{< ext "pltcl" >}} | PL/TclU untrusted procedural language |
| 3260 | {{< ext "plperl" >}} | {{< ext "plperl" >}} | PL/Perl procedural language |
| 3261 | {{< ext "bool_plperl" >}} | {{< ext "plperl" >}} | transform between bool and plperl |
| 3262 | {{< ext "hstore_plperl" >}} | {{< ext "plperl" >}} | transform between hstore and plperl |
| 3263 | {{< ext "jsonb_plperl" >}} | {{< ext "plperl" >}} | transform between jsonb and plperl |
| 3270 | {{< ext "plperlu" >}} | {{< ext "plperlu" >}} | PL/PerlU untrusted procedural language |
| 3271 | {{< ext "bool_plperlu" >}} | {{< ext "plperlu" >}} | transform between bool and plperlu |
| 3272 | {{< ext "jsonb_plperlu" >}} | {{< ext "plperlu" >}} | transform between jsonb and plperlu |
| 3273 | {{< ext "hstore_plperlu" >}} | {{< ext "plperlu" >}} | transform between hstore and plperlu |
| 3280 | {{< ext "plpgsql" >}} | {{< ext "plpgsql" >}} | PL/pgSQL procedural language |
| 3290 | {{< ext "plpython3u" >}} | {{< ext "plpython3u" >}} | PL/Python3U untrusted procedural language |
| 3291 | {{< ext "jsonb_plpython3u" >}} | {{< ext "plpython3u" >}} | transform between jsonb and plpython3u |
| 3292 | {{< ext "ltree_plpython3u" >}} | {{< ext "plpython3u" >}} | transform between ltree and plpython3u |
| 3293 | {{< ext "hstore_plpython3u" >}} | {{< ext "plpython3u" >}} | transform between hstore and plpython3u |
| 3500 | {{< ext "prefix" >}} | {{< ext "prefix" "pg_prefix" >}} | Prefix Range module for PostgreSQL |
| 3510 | {{< ext "semver" >}} | {{< ext "semver" "pg_semver" >}} | Semantic version data type |
| 3520 | {{< ext "unit" >}} | {{< ext "unit" "pgunit" >}} | SI units extension |
| 3530 | {{< ext "pgpdf" >}} | {{< ext "pgpdf" >}} | PDF type with meta admin & Full-Text Search |
| 3550 | {{< ext "md5hash" >}} | {{< ext "md5hash" >}} | type for storing 128-bit binary data inline |
| 3560 | {{< ext "asn1oid" >}} | {{< ext "asn1oid" >}} | asn1oid extension |
| 3570 | {{< ext "roaringbitmap" >}} | {{< ext "roaringbitmap" >}} | support for Roaring Bitmaps |
| 3590 | {{< ext "pg_sphere" >}} | {{< ext "pg_sphere" "pgsphere" >}} | spherical objects with useful functions, operators and index support |
| 3600 | {{< ext "country" >}} | {{< ext "country" "pg_country" >}} | Country data type, ISO 3166-1 |
| 3620 | {{< ext "currency" >}} | {{< ext "currency" "pg_currency" >}} | Custom PostgreSQL currency type in 1Byte |
| 3630 | {{< ext "collection" >}} | {{< ext "collection" "pg_collection" >}} | Memory optimized data type to be used inside of plpglsql func |
| 3700 | {{< ext "pgmp" >}} | {{< ext "pgmp" >}} | Multiple Precision Arithmetic extension |
| 3710 | {{< ext "numeral" >}} | {{< ext "numeral" >}} | numeral datatypes extension |
| 3720 | {{< ext "pg_rational" >}} | {{< ext "pg_rational" >}} | bigint fractions |
| 3730 | {{< ext "uint" >}} | {{< ext "uint" "pguint" >}} | unsigned integer types |
| 3740 | {{< ext "uint128" >}} | {{< ext "uint128" "pg_uint128" >}} | Native uint128 type |
| 3750 | {{< ext "hashtypes" >}} | {{< ext "hashtypes" >}} | sha1, md5 and other data types for PostgreSQL |
| 3820 | {{< ext "ip4r" >}} | {{< ext "ip4r" >}} | IPv4/v6 and IPv4/v6 range index type for PostgreSQL |
| 3830 | {{< ext "pg_duration" >}} | {{< ext "pg_duration" >}} | data type for representing durations |
| 3840 | {{< ext "uri" >}} | {{< ext "uri" "pg_uri" >}} | URI Data type for PostgreSQL |
| 3850 | {{< ext "emailaddr" >}} | {{< ext "emailaddr" "pgemailaddr" >}} | Email address type for PostgreSQL |
| 3860 | {{< ext "acl" >}} | {{< ext "acl" "pg_acl" >}} | ACL Data type |
| 3880 | {{< ext "pg_rrule" >}} | {{< ext "pg_rrule" >}} | RRULE field type for PostgreSQL |
| 3890 | {{< ext "timestamp9" >}} | {{< ext "timestamp9" >}} | timestamp nanosecond resolution |
| 3920 | {{< ext "chkpass" >}} | {{< ext "chkpass" >}} | data type for auto-encrypted passwords |
| 3930 | {{< ext "isn" >}} | {{< ext "isn" >}} | data types for international product numbering standards |
| 3940 | {{< ext "seg" >}} | {{< ext "seg" >}} | data type for representing line segments or floating-point intervals |
| 3950 | {{< ext "cube" >}} | {{< ext "cube" >}} | data type for multidimensional cubes |
| 3960 | {{< ext "ltree" >}} | {{< ext "ltree" >}} | data type for hierarchical tree-like structures |
| 3970 | {{< ext "hstore" >}} | {{< ext "hstore" >}} | data type for storing sets of (key, value) pairs |
| 3980 | {{< ext "citext" >}} | {{< ext "citext" >}} | data type for case-insensitive character strings |
| 3990 | {{< ext "xml2" >}} | {{< ext "xml2" >}} | XPath querying and XSLT |
| 4010 | {{< ext "gzip" >}} | {{< ext "gzip" "pg_gzip" >}} | gzip and gunzip functions. |
| 4020 | {{< ext "bzip" >}} | {{< ext "bzip" "pg_bzip" >}} | Bzip compression and decompression |
| 4030 | {{< ext "zstd" >}} | {{< ext "zstd" "pg_zstd" >}} | Zstandard compression algorithm implementation in PostgreSQL |
| 4070 | {{< ext "http" >}} | {{< ext "http" "pg_http" >}} | HTTP client for PostgreSQL, allows web page retrieval inside the database. |
| 4080 | {{< ext "pg_net" >}} | {{< ext "pg_net" >}} | Async HTTP Requests |
| 4090 | {{< ext "pg_curl" >}} | {{< ext "pg_curl" >}} | Run curl actions for data transfer in URL syntax |
| 4150 | {{< ext "pgjq" >}} | {{< ext "pgjq" >}} | Use jq in Postgres |
| 4190 | {{< ext "url_encode" >}} | {{< ext "url_encode" >}} | url_encode, url_decode functions |
| 4230 | {{< ext "pgpcre" >}} | {{< ext "pgpcre" >}} | Perl Compatible Regular Expression functions |
| 4240 | {{< ext "icu_ext" >}} | {{< ext "icu_ext" >}} | Access ICU functions |
| 4250 | {{< ext "pgqr" >}} | {{< ext "pgqr" >}} | QR Code generator from PostgreSQL |
| 4260 | {{< ext "pg_protobuf" >}} | {{< ext "pg_protobuf" >}} | Protobuf support for PostgreSQL |
| 4270 | {{< ext "envvar" >}} | {{< ext "envvar" >}} | Fetch the value of an environment variable |
| 4280 | {{< ext "floatfile" >}} | {{< ext "floatfile" >}} | Simple file storage for arrays of floats |
| 4300 | {{< ext "pg_readme" >}} | {{< ext "pg_readme" >}} | Generate a README.md document for a database extension or schema |
| 4301 | {{< ext "pg_readme_test_extension" >}} | {{< ext "pg_readme" >}} | Test generating a README.md document for extension or schema |
| 4400 | {{< ext "hashlib" >}} | {{< ext "hashlib" "pg_hashlib" >}} | Stable hash functions for Postgres |
| 4430 | {{< ext "xxhash" >}} | {{< ext "xxhash" "pg_xxhash" >}} | xxhash functions for PostgreSQL |
| 4440 | {{< ext "shacrypt" >}} | {{< ext "shacrypt" >}} | Implements SHA256-CRYPT and SHA512-CRYPT password encryption schemes |
| 4450 | {{< ext "cryptint" >}} | {{< ext "cryptint" >}} | Encryption functions for int and bigint values |
| 4460 | {{< ext "pguecc" >}} | {{< ext "pguecc" "pg_ecdsa" >}} | uECC bindings for Postgres |
| 4540 | {{< ext "pg_uuidv7" >}} | {{< ext "pg_uuidv7" >}} | Create UUIDv7 values in postgres |
| 4550 | {{< ext "permuteseq" >}} | {{< ext "permuteseq" >}} | Pseudo-randomly permute sequences with a format-preserving encryption on elements |
| 4560 | {{< ext "pg_hashids" >}} | {{< ext "pg_hashids" >}} | Short unique id generator for PostgreSQL, using hashids |
| 4570 | {{< ext "sequential_uuids" >}} | {{< ext "sequential_uuids" >}} | generator of sequential UUIDs |
| 4600 | {{< ext "topn" >}} | {{< ext "topn" >}} | type for top-n JSONB |
| 4610 | {{< ext "quantile" >}} | {{< ext "quantile" >}} | Quantile aggregation function |
| 4620 | {{< ext "lower_quantile" >}} | {{< ext "lower_quantile" >}} | Lower quantile aggregate function |
| 4630 | {{< ext "count_distinct" >}} | {{< ext "count_distinct" >}} | An alternative to COUNT(DISTINCT …) aggregate, usable with HashAggregate |
| 4640 | {{< ext "omnisketch" >}} | {{< ext "omnisketch" >}} | data structure for on-line agg of data into approximate sketch |
| 4650 | {{< ext "ddsketch" >}} | {{< ext "ddsketch" >}} | Provides ddsketch aggregate function |
| 4660 | {{< ext "vasco" >}} | {{< ext "vasco" >}} | discover hidden correlations in your data with MIC |
| 4670 | {{< ext "xicor" >}} | {{< ext "xicor" "pgxicor" >}} | XI Correlation Coefficient in Postgres |
| 4700 | {{< ext "tdigest" >}} | {{< ext "tdigest" >}} | Provides tdigest aggregate function. |
| 4710 | {{< ext "first_last_agg" >}} | {{< ext "first_last_agg" >}} | first() and last() aggregate functions |
| 4720 | {{< ext "extra_window_functions" >}} | {{< ext "extra_window_functions" >}} | Extra Window Functions for PostgreSQL |
| 4730 | {{< ext "floatvec" >}} | {{< ext "floatvec" >}} | Math for vectors (arrays) of numbers |
| 4740 | {{< ext "aggs_for_vecs" >}} | {{< ext "aggs_for_vecs" >}} | Aggregate functions for array inputs |
| 4750 | {{< ext "aggs_for_arrays" >}} | {{< ext "aggs_for_arrays" >}} | Various functions for computing statistics on arrays of numbers |
| 4760 | {{< ext "arraymath" >}} | {{< ext "arraymath" "pg_arraymath" >}} | Array math and operators that work element by element on the contents of arrays |
| 4770 | {{< ext "pg_math" >}} | {{< ext "pg_math" >}} | GSL statistical functions for postgresql |
| 4780 | {{< ext "random" >}} | {{< ext "random" "pg_random" >}} | random data generator |
| 4800 | {{< ext "base36" >}} | {{< ext "base36" "pg_base36" >}} | Integer Base36 types |
| 4810 | {{< ext "base62" >}} | {{< ext "base62" "pg_base62" >}} | Base62 extension for PostgreSQL |
| 4840 | {{< ext "financial" >}} | {{< ext "financial" "pg_financial" >}} | Financial aggregate functions |
| 4880 | {{< ext "refint" >}} | {{< ext "refint" >}} | functions for implementing referential integrity (obsolete) |
| 4881 | {{< ext "autoinc" >}} | {{< ext "autoinc" >}} | functions for autoincrementing fields |
| 4882 | {{< ext "insert_username" >}} | {{< ext "insert_username" >}} | functions for tracking who changed a table |
| 4883 | {{< ext "moddatetime" >}} | {{< ext "moddatetime" >}} | functions for tracking last modification time |
| 4890 | {{< ext "tsm_system_time" >}} | {{< ext "tsm_system_time" >}} | TABLESAMPLE method which accepts time in milliseconds as a limit |
| 4900 | {{< ext "dict_xsyn" >}} | {{< ext "dict_xsyn" >}} | text search dictionary template for extended synonym processing |
| 4910 | {{< ext "tsm_system_rows" >}} | {{< ext "tsm_system_rows" >}} | TABLESAMPLE method which accepts number of rows as a limit |
| 4920 | {{< ext "tcn" >}} | {{< ext "tcn" >}} | Triggered change notifications |
| 4930 | {{< ext "uuid-ossp" >}} | {{< ext "uuid-ossp" >}} | generate universally unique identifiers (UUIDs) |
| 4940 | {{< ext "btree_gist" >}} | {{< ext "btree_gist" >}} | support for indexing common datatypes in GiST |
| 4950 | {{< ext "btree_gin" >}} | {{< ext "btree_gin" >}} | support for indexing common datatypes in GIN |
| 4960 | {{< ext "intarray" >}} | {{< ext "intarray" >}} | functions, operators, and index support for 1-D arrays of integers |
| 4970 | {{< ext "intagg" >}} | {{< ext "intagg" >}} | integer aggregator and enumerator (obsolete) |
| 4980 | {{< ext "dict_int" >}} | {{< ext "dict_int" >}} | text search dictionary template for integers |
| 4990 | {{< ext "unaccent" >}} | {{< ext "unaccent" >}} | text search dictionary that removes accents |
| 5010 | {{< ext "pg_repack" >}} | {{< ext "pg_repack" >}} | Reorganize tables in PostgreSQL databases with minimal locks |
| 5020 | {{< ext "pg_rewrite" >}} | {{< ext "pg_rewrite" >}} | Tool allows read write to the table during the rewriting |
| 5040 | {{< ext "pg_squeeze" >}} | {{< ext "pg_squeeze" >}} | A tool to remove unused space from a relation. |
| 5050 | {{< ext "pg_dirtyread" >}} | {{< ext "pg_dirtyread" >}} | Read dead but unvacuumed rows from table |
| 5060 | {{< ext "pgfincore" >}} | {{< ext "pgfincore" >}} | examine and manage the os buffer cache |
| 5070 | {{< ext "pg_cooldown" >}} | {{< ext "pg_cooldown" >}} | remove buffered pages for specific relations |
| 5090 | {{< ext "prioritize" >}} | {{< ext "prioritize" "pg_prioritize" >}} | get and set the priority of PostgreSQL backends |
| 5110 | {{< ext "pg_checksums" >}} | {{< ext "pg_checksums" >}} | Activate/deactivate/verify checksums in offline Postgres clusters |
| 5120 | {{< ext "pg_readonly" >}} | {{< ext "pg_readonly" >}} | cluster database read only |
| 5150 | {{< ext "pgautofailover" >}} | {{< ext "pgautofailover" >}} | pg_auto_failover |
| 5160 | {{< ext "pg_catcheck" >}} | {{< ext "pg_catcheck" >}} | Diagnosing system catalog corruption |
| 5170 | {{< ext "pre_prepare" >}} | {{< ext "pre_prepare" "preprepare" >}} | Pre Prepare your Statement server side |
| 5200 | {{< ext "pg_orphaned" >}} | {{< ext "pg_orphaned" >}} | Deal with orphaned files |
| 5210 | {{< ext "pg_crash" >}} | {{< ext "pg_crash" >}} | Send random signals to random processes |
| 5220 | {{< ext "pg_cheat_funcs" >}} | {{< ext "pg_cheat_funcs" >}} | Provides cheat (but useful) functions |
| 5230 | {{< ext "fio" >}} | {{< ext "fio" "pg_fio" >}} | PostgreSQL File I/O Functions |
| 5810 | {{< ext "pg_savior" >}} | {{< ext "pg_savior" >}} | Postgres extension to save OOPS mistakes |
| 5820 | {{< ext "safeupdate" >}} | {{< ext "safeupdate" >}} | Require criteria for UPDATE and DELETE |
| 5840 | {{< ext "table_log" >}} | {{< ext "table_log" >}} | record table modification logs and PITR for table/row |
| 5880 | {{< ext "pgagent" >}} | {{< ext "pgagent" >}} | A PostgreSQL job scheduler |
| 5890 | {{< ext "pg_prewarm" >}} | {{< ext "pg_prewarm" >}} | prewarm relation data |
| 5900 | {{< ext "pgpool_adm" >}} | {{< ext "pgpool_adm" "pgpool" >}} | Administrative functions for pgPool |
| 5910 | {{< ext "pgpool_recovery" >}} | {{< ext "pgpool_adm" "pgpool" >}} | recovery functions for pgpool-II for V4.3 |
| 5920 | {{< ext "pgpool_regclass" >}} | {{< ext "pgpool_adm" "pgpool" >}} | replacement for regclass |
| 5930 | {{< ext "lo" >}} | {{< ext "lo" >}} | Large Object maintenance |
| 5940 | {{< ext "basic_archive" >}} | {{< ext "basic_archive" >}} | an example of an archive module |
| 5950 | {{< ext "basebackup_to_shell" >}} | {{< ext "basebackup_to_shell" >}} | adds a custom basebackup target called shell |
| 5960 | {{< ext "old_snapshot" >}} | {{< ext "old_snapshot" >}} | utilities in support of old_snapshot_threshold |
| 5970 | {{< ext "adminpack" >}} | {{< ext "adminpack" >}} | administrative functions for PostgreSQL |
| 5980 | {{< ext "amcheck" >}} | {{< ext "amcheck" >}} | functions for verifying relation integrity |
| 5990 | {{< ext "pg_surgery" >}} | {{< ext "pg_surgery" >}} | extension to perform surgery on a damaged relation |
| 6000 | {{< ext "pg_profile" >}} | {{< ext "pg_profile" >}} | PostgreSQL load profile repository and report builder |
| 6010 | {{< ext "pg_tracing" >}} | {{< ext "pg_tracing" >}} | Distributed Tracing for PostgreSQL |
| 6210 | {{< ext "pg_show_plans" >}} | {{< ext "pg_show_plans" >}} | show query plans of all currently running SQL statements |
| 6220 | {{< ext "pg_stat_kcache" >}} | {{< ext "pg_stat_kcache" >}} | Kernel statistics gathering |
| 6230 | {{< ext "pg_stat_monitor" >}} | {{< ext "pg_stat_monitor" >}} | The pg_stat_monitor is a PostgreSQL Query Performance Monitoring tool, based on PostgreSQL contrib module pg_stat_statements. pg_stat_monitor provides aggregated statistics, client information, plan details including plan, and histogram information. |
| 6240 | {{< ext "pg_qualstats" >}} | {{< ext "pg_qualstats" >}} | An extension collecting statistics about quals |
| 6250 | {{< ext "pg_store_plans" >}} | {{< ext "pg_store_plans" >}} | track plan statistics of all SQL statements executed |
| 6270 | {{< ext "pg_wait_sampling" >}} | {{< ext "pg_wait_sampling" >}} | sampling based statistics of wait events |
| 6280 | {{< ext "pgsentinel" >}} | {{< ext "pgsentinel" >}} | active session history |
| 6290 | {{< ext "system_stats" >}} | {{< ext "system_stats" >}} | EnterpriseDB system statistics for PostgreSQL |
| 6310 | {{< ext "pgnodemx" >}} | {{< ext "pgnodemx" >}} | Capture node OS metrics via SQL queries |
| 6320 | {{< ext "pg_proctab" >}} | {{< ext "pgnodemx" >}} | PostgreSQL extension to access the OS process table |
| 6340 | {{< ext "bgw_replstatus" >}} | {{< ext "bgw_replstatus" >}} | Small PostgreSQL background worker to report whether a node is a replication master or standby |
| 6350 | {{< ext "pgmeminfo" >}} | {{< ext "pgmeminfo" >}} | show memory usage |
| 6360 | {{< ext "toastinfo" >}} | {{< ext "toastinfo" >}} | show details on toasted datums |
| 6380 | {{< ext "pg_relusage" >}} | {{< ext "pg_relusage" >}} | Log all the queries that reference a particular column |
| 6880 | {{< ext "pg_overexplain" >}} | {{< ext "pg_overexplain" >}} | Allow EXPLAIN to dump even more details |
| 6890 | {{< ext "pg_logicalinspect" >}} | {{< ext "pg_logicalinspect" >}} | Logical decoding components inspection |
| 6900 | {{< ext "pageinspect" >}} | {{< ext "pageinspect" >}} | inspect the contents of database pages at a low level |
| 6910 | {{< ext "pgrowlocks" >}} | {{< ext "pgrowlocks" >}} | show row-level locking information |
| 6920 | {{< ext "sslinfo" >}} | {{< ext "sslinfo" >}} | information about SSL certificates |
| 6930 | {{< ext "pg_buffercache" >}} | {{< ext "pg_buffercache" >}} | examine the shared buffer cache |
| 6940 | {{< ext "pg_walinspect" >}} | {{< ext "pg_walinspect" >}} | functions to inspect contents of PostgreSQL Write-Ahead Log |
| 6950 | {{< ext "pg_freespacemap" >}} | {{< ext "pg_freespacemap" >}} | examine the free space map (FSM) |
| 6960 | {{< ext "pg_visibility" >}} | {{< ext "pg_visibility" >}} | examine the visibility map (VM) and page-level visibility info |
| 6970 | {{< ext "pgstattuple" >}} | {{< ext "pgstattuple" >}} | show tuple-level statistics |
| 6980 | {{< ext "auto_explain" >}} | {{< ext "auto_explain" >}} | Provides a means for logging execution plans of slow statements automatically |
| 6990 | {{< ext "pg_stat_statements" >}} | {{< ext "pg_stat_statements" >}} | track planning and execution statistics of all SQL statements executed |
| 7000 | {{< ext "passwordcheck_cracklib" >}} | {{< ext "passwordcheck_cracklib" "passwordcheck" >}} | Strengthen PostgreSQL user password checks with cracklib |
| 7010 | {{< ext "supautils" >}} | {{< ext "supautils" >}} | Extension that secures a cluster on a cloud environment |
| 7020 | {{< ext "pgsodium" >}} | {{< ext "pgsodium" >}} | Postgres extension for libsodium functions |
| 7030 | {{< ext "supabase_vault" >}} | {{< ext "supabase_vault" "pg_vault" >}} | Supabase Vault Extension |
| 7060 | {{< ext "pg_tde" >}} | {{< ext "pg_tde" >}} | Percona pg_tde access method |
| 7080 | {{< ext "pgaudit" >}} | {{< ext "pgaudit" >}} | provides auditing functionality |
| 7090 | {{< ext "pgauditlogtofile" >}} | {{< ext "pgauditlogtofile" >}} | pgAudit addon to redirect audit log to an independent file |
| 7100 | {{< ext "pg_auth_mon" >}} | {{< ext "pg_auth_mon" >}} | monitor connection attempts per user |
| 7110 | {{< ext "credcheck" >}} | {{< ext "credcheck" >}} | credcheck - postgresql plain text credential checker |
| 7120 | {{< ext "pgcryptokey" >}} | {{< ext "pgcryptokey" >}} | cryptographic key management |
| 7140 | {{< ext "logerrors" >}} | {{< ext "logerrors" >}} | Function for collecting statistics about messages in logfile |
| 7150 | {{< ext "login_hook" >}} | {{< ext "login_hook" >}} | login_hook - hook to execute login_hook.login() at login time |
| 7160 | {{< ext "set_user" >}} | {{< ext "set_user" >}} | similar to SET ROLE but with added logging |
| 7170 | {{< ext "pg_snakeoil" >}} | {{< ext "pg_snakeoil" >}} | The PostgreSQL Antivirus |
| 7180 | {{< ext "pgextwlist" >}} | {{< ext "pgextwlist" >}} | PostgreSQL Extension Whitelisting |
| 7200 | {{< ext "sslutils" >}} | {{< ext "sslutils" >}} | A Postgres extension for managing SSL certificates through SQL |
| 7210 | {{< ext "noset" >}} | {{< ext "noset" "pg_noset" >}} | Module for blocking SET variables for non-super users. |
| 7960 | {{< ext "sepgsql" >}} | {{< ext "sepgsql" >}} | label-based mandatory access control (MAC) based on SELinux security policy. |
| 7970 | {{< ext "auth_delay" >}} | {{< ext "auth_delay" >}} | pause briefly before reporting authentication failure |
| 7980 | {{< ext "pgcrypto" >}} | {{< ext "pgcrypto" >}} | cryptographic functions |
| 7990 | {{< ext "passwordcheck" >}} | {{< ext "passwordcheck_cracklib" "passwordcheck" >}} | checks user passwords and reject weak password |
| 8510 | {{< ext "multicorn" >}} | {{< ext "multicorn" >}} | Fetch foreign data in Python in your PostgreSQL server. |
| 8520 | {{< ext "odbc_fdw" >}} | {{< ext "odbc_fdw" >}} | Foreign data wrapper for accessing remote databases using ODBC |
| 8530 | {{< ext "jdbc_fdw" >}} | {{< ext "jdbc_fdw" >}} | foreign-data wrapper for remote servers available over JDBC |
| 8540 | {{< ext "pgspider_ext" >}} | {{< ext "pgspider_ext" >}} | foreign-data wrapper for remote PGSpider servers |
| 8600 | {{< ext "mysql_fdw" >}} | {{< ext "mysql_fdw" >}} | Foreign data wrapper for querying a MySQL server |
| 8610 | {{< ext "oracle_fdw" >}} | {{< ext "oracle_fdw" >}} | foreign data wrapper for Oracle access |
| 8620 | {{< ext "tds_fdw" >}} | {{< ext "tds_fdw" >}} | Foreign data wrapper for querying a TDS database (Sybase or Microsoft SQL Server) |
| 8630 | {{< ext "db2_fdw" >}} | {{< ext "db2_fdw" >}} | foreign data wrapper for DB2 access |
| 8640 | {{< ext "sqlite_fdw" >}} | {{< ext "sqlite_fdw" >}} | SQLite Foreign Data Wrapper |
| 8700 | {{< ext "mongo_fdw" >}} | {{< ext "mongo_fdw" >}} | foreign data wrapper for MongoDB access |
| 8710 | {{< ext "redis_fdw" >}} | {{< ext "redis_fdw" >}} | Foreign data wrapper for querying a Redis server |
| 8720 | {{< ext "redis" >}} | {{< ext "redis" "pg_redis_pubsub" >}} | Send redis pub/sub messages to Redis from PostgreSQL Directly |
| 8730 | {{< ext "kafka_fdw" >}} | {{< ext "kafka_fdw" >}} | kafka Foreign Data Wrapper for CSV formatted messages |
| 8740 | {{< ext "hdfs_fdw" >}} | {{< ext "hdfs_fdw" >}} | foreign-data wrapper for remote hdfs servers |
| 8750 | {{< ext "firebird_fdw" >}} | {{< ext "firebird_fdw" >}} | Foreign data wrapper for Firebird |
| 8810 | {{< ext "log_fdw" >}} | {{< ext "log_fdw" >}} | foreign-data wrapper for Postgres log file access |
| 8970 | {{< ext "dblink" >}} | {{< ext "dblink" >}} | connect to other PostgreSQL databases from within a database |
| 8980 | {{< ext "file_fdw" >}} | {{< ext "file_fdw" >}} | foreign-data wrapper for flat file access |
| 8990 | {{< ext "postgres_fdw" >}} | {{< ext "postgres_fdw" >}} | foreign-data wrapper for remote PostgreSQL servers |
| 9000 | {{< ext "documentdb" >}} | {{< ext "documentdb" >}} | API surface for DocumentDB for PostgreSQL |
| 9010 | {{< ext "documentdb_core" >}} | {{< ext "documentdb" >}} | Core API surface for DocumentDB for PostgreSQL |
| 9020 | {{< ext "documentdb_distributed" >}} | {{< ext "documentdb" >}} | Multi-Node API surface for DocumentDB |
| 9100 | {{< ext "orafce" >}} | {{< ext "orafce" >}} | Functions and operators that emulate a subset of functions and packages from the Oracle RDBMS |
| 9110 | {{< ext "pgtt" >}} | {{< ext "pgtt" >}} | Extension to add Global Temporary Tables feature to PostgreSQL |
| 9120 | {{< ext "session_variable" >}} | {{< ext "session_variable" >}} | Registration and manipulation of session variables and constants |
| 9130 | {{< ext "pg_statement_rollback" >}} | {{< ext "pg_statement_rollback" >}} | Server side rollback at statement level for PostgreSQL like Oracle or DB2 |
| 9300 | {{< ext "babelfishpg_common" >}} | {{< ext "babelfishpg_common" >}} | SQL Server Transact SQL Datatype Support |
| 9310 | {{< ext "babelfishpg_tsql" >}} | {{< ext "babelfishpg_tsql" >}} | SQL Server Transact SQL compatibility |
| 9320 | {{< ext "babelfishpg_tds" >}} | {{< ext "babelfishpg_tds" >}} | SQL Server TDS protocol extension |
| 9330 | {{< ext "babelfishpg_money" >}} | {{< ext "babelfishpg_money" >}} | SQL Server Money Data Type |
| 9400 | {{< ext "spat" >}} | {{< ext "spat" >}} | Redis-like In-Memory DB Embedded in Postgres |
| 9410 | {{< ext "pgmemcache" >}} | {{< ext "pgmemcache" >}} | memcached interface |
| 9500 | {{< ext "pglogical" >}} | {{< ext "pglogical" >}} | PostgreSQL Logical Replication |
| 9501 | {{< ext "pglogical_origin" >}} | {{< ext "pglogical" >}} | Dummy extension for compatibility when upgrading from Postgres 9.4 |
| 9510 | {{< ext "pglogical_ticker" >}} | {{< ext "pglogical_ticker" >}} | Have an accurate view on pglogical replication delay |
| 9520 | {{< ext "pgl_ddl_deploy" >}} | {{< ext "pgl_ddl_deploy" >}} | automated ddl deployment using pglogical |
| 9530 | {{< ext "pg_failover_slots" >}} | {{< ext "pg_failover_slots" >}} | PG Failover Slots extension |
| 9550 | {{< ext "pgactive" >}} | {{< ext "pgactive" >}} | Active-Active Replication Extension for PostgreSQL |
| 9630 | {{< ext "wal2json" >}} | {{< ext "wal2json" >}} | Changing data capture in JSON format |
| 9640 | {{< ext "wal2mongo" >}} | {{< ext "wal2mongo" >}} | PostgreSQL logical decoding output plugin for MongoDB |
| 9650 | {{< ext "decoderbufs" >}} | {{< ext "decoderbufs" >}} | Logical decoding plugin that delivers WAL stream changes using a Protocol Buffer format |
| 9660 | {{< ext "decoder_raw" >}} | {{< ext "decoder_raw" >}} | Output plugin for logical replication in Raw SQL format |
| 9710 | {{< ext "repmgr" >}} | {{< ext "repmgr" >}} | Replication manager for PostgreSQL |
| 9820 | {{< ext "pg_fact_loader" >}} | {{< ext "pg_fact_loader" >}} | build fact tables with Postgres |
| 9830 | {{< ext "pg_bulkload" >}} | {{< ext "pg_bulkload" >}} | pg_bulkload is a high speed data loading utility for PostgreSQL |
| 9970 | {{< ext "test_decoding" >}} | {{< ext "test_decoding" >}} | SQL-based test/example module for WAL logical decoding |
| 9980 | {{< ext "pgoutput" >}} | {{< ext "pgoutput" >}} | Logical Replication output plugin |


## SQL

{{< language "SQL" >}} {{< badge content="37 Extensions" color="gray" icon="truck" >}}

Pure SQL extensions and functions

| ID | Extension | Package | Description |
|:---:|:---|:---|:---|
| 1020 | {{< ext "timeseries" >}} | {{< ext "timeseries" "pg_timeseries" >}} | Convenience API for time series stack |
| 1050 | {{< ext "emaj" >}} | {{< ext "emaj" >}} | Enables fine-grained write logging and time travel on subsets of the database. |
| 1060 | {{< ext "table_version" >}} | {{< ext "table_version" >}} | PostgreSQL table versioning extension |
| 1560 | {{< ext "geoip" >}} | {{< ext "geoip" >}} | IP-based geolocation query |
| 2500 | {{< ext "pg_fkpart" >}} | {{< ext "pg_fkpart" >}} | Table partitioning by foreign key utility |
| 2840 | {{< ext "index_advisor" >}} | {{< ext "index_advisor" >}} | Query index advisor |
| 2900 | {{< ext "pgmq" >}} | {{< ext "pgmq" >}} | A lightweight message queue. Like AWS SQS and RSMQ but on Postgres. |
| 3580 | {{< ext "pgfaceting" >}} | {{< ext "pgfaceting" >}} | fast faceting queries using an inverted index |
| 3610 | {{< ext "pg_xenophile" >}} | {{< ext "pg_xenophile" >}} | More than the bare necessities for PostgreSQL i18n and l10n. |
| 3611 | {{< ext "l10n_table_dependent_extension" >}} | {{< ext "pg_xenophile" >}} | PostgreSQL l10n toolbox |
| 3870 | {{< ext "debversion" >}} | {{< ext "debversion" >}} | Debian version number data type |
| 4160 | {{< ext "pgjwt" >}} | {{< ext "pgjwt" >}} | JSON Web Token API for Postgresql |
| 4180 | {{< ext "pg_html5_email_address" >}} | {{< ext "pg_html5_email_address" >}} | PostgreSQL email validation that is consistent with the HTML5 spec |
| 4200 | {{< ext "pgsql_tweaks" >}} | {{< ext "pgsql_tweaks" >}} | Some functions and views for daily usage |
| 4220 | {{< ext "pg_extra_time" >}} | {{< ext "pg_extra_time" >}} | Some date time functions and operators that, |
| 4310 | {{< ext "ddl_historization" >}} | {{< ext "ddl_historization" >}} | Historize the ddl changes inside PostgreSQL database |
| 4320 | {{< ext "data_historization" >}} | {{< ext "data_historization" >}} | PLPGSQL Script to historize data in partitionned table |
| 4330 | {{< ext "schedoc" >}} | {{< ext "schedoc" "pg_schedoc" >}} | Cross documentation between Django and DBT projects |
| 4470 | {{< ext "sparql" >}} | {{< ext "sparql" "pgsparql" >}} | Query SPARQL datasource with SQL |
| 5080 | {{< ext "ddlx" >}} | {{< ext "ddlx" "pg_ddlx" >}} | DDL eXtractor functions |
| 5140 | {{< ext "pg_permissions" >}} | {{< ext "pg_permissions" >}} | view object permissions and compare them with the desired state |
| 5180 | {{< ext "pg_upless" >}} | {{< ext "pg_upless" >}} | Detect Useless UPDATE |
| 5190 | {{< ext "pgcozy" >}} | {{< ext "pgcozy" >}} | Pre-warming shared buffers according to previous pg_buffercache snapshots for PostgreSQL. |
| 5830 | {{< ext "pg_drop_events" >}} | {{< ext "pg_drop_events" >}} | logs transaction ids of drop table, drop column, drop materialized view statements |
| 6260 | {{< ext "pg_track_settings" >}} | {{< ext "pg_track_settings" >}} | Track settings changes |
| 6300 | {{< ext "meta" >}} | {{< ext "meta" "pg_meta" >}} | Normalized, friendlier system catalog for PostgreSQL |
| 6330 | {{< ext "pg_sqlog" >}} | {{< ext "pg_sqlog" >}} | Provide SQL interface to logs |
| 6800 | {{< ext "pagevis" >}} | {{< ext "pagevis" >}} | Visualise database pages in ascii code |
| 7130 | {{< ext "pg_jobmon" >}} | {{< ext "pg_jobmon" >}} | Extension for logging and monitoring functions in PostgreSQL |
| 7190 | {{< ext "pg_auditor" >}} | {{< ext "pg_auditor" >}} | Audit data changes and provide flashback ability |
| 8650 | {{< ext "pgbouncer_fdw" >}} | {{< ext "pgbouncer_fdw" >}} | Extension for querying PgBouncer stats from normal SQL views & running pgbouncer commands from normal SQL functions |
| 8800 | {{< ext "aws_s3" >}} | {{< ext "aws_s3" >}} | aws_s3 postgres extension to import/export data from/to s3 |
| 9240 | {{< ext "pg_dbms_metadata" >}} | {{< ext "pg_dbms_metadata" >}} | Extension to add Oracle DBMS_METADATA compatibility to PostgreSQL |
| 9250 | {{< ext "pg_dbms_lock" >}} | {{< ext "pg_dbms_lock" >}} | Extension to add Oracle DBMS_LOCK full compatibility to PostgreSQL |
| 9260 | {{< ext "pg_dbms_job" >}} | {{< ext "pg_dbms_job" >}} | Extension to add Oracle DBMS_JOB full compatibility to PostgreSQL |
| 9540 | {{< ext "db_migrator" >}} | {{< ext "db_migrator" >}} | Tools to migrate other databases to PostgreSQL |
| 9700 | {{< ext "mimeo" >}} | {{< ext "mimeo" >}} | Extension for specialized, per-table replication between PostgreSQL instances |


## Rust

{{< language "Rust" >}} {{< badge content="33 Extensions" color="gray" icon="truck" >}}

Extensions written in Rust with the pgrx framework

| ID | Extension | Package | PGRX Ver | Description |
|:---:|:---|:---|:---|:---|
| 1010 | {{< ext "timescaledb_toolkit" >}} | {{< ext "timescaledb_toolkit" >}} | 0.12.9 | Library of analytical hyperfunctions, time-series pipelining, and other SQL utilities |
| 1090 | {{< ext "pg_later" >}} | {{< ext "pg_later" >}} | 0.12.5 | Run queries now and get results later |
| 1570 | {{< ext "pg_polyline" >}} | {{< ext "pg_polyline" >}} | 0.12.7 | Fast Google Encoded Polyline encoding & decoding for postgres |
| 1680 | {{< ext "tzf" >}} | {{< ext "tzf" "pg_tzf" >}} | 0.14.1 | Fast lookup timezone name by GPS coordinates |
| 1810 | {{< ext "vchord" >}} | {{< ext "vchord" >}} | 0.16.0 | Vector database plugin for Postgres, written in Rust |
| 1820 | {{< ext "vectorscale" >}} | {{< ext "vectorscale" "pgvectorscale" >}} | 0.12.9 | Advanced indexing for vector data with DiskANN |
| 1830 | {{< ext "vectorize" >}} | {{< ext "vectorize" "pg_vectorize" >}} | 0.13.1 | The simplest way to do vector search on Postgres |
| 1860 | {{< ext "pg_summarize" >}} | {{< ext "pg_summarize" >}} | 0.12.4 | Text Summarization using LLMs. Built using pgrx |
| 1870 | {{< ext "pg_tiktoken" >}} | {{< ext "pg_tiktoken" >}} | 0.12.6 | tiktoken tokenizer for use with OpenAI models in postgres |
| 1890 | {{< ext "pgml" >}} | {{< ext "pgml" >}} | 0.12.9 | Run AL/ML workloads with SQL interface |
| 2100 | {{< ext "pg_search" >}} | {{< ext "pg_search" >}} | 0.15.0 | Full text search for PostgreSQL using BM25 |
| 2140 | {{< ext "pg_bestmatch" >}} | {{< ext "pg_bestmatch" >}} | 0.12.7 | Generate BM25 sparse vector inside PostgreSQL |
| 2150 | {{< ext "vchord_bm25" >}} | {{< ext "vchord_bm25" >}} | 0.13.1 | A postgresql extension for bm25 ranking algorithm |
| 2160 | {{< ext "pg_tokenizer" >}} | {{< ext "pg_tokenizer" >}} | 0.13.1 | Tokenizers for full-text search |
| 2420 | {{< ext "pg_analytics" >}} | {{< ext "pg_analytics" >}} | 0.13.0 | Postgres for analytics, powered by DuckDB |
| 2460 | {{< ext "pg_parquet" >}} | {{< ext "pg_parquet" >}} | 0.16.0 | copy data between Postgres and Parquet |
| 2790 | {{< ext "pg_graphql" >}} | {{< ext "pg_graphql" >}} | 0.12.9 | Add in-database GraphQL support |
| 2800 | {{< ext "pg_jsonschema" >}} | {{< ext "pg_jsonschema" >}} | 0.12.9 | PostgreSQL extension providing JSON Schema validation |
| 2930 | {{< ext "pg_cardano" >}} | {{< ext "pg_cardano" >}} | 0.14.1 | A suite of Cardano-related tools |
| 3040 | {{< ext "plprql" >}} | {{< ext "plprql" >}} | 0.11.4 | Use PRQL in PostgreSQL - Pipelined Relational Query Language |
| 3540 | {{< ext "pglite_fusion" >}} | {{< ext "pglite_fusion" >}} | 0.14.1 | Embed an SQLite database in your PostgreSQL table |
| 4170 | {{< ext "pg_smtp_client" >}} | {{< ext "pg_smtp_client" >}} | 0.12.7 | PostgreSQL extension to send email using SMTP |
| 4290 | {{< ext "pg_render" >}} | {{< ext "pg_render" >}} | 0.12.8 | Render HTML in SQL |
| 4500 | {{< ext "pg_idkit" >}} | {{< ext "pg_idkit" >}} | 0.15.0 | multi-tool for generating new/niche universally unique identifiers (ex. UUIDv6, ULID, KSUID) |
| 4510 | {{< ext "pgx_ulid" >}} | {{< ext "pgx_ulid" >}} | 0.12.7 | ulid type and methods |
| 4830 | {{< ext "pg_base58" >}} | {{< ext "pg_base58" >}} | 0.12.1 | Base58 Encoder/Decoder Extension for PostgreSQL |
| 4850 | {{< ext "convert" >}} | {{< ext "convert" "pg_convert" >}} | 0.14.1 | conversion functions for spatial, routing and other specialized uses |
| 5130 | {{< ext "pgdd" >}} | {{< ext "pgdd" >}} | 0.14.1 | Introspect pg data dictionary via standard SQL |
| 6370 | {{< ext "explain_ui" >}} | {{< ext "explain_ui" "pg_explain_ui" >}} | 0.12.4 | easily jump into a visual plan UI for any SQL query |
| 7040 | {{< ext "pg_session_jwt" >}} | {{< ext "pg_session_jwt" >}} | 0.12.9 | Manage authentication sessions using JWTs |
| 7050 | {{< ext "anon" >}} | {{< ext "anon" "pg_anon" >}} | 0.14.3 | PostgreSQL Anonymizer (anon) extension |
| 7070 | {{< ext "pgsmcrypto" >}} | {{< ext "pgsmcrypto" >}} | 0.12.6 | PostgreSQL SM Algorithm Extension |
| 8500 | {{< ext "wrappers" >}} | {{< ext "wrappers" >}} | 0.14.3 | Foreign data wrappers developed by Supabase |


## Data

{{< language "Data" >}} {{< badge content="10 Extensions" color="gray" icon="truck" >}}

Data-only extensions

| ID | Extension | Package | Description |
|:---:|:---|:---|:---|
| 2170 | {{< ext "hunspell_cs_cz" >}} | {{< ext "hunspell_cs_cz" >}} | Czech Hunspell Dictionary |
| 2171 | {{< ext "hunspell_de_de" >}} | {{< ext "hunspell_de_de" >}} | German Hunspell Dictionary |
| 2172 | {{< ext "hunspell_en_us" >}} | {{< ext "hunspell_en_us" >}} | en_US Hunspell Dictionary |
| 2173 | {{< ext "hunspell_fr" >}} | {{< ext "hunspell_fr" >}} | French Hunspell Dictionary |
| 2174 | {{< ext "hunspell_ne_np" >}} | {{< ext "hunspell_ne_np" >}} | Nepali Hunspell Dictionary |
| 2175 | {{< ext "hunspell_nl_nl" >}} | {{< ext "hunspell_nl_nl" >}} | Dutch Hunspell Dictionary |
| 2176 | {{< ext "hunspell_nn_no" >}} | {{< ext "hunspell_nn_no" >}} | Norwegian (norsk) Hunspell Dictionary |
| 2177 | {{< ext "hunspell_pt_pt" >}} | {{< ext "hunspell_pt_pt" >}} | Portuguese Hunspell Dictionary |
| 2178 | {{< ext "hunspell_ru_ru" >}} | {{< ext "hunspell_ru_ru" >}} | Russian Hunspell Dictionary |
| 2179 | {{< ext "hunspell_ru_ru_aot" >}} | {{< ext "hunspell_ru_ru_aot" >}} | Russian Hunspell Dictionary (from AOT.ru group) |


## C++

{{< language "C++" >}} {{< badge content="6 Extensions" color="gray" icon="truck" >}}

Extensions leveraging C++ features and libraries

| ID | Extension | Package | Description |
|:---:|:---|:---|:---|
| 1510 | {{< ext "pgrouting" >}} | {{< ext "pgrouting" >}} | pgRouting Extension |
| 2430 | {{< ext "pg_duckdb" >}} | {{< ext "pg_duckdb" >}} | DuckDB Embedded in Postgres |
| 2440 | {{< ext "pg_mooncake" >}} | {{< ext "pg_mooncake" >}} | Columnstore Table in Postgres |
| 2770 | {{< ext "hll" >}} | {{< ext "hll" >}} | type for storing hyperloglog data |
| 2940 | {{< ext "rdkit" >}} | {{< ext "rdkit" >}} | Cheminformatics functionality for PostgreSQL. |
| 3010 | {{< ext "plv8" >}} | {{< ext "plv8" >}} | PL/JavaScript (v8) trusted procedural language |


## Python

{{< language "Python" >}} {{< badge content="2 Extensions" color="gray" icon="truck" >}}

Extensions written in Python

| ID | Extension | Package | Description |
|:---:|:---|:---|:---|
| 3210 | {{< ext "faker" >}} | {{< ext "faker" >}} | Wrapper for the Faker Python library |
| 6810 | {{< ext "powa" >}} | {{< ext "powa" >}} | PostgreSQL Workload Analyser-core |


## Java

{{< language "Java" >}} {{< badge content="1 Extensions" color="gray" icon="truck" >}}

Extensions running on JVM

|  ID  | Extension            | Package              | Description                 |
|:----:|:---------------------|:---------------------|:----------------------------|
| 3090 | {{< ext "pljava" >}} | {{< ext "pljava" >}} | PL/Java procedural language |

