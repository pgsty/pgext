---
title: "pageinspect"
linkTitle: "pageinspect"
description: "inspect the contents of database pages at a low level"
weight: 6900
categories: ["STAT"]
width: full
---

inspect the contents of database pages at a low level


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6900** | {{< badge content="pageinspect" link="https://www.postgresql.org/docs/current/pageinspect.html" >}} | {{< ext "pageinspect" >}} | `1.12` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "amcheck" >}} {{< ext "pagevis" >}} {{< ext "pg_visibility" >}} {{< ext "pg_repack" >}} {{< ext "pg_squeeze" >}} {{< ext "pg_dirtyread" >}} {{< ext "pgdd" >}} {{< ext "pg_orphaned" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.12" "PostgreSQL 18: version 1.12" "green" >}} | {{< bg "1.12" "PostgreSQL 17: version 1.12" "green" >}} | {{< bg "1.12" "PostgreSQL 16: version 1.12" "green" >}} | {{< bg "1.12" "PostgreSQL 15: version 1.12" "green" >}} | {{< bg "1.12" "PostgreSQL 14: version 1.12" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pageinspect;
```
