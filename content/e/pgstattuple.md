---
title: "pgstattuple"
linkTitle: "pgstattuple"
description: "show tuple-level statistics"
weight: 6970
categories: ["STAT"]
width: full
---

[**pgstattuple**](https://www.postgresql.org/docs/current/pgstattuple.html) : show tuple-level statistics


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6970** | {{< badge content="pgstattuple" link="https://www.postgresql.org/docs/current/pgstattuple.html" >}} | {{< ext "pgstattuple" >}} | `1.5` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pageinspect" >}} {{< ext "pg_freespacemap" >}} {{< ext "pg_visibility" >}} {{< ext "pg_rewrite" >}} {{< ext "pg_checksums" >}} {{< ext "pg_catcheck" >}} {{< ext "amcheck" >}} {{< ext "toastinfo" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.5" "PostgreSQL 18: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 17: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 16: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 15: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 14: version 1.5" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgstattuple;
```



## Usage

> [pgstattuple: tuple-level statistics](https://www.postgresql.org/docs/current/pgstattuple.html)

pgstattuple provides functions to obtain tuple-level statistics for tables and indexes, including dead tuple counts and free space metrics.

### Table Statistics

```sql
SELECT * FROM pgstattuple('my_table');

 table_len          | 81920
 tuple_count        | 1000
 tuple_len          | 72000
 tuple_percent      | 87.89
 dead_tuple_count   | 50
 dead_tuple_len     | 3600
 dead_tuple_percent | 4.39
 free_space         | 6320
 free_percent       | 7.71
```

**Approximate (faster) version** using visibility map and FSM:

```sql
SELECT * FROM pgstattuple_approx('my_table');
-- Returns: table_len, scanned_percent, approx_tuple_count, approx_tuple_len,
--          dead_tuple_count, dead_tuple_len, approx_free_space, ...
```

### B-Tree Index Statistics

```sql
SELECT * FROM pgstatindex('my_index');

 version            | 4
 tree_level         | 2
 index_size         | 1327104
 root_block_no      | 3
 internal_pages     | 1
 leaf_pages         | 160
 empty_pages        | 0
 deleted_pages      | 0
 avg_leaf_density   | 89.27
 leaf_fragmentation | 0
```

### GIN Index Statistics

```sql
SELECT * FROM pgstatginindex('my_gin_index');
-- Returns: version, pending_pages, pending_tuples
```

### Hash Index Statistics

```sql
SELECT * FROM pgstathashindex('my_hash_index');
-- Returns: version, bucket_pages, overflow_pages, bitmap_pages,
--          unused_pages, live_items, dead_tuples, free_percent
```

### Page Count

```sql
SELECT pg_relpages('my_table');
```

### Access

Restricted to superusers and roles with `pg_stat_scan_tables` privileges.
