---
title: "jsonb_plperl"
linkTitle: "jsonb_plperl"
description: "transform between jsonb and plperl"
weight: 3263
categories: ["LANG"]
width: full
---

[**plperl**](https://www.postgresql.org/docs/current/plperl.html) : transform between jsonb and plperl


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3263** | {{< badge content="jsonb_plperl" link="https://www.postgresql.org/docs/current/plperl.html" >}} | {{< ext "jsonb_plperl" "plperl" >}} | `1.0` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plperl" >}} |
|   **See Also**    | {{< ext "jsquery" >}} {{< ext "jsonb_plperlu" >}} {{< ext "jsonb_plpython3u" >}} {{< ext "pg_jsonschema" >}} {{< ext "plperlu" >}} {{< ext "plpgsql" >}} |
|    **Siblings**   | {{< ext "plperl" >}} {{< ext "bool_plperl" >}} {{< ext "hstore_plperl" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION jsonb_plperl CASCADE; -- requires plperl
```



## Usage

> [jsonb_plperl: Transform between jsonb and PL/Perl](https://www.postgresql.org/docs/current/datatype-json.html)

Provides a transform for the `jsonb` type for PL/Perl. When loaded, `jsonb` values are automatically converted to Perl data structures (hashes, arrays, scalars) and vice versa.

```sql
CREATE EXTENSION jsonb_plperl;

CREATE FUNCTION jsonb_keys_sorted(val jsonb) RETURNS text
LANGUAGE plperl TRANSFORM FOR TYPE jsonb AS $$
  # val is now a Perl hash/array reference
  return join(', ', sort keys %$val);
$$;

CREATE FUNCTION build_jsonb(name text, age integer) RETURNS jsonb
LANGUAGE plperl TRANSFORM FOR TYPE jsonb AS $$
  my ($name, $age) = @_;
  return { name => $name, age => $age };
$$;

SELECT jsonb_keys_sorted('{"b": 2, "a": 1, "c": 3}'::jsonb);
SELECT build_jsonb('Alice', 30);
```
