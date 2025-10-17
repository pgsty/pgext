---
title: "pg_sqlog"
linkTitle: "pg_sqlog"
description: "Provide SQL interface to logs"
weight: 6330
categories: ["Stat"]
width: full
---

Provide SQL interface to logs

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6330** | {{< badge content="pg_sqlog" link="https://github.com/kouber/pg_sqlog" >}} | {{< ext "pg_sqlog" "pg_sqlog" >}} | `1.6` | {{< category "STAT" >}} | {{< license "BSD 3-Clause" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "file_fdw" >}} |
|   **See Also**    | {{< ext "pg_profile" >}} {{< ext "pg_tracing" >}} {{< ext "pg_show_plans" >}} {{< ext "pg_stat_kcache" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_store_plans" >}} |

> [!Note] require certain params


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_sqlog" >}} | `1.6` | {{< badge content="18" color="red" alt="pg_sqlog_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_sqlog_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_sqlog" >}} | `1.6` | {{< badge content="18" color="red" alt="postgresql-18-pg-sqlog" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-sqlog` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_sqlog_18" >}}     | {{< pkg "pg_sqlog_17" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_sqlog_17-1.6-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_sqlog_16" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_sqlog_16-1.6-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_sqlog_15" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_sqlog_15-1.6-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_sqlog_14" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_sqlog_14-1.6-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_sqlog_18" >}}     | {{< pkg "pg_sqlog_17" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_sqlog_17-1.6-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_sqlog_16" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_sqlog_16-1.6-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_sqlog_15" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_sqlog_15-1.6-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_sqlog_14" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_sqlog_14-1.6-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_sqlog_18" >}}     | {{< pkg "pg_sqlog_17" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_sqlog_17-1.6-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_sqlog_16" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_sqlog_16-1.6-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_sqlog_15" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_sqlog_15-1.6-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_sqlog_14" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_sqlog_14-1.6-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_sqlog_18" >}}     | {{< pkg "pg_sqlog_17" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_sqlog_17-1.6-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_sqlog_16" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_sqlog_16-1.6-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_sqlog_15" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_sqlog_15-1.6-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_sqlog_14" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_sqlog_14-1.6-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-sqlog" >}}     | {{< pkg "postgresql-17-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-17-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-16-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-15-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-14-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-sqlog" >}}     | {{< pkg "postgresql-17-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-17-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-16-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-15-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-14-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-sqlog" >}}     | {{< pkg "postgresql-17-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-17-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-16-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-15-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-14-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-sqlog" >}}     | {{< pkg "postgresql-17-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-17-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-16-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-15-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-14-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-sqlog" >}}     | {{< pkg "postgresql-17-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-17-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-16-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-15-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-14-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-sqlog" >}}     | {{< pkg "postgresql-17-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-17-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-16-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-15-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-sqlog" "1.6" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-14-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_sqlog_17` | 1.6 | `el8.x86_64` | pigsty | 15.0 KiB | [pg_sqlog_17-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_sqlog_17-1.6-1PIGSTY.el8.x86_64.rpm) |
| `pg_sqlog_17` | 1.6 | `el8.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_17-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_sqlog_17-1.6-1PIGSTY.el8.aarch64.rpm) |
| `pg_sqlog_17` | 1.6 | `el9.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_17-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_sqlog_17-1.6-1PIGSTY.el9.aarch64.rpm) |
| `pg_sqlog_17` | 1.6 | `el9.x86_64` | pigsty | 15.0 KiB | [pg_sqlog_17-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_sqlog_17-1.6-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-sqlog` | 1.6 | `d12.x86_64` | pigsty | 9.9 KiB | [postgresql-17-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-17-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-sqlog` | 1.6 | `d12.aarch64` | pigsty | 9.9 KiB | [postgresql-17-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-17-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-sqlog` | 1.6 | `u22.x86_64` | pigsty | 9.5 KiB | [postgresql-17-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-17-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-sqlog` | 1.6 | `u22.aarch64` | pigsty | 9.5 KiB | [postgresql-17-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-17-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-sqlog` | 1.6 | `u24.x86_64` | pigsty | 9.5 KiB | [postgresql-17-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-17-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-sqlog` | 1.6 | `u24.aarch64` | pigsty | 9.5 KiB | [postgresql-17-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-17-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_sqlog_16` | 1.6 | `el8.x86_64` | pigsty | 15.0 KiB | [pg_sqlog_16-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_sqlog_16-1.6-1PIGSTY.el8.x86_64.rpm) |
| `pg_sqlog_16` | 1.6 | `el8.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_16-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_sqlog_16-1.6-1PIGSTY.el8.aarch64.rpm) |
| `pg_sqlog_16` | 1.6 | `el9.x86_64` | pigsty | 15.0 KiB | [pg_sqlog_16-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_sqlog_16-1.6-1PIGSTY.el9.x86_64.rpm) |
| `pg_sqlog_16` | 1.6 | `el9.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_16-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_sqlog_16-1.6-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-sqlog` | 1.6 | `d12.x86_64` | pigsty | 9.9 KiB | [postgresql-16-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-16-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-sqlog` | 1.6 | `d12.aarch64` | pigsty | 9.9 KiB | [postgresql-16-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-16-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-sqlog` | 1.6 | `u22.aarch64` | pigsty | 9.5 KiB | [postgresql-16-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-16-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-sqlog` | 1.6 | `u22.x86_64` | pigsty | 9.5 KiB | [postgresql-16-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-16-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-sqlog` | 1.6 | `u24.x86_64` | pigsty | 9.5 KiB | [postgresql-16-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-16-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-sqlog` | 1.6 | `u24.aarch64` | pigsty | 9.5 KiB | [postgresql-16-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-16-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_sqlog_15` | 1.6 | `el8.x86_64` | pigsty | 15.0 KiB | [pg_sqlog_15-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_sqlog_15-1.6-1PIGSTY.el8.x86_64.rpm) |
| `pg_sqlog_15` | 1.6 | `el8.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_15-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_sqlog_15-1.6-1PIGSTY.el8.aarch64.rpm) |
| `pg_sqlog_15` | 1.6 | `el9.x86_64` | pigsty | 15.0 KiB | [pg_sqlog_15-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_sqlog_15-1.6-1PIGSTY.el9.x86_64.rpm) |
| `pg_sqlog_15` | 1.6 | `el9.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_15-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_sqlog_15-1.6-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-sqlog` | 1.6 | `d12.aarch64` | pigsty | 9.9 KiB | [postgresql-15-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-15-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-sqlog` | 1.6 | `d12.x86_64` | pigsty | 9.9 KiB | [postgresql-15-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-15-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-sqlog` | 1.6 | `u22.aarch64` | pigsty | 9.5 KiB | [postgresql-15-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-15-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-sqlog` | 1.6 | `u22.x86_64` | pigsty | 9.5 KiB | [postgresql-15-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-15-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-sqlog` | 1.6 | `u24.x86_64` | pigsty | 9.5 KiB | [postgresql-15-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-15-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-sqlog` | 1.6 | `u24.aarch64` | pigsty | 9.5 KiB | [postgresql-15-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-15-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_sqlog_14` | 1.6 | `el8.x86_64` | pigsty | 15.0 KiB | [pg_sqlog_14-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_sqlog_14-1.6-1PIGSTY.el8.x86_64.rpm) |
| `pg_sqlog_14` | 1.6 | `el8.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_14-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_sqlog_14-1.6-1PIGSTY.el8.aarch64.rpm) |
| `pg_sqlog_14` | 1.6 | `el9.x86_64` | pigsty | 15.0 KiB | [pg_sqlog_14-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_sqlog_14-1.6-1PIGSTY.el9.x86_64.rpm) |
| `pg_sqlog_14` | 1.6 | `el9.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_14-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_sqlog_14-1.6-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-sqlog` | 1.6 | `d12.x86_64` | pigsty | 9.9 KiB | [postgresql-14-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-14-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-sqlog` | 1.6 | `d12.aarch64` | pigsty | 9.9 KiB | [postgresql-14-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-14-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-sqlog` | 1.6 | `u22.x86_64` | pigsty | 9.5 KiB | [postgresql-14-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-14-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-sqlog` | 1.6 | `u22.aarch64` | pigsty | 9.5 KiB | [postgresql-14-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-14-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-sqlog` | 1.6 | `u24.x86_64` | pigsty | 9.5 KiB | [postgresql-14-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-14-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-sqlog` | 1.6 | `u24.aarch64` | pigsty | 9.5 KiB | [postgresql-14-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-14-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_sqlog_13` | 1.6 | `el8.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_13-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_sqlog_13-1.6-1PIGSTY.el8.aarch64.rpm) |
| `pg_sqlog_13` | 1.6 | `el8.x86_64` | pigsty | 15.0 KiB | [pg_sqlog_13-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_sqlog_13-1.6-1PIGSTY.el8.x86_64.rpm) |
| `pg_sqlog_13` | 1.6 | `el9.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_13-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_sqlog_13-1.6-1PIGSTY.el9.aarch64.rpm) |
| `pg_sqlog_13` | 1.6 | `el9.x86_64` | pigsty | 15.0 KiB | [pg_sqlog_13-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_sqlog_13-1.6-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-sqlog` | 1.6 | `d12.aarch64` | pigsty | 9.9 KiB | [postgresql-13-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-13-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-sqlog` | 1.6 | `d12.x86_64` | pigsty | 9.9 KiB | [postgresql-13-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-13-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-sqlog` | 1.6 | `u22.aarch64` | pigsty | 9.5 KiB | [postgresql-13-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-13-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-sqlog` | 1.6 | `u22.x86_64` | pigsty | 9.5 KiB | [postgresql-13-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-13-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-sqlog` | 1.6 | `u24.aarch64` | pigsty | 9.5 KiB | [postgresql-13-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-13-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-sqlog` | 1.6 | `u24.x86_64` | pigsty | 9.5 KiB | [postgresql-13-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-13-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/kouber/pg_sqlog" title="Repository" icon="github" subtitle="github.com/kouber/pg_sqlog" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_sqlog-1.6.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_sqlog; # get pg_sqlog source code
pig build dep pg_sqlog; # install build dependencies
pig build pkg pg_sqlog; # build extension rpm or deb
pig build ext pg_sqlog; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_sqlog; # install by extension name, for the current active PG version
pig ext install pg_sqlog; # install via package alias, for the active PG version
pig ext install pg_sqlog -v 18;   # install for PG 18
pig ext install pg_sqlog -v 17;   # install for PG 17
pig ext install pg_sqlog -v 16;   # install for PG 16
pig ext install pg_sqlog -v 15;   # install for PG 15
pig ext install pg_sqlog -v 14;   # install for PG 14
pig ext install pg_sqlog -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_sqlog CASCADE SCHEMA sqlog;
```

