---
title: "pg_crash"
linkTitle: "pg_crash"
description: "Send random signals to random processes"
weight: 5210
categories: ["Admin"]
width: full
---

Send random signals to random processes

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5210** | {{< badge content="pg_crash" link="https://github.com/cybertec-postgresql/pg_crash" >}} | {{< ext "pg_crash" "pg_crash" >}} | `1.0` | {{< category "ADMIN" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sL---" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="No" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_snakeoil" >}} {{< ext "pg_cheat_funcs" >}} {{< ext "pg_savior" >}} {{< ext "pg_dirtyread" >}} {{< ext "pg_surgery" >}} {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} {{< ext "pg_squeeze" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_crash" >}} | `1.0` | {{< badge content="18" color="red" alt="pg_crash_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_crash_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_crash" >}} | `1.0` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-crash` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_crash_18" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_crash_18-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_crash_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_crash_17-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_crash_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_crash_16-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_crash_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_crash_15-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_crash_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_crash_14-1.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_crash_18" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_crash_18-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_crash_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_crash_17-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_crash_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_crash_16-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_crash_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_crash_15-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_crash_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_crash_14-1.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_crash_18" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_crash_18-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_crash_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_crash_17-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_crash_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_crash_16-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_crash_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_crash_15-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_crash_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_crash_14-1.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_crash_18" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_crash_18-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_crash_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_crash_17-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_crash_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_crash_16-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_crash_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_crash_15-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_crash_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_crash_14-1.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pg-crash" "0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_crash_18` | 1.0 | `el8.aarch64` | pigsty | 13.1 KiB | [pg_crash_18-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_crash_18-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_crash_18` | 1.0 | `el8.x86_64` | pigsty | 13.1 KiB | [pg_crash_18-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_crash_18-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_crash_18` | 1.0 | `el9.aarch64` | pigsty | 12.9 KiB | [pg_crash_18-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_crash_18-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_crash_18` | 1.0 | `el9.x86_64` | pigsty | 13.2 KiB | [pg_crash_18-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_crash_18-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-18-pg-crash` | 0.3 | `d12.x86_64` | pgdg | 12.7 KiB | [postgresql-18-pg-crash_0.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-crash` | 0.3 | `d12.aarch64` | pgdg | 12.8 KiB | [postgresql-18-pg-crash_0.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-crash` | 0.3 | `u22.x86_64` | pgdg | 12.7 KiB | [postgresql-18-pg-crash_0.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-crash` | 0.3 | `u22.aarch64` | pgdg | 12.8 KiB | [postgresql-18-pg-crash_0.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-crash` | 0.3 | `u24.x86_64` | pgdg | 12.8 KiB | [postgresql-18-pg-crash_0.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-crash` | 0.3 | `u24.aarch64` | pgdg | 13.0 KiB | [postgresql-18-pg-crash_0.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_crash_17` | 1.0 | `el8.aarch64` | pigsty | 13.1 KiB | [pg_crash_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_crash_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_crash_17` | 1.0 | `el8.x86_64` | pigsty | 13.1 KiB | [pg_crash_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_crash_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_crash_17` | 1.0 | `el9.x86_64` | pigsty | 13.2 KiB | [pg_crash_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_crash_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_crash_17` | 1.0 | `el9.aarch64` | pigsty | 13.0 KiB | [pg_crash_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_crash_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pg-crash` | 1.0 | `d12.aarch64` | pigsty | 12.9 KiB | [postgresql-17-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-17-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-crash` | 1.0 | `d12.x86_64` | pigsty | 12.8 KiB | [postgresql-17-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-17-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-crash` | 0.3 | `d12.x86_64` | pgdg | 12.7 KiB | [postgresql-17-pg-crash_0.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-crash` | 0.3 | `d12.aarch64` | pgdg | 12.8 KiB | [postgresql-17-pg-crash_0.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-crash` | 1.0 | `u22.x86_64` | pigsty | 13.2 KiB | [postgresql-17-pg-crash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-17-pg-crash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-crash` | 1.0 | `u22.aarch64` | pigsty | 13.4 KiB | [postgresql-17-pg-crash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-17-pg-crash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-crash` | 0.3 | `u22.x86_64` | pgdg | 13.0 KiB | [postgresql-17-pg-crash_0.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-crash` | 0.3 | `u22.aarch64` | pgdg | 13.1 KiB | [postgresql-17-pg-crash_0.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-crash` | 1.0 | `u24.aarch64` | pigsty | 13.0 KiB | [postgresql-17-pg-crash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-17-pg-crash_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-crash` | 1.0 | `u24.x86_64` | pigsty | 13.0 KiB | [postgresql-17-pg-crash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-17-pg-crash_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-crash` | 0.3 | `u24.aarch64` | pgdg | 13.0 KiB | [postgresql-17-pg-crash_0.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pg-crash` | 0.3 | `u24.x86_64` | pgdg | 12.8 KiB | [postgresql-17-pg-crash_0.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_crash_16` | 1.0 | `el8.aarch64` | pigsty | 13.1 KiB | [pg_crash_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_crash_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_crash_16` | 1.0 | `el8.x86_64` | pigsty | 13.1 KiB | [pg_crash_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_crash_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_crash_16` | 1.0 | `el9.x86_64` | pigsty | 13.2 KiB | [pg_crash_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_crash_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_crash_16` | 1.0 | `el9.aarch64` | pigsty | 13.0 KiB | [pg_crash_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_crash_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-crash` | 1.0 | `d12.x86_64` | pigsty | 12.8 KiB | [postgresql-16-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-16-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-crash` | 1.0 | `d12.aarch64` | pigsty | 12.9 KiB | [postgresql-16-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-16-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-crash` | 0.3 | `d12.x86_64` | pgdg | 12.7 KiB | [postgresql-16-pg-crash_0.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-crash` | 0.3 | `d12.aarch64` | pgdg | 12.9 KiB | [postgresql-16-pg-crash_0.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-crash` | 1.0 | `u22.aarch64` | pigsty | 13.4 KiB | [postgresql-16-pg-crash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-16-pg-crash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-crash` | 1.0 | `u22.x86_64` | pigsty | 13.2 KiB | [postgresql-16-pg-crash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-16-pg-crash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-crash` | 0.3 | `u22.aarch64` | pgdg | 13.1 KiB | [postgresql-16-pg-crash_0.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-crash` | 0.3 | `u22.x86_64` | pgdg | 13.0 KiB | [postgresql-16-pg-crash_0.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-crash` | 1.0 | `u24.aarch64` | pigsty | 13.0 KiB | [postgresql-16-pg-crash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-16-pg-crash_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-crash` | 1.0 | `u24.x86_64` | pigsty | 13.0 KiB | [postgresql-16-pg-crash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-16-pg-crash_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-crash` | 0.3 | `u24.x86_64` | pgdg | 12.9 KiB | [postgresql-16-pg-crash_0.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-crash` | 0.3 | `u24.aarch64` | pgdg | 13.0 KiB | [postgresql-16-pg-crash_0.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_crash_15` | 1.0 | `el8.aarch64` | pigsty | 13.1 KiB | [pg_crash_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_crash_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_crash_15` | 1.0 | `el8.x86_64` | pigsty | 13.1 KiB | [pg_crash_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_crash_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_crash_15` | 1.0 | `el9.aarch64` | pigsty | 13.0 KiB | [pg_crash_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_crash_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_crash_15` | 1.0 | `el9.x86_64` | pigsty | 13.2 KiB | [pg_crash_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_crash_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-15-pg-crash` | 1.0 | `d12.aarch64` | pigsty | 12.9 KiB | [postgresql-15-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-15-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-crash` | 1.0 | `d12.x86_64` | pigsty | 12.9 KiB | [postgresql-15-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-15-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-crash` | 0.3 | `d12.aarch64` | pgdg | 12.8 KiB | [postgresql-15-pg-crash_0.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-crash` | 0.3 | `d12.x86_64` | pgdg | 12.7 KiB | [postgresql-15-pg-crash_0.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-crash` | 1.0 | `u22.aarch64` | pigsty | 13.4 KiB | [postgresql-15-pg-crash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-15-pg-crash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-crash` | 1.0 | `u22.x86_64` | pigsty | 13.2 KiB | [postgresql-15-pg-crash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-15-pg-crash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-crash` | 0.3 | `u22.aarch64` | pgdg | 13.1 KiB | [postgresql-15-pg-crash_0.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-crash` | 0.3 | `u22.x86_64` | pgdg | 13.0 KiB | [postgresql-15-pg-crash_0.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-crash` | 1.0 | `u24.x86_64` | pigsty | 13.0 KiB | [postgresql-15-pg-crash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-15-pg-crash_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-crash` | 1.0 | `u24.aarch64` | pigsty | 13.0 KiB | [postgresql-15-pg-crash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-15-pg-crash_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-crash` | 0.3 | `u24.aarch64` | pgdg | 13.0 KiB | [postgresql-15-pg-crash_0.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pg-crash` | 0.3 | `u24.x86_64` | pgdg | 12.9 KiB | [postgresql-15-pg-crash_0.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_crash_14` | 1.0 | `el8.x86_64` | pigsty | 13.1 KiB | [pg_crash_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_crash_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_crash_14` | 1.0 | `el8.aarch64` | pigsty | 13.1 KiB | [pg_crash_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_crash_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_crash_14` | 1.0 | `el9.x86_64` | pigsty | 13.2 KiB | [pg_crash_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_crash_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_crash_14` | 1.0 | `el9.aarch64` | pigsty | 13.0 KiB | [pg_crash_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_crash_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-crash` | 1.0 | `d12.aarch64` | pigsty | 12.9 KiB | [postgresql-14-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-14-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-crash` | 1.0 | `d12.x86_64` | pigsty | 12.8 KiB | [postgresql-14-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-14-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-crash` | 0.3 | `d12.aarch64` | pgdg | 12.8 KiB | [postgresql-14-pg-crash_0.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-crash` | 0.3 | `d12.x86_64` | pgdg | 12.7 KiB | [postgresql-14-pg-crash_0.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-crash` | 1.0 | `u22.aarch64` | pigsty | 13.4 KiB | [postgresql-14-pg-crash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-14-pg-crash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-crash` | 1.0 | `u22.x86_64` | pigsty | 13.2 KiB | [postgresql-14-pg-crash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-14-pg-crash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-crash` | 0.3 | `u22.x86_64` | pgdg | 13.0 KiB | [postgresql-14-pg-crash_0.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-crash` | 0.3 | `u22.aarch64` | pgdg | 13.1 KiB | [postgresql-14-pg-crash_0.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-crash` | 1.0 | `u24.x86_64` | pigsty | 13.0 KiB | [postgresql-14-pg-crash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-14-pg-crash_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-crash` | 1.0 | `u24.aarch64` | pigsty | 13.0 KiB | [postgresql-14-pg-crash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-14-pg-crash_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-crash` | 0.3 | `u24.x86_64` | pgdg | 12.8 KiB | [postgresql-14-pg-crash_0.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-crash` | 0.3 | `u24.aarch64` | pgdg | 13.0 KiB | [postgresql-14-pg-crash_0.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_crash_13` | 1.0 | `el8.x86_64` | pigsty | 13.0 KiB | [pg_crash_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_crash_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_crash_13` | 1.0 | `el8.aarch64` | pigsty | 13.1 KiB | [pg_crash_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_crash_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_crash_13` | 1.0 | `el9.aarch64` | pigsty | 13.0 KiB | [pg_crash_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_crash_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_crash_13` | 1.0 | `el9.x86_64` | pigsty | 13.2 KiB | [pg_crash_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_crash_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-crash` | 1.0 | `d12.x86_64` | pigsty | 12.8 KiB | [postgresql-13-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-13-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-crash` | 1.0 | `d12.aarch64` | pigsty | 12.9 KiB | [postgresql-13-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-13-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-crash` | 0.3 | `d12.aarch64` | pgdg | 12.7 KiB | [postgresql-13-pg-crash_0.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-13-pg-crash_0.3-2.pgdg12+1_arm64.deb) |
| `postgresql-13-pg-crash` | 0.3 | `d12.x86_64` | pgdg | 12.5 KiB | [postgresql-13-pg-crash_0.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-13-pg-crash_0.3-2.pgdg12+1_amd64.deb) |
| `postgresql-13-pg-crash` | 1.0 | `u22.aarch64` | pigsty | 13.4 KiB | [postgresql-13-pg-crash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-13-pg-crash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-crash` | 1.0 | `u22.x86_64` | pigsty | 13.2 KiB | [postgresql-13-pg-crash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-13-pg-crash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-crash` | 0.3 | `u22.aarch64` | pgdg | 12.9 KiB | [postgresql-13-pg-crash_0.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-13-pg-crash_0.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pg-crash` | 0.3 | `u22.x86_64` | pgdg | 12.8 KiB | [postgresql-13-pg-crash_0.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-13-pg-crash_0.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pg-crash` | 1.0 | `u24.aarch64` | pigsty | 13.0 KiB | [postgresql-13-pg-crash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-13-pg-crash_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-crash` | 1.0 | `u24.x86_64` | pigsty | 13.0 KiB | [postgresql-13-pg-crash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-13-pg-crash_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-crash` | 0.3 | `u24.aarch64` | pgdg | 12.7 KiB | [postgresql-13-pg-crash_0.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-13-pg-crash_0.3-2.pgdg24.04+1_arm64.deb) |
| `postgresql-13-pg-crash` | 0.3 | `u24.x86_64` | pgdg | 12.6 KiB | [postgresql-13-pg-crash_0.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-13-pg-crash_0.3-2.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/cybertec-postgresql/pg_crash" title="Repository" icon="github" subtitle="github.com/cybertec-postgresql/pg_crash" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_crash-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_crash; # get pg_crash source code
pig build dep pg_crash; # install build dependencies
pig build pkg pg_crash; # build extension rpm or deb
pig build ext pg_crash; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_crash; # install by extension name, for the current active PG version
pig ext install pg_crash; # install via package alias, for the active PG version
pig ext install pg_crash -v 18;   # install for PG 18
pig ext install pg_crash -v 17;   # install for PG 17
pig ext install pg_crash -v 16;   # install for PG 16
pig ext install pg_crash -v 15;   # install for PG 15
pig ext install pg_crash -v 14;   # install for PG 14
pig ext install pg_crash -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_crash;
```

