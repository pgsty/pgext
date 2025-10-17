---
title: "pgx_ulid"
linkTitle: "pgx_ulid"
description: "ulid type and methods"
weight: 4510
categories: ["Func"]
width: full
---

ulid type and methods

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4510** | {{< badge content="pgx_ulid" link="https://github.com/pksunkara/pgx_ulid" >}} | {{< ext "pgx_ulid" "pgx_ulid" >}} | `0.2.0` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pg_uuidv7" >}} {{< ext "sequential_uuids" >}} {{< ext "uuid-ossp" >}} {{< ext "pg_hashids" >}} {{< ext "permuteseq" >}} |

> [!Note] pgrx=0.12.7


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgx_ulid" >}} | `0.2.0` | {{< badge content="18" color="red" alt="pgx_ulid_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pgx_ulid_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgx_ulid" >}} | `0.2.0` | {{< badge content="18" color="red" alt="postgresql-18-pgx-ulid" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pgx-ulid` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pgx_ulid_18" >}}     | {{< pkg "pgx_ulid_17" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgx_ulid_17-0.2.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pgx_ulid_16" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgx_ulid_16-0.2.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pgx_ulid_15" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgx_ulid_15-0.2.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pgx_ulid_14" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgx_ulid_14-0.2.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pgx_ulid_18" >}}     | {{< pkg "pgx_ulid_17" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgx_ulid_17-0.2.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pgx_ulid_16" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgx_ulid_16-0.2.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pgx_ulid_15" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgx_ulid_15-0.2.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pgx_ulid_14" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgx_ulid_14-0.2.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pgx_ulid_18" >}}     | {{< pkg "pgx_ulid_17" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgx_ulid_17-0.2.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pgx_ulid_16" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgx_ulid_16-0.2.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pgx_ulid_15" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgx_ulid_15-0.2.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pgx_ulid_14" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgx_ulid_14-0.2.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pgx_ulid_18" >}}     | {{< pkg "pgx_ulid_17" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgx_ulid_17-0.2.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pgx_ulid_16" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgx_ulid_16-0.2.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pgx_ulid_15" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgx_ulid_15-0.2.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pgx_ulid_14" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgx_ulid_14-0.2.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pgx-ulid" >}}     | {{< pkg "postgresql-17-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pgx-ulid" >}}     | {{< pkg "postgresql-17-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pgx-ulid" >}}     | {{< pkg "postgresql-17-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pgx-ulid" >}}     | {{< pkg "postgresql-17-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pgx-ulid" >}}     | {{< pkg "postgresql-17-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pgx-ulid" >}}     | {{< pkg "postgresql-17-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pgx-ulid" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgx_ulid_17` | 0.2.0 | `el8.aarch64` | pigsty | 254.6 KiB | [pgx_ulid_17-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgx_ulid_17-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgx_ulid_17` | 0.2.0 | `el8.x86_64` | pigsty | 272.8 KiB | [pgx_ulid_17-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgx_ulid_17-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgx_ulid_17` | 0.2.0 | `el9.aarch64` | pigsty | 272.7 KiB | [pgx_ulid_17-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgx_ulid_17-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgx_ulid_17` | 0.2.0 | `el9.x86_64` | pigsty | 277.4 KiB | [pgx_ulid_17-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgx_ulid_17-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pgx-ulid` | 0.2.0 | `d12.x86_64` | pigsty | 219.6 KiB | [postgresql-17-pgx-ulid_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgx-ulid` | 0.2.0 | `d12.aarch64` | pigsty | 194.7 KiB | [postgresql-17-pgx-ulid_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgx-ulid` | 0.2.0 | `u22.aarch64` | pigsty | 227.0 KiB | [postgresql-17-pgx-ulid_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgx-ulid` | 0.2.0 | `u22.x86_64` | pigsty | 240.8 KiB | [postgresql-17-pgx-ulid_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgx-ulid` | 0.2.0 | `u24.x86_64` | pigsty | 238.7 KiB | [postgresql-17-pgx-ulid_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgx-ulid` | 0.2.0 | `u24.aarch64` | pigsty | 225.0 KiB | [postgresql-17-pgx-ulid_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgx_ulid_16` | 0.2.0 | `el8.x86_64` | pigsty | 272.6 KiB | [pgx_ulid_16-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgx_ulid_16-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgx_ulid_16` | 0.2.0 | `el8.aarch64` | pigsty | 254.4 KiB | [pgx_ulid_16-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgx_ulid_16-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgx_ulid_16` | 0.2.0 | `el9.aarch64` | pigsty | 272.7 KiB | [pgx_ulid_16-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgx_ulid_16-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgx_ulid_16` | 0.2.0 | `el9.x86_64` | pigsty | 277.7 KiB | [pgx_ulid_16-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgx_ulid_16-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-16-pgx-ulid` | 0.2.0 | `d12.aarch64` | pigsty | 194.6 KiB | [postgresql-16-pgx-ulid_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgx-ulid` | 0.2.0 | `d12.x86_64` | pigsty | 219.6 KiB | [postgresql-16-pgx-ulid_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgx-ulid` | 0.2.0 | `u22.x86_64` | pigsty | 240.6 KiB | [postgresql-16-pgx-ulid_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgx-ulid` | 0.2.0 | `u22.aarch64` | pigsty | 226.7 KiB | [postgresql-16-pgx-ulid_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgx-ulid` | 0.2.0 | `u24.aarch64` | pigsty | 225.0 KiB | [postgresql-16-pgx-ulid_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgx-ulid` | 0.2.0 | `u24.x86_64` | pigsty | 238.6 KiB | [postgresql-16-pgx-ulid_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgx_ulid_15` | 0.2.0 | `el8.x86_64` | pigsty | 272.8 KiB | [pgx_ulid_15-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgx_ulid_15-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgx_ulid_15` | 0.2.0 | `el8.aarch64` | pigsty | 254.4 KiB | [pgx_ulid_15-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgx_ulid_15-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgx_ulid_15` | 0.2.0 | `el9.x86_64` | pigsty | 278.0 KiB | [pgx_ulid_15-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgx_ulid_15-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgx_ulid_15` | 0.2.0 | `el9.aarch64` | pigsty | 272.5 KiB | [pgx_ulid_15-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgx_ulid_15-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pgx-ulid` | 0.2.0 | `d12.x86_64` | pigsty | 219.6 KiB | [postgresql-15-pgx-ulid_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgx-ulid` | 0.2.0 | `d12.aarch64` | pigsty | 194.7 KiB | [postgresql-15-pgx-ulid_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgx-ulid` | 0.2.0 | `u22.aarch64` | pigsty | 227.3 KiB | [postgresql-15-pgx-ulid_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgx-ulid` | 0.2.0 | `u22.x86_64` | pigsty | 240.9 KiB | [postgresql-15-pgx-ulid_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgx-ulid` | 0.2.0 | `u24.aarch64` | pigsty | 225.3 KiB | [postgresql-15-pgx-ulid_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgx-ulid` | 0.2.0 | `u24.x86_64` | pigsty | 238.8 KiB | [postgresql-15-pgx-ulid_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgx_ulid_14` | 0.2.0 | `el8.aarch64` | pigsty | 254.4 KiB | [pgx_ulid_14-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgx_ulid_14-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgx_ulid_14` | 0.2.0 | `el8.x86_64` | pigsty | 272.6 KiB | [pgx_ulid_14-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgx_ulid_14-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgx_ulid_14` | 0.2.0 | `el9.x86_64` | pigsty | 277.6 KiB | [pgx_ulid_14-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgx_ulid_14-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgx_ulid_14` | 0.2.0 | `el9.aarch64` | pigsty | 272.3 KiB | [pgx_ulid_14-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgx_ulid_14-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pgx-ulid` | 0.2.0 | `d12.aarch64` | pigsty | 194.5 KiB | [postgresql-14-pgx-ulid_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgx-ulid` | 0.2.0 | `d12.x86_64` | pigsty | 219.8 KiB | [postgresql-14-pgx-ulid_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgx-ulid` | 0.2.0 | `u22.x86_64` | pigsty | 240.8 KiB | [postgresql-14-pgx-ulid_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgx-ulid` | 0.2.0 | `u22.aarch64` | pigsty | 226.6 KiB | [postgresql-14-pgx-ulid_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgx-ulid` | 0.2.0 | `u24.aarch64` | pigsty | 225.1 KiB | [postgresql-14-pgx-ulid_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgx-ulid` | 0.2.0 | `u24.x86_64` | pigsty | 239.2 KiB | [postgresql-14-pgx-ulid_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pksunkara/pgx_ulid" title="Repository" icon="github" subtitle="github.com/pksunkara/pgx_ulid" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pgx_ulid-0.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgx_ulid; # get pgx_ulid source code
pig build dep pgx_ulid; # install build dependencies
pig build pkg pgx_ulid; # build extension rpm or deb
pig build ext pgx_ulid; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgx_ulid; # install by extension name, for the current active PG version
pig ext install pgx_ulid; # install via package alias, for the active PG version
pig ext install pgx_ulid -v 17;   # install for PG 17
pig ext install pgx_ulid -v 16;   # install for PG 16
pig ext install pgx_ulid -v 15;   # install for PG 15
pig ext install pgx_ulid -v 14;   # install for PG 14

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgx_ulid;
```

