---
title: "pg_clickhouse"
linkTitle: "pg_clickhouse"
description: "Interfaces to query ClickHouse databases from PostgreSQL"
weight: 2460
categories: ["OLAP"]
languages: ["C++"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_clickhouse**](https://github.com/ClickHouse/pg_clickhouse) : Interfaces to query ClickHouse databases from PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2460** | {{< badge content="pg_clickhouse" link="https://github.com/ClickHouse/pg_clickhouse" >}} | {{< ext "pg_clickhouse" >}} | `0.10.0` | {{< category "OLAP" >}} | {{< license "Apache-2.0" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_stat_ch" >}} {{< ext "duckdb_fdw" >}} {{< ext "pg_duckdb" >}} {{< ext "pg_mooncake" >}} {{< ext "pg_ducklake" >}} {{< ext "pg_lake" >}} {{< ext "hdfs_fdw" >}} {{< ext "kafka_fdw" >}} {{< ext "aws_s3" >}} {{< ext "pg_parquet" >}} |

> [!Note] Release v0.10.0, control SQL version 0.10; preloading is optional; no llvmjit subpackage on el9.x86_64 in the 2026-08-12 build.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.10.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_clickhouse` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.10.0` | {{< bg "18" "pg_clickhouse_18" "green" >}} {{< bg "17" "pg_clickhouse_17" "green" >}} {{< bg "16" "pg_clickhouse_16" "green" >}} {{< bg "15" "pg_clickhouse_15" "green" >}} {{< bg "14" "pg_clickhouse_14" "green" >}} | `pg_clickhouse_$v` | `openssl`, `libcurl`, `libuuid`, `lz4-libs`, `libzstd` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.10.0` | {{< bg "18" "postgresql-18-clickhouse" "green" >}} {{< bg "17" "postgresql-17-clickhouse" "green" >}} {{< bg "16" "postgresql-16-clickhouse" "green" >}} {{< bg "15" "postgresql-15-clickhouse" "green" >}} {{< bg "14" "postgresql-14-clickhouse" "green" >}} | `postgresql-$v-clickhouse` | `libssl3 | libssl3t64`, `libcurl4 | libcurl4t64`, `libuuid1`, `liblz4-1`, `libzstd1` |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.10.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_18` | `0.10.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 169.8 KiB | [pg_clickhouse_18-0.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_18-0.10.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_18` | `0.10.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 167.9 KiB | [pg_clickhouse_18-0.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_18-0.10.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_18` | `0.10.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 165.7 KiB | [pg_clickhouse_18-0.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_18-0.10.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_18` | `0.10.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 162.4 KiB | [pg_clickhouse_18-0.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_18-0.10.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_18` | `0.10.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 162.8 KiB | [pg_clickhouse_18-0.10.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_18-0.10.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_18` | `0.10.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 165.2 KiB | [pg_clickhouse_18-0.10.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_18-0.10.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-clickhouse` | `0.10.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 446.5 KiB | [postgresql-18-clickhouse_0.10.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.10.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-clickhouse` | `0.10.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 439.9 KiB | [postgresql-18-clickhouse_0.10.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.10.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-clickhouse` | `0.10.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 447.6 KiB | [postgresql-18-clickhouse_0.10.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.10.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-clickhouse` | `0.10.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 442.4 KiB | [postgresql-18-clickhouse_0.10.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.10.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-clickhouse` | `0.10.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 467.8 KiB | [postgresql-18-clickhouse_0.10.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.10.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-clickhouse` | `0.10.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 468.8 KiB | [postgresql-18-clickhouse_0.10.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.10.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-clickhouse` | `0.10.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 446.7 KiB | [postgresql-18-clickhouse_0.10.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.10.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-clickhouse` | `0.10.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 449.1 KiB | [postgresql-18-clickhouse_0.10.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.10.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-clickhouse` | `0.10.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 444.4 KiB | [postgresql-18-clickhouse_0.10.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.10.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-clickhouse` | `0.10.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 445.8 KiB | [postgresql-18-clickhouse_0.10.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.10.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_17` | `0.10.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 169.7 KiB | [pg_clickhouse_17-0.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_17-0.10.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_17` | `0.10.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 167.9 KiB | [pg_clickhouse_17-0.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_17-0.10.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_17` | `0.10.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 165.6 KiB | [pg_clickhouse_17-0.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_17-0.10.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_17` | `0.10.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 162.3 KiB | [pg_clickhouse_17-0.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_17-0.10.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_17` | `0.10.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 162.8 KiB | [pg_clickhouse_17-0.10.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_17-0.10.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_17` | `0.10.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 164.9 KiB | [pg_clickhouse_17-0.10.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_17-0.10.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-clickhouse` | `0.10.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 446.6 KiB | [postgresql-17-clickhouse_0.10.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.10.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-clickhouse` | `0.10.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 439.8 KiB | [postgresql-17-clickhouse_0.10.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.10.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-clickhouse` | `0.10.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 447.4 KiB | [postgresql-17-clickhouse_0.10.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.10.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-clickhouse` | `0.10.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 442.0 KiB | [postgresql-17-clickhouse_0.10.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.10.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-clickhouse` | `0.10.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 524.4 KiB | [postgresql-17-clickhouse_0.10.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.10.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-clickhouse` | `0.10.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 525.1 KiB | [postgresql-17-clickhouse_0.10.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.10.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-clickhouse` | `0.10.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 446.6 KiB | [postgresql-17-clickhouse_0.10.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.10.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-clickhouse` | `0.10.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 448.6 KiB | [postgresql-17-clickhouse_0.10.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.10.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-clickhouse` | `0.10.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 444.1 KiB | [postgresql-17-clickhouse_0.10.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.10.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-clickhouse` | `0.10.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 445.6 KiB | [postgresql-17-clickhouse_0.10.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.10.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_16` | `0.10.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 169.6 KiB | [pg_clickhouse_16-0.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_16-0.10.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_16` | `0.10.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 167.8 KiB | [pg_clickhouse_16-0.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_16-0.10.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_16` | `0.10.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 165.4 KiB | [pg_clickhouse_16-0.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_16-0.10.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_16` | `0.10.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 162.1 KiB | [pg_clickhouse_16-0.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_16-0.10.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_16` | `0.10.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 162.2 KiB | [pg_clickhouse_16-0.10.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_16-0.10.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_16` | `0.10.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 164.8 KiB | [pg_clickhouse_16-0.10.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_16-0.10.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-clickhouse` | `0.10.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 446.2 KiB | [postgresql-16-clickhouse_0.10.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.10.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-clickhouse` | `0.10.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 439.7 KiB | [postgresql-16-clickhouse_0.10.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.10.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-clickhouse` | `0.10.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 447.1 KiB | [postgresql-16-clickhouse_0.10.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.10.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-clickhouse` | `0.10.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 441.6 KiB | [postgresql-16-clickhouse_0.10.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.10.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-clickhouse` | `0.10.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 519.3 KiB | [postgresql-16-clickhouse_0.10.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.10.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-clickhouse` | `0.10.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 520.3 KiB | [postgresql-16-clickhouse_0.10.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.10.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-clickhouse` | `0.10.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 446.2 KiB | [postgresql-16-clickhouse_0.10.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.10.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-clickhouse` | `0.10.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 448.2 KiB | [postgresql-16-clickhouse_0.10.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.10.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-clickhouse` | `0.10.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 443.6 KiB | [postgresql-16-clickhouse_0.10.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.10.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-clickhouse` | `0.10.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 445.2 KiB | [postgresql-16-clickhouse_0.10.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.10.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_15` | `0.10.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 173.3 KiB | [pg_clickhouse_15-0.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_15-0.10.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_15` | `0.10.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 170.9 KiB | [pg_clickhouse_15-0.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_15-0.10.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_15` | `0.10.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 168.5 KiB | [pg_clickhouse_15-0.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_15-0.10.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_15` | `0.10.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 170.9 KiB | [pg_clickhouse_15-0.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_15-0.10.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_15` | `0.10.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 171.1 KiB | [pg_clickhouse_15-0.10.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_15-0.10.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_15` | `0.10.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 173.8 KiB | [pg_clickhouse_15-0.10.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_15-0.10.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-clickhouse` | `0.10.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 449.7 KiB | [postgresql-15-clickhouse_0.10.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.10.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-clickhouse` | `0.10.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 442.3 KiB | [postgresql-15-clickhouse_0.10.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.10.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-clickhouse` | `0.10.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 450.6 KiB | [postgresql-15-clickhouse_0.10.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.10.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-clickhouse` | `0.10.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 444.9 KiB | [postgresql-15-clickhouse_0.10.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.10.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-clickhouse` | `0.10.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 526.5 KiB | [postgresql-15-clickhouse_0.10.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.10.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-clickhouse` | `0.10.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 526.7 KiB | [postgresql-15-clickhouse_0.10.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.10.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-clickhouse` | `0.10.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 453.8 KiB | [postgresql-15-clickhouse_0.10.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.10.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-clickhouse` | `0.10.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 455.4 KiB | [postgresql-15-clickhouse_0.10.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.10.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-clickhouse` | `0.10.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 450.1 KiB | [postgresql-15-clickhouse_0.10.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.10.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-clickhouse` | `0.10.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 452.9 KiB | [postgresql-15-clickhouse_0.10.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.10.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_14` | `0.10.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 173.1 KiB | [pg_clickhouse_14-0.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_14-0.10.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_14` | `0.10.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 170.9 KiB | [pg_clickhouse_14-0.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_14-0.10.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_14` | `0.10.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 168.5 KiB | [pg_clickhouse_14-0.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_14-0.10.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_14` | `0.10.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 170.9 KiB | [pg_clickhouse_14-0.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_14-0.10.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_14` | `0.10.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 171.2 KiB | [pg_clickhouse_14-0.10.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_14-0.10.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_14` | `0.10.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 173.7 KiB | [pg_clickhouse_14-0.10.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_14-0.10.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-clickhouse` | `0.10.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 449.6 KiB | [postgresql-14-clickhouse_0.10.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.10.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-clickhouse` | `0.10.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 442.5 KiB | [postgresql-14-clickhouse_0.10.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.10.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-clickhouse` | `0.10.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 450.5 KiB | [postgresql-14-clickhouse_0.10.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.10.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-clickhouse` | `0.10.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 445.0 KiB | [postgresql-14-clickhouse_0.10.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.10.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-clickhouse` | `0.10.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 526.4 KiB | [postgresql-14-clickhouse_0.10.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.10.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-clickhouse` | `0.10.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 526.6 KiB | [postgresql-14-clickhouse_0.10.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.10.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-clickhouse` | `0.10.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 453.7 KiB | [postgresql-14-clickhouse_0.10.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.10.0-1PGSTY~noble_amd64.deb) |
| `postgresql-14-clickhouse` | `0.10.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 455.4 KiB | [postgresql-14-clickhouse_0.10.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.10.0-1PGSTY~noble_arm64.deb) |
| `postgresql-14-clickhouse` | `0.10.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 450.0 KiB | [postgresql-14-clickhouse_0.10.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.10.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-clickhouse` | `0.10.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 452.8 KiB | [postgresql-14-clickhouse_0.10.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.10.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ClickHouse/pg_clickhouse" title="Repository" icon="github" subtitle="github.com/ClickHouse/pg_clickhouse" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_clickhouse-0.10.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_clickhouse;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_clickhouse;		# install via package name, for the active PG version

pig install pg_clickhouse -v 18;   # install for PG 18
pig install pg_clickhouse -v 17;   # install for PG 17
pig install pg_clickhouse -v 16;   # install for PG 16
pig install pg_clickhouse -v 15;   # install for PG 15
pig install pg_clickhouse -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_clickhouse;
```

## Usage

Sources:

- [pg_clickhouse v0.10.0 README](https://github.com/ClickHouse/pg_clickhouse/blob/v0.10.0/README.md)
- [pg_clickhouse v0.10.0 reference](https://github.com/ClickHouse/pg_clickhouse/blob/v0.10.0/doc/pg_clickhouse.md)
- [pg_clickhouse v0.10.0 tutorial](https://github.com/ClickHouse/pg_clickhouse/blob/v0.10.0/doc/tutorial.md)
- [pg_clickhouse v0.10.0 changelog](https://github.com/ClickHouse/pg_clickhouse/blob/v0.10.0/CHANGELOG.md)
- [pg_clickhouse v0.10.0 control file](https://github.com/ClickHouse/pg_clickhouse/blob/v0.10.0/pg_clickhouse.control)
- [pg_clickhouse 0.3 to 0.10 upgrade SQL](https://github.com/ClickHouse/pg_clickhouse/blob/v0.10.0/sql/pg_clickhouse--0.3--0.10.sql)
- [Pigsty pg_clickhouse package matrix](https://pgext.cloud/ext/pg_clickhouse)

`pg_clickhouse` 0.10.0 exposes ClickHouse tables to PostgreSQL through the `clickhouse_fdw` foreign data wrapper. Upstream targets PostgreSQL 13 or later and ClickHouse 23.3 or later; current Pigsty packages cover PostgreSQL 14–18. No preload is required for normal use; `session_preload_libraries` and `shared_preload_libraries` are optional connection-startup optimizations.

### Connect PostgreSQL to ClickHouse

```sql
CREATE EXTENSION pg_clickhouse;

CREATE SERVER taxi_srv
FOREIGN DATA WRAPPER clickhouse_fdw
OPTIONS (
  driver 'binary',
  host 'localhost',
  dbname 'taxi',
  compression 'lz4'
);

CREATE USER MAPPING FOR CURRENT_USER
SERVER taxi_srv
OPTIONS (user 'default');

CREATE SCHEMA taxi;
IMPORT FOREIGN SCHEMA taxi FROM SERVER taxi_srv INTO taxi;
```

The required `driver` option is `binary` or `http`. Common server options include `host`, `port`, `dbname`, `compression`, `secure`, and `min_tls_version`; user mappings accept `user` and `password`. Version 0.10 deprecates and ignores `fetch_size` because both drivers now stream the same Native format.

`IMPORT FOREIGN SCHEMA` supports `LIMIT TO (...)` and `EXCEPT (...)`. Imported mixed-case identifiers remain quoted and must be referenced with matching quotes.

### Query and Write Foreign Tables

```sql
EXPLAIN (VERBOSE)
SELECT node_id, count(*)
FROM taxi.logs
GROUP BY node_id;

INSERT INTO taxi.nodes(node_id, name)
VALUES (9, 'west-node');

COPY taxi.nodes(node_id, name) FROM STDIN;
```

`SELECT`, `EXPLAIN`, prepared statements, `INSERT`, and `COPY` operate on foreign tables. In version 0.10 the binary driver flushes inserts in bounded 64 MiB batches, so `COPY` is no longer merely expanded into one statement per row. Use `EXPLAIN (VERBOSE)` to inspect remote SQL and verify which filters, joins, aggregates, and functions were pushed down.

### Direct Query and Command APIs

Version 0.10 adds typed arbitrary-query and command interfaces:

```sql
GRANT EXECUTE ON FUNCTION clickhouse_query(text, text) TO analyst;
GRANT EXECUTE ON PROCEDURE clickhouse_perform(text, text) TO operator;

SELECT *
FROM clickhouse_query(
  'taxi_srv',
  'SELECT region, count() FROM taxi GROUP BY region'
) AS t(region text, n bigint);

CALL clickhouse_perform(
  'taxi_srv',
  'OPTIMIZE TABLE taxi.nodes FINAL'
);

SELECT clickhouse_server_version('taxi_srv');
```

`clickhouse_query(server, sql)` returns rows using the caller-provided column definition, while `clickhouse_perform(server, sql)` discards any result. Both can run arbitrary remote SQL, so `EXECUTE` is revoked from `PUBLIC` and should be granted narrowly. `clickhouse_raw_query()` is deprecated in favor of these interfaces.

### Pushdown and Session Settings

Version 0.10 expands aggregate and function pushdown, improves aggregate execution over mixed local and foreign partitions, and fixes several PostgreSQL NULL-semantics mismatches. Subquery pushdown requires ClickHouse 25.8 or later; older servers evaluate those subqueries locally.

The default `pg_clickhouse.session_settings` preserves PostgreSQL-compatible behavior, including `join_use_nulls = 1`, `group_by_use_nulls = 1`, `final = 1`, and `transform_null_in = 0`. If it is overridden, retain the settings needed by the workload—especially `transform_null_in = 0`, which is required for safe `IN` pushdown.

### Upgrade and Operational Boundaries

```sql
ALTER EXTENSION pg_clickhouse UPDATE TO '0.10';
SELECT pgch_version();
```

The extension SQL version is `0.10`, while `pgch_version()` reports the full library version `0.10.0`. An installation upgraded from SQL version `0.3` must run `ALTER EXTENSION` after the new files are installed.

If `pg_clickhouse` is placed in `session_preload_libraries`, new sessions load it automatically. If it is placed in `shared_preload_libraries`, changing the library requires a PostgreSQL restart. Neither setting is mandatory, unlike extensions that register postmaster hooks.

Lightweight `UPDATE` and `DELETE` remain outside the documented write surface. Treat direct remote SQL as privileged, test pushdown with production-shaped NULL and type cases, and validate both PostgreSQL and ClickHouse versions before relying on a version-gated optimization.
