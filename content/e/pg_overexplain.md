---
title: "pg_overexplain"
linkTitle: "pg_overexplain"
description: "Allow EXPLAIN to dump even more details"
weight: 6880
categories: ["STAT"]
width: full
---

[**pg_overexplain**](https://www.postgresql.org/docs/devel/pgoverexplain.html) : Allow EXPLAIN to dump even more details


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6880** | {{< badge content="pg_overexplain" link="https://www.postgresql.org/docs/devel/pgoverexplain.html" >}} | {{< ext "pg_overexplain" >}} | `1.0` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_profile" >}} {{< ext "pg_tracing" >}} {{< ext "pg_show_plans" >}} {{< ext "pg_stat_kcache" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_store_plans" >}} {{< ext "pg_track_settings" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "N/A" "PostgreSQL 17: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 16: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 15: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 14: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 13: not available" "red" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_overexplain';
```


This extension does not need `CREATE EXTENSION` DDL command


