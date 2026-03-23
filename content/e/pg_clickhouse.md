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
| **2460** | {{< badge content="pg_clickhouse" link="https://github.com/ClickHouse/pg_clickhouse" >}} | {{< ext "pg_clickhouse" >}} | `0.1.5` | {{< category "OLAP" >}} | {{< license "Apache-2.0" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_duckdb" >}} {{< ext "duckdb_fdw" >}} {{< ext "citus" >}} {{< ext "columnar" >}} {{< ext "citus_columnar" >}} {{< ext "clickhouse_fdw" >}} {{< ext "postgres_fdw" >}} {{< ext "dblink" >}} |

> [!Note] with submodule


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_clickhouse` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.5` | {{< bg "18" "pg_clickhouse_18" "green" >}} {{< bg "17" "pg_clickhouse_17" "green" >}} {{< bg "16" "pg_clickhouse_16" "green" >}} {{< bg "15" "pg_clickhouse_15" "green" >}} {{< bg "14" "pg_clickhouse_14" "green" >}} | `pg_clickhouse_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.5` | {{< bg "18" "postgresql-18-clickhouse" "green" >}} {{< bg "17" "postgresql-17-clickhouse" "green" >}} {{< bg "16" "postgresql-16-clickhouse" "green" >}} {{< bg "15" "postgresql-15-clickhouse" "green" >}} {{< bg "14" "postgresql-14-clickhouse" "green" >}} | `postgresql-$v-clickhouse` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_18` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 705.7 KiB | [pg_clickhouse_18-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_18-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_18` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 629.7 KiB | [pg_clickhouse_18-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_18-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_18` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 712.5 KiB | [pg_clickhouse_18-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_18-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_18` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 682.5 KiB | [pg_clickhouse_18-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_18-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_18` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 736.6 KiB | [pg_clickhouse_18-0.1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_18-0.1.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_18` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 700.5 KiB | [pg_clickhouse_18-0.1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_18-0.1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-clickhouse` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 820.2 KiB | [postgresql-18-clickhouse_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-clickhouse` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 759.8 KiB | [postgresql-18-clickhouse_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-clickhouse` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 824.1 KiB | [postgresql-18-clickhouse_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-clickhouse` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 763.8 KiB | [postgresql-18-clickhouse_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-clickhouse` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 889.3 KiB | [postgresql-18-clickhouse_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-clickhouse` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 862.3 KiB | [postgresql-18-clickhouse_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-clickhouse` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 902.5 KiB | [postgresql-18-clickhouse_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-clickhouse` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 870.3 KiB | [postgresql-18-clickhouse_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_17` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 705.4 KiB | [pg_clickhouse_17-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_17-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_17` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 629.6 KiB | [pg_clickhouse_17-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_17-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_17` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 712.5 KiB | [pg_clickhouse_17-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_17-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_17` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 682.0 KiB | [pg_clickhouse_17-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_17-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_17` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 736.5 KiB | [pg_clickhouse_17-0.1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_17-0.1.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_17` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 700.2 KiB | [pg_clickhouse_17-0.1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_17-0.1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-clickhouse` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 820.1 KiB | [postgresql-17-clickhouse_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-clickhouse` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 758.5 KiB | [postgresql-17-clickhouse_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-clickhouse` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 823.2 KiB | [postgresql-17-clickhouse_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-clickhouse` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 763.1 KiB | [postgresql-17-clickhouse_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-clickhouse` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 938.5 KiB | [postgresql-17-clickhouse_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-clickhouse` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 910.5 KiB | [postgresql-17-clickhouse_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-clickhouse` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 902.0 KiB | [postgresql-17-clickhouse_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-clickhouse` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 870.2 KiB | [postgresql-17-clickhouse_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_16` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 705.5 KiB | [pg_clickhouse_16-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_16-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_16` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 629.6 KiB | [pg_clickhouse_16-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_16-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_16` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 712.6 KiB | [pg_clickhouse_16-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_16-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_16` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 682.3 KiB | [pg_clickhouse_16-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_16-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_16` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 736.4 KiB | [pg_clickhouse_16-0.1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_16-0.1.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_16` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 700.3 KiB | [pg_clickhouse_16-0.1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_16-0.1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-clickhouse` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 820.0 KiB | [postgresql-16-clickhouse_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-clickhouse` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 758.8 KiB | [postgresql-16-clickhouse_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-clickhouse` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 823.9 KiB | [postgresql-16-clickhouse_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-clickhouse` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 763.2 KiB | [postgresql-16-clickhouse_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-clickhouse` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 935.4 KiB | [postgresql-16-clickhouse_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-clickhouse` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 907.3 KiB | [postgresql-16-clickhouse_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-clickhouse` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 901.6 KiB | [postgresql-16-clickhouse_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-clickhouse` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 870.0 KiB | [postgresql-16-clickhouse_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_15` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 708.4 KiB | [pg_clickhouse_15-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_15-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_15` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 631.3 KiB | [pg_clickhouse_15-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_15-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_15` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 715.2 KiB | [pg_clickhouse_15-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_15-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_15` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 684.5 KiB | [pg_clickhouse_15-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_15-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_15` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 740.2 KiB | [pg_clickhouse_15-0.1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_15-0.1.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_15` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 703.3 KiB | [pg_clickhouse_15-0.1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_15-0.1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-clickhouse` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 821.7 KiB | [postgresql-15-clickhouse_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-clickhouse` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 759.8 KiB | [postgresql-15-clickhouse_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-clickhouse` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 825.0 KiB | [postgresql-15-clickhouse_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-clickhouse` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 764.5 KiB | [postgresql-15-clickhouse_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-clickhouse` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 937.4 KiB | [postgresql-15-clickhouse_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-clickhouse` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 909.0 KiB | [postgresql-15-clickhouse_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-clickhouse` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 904.2 KiB | [postgresql-15-clickhouse_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-clickhouse` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 871.0 KiB | [postgresql-15-clickhouse_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_14` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 708.6 KiB | [pg_clickhouse_14-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_14-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_14` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 631.4 KiB | [pg_clickhouse_14-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_14-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_14` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 715.0 KiB | [pg_clickhouse_14-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_14-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_14` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 684.4 KiB | [pg_clickhouse_14-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_14-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_14` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 740.3 KiB | [pg_clickhouse_14-0.1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_14-0.1.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_14` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 703.2 KiB | [pg_clickhouse_14-0.1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_14-0.1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-clickhouse` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 821.9 KiB | [postgresql-14-clickhouse_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-clickhouse` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 759.4 KiB | [postgresql-14-clickhouse_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-clickhouse` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 825.2 KiB | [postgresql-14-clickhouse_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-clickhouse` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 764.4 KiB | [postgresql-14-clickhouse_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-clickhouse` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 937.1 KiB | [postgresql-14-clickhouse_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-clickhouse` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 908.7 KiB | [postgresql-14-clickhouse_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-clickhouse` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 903.9 KiB | [postgresql-14-clickhouse_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-clickhouse` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 871.1 KiB | [postgresql-14-clickhouse_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ClickHouse/pg_clickhouse" title="Repository" icon="github" subtitle="github.com/ClickHouse/pg_clickhouse" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_clickhouse-0.1.5.tar.gz" >}}
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

> [pg_clickhouse: ClickHouse integration for PostgreSQL](https://github.com/ClickHouse/pg_clickhouse)

`pg_clickhouse` enables analytics queries on ClickHouse directly from PostgreSQL without SQL rewriting.
It supports PostgreSQL 13+ and ClickHouse v23+.

### Create the Extension

```sql
CREATE EXTENSION pg_clickhouse;
```

Or into a specific schema:

```sql
CREATE SCHEMA env;
CREATE EXTENSION pg_clickhouse SCHEMA env;
```

### Query Pushdown

The extension automatically pushes down analytical queries to ClickHouse for execution,
providing significant performance improvements. For example, TPC-H benchmarks show:

- Query 1: 268ms (vs 4,693ms on standard PostgreSQL)
- Query 6: 53ms (vs 764ms on standard PostgreSQL)

When query pushdown is active, ClickHouse handles execution directly, avoiding
data transfer overhead for complex analytical workloads.
