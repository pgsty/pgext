---
title: "duckdb_fdw"
linkTitle: "duckdb_fdw"
description: "DuckDB Foreign Data Wrapper"
weight: 2470
categories: ["OLAP"]
width: full
---

[**duckdb_fdw**](https://github.com/alitrack/duckdb_fdw) : DuckDB Foreign Data Wrapper


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2470** | {{< badge content="duckdb_fdw" link="https://github.com/alitrack/duckdb_fdw" >}} | {{< ext "duckdb_fdw" >}} | `1.4.3` | {{< category "OLAP" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_analytics" >}} {{< ext "pg_duckdb" >}} {{< ext "pg_mooncake" >}} {{< ext "pg_parquet" >}} {{< ext "wrappers" >}} {{< ext "citus_columnar" >}} {{< ext "columnar" >}} {{< ext "citus" >}} |

> [!Note] depend on pg_duckdb's libduckdb, memory mode is break


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `duckdb_fdw` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4.3` | {{< bg "18" "duckdb_fdw_18" "green" >}} {{< bg "17" "duckdb_fdw_17" "green" >}} {{< bg "16" "duckdb_fdw_16" "green" >}} {{< bg "15" "duckdb_fdw_15" "green" >}} {{< bg "14" "duckdb_fdw_14" "green" >}} | `duckdb_fdw_$v` | `pg_duckdb_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4.3` | {{< bg "18" "postgresql-18-duckdb-fdw" "green" >}} {{< bg "17" "postgresql-17-duckdb-fdw" "green" >}} {{< bg "16" "postgresql-16-duckdb-fdw" "green" >}} {{< bg "15" "postgresql-15-duckdb-fdw" "green" >}} {{< bg "14" "postgresql-14-duckdb-fdw" "green" >}} | `postgresql-$v-duckdb-fdw` | `postgresql-$v-pg-duckdb` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-18-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-17-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-16-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-15-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-14-duckdb-fdw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-18-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-17-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-16-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-15-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-14-duckdb-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-18-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-17-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-16-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-15-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-14-duckdb-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-18-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-17-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-16-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-15-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-14-duckdb-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-18-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-17-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-16-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-15-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-14-duckdb-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-18-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-17-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-16-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-15-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-14-duckdb-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-18-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-17-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-16-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-15-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-14-duckdb-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-18-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-17-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-16-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-15-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-14-duckdb-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-duckdb-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-duckdb-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-duckdb-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-duckdb-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-duckdb-fdw : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-duckdb-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-duckdb-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-duckdb-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-duckdb-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-duckdb-fdw : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `duckdb_fdw_18` | `1.4.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 88.1 KiB | [duckdb_fdw_18-1.4.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/duckdb_fdw_18-1.4.3-1PIGSTY.el8.x86_64.rpm) |
| `duckdb_fdw_18` | `1.4.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 82.3 KiB | [duckdb_fdw_18-1.4.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/duckdb_fdw_18-1.4.3-1PIGSTY.el8.aarch64.rpm) |
| `duckdb_fdw_18` | `1.4.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 85.7 KiB | [duckdb_fdw_18-1.4.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/duckdb_fdw_18-1.4.3-1PIGSTY.el9.x86_64.rpm) |
| `duckdb_fdw_18` | `1.4.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 82.6 KiB | [duckdb_fdw_18-1.4.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/duckdb_fdw_18-1.4.3-1PIGSTY.el9.aarch64.rpm) |
| `duckdb_fdw_18` | `1.4.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 80.3 KiB | [duckdb_fdw_18-1.4.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/duckdb_fdw_18-1.4.3-1PIGSTY.el10.x86_64.rpm) |
| `duckdb_fdw_18` | `1.4.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 76.3 KiB | [duckdb_fdw_18-1.4.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/duckdb_fdw_18-1.4.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-duckdb-fdw` | `1.4.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 210.1 KiB | [postgresql-18-duckdb-fdw_1.4.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-18-duckdb-fdw_1.4.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-duckdb-fdw` | `1.4.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 203.9 KiB | [postgresql-18-duckdb-fdw_1.4.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-18-duckdb-fdw_1.4.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-duckdb-fdw` | `1.4.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 211.9 KiB | [postgresql-18-duckdb-fdw_1.4.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/duckdb-fdw/postgresql-18-duckdb-fdw_1.4.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-duckdb-fdw` | `1.4.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 205.2 KiB | [postgresql-18-duckdb-fdw_1.4.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/duckdb-fdw/postgresql-18-duckdb-fdw_1.4.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-duckdb-fdw` | `1.4.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 225.5 KiB | [postgresql-18-duckdb-fdw_1.4.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-18-duckdb-fdw_1.4.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-duckdb-fdw` | `1.4.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 221.6 KiB | [postgresql-18-duckdb-fdw_1.4.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-18-duckdb-fdw_1.4.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-duckdb-fdw` | `1.4.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 214.4 KiB | [postgresql-18-duckdb-fdw_1.4.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-18-duckdb-fdw_1.4.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-duckdb-fdw` | `1.4.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 210.4 KiB | [postgresql-18-duckdb-fdw_1.4.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-18-duckdb-fdw_1.4.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `duckdb_fdw_17` | `1.4.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 88.0 KiB | [duckdb_fdw_17-1.4.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/duckdb_fdw_17-1.4.3-1PIGSTY.el8.x86_64.rpm) |
| `duckdb_fdw_17` | `1.4.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 82.1 KiB | [duckdb_fdw_17-1.4.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/duckdb_fdw_17-1.4.3-1PIGSTY.el8.aarch64.rpm) |
| `duckdb_fdw_17` | `1.4.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 85.6 KiB | [duckdb_fdw_17-1.4.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/duckdb_fdw_17-1.4.3-1PIGSTY.el9.x86_64.rpm) |
| `duckdb_fdw_17` | `1.4.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 82.4 KiB | [duckdb_fdw_17-1.4.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/duckdb_fdw_17-1.4.3-1PIGSTY.el9.aarch64.rpm) |
| `duckdb_fdw_17` | `1.4.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 80.2 KiB | [duckdb_fdw_17-1.4.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/duckdb_fdw_17-1.4.3-1PIGSTY.el10.x86_64.rpm) |
| `duckdb_fdw_17` | `1.4.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 76.1 KiB | [duckdb_fdw_17-1.4.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/duckdb_fdw_17-1.4.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-duckdb-fdw` | `1.4.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 209.7 KiB | [postgresql-17-duckdb-fdw_1.4.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_1.4.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-duckdb-fdw` | `1.4.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 203.5 KiB | [postgresql-17-duckdb-fdw_1.4.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_1.4.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-duckdb-fdw` | `1.4.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 211.4 KiB | [postgresql-17-duckdb-fdw_1.4.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_1.4.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-duckdb-fdw` | `1.4.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 204.8 KiB | [postgresql-17-duckdb-fdw_1.4.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_1.4.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-duckdb-fdw` | `1.4.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 273.1 KiB | [postgresql-17-duckdb-fdw_1.4.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_1.4.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-duckdb-fdw` | `1.4.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 268.8 KiB | [postgresql-17-duckdb-fdw_1.4.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_1.4.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-duckdb-fdw` | `1.4.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 213.5 KiB | [postgresql-17-duckdb-fdw_1.4.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_1.4.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-duckdb-fdw` | `1.4.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 210.1 KiB | [postgresql-17-duckdb-fdw_1.4.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_1.4.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `duckdb_fdw_16` | `1.4.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 89.4 KiB | [duckdb_fdw_16-1.4.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/duckdb_fdw_16-1.4.3-1PIGSTY.el8.x86_64.rpm) |
| `duckdb_fdw_16` | `1.4.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 82.1 KiB | [duckdb_fdw_16-1.4.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/duckdb_fdw_16-1.4.3-1PIGSTY.el8.aarch64.rpm) |
| `duckdb_fdw_16` | `1.4.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 85.0 KiB | [duckdb_fdw_16-1.4.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/duckdb_fdw_16-1.4.3-1PIGSTY.el9.x86_64.rpm) |
| `duckdb_fdw_16` | `1.4.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 82.3 KiB | [duckdb_fdw_16-1.4.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/duckdb_fdw_16-1.4.3-1PIGSTY.el9.aarch64.rpm) |
| `duckdb_fdw_16` | `1.4.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 79.3 KiB | [duckdb_fdw_16-1.4.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/duckdb_fdw_16-1.4.3-1PIGSTY.el10.x86_64.rpm) |
| `duckdb_fdw_16` | `1.4.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 76.0 KiB | [duckdb_fdw_16-1.4.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/duckdb_fdw_16-1.4.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-duckdb-fdw` | `1.4.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 209.9 KiB | [postgresql-16-duckdb-fdw_1.4.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_1.4.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-duckdb-fdw` | `1.4.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 203.4 KiB | [postgresql-16-duckdb-fdw_1.4.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_1.4.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-duckdb-fdw` | `1.4.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 211.5 KiB | [postgresql-16-duckdb-fdw_1.4.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_1.4.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-duckdb-fdw` | `1.4.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 204.8 KiB | [postgresql-16-duckdb-fdw_1.4.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_1.4.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-duckdb-fdw` | `1.4.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 272.2 KiB | [postgresql-16-duckdb-fdw_1.4.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_1.4.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-duckdb-fdw` | `1.4.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 267.9 KiB | [postgresql-16-duckdb-fdw_1.4.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_1.4.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-duckdb-fdw` | `1.4.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 213.5 KiB | [postgresql-16-duckdb-fdw_1.4.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_1.4.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-duckdb-fdw` | `1.4.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 210.2 KiB | [postgresql-16-duckdb-fdw_1.4.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_1.4.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `duckdb_fdw_15` | `1.4.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 93.5 KiB | [duckdb_fdw_15-1.4.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/duckdb_fdw_15-1.4.3-1PIGSTY.el8.x86_64.rpm) |
| `duckdb_fdw_15` | `1.4.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 86.2 KiB | [duckdb_fdw_15-1.4.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/duckdb_fdw_15-1.4.3-1PIGSTY.el8.aarch64.rpm) |
| `duckdb_fdw_15` | `1.4.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 90.0 KiB | [duckdb_fdw_15-1.4.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/duckdb_fdw_15-1.4.3-1PIGSTY.el9.x86_64.rpm) |
| `duckdb_fdw_15` | `1.4.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 87.3 KiB | [duckdb_fdw_15-1.4.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/duckdb_fdw_15-1.4.3-1PIGSTY.el9.aarch64.rpm) |
| `duckdb_fdw_15` | `1.4.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 93.4 KiB | [duckdb_fdw_15-1.4.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/duckdb_fdw_15-1.4.3-1PIGSTY.el10.x86_64.rpm) |
| `duckdb_fdw_15` | `1.4.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 89.8 KiB | [duckdb_fdw_15-1.4.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/duckdb_fdw_15-1.4.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-duckdb-fdw` | `1.4.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 213.6 KiB | [postgresql-15-duckdb-fdw_1.4.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_1.4.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-duckdb-fdw` | `1.4.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 207.2 KiB | [postgresql-15-duckdb-fdw_1.4.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_1.4.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-duckdb-fdw` | `1.4.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 215.5 KiB | [postgresql-15-duckdb-fdw_1.4.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_1.4.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-duckdb-fdw` | `1.4.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 208.5 KiB | [postgresql-15-duckdb-fdw_1.4.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_1.4.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-duckdb-fdw` | `1.4.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 283.8 KiB | [postgresql-15-duckdb-fdw_1.4.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_1.4.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-duckdb-fdw` | `1.4.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 279.6 KiB | [postgresql-15-duckdb-fdw_1.4.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_1.4.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-duckdb-fdw` | `1.4.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 225.5 KiB | [postgresql-15-duckdb-fdw_1.4.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_1.4.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-duckdb-fdw` | `1.4.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 222.3 KiB | [postgresql-15-duckdb-fdw_1.4.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_1.4.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `duckdb_fdw_14` | `1.4.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 93.5 KiB | [duckdb_fdw_14-1.4.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/duckdb_fdw_14-1.4.3-1PIGSTY.el8.x86_64.rpm) |
| `duckdb_fdw_14` | `1.4.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 86.2 KiB | [duckdb_fdw_14-1.4.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/duckdb_fdw_14-1.4.3-1PIGSTY.el8.aarch64.rpm) |
| `duckdb_fdw_14` | `1.4.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 90.0 KiB | [duckdb_fdw_14-1.4.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/duckdb_fdw_14-1.4.3-1PIGSTY.el9.x86_64.rpm) |
| `duckdb_fdw_14` | `1.4.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 87.3 KiB | [duckdb_fdw_14-1.4.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/duckdb_fdw_14-1.4.3-1PIGSTY.el9.aarch64.rpm) |
| `duckdb_fdw_14` | `1.4.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 93.4 KiB | [duckdb_fdw_14-1.4.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/duckdb_fdw_14-1.4.3-1PIGSTY.el10.x86_64.rpm) |
| `duckdb_fdw_14` | `1.4.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 89.6 KiB | [duckdb_fdw_14-1.4.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/duckdb_fdw_14-1.4.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-duckdb-fdw` | `1.4.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 213.4 KiB | [postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-duckdb-fdw` | `1.4.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 207.0 KiB | [postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-duckdb-fdw` | `1.4.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 215.3 KiB | [postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/duckdb-fdw/postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-duckdb-fdw` | `1.4.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 208.2 KiB | [postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/duckdb-fdw/postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-duckdb-fdw` | `1.4.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 283.7 KiB | [postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-duckdb-fdw` | `1.4.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 279.4 KiB | [postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-duckdb-fdw` | `1.4.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 225.4 KiB | [postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-duckdb-fdw` | `1.4.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 222.4 KiB | [postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/alitrack/duckdb_fdw" title="Repository" icon="github" subtitle="github.com/alitrack/duckdb_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="duckdb_fdw-1.4.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg duckdb_fdw;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install duckdb_fdw;		# install via package name, for the active PG version

pig install duckdb_fdw -v 18;   # install for PG 18
pig install duckdb_fdw -v 17;   # install for PG 17
pig install duckdb_fdw -v 16;   # install for PG 16
pig install duckdb_fdw -v 15;   # install for PG 15
pig install duckdb_fdw -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION duckdb_fdw;
```


> [!WARNING] This extension is currently broken and conflict with pg_duckdb & pg_mooncake

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

