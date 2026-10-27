---
title: "duckdb_fdw"
linkTitle: "duckdb_fdw"
description: "DuckDB Foreign Data Wrapper"
weight: 2470
categories: ["OLAP"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**duckdb_fdw**](https://github.com/alitrack/duckdb_fdw) : DuckDB Foreign Data Wrapper


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2470** | {{< badge content="duckdb_fdw" link="https://github.com/alitrack/duckdb_fdw" >}} | {{< ext "duckdb_fdw" >}} | `2.0.2` | {{< category "OLAP" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_duckdb" >}} {{< ext "pg_ducklake" >}} {{< ext "pg_mooncake" >}} {{< ext "pg_clickhouse" >}} {{< ext "pg_parquet" >}} {{< ext "pg_lake" >}} {{< ext "aws_s3" >}} {{< ext "file_fdw" >}} {{< ext "pg_bulkload" >}} {{< ext "pg_stat_ch" >}} |

> [!Note] PGSTY 2.0.2 embeds DuckDB 2.0.0 statically, has no pg_duckdb runtime dependency, and conflicts with pg_duckdb for the same PG major. Switching from pg_duckdb requires removing its preload setting and restarting. Current builds support PG15-18; PG14 fails upstream boolVal/makeBoolean API use despite the README claiming PG13-18. Historical PG14 1.4.3 packages remain available.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `duckdb_fdw` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.2` | {{< bg "18" "duckdb_fdw_18" "green" >}} {{< bg "17" "duckdb_fdw_17" "green" >}} {{< bg "16" "duckdb_fdw_16" "green" >}} {{< bg "15" "duckdb_fdw_15" "green" >}} {{< bg "14" "duckdb_fdw_14" "green" >}} | `duckdb_fdw_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.2` | {{< bg "18" "postgresql-18-duckdb-fdw" "green" >}} {{< bg "17" "postgresql-17-duckdb-fdw" "green" >}} {{< bg "16" "postgresql-16-duckdb-fdw" "green" >}} {{< bg "15" "postgresql-15-duckdb-fdw" "green" >}} {{< bg "14" "postgresql-14-duckdb-fdw" "green" >}} | `postgresql-$v-duckdb-fdw` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "duckdb_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "duckdb_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-14-duckdb-fdw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-14-duckdb-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-14-duckdb-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-14-duckdb-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-14-duckdb-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-14-duckdb-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-14-duckdb-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-14-duckdb-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-14-duckdb-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-duckdb-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.3" "postgresql-14-duckdb-fdw : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `duckdb_fdw_18` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.5 MiB | [duckdb_fdw_18-2.0.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/duckdb_fdw_18-2.0.2-1PGSTY.el8.x86_64.rpm) |
| `duckdb_fdw_18` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.4 MiB | [duckdb_fdw_18-2.0.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/duckdb_fdw_18-2.0.2-1PGSTY.el8.aarch64.rpm) |
| `duckdb_fdw_18` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.5 MiB | [duckdb_fdw_18-2.0.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/duckdb_fdw_18-2.0.2-1PGSTY.el9.x86_64.rpm) |
| `duckdb_fdw_18` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.2 MiB | [duckdb_fdw_18-2.0.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/duckdb_fdw_18-2.0.2-1PGSTY.el9.aarch64.rpm) |
| `duckdb_fdw_18` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.3 MiB | [duckdb_fdw_18-2.0.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/duckdb_fdw_18-2.0.2-1PGSTY.el10.x86_64.rpm) |
| `duckdb_fdw_18` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.6 MiB | [duckdb_fdw_18-2.0.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/duckdb_fdw_18-2.0.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-duckdb-fdw` | `2.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.9 MiB | [postgresql-18-duckdb-fdw_2.0.2-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-18-duckdb-fdw_2.0.2-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-duckdb-fdw` | `2.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.9 MiB | [postgresql-18-duckdb-fdw_2.0.2-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-18-duckdb-fdw_2.0.2-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-duckdb-fdw` | `2.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.5 MiB | [postgresql-18-duckdb-fdw_2.0.2-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/duckdb-fdw/postgresql-18-duckdb-fdw_2.0.2-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-duckdb-fdw` | `2.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.3 MiB | [postgresql-18-duckdb-fdw_2.0.2-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/duckdb-fdw/postgresql-18-duckdb-fdw_2.0.2-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-duckdb-fdw` | `2.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.3 MiB | [postgresql-18-duckdb-fdw_2.0.2-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-18-duckdb-fdw_2.0.2-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-duckdb-fdw` | `2.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.3 MiB | [postgresql-18-duckdb-fdw_2.0.2-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-18-duckdb-fdw_2.0.2-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-duckdb-fdw` | `2.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.7 MiB | [postgresql-18-duckdb-fdw_2.0.2-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-18-duckdb-fdw_2.0.2-1PGSTY~noble_amd64.deb) |
| `postgresql-18-duckdb-fdw` | `2.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.8 MiB | [postgresql-18-duckdb-fdw_2.0.2-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-18-duckdb-fdw_2.0.2-1PGSTY~noble_arm64.deb) |
| `postgresql-18-duckdb-fdw` | `2.0.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.3 MiB | [postgresql-18-duckdb-fdw_2.0.2-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/duckdb-fdw/postgresql-18-duckdb-fdw_2.0.2-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-duckdb-fdw` | `2.0.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 19.0 MiB | [postgresql-18-duckdb-fdw_2.0.2-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/duckdb-fdw/postgresql-18-duckdb-fdw_2.0.2-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `duckdb_fdw_17` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.5 MiB | [duckdb_fdw_17-2.0.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/duckdb_fdw_17-2.0.2-1PGSTY.el8.x86_64.rpm) |
| `duckdb_fdw_17` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.4 MiB | [duckdb_fdw_17-2.0.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/duckdb_fdw_17-2.0.2-1PGSTY.el8.aarch64.rpm) |
| `duckdb_fdw_17` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.5 MiB | [duckdb_fdw_17-2.0.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/duckdb_fdw_17-2.0.2-1PGSTY.el9.x86_64.rpm) |
| `duckdb_fdw_17` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.2 MiB | [duckdb_fdw_17-2.0.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/duckdb_fdw_17-2.0.2-1PGSTY.el9.aarch64.rpm) |
| `duckdb_fdw_17` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.3 MiB | [duckdb_fdw_17-2.0.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/duckdb_fdw_17-2.0.2-1PGSTY.el10.x86_64.rpm) |
| `duckdb_fdw_17` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.6 MiB | [duckdb_fdw_17-2.0.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/duckdb_fdw_17-2.0.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-duckdb-fdw` | `2.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.9 MiB | [postgresql-17-duckdb-fdw_2.0.2-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_2.0.2-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-duckdb-fdw` | `2.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.9 MiB | [postgresql-17-duckdb-fdw_2.0.2-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_2.0.2-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-duckdb-fdw` | `2.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.5 MiB | [postgresql-17-duckdb-fdw_2.0.2-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_2.0.2-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-duckdb-fdw` | `2.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.3 MiB | [postgresql-17-duckdb-fdw_2.0.2-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_2.0.2-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-duckdb-fdw` | `2.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.4 MiB | [postgresql-17-duckdb-fdw_2.0.2-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_2.0.2-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-duckdb-fdw` | `2.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.3 MiB | [postgresql-17-duckdb-fdw_2.0.2-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_2.0.2-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-duckdb-fdw` | `2.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.7 MiB | [postgresql-17-duckdb-fdw_2.0.2-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_2.0.2-1PGSTY~noble_amd64.deb) |
| `postgresql-17-duckdb-fdw` | `2.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.8 MiB | [postgresql-17-duckdb-fdw_2.0.2-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_2.0.2-1PGSTY~noble_arm64.deb) |
| `postgresql-17-duckdb-fdw` | `2.0.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.3 MiB | [postgresql-17-duckdb-fdw_2.0.2-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_2.0.2-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-duckdb-fdw` | `2.0.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 19.1 MiB | [postgresql-17-duckdb-fdw_2.0.2-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/duckdb-fdw/postgresql-17-duckdb-fdw_2.0.2-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `duckdb_fdw_16` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.5 MiB | [duckdb_fdw_16-2.0.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/duckdb_fdw_16-2.0.2-1PGSTY.el8.x86_64.rpm) |
| `duckdb_fdw_16` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.4 MiB | [duckdb_fdw_16-2.0.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/duckdb_fdw_16-2.0.2-1PGSTY.el8.aarch64.rpm) |
| `duckdb_fdw_16` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.5 MiB | [duckdb_fdw_16-2.0.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/duckdb_fdw_16-2.0.2-1PGSTY.el9.x86_64.rpm) |
| `duckdb_fdw_16` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.2 MiB | [duckdb_fdw_16-2.0.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/duckdb_fdw_16-2.0.2-1PGSTY.el9.aarch64.rpm) |
| `duckdb_fdw_16` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.3 MiB | [duckdb_fdw_16-2.0.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/duckdb_fdw_16-2.0.2-1PGSTY.el10.x86_64.rpm) |
| `duckdb_fdw_16` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.6 MiB | [duckdb_fdw_16-2.0.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/duckdb_fdw_16-2.0.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-duckdb-fdw` | `2.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.9 MiB | [postgresql-16-duckdb-fdw_2.0.2-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_2.0.2-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-duckdb-fdw` | `2.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.9 MiB | [postgresql-16-duckdb-fdw_2.0.2-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_2.0.2-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-duckdb-fdw` | `2.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.5 MiB | [postgresql-16-duckdb-fdw_2.0.2-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_2.0.2-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-duckdb-fdw` | `2.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.3 MiB | [postgresql-16-duckdb-fdw_2.0.2-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_2.0.2-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-duckdb-fdw` | `2.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.4 MiB | [postgresql-16-duckdb-fdw_2.0.2-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_2.0.2-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-duckdb-fdw` | `2.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.3 MiB | [postgresql-16-duckdb-fdw_2.0.2-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_2.0.2-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-duckdb-fdw` | `2.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.7 MiB | [postgresql-16-duckdb-fdw_2.0.2-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_2.0.2-1PGSTY~noble_amd64.deb) |
| `postgresql-16-duckdb-fdw` | `2.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.8 MiB | [postgresql-16-duckdb-fdw_2.0.2-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_2.0.2-1PGSTY~noble_arm64.deb) |
| `postgresql-16-duckdb-fdw` | `2.0.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.3 MiB | [postgresql-16-duckdb-fdw_2.0.2-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_2.0.2-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-duckdb-fdw` | `2.0.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 19.1 MiB | [postgresql-16-duckdb-fdw_2.0.2-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/duckdb-fdw/postgresql-16-duckdb-fdw_2.0.2-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `duckdb_fdw_15` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.5 MiB | [duckdb_fdw_15-2.0.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/duckdb_fdw_15-2.0.2-1PGSTY.el8.x86_64.rpm) |
| `duckdb_fdw_15` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.4 MiB | [duckdb_fdw_15-2.0.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/duckdb_fdw_15-2.0.2-1PGSTY.el8.aarch64.rpm) |
| `duckdb_fdw_15` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.5 MiB | [duckdb_fdw_15-2.0.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/duckdb_fdw_15-2.0.2-1PGSTY.el9.x86_64.rpm) |
| `duckdb_fdw_15` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.2 MiB | [duckdb_fdw_15-2.0.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/duckdb_fdw_15-2.0.2-1PGSTY.el9.aarch64.rpm) |
| `duckdb_fdw_15` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.3 MiB | [duckdb_fdw_15-2.0.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/duckdb_fdw_15-2.0.2-1PGSTY.el10.x86_64.rpm) |
| `duckdb_fdw_15` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.6 MiB | [duckdb_fdw_15-2.0.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/duckdb_fdw_15-2.0.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-duckdb-fdw` | `2.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.9 MiB | [postgresql-15-duckdb-fdw_2.0.2-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_2.0.2-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-duckdb-fdw` | `2.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.9 MiB | [postgresql-15-duckdb-fdw_2.0.2-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_2.0.2-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-duckdb-fdw` | `2.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.5 MiB | [postgresql-15-duckdb-fdw_2.0.2-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_2.0.2-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-duckdb-fdw` | `2.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.3 MiB | [postgresql-15-duckdb-fdw_2.0.2-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_2.0.2-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-duckdb-fdw` | `2.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.4 MiB | [postgresql-15-duckdb-fdw_2.0.2-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_2.0.2-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-duckdb-fdw` | `2.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.3 MiB | [postgresql-15-duckdb-fdw_2.0.2-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_2.0.2-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-duckdb-fdw` | `2.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.7 MiB | [postgresql-15-duckdb-fdw_2.0.2-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_2.0.2-1PGSTY~noble_amd64.deb) |
| `postgresql-15-duckdb-fdw` | `2.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.8 MiB | [postgresql-15-duckdb-fdw_2.0.2-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_2.0.2-1PGSTY~noble_arm64.deb) |
| `postgresql-15-duckdb-fdw` | `2.0.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.3 MiB | [postgresql-15-duckdb-fdw_2.0.2-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_2.0.2-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-duckdb-fdw` | `2.0.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 19.0 MiB | [postgresql-15-duckdb-fdw_2.0.2-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/duckdb-fdw/postgresql-15-duckdb-fdw_2.0.2-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

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
| `postgresql-14-duckdb-fdw` | `1.4.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 224.2 KiB | [postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/duckdb-fdw/postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-duckdb-fdw` | `1.4.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 220.6 KiB | [postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/duckdb-fdw/postgresql-14-duckdb-fdw_1.4.3-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/alitrack/duckdb_fdw" title="Repository" icon="github" subtitle="github.com/alitrack/duckdb_fdw" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="duckdb_fdw-2.0.2.tar.gz duckdb-2.0.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg duckdb_fdw;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install duckdb_fdw;		# install via package name, for the active PG version

pig install duckdb_fdw -v 18;   # install for PG 18
pig install duckdb_fdw -v 17;   # install for PG 17
pig install duckdb_fdw -v 16;   # install for PG 16
pig install duckdb_fdw -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION duckdb_fdw;
```

## Usage

Sources:

- [2.0.2 README](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/README.md)
- [2.0.2 control file](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/duckdb_fdw.control)
- [2.0.2 SQL functions and privileges](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/duckdb_fdw--2.0.2.sql)
- [Server, mapping and table options](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/option.c)
- [Scan, write and diagnostic functions](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/duckdb_fdw.c)
- [Connection and transaction handling](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/connection.c)
- [Schema import implementation](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/import.c)
- [2.0.1 to 2.0.2 upgrade](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/duckdb_fdw--2.0.1--2.0.2.sql)

`duckdb_fdw` 2.0.2 exposes DuckDB tables and data sources as PostgreSQL foreign tables through DuckDB's native C API. Use it to query a DuckDB database, read Parquet, or access optional DuckDB catalogs. Foreign tables support reads and inserts; PostgreSQL UPDATE and DELETE through the FDW are not supported.

### Enable and Map a DuckDB Table

Install the extension as a superuser. A new installation does not require preloading or a PostgreSQL restart. Use a persistent DuckDB file in a directory accessible to the PostgreSQL service account; the example assumes that directory already exists.

```sql
CREATE EXTENSION duckdb_fdw;
CREATE SERVER duckdb_server
  FOREIGN DATA WRAPPER duckdb_fdw
  OPTIONS (database '/var/lib/postgresql/duckdb/analytics.duckdb');

SELECT duckdb_execute(
  'duckdb_server',
  'CREATE TABLE measurements (sensor_id INTEGER, reading DOUBLE)'
);

CREATE FOREIGN TABLE measurements (
  sensor_id integer,
  reading double precision
) SERVER duckdb_server OPTIONS (table 'measurements');

INSERT INTO measurements VALUES (1, 21.5), (2, 19.8);
SELECT * FROM measurements WHERE reading > 20;

SELECT extversion FROM pg_extension WHERE extname = 'duckdb_fdw';
SELECT duckdb_fdw_version();
```

The catalog query reports the extension version; despite its name, `duckdb_fdw_version()` reports the linked DuckDB engine version. A user mapping is optional for a local file without credentials. Database-file permissions are those of the PostgreSQL operating-system account.

### Import Tables and Read Parquet

```sql
CREATE SCHEMA duckdb_import;
IMPORT FOREIGN SCHEMA main
  FROM SERVER duckdb_server INTO duckdb_import;
```

Use DuckDB's schema name, normally `main`, rather than PostgreSQL's `public`. To read an existing Parquet file, declare matching PostgreSQL column types:

```sql
CREATE FOREIGN TABLE parquet_measurements (
  sensor_id integer,
  reading double precision
) SERVER duckdb_server
  OPTIONS (table '/var/lib/postgresql/duckdb/measurements.parquet');

SELECT * FROM parquet_measurements WHERE sensor_id = 1;
```

### Options and Privileges

| Object | Purpose |
| --- | --- |
| Server `database` | DuckDB file path, or a connection-local in-memory database |
| Server `force_readonly` | Open file databases read-only and reject FDW writes and non-read-only helper statements |
| Server `extensions`, `attach_catalogs` | Load optional DuckDB extensions and attach catalogs |
| Table `table` | Map a DuckDB table or a Parquet path |
| Column `column_name` | Map a PostgreSQL column to a differently named DuckDB column |
| `duckdb_execute(server, statement)` | Execute DuckDB SQL; returns void, not a result table |
| `duckdb_create_s3_secret(...)` | Create an S3 secret in the DuckDB connection |
| `duckdb_fdw_preflight()` | Inspect runtime compatibility before using DuckDB |

Grant application roles only the foreign-server, schema and foreign-table privileges they need. EXECUTE on the arbitrary-SQL and secret helpers is revoked from PUBLIC on installation; grant it only to trusted roles.

For S3, MotherDuck or Quack credentials, prefer user mappings. Server options are visible in PostgreSQL catalogs and are unsuitable for secrets. Remote examples additionally depend on the relevant DuckDB extensions, network access and credentials; their availability is not implied by installing the FDW.

### Connection and Transaction Boundaries

- Connections close at transaction end by default. With `database ':memory:'`, keep setup and queries in the same explicit transaction or the temporary database disappears.
- `duckdb_fdw.keep_connection` defaults to off. Enabling it retains connections after commit, but a writable file connection holds the DuckDB file lock across transactions and can block other PostgreSQL sessions from opening that file.
- Foreign-table inserts follow the top-level PostgreSQL commit or rollback. Savepoint rollback does not undo individual DuckDB writes: if the outer transaction commits, writes from an aborted subtransaction can remain.
- DuckDB does not participate in PostgreSQL two-phase commit. Do not use prepared transactions or assume crash-atomic commits across PostgreSQL and DuckDB. Arbitrary SQL sent through the helper also does not acquire PostgreSQL transaction semantics automatically.

### Compatibility and Upgrade

The Pigsty 2.0.2 packages target PostgreSQL 15-18. Upstream's README advertises PostgreSQL 13-18, which is broader than this package range. Version 2.0.2 accepts DuckDB 1.5.x and 2.0.x runtimes; it does not change the SQL function set from 2.0.1.

Runtime execution is guarded against loading `pg_duckdb` in the same backend. Keep the two extensions in separate PostgreSQL instances. When switching away from a preloaded `pg_duckdb`, remove its preload entry and restart PostgreSQL; replacing files does not unload libraries from existing processes.

```sql
SELECT duckdb_fdw_runtime_compatibility_status();
SELECT duckdb_fdw_runtime_fingerprint();
SELECT duckdb_fdw_preflight();
```

After installing matching files, use fresh backend sessions and update each database:

```sql
ALTER EXTENSION duckdb_fdw UPDATE TO '2.0.2';
```
