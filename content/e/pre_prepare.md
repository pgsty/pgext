---
title: "pre_prepare"
linkTitle: "pre_prepare"
description: "Pre Prepare your Statement server side"
weight: 5170
categories: ["Admin"]
width: full
---

Pre Prepare your Statement server side

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5170** | {{< badge content="pre_prepare" link="https://github.com/dimitri/preprepare" >}} | {{< ext "pre_prepare" "preprepare" >}} | `0.9` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_store_plans" >}} {{< ext "auto_explain" >}} {{< ext "pg_stat_statements" >}} {{< ext "plpgsql_check" >}} {{< ext "pg_show_plans" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_hint_plan" >}} {{< ext "pgaudit" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pre_prepare" >}} | `0.9` | {{< badge content="18" color="red" alt="preprepare_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `preprepare_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pre_prepare" >}} | `0.9` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-preprepare` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "preprepare_18" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/preprepare_18-0.9-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "preprepare_17" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/preprepare_17-0.9-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "preprepare_16" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/preprepare_16-0.9-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "preprepare_15" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/preprepare_15-0.9-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "preprepare_14" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/preprepare_14-0.9-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "preprepare_18" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/preprepare_18-0.9-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "preprepare_17" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/preprepare_17-0.9-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "preprepare_16" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/preprepare_16-0.9-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "preprepare_15" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/preprepare_15-0.9-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "preprepare_14" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/preprepare_14-0.9-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "preprepare_18" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/preprepare_18-0.9-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "preprepare_17" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/preprepare_17-0.9-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "preprepare_16" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/preprepare_16-0.9-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "preprepare_15" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/preprepare_15-0.9-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "preprepare_14" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/preprepare_14-0.9-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "preprepare_18" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/preprepare_18-0.9-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "preprepare_17" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/preprepare_17-0.9-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "preprepare_16" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/preprepare_16-0.9-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "preprepare_15" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/preprepare_15-0.9-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "preprepare_14" "0.9" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/preprepare_14-0.9-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-18-preprepare_0.9-9.pgdg120+1_amd64.deb" >}} | {{< pkg "postgresql-17-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-17-preprepare_0.9-9.pgdg120+1_amd64.deb" >}} | {{< pkg "postgresql-16-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-16-preprepare_0.9-9.pgdg120+1_amd64.deb" >}} | {{< pkg "postgresql-15-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-15-preprepare_0.9-9.pgdg120+1_amd64.deb" >}} | {{< pkg "postgresql-14-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-14-preprepare_0.9-9.pgdg120+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-18-preprepare_0.9-9.pgdg120+1_arm64.deb" >}} | {{< pkg "postgresql-17-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-17-preprepare_0.9-9.pgdg120+1_arm64.deb" >}} | {{< pkg "postgresql-16-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-16-preprepare_0.9-9.pgdg120+1_arm64.deb" >}} | {{< pkg "postgresql-15-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-15-preprepare_0.9-9.pgdg120+1_arm64.deb" >}} | {{< pkg "postgresql-14-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-14-preprepare_0.9-9.pgdg120+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-18-preprepare_0.9-9.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-17-preprepare_0.9-9.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-16-preprepare_0.9-9.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-15-preprepare_0.9-9.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-14-preprepare_0.9-9.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-18-preprepare_0.9-9.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-17-preprepare_0.9-9.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-16-preprepare_0.9-9.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-15-preprepare_0.9-9.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-14-preprepare_0.9-9.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-18-preprepare_0.9-9.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-17-preprepare_0.9-9.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-16-preprepare_0.9-9.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-15-preprepare_0.9-9.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-14-preprepare_0.9-9.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-18-preprepare_0.9-9.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-17-preprepare_0.9-9.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-16-preprepare_0.9-9.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-15-preprepare_0.9-9.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-preprepare" "0.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-14-preprepare_0.9-9.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `preprepare_18` | 0.9 | `el8.aarch64` | pigsty | 13.4 KiB | [preprepare_18-0.9-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/preprepare_18-0.9-1PIGSTY.el8.aarch64.rpm) |
| `preprepare_18` | 0.9 | `el8.x86_64` | pigsty | 13.4 KiB | [preprepare_18-0.9-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/preprepare_18-0.9-1PIGSTY.el8.x86_64.rpm) |
| `preprepare_18` | 0.9 | `el9.aarch64` | pigsty | 13.1 KiB | [preprepare_18-0.9-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/preprepare_18-0.9-1PIGSTY.el9.aarch64.rpm) |
| `preprepare_18` | 0.9 | `el9.x86_64` | pigsty | 13.5 KiB | [preprepare_18-0.9-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/preprepare_18-0.9-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-18-preprepare` | 0.9 | `d12.aarch64` | pgdg | 16.7 KiB | [postgresql-18-preprepare_0.9-9.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-18-preprepare_0.9-9.pgdg120+1_arm64.deb) |
| `postgresql-18-preprepare` | 0.9 | `d12.x86_64` | pgdg | 16.7 KiB | [postgresql-18-preprepare_0.9-9.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-18-preprepare_0.9-9.pgdg120+1_amd64.deb) |
| `postgresql-18-preprepare` | 0.9 | `u22.x86_64` | pgdg | 17.2 KiB | [postgresql-18-preprepare_0.9-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-18-preprepare_0.9-9.pgdg22.04+1_amd64.deb) |
| `postgresql-18-preprepare` | 0.9 | `u22.aarch64` | pgdg | 17.2 KiB | [postgresql-18-preprepare_0.9-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-18-preprepare_0.9-9.pgdg22.04+1_arm64.deb) |
| `postgresql-18-preprepare` | 0.9 | `u24.x86_64` | pgdg | 16.3 KiB | [postgresql-18-preprepare_0.9-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-18-preprepare_0.9-9.pgdg24.04+1_amd64.deb) |
| `postgresql-18-preprepare` | 0.9 | `u24.aarch64` | pgdg | 16.3 KiB | [postgresql-18-preprepare_0.9-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-18-preprepare_0.9-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `preprepare_17` | 0.9 | `el8.x86_64` | pigsty | 13.4 KiB | [preprepare_17-0.9-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/preprepare_17-0.9-1PIGSTY.el8.x86_64.rpm) |
| `preprepare_17` | 0.9 | `el8.aarch64` | pigsty | 13.4 KiB | [preprepare_17-0.9-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/preprepare_17-0.9-1PIGSTY.el8.aarch64.rpm) |
| `preprepare_17` | 0.9 | `el9.aarch64` | pigsty | 13.2 KiB | [preprepare_17-0.9-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/preprepare_17-0.9-1PIGSTY.el9.aarch64.rpm) |
| `preprepare_17` | 0.9 | `el9.x86_64` | pigsty | 13.5 KiB | [preprepare_17-0.9-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/preprepare_17-0.9-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-preprepare` | 0.9 | `d12.x86_64` | pgdg | 17.7 KiB | [postgresql-17-preprepare_0.9-9.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-17-preprepare_0.9-9.pgdg120+1_amd64.deb) |
| `postgresql-17-preprepare` | 0.9 | `d12.aarch64` | pgdg | 17.7 KiB | [postgresql-17-preprepare_0.9-9.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-17-preprepare_0.9-9.pgdg120+1_arm64.deb) |
| `postgresql-17-preprepare` | 0.9 | `u22.aarch64` | pgdg | 18.1 KiB | [postgresql-17-preprepare_0.9-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-17-preprepare_0.9-9.pgdg22.04+1_arm64.deb) |
| `postgresql-17-preprepare` | 0.9 | `u22.x86_64` | pgdg | 18.1 KiB | [postgresql-17-preprepare_0.9-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-17-preprepare_0.9-9.pgdg22.04+1_amd64.deb) |
| `postgresql-17-preprepare` | 0.9 | `u24.aarch64` | pgdg | 16.3 KiB | [postgresql-17-preprepare_0.9-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-17-preprepare_0.9-9.pgdg24.04+1_arm64.deb) |
| `postgresql-17-preprepare` | 0.9 | `u24.x86_64` | pgdg | 16.4 KiB | [postgresql-17-preprepare_0.9-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-17-preprepare_0.9-9.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `preprepare_16` | 0.9 | `el8.x86_64` | pigsty | 13.4 KiB | [preprepare_16-0.9-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/preprepare_16-0.9-1PIGSTY.el8.x86_64.rpm) |
| `preprepare_16` | 0.9 | `el8.aarch64` | pigsty | 13.4 KiB | [preprepare_16-0.9-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/preprepare_16-0.9-1PIGSTY.el8.aarch64.rpm) |
| `preprepare_16` | 0.9 | `el9.x86_64` | pigsty | 13.5 KiB | [preprepare_16-0.9-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/preprepare_16-0.9-1PIGSTY.el9.x86_64.rpm) |
| `preprepare_16` | 0.9 | `el9.aarch64` | pigsty | 13.2 KiB | [preprepare_16-0.9-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/preprepare_16-0.9-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-preprepare` | 0.9 | `d12.aarch64` | pgdg | 17.7 KiB | [postgresql-16-preprepare_0.9-9.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-16-preprepare_0.9-9.pgdg120+1_arm64.deb) |
| `postgresql-16-preprepare` | 0.9 | `d12.x86_64` | pgdg | 17.7 KiB | [postgresql-16-preprepare_0.9-9.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-16-preprepare_0.9-9.pgdg120+1_amd64.deb) |
| `postgresql-16-preprepare` | 0.9 | `u22.aarch64` | pgdg | 18.1 KiB | [postgresql-16-preprepare_0.9-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-16-preprepare_0.9-9.pgdg22.04+1_arm64.deb) |
| `postgresql-16-preprepare` | 0.9 | `u22.x86_64` | pgdg | 18.1 KiB | [postgresql-16-preprepare_0.9-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-16-preprepare_0.9-9.pgdg22.04+1_amd64.deb) |
| `postgresql-16-preprepare` | 0.9 | `u24.x86_64` | pgdg | 16.4 KiB | [postgresql-16-preprepare_0.9-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-16-preprepare_0.9-9.pgdg24.04+1_amd64.deb) |
| `postgresql-16-preprepare` | 0.9 | `u24.aarch64` | pgdg | 16.3 KiB | [postgresql-16-preprepare_0.9-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-16-preprepare_0.9-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `preprepare_15` | 0.9 | `el8.aarch64` | pigsty | 13.4 KiB | [preprepare_15-0.9-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/preprepare_15-0.9-1PIGSTY.el8.aarch64.rpm) |
| `preprepare_15` | 0.9 | `el8.x86_64` | pigsty | 13.4 KiB | [preprepare_15-0.9-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/preprepare_15-0.9-1PIGSTY.el8.x86_64.rpm) |
| `preprepare_15` | 0.9 | `el9.x86_64` | pigsty | 13.5 KiB | [preprepare_15-0.9-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/preprepare_15-0.9-1PIGSTY.el9.x86_64.rpm) |
| `preprepare_15` | 0.9 | `el9.aarch64` | pigsty | 13.2 KiB | [preprepare_15-0.9-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/preprepare_15-0.9-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-preprepare` | 0.9 | `d12.x86_64` | pgdg | 17.7 KiB | [postgresql-15-preprepare_0.9-9.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-15-preprepare_0.9-9.pgdg120+1_amd64.deb) |
| `postgresql-15-preprepare` | 0.9 | `d12.aarch64` | pgdg | 17.7 KiB | [postgresql-15-preprepare_0.9-9.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-15-preprepare_0.9-9.pgdg120+1_arm64.deb) |
| `postgresql-15-preprepare` | 0.9 | `u22.aarch64` | pgdg | 18.1 KiB | [postgresql-15-preprepare_0.9-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-15-preprepare_0.9-9.pgdg22.04+1_arm64.deb) |
| `postgresql-15-preprepare` | 0.9 | `u22.x86_64` | pgdg | 18.1 KiB | [postgresql-15-preprepare_0.9-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-15-preprepare_0.9-9.pgdg22.04+1_amd64.deb) |
| `postgresql-15-preprepare` | 0.9 | `u24.x86_64` | pgdg | 16.4 KiB | [postgresql-15-preprepare_0.9-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-15-preprepare_0.9-9.pgdg24.04+1_amd64.deb) |
| `postgresql-15-preprepare` | 0.9 | `u24.aarch64` | pgdg | 16.3 KiB | [postgresql-15-preprepare_0.9-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-15-preprepare_0.9-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `preprepare_14` | 0.9 | `el8.x86_64` | pigsty | 13.4 KiB | [preprepare_14-0.9-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/preprepare_14-0.9-1PIGSTY.el8.x86_64.rpm) |
| `preprepare_14` | 0.9 | `el8.aarch64` | pigsty | 13.4 KiB | [preprepare_14-0.9-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/preprepare_14-0.9-1PIGSTY.el8.aarch64.rpm) |
| `preprepare_14` | 0.9 | `el9.x86_64` | pigsty | 13.4 KiB | [preprepare_14-0.9-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/preprepare_14-0.9-1PIGSTY.el9.x86_64.rpm) |
| `preprepare_14` | 0.9 | `el9.aarch64` | pigsty | 13.2 KiB | [preprepare_14-0.9-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/preprepare_14-0.9-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-preprepare` | 0.9 | `d12.x86_64` | pgdg | 17.6 KiB | [postgresql-14-preprepare_0.9-9.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-14-preprepare_0.9-9.pgdg120+1_amd64.deb) |
| `postgresql-14-preprepare` | 0.9 | `d12.aarch64` | pgdg | 17.7 KiB | [postgresql-14-preprepare_0.9-9.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-14-preprepare_0.9-9.pgdg120+1_arm64.deb) |
| `postgresql-14-preprepare` | 0.9 | `u22.x86_64` | pgdg | 18.0 KiB | [postgresql-14-preprepare_0.9-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-14-preprepare_0.9-9.pgdg22.04+1_amd64.deb) |
| `postgresql-14-preprepare` | 0.9 | `u22.aarch64` | pgdg | 18.1 KiB | [postgresql-14-preprepare_0.9-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-14-preprepare_0.9-9.pgdg22.04+1_arm64.deb) |
| `postgresql-14-preprepare` | 0.9 | `u24.aarch64` | pgdg | 16.3 KiB | [postgresql-14-preprepare_0.9-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-14-preprepare_0.9-9.pgdg24.04+1_arm64.deb) |
| `postgresql-14-preprepare` | 0.9 | `u24.x86_64` | pgdg | 16.4 KiB | [postgresql-14-preprepare_0.9-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-14-preprepare_0.9-9.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `preprepare_13` | 0.9 | `el8.aarch64` | pigsty | 13.4 KiB | [preprepare_13-0.9-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/preprepare_13-0.9-1PIGSTY.el8.aarch64.rpm) |
| `preprepare_13` | 0.9 | `el8.x86_64` | pigsty | 13.3 KiB | [preprepare_13-0.9-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/preprepare_13-0.9-1PIGSTY.el8.x86_64.rpm) |
| `preprepare_13` | 0.9 | `el9.aarch64` | pigsty | 13.2 KiB | [preprepare_13-0.9-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/preprepare_13-0.9-1PIGSTY.el9.aarch64.rpm) |
| `preprepare_13` | 0.9 | `el9.x86_64` | pigsty | 13.4 KiB | [preprepare_13-0.9-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/preprepare_13-0.9-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-preprepare` | 0.9 | `d12.aarch64` | pgdg | 17.4 KiB | [postgresql-13-preprepare_0.9-9.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-13-preprepare_0.9-9.pgdg120+1_arm64.deb) |
| `postgresql-13-preprepare` | 0.9 | `d12.x86_64` | pgdg | 17.7 KiB | [postgresql-13-preprepare_0.9-9.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-13-preprepare_0.9-9.pgdg120+1_amd64.deb) |
| `postgresql-13-preprepare` | 0.9 | `u22.aarch64` | pgdg | 17.7 KiB | [postgresql-13-preprepare_0.9-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-13-preprepare_0.9-9.pgdg22.04+1_arm64.deb) |
| `postgresql-13-preprepare` | 0.9 | `u22.x86_64` | pgdg | 18.0 KiB | [postgresql-13-preprepare_0.9-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-13-preprepare_0.9-9.pgdg22.04+1_amd64.deb) |
| `postgresql-13-preprepare` | 0.9 | `u24.x86_64` | pgdg | 16.5 KiB | [postgresql-13-preprepare_0.9-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-13-preprepare_0.9-9.pgdg24.04+1_amd64.deb) |
| `postgresql-13-preprepare` | 0.9 | `u24.aarch64` | pgdg | 16.1 KiB | [postgresql-13-preprepare_0.9-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/preprepare/postgresql-13-preprepare_0.9-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/dimitri/preprepare" title="Repository" icon="github" subtitle="github.com/dimitri/preprepare" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="preprepare-0.9.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pre_prepare; # get pre_prepare source code
pig build dep pre_prepare; # install build dependencies
pig build pkg pre_prepare; # build extension rpm or deb
pig build ext pre_prepare; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pre_prepare; # install by extension name, for the current active PG version
pig ext install preprepare; # install via package alias, for the active PG version
pig ext install pre_prepare -v 18;   # install for PG 18
pig ext install pre_prepare -v 17;   # install for PG 17
pig ext install pre_prepare -v 16;   # install for PG 16
pig ext install pre_prepare -v 15;   # install for PG 15
pig ext install pre_prepare -v 14;   # install for PG 14
pig ext install pre_prepare -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pre_prepare;
```

