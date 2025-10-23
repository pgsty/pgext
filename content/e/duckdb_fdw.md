---
title: "duckdb_fdw"
linkTitle: "duckdb_fdw"
description: "DuckDB Foreign Data Wrapper"
weight: 2450
categories: ["OLAP"]
width: full
---

DuckDB Foreign Data Wrapper


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2450** | {{< badge content="duckdb_fdw" link="https://github.com/alitrack/duckdb_fdw" >}} | {{< ext "duckdb_fdw" >}} | `1.1.2` | {{< category "OLAP" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_analytics" >}} {{< ext "pg_duckdb" >}} {{< ext "pg_mooncake" >}} {{< ext "pg_parquet" >}} {{< ext "wrappers" >}} {{< ext "citus_columnar" >}} {{< ext "columnar" >}} {{< ext "citus" >}} |

> [!Note] conflict on libduckdb with pg_duckdb


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/duckdb_fdw" >}} | `1.1.2` | {{< bg "18" "duckdb_fdw_18*" "red" >}} {{< bg "17" "duckdb_fdw_17*" "green" >}} {{< bg "16" "duckdb_fdw_16*" "green" >}} {{< bg "15" "duckdb_fdw_15*" "green" >}} {{< bg "14" "duckdb_fdw_14*" "green" >}} | `duckdb_fdw_$v*` | `libduckdb` |
| **Debian** | {{< badge content="PIGSTY" link="/e/duckdb_fdw" >}} | `1.1.2` | {{< bg "18" "postgresql-18-duckdb-fdw" "red" >}} {{< bg "17" "postgresql-17-duckdb-fdw" "green" >}} {{< bg "16" "postgresql-16-duckdb-fdw" "green" >}} {{< bg "15" "postgresql-15-duckdb-fdw" "green" >}} {{< bg "14" "postgresql-14-duckdb-fdw" "green" >}} | `postgresql-$v-duckdb-fdw` | `libduckdb` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |  {{< bg "MISS" "duckdb_fdw_18 : HIDE 0" >}}   |  {{< bg "PIGSTY 1.1.2" "duckdb_fdw_17 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "duckdb_fdw_16 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "duckdb_fdw_15 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "duckdb_fdw_14 : HIDE 1" >}}   |
|    `el8.aarch64`    |  {{< bg "MISS" "duckdb_fdw_18 : HIDE 0" >}}   |  {{< bg "PIGSTY 1.1.2" "duckdb_fdw_17 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "duckdb_fdw_16 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "duckdb_fdw_15 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "duckdb_fdw_14 : HIDE 1" >}}   |
|    `el9.x86_64`    |  {{< bg "MISS" "duckdb_fdw_18 : HIDE 0" >}}   |  {{< bg "PIGSTY 1.1.2" "duckdb_fdw_17 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "duckdb_fdw_16 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "duckdb_fdw_15 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "duckdb_fdw_14 : HIDE 1" >}}   |
|    `el9.aarch64`    |  {{< bg "MISS" "duckdb_fdw_18 : HIDE 0" >}}   |  {{< bg "PIGSTY 1.1.2" "duckdb_fdw_17 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "duckdb_fdw_16 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "duckdb_fdw_15 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "duckdb_fdw_14 : HIDE 1" >}}   |
|    `d12.x86_64`    |  {{< bg "MISS" "postgresql-18-duckdb-fdw : HIDE 0" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-17-duckdb-fdw : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-16-duckdb-fdw : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-15-duckdb-fdw : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-14-duckdb-fdw : HIDE 1" >}}   |
|    `d12.aarch64`    |  {{< bg "MISS" "postgresql-18-duckdb-fdw : HIDE 0" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-17-duckdb-fdw : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-16-duckdb-fdw : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-15-duckdb-fdw : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-14-duckdb-fdw : HIDE 1" >}}   |
|    `u22.x86_64`    |  {{< bg "MISS" "postgresql-18-duckdb-fdw : HIDE 0" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-17-duckdb-fdw : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-16-duckdb-fdw : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-15-duckdb-fdw : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-14-duckdb-fdw : HIDE 1" >}}   |
|    `u22.aarch64`    |  {{< bg "MISS" "postgresql-18-duckdb-fdw : HIDE 0" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-17-duckdb-fdw : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-16-duckdb-fdw : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-15-duckdb-fdw : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-14-duckdb-fdw : HIDE 1" >}}   |
|    `u24.x86_64`    |  {{< bg "MISS" "postgresql-18-duckdb-fdw : HIDE 0" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-17-duckdb-fdw : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-16-duckdb-fdw : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-15-duckdb-fdw : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-14-duckdb-fdw : HIDE 1" >}}   |
|    `u24.aarch64`    |  {{< bg "MISS" "postgresql-18-duckdb-fdw : HIDE 0" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-17-duckdb-fdw : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-16-duckdb-fdw : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-15-duckdb-fdw : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-14-duckdb-fdw : HIDE 1" >}}   |


{{< tabs items="PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `duckdb_fdw_17` | 1.1.2 | `el8.x86_64` | pigsty | 83.2 KiB | [duckdb_fdw_17-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/duckdb_fdw_17-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `duckdb_fdw_17` | 1.1.2 | `el8.aarch64` | pigsty | 76.0 KiB | [duckdb_fdw_17-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/duckdb_fdw_17-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `duckdb_fdw_17` | 1.1.2 | `el9.x86_64` | pigsty | 80.6 KiB | [duckdb_fdw_17-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/duckdb_fdw_17-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `duckdb_fdw_17` | 1.1.2 | `el9.aarch64` | pigsty | 78.2 KiB | [duckdb_fdw_17-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/duckdb_fdw_17-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-duckdb-fdw` | 1.1.2 | `d12.x86_64` | pigsty | 256.1 KiB | [postgresql-17-duckdb-fdw_1.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_1.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-duckdb-fdw` | 1.1.2 | `d12.aarch64` | pigsty | 249.8 KiB | [postgresql-17-duckdb-fdw_1.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_1.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-duckdb-fdw` | 1.1.2 | `u22.x86_64` | pigsty | 266.4 KiB | [postgresql-17-duckdb-fdw_1.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_1.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-duckdb-fdw` | 1.1.2 | `u22.aarch64` | pigsty | 262.6 KiB | [postgresql-17-duckdb-fdw_1.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_1.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-duckdb-fdw` | 1.1.2 | `u24.x86_64` | pigsty | 217.8 KiB | [postgresql-17-duckdb-fdw_1.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_1.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-duckdb-fdw` | 1.1.2 | `u24.aarch64` | pigsty | 214.9 KiB | [postgresql-17-duckdb-fdw_1.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_1.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `duckdb_fdw_16` | 1.1.2 | `el8.x86_64` | pigsty | 83.2 KiB | [duckdb_fdw_16-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/duckdb_fdw_16-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `duckdb_fdw_16` | 1.1.2 | `el8.aarch64` | pigsty | 76.0 KiB | [duckdb_fdw_16-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/duckdb_fdw_16-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `duckdb_fdw_16` | 1.1.2 | `el9.x86_64` | pigsty | 80.6 KiB | [duckdb_fdw_16-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/duckdb_fdw_16-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `duckdb_fdw_16` | 1.1.2 | `el9.aarch64` | pigsty | 78.2 KiB | [duckdb_fdw_16-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/duckdb_fdw_16-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-duckdb-fdw` | 1.1.2 | `d12.x86_64` | pigsty | 255.1 KiB | [postgresql-16-duckdb-fdw_1.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_1.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-duckdb-fdw` | 1.1.2 | `d12.aarch64` | pigsty | 248.8 KiB | [postgresql-16-duckdb-fdw_1.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_1.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-duckdb-fdw` | 1.1.2 | `u22.x86_64` | pigsty | 265.6 KiB | [postgresql-16-duckdb-fdw_1.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_1.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-duckdb-fdw` | 1.1.2 | `u22.aarch64` | pigsty | 261.8 KiB | [postgresql-16-duckdb-fdw_1.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_1.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-duckdb-fdw` | 1.1.2 | `u24.x86_64` | pigsty | 217.9 KiB | [postgresql-16-duckdb-fdw_1.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_1.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-duckdb-fdw` | 1.1.2 | `u24.aarch64` | pigsty | 214.9 KiB | [postgresql-16-duckdb-fdw_1.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_1.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `duckdb_fdw_15` | 1.1.2 | `el8.x86_64` | pigsty | 86.0 KiB | [duckdb_fdw_15-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/duckdb_fdw_15-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `duckdb_fdw_15` | 1.1.2 | `el8.aarch64` | pigsty | 78.6 KiB | [duckdb_fdw_15-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/duckdb_fdw_15-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `duckdb_fdw_15` | 1.1.2 | `el9.x86_64` | pigsty | 84.5 KiB | [duckdb_fdw_15-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/duckdb_fdw_15-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `duckdb_fdw_15` | 1.1.2 | `el9.aarch64` | pigsty | 82.4 KiB | [duckdb_fdw_15-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/duckdb_fdw_15-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-duckdb-fdw` | 1.1.2 | `d12.x86_64` | pigsty | 258.9 KiB | [postgresql-15-duckdb-fdw_1.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_1.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-duckdb-fdw` | 1.1.2 | `d12.aarch64` | pigsty | 252.2 KiB | [postgresql-15-duckdb-fdw_1.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_1.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-duckdb-fdw` | 1.1.2 | `u22.x86_64` | pigsty | 277.6 KiB | [postgresql-15-duckdb-fdw_1.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_1.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-duckdb-fdw` | 1.1.2 | `u22.aarch64` | pigsty | 273.5 KiB | [postgresql-15-duckdb-fdw_1.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_1.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-duckdb-fdw` | 1.1.2 | `u24.x86_64` | pigsty | 229.1 KiB | [postgresql-15-duckdb-fdw_1.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_1.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-duckdb-fdw` | 1.1.2 | `u24.aarch64` | pigsty | 227.3 KiB | [postgresql-15-duckdb-fdw_1.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_1.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `duckdb_fdw_14` | 1.1.2 | `el8.x86_64` | pigsty | 86.0 KiB | [duckdb_fdw_14-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/duckdb_fdw_14-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `duckdb_fdw_14` | 1.1.2 | `el8.aarch64` | pigsty | 78.6 KiB | [duckdb_fdw_14-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/duckdb_fdw_14-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `duckdb_fdw_14` | 1.1.2 | `el9.x86_64` | pigsty | 84.5 KiB | [duckdb_fdw_14-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/duckdb_fdw_14-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `duckdb_fdw_14` | 1.1.2 | `el9.aarch64` | pigsty | 82.3 KiB | [duckdb_fdw_14-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/duckdb_fdw_14-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-duckdb-fdw` | 1.1.2 | `d12.x86_64` | pigsty | 258.9 KiB | [postgresql-14-duckdb-fdw_1.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-14-duckdb-fdw_1.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-duckdb-fdw` | 1.1.2 | `d12.aarch64` | pigsty | 252.2 KiB | [postgresql-14-duckdb-fdw_1.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-14-duckdb-fdw_1.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-duckdb-fdw` | 1.1.2 | `u22.x86_64` | pigsty | 277.5 KiB | [postgresql-14-duckdb-fdw_1.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-14-duckdb-fdw_1.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-duckdb-fdw` | 1.1.2 | `u22.aarch64` | pigsty | 273.3 KiB | [postgresql-14-duckdb-fdw_1.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-14-duckdb-fdw_1.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-duckdb-fdw` | 1.1.2 | `u24.x86_64` | pigsty | 229.1 KiB | [postgresql-14-duckdb-fdw_1.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-14-duckdb-fdw_1.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-duckdb-fdw` | 1.1.2 | `u24.aarch64` | pigsty | 227.3 KiB | [postgresql-14-duckdb-fdw_1.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-14-duckdb-fdw_1.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/alitrack/duckdb_fdw" title="Repository" icon="github" subtitle="github.com/alitrack/duckdb_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="duckdb_fdw-1.1.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get duckdb_fdw; # get duckdb_fdw source code
pig build dep duckdb_fdw; # install build dependencies
pig build pkg duckdb_fdw; # build extension rpm or deb
pig build ext duckdb_fdw; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install duckdb_fdw; # install by extension name, for the current active PG version
pig ext install duckdb_fdw; # install via package alias, for the active PG version
pig ext install duckdb_fdw -v 17;   # install for PG 17
pig ext install duckdb_fdw -v 16;   # install for PG 16
pig ext install duckdb_fdw -v 15;   # install for PG 15
pig ext install duckdb_fdw -v 14;   # install for PG 14
pig ext install duckdb_fdw -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION duckdb_fdw;
```



--------

## Usage


### Create Extension

After install the `duckdb_fdw` yum package, you can create the extension inside PostgreSQL database:

```sql
-- create extension
CREATE EXTENSION duckdb_fdw;

-- create duckdb_fdw server
CREATE SERVER duckdb_server FOREIGN DATA WRAPPER duckdb_fdw OPTIONS (database '/tmp/duck.db');

-- create user mapping [OPTIONAL]
-- GRANT USAGE ON FOREIGN SERVER duckdb_server TO PUBLIC;

SELECT duckdb_fdw_version();

-- You can execute duckdb command with `duckdb_execute`, for example, to create a table inside duckdb:
-- create a table in duckdb
SELECT duckdb_execute('duckdb_server', 'CREATE TABLE t1 (a integer,b varchar);');

-- create duckdb foreign table mapping that duckdb table
CREATE FOREIGN TABLE t1 (
    a integer,
    b text
) SERVER duckdb_server OPTIONS (
    table 't1'
);

-- write some data and read it back
INSERT INTO t1 SELECT i AS a,i::text AS b FROM generate_series(1,10) i;
SELECT * FROM t1;
```


You can also import foreign schema from duckdb server, for example, create a table with duckdb cli:

```bash
duckdb /tmp/duck.db

CREATE TABLE t1 (
  a integer,
  b text
);
  
INSERT INTO t1 VALUES (1, 'a'), (2 , 'b'), (3, 'c');
SELECT * FROM t1;
```

Then import the schema into PostgreSQL:

```sql
IMPORT FOREIGN SCHEMA public FROM SERVER duckdb_server INTO public;
```

### Other Resource

- [DuckDB Website](https://duckdb.org/)
- [GitHub: duckdb_fdw](https://github.com/alitrack/duckdb_fdw/)
- [Building libduckdb](https://github.com/digoal/blog/blob/master/202401/20240124_01.md)

