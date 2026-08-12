---
title: "pg_disorder"
linkTitle: "pg_disorder"
description: "Perturb unordered SELECT row order to expose order-dependent tests"
weight: 2880
categories: ["FEAT"]
width: full
---

[**pg_disorder**](https://github.com/viralpraxis/pg_disorder) : Perturb unordered SELECT row order to expose order-dependent tests


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2880** | {{< badge content="pg_disorder" link="https://github.com/viralpraxis/pg_disorder" >}} | {{< ext "pg_disorder" >}} | `0.1.0` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plan_filter" >}} {{< ext "pg_hint_plan" >}} {{< ext "pg_mockable" >}} {{< ext "pgtap" >}} {{< ext "pg_simula" >}} {{< ext "pg_fiu" >}} {{< ext "pg_crash" >}} |

> [!Note] Headless loadable module with no control file and no CREATE EXTENSION step; intended only for test databases; load per session with session_preload_libraries and never enable globally in production.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_disorder` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pg_disorder_18" "green" >}} {{< bg "17" "pg_disorder_17" "green" >}} {{< bg "16" "pg_disorder_16" "green" >}} {{< bg "15" "pg_disorder_15" "green" >}} {{< bg "14" "pg_disorder_14" "green" >}} | `pg_disorder_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-pg-disorder" "green" >}} {{< bg "17" "postgresql-17-pg-disorder" "green" >}} {{< bg "16" "postgresql-16-pg-disorder" "green" >}} {{< bg "15" "postgresql-15-pg-disorder" "green" >}} {{< bg "14" "postgresql-14-pg-disorder" "green" >}} | `postgresql-$v-pg-disorder` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_disorder_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-disorder : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-disorder : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-disorder : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-disorder : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-disorder : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-disorder : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-disorder : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-disorder : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-disorder : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-disorder : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-disorder : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_disorder_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.7 KiB | [pg_disorder_18-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_disorder_18-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_disorder_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.8 KiB | [pg_disorder_18-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_disorder_18-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_disorder_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.7 KiB | [pg_disorder_18-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_disorder_18-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_disorder_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.7 KiB | [pg_disorder_18-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_disorder_18-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_disorder_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.7 KiB | [pg_disorder_18-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_disorder_18-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_disorder_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.9 KiB | [pg_disorder_18-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_disorder_18-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-disorder` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.3 KiB | [postgresql-18-pg-disorder_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-disorder/postgresql-18-pg-disorder_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-disorder` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 21.0 KiB | [postgresql-18-pg-disorder_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-disorder/postgresql-18-pg-disorder_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-disorder` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.3 KiB | [postgresql-18-pg-disorder_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-disorder/postgresql-18-pg-disorder_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-disorder` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.1 KiB | [postgresql-18-pg-disorder_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-disorder/postgresql-18-pg-disorder_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-disorder` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.2 KiB | [postgresql-18-pg-disorder_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-disorder/postgresql-18-pg-disorder_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-disorder` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.0 KiB | [postgresql-18-pg-disorder_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-disorder/postgresql-18-pg-disorder_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-disorder` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.2 KiB | [postgresql-18-pg-disorder_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-disorder/postgresql-18-pg-disorder_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-disorder` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.9 KiB | [postgresql-18-pg-disorder_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-disorder/postgresql-18-pg-disorder_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-disorder` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 22.1 KiB | [postgresql-18-pg-disorder_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-disorder/postgresql-18-pg-disorder_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-disorder` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 21.8 KiB | [postgresql-18-pg-disorder_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-disorder/postgresql-18-pg-disorder_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_disorder_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.7 KiB | [pg_disorder_17-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_disorder_17-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_disorder_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.8 KiB | [pg_disorder_17-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_disorder_17-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_disorder_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.7 KiB | [pg_disorder_17-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_disorder_17-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_disorder_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.7 KiB | [pg_disorder_17-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_disorder_17-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_disorder_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.7 KiB | [pg_disorder_17-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_disorder_17-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_disorder_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.9 KiB | [pg_disorder_17-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_disorder_17-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-disorder` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.2 KiB | [postgresql-17-pg-disorder_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-disorder/postgresql-17-pg-disorder_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-disorder` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 21.0 KiB | [postgresql-17-pg-disorder_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-disorder/postgresql-17-pg-disorder_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-disorder` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.2 KiB | [postgresql-17-pg-disorder_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-disorder/postgresql-17-pg-disorder_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-disorder` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.1 KiB | [postgresql-17-pg-disorder_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-disorder/postgresql-17-pg-disorder_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-disorder` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.0 KiB | [postgresql-17-pg-disorder_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-disorder/postgresql-17-pg-disorder_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-disorder` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.7 KiB | [postgresql-17-pg-disorder_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-disorder/postgresql-17-pg-disorder_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-disorder` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.0 KiB | [postgresql-17-pg-disorder_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-disorder/postgresql-17-pg-disorder_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-disorder` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.9 KiB | [postgresql-17-pg-disorder_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-disorder/postgresql-17-pg-disorder_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-disorder` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 22.2 KiB | [postgresql-17-pg-disorder_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-disorder/postgresql-17-pg-disorder_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-disorder` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 21.8 KiB | [postgresql-17-pg-disorder_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-disorder/postgresql-17-pg-disorder_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_disorder_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.8 KiB | [pg_disorder_16-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_disorder_16-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_disorder_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.8 KiB | [pg_disorder_16-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_disorder_16-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_disorder_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.8 KiB | [pg_disorder_16-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_disorder_16-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_disorder_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.9 KiB | [pg_disorder_16-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_disorder_16-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_disorder_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.8 KiB | [pg_disorder_16-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_disorder_16-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_disorder_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.0 KiB | [pg_disorder_16-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_disorder_16-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-disorder` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.3 KiB | [postgresql-16-pg-disorder_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-disorder/postgresql-16-pg-disorder_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-disorder` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 21.2 KiB | [postgresql-16-pg-disorder_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-disorder/postgresql-16-pg-disorder_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-disorder` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.3 KiB | [postgresql-16-pg-disorder_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-disorder/postgresql-16-pg-disorder_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-disorder` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.2 KiB | [postgresql-16-pg-disorder_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-disorder/postgresql-16-pg-disorder_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-disorder` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.1 KiB | [postgresql-16-pg-disorder_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-disorder/postgresql-16-pg-disorder_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-disorder` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.8 KiB | [postgresql-16-pg-disorder_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-disorder/postgresql-16-pg-disorder_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-disorder` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.2 KiB | [postgresql-16-pg-disorder_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-disorder/postgresql-16-pg-disorder_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-disorder` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 22.1 KiB | [postgresql-16-pg-disorder_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-disorder/postgresql-16-pg-disorder_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-disorder` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 22.3 KiB | [postgresql-16-pg-disorder_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-disorder/postgresql-16-pg-disorder_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-disorder` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 21.9 KiB | [postgresql-16-pg-disorder_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-disorder/postgresql-16-pg-disorder_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_disorder_15` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.7 KiB | [pg_disorder_15-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_disorder_15-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_disorder_15` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.8 KiB | [pg_disorder_15-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_disorder_15-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_disorder_15` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.8 KiB | [pg_disorder_15-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_disorder_15-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_disorder_15` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.8 KiB | [pg_disorder_15-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_disorder_15-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_disorder_15` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.8 KiB | [pg_disorder_15-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_disorder_15-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_disorder_15` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.0 KiB | [pg_disorder_15-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_disorder_15-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-disorder` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.3 KiB | [postgresql-15-pg-disorder_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-disorder/postgresql-15-pg-disorder_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-disorder` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 21.2 KiB | [postgresql-15-pg-disorder_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-disorder/postgresql-15-pg-disorder_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-disorder` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.3 KiB | [postgresql-15-pg-disorder_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-disorder/postgresql-15-pg-disorder_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-disorder` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.2 KiB | [postgresql-15-pg-disorder_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-disorder/postgresql-15-pg-disorder_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-disorder` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.1 KiB | [postgresql-15-pg-disorder_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-disorder/postgresql-15-pg-disorder_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-disorder` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.8 KiB | [postgresql-15-pg-disorder_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-disorder/postgresql-15-pg-disorder_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-disorder` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.2 KiB | [postgresql-15-pg-disorder_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-disorder/postgresql-15-pg-disorder_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-disorder` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 22.1 KiB | [postgresql-15-pg-disorder_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-disorder/postgresql-15-pg-disorder_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-disorder` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 22.3 KiB | [postgresql-15-pg-disorder_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-disorder/postgresql-15-pg-disorder_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-disorder` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 22.0 KiB | [postgresql-15-pg-disorder_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-disorder/postgresql-15-pg-disorder_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_disorder_14` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.7 KiB | [pg_disorder_14-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_disorder_14-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_disorder_14` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.8 KiB | [pg_disorder_14-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_disorder_14-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_disorder_14` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.8 KiB | [pg_disorder_14-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_disorder_14-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_disorder_14` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.8 KiB | [pg_disorder_14-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_disorder_14-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_disorder_14` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.8 KiB | [pg_disorder_14-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_disorder_14-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_disorder_14` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.0 KiB | [pg_disorder_14-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_disorder_14-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-disorder` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.3 KiB | [postgresql-14-pg-disorder_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-disorder/postgresql-14-pg-disorder_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-disorder` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 21.1 KiB | [postgresql-14-pg-disorder_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-disorder/postgresql-14-pg-disorder_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-disorder` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.3 KiB | [postgresql-14-pg-disorder_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-disorder/postgresql-14-pg-disorder_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-disorder` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.2 KiB | [postgresql-14-pg-disorder_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-disorder/postgresql-14-pg-disorder_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-disorder` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.1 KiB | [postgresql-14-pg-disorder_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-disorder/postgresql-14-pg-disorder_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-disorder` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.7 KiB | [postgresql-14-pg-disorder_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-disorder/postgresql-14-pg-disorder_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-disorder` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.2 KiB | [postgresql-14-pg-disorder_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-disorder/postgresql-14-pg-disorder_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-disorder` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 22.1 KiB | [postgresql-14-pg-disorder_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-disorder/postgresql-14-pg-disorder_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-disorder` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 22.3 KiB | [postgresql-14-pg-disorder_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-disorder/postgresql-14-pg-disorder_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-disorder` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 21.9 KiB | [postgresql-14-pg-disorder_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-disorder/postgresql-14-pg-disorder_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/viralpraxis/pg_disorder" title="Repository" icon="github" subtitle="github.com/viralpraxis/pg_disorder" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_disorder-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_disorder;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_disorder;		# install via package name, for the active PG version

pig install pg_disorder -v 18;   # install for PG 18
pig install pg_disorder -v 17;   # install for PG 17
pig install pg_disorder -v 16;   # install for PG 16
pig install pg_disorder -v 15;   # install for PG 15
pig install pg_disorder -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_disorder';
```


This extension does not need `CREATE EXTENSION` DDL command



## Usage

Sources:

- [pg_disorder 0.1.0 README](https://api.pgxn.org/src/pg_disorder/pg_disorder-0.1.0/README.md)
- [pg_disorder 0.1.0 metadata](https://api.pgxn.org/src/pg_disorder/pg_disorder-0.1.0/META.json)
- [pg_disorder 0.1.0 Makefile](https://api.pgxn.org/src/pg_disorder/pg_disorder-0.1.0/Makefile)

`pg_disorder` is a test-only PostgreSQL loadable module that deliberately changes the output order of eligible `SELECT` queries. It helps find applications and tests that accidentally depend on unspecified row order. It is a headless module: there is no control file, SQL install script, or `CREATE EXTENSION pg_disorder` step.

### Enable It for a Test Database

Load the module at session start so its planner hook is available:

```sql
ALTER DATABASE regression_db
  SET session_preload_libraries = 'pg_disorder';

ALTER DATABASE regression_db
  SET pg_disorder.mode = 'reverse';
```

Reconnect after changing `session_preload_libraries`. Do not add this module to a production-wide `shared_preload_libraries` setting.

### Modes

```sql
SET pg_disorder.mode = 'off';
SET pg_disorder.mode = 'reverse';
SET pg_disorder.mode = 'shuffle';
SET pg_disorder.seed = 42;
SET pg_disorder.force_serial = on;
```

- `off` leaves plans unchanged.
- `reverse` deterministically reverses eligible output.
- `shuffle` produces a deterministic permutation for a fixed session seed, submitted query text, and plan. With the default seed of zero, each session first chooses and logs a random seed.
- `force_serial` suppresses parallel plans to make disorder tests reproducible.

Always fix a failing query by adding a semantically correct `ORDER BY`; do not encode the accidental order observed under `off`.

### Eligibility and Caveats

The hook targets top-level `SELECT` statements without `ORDER BY`. It deliberately skips query shapes where reordering is unsafe or changes SQL semantics, including aggregates, grouping, `DISTINCT`, set operations, window functions, recursive queries, row locks, and queries without a `FROM` relation.

- `pg_disorder` is fault-injection tooling, not a production query feature.
- Passing a disorder run does not prove every unordered query is safe; excluded query shapes and planner paths are not rewritten.
- The package installs a server module only. Verify enablement with the GUCs or module load state, not `pg_extension`.
