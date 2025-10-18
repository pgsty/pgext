---
title: By License
description: PostgreSQL extensions organized by open source license
---

{{< license "MIT" >}} {{< license "ISC" >}} {{< license "PostgreSQL" >}} {{< license "BSD 0-Clause" >}} {{< license "BSD 2-Clause" >}} {{< license "BSD 3-Clause" >}} {{< license "Artistic" >}} {{< license "Apache-2.0" >}} {{< license "MPL-2.0" >}}

{{< license "GPL-2.0" >}} {{< license "GPL-3.0" >}} {{< license "LGPL-2.1" >}} {{< license "LGPL-3.0" >}} {{< license "AGPL-3.0" >}} {{< license "Timescale" >}}

## Summary

| License | Count | Reference | Description |
|:--------|:-----:|:-------:|:----------|
| {{< license "PostgreSQL" >}} | 200 | [License Text](https://opensource.org/licenses/postgresql) | Very liberal license based on the BSD license, allowing almost unlimited freedom. |
| {{< license "Apache-2.0" >}} | 69 | [License Text](https://opensource.org/licenses/Apache-2.0) | Permissive license with patent protection and attribution requirements. |
| {{< license "MIT" >}} | 59 | [License Text](https://opensource.org/licenses/MIT) | A permissive license that allows commercial use, modification, and private use. |
| {{< license "BSD 3-Clause" >}} | 30 | [License Text](https://opensource.org/license/bsd-3-clause) | Permissive license with attribution and endorsement restriction clauses. |
| {{< license "BSD 2-Clause" >}} | 14 | [License Text](https://opensource.org/license/bsd-2-clause) | Permissive license requiring attribution but allowing commercial use. |
| {{< license "GPL-2.0" >}} | 14 | [License Text](https://opensource.org/licenses/GPL-2.0) | Strong copyleft license requiring derivative works to be open source. |
| {{< license "GPL-3.0" >}} | 12 | [License Text](https://opensource.org/licenses/GPL-3.0) | Strong copyleft license with additional patent and hardware restrictions. |
| {{< license "AGPL-3.0" >}} | 10 | [License Text](https://opensource.org/licenses/AGPL-3.0) | Network copyleft license extending GPL to cover network-distributed software. |
| {{< license "ISC" >}} | 5 | [License Text](https://opensource.org/licenses/ISC) | A permissive license similar to MIT, allowing commercial use and modification. |
| {{< license "Artistic" >}} | 3 | [License Text](https://opensource.org/license/artistic-2-0) | Copyleft license allowing modification with certain distribution requirements. |
| {{< license "Timescale" >}} | 2 | [License Text](https://www.timescale.com/legal/licenses) | Proprietary license with restrictions on commercial use and distribution. |
| {{< license "BSD 0-Clause" >}} | 2 | [License Text](https://opensource.org/license/0bsd) | Public domain equivalent license with no restrictions on use. |
| {{< license "LGPL-3.0" >}} | 2 | [License Text](https://opensource.org/licenses/LGPL-3.0) | Weak copyleft license with additional patent and hardware provisions. |
| {{< license "MPL-2.0" >}} | 1 | [License Text](https://opensource.org/licenses/MPL-2.0) | Weak copyleft license allowing proprietary combinations with file-level copyleft. |
| {{< license "LGPL-2.1" >}} | 1 | [License Text](https://opensource.org/licenses/LGPL-2.1) | Weak copyleft license allowing proprietary applications to link dynamically. |

---------

## PostgreSQL

{{< license "PostgreSQL" >}} {{< badge content="200 Extensions" color="gray" icon="cube" >}}

[PostgreSQL License Text](https://opensource.org/licenses/postgresql) : Very liberal license based on the BSD license, allowing almost unlimited freedom.

| ID | Extension | Description |
|:---:|:---|:---|
| 1020 | {{< alias "timeseries" "pg_timeseries" >}} | Convenience API for time series stack |
| 1030 | {{< alias "periods" >}} | Provide Standard SQL functionality for PERIODs and SYSTEM VERSIONING |
| 1070 | {{< alias "pg_cron" >}} | Job scheduler for PostgreSQL |
| 1090 | {{< alias "pg_later" >}} | Run queries now and get results later |
| 1690 | {{< alias "earthdistance" >}} | calculate great-circle distances on the surface of the Earth |
| 1800 | {{< alias "vector" "pgvector" >}} | vector data type and ivfflat and hnsw access methods |
| 1820 | {{< alias "vectorscale" "pgvectorscale" >}} | Advanced indexing for vector data with DiskANN |
| 1830 | {{< alias "vectorize" "pg_vectorize" >}} | The simplest way to do vector search on Postgres |
| 1850 | {{< alias "smlar" >}} | Effective similarity search |
| 1860 | {{< alias "pg_summarize" >}} | Text Summarization using LLMs. Built using pgrx |
| 2110 | {{< alias "pgroonga" >}} | Use Groonga as index, fast full text search platform for all languages! |
| 2111 | {{< alias "pgroonga_database" "pgroonga" >}} | PGroonga database management module |
| 2120 | {{< alias "pg_bigm" >}} | create 2-gram (bigram) index for faster full text search. |
| 2130 | {{< alias "zhparser" >}} | a parser for full-text search of Chinese |
| 2170 | {{< alias "hunspell_cs_cz" >}} | Czech Hunspell Dictionary |
| 2171 | {{< alias "hunspell_de_de" >}} | German Hunspell Dictionary |
| 2172 | {{< alias "hunspell_en_us" >}} | en_US Hunspell Dictionary |
| 2173 | {{< alias "hunspell_fr" >}} | French Hunspell Dictionary |
| 2174 | {{< alias "hunspell_ne_np" >}} | Nepali Hunspell Dictionary |
| 2175 | {{< alias "hunspell_nl_nl" >}} | Dutch Hunspell Dictionary |
| 2176 | {{< alias "hunspell_nn_no" >}} | Norwegian (norsk) Hunspell Dictionary |
| 2177 | {{< alias "hunspell_pt_pt" >}} | Portuguese Hunspell Dictionary |
| 2178 | {{< alias "hunspell_ru_ru" >}} | Russian Hunspell Dictionary |
| 2179 | {{< alias "hunspell_ru_ru_aot" >}} | Russian Hunspell Dictionary (from AOT.ru group) |
| 2180 | {{< alias "fuzzystrmatch" >}} | determine similarities and distance between strings |
| 2190 | {{< alias "pg_trgm" >}} | text similarity measurement and index searching based on trigrams |
| 2420 | {{< alias "pg_analytics" >}} | Postgres for analytics, powered by DuckDB |
| 2460 | {{< alias "pg_parquet" >}} | copy data between Postgres and Parquet |
| 2510 | {{< alias "pg_partman" >}} | Extension to manage partitioned tables by time or ID |
| 2530 | {{< alias "pg_strom" >}} | PG-Strom - big-data processing acceleration using GPU and NVME |
| 2590 | {{< alias "tablefunc" >}} | functions that manipulate whole tables, including crosstab |
| 2780 | {{< alias "rum" >}} | RUM index access method |
| 2810 | {{< alias "jsquery" >}} | data type for jsonb inspection |
| 2830 | {{< alias "hypopg" >}} | Hypothetical indexes for PostgreSQL |
| 2840 | {{< alias "index_advisor" >}} | Query index advisor |
| 2850 | {{< alias "plan_filter" "pg_plan_filter" >}} | filter statements by their execution plans. |
| 2860 | {{< alias "imgsmlr" >}} | Image similarity with haar |
| 2870 | {{< alias "pg_ivm" >}} | incremental view maintenance on PostgreSQL |
| 2880 | {{< alias "pg_incremental" >}} | Incremental Processing by Crunchy Data |
| 2900 | {{< alias "pgmq" >}} | A lightweight message queue. Like AWS SQS and RSMQ but on Postgres. |
| 2920 | {{< alias "orioledb" >}} | OrioleDB, the next generation transactional engine |
| 2990 | {{< alias "bloom" >}} | bloom access method - signature file based index |
| 3010 | {{< alias "plv8" >}} | PL/JavaScript (v8) trusted procedural language |
| 3200 | {{< alias "pgtap" >}} | Unit testing for PostgreSQL |
| 3210 | {{< alias "faker" >}} | Wrapper for the Faker Python library |
| 3240 | {{< alias "pltcl" >}} | PL/Tcl procedural language |
| 3250 | {{< alias "pltclu" "pltcl" >}} | PL/TclU untrusted procedural language |
| 3260 | {{< alias "plperl" >}} | PL/Perl procedural language |
| 3261 | {{< alias "bool_plperl" "plperl" >}} | transform between bool and plperl |
| 3262 | {{< alias "hstore_plperl" "plperl" >}} | transform between hstore and plperl |
| 3263 | {{< alias "jsonb_plperl" "plperl" >}} | transform between jsonb and plperl |
| 3270 | {{< alias "plperlu" >}} | PL/PerlU untrusted procedural language |
| 3271 | {{< alias "bool_plperlu" "plperlu" >}} | transform between bool and plperlu |
| 3272 | {{< alias "jsonb_plperlu" "plperlu" >}} | transform between jsonb and plperlu |
| 3273 | {{< alias "hstore_plperlu" "plperlu" >}} | transform between hstore and plperlu |
| 3280 | {{< alias "plpgsql" >}} | PL/pgSQL procedural language |
| 3290 | {{< alias "plpython3u" >}} | PL/Python3U untrusted procedural language |
| 3291 | {{< alias "jsonb_plpython3u" "plpython3u" >}} | transform between jsonb and plpython3u |
| 3292 | {{< alias "ltree_plpython3u" "plpython3u" >}} | transform between ltree and plpython3u |
| 3293 | {{< alias "hstore_plpython3u" "plpython3u" >}} | transform between hstore and plpython3u |
| 3500 | {{< alias "prefix" "pg_prefix" >}} | Prefix Range module for PostgreSQL |
| 3510 | {{< alias "semver" "pg_semver" >}} | Semantic version data type |
| 3600 | {{< alias "country" "pg_country" >}} | Country data type, ISO 3166-1 |
| 3610 | {{< alias "pg_xenophile" >}} | More than the bare necessities for PostgreSQL i18n and l10n. |
| 3611 | {{< alias "l10n_table_dependent_extension" "pg_xenophile" >}} | PostgreSQL l10n toolbox |
| 3730 | {{< alias "uint" "pguint" >}} | unsigned integer types |
| 3740 | {{< alias "uint128" "pg_uint128" >}} | Native uint128 type |
| 3750 | {{< alias "hashtypes" >}} | sha1, md5 and other data types for PostgreSQL |
| 3820 | {{< alias "ip4r" >}} | IPv4/v6 and IPv4/v6 range index type for PostgreSQL |
| 3840 | {{< alias "uri" "pg_uri" >}} | URI Data type for PostgreSQL |
| 3850 | {{< alias "emailaddr" "pgemailaddr" >}} | Email address type for PostgreSQL |
| 3870 | {{< alias "debversion" >}} | Debian version number data type |
| 3920 | {{< alias "chkpass" >}} | data type for auto-encrypted passwords |
| 3930 | {{< alias "isn" >}} | data types for international product numbering standards |
| 3940 | {{< alias "seg" >}} | data type for representing line segments or floating-point intervals |
| 3950 | {{< alias "cube" >}} | data type for multidimensional cubes |
| 3960 | {{< alias "ltree" >}} | data type for hierarchical tree-like structures |
| 3970 | {{< alias "hstore" >}} | data type for storing sets of (key, value) pairs |
| 3980 | {{< alias "citext" >}} | data type for case-insensitive character strings |
| 3990 | {{< alias "xml2" >}} | XPath querying and XSLT |
| 4180 | {{< alias "pg_html5_email_address" >}} | PostgreSQL email validation that is consistent with the HTML5 spec |
| 4190 | {{< alias "url_encode" >}} | url_encode, url_decode functions |
| 4200 | {{< alias "pgsql_tweaks" >}} | Some functions and views for daily usage |
| 4220 | {{< alias "pg_extra_time" >}} | Some date time functions and operators that, |
| 4230 | {{< alias "pgpcre" >}} | Perl Compatible Regular Expression functions |
| 4240 | {{< alias "icu_ext" >}} | Access ICU functions |
| 4270 | {{< alias "envvar" >}} | Fetch the value of an environment variable |
| 4300 | {{< alias "pg_readme" >}} | Generate a README.md document for a database extension or schema |
| 4301 | {{< alias "pg_readme_test_extension" "pg_readme" >}} | Test generating a README.md document for extension or schema |
| 4320 | {{< alias "data_historization" >}} | PLPGSQL Script to historize data in partitionned table |
| 4400 | {{< alias "hashlib" "pg_hashlib" >}} | Stable hash functions for Postgres |
| 4440 | {{< alias "shacrypt" >}} | Implements SHA256-CRYPT and SHA512-CRYPT password encryption schemes |
| 4450 | {{< alias "cryptint" >}} | Encryption functions for int and bigint values |
| 4550 | {{< alias "permuteseq" >}} | Pseudo-randomly permute sequences with a format-preserving encryption on elements |
| 4640 | {{< alias "omnisketch" >}} | data structure for on-line agg of data into approximate sketch |
| 4650 | {{< alias "ddsketch" >}} | Provides ddsketch aggregate function |
| 4710 | {{< alias "first_last_agg" >}} | first() and last() aggregate functions |
| 4720 | {{< alias "extra_window_functions" >}} | Extra Window Functions for PostgreSQL |
| 4780 | {{< alias "random" "pg_random" >}} | random data generator |
| 4840 | {{< alias "financial" "pg_financial" >}} | Financial aggregate functions |
| 4880 | {{< alias "refint" >}} | functions for implementing referential integrity (obsolete) |
| 4881 | {{< alias "autoinc" >}} | functions for autoincrementing fields |
| 4882 | {{< alias "insert_username" >}} | functions for tracking who changed a table |
| 4883 | {{< alias "moddatetime" >}} | functions for tracking last modification time |
| 4890 | {{< alias "tsm_system_time" >}} | TABLESAMPLE method which accepts time in milliseconds as a limit |
| 4900 | {{< alias "dict_xsyn" >}} | text search dictionary template for extended synonym processing |
| 4910 | {{< alias "tsm_system_rows" >}} | TABLESAMPLE method which accepts number of rows as a limit |
| 4920 | {{< alias "tcn" >}} | Triggered change notifications |
| 4930 | {{< alias "uuid-ossp" >}} | generate universally unique identifiers (UUIDs) |
| 4940 | {{< alias "btree_gist" >}} | support for indexing common datatypes in GiST |
| 4950 | {{< alias "btree_gin" >}} | support for indexing common datatypes in GIN |
| 4960 | {{< alias "intarray" >}} | functions, operators, and index support for 1-D arrays of integers |
| 4970 | {{< alias "intagg" >}} | integer aggregator and enumerator (obsolete) |
| 4980 | {{< alias "dict_int" >}} | text search dictionary template for integers |
| 4990 | {{< alias "unaccent" >}} | text search dictionary that removes accents |
| 5010 | {{< alias "pg_repack" >}} | Reorganize tables in PostgreSQL databases with minimal locks |
| 5080 | {{< alias "ddlx" "pg_ddlx" >}} | DDL eXtractor functions |
| 5090 | {{< alias "prioritize" "pg_prioritize" >}} | get and set the priority of PostgreSQL backends |
| 5120 | {{< alias "pg_readonly" >}} | cluster database read only |
| 5150 | {{< alias "pgautofailover" >}} | pg_auto_failover |
| 5170 | {{< alias "pre_prepare" "preprepare" >}} | Pre Prepare your Statement server side |
| 5180 | {{< alias "pg_upless" >}} | Detect Useless UPDATE |
| 5190 | {{< alias "pgcozy" >}} | Pre-warming shared buffers according to previous pg_buffercache snapshots for PostgreSQL. |
| 5200 | {{< alias "pg_orphaned" >}} | Deal with orphaned files |
| 5220 | {{< alias "pg_cheat_funcs" >}} | Provides cheat (but useful) functions |
| 5830 | {{< alias "pg_drop_events" >}} | logs transaction ids of drop table, drop column, drop materialized view statements |
| 5840 | {{< alias "table_log" >}} | record table modification logs and PITR for table/row |
| 5880 | {{< alias "pgagent" >}} | A PostgreSQL job scheduler |
| 5890 | {{< alias "pg_prewarm" >}} | prewarm relation data |
| 5900 | {{< alias "pgpool_adm" "pgpool" >}} | Administrative functions for pgPool |
| 5910 | {{< alias "pgpool_recovery" "pgpool" >}} | recovery functions for pgpool-II for V4.3 |
| 5920 | {{< alias "pgpool_regclass" "pgpool" >}} | replacement for regclass |
| 5930 | {{< alias "lo" >}} | Large Object maintenance |
| 5940 | {{< alias "basic_archive" >}} | an example of an archive module |
| 5950 | {{< alias "basebackup_to_shell" >}} | adds a custom basebackup target called shell |
| 5960 | {{< alias "old_snapshot" >}} | utilities in support of old_snapshot_threshold |
| 5970 | {{< alias "adminpack" >}} | administrative functions for PostgreSQL |
| 5980 | {{< alias "amcheck" >}} | functions for verifying relation integrity |
| 5990 | {{< alias "pg_surgery" >}} | extension to perform surgery on a damaged relation |
| 6210 | {{< alias "pg_show_plans" >}} | show query plans of all currently running SQL statements |
| 6260 | {{< alias "pg_track_settings" >}} | Track settings changes |
| 6270 | {{< alias "pg_wait_sampling" >}} | sampling based statistics of wait events |
| 6280 | {{< alias "pgsentinel" >}} | active session history |
| 6290 | {{< alias "system_stats" >}} | EnterpriseDB system statistics for PostgreSQL |
| 6340 | {{< alias "bgw_replstatus" >}} | Small PostgreSQL background worker to report whether a node is a replication master or standby |
| 6360 | {{< alias "toastinfo" >}} | show details on toasted datums |
| 6370 | {{< alias "explain_ui" "pg_explain_ui" >}} | easily jump into a visual plan UI for any SQL query |
| 6380 | {{< alias "pg_relusage" >}} | Log all the queries that reference a particular column |
| 6810 | {{< alias "powa" >}} | PostgreSQL Workload Analyser-core |
| 6880 | {{< alias "pg_overexplain" >}} | Allow EXPLAIN to dump even more details |
| 6890 | {{< alias "pg_logicalinspect" >}} | Logical decoding components inspection |
| 6900 | {{< alias "pageinspect" >}} | inspect the contents of database pages at a low level |
| 6910 | {{< alias "pgrowlocks" >}} | show row-level locking information |
| 6920 | {{< alias "sslinfo" >}} | information about SSL certificates |
| 6930 | {{< alias "pg_buffercache" >}} | examine the shared buffer cache |
| 6940 | {{< alias "pg_walinspect" >}} | functions to inspect contents of PostgreSQL Write-Ahead Log |
| 6950 | {{< alias "pg_freespacemap" >}} | examine the free space map (FSM) |
| 6960 | {{< alias "pg_visibility" >}} | examine the visibility map (VM) and page-level visibility info |
| 6970 | {{< alias "pgstattuple" >}} | show tuple-level statistics |
| 6980 | {{< alias "auto_explain" >}} | Provides a means for logging execution plans of slow statements automatically |
| 6990 | {{< alias "pg_stat_statements" >}} | track planning and execution statistics of all SQL statements executed |
| 7050 | {{< alias "anon" "pg_anon" >}} | PostgreSQL Anonymizer (anon) extension |
| 7080 | {{< alias "pgaudit" >}} | provides auditing functionality |
| 7090 | {{< alias "pgauditlogtofile" >}} | pgAudit addon to redirect audit log to an independent file |
| 7120 | {{< alias "pgcryptokey" >}} | cryptographic key management |
| 7130 | {{< alias "pg_jobmon" >}} | Extension for logging and monitoring functions in PostgreSQL |
| 7160 | {{< alias "set_user" >}} | similar to SET ROLE but with added logging |
| 7170 | {{< alias "pg_snakeoil" >}} | The PostgreSQL Antivirus |
| 7180 | {{< alias "pgextwlist" >}} | PostgreSQL Extension Whitelisting |
| 7200 | {{< alias "sslutils" >}} | A Postgres extension for managing SSL certificates through SQL |
| 7960 | {{< alias "sepgsql" >}} | label-based mandatory access control (MAC) based on SELinux security policy. |
| 7970 | {{< alias "auth_delay" >}} | pause briefly before reporting authentication failure |
| 7980 | {{< alias "pgcrypto" >}} | cryptographic functions |
| 7990 | {{< alias "passwordcheck" >}} | checks user passwords and reject weak password |
| 8510 | {{< alias "multicorn" >}} | Fetch foreign data in Python in your PostgreSQL server. |
| 8520 | {{< alias "odbc_fdw" >}} | Foreign data wrapper for accessing remote databases using ODBC |
| 8530 | {{< alias "jdbc_fdw" >}} | foreign-data wrapper for remote servers available over JDBC |
| 8540 | {{< alias "pgspider_ext" >}} | foreign-data wrapper for remote PGSpider servers |
| 8610 | {{< alias "oracle_fdw" >}} | foreign data wrapper for Oracle access |
| 8620 | {{< alias "tds_fdw" >}} | Foreign data wrapper for querying a TDS database (Sybase or Microsoft SQL Server) |
| 8630 | {{< alias "db2_fdw" >}} | foreign data wrapper for DB2 access |
| 8640 | {{< alias "sqlite_fdw" >}} | SQLite Foreign Data Wrapper |
| 8650 | {{< alias "pgbouncer_fdw" >}} | Extension for querying PgBouncer stats from normal SQL views & running pgbouncer commands from normal SQL functions |
| 8710 | {{< alias "redis_fdw" >}} | Foreign data wrapper for querying a Redis server |
| 8730 | {{< alias "kafka_fdw" >}} | kafka Foreign Data Wrapper for CSV formatted messages |
| 8750 | {{< alias "firebird_fdw" >}} | Foreign data wrapper for Firebird |
| 8970 | {{< alias "dblink" >}} | connect to other PostgreSQL databases from within a database |
| 8980 | {{< alias "file_fdw" >}} | foreign-data wrapper for flat file access |
| 8990 | {{< alias "postgres_fdw" >}} | foreign-data wrapper for remote PostgreSQL servers |
| 9240 | {{< alias "pg_dbms_metadata" >}} | Extension to add Oracle DBMS_METADATA compatibility to PostgreSQL |
| 9250 | {{< alias "pg_dbms_lock" >}} | Extension to add Oracle DBMS_LOCK full compatibility to PostgreSQL |
| 9260 | {{< alias "pg_dbms_job" >}} | Extension to add Oracle DBMS_JOB full compatibility to PostgreSQL |
| 9500 | {{< alias "pglogical" >}} | PostgreSQL Logical Replication |
| 9501 | {{< alias "pglogical_origin" "pglogical" >}} | Dummy extension for compatibility when upgrading from Postgres 9.4 |
| 9510 | {{< alias "pglogical_ticker" >}} | Have an accurate view on pglogical replication delay |
| 9530 | {{< alias "pg_failover_slots" >}} | PG Failover Slots extension |
| 9660 | {{< alias "decoder_raw" >}} | Output plugin for logical replication in Raw SQL format |
| 9700 | {{< alias "mimeo" >}} | Extension for specialized, per-table replication between PostgreSQL instances |
| 9970 | {{< alias "test_decoding" >}} | SQL-based test/example module for WAL logical decoding |
| 9980 | {{< alias "pgoutput" >}} | Logical Replication output plugin |

## Apache-2.0

{{< license "Apache-2.0" >}} {{< badge content="69 Extensions" color="gray" icon="cube" >}}

[Apache-2.0 License Text](https://opensource.org/licenses/Apache-2.0) : Permissive license with patent protection and attribution requirements.

| ID | Extension | Description |
|:---:|:---|:---|
| 1530 | {{< alias "h3" "pg_h3" >}} | H3 bindings for PostgreSQL |
| 1531 | {{< alias "h3_postgis" "pg_h3" >}} | H3 PostGIS integration |
| 1870 | {{< alias "pg_tiktoken" >}} | tiktoken tokenizer for use with OpenAI models in postgres |
| 2140 | {{< alias "pg_bestmatch" >}} | Generate BM25 sparse vector inside PostgreSQL |
| 2160 | {{< alias "pg_tokenizer" >}} | Tokenizers for full-text search |
| 2760 | {{< alias "age" >}} | AGE graph database extension |
| 2770 | {{< alias "hll" >}} | type for storing hyperloglog data |
| 2790 | {{< alias "pg_graphql" >}} | Add in-database GraphQL support |
| 2800 | {{< alias "pg_jsonschema" >}} | PostgreSQL extension providing JSON Schema validation |
| 2951 | {{< alias "omni" "omnigres" >}} | Advanced adapter for Postgres extensions |
| 2952 | {{< alias "omni_auth" "omnigres" >}} | Basic session management |
| 2953 | {{< alias "omni_aws" "omnigres" >}} | Amazon Web Services APIs (S3) |
| 2954 | {{< alias "omni_cloudevents" "omnigres" >}} | CloudEvents support |
| 2955 | {{< alias "omni_containers" "omnigres" >}} | Docker container management |
| 2956 | {{< alias "omni_credentials" "omnigres" >}} | Application credential management |
| 2958 | {{< alias "omni_email" "omnigres" >}} | E-mail framework |
| 2959 | {{< alias "omni_http" "omnigres" >}} | Basic HTTP types |
| 2960 | {{< alias "omni_httpc" "omnigres" >}} | HTTP client |
| 2961 | {{< alias "omni_httpd" "omnigres" >}} | HTTP server |
| 2962 | {{< alias "omni_id" "omnigres" >}} | Identity types |
| 2963 | {{< alias "omni_json" "omnigres" >}} | JSON toolkit |
| 2964 | {{< alias "omni_kube" "omnigres" >}} | Kubernetes (k8s) integration |
| 2965 | {{< alias "omni_ledger" "omnigres" >}} | Financial ledger |
| 2966 | {{< alias "omni_manifest" "omnigres" >}} | Package installation manifests |
| 2967 | {{< alias "omni_mimetypes" "omnigres" >}} | MIME types |
| 2968 | {{< alias "omni_os" "omnigres" >}} | Operating system integration |
| 2969 | {{< alias "omni_polyfill" "omnigres" >}} | Postgres API polyfills |
| 2970 | {{< alias "omni_python" "omnigres" >}} | First-class Python support |
| 2971 | {{< alias "omni_regex" "omnigres" >}} | PCRE-compatible regular expressions |
| 2972 | {{< alias "omni_rest" "omnigres" >}} | REST API toolkit (with PostgREST support) |
| 2973 | {{< alias "omni_schema" "omnigres" >}} | Advanced schema management tooling |
| 2974 | {{< alias "omni_seq" "omnigres" >}} | Distributed integer sequences |
| 2975 | {{< alias "omni_service" "omnigres" >}} | Service management |
| 2976 | {{< alias "omni_session" "omnigres" >}} | Session management |
| 2977 | {{< alias "omni_sql" "omnigres" >}} | Programmatic SQL manipulation |
| 2979 | {{< alias "omni_sqlite" "omnigres" >}} | Embedded SQLite |
| 2980 | {{< alias "omni_test" "omnigres" >}} | Testing framework |
| 2981 | {{< alias "omni_txn" "omnigres" >}} | Transaction management |
| 2982 | {{< alias "omni_types" "omnigres" >}} | Advanced types |
| 2983 | {{< alias "omni_var" "omnigres" >}} | Scoped variables |
| 2984 | {{< alias "omni_vfs" "omnigres" >}} | Virtual File System |
| 2985 | {{< alias "omni_vfs_types_v1" "omnigres" >}} | Virtual File System types (v1) |
| 2986 | {{< alias "omni_web" "omnigres" >}} | Common web stack primitives |
| 2987 | {{< alias "omni_worker" "omnigres" >}} | Generalized worker pool |
| 2988 | {{< alias "omni_xml" "omnigres" >}} | XML toolkit |
| 2989 | {{< alias "omni_yaml" "omnigres" >}} | YAML toolkit |
| 3000 | {{< alias "pg_tle" >}} | Trusted Language Extensions for PostgreSQL |
| 3040 | {{< alias "plprql" >}} | Use PRQL in PostgreSQL - Pipelined Relational Query Language |
| 3570 | {{< alias "roaringbitmap" >}} | support for Roaring Bitmaps |
| 3630 | {{< alias "collection" "pg_collection" >}} | Memory optimized data type to be used inside of plpglsql func |
| 4080 | {{< alias "pg_net" >}} | Async HTTP Requests |
| 4470 | {{< alias "sparql" "pgsparql" >}} | Query SPARQL datasource with SQL |
| 4500 | {{< alias "pg_idkit" >}} | multi-tool for generating new/niche universally unique identifiers (ex. UUIDv6, ULID, KSUID) |
| 4700 | {{< alias "tdigest" >}} | Provides tdigest aggregate function. |
| 5070 | {{< alias "pg_cooldown" >}} | remove buffered pages for specific relations |
| 5810 | {{< alias "pg_savior" >}} | Postgres extension to save OOPS mistakes |
| 6310 | {{< alias "pgnodemx" >}} | Capture node OS metrics via SQL queries |
| 7010 | {{< alias "supautils" >}} | Extension that secures a cluster on a cloud environment |
| 7030 | {{< alias "supabase_vault" "pg_vault" >}} | Supabase Vault Extension |
| 7040 | {{< alias "pg_session_jwt" >}} | Manage authentication sessions using JWTs |
| 8500 | {{< alias "wrappers" >}} | Foreign data wrappers developed by Supabase |
| 8800 | {{< alias "aws_s3" >}} | aws_s3 postgres extension to import/export data from/to s3 |
| 8810 | {{< alias "log_fdw" >}} | foreign-data wrapper for Postgres log file access |
| 9300 | {{< alias "babelfishpg_common" >}} | SQL Server Transact SQL Datatype Support |
| 9310 | {{< alias "babelfishpg_tsql" >}} | SQL Server Transact SQL compatibility |
| 9320 | {{< alias "babelfishpg_tds" >}} | SQL Server TDS protocol extension |
| 9330 | {{< alias "babelfishpg_money" >}} | SQL Server Money Data Type |
| 9550 | {{< alias "pgactive" >}} | Active-Active Replication Extension for PostgreSQL |
| 9640 | {{< alias "wal2mongo" >}} | PostgreSQL logical decoding output plugin for MongoDB |

## MIT

{{< license "MIT" >}} {{< badge content="59 Extensions" color="gray" icon="cube" >}}

[MIT License Text](https://opensource.org/licenses/MIT) : A permissive license that allows commercial use, modification, and private use.

| ID | Extension | Description |
|:---:|:---|:---|
| 1080 | {{< alias "pg_task" >}} | execute any sql command at any specific time at background |
| 1550 | {{< alias "ogr_fdw" >}} | foreign-data wrapper for GIS data access |
| 1570 | {{< alias "pg_polyline" >}} | Fast Google Encoded Polyline encoding & decoding for postgres |
| 1590 | {{< alias "pg_geohash" >}} | Handle geohash based functionality for spatial coordinates |
| 1680 | {{< alias "tzf" "pg_tzf" >}} | Fast lookup timezone name by GPS coordinates |
| 1890 | {{< alias "pgml" >}} | Run AL/ML workloads with SQL interface |
| 2430 | {{< alias "pg_duckdb" >}} | DuckDB Embedded in Postgres |
| 2440 | {{< alias "pg_mooncake" >}} | Columnstore Table in Postgres |
| 2450 | {{< alias "duckdb_fdw" >}} | DuckDB Foreign Data Wrapper |
| 2930 | {{< alias "pg_cardano" >}} | A suite of Cardano-related tools |
| 3020 | {{< alias "pllua" >}} | Lua as a procedural language |
| 3021 | {{< alias "hstore_pllua" "pllua" >}} | Hstore transform for Lua |
| 3030 | {{< alias "plluau" "pllua" >}} | Lua as an untrusted procedural language |
| 3031 | {{< alias "hstore_plluau" "pllua" >}} | Hstore transform for untrusted Lua |
| 3060 | {{< alias "plpgsql_check" >}} | extended check for plpgsql functions |
| 3080 | {{< alias "plsh" >}} | PL/sh procedural language |
| 3540 | {{< alias "pglite_fusion" >}} | Embed an SQLite database in your PostgreSQL table |
| 3620 | {{< alias "currency" "pg_currency" >}} | Custom PostgreSQL currency type in 1Byte |
| 3720 | {{< alias "pg_rational" >}} | bigint fractions |
| 3830 | {{< alias "pg_duration" >}} | data type for representing durations |
| 3880 | {{< alias "pg_rrule" >}} | RRULE field type for PostgreSQL |
| 3890 | {{< alias "timestamp9" >}} | timestamp nanosecond resolution |
| 4010 | {{< alias "gzip" "pg_gzip" >}} | gzip and gunzip functions. |
| 4020 | {{< alias "bzip" "pg_bzip" >}} | Bzip compression and decompression |
| 4070 | {{< alias "http" "pg_http" >}} | HTTP client for PostgreSQL, allows web page retrieval inside the database. |
| 4090 | {{< alias "pg_curl" >}} | Run curl actions for data transfer in URL syntax |
| 4150 | {{< alias "pgjq" >}} | Use jq in Postgres |
| 4160 | {{< alias "pgjwt" >}} | JSON Web Token API for Postgresql |
| 4170 | {{< alias "pg_smtp_client" >}} | PostgreSQL extension to send email using SMTP |
| 4260 | {{< alias "pg_protobuf" >}} | Protobuf support for PostgreSQL |
| 4280 | {{< alias "floatfile" >}} | Simple file storage for arrays of floats |
| 4290 | {{< alias "pg_render" >}} | Render HTML in SQL |
| 4510 | {{< alias "pgx_ulid" >}} | ulid type and methods |
| 4560 | {{< alias "pg_hashids" >}} | Short unique id generator for PostgreSQL, using hashids |
| 4570 | {{< alias "sequential_uuids" >}} | generator of sequential UUIDs |
| 4730 | {{< alias "floatvec" >}} | Math for vectors (arrays) of numbers |
| 4740 | {{< alias "aggs_for_vecs" >}} | Aggregate functions for array inputs |
| 4750 | {{< alias "aggs_for_arrays" >}} | Various functions for computing statistics on arrays of numbers |
| 4760 | {{< alias "arraymath" "pg_arraymath" >}} | Array math and operators that work element by element on the contents of arrays |
| 4800 | {{< alias "base36" "pg_base36" >}} | Integer Base36 types |
| 4810 | {{< alias "base62" "pg_base62" >}} | Base62 extension for PostgreSQL |
| 4830 | {{< alias "pg_base58" >}} | Base58 Encoder/Decoder Extension for PostgreSQL |
| 4850 | {{< alias "convert" "pg_convert" >}} | conversion functions for spatial, routing and other specialized uses |
| 5130 | {{< alias "pgdd" >}} | Introspect pg data dictionary via standard SQL |
| 6010 | {{< alias "pg_tracing" >}} | Distributed Tracing for PostgreSQL |
| 6350 | {{< alias "pgmeminfo" >}} | show memory usage |
| 6800 | {{< alias "pagevis" >}} | Visualise database pages in ascii code |
| 7060 | {{< alias "pg_tde" >}} | Percona pg_tde access method |
| 7070 | {{< alias "pgsmcrypto" >}} | PostgreSQL SM Algorithm Extension |
| 7100 | {{< alias "pg_auth_mon" >}} | monitor connection attempts per user |
| 7110 | {{< alias "credcheck" >}} | credcheck - postgresql plain text credential checker |
| 8720 | {{< alias "redis" "pg_redis_pubsub" >}} | Send redis pub/sub messages to Redis from PostgreSQL Directly |
| 9000 | {{< alias "documentdb" >}} | API surface for DocumentDB for PostgreSQL |
| 9010 | {{< alias "documentdb_core" "documentdb" >}} | Core API surface for DocumentDB for PostgreSQL |
| 9020 | {{< alias "documentdb_distributed" "documentdb" >}} | Multi-Node API surface for DocumentDB |
| 9410 | {{< alias "pgmemcache" >}} | memcached interface |
| 9520 | {{< alias "pgl_ddl_deploy" >}} | automated ddl deployment using pglogical |
| 9650 | {{< alias "decoderbufs" >}} | Logical decoding plugin that delivers WAL stream changes using a Protocol Buffer format |
| 9820 | {{< alias "pg_fact_loader" >}} | build fact tables with Postgres |

## BSD 3-Clause

{{< license "BSD 3-Clause" >}} {{< badge content="30 Extensions" color="gray" icon="cube" >}}

[BSD 3-Clause License Text](https://opensource.org/license/bsd-3-clause) : Permissive license with attribution and endorsement restriction clauses.

| ID | Extension | Description |
|:---:|:---|:---|
| 1060 | {{< alias "table_version" >}} | PostgreSQL table versioning extension |
| 1520 | {{< alias "pointcloud" >}} | data type for lidar point clouds |
| 1521 | {{< alias "pointcloud_postgis" "pointcloud" >}} | integration for pointcloud LIDAR data and PostGIS geometry data |
| 1840 | {{< alias "pg_similarity" >}} | support similarity queries |
| 2820 | {{< alias "pg_hint_plan" >}} | Give PostgreSQL ability to manually force some decisions in execution plans. |
| 2940 | {{< alias "rdkit" >}} | Cheminformatics functionality for PostgreSQL. |
| 3090 | {{< alias "pljava" >}} | PL/Java procedural language |
| 3580 | {{< alias "pgfaceting" >}} | fast faceting queries using an inverted index |
| 3590 | {{< alias "pg_sphere" "pgsphere" >}} | spherical objects with useful functions, operators and index support |
| 4250 | {{< alias "pgqr" >}} | QR Code generator from PostgreSQL |
| 5020 | {{< alias "pg_rewrite" >}} | Tool allows read write to the table during the rewriting |
| 5050 | {{< alias "pg_dirtyread" >}} | Read dead but unvacuumed rows from table |
| 5060 | {{< alias "pgfincore" >}} | examine and manage the os buffer cache |
| 5160 | {{< alias "pg_catcheck" >}} | Diagnosing system catalog corruption |
| 5210 | {{< alias "pg_crash" >}} | Send random signals to random processes |
| 5230 | {{< alias "fio" "pg_fio" >}} | PostgreSQL File I/O Functions |
| 6220 | {{< alias "pg_stat_kcache" >}} | Kernel statistics gathering |
| 6230 | {{< alias "pg_stat_monitor" >}} | The pg_stat_monitor is a PostgreSQL Query Performance Monitoring tool, based on PostgreSQL contrib module pg_stat_statements. pg_stat_monitor provides aggregated statistics, client information, plan details including plan, and histogram information. |
| 6240 | {{< alias "pg_qualstats" >}} | An extension collecting statistics about quals |
| 6250 | {{< alias "pg_store_plans" >}} | track plan statistics of all SQL statements executed |
| 6320 | {{< alias "pg_proctab" "pgnodemx" >}} | PostgreSQL extension to access the OS process table |
| 6330 | {{< alias "pg_sqlog" >}} | Provide SQL interface to logs |
| 7020 | {{< alias "pgsodium" >}} | Postgres extension for libsodium functions |
| 7140 | {{< alias "logerrors" >}} | Function for collecting statistics about messages in logfile |
| 7190 | {{< alias "pg_auditor" >}} | Audit data changes and provide flashback ability |
| 8600 | {{< alias "mysql_fdw" >}} | Foreign data wrapper for querying a MySQL server |
| 8740 | {{< alias "hdfs_fdw" >}} | foreign-data wrapper for remote hdfs servers |
| 9540 | {{< alias "db_migrator" >}} | Tools to migrate other databases to PostgreSQL |
| 9630 | {{< alias "wal2json" >}} | Changing data capture in JSON format |
| 9830 | {{< alias "pg_bulkload" >}} | pg_bulkload is a high speed data loading utility for PostgreSQL |

## BSD 2-Clause

{{< license "BSD 2-Clause" >}} {{< badge content="14 Extensions" color="gray" icon="cube" >}}

[BSD 2-Clause License Text](https://opensource.org/license/bsd-2-clause) : Permissive license requiring attribution but allowing commercial use.

| ID | Extension | Description |
|:---:|:---|:---|
| 1040 | {{< alias "temporal_tables" >}} | temporal tables |
| 1560 | {{< alias "geoip" >}} | IP-based geolocation query |
| 3550 | {{< alias "md5hash" >}} | type for storing 128-bit binary data inline |
| 3860 | {{< alias "acl" "pg_acl" >}} | ACL Data type |
| 4430 | {{< alias "xxhash" "pg_xxhash" >}} | xxhash functions for PostgreSQL |
| 4460 | {{< alias "pguecc" "pg_ecdsa" >}} | uECC bindings for Postgres |
| 4610 | {{< alias "quantile" >}} | Quantile aggregation function |
| 4620 | {{< alias "lower_quantile" >}} | Lower quantile aggregate function |
| 4630 | {{< alias "count_distinct" >}} | An alternative to COUNT(DISTINCT …) aggregate, usable with HashAggregate |
| 5040 | {{< alias "pg_squeeze" >}} | A tool to remove unused space from a relation. |
| 5110 | {{< alias "pg_checksums" >}} | Activate/deactivate/verify checksums in offline Postgres clusters |
| 5140 | {{< alias "pg_permissions" >}} | view object permissions and compare them with the desired state |
| 6000 | {{< alias "pg_profile" >}} | PostgreSQL load profile repository and report builder |
| 6300 | {{< alias "meta" "pg_meta" >}} | Normalized, friendlier system catalog for PostgreSQL |

## GPL-2.0

{{< license "GPL-2.0" >}} {{< badge content="14 Extensions" color="gray" icon="cube" >}}

[GPL-2.0 License Text](https://opensource.org/licenses/GPL-2.0) : Strong copyleft license requiring derivative works to be open source.

| ID | Extension | Description |
|:---:|:---|:---|
| 1500 | {{< alias "postgis" >}} | PostGIS geometry and geography spatial types and functions |
| 1501 | {{< alias "postgis_topology" "postgis" >}} | PostGIS topology spatial types and functions |
| 1502 | {{< alias "postgis_raster" "postgis" >}} | PostGIS raster types and functions |
| 1503 | {{< alias "postgis_sfcgal" "postgis" >}} | PostGIS SFCGAL functions |
| 1504 | {{< alias "postgis_tiger_geocoder" "postgis" >}} | PostGIS tiger geocoder and reverse geocoder |
| 1505 | {{< alias "address_standardizer" "postgis" >}} | Used to parse an address into constituent elements. Generally used to support geocoding address normalization step. |
| 1506 | {{< alias "address_standardizer_data_us" "postgis" >}} | Address Standardizer US dataset example |
| 1510 | {{< alias "pgrouting" >}} | pgRouting Extension |
| 1540 | {{< alias "q3c" >}} | q3c sky indexing plugin |
| 2500 | {{< alias "pg_fkpart" >}} | Table partitioning by foreign key utility |
| 3100 | {{< alias "plr" >}} | load R interpreter and execute R script from within a database |
| 3520 | {{< alias "unit" "pgunit" >}} | SI units extension |
| 3710 | {{< alias "numeral" >}} | numeral datatypes extension |
| 4310 | {{< alias "ddl_historization" >}} | Historize the ddl changes inside PostgreSQL database |

## GPL-3.0

{{< license "GPL-3.0" >}} {{< badge content="12 Extensions" color="gray" icon="cube" >}}

[GPL-3.0 License Text](https://opensource.org/licenses/GPL-3.0) : Strong copyleft license with additional patent and hardware restrictions.

| ID | Extension | Description |
|:---:|:---|:---|
| 1050 | {{< alias "emaj" >}} | Enables fine-grained write logging and time travel on subsets of the database. |
| 1100 | {{< alias "pg_background" >}} | Run SQL queries in the background |
| 1650 | {{< alias "mobilitydb" >}} | MobilityDB geospatial trajectory data management & analysis platform |
| 3530 | {{< alias "pgpdf" >}} | PDF type with meta admin & Full-Text Search |
| 3560 | {{< alias "asn1oid" >}} | asn1oid extension |
| 4330 | {{< alias "schedoc" "pg_schedoc" >}} | Cross documentation between Django and DBT projects |
| 4660 | {{< alias "vasco" >}} | discover hidden correlations in your data with MIC |
| 4670 | {{< alias "xicor" "pgxicor" >}} | XI Correlation Coefficient in Postgres |
| 4770 | {{< alias "pg_math" >}} | GSL statistical functions for postgresql |
| 7150 | {{< alias "login_hook" >}} | login_hook - hook to execute login_hook.login() at login time |
| 9120 | {{< alias "session_variable" >}} | Registration and manipulation of session variables and constants |
| 9710 | {{< alias "repmgr" >}} | Replication manager for PostgreSQL |

## AGPL-3.0

{{< license "AGPL-3.0" >}} {{< badge content="10 Extensions" color="gray" icon="cube" >}}

[AGPL-3.0 License Text](https://opensource.org/licenses/AGPL-3.0) : Network copyleft license extending GPL to cover network-distributed software.

| ID | Extension | Description |
|:---:|:---|:---|
| 1810 | {{< alias "vchord" >}} | Vector database plugin for Postgres, written in Rust |
| 1880 | {{< alias "pg4ml" >}} | Machine learning framework for PostgreSQL |
| 2100 | {{< alias "pg_search" >}} | Full text search for PostgreSQL using BM25 |
| 2150 | {{< alias "vchord_bm25" >}} | A postgresql extension for bm25 ranking algorithm |
| 2400 | {{< alias "citus" >}} | Distributed PostgreSQL as an extension |
| 2401 | {{< alias "citus_columnar" "citus" >}} | Citus columnar storage engine |
| 2410 | {{< alias "columnar" "hydra" >}} | Hydra Columnar extension |
| 4600 | {{< alias "topn" >}} | type for top-n JSONB |
| 7210 | {{< alias "noset" "pg_noset" >}} | Module for blocking SET variables for non-super users. |
| 9400 | {{< alias "spat" >}} | Redis-like In-Memory DB Embedded in Postgres |

## ISC

{{< license "ISC" >}} {{< badge content="5 Extensions" color="gray" icon="cube" >}}

[ISC License Text](https://opensource.org/licenses/ISC) : A permissive license similar to MIT, allowing commercial use and modification.

| ID | Extension | Description |
|:---:|:---|:---|
| 2910 | {{< alias "pgq" >}} | Generic queue for PostgreSQL |
| 4030 | {{< alias "zstd" "pg_zstd" >}} | Zstandard compression algorithm implementation in PostgreSQL |
| 5820 | {{< alias "safeupdate" >}} | Require criteria for UPDATE and DELETE |
| 9110 | {{< alias "pgtt" >}} | Extension to add Global Temporary Tables feature to PostgreSQL |
| 9130 | {{< alias "pg_statement_rollback" >}} | Server side rollback at statement level for PostgreSQL like Oracle or DB2 |

## Artistic

{{< license "Artistic" >}} {{< badge content="3 Extensions" color="gray" icon="cube" >}}

[Artistic License Text](https://opensource.org/license/artistic-2-0) : Copyleft license allowing modification with certain distribution requirements.

| ID | Extension | Description |
|:---:|:---|:---|
| 3050 | {{< alias "pldbgapi" "pldebugger" >}} | server-side support for debugging PL/pgSQL functions |
| 3070 | {{< alias "plprofiler" >}} | server-side support for profiling PL/pgSQL functions |
| 3220 | {{< alias "dbt2" >}} | OSDL-DBT-2 test kit |

## Timescale

{{< license "Timescale" >}} {{< badge content="2 Extensions" color="gray" icon="cube" >}}

[Timescale License Text](https://www.timescale.com/legal/licenses) : Proprietary license with restrictions on commercial use and distribution.

| ID | Extension | Description |
|:---:|:---|:---|
| 1000 | {{< alias "timescaledb" >}} | Enables scalable inserts and complex queries for time-series data |
| 1010 | {{< alias "timescaledb_toolkit" >}} | Library of analytical hyperfunctions, time-series pipelining, and other SQL utilities |

## BSD 0-Clause

{{< license "BSD 0-Clause" >}} {{< badge content="2 Extensions" color="gray" icon="cube" >}}

[BSD 0-Clause License Text](https://opensource.org/license/0bsd) : Public domain equivalent license with no restrictions on use.

| ID | Extension | Description |
|:---:|:---|:---|
| 2520 | {{< alias "plproxy" >}} | Database partitioning implemented as procedural language |
| 9100 | {{< alias "orafce" >}} | Functions and operators that emulate a subset of functions and packages from the Oracle RDBMS |

## LGPL-3.0

{{< license "LGPL-3.0" >}} {{< badge content="2 Extensions" color="gray" icon="cube" >}}

[LGPL-3.0 License Text](https://opensource.org/licenses/LGPL-3.0) : Weak copyleft license with additional patent and hardware provisions.

| ID | Extension | Description |
|:---:|:---|:---|
| 3700 | {{< alias "pgmp" >}} | Multiple Precision Arithmetic extension |
| 8700 | {{< alias "mongo_fdw" >}} | foreign data wrapper for MongoDB access |

## MPL-2.0

{{< license "MPL-2.0" >}} {{< badge content="1 Extensions" color="gray" icon="cube" >}}

[MPL-2.0 License Text](https://opensource.org/licenses/MPL-2.0) : Weak copyleft license allowing proprietary combinations with file-level copyleft.

| ID | Extension | Description |
|:---:|:---|:---|
| 4540 | {{< alias "pg_uuidv7" >}} | Create UUIDv7 values in postgres |

## LGPL-2.1

{{< license "LGPL-2.1" >}} {{< badge content="1 Extensions" color="gray" icon="cube" >}}

[LGPL-2.1 License Text](https://opensource.org/licenses/LGPL-2.1) : Weak copyleft license allowing proprietary applications to link dynamically.

| ID | Extension | Description |
|:---:|:---|:---|
| 7000 | {{< alias "passwordcheck_cracklib" >}} | Strengthen PostgreSQL user password checks with cracklib |
