---
title: "btree_gist"
linkTitle: "btree_gist"
description: "support for indexing common datatypes in GiST"
weight: 4940
categories: ["FUNC"]
width: full
---

[**btree_gist**](https://www.postgresql.org/docs/current/btree-gist.html) : support for indexing common datatypes in GiST


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4940** | {{< badge content="btree_gist" link="https://www.postgresql.org/docs/current/btree-gist.html" >}} | {{< ext "btree_gist" >}} | `1.7` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "emaj" >}} {{< ext "omni_auth" >}} {{< ext "periods" >}} {{< ext "pgautofailover" >}} {{< ext "powa" >}} |
|   **See Also**    | {{< ext "btree_gin" >}} {{< ext "unaccent" >}} {{< ext "fuzzystrmatch" >}} {{< ext "pg_trgm" >}} {{< ext "prefix" >}} {{< ext "citext" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.7" "PostgreSQL 18: version 1.7" "green" >}} | {{< bg "1.7" "PostgreSQL 17: version 1.7" "green" >}} | {{< bg "1.7" "PostgreSQL 16: version 1.7" "green" >}} | {{< bg "1.7" "PostgreSQL 15: version 1.7" "green" >}} | {{< bg "1.7" "PostgreSQL 14: version 1.7" "green" >}} | {{< bg "1.7" "PostgreSQL 13: version 1.7" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION btree_gist;
```
