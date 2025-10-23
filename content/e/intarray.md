---
title: "intarray"
linkTitle: "intarray"
description: "functions, operators, and index support for 1-D arrays of integers"
weight: 4960
categories: ["FUNC"]
width: full
---

functions, operators, and index support for 1-D arrays of integers


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4960** | {{< badge content="intarray" link="https://www.postgresql.org/docs/current/intarray.html" >}} | {{< ext "intarray" >}} | `1.5` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "aggs_for_arrays" >}} {{< ext "aggs_for_vecs" >}} {{< ext "arraymath" >}} {{< ext "floatvec" >}} {{< ext "vector" >}} {{< ext "vchord" >}} {{< ext "vectorscale" >}} {{< ext "vectorize" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.5" "PostgreSQL 18: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 17: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 16: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 15: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 14: version 1.5" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION intarray;
```
