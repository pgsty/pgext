---
title: "pgroonga"
linkTitle: "pgroonga"
description: "Use Groonga as index, fast full text search platform for all languages!"
weight: 2110
categories: ["FTS"]
width: full
---

Use Groonga as index, fast full text search platform for all languages!


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2110** | {{< badge content="pgroonga" link="https://github.com/pgroonga/pgroonga" >}} | {{< ext "pgroonga" >}} | `4.0.4` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dtr" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pg_bigm" >}} {{< ext "zhparser" >}} {{< ext "pg_bestmatch" >}} {{< ext "pg_tokenizer" >}} {{< ext "pg_trgm" >}} {{< ext "rum" >}} {{< ext "vchord_bm25" >}} |
|    **Siblings**   | {{< ext "pgroonga_database" >}} |

> [!Note] require xxHash vendor repo to build


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgroonga" >}} | `4.0.4` | {{< bg "18" "pgroonga_18*" "red" >}} {{< bg "17" "pgroonga_17*" "green" >}} {{< bg "16" "pgroonga_16*" "green" >}} {{< bg "15" "pgroonga_15*" "green" >}} {{< bg "14" "pgroonga_14*" "green" >}} {{< bg "13" "pgroonga_13*" "green" >}} | `pgroonga_$v*` | `groonga-libs` |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgroonga" >}} | `4.0.4` | {{< bg "18" "postgresql-18-pgroonga" "green" >}} {{< bg "17" "postgresql-17-pgroonga" "green" >}} {{< bg "16" "postgresql-16-pgroonga" "green" >}} {{< bg "15" "postgresql-15-pgroonga" "green" >}} {{< bg "14" "postgresql-14-pgroonga" "green" >}} {{< bg "13" "postgresql-13-pgroonga" "green" >}} | `postgresql-$v-pgroonga` | `libgroonga0` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pgroonga_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pgroonga_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pgroonga_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pgroonga_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgroonga_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    | {{< bg "PIGSTY 4.0.4" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_13 : AVAIL 1" "green" >}} |
|    `el10.aarch64`    | {{< bg "PIGSTY 4.0.4" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_13 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pgroonga : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pgroonga : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pgroonga : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgroonga : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pgroonga : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pgroonga : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pgroonga : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pgroonga : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pgroonga : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgroonga_18` | 4.0.4 | `el10.x86_64` | pigsty | 347.3 KiB | [pgroonga_18-4.0.4-1.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgroonga_18-4.0.4-1.el10.x86_64.rpm) |
| `pgroonga_18` | 4.0.4 | `el10.aarch64` | pigsty | 339.3 KiB | [pgroonga_18-4.0.4-1.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgroonga_18-4.0.4-1.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgroonga_17` | 4.0.0 | `el8.x86_64` | pigsty | 342.0 KiB | [pgroonga_17-4.0.0-1.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_17-4.0.0-1.el8.x86_64.rpm) |
| `pgroonga_17` | 4.0.0 | `el8.aarch64` | pigsty | 330.5 KiB | [pgroonga_17-4.0.0-1.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_17-4.0.0-1.el8.aarch64.rpm) |
| `pgroonga_17` | 4.0.0 | `el9.x86_64` | pigsty | 328.0 KiB | [pgroonga_17-4.0.0-1.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_17-4.0.0-1.el9.x86_64.rpm) |
| `pgroonga_17` | 4.0.0 | `el9.aarch64` | pigsty | 320.1 KiB | [pgroonga_17-4.0.0-1.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_17-4.0.0-1.el9.aarch64.rpm) |
| `pgroonga_17` | 4.0.4 | `el10.x86_64` | pigsty | 347.0 KiB | [pgroonga_17-4.0.4-1.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgroonga_17-4.0.4-1.el10.x86_64.rpm) |
| `pgroonga_17` | 4.0.4 | `el10.aarch64` | pigsty | 339.2 KiB | [pgroonga_17-4.0.4-1.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgroonga_17-4.0.4-1.el10.aarch64.rpm) |
| `postgresql-17-pgroonga` | 4.0.0 | `d12.x86_64` | pigsty | 697.4 KiB | [postgresql-17-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgroonga` | 4.0.0 | `d12.aarch64` | pigsty | 693.0 KiB | [postgresql-17-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgroonga` | 4.0.0 | `u22.x86_64` | pigsty | 736.2 KiB | [postgresql-17-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgroonga` | 4.0.0 | `u22.aarch64` | pigsty | 739.9 KiB | [postgresql-17-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgroonga` | 4.0.0 | `u24.x86_64` | pigsty | 649.9 KiB | [postgresql-17-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgroonga` | 4.0.0 | `u24.aarch64` | pigsty | 653.1 KiB | [postgresql-17-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgroonga_16` | 4.0.0 | `el8.x86_64` | pigsty | 339.7 KiB | [pgroonga_16-4.0.0-1.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_16-4.0.0-1.el8.x86_64.rpm) |
| `pgroonga_16` | 4.0.0 | `el8.aarch64` | pigsty | 328.3 KiB | [pgroonga_16-4.0.0-1.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_16-4.0.0-1.el8.aarch64.rpm) |
| `pgroonga_16` | 4.0.0 | `el9.x86_64` | pigsty | 325.9 KiB | [pgroonga_16-4.0.0-1.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_16-4.0.0-1.el9.x86_64.rpm) |
| `pgroonga_16` | 4.0.0 | `el9.aarch64` | pigsty | 317.5 KiB | [pgroonga_16-4.0.0-1.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_16-4.0.0-1.el9.aarch64.rpm) |
| `pgroonga_16` | 4.0.4 | `el10.x86_64` | pigsty | 344.8 KiB | [pgroonga_16-4.0.4-1.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgroonga_16-4.0.4-1.el10.x86_64.rpm) |
| `pgroonga_16` | 4.0.4 | `el10.aarch64` | pigsty | 337.1 KiB | [pgroonga_16-4.0.4-1.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgroonga_16-4.0.4-1.el10.aarch64.rpm) |
| `postgresql-16-pgroonga` | 4.0.0 | `d12.x86_64` | pigsty | 684.9 KiB | [postgresql-16-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgroonga` | 4.0.0 | `d12.aarch64` | pigsty | 681.1 KiB | [postgresql-16-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgroonga` | 4.0.0 | `u22.x86_64` | pigsty | 723.3 KiB | [postgresql-16-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgroonga` | 4.0.0 | `u22.aarch64` | pigsty | 727.5 KiB | [postgresql-16-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgroonga` | 4.0.0 | `u24.x86_64` | pigsty | 642.4 KiB | [postgresql-16-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgroonga` | 4.0.0 | `u24.aarch64` | pigsty | 646.1 KiB | [postgresql-16-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgroonga_15` | 4.0.0 | `el8.x86_64` | pigsty | 341.9 KiB | [pgroonga_15-4.0.0-1.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_15-4.0.0-1.el8.x86_64.rpm) |
| `pgroonga_15` | 4.0.0 | `el8.aarch64` | pigsty | 330.8 KiB | [pgroonga_15-4.0.0-1.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_15-4.0.0-1.el8.aarch64.rpm) |
| `pgroonga_15` | 4.0.0 | `el9.x86_64` | pigsty | 327.9 KiB | [pgroonga_15-4.0.0-1.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_15-4.0.0-1.el9.x86_64.rpm) |
| `pgroonga_15` | 4.0.0 | `el9.aarch64` | pigsty | 321.7 KiB | [pgroonga_15-4.0.0-1.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_15-4.0.0-1.el9.aarch64.rpm) |
| `pgroonga_15` | 4.0.4 | `el10.x86_64` | pigsty | 349.2 KiB | [pgroonga_15-4.0.4-1.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgroonga_15-4.0.4-1.el10.x86_64.rpm) |
| `pgroonga_15` | 4.0.4 | `el10.aarch64` | pigsty | 339.4 KiB | [pgroonga_15-4.0.4-1.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgroonga_15-4.0.4-1.el10.aarch64.rpm) |
| `postgresql-15-pgroonga` | 4.0.0 | `d12.x86_64` | pigsty | 687.5 KiB | [postgresql-15-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgroonga` | 4.0.0 | `d12.aarch64` | pigsty | 683.7 KiB | [postgresql-15-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgroonga` | 4.0.0 | `u22.x86_64` | pigsty | 729.8 KiB | [postgresql-15-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgroonga` | 4.0.0 | `u22.aarch64` | pigsty | 738.1 KiB | [postgresql-15-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgroonga` | 4.0.0 | `u24.x86_64` | pigsty | 651.6 KiB | [postgresql-15-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgroonga` | 4.0.0 | `u24.aarch64` | pigsty | 654.5 KiB | [postgresql-15-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgroonga_14` | 4.0.0 | `el8.x86_64` | pigsty | 322.8 KiB | [pgroonga_14-4.0.0-1.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_14-4.0.0-1.el8.x86_64.rpm) |
| `pgroonga_14` | 4.0.0 | `el8.aarch64` | pigsty | 314.4 KiB | [pgroonga_14-4.0.0-1.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_14-4.0.0-1.el8.aarch64.rpm) |
| `pgroonga_14` | 4.0.0 | `el9.x86_64` | pigsty | 310.4 KiB | [pgroonga_14-4.0.0-1.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_14-4.0.0-1.el9.x86_64.rpm) |
| `pgroonga_14` | 4.0.0 | `el9.aarch64` | pigsty | 305.4 KiB | [pgroonga_14-4.0.0-1.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_14-4.0.0-1.el9.aarch64.rpm) |
| `pgroonga_14` | 4.0.4 | `el10.x86_64` | pigsty | 330.9 KiB | [pgroonga_14-4.0.4-1.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgroonga_14-4.0.4-1.el10.x86_64.rpm) |
| `pgroonga_14` | 4.0.4 | `el10.aarch64` | pigsty | 322.9 KiB | [pgroonga_14-4.0.4-1.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgroonga_14-4.0.4-1.el10.aarch64.rpm) |
| `postgresql-14-pgroonga` | 4.0.0 | `d12.x86_64` | pigsty | 630.7 KiB | [postgresql-14-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgroonga` | 4.0.0 | `d12.aarch64` | pigsty | 627.3 KiB | [postgresql-14-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgroonga` | 4.0.0 | `u22.x86_64` | pigsty | 670.2 KiB | [postgresql-14-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgroonga` | 4.0.0 | `u22.aarch64` | pigsty | 678.7 KiB | [postgresql-14-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgroonga` | 4.0.0 | `u24.x86_64` | pigsty | 595.1 KiB | [postgresql-14-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgroonga` | 4.0.0 | `u24.aarch64` | pigsty | 600.1 KiB | [postgresql-14-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgroonga_13` | 4.0.0 | `el8.x86_64` | pigsty | 321.9 KiB | [pgroonga_13-4.0.0-1.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_13-4.0.0-1.el8.x86_64.rpm) |
| `pgroonga_13` | 4.0.0 | `el8.aarch64` | pigsty | 314.3 KiB | [pgroonga_13-4.0.0-1.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_13-4.0.0-1.el8.aarch64.rpm) |
| `pgroonga_13` | 4.0.0 | `el9.x86_64` | pigsty | 310.4 KiB | [pgroonga_13-4.0.0-1.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_13-4.0.0-1.el9.x86_64.rpm) |
| `pgroonga_13` | 4.0.0 | `el9.aarch64` | pigsty | 305.2 KiB | [pgroonga_13-4.0.0-1.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_13-4.0.0-1.el9.aarch64.rpm) |
| `pgroonga_13` | 4.0.4 | `el10.x86_64` | pigsty | 330.2 KiB | [pgroonga_13-4.0.4-1.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgroonga_13-4.0.4-1.el10.x86_64.rpm) |
| `pgroonga_13` | 4.0.4 | `el10.aarch64` | pigsty | 323.0 KiB | [pgroonga_13-4.0.4-1.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgroonga_13-4.0.4-1.el10.aarch64.rpm) |
| `postgresql-13-pgroonga` | 4.0.0 | `d12.x86_64` | pigsty | 738.3 KiB | [postgresql-13-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgroonga` | 4.0.0 | `d12.aarch64` | pigsty | 730.7 KiB | [postgresql-13-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgroonga` | 4.0.0 | `u22.x86_64` | pigsty | 786.0 KiB | [postgresql-13-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgroonga` | 4.0.0 | `u22.aarch64` | pigsty | 786.5 KiB | [postgresql-13-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgroonga` | 4.0.0 | `u24.x86_64` | pigsty | 695.4 KiB | [postgresql-13-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgroonga` | 4.0.0 | `u24.aarch64` | pigsty | 695.5 KiB | [postgresql-13-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgroonga/pgroonga" title="Repository" icon="github" subtitle="github.com/pgroonga/pgroonga" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgroonga-4.0.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgroonga; # get pgroonga source code
pig build dep pgroonga; # install build dependencies
pig build pkg pgroonga; # build extension rpm or deb
pig build ext pgroonga; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgroonga; # install by extension name, for the current active PG version
pig ext install pgroonga; # install via package alias, for the active PG version
pig ext install pgroonga -v 18;   # install for PG 18
pig ext install pgroonga -v 17;   # install for PG 17
pig ext install pgroonga -v 16;   # install for PG 16
pig ext install pgroonga -v 15;   # install for PG 15
pig ext install pgroonga -v 14;   # install for PG 14
pig ext install pgroonga -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgroonga;
```




--------

## Usage

- https://pgroonga.github.io/
- [News](https://pgroonga.github.io/news/): It lists release information.
- [Overview](https://pgroonga.github.io/overview/): It describes about PGroonga.
- [Install](https://pgroonga.github.io/install/): It describes how to install PGroonga.
- [Upgrade](https://pgroonga.github.io/upgrade/): It describes how to upgrade PGroonga.
- [Uninstall](https://pgroonga.github.io/uninstall/): It describes how to uninstall PGroonga.
- [Tutorial](https://pgroonga.github.io/tutorial/): It describes how to use PGroonga step by step.
- [FAQ](https://pgroonga.github.io/faq/): Frequently asked questions.
- [How to](https://pgroonga.github.io/how-to/): It describes about useful information for specific situations.
- [Reference](https://pgroonga.github.io/reference/): It describes details for each features such as options, functions and operators.
- [Troubleshooting](https://pgroonga.github.io/troubleshooting/): It describes how to fix troubles.
- [Community](https://pgroonga.github.io/community/): It introduces about PGroonga community.
- [Users](https://pgroonga.github.io/users/): It lists PGroonga users.
- [Development](https://pgroonga.github.io/development/): It describes how to develop PGroonga.

Here's a quick [tutorial](https://pgroonga.github.io/tutorial/) about how to use PGroonga:

```sql
CREATE EXTENSION IF NOT EXISTS pgroonga;

CREATE TABLE memos
(
    id      integer,
    content text
);

CREATE INDEX pgroonga_content_index ON memos USING pgroonga (content);

INSERT INTO memos VALUES (1, 'PostgreSQL is a relational database management system.');
INSERT INTO memos VALUES (2, 'Groonga is a fast full text search engine that supports all languages.');
INSERT INTO memos VALUES (3, 'PGroonga is a PostgreSQL extension that uses Groonga as index.');
INSERT INTO memos VALUES (4, 'There is groonga command.');

SET enable_seqscan = off;

-- now let's query pgroonga

SELECT * FROM memos WHERE content &@ 'engine';
--  id |                                content                                 
-- ----+------------------------------------------------------------------------
--   2 | Groonga is a fast full text search engine that supports all languages.
-- (1 row)

SELECT * FROM memos WHERE content &@~ 'PGroonga OR PostgreSQL';
--  id |                            content                             
-- ----+----------------------------------------------------------------
--   3 | PGroonga is a PostgreSQL extension that uses Groonga as index.
--   1 | PostgreSQL is a relational database management system.
-- (2 rows)

SELECT * FROM memos WHERE content LIKE '%engine%';
--  id |                                content                                 
-- ----+------------------------------------------------------------------------
--   2 | Groonga is a fast full text search engine that supports all languages.
-- (1 row)
```