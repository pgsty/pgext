---
title: "auto_explain"
linkTitle: "auto_explain"
description: "Provides a means for logging execution plans of slow statements automatically"
weight: 6980
categories: ["STAT"]
width: full
---

[**auto_explain**](https://www.postgresql.org/docs/current/auto-explain.html) : Provides a means for logging execution plans of slow statements automatically


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6980** | {{< badge content="auto_explain" link="https://www.postgresql.org/docs/current/auto-explain.html" >}} | {{< ext "auto_explain" >}} | `-` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_show_plans" >}} {{< ext "pg_store_plans" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_hint_plan" >}} {{< ext "plprofiler" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_track_settings" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "-" "PostgreSQL 18: version -" "green" >}} | {{< bg "-" "PostgreSQL 17: version -" "green" >}} | {{< bg "-" "PostgreSQL 16: version -" "green" >}} | {{< bg "-" "PostgreSQL 15: version -" "green" >}} | {{< bg "-" "PostgreSQL 14: version -" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'auto_explain';
```


This extension does not need `CREATE EXTENSION` DDL command





## Usage

> [auto_explain: automatically log slow query plans](https://www.postgresql.org/docs/current/auto-explain.html)

auto_explain automatically logs execution plans of slow statements, eliminating the need to manually run `EXPLAIN`. Plans are sent to the PostgreSQL log.

### Quick Start

```sql
-- Load per-session
LOAD 'auto_explain';
SET auto_explain.log_min_duration = '1s';
SET auto_explain.log_analyze = true;
```

Or in `postgresql.conf` for all sessions:

```
session_preload_libraries = 'auto_explain'
auto_explain.log_min_duration = '3s'
```

### Configuration Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `auto_explain.log_min_duration` | `-1` | Minimum duration to log (ms). `0` = all, `-1` = disabled |
| `auto_explain.log_analyze` | `off` | Use `EXPLAIN ANALYZE` (includes actual timing) |
| `auto_explain.log_buffers` | `off` | Include buffer usage statistics |
| `auto_explain.log_wal` | `off` | Include WAL usage statistics |
| `auto_explain.log_timing` | `on` | Include per-node timing (disable to reduce overhead) |
| `auto_explain.log_triggers` | `off` | Include trigger execution statistics |
| `auto_explain.log_verbose` | `off` | Include verbose output |
| `auto_explain.log_settings` | `off` | Log modified planner-relevant settings |
| `auto_explain.log_format` | `text` | Format: `text`, `xml`, `json`, `yaml` |
| `auto_explain.log_level` | `LOG` | Log level for output |
| `auto_explain.log_nested_statements` | `off` | Log plans for statements inside functions |
| `auto_explain.log_parameter_max_length` | `-1` | Parameter logging: `-1` = full, `0` = none |
| `auto_explain.sample_rate` | `1` | Fraction of statements to explain (0.0 to 1.0) |

### Example Log Output

```
LOG:  duration: 3.651 ms  plan:
  Query Text: SELECT count(*) FROM pg_class, pg_index
              WHERE oid = indrelid AND indisunique;
  Aggregate  (cost=16.79..16.80 rows=1 width=0)
             (actual time=3.626..3.627 rows=1 loops=1)
    ->  Hash Join  (cost=4.17..16.55 rows=92 width=0)
                   (actual time=3.349..3.594 rows=92 loops=1)
```

### Performance Tip

When using `log_analyze`, disable `log_timing` if you only need row counts:

```sql
SET auto_explain.log_analyze = true;
SET auto_explain.log_timing = off;
```
