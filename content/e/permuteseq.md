---
title: "permuteseq"
linkTitle: "permuteseq"
description: "Pseudo-randomly permute sequences with a format-preserving encryption on elements"
weight: 4550
categories: ["Func"]
width: full
---

Pseudo-randomly permute sequences with a format-preserving encryption on elements

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4550** | {{< badge content="permuteseq" link="https://github.com/dverite/permuteseq" >}} | {{< ext "permuteseq" "permuteseq" >}} | `1.2.2` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "random" >}} {{< ext "sequential_uuids" >}} {{< ext "pg_hashids" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "uuid-ossp" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/permuteseq" >}} | `1.2.2` | {{< badge content="18" color="red" alt="permuteseq_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `permuteseq_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/permuteseq" >}} | `1.2.2` | {{< badge content="18" color="red" alt="postgresql-18-permuteseq" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-permuteseq` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "permuteseq_18" >}}     | {{< pkg "permuteseq_17" "1.2.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/permuteseq_17-1.2.2-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "permuteseq_16" "1.2.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/permuteseq_16-1.2.2-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "permuteseq_15" "1.2.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/permuteseq_15-1.2.2-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "permuteseq_14" "1.2.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/permuteseq_14-1.2.2-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "permuteseq_18" >}}     | {{< pkg "permuteseq_17" "1.2.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/permuteseq_17-1.2.2-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "permuteseq_16" "1.2.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/permuteseq_16-1.2.2-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "permuteseq_15" "1.2.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/permuteseq_15-1.2.2-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "permuteseq_14" "1.2.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/permuteseq_14-1.2.2-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "permuteseq_18" >}}     | {{< pkg "permuteseq_17" "1.2.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/permuteseq_17-1.2.2-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "permuteseq_16" "1.2.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/permuteseq_16-1.2.2-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "permuteseq_15" "1.2.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/permuteseq_15-1.2.2-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "permuteseq_14" "1.2.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/permuteseq_14-1.2.2-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "permuteseq_18" >}}     | {{< pkg "permuteseq_17" "1.2.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/permuteseq_17-1.2.2-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "permuteseq_16" "1.2.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/permuteseq_16-1.2.2-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "permuteseq_15" "1.2.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/permuteseq_15-1.2.2-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "permuteseq_14" "1.2.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/permuteseq_14-1.2.2-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-permuteseq" >}}     | {{< pkg "postgresql-17-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-permuteseq" >}}     | {{< pkg "postgresql-17-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-permuteseq" >}}     | {{< pkg "postgresql-17-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-permuteseq" >}}     | {{< pkg "postgresql-17-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-permuteseq" >}}     | {{< pkg "postgresql-17-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-permuteseq" >}}     | {{< pkg "postgresql-17-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-permuteseq" "1.2.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `permuteseq_17` | 1.2.2 | `el8.x86_64` | pigsty | 13.2 KiB | [permuteseq_17-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/permuteseq_17-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `permuteseq_17` | 1.2.2 | `el8.aarch64` | pigsty | 13.4 KiB | [permuteseq_17-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/permuteseq_17-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `permuteseq_17` | 1.2.2 | `el9.aarch64` | pigsty | 13.1 KiB | [permuteseq_17-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/permuteseq_17-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `permuteseq_17` | 1.2.2 | `el9.x86_64` | pigsty | 13.1 KiB | [permuteseq_17-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/permuteseq_17-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-permuteseq` | 1.2.2 | `d12.x86_64` | pigsty | 15.9 KiB | [postgresql-17-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-permuteseq` | 1.2.2 | `d12.aarch64` | pigsty | 15.8 KiB | [postgresql-17-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-permuteseq` | 1.2.2 | `u22.x86_64` | pigsty | 16.3 KiB | [postgresql-17-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-permuteseq` | 1.2.2 | `u22.aarch64` | pigsty | 16.1 KiB | [postgresql-17-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-permuteseq` | 1.2.2 | `u24.x86_64` | pigsty | 15.9 KiB | [postgresql-17-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-permuteseq` | 1.2.2 | `u24.aarch64` | pigsty | 15.8 KiB | [postgresql-17-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `permuteseq_16` | 1.2.2 | `el8.x86_64` | pigsty | 13.2 KiB | [permuteseq_16-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/permuteseq_16-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `permuteseq_16` | 1.2.2 | `el8.aarch64` | pigsty | 13.4 KiB | [permuteseq_16-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/permuteseq_16-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `permuteseq_16` | 1.2.2 | `el9.x86_64` | pigsty | 13.1 KiB | [permuteseq_16-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/permuteseq_16-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `permuteseq_16` | 1.2.2 | `el9.aarch64` | pigsty | 13.1 KiB | [permuteseq_16-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/permuteseq_16-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-permuteseq` | 1.2.2 | `d12.x86_64` | pigsty | 15.8 KiB | [postgresql-16-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-permuteseq` | 1.2.2 | `d12.aarch64` | pigsty | 15.8 KiB | [postgresql-16-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-permuteseq` | 1.2.2 | `u22.aarch64` | pigsty | 16.1 KiB | [postgresql-16-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-permuteseq` | 1.2.2 | `u22.x86_64` | pigsty | 16.3 KiB | [postgresql-16-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-permuteseq` | 1.2.2 | `u24.x86_64` | pigsty | 15.9 KiB | [postgresql-16-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-permuteseq` | 1.2.2 | `u24.aarch64` | pigsty | 15.8 KiB | [postgresql-16-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `permuteseq_15` | 1.2.2 | `el8.x86_64` | pigsty | 13.0 KiB | [permuteseq_15-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/permuteseq_15-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `permuteseq_15` | 1.2.2 | `el8.aarch64` | pigsty | 13.2 KiB | [permuteseq_15-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/permuteseq_15-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `permuteseq_15` | 1.2.2 | `el9.x86_64` | pigsty | 12.8 KiB | [permuteseq_15-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/permuteseq_15-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `permuteseq_15` | 1.2.2 | `el9.aarch64` | pigsty | 12.8 KiB | [permuteseq_15-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/permuteseq_15-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-permuteseq` | 1.2.2 | `d12.aarch64` | pigsty | 14.9 KiB | [postgresql-15-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-permuteseq` | 1.2.2 | `d12.x86_64` | pigsty | 15.0 KiB | [postgresql-15-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-permuteseq` | 1.2.2 | `u22.aarch64` | pigsty | 15.2 KiB | [postgresql-15-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-permuteseq` | 1.2.2 | `u22.x86_64` | pigsty | 15.2 KiB | [postgresql-15-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-permuteseq` | 1.2.2 | `u24.x86_64` | pigsty | 15.0 KiB | [postgresql-15-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-permuteseq` | 1.2.2 | `u24.aarch64` | pigsty | 14.8 KiB | [postgresql-15-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `permuteseq_14` | 1.2.2 | `el8.x86_64` | pigsty | 12.9 KiB | [permuteseq_14-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/permuteseq_14-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `permuteseq_14` | 1.2.2 | `el8.aarch64` | pigsty | 13.2 KiB | [permuteseq_14-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/permuteseq_14-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `permuteseq_14` | 1.2.2 | `el9.x86_64` | pigsty | 12.8 KiB | [permuteseq_14-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/permuteseq_14-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `permuteseq_14` | 1.2.2 | `el9.aarch64` | pigsty | 12.8 KiB | [permuteseq_14-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/permuteseq_14-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-permuteseq` | 1.2.2 | `d12.x86_64` | pigsty | 15.0 KiB | [postgresql-14-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-permuteseq` | 1.2.2 | `d12.aarch64` | pigsty | 14.9 KiB | [postgresql-14-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-permuteseq` | 1.2.2 | `u22.x86_64` | pigsty | 15.1 KiB | [postgresql-14-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-permuteseq` | 1.2.2 | `u22.aarch64` | pigsty | 15.1 KiB | [postgresql-14-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-permuteseq` | 1.2.2 | `u24.x86_64` | pigsty | 15.0 KiB | [postgresql-14-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-permuteseq` | 1.2.2 | `u24.aarch64` | pigsty | 14.8 KiB | [postgresql-14-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `permuteseq_13` | 1.2.2 | `el8.aarch64` | pigsty | 13.2 KiB | [permuteseq_13-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/permuteseq_13-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `permuteseq_13` | 1.2.2 | `el8.x86_64` | pigsty | 12.9 KiB | [permuteseq_13-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/permuteseq_13-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `permuteseq_13` | 1.2.2 | `el9.aarch64` | pigsty | 12.8 KiB | [permuteseq_13-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/permuteseq_13-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `permuteseq_13` | 1.2.2 | `el9.x86_64` | pigsty | 12.7 KiB | [permuteseq_13-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/permuteseq_13-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-permuteseq` | 1.2.2 | `d12.aarch64` | pigsty | 14.9 KiB | [postgresql-13-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-13-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-permuteseq` | 1.2.2 | `d12.x86_64` | pigsty | 14.9 KiB | [postgresql-13-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-13-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-permuteseq` | 1.2.2 | `u22.aarch64` | pigsty | 15.1 KiB | [postgresql-13-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-13-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-permuteseq` | 1.2.2 | `u22.x86_64` | pigsty | 15.1 KiB | [postgresql-13-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-13-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-permuteseq` | 1.2.2 | `u24.aarch64` | pigsty | 14.8 KiB | [postgresql-13-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-13-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-permuteseq` | 1.2.2 | `u24.x86_64` | pigsty | 14.8 KiB | [postgresql-13-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-13-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/dverite/permuteseq" title="Repository" icon="github" subtitle="github.com/dverite/permuteseq" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="permuteseq-1.2.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build get permuteseq; # get permuteseq source code
pig build dep permuteseq; # install build dependencies
pig build pkg permuteseq; # build extension rpm or deb
pig build ext permuteseq; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install permuteseq; # install by extension name, for the current active PG version
pig ext install permuteseq; # install via package alias, for the active PG version
pig ext install permuteseq -v 18;   # install for PG 18
pig ext install permuteseq -v 17;   # install for PG 17
pig ext install permuteseq -v 16;   # install for PG 16
pig ext install permuteseq -v 15;   # install for PG 15
pig ext install permuteseq -v 14;   # install for PG 14
pig ext install permuteseq -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION permuteseq;
```

