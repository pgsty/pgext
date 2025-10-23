---
title: "pg_qualstats"
linkTitle: "pg_qualstats"
description: "An extension collecting statistics about quals"
weight: 6240
categories: ["STAT"]
width: full
---

An extension collecting statistics about quals


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6240** | {{< badge content="pg_qualstats" link="https://github.com/powa-team/pg_qualstats" >}} | {{< ext "pg_qualstats" >}} | `2.1.2` | {{< category "STAT" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "hypopg" >}} {{< ext "pg_stat_kcache" >}} {{< ext "powa" >}} {{< ext "pg_stat_statements" >}} {{< ext "index_advisor" >}} {{< ext "pre_prepare" >}} {{< ext "pg_show_plans" >}} {{< ext "pg_stat_monitor" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_qualstats" >}} | `2.1.1` | {{< bg "18" "pg_qualstats_18*" "green" >}} {{< bg "17" "pg_qualstats_17*" "green" >}} {{< bg "16" "pg_qualstats_16*" "green" >}} {{< bg "15" "pg_qualstats_15*" "green" >}} {{< bg "14" "pg_qualstats_14*" "green" >}} | `pg_qualstats_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pg_qualstats" >}} | `2.1.2` | {{< bg "18" "postgresql-18-pg-qualstats" "green" >}} {{< bg "17" "postgresql-17-pg-qualstats" "green" >}} {{< bg "16" "postgresql-16-pg-qualstats" "green" >}} {{< bg "15" "postgresql-15-pg-qualstats" "green" >}} {{< bg "14" "postgresql-14-pg-qualstats" "green" >}} | `postgresql-$v-pg-qualstats` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 2.1.2" "pg_qualstats_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.1" "pg_qualstats_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.1" "pg_qualstats_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.1" "pg_qualstats_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.1" "pg_qualstats_14 : AVAIL 4" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 2.1.2" "pg_qualstats_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.1" "pg_qualstats_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.1" "pg_qualstats_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.1" "pg_qualstats_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.1" "pg_qualstats_14 : AVAIL 3" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 2.1.2" "pg_qualstats_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.1" "pg_qualstats_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.1" "pg_qualstats_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.1" "pg_qualstats_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.1" "pg_qualstats_14 : AVAIL 3" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 2.1.2" "pg_qualstats_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.1" "pg_qualstats_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.1" "pg_qualstats_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.1" "pg_qualstats_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.1" "pg_qualstats_14 : AVAIL 3" "blue" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 2.1.3" "postgresql-18-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-17-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-16-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-15-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-14-pg-qualstats : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 2.1.3" "postgresql-18-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-17-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-16-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-15-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-14-pg-qualstats : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 2.1.3" "postgresql-18-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-17-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-16-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-15-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-14-pg-qualstats : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 2.1.3" "postgresql-18-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-17-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-16-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-15-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-14-pg-qualstats : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 2.1.3" "postgresql-18-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-17-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-16-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-15-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-14-pg-qualstats : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 2.1.3" "postgresql-18-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-17-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-16-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-15-pg-qualstats : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.3" "postgresql-14-pg-qualstats : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_qualstats_18` | 2.1.2 | `el8.x86_64` | pgdg | 38.0 KiB | [pg_qualstats_18-2.1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_qualstats_18-2.1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_qualstats_18` | 2.1.2 | `el8.aarch64` | pgdg | 37.1 KiB | [pg_qualstats_18-2.1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_qualstats_18-2.1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_qualstats_18` | 2.1.2 | `el9.x86_64` | pgdg | 36.4 KiB | [pg_qualstats_18-2.1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_qualstats_18-2.1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_qualstats_18` | 2.1.2 | `el9.aarch64` | pgdg | 35.7 KiB | [pg_qualstats_18-2.1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_qualstats_18-2.1.2-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-18-pg-qualstats` | 2.1.3 | `d12.x86_64` | pgdg | 56.7 KiB | [postgresql-18-pg-qualstats_2.1.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-18-pg-qualstats_2.1.3-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-qualstats` | 2.1.3 | `d12.aarch64` | pgdg | 55.6 KiB | [postgresql-18-pg-qualstats_2.1.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-18-pg-qualstats_2.1.3-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-qualstats` | 2.1.3 | `u22.x86_64` | pgdg | 56.4 KiB | [postgresql-18-pg-qualstats_2.1.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-18-pg-qualstats_2.1.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-qualstats` | 2.1.3 | `u22.aarch64` | pgdg | 54.8 KiB | [postgresql-18-pg-qualstats_2.1.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-18-pg-qualstats_2.1.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-qualstats` | 2.1.3 | `u24.x86_64` | pgdg | 54.1 KiB | [postgresql-18-pg-qualstats_2.1.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-18-pg-qualstats_2.1.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-qualstats` | 2.1.3 | `u24.aarch64` | pgdg | 53.1 KiB | [postgresql-18-pg-qualstats_2.1.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-18-pg-qualstats_2.1.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_qualstats_17` | 2.1.1 | `el8.x86_64` | pgdg | 37.6 KiB | [pg_qualstats_17-2.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_qualstats_17-2.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_qualstats_17` | 2.1.1 | `el8.aarch64` | pgdg | 36.7 KiB | [pg_qualstats_17-2.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_qualstats_17-2.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_qualstats_17` | 2.1.1 | `el9.x86_64` | pgdg | 36.2 KiB | [pg_qualstats_17-2.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_qualstats_17-2.1.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_qualstats_17` | 2.1.1 | `el9.aarch64` | pgdg | 35.5 KiB | [pg_qualstats_17-2.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_qualstats_17-2.1.1-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-pg-qualstats` | 2.1.3 | `d12.x86_64` | pgdg | 56.8 KiB | [postgresql-17-pg-qualstats_2.1.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-17-pg-qualstats_2.1.3-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-qualstats` | 2.1.3 | `d12.aarch64` | pgdg | 55.7 KiB | [postgresql-17-pg-qualstats_2.1.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-17-pg-qualstats_2.1.3-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-qualstats` | 2.1.3 | `u22.x86_64` | pgdg | 60.6 KiB | [postgresql-17-pg-qualstats_2.1.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-17-pg-qualstats_2.1.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-qualstats` | 2.1.3 | `u22.aarch64` | pgdg | 58.9 KiB | [postgresql-17-pg-qualstats_2.1.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-17-pg-qualstats_2.1.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-qualstats` | 2.1.3 | `u24.x86_64` | pgdg | 54.3 KiB | [postgresql-17-pg-qualstats_2.1.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-17-pg-qualstats_2.1.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-qualstats` | 2.1.3 | `u24.aarch64` | pgdg | 53.3 KiB | [postgresql-17-pg-qualstats_2.1.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-17-pg-qualstats_2.1.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_qualstats_16` | 2.1.1 | `el8.x86_64` | pgdg | 37.6 KiB | [pg_qualstats_16-2.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_qualstats_16-2.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_qualstats_16` | 2.1.0 | `el8.x86_64` | pgdg | 36.9 KiB | [pg_qualstats_16-2.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_qualstats_16-2.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_qualstats_16` | 2.0.4 | `el8.x86_64` | pgdg | 35.5 KiB | [pg_qualstats_16-2.0.4-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_qualstats_16-2.0.4-3PGDG.rhel8.x86_64.rpm) |
| `pg_qualstats_16` | 2.1.1 | `el8.aarch64` | pgdg | 36.7 KiB | [pg_qualstats_16-2.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_qualstats_16-2.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_qualstats_16` | 2.1.0 | `el8.aarch64` | pgdg | 36.0 KiB | [pg_qualstats_16-2.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_qualstats_16-2.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_qualstats_16` | 2.0.4 | `el8.aarch64` | pgdg | 34.7 KiB | [pg_qualstats_16-2.0.4-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_qualstats_16-2.0.4-3PGDG.rhel8.aarch64.rpm) |
| `pg_qualstats_16` | 2.1.1 | `el9.x86_64` | pgdg | 36.1 KiB | [pg_qualstats_16-2.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_qualstats_16-2.1.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_qualstats_16` | 2.1.0 | `el9.x86_64` | pgdg | 35.3 KiB | [pg_qualstats_16-2.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_qualstats_16-2.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_qualstats_16` | 2.0.4 | `el9.x86_64` | pgdg | 34.2 KiB | [pg_qualstats_16-2.0.4-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_qualstats_16-2.0.4-3PGDG.rhel9.x86_64.rpm) |
| `pg_qualstats_16` | 2.1.1 | `el9.aarch64` | pgdg | 35.6 KiB | [pg_qualstats_16-2.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_qualstats_16-2.1.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_qualstats_16` | 2.1.0 | `el9.aarch64` | pgdg | 34.7 KiB | [pg_qualstats_16-2.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_qualstats_16-2.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_qualstats_16` | 2.0.4 | `el9.aarch64` | pgdg | 33.4 KiB | [pg_qualstats_16-2.0.4-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_qualstats_16-2.0.4-3PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-pg-qualstats` | 2.1.3 | `d12.x86_64` | pgdg | 56.8 KiB | [postgresql-16-pg-qualstats_2.1.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-16-pg-qualstats_2.1.3-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-qualstats` | 2.1.3 | `d12.aarch64` | pgdg | 55.8 KiB | [postgresql-16-pg-qualstats_2.1.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-16-pg-qualstats_2.1.3-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-qualstats` | 2.1.3 | `u22.x86_64` | pgdg | 60.4 KiB | [postgresql-16-pg-qualstats_2.1.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-16-pg-qualstats_2.1.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-qualstats` | 2.1.3 | `u22.aarch64` | pgdg | 58.8 KiB | [postgresql-16-pg-qualstats_2.1.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-16-pg-qualstats_2.1.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-qualstats` | 2.1.3 | `u24.x86_64` | pgdg | 54.2 KiB | [postgresql-16-pg-qualstats_2.1.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-16-pg-qualstats_2.1.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-qualstats` | 2.1.3 | `u24.aarch64` | pgdg | 53.3 KiB | [postgresql-16-pg-qualstats_2.1.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-16-pg-qualstats_2.1.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_qualstats_15` | 2.1.1 | `el8.x86_64` | pgdg | 37.6 KiB | [pg_qualstats_15-2.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_qualstats_15-2.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_qualstats_15` | 2.1.0 | `el8.x86_64` | pgdg | 37.0 KiB | [pg_qualstats_15-2.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_qualstats_15-2.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_qualstats_15` | 2.0.4 | `el8.x86_64` | pgdg | 68.1 KiB | [pg_qualstats_15-2.0.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_qualstats_15-2.0.4-1.rhel8.x86_64.rpm) |
| `pg_qualstats_15` | 2.1.1 | `el8.aarch64` | pgdg | 36.7 KiB | [pg_qualstats_15-2.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_qualstats_15-2.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_qualstats_15` | 2.1.0 | `el8.aarch64` | pgdg | 36.1 KiB | [pg_qualstats_15-2.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_qualstats_15-2.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_qualstats_15` | 2.0.4 | `el8.aarch64` | pgdg | 66.9 KiB | [pg_qualstats_15-2.0.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_qualstats_15-2.0.4-1.rhel8.aarch64.rpm) |
| `pg_qualstats_15` | 2.1.1 | `el9.x86_64` | pgdg | 36.2 KiB | [pg_qualstats_15-2.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_qualstats_15-2.1.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_qualstats_15` | 2.1.0 | `el9.x86_64` | pgdg | 35.4 KiB | [pg_qualstats_15-2.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_qualstats_15-2.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_qualstats_15` | 2.0.4 | `el9.x86_64` | pgdg | 68.1 KiB | [pg_qualstats_15-2.0.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_qualstats_15-2.0.4-1.rhel9.x86_64.rpm) |
| `pg_qualstats_15` | 2.1.1 | `el9.aarch64` | pgdg | 35.6 KiB | [pg_qualstats_15-2.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_qualstats_15-2.1.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_qualstats_15` | 2.1.0 | `el9.aarch64` | pgdg | 34.7 KiB | [pg_qualstats_15-2.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_qualstats_15-2.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_qualstats_15` | 2.0.4 | `el9.aarch64` | pgdg | 67.0 KiB | [pg_qualstats_15-2.0.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_qualstats_15-2.0.4-1.rhel9.aarch64.rpm) |
| `postgresql-15-pg-qualstats` | 2.1.3 | `d12.x86_64` | pgdg | 56.8 KiB | [postgresql-15-pg-qualstats_2.1.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-15-pg-qualstats_2.1.3-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-qualstats` | 2.1.3 | `d12.aarch64` | pgdg | 55.8 KiB | [postgresql-15-pg-qualstats_2.1.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-15-pg-qualstats_2.1.3-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-qualstats` | 2.1.3 | `u22.x86_64` | pgdg | 60.5 KiB | [postgresql-15-pg-qualstats_2.1.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-15-pg-qualstats_2.1.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-qualstats` | 2.1.3 | `u22.aarch64` | pgdg | 58.8 KiB | [postgresql-15-pg-qualstats_2.1.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-15-pg-qualstats_2.1.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-qualstats` | 2.1.3 | `u24.x86_64` | pgdg | 54.3 KiB | [postgresql-15-pg-qualstats_2.1.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-15-pg-qualstats_2.1.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-qualstats` | 2.1.3 | `u24.aarch64` | pgdg | 53.3 KiB | [postgresql-15-pg-qualstats_2.1.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-15-pg-qualstats_2.1.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_qualstats_14` | 2.1.1 | `el8.x86_64` | pgdg | 37.7 KiB | [pg_qualstats_14-2.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_qualstats_14-2.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_qualstats_14` | 2.1.0 | `el8.x86_64` | pgdg | 37.1 KiB | [pg_qualstats_14-2.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_qualstats_14-2.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_qualstats_14` | 2.0.4 | `el8.x86_64` | pgdg | 68.7 KiB | [pg_qualstats_14-2.0.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_qualstats_14-2.0.4-1.rhel8.x86_64.rpm) |
| `pg_qualstats_14` | 2.0.3 | `el8.x86_64` | pgdg | 67.2 KiB | [pg_qualstats_14-2.0.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_qualstats_14-2.0.3-1.rhel8.x86_64.rpm) |
| `pg_qualstats_14` | 2.1.1 | `el8.aarch64` | pgdg | 36.8 KiB | [pg_qualstats_14-2.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_qualstats_14-2.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_qualstats_14` | 2.1.0 | `el8.aarch64` | pgdg | 36.1 KiB | [pg_qualstats_14-2.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_qualstats_14-2.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_qualstats_14` | 2.0.4 | `el8.aarch64` | pgdg | 67.0 KiB | [pg_qualstats_14-2.0.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_qualstats_14-2.0.4-1.rhel8.aarch64.rpm) |
| `pg_qualstats_14` | 2.1.1 | `el9.x86_64` | pgdg | 36.2 KiB | [pg_qualstats_14-2.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_qualstats_14-2.1.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_qualstats_14` | 2.1.0 | `el9.x86_64` | pgdg | 35.4 KiB | [pg_qualstats_14-2.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_qualstats_14-2.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_qualstats_14` | 2.0.4 | `el9.x86_64` | pgdg | 68.6 KiB | [pg_qualstats_14-2.0.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_qualstats_14-2.0.4-1.rhel9.x86_64.rpm) |
| `pg_qualstats_14` | 2.1.1 | `el9.aarch64` | pgdg | 35.5 KiB | [pg_qualstats_14-2.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_qualstats_14-2.1.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_qualstats_14` | 2.1.0 | `el9.aarch64` | pgdg | 34.8 KiB | [pg_qualstats_14-2.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_qualstats_14-2.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_qualstats_14` | 2.0.4 | `el9.aarch64` | pgdg | 67.0 KiB | [pg_qualstats_14-2.0.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_qualstats_14-2.0.4-1.rhel9.aarch64.rpm) |
| `postgresql-14-pg-qualstats` | 2.1.3 | `d12.x86_64` | pgdg | 57.2 KiB | [postgresql-14-pg-qualstats_2.1.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-14-pg-qualstats_2.1.3-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-qualstats` | 2.1.3 | `d12.aarch64` | pgdg | 56.1 KiB | [postgresql-14-pg-qualstats_2.1.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-14-pg-qualstats_2.1.3-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-qualstats` | 2.1.3 | `u22.x86_64` | pgdg | 61.0 KiB | [postgresql-14-pg-qualstats_2.1.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-14-pg-qualstats_2.1.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-qualstats` | 2.1.3 | `u22.aarch64` | pgdg | 59.5 KiB | [postgresql-14-pg-qualstats_2.1.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-14-pg-qualstats_2.1.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-qualstats` | 2.1.3 | `u24.x86_64` | pgdg | 54.6 KiB | [postgresql-14-pg-qualstats_2.1.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-14-pg-qualstats_2.1.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-qualstats` | 2.1.3 | `u24.aarch64` | pgdg | 53.7 KiB | [postgresql-14-pg-qualstats_2.1.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-qualstats/postgresql-14-pg-qualstats_2.1.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/powa-team/pg_qualstats" title="Repository" icon="github" subtitle="github.com/powa-team/pg_qualstats" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_qualstats; # install by extension name, for the current active PG version
pig ext install pg_qualstats; # install via package alias, for the active PG version
pig ext install pg_qualstats -v 18;   # install for PG 18
pig ext install pg_qualstats -v 17;   # install for PG 17
pig ext install pg_qualstats -v 16;   # install for PG 16
pig ext install pg_qualstats -v 15;   # install for PG 15
pig ext install pg_qualstats -v 14;   # install for PG 14
pig ext install pg_qualstats -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_qualstats;
```

