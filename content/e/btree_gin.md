---
title: "btree_gin"
linkTitle: "btree_gin"
description: "support for indexing common datatypes in GIN"
weight: 4950
categories: ["FUNC"]
width: full
---

[**btree_gin**](https://www.postgresql.org/docs/current/btree-gin.html) : support for indexing common datatypes in GIN


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4950** | {{< badge content="btree_gin" link="https://www.postgresql.org/docs/current/btree-gin.html" >}} | {{< ext "btree_gin" >}} | `1.3` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "btree_gist" >}} {{< ext "unaccent" >}} {{< ext "fuzzystrmatch" >}} {{< ext "pg_trgm" >}} {{< ext "prefix" >}} {{< ext "citext" >}} {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.3" "PostgreSQL 18: version 1.3" "green" >}} | {{< bg "1.3" "PostgreSQL 17: version 1.3" "green" >}} | {{< bg "1.3" "PostgreSQL 16: version 1.3" "green" >}} | {{< bg "1.3" "PostgreSQL 15: version 1.3" "green" >}} | {{< bg "1.3" "PostgreSQL 14: version 1.3" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION btree_gin;
```



## Usage

> [btree_gin: B-tree equivalent GIN operator classes](https://www.postgresql.org/docs/current/btree-gin.html)

Provides GIN index operator classes for data types that normally only support B-tree indexing. Useful for multicolumn GIN indexes that combine GIN-indexable and B-tree-indexable columns.

```sql
CREATE EXTENSION btree_gin;
```

### Supported Data Types

`int2`, `int4`, `int8`, `float4`, `float8`, `numeric`, `timestamp with time zone`, `timestamp without time zone`, `time with time zone`, `time without time zone`, `date`, `interval`, `oid`, `money`, `char`, `varchar`, `text`, `bytea`, `macaddr`, `macaddr8`, `inet`, `cidr`, `uuid`, `bit`, `varbit`, `bool`, `name`, `bpchar`, and all `enum` types.

### Examples

```sql
-- GIN index on an integer column
CREATE INDEX idx ON test USING GIN (a);
SELECT * FROM test WHERE a < 10;

-- Multicolumn GIN index combining full-text search with a scalar filter
CREATE INDEX idx ON articles USING GIN (body_tsvector, category);
SELECT * FROM articles
WHERE body_tsvector @@ to_tsquery('PostgreSQL')
  AND category = 'tech';
```

Note: btree_gin does not outperform standard B-tree indexes for single-column queries. Its main benefit is combining scalar columns with GIN-native columns (like tsvector or arrays) in a single multicolumn index.
