---
title: "pg_cheat_funcs"
linkTitle: "pg_cheat_funcs"
description: "Provides cheat (but useful) functions"
weight: 5220
categories: ["Admin"]
width: full
---

Provides cheat (but useful) functions

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5220** | {{< badge content="pg_cheat_funcs" link="https://github.com/MasaoFujii/pg_cheat_funcs" >}} | {{< ext "pg_cheat_funcs" "pg_cheat_funcs" >}} | `1.0` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_crash" >}} {{< ext "pg_snakeoil" >}} {{< ext "pg_dirtyread" >}} {{< ext "pg_savior" >}} {{< ext "pg_surgery" >}} {{< ext "adminpack" >}} {{< ext "pageinspect" >}} {{< ext "pg_repack" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_cheat_funcs" >}} | `1.0` | {{< badge content="18" color="red" alt="pg_cheat_funcs_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_cheat_funcs_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_cheat_funcs" >}} | `1.0` | {{< badge content="18" color="red" alt="postgresql-18-pg-cheat-funcs" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-cheat-funcs` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_cheat_funcs_18" >}}     | {{< pkg "pg_cheat_funcs_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cheat_funcs_17-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_cheat_funcs_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cheat_funcs_16-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_cheat_funcs_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cheat_funcs_15-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_cheat_funcs_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cheat_funcs_14-1.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_cheat_funcs_18" >}}     | {{< pkg "pg_cheat_funcs_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cheat_funcs_17-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_cheat_funcs_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cheat_funcs_16-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_cheat_funcs_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cheat_funcs_15-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_cheat_funcs_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cheat_funcs_14-1.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_cheat_funcs_18" >}}     | {{< pkg "pg_cheat_funcs_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cheat_funcs_17-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_cheat_funcs_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cheat_funcs_16-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_cheat_funcs_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cheat_funcs_15-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_cheat_funcs_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cheat_funcs_14-1.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_cheat_funcs_18" >}}     | {{< pkg "pg_cheat_funcs_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cheat_funcs_17-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_cheat_funcs_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cheat_funcs_16-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_cheat_funcs_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cheat_funcs_15-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_cheat_funcs_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cheat_funcs_14-1.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-cheat-funcs" >}}     | {{< pkg "postgresql-17-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-cheat-funcs" >}}     | {{< pkg "postgresql-17-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-cheat-funcs" >}}     | {{< pkg "postgresql-17-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-cheat-funcs" >}}     | {{< pkg "postgresql-17-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-cheat-funcs" >}}     | {{< pkg "postgresql-17-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-cheat-funcs" >}}     | {{< pkg "postgresql-17-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-cheat-funcs" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_cheat_funcs_17` | 1.0 | `el8.x86_64` | pigsty | 46.7 KiB | [pg_cheat_funcs_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cheat_funcs_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cheat_funcs_17` | 1.0 | `el8.aarch64` | pigsty | 46.3 KiB | [pg_cheat_funcs_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cheat_funcs_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cheat_funcs_17` | 1.0 | `el9.aarch64` | pigsty | 52.2 KiB | [pg_cheat_funcs_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cheat_funcs_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_cheat_funcs_17` | 1.0 | `el9.x86_64` | pigsty | 52.2 KiB | [pg_cheat_funcs_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cheat_funcs_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-cheat-funcs` | 1.0 | `d12.x86_64` | pigsty | 99.6 KiB | [postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-cheat-funcs` | 1.0 | `d12.aarch64` | pigsty | 98.7 KiB | [postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-cheat-funcs` | 1.0 | `u22.x86_64` | pigsty | 124.5 KiB | [postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-cheat-funcs` | 1.0 | `u22.aarch64` | pigsty | 123.8 KiB | [postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-cheat-funcs` | 1.0 | `u24.x86_64` | pigsty | 111.8 KiB | [postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-cheat-funcs` | 1.0 | `u24.aarch64` | pigsty | 111.2 KiB | [postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_cheat_funcs_16` | 1.0 | `el8.x86_64` | pigsty | 46.8 KiB | [pg_cheat_funcs_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cheat_funcs_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cheat_funcs_16` | 1.0 | `el8.aarch64` | pigsty | 46.3 KiB | [pg_cheat_funcs_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cheat_funcs_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cheat_funcs_16` | 1.0 | `el9.x86_64` | pigsty | 52.2 KiB | [pg_cheat_funcs_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cheat_funcs_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_cheat_funcs_16` | 1.0 | `el9.aarch64` | pigsty | 52.1 KiB | [pg_cheat_funcs_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cheat_funcs_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-cheat-funcs` | 1.0 | `d12.x86_64` | pigsty | 98.9 KiB | [postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-cheat-funcs` | 1.0 | `d12.aarch64` | pigsty | 98.1 KiB | [postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-cheat-funcs` | 1.0 | `u22.aarch64` | pigsty | 123.0 KiB | [postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-cheat-funcs` | 1.0 | `u22.x86_64` | pigsty | 123.8 KiB | [postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-cheat-funcs` | 1.0 | `u24.x86_64` | pigsty | 111.8 KiB | [postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-cheat-funcs` | 1.0 | `u24.aarch64` | pigsty | 111.2 KiB | [postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_cheat_funcs_15` | 1.0 | `el8.x86_64` | pigsty | 46.9 KiB | [pg_cheat_funcs_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cheat_funcs_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cheat_funcs_15` | 1.0 | `el8.aarch64` | pigsty | 46.5 KiB | [pg_cheat_funcs_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cheat_funcs_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cheat_funcs_15` | 1.0 | `el9.x86_64` | pigsty | 52.4 KiB | [pg_cheat_funcs_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cheat_funcs_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_cheat_funcs_15` | 1.0 | `el9.aarch64` | pigsty | 52.3 KiB | [pg_cheat_funcs_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cheat_funcs_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-cheat-funcs` | 1.0 | `d12.aarch64` | pigsty | 98.7 KiB | [postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-cheat-funcs` | 1.0 | `d12.x86_64` | pigsty | 99.6 KiB | [postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-cheat-funcs` | 1.0 | `u22.aarch64` | pigsty | 124.1 KiB | [postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-cheat-funcs` | 1.0 | `u22.x86_64` | pigsty | 124.4 KiB | [postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-cheat-funcs` | 1.0 | `u24.x86_64` | pigsty | 112.6 KiB | [postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-cheat-funcs` | 1.0 | `u24.aarch64` | pigsty | 111.8 KiB | [postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_cheat_funcs_14` | 1.0 | `el8.x86_64` | pigsty | 46.8 KiB | [pg_cheat_funcs_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cheat_funcs_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cheat_funcs_14` | 1.0 | `el8.aarch64` | pigsty | 46.4 KiB | [pg_cheat_funcs_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cheat_funcs_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cheat_funcs_14` | 1.0 | `el9.x86_64` | pigsty | 52.1 KiB | [pg_cheat_funcs_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cheat_funcs_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_cheat_funcs_14` | 1.0 | `el9.aarch64` | pigsty | 52.2 KiB | [pg_cheat_funcs_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cheat_funcs_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-cheat-funcs` | 1.0 | `d12.x86_64` | pigsty | 98.4 KiB | [postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-cheat-funcs` | 1.0 | `d12.aarch64` | pigsty | 97.7 KiB | [postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-cheat-funcs` | 1.0 | `u22.x86_64` | pigsty | 124.7 KiB | [postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-cheat-funcs` | 1.0 | `u22.aarch64` | pigsty | 123.7 KiB | [postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-cheat-funcs` | 1.0 | `u24.x86_64` | pigsty | 112.3 KiB | [postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-cheat-funcs` | 1.0 | `u24.aarch64` | pigsty | 111.5 KiB | [postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_cheat_funcs_13` | 1.0 | `el8.aarch64` | pigsty | 46.9 KiB | [pg_cheat_funcs_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cheat_funcs_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cheat_funcs_13` | 1.0 | `el8.x86_64` | pigsty | 47.1 KiB | [pg_cheat_funcs_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cheat_funcs_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cheat_funcs_13` | 1.0 | `el9.aarch64` | pigsty | 52.8 KiB | [pg_cheat_funcs_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cheat_funcs_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_cheat_funcs_13` | 1.0 | `el9.x86_64` | pigsty | 52.9 KiB | [pg_cheat_funcs_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cheat_funcs_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-cheat-funcs` | 1.0 | `d12.aarch64` | pigsty | 100.3 KiB | [postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-cheat-funcs` | 1.0 | `d12.x86_64` | pigsty | 101.0 KiB | [postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-cheat-funcs` | 1.0 | `u22.aarch64` | pigsty | 125.8 KiB | [postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-cheat-funcs` | 1.0 | `u22.x86_64` | pigsty | 126.1 KiB | [postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-cheat-funcs` | 1.0 | `u24.aarch64` | pigsty | 113.5 KiB | [postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-cheat-funcs` | 1.0 | `u24.x86_64` | pigsty | 114.1 KiB | [postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/MasaoFujii/pg_cheat_funcs" title="Repository" icon="github" subtitle="github.com/MasaoFujii/pg_cheat_funcs" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_cheat_funcs-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_cheat_funcs; # get pg_cheat_funcs source code
pig build dep pg_cheat_funcs; # install build dependencies
pig build pkg pg_cheat_funcs; # build extension rpm or deb
pig build ext pg_cheat_funcs; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_cheat_funcs; # install by extension name, for the current active PG version
pig ext install pg_cheat_funcs; # install via package alias, for the active PG version
pig ext install pg_cheat_funcs -v 18;   # install for PG 18
pig ext install pg_cheat_funcs -v 17;   # install for PG 17
pig ext install pg_cheat_funcs -v 16;   # install for PG 16
pig ext install pg_cheat_funcs -v 15;   # install for PG 15
pig ext install pg_cheat_funcs -v 14;   # install for PG 14
pig ext install pg_cheat_funcs -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_cheat_funcs;
```

