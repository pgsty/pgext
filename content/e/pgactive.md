---
title: "pgactive"
linkTitle: "pgactive"
description: "Active-Active Replication Extension for PostgreSQL"
weight: 9550
categories: ["ETL"]
width: full
---

Active-Active Replication Extension for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9550** | {{< badge content="pgactive" link="https://github.com/aws/pgactive" >}} | {{< ext "pgactive" >}} | `2.1.7` | {{< category "ETL" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pglogical" >}} {{< ext "pg_failover_slots" >}} {{< ext "repmgr" >}} {{< ext "bgw_replstatus" >}} {{< ext "pglogical_origin" >}} {{< ext "pglogical_ticker" >}} {{< ext "pgl_ddl_deploy" >}} {{< ext "decoderbufs" >}} |

> [!Note] require libpgfeutils


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgactive" >}} | `2.1.7` | {{< bg "18" "pgactive_18*" "green" >}} {{< bg "17" "pgactive_17*" "green" >}} {{< bg "16" "pgactive_16*" "green" >}} {{< bg "15" "pgactive_15*" "green" >}} {{< bg "14" "pgactive_14*" "green" >}} {{< bg "13" "pgactive_13*" "green" >}} | `pgactive_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgactive" >}} | `2.1.7` | {{< bg "18" "postgresql-18-pgactive" "red" >}} {{< bg "17" "postgresql-17-pgactive" "green" >}} {{< bg "16" "postgresql-16-pgactive" "green" >}} {{< bg "15" "postgresql-15-pgactive" "green" >}} {{< bg "14" "postgresql-14-pgactive" "green" >}} {{< bg "13" "postgresql-13-pgactive" "green" >}} | `postgresql-$v-pgactive` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 2.1.7" "pgactive_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_14 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_13 : AVAIL 3" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 2.1.7" "pgactive_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_14 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_13 : AVAIL 3" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 2.1.7" "pgactive_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_14 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_13 : AVAIL 2" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 2.1.7" "pgactive_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_14 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    | {{< bg "PIGSTY 2.1.7" "pgactive_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_13 : AVAIL 1" "green" >}} |
|    `el10.aarch64`    | {{< bg "PIGSTY 2.1.7" "pgactive_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.7" "pgactive_13 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pgactive : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.1.6" "postgresql-17-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-16-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-15-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-14-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-13-pgactive : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pgactive : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.1.6" "postgresql-17-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-16-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-15-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-14-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-13-pgactive : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pgactive : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgactive : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgactive : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgactive : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgactive : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pgactive : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pgactive : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgactive : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgactive : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgactive : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgactive : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pgactive : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pgactive : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.1.6" "postgresql-17-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-16-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-15-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-14-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-13-pgactive : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pgactive : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.1.6" "postgresql-17-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-16-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-15-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-14-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-13-pgactive : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pgactive : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.1.6" "postgresql-17-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-16-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-15-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-14-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-13-pgactive : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pgactive : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.1.6" "postgresql-17-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-16-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-15-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-14-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.6" "postgresql-13-pgactive : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgactive_18` | 2.1.7 | `el8.x86_64` | pigsty | 376.1 KiB | [pgactive_18-2.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_18-2.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pgactive_18` | 2.1.6 | `el8.x86_64` | pigsty | 347.5 KiB | [pgactive_18-2.1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_18-2.1.6-1PIGSTY.el8.x86_64.rpm) |
| `pgactive_18` | 2.1.7 | `el8.aarch64` | pigsty | 360.5 KiB | [pgactive_18-2.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_18-2.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pgactive_18` | 2.1.6 | `el8.aarch64` | pigsty | 329.6 KiB | [pgactive_18-2.1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_18-2.1.6-1PIGSTY.el8.aarch64.rpm) |
| `pgactive_18` | 2.1.7 | `el9.x86_64` | pigsty | 346.6 KiB | [pgactive_18-2.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_18-2.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pgactive_18` | 2.1.6 | `el9.x86_64` | pigsty | 338.1 KiB | [pgactive_18-2.1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_18-2.1.6-1PIGSTY.el9.x86_64.rpm) |
| `pgactive_18` | 2.1.7 | `el9.aarch64` | pigsty | 338.0 KiB | [pgactive_18-2.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgactive_18-2.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pgactive_18` | 2.1.7 | `el10.x86_64` | pigsty | 350.7 KiB | [pgactive_18-2.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgactive_18-2.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pgactive_18` | 2.1.7 | `el10.aarch64` | pigsty | 343.5 KiB | [pgactive_18-2.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgactive_18-2.1.7-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgactive_17` | 2.1.7 | `el8.x86_64` | pigsty | 367.9 KiB | [pgactive_17-2.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_17-2.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pgactive_17` | 2.1.6 | `el8.x86_64` | pigsty | 339.6 KiB | [pgactive_17-2.1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_17-2.1.6-1PIGSTY.el8.x86_64.rpm) |
| `pgactive_17` | 2.1.5 | `el8.x86_64` | pigsty | 332.0 KiB | [pgactive_17-2.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_17-2.1.5-1PIGSTY.el8.x86_64.rpm) |
| `pgactive_17` | 2.1.7 | `el8.aarch64` | pigsty | 352.8 KiB | [pgactive_17-2.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_17-2.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pgactive_17` | 2.1.6 | `el8.aarch64` | pigsty | 322.3 KiB | [pgactive_17-2.1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_17-2.1.6-1PIGSTY.el8.aarch64.rpm) |
| `pgactive_17` | 2.1.5 | `el8.aarch64` | pigsty | 314.9 KiB | [pgactive_17-2.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_17-2.1.5-1PIGSTY.el8.aarch64.rpm) |
| `pgactive_17` | 2.1.7 | `el9.x86_64` | pigsty | 339.2 KiB | [pgactive_17-2.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_17-2.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pgactive_17` | 2.1.6 | `el9.x86_64` | pigsty | 331.3 KiB | [pgactive_17-2.1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_17-2.1.6-1PIGSTY.el9.x86_64.rpm) |
| `pgactive_17` | 2.1.5 | `el9.x86_64` | pigsty | 322.0 KiB | [pgactive_17-2.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_17-2.1.5-1PIGSTY.el9.x86_64.rpm) |
| `pgactive_17` | 2.1.7 | `el9.aarch64` | pigsty | 330.8 KiB | [pgactive_17-2.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgactive_17-2.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pgactive_17` | 2.1.6 | `el9.aarch64` | pigsty | 322.6 KiB | [pgactive_17-2.1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgactive_17-2.1.6-1PIGSTY.el9.aarch64.rpm) |
| `pgactive_17` | 2.1.5 | `el9.aarch64` | pigsty | 313.6 KiB | [pgactive_17-2.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgactive_17-2.1.5-1PIGSTY.el9.aarch64.rpm) |
| `pgactive_17` | 2.1.7 | `el10.x86_64` | pigsty | 343.0 KiB | [pgactive_17-2.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgactive_17-2.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pgactive_17` | 2.1.7 | `el10.aarch64` | pigsty | 335.9 KiB | [pgactive_17-2.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgactive_17-2.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgactive` | 2.1.6 | `d12.x86_64` | pigsty | 586.4 KiB | [postgresql-17-pgactive_2.1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-17-pgactive_2.1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgactive` | 2.1.6 | `d12.aarch64` | pigsty | 560.4 KiB | [postgresql-17-pgactive_2.1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-17-pgactive_2.1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgactive` | 2.1.6 | `u22.x86_64` | pigsty | 711.6 KiB | [postgresql-17-pgactive_2.1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-17-pgactive_2.1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgactive` | 2.1.6 | `u22.aarch64` | pigsty | 702.7 KiB | [postgresql-17-pgactive_2.1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-17-pgactive_2.1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgactive` | 2.1.6 | `u24.x86_64` | pigsty | 618.6 KiB | [postgresql-17-pgactive_2.1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-17-pgactive_2.1.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgactive` | 2.1.6 | `u24.aarch64` | pigsty | 612.8 KiB | [postgresql-17-pgactive_2.1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-17-pgactive_2.1.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgactive_16` | 2.1.7 | `el8.x86_64` | pigsty | 361.5 KiB | [pgactive_16-2.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_16-2.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pgactive_16` | 2.1.6 | `el8.x86_64` | pigsty | 333.7 KiB | [pgactive_16-2.1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_16-2.1.6-1PIGSTY.el8.x86_64.rpm) |
| `pgactive_16` | 2.1.5 | `el8.x86_64` | pigsty | 326.6 KiB | [pgactive_16-2.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_16-2.1.5-1PIGSTY.el8.x86_64.rpm) |
| `pgactive_16` | 2.1.7 | `el8.aarch64` | pigsty | 346.2 KiB | [pgactive_16-2.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_16-2.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pgactive_16` | 2.1.6 | `el8.aarch64` | pigsty | 316.5 KiB | [pgactive_16-2.1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_16-2.1.6-1PIGSTY.el8.aarch64.rpm) |
| `pgactive_16` | 2.1.5 | `el8.aarch64` | pigsty | 309.4 KiB | [pgactive_16-2.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_16-2.1.5-1PIGSTY.el8.aarch64.rpm) |
| `pgactive_16` | 2.1.7 | `el9.x86_64` | pigsty | 333.9 KiB | [pgactive_16-2.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_16-2.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pgactive_16` | 2.1.6 | `el9.x86_64` | pigsty | 326.0 KiB | [pgactive_16-2.1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_16-2.1.6-1PIGSTY.el9.x86_64.rpm) |
| `pgactive_16` | 2.1.5 | `el9.x86_64` | pigsty | 316.2 KiB | [pgactive_16-2.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_16-2.1.5-1PIGSTY.el9.x86_64.rpm) |
| `pgactive_16` | 2.1.7 | `el9.aarch64` | pigsty | 325.5 KiB | [pgactive_16-2.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgactive_16-2.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pgactive_16` | 2.1.6 | `el9.aarch64` | pigsty | 317.0 KiB | [pgactive_16-2.1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgactive_16-2.1.6-1PIGSTY.el9.aarch64.rpm) |
| `pgactive_16` | 2.1.5 | `el9.aarch64` | pigsty | 307.7 KiB | [pgactive_16-2.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgactive_16-2.1.5-1PIGSTY.el9.aarch64.rpm) |
| `pgactive_16` | 2.1.7 | `el10.x86_64` | pigsty | 338.2 KiB | [pgactive_16-2.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgactive_16-2.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pgactive_16` | 2.1.7 | `el10.aarch64` | pigsty | 329.6 KiB | [pgactive_16-2.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgactive_16-2.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgactive` | 2.1.6 | `d12.x86_64` | pigsty | 581.2 KiB | [postgresql-16-pgactive_2.1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-16-pgactive_2.1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgactive` | 2.1.6 | `d12.aarch64` | pigsty | 556.5 KiB | [postgresql-16-pgactive_2.1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-16-pgactive_2.1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgactive` | 2.1.6 | `u22.x86_64` | pigsty | 700.8 KiB | [postgresql-16-pgactive_2.1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-16-pgactive_2.1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgactive` | 2.1.6 | `u22.aarch64` | pigsty | 692.2 KiB | [postgresql-16-pgactive_2.1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-16-pgactive_2.1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgactive` | 2.1.6 | `u24.x86_64` | pigsty | 613.8 KiB | [postgresql-16-pgactive_2.1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-16-pgactive_2.1.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgactive` | 2.1.6 | `u24.aarch64` | pigsty | 607.3 KiB | [postgresql-16-pgactive_2.1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-16-pgactive_2.1.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgactive_15` | 2.1.7 | `el8.x86_64` | pigsty | 353.2 KiB | [pgactive_15-2.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_15-2.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pgactive_15` | 2.1.6 | `el8.x86_64` | pigsty | 325.9 KiB | [pgactive_15-2.1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_15-2.1.6-1PIGSTY.el8.x86_64.rpm) |
| `pgactive_15` | 2.1.5 | `el8.x86_64` | pigsty | 319.0 KiB | [pgactive_15-2.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_15-2.1.5-1PIGSTY.el8.x86_64.rpm) |
| `pgactive_15` | 2.1.7 | `el8.aarch64` | pigsty | 338.5 KiB | [pgactive_15-2.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_15-2.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pgactive_15` | 2.1.6 | `el8.aarch64` | pigsty | 308.7 KiB | [pgactive_15-2.1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_15-2.1.6-1PIGSTY.el8.aarch64.rpm) |
| `pgactive_15` | 2.1.5 | `el8.aarch64` | pigsty | 301.6 KiB | [pgactive_15-2.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_15-2.1.5-1PIGSTY.el8.aarch64.rpm) |
| `pgactive_15` | 2.1.7 | `el9.x86_64` | pigsty | 330.2 KiB | [pgactive_15-2.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_15-2.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pgactive_15` | 2.1.6 | `el9.x86_64` | pigsty | 321.9 KiB | [pgactive_15-2.1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_15-2.1.6-1PIGSTY.el9.x86_64.rpm) |
| `pgactive_15` | 2.1.5 | `el9.x86_64` | pigsty | 312.7 KiB | [pgactive_15-2.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_15-2.1.5-1PIGSTY.el9.x86_64.rpm) |
| `pgactive_15` | 2.1.7 | `el9.aarch64` | pigsty | 322.7 KiB | [pgactive_15-2.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgactive_15-2.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pgactive_15` | 2.1.6 | `el9.aarch64` | pigsty | 313.0 KiB | [pgactive_15-2.1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgactive_15-2.1.6-1PIGSTY.el9.aarch64.rpm) |
| `pgactive_15` | 2.1.5 | `el9.aarch64` | pigsty | 304.9 KiB | [pgactive_15-2.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgactive_15-2.1.5-1PIGSTY.el9.aarch64.rpm) |
| `pgactive_15` | 2.1.7 | `el10.x86_64` | pigsty | 333.6 KiB | [pgactive_15-2.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgactive_15-2.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pgactive_15` | 2.1.7 | `el10.aarch64` | pigsty | 327.7 KiB | [pgactive_15-2.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgactive_15-2.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgactive` | 2.1.6 | `d12.x86_64` | pigsty | 577.4 KiB | [postgresql-15-pgactive_2.1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-15-pgactive_2.1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgactive` | 2.1.6 | `d12.aarch64` | pigsty | 550.9 KiB | [postgresql-15-pgactive_2.1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-15-pgactive_2.1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgactive` | 2.1.6 | `u22.x86_64` | pigsty | 698.1 KiB | [postgresql-15-pgactive_2.1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-15-pgactive_2.1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgactive` | 2.1.6 | `u22.aarch64` | pigsty | 690.5 KiB | [postgresql-15-pgactive_2.1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-15-pgactive_2.1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgactive` | 2.1.6 | `u24.x86_64` | pigsty | 610.0 KiB | [postgresql-15-pgactive_2.1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-15-pgactive_2.1.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgactive` | 2.1.6 | `u24.aarch64` | pigsty | 605.7 KiB | [postgresql-15-pgactive_2.1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-15-pgactive_2.1.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgactive_14` | 2.1.7 | `el8.x86_64` | pigsty | 352.8 KiB | [pgactive_14-2.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_14-2.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pgactive_14` | 2.1.6 | `el8.x86_64` | pigsty | 325.3 KiB | [pgactive_14-2.1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_14-2.1.6-1PIGSTY.el8.x86_64.rpm) |
| `pgactive_14` | 2.1.5 | `el8.x86_64` | pigsty | 318.9 KiB | [pgactive_14-2.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_14-2.1.5-1PIGSTY.el8.x86_64.rpm) |
| `pgactive_14` | 2.1.7 | `el8.aarch64` | pigsty | 339.0 KiB | [pgactive_14-2.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_14-2.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pgactive_14` | 2.1.6 | `el8.aarch64` | pigsty | 309.2 KiB | [pgactive_14-2.1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_14-2.1.6-1PIGSTY.el8.aarch64.rpm) |
| `pgactive_14` | 2.1.5 | `el8.aarch64` | pigsty | 302.9 KiB | [pgactive_14-2.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_14-2.1.5-1PIGSTY.el8.aarch64.rpm) |
| `pgactive_14` | 2.1.7 | `el9.x86_64` | pigsty | 332.9 KiB | [pgactive_14-2.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_14-2.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pgactive_14` | 2.1.6 | `el9.x86_64` | pigsty | 324.3 KiB | [pgactive_14-2.1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_14-2.1.6-1PIGSTY.el9.x86_64.rpm) |
| `pgactive_14` | 2.1.5 | `el9.x86_64` | pigsty | 315.9 KiB | [pgactive_14-2.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_14-2.1.5-1PIGSTY.el9.x86_64.rpm) |
| `pgactive_14` | 2.1.7 | `el9.aarch64` | pigsty | 325.3 KiB | [pgactive_14-2.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgactive_14-2.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pgactive_14` | 2.1.6 | `el9.aarch64` | pigsty | 315.1 KiB | [pgactive_14-2.1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgactive_14-2.1.6-1PIGSTY.el9.aarch64.rpm) |
| `pgactive_14` | 2.1.5 | `el9.aarch64` | pigsty | 307.5 KiB | [pgactive_14-2.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgactive_14-2.1.5-1PIGSTY.el9.aarch64.rpm) |
| `pgactive_14` | 2.1.7 | `el10.x86_64` | pigsty | 338.2 KiB | [pgactive_14-2.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgactive_14-2.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pgactive_14` | 2.1.7 | `el10.aarch64` | pigsty | 329.0 KiB | [pgactive_14-2.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgactive_14-2.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgactive` | 2.1.6 | `d12.x86_64` | pigsty | 580.8 KiB | [postgresql-14-pgactive_2.1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-14-pgactive_2.1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgactive` | 2.1.6 | `d12.aarch64` | pigsty | 555.2 KiB | [postgresql-14-pgactive_2.1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-14-pgactive_2.1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgactive` | 2.1.6 | `u22.x86_64` | pigsty | 701.1 KiB | [postgresql-14-pgactive_2.1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-14-pgactive_2.1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgactive` | 2.1.6 | `u22.aarch64` | pigsty | 692.8 KiB | [postgresql-14-pgactive_2.1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-14-pgactive_2.1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgactive` | 2.1.6 | `u24.x86_64` | pigsty | 613.2 KiB | [postgresql-14-pgactive_2.1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-14-pgactive_2.1.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgactive` | 2.1.6 | `u24.aarch64` | pigsty | 606.9 KiB | [postgresql-14-pgactive_2.1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-14-pgactive_2.1.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgactive_13` | 2.1.7 | `el8.x86_64` | pigsty | 345.8 KiB | [pgactive_13-2.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_13-2.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pgactive_13` | 2.1.6 | `el8.x86_64` | pigsty | 319.0 KiB | [pgactive_13-2.1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_13-2.1.6-1PIGSTY.el8.x86_64.rpm) |
| `pgactive_13` | 2.1.5 | `el8.x86_64` | pigsty | 312.9 KiB | [pgactive_13-2.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_13-2.1.5-1PIGSTY.el8.x86_64.rpm) |
| `pgactive_13` | 2.1.7 | `el8.aarch64` | pigsty | 334.4 KiB | [pgactive_13-2.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_13-2.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pgactive_13` | 2.1.6 | `el8.aarch64` | pigsty | 305.6 KiB | [pgactive_13-2.1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_13-2.1.6-1PIGSTY.el8.aarch64.rpm) |
| `pgactive_13` | 2.1.5 | `el8.aarch64` | pigsty | 299.6 KiB | [pgactive_13-2.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_13-2.1.5-1PIGSTY.el8.aarch64.rpm) |
| `pgactive_13` | 2.1.7 | `el9.x86_64` | pigsty | 328.9 KiB | [pgactive_13-2.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_13-2.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pgactive_13` | 2.1.6 | `el9.x86_64` | pigsty | 320.0 KiB | [pgactive_13-2.1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_13-2.1.6-1PIGSTY.el9.x86_64.rpm) |
| `pgactive_13` | 2.1.7 | `el9.aarch64` | pigsty | 321.7 KiB | [pgactive_13-2.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgactive_13-2.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pgactive_13` | 2.1.7 | `el10.x86_64` | pigsty | 334.1 KiB | [pgactive_13-2.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgactive_13-2.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pgactive_13` | 2.1.7 | `el10.aarch64` | pigsty | 325.2 KiB | [pgactive_13-2.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgactive_13-2.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pgactive` | 2.1.6 | `d12.x86_64` | pigsty | 576.6 KiB | [postgresql-13-pgactive_2.1.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-13-pgactive_2.1.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgactive` | 2.1.6 | `d12.aarch64` | pigsty | 549.9 KiB | [postgresql-13-pgactive_2.1.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-13-pgactive_2.1.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgactive` | 2.1.6 | `u22.x86_64` | pigsty | 694.5 KiB | [postgresql-13-pgactive_2.1.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-13-pgactive_2.1.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgactive` | 2.1.6 | `u22.aarch64` | pigsty | 686.3 KiB | [postgresql-13-pgactive_2.1.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-13-pgactive_2.1.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgactive` | 2.1.6 | `u24.x86_64` | pigsty | 607.9 KiB | [postgresql-13-pgactive_2.1.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-13-pgactive_2.1.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgactive` | 2.1.6 | `u24.aarch64` | pigsty | 601.7 KiB | [postgresql-13-pgactive_2.1.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-13-pgactive_2.1.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/aws/pgactive" title="Repository" icon="github" subtitle="github.com/aws/pgactive" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgactive-2.1.7.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgactive; # get pgactive source code
pig build dep pgactive; # install build dependencies
pig build pkg pgactive; # build extension rpm or deb
pig build ext pgactive; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgactive; # install by extension name, for the current active PG version
pig ext install pgactive; # install via package alias, for the active PG version
pig ext install pgactive -v 18;   # install for PG 18
pig ext install pgactive -v 17;   # install for PG 17
pig ext install pgactive -v 16;   # install for PG 16
pig ext install pgactive -v 15;   # install for PG 15
pig ext install pgactive -v 14;   # install for PG 14
pig ext install pgactive -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgactive CASCADE SCHEMA pg_catalog;
```

