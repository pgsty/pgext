---
title: "pg_dirtyread"
linkTitle: "pg_dirtyread"
description: "Read dead but unvacuumed rows from table"
weight: 5050
categories: ["Admin"]
width: full
---

Read dead but unvacuumed rows from table

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5050** | {{< badge content="pg_dirtyread" link="https://github.com/df7cb/pg_dirtyread" >}} | {{< ext "pg_dirtyread" "pg_dirtyread" >}} | `2.7` | {{< category "ADMIN" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_orphaned" >}} {{< ext "pg_surgery" >}} {{< ext "pageinspect" >}} {{< ext "pg_visibility" >}} {{< ext "pg_cheat_funcs" >}} {{< ext "amcheck" >}} {{< ext "pg_repack" >}} {{< ext "pg_squeeze" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_dirtyread" >}} | `2.7` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_dirtyread_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pg_dirtyread" >}} | `2.7` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-dirtyread` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_dirtyread_18" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_dirtyread_18-2.7-4PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_dirtyread_17" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_dirtyread_17-2.7-2PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_dirtyread_16" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_dirtyread_16-2.7-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_dirtyread_15" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_dirtyread_15-2.7-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_dirtyread_14" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_dirtyread_14-2.7-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_dirtyread_18" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_dirtyread_18-2.7-4PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_dirtyread_17" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_dirtyread_17-2.7-2PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_dirtyread_16" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_dirtyread_16-2.7-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_dirtyread_15" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_dirtyread_15-2.7-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_dirtyread_14" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_dirtyread_14-2.7-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_dirtyread_18" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_dirtyread_18-2.7-4PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_dirtyread_17" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_dirtyread_17-2.7-2PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_dirtyread_16" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_dirtyread_16-2.7-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_dirtyread_15" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_dirtyread_15-2.7-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_dirtyread_14" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dirtyread_14-2.7-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_dirtyread_18" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_dirtyread_18-2.7-4PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_dirtyread_17" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_dirtyread_17-2.7-2PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_dirtyread_16" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_dirtyread_16-2.7-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_dirtyread_15" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_dirtyread_15-2.7-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_dirtyread_14" "2.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_dirtyread_14-2.7-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-dirtyread" "2.7" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_dirtyread_18` | 2.7 | `el8.aarch64` | pgdg | 16.9 KiB | [pg_dirtyread_18-2.7-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_dirtyread_18-2.7-4PGDG.rhel8.aarch64.rpm) |
| `pg_dirtyread_18` | 2.7 | `el8.x86_64` | pgdg | 17.0 KiB | [pg_dirtyread_18-2.7-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_dirtyread_18-2.7-4PGDG.rhel8.x86_64.rpm) |
| `pg_dirtyread_18` | 2.7 | `el9.aarch64` | pgdg | 16.6 KiB | [pg_dirtyread_18-2.7-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_dirtyread_18-2.7-4PGDG.rhel9.aarch64.rpm) |
| `pg_dirtyread_18` | 2.7 | `el9.x86_64` | pgdg | 17.2 KiB | [pg_dirtyread_18-2.7-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_dirtyread_18-2.7-4PGDG.rhel9.x86_64.rpm) |
| `postgresql-18-dirtyread` | 2.7 | `d12.aarch64` | pgdg | 20.9 KiB | [postgresql-18-dirtyread_2.7-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg12+1_arm64.deb) |
| `postgresql-18-dirtyread` | 2.7 | `d12.x86_64` | pgdg | 21.1 KiB | [postgresql-18-dirtyread_2.7-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg12+1_amd64.deb) |
| `postgresql-18-dirtyread` | 2.7 | `u22.x86_64` | pgdg | 22.0 KiB | [postgresql-18-dirtyread_2.7-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-dirtyread` | 2.7 | `u22.aarch64` | pgdg | 21.4 KiB | [postgresql-18-dirtyread_2.7-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-dirtyread` | 2.7 | `u24.x86_64` | pgdg | 21.2 KiB | [postgresql-18-dirtyread_2.7-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-dirtyread` | 2.7 | `u24.aarch64` | pgdg | 20.9 KiB | [postgresql-18-dirtyread_2.7-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_dirtyread_17` | 2.7 | `el8.x86_64` | pgdg | 16.8 KiB | [pg_dirtyread_17-2.7-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_dirtyread_17-2.7-2PGDG.rhel8.x86_64.rpm) |
| `pg_dirtyread_17` | 2.7 | `el8.aarch64` | pgdg | 16.7 KiB | [pg_dirtyread_17-2.7-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_dirtyread_17-2.7-2PGDG.rhel8.aarch64.rpm) |
| `pg_dirtyread_17` | 2.7 | `el9.aarch64` | pgdg | 16.6 KiB | [pg_dirtyread_17-2.7-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_dirtyread_17-2.7-2PGDG.rhel9.aarch64.rpm) |
| `pg_dirtyread_17` | 2.7 | `el9.x86_64` | pgdg | 17.0 KiB | [pg_dirtyread_17-2.7-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_dirtyread_17-2.7-2PGDG.rhel9.x86_64.rpm) |
| `postgresql-17-dirtyread` | 2.7 | `d12.x86_64` | pgdg | 20.9 KiB | [postgresql-17-dirtyread_2.7-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg12+1_amd64.deb) |
| `postgresql-17-dirtyread` | 2.7 | `d12.aarch64` | pgdg | 20.8 KiB | [postgresql-17-dirtyread_2.7-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg12+1_arm64.deb) |
| `postgresql-17-dirtyread` | 2.7 | `u22.aarch64` | pgdg | 25.3 KiB | [postgresql-17-dirtyread_2.7-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-dirtyread` | 2.7 | `u22.x86_64` | pgdg | 26.0 KiB | [postgresql-17-dirtyread_2.7-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-dirtyread` | 2.7 | `u24.aarch64` | pgdg | 20.8 KiB | [postgresql-17-dirtyread_2.7-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg24.04+1_arm64.deb) |
| `postgresql-17-dirtyread` | 2.7 | `u24.x86_64` | pgdg | 21.1 KiB | [postgresql-17-dirtyread_2.7-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_dirtyread_16` | 2.7 | `el8.x86_64` | pgdg | 16.7 KiB | [pg_dirtyread_16-2.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_dirtyread_16-2.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_dirtyread_16` | 2.7 | `el8.aarch64` | pgdg | 16.6 KiB | [pg_dirtyread_16-2.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_dirtyread_16-2.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_dirtyread_16` | 2.7 | `el9.x86_64` | pgdg | 16.9 KiB | [pg_dirtyread_16-2.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_dirtyread_16-2.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_dirtyread_16` | 2.7 | `el9.aarch64` | pgdg | 16.5 KiB | [pg_dirtyread_16-2.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_dirtyread_16-2.7-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-dirtyread` | 2.7 | `d12.aarch64` | pgdg | 20.8 KiB | [postgresql-16-dirtyread_2.7-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg12+1_arm64.deb) |
| `postgresql-16-dirtyread` | 2.7 | `d12.x86_64` | pgdg | 20.9 KiB | [postgresql-16-dirtyread_2.7-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg12+1_amd64.deb) |
| `postgresql-16-dirtyread` | 2.7 | `u22.aarch64` | pgdg | 24.9 KiB | [postgresql-16-dirtyread_2.7-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-dirtyread` | 2.7 | `u22.x86_64` | pgdg | 25.5 KiB | [postgresql-16-dirtyread_2.7-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-dirtyread` | 2.7 | `u24.x86_64` | pgdg | 21.1 KiB | [postgresql-16-dirtyread_2.7-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-dirtyread` | 2.7 | `u24.aarch64` | pgdg | 20.8 KiB | [postgresql-16-dirtyread_2.7-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_dirtyread_15` | 2.7 | `el8.aarch64` | pgdg | 16.6 KiB | [pg_dirtyread_15-2.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_dirtyread_15-2.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_dirtyread_15` | 2.7 | `el8.x86_64` | pgdg | 16.8 KiB | [pg_dirtyread_15-2.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_dirtyread_15-2.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_dirtyread_15` | 2.7 | `el9.x86_64` | pgdg | 17.0 KiB | [pg_dirtyread_15-2.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_dirtyread_15-2.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_dirtyread_15` | 2.7 | `el9.aarch64` | pgdg | 16.8 KiB | [pg_dirtyread_15-2.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_dirtyread_15-2.7-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-15-dirtyread` | 2.7 | `d12.x86_64` | pgdg | 21.1 KiB | [postgresql-15-dirtyread_2.7-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg12+1_amd64.deb) |
| `postgresql-15-dirtyread` | 2.7 | `d12.aarch64` | pgdg | 20.8 KiB | [postgresql-15-dirtyread_2.7-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg12+1_arm64.deb) |
| `postgresql-15-dirtyread` | 2.7 | `u22.aarch64` | pgdg | 25.0 KiB | [postgresql-15-dirtyread_2.7-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-dirtyread` | 2.7 | `u22.x86_64` | pgdg | 25.6 KiB | [postgresql-15-dirtyread_2.7-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-dirtyread` | 2.7 | `u24.x86_64` | pgdg | 21.1 KiB | [postgresql-15-dirtyread_2.7-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-dirtyread` | 2.7 | `u24.aarch64` | pgdg | 20.9 KiB | [postgresql-15-dirtyread_2.7-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_dirtyread_14` | 2.7 | `el8.x86_64` | pgdg | 16.8 KiB | [pg_dirtyread_14-2.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_dirtyread_14-2.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_dirtyread_14` | 2.7 | `el8.aarch64` | pgdg | 16.6 KiB | [pg_dirtyread_14-2.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_dirtyread_14-2.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_dirtyread_14` | 2.7 | `el9.x86_64` | pgdg | 17.0 KiB | [pg_dirtyread_14-2.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dirtyread_14-2.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_dirtyread_14` | 2.7 | `el9.aarch64` | pgdg | 16.8 KiB | [pg_dirtyread_14-2.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_dirtyread_14-2.7-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-14-dirtyread` | 2.7 | `d12.x86_64` | pgdg | 21.0 KiB | [postgresql-14-dirtyread_2.7-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg12+1_amd64.deb) |
| `postgresql-14-dirtyread` | 2.7 | `d12.aarch64` | pgdg | 20.8 KiB | [postgresql-14-dirtyread_2.7-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg12+1_arm64.deb) |
| `postgresql-14-dirtyread` | 2.7 | `u22.x86_64` | pgdg | 25.5 KiB | [postgresql-14-dirtyread_2.7-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-dirtyread` | 2.7 | `u22.aarch64` | pgdg | 25.0 KiB | [postgresql-14-dirtyread_2.7-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-dirtyread` | 2.7 | `u24.aarch64` | pgdg | 20.8 KiB | [postgresql-14-dirtyread_2.7-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg24.04+1_arm64.deb) |
| `postgresql-14-dirtyread` | 2.7 | `u24.x86_64` | pgdg | 21.1 KiB | [postgresql-14-dirtyread_2.7-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_dirtyread_13` | 2.7 | `el8.aarch64` | pgdg | 16.6 KiB | [pg_dirtyread_13-2.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_dirtyread_13-2.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_dirtyread_13` | 2.7 | `el8.x86_64` | pgdg | 16.7 KiB | [pg_dirtyread_13-2.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_dirtyread_13-2.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_dirtyread_13` | 2.7 | `el9.aarch64` | pgdg | 16.7 KiB | [pg_dirtyread_13-2.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_dirtyread_13-2.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_dirtyread_13` | 2.7 | `el9.x86_64` | pgdg | 17.0 KiB | [pg_dirtyread_13-2.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_dirtyread_13-2.7-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-13-dirtyread` | 2.7 | `d12.aarch64` | pgdg | 20.8 KiB | [postgresql-13-dirtyread_2.7-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-13-dirtyread_2.7-3.pgdg12+1_arm64.deb) |
| `postgresql-13-dirtyread` | 2.7 | `d12.x86_64` | pgdg | 20.9 KiB | [postgresql-13-dirtyread_2.7-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-13-dirtyread_2.7-3.pgdg12+1_amd64.deb) |
| `postgresql-13-dirtyread` | 2.7 | `u22.aarch64` | pgdg | 24.9 KiB | [postgresql-13-dirtyread_2.7-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-13-dirtyread_2.7-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-dirtyread` | 2.7 | `u22.x86_64` | pgdg | 25.2 KiB | [postgresql-13-dirtyread_2.7-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-13-dirtyread_2.7-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-dirtyread` | 2.7 | `u24.x86_64` | pgdg | 20.9 KiB | [postgresql-13-dirtyread_2.7-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-13-dirtyread_2.7-3.pgdg24.04+1_amd64.deb) |
| `postgresql-13-dirtyread` | 2.7 | `u24.aarch64` | pgdg | 20.6 KiB | [postgresql-13-dirtyread_2.7-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-13-dirtyread_2.7-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/df7cb/pg_dirtyread" title="Repository" icon="github" subtitle="github.com/df7cb/pg_dirtyread" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_dirtyread-2.7.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_dirtyread; # get pg_dirtyread source code
pig build dep pg_dirtyread; # install build dependencies
pig build pkg pg_dirtyread; # build extension rpm or deb
pig build ext pg_dirtyread; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_dirtyread; # install by extension name, for the current active PG version
pig ext install pg_dirtyread; # install via package alias, for the active PG version
pig ext install pg_dirtyread -v 18;   # install for PG 18
pig ext install pg_dirtyread -v 17;   # install for PG 17
pig ext install pg_dirtyread -v 16;   # install for PG 16
pig ext install pg_dirtyread -v 15;   # install for PG 15
pig ext install pg_dirtyread -v 14;   # install for PG 14
pig ext install pg_dirtyread -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_dirtyread;
```

