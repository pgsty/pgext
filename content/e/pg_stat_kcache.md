---
title: "pg_stat_kcache"
linkTitle: "pg_stat_kcache"
description: "Kernel statistics gathering"
weight: 6220
categories: ["STAT"]
languages: ["C"]
licenses: ["BSD-3-Clause"]
repos: ["PGDG"]
page_width: full
---

[**pg_stat_kcache**](https://github.com/powa-team/pg_stat_kcache) : Kernel statistics gathering


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6220** | {{< badge content="pg_stat_kcache" link="https://github.com/powa-team/pg_stat_kcache" >}} | {{< ext "pg_stat_kcache" >}} | `2.3.2` | {{< category "STAT" >}} | {{< license "BSD-3-Clause" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pg_stat_statements" >}} |
|   **See Also**    | {{< ext "pgnodemx" >}} {{< ext "system_stats" >}} {{< ext "pgmeminfo" >}} {{< ext "pgmonitor" >}} {{< ext "pg_stat_ch" >}} {{< ext "pg_datasentinel" >}} {{< ext "pgsentinel" >}} {{< ext "pg_rusage" >}} |

> [!Note] PGDG 2.3.2 covers listed EL/D/U x86_64/aarch64; requires pg_stat_statements and preload.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.3.2` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pg_stat_kcache` | `pg_stat_statements` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.3.2` | {{< bg "19" "pg_stat_kcache_19" "green" >}} {{< bg "18" "pg_stat_kcache_18" "green" >}} {{< bg "17" "pg_stat_kcache_17" "green" >}} {{< bg "16" "pg_stat_kcache_16" "green" >}} {{< bg "15" "pg_stat_kcache_15" "green" >}} | `pg_stat_kcache_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.3.2` | {{< bg "19" "postgresql-19-pg-stat-kcache" "green" >}} {{< bg "18" "postgresql-18-pg-stat-kcache" "green" >}} {{< bg "17" "postgresql-17-pg-stat-kcache" "green" >}} {{< bg "16" "postgresql-16-pg-stat-kcache" "green" >}} {{< bg "15" "postgresql-15-pg-stat-kcache" "green" >}} | `postgresql-$v-pg-stat-kcache` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_stat_kcache_19 : N/A 0" "gray" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_15 : AVAIL 7" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_stat_kcache_19 : N/A 0" "gray" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_15 : AVAIL 7" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_15 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_15 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_15 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.3.2" "pg_stat_kcache_15 : AVAIL 6" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.3.2" "postgresql-19-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-18-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-17-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-16-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-15-pg-stat-kcache : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.3.2" "postgresql-19-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-18-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-17-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-16-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-15-pg-stat-kcache : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.3.2" "postgresql-19-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-18-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-17-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-16-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-15-pg-stat-kcache : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.3.2" "postgresql-19-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-18-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-17-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-16-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-15-pg-stat-kcache : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.3.2" "postgresql-19-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-18-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-17-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-16-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-15-pg-stat-kcache : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.3.2" "postgresql-19-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-18-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-17-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-16-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-15-pg-stat-kcache : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.3.2" "postgresql-19-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-18-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-17-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-16-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-15-pg-stat-kcache : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.3.2" "postgresql-19-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-18-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-17-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-16-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-15-pg-stat-kcache : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 2.3.2" "postgresql-19-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-18-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-17-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-16-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-15-pg-stat-kcache : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 2.3.2" "postgresql-19-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-18-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-17-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-16-pg-stat-kcache : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.3.2" "postgresql-15-pg-stat-kcache : AVAIL 3" "blue" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_kcache_19` | `2.3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.8 KiB | [pg_stat_kcache_19-2.3.2-7PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-x86_64/pg_stat_kcache_19-2.3.2-7PGDG.rhel9.9.x86_64.rpm) |
| `pg_stat_kcache_19` | `2.3.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 27.9 KiB | [pg_stat_kcache_19-2.3.2-7PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-aarch64/pg_stat_kcache_19-2.3.2-7PGDG.rhel9.9.aarch64.rpm) |
| `pg_stat_kcache_19` | `2.3.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 28.0 KiB | [pg_stat_kcache_19-2.3.2-7PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-x86_64/pg_stat_kcache_19-2.3.2-7PGDG.rhel10.3.x86_64.rpm) |
| `pg_stat_kcache_19` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.4 KiB | [pg_stat_kcache_19-2.3.2-7PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-aarch64/pg_stat_kcache_19-2.3.2-7PGDG.rhel10.3.aarch64.rpm) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.5 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg12+3_amd64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.5 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg12+2_amd64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.5 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg12+1_amd64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 37.1 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg12+3_arm64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 37.1 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg12+2_arm64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 37.0 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg12+1_arm64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 37.5 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg13+3_amd64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 37.5 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg13+2_amd64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 37.6 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg13+1_amd64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 37.1 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg13+3_arm64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 37.1 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg13+2_arm64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 37.1 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg13+1_arm64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 36.2 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg22.04+3_amd64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 36.1 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg22.04+2_amd64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 36.1 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 35.5 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg22.04+3_arm64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 35.4 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg22.04+2_arm64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 35.5 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.6 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg24.04+3_amd64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.7 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg24.04+2_amd64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.6 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 35.0 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg24.04+3_arm64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 35.0 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg24.04+2_arm64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.9 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 35.4 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg26.04+3_amd64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 35.4 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg26.04+2_amd64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 35.5 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 34.5 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg26.04+3_arm64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 34.5 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg26.04+2_arm64.deb) |
| `postgresql-19-pg-stat-kcache` | `2.3.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 34.5 KiB | [postgresql-19-pg-stat-kcache_2.3.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-19-pg-stat-kcache_2.3.2-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_kcache_18` | `2.3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.3 KiB | [pg_stat_kcache_18-2.3.2-7PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_stat_kcache_18-2.3.2-7PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_kcache_18` | `2.3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.0 KiB | [pg_stat_kcache_18-2.3.2-6PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_stat_kcache_18-2.3.2-6PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_kcache_18` | `2.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.8 KiB | [pg_stat_kcache_18-2.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_stat_kcache_18-2.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_18` | `2.3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.2 KiB | [pg_stat_kcache_18-2.3.2-7PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_stat_kcache_18-2.3.2-7PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_kcache_18` | `2.3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.9 KiB | [pg_stat_kcache_18-2.3.2-6PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_stat_kcache_18-2.3.2-6PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_kcache_18` | `2.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.7 KiB | [pg_stat_kcache_18-2.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_stat_kcache_18-2.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_18` | `2.3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.7 KiB | [pg_stat_kcache_18-2.3.2-7PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_stat_kcache_18-2.3.2-7PGDG.rhel9.9.x86_64.rpm) |
| `pg_stat_kcache_18` | `2.3.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 27.9 KiB | [pg_stat_kcache_18-2.3.2-7PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_stat_kcache_18-2.3.2-7PGDG.rhel9.9.aarch64.rpm) |
| `pg_stat_kcache_18` | `2.3.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.9 KiB | [pg_stat_kcache_18-2.3.2-7PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_stat_kcache_18-2.3.2-7PGDG.rhel10.3.x86_64.rpm) |
| `pg_stat_kcache_18` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.5 KiB | [pg_stat_kcache_18-2.3.2-7PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_kcache_18-2.3.2-7PGDG.rhel10.3.aarch64.rpm) |
| `pg_stat_kcache_18` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.5 KiB | [pg_stat_kcache_18-2.3.2-7PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_kcache_18-2.3.2-7PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_kcache_18` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.2 KiB | [pg_stat_kcache_18-2.3.2-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_kcache_18-2.3.2-6PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_kcache_18` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.4 KiB | [pg_stat_kcache_18-2.3.1-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_kcache_18-2.3.1-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_kcache_18` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.5 KiB | [pg_stat_kcache_18-2.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_kcache_18-2.3.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.1 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg12+3_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.1 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg12+2_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.0 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.8 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg12+3_arm64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.7 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg12+2_arm64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.7 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 37.0 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg13+3_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 37.0 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg13+2_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 37.0 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 37.0 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg13+3_arm64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.9 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg13+2_arm64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.9 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 35.9 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg22.04+3_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 35.9 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg22.04+2_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 35.8 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 35.3 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg22.04+3_arm64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 35.3 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg22.04+2_arm64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 35.3 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.4 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg24.04+3_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.2 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg24.04+2_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.2 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.9 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg24.04+3_arm64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.8 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg24.04+2_arm64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.8 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 34.9 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg26.04+3_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 34.9 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg26.04+2_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 34.9 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 34.4 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg26.04+3_arm64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 34.4 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg26.04+2_arm64.deb) |
| `postgresql-18-pg-stat-kcache` | `2.3.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 34.4 KiB | [postgresql-18-pg-stat-kcache_2.3.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-18-pg-stat-kcache_2.3.2-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_kcache_17` | `2.3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.3 KiB | [pg_stat_kcache_17-2.3.2-7PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_kcache_17-2.3.2-7PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_kcache_17` | `2.3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.0 KiB | [pg_stat_kcache_17-2.3.2-6PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_kcache_17-2.3.2-6PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_kcache_17` | `2.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.8 KiB | [pg_stat_kcache_17-2.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_kcache_17-2.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_17` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.3 KiB | [pg_stat_kcache_17-2.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_kcache_17-2.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_17` | `2.3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.2 KiB | [pg_stat_kcache_17-2.3.2-7PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_kcache_17-2.3.2-7PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_kcache_17` | `2.3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.0 KiB | [pg_stat_kcache_17-2.3.2-6PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_kcache_17-2.3.2-6PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_kcache_17` | `2.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.7 KiB | [pg_stat_kcache_17-2.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_kcache_17-2.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_17` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.2 KiB | [pg_stat_kcache_17-2.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_kcache_17-2.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_17` | `2.3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.8 KiB | [pg_stat_kcache_17-2.3.2-7PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_stat_kcache_17-2.3.2-7PGDG.rhel9.9.x86_64.rpm) |
| `pg_stat_kcache_17` | `2.3.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 27.9 KiB | [pg_stat_kcache_17-2.3.2-7PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_stat_kcache_17-2.3.2-7PGDG.rhel9.9.aarch64.rpm) |
| `pg_stat_kcache_17` | `2.3.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 28.0 KiB | [pg_stat_kcache_17-2.3.2-7PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_stat_kcache_17-2.3.2-7PGDG.rhel10.3.x86_64.rpm) |
| `pg_stat_kcache_17` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.5 KiB | [pg_stat_kcache_17-2.3.2-7PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_kcache_17-2.3.2-7PGDG.rhel10.3.aarch64.rpm) |
| `pg_stat_kcache_17` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.5 KiB | [pg_stat_kcache_17-2.3.2-7PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_kcache_17-2.3.2-7PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_kcache_17` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.2 KiB | [pg_stat_kcache_17-2.3.2-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_kcache_17-2.3.2-6PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_kcache_17` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.4 KiB | [pg_stat_kcache_17-2.3.1-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_kcache_17-2.3.1-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_kcache_17` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.5 KiB | [pg_stat_kcache_17-2.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_kcache_17-2.3.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_stat_kcache_17` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.2 KiB | [pg_stat_kcache_17-2.3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_kcache_17-2.3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.1 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg12+3_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.0 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg12+2_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.0 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.8 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg12+3_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.8 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg12+2_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.7 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 37.0 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg13+3_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 36.9 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg13+2_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 36.9 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.9 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg13+3_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.9 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg13+2_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.9 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.2 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg22.04+3_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.2 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg22.04+2_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.2 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.6 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg22.04+3_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.6 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg22.04+2_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.6 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.3 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg24.04+3_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.3 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg24.04+2_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.3 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.8 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg24.04+3_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.8 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg24.04+2_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.8 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 35.0 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg26.04+3_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 35.0 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg26.04+2_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 34.9 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 34.5 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg26.04+3_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 34.4 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg26.04+2_arm64.deb) |
| `postgresql-17-pg-stat-kcache` | `2.3.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 34.4 KiB | [postgresql-17-pg-stat-kcache_2.3.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-17-pg-stat-kcache_2.3.2-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_kcache_16` | `2.3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.3 KiB | [pg_stat_kcache_16-2.3.2-7PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_kcache_16-2.3.2-7PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_kcache_16` | `2.3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.0 KiB | [pg_stat_kcache_16-2.3.2-6PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_kcache_16-2.3.2-6PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_kcache_16` | `2.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.8 KiB | [pg_stat_kcache_16-2.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_kcache_16-2.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_16` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.3 KiB | [pg_stat_kcache_16-2.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_kcache_16-2.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_16` | `2.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.6 KiB | [pg_stat_kcache_16-2.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_kcache_16-2.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_16` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.3 KiB | [pg_stat_kcache_16-2.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_kcache_16-2.2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_16` | `2.3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.2 KiB | [pg_stat_kcache_16-2.3.2-7PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_kcache_16-2.3.2-7PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.0 KiB | [pg_stat_kcache_16-2.3.2-6PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_kcache_16-2.3.2-6PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.7 KiB | [pg_stat_kcache_16-2.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_kcache_16-2.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.2 KiB | [pg_stat_kcache_16-2.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_kcache_16-2.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.6 KiB | [pg_stat_kcache_16-2.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_kcache_16-2.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.3 KiB | [pg_stat_kcache_16-2.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_kcache_16-2.2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.8 KiB | [pg_stat_kcache_16-2.3.2-7PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_kcache_16-2.3.2-7PGDG.rhel9.9.x86_64.rpm) |
| `pg_stat_kcache_16` | `2.3.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 27.9 KiB | [pg_stat_kcache_16-2.3.2-7PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_kcache_16-2.3.2-7PGDG.rhel9.9.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.3.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.9 KiB | [pg_stat_kcache_16-2.3.2-7PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_stat_kcache_16-2.3.2-7PGDG.rhel10.3.x86_64.rpm) |
| `pg_stat_kcache_16` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.5 KiB | [pg_stat_kcache_16-2.3.2-7PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_kcache_16-2.3.2-7PGDG.rhel10.3.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.5 KiB | [pg_stat_kcache_16-2.3.2-7PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_kcache_16-2.3.2-7PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.2 KiB | [pg_stat_kcache_16-2.3.2-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_kcache_16-2.3.2-6PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.4 KiB | [pg_stat_kcache_16-2.3.1-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_kcache_16-2.3.1-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.5 KiB | [pg_stat_kcache_16-2.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_kcache_16-2.3.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_stat_kcache_16` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.2 KiB | [pg_stat_kcache_16-2.3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_kcache_16-2.3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.1 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg12+3_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.1 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg12+2_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.1 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.8 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg12+3_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.8 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg12+2_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.7 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 37.0 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg13+3_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 37.0 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg13+2_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 36.9 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.9 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg13+3_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.9 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg13+2_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.9 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.1 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg22.04+3_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.1 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg22.04+2_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.1 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.6 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg22.04+3_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.6 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg22.04+2_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.6 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.3 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg24.04+3_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.3 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg24.04+2_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.3 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.9 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg24.04+3_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.9 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg24.04+2_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.8 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 35.1 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg26.04+3_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 35.0 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg26.04+2_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 35.0 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 34.4 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg26.04+3_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 34.4 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg26.04+2_arm64.deb) |
| `postgresql-16-pg-stat-kcache` | `2.3.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 34.4 KiB | [postgresql-16-pg-stat-kcache_2.3.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-16-pg-stat-kcache_2.3.2-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_kcache_15` | `2.3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.3 KiB | [pg_stat_kcache_15-2.3.2-7PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_kcache_15-2.3.2-7PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.1 KiB | [pg_stat_kcache_15-2.3.2-6PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_kcache_15-2.3.2-6PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.8 KiB | [pg_stat_kcache_15-2.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_kcache_15-2.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.4 KiB | [pg_stat_kcache_15-2.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_kcache_15-2.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.7 KiB | [pg_stat_kcache_15-2.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_kcache_15-2.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.3 KiB | [pg_stat_kcache_15-2.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_kcache_15-2.2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 47.2 KiB | [pg_stat_kcache_15-2.2.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_stat_kcache_15-2.2.1-2.rhel8.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [pg_stat_kcache_15-2.3.2-7PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_kcache_15-2.3.2-7PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.0 KiB | [pg_stat_kcache_15-2.3.2-6PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_kcache_15-2.3.2-6PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.8 KiB | [pg_stat_kcache_15-2.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_kcache_15-2.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.3 KiB | [pg_stat_kcache_15-2.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_kcache_15-2.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.6 KiB | [pg_stat_kcache_15-2.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_kcache_15-2.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.3 KiB | [pg_stat_kcache_15-2.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_kcache_15-2.2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 46.4 KiB | [pg_stat_kcache_15-2.2.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_stat_kcache_15-2.2.1-2.rhel8.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.8 KiB | [pg_stat_kcache_15-2.3.2-7PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_stat_kcache_15-2.3.2-7PGDG.rhel9.9.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.3.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.0 KiB | [pg_stat_kcache_15-2.3.2-7PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_stat_kcache_15-2.3.2-7PGDG.rhel9.9.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.3.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 28.1 KiB | [pg_stat_kcache_15-2.3.2-7PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_stat_kcache_15-2.3.2-7PGDG.rhel10.3.x86_64.rpm) |
| `pg_stat_kcache_15` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.5 KiB | [pg_stat_kcache_15-2.3.2-7PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_kcache_15-2.3.2-7PGDG.rhel10.3.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.6 KiB | [pg_stat_kcache_15-2.3.2-7PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_kcache_15-2.3.2-7PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.3 KiB | [pg_stat_kcache_15-2.3.2-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_kcache_15-2.3.2-6PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.5 KiB | [pg_stat_kcache_15-2.3.1-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_kcache_15-2.3.1-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.6 KiB | [pg_stat_kcache_15-2.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_kcache_15-2.3.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_stat_kcache_15` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.3 KiB | [pg_stat_kcache_15-2.3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_stat_kcache_15-2.3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.1 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg12+3_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.1 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg12+2_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.1 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.8 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg12+3_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.8 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg12+2_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.7 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 37.0 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg13+3_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 37.0 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg13+2_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 36.9 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.9 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg13+3_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.9 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg13+2_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.9 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.2 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg22.04+3_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.2 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg22.04+2_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.2 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.7 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg22.04+3_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.7 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg22.04+2_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.6 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.4 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg24.04+3_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.3 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg24.04+2_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.3 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 35.0 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg24.04+3_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.9 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg24.04+2_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.9 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 35.2 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg26.04+3_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 35.2 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg26.04+2_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 35.1 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 34.4 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg26.04+3_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 34.4 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg26.04+2_arm64.deb) |
| `postgresql-15-pg-stat-kcache` | `2.3.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 34.4 KiB | [postgresql-15-pg-stat-kcache_2.3.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-kcache/postgresql-15-pg-stat-kcache_2.3.2-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/powa-team/pg_stat_kcache" title="Repository" icon="github" subtitle="github.com/powa-team/pg_stat_kcache" />}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_stat_kcache;		# install via package name, for the active PG version

pig install pg_stat_kcache -v 19;   # install for PG 19
pig install pg_stat_kcache -v 18;   # install for PG 18
pig install pg_stat_kcache -v 17;   # install for PG 17
pig install pg_stat_kcache -v 16;   # install for PG 16
pig install pg_stat_kcache -v 15;   # install for PG 15

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_stat_statements, pg_stat_kcache';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_stat_kcache CASCADE; -- requires pg_stat_statements
```

## Usage

Sources:

- [v2.3.2 release](https://github.com/powa-team/pg_stat_kcache/releases/tag/REL2_3_2)
- [Official README](https://github.com/powa-team/pg_stat_kcache)
- [Extension control file](https://github.com/powa-team/pg_stat_kcache/blob/master/pg_stat_kcache.control)

`pg_stat_kcache` 2.3.2 gathers statistics about real reads and writes done by the filesystem layer, as well as CPU usage per query. It requires `pg_stat_statements`; this release adds PostgreSQL 19 compatibility and fixes installation on RHEL-family systems.

### Views

**`pg_stat_kcache`** -- aggregated per-database statistics:

| Column | Type | Description |
|--------|------|-------------|
| `datname` | name | Database name |
| `exec_user_time` | double precision | User CPU time executing statements (seconds) |
| `exec_system_time` | double precision | System CPU time executing statements (seconds) |
| `exec_reads` | bigint | Bytes read by filesystem layer |
| `exec_reads_blks` | bigint | 8K blocks read by filesystem layer |
| `exec_writes` | bigint | Bytes written by filesystem layer |
| `exec_writes_blks` | bigint | 8K blocks written by filesystem layer |
| `plan_user_time` | double precision | User CPU time planning (if tracking enabled) |
| `plan_system_time` | double precision | System CPU time planning (if tracking enabled) |

**`pg_stat_kcache_detail`** -- per-query statistics including query text, role, and database.

### Functions

```sql
-- Reset all collected statistics
SELECT pg_stat_kcache_reset();
```

### Configuration

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pg_stat_kcache.linux_hz` | -1 | Linux CONFIG_HZ (auto-detected) |
| `pg_stat_kcache.track` | `top` | Track: `top`, `all`, or `none` |
| `pg_stat_kcache.track_planning` | `off` | Track planning statistics (PG 13+) |

### Example

```sql
SELECT datname, exec_user_time, exec_system_time, exec_reads, exec_writes
FROM pg_stat_kcache;

SELECT query, exec_user_time, exec_system_time, exec_reads
FROM pg_stat_kcache_detail
ORDER BY exec_user_time DESC
LIMIT 10;
```
