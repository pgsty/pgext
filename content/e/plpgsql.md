---
title: "plpgsql"
linkTitle: "plpgsql"
description: "PL/pgSQL procedural language"
weight: 3280
categories: ["LANG"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["CONTRIB"]
page_width: full
---

[**plpgsql**](https://www.postgresql.org/docs/current/plpgsql.html) : PL/pgSQL procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3280** | {{< badge content="plpgsql" link="https://www.postgresql.org/docs/current/plpgsql.html" >}} | {{< ext "plpgsql" >}} | `1.0` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|    **Need By**    | `bedquilt` {{< ext "biscuit" >}} {{< ext "cat_tools" >}} `check_orapg` {{< ext "currency" >}} {{< ext "data_historization" >}} {{< ext "db2fce" >}} `dbpatch` `dbstat` {{< ext "ddl_historization" >}} `drop_role_helper` `dsef` `event_manager` `explanation` `firefly` `geekspeak` `generic_plan` `gogudb` `grants_manager` `hello-world` `hybrid_search` `index_analyzer` `istoria` `italian_codes` `job_queue` `json_query` `json_utils` `jsonb_schema` `jx_io` `keyhippo` `kilobase` `kissfft` `lab-orders` `launchql-base32` `launchql-ext-types` `launchql-extension-utils` `launchql-extension-verify` `launchql-inflection` `launchql-jwt-claims` `launchql-stamps` `launchql-totp` `livewire` `medications` `merge_ips` `meta_triggers` `migration` `monitoring_role` `mv_rewrite` `mv_stats` `myhelper` `mypg_sharding` `mysqlcompat` `newsfeeds` `nfiesta_gisdata` `nfiesta_sdesign` `nfiesta_target_data` `nonoms` `norm` `npm` `ollama` `omnidb_plpgsql_debugger` `partman_to_cstore` `pase` `pathman_sharding` `patients` `pg-audit-json` `pg2podg` {{< ext "pg4ml" >}} `pgAutomator` `pg_abris` {{< ext "pg_accumulator" >}} `pg_audit` `pg_audit_tools` `pg_biscuit` `pg_bleve` `pg_bm25` `pg_cache_tree` `pg_calcpi` `pg_catalog_get_defs` {{< ext "pg_column_tetris" >}} `pg_credereum` `pg_datatype_password` `pg_dbo_timestamp` `pg_dbwa` `pg_dms` {{< ext "pg_drop_events" >}} `pg_dropbuffers` `pg_dropcache` `pg_eyes` `pg_fairmlq` {{< ext "pg_fsql" >}} `pg_gen_uid` `pg_git` `pg_graphql_server` `pg_gsl` `pg_idm` `pg_idx_advisor` {{< ext "pg_lake_iceberg" >}} `pg_landmetrics` `pg_ledger` `pg_linegazer` `pg_llm_helper` `pg_lock_pool` `pg_message_queue` `pg_monitoring` `pg_normalize_email` `pg_once` `pg_os` `pg_osgr` `pg_pageprep` `pg_part` `pg_particulous` {{< ext "pg_partman" >}} `pg_pathman` `pg_paxos` `pg_popyramids_datamarts` {{< ext "pg_profile" >}} `pg_prometheus` `pg_prttn_tools` `pg_reversi` `pg_sakila_db` `pg_semantic_cache` `pg_sendmail` `pg_sentence_transformer` `pg_sessions` `pg_shardman` {{< ext "pg_statviz" >}} `pg_tileless` `pg_tms` `pg_turboquant` `pg_twkb` {{< ext "pg_upless" >}} `pg_zlog` `pgaut` `pgcat` `pgeyes` `pgfsm` {{< ext "pgh_consistency" >}} {{< ext "pgh_hgm" >}} {{< ext "pgh_output" >}} {{< ext "pgh_output_en_au" >}} {{< ext "pgh_output_pt_br" >}} {{< ext "pgh_raster" >}} {{< ext "pghydro" >}} {{< ext "pgmemento" >}} `pgmock` `pgnats` `pgparts` `pgpm-base32` `pgpm-defaults` `pgpm-faker` `pgpm-inflection` `pgpm-jwt-claims` `pgpm-measurements` `pgpm-types` `pgpm-verify` `pgrao` `pgrollup` {{< ext "pgrouting" >}} {{< ext "pgsqlmock" >}} {{< ext "pgtap" >}} `pgtap_fixture` {{< ext "pgtelemetry" >}} `pgvroom` `plparrot` {{< ext "plpgsql_check" >}} {{< ext "plpgsql_wrap" >}} `plrust` `postgres_ci` `postpic` {{< ext "powa" >}} `prescriptions` `qgres` `quria` `range_partitioning` `recall` `recursively_delete` `rep_fdw` `rls_helpers` `roleman` `rpg` `rtiles` `scheduling` `session_variables` `short_ids` `skitch-extension-defaults` `skitch-extension-jobs` `skitch-extension-utils` `skitch-extension-verify` `sphinxlink` `sql_saga` `supa_queue` `supabase` `supabase_auth_apikey` `sys_syn_dblink` `tab_tier` `table_log_pl` {{< ext "table_version" >}} `tablelog` `telephone` `test_factory` `time_for_keys` `timestampandtz` `town` `types` {{< ext "unit" >}} `units` `us-states` `uuidv7-sql` `variant` `vectors` `vrprouting` `wasm` `webauthn` `xl_global_views` `zombodb` |
|   **See Also**    | {{< ext "plx" >}} {{< ext "plisql" >}} {{< ext "plpgsql_check" >}} {{< ext "pldbgapi" >}} {{< ext "plpgsql_wrap" >}} {{< ext "orafce" >}} {{< ext "plprofiler" >}} {{< ext "plsh" >}} {{< ext "pljava" >}} {{< ext "pglinter" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION plpgsql;
```




## Usage

> [plpgsql: PL/pgSQL procedural language](https://www.postgresql.org/docs/current/plpgsql.html)

PL/pgSQL is PostgreSQL's default procedural language. It extends SQL with control structures, variables, cursors, and exception handling.

```sql
CREATE EXTENSION plpgsql;  -- installed by default

-- Basic function with variables and control flow
CREATE FUNCTION calculate_discount(price numeric, quantity integer) RETURNS numeric
LANGUAGE plpgsql AS $$
DECLARE
  discount numeric := 0;
BEGIN
  IF quantity >= 100 THEN
    discount := 0.20;
  ELSIF quantity >= 50 THEN
    discount := 0.10;
  ELSIF quantity >= 10 THEN
    discount := 0.05;
  END IF;
  RETURN price * quantity * (1 - discount);
END;
$$;

-- Loop and set-returning function
CREATE FUNCTION fibonacci(n integer) RETURNS SETOF integer
LANGUAGE plpgsql AS $$
DECLARE
  a integer := 0;
  b integer := 1;
  tmp integer;
BEGIN
  FOR i IN 1..n LOOP
    RETURN NEXT a;
    tmp := a + b;
    a := b;
    b := tmp;
  END LOOP;
END;
$$;

SELECT * FROM fibonacci(10);

-- Exception handling
CREATE FUNCTION safe_divide(a numeric, b numeric) RETURNS numeric
LANGUAGE plpgsql AS $$
BEGIN
  RETURN a / b;
EXCEPTION
  WHEN division_by_zero THEN
    RAISE NOTICE 'Division by zero, returning NULL';
    RETURN NULL;
END;
$$;

-- Trigger function
CREATE FUNCTION update_modified_column() RETURNS trigger
LANGUAGE plpgsql AS $$
BEGIN
  NEW.modified_at = now();
  RETURN NEW;
END;
$$;

CREATE TRIGGER set_modified
  BEFORE UPDATE ON my_table
  FOR EACH ROW EXECUTE FUNCTION update_modified_column();

-- Procedure with transaction control (PG 11+)
CREATE PROCEDURE batch_archive(batch_size integer)
LANGUAGE plpgsql AS $$
DECLARE
  rows_moved integer;
BEGIN
  LOOP
    WITH moved AS (
      DELETE FROM orders WHERE status = 'completed'
      RETURNING *
    )
    INSERT INTO orders_archive SELECT * FROM moved;

    GET DIAGNOSTICS rows_moved = ROW_COUNT;
    COMMIT;
    EXIT WHEN rows_moved < batch_size;
  END LOOP;
END;
$$;
```
