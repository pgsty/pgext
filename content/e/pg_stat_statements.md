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

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.11" "PostgreSQL 18: version 1.11" "green" >}} | {{< bg "1.11" "PostgreSQL 17: version 1.11" "green" >}} | {{< bg "1.11" "PostgreSQL 16: version 1.11" "green" >}} | {{< bg "1.11" "PostgreSQL 15: version 1.11" "green" >}} | {{< bg "1.11" "PostgreSQL 14: version 1.11" "green" >}} | {{< bg "1.11" "PostgreSQL 13: version 1.11" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pg_stat_statements';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_stat_statements;
```
