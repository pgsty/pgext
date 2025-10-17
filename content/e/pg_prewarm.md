---
title: "pg_prewarm"
linkTitle: "pg_prewarm"
description: "prewarm relation data"
weight: 5890
categories: ["Admin"]
width: full
---

prewarm relation data

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5890** | {{< badge content="pg_prewarm" link="https://www.postgresql.org/docs/current/pgprewarm.html" >}} | {{< ext "pg_prewarm" "pg_prewarm" >}} | `1.2` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgfincore" >}} {{< ext "pg_cooldown" >}} {{< ext "pgcozy" >}} {{< ext "pg_buffercache" >}} {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} {{< ext "pg_squeeze" >}} {{< ext "old_snapshot" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="1.2" color="green" >}} | {{< badge content="1.2" color="green" >}} | {{< badge content="1.2" color="green" >}} | {{< badge content="1.2" color="green" >}} | {{< badge content="1.2" color="green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

Add this extension to [`shared_preload_libraries`](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pg_prewarm';  -- comma-separated list
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_prewarm;
```
