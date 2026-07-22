---
title: "argm"
linkTitle: "argm"
description: "argmax, argmin, and anyold aggregate functions"
weight: 4755
categories: ["FUNC"]
width: full
---

[**argm**](https://github.com/bashtanov/argm) : argmax, argmin, and anyold aggregate functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4755** | {{< badge content="argm" link="https://github.com/bashtanov/argm" >}} | {{< ext "argm" >}} | `1.1.1` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "first_last_agg" >}} {{< ext "aggs_for_arrays" >}} {{< ext "aggs_for_vecs" >}} {{< ext "topn" >}} |

> [!Note] fix pg16+ varlena header with patch


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `argm` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.1` | {{< bg "18" "argm_18" "green" >}} {{< bg "17" "argm_17" "green" >}} {{< bg "16" "argm_16" "green" >}} {{< bg "15" "argm_15" "green" >}} {{< bg "14" "argm_14" "green" >}} | `argm_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.1` | {{< bg "18" "postgresql-18-argm" "green" >}} {{< bg "17" "postgresql-17-argm" "green" >}} {{< bg "16" "postgresql-16-argm" "green" >}} {{< bg "15" "postgresql-15-argm" "green" >}} {{< bg "14" "postgresql-14-argm" "green" >}} | `postgresql-$v-argm` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "argm_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "argm_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "argm_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "argm_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "argm_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "argm_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "argm_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-argm : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-argm : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-argm : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-argm : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-argm : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-argm : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-argm : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-argm : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-argm : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-argm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-argm : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `argm_18` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.3 KiB | [argm_18-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/argm_18-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `argm_18` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.4 KiB | [argm_18-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/argm_18-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `argm_18` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.2 KiB | [argm_18-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/argm_18-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `argm_18` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.3 KiB | [argm_18-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/argm_18-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `argm_18` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.2 KiB | [argm_18-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/argm_18-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `argm_18` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.5 KiB | [argm_18-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/argm_18-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-argm` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.1 KiB | [postgresql-18-argm_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/argm/postgresql-18-argm_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-argm` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.3 KiB | [postgresql-18-argm_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/argm/postgresql-18-argm_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-argm` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.2 KiB | [postgresql-18-argm_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/argm/postgresql-18-argm_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-argm` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.3 KiB | [postgresql-18-argm_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/argm/postgresql-18-argm_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-argm` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.0 KiB | [postgresql-18-argm_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/argm/postgresql-18-argm_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-argm` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.2 KiB | [postgresql-18-argm_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/argm/postgresql-18-argm_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-argm` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.1 KiB | [postgresql-18-argm_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/argm/postgresql-18-argm_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-argm` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.3 KiB | [postgresql-18-argm_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/argm/postgresql-18-argm_1.1.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-argm` | `1.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 19.9 KiB | [postgresql-18-argm_1.1.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/argm/postgresql-18-argm_1.1.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-argm` | `1.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 20.0 KiB | [postgresql-18-argm_1.1.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/argm/postgresql-18-argm_1.1.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `argm_17` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.3 KiB | [argm_17-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/argm_17-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `argm_17` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.4 KiB | [argm_17-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/argm_17-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `argm_17` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.2 KiB | [argm_17-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/argm_17-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `argm_17` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.3 KiB | [argm_17-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/argm_17-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `argm_17` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.2 KiB | [argm_17-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/argm_17-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `argm_17` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.5 KiB | [argm_17-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/argm_17-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-argm` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.1 KiB | [postgresql-17-argm_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/argm/postgresql-17-argm_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-argm` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.3 KiB | [postgresql-17-argm_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/argm/postgresql-17-argm_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-argm` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.1 KiB | [postgresql-17-argm_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/argm/postgresql-17-argm_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-argm` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.4 KiB | [postgresql-17-argm_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/argm/postgresql-17-argm_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-argm` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.9 KiB | [postgresql-17-argm_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/argm/postgresql-17-argm_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-argm` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.1 KiB | [postgresql-17-argm_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/argm/postgresql-17-argm_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-argm` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.1 KiB | [postgresql-17-argm_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/argm/postgresql-17-argm_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-argm` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.2 KiB | [postgresql-17-argm_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/argm/postgresql-17-argm_1.1.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-argm` | `1.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.0 KiB | [postgresql-17-argm_1.1.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/argm/postgresql-17-argm_1.1.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-argm` | `1.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 20.1 KiB | [postgresql-17-argm_1.1.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/argm/postgresql-17-argm_1.1.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `argm_16` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.3 KiB | [argm_16-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/argm_16-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `argm_16` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.4 KiB | [argm_16-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/argm_16-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `argm_16` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.2 KiB | [argm_16-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/argm_16-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `argm_16` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.3 KiB | [argm_16-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/argm_16-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `argm_16` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.2 KiB | [argm_16-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/argm_16-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `argm_16` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.4 KiB | [argm_16-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/argm_16-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-argm` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.1 KiB | [postgresql-16-argm_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/argm/postgresql-16-argm_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-argm` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.3 KiB | [postgresql-16-argm_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/argm/postgresql-16-argm_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-argm` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.1 KiB | [postgresql-16-argm_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/argm/postgresql-16-argm_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-argm` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.4 KiB | [postgresql-16-argm_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/argm/postgresql-16-argm_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-argm` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.9 KiB | [postgresql-16-argm_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/argm/postgresql-16-argm_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-argm` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.1 KiB | [postgresql-16-argm_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/argm/postgresql-16-argm_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-argm` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.1 KiB | [postgresql-16-argm_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/argm/postgresql-16-argm_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-argm` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.2 KiB | [postgresql-16-argm_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/argm/postgresql-16-argm_1.1.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-argm` | `1.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.0 KiB | [postgresql-16-argm_1.1.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/argm/postgresql-16-argm_1.1.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-argm` | `1.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 20.0 KiB | [postgresql-16-argm_1.1.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/argm/postgresql-16-argm_1.1.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `argm_15` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.3 KiB | [argm_15-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/argm_15-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `argm_15` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.4 KiB | [argm_15-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/argm_15-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `argm_15` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.2 KiB | [argm_15-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/argm_15-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `argm_15` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.3 KiB | [argm_15-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/argm_15-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `argm_15` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.1 KiB | [argm_15-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/argm_15-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `argm_15` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.4 KiB | [argm_15-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/argm_15-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-argm` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.1 KiB | [postgresql-15-argm_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/argm/postgresql-15-argm_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-argm` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.3 KiB | [postgresql-15-argm_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/argm/postgresql-15-argm_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-argm` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.1 KiB | [postgresql-15-argm_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/argm/postgresql-15-argm_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-argm` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.5 KiB | [postgresql-15-argm_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/argm/postgresql-15-argm_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-argm` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.9 KiB | [postgresql-15-argm_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/argm/postgresql-15-argm_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-argm` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.1 KiB | [postgresql-15-argm_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/argm/postgresql-15-argm_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-argm` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.1 KiB | [postgresql-15-argm_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/argm/postgresql-15-argm_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-argm` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.2 KiB | [postgresql-15-argm_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/argm/postgresql-15-argm_1.1.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-argm` | `1.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.0 KiB | [postgresql-15-argm_1.1.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/argm/postgresql-15-argm_1.1.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-argm` | `1.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 20.1 KiB | [postgresql-15-argm_1.1.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/argm/postgresql-15-argm_1.1.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `argm_14` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.3 KiB | [argm_14-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/argm_14-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `argm_14` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.4 KiB | [argm_14-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/argm_14-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `argm_14` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.1 KiB | [argm_14-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/argm_14-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `argm_14` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.2 KiB | [argm_14-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/argm_14-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `argm_14` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.1 KiB | [argm_14-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/argm_14-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `argm_14` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.4 KiB | [argm_14-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/argm_14-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-argm` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.0 KiB | [postgresql-14-argm_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/argm/postgresql-14-argm_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-argm` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.3 KiB | [postgresql-14-argm_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/argm/postgresql-14-argm_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-argm` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.0 KiB | [postgresql-14-argm_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/argm/postgresql-14-argm_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-argm` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.3 KiB | [postgresql-14-argm_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/argm/postgresql-14-argm_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-argm` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.9 KiB | [postgresql-14-argm_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/argm/postgresql-14-argm_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-argm` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.1 KiB | [postgresql-14-argm_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/argm/postgresql-14-argm_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-argm` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.0 KiB | [postgresql-14-argm_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/argm/postgresql-14-argm_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-argm` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.2 KiB | [postgresql-14-argm_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/argm/postgresql-14-argm_1.1.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-argm` | `1.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.0 KiB | [postgresql-14-argm_1.1.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/argm/postgresql-14-argm_1.1.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-argm` | `1.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 20.0 KiB | [postgresql-14-argm_1.1.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/argm/postgresql-14-argm_1.1.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bashtanov/argm" title="Repository" icon="github" subtitle="github.com/bashtanov/argm" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="argm-1.1.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg argm;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install argm;		# install via package name, for the active PG version

pig install argm -v 18;   # install for PG 18
pig install argm -v 17;   # install for PG 17
pig install argm -v 16;   # install for PG 16
pig install argm -v 15;   # install for PG 15
pig install argm -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION argm;
```

## Usage

Sources:

- [argm 1.1.1 README](https://github.com/bashtanov/argm/blob/1.1.1/README.md)
- [Extension control file](https://github.com/bashtanov/argm/blob/1.1.1/argm.control)
- [SQL definitions](https://github.com/bashtanov/argm/blob/1.1.1/argm--1.1.1.sql)

`argm` provides the polymorphic aggregates `argmax`, `argmin`, and `anyold`. They return a value from a selected row while grouping, avoiding a join or window-function pass when the row can be chosen by one or more sortable keys.

### Core Workflow

```sql
CREATE EXTENSION argm;

SELECT customer_id,
       argmax(order_id, total, ordered_at) AS largest_order
FROM orders
GROUP BY customer_id;
```

`argmax(value, key...)` returns the `value` belonging to the lexicographically greatest key tuple. `argmin` selects the least tuple. Additional keys break ties without building a composite value:

```sql
SELECT device_id,
       argmax(reading, measured_at, sequence_no) AS latest_reading
FROM measurements
GROUP BY device_id;
```

Use `anyold(value)` when any member of a group is acceptable:

```sql
SELECT account_id, anyold(display_name)
FROM account_aliases
GROUP BY account_id;
```

### Important Objects

- `argmax(value, key [, key ...])` selects the value associated with the greatest key tuple.
- `argmin(value, key [, key ...])` selects the value associated with the least key tuple.
- `anyold(value)` returns an arbitrary non-null value from the aggregate state.

The aggregates accept any value type; key types must have ordering support. The SQL definitions are parallel-safe and include combine and serialization functions for partial aggregation.

### Semantics and Caveats

Key tuples use one ordering direction and one collation for the whole tuple, with null keys sorted last. If complete key tuples tie, the chosen value is unspecified; add a stable final key when deterministic results matter. As with other PostgreSQL aggregates, empty input produces `NULL`.

`argm` 1.1.x requires PostgreSQL 9.6 or newer. The extension is relocatable. Upgrading from 1.0.3 to 1.1.x requires dropping and recreating the extension because the aggregate state changed; the 1.1.0-to-1.1.1 upgrade does not change the public SQL surface.
