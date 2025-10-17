---
title: "bgw_replstatus"
linkTitle: "bgw_replstatus"
description: "Small PostgreSQL background worker to report whether a node is a replication master or standby"
weight: 6340
categories: ["Stat"]
width: full
---

Small PostgreSQL background worker to report whether a node is a replication master or standby

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6340** | {{< badge content="bgw_replstatus" link="https://github.com/mhagander/bgw_replstatus" >}} | {{< ext "bgw_replstatus" "bgw_replstatus" >}} | `1.0.8` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sL---" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="No" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgautofailover" >}} {{< ext "pglogical" >}} {{< ext "pg_failover_slots" >}} {{< ext "pgpool_recovery" >}} {{< ext "pgsentinel" >}} {{< ext "pglogical_origin" >}} {{< ext "repmgr" >}} {{< ext "pg_jobmon" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/bgw_replstatus" >}} | `1.0.8` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `bgw_replstatus_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/bgw_replstatus" >}} | `1.0.8` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-bgw-replstatus` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "bgw_replstatus_18" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/bgw_replstatus_18-1.0.8-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "bgw_replstatus_17" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/bgw_replstatus_17-1.0.8-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "bgw_replstatus_16" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/bgw_replstatus_16-1.0.8-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "bgw_replstatus_15" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/bgw_replstatus_15-1.0.8-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "bgw_replstatus_14" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/bgw_replstatus_14-1.0.8-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "bgw_replstatus_18" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/bgw_replstatus_18-1.0.8-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "bgw_replstatus_17" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/bgw_replstatus_17-1.0.8-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "bgw_replstatus_16" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/bgw_replstatus_16-1.0.8-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "bgw_replstatus_15" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/bgw_replstatus_15-1.0.8-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "bgw_replstatus_14" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/bgw_replstatus_14-1.0.8-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "bgw_replstatus_18" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/bgw_replstatus_18-1.0.8-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "bgw_replstatus_17" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/bgw_replstatus_17-1.0.8-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "bgw_replstatus_16" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/bgw_replstatus_16-1.0.8-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "bgw_replstatus_15" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/bgw_replstatus_15-1.0.8-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "bgw_replstatus_14" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/bgw_replstatus_14-1.0.8-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "bgw_replstatus_18" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/bgw_replstatus_18-1.0.8-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "bgw_replstatus_17" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/bgw_replstatus_17-1.0.8-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "bgw_replstatus_16" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/bgw_replstatus_16-1.0.8-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "bgw_replstatus_15" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/bgw_replstatus_15-1.0.8-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "bgw_replstatus_14" "1.0.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/bgw_replstatus_14-1.0.8-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-bgw-replstatus" "1.0.8" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `bgw_replstatus_18` | 1.0.8 | `el8.x86_64` | pgdg | 16.0 KiB | [bgw_replstatus_18-1.0.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/bgw_replstatus_18-1.0.8-1PGDG.rhel8.x86_64.rpm) |
| `bgw_replstatus_18` | 1.0.8 | `el8.aarch64` | pgdg | 15.9 KiB | [bgw_replstatus_18-1.0.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/bgw_replstatus_18-1.0.8-1PGDG.rhel8.aarch64.rpm) |
| `bgw_replstatus_18` | 1.0.8 | `el9.aarch64` | pgdg | 14.8 KiB | [bgw_replstatus_18-1.0.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/bgw_replstatus_18-1.0.8-1PGDG.rhel9.aarch64.rpm) |
| `bgw_replstatus_18` | 1.0.8 | `el9.x86_64` | pgdg | 15.2 KiB | [bgw_replstatus_18-1.0.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/bgw_replstatus_18-1.0.8-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-18-bgw-replstatus` | 1.0.8 | `d12.aarch64` | pgdg | 14.9 KiB | [postgresql-18-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb) |
| `postgresql-18-bgw-replstatus` | 1.0.8 | `d12.x86_64` | pgdg | 14.7 KiB | [postgresql-18-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb) |
| `postgresql-18-bgw-replstatus` | 1.0.8 | `u22.aarch64` | pgdg | 14.6 KiB | [postgresql-18-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-bgw-replstatus` | 1.0.8 | `u22.x86_64` | pgdg | 15.0 KiB | [postgresql-18-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-bgw-replstatus` | 1.0.8 | `u24.aarch64` | pgdg | 14.9 KiB | [postgresql-18-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb) |
| `postgresql-18-bgw-replstatus` | 1.0.8 | `u24.x86_64` | pgdg | 14.8 KiB | [postgresql-18-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `bgw_replstatus_17` | 1.0.8 | `el8.x86_64` | pgdg | 16.0 KiB | [bgw_replstatus_17-1.0.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/bgw_replstatus_17-1.0.8-1PGDG.rhel8.x86_64.rpm) |
| `bgw_replstatus_17` | 1.0.8 | `el8.aarch64` | pgdg | 15.9 KiB | [bgw_replstatus_17-1.0.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/bgw_replstatus_17-1.0.8-1PGDG.rhel8.aarch64.rpm) |
| `bgw_replstatus_17` | 1.0.6 | `el8.aarch64` | pgdg | 15.6 KiB | [bgw_replstatus_17-1.0.6-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/bgw_replstatus_17-1.0.6-4PGDG.rhel8.aarch64.rpm) |
| `bgw_replstatus_17` | 1.0.6 | `el8.x86_64` | pgdg | 15.7 KiB | [bgw_replstatus_17-1.0.6-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/bgw_replstatus_17-1.0.6-4PGDG.rhel8.x86_64.rpm) |
| `bgw_replstatus_17` | 1.0.8 | `el9.aarch64` | pgdg | 15.0 KiB | [bgw_replstatus_17-1.0.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/bgw_replstatus_17-1.0.8-1PGDG.rhel9.aarch64.rpm) |
| `bgw_replstatus_17` | 1.0.8 | `el9.x86_64` | pgdg | 15.2 KiB | [bgw_replstatus_17-1.0.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/bgw_replstatus_17-1.0.8-1PGDG.rhel9.x86_64.rpm) |
| `bgw_replstatus_17` | 1.0.6 | `el9.x86_64` | pgdg | 13.5 KiB | [bgw_replstatus_17-1.0.6-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/bgw_replstatus_17-1.0.6-4PGDG.rhel9.x86_64.rpm) |
| `bgw_replstatus_17` | 1.0.6 | `el9.aarch64` | pgdg | 13.5 KiB | [bgw_replstatus_17-1.0.6-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/bgw_replstatus_17-1.0.6-4PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-bgw-replstatus` | 1.0.8 | `d12.aarch64` | pgdg | 14.8 KiB | [postgresql-17-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb) |
| `postgresql-17-bgw-replstatus` | 1.0.8 | `d12.x86_64` | pgdg | 14.7 KiB | [postgresql-17-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb) |
| `postgresql-17-bgw-replstatus` | 1.0.8 | `u22.x86_64` | pgdg | 15.4 KiB | [postgresql-17-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-bgw-replstatus` | 1.0.8 | `u22.aarch64` | pgdg | 15.0 KiB | [postgresql-17-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-bgw-replstatus` | 1.0.8 | `u24.aarch64` | pgdg | 14.8 KiB | [postgresql-17-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb) |
| `postgresql-17-bgw-replstatus` | 1.0.8 | `u24.x86_64` | pgdg | 14.7 KiB | [postgresql-17-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `bgw_replstatus_16` | 1.0.8 | `el8.x86_64` | pgdg | 16.0 KiB | [bgw_replstatus_16-1.0.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/bgw_replstatus_16-1.0.8-1PGDG.rhel8.x86_64.rpm) |
| `bgw_replstatus_16` | 1.0.8 | `el8.aarch64` | pgdg | 15.9 KiB | [bgw_replstatus_16-1.0.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/bgw_replstatus_16-1.0.8-1PGDG.rhel8.aarch64.rpm) |
| `bgw_replstatus_16` | 1.0.6 | `el8.aarch64` | pgdg | 15.5 KiB | [bgw_replstatus_16-1.0.6-2.rhel8.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/bgw_replstatus_16-1.0.6-2.rhel8.1.aarch64.rpm) |
| `bgw_replstatus_16` | 1.0.6 | `el8.x86_64` | pgdg | 15.5 KiB | [bgw_replstatus_16-1.0.6-2.rhel8.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/bgw_replstatus_16-1.0.6-2.rhel8.1.x86_64.rpm) |
| `bgw_replstatus_16` | 1.0.8 | `el9.aarch64` | pgdg | 14.9 KiB | [bgw_replstatus_16-1.0.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/bgw_replstatus_16-1.0.8-1PGDG.rhel9.aarch64.rpm) |
| `bgw_replstatus_16` | 1.0.8 | `el9.x86_64` | pgdg | 15.2 KiB | [bgw_replstatus_16-1.0.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/bgw_replstatus_16-1.0.8-1PGDG.rhel9.x86_64.rpm) |
| `bgw_replstatus_16` | 1.0.6 | `el9.aarch64` | pgdg | 13.2 KiB | [bgw_replstatus_16-1.0.6-2.rhel9.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/bgw_replstatus_16-1.0.6-2.rhel9.1.aarch64.rpm) |
| `bgw_replstatus_16` | 1.0.6 | `el9.x86_64` | pgdg | 13.3 KiB | [bgw_replstatus_16-1.0.6-2.rhel9.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/bgw_replstatus_16-1.0.6-2.rhel9.1.x86_64.rpm) |
| `bgw_replstatus_16` | 1.0.6 | `el9.x86_64` | pgdg | 13.5 KiB | [bgw_replstatus_16-1.0.6-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/bgw_replstatus_16-1.0.6-3PGDG.rhel9.x86_64.rpm) |
| `postgresql-16-bgw-replstatus` | 1.0.8 | `d12.aarch64` | pgdg | 14.8 KiB | [postgresql-16-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb) |
| `postgresql-16-bgw-replstatus` | 1.0.8 | `d12.x86_64` | pgdg | 14.7 KiB | [postgresql-16-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb) |
| `postgresql-16-bgw-replstatus` | 1.0.8 | `u22.x86_64` | pgdg | 15.4 KiB | [postgresql-16-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-bgw-replstatus` | 1.0.8 | `u22.aarch64` | pgdg | 15.0 KiB | [postgresql-16-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-bgw-replstatus` | 1.0.8 | `u24.x86_64` | pgdg | 14.7 KiB | [postgresql-16-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-bgw-replstatus` | 1.0.8 | `u24.aarch64` | pgdg | 14.8 KiB | [postgresql-16-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `bgw_replstatus_15` | 1.0.8 | `el8.aarch64` | pgdg | 15.9 KiB | [bgw_replstatus_15-1.0.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/bgw_replstatus_15-1.0.8-1PGDG.rhel8.aarch64.rpm) |
| `bgw_replstatus_15` | 1.0.8 | `el8.x86_64` | pgdg | 16.0 KiB | [bgw_replstatus_15-1.0.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/bgw_replstatus_15-1.0.8-1PGDG.rhel8.x86_64.rpm) |
| `bgw_replstatus_15` | 1.0.6 | `el8.x86_64` | pgdg | 15.4 KiB | [bgw_replstatus_15-1.0.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/bgw_replstatus_15-1.0.6-1.rhel8.x86_64.rpm) |
| `bgw_replstatus_15` | 1.0.6 | `el8.aarch64` | pgdg | 15.3 KiB | [bgw_replstatus_15-1.0.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/bgw_replstatus_15-1.0.6-1.rhel8.aarch64.rpm) |
| `bgw_replstatus_15` | 1.0.8 | `el9.x86_64` | pgdg | 15.2 KiB | [bgw_replstatus_15-1.0.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/bgw_replstatus_15-1.0.8-1PGDG.rhel9.x86_64.rpm) |
| `bgw_replstatus_15` | 1.0.8 | `el9.aarch64` | pgdg | 15.0 KiB | [bgw_replstatus_15-1.0.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/bgw_replstatus_15-1.0.8-1PGDG.rhel9.aarch64.rpm) |
| `bgw_replstatus_15` | 1.0.6 | `el9.x86_64` | pgdg | 15.2 KiB | [bgw_replstatus_15-1.0.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/bgw_replstatus_15-1.0.6-1.rhel9.x86_64.rpm) |
| `bgw_replstatus_15` | 1.0.6 | `el9.aarch64` | pgdg | 14.9 KiB | [bgw_replstatus_15-1.0.6-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/bgw_replstatus_15-1.0.6-1.rhel9.aarch64.rpm) |
| `postgresql-15-bgw-replstatus` | 1.0.8 | `d12.x86_64` | pgdg | 14.7 KiB | [postgresql-15-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb) |
| `postgresql-15-bgw-replstatus` | 1.0.8 | `d12.aarch64` | pgdg | 14.8 KiB | [postgresql-15-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb) |
| `postgresql-15-bgw-replstatus` | 1.0.8 | `u22.x86_64` | pgdg | 15.4 KiB | [postgresql-15-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-bgw-replstatus` | 1.0.8 | `u22.aarch64` | pgdg | 15.0 KiB | [postgresql-15-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-bgw-replstatus` | 1.0.8 | `u24.x86_64` | pgdg | 14.7 KiB | [postgresql-15-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-bgw-replstatus` | 1.0.8 | `u24.aarch64` | pgdg | 14.8 KiB | [postgresql-15-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `bgw_replstatus_14` | 1.0.8 | `el8.x86_64` | pgdg | 16.0 KiB | [bgw_replstatus_14-1.0.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/bgw_replstatus_14-1.0.8-1PGDG.rhel8.x86_64.rpm) |
| `bgw_replstatus_14` | 1.0.8 | `el8.aarch64` | pgdg | 15.9 KiB | [bgw_replstatus_14-1.0.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/bgw_replstatus_14-1.0.8-1PGDG.rhel8.aarch64.rpm) |
| `bgw_replstatus_14` | 1.0.6 | `el8.aarch64` | pgdg | 15.3 KiB | [bgw_replstatus_14-1.0.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/bgw_replstatus_14-1.0.6-1.rhel8.aarch64.rpm) |
| `bgw_replstatus_14` | 1.0.6 | `el8.x86_64` | pgdg | 15.4 KiB | [bgw_replstatus_14-1.0.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/bgw_replstatus_14-1.0.6-1.rhel8.x86_64.rpm) |
| `bgw_replstatus_14` | 1.0.3 | `el8.x86_64` | pgdg | 23.6 KiB | [bgw_replstatus_14-1.0.3-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/bgw_replstatus_14-1.0.3-3.rhel8.x86_64.rpm) |
| `bgw_replstatus_14` | 1.0.8 | `el9.aarch64` | pgdg | 15.0 KiB | [bgw_replstatus_14-1.0.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/bgw_replstatus_14-1.0.8-1PGDG.rhel9.aarch64.rpm) |
| `bgw_replstatus_14` | 1.0.8 | `el9.x86_64` | pgdg | 15.2 KiB | [bgw_replstatus_14-1.0.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/bgw_replstatus_14-1.0.8-1PGDG.rhel9.x86_64.rpm) |
| `bgw_replstatus_14` | 1.0.6 | `el9.aarch64` | pgdg | 14.9 KiB | [bgw_replstatus_14-1.0.6-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/bgw_replstatus_14-1.0.6-1.rhel9.aarch64.rpm) |
| `bgw_replstatus_14` | 1.0.6 | `el9.x86_64` | pgdg | 15.2 KiB | [bgw_replstatus_14-1.0.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/bgw_replstatus_14-1.0.6-1.rhel9.x86_64.rpm) |
| `postgresql-14-bgw-replstatus` | 1.0.8 | `d12.x86_64` | pgdg | 14.7 KiB | [postgresql-14-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb) |
| `postgresql-14-bgw-replstatus` | 1.0.8 | `d12.aarch64` | pgdg | 14.8 KiB | [postgresql-14-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb) |
| `postgresql-14-bgw-replstatus` | 1.0.8 | `u22.aarch64` | pgdg | 15.0 KiB | [postgresql-14-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-bgw-replstatus` | 1.0.8 | `u22.x86_64` | pgdg | 15.4 KiB | [postgresql-14-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-bgw-replstatus` | 1.0.8 | `u24.aarch64` | pgdg | 14.8 KiB | [postgresql-14-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb) |
| `postgresql-14-bgw-replstatus` | 1.0.8 | `u24.x86_64` | pgdg | 14.7 KiB | [postgresql-14-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `bgw_replstatus_13` | 1.0.8 | `el8.x86_64` | pgdg | 15.9 KiB | [bgw_replstatus_13-1.0.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/bgw_replstatus_13-1.0.8-1PGDG.rhel8.x86_64.rpm) |
| `bgw_replstatus_13` | 1.0.8 | `el8.aarch64` | pgdg | 15.9 KiB | [bgw_replstatus_13-1.0.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/bgw_replstatus_13-1.0.8-1PGDG.rhel8.aarch64.rpm) |
| `bgw_replstatus_13` | 1.0.6 | `el8.aarch64` | pgdg | 15.3 KiB | [bgw_replstatus_13-1.0.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/bgw_replstatus_13-1.0.6-1.rhel8.aarch64.rpm) |
| `bgw_replstatus_13` | 1.0.6 | `el8.x86_64` | pgdg | 15.3 KiB | [bgw_replstatus_13-1.0.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/bgw_replstatus_13-1.0.6-1.rhel8.x86_64.rpm) |
| `bgw_replstatus_13` | 1.0.8 | `el9.aarch64` | pgdg | 14.9 KiB | [bgw_replstatus_13-1.0.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/bgw_replstatus_13-1.0.8-1PGDG.rhel9.aarch64.rpm) |
| `bgw_replstatus_13` | 1.0.8 | `el9.x86_64` | pgdg | 15.2 KiB | [bgw_replstatus_13-1.0.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/bgw_replstatus_13-1.0.8-1PGDG.rhel9.x86_64.rpm) |
| `bgw_replstatus_13` | 1.0.6 | `el9.x86_64` | pgdg | 15.2 KiB | [bgw_replstatus_13-1.0.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/bgw_replstatus_13-1.0.6-1.rhel9.x86_64.rpm) |
| `bgw_replstatus_13` | 1.0.6 | `el9.aarch64` | pgdg | 14.9 KiB | [bgw_replstatus_13-1.0.6-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/bgw_replstatus_13-1.0.6-1.rhel9.aarch64.rpm) |
| `postgresql-13-bgw-replstatus` | 1.0.8 | `d12.x86_64` | pgdg | 14.6 KiB | [postgresql-13-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-13-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb) |
| `postgresql-13-bgw-replstatus` | 1.0.8 | `d12.aarch64` | pgdg | 15.0 KiB | [postgresql-13-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-13-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb) |
| `postgresql-13-bgw-replstatus` | 1.0.8 | `u22.aarch64` | pgdg | 14.9 KiB | [postgresql-13-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-13-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-bgw-replstatus` | 1.0.8 | `u22.x86_64` | pgdg | 15.2 KiB | [postgresql-13-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-13-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-bgw-replstatus` | 1.0.8 | `u24.x86_64` | pgdg | 14.7 KiB | [postgresql-13-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-13-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-bgw-replstatus` | 1.0.8 | `u24.aarch64` | pgdg | 15.1 KiB | [postgresql-13-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-13-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/mhagander/bgw_replstatus" title="Repository" icon="github" subtitle="github.com/mhagander/bgw_replstatus" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install bgw_replstatus; # install by extension name, for the current active PG version
pig ext install bgw_replstatus; # install via package alias, for the active PG version
pig ext install bgw_replstatus -v 18;   # install for PG 18
pig ext install bgw_replstatus -v 17;   # install for PG 17
pig ext install bgw_replstatus -v 16;   # install for PG 16
pig ext install bgw_replstatus -v 15;   # install for PG 15
pig ext install bgw_replstatus -v 14;   # install for PG 14
pig ext install bgw_replstatus -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION bgw_replstatus;
```

