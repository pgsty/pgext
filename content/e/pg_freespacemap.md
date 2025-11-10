---
title: "pg_freespacemap"
linkTitle: "pg_freespacemap"
description: "examine the free space map (FSM)"
weight: 6950
categories: ["STAT"]
width: full
---

[**pg_freespacemap**](https://www.postgresql.org/docs/current/pgfreespacemap.html) : examine the free space map (FSM)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6950** | {{< badge content="pg_freespacemap" link="https://www.postgresql.org/docs/current/pgfreespacemap.html" >}} | {{< ext "pg_freespacemap" >}} | `1.2` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_relusage" >}} {{< ext "pg_visibility" >}} {{< ext "pgstattuple" >}} {{< ext "amcheck" >}} {{< ext "toastinfo" >}} {{< ext "pageinspect" >}} {{< ext "pg_repack" >}} {{< ext "pg_squeeze" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.2" "PostgreSQL 18: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 17: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 16: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 15: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 14: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 13: version 1.2" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_freespacemap;
```
