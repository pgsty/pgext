---
title: "pg_tracing"
linkTitle: "pg_tracing"
description: "Distributed Tracing for PostgreSQL"
weight: 6010
categories: ["STAT"]
width: full
---

Distributed Tracing for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6010** | {{< badge content="pg_tracing" link="https://github.com/DataDog/pg_tracing" >}} | {{< ext "pg_tracing" >}} | `0.1.3` | {{< category "STAT" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_profile" >}} {{< ext "pg_show_plans" >}} {{< ext "pg_stat_kcache" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_store_plans" >}} {{< ext "pg_track_settings" >}} {{< ext "pg_wait_sampling" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_tracing" >}} | `0.1.3` | {{< bg "18" "pg_tracing_18*" "green" >}} {{< bg "17" "pg_tracing_17*" "green" >}} {{< bg "16" "pg_tracing_16*" "green" >}} {{< bg "15" "pg_tracing_15*" "green" >}} {{< bg "14" "pg_tracing_14*" "green" >}} {{< bg "13" "pg_tracing_13*" "red" >}} | `pg_tracing_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_tracing" >}} | `0.1.3` | {{< bg "18" "postgresql-18-pg-tracing" "green" >}} {{< bg "17" "postgresql-17-pg-tracing" "green" >}} {{< bg "16" "postgresql-16-pg-tracing" "green" >}} {{< bg "15" "postgresql-15-pg-tracing" "green" >}} {{< bg "14" "postgresql-14-pg-tracing" "green" >}} {{< bg "13" "postgresql-13-pg-tracing" "red" >}} | `postgresql-$v-pg-tracing` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_tracing_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_tracing_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_tracing_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_tracing_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_tracing_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_tracing_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-18-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-17-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-16-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-15-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-14-pg-tracing : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-tracing : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-18-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-17-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-16-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-15-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-14-pg-tracing : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-tracing : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-tracing : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-tracing : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-tracing : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-tracing : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-tracing : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-tracing : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-tracing : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-tracing : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-tracing : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-tracing : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-tracing : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-tracing : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-18-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-17-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-16-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-15-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-14-pg-tracing : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-tracing : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-18-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-17-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-16-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-15-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-14-pg-tracing : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-tracing : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-18-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-17-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-16-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-15-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-14-pg-tracing : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-tracing : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-18-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-17-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-16-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-15-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-14-pg-tracing : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-tracing : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tracing_18` | 0.1.3 | `el8.x86_64` | pigsty | 46.2 KiB | [pg_tracing_18-0.1.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tracing_18-0.1.3-2PIGSTY.el8.x86_64.rpm) |
| `pg_tracing_18` | 0.1.3 | `el8.aarch64` | pigsty | 44.8 KiB | [pg_tracing_18-0.1.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tracing_18-0.1.3-2PIGSTY.el8.aarch64.rpm) |
| `pg_tracing_18` | 0.1.3 | `el9.x86_64` | pigsty | 43.9 KiB | [pg_tracing_18-0.1.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tracing_18-0.1.3-2PIGSTY.el9.x86_64.rpm) |
| `pg_tracing_18` | 0.1.3 | `el9.aarch64` | pigsty | 43.3 KiB | [pg_tracing_18-0.1.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tracing_18-0.1.3-2PIGSTY.el9.aarch64.rpm) |
| `pg_tracing_18` | 0.1.3 | `el10.x86_64` | pigsty | 44.4 KiB | [pg_tracing_18-0.1.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tracing_18-0.1.3-2PIGSTY.el10.x86_64.rpm) |
| `pg_tracing_18` | 0.1.3 | `el10.aarch64` | pigsty | 43.7 KiB | [pg_tracing_18-0.1.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tracing_18-0.1.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-tracing` | 0.1.3 | `d12.x86_64` | pigsty | 105.3 KiB | [postgresql-18-pg-tracing_0.1.3-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tracing/postgresql-18-pg-tracing_0.1.3-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-tracing` | 0.1.3 | `d12.aarch64` | pigsty | 102.9 KiB | [postgresql-18-pg-tracing_0.1.3-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tracing/postgresql-18-pg-tracing_0.1.3-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-tracing` | 0.1.3 | `u22.x86_64` | pigsty | 114.7 KiB | [postgresql-18-pg-tracing_0.1.3-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tracing/postgresql-18-pg-tracing_0.1.3-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-tracing` | 0.1.3 | `u22.aarch64` | pigsty | 113.7 KiB | [postgresql-18-pg-tracing_0.1.3-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tracing/postgresql-18-pg-tracing_0.1.3-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-tracing` | 0.1.3 | `u24.x86_64` | pigsty | 110.1 KiB | [postgresql-18-pg-tracing_0.1.3-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tracing/postgresql-18-pg-tracing_0.1.3-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-tracing` | 0.1.3 | `u24.aarch64` | pigsty | 109.2 KiB | [postgresql-18-pg-tracing_0.1.3-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tracing/postgresql-18-pg-tracing_0.1.3-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tracing_17` | 0.1.3 | `el8.x86_64` | pigsty | 46.2 KiB | [pg_tracing_17-0.1.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tracing_17-0.1.3-2PIGSTY.el8.x86_64.rpm) |
| `pg_tracing_17` | 0.1.3 | `el8.aarch64` | pigsty | 44.9 KiB | [pg_tracing_17-0.1.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tracing_17-0.1.3-2PIGSTY.el8.aarch64.rpm) |
| `pg_tracing_17` | 0.1.3 | `el9.x86_64` | pigsty | 43.9 KiB | [pg_tracing_17-0.1.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tracing_17-0.1.3-2PIGSTY.el9.x86_64.rpm) |
| `pg_tracing_17` | 0.1.3 | `el9.aarch64` | pigsty | 43.4 KiB | [pg_tracing_17-0.1.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tracing_17-0.1.3-2PIGSTY.el9.aarch64.rpm) |
| `pg_tracing_17` | 0.1.3 | `el10.x86_64` | pigsty | 44.4 KiB | [pg_tracing_17-0.1.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tracing_17-0.1.3-2PIGSTY.el10.x86_64.rpm) |
| `pg_tracing_17` | 0.1.3 | `el10.aarch64` | pigsty | 43.6 KiB | [pg_tracing_17-0.1.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tracing_17-0.1.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-tracing` | 0.1.3 | `d12.x86_64` | pigsty | 105.4 KiB | [postgresql-17-pg-tracing_0.1.3-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tracing/postgresql-17-pg-tracing_0.1.3-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-tracing` | 0.1.3 | `d12.aarch64` | pigsty | 103.0 KiB | [postgresql-17-pg-tracing_0.1.3-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tracing/postgresql-17-pg-tracing_0.1.3-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-tracing` | 0.1.3 | `u22.x86_64` | pigsty | 129.3 KiB | [postgresql-17-pg-tracing_0.1.3-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tracing/postgresql-17-pg-tracing_0.1.3-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-tracing` | 0.1.3 | `u22.aarch64` | pigsty | 128.1 KiB | [postgresql-17-pg-tracing_0.1.3-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tracing/postgresql-17-pg-tracing_0.1.3-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-tracing` | 0.1.3 | `u24.x86_64` | pigsty | 110.2 KiB | [postgresql-17-pg-tracing_0.1.3-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tracing/postgresql-17-pg-tracing_0.1.3-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-tracing` | 0.1.3 | `u24.aarch64` | pigsty | 109.2 KiB | [postgresql-17-pg-tracing_0.1.3-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tracing/postgresql-17-pg-tracing_0.1.3-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tracing_16` | 0.1.3 | `el8.x86_64` | pigsty | 46.2 KiB | [pg_tracing_16-0.1.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tracing_16-0.1.3-2PIGSTY.el8.x86_64.rpm) |
| `pg_tracing_16` | 0.1.3 | `el8.aarch64` | pigsty | 44.9 KiB | [pg_tracing_16-0.1.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tracing_16-0.1.3-2PIGSTY.el8.aarch64.rpm) |
| `pg_tracing_16` | 0.1.3 | `el9.x86_64` | pigsty | 43.9 KiB | [pg_tracing_16-0.1.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tracing_16-0.1.3-2PIGSTY.el9.x86_64.rpm) |
| `pg_tracing_16` | 0.1.3 | `el9.aarch64` | pigsty | 43.3 KiB | [pg_tracing_16-0.1.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tracing_16-0.1.3-2PIGSTY.el9.aarch64.rpm) |
| `pg_tracing_16` | 0.1.3 | `el10.x86_64` | pigsty | 44.5 KiB | [pg_tracing_16-0.1.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tracing_16-0.1.3-2PIGSTY.el10.x86_64.rpm) |
| `pg_tracing_16` | 0.1.3 | `el10.aarch64` | pigsty | 43.5 KiB | [pg_tracing_16-0.1.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tracing_16-0.1.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-tracing` | 0.1.3 | `d12.x86_64` | pigsty | 105.4 KiB | [postgresql-16-pg-tracing_0.1.3-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tracing/postgresql-16-pg-tracing_0.1.3-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-tracing` | 0.1.3 | `d12.aarch64` | pigsty | 102.9 KiB | [postgresql-16-pg-tracing_0.1.3-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tracing/postgresql-16-pg-tracing_0.1.3-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-tracing` | 0.1.3 | `u22.x86_64` | pigsty | 128.8 KiB | [postgresql-16-pg-tracing_0.1.3-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tracing/postgresql-16-pg-tracing_0.1.3-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-tracing` | 0.1.3 | `u22.aarch64` | pigsty | 127.6 KiB | [postgresql-16-pg-tracing_0.1.3-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tracing/postgresql-16-pg-tracing_0.1.3-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-tracing` | 0.1.3 | `u24.x86_64` | pigsty | 110.1 KiB | [postgresql-16-pg-tracing_0.1.3-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tracing/postgresql-16-pg-tracing_0.1.3-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-tracing` | 0.1.3 | `u24.aarch64` | pigsty | 109.1 KiB | [postgresql-16-pg-tracing_0.1.3-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tracing/postgresql-16-pg-tracing_0.1.3-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tracing_15` | 0.1.3 | `el8.x86_64` | pigsty | 47.9 KiB | [pg_tracing_15-0.1.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tracing_15-0.1.3-2PIGSTY.el8.x86_64.rpm) |
| `pg_tracing_15` | 0.1.3 | `el8.aarch64` | pigsty | 46.5 KiB | [pg_tracing_15-0.1.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tracing_15-0.1.3-2PIGSTY.el8.aarch64.rpm) |
| `pg_tracing_15` | 0.1.3 | `el9.x86_64` | pigsty | 46.4 KiB | [pg_tracing_15-0.1.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tracing_15-0.1.3-2PIGSTY.el9.x86_64.rpm) |
| `pg_tracing_15` | 0.1.3 | `el9.aarch64` | pigsty | 45.8 KiB | [pg_tracing_15-0.1.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tracing_15-0.1.3-2PIGSTY.el9.aarch64.rpm) |
| `pg_tracing_15` | 0.1.3 | `el10.x86_64` | pigsty | 46.8 KiB | [pg_tracing_15-0.1.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tracing_15-0.1.3-2PIGSTY.el10.x86_64.rpm) |
| `pg_tracing_15` | 0.1.3 | `el10.aarch64` | pigsty | 46.2 KiB | [pg_tracing_15-0.1.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tracing_15-0.1.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-tracing` | 0.1.3 | `d12.x86_64` | pigsty | 108.2 KiB | [postgresql-15-pg-tracing_0.1.3-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tracing/postgresql-15-pg-tracing_0.1.3-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-tracing` | 0.1.3 | `d12.aarch64` | pigsty | 105.3 KiB | [postgresql-15-pg-tracing_0.1.3-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tracing/postgresql-15-pg-tracing_0.1.3-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-tracing` | 0.1.3 | `u22.x86_64` | pigsty | 132.0 KiB | [postgresql-15-pg-tracing_0.1.3-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tracing/postgresql-15-pg-tracing_0.1.3-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-tracing` | 0.1.3 | `u22.aarch64` | pigsty | 130.8 KiB | [postgresql-15-pg-tracing_0.1.3-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tracing/postgresql-15-pg-tracing_0.1.3-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-tracing` | 0.1.3 | `u24.x86_64` | pigsty | 113.1 KiB | [postgresql-15-pg-tracing_0.1.3-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tracing/postgresql-15-pg-tracing_0.1.3-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-tracing` | 0.1.3 | `u24.aarch64` | pigsty | 112.4 KiB | [postgresql-15-pg-tracing_0.1.3-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tracing/postgresql-15-pg-tracing_0.1.3-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tracing_14` | 0.1.3 | `el8.x86_64` | pigsty | 48.2 KiB | [pg_tracing_14-0.1.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tracing_14-0.1.3-2PIGSTY.el8.x86_64.rpm) |
| `pg_tracing_14` | 0.1.3 | `el8.aarch64` | pigsty | 47.3 KiB | [pg_tracing_14-0.1.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tracing_14-0.1.3-2PIGSTY.el8.aarch64.rpm) |
| `pg_tracing_14` | 0.1.3 | `el9.x86_64` | pigsty | 46.6 KiB | [pg_tracing_14-0.1.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tracing_14-0.1.3-2PIGSTY.el9.x86_64.rpm) |
| `pg_tracing_14` | 0.1.3 | `el9.aarch64` | pigsty | 46.3 KiB | [pg_tracing_14-0.1.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tracing_14-0.1.3-2PIGSTY.el9.aarch64.rpm) |
| `pg_tracing_14` | 0.1.3 | `el10.x86_64` | pigsty | 47.3 KiB | [pg_tracing_14-0.1.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tracing_14-0.1.3-2PIGSTY.el10.x86_64.rpm) |
| `pg_tracing_14` | 0.1.3 | `el10.aarch64` | pigsty | 46.8 KiB | [pg_tracing_14-0.1.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tracing_14-0.1.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-tracing` | 0.1.3 | `d12.x86_64` | pigsty | 109.6 KiB | [postgresql-14-pg-tracing_0.1.3-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tracing/postgresql-14-pg-tracing_0.1.3-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-tracing` | 0.1.3 | `d12.aarch64` | pigsty | 107.0 KiB | [postgresql-14-pg-tracing_0.1.3-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tracing/postgresql-14-pg-tracing_0.1.3-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-tracing` | 0.1.3 | `u22.x86_64` | pigsty | 133.2 KiB | [postgresql-14-pg-tracing_0.1.3-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tracing/postgresql-14-pg-tracing_0.1.3-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-tracing` | 0.1.3 | `u22.aarch64` | pigsty | 132.7 KiB | [postgresql-14-pg-tracing_0.1.3-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tracing/postgresql-14-pg-tracing_0.1.3-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-tracing` | 0.1.3 | `u24.x86_64` | pigsty | 114.1 KiB | [postgresql-14-pg-tracing_0.1.3-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tracing/postgresql-14-pg-tracing_0.1.3-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-tracing` | 0.1.3 | `u24.aarch64` | pigsty | 113.9 KiB | [postgresql-14-pg-tracing_0.1.3-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tracing/postgresql-14-pg-tracing_0.1.3-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/DataDog/pg_tracing" title="Repository" icon="github" subtitle="github.com/DataDog/pg_tracing" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_tracing-0.1.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_tracing; # get pg_tracing source code
pig build dep pg_tracing; # install build dependencies
pig build pkg pg_tracing; # build extension rpm or deb
pig build ext pg_tracing; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_tracing; # install by extension name, for the current active PG version
pig ext install pg_tracing; # install via package alias, for the active PG version
pig ext install pg_tracing -v 18;   # install for PG 18
pig ext install pg_tracing -v 17;   # install for PG 17
pig ext install pg_tracing -v 16;   # install for PG 16
pig ext install pg_tracing -v 15;   # install for PG 15
pig ext install pg_tracing -v 14;   # install for PG 14

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_tracing;
```

