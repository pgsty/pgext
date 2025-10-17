---
title: "cryptint"
linkTitle: "cryptint"
description: "Encryption functions for int and bigint values"
weight: 4450
categories: ["Util"]
width: full
---

Encryption functions for int and bigint values

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4450** | {{< badge content="cryptint" link="https://github.com/dverite/cryptint" >}} | {{< ext "cryptint" "cryptint" >}} | `1.0.0` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "hashlib" >}} {{< ext "shacrypt" >}} {{< ext "pguecc" >}} {{< ext "pgcrypto" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/cryptint" >}} | `1.0.0` | {{< badge content="18" color="red" alt="cryptint_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `cryptint_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/cryptint" >}} | `1.0.0` | {{< badge content="18" color="red" alt="postgresql-18-cryptint" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-cryptint` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "cryptint_18" >}}     | {{< pkg "cryptint_17" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/cryptint_17-1.0.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "cryptint_16" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/cryptint_16-1.0.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "cryptint_15" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/cryptint_15-1.0.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "cryptint_14" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/cryptint_14-1.0.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "cryptint_18" >}}     | {{< pkg "cryptint_17" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/cryptint_17-1.0.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "cryptint_16" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/cryptint_16-1.0.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "cryptint_15" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/cryptint_15-1.0.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "cryptint_14" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/cryptint_14-1.0.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "cryptint_18" >}}     | {{< pkg "cryptint_17" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/cryptint_17-1.0.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "cryptint_16" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/cryptint_16-1.0.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "cryptint_15" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/cryptint_15-1.0.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "cryptint_14" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/cryptint_14-1.0.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "cryptint_18" >}}     | {{< pkg "cryptint_17" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/cryptint_17-1.0.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "cryptint_16" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/cryptint_16-1.0.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "cryptint_15" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/cryptint_15-1.0.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "cryptint_14" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/cryptint_14-1.0.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-cryptint" >}}     | {{< pkg "postgresql-17-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cryptint/postgresql-17-cryptint_1.0.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cryptint/postgresql-16-cryptint_1.0.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cryptint/postgresql-15-cryptint_1.0.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cryptint/postgresql-14-cryptint_1.0.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-cryptint" >}}     | {{< pkg "postgresql-17-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cryptint/postgresql-17-cryptint_1.0.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cryptint/postgresql-16-cryptint_1.0.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cryptint/postgresql-15-cryptint_1.0.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cryptint/postgresql-14-cryptint_1.0.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-cryptint" >}}     | {{< pkg "postgresql-17-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cryptint/postgresql-17-cryptint_1.0.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cryptint/postgresql-16-cryptint_1.0.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cryptint/postgresql-15-cryptint_1.0.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cryptint/postgresql-14-cryptint_1.0.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-cryptint" >}}     | {{< pkg "postgresql-17-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cryptint/postgresql-17-cryptint_1.0.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cryptint/postgresql-16-cryptint_1.0.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cryptint/postgresql-15-cryptint_1.0.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cryptint/postgresql-14-cryptint_1.0.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-cryptint" >}}     | {{< pkg "postgresql-17-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cryptint/postgresql-17-cryptint_1.0.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cryptint/postgresql-16-cryptint_1.0.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cryptint/postgresql-15-cryptint_1.0.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cryptint/postgresql-14-cryptint_1.0.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-cryptint" >}}     | {{< pkg "postgresql-17-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cryptint/postgresql-17-cryptint_1.0.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cryptint/postgresql-16-cryptint_1.0.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cryptint/postgresql-15-cryptint_1.0.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-cryptint" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cryptint/postgresql-14-cryptint_1.0.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `cryptint_17` | 1.0.0 | `el8.x86_64` | pigsty | 13.9 KiB | [cryptint_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cryptint_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `cryptint_17` | 1.0.0 | `el8.aarch64` | pigsty | 14.0 KiB | [cryptint_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cryptint_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `cryptint_17` | 1.0.0 | `el9.aarch64` | pigsty | 13.7 KiB | [cryptint_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cryptint_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `cryptint_17` | 1.0.0 | `el9.x86_64` | pigsty | 13.9 KiB | [cryptint_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cryptint_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-cryptint` | 1.0.0 | `d12.x86_64` | pigsty | 13.4 KiB | [postgresql-17-cryptint_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cryptint/postgresql-17-cryptint_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-cryptint` | 1.0.0 | `d12.aarch64` | pigsty | 13.5 KiB | [postgresql-17-cryptint_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cryptint/postgresql-17-cryptint_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-cryptint` | 1.0.0 | `u22.x86_64` | pigsty | 13.8 KiB | [postgresql-17-cryptint_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cryptint/postgresql-17-cryptint_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-cryptint` | 1.0.0 | `u22.aarch64` | pigsty | 13.8 KiB | [postgresql-17-cryptint_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cryptint/postgresql-17-cryptint_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-cryptint` | 1.0.0 | `u24.x86_64` | pigsty | 13.7 KiB | [postgresql-17-cryptint_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cryptint/postgresql-17-cryptint_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-cryptint` | 1.0.0 | `u24.aarch64` | pigsty | 13.6 KiB | [postgresql-17-cryptint_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cryptint/postgresql-17-cryptint_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `cryptint_16` | 1.0.0 | `el8.x86_64` | pigsty | 13.9 KiB | [cryptint_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cryptint_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `cryptint_16` | 1.0.0 | `el8.aarch64` | pigsty | 13.9 KiB | [cryptint_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cryptint_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `cryptint_16` | 1.0.0 | `el9.x86_64` | pigsty | 13.9 KiB | [cryptint_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cryptint_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `cryptint_16` | 1.0.0 | `el9.aarch64` | pigsty | 13.7 KiB | [cryptint_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cryptint_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-cryptint` | 1.0.0 | `d12.x86_64` | pigsty | 13.4 KiB | [postgresql-16-cryptint_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cryptint/postgresql-16-cryptint_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-cryptint` | 1.0.0 | `d12.aarch64` | pigsty | 13.5 KiB | [postgresql-16-cryptint_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cryptint/postgresql-16-cryptint_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-cryptint` | 1.0.0 | `u22.aarch64` | pigsty | 13.9 KiB | [postgresql-16-cryptint_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cryptint/postgresql-16-cryptint_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-cryptint` | 1.0.0 | `u22.x86_64` | pigsty | 13.8 KiB | [postgresql-16-cryptint_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cryptint/postgresql-16-cryptint_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-cryptint` | 1.0.0 | `u24.x86_64` | pigsty | 13.7 KiB | [postgresql-16-cryptint_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cryptint/postgresql-16-cryptint_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-cryptint` | 1.0.0 | `u24.aarch64` | pigsty | 13.6 KiB | [postgresql-16-cryptint_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cryptint/postgresql-16-cryptint_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `cryptint_15` | 1.0.0 | `el8.x86_64` | pigsty | 14.0 KiB | [cryptint_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cryptint_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `cryptint_15` | 1.0.0 | `el8.aarch64` | pigsty | 14.0 KiB | [cryptint_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cryptint_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `cryptint_15` | 1.0.0 | `el9.x86_64` | pigsty | 14.1 KiB | [cryptint_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cryptint_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `cryptint_15` | 1.0.0 | `el9.aarch64` | pigsty | 14.0 KiB | [cryptint_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cryptint_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-cryptint` | 1.0.0 | `d12.aarch64` | pigsty | 13.6 KiB | [postgresql-15-cryptint_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cryptint/postgresql-15-cryptint_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-cryptint` | 1.0.0 | `d12.x86_64` | pigsty | 13.5 KiB | [postgresql-15-cryptint_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cryptint/postgresql-15-cryptint_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-cryptint` | 1.0.0 | `u22.aarch64` | pigsty | 14.0 KiB | [postgresql-15-cryptint_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cryptint/postgresql-15-cryptint_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-cryptint` | 1.0.0 | `u22.x86_64` | pigsty | 13.9 KiB | [postgresql-15-cryptint_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cryptint/postgresql-15-cryptint_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-cryptint` | 1.0.0 | `u24.x86_64` | pigsty | 13.9 KiB | [postgresql-15-cryptint_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cryptint/postgresql-15-cryptint_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-cryptint` | 1.0.0 | `u24.aarch64` | pigsty | 13.8 KiB | [postgresql-15-cryptint_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cryptint/postgresql-15-cryptint_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `cryptint_14` | 1.0.0 | `el8.x86_64` | pigsty | 13.9 KiB | [cryptint_14-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cryptint_14-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `cryptint_14` | 1.0.0 | `el8.aarch64` | pigsty | 14.0 KiB | [cryptint_14-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cryptint_14-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `cryptint_14` | 1.0.0 | `el9.x86_64` | pigsty | 14.1 KiB | [cryptint_14-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cryptint_14-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `cryptint_14` | 1.0.0 | `el9.aarch64` | pigsty | 14.0 KiB | [cryptint_14-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cryptint_14-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-cryptint` | 1.0.0 | `d12.x86_64` | pigsty | 13.5 KiB | [postgresql-14-cryptint_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cryptint/postgresql-14-cryptint_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-cryptint` | 1.0.0 | `d12.aarch64` | pigsty | 13.6 KiB | [postgresql-14-cryptint_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cryptint/postgresql-14-cryptint_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-cryptint` | 1.0.0 | `u22.x86_64` | pigsty | 13.9 KiB | [postgresql-14-cryptint_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cryptint/postgresql-14-cryptint_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-cryptint` | 1.0.0 | `u22.aarch64` | pigsty | 14.0 KiB | [postgresql-14-cryptint_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cryptint/postgresql-14-cryptint_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-cryptint` | 1.0.0 | `u24.x86_64` | pigsty | 13.9 KiB | [postgresql-14-cryptint_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cryptint/postgresql-14-cryptint_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-cryptint` | 1.0.0 | `u24.aarch64` | pigsty | 13.8 KiB | [postgresql-14-cryptint_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cryptint/postgresql-14-cryptint_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `cryptint_13` | 1.0.0 | `el8.aarch64` | pigsty | 14.0 KiB | [cryptint_13-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cryptint_13-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `cryptint_13` | 1.0.0 | `el8.x86_64` | pigsty | 13.9 KiB | [cryptint_13-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cryptint_13-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `cryptint_13` | 1.0.0 | `el9.aarch64` | pigsty | 14.0 KiB | [cryptint_13-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cryptint_13-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `cryptint_13` | 1.0.0 | `el9.x86_64` | pigsty | 14.1 KiB | [cryptint_13-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cryptint_13-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-cryptint` | 1.0.0 | `d12.aarch64` | pigsty | 13.4 KiB | [postgresql-13-cryptint_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cryptint/postgresql-13-cryptint_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-cryptint` | 1.0.0 | `d12.x86_64` | pigsty | 13.5 KiB | [postgresql-13-cryptint_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cryptint/postgresql-13-cryptint_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-cryptint` | 1.0.0 | `u22.aarch64` | pigsty | 14.0 KiB | [postgresql-13-cryptint_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cryptint/postgresql-13-cryptint_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-cryptint` | 1.0.0 | `u22.x86_64` | pigsty | 13.9 KiB | [postgresql-13-cryptint_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cryptint/postgresql-13-cryptint_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-cryptint` | 1.0.0 | `u24.aarch64` | pigsty | 13.6 KiB | [postgresql-13-cryptint_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cryptint/postgresql-13-cryptint_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-cryptint` | 1.0.0 | `u24.x86_64` | pigsty | 13.9 KiB | [postgresql-13-cryptint_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cryptint/postgresql-13-cryptint_1.0.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/dverite/cryptint" title="Repository" icon="github" subtitle="github.com/dverite/cryptint" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="cryptint-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get cryptint; # get cryptint source code
pig build dep cryptint; # install build dependencies
pig build pkg cryptint; # build extension rpm or deb
pig build ext cryptint; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install cryptint; # install by extension name, for the current active PG version
pig ext install cryptint; # install via package alias, for the active PG version
pig ext install cryptint -v 18;   # install for PG 18
pig ext install cryptint -v 17;   # install for PG 17
pig ext install cryptint -v 16;   # install for PG 16
pig ext install cryptint -v 15;   # install for PG 15
pig ext install cryptint -v 14;   # install for PG 14
pig ext install cryptint -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION cryptint;
```

