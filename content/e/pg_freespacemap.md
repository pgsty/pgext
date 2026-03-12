---
title: "pg_freespacemap"
linkTitle: "pg_freespacemap"
description: "examine the free space map (FSM)"
weight: 6950
categories: ["STAT"]
width: full
---

[**pg_freespacemap**](https://www.postgresql.org/docs/current/pgfreespacemap.html) : examine the free space map (FSM)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6950** | {{< badge content="pg_freespacemap" link="https://www.postgresql.org/docs/current/pgfreespacemap.html" >}} | {{< ext "pg_freespacemap" >}} | `1.2` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_relusage" >}} {{< ext "pg_visibility" >}} {{< ext "pgstattuple" >}} {{< ext "amcheck" >}} {{< ext "toastinfo" >}} {{< ext "pageinspect" >}} {{< ext "pg_repack" >}} {{< ext "pg_squeeze" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.2" "PostgreSQL 18: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 17: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 16: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 15: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 14: version 1.2" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_freespacemap;
```



## Usage

> [pg_freespacemap: examine the free space map](https://www.postgresql.org/docs/current/pgfreespacemap.html)

pg_freespacemap provides functions to examine the free space map (FSM), which tracks available space on each page of a relation.

### Functions

**Free space on a single page:**

```sql
SELECT pg_freespace('my_table'::regclass, 0);  -- block 0
```

**Free space on all pages:**

```sql
SELECT * FROM pg_freespace('my_table'::regclass);

 blkno | avail
-------+-------
     0 |     0
     1 |     0
     2 |   224
     3 |  3456
     4 |  8160
```

### Example: Table Bloat Analysis

```sql
-- Pages with significant free space
SELECT blkno, avail
FROM pg_freespace('my_table'::regclass)
WHERE avail > 1000
ORDER BY avail DESC;

-- Total free space in a relation
SELECT sum(avail) AS total_free_bytes,
       count(*) AS total_pages,
       count(*) FILTER (WHERE avail > 0) AS pages_with_free_space
FROM pg_freespace('my_table'::regclass);
```

### Notes

- FSM values are rounded to 1/256th of `BLCKSZ` (typically 32 bytes)
- FSM is not kept fully up-to-date; values may lag behind actual free space
- For indexes, only entirely unused pages are tracked
- Access restricted to superusers and `pg_stat_scan_tables` members
