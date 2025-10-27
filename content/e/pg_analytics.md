---
title: "pg_analytics"
linkTitle: "pg_analytics"
description: "Postgres for analytics, powered by DuckDB"
weight: 2420
categories: ["OLAP"]
width: full
---

Postgres for analytics, powered by DuckDB


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2420** | {{< badge content="pg_analytics" link="https://github.com/paradedb/pg_analytics" >}} | {{< ext "pg_analytics" >}} | `0.3.7` | {{< category "OLAP" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_duckdb" >}} {{< ext "pg_mooncake" >}} {{< ext "duckdb_fdw" >}} {{< ext "pg_parquet" >}} {{< ext "columnar" >}} {{< ext "citus_columnar" >}} {{< ext "orioledb" >}} {{< ext "citus" >}} |

> [!Note] pgrx=0.13.0, archived, no longer maintained


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_analytics" >}} | `0.3.7` | {{< bg "18" "pg_analytics_18" "red" >}} {{< bg "17" "pg_analytics_17" "green" >}} {{< bg "16" "pg_analytics_16" "green" >}} {{< bg "15" "pg_analytics_15" "green" >}} {{< bg "14" "pg_analytics_14" "green" >}} | `pg_analytics_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_analytics" >}} | `0.3.7` | {{< bg "18" "postgresql-18-pg-analytics" "red" >}} {{< bg "17" "postgresql-17-pg-analytics" "green" >}} {{< bg "16" "postgresql-16-pg-analytics" "green" >}} {{< bg "15" "postgresql-15-pg-analytics" "green" >}} {{< bg "14" "postgresql-14-pg-analytics" "green" >}} | `postgresql-$v-pg-analytics` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pg_analytics_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.7" "pg_analytics_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "pg_analytics_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "pg_analytics_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "pg_analytics_14 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pg_analytics_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.7" "pg_analytics_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "pg_analytics_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "pg_analytics_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "pg_analytics_14 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pg_analytics_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.7" "pg_analytics_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "pg_analytics_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "pg_analytics_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "pg_analytics_14 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pg_analytics_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.7" "pg_analytics_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "pg_analytics_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "pg_analytics_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "pg_analytics_14 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-analytics : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.7" "postgresql-17-pg-analytics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "postgresql-16-pg-analytics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "postgresql-15-pg-analytics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "postgresql-14-pg-analytics : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-analytics : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.7" "postgresql-17-pg-analytics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "postgresql-16-pg-analytics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "postgresql-15-pg-analytics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "postgresql-14-pg-analytics : AVAIL 1" "green" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-analytics : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.7" "postgresql-17-pg-analytics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "postgresql-16-pg-analytics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "postgresql-15-pg-analytics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "postgresql-14-pg-analytics : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-analytics : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.7" "postgresql-17-pg-analytics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "postgresql-16-pg-analytics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "postgresql-15-pg-analytics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "postgresql-14-pg-analytics : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-analytics : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.7" "postgresql-17-pg-analytics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "postgresql-16-pg-analytics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "postgresql-15-pg-analytics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "postgresql-14-pg-analytics : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-analytics : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.7" "postgresql-17-pg-analytics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "postgresql-16-pg-analytics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "postgresql-15-pg-analytics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.7" "postgresql-14-pg-analytics : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_analytics_17` | 0.3.7 | `el8.x86_64` | pigsty | 11.1 MiB | [pg_analytics_17-0.3.7-1PARADEDB.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_analytics_17-0.3.7-1PARADEDB.el8.x86_64.rpm) |
| `pg_analytics_17` | 0.3.7 | `el8.aarch64` | pigsty | 9.7 MiB | [pg_analytics_17-0.3.7-1PARADEDB.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_analytics_17-0.3.7-1PARADEDB.el8.aarch64.rpm) |
| `pg_analytics_17` | 0.3.7 | `el9.x86_64` | pigsty | 11.1 MiB | [pg_analytics_17-0.3.7-1PARADEDB.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_analytics_17-0.3.7-1PARADEDB.el9.x86_64.rpm) |
| `pg_analytics_17` | 0.3.7 | `el9.aarch64` | pigsty | 10.2 MiB | [pg_analytics_17-0.3.7-1PARADEDB.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_analytics_17-0.3.7-1PARADEDB.el9.aarch64.rpm) |
| `postgresql-17-pg-analytics` | 0.3.7 | `d12.x86_64` | pigsty | 10.2 MiB | [postgresql-17-pg-analytics_0.3.7_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-17-pg-analytics/postgresql-17-pg-analytics_0.3.7_amd64.deb) |
| `postgresql-17-pg-analytics` | 0.3.7 | `d12.aarch64` | pigsty | 8.8 MiB | [postgresql-17-pg-analytics_0.3.7_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-17-pg-analytics/postgresql-17-pg-analytics_0.3.7_arm64.deb) |
| `postgresql-17-pg-analytics` | 0.3.7 | `u22.x86_64` | pigsty | 10.3 MiB | [postgresql-17-pg-analytics_0.3.7_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-17-pg-analytics/postgresql-17-pg-analytics_0.3.7_amd64.deb) |
| `postgresql-17-pg-analytics` | 0.3.7 | `u22.aarch64` | pigsty | 9.0 MiB | [postgresql-17-pg-analytics_0.3.7_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-17-pg-analytics/postgresql-17-pg-analytics_0.3.7_arm64.deb) |
| `postgresql-17-pg-analytics` | 0.3.7 | `u24.x86_64` | pigsty | 10.6 MiB | [postgresql-17-pg-analytics_0.3.7_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-17-pg-analytics/postgresql-17-pg-analytics_0.3.7_amd64.deb) |
| `postgresql-17-pg-analytics` | 0.3.7 | `u24.aarch64` | pigsty | 9.2 MiB | [postgresql-17-pg-analytics_0.3.7_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-17-pg-analytics/postgresql-17-pg-analytics_0.3.7_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_analytics_16` | 0.3.7 | `el8.x86_64` | pigsty | 11.1 MiB | [pg_analytics_16-0.3.7-1PARADEDB.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_analytics_16-0.3.7-1PARADEDB.el8.x86_64.rpm) |
| `pg_analytics_16` | 0.3.7 | `el8.aarch64` | pigsty | 9.7 MiB | [pg_analytics_16-0.3.7-1PARADEDB.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_analytics_16-0.3.7-1PARADEDB.el8.aarch64.rpm) |
| `pg_analytics_16` | 0.3.7 | `el9.x86_64` | pigsty | 11.1 MiB | [pg_analytics_16-0.3.7-1PARADEDB.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_analytics_16-0.3.7-1PARADEDB.el9.x86_64.rpm) |
| `pg_analytics_16` | 0.3.7 | `el9.aarch64` | pigsty | 10.2 MiB | [pg_analytics_16-0.3.7-1PARADEDB.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_analytics_16-0.3.7-1PARADEDB.el9.aarch64.rpm) |
| `postgresql-16-pg-analytics` | 0.3.7 | `d12.x86_64` | pigsty | 10.2 MiB | [postgresql-16-pg-analytics_0.3.7_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-16-pg-analytics/postgresql-16-pg-analytics_0.3.7_amd64.deb) |
| `postgresql-16-pg-analytics` | 0.3.7 | `d12.aarch64` | pigsty | 8.8 MiB | [postgresql-16-pg-analytics_0.3.7_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-16-pg-analytics/postgresql-16-pg-analytics_0.3.7_arm64.deb) |
| `postgresql-16-pg-analytics` | 0.3.7 | `u22.x86_64` | pigsty | 10.3 MiB | [postgresql-16-pg-analytics_0.3.7_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-16-pg-analytics/postgresql-16-pg-analytics_0.3.7_amd64.deb) |
| `postgresql-16-pg-analytics` | 0.3.7 | `u22.aarch64` | pigsty | 9.0 MiB | [postgresql-16-pg-analytics_0.3.7_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-16-pg-analytics/postgresql-16-pg-analytics_0.3.7_arm64.deb) |
| `postgresql-16-pg-analytics` | 0.3.7 | `u24.x86_64` | pigsty | 10.6 MiB | [postgresql-16-pg-analytics_0.3.7_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-16-pg-analytics/postgresql-16-pg-analytics_0.3.7_amd64.deb) |
| `postgresql-16-pg-analytics` | 0.3.7 | `u24.aarch64` | pigsty | 9.2 MiB | [postgresql-16-pg-analytics_0.3.7_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-16-pg-analytics/postgresql-16-pg-analytics_0.3.7_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_analytics_15` | 0.3.7 | `el8.x86_64` | pigsty | 11.1 MiB | [pg_analytics_15-0.3.7-1PARADEDB.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_analytics_15-0.3.7-1PARADEDB.el8.x86_64.rpm) |
| `pg_analytics_15` | 0.3.7 | `el8.aarch64` | pigsty | 9.7 MiB | [pg_analytics_15-0.3.7-1PARADEDB.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_analytics_15-0.3.7-1PARADEDB.el8.aarch64.rpm) |
| `pg_analytics_15` | 0.3.7 | `el9.x86_64` | pigsty | 11.1 MiB | [pg_analytics_15-0.3.7-1PARADEDB.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_analytics_15-0.3.7-1PARADEDB.el9.x86_64.rpm) |
| `pg_analytics_15` | 0.3.7 | `el9.aarch64` | pigsty | 10.2 MiB | [pg_analytics_15-0.3.7-1PARADEDB.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_analytics_15-0.3.7-1PARADEDB.el9.aarch64.rpm) |
| `postgresql-15-pg-analytics` | 0.3.7 | `d12.x86_64` | pigsty | 10.2 MiB | [postgresql-15-pg-analytics_0.3.7_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-15-pg-analytics/postgresql-15-pg-analytics_0.3.7_amd64.deb) |
| `postgresql-15-pg-analytics` | 0.3.7 | `d12.aarch64` | pigsty | 8.8 MiB | [postgresql-15-pg-analytics_0.3.7_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-15-pg-analytics/postgresql-15-pg-analytics_0.3.7_arm64.deb) |
| `postgresql-15-pg-analytics` | 0.3.7 | `u22.x86_64` | pigsty | 10.4 MiB | [postgresql-15-pg-analytics_0.3.7_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-15-pg-analytics/postgresql-15-pg-analytics_0.3.7_amd64.deb) |
| `postgresql-15-pg-analytics` | 0.3.7 | `u22.aarch64` | pigsty | 9.0 MiB | [postgresql-15-pg-analytics_0.3.7_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-15-pg-analytics/postgresql-15-pg-analytics_0.3.7_arm64.deb) |
| `postgresql-15-pg-analytics` | 0.3.7 | `u24.x86_64` | pigsty | 10.5 MiB | [postgresql-15-pg-analytics_0.3.7_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-15-pg-analytics/postgresql-15-pg-analytics_0.3.7_amd64.deb) |
| `postgresql-15-pg-analytics` | 0.3.7 | `u24.aarch64` | pigsty | 9.2 MiB | [postgresql-15-pg-analytics_0.3.7_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-15-pg-analytics/postgresql-15-pg-analytics_0.3.7_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_analytics_14` | 0.3.7 | `el8.x86_64` | pigsty | 11.1 MiB | [pg_analytics_14-0.3.7-1PARADEDB.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_analytics_14-0.3.7-1PARADEDB.el8.x86_64.rpm) |
| `pg_analytics_14` | 0.3.7 | `el8.aarch64` | pigsty | 9.7 MiB | [pg_analytics_14-0.3.7-1PARADEDB.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_analytics_14-0.3.7-1PARADEDB.el8.aarch64.rpm) |
| `pg_analytics_14` | 0.3.7 | `el9.x86_64` | pigsty | 11.1 MiB | [pg_analytics_14-0.3.7-1PARADEDB.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_analytics_14-0.3.7-1PARADEDB.el9.x86_64.rpm) |
| `pg_analytics_14` | 0.3.7 | `el9.aarch64` | pigsty | 10.2 MiB | [pg_analytics_14-0.3.7-1PARADEDB.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_analytics_14-0.3.7-1PARADEDB.el9.aarch64.rpm) |
| `postgresql-14-pg-analytics` | 0.3.7 | `d12.x86_64` | pigsty | 10.2 MiB | [postgresql-14-pg-analytics_0.3.7_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-14-pg-analytics/postgresql-14-pg-analytics_0.3.7_amd64.deb) |
| `postgresql-14-pg-analytics` | 0.3.7 | `d12.aarch64` | pigsty | 8.8 MiB | [postgresql-14-pg-analytics_0.3.7_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-14-pg-analytics/postgresql-14-pg-analytics_0.3.7_arm64.deb) |
| `postgresql-14-pg-analytics` | 0.3.7 | `u22.x86_64` | pigsty | 10.3 MiB | [postgresql-14-pg-analytics_0.3.7_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-14-pg-analytics/postgresql-14-pg-analytics_0.3.7_amd64.deb) |
| `postgresql-14-pg-analytics` | 0.3.7 | `u22.aarch64` | pigsty | 9.0 MiB | [postgresql-14-pg-analytics_0.3.7_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-14-pg-analytics/postgresql-14-pg-analytics_0.3.7_arm64.deb) |
| `postgresql-14-pg-analytics` | 0.3.7 | `u24.x86_64` | pigsty | 10.6 MiB | [postgresql-14-pg-analytics_0.3.7_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-14-pg-analytics/postgresql-14-pg-analytics_0.3.7_amd64.deb) |
| `postgresql-14-pg-analytics` | 0.3.7 | `u24.aarch64` | pigsty | 9.1 MiB | [postgresql-14-pg-analytics_0.3.7_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-14-pg-analytics/postgresql-14-pg-analytics_0.3.7_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/paradedb/pg_analytics" title="Repository" icon="github" subtitle="github.com/paradedb/pg_analytics" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_analytics; # install by extension name, for the current active PG version
pig ext install pg_analytics; # install via package alias, for the active PG version
pig ext install pg_analytics -v 17;   # install for PG 17
pig ext install pg_analytics -v 16;   # install for PG 16
pig ext install pg_analytics -v 15;   # install for PG 15
pig ext install pg_analytics -v 14;   # install for PG 14

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_analytics CASCADE SCHEMA paradedb;
```




--------

## Usage

<Callout title="Archived!" type="warning">

This extension is archived and no longer maintained. It is recommended to use other duckdb related extensions instead.

</Callout>


https://github.com/paradedb/pg_analytics

Example, read parquet files from S3:

```bash
CREATE EXTENSION pg_lakehouse;
CREATE FOREIGN DATA WRAPPER parquet_wrapper HANDLER parquet_fdw_handler VALIDATOR parquet_fdw_validator;

-- Provide S3 credentials
CREATE SERVER parquet_server FOREIGN DATA WRAPPER parquet_wrapper;

-- Create foreign table with auto schema creation
CREATE FOREIGN TABLE trips ()
SERVER parquet_server
OPTIONS (files 's3://paradedb-benchmarks/yellow_tripdata_2024-01.parquet');

-- Success! Now you can query the remote Parquet file like a regular Postgres table
SELECT COUNT(*) FROM trips;
  count
---------
 2964624
(1 row)
```

This fdw is read-only for now.



----

## Iceberg Support

```sql
CREATE EXTENSION pg_lakehouse;

CREATE FOREIGN DATA WRAPPER iceberg_wrapper
    HANDLER iceberg_fdw_handler
    VALIDATOR iceberg_fdw_validator;

CREATE SERVER iceberg_server
    FOREIGN DATA WRAPPER iceberg_wrapper;

-- Replace the dummy schema with the actual schema
CREATE FOREIGN TABLE iceberg_table (x INT)
    SERVER iceberg_server
    OPTIONS (files 's3://bucket/iceberg_folder');

-- Success! You can now query the Iceberg table
SELECT COUNT(*) FROM iceberg_table;
```

