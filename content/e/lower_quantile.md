---
title: "lower_quantile"
linkTitle: "lower_quantile"
description: "Lower quantile aggregate function"
weight: 4620
categories: ["FUNC"]
width: full
---

[**lower_quantile**](https://github.com/tvondra/lower_quantile) : Lower quantile aggregate function


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4620** | {{< badge content="lower_quantile" link="https://github.com/tvondra/lower_quantile" >}} | {{< ext "lower_quantile" >}} | `1.0.3` | {{< category "FUNC" >}} | {{< license "BSD 2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "quantile" >}} {{< ext "topn" >}} {{< ext "ddsketch" >}} {{< ext "omnisketch" >}} {{< ext "count_distinct" >}} {{< ext "first_last_agg" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `lower_quantile` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.3` | {{< bg "18" "lower_quantile_18" "green" >}} {{< bg "17" "lower_quantile_17" "green" >}} {{< bg "16" "lower_quantile_16" "green" >}} {{< bg "15" "lower_quantile_15" "green" >}} {{< bg "14" "lower_quantile_14" "green" >}} {{< bg "13" "lower_quantile_13" "green" >}} | `lower_quantile_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.3` | {{< bg "18" "postgresql-18-lower-quantile" "green" >}} {{< bg "17" "postgresql-17-lower-quantile" "green" >}} {{< bg "16" "postgresql-16-lower-quantile" "green" >}} {{< bg "15" "postgresql-15-lower-quantile" "green" >}} {{< bg "14" "postgresql-14-lower-quantile" "green" >}} {{< bg "13" "postgresql-13-lower-quantile" "green" >}} | `postgresql-$v-lower-quantile` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "lower_quantile_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-13-lower-quantile : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-13-lower-quantile : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-13-lower-quantile : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-13-lower-quantile : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-13-lower-quantile : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-13-lower-quantile : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-13-lower-quantile : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-lower-quantile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-13-lower-quantile : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `lower_quantile_18` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.5 KiB | [lower_quantile_18-1.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/lower_quantile_18-1.0.3-1PIGSTY.el8.x86_64.rpm) |
| `lower_quantile_18` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.4 KiB | [lower_quantile_18-1.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/lower_quantile_18-1.0.3-1PIGSTY.el8.aarch64.rpm) |
| `lower_quantile_18` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.4 KiB | [lower_quantile_18-1.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/lower_quantile_18-1.0.3-1PIGSTY.el9.x86_64.rpm) |
| `lower_quantile_18` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.3 KiB | [lower_quantile_18-1.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/lower_quantile_18-1.0.3-1PIGSTY.el9.aarch64.rpm) |
| `lower_quantile_18` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.3 KiB | [lower_quantile_18-1.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/lower_quantile_18-1.0.3-1PIGSTY.el10.x86_64.rpm) |
| `lower_quantile_18` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.4 KiB | [lower_quantile_18-1.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/lower_quantile_18-1.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-lower-quantile` | `1.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.4 KiB | [postgresql-18-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-18-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-lower-quantile` | `1.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.2 KiB | [postgresql-18-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-18-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-lower-quantile` | `1.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.4 KiB | [postgresql-18-lower-quantile_1.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/lower-quantile/postgresql-18-lower-quantile_1.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-lower-quantile` | `1.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.2 KiB | [postgresql-18-lower-quantile_1.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/lower-quantile/postgresql-18-lower-quantile_1.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-lower-quantile` | `1.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.7 KiB | [postgresql-18-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-18-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-lower-quantile` | `1.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 16.3 KiB | [postgresql-18-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-18-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-lower-quantile` | `1.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.9 KiB | [postgresql-18-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-18-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-lower-quantile` | `1.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.5 KiB | [postgresql-18-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-18-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `lower_quantile_17` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.6 KiB | [lower_quantile_17-1.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/lower_quantile_17-1.0.3-1PIGSTY.el8.x86_64.rpm) |
| `lower_quantile_17` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.4 KiB | [lower_quantile_17-1.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/lower_quantile_17-1.0.3-1PIGSTY.el8.aarch64.rpm) |
| `lower_quantile_17` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.4 KiB | [lower_quantile_17-1.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/lower_quantile_17-1.0.3-1PIGSTY.el9.x86_64.rpm) |
| `lower_quantile_17` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.3 KiB | [lower_quantile_17-1.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/lower_quantile_17-1.0.3-1PIGSTY.el9.aarch64.rpm) |
| `lower_quantile_17` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.3 KiB | [lower_quantile_17-1.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/lower_quantile_17-1.0.3-1PIGSTY.el10.x86_64.rpm) |
| `lower_quantile_17` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.3 KiB | [lower_quantile_17-1.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/lower_quantile_17-1.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-lower-quantile` | `1.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.8 KiB | [postgresql-17-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-lower-quantile` | `1.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.6 KiB | [postgresql-17-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-lower-quantile` | `1.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.8 KiB | [postgresql-17-lower-quantile_1.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-lower-quantile` | `1.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.6 KiB | [postgresql-17-lower-quantile_1.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-lower-quantile` | `1.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.7 KiB | [postgresql-17-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-lower-quantile` | `1.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.3 KiB | [postgresql-17-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-lower-quantile` | `1.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.2 KiB | [postgresql-17-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-lower-quantile` | `1.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.9 KiB | [postgresql-17-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `lower_quantile_16` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.6 KiB | [lower_quantile_16-1.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/lower_quantile_16-1.0.3-1PIGSTY.el8.x86_64.rpm) |
| `lower_quantile_16` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.4 KiB | [lower_quantile_16-1.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/lower_quantile_16-1.0.3-1PIGSTY.el8.aarch64.rpm) |
| `lower_quantile_16` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.4 KiB | [lower_quantile_16-1.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/lower_quantile_16-1.0.3-1PIGSTY.el9.x86_64.rpm) |
| `lower_quantile_16` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.3 KiB | [lower_quantile_16-1.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/lower_quantile_16-1.0.3-1PIGSTY.el9.aarch64.rpm) |
| `lower_quantile_16` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.3 KiB | [lower_quantile_16-1.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/lower_quantile_16-1.0.3-1PIGSTY.el10.x86_64.rpm) |
| `lower_quantile_16` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.3 KiB | [lower_quantile_16-1.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/lower_quantile_16-1.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-lower-quantile` | `1.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.4 KiB | [postgresql-16-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-lower-quantile` | `1.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.2 KiB | [postgresql-16-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-lower-quantile` | `1.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.4 KiB | [postgresql-16-lower-quantile_1.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-lower-quantile` | `1.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.2 KiB | [postgresql-16-lower-quantile_1.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-lower-quantile` | `1.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.5 KiB | [postgresql-16-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-lower-quantile` | `1.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.1 KiB | [postgresql-16-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-lower-quantile` | `1.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.9 KiB | [postgresql-16-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-lower-quantile` | `1.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.5 KiB | [postgresql-16-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `lower_quantile_15` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.6 KiB | [lower_quantile_15-1.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/lower_quantile_15-1.0.3-1PIGSTY.el8.x86_64.rpm) |
| `lower_quantile_15` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.4 KiB | [lower_quantile_15-1.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/lower_quantile_15-1.0.3-1PIGSTY.el8.aarch64.rpm) |
| `lower_quantile_15` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.4 KiB | [lower_quantile_15-1.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/lower_quantile_15-1.0.3-1PIGSTY.el9.x86_64.rpm) |
| `lower_quantile_15` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.3 KiB | [lower_quantile_15-1.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/lower_quantile_15-1.0.3-1PIGSTY.el9.aarch64.rpm) |
| `lower_quantile_15` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.3 KiB | [lower_quantile_15-1.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/lower_quantile_15-1.0.3-1PIGSTY.el10.x86_64.rpm) |
| `lower_quantile_15` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.3 KiB | [lower_quantile_15-1.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/lower_quantile_15-1.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-lower-quantile` | `1.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.4 KiB | [postgresql-15-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-lower-quantile` | `1.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.2 KiB | [postgresql-15-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-lower-quantile` | `1.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.4 KiB | [postgresql-15-lower-quantile_1.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-lower-quantile` | `1.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.2 KiB | [postgresql-15-lower-quantile_1.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-lower-quantile` | `1.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.4 KiB | [postgresql-15-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-lower-quantile` | `1.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.0 KiB | [postgresql-15-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-lower-quantile` | `1.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.9 KiB | [postgresql-15-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-lower-quantile` | `1.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.5 KiB | [postgresql-15-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `lower_quantile_14` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.5 KiB | [lower_quantile_14-1.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/lower_quantile_14-1.0.3-1PIGSTY.el8.x86_64.rpm) |
| `lower_quantile_14` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.4 KiB | [lower_quantile_14-1.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/lower_quantile_14-1.0.3-1PIGSTY.el8.aarch64.rpm) |
| `lower_quantile_14` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.4 KiB | [lower_quantile_14-1.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/lower_quantile_14-1.0.3-1PIGSTY.el9.x86_64.rpm) |
| `lower_quantile_14` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.2 KiB | [lower_quantile_14-1.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/lower_quantile_14-1.0.3-1PIGSTY.el9.aarch64.rpm) |
| `lower_quantile_14` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.3 KiB | [lower_quantile_14-1.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/lower_quantile_14-1.0.3-1PIGSTY.el10.x86_64.rpm) |
| `lower_quantile_14` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.3 KiB | [lower_quantile_14-1.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/lower_quantile_14-1.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-lower-quantile` | `1.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.4 KiB | [postgresql-14-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-lower-quantile` | `1.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.2 KiB | [postgresql-14-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-lower-quantile` | `1.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.4 KiB | [postgresql-14-lower-quantile_1.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-lower-quantile` | `1.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.2 KiB | [postgresql-14-lower-quantile_1.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-lower-quantile` | `1.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.4 KiB | [postgresql-14-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-lower-quantile` | `1.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.0 KiB | [postgresql-14-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-lower-quantile` | `1.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.8 KiB | [postgresql-14-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-lower-quantile` | `1.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.5 KiB | [postgresql-14-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `lower_quantile_13` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.4 KiB | [lower_quantile_13-1.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/lower_quantile_13-1.0.3-1PIGSTY.el8.x86_64.rpm) |
| `lower_quantile_13` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.4 KiB | [lower_quantile_13-1.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/lower_quantile_13-1.0.3-1PIGSTY.el8.aarch64.rpm) |
| `lower_quantile_13` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.4 KiB | [lower_quantile_13-1.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/lower_quantile_13-1.0.3-1PIGSTY.el9.x86_64.rpm) |
| `lower_quantile_13` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.2 KiB | [lower_quantile_13-1.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/lower_quantile_13-1.0.3-1PIGSTY.el9.aarch64.rpm) |
| `lower_quantile_13` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.3 KiB | [lower_quantile_13-1.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/lower_quantile_13-1.0.3-1PIGSTY.el10.x86_64.rpm) |
| `lower_quantile_13` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.3 KiB | [lower_quantile_13-1.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/lower_quantile_13-1.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-lower-quantile` | `1.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.1 KiB | [postgresql-13-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-13-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-lower-quantile` | `1.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.9 KiB | [postgresql-13-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-13-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-lower-quantile` | `1.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.1 KiB | [postgresql-13-lower-quantile_1.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/lower-quantile/postgresql-13-lower-quantile_1.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-lower-quantile` | `1.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.9 KiB | [postgresql-13-lower-quantile_1.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/lower-quantile/postgresql-13-lower-quantile_1.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-lower-quantile` | `1.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.2 KiB | [postgresql-13-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-13-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-lower-quantile` | `1.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 16.9 KiB | [postgresql-13-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-13-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-lower-quantile` | `1.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.6 KiB | [postgresql-13-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-13-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-lower-quantile` | `1.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.1 KiB | [postgresql-13-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-13-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/lower_quantile" title="Repository" icon="github" subtitle="github.com/tvondra/lower_quantile" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="lower_quantile-1.0.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg lower_quantile;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install lower_quantile;		# install via package name, for the active PG version

pig install lower_quantile -v 18;   # install for PG 18
pig install lower_quantile -v 17;   # install for PG 17
pig install lower_quantile -v 16;   # install for PG 16
pig install lower_quantile -v 15;   # install for PG 15
pig install lower_quantile -v 14;   # install for PG 14
pig install lower_quantile -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION lower_quantile;
```
