---
title: "floatvec"
linkTitle: "floatvec"
description: "Math for vectors (arrays) of numbers"
weight: 4730
categories: ["FUNC"]
width: full
---

[**floatvec**](https://github.com/pjungwir/floatvec) : Math for vectors (arrays) of numbers


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4730** | {{< badge content="floatvec" link="https://github.com/pjungwir/floatvec" >}} | {{< ext "floatvec" >}} | `1.1.1` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_rational" >}} {{< ext "uint" >}} {{< ext "uint128" >}} {{< ext "numeral" >}} {{< ext "aggs_for_vecs" >}} {{< ext "aggs_for_arrays" >}} {{< ext "arraymath" >}} {{< ext "financial" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `floatvec` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.1` | {{< bg "18" "floatvec_18" "green" >}} {{< bg "17" "floatvec_17" "green" >}} {{< bg "16" "floatvec_16" "green" >}} {{< bg "15" "floatvec_15" "green" >}} {{< bg "14" "floatvec_14" "green" >}} {{< bg "13" "floatvec_13" "green" >}} | `floatvec_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.1` | {{< bg "18" "postgresql-18-floatvec" "green" >}} {{< bg "17" "postgresql-17-floatvec" "green" >}} {{< bg "16" "postgresql-16-floatvec" "green" >}} {{< bg "15" "postgresql-15-floatvec" "green" >}} {{< bg "14" "postgresql-14-floatvec" "green" >}} {{< bg "13" "postgresql-13-floatvec" "green" >}} | `postgresql-$v-floatvec` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "floatvec_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-13-floatvec : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-13-floatvec : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-13-floatvec : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-13-floatvec : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-13-floatvec : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-13-floatvec : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-13-floatvec : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-floatvec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-13-floatvec : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `floatvec_18` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.1 KiB | [floatvec_18-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_18-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `floatvec_18` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.3 KiB | [floatvec_18-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_18-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `floatvec_18` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.0 KiB | [floatvec_18-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_18-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `floatvec_18` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.2 KiB | [floatvec_18-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_18-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `floatvec_18` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.8 KiB | [floatvec_18-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/floatvec_18-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `floatvec_18` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.5 KiB | [floatvec_18-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/floatvec_18-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-floatvec` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.0 KiB | [postgresql-18-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-18-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-floatvec` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 25.3 KiB | [postgresql-18-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-18-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-floatvec` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.0 KiB | [postgresql-18-floatvec_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/floatvec/postgresql-18-floatvec_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-floatvec` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 25.3 KiB | [postgresql-18-floatvec_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/floatvec/postgresql-18-floatvec_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-floatvec` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.2 KiB | [postgresql-18-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-18-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-floatvec` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.7 KiB | [postgresql-18-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-18-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-floatvec` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.2 KiB | [postgresql-18-floatvec_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-18-floatvec_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-floatvec` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.7 KiB | [postgresql-18-floatvec_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-18-floatvec_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `floatvec_17` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.1 KiB | [floatvec_17-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_17-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `floatvec_17` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.3 KiB | [floatvec_17-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_17-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `floatvec_17` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.0 KiB | [floatvec_17-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_17-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `floatvec_17` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.2 KiB | [floatvec_17-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_17-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `floatvec_17` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.8 KiB | [floatvec_17-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/floatvec_17-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `floatvec_17` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.5 KiB | [floatvec_17-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/floatvec_17-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-floatvec` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.0 KiB | [postgresql-17-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-floatvec` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 25.3 KiB | [postgresql-17-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-floatvec` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.0 KiB | [postgresql-17-floatvec_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-floatvec` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 25.2 KiB | [postgresql-17-floatvec_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-floatvec` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.8 KiB | [postgresql-17-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-floatvec` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.3 KiB | [postgresql-17-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-floatvec` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.2 KiB | [postgresql-17-floatvec_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-floatvec` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.6 KiB | [postgresql-17-floatvec_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-17-floatvec_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `floatvec_16` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.1 KiB | [floatvec_16-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_16-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `floatvec_16` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.3 KiB | [floatvec_16-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_16-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `floatvec_16` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.0 KiB | [floatvec_16-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_16-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `floatvec_16` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.2 KiB | [floatvec_16-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_16-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `floatvec_16` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.9 KiB | [floatvec_16-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/floatvec_16-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `floatvec_16` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.5 KiB | [floatvec_16-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/floatvec_16-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-floatvec` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.0 KiB | [postgresql-16-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-floatvec` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 25.3 KiB | [postgresql-16-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-floatvec` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.0 KiB | [postgresql-16-floatvec_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-floatvec` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 25.3 KiB | [postgresql-16-floatvec_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-floatvec` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.8 KiB | [postgresql-16-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-floatvec` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.4 KiB | [postgresql-16-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-floatvec` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.2 KiB | [postgresql-16-floatvec_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-floatvec` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.6 KiB | [postgresql-16-floatvec_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-16-floatvec_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `floatvec_15` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.1 KiB | [floatvec_15-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_15-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `floatvec_15` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.0 KiB | [floatvec_15-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_15-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `floatvec_15` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.0 KiB | [floatvec_15-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_15-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `floatvec_15` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.8 KiB | [floatvec_15-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_15-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `floatvec_15` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.9 KiB | [floatvec_15-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/floatvec_15-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `floatvec_15` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.2 KiB | [floatvec_15-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/floatvec_15-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-floatvec` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.6 KiB | [postgresql-15-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-floatvec` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.4 KiB | [postgresql-15-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-floatvec` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.6 KiB | [postgresql-15-floatvec_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-floatvec` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.5 KiB | [postgresql-15-floatvec_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-floatvec` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.7 KiB | [postgresql-15-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-floatvec` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.8 KiB | [postgresql-15-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-floatvec` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.7 KiB | [postgresql-15-floatvec_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-floatvec` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.8 KiB | [postgresql-15-floatvec_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-15-floatvec_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `floatvec_14` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.1 KiB | [floatvec_14-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_14-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `floatvec_14` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.0 KiB | [floatvec_14-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_14-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `floatvec_14` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.9 KiB | [floatvec_14-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_14-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `floatvec_14` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.8 KiB | [floatvec_14-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_14-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `floatvec_14` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.9 KiB | [floatvec_14-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/floatvec_14-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `floatvec_14` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.2 KiB | [floatvec_14-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/floatvec_14-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-floatvec` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.6 KiB | [postgresql-14-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-floatvec` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.4 KiB | [postgresql-14-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-floatvec` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.6 KiB | [postgresql-14-floatvec_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-floatvec` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.5 KiB | [postgresql-14-floatvec_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-floatvec` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.7 KiB | [postgresql-14-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-floatvec` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.7 KiB | [postgresql-14-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-floatvec` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.7 KiB | [postgresql-14-floatvec_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-floatvec` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.7 KiB | [postgresql-14-floatvec_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-14-floatvec_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `floatvec_13` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.9 KiB | [floatvec_13-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatvec_13-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `floatvec_13` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.9 KiB | [floatvec_13-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatvec_13-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `floatvec_13` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.0 KiB | [floatvec_13-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatvec_13-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `floatvec_13` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.8 KiB | [floatvec_13-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatvec_13-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `floatvec_13` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.8 KiB | [floatvec_13-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/floatvec_13-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `floatvec_13` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.2 KiB | [floatvec_13-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/floatvec_13-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-floatvec` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.1 KiB | [postgresql-13-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-13-floatvec_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-floatvec` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.2 KiB | [postgresql-13-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatvec/postgresql-13-floatvec_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-floatvec` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.2 KiB | [postgresql-13-floatvec_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/floatvec/postgresql-13-floatvec_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-floatvec` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.3 KiB | [postgresql-13-floatvec_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/floatvec/postgresql-13-floatvec_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-floatvec` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.7 KiB | [postgresql-13-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-13-floatvec_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-floatvec` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.4 KiB | [postgresql-13-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatvec/postgresql-13-floatvec_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-floatvec` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.4 KiB | [postgresql-13-floatvec_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-13-floatvec_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-floatvec` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.6 KiB | [postgresql-13-floatvec_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatvec/postgresql-13-floatvec_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pjungwir/floatvec" title="Repository" icon="github" subtitle="github.com/pjungwir/floatvec" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="floatvec-1.1.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg floatvec;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install floatvec;		# install via package name, for the active PG version

pig install floatvec -v 18;   # install for PG 18
pig install floatvec -v 17;   # install for PG 17
pig install floatvec -v 16;   # install for PG 16
pig install floatvec -v 15;   # install for PG 15
pig install floatvec -v 14;   # install for PG 14
pig install floatvec -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION floatvec;
```
