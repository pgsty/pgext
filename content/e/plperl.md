---
title: "plperl"
linkTitle: "plperl"
description: "PL/Perl procedural language"
weight: 3260
categories: ["Lang"]
width: full
---

PL/Perl procedural language

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3260** | {{< badge content="plperl" link="https://www.postgresql.org/docs/current/plperl.html" >}} | {{< ext "plperl" "plperl" >}} | `1.0` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plperl" >}} |
|    **Need By**    | {{< ext "bool_plperl" >}} {{< ext "hstore_plperl" >}} {{< ext "jsonb_plperl" >}} {{< ext "plperl" >}} {{< ext "sparql" >}} |
|   **See Also**    | {{< ext "plperlu" >}} {{< ext "bool_plperlu" >}} {{< ext "jsonb_plperlu" >}} {{< ext "hstore_plperlu" >}} {{< ext "plpgsql" >}} {{< ext "pg_tle" >}} {{< ext "plv8" >}} {{< ext "pllua" >}} |
|    **Siblings**   | {{< ext "bool_plperl" >}} {{< ext "hstore_plperl" >}} {{< ext "jsonb_plperl" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="1.0" color="green" >}} | {{< badge content="1.0" color="green" >}} | {{< badge content="1.0" color="green" >}} | {{< badge content="1.0" color="green" >}} | {{< badge content="1.0" color="green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plperl;
```
