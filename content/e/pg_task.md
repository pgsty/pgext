---
title: "pg_task"
linkTitle: "pg_task"
description: "execute any sql command at any specific time at background"
weight: 1080
categories: ["TIME"]
width: full
---

[**pg_task**](https://github.com/RekGRpth/pg_task) : execute any sql command at any specific time at background


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1080** | {{< badge content="pg_task" link="https://github.com/RekGRpth/pg_task" >}} | {{< ext "pg_task" >}} | `2.1.29` | {{< category "TIME" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "timescaledb" >}} {{< ext "pg_cron" >}} {{< ext "pg_later" >}} {{< ext "pg_background" >}} {{< ext "pg_partman" >}} {{< ext "timescaledb_toolkit" >}} {{< ext "timeseries" >}} {{< ext "periods" >}} |

> [!Note] breaks on many systems


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.29` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_task` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.29` | {{< bg "18" "pg_task_18" "green" >}} {{< bg "17" "pg_task_17" "green" >}} {{< bg "16" "pg_task_16" "green" >}} {{< bg "15" "pg_task_15" "green" >}} {{< bg "14" "pg_task_14" "green" >}} | `pg_task_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.29` | {{< bg "18" "postgresql-18-pg-task" "green" >}} {{< bg "17" "postgresql-17-pg-task" "green" >}} {{< bg "16" "postgresql-16-pg-task" "green" >}} {{< bg "15" "postgresql-15-pg-task" "green" >}} {{< bg "14" "postgresql-14-pg-task" "green" >}} | `postgresql-$v-pg-task` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_14 : AVAIL 3" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_14 : AVAIL 3" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_14 : AVAIL 3" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_14 : AVAIL 3" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_14 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.29" "pg_task_14 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-18-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-17-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-16-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-15-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-14-pg-task : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-18-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-17-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-16-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-15-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-14-pg-task : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-18-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-17-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-16-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-15-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-14-pg-task : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-18-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-17-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-16-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-15-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-14-pg-task : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-18-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-17-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-16-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-15-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-14-pg-task : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-18-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-17-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-16-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-15-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-14-pg-task : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-18-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-17-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-16-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-15-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-14-pg-task : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-18-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-17-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-16-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-15-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-14-pg-task : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-18-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-17-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-16-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-15-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-14-pg-task : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-18-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-17-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-16-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-15-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.29" "postgresql-14-pg-task : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_task_18` | `2.1.29` | [el8.x86_64](/os/el8.x86_64) | pigsty | 54.8 KiB | [pg_task_18-2.1.29-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_task_18-2.1.29-1PIGSTY.el8.x86_64.rpm) |
| `pg_task_18` | `2.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 72.4 KiB | [pg_task_18-2.1.7-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_task_18-2.1.7-3PGDG.rhel8.x86_64.rpm) |
| `pg_task_18` | `2.1.29` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.8 KiB | [pg_task_18-2.1.29-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_task_18-2.1.29-1PIGSTY.el8.aarch64.rpm) |
| `pg_task_18` | `2.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.3 KiB | [pg_task_18-2.1.7-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_task_18-2.1.7-3PGDG.rhel8.aarch64.rpm) |
| `pg_task_18` | `2.1.29` | [el9.x86_64](/os/el9.x86_64) | pigsty | 54.7 KiB | [pg_task_18-2.1.29-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_task_18-2.1.29-1PIGSTY.el9.x86_64.rpm) |
| `pg_task_18` | `2.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 63.5 KiB | [pg_task_18-2.1.7-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_task_18-2.1.7-3PGDG.rhel9.x86_64.rpm) |
| `pg_task_18` | `2.1.29` | [el9.aarch64](/os/el9.aarch64) | pigsty | 52.7 KiB | [pg_task_18-2.1.29-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_task_18-2.1.29-1PIGSTY.el9.aarch64.rpm) |
| `pg_task_18` | `2.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.5 KiB | [pg_task_18-2.1.7-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_task_18-2.1.7-3PGDG.rhel9.aarch64.rpm) |
| `pg_task_18` | `2.1.29` | [el10.x86_64](/os/el10.x86_64) | pigsty | 54.9 KiB | [pg_task_18-2.1.29-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_task_18-2.1.29-1PIGSTY.el10.x86_64.rpm) |
| `pg_task_18` | `2.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 59.7 KiB | [pg_task_18-2.1.7-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_task_18-2.1.7-3PGDG.rhel10.x86_64.rpm) |
| `pg_task_18` | `2.1.29` | [el10.aarch64](/os/el10.aarch64) | pigsty | 52.7 KiB | [pg_task_18-2.1.29-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_task_18-2.1.29-1PIGSTY.el10.aarch64.rpm) |
| `pg_task_18` | `2.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.3 KiB | [pg_task_18-2.1.7-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_task_18-2.1.7-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-task` | `2.1.29` | [d12.x86_64](/os/d12.x86_64) | pigsty | 38.5 KiB | [postgresql-18-pg-task_2.1.29-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-18-pg-task_2.1.29-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-task` | `2.1.29` | [d12.aarch64](/os/d12.aarch64) | pigsty | 35.3 KiB | [postgresql-18-pg-task_2.1.29-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-18-pg-task_2.1.29-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-task` | `2.1.29` | [d13.x86_64](/os/d13.x86_64) | pigsty | 38.8 KiB | [postgresql-18-pg-task_2.1.29-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-18-pg-task_2.1.29-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-task` | `2.1.29` | [d13.aarch64](/os/d13.aarch64) | pigsty | 35.6 KiB | [postgresql-18-pg-task_2.1.29-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-18-pg-task_2.1.29-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-task` | `2.1.29` | [u22.x86_64](/os/u22.x86_64) | pigsty | 42.5 KiB | [postgresql-18-pg-task_2.1.29-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-18-pg-task_2.1.29-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-task` | `2.1.29` | [u22.aarch64](/os/u22.aarch64) | pigsty | 41.0 KiB | [postgresql-18-pg-task_2.1.29-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-18-pg-task_2.1.29-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-task` | `2.1.29` | [u24.x86_64](/os/u24.x86_64) | pigsty | 41.1 KiB | [postgresql-18-pg-task_2.1.29-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-18-pg-task_2.1.29-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-task` | `2.1.29` | [u24.aarch64](/os/u24.aarch64) | pigsty | 39.7 KiB | [postgresql-18-pg-task_2.1.29-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-18-pg-task_2.1.29-2PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-task` | `2.1.29` | [u26.x86_64](/os/u26.x86_64) | pigsty | 41.1 KiB | [postgresql-18-pg-task_2.1.29-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-task/postgresql-18-pg-task_2.1.29-2PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-task` | `2.1.29` | [u26.aarch64](/os/u26.aarch64) | pigsty | 39.6 KiB | [postgresql-18-pg-task_2.1.29-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-task/postgresql-18-pg-task_2.1.29-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_task_17` | `2.1.29` | [el8.x86_64](/os/el8.x86_64) | pigsty | 54.8 KiB | [pg_task_17-2.1.29-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_task_17-2.1.29-1PIGSTY.el8.x86_64.rpm) |
| `pg_task_17` | `2.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 72.5 KiB | [pg_task_17-2.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_task_17-2.1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_task_17` | `2.1.29` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.8 KiB | [pg_task_17-2.1.29-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_task_17-2.1.29-1PIGSTY.el8.aarch64.rpm) |
| `pg_task_17` | `2.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.4 KiB | [pg_task_17-2.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_task_17-2.1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_task_17` | `2.1.29` | [el9.x86_64](/os/el9.x86_64) | pigsty | 54.7 KiB | [pg_task_17-2.1.29-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_task_17-2.1.29-1PIGSTY.el9.x86_64.rpm) |
| `pg_task_17` | `2.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 63.3 KiB | [pg_task_17-2.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_task_17-2.1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_task_17` | `2.1.29` | [el9.aarch64](/os/el9.aarch64) | pigsty | 52.7 KiB | [pg_task_17-2.1.29-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_task_17-2.1.29-1PIGSTY.el9.aarch64.rpm) |
| `pg_task_17` | `2.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.4 KiB | [pg_task_17-2.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_task_17-2.1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_task_17` | `2.1.29` | [el10.x86_64](/os/el10.x86_64) | pigsty | 54.8 KiB | [pg_task_17-2.1.29-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_task_17-2.1.29-1PIGSTY.el10.x86_64.rpm) |
| `pg_task_17` | `2.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 59.6 KiB | [pg_task_17-2.1.7-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_task_17-2.1.7-3PGDG.rhel10.x86_64.rpm) |
| `pg_task_17` | `2.1.29` | [el10.aarch64](/os/el10.aarch64) | pigsty | 52.7 KiB | [pg_task_17-2.1.29-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_task_17-2.1.29-1PIGSTY.el10.aarch64.rpm) |
| `pg_task_17` | `2.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.2 KiB | [pg_task_17-2.1.7-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_task_17-2.1.7-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-task` | `2.1.29` | [d12.x86_64](/os/d12.x86_64) | pigsty | 38.4 KiB | [postgresql-17-pg-task_2.1.29-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-17-pg-task_2.1.29-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-task` | `2.1.29` | [d12.aarch64](/os/d12.aarch64) | pigsty | 35.1 KiB | [postgresql-17-pg-task_2.1.29-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-17-pg-task_2.1.29-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-task` | `2.1.29` | [d13.x86_64](/os/d13.x86_64) | pigsty | 38.8 KiB | [postgresql-17-pg-task_2.1.29-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-17-pg-task_2.1.29-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-task` | `2.1.29` | [d13.aarch64](/os/d13.aarch64) | pigsty | 35.3 KiB | [postgresql-17-pg-task_2.1.29-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-17-pg-task_2.1.29-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-task` | `2.1.29` | [u22.x86_64](/os/u22.x86_64) | pigsty | 42.4 KiB | [postgresql-17-pg-task_2.1.29-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-17-pg-task_2.1.29-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-task` | `2.1.29` | [u22.aarch64](/os/u22.aarch64) | pigsty | 40.9 KiB | [postgresql-17-pg-task_2.1.29-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-17-pg-task_2.1.29-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-task` | `2.1.29` | [u24.x86_64](/os/u24.x86_64) | pigsty | 41.0 KiB | [postgresql-17-pg-task_2.1.29-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-17-pg-task_2.1.29-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-task` | `2.1.29` | [u24.aarch64](/os/u24.aarch64) | pigsty | 39.7 KiB | [postgresql-17-pg-task_2.1.29-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-17-pg-task_2.1.29-2PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-task` | `2.1.29` | [u26.x86_64](/os/u26.x86_64) | pigsty | 41.0 KiB | [postgresql-17-pg-task_2.1.29-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-task/postgresql-17-pg-task_2.1.29-2PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-task` | `2.1.29` | [u26.aarch64](/os/u26.aarch64) | pigsty | 39.6 KiB | [postgresql-17-pg-task_2.1.29-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-task/postgresql-17-pg-task_2.1.29-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_task_16` | `2.1.29` | [el8.x86_64](/os/el8.x86_64) | pigsty | 54.8 KiB | [pg_task_16-2.1.29-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_task_16-2.1.29-1PIGSTY.el8.x86_64.rpm) |
| `pg_task_16` | `2.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 72.3 KiB | [pg_task_16-2.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_task_16-2.1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_task_16` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 72.2 KiB | [pg_task_16-2.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_task_16-2.1.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_task_16` | `2.1.29` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.7 KiB | [pg_task_16-2.1.29-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_task_16-2.1.29-1PIGSTY.el8.aarch64.rpm) |
| `pg_task_16` | `2.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.1 KiB | [pg_task_16-2.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_task_16-2.1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_task_16` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.0 KiB | [pg_task_16-2.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_task_16-2.1.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_task_16` | `2.1.29` | [el9.x86_64](/os/el9.x86_64) | pigsty | 54.7 KiB | [pg_task_16-2.1.29-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_task_16-2.1.29-1PIGSTY.el9.x86_64.rpm) |
| `pg_task_16` | `2.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 62.8 KiB | [pg_task_16-2.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_task_16-2.1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_task_16` | `2.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 62.8 KiB | [pg_task_16-2.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_task_16-2.1.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_task_16` | `2.1.29` | [el9.aarch64](/os/el9.aarch64) | pigsty | 52.7 KiB | [pg_task_16-2.1.29-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_task_16-2.1.29-1PIGSTY.el9.aarch64.rpm) |
| `pg_task_16` | `2.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.8 KiB | [pg_task_16-2.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_task_16-2.1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_task_16` | `2.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.7 KiB | [pg_task_16-2.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_task_16-2.1.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_task_16` | `2.1.29` | [el10.x86_64](/os/el10.x86_64) | pigsty | 54.8 KiB | [pg_task_16-2.1.29-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_task_16-2.1.29-1PIGSTY.el10.x86_64.rpm) |
| `pg_task_16` | `2.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 58.8 KiB | [pg_task_16-2.1.7-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_task_16-2.1.7-3PGDG.rhel10.x86_64.rpm) |
| `pg_task_16` | `2.1.29` | [el10.aarch64](/os/el10.aarch64) | pigsty | 52.7 KiB | [pg_task_16-2.1.29-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_task_16-2.1.29-1PIGSTY.el10.aarch64.rpm) |
| `pg_task_16` | `2.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.0 KiB | [pg_task_16-2.1.7-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_task_16-2.1.7-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-task` | `2.1.29` | [d12.x86_64](/os/d12.x86_64) | pigsty | 38.5 KiB | [postgresql-16-pg-task_2.1.29-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-16-pg-task_2.1.29-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-task` | `2.1.29` | [d12.aarch64](/os/d12.aarch64) | pigsty | 35.2 KiB | [postgresql-16-pg-task_2.1.29-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-16-pg-task_2.1.29-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-task` | `2.1.29` | [d13.x86_64](/os/d13.x86_64) | pigsty | 38.7 KiB | [postgresql-16-pg-task_2.1.29-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-16-pg-task_2.1.29-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-task` | `2.1.29` | [d13.aarch64](/os/d13.aarch64) | pigsty | 35.3 KiB | [postgresql-16-pg-task_2.1.29-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-16-pg-task_2.1.29-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-task` | `2.1.29` | [u22.x86_64](/os/u22.x86_64) | pigsty | 42.4 KiB | [postgresql-16-pg-task_2.1.29-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-16-pg-task_2.1.29-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-task` | `2.1.29` | [u22.aarch64](/os/u22.aarch64) | pigsty | 40.9 KiB | [postgresql-16-pg-task_2.1.29-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-16-pg-task_2.1.29-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-task` | `2.1.29` | [u24.x86_64](/os/u24.x86_64) | pigsty | 41.0 KiB | [postgresql-16-pg-task_2.1.29-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-16-pg-task_2.1.29-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-task` | `2.1.29` | [u24.aarch64](/os/u24.aarch64) | pigsty | 39.6 KiB | [postgresql-16-pg-task_2.1.29-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-16-pg-task_2.1.29-2PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-task` | `2.1.29` | [u26.x86_64](/os/u26.x86_64) | pigsty | 41.0 KiB | [postgresql-16-pg-task_2.1.29-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-task/postgresql-16-pg-task_2.1.29-2PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-task` | `2.1.29` | [u26.aarch64](/os/u26.aarch64) | pigsty | 39.6 KiB | [postgresql-16-pg-task_2.1.29-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-task/postgresql-16-pg-task_2.1.29-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_task_15` | `2.1.29` | [el8.x86_64](/os/el8.x86_64) | pigsty | 55.9 KiB | [pg_task_15-2.1.29-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_task_15-2.1.29-1PIGSTY.el8.x86_64.rpm) |
| `pg_task_15` | `2.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 73.3 KiB | [pg_task_15-2.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_task_15-2.1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_task_15` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 73.2 KiB | [pg_task_15-2.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_task_15-2.1.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_task_15` | `2.1.29` | [el8.aarch64](/os/el8.aarch64) | pigsty | 50.8 KiB | [pg_task_15-2.1.29-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_task_15-2.1.29-1PIGSTY.el8.aarch64.rpm) |
| `pg_task_15` | `2.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.0 KiB | [pg_task_15-2.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_task_15-2.1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_task_15` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.9 KiB | [pg_task_15-2.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_task_15-2.1.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_task_15` | `2.1.29` | [el9.x86_64](/os/el9.x86_64) | pigsty | 56.0 KiB | [pg_task_15-2.1.29-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_task_15-2.1.29-1PIGSTY.el9.x86_64.rpm) |
| `pg_task_15` | `2.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.3 KiB | [pg_task_15-2.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_task_15-2.1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_task_15` | `2.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.2 KiB | [pg_task_15-2.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_task_15-2.1.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_task_15` | `2.1.29` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.0 KiB | [pg_task_15-2.1.29-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_task_15-2.1.29-1PIGSTY.el9.aarch64.rpm) |
| `pg_task_15` | `2.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.3 KiB | [pg_task_15-2.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_task_15-2.1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_task_15` | `2.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.2 KiB | [pg_task_15-2.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_task_15-2.1.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_task_15` | `2.1.29` | [el10.x86_64](/os/el10.x86_64) | pigsty | 56.2 KiB | [pg_task_15-2.1.29-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_task_15-2.1.29-1PIGSTY.el10.x86_64.rpm) |
| `pg_task_15` | `2.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 72.3 KiB | [pg_task_15-2.1.7-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_task_15-2.1.7-3PGDG.rhel10.x86_64.rpm) |
| `pg_task_15` | `2.1.29` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.1 KiB | [pg_task_15-2.1.29-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_task_15-2.1.29-1PIGSTY.el10.aarch64.rpm) |
| `pg_task_15` | `2.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 69.5 KiB | [pg_task_15-2.1.7-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_task_15-2.1.7-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-task` | `2.1.29` | [d12.x86_64](/os/d12.x86_64) | pigsty | 39.7 KiB | [postgresql-15-pg-task_2.1.29-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-15-pg-task_2.1.29-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-task` | `2.1.29` | [d12.aarch64](/os/d12.aarch64) | pigsty | 36.1 KiB | [postgresql-15-pg-task_2.1.29-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-15-pg-task_2.1.29-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-task` | `2.1.29` | [d13.x86_64](/os/d13.x86_64) | pigsty | 39.9 KiB | [postgresql-15-pg-task_2.1.29-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-15-pg-task_2.1.29-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-task` | `2.1.29` | [d13.aarch64](/os/d13.aarch64) | pigsty | 36.4 KiB | [postgresql-15-pg-task_2.1.29-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-15-pg-task_2.1.29-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-task` | `2.1.29` | [u22.x86_64](/os/u22.x86_64) | pigsty | 43.3 KiB | [postgresql-15-pg-task_2.1.29-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-15-pg-task_2.1.29-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-task` | `2.1.29` | [u22.aarch64](/os/u22.aarch64) | pigsty | 41.8 KiB | [postgresql-15-pg-task_2.1.29-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-15-pg-task_2.1.29-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-task` | `2.1.29` | [u24.x86_64](/os/u24.x86_64) | pigsty | 41.9 KiB | [postgresql-15-pg-task_2.1.29-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-15-pg-task_2.1.29-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-task` | `2.1.29` | [u24.aarch64](/os/u24.aarch64) | pigsty | 40.5 KiB | [postgresql-15-pg-task_2.1.29-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-15-pg-task_2.1.29-2PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-task` | `2.1.29` | [u26.x86_64](/os/u26.x86_64) | pigsty | 42.1 KiB | [postgresql-15-pg-task_2.1.29-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-task/postgresql-15-pg-task_2.1.29-2PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-task` | `2.1.29` | [u26.aarch64](/os/u26.aarch64) | pigsty | 40.9 KiB | [postgresql-15-pg-task_2.1.29-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-task/postgresql-15-pg-task_2.1.29-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_task_14` | `2.1.29` | [el8.x86_64](/os/el8.x86_64) | pigsty | 55.8 KiB | [pg_task_14-2.1.29-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_task_14-2.1.29-1PIGSTY.el8.x86_64.rpm) |
| `pg_task_14` | `2.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 73.0 KiB | [pg_task_14-2.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_task_14-2.1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_task_14` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 72.9 KiB | [pg_task_14-2.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_task_14-2.1.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_task_14` | `2.1.29` | [el8.aarch64](/os/el8.aarch64) | pigsty | 50.8 KiB | [pg_task_14-2.1.29-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_task_14-2.1.29-1PIGSTY.el8.aarch64.rpm) |
| `pg_task_14` | `2.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.8 KiB | [pg_task_14-2.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_task_14-2.1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_task_14` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.7 KiB | [pg_task_14-2.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_task_14-2.1.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_task_14` | `2.1.29` | [el9.x86_64](/os/el9.x86_64) | pigsty | 56.0 KiB | [pg_task_14-2.1.29-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_task_14-2.1.29-1PIGSTY.el9.x86_64.rpm) |
| `pg_task_14` | `2.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.9 KiB | [pg_task_14-2.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_task_14-2.1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_task_14` | `2.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.9 KiB | [pg_task_14-2.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_task_14-2.1.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_task_14` | `2.1.29` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.1 KiB | [pg_task_14-2.1.29-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_task_14-2.1.29-1PIGSTY.el9.aarch64.rpm) |
| `pg_task_14` | `2.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.2 KiB | [pg_task_14-2.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_task_14-2.1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_task_14` | `2.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.1 KiB | [pg_task_14-2.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_task_14-2.1.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_task_14` | `2.1.29` | [el10.x86_64](/os/el10.x86_64) | pigsty | 56.1 KiB | [pg_task_14-2.1.29-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_task_14-2.1.29-1PIGSTY.el10.x86_64.rpm) |
| `pg_task_14` | `2.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 72.1 KiB | [pg_task_14-2.1.7-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_task_14-2.1.7-3PGDG.rhel10.x86_64.rpm) |
| `pg_task_14` | `2.1.29` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.1 KiB | [pg_task_14-2.1.29-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_task_14-2.1.29-1PIGSTY.el10.aarch64.rpm) |
| `pg_task_14` | `2.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 69.4 KiB | [pg_task_14-2.1.7-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_task_14-2.1.7-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-task` | `2.1.29` | [d12.x86_64](/os/d12.x86_64) | pigsty | 39.6 KiB | [postgresql-14-pg-task_2.1.29-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-14-pg-task_2.1.29-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-task` | `2.1.29` | [d12.aarch64](/os/d12.aarch64) | pigsty | 36.1 KiB | [postgresql-14-pg-task_2.1.29-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-14-pg-task_2.1.29-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-task` | `2.1.29` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.1 KiB | [postgresql-14-pg-task_2.1.29-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-14-pg-task_2.1.29-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-task` | `2.1.29` | [d13.aarch64](/os/d13.aarch64) | pigsty | 36.2 KiB | [postgresql-14-pg-task_2.1.29-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-14-pg-task_2.1.29-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-task` | `2.1.29` | [u22.x86_64](/os/u22.x86_64) | pigsty | 43.2 KiB | [postgresql-14-pg-task_2.1.29-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-14-pg-task_2.1.29-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-task` | `2.1.29` | [u22.aarch64](/os/u22.aarch64) | pigsty | 41.7 KiB | [postgresql-14-pg-task_2.1.29-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-14-pg-task_2.1.29-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-task` | `2.1.29` | [u24.x86_64](/os/u24.x86_64) | pigsty | 41.8 KiB | [postgresql-14-pg-task_2.1.29-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-14-pg-task_2.1.29-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-task` | `2.1.29` | [u24.aarch64](/os/u24.aarch64) | pigsty | 40.5 KiB | [postgresql-14-pg-task_2.1.29-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-14-pg-task_2.1.29-2PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-task` | `2.1.29` | [u26.x86_64](/os/u26.x86_64) | pigsty | 42.1 KiB | [postgresql-14-pg-task_2.1.29-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-task/postgresql-14-pg-task_2.1.29-2PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-task` | `2.1.29` | [u26.aarch64](/os/u26.aarch64) | pigsty | 40.9 KiB | [postgresql-14-pg-task_2.1.29-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-task/postgresql-14-pg-task_2.1.29-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/RekGRpth/pg_task" title="Repository" icon="github" subtitle="github.com/RekGRpth/pg_task" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_task-2.1.29.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_task;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_task;		# install via package name, for the active PG version

pig install pg_task -v 18;   # install for PG 18
pig install pg_task -v 17;   # install for PG 17
pig install pg_task -v 16;   # install for PG 16
pig install pg_task -v 15;   # install for PG 15
pig install pg_task -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_task';
```


This extension does not need `CREATE EXTENSION` DDL command






## Usage

> Sources: [pg_task upstream README](https://github.com/RekGRpth/pg_task), [PGXN pg_task](https://pgxn.org/dist/pg_task/), [local metadata](../db/extension.csv).

`pg_task` is a background-worker scheduler for running SQL asynchronously at a planned time. Upstream documents PostgreSQL, Greenplum, and Greengage support.

Enable the worker at server start, then create the extension in the database that will own the task table:

```conf
shared_preload_libraries = 'pg_task'
```

```sql
CREATE EXTENSION pg_task;
```

### Schedule Tasks

Schedule work by inserting rows into the configured task table, which defaults to `public.task` in database `postgres` unless changed with GUCs.

```sql
-- Run SQL immediately
INSERT INTO task (input) VALUES ('SELECT now()');

-- Run SQL after 5 minutes
INSERT INTO task (plan, input) VALUES (now() + '5 min'::interval, 'SELECT now()');

-- Run SQL at a specific time
INSERT INTO task (plan, input) VALUES ('2029-07-01 12:51:00', 'SELECT now()');

-- Repeat SQL every 5 minutes
INSERT INTO task (repeat, input) VALUES ('5 min', 'SELECT now()');

-- Exceptions are caught and written to the error column
INSERT INTO task (input) VALUES ('SELECT 1/0');

-- Limit concurrent tasks in a group.
-- max = 1 means one task at a time for this group.
INSERT INTO task ("group", max, input) VALUES ('billing', 1, 'SELECT refresh_billing_cache()');

-- Run SQL on a remote database
INSERT INTO task (input, remote) VALUES ('SELECT now()', 'user=user host=host');
```

### Task Table

The task table is meant to be user-visible. Upstream notes that users may add columns or partition it.

Key columns:

| Name | Type | Default | Description |
|------|------|---------|-------------|
| id | bigserial | autoincrement | Primary key |
| parent | bigint | pg_task.id | Parent task id |
| plan | timestamptz | statement_timestamp() | Planned start time |
| start | timestamptz | | Actual start time |
| stop | timestamptz | | Actual stop time |
| active | interval | 1 hour | Period after plan time when task is active |
| live | interval | 0 sec | Max lifetime of background worker |
| repeat | interval | 0 sec | Auto repeat interval |
| timeout | interval | 0 sec | Allowed time for task run |
| count | int | 0 | Max task count before worker exit |
| max | int | 0 | Max concurrent tasks in group; negative values mean pause between tasks in milliseconds |
| pid | int | | Process id executing task |
| state | enum | PLAN | PLAN, TAKE, WORK, DONE, STOP |
| delete | bool | true | Auto delete when output and error are null |
| drift | bool | false | Compute next repeat by stop time |
| header | bool | true | Show column headers in output |
| group | text | 'group' | Task grouping name |
| input | text | | SQL command(s) to execute |
| output | text | | Received result(s) |
| error | text | | Caught error |
| remote | text | | Remote database connection string |

### Configuration

Key settings:

| Name | Type | Default | Description |
|------|------|---------|-------------|
| pg_task.delete | bool | true | Auto delete completed tasks |
| pg_task.drift | bool | false | Compute repeat by stop time |
| pg_task.header | bool | true | Show column headers in task output |
| pg_task.string | bool | true | Quote only strings in output |
| pg_task.count | int | 0 | Default maximum number of tasks per worker before exit |
| pg_task.fetch | int | 100 | Number of task rows fetched at once |
| pg_task.limit | int | 1000 | Limit task rows at once |
| pg_task.max | int | 0 | Default max concurrent tasks in group |
| pg_task.run | int | 2147483647 | Maximum concurrently executing tasks in work |
| pg_task.sleep | int | 1000 | Check tasks every N milliseconds |
| pg_task.active | interval | 1 hour | Period after plan time when a task remains active for execution |
| pg_task.live | interval | 0 sec | Maximum lifetime of a worker process |
| pg_task.repeat | interval | 0 sec | Default repeat interval |
| pg_task.timeout | interval | 0 sec | Default task timeout |
| pg_task.data | text | postgres | Database name for tasks table |
| pg_task.user | text | postgres | User name for tasks table |
| pg_task.schema | text | public | Schema name for tasks table |
| pg_task.table | text | task | Table name for tasks table |
| pg_task.json | json | `[{"data":"postgres"}]` | Multi-database configuration |
| pg_task.id | bigint | 0 | Current task id, read-only session setting |

### Multi-Database Configuration

To run tasks on multiple databases, configure via JSON:

```conf
pg_task.json = '[{"data":"database1"},{"data":"database2","user":"username2"},{"data":"database3","schema":"schema3"}]'
```

If the specified database, user, schema or table does not exist, `pg_task` will create them.

The local metadata marks this package as `headless`, so it needs `shared_preload_libraries` rather than a user-facing SQL-only install path. Validate background scheduling behavior on the target PostgreSQL version before relying on it for critical jobs.
