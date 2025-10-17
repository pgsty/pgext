---
title: "random"
linkTitle: "random"
description: "random data generator"
weight: 4780
categories: ["Func"]
width: full
---

random data generator

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4780** | {{< badge content="random" link="https://github.com/tvondra/random" >}} | {{< ext "random" "pg_random" >}} | `2.0.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "permuteseq" >}} {{< ext "tsm_system_rows" >}} {{< ext "tsm_system_time" >}} {{< ext "pg_idkit" >}} {{< ext "sequential_uuids" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/random" >}} | `2.0.0` | {{< badge content="18" color="red" alt="pg_random_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_random_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/random" >}} | `2.0.0` | {{< badge content="18" color="red" alt="postgresql-18-random" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-random` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_random_18" >}}     | {{< pkg "pg_random_17" "2.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_random_17-2.0.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_random_16" "2.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_random_16-2.0.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_random_15" "2.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_random_15-2.0.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_random_14" "2.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_random_14-2.0.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_random_18" >}}     | {{< pkg "pg_random_17" "2.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_random_17-2.0.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_random_16" "2.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_random_16-2.0.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_random_15" "2.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_random_15-2.0.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_random_14" "2.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_random_14-2.0.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_random_18" >}}     | {{< pkg "pg_random_17" "2.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_random_17-2.0.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_random_16" "2.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_random_16-2.0.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_random_15" "2.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_random_15-2.0.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_random_14" "2.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_random_14-2.0.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_random_18" >}}     | {{< pkg "pg_random_17" "2.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_random_17-2.0.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_random_16" "2.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_random_16-2.0.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_random_15" "2.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_random_15-2.0.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_random_14" "2.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_random_14-2.0.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-random" >}}     | {{< pkg "postgresql-17-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-random" >}}     | {{< pkg "postgresql-17-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-random" >}}     | {{< pkg "postgresql-17-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-random" >}}     | {{< pkg "postgresql-17-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-random" >}}     | {{< pkg "postgresql-17-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-random" >}}     | {{< pkg "postgresql-17-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-random" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_random_17` | 2.0.0 | `el8.x86_64` | pigsty | 16.7 KiB | [pg_random_17-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_random_17-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_random_17` | 2.0.0 | `el8.aarch64` | pigsty | 16.4 KiB | [pg_random_17-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_random_17-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_random_17` | 2.0.0 | `el9.aarch64` | pigsty | 16.8 KiB | [pg_random_17-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_random_17-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_random_17` | 2.0.0 | `el9.x86_64` | pigsty | 17.3 KiB | [pg_random_17-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_random_17-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-random` | 2.0.0 | `d12.x86_64` | pigsty | 20.3 KiB | [postgresql-17-random_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-random` | 2.0.0 | `d12.aarch64` | pigsty | 20.0 KiB | [postgresql-17-random_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-random` | 2.0.0 | `u22.x86_64` | pigsty | 21.9 KiB | [postgresql-17-random_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-random` | 2.0.0 | `u22.aarch64` | pigsty | 21.4 KiB | [postgresql-17-random_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-random` | 2.0.0 | `u24.x86_64` | pigsty | 21.4 KiB | [postgresql-17-random_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-random` | 2.0.0 | `u24.aarch64` | pigsty | 21.1 KiB | [postgresql-17-random_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_random_16` | 2.0.0 | `el8.x86_64` | pigsty | 16.7 KiB | [pg_random_16-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_random_16-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_random_16` | 2.0.0 | `el8.aarch64` | pigsty | 16.4 KiB | [pg_random_16-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_random_16-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_random_16` | 2.0.0 | `el9.x86_64` | pigsty | 17.3 KiB | [pg_random_16-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_random_16-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_random_16` | 2.0.0 | `el9.aarch64` | pigsty | 16.8 KiB | [pg_random_16-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_random_16-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-random` | 2.0.0 | `d12.x86_64` | pigsty | 20.3 KiB | [postgresql-16-random_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-random` | 2.0.0 | `d12.aarch64` | pigsty | 20.0 KiB | [postgresql-16-random_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-random` | 2.0.0 | `u22.aarch64` | pigsty | 21.4 KiB | [postgresql-16-random_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-random` | 2.0.0 | `u22.x86_64` | pigsty | 21.9 KiB | [postgresql-16-random_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-random` | 2.0.0 | `u24.x86_64` | pigsty | 21.4 KiB | [postgresql-16-random_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-random` | 2.0.0 | `u24.aarch64` | pigsty | 21.1 KiB | [postgresql-16-random_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_random_15` | 2.0.0 | `el8.x86_64` | pigsty | 16.8 KiB | [pg_random_15-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_random_15-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_random_15` | 2.0.0 | `el8.aarch64` | pigsty | 16.5 KiB | [pg_random_15-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_random_15-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_random_15` | 2.0.0 | `el9.x86_64` | pigsty | 17.3 KiB | [pg_random_15-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_random_15-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_random_15` | 2.0.0 | `el9.aarch64` | pigsty | 16.9 KiB | [pg_random_15-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_random_15-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-random` | 2.0.0 | `d12.aarch64` | pigsty | 20.1 KiB | [postgresql-15-random_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-random` | 2.0.0 | `d12.x86_64` | pigsty | 20.3 KiB | [postgresql-15-random_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-random` | 2.0.0 | `u22.aarch64` | pigsty | 21.5 KiB | [postgresql-15-random_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-random` | 2.0.0 | `u22.x86_64` | pigsty | 21.7 KiB | [postgresql-15-random_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-random` | 2.0.0 | `u24.x86_64` | pigsty | 21.2 KiB | [postgresql-15-random_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-random` | 2.0.0 | `u24.aarch64` | pigsty | 21.1 KiB | [postgresql-15-random_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_random_14` | 2.0.0 | `el8.x86_64` | pigsty | 16.8 KiB | [pg_random_14-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_random_14-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_random_14` | 2.0.0 | `el8.aarch64` | pigsty | 16.5 KiB | [pg_random_14-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_random_14-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_random_14` | 2.0.0 | `el9.x86_64` | pigsty | 17.3 KiB | [pg_random_14-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_random_14-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_random_14` | 2.0.0 | `el9.aarch64` | pigsty | 16.9 KiB | [pg_random_14-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_random_14-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-random` | 2.0.0 | `d12.x86_64` | pigsty | 20.3 KiB | [postgresql-14-random_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-random` | 2.0.0 | `d12.aarch64` | pigsty | 20.1 KiB | [postgresql-14-random_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-random` | 2.0.0 | `u22.x86_64` | pigsty | 21.7 KiB | [postgresql-14-random_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-random` | 2.0.0 | `u22.aarch64` | pigsty | 21.4 KiB | [postgresql-14-random_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-random` | 2.0.0 | `u24.x86_64` | pigsty | 21.2 KiB | [postgresql-14-random_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-random` | 2.0.0 | `u24.aarch64` | pigsty | 21.0 KiB | [postgresql-14-random_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_random_13` | 2.0.0 | `el8.aarch64` | pigsty | 16.5 KiB | [pg_random_13-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_random_13-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_random_13` | 2.0.0 | `el8.x86_64` | pigsty | 16.6 KiB | [pg_random_13-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_random_13-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_random_13` | 2.0.0 | `el9.aarch64` | pigsty | 16.9 KiB | [pg_random_13-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_random_13-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_random_13` | 2.0.0 | `el9.x86_64` | pigsty | 17.3 KiB | [pg_random_13-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_random_13-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-random` | 2.0.0 | `d12.aarch64` | pigsty | 19.9 KiB | [postgresql-13-random_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-13-random_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-random` | 2.0.0 | `d12.x86_64` | pigsty | 20.3 KiB | [postgresql-13-random_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-13-random_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-random` | 2.0.0 | `u22.aarch64` | pigsty | 21.4 KiB | [postgresql-13-random_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-13-random_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-random` | 2.0.0 | `u22.x86_64` | pigsty | 21.7 KiB | [postgresql-13-random_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-13-random_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-random` | 2.0.0 | `u24.aarch64` | pigsty | 20.8 KiB | [postgresql-13-random_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-13-random_2.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-random` | 2.0.0 | `u24.x86_64` | pigsty | 21.3 KiB | [postgresql-13-random_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-13-random_2.0.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/random" title="Repository" icon="github" subtitle="github.com/tvondra/random" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="random-2.0.0-dev.tar.gz" >}}
{{< /cards >}}


```bash
pig build get random; # get random source code
pig build dep random; # install build dependencies
pig build pkg random; # build extension rpm or deb
pig build ext random; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install random; # install by extension name, for the current active PG version
pig ext install pg_random; # install via package alias, for the active PG version
pig ext install random -v 18;   # install for PG 18
pig ext install random -v 17;   # install for PG 17
pig ext install random -v 16;   # install for PG 16
pig ext install random -v 15;   # install for PG 15
pig ext install random -v 14;   # install for PG 14
pig ext install random -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION random;
```

