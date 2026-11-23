---
title: "pg_stat_monitor"
linkTitle: "pg_stat_monitor"
description: "The pg_stat_monitor is a PostgreSQL Query Performance Monitoring tool, based on PostgreSQL contrib module pg_stat_statements. pg_stat_monitor provides aggregated statistics, client information, plan details including plan, and histogram information."
weight: 6240
categories: ["STAT"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_stat_monitor**](https://github.com/percona/pg_stat_monitor) : The pg_stat_monitor is a PostgreSQL Query Performance Monitoring tool, based on PostgreSQL contrib module pg_stat_statements. pg_stat_monitor provides aggregated statistics, client information, plan details including plan, and histogram information.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6240** | {{< badge content="pg_stat_monitor" link="https://github.com/percona/pg_stat_monitor" >}} | {{< ext "pg_stat_monitor" >}} | `2.4.1` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_profile" >}} {{< ext "pg_wait_sampling" >}} {{< ext "pgsentinel" >}} {{< ext "pg_datasentinel" >}} {{< ext "powa" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_store_plans" >}} {{< ext "pg_stat_plans" >}} {{< ext "pg_track_settings" >}} {{< ext "pg_track_optimizer" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.4.1` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pg_stat_monitor` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.4.1` | {{< bg "19" "pg_stat_monitor_19" "green" >}} {{< bg "18" "pg_stat_monitor_18" "green" >}} {{< bg "17" "pg_stat_monitor_17" "green" >}} {{< bg "16" "pg_stat_monitor_16" "green" >}} {{< bg "15" "pg_stat_monitor_15" "green" >}} | `pg_stat_monitor_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.4.1` | {{< bg "19" "postgresql-19-pg-stat-monitor" "green" >}} {{< bg "18" "postgresql-18-pg-stat-monitor" "green" >}} {{< bg "17" "postgresql-17-pg-stat-monitor" "green" >}} {{< bg "16" "postgresql-16-pg-stat-monitor" "green" >}} {{< bg "15" "postgresql-15-pg-stat-monitor" "green" >}} | `postgresql-$v-pg-stat-monitor` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_stat_monitor_19 : N/A 0" "gray" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_17 : AVAIL 9" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_16 : AVAIL 11" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_15 : AVAIL 12" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_stat_monitor_19 : N/A 0" "gray" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_17 : AVAIL 9" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_16 : AVAIL 11" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_15 : AVAIL 12" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_15 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_15 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_15 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_18 : AVAIL 11" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_17 : AVAIL 13" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_16 : AVAIL 13" "blue" >}} | {{< bg "PGDG 2.4.1" "pg_stat_monitor_15 : AVAIL 13" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-19-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-18-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-17-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-16-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-15-pg-stat-monitor : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-19-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-18-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-17-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-16-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-15-pg-stat-monitor : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-19-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-18-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-17-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-16-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-15-pg-stat-monitor : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-19-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-18-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-17-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-16-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-15-pg-stat-monitor : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-19-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-18-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-17-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-16-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-15-pg-stat-monitor : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-19-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-18-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-17-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-16-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-15-pg-stat-monitor : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-19-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-18-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-17-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-16-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-15-pg-stat-monitor : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-19-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-18-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-17-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-16-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-15-pg-stat-monitor : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-19-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-18-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-17-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-16-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-15-pg-stat-monitor : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-19-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-18-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-17-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-16-pg-stat-monitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.1" "postgresql-15-pg-stat-monitor : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_monitor_19` | `2.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.3 KiB | [pg_stat_monitor_19-2.4.1-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-x86_64/pg_stat_monitor_19-2.4.1-1PGDG.rhel9.9.x86_64.rpm) |
| `pg_stat_monitor_19` | `2.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.5 KiB | [pg_stat_monitor_19-2.4.1-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-aarch64/pg_stat_monitor_19-2.4.1-1PGDG.rhel9.9.aarch64.rpm) |
| `pg_stat_monitor_19` | `2.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.7 KiB | [pg_stat_monitor_19-2.4.1-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-x86_64/pg_stat_monitor_19-2.4.1-1PGDG.rhel10.3.x86_64.rpm) |
| `pg_stat_monitor_19` | `2.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.9 KiB | [pg_stat_monitor_19-2.4.1-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-aarch64/pg_stat_monitor_19-2.4.1-1PGDG.rhel10.3.aarch64.rpm) |
| `postgresql-19-pg-stat-monitor` | `2.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 73.9 KiB | [postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-19-pg-stat-monitor` | `2.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 71.0 KiB | [postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-19-pg-stat-monitor` | `2.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 74.1 KiB | [postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-19-pg-stat-monitor` | `2.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 71.5 KiB | [postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-19-pg-stat-monitor` | `2.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 78.2 KiB | [postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-19-pg-stat-monitor` | `2.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 75.9 KiB | [postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-19-pg-stat-monitor` | `2.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 75.1 KiB | [postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~noble_amd64.deb) |
| `postgresql-19-pg-stat-monitor` | `2.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 73.3 KiB | [postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~noble_arm64.deb) |
| `postgresql-19-pg-stat-monitor` | `2.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 74.5 KiB | [postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-monitor/postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-19-pg-stat-monitor` | `2.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 72.5 KiB | [postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-monitor/postgresql-19-pg-stat-monitor_2.4.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_monitor_18` | `2.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.7 KiB | [pg_stat_monitor_18-2.4.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_stat_monitor_18-2.4.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_18` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.5 KiB | [pg_stat_monitor_18-2.4.0-4PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_stat_monitor_18-2.4.0-4PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_18` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 82.2 KiB | [pg_stat_monitor_18-2.4.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_stat_monitor_18-2.4.0-1PGSTY.el8.x86_64.rpm) |
| `pg_stat_monitor_18` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.0 KiB | [pg_stat_monitor_18-2.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_stat_monitor_18-2.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_18` | `2.3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.7 KiB | [pg_stat_monitor_18-2.3.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_stat_monitor_18-2.3.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_18` | `2.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.6 KiB | [pg_stat_monitor_18-2.3.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_stat_monitor_18-2.3.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_18` | `2.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.2 KiB | [pg_stat_monitor_18-2.4.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_stat_monitor_18-2.4.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_18` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.1 KiB | [pg_stat_monitor_18-2.4.0-4PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_stat_monitor_18-2.4.0-4PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_18` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 80.9 KiB | [pg_stat_monitor_18-2.4.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_stat_monitor_18-2.4.0-1PGSTY.el8.aarch64.rpm) |
| `pg_stat_monitor_18` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.6 KiB | [pg_stat_monitor_18-2.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_stat_monitor_18-2.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_18` | `2.3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.3 KiB | [pg_stat_monitor_18-2.3.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_stat_monitor_18-2.3.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_18` | `2.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.2 KiB | [pg_stat_monitor_18-2.3.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_stat_monitor_18-2.3.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_18` | `2.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.6 KiB | [pg_stat_monitor_18-2.4.1-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_stat_monitor_18-2.4.1-1PGDG.rhel9.9.x86_64.rpm) |
| `pg_stat_monitor_18` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.5 KiB | [pg_stat_monitor_18-2.4.0-4PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_stat_monitor_18-2.4.0-4PGDG.rhel9.9.x86_64.rpm) |
| `pg_stat_monitor_18` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 83.5 KiB | [pg_stat_monitor_18-2.4.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_monitor_18-2.4.0-1PGSTY.el9.x86_64.rpm) |
| `pg_stat_monitor_18` | `2.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.5 KiB | [pg_stat_monitor_18-2.4.1-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_stat_monitor_18-2.4.1-1PGDG.rhel9.9.aarch64.rpm) |
| `pg_stat_monitor_18` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.4 KiB | [pg_stat_monitor_18-2.4.0-4PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_stat_monitor_18-2.4.0-4PGDG.rhel9.9.aarch64.rpm) |
| `pg_stat_monitor_18` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 82.4 KiB | [pg_stat_monitor_18-2.4.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_monitor_18-2.4.0-1PGSTY.el9.aarch64.rpm) |
| `pg_stat_monitor_18` | `2.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.2 KiB | [pg_stat_monitor_18-2.4.1-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_stat_monitor_18-2.4.1-1PGDG.rhel10.3.x86_64.rpm) |
| `pg_stat_monitor_18` | `2.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.9 KiB | [pg_stat_monitor_18-2.4.0-4PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_stat_monitor_18-2.4.0-4PGDG.rhel10.3.x86_64.rpm) |
| `pg_stat_monitor_18` | `2.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 83.9 KiB | [pg_stat_monitor_18-2.4.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_monitor_18-2.4.0-1PGSTY.el10.x86_64.rpm) |
| `pg_stat_monitor_18` | `2.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.1 KiB | [pg_stat_monitor_18-2.4.1-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_monitor_18-2.4.1-1PGDG.rhel10.3.aarch64.rpm) |
| `pg_stat_monitor_18` | `2.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.1 KiB | [pg_stat_monitor_18-2.4.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_monitor_18-2.4.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_monitor_18` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.0 KiB | [pg_stat_monitor_18-2.4.0-4PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_monitor_18-2.4.0-4PGDG.rhel10.3.aarch64.rpm) |
| `pg_stat_monitor_18` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.0 KiB | [pg_stat_monitor_18-2.4.0-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_monitor_18-2.4.0-4PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_monitor_18` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 83.0 KiB | [pg_stat_monitor_18-2.4.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_monitor_18-2.4.0-1PGSTY.el10.aarch64.rpm) |
| `pg_stat_monitor_18` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.3 KiB | [pg_stat_monitor_18-2.4.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_monitor_18-2.4.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_monitor_18` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.4 KiB | [pg_stat_monitor_18-2.3.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_monitor_18-2.3.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_monitor_18` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.4 KiB | [pg_stat_monitor_18-2.3.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_monitor_18-2.3.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_stat_monitor_18` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.4 KiB | [pg_stat_monitor_18-2.3.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_monitor_18-2.3.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_stat_monitor_18` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.1 KiB | [pg_stat_monitor_18-2.3.1-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_monitor_18-2.3.1-1PGDGrhel10.1.aarch64.rpm) |
| `pg_stat_monitor_18` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.2 KiB | [pg_stat_monitor_18-2.3.1-1PGDGrhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_monitor_18-2.3.1-1PGDGrhel10.0.aarch64.rpm) |
| `postgresql-18-pg-stat-monitor` | `2.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 73.0 KiB | [postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-stat-monitor` | `2.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 71.3 KiB | [postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-stat-monitor` | `2.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 73.1 KiB | [postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-stat-monitor` | `2.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 72.1 KiB | [postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-stat-monitor` | `2.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 77.0 KiB | [postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-stat-monitor` | `2.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 76.1 KiB | [postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-stat-monitor` | `2.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 74.2 KiB | [postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-stat-monitor` | `2.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 73.8 KiB | [postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-stat-monitor` | `2.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 73.5 KiB | [postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-monitor/postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-stat-monitor` | `2.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 73.1 KiB | [postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-monitor/postgresql-18-pg-stat-monitor_2.4.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_monitor_17` | `2.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.6 KiB | [pg_stat_monitor_17-2.4.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_monitor_17-2.4.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.4 KiB | [pg_stat_monitor_17-2.4.0-4PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_monitor_17-2.4.0-4PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 81.7 KiB | [pg_stat_monitor_17-2.4.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_stat_monitor_17-2.4.0-1PGSTY.el8.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.9 KiB | [pg_stat_monitor_17-2.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_monitor_17-2.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.7 KiB | [pg_stat_monitor_17-2.3.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_monitor_17-2.3.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.5 KiB | [pg_stat_monitor_17-2.3.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_monitor_17-2.3.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.2 KiB | [pg_stat_monitor_17-2.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_monitor_17-2.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.7 KiB | [pg_stat_monitor_17-2.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_monitor_17-2.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.2 KiB | [pg_stat_monitor_17-2.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_monitor_17-2.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.0 KiB | [pg_stat_monitor_17-2.4.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_monitor_17-2.4.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.9 KiB | [pg_stat_monitor_17-2.4.0-4PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_monitor_17-2.4.0-4PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 80.4 KiB | [pg_stat_monitor_17-2.4.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_stat_monitor_17-2.4.0-1PGSTY.el8.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.3 KiB | [pg_stat_monitor_17-2.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_monitor_17-2.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.3 KiB | [pg_stat_monitor_17-2.3.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_monitor_17-2.3.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.1 KiB | [pg_stat_monitor_17-2.3.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_monitor_17-2.3.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_stat_monitor_17-2.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_monitor_17-2.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.1 KiB | [pg_stat_monitor_17-2.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_monitor_17-2.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.6 KiB | [pg_stat_monitor_17-2.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_monitor_17-2.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.4 KiB | [pg_stat_monitor_17-2.4.1-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_stat_monitor_17-2.4.1-1PGDG.rhel9.9.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.3 KiB | [pg_stat_monitor_17-2.4.0-4PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_stat_monitor_17-2.4.0-4PGDG.rhel9.9.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 82.8 KiB | [pg_stat_monitor_17-2.4.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_monitor_17-2.4.0-1PGSTY.el9.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.4 KiB | [pg_stat_monitor_17-2.4.1-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_stat_monitor_17-2.4.1-1PGDG.rhel9.9.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.2 KiB | [pg_stat_monitor_17-2.4.0-4PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_stat_monitor_17-2.4.0-4PGDG.rhel9.9.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 81.9 KiB | [pg_stat_monitor_17-2.4.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_monitor_17-2.4.0-1PGSTY.el9.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.0 KiB | [pg_stat_monitor_17-2.4.1-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_stat_monitor_17-2.4.1-1PGDG.rhel10.3.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.8 KiB | [pg_stat_monitor_17-2.4.0-4PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_stat_monitor_17-2.4.0-4PGDG.rhel10.3.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 83.4 KiB | [pg_stat_monitor_17-2.4.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_monitor_17-2.4.0-1PGSTY.el10.x86_64.rpm) |
| `pg_stat_monitor_17` | `2.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.0 KiB | [pg_stat_monitor_17-2.4.1-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_monitor_17-2.4.1-1PGDG.rhel10.3.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.0 KiB | [pg_stat_monitor_17-2.4.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_monitor_17-2.4.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.2 KiB | [pg_stat_monitor_17-2.4.0-4PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_monitor_17-2.4.0-4PGDG.rhel10.3.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.2 KiB | [pg_stat_monitor_17-2.4.0-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_monitor_17-2.4.0-4PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 82.8 KiB | [pg_stat_monitor_17-2.4.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_monitor_17-2.4.0-1PGSTY.el10.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.5 KiB | [pg_stat_monitor_17-2.4.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_monitor_17-2.4.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.2 KiB | [pg_stat_monitor_17-2.3.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_monitor_17-2.3.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.3 KiB | [pg_stat_monitor_17-2.3.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_monitor_17-2.3.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.2 KiB | [pg_stat_monitor_17-2.3.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_monitor_17-2.3.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.9 KiB | [pg_stat_monitor_17-2.3.1-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_monitor_17-2.3.1-1PGDGrhel10.1.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.0 KiB | [pg_stat_monitor_17-2.3.1-1PGDGrhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_monitor_17-2.3.1-1PGDGrhel10.0.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.1 KiB | [pg_stat_monitor_17-2.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_monitor_17-2.2.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_stat_monitor_17` | `2.1.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.8 KiB | [pg_stat_monitor_17-2.1.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_monitor_17-2.1.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-stat-monitor` | `2.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 72.5 KiB | [postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-stat-monitor` | `2.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 70.7 KiB | [postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-stat-monitor` | `2.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 72.6 KiB | [postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-stat-monitor` | `2.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 71.5 KiB | [postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-stat-monitor` | `2.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 82.7 KiB | [postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-stat-monitor` | `2.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 81.7 KiB | [postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-stat-monitor` | `2.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 73.6 KiB | [postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-stat-monitor` | `2.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 73.4 KiB | [postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-stat-monitor` | `2.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 73.1 KiB | [postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-stat-monitor` | `2.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 72.4 KiB | [postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-monitor/postgresql-17-pg-stat-monitor_2.4.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_monitor_16` | `2.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.4 KiB | [pg_stat_monitor_16-2.4.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.4.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.2 KiB | [pg_stat_monitor_16-2.4.0-4PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.4.0-4PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 81.3 KiB | [pg_stat_monitor_16-2.4.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_stat_monitor_16-2.4.0-1PGSTY.el8.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.7 KiB | [pg_stat_monitor_16-2.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.5 KiB | [pg_stat_monitor_16-2.3.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.3.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.4 KiB | [pg_stat_monitor_16-2.3.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.3.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.1 KiB | [pg_stat_monitor_16-2.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.6 KiB | [pg_stat_monitor_16-2.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.0 KiB | [pg_stat_monitor_16-2.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.7 KiB | [pg_stat_monitor_16-2.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.0.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.5 KiB | [pg_stat_monitor_16-2.0.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_monitor_16-2.0.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.9 KiB | [pg_stat_monitor_16-2.4.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.4.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.8 KiB | [pg_stat_monitor_16-2.4.0-4PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.4.0-4PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 80.3 KiB | [pg_stat_monitor_16-2.4.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_stat_monitor_16-2.4.0-1PGSTY.el8.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.2 KiB | [pg_stat_monitor_16-2.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.2 KiB | [pg_stat_monitor_16-2.3.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.3.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.1 KiB | [pg_stat_monitor_16-2.3.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.3.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_stat_monitor_16-2.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.1 KiB | [pg_stat_monitor_16-2.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.6 KiB | [pg_stat_monitor_16-2.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.3 KiB | [pg_stat_monitor_16-2.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.0.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.1 KiB | [pg_stat_monitor_16-2.0.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_monitor_16-2.0.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.2 KiB | [pg_stat_monitor_16-2.4.1-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_monitor_16-2.4.1-1PGDG.rhel9.9.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.1 KiB | [pg_stat_monitor_16-2.4.0-4PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_monitor_16-2.4.0-4PGDG.rhel9.9.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 82.5 KiB | [pg_stat_monitor_16-2.4.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_monitor_16-2.4.0-1PGSTY.el9.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.3 KiB | [pg_stat_monitor_16-2.4.1-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_monitor_16-2.4.1-1PGDG.rhel9.9.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.2 KiB | [pg_stat_monitor_16-2.4.0-4PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_monitor_16-2.4.0-4PGDG.rhel9.9.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 81.7 KiB | [pg_stat_monitor_16-2.4.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_monitor_16-2.4.0-1PGSTY.el9.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.8 KiB | [pg_stat_monitor_16-2.4.1-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_stat_monitor_16-2.4.1-1PGDG.rhel10.3.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.7 KiB | [pg_stat_monitor_16-2.4.0-4PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_stat_monitor_16-2.4.0-4PGDG.rhel10.3.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 83.1 KiB | [pg_stat_monitor_16-2.4.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_monitor_16-2.4.0-1PGSTY.el10.x86_64.rpm) |
| `pg_stat_monitor_16` | `2.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pg_stat_monitor_16-2.4.1-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_monitor_16-2.4.1-1PGDG.rhel10.3.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pg_stat_monitor_16-2.4.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_monitor_16-2.4.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.0 KiB | [pg_stat_monitor_16-2.4.0-4PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_monitor_16-2.4.0-4PGDG.rhel10.3.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.0 KiB | [pg_stat_monitor_16-2.4.0-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_monitor_16-2.4.0-4PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 82.6 KiB | [pg_stat_monitor_16-2.4.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_monitor_16-2.4.0-1PGSTY.el10.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.3 KiB | [pg_stat_monitor_16-2.4.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_monitor_16-2.4.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.1 KiB | [pg_stat_monitor_16-2.3.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_monitor_16-2.3.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.1 KiB | [pg_stat_monitor_16-2.3.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_monitor_16-2.3.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.1 KiB | [pg_stat_monitor_16-2.3.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_monitor_16-2.3.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.8 KiB | [pg_stat_monitor_16-2.3.1-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_monitor_16-2.3.1-1PGDGrhel10.1.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.8 KiB | [pg_stat_monitor_16-2.3.1-1PGDGrhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_monitor_16-2.3.1-1PGDGrhel10.0.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.0 KiB | [pg_stat_monitor_16-2.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_monitor_16-2.2.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_stat_monitor_16` | `2.1.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_stat_monitor_16-2.1.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_monitor_16-2.1.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-stat-monitor` | `2.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 72.2 KiB | [postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-stat-monitor` | `2.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 70.6 KiB | [postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-stat-monitor` | `2.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 72.3 KiB | [postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-stat-monitor` | `2.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 71.2 KiB | [postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-stat-monitor` | `2.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 82.2 KiB | [postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-stat-monitor` | `2.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 81.3 KiB | [postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-stat-monitor` | `2.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 73.4 KiB | [postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-stat-monitor` | `2.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 73.1 KiB | [postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-stat-monitor` | `2.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 72.6 KiB | [postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-stat-monitor` | `2.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 72.2 KiB | [postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-monitor/postgresql-16-pg-stat-monitor_2.4.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_monitor_15` | `2.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.1 KiB | [pg_stat_monitor_15-2.4.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.4.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.9 KiB | [pg_stat_monitor_15-2.4.0-4PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.4.0-4PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 82.6 KiB | [pg_stat_monitor_15-2.4.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_stat_monitor_15-2.4.0-1PGSTY.el8.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.4 KiB | [pg_stat_monitor_15-2.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.6 KiB | [pg_stat_monitor_15-2.3.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.3.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.6 KiB | [pg_stat_monitor_15-2.3.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.3.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.3 KiB | [pg_stat_monitor_15-2.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.8 KiB | [pg_stat_monitor_15-2.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.2 KiB | [pg_stat_monitor_15-2.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.9 KiB | [pg_stat_monitor_15-2.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.0.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.5 KiB | [pg_stat_monitor_15-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_monitor_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.1 KiB | [pg_stat_monitor_15-1.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_monitor_15-1.1.0-1.rhel8.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.3 KiB | [pg_stat_monitor_15-2.4.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.4.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.2 KiB | [pg_stat_monitor_15-2.4.0-4PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.4.0-4PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 81.1 KiB | [pg_stat_monitor_15-2.4.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_stat_monitor_15-2.4.0-1PGSTY.el8.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.7 KiB | [pg_stat_monitor_15-2.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.9 KiB | [pg_stat_monitor_15-2.3.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.3.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.9 KiB | [pg_stat_monitor_15-2.3.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.3.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.5 KiB | [pg_stat_monitor_15-2.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.0 KiB | [pg_stat_monitor_15-2.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.5 KiB | [pg_stat_monitor_15-2.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.1 KiB | [pg_stat_monitor_15-2.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.0.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.8 KiB | [pg_stat_monitor_15-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_monitor_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 84.7 KiB | [pg_stat_monitor_15-1.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_monitor_15-1.1.0-1.rhel8.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.0 KiB | [pg_stat_monitor_15-2.4.1-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_monitor_15-2.4.1-1PGDG.rhel9.9.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.8 KiB | [pg_stat_monitor_15-2.4.0-4PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_monitor_15-2.4.0-4PGDG.rhel9.9.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 84.0 KiB | [pg_stat_monitor_15-2.4.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_monitor_15-2.4.0-1PGSTY.el9.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.0 KiB | [pg_stat_monitor_15-2.4.1-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_monitor_15-2.4.1-1PGDG.rhel9.9.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.9 KiB | [pg_stat_monitor_15-2.4.0-4PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_monitor_15-2.4.0-4PGDG.rhel9.9.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 83.0 KiB | [pg_stat_monitor_15-2.4.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_monitor_15-2.4.0-1PGSTY.el9.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.6 KiB | [pg_stat_monitor_15-2.4.1-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_stat_monitor_15-2.4.1-1PGDG.rhel10.3.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.3 KiB | [pg_stat_monitor_15-2.4.0-4PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_stat_monitor_15-2.4.0-4PGDG.rhel10.3.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 84.3 KiB | [pg_stat_monitor_15-2.4.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_monitor_15-2.4.0-1PGSTY.el10.x86_64.rpm) |
| `pg_stat_monitor_15` | `2.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.8 KiB | [pg_stat_monitor_15-2.4.1-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_monitor_15-2.4.1-1PGDG.rhel10.3.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.8 KiB | [pg_stat_monitor_15-2.4.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_monitor_15-2.4.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.6 KiB | [pg_stat_monitor_15-2.4.0-4PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_monitor_15-2.4.0-4PGDG.rhel10.3.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.6 KiB | [pg_stat_monitor_15-2.4.0-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_monitor_15-2.4.0-4PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 83.8 KiB | [pg_stat_monitor_15-2.4.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_monitor_15-2.4.0-1PGSTY.el10.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pg_stat_monitor_15-2.4.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_monitor_15-2.4.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.6 KiB | [pg_stat_monitor_15-2.3.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_monitor_15-2.3.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.6 KiB | [pg_stat_monitor_15-2.3.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_monitor_15-2.3.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.6 KiB | [pg_stat_monitor_15-2.3.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_monitor_15-2.3.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.6 KiB | [pg_stat_monitor_15-2.3.1-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_monitor_15-2.3.1-1PGDGrhel10.1.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.6 KiB | [pg_stat_monitor_15-2.3.1-1PGDGrhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_monitor_15-2.3.1-1PGDGrhel10.0.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.6 KiB | [pg_stat_monitor_15-2.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_monitor_15-2.2.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_stat_monitor_15` | `2.1.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.3 KiB | [pg_stat_monitor_15-2.1.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_monitor_15-2.1.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-stat-monitor` | `2.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 73.0 KiB | [postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-stat-monitor` | `2.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 71.5 KiB | [postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-stat-monitor` | `2.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 73.3 KiB | [postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-stat-monitor` | `2.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 72.0 KiB | [postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-stat-monitor` | `2.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 83.4 KiB | [postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-stat-monitor` | `2.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 82.5 KiB | [postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-stat-monitor` | `2.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 74.3 KiB | [postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-stat-monitor` | `2.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 74.2 KiB | [postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-stat-monitor` | `2.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 73.6 KiB | [postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-stat-monitor` | `2.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 73.5 KiB | [postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-monitor/postgresql-15-pg-stat-monitor_2.4.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/percona/pg_stat_monitor" title="Repository" icon="github" subtitle="github.com/percona/pg_stat_monitor" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_stat_monitor-2.4.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_stat_monitor;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_stat_monitor;		# install via package name, for the active PG version

pig install pg_stat_monitor -v 19;   # install for PG 19
pig install pg_stat_monitor -v 18;   # install for PG 18
pig install pg_stat_monitor -v 17;   # install for PG 17
pig install pg_stat_monitor -v 16;   # install for PG 16
pig install pg_stat_monitor -v 15;   # install for PG 15

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_stat_monitor';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_stat_monitor;
```




## Usage

> [pg_stat_monitor: query performance monitoring tool for PostgreSQL](https://github.com/percona/pg_stat_monitor)

pg_stat_monitor is an advanced replacement for `pg_stat_statements` that aggregates statistics into configurable time-based buckets, provides query origin information, actual parameter capture, and query plan details.

### Querying Statistics

```sql
-- Basic query monitoring
SELECT application_name, userid::regrole AS user_name,
       datname AS database_name, substr(query, 0, 50) AS query,
       calls, client_ip
FROM pg_stat_monitor;

-- Bucket-based time analysis
SELECT bucket, bucket_start_time, query, calls,
       mean_exec_time, total_exec_time
FROM pg_stat_monitor
ORDER BY total_exec_time DESC;

-- Show query plans
SELECT query, query_plan FROM pg_stat_monitor
WHERE query_plan IS NOT NULL;
```

### Key Features

- **Time-based buckets**: Statistics are grouped into configurable intervals for more accurate analysis
- **Client IP tracking**: Each query records the originating client IP address
- **Actual parameters**: Optionally capture real query parameter values instead of placeholders
- **Query plans**: Each query is accompanied by its actual execution plan
- **Top query tracking**: Identify the most resource-intensive queries per bucket
- **Histogram support**: Visual timing distribution via histogram function

### Functions

```sql
-- Reset all statistics
SELECT pg_stat_monitor_reset();

-- View internal info
SELECT * FROM pg_stat_monitor_info;
```

### Configuration

Key parameters (set in `postgresql.conf`):

| Parameter | Description |
|-----------|-------------|
| `pg_stat_monitor.pgsm_max` | Maximum number of statements tracked |
| `pg_stat_monitor.pgsm_query_max_len` | Maximum query length |
| `pg_stat_monitor.pgsm_bucket_time` | Bucket duration in seconds |
| `pg_stat_monitor.pgsm_max_buckets` | Maximum number of buckets |
| `pg_stat_monitor.pgsm_enable_query_plan` | Enable query plan capture |
| `pg_stat_monitor.pgsm_track` | Track: `top`, `all`, or `none` |
