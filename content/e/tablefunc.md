---
title: "tablefunc"
linkTitle: "tablefunc"
description: "functions that manipulate whole tables, including crosstab"
weight: 2590
categories: ["OLAP"]
width: full
---

[**tablefunc**](https://www.postgresql.org/docs/current/tablefunc.html) : functions that manipulate whole tables, including crosstab


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2590** | {{< badge content="tablefunc" link="https://www.postgresql.org/docs/current/tablefunc.html" >}} | {{< ext "tablefunc" >}} | `1.0` | {{< category "OLAP" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "pg4ml" >}} |
|   **See Also**    | {{< ext "cube" >}} {{< ext "plr" >}} {{< ext "orafce" >}} {{< ext "timescaledb" >}} {{< ext "citus" >}} {{< ext "pg_partman" >}} {{< ext "citus_columnar" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 13: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION tablefunc;
```
