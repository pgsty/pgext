---
title: "dict_xsyn"
linkTitle: "dict_xsyn"
description: "text search dictionary template for extended synonym processing"
weight: 4900
categories: ["FUNC"]
width: full
---

[**dict_xsyn**](https://www.postgresql.org/docs/current/dict-xsyn.html) : text search dictionary template for extended synonym processing


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4900** | {{< badge content="dict_xsyn" link="https://www.postgresql.org/docs/current/dict-xsyn.html" >}} | {{< ext "dict_xsyn" >}} | `1.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "dict_int" >}} {{< ext "unaccent" >}} {{< ext "pg_similarity" >}} {{< ext "smlar" >}} {{< ext "pg_summarize" >}} {{< ext "pg_search" >}} {{< ext "pgroonga" >}} {{< ext "pg_bigm" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION dict_xsyn;
```



## Usage

> [dict_xsyn: extended synonym dictionary for text search](https://www.postgresql.org/docs/current/dict-xsyn.html)

Provides an extended synonym dictionary template for text search, replacing words with groups of synonyms.

```sql
CREATE EXTENSION dict_xsyn;
```

### Configuration Parameters

| Parameter | Description | Default |
|---|---|---|
| `matchorig` | Accept original word | `true` |
| `matchsynonyms` | Accept synonyms as input | `false` |
| `keeporig` | Include original in output | `true` |
| `keepsynonyms` | Include synonyms in output | `true` |
| `rules` | Base name of synonym file in `$SHAREDIR/tsearch_data/` (`.rules` extension) | -- |

### Rules File Format

```
word syn1 syn2 syn3
```

Lines starting with `#` are comments.

### Examples

```sql
-- Configure the dictionary
ALTER TEXT SEARCH DICTIONARY xsyn (RULES='my_rules', KEEPORIG=true);

-- Test the dictionary
SELECT ts_lexize('xsyn', 'word');
-- {word,syn1,syn2,syn3}

-- Match synonyms as input too
ALTER TEXT SEARCH DICTIONARY xsyn (RULES='my_rules', MATCHSYNONYMS=true);
SELECT ts_lexize('xsyn', 'syn1');
-- {syn1,syn2,syn3}

-- Use in a text search configuration
ALTER TEXT SEARCH CONFIGURATION english
  ALTER MAPPING FOR word, asciiword WITH xsyn, english_stem;
```
