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
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_durable_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-durable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-durable : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_durable_18` | `0.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 5.3 MiB | [pg_durable_18-0.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_durable_18-0.2.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_durable_18` | `0.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.9 MiB | [pg_durable_18-0.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_durable_18-0.2.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_durable_18` | `0.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 5.1 MiB | [pg_durable_18-0.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_durable_18-0.2.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_durable_18` | `0.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 5.0 MiB | [pg_durable_18-0.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_durable_18-0.2.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_durable_18` | `0.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 5.1 MiB | [pg_durable_18-0.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_durable_18-0.2.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_durable_18` | `0.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 5.0 MiB | [pg_durable_18-0.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_durable_18-0.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-durable` | `0.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.1 MiB | [postgresql-18-pg-durable_0.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-durable` | `0.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.6 MiB | [postgresql-18-pg-durable_0.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-durable` | `0.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.1 MiB | [postgresql-18-pg-durable_0.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-durable` | `0.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.6 MiB | [postgresql-18-pg-durable_0.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-durable` | `0.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.5 MiB | [postgresql-18-pg-durable_0.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-durable` | `0.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.2 MiB | [postgresql-18-pg-durable_0.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-durable` | `0.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.5 MiB | [postgresql-18-pg-durable_0.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-durable` | `0.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.2 MiB | [postgresql-18-pg-durable_0.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-durable` | `0.2.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.5 MiB | [postgresql-18-pg-durable_0.2.2-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.2-2PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-durable` | `0.2.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.8 MiB | [postgresql-18-pg-durable_0.2.2-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-18-pg-durable_0.2.2-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_durable_17` | `0.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 5.3 MiB | [pg_durable_17-0.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_durable_17-0.2.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_durable_17` | `0.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.9 MiB | [pg_durable_17-0.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_durable_17-0.2.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_durable_17` | `0.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 5.1 MiB | [pg_durable_17-0.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_durable_17-0.2.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_durable_17` | `0.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 5.0 MiB | [pg_durable_17-0.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_durable_17-0.2.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_durable_17` | `0.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 5.1 MiB | [pg_durable_17-0.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_durable_17-0.2.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_durable_17` | `0.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 5.1 MiB | [pg_durable_17-0.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_durable_17-0.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-durable` | `0.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.1 MiB | [postgresql-17-pg-durable_0.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-durable` | `0.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.6 MiB | [postgresql-17-pg-durable_0.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-durable` | `0.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.1 MiB | [postgresql-17-pg-durable_0.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-durable` | `0.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.6 MiB | [postgresql-17-pg-durable_0.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-durable` | `0.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.5 MiB | [postgresql-17-pg-durable_0.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-durable` | `0.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.2 MiB | [postgresql-17-pg-durable_0.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-durable` | `0.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.5 MiB | [postgresql-17-pg-durable_0.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-durable` | `0.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.2 MiB | [postgresql-17-pg-durable_0.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-durable` | `0.2.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.5 MiB | [postgresql-17-pg-durable_0.2.2-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.2-2PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-durable` | `0.2.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.8 MiB | [postgresql-17-pg-durable_0.2.2-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-17-pg-durable_0.2.2-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_durable_16` | `0.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 5.3 MiB | [pg_durable_16-0.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_durable_16-0.2.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_durable_16` | `0.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.9 MiB | [pg_durable_16-0.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_durable_16-0.2.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_durable_16` | `0.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 5.1 MiB | [pg_durable_16-0.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_durable_16-0.2.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_durable_16` | `0.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 5.0 MiB | [pg_durable_16-0.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_durable_16-0.2.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_durable_16` | `0.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 5.1 MiB | [pg_durable_16-0.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_durable_16-0.2.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_durable_16` | `0.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 5.0 MiB | [pg_durable_16-0.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_durable_16-0.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-durable` | `0.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.1 MiB | [postgresql-16-pg-durable_0.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-durable` | `0.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.6 MiB | [postgresql-16-pg-durable_0.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-durable` | `0.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.1 MiB | [postgresql-16-pg-durable_0.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-durable` | `0.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.6 MiB | [postgresql-16-pg-durable_0.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-durable` | `0.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.5 MiB | [postgresql-16-pg-durable_0.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-durable` | `0.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.2 MiB | [postgresql-16-pg-durable_0.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-durable` | `0.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.5 MiB | [postgresql-16-pg-durable_0.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-durable` | `0.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.2 MiB | [postgresql-16-pg-durable_0.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-durable` | `0.2.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.5 MiB | [postgresql-16-pg-durable_0.2.2-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.2-2PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-durable` | `0.2.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.8 MiB | [postgresql-16-pg-durable_0.2.2-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-16-pg-durable_0.2.2-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_durable_15` | `0.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 5.3 MiB | [pg_durable_15-0.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_durable_15-0.2.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_durable_15` | `0.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.9 MiB | [pg_durable_15-0.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_durable_15-0.2.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_durable_15` | `0.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 5.1 MiB | [pg_durable_15-0.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_durable_15-0.2.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_durable_15` | `0.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 5.0 MiB | [pg_durable_15-0.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_durable_15-0.2.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_durable_15` | `0.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 5.1 MiB | [pg_durable_15-0.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_durable_15-0.2.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_durable_15` | `0.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 5.0 MiB | [pg_durable_15-0.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_durable_15-0.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-durable` | `0.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.1 MiB | [postgresql-15-pg-durable_0.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-durable` | `0.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.6 MiB | [postgresql-15-pg-durable_0.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-durable` | `0.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.1 MiB | [postgresql-15-pg-durable_0.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-durable` | `0.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.6 MiB | [postgresql-15-pg-durable_0.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-durable` | `0.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.5 MiB | [postgresql-15-pg-durable_0.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-durable` | `0.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.2 MiB | [postgresql-15-pg-durable_0.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-durable` | `0.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.5 MiB | [postgresql-15-pg-durable_0.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-durable` | `0.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.2 MiB | [postgresql-15-pg-durable_0.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-durable` | `0.2.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.5 MiB | [postgresql-15-pg-durable_0.2.2-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.2-2PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-durable` | `0.2.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.8 MiB | [postgresql-15-pg-durable_0.2.2-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-15-pg-durable_0.2.2-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_durable_14` | `0.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 5.3 MiB | [pg_durable_14-0.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_durable_14-0.2.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_durable_14` | `0.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.9 MiB | [pg_durable_14-0.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_durable_14-0.2.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_durable_14` | `0.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 5.1 MiB | [pg_durable_14-0.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_durable_14-0.2.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_durable_14` | `0.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 5.0 MiB | [pg_durable_14-0.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_durable_14-0.2.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_durable_14` | `0.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 5.1 MiB | [pg_durable_14-0.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_durable_14-0.2.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_durable_14` | `0.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 5.0 MiB | [pg_durable_14-0.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_durable_14-0.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-durable` | `0.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.1 MiB | [postgresql-14-pg-durable_0.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-durable` | `0.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.6 MiB | [postgresql-14-pg-durable_0.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-durable` | `0.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.1 MiB | [postgresql-14-pg-durable_0.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-durable` | `0.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.6 MiB | [postgresql-14-pg-durable_0.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-durable` | `0.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.5 MiB | [postgresql-14-pg-durable_0.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-durable` | `0.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.2 MiB | [postgresql-14-pg-durable_0.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-durable` | `0.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.5 MiB | [postgresql-14-pg-durable_0.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-durable` | `0.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.2 MiB | [postgresql-14-pg-durable_0.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-durable` | `0.2.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.5 MiB | [postgresql-14-pg-durable_0.2.2-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.2-2PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-durable` | `0.2.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.8 MiB | [postgresql-14-pg-durable_0.2.2-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-durable/postgresql-14-pg-durable_0.2.2-2PIGSTY~resolute_arm64.deb) |

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

- [pg_durable README on main](https://github.com/microsoft/pg_durable/blob/main/README.md)
- [User Guide](https://github.com/microsoft/pg_durable/blob/main/USER_GUIDE.md)
- [control file](https://github.com/microsoft/pg_durable/blob/main/pg_durable.control)
- [0.2.3 to 0.2.4 upgrade SQL](https://github.com/microsoft/pg_durable/blob/main/sql/pg_durable--0.2.3--0.2.4.sql)
- [0.2.2 to 0.2.3 upgrade SQL](https://github.com/microsoft/pg_durable/blob/v0.2.3/sql/pg_durable--0.2.2--0.2.3.sql)

`pg_durable` runs durable, fault-tolerant SQL workflows inside PostgreSQL. A workflow is a graph of SQL strings, functions, timers, signals, conditions, and parallel branches submitted with `df.start()`. The extension checkpoints state in PostgreSQL so completed steps are not re-executed after crashes, restarts, or failed steps.

The current main branch reports package version `0.2.4`; the latest tagged release observed for this refresh is `v0.2.3`. Upstream published packages and Docker images focus on PostgreSQL 17 and 18, while local Pigsty metadata may validate a broader PostgreSQL range separately.

### Enable and Grant Access

```sql
CREATE EXTENSION pg_durable;

SELECT df.grant_usage('app_role');
```

`pg_durable` must be loaded through `shared_preload_libraries`, followed by a PostgreSQL restart. The background worker initializes asynchronously after `CREATE EXTENSION`; retry if `df.*` calls report that the worker is not initialized yet. `CREATE EXTENSION` does not grant usage to `PUBLIC`, so run `df.grant_usage()` for application roles and repeat it after extension upgrades so new functions are covered.

The worker connects to `pg_durable.database` and runs as `pg_durable.worker_role`. The worker role must be a superuser because it bypasses RLS to manage all users' durable instances.

### Start and Monitor Workflows

```sql
SELECT df.start('SELECT ''Hello, durable world!''', 'hello-job');

SELECT *
FROM df.list_instances();

SELECT df.status('a1b2c3d4');
SELECT df.result('a1b2c3d4');
SELECT df.cancel('a1b2c3d4', 'No longer needed');
```

`df.start()` returns an instance ID. Status, result, cancel, signal, explain, and await helpers operate on the instance ID, not on the label. `df.list_instances()` has a basic overload and a paginated overload; pass at least three arguments, using `NULL` for skipped filters, when you need `created_at`, `completed_at`, and `next_cursor`.

### Compose SQL Steps

```sql
-- Run one step, name its result, then substitute it in the next step.
SELECT df.start(
  'SELECT 100 AS amount' |=> 'total'
  ~> 'SELECT $total * 2 AS doubled',
  'double-total'
);

-- Pass a multi-row result set between steps.
SELECT df.start(
  'SELECT id FROM documents WHERE processed = false LIMIT 100' |=> 'batch'
  ~> 'UPDATE documents SET processed = true WHERE id IN (SELECT id FROM $batch.*)',
  'process-documents'
);

-- Branch on a SQL condition.
SELECT df.start(
  'SELECT count(*) > 10 FROM orders'
    ?> 'SELECT ''high volume'''
    !> 'SELECT ''low volume''',
  'order-volume'
);
```

Core operators are `~>` for sequence, `|=>` for naming a result, `&` for join, `|` for race, `?>` and `!>` for conditional branches, and `@>` for loops. Result substitution supports `$name`, `$name.column`, null-safe `$name?` / `$name.column?`, and row-set expansion through `$name.*`.

### Timers, Schedules, Signals, and Variables

```sql
SELECT df.start(
  @> (
    df.wait_for_schedule('0 * * * *')
    ~> 'SELECT run_hourly_rollup()'
  ),
  'hourly-rollup'
);

SELECT df.start(
  'SELECT create_invoice()' |=> 'invoice'
  ~> df.wait_for_signal('approval', 86400)
  ~> 'SELECT finalize_invoice($invoice.id)',
  'invoice-approval'
);

SELECT df.setvar('api_url', 'https://example.internal');
SELECT df.getvar('api_url');
```

Useful DSL functions include `df.sleep(seconds)`, `df.wait_for_schedule(cron)`, `df.wait_for_signal(name, timeout)`, `df.signal(id, name, data)`, `df.join()`, `df.join3()`, `df.race()`, `df.if()`, `df.if_rows()`, `df.loop()`, `df.break()`, `df.await_instance()`, `df.explain()`, and durable variable helpers.

### Monitoring and Runtime State

```sql
SELECT * FROM df.instance_nodes('a1b2c3d4');
SELECT * FROM df.instance_executions('a1b2c3d4', 20);

-- Requires a direct admin grant; df.grant_usage() does not include it.
SELECT * FROM df.metrics();
```

`df.instance_nodes()` reports both stored node status and inferred status. The inferred view adds `skipped` for untaken branches or race losers and reinterprets superseded loop-iteration nodes as pending when a newer iteration has taken over. This is a read-time interpretation and does not change workflow execution semantics.

### Configuration and Caveats

- Required preload: add `pg_durable` to `shared_preload_libraries` and restart PostgreSQL.
- `pg_durable.database` must name the database where the extension is created; workflows are not processed in a different database unless you explicitly use the supported database argument where available.
- Fresh installs after the v0.2.3 provider-schema move use `_duroxide`; installs upgraded from <= 0.2.2 keep the legacy `duroxide` schema. `df.duroxide_schema()` reports which schema an install should use.
- `pg_durable.worker_role` must exist and be a superuser.
- Connection-related GUCs include `pg_durable.max_management_connections`, `pg_durable.max_duroxide_connections`, `pg_durable.max_user_connections`, and `pg_durable.execution_acquire_timeout`.
- `df.http()` performs outbound HTTP work. Standard grants exclude HTTP unless `df.grant_usage(..., include_http => true)` is used, and release builds may restrict HTTP egress by feature.
- Upstream status is preview. The published Docker image is intended for evaluation and learning, not production.
