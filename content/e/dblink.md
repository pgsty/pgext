---
title: "dblink"
linkTitle: "dblink"
description: "connect to other PostgreSQL databases from within a database"
weight: 8970
categories: ["Fdw"]
width: full
---

connect to other PostgreSQL databases from within a database

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8970** | {{< badge content="dblink" link="https://www.postgresql.org/docs/current/dblink.html" >}} | {{< ext "dblink" "dblink" >}} | `1.2` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "emaj" >}} {{< ext "mimeo" >}} {{< ext "omni_schema" >}} {{< ext "omni_test" >}} {{< ext "omni_vfs_types_v1" >}} {{< ext "pg_jobmon" >}} {{< ext "pg_profile" >}} |
|   **See Also**    | {{< ext "plproxy" >}} {{< ext "pgbouncer_fdw" >}} {{< ext "postgres_fdw" >}} {{< ext "citus" >}} {{< ext "wrappers" >}} {{< ext "pgspider_ext" >}} {{< ext "pglogical" >}} {{< ext "repmgr" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="1.2" color="green" >}} | {{< badge content="1.2" color="green" >}} | {{< badge content="1.2" color="green" >}} | {{< badge content="1.2" color="green" >}} | {{< badge content="1.2" color="green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION dblink;
```
