---
title: "pg_mooncake"
linkTitle: "pg_mooncake"
description: "Columnstore Table in Postgres"
weight: 2440
categories: ["OLAP"]
width: full
---

Columnstore Table in Postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2440** | {{< badge content="pg_mooncake" link="https://github.com/Mooncake-Labs/pg_mooncake" >}} | {{< ext "pg_mooncake" >}} | `0.1.2` | {{< category "OLAP" >}} | {{< license "MIT" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_duckdb" >}} {{< ext "duckdb_fdw" >}} {{< ext "pg_analytics" >}} {{< ext "columnar" >}} {{< ext "citus_columnar" >}} {{< ext "pg_parquet" >}} {{< ext "orioledb" >}} {{< ext "timescaledb" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_mooncake" >}} | `0.1.2` | {{< bg "18" "pg_mooncake_18*" "red" >}} {{< bg "17" "pg_mooncake_17*" "green" >}} {{< bg "16" "pg_mooncake_16*" "green" >}} {{< bg "15" "pg_mooncake_15*" "green" >}} {{< bg "14" "pg_mooncake_14*" "green" >}} {{< bg "13" "pg_mooncake_13*" "red" >}} | `pg_mooncake_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_mooncake" >}} | `0.1.2` | {{< bg "18" "postgresql-18-pg-mooncake" "red" >}} {{< bg "17" "postgresql-17-pg-mooncake" "green" >}} {{< bg "16" "postgresql-16-pg-mooncake" "green" >}} {{< bg "15" "postgresql-15-pg-mooncake" "green" >}} {{< bg "14" "postgresql-14-pg-mooncake" "green" >}} {{< bg "13" "postgresql-13-pg-mooncake" "red" >}} | `postgresql-$v-pg-mooncake` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pg_mooncake_18 : MISS 0" "red" >}}      |      {{< bg "PIGSTY 0.1.2" "pg_mooncake_17 : BREAK 1" "orange" >}}      |      {{< bg "PIGSTY 0.1.2" "pg_mooncake_16 : BREAK 1" "orange" >}}      |      {{< bg "PIGSTY 0.1.2" "pg_mooncake_15 : BREAK 1" "orange" >}}      |      {{< bg "PIGSTY 0.1.2" "pg_mooncake_14 : BREAK 1" "orange" >}}      |      {{< bg "MISS" "pg_mooncake_13 : MISS 0" "red" >}}      |
|    `el8.aarch64`    |      {{< bg "MISS" "pg_mooncake_18 : MISS 0" "red" >}}      |      {{< bg "PIGSTY 0.1.2" "pg_mooncake_17 : BREAK 1" "orange" >}}      |      {{< bg "PIGSTY 0.1.2" "pg_mooncake_16 : BREAK 1" "orange" >}}      |      {{< bg "PIGSTY 0.1.2" "pg_mooncake_15 : BREAK 1" "orange" >}}      |      {{< bg "PIGSTY 0.1.2" "pg_mooncake_14 : BREAK 1" "orange" >}}      |      {{< bg "MISS" "pg_mooncake_13 : MISS 0" "red" >}}      |
|    `el9.x86_64`    |      {{< bg "MISS" "pg_mooncake_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.2" "pg_mooncake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_mooncake_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_mooncake_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_mooncake_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_mooncake_13 : MISS 0" "red" >}}      |
|    `el9.aarch64`    |      {{< bg "MISS" "pg_mooncake_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.2" "pg_mooncake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_mooncake_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_mooncake_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_mooncake_14 : AVAIL 2" "green" >}} |      {{< bg "MISS" "pg_mooncake_13 : MISS 0" "red" >}}      |
|    `el10.x86_64`    |      {{< bg "MISS" "pg_mooncake_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_mooncake_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_mooncake_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_mooncake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_mooncake_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_mooncake_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "pg_mooncake_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_mooncake_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_mooncake_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_mooncake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_mooncake_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_mooncake_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-mooncake : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.2" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-mooncake : MISS 0" "red" >}}      |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-mooncake : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.2" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-mooncake : MISS 0" "red" >}}      |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-mooncake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-mooncake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-mooncake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-mooncake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-mooncake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-mooncake : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-mooncake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-mooncake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-mooncake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-mooncake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-mooncake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-mooncake : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-mooncake : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.2" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-mooncake : MISS 0" "red" >}}      |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-mooncake : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.2" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-mooncake : MISS 0" "red" >}}      |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-mooncake : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.2" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-mooncake : MISS 0" "red" >}}      |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-mooncake : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.2" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-mooncake : MISS 0" "red" >}}      |


{{< tabs items="PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mooncake_17` | 0.1.2 | `el8.x86_64` | pigsty | 29.9 MiB | [pg_mooncake_17-0.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mooncake_17-0.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_mooncake_17` | 0.1.2 | `el8.aarch64` | pigsty | 27.0 MiB | [pg_mooncake_17-0.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mooncake_17-0.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_mooncake_17` | 0.1.2 | `el9.x86_64` | pigsty | 29.3 MiB | [pg_mooncake_17-0.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mooncake_17-0.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_mooncake_17` | 0.1.2 | `el9.aarch64` | pigsty | 27.6 MiB | [pg_mooncake_17-0.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mooncake_17-0.1.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pg-mooncake` | 0.1.2 | `d12.x86_64` | pigsty | 23.6 MiB | [postgresql-17-pg-mooncake_0.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-mooncake` | 0.1.2 | `d12.aarch64` | pigsty | 20.7 MiB | [postgresql-17-pg-mooncake_0.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-mooncake` | 0.1.2 | `u22.x86_64` | pigsty | 25.1 MiB | [postgresql-17-pg-mooncake_0.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-mooncake` | 0.1.2 | `u22.aarch64` | pigsty | 23.1 MiB | [postgresql-17-pg-mooncake_0.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-mooncake` | 0.1.2 | `u24.x86_64` | pigsty | 25.1 MiB | [postgresql-17-pg-mooncake_0.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-mooncake` | 0.1.2 | `u24.aarch64` | pigsty | 23.3 MiB | [postgresql-17-pg-mooncake_0.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mooncake_16` | 0.1.2 | `el8.x86_64` | pigsty | 29.9 MiB | [pg_mooncake_16-0.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mooncake_16-0.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_mooncake_16` | 0.1.2 | `el8.aarch64` | pigsty | 27.0 MiB | [pg_mooncake_16-0.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mooncake_16-0.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_mooncake_16` | 0.1.2 | `el9.x86_64` | pigsty | 29.3 MiB | [pg_mooncake_16-0.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mooncake_16-0.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_mooncake_16` | 0.1.2 | `el9.aarch64` | pigsty | 27.6 MiB | [pg_mooncake_16-0.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mooncake_16-0.1.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_mooncake_16` | 0.1.1 | `el9.aarch64` | pigsty | 27.7 MiB | [pg_mooncake_16-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mooncake_16-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-mooncake` | 0.1.2 | `d12.x86_64` | pigsty | 23.6 MiB | [postgresql-16-pg-mooncake_0.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-mooncake` | 0.1.2 | `d12.aarch64` | pigsty | 20.7 MiB | [postgresql-16-pg-mooncake_0.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-mooncake` | 0.1.2 | `u22.x86_64` | pigsty | 25.1 MiB | [postgresql-16-pg-mooncake_0.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-mooncake` | 0.1.2 | `u22.aarch64` | pigsty | 23.1 MiB | [postgresql-16-pg-mooncake_0.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-mooncake` | 0.1.2 | `u24.x86_64` | pigsty | 25.1 MiB | [postgresql-16-pg-mooncake_0.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-mooncake` | 0.1.2 | `u24.aarch64` | pigsty | 23.3 MiB | [postgresql-16-pg-mooncake_0.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mooncake_15` | 0.1.2 | `el8.x86_64` | pigsty | 30.0 MiB | [pg_mooncake_15-0.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mooncake_15-0.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_mooncake_15` | 0.1.2 | `el8.aarch64` | pigsty | 27.0 MiB | [pg_mooncake_15-0.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mooncake_15-0.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_mooncake_15` | 0.1.2 | `el9.x86_64` | pigsty | 29.3 MiB | [pg_mooncake_15-0.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mooncake_15-0.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_mooncake_15` | 0.1.2 | `el9.aarch64` | pigsty | 27.7 MiB | [pg_mooncake_15-0.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mooncake_15-0.1.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_mooncake_15` | 0.1.1 | `el9.aarch64` | pigsty | 27.7 MiB | [pg_mooncake_15-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mooncake_15-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-mooncake` | 0.1.2 | `d12.x86_64` | pigsty | 23.6 MiB | [postgresql-15-pg-mooncake_0.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-mooncake` | 0.1.2 | `d12.aarch64` | pigsty | 20.7 MiB | [postgresql-15-pg-mooncake_0.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-mooncake` | 0.1.2 | `u22.x86_64` | pigsty | 25.1 MiB | [postgresql-15-pg-mooncake_0.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-mooncake` | 0.1.2 | `u22.aarch64` | pigsty | 23.1 MiB | [postgresql-15-pg-mooncake_0.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-mooncake` | 0.1.2 | `u24.x86_64` | pigsty | 25.2 MiB | [postgresql-15-pg-mooncake_0.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-mooncake` | 0.1.2 | `u24.aarch64` | pigsty | 23.3 MiB | [postgresql-15-pg-mooncake_0.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mooncake_14` | 0.1.2 | `el8.x86_64` | pigsty | 30.0 MiB | [pg_mooncake_14-0.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mooncake_14-0.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_mooncake_14` | 0.1.2 | `el8.aarch64` | pigsty | 27.0 MiB | [pg_mooncake_14-0.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mooncake_14-0.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_mooncake_14` | 0.1.2 | `el9.x86_64` | pigsty | 29.4 MiB | [pg_mooncake_14-0.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mooncake_14-0.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_mooncake_14` | 0.1.2 | `el9.aarch64` | pigsty | 27.7 MiB | [pg_mooncake_14-0.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mooncake_14-0.1.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_mooncake_14` | 0.1.1 | `el9.aarch64` | pigsty | 27.7 MiB | [pg_mooncake_14-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mooncake_14-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-mooncake` | 0.1.2 | `d12.x86_64` | pigsty | 23.6 MiB | [postgresql-14-pg-mooncake_0.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-mooncake` | 0.1.2 | `d12.aarch64` | pigsty | 20.7 MiB | [postgresql-14-pg-mooncake_0.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-mooncake` | 0.1.2 | `u22.x86_64` | pigsty | 25.1 MiB | [postgresql-14-pg-mooncake_0.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-mooncake` | 0.1.2 | `u22.aarch64` | pigsty | 23.1 MiB | [postgresql-14-pg-mooncake_0.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-mooncake` | 0.1.2 | `u24.x86_64` | pigsty | 25.2 MiB | [postgresql-14-pg-mooncake_0.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-mooncake` | 0.1.2 | `u24.aarch64` | pigsty | 23.4 MiB | [postgresql-14-pg-mooncake_0.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Mooncake-Labs/pg_mooncake" title="Repository" icon="github" subtitle="github.com/Mooncake-Labs/pg_mooncake" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_mooncake-0.1.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_mooncake; # get pg_mooncake source code
pig build dep pg_mooncake; # install build dependencies
pig build pkg pg_mooncake; # build extension rpm or deb
pig build ext pg_mooncake; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_mooncake; # install by extension name, for the current active PG version
pig ext install pg_mooncake; # install via package alias, for the active PG version
pig ext install pg_mooncake -v 17;   # install for PG 17
pig ext install pg_mooncake -v 16;   # install for PG 16
pig ext install pg_mooncake -v 15;   # install for PG 15
pig ext install pg_mooncake -v 14;   # install for PG 14

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_mooncake;
```




--------

## Usage

pg_mooncake creates real-time columnstore mirrors of PostgreSQL tables using Apache Iceberg format, enabling fast analytics queries with sub-second data freshness while maintaining full PostgreSQL compatibility.

> **Warning**: This extension conflicts with `pg_duckdb` & `duckdb_fdw` due to shared `libduckdb.so` library usage.

### Real-time Columnstore Mirrors

Create analytical mirrors of existing PostgreSQL tables that automatically sync data changes.

```sql
-- Create extension
CREATE EXTENSION pg_mooncake;

-- Create source table
CREATE TABLE sales_data(
    id bigint PRIMARY KEY,
    product_id int,
    sale_date timestamp,
    amount decimal(10,2),
    region text
);

-- Create real-time columnstore mirror
CALL mooncake.create_table('sales_analytics', 'sales_data');

-- Insert data (automatically synced to mirror)
INSERT INTO sales_data VALUES 
    (1, 101, '2024-01-15 10:30:00', 1250.00, 'North'),
    (2, 102, '2024-01-15 11:45:00', 890.50, 'South');

-- Query columnstore mirror for analytics
SELECT 
    region,
    DATE_TRUNC('month', sale_date) as month,
    SUM(amount) as total_sales,
    AVG(amount) as avg_sale
FROM sales_analytics 
GROUP BY region, month;
```

### Direct Columnstore Tables

Create columnstore tables directly for analytical workloads.

```sql
-- Create columnstore table
CREATE TABLE user_activity(
    user_id BIGINT,
    activity_type TEXT,
    activity_timestamp TIMESTAMP,
    duration INT
) USING columnstore;

-- Insert data
INSERT INTO user_activity VALUES
    (1, 'login', '2024-01-01 08:00:00', 120),
    (2, 'page_view', '2024-01-01 08:05:00', 30),
    (3, 'logout', '2024-01-01 08:30:00', 60);

-- Run analytical queries
SELECT
    user_id,
    activity_type,
    SUM(duration) AS total_duration,
    COUNT(*) AS activity_count
FROM user_activity
GROUP BY user_id, activity_type
ORDER BY user_id, activity_type;
```

### Cloud Storage Integration

Configure pg_mooncake to use cloud storage backends like S3.

```sql
-- Create S3 credentials
SELECT mooncake.create_secret(
    'my_s3_secret', 
    'S3', 
    'ACCESS_KEY_ID', 
    'SECRET_ACCESS_KEY', 
    '{"REGION": "us-west-2"}'
);

-- Set default S3 bucket
SET mooncake.default_bucket = 's3://my-analytics-bucket';

-- Disable local cache for cloud environments
SET mooncake.enable_local_cache = false;

-- Create columnstore table using S3 storage
CREATE TABLE metrics_data(
    timestamp TIMESTAMP,
    metric_name TEXT,
    value DOUBLE PRECISION,
    tags JSONB
) USING columnstore;
```

### Multi-Format Data Loading

Load data from various file formats into columnstore tables.

```sql
-- Load from Parquet files
COPY metrics_data FROM 's3://data-lake/metrics/*.parquet';

-- Load from CSV files  
COPY user_events FROM '/path/to/events.csv' WITH CSV HEADER;

-- Load from JSON files
COPY log_data FROM 's3://logs/application.json' WITH (FORMAT json);
```

### External Data Lake Access

Query external Iceberg and Delta Lake tables directly.

```sql
-- Query external Iceberg table
SELECT COUNT(*) FROM iceberg_scan('s3://datalake/sales_iceberg/');

-- Query Delta Lake table
SELECT * FROM delta_scan('s3://datalake/user_events_delta/') 
WHERE event_date >= '2024-01-01';
```

### Hybrid Workloads

Combine transactional and analytical queries seamlessly.

```sql
-- Join columnstore table with regular PostgreSQL table
SELECT 
    t.symbol,
    t.price,
    c.company_name,
    c.sector
FROM trades_analytics t 
JOIN companies c ON t.symbol = c.symbol
WHERE t.trade_date >= CURRENT_DATE - INTERVAL '7 days'
  AND t.volume > 1000000;

-- Mixed workload: OLTP inserts with OLAP queries
BEGIN;
INSERT INTO orders (customer_id, product_id, quantity, price) 
VALUES (1001, 5502, 2, 29.99);

-- Immediate analytics on fresh data
SELECT 
    product_id,
    SUM(quantity * price) as revenue,
    COUNT(*) as order_count
FROM orders_analytics
WHERE order_date >= CURRENT_DATE
GROUP BY product_id;
COMMIT;
```


--------

## Performance Features

### Query Acceleration

Queries on columnstore tables use DuckDB execution engine for high performance:

```sql
EXPLAIN SELECT 
    region,
    AVG(amount) as avg_sale,
    PERCENTILE_CONT(0.95) WITHIN GROUP (ORDER BY amount) as p95
FROM sales_analytics 
GROUP BY region;

-- Shows: Custom Scan (DuckDBScan) with optimized execution plan
```

### Configuration Options

```sql
-- Set maximum memory usage (default: 1GB)
SET mooncake.max_memory = '4GB';

-- Configure thread count for parallel processing
SET mooncake.threads = 8;

-- Enable/disable local caching
SET mooncake.enable_local_cache = true;

-- Set data compression level
SET mooncake.compression = 'zstd';
```


--------

## Use Cases

- **Real-time Analytics**: Live dashboards and reports with sub-second data freshness
- **Hybrid Workloads**: Applications requiring both OLTP and OLAP capabilities  
- **Data Lake Integration**: Bridge PostgreSQL with modern data lake architectures
- **Performance Optimization**: Accelerate analytical queries without application changes
- **Operational Analytics**: Real-time monitoring and alerting on transactional data

pg_mooncake provides PostgreSQL-native columnstore capabilities while maintaining compatibility with existing applications and enabling integration with modern data lake ecosystems through Apache Iceberg format.