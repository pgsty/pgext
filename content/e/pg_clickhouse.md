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
| **2460** | {{< badge content="pg_clickhouse" link="https://github.com/ClickHouse/pg_clickhouse" >}} | {{< ext "pg_clickhouse" >}} | `0.2.0` | {{< category "OLAP" >}} | {{< license "Apache-2.0" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_duckdb" >}} {{< ext "duckdb_fdw" >}} {{< ext "citus" >}} {{< ext "columnar" >}} {{< ext "citus_columnar" >}} {{< ext "clickhouse_fdw" >}} {{< ext "postgres_fdw" >}} {{< ext "dblink" >}} |

> [!Note] release 0.2.0; SQL v0.2


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_clickhouse` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "pg_clickhouse_18" "green" >}} {{< bg "17" "pg_clickhouse_17" "green" >}} {{< bg "16" "pg_clickhouse_16" "green" >}} {{< bg "15" "pg_clickhouse_15" "green" >}} {{< bg "14" "pg_clickhouse_14" "green" >}} | `pg_clickhouse_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "postgresql-18-clickhouse" "green" >}} {{< bg "17" "postgresql-17-clickhouse" "green" >}} {{< bg "16" "postgresql-16-clickhouse" "green" >}} {{< bg "15" "postgresql-15-clickhouse" "green" >}} {{< bg "14" "postgresql-14-clickhouse" "green" >}} | `postgresql-$v-clickhouse` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_clickhouse_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-clickhouse : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-clickhouse : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_18` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 724.2 KiB | [pg_clickhouse_18-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_18-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_18` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 645.0 KiB | [pg_clickhouse_18-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_18-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_18` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 729.8 KiB | [pg_clickhouse_18-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_18-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_18` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 699.4 KiB | [pg_clickhouse_18-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_18-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_18` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 753.2 KiB | [pg_clickhouse_18-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_18-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_18` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 717.0 KiB | [pg_clickhouse_18-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_18-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-clickhouse` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 859.5 KiB | [postgresql-18-clickhouse_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-clickhouse` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 796.3 KiB | [postgresql-18-clickhouse_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-clickhouse` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 863.3 KiB | [postgresql-18-clickhouse_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-clickhouse` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 801.1 KiB | [postgresql-18-clickhouse_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-clickhouse` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 927.8 KiB | [postgresql-18-clickhouse_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-clickhouse` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 900.1 KiB | [postgresql-18-clickhouse_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-clickhouse` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 939.7 KiB | [postgresql-18-clickhouse_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-clickhouse` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 906.6 KiB | [postgresql-18-clickhouse_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-18-clickhouse_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_17` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 723.9 KiB | [pg_clickhouse_17-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_17-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_17` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 644.9 KiB | [pg_clickhouse_17-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_17-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_17` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 729.6 KiB | [pg_clickhouse_17-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_17-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_17` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 699.5 KiB | [pg_clickhouse_17-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_17-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_17` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 753.4 KiB | [pg_clickhouse_17-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_17-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_17` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 716.5 KiB | [pg_clickhouse_17-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_17-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-clickhouse` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 859.2 KiB | [postgresql-17-clickhouse_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-clickhouse` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 796.0 KiB | [postgresql-17-clickhouse_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-clickhouse` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 862.4 KiB | [postgresql-17-clickhouse_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-clickhouse` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 801.0 KiB | [postgresql-17-clickhouse_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-clickhouse` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 981.5 KiB | [postgresql-17-clickhouse_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-clickhouse` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 953.0 KiB | [postgresql-17-clickhouse_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-clickhouse` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 939.6 KiB | [postgresql-17-clickhouse_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-clickhouse` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 906.4 KiB | [postgresql-17-clickhouse_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-17-clickhouse_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_16` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 723.7 KiB | [pg_clickhouse_16-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_16-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_16` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 644.8 KiB | [pg_clickhouse_16-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_16-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_16` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 729.7 KiB | [pg_clickhouse_16-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_16-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_16` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 699.4 KiB | [pg_clickhouse_16-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_16-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_16` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 753.1 KiB | [pg_clickhouse_16-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_16-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_16` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 716.7 KiB | [pg_clickhouse_16-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_16-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-clickhouse` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 859.0 KiB | [postgresql-16-clickhouse_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-clickhouse` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 795.9 KiB | [postgresql-16-clickhouse_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-clickhouse` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 862.4 KiB | [postgresql-16-clickhouse_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-clickhouse` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 800.9 KiB | [postgresql-16-clickhouse_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-clickhouse` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 978.6 KiB | [postgresql-16-clickhouse_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-clickhouse` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 950.3 KiB | [postgresql-16-clickhouse_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-clickhouse` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 939.4 KiB | [postgresql-16-clickhouse_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-clickhouse` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 906.0 KiB | [postgresql-16-clickhouse_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-16-clickhouse_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_15` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 727.4 KiB | [pg_clickhouse_15-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_15-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_15` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 649.8 KiB | [pg_clickhouse_15-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_15-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_15` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 732.9 KiB | [pg_clickhouse_15-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_15-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_15` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 702.5 KiB | [pg_clickhouse_15-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_15-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_15` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 757.1 KiB | [pg_clickhouse_15-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_15-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_15` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 720.2 KiB | [pg_clickhouse_15-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_15-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-clickhouse` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 861.2 KiB | [postgresql-15-clickhouse_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-clickhouse` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 797.7 KiB | [postgresql-15-clickhouse_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-clickhouse` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 864.7 KiB | [postgresql-15-clickhouse_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-clickhouse` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 802.1 KiB | [postgresql-15-clickhouse_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-clickhouse` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 980.5 KiB | [postgresql-15-clickhouse_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-clickhouse` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 950.9 KiB | [postgresql-15-clickhouse_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-clickhouse` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 942.3 KiB | [postgresql-15-clickhouse_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-clickhouse` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 909.1 KiB | [postgresql-15-clickhouse_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-15-clickhouse_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_clickhouse_14` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 727.3 KiB | [pg_clickhouse_14-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_clickhouse_14-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_clickhouse_14` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 649.7 KiB | [pg_clickhouse_14-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_clickhouse_14-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_clickhouse_14` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 732.7 KiB | [pg_clickhouse_14-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_clickhouse_14-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_clickhouse_14` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 702.4 KiB | [pg_clickhouse_14-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_clickhouse_14-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_clickhouse_14` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 757.1 KiB | [pg_clickhouse_14-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_clickhouse_14-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_clickhouse_14` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 720.3 KiB | [pg_clickhouse_14-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_clickhouse_14-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-clickhouse` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 861.2 KiB | [postgresql-14-clickhouse_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-clickhouse` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 798.1 KiB | [postgresql-14-clickhouse_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-clickhouse` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 864.2 KiB | [postgresql-14-clickhouse_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-clickhouse` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 802.5 KiB | [postgresql-14-clickhouse_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-clickhouse` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 979.6 KiB | [postgresql-14-clickhouse_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-clickhouse` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 950.1 KiB | [postgresql-14-clickhouse_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-clickhouse` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 942.4 KiB | [postgresql-14-clickhouse_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-clickhouse` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 909.1 KiB | [postgresql-14-clickhouse_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-clickhouse/postgresql-14-clickhouse_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ClickHouse/pg_clickhouse" title="Repository" icon="github" subtitle="github.com/ClickHouse/pg_clickhouse" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_clickhouse-0.2.0.tar.gz" >}}
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

Sources: [README](https://github.com/ClickHouse/pg_clickhouse/blob/master/README.md), [reference](https://github.com/ClickHouse/pg_clickhouse/blob/master/doc/pg_clickhouse.md), [tutorial](https://github.com/ClickHouse/pg_clickhouse/blob/master/doc/tutorial.md), [v0.2.0 release notes](https://github.com/ClickHouse/pg_clickhouse/releases/tag/v0.2.0)

`pg_clickhouse` runs analytics queries on ClickHouse from PostgreSQL through the `clickhouse_fdw` foreign data wrapper. Upstream documents PostgreSQL 13+ and ClickHouse 23+ support.

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

User mapping options:

- `user`
- `password`

### Common operations

```sql
ALTER EXTENSION pg_clickhouse UPDATE;
ALTER EXTENSION pg_clickhouse UPDATE TO '0.2';
DROP SERVER taxi_srv CASCADE;
```

`IMPORT FOREIGN SCHEMA` also supports `LIMIT TO (...)` and `EXCEPT (...)`. The reference warns that imported mixed-case identifiers are double-quoted in PostgreSQL and must be queried with quotes.

### Version and pushdown notes

- The reference documents separate library and extension versions; `pgch_version()` was added in release `v0.2.0`.
- Patch-only releases update the library without requiring `ALTER EXTENSION`.
- Release `v0.2.0` added more pushdown for arrays, regex functions, `split_part()`, array operators, and current date/time expressions, plus the `pg_clickhouse.pushdown_regex` setting.

### Caveats

- This is positioned upstream as an analytics-first extension; the roadmap still lists broader DML support as future work.
- For full examples, follow the official tutorial, which creates a ClickHouse `taxi` database, imports it through `IMPORT FOREIGN SCHEMA`, and queries the resulting foreign tables.
