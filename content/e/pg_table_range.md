---
title: "pg_table_range"
linkTitle: "pg_table_range"
description: "Planning-time partition pruning by per-partition data ranges, including non-key columns."
weight: 2570
categories: ["OLAP"]
languages: ["Rust"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_table_range**](https://github.com/bitner/pg_table_range) : Planning-time partition pruning by per-partition data ranges, including non-key columns.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2570** | {{< badge content="pg_table_range" link="https://github.com/bitner/pg_table_range" >}} | {{< ext "pg_table_range" >}} | `0.1.2` | {{< category "OLAP" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_kpart" >}} `pg_prttn_tools` {{< ext "qdgc" >}} `tab_tier` `brin4postgis` `planinstr` `pg_stash_advice` `pargres` `edb_foreignkeyconstraintmanager` `gogudb` |

> [!Note] Experimental. 0.1.2 fixes missing rows and backend crashes. From 0.1.1: install new binaries, restart all backends, UPDATE extension, REINDEX each table_range index (parent for partitioned tables); old summaries are ignored until rebuilt.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_table_range` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.2` | {{< bg "18" "pg_table_range_18" "green" >}} {{< bg "17" "pg_table_range_17" "green" >}} {{< bg "16" "pg_table_range_16" "green" >}} {{< bg "15" "pg_table_range_15" "red" >}} {{< bg "14" "pg_table_range_14" "red" >}} | `pg_table_range_$v` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.2" "pg_table_range_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_table_range_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_table_range_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_table_range_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.2" "pg_table_range_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_table_range_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_table_range_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_table_range_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.2" "pg_table_range_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_table_range_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_table_range_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_table_range_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.2" "pg_table_range_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_table_range_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_table_range_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_table_range_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.2" "pg_table_range_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_table_range_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_table_range_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_table_range_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.2" "pg_table_range_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_table_range_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_table_range_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_table_range_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} | {{< bg "N/A" "pg_table_range : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_table_range_18` | `0.1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 984.1 KiB | [pg_table_range_18-0.1.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_table_range_18-0.1.2-1PGSTY.el8.x86_64.rpm) |
| `pg_table_range_18` | `0.1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 851.0 KiB | [pg_table_range_18-0.1.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_table_range_18-0.1.2-1PGSTY.el8.aarch64.rpm) |
| `pg_table_range_18` | `0.1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1010.2 KiB | [pg_table_range_18-0.1.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_table_range_18-0.1.2-1PGSTY.el9.x86_64.rpm) |
| `pg_table_range_18` | `0.1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 909.1 KiB | [pg_table_range_18-0.1.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_table_range_18-0.1.2-1PGSTY.el9.aarch64.rpm) |
| `pg_table_range_18` | `0.1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1009.9 KiB | [pg_table_range_18-0.1.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_table_range_18-0.1.2-1PGSTY.el10.x86_64.rpm) |
| `pg_table_range_18` | `0.1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 889.5 KiB | [pg_table_range_18-0.1.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_table_range_18-0.1.2-1PGSTY.el10.aarch64.rpm) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_table_range_17` | `0.1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 983.2 KiB | [pg_table_range_17-0.1.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_table_range_17-0.1.2-1PGSTY.el8.x86_64.rpm) |
| `pg_table_range_17` | `0.1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 848.2 KiB | [pg_table_range_17-0.1.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_table_range_17-0.1.2-1PGSTY.el8.aarch64.rpm) |
| `pg_table_range_17` | `0.1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1000.6 KiB | [pg_table_range_17-0.1.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_table_range_17-0.1.2-1PGSTY.el9.x86_64.rpm) |
| `pg_table_range_17` | `0.1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 906.2 KiB | [pg_table_range_17-0.1.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_table_range_17-0.1.2-1PGSTY.el9.aarch64.rpm) |
| `pg_table_range_17` | `0.1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1000.2 KiB | [pg_table_range_17-0.1.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_table_range_17-0.1.2-1PGSTY.el10.x86_64.rpm) |
| `pg_table_range_17` | `0.1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 887.9 KiB | [pg_table_range_17-0.1.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_table_range_17-0.1.2-1PGSTY.el10.aarch64.rpm) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_table_range_16` | `0.1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 980.6 KiB | [pg_table_range_16-0.1.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_table_range_16-0.1.2-1PGSTY.el8.x86_64.rpm) |
| `pg_table_range_16` | `0.1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 846.0 KiB | [pg_table_range_16-0.1.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_table_range_16-0.1.2-1PGSTY.el8.aarch64.rpm) |
| `pg_table_range_16` | `0.1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1004.9 KiB | [pg_table_range_16-0.1.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_table_range_16-0.1.2-1PGSTY.el9.x86_64.rpm) |
| `pg_table_range_16` | `0.1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 904.7 KiB | [pg_table_range_16-0.1.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_table_range_16-0.1.2-1PGSTY.el9.aarch64.rpm) |
| `pg_table_range_16` | `0.1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1004.8 KiB | [pg_table_range_16-0.1.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_table_range_16-0.1.2-1PGSTY.el10.x86_64.rpm) |
| `pg_table_range_16` | `0.1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 887.9 KiB | [pg_table_range_16-0.1.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_table_range_16-0.1.2-1PGSTY.el10.aarch64.rpm) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bitner/pg_table_range" title="Repository" icon="github" subtitle="github.com/bitner/pg_table_range" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_table_range-0.1.2.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_table_range;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_table_range;		# install via package name, for the active PG version

pig install pg_table_range -v 18;   # install for PG 18
pig install pg_table_range -v 17;   # install for PG 17
pig install pg_table_range -v 16;   # install for PG 16

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_table_range;
```

## Usage

Sources:

- [Upstream README](https://github.com/bitner/pg_table_range/blob/v0.1.2/README.md)
- [Version and PostgreSQL feature matrix](https://github.com/bitner/pg_table_range/blob/v0.1.2/Cargo.toml)
- [Planner-hook and GUC registration](https://github.com/bitner/pg_table_range/blob/v0.1.2/src/lib.rs)
- [0.1.2 upgrade and correctness fixes](https://github.com/bitner/pg_table_range/blob/v0.1.2/CHANGELOG.md)

pg_table_range is an experimental PostgreSQL 16–18 extension for planning-time partition pruning on columns that are not partition keys. Its table_range index access method stores each partition's actual scalar minimum and maximum, range extent, or PostGIS geometry extent. A planner hook removes a partition only when its summary proves that the partition cannot match.

### Basic use

Build the summary index on a partitioned table and compare plans with representative predicates:

```sql
CREATE EXTENSION pg_table_range;

CREATE INDEX events_tr
ON events USING table_range (val, created_at);

EXPLAIN (COSTS OFF)
SELECT *
FROM events
WHERE val >= 250;

SET table_range.enable_pruning = off;
EXPLAIN (COSTS OFF)
SELECT *
FROM events
WHERE val >= 250;
RESET table_range.enable_pruning;
```

The index is never selected to scan rows; it owns compact summaries used by the planner. Inserts widen summaries automatically. Deletes remain safe but can leave them wider than live data until REINDEX tightens them; VACUUM does not rebuild summaries.

Supported pruning predicates include scalar comparisons, BETWEEN, null tests, constant IN or ANY lists, and overlap for range and PostGIS geometry types. Unsupported expressions conservatively keep the partition.

### Upgrade from 0.1.1

Version 0.1.2 fixes missing rows and backend crashes in 0.1.1. Install the new binaries and restart PostgreSQL so no backend still runs the old code. Then update the extension and rebuild every table_range index, using the parent index for partitioned tables:

```sql
ALTER EXTENSION pg_table_range UPDATE TO '0.1.2';
REINDEX INDEX events_tr;
```

Old summaries are ignored until rebuilt, keeping all partitions rather than risking incorrect pruning. Partial indexes are rejected. REINDEX is also the way to tighten summaries after deletes.

### Caveats

- Upstream explicitly labels the project experimental and heavily AI-driven. Qualify correctness, crash recovery, upgrades, and planner behavior on production-shaped data before adoption.
- Planning cost is linear in the partition count. Native partition-key pruning is far cheaper; use pg_table_range only for selective non-key predicates where the execution savings justify planning work.
- Index creation and non-key planning lock every partition. With thousands of partitions, the default max_locks_per_transaction can fail with out-of-shared-memory errors; increasing it consumes shared memory and requires a restart.
- Deletes can reduce selectivity until maintenance rebuilds or tightens summaries. An invalid table_range index silently disables its pruning.
- The control file requires superuser installation. Limit table and index DDL privileges and test the Rust/pgrx build for the exact PostgreSQL major version.
