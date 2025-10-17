---
title: "btree_gin"
linkTitle: "btree_gin"
description: "support for indexing common datatypes in GIN"
weight: 4950
categories: ["Func"]
width: full
---

support for indexing common datatypes in GIN

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4950** | {{< badge content="btree_gin" link="https://www.postgresql.org/docs/current/btree-gin.html" >}} | {{< ext "btree_gin" "btree_gin" >}} | `1.3` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "btree_gist" >}} {{< ext "unaccent" >}} {{< ext "fuzzystrmatch" >}} {{< ext "pg_trgm" >}} {{< ext "prefix" >}} {{< ext "citext" >}} {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="1.3" color="green" >}} | {{< badge content="1.3" color="green" >}} | {{< badge content="1.3" color="green" >}} | {{< badge content="1.3" color="green" >}} | {{< badge content="1.3" color="green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION btree_gin;
```
