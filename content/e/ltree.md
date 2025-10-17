---
title: "ltree"
linkTitle: "ltree"
description: "data type for hierarchical tree-like structures"
weight: 3960
categories: ["Type"]
width: full
---

data type for hierarchical tree-like structures

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3960** | {{< badge content="ltree" link="https://www.postgresql.org/docs/current/ltree.html" >}} | {{< ext "ltree" "ltree" >}} | `1.3` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "ltree_plpython3u" >}} |
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "citext" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="1.3" color="green" >}} | {{< badge content="1.3" color="green" >}} | {{< badge content="1.3" color="green" >}} | {{< badge content="1.3" color="green" >}} | {{< badge content="1.3" color="green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION ltree;
```
