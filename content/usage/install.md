---
title: Install
description: Install PostgreSQL Extension
icon: PackagePlus
---

Pigsty piggyback on standard OS package managers (yum/apt) to install PostgreSQL extensions.

------

## Quick Start

When installing extensions, Pigsty uses the same [**alias mapping**](/usage/pkg) in the [**download**](/usage/download) section.

Install all extensions explicitly specified in the `pg_extensions` parameter, for the cluster `pg-meta`:

```yaml
all:
  children:
    pg-meta:
      hosts: { 10.10.10.10: { pg_seq: 1, pg_role: primary } }
      vars:
        pg_cluster: pg-meta
        pg_extensions: # extensions to be installed on this cluster
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

Or install all extensions by category aliases globally:

```yaml
all:
  vars:
    pg_version: 17   # default postgres version 17, so the pgsql-main is equivalent to pg17-main
    pg_extensions: [ pgsql-main ,pgsql-time ,pgsql-gis ,pgsql-rag ,pgsql-fts ,pgsql-olap ,pgsql-feat ,pgsql-lang ,pgsql-type ,pgsql-util ,pgsql-func ,pgsql-admin ,pgsql-stat ,pgsql-sec ,pgsql-fdw ,pgsql-sim ,pgsql-etl]
```

You can also specify the PG major version explicitly in these alias:

```yaml
all:
  vars:

    pg_extensions: [pg17-time ,pg17-gis ,pg17-rag ,pg17-fts ,pg17-feat ,pg17-lang ,pg17-type ,pg17-util ,pg17-func ,pg17-admin ,pg17-stat ,pg17-sec ,pg17-fdw ,pg17-sim ,pg17-etl ] #,pg17-olap]
```

Install all extensions simultaneously is applicable (except two conflicts in the [`olap`](/cate/olap/) category) but not recommended. Just install the extensions you need by explicitly specifying them in the `pg_extensions` parameter.

------

## Configure

During PGSQL cluster init, Pigsty will automatically install packages (& alias) specified in [**`pg_packages`**](/docs/pgsql/param#pg_packages) and [**`pg_extensions`**](/docs/pgsql/param#pg_extensions).

- [**`pg_packages`**](/docs/pgsql/param#pg_packages)
- [**`pg_extensions`**](/docs/pgsql/param#pg_extensions)

Both parameters can be used to install PostgreSQL-related packages. Typically, `pg_packages` is used to globally specify packages that should be installed across all PostgreSQL clusters in your environment: such as the PostgreSQL kernel, high-availability agent like Patroni, connection pooling with pgBouncer, monitoring with pgExporter, etc.

By default, Pigsty also specifies 3 important extensions here: [`pgvector`](/e/vector), [`pg_repack`](/e/pg_repack), and [`wal2json`](/e/wal2json) for vector search, bloat management, and CDC change extraction.

Meanwhile, `pg_extensions` is usually used to specify extension for a specific cluster. The default is an empty list, indicating no other extensions will be installed by default.

```yaml
pg_packages:                      # pg packages to be installed, alias can be used, state=present
  - postgresql
  - wal2json pg_repack pgvector
  - patroni pgbouncer pgbackrest pg_exporter pgbadger vip-manager
pg_extensions: []                 # pg extensions to be installed, alias can be used, state=latest
```

An important distinction: packages installed via `pg_packages` are merely ensured to be present, whereas those installed via `pg_extensions` are automatically upgraded to the **latest** available version.

When using a local software repo, this distinction isn’t an issue. However, when using upstream internet repo, consider this carefully and move extensions you don’t want automatically upgraded to `pg_packages`.

------

## Install

Extensions pre-defined in the `pg_extensions` (and `pg_packages`) will be installed during cluster provisioning.

To install new extensions on a provisioned PostgreSQL cluster:

First, add extensions to [**`pg_extensions`**](/docs/pgsql/param#pg_extensions), then execute the playbook subtask:

```bash
./pgsql.yml -t pg_extension  # install extensions specified in pg_extensions
```

Note that extension plugins specified in the `pg_extension` task will be **upgraded** to the latest available version in your current environment by default.

------

## Repo

To install extension, you need to ensure one of the following conditions is met:

- **Local Repo**: You have configured using Pigsty’s local repo, and the extensions have already been [**downloaded**](/usage/download/) to the local repo.
- **Online Repo**: You have directly configured upstream internet repo on the target node, and internet access is available on these nodes.

For production environments, we recommend using Pigsty’s local software repo to manage and install extensions uniformly: First [**download**](/usage/download) extensions to the local repo, then install them from there.
This ensures consistent extension versions across your environment and prevents database nodes from directly accessing the internet. You have to do nothing when installed from local repo, just make sure they are downloaded to the local repo.

For development environments, you may choose to directly use upstream internet repo for convenience. Use the following commands to add Internet repo and install extensions on the target cluster directly:

```bash
./node.yml  -l <cls> -t node_repo -e node_repo_modules=local,node,pgsql    # Enable internet repo on target node
./pgsql.yml -l <cls> -t pg_extension                                        # Install extensions using local+internet upstream repos
```

------

## Package Alias

When installing extensions, users can use [**extension aliases**](/usage/pkg) to specify extension.

The aliases will be [**translated**](/usage/pkg) to the current active PG major version and OS environment.

and translated to the corresponding RPM/DEB package names by [**alias translation mechanism**](/usage/pkg).

------

## Caveats

- There are two known conflicts:
  - [`citus`](/e/citus/) and [`hydra`](/e/columnar) are mutually exclusive, since hydra is a fork of citus columnar without renaming
  - Only install one from [`pg_duckdb`](/e/pg_duckdb), [`pg_mooncake`](/e/pg_mooncake), [`duckdb_fdw`](/e/duckdb_fdw), they all use their own libduckdb
- `pgaudit` got a different naming pattern on el for pg 15-: pg16+ = pgaudit, pg15=pgaudit17, pg14=pgaudit16 pg13=pgaudit15 pg12=pgaudit14
- `postgis` got its own version in el package name: postgis35 by default, and postgis33 for legacy el7