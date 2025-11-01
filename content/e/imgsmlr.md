---
title: "imgsmlr"
linkTitle: "imgsmlr"
description: "Image similarity with haar"
weight: 2860
categories: ["FEAT"]
width: full
---

Image similarity with haar


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2860** | {{< badge content="imgsmlr" link="https://github.com/postgrespro/imgsmlr" >}} | {{< ext "imgsmlr" >}} | `1.0` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "age" >}} {{< ext "hll" >}} {{< ext "rum" >}} {{< ext "pg_graphql" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_hint_plan" >}} {{< ext "hypopg" >}} |

> [!Note] breaks on el10


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/imgsmlr" >}} | `1.0` | {{< bg "18" "imgsmlr_18*" "green" >}} {{< bg "17" "imgsmlr_17*" "green" >}} {{< bg "16" "imgsmlr_16*" "green" >}} {{< bg "15" "imgsmlr_15*" "green" >}} {{< bg "14" "imgsmlr_14*" "green" >}} {{< bg "13" "imgsmlr_13*" "green" >}} | `imgsmlr_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/imgsmlr" >}} | `1.0` | {{< bg "18" "postgresql-18-imgsmlr" "green" >}} {{< bg "17" "postgresql-17-imgsmlr" "green" >}} {{< bg "16" "postgresql-16-imgsmlr" "green" >}} {{< bg "15" "postgresql-15-imgsmlr" "green" >}} {{< bg "14" "postgresql-14-imgsmlr" "green" >}} {{< bg "13" "postgresql-13-imgsmlr" "green" >}} | `postgresql-$v-imgsmlr` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 1.0" "imgsmlr_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 1.0" "imgsmlr_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 1.0" "imgsmlr_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 1.0" "imgsmlr_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "imgsmlr_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    |      {{< bg "MISS" "imgsmlr_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "imgsmlr_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "imgsmlr_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "imgsmlr_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "imgsmlr_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "imgsmlr_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "imgsmlr_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "imgsmlr_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "imgsmlr_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "imgsmlr_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "imgsmlr_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "imgsmlr_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-imgsmlr : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-imgsmlr : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-imgsmlr : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-imgsmlr : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-imgsmlr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-imgsmlr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-imgsmlr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-imgsmlr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-imgsmlr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-imgsmlr : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-imgsmlr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-imgsmlr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-imgsmlr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-imgsmlr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-imgsmlr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-imgsmlr : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-imgsmlr : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-imgsmlr : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-imgsmlr : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-imgsmlr : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-imgsmlr : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-imgsmlr : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-imgsmlr : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-imgsmlr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-imgsmlr : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `imgsmlr_18` | 1.0 | `el8.x86_64` | pigsty | 21.6 KiB | [imgsmlr_18-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/imgsmlr_18-1.0-2PIGSTY.el8.x86_64.rpm) |
| `imgsmlr_18` | 1.0 | `el8.aarch64` | pigsty | 21.2 KiB | [imgsmlr_18-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/imgsmlr_18-1.0-2PIGSTY.el8.aarch64.rpm) |
| `imgsmlr_18` | 1.0 | `el9.x86_64` | pigsty | 21.2 KiB | [imgsmlr_18-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/imgsmlr_18-1.0-2PIGSTY.el9.x86_64.rpm) |
| `imgsmlr_18` | 1.0 | `el9.aarch64` | pigsty | 20.9 KiB | [imgsmlr_18-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/imgsmlr_18-1.0-2PIGSTY.el9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `imgsmlr_17` | 1.0 | `el8.x86_64` | pigsty | 21.6 KiB | [imgsmlr_17-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/imgsmlr_17-1.0-2PIGSTY.el8.x86_64.rpm) |
| `imgsmlr_17` | 1.0 | `el8.aarch64` | pigsty | 21.2 KiB | [imgsmlr_17-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/imgsmlr_17-1.0-2PIGSTY.el8.aarch64.rpm) |
| `imgsmlr_17` | 1.0 | `el9.x86_64` | pigsty | 21.2 KiB | [imgsmlr_17-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/imgsmlr_17-1.0-2PIGSTY.el9.x86_64.rpm) |
| `imgsmlr_17` | 1.0 | `el9.aarch64` | pigsty | 20.9 KiB | [imgsmlr_17-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/imgsmlr_17-1.0-2PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-imgsmlr` | 1.0 | `d12.x86_64` | pigsty | 32.8 KiB | [postgresql-17-imgsmlr_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/imgsmlr/postgresql-17-imgsmlr_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-imgsmlr` | 1.0 | `d12.aarch64` | pigsty | 32.3 KiB | [postgresql-17-imgsmlr_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/imgsmlr/postgresql-17-imgsmlr_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-imgsmlr` | 1.0 | `u22.x86_64` | pigsty | 34.1 KiB | [postgresql-17-imgsmlr_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/imgsmlr/postgresql-17-imgsmlr_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-imgsmlr` | 1.0 | `u22.aarch64` | pigsty | 33.2 KiB | [postgresql-17-imgsmlr_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/imgsmlr/postgresql-17-imgsmlr_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-imgsmlr` | 1.0 | `u24.x86_64` | pigsty | 32.7 KiB | [postgresql-17-imgsmlr_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/imgsmlr/postgresql-17-imgsmlr_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-imgsmlr` | 1.0 | `u24.aarch64` | pigsty | 32.3 KiB | [postgresql-17-imgsmlr_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/imgsmlr/postgresql-17-imgsmlr_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `imgsmlr_16` | 1.0 | `el8.x86_64` | pigsty | 21.6 KiB | [imgsmlr_16-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/imgsmlr_16-1.0-2PIGSTY.el8.x86_64.rpm) |
| `imgsmlr_16` | 1.0 | `el8.aarch64` | pigsty | 21.2 KiB | [imgsmlr_16-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/imgsmlr_16-1.0-2PIGSTY.el8.aarch64.rpm) |
| `imgsmlr_16` | 1.0 | `el9.x86_64` | pigsty | 21.2 KiB | [imgsmlr_16-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/imgsmlr_16-1.0-2PIGSTY.el9.x86_64.rpm) |
| `imgsmlr_16` | 1.0 | `el9.aarch64` | pigsty | 20.9 KiB | [imgsmlr_16-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/imgsmlr_16-1.0-2PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-imgsmlr` | 1.0 | `d12.x86_64` | pigsty | 32.8 KiB | [postgresql-16-imgsmlr_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/imgsmlr/postgresql-16-imgsmlr_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-imgsmlr` | 1.0 | `d12.aarch64` | pigsty | 32.2 KiB | [postgresql-16-imgsmlr_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/imgsmlr/postgresql-16-imgsmlr_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-imgsmlr` | 1.0 | `u22.x86_64` | pigsty | 34.1 KiB | [postgresql-16-imgsmlr_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/imgsmlr/postgresql-16-imgsmlr_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-imgsmlr` | 1.0 | `u22.aarch64` | pigsty | 33.2 KiB | [postgresql-16-imgsmlr_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/imgsmlr/postgresql-16-imgsmlr_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-imgsmlr` | 1.0 | `u24.x86_64` | pigsty | 32.7 KiB | [postgresql-16-imgsmlr_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/imgsmlr/postgresql-16-imgsmlr_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-imgsmlr` | 1.0 | `u24.aarch64` | pigsty | 32.3 KiB | [postgresql-16-imgsmlr_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/imgsmlr/postgresql-16-imgsmlr_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `imgsmlr_15` | 1.0 | `el8.x86_64` | pigsty | 21.6 KiB | [imgsmlr_15-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/imgsmlr_15-1.0-2PIGSTY.el8.x86_64.rpm) |
| `imgsmlr_15` | 1.0 | `el8.aarch64` | pigsty | 21.2 KiB | [imgsmlr_15-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/imgsmlr_15-1.0-2PIGSTY.el8.aarch64.rpm) |
| `imgsmlr_15` | 1.0 | `el9.x86_64` | pigsty | 21.2 KiB | [imgsmlr_15-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/imgsmlr_15-1.0-2PIGSTY.el9.x86_64.rpm) |
| `imgsmlr_15` | 1.0 | `el9.aarch64` | pigsty | 20.9 KiB | [imgsmlr_15-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/imgsmlr_15-1.0-2PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-imgsmlr` | 1.0 | `d12.x86_64` | pigsty | 32.9 KiB | [postgresql-15-imgsmlr_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/imgsmlr/postgresql-15-imgsmlr_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-imgsmlr` | 1.0 | `d12.aarch64` | pigsty | 32.3 KiB | [postgresql-15-imgsmlr_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/imgsmlr/postgresql-15-imgsmlr_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-imgsmlr` | 1.0 | `u22.x86_64` | pigsty | 34.1 KiB | [postgresql-15-imgsmlr_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/imgsmlr/postgresql-15-imgsmlr_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-imgsmlr` | 1.0 | `u22.aarch64` | pigsty | 33.3 KiB | [postgresql-15-imgsmlr_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/imgsmlr/postgresql-15-imgsmlr_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-imgsmlr` | 1.0 | `u24.x86_64` | pigsty | 32.7 KiB | [postgresql-15-imgsmlr_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/imgsmlr/postgresql-15-imgsmlr_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-imgsmlr` | 1.0 | `u24.aarch64` | pigsty | 32.3 KiB | [postgresql-15-imgsmlr_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/imgsmlr/postgresql-15-imgsmlr_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `imgsmlr_14` | 1.0 | `el8.x86_64` | pigsty | 21.6 KiB | [imgsmlr_14-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/imgsmlr_14-1.0-2PIGSTY.el8.x86_64.rpm) |
| `imgsmlr_14` | 1.0 | `el8.aarch64` | pigsty | 21.2 KiB | [imgsmlr_14-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/imgsmlr_14-1.0-2PIGSTY.el8.aarch64.rpm) |
| `imgsmlr_14` | 1.0 | `el9.x86_64` | pigsty | 21.2 KiB | [imgsmlr_14-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/imgsmlr_14-1.0-2PIGSTY.el9.x86_64.rpm) |
| `imgsmlr_14` | 1.0 | `el9.aarch64` | pigsty | 20.9 KiB | [imgsmlr_14-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/imgsmlr_14-1.0-2PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-imgsmlr` | 1.0 | `d12.x86_64` | pigsty | 32.8 KiB | [postgresql-14-imgsmlr_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/imgsmlr/postgresql-14-imgsmlr_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-imgsmlr` | 1.0 | `d12.aarch64` | pigsty | 32.3 KiB | [postgresql-14-imgsmlr_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/imgsmlr/postgresql-14-imgsmlr_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-imgsmlr` | 1.0 | `u22.x86_64` | pigsty | 34.1 KiB | [postgresql-14-imgsmlr_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/imgsmlr/postgresql-14-imgsmlr_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-imgsmlr` | 1.0 | `u22.aarch64` | pigsty | 33.3 KiB | [postgresql-14-imgsmlr_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/imgsmlr/postgresql-14-imgsmlr_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-imgsmlr` | 1.0 | `u24.x86_64` | pigsty | 32.7 KiB | [postgresql-14-imgsmlr_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/imgsmlr/postgresql-14-imgsmlr_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-imgsmlr` | 1.0 | `u24.aarch64` | pigsty | 32.3 KiB | [postgresql-14-imgsmlr_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/imgsmlr/postgresql-14-imgsmlr_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `imgsmlr_13` | 1.0 | `el8.x86_64` | pigsty | 21.5 KiB | [imgsmlr_13-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/imgsmlr_13-1.0-2PIGSTY.el8.x86_64.rpm) |
| `imgsmlr_13` | 1.0 | `el8.aarch64` | pigsty | 21.2 KiB | [imgsmlr_13-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/imgsmlr_13-1.0-2PIGSTY.el8.aarch64.rpm) |
| `imgsmlr_13` | 1.0 | `el9.x86_64` | pigsty | 21.2 KiB | [imgsmlr_13-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/imgsmlr_13-1.0-2PIGSTY.el9.x86_64.rpm) |
| `imgsmlr_13` | 1.0 | `el9.aarch64` | pigsty | 20.9 KiB | [imgsmlr_13-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/imgsmlr_13-1.0-2PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-imgsmlr` | 1.0 | `d12.x86_64` | pigsty | 32.8 KiB | [postgresql-13-imgsmlr_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/imgsmlr/postgresql-13-imgsmlr_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-imgsmlr` | 1.0 | `d12.aarch64` | pigsty | 32.0 KiB | [postgresql-13-imgsmlr_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/imgsmlr/postgresql-13-imgsmlr_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-imgsmlr` | 1.0 | `u22.x86_64` | pigsty | 34.0 KiB | [postgresql-13-imgsmlr_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/imgsmlr/postgresql-13-imgsmlr_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-imgsmlr` | 1.0 | `u22.aarch64` | pigsty | 33.1 KiB | [postgresql-13-imgsmlr_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/imgsmlr/postgresql-13-imgsmlr_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-imgsmlr` | 1.0 | `u24.x86_64` | pigsty | 32.6 KiB | [postgresql-13-imgsmlr_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/imgsmlr/postgresql-13-imgsmlr_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-imgsmlr` | 1.0 | `u24.aarch64` | pigsty | 32.3 KiB | [postgresql-13-imgsmlr_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/imgsmlr/postgresql-13-imgsmlr_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/postgrespro/imgsmlr" title="Repository" icon="github" subtitle="github.com/postgrespro/imgsmlr" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="imgsmlr-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get imgsmlr; # get imgsmlr source code
pig build dep imgsmlr; # install build dependencies
pig build pkg imgsmlr; # build extension rpm or deb
pig build ext imgsmlr; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install imgsmlr; # install by extension name, for the current active PG version
pig ext install imgsmlr; # install via package alias, for the active PG version
pig ext install imgsmlr -v 18;   # install for PG 18
pig ext install imgsmlr -v 17;   # install for PG 17
pig ext install imgsmlr -v 16;   # install for PG 16
pig ext install imgsmlr -v 15;   # install for PG 15
pig ext install imgsmlr -v 14;   # install for PG 14
pig ext install imgsmlr -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION imgsmlr;
```

