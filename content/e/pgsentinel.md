---
title: "pgsentinel"
linkTitle: "pgsentinel"
description: "active session history"
weight: 6410
categories: ["STAT"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PGDG"]
page_width: full
---

[**pgsentinel**](https://github.com/pgsentinel/pgsentinel) : active session history


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6410** | {{< badge content="pgsentinel" link="https://github.com/pgsentinel/pgsentinel" >}} | {{< ext "pgsentinel" >}} | `1.5.1` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_datasentinel" >}} {{< ext "pg_profile" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_wait_sampling" >}} {{< ext "pg_stat_ch" >}} {{< ext "pgmonitor" >}} {{< ext "system_stats" >}} {{< ext "pgnodemx" >}} {{< ext "powa" >}} {{< ext "pg_stat_statements" >}} |

> [!Note] Requires preload alongside pg_stat_statements.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.5.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgsentinel` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.5.1` | {{< bg "18" "pgsentinel_18" "green" >}} {{< bg "17" "pgsentinel_17" "green" >}} {{< bg "16" "pgsentinel_16" "green" >}} {{< bg "15" "pgsentinel_15" "green" >}} {{< bg "14" "pgsentinel_14" "green" >}} | `pgsentinel_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.5.1` | {{< bg "18" "postgresql-18-pgsentinel" "green" >}} {{< bg "17" "postgresql-17-pgsentinel" "green" >}} {{< bg "16" "postgresql-16-pgsentinel" "green" >}} {{< bg "15" "postgresql-15-pgsentinel" "green" >}} {{< bg "14" "postgresql-14-pgsentinel" "green" >}} | `postgresql-$v-pgsentinel` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_14 : AVAIL 4" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_14 : AVAIL 4" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_14 : AVAIL 7" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_14 : AVAIL 7" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_14 : AVAIL 7" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.5.1" "pgsentinel_14 : AVAIL 7" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-pgsentinel : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-pgsentinel : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-pgsentinel : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-pgsentinel : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-pgsentinel : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-pgsentinel : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-pgsentinel : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-pgsentinel : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-pgsentinel : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-pgsentinel : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-pgsentinel : AVAIL 3" "blue" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsentinel_18` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.5 KiB | [pgsentinel_18-1.5.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgsentinel_18-1.5.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.7 KiB | [pgsentinel_18-1.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgsentinel_18-1.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_18` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.2 KiB | [pgsentinel_18-1.3.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgsentinel_18-1.3.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_18` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.6 KiB | [pgsentinel_18-1.5.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgsentinel_18-1.5.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.9 KiB | [pgsentinel_18-1.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgsentinel_18-1.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_18` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.4 KiB | [pgsentinel_18-1.3.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgsentinel_18-1.3.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_18` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.9 KiB | [pgsentinel_18-1.5.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgsentinel_18-1.5.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pgsentinel_18` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.3 KiB | [pgsentinel_18-1.4.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgsentinel_18-1.4.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.0 KiB | [pgsentinel_18-1.4.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgsentinel_18-1.4.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.1 KiB | [pgsentinel_18-1.4.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgsentinel_18-1.4.0-1PGDG.rhel9.6.x86_64.rpm) |
| `pgsentinel_18` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.5 KiB | [pgsentinel_18-1.3.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgsentinel_18-1.3.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_18` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.7 KiB | [pgsentinel_18-1.3.1-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgsentinel_18-1.3.1-1PGDG.rhel9.6.x86_64.rpm) |
| `pgsentinel_18` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.1 KiB | [pgsentinel_18-1.5.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgsentinel_18-1.5.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pgsentinel_18` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.6 KiB | [pgsentinel_18-1.4.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgsentinel_18-1.4.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.2 KiB | [pgsentinel_18-1.4.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgsentinel_18-1.4.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.4 KiB | [pgsentinel_18-1.4.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgsentinel_18-1.4.0-1PGDG.rhel9.6.aarch64.rpm) |
| `pgsentinel_18` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.8 KiB | [pgsentinel_18-1.3.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgsentinel_18-1.3.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_18` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.9 KiB | [pgsentinel_18-1.3.1-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgsentinel_18-1.3.1-1PGDG.rhel9.6.aarch64.rpm) |
| `pgsentinel_18` | `1.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.4 KiB | [pgsentinel_18-1.5.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgsentinel_18-1.5.1-1PGDG.rhel10.2.x86_64.rpm) |
| `pgsentinel_18` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.8 KiB | [pgsentinel_18-1.4.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgsentinel_18-1.4.1-1PGDG.rhel10.2.x86_64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.5 KiB | [pgsentinel_18-1.4.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgsentinel_18-1.4.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.9 KiB | [pgsentinel_18-1.4.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgsentinel_18-1.4.0-1PGDG.rhel10.0.x86_64.rpm) |
| `pgsentinel_18` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.9 KiB | [pgsentinel_18-1.3.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgsentinel_18-1.3.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_18` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.3 KiB | [pgsentinel_18-1.3.1-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgsentinel_18-1.3.1-1PGDG.rhel10.0.x86_64.rpm) |
| `pgsentinel_18` | `1.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.3 KiB | [pgsentinel_18-1.5.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsentinel_18-1.5.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pgsentinel_18` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.8 KiB | [pgsentinel_18-1.4.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsentinel_18-1.4.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.4 KiB | [pgsentinel_18-1.4.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsentinel_18-1.4.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.4 KiB | [pgsentinel_18-1.4.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsentinel_18-1.4.0-1PGDG.rhel10.0.aarch64.rpm) |
| `pgsentinel_18` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.0 KiB | [pgsentinel_18-1.3.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsentinel_18-1.3.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pgsentinel_18` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.0 KiB | [pgsentinel_18-1.3.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsentinel_18-1.3.1-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-18-pgsentinel` | `1.5.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.8 KiB | [postgresql-18-pgsentinel_1.5.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.1-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.6 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg12+3_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.6 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg12+2_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.5.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.8 KiB | [postgresql-18-pgsentinel_1.5.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.1-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.5 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg12+3_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.5 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg12+2_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.5.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.8 KiB | [postgresql-18-pgsentinel_1.5.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.1-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.6 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg13+3_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.6 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg13+2_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.5.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 45.0 KiB | [postgresql-18-pgsentinel_1.5.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.1-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.7 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg13+3_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.7 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg13+2_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.5.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 46.9 KiB | [postgresql-18-pgsentinel_1.5.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 46.6 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg22.04+3_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 46.6 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg22.04+2_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.5.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 45.7 KiB | [postgresql-18-pgsentinel_1.5.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 45.5 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg22.04+3_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 45.5 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg22.04+2_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.5.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.7 KiB | [postgresql-18-pgsentinel_1.5.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.5 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg24.04+3_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.5 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg24.04+2_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.5.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.9 KiB | [postgresql-18-pgsentinel_1.5.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.6 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg24.04+3_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.6 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg24.04+2_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.5.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 46.0 KiB | [postgresql-18-pgsentinel_1.5.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 45.8 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg26.04+3_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 45.8 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg26.04+2_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.5.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 44.9 KiB | [postgresql-18-pgsentinel_1.5.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 44.7 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg26.04+3_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 44.7 KiB | [postgresql-18-pgsentinel_1.5.0-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.5.0-1.pgdg26.04+2_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsentinel_17` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.5 KiB | [pgsentinel_17-1.5.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsentinel_17-1.5.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.7 KiB | [pgsentinel_17-1.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsentinel_17-1.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_17` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.3 KiB | [pgsentinel_17-1.3.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsentinel_17-1.3.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_17` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.6 KiB | [pgsentinel_17-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsentinel_17-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsentinel_17` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.6 KiB | [pgsentinel_17-1.5.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsentinel_17-1.5.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.0 KiB | [pgsentinel_17-1.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsentinel_17-1.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_17` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.4 KiB | [pgsentinel_17-1.3.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsentinel_17-1.3.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_17` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.8 KiB | [pgsentinel_17-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsentinel_17-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsentinel_17` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.9 KiB | [pgsentinel_17-1.5.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsentinel_17-1.5.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pgsentinel_17` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.3 KiB | [pgsentinel_17-1.4.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsentinel_17-1.4.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.0 KiB | [pgsentinel_17-1.4.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsentinel_17-1.4.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.1 KiB | [pgsentinel_17-1.4.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsentinel_17-1.4.0-1PGDG.rhel9.6.x86_64.rpm) |
| `pgsentinel_17` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.6 KiB | [pgsentinel_17-1.3.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsentinel_17-1.3.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_17` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.7 KiB | [pgsentinel_17-1.3.1-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsentinel_17-1.3.1-1PGDG.rhel9.6.x86_64.rpm) |
| `pgsentinel_17` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.1 KiB | [pgsentinel_17-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsentinel_17-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsentinel_17` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.2 KiB | [pgsentinel_17-1.5.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsentinel_17-1.5.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pgsentinel_17` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.6 KiB | [pgsentinel_17-1.4.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsentinel_17-1.4.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.3 KiB | [pgsentinel_17-1.4.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsentinel_17-1.4.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.4 KiB | [pgsentinel_17-1.4.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsentinel_17-1.4.0-1PGDG.rhel9.6.aarch64.rpm) |
| `pgsentinel_17` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.9 KiB | [pgsentinel_17-1.3.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsentinel_17-1.3.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_17` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.0 KiB | [pgsentinel_17-1.3.1-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsentinel_17-1.3.1-1PGDG.rhel9.6.aarch64.rpm) |
| `pgsentinel_17` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.2 KiB | [pgsentinel_17-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsentinel_17-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsentinel_17` | `1.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.5 KiB | [pgsentinel_17-1.5.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgsentinel_17-1.5.1-1PGDG.rhel10.2.x86_64.rpm) |
| `pgsentinel_17` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.9 KiB | [pgsentinel_17-1.4.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgsentinel_17-1.4.1-1PGDG.rhel10.2.x86_64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.5 KiB | [pgsentinel_17-1.4.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgsentinel_17-1.4.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.9 KiB | [pgsentinel_17-1.4.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgsentinel_17-1.4.0-1PGDG.rhel10.0.x86_64.rpm) |
| `pgsentinel_17` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.9 KiB | [pgsentinel_17-1.3.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgsentinel_17-1.3.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_17` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.3 KiB | [pgsentinel_17-1.3.1-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgsentinel_17-1.3.1-1PGDG.rhel10.0.x86_64.rpm) |
| `pgsentinel_17` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.7 KiB | [pgsentinel_17-1.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgsentinel_17-1.2.0-1PGDG.rhel10.x86_64.rpm) |
| `pgsentinel_17` | `1.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.4 KiB | [pgsentinel_17-1.5.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsentinel_17-1.5.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pgsentinel_17` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.8 KiB | [pgsentinel_17-1.4.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsentinel_17-1.4.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.5 KiB | [pgsentinel_17-1.4.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsentinel_17-1.4.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.5 KiB | [pgsentinel_17-1.4.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsentinel_17-1.4.0-1PGDG.rhel10.0.aarch64.rpm) |
| `pgsentinel_17` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.1 KiB | [pgsentinel_17-1.3.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsentinel_17-1.3.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pgsentinel_17` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.1 KiB | [pgsentinel_17-1.3.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsentinel_17-1.3.1-1PGDG.rhel10.0.aarch64.rpm) |
| `pgsentinel_17` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.8 KiB | [pgsentinel_17-1.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsentinel_17-1.2.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgsentinel` | `1.5.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.8 KiB | [postgresql-17-pgsentinel_1.5.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.1-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.6 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg12+3_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.6 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg12+2_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.5.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.9 KiB | [postgresql-17-pgsentinel_1.5.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.1-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.6 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg12+3_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.6 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg12+2_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.5.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.9 KiB | [postgresql-17-pgsentinel_1.5.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.1-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.7 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg13+3_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.7 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg13+2_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.5.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 45.0 KiB | [postgresql-17-pgsentinel_1.5.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.1-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.7 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg13+3_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.7 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg13+2_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.5.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 54.5 KiB | [postgresql-17-pgsentinel_1.5.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 54.3 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg22.04+3_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 54.3 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg22.04+2_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.5.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 53.4 KiB | [postgresql-17-pgsentinel_1.5.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 53.0 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg22.04+3_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 53.0 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg22.04+2_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.5.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.8 KiB | [postgresql-17-pgsentinel_1.5.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.5 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg24.04+3_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.5 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg24.04+2_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.5.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 45.0 KiB | [postgresql-17-pgsentinel_1.5.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.7 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg24.04+3_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.7 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg24.04+2_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.5.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 46.1 KiB | [postgresql-17-pgsentinel_1.5.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 45.8 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg26.04+3_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 45.8 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg26.04+2_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.5.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 45.0 KiB | [postgresql-17-pgsentinel_1.5.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 44.8 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg26.04+3_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 44.7 KiB | [postgresql-17-pgsentinel_1.5.0-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.5.0-1.pgdg26.04+2_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsentinel_16` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.5 KiB | [pgsentinel_16-1.5.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsentinel_16-1.5.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.7 KiB | [pgsentinel_16-1.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsentinel_16-1.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_16` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.2 KiB | [pgsentinel_16-1.3.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsentinel_16-1.3.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_16` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.6 KiB | [pgsentinel_16-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsentinel_16-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsentinel_16` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.6 KiB | [pgsentinel_16-1.5.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsentinel_16-1.5.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.9 KiB | [pgsentinel_16-1.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsentinel_16-1.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_16` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.4 KiB | [pgsentinel_16-1.3.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsentinel_16-1.3.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_16` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.8 KiB | [pgsentinel_16-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsentinel_16-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsentinel_16` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.9 KiB | [pgsentinel_16-1.5.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsentinel_16-1.5.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pgsentinel_16` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.3 KiB | [pgsentinel_16-1.4.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsentinel_16-1.4.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.0 KiB | [pgsentinel_16-1.4.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsentinel_16-1.4.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.1 KiB | [pgsentinel_16-1.4.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsentinel_16-1.4.0-1PGDG.rhel9.6.x86_64.rpm) |
| `pgsentinel_16` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.6 KiB | [pgsentinel_16-1.3.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsentinel_16-1.3.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_16` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.7 KiB | [pgsentinel_16-1.3.1-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsentinel_16-1.3.1-1PGDG.rhel9.6.x86_64.rpm) |
| `pgsentinel_16` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.1 KiB | [pgsentinel_16-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsentinel_16-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsentinel_16` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.2 KiB | [pgsentinel_16-1.5.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsentinel_16-1.5.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pgsentinel_16` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.6 KiB | [pgsentinel_16-1.4.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsentinel_16-1.4.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.3 KiB | [pgsentinel_16-1.4.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsentinel_16-1.4.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.4 KiB | [pgsentinel_16-1.4.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsentinel_16-1.4.0-1PGDG.rhel9.6.aarch64.rpm) |
| `pgsentinel_16` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.8 KiB | [pgsentinel_16-1.3.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsentinel_16-1.3.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_16` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.0 KiB | [pgsentinel_16-1.3.1-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsentinel_16-1.3.1-1PGDG.rhel9.6.aarch64.rpm) |
| `pgsentinel_16` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.2 KiB | [pgsentinel_16-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsentinel_16-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsentinel_16` | `1.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.5 KiB | [pgsentinel_16-1.5.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgsentinel_16-1.5.1-1PGDG.rhel10.2.x86_64.rpm) |
| `pgsentinel_16` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.9 KiB | [pgsentinel_16-1.4.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgsentinel_16-1.4.1-1PGDG.rhel10.2.x86_64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.5 KiB | [pgsentinel_16-1.4.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgsentinel_16-1.4.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.9 KiB | [pgsentinel_16-1.4.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgsentinel_16-1.4.0-1PGDG.rhel10.0.x86_64.rpm) |
| `pgsentinel_16` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.9 KiB | [pgsentinel_16-1.3.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgsentinel_16-1.3.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_16` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.3 KiB | [pgsentinel_16-1.3.1-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgsentinel_16-1.3.1-1PGDG.rhel10.0.x86_64.rpm) |
| `pgsentinel_16` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.7 KiB | [pgsentinel_16-1.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgsentinel_16-1.2.0-1PGDG.rhel10.x86_64.rpm) |
| `pgsentinel_16` | `1.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.4 KiB | [pgsentinel_16-1.5.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsentinel_16-1.5.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pgsentinel_16` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.8 KiB | [pgsentinel_16-1.4.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsentinel_16-1.4.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.5 KiB | [pgsentinel_16-1.4.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsentinel_16-1.4.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.5 KiB | [pgsentinel_16-1.4.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsentinel_16-1.4.0-1PGDG.rhel10.0.aarch64.rpm) |
| `pgsentinel_16` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.1 KiB | [pgsentinel_16-1.3.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsentinel_16-1.3.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pgsentinel_16` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.1 KiB | [pgsentinel_16-1.3.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsentinel_16-1.3.1-1PGDG.rhel10.0.aarch64.rpm) |
| `pgsentinel_16` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.8 KiB | [pgsentinel_16-1.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsentinel_16-1.2.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgsentinel` | `1.5.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.8 KiB | [postgresql-16-pgsentinel_1.5.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.1-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.5 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg12+3_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.5 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg12+2_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.5.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.8 KiB | [postgresql-16-pgsentinel_1.5.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.1-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.6 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg12+3_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.6 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg12+2_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.5.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.8 KiB | [postgresql-16-pgsentinel_1.5.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.1-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.5 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg13+3_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.5 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg13+2_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.5.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 45.0 KiB | [postgresql-16-pgsentinel_1.5.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.1-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.7 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg13+3_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.7 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg13+2_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.5.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 54.3 KiB | [postgresql-16-pgsentinel_1.5.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 54.1 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg22.04+3_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 54.1 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg22.04+2_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.5.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 53.1 KiB | [postgresql-16-pgsentinel_1.5.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 52.9 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg22.04+3_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 52.9 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg22.04+2_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.5.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.7 KiB | [postgresql-16-pgsentinel_1.5.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.5 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg24.04+3_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.5 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg24.04+2_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.5.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.9 KiB | [postgresql-16-pgsentinel_1.5.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.6 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg24.04+3_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.6 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg24.04+2_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.5.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 46.1 KiB | [postgresql-16-pgsentinel_1.5.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 45.8 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg26.04+3_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 45.8 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg26.04+2_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.5.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 44.9 KiB | [postgresql-16-pgsentinel_1.5.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 44.7 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg26.04+3_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 44.7 KiB | [postgresql-16-pgsentinel_1.5.0-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.5.0-1.pgdg26.04+2_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsentinel_15` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.6 KiB | [pgsentinel_15-1.5.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsentinel_15-1.5.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.8 KiB | [pgsentinel_15-1.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsentinel_15-1.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_15` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.4 KiB | [pgsentinel_15-1.3.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsentinel_15-1.3.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_15` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.8 KiB | [pgsentinel_15-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsentinel_15-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsentinel_15` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.6 KiB | [pgsentinel_15-1.5.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsentinel_15-1.5.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.9 KiB | [pgsentinel_15-1.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsentinel_15-1.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_15` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.4 KiB | [pgsentinel_15-1.3.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsentinel_15-1.3.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_15` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.8 KiB | [pgsentinel_15-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsentinel_15-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsentinel_15` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.1 KiB | [pgsentinel_15-1.5.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsentinel_15-1.5.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pgsentinel_15` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.5 KiB | [pgsentinel_15-1.4.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsentinel_15-1.4.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.2 KiB | [pgsentinel_15-1.4.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsentinel_15-1.4.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.3 KiB | [pgsentinel_15-1.4.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsentinel_15-1.4.0-1PGDG.rhel9.6.x86_64.rpm) |
| `pgsentinel_15` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.7 KiB | [pgsentinel_15-1.3.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsentinel_15-1.3.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_15` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.8 KiB | [pgsentinel_15-1.3.1-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsentinel_15-1.3.1-1PGDG.rhel9.6.x86_64.rpm) |
| `pgsentinel_15` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.2 KiB | [pgsentinel_15-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsentinel_15-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsentinel_15` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.3 KiB | [pgsentinel_15-1.5.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsentinel_15-1.5.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pgsentinel_15` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.8 KiB | [pgsentinel_15-1.4.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsentinel_15-1.4.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.5 KiB | [pgsentinel_15-1.4.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsentinel_15-1.4.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.6 KiB | [pgsentinel_15-1.4.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsentinel_15-1.4.0-1PGDG.rhel9.6.aarch64.rpm) |
| `pgsentinel_15` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.0 KiB | [pgsentinel_15-1.3.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsentinel_15-1.3.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_15` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.1 KiB | [pgsentinel_15-1.3.1-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsentinel_15-1.3.1-1PGDG.rhel9.6.aarch64.rpm) |
| `pgsentinel_15` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.4 KiB | [pgsentinel_15-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsentinel_15-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsentinel_15` | `1.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.7 KiB | [pgsentinel_15-1.5.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgsentinel_15-1.5.1-1PGDG.rhel10.2.x86_64.rpm) |
| `pgsentinel_15` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.0 KiB | [pgsentinel_15-1.4.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgsentinel_15-1.4.1-1PGDG.rhel10.2.x86_64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.7 KiB | [pgsentinel_15-1.4.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgsentinel_15-1.4.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.1 KiB | [pgsentinel_15-1.4.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgsentinel_15-1.4.0-1PGDG.rhel10.0.x86_64.rpm) |
| `pgsentinel_15` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.1 KiB | [pgsentinel_15-1.3.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgsentinel_15-1.3.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_15` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.5 KiB | [pgsentinel_15-1.3.1-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgsentinel_15-1.3.1-1PGDG.rhel10.0.x86_64.rpm) |
| `pgsentinel_15` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.9 KiB | [pgsentinel_15-1.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgsentinel_15-1.2.0-1PGDG.rhel10.x86_64.rpm) |
| `pgsentinel_15` | `1.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.5 KiB | [pgsentinel_15-1.5.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsentinel_15-1.5.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pgsentinel_15` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.9 KiB | [pgsentinel_15-1.4.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsentinel_15-1.4.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.6 KiB | [pgsentinel_15-1.4.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsentinel_15-1.4.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.6 KiB | [pgsentinel_15-1.4.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsentinel_15-1.4.0-1PGDG.rhel10.0.aarch64.rpm) |
| `pgsentinel_15` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.1 KiB | [pgsentinel_15-1.3.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsentinel_15-1.3.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pgsentinel_15` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.2 KiB | [pgsentinel_15-1.3.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsentinel_15-1.3.1-1PGDG.rhel10.0.aarch64.rpm) |
| `pgsentinel_15` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.9 KiB | [pgsentinel_15-1.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsentinel_15-1.2.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgsentinel` | `1.5.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.4 KiB | [postgresql-15-pgsentinel_1.5.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.1-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.1 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg12+3_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.1 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg12+2_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.5.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.4 KiB | [postgresql-15-pgsentinel_1.5.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.1-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.2 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg12+3_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.2 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg12+2_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.5.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.4 KiB | [postgresql-15-pgsentinel_1.5.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.1-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.2 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg13+3_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.2 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg13+2_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.5.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.5 KiB | [postgresql-15-pgsentinel_1.5.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.1-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.2 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg13+3_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.2 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg13+2_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.5.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 54.8 KiB | [postgresql-15-pgsentinel_1.5.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 54.5 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg22.04+3_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 54.5 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg22.04+2_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.5.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 53.5 KiB | [postgresql-15-pgsentinel_1.5.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 53.2 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg22.04+3_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 53.2 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg22.04+2_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.5.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.4 KiB | [postgresql-15-pgsentinel_1.5.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.3 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg24.04+3_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.3 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg24.04+2_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.5.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.5 KiB | [postgresql-15-pgsentinel_1.5.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.4 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg24.04+3_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.3 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg24.04+2_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.5.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 45.7 KiB | [postgresql-15-pgsentinel_1.5.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 45.5 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg26.04+3_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 45.5 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg26.04+2_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.5.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 44.6 KiB | [postgresql-15-pgsentinel_1.5.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 44.3 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg26.04+3_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 44.3 KiB | [postgresql-15-pgsentinel_1.5.0-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.5.0-1.pgdg26.04+2_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsentinel_14` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.6 KiB | [pgsentinel_14-1.5.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsentinel_14-1.5.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.7 KiB | [pgsentinel_14-1.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsentinel_14-1.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_14` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.3 KiB | [pgsentinel_14-1.3.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsentinel_14-1.3.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_14` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.7 KiB | [pgsentinel_14-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsentinel_14-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsentinel_14` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.5 KiB | [pgsentinel_14-1.5.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsentinel_14-1.5.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.9 KiB | [pgsentinel_14-1.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsentinel_14-1.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_14` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.3 KiB | [pgsentinel_14-1.3.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsentinel_14-1.3.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_14` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.7 KiB | [pgsentinel_14-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsentinel_14-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsentinel_14` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.0 KiB | [pgsentinel_14-1.5.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsentinel_14-1.5.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pgsentinel_14` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.5 KiB | [pgsentinel_14-1.4.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsentinel_14-1.4.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.1 KiB | [pgsentinel_14-1.4.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsentinel_14-1.4.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.2 KiB | [pgsentinel_14-1.4.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsentinel_14-1.4.0-1PGDG.rhel9.6.x86_64.rpm) |
| `pgsentinel_14` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.6 KiB | [pgsentinel_14-1.3.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsentinel_14-1.3.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_14` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.8 KiB | [pgsentinel_14-1.3.1-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsentinel_14-1.3.1-1PGDG.rhel9.6.x86_64.rpm) |
| `pgsentinel_14` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.1 KiB | [pgsentinel_14-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsentinel_14-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsentinel_14` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.1 KiB | [pgsentinel_14-1.5.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsentinel_14-1.5.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pgsentinel_14` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.7 KiB | [pgsentinel_14-1.4.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsentinel_14-1.4.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.3 KiB | [pgsentinel_14-1.4.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsentinel_14-1.4.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.5 KiB | [pgsentinel_14-1.4.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsentinel_14-1.4.0-1PGDG.rhel9.6.aarch64.rpm) |
| `pgsentinel_14` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.9 KiB | [pgsentinel_14-1.3.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsentinel_14-1.3.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_14` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.0 KiB | [pgsentinel_14-1.3.1-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsentinel_14-1.3.1-1PGDG.rhel9.6.aarch64.rpm) |
| `pgsentinel_14` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.3 KiB | [pgsentinel_14-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsentinel_14-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsentinel_14` | `1.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.6 KiB | [pgsentinel_14-1.5.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgsentinel_14-1.5.1-1PGDG.rhel10.2.x86_64.rpm) |
| `pgsentinel_14` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.9 KiB | [pgsentinel_14-1.4.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgsentinel_14-1.4.1-1PGDG.rhel10.2.x86_64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.6 KiB | [pgsentinel_14-1.4.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgsentinel_14-1.4.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.0 KiB | [pgsentinel_14-1.4.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgsentinel_14-1.4.0-1PGDG.rhel10.0.x86_64.rpm) |
| `pgsentinel_14` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.0 KiB | [pgsentinel_14-1.3.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgsentinel_14-1.3.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_14` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.4 KiB | [pgsentinel_14-1.3.1-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgsentinel_14-1.3.1-1PGDG.rhel10.0.x86_64.rpm) |
| `pgsentinel_14` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.8 KiB | [pgsentinel_14-1.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgsentinel_14-1.2.0-1PGDG.rhel10.x86_64.rpm) |
| `pgsentinel_14` | `1.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.3 KiB | [pgsentinel_14-1.5.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgsentinel_14-1.5.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pgsentinel_14` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.8 KiB | [pgsentinel_14-1.4.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgsentinel_14-1.4.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.5 KiB | [pgsentinel_14-1.4.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgsentinel_14-1.4.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.5 KiB | [pgsentinel_14-1.4.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgsentinel_14-1.4.0-1PGDG.rhel10.0.aarch64.rpm) |
| `pgsentinel_14` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.1 KiB | [pgsentinel_14-1.3.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgsentinel_14-1.3.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pgsentinel_14` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.0 KiB | [pgsentinel_14-1.3.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgsentinel_14-1.3.1-1PGDG.rhel10.0.aarch64.rpm) |
| `pgsentinel_14` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.8 KiB | [pgsentinel_14-1.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgsentinel_14-1.2.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgsentinel` | `1.5.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.1 KiB | [postgresql-14-pgsentinel_1.5.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.1-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 44.8 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg12+3_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 44.8 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg12+2_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.5.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.1 KiB | [postgresql-14-pgsentinel_1.5.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.1-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.8 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg12+3_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.8 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg12+2_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.5.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.1 KiB | [postgresql-14-pgsentinel_1.5.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.1-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 44.8 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg13+3_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 44.8 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg13+2_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.5.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.2 KiB | [postgresql-14-pgsentinel_1.5.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.1-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.9 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg13+3_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.9 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg13+2_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.5.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 52.6 KiB | [postgresql-14-pgsentinel_1.5.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 52.3 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg22.04+3_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 52.4 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg22.04+2_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.5.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 51.1 KiB | [postgresql-14-pgsentinel_1.5.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 50.8 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg22.04+3_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 50.9 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg22.04+2_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.5.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.1 KiB | [postgresql-14-pgsentinel_1.5.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 44.8 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg24.04+3_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 44.8 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg24.04+2_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.5.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.2 KiB | [postgresql-14-pgsentinel_1.5.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.9 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg24.04+3_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.9 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg24.04+2_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.5.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 45.4 KiB | [postgresql-14-pgsentinel_1.5.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 45.2 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg26.04+3_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 45.1 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg26.04+2_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.5.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 44.3 KiB | [postgresql-14-pgsentinel_1.5.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 44.0 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg26.04+3_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 44.0 KiB | [postgresql-14-pgsentinel_1.5.0-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.5.0-1.pgdg26.04+2_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgsentinel/pgsentinel" title="Repository" icon="github" subtitle="github.com/pgsentinel/pgsentinel" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgsentinel-1.5.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pgsentinel;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pgsentinel;		# install via package name, for the active PG version

pig install pgsentinel -v 18;   # install for PG 18
pig install pgsentinel -v 17;   # install for PG 17
pig install pgsentinel -v 16;   # install for PG 16
pig install pgsentinel -v 15;   # install for PG 15
pig install pgsentinel -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pgsentinel';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgsentinel;
```

## Usage

Sources:

- [pgsentinel 1.5.1 README](https://github.com/pgsentinel/pgsentinel/blob/v1.5.1/README.md)
- [pgsentinel 1.5.1 release](https://github.com/pgsentinel/pgsentinel/releases/tag/v1.5.1)
- [1.5.1 upgrade SQL](https://github.com/pgsentinel/pgsentinel/blob/v1.5.1/src/pgsentinel--1.5.0--1.5.1.sql)
- [pgsentinel control file](https://github.com/pgsentinel/pgsentinel/blob/v1.5.1/src/pgsentinel.control)

`pgsentinel` records active session history by sampling `pg_stat_activity` at regular intervals and linking activity with `pg_stat_statements` query statistics. It keeps recent samples in shared-memory ring buffers managed by a background worker.

```ini
shared_preload_libraries = 'pg_stat_statements,pgsentinel'
pg_stat_statements.track = all
pgsentinel.db_name = 'postgres'
```

Restart PostgreSQL, then enable both extensions in the database used by the worker:

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION pgsentinel;
```

### Active Session History

```sql
SELECT ash_time, datname, usename, pid, state,
       wait_event_type, wait_event, query, queryid
FROM pg_active_session_history
ORDER BY ash_time DESC;
```

Key columns beyond `pg_stat_activity`:

| Column | Description |
|--------|-------------|
| `ash_time` | Sampling timestamp |
| `top_level_query` | Top-level statement (for PL/pgSQL) |
| `query` | Statement with actual parameter values |
| `cmdtype` | Statement type: SELECT, UPDATE, INSERT, DELETE, UTILITY, UNKNOWN, NOTHING |
| `queryid` | Links to `pg_stat_statements` |
| `blockers` | Number of blocking processes |
| `blockerpid` | PID of a blocking process |
| `blocker_state` | State of the blocker |

### Query Statistics History

When enabled, pgsentinel also samples `pg_stat_statements` concurrently:

```sql
SELECT ash_time, queryid, calls, total_exec_time, rows,
       shared_blks_hit, shared_blks_read
FROM pg_stat_statements_history
ORDER BY ash_time DESC;
```

### Example: Wait Analysis

```sql
-- Top wait events in the last hour
SELECT wait_event_type, wait_event, count(*)
FROM pg_active_session_history
WHERE ash_time > now() - interval '1 hour'
  AND wait_event IS NOT NULL
GROUP BY 1, 2
ORDER BY 3 DESC;

-- Blocking analysis
SELECT blockerpid, blocker_state, count(*)
FROM pg_active_session_history
WHERE blockers > 0
GROUP BY 1, 2
ORDER BY 3 DESC;
```

### Configuration

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pgsentinel_ash.sampling_period` | 1 | Sampling period in seconds |
| `pgsentinel_ash.max_entries` | 1000 | Ring buffer size for ASH |
| `pgsentinel.db_name` | `postgres` | Database for worker connection |
| `pgsentinel_ash.track_idle_trans` | `false` | Track idle-in-transaction sessions |
| `pgsentinel_pgssh.max_entries` | 1000 | Ring buffer for pg_stat_statements history |
| `pgsentinel_pgssh.enable` | `false` | Enable pg_stat_statements history |

### Version and Privilege Changes

Version 1.5.0 separates `queryid` from `nested_queryid`: on PostgreSQL 16 and later the former comes from the activity query identifier, while the latter identifies the inner parsed statement; they match on older versions. The upgrade recreates the history view, so inspect dependent views and grants first.

Version 1.5.1 revokes PUBLIC execution of `get_parsedinfo(int)` because it can expose other backends' query text and literals. Apply the SQL upgrade after installing the new files, and grant access only to roles that need it:

```sql
ALTER EXTENSION pgsentinel UPDATE TO '1.5.1';
```

History remains in finite shared-memory ring buffers; export it for long-term retention. Sampled query text can contain sensitive values, so review access to history views as well. Upstream CI covers PostgreSQL 10–19.
