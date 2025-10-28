---
title: "timescaledb"
linkTitle: "timescaledb"
description: "Enables scalable inserts and complex queries for time-series data"
weight: 1000
categories: ["TIME"]
width: full
---

Enables scalable inserts and complex queries for time-series data


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1000** | {{< badge content="timescaledb" link="https://github.com/timescale/timescaledb" >}} | {{< ext "timescaledb" >}} | `2.22.1` | {{< category "TIME" >}} | {{< license "Timescale" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "timescaledb_toolkit" >}} {{< ext "timeseries" >}} {{< ext "pg_cron" >}} {{< ext "pg_partman" >}} {{< ext "periods" >}} {{< ext "temporal_tables" >}} {{< ext "emaj" >}} {{< ext "pg_task" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/timescaledb" >}} | `2.22.1` | {{< bg "18" "timescaledb-tsl_18*" "red" >}} {{< bg "17" "timescaledb-tsl_17*" "green" >}} {{< bg "16" "timescaledb-tsl_16*" "green" >}} {{< bg "15" "timescaledb-tsl_15*" "green" >}} {{< bg "14" "timescaledb-tsl_14*" "green" >}} {{< bg "13" "timescaledb-tsl_13*" "red" >}} | `timescaledb-tsl_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/timescaledb" >}} | `2.22.1` | {{< bg "18" "postgresql-18-timescaledb-tsl" "red" >}} {{< bg "17" "postgresql-17-timescaledb-tsl" "green" >}} {{< bg "16" "postgresql-16-timescaledb-tsl" "green" >}} {{< bg "15" "postgresql-15-timescaledb-tsl" "green" >}} {{< bg "14" "postgresql-14-timescaledb-tsl" "green" >}} {{< bg "13" "postgresql-13-timescaledb-tsl" "red" >}} | `postgresql-$v-timescaledb-tsl` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "timescaledb-tsl_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.22.1" "timescaledb-tsl_17 : AVAIL 21" "blue" >}} | {{< bg "PGDG 2.22.1" "timescaledb-tsl_16 : AVAIL 27" "blue" >}} | {{< bg "PGDG 2.22.1" "timescaledb-tsl_15 : AVAIL 26" "blue" >}} | {{< bg "PGDG 2.9.3" "timescaledb-tsl_14 : AVAIL 35" "blue" >}} | {{< bg "PGDG 2.9.3" "timescaledb-tsl_13 : AVAIL 25" "blue" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "timescaledb-tsl_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.22.0" "timescaledb-tsl_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.22.0" "timescaledb-tsl_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.22.0" "timescaledb-tsl_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.19.3" "timescaledb-tsl_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "timescaledb-tsl_13 : MISS 0" "red" >}}      |
|    `el9.x86_64`    |      {{< bg "MISS" "timescaledb-tsl_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.22.1" "timescaledb-tsl_17 : AVAIL 20" "blue" >}} | {{< bg "PGDG 2.22.1" "timescaledb-tsl_16 : AVAIL 26" "blue" >}} | {{< bg "PGDG 2.9.3" "timescaledb-tsl_15 : AVAIL 34" "blue" >}} | {{< bg "PGDG 2.9.3" "timescaledb-tsl_14 : AVAIL 29" "blue" >}} | {{< bg "PGDG 2.9.3" "timescaledb-tsl_13 : AVAIL 16" "blue" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "timescaledb-tsl_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.22.0" "timescaledb-tsl_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.22.0" "timescaledb-tsl_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.22.0" "timescaledb-tsl_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.19.3" "timescaledb-tsl_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "timescaledb-tsl_13 : MISS 0" "red" >}}      |
|    `el10.x86_64`    |      {{< bg "MISS" "timescaledb-tsl_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.22.1" "timescaledb-tsl_17 : AVAIL 9" "blue" >}} | {{< bg "PGDG 2.22.1" "timescaledb-tsl_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 2.22.1" "timescaledb-tsl_15 : AVAIL 9" "blue" >}} |      {{< bg "MISS" "timescaledb-tsl_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "timescaledb-tsl_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "timescaledb-tsl_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "timescaledb-tsl_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "timescaledb-tsl_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "timescaledb-tsl_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "timescaledb-tsl_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "timescaledb-tsl_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-timescaledb-tsl : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.22.0" "postgresql-17-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.22.0" "postgresql-16-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.22.0" "postgresql-15-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.19.3" "postgresql-14-timescaledb-tsl : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-timescaledb-tsl : MISS 0" "red" >}}      |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-timescaledb-tsl : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.22.0" "postgresql-17-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.22.0" "postgresql-16-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.22.0" "postgresql-15-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.19.3" "postgresql-14-timescaledb-tsl : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-timescaledb-tsl : MISS 0" "red" >}}      |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-timescaledb-tsl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-timescaledb-tsl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-timescaledb-tsl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-timescaledb-tsl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-timescaledb-tsl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-timescaledb-tsl : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-timescaledb-tsl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-timescaledb-tsl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-timescaledb-tsl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-timescaledb-tsl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-timescaledb-tsl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-timescaledb-tsl : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-timescaledb-tsl : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.22.0" "postgresql-17-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.22.0" "postgresql-16-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.22.0" "postgresql-15-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.19.3" "postgresql-14-timescaledb-tsl : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-timescaledb-tsl : MISS 0" "red" >}}      |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-timescaledb-tsl : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.22.0" "postgresql-17-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.22.0" "postgresql-16-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.22.0" "postgresql-15-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.19.3" "postgresql-14-timescaledb-tsl : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-timescaledb-tsl : MISS 0" "red" >}}      |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-timescaledb-tsl : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.22.0" "postgresql-17-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.22.0" "postgresql-16-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.22.0" "postgresql-15-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.19.3" "postgresql-14-timescaledb-tsl : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-timescaledb-tsl : MISS 0" "red" >}}      |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-timescaledb-tsl : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.22.0" "postgresql-17-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.22.0" "postgresql-16-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.22.0" "postgresql-15-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.19.3" "postgresql-14-timescaledb-tsl : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-timescaledb-tsl : MISS 0" "red" >}}      |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `timescaledb-tsl_17` | 2.22.1 | `el8.x86_64` | pgdg | 733.1 KiB | [timescaledb-tsl_17-2.22.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.22.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.22.0 | `el8.x86_64` | pigsty | 802.5 KiB | [timescaledb-tsl_17-2.22.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-tsl_17-2.22.0-1PIGSTY.el8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.22.0 | `el8.x86_64` | pgdg | 730.7 KiB | [timescaledb-tsl_17-2.22.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.22.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.21.3 | `el8.x86_64` | pgdg | 738.9 KiB | [timescaledb-tsl_17-2.21.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.21.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.21.2 | `el8.x86_64` | pgdg | 738.7 KiB | [timescaledb-tsl_17-2.21.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.21.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.21.1 | `el8.x86_64` | pigsty | 811.0 KiB | [timescaledb-tsl_17-2.21.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-tsl_17-2.21.1-1PIGSTY.el8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.21.1 | `el8.x86_64` | pgdg | 737.8 KiB | [timescaledb-tsl_17-2.21.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.21.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.21.0 | `el8.x86_64` | pgdg | 737.0 KiB | [timescaledb-tsl_17-2.21.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.21.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.20.3 | `el8.x86_64` | pgdg | 722.7 KiB | [timescaledb-tsl_17-2.20.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.20.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.20.2 | `el8.x86_64` | pgdg | 721.1 KiB | [timescaledb-tsl_17-2.20.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.20.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.20.1 | `el8.x86_64` | pgdg | 720.2 KiB | [timescaledb-tsl_17-2.20.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.20.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.19.3 | `el8.x86_64` | pgdg | 700.0 KiB | [timescaledb-tsl_17-2.19.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.19.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.19.2 | `el8.x86_64` | pgdg | 698.9 KiB | [timescaledb-tsl_17-2.19.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.19.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.19.1 | `el8.x86_64` | pgdg | 698.3 KiB | [timescaledb-tsl_17-2.19.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.19.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.19.0 | `el8.x86_64` | pgdg | 773.5 KiB | [timescaledb-tsl_17-2.19.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.19.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.18.2 | `el8.x86_64` | pgdg | 746.3 KiB | [timescaledb-tsl_17-2.18.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.18.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.18.1 | `el8.x86_64` | pgdg | 745.6 KiB | [timescaledb-tsl_17-2.18.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.18.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.18.0 | `el8.x86_64` | pgdg | 743.6 KiB | [timescaledb-tsl_17-2.18.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.18.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.17.2 | `el8.x86_64` | pgdg | 686.4 KiB | [timescaledb-tsl_17-2.17.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.17.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.17.1 | `el8.x86_64` | pgdg | 685.9 KiB | [timescaledb-tsl_17-2.17.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.17.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.17.0 | `el8.x86_64` | pgdg | 683.4 KiB | [timescaledb-tsl_17-2.17.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.17.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.22.0 | `el8.aarch64` | pigsty | 736.0 KiB | [timescaledb-tsl_17-2.22.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-tsl_17-2.22.0-1PIGSTY.el8.aarch64.rpm) |
| `timescaledb-tsl_17` | 2.21.1 | `el8.aarch64` | pigsty | 741.8 KiB | [timescaledb-tsl_17-2.21.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-tsl_17-2.21.1-1PIGSTY.el8.aarch64.rpm) |
| `timescaledb-tsl_17` | 2.22.1 | `el9.x86_64` | pgdg | 712.8 KiB | [timescaledb-tsl_17-2.22.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.22.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.22.0 | `el9.x86_64` | pigsty | 729.1 KiB | [timescaledb-tsl_17-2.22.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/timescaledb-tsl_17-2.22.0-1PIGSTY.el9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.22.0 | `el9.x86_64` | pgdg | 710.1 KiB | [timescaledb-tsl_17-2.22.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.22.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.21.3 | `el9.x86_64` | pgdg | 718.9 KiB | [timescaledb-tsl_17-2.21.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.21.3-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.21.2 | `el9.x86_64` | pgdg | 715.9 KiB | [timescaledb-tsl_17-2.21.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.21.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.21.1 | `el9.x86_64` | pigsty | 737.0 KiB | [timescaledb-tsl_17-2.21.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/timescaledb-tsl_17-2.21.1-1PIGSTY.el9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.21.1 | `el9.x86_64` | pgdg | 721.9 KiB | [timescaledb-tsl_17-2.21.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.21.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.21.0 | `el9.x86_64` | pgdg | 719.9 KiB | [timescaledb-tsl_17-2.21.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.21.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.20.3 | `el9.x86_64` | pgdg | 698.7 KiB | [timescaledb-tsl_17-2.20.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.20.3-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.20.2 | `el9.x86_64` | pgdg | 695.6 KiB | [timescaledb-tsl_17-2.20.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.20.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.19.3 | `el9.x86_64` | pgdg | 660.9 KiB | [timescaledb-tsl_17-2.19.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.19.3-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.19.2 | `el9.x86_64` | pgdg | 658.1 KiB | [timescaledb-tsl_17-2.19.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.19.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.19.1 | `el9.x86_64` | pgdg | 657.8 KiB | [timescaledb-tsl_17-2.19.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.19.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.19.0 | `el9.x86_64` | pgdg | 679.2 KiB | [timescaledb-tsl_17-2.19.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.19.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.18.2 | `el9.x86_64` | pgdg | 659.2 KiB | [timescaledb-tsl_17-2.18.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.18.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.18.1 | `el9.x86_64` | pgdg | 657.9 KiB | [timescaledb-tsl_17-2.18.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.18.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.18.0 | `el9.x86_64` | pgdg | 658.2 KiB | [timescaledb-tsl_17-2.18.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.18.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.17.2 | `el9.x86_64` | pgdg | 607.6 KiB | [timescaledb-tsl_17-2.17.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.17.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.17.1 | `el9.x86_64` | pgdg | 607.4 KiB | [timescaledb-tsl_17-2.17.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.17.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.17.0 | `el9.x86_64` | pgdg | 606.4 KiB | [timescaledb-tsl_17-2.17.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.17.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.22.0 | `el9.aarch64` | pigsty | 691.0 KiB | [timescaledb-tsl_17-2.22.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-tsl_17-2.22.0-1PIGSTY.el9.aarch64.rpm) |
| `timescaledb-tsl_17` | 2.21.1 | `el9.aarch64` | pigsty | 695.3 KiB | [timescaledb-tsl_17-2.21.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-tsl_17-2.21.1-1PIGSTY.el9.aarch64.rpm) |
| `timescaledb-tsl_17` | 2.22.1 | `el10.x86_64` | pgdg | 737.7 KiB | [timescaledb-tsl_17-2.22.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/timescaledb-tsl_17-2.22.1-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.22.0 | `el10.x86_64` | pgdg | 737.4 KiB | [timescaledb-tsl_17-2.22.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/timescaledb-tsl_17-2.22.0-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.21.3 | `el10.x86_64` | pgdg | 741.3 KiB | [timescaledb-tsl_17-2.21.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/timescaledb-tsl_17-2.21.3-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.21.2 | `el10.x86_64` | pgdg | 740.6 KiB | [timescaledb-tsl_17-2.21.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/timescaledb-tsl_17-2.21.2-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.21.1 | `el10.x86_64` | pgdg | 740.2 KiB | [timescaledb-tsl_17-2.21.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/timescaledb-tsl_17-2.21.1-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.21.0 | `el10.x86_64` | pgdg | 739.6 KiB | [timescaledb-tsl_17-2.21.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/timescaledb-tsl_17-2.21.0-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.20.3 | `el10.x86_64` | pgdg | 723.7 KiB | [timescaledb-tsl_17-2.20.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/timescaledb-tsl_17-2.20.3-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.20.2 | `el10.x86_64` | pgdg | 723.4 KiB | [timescaledb-tsl_17-2.20.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/timescaledb-tsl_17-2.20.2-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_17` | 2.20.1 | `el10.x86_64` | pgdg | 722.8 KiB | [timescaledb-tsl_17-2.20.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/timescaledb-tsl_17-2.20.1-1PGDG.rhel10.x86_64.rpm) |
| `postgresql-17-timescaledb-tsl` | 2.22.0 | `d12.x86_64` | pigsty | 715.7 KiB | [postgresql-17-timescaledb-tsl_2.22.0-9PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-tsl/postgresql-17-timescaledb-tsl_2.22.0-9PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-timescaledb-tsl` | 2.22.0 | `d12.aarch64` | pigsty | 654.5 KiB | [postgresql-17-timescaledb-tsl_2.22.0-9PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-tsl/postgresql-17-timescaledb-tsl_2.22.0-9PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-timescaledb-tsl` | 2.22.0 | `u22.x86_64` | pigsty | 781.2 KiB | [postgresql-17-timescaledb-tsl_2.22.0-9PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-tsl/postgresql-17-timescaledb-tsl_2.22.0-9PIGSTY~jammy_amd64.deb) |
| `postgresql-17-timescaledb-tsl` | 2.22.0 | `u22.aarch64` | pigsty | 746.9 KiB | [postgresql-17-timescaledb-tsl_2.22.0-9PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-tsl/postgresql-17-timescaledb-tsl_2.22.0-9PIGSTY~jammy_arm64.deb) |
| `postgresql-17-timescaledb-tsl` | 2.22.0 | `u24.x86_64` | pigsty | 769.0 KiB | [postgresql-17-timescaledb-tsl_2.22.0-9PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-tsl/postgresql-17-timescaledb-tsl_2.22.0-9PIGSTY~noble_amd64.deb) |
| `postgresql-17-timescaledb-tsl` | 2.22.0 | `u24.aarch64` | pigsty | 746.9 KiB | [postgresql-17-timescaledb-tsl_2.22.0-9PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-tsl/postgresql-17-timescaledb-tsl_2.22.0-9PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `timescaledb-tsl_16` | 2.22.1 | `el8.x86_64` | pgdg | 732.6 KiB | [timescaledb-tsl_16-2.22.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.22.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.22.0 | `el8.x86_64` | pigsty | 801.8 KiB | [timescaledb-tsl_16-2.22.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-tsl_16-2.22.0-1PIGSTY.el8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.22.0 | `el8.x86_64` | pgdg | 730.0 KiB | [timescaledb-tsl_16-2.22.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.22.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.21.3 | `el8.x86_64` | pgdg | 737.4 KiB | [timescaledb-tsl_16-2.21.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.21.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.21.2 | `el8.x86_64` | pgdg | 736.7 KiB | [timescaledb-tsl_16-2.21.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.21.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.21.1 | `el8.x86_64` | pigsty | 809.6 KiB | [timescaledb-tsl_16-2.21.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-tsl_16-2.21.1-1PIGSTY.el8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.21.1 | `el8.x86_64` | pgdg | 735.8 KiB | [timescaledb-tsl_16-2.21.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.21.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.21.0 | `el8.x86_64` | pgdg | 735.5 KiB | [timescaledb-tsl_16-2.21.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.21.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.20.3 | `el8.x86_64` | pgdg | 721.2 KiB | [timescaledb-tsl_16-2.20.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.20.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.20.2 | `el8.x86_64` | pgdg | 720.0 KiB | [timescaledb-tsl_16-2.20.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.20.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.20.1 | `el8.x86_64` | pgdg | 719.4 KiB | [timescaledb-tsl_16-2.20.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.20.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.19.3 | `el8.x86_64` | pgdg | 699.6 KiB | [timescaledb-tsl_16-2.19.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.19.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.19.2 | `el8.x86_64` | pgdg | 698.4 KiB | [timescaledb-tsl_16-2.19.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.19.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.19.1 | `el8.x86_64` | pgdg | 697.6 KiB | [timescaledb-tsl_16-2.19.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.19.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.19.0 | `el8.x86_64` | pgdg | 772.7 KiB | [timescaledb-tsl_16-2.19.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.19.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.18.2 | `el8.x86_64` | pgdg | 745.8 KiB | [timescaledb-tsl_16-2.18.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.18.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.18.1 | `el8.x86_64` | pgdg | 744.7 KiB | [timescaledb-tsl_16-2.18.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.18.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.18.0 | `el8.x86_64` | pgdg | 742.4 KiB | [timescaledb-tsl_16-2.18.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.18.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.17.2 | `el8.x86_64` | pgdg | 686.8 KiB | [timescaledb-tsl_16-2.17.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.17.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.17.1 | `el8.x86_64` | pgdg | 686.3 KiB | [timescaledb-tsl_16-2.17.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.17.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.17.0 | `el8.x86_64` | pgdg | 683.7 KiB | [timescaledb-tsl_16-2.17.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.17.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.16.0 | `el8.x86_64` | pgdg | 641.5 KiB | [timescaledb-tsl_16-2.16.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.16.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.15.3 | `el8.x86_64` | pgdg | 638.4 KiB | [timescaledb-tsl_16-2.15.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.15.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.15.2 | `el8.x86_64` | pgdg | 637.7 KiB | [timescaledb-tsl_16-2.15.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.15.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.15.0 | `el8.x86_64` | pgdg | 635.3 KiB | [timescaledb-tsl_16-2.15.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.15.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.13.1 | `el8.x86_64` | pgdg | 758.6 KiB | [timescaledb-tsl_16-2.13.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.13.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.13.0 | `el8.x86_64` | pgdg | 757.4 KiB | [timescaledb-tsl_16-2.13.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.13.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.22.0 | `el8.aarch64` | pigsty | 736.6 KiB | [timescaledb-tsl_16-2.22.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-tsl_16-2.22.0-1PIGSTY.el8.aarch64.rpm) |
| `timescaledb-tsl_16` | 2.21.1 | `el8.aarch64` | pigsty | 741.3 KiB | [timescaledb-tsl_16-2.21.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-tsl_16-2.21.1-1PIGSTY.el8.aarch64.rpm) |
| `timescaledb-tsl_16` | 2.22.1 | `el9.x86_64` | pgdg | 712.5 KiB | [timescaledb-tsl_16-2.22.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.22.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.22.0 | `el9.x86_64` | pigsty | 727.0 KiB | [timescaledb-tsl_16-2.22.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/timescaledb-tsl_16-2.22.0-1PIGSTY.el9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.22.0 | `el9.x86_64` | pgdg | 709.3 KiB | [timescaledb-tsl_16-2.22.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.22.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.21.3 | `el9.x86_64` | pgdg | 715.6 KiB | [timescaledb-tsl_16-2.21.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.21.3-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.21.2 | `el9.x86_64` | pgdg | 713.2 KiB | [timescaledb-tsl_16-2.21.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.21.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.21.1 | `el9.x86_64` | pigsty | 737.3 KiB | [timescaledb-tsl_16-2.21.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/timescaledb-tsl_16-2.21.1-1PIGSTY.el9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.21.1 | `el9.x86_64` | pgdg | 713.6 KiB | [timescaledb-tsl_16-2.21.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.21.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.21.0 | `el9.x86_64` | pgdg | 720.3 KiB | [timescaledb-tsl_16-2.21.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.21.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.20.3 | `el9.x86_64` | pgdg | 695.7 KiB | [timescaledb-tsl_16-2.20.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.20.3-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.20.2 | `el9.x86_64` | pgdg | 694.8 KiB | [timescaledb-tsl_16-2.20.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.20.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.19.3 | `el9.x86_64` | pgdg | 658.9 KiB | [timescaledb-tsl_16-2.19.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.19.3-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.19.2 | `el9.x86_64` | pgdg | 657.6 KiB | [timescaledb-tsl_16-2.19.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.19.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.19.1 | `el9.x86_64` | pgdg | 656.1 KiB | [timescaledb-tsl_16-2.19.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.19.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.19.0 | `el9.x86_64` | pgdg | 677.2 KiB | [timescaledb-tsl_16-2.19.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.19.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.18.2 | `el9.x86_64` | pgdg | 656.5 KiB | [timescaledb-tsl_16-2.18.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.18.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.18.1 | `el9.x86_64` | pgdg | 655.8 KiB | [timescaledb-tsl_16-2.18.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.18.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.18.0 | `el9.x86_64` | pgdg | 653.2 KiB | [timescaledb-tsl_16-2.18.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.18.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.17.2 | `el9.x86_64` | pgdg | 607.8 KiB | [timescaledb-tsl_16-2.17.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.17.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.17.1 | `el9.x86_64` | pgdg | 606.8 KiB | [timescaledb-tsl_16-2.17.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.17.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.17.0 | `el9.x86_64` | pgdg | 606.3 KiB | [timescaledb-tsl_16-2.17.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.17.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.16.0 | `el9.x86_64` | pgdg | 571.2 KiB | [timescaledb-tsl_16-2.16.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.16.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.15.3 | `el9.x86_64` | pgdg | 565.9 KiB | [timescaledb-tsl_16-2.15.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.15.3-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.15.2 | `el9.x86_64` | pgdg | 566.2 KiB | [timescaledb-tsl_16-2.15.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.15.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.15.0 | `el9.x86_64` | pgdg | 565.8 KiB | [timescaledb-tsl_16-2.15.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.15.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.13.1 | `el9.x86_64` | pgdg | 716.4 KiB | [timescaledb-tsl_16-2.13.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.13.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.13.0 | `el9.x86_64` | pgdg | 715.8 KiB | [timescaledb-tsl_16-2.13.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.13.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.22.0 | `el9.aarch64` | pigsty | 691.2 KiB | [timescaledb-tsl_16-2.22.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-tsl_16-2.22.0-1PIGSTY.el9.aarch64.rpm) |
| `timescaledb-tsl_16` | 2.21.1 | `el9.aarch64` | pigsty | 695.7 KiB | [timescaledb-tsl_16-2.21.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-tsl_16-2.21.1-1PIGSTY.el9.aarch64.rpm) |
| `timescaledb-tsl_16` | 2.22.1 | `el10.x86_64` | pgdg | 737.0 KiB | [timescaledb-tsl_16-2.22.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/timescaledb-tsl_16-2.22.1-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.22.0 | `el10.x86_64` | pgdg | 737.8 KiB | [timescaledb-tsl_16-2.22.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/timescaledb-tsl_16-2.22.0-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.21.3 | `el10.x86_64` | pgdg | 741.3 KiB | [timescaledb-tsl_16-2.21.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/timescaledb-tsl_16-2.21.3-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.21.2 | `el10.x86_64` | pgdg | 740.8 KiB | [timescaledb-tsl_16-2.21.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/timescaledb-tsl_16-2.21.2-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.21.1 | `el10.x86_64` | pgdg | 739.4 KiB | [timescaledb-tsl_16-2.21.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/timescaledb-tsl_16-2.21.1-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.21.0 | `el10.x86_64` | pgdg | 739.3 KiB | [timescaledb-tsl_16-2.21.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/timescaledb-tsl_16-2.21.0-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.20.3 | `el10.x86_64` | pgdg | 726.5 KiB | [timescaledb-tsl_16-2.20.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/timescaledb-tsl_16-2.20.3-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.20.2 | `el10.x86_64` | pgdg | 723.1 KiB | [timescaledb-tsl_16-2.20.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/timescaledb-tsl_16-2.20.2-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_16` | 2.20.1 | `el10.x86_64` | pgdg | 723.6 KiB | [timescaledb-tsl_16-2.20.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/timescaledb-tsl_16-2.20.1-1PGDG.rhel10.x86_64.rpm) |
| `postgresql-16-timescaledb-tsl` | 2.22.0 | `d12.x86_64` | pigsty | 713.5 KiB | [postgresql-16-timescaledb-tsl_2.22.0-9PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-tsl/postgresql-16-timescaledb-tsl_2.22.0-9PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-timescaledb-tsl` | 2.22.0 | `d12.aarch64` | pigsty | 651.4 KiB | [postgresql-16-timescaledb-tsl_2.22.0-9PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-tsl/postgresql-16-timescaledb-tsl_2.22.0-9PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-timescaledb-tsl` | 2.22.0 | `u22.x86_64` | pigsty | 778.6 KiB | [postgresql-16-timescaledb-tsl_2.22.0-9PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-tsl/postgresql-16-timescaledb-tsl_2.22.0-9PIGSTY~jammy_amd64.deb) |
| `postgresql-16-timescaledb-tsl` | 2.22.0 | `u22.aarch64` | pigsty | 746.4 KiB | [postgresql-16-timescaledb-tsl_2.22.0-9PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-tsl/postgresql-16-timescaledb-tsl_2.22.0-9PIGSTY~jammy_arm64.deb) |
| `postgresql-16-timescaledb-tsl` | 2.22.0 | `u24.x86_64` | pigsty | 766.2 KiB | [postgresql-16-timescaledb-tsl_2.22.0-9PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-tsl/postgresql-16-timescaledb-tsl_2.22.0-9PIGSTY~noble_amd64.deb) |
| `postgresql-16-timescaledb-tsl` | 2.22.0 | `u24.aarch64` | pigsty | 742.5 KiB | [postgresql-16-timescaledb-tsl_2.22.0-9PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-tsl/postgresql-16-timescaledb-tsl_2.22.0-9PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `timescaledb-tsl_15` | 2.22.1 | `el8.x86_64` | pgdg | 730.8 KiB | [timescaledb-tsl_15-2.22.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.22.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.22.0 | `el8.x86_64` | pigsty | 800.2 KiB | [timescaledb-tsl_15-2.22.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-tsl_15-2.22.0-1PIGSTY.el8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.22.0 | `el8.x86_64` | pgdg | 728.2 KiB | [timescaledb-tsl_15-2.22.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.22.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.21.3 | `el8.x86_64` | pgdg | 736.2 KiB | [timescaledb-tsl_15-2.21.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.21.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.21.2 | `el8.x86_64` | pgdg | 735.8 KiB | [timescaledb-tsl_15-2.21.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.21.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.21.1 | `el8.x86_64` | pigsty | 808.0 KiB | [timescaledb-tsl_15-2.21.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-tsl_15-2.21.1-1PIGSTY.el8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.21.1 | `el8.x86_64` | pgdg | 734.6 KiB | [timescaledb-tsl_15-2.21.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.21.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.21.0 | `el8.x86_64` | pgdg | 734.1 KiB | [timescaledb-tsl_15-2.21.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.21.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.20.3 | `el8.x86_64` | pgdg | 719.9 KiB | [timescaledb-tsl_15-2.20.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.20.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.20.2 | `el8.x86_64` | pgdg | 718.6 KiB | [timescaledb-tsl_15-2.20.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.20.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.20.1 | `el8.x86_64` | pgdg | 718.0 KiB | [timescaledb-tsl_15-2.20.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.20.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.19.3 | `el8.x86_64` | pgdg | 699.1 KiB | [timescaledb-tsl_15-2.19.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.19.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.19.2 | `el8.x86_64` | pgdg | 697.5 KiB | [timescaledb-tsl_15-2.19.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.19.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.19.1 | `el8.x86_64` | pgdg | 696.9 KiB | [timescaledb-tsl_15-2.19.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.19.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.19.0 | `el8.x86_64` | pgdg | 772.4 KiB | [timescaledb-tsl_15-2.19.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.19.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.18.2 | `el8.x86_64` | pgdg | 744.9 KiB | [timescaledb-tsl_15-2.18.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.18.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.18.1 | `el8.x86_64` | pgdg | 743.9 KiB | [timescaledb-tsl_15-2.18.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.18.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.18.0 | `el8.x86_64` | pgdg | 742.0 KiB | [timescaledb-tsl_15-2.18.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.18.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.17.2 | `el8.x86_64` | pgdg | 686.1 KiB | [timescaledb-tsl_15-2.17.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.17.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.17.1 | `el8.x86_64` | pgdg | 685.5 KiB | [timescaledb-tsl_15-2.17.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.17.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.17.0 | `el8.x86_64` | pgdg | 682.7 KiB | [timescaledb-tsl_15-2.17.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.17.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.16.0 | `el8.x86_64` | pgdg | 640.9 KiB | [timescaledb-tsl_15-2.16.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.16.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.15.3 | `el8.x86_64` | pgdg | 637.4 KiB | [timescaledb-tsl_15-2.15.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.15.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.15.2 | `el8.x86_64` | pgdg | 637.0 KiB | [timescaledb-tsl_15-2.15.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.15.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.15.0 | `el8.x86_64` | pgdg | 633.4 KiB | [timescaledb-tsl_15-2.15.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.15.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.13.1 | `el8.x86_64` | pgdg | 768.2 KiB | [timescaledb-tsl_15-2.13.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.13.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.22.0 | `el8.aarch64` | pigsty | 735.3 KiB | [timescaledb-tsl_15-2.22.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-tsl_15-2.22.0-1PIGSTY.el8.aarch64.rpm) |
| `timescaledb-tsl_15` | 2.21.1 | `el8.aarch64` | pigsty | 739.8 KiB | [timescaledb-tsl_15-2.21.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-tsl_15-2.21.1-1PIGSTY.el8.aarch64.rpm) |
| `timescaledb-tsl_15` | 2.9.3 | `el9.x86_64` | pgdg | 643.6 KiB | [timescaledb-tsl_15-2.9.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.9.3-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.9.2 | `el9.x86_64` | pgdg | 642.8 KiB | [timescaledb-tsl_15-2.9.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.9.2-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.9.1 | `el9.x86_64` | pgdg | 642.3 KiB | [timescaledb-tsl_15-2.9.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.9.1-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.22.1 | `el9.x86_64` | pgdg | 709.8 KiB | [timescaledb-tsl_15-2.22.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.22.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.22.0 | `el9.x86_64` | pigsty | 726.4 KiB | [timescaledb-tsl_15-2.22.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/timescaledb-tsl_15-2.22.0-1PIGSTY.el9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.22.0 | `el9.x86_64` | pgdg | 708.3 KiB | [timescaledb-tsl_15-2.22.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.22.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.21.3 | `el9.x86_64` | pgdg | 713.0 KiB | [timescaledb-tsl_15-2.21.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.21.3-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.21.2 | `el9.x86_64` | pgdg | 715.6 KiB | [timescaledb-tsl_15-2.21.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.21.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.21.1 | `el9.x86_64` | pigsty | 731.9 KiB | [timescaledb-tsl_15-2.21.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/timescaledb-tsl_15-2.21.1-1PIGSTY.el9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.21.1 | `el9.x86_64` | pgdg | 718.8 KiB | [timescaledb-tsl_15-2.21.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.21.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.21.0 | `el9.x86_64` | pgdg | 717.8 KiB | [timescaledb-tsl_15-2.21.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.21.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.20.3 | `el9.x86_64` | pgdg | 697.3 KiB | [timescaledb-tsl_15-2.20.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.20.3-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.20.2 | `el9.x86_64` | pgdg | 693.5 KiB | [timescaledb-tsl_15-2.20.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.20.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.19.3 | `el9.x86_64` | pgdg | 657.5 KiB | [timescaledb-tsl_15-2.19.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.19.3-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.19.2 | `el9.x86_64` | pgdg | 655.8 KiB | [timescaledb-tsl_15-2.19.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.19.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.19.1 | `el9.x86_64` | pgdg | 654.6 KiB | [timescaledb-tsl_15-2.19.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.19.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.19.0 | `el9.x86_64` | pgdg | 676.9 KiB | [timescaledb-tsl_15-2.19.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.19.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.18.2 | `el9.x86_64` | pgdg | 655.4 KiB | [timescaledb-tsl_15-2.18.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.18.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.18.1 | `el9.x86_64` | pgdg | 654.6 KiB | [timescaledb-tsl_15-2.18.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.18.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.18.0 | `el9.x86_64` | pgdg | 652.2 KiB | [timescaledb-tsl_15-2.18.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.18.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.17.2 | `el9.x86_64` | pgdg | 607.4 KiB | [timescaledb-tsl_15-2.17.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.17.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.17.1 | `el9.x86_64` | pgdg | 607.7 KiB | [timescaledb-tsl_15-2.17.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.17.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.17.0 | `el9.x86_64` | pgdg | 606.6 KiB | [timescaledb-tsl_15-2.17.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.17.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.16.0 | `el9.x86_64` | pgdg | 570.9 KiB | [timescaledb-tsl_15-2.16.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.16.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.15.3 | `el9.x86_64` | pgdg | 565.1 KiB | [timescaledb-tsl_15-2.15.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.15.3-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.15.2 | `el9.x86_64` | pgdg | 565.2 KiB | [timescaledb-tsl_15-2.15.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.15.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.13.1 | `el9.x86_64` | pgdg | 725.1 KiB | [timescaledb-tsl_15-2.13.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.13.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.13.0 | `el9.x86_64` | pgdg | 726.7 KiB | [timescaledb-tsl_15-2.13.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.13.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.12.2 | `el9.x86_64` | pgdg | 708.6 KiB | [timescaledb-tsl_15-2.12.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.12.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.12.0 | `el9.x86_64` | pgdg | 707.6 KiB | [timescaledb-tsl_15-2.12.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.12.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.11.2 | `el9.x86_64` | pgdg | 677.1 KiB | [timescaledb-tsl_15-2.11.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.11.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.11.1 | `el9.x86_64` | pgdg | 675.4 KiB | [timescaledb-tsl_15-2.11.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.11.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.11.0 | `el9.x86_64` | pgdg | 674.1 KiB | [timescaledb-tsl_15-2.11.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.11.0-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.10.2 | `el9.x86_64` | pgdg | 652.1 KiB | [timescaledb-tsl_15-2.10.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/timescaledb-tsl_15-2.10.2-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.22.0 | `el9.aarch64` | pigsty | 690.1 KiB | [timescaledb-tsl_15-2.22.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-tsl_15-2.22.0-1PIGSTY.el9.aarch64.rpm) |
| `timescaledb-tsl_15` | 2.21.1 | `el9.aarch64` | pigsty | 692.2 KiB | [timescaledb-tsl_15-2.21.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-tsl_15-2.21.1-1PIGSTY.el9.aarch64.rpm) |
| `timescaledb-tsl_15` | 2.22.1 | `el10.x86_64` | pgdg | 736.3 KiB | [timescaledb-tsl_15-2.22.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/timescaledb-tsl_15-2.22.1-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.22.0 | `el10.x86_64` | pgdg | 735.7 KiB | [timescaledb-tsl_15-2.22.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/timescaledb-tsl_15-2.22.0-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.21.3 | `el10.x86_64` | pgdg | 741.2 KiB | [timescaledb-tsl_15-2.21.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/timescaledb-tsl_15-2.21.3-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.21.2 | `el10.x86_64` | pgdg | 739.9 KiB | [timescaledb-tsl_15-2.21.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/timescaledb-tsl_15-2.21.2-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.21.1 | `el10.x86_64` | pgdg | 738.9 KiB | [timescaledb-tsl_15-2.21.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/timescaledb-tsl_15-2.21.1-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.21.0 | `el10.x86_64` | pgdg | 738.6 KiB | [timescaledb-tsl_15-2.21.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/timescaledb-tsl_15-2.21.0-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.20.3 | `el10.x86_64` | pgdg | 724.8 KiB | [timescaledb-tsl_15-2.20.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/timescaledb-tsl_15-2.20.3-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.20.2 | `el10.x86_64` | pgdg | 721.9 KiB | [timescaledb-tsl_15-2.20.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/timescaledb-tsl_15-2.20.2-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_15` | 2.20.1 | `el10.x86_64` | pgdg | 721.6 KiB | [timescaledb-tsl_15-2.20.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/timescaledb-tsl_15-2.20.1-1PGDG.rhel10.x86_64.rpm) |
| `postgresql-15-timescaledb-tsl` | 2.22.0 | `d12.x86_64` | pigsty | 708.2 KiB | [postgresql-15-timescaledb-tsl_2.22.0-9PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-tsl/postgresql-15-timescaledb-tsl_2.22.0-9PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-timescaledb-tsl` | 2.22.0 | `d12.aarch64` | pigsty | 647.1 KiB | [postgresql-15-timescaledb-tsl_2.22.0-9PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-tsl/postgresql-15-timescaledb-tsl_2.22.0-9PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-timescaledb-tsl` | 2.22.0 | `u22.x86_64` | pigsty | 773.3 KiB | [postgresql-15-timescaledb-tsl_2.22.0-9PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-tsl/postgresql-15-timescaledb-tsl_2.22.0-9PIGSTY~jammy_amd64.deb) |
| `postgresql-15-timescaledb-tsl` | 2.22.0 | `u22.aarch64` | pigsty | 742.1 KiB | [postgresql-15-timescaledb-tsl_2.22.0-9PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-tsl/postgresql-15-timescaledb-tsl_2.22.0-9PIGSTY~jammy_arm64.deb) |
| `postgresql-15-timescaledb-tsl` | 2.22.0 | `u24.x86_64` | pigsty | 762.2 KiB | [postgresql-15-timescaledb-tsl_2.22.0-9PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-tsl/postgresql-15-timescaledb-tsl_2.22.0-9PIGSTY~noble_amd64.deb) |
| `postgresql-15-timescaledb-tsl` | 2.22.0 | `u24.aarch64` | pigsty | 739.3 KiB | [postgresql-15-timescaledb-tsl_2.22.0-9PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-tsl/postgresql-15-timescaledb-tsl_2.22.0-9PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `timescaledb-tsl_14` | 2.9.3 | `el8.x86_64` | pgdg | 676.4 KiB | [timescaledb-tsl_14-2.9.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.9.3-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.9.2 | `el8.x86_64` | pgdg | 675.5 KiB | [timescaledb-tsl_14-2.9.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.9.2-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.9.1 | `el8.x86_64` | pgdg | 674.2 KiB | [timescaledb-tsl_14-2.9.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.9.1-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.8.1 | `el8.x86_64` | pgdg | 644.5 KiB | [timescaledb-tsl_14-2.8.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.8.1-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.8.0 | `el8.x86_64` | pgdg | 643.2 KiB | [timescaledb-tsl_14-2.8.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.8.0-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.7.2 | `el8.x86_64` | pgdg | 616.0 KiB | [timescaledb-tsl_14-2.7.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.7.2-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.6.1 | `el8.x86_64` | pgdg | 594.7 KiB | [timescaledb-tsl_14-2.6.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.6.1-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.6.0 | `el8.x86_64` | pgdg | 593.1 KiB | [timescaledb-tsl_14-2.6.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.6.0-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.5.2 | `el8.x86_64` | pgdg | 579.1 KiB | [timescaledb-tsl_14-2.5.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.5.2-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.5.1 | `el8.x86_64` | pgdg | 619.0 KiB | [timescaledb-tsl_14-2.5.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.5.1-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.5.0 | `el8.x86_64` | pgdg | 617.4 KiB | [timescaledb-tsl_14-2.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.5.0-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.19.3 | `el8.x86_64` | pigsty | 763.9 KiB | [timescaledb-tsl_14-2.19.3-9PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-tsl_14-2.19.3-9PIGSTY.el8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.19.3 | `el8.x86_64` | pgdg | 693.2 KiB | [timescaledb-tsl_14-2.19.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.19.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.19.2 | `el8.x86_64` | pgdg | 692.0 KiB | [timescaledb-tsl_14-2.19.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.19.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.19.1 | `el8.x86_64` | pgdg | 691.5 KiB | [timescaledb-tsl_14-2.19.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.19.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.19.0 | `el8.x86_64` | pgdg | 766.7 KiB | [timescaledb-tsl_14-2.19.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.19.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.18.2 | `el8.x86_64` | pgdg | 739.7 KiB | [timescaledb-tsl_14-2.18.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.18.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.18.1 | `el8.x86_64` | pgdg | 738.5 KiB | [timescaledb-tsl_14-2.18.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.18.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.18.0 | `el8.x86_64` | pgdg | 736.7 KiB | [timescaledb-tsl_14-2.18.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.18.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.17.2 | `el8.x86_64` | pgdg | 680.1 KiB | [timescaledb-tsl_14-2.17.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.17.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.17.1 | `el8.x86_64` | pgdg | 679.5 KiB | [timescaledb-tsl_14-2.17.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.17.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.17.0 | `el8.x86_64` | pgdg | 677.4 KiB | [timescaledb-tsl_14-2.17.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.17.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.16.0 | `el8.x86_64` | pgdg | 637.4 KiB | [timescaledb-tsl_14-2.16.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.16.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.15.3 | `el8.x86_64` | pgdg | 634.0 KiB | [timescaledb-tsl_14-2.15.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.15.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.15.2 | `el8.x86_64` | pgdg | 633.8 KiB | [timescaledb-tsl_14-2.15.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.15.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.15.0 | `el8.x86_64` | pgdg | 630.3 KiB | [timescaledb-tsl_14-2.15.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.15.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.13.1 | `el8.x86_64` | pgdg | 764.5 KiB | [timescaledb-tsl_14-2.13.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.13.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.13.0 | `el8.x86_64` | pgdg | 763.3 KiB | [timescaledb-tsl_14-2.13.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.13.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.12.2 | `el8.x86_64` | pgdg | 747.2 KiB | [timescaledb-tsl_14-2.12.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.12.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.12.0 | `el8.x86_64` | pgdg | 745.9 KiB | [timescaledb-tsl_14-2.12.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.12.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.11.2 | `el8.x86_64` | pgdg | 710.0 KiB | [timescaledb-tsl_14-2.11.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.11.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.11.1 | `el8.x86_64` | pgdg | 709.2 KiB | [timescaledb-tsl_14-2.11.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.11.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.11.0 | `el8.x86_64` | pgdg | 707.9 KiB | [timescaledb-tsl_14-2.11.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.11.0-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.10.2 | `el8.x86_64` | pgdg | 683.3 KiB | [timescaledb-tsl_14-2.10.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.10.2-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.10.0 | `el8.x86_64` | pgdg | 679.0 KiB | [timescaledb-tsl_14-2.10.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/timescaledb-tsl_14-2.10.0-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.19.3 | `el8.aarch64` | pigsty | 699.3 KiB | [timescaledb-tsl_14-2.19.3-9PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-tsl_14-2.19.3-9PIGSTY.el8.aarch64.rpm) |
| `timescaledb-tsl_14` | 2.9.3 | `el9.x86_64` | pgdg | 643.2 KiB | [timescaledb-tsl_14-2.9.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.9.3-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.9.2 | `el9.x86_64` | pgdg | 641.7 KiB | [timescaledb-tsl_14-2.9.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.9.2-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.9.1 | `el9.x86_64` | pgdg | 641.8 KiB | [timescaledb-tsl_14-2.9.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.9.1-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.8.1 | `el9.x86_64` | pgdg | 611.2 KiB | [timescaledb-tsl_14-2.8.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.8.1-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.7.2 | `el9.x86_64` | pgdg | 585.0 KiB | [timescaledb-tsl_14-2.7.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.7.2-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.7.0 | `el9.x86_64` | pgdg | 578.1 KiB | [timescaledb-tsl_14-2.7.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.7.0-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.19.3 | `el9.x86_64` | pigsty | 670.4 KiB | [timescaledb-tsl_14-2.19.3-9PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/timescaledb-tsl_14-2.19.3-9PIGSTY.el9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.19.3 | `el9.x86_64` | pgdg | 653.0 KiB | [timescaledb-tsl_14-2.19.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.19.3-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.19.2 | `el9.x86_64` | pgdg | 650.9 KiB | [timescaledb-tsl_14-2.19.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.19.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.19.1 | `el9.x86_64` | pgdg | 650.4 KiB | [timescaledb-tsl_14-2.19.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.19.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.19.0 | `el9.x86_64` | pgdg | 671.3 KiB | [timescaledb-tsl_14-2.19.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.19.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.18.2 | `el9.x86_64` | pgdg | 648.5 KiB | [timescaledb-tsl_14-2.18.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.18.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.18.1 | `el9.x86_64` | pgdg | 647.7 KiB | [timescaledb-tsl_14-2.18.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.18.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.18.0 | `el9.x86_64` | pgdg | 646.8 KiB | [timescaledb-tsl_14-2.18.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.18.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.17.2 | `el9.x86_64` | pgdg | 601.4 KiB | [timescaledb-tsl_14-2.17.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.17.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.17.1 | `el9.x86_64` | pgdg | 601.2 KiB | [timescaledb-tsl_14-2.17.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.17.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.17.0 | `el9.x86_64` | pgdg | 600.8 KiB | [timescaledb-tsl_14-2.17.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.17.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.16.0 | `el9.x86_64` | pgdg | 566.9 KiB | [timescaledb-tsl_14-2.16.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.16.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.15.3 | `el9.x86_64` | pgdg | 563.4 KiB | [timescaledb-tsl_14-2.15.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.15.3-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.15.2 | `el9.x86_64` | pgdg | 560.5 KiB | [timescaledb-tsl_14-2.15.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.15.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.13.1 | `el9.x86_64` | pgdg | 723.9 KiB | [timescaledb-tsl_14-2.13.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.13.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.13.0 | `el9.x86_64` | pgdg | 723.4 KiB | [timescaledb-tsl_14-2.13.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.13.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.12.2 | `el9.x86_64` | pgdg | 706.1 KiB | [timescaledb-tsl_14-2.12.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.12.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.12.0 | `el9.x86_64` | pgdg | 704.9 KiB | [timescaledb-tsl_14-2.12.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.12.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.11.2 | `el9.x86_64` | pgdg | 675.9 KiB | [timescaledb-tsl_14-2.11.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.11.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.11.1 | `el9.x86_64` | pgdg | 674.8 KiB | [timescaledb-tsl_14-2.11.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.11.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.11.0 | `el9.x86_64` | pgdg | 673.0 KiB | [timescaledb-tsl_14-2.11.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.11.0-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.10.2 | `el9.x86_64` | pgdg | 650.7 KiB | [timescaledb-tsl_14-2.10.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.10.2-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.10.0 | `el9.x86_64` | pgdg | 646.2 KiB | [timescaledb-tsl_14-2.10.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/timescaledb-tsl_14-2.10.0-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_14` | 2.19.3 | `el9.aarch64` | pigsty | 637.9 KiB | [timescaledb-tsl_14-2.19.3-9PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-tsl_14-2.19.3-9PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-timescaledb-tsl` | 2.19.3 | `d12.x86_64` | pigsty | 672.5 KiB | [postgresql-14-timescaledb-tsl_2.19.3-9PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-tsl/postgresql-14-timescaledb-tsl_2.19.3-9PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-timescaledb-tsl` | 2.19.3 | `d12.aarch64` | pigsty | 613.9 KiB | [postgresql-14-timescaledb-tsl_2.19.3-9PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-tsl/postgresql-14-timescaledb-tsl_2.19.3-9PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-timescaledb-tsl` | 2.19.3 | `u22.x86_64` | pigsty | 714.3 KiB | [postgresql-14-timescaledb-tsl_2.19.3-9PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-tsl/postgresql-14-timescaledb-tsl_2.19.3-9PIGSTY~jammy_amd64.deb) |
| `postgresql-14-timescaledb-tsl` | 2.19.3 | `u22.aarch64` | pigsty | 686.2 KiB | [postgresql-14-timescaledb-tsl_2.19.3-9PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-tsl/postgresql-14-timescaledb-tsl_2.19.3-9PIGSTY~jammy_arm64.deb) |
| `postgresql-14-timescaledb-tsl` | 2.19.3 | `u24.x86_64` | pigsty | 702.7 KiB | [postgresql-14-timescaledb-tsl_2.19.3-9PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-tsl/postgresql-14-timescaledb-tsl_2.19.3-9PIGSTY~noble_amd64.deb) |
| `postgresql-14-timescaledb-tsl` | 2.19.3 | `u24.aarch64` | pigsty | 680.3 KiB | [postgresql-14-timescaledb-tsl_2.19.3-9PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-tsl/postgresql-14-timescaledb-tsl_2.19.3-9PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `timescaledb-tsl_13` | 2.9.3 | `el8.x86_64` | pgdg | 658.1 KiB | [timescaledb-tsl_13-2.9.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.9.3-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.9.2 | `el8.x86_64` | pgdg | 657.0 KiB | [timescaledb-tsl_13-2.9.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.9.2-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.9.1 | `el8.x86_64` | pgdg | 656.4 KiB | [timescaledb-tsl_13-2.9.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.9.1-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.8.1 | `el8.x86_64` | pgdg | 628.4 KiB | [timescaledb-tsl_13-2.8.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.8.1-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.8.0 | `el8.x86_64` | pgdg | 626.4 KiB | [timescaledb-tsl_13-2.8.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.8.0-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.7.2 | `el8.x86_64` | pgdg | 599.9 KiB | [timescaledb-tsl_13-2.7.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.7.2-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.6.1 | `el8.x86_64` | pgdg | 579.8 KiB | [timescaledb-tsl_13-2.6.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.6.1-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.6.0 | `el8.x86_64` | pgdg | 578.0 KiB | [timescaledb-tsl_13-2.6.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.6.0-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.5.2 | `el8.x86_64` | pgdg | 564.2 KiB | [timescaledb-tsl_13-2.5.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.5.2-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.5.1 | `el8.x86_64` | pgdg | 545.5 KiB | [timescaledb-tsl_13-2.5.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.5.1-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.5.0 | `el8.x86_64` | pgdg | 544.6 KiB | [timescaledb-tsl_13-2.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.5.0-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.4.2 | `el8.x86_64` | pgdg | 533.7 KiB | [timescaledb-tsl_13-2.4.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.4.2-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.4.1 | `el8.x86_64` | pgdg | 530.7 KiB | [timescaledb-tsl_13-2.4.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.4.1-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.3.0 | `el8.x86_64` | pgdg | 3.0 MiB | [timescaledb-tsl_13-2.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.3.0-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.15.3 | `el8.x86_64` | pgdg | 613.6 KiB | [timescaledb-tsl_13-2.15.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.15.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.15.2 | `el8.x86_64` | pgdg | 613.2 KiB | [timescaledb-tsl_13-2.15.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.15.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.15.0 | `el8.x86_64` | pgdg | 610.5 KiB | [timescaledb-tsl_13-2.15.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.15.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.13.1 | `el8.x86_64` | pgdg | 741.2 KiB | [timescaledb-tsl_13-2.13.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.13.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.13.0 | `el8.x86_64` | pgdg | 740.3 KiB | [timescaledb-tsl_13-2.13.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.13.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.12.2 | `el8.x86_64` | pgdg | 723.6 KiB | [timescaledb-tsl_13-2.12.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.12.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.12.0 | `el8.x86_64` | pgdg | 722.1 KiB | [timescaledb-tsl_13-2.12.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.12.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.11.2 | `el8.x86_64` | pgdg | 688.9 KiB | [timescaledb-tsl_13-2.11.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.11.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.11.1 | `el8.x86_64` | pgdg | 688.1 KiB | [timescaledb-tsl_13-2.11.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.11.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.11.0 | `el8.x86_64` | pgdg | 686.6 KiB | [timescaledb-tsl_13-2.11.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.11.0-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.10.2 | `el8.x86_64` | pgdg | 665.9 KiB | [timescaledb-tsl_13-2.10.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/timescaledb-tsl_13-2.10.2-1.rhel8.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.9.3 | `el9.x86_64` | pgdg | 631.1 KiB | [timescaledb-tsl_13-2.9.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/timescaledb-tsl_13-2.9.3-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.9.2 | `el9.x86_64` | pgdg | 629.9 KiB | [timescaledb-tsl_13-2.9.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/timescaledb-tsl_13-2.9.2-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.9.1 | `el9.x86_64` | pgdg | 629.5 KiB | [timescaledb-tsl_13-2.9.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/timescaledb-tsl_13-2.9.1-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.8.1 | `el9.x86_64` | pgdg | 600.4 KiB | [timescaledb-tsl_13-2.8.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/timescaledb-tsl_13-2.8.1-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.7.2 | `el9.x86_64` | pgdg | 575.0 KiB | [timescaledb-tsl_13-2.7.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/timescaledb-tsl_13-2.7.2-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.7.0 | `el9.x86_64` | pgdg | 568.7 KiB | [timescaledb-tsl_13-2.7.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/timescaledb-tsl_13-2.7.0-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.15.3 | `el9.x86_64` | pgdg | 544.8 KiB | [timescaledb-tsl_13-2.15.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/timescaledb-tsl_13-2.15.3-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.15.2 | `el9.x86_64` | pgdg | 547.9 KiB | [timescaledb-tsl_13-2.15.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/timescaledb-tsl_13-2.15.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.13.1 | `el9.x86_64` | pgdg | 709.3 KiB | [timescaledb-tsl_13-2.13.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/timescaledb-tsl_13-2.13.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.13.0 | `el9.x86_64` | pgdg | 707.3 KiB | [timescaledb-tsl_13-2.13.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/timescaledb-tsl_13-2.13.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.12.2 | `el9.x86_64` | pgdg | 689.9 KiB | [timescaledb-tsl_13-2.12.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/timescaledb-tsl_13-2.12.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.12.0 | `el9.x86_64` | pgdg | 688.9 KiB | [timescaledb-tsl_13-2.12.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/timescaledb-tsl_13-2.12.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.11.2 | `el9.x86_64` | pgdg | 665.5 KiB | [timescaledb-tsl_13-2.11.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/timescaledb-tsl_13-2.11.2-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.11.1 | `el9.x86_64` | pgdg | 663.4 KiB | [timescaledb-tsl_13-2.11.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/timescaledb-tsl_13-2.11.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.11.0 | `el9.x86_64` | pgdg | 662.4 KiB | [timescaledb-tsl_13-2.11.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/timescaledb-tsl_13-2.11.0-1.rhel9.x86_64.rpm) |
| `timescaledb-tsl_13` | 2.10.2 | `el9.x86_64` | pgdg | 640.1 KiB | [timescaledb-tsl_13-2.10.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/timescaledb-tsl_13-2.10.2-1.rhel9.x86_64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/timescale/timescaledb" title="Repository" icon="github" subtitle="github.com/timescale/timescaledb" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="timescaledb-2.22.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get timescaledb; # get timescaledb source code
pig build dep timescaledb; # install build dependencies
pig build pkg timescaledb; # build extension rpm or deb
pig build ext timescaledb; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install timescaledb; # install by extension name, for the current active PG version
pig ext install timescaledb; # install via package alias, for the active PG version
pig ext install timescaledb -v 17;   # install for PG 17
pig ext install timescaledb -v 16;   # install for PG 16
pig ext install timescaledb -v 15;   # install for PG 15

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION timescaledb CASCADE SCHEMA timescaledb_information;
```



--------

## Usage

Create a table and turn it into hypertable

```sql
DROP TABLE IF EXISTS ts_test;
CREATE TABLE ts_test
(
    id BIGINT PRIMARY KEY,
    ts TIMESTAMPTZ NOT NULL,
    v  INTEGER -- payload
);
SELECT create_hypertable('ts_test', by_range('id'));

INSERT INTO ts_test 
    SELECT i, now() + (i || ' seconds')::INTERVAL, i % 100 
    FROM generate_series(1, 1000000) i;


ALTER TABLE ts_test SET (timescaledb.compress_chunk_time_interval = '24 hours');
```

Continuous Agg Example:

```sql

CREATE MATERIALIZED VIEW continuous_aggregate_daily( timec, minl, sumt, sumh )
WITH (timescaledb.continuous) AS
  SELECT count(*) FROM ts_test;


SELECT add_job('SELECT 1','1h', initial_start => '2024-07-09 18:52:00+00'::timestamptz);
```