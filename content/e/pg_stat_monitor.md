---
title: "pg_stat_monitor"
linkTitle: "pg_stat_monitor"
description: "The pg_stat_monitor is a PostgreSQL Query Performance Monitoring tool, based on PostgreSQL contrib module pg_stat_statements. pg_stat_monitor provides aggregated statistics, client information, plan details including plan, and histogram information."
weight: 6230
categories: ["Stat"]
width: full
---

The pg_stat_monitor is a PostgreSQL Query Performance Monitoring tool, based on PostgreSQL contrib module pg_stat_statements. pg_stat_monitor provides aggregated statistics, client information, plan details including plan, and histogram information.

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6230** | {{< badge content="pg_stat_monitor" link="https://github.com/percona/pg_stat_monitor" >}} | {{< ext "pg_stat_monitor" "pg_stat_monitor" >}} | `2.2.0` | {{< category "STAT" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sLd-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_show_plans" >}} {{< ext "pg_stat_kcache" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_store_plans" >}} {{< ext "pgsentinel" >}} {{< ext "auto_explain" >}} {{< ext "logerrors" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_stat_monitor" >}} | `2.2.0` | {{< badge content="18" color="red" alt="pg_stat_monitor_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_stat_monitor_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_stat_monitor" >}} | `2.2.0` | {{< badge content="18" color="red" alt="postgresql-18-pg-stat-monitor" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-stat-monitor` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_stat_monitor_18" >}}     | {{< pkg "pg_stat_monitor_17" "2.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_monitor_17-2.2.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_stat_monitor_16" "2.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.2.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_stat_monitor_15" "2.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.2.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_stat_monitor_14" "2.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-2.2.0-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_stat_monitor_18" >}}     | {{< pkg "pg_stat_monitor_17" "2.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_monitor_17-2.2.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_stat_monitor_16" "2.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.2.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_stat_monitor_15" "2.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.2.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_stat_monitor_14" "2.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_monitor_14-2.2.0-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_stat_monitor_18" >}}     | {{< pkg "pg_stat_monitor_17" "2.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_stat_monitor_17-2.2.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_stat_monitor_16" "2.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_monitor_16-2.2.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_stat_monitor_15" "2.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_monitor_15-2.2.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_stat_monitor_14" "2.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_monitor_14-2.2.0-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_stat_monitor_18" >}}     | {{< pkg "pg_stat_monitor_17" "2.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_stat_monitor_17-2.2.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_stat_monitor_16" "2.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_monitor_16-2.2.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_stat_monitor_15" "2.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_monitor_15-2.2.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_stat_monitor_14" "2.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_monitor_14-2.2.0-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-stat-monitor" >}}     | {{< pkg "postgresql-17-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-stat-monitor" >}}     | {{< pkg "postgresql-17-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-stat-monitor" >}}     | {{< pkg "postgresql-17-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.2.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.2.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.2.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.2.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-stat-monitor" >}}     | {{< pkg "postgresql-17-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.2.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.2.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.2.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.2.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-stat-monitor" >}}     | {{< pkg "postgresql-17-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.2.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.2.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.2.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.2.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-stat-monitor" >}}     | {{< pkg "postgresql-17-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.2.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.2.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.2.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-stat-monitor" "2.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.2.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_stat_monitor_17` | 2.2.0 | `el8.aarch64` | pgdg | 40.7 KiB | [pg_stat_monitor_17-2.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_monitor_17-2.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_17` | 2.2.0 | `el8.x86_64` | pgdg | 42.2 KiB | [pg_stat_monitor_17-2.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_monitor_17-2.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_17` | 2.1.1 | `el8.aarch64` | pgdg | 40.1 KiB | [pg_stat_monitor_17-2.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_monitor_17-2.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_17` | 2.1.1 | `el8.x86_64` | pgdg | 41.7 KiB | [pg_stat_monitor_17-2.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_monitor_17-2.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_17` | 2.1.0 | `el8.x86_64` | pgdg | 40.2 KiB | [pg_stat_monitor_17-2.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_monitor_17-2.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_17` | 2.1.0 | `el8.aarch64` | pgdg | 38.6 KiB | [pg_stat_monitor_17-2.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_monitor_17-2.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_17` | 2.2.0 | `el9.x86_64` | pgdg | 41.2 KiB | [pg_stat_monitor_17-2.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_stat_monitor_17-2.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_17` | 2.2.0 | `el9.aarch64` | pgdg | 40.1 KiB | [pg_stat_monitor_17-2.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_stat_monitor_17-2.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_17` | 2.1.1 | `el9.x86_64` | pgdg | 40.9 KiB | [pg_stat_monitor_17-2.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_stat_monitor_17-2.1.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_17` | 2.1.1 | `el9.aarch64` | pgdg | 39.9 KiB | [pg_stat_monitor_17-2.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_stat_monitor_17-2.1.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_17` | 2.1.0 | `el9.aarch64` | pgdg | 38.5 KiB | [pg_stat_monitor_17-2.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_stat_monitor_17-2.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_17` | 2.1.0 | `el9.x86_64` | pgdg | 39.5 KiB | [pg_stat_monitor_17-2.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_stat_monitor_17-2.1.0-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-17-pg-stat-monitor` | 2.2.0 | `d12.x86_64` | pigsty | 72.9 KiB | [postgresql-17-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-stat-monitor` | 2.2.0 | `d12.aarch64` | pigsty | 71.2 KiB | [postgresql-17-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-stat-monitor` | 2.2.0 | `u22.aarch64` | pigsty | 83.5 KiB | [postgresql-17-pg-stat-monitor_2.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-stat-monitor` | 2.2.0 | `u22.x86_64` | pigsty | 85.3 KiB | [postgresql-17-pg-stat-monitor_2.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-stat-monitor` | 2.2.0 | `u24.x86_64` | pigsty | 76.0 KiB | [postgresql-17-pg-stat-monitor_2.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-stat-monitor` | 2.2.0 | `u24.aarch64` | pigsty | 75.0 KiB | [postgresql-17-pg-stat-monitor_2.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_stat_monitor_16` | 2.2.0 | `el8.aarch64` | pgdg | 40.7 KiB | [pg_stat_monitor_16-2.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_16` | 2.2.0 | `el8.x86_64` | pgdg | 42.1 KiB | [pg_stat_monitor_16-2.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_16` | 2.1.1 | `el8.x86_64` | pgdg | 41.6 KiB | [pg_stat_monitor_16-2.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_16` | 2.1.1 | `el8.aarch64` | pgdg | 40.1 KiB | [pg_stat_monitor_16-2.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_16` | 2.1.0 | `el8.x86_64` | pgdg | 40.0 KiB | [pg_stat_monitor_16-2.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_16` | 2.1.0 | `el8.aarch64` | pgdg | 38.6 KiB | [pg_stat_monitor_16-2.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_16` | 2.0.4 | `el8.x86_64` | pgdg | 38.7 KiB | [pg_stat_monitor_16-2.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.0.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_16` | 2.0.4 | `el8.aarch64` | pgdg | 37.3 KiB | [pg_stat_monitor_16-2.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.0.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_16` | 2.0.2 | `el8.x86_64` | pgdg | 38.5 KiB | [pg_stat_monitor_16-2.0.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.0.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_16` | 2.0.2 | `el8.aarch64` | pgdg | 37.1 KiB | [pg_stat_monitor_16-2.0.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.0.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_16` | 2.2.0 | `el9.aarch64` | pgdg | 40.0 KiB | [pg_stat_monitor_16-2.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_monitor_16-2.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_16` | 2.2.0 | `el9.x86_64` | pgdg | 41.0 KiB | [pg_stat_monitor_16-2.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_monitor_16-2.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_16` | 2.1.1 | `el9.aarch64` | pgdg | 39.8 KiB | [pg_stat_monitor_16-2.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_monitor_16-2.1.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_16` | 2.1.1 | `el9.x86_64` | pgdg | 40.7 KiB | [pg_stat_monitor_16-2.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_monitor_16-2.1.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_16` | 2.1.0 | `el9.x86_64` | pgdg | 39.3 KiB | [pg_stat_monitor_16-2.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_monitor_16-2.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_16` | 2.1.0 | `el9.aarch64` | pgdg | 38.3 KiB | [pg_stat_monitor_16-2.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_monitor_16-2.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_16` | 2.0.4 | `el9.aarch64` | pgdg | 37.2 KiB | [pg_stat_monitor_16-2.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_monitor_16-2.0.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_16` | 2.0.4 | `el9.x86_64` | pgdg | 38.1 KiB | [pg_stat_monitor_16-2.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_monitor_16-2.0.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_16` | 2.0.2 | `el9.aarch64` | pgdg | 36.9 KiB | [pg_stat_monitor_16-2.0.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_monitor_16-2.0.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_16` | 2.0.2 | `el9.x86_64` | pgdg | 38.0 KiB | [pg_stat_monitor_16-2.0.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_monitor_16-2.0.2-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-16-pg-stat-monitor` | 2.2.0 | `d12.aarch64` | pigsty | 70.8 KiB | [postgresql-16-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-stat-monitor` | 2.2.0 | `d12.x86_64` | pigsty | 72.5 KiB | [postgresql-16-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-stat-monitor` | 2.2.0 | `u22.x86_64` | pigsty | 84.8 KiB | [postgresql-16-pg-stat-monitor_2.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-stat-monitor` | 2.2.0 | `u22.aarch64` | pigsty | 83.0 KiB | [postgresql-16-pg-stat-monitor_2.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-stat-monitor` | 2.2.0 | `u24.aarch64` | pigsty | 75.9 KiB | [postgresql-16-pg-stat-monitor_2.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-stat-monitor` | 2.2.0 | `u24.x86_64` | pigsty | 76.9 KiB | [postgresql-16-pg-stat-monitor_2.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.2.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_stat_monitor_15` | 2.2.0 | `el8.x86_64` | pgdg | 43.3 KiB | [pg_stat_monitor_15-2.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_15` | 2.2.0 | `el8.aarch64` | pgdg | 41.5 KiB | [pg_stat_monitor_15-2.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_15` | 2.1.1 | `el8.aarch64` | pgdg | 41.0 KiB | [pg_stat_monitor_15-2.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_15` | 2.1.1 | `el8.x86_64` | pgdg | 42.8 KiB | [pg_stat_monitor_15-2.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_15` | 2.1.0 | `el8.x86_64` | pgdg | 41.2 KiB | [pg_stat_monitor_15-2.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_15` | 2.1.0 | `el8.aarch64` | pgdg | 39.5 KiB | [pg_stat_monitor_15-2.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_15` | 2.0.4 | `el8.aarch64` | pgdg | 38.1 KiB | [pg_stat_monitor_15-2.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.0.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_15` | 2.0.4 | `el8.x86_64` | pgdg | 39.9 KiB | [pg_stat_monitor_15-2.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.0.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_15` | 2.0.1 | `el8.x86_64` | pgdg | 39.5 KiB | [pg_stat_monitor_15-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_15` | 2.0.1 | `el8.aarch64` | pgdg | 37.8 KiB | [pg_stat_monitor_15-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_15` | 1.1.0 | `el8.aarch64` | pgdg | 84.7 KiB | [pg_stat_monitor_15-1.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-1.1.0-1.rhel8.aarch64.rpm) |
| `pg_stat_monitor_15` | 1.1.0 | `el8.x86_64` | pgdg | 87.1 KiB | [pg_stat_monitor_15-1.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-1.1.0-1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_15` | 2.2.0 | `el9.aarch64` | pgdg | 41.7 KiB | [pg_stat_monitor_15-2.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_monitor_15-2.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_15` | 2.2.0 | `el9.x86_64` | pgdg | 42.6 KiB | [pg_stat_monitor_15-2.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_monitor_15-2.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_15` | 2.1.1 | `el9.x86_64` | pgdg | 42.5 KiB | [pg_stat_monitor_15-2.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_monitor_15-2.1.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_15` | 2.1.1 | `el9.aarch64` | pgdg | 41.4 KiB | [pg_stat_monitor_15-2.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_monitor_15-2.1.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_15` | 2.1.0 | `el9.aarch64` | pgdg | 40.1 KiB | [pg_stat_monitor_15-2.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_monitor_15-2.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_15` | 2.1.0 | `el9.x86_64` | pgdg | 41.1 KiB | [pg_stat_monitor_15-2.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_monitor_15-2.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_15` | 2.0.4 | `el9.x86_64` | pgdg | 40.0 KiB | [pg_stat_monitor_15-2.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_monitor_15-2.0.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_15` | 2.0.4 | `el9.aarch64` | pgdg | 38.9 KiB | [pg_stat_monitor_15-2.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_monitor_15-2.0.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_15` | 2.0.1 | `el9.x86_64` | pgdg | 39.7 KiB | [pg_stat_monitor_15-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_monitor_15-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_15` | 2.0.1 | `el9.aarch64` | pgdg | 38.6 KiB | [pg_stat_monitor_15-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_monitor_15-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_15` | 1.1.0 | `el9.x86_64` | pgdg | 88.5 KiB | [pg_stat_monitor_15-1.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_monitor_15-1.1.0-1.rhel9.x86_64.rpm) |
| `pg_stat_monitor_15` | 1.1.0 | `el9.aarch64` | pgdg | 87.1 KiB | [pg_stat_monitor_15-1.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_monitor_15-1.1.0-1.rhel9.aarch64.rpm) |
| `postgresql-15-pg-stat-monitor` | 2.2.0 | `d12.aarch64` | pigsty | 79.9 KiB | [postgresql-15-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-stat-monitor` | 2.2.0 | `d12.x86_64` | pigsty | 81.4 KiB | [postgresql-15-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-stat-monitor` | 2.2.0 | `u22.x86_64` | pigsty | 86.9 KiB | [postgresql-15-pg-stat-monitor_2.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-stat-monitor` | 2.2.0 | `u22.aarch64` | pigsty | 85.2 KiB | [postgresql-15-pg-stat-monitor_2.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-stat-monitor` | 2.2.0 | `u24.aarch64` | pigsty | 76.3 KiB | [postgresql-15-pg-stat-monitor_2.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-stat-monitor` | 2.2.0 | `u24.x86_64` | pigsty | 77.2 KiB | [postgresql-15-pg-stat-monitor_2.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.2.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_stat_monitor_14` | 2.2.0 | `el8.aarch64` | pgdg | 41.4 KiB | [pg_stat_monitor_14-2.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_monitor_14-2.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_14` | 2.2.0 | `el8.x86_64` | pgdg | 43.1 KiB | [pg_stat_monitor_14-2.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-2.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | 2.1.1 | `el8.x86_64` | pgdg | 42.6 KiB | [pg_stat_monitor_14-2.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-2.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | 2.1.1 | `el8.aarch64` | pgdg | 40.9 KiB | [pg_stat_monitor_14-2.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_monitor_14-2.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_14` | 2.1.0 | `el8.aarch64` | pgdg | 39.4 KiB | [pg_stat_monitor_14-2.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_monitor_14-2.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_14` | 2.1.0 | `el8.x86_64` | pgdg | 41.1 KiB | [pg_stat_monitor_14-2.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-2.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | 2.0.4 | `el8.x86_64` | pgdg | 39.8 KiB | [pg_stat_monitor_14-2.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-2.0.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | 2.0.4 | `el8.aarch64` | pgdg | 38.0 KiB | [pg_stat_monitor_14-2.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_monitor_14-2.0.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_14` | 2.0.1 | `el8.x86_64` | pgdg | 39.4 KiB | [pg_stat_monitor_14-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | 2.0.1 | `el8.aarch64` | pgdg | 37.7 KiB | [pg_stat_monitor_14-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_monitor_14-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_14` | 1.1.0 | `el8.aarch64` | pgdg | 84.8 KiB | [pg_stat_monitor_14-1.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_monitor_14-1.1.0-1.rhel8.aarch64.rpm) |
| `pg_stat_monitor_14` | 1.1.0 | `el8.x86_64` | pgdg | 87.0 KiB | [pg_stat_monitor_14-1.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-1.1.0-1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | 1.0.1 | `el8.x86_64` | pgdg | 106.7 KiB | [pg_stat_monitor_14-1.0.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-1.0.1-1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | 1.0.0 | `el8.x86_64` | pgdg | 107.0 KiB | [pg_stat_monitor_14-1.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-1.0.0-1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | 1.0.0 | `el8.x86_64` | pgdg | 93.4 KiB | [pg_stat_monitor_14-1.0.0-rc.1_1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-1.0.0-rc.1_1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | 0.9.2 | `el8.x86_64` | pgdg | 86.9 KiB | [pg_stat_monitor_14-0.9.2-beta1_1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-0.9.2-beta1_1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | 2.2.0 | `el9.aarch64` | pgdg | 41.6 KiB | [pg_stat_monitor_14-2.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_monitor_14-2.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_14` | 2.2.0 | `el9.x86_64` | pgdg | 42.4 KiB | [pg_stat_monitor_14-2.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_monitor_14-2.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_14` | 2.1.1 | `el9.x86_64` | pgdg | 42.2 KiB | [pg_stat_monitor_14-2.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_monitor_14-2.1.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_14` | 2.1.1 | `el9.aarch64` | pgdg | 41.2 KiB | [pg_stat_monitor_14-2.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_monitor_14-2.1.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_14` | 2.1.0 | `el9.x86_64` | pgdg | 40.8 KiB | [pg_stat_monitor_14-2.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_monitor_14-2.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_14` | 2.1.0 | `el9.aarch64` | pgdg | 40.0 KiB | [pg_stat_monitor_14-2.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_monitor_14-2.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_14` | 2.0.4 | `el9.x86_64` | pgdg | 39.8 KiB | [pg_stat_monitor_14-2.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_monitor_14-2.0.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_14` | 2.0.4 | `el9.aarch64` | pgdg | 38.8 KiB | [pg_stat_monitor_14-2.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_monitor_14-2.0.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_14` | 2.0.1 | `el9.x86_64` | pgdg | 39.5 KiB | [pg_stat_monitor_14-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_monitor_14-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_14` | 2.0.1 | `el9.aarch64` | pgdg | 38.5 KiB | [pg_stat_monitor_14-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_monitor_14-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_14` | 1.1.0 | `el9.aarch64` | pgdg | 87.2 KiB | [pg_stat_monitor_14-1.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_monitor_14-1.1.0-1.rhel9.aarch64.rpm) |
| `pg_stat_monitor_14` | 1.1.0 | `el9.x86_64` | pgdg | 88.5 KiB | [pg_stat_monitor_14-1.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_monitor_14-1.1.0-1.rhel9.x86_64.rpm) |
| `pg_stat_monitor_14` | 1.0.1 | `el9.x86_64` | pgdg | 108.1 KiB | [pg_stat_monitor_14-1.0.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_monitor_14-1.0.1-1.rhel9.x86_64.rpm) |
| `pg_stat_monitor_14` | 1.0.0 | `el9.x86_64` | pgdg | 94.1 KiB | [pg_stat_monitor_14-1.0.0-rc.1_1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_monitor_14-1.0.0-rc.1_1.rhel9.x86_64.rpm) |
| `pg_stat_monitor_14` | 1.0.0 | `el9.x86_64` | pgdg | 107.5 KiB | [pg_stat_monitor_14-1.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_monitor_14-1.0.0-1.rhel9.x86_64.rpm) |
| `postgresql-14-pg-stat-monitor` | 2.2.0 | `d12.aarch64` | pigsty | 71.8 KiB | [postgresql-14-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-stat-monitor` | 2.2.0 | `d12.x86_64` | pigsty | 73.4 KiB | [postgresql-14-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-stat-monitor` | 2.2.0 | `u22.aarch64` | pigsty | 84.8 KiB | [postgresql-14-pg-stat-monitor_2.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-stat-monitor` | 2.2.0 | `u22.x86_64` | pigsty | 85.8 KiB | [postgresql-14-pg-stat-monitor_2.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-stat-monitor` | 2.2.0 | `u24.aarch64` | pigsty | 75.8 KiB | [postgresql-14-pg-stat-monitor_2.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-stat-monitor` | 2.2.0 | `u24.x86_64` | pigsty | 77.0 KiB | [postgresql-14-pg-stat-monitor_2.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.2.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_stat_monitor_13` | 2.2.0 | `el8.x86_64` | pgdg | 44.9 KiB | [pg_stat_monitor_13-2.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-2.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | 2.2.0 | `el8.aarch64` | pgdg | 43.1 KiB | [pg_stat_monitor_13-2.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_monitor_13-2.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_13` | 2.1.1 | `el8.x86_64` | pgdg | 44.4 KiB | [pg_stat_monitor_13-2.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-2.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | 2.1.1 | `el8.aarch64` | pgdg | 42.6 KiB | [pg_stat_monitor_13-2.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_monitor_13-2.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_13` | 2.1.0 | `el8.x86_64` | pgdg | 42.9 KiB | [pg_stat_monitor_13-2.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-2.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | 2.1.0 | `el8.aarch64` | pgdg | 41.1 KiB | [pg_stat_monitor_13-2.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_monitor_13-2.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_13` | 2.0.4 | `el8.aarch64` | pgdg | 39.7 KiB | [pg_stat_monitor_13-2.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_monitor_13-2.0.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_13` | 2.0.4 | `el8.x86_64` | pgdg | 41.4 KiB | [pg_stat_monitor_13-2.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-2.0.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | 2.0.1 | `el8.aarch64` | pgdg | 39.3 KiB | [pg_stat_monitor_13-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_monitor_13-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_13` | 2.0.1 | `el8.x86_64` | pgdg | 41.1 KiB | [pg_stat_monitor_13-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | 1.1.0 | `el8.x86_64` | pgdg | 95.8 KiB | [pg_stat_monitor_13-1.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-1.1.0-1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | 1.1.0 | `el8.aarch64` | pgdg | 93.5 KiB | [pg_stat_monitor_13-1.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_monitor_13-1.1.0-1.rhel8.aarch64.rpm) |
| `pg_stat_monitor_13` | 1.0.1 | `el8.x86_64` | pgdg | 115.5 KiB | [pg_stat_monitor_13-1.0.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-1.0.1-1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | 1.0.0 | `el8.x86_64` | pgdg | 103.8 KiB | [pg_stat_monitor_13-1.0.0-rc.1_1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-1.0.0-rc.1_1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | 1.0.0 | `el8.x86_64` | pgdg | 117.3 KiB | [pg_stat_monitor_13-1.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-1.0.0-1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | 0.9.1 | `el8.x86_64` | pgdg | 94.2 KiB | [pg_stat_monitor_13-0.9.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-0.9.1-1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | 0.7.2 | `el8.x86_64` | pgdg | 86.6 KiB | [pg_stat_monitor_13-0.7.2-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-0.7.2-2.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | 2.2.0 | `el9.aarch64` | pgdg | 43.3 KiB | [pg_stat_monitor_13-2.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_monitor_13-2.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_13` | 2.2.0 | `el9.x86_64` | pgdg | 44.4 KiB | [pg_stat_monitor_13-2.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_monitor_13-2.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_13` | 2.1.1 | `el9.x86_64` | pgdg | 44.0 KiB | [pg_stat_monitor_13-2.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_monitor_13-2.1.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_13` | 2.1.1 | `el9.aarch64` | pgdg | 43.3 KiB | [pg_stat_monitor_13-2.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_monitor_13-2.1.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_13` | 2.1.0 | `el9.x86_64` | pgdg | 42.7 KiB | [pg_stat_monitor_13-2.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_monitor_13-2.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_13` | 2.1.0 | `el9.aarch64` | pgdg | 41.8 KiB | [pg_stat_monitor_13-2.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_monitor_13-2.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_13` | 2.0.4 | `el9.x86_64` | pgdg | 41.4 KiB | [pg_stat_monitor_13-2.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_monitor_13-2.0.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_13` | 2.0.4 | `el9.aarch64` | pgdg | 40.5 KiB | [pg_stat_monitor_13-2.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_monitor_13-2.0.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_13` | 2.0.1 | `el9.aarch64` | pgdg | 40.2 KiB | [pg_stat_monitor_13-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_monitor_13-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_13` | 2.0.1 | `el9.x86_64` | pgdg | 41.1 KiB | [pg_stat_monitor_13-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_monitor_13-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_13` | 1.1.0 | `el9.aarch64` | pgdg | 96.1 KiB | [pg_stat_monitor_13-1.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_monitor_13-1.1.0-1.rhel9.aarch64.rpm) |
| `pg_stat_monitor_13` | 1.1.0 | `el9.x86_64` | pgdg | 97.4 KiB | [pg_stat_monitor_13-1.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_monitor_13-1.1.0-1.rhel9.x86_64.rpm) |
| `pg_stat_monitor_13` | 1.0.1 | `el9.x86_64` | pgdg | 117.1 KiB | [pg_stat_monitor_13-1.0.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_monitor_13-1.0.1-1.rhel9.x86_64.rpm) |
| `pg_stat_monitor_13` | 1.0.0 | `el9.x86_64` | pgdg | 116.5 KiB | [pg_stat_monitor_13-1.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_monitor_13-1.0.0-1.rhel9.x86_64.rpm) |
| `pg_stat_monitor_13` | 1.0.0 | `el9.x86_64` | pgdg | 102.8 KiB | [pg_stat_monitor_13-1.0.0-rc.1_1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_monitor_13-1.0.0-rc.1_1.rhel9.x86_64.rpm) |
| `postgresql-13-pg-stat-monitor` | 2.2.0 | `d12.aarch64` | pigsty | 83.2 KiB | [postgresql-13-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-13-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-stat-monitor` | 2.2.0 | `d12.x86_64` | pigsty | 85.1 KiB | [postgresql-13-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-13-pg-stat-monitor_2.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-stat-monitor` | 2.2.0 | `u22.x86_64` | pigsty | 98.8 KiB | [postgresql-13-pg-stat-monitor_2.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-13-pg-stat-monitor_2.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-stat-monitor` | 2.2.0 | `u22.aarch64` | pigsty | 97.5 KiB | [postgresql-13-pg-stat-monitor_2.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-13-pg-stat-monitor_2.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-stat-monitor` | 2.2.0 | `u24.x86_64` | pigsty | 89.4 KiB | [postgresql-13-pg-stat-monitor_2.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-13-pg-stat-monitor_2.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-stat-monitor` | 2.2.0 | `u24.aarch64` | pigsty | 88.1 KiB | [postgresql-13-pg-stat-monitor_2.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-13-pg-stat-monitor_2.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/percona/pg_stat_monitor" title="Repository" icon="github" subtitle="github.com/percona/pg_stat_monitor" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_stat_monitor-2.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_stat_monitor; # get pg_stat_monitor source code
pig build dep pg_stat_monitor; # install build dependencies
pig build pkg pg_stat_monitor; # build extension rpm or deb
pig build ext pg_stat_monitor; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_stat_monitor; # install by extension name, for the current active PG version
pig ext install pg_stat_monitor; # install via package alias, for the active PG version
pig ext install pg_stat_monitor -v 17;   # install for PG 17
pig ext install pg_stat_monitor -v 16;   # install for PG 16
pig ext install pg_stat_monitor -v 15;   # install for PG 15
pig ext install pg_stat_monitor -v 14;   # install for PG 14
pig ext install pg_stat_monitor -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_stat_monitor;
```

