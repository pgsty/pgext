---
title: "auto_explain"
linkTitle: "auto_explain"
description: "Provides a means for logging execution plans of slow statements automatically"
weight: 6980
categories: ["Stat"]
width: full
---

Provides a means for logging execution plans of slow statements automatically

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6980** | {{< badge content="auto_explain" link="https://www.postgresql.org/docs/current/auto-explain.html" >}} | {{< ext "auto_explain" "auto_explain" >}} | `-` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sL---" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="No" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_show_plans" >}} {{< ext "pg_store_plans" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_hint_plan" >}} {{< ext "plprofiler" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_track_settings" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="-" color="green" >}} | {{< badge content="-" color="green" >}} | {{< badge content="-" color="green" >}} | {{< badge content="-" color="green" >}} | {{< badge content="-" color="green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

Add this extension to [`shared_preload_libraries`](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'auto_explain';  -- comma-separated list
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION auto_explain;
```
