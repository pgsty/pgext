---
title: "adminpack"
linkTitle: "adminpack"
description: "administrative functions for PostgreSQL"
weight: 5970
categories: ["Admin"]
width: full
---

administrative functions for PostgreSQL

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5970** | {{< badge content="adminpack" link="https://www.postgresql.org/docs/current/adminpack.html" >}} | {{< ext "adminpack" "adminpack" >}} | `2.1` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "fio" >}} {{< ext "lo" >}} {{< ext "file_fdw" >}} {{< ext "ddlx" >}} {{< ext "pgdd" >}} {{< ext "pg_catcheck" >}} {{< ext "pg_cheat_funcs" >}} {{< ext "pg_repack" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="N/A" color="red" >}} | {{< badge content="N/A" color="red" >}} | {{< badge content="2.1" color="green" >}} | {{< badge content="2.1" color="green" >}} | {{< badge content="2.1" color="green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION adminpack;
```
