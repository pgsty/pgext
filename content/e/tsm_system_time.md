---
title: "tsm_system_time"
linkTitle: "tsm_system_time"
description: "TABLESAMPLE method which accepts time in milliseconds as a limit"
weight: 4890
categories: ["FUNC"]
width: full
---

[**tsm_system_time**](https://www.postgresql.org/docs/current/tsm-system-time.html) : TABLESAMPLE method which accepts time in milliseconds as a limit


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4890** | {{< badge content="tsm_system_time" link="https://www.postgresql.org/docs/current/tsm-system-time.html" >}} | {{< ext "tsm_system_time" >}} | `1.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "random" >}} {{< ext "permuteseq" >}} {{< ext "tsm_system_rows" >}} {{< ext "pg_crash" >}} {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "pg_hashids" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION tsm_system_time;
```



## Usage

> [tsm_system_time: time-based TABLESAMPLE method](https://www.postgresql.org/docs/current/tsm-system-time.html)

Provides the `SYSTEM_TIME` table sampling method that returns as many rows as can be read within a specified time limit.

```sql
CREATE EXTENSION tsm_system_time;
```

### TABLESAMPLE Method

`SYSTEM_TIME(milliseconds float)` -- maximum time to spend reading the table.

### Examples

```sql
-- Sample rows readable within 1 second
SELECT * FROM my_table TABLESAMPLE SYSTEM_TIME(1000);

-- Sample from a large table with a 500ms budget
SELECT count(*) FROM large_table TABLESAMPLE SYSTEM_TIME(500);
```

Performs block-level sampling (not row-level). If the entire table can be read within the time limit, all rows are returned. Does not support `REPEATABLE`.
