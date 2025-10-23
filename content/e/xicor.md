---
title: "xicor"
linkTitle: "xicor"
description: "XI Correlation Coefficient in Postgres"
weight: 4670
categories: ["FUNC"]
width: full
---

XI Correlation Coefficient in Postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4670** | {{< badge content="xicor" link="https://github.com/Florents-Tselai/pgxicor" >}} | {{< ext "xicor" "pgxicor" >}} | `0.1.0` | {{< category "FUNC" >}} | {{< license "GPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dtr" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "permuteseq" >}} {{< ext "pg_hashids" >}} {{< ext "sequential_uuids" >}} {{< ext "topn" >}} {{< ext "quantile" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/xicor" >}} | `0.1.0` | {{< bg "18" "pgxicor_18*" "red" >}} {{< bg "17" "pgxicor_17*" "green" >}} {{< bg "16" "pgxicor_16*" "green" >}} {{< bg "15" "pgxicor_15*" "green" >}} {{< bg "14" "pgxicor_14*" "green" >}} | `pgxicor_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/xicor" >}} | `0.1.0` | {{< bg "18" "postgresql-18-pgxicor" "red" >}} {{< bg "17" "postgresql-17-pgxicor" "green" >}} {{< bg "16" "postgresql-16-pgxicor" "green" >}} {{< bg "15" "postgresql-15-pgxicor" "green" >}} {{< bg "14" "postgresql-14-pgxicor" "green" >}} | `postgresql-$v-pgxicor` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pgxicor_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "pgxicor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgxicor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgxicor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgxicor_14 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pgxicor_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "pgxicor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgxicor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgxicor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgxicor_14 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pgxicor_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "pgxicor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgxicor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgxicor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgxicor_14 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pgxicor_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "pgxicor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgxicor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgxicor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgxicor_14 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pgxicor : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgxicor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgxicor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgxicor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgxicor : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pgxicor : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgxicor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgxicor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgxicor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgxicor : AVAIL 1" "green" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pgxicor : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgxicor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgxicor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgxicor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgxicor : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pgxicor : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgxicor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgxicor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgxicor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgxicor : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pgxicor : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgxicor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgxicor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgxicor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgxicor : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pgxicor : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgxicor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgxicor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgxicor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgxicor : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgxicor_17` | 0.1.0 | `el8.x86_64` | pigsty | 26.0 KiB | [pgxicor_17-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgxicor_17-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgxicor_17` | 0.1.0 | `el8.aarch64` | pigsty | 26.0 KiB | [pgxicor_17-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgxicor_17-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgxicor_17` | 0.1.0 | `el9.x86_64` | pigsty | 25.8 KiB | [pgxicor_17-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgxicor_17-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgxicor_17` | 0.1.0 | `el9.aarch64` | pigsty | 25.5 KiB | [pgxicor_17-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgxicor_17-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pgxicor` | 0.1.0 | `d12.x86_64` | pigsty | 25.3 KiB | [postgresql-17-pgxicor_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgxicor/postgresql-17-pgxicor_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgxicor` | 0.1.0 | `d12.aarch64` | pigsty | 25.4 KiB | [postgresql-17-pgxicor_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgxicor/postgresql-17-pgxicor_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgxicor` | 0.1.0 | `u22.x86_64` | pigsty | 26.9 KiB | [postgresql-17-pgxicor_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgxicor/postgresql-17-pgxicor_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgxicor` | 0.1.0 | `u22.aarch64` | pigsty | 26.3 KiB | [postgresql-17-pgxicor_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgxicor/postgresql-17-pgxicor_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgxicor` | 0.1.0 | `u24.x86_64` | pigsty | 25.4 KiB | [postgresql-17-pgxicor_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgxicor/postgresql-17-pgxicor_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgxicor` | 0.1.0 | `u24.aarch64` | pigsty | 24.9 KiB | [postgresql-17-pgxicor_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgxicor/postgresql-17-pgxicor_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgxicor_16` | 0.1.0 | `el8.x86_64` | pigsty | 26.0 KiB | [pgxicor_16-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgxicor_16-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgxicor_16` | 0.1.0 | `el8.aarch64` | pigsty | 26.0 KiB | [pgxicor_16-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgxicor_16-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgxicor_16` | 0.1.0 | `el9.x86_64` | pigsty | 25.8 KiB | [pgxicor_16-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgxicor_16-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgxicor_16` | 0.1.0 | `el9.aarch64` | pigsty | 25.5 KiB | [pgxicor_16-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgxicor_16-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pgxicor` | 0.1.0 | `d12.x86_64` | pigsty | 25.3 KiB | [postgresql-16-pgxicor_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgxicor/postgresql-16-pgxicor_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgxicor` | 0.1.0 | `d12.aarch64` | pigsty | 25.4 KiB | [postgresql-16-pgxicor_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgxicor/postgresql-16-pgxicor_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgxicor` | 0.1.0 | `u22.x86_64` | pigsty | 26.9 KiB | [postgresql-16-pgxicor_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgxicor/postgresql-16-pgxicor_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgxicor` | 0.1.0 | `u22.aarch64` | pigsty | 26.3 KiB | [postgresql-16-pgxicor_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgxicor/postgresql-16-pgxicor_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgxicor` | 0.1.0 | `u24.x86_64` | pigsty | 25.4 KiB | [postgresql-16-pgxicor_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgxicor/postgresql-16-pgxicor_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgxicor` | 0.1.0 | `u24.aarch64` | pigsty | 24.9 KiB | [postgresql-16-pgxicor_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgxicor/postgresql-16-pgxicor_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgxicor_15` | 0.1.0 | `el8.x86_64` | pigsty | 26.1 KiB | [pgxicor_15-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgxicor_15-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgxicor_15` | 0.1.0 | `el8.aarch64` | pigsty | 26.1 KiB | [pgxicor_15-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgxicor_15-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgxicor_15` | 0.1.0 | `el9.x86_64` | pigsty | 26.0 KiB | [pgxicor_15-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgxicor_15-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgxicor_15` | 0.1.0 | `el9.aarch64` | pigsty | 25.9 KiB | [pgxicor_15-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgxicor_15-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pgxicor` | 0.1.0 | `d12.x86_64` | pigsty | 25.4 KiB | [postgresql-15-pgxicor_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgxicor/postgresql-15-pgxicor_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgxicor` | 0.1.0 | `d12.aarch64` | pigsty | 25.5 KiB | [postgresql-15-pgxicor_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgxicor/postgresql-15-pgxicor_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgxicor` | 0.1.0 | `u22.x86_64` | pigsty | 27.0 KiB | [postgresql-15-pgxicor_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgxicor/postgresql-15-pgxicor_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgxicor` | 0.1.0 | `u22.aarch64` | pigsty | 26.7 KiB | [postgresql-15-pgxicor_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgxicor/postgresql-15-pgxicor_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgxicor` | 0.1.0 | `u24.x86_64` | pigsty | 25.6 KiB | [postgresql-15-pgxicor_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgxicor/postgresql-15-pgxicor_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgxicor` | 0.1.0 | `u24.aarch64` | pigsty | 25.4 KiB | [postgresql-15-pgxicor_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgxicor/postgresql-15-pgxicor_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgxicor_14` | 0.1.0 | `el8.x86_64` | pigsty | 26.1 KiB | [pgxicor_14-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgxicor_14-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgxicor_14` | 0.1.0 | `el8.aarch64` | pigsty | 26.1 KiB | [pgxicor_14-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgxicor_14-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgxicor_14` | 0.1.0 | `el9.x86_64` | pigsty | 26.0 KiB | [pgxicor_14-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgxicor_14-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgxicor_14` | 0.1.0 | `el9.aarch64` | pigsty | 25.9 KiB | [pgxicor_14-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgxicor_14-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pgxicor` | 0.1.0 | `d12.x86_64` | pigsty | 25.4 KiB | [postgresql-14-pgxicor_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgxicor/postgresql-14-pgxicor_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgxicor` | 0.1.0 | `d12.aarch64` | pigsty | 25.4 KiB | [postgresql-14-pgxicor_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgxicor/postgresql-14-pgxicor_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgxicor` | 0.1.0 | `u22.x86_64` | pigsty | 27.0 KiB | [postgresql-14-pgxicor_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgxicor/postgresql-14-pgxicor_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgxicor` | 0.1.0 | `u22.aarch64` | pigsty | 26.6 KiB | [postgresql-14-pgxicor_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgxicor/postgresql-14-pgxicor_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgxicor` | 0.1.0 | `u24.x86_64` | pigsty | 25.5 KiB | [postgresql-14-pgxicor_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgxicor/postgresql-14-pgxicor_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgxicor` | 0.1.0 | `u24.aarch64` | pigsty | 25.3 KiB | [postgresql-14-pgxicor_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgxicor/postgresql-14-pgxicor_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Florents-Tselai/pgxicor" title="Repository" icon="github" subtitle="github.com/Florents-Tselai/pgxicor" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgxicor-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get xicor; # get xicor source code
pig build dep xicor; # install build dependencies
pig build pkg xicor; # build extension rpm or deb
pig build ext xicor; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install xicor; # install by extension name, for the current active PG version
pig ext install pgxicor; # install via package alias, for the active PG version
pig ext install xicor -v 18;   # install for PG 18
pig ext install xicor -v 17;   # install for PG 17
pig ext install xicor -v 16;   # install for PG 16
pig ext install xicor -v 15;   # install for PG 15
pig ext install xicor -v 14;   # install for PG 14
pig ext install xicor -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION xicor;
```

