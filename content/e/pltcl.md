---
title: "pltcl"
linkTitle: "pltcl"
description: "PL/Tcl procedural language"
weight: 3240
categories: ["LANG"]
width: full
---

[**pltcl**](https://www.postgresql.org/docs/current/pltcl.html) : PL/Tcl procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3240** | {{< badge content="pltcl" link="https://www.postgresql.org/docs/current/pltcl.html" >}} | {{< ext "pltcl" >}} | `1.0` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|   **See Also**    | {{< ext "plpgsql" >}} {{< ext "plperl" >}} {{< ext "plpython3u" >}} {{< ext "pg_tle" >}} {{< ext "plv8" >}} {{< ext "pllua" >}} {{< ext "pljava" >}} |
|    **Siblings**   | {{< ext "pltclu" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pltcl;
```



## Usage

> [pltcl: PL/Tcl trusted procedural language](https://www.postgresql.org/docs/current/pltcl.html)

PL/Tcl allows writing PostgreSQL functions in the Tcl programming language. As a trusted language, it restricts access to the filesystem and other system resources.

```sql
CREATE EXTENSION pltcl;

-- Simple function returning a greeting
CREATE FUNCTION tcl_hello(name text) RETURNS text
LANGUAGE pltcl AS $$
  return "Hello, $1!"
$$;

SELECT tcl_hello('world');

-- Function with conditional logic
CREATE FUNCTION tcl_max(a integer, b integer) RETURNS integer
LANGUAGE pltcl AS $$
  if {$1 > $2} {return $1}
  return $2
$$;

-- Set-returning function
CREATE FUNCTION tcl_sequence(count integer) RETURNS SETOF integer
LANGUAGE pltcl AS $$
  for {set i 1} {$i <= $1} {incr i} {
    return_next $i
  }
$$;

-- Trigger function
CREATE FUNCTION tcl_audit() RETURNS trigger
LANGUAGE pltcl AS $$
  set NEW(modified_at) [clock format [clock seconds] -format "%Y-%m-%d %H:%M:%S"]
  return [array get NEW]
$$;
```

Database access from PL/Tcl uses the `spi_exec` command:

```sql
CREATE FUNCTION tcl_count_rows(tbl text) RETURNS integer
LANGUAGE pltcl AS $$
  spi_exec "SELECT count(*) AS cnt FROM $1"
  return $cnt
$$;
```
