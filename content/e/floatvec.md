---
title: "floatvec"
linkTitle: "floatvec"
description: "Math for vectors (arrays) of numbers"
weight: 4730
categories: ["Func"]
width: full
---

Math for vectors (arrays) of numbers

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4730** | {{< badge content="floatvec" link="https://github.com/pjungwir/floatvec" >}} | {{< ext "floatvec" "floatvec" >}} | `1.1.1` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_rational" >}} {{< ext "uint" >}} {{< ext "uint128" >}} {{< ext "numeral" >}} {{< ext "aggs_for_vecs" >}} {{< ext "aggs_for_arrays" >}} {{< ext "arraymath" >}} {{< ext "financial" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/floatvec" >}} | `1.1.1` | {{< badge content="18" color="red" alt="floatvec_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `floatvec_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/floatvec" >}} | `1.1.1` | {{< badge content="18" color="red" alt="postgresql-18-floatvec" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-floatvec` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "floatvec_18" >}}     | {{< pkg "floatvec_17" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_17-1.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "floatvec_16" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_16-1.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "floatvec_15" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_15-1.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "floatvec_14" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_14-1.0.1-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "floatvec_18" >}}     | {{< pkg "floatvec_17" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_17-1.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "floatvec_16" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_16-1.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "floatvec_15" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_15-1.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "floatvec_14" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_14-1.0.1-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "floatvec_18" >}}     | {{< pkg "floatvec_17" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_17-1.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "floatvec_16" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_16-1.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "floatvec_15" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_15-1.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "floatvec_14" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_14-1.0.1-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "floatvec_18" >}}     | {{< pkg "floatvec_17" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_17-1.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "floatvec_16" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_16-1.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "floatvec_15" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_15-1.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "floatvec_14" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_14-1.0.1-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-floatvec" >}}     | {{< pkg "postgresql-17-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-floatvec" >}}     | {{< pkg "postgresql-17-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-floatvec" >}}     | {{< pkg "postgresql-17-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-floatvec" >}}     | {{< pkg "postgresql-17-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-floatvec" >}}     | {{< pkg "postgresql-17-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-floatvec" >}}     | {{< pkg "postgresql-17-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-floatvec" "1.1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `floatvec_17` | 1.1.1 | `el8.aarch64` | pigsty | 18.9 KiB | [floatvec_17-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_17-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `floatvec_17` | 1.1.1 | `el8.x86_64` | pigsty | 17.8 KiB | [floatvec_17-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_17-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `floatvec_17` | 1.0.1 | `el8.aarch64` | pigsty | 18.4 KiB | [floatvec_17-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_17-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `floatvec_17` | 1.0.1 | `el8.x86_64` | pigsty | 17.3 KiB | [floatvec_17-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_17-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `floatvec_17` | 1.1.1 | `el9.x86_64` | pigsty | 18.0 KiB | [floatvec_17-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_17-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `floatvec_17` | 1.1.1 | `el9.aarch64` | pigsty | 19.1 KiB | [floatvec_17-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_17-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `floatvec_17` | 1.0.1 | `el9.aarch64` | pigsty | 18.6 KiB | [floatvec_17-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_17-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `floatvec_17` | 1.0.1 | `el9.x86_64` | pigsty | 17.4 KiB | [floatvec_17-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_17-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-floatvec` | 1.1.1 | `d12.aarch64` | pigsty | 26.3 KiB | [postgresql-17-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-floatvec` | 1.1.1 | `d12.x86_64` | pigsty | 24.8 KiB | [postgresql-17-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-floatvec` | 1.1.1 | `u22.aarch64` | pigsty | 27.4 KiB | [postgresql-17-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-floatvec` | 1.1.1 | `u22.x86_64` | pigsty | 25.8 KiB | [postgresql-17-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-floatvec` | 1.1.1 | `u24.x86_64` | pigsty | 25.5 KiB | [postgresql-17-floatvec_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-floatvec` | 1.1.1 | `u24.aarch64` | pigsty | 27.1 KiB | [postgresql-17-floatvec_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `floatvec_16` | 1.1.1 | `el8.x86_64` | pigsty | 17.8 KiB | [floatvec_16-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_16-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `floatvec_16` | 1.1.1 | `el8.aarch64` | pigsty | 18.9 KiB | [floatvec_16-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_16-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `floatvec_16` | 1.0.1 | `el8.x86_64` | pigsty | 17.3 KiB | [floatvec_16-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_16-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `floatvec_16` | 1.0.1 | `el8.aarch64` | pigsty | 18.4 KiB | [floatvec_16-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_16-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `floatvec_16` | 1.1.1 | `el9.x86_64` | pigsty | 18.1 KiB | [floatvec_16-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_16-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `floatvec_16` | 1.1.1 | `el9.aarch64` | pigsty | 19.1 KiB | [floatvec_16-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_16-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `floatvec_16` | 1.0.1 | `el9.aarch64` | pigsty | 18.6 KiB | [floatvec_16-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_16-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `floatvec_16` | 1.0.1 | `el9.x86_64` | pigsty | 17.5 KiB | [floatvec_16-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_16-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-16-floatvec` | 1.1.1 | `d12.x86_64` | pigsty | 24.8 KiB | [postgresql-16-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-floatvec` | 1.1.1 | `d12.aarch64` | pigsty | 26.3 KiB | [postgresql-16-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-floatvec` | 1.1.1 | `u22.aarch64` | pigsty | 27.4 KiB | [postgresql-16-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-floatvec` | 1.1.1 | `u22.x86_64` | pigsty | 25.8 KiB | [postgresql-16-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-floatvec` | 1.1.1 | `u24.aarch64` | pigsty | 27.1 KiB | [postgresql-16-floatvec_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-floatvec` | 1.1.1 | `u24.x86_64` | pigsty | 25.5 KiB | [postgresql-16-floatvec_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `floatvec_15` | 1.1.1 | `el8.aarch64` | pigsty | 18.6 KiB | [floatvec_15-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_15-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `floatvec_15` | 1.1.1 | `el8.x86_64` | pigsty | 17.9 KiB | [floatvec_15-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_15-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `floatvec_15` | 1.0.1 | `el8.aarch64` | pigsty | 18.1 KiB | [floatvec_15-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_15-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `floatvec_15` | 1.0.1 | `el8.x86_64` | pigsty | 17.3 KiB | [floatvec_15-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_15-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `floatvec_15` | 1.1.1 | `el9.aarch64` | pigsty | 18.7 KiB | [floatvec_15-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_15-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `floatvec_15` | 1.1.1 | `el9.x86_64` | pigsty | 18.0 KiB | [floatvec_15-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_15-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `floatvec_15` | 1.0.1 | `el9.aarch64` | pigsty | 18.2 KiB | [floatvec_15-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_15-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `floatvec_15` | 1.0.1 | `el9.x86_64` | pigsty | 17.4 KiB | [floatvec_15-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_15-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-15-floatvec` | 1.1.1 | `d12.x86_64` | pigsty | 24.9 KiB | [postgresql-15-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-floatvec` | 1.1.1 | `d12.aarch64` | pigsty | 25.9 KiB | [postgresql-15-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-floatvec` | 1.1.1 | `u22.x86_64` | pigsty | 25.7 KiB | [postgresql-15-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-floatvec` | 1.1.1 | `u22.aarch64` | pigsty | 26.8 KiB | [postgresql-15-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-floatvec` | 1.1.1 | `u24.x86_64` | pigsty | 24.8 KiB | [postgresql-15-floatvec_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-floatvec` | 1.1.1 | `u24.aarch64` | pigsty | 26.0 KiB | [postgresql-15-floatvec_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `floatvec_14` | 1.1.1 | `el8.aarch64` | pigsty | 18.6 KiB | [floatvec_14-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_14-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `floatvec_14` | 1.1.1 | `el8.x86_64` | pigsty | 17.9 KiB | [floatvec_14-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_14-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `floatvec_14` | 1.0.1 | `el8.x86_64` | pigsty | 17.3 KiB | [floatvec_14-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_14-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `floatvec_14` | 1.0.1 | `el8.aarch64` | pigsty | 18.1 KiB | [floatvec_14-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_14-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `floatvec_14` | 1.1.1 | `el9.aarch64` | pigsty | 18.6 KiB | [floatvec_14-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_14-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `floatvec_14` | 1.1.1 | `el9.x86_64` | pigsty | 18.0 KiB | [floatvec_14-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_14-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `floatvec_14` | 1.0.1 | `el9.aarch64` | pigsty | 18.2 KiB | [floatvec_14-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_14-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `floatvec_14` | 1.0.1 | `el9.x86_64` | pigsty | 17.4 KiB | [floatvec_14-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_14-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-14-floatvec` | 1.1.1 | `d12.aarch64` | pigsty | 25.9 KiB | [postgresql-14-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-floatvec` | 1.1.1 | `d12.x86_64` | pigsty | 24.8 KiB | [postgresql-14-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-floatvec` | 1.1.1 | `u22.x86_64` | pigsty | 25.7 KiB | [postgresql-14-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-floatvec` | 1.1.1 | `u22.aarch64` | pigsty | 26.8 KiB | [postgresql-14-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-floatvec` | 1.1.1 | `u24.x86_64` | pigsty | 25.2 KiB | [postgresql-14-floatvec_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-floatvec` | 1.1.1 | `u24.aarch64` | pigsty | 26.3 KiB | [postgresql-14-floatvec_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `floatvec_13` | 1.1.1 | `el8.aarch64` | pigsty | 18.5 KiB | [floatvec_13-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_13-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `floatvec_13` | 1.1.1 | `el8.x86_64` | pigsty | 17.7 KiB | [floatvec_13-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_13-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `floatvec_13` | 1.0.1 | `el8.aarch64` | pigsty | 18.1 KiB | [floatvec_13-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_13-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `floatvec_13` | 1.0.1 | `el8.x86_64` | pigsty | 17.1 KiB | [floatvec_13-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_13-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `floatvec_13` | 1.1.1 | `el9.aarch64` | pigsty | 18.6 KiB | [floatvec_13-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_13-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `floatvec_13` | 1.1.1 | `el9.x86_64` | pigsty | 18.0 KiB | [floatvec_13-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_13-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `floatvec_13` | 1.0.1 | `el9.x86_64` | pigsty | 17.4 KiB | [floatvec_13-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_13-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `floatvec_13` | 1.0.1 | `el9.aarch64` | pigsty | 18.2 KiB | [floatvec_13-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_13-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-floatvec` | 1.1.1 | `d12.aarch64` | pigsty | 25.5 KiB | [postgresql-13-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-13-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-floatvec` | 1.1.1 | `d12.x86_64` | pigsty | 24.4 KiB | [postgresql-13-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-13-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-floatvec` | 1.1.1 | `u22.x86_64` | pigsty | 25.7 KiB | [postgresql-13-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-13-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-floatvec` | 1.1.1 | `u22.aarch64` | pigsty | 26.5 KiB | [postgresql-13-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-13-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-floatvec` | 1.1.1 | `u24.x86_64` | pigsty | 24.8 KiB | [postgresql-13-floatvec_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-13-floatvec_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-floatvec` | 1.1.1 | `u24.aarch64` | pigsty | 26.0 KiB | [postgresql-13-floatvec_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-13-floatvec_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pjungwir/floatvec" title="Repository" icon="github" subtitle="github.com/pjungwir/floatvec" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="floatvec-1.1.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get floatvec; # get floatvec source code
pig build dep floatvec; # install build dependencies
pig build pkg floatvec; # build extension rpm or deb
pig build ext floatvec; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install floatvec; # install by extension name, for the current active PG version
pig ext install floatvec; # install via package alias, for the active PG version
pig ext install floatvec -v 18;   # install for PG 18
pig ext install floatvec -v 17;   # install for PG 17
pig ext install floatvec -v 16;   # install for PG 16
pig ext install floatvec -v 15;   # install for PG 15
pig ext install floatvec -v 14;   # install for PG 14
pig ext install floatvec -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION floatvec;
```

