---
title: "pg_failover_slots"
linkTitle: "pg_failover_slots"
description: "PG Failover Slots extension"
weight: 9530
categories: ["Etl"]
width: full
---

PG Failover Slots extension

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9530** | {{< badge content="pg_failover_slots" link="https://github.com/EnterpriseDB/pg_failover_slots" >}} | {{< ext "pg_failover_slots" "pg_failover_slots" >}} | `1.1.0` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sL--r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="No" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pglogical" >}} {{< ext "pglogical_origin" >}} {{< ext "pglogical_ticker" >}} {{< ext "pgactive" >}} {{< ext "repmgr" >}} {{< ext "bgw_replstatus" >}} {{< ext "pgl_ddl_deploy" >}} {{< ext "decoderbufs" >}} |

> [!Note] break on pg18


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_failover_slots" >}} | `1.1.0` | {{< badge content="18" color="red" alt="pg_failover_slots_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_failover_slots_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_failover_slots" >}} | `1.1.0` | {{< badge content="18" color="red" alt="postgresql-18-pg-failover-slots" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-failover-slots` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_failover_slots_18" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_failover_slots_18-1.2.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_failover_slots_17" "1.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_failover_slots_17-1.1.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_failover_slots_16" "1.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_failover_slots_16-1.1.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_failover_slots_15" "1.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_failover_slots_15-1.1.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_failover_slots_14" "1.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_failover_slots_14-1.1.0-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_failover_slots_18" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_failover_slots_18-1.2.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_failover_slots_17" "1.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_failover_slots_17-1.1.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_failover_slots_16" "1.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_failover_slots_16-1.1.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_failover_slots_15" "1.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_failover_slots_15-1.1.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_failover_slots_14" "1.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_failover_slots_14-1.1.0-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_failover_slots_18" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_failover_slots_18-1.2.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_failover_slots_17" "1.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_failover_slots_17-1.1.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_failover_slots_16" "1.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_failover_slots_16-1.1.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_failover_slots_15" "1.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_failover_slots_15-1.1.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_failover_slots_14" "1.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_failover_slots_14-1.1.0-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_failover_slots_18" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_failover_slots_18-1.2.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_failover_slots_17" "1.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_failover_slots_17-1.1.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_failover_slots_16" "1.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_failover_slots_16-1.1.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_failover_slots_15" "1.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_failover_slots_15-1.1.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_failover_slots_14" "1.1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_failover_slots_14-1.1.0-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-18-pg-failover-slots_1.2.0-1.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-17-pg-failover-slots_1.2.0-1.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-16-pg-failover-slots_1.2.0-1.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-15-pg-failover-slots_1.2.0-1.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-14-pg-failover-slots_1.2.0-1.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-18-pg-failover-slots_1.2.0-1.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-17-pg-failover-slots_1.2.0-1.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-16-pg-failover-slots_1.2.0-1.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-15-pg-failover-slots_1.2.0-1.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-14-pg-failover-slots_1.2.0-1.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-18-pg-failover-slots_1.2.0-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-17-pg-failover-slots_1.2.0-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-16-pg-failover-slots_1.2.0-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-15-pg-failover-slots_1.2.0-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-14-pg-failover-slots_1.2.0-1.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-18-pg-failover-slots_1.2.0-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-17-pg-failover-slots_1.2.0-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-16-pg-failover-slots_1.2.0-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-15-pg-failover-slots_1.2.0-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-14-pg-failover-slots_1.2.0-1.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-18-pg-failover-slots_1.2.0-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-17-pg-failover-slots_1.2.0-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-16-pg-failover-slots_1.2.0-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-15-pg-failover-slots_1.2.0-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-14-pg-failover-slots_1.2.0-1.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-18-pg-failover-slots_1.2.0-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-17-pg-failover-slots_1.2.0-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-16-pg-failover-slots_1.2.0-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-15-pg-failover-slots_1.2.0-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pg-failover-slots" "1.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-14-pg-failover-slots_1.2.0-1.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_failover_slots_18` | 1.2.0 | `el8.aarch64` | pgdg | 24.2 KiB | [pg_failover_slots_18-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_failover_slots_18-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_failover_slots_18` | 1.2.0 | `el8.x86_64` | pgdg | 24.7 KiB | [pg_failover_slots_18-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_failover_slots_18-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_failover_slots_18` | 1.2.0 | `el9.aarch64` | pgdg | 24.3 KiB | [pg_failover_slots_18-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_failover_slots_18-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_failover_slots_18` | 1.2.0 | `el9.x86_64` | pgdg | 24.7 KiB | [pg_failover_slots_18-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_failover_slots_18-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-18-pg-failover-slots` | 1.2.0 | `d12.x86_64` | pgdg | 39.3 KiB | [postgresql-18-pg-failover-slots_1.2.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-18-pg-failover-slots_1.2.0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-failover-slots` | 1.2.0 | `d12.aarch64` | pgdg | 38.7 KiB | [postgresql-18-pg-failover-slots_1.2.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-18-pg-failover-slots_1.2.0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-failover-slots` | 1.2.0 | `u22.aarch64` | pgdg | 38.8 KiB | [postgresql-18-pg-failover-slots_1.2.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-18-pg-failover-slots_1.2.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-failover-slots` | 1.2.0 | `u22.x86_64` | pgdg | 39.9 KiB | [postgresql-18-pg-failover-slots_1.2.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-18-pg-failover-slots_1.2.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-failover-slots` | 1.2.0 | `u24.aarch64` | pgdg | 38.8 KiB | [postgresql-18-pg-failover-slots_1.2.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-18-pg-failover-slots_1.2.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pg-failover-slots` | 1.2.0 | `u24.x86_64` | pgdg | 39.6 KiB | [postgresql-18-pg-failover-slots_1.2.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-18-pg-failover-slots_1.2.0-1.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_failover_slots_17` | 1.1.0 | `el8.aarch64` | pgdg | 23.8 KiB | [pg_failover_slots_17-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_failover_slots_17-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_failover_slots_17` | 1.1.0 | `el8.aarch64` | pigsty | 23.1 KiB | [pg_failover_slots_17-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_failover_slots_17-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_failover_slots_17` | 1.1.0 | `el8.x86_64` | pigsty | 23.5 KiB | [pg_failover_slots_17-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_failover_slots_17-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_failover_slots_17` | 1.1.0 | `el8.x86_64` | pgdg | 24.2 KiB | [pg_failover_slots_17-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_failover_slots_17-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_failover_slots_17` | 1.1.0 | `el9.aarch64` | pgdg | 24.4 KiB | [pg_failover_slots_17-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_failover_slots_17-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_failover_slots_17` | 1.1.0 | `el9.aarch64` | pigsty | 23.6 KiB | [pg_failover_slots_17-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_failover_slots_17-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_failover_slots_17` | 1.1.0 | `el9.x86_64` | pigsty | 23.7 KiB | [pg_failover_slots_17-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_failover_slots_17-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_failover_slots_17` | 1.1.0 | `el9.x86_64` | pgdg | 24.5 KiB | [pg_failover_slots_17-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_failover_slots_17-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-17-pg-failover-slots` | 1.2.0 | `d12.aarch64` | pgdg | 38.7 KiB | [postgresql-17-pg-failover-slots_1.2.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-17-pg-failover-slots_1.2.0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-failover-slots` | 1.2.0 | `d12.x86_64` | pgdg | 39.4 KiB | [postgresql-17-pg-failover-slots_1.2.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-17-pg-failover-slots_1.2.0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-failover-slots` | 1.1.0 | `d12.x86_64` | pigsty | 43.3 KiB | [postgresql-17-pg-failover-slots_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-failover-slots/postgresql-17-pg-failover-slots_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-failover-slots` | 1.1.0 | `d12.aarch64` | pigsty | 42.1 KiB | [postgresql-17-pg-failover-slots_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-failover-slots/postgresql-17-pg-failover-slots_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-failover-slots` | 1.2.0 | `u22.x86_64` | pgdg | 44.2 KiB | [postgresql-17-pg-failover-slots_1.2.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-17-pg-failover-slots_1.2.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-failover-slots` | 1.2.0 | `u22.aarch64` | pgdg | 43.3 KiB | [postgresql-17-pg-failover-slots_1.2.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-17-pg-failover-slots_1.2.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-failover-slots` | 1.1.0 | `u22.aarch64` | pigsty | 44.7 KiB | [postgresql-17-pg-failover-slots_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-failover-slots/postgresql-17-pg-failover-slots_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-failover-slots` | 1.1.0 | `u22.x86_64` | pigsty | 45.4 KiB | [postgresql-17-pg-failover-slots_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-failover-slots/postgresql-17-pg-failover-slots_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-failover-slots` | 1.2.0 | `u24.aarch64` | pgdg | 38.7 KiB | [postgresql-17-pg-failover-slots_1.2.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-17-pg-failover-slots_1.2.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pg-failover-slots` | 1.2.0 | `u24.x86_64` | pgdg | 39.6 KiB | [postgresql-17-pg-failover-slots_1.2.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-17-pg-failover-slots_1.2.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-failover-slots` | 1.1.0 | `u24.aarch64` | pigsty | 39.9 KiB | [postgresql-17-pg-failover-slots_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-failover-slots/postgresql-17-pg-failover-slots_1.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-failover-slots` | 1.1.0 | `u24.x86_64` | pigsty | 40.5 KiB | [postgresql-17-pg-failover-slots_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-failover-slots/postgresql-17-pg-failover-slots_1.1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_failover_slots_16` | 1.1.0 | `el8.x86_64` | pgdg | 24.2 KiB | [pg_failover_slots_16-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_failover_slots_16-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_failover_slots_16` | 1.1.0 | `el8.x86_64` | pigsty | 23.5 KiB | [pg_failover_slots_16-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_failover_slots_16-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_failover_slots_16` | 1.1.0 | `el8.aarch64` | pigsty | 23.1 KiB | [pg_failover_slots_16-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_failover_slots_16-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_failover_slots_16` | 1.1.0 | `el8.aarch64` | pgdg | 23.8 KiB | [pg_failover_slots_16-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_failover_slots_16-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_failover_slots_16` | 1.0.1 | `el8.aarch64` | pgdg | 23.1 KiB | [pg_failover_slots_16-1.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_failover_slots_16-1.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_failover_slots_16` | 1.0.1 | `el8.x86_64` | pgdg | 23.5 KiB | [pg_failover_slots_16-1.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_failover_slots_16-1.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_failover_slots_16` | 1.1.0 | `el9.x86_64` | pgdg | 24.5 KiB | [pg_failover_slots_16-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_failover_slots_16-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_failover_slots_16` | 1.1.0 | `el9.aarch64` | pigsty | 23.6 KiB | [pg_failover_slots_16-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_failover_slots_16-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_failover_slots_16` | 1.1.0 | `el9.aarch64` | pgdg | 24.4 KiB | [pg_failover_slots_16-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_failover_slots_16-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_failover_slots_16` | 1.1.0 | `el9.x86_64` | pigsty | 23.7 KiB | [pg_failover_slots_16-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_failover_slots_16-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_failover_slots_16` | 1.0.1 | `el9.x86_64` | pgdg | 23.6 KiB | [pg_failover_slots_16-1.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_failover_slots_16-1.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_failover_slots_16` | 1.0.1 | `el9.aarch64` | pgdg | 23.4 KiB | [pg_failover_slots_16-1.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_failover_slots_16-1.0.1-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-pg-failover-slots` | 1.2.0 | `d12.aarch64` | pgdg | 38.7 KiB | [postgresql-16-pg-failover-slots_1.2.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-16-pg-failover-slots_1.2.0-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-failover-slots` | 1.2.0 | `d12.x86_64` | pgdg | 39.3 KiB | [postgresql-16-pg-failover-slots_1.2.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-16-pg-failover-slots_1.2.0-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-failover-slots` | 1.1.0 | `d12.x86_64` | pigsty | 41.9 KiB | [postgresql-16-pg-failover-slots_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-failover-slots/postgresql-16-pg-failover-slots_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-failover-slots` | 1.1.0 | `d12.aarch64` | pigsty | 40.7 KiB | [postgresql-16-pg-failover-slots_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-failover-slots/postgresql-16-pg-failover-slots_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-failover-slots` | 1.2.0 | `u22.x86_64` | pgdg | 42.7 KiB | [postgresql-16-pg-failover-slots_1.2.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-16-pg-failover-slots_1.2.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-failover-slots` | 1.2.0 | `u22.aarch64` | pgdg | 41.9 KiB | [postgresql-16-pg-failover-slots_1.2.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-16-pg-failover-slots_1.2.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-failover-slots` | 1.1.0 | `u22.aarch64` | pigsty | 43.2 KiB | [postgresql-16-pg-failover-slots_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-failover-slots/postgresql-16-pg-failover-slots_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-failover-slots` | 1.1.0 | `u22.x86_64` | pigsty | 43.9 KiB | [postgresql-16-pg-failover-slots_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-failover-slots/postgresql-16-pg-failover-slots_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-failover-slots` | 1.2.0 | `u24.aarch64` | pgdg | 38.7 KiB | [postgresql-16-pg-failover-slots_1.2.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-16-pg-failover-slots_1.2.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pg-failover-slots` | 1.2.0 | `u24.x86_64` | pgdg | 39.4 KiB | [postgresql-16-pg-failover-slots_1.2.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-16-pg-failover-slots_1.2.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-failover-slots` | 1.1.0 | `u24.aarch64` | pigsty | 39.9 KiB | [postgresql-16-pg-failover-slots_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-failover-slots/postgresql-16-pg-failover-slots_1.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-failover-slots` | 1.1.0 | `u24.x86_64` | pigsty | 40.5 KiB | [postgresql-16-pg-failover-slots_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-failover-slots/postgresql-16-pg-failover-slots_1.1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_failover_slots_15` | 1.1.0 | `el8.x86_64` | pgdg | 24.5 KiB | [pg_failover_slots_15-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_failover_slots_15-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_failover_slots_15` | 1.1.0 | `el8.aarch64` | pigsty | 23.3 KiB | [pg_failover_slots_15-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_failover_slots_15-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_failover_slots_15` | 1.1.0 | `el8.x86_64` | pigsty | 23.7 KiB | [pg_failover_slots_15-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_failover_slots_15-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_failover_slots_15` | 1.1.0 | `el8.aarch64` | pgdg | 24.0 KiB | [pg_failover_slots_15-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_failover_slots_15-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_failover_slots_15` | 1.0.1 | `el8.x86_64` | pgdg | 23.7 KiB | [pg_failover_slots_15-1.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_failover_slots_15-1.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_failover_slots_15` | 1.0.1 | `el8.aarch64` | pgdg | 23.2 KiB | [pg_failover_slots_15-1.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_failover_slots_15-1.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_failover_slots_15` | 1.0.0 | `el8.x86_64` | pgdg | 23.0 KiB | [pg_failover_slots_15-1.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_failover_slots_15-1.0.0-1.rhel8.x86_64.rpm) |
| `pg_failover_slots_15` | 1.0.0 | `el8.aarch64` | pgdg | 22.4 KiB | [pg_failover_slots_15-1.0.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_failover_slots_15-1.0.0-1.rhel8.aarch64.rpm) |
| `pg_failover_slots_15` | 1.1.0 | `el9.x86_64` | pigsty | 23.9 KiB | [pg_failover_slots_15-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_failover_slots_15-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_failover_slots_15` | 1.1.0 | `el9.x86_64` | pgdg | 24.6 KiB | [pg_failover_slots_15-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_failover_slots_15-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_failover_slots_15` | 1.1.0 | `el9.aarch64` | pigsty | 23.8 KiB | [pg_failover_slots_15-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_failover_slots_15-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_failover_slots_15` | 1.1.0 | `el9.aarch64` | pgdg | 24.6 KiB | [pg_failover_slots_15-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_failover_slots_15-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_failover_slots_15` | 1.0.1 | `el9.aarch64` | pgdg | 23.5 KiB | [pg_failover_slots_15-1.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_failover_slots_15-1.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_failover_slots_15` | 1.0.1 | `el9.x86_64` | pgdg | 23.7 KiB | [pg_failover_slots_15-1.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_failover_slots_15-1.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_failover_slots_15` | 1.0.0 | `el9.x86_64` | pgdg | 23.1 KiB | [pg_failover_slots_15-1.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_failover_slots_15-1.0.0-1.rhel9.x86_64.rpm) |
| `pg_failover_slots_15` | 1.0.0 | `el9.aarch64` | pgdg | 22.8 KiB | [pg_failover_slots_15-1.0.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_failover_slots_15-1.0.0-1.rhel9.aarch64.rpm) |
| `postgresql-15-pg-failover-slots` | 1.2.0 | `d12.x86_64` | pgdg | 39.5 KiB | [postgresql-15-pg-failover-slots_1.2.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-15-pg-failover-slots_1.2.0-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-failover-slots` | 1.2.0 | `d12.aarch64` | pgdg | 38.7 KiB | [postgresql-15-pg-failover-slots_1.2.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-15-pg-failover-slots_1.2.0-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-failover-slots` | 1.1.0 | `d12.x86_64` | pigsty | 42.0 KiB | [postgresql-15-pg-failover-slots_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-failover-slots/postgresql-15-pg-failover-slots_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-failover-slots` | 1.1.0 | `d12.aarch64` | pigsty | 40.9 KiB | [postgresql-15-pg-failover-slots_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-failover-slots/postgresql-15-pg-failover-slots_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-failover-slots` | 1.2.0 | `u22.x86_64` | pgdg | 42.7 KiB | [postgresql-15-pg-failover-slots_1.2.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-15-pg-failover-slots_1.2.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-failover-slots` | 1.2.0 | `u22.aarch64` | pgdg | 41.8 KiB | [postgresql-15-pg-failover-slots_1.2.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-15-pg-failover-slots_1.2.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-failover-slots` | 1.1.0 | `u22.x86_64` | pigsty | 44.1 KiB | [postgresql-15-pg-failover-slots_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-failover-slots/postgresql-15-pg-failover-slots_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-failover-slots` | 1.1.0 | `u22.aarch64` | pigsty | 43.4 KiB | [postgresql-15-pg-failover-slots_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-failover-slots/postgresql-15-pg-failover-slots_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-failover-slots` | 1.2.0 | `u24.aarch64` | pgdg | 38.7 KiB | [postgresql-15-pg-failover-slots_1.2.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-15-pg-failover-slots_1.2.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pg-failover-slots` | 1.2.0 | `u24.x86_64` | pgdg | 39.6 KiB | [postgresql-15-pg-failover-slots_1.2.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-15-pg-failover-slots_1.2.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-failover-slots` | 1.1.0 | `u24.aarch64` | pigsty | 40.2 KiB | [postgresql-15-pg-failover-slots_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-failover-slots/postgresql-15-pg-failover-slots_1.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-failover-slots` | 1.1.0 | `u24.x86_64` | pigsty | 40.5 KiB | [postgresql-15-pg-failover-slots_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-failover-slots/postgresql-15-pg-failover-slots_1.1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_failover_slots_14` | 1.1.0 | `el8.x86_64` | pgdg | 24.5 KiB | [pg_failover_slots_14-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_failover_slots_14-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_failover_slots_14` | 1.1.0 | `el8.aarch64` | pgdg | 24.0 KiB | [pg_failover_slots_14-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_failover_slots_14-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_failover_slots_14` | 1.1.0 | `el8.x86_64` | pigsty | 23.7 KiB | [pg_failover_slots_14-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_failover_slots_14-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_failover_slots_14` | 1.1.0 | `el8.aarch64` | pigsty | 23.3 KiB | [pg_failover_slots_14-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_failover_slots_14-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_failover_slots_14` | 1.0.1 | `el8.x86_64` | pgdg | 23.7 KiB | [pg_failover_slots_14-1.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_failover_slots_14-1.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_failover_slots_14` | 1.0.1 | `el8.aarch64` | pgdg | 23.2 KiB | [pg_failover_slots_14-1.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_failover_slots_14-1.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_failover_slots_14` | 1.0.0 | `el8.x86_64` | pgdg | 23.0 KiB | [pg_failover_slots_14-1.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_failover_slots_14-1.0.0-1.rhel8.x86_64.rpm) |
| `pg_failover_slots_14` | 1.0.0 | `el8.aarch64` | pgdg | 22.4 KiB | [pg_failover_slots_14-1.0.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_failover_slots_14-1.0.0-1.rhel8.aarch64.rpm) |
| `pg_failover_slots_14` | 1.1.0 | `el9.x86_64` | pgdg | 24.6 KiB | [pg_failover_slots_14-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_failover_slots_14-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_failover_slots_14` | 1.1.0 | `el9.aarch64` | pigsty | 23.8 KiB | [pg_failover_slots_14-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_failover_slots_14-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_failover_slots_14` | 1.1.0 | `el9.aarch64` | pgdg | 24.5 KiB | [pg_failover_slots_14-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_failover_slots_14-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_failover_slots_14` | 1.1.0 | `el9.x86_64` | pigsty | 23.8 KiB | [pg_failover_slots_14-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_failover_slots_14-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_failover_slots_14` | 1.0.1 | `el9.aarch64` | pgdg | 23.5 KiB | [pg_failover_slots_14-1.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_failover_slots_14-1.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_failover_slots_14` | 1.0.1 | `el9.x86_64` | pgdg | 23.7 KiB | [pg_failover_slots_14-1.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_failover_slots_14-1.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_failover_slots_14` | 1.0.0 | `el9.aarch64` | pgdg | 22.8 KiB | [pg_failover_slots_14-1.0.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_failover_slots_14-1.0.0-1.rhel9.aarch64.rpm) |
| `pg_failover_slots_14` | 1.0.0 | `el9.x86_64` | pgdg | 23.1 KiB | [pg_failover_slots_14-1.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_failover_slots_14-1.0.0-1.rhel9.x86_64.rpm) |
| `postgresql-14-pg-failover-slots` | 1.2.0 | `d12.aarch64` | pgdg | 38.7 KiB | [postgresql-14-pg-failover-slots_1.2.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-14-pg-failover-slots_1.2.0-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-failover-slots` | 1.2.0 | `d12.x86_64` | pgdg | 39.4 KiB | [postgresql-14-pg-failover-slots_1.2.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-14-pg-failover-slots_1.2.0-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-failover-slots` | 1.1.0 | `d12.x86_64` | pigsty | 41.9 KiB | [postgresql-14-pg-failover-slots_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-failover-slots/postgresql-14-pg-failover-slots_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-failover-slots` | 1.1.0 | `d12.aarch64` | pigsty | 40.8 KiB | [postgresql-14-pg-failover-slots_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-failover-slots/postgresql-14-pg-failover-slots_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-failover-slots` | 1.2.0 | `u22.x86_64` | pgdg | 42.6 KiB | [postgresql-14-pg-failover-slots_1.2.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-14-pg-failover-slots_1.2.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-failover-slots` | 1.2.0 | `u22.aarch64` | pgdg | 41.7 KiB | [postgresql-14-pg-failover-slots_1.2.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-14-pg-failover-slots_1.2.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-failover-slots` | 1.1.0 | `u22.aarch64` | pigsty | 43.3 KiB | [postgresql-14-pg-failover-slots_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-failover-slots/postgresql-14-pg-failover-slots_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-failover-slots` | 1.1.0 | `u22.x86_64` | pigsty | 44.0 KiB | [postgresql-14-pg-failover-slots_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-failover-slots/postgresql-14-pg-failover-slots_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-failover-slots` | 1.2.0 | `u24.x86_64` | pgdg | 39.5 KiB | [postgresql-14-pg-failover-slots_1.2.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-14-pg-failover-slots_1.2.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-failover-slots` | 1.2.0 | `u24.aarch64` | pgdg | 38.6 KiB | [postgresql-14-pg-failover-slots_1.2.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-14-pg-failover-slots_1.2.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pg-failover-slots` | 1.1.0 | `u24.aarch64` | pigsty | 40.0 KiB | [postgresql-14-pg-failover-slots_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-failover-slots/postgresql-14-pg-failover-slots_1.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-failover-slots` | 1.1.0 | `u24.x86_64` | pigsty | 40.4 KiB | [postgresql-14-pg-failover-slots_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-failover-slots/postgresql-14-pg-failover-slots_1.1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_failover_slots_13` | 1.1.0 | `el8.aarch64` | pigsty | 23.3 KiB | [pg_failover_slots_13-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_failover_slots_13-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_failover_slots_13` | 1.1.0 | `el8.x86_64` | pgdg | 24.3 KiB | [pg_failover_slots_13-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_failover_slots_13-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_failover_slots_13` | 1.1.0 | `el8.x86_64` | pigsty | 23.6 KiB | [pg_failover_slots_13-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_failover_slots_13-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_failover_slots_13` | 1.1.0 | `el8.aarch64` | pgdg | 24.0 KiB | [pg_failover_slots_13-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_failover_slots_13-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_failover_slots_13` | 1.0.1 | `el8.x86_64` | pgdg | 23.5 KiB | [pg_failover_slots_13-1.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_failover_slots_13-1.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_failover_slots_13` | 1.0.1 | `el8.aarch64` | pgdg | 23.2 KiB | [pg_failover_slots_13-1.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_failover_slots_13-1.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_failover_slots_13` | 1.0.0 | `el8.aarch64` | pgdg | 22.4 KiB | [pg_failover_slots_13-1.0.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_failover_slots_13-1.0.0-1.rhel8.aarch64.rpm) |
| `pg_failover_slots_13` | 1.0.0 | `el8.x86_64` | pgdg | 22.8 KiB | [pg_failover_slots_13-1.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_failover_slots_13-1.0.0-1.rhel8.x86_64.rpm) |
| `pg_failover_slots_13` | 1.1.0 | `el9.aarch64` | pgdg | 24.6 KiB | [pg_failover_slots_13-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_failover_slots_13-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_failover_slots_13` | 1.1.0 | `el9.x86_64` | pigsty | 23.9 KiB | [pg_failover_slots_13-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_failover_slots_13-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_failover_slots_13` | 1.1.0 | `el9.x86_64` | pgdg | 24.6 KiB | [pg_failover_slots_13-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_failover_slots_13-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_failover_slots_13` | 1.1.0 | `el9.aarch64` | pigsty | 23.8 KiB | [pg_failover_slots_13-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_failover_slots_13-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_failover_slots_13` | 1.0.1 | `el9.x86_64` | pgdg | 23.7 KiB | [pg_failover_slots_13-1.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_failover_slots_13-1.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_failover_slots_13` | 1.0.1 | `el9.aarch64` | pgdg | 23.5 KiB | [pg_failover_slots_13-1.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_failover_slots_13-1.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_failover_slots_13` | 1.0.0 | `el9.aarch64` | pgdg | 22.8 KiB | [pg_failover_slots_13-1.0.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_failover_slots_13-1.0.0-1.rhel9.aarch64.rpm) |
| `pg_failover_slots_13` | 1.0.0 | `el9.x86_64` | pgdg | 23.1 KiB | [pg_failover_slots_13-1.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_failover_slots_13-1.0.0-1.rhel9.x86_64.rpm) |
| `postgresql-13-pg-failover-slots` | 1.2.0 | `d12.x86_64` | pgdg | 39.4 KiB | [postgresql-13-pg-failover-slots_1.2.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-13-pg-failover-slots_1.2.0-1.pgdg12+1_amd64.deb) |
| `postgresql-13-pg-failover-slots` | 1.2.0 | `d12.aarch64` | pgdg | 38.5 KiB | [postgresql-13-pg-failover-slots_1.2.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-13-pg-failover-slots_1.2.0-1.pgdg12+1_arm64.deb) |
| `postgresql-13-pg-failover-slots` | 1.1.0 | `d12.x86_64` | pigsty | 41.5 KiB | [postgresql-13-pg-failover-slots_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-failover-slots/postgresql-13-pg-failover-slots_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-failover-slots` | 1.1.0 | `d12.aarch64` | pigsty | 40.8 KiB | [postgresql-13-pg-failover-slots_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-failover-slots/postgresql-13-pg-failover-slots_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-failover-slots` | 1.2.0 | `u22.aarch64` | pgdg | 41.8 KiB | [postgresql-13-pg-failover-slots_1.2.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-13-pg-failover-slots_1.2.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pg-failover-slots` | 1.2.0 | `u22.x86_64` | pgdg | 42.6 KiB | [postgresql-13-pg-failover-slots_1.2.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-13-pg-failover-slots_1.2.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pg-failover-slots` | 1.1.0 | `u22.x86_64` | pigsty | 43.9 KiB | [postgresql-13-pg-failover-slots_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-failover-slots/postgresql-13-pg-failover-slots_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-failover-slots` | 1.1.0 | `u22.aarch64` | pigsty | 43.4 KiB | [postgresql-13-pg-failover-slots_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-failover-slots/postgresql-13-pg-failover-slots_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-failover-slots` | 1.2.0 | `u24.aarch64` | pgdg | 39.0 KiB | [postgresql-13-pg-failover-slots_1.2.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-13-pg-failover-slots_1.2.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-13-pg-failover-slots` | 1.2.0 | `u24.x86_64` | pgdg | 39.8 KiB | [postgresql-13-pg-failover-slots_1.2.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-failover-slots/postgresql-13-pg-failover-slots_1.2.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pg-failover-slots` | 1.1.0 | `u24.x86_64` | pigsty | 40.7 KiB | [postgresql-13-pg-failover-slots_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-failover-slots/postgresql-13-pg-failover-slots_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-failover-slots` | 1.1.0 | `u24.aarch64` | pigsty | 39.9 KiB | [postgresql-13-pg-failover-slots_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-failover-slots/postgresql-13-pg-failover-slots_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/EnterpriseDB/pg_failover_slots" title="Repository" icon="github" subtitle="github.com/EnterpriseDB/pg_failover_slots" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_failover_slots-1.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_failover_slots; # get pg_failover_slots source code
pig build dep pg_failover_slots; # install build dependencies
pig build pkg pg_failover_slots; # build extension rpm or deb
pig build ext pg_failover_slots; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_failover_slots; # install by extension name, for the current active PG version
pig ext install pg_failover_slots; # install via package alias, for the active PG version
pig ext install pg_failover_slots -v 17;   # install for PG 17
pig ext install pg_failover_slots -v 16;   # install for PG 16
pig ext install pg_failover_slots -v 15;   # install for PG 15
pig ext install pg_failover_slots -v 14;   # install for PG 14
pig ext install pg_failover_slots -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_failover_slots;
```

