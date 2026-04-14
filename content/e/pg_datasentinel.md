---
title: "pg_datasentinel"
linkTitle: "pg_datasentinel"
description: "Observability and activity monitoring extension for PostgreSQL"
weight: 6400
categories: ["STAT"]
width: full
---

[**pg_datasentinel**](https://github.com/datasentinel/pg_datasentinel) : Observability and activity monitoring extension for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6400** | {{< badge content="pg_datasentinel" link="https://github.com/datasentinel/pg_datasentinel" >}} | {{< ext "pg_datasentinel" >}} | `1.0` | {{< category "STAT" >}} | {{< license "BSD-3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgsentinel" >}} {{< ext "system_stats" >}} {{< ext "pg_profile" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_stat_kcache" >}} {{< ext "powa" >}} |

> [!Note] shared_preload_libraries = pg_datasentinel is required because the extension allocates shared memory and hooks into activity logging.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `pg_datasentinel` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "pg_datasentinel_18" "green" >}} {{< bg "17" "pg_datasentinel_17" "green" >}} {{< bg "16" "pg_datasentinel_16" "green" >}} {{< bg "15" "pg_datasentinel_15" "green" >}} {{< bg "14" "pg_datasentinel_14" "red" >}} | `pg_datasentinel_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-pg-datasentinel" "green" >}} {{< bg "17" "postgresql-17-pg-datasentinel" "green" >}} {{< bg "16" "postgresql-16-pg-datasentinel" "green" >}} {{< bg "15" "postgresql-15-pg-datasentinel" "green" >}} {{< bg "14" "postgresql-14-pg-datasentinel" "red" >}} | `postgresql-$v-pg-datasentinel` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_datasentinel_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_datasentinel_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_datasentinel_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_datasentinel_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_datasentinel_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_datasentinel_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_datasentinel_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-datasentinel : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-datasentinel : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-datasentinel : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-datasentinel : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-datasentinel : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-datasentinel : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-datasentinel : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-datasentinel : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-datasentinel : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-datasentinel : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-datasentinel : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-datasentinel : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-datasentinel : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-datasentinel : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-datasentinel : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-datasentinel : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-datasentinel : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_datasentinel_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 34.6 KiB | [pg_datasentinel_18-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_datasentinel_18-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_datasentinel_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.6 KiB | [pg_datasentinel_18-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_datasentinel_18-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_datasentinel_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.3 KiB | [pg_datasentinel_18-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_datasentinel_18-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_datasentinel_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 34.8 KiB | [pg_datasentinel_18-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_datasentinel_18-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_datasentinel_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.7 KiB | [pg_datasentinel_18-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_datasentinel_18-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_datasentinel_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 35.2 KiB | [pg_datasentinel_18-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_datasentinel_18-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-datasentinel` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 58.8 KiB | [postgresql-18-pg-datasentinel_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-datasentinel/postgresql-18-pg-datasentinel_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-datasentinel` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 58.9 KiB | [postgresql-18-pg-datasentinel_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-datasentinel/postgresql-18-pg-datasentinel_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-datasentinel` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 59.0 KiB | [postgresql-18-pg-datasentinel_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-datasentinel/postgresql-18-pg-datasentinel_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-datasentinel` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 59.0 KiB | [postgresql-18-pg-datasentinel_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-datasentinel/postgresql-18-pg-datasentinel_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-datasentinel` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 63.2 KiB | [postgresql-18-pg-datasentinel_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-datasentinel/postgresql-18-pg-datasentinel_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-datasentinel` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 63.7 KiB | [postgresql-18-pg-datasentinel_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-datasentinel/postgresql-18-pg-datasentinel_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-datasentinel` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 61.4 KiB | [postgresql-18-pg-datasentinel_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-datasentinel/postgresql-18-pg-datasentinel_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-datasentinel` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 62.2 KiB | [postgresql-18-pg-datasentinel_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-datasentinel/postgresql-18-pg-datasentinel_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_datasentinel_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 34.6 KiB | [pg_datasentinel_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_datasentinel_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_datasentinel_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.6 KiB | [pg_datasentinel_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_datasentinel_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_datasentinel_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.3 KiB | [pg_datasentinel_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_datasentinel_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_datasentinel_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 34.8 KiB | [pg_datasentinel_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_datasentinel_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_datasentinel_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.7 KiB | [pg_datasentinel_17-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_datasentinel_17-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_datasentinel_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 35.3 KiB | [pg_datasentinel_17-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_datasentinel_17-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-datasentinel` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 58.7 KiB | [postgresql-17-pg-datasentinel_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-datasentinel/postgresql-17-pg-datasentinel_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-datasentinel` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 58.9 KiB | [postgresql-17-pg-datasentinel_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-datasentinel/postgresql-17-pg-datasentinel_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-datasentinel` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 59.0 KiB | [postgresql-17-pg-datasentinel_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-datasentinel/postgresql-17-pg-datasentinel_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-datasentinel` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 59.0 KiB | [postgresql-17-pg-datasentinel_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-datasentinel/postgresql-17-pg-datasentinel_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-datasentinel` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 71.9 KiB | [postgresql-17-pg-datasentinel_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-datasentinel/postgresql-17-pg-datasentinel_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-datasentinel` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 72.4 KiB | [postgresql-17-pg-datasentinel_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-datasentinel/postgresql-17-pg-datasentinel_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-datasentinel` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 61.4 KiB | [postgresql-17-pg-datasentinel_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-datasentinel/postgresql-17-pg-datasentinel_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-datasentinel` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 62.3 KiB | [postgresql-17-pg-datasentinel_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-datasentinel/postgresql-17-pg-datasentinel_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_datasentinel_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 34.6 KiB | [pg_datasentinel_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_datasentinel_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_datasentinel_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.6 KiB | [pg_datasentinel_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_datasentinel_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_datasentinel_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.3 KiB | [pg_datasentinel_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_datasentinel_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_datasentinel_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 34.8 KiB | [pg_datasentinel_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_datasentinel_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_datasentinel_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.7 KiB | [pg_datasentinel_16-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_datasentinel_16-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_datasentinel_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 35.3 KiB | [pg_datasentinel_16-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_datasentinel_16-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-datasentinel` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 58.7 KiB | [postgresql-16-pg-datasentinel_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-datasentinel/postgresql-16-pg-datasentinel_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-datasentinel` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 58.8 KiB | [postgresql-16-pg-datasentinel_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-datasentinel/postgresql-16-pg-datasentinel_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-datasentinel` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 58.9 KiB | [postgresql-16-pg-datasentinel_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-datasentinel/postgresql-16-pg-datasentinel_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-datasentinel` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 59.0 KiB | [postgresql-16-pg-datasentinel_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-datasentinel/postgresql-16-pg-datasentinel_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-datasentinel` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 71.4 KiB | [postgresql-16-pg-datasentinel_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-datasentinel/postgresql-16-pg-datasentinel_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-datasentinel` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 72.0 KiB | [postgresql-16-pg-datasentinel_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-datasentinel/postgresql-16-pg-datasentinel_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-datasentinel` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 61.4 KiB | [postgresql-16-pg-datasentinel_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-datasentinel/postgresql-16-pg-datasentinel_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-datasentinel` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 62.3 KiB | [postgresql-16-pg-datasentinel_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-datasentinel/postgresql-16-pg-datasentinel_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_datasentinel_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 34.8 KiB | [pg_datasentinel_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_datasentinel_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_datasentinel_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.8 KiB | [pg_datasentinel_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_datasentinel_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_datasentinel_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.6 KiB | [pg_datasentinel_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_datasentinel_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_datasentinel_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 35.0 KiB | [pg_datasentinel_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_datasentinel_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_datasentinel_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 35.1 KiB | [pg_datasentinel_15-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_datasentinel_15-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_datasentinel_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 35.6 KiB | [pg_datasentinel_15-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_datasentinel_15-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-datasentinel` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 59.0 KiB | [postgresql-15-pg-datasentinel_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-datasentinel/postgresql-15-pg-datasentinel_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-datasentinel` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 59.1 KiB | [postgresql-15-pg-datasentinel_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-datasentinel/postgresql-15-pg-datasentinel_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-datasentinel` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 59.2 KiB | [postgresql-15-pg-datasentinel_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-datasentinel/postgresql-15-pg-datasentinel_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-datasentinel` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 59.2 KiB | [postgresql-15-pg-datasentinel_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-datasentinel/postgresql-15-pg-datasentinel_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-datasentinel` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 71.7 KiB | [postgresql-15-pg-datasentinel_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-datasentinel/postgresql-15-pg-datasentinel_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-datasentinel` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 72.2 KiB | [postgresql-15-pg-datasentinel_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-datasentinel/postgresql-15-pg-datasentinel_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-datasentinel` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 61.6 KiB | [postgresql-15-pg-datasentinel_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-datasentinel/postgresql-15-pg-datasentinel_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-datasentinel` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 62.6 KiB | [postgresql-15-pg-datasentinel_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-datasentinel/postgresql-15-pg-datasentinel_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/datasentinel/pg_datasentinel" title="Repository" icon="github" subtitle="github.com/datasentinel/pg_datasentinel" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_datasentinel-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_datasentinel;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_datasentinel;		# install via package name, for the active PG version

pig install pg_datasentinel -v 18;   # install for PG 18
pig install pg_datasentinel -v 17;   # install for PG 17
pig install pg_datasentinel -v 16;   # install for PG 16
pig install pg_datasentinel -v 15;   # install for PG 15

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_datasentinel';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_datasentinel;
```
