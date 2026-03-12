---
title: "pg_overexplain"
linkTitle: "pg_overexplain"
description: "Allow EXPLAIN to dump even more details"
weight: 6880
categories: ["STAT"]
width: full
---

[**pg_overexplain**](https://www.postgresql.org/docs/devel/pgoverexplain.html) : Allow EXPLAIN to dump even more details


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6880** | {{< badge content="pg_overexplain" link="https://www.postgresql.org/docs/devel/pgoverexplain.html" >}} | {{< ext "pg_overexplain" >}} | `1.0` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_profile" >}} {{< ext "pg_tracing" >}} {{< ext "pg_show_plans" >}} {{< ext "pg_stat_kcache" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_store_plans" >}} {{< ext "pg_track_settings" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "N/A" "PostgreSQL 17: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 16: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 15: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 14: not available" "red" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_overexplain';
```


This extension does not need `CREATE EXTENSION` DDL command





## Usage

> [pg_overexplain: extended EXPLAIN with internal planner details](https://www.postgresql.org/docs/current/pgoverexplain.html)

pg_overexplain extends the `EXPLAIN` command with additional debugging options to display internal planner data structures. Primarily intended for planner debugging and development.

### Loading

```sql
LOAD 'pg_overexplain';
-- Or in postgresql.conf:
-- session_preload_libraries = 'pg_overexplain'
```

### EXPLAIN (DEBUG)

Displays internal plan tree information:

```sql
EXPLAIN (DEBUG) SELECT * FROM my_table WHERE id = 1;
```

Shows per-node fields:
- **Disabled Nodes** -- raw counter of disabled nodes
- **Parallel Safe** -- whether the node can appear under Gather
- **Plan Node ID** -- internal ID for parallel query coordination
- **extParam / allParam** -- parameters affecting the node

Shows per-query fields:
- **Command Type** -- query type (select, update, etc.)
- **Flags** -- hasReturning, hasModifyingCTE, canSetTag, transientPlan, etc.
- **Subplans Needing Rewind** -- subplan IDs requiring rewind
- **Relation OIDs** -- OIDs the plan depends on
- **Parse Location** -- location in the query string

### EXPLAIN (RANGE_TABLE)

Displays information about the query's range table entries:

```sql
EXPLAIN (RANGE_TABLE) SELECT * FROM t1 JOIN t2 ON t1.id = t2.id;
```

Shows range table index references (`Scan RTI`, `Nominal RTI`, `Append RTIs`, etc.) and dumps each range table entry with its kind (relation, subquery, join, cte, etc.) and entry-specific fields.

### Notes

- Output reflects internal planner data structures and may require reading source code to interpret
- Output format may change across PostgreSQL versions
- Not recommended for general production use
