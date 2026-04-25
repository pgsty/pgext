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
| **1020** | {{< badge content="timeseries" link="https://github.com/ChuckHend/pg_timeseries" >}} | {{< ext "timeseries" "pg_timeseries" >}} | `0.2.0` | {{< category "TIME" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pg_cron" >}} {{< ext "pg_partman" >}} |
|   **See Also**    | {{< ext "timescaledb" >}} {{< ext "timescaledb_toolkit" >}} {{< ext "periods" >}} {{< ext "temporal_tables" >}} {{< ext "emaj" >}} {{< ext "table_version" >}} {{< ext "pg_task" >}} {{< ext "pg_later" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_timeseries` | `pg_cron`, `pg_partman` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "pg_timeseries_18" "green" >}} {{< bg "17" "pg_timeseries_17" "green" >}} {{< bg "16" "pg_timeseries_16" "green" >}} {{< bg "15" "pg_timeseries_15" "green" >}} {{< bg "14" "pg_timeseries_14" "green" >}} | `pg_timeseries_$v` | `pg_cron_$v`, `pg_partman_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "postgresql-18-pg-timeseries" "green" >}} {{< bg "17" "postgresql-17-pg-timeseries" "green" >}} {{< bg "16" "postgresql-16-pg-timeseries" "green" >}} {{< bg "15" "postgresql-15-pg-timeseries" "green" >}} {{< bg "14" "postgresql-14-pg-timeseries" "green" >}} | `postgresql-$v-pg-timeseries` | `postgresql-$v-cron`, `postgresql-$v-partman` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_timeseries_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-timeseries : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-timeseries : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-timeseries : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-timeseries : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-timeseries : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-timeseries : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-timeseries : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-timeseries : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-timeseries : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-timeseries : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-timeseries : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-timeseries : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-timeseries : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-timeseries : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-timeseries : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-timeseries : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-timeseries : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-timeseries : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-timeseries : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_timeseries_18` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 28.3 KiB | [pg_timeseries_18-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_18-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_18` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.2 KiB | [pg_timeseries_18-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_18-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_18` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.9 KiB | [pg_timeseries_18-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_18-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_18` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.8 KiB | [pg_timeseries_18-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_18-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_18` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.1 KiB | [pg_timeseries_18-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_timeseries_18-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_timeseries_18` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.1 KiB | [pg_timeseries_18-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_timeseries_18-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-timeseries` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.6 KiB | [postgresql-18-pg-timeseries_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-18-pg-timeseries_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-timeseries` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.6 KiB | [postgresql-18-pg-timeseries_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-18-pg-timeseries_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-timeseries` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.6 KiB | [postgresql-18-pg-timeseries_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-18-pg-timeseries_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-timeseries` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.6 KiB | [postgresql-18-pg-timeseries_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-18-pg-timeseries_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-timeseries` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.0 KiB | [postgresql-18-pg-timeseries_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-18-pg-timeseries_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-timeseries` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.0 KiB | [postgresql-18-pg-timeseries_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-18-pg-timeseries_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-timeseries` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.0 KiB | [postgresql-18-pg-timeseries_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-18-pg-timeseries_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-timeseries` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 24.0 KiB | [postgresql-18-pg-timeseries_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-18-pg-timeseries_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_timeseries_17` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 28.3 KiB | [pg_timeseries_17-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_17-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_17` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.2 KiB | [pg_timeseries_17-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_17-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_17` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.9 KiB | [pg_timeseries_17-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_17-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_17` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.8 KiB | [pg_timeseries_17-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_17-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_17` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.1 KiB | [pg_timeseries_17-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_timeseries_17-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_timeseries_17` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.1 KiB | [pg_timeseries_17-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_timeseries_17-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-timeseries` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.6 KiB | [postgresql-17-pg-timeseries_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-timeseries` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.6 KiB | [postgresql-17-pg-timeseries_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-timeseries` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.6 KiB | [postgresql-17-pg-timeseries_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-timeseries` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.6 KiB | [postgresql-17-pg-timeseries_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-timeseries` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.0 KiB | [postgresql-17-pg-timeseries_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-timeseries` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.0 KiB | [postgresql-17-pg-timeseries_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-timeseries` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.0 KiB | [postgresql-17-pg-timeseries_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-timeseries` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 24.0 KiB | [postgresql-17-pg-timeseries_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-17-pg-timeseries_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_timeseries_16` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 28.3 KiB | [pg_timeseries_16-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_16-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_16` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.2 KiB | [pg_timeseries_16-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_16-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_16` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.9 KiB | [pg_timeseries_16-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_16-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_16` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.8 KiB | [pg_timeseries_16-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_16-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_16` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.1 KiB | [pg_timeseries_16-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_timeseries_16-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_timeseries_16` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.1 KiB | [pg_timeseries_16-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_timeseries_16-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-timeseries` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.6 KiB | [postgresql-16-pg-timeseries_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-timeseries` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.6 KiB | [postgresql-16-pg-timeseries_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-timeseries` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.6 KiB | [postgresql-16-pg-timeseries_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-timeseries` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.6 KiB | [postgresql-16-pg-timeseries_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-timeseries` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.0 KiB | [postgresql-16-pg-timeseries_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-timeseries` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.0 KiB | [postgresql-16-pg-timeseries_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-timeseries` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.0 KiB | [postgresql-16-pg-timeseries_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-timeseries` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 24.0 KiB | [postgresql-16-pg-timeseries_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-16-pg-timeseries_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_timeseries_15` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 28.3 KiB | [pg_timeseries_15-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_15-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_15` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.2 KiB | [pg_timeseries_15-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_15-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_15` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.9 KiB | [pg_timeseries_15-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_15-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_15` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.8 KiB | [pg_timeseries_15-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_15-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_15` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.1 KiB | [pg_timeseries_15-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_timeseries_15-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_timeseries_15` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.1 KiB | [pg_timeseries_15-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_timeseries_15-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-timeseries` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.6 KiB | [postgresql-15-pg-timeseries_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-timeseries` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.6 KiB | [postgresql-15-pg-timeseries_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-timeseries` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.6 KiB | [postgresql-15-pg-timeseries_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-timeseries` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.6 KiB | [postgresql-15-pg-timeseries_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-timeseries` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.0 KiB | [postgresql-15-pg-timeseries_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-timeseries` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.0 KiB | [postgresql-15-pg-timeseries_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-timeseries` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.0 KiB | [postgresql-15-pg-timeseries_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-timeseries` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 24.0 KiB | [postgresql-15-pg-timeseries_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-15-pg-timeseries_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_timeseries_14` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 28.3 KiB | [pg_timeseries_14-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_timeseries_14-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_timeseries_14` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.2 KiB | [pg_timeseries_14-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_timeseries_14-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_timeseries_14` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.9 KiB | [pg_timeseries_14-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_timeseries_14-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_timeseries_14` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.8 KiB | [pg_timeseries_14-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_timeseries_14-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_timeseries_14` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.1 KiB | [pg_timeseries_14-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_timeseries_14-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_timeseries_14` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.1 KiB | [pg_timeseries_14-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_timeseries_14-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-timeseries` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.6 KiB | [postgresql-14-pg-timeseries_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-timeseries` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.6 KiB | [postgresql-14-pg-timeseries_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-timeseries` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.6 KiB | [postgresql-14-pg-timeseries_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-timeseries` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.6 KiB | [postgresql-14-pg-timeseries_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-timeseries` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.0 KiB | [postgresql-14-pg-timeseries_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-timeseries` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.0 KiB | [postgresql-14-pg-timeseries_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-timeseries` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.0 KiB | [postgresql-14-pg-timeseries_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-timeseries` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 24.0 KiB | [postgresql-14-pg-timeseries_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-timeseries/postgresql-14-pg-timeseries_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ChuckHend/pg_timeseries" title="Repository" icon="github" subtitle="github.com/ChuckHend/pg_timeseries" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_timeseries-0.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_timeseries;		# build rpm/deb
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

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION timeseries CASCADE; -- requires pg_cron, pg_partman
```



## Usage

> [pg_timeseries: Convenience API for time series stack](https://github.com/ChuckHend/pg_timeseries)

This extension provides a cohesive user experience around the creation, maintenance, and use of time-series tables.

### Getting Started

Assuming you already have a partitioned table created, simply call the `enable_ts_table` function with your table name.

```sql
CREATE EXTENSION timeseries CASCADE;

SELECT enable_ts_table('sensor_readings');
```

With this one call, several things will happen:

  * The table will be restructured as a series of partitions using PostgreSQL's [native PARTITION features](https://www.postgresql.org/docs/current/ddl-partitioning.html)
  * Each partition covers a particular range of time (one week by default)
  * New partitions will be created for some time in the future (one month by default)
  * Once an hour, a maintenance job will create any missing partitions as well as needed future ones


## Using Your Tables

### Indexes

The time-series tables you create start out life as little more than typical [partitioned PostgreSQL tables](https://www.postgresql.org/docs/current/ddl-partitioning.html). All of PostgreSQL's existing functionality will "just work" with them.

[Traditional B-Tree indexes](https://www.postgresql.org/docs/current/btree-intro.html) work well for time-series data, but you may wish to benchmark [BRIN indexes](https://www.postgresql.org/docs/current/brin-intro.html) as well, as they may perform better in specific query scenarios (often queries with _many_ results). Start with B-Tree if you don't anticipate more than a million records in each partition (by default, partitions are one week long).

### Partition Sizing

Because calculating the total size of partitioned tables can be tedious, this extension provides several easy-to-use views surfacing this information.

To examine the table (data), index, and total size for each of your partitions, query the time-series partition information view, `ts_part_info`. A general rule of thumb is that each partition should be able to fit within roughly one quarter of your available memory.

### Retention

Call `set_ts_retention_policy` with your time-series table and an interval (say, `'90 days'`) to establish a retention policy. Once an hour, any partitions falling entirely outside the retention window will be dropped. Use `clear_ts_retention_policy` to revert to the default behavior (infinite retention). Each of these functions will return the previous retention policy when called.

### Compression

By calling `set_ts_compression_policy` on a time-series table with an appropriate interval (perhaps `'1 month'`), this extension will compress partitions (using a columnar storage method) older than the specified interval, once an hour. A function is also provided for clearing any existing policy (existing partitions will not be decompressed, however).

The compression features depend on the [citus](https://github.com/citusdata/citus/tree/main) and [citus_columnar](https://github.com/citusdata/citus/tree/main/src/backend/columnar) extensions:

```sql
CREATE EXTENSION citus;
CREATE EXTENSION citus_columnar;
```


## Analytics Helpers

### `first` and `last`

These two functions help clean up the syntax of a fairly common pattern: a query is grouped by one dimension, but a user wants to know what the first or last row in a group is when ordered by a _different_ dimension.

```sql
SELECT machine_id,
       last(cpu_util, recorded_at)
FROM events
GROUP BY machine_id;
```

### `date_bin_table`

This function automates the tedium of aligning time-series values to a given width, or "stride", and makes sure to include NULL rows for any time periods where the source table has no data points.

```sql
SELECT * FROM date_bin_table(NULL::target_table, '1 hour', '[2024-02-01 00:00, 2024-02-02 15:00]');
```

The output of this query will differ from simply hitting the target table directly in three ways:

  * Rows will be sorted by time, ascending
  * The time column's values will be binned to the provided width
  * Extra rows will be added for periods with no data. They will include the time stamp for that bin and NULL in all other columns

### `make_view_incremental`

This function accepts a view and converts it into a materialized view which is kept up-to-date after every modification. This removes the need for users to pick between always up-to-date `VIEW`s and having to call `REFRESH` on `MATERIALIZED VIEW`s.

The underlying functionality is provided by [a fork](https://github.com/ChuckHend/pg_ivm) of [`pg_ivm`](https://github.com/sraoss/pg_ivm). Enable the `pg_ivm` extension if you want to use this feature:

```sql
CREATE EXTENSION pg_ivm;
```


## Requirements

* [pg_cron](https://github.com/citusdata/pg_cron)
* [pg_partman](https://github.com/pgpartman/pg_partman)

### Optional Dependencies

* [pg_ivm](https://github.com/sraoss/pg_ivm) — for incremental materialized views
* [Citus & Citus Columnar](https://github.com/citusdata/citus) — for compression features
