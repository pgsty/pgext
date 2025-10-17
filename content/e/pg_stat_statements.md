---
title: "pg_stat_statements"
linkTitle: "pg_stat_statements"
description: "track planning and execution statistics of all SQL statements executed"
weight: 6990
categories: ["Stat"]
width: full
---

track planning and execution statistics of all SQL statements executed

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6990** | {{< badge content="pg_stat_statements" link="https://www.postgresql.org/docs/current/pgstatstatements.html" >}} | {{< ext "pg_stat_statements" "pg_stat_statements" >}} | `1.11` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "pg_stat_kcache" >}} {{< ext "powa" >}} |
|   **See Also**    | {{< ext "pg_qualstats" >}} {{< ext "pg_store_plans" >}} {{< ext "pg_track_settings" >}} {{< ext "pg_stat_monitor" >}} {{< ext "auto_explain" >}} {{< ext "pg_profile" >}} {{< ext "pg_show_plans" >}} {{< ext "pg_hint_plan" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="1.11" color="green" >}} | {{< badge content="1.11" color="green" >}} | {{< badge content="1.11" color="green" >}} | {{< badge content="1.11" color="green" >}} | {{< badge content="1.11" color="green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

Add this extension to [`shared_preload_libraries`](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pg_stat_statements';  -- comma-separated list
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_stat_statements;
```
