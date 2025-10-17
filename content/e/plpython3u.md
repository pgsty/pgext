---
title: "plpython3u"
linkTitle: "plpython3u"
description: "PL/Python3U untrusted procedural language"
weight: 3290
categories: ["Lang"]
width: full
---

PL/Python3U untrusted procedural language

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3290** | {{< badge content="plpython3u" link="https://www.postgresql.org/docs/current/plpython.html" >}} | {{< ext "plpython3u" "plpython3u" >}} | `1.0` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "hstore_plpython3u" >}} {{< ext "jsonb_plpython3u" >}} {{< ext "ltree_plpython3u" >}} {{< ext "omni_python" >}} {{< ext "pg4ml" >}} |
|   **See Also**    | {{< ext "faker" >}} {{< ext "plv8" >}} {{< ext "pllua" >}} {{< ext "plluau" >}} {{< ext "pltcl" >}} {{< ext "pltclu" >}} {{< ext "plperl" >}} {{< ext "plperlu" >}} |
|    **Siblings**   | {{< ext "hstore_plpython3u" >}} {{< ext "jsonb_plpython3u" >}} {{< ext "ltree_plpython3u" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="1.0" color="green" >}} | {{< badge content="1.0" color="green" >}} | {{< badge content="1.0" color="green" >}} | {{< badge content="1.0" color="green" >}} | {{< badge content="1.0" color="green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plpython3u CASCADE SCHEMA pg_catalog;
```
