---
title: "plperl"
linkTitle: "plperl"
description: "PL/Perl procedural language"
weight: 3260
categories: ["LANG"]
width: full
---

[**plperl**](https://www.postgresql.org/docs/current/plperl.html) : PL/Perl procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3260** | {{< badge content="plperl" link="https://www.postgresql.org/docs/current/plperl.html" >}} | {{< ext "plperl" >}} | `1.0` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plperl" >}} |
|    **Need By**    | {{< ext "bool_plperl" >}} {{< ext "hstore_plperl" >}} {{< ext "jsonb_plperl" >}} {{< ext "plperl" >}} {{< ext "sparql" >}} |
|   **See Also**    | {{< ext "plperlu" >}} {{< ext "bool_plperlu" >}} {{< ext "jsonb_plperlu" >}} {{< ext "hstore_plperlu" >}} {{< ext "plpgsql" >}} {{< ext "pg_tle" >}} {{< ext "plv8" >}} {{< ext "pllua" >}} |
|    **Siblings**   | {{< ext "bool_plperl" >}} {{< ext "hstore_plperl" >}} {{< ext "jsonb_plperl" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plperl CASCADE; -- requires plperl
```



## Usage

> [plperl: PL/Perl trusted procedural language](https://www.postgresql.org/docs/current/plperl.html)

PL/Perl allows writing PostgreSQL functions in Perl. As a trusted language, it runs in a restricted environment without access to the filesystem or external modules.

```sql
CREATE EXTENSION plperl;

-- Simple scalar function
CREATE FUNCTION perl_hello(text) RETURNS text
LANGUAGE plperl AS $$
  my ($name) = @_;
  return "Hello, $name!";
$$;

SELECT perl_hello('world');

-- Using Perl regex for text processing
CREATE FUNCTION clean_whitespace(text) RETURNS text
LANGUAGE plperl AS $$
  my ($str) = @_;
  $str =~ s/^\s+|\s+$//g;   # trim
  $str =~ s/\s+/ /g;         # collapse internal whitespace
  return $str;
$$;

-- Returning a composite type
CREATE TYPE name_parts AS (first_name text, last_name text);

CREATE FUNCTION split_name(text) RETURNS name_parts
LANGUAGE plperl AS $$
  my ($full) = @_;
  my ($first, $last) = split(/\s+/, $full, 2);
  return { first_name => $first, last_name => $last };
$$;

-- Set-returning function
CREATE FUNCTION perl_generate_series(integer, integer) RETURNS SETOF integer
LANGUAGE plperl AS $$
  my ($start, $stop) = @_;
  for my $i ($start .. $stop) {
    return_next($i);
  }
  return undef;
$$;

-- Trigger function
CREATE FUNCTION perl_audit_trigger() RETURNS trigger
LANGUAGE plperl AS $$
  $_TD->{new}{modified_at} = localtime();
  return "MODIFY";
$$;
```

Database access uses `spi_exec_query`:

```sql
CREATE FUNCTION perl_row_count(text) RETURNS integer
LANGUAGE plperl AS $$
  my ($table) = @_;
  my $rv = spi_exec_query("SELECT count(*) AS cnt FROM $table");
  return $rv->{rows}[0]{cnt};
$$;
```
