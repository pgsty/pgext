---
title: "fio"
linkTitle: "fio"
description: "PostgreSQL File I/O Functions"
weight: 5230
categories: ["Admin"]
width: full
---

PostgreSQL File I/O Functions

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5230** | {{< badge content="fio" link="https://github.com/csimsek/pgsql-fio" >}} | {{< ext "fio" "pg_fio" >}} | `1.0` | {{< category "ADMIN" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgfincore" >}} {{< ext "adminpack" >}} {{< ext "file_fdw" >}} {{< ext "pageinspect" >}} {{< ext "pgstattuple" >}} {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} {{< ext "pg_squeeze" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/fio" >}} | `1.0` | {{< badge content="18" color="red" alt="pg_fio_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_fio_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/fio" >}} | `1.0` | {{< badge content="18" color="red" alt="postgresql-18-pg-fio" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-fio` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_fio_18" >}}     | {{< pkg "pg_fio_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fio_17-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_fio_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fio_16-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_fio_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fio_15-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_fio_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fio_14-1.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_fio_18" >}}     | {{< pkg "pg_fio_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fio_17-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_fio_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fio_16-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_fio_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fio_15-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_fio_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fio_14-1.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_fio_18" >}}     | {{< pkg "pg_fio_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fio_17-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_fio_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fio_16-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_fio_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fio_15-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_fio_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fio_14-1.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_fio_18" >}}     | {{< pkg "pg_fio_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fio_17-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_fio_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fio_16-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_fio_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fio_15-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_fio_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fio_14-1.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-fio" >}}     | {{< pkg "postgresql-17-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fio/postgresql-17-pg-fio_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fio/postgresql-16-pg-fio_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fio/postgresql-15-pg-fio_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fio/postgresql-14-pg-fio_1.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-fio" >}}     | {{< pkg "postgresql-17-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fio/postgresql-17-pg-fio_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fio/postgresql-16-pg-fio_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fio/postgresql-15-pg-fio_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fio/postgresql-14-pg-fio_1.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-fio" >}}     | {{< pkg "postgresql-17-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fio/postgresql-17-pg-fio_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fio/postgresql-16-pg-fio_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fio/postgresql-15-pg-fio_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fio/postgresql-14-pg-fio_1.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-fio" >}}     | {{< pkg "postgresql-17-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fio/postgresql-17-pg-fio_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fio/postgresql-16-pg-fio_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fio/postgresql-15-pg-fio_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fio/postgresql-14-pg-fio_1.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-fio" >}}     | {{< pkg "postgresql-17-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fio/postgresql-17-pg-fio_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fio/postgresql-16-pg-fio_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fio/postgresql-15-pg-fio_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fio/postgresql-14-pg-fio_1.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-fio" >}}     | {{< pkg "postgresql-17-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fio/postgresql-17-pg-fio_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fio/postgresql-16-pg-fio_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fio/postgresql-15-pg-fio_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-fio" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fio/postgresql-14-pg-fio_1.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_fio_17` | 1.0 | `el8.x86_64` | pigsty | 14.4 KiB | [pg_fio_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fio_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_fio_17` | 1.0 | `el8.aarch64` | pigsty | 14.1 KiB | [pg_fio_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fio_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_fio_17` | 1.0 | `el9.aarch64` | pigsty | 14.1 KiB | [pg_fio_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fio_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_fio_17` | 1.0 | `el9.x86_64` | pigsty | 14.5 KiB | [pg_fio_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fio_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-fio` | 1.0 | `d12.x86_64` | pigsty | 26.9 KiB | [postgresql-17-pg-fio_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fio/postgresql-17-pg-fio_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-fio` | 1.0 | `d12.aarch64` | pigsty | 26.8 KiB | [postgresql-17-pg-fio_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fio/postgresql-17-pg-fio_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-fio` | 1.0 | `u22.x86_64` | pigsty | 28.1 KiB | [postgresql-17-pg-fio_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fio/postgresql-17-pg-fio_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-fio` | 1.0 | `u22.aarch64` | pigsty | 27.7 KiB | [postgresql-17-pg-fio_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fio/postgresql-17-pg-fio_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-fio` | 1.0 | `u24.x86_64` | pigsty | 24.4 KiB | [postgresql-17-pg-fio_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fio/postgresql-17-pg-fio_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-fio` | 1.0 | `u24.aarch64` | pigsty | 24.2 KiB | [postgresql-17-pg-fio_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fio/postgresql-17-pg-fio_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_fio_16` | 1.0 | `el8.x86_64` | pigsty | 14.4 KiB | [pg_fio_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fio_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_fio_16` | 1.0 | `el8.aarch64` | pigsty | 14.1 KiB | [pg_fio_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fio_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_fio_16` | 1.0 | `el9.x86_64` | pigsty | 14.5 KiB | [pg_fio_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fio_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_fio_16` | 1.0 | `el9.aarch64` | pigsty | 14.1 KiB | [pg_fio_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fio_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-fio` | 1.0 | `d12.x86_64` | pigsty | 26.8 KiB | [postgresql-16-pg-fio_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fio/postgresql-16-pg-fio_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-fio` | 1.0 | `d12.aarch64` | pigsty | 26.7 KiB | [postgresql-16-pg-fio_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fio/postgresql-16-pg-fio_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-fio` | 1.0 | `u22.aarch64` | pigsty | 27.6 KiB | [postgresql-16-pg-fio_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fio/postgresql-16-pg-fio_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-fio` | 1.0 | `u22.x86_64` | pigsty | 28.1 KiB | [postgresql-16-pg-fio_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fio/postgresql-16-pg-fio_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-fio` | 1.0 | `u24.x86_64` | pigsty | 24.4 KiB | [postgresql-16-pg-fio_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fio/postgresql-16-pg-fio_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-fio` | 1.0 | `u24.aarch64` | pigsty | 24.2 KiB | [postgresql-16-pg-fio_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fio/postgresql-16-pg-fio_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_fio_15` | 1.0 | `el8.x86_64` | pigsty | 14.4 KiB | [pg_fio_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fio_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_fio_15` | 1.0 | `el8.aarch64` | pigsty | 14.2 KiB | [pg_fio_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fio_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_fio_15` | 1.0 | `el9.x86_64` | pigsty | 14.6 KiB | [pg_fio_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fio_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_fio_15` | 1.0 | `el9.aarch64` | pigsty | 14.2 KiB | [pg_fio_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fio_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-fio` | 1.0 | `d12.aarch64` | pigsty | 26.8 KiB | [postgresql-15-pg-fio_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fio/postgresql-15-pg-fio_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-fio` | 1.0 | `d12.x86_64` | pigsty | 26.9 KiB | [postgresql-15-pg-fio_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fio/postgresql-15-pg-fio_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-fio` | 1.0 | `u22.aarch64` | pigsty | 27.8 KiB | [postgresql-15-pg-fio_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fio/postgresql-15-pg-fio_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-fio` | 1.0 | `u22.x86_64` | pigsty | 28.2 KiB | [postgresql-15-pg-fio_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fio/postgresql-15-pg-fio_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-fio` | 1.0 | `u24.x86_64` | pigsty | 24.5 KiB | [postgresql-15-pg-fio_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fio/postgresql-15-pg-fio_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-fio` | 1.0 | `u24.aarch64` | pigsty | 24.3 KiB | [postgresql-15-pg-fio_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fio/postgresql-15-pg-fio_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_fio_14` | 1.0 | `el8.x86_64` | pigsty | 14.4 KiB | [pg_fio_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fio_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_fio_14` | 1.0 | `el8.aarch64` | pigsty | 14.2 KiB | [pg_fio_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fio_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_fio_14` | 1.0 | `el9.x86_64` | pigsty | 14.6 KiB | [pg_fio_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fio_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_fio_14` | 1.0 | `el9.aarch64` | pigsty | 14.2 KiB | [pg_fio_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fio_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-fio` | 1.0 | `d12.x86_64` | pigsty | 26.9 KiB | [postgresql-14-pg-fio_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fio/postgresql-14-pg-fio_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-fio` | 1.0 | `d12.aarch64` | pigsty | 26.8 KiB | [postgresql-14-pg-fio_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fio/postgresql-14-pg-fio_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-fio` | 1.0 | `u22.x86_64` | pigsty | 28.1 KiB | [postgresql-14-pg-fio_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fio/postgresql-14-pg-fio_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-fio` | 1.0 | `u22.aarch64` | pigsty | 27.7 KiB | [postgresql-14-pg-fio_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fio/postgresql-14-pg-fio_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-fio` | 1.0 | `u24.x86_64` | pigsty | 24.4 KiB | [postgresql-14-pg-fio_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fio/postgresql-14-pg-fio_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-fio` | 1.0 | `u24.aarch64` | pigsty | 24.3 KiB | [postgresql-14-pg-fio_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fio/postgresql-14-pg-fio_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_fio_13` | 1.0 | `el8.aarch64` | pigsty | 14.1 KiB | [pg_fio_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fio_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_fio_13` | 1.0 | `el8.x86_64` | pigsty | 14.3 KiB | [pg_fio_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fio_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_fio_13` | 1.0 | `el9.aarch64` | pigsty | 14.2 KiB | [pg_fio_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fio_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_fio_13` | 1.0 | `el9.x86_64` | pigsty | 14.6 KiB | [pg_fio_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fio_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-fio` | 1.0 | `d12.aarch64` | pigsty | 26.5 KiB | [postgresql-13-pg-fio_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fio/postgresql-13-pg-fio_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-fio` | 1.0 | `d12.x86_64` | pigsty | 26.6 KiB | [postgresql-13-pg-fio_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fio/postgresql-13-pg-fio_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-fio` | 1.0 | `u22.aarch64` | pigsty | 27.4 KiB | [postgresql-13-pg-fio_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fio/postgresql-13-pg-fio_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-fio` | 1.0 | `u22.x86_64` | pigsty | 27.9 KiB | [postgresql-13-pg-fio_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fio/postgresql-13-pg-fio_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-fio` | 1.0 | `u24.aarch64` | pigsty | 24.2 KiB | [postgresql-13-pg-fio_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fio/postgresql-13-pg-fio_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-fio` | 1.0 | `u24.x86_64` | pigsty | 24.3 KiB | [postgresql-13-pg-fio_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fio/postgresql-13-pg-fio_1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/csimsek/pgsql-fio" title="Repository" icon="github" subtitle="github.com/csimsek/pgsql-fio" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_fio-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get fio; # get fio source code
pig build dep fio; # install build dependencies
pig build pkg fio; # build extension rpm or deb
pig build ext fio; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install fio; # install by extension name, for the current active PG version
pig ext install pg_fio; # install via package alias, for the active PG version
pig ext install fio -v 18;   # install for PG 18
pig ext install fio -v 17;   # install for PG 17
pig ext install fio -v 16;   # install for PG 16
pig ext install fio -v 15;   # install for PG 15
pig ext install fio -v 14;   # install for PG 14
pig ext install fio -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION fio;
```

