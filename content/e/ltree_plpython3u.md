---
title: "ltree_plpython3u"
linkTitle: "ltree_plpython3u"
description: "transform between ltree and plpython3u"
weight: 3292
categories: ["Lang"]
width: full
---

transform between ltree and plpython3u

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3292** | {{< badge content="ltree_plpython3u" link="https://www.postgresql.org/docs/current/plpython.html" >}} | {{< ext "ltree_plpython3u" "plpython3u" >}} | `1.0` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "ltree" >}} {{< ext "plpython3u" >}} |
|   **See Also**    | {{< ext "faker" >}} {{< ext "plperl" >}} {{< ext "plpgsql" >}} {{< ext "pg_tle" >}} |
|    **Siblings**   | {{< ext "hstore_plpython3u" >}} {{< ext "jsonb_plpython3u" >}} {{< ext "plpython3u" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="1.0" color="green" >}} | {{< badge content="1.0" color="green" >}} | {{< badge content="1.0" color="green" >}} | {{< badge content="1.0" color="green" >}} | {{< badge content="1.0" color="green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION ltree_plpython3u;
```
