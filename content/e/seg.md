---
title: "seg"
linkTitle: "seg"
description: "data type for representing line segments or floating-point intervals"
weight: 3940
categories: ["TYPE"]
width: full
---

[**seg**](https://www.postgresql.org/docs/current/seg.html) : data type for representing line segments or floating-point intervals


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3940** | {{< badge content="seg" link="https://www.postgresql.org/docs/current/seg.html" >}} | {{< ext "seg" >}} | `1.4` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "cube" >}} {{< ext "intarray" >}} {{< ext "intagg" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.4" "PostgreSQL 18: version 1.4" "green" >}} | {{< bg "1.4" "PostgreSQL 17: version 1.4" "green" >}} | {{< bg "1.4" "PostgreSQL 16: version 1.4" "green" >}} | {{< bg "1.4" "PostgreSQL 15: version 1.4" "green" >}} | {{< bg "1.4" "PostgreSQL 14: version 1.4" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION seg;
```



## Usage

> [seg: line segment / floating-point interval data type](https://www.postgresql.org/docs/current/seg.html)

The `seg` extension provides a data type for representing line segments or floating-point intervals, useful for representing measurements with uncertainty.

```sql
CREATE EXTENSION seg;
```

### Input Formats

```sql
SELECT '5.0'::seg;           -- zero-length interval (point)
SELECT '5.0 .. 7.0'::seg;    -- interval from 5.0 to 7.0
SELECT '5(+-)0.3'::seg;      -- interval 4.7 .. 5.3
SELECT '50 ..'::seg;          -- open interval >= 50
SELECT '.. 0'::seg;           -- open interval <= 0
```

Certainty indicators (`<`, `>`, `~`) can be prepended but are treated as comments and ignored by operators.

### Operators

**Spatial operators (GiST-indexed):**

| Operator | Description |
|----------|-------------|
| `<<` | Entirely left of |
| `>>` | Entirely right of |
| `&<` | Does not extend right of |
| `&>` | Does not extend left of |
| `=` | Equal |
| `&&` | Overlaps |
| `@>` | Contains |
| `<@` | Contained in |

**Comparison operators** (`<`, `<=`, `>`, `>=`) are available for sorting.

### Index Support

```sql
CREATE INDEX idx ON measurements USING gist (reading);
```

### Precision

Values are stored as 32-bit floating-point pairs, retaining up to 7 significant digits. Trailing zeroes are preserved.
