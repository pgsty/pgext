---
title: "bloom"
linkTitle: "bloom"
description: "bloom access method - signature file based index"
weight: 2990
categories: ["FEAT"]
width: full
---

[**bloom**](https://www.postgresql.org/docs/current/bloom.html) : bloom access method - signature file based index


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2990** | {{< badge content="bloom" link="https://www.postgresql.org/docs/current/bloom.html" >}} | {{< ext "bloom" >}} | `1.0` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "hll" >}} {{< ext "age" >}} {{< ext "rum" >}} {{< ext "pg_graphql" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_hint_plan" >}} {{< ext "hypopg" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION bloom;
```




## Usage

> [bloom: bloom access method - signature file based index](https://www.postgresql.org/docs/current/bloom.html)

The `bloom` extension provides an index access method based on [Bloom filters](https://en.wikipedia.org/wiki/Bloom_filter). A Bloom filter is a space-efficient data structure that tests whether an element is a member of a set, with possible false positives but no false negatives.

Bloom indexes are particularly useful for tables with many columns where queries test arbitrary combinations of columns. A single bloom index can replace multiple btree indexes while using significantly less space.

### Create Bloom Index

```sql
CREATE INDEX bloomidx ON tbloom USING bloom (i1, i2, i3, i4, i5, i6);
```

With custom parameters:

```sql
CREATE INDEX bloomidx ON tbloom USING bloom (i1, i2, i3)
       WITH (length=80, col1=2, col2=2, col3=4);
```

### Parameters

| Parameter | Default | Max | Description |
|-----------|---------|-----|-------------|
| `length` | 80 | 4096 | Length of each signature in bits (rounded to nearest 16) |
| `col1` - `col32` | 2 | 4095 | Number of bits generated for each index column |

### Operator Class Support

Bloom indexes only support the equality operator (`=`). Built-in operator classes exist for `int4` and `text`. You can define custom operator classes for other types that have a hash function:

```sql
CREATE OPERATOR CLASS text_ops
DEFAULT FOR TYPE text USING bloom AS
    OPERATOR    1   =(text, text),
    FUNCTION    1   hashtext(text);
```

### Example

```sql
-- Create a table with many columns
CREATE TABLE tbloom AS
   SELECT
     (random() * 1000000)::int as i1,
     (random() * 1000000)::int as i2,
     (random() * 1000000)::int as i3,
     (random() * 1000000)::int as i4,
     (random() * 1000000)::int as i5,
     (random() * 1000000)::int as i6
   FROM generate_series(1, 10000000);

-- A single bloom index covers all column combinations
CREATE INDEX bloomidx ON tbloom USING bloom (i1, i2, i3, i4, i5, i6);

-- Queries on any subset of columns can use the index
SELECT * FROM tbloom WHERE i2 = 898732 AND i5 = 123451;
```

### Limitations

- Only equality (`=`) queries are supported (no range queries)
- Does not support `UNIQUE` indexes
- Does not support searching for `NULL` values
- Results require recheck against heap due to false positives
