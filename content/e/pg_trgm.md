---
title: "pg_trgm"
linkTitle: "pg_trgm"
description: "text similarity measurement and index searching based on trigrams"
weight: 2190
categories: ["FTS"]
width: full
---

[**pg_trgm**](https://www.postgresql.org/docs/current/pgtrgm.html)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2190** | {{< badge content="pg_trgm" link="https://www.postgresql.org/docs/current/pgtrgm.html" >}} | {{< ext "pg_trgm" >}} | `1.6` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_similarity" >}} {{< ext "pg_bigm" >}} {{< ext "fuzzystrmatch" >}} {{< ext "unaccent" >}} {{< ext "smlar" >}} {{< ext "pgroonga_database" >}} {{< ext "rum" >}} {{< ext "citext" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.6" "PostgreSQL 18: version 1.6" "green" >}} | {{< bg "1.6" "PostgreSQL 17: version 1.6" "green" >}} | {{< bg "1.6" "PostgreSQL 16: version 1.6" "green" >}} | {{< bg "1.6" "PostgreSQL 15: version 1.6" "green" >}} | {{< bg "1.6" "PostgreSQL 14: version 1.6" "green" >}} | {{< bg "1.6" "PostgreSQL 13: version 1.6" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_trgm;
```
