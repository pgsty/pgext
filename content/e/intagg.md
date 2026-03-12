---
title: "intagg"
linkTitle: "intagg"
description: "integer aggregator and enumerator (obsolete)"
weight: 4970
categories: ["FUNC"]
width: full
---

[**intagg**](https://www.postgresql.org/docs/current/intagg.html) : integer aggregator and enumerator (obsolete)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4970** | {{< badge content="intagg" link="https://www.postgresql.org/docs/current/intagg.html" >}} | {{< ext "intagg" >}} | `1.1` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "topn" >}} {{< ext "quantile" >}} {{< ext "lower_quantile" >}} {{< ext "count_distinct" >}} {{< ext "omnisketch" >}} {{< ext "ddsketch" >}} {{< ext "tdigest" >}} {{< ext "first_last_agg" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.1" "PostgreSQL 18: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 17: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 16: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 15: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 14: version 1.1" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION intagg;
```



## Usage

> [intagg: integer aggregator and enumerator (obsolete)](https://www.postgresql.org/docs/current/intagg.html)

Provides an integer aggregator and enumerator. These are now wrappers around the built-in `array_agg()` and `unnest()` functions.

```sql
CREATE EXTENSION intagg;
```

### Functions

| Function | Description |
|---|---|
| `int_array_aggregate(integer)` | Aggregate integers into an array (wrapper for `array_agg()`) |
| `int_array_enum(integer[])` | Expand array into rows (wrapper for `unnest()`) |

### Examples

```sql
-- Aggregate integers into an array
SELECT id_left, int_array_aggregate(id_right) AS rights
FROM many_to_many
GROUP BY id_left;

-- Expand an integer array into rows
SELECT int_array_enum(ARRAY[1, 2, 3, 4]);
-- Returns: 1, 2, 3, 4 (as separate rows)
```

Note: This module is obsolete. Use the built-in `array_agg()` and `unnest()` functions instead for new code.
