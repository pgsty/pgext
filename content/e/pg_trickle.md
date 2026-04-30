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
| **2860** | {{< badge content="pg_trickle" link="https://github.com/grove/pg-trickle" >}} | {{< ext "pg_trickle" >}} | `0.31.0` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_ivm" >}} {{< ext "pg_incremental" >}} {{< ext "pg_partman" >}} {{< ext "timeseries" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.31.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_trickle` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.31.0` | {{< bg "18" "pg_trickle_18" "green" >}} {{< bg "17" "pg_trickle_17" "red" >}} {{< bg "16" "pg_trickle_16" "red" >}} {{< bg "15" "pg_trickle_15" "red" >}} {{< bg "14" "pg_trickle_14" "red" >}} | `pg_trickle_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.31.0` | {{< bg "18" "postgresql-18-pg-trickle" "green" >}} {{< bg "17" "postgresql-17-pg-trickle" "red" >}} {{< bg "16" "postgresql-16-pg-trickle" "red" >}} {{< bg "15" "postgresql-15-pg-trickle" "red" >}} {{< bg "14" "postgresql-14-pg-trickle" "red" >}} | `postgresql-$v-pg-trickle` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.31.0" "pg_trickle_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_trickle_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.31.0" "pg_trickle_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_trickle_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.31.0" "pg_trickle_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_trickle_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.31.0" "pg_trickle_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_trickle_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.31.0" "pg_trickle_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_trickle_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.31.0" "pg_trickle_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_trickle_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.31.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.31.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.31.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.31.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.31.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.31.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.20.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.20.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.31.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.31.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_trickle_18` | `0.31.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 4.4 MiB | [pg_trickle_18-0.31.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_trickle_18-0.31.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_trickle_18` | `0.31.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.9 MiB | [pg_trickle_18-0.31.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_trickle_18-0.31.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_trickle_18` | `0.31.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 4.3 MiB | [pg_trickle_18-0.31.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_trickle_18-0.31.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_trickle_18` | `0.31.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.1 MiB | [pg_trickle_18-0.31.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_trickle_18-0.31.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_trickle_18` | `0.31.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 4.3 MiB | [pg_trickle_18-0.31.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_trickle_18-0.31.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_trickle_18` | `0.31.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.1 MiB | [pg_trickle_18-0.31.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_trickle_18-0.31.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-trickle` | `0.31.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.6 MiB | [postgresql-18-pg-trickle_0.31.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.31.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-trickle` | `0.31.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.0 MiB | [postgresql-18-pg-trickle_0.31.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.31.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-trickle` | `0.31.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.6 MiB | [postgresql-18-pg-trickle_0.31.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.31.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-trickle` | `0.31.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.9 MiB | [postgresql-18-pg-trickle_0.31.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.31.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-trickle` | `0.31.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.9 MiB | [postgresql-18-pg-trickle_0.31.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.31.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-trickle` | `0.31.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.5 MiB | [postgresql-18-pg-trickle_0.31.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.31.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-trickle` | `0.20.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.1 MiB | [postgresql-18-pg-trickle_0.20.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.20.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-trickle` | `0.20.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.7 MiB | [postgresql-18-pg-trickle_0.20.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.20.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-trickle` | `0.31.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.9 MiB | [postgresql-18-pg-trickle_0.31.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.31.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-trickle` | `0.31.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.5 MiB | [postgresql-18-pg-trickle_0.31.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.31.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/grove/pg-trickle" title="Repository" icon="github" subtitle="github.com/grove/pg-trickle" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_trickle-0.31.0.tar.gz" >}}
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

> Sources: [README v0.20.0](https://github.com/grove/pg-trickle/blob/v0.20.0/README.md), [v0.20.0 release notes](https://github.com/grove/pg-trickle/releases/tag/v0.20.0)

`pg_trickle` provides incrementally maintained stream tables for PostgreSQL 18. It keeps query results fresh using differential refresh when possible, with full recompute and immediate in-transaction modes also documented upstream.

### Enable the Extension

```sql
-- postgresql.conf
shared_preload_libraries = 'pg_trickle'
max_worker_processes = 8

CREATE EXTENSION pg_trickle;
```

The upstream README notes that `wal_level = logical` is not required by default. CDC starts with row-level triggers and can switch to WAL-based capture when `pg_trickle.cdc_mode = 'auto'`.

### Create and Refresh Stream Tables

```sql
SELECT pgtrickle.create_stream_table(
    'regional_totals',
    'SELECT region, SUM(amount) AS total, COUNT(*) AS cnt
     FROM orders GROUP BY region'
);

SELECT * FROM regional_totals;
SELECT pgtrickle.refresh_stream_table('regional_totals');
```

The documented refresh modes are `AUTO`, `DIFFERENTIAL`, `FULL`, and `IMMEDIATE`.

```sql
SELECT pgtrickle.create_stream_table(
    'regional_totals_live',
    'SELECT region, SUM(amount) AS total, COUNT(*) AS cnt
     FROM orders GROUP BY region',
    schedule => NULL,
    refresh_mode => 'IMMEDIATE'
);
```

### Operations and Introspection

```sql
SELECT pgtrickle.alter_stream_table(
    'regional_totals',
    query => 'SELECT region, SUM(amount) AS total FROM orders GROUP BY region'
);

SELECT * FROM pgtrickle.pgt_status();
SELECT * FROM pgtrickle.health_check();
SELECT * FROM pgtrickle.refresh_timeline(20);
SELECT * FROM pgtrickle.change_buffer_sizes();
SELECT * FROM pgtrickle.dependency_tree();
```

The README also documents broad SQL coverage including joins, aggregates, window functions, recursive CTEs, subqueries, set operations, and TopK queries.

### v0.20.0 Monitoring Additions

The `v0.20.0` release adds built-in self-monitoring:

- `pgtrickle.setup_dog_feeding()`
- `pgtrickle.teardown_dog_feeding()`
- `pgtrickle.dog_feeding_status()`

Release notes describe five monitoring stream tables that analyze refresh history and can emit threshold advice and alerts.
