---
title: "pg_surgery"
linkTitle: "pg_surgery"
description: "extension to perform surgery on a damaged relation"
weight: 5990
categories: ["ADMIN"]
width: full
---

[**pg_surgery**](https://www.postgresql.org/docs/current/pgsurgery.html) : extension to perform surgery on a damaged relation


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5990** | {{< badge content="pg_surgery" link="https://www.postgresql.org/docs/current/pgsurgery.html" >}} | {{< ext "pg_surgery" >}} | `1.0` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_dirtyread" >}} {{< ext "amcheck" >}} {{< ext "pageinspect" >}} {{< ext "pg_checksums" >}} {{< ext "pg_catcheck" >}} {{< ext "pg_cheat_funcs" >}} {{< ext "pagevis" >}} {{< ext "pg_visibility" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_surgery;
```




## Usage

> [pg_surgery: extension to perform surgery on a damaged relation](https://www.postgresql.org/docs/current/pgsurgery.html)

The `pg_surgery` extension provides functions to perform low-level surgery on damaged relations. These are last-resort tools that can corrupt data if misused.

### Functions

#### heap_force_kill

Marks line pointers as "dead" without examining the tuples, forcibly removing inaccessible tuples.

```sql
heap_force_kill(regclass, tid[]) RETURNS void
```

```sql
-- Tuple causes error when accessed
SELECT * FROM t1 WHERE ctid = '(0, 1)';
-- ERROR: could not access status of transaction 4007513275

-- Force kill the problematic tuple
SELECT heap_force_kill('t1'::regclass, ARRAY['(0, 1)']::tid[]);

-- Tuple is now removed
SELECT * FROM t1 WHERE ctid = '(0, 1)';
-- (0 rows)
```

#### heap_force_freeze

Marks tuples as frozen without examining tuple data, making tuples accessible when visibility information is corrupted.

```sql
heap_force_freeze(regclass, tid[]) RETURNS void
```

```sql
-- VACUUM fails on corrupted visibility info
VACUUM t1;
-- ERROR: found xmin 507 from before relfrozenxid 515

-- Find the problematic tuple
SELECT ctid FROM t1 WHERE xmin = 507;
--  ctid
-- -------
--  (0,3)

-- Force freeze the tuple
SELECT heap_force_freeze('t1'::regclass, ARRAY['(0, 3)']::tid[]);

-- Tuple is now frozen (xmin becomes 2 = FrozenTransactionId)
SELECT ctid FROM t1 WHERE xmin = 2;
--  ctid
-- -------
--  (0,3)
```

### When to Use

- `heap_force_kill`: When tuples cause errors on access due to corrupted transaction status, and the data can be discarded
- `heap_force_freeze`: When VACUUM fails due to tuples with xmin before relfrozenxid, or when tuples are invisible due to corrupted visibility information

These functions are unsafe by design and should only be used as a last resort when normal recovery methods have failed.
