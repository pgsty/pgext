---
title: "aggs_for_arrays"
linkTitle: "aggs_for_arrays"
description: "Various functions for computing statistics on arrays of numbers"
weight: 4750
categories: ["Func"]
width: full
---

Various functions for computing statistics on arrays of numbers

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4750** | {{< badge content="aggs_for_arrays" link="https://github.com/pjungwir/aggs_for_arrays" >}} | {{< ext "aggs_for_arrays" "aggs_for_arrays" >}} | `1.3.3` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "aggs_for_vecs" >}} {{< ext "first_last_agg" >}} {{< ext "arraymath" >}} {{< ext "intarray" >}} {{< ext "topn" >}} {{< ext "quantile" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/aggs_for_arrays" >}} | `1.3.3` | {{< badge content="18" color="red" alt="aggs_for_arrays_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `aggs_for_arrays_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/aggs_for_arrays" >}} | `1.3.3` | {{< badge content="18" color="red" alt="postgresql-18-aggs-for-arrays" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-aggs-for-arrays` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "aggs_for_arrays_18" >}}     | {{< pkg "aggs_for_arrays_17" "1.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_arrays_17-1.3.3-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "aggs_for_arrays_16" "1.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_arrays_16-1.3.3-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "aggs_for_arrays_15" "1.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_arrays_15-1.3.3-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "aggs_for_arrays_14" "1.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_arrays_14-1.3.3-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "aggs_for_arrays_18" >}}     | {{< pkg "aggs_for_arrays_17" "1.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_arrays_17-1.3.3-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "aggs_for_arrays_16" "1.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_arrays_16-1.3.3-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "aggs_for_arrays_15" "1.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_arrays_15-1.3.3-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "aggs_for_arrays_14" "1.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_arrays_14-1.3.3-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "aggs_for_arrays_18" >}}     | {{< pkg "aggs_for_arrays_17" "1.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_arrays_17-1.3.3-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "aggs_for_arrays_16" "1.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_arrays_16-1.3.3-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "aggs_for_arrays_15" "1.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_arrays_15-1.3.3-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "aggs_for_arrays_14" "1.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_arrays_14-1.3.3-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "aggs_for_arrays_18" >}}     | {{< pkg "aggs_for_arrays_17" "1.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_arrays_17-1.3.3-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "aggs_for_arrays_16" "1.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_arrays_16-1.3.3-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "aggs_for_arrays_15" "1.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_arrays_15-1.3.3-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "aggs_for_arrays_14" "1.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_arrays_14-1.3.3-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-aggs-for-arrays" >}}     | {{< pkg "postgresql-17-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-aggs-for-arrays" >}}     | {{< pkg "postgresql-17-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-aggs-for-arrays" >}}     | {{< pkg "postgresql-17-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-aggs-for-arrays" >}}     | {{< pkg "postgresql-17-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-aggs-for-arrays" >}}     | {{< pkg "postgresql-17-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-aggs-for-arrays" >}}     | {{< pkg "postgresql-17-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-aggs-for-arrays" "1.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `aggs_for_arrays_17` | 1.3.3 | `el8.x86_64` | pigsty | 27.0 KiB | [aggs_for_arrays_17-1.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_arrays_17-1.3.3-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_arrays_17` | 1.3.3 | `el8.aarch64` | pigsty | 30.1 KiB | [aggs_for_arrays_17-1.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_arrays_17-1.3.3-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_arrays_17` | 1.3.3 | `el9.aarch64` | pigsty | 31.0 KiB | [aggs_for_arrays_17-1.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_arrays_17-1.3.3-1PIGSTY.el9.aarch64.rpm) |
| `aggs_for_arrays_17` | 1.3.3 | `el9.x86_64` | pigsty | 26.8 KiB | [aggs_for_arrays_17-1.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_arrays_17-1.3.3-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-aggs-for-arrays` | 1.3.3 | `d12.x86_64` | pigsty | 43.6 KiB | [postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-aggs-for-arrays` | 1.3.3 | `d12.aarch64` | pigsty | 48.5 KiB | [postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-aggs-for-arrays` | 1.3.3 | `u22.x86_64` | pigsty | 45.6 KiB | [postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-aggs-for-arrays` | 1.3.3 | `u22.aarch64` | pigsty | 49.9 KiB | [postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-aggs-for-arrays` | 1.3.3 | `u24.x86_64` | pigsty | 43.7 KiB | [postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-aggs-for-arrays` | 1.3.3 | `u24.aarch64` | pigsty | 49.9 KiB | [postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `aggs_for_arrays_16` | 1.3.3 | `el8.x86_64` | pigsty | 27.0 KiB | [aggs_for_arrays_16-1.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_arrays_16-1.3.3-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_arrays_16` | 1.3.3 | `el8.aarch64` | pigsty | 30.1 KiB | [aggs_for_arrays_16-1.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_arrays_16-1.3.3-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_arrays_16` | 1.3.3 | `el9.x86_64` | pigsty | 26.8 KiB | [aggs_for_arrays_16-1.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_arrays_16-1.3.3-1PIGSTY.el9.x86_64.rpm) |
| `aggs_for_arrays_16` | 1.3.3 | `el9.aarch64` | pigsty | 31.0 KiB | [aggs_for_arrays_16-1.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_arrays_16-1.3.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-aggs-for-arrays` | 1.3.3 | `d12.x86_64` | pigsty | 43.6 KiB | [postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-aggs-for-arrays` | 1.3.3 | `d12.aarch64` | pigsty | 48.4 KiB | [postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-aggs-for-arrays` | 1.3.3 | `u22.aarch64` | pigsty | 49.9 KiB | [postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-aggs-for-arrays` | 1.3.3 | `u22.x86_64` | pigsty | 45.6 KiB | [postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-aggs-for-arrays` | 1.3.3 | `u24.x86_64` | pigsty | 43.7 KiB | [postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-aggs-for-arrays` | 1.3.3 | `u24.aarch64` | pigsty | 49.9 KiB | [postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `aggs_for_arrays_15` | 1.3.3 | `el8.x86_64` | pigsty | 27.1 KiB | [aggs_for_arrays_15-1.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_arrays_15-1.3.3-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_arrays_15` | 1.3.3 | `el8.aarch64` | pigsty | 30.3 KiB | [aggs_for_arrays_15-1.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_arrays_15-1.3.3-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_arrays_15` | 1.3.3 | `el9.x86_64` | pigsty | 26.9 KiB | [aggs_for_arrays_15-1.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_arrays_15-1.3.3-1PIGSTY.el9.x86_64.rpm) |
| `aggs_for_arrays_15` | 1.3.3 | `el9.aarch64` | pigsty | 31.2 KiB | [aggs_for_arrays_15-1.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_arrays_15-1.3.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-aggs-for-arrays` | 1.3.3 | `d12.aarch64` | pigsty | 48.6 KiB | [postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-aggs-for-arrays` | 1.3.3 | `d12.x86_64` | pigsty | 43.8 KiB | [postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-aggs-for-arrays` | 1.3.3 | `u22.aarch64` | pigsty | 50.2 KiB | [postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-aggs-for-arrays` | 1.3.3 | `u22.x86_64` | pigsty | 45.7 KiB | [postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-aggs-for-arrays` | 1.3.3 | `u24.x86_64` | pigsty | 44.1 KiB | [postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-aggs-for-arrays` | 1.3.3 | `u24.aarch64` | pigsty | 50.2 KiB | [postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `aggs_for_arrays_14` | 1.3.3 | `el8.x86_64` | pigsty | 27.1 KiB | [aggs_for_arrays_14-1.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_arrays_14-1.3.3-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_arrays_14` | 1.3.3 | `el8.aarch64` | pigsty | 30.3 KiB | [aggs_for_arrays_14-1.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_arrays_14-1.3.3-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_arrays_14` | 1.3.3 | `el9.x86_64` | pigsty | 26.9 KiB | [aggs_for_arrays_14-1.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_arrays_14-1.3.3-1PIGSTY.el9.x86_64.rpm) |
| `aggs_for_arrays_14` | 1.3.3 | `el9.aarch64` | pigsty | 31.1 KiB | [aggs_for_arrays_14-1.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_arrays_14-1.3.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-aggs-for-arrays` | 1.3.3 | `d12.x86_64` | pigsty | 43.7 KiB | [postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-aggs-for-arrays` | 1.3.3 | `d12.aarch64` | pigsty | 48.6 KiB | [postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-aggs-for-arrays` | 1.3.3 | `u22.x86_64` | pigsty | 45.7 KiB | [postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-aggs-for-arrays` | 1.3.3 | `u22.aarch64` | pigsty | 50.1 KiB | [postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-aggs-for-arrays` | 1.3.3 | `u24.x86_64` | pigsty | 44.1 KiB | [postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-aggs-for-arrays` | 1.3.3 | `u24.aarch64` | pigsty | 50.2 KiB | [postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `aggs_for_arrays_13` | 1.3.3 | `el8.aarch64` | pigsty | 30.3 KiB | [aggs_for_arrays_13-1.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_arrays_13-1.3.3-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_arrays_13` | 1.3.3 | `el8.x86_64` | pigsty | 26.8 KiB | [aggs_for_arrays_13-1.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_arrays_13-1.3.3-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_arrays_13` | 1.3.3 | `el9.aarch64` | pigsty | 31.1 KiB | [aggs_for_arrays_13-1.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_arrays_13-1.3.3-1PIGSTY.el9.aarch64.rpm) |
| `aggs_for_arrays_13` | 1.3.3 | `el9.x86_64` | pigsty | 26.7 KiB | [aggs_for_arrays_13-1.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_arrays_13-1.3.3-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-aggs-for-arrays` | 1.3.3 | `d12.aarch64` | pigsty | 48.5 KiB | [postgresql-13-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-13-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-aggs-for-arrays` | 1.3.3 | `d12.x86_64` | pigsty | 43.5 KiB | [postgresql-13-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-13-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-aggs-for-arrays` | 1.3.3 | `u22.aarch64` | pigsty | 50.0 KiB | [postgresql-13-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-13-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-aggs-for-arrays` | 1.3.3 | `u22.x86_64` | pigsty | 45.6 KiB | [postgresql-13-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-13-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-aggs-for-arrays` | 1.3.3 | `u24.aarch64` | pigsty | 49.9 KiB | [postgresql-13-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-13-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-aggs-for-arrays` | 1.3.3 | `u24.x86_64` | pigsty | 43.9 KiB | [postgresql-13-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-13-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pjungwir/aggs_for_arrays" title="Repository" icon="github" subtitle="github.com/pjungwir/aggs_for_arrays" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="aggs_for_arrays-1.3.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get aggs_for_arrays; # get aggs_for_arrays source code
pig build dep aggs_for_arrays; # install build dependencies
pig build pkg aggs_for_arrays; # build extension rpm or deb
pig build ext aggs_for_arrays; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install aggs_for_arrays; # install by extension name, for the current active PG version
pig ext install aggs_for_arrays; # install via package alias, for the active PG version
pig ext install aggs_for_arrays -v 18;   # install for PG 18
pig ext install aggs_for_arrays -v 17;   # install for PG 17
pig ext install aggs_for_arrays -v 16;   # install for PG 16
pig ext install aggs_for_arrays -v 15;   # install for PG 15
pig ext install aggs_for_arrays -v 14;   # install for PG 14
pig ext install aggs_for_arrays -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION aggs_for_arrays;
```

