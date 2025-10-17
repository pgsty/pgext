---
title: "jsonb_plperl"
linkTitle: "jsonb_plperl"
description: "transform between jsonb and plperl"
weight: 3263
categories: ["Lang"]
width: full
---

transform between jsonb and plperl

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3263** | {{< badge content="jsonb_plperl" link="https://www.postgresql.org/docs/current/plperl.html" >}} | {{< ext "jsonb_plperl" "plperl" >}} | `1.0` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="-----d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plperl" >}} |
|   **See Also**    | {{< ext "jsquery" >}} {{< ext "jsonb_plperlu" >}} {{< ext "jsonb_plpython3u" >}} {{< ext "pg_jsonschema" >}} {{< ext "plperlu" >}} {{< ext "plpgsql" >}} |
|    **Siblings**   | {{< ext "bool_plperl" >}} {{< ext "hstore_plperl" >}} {{< ext "plperl" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="1.0" color="green" >}} | {{< badge content="1.0" color="green" >}} | {{< badge content="1.0" color="green" >}} | {{< badge content="1.0" color="green" >}} | {{< badge content="1.0" color="green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION jsonb_plperl;
```
