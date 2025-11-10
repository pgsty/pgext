---
title: "pg_stat_monitor"
linkTitle: "pg_stat_monitor"
description: "The pg_stat_monitor is a PostgreSQL Query Performance Monitoring tool, based on PostgreSQL contrib module pg_stat_statements. pg_stat_monitor provides aggregated statistics, client information, plan details including plan, and histogram information."
weight: 6230
categories: ["STAT"]
width: full
---

[**pg_stat_monitor**](https://github.com/percona/pg_stat_monitor) : The pg_stat_monitor is a PostgreSQL Query Performance Monitoring tool, based on PostgreSQL contrib module pg_stat_statements. pg_stat_monitor provides aggregated statistics, client information, plan details including plan, and histogram information.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6230** | {{< badge content="pg_stat_monitor" link="https://github.com/percona/pg_stat_monitor" >}} | {{< ext "pg_stat_monitor" >}} | `2.2.0` | {{< category "STAT" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_show_plans" >}} {{< ext "pg_stat_kcache" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_store_plans" >}} {{< ext "pgsentinel" >}} {{< ext "auto_explain" >}} {{< ext "logerrors" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `2.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_stat_monitor` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.2.0` | {{< bg "18" "pg_stat_monitor_18*" "red" >}} {{< bg "17" "pg_stat_monitor_17*" "green" >}} {{< bg "16" "pg_stat_monitor_16*" "green" >}} {{< bg "15" "pg_stat_monitor_15*" "green" >}} {{< bg "14" "pg_stat_monitor_14*" "green" >}} {{< bg "13" "pg_stat_monitor_13*" "green" >}} | `pg_stat_monitor_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.3.0` | {{< bg "18" "postgresql-18-pg-stat-monitor" "green" >}} {{< bg "17" "postgresql-17-pg-stat-monitor" "green" >}} {{< bg "16" "postgresql-16-pg-stat-monitor" "green" >}} {{< bg "15" "postgresql-15-pg-stat-monitor" "green" >}} {{< bg "14" "postgresql-14-pg-stat-monitor" "green" >}} {{< bg "13" "postgresql-13-pg-stat-monitor" "green" >}} | `postgresql-$v-pg-stat-monitor` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pg_stat_monitor_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.0" "pg_stat_monitor_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_14 : AVAIL 10" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_13 : AVAIL 11" "blue" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pg_stat_monitor_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.0" "pg_stat_monitor_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_14 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_13 : AVAIL 6" "blue" >}} |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "pg_stat_monitor_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.0" "pg_stat_monitor_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_14 : AVAIL 9" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_13 : AVAIL 9" "blue" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "pg_stat_monitor_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.0" "pg_stat_monitor_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_14 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_13 : AVAIL 6" "blue" >}} |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "pg_stat_monitor_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.0" "pg_stat_monitor_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "pg_stat_monitor_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.0" "pg_stat_monitor_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2.0" "pg_stat_monitor_13 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-14-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-13-pg-stat-monitor : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-14-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-13-pg-stat-monitor : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-14-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-13-pg-stat-monitor : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-14-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-13-pg-stat-monitor : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-14-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-13-pg-stat-monitor : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-14-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-13-pg-stat-monitor : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-14-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-13-pg-stat-monitor : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-14-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-13-pg-stat-monitor : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-18-pg-stat-monitor` | `2.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 74.0 KiB | [postgresql-18-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-18-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-stat-monitor` | `2.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 72.4 KiB | [postgresql-18-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-18-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-stat-monitor` | `2.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 73.9 KiB | [postgresql-18-pg-stat-monitor_2.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-18-pg-stat-monitor_2.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-stat-monitor` | `2.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 72.6 KiB | [postgresql-18-pg-stat-monitor_2.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-18-pg-stat-monitor_2.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-stat-monitor` | `2.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 80.0 KiB | [postgresql-18-pg-stat-monitor_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-18-pg-stat-monitor_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-stat-monitor` | `2.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 78.9 KiB | [postgresql-18-pg-stat-monitor_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-18-pg-stat-monitor_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-stat-monitor` | `2.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 77.2 KiB | [postgresql-18-pg-stat-monitor_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-18-pg-stat-monitor_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-stat-monitor` | `2.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 76.1 KiB | [postgresql-18-pg-stat-monitor_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-18-pg-stat-monitor_2.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_monitor_17` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.2 KiB | [pg_stat_monitor_17-2.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_monitor_17-2.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.7 KiB | [pg_stat_monitor_17-2.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_monitor_17-2.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.2 KiB | [pg_stat_monitor_17-2.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_monitor_17-2.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_stat_monitor_17-2.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_monitor_17-2.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.1 KiB | [pg_stat_monitor_17-2.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_monitor_17-2.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.6 KiB | [pg_stat_monitor_17-2.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_monitor_17-2.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.2 KiB | [pg_stat_monitor_17-2.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_stat_monitor_17-2.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [pg_stat_monitor_17-2.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_stat_monitor_17-2.1.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.5 KiB | [pg_stat_monitor_17-2.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_stat_monitor_17-2.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.1 KiB | [pg_stat_monitor_17-2.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_stat_monitor_17-2.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.1.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.9 KiB | [pg_stat_monitor_17-2.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_stat_monitor_17-2.1.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.5 KiB | [pg_stat_monitor_17-2.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_stat_monitor_17-2.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.5 KiB | [pg_stat_monitor_17-2.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_stat_monitor_17-2.2.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.1.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.4 KiB | [pg_stat_monitor_17-2.1.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_stat_monitor_17-2.1.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.1 KiB | [pg_stat_monitor_17-2.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_monitor_17-2.2.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.1.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.8 KiB | [pg_stat_monitor_17-2.1.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_monitor_17-2.1.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-stat-monitor` | `2.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 73.8 KiB | [postgresql-17-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-stat-monitor` | `2.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 72.1 KiB | [postgresql-17-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-stat-monitor` | `2.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 73.8 KiB | [postgresql-17-pg-stat-monitor_2.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-stat-monitor` | `2.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 72.4 KiB | [postgresql-17-pg-stat-monitor_2.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-stat-monitor` | `2.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 85.8 KiB | [postgresql-17-pg-stat-monitor_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-stat-monitor` | `2.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 84.4 KiB | [postgresql-17-pg-stat-monitor_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-stat-monitor` | `2.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 76.9 KiB | [postgresql-17-pg-stat-monitor_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-stat-monitor` | `2.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 75.8 KiB | [postgresql-17-pg-stat-monitor_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_monitor_16` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.1 KiB | [pg_stat_monitor_16-2.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.6 KiB | [pg_stat_monitor_16-2.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.0 KiB | [pg_stat_monitor_16-2.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.7 KiB | [pg_stat_monitor_16-2.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.0.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.5 KiB | [pg_stat_monitor_16-2.0.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.0.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_stat_monitor_16-2.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.1 KiB | [pg_stat_monitor_16-2.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.6 KiB | [pg_stat_monitor_16-2.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.3 KiB | [pg_stat_monitor_16-2.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.0.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.1 KiB | [pg_stat_monitor_16-2.0.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.0.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.0 KiB | [pg_stat_monitor_16-2.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_monitor_16-2.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.7 KiB | [pg_stat_monitor_16-2.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_monitor_16-2.1.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.3 KiB | [pg_stat_monitor_16-2.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_monitor_16-2.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 38.1 KiB | [pg_stat_monitor_16-2.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_monitor_16-2.0.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 38.0 KiB | [pg_stat_monitor_16-2.0.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_monitor_16-2.0.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.0 KiB | [pg_stat_monitor_16-2.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_monitor_16-2.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.1.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.8 KiB | [pg_stat_monitor_16-2.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_monitor_16-2.1.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.3 KiB | [pg_stat_monitor_16-2.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_monitor_16-2.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 37.2 KiB | [pg_stat_monitor_16-2.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_monitor_16-2.0.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.9 KiB | [pg_stat_monitor_16-2.0.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_monitor_16-2.0.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.4 KiB | [pg_stat_monitor_16-2.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_stat_monitor_16-2.2.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.1.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.3 KiB | [pg_stat_monitor_16-2.1.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_stat_monitor_16-2.1.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.0 KiB | [pg_stat_monitor_16-2.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_monitor_16-2.2.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.1.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_stat_monitor_16-2.1.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_monitor_16-2.1.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-stat-monitor` | `2.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 73.4 KiB | [postgresql-16-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-stat-monitor` | `2.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 71.7 KiB | [postgresql-16-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-stat-monitor` | `2.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 73.5 KiB | [postgresql-16-pg-stat-monitor_2.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-stat-monitor` | `2.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 72.1 KiB | [postgresql-16-pg-stat-monitor_2.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-stat-monitor` | `2.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 85.4 KiB | [postgresql-16-pg-stat-monitor_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-stat-monitor` | `2.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 84.0 KiB | [postgresql-16-pg-stat-monitor_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-stat-monitor` | `2.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 76.5 KiB | [postgresql-16-pg-stat-monitor_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-stat-monitor` | `2.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 75.4 KiB | [postgresql-16-pg-stat-monitor_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_monitor_15` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.3 KiB | [pg_stat_monitor_15-2.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.8 KiB | [pg_stat_monitor_15-2.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.2 KiB | [pg_stat_monitor_15-2.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.9 KiB | [pg_stat_monitor_15-2.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.0.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.5 KiB | [pg_stat_monitor_15-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.1 KiB | [pg_stat_monitor_15-1.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-1.1.0-1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.5 KiB | [pg_stat_monitor_15-2.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.0 KiB | [pg_stat_monitor_15-2.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.5 KiB | [pg_stat_monitor_15-2.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.1 KiB | [pg_stat_monitor_15-2.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.0.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.8 KiB | [pg_stat_monitor_15-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 84.7 KiB | [pg_stat_monitor_15-1.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-1.1.0-1.rhel8.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.6 KiB | [pg_stat_monitor_15-2.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_monitor_15-2.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.5 KiB | [pg_stat_monitor_15-2.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_monitor_15-2.1.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.1 KiB | [pg_stat_monitor_15-2.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_monitor_15-2.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.0 KiB | [pg_stat_monitor_15-2.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_monitor_15-2.0.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.7 KiB | [pg_stat_monitor_15-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_monitor_15-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 88.5 KiB | [pg_stat_monitor_15-1.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_monitor_15-1.1.0-1.rhel9.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.7 KiB | [pg_stat_monitor_15-2.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_monitor_15-2.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.1.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.4 KiB | [pg_stat_monitor_15-2.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_monitor_15-2.1.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.1 KiB | [pg_stat_monitor_15-2.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_monitor_15-2.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.9 KiB | [pg_stat_monitor_15-2.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_monitor_15-2.0.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.6 KiB | [pg_stat_monitor_15-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_monitor_15-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 87.1 KiB | [pg_stat_monitor_15-1.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_monitor_15-1.1.0-1.rhel9.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.0 KiB | [pg_stat_monitor_15-2.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_stat_monitor_15-2.2.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.1.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.9 KiB | [pg_stat_monitor_15-2.1.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_stat_monitor_15-2.1.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.6 KiB | [pg_stat_monitor_15-2.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_monitor_15-2.2.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.1.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.3 KiB | [pg_stat_monitor_15-2.1.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_monitor_15-2.1.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-stat-monitor` | `2.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 74.8 KiB | [postgresql-15-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-stat-monitor` | `2.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 73.0 KiB | [postgresql-15-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-stat-monitor` | `2.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 74.8 KiB | [postgresql-15-pg-stat-monitor_2.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-stat-monitor` | `2.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 73.2 KiB | [postgresql-15-pg-stat-monitor_2.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-stat-monitor` | `2.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 87.4 KiB | [postgresql-15-pg-stat-monitor_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-stat-monitor` | `2.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 86.1 KiB | [postgresql-15-pg-stat-monitor_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-stat-monitor` | `2.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 78.2 KiB | [postgresql-15-pg-stat-monitor_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-stat-monitor` | `2.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 77.3 KiB | [postgresql-15-pg-stat-monitor_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_monitor_14` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.1 KiB | [pg_stat_monitor_14-2.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-2.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | `2.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.6 KiB | [pg_stat_monitor_14-2.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-2.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.1 KiB | [pg_stat_monitor_14-2.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-2.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.8 KiB | [pg_stat_monitor_14-2.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-2.0.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.4 KiB | [pg_stat_monitor_14-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.0 KiB | [pg_stat_monitor_14-1.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-1.1.0-1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 106.7 KiB | [pg_stat_monitor_14-1.0.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-1.0.1-1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 93.4 KiB | [pg_stat_monitor_14-1.0.0-rc.1_1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-1.0.0-rc.1_1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.0 KiB | [pg_stat_monitor_14-1.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-1.0.0-1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | `0.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 86.9 KiB | [pg_stat_monitor_14-0.9.2-beta1_1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_monitor_14-0.9.2-beta1_1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_14` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.4 KiB | [pg_stat_monitor_14-2.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_monitor_14-2.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_14` | `2.1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.9 KiB | [pg_stat_monitor_14-2.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_monitor_14-2.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_14` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.4 KiB | [pg_stat_monitor_14-2.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_monitor_14-2.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_14` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.0 KiB | [pg_stat_monitor_14-2.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_monitor_14-2.0.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_14` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.7 KiB | [pg_stat_monitor_14-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_monitor_14-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_14` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 84.8 KiB | [pg_stat_monitor_14-1.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_monitor_14-1.1.0-1.rhel8.aarch64.rpm) |
| `pg_stat_monitor_14` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.4 KiB | [pg_stat_monitor_14-2.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_monitor_14-2.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_14` | `2.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.2 KiB | [pg_stat_monitor_14-2.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_monitor_14-2.1.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_14` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.8 KiB | [pg_stat_monitor_14-2.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_monitor_14-2.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_14` | `2.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.8 KiB | [pg_stat_monitor_14-2.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_monitor_14-2.0.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_14` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.5 KiB | [pg_stat_monitor_14-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_monitor_14-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_14` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 88.5 KiB | [pg_stat_monitor_14-1.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_monitor_14-1.1.0-1.rhel9.x86_64.rpm) |
| `pg_stat_monitor_14` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.1 KiB | [pg_stat_monitor_14-1.0.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_monitor_14-1.0.1-1.rhel9.x86_64.rpm) |
| `pg_stat_monitor_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 94.1 KiB | [pg_stat_monitor_14-1.0.0-rc.1_1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_monitor_14-1.0.0-rc.1_1.rhel9.x86_64.rpm) |
| `pg_stat_monitor_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.5 KiB | [pg_stat_monitor_14-1.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_monitor_14-1.0.0-1.rhel9.x86_64.rpm) |
| `pg_stat_monitor_14` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.6 KiB | [pg_stat_monitor_14-2.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_monitor_14-2.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_14` | `2.1.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.2 KiB | [pg_stat_monitor_14-2.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_monitor_14-2.1.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_14` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.0 KiB | [pg_stat_monitor_14-2.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_monitor_14-2.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_14` | `2.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.8 KiB | [pg_stat_monitor_14-2.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_monitor_14-2.0.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_14` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.5 KiB | [pg_stat_monitor_14-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_monitor_14-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_14` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 87.2 KiB | [pg_stat_monitor_14-1.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_monitor_14-1.1.0-1.rhel9.aarch64.rpm) |
| `pg_stat_monitor_14` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.9 KiB | [pg_stat_monitor_14-2.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_stat_monitor_14-2.2.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_stat_monitor_14` | `2.1.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.7 KiB | [pg_stat_monitor_14-2.1.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_stat_monitor_14-2.1.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_stat_monitor_14` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.5 KiB | [pg_stat_monitor_14-2.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_stat_monitor_14-2.2.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_stat_monitor_14` | `2.1.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.2 KiB | [pg_stat_monitor_14-2.1.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_stat_monitor_14-2.1.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-stat-monitor` | `2.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 74.3 KiB | [postgresql-14-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-stat-monitor` | `2.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 72.7 KiB | [postgresql-14-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-stat-monitor` | `2.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 74.3 KiB | [postgresql-14-pg-stat-monitor_2.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-stat-monitor` | `2.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 73.0 KiB | [postgresql-14-pg-stat-monitor_2.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-stat-monitor` | `2.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 86.8 KiB | [postgresql-14-pg-stat-monitor_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-stat-monitor` | `2.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 85.5 KiB | [postgresql-14-pg-stat-monitor_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-stat-monitor` | `2.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 77.5 KiB | [postgresql-14-pg-stat-monitor_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-stat-monitor` | `2.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 76.8 KiB | [postgresql-14-pg-stat-monitor_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-14-pg-stat-monitor_2.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_monitor_13` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.9 KiB | [pg_stat_monitor_13-2.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-2.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | `2.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.4 KiB | [pg_stat_monitor_13-2.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-2.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.9 KiB | [pg_stat_monitor_13-2.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-2.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.4 KiB | [pg_stat_monitor_13-2.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-2.0.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.1 KiB | [pg_stat_monitor_13-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 95.8 KiB | [pg_stat_monitor_13-1.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-1.1.0-1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 115.5 KiB | [pg_stat_monitor_13-1.0.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-1.0.1-1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.8 KiB | [pg_stat_monitor_13-1.0.0-rc.1_1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-1.0.0-rc.1_1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 117.3 KiB | [pg_stat_monitor_13-1.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-1.0.0-1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | `0.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 94.2 KiB | [pg_stat_monitor_13-0.9.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-0.9.1-1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | `0.7.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 86.6 KiB | [pg_stat_monitor_13-0.7.2-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_monitor_13-0.7.2-2.rhel8.x86_64.rpm) |
| `pg_stat_monitor_13` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.1 KiB | [pg_stat_monitor_13-2.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_monitor_13-2.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_13` | `2.1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.6 KiB | [pg_stat_monitor_13-2.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_monitor_13-2.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_13` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.1 KiB | [pg_stat_monitor_13-2.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_monitor_13-2.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_13` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.7 KiB | [pg_stat_monitor_13-2.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_monitor_13-2.0.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_13` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.3 KiB | [pg_stat_monitor_13-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_monitor_13-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_13` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 93.5 KiB | [pg_stat_monitor_13-1.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_monitor_13-1.1.0-1.rhel8.aarch64.rpm) |
| `pg_stat_monitor_13` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.4 KiB | [pg_stat_monitor_13-2.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_monitor_13-2.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_13` | `2.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.0 KiB | [pg_stat_monitor_13-2.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_monitor_13-2.1.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_13` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.7 KiB | [pg_stat_monitor_13-2.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_monitor_13-2.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_13` | `2.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.4 KiB | [pg_stat_monitor_13-2.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_monitor_13-2.0.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_13` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.1 KiB | [pg_stat_monitor_13-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_monitor_13-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_monitor_13` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 97.4 KiB | [pg_stat_monitor_13-1.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_monitor_13-1.1.0-1.rhel9.x86_64.rpm) |
| `pg_stat_monitor_13` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 117.1 KiB | [pg_stat_monitor_13-1.0.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_monitor_13-1.0.1-1.rhel9.x86_64.rpm) |
| `pg_stat_monitor_13` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 116.5 KiB | [pg_stat_monitor_13-1.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_monitor_13-1.0.0-1.rhel9.x86_64.rpm) |
| `pg_stat_monitor_13` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 102.8 KiB | [pg_stat_monitor_13-1.0.0-rc.1_1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_monitor_13-1.0.0-rc.1_1.rhel9.x86_64.rpm) |
| `pg_stat_monitor_13` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.3 KiB | [pg_stat_monitor_13-2.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_monitor_13-2.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_13` | `2.1.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.3 KiB | [pg_stat_monitor_13-2.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_monitor_13-2.1.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_13` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.8 KiB | [pg_stat_monitor_13-2.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_monitor_13-2.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_13` | `2.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.5 KiB | [pg_stat_monitor_13-2.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_monitor_13-2.0.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_13` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.2 KiB | [pg_stat_monitor_13-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_monitor_13-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_monitor_13` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 96.1 KiB | [pg_stat_monitor_13-1.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_monitor_13-1.1.0-1.rhel9.aarch64.rpm) |
| `pg_stat_monitor_13` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.0 KiB | [pg_stat_monitor_13-2.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_stat_monitor_13-2.2.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_stat_monitor_13` | `2.1.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.7 KiB | [pg_stat_monitor_13-2.1.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_stat_monitor_13-2.1.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_stat_monitor_13` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.2 KiB | [pg_stat_monitor_13-2.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_stat_monitor_13-2.2.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_stat_monitor_13` | `2.1.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.0 KiB | [pg_stat_monitor_13-2.1.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_stat_monitor_13-2.1.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pg-stat-monitor` | `2.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 85.9 KiB | [postgresql-13-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-13-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-stat-monitor` | `2.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 83.8 KiB | [postgresql-13-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-13-pg-stat-monitor_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-stat-monitor` | `2.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 85.9 KiB | [postgresql-13-pg-stat-monitor_2.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-13-pg-stat-monitor_2.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-stat-monitor` | `2.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 84.2 KiB | [postgresql-13-pg-stat-monitor_2.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-13-pg-stat-monitor_2.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-stat-monitor` | `2.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 99.5 KiB | [postgresql-13-pg-stat-monitor_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-13-pg-stat-monitor_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-stat-monitor` | `2.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 98.5 KiB | [postgresql-13-pg-stat-monitor_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-13-pg-stat-monitor_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-stat-monitor` | `2.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 89.9 KiB | [postgresql-13-pg-stat-monitor_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-13-pg-stat-monitor_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-stat-monitor` | `2.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 88.8 KiB | [postgresql-13-pg-stat-monitor_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-13-pg-stat-monitor_2.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/percona/pg_stat_monitor" title="Repository" icon="github" subtitle="github.com/percona/pg_stat_monitor" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_stat_monitor-2.3.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_stat_monitor;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_stat_monitor;		# install via package name, for the active PG version

pig install pg_stat_monitor -v 18;   # install for PG 18
pig install pg_stat_monitor -v 17;   # install for PG 17
pig install pg_stat_monitor -v 16;   # install for PG 16
pig install pg_stat_monitor -v 15;   # install for PG 15
pig install pg_stat_monitor -v 14;   # install for PG 14
pig install pg_stat_monitor -v 13;   # install for PG 13

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pg_stat_monitor';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_stat_monitor;
```
