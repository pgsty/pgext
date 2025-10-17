---
title: "pgstattuple"
linkTitle: "pgstattuple"
description: "show tuple-level statistics"
weight: 6970
categories: ["Stat"]
width: full
---

show tuple-level statistics

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6970** | {{< badge content="pgstattuple" link="https://www.postgresql.org/docs/current/pgstattuple.html" >}} | {{< ext "pgstattuple" "pgstattuple" >}} | `1.5` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pageinspect" >}} {{< ext "pg_freespacemap" >}} {{< ext "pg_visibility" >}} {{< ext "pg_rewrite" >}} {{< ext "pg_checksums" >}} {{< ext "pg_catcheck" >}} {{< ext "amcheck" >}} {{< ext "toastinfo" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="1.5" color="green" >}} | {{< badge content="1.5" color="green" >}} | {{< badge content="1.5" color="green" >}} | {{< badge content="1.5" color="green" >}} | {{< badge content="1.5" color="green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgstattuple;
```
