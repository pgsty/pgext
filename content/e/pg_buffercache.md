---
title: "pg_buffercache"
linkTitle: "pg_buffercache"
description: "examine the shared buffer cache"
weight: 6930
categories: ["STAT"]
width: full
---

[**pg_buffercache**](https://www.postgresql.org/docs/current/pgbuffercache.html) : examine the shared buffer cache


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6930** | {{< badge content="pg_buffercache" link="https://www.postgresql.org/docs/current/pgbuffercache.html" >}} | {{< ext "pg_buffercache" >}} | `1.5` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_repack" >}} {{< ext "pgfincore" >}} {{< ext "pgcozy" >}} {{< ext "pg_prewarm" >}} {{< ext "pgmeminfo" >}} {{< ext "pg_squeeze" >}} {{< ext "old_snapshot" >}} {{< ext "system_stats" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.5" "PostgreSQL 18: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 17: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 16: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 15: version 1.5" "green" >}} | {{< bg "1.5" "PostgreSQL 14: version 1.5" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_buffercache;
```



## Usage

> [pg_buffercache: inspect the shared buffer cache](https://www.postgresql.org/docs/current/pgbuffercache.html)

pg_buffercache provides views and functions to examine what is stored in the PostgreSQL shared buffer cache in real time.

### Views

**`pg_buffercache`** -- detailed per-buffer information:

```sql
-- Top 10 relations by buffer usage
SELECT n.nspname, c.relname, count(*) AS buffers
FROM pg_buffercache b
JOIN pg_class c ON b.relfilenode = pg_relation_filenode(c.oid)
  AND b.reldatabase IN (0, (SELECT oid FROM pg_database WHERE datname = current_database()))
JOIN pg_namespace n ON n.oid = c.relnamespace
GROUP BY n.nspname, c.relname
ORDER BY 3 DESC
LIMIT 10;
```

Columns: `bufferid`, `relfilenode`, `reltablespace`, `reldatabase`, `relforknumber`, `relblocknumber`, `isdirty`, `usagecount`, `pinning_backends`.

### Summary Functions

```sql
-- Quick buffer cache summary (cheaper than the view)
SELECT * FROM pg_buffercache_summary();
--  buffers_used | buffers_unused | buffers_dirty | buffers_pinned | usagecount_avg

-- Buffer distribution by usage count
SELECT * FROM pg_buffercache_usage_counts();
--  usage_count | buffers | dirty | pinned
```

### Eviction Functions (Superuser, Developer Testing Only)

```sql
-- Evict a single buffer by ID
SELECT * FROM pg_buffercache_evict(42);

-- Evict all buffers for a relation
SELECT * FROM pg_buffercache_evict_relation('my_table'::regclass);

-- Evict all unpinned buffers
SELECT * FROM pg_buffercache_evict_all();
```

### NUMA View

```sql
-- NUMA node mapping for shared buffers
SELECT * FROM pg_buffercache_numa;
-- Returns: bufferid, os_page_num, numa_node
```

### Access

Restricted to superusers and roles with `pg_monitor` privileges.
