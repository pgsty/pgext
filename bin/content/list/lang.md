---
title: By Language
description: PostgreSQL extensions organized by implementation language
excludeSearch: true
weight: 1
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

{{< language "C" >}} {{< badge content="335 Extensions" color="gray" icon="cube" >}}


The traditional PostgreSQL extension language

| ID | Extension | Package | Description |
|:---:|:---|:---|:---|
| 3860 | [`acl`](/e/acl) | [`pg_acl`](/e/acl) | ACL Data type |
| 1505 | [`address_standardizer`](/e/address_standardizer) | [`postgis`](/e/postgis) | Used to parse an address into constituent elements. Generally used to support geocoding address normalization step. |
| 1506 | [`address_standardizer_data_us`](/e/address_standardizer_data_us) | [`postgis`](/e/postgis) | Address Standardizer US dataset example |
| 5970 | [`adminpack`](/e/adminpack) | [`adminpack`](/e/adminpack) | administrative functions for PostgreSQL |
| 2760 | [`age`](/e/age) | [`age`](/e/age) | AGE graph database extension |
| 4750 | [`aggs_for_arrays`](/e/aggs_for_arrays) | [`aggs_for_arrays`](/e/aggs_for_arrays) | Various functions for computing statistics on arrays of numbers |
| 4740 | [`aggs_for_vecs`](/e/aggs_for_vecs) | [`aggs_for_vecs`](/e/aggs_for_vecs) | Aggregate functions for array inputs |
| 5980 | [`amcheck`](/e/amcheck) | [`amcheck`](/e/amcheck) | functions for verifying relation integrity |
| 4760 | [`arraymath`](/e/arraymath) | [`pg_arraymath`](/e/arraymath) | Array math and operators that work element by element on the contents of arrays |
| 3560 | [`asn1oid`](/e/asn1oid) | [`asn1oid`](/e/asn1oid) | asn1oid extension |
| 7970 | [`auth_delay`](/e/auth_delay) | [`auth_delay`](/e/auth_delay) | pause briefly before reporting authentication failure |
| 6980 | [`auto_explain`](/e/auto_explain) | [`auto_explain`](/e/auto_explain) | Provides a means for logging execution plans of slow statements automatically |
| 4881 | [`autoinc`](/e/autoinc) | [`autoinc`](/e/autoinc) | functions for autoincrementing fields |
| 9300 | [`babelfishpg_common`](/e/babelfishpg_common) | [`babelfishpg_common`](/e/babelfishpg_common) | SQL Server Transact SQL Datatype Support |
| 9330 | [`babelfishpg_money`](/e/babelfishpg_money) | [`babelfishpg_money`](/e/babelfishpg_money) | SQL Server Money Data Type |
| 9320 | [`babelfishpg_tds`](/e/babelfishpg_tds) | [`babelfishpg_tds`](/e/babelfishpg_tds) | SQL Server TDS protocol extension |
| 9310 | [`babelfishpg_tsql`](/e/babelfishpg_tsql) | [`babelfishpg_tsql`](/e/babelfishpg_tsql) | SQL Server Transact SQL compatibility |
| 4800 | [`base36`](/e/base36) | [`pg_base36`](/e/base36) | Integer Base36 types |
| 4810 | [`base62`](/e/base62) | [`pg_base62`](/e/base62) | Base62 extension for PostgreSQL |
| 5950 | [`basebackup_to_shell`](/e/basebackup_to_shell) | [`basebackup_to_shell`](/e/basebackup_to_shell) | adds a custom basebackup target called shell |
| 5940 | [`basic_archive`](/e/basic_archive) | [`basic_archive`](/e/basic_archive) | an example of an archive module |
| 6340 | [`bgw_replstatus`](/e/bgw_replstatus) | [`bgw_replstatus`](/e/bgw_replstatus) | Small PostgreSQL background worker to report whether a node is a replication master or standby |
| 2990 | [`bloom`](/e/bloom) | [`bloom`](/e/bloom) | bloom access method - signature file based index |
| 3261 | [`bool_plperl`](/e/bool_plperl) | [`plperl`](/e/plperl) | transform between bool and plperl |
| 3271 | [`bool_plperlu`](/e/bool_plperlu) | [`plperlu`](/e/plperlu) | transform between bool and plperlu |
| 4950 | [`btree_gin`](/e/btree_gin) | [`btree_gin`](/e/btree_gin) | support for indexing common datatypes in GIN |
| 4940 | [`btree_gist`](/e/btree_gist) | [`btree_gist`](/e/btree_gist) | support for indexing common datatypes in GiST |
| 4020 | [`bzip`](/e/bzip) | [`pg_bzip`](/e/bzip) | Bzip compression and decompression |
| 3920 | [`chkpass`](/e/chkpass) | [`chkpass`](/e/chkpass) | data type for auto-encrypted passwords |
| 3980 | [`citext`](/e/citext) | [`citext`](/e/citext) | data type for case-insensitive character strings |
| 2400 | [`citus`](/e/citus) | [`citus`](/e/citus) | Distributed PostgreSQL as an extension |
| 2401 | [`citus_columnar`](/e/citus_columnar) | [`citus`](/e/citus) | Citus columnar storage engine |
| 3630 | [`collection`](/e/collection) | [`pg_collection`](/e/collection) | Memory optimized data type to be used inside of plpglsql func |
| 2410 | [`columnar`](/e/columnar) | [`hydra`](/e/columnar) | Hydra Columnar extension |
| 4630 | [`count_distinct`](/e/count_distinct) | [`count_distinct`](/e/count_distinct) | An alternative to COUNT(DISTINCT …) aggregate, usable with HashAggregate |
| 3600 | [`country`](/e/country) | [`pg_country`](/e/country) | Country data type, ISO 3166-1 |
| 7110 | [`credcheck`](/e/credcheck) | [`credcheck`](/e/credcheck) | credcheck - postgresql plain text credential checker |
| 4450 | [`cryptint`](/e/cryptint) | [`cryptint`](/e/cryptint) | Encryption functions for int and bigint values |
| 3950 | [`cube`](/e/cube) | [`cube`](/e/cube) | data type for multidimensional cubes |
| 3620 | [`currency`](/e/currency) | [`pg_currency`](/e/currency) | Custom PostgreSQL currency type in 1Byte |
| 8630 | [`db2_fdw`](/e/db2_fdw) | [`db2_fdw`](/e/db2_fdw) | foreign data wrapper for DB2 access |
| 8970 | [`dblink`](/e/dblink) | [`dblink`](/e/dblink) | connect to other PostgreSQL databases from within a database |
| 3220 | [`dbt2`](/e/dbt2) | [`dbt2`](/e/dbt2) | OSDL-DBT-2 test kit |
| 4650 | [`ddsketch`](/e/ddsketch) | [`ddsketch`](/e/ddsketch) | Provides ddsketch aggregate function |
| 9660 | [`decoder_raw`](/e/decoder_raw) | [`decoder_raw`](/e/decoder_raw) | Output plugin for logical replication in Raw SQL format |
| 9650 | [`decoderbufs`](/e/decoderbufs) | [`decoderbufs`](/e/decoderbufs) | Logical decoding plugin that delivers WAL stream changes using a Protocol Buffer format |
| 4980 | [`dict_int`](/e/dict_int) | [`dict_int`](/e/dict_int) | text search dictionary template for integers |
| 4900 | [`dict_xsyn`](/e/dict_xsyn) | [`dict_xsyn`](/e/dict_xsyn) | text search dictionary template for extended synonym processing |
| 9000 | [`documentdb`](/e/documentdb) | [`documentdb`](/e/documentdb) | API surface for DocumentDB for PostgreSQL |
| 9010 | [`documentdb_core`](/e/documentdb_core) | [`documentdb`](/e/documentdb) | Core API surface for DocumentDB for PostgreSQL |
| 9020 | [`documentdb_distributed`](/e/documentdb_distributed) | [`documentdb`](/e/documentdb) | Multi-Node API surface for DocumentDB |
| 2450 | [`duckdb_fdw`](/e/duckdb_fdw) | [`duckdb_fdw`](/e/duckdb_fdw) | DuckDB Foreign Data Wrapper |
| 1690 | [`earthdistance`](/e/earthdistance) | [`earthdistance`](/e/earthdistance) | calculate great-circle distances on the surface of the Earth |
| 3850 | [`emailaddr`](/e/emailaddr) | [`pgemailaddr`](/e/emailaddr) | Email address type for PostgreSQL |
| 4270 | [`envvar`](/e/envvar) | [`envvar`](/e/envvar) | Fetch the value of an environment variable |
| 4720 | [`extra_window_functions`](/e/extra_window_functions) | [`extra_window_functions`](/e/extra_window_functions) | Extra Window Functions for PostgreSQL |
| 8980 | [`file_fdw`](/e/file_fdw) | [`file_fdw`](/e/file_fdw) | foreign-data wrapper for flat file access |
| 4840 | [`financial`](/e/financial) | [`pg_financial`](/e/financial) | Financial aggregate functions |
| 5230 | [`fio`](/e/fio) | [`pg_fio`](/e/fio) | PostgreSQL File I/O Functions |
| 8750 | [`firebird_fdw`](/e/firebird_fdw) | [`firebird_fdw`](/e/firebird_fdw) | Foreign data wrapper for Firebird |
| 4710 | [`first_last_agg`](/e/first_last_agg) | [`first_last_agg`](/e/first_last_agg) | first() and last() aggregate functions |
| 4280 | [`floatfile`](/e/floatfile) | [`floatfile`](/e/floatfile) | Simple file storage for arrays of floats |
| 4730 | [`floatvec`](/e/floatvec) | [`floatvec`](/e/floatvec) | Math for vectors (arrays) of numbers |
| 2180 | [`fuzzystrmatch`](/e/fuzzystrmatch) | [`fuzzystrmatch`](/e/fuzzystrmatch) | determine similarities and distance between strings |
| 4010 | [`gzip`](/e/gzip) | [`pg_gzip`](/e/gzip) | gzip and gunzip functions. |
| 1530 | [`h3`](/e/h3) | [`pg_h3`](/e/h3) | H3 bindings for PostgreSQL |
| 1531 | [`h3_postgis`](/e/h3_postgis) | [`pg_h3`](/e/h3) | H3 PostGIS integration |
| 4400 | [`hashlib`](/e/hashlib) | [`pg_hashlib`](/e/hashlib) | Stable hash functions for Postgres |
| 3750 | [`hashtypes`](/e/hashtypes) | [`hashtypes`](/e/hashtypes) | sha1, md5 and other data types for PostgreSQL |
| 8740 | [`hdfs_fdw`](/e/hdfs_fdw) | [`hdfs_fdw`](/e/hdfs_fdw) | foreign-data wrapper for remote hdfs servers |
| 3970 | [`hstore`](/e/hstore) | [`hstore`](/e/hstore) | data type for storing sets of (key, value) pairs |
| 3021 | [`hstore_pllua`](/e/hstore_pllua) | [`pllua`](/e/pllua) | Hstore transform for Lua |
| 3031 | [`hstore_plluau`](/e/hstore_plluau) | [`pllua`](/e/pllua) | Hstore transform for untrusted Lua |
| 3262 | [`hstore_plperl`](/e/hstore_plperl) | [`plperl`](/e/plperl) | transform between hstore and plperl |
| 3273 | [`hstore_plperlu`](/e/hstore_plperlu) | [`plperlu`](/e/plperlu) | transform between hstore and plperlu |
| 3293 | [`hstore_plpython3u`](/e/hstore_plpython3u) | [`plpython3u`](/e/plpython3u) | transform between hstore and plpython3u |
| 4070 | [`http`](/e/http) | [`pg_http`](/e/http) | HTTP client for PostgreSQL, allows web page retrieval inside the database. |
| 2830 | [`hypopg`](/e/hypopg) | [`hypopg`](/e/hypopg) | Hypothetical indexes for PostgreSQL |
| 4240 | [`icu_ext`](/e/icu_ext) | [`icu_ext`](/e/icu_ext) | Access ICU functions |
| 2860 | [`imgsmlr`](/e/imgsmlr) | [`imgsmlr`](/e/imgsmlr) | Image similarity with haar |
| 4882 | [`insert_username`](/e/insert_username) | [`insert_username`](/e/insert_username) | functions for tracking who changed a table |
| 4970 | [`intagg`](/e/intagg) | [`intagg`](/e/intagg) | integer aggregator and enumerator (obsolete) |
| 4960 | [`intarray`](/e/intarray) | [`intarray`](/e/intarray) | functions, operators, and index support for 1-D arrays of integers |
| 3820 | [`ip4r`](/e/ip4r) | [`ip4r`](/e/ip4r) | IPv4/v6 and IPv4/v6 range index type for PostgreSQL |
| 3930 | [`isn`](/e/isn) | [`isn`](/e/isn) | data types for international product numbering standards |
| 8530 | [`jdbc_fdw`](/e/jdbc_fdw) | [`jdbc_fdw`](/e/jdbc_fdw) | foreign-data wrapper for remote servers available over JDBC |
| 3263 | [`jsonb_plperl`](/e/jsonb_plperl) | [`plperl`](/e/plperl) | transform between jsonb and plperl |
| 3272 | [`jsonb_plperlu`](/e/jsonb_plperlu) | [`plperlu`](/e/plperlu) | transform between jsonb and plperlu |
| 3291 | [`jsonb_plpython3u`](/e/jsonb_plpython3u) | [`plpython3u`](/e/plpython3u) | transform between jsonb and plpython3u |
| 2810 | [`jsquery`](/e/jsquery) | [`jsquery`](/e/jsquery) | data type for jsonb inspection |
| 8730 | [`kafka_fdw`](/e/kafka_fdw) | [`kafka_fdw`](/e/kafka_fdw) | kafka Foreign Data Wrapper for CSV formatted messages |
| 5930 | [`lo`](/e/lo) | [`lo`](/e/lo) | Large Object maintenance |
| 8810 | [`log_fdw`](/e/log_fdw) | [`log_fdw`](/e/log_fdw) | foreign-data wrapper for Postgres log file access |
| 7140 | [`logerrors`](/e/logerrors) | [`logerrors`](/e/logerrors) | Function for collecting statistics about messages in logfile |
| 7150 | [`login_hook`](/e/login_hook) | [`login_hook`](/e/login_hook) | login_hook - hook to execute login_hook.login() at login time |
| 4620 | [`lower_quantile`](/e/lower_quantile) | [`lower_quantile`](/e/lower_quantile) | Lower quantile aggregate function |
| 3960 | [`ltree`](/e/ltree) | [`ltree`](/e/ltree) | data type for hierarchical tree-like structures |
| 3292 | [`ltree_plpython3u`](/e/ltree_plpython3u) | [`plpython3u`](/e/plpython3u) | transform between ltree and plpython3u |
| 3550 | [`md5hash`](/e/md5hash) | [`md5hash`](/e/md5hash) | type for storing 128-bit binary data inline |
| 1650 | [`mobilitydb`](/e/mobilitydb) | [`mobilitydb`](/e/mobilitydb) | MobilityDB geospatial trajectory data management & analysis platform |
| 4883 | [`moddatetime`](/e/moddatetime) | [`moddatetime`](/e/moddatetime) | functions for tracking last modification time |
| 8700 | [`mongo_fdw`](/e/mongo_fdw) | [`mongo_fdw`](/e/mongo_fdw) | foreign data wrapper for MongoDB access |
| 8510 | [`multicorn`](/e/multicorn) | [`multicorn`](/e/multicorn) | Fetch foreign data in Python in your PostgreSQL server. |
| 8600 | [`mysql_fdw`](/e/mysql_fdw) | [`mysql_fdw`](/e/mysql_fdw) | Foreign data wrapper for querying a MySQL server |
| 7210 | [`noset`](/e/noset) | [`pg_noset`](/e/noset) | Module for blocking SET variables for non-super users. |
| 3710 | [`numeral`](/e/numeral) | [`numeral`](/e/numeral) | numeral datatypes extension |
| 8520 | [`odbc_fdw`](/e/odbc_fdw) | [`odbc_fdw`](/e/odbc_fdw) | Foreign data wrapper for accessing remote databases using ODBC |
| 1550 | [`ogr_fdw`](/e/ogr_fdw) | [`ogr_fdw`](/e/ogr_fdw) | foreign-data wrapper for GIS data access |
| 5960 | [`old_snapshot`](/e/old_snapshot) | [`old_snapshot`](/e/old_snapshot) | utilities in support of old_snapshot_threshold |
| 2951 | [`omni`](/e/omni) | [`omnigres`](/e/omni) | Advanced adapter for Postgres extensions |
| 2952 | [`omni_auth`](/e/omni_auth) | [`omnigres`](/e/omni) | Basic session management |
| 2953 | [`omni_aws`](/e/omni_aws) | [`omnigres`](/e/omni) | Amazon Web Services APIs (S3) |
| 2954 | [`omni_cloudevents`](/e/omni_cloudevents) | [`omnigres`](/e/omni) | CloudEvents support |
| 2955 | [`omni_containers`](/e/omni_containers) | [`omnigres`](/e/omni) | Docker container management |
| 2956 | [`omni_credentials`](/e/omni_credentials) | [`omnigres`](/e/omni) | Application credential management |
| 2958 | [`omni_email`](/e/omni_email) | [`omnigres`](/e/omni) | E-mail framework |
| 2959 | [`omni_http`](/e/omni_http) | [`omnigres`](/e/omni) | Basic HTTP types |
| 2960 | [`omni_httpc`](/e/omni_httpc) | [`omnigres`](/e/omni) | HTTP client |
| 2961 | [`omni_httpd`](/e/omni_httpd) | [`omnigres`](/e/omni) | HTTP server |
| 2962 | [`omni_id`](/e/omni_id) | [`omnigres`](/e/omni) | Identity types |
| 2963 | [`omni_json`](/e/omni_json) | [`omnigres`](/e/omni) | JSON toolkit |
| 2964 | [`omni_kube`](/e/omni_kube) | [`omnigres`](/e/omni) | Kubernetes (k8s) integration |
| 2965 | [`omni_ledger`](/e/omni_ledger) | [`omnigres`](/e/omni) | Financial ledger |
| 2966 | [`omni_manifest`](/e/omni_manifest) | [`omnigres`](/e/omni) | Package installation manifests |
| 2967 | [`omni_mimetypes`](/e/omni_mimetypes) | [`omnigres`](/e/omni) | MIME types |
| 2968 | [`omni_os`](/e/omni_os) | [`omnigres`](/e/omni) | Operating system integration |
| 2969 | [`omni_polyfill`](/e/omni_polyfill) | [`omnigres`](/e/omni) | Postgres API polyfills |
| 2970 | [`omni_python`](/e/omni_python) | [`omnigres`](/e/omni) | First-class Python support |
| 2971 | [`omni_regex`](/e/omni_regex) | [`omnigres`](/e/omni) | PCRE-compatible regular expressions |
| 2972 | [`omni_rest`](/e/omni_rest) | [`omnigres`](/e/omni) | REST API toolkit (with PostgREST support) |
| 2973 | [`omni_schema`](/e/omni_schema) | [`omnigres`](/e/omni) | Advanced schema management tooling |
| 2974 | [`omni_seq`](/e/omni_seq) | [`omnigres`](/e/omni) | Distributed integer sequences |
| 2975 | [`omni_service`](/e/omni_service) | [`omnigres`](/e/omni) | Service management |
| 2976 | [`omni_session`](/e/omni_session) | [`omnigres`](/e/omni) | Session management |
| 2977 | [`omni_sql`](/e/omni_sql) | [`omnigres`](/e/omni) | Programmatic SQL manipulation |
| 2979 | [`omni_sqlite`](/e/omni_sqlite) | [`omnigres`](/e/omni) | Embedded SQLite |
| 2980 | [`omni_test`](/e/omni_test) | [`omnigres`](/e/omni) | Testing framework |
| 2981 | [`omni_txn`](/e/omni_txn) | [`omnigres`](/e/omni) | Transaction management |
| 2982 | [`omni_types`](/e/omni_types) | [`omnigres`](/e/omni) | Advanced types |
| 2983 | [`omni_var`](/e/omni_var) | [`omnigres`](/e/omni) | Scoped variables |
| 2984 | [`omni_vfs`](/e/omni_vfs) | [`omnigres`](/e/omni) | Virtual File System |
| 2985 | [`omni_vfs_types_v1`](/e/omni_vfs_types_v1) | [`omnigres`](/e/omni) | Virtual File System types (v1) |
| 2986 | [`omni_web`](/e/omni_web) | [`omnigres`](/e/omni) | Common web stack primitives |
| 2987 | [`omni_worker`](/e/omni_worker) | [`omnigres`](/e/omni) | Generalized worker pool |
| 2988 | [`omni_xml`](/e/omni_xml) | [`omnigres`](/e/omni) | XML toolkit |
| 2989 | [`omni_yaml`](/e/omni_yaml) | [`omnigres`](/e/omni) | YAML toolkit |
| 4640 | [`omnisketch`](/e/omnisketch) | [`omnisketch`](/e/omnisketch) | data structure for on-line agg of data into approximate sketch |
| 8610 | [`oracle_fdw`](/e/oracle_fdw) | [`oracle_fdw`](/e/oracle_fdw) | foreign data wrapper for Oracle access |
| 9100 | [`orafce`](/e/orafce) | [`orafce`](/e/orafce) | Functions and operators that emulate a subset of functions and packages from the Oracle RDBMS |
| 2920 | [`orioledb`](/e/orioledb) | [`orioledb`](/e/orioledb) | OrioleDB, the next generation transactional engine |
| 6900 | [`pageinspect`](/e/pageinspect) | [`pageinspect`](/e/pageinspect) | inspect the contents of database pages at a low level |
| 7990 | [`passwordcheck`](/e/passwordcheck) | [`passwordcheck`](/e/passwordcheck_cracklib) | checks user passwords and reject weak password |
| 7000 | [`passwordcheck_cracklib`](/e/passwordcheck_cracklib) | [`passwordcheck`](/e/passwordcheck_cracklib) | Strengthen PostgreSQL user password checks with cracklib |
| 1030 | [`periods`](/e/periods) | [`periods`](/e/periods) | Provide Standard SQL functionality for PERIODs and SYSTEM VERSIONING |
| 4550 | [`permuteseq`](/e/permuteseq) | [`permuteseq`](/e/permuteseq) | Pseudo-randomly permute sequences with a format-preserving encryption on elements |
| 1880 | [`pg4ml`](/e/pg4ml) | [`pg4ml`](/e/pg4ml) | Machine learning framework for PostgreSQL |
| 7100 | [`pg_auth_mon`](/e/pg_auth_mon) | [`pg_auth_mon`](/e/pg_auth_mon) | monitor connection attempts per user |
| 1100 | [`pg_background`](/e/pg_background) | [`pg_background`](/e/pg_background) | Run SQL queries in the background |
| 2120 | [`pg_bigm`](/e/pg_bigm) | [`pg_bigm`](/e/pg_bigm) | create 2-gram (bigram) index for faster full text search. |
| 6930 | [`pg_buffercache`](/e/pg_buffercache) | [`pg_buffercache`](/e/pg_buffercache) | examine the shared buffer cache |
| 9830 | [`pg_bulkload`](/e/pg_bulkload) | [`pg_bulkload`](/e/pg_bulkload) | pg_bulkload is a high speed data loading utility for PostgreSQL |
| 5160 | [`pg_catcheck`](/e/pg_catcheck) | [`pg_catcheck`](/e/pg_catcheck) | Diagnosing system catalog corruption |
| 5220 | [`pg_cheat_funcs`](/e/pg_cheat_funcs) | [`pg_cheat_funcs`](/e/pg_cheat_funcs) | Provides cheat (but useful) functions |
| 5110 | [`pg_checksums`](/e/pg_checksums) | [`pg_checksums`](/e/pg_checksums) | Activate/deactivate/verify checksums in offline Postgres clusters |
| 5070 | [`pg_cooldown`](/e/pg_cooldown) | [`pg_cooldown`](/e/pg_cooldown) | remove buffered pages for specific relations |
| 5210 | [`pg_crash`](/e/pg_crash) | [`pg_crash`](/e/pg_crash) | Send random signals to random processes |
| 1070 | [`pg_cron`](/e/pg_cron) | [`pg_cron`](/e/pg_cron) | Job scheduler for PostgreSQL |
| 4090 | [`pg_curl`](/e/pg_curl) | [`pg_curl`](/e/pg_curl) | Run curl actions for data transfer in URL syntax |
| 5050 | [`pg_dirtyread`](/e/pg_dirtyread) | [`pg_dirtyread`](/e/pg_dirtyread) | Read dead but unvacuumed rows from table |
| 3830 | [`pg_duration`](/e/pg_duration) | [`pg_duration`](/e/pg_duration) | data type for representing durations |
| 9820 | [`pg_fact_loader`](/e/pg_fact_loader) | [`pg_fact_loader`](/e/pg_fact_loader) | build fact tables with Postgres |
| 9530 | [`pg_failover_slots`](/e/pg_failover_slots) | [`pg_failover_slots`](/e/pg_failover_slots) | PG Failover Slots extension |
| 6950 | [`pg_freespacemap`](/e/pg_freespacemap) | [`pg_freespacemap`](/e/pg_freespacemap) | examine the free space map (FSM) |
| 1590 | [`pg_geohash`](/e/pg_geohash) | [`pg_geohash`](/e/pg_geohash) | Handle geohash based functionality for spatial coordinates |
| 4560 | [`pg_hashids`](/e/pg_hashids) | [`pg_hashids`](/e/pg_hashids) | Short unique id generator for PostgreSQL, using hashids |
| 2820 | [`pg_hint_plan`](/e/pg_hint_plan) | [`pg_hint_plan`](/e/pg_hint_plan) | Give PostgreSQL ability to manually force some decisions in execution plans. |
| 2880 | [`pg_incremental`](/e/pg_incremental) | [`pg_incremental`](/e/pg_incremental) | Incremental Processing by Crunchy Data |
| 2870 | [`pg_ivm`](/e/pg_ivm) | [`pg_ivm`](/e/pg_ivm) | incremental view maintenance on PostgreSQL |
| 6890 | [`pg_logicalinspect`](/e/pg_logicalinspect) | [`pg_logicalinspect`](/e/pg_logicalinspect) | Logical decoding components inspection |
| 4770 | [`pg_math`](/e/pg_math) | [`pg_math`](/e/pg_math) | GSL statistical functions for postgresql |
| 4080 | [`pg_net`](/e/pg_net) | [`pg_net`](/e/pg_net) | Async HTTP Requests |
| 5200 | [`pg_orphaned`](/e/pg_orphaned) | [`pg_orphaned`](/e/pg_orphaned) | Deal with orphaned files |
| 6880 | [`pg_overexplain`](/e/pg_overexplain) | [`pg_overexplain`](/e/pg_overexplain) | Allow EXPLAIN to dump even more details |
| 2510 | [`pg_partman`](/e/pg_partman) | [`pg_partman`](/e/pg_partman) | Extension to manage partitioned tables by time or ID |
| 5890 | [`pg_prewarm`](/e/pg_prewarm) | [`pg_prewarm`](/e/pg_prewarm) | prewarm relation data |
| 6320 | [`pg_proctab`](/e/pg_proctab) | [`pgnodemx`](/e/pgnodemx) | PostgreSQL extension to access the OS process table |
| 6000 | [`pg_profile`](/e/pg_profile) | [`pg_profile`](/e/pg_profile) | PostgreSQL load profile repository and report builder |
| 4260 | [`pg_protobuf`](/e/pg_protobuf) | [`pg_protobuf`](/e/pg_protobuf) | Protobuf support for PostgreSQL |
| 6240 | [`pg_qualstats`](/e/pg_qualstats) | [`pg_qualstats`](/e/pg_qualstats) | An extension collecting statistics about quals |
| 3720 | [`pg_rational`](/e/pg_rational) | [`pg_rational`](/e/pg_rational) | bigint fractions |
| 4300 | [`pg_readme`](/e/pg_readme) | [`pg_readme`](/e/pg_readme) | Generate a README.md document for a database extension or schema |
| 4301 | [`pg_readme_test_extension`](/e/pg_readme_test_extension) | [`pg_readme`](/e/pg_readme) | Test generating a README.md document for extension or schema |
| 5120 | [`pg_readonly`](/e/pg_readonly) | [`pg_readonly`](/e/pg_readonly) | cluster database read only |
| 6380 | [`pg_relusage`](/e/pg_relusage) | [`pg_relusage`](/e/pg_relusage) | Log all the queries that reference a particular column |
| 5010 | [`pg_repack`](/e/pg_repack) | [`pg_repack`](/e/pg_repack) | Reorganize tables in PostgreSQL databases with minimal locks |
| 5020 | [`pg_rewrite`](/e/pg_rewrite) | [`pg_rewrite`](/e/pg_rewrite) | Tool allows read write to the table during the rewriting |
| 3880 | [`pg_rrule`](/e/pg_rrule) | [`pg_rrule`](/e/pg_rrule) | RRULE field type for PostgreSQL |
| 5810 | [`pg_savior`](/e/pg_savior) | [`pg_savior`](/e/pg_savior) | Postgres extension to save OOPS mistakes |
| 6210 | [`pg_show_plans`](/e/pg_show_plans) | [`pg_show_plans`](/e/pg_show_plans) | show query plans of all currently running SQL statements |
| 1840 | [`pg_similarity`](/e/pg_similarity) | [`pg_similarity`](/e/pg_similarity) | support similarity queries |
| 7170 | [`pg_snakeoil`](/e/pg_snakeoil) | [`pg_snakeoil`](/e/pg_snakeoil) | The PostgreSQL Antivirus |
| 3590 | [`pg_sphere`](/e/pg_sphere) | [`pgsphere`](/e/pg_sphere) | spherical objects with useful functions, operators and index support |
| 5040 | [`pg_squeeze`](/e/pg_squeeze) | [`pg_squeeze`](/e/pg_squeeze) | A tool to remove unused space from a relation. |
| 6220 | [`pg_stat_kcache`](/e/pg_stat_kcache) | [`pg_stat_kcache`](/e/pg_stat_kcache) | Kernel statistics gathering |
| 6230 | [`pg_stat_monitor`](/e/pg_stat_monitor) | [`pg_stat_monitor`](/e/pg_stat_monitor) | The pg_stat_monitor is a PostgreSQL Query Performance Monitoring tool, based on PostgreSQL contrib module pg_stat_statements. pg_stat_monitor provides aggregated statistics, client information, plan details including plan, and histogram information. |
| 6990 | [`pg_stat_statements`](/e/pg_stat_statements) | [`pg_stat_statements`](/e/pg_stat_statements) | track planning and execution statistics of all SQL statements executed |
| 9130 | [`pg_statement_rollback`](/e/pg_statement_rollback) | [`pg_statement_rollback`](/e/pg_statement_rollback) | Server side rollback at statement level for PostgreSQL like Oracle or DB2 |
| 6250 | [`pg_store_plans`](/e/pg_store_plans) | [`pg_store_plans`](/e/pg_store_plans) | track plan statistics of all SQL statements executed |
| 2530 | [`pg_strom`](/e/pg_strom) | [`pg_strom`](/e/pg_strom) | PG-Strom - big-data processing acceleration using GPU and NVME |
| 5990 | [`pg_surgery`](/e/pg_surgery) | [`pg_surgery`](/e/pg_surgery) | extension to perform surgery on a damaged relation |
| 1080 | [`pg_task`](/e/pg_task) | [`pg_task`](/e/pg_task) | execute any sql command at any specific time at background |
| 7060 | [`pg_tde`](/e/pg_tde) | [`pg_tde`](/e/pg_tde) | Percona pg_tde access method |
| 3000 | [`pg_tle`](/e/pg_tle) | [`pg_tle`](/e/pg_tle) | Trusted Language Extensions for PostgreSQL |
| 6010 | [`pg_tracing`](/e/pg_tracing) | [`pg_tracing`](/e/pg_tracing) | Distributed Tracing for PostgreSQL |
| 2190 | [`pg_trgm`](/e/pg_trgm) | [`pg_trgm`](/e/pg_trgm) | text similarity measurement and index searching based on trigrams |
| 4540 | [`pg_uuidv7`](/e/pg_uuidv7) | [`pg_uuidv7`](/e/pg_uuidv7) | Create UUIDv7 values in postgres |
| 6960 | [`pg_visibility`](/e/pg_visibility) | [`pg_visibility`](/e/pg_visibility) | examine the visibility map (VM) and page-level visibility info |
| 6270 | [`pg_wait_sampling`](/e/pg_wait_sampling) | [`pg_wait_sampling`](/e/pg_wait_sampling) | sampling based statistics of wait events |
| 6940 | [`pg_walinspect`](/e/pg_walinspect) | [`pg_walinspect`](/e/pg_walinspect) | functions to inspect contents of PostgreSQL Write-Ahead Log |
| 9550 | [`pgactive`](/e/pgactive) | [`pgactive`](/e/pgactive) | Active-Active Replication Extension for PostgreSQL |
| 5880 | [`pgagent`](/e/pgagent) | [`pgagent`](/e/pgagent) | A PostgreSQL job scheduler |
| 7080 | [`pgaudit`](/e/pgaudit) | [`pgaudit`](/e/pgaudit) | provides auditing functionality |
| 7090 | [`pgauditlogtofile`](/e/pgauditlogtofile) | [`pgauditlogtofile`](/e/pgauditlogtofile) | pgAudit addon to redirect audit log to an independent file |
| 5150 | [`pgautofailover`](/e/pgautofailover) | [`pgautofailover`](/e/pgautofailover) | pg_auto_failover |
| 7980 | [`pgcrypto`](/e/pgcrypto) | [`pgcrypto`](/e/pgcrypto) | cryptographic functions |
| 7120 | [`pgcryptokey`](/e/pgcryptokey) | [`pgcryptokey`](/e/pgcryptokey) | cryptographic key management |
| 7180 | [`pgextwlist`](/e/pgextwlist) | [`pgextwlist`](/e/pgextwlist) | PostgreSQL Extension Whitelisting |
| 5060 | [`pgfincore`](/e/pgfincore) | [`pgfincore`](/e/pgfincore) | examine and manage the os buffer cache |
| 4150 | [`pgjq`](/e/pgjq) | [`pgjq`](/e/pgjq) | Use jq in Postgres |
| 9520 | [`pgl_ddl_deploy`](/e/pgl_ddl_deploy) | [`pgl_ddl_deploy`](/e/pgl_ddl_deploy) | automated ddl deployment using pglogical |
| 9500 | [`pglogical`](/e/pglogical) | [`pglogical`](/e/pglogical) | PostgreSQL Logical Replication |
| 9501 | [`pglogical_origin`](/e/pglogical_origin) | [`pglogical`](/e/pglogical) | Dummy extension for compatibility when upgrading from Postgres 9.4 |
| 9510 | [`pglogical_ticker`](/e/pglogical_ticker) | [`pglogical_ticker`](/e/pglogical_ticker) | Have an accurate view on pglogical replication delay |
| 9410 | [`pgmemcache`](/e/pgmemcache) | [`pgmemcache`](/e/pgmemcache) | memcached interface |
| 6350 | [`pgmeminfo`](/e/pgmeminfo) | [`pgmeminfo`](/e/pgmeminfo) | show memory usage |
| 3700 | [`pgmp`](/e/pgmp) | [`pgmp`](/e/pgmp) | Multiple Precision Arithmetic extension |
| 6310 | [`pgnodemx`](/e/pgnodemx) | [`pgnodemx`](/e/pgnodemx) | Capture node OS metrics via SQL queries |
| 9980 | [`pgoutput`](/e/pgoutput) | [`pgoutput`](/e/pgoutput) | Logical Replication output plugin |
| 4230 | [`pgpcre`](/e/pgpcre) | [`pgpcre`](/e/pgpcre) | Perl Compatible Regular Expression functions |
| 3530 | [`pgpdf`](/e/pgpdf) | [`pgpdf`](/e/pgpdf) | PDF type with meta admin & Full-Text Search |
| 5900 | [`pgpool_adm`](/e/pgpool_adm) | [`pgpool`](/e/pgpool_adm) | Administrative functions for pgPool |
| 5910 | [`pgpool_recovery`](/e/pgpool_recovery) | [`pgpool`](/e/pgpool_adm) | recovery functions for pgpool-II for V4.3 |
| 5920 | [`pgpool_regclass`](/e/pgpool_regclass) | [`pgpool`](/e/pgpool_adm) | replacement for regclass |
| 2910 | [`pgq`](/e/pgq) | [`pgq`](/e/pgq) | Generic queue for PostgreSQL |
| 4250 | [`pgqr`](/e/pgqr) | [`pgqr`](/e/pgqr) | QR Code generator from PostgreSQL |
| 2110 | [`pgroonga`](/e/pgroonga) | [`pgroonga`](/e/pgroonga) | Use Groonga as index, fast full text search platform for all languages! |
| 2111 | [`pgroonga_database`](/e/pgroonga_database) | [`pgroonga`](/e/pgroonga) | PGroonga database management module |
| 6910 | [`pgrowlocks`](/e/pgrowlocks) | [`pgrowlocks`](/e/pgrowlocks) | show row-level locking information |
| 6280 | [`pgsentinel`](/e/pgsentinel) | [`pgsentinel`](/e/pgsentinel) | active session history |
| 7020 | [`pgsodium`](/e/pgsodium) | [`pgsodium`](/e/pgsodium) | Postgres extension for libsodium functions |
| 8540 | [`pgspider_ext`](/e/pgspider_ext) | [`pgspider_ext`](/e/pgspider_ext) | foreign-data wrapper for remote PGSpider servers |
| 6970 | [`pgstattuple`](/e/pgstattuple) | [`pgstattuple`](/e/pgstattuple) | show tuple-level statistics |
| 3200 | [`pgtap`](/e/pgtap) | [`pgtap`](/e/pgtap) | Unit testing for PostgreSQL |
| 9110 | [`pgtt`](/e/pgtt) | [`pgtt`](/e/pgtt) | Extension to add Global Temporary Tables feature to PostgreSQL |
| 4460 | [`pguecc`](/e/pguecc) | [`pg_ecdsa`](/e/pguecc) | uECC bindings for Postgres |
| 2850 | [`plan_filter`](/e/plan_filter) | [`pg_plan_filter`](/e/plan_filter) | filter statements by their execution plans. |
| 3050 | [`pldbgapi`](/e/pldbgapi) | [`pldebugger`](/e/pldbgapi) | server-side support for debugging PL/pgSQL functions |
| 3020 | [`pllua`](/e/pllua) | [`pllua`](/e/pllua) | Lua as a procedural language |
| 3030 | [`plluau`](/e/plluau) | [`pllua`](/e/pllua) | Lua as an untrusted procedural language |
| 3260 | [`plperl`](/e/plperl) | [`plperl`](/e/plperl) | PL/Perl procedural language |
| 3270 | [`plperlu`](/e/plperlu) | [`plperlu`](/e/plperlu) | PL/PerlU untrusted procedural language |
| 3280 | [`plpgsql`](/e/plpgsql) | [`plpgsql`](/e/plpgsql) | PL/pgSQL procedural language |
| 3060 | [`plpgsql_check`](/e/plpgsql_check) | [`plpgsql_check`](/e/plpgsql_check) | extended check for plpgsql functions |
| 3070 | [`plprofiler`](/e/plprofiler) | [`plprofiler`](/e/plprofiler) | server-side support for profiling PL/pgSQL functions |
| 2520 | [`plproxy`](/e/plproxy) | [`plproxy`](/e/plproxy) | Database partitioning implemented as procedural language |
| 3290 | [`plpython3u`](/e/plpython3u) | [`plpython3u`](/e/plpython3u) | PL/Python3U untrusted procedural language |
| 3100 | [`plr`](/e/plr) | [`plr`](/e/plr) | load R interpreter and execute R script from within a database |
| 3080 | [`plsh`](/e/plsh) | [`plsh`](/e/plsh) | PL/sh procedural language |
| 3240 | [`pltcl`](/e/pltcl) | [`pltcl`](/e/pltcl) | PL/Tcl procedural language |
| 3250 | [`pltclu`](/e/pltclu) | [`pltcl`](/e/pltcl) | PL/TclU untrusted procedural language |
| 1520 | [`pointcloud`](/e/pointcloud) | [`pointcloud`](/e/pointcloud) | data type for lidar point clouds |
| 1521 | [`pointcloud_postgis`](/e/pointcloud_postgis) | [`pointcloud`](/e/pointcloud) | integration for pointcloud LIDAR data and PostGIS geometry data |
| 1500 | [`postgis`](/e/postgis) | [`postgis`](/e/postgis) | PostGIS geometry and geography spatial types and functions |
| 1502 | [`postgis_raster`](/e/postgis_raster) | [`postgis`](/e/postgis) | PostGIS raster types and functions |
| 1503 | [`postgis_sfcgal`](/e/postgis_sfcgal) | [`postgis`](/e/postgis) | PostGIS SFCGAL functions |
| 1504 | [`postgis_tiger_geocoder`](/e/postgis_tiger_geocoder) | [`postgis`](/e/postgis) | PostGIS tiger geocoder and reverse geocoder |
| 1501 | [`postgis_topology`](/e/postgis_topology) | [`postgis`](/e/postgis) | PostGIS topology spatial types and functions |
| 8990 | [`postgres_fdw`](/e/postgres_fdw) | [`postgres_fdw`](/e/postgres_fdw) | foreign-data wrapper for remote PostgreSQL servers |
| 5170 | [`pre_prepare`](/e/pre_prepare) | [`preprepare`](/e/pre_prepare) | Pre Prepare your Statement server side |
| 3500 | [`prefix`](/e/prefix) | [`pg_prefix`](/e/prefix) | Prefix Range module for PostgreSQL |
| 5090 | [`prioritize`](/e/prioritize) | [`pg_prioritize`](/e/prioritize) | get and set the priority of PostgreSQL backends |
| 1540 | [`q3c`](/e/q3c) | [`q3c`](/e/q3c) | q3c sky indexing plugin |
| 4610 | [`quantile`](/e/quantile) | [`quantile`](/e/quantile) | Quantile aggregation function |
| 4780 | [`random`](/e/random) | [`pg_random`](/e/random) | random data generator |
| 8720 | [`redis`](/e/redis) | [`pg_redis_pubsub`](/e/redis) | Send redis pub/sub messages to Redis from PostgreSQL Directly |
| 8710 | [`redis_fdw`](/e/redis_fdw) | [`redis_fdw`](/e/redis_fdw) | Foreign data wrapper for querying a Redis server |
| 4880 | [`refint`](/e/refint) | [`refint`](/e/refint) | functions for implementing referential integrity (obsolete) |
| 9710 | [`repmgr`](/e/repmgr) | [`repmgr`](/e/repmgr) | Replication manager for PostgreSQL |
| 3570 | [`roaringbitmap`](/e/roaringbitmap) | [`roaringbitmap`](/e/roaringbitmap) | support for Roaring Bitmaps |
| 2780 | [`rum`](/e/rum) | [`rum`](/e/rum) | RUM index access method |
| 5820 | [`safeupdate`](/e/safeupdate) | [`safeupdate`](/e/safeupdate) | Require criteria for UPDATE and DELETE |
| 3940 | [`seg`](/e/seg) | [`seg`](/e/seg) | data type for representing line segments or floating-point intervals |
| 3510 | [`semver`](/e/semver) | [`pg_semver`](/e/semver) | Semantic version data type |
| 7960 | [`sepgsql`](/e/sepgsql) | [`sepgsql`](/e/sepgsql) | label-based mandatory access control (MAC) based on SELinux security policy. |
| 4570 | [`sequential_uuids`](/e/sequential_uuids) | [`sequential_uuids`](/e/sequential_uuids) | generator of sequential UUIDs |
| 9120 | [`session_variable`](/e/session_variable) | [`session_variable`](/e/session_variable) | Registration and manipulation of session variables and constants |
| 7160 | [`set_user`](/e/set_user) | [`set_user`](/e/set_user) | similar to SET ROLE but with added logging |
| 4440 | [`shacrypt`](/e/shacrypt) | [`shacrypt`](/e/shacrypt) | Implements SHA256-CRYPT and SHA512-CRYPT password encryption schemes |
| 1850 | [`smlar`](/e/smlar) | [`smlar`](/e/smlar) | Effective similarity search |
| 9400 | [`spat`](/e/spat) | [`spat`](/e/spat) | Redis-like In-Memory DB Embedded in Postgres |
| 8640 | [`sqlite_fdw`](/e/sqlite_fdw) | [`sqlite_fdw`](/e/sqlite_fdw) | SQLite Foreign Data Wrapper |
| 6920 | [`sslinfo`](/e/sslinfo) | [`sslinfo`](/e/sslinfo) | information about SSL certificates |
| 7200 | [`sslutils`](/e/sslutils) | [`sslutils`](/e/sslutils) | A Postgres extension for managing SSL certificates through SQL |
| 7030 | [`supabase_vault`](/e/supabase_vault) | [`pg_vault`](/e/supabase_vault) | Supabase Vault Extension |
| 7010 | [`supautils`](/e/supautils) | [`supautils`](/e/supautils) | Extension that secures a cluster on a cloud environment |
| 6290 | [`system_stats`](/e/system_stats) | [`system_stats`](/e/system_stats) | EnterpriseDB system statistics for PostgreSQL |
| 5840 | [`table_log`](/e/table_log) | [`table_log`](/e/table_log) | record table modification logs and PITR for table/row |
| 2590 | [`tablefunc`](/e/tablefunc) | [`tablefunc`](/e/tablefunc) | functions that manipulate whole tables, including crosstab |
| 4920 | [`tcn`](/e/tcn) | [`tcn`](/e/tcn) | Triggered change notifications |
| 4700 | [`tdigest`](/e/tdigest) | [`tdigest`](/e/tdigest) | Provides tdigest aggregate function. |
| 8620 | [`tds_fdw`](/e/tds_fdw) | [`tds_fdw`](/e/tds_fdw) | Foreign data wrapper for querying a TDS database (Sybase or Microsoft SQL Server) |
| 1040 | [`temporal_tables`](/e/temporal_tables) | [`temporal_tables`](/e/temporal_tables) | temporal tables |
| 9970 | [`test_decoding`](/e/test_decoding) | [`test_decoding`](/e/test_decoding) | SQL-based test/example module for WAL logical decoding |
| 1000 | [`timescaledb`](/e/timescaledb) | [`timescaledb`](/e/timescaledb) | Enables scalable inserts and complex queries for time-series data |
| 3890 | [`timestamp9`](/e/timestamp9) | [`timestamp9`](/e/timestamp9) | timestamp nanosecond resolution |
| 6360 | [`toastinfo`](/e/toastinfo) | [`toastinfo`](/e/toastinfo) | show details on toasted datums |
| 4600 | [`topn`](/e/topn) | [`topn`](/e/topn) | type for top-n JSONB |
| 4910 | [`tsm_system_rows`](/e/tsm_system_rows) | [`tsm_system_rows`](/e/tsm_system_rows) | TABLESAMPLE method which accepts number of rows as a limit |
| 4890 | [`tsm_system_time`](/e/tsm_system_time) | [`tsm_system_time`](/e/tsm_system_time) | TABLESAMPLE method which accepts time in milliseconds as a limit |
| 3730 | [`uint`](/e/uint) | [`pguint`](/e/uint) | unsigned integer types |
| 3740 | [`uint128`](/e/uint128) | [`pg_uint128`](/e/uint128) | Native uint128 type |
| 4990 | [`unaccent`](/e/unaccent) | [`unaccent`](/e/unaccent) | text search dictionary that removes accents |
| 3520 | [`unit`](/e/unit) | [`pgunit`](/e/unit) | SI units extension |
| 3840 | [`uri`](/e/uri) | [`pg_uri`](/e/uri) | URI Data type for PostgreSQL |
| 4190 | [`url_encode`](/e/url_encode) | [`url_encode`](/e/url_encode) | url_encode, url_decode functions |
| 4930 | [`uuid-ossp`](/e/uuid-ossp) | [`uuid-ossp`](/e/uuid-ossp) | generate universally unique identifiers (UUIDs) |
| 4660 | [`vasco`](/e/vasco) | [`vasco`](/e/vasco) | discover hidden correlations in your data with MIC |
| 1800 | [`vector`](/e/vector) | [`pgvector`](/e/vector) | vector data type and ivfflat and hnsw access methods |
| 9630 | [`wal2json`](/e/wal2json) | [`wal2json`](/e/wal2json) | Changing data capture in JSON format |
| 9640 | [`wal2mongo`](/e/wal2mongo) | [`wal2mongo`](/e/wal2mongo) | PostgreSQL logical decoding output plugin for MongoDB |
| 4670 | [`xicor`](/e/xicor) | [`pgxicor`](/e/xicor) | XI Correlation Coefficient in Postgres |
| 3990 | [`xml2`](/e/xml2) | [`xml2`](/e/xml2) | XPath querying and XSLT |
| 4430 | [`xxhash`](/e/xxhash) | [`pg_xxhash`](/e/xxhash) | xxhash functions for PostgreSQL |
| 2130 | [`zhparser`](/e/zhparser) | [`zhparser`](/e/zhparser) | a parser for full-text search of Chinese |
| 4030 | [`zstd`](/e/zstd) | [`pg_zstd`](/e/zstd) | Zstandard compression algorithm implementation in PostgreSQL |

