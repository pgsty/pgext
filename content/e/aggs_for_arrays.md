---
title: "aggs_for_arrays"
linkTitle: "aggs_for_arrays"
description: "Various functions for computing statistics on arrays of numbers"
weight: 4750
categories: ["FUNC"]
width: full
---

[**aggs_for_arrays**](https://github.com/pjungwir/aggs_for_arrays) : Various functions for computing statistics on arrays of numbers


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4750** | {{< badge content="aggs_for_arrays" link="https://github.com/pjungwir/aggs_for_arrays" >}} | {{< ext "aggs_for_arrays" >}} | `1.3.3` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "aggs_for_vecs" >}} {{< ext "first_last_agg" >}} {{< ext "arraymath" >}} {{< ext "intarray" >}} {{< ext "topn" >}} {{< ext "quantile" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `aggs_for_arrays` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.3` | {{< bg "18" "aggs_for_arrays_18" "green" >}} {{< bg "17" "aggs_for_arrays_17" "green" >}} {{< bg "16" "aggs_for_arrays_16" "green" >}} {{< bg "15" "aggs_for_arrays_15" "green" >}} {{< bg "14" "aggs_for_arrays_14" "green" >}} | `aggs_for_arrays_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.3` | {{< bg "18" "postgresql-18-aggs-for-arrays" "green" >}} {{< bg "17" "postgresql-17-aggs-for-arrays" "green" >}} {{< bg "16" "postgresql-16-aggs-for-arrays" "green" >}} {{< bg "15" "postgresql-15-aggs-for-arrays" "green" >}} {{< bg "14" "postgresql-14-aggs-for-arrays" "green" >}} | `postgresql-$v-aggs-for-arrays` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "aggs_for_arrays_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-18-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-17-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-16-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-15-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-14-aggs-for-arrays : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-18-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-17-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-16-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-15-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-14-aggs-for-arrays : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-18-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-17-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-16-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-15-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-14-aggs-for-arrays : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-18-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-17-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-16-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-15-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-14-aggs-for-arrays : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-18-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-17-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-16-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-15-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-14-aggs-for-arrays : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-18-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-17-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-16-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-15-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-14-aggs-for-arrays : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-18-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-17-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-16-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-15-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-14-aggs-for-arrays : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-18-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-17-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-16-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-15-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-14-aggs-for-arrays : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-18-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-17-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-16-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-15-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-14-aggs-for-arrays : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-18-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-17-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-16-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-15-aggs-for-arrays : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.3" "postgresql-14-aggs-for-arrays : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aggs_for_arrays_18` | `1.3.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.6 KiB | [aggs_for_arrays_18-1.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_arrays_18-1.3.3-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_arrays_18` | `1.3.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.8 KiB | [aggs_for_arrays_18-1.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_arrays_18-1.3.3-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_arrays_18` | `1.3.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 26.9 KiB | [aggs_for_arrays_18-1.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_arrays_18-1.3.3-1PIGSTY.el9.x86_64.rpm) |
| `aggs_for_arrays_18` | `1.3.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.3 KiB | [aggs_for_arrays_18-1.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_arrays_18-1.3.3-1PIGSTY.el9.aarch64.rpm) |
| `aggs_for_arrays_18` | `1.3.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.8 KiB | [aggs_for_arrays_18-1.3.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aggs_for_arrays_18-1.3.3-1PIGSTY.el10.x86_64.rpm) |
| `aggs_for_arrays_18` | `1.3.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.7 KiB | [aggs_for_arrays_18-1.3.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aggs_for_arrays_18-1.3.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-aggs-for-arrays` | `1.3.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 42.3 KiB | [postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-aggs-for-arrays` | `1.3.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 46.8 KiB | [postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-aggs-for-arrays` | `1.3.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 42.9 KiB | [postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-arrays/postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-aggs-for-arrays` | `1.3.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 47.7 KiB | [postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-arrays/postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-aggs-for-arrays` | `1.3.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 44.3 KiB | [postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-aggs-for-arrays` | `1.3.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 48.4 KiB | [postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-aggs-for-arrays` | `1.3.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 43.6 KiB | [postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-aggs-for-arrays` | `1.3.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 49.7 KiB | [postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-aggs-for-arrays` | `1.3.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 44.3 KiB | [postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/aggs-for-arrays/postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-aggs-for-arrays` | `1.3.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 49.8 KiB | [postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/aggs-for-arrays/postgresql-18-aggs-for-arrays_1.3.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aggs_for_arrays_17` | `1.3.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.6 KiB | [aggs_for_arrays_17-1.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_arrays_17-1.3.3-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_arrays_17` | `1.3.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.8 KiB | [aggs_for_arrays_17-1.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_arrays_17-1.3.3-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_arrays_17` | `1.3.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 26.9 KiB | [aggs_for_arrays_17-1.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_arrays_17-1.3.3-1PIGSTY.el9.x86_64.rpm) |
| `aggs_for_arrays_17` | `1.3.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.2 KiB | [aggs_for_arrays_17-1.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_arrays_17-1.3.3-1PIGSTY.el9.aarch64.rpm) |
| `aggs_for_arrays_17` | `1.3.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.8 KiB | [aggs_for_arrays_17-1.3.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aggs_for_arrays_17-1.3.3-1PIGSTY.el10.x86_64.rpm) |
| `aggs_for_arrays_17` | `1.3.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.7 KiB | [aggs_for_arrays_17-1.3.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aggs_for_arrays_17-1.3.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-aggs-for-arrays` | `1.3.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 42.4 KiB | [postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-aggs-for-arrays` | `1.3.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 46.9 KiB | [postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-aggs-for-arrays` | `1.3.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 42.9 KiB | [postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-aggs-for-arrays` | `1.3.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 47.8 KiB | [postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-aggs-for-arrays` | `1.3.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 45.6 KiB | [postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-aggs-for-arrays` | `1.3.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 49.8 KiB | [postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-aggs-for-arrays` | `1.3.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 43.6 KiB | [postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-aggs-for-arrays` | `1.3.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 49.7 KiB | [postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-aggs-for-arrays` | `1.3.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 44.4 KiB | [postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-aggs-for-arrays` | `1.3.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 49.9 KiB | [postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/aggs-for-arrays/postgresql-17-aggs-for-arrays_1.3.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aggs_for_arrays_16` | `1.3.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.6 KiB | [aggs_for_arrays_16-1.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_arrays_16-1.3.3-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_arrays_16` | `1.3.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.9 KiB | [aggs_for_arrays_16-1.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_arrays_16-1.3.3-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_arrays_16` | `1.3.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 26.8 KiB | [aggs_for_arrays_16-1.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_arrays_16-1.3.3-1PIGSTY.el9.x86_64.rpm) |
| `aggs_for_arrays_16` | `1.3.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.3 KiB | [aggs_for_arrays_16-1.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_arrays_16-1.3.3-1PIGSTY.el9.aarch64.rpm) |
| `aggs_for_arrays_16` | `1.3.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.8 KiB | [aggs_for_arrays_16-1.3.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aggs_for_arrays_16-1.3.3-1PIGSTY.el10.x86_64.rpm) |
| `aggs_for_arrays_16` | `1.3.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.7 KiB | [aggs_for_arrays_16-1.3.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aggs_for_arrays_16-1.3.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-aggs-for-arrays` | `1.3.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 42.4 KiB | [postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-aggs-for-arrays` | `1.3.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 46.9 KiB | [postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-aggs-for-arrays` | `1.3.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 42.9 KiB | [postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-aggs-for-arrays` | `1.3.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 47.9 KiB | [postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-aggs-for-arrays` | `1.3.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 45.6 KiB | [postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-aggs-for-arrays` | `1.3.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 49.8 KiB | [postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-aggs-for-arrays` | `1.3.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 43.6 KiB | [postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-aggs-for-arrays` | `1.3.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 49.7 KiB | [postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-aggs-for-arrays` | `1.3.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 44.4 KiB | [postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-aggs-for-arrays` | `1.3.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 49.9 KiB | [postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/aggs-for-arrays/postgresql-16-aggs-for-arrays_1.3.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aggs_for_arrays_15` | `1.3.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.7 KiB | [aggs_for_arrays_15-1.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_arrays_15-1.3.3-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_arrays_15` | `1.3.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.0 KiB | [aggs_for_arrays_15-1.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_arrays_15-1.3.3-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_arrays_15` | `1.3.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.0 KiB | [aggs_for_arrays_15-1.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_arrays_15-1.3.3-1PIGSTY.el9.x86_64.rpm) |
| `aggs_for_arrays_15` | `1.3.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.4 KiB | [aggs_for_arrays_15-1.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_arrays_15-1.3.3-1PIGSTY.el9.aarch64.rpm) |
| `aggs_for_arrays_15` | `1.3.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.9 KiB | [aggs_for_arrays_15-1.3.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aggs_for_arrays_15-1.3.3-1PIGSTY.el10.x86_64.rpm) |
| `aggs_for_arrays_15` | `1.3.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.9 KiB | [aggs_for_arrays_15-1.3.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aggs_for_arrays_15-1.3.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-aggs-for-arrays` | `1.3.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 42.6 KiB | [postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-aggs-for-arrays` | `1.3.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 47.1 KiB | [postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-aggs-for-arrays` | `1.3.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 43.1 KiB | [postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-aggs-for-arrays` | `1.3.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 48.0 KiB | [postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-aggs-for-arrays` | `1.3.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 45.7 KiB | [postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-aggs-for-arrays` | `1.3.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 50.1 KiB | [postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-aggs-for-arrays` | `1.3.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 43.8 KiB | [postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-aggs-for-arrays` | `1.3.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 50.0 KiB | [postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-aggs-for-arrays` | `1.3.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 44.7 KiB | [postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-aggs-for-arrays` | `1.3.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 50.2 KiB | [postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/aggs-for-arrays/postgresql-15-aggs-for-arrays_1.3.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `aggs_for_arrays_14` | `1.3.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.7 KiB | [aggs_for_arrays_14-1.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/aggs_for_arrays_14-1.3.3-1PIGSTY.el8.x86_64.rpm) |
| `aggs_for_arrays_14` | `1.3.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.0 KiB | [aggs_for_arrays_14-1.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/aggs_for_arrays_14-1.3.3-1PIGSTY.el8.aarch64.rpm) |
| `aggs_for_arrays_14` | `1.3.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.0 KiB | [aggs_for_arrays_14-1.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/aggs_for_arrays_14-1.3.3-1PIGSTY.el9.x86_64.rpm) |
| `aggs_for_arrays_14` | `1.3.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.4 KiB | [aggs_for_arrays_14-1.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/aggs_for_arrays_14-1.3.3-1PIGSTY.el9.aarch64.rpm) |
| `aggs_for_arrays_14` | `1.3.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.9 KiB | [aggs_for_arrays_14-1.3.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/aggs_for_arrays_14-1.3.3-1PIGSTY.el10.x86_64.rpm) |
| `aggs_for_arrays_14` | `1.3.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.9 KiB | [aggs_for_arrays_14-1.3.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/aggs_for_arrays_14-1.3.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-aggs-for-arrays` | `1.3.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 42.6 KiB | [postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-aggs-for-arrays` | `1.3.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 47.2 KiB | [postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-aggs-for-arrays` | `1.3.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 43.1 KiB | [postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-aggs-for-arrays` | `1.3.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 48.0 KiB | [postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-aggs-for-arrays` | `1.3.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 45.6 KiB | [postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-aggs-for-arrays` | `1.3.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 50.0 KiB | [postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-aggs-for-arrays` | `1.3.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 43.8 KiB | [postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-aggs-for-arrays` | `1.3.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 50.1 KiB | [postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-aggs-for-arrays` | `1.3.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 44.7 KiB | [postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-aggs-for-arrays` | `1.3.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 50.2 KiB | [postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/aggs-for-arrays/postgresql-14-aggs-for-arrays_1.3.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pjungwir/aggs_for_arrays" title="Repository" icon="github" subtitle="github.com/pjungwir/aggs_for_arrays" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="aggs_for_arrays-1.3.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg aggs_for_arrays;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install aggs_for_arrays;		# install via package name, for the active PG version

pig install aggs_for_arrays -v 18;   # install for PG 18
pig install aggs_for_arrays -v 17;   # install for PG 17
pig install aggs_for_arrays -v 16;   # install for PG 16
pig install aggs_for_arrays -v 15;   # install for PG 15
pig install aggs_for_arrays -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION aggs_for_arrays;
```



## Usage

> [aggs_for_arrays: aggregate-like functions for single arrays (column-based)](https://github.com/pjungwir/aggs_for_arrays)

Provides functions that compute statistics on a single array input. Supports `SMALLINT`, `INTEGER`, `BIGINT`, `REAL`, and `DOUBLE PRECISION`.

```sql
CREATE EXTENSION aggs_for_arrays;
```

### Functions

| Function | Description |
|---|---|
| `array_to_hist(values T[], start T, width T, count INT)` | Compute histogram bucket counts |
| `array_to_hist_2d(x[], y[], ...)` | 2-D histogram |
| `array_to_mean(values T[])` | Mean of array elements |
| `array_to_median(values T[])` | Median (unsorted input OK) |
| `sorted_array_to_median(values T[])` | Median (pre-sorted input) |
| `array_to_mode(values T[])` | Mode of array elements |
| `sorted_array_to_mode(values T[])` | Mode (pre-sorted input) |
| `array_to_percentile(values T[], pct FLOAT)` | Percentile (0 to 1) |
| `sorted_array_to_percentile(values T[], pct FLOAT)` | Percentile (pre-sorted input) |
| `array_to_percentiles(values T[], pcts FLOAT[])` | Multiple percentiles |
| `sorted_array_to_percentiles(values T[], pcts FLOAT[])` | Multiple percentiles (pre-sorted) |
| `array_to_max(values T[])` | Maximum element |
| `array_to_min(values T[])` | Minimum element |
| `array_to_min_max(values T[])` | `{min, max}` tuple |
| `array_to_skewness(values T[])` | Skewness |
| `array_to_kurtosis(values T[])` | Kurtosis |

### Examples

```sql
-- Histogram with 10 buckets starting at 0, width 10
SELECT array_to_hist(ARRAY[5,15,25,35,45], 0, 10, 5);
-- {1,1,1,1,1}

-- Mean and median
SELECT array_to_mean(ARRAY[1,2,3,4,5]);   -- 3.0
SELECT array_to_median(ARRAY[1,2,3,4,5]); -- 3.0

-- Percentiles
SELECT array_to_percentile(ARRAY[1,2,3,4,5,6,7,8,9,10], 0.95);

-- Multiple percentiles at once
SELECT array_to_percentiles(ARRAY[1,2,3,4,5], ARRAY[0.25, 0.5, 0.75]);
```
