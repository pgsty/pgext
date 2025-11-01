---
title: "vchord_bm25"
linkTitle: "vchord_bm25"
description: "A postgresql extension for bm25 ranking algorithm"
weight: 2150
categories: ["FTS"]
width: full
---

A postgresql extension for bm25 ranking algorithm


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2150** | {{< badge content="vchord_bm25" link="https://github.com/tensorchord/VectorChord-bm25" >}} | {{< ext "vchord_bm25" >}} | `0.2.2` | {{< category "FTS" >}} | {{< license "AGPL-3.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "vector" >}} {{< ext "vchord" >}} {{< ext "pg_search" >}} {{< ext "pg_bestmatch" >}} {{< ext "vectorscale" >}} {{< ext "zhparser" >}} {{< ext "pg_tokenizer" >}} {{< ext "pgroonga" >}} |

> [!Note] pgrx=0.16.1


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/vchord_bm25" >}} | `0.2.2` | {{< bg "18" "vchord_bm25_18" "green" >}} {{< bg "17" "vchord_bm25_17" "green" >}} {{< bg "16" "vchord_bm25_16" "green" >}} {{< bg "15" "vchord_bm25_15" "green" >}} {{< bg "14" "vchord_bm25_14" "green" >}} {{< bg "13" "vchord_bm25_13" "green" >}} | `vchord_bm25_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/vchord_bm25" >}} | `0.2.2` | {{< bg "18" "postgresql-18-vchord-bm25" "green" >}} {{< bg "17" "postgresql-17-vchord-bm25" "green" >}} {{< bg "16" "postgresql-16-vchord-bm25" "green" >}} {{< bg "15" "postgresql-15-vchord-bm25" "green" >}} {{< bg "14" "postgresql-14-vchord-bm25" "green" >}} {{< bg "13" "postgresql-13-vchord-bm25" "green" >}} | `postgresql-$v-vchord-bm25` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 0.2.2" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 0.2.2" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 0.2.2" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 0.2.2" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    | {{< bg "PIGSTY 0.2.2" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_13 : AVAIL 1" "green" >}} |
|    `el10.aarch64`    | {{< bg "PIGSTY 0.2.2" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "vchord_bm25_13 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-vchord-bm25 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2.1" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-vchord-bm25 : MISS 0" "red" >}}      |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-vchord-bm25 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2.1" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-vchord-bm25 : MISS 0" "red" >}}      |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-vchord-bm25 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-vchord-bm25 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-vchord-bm25 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-vchord-bm25 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-vchord-bm25 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-vchord-bm25 : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-vchord-bm25 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-vchord-bm25 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-vchord-bm25 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-vchord-bm25 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-vchord-bm25 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-vchord-bm25 : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-vchord-bm25 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2.1" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-vchord-bm25 : MISS 0" "red" >}}      |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-vchord-bm25 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2.1" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-vchord-bm25 : MISS 0" "red" >}}      |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-vchord-bm25 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2.1" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-vchord-bm25 : MISS 0" "red" >}}      |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-vchord-bm25 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2.1" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-vchord-bm25 : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_18` | 0.2.2 | `el8.x86_64` | pigsty | 516.0 KiB | [vchord_bm25_18-0.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_18-0.2.2-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_18` | 0.2.2 | `el8.aarch64` | pigsty | 408.7 KiB | [vchord_bm25_18-0.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_18-0.2.2-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_18` | 0.2.2 | `el9.x86_64` | pigsty | 532.0 KiB | [vchord_bm25_18-0.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_18-0.2.2-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_18` | 0.2.2 | `el9.aarch64` | pigsty | 440.5 KiB | [vchord_bm25_18-0.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_18-0.2.2-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_18` | 0.2.2 | `el10.x86_64` | pigsty | 535.1 KiB | [vchord_bm25_18-0.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_18-0.2.2-1PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_18` | 0.2.2 | `el10.aarch64` | pigsty | 446.8 KiB | [vchord_bm25_18-0.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_18-0.2.2-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_17` | 0.2.2 | `el8.x86_64` | pigsty | 515.6 KiB | [vchord_bm25_17-0.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_17-0.2.2-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_17` | 0.2.1 | `el8.x86_64` | pigsty | 416.4 KiB | [vchord_bm25_17-0.2.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_17-0.2.1-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_17` | 0.2.2 | `el8.aarch64` | pigsty | 408.6 KiB | [vchord_bm25_17-0.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_17-0.2.2-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_17` | 0.2.1 | `el8.aarch64` | pigsty | 404.9 KiB | [vchord_bm25_17-0.2.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_17-0.2.1-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_17` | 0.2.2 | `el9.x86_64` | pigsty | 531.8 KiB | [vchord_bm25_17-0.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_17-0.2.2-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_17` | 0.2.1 | `el9.x86_64` | pigsty | 422.4 KiB | [vchord_bm25_17-0.2.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_17-0.2.1-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_17` | 0.2.2 | `el9.aarch64` | pigsty | 440.1 KiB | [vchord_bm25_17-0.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_17-0.2.2-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_17` | 0.2.1 | `el9.aarch64` | pigsty | 437.1 KiB | [vchord_bm25_17-0.2.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_17-0.2.1-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_17` | 0.2.2 | `el10.x86_64` | pigsty | 535.5 KiB | [vchord_bm25_17-0.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_17-0.2.2-1PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_17` | 0.2.2 | `el10.aarch64` | pigsty | 447.0 KiB | [vchord_bm25_17-0.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_17-0.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-vchord-bm25` | 0.2.1 | `d12.x86_64` | pigsty | 333.0 KiB | [postgresql-17-vchord-bm25_0.2.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.2.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-vchord-bm25` | 0.2.1 | `d12.aarch64` | pigsty | 318.4 KiB | [postgresql-17-vchord-bm25_0.2.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.2.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-vchord-bm25` | 0.2.1 | `u22.x86_64` | pigsty | 369.4 KiB | [postgresql-17-vchord-bm25_0.2.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.2.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-vchord-bm25` | 0.2.1 | `u22.aarch64` | pigsty | 378.6 KiB | [postgresql-17-vchord-bm25_0.2.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.2.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-vchord-bm25` | 0.2.1 | `u24.x86_64` | pigsty | 365.5 KiB | [postgresql-17-vchord-bm25_0.2.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.2.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-vchord-bm25` | 0.2.1 | `u24.aarch64` | pigsty | 372.8 KiB | [postgresql-17-vchord-bm25_0.2.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.2.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_16` | 0.2.2 | `el8.x86_64` | pigsty | 515.6 KiB | [vchord_bm25_16-0.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_16-0.2.2-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_16` | 0.2.1 | `el8.x86_64` | pigsty | 416.3 KiB | [vchord_bm25_16-0.2.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_16-0.2.1-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_16` | 0.2.2 | `el8.aarch64` | pigsty | 408.5 KiB | [vchord_bm25_16-0.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_16-0.2.2-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_16` | 0.2.1 | `el8.aarch64` | pigsty | 404.6 KiB | [vchord_bm25_16-0.2.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_16-0.2.1-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_16` | 0.2.2 | `el9.x86_64` | pigsty | 531.8 KiB | [vchord_bm25_16-0.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_16-0.2.2-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_16` | 0.2.1 | `el9.x86_64` | pigsty | 422.2 KiB | [vchord_bm25_16-0.2.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_16-0.2.1-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_16` | 0.2.2 | `el9.aarch64` | pigsty | 440.1 KiB | [vchord_bm25_16-0.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_16-0.2.2-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_16` | 0.2.1 | `el9.aarch64` | pigsty | 437.0 KiB | [vchord_bm25_16-0.2.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_16-0.2.1-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_16` | 0.2.2 | `el10.x86_64` | pigsty | 535.2 KiB | [vchord_bm25_16-0.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_16-0.2.2-1PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_16` | 0.2.2 | `el10.aarch64` | pigsty | 446.6 KiB | [vchord_bm25_16-0.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_16-0.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-vchord-bm25` | 0.2.1 | `d12.x86_64` | pigsty | 333.1 KiB | [postgresql-16-vchord-bm25_0.2.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.2.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-vchord-bm25` | 0.2.1 | `d12.aarch64` | pigsty | 318.2 KiB | [postgresql-16-vchord-bm25_0.2.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.2.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-vchord-bm25` | 0.2.1 | `u22.x86_64` | pigsty | 369.4 KiB | [postgresql-16-vchord-bm25_0.2.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.2.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-vchord-bm25` | 0.2.1 | `u22.aarch64` | pigsty | 378.4 KiB | [postgresql-16-vchord-bm25_0.2.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.2.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-vchord-bm25` | 0.2.1 | `u24.x86_64` | pigsty | 365.6 KiB | [postgresql-16-vchord-bm25_0.2.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.2.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-vchord-bm25` | 0.2.1 | `u24.aarch64` | pigsty | 372.3 KiB | [postgresql-16-vchord-bm25_0.2.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.2.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_15` | 0.2.2 | `el8.x86_64` | pigsty | 518.0 KiB | [vchord_bm25_15-0.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_15-0.2.2-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_15` | 0.2.1 | `el8.x86_64` | pigsty | 417.4 KiB | [vchord_bm25_15-0.2.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_15-0.2.1-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_15` | 0.2.2 | `el8.aarch64` | pigsty | 410.4 KiB | [vchord_bm25_15-0.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_15-0.2.2-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_15` | 0.2.1 | `el8.aarch64` | pigsty | 405.9 KiB | [vchord_bm25_15-0.2.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_15-0.2.1-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_15` | 0.2.2 | `el9.x86_64` | pigsty | 533.0 KiB | [vchord_bm25_15-0.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_15-0.2.2-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_15` | 0.2.1 | `el9.x86_64` | pigsty | 423.3 KiB | [vchord_bm25_15-0.2.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_15-0.2.1-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_15` | 0.2.2 | `el9.aarch64` | pigsty | 442.7 KiB | [vchord_bm25_15-0.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_15-0.2.2-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_15` | 0.2.1 | `el9.aarch64` | pigsty | 437.5 KiB | [vchord_bm25_15-0.2.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_15-0.2.1-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_15` | 0.2.2 | `el10.x86_64` | pigsty | 536.6 KiB | [vchord_bm25_15-0.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_15-0.2.2-1PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_15` | 0.2.2 | `el10.aarch64` | pigsty | 449.1 KiB | [vchord_bm25_15-0.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_15-0.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-vchord-bm25` | 0.2.1 | `d12.x86_64` | pigsty | 334.2 KiB | [postgresql-15-vchord-bm25_0.2.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.2.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-vchord-bm25` | 0.2.1 | `d12.aarch64` | pigsty | 319.4 KiB | [postgresql-15-vchord-bm25_0.2.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.2.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-vchord-bm25` | 0.2.1 | `u22.x86_64` | pigsty | 370.8 KiB | [postgresql-15-vchord-bm25_0.2.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.2.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-vchord-bm25` | 0.2.1 | `u22.aarch64` | pigsty | 379.7 KiB | [postgresql-15-vchord-bm25_0.2.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.2.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-vchord-bm25` | 0.2.1 | `u24.x86_64` | pigsty | 367.2 KiB | [postgresql-15-vchord-bm25_0.2.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.2.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-vchord-bm25` | 0.2.1 | `u24.aarch64` | pigsty | 373.8 KiB | [postgresql-15-vchord-bm25_0.2.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.2.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_14` | 0.2.2 | `el8.x86_64` | pigsty | 517.5 KiB | [vchord_bm25_14-0.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_14-0.2.2-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_14` | 0.2.1 | `el8.x86_64` | pigsty | 417.4 KiB | [vchord_bm25_14-0.2.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_14-0.2.1-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_14` | 0.2.2 | `el8.aarch64` | pigsty | 410.6 KiB | [vchord_bm25_14-0.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_14-0.2.2-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_14` | 0.2.1 | `el8.aarch64` | pigsty | 405.8 KiB | [vchord_bm25_14-0.2.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_14-0.2.1-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_14` | 0.2.2 | `el9.x86_64` | pigsty | 533.0 KiB | [vchord_bm25_14-0.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_14-0.2.2-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_14` | 0.2.1 | `el9.x86_64` | pigsty | 423.4 KiB | [vchord_bm25_14-0.2.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_14-0.2.1-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_14` | 0.2.2 | `el9.aarch64` | pigsty | 442.9 KiB | [vchord_bm25_14-0.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_14-0.2.2-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_14` | 0.2.1 | `el9.aarch64` | pigsty | 437.3 KiB | [vchord_bm25_14-0.2.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_14-0.2.1-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_14` | 0.2.2 | `el10.x86_64` | pigsty | 536.8 KiB | [vchord_bm25_14-0.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_14-0.2.2-1PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_14` | 0.2.2 | `el10.aarch64` | pigsty | 449.3 KiB | [vchord_bm25_14-0.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_14-0.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-vchord-bm25` | 0.2.1 | `d12.x86_64` | pigsty | 333.9 KiB | [postgresql-14-vchord-bm25_0.2.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.2.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-vchord-bm25` | 0.2.1 | `d12.aarch64` | pigsty | 319.3 KiB | [postgresql-14-vchord-bm25_0.2.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.2.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-vchord-bm25` | 0.2.1 | `u22.x86_64` | pigsty | 370.5 KiB | [postgresql-14-vchord-bm25_0.2.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.2.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-vchord-bm25` | 0.2.1 | `u22.aarch64` | pigsty | 379.3 KiB | [postgresql-14-vchord-bm25_0.2.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.2.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-vchord-bm25` | 0.2.1 | `u24.x86_64` | pigsty | 367.6 KiB | [postgresql-14-vchord-bm25_0.2.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.2.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-vchord-bm25` | 0.2.1 | `u24.aarch64` | pigsty | 373.7 KiB | [postgresql-14-vchord-bm25_0.2.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.2.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_13` | 0.2.2 | `el8.x86_64` | pigsty | 517.2 KiB | [vchord_bm25_13-0.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_13-0.2.2-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_13` | 0.2.2 | `el8.aarch64` | pigsty | 410.4 KiB | [vchord_bm25_13-0.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_13-0.2.2-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_13` | 0.2.2 | `el9.x86_64` | pigsty | 533.3 KiB | [vchord_bm25_13-0.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_13-0.2.2-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_13` | 0.2.2 | `el9.aarch64` | pigsty | 442.5 KiB | [vchord_bm25_13-0.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_13-0.2.2-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_13` | 0.2.2 | `el10.x86_64` | pigsty | 536.6 KiB | [vchord_bm25_13-0.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_13-0.2.2-1PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_13` | 0.2.2 | `el10.aarch64` | pigsty | 448.5 KiB | [vchord_bm25_13-0.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_13-0.2.2-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tensorchord/VectorChord-bm25" title="Repository" icon="github" subtitle="github.com/tensorchord/VectorChord-bm25" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="VectorChord-bm25-0.2.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build get vchord_bm25; # get vchord_bm25 source code
pig build dep vchord_bm25; # install build dependencies
pig build pkg vchord_bm25; # build extension rpm or deb
pig build ext vchord_bm25; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install vchord_bm25; # install by extension name, for the current active PG version
pig ext install vchord_bm25; # install via package alias, for the active PG version
pig ext install vchord_bm25 -v 18;   # install for PG 18
pig ext install vchord_bm25 -v 17;   # install for PG 17
pig ext install vchord_bm25 -v 16;   # install for PG 16
pig ext install vchord_bm25 -v 15;   # install for PG 15
pig ext install vchord_bm25 -v 14;   # install for PG 14
pig ext install vchord_bm25 -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION vchord_bm25 CASCADE SCHEMA bm25_catalog;
```

