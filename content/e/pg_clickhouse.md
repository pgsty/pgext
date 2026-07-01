---
title: "pg_clickhouse"
linkTitle: "pg_clickhouse"
description: "Interfaces to query ClickHouse databases from PostgreSQL"
weight: 2460
categories: ["OLAP"]
width: full
---

[**pg_clickhouse**](https://github.com/ClickHouse/pg_clickhouse) : Interfaces to query ClickHouse databases from PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2460** | {{< badge content="pg_clickhouse" link="https://github.com/ClickHouse/pg_clickhouse" >}} | {{< ext "pg_clickhouse" >}} | `0.3.2` | {{< category "OLAP" >}} | {{< license "Apache-2.0" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_duckdb" >}} {{< ext "duckdb_fdw" >}} {{< ext "citus" >}} {{< ext "columnar" >}} {{< ext "citus_columnar" >}} {{< ext "clickhouse_fdw" >}} {{< ext "postgres_fdw" >}} {{< ext "dblink" >}} |

> [!Note] release 0.3.1; SQL v0.3


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_clickhouse` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.2` | {{< bg "18" "pg_clickhouse_18" "green" >}} {{< bg "17" "pg_clickhouse_17" "green" >}} {{< bg "16" "pg_clickhouse_16" "green" >}} {{< bg "15" "pg_clickhouse_15" "green" >}} {{< bg "14" "pg_clickhouse_14" "green" >}} | `pg_clickhouse_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.2` | {{< bg "18" "postgresql-18-clickhouse" "green" >}} {{< bg "17" "postgresql-17-clickhouse" "green" >}} {{< bg "16" "postgresql-16-clickhouse" "green" >}} {{< bg "15" "postgresql-15-clickhouse" "green" >}} {{< bg "14" "postgresql-14-clickhouse" "green" >}} | `postgresql-$v-clickhouse` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.2" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_18` | `0.3.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 147.6 KiB | [pg_clickhouse_18-0.3.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_18-0.3.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_18` | `0.3.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 146.6 KiB | [pg_clickhouse_18-0.3.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_18-0.3.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_18` | `0.3.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 143.8 KiB | [pg_clickhouse_18-0.3.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_18-0.3.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_18` | `0.3.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 142.0 KiB | [pg_clickhouse_18-0.3.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_18-0.3.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_18` | `0.3.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 141.1 KiB | [pg_clickhouse_18-0.3.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_18-0.3.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_18` | `0.3.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 144.7 KiB | [pg_clickhouse_18-0.3.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_18-0.3.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-clickhouse` | `0.3.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 392.4 KiB | [postgresql-18-clickhouse_0.3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-clickhouse` | `0.3.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 387.6 KiB | [postgresql-18-clickhouse_0.3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-clickhouse` | `0.3.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 392.8 KiB | [postgresql-18-clickhouse_0.3.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.3.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-clickhouse` | `0.3.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 389.1 KiB | [postgresql-18-clickhouse_0.3.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.3.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-clickhouse` | `0.3.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 415.1 KiB | [postgresql-18-clickhouse_0.3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-clickhouse` | `0.3.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 417.3 KiB | [postgresql-18-clickhouse_0.3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-clickhouse` | `0.3.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 396.5 KiB | [postgresql-18-clickhouse_0.3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-clickhouse` | `0.3.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 399.0 KiB | [postgresql-18-clickhouse_0.3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.3.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-clickhouse` | `0.3.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 392.8 KiB | [postgresql-18-clickhouse_0.3.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.3.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-clickhouse` | `0.3.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 394.9 KiB | [postgresql-18-clickhouse_0.3.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.3.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_17` | `0.3.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 147.4 KiB | [pg_clickhouse_17-0.3.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_17-0.3.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_17` | `0.3.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 146.3 KiB | [pg_clickhouse_17-0.3.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_17-0.3.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_17` | `0.3.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 143.6 KiB | [pg_clickhouse_17-0.3.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_17-0.3.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_17` | `0.3.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 141.9 KiB | [pg_clickhouse_17-0.3.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_17-0.3.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_17` | `0.3.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 140.9 KiB | [pg_clickhouse_17-0.3.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_17-0.3.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_17` | `0.3.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 144.5 KiB | [pg_clickhouse_17-0.3.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_17-0.3.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-clickhouse` | `0.3.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 392.2 KiB | [postgresql-17-clickhouse_0.3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-clickhouse` | `0.3.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 387.1 KiB | [postgresql-17-clickhouse_0.3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-clickhouse` | `0.3.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 392.3 KiB | [postgresql-17-clickhouse_0.3.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.3.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-clickhouse` | `0.3.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 388.9 KiB | [postgresql-17-clickhouse_0.3.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.3.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-clickhouse` | `0.3.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 465.6 KiB | [postgresql-17-clickhouse_0.3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-clickhouse` | `0.3.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 466.8 KiB | [postgresql-17-clickhouse_0.3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-clickhouse` | `0.3.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 396.2 KiB | [postgresql-17-clickhouse_0.3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-clickhouse` | `0.3.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 398.8 KiB | [postgresql-17-clickhouse_0.3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.3.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-clickhouse` | `0.3.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 392.1 KiB | [postgresql-17-clickhouse_0.3.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.3.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-clickhouse` | `0.3.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 394.3 KiB | [postgresql-17-clickhouse_0.3.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.3.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_16` | `0.3.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 147.4 KiB | [pg_clickhouse_16-0.3.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_16-0.3.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_16` | `0.3.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 146.3 KiB | [pg_clickhouse_16-0.3.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_16-0.3.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_16` | `0.3.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 143.4 KiB | [pg_clickhouse_16-0.3.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_16-0.3.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_16` | `0.3.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 141.8 KiB | [pg_clickhouse_16-0.3.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_16-0.3.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_16` | `0.3.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 140.5 KiB | [pg_clickhouse_16-0.3.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_16-0.3.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_16` | `0.3.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 144.4 KiB | [pg_clickhouse_16-0.3.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_16-0.3.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-clickhouse` | `0.3.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 391.7 KiB | [postgresql-16-clickhouse_0.3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-clickhouse` | `0.3.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 387.5 KiB | [postgresql-16-clickhouse_0.3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-clickhouse` | `0.3.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 392.0 KiB | [postgresql-16-clickhouse_0.3.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.3.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-clickhouse` | `0.3.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 388.9 KiB | [postgresql-16-clickhouse_0.3.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.3.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-clickhouse` | `0.3.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 460.9 KiB | [postgresql-16-clickhouse_0.3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-clickhouse` | `0.3.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 461.8 KiB | [postgresql-16-clickhouse_0.3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-clickhouse` | `0.3.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 395.6 KiB | [postgresql-16-clickhouse_0.3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-clickhouse` | `0.3.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 398.5 KiB | [postgresql-16-clickhouse_0.3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.3.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-clickhouse` | `0.3.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 391.9 KiB | [postgresql-16-clickhouse_0.3.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.3.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-clickhouse` | `0.3.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 394.4 KiB | [postgresql-16-clickhouse_0.3.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.3.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_15` | `0.3.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 151.3 KiB | [pg_clickhouse_15-0.3.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_15-0.3.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_15` | `0.3.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 150.0 KiB | [pg_clickhouse_15-0.3.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_15-0.3.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_15` | `0.3.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 147.3 KiB | [pg_clickhouse_15-0.3.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_15-0.3.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_15` | `0.3.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 150.7 KiB | [pg_clickhouse_15-0.3.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_15-0.3.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_15` | `0.3.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 149.5 KiB | [pg_clickhouse_15-0.3.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_15-0.3.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_15` | `0.3.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 153.3 KiB | [pg_clickhouse_15-0.3.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_15-0.3.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-clickhouse` | `0.3.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 396.1 KiB | [postgresql-15-clickhouse_0.3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-clickhouse` | `0.3.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 390.9 KiB | [postgresql-15-clickhouse_0.3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-clickhouse` | `0.3.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 396.6 KiB | [postgresql-15-clickhouse_0.3.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.3.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-clickhouse` | `0.3.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 393.5 KiB | [postgresql-15-clickhouse_0.3.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.3.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-clickhouse` | `0.3.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 468.6 KiB | [postgresql-15-clickhouse_0.3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-clickhouse` | `0.3.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 469.7 KiB | [postgresql-15-clickhouse_0.3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-clickhouse` | `0.3.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 402.0 KiB | [postgresql-15-clickhouse_0.3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-clickhouse` | `0.3.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 406.3 KiB | [postgresql-15-clickhouse_0.3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.3.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-clickhouse` | `0.3.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 398.5 KiB | [postgresql-15-clickhouse_0.3.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.3.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-clickhouse` | `0.3.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 401.5 KiB | [postgresql-15-clickhouse_0.3.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.3.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_14` | `0.3.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 151.2 KiB | [pg_clickhouse_14-0.3.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_14-0.3.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_14` | `0.3.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 149.9 KiB | [pg_clickhouse_14-0.3.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_14-0.3.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_14` | `0.3.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 147.3 KiB | [pg_clickhouse_14-0.3.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_14-0.3.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_14` | `0.3.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 150.5 KiB | [pg_clickhouse_14-0.3.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_14-0.3.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_14` | `0.3.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 149.4 KiB | [pg_clickhouse_14-0.3.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_14-0.3.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_14` | `0.3.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 153.2 KiB | [pg_clickhouse_14-0.3.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_14-0.3.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-clickhouse` | `0.3.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 395.9 KiB | [postgresql-14-clickhouse_0.3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-clickhouse` | `0.3.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 390.9 KiB | [postgresql-14-clickhouse_0.3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-clickhouse` | `0.3.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 396.2 KiB | [postgresql-14-clickhouse_0.3.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.3.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-clickhouse` | `0.3.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 392.4 KiB | [postgresql-14-clickhouse_0.3.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.3.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-clickhouse` | `0.3.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 468.0 KiB | [postgresql-14-clickhouse_0.3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-clickhouse` | `0.3.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 469.0 KiB | [postgresql-14-clickhouse_0.3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-clickhouse` | `0.3.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 401.6 KiB | [postgresql-14-clickhouse_0.3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-clickhouse` | `0.3.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 405.8 KiB | [postgresql-14-clickhouse_0.3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.3.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-clickhouse` | `0.3.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 398.2 KiB | [postgresql-14-clickhouse_0.3.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.3.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-clickhouse` | `0.3.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 401.4 KiB | [postgresql-14-clickhouse_0.3.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.3.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ClickHouse/pg_clickhouse" title="Repository" icon="github" subtitle="github.com/ClickHouse/pg_clickhouse" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_clickhouse-0.3.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_clickhouse;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_clickhouse;		# install via package name, for the active PG version

pig install pg_clickhouse -v 18;   # install for PG 18
pig install pg_clickhouse -v 17;   # install for PG 17
pig install pg_clickhouse -v 16;   # install for PG 16
pig install pg_clickhouse -v 15;   # install for PG 15
pig install pg_clickhouse -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_clickhouse';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_clickhouse;
```




## Usage

Sources: [README](https://github.com/ClickHouse/pg_clickhouse/blob/main/README.md), [reference](https://github.com/ClickHouse/pg_clickhouse/blob/main/doc/pg_clickhouse.md), [tutorial](https://github.com/ClickHouse/pg_clickhouse/blob/main/doc/tutorial.md), [v0.3.2 release notes](https://github.com/ClickHouse/pg_clickhouse/releases/tag/v0.3.2), [changelog](https://github.com/ClickHouse/pg_clickhouse/blob/main/CHANGELOG.md)

`pg_clickhouse` runs analytics queries on ClickHouse from PostgreSQL through the `clickhouse_fdw` foreign data wrapper. Upstream documents PostgreSQL 13+ and ClickHouse 23+ support; the current version is 0.3.2.

### Connect PostgreSQL to ClickHouse

```sql
CREATE EXTENSION pg_clickhouse;

CREATE SERVER taxi_srv
FOREIGN DATA WRAPPER clickhouse_fdw
OPTIONS (driver 'binary', host 'localhost', dbname 'taxi');

CREATE USER MAPPING FOR CURRENT_USER
SERVER taxi_srv
OPTIONS (user 'default');

CREATE SCHEMA taxi;
IMPORT FOREIGN SCHEMA taxi FROM SERVER taxi_srv INTO taxi;
```

Server options documented upstream:

- `driver`: required, `binary` or `http`
- `host`
- `port`
- `dbname`
- `fetch_size`: HTTP streaming batch size; `0` disables streaming
- `compression`: binary-driver compression, `none`, `lz4`, or `zstd`; v0.3.2 defaults to `lz4`
- `secure`: explicit TLS mode, `on`, `off`, or `auto`
- `min_tls_version`: minimum TLS protocol version for both binary and HTTP drivers

User mapping options:

- `user`
- `password`

### Common operations

```sql
ALTER EXTENSION pg_clickhouse UPDATE;
ALTER EXTENSION pg_clickhouse UPDATE TO '0.3';
SELECT pgch_version();
DROP SERVER taxi_srv CASCADE;
```

`IMPORT FOREIGN SCHEMA` also supports `LIMIT TO (...)` and `EXCEPT (...)`. The reference warns that imported mixed-case identifiers are double-quoted in PostgreSQL and must be queried with quotes.

### Query and write notes

`SELECT`, `EXPLAIN`, prepared statements, `INSERT`, and `COPY` can operate on `pg_clickhouse` foreign tables. Use `EXPLAIN (VERBOSE)` to inspect the remote SQL that will be sent to ClickHouse.

```sql
EXPLAIN (VERBOSE)
SELECT node_id, count(*)
FROM logs
GROUP BY node_id;

INSERT INTO nodes(node_id, name, region, arch, os)
VALUES (9, 'west-node', 'us-west-2', 'amd64', 'Linux');
```

`COPY` into a foreign table is documented, but upstream notes that it currently uses `INSERT` statements because FDW batch insertion is still future work.

### Version and pushdown notes

- The reference documents separate library and extension versions; `pgch_version()` reports the loaded library version.
- Release `v0.3.2` is binary-only for existing SQL version `0.3`; it does not require `ALTER EXTENSION UPDATE` when upgrading from another 0.3 build.
- Release `v0.3.2` adds the `compression`, `secure`, and `min_tls_version` server options, adds `regexp_match()` pushdown, and adds PostgreSQL 19beta1 distribution support.
- Release `v0.3.2` also fixes regular-expression flag pushdown and avoids pushing down regex calls when the regex argument is not a constant.
- Release `v0.3.1` is binary-only for existing SQL version `0.3`; it does not require `ALTER EXTENSION UPDATE` when upgrading from `v0.3.0`.
- Release `v0.3.1` replaces the client library with `ClickHouse/clickhouse-c`, streams result blocks, and adds rectangular multidimensional array support for `SELECT` and `INSERT`.
- Release `v0.3.1` also adds pushdown for `pg_re2` 0.3.0 functions such as `re2extractallgroupshorizontal`, `re2extractallgroupsvertical`, `re2regexpquotemeta`, and `re2splitbyregexp`, and fixes `UInt16` casts to PostgreSQL `int4`.
- Release `v0.3.0` uses SQL version `0.3`; run `ALTER EXTENSION pg_clickhouse UPDATE TO '0.3'` to apply its SQL-level privilege change.
- Release `v0.3.0` adds pushdown for `re2` functions, `soundex()`, two-argument `levenshtein()`, compatible `to_char(timestamp[tz], fmt)`, selected builtin functions, and JSON/JSONB path operations.
- ClickHouse `JSON` maps to PostgreSQL `jsonb` or `json`; the binary driver's `JSON` mapping requires ClickHouse 24.10 or later.
- `pg_clickhouse.pushdown_regex` controls built-in PostgreSQL regex pushdown. Upstream recommends considering the `re2` extension for regex work that should push down directly.

### Caveats

- In 0.3.0, `clickhouse_raw_query(text, text)` is no longer executable by `PUBLIC`; grant it only to roles that need ad-hoc ClickHouse queries.
- This is positioned upstream as an analytics-first extension; lightweight `DELETE` and `UPDATE` support remain on the roadmap.
- For full examples, follow the official tutorial, which creates a ClickHouse `taxi` database, imports it through `IMPORT FOREIGN SCHEMA`, and queries the resulting foreign tables.
