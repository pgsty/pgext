---
title: "pg_math"
linkTitle: "pg_math"
description: "GSL statistical functions for postgresql"
weight: 4770
categories: ["FUNC"]
width: full
---

GSL statistical functions for postgresql


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4770** | {{< badge content="pg_math" link="https://github.com/chanukyasds/pg_math" >}} | {{< ext "pg_math" >}} | `1.0` | {{< category "FUNC" >}} | {{< license "GPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "permuteseq" >}} {{< ext "pg_hashids" >}} {{< ext "sequential_uuids" >}} {{< ext "topn" >}} {{< ext "quantile" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_math" >}} | `1.0` | {{< bg "18" "pg_math_18*" "green" >}} {{< bg "17" "pg_math_17*" "green" >}} {{< bg "16" "pg_math_16*" "green" >}} {{< bg "15" "pg_math_15*" "green" >}} {{< bg "14" "pg_math_14*" "green" >}} {{< bg "13" "pg_math_13*" "green" >}} | `pg_math_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_math" >}} | `1.0` | {{< bg "18" "postgresql-18-pg-math" "green" >}} {{< bg "17" "postgresql-17-pg-math" "green" >}} {{< bg "16" "postgresql-16-pg-math" "green" >}} {{< bg "15" "postgresql-15-pg-math" "green" >}} {{< bg "14" "postgresql-14-pg-math" "green" >}} {{< bg "13" "postgresql-13-pg-math" "green" >}} | `postgresql-$v-pg-math` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pg_math_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "pg_math_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_math_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_math_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_math_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_math_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pg_math_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "pg_math_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_math_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_math_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_math_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_math_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pg_math_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "pg_math_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_math_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_math_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_math_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_math_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pg_math_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "pg_math_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_math_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_math_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_math_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_math_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    |      {{< bg "MISS" "pg_math_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_math_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_math_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_math_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_math_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_math_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "pg_math_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_math_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_math_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_math_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_math_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_math_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-math : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-math : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-math : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-math : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-math : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-math : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-math : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-math : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-math : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-math : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-math : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-math : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-math : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-math : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-math : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-math : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-math : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-math : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-math : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-math : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-math : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-math : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-math : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-math : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-math : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_math_17` | 1.0 | `el8.x86_64` | pigsty | 28.0 KiB | [pg_math_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_math_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_math_17` | 1.0 | `el8.aarch64` | pigsty | 26.5 KiB | [pg_math_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_math_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_math_17` | 1.0 | `el9.x86_64` | pigsty | 28.6 KiB | [pg_math_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_math_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_math_17` | 1.0 | `el9.aarch64` | pigsty | 27.4 KiB | [pg_math_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_math_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pg-math` | 1.0 | `d12.x86_64` | pigsty | 61.1 KiB | [postgresql-17-pg-math_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-math/postgresql-17-pg-math_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-math` | 1.0 | `d12.aarch64` | pigsty | 60.1 KiB | [postgresql-17-pg-math_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-math/postgresql-17-pg-math_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-math` | 1.0 | `u22.x86_64` | pigsty | 69.0 KiB | [postgresql-17-pg-math_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-math/postgresql-17-pg-math_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-math` | 1.0 | `u22.aarch64` | pigsty | 67.8 KiB | [postgresql-17-pg-math_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-math/postgresql-17-pg-math_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-math` | 1.0 | `u24.x86_64` | pigsty | 64.8 KiB | [postgresql-17-pg-math_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-math/postgresql-17-pg-math_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-math` | 1.0 | `u24.aarch64` | pigsty | 63.4 KiB | [postgresql-17-pg-math_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-math/postgresql-17-pg-math_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_math_16` | 1.0 | `el8.x86_64` | pigsty | 28.0 KiB | [pg_math_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_math_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_math_16` | 1.0 | `el8.aarch64` | pigsty | 26.5 KiB | [pg_math_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_math_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_math_16` | 1.0 | `el9.x86_64` | pigsty | 29.0 KiB | [pg_math_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_math_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_math_16` | 1.0 | `el9.aarch64` | pigsty | 27.4 KiB | [pg_math_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_math_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-math` | 1.0 | `d12.x86_64` | pigsty | 61.1 KiB | [postgresql-16-pg-math_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-math/postgresql-16-pg-math_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-math` | 1.0 | `d12.aarch64` | pigsty | 60.1 KiB | [postgresql-16-pg-math_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-math/postgresql-16-pg-math_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-math` | 1.0 | `u22.x86_64` | pigsty | 69.0 KiB | [postgresql-16-pg-math_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-math/postgresql-16-pg-math_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-math` | 1.0 | `u22.aarch64` | pigsty | 67.8 KiB | [postgresql-16-pg-math_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-math/postgresql-16-pg-math_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-math` | 1.0 | `u24.x86_64` | pigsty | 64.8 KiB | [postgresql-16-pg-math_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-math/postgresql-16-pg-math_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-math` | 1.0 | `u24.aarch64` | pigsty | 63.3 KiB | [postgresql-16-pg-math_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-math/postgresql-16-pg-math_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_math_15` | 1.0 | `el8.x86_64` | pigsty | 28.0 KiB | [pg_math_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_math_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_math_15` | 1.0 | `el8.aarch64` | pigsty | 26.5 KiB | [pg_math_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_math_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_math_15` | 1.0 | `el9.x86_64` | pigsty | 28.6 KiB | [pg_math_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_math_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_math_15` | 1.0 | `el9.aarch64` | pigsty | 27.4 KiB | [pg_math_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_math_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-math` | 1.0 | `d12.x86_64` | pigsty | 61.0 KiB | [postgresql-15-pg-math_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-math/postgresql-15-pg-math_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-math` | 1.0 | `d12.aarch64` | pigsty | 60.2 KiB | [postgresql-15-pg-math_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-math/postgresql-15-pg-math_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-math` | 1.0 | `u22.x86_64` | pigsty | 69.0 KiB | [postgresql-15-pg-math_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-math/postgresql-15-pg-math_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-math` | 1.0 | `u22.aarch64` | pigsty | 67.8 KiB | [postgresql-15-pg-math_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-math/postgresql-15-pg-math_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-math` | 1.0 | `u24.x86_64` | pigsty | 64.8 KiB | [postgresql-15-pg-math_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-math/postgresql-15-pg-math_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-math` | 1.0 | `u24.aarch64` | pigsty | 63.4 KiB | [postgresql-15-pg-math_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-math/postgresql-15-pg-math_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_math_14` | 1.0 | `el8.x86_64` | pigsty | 28.0 KiB | [pg_math_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_math_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_math_14` | 1.0 | `el8.aarch64` | pigsty | 26.4 KiB | [pg_math_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_math_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_math_14` | 1.0 | `el9.x86_64` | pigsty | 28.6 KiB | [pg_math_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_math_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_math_14` | 1.0 | `el9.aarch64` | pigsty | 27.4 KiB | [pg_math_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_math_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-math` | 1.0 | `d12.x86_64` | pigsty | 60.9 KiB | [postgresql-14-pg-math_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-math/postgresql-14-pg-math_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-math` | 1.0 | `d12.aarch64` | pigsty | 59.9 KiB | [postgresql-14-pg-math_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-math/postgresql-14-pg-math_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-math` | 1.0 | `u22.x86_64` | pigsty | 68.9 KiB | [postgresql-14-pg-math_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-math/postgresql-14-pg-math_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-math` | 1.0 | `u22.aarch64` | pigsty | 67.6 KiB | [postgresql-14-pg-math_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-math/postgresql-14-pg-math_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-math` | 1.0 | `u24.x86_64` | pigsty | 64.8 KiB | [postgresql-14-pg-math_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-math/postgresql-14-pg-math_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-math` | 1.0 | `u24.aarch64` | pigsty | 63.5 KiB | [postgresql-14-pg-math_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-math/postgresql-14-pg-math_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_math_13` | 1.0 | `el8.x86_64` | pigsty | 27.0 KiB | [pg_math_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_math_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_math_13` | 1.0 | `el8.aarch64` | pigsty | 26.4 KiB | [pg_math_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_math_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_math_13` | 1.0 | `el9.x86_64` | pigsty | 28.9 KiB | [pg_math_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_math_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_math_13` | 1.0 | `el9.aarch64` | pigsty | 27.4 KiB | [pg_math_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_math_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-pg-math` | 1.0 | `d12.x86_64` | pigsty | 61.0 KiB | [postgresql-13-pg-math_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-math/postgresql-13-pg-math_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-math` | 1.0 | `d12.aarch64` | pigsty | 59.8 KiB | [postgresql-13-pg-math_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-math/postgresql-13-pg-math_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-math` | 1.0 | `u22.x86_64` | pigsty | 68.7 KiB | [postgresql-13-pg-math_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-math/postgresql-13-pg-math_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-math` | 1.0 | `u22.aarch64` | pigsty | 67.6 KiB | [postgresql-13-pg-math_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-math/postgresql-13-pg-math_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-math` | 1.0 | `u24.x86_64` | pigsty | 64.6 KiB | [postgresql-13-pg-math_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-math/postgresql-13-pg-math_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-math` | 1.0 | `u24.aarch64` | pigsty | 63.3 KiB | [postgresql-13-pg-math_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-math/postgresql-13-pg-math_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/chanukyasds/pg_math" title="Repository" icon="github" subtitle="github.com/chanukyasds/pg_math" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_math-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_math; # get pg_math source code
pig build dep pg_math; # install build dependencies
pig build pkg pg_math; # build extension rpm or deb
pig build ext pg_math; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_math; # install by extension name, for the current active PG version
pig ext install pg_math; # install via package alias, for the active PG version
pig ext install pg_math -v 18;   # install for PG 18
pig ext install pg_math -v 17;   # install for PG 17
pig ext install pg_math -v 16;   # install for PG 16
pig ext install pg_math -v 15;   # install for PG 15
pig ext install pg_math -v 14;   # install for PG 14
pig ext install pg_math -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_math;
```

