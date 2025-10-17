---
title: "pguecc"
linkTitle: "pguecc"
description: "uECC bindings for Postgres"
weight: 4460
categories: ["Util"]
width: full
---

uECC bindings for Postgres

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4460** | {{< badge content="pguecc" link="https://github.com/ameensol/pg-ecdsa" >}} | {{< ext "pguecc" "pg_ecdsa" >}} | `1.0` | {{< category "UTIL" >}} | {{< license "BSD 2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "hashlib" >}} {{< ext "shacrypt" >}} {{< ext "cryptint" >}} {{< ext "pgcrypto" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pguecc" >}} | `1.0` | {{< badge content="18" color="red" alt="pg_ecdsa_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_ecdsa_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pguecc" >}} | `1.0` | {{< badge content="18" color="red" alt="postgresql-18-pg-ecdsa" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-ecdsa` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_ecdsa_18" >}}     | {{< pkg "pg_ecdsa_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ecdsa_17-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_ecdsa_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ecdsa_16-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_ecdsa_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ecdsa_15-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_ecdsa_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ecdsa_14-1.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_ecdsa_18" >}}     | {{< pkg "pg_ecdsa_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ecdsa_17-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_ecdsa_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ecdsa_16-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_ecdsa_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ecdsa_15-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_ecdsa_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ecdsa_14-1.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_ecdsa_18" >}}     | {{< pkg "pg_ecdsa_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ecdsa_17-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_ecdsa_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ecdsa_16-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_ecdsa_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ecdsa_15-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_ecdsa_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ecdsa_14-1.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_ecdsa_18" >}}     | {{< pkg "pg_ecdsa_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ecdsa_17-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_ecdsa_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ecdsa_16-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_ecdsa_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ecdsa_15-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_ecdsa_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ecdsa_14-1.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-ecdsa" >}}     | {{< pkg "postgresql-17-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-17-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-16-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-15-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-14-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-ecdsa" >}}     | {{< pkg "postgresql-17-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-17-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-16-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-15-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-14-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-ecdsa" >}}     | {{< pkg "postgresql-17-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-17-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-16-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-15-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-14-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-ecdsa" >}}     | {{< pkg "postgresql-17-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-17-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-16-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-15-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-14-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-ecdsa" >}}     | {{< pkg "postgresql-17-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-17-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-16-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-15-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-14-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-ecdsa" >}}     | {{< pkg "postgresql-17-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-17-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-16-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-15-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-ecdsa" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-14-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_ecdsa_17` | 1.0 | `el8.x86_64` | pigsty | 26.9 KiB | [pg_ecdsa_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ecdsa_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_ecdsa_17` | 1.0 | `el8.aarch64` | pigsty | 26.3 KiB | [pg_ecdsa_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ecdsa_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_ecdsa_17` | 1.0 | `el9.aarch64` | pigsty | 24.8 KiB | [pg_ecdsa_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ecdsa_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_ecdsa_17` | 1.0 | `el9.x86_64` | pigsty | 25.5 KiB | [pg_ecdsa_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ecdsa_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-ecdsa` | 1.0 | `d12.x86_64` | pigsty | 67.8 KiB | [postgresql-17-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-17-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-ecdsa` | 1.0 | `d12.aarch64` | pigsty | 66.4 KiB | [postgresql-17-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-17-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-ecdsa` | 1.0 | `u22.x86_64` | pigsty | 69.4 KiB | [postgresql-17-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-17-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-ecdsa` | 1.0 | `u22.aarch64` | pigsty | 68.3 KiB | [postgresql-17-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-17-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-ecdsa` | 1.0 | `u24.x86_64` | pigsty | 66.4 KiB | [postgresql-17-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-17-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-ecdsa` | 1.0 | `u24.aarch64` | pigsty | 65.0 KiB | [postgresql-17-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-17-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_ecdsa_16` | 1.0 | `el8.x86_64` | pigsty | 26.9 KiB | [pg_ecdsa_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ecdsa_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_ecdsa_16` | 1.0 | `el8.aarch64` | pigsty | 26.3 KiB | [pg_ecdsa_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ecdsa_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_ecdsa_16` | 1.0 | `el9.x86_64` | pigsty | 25.5 KiB | [pg_ecdsa_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ecdsa_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_ecdsa_16` | 1.0 | `el9.aarch64` | pigsty | 24.9 KiB | [pg_ecdsa_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ecdsa_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-ecdsa` | 1.0 | `d12.x86_64` | pigsty | 67.8 KiB | [postgresql-16-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-16-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-ecdsa` | 1.0 | `d12.aarch64` | pigsty | 66.4 KiB | [postgresql-16-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-16-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-ecdsa` | 1.0 | `u22.aarch64` | pigsty | 68.3 KiB | [postgresql-16-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-16-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-ecdsa` | 1.0 | `u22.x86_64` | pigsty | 69.4 KiB | [postgresql-16-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-16-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-ecdsa` | 1.0 | `u24.x86_64` | pigsty | 66.4 KiB | [postgresql-16-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-16-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-ecdsa` | 1.0 | `u24.aarch64` | pigsty | 65.0 KiB | [postgresql-16-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-16-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_ecdsa_15` | 1.0 | `el8.x86_64` | pigsty | 27.2 KiB | [pg_ecdsa_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ecdsa_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_ecdsa_15` | 1.0 | `el8.aarch64` | pigsty | 26.6 KiB | [pg_ecdsa_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ecdsa_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_ecdsa_15` | 1.0 | `el9.x86_64` | pigsty | 27.9 KiB | [pg_ecdsa_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ecdsa_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_ecdsa_15` | 1.0 | `el9.aarch64` | pigsty | 27.3 KiB | [pg_ecdsa_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ecdsa_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-ecdsa` | 1.0 | `d12.aarch64` | pigsty | 66.8 KiB | [postgresql-15-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-15-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-ecdsa` | 1.0 | `d12.x86_64` | pigsty | 68.2 KiB | [postgresql-15-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-15-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-ecdsa` | 1.0 | `u22.aarch64` | pigsty | 70.4 KiB | [postgresql-15-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-15-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-ecdsa` | 1.0 | `u22.x86_64` | pigsty | 71.4 KiB | [postgresql-15-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-15-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-ecdsa` | 1.0 | `u24.x86_64` | pigsty | 68.5 KiB | [postgresql-15-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-15-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-ecdsa` | 1.0 | `u24.aarch64` | pigsty | 67.4 KiB | [postgresql-15-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-15-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_ecdsa_14` | 1.0 | `el8.x86_64` | pigsty | 27.2 KiB | [pg_ecdsa_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ecdsa_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_ecdsa_14` | 1.0 | `el8.aarch64` | pigsty | 26.6 KiB | [pg_ecdsa_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ecdsa_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_ecdsa_14` | 1.0 | `el9.x86_64` | pigsty | 27.9 KiB | [pg_ecdsa_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ecdsa_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_ecdsa_14` | 1.0 | `el9.aarch64` | pigsty | 27.3 KiB | [pg_ecdsa_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ecdsa_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-ecdsa` | 1.0 | `d12.x86_64` | pigsty | 68.1 KiB | [postgresql-14-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-14-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-ecdsa` | 1.0 | `d12.aarch64` | pigsty | 66.8 KiB | [postgresql-14-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-14-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-ecdsa` | 1.0 | `u22.x86_64` | pigsty | 71.4 KiB | [postgresql-14-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-14-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-ecdsa` | 1.0 | `u22.aarch64` | pigsty | 70.4 KiB | [postgresql-14-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-14-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-ecdsa` | 1.0 | `u24.x86_64` | pigsty | 68.4 KiB | [postgresql-14-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-14-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-ecdsa` | 1.0 | `u24.aarch64` | pigsty | 67.4 KiB | [postgresql-14-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-14-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_ecdsa_13` | 1.0 | `el8.aarch64` | pigsty | 26.6 KiB | [pg_ecdsa_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ecdsa_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_ecdsa_13` | 1.0 | `el8.x86_64` | pigsty | 27.1 KiB | [pg_ecdsa_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ecdsa_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_ecdsa_13` | 1.0 | `el9.aarch64` | pigsty | 27.3 KiB | [pg_ecdsa_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ecdsa_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_ecdsa_13` | 1.0 | `el9.x86_64` | pigsty | 27.8 KiB | [pg_ecdsa_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ecdsa_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-ecdsa` | 1.0 | `d12.aarch64` | pigsty | 66.7 KiB | [postgresql-13-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-13-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-ecdsa` | 1.0 | `d12.x86_64` | pigsty | 68.1 KiB | [postgresql-13-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-13-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-ecdsa` | 1.0 | `u22.aarch64` | pigsty | 70.4 KiB | [postgresql-13-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-13-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-ecdsa` | 1.0 | `u22.x86_64` | pigsty | 71.4 KiB | [postgresql-13-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-13-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-ecdsa` | 1.0 | `u24.aarch64` | pigsty | 67.3 KiB | [postgresql-13-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-13-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-ecdsa` | 1.0 | `u24.x86_64` | pigsty | 68.4 KiB | [postgresql-13-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-13-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ameensol/pg-ecdsa" title="Repository" icon="github" subtitle="github.com/ameensol/pg-ecdsa" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg-ecdsa-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pguecc; # get pguecc source code
pig build dep pguecc; # install build dependencies
pig build pkg pguecc; # build extension rpm or deb
pig build ext pguecc; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pguecc; # install by extension name, for the current active PG version
pig ext install pg_ecdsa; # install via package alias, for the active PG version
pig ext install pguecc -v 18;   # install for PG 18
pig ext install pguecc -v 17;   # install for PG 17
pig ext install pguecc -v 16;   # install for PG 16
pig ext install pguecc -v 15;   # install for PG 15
pig ext install pguecc -v 14;   # install for PG 14
pig ext install pguecc -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pguecc;
```

