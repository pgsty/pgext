---
title: "pg_stat_kcache"
linkTitle: "pg_stat_kcache"
description: "Kernel statistics gathering"
weight: 6220
categories: ["STAT"]
width: full
---

[**pg_stat_kcache**](https://github.com/powa-team/pg_stat_kcache) : Kernel statistics gathering


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6220** | {{< badge content="pg_stat_kcache" link="https://github.com/powa-team/pg_stat_kcache" >}} | {{< ext "pg_stat_kcache" >}} | `2.3.0` | {{< category "STAT" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pg_stat_statements" >}} |
|   **See Also**    | {{< ext "pg_profile" >}} {{< ext "powa" >}} {{< ext "plprofiler" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_track_settings" >}} {{< ext "pg_wait_sampling" >}} {{< ext "system_stats" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_stat_kcache` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.3.0` | {{< bg "18" "pg_stat_kcache_18*" "green" >}} {{< bg "17" "pg_stat_kcache_17*" "green" >}} {{< bg "16" "pg_stat_kcache_16*" "green" >}} {{< bg "15" "pg_stat_kcache_15*" "green" >}} {{< bg "14" "pg_stat_kcache_14*" "green" >}} {{< bg "13" "pg_stat_kcache_13*" "green" >}} | `pg_stat_kcache_$v*` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.3.0` | {{< bg "18" "postgresql-18-pg-stat-kcache" "green" >}} {{< bg "17" "postgresql-17-pg-stat-kcache" "green" >}} {{< bg "16" "postgresql-16-pg-stat-kcache" "green" >}} {{< bg "15" "postgresql-15-pg-stat-kcache" "green" >}} {{< bg "14" "postgresql-14-pg-stat-kcache" "green" >}} {{< bg "13" "postgresql-13-pg-stat-kcache" "green" >}} | `postgresql-$v-pg-stat-kcache` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_14 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_13 : AVAIL 6" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_14 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_13 : AVAIL 5" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_14 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_13 : AVAIL 5" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_14 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_13 : AVAIL 5" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.1" "pg_stat_kcache_13 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.3.1" "postgresql-18-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-17-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-16-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-15-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-14-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-13-pg-stat-kcache : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.3.1" "postgresql-18-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-17-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-16-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-15-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-14-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-13-pg-stat-kcache : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.3.1" "postgresql-18-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-17-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-16-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-15-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-14-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-13-pg-stat-kcache : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.3.1" "postgresql-18-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-17-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-16-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-15-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-14-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-13-pg-stat-kcache : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.3.1" "postgresql-18-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-17-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-16-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-15-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-14-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-13-pg-stat-kcache : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.3.1" "postgresql-18-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-17-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-16-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-15-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-14-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-13-pg-stat-kcache : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.3.1" "postgresql-18-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-17-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-16-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-15-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-14-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-13-pg-stat-kcache : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.3.1" "postgresql-18-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-17-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-16-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-15-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-14-pg-stat-kcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.1" "postgresql-13-pg-stat-kcache : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_kcache_18` | `2.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.8 KiB | [pg_stat_kcache_18-2.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_stat_kcache_18-2.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_18` | `2.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.7 KiB | [pg_stat_kcache_18-2.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_stat_kcache_18-2.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_18` | `2.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.5 KiB | [pg_stat_kcache_18-2.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_stat_kcache_18-2.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_18` | `2.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.6 KiB | [pg_stat_kcache_18-2.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_stat_kcache_18-2.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_18` | `2.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.9 KiB | [pg_stat_kcache_18-2.3.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_stat_kcache_18-2.3.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_stat_kcache_18` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.5 KiB | [pg_stat_kcache_18-2.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_kcache_18-2.3.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-stat-kcache` | `2.3.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 36.9 KiB | [postgresql-18-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.7 KiB | [postgresql-18-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 36.9 KiB | [postgresql-18-pg-stat-kcache_2.3.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.8 KiB | [postgresql-18-pg-stat-kcache_2.3.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 35.8 KiB | [postgresql-18-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 35.2 KiB | [postgresql-18-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.2 KiB | [postgresql-18-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.8 KiB | [postgresql-18-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_kcache_17` | `2.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.8 KiB | [pg_stat_kcache_17-2.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_kcache_17-2.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_17` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.3 KiB | [pg_stat_kcache_17-2.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_kcache_17-2.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_17` | `2.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.7 KiB | [pg_stat_kcache_17-2.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_kcache_17-2.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_17` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.2 KiB | [pg_stat_kcache_17-2.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_kcache_17-2.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_17` | `2.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.5 KiB | [pg_stat_kcache_17-2.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_stat_kcache_17-2.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_17` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.4 KiB | [pg_stat_kcache_17-2.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_stat_kcache_17-2.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_17` | `2.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.5 KiB | [pg_stat_kcache_17-2.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_stat_kcache_17-2.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_17` | `2.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.5 KiB | [pg_stat_kcache_17-2.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_stat_kcache_17-2.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_17` | `2.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.0 KiB | [pg_stat_kcache_17-2.3.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_stat_kcache_17-2.3.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_stat_kcache_17` | `2.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.7 KiB | [pg_stat_kcache_17-2.3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_stat_kcache_17-2.3.0-2PGDG.rhel10.x86_64.rpm) |
| `pg_stat_kcache_17` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.5 KiB | [pg_stat_kcache_17-2.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_kcache_17-2.3.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_stat_kcache_17` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.2 KiB | [pg_stat_kcache_17-2.3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_kcache_17-2.3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-stat-kcache` | `2.3.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.0 KiB | [postgresql-17-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.6 KiB | [postgresql-17-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 36.9 KiB | [postgresql-17-pg-stat-kcache_2.3.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.7 KiB | [postgresql-17-pg-stat-kcache_2.3.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.1 KiB | [postgresql-17-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.6 KiB | [postgresql-17-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.2 KiB | [postgresql-17-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.8 KiB | [postgresql-17-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_kcache_16` | `2.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.8 KiB | [pg_stat_kcache_16-2.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_kcache_16-2.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_16` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.3 KiB | [pg_stat_kcache_16-2.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_kcache_16-2.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_16` | `2.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.6 KiB | [pg_stat_kcache_16-2.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_kcache_16-2.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_16` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.3 KiB | [pg_stat_kcache_16-2.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_kcache_16-2.2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_16` | `2.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.7 KiB | [pg_stat_kcache_16-2.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_kcache_16-2.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.2 KiB | [pg_stat_kcache_16-2.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_kcache_16-2.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.6 KiB | [pg_stat_kcache_16-2.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_kcache_16-2.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.3 KiB | [pg_stat_kcache_16-2.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_kcache_16-2.2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.5 KiB | [pg_stat_kcache_16-2.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_kcache_16-2.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_16` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.4 KiB | [pg_stat_kcache_16-2.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_kcache_16-2.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_16` | `2.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.6 KiB | [pg_stat_kcache_16-2.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_kcache_16-2.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_16` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.2 KiB | [pg_stat_kcache_16-2.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_kcache_16-2.2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_16` | `2.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.6 KiB | [pg_stat_kcache_16-2.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_kcache_16-2.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.5 KiB | [pg_stat_kcache_16-2.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_kcache_16-2.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.8 KiB | [pg_stat_kcache_16-2.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_kcache_16-2.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.3 KiB | [pg_stat_kcache_16-2.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_kcache_16-2.2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.0 KiB | [pg_stat_kcache_16-2.3.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_stat_kcache_16-2.3.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_stat_kcache_16` | `2.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.7 KiB | [pg_stat_kcache_16-2.3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_stat_kcache_16-2.3.0-2PGDG.rhel10.x86_64.rpm) |
| `pg_stat_kcache_16` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.5 KiB | [pg_stat_kcache_16-2.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_kcache_16-2.3.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.2 KiB | [pg_stat_kcache_16-2.3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_kcache_16-2.3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-stat-kcache` | `2.3.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.0 KiB | [postgresql-16-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.6 KiB | [postgresql-16-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 36.8 KiB | [postgresql-16-pg-stat-kcache_2.3.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.8 KiB | [postgresql-16-pg-stat-kcache_2.3.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.1 KiB | [postgresql-16-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.5 KiB | [postgresql-16-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.2 KiB | [postgresql-16-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.7 KiB | [postgresql-16-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_kcache_15` | `2.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.8 KiB | [pg_stat_kcache_15-2.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_kcache_15-2.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.4 KiB | [pg_stat_kcache_15-2.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_kcache_15-2.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.7 KiB | [pg_stat_kcache_15-2.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_kcache_15-2.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.3 KiB | [pg_stat_kcache_15-2.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_kcache_15-2.2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 47.2 KiB | [pg_stat_kcache_15-2.2.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_kcache_15-2.2.1-2.rhel8.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.8 KiB | [pg_stat_kcache_15-2.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_kcache_15-2.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.3 KiB | [pg_stat_kcache_15-2.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_kcache_15-2.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.6 KiB | [pg_stat_kcache_15-2.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_kcache_15-2.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.3 KiB | [pg_stat_kcache_15-2.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_kcache_15-2.2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 46.4 KiB | [pg_stat_kcache_15-2.2.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_kcache_15-2.2.1-2.rhel8.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.6 KiB | [pg_stat_kcache_15-2.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_kcache_15-2.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.5 KiB | [pg_stat_kcache_15-2.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_kcache_15-2.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.7 KiB | [pg_stat_kcache_15-2.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_kcache_15-2.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.3 KiB | [pg_stat_kcache_15-2.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_kcache_15-2.2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 47.4 KiB | [pg_stat_kcache_15-2.2.1-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_kcache_15-2.2.1-2.rhel9.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.6 KiB | [pg_stat_kcache_15-2.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_kcache_15-2.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.6 KiB | [pg_stat_kcache_15-2.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_kcache_15-2.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.9 KiB | [pg_stat_kcache_15-2.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_kcache_15-2.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.4 KiB | [pg_stat_kcache_15-2.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_kcache_15-2.2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 46.6 KiB | [pg_stat_kcache_15-2.2.1-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_kcache_15-2.2.1-2.rhel9.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.2 KiB | [pg_stat_kcache_15-2.3.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_stat_kcache_15-2.3.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.9 KiB | [pg_stat_kcache_15-2.3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_stat_kcache_15-2.3.0-2PGDG.rhel10.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.6 KiB | [pg_stat_kcache_15-2.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_kcache_15-2.3.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.3 KiB | [pg_stat_kcache_15-2.3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_kcache_15-2.3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-stat-kcache` | `2.3.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.0 KiB | [postgresql-15-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.6 KiB | [postgresql-15-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 36.9 KiB | [postgresql-15-pg-stat-kcache_2.3.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.8 KiB | [postgresql-15-pg-stat-kcache_2.3.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.1 KiB | [postgresql-15-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.6 KiB | [postgresql-15-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.3 KiB | [postgresql-15-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.8 KiB | [postgresql-15-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_kcache_14` | `2.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.8 KiB | [pg_stat_kcache_14-2.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_kcache_14-2.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_14` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.4 KiB | [pg_stat_kcache_14-2.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_kcache_14-2.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_14` | `2.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.7 KiB | [pg_stat_kcache_14-2.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_kcache_14-2.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_14` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.3 KiB | [pg_stat_kcache_14-2.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_kcache_14-2.2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_14` | `2.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 47.0 KiB | [pg_stat_kcache_14-2.2.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_kcache_14-2.2.1-2.rhel8.x86_64.rpm) |
| `pg_stat_kcache_14` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.6 KiB | [pg_stat_kcache_14-2.2.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_stat_kcache_14-2.2.0-2.rhel8.x86_64.rpm) |
| `pg_stat_kcache_14` | `2.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.7 KiB | [pg_stat_kcache_14-2.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_kcache_14-2.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_14` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.3 KiB | [pg_stat_kcache_14-2.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_kcache_14-2.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_14` | `2.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.6 KiB | [pg_stat_kcache_14-2.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_kcache_14-2.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_14` | `2.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.3 KiB | [pg_stat_kcache_14-2.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_kcache_14-2.2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_14` | `2.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 46.2 KiB | [pg_stat_kcache_14-2.2.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_stat_kcache_14-2.2.1-2.rhel8.aarch64.rpm) |
| `pg_stat_kcache_14` | `2.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.6 KiB | [pg_stat_kcache_14-2.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_kcache_14-2.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_14` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.4 KiB | [pg_stat_kcache_14-2.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_kcache_14-2.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_14` | `2.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.7 KiB | [pg_stat_kcache_14-2.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_kcache_14-2.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_14` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.3 KiB | [pg_stat_kcache_14-2.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_kcache_14-2.2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_14` | `2.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 47.2 KiB | [pg_stat_kcache_14-2.2.1-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_stat_kcache_14-2.2.1-2.rhel9.x86_64.rpm) |
| `pg_stat_kcache_14` | `2.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.6 KiB | [pg_stat_kcache_14-2.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_kcache_14-2.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_14` | `2.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.6 KiB | [pg_stat_kcache_14-2.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_kcache_14-2.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_14` | `2.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.9 KiB | [pg_stat_kcache_14-2.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_kcache_14-2.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_14` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.4 KiB | [pg_stat_kcache_14-2.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_kcache_14-2.2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_14` | `2.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 46.5 KiB | [pg_stat_kcache_14-2.2.1-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_stat_kcache_14-2.2.1-2.rhel9.aarch64.rpm) |
| `pg_stat_kcache_14` | `2.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.2 KiB | [pg_stat_kcache_14-2.3.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_stat_kcache_14-2.3.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_stat_kcache_14` | `2.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.9 KiB | [pg_stat_kcache_14-2.3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_stat_kcache_14-2.3.0-2PGDG.rhel10.x86_64.rpm) |
| `pg_stat_kcache_14` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.5 KiB | [pg_stat_kcache_14-2.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_stat_kcache_14-2.3.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_stat_kcache_14` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.2 KiB | [pg_stat_kcache_14-2.3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_stat_kcache_14-2.3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-stat-kcache` | `2.3.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 36.9 KiB | [postgresql-14-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-stat-kcache` | `2.3.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.6 KiB | [postgresql-14-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-stat-kcache` | `2.3.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 36.9 KiB | [postgresql-14-pg-stat-kcache_2.3.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-stat-kcache` | `2.3.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.8 KiB | [postgresql-14-pg-stat-kcache_2.3.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-stat-kcache` | `2.3.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 38.9 KiB | [postgresql-14-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-stat-kcache` | `2.3.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.4 KiB | [postgresql-14-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-stat-kcache` | `2.3.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.3 KiB | [postgresql-14-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-stat-kcache` | `2.3.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.8 KiB | [postgresql-14-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-14-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_kcache_13` | `2.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.7 KiB | [pg_stat_kcache_13-2.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_kcache_13-2.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_13` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.3 KiB | [pg_stat_kcache_13-2.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_kcache_13-2.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_13` | `2.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.6 KiB | [pg_stat_kcache_13-2.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_kcache_13-2.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_13` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.2 KiB | [pg_stat_kcache_13-2.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_kcache_13-2.2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_13` | `2.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.7 KiB | [pg_stat_kcache_13-2.2.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_kcache_13-2.2.1-2.rhel8.x86_64.rpm) |
| `pg_stat_kcache_13` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.4 KiB | [pg_stat_kcache_13-2.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_stat_kcache_13-2.2.0-1.rhel8.x86_64.rpm) |
| `pg_stat_kcache_13` | `2.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.8 KiB | [pg_stat_kcache_13-2.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_kcache_13-2.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_13` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.3 KiB | [pg_stat_kcache_13-2.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_kcache_13-2.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_13` | `2.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.7 KiB | [pg_stat_kcache_13-2.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_kcache_13-2.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_13` | `2.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.3 KiB | [pg_stat_kcache_13-2.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_kcache_13-2.2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_13` | `2.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 46.1 KiB | [pg_stat_kcache_13-2.2.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_stat_kcache_13-2.2.1-2.rhel8.aarch64.rpm) |
| `pg_stat_kcache_13` | `2.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.6 KiB | [pg_stat_kcache_13-2.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_kcache_13-2.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_13` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.4 KiB | [pg_stat_kcache_13-2.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_kcache_13-2.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_13` | `2.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.7 KiB | [pg_stat_kcache_13-2.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_kcache_13-2.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_13` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.3 KiB | [pg_stat_kcache_13-2.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_kcache_13-2.2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_stat_kcache_13` | `2.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 47.0 KiB | [pg_stat_kcache_13-2.2.1-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_stat_kcache_13-2.2.1-2.rhel9.x86_64.rpm) |
| `pg_stat_kcache_13` | `2.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.7 KiB | [pg_stat_kcache_13-2.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_kcache_13-2.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_13` | `2.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.6 KiB | [pg_stat_kcache_13-2.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_kcache_13-2.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_13` | `2.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.9 KiB | [pg_stat_kcache_13-2.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_kcache_13-2.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_13` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.4 KiB | [pg_stat_kcache_13-2.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_kcache_13-2.2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_stat_kcache_13` | `2.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 46.2 KiB | [pg_stat_kcache_13-2.2.1-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_stat_kcache_13-2.2.1-2.rhel9.aarch64.rpm) |
| `pg_stat_kcache_13` | `2.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.2 KiB | [pg_stat_kcache_13-2.3.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_stat_kcache_13-2.3.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_stat_kcache_13` | `2.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.9 KiB | [pg_stat_kcache_13-2.3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_stat_kcache_13-2.3.0-2PGDG.rhel10.x86_64.rpm) |
| `pg_stat_kcache_13` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.6 KiB | [pg_stat_kcache_13-2.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_stat_kcache_13-2.3.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_stat_kcache_13` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.3 KiB | [pg_stat_kcache_13-2.3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_stat_kcache_13-2.3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pg-stat-kcache` | `2.3.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 36.7 KiB | [postgresql-13-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-13-pg-stat-kcache_2.3.1-1.pgdg12+1_amd64.deb) |
| `postgresql-13-pg-stat-kcache` | `2.3.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.5 KiB | [postgresql-13-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-13-pg-stat-kcache_2.3.1-1.pgdg12+1_arm64.deb) |
| `postgresql-13-pg-stat-kcache` | `2.3.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 37.0 KiB | [postgresql-13-pg-stat-kcache_2.3.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-13-pg-stat-kcache_2.3.1-1.pgdg13+1_amd64.deb) |
| `postgresql-13-pg-stat-kcache` | `2.3.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.7 KiB | [postgresql-13-pg-stat-kcache_2.3.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-13-pg-stat-kcache_2.3.1-1.pgdg13+1_arm64.deb) |
| `postgresql-13-pg-stat-kcache` | `2.3.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.0 KiB | [postgresql-13-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-13-pg-stat-kcache_2.3.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pg-stat-kcache` | `2.3.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 37.9 KiB | [postgresql-13-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-13-pg-stat-kcache_2.3.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pg-stat-kcache` | `2.3.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.0 KiB | [postgresql-13-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-13-pg-stat-kcache_2.3.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pg-stat-kcache` | `2.3.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.8 KiB | [postgresql-13-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-13-pg-stat-kcache_2.3.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/powa-team/pg_stat_kcache" title="Repository" icon="github" subtitle="github.com/powa-team/pg_stat_kcache" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_stat_kcache;		# install via package name, for the active PG version

pig install pg_stat_kcache -v 18;   # install for PG 18
pig install pg_stat_kcache -v 17;   # install for PG 17
pig install pg_stat_kcache -v 16;   # install for PG 16
pig install pg_stat_kcache -v 15;   # install for PG 15
pig install pg_stat_kcache -v 14;   # install for PG 14
pig install pg_stat_kcache -v 13;   # install for PG 13

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pg_stat_statements, pg_stat_kcache';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_stat_kcache CASCADE; -- requires pg_stat_statements
```
