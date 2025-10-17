---
title: "convert"
linkTitle: "convert"
description: "conversion functions for spatial, routing and other specialized uses"
weight: 4850
categories: ["Func"]
width: full
---

conversion functions for spatial, routing and other specialized uses

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4850** | {{< badge content="convert" link="https://github.com/rustprooflabs/convert" >}} | {{< ext "convert" "pg_convert" >}} | `0.0.4` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "permuteseq" >}} {{< ext "pg_hashids" >}} {{< ext "sequential_uuids" >}} {{< ext "topn" >}} {{< ext "quantile" >}} |

> [!Note] pgrx=0.14.1


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/convert" >}} | `0.0.4` | {{< badge content="18" color="red" alt="pg_convert_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_convert_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/convert" >}} | `0.0.4` | {{< badge content="18" color="red" alt="postgresql-18-convert" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-convert` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_convert_18" >}}     | {{< pkg "pg_convert_17" "0.0.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_17-0.0.4-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_convert_16" "0.0.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_16-0.0.4-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_convert_15" "0.0.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_15-0.0.4-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_convert_14" "0.0.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_14-0.0.4-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_convert_18" >}}     | {{< pkg "pg_convert_17" "0.0.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_17-0.0.4-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_convert_16" "0.0.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_16-0.0.4-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_convert_15" "0.0.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_15-0.0.4-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_convert_14" "0.0.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_14-0.0.4-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_convert_18" >}}     | {{< pkg "pg_convert_17" "0.0.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_17-0.0.4-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_convert_16" "0.0.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_16-0.0.4-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_convert_15" "0.0.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_15-0.0.4-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_convert_14" "0.0.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_14-0.0.4-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_convert_18" >}}     | {{< pkg "pg_convert_17" "0.0.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_17-0.0.4-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_convert_16" "0.0.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_16-0.0.4-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_convert_15" "0.0.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_15-0.0.4-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_convert_14" "0.0.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_14-0.0.4-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-convert" >}}     | {{< pkg "postgresql-17-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-17-convert_0.0.4-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-16-convert_0.0.4-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-15-convert_0.0.4-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-14-convert_0.0.4-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-convert" >}}     | {{< pkg "postgresql-17-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-17-convert_0.0.4-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-16-convert_0.0.4-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-15-convert_0.0.4-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-14-convert_0.0.4-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-convert" >}}     | {{< pkg "postgresql-17-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-17-convert_0.0.4-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-16-convert_0.0.4-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-15-convert_0.0.4-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-14-convert_0.0.4-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-convert" >}}     | {{< pkg "postgresql-17-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-17-convert_0.0.4-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-16-convert_0.0.4-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-15-convert_0.0.4-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-14-convert_0.0.4-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-convert" >}}     | {{< pkg "postgresql-17-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-17-convert_0.0.4-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-16-convert_0.0.4-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-15-convert_0.0.4-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-14-convert_0.0.4-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-convert" >}}     | {{< pkg "postgresql-17-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-17-convert_0.0.4-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-16-convert_0.0.4-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-15-convert_0.0.4-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-convert" "0.0.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-14-convert_0.0.4-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_convert_17` | 0.0.4 | `el8.x86_64` | pigsty | 203.3 KiB | [pg_convert_17-0.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_17-0.0.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_convert_17` | 0.0.4 | `el8.aarch64` | pigsty | 191.5 KiB | [pg_convert_17-0.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_17-0.0.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_convert_17` | 0.0.4 | `el9.aarch64` | pigsty | 205.4 KiB | [pg_convert_17-0.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_17-0.0.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_convert_17` | 0.0.4 | `el9.x86_64` | pigsty | 208.5 KiB | [pg_convert_17-0.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_17-0.0.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-convert` | 0.0.4 | `d12.x86_64` | pigsty | 164.2 KiB | [postgresql-17-convert_0.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-17-convert_0.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-convert` | 0.0.4 | `d12.aarch64` | pigsty | 148.3 KiB | [postgresql-17-convert_0.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-17-convert_0.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-convert` | 0.0.4 | `u22.x86_64` | pigsty | 180.6 KiB | [postgresql-17-convert_0.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-17-convert_0.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-convert` | 0.0.4 | `u22.aarch64` | pigsty | 172.2 KiB | [postgresql-17-convert_0.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-17-convert_0.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-convert` | 0.0.4 | `u24.x86_64` | pigsty | 179.4 KiB | [postgresql-17-convert_0.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-17-convert_0.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-convert` | 0.0.4 | `u24.aarch64` | pigsty | 171.4 KiB | [postgresql-17-convert_0.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-17-convert_0.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_convert_16` | 0.0.4 | `el8.x86_64` | pigsty | 203.3 KiB | [pg_convert_16-0.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_16-0.0.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_convert_16` | 0.0.4 | `el8.aarch64` | pigsty | 191.6 KiB | [pg_convert_16-0.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_16-0.0.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_convert_16` | 0.0.4 | `el9.x86_64` | pigsty | 208.6 KiB | [pg_convert_16-0.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_16-0.0.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_convert_16` | 0.0.4 | `el9.aarch64` | pigsty | 205.5 KiB | [pg_convert_16-0.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_16-0.0.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-convert` | 0.0.4 | `d12.x86_64` | pigsty | 164.2 KiB | [postgresql-16-convert_0.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-16-convert_0.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-convert` | 0.0.4 | `d12.aarch64` | pigsty | 148.3 KiB | [postgresql-16-convert_0.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-16-convert_0.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-convert` | 0.0.4 | `u22.aarch64` | pigsty | 172.1 KiB | [postgresql-16-convert_0.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-16-convert_0.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-convert` | 0.0.4 | `u22.x86_64` | pigsty | 180.7 KiB | [postgresql-16-convert_0.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-16-convert_0.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-convert` | 0.0.4 | `u24.x86_64` | pigsty | 179.5 KiB | [postgresql-16-convert_0.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-16-convert_0.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-convert` | 0.0.4 | `u24.aarch64` | pigsty | 171.6 KiB | [postgresql-16-convert_0.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-16-convert_0.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_convert_15` | 0.0.4 | `el8.x86_64` | pigsty | 203.3 KiB | [pg_convert_15-0.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_15-0.0.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_convert_15` | 0.0.4 | `el8.aarch64` | pigsty | 191.6 KiB | [pg_convert_15-0.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_15-0.0.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_convert_15` | 0.0.4 | `el9.x86_64` | pigsty | 208.7 KiB | [pg_convert_15-0.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_15-0.0.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_convert_15` | 0.0.4 | `el9.aarch64` | pigsty | 205.5 KiB | [pg_convert_15-0.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_15-0.0.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-convert` | 0.0.4 | `d12.aarch64` | pigsty | 148.4 KiB | [postgresql-15-convert_0.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-15-convert_0.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-convert` | 0.0.4 | `d12.x86_64` | pigsty | 164.0 KiB | [postgresql-15-convert_0.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-15-convert_0.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-convert` | 0.0.4 | `u22.aarch64` | pigsty | 172.2 KiB | [postgresql-15-convert_0.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-15-convert_0.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-convert` | 0.0.4 | `u22.x86_64` | pigsty | 181.0 KiB | [postgresql-15-convert_0.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-15-convert_0.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-convert` | 0.0.4 | `u24.x86_64` | pigsty | 179.5 KiB | [postgresql-15-convert_0.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-15-convert_0.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-convert` | 0.0.4 | `u24.aarch64` | pigsty | 171.5 KiB | [postgresql-15-convert_0.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-15-convert_0.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_convert_14` | 0.0.4 | `el8.x86_64` | pigsty | 203.4 KiB | [pg_convert_14-0.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_14-0.0.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_convert_14` | 0.0.4 | `el8.aarch64` | pigsty | 191.6 KiB | [pg_convert_14-0.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_14-0.0.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_convert_14` | 0.0.4 | `el9.x86_64` | pigsty | 208.6 KiB | [pg_convert_14-0.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_14-0.0.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_convert_14` | 0.0.4 | `el9.aarch64` | pigsty | 205.5 KiB | [pg_convert_14-0.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_14-0.0.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-convert` | 0.0.4 | `d12.x86_64` | pigsty | 164.1 KiB | [postgresql-14-convert_0.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-14-convert_0.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-convert` | 0.0.4 | `d12.aarch64` | pigsty | 148.3 KiB | [postgresql-14-convert_0.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-14-convert_0.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-convert` | 0.0.4 | `u22.x86_64` | pigsty | 181.0 KiB | [postgresql-14-convert_0.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-14-convert_0.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-convert` | 0.0.4 | `u22.aarch64` | pigsty | 172.2 KiB | [postgresql-14-convert_0.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-14-convert_0.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-convert` | 0.0.4 | `u24.x86_64` | pigsty | 179.6 KiB | [postgresql-14-convert_0.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-14-convert_0.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-convert` | 0.0.4 | `u24.aarch64` | pigsty | 171.4 KiB | [postgresql-14-convert_0.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-14-convert_0.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_convert_13` | 0.0.4 | `el8.aarch64` | pigsty | 191.7 KiB | [pg_convert_13-0.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_13-0.0.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_convert_13` | 0.0.4 | `el8.x86_64` | pigsty | 203.4 KiB | [pg_convert_13-0.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_13-0.0.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_convert_13` | 0.0.4 | `el9.aarch64` | pigsty | 205.5 KiB | [pg_convert_13-0.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_13-0.0.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_convert_13` | 0.0.4 | `el9.x86_64` | pigsty | 208.6 KiB | [pg_convert_13-0.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_13-0.0.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-convert` | 0.0.4 | `d12.aarch64` | pigsty | 148.4 KiB | [postgresql-13-convert_0.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-13-convert_0.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-convert` | 0.0.4 | `d12.x86_64` | pigsty | 164.1 KiB | [postgresql-13-convert_0.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-13-convert_0.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-convert` | 0.0.4 | `u22.aarch64` | pigsty | 172.2 KiB | [postgresql-13-convert_0.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-13-convert_0.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-convert` | 0.0.4 | `u22.x86_64` | pigsty | 181.1 KiB | [postgresql-13-convert_0.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-13-convert_0.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-convert` | 0.0.4 | `u24.aarch64` | pigsty | 171.5 KiB | [postgresql-13-convert_0.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-13-convert_0.0.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-convert` | 0.0.4 | `u24.x86_64` | pigsty | 179.6 KiB | [postgresql-13-convert_0.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-13-convert_0.0.4-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rustprooflabs/convert" title="Repository" icon="github" subtitle="github.com/rustprooflabs/convert" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="convert-0.0.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build get convert; # get convert source code
pig build dep convert; # install build dependencies
pig build pkg convert; # build extension rpm or deb
pig build ext convert; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install convert; # install by extension name, for the current active PG version
pig ext install pg_convert; # install via package alias, for the active PG version
pig ext install convert -v 17;   # install for PG 17
pig ext install convert -v 16;   # install for PG 16
pig ext install convert -v 15;   # install for PG 15
pig ext install convert -v 14;   # install for PG 14
pig ext install convert -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION convert;
```