## SQL

{{< language "SQL" >}} {{< badge content="37 Extensions" color="gray" >}}


Pure SQL extensions and functions

| ID | Extension | Package | Description |
|:---:|:---|:---|:---|
| 8800 | [`aws_s3`](/e/aws_s3) | [`aws_s3`](/e/aws_s3) | aws_s3 postgres extension to import/export data from/to s3 |
| 4320 | [`data_historization`](/e/data_historization) | [`data_historization`](/e/data_historization) | PLPGSQL Script to historize data in partitionned table |
| 9540 | [`db_migrator`](/e/db_migrator) | [`db_migrator`](/e/db_migrator) | Tools to migrate other databases to PostgreSQL |
| 4310 | [`ddl_historization`](/e/ddl_historization) | [`ddl_historization`](/e/ddl_historization) | Historize the ddl changes inside PostgreSQL database |
| 5080 | [`ddlx`](/e/ddlx) | [`pg_ddlx`](/e/ddlx) | DDL eXtractor functions |
| 3870 | [`debversion`](/e/debversion) | [`debversion`](/e/debversion) | Debian version number data type |
| 1050 | [`emaj`](/e/emaj) | [`emaj`](/e/emaj) | Enables fine-grained write logging and time travel on subsets of the database. |
| 1560 | [`geoip`](/e/geoip) | [`geoip`](/e/geoip) | IP-based geolocation query |
| 2840 | [`index_advisor`](/e/index_advisor) | [`index_advisor`](/e/index_advisor) | Query index advisor |
| 3611 | [`l10n_table_dependent_extension`](/e/l10n_table_dependent_extension) | [`pg_xenophile`](/e/pg_xenophile) | PostgreSQL l10n toolbox |
| 6300 | [`meta`](/e/meta) | [`pg_meta`](/e/meta) | Normalized, friendlier system catalog for PostgreSQL |
| 9700 | [`mimeo`](/e/mimeo) | [`mimeo`](/e/mimeo) | Extension for specialized, per-table replication between PostgreSQL instances |
| 6800 | [`pagevis`](/e/pagevis) | [`pagevis`](/e/pagevis) | Visualise database pages in ascii code |
| 7190 | [`pg_auditor`](/e/pg_auditor) | [`pg_auditor`](/e/pg_auditor) | Audit data changes and provide flashback ability |
| 9260 | [`pg_dbms_job`](/e/pg_dbms_job) | [`pg_dbms_job`](/e/pg_dbms_job) | Extension to add Oracle DBMS_JOB full compatibility to PostgreSQL |
| 9250 | [`pg_dbms_lock`](/e/pg_dbms_lock) | [`pg_dbms_lock`](/e/pg_dbms_lock) | Extension to add Oracle DBMS_LOCK full compatibility to PostgreSQL |
| 9240 | [`pg_dbms_metadata`](/e/pg_dbms_metadata) | [`pg_dbms_metadata`](/e/pg_dbms_metadata) | Extension to add Oracle DBMS_METADATA compatibility to PostgreSQL |
| 5830 | [`pg_drop_events`](/e/pg_drop_events) | [`pg_drop_events`](/e/pg_drop_events) | logs transaction ids of drop table, drop column, drop materialized view statements |
| 4220 | [`pg_extra_time`](/e/pg_extra_time) | [`pg_extra_time`](/e/pg_extra_time) | Some date time functions and operators that, |
| 2500 | [`pg_fkpart`](/e/pg_fkpart) | [`pg_fkpart`](/e/pg_fkpart) | Table partitioning by foreign key utility |
| 4180 | [`pg_html5_email_address`](/e/pg_html5_email_address) | [`pg_html5_email_address`](/e/pg_html5_email_address) | PostgreSQL email validation that is consistent with the HTML5 spec |
| 7130 | [`pg_jobmon`](/e/pg_jobmon) | [`pg_jobmon`](/e/pg_jobmon) | Extension for logging and monitoring functions in PostgreSQL |
| 5140 | [`pg_permissions`](/e/pg_permissions) | [`pg_permissions`](/e/pg_permissions) | view object permissions and compare them with the desired state |
| 6330 | [`pg_sqlog`](/e/pg_sqlog) | [`pg_sqlog`](/e/pg_sqlog) | Provide SQL interface to logs |
| 6260 | [`pg_track_settings`](/e/pg_track_settings) | [`pg_track_settings`](/e/pg_track_settings) | Track settings changes |
| 5180 | [`pg_upless`](/e/pg_upless) | [`pg_upless`](/e/pg_upless) | Detect Useless UPDATE |
| 3610 | [`pg_xenophile`](/e/pg_xenophile) | [`pg_xenophile`](/e/pg_xenophile) | More than the bare necessities for PostgreSQL i18n and l10n. |
| 8650 | [`pgbouncer_fdw`](/e/pgbouncer_fdw) | [`pgbouncer_fdw`](/e/pgbouncer_fdw) | Extension for querying PgBouncer stats from normal SQL views & running pgbouncer commands from normal SQL functions |
| 5190 | [`pgcozy`](/e/pgcozy) | [`pgcozy`](/e/pgcozy) | Pre-warming shared buffers according to previous pg_buffercache snapshots for PostgreSQL. |
| 3580 | [`pgfaceting`](/e/pgfaceting) | [`pgfaceting`](/e/pgfaceting) | fast faceting queries using an inverted index |
| 4160 | [`pgjwt`](/e/pgjwt) | [`pgjwt`](/e/pgjwt) | JSON Web Token API for Postgresql |
| 2900 | [`pgmq`](/e/pgmq) | [`pgmq`](/e/pgmq) | A lightweight message queue. Like AWS SQS and RSMQ but on Postgres. |
| 4200 | [`pgsql_tweaks`](/e/pgsql_tweaks) | [`pgsql_tweaks`](/e/pgsql_tweaks) | Some functions and views for daily usage |
| 4330 | [`schedoc`](/e/schedoc) | [`pg_schedoc`](/e/schedoc) | Cross documentation between Django and DBT projects |
| 4470 | [`sparql`](/e/sparql) | [`pgsparql`](/e/sparql) | Query SPARQL datasource with SQL |
| 1060 | [`table_version`](/e/table_version) | [`table_version`](/e/table_version) | PostgreSQL table versioning extension |
| 1020 | [`timeseries`](/e/timeseries) | [`pg_timeseries`](/e/timeseries) | Convenience API for time series stack |

