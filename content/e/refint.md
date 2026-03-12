---
title: "refint"
linkTitle: "refint"
description: "functions for implementing referential integrity (obsolete)"
weight: 4880
categories: ["FUNC"]
width: full
---

[**refint**](https://www.postgresql.org/docs/current/contrib-spi.html#CONTRIB-SPI-REFINT) : functions for implementing referential integrity (obsolete)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4880** | {{< badge content="refint" link="https://www.postgresql.org/docs/current/contrib-spi.html#CONTRIB-SPI-REFINT" >}} | {{< ext "refint" >}} | `1.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


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
CREATE EXTENSION refint;
```



## Usage

> [refint: referential integrity trigger functions](https://www.postgresql.org/docs/current/contrib-spi.html)

Provides trigger functions for implementing referential integrity checks (largely superseded by built-in foreign key constraints).

```sql
CREATE EXTENSION refint;
```

### Trigger Functions

| Function | Description |
|---|---|
| `check_primary_key()` | Check referencing table -- ensures referenced row exists |
| `check_foreign_key()` | Check referenced table -- handles cascading actions on delete/update |

### check_primary_key

Trigger on the referencing table (`AFTER INSERT OR UPDATE`). Parameters: referencing column name(s), referenced table name, referenced column name(s).

```sql
CREATE TRIGGER check_fk
  AFTER INSERT OR UPDATE ON orders
  FOR EACH ROW
  EXECUTE FUNCTION check_primary_key('customer_id', 'customers', 'id');
```

### check_foreign_key

Trigger on the referenced table (`AFTER DELETE OR UPDATE`). Parameters: number of referencing tables, action (`cascade`, `restrict`, or `setnull`), primary key column(s), then referencing table and column pairs.

```sql
CREATE TRIGGER check_ref
  AFTER DELETE OR UPDATE ON customers
  FOR EACH ROW
  EXECUTE FUNCTION check_foreign_key(1, 'cascade', 'id', 'orders', 'customer_id');
```
