---
title: "postgres_fdw"
linkTitle: "postgres_fdw"
description: "foreign-data wrapper for remote PostgreSQL servers"
weight: 8990
categories: ["FDW"]
width: full
---

[**postgres_fdw**](https://www.postgresql.org/docs/current/postgres-fdw.html)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8990** | {{< badge content="postgres_fdw" link="https://www.postgresql.org/docs/current/postgres-fdw.html" >}} | {{< ext "postgres_fdw" >}} | `1.1` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "omni_schema" >}} |
|   **See Also**    | {{< ext "citus" >}} {{< ext "plproxy" >}} {{< ext "wrappers" >}} {{< ext "pgspider_ext" >}} {{< ext "dblink" >}} {{< ext "mimeo" >}} {{< ext "multicorn" >}} {{< ext "mysql_fdw" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.1" "PostgreSQL 18: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 17: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 16: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 15: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 14: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 13: version 1.1" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION postgres_fdw;
```
