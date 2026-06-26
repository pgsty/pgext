---
title: "pg_extra_time"
linkTitle: "pg_extra_time"
description: "Some date time functions and operators that,"
weight: 4220
categories: ["UTIL"]
width: full
---

[**pg_extra_time**](https://github.com/bigsmoke/pg_extra_time) : Some date time functions and operators that,


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4220** | {{< badge content="pg_extra_time" link="https://github.com/bigsmoke/pg_extra_time" >}} | {{< ext "pg_extra_time" >}} | `2.1.0` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgsql_tweaks" >}} {{< ext "periods" >}} {{< ext "temporal_tables" >}} {{< ext "pg_cron" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_extra_time` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.0` | {{< bg "18" "pg_extra_time_18" "green" >}} {{< bg "17" "pg_extra_time_17" "green" >}} {{< bg "16" "pg_extra_time_16" "green" >}} {{< bg "15" "pg_extra_time_15" "green" >}} {{< bg "14" "pg_extra_time_14" "green" >}} | `pg_extra_time_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.0` | {{< bg "18" "postgresql-18-pg-extra-time" "green" >}} {{< bg "17" "postgresql-17-pg-extra-time" "green" >}} {{< bg "16" "postgresql-16-pg-extra-time" "green" >}} {{< bg "15" "postgresql-15-pg-extra-time" "green" >}} {{< bg "14" "postgresql-14-pg-extra-time" "green" >}} | `postgresql-$v-pg-extra-time` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_14 : AVAIL 4" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_14 : AVAIL 4" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_14 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_14 : AVAIL 5" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_14 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.0" "pg_extra_time_14 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pg-extra-time : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pg-extra-time : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pg-extra-time : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pg-extra-time : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pg-extra-time : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pg-extra-time : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pg-extra-time : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pg-extra-time : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pg-extra-time : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pg-extra-time : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pg-extra-time : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_extra_time_18` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.6 KiB | [pg_extra_time_18-2.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_extra_time_18-2.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_extra_time_18` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.8 KiB | [pg_extra_time_18-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_extra_time_18-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_18` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.6 KiB | [pg_extra_time_18-2.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_extra_time_18-2.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_extra_time_18` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.7 KiB | [pg_extra_time_18-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_extra_time_18-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_18` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.9 KiB | [pg_extra_time_18-2.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_extra_time_18-2.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_extra_time_18` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.3 KiB | [pg_extra_time_18-2.0.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_extra_time_18-2.0.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_extra_time_18` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.8 KiB | [pg_extra_time_18-2.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_extra_time_18-2.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_extra_time_18` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.3 KiB | [pg_extra_time_18-2.0.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_extra_time_18-2.0.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_extra_time_18` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.1 KiB | [pg_extra_time_18-2.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_extra_time_18-2.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_extra_time_18` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.5 KiB | [pg_extra_time_18-2.0.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_extra_time_18-2.0.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_extra_time_18` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 34.0 KiB | [pg_extra_time_18-2.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_extra_time_18-2.1.0-1PIGSTY.el10.aarch64.rpm) |
| `pg_extra_time_18` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.5 KiB | [pg_extra_time_18-2.0.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_extra_time_18-2.0.0-1PGDG.rhel10.2.noarch.rpm) |
| `postgresql-18-pg-extra-time` | `2.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.9 KiB | [postgresql-18-pg-extra-time_2.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-18-pg-extra-time_2.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-extra-time` | `2.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 40.9 KiB | [postgresql-18-pg-extra-time_2.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-18-pg-extra-time_2.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-extra-time` | `2.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.9 KiB | [postgresql-18-pg-extra-time_2.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-extra-time/postgresql-18-pg-extra-time_2.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-extra-time` | `2.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 40.9 KiB | [postgresql-18-pg-extra-time_2.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-extra-time/postgresql-18-pg-extra-time_2.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-extra-time` | `2.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 37.3 KiB | [postgresql-18-pg-extra-time_2.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-18-pg-extra-time_2.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-extra-time` | `2.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 37.3 KiB | [postgresql-18-pg-extra-time_2.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-18-pg-extra-time_2.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-extra-time` | `2.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 37.1 KiB | [postgresql-18-pg-extra-time_2.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-18-pg-extra-time_2.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-extra-time` | `2.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 37.1 KiB | [postgresql-18-pg-extra-time_2.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-18-pg-extra-time_2.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-extra-time` | `2.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 37.1 KiB | [postgresql-18-pg-extra-time_2.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-extra-time/postgresql-18-pg-extra-time_2.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-extra-time` | `2.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 37.1 KiB | [postgresql-18-pg-extra-time_2.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-extra-time/postgresql-18-pg-extra-time_2.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_extra_time_17` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.6 KiB | [pg_extra_time_17-2.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_extra_time_17-2.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_extra_time_17` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.8 KiB | [pg_extra_time_17-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_extra_time_17-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_17` | `1.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.8 KiB | [pg_extra_time_17-1.1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_extra_time_17-1.1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_17` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.6 KiB | [pg_extra_time_17-2.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_extra_time_17-2.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_extra_time_17` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.7 KiB | [pg_extra_time_17-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_extra_time_17-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_17` | `1.1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.8 KiB | [pg_extra_time_17-1.1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_extra_time_17-1.1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_17` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.9 KiB | [pg_extra_time_17-2.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_extra_time_17-2.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_extra_time_17` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.3 KiB | [pg_extra_time_17-2.0.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_extra_time_17-2.0.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_extra_time_17` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.8 KiB | [pg_extra_time_17-2.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_extra_time_17-2.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_extra_time_17` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.3 KiB | [pg_extra_time_17-2.0.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_extra_time_17-2.0.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_extra_time_17` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.1 KiB | [pg_extra_time_17-2.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_extra_time_17-2.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_extra_time_17` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.5 KiB | [pg_extra_time_17-2.0.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_extra_time_17-2.0.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_extra_time_17` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 34.0 KiB | [pg_extra_time_17-2.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_extra_time_17-2.1.0-1PIGSTY.el10.aarch64.rpm) |
| `pg_extra_time_17` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.5 KiB | [pg_extra_time_17-2.0.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_extra_time_17-2.0.0-1PGDG.rhel10.2.noarch.rpm) |
| `postgresql-17-pg-extra-time` | `2.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.9 KiB | [postgresql-17-pg-extra-time_2.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-extra-time` | `2.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 40.9 KiB | [postgresql-17-pg-extra-time_2.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-extra-time` | `2.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.9 KiB | [postgresql-17-pg-extra-time_2.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-extra-time` | `2.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 40.9 KiB | [postgresql-17-pg-extra-time_2.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-extra-time` | `2.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 37.3 KiB | [postgresql-17-pg-extra-time_2.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-extra-time` | `2.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 37.3 KiB | [postgresql-17-pg-extra-time_2.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-extra-time` | `2.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 37.1 KiB | [postgresql-17-pg-extra-time_2.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-extra-time` | `2.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 37.1 KiB | [postgresql-17-pg-extra-time_2.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-extra-time` | `2.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 37.1 KiB | [postgresql-17-pg-extra-time_2.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-extra-time` | `2.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 37.1 KiB | [postgresql-17-pg-extra-time_2.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_extra_time_16` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.6 KiB | [pg_extra_time_16-2.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_extra_time_16-2.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_extra_time_16` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.8 KiB | [pg_extra_time_16-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_extra_time_16-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_16` | `1.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.8 KiB | [pg_extra_time_16-1.1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_extra_time_16-1.1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_16` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.3 KiB | [pg_extra_time_16-1.1.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_extra_time_16-1.1.2-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_16` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.6 KiB | [pg_extra_time_16-2.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_extra_time_16-2.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_extra_time_16` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.7 KiB | [pg_extra_time_16-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_extra_time_16-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_16` | `1.1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.8 KiB | [pg_extra_time_16-1.1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_extra_time_16-1.1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_16` | `1.1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.3 KiB | [pg_extra_time_16-1.1.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_extra_time_16-1.1.2-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_16` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.9 KiB | [pg_extra_time_16-2.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_extra_time_16-2.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_extra_time_16` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.3 KiB | [pg_extra_time_16-2.0.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_extra_time_16-2.0.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_extra_time_16` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.8 KiB | [pg_extra_time_16-2.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_extra_time_16-2.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_extra_time_16` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.3 KiB | [pg_extra_time_16-2.0.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_extra_time_16-2.0.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_extra_time_16` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.1 KiB | [pg_extra_time_16-2.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_extra_time_16-2.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_extra_time_16` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.5 KiB | [pg_extra_time_16-2.0.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_extra_time_16-2.0.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_extra_time_16` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 34.0 KiB | [pg_extra_time_16-2.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_extra_time_16-2.1.0-1PIGSTY.el10.aarch64.rpm) |
| `pg_extra_time_16` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.5 KiB | [pg_extra_time_16-2.0.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_extra_time_16-2.0.0-1PGDG.rhel10.2.noarch.rpm) |
| `postgresql-16-pg-extra-time` | `2.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.9 KiB | [postgresql-16-pg-extra-time_2.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-extra-time` | `2.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 40.9 KiB | [postgresql-16-pg-extra-time_2.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-extra-time` | `2.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.9 KiB | [postgresql-16-pg-extra-time_2.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-extra-time` | `2.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 40.9 KiB | [postgresql-16-pg-extra-time_2.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-extra-time` | `2.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 37.3 KiB | [postgresql-16-pg-extra-time_2.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-extra-time` | `2.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 37.3 KiB | [postgresql-16-pg-extra-time_2.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-extra-time` | `2.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 37.1 KiB | [postgresql-16-pg-extra-time_2.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-extra-time` | `2.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 37.1 KiB | [postgresql-16-pg-extra-time_2.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-extra-time` | `2.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 37.1 KiB | [postgresql-16-pg-extra-time_2.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-extra-time` | `2.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 37.1 KiB | [postgresql-16-pg-extra-time_2.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_extra_time_15` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.6 KiB | [pg_extra_time_15-2.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_extra_time_15-2.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_extra_time_15` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.8 KiB | [pg_extra_time_15-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_extra_time_15-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_15` | `1.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.8 KiB | [pg_extra_time_15-1.1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_extra_time_15-1.1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_15` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.3 KiB | [pg_extra_time_15-1.1.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_extra_time_15-1.1.2-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_15` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.6 KiB | [pg_extra_time_15-2.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_extra_time_15-2.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_extra_time_15` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.7 KiB | [pg_extra_time_15-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_extra_time_15-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_15` | `1.1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.8 KiB | [pg_extra_time_15-1.1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_extra_time_15-1.1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_15` | `1.1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.3 KiB | [pg_extra_time_15-1.1.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_extra_time_15-1.1.2-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_15` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.9 KiB | [pg_extra_time_15-2.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_extra_time_15-2.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_extra_time_15` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.3 KiB | [pg_extra_time_15-2.0.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_extra_time_15-2.0.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_extra_time_15` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.8 KiB | [pg_extra_time_15-2.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_extra_time_15-2.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_extra_time_15` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.3 KiB | [pg_extra_time_15-2.0.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_extra_time_15-2.0.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_extra_time_15` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.1 KiB | [pg_extra_time_15-2.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_extra_time_15-2.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_extra_time_15` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.5 KiB | [pg_extra_time_15-2.0.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_extra_time_15-2.0.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_extra_time_15` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 34.0 KiB | [pg_extra_time_15-2.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_extra_time_15-2.1.0-1PIGSTY.el10.aarch64.rpm) |
| `pg_extra_time_15` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.5 KiB | [pg_extra_time_15-2.0.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_extra_time_15-2.0.0-1PGDG.rhel10.2.noarch.rpm) |
| `postgresql-15-pg-extra-time` | `2.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.9 KiB | [postgresql-15-pg-extra-time_2.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-extra-time` | `2.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 40.9 KiB | [postgresql-15-pg-extra-time_2.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-extra-time` | `2.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.9 KiB | [postgresql-15-pg-extra-time_2.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-extra-time` | `2.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 40.9 KiB | [postgresql-15-pg-extra-time_2.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-extra-time` | `2.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 37.2 KiB | [postgresql-15-pg-extra-time_2.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-extra-time` | `2.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 37.2 KiB | [postgresql-15-pg-extra-time_2.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-extra-time` | `2.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 37.1 KiB | [postgresql-15-pg-extra-time_2.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-extra-time` | `2.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 37.1 KiB | [postgresql-15-pg-extra-time_2.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-extra-time` | `2.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 37.1 KiB | [postgresql-15-pg-extra-time_2.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-extra-time` | `2.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 37.1 KiB | [postgresql-15-pg-extra-time_2.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_extra_time_14` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.6 KiB | [pg_extra_time_14-2.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_extra_time_14-2.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_extra_time_14` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.8 KiB | [pg_extra_time_14-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_extra_time_14-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_14` | `1.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.8 KiB | [pg_extra_time_14-1.1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_extra_time_14-1.1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_14` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.3 KiB | [pg_extra_time_14-1.1.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_extra_time_14-1.1.2-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_14` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.6 KiB | [pg_extra_time_14-2.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_extra_time_14-2.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_extra_time_14` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.7 KiB | [pg_extra_time_14-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_extra_time_14-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_14` | `1.1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.8 KiB | [pg_extra_time_14-1.1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_extra_time_14-1.1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_14` | `1.1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.3 KiB | [pg_extra_time_14-1.1.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_extra_time_14-1.1.2-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_14` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.9 KiB | [pg_extra_time_14-2.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_extra_time_14-2.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_extra_time_14` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.3 KiB | [pg_extra_time_14-2.0.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_extra_time_14-2.0.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_extra_time_14` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.8 KiB | [pg_extra_time_14-2.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_extra_time_14-2.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_extra_time_14` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.3 KiB | [pg_extra_time_14-2.0.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_extra_time_14-2.0.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_extra_time_14` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.3 KiB | [pg_extra_time_14-2.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_extra_time_14-2.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_14` | `1.1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.6 KiB | [pg_extra_time_14-1.1.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_extra_time_14-1.1.3-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_14` | `1.1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.0 KiB | [pg_extra_time_14-1.1.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_extra_time_14-1.1.2-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_14` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.1 KiB | [pg_extra_time_14-2.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_extra_time_14-2.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_extra_time_14` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.5 KiB | [pg_extra_time_14-2.0.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_extra_time_14-2.0.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_extra_time_14` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 34.0 KiB | [pg_extra_time_14-2.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_extra_time_14-2.1.0-1PIGSTY.el10.aarch64.rpm) |
| `pg_extra_time_14` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.5 KiB | [pg_extra_time_14-2.0.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_extra_time_14-2.0.0-1PGDG.rhel10.2.noarch.rpm) |
| `postgresql-14-pg-extra-time` | `2.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.9 KiB | [postgresql-14-pg-extra-time_2.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-extra-time` | `2.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 40.9 KiB | [postgresql-14-pg-extra-time_2.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-extra-time` | `2.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.9 KiB | [postgresql-14-pg-extra-time_2.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-extra-time` | `2.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 40.9 KiB | [postgresql-14-pg-extra-time_2.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-extra-time` | `2.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 37.3 KiB | [postgresql-14-pg-extra-time_2.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-extra-time` | `2.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 37.3 KiB | [postgresql-14-pg-extra-time_2.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-extra-time` | `2.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 37.1 KiB | [postgresql-14-pg-extra-time_2.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-extra-time` | `2.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 37.1 KiB | [postgresql-14-pg-extra-time_2.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-extra-time` | `2.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 37.1 KiB | [postgresql-14-pg-extra-time_2.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-extra-time` | `2.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 37.1 KiB | [postgresql-14-pg-extra-time_2.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bigsmoke/pg_extra_time" title="Repository" icon="github" subtitle="github.com/bigsmoke/pg_extra_time" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_extra_time-2.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_extra_time;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_extra_time;		# install via package name, for the active PG version

pig install pg_extra_time -v 18;   # install for PG 18
pig install pg_extra_time -v 17;   # install for PG 17
pig install pg_extra_time -v 16;   # install for PG 16
pig install pg_extra_time -v 15;   # install for PG 15
pig install pg_extra_time -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_extra_time;
```


## Usage

> Sources: [pg_extra_time upstream README](https://github.com/bigsmoke/pg_extra_time), [PGXN pg_extra_time](https://pgxn.org/dist/pg_extra_time/), [local metadata](../db/extension.csv).

`pg_extra_time` provides small SQL functions and casts for date/time, interval, and range calculations that are awkward with PostgreSQL core functions alone.

```sql
CREATE EXTENSION pg_extra_time;
```

### Convert to Seconds (float)

Use `to_float(...)` or explicit casts to `float`/`double precision` for timestamps, timestamp ranges, and intervals. Timestamp values are measured from the Unix epoch; ranges and intervals are measured by duration in seconds.

```sql
SELECT to_float('1970-01-01 00:00:00+0'::timestamptz);  -- 0.0
SELECT to_float('1970-01-01 00:00:00+0'::timestamp);    -- 0.0
SELECT to_float('1 day 1 sec'::interval);                -- 86401.0
SELECT to_float('[2024-06-06 05:58:00,2024-06-06 06:00:10]'::tstzrange);  -- 130.0
SELECT to_float('[2024-06-06 05:58:00,2024-06-06 06:00:10]'::tsrange);    -- 130.0
```

Cast syntax also works:

```sql
SELECT '1970-01-01 01:03:01+00'::timestamptz::float;    -- 3181.00
SELECT '1 day 1 sec 200 ms'::interval::float;            -- 86401.2
SELECT '[epoch,1970-01-01T01:03:01+00]'::tstzrange::float;  -- 3181.00
```

### Convert to Days

Use `days(...)` when fractions matter and `whole_days(...)` when an integer number of complete days is needed.

```sql
SELECT days('[2024-06-06,2024-06-09)'::daterange);       -- 2
SELECT days('[2024-06-06,2024-06-08 06:00]'::tstzrange);  -- 3.25 (fractional days)
SELECT whole_days('[2024-06-06,2024-06-08 18:00]'::tstzrange);  -- 2
SELECT days('10 days 12 hours'::interval);                -- 10.5
SELECT whole_days('10 days 20 hours'::interval);          -- 10
```

`whole_days(interval)` handles negative intervals by applying the sign after flooring the absolute day count.

### Count Date Parts

`date_part_parts(part, subpart, timestamp with time zone, timezone)` returns how many smaller date parts exist in a larger date part at a given timestamp and timezone. This helps with calculations where a day is not always 24 hours because of DST.

```sql
SELECT date_part_parts('month', 'days', '2024-02-12'::timestamptz, 'UTC');  -- 29
SELECT date_part_parts('year', 'days', '2024-08-23'::timestamptz, 'UTC');   -- 366
```

### Build And Split Ranges

Use `make_tstzrange` or `make_tsrange` to build ranges from a timestamp and interval, including negative intervals.

```sql
SELECT make_tstzrange('2024-01-05 00:00+00'::timestamptz, '-4 days'::interval);
SELECT make_tsrange('2024-01-01 00:00'::timestamp, '4 days'::interval, '[)');
```

`each_subperiod(tstzrange, interval, round_remainder integer DEFAULT 0)` splits a timestamp range into interval-sized chunks. The remainder policy is: `1` rounds up to a full chunk, `0` keeps a partial final chunk, and `-1` discards the remainder.

```sql
SELECT *
FROM each_subperiod(
  '[2023-01-01,2023-04-02)'::tstzrange,
  '1 month'::interval,
  0
);
```

### Extract And Remainder Intervals

`to_interval(tstzrange)` extracts an interval from a timestamp range using month, day, and microsecond units. `to_interval(tstzrange, interval[])` accepts explicit units in greatest-first order and rounds down by discarding the remainder.

```sql
SELECT to_interval('[2024-01-01,2024-01-05]'::tstzrange);  -- 4 days
SELECT to_interval(
  '[2024-01-01,2024-04-13 01:10]'::tstzrange,
  ARRAY['1 mon'::interval, '1 day'::interval, '1 hour'::interval]
);
```

Use `%` or `modulo(...)` when the remainder matters.

```sql
SELECT '10 seconds 100 milliseconds'::interval % '3 seconds'::interval;
SELECT '[2024-01-01,2024-01-10)'::tstzrange % '4 days'::interval;
```

### Caveats

`to_float(tstzrange)` and `to_float(tsrange)` return positive or negative infinity for unbounded ranges and `0` for empty ranges. Integer casts are intentionally not provided; use `whole_days(...)` when you need integer days. Deprecated aliases such as `extract_days(interval)` and `extract_interval(tstzrange, interval[])` remain for compatibility, but upstream recommends `whole_days(...)` and `to_interval(...)` instead.

### Reference

Common public functions:

| Function | Use |
|----------|-----|
| `current_timezone()` | Return the active `pg_timezone_names` row |
| `date_part_parts(...)` | Count smaller date parts inside larger date parts |
| `days(...)` | Fractional or integer day count, depending on input type |
| `whole_days(...)` | Whole days from intervals or timestamp ranges |
| `to_float(...)` | Seconds from timestamps, timestamp ranges, or intervals |
| `to_interval(...)` | Interval extracted from a `tstzrange` |
| `make_tsrange(...)` / `make_tstzrange(...)` | Build ranges from timestamp plus interval |
| `each_subperiod(...)` | Split a `tstzrange` into subranges |
| `modulo(...)` / `%` | Remainder after dividing intervals or ranges |
