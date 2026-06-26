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
| **1110** | {{< badge content="pg_background" link="https://github.com/vibhorkum/pg_background" >}} | {{< ext "pg_background" >}} | `2.0.2` | {{< category "TIME" >}} | {{< license "GPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_cron" >}} {{< ext "pg_task" >}} {{< ext "pg_later" >}} {{< ext "pgq" >}} {{< ext "timescaledb" >}} {{< ext "timescaledb_toolkit" >}} {{< ext "timeseries" >}} {{< ext "periods" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_background` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.2` | {{< bg "18" "pg_background_18" "green" >}} {{< bg "17" "pg_background_17" "green" >}} {{< bg "16" "pg_background_16" "green" >}} {{< bg "15" "pg_background_15" "green" >}} {{< bg "14" "pg_background_14" "green" >}} | `pg_background_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.2` | {{< bg "18" "postgresql-18-pg-background" "green" >}} {{< bg "17" "postgresql-17-pg-background" "green" >}} {{< bg "16" "postgresql-16-pg-background" "green" >}} {{< bg "15" "postgresql-15-pg-background" "green" >}} {{< bg "14" "postgresql-14-pg-background" "green" >}} | `postgresql-$v-pg-background` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.0.2" "pg_background_18 : AVAIL 8" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_17 : AVAIL 9" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_14 : AVAIL 9" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.0.2" "pg_background_18 : AVAIL 9" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_17 : AVAIL 10" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_16 : AVAIL 10" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_15 : AVAIL 11" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_14 : AVAIL 10" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.0.2" "pg_background_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_14 : AVAIL 4" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.0.2" "pg_background_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_14 : AVAIL 20" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.0.2" "pg_background_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_14 : AVAIL 4" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.0.2" "pg_background_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "pg_background_14 : AVAIL 4" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-pg-background : AVAIL 4" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-pg-background : AVAIL 4" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-pg-background : AVAIL 4" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-pg-background : AVAIL 4" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-pg-background : AVAIL 4" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-pg-background : AVAIL 4" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-pg-background : AVAIL 4" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-pg-background : AVAIL 4" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-pg-background : AVAIL 4" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-pg-background : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-pg-background : AVAIL 4" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_18` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 65.0 KiB | [pg_background_18-2.0.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-2.0.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 63.7 KiB | [pg_background_18-2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_background_18-2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_background_18` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.1 KiB | [pg_background_18-2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `1.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.5 KiB | [pg_background_18-1.9.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.9.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.4 KiB | [pg_background_18-1.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.1 KiB | [pg_background_18-1.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pg_background_18-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.5 KiB | [pg_background_18-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_18` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.8 KiB | [pg_background_18-2.0.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-2.0.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.5 KiB | [pg_background_18-2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_background_18-2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_background_18` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 61.8 KiB | [pg_background_18-2.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-2.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 57.3 KiB | [pg_background_18-1.9.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.9.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.1 KiB | [pg_background_18-1.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pg_background_18-1.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.8 KiB | [pg_background_18-1.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_background_18-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.0 KiB | [pg_background_18-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_18` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 63.1 KiB | [pg_background_18-2.0.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_background_18-2.0.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_background_18` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 62.0 KiB | [pg_background_18-2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_background_18-2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_background_18` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 61.4 KiB | [pg_background_18-2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_background_18-2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_background_18` | `1.9.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.4 KiB | [pg_background_18-1.9.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_background_18-1.9.3-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_background_18` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.6 KiB | [pg_background_18-2.0.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_background_18-2.0.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_background_18` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 61.2 KiB | [pg_background_18-2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_background_18-2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_background_18` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 60.8 KiB | [pg_background_18-2.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_background_18-2.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_background_18` | `1.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 56.4 KiB | [pg_background_18-1.9.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_background_18-1.9.3-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_background_18` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 63.7 KiB | [pg_background_18-2.0.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_background_18-2.0.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_background_18` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.2 KiB | [pg_background_18-2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_background_18-2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_background_18` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 61.9 KiB | [pg_background_18-2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_background_18-2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_background_18` | `1.9.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.6 KiB | [pg_background_18-1.9.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_background_18-1.9.3-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_background_18` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.2 KiB | [pg_background_18-2.0.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-2.0.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_18` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 61.6 KiB | [pg_background_18-2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_background_18-2.0-1PIGSTY.el10.aarch64.rpm) |
| `pg_background_18` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.3 KiB | [pg_background_18-2.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-2.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_18` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_18-1.9.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.9.3-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-18-pg-background` | `2.0.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 67.6 KiB | [postgresql-18-pg-background_2.0.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.2-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 67.9 KiB | [postgresql-18-pg-background_2.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0-2.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 100.9 KiB | [postgresql-18-pg-background_2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-18-pg-background_2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 58.9 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 66.3 KiB | [postgresql-18-pg-background_2.0.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.2-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-background` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 66.3 KiB | [postgresql-18-pg-background_2.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0-2.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-background` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 99.4 KiB | [postgresql-18-pg-background_2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-18-pg-background_2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 57.0 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 68.0 KiB | [postgresql-18-pg-background_2.0.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.2-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 68.0 KiB | [postgresql-18-pg-background_2.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0-2.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 101.3 KiB | [postgresql-18-pg-background_2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-18-pg-background_2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 58.7 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 66.5 KiB | [postgresql-18-pg-background_2.0.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.2-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-background` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 66.5 KiB | [postgresql-18-pg-background_2.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0-2.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-background` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 99.7 KiB | [postgresql-18-pg-background_2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-18-pg-background_2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 57.3 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 69.1 KiB | [postgresql-18-pg-background_2.0.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 69.2 KiB | [postgresql-18-pg-background_2.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 107.7 KiB | [postgresql-18-pg-background_2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-18-pg-background_2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 59.7 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 67.5 KiB | [postgresql-18-pg-background_2.0.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-background` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 67.4 KiB | [postgresql-18-pg-background_2.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-background` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 105.9 KiB | [postgresql-18-pg-background_2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-18-pg-background_2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 58.0 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 68.2 KiB | [postgresql-18-pg-background_2.0.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 68.1 KiB | [postgresql-18-pg-background_2.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 104.3 KiB | [postgresql-18-pg-background_2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-18-pg-background_2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 58.8 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 66.2 KiB | [postgresql-18-pg-background_2.0.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pg-background` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 66.2 KiB | [postgresql-18-pg-background_2.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0-2.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pg-background` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 103.1 KiB | [postgresql-18-pg-background_2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-18-pg-background_2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 57.2 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 67.4 KiB | [postgresql-18-pg-background_2.0.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 67.5 KiB | [postgresql-18-pg-background_2.0-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0-2.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 103.4 KiB | [postgresql-18-pg-background_2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-18-pg-background_2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 58.4 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 65.8 KiB | [postgresql-18-pg-background_2.0.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.2-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-pg-background` | `2.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 65.7 KiB | [postgresql-18-pg-background_2.0-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0-2.pgdg26.04+1_arm64.deb) |
| `postgresql-18-pg-background` | `2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 102.4 KiB | [postgresql-18-pg-background_2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-18-pg-background_2.0-1PIGSTY~resolute_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 56.3 KiB | [postgresql-18-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_17` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 65.0 KiB | [pg_background_17-2.0.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-2.0.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 63.7 KiB | [pg_background_17-2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_background_17-2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_background_17` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.1 KiB | [pg_background_17-2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `1.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.5 KiB | [pg_background_17-1.9.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.9.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.4 KiB | [pg_background_17-1.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.1 KiB | [pg_background_17-1.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [pg_background_17-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.9 KiB | [pg_background_17-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_17` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.1 KiB | [pg_background_17-1.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.2-2PGDG.rhel8.x86_64.rpm) |
| `pg_background_17` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.6 KiB | [pg_background_17-2.0.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-2.0.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.4 KiB | [pg_background_17-2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_background_17-2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_background_17` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 61.7 KiB | [pg_background_17-2.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-2.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 57.2 KiB | [pg_background_17-1.9.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.9.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.1 KiB | [pg_background_17-1.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pg_background_17-1.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.8 KiB | [pg_background_17-1.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_background_17-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.3 KiB | [pg_background_17-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_17` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.6 KiB | [pg_background_17-1.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.2-2PGDG.rhel8.aarch64.rpm) |
| `pg_background_17` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 63.1 KiB | [pg_background_17-2.0.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_background_17-2.0.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_background_17` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 62.0 KiB | [pg_background_17-2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_background_17-2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_background_17` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 61.4 KiB | [pg_background_17-2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_background_17-2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_background_17` | `1.9.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.3 KiB | [pg_background_17-1.9.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_background_17-1.9.3-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_background_17` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.6 KiB | [pg_background_17-2.0.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_background_17-2.0.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_background_17` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 61.2 KiB | [pg_background_17-2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_background_17-2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_background_17` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 60.7 KiB | [pg_background_17-2.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_background_17-2.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_background_17` | `1.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 56.4 KiB | [pg_background_17-1.9.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_background_17-1.9.3-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_background_17` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 63.6 KiB | [pg_background_17-2.0.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_background_17-2.0.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_background_17` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.4 KiB | [pg_background_17-2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_background_17-2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_background_17` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 61.9 KiB | [pg_background_17-2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_background_17-2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_background_17` | `1.9.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.5 KiB | [pg_background_17-1.9.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_background_17-1.9.3-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_background_17` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.1 KiB | [pg_background_17-2.0.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-2.0.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_17` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 61.5 KiB | [pg_background_17-2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_background_17-2.0-1PIGSTY.el10.aarch64.rpm) |
| `pg_background_17` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.2 KiB | [pg_background_17-2.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-2.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_17` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_17-1.9.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.9.3-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-17-pg-background` | `2.0.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 67.7 KiB | [postgresql-17-pg-background_2.0.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.2-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 67.6 KiB | [postgresql-17-pg-background_2.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0-2.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 100.9 KiB | [postgresql-17-pg-background_2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-17-pg-background_2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 58.5 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 66.3 KiB | [postgresql-17-pg-background_2.0.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.2-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-background` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 66.2 KiB | [postgresql-17-pg-background_2.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0-2.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-background` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 99.2 KiB | [postgresql-17-pg-background_2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-17-pg-background_2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 57.0 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 67.9 KiB | [postgresql-17-pg-background_2.0.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.2-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 68.0 KiB | [postgresql-17-pg-background_2.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0-2.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 101.3 KiB | [postgresql-17-pg-background_2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-17-pg-background_2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 58.7 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 66.4 KiB | [postgresql-17-pg-background_2.0.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.2-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-background` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 66.5 KiB | [postgresql-17-pg-background_2.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0-2.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-background` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 99.6 KiB | [postgresql-17-pg-background_2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-17-pg-background_2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 57.1 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 76.6 KiB | [postgresql-17-pg-background_2.0.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 76.6 KiB | [postgresql-17-pg-background_2.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 115.2 KiB | [postgresql-17-pg-background_2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-17-pg-background_2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 64.8 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 74.9 KiB | [postgresql-17-pg-background_2.0.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-background` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 74.9 KiB | [postgresql-17-pg-background_2.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-background` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 113.6 KiB | [postgresql-17-pg-background_2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-17-pg-background_2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 63.1 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 68.1 KiB | [postgresql-17-pg-background_2.0.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 68.0 KiB | [postgresql-17-pg-background_2.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 104.3 KiB | [postgresql-17-pg-background_2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-17-pg-background_2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 58.7 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 66.2 KiB | [postgresql-17-pg-background_2.0.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pg-background` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 66.1 KiB | [postgresql-17-pg-background_2.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0-2.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pg-background` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 103.0 KiB | [postgresql-17-pg-background_2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-17-pg-background_2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 57.2 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 67.3 KiB | [postgresql-17-pg-background_2.0.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 67.3 KiB | [postgresql-17-pg-background_2.0-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0-2.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 103.2 KiB | [postgresql-17-pg-background_2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-17-pg-background_2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 58.3 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 65.7 KiB | [postgresql-17-pg-background_2.0.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.2-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-pg-background` | `2.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 65.7 KiB | [postgresql-17-pg-background_2.0-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0-2.pgdg26.04+1_arm64.deb) |
| `postgresql-17-pg-background` | `2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 102.3 KiB | [postgresql-17-pg-background_2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-17-pg-background_2.0-1PIGSTY~resolute_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 56.2 KiB | [postgresql-17-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_16` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 65.0 KiB | [pg_background_16-2.0.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-2.0.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 63.6 KiB | [pg_background_16-2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_background_16-2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_background_16` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.1 KiB | [pg_background_16-2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `1.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.5 KiB | [pg_background_16-1.9.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.9.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.3 KiB | [pg_background_16-1.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.1 KiB | [pg_background_16-1.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pg_background_16-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.9 KiB | [pg_background_16-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_16` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.7 KiB | [pg_background_16-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_16` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.6 KiB | [pg_background_16-2.0.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-2.0.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.4 KiB | [pg_background_16-2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_background_16-2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_background_16` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 61.6 KiB | [pg_background_16-2.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-2.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 57.2 KiB | [pg_background_16-1.9.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.9.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.1 KiB | [pg_background_16-1.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.8 KiB | [pg_background_16-1.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.7 KiB | [pg_background_16-1.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_background_16-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.3 KiB | [pg_background_16-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_16` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.2 KiB | [pg_background_16-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_16` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 63.1 KiB | [pg_background_16-2.0.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_background_16-2.0.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_background_16` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 61.8 KiB | [pg_background_16-2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_background_16-2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_background_16` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 61.4 KiB | [pg_background_16-2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_background_16-2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_background_16` | `1.9.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.3 KiB | [pg_background_16-1.9.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_background_16-1.9.3-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_background_16` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.6 KiB | [pg_background_16-2.0.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_background_16-2.0.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_background_16` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 61.2 KiB | [pg_background_16-2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_background_16-2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_background_16` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 60.8 KiB | [pg_background_16-2.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_background_16-2.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_background_16` | `1.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 56.4 KiB | [pg_background_16-1.9.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_background_16-1.9.3-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_background_16` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 63.6 KiB | [pg_background_16-2.0.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_background_16-2.0.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_background_16` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.4 KiB | [pg_background_16-2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_background_16-2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_background_16` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 61.9 KiB | [pg_background_16-2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_background_16-2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_background_16` | `1.9.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.5 KiB | [pg_background_16-1.9.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_background_16-1.9.3-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_background_16` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.1 KiB | [pg_background_16-2.0.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-2.0.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_16` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 61.5 KiB | [pg_background_16-2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_background_16-2.0-1PIGSTY.el10.aarch64.rpm) |
| `pg_background_16` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.3 KiB | [pg_background_16-2.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-2.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_16` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_16-1.9.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.9.3-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-16-pg-background` | `2.0.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 67.6 KiB | [postgresql-16-pg-background_2.0.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.2-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 67.6 KiB | [postgresql-16-pg-background_2.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0-2.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 100.9 KiB | [postgresql-16-pg-background_2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-16-pg-background_2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 58.8 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 66.2 KiB | [postgresql-16-pg-background_2.0.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.2-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-background` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 66.1 KiB | [postgresql-16-pg-background_2.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0-2.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-background` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 99.2 KiB | [postgresql-16-pg-background_2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-16-pg-background_2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 56.9 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 68.0 KiB | [postgresql-16-pg-background_2.0.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.2-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 68.1 KiB | [postgresql-16-pg-background_2.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0-2.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 101.4 KiB | [postgresql-16-pg-background_2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-16-pg-background_2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 58.5 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 66.5 KiB | [postgresql-16-pg-background_2.0.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.2-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-background` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 66.5 KiB | [postgresql-16-pg-background_2.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0-2.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-background` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 99.7 KiB | [postgresql-16-pg-background_2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-16-pg-background_2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 57.2 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 76.6 KiB | [postgresql-16-pg-background_2.0.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 76.5 KiB | [postgresql-16-pg-background_2.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 115.1 KiB | [postgresql-16-pg-background_2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-16-pg-background_2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 64.8 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 74.9 KiB | [postgresql-16-pg-background_2.0.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-background` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 74.9 KiB | [postgresql-16-pg-background_2.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-background` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 113.5 KiB | [postgresql-16-pg-background_2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-16-pg-background_2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 63.0 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 68.2 KiB | [postgresql-16-pg-background_2.0.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 68.2 KiB | [postgresql-16-pg-background_2.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 104.3 KiB | [postgresql-16-pg-background_2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-16-pg-background_2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 58.5 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 66.2 KiB | [postgresql-16-pg-background_2.0.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pg-background` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 66.3 KiB | [postgresql-16-pg-background_2.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0-2.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pg-background` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 103.0 KiB | [postgresql-16-pg-background_2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-16-pg-background_2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 57.2 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 67.1 KiB | [postgresql-16-pg-background_2.0.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 67.1 KiB | [postgresql-16-pg-background_2.0-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0-2.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 103.2 KiB | [postgresql-16-pg-background_2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-16-pg-background_2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 58.2 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 65.7 KiB | [postgresql-16-pg-background_2.0.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.2-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-pg-background` | `2.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 65.7 KiB | [postgresql-16-pg-background_2.0-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0-2.pgdg26.04+1_arm64.deb) |
| `postgresql-16-pg-background` | `2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 102.3 KiB | [postgresql-16-pg-background_2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-16-pg-background_2.0-1PIGSTY~resolute_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 56.2 KiB | [postgresql-16-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_15` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 65.1 KiB | [pg_background_15-2.0.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-2.0.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 63.7 KiB | [pg_background_15-2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_background_15-2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_background_15` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.2 KiB | [pg_background_15-2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `1.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.5 KiB | [pg_background_15-1.9.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.9.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.3 KiB | [pg_background_15-1.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.1 KiB | [pg_background_15-1.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pg_background_15-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.9 KiB | [pg_background_15-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.6 KiB | [pg_background_15-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.3 KiB | [pg_background_15-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.0-1.rhel8.x86_64.rpm) |
| `pg_background_15` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.6 KiB | [pg_background_15-2.0.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-2.0.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.4 KiB | [pg_background_15-2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_background_15-2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_background_15` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 61.7 KiB | [pg_background_15-2.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-2.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 57.2 KiB | [pg_background_15-1.9.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.9.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.1 KiB | [pg_background_15-1.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.8 KiB | [pg_background_15-1.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.7 KiB | [pg_background_15-1.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_background_15-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.2 KiB | [pg_background_15-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.2 KiB | [pg_background_15-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.7 KiB | [pg_background_15-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.0-1.rhel8.aarch64.rpm) |
| `pg_background_15` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 63.4 KiB | [pg_background_15-2.0.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-2.0.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_background_15` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 62.1 KiB | [pg_background_15-2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_background_15-2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_background_15` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 61.7 KiB | [pg_background_15-2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_background_15` | `1.9.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.3 KiB | [pg_background_15-1.9.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-1.9.3-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_background_15` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.7 KiB | [pg_background_15-2.0.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-2.0.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_background_15` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 61.3 KiB | [pg_background_15-2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_background_15-2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_background_15` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 60.9 KiB | [pg_background_15-2.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-2.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_background_15` | `1.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 56.3 KiB | [pg_background_15-1.9.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-1.9.3-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_background_15` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 63.8 KiB | [pg_background_15-2.0.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_background_15-2.0.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_background_15` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.6 KiB | [pg_background_15-2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_background_15-2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_background_15` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 62.0 KiB | [pg_background_15-2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_background_15-2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_background_15` | `1.9.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.5 KiB | [pg_background_15-1.9.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_background_15-1.9.3-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_background_15` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.2 KiB | [pg_background_15-2.0.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-2.0.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_15` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 61.7 KiB | [pg_background_15-2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_background_15-2.0-1PIGSTY.el10.aarch64.rpm) |
| `pg_background_15` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.4 KiB | [pg_background_15-2.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-2.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_15` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_15-1.9.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.9.3-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-15-pg-background` | `2.0.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 67.8 KiB | [postgresql-15-pg-background_2.0.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.2-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 67.8 KiB | [postgresql-15-pg-background_2.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0-2.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 101.1 KiB | [postgresql-15-pg-background_2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-15-pg-background_2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 58.5 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 66.2 KiB | [postgresql-15-pg-background_2.0.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.2-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-background` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 66.2 KiB | [postgresql-15-pg-background_2.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0-2.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-background` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 99.4 KiB | [postgresql-15-pg-background_2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-15-pg-background_2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 56.9 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 67.9 KiB | [postgresql-15-pg-background_2.0.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.2-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 68.0 KiB | [postgresql-15-pg-background_2.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0-2.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 101.1 KiB | [postgresql-15-pg-background_2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-15-pg-background_2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 58.5 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 66.5 KiB | [postgresql-15-pg-background_2.0.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.2-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-background` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 66.4 KiB | [postgresql-15-pg-background_2.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0-2.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-background` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 99.7 KiB | [postgresql-15-pg-background_2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-15-pg-background_2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 57.3 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 76.8 KiB | [postgresql-15-pg-background_2.0.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 76.8 KiB | [postgresql-15-pg-background_2.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 115.4 KiB | [postgresql-15-pg-background_2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-15-pg-background_2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 64.7 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 75.0 KiB | [postgresql-15-pg-background_2.0.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-background` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 75.1 KiB | [postgresql-15-pg-background_2.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-background` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 113.6 KiB | [postgresql-15-pg-background_2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-15-pg-background_2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 63.0 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 68.2 KiB | [postgresql-15-pg-background_2.0.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 68.1 KiB | [postgresql-15-pg-background_2.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 104.6 KiB | [postgresql-15-pg-background_2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-15-pg-background_2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 58.8 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 66.3 KiB | [postgresql-15-pg-background_2.0.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pg-background` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 66.1 KiB | [postgresql-15-pg-background_2.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0-2.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pg-background` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 103.1 KiB | [postgresql-15-pg-background_2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-15-pg-background_2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 57.3 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 67.3 KiB | [postgresql-15-pg-background_2.0.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 67.3 KiB | [postgresql-15-pg-background_2.0-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0-2.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 103.4 KiB | [postgresql-15-pg-background_2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-15-pg-background_2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 58.3 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 65.8 KiB | [postgresql-15-pg-background_2.0.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.2-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-pg-background` | `2.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 65.7 KiB | [postgresql-15-pg-background_2.0-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0-2.pgdg26.04+1_arm64.deb) |
| `postgresql-15-pg-background` | `2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 102.4 KiB | [postgresql-15-pg-background_2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-15-pg-background_2.0-1PIGSTY~resolute_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 56.4 KiB | [postgresql-15-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_14` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 65.1 KiB | [pg_background_14-2.0.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-2.0.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_14` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 63.7 KiB | [pg_background_14-2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_background_14-2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_background_14` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.1 KiB | [pg_background_14-2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_14` | `1.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.4 KiB | [pg_background_14-1.9.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-1.9.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_14` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.3 KiB | [pg_background_14-1.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-1.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_14` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.1 KiB | [pg_background_14-1.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-1.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_14` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pg_background_14-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_14` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.9 KiB | [pg_background_14-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.6 KiB | [pg_background_14-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-1.0-1.rhel8.x86_64.rpm) |
| `pg_background_14` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.6 KiB | [pg_background_14-2.0.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-2.0.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_14` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.4 KiB | [pg_background_14-2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_background_14-2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_background_14` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 61.6 KiB | [pg_background_14-2.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-2.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_14` | `1.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 57.1 KiB | [pg_background_14-1.9.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.9.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_14` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.0 KiB | [pg_background_14-1.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_14` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.8 KiB | [pg_background_14-1.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_14` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.6 KiB | [pg_background_14-1.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_14` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.6 KiB | [pg_background_14-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_14` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.2 KiB | [pg_background_14-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.6 KiB | [pg_background_14-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.0-1.rhel8.aarch64.rpm) |
| `pg_background_14` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 63.3 KiB | [pg_background_14-2.0.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_background_14-2.0.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_background_14` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 62.0 KiB | [pg_background_14-2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_background_14-2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_background_14` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 61.6 KiB | [pg_background_14-2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_background_14-2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_background_14` | `1.9.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.3 KiB | [pg_background_14-1.9.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_background_14-1.9.3-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_background_14` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.6 KiB | [pg_background_14-2.0.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-2.0.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_background_14` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.6 KiB | [pg_background_14-2.0.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-2.0.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_14` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.7 KiB | [pg_background_14-2.0.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-2.0.2-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_background_14` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 61.2 KiB | [pg_background_14-2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_background_14-2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_background_14` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 60.8 KiB | [pg_background_14-2.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-2.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_background_14` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 60.8 KiB | [pg_background_14-2.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-2.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_14` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 61.0 KiB | [pg_background_14-2.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-2.0-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_background_14` | `1.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 56.3 KiB | [pg_background_14-1.9.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.9.3-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_background_14` | `1.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 56.3 KiB | [pg_background_14-1.9.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.9.3-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_14` | `1.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 56.4 KiB | [pg_background_14-1.9.3-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.9.3-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_background_14` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.3 KiB | [pg_background_14-1.9.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.9.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_14` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.4 KiB | [pg_background_14-1.9.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.9.2-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_background_14` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.0 KiB | [pg_background_14-1.9.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.9.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_14` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.2 KiB | [pg_background_14-1.9.1-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.9.1-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_background_14` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.4 KiB | [pg_background_14-1.8-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.8-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_14` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.5 KiB | [pg_background_14-1.8-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.8-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_background_14` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [pg_background_14-1.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.6-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_14` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.5 KiB | [pg_background_14-1.6-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.6-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_background_14` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.6 KiB | [pg_background_14-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.5 KiB | [pg_background_14-1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.0-1.rhel9.aarch64.rpm) |
| `pg_background_14` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 63.7 KiB | [pg_background_14-2.0.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_background_14-2.0.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_background_14` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.6 KiB | [pg_background_14-2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_background_14-2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_background_14` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 62.1 KiB | [pg_background_14-2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_background_14-2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_background_14` | `1.9.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.5 KiB | [pg_background_14-1.9.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_background_14-1.9.3-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_background_14` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.2 KiB | [pg_background_14-2.0.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_background_14-2.0.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_14` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 61.6 KiB | [pg_background_14-2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_background_14-2.0-1PIGSTY.el10.aarch64.rpm) |
| `pg_background_14` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.3 KiB | [pg_background_14-2.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_background_14-2.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_14` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_14-1.9.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_background_14-1.9.3-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-14-pg-background` | `2.0.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 92.7 KiB | [postgresql-14-pg-background_2.0.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0.2-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-background` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 92.7 KiB | [postgresql-14-pg-background_2.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0-2.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-background` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 100.9 KiB | [postgresql-14-pg-background_2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-14-pg-background_2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 83.2 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-background` | `2.0.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 91.0 KiB | [postgresql-14-pg-background_2.0.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0.2-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-background` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 90.9 KiB | [postgresql-14-pg-background_2.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0-2.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-background` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 99.4 KiB | [postgresql-14-pg-background_2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-14-pg-background_2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 81.7 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-background` | `2.0.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 92.7 KiB | [postgresql-14-pg-background_2.0.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0.2-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-background` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 92.8 KiB | [postgresql-14-pg-background_2.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0-2.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-background` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 101.0 KiB | [postgresql-14-pg-background_2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-14-pg-background_2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 83.4 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-background` | `2.0.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 91.3 KiB | [postgresql-14-pg-background_2.0.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0.2-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-background` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 91.3 KiB | [postgresql-14-pg-background_2.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0-2.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-background` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 99.6 KiB | [postgresql-14-pg-background_2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-14-pg-background_2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 82.0 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-background` | `2.0.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 101.6 KiB | [postgresql-14-pg-background_2.0.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-background` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 101.7 KiB | [postgresql-14-pg-background_2.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-background` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 115.3 KiB | [postgresql-14-pg-background_2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-14-pg-background_2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 89.6 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-background` | `2.0.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 99.7 KiB | [postgresql-14-pg-background_2.0.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-background` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 99.6 KiB | [postgresql-14-pg-background_2.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-background` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 113.6 KiB | [postgresql-14-pg-background_2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-14-pg-background_2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 87.8 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-background` | `2.0.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 93.0 KiB | [postgresql-14-pg-background_2.0.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-background` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 93.0 KiB | [postgresql-14-pg-background_2.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-background` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 104.4 KiB | [postgresql-14-pg-background_2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-14-pg-background_2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 83.3 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-background` | `2.0.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 91.0 KiB | [postgresql-14-pg-background_2.0.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pg-background` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 91.0 KiB | [postgresql-14-pg-background_2.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0-2.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pg-background` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 103.1 KiB | [postgresql-14-pg-background_2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-14-pg-background_2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 82.0 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pg-background` | `2.0.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 92.1 KiB | [postgresql-14-pg-background_2.0.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pg-background` | `2.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 92.1 KiB | [postgresql-14-pg-background_2.0-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0-2.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pg-background` | `2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 103.1 KiB | [postgresql-14-pg-background_2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-14-pg-background_2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 83.2 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pg-background` | `2.0.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 90.5 KiB | [postgresql-14-pg-background_2.0.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0.2-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-pg-background` | `2.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 90.6 KiB | [postgresql-14-pg-background_2.0-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_2.0-2.pgdg26.04+1_arm64.deb) |
| `postgresql-14-pg-background` | `2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 102.3 KiB | [postgresql-14-pg-background_2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-14-pg-background_2.0-1PIGSTY~resolute_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 81.1 KiB | [postgresql-14-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/vibhorkum/pg_background" title="Repository" icon="github" subtitle="github.com/vibhorkum/pg_background" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_background-2.0.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_background;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
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

Sources: [official README](https://github.com/vibhorkum/pg_background/blob/master/README.md), [v2.0 release notes](https://github.com/vibhorkum/pg_background/releases/tag/v2.0), [migration guide](https://github.com/vibhorkum/pg_background/blob/v2.0/docs/MIGRATION.md).

`pg_background` executes SQL inside PostgreSQL background worker processes. Workers run independent transactions inside the server, which is useful for asynchronous maintenance, autonomous side effects, bounded long-running tasks, and progress-tracked jobs.

Version 2.0 makes the unsuffixed API canonical. The old `_v2` names remain deprecated aliases through the 2.x line, but new code should use names such as `pg_background_launch`, `pg_background_result`, and `pg_background_run`.

### One-Shot Execution

Use `pg_background_run` when the SQL has side effects and you only need execution metadata:

```sql
CREATE EXTENSION pg_background;

SELECT completed, has_error, sqlstate, error_message,
       row_count, command_tag, elapsed_ms, timed_out
FROM pg_background_run(
  'INSERT INTO audit_log(ts, who) VALUES (clock_timestamp(), current_user)',
  queue_size := 0,
  timeout_ms := 30000,
  label := 'audit-login'
);
```

### Launch And Fetch Results

Use the launch/result pattern when the background SQL returns rows:

```sql
SELECT * FROM pg_background_launch(
  'SELECT count(*) FROM large_table',
  queue_size := 65536,
  label := 'count-large-table'
) AS h;

SELECT * FROM pg_background_result(h.pid, h.cookie) AS (count bigint);
```

Results can be consumed once. Keep both `pid` and `cookie`; the cookie protects later calls from PID reuse.

### Fire And Forget

For side effects where no result rows need to be consumed:

```sql
SELECT * FROM pg_background_submit(
  $$VACUUM (ANALYZE) public.events$$,
  queue_size := 65536,
  label := 'vacuum-events'
);
```

### Core API

- `pg_background_launch(sql, queue_size, label)` launches a worker and returns `pg_background_handle(pid, cookie)`.
- `pg_background_submit(sql, queue_size, label)` launches fire-and-forget work and returns a handle.
- `pg_background_result(pid, cookie)` consumes result rows once.
- `pg_background_result_info(pid, cookie)` returns completion and row-count metadata without consuming rows.
- `pg_background_error_info(pid, cookie)` returns structured SQL error details.
- `pg_background_wait(pid, cookie, timeout_ms DEFAULT 0)` waits for completion; `timeout_ms <= 0` blocks indefinitely.
- `pg_background_cancel(pid, cookie, grace_ms DEFAULT 0)` requests cooperative cancellation.
- `pg_background_detach(pid, cookie)` stops tracking a worker while letting it continue.
- `pg_background_outcome(pid, cookie)` returns a combined status snapshot without raising on missing state.
- `pg_background_list` and `pg_background_activity` are monitoring views; `pg_background_stats()` returns session counters.

Convenience helpers include `pg_background_run_query`, `pg_background_drain`, `pg_background_wait_any`, `pg_background_cancel_by_label`, and `pg_background_purge`.

### Progress Reporting

Report progress from inside the worker SQL, then poll it from the launcher:

```sql
SELECT * FROM pg_background_launch($$
  SELECT pg_background_report_progress(0, 'starting');
  SELECT pg_sleep(1);
  SELECT pg_background_report_progress(100, 'done');
$$) AS h;

SELECT * FROM pg_background_get_progress(h.pid, h.cookie);
```

`pg_background_report_progress` is the 2.0 name; the earlier `pg_background_progress` name was hard-renamed.

### GUCs And Loading

`pg_background` does not require `shared_preload_libraries`. Preloading is optional and mainly useful when you want its GUCs available before the extension is first loaded in a session.

```sql
SET pg_background.max_workers = 10;
SET pg_background.default_queue_size = '256KB';
SET pg_background.worker_timeout = '5min';
```

- `pg_background.max_workers` defaults to `16`.
- `pg_background.default_queue_size` defaults to `65536` bytes.
- `pg_background.worker_timeout` defaults to `0`, meaning no execution timeout.

### Caveats

- Pigsty packages `pg_background` 2.0 for PostgreSQL 14-18; upstream 2.0 also validates PostgreSQL 19 beta.
- Upgrades from pre-1.8 installs must first reach the 1.8/1.10 release line before updating to 2.0.
- The original v1 PID-only API was removed. Unsuffixed names now have cookie-protected semantics and return/use `(pid, cookie)`.
- `pg_background_cancel_v2_grace` and `pg_background_wait_v2_timeout` are folded into `pg_background_cancel(..., grace_ms)` and `pg_background_wait(..., timeout_ms)`.
- `pg_background_status_v2` was removed; use `pg_background_outcome(pid, cookie)`.
