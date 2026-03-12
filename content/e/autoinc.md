---
title: "autoinc"
linkTitle: "autoinc"
description: "functions for autoincrementing fields"
weight: 4881
categories: ["FUNC"]
width: full
---

[**autoinc**](https://www.postgresql.org/docs/current/contrib-spi.html#CONTRIB-SPI-AUTOINC) : functions for autoincrementing fields


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4881** | {{< badge content="autoinc" link="https://www.postgresql.org/docs/current/contrib-spi.html#CONTRIB-SPI-AUTOINC" >}} | {{< ext "autoinc" >}} | `1.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "permuteseq" >}} {{< ext "pg_hashids" >}} {{< ext "sequential_uuids" >}} {{< ext "topn" >}} {{< ext "quantile" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION autoinc;
```



## Usage

> [autoinc: auto-increment trigger function](https://www.postgresql.org/docs/current/contrib-spi.html)

Provides a trigger function that auto-increments specified columns using sequences.

```sql
CREATE EXTENSION autoinc;
```

### Trigger Function

| Function | Description |
|---|---|
| `autoinc()` | Auto-increment columns from sequences on insert (and optionally update) |

Parameters are pairs of (column name, sequence name). The value is replaced only if initially zero or NULL.

### Examples

```sql
CREATE SEQUENCE next_id;
CREATE TABLE ids (id int4, idesc text);

CREATE TRIGGER ids_autoinc
  BEFORE INSERT ON ids
  FOR EACH ROW
  EXECUTE FUNCTION autoinc('id', 'next_id');

INSERT INTO ids (idesc) VALUES ('first');  -- id is auto-assigned from next_id

-- Multiple auto-increment columns
CREATE TRIGGER multi_autoinc
  BEFORE INSERT ON my_table
  FOR EACH ROW
  EXECUTE FUNCTION autoinc('col1', 'seq1', 'col2', 'seq2');
```
