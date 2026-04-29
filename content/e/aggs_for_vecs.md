---
title: "aggs_for_vecs"
linkTitle: "aggs_for_vecs"
description: "Aggregate functions for array inputs"
weight: 4740
categories: ["FUNC"]
width: full
---

[**aggs_for_vecs**](https://github.com/pjungwir/aggs_for_vecs) : Aggregate functions for array inputs


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4740** | {{< badge content="aggs_for_vecs" link="https://github.com/pjungwir/aggs_for_vecs" >}} | {{< ext "aggs_for_vecs" >}} | `1.4.1` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "aggs_for_arrays" >}} {{< ext "first_last_agg" >}} {{< ext "arraymath" >}} {{< ext "floatvec" >}} {{< ext "vector" >}} {{< ext "topn" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `aggs_for_vecs` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4.1` | {{< bg "18" "aggs_for_vecs_18" "green" >}} {{< bg "17" "aggs_for_vecs_17" "green" >}} {{< bg "16" "aggs_for_vecs_16" "green" >}} {{< bg "15" "aggs_for_vecs_15" "green" >}} {{< bg "14" "aggs_for_vecs_14" "green" >}} | `aggs_for_vecs_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4.1` | {{< bg "18" "postgresql-18-aggs-for-vecs" "green" >}} {{< bg "17" "postgresql-17-aggs-for-vecs" "green" >}} {{< bg "16" "postgresql-16-aggs-for-vecs" "green" >}} {{< bg "15" "postgresql-15-aggs-for-vecs" "green" >}} {{< bg "14" "postgresql-14-aggs-for-vecs" "green" >}} | `postgresql-$v-aggs-for-vecs` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "aggs_for_vecs_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-aggs-for-vecs : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-aggs-for-vecs : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-aggs-for-vecs : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-aggs-for-vecs : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-aggs-for-vecs : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-aggs-for-vecs : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-aggs-for-vecs : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-aggs-for-vecs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-aggs-for-vecs : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-aggs-for-vecs : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-aggs-for-vecs : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-aggs-for-vecs : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-aggs-for-vecs : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-aggs-for-vecs : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-aggs-for-vecs : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-aggs-for-vecs : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-aggs-for-vecs : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-aggs-for-vecs : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-aggs-for-vecs : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aggs_for_vecs_18` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.9 KiB | [aggs_for_vecs_18-1.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_vecs_18-1.4.1-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_vecs_18` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.9 KiB | [aggs_for_vecs_18-1.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_vecs_18-1.4.1-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_vecs_18` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.4 KiB | [aggs_for_vecs_18-1.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_vecs_18-1.4.1-1PIGSTY.el9.x86_64.rpm) |
| `aggs_for_vecs_18` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.4 KiB | [aggs_for_vecs_18-1.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_vecs_18-1.4.1-1PIGSTY.el9.aarch64.rpm) |
| `aggs_for_vecs_18` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.4 KiB | [aggs_for_vecs_18-1.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aggs_for_vecs_18-1.4.1-1PIGSTY.el10.x86_64.rpm) |
| `aggs_for_vecs_18` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.7 KiB | [aggs_for_vecs_18-1.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aggs_for_vecs_18-1.4.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-aggs-for-vecs` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 82.1 KiB | [postgresql-18-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-18-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-aggs-for-vecs` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 82.2 KiB | [postgresql-18-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-18-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-aggs-for-vecs` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 81.9 KiB | [postgresql-18-aggs-for-vecs_1.4.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-18-aggs-for-vecs_1.4.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-aggs-for-vecs` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.2 KiB | [postgresql-18-aggs-for-vecs_1.4.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-18-aggs-for-vecs_1.4.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-aggs-for-vecs` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 88.6 KiB | [postgresql-18-aggs-for-vecs_1.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-18-aggs-for-vecs_1.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-aggs-for-vecs` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 89.2 KiB | [postgresql-18-aggs-for-vecs_1.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-18-aggs-for-vecs_1.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-aggs-for-vecs` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 84.7 KiB | [postgresql-18-aggs-for-vecs_1.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-18-aggs-for-vecs_1.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-aggs-for-vecs` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 86.4 KiB | [postgresql-18-aggs-for-vecs_1.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-18-aggs-for-vecs_1.4.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aggs_for_vecs_17` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.9 KiB | [aggs_for_vecs_17-1.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_vecs_17-1.4.1-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_vecs_17` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.9 KiB | [aggs_for_vecs_17-1.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_vecs_17-1.4.1-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_vecs_17` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.6 KiB | [aggs_for_vecs_17-1.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_vecs_17-1.4.1-1PIGSTY.el9.x86_64.rpm) |
| `aggs_for_vecs_17` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.5 KiB | [aggs_for_vecs_17-1.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_vecs_17-1.4.1-1PIGSTY.el9.aarch64.rpm) |
| `aggs_for_vecs_17` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.3 KiB | [aggs_for_vecs_17-1.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aggs_for_vecs_17-1.4.1-1PIGSTY.el10.x86_64.rpm) |
| `aggs_for_vecs_17` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.8 KiB | [aggs_for_vecs_17-1.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aggs_for_vecs_17-1.4.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-aggs-for-vecs` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 82.3 KiB | [postgresql-17-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-17-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-aggs-for-vecs` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 82.2 KiB | [postgresql-17-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-17-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-aggs-for-vecs` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 81.7 KiB | [postgresql-17-aggs-for-vecs_1.4.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-17-aggs-for-vecs_1.4.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-aggs-for-vecs` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.3 KiB | [postgresql-17-aggs-for-vecs_1.4.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-17-aggs-for-vecs_1.4.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-aggs-for-vecs` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 93.0 KiB | [postgresql-17-aggs-for-vecs_1.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-17-aggs-for-vecs_1.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-aggs-for-vecs` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 93.7 KiB | [postgresql-17-aggs-for-vecs_1.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-17-aggs-for-vecs_1.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-aggs-for-vecs` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 84.7 KiB | [postgresql-17-aggs-for-vecs_1.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-17-aggs-for-vecs_1.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-aggs-for-vecs` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 86.4 KiB | [postgresql-17-aggs-for-vecs_1.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-17-aggs-for-vecs_1.4.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aggs_for_vecs_16` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.8 KiB | [aggs_for_vecs_16-1.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_vecs_16-1.4.1-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_vecs_16` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.9 KiB | [aggs_for_vecs_16-1.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_vecs_16-1.4.1-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_vecs_16` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.6 KiB | [aggs_for_vecs_16-1.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_vecs_16-1.4.1-1PIGSTY.el9.x86_64.rpm) |
| `aggs_for_vecs_16` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.5 KiB | [aggs_for_vecs_16-1.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_vecs_16-1.4.1-1PIGSTY.el9.aarch64.rpm) |
| `aggs_for_vecs_16` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.3 KiB | [aggs_for_vecs_16-1.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aggs_for_vecs_16-1.4.1-1PIGSTY.el10.x86_64.rpm) |
| `aggs_for_vecs_16` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.8 KiB | [aggs_for_vecs_16-1.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aggs_for_vecs_16-1.4.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-aggs-for-vecs` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 82.3 KiB | [postgresql-16-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-16-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-aggs-for-vecs` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 82.2 KiB | [postgresql-16-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-16-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-aggs-for-vecs` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 81.7 KiB | [postgresql-16-aggs-for-vecs_1.4.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-16-aggs-for-vecs_1.4.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-aggs-for-vecs` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.2 KiB | [postgresql-16-aggs-for-vecs_1.4.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-16-aggs-for-vecs_1.4.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-aggs-for-vecs` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 93.0 KiB | [postgresql-16-aggs-for-vecs_1.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-16-aggs-for-vecs_1.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-aggs-for-vecs` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 93.7 KiB | [postgresql-16-aggs-for-vecs_1.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-16-aggs-for-vecs_1.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-aggs-for-vecs` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 85.0 KiB | [postgresql-16-aggs-for-vecs_1.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-16-aggs-for-vecs_1.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-aggs-for-vecs` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 86.4 KiB | [postgresql-16-aggs-for-vecs_1.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-16-aggs-for-vecs_1.4.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aggs_for_vecs_15` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 44.0 KiB | [aggs_for_vecs_15-1.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_vecs_15-1.4.1-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_vecs_15` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.7 KiB | [aggs_for_vecs_15-1.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_vecs_15-1.4.1-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_vecs_15` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.9 KiB | [aggs_for_vecs_15-1.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_vecs_15-1.4.1-1PIGSTY.el9.x86_64.rpm) |
| `aggs_for_vecs_15` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.5 KiB | [aggs_for_vecs_15-1.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_vecs_15-1.4.1-1PIGSTY.el9.aarch64.rpm) |
| `aggs_for_vecs_15` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.6 KiB | [aggs_for_vecs_15-1.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aggs_for_vecs_15-1.4.1-1PIGSTY.el10.x86_64.rpm) |
| `aggs_for_vecs_15` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.8 KiB | [aggs_for_vecs_15-1.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aggs_for_vecs_15-1.4.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-aggs-for-vecs` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 82.5 KiB | [postgresql-15-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-15-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-aggs-for-vecs` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 82.2 KiB | [postgresql-15-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-15-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-aggs-for-vecs` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 82.1 KiB | [postgresql-15-aggs-for-vecs_1.4.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-15-aggs-for-vecs_1.4.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-aggs-for-vecs` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.3 KiB | [postgresql-15-aggs-for-vecs_1.4.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-15-aggs-for-vecs_1.4.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-aggs-for-vecs` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 93.2 KiB | [postgresql-15-aggs-for-vecs_1.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-15-aggs-for-vecs_1.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-aggs-for-vecs` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 93.5 KiB | [postgresql-15-aggs-for-vecs_1.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-15-aggs-for-vecs_1.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-aggs-for-vecs` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 85.3 KiB | [postgresql-15-aggs-for-vecs_1.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-15-aggs-for-vecs_1.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-aggs-for-vecs` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 86.5 KiB | [postgresql-15-aggs-for-vecs_1.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-15-aggs-for-vecs_1.4.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aggs_for_vecs_14` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 44.0 KiB | [aggs_for_vecs_14-1.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_vecs_14-1.4.1-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_vecs_14` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.7 KiB | [aggs_for_vecs_14-1.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_vecs_14-1.4.1-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_vecs_14` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.8 KiB | [aggs_for_vecs_14-1.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_vecs_14-1.4.1-1PIGSTY.el9.x86_64.rpm) |
| `aggs_for_vecs_14` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.5 KiB | [aggs_for_vecs_14-1.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_vecs_14-1.4.1-1PIGSTY.el9.aarch64.rpm) |
| `aggs_for_vecs_14` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.6 KiB | [aggs_for_vecs_14-1.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aggs_for_vecs_14-1.4.1-1PIGSTY.el10.x86_64.rpm) |
| `aggs_for_vecs_14` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.8 KiB | [aggs_for_vecs_14-1.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aggs_for_vecs_14-1.4.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-aggs-for-vecs` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 82.3 KiB | [postgresql-14-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-14-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-aggs-for-vecs` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 82.5 KiB | [postgresql-14-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-vecs/postgresql-14-aggs-for-vecs_1.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-aggs-for-vecs` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 81.9 KiB | [postgresql-14-aggs-for-vecs_1.4.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-14-aggs-for-vecs_1.4.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-aggs-for-vecs` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.1 KiB | [postgresql-14-aggs-for-vecs_1.4.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-vecs/postgresql-14-aggs-for-vecs_1.4.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-aggs-for-vecs` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 93.0 KiB | [postgresql-14-aggs-for-vecs_1.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-14-aggs-for-vecs_1.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-aggs-for-vecs` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 93.5 KiB | [postgresql-14-aggs-for-vecs_1.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-vecs/postgresql-14-aggs-for-vecs_1.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-aggs-for-vecs` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 85.2 KiB | [postgresql-14-aggs-for-vecs_1.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-14-aggs-for-vecs_1.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-aggs-for-vecs` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 86.3 KiB | [postgresql-14-aggs-for-vecs_1.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-vecs/postgresql-14-aggs-for-vecs_1.4.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pjungwir/aggs_for_vecs" title="Repository" icon="github" subtitle="github.com/pjungwir/aggs_for_vecs" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="aggs_for_vecs-1.4.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg aggs_for_vecs;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install aggs_for_vecs;		# install via package name, for the active PG version

pig install aggs_for_vecs -v 18;   # install for PG 18
pig install aggs_for_vecs -v 17;   # install for PG 17
pig install aggs_for_vecs -v 16;   # install for PG 16
pig install aggs_for_vecs -v 15;   # install for PG 15
pig install aggs_for_vecs -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION aggs_for_vecs;
```



## Usage

> [aggs_for_vecs: aggregate functions for arrays (vector/row-based)](https://github.com/pjungwir/aggs_for_vecs)

Provides aggregate functions that operate on arrays as vectors, computing position-wise statistics across multiple rows. Supports `SMALLINT`, `INTEGER`, `BIGINT`, `REAL`, and `DOUBLE PRECISION`.

```sql
CREATE EXTENSION aggs_for_vecs;
```

### Aggregate Functions

| Function | Description |
|---|---|
| `vec_to_count(anyarray)` | Count of non-nulls in each position |
| `vec_to_sum(anyarray)` | Sum in each position |
| `vec_to_min(anyarray)` | Minimum in each position |
| `vec_to_max(anyarray)` | Maximum in each position |
| `vec_to_mean(anyarray)` | Average in each position (returns `FLOAT[]`) |
| `vec_to_weighted_mean(values, weights)` | Weighted average in each position |
| `vec_to_var_samp(anyarray)` | Sample variance in each position |
| `vec_to_first(anyarray)` | First non-null in each position (use with ORDER BY) |
| `vec_to_last(anyarray)` | Last non-null in each position (use with ORDER BY) |
| `hist_2d(x, y, ...)` | 2-D histogram |
| `hist_md(vals, indexes, ...)` | N-dimensional histogram |

### Non-Aggregate Functions

| Function | Description |
|---|---|
| `vec_add(l, r)` | Element-wise addition |
| `vec_sub(l, r)` | Element-wise subtraction |
| `vec_mul(l, r)` | Element-wise multiplication |
| `vec_div(l, r)` | Element-wise division |
| `vec_elements(array, indexes)` | Select elements at given indexes |
| `pad_vec(array, length)` | Extend array to given length with NULLs |
| `vec_coalesce(array, default)` | Replace NULLs with default |
| `vec_trim_scale(numeric[])` | Trim trailing zeros from NUMERIC elements |
| `vec_without_outliers(vals, mins, maxs)` | Replace outliers with NULL |

### Examples

```sql
-- Position-wise minimum across rows
SELECT vec_to_min(vals) FROM (VALUES
  (ARRAY[1,2,3,4]),
  (ARRAY[5,0,-5,0]),
  (ARRAY[3,6,0,9])
) AS t(vals);
-- {1,0,-5,0}

-- Position-wise average
SELECT vec_to_mean(vals) FROM my_table;
```
