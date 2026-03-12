---
title: "intarray"
linkTitle: "intarray"
description: "functions, operators, and index support for 1-D arrays of integers"
weight: 4960
categories: ["FUNC"]
width: full
---

[**intarray**](https://www.postgresql.org/docs/current/intarray.html) : functions, operators, and index support for 1-D arrays of integers


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4960** | {{< badge content="intarray" link="https://www.postgresql.org/docs/current/intarray.html" >}} | {{< ext "intarray" >}} | `1.5` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "aggs_for_arrays" >}} {{< ext "aggs_for_vecs" >}} {{< ext "arraymath" >}} {{< ext "floatvec" >}} {{< ext "vector" >}} {{< ext "vchord" >}} {{< ext "vectorscale" >}} {{< ext "vectorize" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.5" "PostgreSQL 18: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 17: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 16: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 15: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 14: version 1.5" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION intarray;
```



## Usage

> [intarray: integer array functions and operators with index support](https://www.postgresql.org/docs/current/intarray.html)

Provides functions and operators for manipulating null-free integer arrays, with GiST and GIN index support for fast array searches.

```sql
CREATE EXTENSION intarray;
```

### Functions

| Function | Description | Example |
|---|---|---|
| `icount(int[])` | Number of elements | `icount('{1,2,3}')` -- 3 |
| `sort(int[], dir)` | Sort array (`'asc'` or `'desc'`) | `sort('{3,1,2}','asc')` -- `{1,2,3}` |
| `sort_asc(int[])` | Sort ascending | `sort_asc('{3,1,2}')` -- `{1,2,3}` |
| `sort_desc(int[])` | Sort descending | `sort_desc('{3,1,2}')` -- `{3,2,1}` |
| `uniq(int[])` | Remove adjacent duplicates | `uniq(sort('{1,2,3,2,1}'))` -- `{1,2,3}` |
| `idx(int[], item)` | Index of first match | `idx('{11,22,33}', 22)` -- 2 |
| `subarray(int[], start, len)` | Extract sub-array | `subarray('{1,2,3,4}', 2, 2)` -- `{2,3}` |
| `intset(int)` | Make single-element array | `intset(42)` -- `{42}` |

### Operators

| Operator | Description |
|---|---|
| `&&` | Arrays overlap (have common elements) |
| `@>` | Left array contains right |
| `<@` | Left array is contained in right |
| `#` | Number of elements |
| `+` | Array concatenation / append element |
| `-` | Remove elements |
| `\|` | Union of arrays |
| `&` | Intersection of arrays |
| `@@` | Array matches a query expression |
| `~~` | Query expression matches array |

### Index Support

```sql
-- GiST index for array containment/overlap queries
CREATE INDEX idx ON messages USING GIST (tags gist__intbig_ops);

-- GIN index (alternative)
CREATE INDEX idx ON messages USING GIN (tags gin__int_ops);

-- Query with index support
SELECT * FROM messages WHERE tags && '{1,2}';     -- overlap
SELECT * FROM messages WHERE tags @> '{1,2}';     -- contains
SELECT * FROM messages WHERE tags @@ '1&(2|3)';  -- query expression
```
