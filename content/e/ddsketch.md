---
title: "ddsketch"
linkTitle: "ddsketch"
description: "Provides ddsketch aggregate function"
weight: 4650
categories: ["FUNC"]
width: full
---

[**ddsketch**](https://github.com/tvondra/ddsketch) : Provides ddsketch aggregate function


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4650** | {{< badge content="ddsketch" link="https://github.com/tvondra/ddsketch" >}} | {{< ext "ddsketch" >}} | `1.0.1` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "omnisketch" >}} {{< ext "quantile" >}} {{< ext "lower_quantile" >}} {{< ext "topn" >}} {{< ext "count_distinct" >}} {{< ext "hll" >}} {{< ext "first_last_agg" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `ddsketch` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.1` | {{< bg "18" "ddsketch_18" "green" >}} {{< bg "17" "ddsketch_17" "green" >}} {{< bg "16" "ddsketch_16" "green" >}} {{< bg "15" "ddsketch_15" "green" >}} {{< bg "14" "ddsketch_14" "green" >}} | `ddsketch_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.1` | {{< bg "18" "postgresql-18-ddsketch" "green" >}} {{< bg "17" "postgresql-17-ddsketch" "green" >}} {{< bg "16" "postgresql-16-ddsketch" "green" >}} {{< bg "15" "postgresql-15-ddsketch" "green" >}} {{< bg "14" "postgresql-14-ddsketch" "green" >}} | `postgresql-$v-ddsketch` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-ddsketch : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-ddsketch : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-ddsketch : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-ddsketch : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-ddsketch : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-ddsketch : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-ddsketch : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-ddsketch : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-ddsketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-ddsketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-ddsketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-ddsketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-ddsketch : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-ddsketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-ddsketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-ddsketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-ddsketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-ddsketch : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddsketch_18` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 34.7 KiB | [ddsketch_18-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_18-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `ddsketch_18` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.0 KiB | [ddsketch_18-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_18-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `ddsketch_18` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.0 KiB | [ddsketch_18-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_18-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `ddsketch_18` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 32.9 KiB | [ddsketch_18-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_18-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `ddsketch_18` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.2 KiB | [ddsketch_18-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddsketch_18-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `ddsketch_18` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.3 KiB | [ddsketch_18-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddsketch_18-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-ddsketch` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 60.7 KiB | [postgresql-18-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-18-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-ddsketch` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 60.4 KiB | [postgresql-18-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-18-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-ddsketch` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 60.6 KiB | [postgresql-18-ddsketch_1.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-18-ddsketch_1.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-ddsketch` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 60.7 KiB | [postgresql-18-ddsketch_1.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-18-ddsketch_1.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-ddsketch` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 64.8 KiB | [postgresql-18-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-18-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-ddsketch` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 65.6 KiB | [postgresql-18-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-18-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-ddsketch` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 63.5 KiB | [postgresql-18-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-18-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-ddsketch` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 64.0 KiB | [postgresql-18-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-18-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddsketch_17` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 34.7 KiB | [ddsketch_17-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_17-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `ddsketch_17` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.0 KiB | [ddsketch_17-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_17-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `ddsketch_17` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.0 KiB | [ddsketch_17-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_17-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `ddsketch_17` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 32.9 KiB | [ddsketch_17-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_17-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `ddsketch_17` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.2 KiB | [ddsketch_17-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddsketch_17-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `ddsketch_17` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.3 KiB | [ddsketch_17-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddsketch_17-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-ddsketch` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 60.7 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-ddsketch` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 60.4 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-ddsketch` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 60.5 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-ddsketch` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 60.7 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-ddsketch` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 69.7 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-ddsketch` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 69.2 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-ddsketch` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 63.6 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-ddsketch` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 64.0 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddsketch_16` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 34.7 KiB | [ddsketch_16-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_16-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `ddsketch_16` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.0 KiB | [ddsketch_16-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_16-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `ddsketch_16` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.0 KiB | [ddsketch_16-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_16-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `ddsketch_16` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 32.9 KiB | [ddsketch_16-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_16-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `ddsketch_16` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.2 KiB | [ddsketch_16-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddsketch_16-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `ddsketch_16` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.3 KiB | [ddsketch_16-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddsketch_16-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-ddsketch` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 60.7 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-ddsketch` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 60.4 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-ddsketch` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 60.6 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-ddsketch` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 60.8 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-ddsketch` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 69.7 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-ddsketch` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 69.2 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-ddsketch` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 63.6 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-ddsketch` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 64.0 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddsketch_15` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 34.7 KiB | [ddsketch_15-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_15-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `ddsketch_15` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.0 KiB | [ddsketch_15-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_15-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `ddsketch_15` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.0 KiB | [ddsketch_15-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_15-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `ddsketch_15` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 32.8 KiB | [ddsketch_15-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_15-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `ddsketch_15` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.2 KiB | [ddsketch_15-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddsketch_15-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `ddsketch_15` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.3 KiB | [ddsketch_15-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddsketch_15-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-ddsketch` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 60.7 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-ddsketch` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 60.4 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-ddsketch` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 60.4 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-ddsketch` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 60.6 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-ddsketch` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 69.7 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-ddsketch` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 69.4 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-ddsketch` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 63.5 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-ddsketch` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 64.0 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddsketch_14` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 34.7 KiB | [ddsketch_14-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_14-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `ddsketch_14` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.0 KiB | [ddsketch_14-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_14-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `ddsketch_14` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.0 KiB | [ddsketch_14-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_14-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `ddsketch_14` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 32.8 KiB | [ddsketch_14-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_14-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `ddsketch_14` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.2 KiB | [ddsketch_14-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddsketch_14-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `ddsketch_14` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.3 KiB | [ddsketch_14-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddsketch_14-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-ddsketch` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 60.6 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-ddsketch` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 60.3 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-ddsketch` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 60.5 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-ddsketch` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 60.6 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-ddsketch` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 69.7 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-ddsketch` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 69.3 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-ddsketch` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 63.5 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-ddsketch` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 63.9 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/ddsketch" title="Repository" icon="github" subtitle="github.com/tvondra/ddsketch" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="ddsketch-1.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg ddsketch;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install ddsketch;		# install via package name, for the active PG version

pig install ddsketch -v 18;   # install for PG 18
pig install ddsketch -v 17;   # install for PG 17
pig install ddsketch -v 16;   # install for PG 16
pig install ddsketch -v 15;   # install for PG 15
pig install ddsketch -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION ddsketch;
```



## Usage

> [ddsketch: DDSketch quantile estimation for PostgreSQL](https://github.com/tvondra/ddsketch)

Implements DDSketch, a fully-mergeable quantile sketch with relative-error guarantees. Much faster than `percentile_cont` and supports parallelism.

```sql
CREATE EXTENSION ddsketch;
```

### Direct Aggregation Functions

| Function | Description |
|---|---|
| `ddsketch_percentile(value, alpha, nbuckets, quantile)` | Estimate a single percentile |
| `ddsketch_percentile(value, alpha, nbuckets, quantiles[])` | Estimate multiple percentiles |
| `ddsketch_percentile_of(value, alpha, nbuckets, value)` | Estimate percentile rank of a value |
| `ddsketch_percentile_of(value, alpha, nbuckets, values[])` | Estimate percentile ranks of multiple values |

### Pre-aggregation Functions

| Function | Description |
|---|---|
| `ddsketch(value, alpha, nbuckets)` | Build a ddsketch from values |
| `ddsketch_percentile(sketch, quantile)` | Estimate percentile from a pre-built sketch |
| `ddsketch_percentile(sketch, quantiles[])` | Estimate multiple percentiles from a pre-built sketch |

### Utility Functions

| Function | Description |
|---|---|
| `ddsketch_count(sketch)` | Return the number of items in the sketch |
| `ddsketch_sum(sketch, low, high)` | Trimmed sum within a value range |
| `ddsketch_avg(sketch, low, high)` | Trimmed average within a value range |

### Parameters

- `alpha` -- controls accuracy and sketch size (lower = more accurate, larger)
- `nbuckets` -- maximum number of buckets (each 8 bytes)

### Examples

```sql
-- Instead of: SELECT percentile_cont(0.95) WITHIN GROUP (ORDER BY a) FROM t;
SELECT ddsketch_percentile(a, 0.05, 1024, 0.95) FROM t;

-- Multiple percentiles at once
SELECT ddsketch_percentile(a, 0.05, 1024, ARRAY[0.5, 0.95, 0.99]) FROM t;

-- Pre-aggregate for fast repeated queries
CREATE TABLE p AS SELECT a, b, ddsketch(c, 0.05, 1024) AS d FROM t GROUP BY a, b;

-- Query pre-aggregated data (~1.5ms vs ~7s for exact)
SELECT a, ddsketch_percentile(d, 0.95) FROM p GROUP BY a ORDER BY a;
```
