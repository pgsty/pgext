---
title: "pg_stat_backtrace"
linkTitle: "pg_stat_backtrace"
description: "Capture or log C-level stack backtraces of PostgreSQL processes"
weight: 6030
categories: ["STAT"]
width: full
---

[**pg_stat_backtrace**](https://github.com/Nickyoung0/pg_stat_backtrace) : Capture or log C-level stack backtraces of PostgreSQL processes


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6030** | {{< badge content="pg_stat_backtrace" link="https://github.com/Nickyoung0/pg_stat_backtrace" >}} | {{< ext "pg_stat_backtrace" >}} | `1.0.0` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |

> [!Note] GitHub v1.0.0; C PGXS extension using ptrace(PTRACE_SEIZE) and libunwind; Linux only; runtime may need kernel.yama.ptrace_scope=0


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_stat_backtrace` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pg_stat_backtrace_18" "green" >}} {{< bg "17" "pg_stat_backtrace_17" "green" >}} {{< bg "16" "pg_stat_backtrace_16" "green" >}} {{< bg "15" "pg_stat_backtrace_15" "green" >}} {{< bg "14" "pg_stat_backtrace_14" "green" >}} | `pg_stat_backtrace_$v` | `libunwind` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-pg-stat-backtrace" "green" >}} {{< bg "17" "postgresql-17-pg-stat-backtrace" "green" >}} {{< bg "16" "postgresql-16-pg-stat-backtrace" "green" >}} {{< bg "15" "postgresql-15-pg-stat-backtrace" "green" >}} {{< bg "14" "postgresql-14-pg-stat-backtrace" "green" >}} | `postgresql-$v-pg-stat-backtrace` | `libunwind8` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stat_backtrace_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-stat-backtrace : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-stat-backtrace : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-stat-backtrace : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-stat-backtrace : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-stat-backtrace : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-stat-backtrace : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-stat-backtrace : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-stat-backtrace : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-stat-backtrace : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-stat-backtrace : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-stat-backtrace : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_backtrace_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.9 KiB | [pg_stat_backtrace_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_stat_backtrace_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_stat_backtrace_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.0 KiB | [pg_stat_backtrace_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_stat_backtrace_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_stat_backtrace_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.0 KiB | [pg_stat_backtrace_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_backtrace_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_stat_backtrace_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.2 KiB | [pg_stat_backtrace_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_backtrace_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_stat_backtrace_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.2 KiB | [pg_stat_backtrace_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_backtrace_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_stat_backtrace_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 32.1 KiB | [pg_stat_backtrace_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_backtrace_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-stat-backtrace` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.6 KiB | [postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-backtrace/postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-stat-backtrace` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.4 KiB | [postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-backtrace/postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-stat-backtrace` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.6 KiB | [postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-backtrace/postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-stat-backtrace` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.5 KiB | [postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-backtrace/postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-stat-backtrace` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.0 KiB | [postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-backtrace/postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-stat-backtrace` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.7 KiB | [postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-backtrace/postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-stat-backtrace` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.1 KiB | [postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-backtrace/postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-stat-backtrace` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.7 KiB | [postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-backtrace/postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-stat-backtrace` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.1 KiB | [postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-backtrace/postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-stat-backtrace` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 17.7 KiB | [postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-backtrace/postgresql-18-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_backtrace_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.9 KiB | [pg_stat_backtrace_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_stat_backtrace_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_stat_backtrace_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.0 KiB | [pg_stat_backtrace_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_stat_backtrace_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_stat_backtrace_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.9 KiB | [pg_stat_backtrace_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_backtrace_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_stat_backtrace_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.3 KiB | [pg_stat_backtrace_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_backtrace_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_stat_backtrace_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.2 KiB | [pg_stat_backtrace_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_backtrace_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_stat_backtrace_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 32.2 KiB | [pg_stat_backtrace_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_backtrace_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-stat-backtrace` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.6 KiB | [postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-backtrace/postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-stat-backtrace` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.4 KiB | [postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-backtrace/postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-stat-backtrace` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.6 KiB | [postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-backtrace/postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-stat-backtrace` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.4 KiB | [postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-backtrace/postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-stat-backtrace` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.0 KiB | [postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-backtrace/postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-stat-backtrace` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.7 KiB | [postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-backtrace/postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-stat-backtrace` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.0 KiB | [postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-backtrace/postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-stat-backtrace` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.7 KiB | [postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-backtrace/postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-stat-backtrace` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.1 KiB | [postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-backtrace/postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-stat-backtrace` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 17.7 KiB | [postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-backtrace/postgresql-17-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_backtrace_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.9 KiB | [pg_stat_backtrace_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_stat_backtrace_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_stat_backtrace_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.0 KiB | [pg_stat_backtrace_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_stat_backtrace_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_stat_backtrace_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.9 KiB | [pg_stat_backtrace_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_backtrace_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_stat_backtrace_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.2 KiB | [pg_stat_backtrace_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_backtrace_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_stat_backtrace_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.2 KiB | [pg_stat_backtrace_16-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_backtrace_16-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_stat_backtrace_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 32.2 KiB | [pg_stat_backtrace_16-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_backtrace_16-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-stat-backtrace` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.6 KiB | [postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-backtrace/postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-stat-backtrace` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.4 KiB | [postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-backtrace/postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-stat-backtrace` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.6 KiB | [postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-backtrace/postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-stat-backtrace` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.5 KiB | [postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-backtrace/postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-stat-backtrace` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.0 KiB | [postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-backtrace/postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-stat-backtrace` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.7 KiB | [postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-backtrace/postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-stat-backtrace` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.0 KiB | [postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-backtrace/postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-stat-backtrace` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.6 KiB | [postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-backtrace/postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-stat-backtrace` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.1 KiB | [postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-backtrace/postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-stat-backtrace` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 17.7 KiB | [postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-backtrace/postgresql-16-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_backtrace_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.9 KiB | [pg_stat_backtrace_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_stat_backtrace_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_stat_backtrace_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.0 KiB | [pg_stat_backtrace_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_stat_backtrace_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_stat_backtrace_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.0 KiB | [pg_stat_backtrace_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_backtrace_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_stat_backtrace_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.2 KiB | [pg_stat_backtrace_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_backtrace_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_stat_backtrace_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.2 KiB | [pg_stat_backtrace_15-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_backtrace_15-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_stat_backtrace_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 32.2 KiB | [pg_stat_backtrace_15-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_backtrace_15-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-stat-backtrace` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.6 KiB | [postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-backtrace/postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-stat-backtrace` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.4 KiB | [postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-backtrace/postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-stat-backtrace` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.6 KiB | [postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-backtrace/postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-stat-backtrace` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.5 KiB | [postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-backtrace/postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-stat-backtrace` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.0 KiB | [postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-backtrace/postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-stat-backtrace` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.7 KiB | [postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-backtrace/postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-stat-backtrace` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.0 KiB | [postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-backtrace/postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-stat-backtrace` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.7 KiB | [postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-backtrace/postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-stat-backtrace` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.1 KiB | [postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-backtrace/postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-stat-backtrace` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 17.7 KiB | [postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-backtrace/postgresql-15-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_backtrace_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.9 KiB | [pg_stat_backtrace_14-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_stat_backtrace_14-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_stat_backtrace_14` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.0 KiB | [pg_stat_backtrace_14-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_stat_backtrace_14-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_stat_backtrace_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.9 KiB | [pg_stat_backtrace_14-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_backtrace_14-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_stat_backtrace_14` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.2 KiB | [pg_stat_backtrace_14-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_backtrace_14-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_stat_backtrace_14` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.2 KiB | [pg_stat_backtrace_14-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_backtrace_14-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_stat_backtrace_14` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 32.1 KiB | [pg_stat_backtrace_14-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_backtrace_14-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-stat-backtrace` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.6 KiB | [postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-backtrace/postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-stat-backtrace` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.4 KiB | [postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-backtrace/postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-stat-backtrace` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.6 KiB | [postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-backtrace/postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-stat-backtrace` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.4 KiB | [postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-backtrace/postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-stat-backtrace` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.0 KiB | [postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-backtrace/postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-stat-backtrace` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.7 KiB | [postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-backtrace/postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-stat-backtrace` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.0 KiB | [postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-backtrace/postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-stat-backtrace` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.6 KiB | [postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-backtrace/postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-stat-backtrace` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.1 KiB | [postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-backtrace/postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-stat-backtrace` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 17.7 KiB | [postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-backtrace/postgresql-14-pg-stat-backtrace_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Nickyoung0/pg_stat_backtrace" title="Repository" icon="github" subtitle="github.com/Nickyoung0/pg_stat_backtrace" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_stat_backtrace-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_stat_backtrace;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_stat_backtrace;		# install via package name, for the active PG version

pig install pg_stat_backtrace -v 18;   # install for PG 18
pig install pg_stat_backtrace -v 17;   # install for PG 17
pig install pg_stat_backtrace -v 16;   # install for PG 16
pig install pg_stat_backtrace -v 15;   # install for PG 15
pig install pg_stat_backtrace -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_stat_backtrace;
```

## Usage

> Sources: [pg_stat_backtrace upstream README](https://github.com/Nickyoung0/pg_stat_backtrace), [upstream changelog](https://github.com/Nickyoung0/pg_stat_backtrace/blob/main/CHANGELOG.md), [local metadata](../db/extension.csv), local source tarball `pg_stat_backtrace-1.0.0.tar.gz`.

`pg_stat_backtrace` captures or logs the C-level stack backtrace of a PostgreSQL backend or auxiliary process on the same Linux host. It uses `ptrace(PTRACE_SEIZE)` plus `libunwind`; it does not use `shared_preload_libraries` and does not send `SIGSTOP` to the target.

```sql
CREATE EXTENSION pg_stat_backtrace;
```

### Capture A Backtrace

Find a target process from PostgreSQL views, then call `pg_get_backtrace(pid)`:

```sql
SELECT pid, backend_type, state, wait_event, query
FROM pg_stat_activity
WHERE backend_type = 'autovacuum worker';

SELECT pg_get_backtrace(123456);
```

The returned text uses a `pstack(1)`-style format:

```text
#0  0x00007f5e6c1a4d9e in __epoll_wait+0x4e
#1  0x000055f1a8c2a213 in WaitEventSetWaitBlock+0x83
#2  0x000055f1a8c2a04e in WaitEventSetWait+0xfe
```

### Write To The Server Log

Use `pg_log_backtrace(pid)` when the result should go through the normal PostgreSQL log pipeline:

```sql
SELECT pg_log_backtrace(123456);

SELECT pid, pg_log_backtrace(pid)
FROM pg_stat_activity
WHERE backend_type = 'walsender';
```

The function returns `true` on success.

### Permissions

By default, execute privilege is revoked from `PUBLIC` for both functions. Grant access only to trusted monitoring roles:

```sql
GRANT EXECUTE ON FUNCTION pg_get_backtrace(int) TO observability;
GRANT EXECUTE ON FUNCTION pg_log_backtrace(int) TO observability;
```

The C code still enforces target checks:

- Superusers may target any PostgreSQL process in the instance, including auxiliary processes such as `walwriter`, `checkpointer`, `walsender`, autovacuum workers, startup, and archiver processes.
- Non-superusers may target only regular backends owned by roles they are members of.
- Auxiliary processes have no role ownership and are rejected for non-superusers.
- A non-superuser may not target a superuser-owned backend, even with role membership.

### Input And Error Behavior

Both functions are `VOLATILE STRICT PARALLEL RESTRICTED`.

```sql
SELECT pg_get_backtrace(NULL::int);  -- NULL, no ptrace attempt
SELECT pg_log_backtrace(NULL::int);  -- NULL, no ptrace attempt
SELECT pg_get_backtrace(0);          -- WARNING, then NULL
SELECT pg_log_backtrace(-1);         -- WARNING, then false
```

Self-targeting is rejected because a Linux process cannot ptrace itself:

```sql
SELECT pg_get_backtrace(pg_backend_pid());
```

### Operational Caveats

- Pigsty packages `pg_stat_backtrace` 1.0.0 for PostgreSQL 14-18. Upstream 1.0.0 also advertises PostgreSQL 19 compatibility.
- The extension is Linux-only and depends on `libunwind` / `libunwind-ptrace` at build and runtime.
- On hosts with Yama ptrace restrictions, backend-to-backend capture may require `kernel.yama.ptrace_scope = 0`.
- The target process is briefly paused while the stack is unwound. Avoid tight loops against critical processes such as `walwriter`, `checkpointer`, or synchronous-replication `walsender` on busy primaries.
- Linux permits only one tracer per target process. Concurrent calls against the same PID can fail with `EPERM`; retry after the in-flight call finishes.
