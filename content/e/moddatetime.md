---
title: "moddatetime"
linkTitle: "moddatetime"
description: "functions for tracking last modification time"
weight: 4883
categories: ["FUNC"]
width: full
---

[**moddatetime**](https://www.postgresql.org/docs/current/contrib-spi.html#CONTRIB-SPI-MODDATETIME) : functions for tracking last modification time


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4883** | {{< badge content="moddatetime" link="https://www.postgresql.org/docs/current/contrib-spi.html#CONTRIB-SPI-MODDATETIME" >}} | {{< ext "moddatetime" >}} | `1.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "topn" >}} {{< ext "quantile" >}} {{< ext "lower_quantile" >}} {{< ext "count_distinct" >}} {{< ext "omnisketch" >}} {{< ext "ddsketch" >}} {{< ext "tdigest" >}} {{< ext "first_last_agg" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION moddatetime;
```



## Usage

> [moddatetime: track modification timestamp](https://www.postgresql.org/docs/current/contrib-spi.html)

Provides a trigger function that stores the current timestamp when a row is modified.

```sql
CREATE EXTENSION moddatetime;
```

### Trigger Function

| Function | Description |
|---|---|
| `moddatetime()` | Store current timestamp in the specified column on UPDATE |

Parameter: name of the `timestamp` or `timestamp with time zone` column to update.

### Examples

```sql
CREATE TABLE documents (
  id serial PRIMARY KEY,
  content text,
  modified_at timestamp with time zone
);

CREATE TRIGGER set_modified
  BEFORE UPDATE ON documents
  FOR EACH ROW
  EXECUTE FUNCTION moddatetime('modified_at');

UPDATE documents SET content = 'new content' WHERE id = 1;
-- modified_at is automatically set to current timestamp
```
