---
title: "pltclu"
linkTitle: "pltclu"
description: "PL/TclU untrusted procedural language"
weight: 3250
categories: ["LANG"]
width: full
---

[**pltcl**](https://www.postgresql.org/docs/current/pltcl.html) : PL/TclU untrusted procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3250** | {{< badge content="pltclu" link="https://www.postgresql.org/docs/current/pltcl.html" >}} | {{< ext "pltclu" "pltcl" >}} | `1.0` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|   **See Also**    | {{< ext "plpgsql" >}} {{< ext "plperlu" >}} {{< ext "plpython3u" >}} {{< ext "plv8" >}} {{< ext "plluau" >}} {{< ext "pljava" >}} {{< ext "pg_tle" >}} |
|    **Siblings**   | {{< ext "pltcl" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pltclu;
```



## Usage

> [pltclu: PL/Tcl untrusted procedural language](https://www.postgresql.org/docs/current/pltcl.html)

PL/Tcl Untrusted provides full Tcl capabilities including filesystem access and external program execution. Only superusers can create functions in this language.

```sql
CREATE EXTENSION pltclu;

-- Read a file from the server filesystem
CREATE FUNCTION read_file(filename text) RETURNS text
LANGUAGE pltclu AS $$
  set fd [open $1 r]
  set content [read $fd]
  close $fd
  return $content
$$;

-- Execute an external command
CREATE FUNCTION run_command(cmd text) RETURNS text
LANGUAGE pltclu AS $$
  return [exec {*}$1]
$$;

-- Access environment variables
CREATE FUNCTION get_env(varname text) RETURNS text
LANGUAGE pltclu AS $$
  if {[info exists ::env($1)]} {
    return $::env($1)
  }
  return ""
$$;

SELECT get_env('HOME');
```
