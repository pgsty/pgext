---
title: "pg_sqlog"
linkTitle: "pg_sqlog"
description: "Provide SQL interface to logs"
weight: 6330
categories: ["STAT"]
width: full
---

Provide SQL interface to logs


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6330** | {{< badge content="pg_sqlog" link="https://github.com/kouber/pg_sqlog" >}} | {{< ext "pg_sqlog" >}} | `1.6` | {{< category "STAT" >}} | {{< license "BSD 3-Clause" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "file_fdw" >}} |
|   **See Also**    | {{< ext "pg_profile" >}} {{< ext "pg_tracing" >}} {{< ext "pg_show_plans" >}} {{< ext "pg_stat_kcache" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_store_plans" >}} |

> [!Note] require certain params


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_sqlog" >}} | `1.6` | {{< bg "18" "pg_sqlog_18" "green" >}} {{< bg "17" "pg_sqlog_17" "green" >}} {{< bg "16" "pg_sqlog_16" "green" >}} {{< bg "15" "pg_sqlog_15" "green" >}} {{< bg "14" "pg_sqlog_14" "green" >}} {{< bg "13" "pg_sqlog_13" "green" >}} | `pg_sqlog_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_sqlog" >}} | `1.6` | {{< bg "18" "postgresql-18-pg-sqlog" "green" >}} {{< bg "17" "postgresql-17-pg-sqlog" "green" >}} {{< bg "16" "postgresql-16-pg-sqlog" "green" >}} {{< bg "15" "postgresql-15-pg-sqlog" "green" >}} {{< bg "14" "postgresql-14-pg-sqlog" "green" >}} {{< bg "13" "postgresql-13-pg-sqlog" "green" >}} | `postgresql-$v-pg-sqlog` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 1.6" "pg_sqlog_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 1.6" "pg_sqlog_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 1.6" "pg_sqlog_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 1.6" "pg_sqlog_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    | {{< bg "PIGSTY 1.6" "pg_sqlog_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_13 : AVAIL 1" "green" >}} |
|    `el10.aarch64`    | {{< bg "PIGSTY 1.6" "pg_sqlog_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "pg_sqlog_13 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-sqlog : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.6" "postgresql-17-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-16-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-15-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-14-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-13-pg-sqlog : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-sqlog : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.6" "postgresql-17-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-16-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-15-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-14-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-13-pg-sqlog : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-sqlog : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-sqlog : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-sqlog : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-sqlog : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-sqlog : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-sqlog : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-sqlog : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-sqlog : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-sqlog : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-sqlog : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-sqlog : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-sqlog : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-sqlog : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.6" "postgresql-17-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-16-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-15-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-14-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-13-pg-sqlog : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-sqlog : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.6" "postgresql-17-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-16-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-15-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-14-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-13-pg-sqlog : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-sqlog : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.6" "postgresql-17-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-16-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-15-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-14-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-13-pg-sqlog : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-sqlog : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.6" "postgresql-17-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-16-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-15-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-14-pg-sqlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "postgresql-13-pg-sqlog : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_sqlog_18` | 1.6 | `el8.x86_64` | pigsty | 15.1 KiB | [pg_sqlog_18-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_sqlog_18-1.6-1PIGSTY.el8.x86_64.rpm) |
| `pg_sqlog_18` | 1.6 | `el8.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_18-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_sqlog_18-1.6-1PIGSTY.el8.aarch64.rpm) |
| `pg_sqlog_18` | 1.6 | `el9.x86_64` | pigsty | 15.0 KiB | [pg_sqlog_18-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_sqlog_18-1.6-1PIGSTY.el9.x86_64.rpm) |
| `pg_sqlog_18` | 1.6 | `el9.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_18-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_sqlog_18-1.6-1PIGSTY.el9.aarch64.rpm) |
| `pg_sqlog_18` | 1.6 | `el10.x86_64` | pigsty | 15.2 KiB | [pg_sqlog_18-1.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_sqlog_18-1.6-1PIGSTY.el10.x86_64.rpm) |
| `pg_sqlog_18` | 1.6 | `el10.aarch64` | pigsty | 15.2 KiB | [pg_sqlog_18-1.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_sqlog_18-1.6-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_sqlog_17` | 1.6 | `el8.x86_64` | pigsty | 15.1 KiB | [pg_sqlog_17-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_sqlog_17-1.6-1PIGSTY.el8.x86_64.rpm) |
| `pg_sqlog_17` | 1.6 | `el8.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_17-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_sqlog_17-1.6-1PIGSTY.el8.aarch64.rpm) |
| `pg_sqlog_17` | 1.6 | `el9.x86_64` | pigsty | 15.0 KiB | [pg_sqlog_17-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_sqlog_17-1.6-1PIGSTY.el9.x86_64.rpm) |
| `pg_sqlog_17` | 1.6 | `el9.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_17-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_sqlog_17-1.6-1PIGSTY.el9.aarch64.rpm) |
| `pg_sqlog_17` | 1.6 | `el10.x86_64` | pigsty | 15.2 KiB | [pg_sqlog_17-1.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_sqlog_17-1.6-1PIGSTY.el10.x86_64.rpm) |
| `pg_sqlog_17` | 1.6 | `el10.aarch64` | pigsty | 15.2 KiB | [pg_sqlog_17-1.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_sqlog_17-1.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-sqlog` | 1.6 | `d12.x86_64` | pigsty | 9.9 KiB | [postgresql-17-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-17-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-sqlog` | 1.6 | `d12.aarch64` | pigsty | 9.9 KiB | [postgresql-17-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-17-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-sqlog` | 1.6 | `u22.x86_64` | pigsty | 9.5 KiB | [postgresql-17-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-17-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-sqlog` | 1.6 | `u22.aarch64` | pigsty | 9.5 KiB | [postgresql-17-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-17-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-sqlog` | 1.6 | `u24.x86_64` | pigsty | 9.5 KiB | [postgresql-17-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-17-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-sqlog` | 1.6 | `u24.aarch64` | pigsty | 9.5 KiB | [postgresql-17-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-17-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_sqlog_16` | 1.6 | `el8.x86_64` | pigsty | 15.1 KiB | [pg_sqlog_16-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_sqlog_16-1.6-1PIGSTY.el8.x86_64.rpm) |
| `pg_sqlog_16` | 1.6 | `el8.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_16-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_sqlog_16-1.6-1PIGSTY.el8.aarch64.rpm) |
| `pg_sqlog_16` | 1.6 | `el9.x86_64` | pigsty | 15.0 KiB | [pg_sqlog_16-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_sqlog_16-1.6-1PIGSTY.el9.x86_64.rpm) |
| `pg_sqlog_16` | 1.6 | `el9.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_16-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_sqlog_16-1.6-1PIGSTY.el9.aarch64.rpm) |
| `pg_sqlog_16` | 1.6 | `el10.x86_64` | pigsty | 15.2 KiB | [pg_sqlog_16-1.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_sqlog_16-1.6-1PIGSTY.el10.x86_64.rpm) |
| `pg_sqlog_16` | 1.6 | `el10.aarch64` | pigsty | 15.2 KiB | [pg_sqlog_16-1.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_sqlog_16-1.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-sqlog` | 1.6 | `d12.x86_64` | pigsty | 9.9 KiB | [postgresql-16-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-16-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-sqlog` | 1.6 | `d12.aarch64` | pigsty | 9.9 KiB | [postgresql-16-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-16-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-sqlog` | 1.6 | `u22.x86_64` | pigsty | 9.5 KiB | [postgresql-16-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-16-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-sqlog` | 1.6 | `u22.aarch64` | pigsty | 9.5 KiB | [postgresql-16-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-16-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-sqlog` | 1.6 | `u24.x86_64` | pigsty | 9.5 KiB | [postgresql-16-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-16-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-sqlog` | 1.6 | `u24.aarch64` | pigsty | 9.5 KiB | [postgresql-16-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-16-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_sqlog_15` | 1.6 | `el8.x86_64` | pigsty | 15.1 KiB | [pg_sqlog_15-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_sqlog_15-1.6-1PIGSTY.el8.x86_64.rpm) |
| `pg_sqlog_15` | 1.6 | `el8.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_15-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_sqlog_15-1.6-1PIGSTY.el8.aarch64.rpm) |
| `pg_sqlog_15` | 1.6 | `el9.x86_64` | pigsty | 15.0 KiB | [pg_sqlog_15-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_sqlog_15-1.6-1PIGSTY.el9.x86_64.rpm) |
| `pg_sqlog_15` | 1.6 | `el9.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_15-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_sqlog_15-1.6-1PIGSTY.el9.aarch64.rpm) |
| `pg_sqlog_15` | 1.6 | `el10.x86_64` | pigsty | 15.2 KiB | [pg_sqlog_15-1.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_sqlog_15-1.6-1PIGSTY.el10.x86_64.rpm) |
| `pg_sqlog_15` | 1.6 | `el10.aarch64` | pigsty | 15.2 KiB | [pg_sqlog_15-1.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_sqlog_15-1.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-sqlog` | 1.6 | `d12.x86_64` | pigsty | 9.9 KiB | [postgresql-15-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-15-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-sqlog` | 1.6 | `d12.aarch64` | pigsty | 9.9 KiB | [postgresql-15-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-15-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-sqlog` | 1.6 | `u22.x86_64` | pigsty | 9.5 KiB | [postgresql-15-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-15-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-sqlog` | 1.6 | `u22.aarch64` | pigsty | 9.5 KiB | [postgresql-15-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-15-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-sqlog` | 1.6 | `u24.x86_64` | pigsty | 9.5 KiB | [postgresql-15-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-15-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-sqlog` | 1.6 | `u24.aarch64` | pigsty | 9.5 KiB | [postgresql-15-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-15-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_sqlog_14` | 1.6 | `el8.x86_64` | pigsty | 15.1 KiB | [pg_sqlog_14-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_sqlog_14-1.6-1PIGSTY.el8.x86_64.rpm) |
| `pg_sqlog_14` | 1.6 | `el8.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_14-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_sqlog_14-1.6-1PIGSTY.el8.aarch64.rpm) |
| `pg_sqlog_14` | 1.6 | `el9.x86_64` | pigsty | 15.0 KiB | [pg_sqlog_14-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_sqlog_14-1.6-1PIGSTY.el9.x86_64.rpm) |
| `pg_sqlog_14` | 1.6 | `el9.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_14-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_sqlog_14-1.6-1PIGSTY.el9.aarch64.rpm) |
| `pg_sqlog_14` | 1.6 | `el10.x86_64` | pigsty | 15.2 KiB | [pg_sqlog_14-1.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_sqlog_14-1.6-1PIGSTY.el10.x86_64.rpm) |
| `pg_sqlog_14` | 1.6 | `el10.aarch64` | pigsty | 15.2 KiB | [pg_sqlog_14-1.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_sqlog_14-1.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-sqlog` | 1.6 | `d12.x86_64` | pigsty | 9.9 KiB | [postgresql-14-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-14-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-sqlog` | 1.6 | `d12.aarch64` | pigsty | 9.9 KiB | [postgresql-14-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-14-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-sqlog` | 1.6 | `u22.x86_64` | pigsty | 9.5 KiB | [postgresql-14-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-14-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-sqlog` | 1.6 | `u22.aarch64` | pigsty | 9.5 KiB | [postgresql-14-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-14-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-sqlog` | 1.6 | `u24.x86_64` | pigsty | 9.5 KiB | [postgresql-14-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-14-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-sqlog` | 1.6 | `u24.aarch64` | pigsty | 9.5 KiB | [postgresql-14-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-14-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_sqlog_13` | 1.6 | `el8.x86_64` | pigsty | 15.1 KiB | [pg_sqlog_13-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_sqlog_13-1.6-1PIGSTY.el8.x86_64.rpm) |
| `pg_sqlog_13` | 1.6 | `el8.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_13-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_sqlog_13-1.6-1PIGSTY.el8.aarch64.rpm) |
| `pg_sqlog_13` | 1.6 | `el9.x86_64` | pigsty | 15.0 KiB | [pg_sqlog_13-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_sqlog_13-1.6-1PIGSTY.el9.x86_64.rpm) |
| `pg_sqlog_13` | 1.6 | `el9.aarch64` | pigsty | 15.0 KiB | [pg_sqlog_13-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_sqlog_13-1.6-1PIGSTY.el9.aarch64.rpm) |
| `pg_sqlog_13` | 1.6 | `el10.x86_64` | pigsty | 15.2 KiB | [pg_sqlog_13-1.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_sqlog_13-1.6-1PIGSTY.el10.x86_64.rpm) |
| `pg_sqlog_13` | 1.6 | `el10.aarch64` | pigsty | 15.2 KiB | [pg_sqlog_13-1.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_sqlog_13-1.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-sqlog` | 1.6 | `d12.x86_64` | pigsty | 9.9 KiB | [postgresql-13-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-13-pg-sqlog_1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-sqlog` | 1.6 | `d12.aarch64` | pigsty | 9.9 KiB | [postgresql-13-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sqlog/postgresql-13-pg-sqlog_1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-sqlog` | 1.6 | `u22.x86_64` | pigsty | 9.5 KiB | [postgresql-13-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-13-pg-sqlog_1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-sqlog` | 1.6 | `u22.aarch64` | pigsty | 9.5 KiB | [postgresql-13-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sqlog/postgresql-13-pg-sqlog_1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-sqlog` | 1.6 | `u24.x86_64` | pigsty | 9.5 KiB | [postgresql-13-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-13-pg-sqlog_1.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-sqlog` | 1.6 | `u24.aarch64` | pigsty | 9.5 KiB | [postgresql-13-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sqlog/postgresql-13-pg-sqlog_1.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/kouber/pg_sqlog" title="Repository" icon="github" subtitle="github.com/kouber/pg_sqlog" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_sqlog-1.6.tar.gz" >}}
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

