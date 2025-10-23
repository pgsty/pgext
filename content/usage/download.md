---
title: Download
description: Download PostgreSQL Extension
icon: Download
---



In Pigsty, downloading and [**installing**](/usage/install) extensions are separate steps. During [**`INFRA`**](/docs/infra/) module installation, Pigsty downloads all required software to the local machine and creates a local YUM/APT repo for the entire deployment.

This approach accelerates installation, eliminates redundant downloads, removes the need for database nodes to access the internet, reduces network traffic, improves delivery reliability, and ensures consistent versions across your environment - all best practices for production deployments.

For development environments, installing extensions directly from internet repo is also acceptable

------

## Quick Start

Packages defined in [**`repo_packages`**](/docs/infra/param/#repo_packages) and [**`repo_extra_packages`**](/docs/infra/param/#repo_extra_packages) are automatically downloaded to your local repo during Pigsty installation.

For PostgreSQL-related packages (core and extensions), typically put them in [**`repo_extra_packages`**](/docs/infra/param/#repo_extra_packages) while leaving `repo_packages` with its os-specific global defaults.

The default value for `repo_extra_packages` is `[pgsql-main]`, an alias representing core PostgreSQL and critical extensions for the current active major version.

```yaml
repo_extra_packages: [ pgsql-main ]  # main packages (kernel + 3 extension) for current pg major 17
```

To add specific extensions, simply add Pigsty [**extension package name (`pkg`)**](/usage/pkg) to this parameter. Pigsty automatically downloads the appropriate packages for your active PG version and current OS distro.

```yaml
repo_extra_packages: [ pgsql-main, documentdb, citus, postgis, pgvector, pg_cron, rum ]
```

To download **all available** extensions for the current PG version, add all **16** extension category aliases (as in the [`rich`](/docs/config/template) config template):

```yaml
repo_extra_packages: [ pgsql-main ,pgsql-time ,pgsql-gis ,pgsql-rag ,pgsql-fts ,pgsql-olap ,pgsql-feat ,pgsql-lang ,pgsql-type ,pgsql-util ,pgsql-func ,pgsql-admin ,pgsql-stat ,pgsql-sec ,pgsql-fdw ,pgsql-sim ,pgsql-etl]
```

Alternatively, use version-specific aliases to download extensions for multiple PostgreSQL versions:

```yaml
repo_extra_packages: [
    pg17-core,pg17-time,pg17-gis,pg17-rag,pg17-fts,pg17-olap,pg17-feat,pg17-lang,pg17-type,pg17-util,pg17-func,pg17-admin,pg17-stat,pg17-sec,pg17-fdw,pg17-sim,pg17-etl,
    pg16-core,pg16-time,pg16-gis,pg16-rag,pg16-fts,pg16-olap,pg16-feat,pg16-lang,pg16-type,pg16-util,pg16-func,pg16-admin,pg16-stat,pg16-sec,pg16-fdw,pg16-sim,pg16-etl,
    pg15-core,pg15-time,pg15-gis,pg15-rag,pg15-fts,pg15-olap,pg15-feat,pg15-lang,pg15-type,pg15-util,pg15-func,pg15-admin,pg15-stat,pg15-sec,pg15-fdw,pg15-sim,pg15-etl,
    pg14-core,pg14-time,pg14-gis,pg14-rag,pg14-fts,pg14-olap,pg14-feat,pg14-lang,pg14-type,pg14-util,pg14-func,pg14-admin,pg14-stat,pg14-sec,pg14-fdw,pg14-sim,pg14-etl,
    pg13-core,pg13-time,pg13-gis,pg13-rag,pg13-fts,pg13-olap,pg13-feat,pg13-lang,pg13-type,pg13-util,pg13-func,pg13-admin,pg13-stat,pg13-sec,pg13-fdw,pg13-sim,pg13-etl,
]
```

To add new extensions to your local repo, modify the parameters above and run:

```bash
./infra.yml -t repo_build   # Re-download and rebuild local repo
```

To refresh the repo metadata on all other nodes in your environment, run:

```bash
./node.yml  -t node_repo    # [Optional] apt update / yum makecache
```

------

## Alias Mapping

PostgreSQL has a rich open-source ecosystem with numerous packages across different systems and architectures.

Pigsty provides an abstraction layer that categorizes PostgreSQL packages into “aliases,” hiding differences between systems, architectures, and PG versions.

In the [**Quick Start**](/usage/download#quick-start) section, we used aliases like `pgsql-main` and `pgsql-core`. These aliases are translated into specific package names based on your system and architecture. For EL systems, `pgsql-main` expands to `postgresql$v*` kernel packages with `pgvector_$v*`, `pg_repack_$v*`, and `wal2json_$v*` extension packages.

```yaml
pgsql-main:   "postgresql$v* pg_repack_$v* wal2json_$v* pgvector_$v*"
```

The `$v` placeholder is replaced by the **`pg_version`** value (default: 17) to target the correct version. The `*` wildcard expands to include all package variants (e.g., server, libs, contrib, devel). Pigsty handles these details automatically.

The complete list of available packages and aliases is in [`roles/node_id/vars/.yml`](https://github.com/pgsty/pigsty/tree/main/roles/node_id/vars). Here are commonly used aliases available across all supported systems:

```yaml
postgresql:   "postgresql$v*"
pgsql-main:   "postgresql$v* pg_repack_$v* wal2json_$v* pgvector_$v*"
pgsql-core:   "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl postgresql$v-test postgresql$v-devel postgresql$v-llvmjit"
pgsql-simple: "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl"
pgsql-client: "postgresql$v"
pgsql-server: "postgresql$v-server postgresql$v-libs postgresql$v-contrib"
pgsql-devel:  "postgresql$v-devel"
pgsql-basic:  "pg_repack_$v* wal2json_$v* pgvector_$v*"

pgsql-time:   "timescaledb-tsl_$v* timescaledb-toolkit_$v pg_timeseries_$v periods_$v* temporal_tables_$v* e-maj_$v table_version_$v pg_cron_$v* pg_task_$v* pg_later_$v pg_background_$v*"
pgsql-gis:    "postgis35_$v* pgrouting_$v* pointcloud_$v* h3-pg_$v* q3c_$v* ogr_fdw_$v* geoip_$v pg_polyline_$v pg_geohash_$v*"
pgsql-rag:    "pgvector_$v* vchord_$v pgvectorscale_$v pg_vectorize_$v pg_similarity_$v* smlar_$v* pg_summarize_$v pg_tiktoken_$v pg4ml_$v"
pgsql-fts:    "pg_search_$v pgroonga_$v* pg_bigm_$v* zhparser_$v* pg_bestmatch_$v vchord_bm25_$v hunspell_cs_cz_$v hunspell_de_de_$v hunspell_en_us_$v hunspell_fr_$v hunspell_ne_np_$v hunspell_nl_nl_$v hunspell_nn_no_$v hunspell_ru_ru_$v hunspell_ru_ru_aot_$v"
pgsql-olap:   "citus_$v* pg_analytics_$v pg_duckdb_$v* pg_mooncake_$v* duckdb_fdw_$v* pg_parquet_$v pg_fkpart_$v pg_partman_$v* plproxy_$v*" #hydra_$v* #pg_strom_$v*
pgsql-feat:   "hll_$v* rum_$v pg_graphql_$v pg_jsonschema_$v jsquery_$v* pg_hint_plan_$v* hypopg_$v* index_advisor_$v pg_plan_filter_$v* imgsmlr_$v* pg_ivm_$v* pg_incremental_$v* pgmq_$v pgq_$v* pg_cardano_$v omnigres_$v" #apache-age_$v*
pgsql-lang:   "pg_tle_$v* plv8_$v* pllua_$v* pldebugger_$v* plpgsql_check_$v* plprofiler_$v* plsh_$v* pljava_$v*" #plprql_$v #plr_$v* #pgtap_$v* #postgresql_faker_$v* #dbt2-pgsql-extensions*
pgsql-type:   "prefix_$v* semver_$v* postgresql-unit_$v* pgpdf_$v* pglite_fusion_$v md5hash_$v* asn1oid_$v* pg_roaringbitmap_$v* pgfaceting_$v pgsphere_$v* pg_country_$v* pg_xenophile_$v pg_currency_$v* pgcollection_$v* pgmp_$v* numeral_$v* pg_rational_$v* pguint_$v* pg_uint128_$v* hashtypes_$v* ip4r_$v* pg_duration_$v* pg_uri_$v* pg_emailaddr_$v* acl_$v* timestamp9_$v* chkpass_$v*"
pgsql-util:   "pgsql_gzip_$v* pg_bzip_$v* pg_zstd_$v* pgsql_http_$v* pg_net_$v* pg_curl_$v* pgjq_$v* pgjwt_$v pg_smtp_client_$v pg_html5_email_address_$v url_encode_$v* pgsql_tweaks_$v pg_extra_time_$v pgpcre_$v icu_ext_$v* pgqr_$v* pg_protobuf_$v pg_envvar_$v* floatfile_$v* pg_readme_$v ddl_historization_$v data_historization_$v pg_schedoc_$v pg_hashlib_$v pg_xxhash_$v* postgres_shacrypt_$v* cryptint_$v* pg_ecdsa_$v* pgsparql_$v"
pgsql-func:   "pg_idkit_$v pg_uuidv7_$v* permuteseq_$v* pg_hashids_$v* sequential_uuids_$v topn_$v* quantile_$v* lower_quantile_$v* count_distinct_$v* omnisketch_$v* ddsketch_$v* vasco_$v* pgxicor_$v* tdigest_$v* first_last_agg_$v extra_window_functions_$v* floatvec_$v* aggs_for_vecs_$v* aggs_for_arrays_$v* pg_arraymath_$v* pg_math_$v* pg_random_$v* pg_base36_$v* pg_base62_$v* pg_base58_$v pg_financial_$v*"
pgsql-admin:  "pg_repack_$v* pg_squeeze_$v* pg_dirtyread_$v* pgfincore_$v* pg_cooldown_$v* ddlx_$v pg_prioritize_$v* pg_readonly_$v* pg_upless_$v pg_permissions_$v pg_catcheck_$v* preprepare_$v* pgcozy_$v pg_orphaned_$v* pg_crash_$v* pg_cheat_funcs_$v* pg_fio_$v pg_savior_$v* safeupdate_$v* pg_drop_events_$v table_log_$v" #pg_checksums_$v* #pg_auto_failover_$v* #pgagent_$v* #pgpool-II-pgsql-extensions
pgsql-stat:   "pg_profile_$v* pg_tracing_$v* pg_show_plans_$v* pg_stat_kcache_$v* pg_stat_monitor_$v* pg_qualstats_$v* pg_store_plans_$v* pg_track_settings_$v pg_wait_sampling_$v* system_stats_$v* pg_meta_$v pgnodemx_$v pg_sqlog_$v bgw_replstatus_$v* pgmeminfo_$v* toastinfo_$v* pg_explain_ui_$v pg_relusage_$v pagevis_$v powa_$v*"
pgsql-sec:    "passwordcheck_cracklib_$v* supautils_$v* pgsodium_$v* vault_$v* pg_session_jwt_$v pg_anon_$v pgsmcrypto_$v pgaudit_$v* pgauditlogtofile_$v* pg_auth_mon_$v* credcheck_$v* pgcryptokey_$v pg_jobmon_$v logerrors_$v* login_hook_$v* set_user_$v* pg_snakeoil_$v* pgextwlist_$v* pg_auditor_$v sslutils_$v* noset_$v*" #pg_tde_$v*
pgsql-fdw:    "wrappers_$v multicorn2_$v* odbc_fdw_$v* mysql_fdw_$v* tds_fdw_$v* sqlite_fdw_$v* pgbouncer_fdw_$v redis_fdw_$v* pg_redis_pubsub_$v* hdfs_fdw_$v* firebird_fdw_$v aws_s3_$v log_fdw_$v*" #jdbc_fdw_$v* #oracle_fdw_$v* #db2_fdw_$v* #mongo_fdw_$v* #kafka_fdw_$v
pgsql-sim:    "documentdb_$v* orafce_$v pgtt_$v* session_variable_$v* pg_statement_rollback_$v* pg_dbms_metadata_$v pg_dbms_lock_$v pgmemcache_$v*" #pg_dbms_job_$v #wiltondb
pgsql-etl:    "pglogical_$v* pglogical_ticker_$v* pgl_ddl_deploy_$v* pg_failover_slots_$v* db_migrator_$v wal2json_$v* postgres-decoderbufs_$v* decoder_raw_$v* mimeo_$v pg_fact_loader_$v* pg_bulkload_$v*" #wal2mongo_$v* #repmgr_$v*
```

When using these aliases, the `$v` placeholder is replaced with the PostgreSQL major version number from [`pg_version`](/docs/pgsql/param#pg_version) (default: 17).

To download packages for different PostgreSQL versions, either:

- Change the [**`pg_version`**](/docs/pgsql/param#pg_version) parameter, or
- Use version-specific aliases by replacing the `pgsql-` prefix with `pg17-`, `pg16-`, `pg15-`, etc.

Not all extensions are available on all systems. Some extensions are commented out in the aliases because they:

- Are unavailable on specific systems
- Have extensive dependencies (like [**`pl/R`**](/e/plr))
- Depend on commercial software (like [**`oracle_fdw`**](/e/oracle_fdw))
- Are unavailable in the latest PG 17 but available in earlier versions

You can still manually add these extensions if needed.
