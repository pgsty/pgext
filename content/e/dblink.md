---
title: "dblink"
linkTitle: "dblink"
description: "connect to other PostgreSQL databases from within a database"
weight: 8970
categories: ["FDW"]
width: full
---

[**dblink**](https://www.postgresql.org/docs/current/dblink.html) : connect to other PostgreSQL databases from within a database


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8970** | {{< badge content="dblink" link="https://www.postgresql.org/docs/current/dblink.html" >}} | {{< ext "dblink" >}} | `1.2` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "emaj" >}} {{< ext "mimeo" >}} {{< ext "omni_schema" >}} {{< ext "omni_test" >}} {{< ext "omni_vfs" >}} {{< ext "pg_jobmon" >}} {{< ext "pg_profile" >}} |
|   **See Also**    | {{< ext "plproxy" >}} {{< ext "pgbouncer_fdw" >}} {{< ext "postgres_fdw" >}} {{< ext "citus" >}} {{< ext "wrappers" >}} {{< ext "pgspider_ext" >}} {{< ext "pglogical" >}} {{< ext "repmgr" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.2" "PostgreSQL 18: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 17: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 16: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 15: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 14: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 13: version 1.2" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION dblink;
```
