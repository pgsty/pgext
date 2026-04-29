---
title: "pg_stat_ch"
linkTitle: "pg_stat_ch"
description: "Export PostgreSQL query telemetry to ClickHouse"
weight: 6020
categories: ["STAT"]
width: full
---

[**pg_stat_ch**](https://github.com/ClickHouse/pg_stat_ch) : Export PostgreSQL query telemetry to ClickHouse


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6020** | {{< badge content="pg_stat_ch" link="https://github.com/ClickHouse/pg_stat_ch" >}} | {{< ext "pg_stat_ch" >}} | `0.3.6` | {{< category "STAT" >}} | {{< license "Apache-2.0" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_tracing" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_stat_kcache" >}} {{< ext "powa" >}} |

> [!Note] release 0.3.6; SQL v0.1


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.6` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_stat_ch` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.6` | {{< bg "18" "pg_stat_ch_18" "green" >}} {{< bg "17" "pg_stat_ch_17" "green" >}} {{< bg "16" "pg_stat_ch_16" "green" >}} {{< bg "15" "pg_stat_ch_15" "red" >}} {{< bg "14" "pg_stat_ch_14" "red" >}} | `pg_stat_ch_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.6` | {{< bg "18" "postgresql-18-pg-stat-ch" "green" >}} {{< bg "17" "postgresql-17-pg-stat-ch" "green" >}} {{< bg "16" "postgresql-16-pg-stat-ch" "green" >}} {{< bg "15" "postgresql-15-pg-stat-ch" "red" >}} {{< bg "14" "postgresql-14-pg-stat-ch" "red" >}} | `postgresql-$v-pg-stat-ch` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pg_stat_ch_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pg_stat_ch_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.6" "pg_stat_ch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "pg_stat_ch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "pg_stat_ch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_stat_ch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.6" "pg_stat_ch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "pg_stat_ch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "pg_stat_ch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_stat_ch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.6" "pg_stat_ch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "pg_stat_ch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "pg_stat_ch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_stat_ch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.6" "pg_stat_ch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "pg_stat_ch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "pg_stat_ch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_stat_ch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.6" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_ch_18` | `0.3.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 871.9 KiB | [pg_stat_ch_18-0.3.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_ch_18-0.3.6-1PIGSTY.el9.x86_64.rpm) |
| `pg_stat_ch_18` | `0.3.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 826.6 KiB | [pg_stat_ch_18-0.3.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_ch_18-0.3.6-1PIGSTY.el9.aarch64.rpm) |
| `pg_stat_ch_18` | `0.3.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 819.7 KiB | [pg_stat_ch_18-0.3.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_ch_18-0.3.6-1PIGSTY.el10.x86_64.rpm) |
| `pg_stat_ch_18` | `0.3.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 773.7 KiB | [pg_stat_ch_18-0.3.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_ch_18-0.3.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-stat-ch` | `0.3.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 720.3 KiB | [postgresql-18-pg-stat-ch_0.3.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 649.5 KiB | [postgresql-18-pg-stat-ch_0.3.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 731.0 KiB | [postgresql-18-pg-stat-ch_0.3.6-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.6-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 656.1 KiB | [postgresql-18-pg-stat-ch_0.3.6-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.6-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.4 MiB | [postgresql-18-pg-stat-ch_0.3.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.1 MiB | [postgresql-18-pg-stat-ch_0.3.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 716.9 KiB | [postgresql-18-pg-stat-ch_0.3.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 683.3 KiB | [postgresql-18-pg-stat-ch_0.3.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_ch_17` | `0.3.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 871.4 KiB | [pg_stat_ch_17-0.3.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_ch_17-0.3.6-1PIGSTY.el9.x86_64.rpm) |
| `pg_stat_ch_17` | `0.3.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 826.5 KiB | [pg_stat_ch_17-0.3.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_ch_17-0.3.6-1PIGSTY.el9.aarch64.rpm) |
| `pg_stat_ch_17` | `0.3.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 822.2 KiB | [pg_stat_ch_17-0.3.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_ch_17-0.3.6-1PIGSTY.el10.x86_64.rpm) |
| `pg_stat_ch_17` | `0.3.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 773.0 KiB | [pg_stat_ch_17-0.3.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_ch_17-0.3.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-stat-ch` | `0.3.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 719.4 KiB | [postgresql-17-pg-stat-ch_0.3.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 648.1 KiB | [postgresql-17-pg-stat-ch_0.3.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 730.8 KiB | [postgresql-17-pg-stat-ch_0.3.6-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.6-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 655.8 KiB | [postgresql-17-pg-stat-ch_0.3.6-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.6-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.4 MiB | [postgresql-17-pg-stat-ch_0.3.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.1 MiB | [postgresql-17-pg-stat-ch_0.3.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 716.7 KiB | [postgresql-17-pg-stat-ch_0.3.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 683.1 KiB | [postgresql-17-pg-stat-ch_0.3.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_ch_16` | `0.3.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 871.4 KiB | [pg_stat_ch_16-0.3.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_ch_16-0.3.6-1PIGSTY.el9.x86_64.rpm) |
| `pg_stat_ch_16` | `0.3.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 828.7 KiB | [pg_stat_ch_16-0.3.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_ch_16-0.3.6-1PIGSTY.el9.aarch64.rpm) |
| `pg_stat_ch_16` | `0.3.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 822.4 KiB | [pg_stat_ch_16-0.3.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_ch_16-0.3.6-1PIGSTY.el10.x86_64.rpm) |
| `pg_stat_ch_16` | `0.3.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 773.4 KiB | [pg_stat_ch_16-0.3.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_ch_16-0.3.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-stat-ch` | `0.3.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 719.6 KiB | [postgresql-16-pg-stat-ch_0.3.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 649.4 KiB | [postgresql-16-pg-stat-ch_0.3.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 729.5 KiB | [postgresql-16-pg-stat-ch_0.3.6-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.6-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 656.6 KiB | [postgresql-16-pg-stat-ch_0.3.6-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.6-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.4 MiB | [postgresql-16-pg-stat-ch_0.3.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.1 MiB | [postgresql-16-pg-stat-ch_0.3.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 716.4 KiB | [postgresql-16-pg-stat-ch_0.3.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 681.9 KiB | [postgresql-16-pg-stat-ch_0.3.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ClickHouse/pg_stat_ch" title="Repository" icon="github" subtitle="github.com/ClickHouse/pg_stat_ch" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_stat_ch-0.3.6.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_stat_ch;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_stat_ch;		# install via package name, for the active PG version

pig install pg_stat_ch -v 18;   # install for PG 18
pig install pg_stat_ch -v 17;   # install for PG 17
pig install pg_stat_ch -v 16;   # install for PG 16

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_stat_ch';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_stat_ch;
```


## Usage

Sources: [README](https://github.com/ClickHouse/pg_stat_ch/blob/main/README.md), [release 0.3.6](https://github.com/ClickHouse/pg_stat_ch/releases/tag/v0.3.6)

`pg_stat_ch` captures per-query PostgreSQL execution telemetry and exports raw events to ClickHouse through a shared-memory queue and background worker. The upstream project positions it as a raw-event alternative to `pg_stat_statements`, with aggregation and dashboards handled in ClickHouse instead of PostgreSQL.

```sql
CREATE EXTENSION pg_stat_ch;
```

### Required Setup

`pg_stat_ch` must be preloaded and pointed at a ClickHouse database:

```conf
shared_preload_libraries = 'pg_stat_ch'
track_io_timing = on

pg_stat_ch.clickhouse_host = 'localhost'
pg_stat_ch.clickhouse_port = 9000
pg_stat_ch.clickhouse_database = 'pg_stat_ch'
pg_stat_ch.clickhouse_use_tls = on
pg_stat_ch.clickhouse_skip_tls_verify = off
```

The README says PostgreSQL 16, 17, and 18 are fully supported; the latest release is `0.3.6` from April 15, 2026.

### SQL API

- `pg_stat_ch_version()` returns the extension version.
- `pg_stat_ch_stats()` exposes queue and exporter counters.
- `pg_stat_ch_reset()` clears queue counters.
- `pg_stat_ch_flush()` triggers an immediate export flush.

```sql
SELECT pg_stat_ch_version();
SELECT * FROM pg_stat_ch_stats();
SELECT pg_stat_ch_flush();
```

### Important GUCs

- `pg_stat_ch.enabled` toggles collection.
- `pg_stat_ch.queue_capacity` and `pg_stat_ch.string_area_size` size the shared-memory buffers.
- `pg_stat_ch.flush_interval_ms` and `pg_stat_ch.batch_max` control exporter cadence and batch size.
- `pg_stat_ch.log_min_elevel` controls which errors are captured.

### What Gets Captured

- Query timing, row counts, buffer usage, WAL usage, and CPU time.
- DML, DDL, and utility statements.
- SQLSTATE and error levels.
- JIT metrics on PostgreSQL 15+.
- Parallel worker stats on PostgreSQL 18+.
- Application name, client IP, and query text up to the upstream truncation limit.

### Caveats

- The design intentionally drops events on queue overflow instead of blocking the foreground query path.
- ClickHouse schema creation is a required part of setup; upstream quickstart scripts preload it automatically, but manual deployments must load the schema separately.
