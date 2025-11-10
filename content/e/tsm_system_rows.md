---
title: "tsm_system_rows"
linkTitle: "tsm_system_rows"
description: "TABLESAMPLE method which accepts number of rows as a limit"
weight: 4910
categories: ["FUNC"]
width: full
---

[**tsm_system_rows**](https://www.postgresql.org/docs/current/tsm-system-rows.html) : TABLESAMPLE method which accepts number of rows as a limit


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4910** | {{< badge content="tsm_system_rows" link="https://www.postgresql.org/docs/current/tsm-system-rows.html" >}} | {{< ext "tsm_system_rows" >}} | `1.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "documentdb" >}} |
|   **See Also**    | {{< ext "random" >}} {{< ext "permuteseq" >}} {{< ext "tsm_system_time" >}} {{< ext "pg_crash" >}} {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "pg_hashids" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 13: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION tsm_system_rows;
```
