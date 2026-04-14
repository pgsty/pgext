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
| **6020** | {{< badge content="pg_stat_ch" link="https://github.com/ClickHouse/pg_stat_ch" >}} | {{< ext "pg_stat_ch" >}} | `0.3.4` | {{< category "STAT" >}} | {{< license "Apache-2.0" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_tracing" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_stat_kcache" >}} {{< ext "powa" >}} |

> [!Note] Release tag 0.3.4 still ships extension SQL version 0.1 by upstream design; source tarball bundles third-party dependencies.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_stat_ch` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.4` | {{< bg "18" "pg_stat_ch_18" "green" >}} {{< bg "17" "pg_stat_ch_17" "green" >}} {{< bg "16" "pg_stat_ch_16" "green" >}} {{< bg "15" "pg_stat_ch_15" "red" >}} {{< bg "14" "pg_stat_ch_14" "red" >}} | `pg_stat_ch_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.4` | {{< bg "18" "postgresql-18-pg-stat-ch" "green" >}} {{< bg "17" "postgresql-17-pg-stat-ch" "green" >}} {{< bg "16" "postgresql-16-pg-stat-ch" "green" >}} {{< bg "15" "postgresql-15-pg-stat-ch" "red" >}} {{< bg "14" "postgresql-14-pg-stat-ch" "red" >}} | `postgresql-$v-pg-stat-ch` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pg_stat_ch_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pg_stat_ch_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.4" "pg_stat_ch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "pg_stat_ch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "pg_stat_ch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_stat_ch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.4" "pg_stat_ch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "pg_stat_ch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "pg_stat_ch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_stat_ch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.4" "pg_stat_ch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "pg_stat_ch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "pg_stat_ch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_stat_ch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.4" "pg_stat_ch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "pg_stat_ch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "pg_stat_ch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_stat_ch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_ch_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-18-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-17-pg-stat-ch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.4" "postgresql-16-pg-stat-ch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-ch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-ch : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_ch_18` | `0.3.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_stat_ch_18-0.3.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_ch_18-0.3.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_stat_ch_18` | `0.3.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [pg_stat_ch_18-0.3.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_ch_18-0.3.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_stat_ch_18` | `0.3.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.0 MiB | [pg_stat_ch_18-0.3.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_ch_18-0.3.4-1PIGSTY.el10.x86_64.rpm) |
| `pg_stat_ch_18` | `0.3.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 994.5 KiB | [pg_stat_ch_18-0.3.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_ch_18-0.3.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-stat-ch` | `0.3.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 961.9 KiB | [postgresql-18-pg-stat-ch_0.3.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 865.7 KiB | [postgresql-18-pg-stat-ch_0.3.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 988.1 KiB | [postgresql-18-pg-stat-ch_0.3.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 884.3 KiB | [postgresql-18-pg-stat-ch_0.3.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.6 MiB | [postgresql-18-pg-stat-ch_0.3.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.3 MiB | [postgresql-18-pg-stat-ch_0.3.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 922.2 KiB | [postgresql-18-pg-stat-ch_0.3.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-stat-ch` | `0.3.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 880.8 KiB | [postgresql-18-pg-stat-ch_0.3.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-ch/postgresql-18-pg-stat-ch_0.3.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_ch_17` | `0.3.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_stat_ch_17-0.3.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_ch_17-0.3.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_stat_ch_17` | `0.3.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [pg_stat_ch_17-0.3.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_ch_17-0.3.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_stat_ch_17` | `0.3.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.0 MiB | [pg_stat_ch_17-0.3.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_ch_17-0.3.4-1PIGSTY.el10.x86_64.rpm) |
| `pg_stat_ch_17` | `0.3.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 992.7 KiB | [pg_stat_ch_17-0.3.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_ch_17-0.3.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-stat-ch` | `0.3.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 962.3 KiB | [postgresql-17-pg-stat-ch_0.3.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 865.7 KiB | [postgresql-17-pg-stat-ch_0.3.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 988.7 KiB | [postgresql-17-pg-stat-ch_0.3.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 884.5 KiB | [postgresql-17-pg-stat-ch_0.3.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.6 MiB | [postgresql-17-pg-stat-ch_0.3.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.3 MiB | [postgresql-17-pg-stat-ch_0.3.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 922.6 KiB | [postgresql-17-pg-stat-ch_0.3.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-stat-ch` | `0.3.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 879.4 KiB | [postgresql-17-pg-stat-ch_0.3.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-ch/postgresql-17-pg-stat-ch_0.3.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_ch_16` | `0.3.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_stat_ch_16-0.3.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_ch_16-0.3.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_stat_ch_16` | `0.3.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [pg_stat_ch_16-0.3.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_ch_16-0.3.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_stat_ch_16` | `0.3.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.0 MiB | [pg_stat_ch_16-0.3.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_ch_16-0.3.4-1PIGSTY.el10.x86_64.rpm) |
| `pg_stat_ch_16` | `0.3.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 994.7 KiB | [pg_stat_ch_16-0.3.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_ch_16-0.3.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-stat-ch` | `0.3.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 962.4 KiB | [postgresql-16-pg-stat-ch_0.3.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 866.3 KiB | [postgresql-16-pg-stat-ch_0.3.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 988.1 KiB | [postgresql-16-pg-stat-ch_0.3.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 884.6 KiB | [postgresql-16-pg-stat-ch_0.3.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.6 MiB | [postgresql-16-pg-stat-ch_0.3.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.3 MiB | [postgresql-16-pg-stat-ch_0.3.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 920.3 KiB | [postgresql-16-pg-stat-ch_0.3.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-stat-ch` | `0.3.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 877.2 KiB | [postgresql-16-pg-stat-ch_0.3.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-ch/postgresql-16-pg-stat-ch_0.3.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ClickHouse/pg_stat_ch" title="Repository" icon="github" subtitle="github.com/ClickHouse/pg_stat_ch" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_stat_ch-0.3.4.tar.gz" >}}
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

> Syntax:
>
> ```sql
> CREATE EXTENSION pg_stat_ch;
> SELECT pg_stat_ch_version();
> SELECT * FROM pg_stat_ch_stats();
> ```
>
> Sources: [README](https://github.com/ClickHouse/pg_stat_ch), [Blog post](https://clickhouse.com/blog/pg_stat_ch-postgres-extension-stats-to-clickhouse)

`pg_stat_ch` captures per-query execution telemetry in PostgreSQL and exports raw events to ClickHouse in real time. The upstream project contrasts this with `pg_stat_statements`: instead of aggregating inside PostgreSQL, it sends raw events to ClickHouse for downstream analysis.

## Architecture

The README describes a single pipeline:

```text
PostgreSQL hooks -> shared memory queue -> background worker -> ClickHouse
```

Design goals called out upstream include:

- no network I/O on the query path
- bounded memory via a fixed-size ring buffer
- raw event export instead of local aggregation
- graceful degradation when the queue overflows or ClickHouse is unavailable

## Setup

The extension must be preloaded and configured with ClickHouse connection settings:

```ini
shared_preload_libraries = 'pg_stat_ch'
track_io_timing = on

pg_stat_ch.clickhouse_host = 'localhost'
pg_stat_ch.clickhouse_port = 9000
pg_stat_ch.clickhouse_database = 'pg_stat_ch'
pg_stat_ch.clickhouse_use_tls = on
pg_stat_ch.clickhouse_skip_tls_verify = off
```

After PostgreSQL restart and ClickHouse schema setup:

```sql
CREATE EXTENSION pg_stat_ch;
```

## SQL API

The README documents these SQL functions:

- `pg_stat_ch_version()`
- `pg_stat_ch_stats()`
- `pg_stat_ch_reset()`

`pg_stat_ch_stats()` exposes queue and exporter counters so you can verify that events are being captured and flushed.

## What Gets Captured

The current README lists support for:

- query timing and row counts
- buffer usage and WAL usage
- CPU time
- DML, DDL, and utility statements
- SQLSTATE and error levels
- JIT instrumentation on PostgreSQL 15+
- parallel worker statistics on PostgreSQL 18+
- client context such as application name and client IP

The project currently states full support for PostgreSQL 16, 17, and 18.
