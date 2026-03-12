---
title: "redis_fdw"
linkTitle: "redis_fdw"
description: "Foreign data wrapper for querying a Redis server"
weight: 8710
categories: ["FDW"]
width: full
---

[**redis_fdw**](https://github.com/pg-redis-fdw/redis_fdw) : Foreign data wrapper for querying a Redis server


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8710** | {{< badge content="redis_fdw" link="https://github.com/pg-redis-fdw/redis_fdw" >}} | {{< ext "redis_fdw" >}} | `1.0` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "mongo_fdw" >}} {{< ext "redis" >}} {{< ext "kafka_fdw" >}} {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "spat" >}} {{< ext "pgmemcache" >}} {{< ext "odbc_fdw" >}} |

> [!Note] multiple branch for different pg major versions


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `redis_fdw` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "redis_fdw_18" "green" >}} {{< bg "17" "redis_fdw_17" "green" >}} {{< bg "16" "redis_fdw_16" "green" >}} {{< bg "15" "redis_fdw_15" "green" >}} {{< bg "14" "redis_fdw_14" "green" >}} | `redis_fdw_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-redis-fdw" "green" >}} {{< bg "17" "postgresql-17-redis-fdw" "green" >}} {{< bg "16" "postgresql-16-redis-fdw" "green" >}} {{< bg "15" "postgresql-15-redis-fdw" "green" >}} {{< bg "14" "postgresql-14-redis-fdw" "green" >}} | `postgresql-$v-redis-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-redis-fdw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-redis-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-redis-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-redis-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-redis-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-redis-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-redis-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-redis-fdw : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `redis_fdw_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.6 KiB | [redis_fdw_18-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/redis_fdw_18-1.0-2PIGSTY.el8.x86_64.rpm) |
| `redis_fdw_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.9 KiB | [redis_fdw_18-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/redis_fdw_18-1.0-2PIGSTY.el8.aarch64.rpm) |
| `redis_fdw_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.4 KiB | [redis_fdw_18-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/redis_fdw_18-1.0-2PIGSTY.el9.x86_64.rpm) |
| `redis_fdw_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.9 KiB | [redis_fdw_18-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/redis_fdw_18-1.0-2PIGSTY.el9.aarch64.rpm) |
| `redis_fdw_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.7 KiB | [redis_fdw_18-1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/redis_fdw_18-1.0-2PIGSTY.el10.x86_64.rpm) |
| `redis_fdw_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.1 KiB | [redis_fdw_18-1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/redis_fdw_18-1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-redis-fdw` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.0 KiB | [postgresql-18-redis-fdw_1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-18-redis-fdw_1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-redis-fdw` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 50.1 KiB | [postgresql-18-redis-fdw_1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-18-redis-fdw_1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-redis-fdw` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 51.9 KiB | [postgresql-18-redis-fdw_1.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/redis-fdw/postgresql-18-redis-fdw_1.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-redis-fdw` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 50.5 KiB | [postgresql-18-redis-fdw_1.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/redis-fdw/postgresql-18-redis-fdw_1.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-redis-fdw` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 55.7 KiB | [postgresql-18-redis-fdw_1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-18-redis-fdw_1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-redis-fdw` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 54.6 KiB | [postgresql-18-redis-fdw_1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-18-redis-fdw_1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-redis-fdw` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.1 KiB | [postgresql-18-redis-fdw_1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-18-redis-fdw_1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-redis-fdw` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.0 KiB | [postgresql-18-redis-fdw_1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-18-redis-fdw_1.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `redis_fdw_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.3 KiB | [redis_fdw_17-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/redis_fdw_17-1.0-2PIGSTY.el8.x86_64.rpm) |
| `redis_fdw_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 26.6 KiB | [redis_fdw_17-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/redis_fdw_17-1.0-2PIGSTY.el8.aarch64.rpm) |
| `redis_fdw_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.4 KiB | [redis_fdw_17-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/redis_fdw_17-1.0-2PIGSTY.el9.x86_64.rpm) |
| `redis_fdw_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.7 KiB | [redis_fdw_17-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/redis_fdw_17-1.0-2PIGSTY.el9.aarch64.rpm) |
| `redis_fdw_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.7 KiB | [redis_fdw_17-1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/redis_fdw_17-1.0-2PIGSTY.el10.x86_64.rpm) |
| `redis_fdw_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.1 KiB | [redis_fdw_17-1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/redis_fdw_17-1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-redis-fdw` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 51.6 KiB | [postgresql-17-redis-fdw_1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-redis-fdw` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 50.1 KiB | [postgresql-17-redis-fdw_1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-redis-fdw` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.2 KiB | [postgresql-17-redis-fdw_1.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-redis-fdw` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 50.4 KiB | [postgresql-17-redis-fdw_1.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-redis-fdw` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 62.2 KiB | [postgresql-17-redis-fdw_1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-redis-fdw` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 60.9 KiB | [postgresql-17-redis-fdw_1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-redis-fdw` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.2 KiB | [postgresql-17-redis-fdw_1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-redis-fdw` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.1 KiB | [postgresql-17-redis-fdw_1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-17-redis-fdw_1.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `redis_fdw_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.3 KiB | [redis_fdw_16-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/redis_fdw_16-1.0-2PIGSTY.el8.x86_64.rpm) |
| `redis_fdw_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 26.7 KiB | [redis_fdw_16-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/redis_fdw_16-1.0-2PIGSTY.el8.aarch64.rpm) |
| `redis_fdw_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.4 KiB | [redis_fdw_16-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/redis_fdw_16-1.0-2PIGSTY.el9.x86_64.rpm) |
| `redis_fdw_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.8 KiB | [redis_fdw_16-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/redis_fdw_16-1.0-2PIGSTY.el9.aarch64.rpm) |
| `redis_fdw_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.7 KiB | [redis_fdw_16-1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/redis_fdw_16-1.0-2PIGSTY.el10.x86_64.rpm) |
| `redis_fdw_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.2 KiB | [redis_fdw_16-1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/redis_fdw_16-1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-redis-fdw` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 51.7 KiB | [postgresql-16-redis-fdw_1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-redis-fdw` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 50.4 KiB | [postgresql-16-redis-fdw_1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-redis-fdw` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.1 KiB | [postgresql-16-redis-fdw_1.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-redis-fdw` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 50.4 KiB | [postgresql-16-redis-fdw_1.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-redis-fdw` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 61.9 KiB | [postgresql-16-redis-fdw_1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-redis-fdw` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 61.0 KiB | [postgresql-16-redis-fdw_1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-redis-fdw` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.1 KiB | [postgresql-16-redis-fdw_1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-redis-fdw` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.2 KiB | [postgresql-16-redis-fdw_1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-16-redis-fdw_1.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `redis_fdw_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.3 KiB | [redis_fdw_15-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/redis_fdw_15-1.0-2PIGSTY.el8.x86_64.rpm) |
| `redis_fdw_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 26.6 KiB | [redis_fdw_15-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/redis_fdw_15-1.0-2PIGSTY.el8.aarch64.rpm) |
| `redis_fdw_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.4 KiB | [redis_fdw_15-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/redis_fdw_15-1.0-2PIGSTY.el9.x86_64.rpm) |
| `redis_fdw_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.7 KiB | [redis_fdw_15-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/redis_fdw_15-1.0-2PIGSTY.el9.aarch64.rpm) |
| `redis_fdw_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.7 KiB | [redis_fdw_15-1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/redis_fdw_15-1.0-2PIGSTY.el10.x86_64.rpm) |
| `redis_fdw_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.1 KiB | [redis_fdw_15-1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/redis_fdw_15-1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-redis-fdw` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 51.8 KiB | [postgresql-15-redis-fdw_1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-redis-fdw` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 50.3 KiB | [postgresql-15-redis-fdw_1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-redis-fdw` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 51.9 KiB | [postgresql-15-redis-fdw_1.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-redis-fdw` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 50.2 KiB | [postgresql-15-redis-fdw_1.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-redis-fdw` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 61.7 KiB | [postgresql-15-redis-fdw_1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-redis-fdw` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 60.6 KiB | [postgresql-15-redis-fdw_1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-redis-fdw` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.0 KiB | [postgresql-15-redis-fdw_1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-redis-fdw` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.0 KiB | [postgresql-15-redis-fdw_1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-15-redis-fdw_1.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `redis_fdw_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.3 KiB | [redis_fdw_14-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/redis_fdw_14-1.0-2PIGSTY.el8.x86_64.rpm) |
| `redis_fdw_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 26.6 KiB | [redis_fdw_14-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/redis_fdw_14-1.0-2PIGSTY.el8.aarch64.rpm) |
| `redis_fdw_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.4 KiB | [redis_fdw_14-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/redis_fdw_14-1.0-2PIGSTY.el9.x86_64.rpm) |
| `redis_fdw_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.7 KiB | [redis_fdw_14-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/redis_fdw_14-1.0-2PIGSTY.el9.aarch64.rpm) |
| `redis_fdw_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.6 KiB | [redis_fdw_14-1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/redis_fdw_14-1.0-2PIGSTY.el10.x86_64.rpm) |
| `redis_fdw_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.1 KiB | [redis_fdw_14-1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/redis_fdw_14-1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-redis-fdw` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 51.8 KiB | [postgresql-14-redis-fdw_1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-redis-fdw` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 50.2 KiB | [postgresql-14-redis-fdw_1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-redis-fdw` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.0 KiB | [postgresql-14-redis-fdw_1.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-redis-fdw` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 50.2 KiB | [postgresql-14-redis-fdw_1.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-redis-fdw` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 61.7 KiB | [postgresql-14-redis-fdw_1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-redis-fdw` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 60.6 KiB | [postgresql-14-redis-fdw_1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-redis-fdw` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.0 KiB | [postgresql-14-redis-fdw_1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-redis-fdw` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.0 KiB | [postgresql-14-redis-fdw_1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-14-redis-fdw_1.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pg-redis-fdw/redis_fdw" title="Repository" icon="github" subtitle="github.com/pg-redis-fdw/redis_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="redis_fdw-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg redis_fdw;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install redis_fdw;		# install via package name, for the active PG version

pig install redis_fdw -v 18;   # install for PG 18
pig install redis_fdw -v 17;   # install for PG 17
pig install redis_fdw -v 16;   # install for PG 16
pig install redis_fdw -v 15;   # install for PG 15
pig install redis_fdw -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION redis_fdw;
```




## Usage

> [redis_fdw: Foreign data wrapper for querying a Redis server](https://github.com/pg-redis-fdw/redis_fdw)

### Create Server

```sql
CREATE EXTENSION redis_fdw;

CREATE SERVER redis_server FOREIGN DATA WRAPPER redis_fdw
  OPTIONS (address '127.0.0.1', port '6379');
```

**Server Options:** `address` (default `127.0.0.1`), `port` (default `6379`).

### Create User Mapping

```sql
CREATE USER MAPPING FOR pguser SERVER redis_server
  OPTIONS (password 'secret');
```

### Scalar Key-Value Pairs

```sql
CREATE FOREIGN TABLE redis_db0 (
  key text,
  val text
)
SERVER redis_server
OPTIONS (database '0');

SELECT * FROM redis_db0;
```

### Hash Tables (with Key Prefix)

```sql
CREATE FOREIGN TABLE redis_hash (
  key text,
  val text[]
)
SERVER redis_server
OPTIONS (database '0', tabletype 'hash', tablekeyprefix 'mytable:');

INSERT INTO redis_hash VALUES ('mytable:r1', '{prop1,val1,prop2,val2}');
UPDATE redis_hash SET val = '{prop3,val3}' WHERE key = 'mytable:r1';
DELETE FROM redis_hash WHERE key = 'mytable:r1';
SELECT * FROM redis_hash;
```

### Hash Tables (Singleton Key)

```sql
CREATE FOREIGN TABLE redis_singleton (
  key text,
  val text
)
SERVER redis_server
OPTIONS (database '0', tabletype 'hash', singleton_key 'myhash');

INSERT INTO redis_singleton VALUES ('field1', 'value1');
UPDATE redis_singleton SET val = 'newvalue' WHERE key = 'field1';
DELETE FROM redis_singleton WHERE key = 'field1';
```

### Table Options

| Option | Description |
|--------|-------------|
| `database` | Redis database number (default `0`) |
| `tabletype` | `hash`, `list`, `set`, or `zset` (omit for scalar key-value) |
| `tablekeyprefix` | Filter items by key prefix |
| `tablekeyset` | Fetch keys from a specific Redis set |
| `singleton_key` | Access all values from a single Redis key |

Use only one of `tablekeyset` or `tablekeyprefix`. Do not combine them with `singleton_key`.

Hash values are returned as alternating key-value pairs in a `text[]` array. Lists, sets, and sorted sets also return values as arrays.
