---
title: "mimeo"
linkTitle: "mimeo"
description: "Extension for specialized, per-table replication between PostgreSQL instances"
weight: 9700
categories: ["ETL"]
width: full
---

[**mimeo**](https://github.com/omniti-labs/mimeo) : Extension for specialized, per-table replication between PostgreSQL instances


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9700** | {{< badge content="mimeo" link="https://github.com/omniti-labs/mimeo" >}} | {{< ext "mimeo" >}} | `1.5.1` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "dblink" >}} |
|   **See Also**    | {{< ext "pg_jobmon" >}} {{< ext "postgres_fdw" >}} {{< ext "pglogical" >}} {{< ext "pg_cron" >}} {{< ext "pg_partman" >}} {{< ext "repmgr" >}} {{< ext "pg_fact_loader" >}} {{< ext "pg_failover_slots" >}} |

> [!Note] name conflict with pg_partman


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.5.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `mimeo` | `dblink` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5.1` | {{< bg "18" "mimeo_18" "green" >}} {{< bg "17" "mimeo_17" "green" >}} {{< bg "16" "mimeo_16" "green" >}} {{< bg "15" "mimeo_15" "green" >}} {{< bg "14" "mimeo_14" "green" >}} | `mimeo_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.5.1` | {{< bg "18" "postgresql-18-mimeo" "green" >}} {{< bg "17" "postgresql-17-mimeo" "green" >}} {{< bg "16" "postgresql-16-mimeo" "green" >}} {{< bg "15" "postgresql-15-mimeo" "green" >}} {{< bg "14" "postgresql-14-mimeo" "green" >}} | `postgresql-$v-mimeo` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-mimeo : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-mimeo : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-mimeo : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-mimeo : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-mimeo : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-mimeo : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-mimeo : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-mimeo : AVAIL 1" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-mimeo : AVAIL 1" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-mimeo : AVAIL 1" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mimeo_18` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 139.7 KiB | [mimeo_18-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/mimeo_18-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `mimeo_18` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 139.7 KiB | [mimeo_18-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/mimeo_18-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `mimeo_18` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 113.0 KiB | [mimeo_18-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/mimeo_18-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `mimeo_18` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 113.0 KiB | [mimeo_18-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/mimeo_18-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `mimeo_18` | `1.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 114.5 KiB | [mimeo_18-1.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/mimeo_18-1.5.1-1PIGSTY.el10.x86_64.rpm) |
| `mimeo_18` | `1.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 114.5 KiB | [mimeo_18-1.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/mimeo_18-1.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-mimeo` | `1.5.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 125.6 KiB | [postgresql-18-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-18-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-18-mimeo` | `1.5.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 125.6 KiB | [postgresql-18-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-18-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-18-mimeo` | `1.5.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 125.6 KiB | [postgresql-18-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-18-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-18-mimeo` | `1.5.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 125.6 KiB | [postgresql-18-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-18-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-18-mimeo` | `1.5.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 108.0 KiB | [postgresql-18-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-18-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-18-mimeo` | `1.5.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 108.0 KiB | [postgresql-18-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-18-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-18-mimeo` | `1.5.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 107.7 KiB | [postgresql-18-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-18-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |
| `postgresql-18-mimeo` | `1.5.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 107.7 KiB | [postgresql-18-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-18-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |
| `postgresql-18-mimeo` | `1.5.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 107.6 KiB | [postgresql-18-mimeo_1.5.1-20.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-18-mimeo_1.5.1-20.pgdg26.04+1_all.deb) |
| `postgresql-18-mimeo` | `1.5.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 107.6 KiB | [postgresql-18-mimeo_1.5.1-20.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-18-mimeo_1.5.1-20.pgdg26.04+1_all.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mimeo_17` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 139.7 KiB | [mimeo_17-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/mimeo_17-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `mimeo_17` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 139.7 KiB | [mimeo_17-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/mimeo_17-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `mimeo_17` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 113.0 KiB | [mimeo_17-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/mimeo_17-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `mimeo_17` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 113.0 KiB | [mimeo_17-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/mimeo_17-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `mimeo_17` | `1.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 114.5 KiB | [mimeo_17-1.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/mimeo_17-1.5.1-1PIGSTY.el10.x86_64.rpm) |
| `mimeo_17` | `1.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 114.5 KiB | [mimeo_17-1.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/mimeo_17-1.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-mimeo` | `1.5.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 125.7 KiB | [postgresql-17-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-17-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-17-mimeo` | `1.5.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 125.7 KiB | [postgresql-17-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-17-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-17-mimeo` | `1.5.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 125.6 KiB | [postgresql-17-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-17-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-17-mimeo` | `1.5.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 125.6 KiB | [postgresql-17-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-17-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-17-mimeo` | `1.5.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 108.1 KiB | [postgresql-17-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-17-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-17-mimeo` | `1.5.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 108.1 KiB | [postgresql-17-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-17-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-17-mimeo` | `1.5.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 107.7 KiB | [postgresql-17-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-17-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |
| `postgresql-17-mimeo` | `1.5.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 107.7 KiB | [postgresql-17-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-17-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |
| `postgresql-17-mimeo` | `1.5.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 107.7 KiB | [postgresql-17-mimeo_1.5.1-20.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-17-mimeo_1.5.1-20.pgdg26.04+1_all.deb) |
| `postgresql-17-mimeo` | `1.5.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 107.7 KiB | [postgresql-17-mimeo_1.5.1-20.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-17-mimeo_1.5.1-20.pgdg26.04+1_all.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mimeo_16` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 139.7 KiB | [mimeo_16-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/mimeo_16-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `mimeo_16` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 139.7 KiB | [mimeo_16-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/mimeo_16-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `mimeo_16` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 113.0 KiB | [mimeo_16-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/mimeo_16-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `mimeo_16` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 113.0 KiB | [mimeo_16-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/mimeo_16-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `mimeo_16` | `1.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 114.5 KiB | [mimeo_16-1.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/mimeo_16-1.5.1-1PIGSTY.el10.x86_64.rpm) |
| `mimeo_16` | `1.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 114.5 KiB | [mimeo_16-1.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/mimeo_16-1.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-mimeo` | `1.5.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 125.7 KiB | [postgresql-16-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-16-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-16-mimeo` | `1.5.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 125.7 KiB | [postgresql-16-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-16-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-16-mimeo` | `1.5.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 125.7 KiB | [postgresql-16-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-16-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-16-mimeo` | `1.5.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 125.7 KiB | [postgresql-16-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-16-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-16-mimeo` | `1.5.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 108.0 KiB | [postgresql-16-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-16-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-16-mimeo` | `1.5.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 108.0 KiB | [postgresql-16-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-16-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-16-mimeo` | `1.5.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 107.7 KiB | [postgresql-16-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-16-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |
| `postgresql-16-mimeo` | `1.5.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 107.7 KiB | [postgresql-16-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-16-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |
| `postgresql-16-mimeo` | `1.5.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 107.6 KiB | [postgresql-16-mimeo_1.5.1-20.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-16-mimeo_1.5.1-20.pgdg26.04+1_all.deb) |
| `postgresql-16-mimeo` | `1.5.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 107.6 KiB | [postgresql-16-mimeo_1.5.1-20.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-16-mimeo_1.5.1-20.pgdg26.04+1_all.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mimeo_15` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 139.7 KiB | [mimeo_15-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/mimeo_15-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `mimeo_15` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 139.7 KiB | [mimeo_15-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/mimeo_15-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `mimeo_15` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 113.0 KiB | [mimeo_15-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/mimeo_15-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `mimeo_15` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 113.0 KiB | [mimeo_15-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/mimeo_15-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `mimeo_15` | `1.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 114.5 KiB | [mimeo_15-1.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/mimeo_15-1.5.1-1PIGSTY.el10.x86_64.rpm) |
| `mimeo_15` | `1.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 114.5 KiB | [mimeo_15-1.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/mimeo_15-1.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-mimeo` | `1.5.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 125.7 KiB | [postgresql-15-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-15-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-15-mimeo` | `1.5.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 125.7 KiB | [postgresql-15-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-15-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-15-mimeo` | `1.5.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 125.6 KiB | [postgresql-15-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-15-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-15-mimeo` | `1.5.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 125.6 KiB | [postgresql-15-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-15-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-15-mimeo` | `1.5.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 108.1 KiB | [postgresql-15-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-15-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-15-mimeo` | `1.5.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 108.1 KiB | [postgresql-15-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-15-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-15-mimeo` | `1.5.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 107.7 KiB | [postgresql-15-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-15-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |
| `postgresql-15-mimeo` | `1.5.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 107.7 KiB | [postgresql-15-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-15-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |
| `postgresql-15-mimeo` | `1.5.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 107.7 KiB | [postgresql-15-mimeo_1.5.1-20.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-15-mimeo_1.5.1-20.pgdg26.04+1_all.deb) |
| `postgresql-15-mimeo` | `1.5.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 107.7 KiB | [postgresql-15-mimeo_1.5.1-20.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-15-mimeo_1.5.1-20.pgdg26.04+1_all.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mimeo_14` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 139.7 KiB | [mimeo_14-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/mimeo_14-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `mimeo_14` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 139.7 KiB | [mimeo_14-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/mimeo_14-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `mimeo_14` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 113.0 KiB | [mimeo_14-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/mimeo_14-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `mimeo_14` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 113.0 KiB | [mimeo_14-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/mimeo_14-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `mimeo_14` | `1.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 114.5 KiB | [mimeo_14-1.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/mimeo_14-1.5.1-1PIGSTY.el10.x86_64.rpm) |
| `mimeo_14` | `1.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 114.5 KiB | [mimeo_14-1.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/mimeo_14-1.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-mimeo` | `1.5.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 125.7 KiB | [postgresql-14-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-14-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-14-mimeo` | `1.5.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 125.7 KiB | [postgresql-14-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-14-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-14-mimeo` | `1.5.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 125.7 KiB | [postgresql-14-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-14-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-14-mimeo` | `1.5.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 125.7 KiB | [postgresql-14-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-14-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-14-mimeo` | `1.5.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 108.1 KiB | [postgresql-14-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-14-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-14-mimeo` | `1.5.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 108.1 KiB | [postgresql-14-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-14-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-14-mimeo` | `1.5.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 107.6 KiB | [postgresql-14-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-14-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |
| `postgresql-14-mimeo` | `1.5.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 107.6 KiB | [postgresql-14-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-14-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |
| `postgresql-14-mimeo` | `1.5.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 107.7 KiB | [postgresql-14-mimeo_1.5.1-20.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-14-mimeo_1.5.1-20.pgdg26.04+1_all.deb) |
| `postgresql-14-mimeo` | `1.5.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 107.7 KiB | [postgresql-14-mimeo_1.5.1-20.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-14-mimeo_1.5.1-20.pgdg26.04+1_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/omniti-labs/mimeo" title="Repository" icon="github" subtitle="github.com/omniti-labs/mimeo" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="mimeo-1.5.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg mimeo;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install mimeo;		# install via package name, for the active PG version

pig install mimeo -v 18;   # install for PG 18
pig install mimeo -v 17;   # install for PG 17
pig install mimeo -v 16;   # install for PG 16
pig install mimeo -v 15;   # install for PG 15
pig install mimeo -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION mimeo CASCADE; -- requires dblink
```



## Usage

> [mimeo: Extension for specialized, per-table replication between PostgreSQL instances](https://github.com/omniti-labs/mimeo)

Provides per-table replication between PostgreSQL instances with snapshot (full copy), incremental (timestamp/id based), and DML (insert/update/delete) modes.

### Enabling

```sql
CREATE SCHEMA mimeo;
CREATE EXTENSION mimeo SCHEMA mimeo;
```

Requires the `dblink` extension. Optionally install `pg_jobmon` for monitoring.

### Setting Up a Data Source

```sql
-- Create a dblink connection to the source database
SELECT mimeo.dblink_mapping_create(
    p_mapping_name := 'source_db',
    p_data_source := 'host=sourcehost dbname=sourcedb user=replicator password=secret',
    p_superuser := true
);
```

### Snapshot Replication (Full Table Copy)

Copies the entire source table each time it runs:

```sql
SELECT mimeo.snapshot_maker(
    p_src_table := 'public.my_table',
    p_dblink_id := 1  -- from dblink_mapping
);

-- Refresh the snapshot
SELECT mimeo.refresh_snap('public.my_table');
```

### Incremental Replication (Timestamp-Based)

Replicates rows based on an incrementing timestamp column:

```sql
SELECT mimeo.inserter_maker(
    p_src_table := 'public.events',
    p_control := 'created_at',  -- timestamp column
    p_dblink_id := 1
);

-- Refresh incrementally
SELECT mimeo.refresh_inserter('public.events');
```

For tables with updates (not just inserts):

```sql
SELECT mimeo.updater_maker(
    p_src_table := 'public.orders',
    p_control := 'updated_at',
    p_dblink_id := 1
);

SELECT mimeo.refresh_updater('public.orders');
```

### DML Replication (Insert/Update/Delete)

Full DML tracking via triggers on the source:

```sql
SELECT mimeo.dml_maker(
    p_src_table := 'public.accounts',
    p_dblink_id := 1
);

SELECT mimeo.refresh_dml('public.accounts');
```

### Scheduling Refreshes

Use `pg_jobmon` or cron to schedule periodic calls to the appropriate `refresh_*` function.

### Key Features

- Three replication modes: snapshot, incremental, DML
- Per-table replication (no need to replicate entire database)
- Works between different PostgreSQL versions
- Built on `dblink` for cross-database communication
