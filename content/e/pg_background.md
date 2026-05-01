---
title: "pg_background"
linkTitle: "pg_background"
description: "Run SQL queries in the background"
weight: 1110
categories: ["TIME"]
width: full
---

[**pg_background**](https://github.com/vibhorkum/pg_background) : Run SQL queries in the background


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1110** | {{< badge content="pg_background" link="https://github.com/vibhorkum/pg_background" >}} | {{< ext "pg_background" >}} | `1.9.2` | {{< category "TIME" >}} | {{< license "GPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_cron" >}} {{< ext "pg_task" >}} {{< ext "pg_later" >}} {{< ext "pgq" >}} {{< ext "timescaledb" >}} {{< ext "timescaledb_toolkit" >}} {{< ext "timeseries" >}} {{< ext "periods" >}} |

> [!Note] Release tag 1.9.2 still ships extension SQL version 1.9.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.9.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_background` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.9.3` | {{< bg "18" "pg_background_18" "green" >}} {{< bg "17" "pg_background_17" "green" >}} {{< bg "16" "pg_background_16" "green" >}} {{< bg "15" "pg_background_15" "green" >}} {{< bg "14" "pg_background_14" "green" >}} | `pg_background_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.9.3` | {{< bg "18" "postgresql-18-pg-background" "green" >}} {{< bg "17" "postgresql-17-pg-background" "green" >}} {{< bg "16" "postgresql-16-pg-background" "green" >}} {{< bg "15" "postgresql-15-pg-background" "green" >}} {{< bg "14" "postgresql-14-pg-background" "green" >}} | `postgresql-$v-pg-background` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.9.3" "pg_background_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_15 : AVAIL 8" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_14 : AVAIL 7" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.9.3" "pg_background_18 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_17 : AVAIL 8" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_16 : AVAIL 8" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_15 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_14 : AVAIL 8" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.9.3" "pg_background_18 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_17 : AVAIL 8" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_16 : AVAIL 8" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_15 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_14 : AVAIL 8" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.9.3" "pg_background_18 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_17 : AVAIL 8" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_16 : AVAIL 8" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_15 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_14 : AVAIL 8" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.9.3" "pg_background_18 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_14 : AVAIL 7" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.9.3" "pg_background_18 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.3" "pg_background_14 : AVAIL 7" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.9.2" "postgresql-18-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-17-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-16-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-15-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-14-pg-background : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.9.2" "postgresql-18-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-17-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-16-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-15-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-14-pg-background : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.9.2" "postgresql-18-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-17-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-16-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-15-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-14-pg-background : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.9.2" "postgresql-18-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-17-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-16-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-15-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-14-pg-background : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.9.2" "postgresql-18-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-17-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-16-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-15-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-14-pg-background : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.9.2" "postgresql-18-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-17-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-16-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-15-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-14-pg-background : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.9.2" "postgresql-18-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-17-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-16-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-15-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-14-pg-background : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.9.2" "postgresql-18-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-17-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-16-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-15-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-14-pg-background : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.9.2" "postgresql-18-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-17-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-16-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-15-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-14-pg-background : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.9.2" "postgresql-18-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-17-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-16-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-15-pg-background : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.2" "postgresql-14-pg-background : AVAIL 3" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_18` | `1.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.5 KiB | [pg_background_18-1.9.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.9.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 56.3 KiB | [pg_background_18-1.9.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_background_18-1.9.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_background_18` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.4 KiB | [pg_background_18-1.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.1 KiB | [pg_background_18-1.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pg_background_18-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.5 KiB | [pg_background_18-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_18` | `1.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 57.3 KiB | [pg_background_18-1.9.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.9.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 55.2 KiB | [pg_background_18-1.9.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_background_18-1.9.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_background_18` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.1 KiB | [pg_background_18-1.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pg_background_18-1.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.8 KiB | [pg_background_18-1.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_background_18-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.0 KiB | [pg_background_18-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_18` | `1.9.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.4 KiB | [pg_background_18-1.9.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_background_18-1.9.3-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_18` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.3 KiB | [pg_background_18-1.9.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_background_18-1.9.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_background_18` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 54.3 KiB | [pg_background_18-1.9.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_background_18-1.9.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_18` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 54.0 KiB | [pg_background_18-1.9.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_background_18-1.9.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_18` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.5 KiB | [pg_background_18-1.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_background_18-1.8-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_18` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.3 KiB | [pg_background_18-1.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_background_18-1.6-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_18` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.7 KiB | [pg_background_18-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_background_18-1.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_18` | `1.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 56.4 KiB | [pg_background_18-1.9.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_background_18-1.9.3-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_18` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.3 KiB | [pg_background_18-1.9.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_background_18-1.9.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_background_18` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.4 KiB | [pg_background_18-1.9.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_background_18-1.9.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_18` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.1 KiB | [pg_background_18-1.9.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_background_18-1.9.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_18` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.4 KiB | [pg_background_18-1.8-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_background_18-1.8-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_18` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [pg_background_18-1.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_background_18-1.6-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_18` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.0 KiB | [pg_background_18-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_background_18-1.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_18` | `1.9.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.6 KiB | [pg_background_18-1.9.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_background_18-1.9.3-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_18` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.5 KiB | [pg_background_18-1.9.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_background_18-1.9.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_background_18` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.5 KiB | [pg_background_18-1.9.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_background_18-1.9.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_18` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.2 KiB | [pg_background_18-1.9.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_background_18-1.9.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_18` | `1.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.8 KiB | [pg_background_18-1.8-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_background_18-1.8-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_18` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.6 KiB | [pg_background_18-1.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_background_18-1.6-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_18` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.1 KiB | [pg_background_18-1.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_background_18-1.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_background_18` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_18-1.9.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.9.3-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_18` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.7 KiB | [pg_background_18-1.9.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_background_18-1.9.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_background_18` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.8 KiB | [pg_background_18-1.9.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.9.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_18` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.5 KiB | [pg_background_18-1.9.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.9.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_18` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_background_18-1.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.8-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_18` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_background_18-1.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.6-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_18` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.6 KiB | [pg_background_18-1.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 58.9 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 89.2 KiB | [postgresql-18-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 58.6 KiB | [postgresql-18-pg-background_1.9.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.1-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 57.0 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 87.4 KiB | [postgresql-18-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 56.8 KiB | [postgresql-18-pg-background_1.9.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.1-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 58.7 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 89.3 KiB | [postgresql-18-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 58.7 KiB | [postgresql-18-pg-background_1.9.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.1-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 57.3 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 87.9 KiB | [postgresql-18-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 57.2 KiB | [postgresql-18-pg-background_1.9.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.1-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 59.7 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 94.9 KiB | [postgresql-18-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 59.8 KiB | [postgresql-18-pg-background_1.9.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 58.0 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 93.6 KiB | [postgresql-18-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 58.1 KiB | [postgresql-18-pg-background_1.9.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 58.8 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 92.0 KiB | [postgresql-18-pg-background_1.9.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 58.7 KiB | [postgresql-18-pg-background_1.9.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 57.2 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 91.1 KiB | [postgresql-18-pg-background_1.9.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 57.1 KiB | [postgresql-18-pg-background_1.9.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 58.4 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 91.4 KiB | [postgresql-18-pg-background_1.9.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 58.3 KiB | [postgresql-18-pg-background_1.9.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 56.3 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 89.9 KiB | [postgresql-18-pg-background_1.9.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1PIGSTY~resolute_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 56.2 KiB | [postgresql-18-pg-background_1.9.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.1-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_17` | `1.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.5 KiB | [pg_background_17-1.9.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.9.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 56.3 KiB | [pg_background_17-1.9.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_background_17-1.9.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_background_17` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.4 KiB | [pg_background_17-1.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.1 KiB | [pg_background_17-1.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [pg_background_17-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.9 KiB | [pg_background_17-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_17` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.1 KiB | [pg_background_17-1.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.2-2PGDG.rhel8.x86_64.rpm) |
| `pg_background_17` | `1.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 57.2 KiB | [pg_background_17-1.9.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.9.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 55.1 KiB | [pg_background_17-1.9.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_background_17-1.9.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_background_17` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.1 KiB | [pg_background_17-1.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pg_background_17-1.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.8 KiB | [pg_background_17-1.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_background_17-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.3 KiB | [pg_background_17-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_17` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.6 KiB | [pg_background_17-1.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.2-2PGDG.rhel8.aarch64.rpm) |
| `pg_background_17` | `1.9.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.3 KiB | [pg_background_17-1.9.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_background_17-1.9.3-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_17` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.5 KiB | [pg_background_17-1.9.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_background_17-1.9.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_background_17` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 54.1 KiB | [pg_background_17-1.9.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_background_17-1.9.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_17` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.9 KiB | [pg_background_17-1.9.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_background_17-1.9.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_17` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.4 KiB | [pg_background_17-1.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_background_17-1.8-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_17` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.3 KiB | [pg_background_17-1.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_background_17-1.6-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_17` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.3 KiB | [pg_background_17-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_background_17-1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_17` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.5 KiB | [pg_background_17-1.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_background_17-1.2-2PGDG.rhel9.x86_64.rpm) |
| `pg_background_17` | `1.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 56.4 KiB | [pg_background_17-1.9.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_background_17-1.9.3-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_17` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.3 KiB | [pg_background_17-1.9.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_background_17-1.9.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_background_17` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.3 KiB | [pg_background_17-1.9.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_background_17-1.9.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_17` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.1 KiB | [pg_background_17-1.9.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_background_17-1.9.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_17` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.4 KiB | [pg_background_17-1.8-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_background_17-1.8-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_17` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [pg_background_17-1.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_background_17-1.6-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_17` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.6 KiB | [pg_background_17-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_background_17-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_17` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.9 KiB | [pg_background_17-1.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_background_17-1.2-2PGDG.rhel9.aarch64.rpm) |
| `pg_background_17` | `1.9.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.5 KiB | [pg_background_17-1.9.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_background_17-1.9.3-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_17` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.4 KiB | [pg_background_17-1.9.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_background_17-1.9.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_background_17` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.3 KiB | [pg_background_17-1.9.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_background_17-1.9.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_17` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.1 KiB | [pg_background_17-1.9.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_background_17-1.9.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_17` | `1.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.8 KiB | [pg_background_17-1.8-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_background_17-1.8-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_17` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.6 KiB | [pg_background_17-1.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_background_17-1.6-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_17` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.8 KiB | [pg_background_17-1.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_background_17-1.3-3PGDG.rhel10.x86_64.rpm) |
| `pg_background_17` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_17-1.9.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.9.3-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_17` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.7 KiB | [pg_background_17-1.9.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_background_17-1.9.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_background_17` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.8 KiB | [pg_background_17-1.9.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.9.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_17` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.6 KiB | [pg_background_17-1.9.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.9.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_17` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_background_17-1.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.8-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_17` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_background_17-1.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.6-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_17` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.3 KiB | [pg_background_17-1.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.3-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 58.5 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 89.0 KiB | [postgresql-17-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 58.4 KiB | [postgresql-17-pg-background_1.9.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.1-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 57.0 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 87.5 KiB | [postgresql-17-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 56.7 KiB | [postgresql-17-pg-background_1.9.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.1-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 58.7 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 89.1 KiB | [postgresql-17-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 58.4 KiB | [postgresql-17-pg-background_1.9.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.1-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 57.1 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 87.8 KiB | [postgresql-17-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 57.1 KiB | [postgresql-17-pg-background_1.9.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.1-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 64.8 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 100.2 KiB | [postgresql-17-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 64.7 KiB | [postgresql-17-pg-background_1.9.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 63.1 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 98.6 KiB | [postgresql-17-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 62.9 KiB | [postgresql-17-pg-background_1.9.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 58.7 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 92.0 KiB | [postgresql-17-pg-background_1.9.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 58.6 KiB | [postgresql-17-pg-background_1.9.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 57.2 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 91.0 KiB | [postgresql-17-pg-background_1.9.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 57.0 KiB | [postgresql-17-pg-background_1.9.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 58.3 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 91.4 KiB | [postgresql-17-pg-background_1.9.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 58.3 KiB | [postgresql-17-pg-background_1.9.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 56.2 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 89.9 KiB | [postgresql-17-pg-background_1.9.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1PIGSTY~resolute_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 56.2 KiB | [postgresql-17-pg-background_1.9.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.1-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_16` | `1.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.5 KiB | [pg_background_16-1.9.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.9.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 56.2 KiB | [pg_background_16-1.9.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_background_16-1.9.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_background_16` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.3 KiB | [pg_background_16-1.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.1 KiB | [pg_background_16-1.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pg_background_16-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.9 KiB | [pg_background_16-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_16` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.7 KiB | [pg_background_16-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_16` | `1.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 57.2 KiB | [pg_background_16-1.9.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.9.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 55.1 KiB | [pg_background_16-1.9.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_background_16-1.9.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_background_16` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.1 KiB | [pg_background_16-1.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.8 KiB | [pg_background_16-1.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.7 KiB | [pg_background_16-1.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_background_16-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.3 KiB | [pg_background_16-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_16` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.2 KiB | [pg_background_16-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_16` | `1.9.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.3 KiB | [pg_background_16-1.9.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_background_16-1.9.3-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_16` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.2 KiB | [pg_background_16-1.9.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_background_16-1.9.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_background_16` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 54.1 KiB | [pg_background_16-1.9.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_background_16-1.9.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_16` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.9 KiB | [pg_background_16-1.9.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_background_16-1.9.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_16` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.4 KiB | [pg_background_16-1.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_background_16-1.8-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_16` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.2 KiB | [pg_background_16-1.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_background_16-1.6-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_16` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.3 KiB | [pg_background_16-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_background_16-1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_16` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.9 KiB | [pg_background_16-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_background_16-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_16` | `1.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 56.4 KiB | [pg_background_16-1.9.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_background_16-1.9.3-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_16` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.4 KiB | [pg_background_16-1.9.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_background_16-1.9.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_background_16` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.3 KiB | [pg_background_16-1.9.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_background_16-1.9.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_16` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.1 KiB | [pg_background_16-1.9.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_background_16-1.9.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_16` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.3 KiB | [pg_background_16-1.8-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_background_16-1.8-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_16` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [pg_background_16-1.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_background_16-1.6-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_16` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.6 KiB | [pg_background_16-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_background_16-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_16` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.2 KiB | [pg_background_16-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_background_16-1.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_16` | `1.9.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.5 KiB | [pg_background_16-1.9.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_background_16-1.9.3-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_16` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.4 KiB | [pg_background_16-1.9.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_background_16-1.9.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_background_16` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.3 KiB | [pg_background_16-1.9.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_background_16-1.9.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_16` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.2 KiB | [pg_background_16-1.9.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_background_16-1.9.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_16` | `1.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.7 KiB | [pg_background_16-1.8-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_background_16-1.8-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_16` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.7 KiB | [pg_background_16-1.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_background_16-1.6-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_16` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.8 KiB | [pg_background_16-1.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_background_16-1.3-3PGDG.rhel10.x86_64.rpm) |
| `pg_background_16` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_16-1.9.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.9.3-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_16` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.7 KiB | [pg_background_16-1.9.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_background_16-1.9.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_background_16` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.8 KiB | [pg_background_16-1.9.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.9.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_16` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.5 KiB | [pg_background_16-1.9.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.9.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_16` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_background_16-1.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.8-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_16` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_background_16-1.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.6-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_16` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.4 KiB | [pg_background_16-1.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.3-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 58.8 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 89.2 KiB | [postgresql-16-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 58.4 KiB | [postgresql-16-pg-background_1.9.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.1-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 56.9 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 87.4 KiB | [postgresql-16-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 56.8 KiB | [postgresql-16-pg-background_1.9.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.1-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 58.5 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 89.0 KiB | [postgresql-16-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 58.4 KiB | [postgresql-16-pg-background_1.9.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.1-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 57.2 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 87.7 KiB | [postgresql-16-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 57.1 KiB | [postgresql-16-pg-background_1.9.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.1-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 64.8 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 100.1 KiB | [postgresql-16-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 64.6 KiB | [postgresql-16-pg-background_1.9.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 63.0 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 98.5 KiB | [postgresql-16-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 62.9 KiB | [postgresql-16-pg-background_1.9.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 58.5 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 91.9 KiB | [postgresql-16-pg-background_1.9.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 58.7 KiB | [postgresql-16-pg-background_1.9.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 57.2 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 91.1 KiB | [postgresql-16-pg-background_1.9.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 57.2 KiB | [postgresql-16-pg-background_1.9.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 58.2 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 91.4 KiB | [postgresql-16-pg-background_1.9.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 58.3 KiB | [postgresql-16-pg-background_1.9.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 56.2 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 89.9 KiB | [postgresql-16-pg-background_1.9.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1PIGSTY~resolute_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 56.2 KiB | [postgresql-16-pg-background_1.9.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.1-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_15` | `1.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.5 KiB | [pg_background_15-1.9.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.9.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 56.3 KiB | [pg_background_15-1.9.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_background_15-1.9.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_background_15` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.3 KiB | [pg_background_15-1.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.1 KiB | [pg_background_15-1.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pg_background_15-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.9 KiB | [pg_background_15-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.6 KiB | [pg_background_15-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.3 KiB | [pg_background_15-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.0-1.rhel8.x86_64.rpm) |
| `pg_background_15` | `1.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 57.2 KiB | [pg_background_15-1.9.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.9.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 55.1 KiB | [pg_background_15-1.9.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_background_15-1.9.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_background_15` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.1 KiB | [pg_background_15-1.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.8 KiB | [pg_background_15-1.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.7 KiB | [pg_background_15-1.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_background_15-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.2 KiB | [pg_background_15-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.2 KiB | [pg_background_15-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.7 KiB | [pg_background_15-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.0-1.rhel8.aarch64.rpm) |
| `pg_background_15` | `1.9.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.3 KiB | [pg_background_15-1.9.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-1.9.3-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_15` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.2 KiB | [pg_background_15-1.9.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_background_15-1.9.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_background_15` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 54.1 KiB | [pg_background_15-1.9.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-1.9.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_15` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.9 KiB | [pg_background_15-1.9.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-1.9.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_15` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.4 KiB | [pg_background_15-1.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-1.8-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_15` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.2 KiB | [pg_background_15-1.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-1.6-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_15` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.2 KiB | [pg_background_15-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_15` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.9 KiB | [pg_background_15-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.6 KiB | [pg_background_15-1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-1.0-1.rhel9.x86_64.rpm) |
| `pg_background_15` | `1.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 56.3 KiB | [pg_background_15-1.9.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-1.9.3-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_15` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.3 KiB | [pg_background_15-1.9.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_background_15-1.9.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_background_15` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.3 KiB | [pg_background_15-1.9.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-1.9.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_15` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.1 KiB | [pg_background_15-1.9.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-1.9.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_15` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.4 KiB | [pg_background_15-1.8-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-1.8-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_15` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [pg_background_15-1.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-1.6-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_15` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.6 KiB | [pg_background_15-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_15` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.1 KiB | [pg_background_15-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-1.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.7 KiB | [pg_background_15-1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-1.0-1.rhel9.aarch64.rpm) |
| `pg_background_15` | `1.9.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.5 KiB | [pg_background_15-1.9.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_background_15-1.9.3-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_15` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.4 KiB | [pg_background_15-1.9.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_background_15-1.9.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_background_15` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.4 KiB | [pg_background_15-1.9.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_background_15-1.9.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_15` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.2 KiB | [pg_background_15-1.9.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_background_15-1.9.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_15` | `1.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.8 KiB | [pg_background_15-1.8-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_background_15-1.8-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_15` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.8 KiB | [pg_background_15-1.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_background_15-1.6-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_15` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.8 KiB | [pg_background_15-1.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_background_15-1.3-3PGDG.rhel10.x86_64.rpm) |
| `pg_background_15` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_15-1.9.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.9.3-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_15` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.8 KiB | [pg_background_15-1.9.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_background_15-1.9.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_background_15` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.8 KiB | [pg_background_15-1.9.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.9.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_15` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.6 KiB | [pg_background_15-1.9.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.9.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_15` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_background_15-1.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.8-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_15` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_background_15-1.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.6-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_15` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.3 KiB | [pg_background_15-1.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.3-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 58.5 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 89.1 KiB | [postgresql-15-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 58.4 KiB | [postgresql-15-pg-background_1.9.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.1-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 56.9 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 87.4 KiB | [postgresql-15-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 56.8 KiB | [postgresql-15-pg-background_1.9.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.1-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 58.5 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 89.3 KiB | [postgresql-15-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 58.5 KiB | [postgresql-15-pg-background_1.9.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.1-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 57.3 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 87.7 KiB | [postgresql-15-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 57.2 KiB | [postgresql-15-pg-background_1.9.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.1-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 64.7 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 100.1 KiB | [postgresql-15-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 64.6 KiB | [postgresql-15-pg-background_1.9.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 63.0 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 98.5 KiB | [postgresql-15-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 62.9 KiB | [postgresql-15-pg-background_1.9.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 58.8 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 91.9 KiB | [postgresql-15-pg-background_1.9.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 58.7 KiB | [postgresql-15-pg-background_1.9.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 57.3 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 91.1 KiB | [postgresql-15-pg-background_1.9.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 57.2 KiB | [postgresql-15-pg-background_1.9.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 58.3 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 91.4 KiB | [postgresql-15-pg-background_1.9.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 58.1 KiB | [postgresql-15-pg-background_1.9.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 56.4 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 89.9 KiB | [postgresql-15-pg-background_1.9.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1PIGSTY~resolute_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 56.3 KiB | [postgresql-15-pg-background_1.9.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.1-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_14` | `1.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.4 KiB | [pg_background_14-1.9.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-1.9.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_14` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 56.2 KiB | [pg_background_14-1.9.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_background_14-1.9.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_background_14` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.3 KiB | [pg_background_14-1.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-1.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_14` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.1 KiB | [pg_background_14-1.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-1.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_14` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pg_background_14-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_14` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.9 KiB | [pg_background_14-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.6 KiB | [pg_background_14-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-1.0-1.rhel8.x86_64.rpm) |
| `pg_background_14` | `1.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 57.1 KiB | [pg_background_14-1.9.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.9.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_14` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 55.1 KiB | [pg_background_14-1.9.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_background_14-1.9.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_background_14` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.0 KiB | [pg_background_14-1.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_14` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.8 KiB | [pg_background_14-1.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_14` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.6 KiB | [pg_background_14-1.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_14` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.6 KiB | [pg_background_14-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_14` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.2 KiB | [pg_background_14-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.6 KiB | [pg_background_14-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.0-1.rhel8.aarch64.rpm) |
| `pg_background_14` | `1.9.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.2 KiB | [pg_background_14-1.9.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_background_14-1.9.3-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_14` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.2 KiB | [pg_background_14-1.9.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_background_14-1.9.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_background_14` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 54.3 KiB | [pg_background_14-1.9.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_background_14-1.9.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_14` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.9 KiB | [pg_background_14-1.9.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_background_14-1.9.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_14` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.3 KiB | [pg_background_14-1.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_background_14-1.8-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_14` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.2 KiB | [pg_background_14-1.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_background_14-1.6-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_14` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.2 KiB | [pg_background_14-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_background_14-1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_14` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.9 KiB | [pg_background_14-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_background_14-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_14` | `1.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 56.3 KiB | [pg_background_14-1.9.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.9.3-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_14` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.2 KiB | [pg_background_14-1.9.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_background_14-1.9.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_background_14` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.3 KiB | [pg_background_14-1.9.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.9.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_14` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.0 KiB | [pg_background_14-1.9.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.9.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_14` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.4 KiB | [pg_background_14-1.8-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.8-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_14` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [pg_background_14-1.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.6-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_14` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.6 KiB | [pg_background_14-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.5 KiB | [pg_background_14-1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.0-1.rhel9.aarch64.rpm) |
| `pg_background_14` | `1.9.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.5 KiB | [pg_background_14-1.9.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_background_14-1.9.3-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_14` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.4 KiB | [pg_background_14-1.9.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_background_14-1.9.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_background_14` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.4 KiB | [pg_background_14-1.9.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_background_14-1.9.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_14` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.2 KiB | [pg_background_14-1.9.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_background_14-1.9.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_14` | `1.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.7 KiB | [pg_background_14-1.8-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_background_14-1.8-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_14` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.6 KiB | [pg_background_14-1.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_background_14-1.6-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_14` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.8 KiB | [pg_background_14-1.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_background_14-1.3-3PGDG.rhel10.x86_64.rpm) |
| `pg_background_14` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_14-1.9.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_background_14-1.9.3-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_14` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.7 KiB | [pg_background_14-1.9.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_background_14-1.9.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_background_14` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.8 KiB | [pg_background_14-1.9.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_background_14-1.9.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_14` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.5 KiB | [pg_background_14-1.9.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_background_14-1.9.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_14` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_background_14-1.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_background_14-1.8-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_14` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_background_14-1.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_background_14-1.6-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_14` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.4 KiB | [pg_background_14-1.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_background_14-1.3-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 83.2 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 88.9 KiB | [postgresql-14-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 83.1 KiB | [postgresql-14-pg-background_1.9.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.1-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 81.7 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 87.2 KiB | [postgresql-14-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 81.5 KiB | [postgresql-14-pg-background_1.9.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.1-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 83.4 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 89.0 KiB | [postgresql-14-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 83.1 KiB | [postgresql-14-pg-background_1.9.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.1-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 82.0 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 87.6 KiB | [postgresql-14-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 82.0 KiB | [postgresql-14-pg-background_1.9.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.1-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 89.6 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 100.0 KiB | [postgresql-14-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 89.4 KiB | [postgresql-14-pg-background_1.9.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 87.8 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 98.4 KiB | [postgresql-14-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 87.7 KiB | [postgresql-14-pg-background_1.9.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 83.3 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 91.8 KiB | [postgresql-14-pg-background_1.9.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 83.2 KiB | [postgresql-14-pg-background_1.9.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 82.0 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 90.9 KiB | [postgresql-14-pg-background_1.9.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 81.8 KiB | [postgresql-14-pg-background_1.9.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 83.2 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 91.2 KiB | [postgresql-14-pg-background_1.9.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 82.9 KiB | [postgresql-14-pg-background_1.9.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 81.1 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 89.8 KiB | [postgresql-14-pg-background_1.9.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1PIGSTY~resolute_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 80.9 KiB | [postgresql-14-pg-background_1.9.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.1-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/vibhorkum/pg_background" title="Repository" icon="github" subtitle="github.com/vibhorkum/pg_background" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_background-1.9.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_background;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_background;		# install via package name, for the active PG version

pig install pg_background -v 18;   # install for PG 18
pig install pg_background -v 17;   # install for PG 17
pig install pg_background -v 16;   # install for PG 16
pig install pg_background -v 15;   # install for PG 15
pig install pg_background -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_background;
```

## Usage

Sources: [official README](https://github.com/vibhorkum/pg_background/blob/master/README.md), [v1.9.2 release](https://github.com/vibhorkum/pg_background/releases/tag/v1.9.2)

`pg_background` executes SQL in PostgreSQL background worker processes. The worker runs inside the server and uses its own transaction, which makes it useful for asynchronous maintenance, autonomous side effects, and long-running tasks that should not block the caller.

```sql
CREATE EXTENSION pg_background;

SELECT * FROM pg_background_launch_v2(
  'SELECT count(*) FROM large_table',
  65536,
  'count-large-table'
) AS h;

SELECT * FROM pg_background_result_v2(h.pid, h.cookie) AS (count bigint);
```

### Core API

- `pg_background_launch_v2(sql, queue_size, label)` launches a tracked worker and returns `(pid, cookie)`.
- `pg_background_submit_v2(sql, queue_size, label)` is fire-and-forget for side-effect SQL.
- `pg_background_result_v2(pid, cookie)` consumes the worker result set once.
- `pg_background_wait_v2(...)` and `pg_background_wait_v2_timeout(...)` wait for completion.
- `pg_background_cancel_v2(...)` stops execution; `pg_background_detach_v2(...)` stops tracking but lets work continue.
- `pg_background_list_v2()`, `pg_background_stats_v2()`, and `pg_background_get_progress_v2(...)` expose worker state and progress.

### Typical Patterns

Run maintenance without holding the client session open:

```sql
SELECT * FROM pg_background_submit_v2(
  'VACUUM (ANALYZE) public.events',
  65536,
  'vacuum-events'
);
```

Use an autonomous side effect from application SQL:

```sql
SELECT * FROM pg_background_submit_v2(
  $$INSERT INTO audit_log(ts, event) VALUES (clock_timestamp(), 'job queued')$$
);
```

### GUCs And Security

- `pg_background.max_workers` limits concurrent workers per session.
- `pg_background.default_queue_size` controls the shared-memory queue size.
- `pg_background.worker_timeout` sets an execution timeout; `0` means no limit.
- The extension creates a dedicated `pg_background_worker` NOLOGIN role and ships helper functions to grant or revoke execution privileges.

### Caveats

- Prefer the V2 API. The older V1 API is still present for compatibility but lacks cookie-based PID reuse protection.
- The `v1.9.2` release is a binary-only patch for assert-enabled PostgreSQL builds. The SQL extension version remains `1.9`, so there is no new SQL upgrade script or user-facing function delta from `1.9.1`.
