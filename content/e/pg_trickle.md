---
title: "pg_trickle"
linkTitle: "pg_trickle"
description: "Streaming tables and differential view maintenance for PostgreSQL 18"
weight: 2860
categories: ["FEAT"]
languages: ["Rust"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_trickle**](https://github.com/trickle-labs/pg-trickle) : Streaming tables and differential view maintenance for PostgreSQL 18


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2860** | {{< badge content="pg_trickle" link="https://github.com/trickle-labs/pg-trickle" >}} | {{< ext "pg_trickle" >}} | `0.108.1` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgtrickle` `pgtrickle_changes` |
|   **See Also**    | {{< ext "pg_ivm" >}} {{< ext "pg_incremental" >}} {{< ext "timescaledb" >}} {{< ext "pg_duckdb" >}} {{< ext "pg_partman" >}} {{< ext "pg_ttl_index" >}} {{< ext "duckdb_fdw" >}} {{< ext "pg_lake" >}} |

> [!Note] PG18 only; requires preload and ships pg_trickle_dump. Follow the packaged upgrade guide. v0.108.2 is tag-only as of 2026-10-04; the current GitHub Release and packaged baseline remain 0.108.1.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.108.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_trickle` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.108.1` | {{< bg "18" "pg_trickle_18" "green" >}} {{< bg "17" "pg_trickle_17" "red" >}} {{< bg "16" "pg_trickle_16" "red" >}} {{< bg "15" "pg_trickle_15" "red" >}} {{< bg "14" "pg_trickle_14" "red" >}} | `pg_trickle_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.108.1` | {{< bg "18" "postgresql-18-pg-trickle" "green" >}} {{< bg "17" "postgresql-17-pg-trickle" "red" >}} {{< bg "16" "postgresql-16-pg-trickle" "red" >}} {{< bg "15" "postgresql-15-pg-trickle" "red" >}} {{< bg "14" "postgresql-14-pg-trickle" "red" >}} | `postgresql-$v-pg-trickle` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.108.1" "pg_trickle_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_trickle_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_trickle_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_trickle_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_trickle_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.108.1" "pg_trickle_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_trickle_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_trickle_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_trickle_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_trickle_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.108.1" "pg_trickle_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_trickle_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_trickle_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_trickle_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_trickle_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.108.1" "pg_trickle_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_trickle_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_trickle_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_trickle_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_trickle_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.108.1" "pg_trickle_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_trickle_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_trickle_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_trickle_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_trickle_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.108.1" "pg_trickle_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_trickle_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_trickle_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_trickle_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_trickle_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.108.1" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-trickle : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.108.1" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-trickle : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.108.1" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-trickle : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.108.1" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-trickle : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.108.1" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-trickle : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.108.1" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-trickle : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.108.1" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-trickle : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.108.1" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-trickle : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.108.1" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-trickle : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.108.1" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-trickle : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-trickle : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_trickle_18` | `0.108.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 6.5 MiB | [pg_trickle_18-0.108.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_trickle_18-0.108.1-1PGSTY.el8.x86_64.rpm) |
| `pg_trickle_18` | `0.108.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 5.4 MiB | [pg_trickle_18-0.108.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_trickle_18-0.108.1-1PGSTY.el8.aarch64.rpm) |
| `pg_trickle_18` | `0.108.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 6.4 MiB | [pg_trickle_18-0.108.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_trickle_18-0.108.1-1PGSTY.el9.x86_64.rpm) |
| `pg_trickle_18` | `0.108.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 5.6 MiB | [pg_trickle_18-0.108.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_trickle_18-0.108.1-1PGSTY.el9.aarch64.rpm) |
| `pg_trickle_18` | `0.108.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 6.4 MiB | [pg_trickle_18-0.108.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_trickle_18-0.108.1-1PGSTY.el10.x86_64.rpm) |
| `pg_trickle_18` | `0.108.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 5.6 MiB | [pg_trickle_18-0.108.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_trickle_18-0.108.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-trickle` | `0.108.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.6 MiB | [postgresql-18-pg-trickle_0.108.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.108.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-trickle` | `0.108.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.6 MiB | [postgresql-18-pg-trickle_0.108.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.108.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-trickle` | `0.108.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.6 MiB | [postgresql-18-pg-trickle_0.108.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.108.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-trickle` | `0.108.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.6 MiB | [postgresql-18-pg-trickle_0.108.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.108.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-trickle` | `0.108.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.1 MiB | [postgresql-18-pg-trickle_0.108.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.108.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-trickle` | `0.108.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.4 MiB | [postgresql-18-pg-trickle_0.108.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.108.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-trickle` | `0.108.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.1 MiB | [postgresql-18-pg-trickle_0.108.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.108.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-trickle` | `0.108.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.4 MiB | [postgresql-18-pg-trickle_0.108.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.108.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-trickle` | `0.108.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.1 MiB | [postgresql-18-pg-trickle_0.108.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.108.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-trickle` | `0.108.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.3 MiB | [postgresql-18-pg-trickle_0.108.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.108.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/trickle-labs/pg-trickle" title="Repository" icon="github" subtitle="github.com/trickle-labs/pg-trickle" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_trickle-0.108.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_trickle;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

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

Sources:

- [sql/pg_trickle--0.108.0--0.108.1.sql](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/sql/pg_trickle--0.108.0--0.108.1.sql)
- [Version 0.108.1 README](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/README.md)
- [SQL reference](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/docs/SQL_REFERENCE.md)
- [Configuration](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/docs/CONFIGURATION.md)
- [GUC catalog](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/docs/GUC_CATALOG.md)
- [Upgrade guide](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/docs/UPGRADING.md)
- [Control file](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/pg_trickle.control)
- [0.107.0 to 0.108.0 migration](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/sql/pg_trickle--0.107.0--0.108.0.sql)

`pg_trickle` 0.108.1 maintains stream tables on PostgreSQL 18: ordinary queryable tables derived from a SQL query, refreshed incrementally when supported or recomputed in full. Tables may depend on other stream tables, forming a dependency graph. Same-transaction maintenance is also available.

### Enable the Extension

Append `pg_trickle` to `shared_preload_libraries` and restart PostgreSQL, then install as a superuser. Size the background-worker pool for the deployment; the upstream example uses eight workers.

```ini
shared_preload_libraries = 'pg_trickle'
max_worker_processes = 8
```

```sql
CREATE EXTENSION pg_trickle;
```

The default `pg_trickle.cdc_mode` is `trigger`: transactional change capture needs neither logical WAL nor replication slots. Opt-in `auto` starts with triggers and can move eligible sources to receipt-backed WAL capture. Explicit `wal` also falls back to triggers when admission or logical-decoding prerequisites fail. These choices have different write overhead and operational requirements.

### Create and Refresh a Stream Table

```sql
CREATE TABLE orders (id bigint PRIMARY KEY, region text, amount numeric);
SELECT pgtrickle.create_stream_table(
    name => 'regional_totals',
    query => 'SELECT region, SUM(amount) AS total, COUNT(*) AS cnt FROM orders GROUP BY region',
    schedule => '30s',
    refresh_mode => 'AUTO'
);
INSERT INTO orders VALUES (1, 'east', 10);
SELECT pgtrickle.refresh_stream_table('regional_totals');
SELECT * FROM regional_totals;
```

`initialize` defaults to true, so creation populates the result. `schedule` accepts durations, cron expressions such as `@hourly`, or the default `calculated` schedule inherited from downstream dependents. `AUTO` selects differential maintenance where possible and can fall back to full refresh. `DIFFERENTIAL` rejects queries it cannot maintain incrementally; `FULL` truncates and reloads the result.

`IMMEDIATE` uses statement-level triggers inside the base-table write transaction. It does not use WAL capture and rejects an effective explicit WAL request. Use the documented query-admission rules for joins, aggregates, subqueries, recursive queries, and other supported shapes; support is not a promise that every SQL expression is differentiable.

### Lifecycle and Monitoring

`pgtrickle.alter_stream_table` changes definitions or refresh policy, and `pgtrickle.drop_stream_table` removes managed tables. `pgtrickle.repair_stream_table` repairs missing capture infrastructure and resets maintenance state after events such as restore or operator DDL. Use lifecycle APIs instead of direct writes or foreign keys against managed stream tables.

```sql
SELECT * FROM pgtrickle.pgt_status();
SELECT * FROM pgtrickle.health_check();
SELECT * FROM pgtrickle.dependency_tree();
SELECT * FROM pgtrickle.explain_st('regional_totals');
```

Lifecycle functions require explicit execution grants and ownership checks; administrator-wide operations are restricted to the extension owner or superuser. Arbitrary-SQL helpers preserve caller privileges. Capture triggers add work to source writes, and refresh failures can accumulate change buffers, so monitor health and storage instead of treating a schedule as a hard freshness guarantee.

### External Coordination and Output Deltas

`orchestration_mode` selects `MANAGED` scheduling or `EXTERNAL` coordination. External coordination cannot be combined with immediate maintenance. `pgtrickle.integration_capabilities` advertises the available contracts; version 0.108.0 exposes Graph V1 1.2 and Delta V1 1.1.

`pgtrickle.output_delta_consumer_status` reports consumers, while `pgtrickle.validate_output_delta_consumer` checks whether a consumer can resume. `pgtrickle.request_output_delta_resnapshot`, `pgtrickle.begin_output_delta_resnapshot`, and `pgtrickle.ack_output_delta_resnapshot` manage rebuilding a baseline. A resnapshot is fenced by database-instance identity, output-contract digest, and row-identity version. Follow the exact SQL reference signatures and acknowledgement protocol before advancing external delivery.

### Upgrade to 0.108.1

Install the new library and extension files before applying the packaged migration:

```sql
ALTER EXTENSION pg_trickle UPDATE TO '0.108.1';
SELECT * FROM pgtrickle.output_delta_consumer_status();
```

The upgrade preserves consumers, cursors, batches, and typed payload while adding resnapshot fences. A 0.106.1 installation can traverse the packaged 0.107.0 migration. Validate every consumer before resuming delivery; `INVALIDATED` or `RESNAPSHOT_REQUIRED` requires a new acknowledged baseline. Older releases through 0.105.2 had documented differential-result bugs for certain query shapes; upgrading does not automatically repair previously materialized rows. Use the upgrade guide’s comparison and repair procedure where applicable.

### Version 0.108.1

This patch keeps downstream `IMMEDIATE` tables current after upstream FULL refresh or truncation, recovers missing change buffers, and fixes unintended suspension after source schema changes and scalar-subquery differential refresh. PostgreSQL 18.6 support is added. After installing matching files run `ALTER EXTENSION pg_trickle UPDATE TO '0.108.1'` and verify dependent stream-table results and capture health.
