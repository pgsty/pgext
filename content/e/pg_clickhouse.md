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
| **2460** | {{< badge content="pg_clickhouse" link="https://github.com/ClickHouse/pg_clickhouse" >}} | {{< ext "pg_clickhouse" >}} | `0.1.0` | {{< category "OLAP" >}} | {{< license "Apache-2.0" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_duckdb" >}} {{< ext "duckdb_fdw" >}} {{< ext "citus" >}} {{< ext "columnar" >}} {{< ext "citus_columnar" >}} {{< ext "clickhouse_fdw" >}} {{< ext "postgres_fdw" >}} {{< ext "dblink" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_clickhouse` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pg_clickhouse_18*" "green" >}} {{< bg "17" "pg_clickhouse_17*" "green" >}} {{< bg "16" "pg_clickhouse_16*" "green" >}} {{< bg "15" "pg_clickhouse_15*" "green" >}} {{< bg "14" "pg_clickhouse_14*" "green" >}} {{< bg "13" "pg_clickhouse_13*" "green" >}} | `pg_clickhouse_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-clickhouse" "green" >}} {{< bg "17" "postgresql-17-clickhouse" "green" >}} {{< bg "16" "postgresql-16-clickhouse" "green" >}} {{< bg "15" "postgresql-15-clickhouse" "green" >}} {{< bg "14" "postgresql-14-clickhouse" "green" >}} {{< bg "13" "postgresql-13-clickhouse" "green" >}} | `postgresql-$v-clickhouse` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_clickhouse_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-13-clickhouse : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-13-clickhouse : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-13-clickhouse : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-13-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-13-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-13-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-13-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-13-clickhouse : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 699.7 KiB | [pg_clickhouse_18-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_18-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 624.2 KiB | [pg_clickhouse_18-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_18-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 705.0 KiB | [pg_clickhouse_18-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_18-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 675.5 KiB | [pg_clickhouse_18-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_18-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 723.8 KiB | [pg_clickhouse_18-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_18-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 688.8 KiB | [pg_clickhouse_18-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_18-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-clickhouse` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 807.5 KiB | [postgresql-18-clickhouse_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-clickhouse` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 744.7 KiB | [postgresql-18-clickhouse_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-clickhouse` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 809.7 KiB | [postgresql-18-clickhouse_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-clickhouse` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 749.3 KiB | [postgresql-18-clickhouse_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-clickhouse` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 876.7 KiB | [postgresql-18-clickhouse_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-clickhouse` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 847.1 KiB | [postgresql-18-clickhouse_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-clickhouse` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 887.0 KiB | [postgresql-18-clickhouse_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-clickhouse` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 855.9 KiB | [postgresql-18-clickhouse_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 699.5 KiB | [pg_clickhouse_17-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_17-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 623.9 KiB | [pg_clickhouse_17-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_17-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 705.0 KiB | [pg_clickhouse_17-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_17-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 675.5 KiB | [pg_clickhouse_17-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_17-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 723.6 KiB | [pg_clickhouse_17-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_17-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 688.6 KiB | [pg_clickhouse_17-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_17-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-clickhouse` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 806.2 KiB | [postgresql-17-clickhouse_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-clickhouse` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 744.0 KiB | [postgresql-17-clickhouse_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-clickhouse` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 809.5 KiB | [postgresql-17-clickhouse_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-clickhouse` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 748.5 KiB | [postgresql-17-clickhouse_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-clickhouse` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 925.5 KiB | [postgresql-17-clickhouse_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-clickhouse` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 895.0 KiB | [postgresql-17-clickhouse_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-clickhouse` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 886.7 KiB | [postgresql-17-clickhouse_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-clickhouse` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 854.9 KiB | [postgresql-17-clickhouse_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 699.6 KiB | [pg_clickhouse_16-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_16-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 624.0 KiB | [pg_clickhouse_16-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_16-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 704.9 KiB | [pg_clickhouse_16-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_16-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 675.4 KiB | [pg_clickhouse_16-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_16-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 723.6 KiB | [pg_clickhouse_16-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_16-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 688.7 KiB | [pg_clickhouse_16-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_16-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-clickhouse` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 806.2 KiB | [postgresql-16-clickhouse_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-clickhouse` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 744.8 KiB | [postgresql-16-clickhouse_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-clickhouse` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 811.1 KiB | [postgresql-16-clickhouse_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-clickhouse` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 748.5 KiB | [postgresql-16-clickhouse_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-clickhouse` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 922.0 KiB | [postgresql-16-clickhouse_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-clickhouse` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 891.9 KiB | [postgresql-16-clickhouse_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-clickhouse` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 886.7 KiB | [postgresql-16-clickhouse_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-clickhouse` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 855.0 KiB | [postgresql-16-clickhouse_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_15` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 701.6 KiB | [pg_clickhouse_15-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_15-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_15` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 624.9 KiB | [pg_clickhouse_15-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_15-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_15` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 705.8 KiB | [pg_clickhouse_15-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_15-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_15` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 676.7 KiB | [pg_clickhouse_15-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_15-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_15` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 726.1 KiB | [pg_clickhouse_15-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_15-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_15` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 691.6 KiB | [pg_clickhouse_15-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_15-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-clickhouse` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 807.8 KiB | [postgresql-15-clickhouse_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-clickhouse` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 745.9 KiB | [postgresql-15-clickhouse_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-clickhouse` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 810.5 KiB | [postgresql-15-clickhouse_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-clickhouse` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 749.4 KiB | [postgresql-15-clickhouse_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-clickhouse` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 923.5 KiB | [postgresql-15-clickhouse_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-clickhouse` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 893.2 KiB | [postgresql-15-clickhouse_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-clickhouse` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 888.3 KiB | [postgresql-15-clickhouse_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-clickhouse` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 857.1 KiB | [postgresql-15-clickhouse_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_14` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 701.6 KiB | [pg_clickhouse_14-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_14-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_14` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 624.9 KiB | [pg_clickhouse_14-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_14-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_14` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 705.6 KiB | [pg_clickhouse_14-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_14-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_14` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 676.4 KiB | [pg_clickhouse_14-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_14-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_14` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 726.0 KiB | [pg_clickhouse_14-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_14-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_14` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 691.4 KiB | [pg_clickhouse_14-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_14-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-clickhouse` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 807.4 KiB | [postgresql-14-clickhouse_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-clickhouse` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 745.2 KiB | [postgresql-14-clickhouse_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-clickhouse` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 810.5 KiB | [postgresql-14-clickhouse_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-clickhouse` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 749.7 KiB | [postgresql-14-clickhouse_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-clickhouse` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 923.2 KiB | [postgresql-14-clickhouse_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-clickhouse` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 892.9 KiB | [postgresql-14-clickhouse_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-clickhouse` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 888.2 KiB | [postgresql-14-clickhouse_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-clickhouse` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 856.5 KiB | [postgresql-14-clickhouse_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_13` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 701.0 KiB | [pg_clickhouse_13-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_13-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_13` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 625.0 KiB | [pg_clickhouse_13-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_13-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_13` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 706.2 KiB | [pg_clickhouse_13-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_13-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_13` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 676.9 KiB | [pg_clickhouse_13-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_13-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_13` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 726.0 KiB | [pg_clickhouse_13-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_13-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_13` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 691.8 KiB | [pg_clickhouse_13-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_13-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-clickhouse` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 807.8 KiB | [postgresql-13-clickhouse_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-13-clickhouse_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-clickhouse` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 745.0 KiB | [postgresql-13-clickhouse_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-13-clickhouse_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-clickhouse` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 810.4 KiB | [postgresql-13-clickhouse_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-13-clickhouse_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-clickhouse` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 749.2 KiB | [postgresql-13-clickhouse_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-13-clickhouse_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-clickhouse` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 922.1 KiB | [postgresql-13-clickhouse_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-13-clickhouse_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-clickhouse` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 892.5 KiB | [postgresql-13-clickhouse_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-13-clickhouse_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-clickhouse` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 888.0 KiB | [postgresql-13-clickhouse_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-13-clickhouse_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-clickhouse` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 856.4 KiB | [postgresql-13-clickhouse_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-13-clickhouse_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ClickHouse/pg_clickhouse" title="Repository" icon="github" subtitle="github.com/ClickHouse/pg_clickhouse" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_clickhouse-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_clickhouse;		# build rpm / deb with pig
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
pig install pg_clickhouse -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_clickhouse;
```
