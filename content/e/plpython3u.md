---
title: "plpython3u"
linkTitle: "plpython3u"
description: "PL/Python3U untrusted procedural language"
weight: 3290
categories: ["LANG"]
width: full
---

[**plpython3u**](https://www.postgresql.org/docs/current/plpython.html) : PL/Python3U untrusted procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3290** | {{< badge content="plpython3u" link="https://www.postgresql.org/docs/current/plpython.html" >}} | {{< ext "plpython3u" >}} | `1.0` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "hstore_plpython3u" >}} {{< ext "jsonb_plpython3u" >}} {{< ext "ltree_plpython3u" >}} {{< ext "omni_python" >}} {{< ext "pg4ml" >}} |
|   **See Also**    | {{< ext "faker" >}} {{< ext "plv8" >}} {{< ext "pllua" >}} {{< ext "plluau" >}} {{< ext "pltcl" >}} {{< ext "pltclu" >}} {{< ext "plperl" >}} {{< ext "plperlu" >}} |
|    **Siblings**   | {{< ext "jsonb_plpython3u" >}} {{< ext "ltree_plpython3u" >}} {{< ext "hstore_plpython3u" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 13: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plpython3u;
```
