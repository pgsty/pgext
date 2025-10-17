---
title: "pgfincore"
linkTitle: "pgfincore"
description: "examine and manage the os buffer cache"
weight: 5060
categories: ["Admin"]
width: full
---

examine and manage the os buffer cache

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5060** | {{< badge content="pgfincore" link="https://github.com/klando/pgfincore" >}} | {{< ext "pgfincore" "pgfincore" >}} | `1.3.1` | {{< category "ADMIN" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_cooldown" >}} {{< ext "pgcozy" >}} {{< ext "fio" >}} {{< ext "pg_prewarm" >}} {{< ext "pgmeminfo" >}} {{< ext "pg_buffercache" >}} {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pgfincore" >}} | `1.3.1` | {{< badge content="18" color="red" alt="pgfincore_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pgfincore_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pgfincore" >}} | `1.3.1` | {{< badge content="18" color="red" alt="postgresql-18-pgfincore" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pgfincore` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pgfincore_18" >}}     | {{< pkg "pgfincore_17" "1.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgfincore_17-1.3.1-3PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgfincore_16" "1.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgfincore_16-1.3.1-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgfincore_15" "1.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgfincore_15-1.3.1-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgfincore_14" "1.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgfincore_14-1.3.1-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pgfincore_18" >}}     | {{< pkg "pgfincore_17" "1.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgfincore_17-1.3.1-3PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgfincore_16" "1.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgfincore_16-1.3.1-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgfincore_15" "1.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgfincore_15-1.3.1-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgfincore_14" "1.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgfincore_14-1.3.1-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pgfincore_18" >}}     | {{< pkg "pgfincore_17" "1.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgfincore_17-1.3.1-3PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgfincore_16" "1.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgfincore_16-1.3.1-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgfincore_15" "1.2.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgfincore_15-1.2.4-1.rhel9.x86_64.rpm" >}} | {{< pkg "pgfincore_14" "1.2.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgfincore_14-1.2.4-1.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pgfincore_18" >}}     | {{< pkg "pgfincore_17" "1.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgfincore_17-1.3.1-3PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgfincore_16" "1.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgfincore_16-1.3.1-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgfincore_15" "1.2.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgfincore_15-1.2.4-1.rhel9.aarch64.rpm" >}} | {{< pkg "pgfincore_14" "1.2.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgfincore_14-1.2.4-1.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pgfincore" >}}     | {{< pkg "postgresql-17-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-2.pgdg120+1_amd64.deb" >}} | {{< pkg "postgresql-16-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-2.pgdg120+1_amd64.deb" >}} | {{< pkg "postgresql-15-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-2.pgdg120+1_amd64.deb" >}} | {{< pkg "postgresql-14-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-2.pgdg120+1_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pgfincore" >}}     | {{< pkg "postgresql-17-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-2.pgdg120+1_arm64.deb" >}} | {{< pkg "postgresql-16-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-2.pgdg120+1_arm64.deb" >}} | {{< pkg "postgresql-15-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-2.pgdg120+1_arm64.deb" >}} | {{< pkg "postgresql-14-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-2.pgdg120+1_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pgfincore" >}}     | {{< pkg "postgresql-17-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-2.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pgfincore" >}}     | {{< pkg "postgresql-17-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-2.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pgfincore" >}}     | {{< pkg "postgresql-17-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-2.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pgfincore" >}}     | {{< pkg "postgresql-17-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pgfincore" "1.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-2.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgfincore_17` | 1.3.1 | `el8.x86_64` | pgdg | 24.5 KiB | [pgfincore_17-1.3.1-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgfincore_17-1.3.1-3PGDG.rhel8.x86_64.rpm) |
| `pgfincore_17` | 1.3.1 | `el8.aarch64` | pgdg | 24.2 KiB | [pgfincore_17-1.3.1-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgfincore_17-1.3.1-3PGDG.rhel8.aarch64.rpm) |
| `pgfincore_17` | 1.3.1 | `el9.x86_64` | pgdg | 23.8 KiB | [pgfincore_17-1.3.1-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgfincore_17-1.3.1-3PGDG.rhel9.x86_64.rpm) |
| `pgfincore_17` | 1.3.1 | `el9.aarch64` | pgdg | 23.2 KiB | [pgfincore_17-1.3.1-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgfincore_17-1.3.1-3PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-pgfincore` | 1.3.1 | `d12.aarch64` | pgdg | 32.1 KiB | [postgresql-17-pgfincore_1.3.1-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-2.pgdg120+1_arm64.deb) |
| `postgresql-17-pgfincore` | 1.3.1 | `d12.x86_64` | pgdg | 32.5 KiB | [postgresql-17-pgfincore_1.3.1-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-2.pgdg120+1_amd64.deb) |
| `postgresql-17-pgfincore` | 1.3.1 | `u22.x86_64` | pgdg | 31.6 KiB | [postgresql-17-pgfincore_1.3.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgfincore` | 1.3.1 | `u22.aarch64` | pgdg | 31.1 KiB | [postgresql-17-pgfincore_1.3.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgfincore` | 1.3.1 | `u24.x86_64` | pgdg | 26.8 KiB | [postgresql-17-pgfincore_1.3.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgfincore` | 1.3.1 | `u24.aarch64` | pgdg | 26.2 KiB | [postgresql-17-pgfincore_1.3.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgfincore_16` | 1.3.1 | `el8.aarch64` | pgdg | 24.0 KiB | [pgfincore_16-1.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgfincore_16-1.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pgfincore_16` | 1.3.1 | `el8.x86_64` | pgdg | 24.3 KiB | [pgfincore_16-1.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgfincore_16-1.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pgfincore_16` | 1.3.1 | `el9.aarch64` | pgdg | 22.8 KiB | [pgfincore_16-1.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgfincore_16-1.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pgfincore_16` | 1.3.1 | `el9.x86_64` | pgdg | 23.5 KiB | [pgfincore_16-1.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgfincore_16-1.3.1-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-16-pgfincore` | 1.3.1 | `d12.x86_64` | pgdg | 32.0 KiB | [postgresql-16-pgfincore_1.3.1-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-2.pgdg120+1_amd64.deb) |
| `postgresql-16-pgfincore` | 1.3.1 | `d12.aarch64` | pgdg | 31.7 KiB | [postgresql-16-pgfincore_1.3.1-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-2.pgdg120+1_arm64.deb) |
| `postgresql-16-pgfincore` | 1.3.1 | `u22.aarch64` | pgdg | 30.6 KiB | [postgresql-16-pgfincore_1.3.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgfincore` | 1.3.1 | `u22.x86_64` | pgdg | 31.1 KiB | [postgresql-16-pgfincore_1.3.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgfincore` | 1.3.1 | `u24.x86_64` | pgdg | 26.8 KiB | [postgresql-16-pgfincore_1.3.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgfincore` | 1.3.1 | `u24.aarch64` | pgdg | 26.2 KiB | [postgresql-16-pgfincore_1.3.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgfincore_15` | 1.3.1 | `el8.x86_64` | pgdg | 24.3 KiB | [pgfincore_15-1.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgfincore_15-1.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pgfincore_15` | 1.3.1 | `el8.aarch64` | pgdg | 24.0 KiB | [pgfincore_15-1.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgfincore_15-1.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pgfincore_15` | 1.2.4 | `el8.aarch64` | pgdg | 23.7 KiB | [pgfincore_15-1.2.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgfincore_15-1.2.4-1.rhel8.aarch64.rpm) |
| `pgfincore_15` | 1.2.4 | `el8.x86_64` | pgdg | 24.0 KiB | [pgfincore_15-1.2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgfincore_15-1.2.4-1.rhel8.x86_64.rpm) |
| `pgfincore_15` | 1.2.4 | `el9.aarch64` | pgdg | 23.1 KiB | [pgfincore_15-1.2.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgfincore_15-1.2.4-1.rhel9.aarch64.rpm) |
| `pgfincore_15` | 1.2.4 | `el9.x86_64` | pgdg | 23.7 KiB | [pgfincore_15-1.2.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgfincore_15-1.2.4-1.rhel9.x86_64.rpm) |
| `postgresql-15-pgfincore` | 1.3.1 | `d12.x86_64` | pgdg | 32.0 KiB | [postgresql-15-pgfincore_1.3.1-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-2.pgdg120+1_amd64.deb) |
| `postgresql-15-pgfincore` | 1.3.1 | `d12.aarch64` | pgdg | 31.7 KiB | [postgresql-15-pgfincore_1.3.1-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-2.pgdg120+1_arm64.deb) |
| `postgresql-15-pgfincore` | 1.3.1 | `u22.x86_64` | pgdg | 31.2 KiB | [postgresql-15-pgfincore_1.3.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgfincore` | 1.3.1 | `u22.aarch64` | pgdg | 30.6 KiB | [postgresql-15-pgfincore_1.3.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgfincore` | 1.3.1 | `u24.aarch64` | pgdg | 26.2 KiB | [postgresql-15-pgfincore_1.3.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-2.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgfincore` | 1.3.1 | `u24.x86_64` | pgdg | 26.8 KiB | [postgresql-15-pgfincore_1.3.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-2.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgfincore_14` | 1.3.1 | `el8.x86_64` | pgdg | 24.3 KiB | [pgfincore_14-1.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgfincore_14-1.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pgfincore_14` | 1.3.1 | `el8.aarch64` | pgdg | 24.0 KiB | [pgfincore_14-1.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgfincore_14-1.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pgfincore_14` | 1.2.4 | `el8.aarch64` | pgdg | 23.7 KiB | [pgfincore_14-1.2.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgfincore_14-1.2.4-1.rhel8.aarch64.rpm) |
| `pgfincore_14` | 1.2.4 | `el8.x86_64` | pgdg | 24.0 KiB | [pgfincore_14-1.2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgfincore_14-1.2.4-1.rhel8.x86_64.rpm) |
| `pgfincore_14` | 1.2.2 | `el8.x86_64` | pgdg | 41.0 KiB | [pgfincore_14-1.2.2-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgfincore_14-1.2.2-3.rhel8.x86_64.rpm) |
| `pgfincore_14` | 1.2.4 | `el9.aarch64` | pgdg | 23.1 KiB | [pgfincore_14-1.2.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgfincore_14-1.2.4-1.rhel9.aarch64.rpm) |
| `pgfincore_14` | 1.2.4 | `el9.x86_64` | pgdg | 23.7 KiB | [pgfincore_14-1.2.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgfincore_14-1.2.4-1.rhel9.x86_64.rpm) |
| `postgresql-14-pgfincore` | 1.3.1 | `d12.aarch64` | pgdg | 31.6 KiB | [postgresql-14-pgfincore_1.3.1-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-2.pgdg120+1_arm64.deb) |
| `postgresql-14-pgfincore` | 1.3.1 | `d12.x86_64` | pgdg | 32.0 KiB | [postgresql-14-pgfincore_1.3.1-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-2.pgdg120+1_amd64.deb) |
| `postgresql-14-pgfincore` | 1.3.1 | `u22.x86_64` | pgdg | 31.1 KiB | [postgresql-14-pgfincore_1.3.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgfincore` | 1.3.1 | `u22.aarch64` | pgdg | 30.6 KiB | [postgresql-14-pgfincore_1.3.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgfincore` | 1.3.1 | `u24.aarch64` | pgdg | 26.2 KiB | [postgresql-14-pgfincore_1.3.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-2.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgfincore` | 1.3.1 | `u24.x86_64` | pgdg | 26.7 KiB | [postgresql-14-pgfincore_1.3.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-2.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgfincore_13` | 1.3.1 | `el8.aarch64` | pgdg | 24.0 KiB | [pgfincore_13-1.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgfincore_13-1.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pgfincore_13` | 1.3.1 | `el8.x86_64` | pgdg | 24.1 KiB | [pgfincore_13-1.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgfincore_13-1.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pgfincore_13` | 1.2.4 | `el8.aarch64` | pgdg | 23.7 KiB | [pgfincore_13-1.2.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgfincore_13-1.2.4-1.rhel8.aarch64.rpm) |
| `pgfincore_13` | 1.2.4 | `el8.x86_64` | pgdg | 23.8 KiB | [pgfincore_13-1.2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgfincore_13-1.2.4-1.rhel8.x86_64.rpm) |
| `pgfincore_13` | 1.2.4 | `el9.aarch64` | pgdg | 23.1 KiB | [pgfincore_13-1.2.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgfincore_13-1.2.4-1.rhel9.aarch64.rpm) |
| `pgfincore_13` | 1.2.4 | `el9.x86_64` | pgdg | 23.7 KiB | [pgfincore_13-1.2.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgfincore_13-1.2.4-1.rhel9.x86_64.rpm) |
| `postgresql-13-pgfincore` | 1.3.1 | `d12.x86_64` | pgdg | 31.8 KiB | [postgresql-13-pgfincore_1.3.1-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-13-pgfincore_1.3.1-2.pgdg120+1_amd64.deb) |
| `postgresql-13-pgfincore` | 1.3.1 | `d12.aarch64` | pgdg | 31.4 KiB | [postgresql-13-pgfincore_1.3.1-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-13-pgfincore_1.3.1-2.pgdg120+1_arm64.deb) |
| `postgresql-13-pgfincore` | 1.3.1 | `u22.x86_64` | pgdg | 30.9 KiB | [postgresql-13-pgfincore_1.3.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-13-pgfincore_1.3.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pgfincore` | 1.3.1 | `u22.aarch64` | pgdg | 30.3 KiB | [postgresql-13-pgfincore_1.3.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-13-pgfincore_1.3.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pgfincore` | 1.3.1 | `u24.x86_64` | pgdg | 26.7 KiB | [postgresql-13-pgfincore_1.3.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-13-pgfincore_1.3.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pgfincore` | 1.3.1 | `u24.aarch64` | pgdg | 26.1 KiB | [postgresql-13-pgfincore_1.3.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-13-pgfincore_1.3.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/klando/pgfincore" title="Repository" icon="github" subtitle="github.com/klando/pgfincore" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgfincore; # install by extension name, for the current active PG version
pig ext install pgfincore; # install via package alias, for the active PG version
pig ext install pgfincore -v 17;   # install for PG 17
pig ext install pgfincore -v 16;   # install for PG 16
pig ext install pgfincore -v 15;   # install for PG 15
pig ext install pgfincore -v 14;   # install for PG 14
pig ext install pgfincore -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgfincore;
```

