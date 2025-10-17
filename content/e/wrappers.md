---
title: "wrappers"
linkTitle: "wrappers"
description: "Foreign data wrappers developed by Supabase"
weight: 8500
categories: ["Fdw"]
width: full
---

Foreign data wrappers developed by Supabase

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8500** | {{< badge content="wrappers" link="https://github.com/supabase/wrappers" >}} | {{< ext "wrappers" "wrappers" >}} | `0.5.4` | {{< category "FDW" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "multicorn" >}} {{< ext "odbc_fdw" >}} {{< ext "jdbc_fdw" >}} {{< ext "pgspider_ext" >}} {{< ext "mysql_fdw" >}} {{< ext "db2_fdw" >}} {{< ext "pgbouncer_fdw" >}} {{< ext "mongo_fdw" >}} |

> [!Note] pgrx=0.14.3


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/wrappers" >}} | `0.5.4` | {{< badge content="18" color="red" alt="wrappers_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `wrappers_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/wrappers" >}} | `0.5.4` | {{< badge content="18" color="red" alt="postgresql-18-wrappers" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-wrappers` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "wrappers_18" >}}     | {{< pkg "wrappers_17" "0.5.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_17-0.5.4-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "wrappers_16" "0.5.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_16-0.5.4-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "wrappers_15" "0.5.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_15-0.5.4-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "wrappers_14" "0.5.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_14-0.5.4-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "wrappers_18" >}}     | {{< pkg "wrappers_17" "0.5.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_17-0.5.4-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "wrappers_16" "0.5.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_16-0.5.4-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "wrappers_15" "0.5.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_15-0.5.4-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "wrappers_14" "0.5.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_14-0.5.4-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "wrappers_18" >}}     | {{< pkg "wrappers_17" "0.5.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_17-0.5.4-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "wrappers_16" "0.5.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_16-0.5.4-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "wrappers_15" "0.5.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_15-0.5.4-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "wrappers_14" "0.5.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_14-0.5.4-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "wrappers_18" >}}     | {{< pkg "wrappers_17" "0.5.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_17-0.5.4-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "wrappers_16" "0.5.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_16-0.5.4-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "wrappers_15" "0.5.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_15-0.5.4-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "wrappers_14" "0.5.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_14-0.5.4-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-wrappers" >}}     | {{< pkg "postgresql-17-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-17-wrappers_0.5.4-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-16-wrappers_0.5.4-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-15-wrappers_0.5.4-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-14-wrappers_0.5.4-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-wrappers" >}}     | {{< pkg "postgresql-17-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-17-wrappers_0.5.4-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-16-wrappers_0.5.4-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-15-wrappers_0.5.4-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-14-wrappers_0.5.4-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-wrappers" >}}     | {{< pkg "postgresql-17-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-17-wrappers_0.5.4-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-16-wrappers_0.5.4-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-15-wrappers_0.5.4-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-14-wrappers_0.5.4-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-wrappers" >}}     | {{< pkg "postgresql-17-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-17-wrappers_0.5.4-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-16-wrappers_0.5.4-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-15-wrappers_0.5.4-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-14-wrappers_0.5.4-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-wrappers" >}}     | {{< pkg "postgresql-17-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-17-wrappers_0.5.4-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-16-wrappers_0.5.4-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-15-wrappers_0.5.4-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-14-wrappers_0.5.4-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-wrappers" >}}     | {{< pkg "postgresql-17-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-17-wrappers_0.5.4-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-16-wrappers_0.5.4-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-15-wrappers_0.5.4-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-wrappers" "0.5.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-14-wrappers_0.5.4-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `wrappers_17` | 0.5.4 | `el8.aarch64` | pigsty | 156.1 KiB | [wrappers_17-0.5.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_17-0.5.4-1PIGSTY.el8.aarch64.rpm) |
| `wrappers_17` | 0.5.4 | `el8.x86_64` | pigsty | 167.9 KiB | [wrappers_17-0.5.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_17-0.5.4-1PIGSTY.el8.x86_64.rpm) |
| `wrappers_17` | 0.5.4 | `el9.aarch64` | pigsty | 167.0 KiB | [wrappers_17-0.5.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_17-0.5.4-1PIGSTY.el9.aarch64.rpm) |
| `wrappers_17` | 0.5.4 | `el9.x86_64` | pigsty | 171.9 KiB | [wrappers_17-0.5.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_17-0.5.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-wrappers` | 0.5.4 | `d12.x86_64` | pigsty | 138.7 KiB | [postgresql-17-wrappers_0.5.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-17-wrappers_0.5.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-wrappers` | 0.5.4 | `d12.aarch64` | pigsty | 123.0 KiB | [postgresql-17-wrappers_0.5.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-17-wrappers_0.5.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-wrappers` | 0.5.4 | `u22.aarch64` | pigsty | 142.2 KiB | [postgresql-17-wrappers_0.5.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-17-wrappers_0.5.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-wrappers` | 0.5.4 | `u22.x86_64` | pigsty | 150.8 KiB | [postgresql-17-wrappers_0.5.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-17-wrappers_0.5.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-wrappers` | 0.5.4 | `u24.x86_64` | pigsty | 149.8 KiB | [postgresql-17-wrappers_0.5.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-17-wrappers_0.5.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-wrappers` | 0.5.4 | `u24.aarch64` | pigsty | 140.6 KiB | [postgresql-17-wrappers_0.5.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-17-wrappers_0.5.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `wrappers_16` | 0.5.4 | `el8.x86_64` | pigsty | 167.9 KiB | [wrappers_16-0.5.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_16-0.5.4-1PIGSTY.el8.x86_64.rpm) |
| `wrappers_16` | 0.5.4 | `el8.aarch64` | pigsty | 156.2 KiB | [wrappers_16-0.5.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_16-0.5.4-1PIGSTY.el8.aarch64.rpm) |
| `wrappers_16` | 0.5.4 | `el9.aarch64` | pigsty | 166.9 KiB | [wrappers_16-0.5.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_16-0.5.4-1PIGSTY.el9.aarch64.rpm) |
| `wrappers_16` | 0.5.4 | `el9.x86_64` | pigsty | 171.8 KiB | [wrappers_16-0.5.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_16-0.5.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-16-wrappers` | 0.5.4 | `d12.aarch64` | pigsty | 123.0 KiB | [postgresql-16-wrappers_0.5.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-16-wrappers_0.5.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-wrappers` | 0.5.4 | `d12.x86_64` | pigsty | 138.6 KiB | [postgresql-16-wrappers_0.5.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-16-wrappers_0.5.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-wrappers` | 0.5.4 | `u22.x86_64` | pigsty | 150.8 KiB | [postgresql-16-wrappers_0.5.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-16-wrappers_0.5.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-wrappers` | 0.5.4 | `u22.aarch64` | pigsty | 142.0 KiB | [postgresql-16-wrappers_0.5.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-16-wrappers_0.5.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-wrappers` | 0.5.4 | `u24.aarch64` | pigsty | 140.6 KiB | [postgresql-16-wrappers_0.5.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-16-wrappers_0.5.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-wrappers` | 0.5.4 | `u24.x86_64` | pigsty | 149.8 KiB | [postgresql-16-wrappers_0.5.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-16-wrappers_0.5.4-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `wrappers_15` | 0.5.4 | `el8.x86_64` | pigsty | 167.9 KiB | [wrappers_15-0.5.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_15-0.5.4-1PIGSTY.el8.x86_64.rpm) |
| `wrappers_15` | 0.5.4 | `el8.aarch64` | pigsty | 156.1 KiB | [wrappers_15-0.5.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_15-0.5.4-1PIGSTY.el8.aarch64.rpm) |
| `wrappers_15` | 0.5.4 | `el9.x86_64` | pigsty | 171.9 KiB | [wrappers_15-0.5.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_15-0.5.4-1PIGSTY.el9.x86_64.rpm) |
| `wrappers_15` | 0.5.4 | `el9.aarch64` | pigsty | 166.9 KiB | [wrappers_15-0.5.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_15-0.5.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-wrappers` | 0.5.4 | `d12.x86_64` | pigsty | 138.7 KiB | [postgresql-15-wrappers_0.5.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-15-wrappers_0.5.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-wrappers` | 0.5.4 | `d12.aarch64` | pigsty | 123.0 KiB | [postgresql-15-wrappers_0.5.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-15-wrappers_0.5.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-wrappers` | 0.5.4 | `u22.aarch64` | pigsty | 142.2 KiB | [postgresql-15-wrappers_0.5.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-15-wrappers_0.5.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-wrappers` | 0.5.4 | `u22.x86_64` | pigsty | 150.9 KiB | [postgresql-15-wrappers_0.5.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-15-wrappers_0.5.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-wrappers` | 0.5.4 | `u24.aarch64` | pigsty | 140.6 KiB | [postgresql-15-wrappers_0.5.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-15-wrappers_0.5.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-wrappers` | 0.5.4 | `u24.x86_64` | pigsty | 149.8 KiB | [postgresql-15-wrappers_0.5.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-15-wrappers_0.5.4-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `wrappers_14` | 0.5.4 | `el8.aarch64` | pigsty | 156.1 KiB | [wrappers_14-0.5.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_14-0.5.4-1PIGSTY.el8.aarch64.rpm) |
| `wrappers_14` | 0.5.4 | `el8.x86_64` | pigsty | 167.9 KiB | [wrappers_14-0.5.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_14-0.5.4-1PIGSTY.el8.x86_64.rpm) |
| `wrappers_14` | 0.5.4 | `el9.x86_64` | pigsty | 171.9 KiB | [wrappers_14-0.5.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_14-0.5.4-1PIGSTY.el9.x86_64.rpm) |
| `wrappers_14` | 0.5.4 | `el9.aarch64` | pigsty | 166.9 KiB | [wrappers_14-0.5.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_14-0.5.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-wrappers` | 0.5.4 | `d12.aarch64` | pigsty | 123.0 KiB | [postgresql-14-wrappers_0.5.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-14-wrappers_0.5.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-wrappers` | 0.5.4 | `d12.x86_64` | pigsty | 138.7 KiB | [postgresql-14-wrappers_0.5.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-14-wrappers_0.5.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-wrappers` | 0.5.4 | `u22.x86_64` | pigsty | 150.8 KiB | [postgresql-14-wrappers_0.5.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-14-wrappers_0.5.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-wrappers` | 0.5.4 | `u22.aarch64` | pigsty | 141.8 KiB | [postgresql-14-wrappers_0.5.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-14-wrappers_0.5.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-wrappers` | 0.5.4 | `u24.aarch64` | pigsty | 140.7 KiB | [postgresql-14-wrappers_0.5.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-14-wrappers_0.5.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-wrappers` | 0.5.4 | `u24.x86_64` | pigsty | 149.8 KiB | [postgresql-14-wrappers_0.5.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-14-wrappers_0.5.4-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supabase/wrappers" title="Repository" icon="github" subtitle="github.com/supabase/wrappers" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="wrappers-0.5.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build get wrappers; # get wrappers source code
pig build dep wrappers; # install build dependencies
pig build pkg wrappers; # build extension rpm or deb
pig build ext wrappers; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install wrappers; # install by extension name, for the current active PG version
pig ext install wrappers; # install via package alias, for the active PG version
pig ext install wrappers -v 17;   # install for PG 17
pig ext install wrappers -v 16;   # install for PG 16
pig ext install wrappers -v 15;   # install for PG 15
pig ext install wrappers -v 14;   # install for PG 14

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION wrappers;
```

