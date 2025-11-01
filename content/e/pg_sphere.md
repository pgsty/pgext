---
title: "pg_sphere"
linkTitle: "pg_sphere"
description: "spherical objects with useful functions, operators and index support"
weight: 3590
categories: ["TYPE"]
width: full
---

spherical objects with useful functions, operators and index support


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3590** | {{< badge content="pg_sphere" link="https://github.com/postgrespro/pgsphere" >}} | {{< ext "pg_sphere" "pgsphere" >}} | `1.5.2` | {{< category "TYPE" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "postgis" >}} {{< ext "q3c" >}} {{< ext "earthdistance" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_sphere" >}} | `1.5.2` | {{< bg "18" "pgsphere_18*" "green" >}} {{< bg "17" "pgsphere_17*" "green" >}} {{< bg "16" "pgsphere_16*" "green" >}} {{< bg "15" "pgsphere_15*" "green" >}} {{< bg "14" "pgsphere_14*" "green" >}} {{< bg "13" "pgsphere_13*" "green" >}} | `pgsphere_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pg_sphere" >}} | `1.5.2` | {{< bg "18" "postgresql-18-pgsphere" "green" >}} {{< bg "17" "postgresql-17-pgsphere" "green" >}} {{< bg "16" "postgresql-16-pgsphere" "green" >}} {{< bg "15" "postgresql-15-pgsphere" "green" >}} {{< bg "14" "postgresql-14-pgsphere" "green" >}} {{< bg "13" "postgresql-13-pgsphere" "green" >}} | `postgresql-$v-pgsphere` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 1.5.2" "pgsphere_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_13 : AVAIL 2" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 1.5.2" "pgsphere_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_13 : AVAIL 2" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 1.5.2" "pgsphere_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_13 : AVAIL 2" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 1.5.2" "pgsphere_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_13 : AVAIL 2" "green" >}} |
|    `el10.x86_64`    | {{< bg "PIGSTY 1.5.2" "pgsphere_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_13 : AVAIL 1" "green" >}} |
|    `el10.aarch64`    | {{< bg "PIGSTY 1.5.2" "pgsphere_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "pgsphere_13 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 1.5.2" "postgresql-18-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-17-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-16-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-15-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-14-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-13-pgsphere : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 1.5.2" "postgresql-18-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-17-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-16-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-15-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-14-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-13-pgsphere : AVAIL 1" "blue" >}} |
|    `d13.x86_64`    | {{< bg "PGDG 1.5.2" "postgresql-18-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-17-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-16-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-15-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-14-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-13-pgsphere : AVAIL 1" "blue" >}} |
|    `d13.aarch64`    | {{< bg "PGDG 1.5.2" "postgresql-18-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-17-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-16-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-15-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-14-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-13-pgsphere : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 1.5.2" "postgresql-18-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-17-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-16-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-15-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-14-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-13-pgsphere : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 1.5.2" "postgresql-18-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-17-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-16-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-15-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-14-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-13-pgsphere : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 1.5.2" "postgresql-18-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-17-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-16-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-15-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-14-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-13-pgsphere : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 1.5.2" "postgresql-18-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-17-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-16-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-15-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-14-pgsphere : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-13-pgsphere : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsphere_18` | 1.5.2 | `el8.x86_64` | pigsty | 125.2 KiB | [pgsphere_18-1.5.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_18-1.5.2-1PIGSTY.el8.x86_64.rpm) |
| `pgsphere_18` | 1.5.2 | `el8.aarch64` | pigsty | 122.2 KiB | [pgsphere_18-1.5.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_18-1.5.2-1PIGSTY.el8.aarch64.rpm) |
| `pgsphere_18` | 1.5.2 | `el9.x86_64` | pigsty | 118.2 KiB | [pgsphere_18-1.5.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_18-1.5.2-1PIGSTY.el9.x86_64.rpm) |
| `pgsphere_18` | 1.5.2 | `el9.aarch64` | pigsty | 116.7 KiB | [pgsphere_18-1.5.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_18-1.5.2-1PIGSTY.el9.aarch64.rpm) |
| `pgsphere_18` | 1.5.2 | `el10.x86_64` | pigsty | 119.3 KiB | [pgsphere_18-1.5.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsphere_18-1.5.2-1PIGSTY.el10.x86_64.rpm) |
| `pgsphere_18` | 1.5.2 | `el10.aarch64` | pigsty | 119.5 KiB | [pgsphere_18-1.5.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsphere_18-1.5.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgsphere` | 1.5.2 | `d12.x86_64` | pgdg | 405.0 KiB | [postgresql-18-pgsphere_1.5.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-18-pgsphere_1.5.2-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pgsphere` | 1.5.2 | `d12.aarch64` | pgdg | 400.2 KiB | [postgresql-18-pgsphere_1.5.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-18-pgsphere_1.5.2-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pgsphere` | 1.5.2 | `d13.x86_64` | pgdg | 405.1 KiB | [postgresql-18-pgsphere_1.5.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-18-pgsphere_1.5.2-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pgsphere` | 1.5.2 | `d13.aarch64` | pgdg | 402.7 KiB | [postgresql-18-pgsphere_1.5.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-18-pgsphere_1.5.2-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pgsphere` | 1.5.2 | `u22.x86_64` | pgdg | 413.8 KiB | [postgresql-18-pgsphere_1.5.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-18-pgsphere_1.5.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgsphere` | 1.5.2 | `u22.aarch64` | pgdg | 407.4 KiB | [postgresql-18-pgsphere_1.5.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-18-pgsphere_1.5.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgsphere` | 1.5.2 | `u24.x86_64` | pgdg | 406.4 KiB | [postgresql-18-pgsphere_1.5.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-18-pgsphere_1.5.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgsphere` | 1.5.2 | `u24.aarch64` | pgdg | 401.2 KiB | [postgresql-18-pgsphere_1.5.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-18-pgsphere_1.5.2-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsphere_17` | 1.5.2 | `el8.x86_64` | pigsty | 125.2 KiB | [pgsphere_17-1.5.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_17-1.5.2-1PIGSTY.el8.x86_64.rpm) |
| `pgsphere_17` | 1.5.1 | `el8.x86_64` | pigsty | 105.8 KiB | [pgsphere_17-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_17-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsphere_17` | 1.5.2 | `el8.aarch64` | pigsty | 122.2 KiB | [pgsphere_17-1.5.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_17-1.5.2-1PIGSTY.el8.aarch64.rpm) |
| `pgsphere_17` | 1.5.1 | `el8.aarch64` | pigsty | 103.4 KiB | [pgsphere_17-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_17-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsphere_17` | 1.5.2 | `el9.x86_64` | pigsty | 118.2 KiB | [pgsphere_17-1.5.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_17-1.5.2-1PIGSTY.el9.x86_64.rpm) |
| `pgsphere_17` | 1.5.1 | `el9.x86_64` | pigsty | 112.1 KiB | [pgsphere_17-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_17-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsphere_17` | 1.5.2 | `el9.aarch64` | pigsty | 116.7 KiB | [pgsphere_17-1.5.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_17-1.5.2-1PIGSTY.el9.aarch64.rpm) |
| `pgsphere_17` | 1.5.1 | `el9.aarch64` | pigsty | 110.2 KiB | [pgsphere_17-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_17-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsphere_17` | 1.5.2 | `el10.x86_64` | pigsty | 119.3 KiB | [pgsphere_17-1.5.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsphere_17-1.5.2-1PIGSTY.el10.x86_64.rpm) |
| `pgsphere_17` | 1.5.2 | `el10.aarch64` | pigsty | 119.6 KiB | [pgsphere_17-1.5.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsphere_17-1.5.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgsphere` | 1.5.2 | `d12.x86_64` | pgdg | 404.8 KiB | [postgresql-17-pgsphere_1.5.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.2-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pgsphere` | 1.5.2 | `d12.aarch64` | pgdg | 400.1 KiB | [postgresql-17-pgsphere_1.5.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.2-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pgsphere` | 1.5.2 | `d13.x86_64` | pgdg | 405.5 KiB | [postgresql-17-pgsphere_1.5.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.2-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pgsphere` | 1.5.2 | `d13.aarch64` | pgdg | 402.5 KiB | [postgresql-17-pgsphere_1.5.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.2-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pgsphere` | 1.5.2 | `u22.x86_64` | pgdg | 434.0 KiB | [postgresql-17-pgsphere_1.5.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgsphere` | 1.5.2 | `u22.aarch64` | pgdg | 427.1 KiB | [postgresql-17-pgsphere_1.5.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgsphere` | 1.5.2 | `u24.x86_64` | pgdg | 406.3 KiB | [postgresql-17-pgsphere_1.5.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgsphere` | 1.5.2 | `u24.aarch64` | pgdg | 401.0 KiB | [postgresql-17-pgsphere_1.5.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-17-pgsphere_1.5.2-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsphere_16` | 1.5.2 | `el8.x86_64` | pigsty | 125.3 KiB | [pgsphere_16-1.5.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_16-1.5.2-1PIGSTY.el8.x86_64.rpm) |
| `pgsphere_16` | 1.5.1 | `el8.x86_64` | pigsty | 105.8 KiB | [pgsphere_16-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_16-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsphere_16` | 1.5.2 | `el8.aarch64` | pigsty | 122.2 KiB | [pgsphere_16-1.5.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_16-1.5.2-1PIGSTY.el8.aarch64.rpm) |
| `pgsphere_16` | 1.5.1 | `el8.aarch64` | pigsty | 103.4 KiB | [pgsphere_16-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_16-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsphere_16` | 1.5.2 | `el9.x86_64` | pigsty | 118.2 KiB | [pgsphere_16-1.5.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_16-1.5.2-1PIGSTY.el9.x86_64.rpm) |
| `pgsphere_16` | 1.5.1 | `el9.x86_64` | pigsty | 112.1 KiB | [pgsphere_16-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_16-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsphere_16` | 1.5.2 | `el9.aarch64` | pigsty | 116.7 KiB | [pgsphere_16-1.5.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_16-1.5.2-1PIGSTY.el9.aarch64.rpm) |
| `pgsphere_16` | 1.5.1 | `el9.aarch64` | pigsty | 110.2 KiB | [pgsphere_16-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_16-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsphere_16` | 1.5.2 | `el10.x86_64` | pigsty | 119.3 KiB | [pgsphere_16-1.5.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsphere_16-1.5.2-1PIGSTY.el10.x86_64.rpm) |
| `pgsphere_16` | 1.5.2 | `el10.aarch64` | pigsty | 119.5 KiB | [pgsphere_16-1.5.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsphere_16-1.5.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgsphere` | 1.5.2 | `d12.x86_64` | pgdg | 404.8 KiB | [postgresql-16-pgsphere_1.5.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.2-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pgsphere` | 1.5.2 | `d12.aarch64` | pgdg | 400.2 KiB | [postgresql-16-pgsphere_1.5.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.2-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pgsphere` | 1.5.2 | `d13.x86_64` | pgdg | 405.4 KiB | [postgresql-16-pgsphere_1.5.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.2-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pgsphere` | 1.5.2 | `d13.aarch64` | pgdg | 402.6 KiB | [postgresql-16-pgsphere_1.5.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.2-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pgsphere` | 1.5.2 | `u22.x86_64` | pgdg | 433.6 KiB | [postgresql-16-pgsphere_1.5.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgsphere` | 1.5.2 | `u22.aarch64` | pgdg | 427.1 KiB | [postgresql-16-pgsphere_1.5.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgsphere` | 1.5.2 | `u24.x86_64` | pgdg | 406.4 KiB | [postgresql-16-pgsphere_1.5.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgsphere` | 1.5.2 | `u24.aarch64` | pgdg | 401.0 KiB | [postgresql-16-pgsphere_1.5.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-16-pgsphere_1.5.2-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsphere_15` | 1.5.2 | `el8.x86_64` | pigsty | 127.2 KiB | [pgsphere_15-1.5.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_15-1.5.2-1PIGSTY.el8.x86_64.rpm) |
| `pgsphere_15` | 1.5.1 | `el8.x86_64` | pigsty | 108.1 KiB | [pgsphere_15-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_15-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsphere_15` | 1.5.2 | `el8.aarch64` | pigsty | 124.1 KiB | [pgsphere_15-1.5.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_15-1.5.2-1PIGSTY.el8.aarch64.rpm) |
| `pgsphere_15` | 1.5.1 | `el8.aarch64` | pigsty | 105.3 KiB | [pgsphere_15-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_15-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsphere_15` | 1.5.2 | `el9.x86_64` | pigsty | 114.7 KiB | [pgsphere_15-1.5.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_15-1.5.2-1PIGSTY.el9.x86_64.rpm) |
| `pgsphere_15` | 1.5.1 | `el9.x86_64` | pigsty | 108.8 KiB | [pgsphere_15-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_15-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsphere_15` | 1.5.2 | `el9.aarch64` | pigsty | 114.8 KiB | [pgsphere_15-1.5.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_15-1.5.2-1PIGSTY.el9.aarch64.rpm) |
| `pgsphere_15` | 1.5.1 | `el9.aarch64` | pigsty | 108.7 KiB | [pgsphere_15-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_15-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsphere_15` | 1.5.2 | `el10.x86_64` | pigsty | 114.5 KiB | [pgsphere_15-1.5.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsphere_15-1.5.2-1PIGSTY.el10.x86_64.rpm) |
| `pgsphere_15` | 1.5.2 | `el10.aarch64` | pigsty | 116.8 KiB | [pgsphere_15-1.5.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsphere_15-1.5.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgsphere` | 1.5.2 | `d12.x86_64` | pgdg | 405.9 KiB | [postgresql-15-pgsphere_1.5.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.2-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pgsphere` | 1.5.2 | `d12.aarch64` | pgdg | 401.2 KiB | [postgresql-15-pgsphere_1.5.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.2-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pgsphere` | 1.5.2 | `d13.x86_64` | pgdg | 406.2 KiB | [postgresql-15-pgsphere_1.5.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.2-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pgsphere` | 1.5.2 | `d13.aarch64` | pgdg | 403.3 KiB | [postgresql-15-pgsphere_1.5.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.2-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pgsphere` | 1.5.2 | `u22.x86_64` | pgdg | 434.2 KiB | [postgresql-15-pgsphere_1.5.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgsphere` | 1.5.2 | `u22.aarch64` | pgdg | 428.2 KiB | [postgresql-15-pgsphere_1.5.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgsphere` | 1.5.2 | `u24.x86_64` | pgdg | 404.9 KiB | [postgresql-15-pgsphere_1.5.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgsphere` | 1.5.2 | `u24.aarch64` | pgdg | 402.1 KiB | [postgresql-15-pgsphere_1.5.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-15-pgsphere_1.5.2-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsphere_14` | 1.5.2 | `el8.x86_64` | pigsty | 127.1 KiB | [pgsphere_14-1.5.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_14-1.5.2-1PIGSTY.el8.x86_64.rpm) |
| `pgsphere_14` | 1.5.1 | `el8.x86_64` | pigsty | 108.2 KiB | [pgsphere_14-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_14-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsphere_14` | 1.5.2 | `el8.aarch64` | pigsty | 124.1 KiB | [pgsphere_14-1.5.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_14-1.5.2-1PIGSTY.el8.aarch64.rpm) |
| `pgsphere_14` | 1.5.1 | `el8.aarch64` | pigsty | 105.3 KiB | [pgsphere_14-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_14-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsphere_14` | 1.5.2 | `el9.x86_64` | pigsty | 114.6 KiB | [pgsphere_14-1.5.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_14-1.5.2-1PIGSTY.el9.x86_64.rpm) |
| `pgsphere_14` | 1.5.1 | `el9.x86_64` | pigsty | 108.8 KiB | [pgsphere_14-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_14-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsphere_14` | 1.5.2 | `el9.aarch64` | pigsty | 114.8 KiB | [pgsphere_14-1.5.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_14-1.5.2-1PIGSTY.el9.aarch64.rpm) |
| `pgsphere_14` | 1.5.1 | `el9.aarch64` | pigsty | 108.7 KiB | [pgsphere_14-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_14-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsphere_14` | 1.5.2 | `el10.x86_64` | pigsty | 114.5 KiB | [pgsphere_14-1.5.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsphere_14-1.5.2-1PIGSTY.el10.x86_64.rpm) |
| `pgsphere_14` | 1.5.2 | `el10.aarch64` | pigsty | 116.7 KiB | [pgsphere_14-1.5.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsphere_14-1.5.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgsphere` | 1.5.2 | `d12.x86_64` | pgdg | 406.0 KiB | [postgresql-14-pgsphere_1.5.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.2-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pgsphere` | 1.5.2 | `d12.aarch64` | pgdg | 400.9 KiB | [postgresql-14-pgsphere_1.5.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.2-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pgsphere` | 1.5.2 | `d13.x86_64` | pgdg | 406.3 KiB | [postgresql-14-pgsphere_1.5.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.2-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pgsphere` | 1.5.2 | `d13.aarch64` | pgdg | 403.9 KiB | [postgresql-14-pgsphere_1.5.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.2-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pgsphere` | 1.5.2 | `u22.x86_64` | pgdg | 433.6 KiB | [postgresql-14-pgsphere_1.5.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgsphere` | 1.5.2 | `u22.aarch64` | pgdg | 428.1 KiB | [postgresql-14-pgsphere_1.5.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgsphere` | 1.5.2 | `u24.x86_64` | pgdg | 405.1 KiB | [postgresql-14-pgsphere_1.5.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgsphere` | 1.5.2 | `u24.aarch64` | pgdg | 402.0 KiB | [postgresql-14-pgsphere_1.5.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-14-pgsphere_1.5.2-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsphere_13` | 1.5.2 | `el8.x86_64` | pigsty | 126.8 KiB | [pgsphere_13-1.5.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_13-1.5.2-1PIGSTY.el8.x86_64.rpm) |
| `pgsphere_13` | 1.5.1 | `el8.x86_64` | pigsty | 108.0 KiB | [pgsphere_13-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsphere_13-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsphere_13` | 1.5.2 | `el8.aarch64` | pigsty | 124.0 KiB | [pgsphere_13-1.5.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_13-1.5.2-1PIGSTY.el8.aarch64.rpm) |
| `pgsphere_13` | 1.5.1 | `el8.aarch64` | pigsty | 105.3 KiB | [pgsphere_13-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsphere_13-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsphere_13` | 1.5.2 | `el9.x86_64` | pigsty | 114.9 KiB | [pgsphere_13-1.5.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_13-1.5.2-1PIGSTY.el9.x86_64.rpm) |
| `pgsphere_13` | 1.5.1 | `el9.x86_64` | pigsty | 108.9 KiB | [pgsphere_13-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsphere_13-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsphere_13` | 1.5.2 | `el9.aarch64` | pigsty | 114.7 KiB | [pgsphere_13-1.5.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_13-1.5.2-1PIGSTY.el9.aarch64.rpm) |
| `pgsphere_13` | 1.5.1 | `el9.aarch64` | pigsty | 108.6 KiB | [pgsphere_13-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsphere_13-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsphere_13` | 1.5.2 | `el10.x86_64` | pigsty | 114.5 KiB | [pgsphere_13-1.5.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsphere_13-1.5.2-1PIGSTY.el10.x86_64.rpm) |
| `pgsphere_13` | 1.5.2 | `el10.aarch64` | pigsty | 116.7 KiB | [pgsphere_13-1.5.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsphere_13-1.5.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pgsphere` | 1.5.2 | `d12.x86_64` | pgdg | 405.4 KiB | [postgresql-13-pgsphere_1.5.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-13-pgsphere_1.5.2-1.pgdg12+1_amd64.deb) |
| `postgresql-13-pgsphere` | 1.5.2 | `d12.aarch64` | pgdg | 400.8 KiB | [postgresql-13-pgsphere_1.5.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-13-pgsphere_1.5.2-1.pgdg12+1_arm64.deb) |
| `postgresql-13-pgsphere` | 1.5.2 | `d13.x86_64` | pgdg | 405.5 KiB | [postgresql-13-pgsphere_1.5.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-13-pgsphere_1.5.2-1.pgdg13+1_amd64.deb) |
| `postgresql-13-pgsphere` | 1.5.2 | `d13.aarch64` | pgdg | 402.7 KiB | [postgresql-13-pgsphere_1.5.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-13-pgsphere_1.5.2-1.pgdg13+1_arm64.deb) |
| `postgresql-13-pgsphere` | 1.5.2 | `u22.x86_64` | pgdg | 432.6 KiB | [postgresql-13-pgsphere_1.5.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-13-pgsphere_1.5.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pgsphere` | 1.5.2 | `u22.aarch64` | pgdg | 427.4 KiB | [postgresql-13-pgsphere_1.5.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-13-pgsphere_1.5.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pgsphere` | 1.5.2 | `u24.x86_64` | pgdg | 404.5 KiB | [postgresql-13-pgsphere_1.5.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-13-pgsphere_1.5.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pgsphere` | 1.5.2 | `u24.aarch64` | pgdg | 401.3 KiB | [postgresql-13-pgsphere_1.5.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsphere/postgresql-13-pgsphere_1.5.2-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/postgrespro/pgsphere" title="Repository" icon="github" subtitle="github.com/postgrespro/pgsphere" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgsphere-1.5.2.tar.gz" >}}
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

