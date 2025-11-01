---
title: Menu
description: Building All the extension provided by Pigsty
weight: 999
icon: CircleHelp
breadcrumbs: false
---

## Prepare Environment

```bash
curl https://repo.pigsty.cc/pig | bash -s 0.7.0
pig build repo
pig build tool
pig build spec
pig build rust
pig build pgrx -v 0.16.1

```

pig b get supautils-3.0.2.tar.gz


## Rust Extensions

```bash
pig build pkg timescaledb_toolkit pg_polyline pg_tzf 
pig build pkg vchord pgvectorscale pg_summarize pg_tiktoken
pig build pkg vchord_bm25 pg_tokenizer pg_bestmatch  
pig build pkg pg_graphql pg_jsonschema wrappers 
pig build pkg pg_parquet pg_render
pig build pkg plprql pg_smtp_client pg_idkit pgx_ulid
pig build pkg pg_base58 pg_convert pgdd pg_explain_ui 
pig build pkg pg_session_jwt pg_anon pgsmcrypto 
```

## Big Extensions

```bash
pig build pkg citus
pig build pkg timescaledb
pig build pkg pg_curl
pig build pkg pg_duckdb
pig build pkg plv8        # broken?
pig build pkg hydra
pig build pkg age
pig build documentdb
pig build pg_vault
```

## Common Extensions

```bash
#pig build pkg pg_timeseries
pig build pkg periods temporal_tables
pig build pkg q3c geoip pg_geohash 
pig build pkg pg_similarity smlar pg4ml
pig build pkg pg_bigm zhparser
pig build pkg hunspell_cs_cz hunspell_de_de hunspell_en_us hunspell_fr hunspell_ne_np hunspell_nl_nl hunspell_nn_no hunspell_ru_ru hunspell_ru_ru_aot hunspell_pt_pt

pig build pkg index_advisor pg_plan_filter imgsmlr pg_incremental pgmq pg_tle 
pig build pkg prefix pg_semver pgpdf md5hash asn1oid pg_roaringbitmap pgfaceting pgsphere

pig build pkg pg_country pg_xenophile pg_currency pgcollection numeral
pig build pkg pg_rational pguint pg_uint128 hashtypes pg_duration pg_uri pg_emailaddr pg_acl chkpass

pig build pkg pg_net 
pig build pkg pg_gzip pg_bzip pg_zstd pg_http pg_curl pgjq pgjwt pg_html5_email_address url_encode
pig build pkg pgpcre icu_ext pgqr pg_protobuf pg_envvar ddl_historization data_historization pg_schedoc
pig build pkg pg_hashlib pg_xxhash shacrypt cryptint pg_ecdsa pgsparql

pig build pkg permuteseq pg_hashids sequential_uuids quantile lower_quantile count_distinct omnisketch ddsketch
pig build pkg vasco pgxicor first_last_agg floatvec aggs_for_vecs aggs_for_arrays 
pig build pkg pg_arraymath pg_math pg_random pg_base36 pg_base62 pg_financial 

pig build pkg pg_cooldown ddlx preprepare pg_upless pgcozy pg_orphaned pg_crash pg_cheat_funcs pg_fio pg_savior pg_drop_events table_log
pig build pkg pg_tracing pg_meta pgnodemx pg_sqlog pgmeminfo toastinfo pg_relusage pagevis pgsentinel

pig build pkg supautils pgsodium pgcryptokey pg_snakeoil pgextwlist pg_auditor sslutils
pig build pkg sqlite_fdw redis_fdw pg_redis_pubsub kafka_fdw aws_s3 firebird_fdw session_variable #spat
pig build pkg pglogical_ticker pg_failover_slots db_migrator pgactive decoder_raw mimeo pg_bulkload
```


## Problematic

- redis_fdw
- kafka_fdw
- login_hook pg18
- supautils 18 （el10）
- pig build pkg supautils 18 breaks
- pg-meta
- floatfile
- envvar
- pgshere
- numeral
- vault
- asn1oid
- md5hash
- pglite_fusion
- pg_cardano
- imgsmlr
- noset
- vectorize
- vectorscale (missing)
- pg_summarize (pg18)
- pg_tiktoken (pg18)
- log_fdw (18 break, 14,15,16,17)
- pg_bestmatch
- pg_tiktoken
- pg_later
- pg_redis_pubsub-0.0.1.tar.gz



pig build pkg noset

```bash
pig build pkg pg_savior --pg 18     -- breaks

pig build pkg pg_cooldown --pg 14,15,16,17,18

pig build pkg floatfile
pig build[pgextwlist.spec](../rpmbuild/SPECS/pgextwlist.spec) all pg_envvar

pig build pkg emailaddr --pg 18

pig build get numeral
pig build pkg log_fdw
```

## Prepare

```bash
curl https://repo.pigsty.cc/pig | bash -s 0.7.0
pig build repo
pig build tool
pig build spec
pig build rust
pig build pgrx -v 0.16.1

pig build pgrx -v 0.12.9
cargo pgrx init --pg13=/usr/pgsql-13/bin/pg_config --pg14=/usr/pgsql-14/bin/pg_config --pg15=/usr/pgsql-15/bin/pg_config --pg16=/usr/pgsql-16/bin/pg_config --pg17=/usr/pgsql-17/bin/pg_config
pig build pkg pgvectorscale
```


## Rust Extensions

```bash
pig build pkg timescaledb_toolkit pg_polyline pg_tzf 
pig build pkg vchord pg_summarize pg_tiktoken
pig build pkg vchord_bm25 pg_tokenizer pg_bestmatch
pig build pkg pg_graphql pg_jsonschema wrappers 
pig build pkg pg_parquet pg_render
pig build pkg plprql pg_smtp_client pg_idkit pgx_ulid
pig build pkg pg_base58 pg_convert pgdd pg_explain_ui 
pig build pkg pg_session_jwt pg_anon pgsmcrypto 
```


## Common Extensions

## Deps

```bash
make scws-install
make libpgfeutils-install
pig build pkg zhparser
pig build pkg pgactive
```

## Big Extensions

```bash
pig build pkg citus
pig build pkg timescaledb
pig build pkg pg_curl
pig build pkg pg_duckdb
pig build pkg plv8        # broken?
pig build pkg hydra
pig build pkg age
pig build documentdb
pig build pg_vault
```





## Broken Extensions

- pg_timeseries
- groonga



## Exotic Kernel

```bash
pig build get orioledb-beta12.tar.gz oriolepg-17.11.tar.gz openhalodb-1.0.tar.gz
cd rpmbuild

make orioledb
make oriolepg
make orioledb
```


## DEB TODO

- pg_task 1.0.0 -> 2.1.7

## problem record

- sqlite_fdw el8 
- sslutils el8
- pg_uint128 el8
