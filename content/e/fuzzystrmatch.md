---
title: "fuzzystrmatch"
linkTitle: "fuzzystrmatch"
description: "determine similarities and distance between strings"
weight: 2180
categories: ["FTS"]
width: full
---

[**fuzzystrmatch**](https://www.postgresql.org/docs/current/fuzzystrmatch.html)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2180** | {{< badge content="fuzzystrmatch" link="https://www.postgresql.org/docs/current/fuzzystrmatch.html" >}} | {{< ext "fuzzystrmatch" >}} | `1.2` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "postgis_tiger_geocoder" >}} |
|   **See Also**    | {{< ext "pg_similarity" >}} {{< ext "smlar" >}} {{< ext "pg_trgm" >}} {{< ext "unaccent" >}} {{< ext "pg_bigm" >}} {{< ext "citext" >}} {{< ext "btree_gist" >}} {{< ext "btree_gin" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.2" "PostgreSQL 18: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 17: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 16: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 15: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 14: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 13: version 1.2" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION fuzzystrmatch;
```
