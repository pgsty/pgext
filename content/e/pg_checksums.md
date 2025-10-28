---
title: "pg_checksums"
linkTitle: "pg_checksums"
description: "Activate/deactivate/verify checksums in offline Postgres clusters"
weight: 5110
categories: ["ADMIN"]
width: full
---

Activate/deactivate/verify checksums in offline Postgres clusters


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5110** | {{< badge content="pg_checksums" link="https://github.com/credativ/pg_checksums" >}} | {{< ext "pg_checksums" >}} | `1.3` | {{< category "ADMIN" >}} | {{< license "BSD 2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s---r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_catcheck" >}} {{< ext "amcheck" >}} {{< ext "pg_surgery" >}} {{< ext "pageinspect" >}} {{< ext "pg_visibility" >}} {{< ext "pgstattuple" >}} {{< ext "pg_repack" >}} {{< ext "pg_squeeze" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_checksums" >}} | `1.3` | {{< bg "18" "pg_checksums_18*" "green" >}} {{< bg "17" "pg_checksums_17*" "green" >}} {{< bg "16" "pg_checksums_16*" "green" >}} {{< bg "15" "pg_checksums_15*" "green" >}} {{< bg "14" "pg_checksums_14*" "green" >}} {{< bg "13" "pg_checksums_13*" "green" >}} | `pg_checksums_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pg_checksums" >}} | `1.3` | {{< bg "18" "postgresql-18-pg-checksums" "green" >}} {{< bg "17" "postgresql-17-pg-checksums" "green" >}} {{< bg "16" "postgresql-16-pg-checksums" "green" >}} {{< bg "15" "postgresql-15-pg-checksums" "green" >}} {{< bg "14" "postgresql-14-pg-checksums" "green" >}} {{< bg "13" "postgresql-13-pg-checksums" "green" >}} | `postgresql-$v-pg-checksums` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 1.3" "pg_checksums_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_13 : AVAIL 3" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 1.3" "pg_checksums_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_13 : AVAIL 3" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 1.3" "pg_checksums_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_13 : AVAIL 3" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 1.3" "pg_checksums_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_13 : AVAIL 3" "blue" >}} |
|    `el10.x86_64`    | {{< bg "PGDG 1.3" "pg_checksums_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_13 : AVAIL 2" "blue" >}} |
|    `el10.aarch64`    | {{< bg "PGDG 1.3" "pg_checksums_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "pg_checksums_13 : AVAIL 2" "blue" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 1.3" "postgresql-18-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-13-pg-checksums : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 1.3" "postgresql-18-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-13-pg-checksums : AVAIL 1" "blue" >}} |
|    `d13.x86_64`    | {{< bg "PGDG 1.3" "postgresql-18-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-13-pg-checksums : AVAIL 1" "blue" >}} |
|    `d13.aarch64`    | {{< bg "PGDG 1.3" "postgresql-18-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-13-pg-checksums : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 1.3" "postgresql-18-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-13-pg-checksums : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 1.3" "postgresql-18-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-13-pg-checksums : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 1.3" "postgresql-18-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-13-pg-checksums : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 1.3" "postgresql-18-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-pg-checksums : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-13-pg-checksums : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_checksums_18` | 1.3 | `el8.x86_64` | pgdg | 47.3 KiB | [pg_checksums_18-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_checksums_18-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_checksums_18` | 1.3 | `el8.aarch64` | pgdg | 46.8 KiB | [pg_checksums_18-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_checksums_18-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_checksums_18` | 1.3 | `el9.x86_64` | pgdg | 32.4 KiB | [pg_checksums_18-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_checksums_18-1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_checksums_18` | 1.3 | `el9.aarch64` | pgdg | 39.8 KiB | [pg_checksums_18-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_checksums_18-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_checksums_18` | 1.3 | `el10.x86_64` | pgdg | 33.2 KiB | [pg_checksums_18-1.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_checksums_18-1.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_checksums_18` | 1.3 | `el10.aarch64` | pgdg | 40.6 KiB | [pg_checksums_18-1.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_checksums_18-1.3-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-checksums` | 1.3 | `d12.x86_64` | pgdg | 36.8 KiB | [postgresql-18-pg-checksums_1.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.3-2.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-checksums` | 1.3 | `d12.aarch64` | pgdg | 35.5 KiB | [postgresql-18-pg-checksums_1.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.3-2.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-checksums` | 1.3 | `d13.x86_64` | pgdg | 37.1 KiB | [postgresql-18-pg-checksums_1.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.3-2.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-checksums` | 1.3 | `d13.aarch64` | pgdg | 35.9 KiB | [postgresql-18-pg-checksums_1.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.3-2.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-checksums` | 1.3 | `u22.x86_64` | pgdg | 37.9 KiB | [postgresql-18-pg-checksums_1.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-checksums` | 1.3 | `u22.aarch64` | pgdg | 36.3 KiB | [postgresql-18-pg-checksums_1.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-checksums` | 1.3 | `u24.x86_64` | pgdg | 36.8 KiB | [postgresql-18-pg-checksums_1.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-checksums` | 1.3 | `u24.aarch64` | pgdg | 35.8 KiB | [postgresql-18-pg-checksums_1.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_checksums_17` | 1.3 | `el8.x86_64` | pgdg | 47.4 KiB | [pg_checksums_17-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_checksums_17-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_checksums_17` | 1.3 | `el8.aarch64` | pgdg | 47.0 KiB | [pg_checksums_17-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_checksums_17-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_checksums_17` | 1.3 | `el9.x86_64` | pgdg | 31.8 KiB | [pg_checksums_17-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_checksums_17-1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_checksums_17` | 1.3 | `el9.aarch64` | pgdg | 40.1 KiB | [pg_checksums_17-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_checksums_17-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_checksums_17` | 1.3 | `el10.x86_64` | pgdg | 32.2 KiB | [pg_checksums_17-1.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_checksums_17-1.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_checksums_17` | 1.3 | `el10.aarch64` | pgdg | 40.9 KiB | [pg_checksums_17-1.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_checksums_17-1.3-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-checksums` | 1.3 | `d12.x86_64` | pgdg | 36.4 KiB | [postgresql-17-pg-checksums_1.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.3-2.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-checksums` | 1.3 | `d12.aarch64` | pgdg | 35.7 KiB | [postgresql-17-pg-checksums_1.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.3-2.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-checksums` | 1.3 | `d13.x86_64` | pgdg | 36.7 KiB | [postgresql-17-pg-checksums_1.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.3-2.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-checksums` | 1.3 | `d13.aarch64` | pgdg | 36.2 KiB | [postgresql-17-pg-checksums_1.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.3-2.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-checksums` | 1.3 | `u22.x86_64` | pgdg | 37.5 KiB | [postgresql-17-pg-checksums_1.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-checksums` | 1.3 | `u22.aarch64` | pgdg | 36.5 KiB | [postgresql-17-pg-checksums_1.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-checksums` | 1.3 | `u24.x86_64` | pgdg | 36.4 KiB | [postgresql-17-pg-checksums_1.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-checksums` | 1.3 | `u24.aarch64` | pgdg | 36.1 KiB | [postgresql-17-pg-checksums_1.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_checksums_16` | 1.3 | `el8.x86_64` | pgdg | 45.2 KiB | [pg_checksums_16-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_checksums_16-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_checksums_16` | 1.1 | `el8.x86_64` | pgdg | 45.5 KiB | [pg_checksums_16-1.1-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_checksums_16-1.1-3PGDG.rhel8.x86_64.rpm) |
| `pg_checksums_16` | 1.3 | `el8.aarch64` | pgdg | 44.7 KiB | [pg_checksums_16-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_checksums_16-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_checksums_16` | 1.1 | `el8.aarch64` | pgdg | 45.1 KiB | [pg_checksums_16-1.1-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_checksums_16-1.1-3PGDG.rhel8.aarch64.rpm) |
| `pg_checksums_16` | 1.3 | `el9.x86_64` | pgdg | 31.1 KiB | [pg_checksums_16-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_checksums_16-1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_checksums_16` | 1.1 | `el9.x86_64` | pgdg | 31.2 KiB | [pg_checksums_16-1.1-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_checksums_16-1.1-3PGDG.rhel9.x86_64.rpm) |
| `pg_checksums_16` | 1.3 | `el9.aarch64` | pgdg | 39.4 KiB | [pg_checksums_16-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_checksums_16-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_checksums_16` | 1.1 | `el9.aarch64` | pgdg | 39.7 KiB | [pg_checksums_16-1.1-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_checksums_16-1.1-3PGDG.rhel9.aarch64.rpm) |
| `pg_checksums_16` | 1.3 | `el10.x86_64` | pgdg | 31.6 KiB | [pg_checksums_16-1.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_checksums_16-1.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_checksums_16` | 1.2 | `el10.x86_64` | pgdg | 31.6 KiB | [pg_checksums_16-1.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_checksums_16-1.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_checksums_16` | 1.3 | `el10.aarch64` | pgdg | 40.2 KiB | [pg_checksums_16-1.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_checksums_16-1.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_checksums_16` | 1.2 | `el10.aarch64` | pgdg | 40.3 KiB | [pg_checksums_16-1.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_checksums_16-1.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-checksums` | 1.3 | `d12.x86_64` | pgdg | 34.3 KiB | [postgresql-16-pg-checksums_1.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.3-2.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-checksums` | 1.3 | `d12.aarch64` | pgdg | 33.7 KiB | [postgresql-16-pg-checksums_1.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.3-2.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-checksums` | 1.3 | `d13.x86_64` | pgdg | 34.7 KiB | [postgresql-16-pg-checksums_1.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.3-2.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-checksums` | 1.3 | `d13.aarch64` | pgdg | 34.2 KiB | [postgresql-16-pg-checksums_1.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.3-2.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-checksums` | 1.3 | `u22.x86_64` | pgdg | 35.6 KiB | [postgresql-16-pg-checksums_1.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-checksums` | 1.3 | `u22.aarch64` | pgdg | 34.6 KiB | [postgresql-16-pg-checksums_1.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-checksums` | 1.3 | `u24.x86_64` | pgdg | 34.5 KiB | [postgresql-16-pg-checksums_1.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-checksums` | 1.3 | `u24.aarch64` | pgdg | 34.1 KiB | [postgresql-16-pg-checksums_1.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_checksums_15` | 1.3 | `el8.x86_64` | pgdg | 44.8 KiB | [pg_checksums_15-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_checksums_15-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_checksums_15` | 1.1 | `el8.x86_64` | pgdg | 45.0 KiB | [pg_checksums_15-1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_checksums_15-1.1-1.rhel8.x86_64.rpm) |
| `pg_checksums_15` | 1.1 | `el8.x86_64` | pgdg | 45.1 KiB | [pg_checksums_15-1.1-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_checksums_15-1.1-3PGDG.rhel8.x86_64.rpm) |
| `pg_checksums_15` | 1.3 | `el8.aarch64` | pgdg | 44.3 KiB | [pg_checksums_15-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_checksums_15-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_checksums_15` | 1.1 | `el8.aarch64` | pgdg | 44.7 KiB | [pg_checksums_15-1.1-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_checksums_15-1.1-3PGDG.rhel8.aarch64.rpm) |
| `pg_checksums_15` | 1.1 | `el8.aarch64` | pgdg | 44.5 KiB | [pg_checksums_15-1.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_checksums_15-1.1-1.rhel8.aarch64.rpm) |
| `pg_checksums_15` | 1.3 | `el9.x86_64` | pgdg | 31.2 KiB | [pg_checksums_15-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_checksums_15-1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_checksums_15` | 1.1 | `el9.x86_64` | pgdg | 31.4 KiB | [pg_checksums_15-1.1-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_checksums_15-1.1-3PGDG.rhel9.x86_64.rpm) |
| `pg_checksums_15` | 1.1 | `el9.x86_64` | pgdg | 31.8 KiB | [pg_checksums_15-1.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_checksums_15-1.1-1.rhel9.x86_64.rpm) |
| `pg_checksums_15` | 1.3 | `el9.aarch64` | pgdg | 39.8 KiB | [pg_checksums_15-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_checksums_15-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_checksums_15` | 1.1 | `el9.aarch64` | pgdg | 40.0 KiB | [pg_checksums_15-1.1-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_checksums_15-1.1-3PGDG.rhel9.aarch64.rpm) |
| `pg_checksums_15` | 1.1 | `el9.aarch64` | pgdg | 40.4 KiB | [pg_checksums_15-1.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_checksums_15-1.1-1.rhel9.aarch64.rpm) |
| `pg_checksums_15` | 1.3 | `el10.x86_64` | pgdg | 31.8 KiB | [pg_checksums_15-1.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_checksums_15-1.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_checksums_15` | 1.2 | `el10.x86_64` | pgdg | 31.9 KiB | [pg_checksums_15-1.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_checksums_15-1.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_checksums_15` | 1.3 | `el10.aarch64` | pgdg | 40.7 KiB | [pg_checksums_15-1.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_checksums_15-1.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_checksums_15` | 1.2 | `el10.aarch64` | pgdg | 40.7 KiB | [pg_checksums_15-1.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_checksums_15-1.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-checksums` | 1.3 | `d12.x86_64` | pgdg | 34.0 KiB | [postgresql-15-pg-checksums_1.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.3-2.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-checksums` | 1.3 | `d12.aarch64` | pgdg | 33.4 KiB | [postgresql-15-pg-checksums_1.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.3-2.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-checksums` | 1.3 | `d13.x86_64` | pgdg | 34.4 KiB | [postgresql-15-pg-checksums_1.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.3-2.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-checksums` | 1.3 | `d13.aarch64` | pgdg | 34.0 KiB | [postgresql-15-pg-checksums_1.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.3-2.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-checksums` | 1.3 | `u22.x86_64` | pgdg | 35.2 KiB | [postgresql-15-pg-checksums_1.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-checksums` | 1.3 | `u22.aarch64` | pgdg | 34.3 KiB | [postgresql-15-pg-checksums_1.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-checksums` | 1.3 | `u24.x86_64` | pgdg | 34.2 KiB | [postgresql-15-pg-checksums_1.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-checksums` | 1.3 | `u24.aarch64` | pgdg | 33.7 KiB | [postgresql-15-pg-checksums_1.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_checksums_14` | 1.3 | `el8.x86_64` | pgdg | 43.4 KiB | [pg_checksums_14-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_checksums_14-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_checksums_14` | 1.1 | `el8.x86_64` | pgdg | 43.6 KiB | [pg_checksums_14-1.1-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_checksums_14-1.1-3PGDG.rhel8.x86_64.rpm) |
| `pg_checksums_14` | 1.1 | `el8.x86_64` | pgdg | 43.5 KiB | [pg_checksums_14-1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_checksums_14-1.1-1.rhel8.x86_64.rpm) |
| `pg_checksums_14` | 1.3 | `el8.aarch64` | pgdg | 43.1 KiB | [pg_checksums_14-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_checksums_14-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_checksums_14` | 1.1 | `el8.aarch64` | pgdg | 43.4 KiB | [pg_checksums_14-1.1-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_checksums_14-1.1-3PGDG.rhel8.aarch64.rpm) |
| `pg_checksums_14` | 1.1 | `el8.aarch64` | pgdg | 43.2 KiB | [pg_checksums_14-1.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_checksums_14-1.1-1.rhel8.aarch64.rpm) |
| `pg_checksums_14` | 1.3 | `el9.x86_64` | pgdg | 30.8 KiB | [pg_checksums_14-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_checksums_14-1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_checksums_14` | 1.1 | `el9.x86_64` | pgdg | 31.0 KiB | [pg_checksums_14-1.1-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_checksums_14-1.1-3PGDG.rhel9.x86_64.rpm) |
| `pg_checksums_14` | 1.1 | `el9.x86_64` | pgdg | 31.3 KiB | [pg_checksums_14-1.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_checksums_14-1.1-1.rhel9.x86_64.rpm) |
| `pg_checksums_14` | 1.3 | `el9.aarch64` | pgdg | 39.5 KiB | [pg_checksums_14-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_checksums_14-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_checksums_14` | 1.1 | `el9.aarch64` | pgdg | 39.7 KiB | [pg_checksums_14-1.1-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_checksums_14-1.1-3PGDG.rhel9.aarch64.rpm) |
| `pg_checksums_14` | 1.1 | `el9.aarch64` | pgdg | 40.1 KiB | [pg_checksums_14-1.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_checksums_14-1.1-1.rhel9.aarch64.rpm) |
| `pg_checksums_14` | 1.3 | `el10.x86_64` | pgdg | 31.6 KiB | [pg_checksums_14-1.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_checksums_14-1.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_checksums_14` | 1.2 | `el10.x86_64` | pgdg | 31.5 KiB | [pg_checksums_14-1.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_checksums_14-1.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_checksums_14` | 1.3 | `el10.aarch64` | pgdg | 40.3 KiB | [pg_checksums_14-1.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_checksums_14-1.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_checksums_14` | 1.2 | `el10.aarch64` | pgdg | 40.4 KiB | [pg_checksums_14-1.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_checksums_14-1.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-checksums` | 1.3 | `d12.x86_64` | pgdg | 33.9 KiB | [postgresql-14-pg-checksums_1.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-14-pg-checksums_1.3-2.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-checksums` | 1.3 | `d12.aarch64` | pgdg | 33.3 KiB | [postgresql-14-pg-checksums_1.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-14-pg-checksums_1.3-2.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-checksums` | 1.3 | `d13.x86_64` | pgdg | 34.1 KiB | [postgresql-14-pg-checksums_1.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-14-pg-checksums_1.3-2.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-checksums` | 1.3 | `d13.aarch64` | pgdg | 33.7 KiB | [postgresql-14-pg-checksums_1.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-14-pg-checksums_1.3-2.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-checksums` | 1.3 | `u22.x86_64` | pgdg | 34.8 KiB | [postgresql-14-pg-checksums_1.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-14-pg-checksums_1.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-checksums` | 1.3 | `u22.aarch64` | pgdg | 33.9 KiB | [postgresql-14-pg-checksums_1.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-14-pg-checksums_1.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-checksums` | 1.3 | `u24.x86_64` | pgdg | 34.1 KiB | [postgresql-14-pg-checksums_1.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-14-pg-checksums_1.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-checksums` | 1.3 | `u24.aarch64` | pgdg | 33.7 KiB | [postgresql-14-pg-checksums_1.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-14-pg-checksums_1.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_checksums_13` | 1.3 | `el8.x86_64` | pgdg | 43.2 KiB | [pg_checksums_13-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_checksums_13-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_checksums_13` | 1.1 | `el8.x86_64` | pgdg | 43.6 KiB | [pg_checksums_13-1.1-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_checksums_13-1.1-3PGDG.rhel8.x86_64.rpm) |
| `pg_checksums_13` | 1.1 | `el8.x86_64` | pgdg | 43.4 KiB | [pg_checksums_13-1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_checksums_13-1.1-1.rhel8.x86_64.rpm) |
| `pg_checksums_13` | 1.3 | `el8.aarch64` | pgdg | 43.0 KiB | [pg_checksums_13-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_checksums_13-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_checksums_13` | 1.1 | `el8.aarch64` | pgdg | 43.3 KiB | [pg_checksums_13-1.1-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_checksums_13-1.1-3PGDG.rhel8.aarch64.rpm) |
| `pg_checksums_13` | 1.1 | `el8.aarch64` | pgdg | 43.1 KiB | [pg_checksums_13-1.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_checksums_13-1.1-1.rhel8.aarch64.rpm) |
| `pg_checksums_13` | 1.3 | `el9.x86_64` | pgdg | 30.9 KiB | [pg_checksums_13-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_checksums_13-1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_checksums_13` | 1.1 | `el9.x86_64` | pgdg | 31.0 KiB | [pg_checksums_13-1.1-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_checksums_13-1.1-3PGDG.rhel9.x86_64.rpm) |
| `pg_checksums_13` | 1.1 | `el9.x86_64` | pgdg | 31.4 KiB | [pg_checksums_13-1.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_checksums_13-1.1-1.rhel9.x86_64.rpm) |
| `pg_checksums_13` | 1.3 | `el9.aarch64` | pgdg | 39.5 KiB | [pg_checksums_13-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_checksums_13-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_checksums_13` | 1.1 | `el9.aarch64` | pgdg | 40.2 KiB | [pg_checksums_13-1.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_checksums_13-1.1-1.rhel9.aarch64.rpm) |
| `pg_checksums_13` | 1.1 | `el9.aarch64` | pgdg | 39.8 KiB | [pg_checksums_13-1.1-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_checksums_13-1.1-3PGDG.rhel9.aarch64.rpm) |
| `pg_checksums_13` | 1.3 | `el10.x86_64` | pgdg | 31.6 KiB | [pg_checksums_13-1.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_checksums_13-1.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_checksums_13` | 1.2 | `el10.x86_64` | pgdg | 31.6 KiB | [pg_checksums_13-1.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_checksums_13-1.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_checksums_13` | 1.3 | `el10.aarch64` | pgdg | 40.4 KiB | [pg_checksums_13-1.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_checksums_13-1.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_checksums_13` | 1.2 | `el10.aarch64` | pgdg | 40.4 KiB | [pg_checksums_13-1.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_checksums_13-1.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pg-checksums` | 1.3 | `d12.x86_64` | pgdg | 33.8 KiB | [postgresql-13-pg-checksums_1.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-13-pg-checksums_1.3-2.pgdg12+1_amd64.deb) |
| `postgresql-13-pg-checksums` | 1.3 | `d12.aarch64` | pgdg | 33.2 KiB | [postgresql-13-pg-checksums_1.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-13-pg-checksums_1.3-2.pgdg12+1_arm64.deb) |
| `postgresql-13-pg-checksums` | 1.3 | `d13.x86_64` | pgdg | 34.1 KiB | [postgresql-13-pg-checksums_1.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-13-pg-checksums_1.3-2.pgdg13+1_amd64.deb) |
| `postgresql-13-pg-checksums` | 1.3 | `d13.aarch64` | pgdg | 33.6 KiB | [postgresql-13-pg-checksums_1.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-13-pg-checksums_1.3-2.pgdg13+1_arm64.deb) |
| `postgresql-13-pg-checksums` | 1.3 | `u22.x86_64` | pgdg | 34.8 KiB | [postgresql-13-pg-checksums_1.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-13-pg-checksums_1.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pg-checksums` | 1.3 | `u22.aarch64` | pgdg | 33.9 KiB | [postgresql-13-pg-checksums_1.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-13-pg-checksums_1.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pg-checksums` | 1.3 | `u24.x86_64` | pgdg | 34.1 KiB | [postgresql-13-pg-checksums_1.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-13-pg-checksums_1.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pg-checksums` | 1.3 | `u24.aarch64` | pgdg | 33.6 KiB | [postgresql-13-pg-checksums_1.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-13-pg-checksums_1.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/credativ/pg_checksums" title="Repository" icon="github" subtitle="github.com/credativ/pg_checksums" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_checksums-1.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_checksums; # get pg_checksums source code
pig build dep pg_checksums; # install build dependencies
pig build pkg pg_checksums; # build extension rpm or deb
pig build ext pg_checksums; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_checksums; # install by extension name, for the current active PG version
pig ext install pg_checksums; # install via package alias, for the active PG version
pig ext install pg_checksums -v 18;   # install for PG 18
pig ext install pg_checksums -v 17;   # install for PG 17
pig ext install pg_checksums -v 16;   # install for PG 16
pig ext install pg_checksums -v 15;   # install for PG 15
pig ext install pg_checksums -v 14;   # install for PG 14
pig ext install pg_checksums -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_checksums;
```

