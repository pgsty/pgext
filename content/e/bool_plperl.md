---
title: "bool_plperl"
linkTitle: "bool_plperl"
description: "transform between bool and plperl"
weight: 3261
categories: ["LANG"]
width: full
---

[**plperl**](https://www.postgresql.org/docs/current/plperl.html) : transform between bool and plperl


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3261** | {{< badge content="bool_plperl" link="https://www.postgresql.org/docs/current/plperl.html" >}} | {{< ext "bool_plperl" "plperl" >}} | `1.0` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plperl" >}} |
|   **See Also**    | {{< ext "plperlu" >}} {{< ext "bool_plperlu" >}} {{< ext "plpgsql" >}} {{< ext "pg_tle" >}} {{< ext "plv8" >}} |
|    **Siblings**   | {{< ext "plperl" >}} {{< ext "hstore_plperl" >}} {{< ext "jsonb_plperl" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION bool_plperl CASCADE; -- requires plperl
```



## Usage

> [bool_plperl: Transform between bool and PL/Perl](https://www.postgresql.org/docs/current/plperl.html)

Provides a transform for the `bool` type for PL/Perl. When loaded, PostgreSQL `boolean` values are automatically converted to Perl native boolean representations and vice versa, instead of being passed as strings.

```sql
CREATE EXTENSION bool_plperl;

CREATE FUNCTION check_flag(val boolean) RETURNS text
LANGUAGE plperl TRANSFORM FOR TYPE boolean AS $$
  # val is a native Perl boolean (1 or undef), not a string
  if ($_[0]) {
    return "flag is set";
  }
  return "flag is not set";
$$;

SELECT check_flag(true);   -- flag is set
SELECT check_flag(false);  -- flag is not set
```
