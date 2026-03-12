---
title: "pageinspect"
linkTitle: "pageinspect"
description: "inspect the contents of database pages at a low level"
weight: 6900
categories: ["STAT"]
width: full
---

[**pageinspect**](https://www.postgresql.org/docs/current/pageinspect.html) : inspect the contents of database pages at a low level


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6900** | {{< badge content="pageinspect" link="https://www.postgresql.org/docs/current/pageinspect.html" >}} | {{< ext "pageinspect" >}} | `1.12` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "amcheck" >}} {{< ext "pagevis" >}} {{< ext "pg_visibility" >}} {{< ext "pg_repack" >}} {{< ext "pg_squeeze" >}} {{< ext "pg_dirtyread" >}} {{< ext "pgdd" >}} {{< ext "pg_orphaned" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.12" "PostgreSQL 18: version 1.12" "green" >}} | {{< bg "1.12" "PostgreSQL 17: version 1.12" "green" >}} | {{< bg "1.12" "PostgreSQL 16: version 1.12" "green" >}} | {{< bg "1.12" "PostgreSQL 15: version 1.12" "green" >}} | {{< bg "1.12" "PostgreSQL 14: version 1.12" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pageinspect;
```



## Usage

> [pageinspect: low-level page inspection functions](https://www.postgresql.org/docs/current/pageinspect.html)

pageinspect provides functions to inspect the contents of database pages at a low level. Useful for debugging and educational purposes. Superuser only.

### General Functions

```sql
-- Read a raw page (main fork)
SELECT get_raw_page('my_table', 0);

-- Read from a specific fork (main, fsm, vm, init)
SELECT get_raw_page('my_table', 'fsm', 0);

-- Page header information
SELECT * FROM page_header(get_raw_page('my_table', 0));
-- Returns: lsn, checksum, flags, lower, upper, special, pagesize, version, prune_xid

-- Compute page checksum
SELECT page_checksum(get_raw_page('my_table', 0), 0);
```

### Heap Functions

```sql
-- Line pointers and tuple data on a heap page
SELECT * FROM heap_page_items(get_raw_page('my_table', 0));

-- Tuple data split into attributes
SELECT * FROM heap_page_item_attrs(get_raw_page('my_table', 0), 'my_table'::regclass);

-- Decode tuple infomask flags
SELECT t_ctid, raw_flags, combined_flags
FROM heap_page_items(get_raw_page('my_table', 0)),
     LATERAL heap_tuple_infomask_flags(t_infomask, t_infomask2)
WHERE t_infomask IS NOT NULL;
```

### B-Tree Index Functions

```sql
-- Index metapage
SELECT * FROM bt_metap('my_index');

-- Page-level statistics
SELECT * FROM bt_page_stats('my_index', 1);

-- Multi-page statistics
SELECT * FROM bt_multi_page_stats('my_index', 1, 10);

-- Page items (index entries)
SELECT itemoffset, ctid, itemlen, data FROM bt_page_items('my_index', 1);
```

### BRIN Index Functions

```sql
SELECT brin_page_type(get_raw_page('brin_idx', 0));
SELECT * FROM brin_metapage_info(get_raw_page('brin_idx', 0));
SELECT * FROM brin_revmap_data(get_raw_page('brin_idx', 2));
SELECT * FROM brin_page_items(get_raw_page('brin_idx', 5), 'brin_idx');
```

### GIN Index Functions

```sql
SELECT * FROM gin_metapage_info(get_raw_page('gin_idx', 0));
SELECT * FROM gin_page_opaque_info(get_raw_page('gin_idx', 2));
SELECT * FROM gin_leafpage_items(get_raw_page('gin_idx', 2));
```

### GiST Index Functions

```sql
SELECT * FROM gist_page_opaque_info(get_raw_page('gist_idx', 2));
SELECT * FROM gist_page_items(get_raw_page('gist_idx', 0), 'gist_idx');
SELECT * FROM gist_page_items_bytea(get_raw_page('gist_idx', 0));
```

### Hash Index Functions

```sql
SELECT hash_page_type(get_raw_page('hash_idx', 0));
SELECT * FROM hash_page_stats(get_raw_page('hash_idx', 1));
SELECT * FROM hash_page_items(get_raw_page('hash_idx', 1));
SELECT * FROM hash_bitmap_info('hash_idx', 2052);
SELECT * FROM hash_metapage_info(get_raw_page('hash_idx', 0));
```
