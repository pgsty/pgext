---
title: 下载
description: 下载 PostgreSQL 扩展
icon: Download
---



在 Pigsty 中，下载和[**安装**](/zh/usage/install)扩展是两个独立的步骤。在[**`INFRA`**](/zh/docs/infra/)模块安装期间，Pigsty 会将所有所需软件包下载到本地，并为整个部署环境创建本地 YUM/APT 仓库。

这种方式能够加速安装过程，避免重复下载，无需数据库节点访问互联网，减少网络流量，提高交付可靠性，并确保环境中版本一致——这些都是生产环境的最佳实践。

对于开发环境，也可以直接从互联网仓库安装扩展。

------

## 快速开始

在 Pigsty 安装过程中，[**`repo_packages`**](/zh/docs/infra/param/#repo_packages) 和 [**`repo_extra_packages`**](/zh/docs/infra/param/#repo_extra_packages) 中定义的软件包会自动下载到本地仓库。

与 PostgreSQL 相关的软件包（核心及扩展）通常放在 [**`repo_extra_packages`**](/zh/docs/infra/param/#repo_extra_packages) 中，而 `repo_packages` 保持其操作系统相关的全局默认值。

`repo_extra_packages` 的默认值为 `[pgsql-main]`，这是一个别名，代表当前主用 PG 版本的核心及关键扩展。

```yaml
repo_extra_packages: [ pgsql-main ]  # 当前 PG 主版本 17 的核心包（内核+3个扩展）
```

如需添加特定扩展，只需将 Pigsty 的[**扩展包名（`pkg`）**](/zh/usage/pkg)加入该参数。Pigsty 会自动为当前 PG 版本和操作系统下载合适的软件包。

```yaml
repo_extra_packages: [ pgsql-main, documentdb, citus, postgis, pgvector, pg_cron, rum ]
```

如需为当前 PG 版本下载**所有可用**扩展，可添加全部 **16** 个扩展类别别名（如[`rich`](/zh/docs/config/template)配置模板所示）：

```yaml
repo_extra_packages: [ pgsql-main ,pgsql-time ,pgsql-gis ,pgsql-rag ,pgsql-fts ,pgsql-olap ,pgsql-feat ,pgsql-lang ,pgsql-type ,pgsql-util ,pgsql-func ,pgsql-admin ,pgsql-stat ,pgsql-sec ,pgsql-fdw ,pgsql-sim ,pgsql-etl]
```

也可使用版本专属别名，一次性下载多个 PostgreSQL 版本的扩展：

```yaml
repo_extra_packages: [
    pg17-core,pg17-time,pg17-gis,pg17-rag,pg17-fts,pg17-olap,pg17-feat,pg17-lang,pg17-type,pg17-util,pg17-func,pg17-admin,pg17-stat,pg17-sec,pg17-fdw,pg17-sim,pg17-etl,
    pg16-core,pg16-time,pg16-gis,pg16-rag,pg16-fts,pg16-olap,pg16-feat,pg16-lang,pg16-type,pg16-util,pg16-func,pg16-admin,pg16-stat,pg16-sec,pg16-fdw,pg16-sim,pg16-etl,
    pg15-core,pg15-time,pg15-gis,pg15-rag,pg15-fts,pg15-olap,pg15-feat,pg15-lang,pg15-type,pg15-util,pg15-func,pg15-admin,pg15-stat,pg15-sec,pg15-fdw,pg15-sim,pg15-etl,
    pg14-core,pg14-time,pg14-gis,pg14-rag,pg14-fts,pg14-olap,pg14-feat,pg14-lang,pg14-type,pg14-util,pg14-func,pg14-admin,pg14-stat,pg14-sec,pg14-fdw,pg14-sim,pg14-etl,
    pg13-core,pg13-time,pg13-gis,pg13-rag,pg13-fts,pg13-olap,pg13-feat,pg13-lang,pg13-type,pg13-util,pg13-func,pg13-admin,pg13-stat,pg13-sec,pg13-fdw,pg13-sim,pg13-etl,
]
```

如需将新扩展添加到本地仓库，修改上述参数后执行：

```bash
./infra.yml -t repo_build   # 重新下载并重建本地仓库
```

如需在环境中其他节点刷新仓库元数据，执行：

```bash
./node.yml  -t node_repo    # 【可选】apt update / yum makecache
```

------

## 别名映射

PostgreSQL 拥有丰富的开源生态，涵盖多种系统与架构的软件包。

Pigsty 提供了抽象层，将 PostgreSQL 软件包按“别名”分类，屏蔽了系统、架构和 PG 版本的差异。

在[**快速开始**](/zh/usage/download#quick-start)中，我们用到了如 `pgsql-main`、`pgsql-core` 等别名。这些别名会根据系统和架构自动映射为具体包名。以 EL 系统为例，`pgsql-main` 会展开为 `postgresql$v*` 内核包及 `pgvector_$v*`、`pg_repack_$v*`、`wal2json_$v*` 等扩展包。

```yaml
pgsql-main:   "postgresql$v* pg_repack_$v* wal2json_$v* pgvector_$v*"
```

其中 `$v` 占位符会被**`pg_version`**（默认：17）替换为目标版本号，`*` 通配符会展开为所有包变体（如 server、libs、contrib、devel）。Pigsty 会自动处理这些细节。

完整的可用包与别名列表见 [`roles/node_id/vars/.yml`](https://github.com/pgsty/pigsty/tree/main/roles/node_id/vars)。以下是所有支持系统通用的常用别名：

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

使用这些别名时，`$v` 会被 [`pg_version`](/zh/docs/pgsql/param#pg_version)（默认：17）替换为 PostgreSQL 主版本号。

如需下载不同 PostgreSQL 版本的软件包，可：

- 修改 [**`pg_version`**](/zh/docs/pgsql/param#pg_version) 参数，或
- 使用版本专属别名，将 `pgsql-` 前缀替换为 `pg17-`、`pg16-`、`pg15-` 等。

并非所有扩展在所有系统上都可用。部分扩展在别名中被注释，原因包括：

- 某些系统不可用
- 依赖复杂（如 [**`pl/R`**](/zh/e/plr)）
- 依赖商业软件（如 [**`oracle_fdw`**](/zh/e/oracle_fdw)）
- 在最新 PG 17 不可用，但早期版本可用

如有需要，仍可手动添加这些扩展。
