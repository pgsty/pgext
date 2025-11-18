---
title: "pg_later"
linkTitle: "pg_later"
description: "Run queries now and get results later"
weight: 1090
categories: ["TIME"]
width: full
---

[**pg_later**](https://github.com/ChuckHend/pg_later) : Run queries now and get results later


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1090** | {{< badge content="pg_later" link="https://github.com/ChuckHend/pg_later" >}} | {{< ext "pg_later" >}} | `0.4.0` | {{< category "TIME" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pgmq" >}} |
|   **See Also**    | {{< ext "pg_cron" >}} {{< ext "pg_task" >}} {{< ext "pg_background" >}} {{< ext "timescaledb" >}} {{< ext "timescaledb_toolkit" >}} {{< ext "timeseries" >}} {{< ext "periods" >}} {{< ext "temporal_tables" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_later` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.0` | {{< bg "18" "pg_later_18" "green" >}} {{< bg "17" "pg_later_17" "green" >}} {{< bg "16" "pg_later_16" "green" >}} {{< bg "15" "pg_later_15" "green" >}} {{< bg "14" "pg_later_14" "green" >}} {{< bg "13" "pg_later_13" "green" >}} | `pg_later_$v` | `pgmq_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.0` | {{< bg "18" "postgresql-18-pg-later" "green" >}} {{< bg "17" "postgresql-17-pg-later" "green" >}} {{< bg "16" "postgresql-16-pg-later" "green" >}} {{< bg "15" "postgresql-15-pg-later" "green" >}} {{< bg "14" "postgresql-14-pg-later" "green" >}} {{< bg "13" "postgresql-13-pg-later" "green" >}} | `postgresql-$v-pg-later` | `postgresql-$v-pgmq` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_13 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "pg_later_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-later : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-later : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-later : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-later : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-later : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-later : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-later : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-later : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-later : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_later_18` | `0.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.4 MiB | [pg_later_18-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_later_18-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_later_18` | `0.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.2 MiB | [pg_later_18-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_later_18-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_later_18` | `0.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [pg_later_18-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_later_18-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_later_18` | `0.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.3 MiB | [pg_later_18-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_18-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_later_18` | `0.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.4 MiB | [pg_later_18-0.3.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_later_18-0.3.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_later_18` | `0.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.3 MiB | [pg_later_18-0.3.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_later_18-0.3.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-later` | `0.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.2 MiB | [postgresql-18-pg-later_0.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-18-pg-later_0.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-later` | `0.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 991.3 KiB | [postgresql-18-pg-later_0.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-18-pg-later_0.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-later` | `0.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-18-pg-later_0.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-later/postgresql-18-pg-later_0.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-later` | `0.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 991.1 KiB | [postgresql-18-pg-later_0.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-later/postgresql-18-pg-later_0.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-later` | `0.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.3 MiB | [postgresql-18-pg-later_0.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-18-pg-later_0.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-later` | `0.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-18-pg-later_0.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-18-pg-later_0.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-later` | `0.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.3 MiB | [postgresql-18-pg-later_0.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-18-pg-later_0.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-later` | `0.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.1 MiB | [postgresql-18-pg-later_0.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-18-pg-later_0.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_later_17` | `0.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.4 MiB | [pg_later_17-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_later_17-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_later_17` | `0.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.2 MiB | [pg_later_17-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_later_17-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_later_17` | `0.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [pg_later_17-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_later_17-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_later_17` | `0.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.3 MiB | [pg_later_17-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_17-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_later_17` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.3 MiB | [pg_later_17-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_17-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_later_17` | `0.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.4 MiB | [pg_later_17-0.3.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_later_17-0.3.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_later_17` | `0.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.3 MiB | [pg_later_17-0.3.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_later_17-0.3.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-later` | `0.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.2 MiB | [postgresql-17-pg-later_0.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-17-pg-later_0.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-later` | `0.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 990.9 KiB | [postgresql-17-pg-later_0.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-17-pg-later_0.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-later` | `0.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-17-pg-later_0.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-later/postgresql-17-pg-later_0.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-later` | `0.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 991.4 KiB | [postgresql-17-pg-later_0.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-later/postgresql-17-pg-later_0.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-later` | `0.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.3 MiB | [postgresql-17-pg-later_0.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-17-pg-later_0.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-later` | `0.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-17-pg-later_0.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-17-pg-later_0.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-later` | `0.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.3 MiB | [postgresql-17-pg-later_0.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-17-pg-later_0.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-later` | `0.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.1 MiB | [postgresql-17-pg-later_0.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-17-pg-later_0.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_later_16` | `0.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.4 MiB | [pg_later_16-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_later_16-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_later_16` | `0.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.2 MiB | [pg_later_16-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_later_16-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_later_16` | `0.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [pg_later_16-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_later_16-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_later_16` | `0.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.3 MiB | [pg_later_16-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_16-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_later_16` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.3 MiB | [pg_later_16-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_16-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_later_16` | `0.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.4 MiB | [pg_later_16-0.3.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_later_16-0.3.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_later_16` | `0.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.3 MiB | [pg_later_16-0.3.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_later_16-0.3.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-later` | `0.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.2 MiB | [postgresql-16-pg-later_0.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-16-pg-later_0.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-later` | `0.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 990.6 KiB | [postgresql-16-pg-later_0.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-16-pg-later_0.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-later` | `0.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-16-pg-later_0.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-later/postgresql-16-pg-later_0.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-later` | `0.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 990.7 KiB | [postgresql-16-pg-later_0.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-later/postgresql-16-pg-later_0.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-later` | `0.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.3 MiB | [postgresql-16-pg-later_0.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-16-pg-later_0.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-later` | `0.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-16-pg-later_0.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-16-pg-later_0.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-later` | `0.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.3 MiB | [postgresql-16-pg-later_0.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-16-pg-later_0.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-later` | `0.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.1 MiB | [postgresql-16-pg-later_0.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-16-pg-later_0.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_later_15` | `0.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.4 MiB | [pg_later_15-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_later_15-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_later_15` | `0.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.2 MiB | [pg_later_15-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_later_15-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_later_15` | `0.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [pg_later_15-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_later_15-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_later_15` | `0.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.3 MiB | [pg_later_15-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_15-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_later_15` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.3 MiB | [pg_later_15-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_15-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_later_15` | `0.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.4 MiB | [pg_later_15-0.3.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_later_15-0.3.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_later_15` | `0.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.3 MiB | [pg_later_15-0.3.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_later_15-0.3.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-later` | `0.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.2 MiB | [postgresql-15-pg-later_0.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-15-pg-later_0.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-later` | `0.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 992.4 KiB | [postgresql-15-pg-later_0.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-15-pg-later_0.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-later` | `0.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-15-pg-later_0.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-later/postgresql-15-pg-later_0.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-later` | `0.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 990.8 KiB | [postgresql-15-pg-later_0.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-later/postgresql-15-pg-later_0.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-later` | `0.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.3 MiB | [postgresql-15-pg-later_0.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-15-pg-later_0.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-later` | `0.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-15-pg-later_0.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-15-pg-later_0.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-later` | `0.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.3 MiB | [postgresql-15-pg-later_0.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-15-pg-later_0.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-later` | `0.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.1 MiB | [postgresql-15-pg-later_0.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-15-pg-later_0.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_later_14` | `0.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.4 MiB | [pg_later_14-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_later_14-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_later_14` | `0.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.2 MiB | [pg_later_14-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_later_14-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_later_14` | `0.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [pg_later_14-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_later_14-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_later_14` | `0.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.3 MiB | [pg_later_14-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_14-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_later_14` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.3 MiB | [pg_later_14-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_14-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_later_14` | `0.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.4 MiB | [pg_later_14-0.3.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_later_14-0.3.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_later_14` | `0.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.3 MiB | [pg_later_14-0.3.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_later_14-0.3.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-later` | `0.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.2 MiB | [postgresql-14-pg-later_0.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-14-pg-later_0.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-later` | `0.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 990.9 KiB | [postgresql-14-pg-later_0.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-14-pg-later_0.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-later` | `0.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-14-pg-later_0.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-later/postgresql-14-pg-later_0.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-later` | `0.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 991.7 KiB | [postgresql-14-pg-later_0.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-later/postgresql-14-pg-later_0.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-later` | `0.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.3 MiB | [postgresql-14-pg-later_0.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-14-pg-later_0.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-later` | `0.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-14-pg-later_0.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-14-pg-later_0.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-later` | `0.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.3 MiB | [postgresql-14-pg-later_0.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-14-pg-later_0.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-later` | `0.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.1 MiB | [postgresql-14-pg-later_0.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-14-pg-later_0.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_later_13` | `0.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.4 MiB | [pg_later_13-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_later_13-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_later_13` | `0.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.2 MiB | [pg_later_13-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_later_13-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_later_13` | `0.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [pg_later_13-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_later_13-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_later_13` | `0.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.3 MiB | [pg_later_13-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_13-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_later_13` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.3 MiB | [pg_later_13-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_later_13-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_later_13` | `0.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.4 MiB | [pg_later_13-0.3.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_later_13-0.3.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_later_13` | `0.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.3 MiB | [pg_later_13-0.3.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_later_13-0.3.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-later` | `0.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.2 MiB | [postgresql-13-pg-later_0.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-13-pg-later_0.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-later` | `0.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 991.2 KiB | [postgresql-13-pg-later_0.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-later/postgresql-13-pg-later_0.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-later` | `0.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-13-pg-later_0.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-later/postgresql-13-pg-later_0.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-later` | `0.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 991.1 KiB | [postgresql-13-pg-later_0.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-later/postgresql-13-pg-later_0.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-later` | `0.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.3 MiB | [postgresql-13-pg-later_0.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-13-pg-later_0.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-later` | `0.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-13-pg-later_0.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-later/postgresql-13-pg-later_0.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-later` | `0.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.3 MiB | [postgresql-13-pg-later_0.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-13-pg-later_0.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-later` | `0.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.1 MiB | [postgresql-13-pg-later_0.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-later/postgresql-13-pg-later_0.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ChuckHend/pg_later" title="Repository" icon="github" subtitle="github.com/ChuckHend/pg_later" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_later-0.4.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_later;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_later;		# install via package name, for the active PG version

pig install pg_later -v 18;   # install for PG 18
pig install pg_later -v 17;   # install for PG 17
pig install pg_later -v 16;   # install for PG 16
pig install pg_later -v 15;   # install for PG 15
pig install pg_later -v 14;   # install for PG 14
pig install pg_later -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_later CASCADE; -- requires pgmq
```
