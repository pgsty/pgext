---
title: "pglogical_ticker"
linkTitle: "pglogical_ticker"
description: "Have an accurate view on pglogical replication delay"
weight: 9510
categories: ["ETL"]
width: full
---

Have an accurate view on pglogical replication delay


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9510** | {{< badge content="pglogical_ticker" link="https://github.com/enova/pglogical_ticker" >}} | {{< ext "pglogical_ticker" >}} | `1.4.1` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pglogical" >}} |
|   **See Also**    | {{< ext "pglogical_origin" >}} {{< ext "pgl_ddl_deploy" >}} {{< ext "pg_failover_slots" >}} {{< ext "pgactive" >}} {{< ext "wal2json" >}} {{< ext "decoderbufs" >}} {{< ext "repmgr" >}} {{< ext "decoder_raw" >}} |

> [!Note] require a patch on el, pg18 break on el


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pglogical_ticker" >}} | `1.4.1` | {{< bg "18" "pglogical_ticker_18*" "red" >}} {{< bg "17" "pglogical_ticker_17*" "green" >}} {{< bg "16" "pglogical_ticker_16*" "green" >}} {{< bg "15" "pglogical_ticker_15*" "green" >}} {{< bg "14" "pglogical_ticker_14*" "green" >}} {{< bg "13" "pglogical_ticker_13*" "green" >}} | `pglogical_ticker_$v*` | `pglogical_$v` |
| **Debian** | {{< badge content="PGDG" link="/e/pglogical_ticker" >}} | `1.4.1` | {{< bg "18" "postgresql-18-pglogical-ticker" "red" >}} {{< bg "17" "postgresql-17-pglogical-ticker" "green" >}} {{< bg "16" "postgresql-16-pglogical-ticker" "green" >}} {{< bg "15" "postgresql-15-pglogical-ticker" "green" >}} {{< bg "14" "postgresql-14-pglogical-ticker" "green" >}} {{< bg "13" "postgresql-13-pglogical-ticker" "green" >}} | `postgresql-$v-pglogical-ticker` | `postgresql-$v-pglogical` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pglogical_ticker_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pglogical_ticker_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pglogical_ticker_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pglogical_ticker_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pglogical_ticker_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    |      {{< bg "MISS" "pglogical_ticker_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglogical_ticker_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglogical_ticker_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglogical_ticker_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglogical_ticker_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglogical_ticker_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "pglogical_ticker_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglogical_ticker_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglogical_ticker_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglogical_ticker_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglogical_ticker_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pglogical_ticker_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pglogical-ticker : MISS 0" "red" >}}      | {{< bg "PGDG 1.4.1" "postgresql-17-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-16-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-15-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-14-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-13-pglogical-ticker : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pglogical-ticker : MISS 0" "red" >}}      | {{< bg "PGDG 1.4.1" "postgresql-17-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-16-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-15-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-14-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-13-pglogical-ticker : AVAIL 1" "blue" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pglogical-ticker : MISS 0" "red" >}}      | {{< bg "PGDG 1.4.1" "postgresql-17-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-16-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-15-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-14-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-13-pglogical-ticker : AVAIL 1" "blue" >}} |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pglogical-ticker : MISS 0" "red" >}}      | {{< bg "PGDG 1.4.1" "postgresql-17-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-16-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-15-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-14-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-13-pglogical-ticker : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pglogical-ticker : MISS 0" "red" >}}      | {{< bg "PGDG 1.4.1" "postgresql-17-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-16-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-15-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-14-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-13-pglogical-ticker : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pglogical-ticker : MISS 0" "red" >}}      | {{< bg "PGDG 1.4.1" "postgresql-17-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-16-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-15-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-14-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-13-pglogical-ticker : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pglogical-ticker : MISS 0" "red" >}}      | {{< bg "PGDG 1.4.1" "postgresql-17-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-16-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-15-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-14-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-13-pglogical-ticker : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pglogical-ticker : MISS 0" "red" >}}      | {{< bg "PGDG 1.4.1" "postgresql-17-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-16-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-15-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-14-pglogical-ticker : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-13-pglogical-ticker : AVAIL 1" "blue" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglogical_ticker_17` | 1.4.1 | `el8.x86_64` | pigsty | 17.4 KiB | [pglogical_ticker_17-1.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglogical_ticker_17-1.4.1-1PIGSTY.el8.x86_64.rpm) |
| `pglogical_ticker_17` | 1.4.1 | `el8.aarch64` | pigsty | 17.4 KiB | [pglogical_ticker_17-1.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglogical_ticker_17-1.4.1-1PIGSTY.el8.aarch64.rpm) |
| `pglogical_ticker_17` | 1.4.1 | `el9.x86_64` | pigsty | 17.6 KiB | [pglogical_ticker_17-1.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglogical_ticker_17-1.4.1-1PIGSTY.el9.x86_64.rpm) |
| `pglogical_ticker_17` | 1.4.1 | `el9.aarch64` | pigsty | 17.4 KiB | [pglogical_ticker_17-1.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglogical_ticker_17-1.4.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pglogical-ticker` | 1.4.1 | `d12.x86_64` | pgdg | 20.9 KiB | [postgresql-17-pglogical-ticker_1.4.1-8.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-17-pglogical-ticker_1.4.1-8.pgdg120+1_amd64.deb) |
| `postgresql-17-pglogical-ticker` | 1.4.1 | `d12.aarch64` | pgdg | 20.6 KiB | [postgresql-17-pglogical-ticker_1.4.1-8.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-17-pglogical-ticker_1.4.1-8.pgdg120+1_arm64.deb) |
| `postgresql-17-pglogical-ticker` | 1.4.1 | `d13.x86_64` | pgdg | 20.8 KiB | [postgresql-17-pglogical-ticker_1.4.1-8.pgdg130+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-17-pglogical-ticker_1.4.1-8.pgdg130+2_amd64.deb) |
| `postgresql-17-pglogical-ticker` | 1.4.1 | `d13.aarch64` | pgdg | 20.7 KiB | [postgresql-17-pglogical-ticker_1.4.1-8.pgdg130+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-17-pglogical-ticker_1.4.1-8.pgdg130+2_arm64.deb) |
| `postgresql-17-pglogical-ticker` | 1.4.1 | `u22.x86_64` | pgdg | 20.9 KiB | [postgresql-17-pglogical-ticker_1.4.1-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-17-pglogical-ticker_1.4.1-8.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pglogical-ticker` | 1.4.1 | `u22.aarch64` | pgdg | 20.6 KiB | [postgresql-17-pglogical-ticker_1.4.1-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-17-pglogical-ticker_1.4.1-8.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pglogical-ticker` | 1.4.1 | `u24.x86_64` | pgdg | 20.4 KiB | [postgresql-17-pglogical-ticker_1.4.1-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-17-pglogical-ticker_1.4.1-8.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pglogical-ticker` | 1.4.1 | `u24.aarch64` | pgdg | 20.0 KiB | [postgresql-17-pglogical-ticker_1.4.1-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-17-pglogical-ticker_1.4.1-8.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglogical_ticker_16` | 1.4.1 | `el8.x86_64` | pigsty | 17.4 KiB | [pglogical_ticker_16-1.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglogical_ticker_16-1.4.1-1PIGSTY.el8.x86_64.rpm) |
| `pglogical_ticker_16` | 1.4.1 | `el8.aarch64` | pigsty | 17.4 KiB | [pglogical_ticker_16-1.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglogical_ticker_16-1.4.1-1PIGSTY.el8.aarch64.rpm) |
| `pglogical_ticker_16` | 1.4.1 | `el9.x86_64` | pigsty | 17.6 KiB | [pglogical_ticker_16-1.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglogical_ticker_16-1.4.1-1PIGSTY.el9.x86_64.rpm) |
| `pglogical_ticker_16` | 1.4.1 | `el9.aarch64` | pigsty | 17.4 KiB | [pglogical_ticker_16-1.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglogical_ticker_16-1.4.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pglogical-ticker` | 1.4.1 | `d12.x86_64` | pgdg | 20.8 KiB | [postgresql-16-pglogical-ticker_1.4.1-8.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-16-pglogical-ticker_1.4.1-8.pgdg120+1_amd64.deb) |
| `postgresql-16-pglogical-ticker` | 1.4.1 | `d12.aarch64` | pgdg | 20.7 KiB | [postgresql-16-pglogical-ticker_1.4.1-8.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-16-pglogical-ticker_1.4.1-8.pgdg120+1_arm64.deb) |
| `postgresql-16-pglogical-ticker` | 1.4.1 | `d13.x86_64` | pgdg | 20.8 KiB | [postgresql-16-pglogical-ticker_1.4.1-8.pgdg130+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-16-pglogical-ticker_1.4.1-8.pgdg130+2_amd64.deb) |
| `postgresql-16-pglogical-ticker` | 1.4.1 | `d13.aarch64` | pgdg | 20.7 KiB | [postgresql-16-pglogical-ticker_1.4.1-8.pgdg130+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-16-pglogical-ticker_1.4.1-8.pgdg130+2_arm64.deb) |
| `postgresql-16-pglogical-ticker` | 1.4.1 | `u22.x86_64` | pgdg | 20.9 KiB | [postgresql-16-pglogical-ticker_1.4.1-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-16-pglogical-ticker_1.4.1-8.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pglogical-ticker` | 1.4.1 | `u22.aarch64` | pgdg | 20.6 KiB | [postgresql-16-pglogical-ticker_1.4.1-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-16-pglogical-ticker_1.4.1-8.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pglogical-ticker` | 1.4.1 | `u24.x86_64` | pgdg | 20.4 KiB | [postgresql-16-pglogical-ticker_1.4.1-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-16-pglogical-ticker_1.4.1-8.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pglogical-ticker` | 1.4.1 | `u24.aarch64` | pgdg | 20.0 KiB | [postgresql-16-pglogical-ticker_1.4.1-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-16-pglogical-ticker_1.4.1-8.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglogical_ticker_15` | 1.4.1 | `el8.x86_64` | pigsty | 17.4 KiB | [pglogical_ticker_15-1.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglogical_ticker_15-1.4.1-1PIGSTY.el8.x86_64.rpm) |
| `pglogical_ticker_15` | 1.4.1 | `el8.aarch64` | pigsty | 17.4 KiB | [pglogical_ticker_15-1.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglogical_ticker_15-1.4.1-1PIGSTY.el8.aarch64.rpm) |
| `pglogical_ticker_15` | 1.4.1 | `el9.x86_64` | pigsty | 17.6 KiB | [pglogical_ticker_15-1.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglogical_ticker_15-1.4.1-1PIGSTY.el9.x86_64.rpm) |
| `pglogical_ticker_15` | 1.4.1 | `el9.aarch64` | pigsty | 17.4 KiB | [pglogical_ticker_15-1.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglogical_ticker_15-1.4.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pglogical-ticker` | 1.4.1 | `d12.x86_64` | pgdg | 20.8 KiB | [postgresql-15-pglogical-ticker_1.4.1-8.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-15-pglogical-ticker_1.4.1-8.pgdg120+1_amd64.deb) |
| `postgresql-15-pglogical-ticker` | 1.4.1 | `d12.aarch64` | pgdg | 20.6 KiB | [postgresql-15-pglogical-ticker_1.4.1-8.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-15-pglogical-ticker_1.4.1-8.pgdg120+1_arm64.deb) |
| `postgresql-15-pglogical-ticker` | 1.4.1 | `d13.x86_64` | pgdg | 20.8 KiB | [postgresql-15-pglogical-ticker_1.4.1-8.pgdg130+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-15-pglogical-ticker_1.4.1-8.pgdg130+2_amd64.deb) |
| `postgresql-15-pglogical-ticker` | 1.4.1 | `d13.aarch64` | pgdg | 20.7 KiB | [postgresql-15-pglogical-ticker_1.4.1-8.pgdg130+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-15-pglogical-ticker_1.4.1-8.pgdg130+2_arm64.deb) |
| `postgresql-15-pglogical-ticker` | 1.4.1 | `u22.x86_64` | pgdg | 20.9 KiB | [postgresql-15-pglogical-ticker_1.4.1-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-15-pglogical-ticker_1.4.1-8.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pglogical-ticker` | 1.4.1 | `u22.aarch64` | pgdg | 20.6 KiB | [postgresql-15-pglogical-ticker_1.4.1-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-15-pglogical-ticker_1.4.1-8.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pglogical-ticker` | 1.4.1 | `u24.x86_64` | pgdg | 20.4 KiB | [postgresql-15-pglogical-ticker_1.4.1-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-15-pglogical-ticker_1.4.1-8.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pglogical-ticker` | 1.4.1 | `u24.aarch64` | pgdg | 20.1 KiB | [postgresql-15-pglogical-ticker_1.4.1-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-15-pglogical-ticker_1.4.1-8.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglogical_ticker_14` | 1.4.1 | `el8.x86_64` | pigsty | 17.4 KiB | [pglogical_ticker_14-1.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglogical_ticker_14-1.4.1-1PIGSTY.el8.x86_64.rpm) |
| `pglogical_ticker_14` | 1.4.1 | `el8.aarch64` | pigsty | 17.4 KiB | [pglogical_ticker_14-1.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglogical_ticker_14-1.4.1-1PIGSTY.el8.aarch64.rpm) |
| `pglogical_ticker_14` | 1.4.1 | `el9.x86_64` | pigsty | 17.5 KiB | [pglogical_ticker_14-1.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglogical_ticker_14-1.4.1-1PIGSTY.el9.x86_64.rpm) |
| `pglogical_ticker_14` | 1.4.1 | `el9.aarch64` | pigsty | 17.4 KiB | [pglogical_ticker_14-1.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglogical_ticker_14-1.4.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pglogical-ticker` | 1.4.1 | `d12.x86_64` | pgdg | 20.8 KiB | [postgresql-14-pglogical-ticker_1.4.1-8.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-14-pglogical-ticker_1.4.1-8.pgdg120+1_amd64.deb) |
| `postgresql-14-pglogical-ticker` | 1.4.1 | `d12.aarch64` | pgdg | 20.6 KiB | [postgresql-14-pglogical-ticker_1.4.1-8.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-14-pglogical-ticker_1.4.1-8.pgdg120+1_arm64.deb) |
| `postgresql-14-pglogical-ticker` | 1.4.1 | `d13.x86_64` | pgdg | 20.8 KiB | [postgresql-14-pglogical-ticker_1.4.1-8.pgdg130+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-14-pglogical-ticker_1.4.1-8.pgdg130+2_amd64.deb) |
| `postgresql-14-pglogical-ticker` | 1.4.1 | `d13.aarch64` | pgdg | 20.7 KiB | [postgresql-14-pglogical-ticker_1.4.1-8.pgdg130+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-14-pglogical-ticker_1.4.1-8.pgdg130+2_arm64.deb) |
| `postgresql-14-pglogical-ticker` | 1.4.1 | `u22.x86_64` | pgdg | 20.8 KiB | [postgresql-14-pglogical-ticker_1.4.1-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-14-pglogical-ticker_1.4.1-8.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pglogical-ticker` | 1.4.1 | `u22.aarch64` | pgdg | 20.6 KiB | [postgresql-14-pglogical-ticker_1.4.1-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-14-pglogical-ticker_1.4.1-8.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pglogical-ticker` | 1.4.1 | `u24.x86_64` | pgdg | 20.4 KiB | [postgresql-14-pglogical-ticker_1.4.1-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-14-pglogical-ticker_1.4.1-8.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pglogical-ticker` | 1.4.1 | `u24.aarch64` | pgdg | 20.0 KiB | [postgresql-14-pglogical-ticker_1.4.1-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-14-pglogical-ticker_1.4.1-8.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglogical_ticker_13` | 1.4.1 | `el8.x86_64` | pigsty | 17.4 KiB | [pglogical_ticker_13-1.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglogical_ticker_13-1.4.1-1PIGSTY.el8.x86_64.rpm) |
| `pglogical_ticker_13` | 1.4.1 | `el8.aarch64` | pigsty | 17.4 KiB | [pglogical_ticker_13-1.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglogical_ticker_13-1.4.1-1PIGSTY.el8.aarch64.rpm) |
| `pglogical_ticker_13` | 1.4.1 | `el9.x86_64` | pigsty | 17.5 KiB | [pglogical_ticker_13-1.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglogical_ticker_13-1.4.1-1PIGSTY.el9.x86_64.rpm) |
| `pglogical_ticker_13` | 1.4.1 | `el9.aarch64` | pigsty | 17.3 KiB | [pglogical_ticker_13-1.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglogical_ticker_13-1.4.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-pglogical-ticker` | 1.4.1 | `d12.x86_64` | pgdg | 20.6 KiB | [postgresql-13-pglogical-ticker_1.4.1-8.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-13-pglogical-ticker_1.4.1-8.pgdg120+1_amd64.deb) |
| `postgresql-13-pglogical-ticker` | 1.4.1 | `d12.aarch64` | pgdg | 20.4 KiB | [postgresql-13-pglogical-ticker_1.4.1-8.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-13-pglogical-ticker_1.4.1-8.pgdg120+1_arm64.deb) |
| `postgresql-13-pglogical-ticker` | 1.4.1 | `d13.x86_64` | pgdg | 20.7 KiB | [postgresql-13-pglogical-ticker_1.4.1-8.pgdg130+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-13-pglogical-ticker_1.4.1-8.pgdg130+2_amd64.deb) |
| `postgresql-13-pglogical-ticker` | 1.4.1 | `d13.aarch64` | pgdg | 20.5 KiB | [postgresql-13-pglogical-ticker_1.4.1-8.pgdg130+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-13-pglogical-ticker_1.4.1-8.pgdg130+2_arm64.deb) |
| `postgresql-13-pglogical-ticker` | 1.4.1 | `u22.x86_64` | pgdg | 20.7 KiB | [postgresql-13-pglogical-ticker_1.4.1-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-13-pglogical-ticker_1.4.1-8.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pglogical-ticker` | 1.4.1 | `u22.aarch64` | pgdg | 20.5 KiB | [postgresql-13-pglogical-ticker_1.4.1-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-13-pglogical-ticker_1.4.1-8.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pglogical-ticker` | 1.4.1 | `u24.x86_64` | pgdg | 20.2 KiB | [postgresql-13-pglogical-ticker_1.4.1-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-13-pglogical-ticker_1.4.1-8.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pglogical-ticker` | 1.4.1 | `u24.aarch64` | pgdg | 19.9 KiB | [postgresql-13-pglogical-ticker_1.4.1-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical-ticker/postgresql-13-pglogical-ticker_1.4.1-8.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/enova/pglogical_ticker" title="Repository" icon="github" subtitle="github.com/enova/pglogical_ticker" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pglogical_ticker-1.4.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pglogical_ticker; # get pglogical_ticker source code
pig build dep pglogical_ticker; # install build dependencies
pig build pkg pglogical_ticker; # build extension rpm or deb
pig build ext pglogical_ticker; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pglogical_ticker; # install by extension name, for the current active PG version
pig ext install pglogical_ticker; # install via package alias, for the active PG version
pig ext install pglogical_ticker -v 17;   # install for PG 17
pig ext install pglogical_ticker -v 16;   # install for PG 16
pig ext install pglogical_ticker -v 15;   # install for PG 15
pig ext install pglogical_ticker -v 14;   # install for PG 14
pig ext install pglogical_ticker -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pglogical_ticker CASCADE SCHEMA pglogical_ticker;
```

