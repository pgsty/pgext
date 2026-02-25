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
| **1080** | {{< badge content="pg_task" link="https://github.com/RekGRpth/pg_task" >}} | {{< ext "pg_task" >}} | `1.0.0` | {{< category "TIME" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


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
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_task` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.1.7` | {{< bg "18" "pg_task_18" "green" >}} {{< bg "17" "pg_task_17" "green" >}} {{< bg "16" "pg_task_16" "green" >}} {{< bg "15" "pg_task_15" "green" >}} {{< bg "14" "pg_task_14" "green" >}} {{< bg "13" "pg_task_13" "green" >}} | `pg_task_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.12` | {{< bg "18" "postgresql-18-pg-task" "green" >}} {{< bg "17" "postgresql-17-pg-task" "green" >}} {{< bg "16" "postgresql-16-pg-task" "green" >}} {{< bg "15" "postgresql-15-pg-task" "green" >}} {{< bg "14" "postgresql-14-pg-task" "green" >}} {{< bg "13" "postgresql-13-pg-task" "green" >}} | `postgresql-$v-pg-task` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.1.7" "pg_task_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_13 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.1.7" "pg_task_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.1.7" "pg_task_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.1.7" "pg_task_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.1.7" "pg_task_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.1.7" "pg_task_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "pg_task_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-18-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-17-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-16-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-15-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-14-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-13-pg-task : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-18-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-17-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-16-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-15-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-14-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-13-pg-task : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-18-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-17-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-16-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-15-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-14-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-13-pg-task : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-18-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-17-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-16-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-15-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-14-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-13-pg-task : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-18-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-17-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-16-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-15-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-14-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-13-pg-task : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-18-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-17-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-16-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-15-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-14-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-13-pg-task : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-18-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-17-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-16-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-15-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-14-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-13-pg-task : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-18-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-17-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-16-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-15-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-14-pg-task : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.12" "postgresql-13-pg-task : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_task_18` | `2.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 72.4 KiB | [pg_task_18-2.1.7-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_task_18-2.1.7-3PGDG.rhel8.x86_64.rpm) |
| `pg_task_18` | `2.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.3 KiB | [pg_task_18-2.1.7-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_task_18-2.1.7-3PGDG.rhel8.aarch64.rpm) |
| `pg_task_18` | `2.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 63.5 KiB | [pg_task_18-2.1.7-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_task_18-2.1.7-3PGDG.rhel9.x86_64.rpm) |
| `pg_task_18` | `2.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.5 KiB | [pg_task_18-2.1.7-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_task_18-2.1.7-3PGDG.rhel9.aarch64.rpm) |
| `pg_task_18` | `2.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 59.7 KiB | [pg_task_18-2.1.7-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_task_18-2.1.7-3PGDG.rhel10.x86_64.rpm) |
| `pg_task_18` | `2.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.3 KiB | [pg_task_18-2.1.7-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_task_18-2.1.7-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-task` | `2.1.12` | [d12.x86_64](/os/d12.x86_64) | pigsty | 192.2 KiB | [postgresql-18-pg-task_2.1.12-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-18-pg-task_2.1.12-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-task` | `2.1.12` | [d12.aarch64](/os/d12.aarch64) | pigsty | 182.8 KiB | [postgresql-18-pg-task_2.1.12-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-18-pg-task_2.1.12-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-task` | `2.1.12` | [d13.x86_64](/os/d13.x86_64) | pigsty | 191.6 KiB | [postgresql-18-pg-task_2.1.12-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-18-pg-task_2.1.12-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-task` | `2.1.12` | [d13.aarch64](/os/d13.aarch64) | pigsty | 184.7 KiB | [postgresql-18-pg-task_2.1.12-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-18-pg-task_2.1.12-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-task` | `2.1.12` | [u22.x86_64](/os/u22.x86_64) | pigsty | 201.0 KiB | [postgresql-18-pg-task_2.1.12-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-18-pg-task_2.1.12-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-task` | `2.1.12` | [u22.aarch64](/os/u22.aarch64) | pigsty | 190.4 KiB | [postgresql-18-pg-task_2.1.12-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-18-pg-task_2.1.12-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-task` | `2.1.12` | [u24.x86_64](/os/u24.x86_64) | pigsty | 193.1 KiB | [postgresql-18-pg-task_2.1.12-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-18-pg-task_2.1.12-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-task` | `2.1.12` | [u24.aarch64](/os/u24.aarch64) | pigsty | 183.7 KiB | [postgresql-18-pg-task_2.1.12-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-18-pg-task_2.1.12-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_task_17` | `2.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 72.5 KiB | [pg_task_17-2.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_task_17-2.1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_task_17` | `2.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.4 KiB | [pg_task_17-2.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_task_17-2.1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_task_17` | `2.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 63.3 KiB | [pg_task_17-2.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_task_17-2.1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_task_17` | `2.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.4 KiB | [pg_task_17-2.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_task_17-2.1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_task_17` | `2.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 59.6 KiB | [pg_task_17-2.1.7-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_task_17-2.1.7-3PGDG.rhel10.x86_64.rpm) |
| `pg_task_17` | `2.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.2 KiB | [pg_task_17-2.1.7-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_task_17-2.1.7-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-task` | `2.1.12` | [d12.x86_64](/os/d12.x86_64) | pigsty | 192.8 KiB | [postgresql-17-pg-task_2.1.12-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-17-pg-task_2.1.12-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-task` | `2.1.12` | [d12.aarch64](/os/d12.aarch64) | pigsty | 183.6 KiB | [postgresql-17-pg-task_2.1.12-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-17-pg-task_2.1.12-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-task` | `2.1.12` | [d13.x86_64](/os/d13.x86_64) | pigsty | 192.0 KiB | [postgresql-17-pg-task_2.1.12-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-17-pg-task_2.1.12-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-task` | `2.1.12` | [d13.aarch64](/os/d13.aarch64) | pigsty | 185.4 KiB | [postgresql-17-pg-task_2.1.12-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-17-pg-task_2.1.12-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-task` | `2.1.12` | [u22.x86_64](/os/u22.x86_64) | pigsty | 229.8 KiB | [postgresql-17-pg-task_2.1.12-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-17-pg-task_2.1.12-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-task` | `2.1.12` | [u22.aarch64](/os/u22.aarch64) | pigsty | 219.4 KiB | [postgresql-17-pg-task_2.1.12-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-17-pg-task_2.1.12-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-task` | `2.1.12` | [u24.x86_64](/os/u24.x86_64) | pigsty | 193.9 KiB | [postgresql-17-pg-task_2.1.12-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-17-pg-task_2.1.12-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-task` | `2.1.12` | [u24.aarch64](/os/u24.aarch64) | pigsty | 184.0 KiB | [postgresql-17-pg-task_2.1.12-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-17-pg-task_2.1.12-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_task_16` | `2.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 72.3 KiB | [pg_task_16-2.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_task_16-2.1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_task_16` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 72.2 KiB | [pg_task_16-2.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_task_16-2.1.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_task_16` | `2.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.1 KiB | [pg_task_16-2.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_task_16-2.1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_task_16` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.0 KiB | [pg_task_16-2.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_task_16-2.1.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_task_16` | `2.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 62.8 KiB | [pg_task_16-2.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_task_16-2.1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_task_16` | `2.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 62.8 KiB | [pg_task_16-2.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_task_16-2.1.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_task_16` | `2.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.8 KiB | [pg_task_16-2.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_task_16-2.1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_task_16` | `2.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.7 KiB | [pg_task_16-2.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_task_16-2.1.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_task_16` | `2.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 58.8 KiB | [pg_task_16-2.1.7-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_task_16-2.1.7-3PGDG.rhel10.x86_64.rpm) |
| `pg_task_16` | `2.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.0 KiB | [pg_task_16-2.1.7-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_task_16-2.1.7-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-task` | `2.1.12` | [d12.x86_64](/os/d12.x86_64) | pigsty | 192.4 KiB | [postgresql-16-pg-task_2.1.12-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-16-pg-task_2.1.12-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-task` | `2.1.12` | [d12.aarch64](/os/d12.aarch64) | pigsty | 183.5 KiB | [postgresql-16-pg-task_2.1.12-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-16-pg-task_2.1.12-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-task` | `2.1.12` | [d13.x86_64](/os/d13.x86_64) | pigsty | 191.6 KiB | [postgresql-16-pg-task_2.1.12-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-16-pg-task_2.1.12-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-task` | `2.1.12` | [d13.aarch64](/os/d13.aarch64) | pigsty | 184.8 KiB | [postgresql-16-pg-task_2.1.12-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-16-pg-task_2.1.12-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-task` | `2.1.12` | [u22.x86_64](/os/u22.x86_64) | pigsty | 226.9 KiB | [postgresql-16-pg-task_2.1.12-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-16-pg-task_2.1.12-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-task` | `2.1.12` | [u22.aarch64](/os/u22.aarch64) | pigsty | 216.6 KiB | [postgresql-16-pg-task_2.1.12-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-16-pg-task_2.1.12-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-task` | `2.1.12` | [u24.x86_64](/os/u24.x86_64) | pigsty | 193.1 KiB | [postgresql-16-pg-task_2.1.12-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-16-pg-task_2.1.12-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-task` | `2.1.12` | [u24.aarch64](/os/u24.aarch64) | pigsty | 183.2 KiB | [postgresql-16-pg-task_2.1.12-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-16-pg-task_2.1.12-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_task_15` | `2.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 73.3 KiB | [pg_task_15-2.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_task_15-2.1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_task_15` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 73.2 KiB | [pg_task_15-2.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_task_15-2.1.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_task_15` | `2.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.0 KiB | [pg_task_15-2.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_task_15-2.1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_task_15` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.9 KiB | [pg_task_15-2.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_task_15-2.1.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_task_15` | `2.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.3 KiB | [pg_task_15-2.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_task_15-2.1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_task_15` | `2.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.2 KiB | [pg_task_15-2.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_task_15-2.1.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_task_15` | `2.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.3 KiB | [pg_task_15-2.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_task_15-2.1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_task_15` | `2.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.2 KiB | [pg_task_15-2.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_task_15-2.1.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_task_15` | `2.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 72.3 KiB | [pg_task_15-2.1.7-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_task_15-2.1.7-3PGDG.rhel10.x86_64.rpm) |
| `pg_task_15` | `2.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 69.5 KiB | [pg_task_15-2.1.7-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_task_15-2.1.7-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-task` | `2.1.12` | [d12.x86_64](/os/d12.x86_64) | pigsty | 193.0 KiB | [postgresql-15-pg-task_2.1.12-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-15-pg-task_2.1.12-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-task` | `2.1.12` | [d12.aarch64](/os/d12.aarch64) | pigsty | 183.3 KiB | [postgresql-15-pg-task_2.1.12-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-15-pg-task_2.1.12-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-task` | `2.1.12` | [d13.x86_64](/os/d13.x86_64) | pigsty | 192.1 KiB | [postgresql-15-pg-task_2.1.12-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-15-pg-task_2.1.12-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-task` | `2.1.12` | [d13.aarch64](/os/d13.aarch64) | pigsty | 184.7 KiB | [postgresql-15-pg-task_2.1.12-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-15-pg-task_2.1.12-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-task` | `2.1.12` | [u22.x86_64](/os/u22.x86_64) | pigsty | 236.4 KiB | [postgresql-15-pg-task_2.1.12-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-15-pg-task_2.1.12-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-task` | `2.1.12` | [u22.aarch64](/os/u22.aarch64) | pigsty | 229.2 KiB | [postgresql-15-pg-task_2.1.12-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-15-pg-task_2.1.12-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-task` | `2.1.12` | [u24.x86_64](/os/u24.x86_64) | pigsty | 202.5 KiB | [postgresql-15-pg-task_2.1.12-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-15-pg-task_2.1.12-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-task` | `2.1.12` | [u24.aarch64](/os/u24.aarch64) | pigsty | 195.6 KiB | [postgresql-15-pg-task_2.1.12-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-15-pg-task_2.1.12-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_task_14` | `2.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 73.0 KiB | [pg_task_14-2.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_task_14-2.1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_task_14` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 72.9 KiB | [pg_task_14-2.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_task_14-2.1.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_task_14` | `2.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.8 KiB | [pg_task_14-2.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_task_14-2.1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_task_14` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.7 KiB | [pg_task_14-2.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_task_14-2.1.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_task_14` | `2.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.9 KiB | [pg_task_14-2.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_task_14-2.1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_task_14` | `2.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.9 KiB | [pg_task_14-2.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_task_14-2.1.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_task_14` | `2.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.2 KiB | [pg_task_14-2.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_task_14-2.1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_task_14` | `2.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.1 KiB | [pg_task_14-2.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_task_14-2.1.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_task_14` | `2.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 72.1 KiB | [pg_task_14-2.1.7-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_task_14-2.1.7-3PGDG.rhel10.x86_64.rpm) |
| `pg_task_14` | `2.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 69.4 KiB | [pg_task_14-2.1.7-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_task_14-2.1.7-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-task` | `2.1.12` | [d12.x86_64](/os/d12.x86_64) | pigsty | 192.6 KiB | [postgresql-14-pg-task_2.1.12-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-14-pg-task_2.1.12-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-task` | `2.1.12` | [d12.aarch64](/os/d12.aarch64) | pigsty | 182.7 KiB | [postgresql-14-pg-task_2.1.12-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-14-pg-task_2.1.12-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-task` | `2.1.12` | [d13.x86_64](/os/d13.x86_64) | pigsty | 191.8 KiB | [postgresql-14-pg-task_2.1.12-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-14-pg-task_2.1.12-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-task` | `2.1.12` | [d13.aarch64](/os/d13.aarch64) | pigsty | 183.9 KiB | [postgresql-14-pg-task_2.1.12-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-14-pg-task_2.1.12-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-task` | `2.1.12` | [u22.x86_64](/os/u22.x86_64) | pigsty | 232.1 KiB | [postgresql-14-pg-task_2.1.12-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-14-pg-task_2.1.12-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-task` | `2.1.12` | [u22.aarch64](/os/u22.aarch64) | pigsty | 224.7 KiB | [postgresql-14-pg-task_2.1.12-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-14-pg-task_2.1.12-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-task` | `2.1.12` | [u24.x86_64](/os/u24.x86_64) | pigsty | 201.8 KiB | [postgresql-14-pg-task_2.1.12-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-14-pg-task_2.1.12-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-task` | `2.1.12` | [u24.aarch64](/os/u24.aarch64) | pigsty | 195.2 KiB | [postgresql-14-pg-task_2.1.12-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-14-pg-task_2.1.12-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_task_13` | `2.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 71.7 KiB | [pg_task_13-2.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_task_13-2.1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_task_13` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 71.6 KiB | [pg_task_13-2.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_task_13-2.1.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_task_13` | `2.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.0 KiB | [pg_task_13-2.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_task_13-2.1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_task_13` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.0 KiB | [pg_task_13-2.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_task_13-2.1.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_task_13` | `2.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.6 KiB | [pg_task_13-2.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_task_13-2.1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_task_13` | `2.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.5 KiB | [pg_task_13-2.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_task_13-2.1.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_task_13` | `2.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.7 KiB | [pg_task_13-2.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_task_13-2.1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_task_13` | `2.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.7 KiB | [pg_task_13-2.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_task_13-2.1.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_task_13` | `2.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 71.4 KiB | [pg_task_13-2.1.7-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_task_13-2.1.7-3PGDG.rhel10.x86_64.rpm) |
| `pg_task_13` | `2.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 69.0 KiB | [pg_task_13-2.1.7-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_task_13-2.1.7-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pg-task` | `2.1.12` | [d12.x86_64](/os/d12.x86_64) | pigsty | 191.3 KiB | [postgresql-13-pg-task_2.1.12-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-13-pg-task_2.1.12-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-task` | `2.1.12` | [d12.aarch64](/os/d12.aarch64) | pigsty | 180.9 KiB | [postgresql-13-pg-task_2.1.12-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-task/postgresql-13-pg-task_2.1.12-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-task` | `2.1.12` | [d13.x86_64](/os/d13.x86_64) | pigsty | 190.2 KiB | [postgresql-13-pg-task_2.1.12-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-13-pg-task_2.1.12-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-task` | `2.1.12` | [d13.aarch64](/os/d13.aarch64) | pigsty | 182.5 KiB | [postgresql-13-pg-task_2.1.12-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-task/postgresql-13-pg-task_2.1.12-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-task` | `2.1.12` | [u22.x86_64](/os/u22.x86_64) | pigsty | 229.1 KiB | [postgresql-13-pg-task_2.1.12-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-13-pg-task_2.1.12-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-task` | `2.1.12` | [u22.aarch64](/os/u22.aarch64) | pigsty | 221.5 KiB | [postgresql-13-pg-task_2.1.12-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-task/postgresql-13-pg-task_2.1.12-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-task` | `2.1.12` | [u24.x86_64](/os/u24.x86_64) | pigsty | 200.1 KiB | [postgresql-13-pg-task_2.1.12-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-13-pg-task_2.1.12-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-task` | `2.1.12` | [u24.aarch64](/os/u24.aarch64) | pigsty | 192.9 KiB | [postgresql-13-pg-task_2.1.12-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-task/postgresql-13-pg-task_2.1.12-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/RekGRpth/pg_task" title="Repository" icon="github" subtitle="github.com/RekGRpth/pg_task" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_task-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_task;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_task;		# install via package name, for the active PG version

pig install pg_task -v 18;   # install for PG 18
pig install pg_task -v 17;   # install for PG 17
pig install pg_task -v 16;   # install for PG 16
pig install pg_task -v 15;   # install for PG 15
pig install pg_task -v 14;   # install for PG 14
pig install pg_task -v 13;   # install for PG 13

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_task';
```


This extension does not need `CREATE EXTENSION` DDL command


