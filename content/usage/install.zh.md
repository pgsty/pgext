---
title: 安装
description: 安装 PostgreSQL 扩展
icon: PackagePlus
---

Pigsty 通过标准操作系统包管理器（yum/apt）来安装 PostgreSQL 扩展。

------

## 快速开始

安装扩展时，Pigsty 使用与[**别名映射**](/zh/usage/pkg)和[**下载**](/zh/usage/download)章节相同的机制。

为集群 `pg-meta` 安装在 `pg_extensions` 参数中显式指定的所有扩展：

```yaml
all:
  children:
    pg-meta:
      hosts: { 10.10.10.10: { pg_seq: 1, pg_role: primary } }
      vars:
        pg_cluster: pg-meta
        pg_extensions: # 需在该集群安装的扩展
          - timescaledb timescaledb_toolkit pg_timeseries periods temporal_tables emaj table_version pg_cron pg_task pg_later pg_background
          - postgis pgrouting pointcloud pg_h3 q3c ogr_fdw geoip pg_polyline pg_geohash #mobilitydb
          - pgvector vchord pgvectorscale pg_vectorize pg_similarity smlar pg_summarize pg_tiktoken pg4ml #pgml
          - pg_search pgroonga pg_bigm zhparser pg_bestmatch vchord_bm25 hunspell
          - citus hydra pg_analytics pg_duckdb pg_mooncake duckdb_fdw pg_parquet pg_fkpart pg_partman plproxy #pg_strom
          - age hll rum pg_graphql pg_jsonschema jsquery pg_hint_plan hypopg index_advisor pg_plan_filter imgsmlr pg_ivm pg_incremental pgmq pgq pg_cardano omnigres #rdkit
          - pg_tle plv8 pllua plprql pldebugger plpgsql_check plprofiler plsh pljava #plr #pgtap #faker #dbt2
          - pg_prefix pg_semver pgunit pgpdf pglite_fusion md5hash asn1oid roaringbitmap pgfaceting pgsphere pg_country pg_xenophile pg_currency pg_collection pgmp numeral pg_rational pguint pg_uint128 hashtypes ip4r pg_uri pgemailaddr pg_acl timestamp9 chkpass #pg_duration #debversion #pg_rrule
          - pg_gzip pg_bzip pg_zstd pg_http pg_net pg_curl pgjq pgjwt pg_smtp_client pg_html5_email_address url_encode pgsql_tweaks pg_extra_time pgpcre icu_ext pgqr pg_protobuf envvar floatfile pg_readme ddl_historization data_historization pg_schedoc pg_hashlib pg_xxhash shacrypt cryptint pg_ecdsa pgsparql
          - pg_idkit pg_uuidv7 permuteseq pg_hashids sequential_uuids topn quantile lower_quantile count_distinct omnisketch ddsketch vasco pgxicor tdigest first_last_agg extra_window_functions floatvec aggs_for_vecs aggs_for_arrays pg_arraymath pg_math pg_random pg_base36 pg_base62 pg_base58 pg_financial
          - pg_repack pg_squeeze pg_dirtyread pgfincore pg_cooldown pg_ddlx pg_prioritize pg_checksums pg_readonly pg_upless pg_permissions pgautofailover pg_catcheck preprepare pgcozy pg_orphaned pg_crash pg_cheat_funcs pg_fio pg_savior safeupdate pg_drop_events table_log #pgagent #pgpool
          - pg_profile pg_tracing pg_show_plans pg_stat_kcache pg_stat_monitor pg_qualstats pg_store_plans pg_track_settings pg_wait_sampling system_stats pg_meta pgnodemx pg_sqlog bgw_replstatus pgmeminfo toastinfo pg_explain_ui pg_relusage pagevis powa
          - passwordcheck supautils pgsodium pg_vault pg_session_jwt pg_anon pg_tde pgsmcrypto pgaudit pgauditlogtofile pg_auth_mon credcheck pgcryptokey pg_jobmon logerrors login_hook set_user pg_snakeoil pgextwlist pg_auditor sslutils pg_noset
          - wrappers multicorn odbc_fdw jdbc_fdw mysql_fdw tds_fdw sqlite_fdw pgbouncer_fdw mongo_fdw redis_fdw pg_redis_pubsub kafka_fdw hdfs_fdw firebird_fdw aws_s3 log_fdw #oracle_fdw #db2_fdw
          - documentdb orafce pgtt session_variable pg_statement_rollback pg_dbms_metadata pg_dbms_lock pgmemcache #pg_dbms_job #wiltondb
          - pglogical pglogical_ticker pgl_ddl_deploy pg_failover_slots db_migrator wal2json wal2mongo decoderbufs decoder_raw mimeo pg_fact_loader pg_bulkload #repmgr
```

或全局通过类别别名安装所有扩展：

```yaml
all:
  vars:
    pg_version: 17   # 默认 PostgreSQL 版本 17，pgsql-main 等价于 pg17-main
    pg_extensions: [ pgsql-main ,pgsql-time ,pgsql-gis ,pgsql-rag ,pgsql-fts ,pgsql-olap ,pgsql-feat ,pgsql-lang ,pgsql-type ,pgsql-util ,pgsql-func ,pgsql-admin ,pgsql-stat ,pgsql-sec ,pgsql-fdw ,pgsql-sim ,pgsql-etl]
```

