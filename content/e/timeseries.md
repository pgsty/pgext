---
title: "timeseries"
linkTitle: "timeseries"
description: "Convenience API for time series stack"
weight: 1020
categories: ["TIME"]
width: full
---

Convenience API for time series stack


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1020** | {{< badge content="timeseries" link="https://github.com/ChuckHend/pg_timeseries" >}} | {{< ext "timeseries" "pg_timeseries" >}} | `0.1.7` | {{< category "TIME" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "columnar" >}} {{< ext "pg_cron" >}} {{< ext "pg_ivm" >}} {{< ext "pg_partman" >}} |
|   **See Also**    | {{< ext "timescaledb" >}} {{< ext "timescaledb_toolkit" >}} {{< ext "periods" >}} {{< ext "temporal_tables" >}} {{< ext "emaj" >}} {{< ext "table_version" >}} {{< ext "pg_task" >}} {{< ext "pg_later" >}} |

> [!Note] unmet deps: hydra17 not ready, pg_partman17/pg_ivm12 on el not ready, 0.1.7 break, pg18 break


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/timeseries" >}} | `0.1.7` | {{< bg "18" "pg_timeseries_18" "green" >}} {{< bg "17" "pg_timeseries_17" "green" >}} {{< bg "16" "pg_timeseries_16" "green" >}} {{< bg "15" "pg_timeseries_15" "green" >}} {{< bg "14" "pg_timeseries_14" "green" >}} {{< bg "13" "pg_timeseries_13" "green" >}} | `pg_timeseries_$v` | `hydra_$v`, `pg_cron_$v`, `pg_ivm_$v`, `pg_partman_$v` |
| **Debian** | {{< badge content="PIGSTY" link="/e/timeseries" >}} | `0.1.7` | {{< bg "18" "postgresql-18-pg-timeseries" "green" >}} {{< bg "17" "postgresql-17-pg-timeseries" "green" >}} {{< bg "16" "postgresql-16-pg-timeseries" "green" >}} {{< bg "15" "postgresql-15-pg-timeseries" "green" >}} {{< bg "14" "postgresql-14-pg-timeseries" "green" >}} {{< bg "13" "postgresql-13-pg-timeseries" "green" >}} | `postgresql-$v-pg-timeseries` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_18 : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_17 : HIDE 2" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_16 : HIDE 2" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_15 : HIDE 2" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_14 : HIDE 2" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_13 : HIDE 2" >}}   |
|    `el8.aarch64`    |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_18 : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_17 : HIDE 2" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_16 : HIDE 2" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_15 : HIDE 2" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_14 : HIDE 2" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_13 : HIDE 2" >}}   |
|    `el9.x86_64`    |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_18 : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_17 : HIDE 2" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_16 : HIDE 2" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_15 : HIDE 2" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_14 : HIDE 2" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_13 : HIDE 2" >}}   |
|    `el9.aarch64`    |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_18 : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_17 : HIDE 2" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_16 : HIDE 2" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_15 : HIDE 2" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_14 : HIDE 2" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_13 : HIDE 2" >}}   |
|    `el10.x86_64`    |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_18 : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_17 : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_16 : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_15 : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_14 : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_13 : HIDE 1" >}}   |
|    `el10.aarch64`    |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_18 : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_17 : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_16 : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_15 : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_14 : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.7" "pg_timeseries_13 : HIDE 1" >}}   |
|    `d12.x86_64`    |  {{< bg "MISS" "postgresql-18-pg-timeseries : HIDE 0" "red" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-17-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-16-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-15-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-14-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-13-pg-timeseries : HIDE 1" >}}   |
|    `d12.aarch64`    |  {{< bg "MISS" "postgresql-18-pg-timeseries : HIDE 0" "red" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-17-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-16-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-15-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-14-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-13-pg-timeseries : HIDE 1" >}}   |
|    `d13.x86_64`    |  {{< bg "MISS" "postgresql-18-pg-timeseries : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-17-pg-timeseries : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-16-pg-timeseries : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-15-pg-timeseries : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-14-pg-timeseries : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-13-pg-timeseries : HIDE 0" "red" >}}   |
|    `d13.aarch64`    |  {{< bg "MISS" "postgresql-18-pg-timeseries : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-17-pg-timeseries : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-16-pg-timeseries : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-15-pg-timeseries : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-14-pg-timeseries : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-13-pg-timeseries : HIDE 0" "red" >}}   |
|    `u22.x86_64`    |  {{< bg "MISS" "postgresql-18-pg-timeseries : HIDE 0" "red" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-17-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-16-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-15-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-14-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-13-pg-timeseries : HIDE 1" >}}   |
|    `u22.aarch64`    |  {{< bg "MISS" "postgresql-18-pg-timeseries : HIDE 0" "red" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-17-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-16-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-15-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-14-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-13-pg-timeseries : HIDE 1" >}}   |
|    `u24.x86_64`    |  {{< bg "MISS" "postgresql-18-pg-timeseries : HIDE 0" "red" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-17-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-16-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-15-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-14-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-13-pg-timeseries : HIDE 1" >}}   |
|    `u24.aarch64`    |  {{< bg "MISS" "postgresql-18-pg-timeseries : HIDE 0" "red" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-17-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-16-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-15-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-14-pg-timeseries : HIDE 1" >}}   |  {{< bg "PIGSTY 0.1.6" "postgresql-13-pg-timeseries : HIDE 1" >}}   |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_timeseries_18` | 0.1.7 | `el8.x86_64` | pigsty | 27.8 KiB | [pg_timeseries_18-0.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_18-0.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_18` | 0.1.7 | `el8.aarch64` | pigsty | 27.8 KiB | [pg_timeseries_18-0.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_18-0.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_18` | 0.1.7 | `el9.x86_64` | pigsty | 27.4 KiB | [pg_timeseries_18-0.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_18-0.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_18` | 0.1.7 | `el9.aarch64` | pigsty | 27.4 KiB | [pg_timeseries_18-0.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_18-0.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_18` | 0.1.7 | `el10.x86_64` | pigsty | 27.6 KiB | [pg_timeseries_18-0.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_timeseries_18-0.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_timeseries_18` | 0.1.7 | `el10.aarch64` | pigsty | 27.6 KiB | [pg_timeseries_18-0.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_timeseries_18-0.1.7-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_timeseries_17` | 0.1.7 | `el8.x86_64` | pigsty | 27.8 KiB | [pg_timeseries_17-0.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_17-0.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_17` | 0.1.6 | `el8.x86_64` | pigsty | 26.2 KiB | [pg_timeseries_17-0.1.6-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_17-0.1.6-2PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_17` | 0.1.7 | `el8.aarch64` | pigsty | 27.8 KiB | [pg_timeseries_17-0.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_17-0.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_17` | 0.1.6 | `el8.aarch64` | pigsty | 26.1 KiB | [pg_timeseries_17-0.1.6-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_17-0.1.6-2PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_17` | 0.1.7 | `el9.x86_64` | pigsty | 27.4 KiB | [pg_timeseries_17-0.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_17-0.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_17` | 0.1.6 | `el9.x86_64` | pigsty | 25.8 KiB | [pg_timeseries_17-0.1.6-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_17-0.1.6-2PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_17` | 0.1.7 | `el9.aarch64` | pigsty | 27.4 KiB | [pg_timeseries_17-0.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_17-0.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_17` | 0.1.6 | `el9.aarch64` | pigsty | 25.8 KiB | [pg_timeseries_17-0.1.6-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_17-0.1.6-2PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_17` | 0.1.7 | `el10.x86_64` | pigsty | 27.6 KiB | [pg_timeseries_17-0.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_timeseries_17-0.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_timeseries_17` | 0.1.7 | `el10.aarch64` | pigsty | 27.6 KiB | [pg_timeseries_17-0.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_timeseries_17-0.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-timeseries` | 0.1.6 | `d12.x86_64` | pigsty | 22.4 KiB | [postgresql-17-pg-timeseries_0.1.6-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.1.6-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-timeseries` | 0.1.6 | `d12.aarch64` | pigsty | 22.4 KiB | [postgresql-17-pg-timeseries_0.1.6-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.1.6-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-timeseries` | 0.1.6 | `u22.x86_64` | pigsty | 22.8 KiB | [postgresql-17-pg-timeseries_0.1.6-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.1.6-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-timeseries` | 0.1.6 | `u22.aarch64` | pigsty | 22.8 KiB | [postgresql-17-pg-timeseries_0.1.6-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.1.6-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-timeseries` | 0.1.6 | `u24.x86_64` | pigsty | 22.8 KiB | [postgresql-17-pg-timeseries_0.1.6-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.1.6-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-timeseries` | 0.1.6 | `u24.aarch64` | pigsty | 22.8 KiB | [postgresql-17-pg-timeseries_0.1.6-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.1.6-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_timeseries_16` | 0.1.7 | `el8.x86_64` | pigsty | 27.8 KiB | [pg_timeseries_16-0.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_16-0.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_16` | 0.1.6 | `el8.x86_64` | pigsty | 26.2 KiB | [pg_timeseries_16-0.1.6-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_16-0.1.6-2PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_16` | 0.1.7 | `el8.aarch64` | pigsty | 27.8 KiB | [pg_timeseries_16-0.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_16-0.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_16` | 0.1.6 | `el8.aarch64` | pigsty | 26.1 KiB | [pg_timeseries_16-0.1.6-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_16-0.1.6-2PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_16` | 0.1.7 | `el9.x86_64` | pigsty | 27.4 KiB | [pg_timeseries_16-0.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_16-0.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_16` | 0.1.6 | `el9.x86_64` | pigsty | 25.8 KiB | [pg_timeseries_16-0.1.6-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_16-0.1.6-2PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_16` | 0.1.7 | `el9.aarch64` | pigsty | 27.4 KiB | [pg_timeseries_16-0.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_16-0.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_16` | 0.1.6 | `el9.aarch64` | pigsty | 25.8 KiB | [pg_timeseries_16-0.1.6-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_16-0.1.6-2PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_16` | 0.1.7 | `el10.x86_64` | pigsty | 27.6 KiB | [pg_timeseries_16-0.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_timeseries_16-0.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_timeseries_16` | 0.1.7 | `el10.aarch64` | pigsty | 27.6 KiB | [pg_timeseries_16-0.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_timeseries_16-0.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-timeseries` | 0.1.6 | `d12.x86_64` | pigsty | 22.4 KiB | [postgresql-16-pg-timeseries_0.1.6-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.1.6-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-timeseries` | 0.1.6 | `d12.aarch64` | pigsty | 22.4 KiB | [postgresql-16-pg-timeseries_0.1.6-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.1.6-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-timeseries` | 0.1.6 | `u22.x86_64` | pigsty | 22.8 KiB | [postgresql-16-pg-timeseries_0.1.6-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.1.6-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-timeseries` | 0.1.6 | `u22.aarch64` | pigsty | 22.8 KiB | [postgresql-16-pg-timeseries_0.1.6-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.1.6-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-timeseries` | 0.1.6 | `u24.x86_64` | pigsty | 22.8 KiB | [postgresql-16-pg-timeseries_0.1.6-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.1.6-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-timeseries` | 0.1.6 | `u24.aarch64` | pigsty | 22.8 KiB | [postgresql-16-pg-timeseries_0.1.6-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.1.6-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_timeseries_15` | 0.1.7 | `el8.x86_64` | pigsty | 27.8 KiB | [pg_timeseries_15-0.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_15-0.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_15` | 0.1.6 | `el8.x86_64` | pigsty | 26.2 KiB | [pg_timeseries_15-0.1.6-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_15-0.1.6-2PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_15` | 0.1.7 | `el8.aarch64` | pigsty | 27.8 KiB | [pg_timeseries_15-0.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_15-0.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_15` | 0.1.6 | `el8.aarch64` | pigsty | 26.1 KiB | [pg_timeseries_15-0.1.6-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_15-0.1.6-2PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_15` | 0.1.7 | `el9.x86_64` | pigsty | 27.4 KiB | [pg_timeseries_15-0.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_15-0.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_15` | 0.1.6 | `el9.x86_64` | pigsty | 25.8 KiB | [pg_timeseries_15-0.1.6-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_15-0.1.6-2PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_15` | 0.1.7 | `el9.aarch64` | pigsty | 27.4 KiB | [pg_timeseries_15-0.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_15-0.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_15` | 0.1.6 | `el9.aarch64` | pigsty | 25.8 KiB | [pg_timeseries_15-0.1.6-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_15-0.1.6-2PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_15` | 0.1.7 | `el10.x86_64` | pigsty | 27.6 KiB | [pg_timeseries_15-0.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_timeseries_15-0.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_timeseries_15` | 0.1.7 | `el10.aarch64` | pigsty | 27.6 KiB | [pg_timeseries_15-0.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_timeseries_15-0.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-timeseries` | 0.1.6 | `d12.x86_64` | pigsty | 22.4 KiB | [postgresql-15-pg-timeseries_0.1.6-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.1.6-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-timeseries` | 0.1.6 | `d12.aarch64` | pigsty | 22.4 KiB | [postgresql-15-pg-timeseries_0.1.6-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.1.6-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-timeseries` | 0.1.6 | `u22.x86_64` | pigsty | 22.8 KiB | [postgresql-15-pg-timeseries_0.1.6-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.1.6-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-timeseries` | 0.1.6 | `u22.aarch64` | pigsty | 22.8 KiB | [postgresql-15-pg-timeseries_0.1.6-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.1.6-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-timeseries` | 0.1.6 | `u24.x86_64` | pigsty | 22.8 KiB | [postgresql-15-pg-timeseries_0.1.6-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.1.6-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-timeseries` | 0.1.6 | `u24.aarch64` | pigsty | 22.8 KiB | [postgresql-15-pg-timeseries_0.1.6-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.1.6-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_timeseries_14` | 0.1.7 | `el8.x86_64` | pigsty | 27.8 KiB | [pg_timeseries_14-0.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_14-0.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_14` | 0.1.6 | `el8.x86_64` | pigsty | 26.2 KiB | [pg_timeseries_14-0.1.6-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_14-0.1.6-2PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_14` | 0.1.7 | `el8.aarch64` | pigsty | 27.8 KiB | [pg_timeseries_14-0.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_14-0.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_14` | 0.1.6 | `el8.aarch64` | pigsty | 26.1 KiB | [pg_timeseries_14-0.1.6-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_14-0.1.6-2PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_14` | 0.1.7 | `el9.x86_64` | pigsty | 27.4 KiB | [pg_timeseries_14-0.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_14-0.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_14` | 0.1.6 | `el9.x86_64` | pigsty | 25.8 KiB | [pg_timeseries_14-0.1.6-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_14-0.1.6-2PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_14` | 0.1.7 | `el9.aarch64` | pigsty | 27.4 KiB | [pg_timeseries_14-0.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_14-0.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_14` | 0.1.6 | `el9.aarch64` | pigsty | 25.8 KiB | [pg_timeseries_14-0.1.6-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_14-0.1.6-2PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_14` | 0.1.7 | `el10.x86_64` | pigsty | 27.6 KiB | [pg_timeseries_14-0.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_timeseries_14-0.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_timeseries_14` | 0.1.7 | `el10.aarch64` | pigsty | 27.6 KiB | [pg_timeseries_14-0.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_timeseries_14-0.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-timeseries` | 0.1.6 | `d12.x86_64` | pigsty | 22.4 KiB | [postgresql-14-pg-timeseries_0.1.6-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.1.6-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-timeseries` | 0.1.6 | `d12.aarch64` | pigsty | 22.4 KiB | [postgresql-14-pg-timeseries_0.1.6-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.1.6-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-timeseries` | 0.1.6 | `u22.x86_64` | pigsty | 22.8 KiB | [postgresql-14-pg-timeseries_0.1.6-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.1.6-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-timeseries` | 0.1.6 | `u22.aarch64` | pigsty | 22.8 KiB | [postgresql-14-pg-timeseries_0.1.6-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.1.6-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-timeseries` | 0.1.6 | `u24.x86_64` | pigsty | 22.8 KiB | [postgresql-14-pg-timeseries_0.1.6-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.1.6-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-timeseries` | 0.1.6 | `u24.aarch64` | pigsty | 22.8 KiB | [postgresql-14-pg-timeseries_0.1.6-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.1.6-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_timeseries_13` | 0.1.7 | `el8.x86_64` | pigsty | 27.8 KiB | [pg_timeseries_13-0.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_13-0.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_13` | 0.1.6 | `el8.x86_64` | pigsty | 26.2 KiB | [pg_timeseries_13-0.1.6-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_13-0.1.6-2PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_13` | 0.1.7 | `el8.aarch64` | pigsty | 27.8 KiB | [pg_timeseries_13-0.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_13-0.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_13` | 0.1.6 | `el8.aarch64` | pigsty | 26.1 KiB | [pg_timeseries_13-0.1.6-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_13-0.1.6-2PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_13` | 0.1.7 | `el9.x86_64` | pigsty | 27.4 KiB | [pg_timeseries_13-0.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_13-0.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_13` | 0.1.6 | `el9.x86_64` | pigsty | 25.8 KiB | [pg_timeseries_13-0.1.6-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_13-0.1.6-2PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_13` | 0.1.7 | `el9.aarch64` | pigsty | 27.4 KiB | [pg_timeseries_13-0.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_13-0.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_13` | 0.1.6 | `el9.aarch64` | pigsty | 25.8 KiB | [pg_timeseries_13-0.1.6-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_13-0.1.6-2PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_13` | 0.1.7 | `el10.x86_64` | pigsty | 27.6 KiB | [pg_timeseries_13-0.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_timeseries_13-0.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_timeseries_13` | 0.1.7 | `el10.aarch64` | pigsty | 27.6 KiB | [pg_timeseries_13-0.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_timeseries_13-0.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-timeseries` | 0.1.6 | `d12.x86_64` | pigsty | 22.4 KiB | [postgresql-13-pg-timeseries_0.1.6-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-13-pg-timeseries_0.1.6-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-timeseries` | 0.1.6 | `d12.aarch64` | pigsty | 22.4 KiB | [postgresql-13-pg-timeseries_0.1.6-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-13-pg-timeseries_0.1.6-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-timeseries` | 0.1.6 | `u22.x86_64` | pigsty | 22.8 KiB | [postgresql-13-pg-timeseries_0.1.6-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-13-pg-timeseries_0.1.6-2PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-timeseries` | 0.1.6 | `u22.aarch64` | pigsty | 22.8 KiB | [postgresql-13-pg-timeseries_0.1.6-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-13-pg-timeseries_0.1.6-2PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-timeseries` | 0.1.6 | `u24.x86_64` | pigsty | 22.8 KiB | [postgresql-13-pg-timeseries_0.1.6-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-13-pg-timeseries_0.1.6-2PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-timeseries` | 0.1.6 | `u24.aarch64` | pigsty | 22.8 KiB | [postgresql-13-pg-timeseries_0.1.6-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-13-pg-timeseries_0.1.6-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ChuckHend/pg_timeseries" title="Repository" icon="github" subtitle="github.com/ChuckHend/pg_timeseries" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_timeseries-0.1.7.tar.gz" >}}
{{< /cards >}}


```bash
pig build get timeseries; # get timeseries source code
pig build dep timeseries; # install build dependencies
pig build pkg timeseries; # build extension rpm or deb
pig build ext timeseries; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install timeseries; # install by extension name, for the current active PG version
pig ext install pg_timeseries; # install via package alias, for the active PG version
pig ext install timeseries -v 18;   # install for PG 18
pig ext install timeseries -v 17;   # install for PG 17
pig ext install timeseries -v 16;   # install for PG 16
pig ext install timeseries -v 15;   # install for PG 15
pig ext install timeseries -v 14;   # install for PG 14
pig ext install timeseries -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION timeseries;
```

