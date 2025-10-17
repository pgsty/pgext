---
title: "log_fdw"
linkTitle: "log_fdw"
description: "foreign-data wrapper for Postgres log file access"
weight: 8810
categories: ["Fdw"]
width: full
---

foreign-data wrapper for Postgres log file access

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8810** | {{< badge content="log_fdw" link="https://github.com/aws/postgresql-logfdw" >}} | {{< ext "log_fdw" "log_fdw" >}} | `1.4` | {{< category "FDW" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_sqlog" >}} {{< ext "pgaudit" >}} {{< ext "file_fdw" >}} {{< ext "auto_explain" >}} {{< ext "pgauditlogtofile" >}} {{< ext "logerrors" >}} {{< ext "wrappers" >}} {{< ext "multicorn" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/log_fdw" >}} | `1.4` | {{< badge content="18" color="red" alt="log_fdw_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `log_fdw_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/log_fdw" >}} | `1.4` | {{< badge content="18" color="red" alt="postgresql-18-log-fdw" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-log-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "log_fdw_18" >}}     | {{< pkg "log_fdw_17" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/log_fdw_17-1.4-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "log_fdw_16" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/log_fdw_16-1.4-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "log_fdw_15" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/log_fdw_15-1.4-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "log_fdw_14" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/log_fdw_14-1.4-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "log_fdw_18" >}}     | {{< pkg "log_fdw_17" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/log_fdw_17-1.4-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "log_fdw_16" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/log_fdw_16-1.4-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "log_fdw_15" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/log_fdw_15-1.4-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "log_fdw_14" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/log_fdw_14-1.4-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "log_fdw_18" >}}     | {{< pkg "log_fdw_17" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/log_fdw_17-1.4-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "log_fdw_16" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/log_fdw_16-1.4-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "log_fdw_15" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/log_fdw_15-1.4-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "log_fdw_14" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/log_fdw_14-1.4-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "log_fdw_18" >}}     | {{< pkg "log_fdw_17" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/log_fdw_17-1.4-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "log_fdw_16" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/log_fdw_16-1.4-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "log_fdw_15" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/log_fdw_15-1.4-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "log_fdw_14" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/log_fdw_14-1.4-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-log-fdw" >}}     | {{< pkg "postgresql-17-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/log-fdw/postgresql-17-log-fdw_1.4-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/log-fdw/postgresql-16-log-fdw_1.4-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/log-fdw/postgresql-15-log-fdw_1.4-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/log-fdw/postgresql-14-log-fdw_1.4-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-log-fdw" >}}     | {{< pkg "postgresql-17-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/log-fdw/postgresql-17-log-fdw_1.4-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/log-fdw/postgresql-16-log-fdw_1.4-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/log-fdw/postgresql-15-log-fdw_1.4-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/log-fdw/postgresql-14-log-fdw_1.4-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-log-fdw" >}}     | {{< pkg "postgresql-17-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/log-fdw/postgresql-17-log-fdw_1.4-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/log-fdw/postgresql-16-log-fdw_1.4-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/log-fdw/postgresql-15-log-fdw_1.4-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/log-fdw/postgresql-14-log-fdw_1.4-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-log-fdw" >}}     | {{< pkg "postgresql-17-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/log-fdw/postgresql-17-log-fdw_1.4-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/log-fdw/postgresql-16-log-fdw_1.4-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/log-fdw/postgresql-15-log-fdw_1.4-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/log-fdw/postgresql-14-log-fdw_1.4-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-log-fdw" >}}     | {{< pkg "postgresql-17-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/log-fdw/postgresql-17-log-fdw_1.4-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/log-fdw/postgresql-16-log-fdw_1.4-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/log-fdw/postgresql-15-log-fdw_1.4-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/log-fdw/postgresql-14-log-fdw_1.4-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-log-fdw" >}}     | {{< pkg "postgresql-17-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/log-fdw/postgresql-17-log-fdw_1.4-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/log-fdw/postgresql-16-log-fdw_1.4-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/log-fdw/postgresql-15-log-fdw_1.4-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-log-fdw" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/log-fdw/postgresql-14-log-fdw_1.4-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `log_fdw_17` | 1.4 | `el8.aarch64` | pigsty | 19.2 KiB | [log_fdw_17-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/log_fdw_17-1.4-1PIGSTY.el8.aarch64.rpm) |
| `log_fdw_17` | 1.4 | `el8.x86_64` | pigsty | 19.4 KiB | [log_fdw_17-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/log_fdw_17-1.4-1PIGSTY.el8.x86_64.rpm) |
| `log_fdw_17` | 1.4 | `el9.aarch64` | pigsty | 19.4 KiB | [log_fdw_17-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/log_fdw_17-1.4-1PIGSTY.el9.aarch64.rpm) |
| `log_fdw_17` | 1.4 | `el9.x86_64` | pigsty | 19.8 KiB | [log_fdw_17-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/log_fdw_17-1.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-log-fdw` | 1.4 | `d12.x86_64` | pigsty | 32.6 KiB | [postgresql-17-log-fdw_1.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/log-fdw/postgresql-17-log-fdw_1.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-log-fdw` | 1.4 | `d12.aarch64` | pigsty | 32.4 KiB | [postgresql-17-log-fdw_1.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/log-fdw/postgresql-17-log-fdw_1.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-log-fdw` | 1.4 | `u22.aarch64` | pigsty | 33.6 KiB | [postgresql-17-log-fdw_1.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/log-fdw/postgresql-17-log-fdw_1.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-log-fdw` | 1.4 | `u22.x86_64` | pigsty | 34.1 KiB | [postgresql-17-log-fdw_1.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/log-fdw/postgresql-17-log-fdw_1.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-log-fdw` | 1.4 | `u24.x86_64` | pigsty | 28.8 KiB | [postgresql-17-log-fdw_1.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/log-fdw/postgresql-17-log-fdw_1.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-log-fdw` | 1.4 | `u24.aarch64` | pigsty | 28.5 KiB | [postgresql-17-log-fdw_1.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/log-fdw/postgresql-17-log-fdw_1.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `log_fdw_16` | 1.4 | `el8.x86_64` | pigsty | 19.3 KiB | [log_fdw_16-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/log_fdw_16-1.4-1PIGSTY.el8.x86_64.rpm) |
| `log_fdw_16` | 1.4 | `el8.aarch64` | pigsty | 19.2 KiB | [log_fdw_16-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/log_fdw_16-1.4-1PIGSTY.el8.aarch64.rpm) |
| `log_fdw_16` | 1.4 | `el9.aarch64` | pigsty | 19.4 KiB | [log_fdw_16-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/log_fdw_16-1.4-1PIGSTY.el9.aarch64.rpm) |
| `log_fdw_16` | 1.4 | `el9.x86_64` | pigsty | 19.6 KiB | [log_fdw_16-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/log_fdw_16-1.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-16-log-fdw` | 1.4 | `d12.aarch64` | pigsty | 32.3 KiB | [postgresql-16-log-fdw_1.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/log-fdw/postgresql-16-log-fdw_1.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-log-fdw` | 1.4 | `d12.x86_64` | pigsty | 32.6 KiB | [postgresql-16-log-fdw_1.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/log-fdw/postgresql-16-log-fdw_1.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-log-fdw` | 1.4 | `u22.x86_64` | pigsty | 34.1 KiB | [postgresql-16-log-fdw_1.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/log-fdw/postgresql-16-log-fdw_1.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-log-fdw` | 1.4 | `u22.aarch64` | pigsty | 33.7 KiB | [postgresql-16-log-fdw_1.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/log-fdw/postgresql-16-log-fdw_1.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-log-fdw` | 1.4 | `u24.aarch64` | pigsty | 28.5 KiB | [postgresql-16-log-fdw_1.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/log-fdw/postgresql-16-log-fdw_1.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-log-fdw` | 1.4 | `u24.x86_64` | pigsty | 28.7 KiB | [postgresql-16-log-fdw_1.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/log-fdw/postgresql-16-log-fdw_1.4-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `log_fdw_15` | 1.4 | `el8.x86_64` | pigsty | 19.4 KiB | [log_fdw_15-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/log_fdw_15-1.4-1PIGSTY.el8.x86_64.rpm) |
| `log_fdw_15` | 1.4 | `el8.aarch64` | pigsty | 19.2 KiB | [log_fdw_15-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/log_fdw_15-1.4-1PIGSTY.el8.aarch64.rpm) |
| `log_fdw_15` | 1.4 | `el9.x86_64` | pigsty | 19.9 KiB | [log_fdw_15-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/log_fdw_15-1.4-1PIGSTY.el9.x86_64.rpm) |
| `log_fdw_15` | 1.4 | `el9.aarch64` | pigsty | 19.4 KiB | [log_fdw_15-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/log_fdw_15-1.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-log-fdw` | 1.4 | `d12.x86_64` | pigsty | 32.5 KiB | [postgresql-15-log-fdw_1.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/log-fdw/postgresql-15-log-fdw_1.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-log-fdw` | 1.4 | `d12.aarch64` | pigsty | 32.2 KiB | [postgresql-15-log-fdw_1.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/log-fdw/postgresql-15-log-fdw_1.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-log-fdw` | 1.4 | `u22.aarch64` | pigsty | 33.6 KiB | [postgresql-15-log-fdw_1.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/log-fdw/postgresql-15-log-fdw_1.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-log-fdw` | 1.4 | `u22.x86_64` | pigsty | 33.9 KiB | [postgresql-15-log-fdw_1.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/log-fdw/postgresql-15-log-fdw_1.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-log-fdw` | 1.4 | `u24.aarch64` | pigsty | 28.5 KiB | [postgresql-15-log-fdw_1.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/log-fdw/postgresql-15-log-fdw_1.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-log-fdw` | 1.4 | `u24.x86_64` | pigsty | 28.8 KiB | [postgresql-15-log-fdw_1.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/log-fdw/postgresql-15-log-fdw_1.4-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `log_fdw_14` | 1.4 | `el8.aarch64` | pigsty | 19.2 KiB | [log_fdw_14-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/log_fdw_14-1.4-1PIGSTY.el8.aarch64.rpm) |
| `log_fdw_14` | 1.4 | `el8.x86_64` | pigsty | 19.4 KiB | [log_fdw_14-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/log_fdw_14-1.4-1PIGSTY.el8.x86_64.rpm) |
| `log_fdw_14` | 1.4 | `el9.x86_64` | pigsty | 19.7 KiB | [log_fdw_14-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/log_fdw_14-1.4-1PIGSTY.el9.x86_64.rpm) |
| `log_fdw_14` | 1.4 | `el9.aarch64` | pigsty | 19.4 KiB | [log_fdw_14-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/log_fdw_14-1.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-log-fdw` | 1.4 | `d12.aarch64` | pigsty | 32.1 KiB | [postgresql-14-log-fdw_1.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/log-fdw/postgresql-14-log-fdw_1.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-log-fdw` | 1.4 | `d12.x86_64` | pigsty | 32.4 KiB | [postgresql-14-log-fdw_1.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/log-fdw/postgresql-14-log-fdw_1.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-log-fdw` | 1.4 | `u22.x86_64` | pigsty | 33.8 KiB | [postgresql-14-log-fdw_1.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/log-fdw/postgresql-14-log-fdw_1.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-log-fdw` | 1.4 | `u22.aarch64` | pigsty | 33.5 KiB | [postgresql-14-log-fdw_1.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/log-fdw/postgresql-14-log-fdw_1.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-log-fdw` | 1.4 | `u24.aarch64` | pigsty | 28.5 KiB | [postgresql-14-log-fdw_1.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/log-fdw/postgresql-14-log-fdw_1.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-log-fdw` | 1.4 | `u24.x86_64` | pigsty | 28.8 KiB | [postgresql-14-log-fdw_1.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/log-fdw/postgresql-14-log-fdw_1.4-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/aws/postgresql-logfdw" title="Repository" icon="github" subtitle="github.com/aws/postgresql-logfdw" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="log-fdw-1.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build get log_fdw; # get log_fdw source code
pig build dep log_fdw; # install build dependencies
pig build pkg log_fdw; # build extension rpm or deb
pig build ext log_fdw; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install log_fdw; # install by extension name, for the current active PG version
pig ext install log_fdw; # install via package alias, for the active PG version
pig ext install log_fdw -v 18;   # install for PG 18
pig ext install log_fdw -v 17;   # install for PG 17
pig ext install log_fdw -v 16;   # install for PG 16
pig ext install log_fdw -v 15;   # install for PG 15
pig ext install log_fdw -v 14;   # install for PG 14

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION log_fdw;
```

