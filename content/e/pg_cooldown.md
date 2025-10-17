---
title: "pg_cooldown"
linkTitle: "pg_cooldown"
description: "remove buffered pages for specific relations"
weight: 5070
categories: ["Admin"]
width: full
---

remove buffered pages for specific relations

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5070** | {{< badge content="pg_cooldown" link="https://github.com/rbergm/pg_cooldown" >}} | {{< ext "pg_cooldown" "pg_cooldown" >}} | `0.1` | {{< category "ADMIN" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgfincore" >}} {{< ext "pgcozy" >}} {{< ext "pg_prewarm" >}} {{< ext "pg_buffercache" >}} {{< ext "pg_repack" >}} {{< ext "pg_squeeze" >}} {{< ext "system_stats" >}} {{< ext "pgmeminfo" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_cooldown" >}} | `0.1` | {{< badge content="18" color="red" alt="pg_cooldown_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_cooldown_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_cooldown" >}} | `0.1` | {{< badge content="18" color="red" alt="postgresql-18-pg-cooldown" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-cooldown` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_cooldown_18" >}}     | {{< pkg "pg_cooldown_17" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cooldown_17-0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_cooldown_16" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cooldown_16-0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_cooldown_15" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cooldown_15-0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_cooldown_14" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cooldown_14-0.1-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_cooldown_18" >}}     | {{< pkg "pg_cooldown_17" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cooldown_17-0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_cooldown_16" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cooldown_16-0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_cooldown_15" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cooldown_15-0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_cooldown_14" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cooldown_14-0.1-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_cooldown_18" >}}     | {{< pkg "pg_cooldown_17" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cooldown_17-0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_cooldown_16" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cooldown_16-0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_cooldown_15" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cooldown_15-0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_cooldown_14" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cooldown_14-0.1-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_cooldown_18" >}}     | {{< pkg "pg_cooldown_17" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cooldown_17-0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_cooldown_16" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cooldown_16-0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_cooldown_15" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cooldown_15-0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_cooldown_14" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cooldown_14-0.1-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-cooldown" >}}     | {{< pkg "postgresql-17-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-cooldown" >}}     | {{< pkg "postgresql-17-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-cooldown" >}}     | {{< pkg "postgresql-17-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-cooldown" >}}     | {{< pkg "postgresql-17-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-cooldown" >}}     | {{< pkg "postgresql-17-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-cooldown" >}}     | {{< pkg "postgresql-17-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-cooldown" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_cooldown_17` | 0.1 | `el8.x86_64` | pigsty | 16.4 KiB | [pg_cooldown_17-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cooldown_17-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_cooldown_17` | 0.1 | `el8.aarch64` | pigsty | 16.4 KiB | [pg_cooldown_17-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cooldown_17-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_cooldown_17` | 0.1 | `el9.aarch64` | pigsty | 16.2 KiB | [pg_cooldown_17-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cooldown_17-0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_cooldown_17` | 0.1 | `el9.x86_64` | pigsty | 16.4 KiB | [pg_cooldown_17-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cooldown_17-0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-cooldown` | 0.1 | `d12.x86_64` | pigsty | 12.9 KiB | [postgresql-17-pg-cooldown_0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-cooldown` | 0.1 | `d12.aarch64` | pigsty | 12.8 KiB | [postgresql-17-pg-cooldown_0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-cooldown` | 0.1 | `u22.x86_64` | pigsty | 13.1 KiB | [postgresql-17-pg-cooldown_0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-cooldown` | 0.1 | `u22.aarch64` | pigsty | 13.2 KiB | [postgresql-17-pg-cooldown_0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-cooldown` | 0.1 | `u24.x86_64` | pigsty | 12.7 KiB | [postgresql-17-pg-cooldown_0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-cooldown` | 0.1 | `u24.aarch64` | pigsty | 12.3 KiB | [postgresql-17-pg-cooldown_0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_cooldown_16` | 0.1 | `el8.x86_64` | pigsty | 16.4 KiB | [pg_cooldown_16-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cooldown_16-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_cooldown_16` | 0.1 | `el8.aarch64` | pigsty | 16.4 KiB | [pg_cooldown_16-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cooldown_16-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_cooldown_16` | 0.1 | `el9.x86_64` | pigsty | 16.4 KiB | [pg_cooldown_16-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cooldown_16-0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_cooldown_16` | 0.1 | `el9.aarch64` | pigsty | 16.3 KiB | [pg_cooldown_16-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cooldown_16-0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-cooldown` | 0.1 | `d12.x86_64` | pigsty | 12.9 KiB | [postgresql-16-pg-cooldown_0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-cooldown` | 0.1 | `d12.aarch64` | pigsty | 12.8 KiB | [postgresql-16-pg-cooldown_0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-cooldown` | 0.1 | `u22.aarch64` | pigsty | 13.2 KiB | [postgresql-16-pg-cooldown_0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-cooldown` | 0.1 | `u22.x86_64` | pigsty | 13.2 KiB | [postgresql-16-pg-cooldown_0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-cooldown` | 0.1 | `u24.x86_64` | pigsty | 12.7 KiB | [postgresql-16-pg-cooldown_0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-cooldown` | 0.1 | `u24.aarch64` | pigsty | 12.3 KiB | [postgresql-16-pg-cooldown_0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_cooldown_15` | 0.1 | `el8.x86_64` | pigsty | 16.4 KiB | [pg_cooldown_15-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cooldown_15-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_cooldown_15` | 0.1 | `el8.aarch64` | pigsty | 16.4 KiB | [pg_cooldown_15-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cooldown_15-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_cooldown_15` | 0.1 | `el9.x86_64` | pigsty | 16.4 KiB | [pg_cooldown_15-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cooldown_15-0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_cooldown_15` | 0.1 | `el9.aarch64` | pigsty | 16.3 KiB | [pg_cooldown_15-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cooldown_15-0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-cooldown` | 0.1 | `d12.aarch64` | pigsty | 12.8 KiB | [postgresql-15-pg-cooldown_0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-cooldown` | 0.1 | `d12.x86_64` | pigsty | 12.9 KiB | [postgresql-15-pg-cooldown_0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-cooldown` | 0.1 | `u22.aarch64` | pigsty | 13.2 KiB | [postgresql-15-pg-cooldown_0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-cooldown` | 0.1 | `u22.x86_64` | pigsty | 13.1 KiB | [postgresql-15-pg-cooldown_0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-cooldown` | 0.1 | `u24.x86_64` | pigsty | 12.7 KiB | [postgresql-15-pg-cooldown_0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-cooldown` | 0.1 | `u24.aarch64` | pigsty | 12.3 KiB | [postgresql-15-pg-cooldown_0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_cooldown_14` | 0.1 | `el8.x86_64` | pigsty | 16.4 KiB | [pg_cooldown_14-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cooldown_14-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_cooldown_14` | 0.1 | `el8.aarch64` | pigsty | 16.4 KiB | [pg_cooldown_14-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cooldown_14-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_cooldown_14` | 0.1 | `el9.x86_64` | pigsty | 16.4 KiB | [pg_cooldown_14-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cooldown_14-0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_cooldown_14` | 0.1 | `el9.aarch64` | pigsty | 16.3 KiB | [pg_cooldown_14-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cooldown_14-0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-cooldown` | 0.1 | `d12.x86_64` | pigsty | 12.9 KiB | [postgresql-14-pg-cooldown_0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-cooldown` | 0.1 | `d12.aarch64` | pigsty | 12.8 KiB | [postgresql-14-pg-cooldown_0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-cooldown` | 0.1 | `u22.x86_64` | pigsty | 13.1 KiB | [postgresql-14-pg-cooldown_0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-cooldown` | 0.1 | `u22.aarch64` | pigsty | 13.1 KiB | [postgresql-14-pg-cooldown_0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-cooldown` | 0.1 | `u24.x86_64` | pigsty | 12.7 KiB | [postgresql-14-pg-cooldown_0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-cooldown` | 0.1 | `u24.aarch64` | pigsty | 12.3 KiB | [postgresql-14-pg-cooldown_0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_cooldown_13` | 0.1 | `el8.aarch64` | pigsty | 16.4 KiB | [pg_cooldown_13-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cooldown_13-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_cooldown_13` | 0.1 | `el8.x86_64` | pigsty | 16.3 KiB | [pg_cooldown_13-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cooldown_13-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_cooldown_13` | 0.1 | `el9.aarch64` | pigsty | 16.3 KiB | [pg_cooldown_13-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cooldown_13-0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_cooldown_13` | 0.1 | `el9.x86_64` | pigsty | 16.4 KiB | [pg_cooldown_13-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cooldown_13-0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-cooldown` | 0.1 | `d12.aarch64` | pigsty | 12.8 KiB | [postgresql-13-pg-cooldown_0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-13-pg-cooldown_0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-cooldown` | 0.1 | `d12.x86_64` | pigsty | 12.7 KiB | [postgresql-13-pg-cooldown_0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-13-pg-cooldown_0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-cooldown` | 0.1 | `u22.aarch64` | pigsty | 13.1 KiB | [postgresql-13-pg-cooldown_0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-13-pg-cooldown_0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-cooldown` | 0.1 | `u22.x86_64` | pigsty | 13.1 KiB | [postgresql-13-pg-cooldown_0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-13-pg-cooldown_0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-cooldown` | 0.1 | `u24.aarch64` | pigsty | 12.3 KiB | [postgresql-13-pg-cooldown_0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-13-pg-cooldown_0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-cooldown` | 0.1 | `u24.x86_64` | pigsty | 12.5 KiB | [postgresql-13-pg-cooldown_0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-13-pg-cooldown_0.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rbergm/pg_cooldown" title="Repository" icon="github" subtitle="github.com/rbergm/pg_cooldown" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_cooldown-0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_cooldown; # get pg_cooldown source code
pig build dep pg_cooldown; # install build dependencies
pig build pkg pg_cooldown; # build extension rpm or deb
pig build ext pg_cooldown; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_cooldown; # install by extension name, for the current active PG version
pig ext install pg_cooldown; # install via package alias, for the active PG version
pig ext install pg_cooldown -v 18;   # install for PG 18
pig ext install pg_cooldown -v 17;   # install for PG 17
pig ext install pg_cooldown -v 16;   # install for PG 16
pig ext install pg_cooldown -v 15;   # install for PG 15
pig ext install pg_cooldown -v 14;   # install for PG 14
pig ext install pg_cooldown -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_cooldown;
```

