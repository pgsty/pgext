---
title: "pg_savior"
linkTitle: "pg_savior"
description: "Postgres extension to save OOPS mistakes"
weight: 5810
categories: ["Admin"]
width: full
---

Postgres extension to save OOPS mistakes

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5810** | {{< badge content="pg_savior" link="https://github.com/viggy28/pg_savior" >}} | {{< ext "pg_savior" "pg_savior" >}} | `0.0.1` | {{< category "ADMIN" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_upless" >}} {{< ext "safeupdate" >}} {{< ext "pg_drop_events" >}} {{< ext "pg_cheat_funcs" >}} {{< ext "table_log" >}} {{< ext "pg_snakeoil" >}} {{< ext "pg_auditor" >}} {{< ext "temporal_tables" >}} |

> [!Note] -tuplestore_donestoring , breaks on pg18 @ el


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_savior" >}} | `0.0.1` | {{< badge content="18" color="red" alt="pg_savior_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_savior_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_savior" >}} | `0.0.1` | {{< badge content="18" color="red" alt="postgresql-18-pg-savior" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-savior` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_savior_18" >}}     | {{< pkg "pg_savior_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_17-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_savior_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_16-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_savior_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_15-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_savior_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_14-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_savior_18" >}}     | {{< pkg "pg_savior_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_17-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_savior_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_16-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_savior_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_15-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_savior_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_14-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_savior_18" >}}     | {{< pkg "pg_savior_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_17-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_savior_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_16-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_savior_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_15-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_savior_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_14-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_savior_18" >}}     | {{< pkg "pg_savior_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_17-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_savior_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_16-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_savior_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_15-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_savior_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_14-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-savior" >}}     | {{< pkg "postgresql-17-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-savior" >}}     | {{< pkg "postgresql-17-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-savior" >}}     | {{< pkg "postgresql-17-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-savior" >}}     | {{< pkg "postgresql-17-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-savior" >}}     | {{< pkg "postgresql-17-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-savior" >}}     | {{< pkg "postgresql-17-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-savior" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_savior_17` | 0.0.1 | `el8.x86_64` | pigsty | 13.5 KiB | [pg_savior_17-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_17-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_savior_17` | 0.0.1 | `el8.aarch64` | pigsty | 13.6 KiB | [pg_savior_17-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_17-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_savior_17` | 0.0.1 | `el9.aarch64` | pigsty | 13.5 KiB | [pg_savior_17-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_17-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_savior_17` | 0.0.1 | `el9.x86_64` | pigsty | 13.6 KiB | [pg_savior_17-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_17-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-savior` | 0.0.1 | `d12.x86_64` | pigsty | 19.5 KiB | [postgresql-17-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-savior` | 0.0.1 | `d12.aarch64` | pigsty | 19.4 KiB | [postgresql-17-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-savior` | 0.0.1 | `u22.x86_64` | pigsty | 20.3 KiB | [postgresql-17-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-savior` | 0.0.1 | `u22.aarch64` | pigsty | 20.0 KiB | [postgresql-17-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-savior` | 0.0.1 | `u24.x86_64` | pigsty | 16.8 KiB | [postgresql-17-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-savior` | 0.0.1 | `u24.aarch64` | pigsty | 16.3 KiB | [postgresql-17-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_savior_16` | 0.0.1 | `el8.x86_64` | pigsty | 13.5 KiB | [pg_savior_16-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_16-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_savior_16` | 0.0.1 | `el8.aarch64` | pigsty | 13.6 KiB | [pg_savior_16-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_16-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_savior_16` | 0.0.1 | `el9.x86_64` | pigsty | 13.7 KiB | [pg_savior_16-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_16-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_savior_16` | 0.0.1 | `el9.aarch64` | pigsty | 13.5 KiB | [pg_savior_16-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_16-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-savior` | 0.0.1 | `d12.x86_64` | pigsty | 19.0 KiB | [postgresql-16-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-savior` | 0.0.1 | `d12.aarch64` | pigsty | 18.9 KiB | [postgresql-16-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-savior` | 0.0.1 | `u22.aarch64` | pigsty | 19.5 KiB | [postgresql-16-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-savior` | 0.0.1 | `u22.x86_64` | pigsty | 19.8 KiB | [postgresql-16-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-savior` | 0.0.1 | `u24.x86_64` | pigsty | 16.8 KiB | [postgresql-16-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-savior` | 0.0.1 | `u24.aarch64` | pigsty | 16.3 KiB | [postgresql-16-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_savior_15` | 0.0.1 | `el8.x86_64` | pigsty | 13.5 KiB | [pg_savior_15-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_15-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_savior_15` | 0.0.1 | `el8.aarch64` | pigsty | 13.6 KiB | [pg_savior_15-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_15-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_savior_15` | 0.0.1 | `el9.x86_64` | pigsty | 13.7 KiB | [pg_savior_15-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_15-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_savior_15` | 0.0.1 | `el9.aarch64` | pigsty | 13.6 KiB | [pg_savior_15-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_15-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-savior` | 0.0.1 | `d12.aarch64` | pigsty | 18.9 KiB | [postgresql-15-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-savior` | 0.0.1 | `d12.x86_64` | pigsty | 19.1 KiB | [postgresql-15-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-savior` | 0.0.1 | `u22.aarch64` | pigsty | 19.5 KiB | [postgresql-15-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-savior` | 0.0.1 | `u22.x86_64` | pigsty | 19.8 KiB | [postgresql-15-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-savior` | 0.0.1 | `u24.x86_64` | pigsty | 16.8 KiB | [postgresql-15-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-savior` | 0.0.1 | `u24.aarch64` | pigsty | 16.4 KiB | [postgresql-15-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_savior_14` | 0.0.1 | `el8.x86_64` | pigsty | 13.5 KiB | [pg_savior_14-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_14-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_savior_14` | 0.0.1 | `el8.aarch64` | pigsty | 13.6 KiB | [pg_savior_14-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_14-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_savior_14` | 0.0.1 | `el9.x86_64` | pigsty | 13.7 KiB | [pg_savior_14-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_14-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_savior_14` | 0.0.1 | `el9.aarch64` | pigsty | 13.6 KiB | [pg_savior_14-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_14-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-savior` | 0.0.1 | `d12.x86_64` | pigsty | 17.9 KiB | [postgresql-14-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-savior` | 0.0.1 | `d12.aarch64` | pigsty | 17.7 KiB | [postgresql-14-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-savior` | 0.0.1 | `u22.x86_64` | pigsty | 18.6 KiB | [postgresql-14-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-savior` | 0.0.1 | `u22.aarch64` | pigsty | 18.3 KiB | [postgresql-14-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-savior` | 0.0.1 | `u24.x86_64` | pigsty | 16.8 KiB | [postgresql-14-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-savior` | 0.0.1 | `u24.aarch64` | pigsty | 16.3 KiB | [postgresql-14-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_savior_13` | 0.0.1 | `el8.aarch64` | pigsty | 13.6 KiB | [pg_savior_13-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_13-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_savior_13` | 0.0.1 | `el8.x86_64` | pigsty | 13.5 KiB | [pg_savior_13-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_13-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_savior_13` | 0.0.1 | `el9.aarch64` | pigsty | 13.6 KiB | [pg_savior_13-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_13-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_savior_13` | 0.0.1 | `el9.x86_64` | pigsty | 13.7 KiB | [pg_savior_13-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_13-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-savior` | 0.0.1 | `d12.aarch64` | pigsty | 17.4 KiB | [postgresql-13-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-13-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-savior` | 0.0.1 | `d12.x86_64` | pigsty | 17.6 KiB | [postgresql-13-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-13-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-savior` | 0.0.1 | `u22.aarch64` | pigsty | 18.2 KiB | [postgresql-13-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-13-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-savior` | 0.0.1 | `u22.x86_64` | pigsty | 18.5 KiB | [postgresql-13-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-13-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-savior` | 0.0.1 | `u24.aarch64` | pigsty | 16.1 KiB | [postgresql-13-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-13-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-savior` | 0.0.1 | `u24.x86_64` | pigsty | 16.6 KiB | [postgresql-13-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-13-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/viggy28/pg_savior" title="Repository" icon="github" subtitle="github.com/viggy28/pg_savior" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_savior-0.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_savior; # get pg_savior source code
pig build dep pg_savior; # install build dependencies
pig build pkg pg_savior; # build extension rpm or deb
pig build ext pg_savior; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_savior; # install by extension name, for the current active PG version
pig ext install pg_savior; # install via package alias, for the active PG version
pig ext install pg_savior -v 17;   # install for PG 17
pig ext install pg_savior -v 16;   # install for PG 16
pig ext install pg_savior -v 15;   # install for PG 15
pig ext install pg_savior -v 14;   # install for PG 14
pig ext install pg_savior -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_savior;
```