你也可以在别名中显式指定 PG 主版本：

```yaml
all:
  vars:
    pg_extensions: [pg17-time ,pg17-gis ,pg17-rag ,pg17-fts ,pg17-feat ,pg17-lang ,pg17-type ,pg17-util ,pg17-func ,pg17-admin ,pg17-stat ,pg17-sec ,pg17-fdw ,pg17-sim ,pg17-etl ] #,pg17-olap]
```

理论上可以一次性安装所有扩展（[`olap`](/zh/cate/olap/) 类别有两个冲突除外），但并不推荐。建议仅安装你实际需要的扩展，并在 `pg_extensions` 参数中显式指定。

------

## 配置

在 PGSQL 集群初始化期间，Pigsty 会自动安装 [**`pg_packages`**](/zh/docs/pgsql/param#pg_packages) 和 [**`pg_extensions`**](/zh/docs/pgsql/param#pg_extensions) 中指定的软件包及扩展。

- [**`pg_packages`**](/zh/docs/pgsql/param#pg_packages)
- [**`pg_extensions`**](/zh/docs/pgsql/param#pg_extensions)

这两个参数均可用于安装 PostgreSQL 相关软件包。通常，`pg_packages` 用于全局指定所有集群都应安装的软件包，如 PostgreSQL 内核、高可用组件 Patroni、连接池 pgBouncer、监控工具 pgExporter 等。

默认情况下，Pigsty 还会在此处指定 3 个重要扩展：[`pgvector`](/zh/e/vector)、[`pg_repack`](/zh/e/pg_repack)、[`wal2json`](/zh/e/wal2json)，分别用于向量检索、膨胀治理和 CDC 变更捕获。

而 `pg_extensions` 通常用于为特定集群指定扩展。默认值为空列表，表示默认不安装其他扩展。

```yaml
pg_packages:                      # 需安装的 PG 相关软件包，可用别名，state=present
  - postgresql
  - wal2json pg_repack pgvector
  - patroni pgbouncer pgbackrest pg_exporter pgbadger vip-manager
pg_extensions: []                 # 需安装的 PG 扩展，可用别名，state=latest
```

需要注意的是：通过 `pg_packages` 安装的软件包仅保证已安装，而 `pg_extensions` 安装的扩展会自动升级到**最新**可用版本。

如果使用本地软件仓库，这一区别无影响。但若使用互联网源，请谨慎考虑，将不希望自动升级的扩展移至 `pg_packages`。

------

## 安装

在集群初始化时，`pg_extensions`（及 `pg_packages`）中预定义的扩展会自动安装。

如需在已部署的 PostgreSQL 集群上安装新扩展：

首先将扩展添加到 [**`pg_extensions`**](/zh/docs/pgsql/param#pg_extensions)，然后执行如下子任务：

```bash
./pgsql.yml -t pg_extension  # 安装 pg_extensions 中指定的扩展
```

注意，通过 `pg_extension` 任务安装的扩展会**自动升级**到当前环境下的最新版本。

------

## 软件源

安装扩展需满足以下任一条件：

- **本地仓库**：已配置 Pigsty 本地软件仓库，且扩展已[**下载**](/zh/usage/download/)至本地仓库。
- **在线仓库**：目标节点已配置互联网源，且可访问外网。

生产环境推荐使用 Pigsty 本地软件仓库统一管理和安装扩展：先将扩展[**下载**](/zh/usage/download)到本地仓库，再从本地安装。
这样可确保环境中扩展版本一致，且数据库节点无需直接访问互联网。只需确保扩展已下载到本地仓库，无需额外操作。

开发环境可直接使用互联网源以便捷安装。可用如下命令为目标集群启用互联网源并直接安装扩展：

```bash
./node.yml  -l <cls> -t node_repo -e node_repo_modules=local,node,pgsql    # 目标节点启用互联网源
./pgsql.yml -l <cls> -t pg_extension                                        # 使用本地+互联网源安装扩展
```

------

## 包别名

安装扩展时，用户可通过[**扩展别名**](/zh/usage/pkg)指定扩展。

别名会根据当前激活的 PG 主版本和操作系统环境[**自动转换**](/zh/usage/pkg)。

并通过[**别名转换机制**](/zh/usage/pkg)映射为对应的 RPM/DEB 包名。

------

## 注意事项

- 存在两个已知冲突：
  - [`citus`](/zh/e/citus/) 与 [`hydra`](/zh/e/columnar) 互斥，hydra 是 citus columnar 的分支但未重命名
  - [`pg_duckdb`](/zh/e/pg_duckdb)、[`pg_mooncake`](/zh/e/pg_mooncake)、[`duckdb_fdw`](/zh/e/duckdb_fdw) 只能三选一，因各自依赖独立 libduckdb
- `pgaudit` 在 el 系列 pg15- 版本包名不同：pg16+ = pgaudit，pg15=pgaudit17，pg14=pgaudit16，pg13=pgaudit15，pg12=pgaudit14
- `postgis` 在 el 包名中自带版本号：默认 postgis35，el7 为 postgis33