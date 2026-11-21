---
title: "pg_background"
linkTitle: "pg_background"
description: "Run SQL queries in the background"
weight: 1110
categories: ["TIME"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_background**](https://github.com/vibhorkum/pg_background) : Run SQL queries in the background


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1110** | {{< badge content="pg_background" link="https://github.com/vibhorkum/pg_background" >}} | {{< ext "pg_background" >}} | `2.0.5` | {{< category "TIME" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "pg_dbms_lock" >}} |
|   **See Also**    | {{< ext "pg_task" >}} {{< ext "pg_durable" >}} {{< ext "pg_later" >}} {{< ext "pg_dispatch" >}} {{< ext "pg_retry" >}} {{< ext "pglock" >}} {{< ext "pgmq" >}} {{< ext "ulak" >}} {{< ext "omni_worker" >}} {{< ext "pg_cron" >}} |

> [!Note] SQL 2.0.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.5` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pg_background` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.5` | {{< bg "19" "pg_background_19" "green" >}} {{< bg "18" "pg_background_18" "green" >}} {{< bg "17" "pg_background_17" "green" >}} {{< bg "16" "pg_background_16" "green" >}} {{< bg "15" "pg_background_15" "green" >}} | `pg_background_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.5` | {{< bg "19" "postgresql-19-pg-background" "green" >}} {{< bg "18" "postgresql-18-pg-background" "green" >}} {{< bg "17" "postgresql-17-pg-background" "green" >}} {{< bg "16" "postgresql-16-pg-background" "green" >}} {{< bg "15" "postgresql-15-pg-background" "green" >}} | `postgresql-$v-pg-background` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_background_19 : N/A 0" "gray" >}} | {{< bg "PGDG 2.0.5" "pg_background_18 : AVAIL 12" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_17 : AVAIL 13" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_16 : AVAIL 13" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_15 : AVAIL 14" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_background_19 : N/A 0" "gray" >}} | {{< bg "PGDG 2.0.5" "pg_background_18 : AVAIL 13" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_17 : AVAIL 14" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_16 : AVAIL 14" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_15 : AVAIL 15" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.0.5" "pg_background_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_15 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.0.5" "pg_background_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_15 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.0.5" "pg_background_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_15 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.0.5" "pg_background_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_18 : AVAIL 24" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_17 : AVAIL 24" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_16 : AVAIL 24" "blue" >}} | {{< bg "PGDG 2.0.5" "pg_background_15 : AVAIL 24" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.0.3" "postgresql-19-pg-background : AVAIL 3" "blue" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-background : AVAIL 4" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.0.3" "postgresql-19-pg-background : AVAIL 3" "blue" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-background : AVAIL 4" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.0.3" "postgresql-19-pg-background : AVAIL 3" "blue" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-background : AVAIL 4" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.0.3" "postgresql-19-pg-background : AVAIL 3" "blue" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-background : AVAIL 4" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.0.3" "postgresql-19-pg-background : AVAIL 3" "blue" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-background : AVAIL 4" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.0.3" "postgresql-19-pg-background : AVAIL 3" "blue" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-background : AVAIL 4" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.0.3" "postgresql-19-pg-background : AVAIL 3" "blue" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-background : AVAIL 4" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.0.3" "postgresql-19-pg-background : AVAIL 3" "blue" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-background : AVAIL 4" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 2.0.3" "postgresql-19-pg-background : AVAIL 3" "blue" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-background : AVAIL 4" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 2.0.3" "postgresql-19-pg-background : AVAIL 3" "blue" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-background : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-background : AVAIL 4" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_19` | `2.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 67.0 KiB | [pg_background_19-2.0.5-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-x86_64/pg_background_19-2.0.5-1PGDG.rhel9.9.x86_64.rpm) |
| `pg_background_19` | `2.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 65.9 KiB | [pg_background_19-2.0.5-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-aarch64/pg_background_19-2.0.5-1PGDG.rhel9.9.aarch64.rpm) |
| `pg_background_19` | `2.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 67.4 KiB | [pg_background_19-2.0.5-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-x86_64/pg_background_19-2.0.5-1PGDG.rhel10.3.x86_64.rpm) |
| `pg_background_19` | `2.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 66.4 KiB | [pg_background_19-2.0.5-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-aarch64/pg_background_19-2.0.5-1PGDG.rhel10.3.aarch64.rpm) |
| `postgresql-19-pg-background` | `2.0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 94.8 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg12+3_amd64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 69.9 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg12+2_amd64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 94.8 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg12+1_amd64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 93.1 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg12+3_arm64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 68.1 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg12+2_arm64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 93.1 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg12+1_arm64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 95.0 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg13+3_amd64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 70.1 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg13+2_amd64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 95.0 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg13+1_amd64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 93.4 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg13+3_arm64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 68.4 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg13+2_arm64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 93.4 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg13+1_arm64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 96.6 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg22.04+3_amd64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 71.5 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg22.04+2_amd64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 96.6 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 94.3 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg22.04+3_arm64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 69.3 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg22.04+2_arm64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 94.3 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 95.1 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg24.04+3_amd64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 70.1 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg24.04+2_amd64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 95.0 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 93.2 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg24.04+3_arm64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 68.2 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg24.04+2_arm64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 93.2 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 94.2 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg26.04+3_amd64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 69.3 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg26.04+2_amd64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 94.1 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 92.8 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg26.04+3_arm64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 67.8 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg26.04+2_arm64.deb) |
| `postgresql-19-pg-background` | `2.0.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 92.9 KiB | [postgresql-19-pg-background_2.0.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-19-pg-background_2.0.3-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_18` | `2.0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.7 KiB | [pg_background_18-2.0.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-2.0.5-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.0 KiB | [pg_background_18-2.0.4-3PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-2.0.4-3PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 106.1 KiB | [pg_background_18-2.0.4-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_background_18-2.0.4-1PGSTY.el8.x86_64.rpm) |
| `pg_background_18` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 67.5 KiB | [pg_background_18-2.0.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-2.0.4-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `2.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 66.1 KiB | [pg_background_18-2.0.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-2.0.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 65.0 KiB | [pg_background_18-2.0.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-2.0.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.1 KiB | [pg_background_18-2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `1.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.5 KiB | [pg_background_18-1.9.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.9.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.4 KiB | [pg_background_18-1.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.1 KiB | [pg_background_18-1.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pg_background_18-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_18` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.5 KiB | [pg_background_18-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_18` | `2.0.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 67.2 KiB | [pg_background_18-2.0.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-2.0.5-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 66.6 KiB | [pg_background_18-2.0.4-3PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-2.0.4-3PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 104.7 KiB | [pg_background_18-2.0.4-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_background_18-2.0.4-1PGSTY.el8.aarch64.rpm) |
| `pg_background_18` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 66.1 KiB | [pg_background_18-2.0.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-2.0.4-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `2.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.7 KiB | [pg_background_18-2.0.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-2.0.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.8 KiB | [pg_background_18-2.0.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-2.0.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 61.8 KiB | [pg_background_18-2.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-2.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 57.3 KiB | [pg_background_18-1.9.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.9.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.1 KiB | [pg_background_18-1.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pg_background_18-1.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.8 KiB | [pg_background_18-1.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_background_18-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_18` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.0 KiB | [pg_background_18-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_18` | `2.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 67.0 KiB | [pg_background_18-2.0.5-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_background_18-2.0.5-1PGDG.rhel9.9.x86_64.rpm) |
| `pg_background_18` | `2.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 105.9 KiB | [pg_background_18-2.0.4-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_background_18-2.0.4-1PGSTY.el9.x86_64.rpm) |
| `pg_background_18` | `2.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 65.9 KiB | [pg_background_18-2.0.5-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_background_18-2.0.5-1PGDG.rhel9.9.aarch64.rpm) |
| `pg_background_18` | `2.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 105.4 KiB | [pg_background_18-2.0.4-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_background_18-2.0.4-1PGSTY.el9.aarch64.rpm) |
| `pg_background_18` | `2.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 67.4 KiB | [pg_background_18-2.0.5-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_background_18-2.0.5-1PGDG.rhel10.3.x86_64.rpm) |
| `pg_background_18` | `2.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 106.4 KiB | [pg_background_18-2.0.4-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_background_18-2.0.4-1PGSTY.el10.x86_64.rpm) |
| `pg_background_18` | `2.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 66.4 KiB | [pg_background_18-2.0.5-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-2.0.5-1PGDG.rhel10.3.aarch64.rpm) |
| `pg_background_18` | `2.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 66.4 KiB | [pg_background_18-2.0.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-2.0.5-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_18` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 65.8 KiB | [pg_background_18-2.0.4-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-2.0.4-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_18` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 105.5 KiB | [pg_background_18-2.0.4-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_background_18-2.0.4-1PGSTY.el10.aarch64.rpm) |
| `pg_background_18` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 65.5 KiB | [pg_background_18-2.0.4-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-2.0.4-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_18` | `2.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.8 KiB | [pg_background_18-2.0.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-2.0.3-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_18` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.2 KiB | [pg_background_18-2.0.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-2.0.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_18` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.2 KiB | [pg_background_18-2.0.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-2.0.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_18` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.2 KiB | [pg_background_18-2.0.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-2.0.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_18` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.3 KiB | [pg_background_18-2.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-2.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_18` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.3 KiB | [pg_background_18-2.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-2.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_18` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.3 KiB | [pg_background_18-2.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-2.0-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_18` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_18-1.9.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.9.3-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_18` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_18-1.9.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.9.3-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_18` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_18-1.9.3-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.9.3-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_18` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.8 KiB | [pg_background_18-1.9.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.9.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_18` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.8 KiB | [pg_background_18-1.9.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.9.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_18` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.5 KiB | [pg_background_18-1.9.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.9.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_18` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.5 KiB | [pg_background_18-1.9.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.9.1-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_18` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_background_18-1.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.8-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_18` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_background_18-1.8-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.8-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_18` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_background_18-1.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.6-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_18` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_background_18-1.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.6-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_18` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.6 KiB | [pg_background_18-1.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_background_18-1.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-background` | `2.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 97.5 KiB | [postgresql-18-pg-background_2.0.4-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-18-pg-background_2.0.4-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 69.8 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg12+3_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 69.7 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg12+2_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 69.7 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 96.0 KiB | [postgresql-18-pg-background_2.0.4-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-18-pg-background_2.0.4-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 68.2 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg12+3_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 68.1 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg12+2_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 68.1 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 97.7 KiB | [postgresql-18-pg-background_2.0.4-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-18-pg-background_2.0.4-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 69.7 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg13+3_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 69.8 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg13+2_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 69.8 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 96.2 KiB | [postgresql-18-pg-background_2.0.4-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-18-pg-background_2.0.4-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 68.4 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg13+3_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 68.5 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg13+2_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 68.5 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 103.9 KiB | [postgresql-18-pg-background_2.0.4-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-18-pg-background_2.0.4-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 71.3 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg22.04+3_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 71.2 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg22.04+2_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 71.3 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 102.1 KiB | [postgresql-18-pg-background_2.0.4-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-18-pg-background_2.0.4-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 69.3 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg22.04+3_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 69.3 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg22.04+2_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 69.3 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 100.7 KiB | [postgresql-18-pg-background_2.0.4-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-18-pg-background_2.0.4-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 70.1 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg24.04+3_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 70.0 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg24.04+2_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 70.0 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 99.9 KiB | [postgresql-18-pg-background_2.0.4-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-18-pg-background_2.0.4-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 68.3 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg24.04+3_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 68.2 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg24.04+2_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 68.2 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 99.4 KiB | [postgresql-18-pg-background_2.0.4-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-18-pg-background_2.0.4-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 69.2 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg26.04+3_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 69.1 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg26.04+2_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 69.2 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pg-background` | `2.0.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 98.7 KiB | [postgresql-18-pg-background_2.0.4-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-18-pg-background_2.0.4-1PGSTY~resolute_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 67.7 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg26.04+3_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 67.7 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg26.04+2_arm64.deb) |
| `postgresql-18-pg-background` | `2.0.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 67.7 KiB | [postgresql-18-pg-background_2.0.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-18-pg-background_2.0.3-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_17` | `2.0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.7 KiB | [pg_background_17-2.0.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-2.0.5-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.0 KiB | [pg_background_17-2.0.4-3PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-2.0.4-3PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 106.1 KiB | [pg_background_17-2.0.4-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_background_17-2.0.4-1PGSTY.el8.x86_64.rpm) |
| `pg_background_17` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 67.5 KiB | [pg_background_17-2.0.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-2.0.4-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `2.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 66.1 KiB | [pg_background_17-2.0.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-2.0.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 65.0 KiB | [pg_background_17-2.0.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-2.0.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.1 KiB | [pg_background_17-2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `1.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.5 KiB | [pg_background_17-1.9.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.9.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.4 KiB | [pg_background_17-1.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.1 KiB | [pg_background_17-1.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [pg_background_17-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_17` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.9 KiB | [pg_background_17-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_17` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.1 KiB | [pg_background_17-1.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.2-2PGDG.rhel8.x86_64.rpm) |
| `pg_background_17` | `2.0.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 67.1 KiB | [pg_background_17-2.0.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-2.0.5-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 66.4 KiB | [pg_background_17-2.0.4-3PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-2.0.4-3PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 104.5 KiB | [pg_background_17-2.0.4-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_background_17-2.0.4-1PGSTY.el8.aarch64.rpm) |
| `pg_background_17` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 66.0 KiB | [pg_background_17-2.0.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-2.0.4-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `2.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.6 KiB | [pg_background_17-2.0.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-2.0.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.6 KiB | [pg_background_17-2.0.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-2.0.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 61.7 KiB | [pg_background_17-2.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-2.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 57.2 KiB | [pg_background_17-1.9.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.9.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.1 KiB | [pg_background_17-1.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pg_background_17-1.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.8 KiB | [pg_background_17-1.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_background_17-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_17` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.3 KiB | [pg_background_17-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_17` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.6 KiB | [pg_background_17-1.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.2-2PGDG.rhel8.aarch64.rpm) |
| `pg_background_17` | `2.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 66.9 KiB | [pg_background_17-2.0.5-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_background_17-2.0.5-1PGDG.rhel9.9.x86_64.rpm) |
| `pg_background_17` | `2.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 106.1 KiB | [pg_background_17-2.0.4-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_background_17-2.0.4-1PGSTY.el9.x86_64.rpm) |
| `pg_background_17` | `2.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 65.9 KiB | [pg_background_17-2.0.5-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_background_17-2.0.5-1PGDG.rhel9.9.aarch64.rpm) |
| `pg_background_17` | `2.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 105.2 KiB | [pg_background_17-2.0.4-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_background_17-2.0.4-1PGSTY.el9.aarch64.rpm) |
| `pg_background_17` | `2.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 67.4 KiB | [pg_background_17-2.0.5-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_background_17-2.0.5-1PGDG.rhel10.3.x86_64.rpm) |
| `pg_background_17` | `2.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 106.5 KiB | [pg_background_17-2.0.4-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_background_17-2.0.4-1PGSTY.el10.x86_64.rpm) |
| `pg_background_17` | `2.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 66.4 KiB | [pg_background_17-2.0.5-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-2.0.5-1PGDG.rhel10.3.aarch64.rpm) |
| `pg_background_17` | `2.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 66.4 KiB | [pg_background_17-2.0.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-2.0.5-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_17` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 65.7 KiB | [pg_background_17-2.0.4-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-2.0.4-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_17` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 105.1 KiB | [pg_background_17-2.0.4-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_background_17-2.0.4-1PGSTY.el10.aarch64.rpm) |
| `pg_background_17` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 65.4 KiB | [pg_background_17-2.0.4-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-2.0.4-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_17` | `2.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.8 KiB | [pg_background_17-2.0.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-2.0.3-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_17` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.1 KiB | [pg_background_17-2.0.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-2.0.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_17` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.1 KiB | [pg_background_17-2.0.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-2.0.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_17` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.1 KiB | [pg_background_17-2.0.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-2.0.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_17` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.2 KiB | [pg_background_17-2.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-2.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_17` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.2 KiB | [pg_background_17-2.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-2.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_17` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.2 KiB | [pg_background_17-2.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-2.0-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_17` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_17-1.9.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.9.3-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_17` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_17-1.9.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.9.3-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_17` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_17-1.9.3-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.9.3-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_17` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.8 KiB | [pg_background_17-1.9.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.9.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_17` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.8 KiB | [pg_background_17-1.9.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.9.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_17` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.6 KiB | [pg_background_17-1.9.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.9.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_17` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.6 KiB | [pg_background_17-1.9.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.9.1-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_17` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_background_17-1.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.8-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_17` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_background_17-1.8-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.8-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_17` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_background_17-1.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.6-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_17` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_background_17-1.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.6-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_17` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.3 KiB | [pg_background_17-1.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_background_17-1.3-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-background` | `2.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 97.4 KiB | [postgresql-17-pg-background_2.0.4-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-17-pg-background_2.0.4-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 69.7 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg12+3_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 69.7 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg12+2_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 69.7 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 95.9 KiB | [postgresql-17-pg-background_2.0.4-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-17-pg-background_2.0.4-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 68.2 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg12+3_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 68.1 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg12+2_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 68.1 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 97.6 KiB | [postgresql-17-pg-background_2.0.4-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-17-pg-background_2.0.4-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 69.9 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg13+3_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 70.0 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg13+2_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 69.9 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 96.1 KiB | [postgresql-17-pg-background_2.0.4-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-17-pg-background_2.0.4-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 68.4 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg13+3_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 68.4 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg13+2_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 68.4 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 111.7 KiB | [postgresql-17-pg-background_2.0.4-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-17-pg-background_2.0.4-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 78.9 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg22.04+3_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 78.8 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg22.04+2_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 78.8 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 110.1 KiB | [postgresql-17-pg-background_2.0.4-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-17-pg-background_2.0.4-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 77.0 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg22.04+3_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 76.9 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg22.04+2_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 76.9 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 100.6 KiB | [postgresql-17-pg-background_2.0.4-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-17-pg-background_2.0.4-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 70.0 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg24.04+3_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 70.0 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg24.04+2_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 70.1 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 99.7 KiB | [postgresql-17-pg-background_2.0.4-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-17-pg-background_2.0.4-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 68.3 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg24.04+3_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 68.2 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg24.04+2_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 68.2 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 99.3 KiB | [postgresql-17-pg-background_2.0.4-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-17-pg-background_2.0.4-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 68.8 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg26.04+3_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 69.0 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg26.04+2_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 68.9 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pg-background` | `2.0.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 98.5 KiB | [postgresql-17-pg-background_2.0.4-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-17-pg-background_2.0.4-1PGSTY~resolute_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 67.7 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg26.04+3_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 67.6 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg26.04+2_arm64.deb) |
| `postgresql-17-pg-background` | `2.0.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 67.6 KiB | [postgresql-17-pg-background_2.0.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-17-pg-background_2.0.3-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_16` | `2.0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.7 KiB | [pg_background_16-2.0.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-2.0.5-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.0 KiB | [pg_background_16-2.0.4-3PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-2.0.4-3PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 106.2 KiB | [pg_background_16-2.0.4-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_background_16-2.0.4-1PGSTY.el8.x86_64.rpm) |
| `pg_background_16` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 67.5 KiB | [pg_background_16-2.0.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-2.0.4-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `2.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 66.1 KiB | [pg_background_16-2.0.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-2.0.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 65.0 KiB | [pg_background_16-2.0.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-2.0.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.1 KiB | [pg_background_16-2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `1.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.5 KiB | [pg_background_16-1.9.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.9.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.3 KiB | [pg_background_16-1.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.1 KiB | [pg_background_16-1.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pg_background_16-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_16` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.9 KiB | [pg_background_16-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_16` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.7 KiB | [pg_background_16-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_16` | `2.0.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 67.1 KiB | [pg_background_16-2.0.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-2.0.5-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 66.4 KiB | [pg_background_16-2.0.4-3PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-2.0.4-3PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 104.5 KiB | [pg_background_16-2.0.4-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_background_16-2.0.4-1PGSTY.el8.aarch64.rpm) |
| `pg_background_16` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 65.9 KiB | [pg_background_16-2.0.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-2.0.4-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `2.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.5 KiB | [pg_background_16-2.0.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-2.0.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.6 KiB | [pg_background_16-2.0.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-2.0.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 61.6 KiB | [pg_background_16-2.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-2.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 57.2 KiB | [pg_background_16-1.9.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.9.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.1 KiB | [pg_background_16-1.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.8 KiB | [pg_background_16-1.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.7 KiB | [pg_background_16-1.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_background_16-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_16` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.3 KiB | [pg_background_16-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_16` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.2 KiB | [pg_background_16-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_16` | `2.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 66.9 KiB | [pg_background_16-2.0.5-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_background_16-2.0.5-1PGDG.rhel9.9.x86_64.rpm) |
| `pg_background_16` | `2.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 106.1 KiB | [pg_background_16-2.0.4-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_background_16-2.0.4-1PGSTY.el9.x86_64.rpm) |
| `pg_background_16` | `2.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 65.9 KiB | [pg_background_16-2.0.5-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_background_16-2.0.5-1PGDG.rhel9.9.aarch64.rpm) |
| `pg_background_16` | `2.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 105.2 KiB | [pg_background_16-2.0.4-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_background_16-2.0.4-1PGSTY.el9.aarch64.rpm) |
| `pg_background_16` | `2.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 67.4 KiB | [pg_background_16-2.0.5-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_background_16-2.0.5-1PGDG.rhel10.3.x86_64.rpm) |
| `pg_background_16` | `2.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 106.5 KiB | [pg_background_16-2.0.4-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_background_16-2.0.4-1PGSTY.el10.x86_64.rpm) |
| `pg_background_16` | `2.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 66.4 KiB | [pg_background_16-2.0.5-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-2.0.5-1PGDG.rhel10.3.aarch64.rpm) |
| `pg_background_16` | `2.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 66.4 KiB | [pg_background_16-2.0.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-2.0.5-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_16` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 65.8 KiB | [pg_background_16-2.0.4-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-2.0.4-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_16` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 105.2 KiB | [pg_background_16-2.0.4-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_background_16-2.0.4-1PGSTY.el10.aarch64.rpm) |
| `pg_background_16` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 65.4 KiB | [pg_background_16-2.0.4-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-2.0.4-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_16` | `2.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.8 KiB | [pg_background_16-2.0.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-2.0.3-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_16` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.1 KiB | [pg_background_16-2.0.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-2.0.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_16` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.1 KiB | [pg_background_16-2.0.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-2.0.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_16` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.1 KiB | [pg_background_16-2.0.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-2.0.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_16` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.3 KiB | [pg_background_16-2.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-2.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_16` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.3 KiB | [pg_background_16-2.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-2.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_16` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.3 KiB | [pg_background_16-2.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-2.0-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_16` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_16-1.9.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.9.3-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_16` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_16-1.9.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.9.3-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_16` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_16-1.9.3-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.9.3-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_16` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.8 KiB | [pg_background_16-1.9.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.9.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_16` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.8 KiB | [pg_background_16-1.9.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.9.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_16` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.5 KiB | [pg_background_16-1.9.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.9.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_16` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.6 KiB | [pg_background_16-1.9.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.9.1-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_16` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_background_16-1.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.8-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_16` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_background_16-1.8-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.8-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_16` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_background_16-1.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.6-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_16` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.8 KiB | [pg_background_16-1.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.6-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_16` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.4 KiB | [pg_background_16-1.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_background_16-1.3-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-background` | `2.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 97.3 KiB | [postgresql-16-pg-background_2.0.4-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-16-pg-background_2.0.4-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 69.7 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg12+3_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 69.6 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg12+2_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 69.6 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 95.9 KiB | [postgresql-16-pg-background_2.0.4-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-16-pg-background_2.0.4-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 68.2 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg12+3_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 68.1 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg12+2_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 68.1 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 97.6 KiB | [postgresql-16-pg-background_2.0.4-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-16-pg-background_2.0.4-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 70.0 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg13+3_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 69.9 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg13+2_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 69.9 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 96.1 KiB | [postgresql-16-pg-background_2.0.4-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-16-pg-background_2.0.4-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 68.4 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg13+3_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 68.3 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg13+2_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 68.3 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 111.6 KiB | [postgresql-16-pg-background_2.0.4-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-16-pg-background_2.0.4-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 78.8 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg22.04+3_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 78.8 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg22.04+2_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 78.8 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 110.0 KiB | [postgresql-16-pg-background_2.0.4-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-16-pg-background_2.0.4-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 76.9 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg22.04+3_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 76.8 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg22.04+2_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 76.8 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 100.6 KiB | [postgresql-16-pg-background_2.0.4-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-16-pg-background_2.0.4-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 69.9 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg24.04+3_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 70.0 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg24.04+2_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 70.0 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 99.8 KiB | [postgresql-16-pg-background_2.0.4-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-16-pg-background_2.0.4-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 68.3 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg24.04+3_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 68.3 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg24.04+2_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 68.2 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 99.3 KiB | [postgresql-16-pg-background_2.0.4-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-16-pg-background_2.0.4-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 69.1 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg26.04+3_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 69.0 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg26.04+2_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 69.0 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pg-background` | `2.0.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 98.6 KiB | [postgresql-16-pg-background_2.0.4-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-16-pg-background_2.0.4-1PGSTY~resolute_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 67.7 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg26.04+3_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 67.7 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg26.04+2_arm64.deb) |
| `postgresql-16-pg-background` | `2.0.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 67.6 KiB | [postgresql-16-pg-background_2.0.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-16-pg-background_2.0.3-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_background_15` | `2.0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.7 KiB | [pg_background_15-2.0.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-2.0.5-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.1 KiB | [pg_background_15-2.0.4-3PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-2.0.4-3PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 105.9 KiB | [pg_background_15-2.0.4-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_background_15-2.0.4-1PGSTY.el8.x86_64.rpm) |
| `pg_background_15` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 67.6 KiB | [pg_background_15-2.0.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-2.0.4-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `2.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 66.1 KiB | [pg_background_15-2.0.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-2.0.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 65.1 KiB | [pg_background_15-2.0.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-2.0.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.2 KiB | [pg_background_15-2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `1.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.5 KiB | [pg_background_15-1.9.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.9.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.3 KiB | [pg_background_15-1.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.1 KiB | [pg_background_15-1.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pg_background_15-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_background_15` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.9 KiB | [pg_background_15-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.6 KiB | [pg_background_15-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.3 KiB | [pg_background_15-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.0-1.rhel8.x86_64.rpm) |
| `pg_background_15` | `2.0.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 67.1 KiB | [pg_background_15-2.0.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-2.0.5-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 66.4 KiB | [pg_background_15-2.0.4-3PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-2.0.4-3PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 104.5 KiB | [pg_background_15-2.0.4-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_background_15-2.0.4-1PGSTY.el8.aarch64.rpm) |
| `pg_background_15` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 65.9 KiB | [pg_background_15-2.0.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-2.0.4-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `2.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.6 KiB | [pg_background_15-2.0.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-2.0.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.6 KiB | [pg_background_15-2.0.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-2.0.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 61.7 KiB | [pg_background_15-2.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-2.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 57.2 KiB | [pg_background_15-1.9.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.9.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.1 KiB | [pg_background_15-1.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.8 KiB | [pg_background_15-1.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.7 KiB | [pg_background_15-1.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_background_15-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_background_15` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.2 KiB | [pg_background_15-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.2 KiB | [pg_background_15-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.7 KiB | [pg_background_15-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.0-1.rhel8.aarch64.rpm) |
| `pg_background_15` | `2.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 67.0 KiB | [pg_background_15-2.0.5-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-2.0.5-1PGDG.rhel9.9.x86_64.rpm) |
| `pg_background_15` | `2.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 105.8 KiB | [pg_background_15-2.0.4-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_background_15-2.0.4-1PGSTY.el9.x86_64.rpm) |
| `pg_background_15` | `2.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 65.9 KiB | [pg_background_15-2.0.5-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-2.0.5-1PGDG.rhel9.9.aarch64.rpm) |
| `pg_background_15` | `2.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 105.2 KiB | [pg_background_15-2.0.4-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_background_15-2.0.4-1PGSTY.el9.aarch64.rpm) |
| `pg_background_15` | `2.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 67.4 KiB | [pg_background_15-2.0.5-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_background_15-2.0.5-1PGDG.rhel10.3.x86_64.rpm) |
| `pg_background_15` | `2.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 106.3 KiB | [pg_background_15-2.0.4-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_background_15-2.0.4-1PGSTY.el10.x86_64.rpm) |
| `pg_background_15` | `2.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 66.5 KiB | [pg_background_15-2.0.5-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-2.0.5-1PGDG.rhel10.3.aarch64.rpm) |
| `pg_background_15` | `2.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 66.5 KiB | [pg_background_15-2.0.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-2.0.5-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_15` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 65.8 KiB | [pg_background_15-2.0.4-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-2.0.4-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_15` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 105.3 KiB | [pg_background_15-2.0.4-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_background_15-2.0.4-1PGSTY.el10.aarch64.rpm) |
| `pg_background_15` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 65.5 KiB | [pg_background_15-2.0.4-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-2.0.4-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_15` | `2.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.8 KiB | [pg_background_15-2.0.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-2.0.3-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_15` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.2 KiB | [pg_background_15-2.0.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-2.0.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_15` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.2 KiB | [pg_background_15-2.0.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-2.0.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_15` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.2 KiB | [pg_background_15-2.0.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-2.0.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_15` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.4 KiB | [pg_background_15-2.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-2.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_15` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.4 KiB | [pg_background_15-2.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-2.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_15` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.4 KiB | [pg_background_15-2.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-2.0-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_15` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_15-1.9.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.9.3-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_background_15` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_15-1.9.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.9.3-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_15` | `1.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.8 KiB | [pg_background_15-1.9.3-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.9.3-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_15` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.8 KiB | [pg_background_15-1.9.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.9.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_15` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.8 KiB | [pg_background_15-1.9.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.9.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_15` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.6 KiB | [pg_background_15-1.9.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.9.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_15` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.6 KiB | [pg_background_15-1.9.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.9.1-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_15` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_background_15-1.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.8-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_15` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_background_15-1.8-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.8-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_15` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_background_15-1.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.6-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_background_15` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_background_15-1.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.6-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_background_15` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.3 KiB | [pg_background_15-1.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_background_15-1.3-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-background` | `2.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 97.6 KiB | [postgresql-15-pg-background_2.0.4-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-15-pg-background_2.0.4-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 70.1 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg12+3_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 69.9 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg12+2_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 69.9 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 95.6 KiB | [postgresql-15-pg-background_2.0.4-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-15-pg-background_2.0.4-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 68.0 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg12+3_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 68.0 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg12+2_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 68.0 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 97.6 KiB | [postgresql-15-pg-background_2.0.4-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-15-pg-background_2.0.4-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 70.0 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg13+3_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 70.0 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg13+2_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 70.0 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 95.9 KiB | [postgresql-15-pg-background_2.0.4-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-background/postgresql-15-pg-background_2.0.4-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 68.2 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg13+3_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 68.2 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg13+2_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 68.2 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 111.8 KiB | [postgresql-15-pg-background_2.0.4-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-15-pg-background_2.0.4-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 79.0 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg22.04+3_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 78.9 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg22.04+2_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 79.0 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 110.0 KiB | [postgresql-15-pg-background_2.0.4-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-15-pg-background_2.0.4-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 77.0 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg22.04+3_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 76.8 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg22.04+2_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 76.8 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 100.6 KiB | [postgresql-15-pg-background_2.0.4-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-15-pg-background_2.0.4-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 69.8 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg24.04+3_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 69.8 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg24.04+2_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 69.8 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 99.8 KiB | [postgresql-15-pg-background_2.0.4-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-15-pg-background_2.0.4-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 68.4 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg24.04+3_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 68.2 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg24.04+2_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 68.2 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 99.3 KiB | [postgresql-15-pg-background_2.0.4-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-15-pg-background_2.0.4-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 69.2 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg26.04+3_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 69.2 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg26.04+2_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 69.2 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pg-background` | `2.0.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 98.7 KiB | [postgresql-15-pg-background_2.0.4-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-background/postgresql-15-pg-background_2.0.4-1PGSTY~resolute_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 67.6 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg26.04+3_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 67.5 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg26.04+2_arm64.deb) |
| `postgresql-15-pg-background` | `2.0.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 67.5 KiB | [postgresql-15-pg-background_2.0.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-background/postgresql-15-pg-background_2.0.3-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/vibhorkum/pg_background" title="Repository" icon="github" subtitle="github.com/vibhorkum/pg_background" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_background-2.0.5.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_background;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_background;		# install via package name, for the active PG version

pig install pg_background -v 19;   # install for PG 19
pig install pg_background -v 18;   # install for PG 18
pig install pg_background -v 17;   # install for PG 17
pig install pg_background -v 16;   # install for PG 16
pig install pg_background -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_background;
```

## Usage

Sources:

- [v2.0.5 README](https://github.com/vibhorkum/pg_background/blob/v2.0.5/README.md)
- [v2.0.5 release](https://github.com/vibhorkum/pg_background/releases/tag/v2.0.5)
- [v2.0.5 SQL API](https://github.com/vibhorkum/pg_background/blob/v2.0.5/extension/pg_background--2.0.sql)

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

Use the launch/result pattern when the background SQL returns rows. In psql, `\gset` saves the handle for subsequent commands:

```sql
SELECT pid, cookie FROM pg_background_launch(
  'SELECT count(*) FROM large_table',
  queue_size := 65536,
  label := 'count-large-table'
) \gset bg_

SELECT * FROM pg_background_result(:bg_pid, :bg_cookie) AS (count bigint);
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
SELECT pid, cookie FROM pg_background_launch($$
  SELECT pg_background_report_progress(0, 'starting');
  SELECT pg_sleep(1);
  SELECT pg_background_report_progress(100, 'done');
$$) \gset bg_

SELECT * FROM pg_background_get_progress(:bg_pid, :bg_cookie);
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

- Upstream release 2.0.5 retains SQL/control version 2.0 and targets PostgreSQL 14-18 plus PostgreSQL 19 beta.
- Upgrades from pre-1.8 installs must first reach the 1.8/1.10 release line before updating to 2.0.
- The original v1 PID-only API was removed. Unsuffixed names now have cookie-protected semantics and return/use `(pid, cookie)`.
- `pg_background_cancel_v2_grace` and `pg_background_wait_v2_timeout` are folded into `pg_background_cancel(..., grace_ms)` and `pg_background_wait(..., timeout_ms)`.
- `pg_background_status_v2` was removed; use `pg_background_outcome(pid, cookie)`.

### 2.0.4 Changes and Privileges

2.0.4 fixes visibility between commands in one worker, complex result-type transport, and command completion tags. A nonzero `pg_background.worker_timeout` bounds the whole SQL string, excluding commit; when zero, `statement_timeout` applies to each statement separately. Start new workers after installing the updated library to use these fixes.

Installation requires a superuser. PUBLIC receives no execution rights by default; an administrator grants `pgbackground_role` or uses `pg_background_grant_privileges`. Workers commit independently, so rolling back the caller does not undo completed background operations.

### 2.0.5 Launch Handle Fix

2.0.5 fixes a race where a worker finished before its startup check and launch returned the wrong PID. Use the updated library for new workers and continue retaining both handle fields. SQL/control stays at 2.0, with no new API or SQL migration for this patch.
