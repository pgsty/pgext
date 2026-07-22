---
title: "pg_durable"
linkTitle: "pg_durable"
description: "Durable SQL functions for PostgreSQL"
weight: 2870
categories: ["FEAT"]
width: full
---

[**pg_durable**](https://github.com/microsoft/pg_durable) : Durable SQL functions for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2870** | {{< badge content="pg_durable" link="https://github.com/microsoft/pg_durable" >}} | {{< ext "pg_durable" >}} | `0.2.3` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `df` `duroxide` |

> [!Note] Requires shared_preload_libraries=pg_durable and a superuser worker role.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_durable` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.3` | {{< bg "18" "pg_durable_18" "green" >}} {{< bg "17" "pg_durable_17" "green" >}} {{< bg "16" "pg_durable_16" "green" >}} {{< bg "15" "pg_durable_15" "green" >}} {{< bg "14" "pg_durable_14" "green" >}} | `pg_durable_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.3` | {{< bg "18" "postgresql-18-pg-durable" "green" >}} {{< bg "17" "postgresql-17-pg-durable" "green" >}} {{< bg "16" "postgresql-16-pg-durable" "green" >}} {{< bg "15" "postgresql-15-pg-durable" "green" >}} {{< bg "14" "postgresql-14-pg-durable" "green" >}} | `postgresql-$v-pg-durable` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_durable_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_durable_18` | `0.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 4.7 MiB | [pg_durable_18-0.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_durable_18-0.2.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_durable_18` | `0.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.3 MiB | [pg_durable_18-0.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_durable_18-0.2.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_durable_18` | `0.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 4.5 MiB | [pg_durable_18-0.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_durable_18-0.2.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_durable_18` | `0.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.4 MiB | [pg_durable_18-0.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_durable_18-0.2.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_durable_18` | `0.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 4.5 MiB | [pg_durable_18-0.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_durable_18-0.2.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_durable_18` | `0.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.5 MiB | [pg_durable_18-0.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_durable_18-0.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-durable` | `0.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.7 MiB | [postgresql-18-pg-durable_0.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-durable` | `0.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.2 MiB | [postgresql-18-pg-durable_0.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-durable` | `0.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.7 MiB | [postgresql-18-pg-durable_0.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-durable` | `0.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.2 MiB | [postgresql-18-pg-durable_0.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-durable` | `0.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.0 MiB | [postgresql-18-pg-durable_0.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-durable` | `0.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.7 MiB | [postgresql-18-pg-durable_0.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-durable` | `0.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.0 MiB | [postgresql-18-pg-durable_0.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-durable` | `0.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.8 MiB | [postgresql-18-pg-durable_0.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-durable` | `0.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.0 MiB | [postgresql-18-pg-durable_0.2.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-durable` | `0.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.7 MiB | [postgresql-18-pg-durable_0.2.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_durable_17` | `0.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 4.7 MiB | [pg_durable_17-0.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_durable_17-0.2.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_durable_17` | `0.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.3 MiB | [pg_durable_17-0.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_durable_17-0.2.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_durable_17` | `0.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 4.5 MiB | [pg_durable_17-0.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_durable_17-0.2.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_durable_17` | `0.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.4 MiB | [pg_durable_17-0.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_durable_17-0.2.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_durable_17` | `0.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 4.5 MiB | [pg_durable_17-0.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_durable_17-0.2.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_durable_17` | `0.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.4 MiB | [pg_durable_17-0.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_durable_17-0.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-durable` | `0.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.7 MiB | [postgresql-17-pg-durable_0.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-durable` | `0.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.2 MiB | [postgresql-17-pg-durable_0.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-durable` | `0.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.7 MiB | [postgresql-17-pg-durable_0.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-durable` | `0.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.2 MiB | [postgresql-17-pg-durable_0.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-durable` | `0.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.0 MiB | [postgresql-17-pg-durable_0.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-durable` | `0.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.8 MiB | [postgresql-17-pg-durable_0.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-durable` | `0.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.0 MiB | [postgresql-17-pg-durable_0.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-durable` | `0.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.8 MiB | [postgresql-17-pg-durable_0.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-durable` | `0.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.0 MiB | [postgresql-17-pg-durable_0.2.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-durable` | `0.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.7 MiB | [postgresql-17-pg-durable_0.2.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_durable_16` | `0.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 4.7 MiB | [pg_durable_16-0.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_durable_16-0.2.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_durable_16` | `0.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.3 MiB | [pg_durable_16-0.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_durable_16-0.2.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_durable_16` | `0.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 4.5 MiB | [pg_durable_16-0.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_durable_16-0.2.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_durable_16` | `0.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.4 MiB | [pg_durable_16-0.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_durable_16-0.2.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_durable_16` | `0.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 4.5 MiB | [pg_durable_16-0.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_durable_16-0.2.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_durable_16` | `0.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.5 MiB | [pg_durable_16-0.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_durable_16-0.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-durable` | `0.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.7 MiB | [postgresql-16-pg-durable_0.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-durable` | `0.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.2 MiB | [postgresql-16-pg-durable_0.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-durable` | `0.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.7 MiB | [postgresql-16-pg-durable_0.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-durable` | `0.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.2 MiB | [postgresql-16-pg-durable_0.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-durable` | `0.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.0 MiB | [postgresql-16-pg-durable_0.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-durable` | `0.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.8 MiB | [postgresql-16-pg-durable_0.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-durable` | `0.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.0 MiB | [postgresql-16-pg-durable_0.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-durable` | `0.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.8 MiB | [postgresql-16-pg-durable_0.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-durable` | `0.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.0 MiB | [postgresql-16-pg-durable_0.2.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-durable` | `0.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.7 MiB | [postgresql-16-pg-durable_0.2.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_durable_15` | `0.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 4.7 MiB | [pg_durable_15-0.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_durable_15-0.2.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_durable_15` | `0.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.3 MiB | [pg_durable_15-0.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_durable_15-0.2.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_durable_15` | `0.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 4.5 MiB | [pg_durable_15-0.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_durable_15-0.2.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_durable_15` | `0.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.4 MiB | [pg_durable_15-0.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_durable_15-0.2.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_durable_15` | `0.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 4.5 MiB | [pg_durable_15-0.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_durable_15-0.2.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_durable_15` | `0.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.4 MiB | [pg_durable_15-0.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_durable_15-0.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-durable` | `0.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.7 MiB | [postgresql-15-pg-durable_0.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-durable` | `0.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.2 MiB | [postgresql-15-pg-durable_0.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-durable` | `0.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.7 MiB | [postgresql-15-pg-durable_0.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-durable` | `0.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.2 MiB | [postgresql-15-pg-durable_0.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-durable` | `0.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.0 MiB | [postgresql-15-pg-durable_0.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-durable` | `0.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.7 MiB | [postgresql-15-pg-durable_0.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-durable` | `0.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.0 MiB | [postgresql-15-pg-durable_0.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-durable` | `0.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.8 MiB | [postgresql-15-pg-durable_0.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-durable` | `0.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.0 MiB | [postgresql-15-pg-durable_0.2.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-durable` | `0.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.7 MiB | [postgresql-15-pg-durable_0.2.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_durable_14` | `0.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 4.7 MiB | [pg_durable_14-0.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_durable_14-0.2.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_durable_14` | `0.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.3 MiB | [pg_durable_14-0.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_durable_14-0.2.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_durable_14` | `0.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 4.5 MiB | [pg_durable_14-0.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_durable_14-0.2.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_durable_14` | `0.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.5 MiB | [pg_durable_14-0.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_durable_14-0.2.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_durable_14` | `0.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 4.5 MiB | [pg_durable_14-0.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_durable_14-0.2.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_durable_14` | `0.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.4 MiB | [pg_durable_14-0.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_durable_14-0.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-durable` | `0.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.6 MiB | [postgresql-14-pg-durable_0.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-durable` | `0.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.2 MiB | [postgresql-14-pg-durable_0.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-durable` | `0.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.6 MiB | [postgresql-14-pg-durable_0.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-durable` | `0.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.2 MiB | [postgresql-14-pg-durable_0.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-durable` | `0.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.0 MiB | [postgresql-14-pg-durable_0.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-durable` | `0.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.7 MiB | [postgresql-14-pg-durable_0.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-durable` | `0.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.0 MiB | [postgresql-14-pg-durable_0.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-durable` | `0.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.8 MiB | [postgresql-14-pg-durable_0.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-durable` | `0.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.0 MiB | [postgresql-14-pg-durable_0.2.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-durable` | `0.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.7 MiB | [postgresql-14-pg-durable_0.2.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/microsoft/pg_durable" title="Repository" icon="github" subtitle="github.com/microsoft/pg_durable" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_durable-0.2.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_durable;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

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

- [Official v0.2.3 README](https://github.com/microsoft/pg_durable/blob/v0.2.3/README.md)
- [v0.2.3 user guide](https://github.com/microsoft/pg_durable/blob/v0.2.3/USER_GUIDE.md)
- [v0.2.3 release notes](https://github.com/microsoft/pg_durable/releases/tag/v0.2.3)
- [v0.2.2 to v0.2.3 upgrade SQL](https://github.com/microsoft/pg_durable/blob/v0.2.3/sql/pg_durable--0.2.2--0.2.3.sql)

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

### Version 0.2.3 Boundaries

- Fresh v0.2.3 installs place provider objects in `_duroxide`; installations upgraded from 0.2.2 or earlier keep `duroxide`. `df.duroxide_schema()` reports the active schema.
- Graphs deeper than 256 levels or larger than 10,000 nodes are rejected. A condition query returning no rows evaluates as false.
- Re-run `df.grant_usage()` after `ALTER EXTENSION ... UPDATE`, because grants on all functions do not automatically include functions added later.
- Variable `{name}` substitution is raw SQL text substitution; never place untrusted input in such variables. Named step-result substitution through `$name` performs SQL escaping.
- `df.http()` availability and egress policy are compile-time features. Its restrictions do not sandbox arbitrary SQL or other installed extensions.
- Upstream labels the project preview, and the published v0.2.3 Docker images are for evaluation and learning rather than production.
