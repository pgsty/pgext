---
title: "pg_walinspect"
linkTitle: "pg_walinspect"
description: "functions to inspect contents of PostgreSQL Write-Ahead Log"
weight: 6940
categories: ["STAT"]
width: full
---

[**pg_walinspect**](https://www.postgresql.org/docs/current/pgwalinspect.html) : functions to inspect contents of PostgreSQL Write-Ahead Log


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6940** | {{< badge content="pg_walinspect" link="https://www.postgresql.org/docs/current/pgwalinspect.html" >}} | {{< ext "pg_walinspect" >}} | `1.1` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "basic_archive" >}} {{< ext "pglogical" >}} {{< ext "pg_failover_slots" >}} {{< ext "wal2json" >}} {{< ext "basebackup_to_shell" >}} {{< ext "decoderbufs" >}} {{< ext "test_decoding" >}} {{< ext "pg_profile" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.1" "PostgreSQL 18: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 17: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 16: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 15: version 1.1" "green" >}} | {{< bg "N/A" "PostgreSQL 14: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 13: not available" "red" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_walinspect;
```
