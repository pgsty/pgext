---
title: "plperlu"
linkTitle: "plperlu"
description: "PL/PerlU untrusted procedural language"
weight: 3270
categories: ["LANG"]
width: full
---

PL/PerlU untrusted procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3270** | {{< badge content="plperlu" link="https://www.postgresql.org/docs/current/plperl.html" >}} | {{< ext "plperlu" >}} | `1.0` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plperlu" >}} |
|    **Need By**    | {{< ext "bool_plperlu" >}} {{< ext "hstore_plperlu" >}} {{< ext "jsonb_plperlu" >}} {{< ext "plperlu" >}} {{< ext "sparql" >}} |
|   **See Also**    | {{< ext "plperl" >}} {{< ext "plluau" >}} {{< ext "pltclu" >}} {{< ext "bool_plperl" >}} {{< ext "hstore_plperl" >}} {{< ext "jsonb_plperl" >}} {{< ext "plpgsql" >}} {{< ext "plv8" >}} |
|    **Siblings**   | {{< ext "bool_plperlu" >}} {{< ext "jsonb_plperlu" >}} {{< ext "hstore_plperlu" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 13: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plperlu;
```
