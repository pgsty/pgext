---
title: "timeseries"
linkTitle: "timeseries"
description: "Convenience API for time series stack"
weight: 1020
categories: ["TIME"]
width: full
---

[**pg_timeseries**](https://github.com/ChuckHend/pg_timeseries) : Convenience API for time series stack


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1020** | {{< badge content="timeseries" link="https://github.com/ChuckHend/pg_timeseries" >}} | {{< ext "timeseries" "pg_timeseries" >}} | `0.1.7` | {{< category "TIME" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "columnar" >}} {{< ext "pg_cron" >}} {{< ext "pg_ivm" >}} {{< ext "pg_partman" >}} |
|   **See Also**    | {{< ext "timescaledb" >}} {{< ext "timescaledb_toolkit" >}} {{< ext "periods" >}} {{< ext "temporal_tables" >}} {{< ext "emaj" >}} {{< ext "table_version" >}} {{< ext "pg_task" >}} {{< ext "pg_later" >}} |

> [!Note] unmet deps: hydra17 not ready, pg_partman17/pg_ivm12 on el not ready, 0.1.7 break, pg18 break


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.7` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_timeseries` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.7` | {{< bg "18" "pg_timeseries_18" "green" >}} {{< bg "17" "pg_timeseries_17" "green" >}} {{< bg "16" "pg_timeseries_16" "green" >}} {{< bg "15" "pg_timeseries_15" "green" >}} {{< bg "14" "pg_timeseries_14" "green" >}} {{< bg "13" "pg_timeseries_13" "green" >}} | `pg_timeseries_$v` | `hydra_$v`, `pg_cron_$v`, `pg_ivm_$v`, `pg_partman_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.7` | {{< bg "18" "postgresql-18-pg-timeseries" "green" >}} {{< bg "17" "postgresql-17-pg-timeseries" "green" >}} {{< bg "16" "postgresql-16-pg-timeseries" "green" >}} {{< bg "15" "postgresql-15-pg-timeseries" "green" >}} {{< bg "14" "postgresql-14-pg-timeseries" "green" >}} {{< bg "13" "postgresql-13-pg-timeseries" "green" >}} | `postgresql-$v-pg-timeseries` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_timeseries_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-18-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-17-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-16-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-15-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-14-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-13-pg-timeseries : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-18-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-17-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-16-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-15-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-14-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-13-pg-timeseries : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-18-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-17-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-16-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-15-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-14-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-13-pg-timeseries : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-18-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-17-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-16-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-15-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-14-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-13-pg-timeseries : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-18-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-17-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-16-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-15-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-14-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-13-pg-timeseries : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-18-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-17-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-16-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-15-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-14-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-13-pg-timeseries : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-18-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-17-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-16-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-15-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-14-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-13-pg-timeseries : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-18-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-17-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-16-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-15-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-14-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-13-pg-timeseries : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_timeseries_18` | `0.1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.8 KiB | [pg_timeseries_18-0.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_18-0.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_18` | `0.1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.8 KiB | [pg_timeseries_18-0.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_18-0.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_18` | `0.1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.4 KiB | [pg_timeseries_18-0.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_18-0.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_18` | `0.1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.4 KiB | [pg_timeseries_18-0.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_18-0.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_18` | `0.1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.6 KiB | [pg_timeseries_18-0.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_timeseries_18-0.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_timeseries_18` | `0.1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.6 KiB | [pg_timeseries_18-0.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_timeseries_18-0.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-timeseries` | `0.1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.4 KiB | [postgresql-18-pg-timeseries_0.1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-18-pg-timeseries_0.1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-timeseries` | `0.1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.4 KiB | [postgresql-18-pg-timeseries_0.1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-18-pg-timeseries_0.1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-timeseries` | `0.1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.4 KiB | [postgresql-18-pg-timeseries_0.1.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-18-pg-timeseries_0.1.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-timeseries` | `0.1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.4 KiB | [postgresql-18-pg-timeseries_0.1.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-18-pg-timeseries_0.1.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-timeseries` | `0.1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.8 KiB | [postgresql-18-pg-timeseries_0.1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-18-pg-timeseries_0.1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-timeseries` | `0.1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 23.8 KiB | [postgresql-18-pg-timeseries_0.1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-18-pg-timeseries_0.1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-timeseries` | `0.1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.8 KiB | [postgresql-18-pg-timeseries_0.1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-18-pg-timeseries_0.1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-timeseries` | `0.1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.8 KiB | [postgresql-18-pg-timeseries_0.1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-18-pg-timeseries_0.1.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_timeseries_17` | `0.1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.8 KiB | [pg_timeseries_17-0.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_17-0.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_17` | `0.1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.8 KiB | [pg_timeseries_17-0.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_17-0.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_17` | `0.1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.4 KiB | [pg_timeseries_17-0.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_17-0.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_17` | `0.1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.4 KiB | [pg_timeseries_17-0.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_17-0.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_17` | `0.1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.6 KiB | [pg_timeseries_17-0.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_timeseries_17-0.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_timeseries_17` | `0.1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.6 KiB | [pg_timeseries_17-0.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_timeseries_17-0.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-timeseries` | `0.1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.4 KiB | [postgresql-17-pg-timeseries_0.1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-timeseries` | `0.1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.4 KiB | [postgresql-17-pg-timeseries_0.1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-timeseries` | `0.1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.4 KiB | [postgresql-17-pg-timeseries_0.1.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.1.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-timeseries` | `0.1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.4 KiB | [postgresql-17-pg-timeseries_0.1.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.1.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-timeseries` | `0.1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.8 KiB | [postgresql-17-pg-timeseries_0.1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-timeseries` | `0.1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 23.8 KiB | [postgresql-17-pg-timeseries_0.1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-timeseries` | `0.1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.8 KiB | [postgresql-17-pg-timeseries_0.1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-timeseries` | `0.1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.8 KiB | [postgresql-17-pg-timeseries_0.1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.1.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_timeseries_16` | `0.1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.8 KiB | [pg_timeseries_16-0.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_16-0.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_16` | `0.1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.8 KiB | [pg_timeseries_16-0.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_16-0.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_16` | `0.1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.4 KiB | [pg_timeseries_16-0.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_16-0.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_16` | `0.1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.4 KiB | [pg_timeseries_16-0.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_16-0.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_16` | `0.1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.6 KiB | [pg_timeseries_16-0.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_timeseries_16-0.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_timeseries_16` | `0.1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.6 KiB | [pg_timeseries_16-0.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_timeseries_16-0.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-timeseries` | `0.1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.4 KiB | [postgresql-16-pg-timeseries_0.1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-timeseries` | `0.1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.4 KiB | [postgresql-16-pg-timeseries_0.1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-timeseries` | `0.1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.4 KiB | [postgresql-16-pg-timeseries_0.1.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.1.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-timeseries` | `0.1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.4 KiB | [postgresql-16-pg-timeseries_0.1.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.1.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-timeseries` | `0.1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.8 KiB | [postgresql-16-pg-timeseries_0.1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-timeseries` | `0.1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 23.8 KiB | [postgresql-16-pg-timeseries_0.1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-timeseries` | `0.1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.8 KiB | [postgresql-16-pg-timeseries_0.1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-timeseries` | `0.1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.8 KiB | [postgresql-16-pg-timeseries_0.1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.1.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_timeseries_15` | `0.1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.8 KiB | [pg_timeseries_15-0.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_15-0.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_15` | `0.1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.8 KiB | [pg_timeseries_15-0.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_15-0.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_15` | `0.1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.4 KiB | [pg_timeseries_15-0.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_15-0.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_15` | `0.1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.4 KiB | [pg_timeseries_15-0.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_15-0.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_15` | `0.1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.6 KiB | [pg_timeseries_15-0.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_timeseries_15-0.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_timeseries_15` | `0.1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.6 KiB | [pg_timeseries_15-0.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_timeseries_15-0.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-timeseries` | `0.1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.4 KiB | [postgresql-15-pg-timeseries_0.1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-timeseries` | `0.1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.4 KiB | [postgresql-15-pg-timeseries_0.1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-timeseries` | `0.1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.4 KiB | [postgresql-15-pg-timeseries_0.1.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.1.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-timeseries` | `0.1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.4 KiB | [postgresql-15-pg-timeseries_0.1.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.1.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-timeseries` | `0.1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.8 KiB | [postgresql-15-pg-timeseries_0.1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-timeseries` | `0.1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 23.8 KiB | [postgresql-15-pg-timeseries_0.1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-timeseries` | `0.1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.8 KiB | [postgresql-15-pg-timeseries_0.1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-timeseries` | `0.1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.8 KiB | [postgresql-15-pg-timeseries_0.1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.1.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_timeseries_14` | `0.1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.8 KiB | [pg_timeseries_14-0.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_14-0.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_14` | `0.1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.8 KiB | [pg_timeseries_14-0.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_14-0.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_14` | `0.1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.4 KiB | [pg_timeseries_14-0.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_14-0.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_14` | `0.1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.4 KiB | [pg_timeseries_14-0.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_14-0.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_14` | `0.1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.6 KiB | [pg_timeseries_14-0.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_timeseries_14-0.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_timeseries_14` | `0.1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.6 KiB | [pg_timeseries_14-0.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_timeseries_14-0.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-timeseries` | `0.1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.4 KiB | [postgresql-14-pg-timeseries_0.1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-timeseries` | `0.1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.4 KiB | [postgresql-14-pg-timeseries_0.1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-timeseries` | `0.1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.4 KiB | [postgresql-14-pg-timeseries_0.1.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.1.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-timeseries` | `0.1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.4 KiB | [postgresql-14-pg-timeseries_0.1.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.1.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-timeseries` | `0.1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.8 KiB | [postgresql-14-pg-timeseries_0.1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-timeseries` | `0.1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 23.8 KiB | [postgresql-14-pg-timeseries_0.1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-timeseries` | `0.1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.8 KiB | [postgresql-14-pg-timeseries_0.1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-timeseries` | `0.1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.8 KiB | [postgresql-14-pg-timeseries_0.1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.1.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_timeseries_13` | `0.1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.8 KiB | [pg_timeseries_13-0.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_13-0.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_13` | `0.1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.8 KiB | [pg_timeseries_13-0.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_13-0.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_13` | `0.1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.4 KiB | [pg_timeseries_13-0.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_13-0.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_13` | `0.1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.4 KiB | [pg_timeseries_13-0.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_13-0.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_13` | `0.1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.6 KiB | [pg_timeseries_13-0.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_timeseries_13-0.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_timeseries_13` | `0.1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.6 KiB | [pg_timeseries_13-0.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_timeseries_13-0.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-timeseries` | `0.1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.4 KiB | [postgresql-13-pg-timeseries_0.1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-13-pg-timeseries_0.1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-timeseries` | `0.1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.4 KiB | [postgresql-13-pg-timeseries_0.1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-13-pg-timeseries_0.1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-timeseries` | `0.1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.4 KiB | [postgresql-13-pg-timeseries_0.1.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-13-pg-timeseries_0.1.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-timeseries` | `0.1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.4 KiB | [postgresql-13-pg-timeseries_0.1.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-13-pg-timeseries_0.1.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-timeseries` | `0.1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.8 KiB | [postgresql-13-pg-timeseries_0.1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-13-pg-timeseries_0.1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-timeseries` | `0.1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 23.8 KiB | [postgresql-13-pg-timeseries_0.1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-13-pg-timeseries_0.1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-timeseries` | `0.1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.8 KiB | [postgresql-13-pg-timeseries_0.1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-13-pg-timeseries_0.1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-timeseries` | `0.1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.8 KiB | [postgresql-13-pg-timeseries_0.1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-13-pg-timeseries_0.1.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ChuckHend/pg_timeseries" title="Repository" icon="github" subtitle="github.com/ChuckHend/pg_timeseries" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_timeseries-0.1.7.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_timeseries;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_timeseries;		# install via package name, for the active PG version
pig install timeseries;		# install by extension name, for the current active PG version

pig install timeseries -v 18;   # install for PG 18
pig install timeseries -v 17;   # install for PG 17
pig install timeseries -v 16;   # install for PG 16
pig install timeseries -v 15;   # install for PG 15
pig install timeseries -v 14;   # install for PG 14
pig install timeseries -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION timeseries CASCADE; -- requires columnar, pg_cron, pg_ivm, pg_partman
```
