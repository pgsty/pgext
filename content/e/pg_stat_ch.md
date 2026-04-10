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
| **6020** | {{< badge content="pg_stat_ch" link="https://github.com/ClickHouse/pg_stat_ch" >}} | {{< ext "pg_stat_ch" >}} | `0.3.3` | {{< category "STAT" >}} | {{< license "Apache-2.0" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_tracing" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_stat_kcache" >}} {{< ext "powa" >}} |

> [!Note] Requires shared_preload_libraries = pg_stat_ch; README recommends track_io_timing = on.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_stat_ch` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.3` | {{< bg "18" "pg_stat_ch_18" "green" >}} {{< bg "17" "pg_stat_ch_17" "green" >}} {{< bg "16" "pg_stat_ch_16" "green" >}} {{< bg "15" "pg_stat_ch_15" "red" >}} {{< bg "14" "pg_stat_ch_14" "red" >}} | `pg_stat_ch_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.3` | {{< bg "18" "postgresql-18-pg-stat-ch" "green" >}} {{< bg "17" "postgresql-17-pg-stat-ch" "green" >}} {{< bg "16" "postgresql-16-pg-stat-ch" "green" >}} {{< bg "15" "postgresql-15-pg-stat-ch" "red" >}} {{< bg "14" "postgresql-14-pg-stat-ch" "red" >}} | `postgresql-$v-pg-stat-ch` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pg_stat_ch_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pg_stat_ch_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.3" "pg_stat_ch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_stat_ch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_stat_ch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_stat_ch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.3" "pg_stat_ch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_stat_ch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_stat_ch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_stat_ch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.3" "pg_stat_ch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_stat_ch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_stat_ch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_stat_ch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.3" "pg_stat_ch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_stat_ch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_stat_ch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_stat_ch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_ch_18` | `0.3.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_stat_ch_18-0.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_ch_18-0.3.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_stat_ch_18` | `0.3.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [pg_stat_ch_18-0.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_ch_18-0.3.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_stat_ch_18` | `0.3.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.0 MiB | [pg_stat_ch_18-0.3.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_ch_18-0.3.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_stat_ch_18` | `0.3.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 994.3 KiB | [pg_stat_ch_18-0.3.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_ch_18-0.3.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-stat-ch` | `0.3.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 962.0 KiB | [postgresql-18-pg-stat-ch_0.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 865.7 KiB | [postgresql-18-pg-stat-ch_0.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 987.7 KiB | [postgresql-18-pg-stat-ch_0.3.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 885.3 KiB | [postgresql-18-pg-stat-ch_0.3.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.6 MiB | [postgresql-18-pg-stat-ch_0.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.3 MiB | [postgresql-18-pg-stat-ch_0.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 922.1 KiB | [postgresql-18-pg-stat-ch_0.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 880.8 KiB | [postgresql-18-pg-stat-ch_0.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_ch_17` | `0.3.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_stat_ch_17-0.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_ch_17-0.3.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_stat_ch_17` | `0.3.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [pg_stat_ch_17-0.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_ch_17-0.3.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_stat_ch_17` | `0.3.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.0 MiB | [pg_stat_ch_17-0.3.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_ch_17-0.3.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_stat_ch_17` | `0.3.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 991.6 KiB | [pg_stat_ch_17-0.3.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_ch_17-0.3.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-stat-ch` | `0.3.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 961.6 KiB | [postgresql-17-pg-stat-ch_0.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 865.8 KiB | [postgresql-17-pg-stat-ch_0.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 988.7 KiB | [postgresql-17-pg-stat-ch_0.3.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 884.0 KiB | [postgresql-17-pg-stat-ch_0.3.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.6 MiB | [postgresql-17-pg-stat-ch_0.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.3 MiB | [postgresql-17-pg-stat-ch_0.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 922.5 KiB | [postgresql-17-pg-stat-ch_0.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 879.8 KiB | [postgresql-17-pg-stat-ch_0.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_ch_16` | `0.3.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_stat_ch_16-0.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_ch_16-0.3.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_stat_ch_16` | `0.3.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [pg_stat_ch_16-0.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_ch_16-0.3.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_stat_ch_16` | `0.3.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.0 MiB | [pg_stat_ch_16-0.3.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_ch_16-0.3.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_stat_ch_16` | `0.3.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 994.6 KiB | [pg_stat_ch_16-0.3.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_ch_16-0.3.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-stat-ch` | `0.3.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 961.9 KiB | [postgresql-16-pg-stat-ch_0.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 865.5 KiB | [postgresql-16-pg-stat-ch_0.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 989.7 KiB | [postgresql-16-pg-stat-ch_0.3.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 884.8 KiB | [postgresql-16-pg-stat-ch_0.3.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.6 MiB | [postgresql-16-pg-stat-ch_0.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.3 MiB | [postgresql-16-pg-stat-ch_0.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 921.8 KiB | [postgresql-16-pg-stat-ch_0.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 879.5 KiB | [postgresql-16-pg-stat-ch_0.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ClickHouse/pg_stat_ch" title="Repository" icon="github" subtitle="github.com/ClickHouse/pg_stat_ch" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_stat_ch-0.3.3-vendor.tar.gz" >}}
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
