---
title: "pg_clickhouse"
linkTitle: "pg_clickhouse"
description: "Interfaces to query ClickHouse databases from PostgreSQL"
weight: 2460
categories: ["OLAP"]
languages: ["C"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_clickhouse**](https://github.com/ClickHouse/pg_clickhouse) : Interfaces to query ClickHouse databases from PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2460** | {{< badge content="pg_clickhouse" link="https://github.com/ClickHouse/pg_clickhouse" >}} | {{< ext "pg_clickhouse" >}} | `0.11` | {{< category "OLAP" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_stat_ch" >}} {{< ext "duckdb_fdw" >}} {{< ext "pg_duckdb" >}} {{< ext "pg_mooncake" >}} {{< ext "pg_ducklake" >}} {{< ext "pg_lake" >}} {{< ext "hdfs_fdw" >}} {{< ext "kafka_fdw" >}} {{< ext "aws_s3" >}} {{< ext "pg_parquet" >}} |

> [!Note] pkg 0.11.1; SQL 0.11; update drops clickhouse_raw_query(); fetch_size removed; encoding_check=fail.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.11` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_clickhouse` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.11.1` | {{< bg "18" "pg_clickhouse_18" "green" >}} {{< bg "17" "pg_clickhouse_17" "green" >}} {{< bg "16" "pg_clickhouse_16" "green" >}} {{< bg "15" "pg_clickhouse_15" "green" >}} {{< bg "14" "pg_clickhouse_14" "green" >}} | `pg_clickhouse_$v` | `openssl`, `libcurl`, `libuuid`, `lz4-libs`, `libzstd` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.11.1` | {{< bg "18" "postgresql-18-clickhouse" "green" >}} {{< bg "17" "postgresql-17-clickhouse" "green" >}} {{< bg "16" "postgresql-16-clickhouse" "green" >}} {{< bg "15" "postgresql-15-clickhouse" "green" >}} {{< bg "14" "postgresql-14-clickhouse" "green" >}} | `postgresql-$v-clickhouse` | `libssl3 | libssl3t64`, `libcurl4 | libcurl4t64`, `libuuid1`, `liblz4-1`, `libzstd1` |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.11.1" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_18` | `0.11.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 514.5 KiB | [pg_clickhouse_18-0.11.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_18-0.11.1-1PGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_18` | `0.11.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 509.9 KiB | [pg_clickhouse_18-0.11.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_18-0.11.1-1PGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_18` | `0.11.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 520.2 KiB | [pg_clickhouse_18-0.11.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_18-0.11.1-1PGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_18` | `0.11.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 515.5 KiB | [pg_clickhouse_18-0.11.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_18-0.11.1-1PGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_18` | `0.11.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 518.5 KiB | [pg_clickhouse_18-0.11.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_18-0.11.1-1PGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_18` | `0.11.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 519.5 KiB | [pg_clickhouse_18-0.11.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_18-0.11.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-clickhouse` | `0.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 518.1 KiB | [postgresql-18-clickhouse_0.11.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.11.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-clickhouse` | `0.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 510.8 KiB | [postgresql-18-clickhouse_0.11.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.11.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-clickhouse` | `0.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 520.1 KiB | [postgresql-18-clickhouse_0.11.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.11.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-clickhouse` | `0.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 513.6 KiB | [postgresql-18-clickhouse_0.11.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.11.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-clickhouse` | `0.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 540.7 KiB | [postgresql-18-clickhouse_0.11.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.11.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-clickhouse` | `0.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 541.4 KiB | [postgresql-18-clickhouse_0.11.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.11.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-clickhouse` | `0.11.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 515.8 KiB | [postgresql-18-clickhouse_0.11.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.11.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-clickhouse` | `0.11.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 516.9 KiB | [postgresql-18-clickhouse_0.11.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.11.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-clickhouse` | `0.11.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 512.4 KiB | [postgresql-18-clickhouse_0.11.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.11.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-clickhouse` | `0.11.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 513.6 KiB | [postgresql-18-clickhouse_0.11.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.11.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_17` | `0.11.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 514.1 KiB | [pg_clickhouse_17-0.11.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_17-0.11.1-1PGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_17` | `0.11.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 509.3 KiB | [pg_clickhouse_17-0.11.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_17-0.11.1-1PGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_17` | `0.11.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 519.4 KiB | [pg_clickhouse_17-0.11.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_17-0.11.1-1PGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_17` | `0.11.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 515.0 KiB | [pg_clickhouse_17-0.11.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_17-0.11.1-1PGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_17` | `0.11.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 517.8 KiB | [pg_clickhouse_17-0.11.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_17-0.11.1-1PGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_17` | `0.11.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 518.9 KiB | [pg_clickhouse_17-0.11.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_17-0.11.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-clickhouse` | `0.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 517.4 KiB | [postgresql-17-clickhouse_0.11.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.11.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-clickhouse` | `0.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 509.5 KiB | [postgresql-17-clickhouse_0.11.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.11.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-clickhouse` | `0.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 519.6 KiB | [postgresql-17-clickhouse_0.11.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.11.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-clickhouse` | `0.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 513.4 KiB | [postgresql-17-clickhouse_0.11.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.11.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-clickhouse` | `0.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 606.5 KiB | [postgresql-17-clickhouse_0.11.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.11.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-clickhouse` | `0.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 606.8 KiB | [postgresql-17-clickhouse_0.11.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.11.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-clickhouse` | `0.11.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 514.7 KiB | [postgresql-17-clickhouse_0.11.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.11.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-clickhouse` | `0.11.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 517.0 KiB | [postgresql-17-clickhouse_0.11.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.11.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-clickhouse` | `0.11.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 512.4 KiB | [postgresql-17-clickhouse_0.11.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.11.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-clickhouse` | `0.11.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 513.3 KiB | [postgresql-17-clickhouse_0.11.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.11.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_16` | `0.11.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 511.7 KiB | [pg_clickhouse_16-0.11.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_16-0.11.1-1PGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_16` | `0.11.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 507.2 KiB | [pg_clickhouse_16-0.11.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_16-0.11.1-1PGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_16` | `0.11.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 517.6 KiB | [pg_clickhouse_16-0.11.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_16-0.11.1-1PGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_16` | `0.11.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 513.4 KiB | [pg_clickhouse_16-0.11.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_16-0.11.1-1PGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_16` | `0.11.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 516.1 KiB | [pg_clickhouse_16-0.11.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_16-0.11.1-1PGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_16` | `0.11.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 517.5 KiB | [pg_clickhouse_16-0.11.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_16-0.11.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-clickhouse` | `0.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 515.9 KiB | [postgresql-16-clickhouse_0.11.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.11.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-clickhouse` | `0.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 508.5 KiB | [postgresql-16-clickhouse_0.11.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.11.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-clickhouse` | `0.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 516.9 KiB | [postgresql-16-clickhouse_0.11.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.11.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-clickhouse` | `0.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 510.8 KiB | [postgresql-16-clickhouse_0.11.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.11.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-clickhouse` | `0.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 599.3 KiB | [postgresql-16-clickhouse_0.11.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.11.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-clickhouse` | `0.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 599.2 KiB | [postgresql-16-clickhouse_0.11.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.11.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-clickhouse` | `0.11.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 512.8 KiB | [postgresql-16-clickhouse_0.11.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.11.1-1PGSTY~noble_amd64.deb) |
| `postgresql-16-clickhouse` | `0.11.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 514.9 KiB | [postgresql-16-clickhouse_0.11.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.11.1-1PGSTY~noble_arm64.deb) |
| `postgresql-16-clickhouse` | `0.11.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 510.1 KiB | [postgresql-16-clickhouse_0.11.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.11.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-clickhouse` | `0.11.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 511.9 KiB | [postgresql-16-clickhouse_0.11.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.11.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_15` | `0.11.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 515.1 KiB | [pg_clickhouse_15-0.11.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_15-0.11.1-1PGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_15` | `0.11.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 509.6 KiB | [pg_clickhouse_15-0.11.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_15-0.11.1-1PGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_15` | `0.11.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 521.3 KiB | [pg_clickhouse_15-0.11.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_15-0.11.1-1PGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_15` | `0.11.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 522.1 KiB | [pg_clickhouse_15-0.11.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_15-0.11.1-1PGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_15` | `0.11.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 524.5 KiB | [pg_clickhouse_15-0.11.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_15-0.11.1-1PGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_15` | `0.11.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 525.4 KiB | [pg_clickhouse_15-0.11.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_15-0.11.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-clickhouse` | `0.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 520.8 KiB | [postgresql-15-clickhouse_0.11.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.11.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-clickhouse` | `0.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 512.5 KiB | [postgresql-15-clickhouse_0.11.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.11.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-clickhouse` | `0.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 521.3 KiB | [postgresql-15-clickhouse_0.11.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.11.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-clickhouse` | `0.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 514.5 KiB | [postgresql-15-clickhouse_0.11.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.11.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-clickhouse` | `0.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 608.0 KiB | [postgresql-15-clickhouse_0.11.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.11.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-clickhouse` | `0.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 608.3 KiB | [postgresql-15-clickhouse_0.11.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.11.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-clickhouse` | `0.11.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 521.0 KiB | [postgresql-15-clickhouse_0.11.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.11.1-1PGSTY~noble_amd64.deb) |
| `postgresql-15-clickhouse` | `0.11.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 522.8 KiB | [postgresql-15-clickhouse_0.11.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.11.1-1PGSTY~noble_arm64.deb) |
| `postgresql-15-clickhouse` | `0.11.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 517.8 KiB | [postgresql-15-clickhouse_0.11.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.11.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-clickhouse` | `0.11.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 519.8 KiB | [postgresql-15-clickhouse_0.11.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.11.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_14` | `0.11.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 514.9 KiB | [pg_clickhouse_14-0.11.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_14-0.11.1-1PGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_14` | `0.11.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 509.5 KiB | [pg_clickhouse_14-0.11.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_14-0.11.1-1PGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_14` | `0.11.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 521.3 KiB | [pg_clickhouse_14-0.11.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_14-0.11.1-1PGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_14` | `0.11.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 521.9 KiB | [pg_clickhouse_14-0.11.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_14-0.11.1-1PGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_14` | `0.11.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 524.4 KiB | [pg_clickhouse_14-0.11.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_14-0.11.1-1PGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_14` | `0.11.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 525.3 KiB | [pg_clickhouse_14-0.11.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_14-0.11.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-clickhouse` | `0.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 520.6 KiB | [postgresql-14-clickhouse_0.11.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.11.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-clickhouse` | `0.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 511.8 KiB | [postgresql-14-clickhouse_0.11.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.11.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-clickhouse` | `0.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 520.7 KiB | [postgresql-14-clickhouse_0.11.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.11.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-clickhouse` | `0.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 514.5 KiB | [postgresql-14-clickhouse_0.11.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.11.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-clickhouse` | `0.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 607.7 KiB | [postgresql-14-clickhouse_0.11.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.11.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-clickhouse` | `0.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 607.8 KiB | [postgresql-14-clickhouse_0.11.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.11.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-clickhouse` | `0.11.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 520.7 KiB | [postgresql-14-clickhouse_0.11.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.11.1-1PGSTY~noble_amd64.deb) |
| `postgresql-14-clickhouse` | `0.11.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 522.9 KiB | [postgresql-14-clickhouse_0.11.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.11.1-1PGSTY~noble_arm64.deb) |
| `postgresql-14-clickhouse` | `0.11.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 517.7 KiB | [postgresql-14-clickhouse_0.11.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.11.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-clickhouse` | `0.11.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 519.9 KiB | [postgresql-14-clickhouse_0.11.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.11.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ClickHouse/pg_clickhouse" title="Repository" icon="github" subtitle="github.com/ClickHouse/pg_clickhouse" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_clickhouse-0.11.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_clickhouse;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_clickhouse;		# install via package name, for the active PG version

pig install pg_clickhouse -v 18;   # install for PG 18
pig install pg_clickhouse -v 17;   # install for PG 17
pig install pg_clickhouse -v 16;   # install for PG 16
pig install pg_clickhouse -v 15;   # install for PG 15
pig install pg_clickhouse -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_clickhouse;
```

## Usage

Sources:

- [PGXN README](https://api.pgxn.org/src/pg_clickhouse/pg_clickhouse-0.11.1/README.md)
- [PGXN reference](https://api.pgxn.org/src/pg_clickhouse/pg_clickhouse-0.11.1/doc/pg_clickhouse.md)
- [PGXN tutorial](https://api.pgxn.org/src/pg_clickhouse/pg_clickhouse-0.11.1/doc/tutorial.md)
- [PGXN changelog](https://api.pgxn.org/dist/pg_clickhouse/0.11.1/META.json)
- [PGXN control file](https://api.pgxn.org/src/pg_clickhouse/pg_clickhouse-0.11.1/pg_clickhouse.control)
- [0.10 to 0.11 upgrade SQL](https://api.pgxn.org/src/pg_clickhouse/pg_clickhouse-0.11.1/sql/pg_clickhouse--0.10--0.11.sql)

`pg_clickhouse` 0.11.1 exposes ClickHouse tables to PostgreSQL through the `clickhouse_fdw` foreign data wrapper. Upstream targets PostgreSQL 14 or later and ClickHouse 23.3 or later; package availability is tracked separately. No preload is required for normal use; `session_preload_libraries` and `shared_preload_libraries` are optional connection-startup optimizations.

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

The required `driver` option is `binary` or `http`. Common server options include `host`, `port`, `dbname`, `compression`, `secure`, and `min_tls_version`; user mappings accept `user` and `password`. Version 0.11 removes `fetch_size`; both drivers stream the Native format.

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

`clickhouse_query(server, sql)` returns rows using the caller-provided column definition, while `clickhouse_perform(server, sql)` discards any result. Both can run arbitrary remote SQL, so `EXECUTE` is revoked from `PUBLIC` and should be granted narrowly. `clickhouse_raw_query()` is removed in 0.11; migrate to these interfaces.

### Pushdown and Session Settings

Version 0.10 expands aggregate and function pushdown, improves aggregate execution over mixed local and foreign partitions, and fixes several PostgreSQL NULL-semantics mismatches. Subquery pushdown requires ClickHouse 25.8 or later; older servers evaluate those subqueries locally.

The default `pg_clickhouse.session_settings` preserves PostgreSQL-compatible behavior, including `join_use_nulls = 1`, `group_by_use_nulls = 1`, `final = 1`, and `transform_null_in = 0`. If it is overridden, retain the settings needed by the workload—especially `transform_null_in = 0`, which is required for safe `IN` pushdown.

### Upgrade and Operational Boundaries

```sql
ALTER EXTENSION pg_clickhouse UPDATE TO '0.11';
SELECT pgch_version();
```

The extension SQL version is `0.11`, while `pgch_version()` reports library version `0.11.1`. Run `ALTER EXTENSION` after installing the matching files to remove the obsolete function from the SQL catalog.

If `pg_clickhouse` is placed in `session_preload_libraries`, new sessions load it automatically. If it is placed in `shared_preload_libraries`, changing the library requires a PostgreSQL restart. Neither setting is mandatory, unlike extensions that register postmaster hooks.

Lightweight `UPDATE` and `DELETE` remain outside the documented write surface. Treat direct remote SQL as privileged, test pushdown with production-shaped NULL and type cases, and validate both PostgreSQL and ClickHouse versions before relying on a version-gated optimization.

Version 0.11 rejects text or JSON with invalid database encoding by default. The server option `encoding_check` accepts `fail`, `replace`, `remove` or `truncate`; the latter three alter returned data and require a deliberate policy. Import mappings also improve large integers, intervals and Array/Tuple/Map/Nested types. Existing foreign-table definitions do not automatically follow new inference rules; review their column mappings.

### 0.11.1 Settings and CloudNativePG

Version 0.11.1 documents `settings` options on servers and foreign tables. Table settings override matching server settings; session settings override both, and conflicting table values fall back with a NOTICE. The release also documents CloudNativePG ExtensionConfiguration and Database resources for mounting the matching image.
