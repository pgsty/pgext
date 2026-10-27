---
title: "pg_stat_log"
linkTitle: "pg_stat_log"
description: "Track cumulative PostgreSQL log message statistics by backend, database, user, level, and SQLSTATE."
weight: 6040
categories: ["STAT"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PGDG"]
page_width: full
---

[**pg_stat_log**](https://github.com/fabriziomello/pg_stat_log) : Track cumulative PostgreSQL log message statistics by backend, database, user, level, and SQLSTATE.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6040** | {{< badge content="pg_stat_log" link="https://github.com/fabriziomello/pg_stat_log" >}} | {{< ext "pg_stat_log" >}} | `0.3` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgelog" >}} {{< ext "logerrors" >}} {{< ext "pg_sqlog" >}} {{< ext "log_fdw" >}} {{< ext "pgaudit" >}} {{< ext "pgdisablelogerror" >}} {{< ext "pg_auth_mon" >}} {{< ext "pg_tracing" >}} {{< ext "pg_stat_backtrace" >}} {{< ext "pgfr_record" >}} |

> [!Note] PostgreSQL 18+; shared preload and restart required. Upgrading 0.1/0.2 to 0.3 resets all cumulative statistics; SQL upgrade recreates the view/function, requiring dependent views and custom grants to be restored.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_stat_log` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.3` | {{< bg "18" "pg_stat_log_18" "green" >}} {{< bg "17" "pg_stat_log_17" "red" >}} {{< bg "16" "pg_stat_log_16" "red" >}} {{< bg "15" "pg_stat_log_15" "red" >}} {{< bg "14" "pg_stat_log_14" "red" >}} | `pg_stat_log_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.3` | {{< bg "18" "postgresql-18-stat-log" "green" >}} {{< bg "17" "postgresql-17-stat-log" "red" >}} {{< bg "16" "postgresql-16-stat-log" "red" >}} {{< bg "15" "postgresql-15-stat-log" "red" >}} {{< bg "14" "postgresql-14-stat-log" "red" >}} | `postgresql-$v-stat-log` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 0.3" "pg_stat_log_18 : AVAIL 3" "blue" >}} | {{< bg "N/A" "pg_stat_log_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_stat_log_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_stat_log_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_stat_log_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 0.3" "pg_stat_log_18 : AVAIL 3" "blue" >}} | {{< bg "N/A" "pg_stat_log_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_stat_log_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_stat_log_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_stat_log_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 0.3" "pg_stat_log_18 : AVAIL 5" "blue" >}} | {{< bg "N/A" "pg_stat_log_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_stat_log_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_stat_log_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_stat_log_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 0.3" "pg_stat_log_18 : AVAIL 5" "blue" >}} | {{< bg "N/A" "pg_stat_log_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_stat_log_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_stat_log_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_stat_log_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.3" "pg_stat_log_18 : AVAIL 5" "blue" >}} | {{< bg "N/A" "pg_stat_log_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_stat_log_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_stat_log_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_stat_log_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.3" "pg_stat_log_18 : AVAIL 5" "blue" >}} | {{< bg "N/A" "pg_stat_log_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_stat_log_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_stat_log_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_stat_log_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 0.3" "postgresql-18-stat-log : AVAIL 3" "blue" >}} | {{< bg "N/A" "postgresql-17-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stat-log : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 0.3" "postgresql-18-stat-log : AVAIL 3" "blue" >}} | {{< bg "N/A" "postgresql-17-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stat-log : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 0.3" "postgresql-18-stat-log : AVAIL 3" "blue" >}} | {{< bg "N/A" "postgresql-17-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stat-log : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 0.3" "postgresql-18-stat-log : AVAIL 3" "blue" >}} | {{< bg "N/A" "postgresql-17-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stat-log : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 0.3" "postgresql-18-stat-log : AVAIL 3" "blue" >}} | {{< bg "N/A" "postgresql-17-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stat-log : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 0.3" "postgresql-18-stat-log : AVAIL 3" "blue" >}} | {{< bg "N/A" "postgresql-17-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stat-log : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 0.3" "postgresql-18-stat-log : AVAIL 3" "blue" >}} | {{< bg "N/A" "postgresql-17-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stat-log : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 0.3" "postgresql-18-stat-log : AVAIL 3" "blue" >}} | {{< bg "N/A" "postgresql-17-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stat-log : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 0.3" "postgresql-18-stat-log : AVAIL 3" "blue" >}} | {{< bg "N/A" "postgresql-17-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stat-log : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 0.3" "postgresql-18-stat-log : AVAIL 3" "blue" >}} | {{< bg "N/A" "postgresql-17-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stat-log : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stat-log : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_log_18` | `0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.0 KiB | [pg_stat_log_18-0.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_stat_log_18-0.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_log_18` | `0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.0 KiB | [pg_stat_log_18-0.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_stat_log_18-0.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_log_18` | `0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.2 KiB | [pg_stat_log_18-0.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_stat_log_18-0.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_log_18` | `0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.7 KiB | [pg_stat_log_18-0.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_stat_log_18-0.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_log_18` | `0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.9 KiB | [pg_stat_log_18-0.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_stat_log_18-0.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_log_18` | `0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.1 KiB | [pg_stat_log_18-0.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_stat_log_18-0.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_log_18` | `0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.6 KiB | [pg_stat_log_18-0.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_stat_log_18-0.3-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_stat_log_18` | `0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.3 KiB | [pg_stat_log_18-0.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_stat_log_18-0.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_stat_log_18` | `0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.4 KiB | [pg_stat_log_18-0.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_stat_log_18-0.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_stat_log_18` | `0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.5 KiB | [pg_stat_log_18-0.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_stat_log_18-0.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_stat_log_18` | `0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.6 KiB | [pg_stat_log_18-0.1-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_stat_log_18-0.1-1PGDG.rhel9.6.x86_64.rpm) |
| `pg_stat_log_18` | `0.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.4 KiB | [pg_stat_log_18-0.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_stat_log_18-0.3-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_stat_log_18` | `0.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.1 KiB | [pg_stat_log_18-0.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_stat_log_18-0.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_stat_log_18` | `0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.4 KiB | [pg_stat_log_18-0.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_stat_log_18-0.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_stat_log_18` | `0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.4 KiB | [pg_stat_log_18-0.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_stat_log_18-0.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_stat_log_18` | `0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.5 KiB | [pg_stat_log_18-0.1-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_stat_log_18-0.1-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_stat_log_18` | `0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 29.6 KiB | [pg_stat_log_18-0.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_stat_log_18-0.3-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_stat_log_18` | `0.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.4 KiB | [pg_stat_log_18-0.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_stat_log_18-0.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_stat_log_18` | `0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.5 KiB | [pg_stat_log_18-0.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_stat_log_18-0.1-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_stat_log_18` | `0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.5 KiB | [pg_stat_log_18-0.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_stat_log_18-0.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_stat_log_18` | `0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.9 KiB | [pg_stat_log_18-0.1-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_stat_log_18-0.1-1PGDG.rhel10.0.x86_64.rpm) |
| `pg_stat_log_18` | `0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.7 KiB | [pg_stat_log_18-0.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_log_18-0.3-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_log_18` | `0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.4 KiB | [pg_stat_log_18-0.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_log_18-0.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_log_18` | `0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.7 KiB | [pg_stat_log_18-0.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_log_18-0.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_log_18` | `0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.7 KiB | [pg_stat_log_18-0.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_log_18-0.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_stat_log_18` | `0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.7 KiB | [pg_stat_log_18-0.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_log_18-0.1-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-18-stat-log` | `0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 50.0 KiB | [postgresql-18-stat-log_0.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.3-1.pgdg12+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 42.7 KiB | [postgresql-18-stat-log_0.2-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg12+2_amd64.deb) |
| `postgresql-18-stat-log` | `0.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 42.7 KiB | [postgresql-18-stat-log_0.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg12+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 50.3 KiB | [postgresql-18-stat-log_0.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.3-1.pgdg12+1_arm64.deb) |
| `postgresql-18-stat-log` | `0.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 42.7 KiB | [postgresql-18-stat-log_0.2-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg12+2_arm64.deb) |
| `postgresql-18-stat-log` | `0.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 42.7 KiB | [postgresql-18-stat-log_0.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg12+1_arm64.deb) |
| `postgresql-18-stat-log` | `0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 50.1 KiB | [postgresql-18-stat-log_0.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.3-1.pgdg13+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 42.8 KiB | [postgresql-18-stat-log_0.2-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg13+2_amd64.deb) |
| `postgresql-18-stat-log` | `0.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 42.8 KiB | [postgresql-18-stat-log_0.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg13+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 50.4 KiB | [postgresql-18-stat-log_0.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.3-1.pgdg13+1_arm64.deb) |
| `postgresql-18-stat-log` | `0.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 42.8 KiB | [postgresql-18-stat-log_0.2-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg13+2_arm64.deb) |
| `postgresql-18-stat-log` | `0.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 42.9 KiB | [postgresql-18-stat-log_0.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg13+1_arm64.deb) |
| `postgresql-18-stat-log` | `0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 50.5 KiB | [postgresql-18-stat-log_0.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 43.1 KiB | [postgresql-18-stat-log_0.2-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg22.04+2_amd64.deb) |
| `postgresql-18-stat-log` | `0.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 43.1 KiB | [postgresql-18-stat-log_0.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 50.7 KiB | [postgresql-18-stat-log_0.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-stat-log` | `0.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.8 KiB | [postgresql-18-stat-log_0.2-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg22.04+2_arm64.deb) |
| `postgresql-18-stat-log` | `0.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.8 KiB | [postgresql-18-stat-log_0.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-stat-log` | `0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 50.2 KiB | [postgresql-18-stat-log_0.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 42.9 KiB | [postgresql-18-stat-log_0.2-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg24.04+2_amd64.deb) |
| `postgresql-18-stat-log` | `0.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 42.9 KiB | [postgresql-18-stat-log_0.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 50.3 KiB | [postgresql-18-stat-log_0.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-stat-log` | `0.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 42.8 KiB | [postgresql-18-stat-log_0.2-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg24.04+2_arm64.deb) |
| `postgresql-18-stat-log` | `0.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 42.8 KiB | [postgresql-18-stat-log_0.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-stat-log` | `0.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 49.8 KiB | [postgresql-18-stat-log_0.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 42.8 KiB | [postgresql-18-stat-log_0.2-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg26.04+2_amd64.deb) |
| `postgresql-18-stat-log` | `0.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 42.8 KiB | [postgresql-18-stat-log_0.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 49.8 KiB | [postgresql-18-stat-log_0.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-stat-log` | `0.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 42.8 KiB | [postgresql-18-stat-log_0.2-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg26.04+2_arm64.deb) |
| `postgresql-18-stat-log` | `0.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 42.8 KiB | [postgresql-18-stat-log_0.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.2-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/fabriziomello/pg_stat_log" title="Repository" icon="github" subtitle="github.com/fabriziomello/pg_stat_log" />}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_stat_log;		# install via package name, for the active PG version

pig install pg_stat_log -v 18;   # install for PG 18

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = '$libdir/pg_stat_log';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_stat_log;
```

## Usage

Sources:

- [README.md](https://github.com/fabriziomello/pg_stat_log/blob/20b57e2e2eff884bfcf3ede3c095762e614c51d1/README.md)
- [pg_stat_log.control](https://github.com/fabriziomello/pg_stat_log/blob/20b57e2e2eff884bfcf3ede3c095762e614c51d1/pg_stat_log.control)
- [pg_stat_log--0.1.sql](https://github.com/fabriziomello/pg_stat_log/blob/20b57e2e2eff884bfcf3ede3c095762e614c51d1/pg_stat_log--0.1.sql)
- [pg_stat_log--0.1--0.3.sql](https://github.com/fabriziomello/pg_stat_log/blob/20b57e2e2eff884bfcf3ede3c095762e614c51d1/pg_stat_log--0.1--0.3.sql)

`pg_stat_log` 0.3 collects PostgreSQL 18+ log-message counts by backend, database, role, severity and SQLSTATE, with optional grouping by query identifier. Add the library to the existing preload list and restart before creating the SQL extension. The default SQL grants allow `pg_read_all_stats` to inspect the counters.

### Core Workflow

```ini
shared_preload_libraries = 'pg_stat_log'
```

```sql
CREATE EXTENSION pg_stat_log;
SELECT database_name, user_name, query_id, elevel, sqlerrcode_name, count
FROM pg_stat_log
ORDER BY count DESC;
SELECT * FROM pg_stat_log_info();
```

### Operational Boundaries

The `pg_stat_log` view resolves names from `pg_stat_log_data()`. `pg_stat_log_info()` reports `max_entries`, `num_entries`, `n_dropped` and `stats_reset`; `pg_stat_log_reset()` clears counters and frees capacity, and is restricted by default. `pg_stat_log.enabled` and `pg_stat_log.min_error_level` control collection, with defaults true and warning. Only messages reaching the server log are counted: `log_min_messages` can filter messages before the hook runs, and caught exceptions are absent. Parallel-worker messages can count the same logical failure more than once.

`pg_stat_log.compute_query_id` defaults to off and requires restart to change; PostgreSQL's `compute_query_id` must not be off. A missing query identifier appears as NULL. `pg_stat_log.max_entries` defaults to 1024, requires restart, and drops new combinations when full; inspect `n_dropped`. Changing this capacity resets all PostgreSQL cumulative statistics on the next start. Counters survive clean restarts, but drop telemetry does not, and crash recovery discards cumulative statistics.

Version 0.3 maps connection authentication, authorization and disconnection to `PSL01`, `PSL02` and `PSL03`. Collection needs the corresponding server logging plus `pg_stat_log.min_error_level = 'log'` or lower; the role attribution is the login role.

Upstream releases 0.1 and 0.2 both install SQL version 0.1. Moving to 0.3 requires matching library files and a restart; the first start resets all PostgreSQL cumulative statistics. Then run `ALTER EXTENSION pg_stat_log UPDATE TO '0.3'` in each installed database. The update recreates the reporting view and function: drop and recreate dependent views and reapply custom grants; the default grants to `pg_read_all_stats` are restored.
