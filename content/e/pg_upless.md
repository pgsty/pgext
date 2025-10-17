---
title: "pg_upless"
linkTitle: "pg_upless"
description: "Detect Useless UPDATE"
weight: 5180
categories: ["Admin"]
width: full
---

Detect Useless UPDATE

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5180** | {{< badge content="pg_upless" link="https://github.com/rodo/pg_upless" >}} | {{< ext "pg_upless" "pg_upless" >}} | `0.0.3` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="-----dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "pg_readonly" >}} {{< ext "pg_savior" >}} {{< ext "safeupdate" >}} {{< ext "pg_permissions" >}} {{< ext "pgaudit" >}} {{< ext "set_user" >}} {{< ext "pg_drop_events" >}} {{< ext "table_log" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_upless" >}} | `0.0.3` | {{< badge content="18" color="red" alt="pg_upless_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_upless_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_upless" >}} | `0.0.3` | {{< badge content="18" color="red" alt="postgresql-18-pg-upless" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-upless` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_upless_18" >}}     | {{< pkg "pg_upless_17" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_upless_17-0.0.3-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_upless_16" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_upless_16-0.0.3-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_upless_15" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_upless_15-0.0.3-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_upless_14" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_upless_14-0.0.3-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_upless_18" >}}     | {{< pkg "pg_upless_17" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_upless_17-0.0.3-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_upless_16" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_upless_16-0.0.3-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_upless_15" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_upless_15-0.0.3-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_upless_14" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_upless_14-0.0.3-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_upless_18" >}}     | {{< pkg "pg_upless_17" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_upless_17-0.0.3-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_upless_16" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_upless_16-0.0.3-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_upless_15" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_upless_15-0.0.3-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_upless_14" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_upless_14-0.0.3-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_upless_18" >}}     | {{< pkg "pg_upless_17" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_upless_17-0.0.3-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_upless_16" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_upless_16-0.0.3-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_upless_15" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_upless_15-0.0.3-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_upless_14" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_upless_14-0.0.3-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-upless" >}}     | {{< pkg "postgresql-17-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-upless" >}}     | {{< pkg "postgresql-17-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-upless" >}}     | {{< pkg "postgresql-17-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-upless" >}}     | {{< pkg "postgresql-17-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-upless" >}}     | {{< pkg "postgresql-17-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-upless" >}}     | {{< pkg "postgresql-17-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-upless" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_upless_17` | 0.0.3 | `el8.x86_64` | pigsty | 11.6 KiB | [pg_upless_17-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_upless_17-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_upless_17` | 0.0.3 | `el8.aarch64` | pigsty | 11.5 KiB | [pg_upless_17-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_upless_17-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_upless_17` | 0.0.3 | `el9.aarch64` | pigsty | 11.5 KiB | [pg_upless_17-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_upless_17-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_upless_17` | 0.0.3 | `el9.x86_64` | pigsty | 11.6 KiB | [pg_upless_17-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_upless_17-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-upless` | 0.0.3 | `d12.x86_64` | pigsty | 5.1 KiB | [postgresql-17-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-upless` | 0.0.3 | `d12.aarch64` | pigsty | 5.1 KiB | [postgresql-17-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-upless` | 0.0.3 | `u22.x86_64` | pigsty | 4.9 KiB | [postgresql-17-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-upless` | 0.0.3 | `u22.aarch64` | pigsty | 4.9 KiB | [postgresql-17-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-upless` | 0.0.3 | `u24.x86_64` | pigsty | 4.9 KiB | [postgresql-17-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-upless` | 0.0.3 | `u24.aarch64` | pigsty | 4.9 KiB | [postgresql-17-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_upless_16` | 0.0.3 | `el8.x86_64` | pigsty | 11.6 KiB | [pg_upless_16-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_upless_16-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_upless_16` | 0.0.3 | `el8.aarch64` | pigsty | 11.5 KiB | [pg_upless_16-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_upless_16-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_upless_16` | 0.0.3 | `el9.x86_64` | pigsty | 11.6 KiB | [pg_upless_16-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_upless_16-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_upless_16` | 0.0.3 | `el9.aarch64` | pigsty | 11.5 KiB | [pg_upless_16-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_upless_16-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-upless` | 0.0.3 | `d12.x86_64` | pigsty | 5.1 KiB | [postgresql-16-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-upless` | 0.0.3 | `d12.aarch64` | pigsty | 5.1 KiB | [postgresql-16-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-upless` | 0.0.3 | `u22.aarch64` | pigsty | 4.9 KiB | [postgresql-16-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-upless` | 0.0.3 | `u22.x86_64` | pigsty | 4.8 KiB | [postgresql-16-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-upless` | 0.0.3 | `u24.x86_64` | pigsty | 4.9 KiB | [postgresql-16-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-upless` | 0.0.3 | `u24.aarch64` | pigsty | 4.9 KiB | [postgresql-16-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_upless_15` | 0.0.3 | `el8.x86_64` | pigsty | 11.6 KiB | [pg_upless_15-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_upless_15-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_upless_15` | 0.0.3 | `el8.aarch64` | pigsty | 11.5 KiB | [pg_upless_15-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_upless_15-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_upless_15` | 0.0.3 | `el9.x86_64` | pigsty | 11.6 KiB | [pg_upless_15-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_upless_15-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_upless_15` | 0.0.3 | `el9.aarch64` | pigsty | 11.5 KiB | [pg_upless_15-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_upless_15-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-upless` | 0.0.3 | `d12.aarch64` | pigsty | 5.1 KiB | [postgresql-15-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-upless` | 0.0.3 | `d12.x86_64` | pigsty | 5.1 KiB | [postgresql-15-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-upless` | 0.0.3 | `u22.aarch64` | pigsty | 4.9 KiB | [postgresql-15-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-upless` | 0.0.3 | `u22.x86_64` | pigsty | 4.9 KiB | [postgresql-15-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-upless` | 0.0.3 | `u24.x86_64` | pigsty | 4.9 KiB | [postgresql-15-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-upless` | 0.0.3 | `u24.aarch64` | pigsty | 4.9 KiB | [postgresql-15-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_upless_14` | 0.0.3 | `el8.x86_64` | pigsty | 11.6 KiB | [pg_upless_14-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_upless_14-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_upless_14` | 0.0.3 | `el8.aarch64` | pigsty | 11.5 KiB | [pg_upless_14-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_upless_14-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_upless_14` | 0.0.3 | `el9.x86_64` | pigsty | 11.6 KiB | [pg_upless_14-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_upless_14-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_upless_14` | 0.0.3 | `el9.aarch64` | pigsty | 11.5 KiB | [pg_upless_14-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_upless_14-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-upless` | 0.0.3 | `d12.x86_64` | pigsty | 5.1 KiB | [postgresql-14-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-upless` | 0.0.3 | `d12.aarch64` | pigsty | 5.1 KiB | [postgresql-14-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-upless` | 0.0.3 | `u22.x86_64` | pigsty | 4.9 KiB | [postgresql-14-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-upless` | 0.0.3 | `u22.aarch64` | pigsty | 4.9 KiB | [postgresql-14-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-upless` | 0.0.3 | `u24.x86_64` | pigsty | 4.9 KiB | [postgresql-14-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-upless` | 0.0.3 | `u24.aarch64` | pigsty | 4.9 KiB | [postgresql-14-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_upless_13` | 0.0.3 | `el8.aarch64` | pigsty | 11.5 KiB | [pg_upless_13-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_upless_13-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_upless_13` | 0.0.3 | `el8.x86_64` | pigsty | 11.6 KiB | [pg_upless_13-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_upless_13-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_upless_13` | 0.0.3 | `el9.aarch64` | pigsty | 11.5 KiB | [pg_upless_13-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_upless_13-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_upless_13` | 0.0.3 | `el9.x86_64` | pigsty | 11.6 KiB | [pg_upless_13-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_upless_13-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-upless` | 0.0.3 | `d12.aarch64` | pigsty | 5.1 KiB | [postgresql-13-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-13-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-upless` | 0.0.3 | `d12.x86_64` | pigsty | 5.1 KiB | [postgresql-13-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-13-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-upless` | 0.0.3 | `u22.aarch64` | pigsty | 4.9 KiB | [postgresql-13-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-13-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-upless` | 0.0.3 | `u22.x86_64` | pigsty | 4.9 KiB | [postgresql-13-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-13-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-upless` | 0.0.3 | `u24.aarch64` | pigsty | 4.9 KiB | [postgresql-13-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-13-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-upless` | 0.0.3 | `u24.x86_64` | pigsty | 4.9 KiB | [postgresql-13-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-13-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rodo/pg_upless" title="Repository" icon="github" subtitle="github.com/rodo/pg_upless" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_upless-0.0.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_upless; # get pg_upless source code
pig build dep pg_upless; # install build dependencies
pig build pkg pg_upless; # build extension rpm or deb
pig build ext pg_upless; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_upless; # install by extension name, for the current active PG version
pig ext install pg_upless; # install via package alias, for the active PG version
pig ext install pg_upless -v 18;   # install for PG 18
pig ext install pg_upless -v 17;   # install for PG 17
pig ext install pg_upless -v 16;   # install for PG 16
pig ext install pg_upless -v 15;   # install for PG 15
pig ext install pg_upless -v 14;   # install for PG 14
pig ext install pg_upless -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_upless;
```

