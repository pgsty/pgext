---
title: "arraymath"
linkTitle: "arraymath"
description: "Array math and operators that work element by element on the contents of arrays"
weight: 4760
categories: ["Func"]
width: full
---

Array math and operators that work element by element on the contents of arrays

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4760** | {{< badge content="arraymath" link="https://github.com/pramsey/pgsql-arraymath" >}} | {{< ext "arraymath" "pg_arraymath" >}} | `1.1` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "aggs_for_arrays" >}} {{< ext "aggs_for_vecs" >}} {{< ext "intarray" >}} {{< ext "first_last_agg" >}} {{< ext "floatvec" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/arraymath" >}} | `1.1` | {{< badge content="18" color="red" alt="pg_arraymath_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_arraymath_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/arraymath" >}} | `1.1` | {{< badge content="18" color="red" alt="postgresql-18-pg-arraymath" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-arraymath` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_arraymath_18" >}}     | {{< pkg "pg_arraymath_17" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_arraymath_17-1.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_arraymath_16" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_arraymath_16-1.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_arraymath_15" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_arraymath_15-1.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_arraymath_14" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_arraymath_14-1.1-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_arraymath_18" >}}     | {{< pkg "pg_arraymath_17" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_arraymath_17-1.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_arraymath_16" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_arraymath_16-1.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_arraymath_15" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_arraymath_15-1.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_arraymath_14" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_arraymath_14-1.1-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_arraymath_18" >}}     | {{< pkg "pg_arraymath_17" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_arraymath_17-1.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_arraymath_16" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_arraymath_16-1.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_arraymath_15" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_arraymath_15-1.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_arraymath_14" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_arraymath_14-1.1-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_arraymath_18" >}}     | {{< pkg "pg_arraymath_17" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_arraymath_17-1.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_arraymath_16" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_arraymath_16-1.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_arraymath_15" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_arraymath_15-1.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_arraymath_14" "1.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_arraymath_14-1.1-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-arraymath" >}}     | {{< pkg "postgresql-17-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-arraymath" >}}     | {{< pkg "postgresql-17-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-arraymath" >}}     | {{< pkg "postgresql-17-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-arraymath" >}}     | {{< pkg "postgresql-17-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-arraymath" >}}     | {{< pkg "postgresql-17-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-arraymath" >}}     | {{< pkg "postgresql-17-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-arraymath" "1.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_arraymath_17` | 1.1 | `el8.x86_64` | pigsty | 19.2 KiB | [pg_arraymath_17-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_arraymath_17-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_arraymath_17` | 1.1 | `el8.aarch64` | pigsty | 18.7 KiB | [pg_arraymath_17-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_arraymath_17-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_arraymath_17` | 1.1 | `el9.aarch64` | pigsty | 18.8 KiB | [pg_arraymath_17-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_arraymath_17-1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_arraymath_17` | 1.1 | `el9.x86_64` | pigsty | 19.5 KiB | [pg_arraymath_17-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_arraymath_17-1.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-arraymath` | 1.1 | `d12.x86_64` | pigsty | 26.3 KiB | [postgresql-17-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-arraymath` | 1.1 | `d12.aarch64` | pigsty | 26.3 KiB | [postgresql-17-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-arraymath` | 1.1 | `u22.x86_64` | pigsty | 27.9 KiB | [postgresql-17-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-arraymath` | 1.1 | `u22.aarch64` | pigsty | 27.6 KiB | [postgresql-17-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-arraymath` | 1.1 | `u24.x86_64` | pigsty | 26.4 KiB | [postgresql-17-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-arraymath` | 1.1 | `u24.aarch64` | pigsty | 26.4 KiB | [postgresql-17-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_arraymath_16` | 1.1 | `el8.x86_64` | pigsty | 19.2 KiB | [pg_arraymath_16-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_arraymath_16-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_arraymath_16` | 1.1 | `el8.aarch64` | pigsty | 18.7 KiB | [pg_arraymath_16-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_arraymath_16-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_arraymath_16` | 1.1 | `el9.x86_64` | pigsty | 19.5 KiB | [pg_arraymath_16-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_arraymath_16-1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_arraymath_16` | 1.1 | `el9.aarch64` | pigsty | 18.9 KiB | [pg_arraymath_16-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_arraymath_16-1.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-arraymath` | 1.1 | `d12.x86_64` | pigsty | 26.3 KiB | [postgresql-16-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-arraymath` | 1.1 | `d12.aarch64` | pigsty | 26.3 KiB | [postgresql-16-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-arraymath` | 1.1 | `u22.aarch64` | pigsty | 27.6 KiB | [postgresql-16-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-arraymath` | 1.1 | `u22.x86_64` | pigsty | 27.9 KiB | [postgresql-16-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-arraymath` | 1.1 | `u24.x86_64` | pigsty | 26.5 KiB | [postgresql-16-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-arraymath` | 1.1 | `u24.aarch64` | pigsty | 26.4 KiB | [postgresql-16-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_arraymath_15` | 1.1 | `el8.x86_64` | pigsty | 19.1 KiB | [pg_arraymath_15-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_arraymath_15-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_arraymath_15` | 1.1 | `el8.aarch64` | pigsty | 18.6 KiB | [pg_arraymath_15-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_arraymath_15-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_arraymath_15` | 1.1 | `el9.x86_64` | pigsty | 19.5 KiB | [pg_arraymath_15-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_arraymath_15-1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_arraymath_15` | 1.1 | `el9.aarch64` | pigsty | 18.8 KiB | [pg_arraymath_15-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_arraymath_15-1.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-arraymath` | 1.1 | `d12.aarch64` | pigsty | 25.9 KiB | [postgresql-15-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-arraymath` | 1.1 | `d12.x86_64` | pigsty | 26.1 KiB | [postgresql-15-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-arraymath` | 1.1 | `u22.aarch64` | pigsty | 27.3 KiB | [postgresql-15-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-arraymath` | 1.1 | `u22.x86_64` | pigsty | 27.5 KiB | [postgresql-15-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-arraymath` | 1.1 | `u24.x86_64` | pigsty | 26.2 KiB | [postgresql-15-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-arraymath` | 1.1 | `u24.aarch64` | pigsty | 26.1 KiB | [postgresql-15-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_arraymath_14` | 1.1 | `el8.x86_64` | pigsty | 19.1 KiB | [pg_arraymath_14-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_arraymath_14-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_arraymath_14` | 1.1 | `el8.aarch64` | pigsty | 18.6 KiB | [pg_arraymath_14-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_arraymath_14-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_arraymath_14` | 1.1 | `el9.x86_64` | pigsty | 19.4 KiB | [pg_arraymath_14-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_arraymath_14-1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_arraymath_14` | 1.1 | `el9.aarch64` | pigsty | 18.8 KiB | [pg_arraymath_14-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_arraymath_14-1.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-arraymath` | 1.1 | `d12.x86_64` | pigsty | 26.0 KiB | [postgresql-14-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-arraymath` | 1.1 | `d12.aarch64` | pigsty | 25.9 KiB | [postgresql-14-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-arraymath` | 1.1 | `u22.x86_64` | pigsty | 27.4 KiB | [postgresql-14-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-arraymath` | 1.1 | `u22.aarch64` | pigsty | 27.3 KiB | [postgresql-14-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-arraymath` | 1.1 | `u24.x86_64` | pigsty | 26.1 KiB | [postgresql-14-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-arraymath` | 1.1 | `u24.aarch64` | pigsty | 26.1 KiB | [postgresql-14-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_arraymath_13` | 1.1 | `el8.aarch64` | pigsty | 18.5 KiB | [pg_arraymath_13-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_arraymath_13-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_arraymath_13` | 1.1 | `el8.x86_64` | pigsty | 18.7 KiB | [pg_arraymath_13-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_arraymath_13-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_arraymath_13` | 1.1 | `el9.aarch64` | pigsty | 18.7 KiB | [pg_arraymath_13-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_arraymath_13-1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_arraymath_13` | 1.1 | `el9.x86_64` | pigsty | 19.4 KiB | [pg_arraymath_13-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_arraymath_13-1.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-arraymath` | 1.1 | `d12.aarch64` | pigsty | 25.3 KiB | [postgresql-13-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-13-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-arraymath` | 1.1 | `d12.x86_64` | pigsty | 25.8 KiB | [postgresql-13-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-13-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-arraymath` | 1.1 | `u22.aarch64` | pigsty | 26.6 KiB | [postgresql-13-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-13-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-arraymath` | 1.1 | `u22.x86_64` | pigsty | 27.2 KiB | [postgresql-13-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-13-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-arraymath` | 1.1 | `u24.aarch64` | pigsty | 25.4 KiB | [postgresql-13-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-13-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-arraymath` | 1.1 | `u24.x86_64` | pigsty | 26.0 KiB | [postgresql-13-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-13-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pramsey/pgsql-arraymath" title="Repository" icon="github" subtitle="github.com/pramsey/pgsql-arraymath" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pgsql-arraymath-1.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get arraymath; # get arraymath source code
pig build dep arraymath; # install build dependencies
pig build pkg arraymath; # build extension rpm or deb
pig build ext arraymath; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install arraymath; # install by extension name, for the current active PG version
pig ext install pg_arraymath; # install via package alias, for the active PG version
pig ext install arraymath -v 18;   # install for PG 18
pig ext install arraymath -v 17;   # install for PG 17
pig ext install arraymath -v 16;   # install for PG 16
pig ext install arraymath -v 15;   # install for PG 15
pig ext install arraymath -v 14;   # install for PG 14
pig ext install arraymath -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION arraymath;
```

