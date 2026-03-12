---
title: "cube"
linkTitle: "cube"
description: "data type for multidimensional cubes"
weight: 3950
categories: ["TYPE"]
width: full
---

[**cube**](https://www.postgresql.org/docs/current/cube.html) : data type for multidimensional cubes


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3950** | {{< badge content="cube" link="https://www.postgresql.org/docs/current/cube.html" >}} | {{< ext "cube" >}} | `1.5` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "earthdistance" >}} {{< ext "pg4ml" >}} |
|   **See Also**    | {{< ext "seg" >}} {{< ext "intarray" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.5" "PostgreSQL 18: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 17: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 16: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 15: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 14: version 1.5" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION cube;
```



## Usage

> [cube: multi-dimensional cube data type](https://www.postgresql.org/docs/current/cube.html)

The `cube` extension provides a data type for representing multi-dimensional cubes (up to 100 dimensions).

```sql
CREATE EXTENSION cube;
```

### Input Formats

```sql
SELECT '(1,2),(3,4)'::cube;         -- 2D cube
SELECT '(1,2,3)'::cube;              -- 3D point (zero-volume)
SELECT cube(1.0, 2.0);               -- 1D cube from 1 to 2
SELECT cube(ARRAY[1,2], ARRAY[3,4]); -- 2D cube from arrays
```

### Operators

| Operator | Description |
|----------|-------------|
| `&&` | Overlaps |
| `@>` | Contains |
| `<@` | Contained in |
| `->` | Extract n-th coordinate |
| `<->` | Euclidean distance |
| `<#>` | Taxicab (L-1) distance |
| `<=>` | Chebyshev (L-inf) distance |

### Functions

```sql
SELECT cube_dim('(1,2),(3,4)'::cube);                  -- 2
SELECT cube_ll_coord('(1,2),(3,4)'::cube, 1);          -- 1
SELECT cube_ur_coord('(1,2),(3,4)'::cube, 1);          -- 3
SELECT cube_is_point(cube(1,1));                        -- true
SELECT cube_distance('(1,2)'::cube, '(3,4)'::cube);    -- 2.828...
SELECT cube_union('(1,2)'::cube, '(3,4)'::cube);       -- (1,2),(3,4)
SELECT cube_inter('(0,0),(2,2)'::cube, '(1,1),(3,3)'::cube);
SELECT cube_subset(cube('(1,3,5),(6,7,8)'), ARRAY[2]); -- (3),(7)
SELECT cube_enlarge('(1,2),(3,4)'::cube, 0.5, 2);      -- expand by radius
```

### GiST Index with KNN Search

```sql
CREATE INDEX idx ON points USING gist (location);

-- Find nearest point to (0.5, 0.5, 0.5)
SELECT * FROM points
ORDER BY location <-> cube(ARRAY[0.5, 0.5, 0.5])
LIMIT 1;
```

### Dimension Handling

When operating on cubes of different dimensions, the lower-dimensional cube is treated as having zeroes for missing coordinates.