## Rust

{{< language "Rust" >}} {{< badge content="33 Extensions" color="gray" >}}


Extensions written in Rust with the pgrx framework

| ID | Extension | Package | Description |
|:---:|:---|:---|:---|
| 7050 | [`anon`](/e/anon) | [`pg_anon`](/e/anon) | PostgreSQL Anonymizer (anon) extension |
| 4850 | [`convert`](/e/convert) | [`pg_convert`](/e/convert) | conversion functions for spatial, routing and other specialized uses |
| 6370 | [`explain_ui`](/e/explain_ui) | [`pg_explain_ui`](/e/explain_ui) | easily jump into a visual plan UI for any SQL query |
| 2420 | [`pg_analytics`](/e/pg_analytics) | [`pg_analytics`](/e/pg_analytics) | Postgres for analytics, powered by DuckDB |
| 4830 | [`pg_base58`](/e/pg_base58) | [`pg_base58`](/e/pg_base58) | Base58 Encoder/Decoder Extension for PostgreSQL |
| 2140 | [`pg_bestmatch`](/e/pg_bestmatch) | [`pg_bestmatch`](/e/pg_bestmatch) | Generate BM25 sparse vector inside PostgreSQL |
| 2930 | [`pg_cardano`](/e/pg_cardano) | [`pg_cardano`](/e/pg_cardano) | A suite of Cardano-related tools |
| 2790 | [`pg_graphql`](/e/pg_graphql) | [`pg_graphql`](/e/pg_graphql) | Add in-database GraphQL support |
| 4500 | [`pg_idkit`](/e/pg_idkit) | [`pg_idkit`](/e/pg_idkit) | multi-tool for generating new/niche universally unique identifiers (ex. UUIDv6, ULID, KSUID) |
| 2800 | [`pg_jsonschema`](/e/pg_jsonschema) | [`pg_jsonschema`](/e/pg_jsonschema) | PostgreSQL extension providing JSON Schema validation |
| 1090 | [`pg_later`](/e/pg_later) | [`pg_later`](/e/pg_later) | Run queries now and get results later |
| 2460 | [`pg_parquet`](/e/pg_parquet) | [`pg_parquet`](/e/pg_parquet) | copy data between Postgres and Parquet |
| 1570 | [`pg_polyline`](/e/pg_polyline) | [`pg_polyline`](/e/pg_polyline) | Fast Google Encoded Polyline encoding & decoding for postgres |
| 4290 | [`pg_render`](/e/pg_render) | [`pg_render`](/e/pg_render) | Render HTML in SQL |
| 2100 | [`pg_search`](/e/pg_search) | [`pg_search`](/e/pg_search) | Full text search for PostgreSQL using BM25 |
| 7040 | [`pg_session_jwt`](/e/pg_session_jwt) | [`pg_session_jwt`](/e/pg_session_jwt) | Manage authentication sessions using JWTs |
| 4170 | [`pg_smtp_client`](/e/pg_smtp_client) | [`pg_smtp_client`](/e/pg_smtp_client) | PostgreSQL extension to send email using SMTP |
| 1860 | [`pg_summarize`](/e/pg_summarize) | [`pg_summarize`](/e/pg_summarize) | Text Summarization using LLMs. Built using pgrx |
| 1870 | [`pg_tiktoken`](/e/pg_tiktoken) | [`pg_tiktoken`](/e/pg_tiktoken) | tiktoken tokenizer for use with OpenAI models in postgres |
| 2160 | [`pg_tokenizer`](/e/pg_tokenizer) | [`pg_tokenizer`](/e/pg_tokenizer) | Tokenizers for full-text search |
| 5130 | [`pgdd`](/e/pgdd) | [`pgdd`](/e/pgdd) | Introspect pg data dictionary via standard SQL |
| 3540 | [`pglite_fusion`](/e/pglite_fusion) | [`pglite_fusion`](/e/pglite_fusion) | Embed an SQLite database in your PostgreSQL table |
| 1890 | [`pgml`](/e/pgml) | [`pgml`](/e/pgml) | Run AL/ML workloads with SQL interface |
| 7070 | [`pgsmcrypto`](/e/pgsmcrypto) | [`pgsmcrypto`](/e/pgsmcrypto) | PostgreSQL SM Algorithm Extension |
| 4510 | [`pgx_ulid`](/e/pgx_ulid) | [`pgx_ulid`](/e/pgx_ulid) | ulid type and methods |
| 3040 | [`plprql`](/e/plprql) | [`plprql`](/e/plprql) | Use PRQL in PostgreSQL - Pipelined Relational Query Language |
| 1010 | [`timescaledb_toolkit`](/e/timescaledb_toolkit) | [`timescaledb_toolkit`](/e/timescaledb_toolkit) | Library of analytical hyperfunctions, time-series pipelining, and other SQL utilities |
| 1680 | [`tzf`](/e/tzf) | [`pg_tzf`](/e/tzf) | Fast lookup timezone name by GPS coordinates |
| 1810 | [`vchord`](/e/vchord) | [`vchord`](/e/vchord) | Vector database plugin for Postgres, written in Rust |
| 2150 | [`vchord_bm25`](/e/vchord_bm25) | [`vchord_bm25`](/e/vchord_bm25) | A postgresql extension for bm25 ranking algorithm |
| 1830 | [`vectorize`](/e/vectorize) | [`pg_vectorize`](/e/vectorize) | The simplest way to do vector search on Postgres |
| 1820 | [`vectorscale`](/e/vectorscale) | [`pgvectorscale`](/e/vectorscale) | Advanced indexing for vector data with DiskANN |
| 8500 | [`wrappers`](/e/wrappers) | [`wrappers`](/e/wrappers) | Foreign data wrappers developed by Supabase |

