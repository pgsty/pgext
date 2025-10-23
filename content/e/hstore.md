---
title: "hstore"
linkTitle: "hstore"
description: "data type for storing sets of (key, value) pairs"
weight: 3970
categories: ["TYPE"]
width: full
---

data type for storing sets of (key, value) pairs


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3970** | {{< badge content="hstore" link="https://www.postgresql.org/docs/current/hstore.html" >}} | {{< ext "hstore" >}} | `1.8` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "hstore_pllua" >}} {{< ext "hstore_plluau" >}} {{< ext "hstore_plpython3u" >}} {{< ext "pg_readme" >}} {{< ext "pg_readme_test_extension" >}} |
|   **See Also**    | {{< ext "intarray" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.8" "PostgreSQL 18: version 1.8" "green" >}} | {{< bg "1.8" "PostgreSQL 17: version 1.8" "green" >}} | {{< bg "1.8" "PostgreSQL 16: version 1.8" "green" >}} | {{< bg "1.8" "PostgreSQL 15: version 1.8" "green" >}} | {{< bg "1.8" "PostgreSQL 14: version 1.8" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION hstore;
```
