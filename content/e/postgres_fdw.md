---
title: "postgres_fdw"
linkTitle: "postgres_fdw"
description: "foreign-data wrapper for remote PostgreSQL servers"
weight: 8990
categories: ["Fdw"]
width: full
---

foreign-data wrapper for remote PostgreSQL servers

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8990** | {{< badge content="postgres_fdw" link="https://www.postgresql.org/docs/current/postgres-fdw.html" >}} | {{< ext "postgres_fdw" "postgres_fdw" >}} | `1.1` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "citus" >}} {{< ext "plproxy" >}} {{< ext "wrappers" >}} {{< ext "pgspider_ext" >}} {{< ext "dblink" >}} {{< ext "mimeo" >}} {{< ext "multicorn" >}} {{< ext "mysql_fdw" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="1.1" color="green" >}} | {{< badge content="1.1" color="green" >}} | {{< badge content="1.1" color="green" >}} | {{< badge content="1.1" color="green" >}} | {{< badge content="1.1" color="green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION postgres_fdw;
```
