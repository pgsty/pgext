---
title: "pg_stat_tcpinfo"
linkTitle: "pg_stat_tcpinfo"
description: "Linux TCP connection statistics, latency and retransmissions"
weight: 6480
categories: ["STAT"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PGDG"]
page_width: full
---

[**pg_stat_tcpinfo**](https://github.com/jakubwartakEDB/pg_stat_tcpinfo) : Linux TCP connection statistics, latency and retransmissions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6480** | {{< badge content="pg_stat_tcpinfo" link="https://github.com/jakubwartakEDB/pg_stat_tcpinfo" >}} | {{< ext "pg_stat_tcpinfo" >}} | `1.1` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgnodemx" >}} {{< ext "system_stats" >}} {{< ext "pg_stat_kcache" >}} {{< ext "pg_statvfs" >}} |

> [!Note] Linux >=4.18; procfs and NETLINK_INET_DIAG required. Monitoring roles can see all server-OS-user sockets. No preload or upstream 1.0-to-1.1 upgrade script. PG14-18; PGDG remains principal supplier; PGSTY 1.1 batch supplements published.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_stat_tcpinfo` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1` | {{< bg "18" "pg_stat_tcpinfo_18" "green" >}} {{< bg "17" "pg_stat_tcpinfo_17" "green" >}} {{< bg "16" "pg_stat_tcpinfo_16" "green" >}} {{< bg "15" "pg_stat_tcpinfo_15" "green" >}} {{< bg "14" "pg_stat_tcpinfo_14" "green" >}} | `pg_stat_tcpinfo_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1` | {{< bg "18" "postgresql-18-pg-stat-tcpinfo" "green" >}} {{< bg "17" "postgresql-17-pg-stat-tcpinfo" "green" >}} {{< bg "16" "postgresql-16-pg-stat-tcpinfo" "green" >}} {{< bg "15" "postgresql-15-pg-stat-tcpinfo" "green" >}} {{< bg "14" "postgresql-14-pg-stat-tcpinfo" "green" >}} | `postgresql-$v-pg-stat-tcpinfo` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_14 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_14 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_14 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_14 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1" "pg_stat_tcpinfo_14 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.1" "postgresql-18-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-17-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-16-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-15-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-14-pg-stat-tcpinfo : AVAIL 2" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.1" "postgresql-18-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-17-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-16-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-15-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-14-pg-stat-tcpinfo : AVAIL 2" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.1" "postgresql-18-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-17-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-16-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-15-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-14-pg-stat-tcpinfo : AVAIL 2" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.1" "postgresql-18-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-17-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-16-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-15-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-14-pg-stat-tcpinfo : AVAIL 2" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.1" "postgresql-18-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-17-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-16-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-15-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-14-pg-stat-tcpinfo : AVAIL 2" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.1" "postgresql-18-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-17-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-16-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-15-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-14-pg-stat-tcpinfo : AVAIL 2" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.1" "postgresql-18-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-17-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-16-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-15-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-14-pg-stat-tcpinfo : AVAIL 2" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.1" "postgresql-18-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-17-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-16-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-15-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-14-pg-stat-tcpinfo : AVAIL 2" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.1" "postgresql-18-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-17-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-16-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-15-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-14-pg-stat-tcpinfo : AVAIL 2" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.1" "postgresql-18-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-17-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-16-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-15-pg-stat-tcpinfo : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-14-pg-stat-tcpinfo : AVAIL 2" "blue" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_tcpinfo_18` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 41.0 KiB | [pg_stat_tcpinfo_18-1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_stat_tcpinfo_18-1.1-1PGSTY.el8.x86_64.rpm) |
| `pg_stat_tcpinfo_18` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.7 KiB | [pg_stat_tcpinfo_18-1.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_stat_tcpinfo_18-1.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_tcpinfo_18` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 40.4 KiB | [pg_stat_tcpinfo_18-1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_stat_tcpinfo_18-1.1-1PGSTY.el8.aarch64.rpm) |
| `pg_stat_tcpinfo_18` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.7 KiB | [pg_stat_tcpinfo_18-1.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_stat_tcpinfo_18-1.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_tcpinfo_18` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.1 KiB | [pg_stat_tcpinfo_18-1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_tcpinfo_18-1.1-1PGSTY.el9.x86_64.rpm) |
| `pg_stat_tcpinfo_18` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.2 KiB | [pg_stat_tcpinfo_18-1.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_stat_tcpinfo_18-1.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_stat_tcpinfo_18` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.5 KiB | [pg_stat_tcpinfo_18-1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_tcpinfo_18-1.1-1PGSTY.el9.aarch64.rpm) |
| `pg_stat_tcpinfo_18` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.3 KiB | [pg_stat_tcpinfo_18-1.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_stat_tcpinfo_18-1.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_stat_tcpinfo_18` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.1 KiB | [pg_stat_tcpinfo_18-1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_tcpinfo_18-1.1-1PGSTY.el10.x86_64.rpm) |
| `pg_stat_tcpinfo_18` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 41.8 KiB | [pg_stat_tcpinfo_18-1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_tcpinfo_18-1.1-1PGSTY.el10.aarch64.rpm) |
| `pg_stat_tcpinfo_18` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.6 KiB | [pg_stat_tcpinfo_18-1.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_tcpinfo_18-1.1-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 32.7 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 32.6 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 31.9 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 31.8 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 32.8 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 32.7 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 32.2 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 32.1 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 33.4 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 35.3 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 32.5 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 34.8 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 32.8 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 34.2 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 32.1 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 33.6 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 32.7 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 34.0 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 32.3 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-pg-stat-tcpinfo` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 33.7 KiB | [postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-tcpinfo/postgresql-18-pg-stat-tcpinfo_1.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_tcpinfo_17` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 40.9 KiB | [pg_stat_tcpinfo_17-1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_stat_tcpinfo_17-1.1-1PGSTY.el8.x86_64.rpm) |
| `pg_stat_tcpinfo_17` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.7 KiB | [pg_stat_tcpinfo_17-1.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_tcpinfo_17-1.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_tcpinfo_17` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 40.2 KiB | [pg_stat_tcpinfo_17-1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_stat_tcpinfo_17-1.1-1PGSTY.el8.aarch64.rpm) |
| `pg_stat_tcpinfo_17` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.7 KiB | [pg_stat_tcpinfo_17-1.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_tcpinfo_17-1.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_tcpinfo_17` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 41.8 KiB | [pg_stat_tcpinfo_17-1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_tcpinfo_17-1.1-1PGSTY.el9.x86_64.rpm) |
| `pg_stat_tcpinfo_17` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.1 KiB | [pg_stat_tcpinfo_17-1.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_stat_tcpinfo_17-1.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_stat_tcpinfo_17` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.3 KiB | [pg_stat_tcpinfo_17-1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_tcpinfo_17-1.1-1PGSTY.el9.aarch64.rpm) |
| `pg_stat_tcpinfo_17` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.3 KiB | [pg_stat_tcpinfo_17-1.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_stat_tcpinfo_17-1.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_stat_tcpinfo_17` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.0 KiB | [pg_stat_tcpinfo_17-1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_tcpinfo_17-1.1-1PGSTY.el10.x86_64.rpm) |
| `pg_stat_tcpinfo_17` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 41.6 KiB | [pg_stat_tcpinfo_17-1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_tcpinfo_17-1.1-1PGSTY.el10.aarch64.rpm) |
| `pg_stat_tcpinfo_17` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.6 KiB | [pg_stat_tcpinfo_17-1.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_tcpinfo_17-1.1-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 32.5 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 32.4 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 31.9 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 31.7 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 32.6 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 32.4 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 32.2 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 32.0 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 36.4 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 38.5 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 35.7 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 37.9 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 32.6 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 34.0 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 31.9 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 33.5 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 32.6 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 33.8 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 32.1 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-pg-stat-tcpinfo` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 33.7 KiB | [postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-tcpinfo/postgresql-17-pg-stat-tcpinfo_1.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_tcpinfo_16` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 40.9 KiB | [pg_stat_tcpinfo_16-1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_stat_tcpinfo_16-1.1-1PGSTY.el8.x86_64.rpm) |
| `pg_stat_tcpinfo_16` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.7 KiB | [pg_stat_tcpinfo_16-1.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_tcpinfo_16-1.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_tcpinfo_16` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 40.2 KiB | [pg_stat_tcpinfo_16-1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_stat_tcpinfo_16-1.1-1PGSTY.el8.aarch64.rpm) |
| `pg_stat_tcpinfo_16` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.7 KiB | [pg_stat_tcpinfo_16-1.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_tcpinfo_16-1.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_tcpinfo_16` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 41.8 KiB | [pg_stat_tcpinfo_16-1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_tcpinfo_16-1.1-1PGSTY.el9.x86_64.rpm) |
| `pg_stat_tcpinfo_16` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.1 KiB | [pg_stat_tcpinfo_16-1.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_tcpinfo_16-1.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_stat_tcpinfo_16` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.3 KiB | [pg_stat_tcpinfo_16-1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_tcpinfo_16-1.1-1PGSTY.el9.aarch64.rpm) |
| `pg_stat_tcpinfo_16` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.3 KiB | [pg_stat_tcpinfo_16-1.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_tcpinfo_16-1.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_stat_tcpinfo_16` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.0 KiB | [pg_stat_tcpinfo_16-1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_tcpinfo_16-1.1-1PGSTY.el10.x86_64.rpm) |
| `pg_stat_tcpinfo_16` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 41.6 KiB | [pg_stat_tcpinfo_16-1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_tcpinfo_16-1.1-1PGSTY.el10.aarch64.rpm) |
| `pg_stat_tcpinfo_16` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.6 KiB | [pg_stat_tcpinfo_16-1.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_tcpinfo_16-1.1-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 32.5 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 32.4 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 31.9 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 31.7 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 32.7 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 32.4 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 32.2 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 32.0 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 36.4 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 38.4 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 35.6 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 37.9 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 32.7 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 34.0 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 32.0 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 34.0 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 32.6 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 33.8 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 32.1 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-pg-stat-tcpinfo` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 33.7 KiB | [postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-tcpinfo/postgresql-16-pg-stat-tcpinfo_1.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_tcpinfo_15` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 40.9 KiB | [pg_stat_tcpinfo_15-1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_stat_tcpinfo_15-1.1-1PGSTY.el8.x86_64.rpm) |
| `pg_stat_tcpinfo_15` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.7 KiB | [pg_stat_tcpinfo_15-1.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_tcpinfo_15-1.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_tcpinfo_15` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 40.2 KiB | [pg_stat_tcpinfo_15-1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_stat_tcpinfo_15-1.1-1PGSTY.el8.aarch64.rpm) |
| `pg_stat_tcpinfo_15` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.7 KiB | [pg_stat_tcpinfo_15-1.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_tcpinfo_15-1.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_tcpinfo_15` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 41.8 KiB | [pg_stat_tcpinfo_15-1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_tcpinfo_15-1.1-1PGSTY.el9.x86_64.rpm) |
| `pg_stat_tcpinfo_15` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.1 KiB | [pg_stat_tcpinfo_15-1.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_tcpinfo_15-1.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_stat_tcpinfo_15` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.3 KiB | [pg_stat_tcpinfo_15-1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_tcpinfo_15-1.1-1PGSTY.el9.aarch64.rpm) |
| `pg_stat_tcpinfo_15` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.3 KiB | [pg_stat_tcpinfo_15-1.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_tcpinfo_15-1.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_stat_tcpinfo_15` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.0 KiB | [pg_stat_tcpinfo_15-1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_tcpinfo_15-1.1-1PGSTY.el10.x86_64.rpm) |
| `pg_stat_tcpinfo_15` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 41.6 KiB | [pg_stat_tcpinfo_15-1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_tcpinfo_15-1.1-1PGSTY.el10.aarch64.rpm) |
| `pg_stat_tcpinfo_15` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.6 KiB | [pg_stat_tcpinfo_15-1.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_tcpinfo_15-1.1-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 32.6 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 32.4 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 31.9 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 35.6 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 32.6 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 32.4 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 32.2 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 32.0 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 36.3 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 38.4 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 35.6 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 37.9 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 32.7 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 34.0 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 31.9 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 33.5 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 32.6 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 33.8 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 32.1 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-pg-stat-tcpinfo` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 33.7 KiB | [postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-tcpinfo/postgresql-15-pg-stat-tcpinfo_1.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_tcpinfo_14` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 40.8 KiB | [pg_stat_tcpinfo_14-1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_stat_tcpinfo_14-1.1-1PGSTY.el8.x86_64.rpm) |
| `pg_stat_tcpinfo_14` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.7 KiB | [pg_stat_tcpinfo_14-1.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_tcpinfo_14-1.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_tcpinfo_14` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 40.1 KiB | [pg_stat_tcpinfo_14-1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_stat_tcpinfo_14-1.1-1PGSTY.el8.aarch64.rpm) |
| `pg_stat_tcpinfo_14` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.7 KiB | [pg_stat_tcpinfo_14-1.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_tcpinfo_14-1.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_tcpinfo_14` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 41.9 KiB | [pg_stat_tcpinfo_14-1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stat_tcpinfo_14-1.1-1PGSTY.el9.x86_64.rpm) |
| `pg_stat_tcpinfo_14` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.1 KiB | [pg_stat_tcpinfo_14-1.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_tcpinfo_14-1.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_stat_tcpinfo_14` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.3 KiB | [pg_stat_tcpinfo_14-1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stat_tcpinfo_14-1.1-1PGSTY.el9.aarch64.rpm) |
| `pg_stat_tcpinfo_14` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.3 KiB | [pg_stat_tcpinfo_14-1.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_tcpinfo_14-1.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_stat_tcpinfo_14` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 41.9 KiB | [pg_stat_tcpinfo_14-1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stat_tcpinfo_14-1.1-1PGSTY.el10.x86_64.rpm) |
| `pg_stat_tcpinfo_14` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 41.5 KiB | [pg_stat_tcpinfo_14-1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stat_tcpinfo_14-1.1-1PGSTY.el10.aarch64.rpm) |
| `pg_stat_tcpinfo_14` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.6 KiB | [pg_stat_tcpinfo_14-1.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_stat_tcpinfo_14-1.1-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 32.6 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 32.4 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 31.8 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 31.7 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 32.7 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 32.5 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 32.1 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 31.9 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 36.3 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 38.4 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 35.5 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 38.0 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 32.6 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 34.0 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 31.9 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 33.4 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 32.5 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 33.8 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 32.1 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-pg-stat-tcpinfo` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 33.7 KiB | [postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stat-tcpinfo/postgresql-14-pg-stat-tcpinfo_1.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/jakubwartakEDB/pg_stat_tcpinfo" title="Repository" icon="github" subtitle="github.com/jakubwartakEDB/pg_stat_tcpinfo" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_stat_tcpinfo-1.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_stat_tcpinfo;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_stat_tcpinfo;		# install via package name, for the active PG version

pig install pg_stat_tcpinfo -v 18;   # install for PG 18
pig install pg_stat_tcpinfo -v 17;   # install for PG 17
pig install pg_stat_tcpinfo -v 16;   # install for PG 16
pig install pg_stat_tcpinfo -v 15;   # install for PG 15
pig install pg_stat_tcpinfo -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_stat_tcpinfo;
```

## Usage

Sources:

- [README](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/v1.1/README.md)
- [Installation SQL](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/v1.1/pg_stat_tcpinfo--1.1.sql)
- [Extension control](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/v1.1/pg_stat_tcpinfo.control)
- [Release 1.1](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/releases/tag/v1.1)
- [License](https://github.com/jakubwartakEDB/pg_stat_tcpinfo/blob/v1.1/LICENSE)

`pg_stat_tcpinfo` 1.1 exposes Linux TCP socket statistics through SQL. It helps investigate slow client connections, replication links and foreign-server traffic using round-trip time, retransmissions, congestion control and queue sizes, without shell access to the host.

It requires Linux 4.18 or later, access to `/proc`, and the `NETLINK_INET_DIAG` interface. Install as a superuser; no preload or server restart is required.

### Core Workflow

```sql
CREATE EXTENSION pg_stat_tcpinfo;

SELECT a.pid, a.usename, a.application_name,
       t.dst_addr, t.dst_port, t.state, t.recvq, t.sendq,
       t.tcpinfo->>'rtt' AS rtt_ms,
       t.tcpinfo->>'total_retrans' AS retransmissions,
       t.tcpinfo->>'congestion' AS congestion_control
FROM pg_stat_activity a
JOIN pg_stat_tcpinfo t ON t.pid = a.pid
WHERE t.state = 'ESTABLISHED';
```

### Objects and Metrics

The `pg_stat_tcpinfo` view calls `pg_stat_get_tcpinfo()`. Each row contains a process and OS user ID, local and remote addresses and ports, connection state, receive and send queues, and a `tcpinfo` JSONB object. IPv4 and IPv6 sockets are supported. The `rtt` field is expressed in milliseconds; `total_retrans` is the cumulative retransmission counter.

The view includes every TCP socket owned by the server's OS user, including sockets from other database clusters or programs running under that account. Joining on the backend PID narrows the result to connections visible in the current cluster. Unix-domain connections are not TCP sockets and do not appear; TIME-WAIT sockets have no owner and are omitted. Metrics unavailable on an older kernel may be null.

### Permissions

The installation revokes public execution of the function and grants function and view access to `pg_read_all_stats`; `pg_monitor` inherits that access. Ordinary application roles cannot read the metrics by default. Additional grants expose all sockets owned by the server OS user, not only the caller's database connections.

### Operational Boundaries

Each call scans process descriptors and queries the kernel; it is an on-demand observation, not a historical sampler. Container PID and network namespaces and procfs restrictions determine which sockets and PIDs are visible. Choose a sampling interval appropriate to the number of processes and connections.

Use a package built for the server's PostgreSQL major version. The 1.1 release includes fresh-install SQL for versions 1.0 and 1.1 but no upgrade script between them, so an existing 1.0 installation has no upstream-provided `ALTER EXTENSION pg_stat_tcpinfo UPDATE` path.
