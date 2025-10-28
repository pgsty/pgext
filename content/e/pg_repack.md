---
title: "pg_repack"
linkTitle: "pg_repack"
description: "Reorganize tables in PostgreSQL databases with minimal locks"
weight: 5010
categories: ["ADMIN"]
width: full
---

Reorganize tables in PostgreSQL databases with minimal locks


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5010** | {{< badge content="pg_repack" link="https://github.com/reorg/pg_repack" >}} | {{< ext "pg_repack" >}} | `1.5.2` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_rewrite" >}} {{< ext "pg_squeeze" >}} {{< ext "pgfincore" >}} {{< ext "pg_prewarm" >}} {{< ext "pg_buffercache" >}} {{< ext "pgstattuple" >}} {{< ext "pg_cooldown" >}} {{< ext "pgcozy" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_repack" >}} | `1.5.2` | {{< bg "18" "pg_repack_18*" "green" >}} {{< bg "17" "pg_repack_17*" "green" >}} {{< bg "16" "pg_repack_16*" "green" >}} {{< bg "15" "pg_repack_15*" "green" >}} {{< bg "14" "pg_repack_14*" "green" >}} {{< bg "13" "pg_repack_13*" "green" >}} | `pg_repack_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pg_repack" >}} | `1.5.2` | {{< bg "18" "postgresql-18-repack" "green" >}} {{< bg "17" "postgresql-17-repack" "green" >}} {{< bg "16" "postgresql-16-repack" "green" >}} {{< bg "15" "postgresql-15-repack" "green" >}} {{< bg "14" "postgresql-14-repack" "green" >}} {{< bg "13" "postgresql-13-repack" "green" >}} | `postgresql-$v-repack` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 1.5.3" "pg_repack_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_14 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_13 : AVAIL 5" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 1.5.3" "pg_repack_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_13 : AVAIL 4" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 1.5.3" "pg_repack_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_14 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_13 : AVAIL 5" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 1.5.3" "pg_repack_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_13 : AVAIL 4" "blue" >}} |
|    `el10.x86_64`    | {{< bg "PGDG 1.5.3" "pg_repack_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_13 : AVAIL 2" "blue" >}} |
|    `el10.aarch64`    | {{< bg "PGDG 1.5.3" "pg_repack_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_13 : AVAIL 2" "blue" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 1.5.2" "postgresql-18-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-17-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-16-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-15-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-14-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-13-repack : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 1.5.2" "postgresql-18-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-17-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-16-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-15-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-14-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-13-repack : AVAIL 1" "blue" >}} |
|    `d13.x86_64`    | {{< bg "PGDG 1.5.2" "postgresql-18-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-17-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-16-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-15-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-14-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-13-repack : AVAIL 1" "blue" >}} |
|    `d13.aarch64`    | {{< bg "PGDG 1.5.2" "postgresql-18-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-17-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-16-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-15-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-14-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-13-repack : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 1.5.2" "postgresql-18-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-17-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-16-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-15-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-14-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-13-repack : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 1.5.2" "postgresql-18-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-17-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-16-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-15-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-14-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-13-repack : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 1.5.2" "postgresql-18-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-17-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-16-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-15-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-14-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-13-repack : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 1.5.2" "postgresql-18-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-17-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-16-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-15-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-14-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.2" "postgresql-13-repack : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_repack_18` | 1.5.3 | `el8.x86_64` | pgdg | 76.4 KiB | [pg_repack_18-1.5.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_repack_18-1.5.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_18` | 1.5.2 | `el8.x86_64` | pgdg | 75.8 KiB | [pg_repack_18-1.5.2-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_repack_18-1.5.2-5PGDG.rhel8.x86_64.rpm) |
| `pg_repack_18` | 1.5.3 | `el8.aarch64` | pgdg | 75.3 KiB | [pg_repack_18-1.5.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_repack_18-1.5.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_18` | 1.5.2 | `el8.aarch64` | pgdg | 74.5 KiB | [pg_repack_18-1.5.2-5PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_repack_18-1.5.2-5PGDG.rhel8.aarch64.rpm) |
| `pg_repack_18` | 1.5.3 | `el9.x86_64` | pgdg | 66.9 KiB | [pg_repack_18-1.5.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_repack_18-1.5.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_18` | 1.5.2 | `el9.x86_64` | pgdg | 66.2 KiB | [pg_repack_18-1.5.2-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_repack_18-1.5.2-5PGDG.rhel9.x86_64.rpm) |
| `pg_repack_18` | 1.5.3 | `el9.aarch64` | pgdg | 66.3 KiB | [pg_repack_18-1.5.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_repack_18-1.5.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_18` | 1.5.2 | `el9.aarch64` | pgdg | 65.5 KiB | [pg_repack_18-1.5.2-5PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_repack_18-1.5.2-5PGDG.rhel9.aarch64.rpm) |
| `pg_repack_18` | 1.5.3 | `el10.x86_64` | pgdg | 67.9 KiB | [pg_repack_18-1.5.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_repack_18-1.5.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_repack_18` | 1.5.2 | `el10.x86_64` | pgdg | 67.1 KiB | [pg_repack_18-1.5.2-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_repack_18-1.5.2-5PGDG.rhel10.x86_64.rpm) |
| `pg_repack_18` | 1.5.3 | `el10.aarch64` | pgdg | 67.8 KiB | [pg_repack_18-1.5.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_repack_18-1.5.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_repack_18` | 1.5.2 | `el10.aarch64` | pgdg | 67.1 KiB | [pg_repack_18-1.5.2-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_repack_18-1.5.2-5PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-repack` | 1.5.2 | `d12.x86_64` | pgdg | 101.6 KiB | [postgresql-18-repack_1.5.2-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg12+1_amd64.deb) |
| `postgresql-18-repack` | 1.5.2 | `d12.aarch64` | pgdg | 99.1 KiB | [postgresql-18-repack_1.5.2-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg12+1_arm64.deb) |
| `postgresql-18-repack` | 1.5.2 | `d13.x86_64` | pgdg | 102.0 KiB | [postgresql-18-repack_1.5.2-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg13+1_amd64.deb) |
| `postgresql-18-repack` | 1.5.2 | `d13.aarch64` | pgdg | 100.0 KiB | [postgresql-18-repack_1.5.2-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg13+1_arm64.deb) |
| `postgresql-18-repack` | 1.5.2 | `u22.x86_64` | pgdg | 100.6 KiB | [postgresql-18-repack_1.5.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-repack` | 1.5.2 | `u22.aarch64` | pgdg | 97.4 KiB | [postgresql-18-repack_1.5.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-repack` | 1.5.2 | `u24.x86_64` | pgdg | 98.6 KiB | [postgresql-18-repack_1.5.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-repack` | 1.5.2 | `u24.aarch64` | pgdg | 96.4 KiB | [postgresql-18-repack_1.5.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.2-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_repack_17` | 1.5.3 | `el8.x86_64` | pgdg | 76.7 KiB | [pg_repack_17-1.5.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_repack_17-1.5.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_17` | 1.5.2 | `el8.x86_64` | pgdg | 75.3 KiB | [pg_repack_17-1.5.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_repack_17-1.5.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_17` | 1.5.1 | `el8.x86_64` | pgdg | 74.9 KiB | [pg_repack_17-1.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_repack_17-1.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_17` | 1.5.3 | `el8.aarch64` | pgdg | 75.5 KiB | [pg_repack_17-1.5.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_repack_17-1.5.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_17` | 1.5.2 | `el8.aarch64` | pgdg | 74.0 KiB | [pg_repack_17-1.5.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_repack_17-1.5.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_17` | 1.5.1 | `el8.aarch64` | pgdg | 73.6 KiB | [pg_repack_17-1.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_repack_17-1.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_17` | 1.5.3 | `el9.x86_64` | pgdg | 67.0 KiB | [pg_repack_17-1.5.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_repack_17-1.5.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_17` | 1.5.2 | `el9.x86_64` | pgdg | 65.8 KiB | [pg_repack_17-1.5.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_repack_17-1.5.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_17` | 1.5.1 | `el9.x86_64` | pgdg | 65.7 KiB | [pg_repack_17-1.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_repack_17-1.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_17` | 1.5.3 | `el9.aarch64` | pgdg | 66.3 KiB | [pg_repack_17-1.5.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_repack_17-1.5.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_17` | 1.5.2 | `el9.aarch64` | pgdg | 65.2 KiB | [pg_repack_17-1.5.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_repack_17-1.5.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_17` | 1.5.1 | `el9.aarch64` | pgdg | 65.1 KiB | [pg_repack_17-1.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_repack_17-1.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_17` | 1.5.3 | `el10.x86_64` | pgdg | 68.0 KiB | [pg_repack_17-1.5.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_repack_17-1.5.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_repack_17` | 1.5.2 | `el10.x86_64` | pgdg | 67.2 KiB | [pg_repack_17-1.5.2-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_repack_17-1.5.2-4PGDG.rhel10.x86_64.rpm) |
| `pg_repack_17` | 1.5.3 | `el10.aarch64` | pgdg | 67.9 KiB | [pg_repack_17-1.5.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_repack_17-1.5.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_repack_17` | 1.5.2 | `el10.aarch64` | pgdg | 67.1 KiB | [pg_repack_17-1.5.2-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_repack_17-1.5.2-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-repack` | 1.5.2 | `d12.x86_64` | pgdg | 101.7 KiB | [postgresql-17-repack_1.5.2-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg12+1_amd64.deb) |
| `postgresql-17-repack` | 1.5.2 | `d12.aarch64` | pgdg | 99.1 KiB | [postgresql-17-repack_1.5.2-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg12+1_arm64.deb) |
| `postgresql-17-repack` | 1.5.2 | `d13.x86_64` | pgdg | 102.1 KiB | [postgresql-17-repack_1.5.2-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg13+1_amd64.deb) |
| `postgresql-17-repack` | 1.5.2 | `d13.aarch64` | pgdg | 100.1 KiB | [postgresql-17-repack_1.5.2-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg13+1_arm64.deb) |
| `postgresql-17-repack` | 1.5.2 | `u22.x86_64` | pgdg | 105.8 KiB | [postgresql-17-repack_1.5.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-repack` | 1.5.2 | `u22.aarch64` | pgdg | 102.1 KiB | [postgresql-17-repack_1.5.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-repack` | 1.5.2 | `u24.x86_64` | pgdg | 98.7 KiB | [postgresql-17-repack_1.5.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-repack` | 1.5.2 | `u24.aarch64` | pgdg | 96.5 KiB | [postgresql-17-repack_1.5.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.2-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_repack_16` | 1.5.3 | `el8.x86_64` | pgdg | 76.6 KiB | [pg_repack_16-1.5.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_repack_16-1.5.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_16` | 1.5.2 | `el8.x86_64` | pgdg | 75.3 KiB | [pg_repack_16-1.5.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_repack_16-1.5.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_16` | 1.5.0 | `el8.x86_64` | pgdg | 79.5 KiB | [pg_repack_16-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_repack_16-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_16` | 1.5.3 | `el8.aarch64` | pgdg | 75.4 KiB | [pg_repack_16-1.5.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_repack_16-1.5.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_16` | 1.5.2 | `el8.aarch64` | pgdg | 74.0 KiB | [pg_repack_16-1.5.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_repack_16-1.5.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_16` | 1.5.0 | `el8.aarch64` | pgdg | 79.0 KiB | [pg_repack_16-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_repack_16-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_16` | 1.5.3 | `el9.x86_64` | pgdg | 66.9 KiB | [pg_repack_16-1.5.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_repack_16-1.5.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_16` | 1.5.2 | `el9.x86_64` | pgdg | 65.8 KiB | [pg_repack_16-1.5.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_repack_16-1.5.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_16` | 1.5.0 | `el9.x86_64` | pgdg | 68.0 KiB | [pg_repack_16-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_repack_16-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_16` | 1.5.3 | `el9.aarch64` | pgdg | 66.2 KiB | [pg_repack_16-1.5.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_repack_16-1.5.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_16` | 1.5.2 | `el9.aarch64` | pgdg | 65.2 KiB | [pg_repack_16-1.5.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_repack_16-1.5.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_16` | 1.5.0 | `el9.aarch64` | pgdg | 67.4 KiB | [pg_repack_16-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_repack_16-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_16` | 1.5.3 | `el10.x86_64` | pgdg | 68.0 KiB | [pg_repack_16-1.5.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_repack_16-1.5.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_repack_16` | 1.5.2 | `el10.x86_64` | pgdg | 67.1 KiB | [pg_repack_16-1.5.2-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_repack_16-1.5.2-4PGDG.rhel10.x86_64.rpm) |
| `pg_repack_16` | 1.5.3 | `el10.aarch64` | pgdg | 67.8 KiB | [pg_repack_16-1.5.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_repack_16-1.5.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_repack_16` | 1.5.2 | `el10.aarch64` | pgdg | 67.0 KiB | [pg_repack_16-1.5.2-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_repack_16-1.5.2-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-repack` | 1.5.2 | `d12.x86_64` | pgdg | 101.8 KiB | [postgresql-16-repack_1.5.2-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg12+1_amd64.deb) |
| `postgresql-16-repack` | 1.5.2 | `d12.aarch64` | pgdg | 99.4 KiB | [postgresql-16-repack_1.5.2-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg12+1_arm64.deb) |
| `postgresql-16-repack` | 1.5.2 | `d13.x86_64` | pgdg | 102.2 KiB | [postgresql-16-repack_1.5.2-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg13+1_amd64.deb) |
| `postgresql-16-repack` | 1.5.2 | `d13.aarch64` | pgdg | 100.2 KiB | [postgresql-16-repack_1.5.2-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg13+1_arm64.deb) |
| `postgresql-16-repack` | 1.5.2 | `u22.x86_64` | pgdg | 105.4 KiB | [postgresql-16-repack_1.5.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-repack` | 1.5.2 | `u22.aarch64` | pgdg | 102.0 KiB | [postgresql-16-repack_1.5.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-repack` | 1.5.2 | `u24.x86_64` | pgdg | 99.1 KiB | [postgresql-16-repack_1.5.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-repack` | 1.5.2 | `u24.aarch64` | pgdg | 96.7 KiB | [postgresql-16-repack_1.5.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.2-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_repack_15` | 1.5.3 | `el8.x86_64` | pgdg | 76.8 KiB | [pg_repack_15-1.5.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_repack_15-1.5.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_15` | 1.5.2 | `el8.x86_64` | pgdg | 75.6 KiB | [pg_repack_15-1.5.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_repack_15-1.5.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_15` | 1.5.0 | `el8.x86_64` | pgdg | 79.6 KiB | [pg_repack_15-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_repack_15-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_15` | 1.4.8 | `el8.x86_64` | pgdg | 106.6 KiB | [pg_repack_15-1.4.8-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_repack_15-1.4.8-1.rhel8.x86_64.rpm) |
| `pg_repack_15` | 1.5.3 | `el8.aarch64` | pgdg | 75.7 KiB | [pg_repack_15-1.5.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_repack_15-1.5.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_15` | 1.5.2 | `el8.aarch64` | pgdg | 74.3 KiB | [pg_repack_15-1.5.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_repack_15-1.5.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_15` | 1.5.0 | `el8.aarch64` | pgdg | 79.2 KiB | [pg_repack_15-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_repack_15-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_15` | 1.4.8 | `el8.aarch64` | pgdg | 106.0 KiB | [pg_repack_15-1.4.8-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_repack_15-1.4.8-1.rhel8.aarch64.rpm) |
| `pg_repack_15` | 1.5.3 | `el9.x86_64` | pgdg | 67.3 KiB | [pg_repack_15-1.5.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_repack_15-1.5.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_15` | 1.5.2 | `el9.x86_64` | pgdg | 66.2 KiB | [pg_repack_15-1.5.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_repack_15-1.5.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_15` | 1.5.0 | `el9.x86_64` | pgdg | 68.3 KiB | [pg_repack_15-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_repack_15-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_15` | 1.4.8 | `el9.x86_64` | pgdg | 96.1 KiB | [pg_repack_15-1.4.8-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_repack_15-1.4.8-1.rhel9.x86_64.rpm) |
| `pg_repack_15` | 1.5.3 | `el9.aarch64` | pgdg | 66.7 KiB | [pg_repack_15-1.5.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_repack_15-1.5.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_15` | 1.5.2 | `el9.aarch64` | pgdg | 65.7 KiB | [pg_repack_15-1.5.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_repack_15-1.5.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_15` | 1.5.0 | `el9.aarch64` | pgdg | 67.7 KiB | [pg_repack_15-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_repack_15-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_15` | 1.4.8 | `el9.aarch64` | pgdg | 95.0 KiB | [pg_repack_15-1.4.8-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_repack_15-1.4.8-1.rhel9.aarch64.rpm) |
| `pg_repack_15` | 1.5.3 | `el10.x86_64` | pgdg | 68.3 KiB | [pg_repack_15-1.5.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_repack_15-1.5.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_repack_15` | 1.5.2 | `el10.x86_64` | pgdg | 67.5 KiB | [pg_repack_15-1.5.2-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_repack_15-1.5.2-4PGDG.rhel10.x86_64.rpm) |
| `pg_repack_15` | 1.5.3 | `el10.aarch64` | pgdg | 68.2 KiB | [pg_repack_15-1.5.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_repack_15-1.5.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_repack_15` | 1.5.2 | `el10.aarch64` | pgdg | 67.5 KiB | [pg_repack_15-1.5.2-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_repack_15-1.5.2-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-repack` | 1.5.2 | `d12.x86_64` | pgdg | 102.2 KiB | [postgresql-15-repack_1.5.2-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg12+1_amd64.deb) |
| `postgresql-15-repack` | 1.5.2 | `d12.aarch64` | pgdg | 99.6 KiB | [postgresql-15-repack_1.5.2-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg12+1_arm64.deb) |
| `postgresql-15-repack` | 1.5.2 | `d13.x86_64` | pgdg | 102.3 KiB | [postgresql-15-repack_1.5.2-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg13+1_amd64.deb) |
| `postgresql-15-repack` | 1.5.2 | `d13.aarch64` | pgdg | 100.9 KiB | [postgresql-15-repack_1.5.2-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg13+1_arm64.deb) |
| `postgresql-15-repack` | 1.5.2 | `u22.x86_64` | pgdg | 105.3 KiB | [postgresql-15-repack_1.5.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-repack` | 1.5.2 | `u22.aarch64` | pgdg | 102.1 KiB | [postgresql-15-repack_1.5.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-repack` | 1.5.2 | `u24.x86_64` | pgdg | 99.1 KiB | [postgresql-15-repack_1.5.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-repack` | 1.5.2 | `u24.aarch64` | pgdg | 97.1 KiB | [postgresql-15-repack_1.5.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.2-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_repack_14` | 1.5.3 | `el8.x86_64` | pgdg | 75.1 KiB | [pg_repack_14-1.5.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_repack_14-1.5.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_14` | 1.5.2 | `el8.x86_64` | pgdg | 73.9 KiB | [pg_repack_14-1.5.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_repack_14-1.5.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_14` | 1.5.0 | `el8.x86_64` | pgdg | 78.0 KiB | [pg_repack_14-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_repack_14-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_14` | 1.4.8 | `el8.x86_64` | pgdg | 103.7 KiB | [pg_repack_14-1.4.8-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_repack_14-1.4.8-1.rhel8.x86_64.rpm) |
| `pg_repack_14` | 1.4.7 | `el8.x86_64` | pgdg | 103.7 KiB | [pg_repack_14-1.4.7-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_repack_14-1.4.7-1.rhel8.x86_64.rpm) |
| `pg_repack_14` | 1.5.3 | `el8.aarch64` | pgdg | 74.3 KiB | [pg_repack_14-1.5.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_repack_14-1.5.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_14` | 1.5.2 | `el8.aarch64` | pgdg | 73.0 KiB | [pg_repack_14-1.5.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_repack_14-1.5.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_14` | 1.5.0 | `el8.aarch64` | pgdg | 77.8 KiB | [pg_repack_14-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_repack_14-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_14` | 1.4.8 | `el8.aarch64` | pgdg | 103.7 KiB | [pg_repack_14-1.4.8-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_repack_14-1.4.8-1.rhel8.aarch64.rpm) |
| `pg_repack_14` | 1.5.3 | `el9.x86_64` | pgdg | 66.3 KiB | [pg_repack_14-1.5.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_repack_14-1.5.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_14` | 1.5.2 | `el9.x86_64` | pgdg | 65.3 KiB | [pg_repack_14-1.5.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_repack_14-1.5.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_14` | 1.5.0 | `el9.x86_64` | pgdg | 67.6 KiB | [pg_repack_14-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_repack_14-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_14` | 1.4.8 | `el9.x86_64` | pgdg | 94.2 KiB | [pg_repack_14-1.4.8-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_repack_14-1.4.8-1.rhel9.x86_64.rpm) |
| `pg_repack_14` | 1.4.7 | `el9.x86_64` | pgdg | 93.7 KiB | [pg_repack_14-1.4.7-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_repack_14-1.4.7-1.rhel9.x86_64.rpm) |
| `pg_repack_14` | 1.5.3 | `el9.aarch64` | pgdg | 66.0 KiB | [pg_repack_14-1.5.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_repack_14-1.5.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_14` | 1.5.2 | `el9.aarch64` | pgdg | 65.0 KiB | [pg_repack_14-1.5.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_repack_14-1.5.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_14` | 1.5.0 | `el9.aarch64` | pgdg | 67.1 KiB | [pg_repack_14-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_repack_14-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_14` | 1.4.8 | `el9.aarch64` | pgdg | 93.4 KiB | [pg_repack_14-1.4.8-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_repack_14-1.4.8-1.rhel9.aarch64.rpm) |
| `pg_repack_14` | 1.5.3 | `el10.x86_64` | pgdg | 67.6 KiB | [pg_repack_14-1.5.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_repack_14-1.5.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_repack_14` | 1.5.2 | `el10.x86_64` | pgdg | 66.8 KiB | [pg_repack_14-1.5.2-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_repack_14-1.5.2-4PGDG.rhel10.x86_64.rpm) |
| `pg_repack_14` | 1.5.3 | `el10.aarch64` | pgdg | 67.6 KiB | [pg_repack_14-1.5.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_repack_14-1.5.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_repack_14` | 1.5.2 | `el10.aarch64` | pgdg | 66.9 KiB | [pg_repack_14-1.5.2-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_repack_14-1.5.2-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-repack` | 1.5.2 | `d12.x86_64` | pgdg | 101.4 KiB | [postgresql-14-repack_1.5.2-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg12+1_amd64.deb) |
| `postgresql-14-repack` | 1.5.2 | `d12.aarch64` | pgdg | 99.3 KiB | [postgresql-14-repack_1.5.2-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg12+1_arm64.deb) |
| `postgresql-14-repack` | 1.5.2 | `d13.x86_64` | pgdg | 101.7 KiB | [postgresql-14-repack_1.5.2-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg13+1_amd64.deb) |
| `postgresql-14-repack` | 1.5.2 | `d13.aarch64` | pgdg | 100.2 KiB | [postgresql-14-repack_1.5.2-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg13+1_arm64.deb) |
| `postgresql-14-repack` | 1.5.2 | `u22.x86_64` | pgdg | 103.9 KiB | [postgresql-14-repack_1.5.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-repack` | 1.5.2 | `u22.aarch64` | pgdg | 100.8 KiB | [postgresql-14-repack_1.5.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-repack` | 1.5.2 | `u24.x86_64` | pgdg | 98.4 KiB | [postgresql-14-repack_1.5.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-repack` | 1.5.2 | `u24.aarch64` | pgdg | 96.6 KiB | [postgresql-14-repack_1.5.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.2-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_repack_13` | 1.5.3 | `el8.x86_64` | pgdg | 73.6 KiB | [pg_repack_13-1.5.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_repack_13-1.5.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_13` | 1.5.2 | `el8.x86_64` | pgdg | 72.4 KiB | [pg_repack_13-1.5.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_repack_13-1.5.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_13` | 1.5.0 | `el8.x86_64` | pgdg | 76.0 KiB | [pg_repack_13-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_repack_13-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_13` | 1.4.8 | `el8.x86_64` | pgdg | 101.4 KiB | [pg_repack_13-1.4.8-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_repack_13-1.4.8-1.rhel8.x86_64.rpm) |
| `pg_repack_13` | 1.4.7 | `el8.x86_64` | pgdg | 101.0 KiB | [pg_repack_13-1.4.7-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_repack_13-1.4.7-1.rhel8.x86_64.rpm) |
| `pg_repack_13` | 1.5.3 | `el8.aarch64` | pgdg | 73.1 KiB | [pg_repack_13-1.5.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_repack_13-1.5.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_13` | 1.5.2 | `el8.aarch64` | pgdg | 71.7 KiB | [pg_repack_13-1.5.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_repack_13-1.5.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_13` | 1.5.0 | `el8.aarch64` | pgdg | 76.1 KiB | [pg_repack_13-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_repack_13-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_13` | 1.4.8 | `el8.aarch64` | pgdg | 101.7 KiB | [pg_repack_13-1.4.8-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_repack_13-1.4.8-1.rhel8.aarch64.rpm) |
| `pg_repack_13` | 1.5.3 | `el9.x86_64` | pgdg | 66.1 KiB | [pg_repack_13-1.5.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_repack_13-1.5.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_13` | 1.5.2 | `el9.x86_64` | pgdg | 64.9 KiB | [pg_repack_13-1.5.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_repack_13-1.5.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_13` | 1.5.0 | `el9.x86_64` | pgdg | 67.2 KiB | [pg_repack_13-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_repack_13-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_13` | 1.4.8 | `el9.x86_64` | pgdg | 93.6 KiB | [pg_repack_13-1.4.8-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_repack_13-1.4.8-1.rhel9.x86_64.rpm) |
| `pg_repack_13` | 1.4.7 | `el9.x86_64` | pgdg | 93.2 KiB | [pg_repack_13-1.4.7-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_repack_13-1.4.7-1.rhel9.x86_64.rpm) |
| `pg_repack_13` | 1.5.3 | `el9.aarch64` | pgdg | 65.8 KiB | [pg_repack_13-1.5.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_repack_13-1.5.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_13` | 1.5.2 | `el9.aarch64` | pgdg | 64.7 KiB | [pg_repack_13-1.5.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_repack_13-1.5.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_13` | 1.5.0 | `el9.aarch64` | pgdg | 66.8 KiB | [pg_repack_13-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_repack_13-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_13` | 1.4.8 | `el9.aarch64` | pgdg | 92.7 KiB | [pg_repack_13-1.4.8-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_repack_13-1.4.8-1.rhel9.aarch64.rpm) |
| `pg_repack_13` | 1.5.3 | `el10.x86_64` | pgdg | 67.2 KiB | [pg_repack_13-1.5.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_repack_13-1.5.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_repack_13` | 1.5.2 | `el10.x86_64` | pgdg | 66.4 KiB | [pg_repack_13-1.5.2-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_repack_13-1.5.2-4PGDG.rhel10.x86_64.rpm) |
| `pg_repack_13` | 1.5.3 | `el10.aarch64` | pgdg | 67.4 KiB | [pg_repack_13-1.5.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_repack_13-1.5.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_repack_13` | 1.5.2 | `el10.aarch64` | pgdg | 66.6 KiB | [pg_repack_13-1.5.2-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_repack_13-1.5.2-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-repack` | 1.5.2 | `d12.x86_64` | pgdg | 100.1 KiB | [postgresql-13-repack_1.5.2-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-13-repack_1.5.2-2.pgdg12+1_amd64.deb) |
| `postgresql-13-repack` | 1.5.2 | `d12.aarch64` | pgdg | 97.9 KiB | [postgresql-13-repack_1.5.2-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-13-repack_1.5.2-2.pgdg12+1_arm64.deb) |
| `postgresql-13-repack` | 1.5.2 | `d13.x86_64` | pgdg | 100.4 KiB | [postgresql-13-repack_1.5.2-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-13-repack_1.5.2-2.pgdg13+1_amd64.deb) |
| `postgresql-13-repack` | 1.5.2 | `d13.aarch64` | pgdg | 98.6 KiB | [postgresql-13-repack_1.5.2-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-13-repack_1.5.2-2.pgdg13+1_arm64.deb) |
| `postgresql-13-repack` | 1.5.2 | `u22.x86_64` | pgdg | 102.4 KiB | [postgresql-13-repack_1.5.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-13-repack_1.5.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-repack` | 1.5.2 | `u22.aarch64` | pgdg | 99.7 KiB | [postgresql-13-repack_1.5.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-13-repack_1.5.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-repack` | 1.5.2 | `u24.x86_64` | pgdg | 97.4 KiB | [postgresql-13-repack_1.5.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-13-repack_1.5.2-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-repack` | 1.5.2 | `u24.aarch64` | pgdg | 95.0 KiB | [postgresql-13-repack_1.5.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-13-repack_1.5.2-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

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

