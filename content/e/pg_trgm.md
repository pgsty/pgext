---
title: "pg_trgm"
linkTitle: "pg_trgm"
description: "text similarity measurement and index searching based on trigrams"
weight: 2390
categories: ["FTS"]
width: full
---

[**pg_trgm**](https://www.postgresql.org/docs/current/pgtrgm.html) : text similarity measurement and index searching based on trigrams


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2390** | {{< badge content="pg_trgm" link="https://www.postgresql.org/docs/current/pgtrgm.html" >}} | {{< ext "pg_trgm" >}} | `1.6` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_similarity" >}} {{< ext "pg_bigm" >}} {{< ext "fuzzystrmatch" >}} {{< ext "unaccent" >}} {{< ext "smlar" >}} {{< ext "pgroonga_database" >}} {{< ext "rum" >}} {{< ext "citext" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.6" "PostgreSQL 18: version 1.6" "green" >}} | {{< bg "1.6" "PostgreSQL 17: version 1.6" "green" >}} | {{< bg "1.6" "PostgreSQL 16: version 1.6" "green" >}} | {{< bg "1.6" "PostgreSQL 15: version 1.6" "green" >}} | {{< bg "1.6" "PostgreSQL 14: version 1.6" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_trgm;
```



## Usage

> [pg_trgm: Text similarity measurement and index searching based on trigrams](https://www.postgresql.org/docs/current/pgtrgm.html)

The `pg_trgm` module provides functions and operators for determining the similarity of alphanumeric text based on trigram matching, plus index operator classes for fast string similarity searches.

```sql
CREATE EXTENSION pg_trgm;
```

A trigram is a group of three consecutive characters from a string. Two strings are similar if they share many trigrams.

### Functions

| Function | Description |
|----------|-------------|
| `similarity(text, text)` → `real` | Returns similarity between 0 and 1 |
| `show_trgm(text)` → `text[]` | Returns array of all trigrams in the string |
| `word_similarity(text, text)` → `real` | Similarity of first string to most similar word in second |
| `strict_word_similarity(text, text)` → `real` | Similar but with stricter word boundary matching |
| `show_limit()` → `real` | *(Deprecated)* Returns `pg_trgm.similarity_threshold` |
| `set_limit(real)` → `real` | *(Deprecated)* Sets `pg_trgm.similarity_threshold` |

```sql
SELECT similarity('word', 'two words');
-- 0.36363637

SELECT show_trgm('word');
-- {"  w"," wo",ord,"rd ",wor}
```

### Operators

| Operator | Description |
|----------|-------------|
| `text % text` → `boolean` | True if similarity > `pg_trgm.similarity_threshold` |
| `text <% text` → `boolean` | True if word similarity > `pg_trgm.word_similarity_threshold` |
| `text %> text` → `boolean` | Commutator of `<%` |
| `text <<% text` → `boolean` | True if strict word similarity > threshold |
| `text %>> text` → `boolean` | Commutator of `<<%` |
| `text <-> text` → `real` | Distance (1 - similarity) |
| `text <<-> text` → `real` | Word distance (1 - word_similarity) |
| `text <->> text` → `real` | Commutator of `<<->` |
| `text <<<-> text` → `real` | Strict word distance |
| `text <->>> text` → `real` | Commutator of `<<<->` |

### GUC Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pg_trgm.similarity_threshold` | 0.3 | Threshold for `%` operator |
| `pg_trgm.word_similarity_threshold` | 0.6 | Threshold for `<%` and `%>` operators |
| `pg_trgm.strict_word_similarity_threshold` | 0.5 | Threshold for `<<%` and `%>>` operators |

### Index Support

GiST and GIN indexes support the similarity operators:

```sql
-- GIN index (faster lookups, slower builds)
CREATE INDEX trgm_idx ON test_trgm USING GIN (t gin_trgm_ops);

-- GiST index (supports distance operators for KNN)
CREATE INDEX trgm_idx ON test_trgm USING GIST (t gist_trgm_ops);

-- GiST with custom signature length
CREATE INDEX trgm_idx ON test_trgm USING GIST (t gist_trgm_ops(siglen=32));
```

### Text Search Example

Using trigram indexes to speed up `LIKE` / `ILIKE` / regex queries:

```sql
SELECT t, similarity(t, 'word') AS sml
FROM test_trgm
WHERE t % 'word'
ORDER BY sml DESC, t;

-- KNN search using distance operator
SELECT t, t <-> 'word' AS dist
FROM test_trgm
ORDER BY dist
LIMIT 10;
```

GIN and GiST trigram indexes also accelerate `LIKE`, `ILIKE`, `~`, and `~*` queries automatically.
