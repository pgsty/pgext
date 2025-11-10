---
title: "seg"
linkTitle: "seg"
description: "data type for representing line segments or floating-point intervals"
weight: 3940
categories: ["TYPE"]
width: full
---

[**seg**](https://www.postgresql.org/docs/current/seg.html)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3940** | {{< badge content="seg" link="https://www.postgresql.org/docs/current/seg.html" >}} | {{< ext "seg" >}} | `1.4` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "cube" >}} {{< ext "intarray" >}} {{< ext "intagg" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.4" "PostgreSQL 18: version 1.4" "green" >}} | {{< bg "1.4" "PostgreSQL 17: version 1.4" "green" >}} | {{< bg "1.4" "PostgreSQL 16: version 1.4" "green" >}} | {{< bg "1.4" "PostgreSQL 15: version 1.4" "green" >}} | {{< bg "1.4" "PostgreSQL 14: version 1.4" "green" >}} | {{< bg "1.4" "PostgreSQL 13: version 1.4" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION seg;
```
