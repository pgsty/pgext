---
title: "smlar"
linkTitle: "smlar"
description: "Effective similarity search"
weight: 1850
categories: ["RAG"]
width: full
---

Effective similarity search


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1850** | {{< badge content="smlar" link="https://github.com/jirutka/smlar" >}} | {{< ext "smlar" >}} | `1.0` | {{< category "RAG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_similarity" >}} {{< ext "fuzzystrmatch" >}} {{< ext "pg_trgm" >}} {{< ext "intarray" >}} {{< ext "vector" >}} {{< ext "pg_bigm" >}} {{< ext "unaccent" >}} {{< ext "vchord" >}} |

> [!Note] fix math.abs, gist pointer problem, pg18 break


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/smlar" >}} | `1.0` | {{< bg "18" "smlar_18*" "red" >}} {{< bg "17" "smlar_17*" "green" >}} {{< bg "16" "smlar_16*" "green" >}} {{< bg "15" "smlar_15*" "green" >}} {{< bg "14" "smlar_14*" "green" >}} {{< bg "13" "smlar_13*" "green" >}} | `smlar_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/smlar" >}} | `1.0` | {{< bg "18" "postgresql-18-smlar" "red" >}} {{< bg "17" "postgresql-17-smlar" "green" >}} {{< bg "16" "postgresql-16-smlar" "green" >}} {{< bg "15" "postgresql-15-smlar" "green" >}} {{< bg "14" "postgresql-14-smlar" "green" >}} {{< bg "13" "postgresql-13-smlar" "green" >}} | `postgresql-$v-smlar` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "smlar_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "smlar_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "smlar_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "smlar_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "smlar_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "smlar_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "smlar_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "smlar_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    |      {{< bg "MISS" "smlar_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "smlar_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "smlar_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "smlar_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "smlar_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "smlar_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "smlar_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "smlar_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "smlar_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "smlar_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "smlar_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "smlar_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-smlar : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-smlar : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-smlar : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-smlar : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-smlar : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-smlar : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-smlar : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-smlar : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-smlar : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-smlar : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-smlar : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-smlar : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-smlar : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-smlar : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-smlar : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-smlar : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-smlar : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-smlar : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-smlar : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-smlar : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-smlar : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-smlar : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-smlar : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-smlar : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `smlar_17` | 1.0 | `el8.x86_64` | pigsty | 32.5 KiB | [smlar_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/smlar_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `smlar_17` | 1.0 | `el8.aarch64` | pigsty | 30.5 KiB | [smlar_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/smlar_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `smlar_17` | 1.0 | `el9.x86_64` | pigsty | 33.0 KiB | [smlar_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/smlar_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `smlar_17` | 1.0 | `el9.aarch64` | pigsty | 31.8 KiB | [smlar_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/smlar_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-smlar` | 1.0 | `d12.x86_64` | pigsty | 81.8 KiB | [postgresql-17-smlar_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-17-smlar_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-smlar` | 1.0 | `d12.aarch64` | pigsty | 79.9 KiB | [postgresql-17-smlar_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-17-smlar_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-smlar` | 1.0 | `u22.x86_64` | pigsty | 85.8 KiB | [postgresql-17-smlar_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-17-smlar_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-smlar` | 1.0 | `u22.aarch64` | pigsty | 84.4 KiB | [postgresql-17-smlar_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-17-smlar_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-smlar` | 1.0 | `u24.x86_64` | pigsty | 76.6 KiB | [postgresql-17-smlar_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-17-smlar_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-smlar` | 1.0 | `u24.aarch64` | pigsty | 76.1 KiB | [postgresql-17-smlar_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-17-smlar_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `smlar_16` | 1.0 | `el8.x86_64` | pigsty | 32.5 KiB | [smlar_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/smlar_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `smlar_16` | 1.0 | `el8.aarch64` | pigsty | 30.5 KiB | [smlar_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/smlar_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `smlar_16` | 1.0 | `el9.x86_64` | pigsty | 33.0 KiB | [smlar_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/smlar_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `smlar_16` | 1.0 | `el9.aarch64` | pigsty | 31.8 KiB | [smlar_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/smlar_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-smlar` | 1.0 | `d12.x86_64` | pigsty | 81.6 KiB | [postgresql-16-smlar_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-16-smlar_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-smlar` | 1.0 | `d12.aarch64` | pigsty | 79.7 KiB | [postgresql-16-smlar_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-16-smlar_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-smlar` | 1.0 | `u22.x86_64` | pigsty | 85.6 KiB | [postgresql-16-smlar_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-16-smlar_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-smlar` | 1.0 | `u22.aarch64` | pigsty | 84.2 KiB | [postgresql-16-smlar_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-16-smlar_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-smlar` | 1.0 | `u24.x86_64` | pigsty | 76.6 KiB | [postgresql-16-smlar_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-16-smlar_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-smlar` | 1.0 | `u24.aarch64` | pigsty | 76.1 KiB | [postgresql-16-smlar_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-16-smlar_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `smlar_15` | 1.0 | `el8.x86_64` | pigsty | 32.9 KiB | [smlar_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/smlar_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `smlar_15` | 1.0 | `el8.aarch64` | pigsty | 30.9 KiB | [smlar_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/smlar_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `smlar_15` | 1.0 | `el9.x86_64` | pigsty | 33.6 KiB | [smlar_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/smlar_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `smlar_15` | 1.0 | `el9.aarch64` | pigsty | 32.2 KiB | [smlar_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/smlar_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-smlar` | 1.0 | `d12.x86_64` | pigsty | 82.3 KiB | [postgresql-15-smlar_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-15-smlar_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-smlar` | 1.0 | `d12.aarch64` | pigsty | 80.5 KiB | [postgresql-15-smlar_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-15-smlar_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-smlar` | 1.0 | `u22.x86_64` | pigsty | 86.2 KiB | [postgresql-15-smlar_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-15-smlar_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-smlar` | 1.0 | `u22.aarch64` | pigsty | 84.5 KiB | [postgresql-15-smlar_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-15-smlar_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-smlar` | 1.0 | `u24.x86_64` | pigsty | 76.6 KiB | [postgresql-15-smlar_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-15-smlar_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-smlar` | 1.0 | `u24.aarch64` | pigsty | 75.7 KiB | [postgresql-15-smlar_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-15-smlar_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `smlar_14` | 1.0 | `el8.x86_64` | pigsty | 32.9 KiB | [smlar_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/smlar_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `smlar_14` | 1.0 | `el8.aarch64` | pigsty | 30.9 KiB | [smlar_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/smlar_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `smlar_14` | 1.0 | `el9.x86_64` | pigsty | 33.6 KiB | [smlar_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/smlar_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `smlar_14` | 1.0 | `el9.aarch64` | pigsty | 32.2 KiB | [smlar_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/smlar_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-smlar` | 1.0 | `d12.x86_64` | pigsty | 82.2 KiB | [postgresql-14-smlar_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-14-smlar_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-smlar` | 1.0 | `d12.aarch64` | pigsty | 80.4 KiB | [postgresql-14-smlar_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-14-smlar_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-smlar` | 1.0 | `u22.x86_64` | pigsty | 86.1 KiB | [postgresql-14-smlar_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-14-smlar_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-smlar` | 1.0 | `u22.aarch64` | pigsty | 84.4 KiB | [postgresql-14-smlar_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-14-smlar_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-smlar` | 1.0 | `u24.x86_64` | pigsty | 76.6 KiB | [postgresql-14-smlar_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-14-smlar_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-smlar` | 1.0 | `u24.aarch64` | pigsty | 75.7 KiB | [postgresql-14-smlar_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-14-smlar_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `smlar_13` | 1.0 | `el8.x86_64` | pigsty | 32.6 KiB | [smlar_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/smlar_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `smlar_13` | 1.0 | `el8.aarch64` | pigsty | 30.8 KiB | [smlar_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/smlar_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `smlar_13` | 1.0 | `el9.x86_64` | pigsty | 33.7 KiB | [smlar_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/smlar_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `smlar_13` | 1.0 | `el9.aarch64` | pigsty | 32.3 KiB | [smlar_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/smlar_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-smlar` | 1.0 | `d12.x86_64` | pigsty | 81.7 KiB | [postgresql-13-smlar_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-13-smlar_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-smlar` | 1.0 | `d12.aarch64` | pigsty | 79.9 KiB | [postgresql-13-smlar_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-13-smlar_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-smlar` | 1.0 | `u22.x86_64` | pigsty | 85.8 KiB | [postgresql-13-smlar_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-13-smlar_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-smlar` | 1.0 | `u22.aarch64` | pigsty | 84.1 KiB | [postgresql-13-smlar_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-13-smlar_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-smlar` | 1.0 | `u24.x86_64` | pigsty | 76.5 KiB | [postgresql-13-smlar_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-13-smlar_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-smlar` | 1.0 | `u24.aarch64` | pigsty | 75.7 KiB | [postgresql-13-smlar_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-13-smlar_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/jirutka/smlar" title="Repository" icon="github" subtitle="github.com/jirutka/smlar" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="smlar-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get smlar; # get smlar source code
pig build dep smlar; # install build dependencies
pig build pkg smlar; # build extension rpm or deb
pig build ext smlar; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install smlar; # install by extension name, for the current active PG version
pig ext install smlar; # install via package alias, for the active PG version
pig ext install smlar -v 17;   # install for PG 17
pig ext install smlar -v 16;   # install for PG 16
pig ext install smlar -v 15;   # install for PG 15
pig ext install smlar -v 14;   # install for PG 14
pig ext install smlar -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION smlar;
```

