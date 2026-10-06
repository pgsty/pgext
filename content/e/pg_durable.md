---
title: "pg_durable"
linkTitle: "pg_durable"
description: "Durable SQL functions for PostgreSQL"
weight: 2870
categories: ["FEAT"]
languages: ["Rust"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_durable**](https://github.com/microsoft/pg_durable) : Durable SQL functions for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2870** | {{< badge content="pg_durable" link="https://github.com/microsoft/pg_durable" >}} | {{< ext "pg_durable" >}} | `0.2.8` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` `df` `_duroxide` |
|   **See Also**    | {{< ext "pg_task" >}} {{< ext "pgmq" >}} {{< ext "pg_background" >}} {{< ext "ulak" >}} {{< ext "pgmb" >}} {{< ext "pg_later" >}} {{< ext "pg_dispatch" >}} {{< ext "pg_retry" >}} {{< ext "fsm_core" >}} {{< ext "pglock" >}} |

> [!Note] Requires preload and a superuser worker role; pgrx 0.19.2.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.8` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_durable` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.8` | {{< bg "18" "pg_durable_18" "green" >}} {{< bg "17" "pg_durable_17" "green" >}} {{< bg "16" "pg_durable_16" "green" >}} {{< bg "15" "pg_durable_15" "green" >}} {{< bg "14" "pg_durable_14" "green" >}} | `pg_durable_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.8` | {{< bg "18" "postgresql-18-pg-durable" "green" >}} {{< bg "17" "postgresql-17-pg-durable" "green" >}} {{< bg "16" "postgresql-16-pg-durable" "green" >}} {{< bg "15" "postgresql-15-pg-durable" "green" >}} {{< bg "14" "postgresql-14-pg-durable" "green" >}} | `postgresql-$v-pg-durable` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "pg_durable_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.8" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_durable_18` | `0.2.8` | [el8.x86_64](/os/el8.x86_64) | pigsty | 5.1 MiB | [pg_durable_18-0.2.8-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_durable_18-0.2.8-1PGSTY.el8.x86_64.rpm) |
| `pg_durable_18` | `0.2.8` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.1 MiB | [pg_durable_18-0.2.8-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_durable_18-0.2.8-1PGSTY.el8.aarch64.rpm) |
| `pg_durable_18` | `0.2.8` | [el9.x86_64](/os/el9.x86_64) | pigsty | 5.0 MiB | [pg_durable_18-0.2.8-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_durable_18-0.2.8-1PGSTY.el9.x86_64.rpm) |
| `pg_durable_18` | `0.2.8` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.3 MiB | [pg_durable_18-0.2.8-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_durable_18-0.2.8-1PGSTY.el9.aarch64.rpm) |
| `pg_durable_18` | `0.2.8` | [el10.x86_64](/os/el10.x86_64) | pigsty | 5.0 MiB | [pg_durable_18-0.2.8-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_durable_18-0.2.8-1PGSTY.el10.x86_64.rpm) |
| `pg_durable_18` | `0.2.8` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.3 MiB | [pg_durable_18-0.2.8-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_durable_18-0.2.8-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-durable` | `0.2.8` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.2 MiB | [postgresql-18-pg-durable_0.2.8-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.8-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-durable` | `0.2.8` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.3 MiB | [postgresql-18-pg-durable_0.2.8-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.8-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-durable` | `0.2.8` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.2 MiB | [postgresql-18-pg-durable_0.2.8-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.8-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-durable` | `0.2.8` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.3 MiB | [postgresql-18-pg-durable_0.2.8-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.8-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-durable` | `0.2.8` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.6 MiB | [postgresql-18-pg-durable_0.2.8-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.8-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-durable` | `0.2.8` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.9 MiB | [postgresql-18-pg-durable_0.2.8-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.8-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-durable` | `0.2.8` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.6 MiB | [postgresql-18-pg-durable_0.2.8-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.8-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-durable` | `0.2.8` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.9 MiB | [postgresql-18-pg-durable_0.2.8-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.8-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-durable` | `0.2.8` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.6 MiB | [postgresql-18-pg-durable_0.2.8-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.8-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-durable` | `0.2.8` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.9 MiB | [postgresql-18-pg-durable_0.2.8-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.8-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_durable_17` | `0.2.8` | [el8.x86_64](/os/el8.x86_64) | pigsty | 5.1 MiB | [pg_durable_17-0.2.8-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_durable_17-0.2.8-1PGSTY.el8.x86_64.rpm) |
| `pg_durable_17` | `0.2.8` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.0 MiB | [pg_durable_17-0.2.8-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_durable_17-0.2.8-1PGSTY.el8.aarch64.rpm) |
| `pg_durable_17` | `0.2.8` | [el9.x86_64](/os/el9.x86_64) | pigsty | 5.0 MiB | [pg_durable_17-0.2.8-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_durable_17-0.2.8-1PGSTY.el9.x86_64.rpm) |
| `pg_durable_17` | `0.2.8` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.3 MiB | [pg_durable_17-0.2.8-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_durable_17-0.2.8-1PGSTY.el9.aarch64.rpm) |
| `pg_durable_17` | `0.2.8` | [el10.x86_64](/os/el10.x86_64) | pigsty | 5.0 MiB | [pg_durable_17-0.2.8-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_durable_17-0.2.8-1PGSTY.el10.x86_64.rpm) |
| `pg_durable_17` | `0.2.8` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.3 MiB | [pg_durable_17-0.2.8-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_durable_17-0.2.8-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-durable` | `0.2.8` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.2 MiB | [postgresql-17-pg-durable_0.2.8-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.8-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-durable` | `0.2.8` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.3 MiB | [postgresql-17-pg-durable_0.2.8-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.8-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-durable` | `0.2.8` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.2 MiB | [postgresql-17-pg-durable_0.2.8-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.8-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-durable` | `0.2.8` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.3 MiB | [postgresql-17-pg-durable_0.2.8-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.8-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-durable` | `0.2.8` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.6 MiB | [postgresql-17-pg-durable_0.2.8-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.8-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-durable` | `0.2.8` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.9 MiB | [postgresql-17-pg-durable_0.2.8-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.8-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-durable` | `0.2.8` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.6 MiB | [postgresql-17-pg-durable_0.2.8-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.8-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-durable` | `0.2.8` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.9 MiB | [postgresql-17-pg-durable_0.2.8-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.8-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-durable` | `0.2.8` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.6 MiB | [postgresql-17-pg-durable_0.2.8-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.8-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-durable` | `0.2.8` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.9 MiB | [postgresql-17-pg-durable_0.2.8-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.8-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_durable_16` | `0.2.8` | [el8.x86_64](/os/el8.x86_64) | pigsty | 5.1 MiB | [pg_durable_16-0.2.8-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_durable_16-0.2.8-1PGSTY.el8.x86_64.rpm) |
| `pg_durable_16` | `0.2.8` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.0 MiB | [pg_durable_16-0.2.8-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_durable_16-0.2.8-1PGSTY.el8.aarch64.rpm) |
| `pg_durable_16` | `0.2.8` | [el9.x86_64](/os/el9.x86_64) | pigsty | 5.0 MiB | [pg_durable_16-0.2.8-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_durable_16-0.2.8-1PGSTY.el9.x86_64.rpm) |
| `pg_durable_16` | `0.2.8` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.3 MiB | [pg_durable_16-0.2.8-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_durable_16-0.2.8-1PGSTY.el9.aarch64.rpm) |
| `pg_durable_16` | `0.2.8` | [el10.x86_64](/os/el10.x86_64) | pigsty | 5.0 MiB | [pg_durable_16-0.2.8-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_durable_16-0.2.8-1PGSTY.el10.x86_64.rpm) |
| `pg_durable_16` | `0.2.8` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.3 MiB | [pg_durable_16-0.2.8-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_durable_16-0.2.8-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-durable` | `0.2.8` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.2 MiB | [postgresql-16-pg-durable_0.2.8-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.8-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-durable` | `0.2.8` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.3 MiB | [postgresql-16-pg-durable_0.2.8-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.8-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-durable` | `0.2.8` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.2 MiB | [postgresql-16-pg-durable_0.2.8-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.8-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-durable` | `0.2.8` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.3 MiB | [postgresql-16-pg-durable_0.2.8-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.8-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-durable` | `0.2.8` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.6 MiB | [postgresql-16-pg-durable_0.2.8-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.8-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-durable` | `0.2.8` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.9 MiB | [postgresql-16-pg-durable_0.2.8-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.8-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-durable` | `0.2.8` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.6 MiB | [postgresql-16-pg-durable_0.2.8-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.8-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-durable` | `0.2.8` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.9 MiB | [postgresql-16-pg-durable_0.2.8-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.8-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-durable` | `0.2.8` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.6 MiB | [postgresql-16-pg-durable_0.2.8-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.8-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-durable` | `0.2.8` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.9 MiB | [postgresql-16-pg-durable_0.2.8-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.8-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_durable_15` | `0.2.8` | [el8.x86_64](/os/el8.x86_64) | pigsty | 5.1 MiB | [pg_durable_15-0.2.8-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_durable_15-0.2.8-1PGSTY.el8.x86_64.rpm) |
| `pg_durable_15` | `0.2.8` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.0 MiB | [pg_durable_15-0.2.8-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_durable_15-0.2.8-1PGSTY.el8.aarch64.rpm) |
| `pg_durable_15` | `0.2.8` | [el9.x86_64](/os/el9.x86_64) | pigsty | 5.0 MiB | [pg_durable_15-0.2.8-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_durable_15-0.2.8-1PGSTY.el9.x86_64.rpm) |
| `pg_durable_15` | `0.2.8` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.2 MiB | [pg_durable_15-0.2.8-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_durable_15-0.2.8-1PGSTY.el9.aarch64.rpm) |
| `pg_durable_15` | `0.2.8` | [el10.x86_64](/os/el10.x86_64) | pigsty | 5.0 MiB | [pg_durable_15-0.2.8-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_durable_15-0.2.8-1PGSTY.el10.x86_64.rpm) |
| `pg_durable_15` | `0.2.8` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.3 MiB | [pg_durable_15-0.2.8-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_durable_15-0.2.8-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-durable` | `0.2.8` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.2 MiB | [postgresql-15-pg-durable_0.2.8-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.8-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-durable` | `0.2.8` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.3 MiB | [postgresql-15-pg-durable_0.2.8-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.8-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-durable` | `0.2.8` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.2 MiB | [postgresql-15-pg-durable_0.2.8-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.8-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-durable` | `0.2.8` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.3 MiB | [postgresql-15-pg-durable_0.2.8-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.8-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-durable` | `0.2.8` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.6 MiB | [postgresql-15-pg-durable_0.2.8-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.8-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-durable` | `0.2.8` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.9 MiB | [postgresql-15-pg-durable_0.2.8-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.8-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-durable` | `0.2.8` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.6 MiB | [postgresql-15-pg-durable_0.2.8-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.8-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-durable` | `0.2.8` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.9 MiB | [postgresql-15-pg-durable_0.2.8-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.8-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-durable` | `0.2.8` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.6 MiB | [postgresql-15-pg-durable_0.2.8-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.8-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-durable` | `0.2.8` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.9 MiB | [postgresql-15-pg-durable_0.2.8-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.8-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_durable_14` | `0.2.8` | [el8.x86_64](/os/el8.x86_64) | pigsty | 5.1 MiB | [pg_durable_14-0.2.8-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_durable_14-0.2.8-1PGSTY.el8.x86_64.rpm) |
| `pg_durable_14` | `0.2.8` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.0 MiB | [pg_durable_14-0.2.8-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_durable_14-0.2.8-1PGSTY.el8.aarch64.rpm) |
| `pg_durable_14` | `0.2.8` | [el9.x86_64](/os/el9.x86_64) | pigsty | 5.0 MiB | [pg_durable_14-0.2.8-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_durable_14-0.2.8-1PGSTY.el9.x86_64.rpm) |
| `pg_durable_14` | `0.2.8` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.2 MiB | [pg_durable_14-0.2.8-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_durable_14-0.2.8-1PGSTY.el9.aarch64.rpm) |
| `pg_durable_14` | `0.2.8` | [el10.x86_64](/os/el10.x86_64) | pigsty | 5.0 MiB | [pg_durable_14-0.2.8-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_durable_14-0.2.8-1PGSTY.el10.x86_64.rpm) |
| `pg_durable_14` | `0.2.8` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.3 MiB | [pg_durable_14-0.2.8-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_durable_14-0.2.8-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-durable` | `0.2.8` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.2 MiB | [postgresql-14-pg-durable_0.2.8-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.8-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-durable` | `0.2.8` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.3 MiB | [postgresql-14-pg-durable_0.2.8-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.8-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-durable` | `0.2.8` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.2 MiB | [postgresql-14-pg-durable_0.2.8-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.8-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-durable` | `0.2.8` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.3 MiB | [postgresql-14-pg-durable_0.2.8-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.8-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-durable` | `0.2.8` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.6 MiB | [postgresql-14-pg-durable_0.2.8-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.8-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-durable` | `0.2.8` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.9 MiB | [postgresql-14-pg-durable_0.2.8-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.8-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-durable` | `0.2.8` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.6 MiB | [postgresql-14-pg-durable_0.2.8-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.8-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-durable` | `0.2.8` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.9 MiB | [postgresql-14-pg-durable_0.2.8-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.8-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-durable` | `0.2.8` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.6 MiB | [postgresql-14-pg-durable_0.2.8-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.8-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-durable` | `0.2.8` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.8 MiB | [postgresql-14-pg-durable_0.2.8-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.8-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/microsoft/pg_durable" title="Repository" icon="github" subtitle="github.com/microsoft/pg_durable" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_durable-0.2.8.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_durable;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_durable;		# install via package name, for the active PG version

pig install pg_durable -v 18;   # install for PG 18
pig install pg_durable -v 17;   # install for PG 17
pig install pg_durable -v 16;   # install for PG 16
pig install pg_durable -v 15;   # install for PG 15
pig install pg_durable -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_durable';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_durable;
```

## Usage

Sources:

- [PGXN 0.2.6 README](https://pgxn.org/dist/pg_durable/0.2.6/README.html)
- [0.2.6 user guide](https://api.pgxn.org/src/pg_durable/pg_durable-0.2.6/USER_GUIDE.md)
- [0.2.6 changelog](https://api.pgxn.org/src/pg_durable/pg_durable-0.2.6/CHANGELOG.md)
- [pg_durable control file](https://api.pgxn.org/src/pg_durable/pg_durable-0.2.6/pg_durable.control)
- [0.2.5 to 0.2.6 upgrade SQL](https://api.pgxn.org/src/pg_durable/pg_durable-0.2.6/sql/pg_durable--0.2.5--0.2.6.sql)

`pg_durable` runs durable, fault-tolerant SQL workflows inside PostgreSQL. A workflow is a graph of SQL steps, timers, signals, conditions, and parallel branches submitted with `df.start()`. Execution state is checkpointed in PostgreSQL so completed steps are not repeated after a crash, restart, or retry.

### Enable and Grant Access

Preload the worker, select its database and superuser role if the defaults are unsuitable, then restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_durable'
pg_durable.database = 'postgres'
pg_durable.worker_role = 'postgres'
```

Create the extension in `pg_durable.database` and grant an application login role access:

```sql
CREATE EXTENSION pg_durable;
SELECT df.grant_usage('app_role');
```

The worker role must be a superuser because it manages all users' instances while bypassing row-level security. The role that calls `df.start()` must have `LOGIN`, because workflow SQL is executed through a connection authenticated as that captured role.

### Build and Run a Workflow

```sql
SELECT df.start(
    'SELECT 100 AS amount' |=> 'total'
    ~> 'SELECT $total.amount * 2 AS doubled',
    'double-total'
);
```

`df.start()` returns an instance ID. Use it to monitor or control the run:

```sql
SELECT df.status('a1b2c3d4');
SELECT df.result('a1b2c3d4');
SELECT * FROM df.instance_nodes('a1b2c3d4');
SELECT * FROM df.instance_executions('a1b2c3d4', 20);
SELECT df.cancel('a1b2c3d4', 'No longer needed');
```

### DSL Index

- `~>` sequences steps; `|=>` names a result for `$name`, `$name.column`, or `$name.*` substitution.
- `&` / `df.join()` waits for parallel branches; `|` / `df.race()` keeps the first result.
- `?>` and `!>` / `df.if()` select conditional branches; `@>` / `df.loop()` repeats a graph.
- `df.sleep()`, `df.wait_for_schedule()`, and `df.wait_for_signal()` make waits durable.
- `df.signal()`, `df.wait_for_completion()`, `df.explain()`, and the instance-inspection functions operate on running or stored instances.
- `df.setvar()`, `df.getvar()`, `df.unsetvar()`, and `df.clearvars()` manage per-user variables captured when `df.start()` is called.

### Version 0.2.6 Boundaries

- Upstream source installation and published images support PostgreSQL 17 and 18 with `pgrx` 0.16.1. The extension still requires `shared_preload_libraries`, a restart, and a superuser worker role.
- Upgrades through 0.2.4 and 0.2.5 contain replay-breaking workflow changes. Drain or cancel in-flight JOIN, RACE, loop, and `df.wait_for_schedule()` work before upgrading; the 0.2.4 `df.nodes` key migration also takes an `ACCESS EXCLUSIVE` lock.
- `df.start(..., transaction_mode => 'new')` persists an independent start outside the caller transaction. Cluster-wide admission defaults to two concurrent starts and is controlled by `pg_durable.max_new_transaction_starts` and `pg_durable.new_transaction_start_timeout`.
- Variable substitution is resolved once from left to right in 0.2.6, so token-shaped text introduced by a value is not rescanned. It remains raw SQL substitution; never place untrusted input in `{name}` variables. Named step-result substitution through `$name` performs SQL escaping.
- The undocumented `df.ensure_durofut(text)` helper was removed. Drop or rewrite customer-owned dependent objects before upgrading.
- Re-run `df.grant_usage()` after `ALTER EXTENSION ... UPDATE`, because grants on all functions do not automatically include functions added later.
- `df.http()` and `df.http_multipart()` availability and egress policy are compile-time features. Their restrictions do not sandbox arbitrary SQL or other installed extensions.
- The project remains pre-1.0, and upstream's published Docker images are for evaluation and learning rather than production. Read every adjacent upgrade warning instead of assuming an untested multi-version jump is replay-safe.
