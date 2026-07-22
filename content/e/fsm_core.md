---
title: "fsm_core"
linkTitle: "fsm_core"
description: "Finite state machine toolkit for PostgreSQL"
weight: 2690
categories: ["FEAT"]
width: full
---

[**fsm_core**](https://github.com/Nirajkashyap/fsm/tree/main/packages/database-src-extension/fsm_core) : Finite state machine toolkit for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2690** | {{< badge content="fsm_core" link="https://github.com/Nirajkashyap/fsm/tree/main/packages/database-src-extension/fsm_core" >}} | {{< ext "fsm_core" >}} | `1.1.0` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `fsm_core` |
|   **Requires**    | {{< ext "ltree" >}} {{< ext "pgmq" >}} {{< ext "pg_jsonschema" >}} |
|   **See Also**    | {{< ext "pgmq" >}} {{< ext "pg_jsonschema" >}} {{< ext "pg_task" >}} {{< ext "pg_later" >}} {{< ext "pg_cron" >}} |

> [!Note] PG15+; requires ltree, pgmq, and pg_jsonschema


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `fsm_core` | `ltree`, `pgmq`, `pg_jsonschema` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "fsm_core_18" "green" >}} {{< bg "17" "fsm_core_17" "green" >}} {{< bg "16" "fsm_core_16" "green" >}} {{< bg "15" "fsm_core_15" "green" >}} {{< bg "14" "fsm_core_14" "red" >}} | `fsm_core_$v` | `pgmq_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "postgresql-18-fsm-core" "green" >}} {{< bg "17" "postgresql-17-fsm-core" "green" >}} {{< bg "16" "postgresql-16-fsm-core" "green" >}} {{< bg "15" "postgresql-15-fsm-core" "green" >}} {{< bg "14" "postgresql-14-fsm-core" "red" >}} | `postgresql-$v-fsm-core` | `postgresql-$v-pgmq`, `postgresql-$v-pg-jsonschema` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "fsm_core_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "fsm_core_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "fsm_core_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "fsm_core_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "fsm_core_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "fsm_core_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "fsm_core_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-fsm-core : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-fsm-core : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-fsm-core : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-fsm-core : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-fsm-core : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-fsm-core : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-fsm-core : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-fsm-core : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-fsm-core : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-fsm-core : N/A 0" "gray" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `fsm_core_18` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 33.0 KiB | [fsm_core_18-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/fsm_core_18-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `fsm_core_18` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.0 KiB | [fsm_core_18-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/fsm_core_18-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `fsm_core_18` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.8 KiB | [fsm_core_18-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/fsm_core_18-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `fsm_core_18` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.7 KiB | [fsm_core_18-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/fsm_core_18-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `fsm_core_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.9 KiB | [fsm_core_18-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/fsm_core_18-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `fsm_core_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.8 KiB | [fsm_core_18-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/fsm_core_18-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-fsm-core` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.4 KiB | [postgresql-18-fsm-core_1.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fsm-core/postgresql-18-fsm-core_1.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-fsm-core` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.4 KiB | [postgresql-18-fsm-core_1.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fsm-core/postgresql-18-fsm-core_1.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-fsm-core` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.4 KiB | [postgresql-18-fsm-core_1.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fsm-core/postgresql-18-fsm-core_1.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-fsm-core` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.4 KiB | [postgresql-18-fsm-core_1.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fsm-core/postgresql-18-fsm-core_1.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-fsm-core` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.7 KiB | [postgresql-18-fsm-core_1.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fsm-core/postgresql-18-fsm-core_1.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-fsm-core` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.7 KiB | [postgresql-18-fsm-core_1.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fsm-core/postgresql-18-fsm-core_1.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-fsm-core` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.7 KiB | [postgresql-18-fsm-core_1.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fsm-core/postgresql-18-fsm-core_1.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-fsm-core` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 24.7 KiB | [postgresql-18-fsm-core_1.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fsm-core/postgresql-18-fsm-core_1.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-fsm-core` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 24.7 KiB | [postgresql-18-fsm-core_1.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fsm-core/postgresql-18-fsm-core_1.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-18-fsm-core` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 24.7 KiB | [postgresql-18-fsm-core_1.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fsm-core/postgresql-18-fsm-core_1.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `fsm_core_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 33.0 KiB | [fsm_core_17-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/fsm_core_17-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `fsm_core_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.0 KiB | [fsm_core_17-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/fsm_core_17-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `fsm_core_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.7 KiB | [fsm_core_17-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/fsm_core_17-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `fsm_core_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.7 KiB | [fsm_core_17-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/fsm_core_17-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `fsm_core_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.9 KiB | [fsm_core_17-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/fsm_core_17-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `fsm_core_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.8 KiB | [fsm_core_17-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/fsm_core_17-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-fsm-core` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.4 KiB | [postgresql-17-fsm-core_1.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fsm-core/postgresql-17-fsm-core_1.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-fsm-core` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.4 KiB | [postgresql-17-fsm-core_1.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fsm-core/postgresql-17-fsm-core_1.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-fsm-core` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.4 KiB | [postgresql-17-fsm-core_1.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fsm-core/postgresql-17-fsm-core_1.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-fsm-core` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.4 KiB | [postgresql-17-fsm-core_1.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fsm-core/postgresql-17-fsm-core_1.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-fsm-core` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.7 KiB | [postgresql-17-fsm-core_1.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fsm-core/postgresql-17-fsm-core_1.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-fsm-core` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.7 KiB | [postgresql-17-fsm-core_1.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fsm-core/postgresql-17-fsm-core_1.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-fsm-core` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.7 KiB | [postgresql-17-fsm-core_1.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fsm-core/postgresql-17-fsm-core_1.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-fsm-core` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 24.7 KiB | [postgresql-17-fsm-core_1.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fsm-core/postgresql-17-fsm-core_1.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-fsm-core` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 24.7 KiB | [postgresql-17-fsm-core_1.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fsm-core/postgresql-17-fsm-core_1.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-17-fsm-core` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 24.7 KiB | [postgresql-17-fsm-core_1.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fsm-core/postgresql-17-fsm-core_1.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `fsm_core_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 33.0 KiB | [fsm_core_16-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/fsm_core_16-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `fsm_core_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.0 KiB | [fsm_core_16-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/fsm_core_16-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `fsm_core_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.8 KiB | [fsm_core_16-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/fsm_core_16-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `fsm_core_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.7 KiB | [fsm_core_16-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/fsm_core_16-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `fsm_core_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.9 KiB | [fsm_core_16-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/fsm_core_16-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `fsm_core_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.8 KiB | [fsm_core_16-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/fsm_core_16-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-fsm-core` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.4 KiB | [postgresql-16-fsm-core_1.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fsm-core/postgresql-16-fsm-core_1.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-fsm-core` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.4 KiB | [postgresql-16-fsm-core_1.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fsm-core/postgresql-16-fsm-core_1.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-fsm-core` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.4 KiB | [postgresql-16-fsm-core_1.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fsm-core/postgresql-16-fsm-core_1.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-fsm-core` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.4 KiB | [postgresql-16-fsm-core_1.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fsm-core/postgresql-16-fsm-core_1.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-fsm-core` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.7 KiB | [postgresql-16-fsm-core_1.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fsm-core/postgresql-16-fsm-core_1.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-fsm-core` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.7 KiB | [postgresql-16-fsm-core_1.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fsm-core/postgresql-16-fsm-core_1.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-fsm-core` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.7 KiB | [postgresql-16-fsm-core_1.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fsm-core/postgresql-16-fsm-core_1.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-fsm-core` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 24.7 KiB | [postgresql-16-fsm-core_1.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fsm-core/postgresql-16-fsm-core_1.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-fsm-core` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 24.7 KiB | [postgresql-16-fsm-core_1.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fsm-core/postgresql-16-fsm-core_1.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-16-fsm-core` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 24.7 KiB | [postgresql-16-fsm-core_1.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fsm-core/postgresql-16-fsm-core_1.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `fsm_core_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 33.0 KiB | [fsm_core_15-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/fsm_core_15-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `fsm_core_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.0 KiB | [fsm_core_15-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/fsm_core_15-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `fsm_core_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.7 KiB | [fsm_core_15-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/fsm_core_15-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `fsm_core_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.7 KiB | [fsm_core_15-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/fsm_core_15-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `fsm_core_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.9 KiB | [fsm_core_15-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/fsm_core_15-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `fsm_core_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.8 KiB | [fsm_core_15-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/fsm_core_15-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-fsm-core` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.4 KiB | [postgresql-15-fsm-core_1.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fsm-core/postgresql-15-fsm-core_1.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-fsm-core` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.4 KiB | [postgresql-15-fsm-core_1.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fsm-core/postgresql-15-fsm-core_1.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-fsm-core` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.4 KiB | [postgresql-15-fsm-core_1.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fsm-core/postgresql-15-fsm-core_1.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-fsm-core` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.4 KiB | [postgresql-15-fsm-core_1.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fsm-core/postgresql-15-fsm-core_1.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-fsm-core` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.7 KiB | [postgresql-15-fsm-core_1.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fsm-core/postgresql-15-fsm-core_1.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-fsm-core` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.7 KiB | [postgresql-15-fsm-core_1.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fsm-core/postgresql-15-fsm-core_1.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-fsm-core` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.7 KiB | [postgresql-15-fsm-core_1.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fsm-core/postgresql-15-fsm-core_1.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-fsm-core` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 24.7 KiB | [postgresql-15-fsm-core_1.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fsm-core/postgresql-15-fsm-core_1.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-fsm-core` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 24.7 KiB | [postgresql-15-fsm-core_1.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fsm-core/postgresql-15-fsm-core_1.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-15-fsm-core` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 24.7 KiB | [postgresql-15-fsm-core_1.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fsm-core/postgresql-15-fsm-core_1.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Nirajkashyap/fsm/tree/main/packages/database-src-extension/fsm_core" title="Repository" icon="github" subtitle="github.com/Nirajkashyap/fsm/tree/main/packages/database-src-extension/fsm_core" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="fsm_core-1.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg fsm_core;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install fsm_core;		# install via package name, for the active PG version

pig install fsm_core -v 18;   # install for PG 18
pig install fsm_core -v 17;   # install for PG 17
pig install fsm_core -v 16;   # install for PG 16
pig install fsm_core -v 15;   # install for PG 15

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION fsm_core CASCADE; -- requires ltree, pgmq, pg_jsonschema
```




## Usage

Sources:

- [fsm_core PGXN README](https://github.com/pgfsm/fsm/blob/main/packages/database-src/pgxn-dist/README.md)
- [PGXN control file](https://github.com/pgfsm/fsm/blob/main/packages/database-src/pgxn-dist/fsm_core.control)
- [1.1.0 SQL definition](https://github.com/pgfsm/fsm/blob/main/packages/database-src/pgxn-dist/fsm_core--1.1.0.sql)
- [Supabase 1.1.1-1.1.3 migrations](https://github.com/pgfsm/fsm/tree/main/packages/database-src/supabase/migrations)
- [worker execution ADR](https://github.com/pgfsm/fsm/blob/main/apps/fsm-core-worker-ts/docs/ADR-001-worker-execution-model.md)
- [example definitions README](https://github.com/pgfsm/fsm/blob/main/apps/fsm-core-example/README.md)

`fsm_core` is a finite-state-machine toolkit that stores FSM definitions, instances, transitions, dispatch queues, and event logs inside PostgreSQL. A machine definition is loaded from JSON, instances are created by name and version, and events are sent through SQL functions with optional `pgmq` queues or the newer scheduler dispatch path.

The PGXN packaged extension still declares `default_version = '1.1.0'` and requires PostgreSQL 15+, `ltree`, and `pgmq`. The repository's `packages/database-src` project is at `1.1.3` and contains Supabase migration files through `fsm_core--1.1.2--1.1.3.sql`; the Rust/pgrx directory is explicitly marked not currently used. Treat the PGXN distribution as the packaged source and the Supabase migrations as main-branch material.

### Core Tables and Types

`fsm_core` creates an enum `fsm_state_type` with `atomic`, `compound`, `parallel`, `final`, and `history`, plus tables including:

- `fsm_core.fsm_json` for loaded FSM definitions.
- `fsm_core.fsm_states` for expanded state nodes and ltree paths.
- `fsm_core.fsm_transitions` for transition rules.
- `fsm_core.fsm_instance` for running instances.
- `fsm_core.fsm_instance_lock` for advisory/concurrency state.
- `fsm_core.fsm_instance_queue_event_logs` and `fsm_core.fsm_promise_queue_event_logs` for queued event history.
- `fsm_core.fsm_dispatch_queue` and `fsm_core.fsm_daemon_node` in the main-branch scheduler path.

### Load a Machine Definition

```sql
SELECT fsm_core.load_fsm_from_json_v2(
  json_input        := :'fsm_json'::jsonb,
  root_node_text    := 'root',
  input_fsm_type    := 'workflow',
  input_fsm_name    := 'creditCheck',
  input_fsm_version := 'v01'
);
```

`load_fsm_from_json_v2()` validates JSON with `fsm_core.fsm_json_schema()`, caches the raw definition in `fsm_json`, then expands states and transitions. Upstream examples keep immutable version folders such as `fsm/creditCheck/v01/xstate-fsm.json` and `fsm/creditCheck/v02/xstate-fsm.json`; keep deployed versions immutable so existing instances continue against their original definition.

### Create an Instance

```sql
SELECT fsm_core.create_fsm_instance_from_name_v2(
  input_fsm_name     := 'creditCheck',
  input_fsm_version  := 'v01',
  input_fsm_context  := '{"applicant_id":"a-42"}'::jsonb,
  create_pgmq_queue  := true
);
```

The PGXN 1.1.0 function checks that the named FSM exists, inserts an `fsm_instance`, copies transition authorization rows for that instance, and, when `create_pgmq_queue` is true, creates a `pgmq` queue named by the instance UUID and sends `initialTransition_event`.

The main-branch Supabase schema also has a scheduler-oriented path where instance creation can enqueue work through `fsm_core.enqueue_fsm_dispatch_v2()` instead of relying only on per-instance `pgmq` queues.

### Send Events

```sql
SELECT fsm_core.send_event_to_fsm_queue_with_event_logs_v2(
  input_fsm_instance_id                 := '00000000-0000-0000-0000-000000000000'::uuid,
  input_fsm_instance_id_fsm_type         := 'workflow',
  input_fsm_instance_id_fsm_version      := 'v01',
  input_send_to_parent_queue_id          := fsm_core.pg_system_queue_uuid(),
  input_send_to_parent_queue_type        := fsm_core.pg_system_queue_type(),
  input_send_to_parent_queue_id_event_name := fsm_core.pg_system_event_name(),
  input_event_name                       := 'APPROVE',
  input_event_action_type                := 'user',
  input_event_data                       := '{"approved_by":"manager"}'::jsonb,
  input_event_delay                      := 0
);
```

This helper writes to the instance queue with `pgmq.send()` and records the event in `fsm_instance_queue_event_logs`. For nested FSM and promise flows, `send_event_to_queue_from_fsm_instance_id_v2()` dispatches to the child-FSM or promise queue helper based on `fsmtype`.

### Scheduler Dispatch Path

```sql
SELECT fsm_core.enqueue_fsm_dispatch_v2(
  input_instance_id   := '00000000-0000-0000-0000-000000000000'::uuid,
  input_fsm_name      := 'creditCheck',
  input_fsm_version   := 'v01',
  input_dispatch_type := 'start'
);

SELECT fsm_core.schedule_next_pending();
```

The newer worker design uses `fsm_dispatch_queue` as the source of pending work and `fsm_daemon_node` as the registry of available fsmlet nodes. `enqueue_fsm_dispatch_v2()` inserts a pending dispatch row and emits `pg_notify('fsm_scheduler_work', instance_id)`. `schedule_next_pending()` selects pending rows with `FOR UPDATE SKIP LOCKED`, chooses a daemon under its concurrency limit, marks the row scheduled, and notifies `fsm_fsmlet_work_<daemon_id>`.

This scheduler path appears in the main-branch Supabase schema and worker ADR, not in the PGXN 1.1.0 packaged SQL.

### Resolve and Step State

```sql
SELECT fsm_core.resolve_state_value_v2(
  input_json        := '{"value":"pending"}'::jsonb,
  input_fsm_name    := 'creditCheck',
  input_fsm_version := 'v01'
);

SELECT fsm_core.macrostep_v2(
  event_name        := 'APPROVE',
  input_state_value := ARRAY['pending']::text[],
  fsm_name_param    := 'creditCheck',
  fsm_version_param := 'v01'
);
```

The SQL surface also includes lower-level `microstep_v2()`, `fsm_worker_v2()`, lock helpers, archive helpers, and v1 compatibility functions. Prefer v2 entry points for new usage when both versions are present.

### Dependency and Source Caveats

- Enable `ltree` and `pgmq` before `fsm_core`; account for `pg_jsonschema` because upstream Supabase setup and Pigsty package dependencies list it, even though the PGXN control file only declares `ltree, pgmq`.
- `pgmq` queues and `fsm_dispatch_queue` rows are part of the async execution model, so queue names, retention, and scheduler cleanup should be operated like application data.
- The upstream repo has no release tag for `1.1.0` or `1.1.3`; the authoritative packaged source for PGXN is `packages/database-src/pgxn-dist`.
- The `packages/database-src-extension` Rust/pgrx tree is exploratory and marked not currently used. Do not treat it as the active extension implementation.
- Public docs are sparse compared with the SQL surface. Treat unlisted helper functions as internal unless the SQL definition, Supabase migration, or worker ADR demonstrates their role.
