---
title: "redis_fdw"
linkTitle: "redis_fdw"
description: "Foreign data wrapper for querying a Redis server"
weight: 8710
categories: ["FDW"]
width: full
---

[**redis_fdw**](https://github.com/pg-redis-fdw/redis_fdw)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8710** | {{< badge content="redis_fdw" link="https://github.com/pg-redis-fdw/redis_fdw" >}} | {{< ext "redis_fdw" >}} | `1.0` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "mongo_fdw" >}} {{< ext "redis" >}} {{< ext "kafka_fdw" >}} {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "spat" >}} {{< ext "pgmemcache" >}} {{< ext "odbc_fdw" >}} |

> [!Note] multiple branch for different pg major versions


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/redis_fdw" >}} | `1.0` | {{< bg "18" "redis_fdw_18*" "green" >}} {{< bg "17" "redis_fdw_17*" "green" >}} {{< bg "16" "redis_fdw_16*" "green" >}} {{< bg "15" "redis_fdw_15*" "green" >}} {{< bg "14" "redis_fdw_14*" "green" >}} {{< bg "13" "redis_fdw_13*" "green" >}} | `redis_fdw_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/redis_fdw" >}} | `1.0` | {{< bg "18" "postgresql-18-redis-fdw" "green" >}} {{< bg "17" "postgresql-17-redis-fdw" "green" >}} {{< bg "16" "postgresql-16-redis-fdw" "green" >}} {{< bg "15" "postgresql-15-redis-fdw" "green" >}} {{< bg "14" "postgresql-14-redis-fdw" "green" >}} {{< bg "13" "postgresql-13-redis-fdw" "green" >}} | `postgresql-$v-redis-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_14 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.1" "redis_fdw_13 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_14 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.1" "redis_fdw_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "redis_fdw_14 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.1" "redis_fdw_13 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-redis-fdw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-redis-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-redis-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-redis-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-redis-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-redis-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-redis-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-redis-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-redis-fdw : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
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

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `redis_fdw_13` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 198.5 KiB | [redis_fdw_13-1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/redis_fdw_13-1.1-1.rhel8.x86_64.rpm) |
| `redis_fdw_13` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.0 KiB | [redis_fdw_13-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/redis_fdw_13-1.0-2PIGSTY.el8.x86_64.rpm) |
| `redis_fdw_13` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 26.5 KiB | [redis_fdw_13-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/redis_fdw_13-1.0-2PIGSTY.el8.aarch64.rpm) |
| `redis_fdw_13` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.4 KiB | [redis_fdw_13-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/redis_fdw_13-1.0-2PIGSTY.el9.x86_64.rpm) |
| `redis_fdw_13` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.8 KiB | [redis_fdw_13-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/redis_fdw_13-1.0-2PIGSTY.el9.aarch64.rpm) |
| `redis_fdw_13` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.0 KiB | [redis_fdw_13-1.1-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/redis_fdw_13-1.1-5PGDG.rhel10.x86_64.rpm) |
| `redis_fdw_13` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.7 KiB | [redis_fdw_13-1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/redis_fdw_13-1.0-2PIGSTY.el10.x86_64.rpm) |
| `redis_fdw_13` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.4 KiB | [redis_fdw_13-1.1-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/redis_fdw_13-1.1-5PGDG.rhel10.aarch64.rpm) |
| `redis_fdw_13` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.3 KiB | [redis_fdw_13-1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/redis_fdw_13-1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-redis-fdw` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 51.8 KiB | [postgresql-13-redis-fdw_1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-13-redis-fdw_1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-redis-fdw` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 50.2 KiB | [postgresql-13-redis-fdw_1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/redis-fdw/postgresql-13-redis-fdw_1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-redis-fdw` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.2 KiB | [postgresql-13-redis-fdw_1.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/redis-fdw/postgresql-13-redis-fdw_1.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-13-redis-fdw` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 50.5 KiB | [postgresql-13-redis-fdw_1.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/redis-fdw/postgresql-13-redis-fdw_1.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-13-redis-fdw` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 61.5 KiB | [postgresql-13-redis-fdw_1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-13-redis-fdw_1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-13-redis-fdw` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 60.3 KiB | [postgresql-13-redis-fdw_1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/redis-fdw/postgresql-13-redis-fdw_1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-13-redis-fdw` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.1 KiB | [postgresql-13-redis-fdw_1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-13-redis-fdw_1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-13-redis-fdw` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.3 KiB | [postgresql-13-redis-fdw_1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/redis-fdw/postgresql-13-redis-fdw_1.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pg-redis-fdw/redis_fdw" title="Repository" icon="github" subtitle="github.com/pg-redis-fdw/redis_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="redis_fdw-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get redis_fdw; # get redis_fdw source code
pig build dep redis_fdw; # install build dependencies
pig build pkg redis_fdw; # build extension rpm or deb
pig build ext redis_fdw; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install redis_fdw; # install by extension name, for the current active PG version
pig ext install redis_fdw; # install via package alias, for the active PG version
pig ext install redis_fdw -v 18;   # install for PG 18
pig ext install redis_fdw -v 17;   # install for PG 17
pig ext install redis_fdw -v 16;   # install for PG 16
pig ext install redis_fdw -v 15;   # install for PG 15
pig ext install redis_fdw -v 14;   # install for PG 14
pig ext install redis_fdw -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION redis_fdw;
```

