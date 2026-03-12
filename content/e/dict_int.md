---
title: "dict_int"
linkTitle: "dict_int"
description: "text search dictionary template for integers"
weight: 4980
categories: ["FUNC"]
width: full
---

[**dict_int**](https://www.postgresql.org/docs/current/dict-int.html) : text search dictionary template for integers


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4980** | {{< badge content="dict_int" link="https://www.postgresql.org/docs/current/dict-int.html" >}} | {{< ext "dict_int" >}} | `1.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "dict_xsyn" >}} {{< ext "unaccent" >}} {{< ext "pg_similarity" >}} {{< ext "smlar" >}} {{< ext "pg_summarize" >}} {{< ext "pg_search" >}} {{< ext "pgroonga" >}} {{< ext "pg_bigm" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION dict_int;
```



## Usage

> [dict_int: integer text search dictionary](https://www.postgresql.org/docs/current/dict-int.html)

Controls the indexing of integers in full-text search to prevent excessive unique-word growth that degrades search performance.

```sql
CREATE EXTENSION dict_int;
```

### Configuration Parameters

| Parameter | Description | Default |
|---|---|---|
| `maxlen` | Maximum number of digits allowed | 6 |
| `rejectlong` | If `true`, reject overlength integers (stop word). If `false`, truncate. | `false` |
| `absval` | If `true`, strip leading `+`/`-` before applying maxlen | `false` |

### Examples

```sql
-- Test the default dictionary
SELECT ts_lexize('intdict', '12345678');
-- {123456}  (truncated to 6 digits by default)

-- Configure to reject long integers
ALTER TEXT SEARCH DICTIONARY intdict (MAXLEN = 4, REJECTLONG = true);
SELECT ts_lexize('intdict', '12345');
-- {}  (rejected as stop word)

SELECT ts_lexize('intdict', '1234');
-- {1234}  (accepted)

-- Apply to a text search configuration
ALTER TEXT SEARCH CONFIGURATION english
  ALTER MAPPING FOR int, uint WITH intdict;
```
