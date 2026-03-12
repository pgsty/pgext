---
title: "pg_logicalinspect"
linkTitle: "pg_logicalinspect"
description: "Logical decoding components inspection"
weight: 6890
categories: ["STAT"]
width: full
---

[**pg_logicalinspect**](https://www.postgresql.org/docs/devel/pglogicalinspect.html) : Logical decoding components inspection


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6890** | {{< badge content="pg_logicalinspect" link="https://www.postgresql.org/docs/devel/pglogicalinspect.html" >}} | {{< ext "pg_logicalinspect" >}} | `1.0` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_profile" >}} {{< ext "pg_tracing" >}} {{< ext "pg_show_plans" >}} {{< ext "pg_stat_kcache" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_store_plans" >}} {{< ext "pg_track_settings" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "N/A" "PostgreSQL 17: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 16: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 15: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 14: not available" "red" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_logicalinspect;
```



## Usage

> [pg_logicalinspect: inspect logical decoding snapshot files](https://www.postgresql.org/docs/current/pglogicalinspect.html)

pg_logicalinspect provides SQL functions to examine serialized logical decoding snapshots stored in `pg_logical/snapshots/`, useful for debugging and educational purposes.

### Functions

**`pg_get_logical_snapshot_meta(filename text)`** -- snapshot file metadata:

```sql
SELECT * FROM pg_get_logical_snapshot_meta('0-40796E18.snap');

 magic      | 1369563137
 checksum   | 1028045905
 version    | 6
```

**`pg_get_logical_snapshot_info(filename text)`** -- detailed snapshot information:

```sql
SELECT * FROM pg_get_logical_snapshot_info('0-40796E18.snap');

 state                    | consistent
 xmin                     | 751
 xmax                     | 751
 start_decoding_at        | 0/40796AF8
 two_phase_at             | 0/40796AF8
 initial_xmin_horizon     | 0
 building_full_snapshot   | f
 in_slot_creation         | f
 last_serialized_snapshot | 0/0
 committed_count          | 0
 committed_xip            |
 catchange_count          | 2
 catchange_xip            | {751,752}
```

### Listing Available Snapshots

Combine with `pg_ls_logicalsnapdir()` to discover and inspect all snapshots:

```sql
SELECT ss.name, meta.*
FROM pg_ls_logicalsnapdir() AS ss,
     pg_get_logical_snapshot_meta(ss.name) AS meta;

SELECT ss.name, info.*
FROM pg_ls_logicalsnapdir() AS ss,
     pg_get_logical_snapshot_info(ss.name) AS info;
```

### Access Control

Restricted to superusers and members of `pg_read_server_files` by default.
