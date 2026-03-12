---
title: "unaccent"
linkTitle: "unaccent"
description: "text search dictionary that removes accents"
weight: 4990
categories: ["FUNC"]
width: full
---

[**unaccent**](https://www.postgresql.org/docs/current/unaccent.html) : text search dictionary that removes accents


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4990** | {{< badge content="unaccent" link="https://www.postgresql.org/docs/current/unaccent.html" >}} | {{< ext "unaccent" >}} | `1.1` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_trgm" >}} {{< ext "fuzzystrmatch" >}} {{< ext "citext" >}} {{< ext "btree_gist" >}} {{< ext "btree_gin" >}} {{< ext "prefix" >}} {{< ext "dict_xsyn" >}} {{< ext "dict_int" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.1" "PostgreSQL 18: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 17: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 16: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 15: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 14: version 1.1" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION unaccent;
```



## Usage

> [unaccent: text search dictionary for accent removal](https://www.postgresql.org/docs/current/unaccent.html)

Removes accents (diacritic signs) from text. Can be used as a text search dictionary or as a standalone function.

```sql
CREATE EXTENSION unaccent;
```

### Functions

| Function | Description |
|---|---|
| `unaccent(text)` | Remove accents using the default dictionary |
| `unaccent(dictionary regdictionary, text)` | Remove accents using a specific dictionary |

### Text Search Usage

```sql
-- Test the dictionary
SELECT ts_lexize('unaccent', 'Hôtel');
-- {Hotel}

-- Create an accent-insensitive French text search config
CREATE TEXT SEARCH CONFIGURATION fr (COPY = french);
ALTER TEXT SEARCH CONFIGURATION fr
  ALTER MAPPING FOR hword, hword_part, word
  WITH unaccent, french_stem;

SELECT to_tsvector('fr', 'Hôtels de la Mer');
-- 'hotel':1 'mer':4

-- Accent-insensitive search
SELECT to_tsvector('fr', 'Hôtel de la Mer') @@ to_tsquery('fr', 'Hotels');
-- true
```

### Standalone Function Usage

```sql
SELECT unaccent('Hôtel');           -- Hotel
SELECT unaccent('café');            -- cafe
SELECT unaccent('résumé');          -- resume
SELECT unaccent('niño');            -- nino
```
