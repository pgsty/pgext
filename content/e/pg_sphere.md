---
title: "pg_sphere"
linkTitle: "pg_sphere"
description: "spherical objects with useful functions, operators and index support"
weight: 3590
categories: ["Type"]
width: full
---

spherical objects with useful functions, operators and index support

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3590** | {{< badge content="pg_sphere" link="https://github.com/postgrespro/pgsphere" >}} | {{< ext "pg_sphere" "pgsphere" >}} | `1.5.1` | {{< category "TYPE" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "postgis" >}} {{< ext "q3c" >}} {{< ext "earthdistance" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_sphere" >}} | `1.5.1` | {{< badge content="18" color="red" alt="pgsphere_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pgsphere_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pg_sphere" >}} | `1.5.1` | {{< badge content="18" color="red" alt="postgresql-18-pgsphere" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pgsphere` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pgsphere_18" >}}     | {{< pkg "pgsphere_17" "1.5.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_17-1.5.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pgsphere_16" "1.5.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_16-1.5.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pgsphere_15" "1.5.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_15-1.5.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pgsphere_14" "1.5.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_14-1.5.1-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pgsphere_18" >}}     | {{< pkg "pgsphere_17" "1.5.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_17-1.5.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pgsphere_16" "1.5.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_16-1.5.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pgsphere_15" "1.5.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_15-1.5.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pgsphere_14" "1.5.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_14-1.5.1-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pgsphere_18" >}}     | {{< pkg "pgsphere_17" "1.5.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_17-1.5.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pgsphere_16" "1.5.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_16-1.5.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pgsphere_15" "1.5.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_15-1.5.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pgsphere_14" "1.5.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_14-1.5.1-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pgsphere_18" >}}     | {{< pkg "pgsphere_17" "1.5.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_17-1.5.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pgsphere_16" "1.5.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_16-1.5.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pgsphere_15" "1.5.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_15-1.5.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pgsphere_14" "1.5.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_14-1.5.1-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pgsphere" >}}     | {{< pkg "postgresql-17-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.1-2.pgdg120+1_amd64.deb" >}} | {{< pkg "postgresql-16-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.1-2.pgdg120+1_amd64.deb" >}} | {{< pkg "postgresql-15-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.1-2.pgdg120+1_amd64.deb" >}} | {{< pkg "postgresql-14-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.1-2.pgdg120+1_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pgsphere" >}}     | {{< pkg "postgresql-17-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.1-2.pgdg120+1_arm64.deb" >}} | {{< pkg "postgresql-16-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.1-2.pgdg120+1_arm64.deb" >}} | {{< pkg "postgresql-15-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.1-2.pgdg120+1_arm64.deb" >}} | {{< pkg "postgresql-14-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.1-2.pgdg120+1_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pgsphere" >}}     | {{< pkg "postgresql-17-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.1-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.1-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.1-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.1-2.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pgsphere" >}}     | {{< pkg "postgresql-17-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.1-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.1-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.1-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.1-2.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pgsphere" >}}     | {{< pkg "postgresql-17-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.1-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.1-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.1-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.1-2.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pgsphere" >}}     | {{< pkg "postgresql-17-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.1-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.1-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.1-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pgsphere" "1.5.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.1-2.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsphere_17` | 1.5.1 | `el8.x86_64` | pigsty | 105.8 KiB | [pgsphere_17-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_17-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsphere_17` | 1.5.1 | `el8.aarch64` | pigsty | 103.4 KiB | [pgsphere_17-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_17-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsphere_17` | 1.5.1 | `el9.aarch64` | pigsty | 110.2 KiB | [pgsphere_17-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_17-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsphere_17` | 1.5.1 | `el9.x86_64` | pigsty | 112.1 KiB | [pgsphere_17-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_17-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pgsphere` | 1.5.1 | `d12.x86_64` | pgdg | 427.6 KiB | [postgresql-17-pgsphere_1.5.1-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.1-2.pgdg120+1_amd64.deb) |
| `postgresql-17-pgsphere` | 1.5.1 | `d12.aarch64` | pgdg | 422.4 KiB | [postgresql-17-pgsphere_1.5.1-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.1-2.pgdg120+1_arm64.deb) |
| `postgresql-17-pgsphere` | 1.5.1 | `u22.x86_64` | pgdg | 433.3 KiB | [postgresql-17-pgsphere_1.5.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgsphere` | 1.5.1 | `u22.aarch64` | pgdg | 426.8 KiB | [postgresql-17-pgsphere_1.5.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgsphere` | 1.5.1 | `u24.x86_64` | pgdg | 408.1 KiB | [postgresql-17-pgsphere_1.5.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgsphere` | 1.5.1 | `u24.aarch64` | pgdg | 403.5 KiB | [postgresql-17-pgsphere_1.5.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsphere_16` | 1.5.1 | `el8.x86_64` | pigsty | 105.8 KiB | [pgsphere_16-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_16-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsphere_16` | 1.5.1 | `el8.aarch64` | pigsty | 103.4 KiB | [pgsphere_16-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_16-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsphere_16` | 1.5.1 | `el9.x86_64` | pigsty | 112.1 KiB | [pgsphere_16-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_16-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsphere_16` | 1.5.1 | `el9.aarch64` | pigsty | 110.2 KiB | [pgsphere_16-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_16-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pgsphere` | 1.5.1 | `d12.x86_64` | pgdg | 427.4 KiB | [postgresql-16-pgsphere_1.5.1-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.1-2.pgdg120+1_amd64.deb) |
| `postgresql-16-pgsphere` | 1.5.1 | `d12.aarch64` | pgdg | 422.3 KiB | [postgresql-16-pgsphere_1.5.1-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.1-2.pgdg120+1_arm64.deb) |
| `postgresql-16-pgsphere` | 1.5.1 | `u22.aarch64` | pgdg | 426.6 KiB | [postgresql-16-pgsphere_1.5.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgsphere` | 1.5.1 | `u22.x86_64` | pgdg | 433.4 KiB | [postgresql-16-pgsphere_1.5.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgsphere` | 1.5.1 | `u24.x86_64` | pgdg | 408.6 KiB | [postgresql-16-pgsphere_1.5.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgsphere` | 1.5.1 | `u24.aarch64` | pgdg | 403.4 KiB | [postgresql-16-pgsphere_1.5.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsphere_15` | 1.5.1 | `el8.x86_64` | pigsty | 108.1 KiB | [pgsphere_15-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_15-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsphere_15` | 1.5.1 | `el8.aarch64` | pigsty | 105.3 KiB | [pgsphere_15-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_15-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsphere_15` | 1.5.1 | `el9.x86_64` | pigsty | 108.8 KiB | [pgsphere_15-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_15-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsphere_15` | 1.5.1 | `el9.aarch64` | pigsty | 108.7 KiB | [pgsphere_15-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_15-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pgsphere` | 1.5.1 | `d12.aarch64` | pgdg | 424.6 KiB | [postgresql-15-pgsphere_1.5.1-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.1-2.pgdg120+1_arm64.deb) |
| `postgresql-15-pgsphere` | 1.5.1 | `d12.x86_64` | pgdg | 429.8 KiB | [postgresql-15-pgsphere_1.5.1-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.1-2.pgdg120+1_amd64.deb) |
| `postgresql-15-pgsphere` | 1.5.1 | `u22.aarch64` | pgdg | 427.7 KiB | [postgresql-15-pgsphere_1.5.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgsphere` | 1.5.1 | `u22.x86_64` | pgdg | 433.3 KiB | [postgresql-15-pgsphere_1.5.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgsphere` | 1.5.1 | `u24.x86_64` | pgdg | 407.3 KiB | [postgresql-15-pgsphere_1.5.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgsphere` | 1.5.1 | `u24.aarch64` | pgdg | 404.0 KiB | [postgresql-15-pgsphere_1.5.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsphere_14` | 1.5.1 | `el8.x86_64` | pigsty | 108.2 KiB | [pgsphere_14-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_14-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsphere_14` | 1.5.1 | `el8.aarch64` | pigsty | 105.3 KiB | [pgsphere_14-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_14-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsphere_14` | 1.5.1 | `el9.x86_64` | pigsty | 108.8 KiB | [pgsphere_14-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_14-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsphere_14` | 1.5.1 | `el9.aarch64` | pigsty | 108.7 KiB | [pgsphere_14-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_14-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pgsphere` | 1.5.1 | `d12.x86_64` | pgdg | 430.0 KiB | [postgresql-14-pgsphere_1.5.1-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.1-2.pgdg120+1_amd64.deb) |
| `postgresql-14-pgsphere` | 1.5.1 | `d12.aarch64` | pgdg | 424.7 KiB | [postgresql-14-pgsphere_1.5.1-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.1-2.pgdg120+1_arm64.deb) |
| `postgresql-14-pgsphere` | 1.5.1 | `u22.x86_64` | pgdg | 433.6 KiB | [postgresql-14-pgsphere_1.5.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgsphere` | 1.5.1 | `u22.aarch64` | pgdg | 427.6 KiB | [postgresql-14-pgsphere_1.5.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgsphere` | 1.5.1 | `u24.x86_64` | pgdg | 407.2 KiB | [postgresql-14-pgsphere_1.5.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgsphere` | 1.5.1 | `u24.aarch64` | pgdg | 404.2 KiB | [postgresql-14-pgsphere_1.5.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsphere_13` | 1.5.1 | `el8.aarch64` | pigsty | 105.3 KiB | [pgsphere_13-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_13-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsphere_13` | 1.5.1 | `el8.x86_64` | pigsty | 108.0 KiB | [pgsphere_13-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_13-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsphere_13` | 1.5.1 | `el9.aarch64` | pigsty | 108.6 KiB | [pgsphere_13-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_13-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsphere_13` | 1.5.1 | `el9.x86_64` | pigsty | 108.9 KiB | [pgsphere_13-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_13-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pgsphere` | 1.5.1 | `d12.aarch64` | pgdg | 423.4 KiB | [postgresql-13-pgsphere_1.5.1-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-13-pgsphere_1.5.1-2.pgdg120+1_arm64.deb) |
| `postgresql-13-pgsphere` | 1.5.1 | `d12.x86_64` | pgdg | 428.7 KiB | [postgresql-13-pgsphere_1.5.1-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-13-pgsphere_1.5.1-2.pgdg120+1_amd64.deb) |
| `postgresql-13-pgsphere` | 1.5.1 | `u22.aarch64` | pgdg | 427.3 KiB | [postgresql-13-pgsphere_1.5.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-13-pgsphere_1.5.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pgsphere` | 1.5.1 | `u22.x86_64` | pgdg | 432.5 KiB | [postgresql-13-pgsphere_1.5.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-13-pgsphere_1.5.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pgsphere` | 1.5.1 | `u24.aarch64` | pgdg | 403.7 KiB | [postgresql-13-pgsphere_1.5.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-13-pgsphere_1.5.1-2.pgdg24.04+1_arm64.deb) |
| `postgresql-13-pgsphere` | 1.5.1 | `u24.x86_64` | pgdg | 406.6 KiB | [postgresql-13-pgsphere_1.5.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-13-pgsphere_1.5.1-2.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/postgrespro/pgsphere" title="Repository" icon="github" subtitle="github.com/postgrespro/pgsphere" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pgsphere-1.5.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_sphere; # get pg_sphere source code
pig build dep pg_sphere; # install build dependencies
pig build pkg pg_sphere; # build extension rpm or deb
pig build ext pg_sphere; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_sphere; # install by extension name, for the current active PG version
pig ext install pgsphere; # install via package alias, for the active PG version
pig ext install pg_sphere -v 18;   # install for PG 18
pig ext install pg_sphere -v 17;   # install for PG 17
pig ext install pg_sphere -v 16;   # install for PG 16
pig ext install pg_sphere -v 15;   # install for PG 15
pig ext install pg_sphere -v 14;   # install for PG 14
pig ext install pg_sphere -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_sphere;
```

