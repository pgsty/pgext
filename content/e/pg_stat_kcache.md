---
title: "pg_stat_kcache"
linkTitle: "pg_stat_kcache"
description: "Kernel statistics gathering"
weight: 6220
categories: ["Stat"]
width: full
---

Kernel statistics gathering

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6220** | {{< badge content="pg_stat_kcache" link="https://github.com/powa-team/pg_stat_kcache" >}} | {{< ext "pg_stat_kcache" "pg_stat_kcache" >}} | `2.3.0` | {{< category "STAT" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sLd-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pg_stat_statements" >}} |
|   **See Also**    | {{< ext "pg_profile" >}} {{< ext "powa" >}} {{< ext "plprofiler" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_track_settings" >}} {{< ext "pg_wait_sampling" >}} {{< ext "system_stats" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_stat_kcache" >}} | `2.3.0` | {{< badge content="18" color="red" alt="pg_stat_kcache_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_stat_kcache_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pg_stat_kcache" >}} | `2.3.0` | {{< badge content="18" color="red" alt="postgresql-18-pg-stat-kcache" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-stat-kcache` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_stat_kcache_18" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_stat_kcache_18-2.3.1-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_stat_kcache_17" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_kcache_17-2.3.1-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_stat_kcache_16" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_kcache_16-2.3.1-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_stat_kcache_15" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_kcache_15-2.3.1-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_stat_kcache_14" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_kcache_14-2.3.1-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_stat_kcache_18" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_stat_kcache_18-2.3.1-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_stat_kcache_17" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_kcache_17-2.3.1-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_stat_kcache_16" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_kcache_16-2.3.1-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_stat_kcache_15" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_kcache_15-2.3.1-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_stat_kcache_14" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_kcache_14-2.3.1-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_stat_kcache_18" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_stat_kcache_18-2.3.1-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_stat_kcache_17" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_stat_kcache_17-2.3.1-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_stat_kcache_16" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_kcache_16-2.3.1-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_stat_kcache_15" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_kcache_15-2.3.1-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_stat_kcache_14" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_kcache_14-2.3.1-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_stat_kcache_18" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_stat_kcache_18-2.3.1-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_stat_kcache_17" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_stat_kcache_17-2.3.1-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_stat_kcache_16" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_kcache_16-2.3.1-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_stat_kcache_15" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_kcache_15-2.3.1-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_stat_kcache_14" "2.3.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_kcache_14-2.3.1-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pg-stat-kcache" "2.3.1" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_stat_kcache_18` | 2.3.1 | `el8.x86_64` | pgdg | 27.8 KiB | [pg_stat_kcache_18-2.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_stat_kcache_18-2.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_18` | 2.3.1 | `el8.aarch64` | pgdg | 27.7 KiB | [pg_stat_kcache_18-2.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_stat_kcache_18-2.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_18` | 2.3.1 | `el9.x86_64` | pgdg | 26.5 KiB | [pg_stat_kcache_18-2.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_stat_kcache_18-2.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_18` | 2.3.1 | `el9.aarch64` | pgdg | 26.6 KiB | [pg_stat_kcache_18-2.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_stat_kcache_18-2.3.1-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-18-pg-stat-kcache` | 2.3.1 | `d12.x86_64` | pgdg | 36.9 KiB | [postgresql-18-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | 2.3.1 | `d12.aarch64` | pgdg | 36.7 KiB | [postgresql-18-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-stat-kcache` | 2.3.1 | `u22.aarch64` | pgdg | 35.2 KiB | [postgresql-18-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-stat-kcache` | 2.3.1 | `u22.x86_64` | pgdg | 35.8 KiB | [postgresql-18-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | 2.3.1 | `u24.x86_64` | pgdg | 35.2 KiB | [postgresql-18-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | 2.3.1 | `u24.aarch64` | pgdg | 34.8 KiB | [postgresql-18-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_stat_kcache_17` | 2.3.1 | `el8.aarch64` | pgdg | 27.7 KiB | [pg_stat_kcache_17-2.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_kcache_17-2.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_17` | 2.3.1 | `el8.x86_64` | pgdg | 27.8 KiB | [pg_stat_kcache_17-2.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_kcache_17-2.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_17` | 2.3.0 | `el8.x86_64` | pgdg | 27.3 KiB | [pg_stat_kcache_17-2.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_kcache_17-2.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_17` | 2.3.0 | `el8.aarch64` | pgdg | 27.2 KiB | [pg_stat_kcache_17-2.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_kcache_17-2.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_17` | 2.3.1 | `el9.x86_64` | pgdg | 26.5 KiB | [pg_stat_kcache_17-2.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_stat_kcache_17-2.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_17` | 2.3.1 | `el9.aarch64` | pgdg | 26.5 KiB | [pg_stat_kcache_17-2.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_stat_kcache_17-2.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_17` | 2.3.0 | `el9.aarch64` | pgdg | 26.5 KiB | [pg_stat_kcache_17-2.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_stat_kcache_17-2.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_17` | 2.3.0 | `el9.x86_64` | pgdg | 26.4 KiB | [pg_stat_kcache_17-2.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_stat_kcache_17-2.3.0-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-17-pg-stat-kcache` | 2.3.1 | `d12.aarch64` | pgdg | 36.6 KiB | [postgresql-17-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | 2.3.1 | `d12.x86_64` | pgdg | 37.0 KiB | [postgresql-17-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | 2.3.1 | `u22.aarch64` | pgdg | 38.6 KiB | [postgresql-17-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | 2.3.1 | `u22.x86_64` | pgdg | 39.1 KiB | [postgresql-17-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | 2.3.1 | `u24.aarch64` | pgdg | 34.8 KiB | [postgresql-17-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | 2.3.1 | `u24.x86_64` | pgdg | 35.2 KiB | [postgresql-17-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_stat_kcache_16` | 2.3.1 | `el8.x86_64` | pgdg | 27.8 KiB | [pg_stat_kcache_16-2.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_kcache_16-2.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_16` | 2.3.1 | `el8.aarch64` | pgdg | 27.7 KiB | [pg_stat_kcache_16-2.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_kcache_16-2.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_16` | 2.3.0 | `el8.aarch64` | pgdg | 27.2 KiB | [pg_stat_kcache_16-2.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_kcache_16-2.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_16` | 2.3.0 | `el8.x86_64` | pgdg | 27.3 KiB | [pg_stat_kcache_16-2.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_kcache_16-2.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_16` | 2.2.3 | `el8.aarch64` | pgdg | 26.6 KiB | [pg_stat_kcache_16-2.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_kcache_16-2.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_16` | 2.2.3 | `el8.x86_64` | pgdg | 26.6 KiB | [pg_stat_kcache_16-2.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_kcache_16-2.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_16` | 2.2.2 | `el8.aarch64` | pgdg | 26.3 KiB | [pg_stat_kcache_16-2.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_kcache_16-2.2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_16` | 2.2.2 | `el8.x86_64` | pgdg | 26.3 KiB | [pg_stat_kcache_16-2.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_kcache_16-2.2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_16` | 2.3.1 | `el9.aarch64` | pgdg | 26.6 KiB | [pg_stat_kcache_16-2.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_kcache_16-2.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_16` | 2.3.1 | `el9.x86_64` | pgdg | 26.5 KiB | [pg_stat_kcache_16-2.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_kcache_16-2.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_16` | 2.3.0 | `el9.x86_64` | pgdg | 26.4 KiB | [pg_stat_kcache_16-2.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_kcache_16-2.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_16` | 2.3.0 | `el9.aarch64` | pgdg | 26.5 KiB | [pg_stat_kcache_16-2.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_kcache_16-2.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_16` | 2.2.3 | `el9.x86_64` | pgdg | 25.6 KiB | [pg_stat_kcache_16-2.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_kcache_16-2.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_16` | 2.2.3 | `el9.aarch64` | pgdg | 25.8 KiB | [pg_stat_kcache_16-2.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_kcache_16-2.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_16` | 2.2.2 | `el9.x86_64` | pgdg | 25.2 KiB | [pg_stat_kcache_16-2.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_kcache_16-2.2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_16` | 2.2.2 | `el9.aarch64` | pgdg | 25.3 KiB | [pg_stat_kcache_16-2.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_kcache_16-2.2.2-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-pg-stat-kcache` | 2.3.1 | `d12.aarch64` | pgdg | 36.6 KiB | [postgresql-16-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | 2.3.1 | `d12.x86_64` | pgdg | 37.0 KiB | [postgresql-16-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | 2.3.1 | `u22.aarch64` | pgdg | 38.5 KiB | [postgresql-16-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | 2.3.1 | `u22.x86_64` | pgdg | 39.1 KiB | [postgresql-16-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | 2.3.1 | `u24.aarch64` | pgdg | 34.7 KiB | [postgresql-16-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | 2.3.1 | `u24.x86_64` | pgdg | 35.2 KiB | [postgresql-16-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_stat_kcache_15` | 2.3.1 | `el8.x86_64` | pgdg | 27.8 KiB | [pg_stat_kcache_15-2.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_kcache_15-2.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_15` | 2.3.1 | `el8.aarch64` | pgdg | 27.8 KiB | [pg_stat_kcache_15-2.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_kcache_15-2.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_15` | 2.3.0 | `el8.aarch64` | pgdg | 27.3 KiB | [pg_stat_kcache_15-2.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_kcache_15-2.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_15` | 2.3.0 | `el8.x86_64` | pgdg | 27.4 KiB | [pg_stat_kcache_15-2.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_kcache_15-2.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_15` | 2.2.3 | `el8.x86_64` | pgdg | 26.7 KiB | [pg_stat_kcache_15-2.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_kcache_15-2.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_15` | 2.2.3 | `el8.aarch64` | pgdg | 26.6 KiB | [pg_stat_kcache_15-2.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_kcache_15-2.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_15` | 2.2.2 | `el8.x86_64` | pgdg | 26.3 KiB | [pg_stat_kcache_15-2.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_kcache_15-2.2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_15` | 2.2.2 | `el8.aarch64` | pgdg | 26.3 KiB | [pg_stat_kcache_15-2.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_kcache_15-2.2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_15` | 2.2.1 | `el8.x86_64` | pgdg | 47.2 KiB | [pg_stat_kcache_15-2.2.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_kcache_15-2.2.1-2.rhel8.x86_64.rpm) |
| `pg_stat_kcache_15` | 2.2.1 | `el8.aarch64` | pgdg | 46.4 KiB | [pg_stat_kcache_15-2.2.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_kcache_15-2.2.1-2.rhel8.aarch64.rpm) |
| `pg_stat_kcache_15` | 2.3.1 | `el9.x86_64` | pgdg | 26.6 KiB | [pg_stat_kcache_15-2.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_kcache_15-2.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_15` | 2.3.1 | `el9.aarch64` | pgdg | 26.6 KiB | [pg_stat_kcache_15-2.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_kcache_15-2.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_15` | 2.3.0 | `el9.x86_64` | pgdg | 26.5 KiB | [pg_stat_kcache_15-2.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_kcache_15-2.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_15` | 2.3.0 | `el9.aarch64` | pgdg | 26.6 KiB | [pg_stat_kcache_15-2.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_kcache_15-2.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_15` | 2.2.3 | `el9.aarch64` | pgdg | 25.9 KiB | [pg_stat_kcache_15-2.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_kcache_15-2.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_15` | 2.2.3 | `el9.x86_64` | pgdg | 25.7 KiB | [pg_stat_kcache_15-2.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_kcache_15-2.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_15` | 2.2.2 | `el9.aarch64` | pgdg | 25.4 KiB | [pg_stat_kcache_15-2.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_kcache_15-2.2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_15` | 2.2.2 | `el9.x86_64` | pgdg | 25.3 KiB | [pg_stat_kcache_15-2.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_kcache_15-2.2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_15` | 2.2.1 | `el9.aarch64` | pgdg | 46.6 KiB | [pg_stat_kcache_15-2.2.1-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_kcache_15-2.2.1-2.rhel9.aarch64.rpm) |
| `pg_stat_kcache_15` | 2.2.1 | `el9.x86_64` | pgdg | 47.4 KiB | [pg_stat_kcache_15-2.2.1-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_kcache_15-2.2.1-2.rhel9.x86_64.rpm) |
| `postgresql-15-pg-stat-kcache` | 2.3.1 | `d12.x86_64` | pgdg | 37.0 KiB | [postgresql-15-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | 2.3.1 | `d12.aarch64` | pgdg | 36.6 KiB | [postgresql-15-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | 2.3.1 | `u22.x86_64` | pgdg | 39.1 KiB | [postgresql-15-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | 2.3.1 | `u22.aarch64` | pgdg | 38.6 KiB | [postgresql-15-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | 2.3.1 | `u24.aarch64` | pgdg | 34.8 KiB | [postgresql-15-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | 2.3.1 | `u24.x86_64` | pgdg | 35.3 KiB | [postgresql-15-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_stat_kcache_14` | 2.3.1 | `el8.x86_64` | pgdg | 27.8 KiB | [pg_stat_kcache_14-2.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_kcache_14-2.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_14` | 2.3.1 | `el8.aarch64` | pgdg | 27.7 KiB | [pg_stat_kcache_14-2.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_kcache_14-2.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_14` | 2.3.0 | `el8.x86_64` | pgdg | 27.4 KiB | [pg_stat_kcache_14-2.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_kcache_14-2.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_14` | 2.3.0 | `el8.aarch64` | pgdg | 27.3 KiB | [pg_stat_kcache_14-2.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_kcache_14-2.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_14` | 2.2.3 | `el8.aarch64` | pgdg | 26.6 KiB | [pg_stat_kcache_14-2.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_kcache_14-2.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_14` | 2.2.3 | `el8.x86_64` | pgdg | 26.7 KiB | [pg_stat_kcache_14-2.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_kcache_14-2.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_14` | 2.2.2 | `el8.x86_64` | pgdg | 26.3 KiB | [pg_stat_kcache_14-2.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_kcache_14-2.2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_14` | 2.2.2 | `el8.aarch64` | pgdg | 26.3 KiB | [pg_stat_kcache_14-2.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_kcache_14-2.2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_14` | 2.2.1 | `el8.aarch64` | pgdg | 46.2 KiB | [pg_stat_kcache_14-2.2.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_kcache_14-2.2.1-2.rhel8.aarch64.rpm) |
| `pg_stat_kcache_14` | 2.2.1 | `el8.x86_64` | pgdg | 47.0 KiB | [pg_stat_kcache_14-2.2.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_kcache_14-2.2.1-2.rhel8.x86_64.rpm) |
| `pg_stat_kcache_14` | 2.2.0 | `el8.x86_64` | pgdg | 46.6 KiB | [pg_stat_kcache_14-2.2.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_kcache_14-2.2.0-2.rhel8.x86_64.rpm) |
| `pg_stat_kcache_14` | 2.3.1 | `el9.x86_64` | pgdg | 26.6 KiB | [pg_stat_kcache_14-2.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_kcache_14-2.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_14` | 2.3.1 | `el9.aarch64` | pgdg | 26.6 KiB | [pg_stat_kcache_14-2.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_kcache_14-2.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_14` | 2.3.0 | `el9.x86_64` | pgdg | 26.4 KiB | [pg_stat_kcache_14-2.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_kcache_14-2.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_14` | 2.3.0 | `el9.aarch64` | pgdg | 26.6 KiB | [pg_stat_kcache_14-2.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_kcache_14-2.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_14` | 2.2.3 | `el9.aarch64` | pgdg | 25.9 KiB | [pg_stat_kcache_14-2.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_kcache_14-2.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_14` | 2.2.3 | `el9.x86_64` | pgdg | 25.7 KiB | [pg_stat_kcache_14-2.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_kcache_14-2.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_14` | 2.2.2 | `el9.x86_64` | pgdg | 25.3 KiB | [pg_stat_kcache_14-2.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_kcache_14-2.2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_14` | 2.2.2 | `el9.aarch64` | pgdg | 25.4 KiB | [pg_stat_kcache_14-2.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_kcache_14-2.2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_14` | 2.2.1 | `el9.x86_64` | pgdg | 47.2 KiB | [pg_stat_kcache_14-2.2.1-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_kcache_14-2.2.1-2.rhel9.x86_64.rpm) |
| `pg_stat_kcache_14` | 2.2.1 | `el9.aarch64` | pgdg | 46.5 KiB | [pg_stat_kcache_14-2.2.1-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_kcache_14-2.2.1-2.rhel9.aarch64.rpm) |
| `postgresql-14-pg-stat-kcache` | 2.3.1 | `d12.x86_64` | pgdg | 36.9 KiB | [postgresql-14-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-stat-kcache` | 2.3.1 | `d12.aarch64` | pgdg | 36.6 KiB | [postgresql-14-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-stat-kcache` | 2.3.1 | `u22.x86_64` | pgdg | 38.9 KiB | [postgresql-14-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-stat-kcache` | 2.3.1 | `u22.aarch64` | pgdg | 38.4 KiB | [postgresql-14-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-stat-kcache` | 2.3.1 | `u24.x86_64` | pgdg | 35.3 KiB | [postgresql-14-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-stat-kcache` | 2.3.1 | `u24.aarch64` | pgdg | 34.8 KiB | [postgresql-14-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_stat_kcache_13` | 2.3.1 | `el8.aarch64` | pgdg | 27.8 KiB | [pg_stat_kcache_13-2.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_kcache_13-2.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_13` | 2.3.1 | `el8.x86_64` | pgdg | 27.7 KiB | [pg_stat_kcache_13-2.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_kcache_13-2.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_13` | 2.3.0 | `el8.aarch64` | pgdg | 27.3 KiB | [pg_stat_kcache_13-2.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_kcache_13-2.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_13` | 2.3.0 | `el8.x86_64` | pgdg | 27.3 KiB | [pg_stat_kcache_13-2.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_kcache_13-2.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_13` | 2.2.3 | `el8.x86_64` | pgdg | 26.6 KiB | [pg_stat_kcache_13-2.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_kcache_13-2.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_13` | 2.2.3 | `el8.aarch64` | pgdg | 26.7 KiB | [pg_stat_kcache_13-2.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_kcache_13-2.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_13` | 2.2.2 | `el8.x86_64` | pgdg | 26.2 KiB | [pg_stat_kcache_13-2.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_kcache_13-2.2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_13` | 2.2.2 | `el8.aarch64` | pgdg | 26.3 KiB | [pg_stat_kcache_13-2.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_kcache_13-2.2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_13` | 2.2.1 | `el8.x86_64` | pgdg | 46.7 KiB | [pg_stat_kcache_13-2.2.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_kcache_13-2.2.1-2.rhel8.x86_64.rpm) |
| `pg_stat_kcache_13` | 2.2.1 | `el8.aarch64` | pgdg | 46.1 KiB | [pg_stat_kcache_13-2.2.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_kcache_13-2.2.1-2.rhel8.aarch64.rpm) |
| `pg_stat_kcache_13` | 2.2.0 | `el8.x86_64` | pgdg | 46.4 KiB | [pg_stat_kcache_13-2.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_kcache_13-2.2.0-1.rhel8.x86_64.rpm) |
| `pg_stat_kcache_13` | 2.3.1 | `el9.x86_64` | pgdg | 26.6 KiB | [pg_stat_kcache_13-2.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_kcache_13-2.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_13` | 2.3.1 | `el9.aarch64` | pgdg | 26.7 KiB | [pg_stat_kcache_13-2.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_kcache_13-2.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_13` | 2.3.0 | `el9.x86_64` | pgdg | 26.4 KiB | [pg_stat_kcache_13-2.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_kcache_13-2.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_13` | 2.3.0 | `el9.aarch64` | pgdg | 26.6 KiB | [pg_stat_kcache_13-2.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_kcache_13-2.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_13` | 2.2.3 | `el9.aarch64` | pgdg | 25.9 KiB | [pg_stat_kcache_13-2.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_kcache_13-2.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_13` | 2.2.3 | `el9.x86_64` | pgdg | 25.7 KiB | [pg_stat_kcache_13-2.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_kcache_13-2.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_13` | 2.2.2 | `el9.x86_64` | pgdg | 25.3 KiB | [pg_stat_kcache_13-2.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_kcache_13-2.2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_13` | 2.2.2 | `el9.aarch64` | pgdg | 25.4 KiB | [pg_stat_kcache_13-2.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_kcache_13-2.2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_13` | 2.2.1 | `el9.aarch64` | pgdg | 46.2 KiB | [pg_stat_kcache_13-2.2.1-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_kcache_13-2.2.1-2.rhel9.aarch64.rpm) |
| `pg_stat_kcache_13` | 2.2.1 | `el9.x86_64` | pgdg | 47.0 KiB | [pg_stat_kcache_13-2.2.1-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_kcache_13-2.2.1-2.rhel9.x86_64.rpm) |
| `postgresql-13-pg-stat-kcache` | 2.3.1 | `d12.x86_64` | pgdg | 36.7 KiB | [postgresql-13-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-13-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb) |
| `postgresql-13-pg-stat-kcache` | 2.3.1 | `d12.aarch64` | pgdg | 36.5 KiB | [postgresql-13-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-13-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb) |
| `postgresql-13-pg-stat-kcache` | 2.3.1 | `u22.aarch64` | pgdg | 37.9 KiB | [postgresql-13-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-13-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pg-stat-kcache` | 2.3.1 | `u22.x86_64` | pgdg | 39.0 KiB | [postgresql-13-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-13-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pg-stat-kcache` | 2.3.1 | `u24.x86_64` | pgdg | 35.0 KiB | [postgresql-13-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-13-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pg-stat-kcache` | 2.3.1 | `u24.aarch64` | pgdg | 34.8 KiB | [postgresql-13-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-13-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/powa-team/pg_stat_kcache" title="Repository" icon="github" subtitle="github.com/powa-team/pg_stat_kcache" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_stat_kcache; # install by extension name, for the current active PG version
pig ext install pg_stat_kcache; # install via package alias, for the active PG version
pig ext install pg_stat_kcache -v 17;   # install for PG 17
pig ext install pg_stat_kcache -v 16;   # install for PG 16
pig ext install pg_stat_kcache -v 15;   # install for PG 15
pig ext install pg_stat_kcache -v 14;   # install for PG 14
pig ext install pg_stat_kcache -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_stat_kcache;
```

