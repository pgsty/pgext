---
title: "pg_buffercache"
linkTitle: "pg_buffercache"
description: "examine the shared buffer cache"
weight: 6930
categories: ["Stat"]
width: full
---

examine the shared buffer cache

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6930** | {{< badge content="pg_buffercache" link="https://www.postgresql.org/docs/current/pgbuffercache.html" >}} | {{< ext "pg_buffercache" "pg_buffercache" >}} | `1.5` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_repack" >}} {{< ext "pgfincore" >}} {{< ext "pgcozy" >}} {{< ext "pg_prewarm" >}} {{< ext "pgmeminfo" >}} {{< ext "pg_squeeze" >}} {{< ext "old_snapshot" >}} {{< ext "system_stats" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="1.5" color="green" >}} | {{< badge content="1.5" color="green" >}} | {{< badge content="1.5" color="green" >}} | {{< badge content="1.5" color="green" >}} | {{< badge content="1.5" color="green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_buffercache;
```
