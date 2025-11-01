---
title: "ddsketch"
linkTitle: "ddsketch"
description: "Provides ddsketch aggregate function"
weight: 4650
categories: ["FUNC"]
width: full
---

Provides ddsketch aggregate function


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4650** | {{< badge content="ddsketch" link="https://github.com/tvondra/ddsketch" >}} | {{< ext "ddsketch" >}} | `1.0.1` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "omnisketch" >}} {{< ext "quantile" >}} {{< ext "lower_quantile" >}} {{< ext "topn" >}} {{< ext "count_distinct" >}} {{< ext "hll" >}} {{< ext "first_last_agg" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/ddsketch" >}} | `1.0.1` | {{< bg "18" "ddsketch_18*" "green" >}} {{< bg "17" "ddsketch_17*" "green" >}} {{< bg "16" "ddsketch_16*" "green" >}} {{< bg "15" "ddsketch_15*" "green" >}} {{< bg "14" "ddsketch_14*" "green" >}} {{< bg "13" "ddsketch_13*" "green" >}} | `ddsketch_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/ddsketch" >}} | `1.0.1` | {{< bg "18" "postgresql-18-ddsketch" "green" >}} {{< bg "17" "postgresql-17-ddsketch" "green" >}} {{< bg "16" "postgresql-16-ddsketch" "green" >}} {{< bg "15" "postgresql-15-ddsketch" "green" >}} {{< bg "14" "postgresql-14-ddsketch" "green" >}} {{< bg "13" "postgresql-13-ddsketch" "green" >}} | `postgresql-$v-ddsketch` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 1.0.1" "ddsketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 1.0.1" "ddsketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 1.0.1" "ddsketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 1.0.1" "ddsketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    | {{< bg "PIGSTY 1.0.1" "ddsketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_13 : AVAIL 1" "green" >}} |
|    `el10.aarch64`    | {{< bg "PIGSTY 1.0.1" "ddsketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_13 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-ddsketch : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-ddsketch : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-ddsketch : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-ddsketch : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-ddsketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-ddsketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-ddsketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-ddsketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-ddsketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-ddsketch : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-ddsketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-ddsketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-ddsketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-ddsketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-ddsketch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-ddsketch : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-ddsketch : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-ddsketch : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-ddsketch : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-ddsketch : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-ddsketch : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-ddsketch : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-ddsketch : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-ddsketch : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddsketch_18` | 1.0.1 | `el8.x86_64` | pigsty | 34.7 KiB | [ddsketch_18-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_18-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `ddsketch_18` | 1.0.1 | `el8.aarch64` | pigsty | 33.0 KiB | [ddsketch_18-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_18-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `ddsketch_18` | 1.0.1 | `el9.x86_64` | pigsty | 34.0 KiB | [ddsketch_18-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_18-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `ddsketch_18` | 1.0.1 | `el9.aarch64` | pigsty | 32.9 KiB | [ddsketch_18-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_18-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `ddsketch_18` | 1.0.1 | `el10.x86_64` | pigsty | 34.2 KiB | [ddsketch_18-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddsketch_18-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `ddsketch_18` | 1.0.1 | `el10.aarch64` | pigsty | 33.3 KiB | [ddsketch_18-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddsketch_18-1.0.1-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddsketch_17` | 1.0.1 | `el8.x86_64` | pigsty | 34.7 KiB | [ddsketch_17-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_17-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `ddsketch_17` | 1.0.1 | `el8.aarch64` | pigsty | 33.0 KiB | [ddsketch_17-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_17-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `ddsketch_17` | 1.0.1 | `el9.x86_64` | pigsty | 34.0 KiB | [ddsketch_17-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_17-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `ddsketch_17` | 1.0.1 | `el9.aarch64` | pigsty | 32.9 KiB | [ddsketch_17-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_17-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `ddsketch_17` | 1.0.1 | `el10.x86_64` | pigsty | 34.2 KiB | [ddsketch_17-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddsketch_17-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `ddsketch_17` | 1.0.1 | `el10.aarch64` | pigsty | 33.3 KiB | [ddsketch_17-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddsketch_17-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-ddsketch` | 1.0.1 | `d12.x86_64` | pigsty | 66.0 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-ddsketch` | 1.0.1 | `d12.aarch64` | pigsty | 65.1 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-ddsketch` | 1.0.1 | `u22.x86_64` | pigsty | 69.7 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-ddsketch` | 1.0.1 | `u22.aarch64` | pigsty | 69.2 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-ddsketch` | 1.0.1 | `u24.x86_64` | pigsty | 64.2 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-ddsketch` | 1.0.1 | `u24.aarch64` | pigsty | 65.3 KiB | [postgresql-17-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddsketch_16` | 1.0.1 | `el8.x86_64` | pigsty | 34.7 KiB | [ddsketch_16-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_16-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `ddsketch_16` | 1.0.1 | `el8.aarch64` | pigsty | 33.0 KiB | [ddsketch_16-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_16-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `ddsketch_16` | 1.0.1 | `el9.x86_64` | pigsty | 34.0 KiB | [ddsketch_16-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_16-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `ddsketch_16` | 1.0.1 | `el9.aarch64` | pigsty | 32.9 KiB | [ddsketch_16-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_16-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `ddsketch_16` | 1.0.1 | `el10.x86_64` | pigsty | 34.2 KiB | [ddsketch_16-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddsketch_16-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `ddsketch_16` | 1.0.1 | `el10.aarch64` | pigsty | 33.3 KiB | [ddsketch_16-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddsketch_16-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-ddsketch` | 1.0.1 | `d12.x86_64` | pigsty | 66.0 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-ddsketch` | 1.0.1 | `d12.aarch64` | pigsty | 65.1 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-ddsketch` | 1.0.1 | `u22.x86_64` | pigsty | 69.7 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-ddsketch` | 1.0.1 | `u22.aarch64` | pigsty | 69.2 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-ddsketch` | 1.0.1 | `u24.x86_64` | pigsty | 64.2 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-ddsketch` | 1.0.1 | `u24.aarch64` | pigsty | 65.3 KiB | [postgresql-16-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddsketch_15` | 1.0.1 | `el8.x86_64` | pigsty | 34.7 KiB | [ddsketch_15-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_15-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `ddsketch_15` | 1.0.1 | `el8.aarch64` | pigsty | 33.0 KiB | [ddsketch_15-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_15-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `ddsketch_15` | 1.0.1 | `el9.x86_64` | pigsty | 34.0 KiB | [ddsketch_15-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_15-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `ddsketch_15` | 1.0.1 | `el9.aarch64` | pigsty | 32.8 KiB | [ddsketch_15-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_15-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `ddsketch_15` | 1.0.1 | `el10.x86_64` | pigsty | 34.2 KiB | [ddsketch_15-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddsketch_15-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `ddsketch_15` | 1.0.1 | `el10.aarch64` | pigsty | 33.3 KiB | [ddsketch_15-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddsketch_15-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-ddsketch` | 1.0.1 | `d12.x86_64` | pigsty | 66.0 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-ddsketch` | 1.0.1 | `d12.aarch64` | pigsty | 65.3 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-ddsketch` | 1.0.1 | `u22.x86_64` | pigsty | 69.7 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-ddsketch` | 1.0.1 | `u22.aarch64` | pigsty | 69.4 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-ddsketch` | 1.0.1 | `u24.x86_64` | pigsty | 64.2 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-ddsketch` | 1.0.1 | `u24.aarch64` | pigsty | 65.3 KiB | [postgresql-15-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddsketch_14` | 1.0.1 | `el8.x86_64` | pigsty | 34.7 KiB | [ddsketch_14-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_14-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `ddsketch_14` | 1.0.1 | `el8.aarch64` | pigsty | 33.0 KiB | [ddsketch_14-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_14-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `ddsketch_14` | 1.0.1 | `el9.x86_64` | pigsty | 34.0 KiB | [ddsketch_14-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_14-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `ddsketch_14` | 1.0.1 | `el9.aarch64` | pigsty | 32.8 KiB | [ddsketch_14-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_14-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `ddsketch_14` | 1.0.1 | `el10.x86_64` | pigsty | 34.2 KiB | [ddsketch_14-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddsketch_14-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `ddsketch_14` | 1.0.1 | `el10.aarch64` | pigsty | 33.3 KiB | [ddsketch_14-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddsketch_14-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-ddsketch` | 1.0.1 | `d12.x86_64` | pigsty | 65.9 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-ddsketch` | 1.0.1 | `d12.aarch64` | pigsty | 65.3 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-ddsketch` | 1.0.1 | `u22.x86_64` | pigsty | 69.7 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-ddsketch` | 1.0.1 | `u22.aarch64` | pigsty | 69.3 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-ddsketch` | 1.0.1 | `u24.x86_64` | pigsty | 64.2 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-ddsketch` | 1.0.1 | `u24.aarch64` | pigsty | 65.2 KiB | [postgresql-14-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-14-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddsketch_13` | 1.0.1 | `el8.x86_64` | pigsty | 34.1 KiB | [ddsketch_13-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_13-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `ddsketch_13` | 1.0.1 | `el8.aarch64` | pigsty | 33.0 KiB | [ddsketch_13-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_13-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `ddsketch_13` | 1.0.1 | `el9.x86_64` | pigsty | 34.0 KiB | [ddsketch_13-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_13-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `ddsketch_13` | 1.0.1 | `el9.aarch64` | pigsty | 32.8 KiB | [ddsketch_13-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_13-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `ddsketch_13` | 1.0.1 | `el10.x86_64` | pigsty | 34.1 KiB | [ddsketch_13-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddsketch_13-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `ddsketch_13` | 1.0.1 | `el10.aarch64` | pigsty | 33.3 KiB | [ddsketch_13-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddsketch_13-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-ddsketch` | 1.0.1 | `d12.x86_64` | pigsty | 65.8 KiB | [postgresql-13-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-13-ddsketch_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-ddsketch` | 1.0.1 | `d12.aarch64` | pigsty | 65.5 KiB | [postgresql-13-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-13-ddsketch_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-ddsketch` | 1.0.1 | `u22.x86_64` | pigsty | 69.4 KiB | [postgresql-13-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-13-ddsketch_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-ddsketch` | 1.0.1 | `u22.aarch64` | pigsty | 69.0 KiB | [postgresql-13-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-13-ddsketch_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-ddsketch` | 1.0.1 | `u24.x86_64` | pigsty | 64.0 KiB | [postgresql-13-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-13-ddsketch_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-ddsketch` | 1.0.1 | `u24.aarch64` | pigsty | 65.3 KiB | [postgresql-13-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-13-ddsketch_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/ddsketch" title="Repository" icon="github" subtitle="github.com/tvondra/ddsketch" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="ddsketch-1.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get ddsketch; # get ddsketch source code
pig build dep ddsketch; # install build dependencies
pig build pkg ddsketch; # build extension rpm or deb
pig build ext ddsketch; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install ddsketch; # install by extension name, for the current active PG version
pig ext install ddsketch; # install via package alias, for the active PG version
pig ext install ddsketch -v 18;   # install for PG 18
pig ext install ddsketch -v 17;   # install for PG 17
pig ext install ddsketch -v 16;   # install for PG 16
pig ext install ddsketch -v 15;   # install for PG 15
pig ext install ddsketch -v 14;   # install for PG 14
pig ext install ddsketch -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION ddsketch;
```

