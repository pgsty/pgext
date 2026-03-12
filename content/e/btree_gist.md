---
title: "btree_gist"
linkTitle: "btree_gist"
description: "support for indexing common datatypes in GiST"
weight: 4940
categories: ["FUNC"]
width: full
---

[**btree_gist**](https://www.postgresql.org/docs/current/btree-gist.html) : support for indexing common datatypes in GiST


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4940** | {{< badge content="btree_gist" link="https://www.postgresql.org/docs/current/btree-gist.html" >}} | {{< ext "btree_gist" >}} | `1.7` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "emaj" >}} {{< ext "omni_auth" >}} {{< ext "periods" >}} {{< ext "pgautofailover" >}} {{< ext "powa" >}} |
|   **See Also**    | {{< ext "btree_gin" >}} {{< ext "unaccent" >}} {{< ext "fuzzystrmatch" >}} {{< ext "pg_trgm" >}} {{< ext "prefix" >}} {{< ext "citext" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.7" "PostgreSQL 18: version 1.7" "green" >}} | {{< bg "1.7" "PostgreSQL 17: version 1.7" "green" >}} | {{< bg "1.7" "PostgreSQL 16: version 1.7" "green" >}} | {{< bg "1.7" "PostgreSQL 15: version 1.7" "green" >}} | {{< bg "1.7" "PostgreSQL 14: version 1.7" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION btree_gist;
```



## Usage

> [btree_gist: B-tree equivalent GiST operator classes](https://www.postgresql.org/docs/current/btree-gist.html)

Provides GiST index operator classes for data types that normally only support B-tree indexing. Enables exclusion constraints combining equality with range operators.

```sql
CREATE EXTENSION btree_gist;
```

### Supported Data Types

`int2`, `int4`, `int8`, `float4`, `float8`, `numeric`, `timestamp with time zone`, `timestamp without time zone`, `time with time zone`, `time without time zone`, `date`, `interval`, `oid`, `money`, `char`, `varchar`, `text`, `bytea`, `bit`, `varbit`, `macaddr`, `macaddr8`, `inet`, `cidr`, `uuid`, `bool`, and all `enum` types.

### Distance Operator

The `<->` operator is provided for nearest-neighbor searches on numeric and temporal types.

### Examples

```sql
-- GiST index on integer column
CREATE INDEX idx ON test USING GIST (a);
SELECT * FROM test WHERE a < 10;

-- Nearest-neighbor search
SELECT *, a <-> 42 AS dist FROM test ORDER BY a <-> 42 LIMIT 10;

-- Exclusion constraint: each cage can only contain one type of animal
CREATE TABLE zoo (
  cage   integer,
  animal text,
  EXCLUDE USING GIST (cage WITH =, animal WITH <>)
);

INSERT INTO zoo VALUES (1, 'lion');    -- OK
INSERT INTO zoo VALUES (1, 'tiger');   -- ERROR: conflicting key value
INSERT INTO zoo VALUES (2, 'tiger');   -- OK

-- Exclusion constraint for non-overlapping time ranges per room
CREATE TABLE reservations (
  room int,
  during tsrange,
  EXCLUDE USING GIST (room WITH =, during WITH &&)
);
```
