---
title: "redis"
linkTitle: "redis"
description: "Send redis pub/sub messages to Redis from PostgreSQL Directly"
weight: 8720
categories: ["FDW"]
width: full
---

[**pg_redis_pubsub**](https://github.com/brettlaforge/pg_redis_pubsub) : Send redis pub/sub messages to Redis from PostgreSQL Directly


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8720** | {{< badge content="redis" link="https://github.com/brettlaforge/pg_redis_pubsub" >}} | {{< ext "redis" "pg_redis_pubsub" >}} | `0.0.1` | {{< category "FDW" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "redis_fdw" >}} {{< ext "spat" >}} {{< ext "pgmemcache" >}} {{< ext "pg_net" >}} {{< ext "wrappers" >}} {{< ext "kafka_fdw" >}} {{< ext "pgmq" >}} {{< ext "multicorn" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_redis_pubsub` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "pg_redis_pubsub_18" "green" >}} {{< bg "17" "pg_redis_pubsub_17" "green" >}} {{< bg "16" "pg_redis_pubsub_16" "green" >}} {{< bg "15" "pg_redis_pubsub_15" "green" >}} {{< bg "14" "pg_redis_pubsub_14" "green" >}} {{< bg "13" "pg_redis_pubsub_13" "green" >}} | `pg_redis_pubsub_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "postgresql-18-pg-redis-pubsub" "green" >}} {{< bg "17" "postgresql-17-pg-redis-pubsub" "green" >}} {{< bg "16" "postgresql-16-pg-redis-pubsub" "green" >}} {{< bg "15" "postgresql-15-pg-redis-pubsub" "green" >}} {{< bg "14" "postgresql-14-pg-redis-pubsub" "green" >}} {{< bg "13" "postgresql-13-pg-redis-pubsub" "green" >}} | `postgresql-$v-pg-redis-pubsub` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_redis_pubsub_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-redis-pubsub : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-redis-pubsub : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-redis-pubsub : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-redis-pubsub : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-redis-pubsub : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-redis-pubsub : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-redis-pubsub : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-redis-pubsub : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-redis-pubsub : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_redis_pubsub_18` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.7 KiB | [pg_redis_pubsub_18-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_redis_pubsub_18-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_redis_pubsub_18` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.0 KiB | [pg_redis_pubsub_18-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_redis_pubsub_18-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_redis_pubsub_18` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.7 KiB | [pg_redis_pubsub_18-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_redis_pubsub_18-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_redis_pubsub_18` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.7 KiB | [pg_redis_pubsub_18-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_redis_pubsub_18-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_redis_pubsub_18` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.8 KiB | [pg_redis_pubsub_18-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_redis_pubsub_18-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_redis_pubsub_18` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.9 KiB | [pg_redis_pubsub_18-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_redis_pubsub_18-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-redis-pubsub` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.3 KiB | [postgresql-18-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-redis-pubsub/postgresql-18-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-redis-pubsub` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.3 KiB | [postgresql-18-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-redis-pubsub/postgresql-18-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-redis-pubsub` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.2 KiB | [postgresql-18-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-redis-pubsub/postgresql-18-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-redis-pubsub` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.3 KiB | [postgresql-18-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-redis-pubsub/postgresql-18-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-redis-pubsub` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.6 KiB | [postgresql-18-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-redis-pubsub/postgresql-18-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-redis-pubsub` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.5 KiB | [postgresql-18-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-redis-pubsub/postgresql-18-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-redis-pubsub` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.8 KiB | [postgresql-18-pg-redis-pubsub_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-redis-pubsub/postgresql-18-pg-redis-pubsub_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-redis-pubsub` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.8 KiB | [postgresql-18-pg-redis-pubsub_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-redis-pubsub/postgresql-18-pg-redis-pubsub_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_redis_pubsub_17` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.7 KiB | [pg_redis_pubsub_17-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_redis_pubsub_17-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_redis_pubsub_17` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.0 KiB | [pg_redis_pubsub_17-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_redis_pubsub_17-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_redis_pubsub_17` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.7 KiB | [pg_redis_pubsub_17-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_redis_pubsub_17-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_redis_pubsub_17` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.7 KiB | [pg_redis_pubsub_17-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_redis_pubsub_17-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_redis_pubsub_17` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.8 KiB | [pg_redis_pubsub_17-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_redis_pubsub_17-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_redis_pubsub_17` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.9 KiB | [pg_redis_pubsub_17-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_redis_pubsub_17-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-redis-pubsub` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.2 KiB | [postgresql-17-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-redis-pubsub/postgresql-17-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-redis-pubsub` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.3 KiB | [postgresql-17-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-redis-pubsub/postgresql-17-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-redis-pubsub` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.2 KiB | [postgresql-17-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-redis-pubsub/postgresql-17-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-redis-pubsub` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.3 KiB | [postgresql-17-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-redis-pubsub/postgresql-17-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-redis-pubsub` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.1 KiB | [postgresql-17-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-redis-pubsub/postgresql-17-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-redis-pubsub` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.0 KiB | [postgresql-17-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-redis-pubsub/postgresql-17-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-redis-pubsub` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.8 KiB | [postgresql-17-pg-redis-pubsub_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-redis-pubsub/postgresql-17-pg-redis-pubsub_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-redis-pubsub` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.7 KiB | [postgresql-17-pg-redis-pubsub_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-redis-pubsub/postgresql-17-pg-redis-pubsub_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_redis_pubsub_16` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.7 KiB | [pg_redis_pubsub_16-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_redis_pubsub_16-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_redis_pubsub_16` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.0 KiB | [pg_redis_pubsub_16-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_redis_pubsub_16-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_redis_pubsub_16` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.7 KiB | [pg_redis_pubsub_16-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_redis_pubsub_16-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_redis_pubsub_16` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.7 KiB | [pg_redis_pubsub_16-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_redis_pubsub_16-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_redis_pubsub_16` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.8 KiB | [pg_redis_pubsub_16-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_redis_pubsub_16-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_redis_pubsub_16` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.9 KiB | [pg_redis_pubsub_16-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_redis_pubsub_16-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-redis-pubsub` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.2 KiB | [postgresql-16-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-redis-pubsub/postgresql-16-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-redis-pubsub` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.3 KiB | [postgresql-16-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-redis-pubsub/postgresql-16-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-redis-pubsub` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.2 KiB | [postgresql-16-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-redis-pubsub/postgresql-16-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-redis-pubsub` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.3 KiB | [postgresql-16-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-redis-pubsub/postgresql-16-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-redis-pubsub` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.1 KiB | [postgresql-16-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-redis-pubsub/postgresql-16-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-redis-pubsub` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.0 KiB | [postgresql-16-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-redis-pubsub/postgresql-16-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-redis-pubsub` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.8 KiB | [postgresql-16-pg-redis-pubsub_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-redis-pubsub/postgresql-16-pg-redis-pubsub_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-redis-pubsub` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.7 KiB | [postgresql-16-pg-redis-pubsub_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-redis-pubsub/postgresql-16-pg-redis-pubsub_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_redis_pubsub_15` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.8 KiB | [pg_redis_pubsub_15-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_redis_pubsub_15-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_redis_pubsub_15` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.1 KiB | [pg_redis_pubsub_15-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_redis_pubsub_15-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_redis_pubsub_15` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.8 KiB | [pg_redis_pubsub_15-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_redis_pubsub_15-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_redis_pubsub_15` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.8 KiB | [pg_redis_pubsub_15-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_redis_pubsub_15-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_redis_pubsub_15` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.9 KiB | [pg_redis_pubsub_15-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_redis_pubsub_15-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_redis_pubsub_15` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.0 KiB | [pg_redis_pubsub_15-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_redis_pubsub_15-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-redis-pubsub` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.3 KiB | [postgresql-15-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-redis-pubsub/postgresql-15-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-redis-pubsub` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.4 KiB | [postgresql-15-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-redis-pubsub/postgresql-15-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-redis-pubsub` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.3 KiB | [postgresql-15-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-redis-pubsub/postgresql-15-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-redis-pubsub` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.4 KiB | [postgresql-15-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-redis-pubsub/postgresql-15-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-redis-pubsub` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.1 KiB | [postgresql-15-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-redis-pubsub/postgresql-15-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-redis-pubsub` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.0 KiB | [postgresql-15-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-redis-pubsub/postgresql-15-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-redis-pubsub` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.9 KiB | [postgresql-15-pg-redis-pubsub_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-redis-pubsub/postgresql-15-pg-redis-pubsub_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-redis-pubsub` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.8 KiB | [postgresql-15-pg-redis-pubsub_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-redis-pubsub/postgresql-15-pg-redis-pubsub_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_redis_pubsub_14` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.8 KiB | [pg_redis_pubsub_14-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_redis_pubsub_14-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_redis_pubsub_14` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.1 KiB | [pg_redis_pubsub_14-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_redis_pubsub_14-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_redis_pubsub_14` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.8 KiB | [pg_redis_pubsub_14-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_redis_pubsub_14-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_redis_pubsub_14` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.8 KiB | [pg_redis_pubsub_14-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_redis_pubsub_14-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_redis_pubsub_14` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.9 KiB | [pg_redis_pubsub_14-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_redis_pubsub_14-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_redis_pubsub_14` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.0 KiB | [pg_redis_pubsub_14-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_redis_pubsub_14-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-redis-pubsub` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.3 KiB | [postgresql-14-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-redis-pubsub/postgresql-14-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-redis-pubsub` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.3 KiB | [postgresql-14-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-redis-pubsub/postgresql-14-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-redis-pubsub` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.3 KiB | [postgresql-14-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-redis-pubsub/postgresql-14-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-redis-pubsub` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.4 KiB | [postgresql-14-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-redis-pubsub/postgresql-14-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-redis-pubsub` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.1 KiB | [postgresql-14-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-redis-pubsub/postgresql-14-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-redis-pubsub` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.0 KiB | [postgresql-14-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-redis-pubsub/postgresql-14-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-redis-pubsub` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.9 KiB | [postgresql-14-pg-redis-pubsub_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-redis-pubsub/postgresql-14-pg-redis-pubsub_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-redis-pubsub` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.8 KiB | [postgresql-14-pg-redis-pubsub_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-redis-pubsub/postgresql-14-pg-redis-pubsub_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_redis_pubsub_13` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.8 KiB | [pg_redis_pubsub_13-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_redis_pubsub_13-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_redis_pubsub_13` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.1 KiB | [pg_redis_pubsub_13-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_redis_pubsub_13-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_redis_pubsub_13` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.7 KiB | [pg_redis_pubsub_13-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_redis_pubsub_13-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_redis_pubsub_13` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.8 KiB | [pg_redis_pubsub_13-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_redis_pubsub_13-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_redis_pubsub_13` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.9 KiB | [pg_redis_pubsub_13-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_redis_pubsub_13-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_redis_pubsub_13` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.0 KiB | [pg_redis_pubsub_13-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_redis_pubsub_13-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-redis-pubsub` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.1 KiB | [postgresql-13-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-redis-pubsub/postgresql-13-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-redis-pubsub` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.1 KiB | [postgresql-13-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-redis-pubsub/postgresql-13-pg-redis-pubsub_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-redis-pubsub` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.1 KiB | [postgresql-13-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-redis-pubsub/postgresql-13-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-redis-pubsub` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.1 KiB | [postgresql-13-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-redis-pubsub/postgresql-13-pg-redis-pubsub_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-redis-pubsub` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.9 KiB | [postgresql-13-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-redis-pubsub/postgresql-13-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-redis-pubsub` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.0 KiB | [postgresql-13-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-redis-pubsub/postgresql-13-pg-redis-pubsub_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-redis-pubsub` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.7 KiB | [postgresql-13-pg-redis-pubsub_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-redis-pubsub/postgresql-13-pg-redis-pubsub_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-redis-pubsub` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.6 KiB | [postgresql-13-pg-redis-pubsub_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-redis-pubsub/postgresql-13-pg-redis-pubsub_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/brettlaforge/pg_redis_pubsub" title="Repository" icon="github" subtitle="github.com/brettlaforge/pg_redis_pubsub" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_redis_pubsub-0.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_redis_pubsub;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_redis_pubsub;		# install via package name, for the active PG version
pig install redis;		# install by extension name, for the current active PG version

pig install redis -v 18;   # install for PG 18
pig install redis -v 17;   # install for PG 17
pig install redis -v 16;   # install for PG 16
pig install redis -v 15;   # install for PG 15
pig install redis -v 14;   # install for PG 14
pig install redis -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION redis;
```
