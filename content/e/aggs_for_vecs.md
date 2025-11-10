---
title: "aggs_for_vecs"
linkTitle: "aggs_for_vecs"
description: "Aggregate functions for array inputs"
weight: 4740
categories: ["FUNC"]
width: full
---

[**aggs_for_vecs**](https://github.com/pjungwir/aggs_for_vecs)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4740** | {{< badge content="aggs_for_vecs" link="https://github.com/pjungwir/aggs_for_vecs" >}} | {{< ext "aggs_for_vecs" >}} | `1.4.0` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "aggs_for_arrays" >}} {{< ext "first_last_agg" >}} {{< ext "arraymath" >}} {{< ext "floatvec" >}} {{< ext "vector" >}} {{< ext "topn" >}} |

> [!Note] fix patch not merging


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/aggs_for_vecs" >}} | `1.4.0` | {{< bg "18" "aggs_for_vecs_18*" "green" >}} {{< bg "17" "aggs_for_vecs_17*" "green" >}} {{< bg "16" "aggs_for_vecs_16*" "green" >}} {{< bg "15" "aggs_for_vecs_15*" "green" >}} {{< bg "14" "aggs_for_vecs_14*" "green" >}} {{< bg "13" "aggs_for_vecs_13*" "green" >}} | `aggs_for_vecs_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/aggs_for_vecs" >}} | `1.4.0` | {{< bg "18" "postgresql-18-aggs-for-vecs" "green" >}} {{< bg "17" "postgresql-17-aggs-for-vecs" "green" >}} {{< bg "16" "postgresql-16-aggs-for-vecs" "green" >}} {{< bg "15" "postgresql-15-aggs-for-vecs" "green" >}} {{< bg "14" "postgresql-14-aggs-for-vecs" "green" >}} {{< bg "13" "postgresql-13-aggs-for-vecs" "green" >}} | `postgresql-$v-aggs-for-vecs` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "aggs_for_vecs_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-18-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-17-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-13-aggs-for-vecs : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-18-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-17-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-13-aggs-for-vecs : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-18-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-17-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-13-aggs-for-vecs : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-18-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-17-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-13-aggs-for-vecs : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-18-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-17-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-13-aggs-for-vecs : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-18-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-17-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-13-aggs-for-vecs : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-18-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-17-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-13-aggs-for-vecs : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-18-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-17-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-13-aggs-for-vecs : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aggs_for_vecs_18` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.6 KiB | [aggs_for_vecs_18-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_vecs_18-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_vecs_18` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.6 KiB | [aggs_for_vecs_18-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_vecs_18-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_vecs_18` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.2 KiB | [aggs_for_vecs_18-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_vecs_18-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `aggs_for_vecs_18` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.2 KiB | [aggs_for_vecs_18-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_vecs_18-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `aggs_for_vecs_18` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.1 KiB | [aggs_for_vecs_18-1.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aggs_for_vecs_18-1.4.0-1PIGSTY.el10.x86_64.rpm) |
| `aggs_for_vecs_18` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.4 KiB | [aggs_for_vecs_18-1.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aggs_for_vecs_18-1.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-aggs-for-vecs` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 82.1 KiB | [postgresql-18-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-18-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-aggs-for-vecs` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 82.2 KiB | [postgresql-18-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-18-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-aggs-for-vecs` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 81.9 KiB | [postgresql-18-aggs-for-vecs_1.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-18-aggs-for-vecs_1.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-aggs-for-vecs` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.1 KiB | [postgresql-18-aggs-for-vecs_1.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-18-aggs-for-vecs_1.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-aggs-for-vecs` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 88.6 KiB | [postgresql-18-aggs-for-vecs_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-18-aggs-for-vecs_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-aggs-for-vecs` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 89.1 KiB | [postgresql-18-aggs-for-vecs_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-18-aggs-for-vecs_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-aggs-for-vecs` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 84.7 KiB | [postgresql-18-aggs-for-vecs_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-18-aggs-for-vecs_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-aggs-for-vecs` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 86.3 KiB | [postgresql-18-aggs-for-vecs_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-18-aggs-for-vecs_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aggs_for_vecs_17` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.6 KiB | [aggs_for_vecs_17-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_vecs_17-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_vecs_17` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.6 KiB | [aggs_for_vecs_17-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_vecs_17-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_vecs_17` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.3 KiB | [aggs_for_vecs_17-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_vecs_17-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `aggs_for_vecs_17` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.3 KiB | [aggs_for_vecs_17-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_vecs_17-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `aggs_for_vecs_17` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.0 KiB | [aggs_for_vecs_17-1.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aggs_for_vecs_17-1.4.0-1PIGSTY.el10.x86_64.rpm) |
| `aggs_for_vecs_17` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.5 KiB | [aggs_for_vecs_17-1.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aggs_for_vecs_17-1.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-aggs-for-vecs` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 82.1 KiB | [postgresql-17-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-17-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-aggs-for-vecs` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 82.4 KiB | [postgresql-17-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-17-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-aggs-for-vecs` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 81.7 KiB | [postgresql-17-aggs-for-vecs_1.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-17-aggs-for-vecs_1.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-aggs-for-vecs` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.2 KiB | [postgresql-17-aggs-for-vecs_1.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-17-aggs-for-vecs_1.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-aggs-for-vecs` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 93.0 KiB | [postgresql-17-aggs-for-vecs_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-17-aggs-for-vecs_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-aggs-for-vecs` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 93.7 KiB | [postgresql-17-aggs-for-vecs_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-17-aggs-for-vecs_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-aggs-for-vecs` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 84.7 KiB | [postgresql-17-aggs-for-vecs_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-17-aggs-for-vecs_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-aggs-for-vecs` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 86.4 KiB | [postgresql-17-aggs-for-vecs_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-17-aggs-for-vecs_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aggs_for_vecs_16` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.6 KiB | [aggs_for_vecs_16-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_vecs_16-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_vecs_16` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.6 KiB | [aggs_for_vecs_16-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_vecs_16-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_vecs_16` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.3 KiB | [aggs_for_vecs_16-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_vecs_16-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `aggs_for_vecs_16` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.3 KiB | [aggs_for_vecs_16-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_vecs_16-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `aggs_for_vecs_16` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.0 KiB | [aggs_for_vecs_16-1.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aggs_for_vecs_16-1.4.0-1PIGSTY.el10.x86_64.rpm) |
| `aggs_for_vecs_16` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.5 KiB | [aggs_for_vecs_16-1.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aggs_for_vecs_16-1.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-aggs-for-vecs` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 82.0 KiB | [postgresql-16-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-16-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-aggs-for-vecs` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 82.2 KiB | [postgresql-16-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-16-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-aggs-for-vecs` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 81.8 KiB | [postgresql-16-aggs-for-vecs_1.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-16-aggs-for-vecs_1.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-aggs-for-vecs` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.2 KiB | [postgresql-16-aggs-for-vecs_1.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-16-aggs-for-vecs_1.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-aggs-for-vecs` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 93.0 KiB | [postgresql-16-aggs-for-vecs_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-16-aggs-for-vecs_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-aggs-for-vecs` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 93.6 KiB | [postgresql-16-aggs-for-vecs_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-16-aggs-for-vecs_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-aggs-for-vecs` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 84.9 KiB | [postgresql-16-aggs-for-vecs_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-16-aggs-for-vecs_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-aggs-for-vecs` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 86.4 KiB | [postgresql-16-aggs-for-vecs_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-16-aggs-for-vecs_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aggs_for_vecs_15` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.8 KiB | [aggs_for_vecs_15-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_vecs_15-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_vecs_15` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.5 KiB | [aggs_for_vecs_15-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_vecs_15-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_vecs_15` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.7 KiB | [aggs_for_vecs_15-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_vecs_15-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `aggs_for_vecs_15` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.3 KiB | [aggs_for_vecs_15-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_vecs_15-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `aggs_for_vecs_15` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.3 KiB | [aggs_for_vecs_15-1.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aggs_for_vecs_15-1.4.0-1PIGSTY.el10.x86_64.rpm) |
| `aggs_for_vecs_15` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.5 KiB | [aggs_for_vecs_15-1.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aggs_for_vecs_15-1.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-aggs-for-vecs` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 82.5 KiB | [postgresql-15-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-15-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-aggs-for-vecs` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 82.2 KiB | [postgresql-15-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-15-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-aggs-for-vecs` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 82.0 KiB | [postgresql-15-aggs-for-vecs_1.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-15-aggs-for-vecs_1.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-aggs-for-vecs` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.3 KiB | [postgresql-15-aggs-for-vecs_1.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-15-aggs-for-vecs_1.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-aggs-for-vecs` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 93.0 KiB | [postgresql-15-aggs-for-vecs_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-15-aggs-for-vecs_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-aggs-for-vecs` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 93.5 KiB | [postgresql-15-aggs-for-vecs_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-15-aggs-for-vecs_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-aggs-for-vecs` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 85.2 KiB | [postgresql-15-aggs-for-vecs_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-15-aggs-for-vecs_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-aggs-for-vecs` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 86.4 KiB | [postgresql-15-aggs-for-vecs_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-15-aggs-for-vecs_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aggs_for_vecs_14` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.8 KiB | [aggs_for_vecs_14-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_vecs_14-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_vecs_14` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.5 KiB | [aggs_for_vecs_14-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_vecs_14-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_vecs_14` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.6 KiB | [aggs_for_vecs_14-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_vecs_14-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `aggs_for_vecs_14` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.2 KiB | [aggs_for_vecs_14-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_vecs_14-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `aggs_for_vecs_14` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.3 KiB | [aggs_for_vecs_14-1.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aggs_for_vecs_14-1.4.0-1PIGSTY.el10.x86_64.rpm) |
| `aggs_for_vecs_14` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.5 KiB | [aggs_for_vecs_14-1.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aggs_for_vecs_14-1.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-aggs-for-vecs` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 82.3 KiB | [postgresql-14-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-14-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-aggs-for-vecs` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 82.2 KiB | [postgresql-14-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-14-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-aggs-for-vecs` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 81.9 KiB | [postgresql-14-aggs-for-vecs_1.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-14-aggs-for-vecs_1.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-aggs-for-vecs` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.0 KiB | [postgresql-14-aggs-for-vecs_1.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-14-aggs-for-vecs_1.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-aggs-for-vecs` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 93.0 KiB | [postgresql-14-aggs-for-vecs_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-14-aggs-for-vecs_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-aggs-for-vecs` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 93.4 KiB | [postgresql-14-aggs-for-vecs_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-14-aggs-for-vecs_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-aggs-for-vecs` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 85.1 KiB | [postgresql-14-aggs-for-vecs_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-14-aggs-for-vecs_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-aggs-for-vecs` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 86.3 KiB | [postgresql-14-aggs-for-vecs_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-14-aggs-for-vecs_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aggs_for_vecs_13` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.2 KiB | [aggs_for_vecs_13-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_vecs_13-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_vecs_13` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.5 KiB | [aggs_for_vecs_13-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_vecs_13-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_vecs_13` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.6 KiB | [aggs_for_vecs_13-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_vecs_13-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `aggs_for_vecs_13` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.3 KiB | [aggs_for_vecs_13-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_vecs_13-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `aggs_for_vecs_13` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.3 KiB | [aggs_for_vecs_13-1.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aggs_for_vecs_13-1.4.0-1PIGSTY.el10.x86_64.rpm) |
| `aggs_for_vecs_13` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.7 KiB | [aggs_for_vecs_13-1.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aggs_for_vecs_13-1.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-aggs-for-vecs` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 82.0 KiB | [postgresql-13-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-13-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-aggs-for-vecs` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 82.0 KiB | [postgresql-13-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-13-aggs-for-vecs_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-aggs-for-vecs` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 81.7 KiB | [postgresql-13-aggs-for-vecs_1.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-13-aggs-for-vecs_1.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-aggs-for-vecs` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 81.9 KiB | [postgresql-13-aggs-for-vecs_1.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-13-aggs-for-vecs_1.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-aggs-for-vecs` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 92.9 KiB | [postgresql-13-aggs-for-vecs_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-13-aggs-for-vecs_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-aggs-for-vecs` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 93.3 KiB | [postgresql-13-aggs-for-vecs_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-13-aggs-for-vecs_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-aggs-for-vecs` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 84.6 KiB | [postgresql-13-aggs-for-vecs_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-13-aggs-for-vecs_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-aggs-for-vecs` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 86.0 KiB | [postgresql-13-aggs-for-vecs_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-13-aggs-for-vecs_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pjungwir/aggs_for_vecs" title="Repository" icon="github" subtitle="github.com/pjungwir/aggs_for_vecs" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="aggs_for_vecs-1.4.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get aggs_for_vecs; # get aggs_for_vecs source code
pig build dep aggs_for_vecs; # install build dependencies
pig build pkg aggs_for_vecs; # build extension rpm or deb
pig build ext aggs_for_vecs; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install aggs_for_vecs; # install by extension name, for the current active PG version
pig ext install aggs_for_vecs; # install via package alias, for the active PG version
pig ext install aggs_for_vecs -v 18;   # install for PG 18
pig ext install aggs_for_vecs -v 17;   # install for PG 17
pig ext install aggs_for_vecs -v 16;   # install for PG 16
pig ext install aggs_for_vecs -v 15;   # install for PG 15
pig ext install aggs_for_vecs -v 14;   # install for PG 14
pig ext install aggs_for_vecs -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION aggs_for_vecs;
```

