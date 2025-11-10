---
title: "hstore_plperl"
linkTitle: "hstore_plperl"
description: "transform between hstore and plperl"
weight: 3262
categories: ["LANG"]
width: full
---

[**plperl**](https://www.postgresql.org/docs/current/plperl.html)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3262** | {{< badge content="hstore_plperl" link="https://www.postgresql.org/docs/current/plperl.html" >}} | {{< ext "hstore_plperl" "plperl" >}} | `1.0` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plperl" >}} |
|   **See Also**    | {{< ext "hstore_pllua" >}} {{< ext "hstore_plluau" >}} {{< ext "hstore_plperlu" >}} {{< ext "hstore_plpython3u" >}} {{< ext "hstore" >}} {{< ext "plperlu" >}} {{< ext "plpgsql" >}} |
|    **Siblings**   | {{< ext "plperl" >}} {{< ext "bool_plperl" >}} {{< ext "jsonb_plperl" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 13: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION hstore_plperl;
```
