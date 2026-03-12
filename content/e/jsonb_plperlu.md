---
title: "jsonb_plperlu"
linkTitle: "jsonb_plperlu"
description: "transform between jsonb and plperlu"
weight: 3272
categories: ["LANG"]
width: full
---

[**plperlu**](https://www.postgresql.org/docs/current/plperl.html) : transform between jsonb and plperlu


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3272** | {{< badge content="jsonb_plperlu" link="https://www.postgresql.org/docs/current/plperl.html" >}} | {{< ext "jsonb_plperlu" "plperlu" >}} | `1.0` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plperlu" >}} |
|   **See Also**    | {{< ext "jsquery" >}} {{< ext "jsonb_plperl" >}} {{< ext "jsonb_plpython3u" >}} {{< ext "pg_jsonschema" >}} {{< ext "plperl" >}} {{< ext "plpgsql" >}} |
|    **Siblings**   | {{< ext "plperlu" >}} {{< ext "bool_plperlu" >}} {{< ext "hstore_plperlu" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION jsonb_plperlu CASCADE; -- requires plperlu
```



## Usage

> [jsonb_plperlu: Transform between jsonb and PL/Perl untrusted](https://www.postgresql.org/docs/current/datatype-json.html)

Provides a transform for the `jsonb` type for PL/Perl Untrusted. When loaded, `jsonb` values are automatically converted to Perl data structures (hashes, arrays, scalars) and vice versa.

```sql
CREATE EXTENSION jsonb_plperlu;

CREATE FUNCTION process_json_u(val jsonb) RETURNS text
LANGUAGE plperlu TRANSFORM FOR TYPE jsonb AS $$
  use JSON;
  # val is already a Perl data structure, re-encode with sorting
  return encode_json($val);
$$;

CREATE FUNCTION build_jsonb_u(name text, age integer) RETURNS jsonb
LANGUAGE plperlu TRANSFORM FOR TYPE jsonb AS $$
  my ($name, $age) = @_;
  return { name => $name, age => $age };
$$;

SELECT process_json_u('{"b": 2, "a": 1}'::jsonb);
SELECT build_jsonb_u('Alice', 30);
```
