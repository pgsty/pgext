---
title: "file_fdw"
linkTitle: "file_fdw"
description: "foreign-data wrapper for flat file access"
weight: 8980
categories: ["FDW"]
width: full
---

[**file_fdw**](https://www.postgresql.org/docs/current/file-fdw.html) : foreign-data wrapper for flat file access


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8980** | {{< badge content="file_fdw" link="https://www.postgresql.org/docs/current/file-fdw.html" >}} | {{< ext "file_fdw" >}} | `1.0` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "pg_sqlog" >}} |
|   **See Also**    | {{< ext "log_fdw" >}} {{< ext "wrappers" >}} {{< ext "sqlite_fdw" >}} {{< ext "aws_s3" >}} {{< ext "pg_bulkload" >}} {{< ext "multicorn" >}} {{< ext "hdfs_fdw" >}} {{< ext "postgres_fdw" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 13: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION file_fdw;
```
