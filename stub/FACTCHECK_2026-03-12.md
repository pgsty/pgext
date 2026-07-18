# Stub Fact Check Report (2026-03-12)

## Scope

- Target: `stub/*.md` (454 extension stubs)
- Method: 5-way parallel verification against upstream README/homepage + heuristic SQL/API overlap checks
- Tracking file updated: `task.md`

## Final Marking

- `checked` (verified pass): 273
- `reviewed` (needs follow-up): 151
- untouched rows (no matching stub/task): 11

## Automated Check Snapshot

- PASS: 292
- FAIL: 162
- Failure reasons: usecase_coverage_low=150, example_mismatch=52, upstream_unreachable=4, missing_usecase_section=1
- Manual corrections applied after run: `bool_plperl` source link fixed, `old_snapshot` source moved to PG16 docs, `pg_explain_ui` source links added, and `file_fdw`/`old_snapshot` revalidated as pass.

## High-Risk Follow-Up List

Extensions still flagged with `example_mismatch` or `upstream_unreachable` (excluding manually revalidated entries):

- `biscuit`: upstream_unreachable ()
- `documentdb`: example_mismatch,usecase_coverage_low (https://github.com/documentdb/documentdb)
- `documentdb_distributed`: example_mismatch,usecase_coverage_low (https://github.com/documentdb/documentdb)
- `documentdb_extended_rum`: example_mismatch,usecase_coverage_low (https://github.com/documentdb/documentdb)
- `duckdb_fdw`: example_mismatch,usecase_coverage_low (https://github.com/alitrack/duckdb_fdw)
- `firebird_fdw`: example_mismatch,usecase_coverage_low (https://github.com/ibarwick/firebird_fdw)
- `gzip`: example_mismatch,usecase_coverage_low (https://github.com/pramsey/pgsql-gzip)
- `hdfs_fdw`: example_mismatch,usecase_coverage_low (https://github.com/EnterpriseDB/hdfs_fdw)
- `hstore_pllua`: example_mismatch,usecase_coverage_low (https://github.com/pllua/pllua)
- `hstore_plluau`: example_mismatch,usecase_coverage_low (https://github.com/pllua/pllua)
- `hstore_plperl`: example_mismatch (https://www.postgresql.org/docs/current/hstore.html)
- `hstore_plperlu`: example_mismatch (https://www.postgresql.org/docs/current/hstore.html)
- `hstore_plpython3u`: example_mismatch (https://www.postgresql.org/docs/current/hstore.html)
- `intarray`: example_mismatch (https://www.postgresql.org/docs/current/intarray.html)
- `ip4r`: example_mismatch,usecase_coverage_low (https://github.com/RhodiumToad/ip4r)
- `jsquery`: example_mismatch,usecase_coverage_low (https://github.com/postgrespro/jsquery)
- `l10n_table_dependent_extension`: example_mismatch,usecase_coverage_low (https://github.com/bigsmoke/pg_xenophile)
- `md5hash`: example_mismatch,usecase_coverage_low (https://github.com/tvondra/md5hash)
- `omni`: example_mismatch,usecase_coverage_low (https://github.com/omnigres/omnigres)
- `omni_http`: example_mismatch (https://docs.omnigres.org/omni_httpc/reference/)
- `omni_id`: example_mismatch (https://docs.omnigres.org/omni_id/identity_type/)
- `omni_sql`: example_mismatch,usecase_coverage_low (https://github.com/omnigres/omnigres)
- `omni_txn`: example_mismatch,usecase_coverage_low (https://docs.omnigres.org/omni_txn/linearize/)
- `omni_vfs_types_v1`: example_mismatch,usecase_coverage_low (https://github.com/omnigres/omnigres)
- `pagevis`: example_mismatch,usecase_coverage_low (https://github.com/hollobon/pagevis)
- `pg4ml`: example_mismatch,usecase_coverage_low (https://gitee.com/guotiecheng/plpgsql_pg4ml)
- `pg_analytics`: example_mismatch,usecase_coverage_low (https://github.com/paradedb/pg_analytics)
- `pg_dbms_job`: example_mismatch,usecase_coverage_low (https://github.com/MigOpsRepos/pg_dbms_job)
- `pg_explain_ui`: example_mismatch,usecase_coverage_low (https://github.com/davidgomes/pg-explain-ui)
- `pg_extra_time`: example_mismatch,usecase_coverage_low (https://github.com/bigsmoke/pg_extra_time)
- `pg_savior`: example_mismatch,usecase_coverage_low (https://github.com/viggy28/pg_savior)
- `pg_search`: example_mismatch,usecase_coverage_low (https://github.com/paradedb/paradedb/tree/dev/pg_search)
- `pg_show_plans`: example_mismatch,usecase_coverage_low (https://github.com/cybertec-postgresql/pg_show_plans)
- `pg_sphere`: example_mismatch,usecase_coverage_low (https://github.com/postgrespro/pgsphere)
- `pg_stat_monitor`: example_mismatch,usecase_coverage_low (https://github.com/percona/pg_stat_monitor)
- `pg_stat_statements`: example_mismatch (https://www.postgresql.org/docs/current/pgstatstatements.html)
- `pg_statement_rollback`: example_mismatch,usecase_coverage_low (https://github.com/lzlabs/pg_statement_rollback)
- `pg_store_plans`: example_mismatch,usecase_coverage_low (https://github.com/ossc-db/pg_store_plans)
- `pg_xenophile`: example_mismatch,usecase_coverage_low (https://github.com/bigsmoke/pg_xenophile)
- `pgauditlogtofile`: example_mismatch,usecase_coverage_low (https://github.com/fmbiete/pgauditlogtofile)
- `pgautofailover`: example_mismatch,usecase_coverage_low (https://github.com/hapostgres/pg_auto_failover)
- `pgl_ddl_deploy`: example_mismatch,usecase_coverage_low (https://github.com/enova/pgl_ddl_deploy)
- `pgsentinel`: example_mismatch,usecase_coverage_low (https://github.com/pgsentinel/pgsentinel)
- `pgtt`: example_mismatch,usecase_coverage_low (https://github.com/darold/pgtt)
- `plluau`: example_mismatch,usecase_coverage_low (https://github.com/pllua/pllua)
- `plpgsql_check`: example_mismatch,usecase_coverage_low (https://github.com/okbob/plpgsql_check)
- `random`: example_mismatch,usecase_coverage_low (https://github.com/tvondra/random)
- `rum`: example_mismatch,usecase_coverage_low (https://github.com/postgrespro/rum)
- `sslutils`: example_mismatch,usecase_coverage_low (https://github.com/EnterpriseDB/sslutils)
- `supautils`: example_mismatch,usecase_coverage_low (https://github.com/supabase/supautils)
- `unit`: example_mismatch,usecase_coverage_low (https://github.com/df7cb/postgresql-unit)
- `vectorscale`: example_mismatch,usecase_coverage_low (https://github.com/timescale/pgvectorscale)
- `wal2json`: example_mismatch,usecase_coverage_low (https://github.com/eulerto/wal2json)
- `xml2`: example_mismatch (https://www.postgresql.org/docs/current/xml2.html)

## Raw Outputs

- `tmp/factcheck/report.jsonl`
- `tmp/factcheck/summary.md`
