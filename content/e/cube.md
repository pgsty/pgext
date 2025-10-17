---
title: "cube"
linkTitle: "cube"
description: "data type for multidimensional cubes"
weight: 3950
categories: ["Type"]
width: full
---

data type for multidimensional cubes

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3950** | {{< badge content="cube" link="https://www.postgresql.org/docs/current/cube.html" >}} | {{< ext "cube" "cube" >}} | `1.5` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "earthdistance" >}} {{< ext "pg4ml" >}} |
|   **See Also**    | {{< ext "seg" >}} {{< ext "intarray" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="1.5" color="green" >}} | {{< badge content="1.5" color="green" >}} | {{< badge content="1.5" color="green" >}} | {{< badge content="1.5" color="green" >}} | {{< badge content="1.5" color="green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION cube;
```
