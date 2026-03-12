---
title: "hstore_plperlu"
linkTitle: "hstore_plperlu"
description: "transform between hstore and plperlu"
weight: 3273
categories: ["LANG"]
width: full
---

[**plperlu**](https://www.postgresql.org/docs/current/plperl.html) : transform between hstore and plperlu


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3273** | {{< badge content="hstore_plperlu" link="https://www.postgresql.org/docs/current/plperl.html" >}} | {{< ext "hstore_plperlu" "plperlu" >}} | `1.0` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plperlu" >}} |
|   **See Also**    | {{< ext "hstore_pllua" >}} {{< ext "hstore_plluau" >}} {{< ext "hstore_plperl" >}} {{< ext "hstore_plpython3u" >}} {{< ext "hstore" >}} {{< ext "plperl" >}} {{< ext "plpgsql" >}} |
|    **Siblings**   | {{< ext "plperlu" >}} {{< ext "bool_plperlu" >}} {{< ext "jsonb_plperlu" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION hstore_plperlu CASCADE; -- requires plperlu
```



## Usage

> [hstore_plperlu: Transform between hstore and PL/Perl untrusted](https://www.postgresql.org/docs/current/hstore.html)

Provides a transform for the `hstore` type for PL/Perl Untrusted. When loaded, `hstore` values are automatically converted to Perl hashes and vice versa.

```sql
CREATE EXTENSION hstore_plperlu;

CREATE FUNCTION hstore_to_json_u(val hstore) RETURNS text
LANGUAGE plperlu TRANSFORM FOR TYPE hstore AS $$
  use JSON;
  # val is now a Perl hash reference
  return encode_json($val);
$$;

CREATE FUNCTION make_hstore_u(key text, value text) RETURNS hstore
LANGUAGE plperlu TRANSFORM FOR TYPE hstore AS $$
  my ($k, $v) = @_;
  return { $k => $v };
$$;

SELECT hstore_to_json_u('a=>1, b=>2'::hstore);
SELECT make_hstore_u('color', 'blue');
```
