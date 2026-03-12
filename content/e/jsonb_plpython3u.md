---
title: "jsonb_plpython3u"
linkTitle: "jsonb_plpython3u"
description: "transform between jsonb and plpython3u"
weight: 3291
categories: ["LANG"]
width: full
---

[**plpython3u**](https://www.postgresql.org/docs/current/plpython.html) : transform between jsonb and plpython3u


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3291** | {{< badge content="jsonb_plpython3u" link="https://www.postgresql.org/docs/current/plpython.html" >}} | {{< ext "jsonb_plpython3u" "plpython3u" >}} | `1.0` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpython3u" >}} |
|   **See Also**    | {{< ext "faker" >}} {{< ext "jsonb_plperl" >}} {{< ext "jsonb_plperlu" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "plpgsql" >}} |
|    **Siblings**   | {{< ext "plpython3u" >}} {{< ext "ltree_plpython3u" >}} {{< ext "hstore_plpython3u" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION jsonb_plpython3u CASCADE; -- requires plpython3u
```



## Usage

> [jsonb_plpython3u: Transform between jsonb and PL/Python3](https://www.postgresql.org/docs/current/datatype-json.html)

Provides a transform for the `jsonb` type for PL/Python3U. When loaded, `jsonb` values are automatically converted to Python dicts, lists, and scalars, and vice versa.

```sql
CREATE EXTENSION jsonb_plpython3u;

CREATE FUNCTION jsonb_keys_sorted(val jsonb) RETURNS text[]
LANGUAGE plpython3u TRANSFORM FOR TYPE jsonb AS $$
  # val is now a Python dict
  return sorted(val.keys())
$$;

CREATE FUNCTION build_jsonb(name text, age integer) RETURNS jsonb
LANGUAGE plpython3u TRANSFORM FOR TYPE jsonb AS $$
  return {"name": name, "age": age}
$$;

SELECT jsonb_keys_sorted('{"b": 2, "a": 1, "c": 3}'::jsonb);
SELECT build_jsonb('Alice', 30);
```
