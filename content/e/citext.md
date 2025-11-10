---
title: "citext"
linkTitle: "citext"
description: "data type for case-insensitive character strings"
weight: 3980
categories: ["TYPE"]
width: full
---

[**citext**](https://www.postgresql.org/docs/current/citext.html)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3980** | {{< badge content="citext" link="https://www.postgresql.org/docs/current/citext.html" >}} | {{< ext "citext" >}} | `1.6` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "ltree" >}} {{< ext "unaccent" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.6" "PostgreSQL 18: version 1.6" "green" >}} | {{< bg "1.6" "PostgreSQL 17: version 1.6" "green" >}} | {{< bg "1.6" "PostgreSQL 16: version 1.6" "green" >}} | {{< bg "1.6" "PostgreSQL 15: version 1.6" "green" >}} | {{< bg "1.6" "PostgreSQL 14: version 1.6" "green" >}} | {{< bg "1.6" "PostgreSQL 13: version 1.6" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION citext;
```
