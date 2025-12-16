---
title: "pg_parquet"
linkTitle: "pg_parquet"
description: "copy data between Postgres and Parquet"
weight: 2480
categories: ["OLAP"]
width: full
---

[**pg_parquet**](https://github.com/CrunchyData/pg_parquet/) : copy data between Postgres and Parquet


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2480** | {{< badge content="pg_parquet" link="https://github.com/CrunchyData/pg_parquet/" >}} | {{< ext "pg_parquet" >}} | `0.5.1` | {{< category "OLAP" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLdt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_analytics" >}} {{< ext "pg_duckdb" >}} {{< ext "duckdb_fdw" >}} {{< ext "citus_columnar" >}} {{< ext "columnar" >}} {{< ext "pg_mooncake" >}} {{< ext "aws_s3" >}} {{< ext "citus" >}} |

> [!Note] manual update from 0.16.0


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `pg_parquet` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.1` | {{< bg "18" "pg_parquet_18" "green" >}} {{< bg "17" "pg_parquet_17" "green" >}} {{< bg "16" "pg_parquet_16" "green" >}} {{< bg "15" "pg_parquet_15" "green" >}} {{< bg "14" "pg_parquet_14" "green" >}} {{< bg "13" "pg_parquet_13" "red" >}} | `pg_parquet_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.1` | {{< bg "18" "postgresql-18-pg-parquet" "green" >}} {{< bg "17" "postgresql-17-pg-parquet" "green" >}} {{< bg "16" "postgresql-16-pg-parquet" "green" >}} {{< bg "15" "postgresql-15-pg-parquet" "green" >}} {{< bg "14" "postgresql-14-pg-parquet" "green" >}} {{< bg "13" "postgresql-13-pg-parquet" "red" >}} | `postgresql-$v-pg-parquet` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_parquet_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_parquet_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_parquet_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_parquet_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_parquet_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_parquet_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-parquet : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-parquet : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-parquet : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-parquet : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-parquet : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-parquet : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-parquet : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-parquet : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-parquet : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-parquet : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-parquet : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-parquet : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-parquet : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-parquet : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-parquet : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-parquet : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_parquet_18` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.4 MiB | [pg_parquet_18-0.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_parquet_18-0.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.5 MiB | [pg_parquet_18-0.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_parquet_18-0.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.2 MiB | [pg_parquet_18-0.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_parquet_18-0.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.8 MiB | [pg_parquet_18-0.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_parquet_18-0.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.1 MiB | [pg_parquet_18-0.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_parquet_18-0.5.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.7 MiB | [pg_parquet_18-0.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_parquet_18-0.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-parquet` | `0.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.0 MiB | [postgresql-18-pg-parquet_0.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-parquet/postgresql-18-pg-parquet_0.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-parquet` | `0.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.1 MiB | [postgresql-18-pg-parquet_0.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-parquet/postgresql-18-pg-parquet_0.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-parquet` | `0.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.0 MiB | [postgresql-18-pg-parquet_0.5.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-parquet/postgresql-18-pg-parquet_0.5.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-parquet` | `0.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.1 MiB | [postgresql-18-pg-parquet_0.5.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-parquet/postgresql-18-pg-parquet_0.5.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-parquet` | `0.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-18-pg-parquet_0.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-parquet/postgresql-18-pg-parquet_0.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-parquet` | `0.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.1 MiB | [postgresql-18-pg-parquet_0.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-parquet/postgresql-18-pg-parquet_0.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-parquet` | `0.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.6 MiB | [postgresql-18-pg-parquet_0.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-parquet/postgresql-18-pg-parquet_0.5.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-parquet` | `0.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.0 MiB | [postgresql-18-pg-parquet_0.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-parquet/postgresql-18-pg-parquet_0.5.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_parquet_17` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.4 MiB | [pg_parquet_17-0.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_parquet_17-0.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.5 MiB | [pg_parquet_17-0.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_parquet_17-0.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.2 MiB | [pg_parquet_17-0.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_parquet_17-0.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.8 MiB | [pg_parquet_17-0.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_parquet_17-0.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.1 MiB | [pg_parquet_17-0.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_parquet_17-0.5.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.7 MiB | [pg_parquet_17-0.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_parquet_17-0.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-parquet` | `0.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.0 MiB | [postgresql-17-pg-parquet_0.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-parquet/postgresql-17-pg-parquet_0.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-parquet` | `0.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.1 MiB | [postgresql-17-pg-parquet_0.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-parquet/postgresql-17-pg-parquet_0.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-parquet` | `0.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.0 MiB | [postgresql-17-pg-parquet_0.5.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-parquet/postgresql-17-pg-parquet_0.5.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-parquet` | `0.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.1 MiB | [postgresql-17-pg-parquet_0.5.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-parquet/postgresql-17-pg-parquet_0.5.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-parquet` | `0.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-17-pg-parquet_0.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-parquet/postgresql-17-pg-parquet_0.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-parquet` | `0.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.1 MiB | [postgresql-17-pg-parquet_0.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-parquet/postgresql-17-pg-parquet_0.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-parquet` | `0.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.6 MiB | [postgresql-17-pg-parquet_0.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-parquet/postgresql-17-pg-parquet_0.5.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-parquet` | `0.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.0 MiB | [postgresql-17-pg-parquet_0.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-parquet/postgresql-17-pg-parquet_0.5.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_parquet_16` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.4 MiB | [pg_parquet_16-0.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_parquet_16-0.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.5 MiB | [pg_parquet_16-0.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_parquet_16-0.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.2 MiB | [pg_parquet_16-0.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_parquet_16-0.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.8 MiB | [pg_parquet_16-0.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_parquet_16-0.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.1 MiB | [pg_parquet_16-0.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_parquet_16-0.5.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.7 MiB | [pg_parquet_16-0.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_parquet_16-0.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-parquet` | `0.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.0 MiB | [postgresql-16-pg-parquet_0.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-parquet/postgresql-16-pg-parquet_0.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-parquet` | `0.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.1 MiB | [postgresql-16-pg-parquet_0.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-parquet/postgresql-16-pg-parquet_0.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-parquet` | `0.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.0 MiB | [postgresql-16-pg-parquet_0.5.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-parquet/postgresql-16-pg-parquet_0.5.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-parquet` | `0.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.1 MiB | [postgresql-16-pg-parquet_0.5.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-parquet/postgresql-16-pg-parquet_0.5.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-parquet` | `0.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-16-pg-parquet_0.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-parquet/postgresql-16-pg-parquet_0.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-parquet` | `0.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.1 MiB | [postgresql-16-pg-parquet_0.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-parquet/postgresql-16-pg-parquet_0.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-parquet` | `0.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.6 MiB | [postgresql-16-pg-parquet_0.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-parquet/postgresql-16-pg-parquet_0.5.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-parquet` | `0.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.0 MiB | [postgresql-16-pg-parquet_0.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-parquet/postgresql-16-pg-parquet_0.5.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_parquet_15` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.4 MiB | [pg_parquet_15-0.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_parquet_15-0.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.5 MiB | [pg_parquet_15-0.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_parquet_15-0.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.2 MiB | [pg_parquet_15-0.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_parquet_15-0.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.8 MiB | [pg_parquet_15-0.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_parquet_15-0.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.1 MiB | [pg_parquet_15-0.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_parquet_15-0.5.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.7 MiB | [pg_parquet_15-0.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_parquet_15-0.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-parquet` | `0.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.0 MiB | [postgresql-15-pg-parquet_0.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-parquet/postgresql-15-pg-parquet_0.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-parquet` | `0.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.1 MiB | [postgresql-15-pg-parquet_0.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-parquet/postgresql-15-pg-parquet_0.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-parquet` | `0.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.0 MiB | [postgresql-15-pg-parquet_0.5.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-parquet/postgresql-15-pg-parquet_0.5.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-parquet` | `0.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.1 MiB | [postgresql-15-pg-parquet_0.5.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-parquet/postgresql-15-pg-parquet_0.5.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-parquet` | `0.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-15-pg-parquet_0.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-parquet/postgresql-15-pg-parquet_0.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-parquet` | `0.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.1 MiB | [postgresql-15-pg-parquet_0.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-parquet/postgresql-15-pg-parquet_0.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-parquet` | `0.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.6 MiB | [postgresql-15-pg-parquet_0.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-parquet/postgresql-15-pg-parquet_0.5.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-parquet` | `0.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.0 MiB | [postgresql-15-pg-parquet_0.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-parquet/postgresql-15-pg-parquet_0.5.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_parquet_14` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.4 MiB | [pg_parquet_14-0.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_parquet_14-0.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_parquet_14` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.5 MiB | [pg_parquet_14-0.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_parquet_14-0.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_parquet_14` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.2 MiB | [pg_parquet_14-0.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_parquet_14-0.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_parquet_14` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.8 MiB | [pg_parquet_14-0.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_parquet_14-0.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_parquet_14` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.1 MiB | [pg_parquet_14-0.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_parquet_14-0.5.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_parquet_14` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.7 MiB | [pg_parquet_14-0.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_parquet_14-0.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-parquet` | `0.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.0 MiB | [postgresql-14-pg-parquet_0.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-parquet/postgresql-14-pg-parquet_0.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-parquet` | `0.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.1 MiB | [postgresql-14-pg-parquet_0.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-parquet/postgresql-14-pg-parquet_0.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-parquet` | `0.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.0 MiB | [postgresql-14-pg-parquet_0.5.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-parquet/postgresql-14-pg-parquet_0.5.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-parquet` | `0.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.1 MiB | [postgresql-14-pg-parquet_0.5.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-parquet/postgresql-14-pg-parquet_0.5.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-parquet` | `0.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-14-pg-parquet_0.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-parquet/postgresql-14-pg-parquet_0.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-parquet` | `0.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.1 MiB | [postgresql-14-pg-parquet_0.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-parquet/postgresql-14-pg-parquet_0.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-parquet` | `0.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.6 MiB | [postgresql-14-pg-parquet_0.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-parquet/postgresql-14-pg-parquet_0.5.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-parquet` | `0.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.0 MiB | [postgresql-14-pg-parquet_0.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-parquet/postgresql-14-pg-parquet_0.5.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/CrunchyData/pg_parquet/" title="Repository" icon="github" subtitle="github.com/CrunchyData/pg_parquet/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_parquet-0.5.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_parquet;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_parquet;		# install via package name, for the active PG version

pig install pg_parquet -v 18;   # install for PG 18
pig install pg_parquet -v 17;   # install for PG 17
pig install pg_parquet -v 16;   # install for PG 16
pig install pg_parquet -v 15;   # install for PG 15
pig install pg_parquet -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pg_parquet';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_parquet;
```


## Usage

There are mainly 3 things that you can do with `pg_parquet`:

1. You can export Postgres tables/queries to Parquet files,
2. You can ingest data from Parquet files to Postgres tables,
3. You can inspect the schema and metadata of Parquet files.

### COPY to/from Parquet files from/to Postgres tables



You can use PostgreSQL's `COPY` command to read and write from/to Parquet files. Below is an example of how to write a PostgreSQL table, with complex types, into a Parquet file and then to read the Parquet file content back into the same table.

```sql
-- create composite types
CREATE TYPE product_item AS (id INT, name TEXT, price float4);
CREATE TYPE product AS (id INT, name TEXT, items product_item[]);

-- create a table with complex types
CREATE TABLE product_example (
    id int,
    product product,
    products product[],
    created_at TIMESTAMP,
    updated_at TIMESTAMPTZ
);

-- insert some rows into the table
insert into product_example values (
    1,
    ROW(1, 'product 1', ARRAY[ROW(1, 'item 1', 1.0), ROW(2, 'item 2', 2.0), NULL]::product_item[])::product,
    ARRAY[ROW(1, NULL, NULL)::product, NULL],
    now(),
    '2022-05-01 12:00:00-04'
);

-- copy the table to a parquet file
COPY product_example TO '/tmp/product_example.parquet' (format 'parquet', compression 'gzip');

-- show table
SELECT * FROM product_example;

-- copy the parquet file to the table
COPY product_example FROM '/tmp/product_example.parquet';

-- show table
SELECT * FROM product_example;
```



You can also use `COPY` command to read and write Parquet stream from/to standard input and output. Below is an example usage (you have to specify `format = parquet`):

```bash
psql -d pg_parquet -p 28817 -h localhost -c "create table product_example_reconstructed (like product_example);"
 CREATE TABLE

psql -d pg_parquet -p 28817 -h localhost -c "copy product_example to stdout (format parquet);" | psql -d pg_parquet -p 28817 -h localhost -c "copy product_example_reconstructed from stdin (format parquet);"
 COPY 2
```



### Inspect Parquet schema

You can call `SELECT * FROM parquet.schema(<uri>)` to discover the schema of the Parquet file at given uri.

```sql
SELECT * FROM parquet.schema('/tmp/product_example.parquet') LIMIT 10;
             uri              |     name     | type_name  | type_length | repetition_type | num_children | converted_type | scale | precision | field_id | logical_type 
------------------------------+--------------+------------+-------------+-----------------+--------------+----------------+-------+-----------+----------+--------------
 /tmp/product_example.parquet | arrow_schema |            |             |                 |            5 |                |       |           |          | 
 /tmp/product_example.parquet | id           | INT32      |             | OPTIONAL        |              |                |       |           |        0 | 
 /tmp/product_example.parquet | product      |            |             | OPTIONAL        |            3 |                |       |           |        1 | 
 /tmp/product_example.parquet | id           | INT32      |             | OPTIONAL        |              |                |       |           |        2 | 
 /tmp/product_example.parquet | name         | BYTE_ARRAY |             | OPTIONAL        |              | UTF8           |       |           |        3 | STRING
 /tmp/product_example.parquet | items        |            |             | OPTIONAL        |            1 | LIST           |       |           |        4 | LIST
 /tmp/product_example.parquet | list         |            |             | REPEATED        |            1 |                |       |           |          | 
 /tmp/product_example.parquet | element        |            |             | OPTIONAL        |            3 |                |       |           |        5 | 
 /tmp/product_example.parquet | id           | INT32      |             | OPTIONAL        |              |                |       |           |        6 | 
 /tmp/product_example.parquet | name         | BYTE_ARRAY |             | OPTIONAL        |              | UTF8           |       |           |        7 | STRING
(10 rows)
```



### Inspect Parquet metadata

You can call `SELECT * FROM parquet.metadata(<uri>)` to discover the detailed metadata of the Parquet file, such as column statistics, at given uri.

```
SELECT uri, row_group_id, row_group_num_rows, row_group_num_columns, row_group_bytes, column_id, file_offset, num_values, path_in_schema, type_name FROM parquet.metadata('/tmp/product_example.parquet') LIMIT 1;
             uri              | row_group_id | row_group_num_rows | row_group_num_columns | row_group_bytes | column_id | file_offset | num_values | path_in_schema | type_name 
------------------------------+--------------+--------------------+-----------------------+-----------------+-----------+-------------+------------+----------------+-----------
 /tmp/product_example.parquet |            0 |                  1 |                    13 |             842 |         0 |           0 |          1 | id             | INT32
(1 row)
```



```
SELECT stats_null_count, stats_distinct_count, stats_min, stats_max, compression, encodings, index_page_offset, dictionary_page_offset, data_page_offset, total_compressed_size, total_uncompressed_size FROM parquet.metadata('/tmp/product_example.parquet') LIMIT 1;
 stats_null_count | stats_distinct_count | stats_min | stats_max |    compression     |        encodings         | index_page_offset | dictionary_page_offset | data_page_offset | total_compressed_size | total_uncompressed_size 
------------------+----------------------+-----------+-----------+--------------------+--------------------------+-------------------+------------------------+------------------+-----------------------+-------------------------
                0 |                      | 1         | 1         | GZIP(GzipLevel(6)) | PLAIN,RLE,RLE_DICTIONARY |                   |                      4 |               42 |                   101 |                      61
(1 row)
```



You can call `SELECT * FROM parquet.file_metadata(<uri>)` to discover file level metadata of the Parquet file, such as format version, at given uri.

```
SELECT * FROM parquet.file_metadata('/tmp/product_example.parquet')
             uri              | created_by | num_rows | num_row_groups | format_version 
------------------------------+------------+----------+----------------+----------------
 /tmp/product_example.parquet | pg_parquet |        1 |              1 | 1
(1 row)
```



You can call `SELECT * FROM parquet.kv_metadata(<uri>)` to query custom key-value metadata of the Parquet file at given uri.

```
SELECT uri, encode(key, 'escape') as key, encode(value, 'escape') as value FROM parquet.kv_metadata('/tmp/product_example.parquet');
             uri              |     key      |    value
------------------------------+--------------+---------------------
 /tmp/product_example.parquet | ARROW:schema | /////5gIAAAQAAAA ...
(1 row)
```



### Inspect Parquet column statistics

You can call `SELECT * FROM parquet.column_stats(<uri>)` to discover the column statistics of the Parquet file, such as min and max value for the column, at given uri.

```
SELECT * FROM parquet.column_stats('/tmp/product_example.parquet')
 column_id | field_id |         stats_min          |         stats_max          | stats_null_count | stats_distinct_count 
-----------+----------+----------------------------+----------------------------+------------------+----------------------
         4 |        7 | item 1                     | item 2                     |                1 |                     
         6 |       11 | 1                          | 1                          |                1 |                     
         7 |       12 |                            |                            |                2 |                     
        10 |       17 |                            |                            |                2 |                     
         0 |        0 | 1                          | 1                          |                0 |                     
        11 |       18 | 2025-03-11 14:01:22.045739 | 2025-03-11 14:01:22.045739 |                0 |                     
         3 |        6 | 1                          | 2                          |                1 |                     
        12 |       19 | 2022-05-01 19:00:00+03     | 2022-05-01 19:00:00+03     |                0 |                     
         8 |       15 |                            |                            |                2 |                     
         5 |        8 | 1                          | 2                          |                1 |                     
         9 |       16 |                            |                            |                2 |                     
         1 |        2 | 1                          | 1                          |                0 |                     
         2 |        3 | product 1                  | product 1                  |                0 |                     
(13 rows)
```