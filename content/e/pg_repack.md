---
title: "pg_repack"
linkTitle: "pg_repack"
description: "Reorganize tables in PostgreSQL databases with minimal locks"
weight: 5010
categories: ["Admin"]
width: full
---

Reorganize tables in PostgreSQL databases with minimal locks

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5010** | {{< badge content="pg_repack" link="https://github.com/reorg/pg_repack" >}} | {{< ext "pg_repack" "pg_repack" >}} | `1.5.2` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--bs-d--" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_rewrite" >}} {{< ext "pg_squeeze" >}} {{< ext "pgfincore" >}} {{< ext "pg_prewarm" >}} {{< ext "pg_buffercache" >}} {{< ext "pgstattuple" >}} {{< ext "pg_cooldown" >}} {{< ext "pgcozy" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_repack" >}} | `1.5.2` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_repack_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pg_repack" >}} | `1.5.2` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-repack` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_repack_18" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_repack_18-1.5.2-5PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_repack_17" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_repack_17-1.5.2-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_repack_16" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_repack_16-1.5.2-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_repack_15" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_repack_15-1.5.2-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_repack_14" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_repack_14-1.5.2-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_repack_18" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_repack_18-1.5.2-5PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_repack_17" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_repack_17-1.5.2-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_repack_16" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_repack_16-1.5.2-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_repack_15" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_repack_15-1.5.2-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_repack_14" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_repack_14-1.5.2-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_repack_18" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_repack_18-1.5.2-5PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_repack_17" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_repack_17-1.5.2-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_repack_16" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_repack_16-1.5.2-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_repack_15" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_repack_15-1.5.2-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_repack_14" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_repack_14-1.5.2-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_repack_18" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_repack_18-1.5.2-5PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_repack_17" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_repack_17-1.5.2-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_repack_16" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_repack_16-1.5.2-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_repack_15" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_repack_15-1.5.2-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_repack_14" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_repack_14-1.5.2-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-repack" "1.5.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_repack_18` | 1.5.2 | `el8.aarch64` | pgdg | 74.5 KiB | [pg_repack_18-1.5.2-5PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_repack_18-1.5.2-5PGDG.rhel8.aarch64.rpm) |
| `pg_repack_18` | 1.5.2 | `el8.x86_64` | pgdg | 75.8 KiB | [pg_repack_18-1.5.2-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_repack_18-1.5.2-5PGDG.rhel8.x86_64.rpm) |
| `pg_repack_18` | 1.5.2 | `el9.x86_64` | pgdg | 66.2 KiB | [pg_repack_18-1.5.2-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_repack_18-1.5.2-5PGDG.rhel9.x86_64.rpm) |
| `pg_repack_18` | 1.5.2 | `el9.aarch64` | pgdg | 65.5 KiB | [pg_repack_18-1.5.2-5PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_repack_18-1.5.2-5PGDG.rhel9.aarch64.rpm) |
| `postgresql-18-repack` | 1.5.2 | `d12.x86_64` | pgdg | 101.6 KiB | [postgresql-18-repack_1.5.2-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg12+1_amd64.deb) |
| `postgresql-18-repack` | 1.5.2 | `d12.aarch64` | pgdg | 99.1 KiB | [postgresql-18-repack_1.5.2-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg12+1_arm64.deb) |
| `postgresql-18-repack` | 1.5.2 | `u22.aarch64` | pgdg | 97.4 KiB | [postgresql-18-repack_1.5.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-repack` | 1.5.2 | `u22.x86_64` | pgdg | 100.6 KiB | [postgresql-18-repack_1.5.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-repack` | 1.5.2 | `u24.aarch64` | pgdg | 96.4 KiB | [postgresql-18-repack_1.5.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg24.04+1_arm64.deb) |
| `postgresql-18-repack` | 1.5.2 | `u24.x86_64` | pgdg | 98.6 KiB | [postgresql-18-repack_1.5.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_repack_17` | 1.5.2 | `el8.aarch64` | pgdg | 74.0 KiB | [pg_repack_17-1.5.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_repack_17-1.5.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_17` | 1.5.2 | `el8.x86_64` | pgdg | 75.3 KiB | [pg_repack_17-1.5.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_repack_17-1.5.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_17` | 1.5.1 | `el8.aarch64` | pgdg | 73.6 KiB | [pg_repack_17-1.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_repack_17-1.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_17` | 1.5.1 | `el8.x86_64` | pgdg | 74.9 KiB | [pg_repack_17-1.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_repack_17-1.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_17` | 1.5.2 | `el9.x86_64` | pgdg | 65.8 KiB | [pg_repack_17-1.5.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_repack_17-1.5.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_17` | 1.5.2 | `el9.aarch64` | pgdg | 65.2 KiB | [pg_repack_17-1.5.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_repack_17-1.5.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_17` | 1.5.1 | `el9.aarch64` | pgdg | 65.1 KiB | [pg_repack_17-1.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_repack_17-1.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_17` | 1.5.1 | `el9.x86_64` | pgdg | 65.7 KiB | [pg_repack_17-1.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_repack_17-1.5.1-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-17-repack` | 1.5.2 | `d12.x86_64` | pgdg | 101.7 KiB | [postgresql-17-repack_1.5.2-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg12+1_amd64.deb) |
| `postgresql-17-repack` | 1.5.2 | `d12.aarch64` | pgdg | 99.1 KiB | [postgresql-17-repack_1.5.2-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg12+1_arm64.deb) |
| `postgresql-17-repack` | 1.5.2 | `u22.aarch64` | pgdg | 102.1 KiB | [postgresql-17-repack_1.5.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-repack` | 1.5.2 | `u22.x86_64` | pgdg | 105.8 KiB | [postgresql-17-repack_1.5.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-repack` | 1.5.2 | `u24.x86_64` | pgdg | 98.7 KiB | [postgresql-17-repack_1.5.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-repack` | 1.5.2 | `u24.aarch64` | pgdg | 96.5 KiB | [postgresql-17-repack_1.5.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_repack_16` | 1.5.2 | `el8.x86_64` | pgdg | 75.3 KiB | [pg_repack_16-1.5.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_repack_16-1.5.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_16` | 1.5.2 | `el8.aarch64` | pgdg | 74.0 KiB | [pg_repack_16-1.5.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_repack_16-1.5.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_16` | 1.5.0 | `el8.x86_64` | pgdg | 79.5 KiB | [pg_repack_16-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_repack_16-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_16` | 1.5.0 | `el8.aarch64` | pgdg | 79.0 KiB | [pg_repack_16-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_repack_16-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_16` | 1.5.2 | `el9.x86_64` | pgdg | 65.8 KiB | [pg_repack_16-1.5.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_repack_16-1.5.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_16` | 1.5.2 | `el9.aarch64` | pgdg | 65.2 KiB | [pg_repack_16-1.5.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_repack_16-1.5.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_16` | 1.5.0 | `el9.aarch64` | pgdg | 67.4 KiB | [pg_repack_16-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_repack_16-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_16` | 1.5.0 | `el9.x86_64` | pgdg | 68.0 KiB | [pg_repack_16-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_repack_16-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-16-repack` | 1.5.2 | `d12.x86_64` | pgdg | 101.8 KiB | [postgresql-16-repack_1.5.2-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg12+1_amd64.deb) |
| `postgresql-16-repack` | 1.5.2 | `d12.aarch64` | pgdg | 99.4 KiB | [postgresql-16-repack_1.5.2-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg12+1_arm64.deb) |
| `postgresql-16-repack` | 1.5.2 | `u22.aarch64` | pgdg | 102.0 KiB | [postgresql-16-repack_1.5.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-repack` | 1.5.2 | `u22.x86_64` | pgdg | 105.4 KiB | [postgresql-16-repack_1.5.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-repack` | 1.5.2 | `u24.aarch64` | pgdg | 96.7 KiB | [postgresql-16-repack_1.5.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg24.04+1_arm64.deb) |
| `postgresql-16-repack` | 1.5.2 | `u24.x86_64` | pgdg | 99.1 KiB | [postgresql-16-repack_1.5.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_repack_15` | 1.5.2 | `el8.aarch64` | pgdg | 74.3 KiB | [pg_repack_15-1.5.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_repack_15-1.5.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_15` | 1.5.2 | `el8.x86_64` | pgdg | 75.6 KiB | [pg_repack_15-1.5.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_repack_15-1.5.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_15` | 1.5.0 | `el8.x86_64` | pgdg | 79.6 KiB | [pg_repack_15-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_repack_15-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_15` | 1.5.0 | `el8.aarch64` | pgdg | 79.2 KiB | [pg_repack_15-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_repack_15-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_15` | 1.4.8 | `el8.x86_64` | pgdg | 106.6 KiB | [pg_repack_15-1.4.8-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_repack_15-1.4.8-1.rhel8.x86_64.rpm) |
| `pg_repack_15` | 1.4.8 | `el8.aarch64` | pgdg | 106.0 KiB | [pg_repack_15-1.4.8-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_repack_15-1.4.8-1.rhel8.aarch64.rpm) |
| `pg_repack_15` | 1.5.2 | `el9.aarch64` | pgdg | 65.7 KiB | [pg_repack_15-1.5.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_repack_15-1.5.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_15` | 1.5.2 | `el9.x86_64` | pgdg | 66.2 KiB | [pg_repack_15-1.5.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_repack_15-1.5.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_15` | 1.5.0 | `el9.aarch64` | pgdg | 67.7 KiB | [pg_repack_15-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_repack_15-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_15` | 1.5.0 | `el9.x86_64` | pgdg | 68.3 KiB | [pg_repack_15-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_repack_15-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_15` | 1.4.8 | `el9.x86_64` | pgdg | 96.1 KiB | [pg_repack_15-1.4.8-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_repack_15-1.4.8-1.rhel9.x86_64.rpm) |
| `pg_repack_15` | 1.4.8 | `el9.aarch64` | pgdg | 95.0 KiB | [pg_repack_15-1.4.8-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_repack_15-1.4.8-1.rhel9.aarch64.rpm) |
| `postgresql-15-repack` | 1.5.2 | `d12.aarch64` | pgdg | 99.6 KiB | [postgresql-15-repack_1.5.2-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg12+1_arm64.deb) |
| `postgresql-15-repack` | 1.5.2 | `d12.x86_64` | pgdg | 102.2 KiB | [postgresql-15-repack_1.5.2-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg12+1_amd64.deb) |
| `postgresql-15-repack` | 1.5.2 | `u22.x86_64` | pgdg | 105.3 KiB | [postgresql-15-repack_1.5.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-repack` | 1.5.2 | `u22.aarch64` | pgdg | 102.1 KiB | [postgresql-15-repack_1.5.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-repack` | 1.5.2 | `u24.aarch64` | pgdg | 97.1 KiB | [postgresql-15-repack_1.5.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg24.04+1_arm64.deb) |
| `postgresql-15-repack` | 1.5.2 | `u24.x86_64` | pgdg | 99.1 KiB | [postgresql-15-repack_1.5.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_repack_14` | 1.5.2 | `el8.x86_64` | pgdg | 73.9 KiB | [pg_repack_14-1.5.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_repack_14-1.5.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_14` | 1.5.2 | `el8.aarch64` | pgdg | 73.0 KiB | [pg_repack_14-1.5.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_repack_14-1.5.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_14` | 1.5.0 | `el8.aarch64` | pgdg | 77.8 KiB | [pg_repack_14-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_repack_14-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_14` | 1.5.0 | `el8.x86_64` | pgdg | 78.0 KiB | [pg_repack_14-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_repack_14-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_14` | 1.4.8 | `el8.aarch64` | pgdg | 103.7 KiB | [pg_repack_14-1.4.8-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_repack_14-1.4.8-1.rhel8.aarch64.rpm) |
| `pg_repack_14` | 1.4.8 | `el8.x86_64` | pgdg | 103.7 KiB | [pg_repack_14-1.4.8-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_repack_14-1.4.8-1.rhel8.x86_64.rpm) |
| `pg_repack_14` | 1.4.7 | `el8.x86_64` | pgdg | 103.7 KiB | [pg_repack_14-1.4.7-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_repack_14-1.4.7-1.rhel8.x86_64.rpm) |
| `pg_repack_14` | 1.5.2 | `el9.x86_64` | pgdg | 65.3 KiB | [pg_repack_14-1.5.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_repack_14-1.5.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_14` | 1.5.2 | `el9.aarch64` | pgdg | 65.0 KiB | [pg_repack_14-1.5.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_repack_14-1.5.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_14` | 1.5.0 | `el9.aarch64` | pgdg | 67.1 KiB | [pg_repack_14-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_repack_14-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_14` | 1.5.0 | `el9.x86_64` | pgdg | 67.6 KiB | [pg_repack_14-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_repack_14-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_14` | 1.4.8 | `el9.x86_64` | pgdg | 94.2 KiB | [pg_repack_14-1.4.8-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_repack_14-1.4.8-1.rhel9.x86_64.rpm) |
| `pg_repack_14` | 1.4.8 | `el9.aarch64` | pgdg | 93.4 KiB | [pg_repack_14-1.4.8-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_repack_14-1.4.8-1.rhel9.aarch64.rpm) |
| `pg_repack_14` | 1.4.7 | `el9.x86_64` | pgdg | 93.7 KiB | [pg_repack_14-1.4.7-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_repack_14-1.4.7-1.rhel9.x86_64.rpm) |
| `postgresql-14-repack` | 1.5.2 | `d12.x86_64` | pgdg | 101.4 KiB | [postgresql-14-repack_1.5.2-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg12+1_amd64.deb) |
| `postgresql-14-repack` | 1.5.2 | `d12.aarch64` | pgdg | 99.3 KiB | [postgresql-14-repack_1.5.2-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg12+1_arm64.deb) |
| `postgresql-14-repack` | 1.5.2 | `u22.aarch64` | pgdg | 100.8 KiB | [postgresql-14-repack_1.5.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-repack` | 1.5.2 | `u22.x86_64` | pgdg | 103.9 KiB | [postgresql-14-repack_1.5.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-repack` | 1.5.2 | `u24.x86_64` | pgdg | 98.4 KiB | [postgresql-14-repack_1.5.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-repack` | 1.5.2 | `u24.aarch64` | pgdg | 96.6 KiB | [postgresql-14-repack_1.5.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_repack_13` | 1.5.2 | `el8.aarch64` | pgdg | 71.7 KiB | [pg_repack_13-1.5.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_repack_13-1.5.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_13` | 1.5.2 | `el8.x86_64` | pgdg | 72.4 KiB | [pg_repack_13-1.5.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_repack_13-1.5.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_13` | 1.5.0 | `el8.aarch64` | pgdg | 76.1 KiB | [pg_repack_13-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_repack_13-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_13` | 1.5.0 | `el8.x86_64` | pgdg | 76.0 KiB | [pg_repack_13-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_repack_13-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_13` | 1.4.8 | `el8.aarch64` | pgdg | 101.7 KiB | [pg_repack_13-1.4.8-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_repack_13-1.4.8-1.rhel8.aarch64.rpm) |
| `pg_repack_13` | 1.4.8 | `el8.x86_64` | pgdg | 101.4 KiB | [pg_repack_13-1.4.8-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_repack_13-1.4.8-1.rhel8.x86_64.rpm) |
| `pg_repack_13` | 1.4.7 | `el8.x86_64` | pgdg | 101.0 KiB | [pg_repack_13-1.4.7-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_repack_13-1.4.7-1.rhel8.x86_64.rpm) |
| `pg_repack_13` | 1.5.2 | `el9.aarch64` | pgdg | 64.7 KiB | [pg_repack_13-1.5.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_repack_13-1.5.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_13` | 1.5.2 | `el9.x86_64` | pgdg | 64.9 KiB | [pg_repack_13-1.5.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_repack_13-1.5.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_13` | 1.5.0 | `el9.aarch64` | pgdg | 66.8 KiB | [pg_repack_13-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_repack_13-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_13` | 1.5.0 | `el9.x86_64` | pgdg | 67.2 KiB | [pg_repack_13-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_repack_13-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_13` | 1.4.8 | `el9.x86_64` | pgdg | 93.6 KiB | [pg_repack_13-1.4.8-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_repack_13-1.4.8-1.rhel9.x86_64.rpm) |
| `pg_repack_13` | 1.4.8 | `el9.aarch64` | pgdg | 92.7 KiB | [pg_repack_13-1.4.8-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_repack_13-1.4.8-1.rhel9.aarch64.rpm) |
| `pg_repack_13` | 1.4.7 | `el9.x86_64` | pgdg | 93.2 KiB | [pg_repack_13-1.4.7-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_repack_13-1.4.7-1.rhel9.x86_64.rpm) |
| `postgresql-13-repack` | 1.5.2 | `d12.aarch64` | pgdg | 97.9 KiB | [postgresql-13-repack_1.5.2-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-13-repack_1.5.2-2.pgdg12+1_arm64.deb) |
| `postgresql-13-repack` | 1.5.2 | `d12.x86_64` | pgdg | 100.1 KiB | [postgresql-13-repack_1.5.2-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-13-repack_1.5.2-2.pgdg12+1_amd64.deb) |
| `postgresql-13-repack` | 1.5.2 | `u22.aarch64` | pgdg | 99.7 KiB | [postgresql-13-repack_1.5.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-13-repack_1.5.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-repack` | 1.5.2 | `u22.x86_64` | pgdg | 102.4 KiB | [postgresql-13-repack_1.5.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-13-repack_1.5.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-repack` | 1.5.2 | `u24.x86_64` | pgdg | 97.4 KiB | [postgresql-13-repack_1.5.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-13-repack_1.5.2-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-repack` | 1.5.2 | `u24.aarch64` | pgdg | 95.0 KiB | [postgresql-13-repack_1.5.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-13-repack_1.5.2-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/reorg/pg_repack" title="Repository" icon="github" subtitle="github.com/reorg/pg_repack" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_repack; # install by extension name, for the current active PG version
pig ext install pg_repack; # install via package alias, for the active PG version
pig ext install pg_repack -v 18;   # install for PG 18
pig ext install pg_repack -v 17;   # install for PG 17
pig ext install pg_repack -v 16;   # install for PG 16
pig ext install pg_repack -v 15;   # install for PG 15
pig ext install pg_repack -v 14;   # install for PG 14
pig ext install pg_repack -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_repack;
```