## Data

{{< language "Data" >}} {{< badge content="10 Extensions" color="gray" >}}


Data-only extensions

| ID | Extension | Package | Description |
|:---:|:---|:---|:---|
| 2170 | [`hunspell_cs_cz`](/e/hunspell_cs_cz) | [`hunspell_cs_cz`](/e/hunspell_cs_cz) | Czech Hunspell Dictionary |
| 2171 | [`hunspell_de_de`](/e/hunspell_de_de) | [`hunspell_de_de`](/e/hunspell_de_de) | German Hunspell Dictionary |
| 2172 | [`hunspell_en_us`](/e/hunspell_en_us) | [`hunspell_en_us`](/e/hunspell_en_us) | en_US Hunspell Dictionary |
| 2173 | [`hunspell_fr`](/e/hunspell_fr) | [`hunspell_fr`](/e/hunspell_fr) | French Hunspell Dictionary |
| 2174 | [`hunspell_ne_np`](/e/hunspell_ne_np) | [`hunspell_ne_np`](/e/hunspell_ne_np) | Nepali Hunspell Dictionary |
| 2175 | [`hunspell_nl_nl`](/e/hunspell_nl_nl) | [`hunspell_nl_nl`](/e/hunspell_nl_nl) | Dutch Hunspell Dictionary |
| 2176 | [`hunspell_nn_no`](/e/hunspell_nn_no) | [`hunspell_nn_no`](/e/hunspell_nn_no) | Norwegian (norsk) Hunspell Dictionary |
| 2177 | [`hunspell_pt_pt`](/e/hunspell_pt_pt) | [`hunspell_pt_pt`](/e/hunspell_pt_pt) | Portuguese Hunspell Dictionary |
| 2178 | [`hunspell_ru_ru`](/e/hunspell_ru_ru) | [`hunspell_ru_ru`](/e/hunspell_ru_ru) | Russian Hunspell Dictionary |
| 2179 | [`hunspell_ru_ru_aot`](/e/hunspell_ru_ru_aot) | [`hunspell_ru_ru_aot`](/e/hunspell_ru_ru_aot) | Russian Hunspell Dictionary (from AOT.ru group) |

