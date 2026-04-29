---
title: "quantile"
linkTitle: "quantile"
description: "Quantile aggregation function"
weight: 4610
categories: ["FUNC"]
width: full
---

[**quantile**](https://github.com/tvondra/quantile) : Quantile aggregation function


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4610** | {{< badge content="quantile" link="https://github.com/tvondra/quantile" >}} | {{< ext "quantile" >}} | `1.1.8` | {{< category "FUNC" >}} | {{< license "BSD 2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "lower_quantile" >}} {{< ext "topn" >}} {{< ext "ddsketch" >}} {{< ext "omnisketch" >}} {{< ext "count_distinct" >}} {{< ext "first_last_agg" >}} {{< ext "aggs_for_arrays" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.8` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `quantile` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.8` | {{< bg "18" "quantile_18" "green" >}} {{< bg "17" "quantile_17" "green" >}} {{< bg "16" "quantile_16" "green" >}} {{< bg "15" "quantile_15" "green" >}} {{< bg "14" "quantile_14" "green" >}} | `quantile_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.8` | {{< bg "18" "postgresql-18-quantile" "green" >}} {{< bg "17" "postgresql-17-quantile" "green" >}} {{< bg "16" "postgresql-16-quantile" "green" >}} {{< bg "15" "postgresql-15-quantile" "green" >}} {{< bg "14" "postgresql-14-quantile" "green" >}} | `postgresql-$v-quantile` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1.8" "quantile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1.8" "quantile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1.8" "quantile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1.8" "quantile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1.8" "quantile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1.8" "quantile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "quantile_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-18-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-17-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-16-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-15-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-14-quantile : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-18-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-17-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-16-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-15-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-14-quantile : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-18-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-17-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-16-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-15-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-14-quantile : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-18-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-17-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-16-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-15-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-14-quantile : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-18-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-17-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-16-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-15-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-14-quantile : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-18-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-17-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-16-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-15-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-14-quantile : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-18-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-17-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-16-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-15-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-14-quantile : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-18-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-17-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-16-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-15-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.8" "postgresql-14-quantile : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-quantile : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-quantile : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-quantile : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-quantile : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-quantile : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-quantile : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-quantile : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-quantile : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-quantile : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-quantile : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `quantile_18` | `1.1.8` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.5 KiB | [quantile_18-1.1.8-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/quantile_18-1.1.8-1PIGSTY.el8.x86_64.rpm) |
| `quantile_18` | `1.1.8` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.8 KiB | [quantile_18-1.1.8-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/quantile_18-1.1.8-1PIGSTY.el8.aarch64.rpm) |
| `quantile_18` | `1.1.8` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.1 KiB | [quantile_18-1.1.8-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/quantile_18-1.1.8-1PIGSTY.el9.x86_64.rpm) |
| `quantile_18` | `1.1.8` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.5 KiB | [quantile_18-1.1.8-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/quantile_18-1.1.8-1PIGSTY.el9.aarch64.rpm) |
| `quantile_18` | `1.1.8` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.1 KiB | [quantile_18-1.1.8-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/quantile_18-1.1.8-1PIGSTY.el10.x86_64.rpm) |
| `quantile_18` | `1.1.8` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.8 KiB | [quantile_18-1.1.8-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/quantile_18-1.1.8-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-quantile` | `1.1.8` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.5 KiB | [postgresql-18-quantile_1.1.8-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/quantile/postgresql-18-quantile_1.1.8-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-quantile` | `1.1.8` | [d12.aarch64](/os/d12.aarch64) | pigsty | 21.2 KiB | [postgresql-18-quantile_1.1.8-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/quantile/postgresql-18-quantile_1.1.8-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-quantile` | `1.1.8` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.5 KiB | [postgresql-18-quantile_1.1.8-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/quantile/postgresql-18-quantile_1.1.8-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-quantile` | `1.1.8` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.3 KiB | [postgresql-18-quantile_1.1.8-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/quantile/postgresql-18-quantile_1.1.8-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-quantile` | `1.1.8` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.7 KiB | [postgresql-18-quantile_1.1.8-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/quantile/postgresql-18-quantile_1.1.8-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-quantile` | `1.1.8` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.8 KiB | [postgresql-18-quantile_1.1.8-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/quantile/postgresql-18-quantile_1.1.8-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-quantile` | `1.1.8` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.6 KiB | [postgresql-18-quantile_1.1.8-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/quantile/postgresql-18-quantile_1.1.8-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-quantile` | `1.1.8` | [u24.aarch64](/os/u24.aarch64) | pigsty | 22.2 KiB | [postgresql-18-quantile_1.1.8-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/quantile/postgresql-18-quantile_1.1.8-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `quantile_17` | `1.1.8` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.5 KiB | [quantile_17-1.1.8-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/quantile_17-1.1.8-1PIGSTY.el8.x86_64.rpm) |
| `quantile_17` | `1.1.8` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.8 KiB | [quantile_17-1.1.8-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/quantile_17-1.1.8-1PIGSTY.el8.aarch64.rpm) |
| `quantile_17` | `1.1.8` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.2 KiB | [quantile_17-1.1.8-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/quantile_17-1.1.8-1PIGSTY.el9.x86_64.rpm) |
| `quantile_17` | `1.1.8` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.5 KiB | [quantile_17-1.1.8-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/quantile_17-1.1.8-1PIGSTY.el9.aarch64.rpm) |
| `quantile_17` | `1.1.8` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.1 KiB | [quantile_17-1.1.8-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/quantile_17-1.1.8-1PIGSTY.el10.x86_64.rpm) |
| `quantile_17` | `1.1.8` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.8 KiB | [quantile_17-1.1.8-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/quantile_17-1.1.8-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-quantile` | `1.1.8` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.2 KiB | [postgresql-17-quantile_1.1.8-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/quantile/postgresql-17-quantile_1.1.8-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-quantile` | `1.1.8` | [d12.aarch64](/os/d12.aarch64) | pigsty | 20.9 KiB | [postgresql-17-quantile_1.1.8-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/quantile/postgresql-17-quantile_1.1.8-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-quantile` | `1.1.8` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.2 KiB | [postgresql-17-quantile_1.1.8-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/quantile/postgresql-17-quantile_1.1.8-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-quantile` | `1.1.8` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.0 KiB | [postgresql-17-quantile_1.1.8-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/quantile/postgresql-17-quantile_1.1.8-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-quantile` | `1.1.8` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.9 KiB | [postgresql-17-quantile_1.1.8-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/quantile/postgresql-17-quantile_1.1.8-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-quantile` | `1.1.8` | [u22.aarch64](/os/u22.aarch64) | pigsty | 23.0 KiB | [postgresql-17-quantile_1.1.8-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/quantile/postgresql-17-quantile_1.1.8-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-quantile` | `1.1.8` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.3 KiB | [postgresql-17-quantile_1.1.8-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/quantile/postgresql-17-quantile_1.1.8-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-quantile` | `1.1.8` | [u24.aarch64](/os/u24.aarch64) | pigsty | 22.0 KiB | [postgresql-17-quantile_1.1.8-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/quantile/postgresql-17-quantile_1.1.8-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `quantile_16` | `1.1.8` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.5 KiB | [quantile_16-1.1.8-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/quantile_16-1.1.8-1PIGSTY.el8.x86_64.rpm) |
| `quantile_16` | `1.1.8` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.8 KiB | [quantile_16-1.1.8-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/quantile_16-1.1.8-1PIGSTY.el8.aarch64.rpm) |
| `quantile_16` | `1.1.8` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.2 KiB | [quantile_16-1.1.8-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/quantile_16-1.1.8-1PIGSTY.el9.x86_64.rpm) |
| `quantile_16` | `1.1.8` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.5 KiB | [quantile_16-1.1.8-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/quantile_16-1.1.8-1PIGSTY.el9.aarch64.rpm) |
| `quantile_16` | `1.1.8` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.1 KiB | [quantile_16-1.1.8-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/quantile_16-1.1.8-1PIGSTY.el10.x86_64.rpm) |
| `quantile_16` | `1.1.8` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.8 KiB | [quantile_16-1.1.8-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/quantile_16-1.1.8-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-quantile` | `1.1.8` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.2 KiB | [postgresql-16-quantile_1.1.8-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/quantile/postgresql-16-quantile_1.1.8-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-quantile` | `1.1.8` | [d12.aarch64](/os/d12.aarch64) | pigsty | 20.9 KiB | [postgresql-16-quantile_1.1.8-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/quantile/postgresql-16-quantile_1.1.8-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-quantile` | `1.1.8` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.2 KiB | [postgresql-16-quantile_1.1.8-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/quantile/postgresql-16-quantile_1.1.8-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-quantile` | `1.1.8` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.0 KiB | [postgresql-16-quantile_1.1.8-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/quantile/postgresql-16-quantile_1.1.8-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-quantile` | `1.1.8` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.8 KiB | [postgresql-16-quantile_1.1.8-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/quantile/postgresql-16-quantile_1.1.8-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-quantile` | `1.1.8` | [u22.aarch64](/os/u22.aarch64) | pigsty | 23.0 KiB | [postgresql-16-quantile_1.1.8-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/quantile/postgresql-16-quantile_1.1.8-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-quantile` | `1.1.8` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.2 KiB | [postgresql-16-quantile_1.1.8-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/quantile/postgresql-16-quantile_1.1.8-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-quantile` | `1.1.8` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.9 KiB | [postgresql-16-quantile_1.1.8-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/quantile/postgresql-16-quantile_1.1.8-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `quantile_15` | `1.1.8` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.5 KiB | [quantile_15-1.1.8-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/quantile_15-1.1.8-1PIGSTY.el8.x86_64.rpm) |
| `quantile_15` | `1.1.8` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.8 KiB | [quantile_15-1.1.8-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/quantile_15-1.1.8-1PIGSTY.el8.aarch64.rpm) |
| `quantile_15` | `1.1.8` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.2 KiB | [quantile_15-1.1.8-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/quantile_15-1.1.8-1PIGSTY.el9.x86_64.rpm) |
| `quantile_15` | `1.1.8` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.5 KiB | [quantile_15-1.1.8-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/quantile_15-1.1.8-1PIGSTY.el9.aarch64.rpm) |
| `quantile_15` | `1.1.8` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.1 KiB | [quantile_15-1.1.8-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/quantile_15-1.1.8-1PIGSTY.el10.x86_64.rpm) |
| `quantile_15` | `1.1.8` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.8 KiB | [quantile_15-1.1.8-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/quantile_15-1.1.8-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-quantile` | `1.1.8` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.2 KiB | [postgresql-15-quantile_1.1.8-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/quantile/postgresql-15-quantile_1.1.8-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-quantile` | `1.1.8` | [d12.aarch64](/os/d12.aarch64) | pigsty | 20.9 KiB | [postgresql-15-quantile_1.1.8-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/quantile/postgresql-15-quantile_1.1.8-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-quantile` | `1.1.8` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.2 KiB | [postgresql-15-quantile_1.1.8-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/quantile/postgresql-15-quantile_1.1.8-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-quantile` | `1.1.8` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.0 KiB | [postgresql-15-quantile_1.1.8-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/quantile/postgresql-15-quantile_1.1.8-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-quantile` | `1.1.8` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.6 KiB | [postgresql-15-quantile_1.1.8-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/quantile/postgresql-15-quantile_1.1.8-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-quantile` | `1.1.8` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.7 KiB | [postgresql-15-quantile_1.1.8-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/quantile/postgresql-15-quantile_1.1.8-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-quantile` | `1.1.8` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.2 KiB | [postgresql-15-quantile_1.1.8-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/quantile/postgresql-15-quantile_1.1.8-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-quantile` | `1.1.8` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.9 KiB | [postgresql-15-quantile_1.1.8-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/quantile/postgresql-15-quantile_1.1.8-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `quantile_14` | `1.1.8` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.5 KiB | [quantile_14-1.1.8-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/quantile_14-1.1.8-1PIGSTY.el8.x86_64.rpm) |
| `quantile_14` | `1.1.8` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.8 KiB | [quantile_14-1.1.8-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/quantile_14-1.1.8-1PIGSTY.el8.aarch64.rpm) |
| `quantile_14` | `1.1.8` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.1 KiB | [quantile_14-1.1.8-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/quantile_14-1.1.8-1PIGSTY.el9.x86_64.rpm) |
| `quantile_14` | `1.1.8` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.5 KiB | [quantile_14-1.1.8-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/quantile_14-1.1.8-1PIGSTY.el9.aarch64.rpm) |
| `quantile_14` | `1.1.8` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.1 KiB | [quantile_14-1.1.8-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/quantile_14-1.1.8-1PIGSTY.el10.x86_64.rpm) |
| `quantile_14` | `1.1.8` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.8 KiB | [quantile_14-1.1.8-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/quantile_14-1.1.8-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-quantile` | `1.1.8` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.3 KiB | [postgresql-14-quantile_1.1.8-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/quantile/postgresql-14-quantile_1.1.8-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-quantile` | `1.1.8` | [d12.aarch64](/os/d12.aarch64) | pigsty | 21.0 KiB | [postgresql-14-quantile_1.1.8-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/quantile/postgresql-14-quantile_1.1.8-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-quantile` | `1.1.8` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.3 KiB | [postgresql-14-quantile_1.1.8-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/quantile/postgresql-14-quantile_1.1.8-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-quantile` | `1.1.8` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.1 KiB | [postgresql-14-quantile_1.1.8-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/quantile/postgresql-14-quantile_1.1.8-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-quantile` | `1.1.8` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.6 KiB | [postgresql-14-quantile_1.1.8-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/quantile/postgresql-14-quantile_1.1.8-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-quantile` | `1.1.8` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.7 KiB | [postgresql-14-quantile_1.1.8-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/quantile/postgresql-14-quantile_1.1.8-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-quantile` | `1.1.8` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.3 KiB | [postgresql-14-quantile_1.1.8-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/quantile/postgresql-14-quantile_1.1.8-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-quantile` | `1.1.8` | [u24.aarch64](/os/u24.aarch64) | pigsty | 22.0 KiB | [postgresql-14-quantile_1.1.8-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/quantile/postgresql-14-quantile_1.1.8-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/quantile" title="Repository" icon="github" subtitle="github.com/tvondra/quantile" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="quantile-1.1.8.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg quantile;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install quantile;		# install via package name, for the active PG version

pig install quantile -v 18;   # install for PG 18
pig install quantile -v 17;   # install for PG 17
pig install quantile -v 16;   # install for PG 16
pig install quantile -v 15;   # install for PG 15
pig install quantile -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION quantile;
```



## Usage

> [quantile: quantile aggregate functions for PostgreSQL](https://github.com/tvondra/quantile)

Provides aggregate functions to compute quantiles. Overloaded for `int`, `bigint`, `double precision`, and `numeric`.

```sql
CREATE EXTENSION quantile;
```

### Functions

| Function | Description |
|---|---|
| `quantile(value, quantile float)` | Compute a single quantile (0 to 1) |
| `quantile(value, quantiles float[])` | Compute multiple quantiles at once, returns array |

### Examples

```sql
-- Compute the median (0.5 quantile)
SELECT quantile(i, 0.5) FROM generate_series(1, 1000) s(i);
-- 500

-- Compute the 95th percentile
SELECT quantile(i, 0.95) FROM generate_series(1, 1000) s(i);

-- Compute all quartiles at once (more efficient than separate calls)
SELECT quantile(i, ARRAY[0.25, 0.5, 0.75])
FROM generate_series(1, 1000) s(i);
-- {250, 500, 750}
```

Note: Since PostgreSQL 9.4, built-in `percentile_cont` and `percentile_disc` functions are available. Consider using those first and only use this extension if it provides measurably better performance for your workload.
