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


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.9.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_background` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.9.2` | {{< bg "18" "pg_background_18" "green" >}} {{< bg "17" "pg_background_17" "green" >}} {{< bg "16" "pg_background_16" "green" >}} {{< bg "15" "pg_background_15" "green" >}} {{< bg "14" "pg_background_14" "green" >}} | `pg_background_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.9.2` | {{< bg "18" "postgresql-18-pg-background" "green" >}} {{< bg "17" "postgresql-17-pg-background" "green" >}} {{< bg "16" "postgresql-16-pg-background" "green" >}} {{< bg "15" "postgresql-15-pg-background" "green" >}} {{< bg "14" "postgresql-14-pg-background" "green" >}} | `postgresql-$v-pg-background` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_16 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_15 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_14 : AVAIL 5" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_15 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_14 : AVAIL 6" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_15 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_14 : AVAIL 6" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_15 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_14 : AVAIL 6" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_16 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_15 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_14 : AVAIL 5" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_16 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_15 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.9.2" "pg_background_14 : AVAIL 5" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-18-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-17-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-16-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-15-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-14-pg-background : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-18-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-17-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-16-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-15-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-14-pg-background : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-18-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-17-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-16-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-15-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-14-pg-background : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-18-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-17-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-16-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-15-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-14-pg-background : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-18-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-17-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-16-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-15-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-14-pg-background : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-18-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-17-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-16-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-15-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-14-pg-background : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-18-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-17-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-16-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-15-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-14-pg-background : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-18-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-17-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-16-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-15-pg-background : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.2" "postgresql-14-pg-background : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_18` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 56.3 KiB | [pg_background_18-1.9.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_background_18-1.9.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_background_18` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.1 KiB | [pg_background_18-1.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pg_background_18-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.5 KiB | [pg_background_18-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_18` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 55.2 KiB | [pg_background_18-1.9.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_background_18-1.9.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_background_18` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pg_background_18-1.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.8 KiB | [pg_background_18-1.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_background_18-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.0 KiB | [pg_background_18-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_18` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.3 KiB | [pg_background_18-1.9.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_background_18-1.9.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_background_18` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 54.0 KiB | [pg_background_18-1.9.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_background_18-1.9.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_18` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.5 KiB | [pg_background_18-1.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_background_18-1.8-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_18` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.3 KiB | [pg_background_18-1.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_background_18-1.6-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_18` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.7 KiB | [pg_background_18-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_background_18-1.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_18` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.3 KiB | [pg_background_18-1.9.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_background_18-1.9.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_background_18` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.1 KiB | [pg_background_18-1.9.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_background_18-1.9.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_18` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.4 KiB | [pg_background_18-1.8-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_background_18-1.8-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_18` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [pg_background_18-1.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_background_18-1.6-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_18` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.0 KiB | [pg_background_18-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_background_18-1.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_18` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.5 KiB | [pg_background_18-1.9.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_background_18-1.9.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_background_18` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.2 KiB | [pg_background_18-1.9.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_background_18-1.9.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_18` | `1.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.8 KiB | [pg_background_18-1.8-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_background_18-1.8-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_18` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.6 KiB | [pg_background_18-1.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_background_18-1.6-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_18` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.1 KiB | [pg_background_18-1.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_background_18-1.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_background_18` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.7 KiB | [pg_background_18-1.9.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_background_18-1.9.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_background_18` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.5 KiB | [pg_background_18-1.9.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.9.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_18` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_background_18-1.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.8-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_18` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_background_18-1.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.6-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_18` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.6 KiB | [pg_background_18-1.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 89.2 KiB | [postgresql-18-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 87.4 KiB | [postgresql-18-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 89.3 KiB | [postgresql-18-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 87.9 KiB | [postgresql-18-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 94.9 KiB | [postgresql-18-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 93.6 KiB | [postgresql-18-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 92.0 KiB | [postgresql-18-pg-background_1.9.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 91.1 KiB | [postgresql-18-pg-background_1.9.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-18-pg-background_1.9.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_17` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 56.3 KiB | [pg_background_17-1.9.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_background_17-1.9.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_background_17` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.1 KiB | [pg_background_17-1.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [pg_background_17-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.9 KiB | [pg_background_17-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_17` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.1 KiB | [pg_background_17-1.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.2-2PGDG.rhel8.x86_64.rpm) |
| `pg_background_17` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 55.1 KiB | [pg_background_17-1.9.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_background_17-1.9.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_background_17` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pg_background_17-1.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.8 KiB | [pg_background_17-1.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_background_17-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.3 KiB | [pg_background_17-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_17` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.6 KiB | [pg_background_17-1.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.2-2PGDG.rhel8.aarch64.rpm) |
| `pg_background_17` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.5 KiB | [pg_background_17-1.9.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_background_17-1.9.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_background_17` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.9 KiB | [pg_background_17-1.9.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_background_17-1.9.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_17` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.4 KiB | [pg_background_17-1.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_background_17-1.8-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_17` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.3 KiB | [pg_background_17-1.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_background_17-1.6-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_17` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.3 KiB | [pg_background_17-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_background_17-1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_17` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.5 KiB | [pg_background_17-1.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_background_17-1.2-2PGDG.rhel9.x86_64.rpm) |
| `pg_background_17` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.3 KiB | [pg_background_17-1.9.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_background_17-1.9.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_background_17` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.1 KiB | [pg_background_17-1.9.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_background_17-1.9.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_17` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.4 KiB | [pg_background_17-1.8-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_background_17-1.8-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_17` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [pg_background_17-1.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_background_17-1.6-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_17` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.6 KiB | [pg_background_17-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_background_17-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_17` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.9 KiB | [pg_background_17-1.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_background_17-1.2-2PGDG.rhel9.aarch64.rpm) |
| `pg_background_17` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.4 KiB | [pg_background_17-1.9.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_background_17-1.9.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_background_17` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.1 KiB | [pg_background_17-1.9.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_background_17-1.9.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_17` | `1.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.8 KiB | [pg_background_17-1.8-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_background_17-1.8-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_17` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.6 KiB | [pg_background_17-1.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_background_17-1.6-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_17` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.8 KiB | [pg_background_17-1.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_background_17-1.3-3PGDG.rhel10.x86_64.rpm) |
| `pg_background_17` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.7 KiB | [pg_background_17-1.9.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_background_17-1.9.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_background_17` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.6 KiB | [pg_background_17-1.9.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.9.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_17` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_background_17-1.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.8-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_17` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_background_17-1.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.6-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_17` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.3 KiB | [pg_background_17-1.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.3-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 89.0 KiB | [postgresql-17-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 87.5 KiB | [postgresql-17-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 89.1 KiB | [postgresql-17-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 87.8 KiB | [postgresql-17-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 100.2 KiB | [postgresql-17-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 98.6 KiB | [postgresql-17-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 92.0 KiB | [postgresql-17-pg-background_1.9.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 91.0 KiB | [postgresql-17-pg-background_1.9.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-17-pg-background_1.9.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_16` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 56.2 KiB | [pg_background_16-1.9.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_background_16-1.9.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_background_16` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.1 KiB | [pg_background_16-1.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pg_background_16-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.9 KiB | [pg_background_16-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_16` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.7 KiB | [pg_background_16-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_16` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 55.1 KiB | [pg_background_16-1.9.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_background_16-1.9.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_background_16` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.8 KiB | [pg_background_16-1.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.7 KiB | [pg_background_16-1.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_background_16-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.3 KiB | [pg_background_16-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_16` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.2 KiB | [pg_background_16-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_16` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.2 KiB | [pg_background_16-1.9.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_background_16-1.9.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_background_16` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.9 KiB | [pg_background_16-1.9.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_background_16-1.9.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_16` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.4 KiB | [pg_background_16-1.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_background_16-1.8-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_16` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.2 KiB | [pg_background_16-1.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_background_16-1.6-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_16` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.3 KiB | [pg_background_16-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_background_16-1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_16` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.9 KiB | [pg_background_16-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_background_16-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_16` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.4 KiB | [pg_background_16-1.9.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_background_16-1.9.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_background_16` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.1 KiB | [pg_background_16-1.9.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_background_16-1.9.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_16` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.3 KiB | [pg_background_16-1.8-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_background_16-1.8-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_16` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [pg_background_16-1.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_background_16-1.6-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_16` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.6 KiB | [pg_background_16-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_background_16-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_16` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.2 KiB | [pg_background_16-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_background_16-1.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_16` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.4 KiB | [pg_background_16-1.9.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_background_16-1.9.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_background_16` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.2 KiB | [pg_background_16-1.9.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_background_16-1.9.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_16` | `1.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.7 KiB | [pg_background_16-1.8-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_background_16-1.8-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_16` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.7 KiB | [pg_background_16-1.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_background_16-1.6-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_16` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.8 KiB | [pg_background_16-1.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_background_16-1.3-3PGDG.rhel10.x86_64.rpm) |
| `pg_background_16` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.7 KiB | [pg_background_16-1.9.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_background_16-1.9.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_background_16` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.5 KiB | [pg_background_16-1.9.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.9.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_16` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_background_16-1.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.8-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_16` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_background_16-1.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.6-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_16` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.4 KiB | [pg_background_16-1.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.3-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 89.2 KiB | [postgresql-16-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 87.4 KiB | [postgresql-16-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 89.0 KiB | [postgresql-16-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 87.7 KiB | [postgresql-16-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 100.1 KiB | [postgresql-16-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 98.5 KiB | [postgresql-16-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 91.9 KiB | [postgresql-16-pg-background_1.9.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 91.1 KiB | [postgresql-16-pg-background_1.9.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-16-pg-background_1.9.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_15` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 56.3 KiB | [pg_background_15-1.9.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_background_15-1.9.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_background_15` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.1 KiB | [pg_background_15-1.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pg_background_15-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.9 KiB | [pg_background_15-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.6 KiB | [pg_background_15-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.3 KiB | [pg_background_15-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.0-1.rhel8.x86_64.rpm) |
| `pg_background_15` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 55.1 KiB | [pg_background_15-1.9.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_background_15-1.9.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_background_15` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.8 KiB | [pg_background_15-1.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.7 KiB | [pg_background_15-1.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_background_15-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.2 KiB | [pg_background_15-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.2 KiB | [pg_background_15-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.7 KiB | [pg_background_15-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.0-1.rhel8.aarch64.rpm) |
| `pg_background_15` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.2 KiB | [pg_background_15-1.9.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_background_15-1.9.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_background_15` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.9 KiB | [pg_background_15-1.9.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-1.9.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_15` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.4 KiB | [pg_background_15-1.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-1.8-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_15` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.2 KiB | [pg_background_15-1.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-1.6-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_15` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.2 KiB | [pg_background_15-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_15` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.9 KiB | [pg_background_15-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.6 KiB | [pg_background_15-1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-1.0-1.rhel9.x86_64.rpm) |
| `pg_background_15` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.3 KiB | [pg_background_15-1.9.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_background_15-1.9.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_background_15` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.1 KiB | [pg_background_15-1.9.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-1.9.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_15` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.4 KiB | [pg_background_15-1.8-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-1.8-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_15` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [pg_background_15-1.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-1.6-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_15` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.6 KiB | [pg_background_15-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_15` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.1 KiB | [pg_background_15-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-1.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.7 KiB | [pg_background_15-1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-1.0-1.rhel9.aarch64.rpm) |
| `pg_background_15` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.4 KiB | [pg_background_15-1.9.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_background_15-1.9.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_background_15` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.2 KiB | [pg_background_15-1.9.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_background_15-1.9.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_15` | `1.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.8 KiB | [pg_background_15-1.8-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_background_15-1.8-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_15` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.8 KiB | [pg_background_15-1.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_background_15-1.6-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_15` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.8 KiB | [pg_background_15-1.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_background_15-1.3-3PGDG.rhel10.x86_64.rpm) |
| `pg_background_15` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.8 KiB | [pg_background_15-1.9.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_background_15-1.9.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_background_15` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.6 KiB | [pg_background_15-1.9.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.9.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_15` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_background_15-1.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.8-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_15` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_background_15-1.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.6-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_15` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.3 KiB | [pg_background_15-1.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.3-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 89.1 KiB | [postgresql-15-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 87.4 KiB | [postgresql-15-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 89.3 KiB | [postgresql-15-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 87.7 KiB | [postgresql-15-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 100.1 KiB | [postgresql-15-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 98.5 KiB | [postgresql-15-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 91.9 KiB | [postgresql-15-pg-background_1.9.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 91.1 KiB | [postgresql-15-pg-background_1.9.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-15-pg-background_1.9.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_14` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 56.2 KiB | [pg_background_14-1.9.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_background_14-1.9.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_background_14` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.1 KiB | [pg_background_14-1.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-1.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_14` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pg_background_14-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_14` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.9 KiB | [pg_background_14-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.6 KiB | [pg_background_14-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-1.0-1.rhel8.x86_64.rpm) |
| `pg_background_14` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 55.1 KiB | [pg_background_14-1.9.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_background_14-1.9.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_background_14` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.8 KiB | [pg_background_14-1.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_14` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.6 KiB | [pg_background_14-1.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_14` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.6 KiB | [pg_background_14-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_14` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.2 KiB | [pg_background_14-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.6 KiB | [pg_background_14-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.0-1.rhel8.aarch64.rpm) |
| `pg_background_14` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.2 KiB | [pg_background_14-1.9.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_background_14-1.9.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_background_14` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.9 KiB | [pg_background_14-1.9.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_background_14-1.9.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_14` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.3 KiB | [pg_background_14-1.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_background_14-1.8-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_14` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.2 KiB | [pg_background_14-1.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_background_14-1.6-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_background_14` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.2 KiB | [pg_background_14-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_background_14-1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_14` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.9 KiB | [pg_background_14-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_background_14-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_14` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.2 KiB | [pg_background_14-1.9.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_background_14-1.9.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_background_14` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.0 KiB | [pg_background_14-1.9.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.9.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_14` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.4 KiB | [pg_background_14-1.8-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.8-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_14` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [pg_background_14-1.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.6-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_background_14` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.6 KiB | [pg_background_14-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.5 KiB | [pg_background_14-1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.0-1.rhel9.aarch64.rpm) |
| `pg_background_14` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.4 KiB | [pg_background_14-1.9.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_background_14-1.9.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_background_14` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.2 KiB | [pg_background_14-1.9.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_background_14-1.9.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_14` | `1.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.7 KiB | [pg_background_14-1.8-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_background_14-1.8-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_14` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.6 KiB | [pg_background_14-1.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_background_14-1.6-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_background_14` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.8 KiB | [pg_background_14-1.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_background_14-1.3-3PGDG.rhel10.x86_64.rpm) |
| `pg_background_14` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.7 KiB | [pg_background_14-1.9.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_background_14-1.9.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_background_14` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.5 KiB | [pg_background_14-1.9.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_background_14-1.9.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_14` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_background_14-1.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_background_14-1.8-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_14` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_background_14-1.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_background_14-1.6-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_14` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.4 KiB | [pg_background_14-1.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_background_14-1.3-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-background` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 88.9 KiB | [postgresql-14-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 87.2 KiB | [postgresql-14-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 89.0 KiB | [postgresql-14-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 87.6 KiB | [postgresql-14-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 100.0 KiB | [postgresql-14-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 98.4 KiB | [postgresql-14-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 91.8 KiB | [postgresql-14-pg-background_1.9.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-background` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 90.9 KiB | [postgresql-14-pg-background_1.9.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-14-pg-background_1.9.2-1PIGSTY~noble_arm64.deb) |

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

> [pg_background: Execute SQL in background worker processes](https://github.com/vibhorkum/pg_background)

Execute arbitrary SQL commands in **background worker processes** within PostgreSQL. Unlike `dblink` (which creates a separate connection), `pg_background` workers run **inside** the database server in **independent transactions**.

**Use Cases:**
- Background maintenance (VACUUM, ANALYZE, REINDEX)
- Asynchronous audit logging
- Long-running ETL pipelines
- Independent notification delivery
- Parallel query patterns

### Quick Start (V2 API)

```sql
CREATE EXTENSION pg_background;

-- Launch a background job
SELECT * FROM pg_background_launch_v2(
  'SELECT count(*) FROM large_table'
) AS handle;
--   pid  |      cookie
-- -------+-------------------
--  12345 | 1234567890123456

-- Retrieve results (one-time consumption)
SELECT * FROM pg_background_result_v2(12345, 1234567890123456) AS (count BIGINT);

-- Fire-and-forget (no result needed)
SELECT * FROM pg_background_submit_v2(
  'INSERT INTO audit_log (ts, event) VALUES (now(), ''system_check'')'
) AS handle;
```


## V2 API Reference

| Function | Returns | Description |
|----------|---------|-------------|
| `pg_background_launch_v2(sql, queue_size)` | `pg_background_handle` | Launch worker, return cookie-protected handle |
| `pg_background_submit_v2(sql, queue_size)` | `pg_background_handle` | Fire-and-forget (no result consumption) |
| `pg_background_result_v2(pid, cookie)` | `SETOF record` | Retrieve results (one-time consumption) |
| `pg_background_detach_v2(pid, cookie)` | `void` | Stop tracking worker (worker continues) |
| `pg_background_cancel_v2(pid, cookie)` | `void` | Request cancellation |
| `pg_background_cancel_v2_grace(pid, cookie, grace_ms)` | `void` | Cancel with grace period |
| `pg_background_wait_v2(pid, cookie)` | `void` | Block until worker completes |
| `pg_background_wait_v2_timeout(pid, cookie, timeout_ms)` | `bool` | Wait with timeout |
| `pg_background_list_v2()` | `SETOF record` | List known workers in current session |
| `pg_background_stats_v2()` | `pg_background_stats` | Session statistics (v1.8+) |
| `pg_background_progress(pct, msg)` | `void` | Report progress from worker (v1.8+) |
| `pg_background_get_progress_v2(pid, cookie)` | `pg_background_progress` | Get worker progress (v1.8+) |

### Cancel vs Detach

| Operation | Stops Execution | Removes Tracking |
|-----------|-----------------|------------------|
| `cancel_v2()` | Yes (best-effort) | No |
| `detach_v2()` | No | Yes |

- Use **`cancel_v2()`** to **stop work** (terminate execution, prevent commit)
- Use **`detach_v2()`** to **stop tracking** (free bookkeeping while worker continues)

### Worker Lifecycle

```sql
-- Cancel a running job
SELECT pg_background_cancel_v2(pid, cookie);

-- Wait for completion
SELECT pg_background_wait_v2(pid, cookie);

-- Wait with timeout (returns true if completed)
SELECT pg_background_wait_v2_timeout(pid, cookie, 5000);

-- List active workers
SELECT * FROM pg_background_list_v2() AS (
  pid int4, cookie int8, launched_at timestamptz,
  user_id oid, queue_size int4, state text,
  sql_preview text, last_error text, consumed bool
);
```

Worker states: `running`, `stopped`, `canceled`, `error`

### Progress Reporting (v1.8+)

```sql
-- From within worker SQL
SELECT pg_background_progress(50, 'Halfway done');

-- From launcher (check progress)
SELECT * FROM pg_background_get_progress_v2(pid, cookie);
```


## GUC Settings (v1.8+)

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pg_background.max_workers` | 16 | Max concurrent workers per session |
| `pg_background.default_queue_size` | 65536 | Default shared memory queue size |
| `pg_background.worker_timeout` | 0 | Worker execution timeout (0 = no limit) |


## V1 API (Legacy)

The v1 API is retained for backward compatibility but lacks cookie-based PID reuse protection:

```sql
SELECT pg_background_launch('VACUUM VERBOSE my_table') AS pid \gset
SELECT * FROM pg_background_result(:pid) AS (result TEXT);
SELECT pg_background_detach(:pid);
```

The V2 API is recommended for production use due to cookie-based PID reuse protection.


## Security Model

- Extension is installed by superusers, with **no PUBLIC grants** by default
- A dedicated `pg_background_worker` NOLOGIN role is created
- Helper functions manage privileges: `grant_pg_background_privileges(role, include_v1)`
- Workers execute as the **launching user** (not superuser)
