---
title: "zhparser"
linkTitle: "zhparser"
description: "a parser for full-text search of Chinese"
weight: 2130
categories: ["FTS"]
width: full
---

a parser for full-text search of Chinese


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2130** | {{< badge content="zhparser" link="https://github.com/amutu/zhparser" >}} | {{< ext "zhparser" >}} | `2.3` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_trgm" >}} {{< ext "rum" >}} {{< ext "pg_search" >}} {{< ext "pgroonga" >}} {{< ext "pgroonga_database" >}} {{< ext "pg_bigm" >}} {{< ext "pg_tokenizer" >}} {{< ext "vchord_bm25" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/zhparser" >}} | `2.3` | {{< bg "18" "zhparser_18*" "green" >}} {{< bg "17" "zhparser_17*" "green" >}} {{< bg "16" "zhparser_16*" "green" >}} {{< bg "15" "zhparser_15*" "green" >}} {{< bg "14" "zhparser_14*" "green" >}} {{< bg "13" "zhparser_13*" "green" >}} | `zhparser_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/zhparser" >}} | `2.3` | {{< bg "18" "postgresql-18-zhparser" "green" >}} {{< bg "17" "postgresql-17-zhparser" "green" >}} {{< bg "16" "postgresql-16-zhparser" "green" >}} {{< bg "15" "postgresql-15-zhparser" "green" >}} {{< bg "14" "postgresql-14-zhparser" "green" >}} {{< bg "13" "postgresql-13-zhparser" "green" >}} | `postgresql-$v-zhparser` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 2.3" "zhparser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 2.3" "zhparser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 2.3" "zhparser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 2.3" "zhparser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    |      {{< bg "MISS" "zhparser_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "zhparser_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "zhparser_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "zhparser_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "zhparser_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "zhparser_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "zhparser_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "zhparser_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "zhparser_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "zhparser_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "zhparser_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "zhparser_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-zhparser : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-13-zhparser : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-zhparser : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-13-zhparser : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-zhparser : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-zhparser : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-zhparser : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-zhparser : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-zhparser : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-zhparser : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-zhparser : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-zhparser : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-zhparser : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-zhparser : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-zhparser : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-zhparser : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-zhparser : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-13-zhparser : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-zhparser : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-13-zhparser : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-zhparser : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-13-zhparser : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-zhparser : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-13-zhparser : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `zhparser_18` | 2.3 | `el8.x86_64` | pigsty | 4.7 MiB | [zhparser_18-2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/zhparser_18-2.3-1PIGSTY.el8.x86_64.rpm) |
| `zhparser_18` | 2.3 | `el8.aarch64` | pigsty | 4.7 MiB | [zhparser_18-2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/zhparser_18-2.3-1PIGSTY.el8.aarch64.rpm) |
| `zhparser_18` | 2.3 | `el9.x86_64` | pigsty | 4.3 MiB | [zhparser_18-2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/zhparser_18-2.3-1PIGSTY.el9.x86_64.rpm) |
| `zhparser_18` | 2.3 | `el9.aarch64` | pigsty | 4.3 MiB | [zhparser_18-2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/zhparser_18-2.3-1PIGSTY.el9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `zhparser_17` | 2.3 | `el8.x86_64` | pigsty | 4.7 MiB | [zhparser_17-2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/zhparser_17-2.3-1PIGSTY.el8.x86_64.rpm) |
| `zhparser_17` | 2.3 | `el8.aarch64` | pigsty | 4.7 MiB | [zhparser_17-2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/zhparser_17-2.3-1PIGSTY.el8.aarch64.rpm) |
| `zhparser_17` | 2.3 | `el9.x86_64` | pigsty | 4.3 MiB | [zhparser_17-2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/zhparser_17-2.3-1PIGSTY.el9.x86_64.rpm) |
| `zhparser_17` | 2.3 | `el9.aarch64` | pigsty | 4.3 MiB | [zhparser_17-2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/zhparser_17-2.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-zhparser` | 2.3 | `d12.x86_64` | pigsty | 4.0 MiB | [postgresql-17-zhparser_2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-17-zhparser_2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-zhparser` | 2.3 | `d12.aarch64` | pigsty | 4.0 MiB | [postgresql-17-zhparser_2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-17-zhparser_2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-zhparser` | 2.3 | `u22.x86_64` | pigsty | 4.3 MiB | [postgresql-17-zhparser_2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-17-zhparser_2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-zhparser` | 2.3 | `u22.aarch64` | pigsty | 4.3 MiB | [postgresql-17-zhparser_2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-17-zhparser_2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-zhparser` | 2.3 | `u24.x86_64` | pigsty | 4.3 MiB | [postgresql-17-zhparser_2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-17-zhparser_2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-zhparser` | 2.3 | `u24.aarch64` | pigsty | 4.3 MiB | [postgresql-17-zhparser_2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-17-zhparser_2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `zhparser_16` | 2.3 | `el8.x86_64` | pigsty | 4.7 MiB | [zhparser_16-2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/zhparser_16-2.3-1PIGSTY.el8.x86_64.rpm) |
| `zhparser_16` | 2.3 | `el8.aarch64` | pigsty | 4.7 MiB | [zhparser_16-2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/zhparser_16-2.3-1PIGSTY.el8.aarch64.rpm) |
| `zhparser_16` | 2.3 | `el9.x86_64` | pigsty | 4.3 MiB | [zhparser_16-2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/zhparser_16-2.3-1PIGSTY.el9.x86_64.rpm) |
| `zhparser_16` | 2.3 | `el9.aarch64` | pigsty | 4.3 MiB | [zhparser_16-2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/zhparser_16-2.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-zhparser` | 2.3 | `d12.x86_64` | pigsty | 4.1 MiB | [postgresql-16-zhparser_2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-16-zhparser_2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-zhparser` | 2.3 | `d12.aarch64` | pigsty | 4.0 MiB | [postgresql-16-zhparser_2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-16-zhparser_2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-zhparser` | 2.3 | `u22.x86_64` | pigsty | 4.3 MiB | [postgresql-16-zhparser_2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-16-zhparser_2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-zhparser` | 2.3 | `u22.aarch64` | pigsty | 4.3 MiB | [postgresql-16-zhparser_2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-16-zhparser_2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-zhparser` | 2.3 | `u24.x86_64` | pigsty | 4.3 MiB | [postgresql-16-zhparser_2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-16-zhparser_2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-zhparser` | 2.3 | `u24.aarch64` | pigsty | 4.3 MiB | [postgresql-16-zhparser_2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-16-zhparser_2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `zhparser_15` | 2.3 | `el8.x86_64` | pigsty | 4.7 MiB | [zhparser_15-2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/zhparser_15-2.3-1PIGSTY.el8.x86_64.rpm) |
| `zhparser_15` | 2.3 | `el8.aarch64` | pigsty | 4.7 MiB | [zhparser_15-2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/zhparser_15-2.3-1PIGSTY.el8.aarch64.rpm) |
| `zhparser_15` | 2.3 | `el9.x86_64` | pigsty | 4.3 MiB | [zhparser_15-2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/zhparser_15-2.3-1PIGSTY.el9.x86_64.rpm) |
| `zhparser_15` | 2.3 | `el9.aarch64` | pigsty | 4.3 MiB | [zhparser_15-2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/zhparser_15-2.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-zhparser` | 2.3 | `d12.x86_64` | pigsty | 4.0 MiB | [postgresql-15-zhparser_2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-15-zhparser_2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-zhparser` | 2.3 | `d12.aarch64` | pigsty | 4.1 MiB | [postgresql-15-zhparser_2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-15-zhparser_2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-zhparser` | 2.3 | `u22.x86_64` | pigsty | 4.3 MiB | [postgresql-15-zhparser_2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-15-zhparser_2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-zhparser` | 2.3 | `u22.aarch64` | pigsty | 4.3 MiB | [postgresql-15-zhparser_2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-15-zhparser_2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-zhparser` | 2.3 | `u24.x86_64` | pigsty | 4.3 MiB | [postgresql-15-zhparser_2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-15-zhparser_2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-zhparser` | 2.3 | `u24.aarch64` | pigsty | 4.3 MiB | [postgresql-15-zhparser_2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-15-zhparser_2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `zhparser_14` | 2.3 | `el8.x86_64` | pigsty | 4.7 MiB | [zhparser_14-2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/zhparser_14-2.3-1PIGSTY.el8.x86_64.rpm) |
| `zhparser_14` | 2.3 | `el8.aarch64` | pigsty | 4.7 MiB | [zhparser_14-2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/zhparser_14-2.3-1PIGSTY.el8.aarch64.rpm) |
| `zhparser_14` | 2.3 | `el9.x86_64` | pigsty | 4.3 MiB | [zhparser_14-2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/zhparser_14-2.3-1PIGSTY.el9.x86_64.rpm) |
| `zhparser_14` | 2.3 | `el9.aarch64` | pigsty | 4.3 MiB | [zhparser_14-2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/zhparser_14-2.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-zhparser` | 2.3 | `d12.x86_64` | pigsty | 4.0 MiB | [postgresql-14-zhparser_2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-14-zhparser_2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-zhparser` | 2.3 | `d12.aarch64` | pigsty | 4.0 MiB | [postgresql-14-zhparser_2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-14-zhparser_2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-zhparser` | 2.3 | `u22.x86_64` | pigsty | 4.3 MiB | [postgresql-14-zhparser_2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-14-zhparser_2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-zhparser` | 2.3 | `u22.aarch64` | pigsty | 4.3 MiB | [postgresql-14-zhparser_2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-14-zhparser_2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-zhparser` | 2.3 | `u24.x86_64` | pigsty | 4.3 MiB | [postgresql-14-zhparser_2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-14-zhparser_2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-zhparser` | 2.3 | `u24.aarch64` | pigsty | 4.3 MiB | [postgresql-14-zhparser_2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-14-zhparser_2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `zhparser_13` | 2.3 | `el8.x86_64` | pigsty | 4.7 MiB | [zhparser_13-2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/zhparser_13-2.3-1PIGSTY.el8.x86_64.rpm) |
| `zhparser_13` | 2.3 | `el8.aarch64` | pigsty | 4.7 MiB | [zhparser_13-2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/zhparser_13-2.3-1PIGSTY.el8.aarch64.rpm) |
| `zhparser_13` | 2.3 | `el9.x86_64` | pigsty | 4.3 MiB | [zhparser_13-2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/zhparser_13-2.3-1PIGSTY.el9.x86_64.rpm) |
| `zhparser_13` | 2.3 | `el9.aarch64` | pigsty | 4.3 MiB | [zhparser_13-2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/zhparser_13-2.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-zhparser` | 2.3 | `d12.x86_64` | pigsty | 4.0 MiB | [postgresql-13-zhparser_2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-13-zhparser_2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-zhparser` | 2.3 | `d12.aarch64` | pigsty | 4.1 MiB | [postgresql-13-zhparser_2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-13-zhparser_2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-zhparser` | 2.3 | `u22.x86_64` | pigsty | 4.3 MiB | [postgresql-13-zhparser_2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-13-zhparser_2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-zhparser` | 2.3 | `u22.aarch64` | pigsty | 4.3 MiB | [postgresql-13-zhparser_2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-13-zhparser_2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-zhparser` | 2.3 | `u24.x86_64` | pigsty | 4.3 MiB | [postgresql-13-zhparser_2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-13-zhparser_2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-zhparser` | 2.3 | `u24.aarch64` | pigsty | 4.3 MiB | [postgresql-13-zhparser_2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-13-zhparser_2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/amutu/zhparser" title="Repository" icon="github" subtitle="github.com/amutu/zhparser" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="zhparser-2.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get zhparser; # get zhparser source code
pig build dep zhparser; # install build dependencies
pig build pkg zhparser; # build extension rpm or deb
pig build ext zhparser; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install zhparser; # install by extension name, for the current active PG version
pig ext install zhparser; # install via package alias, for the active PG version
pig ext install zhparser -v 18;   # install for PG 18
pig ext install zhparser -v 17;   # install for PG 17
pig ext install zhparser -v 16;   # install for PG 16
pig ext install zhparser -v 15;   # install for PG 15
pig ext install zhparser -v 14;   # install for PG 14
pig ext install zhparser -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION zhparser;
```

