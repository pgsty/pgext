---
title: "mimeo"
linkTitle: "mimeo"
description: "Extension for specialized, per-table replication between PostgreSQL instances"
weight: 9700
categories: ["ETL"]
width: full
---

Extension for specialized, per-table replication between PostgreSQL instances


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9700** | {{< badge content="mimeo" link="https://github.com/omniti-labs/mimeo" >}} | {{< ext "mimeo" >}} | `1.5.1` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "dblink" >}} |
|   **See Also**    | {{< ext "pg_jobmon" >}} {{< ext "postgres_fdw" >}} {{< ext "pglogical" >}} {{< ext "pg_cron" >}} {{< ext "pg_partman" >}} {{< ext "repmgr" >}} {{< ext "pg_fact_loader" >}} {{< ext "pg_failover_slots" >}} |

> [!Note] name conflict with pg_partman


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/mimeo" >}} | `1.5.1` | {{< bg "18" "mimeo_18" "green" >}} {{< bg "17" "mimeo_17" "green" >}} {{< bg "16" "mimeo_16" "green" >}} {{< bg "15" "mimeo_15" "green" >}} {{< bg "14" "mimeo_14" "green" >}} {{< bg "13" "mimeo_13" "green" >}} | `mimeo_$v` | - |
| **Debian** | {{< badge content="PGDG" link="/e/mimeo" >}} | `1.5.1` | {{< bg "18" "postgresql-18-mimeo" "green" >}} {{< bg "17" "postgresql-17-mimeo" "green" >}} {{< bg "16" "postgresql-16-mimeo" "green" >}} {{< bg "15" "postgresql-15-mimeo" "green" >}} {{< bg "14" "postgresql-14-mimeo" "green" >}} {{< bg "13" "postgresql-13-mimeo" "green" >}} | `postgresql-$v-mimeo` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "mimeo_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-13-mimeo : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-13-mimeo : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-13-mimeo : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-13-mimeo : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-13-mimeo : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-13-mimeo : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-13-mimeo : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.5.1" "postgresql-18-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-17-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-16-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-15-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-14-mimeo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.1" "postgresql-13-mimeo : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mimeo_18` | 1.5.1 | `el8.x86_64` | pigsty | 139.7 KiB | [mimeo_18-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/mimeo_18-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `mimeo_18` | 1.5.1 | `el8.aarch64` | pigsty | 139.7 KiB | [mimeo_18-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/mimeo_18-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `mimeo_18` | 1.5.1 | `el9.x86_64` | pigsty | 113.0 KiB | [mimeo_18-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/mimeo_18-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `mimeo_18` | 1.5.1 | `el9.aarch64` | pigsty | 113.0 KiB | [mimeo_18-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/mimeo_18-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `mimeo_18` | 1.5.1 | `el10.x86_64` | pigsty | 114.5 KiB | [mimeo_18-1.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/mimeo_18-1.5.1-1PIGSTY.el10.x86_64.rpm) |
| `mimeo_18` | 1.5.1 | `el10.aarch64` | pigsty | 114.5 KiB | [mimeo_18-1.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/mimeo_18-1.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-mimeo` | 1.5.1 | `d12.x86_64` | pgdg | 125.6 KiB | [postgresql-18-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-18-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-18-mimeo` | 1.5.1 | `d12.aarch64` | pgdg | 125.6 KiB | [postgresql-18-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-18-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-18-mimeo` | 1.5.1 | `d13.x86_64` | pgdg | 125.6 KiB | [postgresql-18-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-18-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-18-mimeo` | 1.5.1 | `d13.aarch64` | pgdg | 125.6 KiB | [postgresql-18-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-18-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-18-mimeo` | 1.5.1 | `u22.x86_64` | pgdg | 108.0 KiB | [postgresql-18-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-18-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-18-mimeo` | 1.5.1 | `u22.aarch64` | pgdg | 108.0 KiB | [postgresql-18-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-18-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-18-mimeo` | 1.5.1 | `u24.x86_64` | pgdg | 107.7 KiB | [postgresql-18-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-18-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |
| `postgresql-18-mimeo` | 1.5.1 | `u24.aarch64` | pgdg | 107.7 KiB | [postgresql-18-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-18-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mimeo_17` | 1.5.1 | `el8.x86_64` | pigsty | 139.7 KiB | [mimeo_17-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/mimeo_17-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `mimeo_17` | 1.5.1 | `el8.aarch64` | pigsty | 139.7 KiB | [mimeo_17-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/mimeo_17-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `mimeo_17` | 1.5.1 | `el9.x86_64` | pigsty | 113.0 KiB | [mimeo_17-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/mimeo_17-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `mimeo_17` | 1.5.1 | `el9.aarch64` | pigsty | 113.0 KiB | [mimeo_17-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/mimeo_17-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `mimeo_17` | 1.5.1 | `el10.x86_64` | pigsty | 114.5 KiB | [mimeo_17-1.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/mimeo_17-1.5.1-1PIGSTY.el10.x86_64.rpm) |
| `mimeo_17` | 1.5.1 | `el10.aarch64` | pigsty | 114.5 KiB | [mimeo_17-1.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/mimeo_17-1.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-mimeo` | 1.5.1 | `d12.x86_64` | pgdg | 125.7 KiB | [postgresql-17-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-17-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-17-mimeo` | 1.5.1 | `d12.aarch64` | pgdg | 125.7 KiB | [postgresql-17-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-17-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-17-mimeo` | 1.5.1 | `d13.x86_64` | pgdg | 125.6 KiB | [postgresql-17-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-17-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-17-mimeo` | 1.5.1 | `d13.aarch64` | pgdg | 125.6 KiB | [postgresql-17-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-17-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-17-mimeo` | 1.5.1 | `u22.x86_64` | pgdg | 108.1 KiB | [postgresql-17-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-17-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-17-mimeo` | 1.5.1 | `u22.aarch64` | pgdg | 108.1 KiB | [postgresql-17-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-17-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-17-mimeo` | 1.5.1 | `u24.x86_64` | pgdg | 107.7 KiB | [postgresql-17-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-17-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |
| `postgresql-17-mimeo` | 1.5.1 | `u24.aarch64` | pgdg | 107.7 KiB | [postgresql-17-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-17-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mimeo_16` | 1.5.1 | `el8.x86_64` | pigsty | 139.7 KiB | [mimeo_16-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/mimeo_16-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `mimeo_16` | 1.5.1 | `el8.aarch64` | pigsty | 139.7 KiB | [mimeo_16-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/mimeo_16-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `mimeo_16` | 1.5.1 | `el9.x86_64` | pigsty | 113.0 KiB | [mimeo_16-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/mimeo_16-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `mimeo_16` | 1.5.1 | `el9.aarch64` | pigsty | 113.0 KiB | [mimeo_16-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/mimeo_16-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `mimeo_16` | 1.5.1 | `el10.x86_64` | pigsty | 114.5 KiB | [mimeo_16-1.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/mimeo_16-1.5.1-1PIGSTY.el10.x86_64.rpm) |
| `mimeo_16` | 1.5.1 | `el10.aarch64` | pigsty | 114.5 KiB | [mimeo_16-1.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/mimeo_16-1.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-mimeo` | 1.5.1 | `d12.x86_64` | pgdg | 125.7 KiB | [postgresql-16-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-16-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-16-mimeo` | 1.5.1 | `d12.aarch64` | pgdg | 125.7 KiB | [postgresql-16-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-16-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-16-mimeo` | 1.5.1 | `d13.x86_64` | pgdg | 125.7 KiB | [postgresql-16-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-16-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-16-mimeo` | 1.5.1 | `d13.aarch64` | pgdg | 125.7 KiB | [postgresql-16-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-16-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-16-mimeo` | 1.5.1 | `u22.x86_64` | pgdg | 108.0 KiB | [postgresql-16-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-16-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-16-mimeo` | 1.5.1 | `u22.aarch64` | pgdg | 108.0 KiB | [postgresql-16-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-16-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-16-mimeo` | 1.5.1 | `u24.x86_64` | pgdg | 107.7 KiB | [postgresql-16-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-16-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |
| `postgresql-16-mimeo` | 1.5.1 | `u24.aarch64` | pgdg | 107.7 KiB | [postgresql-16-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-16-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mimeo_15` | 1.5.1 | `el8.x86_64` | pigsty | 139.7 KiB | [mimeo_15-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/mimeo_15-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `mimeo_15` | 1.5.1 | `el8.aarch64` | pigsty | 139.7 KiB | [mimeo_15-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/mimeo_15-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `mimeo_15` | 1.5.1 | `el9.x86_64` | pigsty | 113.0 KiB | [mimeo_15-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/mimeo_15-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `mimeo_15` | 1.5.1 | `el9.aarch64` | pigsty | 113.0 KiB | [mimeo_15-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/mimeo_15-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `mimeo_15` | 1.5.1 | `el10.x86_64` | pigsty | 114.5 KiB | [mimeo_15-1.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/mimeo_15-1.5.1-1PIGSTY.el10.x86_64.rpm) |
| `mimeo_15` | 1.5.1 | `el10.aarch64` | pigsty | 114.5 KiB | [mimeo_15-1.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/mimeo_15-1.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-mimeo` | 1.5.1 | `d12.x86_64` | pgdg | 125.7 KiB | [postgresql-15-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-15-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-15-mimeo` | 1.5.1 | `d12.aarch64` | pgdg | 125.7 KiB | [postgresql-15-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-15-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-15-mimeo` | 1.5.1 | `d13.x86_64` | pgdg | 125.6 KiB | [postgresql-15-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-15-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-15-mimeo` | 1.5.1 | `d13.aarch64` | pgdg | 125.6 KiB | [postgresql-15-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-15-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-15-mimeo` | 1.5.1 | `u22.x86_64` | pgdg | 108.1 KiB | [postgresql-15-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-15-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-15-mimeo` | 1.5.1 | `u22.aarch64` | pgdg | 108.1 KiB | [postgresql-15-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-15-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-15-mimeo` | 1.5.1 | `u24.x86_64` | pgdg | 107.7 KiB | [postgresql-15-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-15-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |
| `postgresql-15-mimeo` | 1.5.1 | `u24.aarch64` | pgdg | 107.7 KiB | [postgresql-15-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-15-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mimeo_14` | 1.5.1 | `el8.x86_64` | pigsty | 139.7 KiB | [mimeo_14-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/mimeo_14-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `mimeo_14` | 1.5.1 | `el8.aarch64` | pigsty | 139.7 KiB | [mimeo_14-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/mimeo_14-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `mimeo_14` | 1.5.1 | `el9.x86_64` | pigsty | 113.0 KiB | [mimeo_14-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/mimeo_14-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `mimeo_14` | 1.5.1 | `el9.aarch64` | pigsty | 113.0 KiB | [mimeo_14-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/mimeo_14-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `mimeo_14` | 1.5.1 | `el10.x86_64` | pigsty | 114.5 KiB | [mimeo_14-1.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/mimeo_14-1.5.1-1PIGSTY.el10.x86_64.rpm) |
| `mimeo_14` | 1.5.1 | `el10.aarch64` | pigsty | 114.5 KiB | [mimeo_14-1.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/mimeo_14-1.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-mimeo` | 1.5.1 | `d12.x86_64` | pgdg | 125.7 KiB | [postgresql-14-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-14-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-14-mimeo` | 1.5.1 | `d12.aarch64` | pgdg | 125.7 KiB | [postgresql-14-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-14-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-14-mimeo` | 1.5.1 | `d13.x86_64` | pgdg | 125.7 KiB | [postgresql-14-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-14-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-14-mimeo` | 1.5.1 | `d13.aarch64` | pgdg | 125.7 KiB | [postgresql-14-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-14-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-14-mimeo` | 1.5.1 | `u22.x86_64` | pgdg | 108.1 KiB | [postgresql-14-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-14-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-14-mimeo` | 1.5.1 | `u22.aarch64` | pgdg | 108.1 KiB | [postgresql-14-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-14-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-14-mimeo` | 1.5.1 | `u24.x86_64` | pgdg | 107.6 KiB | [postgresql-14-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-14-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |
| `postgresql-14-mimeo` | 1.5.1 | `u24.aarch64` | pgdg | 107.6 KiB | [postgresql-14-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-14-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mimeo_13` | 1.5.1 | `el8.x86_64` | pigsty | 139.7 KiB | [mimeo_13-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/mimeo_13-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `mimeo_13` | 1.5.1 | `el8.aarch64` | pigsty | 139.7 KiB | [mimeo_13-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/mimeo_13-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `mimeo_13` | 1.5.1 | `el9.x86_64` | pigsty | 113.1 KiB | [mimeo_13-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/mimeo_13-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `mimeo_13` | 1.5.1 | `el9.aarch64` | pigsty | 113.0 KiB | [mimeo_13-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/mimeo_13-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `mimeo_13` | 1.5.1 | `el10.x86_64` | pigsty | 114.5 KiB | [mimeo_13-1.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/mimeo_13-1.5.1-1PIGSTY.el10.x86_64.rpm) |
| `mimeo_13` | 1.5.1 | `el10.aarch64` | pigsty | 114.5 KiB | [mimeo_13-1.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/mimeo_13-1.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-mimeo` | 1.5.1 | `d12.x86_64` | pgdg | 125.7 KiB | [postgresql-13-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-13-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-13-mimeo` | 1.5.1 | `d12.aarch64` | pgdg | 125.7 KiB | [postgresql-13-mimeo_1.5.1-20.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-13-mimeo_1.5.1-20.pgdg12+1_all.deb) |
| `postgresql-13-mimeo` | 1.5.1 | `d13.x86_64` | pgdg | 125.6 KiB | [postgresql-13-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-13-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-13-mimeo` | 1.5.1 | `d13.aarch64` | pgdg | 125.6 KiB | [postgresql-13-mimeo_1.5.1-20.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-13-mimeo_1.5.1-20.pgdg13+1_all.deb) |
| `postgresql-13-mimeo` | 1.5.1 | `u22.x86_64` | pgdg | 108.1 KiB | [postgresql-13-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-13-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-13-mimeo` | 1.5.1 | `u22.aarch64` | pgdg | 108.1 KiB | [postgresql-13-mimeo_1.5.1-20.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-13-mimeo_1.5.1-20.pgdg22.04+1_all.deb) |
| `postgresql-13-mimeo` | 1.5.1 | `u24.x86_64` | pgdg | 107.6 KiB | [postgresql-13-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-13-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |
| `postgresql-13-mimeo` | 1.5.1 | `u24.aarch64` | pgdg | 107.6 KiB | [postgresql-13-mimeo_1.5.1-20.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mimeo/postgresql-13-mimeo_1.5.1-20.pgdg24.04+1_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/omniti-labs/mimeo" title="Repository" icon="github" subtitle="github.com/omniti-labs/mimeo" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="mimeo-1.5.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get mimeo; # get mimeo source code
pig build dep mimeo; # install build dependencies
pig build pkg mimeo; # build extension rpm or deb
pig build ext mimeo; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install mimeo; # install by extension name, for the current active PG version
pig ext install mimeo; # install via package alias, for the active PG version
pig ext install mimeo -v 18;   # install for PG 18
pig ext install mimeo -v 17;   # install for PG 17
pig ext install mimeo -v 16;   # install for PG 16
pig ext install mimeo -v 15;   # install for PG 15
pig ext install mimeo -v 14;   # install for PG 14
pig ext install mimeo -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION mimeo;
```

