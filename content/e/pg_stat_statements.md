---
title: "pg_stat_statements"
linkTitle: "pg_stat_statements"
description: "track planning and execution statistics of all SQL statements executed"
weight: 6990
categories: ["STAT"]
width: full
---

[**pg_stat_statements**](https://www.postgresql.org/docs/current/pgstatstatements.html) : track planning and execution statistics of all SQL statements executed


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6990** | {{< badge content="pg_stat_statements" link="https://www.postgresql.org/docs/current/pgstatstatements.html" >}} | {{< ext "pg_stat_statements" >}} | `1.11` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "pg_stat_kcache" >}} {{< ext "powa" >}} |
|   **See Also**    | {{< ext "pg_qualstats" >}} {{< ext "pg_store_plans" >}} {{< ext "pg_track_settings" >}} {{< ext "pg_stat_monitor" >}} {{< ext "auto_explain" >}} {{< ext "pg_profile" >}} {{< ext "pg_show_plans" >}} {{< ext "pg_hint_plan" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.11" "PostgreSQL 18: version 1.11" "green" >}} | {{< bg "1.11" "PostgreSQL 17: version 1.11" "green" >}} | {{< bg "1.11" "PostgreSQL 16: version 1.11" "green" >}} | {{< bg "1.11" "PostgreSQL 15: version 1.11" "green" >}} | {{< bg "1.11" "PostgreSQL 14: version 1.11" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_stat_statements';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_stat_statements;
```



## Usage

> [pg_stat_statements: track cumulative query execution statistics](https://www.postgresql.org/docs/current/pgstatstatements.html)

pg_stat_statements tracks planning and execution statistics of all SQL statements executed by a server.

### Querying Statistics

```sql
-- Top queries by total execution time
SELECT query, calls, total_exec_time, mean_exec_time, rows
FROM pg_stat_statements
ORDER BY total_exec_time DESC
LIMIT 10;

-- Top queries by shared buffer reads (I/O intensive)
SELECT query, calls, shared_blks_read, shared_blks_hit,
       shared_blk_read_time
FROM pg_stat_statements
ORDER BY shared_blks_read DESC
LIMIT 10;

-- Extension status (deallocations, last reset)
SELECT * FROM pg_stat_statements_info;
```

### Key View Columns

| Column | Type | Description |
|--------|------|-------------|
| `queryid` | bigint | Hash identifying normalized queries |
| `query` | text | Representative query text |
| `calls` | bigint | Execution count |
| `total_exec_time` | double precision | Total execution time (ms) |
| `mean_exec_time` | double precision | Mean execution time (ms) |
| `rows` | bigint | Total rows retrieved/affected |
| `shared_blks_hit` | bigint | Shared buffer cache hits |
| `shared_blks_read` | bigint | Shared blocks read from disk |
| `shared_blk_read_time` | double precision | Time reading shared blocks (ms) |
| `wal_records` | bigint | WAL records generated |
| `wal_bytes` | numeric | Total WAL generated (bytes) |
| `plans` | bigint | Times planned |
| `total_plan_time` | double precision | Total planning time (ms) |

### Functions

```sql
-- Reset all statistics
SELECT pg_stat_statements_reset();

-- Reset for a specific query
SELECT pg_stat_statements_reset(0, 0, queryid)
FROM pg_stat_statements
WHERE query LIKE '%my_table%';

-- Reset only min/max values
SELECT pg_stat_statements_reset(0, 0, 0, true);

-- Query without text (less I/O)
SELECT * FROM pg_stat_statements(showtext := false);
```

### Configuration

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pg_stat_statements.max` | 5000 | Maximum tracked statements (server start only) |
| `pg_stat_statements.track` | `top` | `top`, `all` (nested), or `none` |
| `pg_stat_statements.track_utility` | `on` | Track utility commands |
| `pg_stat_statements.track_planning` | `off` | Track planning statistics |
| `pg_stat_statements.save` | `on` | Persist across server restarts |

Requires `shared_preload_libraries = 'pg_stat_statements'` and `compute_query_id = on`.
