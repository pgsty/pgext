---
title: "pg_strict"
linkTitle: "pg_strict"
description: "Prevent dangerous UPDATE and DELETE without WHERE clause"
weight: 5830
categories: ["ADMIN"]
width: full
---

[**pg_strict**](https://github.com/spa5k/pg_strict) : Prevent dangerous UPDATE and DELETE without WHERE clause


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5830** | {{< badge content="pg_strict" link="https://github.com/spa5k/pg_strict" >}} | {{< ext "pg_strict" >}} | `1.0.5` | {{< category "ADMIN" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "safeupdate" >}} {{< ext "pg_savior" >}} {{< ext "pg_upless" >}} {{< ext "pg_drop_events" >}} {{< ext "pg_readonly" >}} {{< ext "table_log" >}} {{< ext "pgaudit" >}} {{< ext "pg_permissions" >}} |

> [!Note] manually upgraded PGRX from 0.16.1 to 0.17.0 by Vonng


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_strict` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.5` | {{< bg "18" "pg_strict_18" "green" >}} {{< bg "17" "pg_strict_17" "green" >}} {{< bg "16" "pg_strict_16" "green" >}} {{< bg "15" "pg_strict_15" "green" >}} {{< bg "14" "pg_strict_14" "green" >}} | `pg_strict_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.5` | {{< bg "18" "postgresql-18-pg-strict" "green" >}} {{< bg "17" "postgresql-17-pg-strict" "green" >}} {{< bg "16" "postgresql-16-pg-strict" "green" >}} {{< bg "15" "postgresql-15-pg-strict" "green" >}} {{< bg "14" "postgresql-14-pg-strict" "green" >}} | `postgresql-$v-pg-strict` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_strict_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-18-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-17-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-16-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-15-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-14-pg-strict : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-18-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-17-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-16-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-15-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-14-pg-strict : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-18-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-17-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-16-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-15-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-14-pg-strict : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-18-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-17-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-16-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-15-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-14-pg-strict : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-18-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-17-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-16-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-15-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-14-pg-strict : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-18-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-17-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-16-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-15-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-14-pg-strict : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-18-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-17-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-16-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-15-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-14-pg-strict : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-18-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-17-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-16-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-15-pg-strict : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-14-pg-strict : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-strict : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-strict : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-strict : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-strict : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-strict : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-strict : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-strict : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-strict : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-strict : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-strict : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_strict_18` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 328.1 KiB | [pg_strict_18-1.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_strict_18-1.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_strict_18` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 220.8 KiB | [pg_strict_18-1.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_strict_18-1.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_strict_18` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 344.2 KiB | [pg_strict_18-1.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_strict_18-1.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_strict_18` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 236.3 KiB | [pg_strict_18-1.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_strict_18-1.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_strict_18` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 344.4 KiB | [pg_strict_18-1.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_strict_18-1.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_strict_18` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 236.1 KiB | [pg_strict_18-1.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_strict_18-1.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-strict` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 269.8 KiB | [postgresql-18-pg-strict_1.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-strict/postgresql-18-pg-strict_1.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-strict` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 170.7 KiB | [postgresql-18-pg-strict_1.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-strict/postgresql-18-pg-strict_1.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-strict` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 270.2 KiB | [postgresql-18-pg-strict_1.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-strict/postgresql-18-pg-strict_1.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-strict` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 170.7 KiB | [postgresql-18-pg-strict_1.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-strict/postgresql-18-pg-strict_1.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-strict` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 306.2 KiB | [postgresql-18-pg-strict_1.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-strict/postgresql-18-pg-strict_1.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-strict` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 198.2 KiB | [postgresql-18-pg-strict_1.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-strict/postgresql-18-pg-strict_1.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-strict` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 303.1 KiB | [postgresql-18-pg-strict_1.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-strict/postgresql-18-pg-strict_1.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-strict` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 196.8 KiB | [postgresql-18-pg-strict_1.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-strict/postgresql-18-pg-strict_1.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_strict_17` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 328.1 KiB | [pg_strict_17-1.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_strict_17-1.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_strict_17` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 220.7 KiB | [pg_strict_17-1.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_strict_17-1.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_strict_17` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 344.4 KiB | [pg_strict_17-1.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_strict_17-1.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_strict_17` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 236.0 KiB | [pg_strict_17-1.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_strict_17-1.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_strict_17` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 344.4 KiB | [pg_strict_17-1.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_strict_17-1.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_strict_17` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 236.0 KiB | [pg_strict_17-1.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_strict_17-1.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-strict` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 269.8 KiB | [postgresql-17-pg-strict_1.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-strict/postgresql-17-pg-strict_1.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-strict` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 170.7 KiB | [postgresql-17-pg-strict_1.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-strict/postgresql-17-pg-strict_1.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-strict` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 269.6 KiB | [postgresql-17-pg-strict_1.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-strict/postgresql-17-pg-strict_1.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-strict` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 170.7 KiB | [postgresql-17-pg-strict_1.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-strict/postgresql-17-pg-strict_1.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-strict` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 306.3 KiB | [postgresql-17-pg-strict_1.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-strict/postgresql-17-pg-strict_1.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-strict` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 198.4 KiB | [postgresql-17-pg-strict_1.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-strict/postgresql-17-pg-strict_1.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-strict` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 303.2 KiB | [postgresql-17-pg-strict_1.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-strict/postgresql-17-pg-strict_1.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-strict` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 197.0 KiB | [postgresql-17-pg-strict_1.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-strict/postgresql-17-pg-strict_1.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_strict_16` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 328.1 KiB | [pg_strict_16-1.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_strict_16-1.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_strict_16` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 220.7 KiB | [pg_strict_16-1.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_strict_16-1.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_strict_16` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 344.1 KiB | [pg_strict_16-1.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_strict_16-1.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_strict_16` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 236.0 KiB | [pg_strict_16-1.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_strict_16-1.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_strict_16` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 344.4 KiB | [pg_strict_16-1.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_strict_16-1.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_strict_16` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 236.0 KiB | [pg_strict_16-1.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_strict_16-1.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-strict` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 269.8 KiB | [postgresql-16-pg-strict_1.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-strict/postgresql-16-pg-strict_1.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-strict` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 170.6 KiB | [postgresql-16-pg-strict_1.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-strict/postgresql-16-pg-strict_1.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-strict` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 269.8 KiB | [postgresql-16-pg-strict_1.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-strict/postgresql-16-pg-strict_1.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-strict` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 170.6 KiB | [postgresql-16-pg-strict_1.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-strict/postgresql-16-pg-strict_1.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-strict` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 306.2 KiB | [postgresql-16-pg-strict_1.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-strict/postgresql-16-pg-strict_1.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-strict` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 198.1 KiB | [postgresql-16-pg-strict_1.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-strict/postgresql-16-pg-strict_1.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-strict` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 303.0 KiB | [postgresql-16-pg-strict_1.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-strict/postgresql-16-pg-strict_1.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-strict` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 196.8 KiB | [postgresql-16-pg-strict_1.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-strict/postgresql-16-pg-strict_1.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_strict_15` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 327.9 KiB | [pg_strict_15-1.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_strict_15-1.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_strict_15` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 220.8 KiB | [pg_strict_15-1.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_strict_15-1.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_strict_15` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 344.0 KiB | [pg_strict_15-1.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_strict_15-1.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_strict_15` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 236.4 KiB | [pg_strict_15-1.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_strict_15-1.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_strict_15` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 344.3 KiB | [pg_strict_15-1.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_strict_15-1.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_strict_15` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 236.3 KiB | [pg_strict_15-1.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_strict_15-1.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-strict` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 269.9 KiB | [postgresql-15-pg-strict_1.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-strict/postgresql-15-pg-strict_1.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-strict` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 170.7 KiB | [postgresql-15-pg-strict_1.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-strict/postgresql-15-pg-strict_1.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-strict` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 270.5 KiB | [postgresql-15-pg-strict_1.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-strict/postgresql-15-pg-strict_1.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-strict` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 170.6 KiB | [postgresql-15-pg-strict_1.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-strict/postgresql-15-pg-strict_1.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-strict` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 306.0 KiB | [postgresql-15-pg-strict_1.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-strict/postgresql-15-pg-strict_1.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-strict` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 198.2 KiB | [postgresql-15-pg-strict_1.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-strict/postgresql-15-pg-strict_1.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-strict` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 303.4 KiB | [postgresql-15-pg-strict_1.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-strict/postgresql-15-pg-strict_1.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-strict` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 197.1 KiB | [postgresql-15-pg-strict_1.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-strict/postgresql-15-pg-strict_1.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_strict_14` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 327.9 KiB | [pg_strict_14-1.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_strict_14-1.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_strict_14` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 220.7 KiB | [pg_strict_14-1.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_strict_14-1.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_strict_14` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 344.0 KiB | [pg_strict_14-1.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_strict_14-1.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_strict_14` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 236.2 KiB | [pg_strict_14-1.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_strict_14-1.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_strict_14` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 344.6 KiB | [pg_strict_14-1.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_strict_14-1.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_strict_14` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 236.3 KiB | [pg_strict_14-1.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_strict_14-1.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-strict` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 270.4 KiB | [postgresql-14-pg-strict_1.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-strict/postgresql-14-pg-strict_1.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-strict` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 170.7 KiB | [postgresql-14-pg-strict_1.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-strict/postgresql-14-pg-strict_1.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-strict` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 270.3 KiB | [postgresql-14-pg-strict_1.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-strict/postgresql-14-pg-strict_1.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-strict` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 170.6 KiB | [postgresql-14-pg-strict_1.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-strict/postgresql-14-pg-strict_1.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-strict` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 306.1 KiB | [postgresql-14-pg-strict_1.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-strict/postgresql-14-pg-strict_1.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-strict` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 198.1 KiB | [postgresql-14-pg-strict_1.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-strict/postgresql-14-pg-strict_1.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-strict` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 303.0 KiB | [postgresql-14-pg-strict_1.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-strict/postgresql-14-pg-strict_1.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-strict` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 197.0 KiB | [postgresql-14-pg-strict_1.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-strict/postgresql-14-pg-strict_1.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/spa5k/pg_strict" title="Repository" icon="github" subtitle="github.com/spa5k/pg_strict" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_strict-1.0.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_strict;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_strict;		# install via package name, for the active PG version

pig install pg_strict -v 18;   # install for PG 18
pig install pg_strict -v 17;   # install for PG 17
pig install pg_strict -v 16;   # install for PG 16
pig install pg_strict -v 15;   # install for PG 15
pig install pg_strict -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_strict';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_strict;
```

## Usage

Source: [README](https://github.com/spa5k/pg_strict/blob/master/README.md), [Release v1.0.5](https://github.com/spa5k/pg_strict/releases/tag/v1.0.5), [API source](https://github.com/spa5k/pg_strict/blob/master/src/api.rs)

`pg_strict` blocks or warns on `UPDATE` and `DELETE` statements that omit a `WHERE` clause. It installs a `post_parse_analyze_hook`, so it must be loaded from `shared_preload_libraries`.

### Required setup

```sql
-- postgresql.conf
shared_preload_libraries = 'pg_strict'

CREATE EXTENSION pg_strict;
```

### GUCs

- `pg_strict.require_where_on_update`
- `pg_strict.require_where_on_delete`

Each setting supports `off`, `warn`, and `on`.

```sql
SET pg_strict.require_where_on_update = 'on';
SET pg_strict.require_where_on_delete = 'warn';
```

### Helper functions

```sql
SELECT pg_strict_version();
SELECT * FROM pg_strict_config();

SELECT pg_strict_check_where_clause('DELETE FROM t', 'DELETE');
SELECT pg_strict_validate_update('UPDATE t SET x = 1 WHERE id = 42');
SELECT pg_strict_validate_delete('DELETE FROM t WHERE id = 42');

SELECT pg_strict_enable_update();
SELECT pg_strict_warn_delete();
SELECT pg_strict_disable_delete();
```

- `pg_strict_set_update_mode(mode)` and `pg_strict_set_delete_mode(mode)` provide generic mode setters.
- `SET LOCAL` works for one-off bulk operations inside a transaction.

### Caveats

- Enforcement is presence-based, not intent-based: any non-null `WHERE` clause satisfies the rule.
- Only `UPDATE` and `DELETE` are checked.
- Current upstream release is `1.0.5`; the Pigsty note about `pgrx` 0.17.0 is packaging/build metadata, not a documented user-facing feature change.
