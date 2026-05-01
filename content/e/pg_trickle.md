---
title: "pg_trickle"
linkTitle: "pg_trickle"
description: "Streaming tables and differential view maintenance for PostgreSQL 18"
weight: 2860
categories: ["FEAT"]
width: full
---

[**pg_trickle**](https://github.com/grove/pg-trickle) : Streaming tables and differential view maintenance for PostgreSQL 18


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2860** | {{< badge content="pg_trickle" link="https://github.com/grove/pg-trickle" >}} | {{< ext "pg_trickle" >}} | `0.40.0` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_ivm" >}} {{< ext "pg_incremental" >}} {{< ext "pg_partman" >}} {{< ext "timeseries" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.40.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_trickle` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.40.0` | {{< bg "18" "pg_trickle_18" "green" >}} {{< bg "17" "pg_trickle_17" "red" >}} {{< bg "16" "pg_trickle_16" "red" >}} {{< bg "15" "pg_trickle_15" "red" >}} {{< bg "14" "pg_trickle_14" "red" >}} | `pg_trickle_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.40.0` | {{< bg "18" "postgresql-18-pg-trickle" "green" >}} {{< bg "17" "postgresql-17-pg-trickle" "red" >}} {{< bg "16" "postgresql-16-pg-trickle" "red" >}} {{< bg "15" "postgresql-15-pg-trickle" "red" >}} {{< bg "14" "postgresql-14-pg-trickle" "red" >}} | `postgresql-$v-pg-trickle` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.40.0" "pg_trickle_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_trickle_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.40.0" "pg_trickle_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_trickle_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.40.0" "pg_trickle_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_trickle_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.40.0" "pg_trickle_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_trickle_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.40.0" "pg_trickle_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_trickle_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.40.0" "pg_trickle_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_trickle_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.40.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.40.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.40.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.40.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.40.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.40.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.40.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.40.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.40.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.40.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_trickle_18` | `0.40.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 4.5 MiB | [pg_trickle_18-0.40.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_trickle_18-0.40.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_trickle_18` | `0.40.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.0 MiB | [pg_trickle_18-0.40.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_trickle_18-0.40.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_trickle_18` | `0.40.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 4.5 MiB | [pg_trickle_18-0.40.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_trickle_18-0.40.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_trickle_18` | `0.40.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.2 MiB | [pg_trickle_18-0.40.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_trickle_18-0.40.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_trickle_18` | `0.40.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 4.5 MiB | [pg_trickle_18-0.40.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_trickle_18-0.40.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_trickle_18` | `0.40.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.2 MiB | [pg_trickle_18-0.40.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_trickle_18-0.40.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-trickle` | `0.40.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.7 MiB | [postgresql-18-pg-trickle_0.40.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.40.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-trickle` | `0.40.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.1 MiB | [postgresql-18-pg-trickle_0.40.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.40.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-trickle` | `0.40.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.7 MiB | [postgresql-18-pg-trickle_0.40.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.40.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-trickle` | `0.40.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.1 MiB | [postgresql-18-pg-trickle_0.40.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.40.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-trickle` | `0.40.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.1 MiB | [postgresql-18-pg-trickle_0.40.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.40.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-trickle` | `0.40.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.6 MiB | [postgresql-18-pg-trickle_0.40.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.40.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-trickle` | `0.40.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.1 MiB | [postgresql-18-pg-trickle_0.40.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.40.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-trickle` | `0.40.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.6 MiB | [postgresql-18-pg-trickle_0.40.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.40.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-trickle` | `0.40.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.0 MiB | [postgresql-18-pg-trickle_0.40.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.40.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-trickle` | `0.40.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.6 MiB | [postgresql-18-pg-trickle_0.40.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.40.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/grove/pg-trickle" title="Repository" icon="github" subtitle="github.com/grove/pg-trickle" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_trickle-0.40.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_trickle;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_trickle;		# install via package name, for the active PG version

pig install pg_trickle -v 18;   # install for PG 18

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_trickle';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_trickle;
```

## Usage

Sources: [README v0.40.0](https://github.com/grove/pg-trickle/blob/v0.40.0/README.md), [v0.40.0 release notes](https://github.com/grove/pg-trickle/releases/tag/v0.40.0), [SQL reference](https://github.com/grove/pg-trickle/blob/v0.40.0/docs/SQL_REFERENCE.md), [configuration guide](https://github.com/grove/pg-trickle/blob/v0.40.0/docs/CONFIGURATION.md), [GUC catalog](https://github.com/grove/pg-trickle/blob/v0.40.0/docs/GUC_CATALOG.md), [SQL API catalog](https://github.com/grove/pg-trickle/blob/v0.40.0/docs/SQL_API_CATALOG.md), [Cargo.toml](https://github.com/grove/pg-trickle/blob/v0.40.0/Cargo.toml)

`pg_trickle` provides stream tables for PostgreSQL 18: regular queryable tables whose contents are maintained from a defining SQL query. It uses incremental view maintenance when possible, can fall back to full recompute, and also supports `IMMEDIATE` mode for same-transaction maintenance.

Upstream v0.40.0 is still pre-1.0 and says APIs and configuration options may change before a stable 1.0 release. The Rust package is `pg_trickle` version `0.40.0`, uses Rust edition 2024, defaults to the `pg18` feature, and pins `pgrx = 0.18.0`. Build prerequisites in the README are PostgreSQL 18.x plus Rust 1.85+ with pgrx 0.18.x.

### Enable the Extension and Build Scope

Install from release packages when available, or build with pgrx:

```bash
cargo pgrx install --release
cargo pgrx package
```

Add the extension to PostgreSQL startup configuration and restart:

```sql
-- postgresql.conf
shared_preload_libraries = 'pg_trickle'
max_worker_processes = 8
```

```sql
CREATE EXTENSION pg_trickle;
```

`shared_preload_libraries` is required because the extension registers GUCs and a background worker at startup. `wal_level = logical` and replication slots are not required by default: `pg_trickle.cdc_mode = 'auto'` starts with trigger-based CDC and transitions to WAL-based capture only when logical WAL is available.

### Create and Refresh Stream Tables

```sql
CREATE TABLE orders (
    id int PRIMARY KEY,
    region text,
    amount numeric
);

SELECT pgtrickle.create_stream_table(
    'regional_totals',
    'SELECT region, SUM(amount) AS total, COUNT(*) AS cnt
     FROM orders GROUP BY region'
);

SELECT * FROM regional_totals;
SELECT pgtrickle.refresh_stream_table('regional_totals');
```

The main refresh modes are `AUTO`, `DIFFERENTIAL`, `FULL`, and `IMMEDIATE`. `AUTO` chooses differential maintenance when the query is differentiable and falls back to full recompute when needed. `DIFFERENTIAL` applies deltas only. `FULL` truncates and reloads from the defining query. `IMMEDIATE` uses statement-level IVM triggers and is maintained inside the same transaction as base-table DML.

```sql
SELECT pgtrickle.create_stream_table(
    'regional_totals_live',
    'SELECT region, SUM(amount) AS total, COUNT(*) AS cnt
     FROM orders GROUP BY region',
    schedule => NULL,
    refresh_mode => 'IMMEDIATE'
);
```

Schedules accept duration strings such as `'30s'`, `'5m'`, `'1h'`, cron expressions such as `'@hourly'`, or the default `'calculated'` schedule inherited from downstream dependents.

```sql
SELECT pgtrickle.create_stream_table(
    name         => 'hourly_totals',
    query        => 'SELECT region, SUM(amount) AS total FROM orders GROUP BY region',
    schedule     => '@hourly',
    refresh_mode => 'FULL'
);
```

### Lifecycle, SQL Coverage, and Operators

```sql
SELECT pgtrickle.alter_stream_table(
    'regional_totals',
    query => 'SELECT region, SUM(amount) AS total FROM orders GROUP BY region'
);

SELECT pgtrickle.drop_stream_table('regional_totals');
```

The SQL reference documents lifecycle calls such as `pgtrickle.create_stream_table()`, `pgtrickle.create_stream_table_if_not_exists()`, `pgtrickle.create_or_replace_stream_table()`, `pgtrickle.bulk_create()`, `pgtrickle.alter_stream_table()`, `pgtrickle.drop_stream_table()`, `pgtrickle.resume_stream_table()`, `pgtrickle.refresh_stream_table()`, and `pgtrickle.repair_stream_table()`.

The documented SQL coverage includes joins, aggregates, window functions, set operations, scalar and table subqueries, CTEs including `WITH RECURSIVE`, LATERAL/SRFs, `JSON_TABLE`, TopK queries with `ORDER BY ... LIMIT`, views as sources, tables without primary keys, and stream-table dependency DAGs. No custom SQL operator is the main user-facing API; users primarily interact through functions, views, catalog tables, GUCs, and normal SQL queries over stream tables.

### Operations and Introspection

```sql
SELECT * FROM pgtrickle.pgt_status();
SELECT * FROM pgtrickle.refresh_timeline(20);
SELECT * FROM pgtrickle.health_check();
SELECT * FROM pgtrickle.health_summary();
SELECT * FROM pgtrickle.pg_stat_stream_tables;
SELECT * FROM pgtrickle.change_buffer_sizes();
SELECT * FROM pgtrickle.dependency_tree();
SELECT * FROM pgtrickle.explain_st('regional_totals');
SELECT * FROM pgtrickle.slot_health();
SELECT * FROM pgtrickle.check_cdc_health();
```

Other documented views and catalog tables include `pgtrickle.stream_tables_info`, `pgtrickle.quick_health`, `pgtrickle.pgt_cdc_status`, `pgtrickle.pgt_stream_tables`, `pgtrickle.pgt_dependencies`, `pgtrickle.pgt_refresh_history`, `pgtrickle.pgt_change_tracking`, `pgtrickle.pgt_source_gates`, and `pgtrickle.pgt_refresh_groups`.

### Outbox, Inbox, Relay, and Snapshots

`pg_trickle` can publish stream-table deltas through the transactional outbox pattern and consume idempotent inbox tables.

```sql
SELECT pgtrickle.enable_outbox('public.regional_totals');
SELECT pgtrickle.create_consumer_group('billing_workers', 'public.regional_totals');
SELECT * FROM pgtrickle.poll_outbox('billing_workers', 'worker-1');
SELECT pgtrickle.commit_offset('billing_workers', 'worker-1', 42);

SELECT pgtrickle.create_inbox('orders_inbox');
SELECT pgtrickle.inbox_health('orders_inbox');
```

The SQL reference also documents snapshot operations and relay configuration helpers:

```sql
SELECT pgtrickle.snapshot_stream_table('public.regional_totals');
SELECT pgtrickle.restore_from_snapshot(
    'public.regional_totals',
    'pgtrickle.regional_totals_snapshot'
);

SELECT pgtrickle.set_relay_outbox(
    'orders-to-nats',
    'public.regional_totals',
    'relay_group_1',
    '{"type": "nats", "subject": "orders.deltas", "url": "nats://nats:4222"}'::jsonb
);
```

### Important GUCs

The v0.40.0 release adds generated docs for 125 configuration parameters. Common operational GUCs include:

- `pg_trickle.enabled`
- `pg_trickle.cdc_mode`
- `pg_trickle.scheduler_interval_ms`
- `pg_trickle.min_schedule_seconds`
- `pg_trickle.default_schedule_seconds`
- `pg_trickle.max_consecutive_errors`
- `pg_trickle.wal_transition_timeout`
- `pg_trickle.slot_lag_warning_threshold_mb`
- `pg_trickle.slot_lag_critical_threshold_mb`
- `pg_trickle.differential_max_change_ratio`
- `pg_trickle.refresh_strategy`
- `pg_trickle.cost_model_safety_margin`
- `pg_trickle.planner_aggressive`
- `pg_trickle.merge_join_strategy`
- `pg_trickle.merge_strategy`
- `pg_trickle.auto_backoff`
- `pg_trickle.tiered_scheduling`
- `pg_trickle.cleanup_use_truncate`
- `pg_trickle.block_source_ddl`
- `pg_trickle.buffer_alert_threshold`
- `pg_trickle.compact_threshold`
- `pg_trickle.max_buffer_rows`
- `pg_trickle.auto_index`
- `pg_trickle.aggregate_fast_path`
- `pg_trickle.template_cache`
- `pg_trickle.buffer_partitioning`
- `pg_trickle.ivm_topk_max_limit`
- `pg_trickle.ivm_recursive_max_depth`
- `pg_trickle.parallel_refresh_mode`
- `pg_trickle.max_dynamic_refresh_workers`
- `pg_trickle.max_concurrent_refreshes`
- `pg_trickle.change_buffer_schema`
- `pg_trickle.foreign_table_polling`
- `pg_trickle.matview_polling`
- `pg_trickle.log_delta_sql`
- `pg_trickle.metrics_port`
- `pg_trickle.outbox_enabled`
- `pg_trickle.inbox_enabled`
- `pg_trickle.citus_st_lock_lease_ms`
- `pg_trickle.citus_worker_retry_ticks`
- `pg_trickle.enable_vector_agg`
- `pg_trickle.enable_trace_propagation`
- `pg_trickle.otel_endpoint`
- `pg_trickle.trace_id`
- `pg_trickle.cdc_capture_mode`

`pg_trickle.event_driven_wake` and `pg_trickle.wake_debounce_ms` are preserved for upgrade compatibility in v0.40.0 but are formally deprecated and have no effect, because PostgreSQL background workers cannot use `LISTEN`; the scheduler uses latch-based polling.

### v0.40.0 Operator Notes

The v0.40.0 release focuses on operator confidence. It adds generated GUC and SQL API catalogs, a security model and secret-handling guide, a drain-mode runbook with end-to-end tests, expanded Prometheus alert rules, dbt and relay parity helpers, strict unsafe-block CI gating, clearer template-cache documentation, formal deprecation of `event_driven_wake`, and secret scanning in CI.

The generated SQL API catalog lists 24 SQL-callable `#[pg_extern]` functions, including `pgtrickle.metrics_summary()`, `pgtrickle.outbox_status()`, `pgtrickle.outbox_rows_consumed()`, `pgtrickle.commit_offset()`, `pgtrickle.consumer_heartbeat()`, `pgtrickle.seek_offset()`, `pgtrickle.inbox_health()`, `pgtrickle.inbox_is_my_partition()`, `pgtrickle.snapshot_stream_table()`, `pgtrickle.restore_from_snapshot()`, `pgtrickle.restore_stream_tables()`, `pgtrickle.recommend_schedule()`, and `pgtrickle.schedule_recommendations()`.

### Caveats

- `pg_trickle` v0.40.0 is PostgreSQL 18 only; the release packages are named for `pg18`, and Cargo defaults to the `pg18` pgrx feature.
- The extension control file marks it `superuser = true` and `trusted = false`.
- Direct DML on stream tables is not allowed because their contents are managed by the refresh engine.
- `IMMEDIATE` mode bypasses CDC and uses statement-level IVM triggers; WAL CDC is asynchronous and incompatible with in-transaction maintenance.
- Materialized views in `DIFFERENTIAL` mode require `pg_trickle.matview_polling = on`; `FULL` mode works without that snapshot-comparison path.
- `LIMIT` or `OFFSET` without `ORDER BY` is rejected for stream-table definitions; use `ORDER BY ... LIMIT` for TopK.
- Volatile functions are rejected by default in defining queries according to `pg_trickle.volatile_function_policy`.
