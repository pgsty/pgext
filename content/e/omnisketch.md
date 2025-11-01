---
title: "omnisketch"
linkTitle: "omnisketch"
description: "data structure for on-line agg of data into approximate sketch"
weight: 4640
categories: ["FUNC"]
width: full
---

data structure for on-line agg of data into approximate sketch


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4640** | {{< badge content="omnisketch" link="https://github.com/tvondra/omnisketch" >}} | {{< ext "omnisketch" >}} | `1.0.2` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "ddsketch" >}} {{< ext "hll" >}} {{< ext "count_distinct" >}} {{< ext "topn" >}} {{< ext "quantile" >}} {{< ext "lower_quantile" >}} {{< ext "first_last_agg" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/omnisketch" >}} | `1.0.2` | {{< bg "18" "omnisketch_18*" "green" >}} {{< bg "17" "omnisketch_17*" "green" >}} {{< bg "16" "omnisketch_16*" "green" >}} {{< bg "15" "omnisketch_15*" "green" >}} {{< bg "14" "omnisketch_14*" "green" >}} {{< bg "13" "omnisketch_13*" "green" >}} | `omnisketch_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/omnisketch" >}} | `1.0.2` | {{< bg "18" "postgresql-18-omnisketch" "green" >}} {{< bg "17" "postgresql-17-omnisketch" "green" >}} {{< bg "16" "postgresql-16-omnisketch" "green" >}} {{< bg "15" "postgresql-15-omnisketch" "green" >}} {{< bg "14" "postgresql-14-omnisketch" "green" >}} {{< bg "13" "postgresql-13-omnisketch" "green" >}} | `postgresql-$v-omnisketch` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 1.0.2" "omnisketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 1.0.2" "omnisketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 1.0.2" "omnisketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 1.0.2" "omnisketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    | {{< bg "PIGSTY 1.0.2" "omnisketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_13 : AVAIL 1" "green" >}} |
|    `el10.aarch64`    | {{< bg "PIGSTY 1.0.2" "omnisketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "omnisketch_13 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-omnisketch : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.2" "postgresql-17-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-13-omnisketch : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-omnisketch : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.2" "postgresql-17-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-13-omnisketch : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-omnisketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-omnisketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-omnisketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-omnisketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-omnisketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-omnisketch : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-omnisketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-omnisketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-omnisketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-omnisketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-omnisketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-omnisketch : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-omnisketch : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.2" "postgresql-17-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-13-omnisketch : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-omnisketch : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.2" "postgresql-17-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-13-omnisketch : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-omnisketch : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.2" "postgresql-17-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-13-omnisketch : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-omnisketch : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.2" "postgresql-17-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-omnisketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-13-omnisketch : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `omnisketch_18` | 1.0.2 | `el8.x86_64` | pigsty | 20.4 KiB | [omnisketch_18-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/omnisketch_18-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `omnisketch_18` | 1.0.2 | `el8.aarch64` | pigsty | 20.3 KiB | [omnisketch_18-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/omnisketch_18-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `omnisketch_18` | 1.0.2 | `el9.x86_64` | pigsty | 20.1 KiB | [omnisketch_18-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/omnisketch_18-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `omnisketch_18` | 1.0.2 | `el9.aarch64` | pigsty | 20.3 KiB | [omnisketch_18-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/omnisketch_18-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `omnisketch_18` | 1.0.2 | `el10.x86_64` | pigsty | 20.2 KiB | [omnisketch_18-1.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/omnisketch_18-1.0.2-1PIGSTY.el10.x86_64.rpm) |
| `omnisketch_18` | 1.0.2 | `el10.aarch64` | pigsty | 20.5 KiB | [omnisketch_18-1.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/omnisketch_18-1.0.2-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `omnisketch_17` | 1.0.2 | `el8.x86_64` | pigsty | 20.4 KiB | [omnisketch_17-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/omnisketch_17-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `omnisketch_17` | 1.0.2 | `el8.aarch64` | pigsty | 20.4 KiB | [omnisketch_17-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/omnisketch_17-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `omnisketch_17` | 1.0.2 | `el9.x86_64` | pigsty | 20.1 KiB | [omnisketch_17-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/omnisketch_17-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `omnisketch_17` | 1.0.2 | `el9.aarch64` | pigsty | 20.2 KiB | [omnisketch_17-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/omnisketch_17-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `omnisketch_17` | 1.0.2 | `el10.x86_64` | pigsty | 20.2 KiB | [omnisketch_17-1.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/omnisketch_17-1.0.2-1PIGSTY.el10.x86_64.rpm) |
| `omnisketch_17` | 1.0.2 | `el10.aarch64` | pigsty | 20.5 KiB | [omnisketch_17-1.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/omnisketch_17-1.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-omnisketch` | 1.0.2 | `d12.x86_64` | pigsty | 27.6 KiB | [postgresql-17-omnisketch_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnisketch/postgresql-17-omnisketch_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-omnisketch` | 1.0.2 | `d12.aarch64` | pigsty | 27.6 KiB | [postgresql-17-omnisketch_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnisketch/postgresql-17-omnisketch_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-omnisketch` | 1.0.2 | `u22.x86_64` | pigsty | 28.9 KiB | [postgresql-17-omnisketch_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnisketch/postgresql-17-omnisketch_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-omnisketch` | 1.0.2 | `u22.aarch64` | pigsty | 28.5 KiB | [postgresql-17-omnisketch_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnisketch/postgresql-17-omnisketch_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-omnisketch` | 1.0.2 | `u24.x86_64` | pigsty | 27.1 KiB | [postgresql-17-omnisketch_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnisketch/postgresql-17-omnisketch_1.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-omnisketch` | 1.0.2 | `u24.aarch64` | pigsty | 27.3 KiB | [postgresql-17-omnisketch_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnisketch/postgresql-17-omnisketch_1.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `omnisketch_16` | 1.0.2 | `el8.x86_64` | pigsty | 20.4 KiB | [omnisketch_16-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/omnisketch_16-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `omnisketch_16` | 1.0.2 | `el8.aarch64` | pigsty | 20.4 KiB | [omnisketch_16-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/omnisketch_16-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `omnisketch_16` | 1.0.2 | `el9.x86_64` | pigsty | 20.1 KiB | [omnisketch_16-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/omnisketch_16-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `omnisketch_16` | 1.0.2 | `el9.aarch64` | pigsty | 20.2 KiB | [omnisketch_16-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/omnisketch_16-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `omnisketch_16` | 1.0.2 | `el10.x86_64` | pigsty | 20.2 KiB | [omnisketch_16-1.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/omnisketch_16-1.0.2-1PIGSTY.el10.x86_64.rpm) |
| `omnisketch_16` | 1.0.2 | `el10.aarch64` | pigsty | 20.5 KiB | [omnisketch_16-1.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/omnisketch_16-1.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-omnisketch` | 1.0.2 | `d12.x86_64` | pigsty | 27.6 KiB | [postgresql-16-omnisketch_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnisketch/postgresql-16-omnisketch_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-omnisketch` | 1.0.2 | `d12.aarch64` | pigsty | 27.6 KiB | [postgresql-16-omnisketch_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnisketch/postgresql-16-omnisketch_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-omnisketch` | 1.0.2 | `u22.x86_64` | pigsty | 28.9 KiB | [postgresql-16-omnisketch_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnisketch/postgresql-16-omnisketch_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-omnisketch` | 1.0.2 | `u22.aarch64` | pigsty | 28.4 KiB | [postgresql-16-omnisketch_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnisketch/postgresql-16-omnisketch_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-omnisketch` | 1.0.2 | `u24.x86_64` | pigsty | 27.2 KiB | [postgresql-16-omnisketch_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnisketch/postgresql-16-omnisketch_1.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-omnisketch` | 1.0.2 | `u24.aarch64` | pigsty | 27.4 KiB | [postgresql-16-omnisketch_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnisketch/postgresql-16-omnisketch_1.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `omnisketch_15` | 1.0.2 | `el8.x86_64` | pigsty | 20.4 KiB | [omnisketch_15-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/omnisketch_15-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `omnisketch_15` | 1.0.2 | `el8.aarch64` | pigsty | 20.4 KiB | [omnisketch_15-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/omnisketch_15-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `omnisketch_15` | 1.0.2 | `el9.x86_64` | pigsty | 20.1 KiB | [omnisketch_15-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/omnisketch_15-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `omnisketch_15` | 1.0.2 | `el9.aarch64` | pigsty | 20.2 KiB | [omnisketch_15-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/omnisketch_15-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `omnisketch_15` | 1.0.2 | `el10.x86_64` | pigsty | 20.2 KiB | [omnisketch_15-1.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/omnisketch_15-1.0.2-1PIGSTY.el10.x86_64.rpm) |
| `omnisketch_15` | 1.0.2 | `el10.aarch64` | pigsty | 20.5 KiB | [omnisketch_15-1.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/omnisketch_15-1.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-omnisketch` | 1.0.2 | `d12.x86_64` | pigsty | 27.7 KiB | [postgresql-15-omnisketch_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnisketch/postgresql-15-omnisketch_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-omnisketch` | 1.0.2 | `d12.aarch64` | pigsty | 27.6 KiB | [postgresql-15-omnisketch_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnisketch/postgresql-15-omnisketch_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-omnisketch` | 1.0.2 | `u22.x86_64` | pigsty | 28.9 KiB | [postgresql-15-omnisketch_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnisketch/postgresql-15-omnisketch_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-omnisketch` | 1.0.2 | `u22.aarch64` | pigsty | 28.5 KiB | [postgresql-15-omnisketch_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnisketch/postgresql-15-omnisketch_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-omnisketch` | 1.0.2 | `u24.x86_64` | pigsty | 27.2 KiB | [postgresql-15-omnisketch_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnisketch/postgresql-15-omnisketch_1.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-omnisketch` | 1.0.2 | `u24.aarch64` | pigsty | 27.4 KiB | [postgresql-15-omnisketch_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnisketch/postgresql-15-omnisketch_1.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `omnisketch_14` | 1.0.2 | `el8.x86_64` | pigsty | 20.4 KiB | [omnisketch_14-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/omnisketch_14-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `omnisketch_14` | 1.0.2 | `el8.aarch64` | pigsty | 20.3 KiB | [omnisketch_14-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/omnisketch_14-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `omnisketch_14` | 1.0.2 | `el9.x86_64` | pigsty | 20.1 KiB | [omnisketch_14-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/omnisketch_14-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `omnisketch_14` | 1.0.2 | `el9.aarch64` | pigsty | 20.2 KiB | [omnisketch_14-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/omnisketch_14-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `omnisketch_14` | 1.0.2 | `el10.x86_64` | pigsty | 20.2 KiB | [omnisketch_14-1.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/omnisketch_14-1.0.2-1PIGSTY.el10.x86_64.rpm) |
| `omnisketch_14` | 1.0.2 | `el10.aarch64` | pigsty | 20.5 KiB | [omnisketch_14-1.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/omnisketch_14-1.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-omnisketch` | 1.0.2 | `d12.x86_64` | pigsty | 27.6 KiB | [postgresql-14-omnisketch_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnisketch/postgresql-14-omnisketch_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-omnisketch` | 1.0.2 | `d12.aarch64` | pigsty | 27.7 KiB | [postgresql-14-omnisketch_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnisketch/postgresql-14-omnisketch_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-omnisketch` | 1.0.2 | `u22.x86_64` | pigsty | 28.8 KiB | [postgresql-14-omnisketch_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnisketch/postgresql-14-omnisketch_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-omnisketch` | 1.0.2 | `u22.aarch64` | pigsty | 28.6 KiB | [postgresql-14-omnisketch_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnisketch/postgresql-14-omnisketch_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-omnisketch` | 1.0.2 | `u24.x86_64` | pigsty | 27.2 KiB | [postgresql-14-omnisketch_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnisketch/postgresql-14-omnisketch_1.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-omnisketch` | 1.0.2 | `u24.aarch64` | pigsty | 27.0 KiB | [postgresql-14-omnisketch_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnisketch/postgresql-14-omnisketch_1.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `omnisketch_13` | 1.0.2 | `el8.x86_64` | pigsty | 20.2 KiB | [omnisketch_13-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/omnisketch_13-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `omnisketch_13` | 1.0.2 | `el8.aarch64` | pigsty | 20.3 KiB | [omnisketch_13-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/omnisketch_13-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `omnisketch_13` | 1.0.2 | `el9.x86_64` | pigsty | 20.1 KiB | [omnisketch_13-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/omnisketch_13-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `omnisketch_13` | 1.0.2 | `el9.aarch64` | pigsty | 20.2 KiB | [omnisketch_13-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/omnisketch_13-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `omnisketch_13` | 1.0.2 | `el10.x86_64` | pigsty | 20.2 KiB | [omnisketch_13-1.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/omnisketch_13-1.0.2-1PIGSTY.el10.x86_64.rpm) |
| `omnisketch_13` | 1.0.2 | `el10.aarch64` | pigsty | 20.4 KiB | [omnisketch_13-1.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/omnisketch_13-1.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-omnisketch` | 1.0.2 | `d12.x86_64` | pigsty | 27.6 KiB | [postgresql-13-omnisketch_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnisketch/postgresql-13-omnisketch_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-omnisketch` | 1.0.2 | `d12.aarch64` | pigsty | 27.4 KiB | [postgresql-13-omnisketch_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/omnisketch/postgresql-13-omnisketch_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-omnisketch` | 1.0.2 | `u22.x86_64` | pigsty | 28.8 KiB | [postgresql-13-omnisketch_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnisketch/postgresql-13-omnisketch_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-omnisketch` | 1.0.2 | `u22.aarch64` | pigsty | 28.4 KiB | [postgresql-13-omnisketch_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/omnisketch/postgresql-13-omnisketch_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-omnisketch` | 1.0.2 | `u24.x86_64` | pigsty | 27.1 KiB | [postgresql-13-omnisketch_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnisketch/postgresql-13-omnisketch_1.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-omnisketch` | 1.0.2 | `u24.aarch64` | pigsty | 27.3 KiB | [postgresql-13-omnisketch_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/omnisketch/postgresql-13-omnisketch_1.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/omnisketch" title="Repository" icon="github" subtitle="github.com/tvondra/omnisketch" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="omnisketch-1.0.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build get omnisketch; # get omnisketch source code
pig build dep omnisketch; # install build dependencies
pig build pkg omnisketch; # build extension rpm or deb
pig build ext omnisketch; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install omnisketch; # install by extension name, for the current active PG version
pig ext install omnisketch; # install via package alias, for the active PG version
pig ext install omnisketch -v 18;   # install for PG 18
pig ext install omnisketch -v 17;   # install for PG 17
pig ext install omnisketch -v 16;   # install for PG 16
pig ext install omnisketch -v 15;   # install for PG 15
pig ext install omnisketch -v 14;   # install for PG 14
pig ext install omnisketch -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION omnisketch;
```

