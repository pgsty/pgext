---
title: "old_snapshot"
linkTitle: "old_snapshot"
description: "utilities in support of old_snapshot_threshold"
weight: 5960
categories: ["ADMIN"]
width: full
---

[**old_snapshot**](https://www.postgresql.org/docs/current/oldsnapshot.html) : utilities in support of old_snapshot_threshold


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5960** | {{< badge content="old_snapshot" link="https://www.postgresql.org/docs/current/oldsnapshot.html" >}} | {{< ext "old_snapshot" >}} | `1.0` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pageinspect" >}} {{< ext "pg_visibility" >}} {{< ext "pgstattuple" >}} {{< ext "pg_prewarm" >}} {{< ext "pg_buffercache" >}} {{< ext "amcheck" >}} {{< ext "pg_surgery" >}} {{< ext "toastinfo" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "N/A" "PostgreSQL 18: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 17: not available" "red" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION old_snapshot;
```




## Usage

> [old_snapshot: utilities in support of old_snapshot_threshold](https://www.postgresql.org/docs/16/oldsnapshot.html)

The `old_snapshot` extension provides inspection functions for the server state related to the `old_snapshot_threshold` configuration parameter.

Note: this chapter is removed from PostgreSQL current docs in newer releases; use versioned docs when needed.

### Function

```sql
-- View the timestamp-to-XID mapping table
SELECT * FROM pg_old_snapshot_time_mapping();
```

### Function Signature

```sql
pg_old_snapshot_time_mapping(
    array_offset OUT int4,
    end_timestamp OUT timestamptz,
    newest_xmin OUT xid
) RETURNS SETOF record
```

### Output Columns

| Column | Type | Description |
|--------|------|-------------|
| `array_offset` | int4 | Index position in the mapping array |
| `end_timestamp` | timestamptz | End of the corresponding one-minute interval |
| `newest_xmin` | xid | Newest xmin of any snapshot taken during that minute |

### Context

PostgreSQL's `old_snapshot_threshold` parameter controls how long a snapshot can remain valid. The server maintains an internal mapping of timestamps to transaction IDs to implement this feature. This extension exposes that mapping for inspection and debugging.

```sql
-- Check the old_snapshot_threshold setting
SHOW old_snapshot_threshold;

-- Inspect the current mapping entries
SELECT array_offset, end_timestamp, newest_xmin
FROM pg_old_snapshot_time_mapping()
ORDER BY array_offset;
```
