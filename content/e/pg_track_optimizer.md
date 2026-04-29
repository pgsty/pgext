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
| **6270** | {{< badge content="pg_track_optimizer" link="https://github.com/danolivo/pg_track_optimizer" >}} | {{< ext "pg_track_optimizer" >}} | `0.9.2` | {{< category "STAT" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_track_settings" >}} {{< ext "pg_show_plans" >}} {{< ext "powa" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_store_plans" >}} {{< ext "auto_explain" >}} {{< ext "pg_stat_kcache" >}} {{< ext "pg_qualstats" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.9.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_track_optimizer` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.9.2` | {{< bg "18" "pg_track_optimizer_18" "green" >}} {{< bg "17" "pg_track_optimizer_17" "green" >}} {{< bg "16" "pg_track_optimizer_16" "red" >}} {{< bg "15" "pg_track_optimizer_15" "red" >}} {{< bg "14" "pg_track_optimizer_14" "red" >}} | `pg_track_optimizer_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.9.2` | {{< bg "18" "postgresql-18-pg-track-optimizer" "green" >}} {{< bg "17" "postgresql-17-pg-track-optimizer" "green" >}} {{< bg "16" "postgresql-16-pg-track-optimizer" "red" >}} {{< bg "15" "postgresql-15-pg-track-optimizer" "red" >}} {{< bg "14" "postgresql-14-pg-track-optimizer" "red" >}} | `postgresql-$v-pg-track-optimizer` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.9.2" "pg_track_optimizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_track_optimizer_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_track_optimizer_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.9.2" "pg_track_optimizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_track_optimizer_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_track_optimizer_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.9.2" "pg_track_optimizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_track_optimizer_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_track_optimizer_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.9.2" "pg_track_optimizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_track_optimizer_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_track_optimizer_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.9.2" "pg_track_optimizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_track_optimizer_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_track_optimizer_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.9.2" "pg_track_optimizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_track_optimizer_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_track_optimizer_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_track_optimizer_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-18-pg-track-optimizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-17-pg-track-optimizer : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-track-optimizer : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-18-pg-track-optimizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-17-pg-track-optimizer : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-track-optimizer : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-18-pg-track-optimizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-17-pg-track-optimizer : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-track-optimizer : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-18-pg-track-optimizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-17-pg-track-optimizer : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-track-optimizer : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-18-pg-track-optimizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-17-pg-track-optimizer : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-track-optimizer : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-18-pg-track-optimizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-17-pg-track-optimizer : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-track-optimizer : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-18-pg-track-optimizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-17-pg-track-optimizer : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-track-optimizer : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-18-pg-track-optimizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-17-pg-track-optimizer : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-track-optimizer : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-track-optimizer : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-track-optimizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-track-optimizer : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_track_optimizer_18` | `0.9.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.6 KiB | [pg_track_optimizer_18-0.9.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_track_optimizer_18-0.9.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_track_optimizer_18` | `0.9.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.8 KiB | [pg_track_optimizer_18-0.9.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_track_optimizer_18-0.9.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_track_optimizer_18` | `0.9.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 35.3 KiB | [pg_track_optimizer_18-0.9.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_track_optimizer_18-0.9.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_track_optimizer_18` | `0.9.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 35.1 KiB | [pg_track_optimizer_18-0.9.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_track_optimizer_18-0.9.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_track_optimizer_18` | `0.9.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 35.3 KiB | [pg_track_optimizer_18-0.9.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_track_optimizer_18-0.9.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_track_optimizer_18` | `0.9.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 35.4 KiB | [pg_track_optimizer_18-0.9.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_track_optimizer_18-0.9.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-track-optimizer` | `0.9.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 58.2 KiB | [postgresql-18-pg-track-optimizer_0.9.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-track-optimizer/postgresql-18-pg-track-optimizer_0.9.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-track-optimizer` | `0.9.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 57.5 KiB | [postgresql-18-pg-track-optimizer_0.9.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-track-optimizer/postgresql-18-pg-track-optimizer_0.9.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-track-optimizer` | `0.9.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 58.5 KiB | [postgresql-18-pg-track-optimizer_0.9.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-track-optimizer/postgresql-18-pg-track-optimizer_0.9.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-track-optimizer` | `0.9.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 57.6 KiB | [postgresql-18-pg-track-optimizer_0.9.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-track-optimizer/postgresql-18-pg-track-optimizer_0.9.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-track-optimizer` | `0.9.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 63.2 KiB | [postgresql-18-pg-track-optimizer_0.9.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-track-optimizer/postgresql-18-pg-track-optimizer_0.9.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-track-optimizer` | `0.9.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 62.5 KiB | [postgresql-18-pg-track-optimizer_0.9.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-track-optimizer/postgresql-18-pg-track-optimizer_0.9.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-track-optimizer` | `0.9.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 60.9 KiB | [postgresql-18-pg-track-optimizer_0.9.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-track-optimizer/postgresql-18-pg-track-optimizer_0.9.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-track-optimizer` | `0.9.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 60.3 KiB | [postgresql-18-pg-track-optimizer_0.9.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-track-optimizer/postgresql-18-pg-track-optimizer_0.9.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_track_optimizer_17` | `0.9.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.6 KiB | [pg_track_optimizer_17-0.9.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_track_optimizer_17-0.9.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_track_optimizer_17` | `0.9.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.8 KiB | [pg_track_optimizer_17-0.9.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_track_optimizer_17-0.9.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_track_optimizer_17` | `0.9.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 35.3 KiB | [pg_track_optimizer_17-0.9.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_track_optimizer_17-0.9.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_track_optimizer_17` | `0.9.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 35.0 KiB | [pg_track_optimizer_17-0.9.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_track_optimizer_17-0.9.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_track_optimizer_17` | `0.9.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 35.2 KiB | [pg_track_optimizer_17-0.9.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_track_optimizer_17-0.9.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_track_optimizer_17` | `0.9.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 35.4 KiB | [pg_track_optimizer_17-0.9.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_track_optimizer_17-0.9.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-track-optimizer` | `0.9.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 58.1 KiB | [postgresql-17-pg-track-optimizer_0.9.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-track-optimizer/postgresql-17-pg-track-optimizer_0.9.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-track-optimizer` | `0.9.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 57.3 KiB | [postgresql-17-pg-track-optimizer_0.9.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-track-optimizer/postgresql-17-pg-track-optimizer_0.9.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-track-optimizer` | `0.9.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 58.3 KiB | [postgresql-17-pg-track-optimizer_0.9.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-track-optimizer/postgresql-17-pg-track-optimizer_0.9.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-track-optimizer` | `0.9.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 57.4 KiB | [postgresql-17-pg-track-optimizer_0.9.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-track-optimizer/postgresql-17-pg-track-optimizer_0.9.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-track-optimizer` | `0.9.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 69.8 KiB | [postgresql-17-pg-track-optimizer_0.9.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-track-optimizer/postgresql-17-pg-track-optimizer_0.9.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-track-optimizer` | `0.9.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 69.0 KiB | [postgresql-17-pg-track-optimizer_0.9.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-track-optimizer/postgresql-17-pg-track-optimizer_0.9.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-track-optimizer` | `0.9.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 60.8 KiB | [postgresql-17-pg-track-optimizer_0.9.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-track-optimizer/postgresql-17-pg-track-optimizer_0.9.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-track-optimizer` | `0.9.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 60.1 KiB | [postgresql-17-pg-track-optimizer_0.9.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-track-optimizer/postgresql-17-pg-track-optimizer_0.9.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/danolivo/pg_track_optimizer" title="Repository" icon="github" subtitle="github.com/danolivo/pg_track_optimizer" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_track_optimizer-0.9.2.tar.gz" >}}
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



## Usage

> [pg_track_optimizer: detect suboptimal query plans via cardinality estimation errors](https://github.com/danolivo/pg_track_optimizer)

pg_track_optimizer automatically detects queries with poor cardinality estimates by comparing planner predictions to actual execution results. It calculates multiple error metrics using logarithmic scale.

### Enable Tracking

```sql
-- Track only problematic queries in production
SET pg_track_optimizer.mode = 'normal';

-- Track all queries during debugging
SET pg_track_optimizer.mode = 'forced';

-- Log EXPLAIN for queries exceeding error threshold
SET pg_track_optimizer.log_min_error = 2.0;
```

### Viewing Tracked Queries

```sql
SELECT queryid, query,
       avg_avg, avg_min, avg_max,
       rms_avg, rms_min, rms_max,
       time_avg, blks_avg, nexecs
FROM pg_track_optimizer
ORDER BY avg_avg DESC
LIMIT 10;

-- Using the RStats type directly
SELECT queryid, query,
       wca_error -> 'mean' AS avg_wca_error,
       blks_accessed -> 'mean' AS avg_blocks
FROM pg_track_optimizer()
WHERE blks_accessed -> 'mean' > 1000
ORDER BY wca_error -> 'mean' DESC;
```

### Error Metrics

| Metric | Description |
|--------|-------------|
| `avg_error` | Simple average of log-scale errors across plan nodes |
| `rms_error` | Root Mean Square, emphasizes large errors |
| `twa_error` | Time-Weighted Average, highlights slow nodes |
| `wca_error` | Cost-Weighted Average, highlights high-cost nodes |
| `f_join_filter` | JOIN filtering overhead factor |
| `f_scan_filter` | Scan filtering overhead factor |

### Managing Statistics

```sql
-- Save statistics to disk
SELECT pg_track_optimizer_flush();

-- Clear all tracked statistics
SELECT pg_track_optimizer_reset();

-- Check extension status
SELECT * FROM pg_track_optimizer_status;
```

### Configuration

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pg_track_optimizer.mode` | `disabled` | `disabled`, `normal`, `forced` |
| `pg_track_optimizer.log_min_error` | (none) | Error threshold for logging EXPLAIN |
| `pg_track_optimizer.hash_mem` | (default) | Shared memory limit in KB |
| `pg_track_optimizer.auto_flush` | `on` | Auto-save stats on backend shutdown |
