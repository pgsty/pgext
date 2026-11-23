---
title: "fsm_core"
linkTitle: "fsm_core"
description: "Finite state machine toolkit for PostgreSQL"
weight: 2690
categories: ["FEAT"]
languages: ["SQL"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**fsm_core**](https://github.com/pgfsm/fsm/tree/main/packages/database-src-extension/fsm_core) : Finite state machine toolkit for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2690** | {{< badge content="fsm_core" link="https://github.com/pgfsm/fsm/tree/main/packages/database-src-extension/fsm_core" >}} | {{< ext "fsm_core" >}} | `1.2.0` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "SQL" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `fsm_core` |
|   **Requires**    | {{< ext "ltree" >}} {{< ext "pgmq" >}} |
|   **See Also**    | {{< ext "pgmq" >}} {{< ext "ulak" >}} {{< ext "pgmb" >}} {{< ext "pg_durable" >}} {{< ext "redis" >}} {{< ext "pg_task" >}} {{< ext "pg_background" >}} {{< ext "pgq" >}} {{< ext "redis_fdw" >}} {{< ext "tcn" >}} |

> [!Note] PG15+; requires ltree, pgmq, and pg_jsonschema


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.0` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `fsm_core` | `ltree`, `pgmq` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.0` | {{< bg "19" "fsm_core_19" "green" >}} {{< bg "18" "fsm_core_18" "green" >}} {{< bg "17" "fsm_core_17" "green" >}} {{< bg "16" "fsm_core_16" "green" >}} {{< bg "15" "fsm_core_15" "green" >}} | `fsm_core_$v` | `pgmq_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.0` | {{< bg "19" "postgresql-19-fsm-core" "green" >}} {{< bg "18" "postgresql-18-fsm-core" "green" >}} {{< bg "17" "postgresql-17-fsm-core" "green" >}} {{< bg "16" "postgresql-16-fsm-core" "green" >}} {{< bg "15" "postgresql-15-fsm-core" "green" >}} | `postgresql-$v-fsm-core` | `postgresql-$v-pgmq`, `postgresql-$v-pg-jsonschema` |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "fsm_core_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_15 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "fsm_core_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_15 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_15 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_15 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_15 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "fsm_core_15 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-19-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-19-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-19-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-19-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-19-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-19-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-19-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-19-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-19-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-19-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-fsm-core : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-fsm-core : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `fsm_core_19` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 58.5 KiB | [fsm_core_19-1.2.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/fsm_core_19-1.2.0-1PGSTY.el9.noarch.rpm) |
| `fsm_core_19` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 58.5 KiB | [fsm_core_19-1.2.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/fsm_core_19-1.2.0-1PGSTY.el9.noarch.rpm) |
| `fsm_core_19` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 58.7 KiB | [fsm_core_19-1.2.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/fsm_core_19-1.2.0-1PGSTY.el10.noarch.rpm) |
| `fsm_core_19` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 58.7 KiB | [fsm_core_19-1.2.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/fsm_core_19-1.2.0-1PGSTY.el10.noarch.rpm) |
| `postgresql-19-fsm-core` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.5 KiB | [postgresql-19-fsm-core_1.2.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fsm-core/postgresql-19-fsm-core_1.2.0-1PGSTY~bookworm_all.deb) |
| `postgresql-19-fsm-core` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 52.5 KiB | [postgresql-19-fsm-core_1.2.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fsm-core/postgresql-19-fsm-core_1.2.0-1PGSTY~bookworm_all.deb) |
| `postgresql-19-fsm-core` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.6 KiB | [postgresql-19-fsm-core_1.2.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fsm-core/postgresql-19-fsm-core_1.2.0-1PGSTY~trixie_all.deb) |
| `postgresql-19-fsm-core` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 52.6 KiB | [postgresql-19-fsm-core_1.2.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fsm-core/postgresql-19-fsm-core_1.2.0-1PGSTY~trixie_all.deb) |
| `postgresql-19-fsm-core` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 53.6 KiB | [postgresql-19-fsm-core_1.2.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fsm-core/postgresql-19-fsm-core_1.2.0-1PGSTY~jammy_all.deb) |
| `postgresql-19-fsm-core` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 53.6 KiB | [postgresql-19-fsm-core_1.2.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fsm-core/postgresql-19-fsm-core_1.2.0-1PGSTY~jammy_all.deb) |
| `postgresql-19-fsm-core` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 53.4 KiB | [postgresql-19-fsm-core_1.2.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fsm-core/postgresql-19-fsm-core_1.2.0-1PGSTY~noble_all.deb) |
| `postgresql-19-fsm-core` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.4 KiB | [postgresql-19-fsm-core_1.2.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fsm-core/postgresql-19-fsm-core_1.2.0-1PGSTY~noble_all.deb) |
| `postgresql-19-fsm-core` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 53.5 KiB | [postgresql-19-fsm-core_1.2.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fsm-core/postgresql-19-fsm-core_1.2.0-1PGSTY~resolute_all.deb) |
| `postgresql-19-fsm-core` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 53.5 KiB | [postgresql-19-fsm-core_1.2.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fsm-core/postgresql-19-fsm-core_1.2.0-1PGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `fsm_core_18` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 63.4 KiB | [fsm_core_18-1.2.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/fsm_core_18-1.2.0-1PGSTY.el8.noarch.rpm) |
| `fsm_core_18` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 63.4 KiB | [fsm_core_18-1.2.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/fsm_core_18-1.2.0-1PGSTY.el8.noarch.rpm) |
| `fsm_core_18` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 58.5 KiB | [fsm_core_18-1.2.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/fsm_core_18-1.2.0-1PGSTY.el9.noarch.rpm) |
| `fsm_core_18` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 58.5 KiB | [fsm_core_18-1.2.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/fsm_core_18-1.2.0-1PGSTY.el9.noarch.rpm) |
| `fsm_core_18` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 58.7 KiB | [fsm_core_18-1.2.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/fsm_core_18-1.2.0-1PGSTY.el10.noarch.rpm) |
| `fsm_core_18` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 58.7 KiB | [fsm_core_18-1.2.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/fsm_core_18-1.2.0-1PGSTY.el10.noarch.rpm) |
| `postgresql-18-fsm-core` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.5 KiB | [postgresql-18-fsm-core_1.2.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fsm-core/postgresql-18-fsm-core_1.2.0-1PGSTY~bookworm_all.deb) |
| `postgresql-18-fsm-core` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 52.5 KiB | [postgresql-18-fsm-core_1.2.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fsm-core/postgresql-18-fsm-core_1.2.0-1PGSTY~bookworm_all.deb) |
| `postgresql-18-fsm-core` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.5 KiB | [postgresql-18-fsm-core_1.2.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fsm-core/postgresql-18-fsm-core_1.2.0-1PGSTY~trixie_all.deb) |
| `postgresql-18-fsm-core` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 52.5 KiB | [postgresql-18-fsm-core_1.2.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fsm-core/postgresql-18-fsm-core_1.2.0-1PGSTY~trixie_all.deb) |
| `postgresql-18-fsm-core` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 53.6 KiB | [postgresql-18-fsm-core_1.2.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fsm-core/postgresql-18-fsm-core_1.2.0-1PGSTY~jammy_all.deb) |
| `postgresql-18-fsm-core` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 53.6 KiB | [postgresql-18-fsm-core_1.2.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fsm-core/postgresql-18-fsm-core_1.2.0-1PGSTY~jammy_all.deb) |
| `postgresql-18-fsm-core` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 53.5 KiB | [postgresql-18-fsm-core_1.2.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fsm-core/postgresql-18-fsm-core_1.2.0-1PGSTY~noble_all.deb) |
| `postgresql-18-fsm-core` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.5 KiB | [postgresql-18-fsm-core_1.2.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fsm-core/postgresql-18-fsm-core_1.2.0-1PGSTY~noble_all.deb) |
| `postgresql-18-fsm-core` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 53.5 KiB | [postgresql-18-fsm-core_1.2.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fsm-core/postgresql-18-fsm-core_1.2.0-1PGSTY~resolute_all.deb) |
| `postgresql-18-fsm-core` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 53.5 KiB | [postgresql-18-fsm-core_1.2.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fsm-core/postgresql-18-fsm-core_1.2.0-1PGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `fsm_core_17` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 63.4 KiB | [fsm_core_17-1.2.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/fsm_core_17-1.2.0-1PGSTY.el8.noarch.rpm) |
| `fsm_core_17` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 63.4 KiB | [fsm_core_17-1.2.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/fsm_core_17-1.2.0-1PGSTY.el8.noarch.rpm) |
| `fsm_core_17` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 58.5 KiB | [fsm_core_17-1.2.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/fsm_core_17-1.2.0-1PGSTY.el9.noarch.rpm) |
| `fsm_core_17` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 58.5 KiB | [fsm_core_17-1.2.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/fsm_core_17-1.2.0-1PGSTY.el9.noarch.rpm) |
| `fsm_core_17` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 58.7 KiB | [fsm_core_17-1.2.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/fsm_core_17-1.2.0-1PGSTY.el10.noarch.rpm) |
| `fsm_core_17` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 58.7 KiB | [fsm_core_17-1.2.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/fsm_core_17-1.2.0-1PGSTY.el10.noarch.rpm) |
| `postgresql-17-fsm-core` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.5 KiB | [postgresql-17-fsm-core_1.2.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fsm-core/postgresql-17-fsm-core_1.2.0-1PGSTY~bookworm_all.deb) |
| `postgresql-17-fsm-core` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 52.5 KiB | [postgresql-17-fsm-core_1.2.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fsm-core/postgresql-17-fsm-core_1.2.0-1PGSTY~bookworm_all.deb) |
| `postgresql-17-fsm-core` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.5 KiB | [postgresql-17-fsm-core_1.2.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fsm-core/postgresql-17-fsm-core_1.2.0-1PGSTY~trixie_all.deb) |
| `postgresql-17-fsm-core` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 52.5 KiB | [postgresql-17-fsm-core_1.2.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fsm-core/postgresql-17-fsm-core_1.2.0-1PGSTY~trixie_all.deb) |
| `postgresql-17-fsm-core` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 53.6 KiB | [postgresql-17-fsm-core_1.2.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fsm-core/postgresql-17-fsm-core_1.2.0-1PGSTY~jammy_all.deb) |
| `postgresql-17-fsm-core` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 53.6 KiB | [postgresql-17-fsm-core_1.2.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fsm-core/postgresql-17-fsm-core_1.2.0-1PGSTY~jammy_all.deb) |
| `postgresql-17-fsm-core` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 53.4 KiB | [postgresql-17-fsm-core_1.2.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fsm-core/postgresql-17-fsm-core_1.2.0-1PGSTY~noble_all.deb) |
| `postgresql-17-fsm-core` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.4 KiB | [postgresql-17-fsm-core_1.2.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fsm-core/postgresql-17-fsm-core_1.2.0-1PGSTY~noble_all.deb) |
| `postgresql-17-fsm-core` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 53.5 KiB | [postgresql-17-fsm-core_1.2.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fsm-core/postgresql-17-fsm-core_1.2.0-1PGSTY~resolute_all.deb) |
| `postgresql-17-fsm-core` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 53.5 KiB | [postgresql-17-fsm-core_1.2.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fsm-core/postgresql-17-fsm-core_1.2.0-1PGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `fsm_core_16` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 63.4 KiB | [fsm_core_16-1.2.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/fsm_core_16-1.2.0-1PGSTY.el8.noarch.rpm) |
| `fsm_core_16` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 63.4 KiB | [fsm_core_16-1.2.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/fsm_core_16-1.2.0-1PGSTY.el8.noarch.rpm) |
| `fsm_core_16` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 58.5 KiB | [fsm_core_16-1.2.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/fsm_core_16-1.2.0-1PGSTY.el9.noarch.rpm) |
| `fsm_core_16` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 58.5 KiB | [fsm_core_16-1.2.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/fsm_core_16-1.2.0-1PGSTY.el9.noarch.rpm) |
| `fsm_core_16` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 58.7 KiB | [fsm_core_16-1.2.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/fsm_core_16-1.2.0-1PGSTY.el10.noarch.rpm) |
| `fsm_core_16` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 58.7 KiB | [fsm_core_16-1.2.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/fsm_core_16-1.2.0-1PGSTY.el10.noarch.rpm) |
| `postgresql-16-fsm-core` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.5 KiB | [postgresql-16-fsm-core_1.2.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fsm-core/postgresql-16-fsm-core_1.2.0-1PGSTY~bookworm_all.deb) |
| `postgresql-16-fsm-core` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 52.5 KiB | [postgresql-16-fsm-core_1.2.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fsm-core/postgresql-16-fsm-core_1.2.0-1PGSTY~bookworm_all.deb) |
| `postgresql-16-fsm-core` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.6 KiB | [postgresql-16-fsm-core_1.2.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fsm-core/postgresql-16-fsm-core_1.2.0-1PGSTY~trixie_all.deb) |
| `postgresql-16-fsm-core` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 52.6 KiB | [postgresql-16-fsm-core_1.2.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fsm-core/postgresql-16-fsm-core_1.2.0-1PGSTY~trixie_all.deb) |
| `postgresql-16-fsm-core` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 53.6 KiB | [postgresql-16-fsm-core_1.2.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fsm-core/postgresql-16-fsm-core_1.2.0-1PGSTY~jammy_all.deb) |
| `postgresql-16-fsm-core` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 53.6 KiB | [postgresql-16-fsm-core_1.2.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fsm-core/postgresql-16-fsm-core_1.2.0-1PGSTY~jammy_all.deb) |
| `postgresql-16-fsm-core` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 53.4 KiB | [postgresql-16-fsm-core_1.2.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fsm-core/postgresql-16-fsm-core_1.2.0-1PGSTY~noble_all.deb) |
| `postgresql-16-fsm-core` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.4 KiB | [postgresql-16-fsm-core_1.2.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fsm-core/postgresql-16-fsm-core_1.2.0-1PGSTY~noble_all.deb) |
| `postgresql-16-fsm-core` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 53.5 KiB | [postgresql-16-fsm-core_1.2.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fsm-core/postgresql-16-fsm-core_1.2.0-1PGSTY~resolute_all.deb) |
| `postgresql-16-fsm-core` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 53.5 KiB | [postgresql-16-fsm-core_1.2.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fsm-core/postgresql-16-fsm-core_1.2.0-1PGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `fsm_core_15` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 63.4 KiB | [fsm_core_15-1.2.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/fsm_core_15-1.2.0-1PGSTY.el8.noarch.rpm) |
| `fsm_core_15` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 63.4 KiB | [fsm_core_15-1.2.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/fsm_core_15-1.2.0-1PGSTY.el8.noarch.rpm) |
| `fsm_core_15` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 58.5 KiB | [fsm_core_15-1.2.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/fsm_core_15-1.2.0-1PGSTY.el9.noarch.rpm) |
| `fsm_core_15` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 58.5 KiB | [fsm_core_15-1.2.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/fsm_core_15-1.2.0-1PGSTY.el9.noarch.rpm) |
| `fsm_core_15` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 58.7 KiB | [fsm_core_15-1.2.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/fsm_core_15-1.2.0-1PGSTY.el10.noarch.rpm) |
| `fsm_core_15` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 58.7 KiB | [fsm_core_15-1.2.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/fsm_core_15-1.2.0-1PGSTY.el10.noarch.rpm) |
| `postgresql-15-fsm-core` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.5 KiB | [postgresql-15-fsm-core_1.2.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fsm-core/postgresql-15-fsm-core_1.2.0-1PGSTY~bookworm_all.deb) |
| `postgresql-15-fsm-core` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 52.5 KiB | [postgresql-15-fsm-core_1.2.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fsm-core/postgresql-15-fsm-core_1.2.0-1PGSTY~bookworm_all.deb) |
| `postgresql-15-fsm-core` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.6 KiB | [postgresql-15-fsm-core_1.2.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fsm-core/postgresql-15-fsm-core_1.2.0-1PGSTY~trixie_all.deb) |
| `postgresql-15-fsm-core` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 52.6 KiB | [postgresql-15-fsm-core_1.2.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fsm-core/postgresql-15-fsm-core_1.2.0-1PGSTY~trixie_all.deb) |
| `postgresql-15-fsm-core` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 53.6 KiB | [postgresql-15-fsm-core_1.2.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fsm-core/postgresql-15-fsm-core_1.2.0-1PGSTY~jammy_all.deb) |
| `postgresql-15-fsm-core` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 53.6 KiB | [postgresql-15-fsm-core_1.2.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fsm-core/postgresql-15-fsm-core_1.2.0-1PGSTY~jammy_all.deb) |
| `postgresql-15-fsm-core` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 53.5 KiB | [postgresql-15-fsm-core_1.2.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fsm-core/postgresql-15-fsm-core_1.2.0-1PGSTY~noble_all.deb) |
| `postgresql-15-fsm-core` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.5 KiB | [postgresql-15-fsm-core_1.2.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fsm-core/postgresql-15-fsm-core_1.2.0-1PGSTY~noble_all.deb) |
| `postgresql-15-fsm-core` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 53.5 KiB | [postgresql-15-fsm-core_1.2.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fsm-core/postgresql-15-fsm-core_1.2.0-1PGSTY~resolute_all.deb) |
| `postgresql-15-fsm-core` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 53.5 KiB | [postgresql-15-fsm-core_1.2.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fsm-core/postgresql-15-fsm-core_1.2.0-1PGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgfsm/fsm/tree/main/packages/database-src-extension/fsm_core" title="Repository" icon="github" subtitle="github.com/pgfsm/fsm/tree/main/packages/database-src-extension/fsm_core" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="fsm_core-1.2.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg fsm_core;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install fsm_core;		# install via package name, for the active PG version

pig install fsm_core -v 19;   # install for PG 19
pig install fsm_core -v 18;   # install for PG 18
pig install fsm_core -v 17;   # install for PG 17
pig install fsm_core -v 16;   # install for PG 16
pig install fsm_core -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION fsm_core CASCADE; -- requires ltree, pgmq
```

## Usage

Sources:

- [fsm_core-v1.2.0](https://github.com/pgfsm/fsm/releases/tag/fsm_core-v1.2.0)
- [README.md](https://github.com/pgfsm/fsm/blob/fsm_core-v1.2.0/packages/database-src/README.md)
- [fsm_core-1.2.0.zip](https://github.com/pgfsm/fsm/releases/download/fsm_core-v1.2.0/fsm_core-1.2.0.zip)


`fsm_core` 1.2.0 stores finite-state-machine definitions, instances, transitions and event history in PostgreSQL. Its SQL interface loads validated JSON, creates instances by name and version, and submits events through `pgmq` queues.

### Install and Configure

The release requires PostgreSQL 15 or later and was tested on 15–18. Required extensions are `ltree` 1.2+, `pgmq` 1.4.4+ and `pg_cron` 1.5+. Add `pg_cron` to the existing `shared_preload_libraries`, restart PostgreSQL if that setting changes, and install into the database selected by `cron.database_name`.

After installing the matching extension files, a superuser can enable the fixed, nonrelocatable `fsm_core` schema:

```sql
CREATE EXTENSION fsm_core CASCADE;
SELECT fsm_core.version();
```

Installation creates `fsm_*` roles and the `fsm_schedule_all_pending` scheduler job, which runs every five seconds. Inspect role memberships, queue permissions and the scheduler before granting application access. `pg_jsonschema` is optional: if its validation helper exists in `fsm_core`, the loader reports schema problems but does not reject a definition on that check alone.

### Load a Definition

Validate the machine with the upstream compiler and set the psql variable fsm_json to that JSON before this call. The 1.2.0 signature has no input_fsm_type argument:

```sql
SELECT fsm_core.load_fsm_from_json_v2(
  json_input        := :'fsm_json'::jsonb,
  root_node_text    := 'root',
  input_fsm_name    := 'creditCheck',
  input_fsm_version := 'v01'
);
```

Concurrent loads of the same name and version are serialized by an advisory transaction lock. Identical JSON returns the cached definition; different JSON under an existing name and version raises an error. Use a new version for a changed machine. The optional input_dependent_children argument records child-machine dependencies.

### Create an Instance and Send Events

```sql
SELECT fsm_core.create_fsm_instance_from_name_v2(
  input_fsm_name    := 'creditCheck',
  input_fsm_version := 'v01',
  input_fsm_context := '{"applicant_id":"a-42"}'::jsonb,
  create_pgmq_queue := true
) AS creation_result
\gset
SELECT :'creation_result'::jsonb AS creation_status;
SELECT :'creation_result'::jsonb ->> 'fsm_instance_id' AS fsm_instance_id
\gset
```

The returned JSON includes `fsm_instance_id`, `queue_created`, `send_event_result`, `message` and `extra_message`. Check these before submitting work: queue creation and initial-event failures can be reported in the result. Initial context is shallow-merged with the definition's static context, with supplied object keys taking precedence.

For this top-level machine, submit an event declared by its definition:

```sql
SELECT fsm_core.send_event_to_fsm_queue_with_event_logs_v2(
  input_fsm_instance_id            := :'fsm_instance_id'::uuid,
  input_fsm_instance_id_fsm_type   := 'fsm',
  input_fsm_instance_id_fsm_version := 'v01',
  input_event_name                 := 'APPROVE',
  input_event_action_type          := 'user',
  input_event_data                 := '{"approved_by":"manager"}'::jsonb,
  input_event_delay                := 0
);
```

The helper queues the event and writes event history. The old parent-queue parameters are absent from this release's signature. Queuing an event is not proof that a worker has processed it; monitor scheduler activity, queue backlog and instance state.

### Objects and Upgrade Boundaries

`fsm_json`, `fsm_states` and `fsm_transitions` hold definitions; `fsm_instance` holds instances; `fsm_instance_queue_event_logs` records queue events. `resolve_state_value_v2` and `macrostep_v2` expose lower-level state resolution and transition logic. Prefer the verified v2 entry points and restrict access to security-definer functions and underlying tables.

The release notes describe sequential upgrade scripts guarded by `require_version`, but the reviewed 1.2.0 ZIP contains the full install script and no version-to-version upgrade scripts. Preserve populated databases and obtain a verified upgrade artifact before upgrading an existing installation. Do not rerun the full install or assume `ALTER EXTENSION ... UPDATE` has an available upgrade path.
