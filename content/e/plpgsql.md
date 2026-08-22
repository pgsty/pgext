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
|    **Need By**    | {{< ext "bedquilt" >}} {{< ext "biscuit" >}} {{< ext "cat_tools" >}} {{< ext "check_orapg" >}} {{< ext "currency" >}} {{< ext "data_historization" >}} {{< ext "db2fce" >}} {{< ext "dbpatch" >}} {{< ext "dbstat" >}} {{< ext "ddl_historization" >}} {{< ext "drop_role_helper" >}} {{< ext "dsef" >}} {{< ext "event_manager" >}} {{< ext "explanation" >}} {{< ext "firefly" >}} {{< ext "geekspeak" >}} {{< ext "generic_plan" >}} {{< ext "gogudb" >}} {{< ext "grants_manager" >}} {{< ext "hello-world" >}} {{< ext "hybrid_search" >}} {{< ext "index_analyzer" >}} {{< ext "istoria" >}} {{< ext "italian_codes" >}} {{< ext "job_queue" >}} {{< ext "json_query" >}} {{< ext "json_utils" >}} {{< ext "jsonb_schema" >}} {{< ext "jx_io" >}} {{< ext "keyhippo" >}} {{< ext "kilobase" >}} {{< ext "kissfft" >}} {{< ext "lab-orders" >}} {{< ext "launchql-base32" >}} {{< ext "launchql-ext-types" >}} {{< ext "launchql-extension-utils" >}} {{< ext "launchql-extension-verify" >}} {{< ext "launchql-inflection" >}} {{< ext "launchql-jwt-claims" >}} {{< ext "launchql-stamps" >}} {{< ext "launchql-totp" >}} {{< ext "livewire" >}} {{< ext "medications" >}} {{< ext "merge_ips" >}} {{< ext "meta_triggers" >}} {{< ext "migration" >}} {{< ext "monitoring_role" >}} {{< ext "mv_rewrite" >}} {{< ext "mv_stats" >}} {{< ext "myhelper" >}} {{< ext "mypg_sharding" >}} {{< ext "mysqlcompat" >}} {{< ext "newsfeeds" >}} {{< ext "nfiesta_gisdata" >}} {{< ext "nfiesta_sdesign" >}} {{< ext "nfiesta_target_data" >}} {{< ext "nonoms" >}} {{< ext "norm" >}} {{< ext "npm" >}} {{< ext "ollama" >}} {{< ext "omnidb_plpgsql_debugger" >}} {{< ext "partman_to_cstore" >}} {{< ext "pase" >}} {{< ext "pathman_sharding" >}} {{< ext "patients" >}} {{< ext "pg-audit-json" >}} {{< ext "pg2podg" >}} {{< ext "pg4ml" >}} {{< ext "pgAutomator" >}} {{< ext "pg_abris" >}} {{< ext "pg_accumulator" >}} {{< ext "pg_audit" >}} {{< ext "pg_audit_tools" >}} {{< ext "pg_biscuit" >}} {{< ext "pg_bleve" >}} {{< ext "pg_bm25" >}} {{< ext "pg_cache_tree" >}} {{< ext "pg_calcpi" >}} {{< ext "pg_catalog_get_defs" >}} {{< ext "pg_column_tetris" >}} {{< ext "pg_credereum" >}} {{< ext "pg_datatype_password" >}} {{< ext "pg_dbo_timestamp" >}} {{< ext "pg_dbwa" >}} {{< ext "pg_dms" >}} {{< ext "pg_drop_events" >}} {{< ext "pg_dropbuffers" >}} {{< ext "pg_dropcache" >}} {{< ext "pg_eyes" >}} {{< ext "pg_fairmlq" >}} {{< ext "pg_fsql" >}} {{< ext "pg_gen_uid" >}} {{< ext "pg_git" >}} {{< ext "pg_graphql_server" >}} {{< ext "pg_gsl" >}} {{< ext "pg_idm" >}} {{< ext "pg_idx_advisor" >}} {{< ext "pg_lake_iceberg" >}} {{< ext "pg_landmetrics" >}} {{< ext "pg_ledger" >}} {{< ext "pg_linegazer" >}} {{< ext "pg_llm_helper" >}} {{< ext "pg_lock_pool" >}} {{< ext "pg_message_queue" >}} {{< ext "pg_monitoring" >}} {{< ext "pg_normalize_email" >}} {{< ext "pg_once" >}} {{< ext "pg_os" >}} {{< ext "pg_osgr" >}} {{< ext "pg_pageprep" >}} {{< ext "pg_part" >}} {{< ext "pg_particulous" >}} {{< ext "pg_partman" >}} {{< ext "pg_pathman" >}} {{< ext "pg_paxos" >}} {{< ext "pg_popyramids_datamarts" >}} {{< ext "pg_profile" >}} {{< ext "pg_prometheus" >}} {{< ext "pg_prttn_tools" >}} {{< ext "pg_reversi" >}} {{< ext "pg_sakila_db" >}} {{< ext "pg_semantic_cache" >}} {{< ext "pg_sendmail" >}} {{< ext "pg_sentence_transformer" >}} {{< ext "pg_sessions" >}} {{< ext "pg_shardman" >}} {{< ext "pg_statviz" >}} {{< ext "pg_tileless" >}} {{< ext "pg_tms" >}} {{< ext "pg_turboquant" >}} {{< ext "pg_twkb" >}} {{< ext "pg_upless" >}} {{< ext "pg_zlog" >}} {{< ext "pgaut" >}} {{< ext "pgcat" >}} {{< ext "pgeyes" >}} {{< ext "pgfsm" >}} {{< ext "pgh_consistency" >}} {{< ext "pgh_hgm" >}} {{< ext "pgh_output" >}} {{< ext "pgh_output_en_au" >}} {{< ext "pgh_output_pt_br" >}} {{< ext "pgh_raster" >}} {{< ext "pghydro" >}} {{< ext "pgmemento" >}} {{< ext "pgmock" >}} {{< ext "pgnats" >}} {{< ext "pgparts" >}} {{< ext "pgpm-base32" >}} {{< ext "pgpm-defaults" >}} {{< ext "pgpm-faker" >}} {{< ext "pgpm-inflection" >}} {{< ext "pgpm-jwt-claims" >}} {{< ext "pgpm-measurements" >}} {{< ext "pgpm-types" >}} {{< ext "pgpm-verify" >}} {{< ext "pgrao" >}} {{< ext "pgrollup" >}} {{< ext "pgrouting" >}} {{< ext "pgsqlmock" >}} {{< ext "pgtap" >}} {{< ext "pgtap_fixture" >}} {{< ext "pgtelemetry" >}} {{< ext "pgvroom" >}} {{< ext "plparrot" >}} {{< ext "plpgsql_check" >}} {{< ext "plpgsql_wrap" >}} {{< ext "plrust" >}} {{< ext "postgres_ci" >}} {{< ext "postpic" >}} {{< ext "powa" >}} {{< ext "prescriptions" >}} {{< ext "qgres" >}} {{< ext "quria" >}} {{< ext "range_partitioning" >}} {{< ext "recall" >}} {{< ext "recursively_delete" >}} {{< ext "rep_fdw" >}} {{< ext "rls_helpers" >}} {{< ext "roleman" >}} {{< ext "rpg" >}} {{< ext "rtiles" >}} {{< ext "scheduling" >}} {{< ext "session_variables" >}} {{< ext "short_ids" >}} {{< ext "skitch-extension-defaults" >}} {{< ext "skitch-extension-jobs" >}} {{< ext "skitch-extension-utils" >}} {{< ext "skitch-extension-verify" >}} {{< ext "sphinxlink" >}} {{< ext "sql_saga" >}} {{< ext "supa_queue" >}} {{< ext "supabase" >}} {{< ext "supabase_auth_apikey" >}} {{< ext "sys_syn_dblink" >}} {{< ext "tab_tier" >}} {{< ext "table_log_pl" >}} {{< ext "table_version" >}} {{< ext "tablelog" >}} {{< ext "telephone" >}} {{< ext "test_factory" >}} {{< ext "time_for_keys" >}} {{< ext "timestampandtz" >}} {{< ext "town" >}} {{< ext "types" >}} {{< ext "unit" >}} {{< ext "units" >}} {{< ext "us-states" >}} {{< ext "uuidv7-sql" >}} {{< ext "variant" >}} {{< ext "vectors" >}} {{< ext "vrprouting" >}} {{< ext "wasm" >}} {{< ext "webauthn" >}} {{< ext "xl_global_views" >}} {{< ext "zombodb" >}} |
|   **See Also**    | {{< ext "plx" >}} {{< ext "plisql" >}} {{< ext "plpgsql_check" >}} {{< ext "pldbgapi" >}} {{< ext "plpgsql_wrap" >}} {{< ext "orafce" >}} {{< ext "plprofiler" >}} {{< ext "plsh" >}} {{< ext "pljava" >}} {{< ext "pglinter" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

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
