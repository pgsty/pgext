---
title: "pgroonga_database"
linkTitle: "pgroonga_database"
description: "PGroonga database management module"
weight: 2111
categories: ["Fts"]
width: full
---

PGroonga database management module

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2111** | {{< badge content="pgroonga_database" link="https://github.com/pgroonga/pgroonga" >}} | {{< ext "pgroonga_database" "pgroonga" >}} | `4.0.0` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-dtr" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "zhparser" >}} {{< ext "pg_bigm" >}} {{< ext "pg_tokenizer" >}} {{< ext "pg_trgm" >}} {{< ext "fuzzystrmatch" >}} {{< ext "rum" >}} {{< ext "unaccent" >}} |
|    **Siblings**   | {{< ext "pgroonga" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgroonga" >}} | `4.0.0` | {{< badge content="18" color="red" alt="pgroonga_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pgroonga_$v*` | `groonga-libs` |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgroonga" >}} | `4.0.0` | {{< badge content="18" color="red" alt="postgresql-18-pgroonga" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pgroonga` | `libgroonga0` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pgroonga_18" >}}     | {{< pkg "pgroonga_17" "4.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_17-4.0.0-1.el8.x86_64.rpm" >}} | {{< pkg "pgroonga_16" "4.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_16-4.0.0-1.el8.x86_64.rpm" >}} | {{< pkg "pgroonga_15" "4.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_15-4.0.0-1.el8.x86_64.rpm" >}} | {{< pkg "pgroonga_14" "4.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_14-4.0.0-1.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pgroonga_18" >}}     | {{< pkg "pgroonga_17" "4.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_17-4.0.0-1.el8.aarch64.rpm" >}} | {{< pkg "pgroonga_16" "4.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_16-4.0.0-1.el8.aarch64.rpm" >}} | {{< pkg "pgroonga_15" "4.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_15-4.0.0-1.el8.aarch64.rpm" >}} | {{< pkg "pgroonga_14" "4.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_14-4.0.0-1.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pgroonga_18" >}}     | {{< pkg "pgroonga_17" "4.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_17-4.0.0-1.el9.x86_64.rpm" >}} | {{< pkg "pgroonga_16" "4.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_16-4.0.0-1.el9.x86_64.rpm" >}} | {{< pkg "pgroonga_15" "4.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_15-4.0.0-1.el9.x86_64.rpm" >}} | {{< pkg "pgroonga_14" "4.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_14-4.0.0-1.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pgroonga_18" >}}     | {{< pkg "pgroonga_17" "4.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_17-4.0.0-1.el9.aarch64.rpm" >}} | {{< pkg "pgroonga_16" "4.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_16-4.0.0-1.el9.aarch64.rpm" >}} | {{< pkg "pgroonga_15" "4.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_15-4.0.0-1.el9.aarch64.rpm" >}} | {{< pkg "pgroonga_14" "4.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_14-4.0.0-1.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pgroonga" >}}     | {{< pkg "postgresql-17-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pgroonga" >}}     | {{< pkg "postgresql-17-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pgroonga" >}}     | {{< pkg "postgresql-17-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pgroonga" >}}     | {{< pkg "postgresql-17-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pgroonga" >}}     | {{< pkg "postgresql-17-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pgroonga" >}}     | {{< pkg "postgresql-17-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pgroonga" "4.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgroonga_17` | 4.0.0 | `el8.x86_64` | pigsty | 342.0 KiB | [pgroonga_17-4.0.0-1.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_17-4.0.0-1.el8.x86_64.rpm) |
| `pgroonga_17` | 4.0.0 | `el8.aarch64` | pigsty | 330.5 KiB | [pgroonga_17-4.0.0-1.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_17-4.0.0-1.el8.aarch64.rpm) |
| `pgroonga_17` | 4.0.0 | `el9.aarch64` | pigsty | 320.1 KiB | [pgroonga_17-4.0.0-1.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_17-4.0.0-1.el9.aarch64.rpm) |
| `pgroonga_17` | 4.0.0 | `el9.x86_64` | pigsty | 328.0 KiB | [pgroonga_17-4.0.0-1.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_17-4.0.0-1.el9.x86_64.rpm) |
| `postgresql-17-pgroonga` | 4.0.0 | `d12.x86_64` | pigsty | 697.4 KiB | [postgresql-17-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgroonga` | 4.0.0 | `d12.aarch64` | pigsty | 693.0 KiB | [postgresql-17-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgroonga` | 4.0.0 | `u22.x86_64` | pigsty | 736.2 KiB | [postgresql-17-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgroonga` | 4.0.0 | `u22.aarch64` | pigsty | 739.9 KiB | [postgresql-17-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgroonga` | 4.0.0 | `u24.x86_64` | pigsty | 649.9 KiB | [postgresql-17-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgroonga` | 4.0.0 | `u24.aarch64` | pigsty | 653.1 KiB | [postgresql-17-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgroonga_16` | 4.0.0 | `el8.x86_64` | pigsty | 339.7 KiB | [pgroonga_16-4.0.0-1.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_16-4.0.0-1.el8.x86_64.rpm) |
| `pgroonga_16` | 4.0.0 | `el8.aarch64` | pigsty | 328.3 KiB | [pgroonga_16-4.0.0-1.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_16-4.0.0-1.el8.aarch64.rpm) |
| `pgroonga_16` | 4.0.0 | `el9.x86_64` | pigsty | 325.9 KiB | [pgroonga_16-4.0.0-1.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_16-4.0.0-1.el9.x86_64.rpm) |
| `pgroonga_16` | 4.0.0 | `el9.aarch64` | pigsty | 317.5 KiB | [pgroonga_16-4.0.0-1.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_16-4.0.0-1.el9.aarch64.rpm) |
| `postgresql-16-pgroonga` | 4.0.0 | `d12.x86_64` | pigsty | 684.9 KiB | [postgresql-16-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgroonga` | 4.0.0 | `d12.aarch64` | pigsty | 681.1 KiB | [postgresql-16-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgroonga` | 4.0.0 | `u22.aarch64` | pigsty | 727.5 KiB | [postgresql-16-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgroonga` | 4.0.0 | `u22.x86_64` | pigsty | 723.3 KiB | [postgresql-16-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgroonga` | 4.0.0 | `u24.x86_64` | pigsty | 642.4 KiB | [postgresql-16-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgroonga` | 4.0.0 | `u24.aarch64` | pigsty | 646.1 KiB | [postgresql-16-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgroonga_15` | 4.0.0 | `el8.x86_64` | pigsty | 341.9 KiB | [pgroonga_15-4.0.0-1.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_15-4.0.0-1.el8.x86_64.rpm) |
| `pgroonga_15` | 4.0.0 | `el8.aarch64` | pigsty | 330.8 KiB | [pgroonga_15-4.0.0-1.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_15-4.0.0-1.el8.aarch64.rpm) |
| `pgroonga_15` | 4.0.0 | `el9.x86_64` | pigsty | 327.9 KiB | [pgroonga_15-4.0.0-1.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_15-4.0.0-1.el9.x86_64.rpm) |
| `pgroonga_15` | 4.0.0 | `el9.aarch64` | pigsty | 321.7 KiB | [pgroonga_15-4.0.0-1.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_15-4.0.0-1.el9.aarch64.rpm) |
| `postgresql-15-pgroonga` | 4.0.0 | `d12.aarch64` | pigsty | 683.7 KiB | [postgresql-15-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgroonga` | 4.0.0 | `d12.x86_64` | pigsty | 687.5 KiB | [postgresql-15-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgroonga` | 4.0.0 | `u22.aarch64` | pigsty | 738.1 KiB | [postgresql-15-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgroonga` | 4.0.0 | `u22.x86_64` | pigsty | 729.8 KiB | [postgresql-15-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgroonga` | 4.0.0 | `u24.x86_64` | pigsty | 651.6 KiB | [postgresql-15-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgroonga` | 4.0.0 | `u24.aarch64` | pigsty | 654.5 KiB | [postgresql-15-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgroonga_14` | 4.0.0 | `el8.x86_64` | pigsty | 322.8 KiB | [pgroonga_14-4.0.0-1.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_14-4.0.0-1.el8.x86_64.rpm) |
| `pgroonga_14` | 4.0.0 | `el8.aarch64` | pigsty | 314.4 KiB | [pgroonga_14-4.0.0-1.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_14-4.0.0-1.el8.aarch64.rpm) |
| `pgroonga_14` | 4.0.0 | `el9.x86_64` | pigsty | 310.4 KiB | [pgroonga_14-4.0.0-1.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_14-4.0.0-1.el9.x86_64.rpm) |
| `pgroonga_14` | 4.0.0 | `el9.aarch64` | pigsty | 305.4 KiB | [pgroonga_14-4.0.0-1.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_14-4.0.0-1.el9.aarch64.rpm) |
| `postgresql-14-pgroonga` | 4.0.0 | `d12.x86_64` | pigsty | 630.7 KiB | [postgresql-14-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgroonga` | 4.0.0 | `d12.aarch64` | pigsty | 627.3 KiB | [postgresql-14-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgroonga` | 4.0.0 | `u22.x86_64` | pigsty | 670.2 KiB | [postgresql-14-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgroonga` | 4.0.0 | `u22.aarch64` | pigsty | 678.7 KiB | [postgresql-14-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgroonga` | 4.0.0 | `u24.x86_64` | pigsty | 595.1 KiB | [postgresql-14-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgroonga` | 4.0.0 | `u24.aarch64` | pigsty | 600.1 KiB | [postgresql-14-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgroonga_13` | 4.0.0 | `el8.aarch64` | pigsty | 314.3 KiB | [pgroonga_13-4.0.0-1.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_13-4.0.0-1.el8.aarch64.rpm) |
| `pgroonga_13` | 4.0.0 | `el8.x86_64` | pigsty | 321.9 KiB | [pgroonga_13-4.0.0-1.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_13-4.0.0-1.el8.x86_64.rpm) |
| `pgroonga_13` | 4.0.0 | `el9.aarch64` | pigsty | 305.2 KiB | [pgroonga_13-4.0.0-1.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_13-4.0.0-1.el9.aarch64.rpm) |
| `pgroonga_13` | 4.0.0 | `el9.x86_64` | pigsty | 310.4 KiB | [pgroonga_13-4.0.0-1.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_13-4.0.0-1.el9.x86_64.rpm) |
| `postgresql-13-pgroonga` | 4.0.0 | `d12.aarch64` | pigsty | 730.7 KiB | [postgresql-13-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgroonga` | 4.0.0 | `d12.x86_64` | pigsty | 738.3 KiB | [postgresql-13-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgroonga` | 4.0.0 | `u22.aarch64` | pigsty | 786.5 KiB | [postgresql-13-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgroonga` | 4.0.0 | `u22.x86_64` | pigsty | 786.0 KiB | [postgresql-13-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgroonga` | 4.0.0 | `u24.aarch64` | pigsty | 695.5 KiB | [postgresql-13-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pgroonga` | 4.0.0 | `u24.x86_64` | pigsty | 695.4 KiB | [postgresql-13-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgroonga/pgroonga" title="Repository" icon="github" subtitle="github.com/pgroonga/pgroonga" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pgroonga-4.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgroonga_database; # get pgroonga_database source code
pig build dep pgroonga_database; # install build dependencies
pig build pkg pgroonga_database; # build extension rpm or deb
pig build ext pgroonga_database; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgroonga_database; # install by extension name, for the current active PG version
pig ext install pgroonga; # install via package alias, for the active PG version
pig ext install pgroonga_database -v 17;   # install for PG 17
pig ext install pgroonga_database -v 16;   # install for PG 16
pig ext install pgroonga_database -v 15;   # install for PG 15
pig ext install pgroonga_database -v 14;   # install for PG 14
pig ext install pgroonga_database -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgroonga_database;
```

