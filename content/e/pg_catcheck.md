---
title: "pg_catcheck"
linkTitle: "pg_catcheck"
description: "Diagnosing system catalog corruption"
weight: 5160
categories: ["ADMIN"]
width: full
---

Diagnosing system catalog corruption


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5160** | {{< badge content="pg_catcheck" link="https://github.com/EnterpriseDB/pg_catcheck" >}} | {{< ext "pg_catcheck" >}} | `1.6.0` | {{< category "ADMIN" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_checksums" >}} {{< ext "amcheck" >}} {{< ext "pg_surgery" >}} {{< ext "pageinspect" >}} {{< ext "pg_visibility" >}} {{< ext "pgstattuple" >}} {{< ext "ddlx" >}} {{< ext "pgdd" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_catcheck" >}} | `1.6.0` | {{< bg "18" "pg_catcheck_18*" "green" >}} {{< bg "17" "pg_catcheck_17*" "green" >}} {{< bg "16" "pg_catcheck_16*" "green" >}} {{< bg "15" "pg_catcheck_15*" "green" >}} {{< bg "14" "pg_catcheck_14*" "green" >}} {{< bg "13" "pg_catcheck_13*" "green" >}} | `pg_catcheck_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pg_catcheck" >}} | `1.6.0` | {{< bg "18" "postgresql-18-pg-catcheck" "green" >}} {{< bg "17" "postgresql-17-pg-catcheck" "green" >}} {{< bg "16" "postgresql-16-pg-catcheck" "green" >}} {{< bg "15" "postgresql-15-pg-catcheck" "green" >}} {{< bg "14" "postgresql-14-pg-catcheck" "green" >}} {{< bg "13" "postgresql-13-pg-catcheck" "green" >}} | `postgresql-$v-pg-catcheck` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 1.6.0" "pg_catcheck_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_13 : AVAIL 2" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 1.6.0" "pg_catcheck_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_13 : AVAIL 2" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 1.6.0" "pg_catcheck_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_13 : AVAIL 2" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 1.6.0" "pg_catcheck_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_13 : AVAIL 2" "blue" >}} |
|    `el10.x86_64`    | {{< bg "PGDG 1.6.0" "pg_catcheck_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_13 : AVAIL 1" "blue" >}} |
|    `el10.aarch64`    | {{< bg "PGDG 1.6.0" "pg_catcheck_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_catcheck_13 : AVAIL 1" "blue" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 1.6.0" "postgresql-18-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-17-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-15-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-14-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-13-pg-catcheck : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 1.6.0" "postgresql-18-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-17-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-15-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-14-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-13-pg-catcheck : AVAIL 1" "blue" >}} |
|    `d13.x86_64`    | {{< bg "PGDG 1.6.0" "postgresql-18-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-17-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-15-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-14-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-13-pg-catcheck : AVAIL 1" "blue" >}} |
|    `d13.aarch64`    | {{< bg "PGDG 1.6.0" "postgresql-18-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-17-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-15-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-14-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-13-pg-catcheck : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 1.6.0" "postgresql-18-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-17-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-15-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-14-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-13-pg-catcheck : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 1.6.0" "postgresql-18-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-17-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-15-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-14-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-13-pg-catcheck : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 1.6.0" "postgresql-18-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-17-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-15-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-14-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-13-pg-catcheck : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 1.6.0" "postgresql-18-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-17-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-15-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-14-pg-catcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-13-pg-catcheck : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_catcheck_18` | 1.6.0 | `el8.x86_64` | pgdg | 43.4 KiB | [pg_catcheck_18-1.6.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_catcheck_18-1.6.0-3PGDG.rhel8.x86_64.rpm) |
| `pg_catcheck_18` | 1.6.0 | `el8.aarch64` | pgdg | 42.6 KiB | [pg_catcheck_18-1.6.0-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_catcheck_18-1.6.0-3PGDG.rhel8.aarch64.rpm) |
| `pg_catcheck_18` | 1.6.0 | `el9.x86_64` | pgdg | 37.8 KiB | [pg_catcheck_18-1.6.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_catcheck_18-1.6.0-3PGDG.rhel9.x86_64.rpm) |
| `pg_catcheck_18` | 1.6.0 | `el9.aarch64` | pgdg | 37.2 KiB | [pg_catcheck_18-1.6.0-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_catcheck_18-1.6.0-3PGDG.rhel9.aarch64.rpm) |
| `pg_catcheck_18` | 1.6.0 | `el10.x86_64` | pgdg | 38.9 KiB | [pg_catcheck_18-1.6.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_catcheck_18-1.6.0-3PGDG.rhel10.x86_64.rpm) |
| `pg_catcheck_18` | 1.6.0 | `el10.aarch64` | pgdg | 38.6 KiB | [pg_catcheck_18-1.6.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_catcheck_18-1.6.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-catcheck` | 1.6.0 | `d12.x86_64` | pgdg | 34.4 KiB | [postgresql-18-pg-catcheck_1.6.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-18-pg-catcheck_1.6.0-2.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-catcheck` | 1.6.0 | `d12.aarch64` | pgdg | 33.3 KiB | [postgresql-18-pg-catcheck_1.6.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-18-pg-catcheck_1.6.0-2.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-catcheck` | 1.6.0 | `d13.x86_64` | pgdg | 34.6 KiB | [postgresql-18-pg-catcheck_1.6.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-18-pg-catcheck_1.6.0-2.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-catcheck` | 1.6.0 | `d13.aarch64` | pgdg | 33.7 KiB | [postgresql-18-pg-catcheck_1.6.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-18-pg-catcheck_1.6.0-2.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-catcheck` | 1.6.0 | `u22.x86_64` | pgdg | 35.4 KiB | [postgresql-18-pg-catcheck_1.6.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-18-pg-catcheck_1.6.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-catcheck` | 1.6.0 | `u22.aarch64` | pgdg | 33.3 KiB | [postgresql-18-pg-catcheck_1.6.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-18-pg-catcheck_1.6.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-catcheck` | 1.6.0 | `u24.x86_64` | pgdg | 35.3 KiB | [postgresql-18-pg-catcheck_1.6.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-18-pg-catcheck_1.6.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-catcheck` | 1.6.0 | `u24.aarch64` | pgdg | 33.6 KiB | [postgresql-18-pg-catcheck_1.6.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-18-pg-catcheck_1.6.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_catcheck_17` | 1.6.0 | `el8.x86_64` | pgdg | 43.3 KiB | [pg_catcheck_17-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_catcheck_17-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_catcheck_17` | 1.5.0 | `el8.x86_64` | pgdg | 43.1 KiB | [pg_catcheck_17-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_catcheck_17-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_catcheck_17` | 1.6.0 | `el8.aarch64` | pgdg | 42.5 KiB | [pg_catcheck_17-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_catcheck_17-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_catcheck_17` | 1.5.0 | `el8.aarch64` | pgdg | 42.4 KiB | [pg_catcheck_17-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_catcheck_17-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_catcheck_17` | 1.6.0 | `el9.x86_64` | pgdg | 37.9 KiB | [pg_catcheck_17-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_catcheck_17-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_catcheck_17` | 1.5.0 | `el9.x86_64` | pgdg | 37.7 KiB | [pg_catcheck_17-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_catcheck_17-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_catcheck_17` | 1.6.0 | `el9.aarch64` | pgdg | 37.3 KiB | [pg_catcheck_17-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_catcheck_17-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_catcheck_17` | 1.5.0 | `el9.aarch64` | pgdg | 37.1 KiB | [pg_catcheck_17-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_catcheck_17-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_catcheck_17` | 1.6.0 | `el10.x86_64` | pgdg | 38.9 KiB | [pg_catcheck_17-1.6.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_catcheck_17-1.6.0-3PGDG.rhel10.x86_64.rpm) |
| `pg_catcheck_17` | 1.6.0 | `el10.aarch64` | pgdg | 38.6 KiB | [pg_catcheck_17-1.6.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_catcheck_17-1.6.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-catcheck` | 1.6.0 | `d12.x86_64` | pgdg | 34.4 KiB | [postgresql-17-pg-catcheck_1.6.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-17-pg-catcheck_1.6.0-2.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-catcheck` | 1.6.0 | `d12.aarch64` | pgdg | 33.4 KiB | [postgresql-17-pg-catcheck_1.6.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-17-pg-catcheck_1.6.0-2.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-catcheck` | 1.6.0 | `d13.x86_64` | pgdg | 34.8 KiB | [postgresql-17-pg-catcheck_1.6.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-17-pg-catcheck_1.6.0-2.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-catcheck` | 1.6.0 | `d13.aarch64` | pgdg | 33.7 KiB | [postgresql-17-pg-catcheck_1.6.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-17-pg-catcheck_1.6.0-2.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-catcheck` | 1.6.0 | `u22.x86_64` | pgdg | 35.6 KiB | [postgresql-17-pg-catcheck_1.6.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-17-pg-catcheck_1.6.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-catcheck` | 1.6.0 | `u22.aarch64` | pgdg | 33.6 KiB | [postgresql-17-pg-catcheck_1.6.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-17-pg-catcheck_1.6.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-catcheck` | 1.6.0 | `u24.x86_64` | pgdg | 35.3 KiB | [postgresql-17-pg-catcheck_1.6.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-17-pg-catcheck_1.6.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-catcheck` | 1.6.0 | `u24.aarch64` | pgdg | 33.7 KiB | [postgresql-17-pg-catcheck_1.6.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-17-pg-catcheck_1.6.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_catcheck_16` | 1.6.0 | `el8.x86_64` | pgdg | 43.3 KiB | [pg_catcheck_16-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_catcheck_16-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_catcheck_16` | 1.4.0 | `el8.x86_64` | pgdg | 43.1 KiB | [pg_catcheck_16-1.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_catcheck_16-1.4.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_catcheck_16` | 1.6.0 | `el8.aarch64` | pgdg | 42.6 KiB | [pg_catcheck_16-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_catcheck_16-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_catcheck_16` | 1.4.0 | `el8.aarch64` | pgdg | 42.3 KiB | [pg_catcheck_16-1.4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_catcheck_16-1.4.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_catcheck_16` | 1.6.0 | `el9.x86_64` | pgdg | 37.9 KiB | [pg_catcheck_16-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_catcheck_16-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_catcheck_16` | 1.4.0 | `el9.x86_64` | pgdg | 37.5 KiB | [pg_catcheck_16-1.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_catcheck_16-1.4.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_catcheck_16` | 1.6.0 | `el9.aarch64` | pgdg | 37.3 KiB | [pg_catcheck_16-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_catcheck_16-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_catcheck_16` | 1.4.0 | `el9.aarch64` | pgdg | 36.9 KiB | [pg_catcheck_16-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_catcheck_16-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_catcheck_16` | 1.6.0 | `el10.x86_64` | pgdg | 38.9 KiB | [pg_catcheck_16-1.6.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_catcheck_16-1.6.0-3PGDG.rhel10.x86_64.rpm) |
| `pg_catcheck_16` | 1.6.0 | `el10.aarch64` | pgdg | 38.6 KiB | [pg_catcheck_16-1.6.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_catcheck_16-1.6.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-catcheck` | 1.6.0 | `d12.x86_64` | pgdg | 34.4 KiB | [postgresql-16-pg-catcheck_1.6.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-16-pg-catcheck_1.6.0-2.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-catcheck` | 1.6.0 | `d12.aarch64` | pgdg | 33.4 KiB | [postgresql-16-pg-catcheck_1.6.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-16-pg-catcheck_1.6.0-2.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-catcheck` | 1.6.0 | `d13.x86_64` | pgdg | 34.8 KiB | [postgresql-16-pg-catcheck_1.6.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-16-pg-catcheck_1.6.0-2.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-catcheck` | 1.6.0 | `d13.aarch64` | pgdg | 34.1 KiB | [postgresql-16-pg-catcheck_1.6.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-16-pg-catcheck_1.6.0-2.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-catcheck` | 1.6.0 | `u22.x86_64` | pgdg | 35.5 KiB | [postgresql-16-pg-catcheck_1.6.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-16-pg-catcheck_1.6.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-catcheck` | 1.6.0 | `u22.aarch64` | pgdg | 33.6 KiB | [postgresql-16-pg-catcheck_1.6.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-16-pg-catcheck_1.6.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-catcheck` | 1.6.0 | `u24.x86_64` | pgdg | 35.3 KiB | [postgresql-16-pg-catcheck_1.6.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-16-pg-catcheck_1.6.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-catcheck` | 1.6.0 | `u24.aarch64` | pgdg | 33.7 KiB | [postgresql-16-pg-catcheck_1.6.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-16-pg-catcheck_1.6.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_catcheck_15` | 1.6.0 | `el8.x86_64` | pgdg | 43.2 KiB | [pg_catcheck_15-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_catcheck_15-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_catcheck_15` | 1.3.0 | `el8.x86_64` | pgdg | 42.8 KiB | [pg_catcheck_15-1.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_catcheck_15-1.3.0-1.rhel8.x86_64.rpm) |
| `pg_catcheck_15` | 1.6.0 | `el8.aarch64` | pgdg | 42.5 KiB | [pg_catcheck_15-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_catcheck_15-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_catcheck_15` | 1.3.0 | `el8.aarch64` | pgdg | 41.9 KiB | [pg_catcheck_15-1.3.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_catcheck_15-1.3.0-1.rhel8.aarch64.rpm) |
| `pg_catcheck_15` | 1.6.0 | `el9.x86_64` | pgdg | 37.9 KiB | [pg_catcheck_15-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_catcheck_15-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_catcheck_15` | 1.3.0 | `el9.x86_64` | pgdg | 37.7 KiB | [pg_catcheck_15-1.3.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_catcheck_15-1.3.0-1.rhel9.x86_64.rpm) |
| `pg_catcheck_15` | 1.6.0 | `el9.aarch64` | pgdg | 37.3 KiB | [pg_catcheck_15-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_catcheck_15-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_catcheck_15` | 1.3.0 | `el9.aarch64` | pgdg | 37.0 KiB | [pg_catcheck_15-1.3.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_catcheck_15-1.3.0-1.rhel9.aarch64.rpm) |
| `pg_catcheck_15` | 1.6.0 | `el10.x86_64` | pgdg | 38.9 KiB | [pg_catcheck_15-1.6.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_catcheck_15-1.6.0-3PGDG.rhel10.x86_64.rpm) |
| `pg_catcheck_15` | 1.6.0 | `el10.aarch64` | pgdg | 38.6 KiB | [pg_catcheck_15-1.6.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_catcheck_15-1.6.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-catcheck` | 1.6.0 | `d12.x86_64` | pgdg | 34.4 KiB | [postgresql-15-pg-catcheck_1.6.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-15-pg-catcheck_1.6.0-2.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-catcheck` | 1.6.0 | `d12.aarch64` | pgdg | 33.4 KiB | [postgresql-15-pg-catcheck_1.6.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-15-pg-catcheck_1.6.0-2.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-catcheck` | 1.6.0 | `d13.x86_64` | pgdg | 34.8 KiB | [postgresql-15-pg-catcheck_1.6.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-15-pg-catcheck_1.6.0-2.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-catcheck` | 1.6.0 | `d13.aarch64` | pgdg | 33.9 KiB | [postgresql-15-pg-catcheck_1.6.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-15-pg-catcheck_1.6.0-2.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-catcheck` | 1.6.0 | `u22.x86_64` | pgdg | 35.5 KiB | [postgresql-15-pg-catcheck_1.6.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-15-pg-catcheck_1.6.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-catcheck` | 1.6.0 | `u22.aarch64` | pgdg | 33.6 KiB | [postgresql-15-pg-catcheck_1.6.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-15-pg-catcheck_1.6.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-catcheck` | 1.6.0 | `u24.x86_64` | pgdg | 35.2 KiB | [postgresql-15-pg-catcheck_1.6.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-15-pg-catcheck_1.6.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-catcheck` | 1.6.0 | `u24.aarch64` | pgdg | 33.6 KiB | [postgresql-15-pg-catcheck_1.6.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-15-pg-catcheck_1.6.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_catcheck_14` | 1.6.0 | `el8.x86_64` | pgdg | 41.6 KiB | [pg_catcheck_14-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_catcheck_14-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_catcheck_14` | 1.3.0 | `el8.x86_64` | pgdg | 41.0 KiB | [pg_catcheck_14-1.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_catcheck_14-1.3.0-1.rhel8.x86_64.rpm) |
| `pg_catcheck_14` | 1.2.0 | `el8.x86_64` | pgdg | 40.9 KiB | [pg_catcheck_14-1.2.0-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_catcheck_14-1.2.0-3.rhel8.x86_64.rpm) |
| `pg_catcheck_14` | 1.6.0 | `el8.aarch64` | pgdg | 41.2 KiB | [pg_catcheck_14-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_catcheck_14-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_catcheck_14` | 1.3.0 | `el8.aarch64` | pgdg | 40.6 KiB | [pg_catcheck_14-1.3.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_catcheck_14-1.3.0-1.rhel8.aarch64.rpm) |
| `pg_catcheck_14` | 1.6.0 | `el9.x86_64` | pgdg | 37.6 KiB | [pg_catcheck_14-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_catcheck_14-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_catcheck_14` | 1.3.0 | `el9.x86_64` | pgdg | 37.7 KiB | [pg_catcheck_14-1.3.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_catcheck_14-1.3.0-1.rhel9.x86_64.rpm) |
| `pg_catcheck_14` | 1.6.0 | `el9.aarch64` | pgdg | 37.0 KiB | [pg_catcheck_14-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_catcheck_14-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_catcheck_14` | 1.3.0 | `el9.aarch64` | pgdg | 37.1 KiB | [pg_catcheck_14-1.3.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_catcheck_14-1.3.0-1.rhel9.aarch64.rpm) |
| `pg_catcheck_14` | 1.6.0 | `el10.x86_64` | pgdg | 38.4 KiB | [pg_catcheck_14-1.6.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_catcheck_14-1.6.0-3PGDG.rhel10.x86_64.rpm) |
| `pg_catcheck_14` | 1.6.0 | `el10.aarch64` | pgdg | 38.1 KiB | [pg_catcheck_14-1.6.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_catcheck_14-1.6.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-catcheck` | 1.6.0 | `d12.x86_64` | pgdg | 33.9 KiB | [postgresql-14-pg-catcheck_1.6.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-14-pg-catcheck_1.6.0-2.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-catcheck` | 1.6.0 | `d12.aarch64` | pgdg | 32.9 KiB | [postgresql-14-pg-catcheck_1.6.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-14-pg-catcheck_1.6.0-2.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-catcheck` | 1.6.0 | `d13.x86_64` | pgdg | 34.0 KiB | [postgresql-14-pg-catcheck_1.6.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-14-pg-catcheck_1.6.0-2.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-catcheck` | 1.6.0 | `d13.aarch64` | pgdg | 33.2 KiB | [postgresql-14-pg-catcheck_1.6.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-14-pg-catcheck_1.6.0-2.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-catcheck` | 1.6.0 | `u22.x86_64` | pgdg | 34.6 KiB | [postgresql-14-pg-catcheck_1.6.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-14-pg-catcheck_1.6.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-catcheck` | 1.6.0 | `u22.aarch64` | pgdg | 33.0 KiB | [postgresql-14-pg-catcheck_1.6.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-14-pg-catcheck_1.6.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-catcheck` | 1.6.0 | `u24.x86_64` | pgdg | 34.5 KiB | [postgresql-14-pg-catcheck_1.6.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-14-pg-catcheck_1.6.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-catcheck` | 1.6.0 | `u24.aarch64` | pgdg | 33.2 KiB | [postgresql-14-pg-catcheck_1.6.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-14-pg-catcheck_1.6.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_catcheck_13` | 1.6.0 | `el8.x86_64` | pgdg | 40.1 KiB | [pg_catcheck_13-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_catcheck_13-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_catcheck_13` | 1.3.0 | `el8.x86_64` | pgdg | 39.5 KiB | [pg_catcheck_13-1.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_catcheck_13-1.3.0-1.rhel8.x86_64.rpm) |
| `pg_catcheck_13` | 1.6.0 | `el8.aarch64` | pgdg | 39.5 KiB | [pg_catcheck_13-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_catcheck_13-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_catcheck_13` | 1.3.0 | `el8.aarch64` | pgdg | 38.8 KiB | [pg_catcheck_13-1.3.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_catcheck_13-1.3.0-1.rhel8.aarch64.rpm) |
| `pg_catcheck_13` | 1.6.0 | `el9.x86_64` | pgdg | 37.3 KiB | [pg_catcheck_13-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_catcheck_13-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_catcheck_13` | 1.3.0 | `el9.x86_64` | pgdg | 37.4 KiB | [pg_catcheck_13-1.3.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_catcheck_13-1.3.0-1.rhel9.x86_64.rpm) |
| `pg_catcheck_13` | 1.6.0 | `el9.aarch64` | pgdg | 36.9 KiB | [pg_catcheck_13-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_catcheck_13-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_catcheck_13` | 1.3.0 | `el9.aarch64` | pgdg | 36.8 KiB | [pg_catcheck_13-1.3.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_catcheck_13-1.3.0-1.rhel9.aarch64.rpm) |
| `pg_catcheck_13` | 1.6.0 | `el10.x86_64` | pgdg | 38.2 KiB | [pg_catcheck_13-1.6.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_catcheck_13-1.6.0-3PGDG.rhel10.x86_64.rpm) |
| `pg_catcheck_13` | 1.6.0 | `el10.aarch64` | pgdg | 37.9 KiB | [pg_catcheck_13-1.6.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_catcheck_13-1.6.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pg-catcheck` | 1.6.0 | `d12.x86_64` | pgdg | 32.5 KiB | [postgresql-13-pg-catcheck_1.6.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-13-pg-catcheck_1.6.0-2.pgdg12+1_amd64.deb) |
| `postgresql-13-pg-catcheck` | 1.6.0 | `d12.aarch64` | pgdg | 31.5 KiB | [postgresql-13-pg-catcheck_1.6.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-13-pg-catcheck_1.6.0-2.pgdg12+1_arm64.deb) |
| `postgresql-13-pg-catcheck` | 1.6.0 | `d13.x86_64` | pgdg | 32.7 KiB | [postgresql-13-pg-catcheck_1.6.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-13-pg-catcheck_1.6.0-2.pgdg13+1_amd64.deb) |
| `postgresql-13-pg-catcheck` | 1.6.0 | `d13.aarch64` | pgdg | 31.8 KiB | [postgresql-13-pg-catcheck_1.6.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-13-pg-catcheck_1.6.0-2.pgdg13+1_arm64.deb) |
| `postgresql-13-pg-catcheck` | 1.6.0 | `u22.x86_64` | pgdg | 33.3 KiB | [postgresql-13-pg-catcheck_1.6.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-13-pg-catcheck_1.6.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pg-catcheck` | 1.6.0 | `u22.aarch64` | pgdg | 31.3 KiB | [postgresql-13-pg-catcheck_1.6.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-13-pg-catcheck_1.6.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pg-catcheck` | 1.6.0 | `u24.x86_64` | pgdg | 33.3 KiB | [postgresql-13-pg-catcheck_1.6.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-13-pg-catcheck_1.6.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pg-catcheck` | 1.6.0 | `u24.aarch64` | pgdg | 31.9 KiB | [postgresql-13-pg-catcheck_1.6.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-catcheck/postgresql-13-pg-catcheck_1.6.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/EnterpriseDB/pg_catcheck" title="Repository" icon="github" subtitle="github.com/EnterpriseDB/pg_catcheck" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_catcheck; # install by extension name, for the current active PG version
pig ext install pg_catcheck; # install via package alias, for the active PG version
pig ext install pg_catcheck -v 18;   # install for PG 18
pig ext install pg_catcheck -v 17;   # install for PG 17
pig ext install pg_catcheck -v 16;   # install for PG 16
pig ext install pg_catcheck -v 15;   # install for PG 15
pig ext install pg_catcheck -v 14;   # install for PG 14
pig ext install pg_catcheck -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_catcheck;
```

