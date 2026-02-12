---
title: "pg_track_optimizer"
linkTitle: "pg_track_optimizer"
description: "Track planning decisions in comparison with
  execution reality"
weight: 6270
categories: ["STAT"]
width: full
---

[**pg_track_optimizer**](https://github.com/danolivo/pg_track_optimizer) : Track planning decisions in comparison with
  execution reality


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6270** | {{< badge content="pg_track_optimizer" link="https://github.com/danolivo/pg_track_optimizer" >}} | {{< ext "pg_track_optimizer" >}} | `0.9.1` | {{< category "STAT" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_track_settings" >}} {{< ext "pg_show_plans" >}} {{< ext "powa" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_store_plans" >}} {{< ext "auto_explain" >}} {{< ext "pg_stat_kcache" >}} {{< ext "pg_qualstats" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.9.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} {{< bg "13" "" "red" >}} | `pg_track_optimizer` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.9.1` | {{< bg "18" "pg_track_optimizer_18" "green" >}} {{< bg "17" "pg_track_optimizer_17" "green" >}} {{< bg "16" "pg_track_optimizer_16" "red" >}} {{< bg "15" "pg_track_optimizer_15" "red" >}} {{< bg "14" "pg_track_optimizer_14" "red" >}} {{< bg "13" "pg_track_optimizer_13" "red" >}} | `pg_track_optimizer_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.9.1` | {{< bg "18" "postgresql-18-pg-track-optimizer" "green" >}} {{< bg "17" "postgresql-17-pg-track-optimizer" "green" >}} {{< bg "16" "postgresql-16-pg-track-optimizer" "red" >}} {{< bg "15" "postgresql-15-pg-track-optimizer" "red" >}} {{< bg "14" "postgresql-14-pg-track-optimizer" "red" >}} {{< bg "13" "postgresql-13-pg-track-optimizer" "red" >}} | `postgresql-$v-pg-track-optimizer` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "pg_track_optimizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pg_track_optimizer_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_track_optimizer_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "pg_track_optimizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pg_track_optimizer_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_track_optimizer_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "pg_track_optimizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pg_track_optimizer_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_track_optimizer_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "pg_track_optimizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pg_track_optimizer_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_track_optimizer_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "pg_track_optimizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pg_track_optimizer_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_track_optimizer_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "pg_track_optimizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pg_track_optimizer_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_track_optimizer_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pg-track-optimizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pg-track-optimizer : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-track-optimizer : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pg-track-optimizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pg-track-optimizer : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-track-optimizer : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pg-track-optimizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pg-track-optimizer : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-track-optimizer : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pg-track-optimizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pg-track-optimizer : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-track-optimizer : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pg-track-optimizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pg-track-optimizer : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-track-optimizer : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pg-track-optimizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pg-track-optimizer : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-track-optimizer : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pg-track-optimizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pg-track-optimizer : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-track-optimizer : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pg-track-optimizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pg-track-optimizer : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-track-optimizer : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_track_optimizer_18` | `0.9.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.2 KiB | [pg_track_optimizer_18-0.9.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_track_optimizer_18-0.9.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_track_optimizer_18` | `0.9.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.4 KiB | [pg_track_optimizer_18-0.9.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_track_optimizer_18-0.9.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_track_optimizer_18` | `0.9.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.9 KiB | [pg_track_optimizer_18-0.9.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_track_optimizer_18-0.9.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_track_optimizer_18` | `0.9.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 34.6 KiB | [pg_track_optimizer_18-0.9.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_track_optimizer_18-0.9.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_track_optimizer_18` | `0.9.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.9 KiB | [pg_track_optimizer_18-0.9.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_track_optimizer_18-0.9.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_track_optimizer_18` | `0.9.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 35.0 KiB | [pg_track_optimizer_18-0.9.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_track_optimizer_18-0.9.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-track-optimizer` | `0.9.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 57.8 KiB | [postgresql-18-pg-track-optimizer_0.9.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-track-optimizer/postgresql-18-pg-track-optimizer_0.9.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-track-optimizer` | `0.9.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 56.9 KiB | [postgresql-18-pg-track-optimizer_0.9.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-track-optimizer/postgresql-18-pg-track-optimizer_0.9.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-track-optimizer` | `0.9.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 57.9 KiB | [postgresql-18-pg-track-optimizer_0.9.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-track-optimizer/postgresql-18-pg-track-optimizer_0.9.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-track-optimizer` | `0.9.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 57.1 KiB | [postgresql-18-pg-track-optimizer_0.9.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-track-optimizer/postgresql-18-pg-track-optimizer_0.9.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-track-optimizer` | `0.9.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 62.6 KiB | [postgresql-18-pg-track-optimizer_0.9.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-track-optimizer/postgresql-18-pg-track-optimizer_0.9.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-track-optimizer` | `0.9.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 62.0 KiB | [postgresql-18-pg-track-optimizer_0.9.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-track-optimizer/postgresql-18-pg-track-optimizer_0.9.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-track-optimizer` | `0.9.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 60.4 KiB | [postgresql-18-pg-track-optimizer_0.9.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-track-optimizer/postgresql-18-pg-track-optimizer_0.9.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-track-optimizer` | `0.9.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 59.8 KiB | [postgresql-18-pg-track-optimizer_0.9.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-track-optimizer/postgresql-18-pg-track-optimizer_0.9.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_track_optimizer_17` | `0.9.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.2 KiB | [pg_track_optimizer_17-0.9.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_track_optimizer_17-0.9.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_track_optimizer_17` | `0.9.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.4 KiB | [pg_track_optimizer_17-0.9.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_track_optimizer_17-0.9.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_track_optimizer_17` | `0.9.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.9 KiB | [pg_track_optimizer_17-0.9.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_track_optimizer_17-0.9.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_track_optimizer_17` | `0.9.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 34.6 KiB | [pg_track_optimizer_17-0.9.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_track_optimizer_17-0.9.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_track_optimizer_17` | `0.9.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.8 KiB | [pg_track_optimizer_17-0.9.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_track_optimizer_17-0.9.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_track_optimizer_17` | `0.9.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 35.0 KiB | [pg_track_optimizer_17-0.9.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_track_optimizer_17-0.9.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-track-optimizer` | `0.9.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 57.6 KiB | [postgresql-17-pg-track-optimizer_0.9.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-track-optimizer/postgresql-17-pg-track-optimizer_0.9.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-track-optimizer` | `0.9.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 56.8 KiB | [postgresql-17-pg-track-optimizer_0.9.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-track-optimizer/postgresql-17-pg-track-optimizer_0.9.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-track-optimizer` | `0.9.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 57.8 KiB | [postgresql-17-pg-track-optimizer_0.9.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-track-optimizer/postgresql-17-pg-track-optimizer_0.9.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-track-optimizer` | `0.9.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 56.9 KiB | [postgresql-17-pg-track-optimizer_0.9.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-track-optimizer/postgresql-17-pg-track-optimizer_0.9.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-track-optimizer` | `0.9.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 69.2 KiB | [postgresql-17-pg-track-optimizer_0.9.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-track-optimizer/postgresql-17-pg-track-optimizer_0.9.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-track-optimizer` | `0.9.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 68.4 KiB | [postgresql-17-pg-track-optimizer_0.9.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-track-optimizer/postgresql-17-pg-track-optimizer_0.9.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-track-optimizer` | `0.9.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 60.3 KiB | [postgresql-17-pg-track-optimizer_0.9.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-track-optimizer/postgresql-17-pg-track-optimizer_0.9.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-track-optimizer` | `0.9.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 59.6 KiB | [postgresql-17-pg-track-optimizer_0.9.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-track-optimizer/postgresql-17-pg-track-optimizer_0.9.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/danolivo/pg_track_optimizer" title="Repository" icon="github" subtitle="github.com/danolivo/pg_track_optimizer" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_track_optimizer-0.9.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_track_optimizer;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_track_optimizer;		# install via package name, for the active PG version

pig install pg_track_optimizer -v 18;   # install for PG 18
pig install pg_track_optimizer -v 17;   # install for PG 17

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_track_optimizer';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_track_optimizer;
```