## C++

{{< language "C++" >}} {{< badge content="6 Extensions" color="gray" >}}


Extensions leveraging C++ features and libraries

| ID | Extension | Package | Description |
|:---:|:---|:---|:---|
| 2770 | [`hll`](/e/hll) | [`hll`](/e/hll) | type for storing hyperloglog data |
| 2430 | [`pg_duckdb`](/e/pg_duckdb) | [`pg_duckdb`](/e/pg_duckdb) | DuckDB Embedded in Postgres |
| 2440 | [`pg_mooncake`](/e/pg_mooncake) | [`pg_mooncake`](/e/pg_mooncake) | Columnstore Table in Postgres |
| 1510 | [`pgrouting`](/e/pgrouting) | [`pgrouting`](/e/pgrouting) | pgRouting Extension |
| 3010 | [`plv8`](/e/plv8) | [`plv8`](/e/plv8) | PL/JavaScript (v8) trusted procedural language |
| 2940 | [`rdkit`](/e/rdkit) | [`rdkit`](/e/rdkit) | Cheminformatics functionality for PostgreSQL. |

## Python

{{< language "Python" >}} {{< badge content="2 Extensions" color="gray" >}}


Extensions written in Python

| ID | Extension | Package | Description |
|:---:|:---|:---|:---|
| 3210 | [`faker`](/e/faker) | [`faker`](/e/faker) | Wrapper for the Faker Python library |
| 6810 | [`powa`](/e/powa) | [`powa`](/e/powa) | PostgreSQL Workload Analyser-core |

## Java

{{< language "Java" >}} {{< badge content="1 Extensions" color="gray" >}}


Extensions running on JVM

| ID | Extension | Package | Description |
|:---:|:---|:---|:---|
| 3090 | [`pljava`](/e/pljava) | [`pljava`](/e/pljava) | PL/Java procedural language |
