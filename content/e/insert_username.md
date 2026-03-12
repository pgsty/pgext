---
title: "insert_username"
linkTitle: "insert_username"
description: "functions for tracking who changed a table"
weight: 4882
categories: ["FUNC"]
width: full
---

[**insert_username**](https://www.postgresql.org/docs/current/contrib-spi.html#CONTRIB-SPI-INSERT-USERNAME) : functions for tracking who changed a table


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4882** | {{< badge content="insert_username" link="https://www.postgresql.org/docs/current/contrib-spi.html#CONTRIB-SPI-INSERT-USERNAME" >}} | {{< ext "insert_username" >}} | `1.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


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
CREATE EXTENSION insert_username;
```



## Usage

> [insert_username: track who modified a table row](https://www.postgresql.org/docs/current/contrib-spi.html)

Provides a trigger function that stores the current user's name into a specified text column.

```sql
CREATE EXTENSION insert_username;
```

### Trigger Function

| Function | Description |
|---|---|
| `insert_username()` | Store current username in the specified column |

Parameter: name of the text column to store the username.

### Examples

```sql
CREATE TABLE audit_log (
  id serial PRIMARY KEY,
  data text,
  modified_by text
);

-- Track who inserts
CREATE TRIGGER set_insert_user
  BEFORE INSERT ON audit_log
  FOR EACH ROW
  EXECUTE FUNCTION insert_username('modified_by');

-- Track who updates
CREATE TRIGGER set_update_user
  BEFORE UPDATE ON audit_log
  FOR EACH ROW
  EXECUTE FUNCTION insert_username('modified_by');

INSERT INTO audit_log (data) VALUES ('test');
SELECT modified_by FROM audit_log;  -- returns current user
```
