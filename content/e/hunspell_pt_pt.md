---
title: "hunspell_pt_pt"
linkTitle: "hunspell_pt_pt"
description: "Portuguese Hunspell Dictionary"
weight: 2177
categories: ["FTS"]
width: full
---

Portuguese Hunspell Dictionary


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2177** | {{< badge content="hunspell_pt_pt" link="https://github.com/postgrespro/hunspell_dicts" >}} | {{< ext "hunspell_pt_pt" >}} | `1.0` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "Data" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "hunspell_en_us" >}} {{< ext "hunspell_ne_np" >}} {{< ext "hunspell_nn_no" >}} {{< ext "hunspell_ru_ru" >}} {{< ext "hunspell_ru_ru_aot" >}} {{< ext "hunspell_cs_cz" >}} {{< ext "hunspell_de_de" >}} {{< ext "hunspell_fr" >}} |

> [!Note] WARNING, conflict with pg built-in dict file, not recommended


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/hunspell_pt_pt" >}} | `1.0` | {{< bg "18" "hunspell_pt_pt_18" "green" >}} {{< bg "17" "hunspell_pt_pt_17" "green" >}} {{< bg "16" "hunspell_pt_pt_16" "green" >}} {{< bg "15" "hunspell_pt_pt_15" "green" >}} {{< bg "14" "hunspell_pt_pt_14" "green" >}} {{< bg "13" "hunspell_pt_pt_13" "green" >}} | `hunspell_pt_pt_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/hunspell_pt_pt" >}} | `1.0` | {{< bg "18" "postgresql-18-hunspell-pt-pt" "green" >}} {{< bg "17" "postgresql-17-hunspell-pt-pt" "green" >}} {{< bg "16" "postgresql-16-hunspell-pt-pt" "green" >}} {{< bg "15" "postgresql-15-hunspell-pt-pt" "green" >}} {{< bg "14" "postgresql-14-hunspell-pt-pt" "green" >}} {{< bg "13" "postgresql-13-hunspell-pt-pt" "green" >}} | `postgresql-$v-hunspell-pt-pt` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "hunspell_pt_pt_18 : THROW 0" "purple" >}}      |      {{< bg "PIGSTY 1.0" "hunspell_pt_pt_17 : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "hunspell_pt_pt_16 : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "hunspell_pt_pt_15 : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "hunspell_pt_pt_14 : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "hunspell_pt_pt_13 : THROW 1" "purple" >}}      |
|    `el8.aarch64`    |      {{< bg "MISS" "hunspell_pt_pt_18 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_17 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_16 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_15 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_14 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_13 : THROW 0" "purple" >}}      |
|    `el9.x86_64`    |      {{< bg "MISS" "hunspell_pt_pt_18 : THROW 0" "purple" >}}      |      {{< bg "PIGSTY 1.0" "hunspell_pt_pt_17 : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "hunspell_pt_pt_16 : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "hunspell_pt_pt_15 : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "hunspell_pt_pt_14 : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "hunspell_pt_pt_13 : THROW 1" "purple" >}}      |
|    `el9.aarch64`    |      {{< bg "MISS" "hunspell_pt_pt_18 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_17 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_16 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_15 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_14 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_13 : THROW 0" "purple" >}}      |
|    `el10.x86_64`    |      {{< bg "MISS" "hunspell_pt_pt_18 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_17 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_16 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_15 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_14 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_13 : THROW 0" "purple" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "hunspell_pt_pt_18 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_17 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_16 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_15 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_14 : THROW 0" "purple" >}}      |      {{< bg "MISS" "hunspell_pt_pt_13 : THROW 0" "purple" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-hunspell-pt-pt : THROW 0" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-17-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-16-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-15-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-14-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-13-hunspell-pt-pt : THROW 1" "purple" >}}      |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-hunspell-pt-pt : THROW 0" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-17-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-16-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-15-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-14-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-13-hunspell-pt-pt : THROW 1" "purple" >}}      |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-hunspell-pt-pt : THROW 0" "purple" >}}      |      {{< bg "MISS" "postgresql-17-hunspell-pt-pt : THROW 0" "purple" >}}      |      {{< bg "MISS" "postgresql-16-hunspell-pt-pt : THROW 0" "purple" >}}      |      {{< bg "MISS" "postgresql-15-hunspell-pt-pt : THROW 0" "purple" >}}      |      {{< bg "MISS" "postgresql-14-hunspell-pt-pt : THROW 0" "purple" >}}      |      {{< bg "MISS" "postgresql-13-hunspell-pt-pt : THROW 0" "purple" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-hunspell-pt-pt : THROW 0" "purple" >}}      |      {{< bg "MISS" "postgresql-17-hunspell-pt-pt : THROW 0" "purple" >}}      |      {{< bg "MISS" "postgresql-16-hunspell-pt-pt : THROW 0" "purple" >}}      |      {{< bg "MISS" "postgresql-15-hunspell-pt-pt : THROW 0" "purple" >}}      |      {{< bg "MISS" "postgresql-14-hunspell-pt-pt : THROW 0" "purple" >}}      |      {{< bg "MISS" "postgresql-13-hunspell-pt-pt : THROW 0" "purple" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-hunspell-pt-pt : THROW 0" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-17-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-16-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-15-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-14-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-13-hunspell-pt-pt : THROW 1" "purple" >}}      |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-hunspell-pt-pt : THROW 0" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-17-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-16-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-15-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-14-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-13-hunspell-pt-pt : THROW 1" "purple" >}}      |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-hunspell-pt-pt : THROW 0" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-17-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-16-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-15-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-14-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-13-hunspell-pt-pt : THROW 1" "purple" >}}      |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-hunspell-pt-pt : THROW 0" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-17-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-16-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-15-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-14-hunspell-pt-pt : THROW 1" "purple" >}}      |      {{< bg "PIGSTY 1.0" "postgresql-13-hunspell-pt-pt : THROW 1" "purple" >}}      |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hunspell_pt_pt_17` | 1.0 | `el8.x86_64` | pigsty | 224.6 KiB | [hunspell_pt_pt_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hunspell_pt_pt_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `hunspell_pt_pt_17` | 1.0 | `el9.x86_64` | pigsty | 207.9 KiB | [hunspell_pt_pt_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hunspell_pt_pt_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-hunspell-pt-pt` | 1.0 | `d12.x86_64` | pigsty | 189.7 KiB | [postgresql-17-hunspell-pt-pt_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-pt-pt/postgresql-17-hunspell-pt-pt_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-hunspell-pt-pt` | 1.0 | `d12.aarch64` | pigsty | 189.7 KiB | [postgresql-17-hunspell-pt-pt_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-pt-pt/postgresql-17-hunspell-pt-pt_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-hunspell-pt-pt` | 1.0 | `u22.x86_64` | pigsty | 200.5 KiB | [postgresql-17-hunspell-pt-pt_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-pt-pt/postgresql-17-hunspell-pt-pt_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-hunspell-pt-pt` | 1.0 | `u22.aarch64` | pigsty | 200.5 KiB | [postgresql-17-hunspell-pt-pt_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-pt-pt/postgresql-17-hunspell-pt-pt_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-hunspell-pt-pt` | 1.0 | `u24.x86_64` | pigsty | 199.9 KiB | [postgresql-17-hunspell-pt-pt_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-pt-pt/postgresql-17-hunspell-pt-pt_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-hunspell-pt-pt` | 1.0 | `u24.aarch64` | pigsty | 199.9 KiB | [postgresql-17-hunspell-pt-pt_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-pt-pt/postgresql-17-hunspell-pt-pt_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hunspell_pt_pt_16` | 1.0 | `el8.x86_64` | pigsty | 224.6 KiB | [hunspell_pt_pt_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hunspell_pt_pt_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `hunspell_pt_pt_16` | 1.0 | `el9.x86_64` | pigsty | 208.2 KiB | [hunspell_pt_pt_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hunspell_pt_pt_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-16-hunspell-pt-pt` | 1.0 | `d12.x86_64` | pigsty | 188.7 KiB | [postgresql-16-hunspell-pt-pt_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-pt-pt/postgresql-16-hunspell-pt-pt_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-hunspell-pt-pt` | 1.0 | `d12.aarch64` | pigsty | 188.7 KiB | [postgresql-16-hunspell-pt-pt_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-pt-pt/postgresql-16-hunspell-pt-pt_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-hunspell-pt-pt` | 1.0 | `u22.x86_64` | pigsty | 199.7 KiB | [postgresql-16-hunspell-pt-pt_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-pt-pt/postgresql-16-hunspell-pt-pt_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-hunspell-pt-pt` | 1.0 | `u22.aarch64` | pigsty | 199.7 KiB | [postgresql-16-hunspell-pt-pt_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-pt-pt/postgresql-16-hunspell-pt-pt_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-hunspell-pt-pt` | 1.0 | `u24.x86_64` | pigsty | 199.8 KiB | [postgresql-16-hunspell-pt-pt_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-pt-pt/postgresql-16-hunspell-pt-pt_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-hunspell-pt-pt` | 1.0 | `u24.aarch64` | pigsty | 199.8 KiB | [postgresql-16-hunspell-pt-pt_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-pt-pt/postgresql-16-hunspell-pt-pt_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hunspell_pt_pt_15` | 1.0 | `el8.x86_64` | pigsty | 224.6 KiB | [hunspell_pt_pt_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hunspell_pt_pt_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `hunspell_pt_pt_15` | 1.0 | `el9.x86_64` | pigsty | 208.3 KiB | [hunspell_pt_pt_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hunspell_pt_pt_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-15-hunspell-pt-pt` | 1.0 | `d12.x86_64` | pigsty | 189.8 KiB | [postgresql-15-hunspell-pt-pt_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-pt-pt/postgresql-15-hunspell-pt-pt_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-hunspell-pt-pt` | 1.0 | `d12.aarch64` | pigsty | 189.8 KiB | [postgresql-15-hunspell-pt-pt_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-pt-pt/postgresql-15-hunspell-pt-pt_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-hunspell-pt-pt` | 1.0 | `u22.x86_64` | pigsty | 200.0 KiB | [postgresql-15-hunspell-pt-pt_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-pt-pt/postgresql-15-hunspell-pt-pt_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-hunspell-pt-pt` | 1.0 | `u22.aarch64` | pigsty | 200.0 KiB | [postgresql-15-hunspell-pt-pt_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-pt-pt/postgresql-15-hunspell-pt-pt_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-hunspell-pt-pt` | 1.0 | `u24.x86_64` | pigsty | 199.9 KiB | [postgresql-15-hunspell-pt-pt_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-pt-pt/postgresql-15-hunspell-pt-pt_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-hunspell-pt-pt` | 1.0 | `u24.aarch64` | pigsty | 199.9 KiB | [postgresql-15-hunspell-pt-pt_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-pt-pt/postgresql-15-hunspell-pt-pt_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hunspell_pt_pt_14` | 1.0 | `el8.x86_64` | pigsty | 224.6 KiB | [hunspell_pt_pt_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hunspell_pt_pt_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `hunspell_pt_pt_14` | 1.0 | `el9.x86_64` | pigsty | 208.2 KiB | [hunspell_pt_pt_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hunspell_pt_pt_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-14-hunspell-pt-pt` | 1.0 | `d12.x86_64` | pigsty | 188.6 KiB | [postgresql-14-hunspell-pt-pt_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-pt-pt/postgresql-14-hunspell-pt-pt_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-hunspell-pt-pt` | 1.0 | `d12.aarch64` | pigsty | 188.6 KiB | [postgresql-14-hunspell-pt-pt_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-pt-pt/postgresql-14-hunspell-pt-pt_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-hunspell-pt-pt` | 1.0 | `u22.x86_64` | pigsty | 200.1 KiB | [postgresql-14-hunspell-pt-pt_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-pt-pt/postgresql-14-hunspell-pt-pt_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-hunspell-pt-pt` | 1.0 | `u22.aarch64` | pigsty | 200.1 KiB | [postgresql-14-hunspell-pt-pt_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-pt-pt/postgresql-14-hunspell-pt-pt_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-hunspell-pt-pt` | 1.0 | `u24.x86_64` | pigsty | 199.8 KiB | [postgresql-14-hunspell-pt-pt_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-pt-pt/postgresql-14-hunspell-pt-pt_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-hunspell-pt-pt` | 1.0 | `u24.aarch64` | pigsty | 199.8 KiB | [postgresql-14-hunspell-pt-pt_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-pt-pt/postgresql-14-hunspell-pt-pt_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hunspell_pt_pt_13` | 1.0 | `el8.x86_64` | pigsty | 224.6 KiB | [hunspell_pt_pt_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hunspell_pt_pt_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `hunspell_pt_pt_13` | 1.0 | `el9.x86_64` | pigsty | 208.2 KiB | [hunspell_pt_pt_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hunspell_pt_pt_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-hunspell-pt-pt` | 1.0 | `d12.x86_64` | pigsty | 189.9 KiB | [postgresql-13-hunspell-pt-pt_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-pt-pt/postgresql-13-hunspell-pt-pt_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-hunspell-pt-pt` | 1.0 | `d12.aarch64` | pigsty | 189.9 KiB | [postgresql-13-hunspell-pt-pt_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-pt-pt/postgresql-13-hunspell-pt-pt_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-hunspell-pt-pt` | 1.0 | `u22.x86_64` | pigsty | 199.7 KiB | [postgresql-13-hunspell-pt-pt_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-pt-pt/postgresql-13-hunspell-pt-pt_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-hunspell-pt-pt` | 1.0 | `u22.aarch64` | pigsty | 199.7 KiB | [postgresql-13-hunspell-pt-pt_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-pt-pt/postgresql-13-hunspell-pt-pt_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-hunspell-pt-pt` | 1.0 | `u24.x86_64` | pigsty | 199.7 KiB | [postgresql-13-hunspell-pt-pt_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-pt-pt/postgresql-13-hunspell-pt-pt_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-hunspell-pt-pt` | 1.0 | `u24.aarch64` | pigsty | 199.7 KiB | [postgresql-13-hunspell-pt-pt_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-pt-pt/postgresql-13-hunspell-pt-pt_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/postgrespro/hunspell_dicts" title="Repository" icon="github" subtitle="github.com/postgrespro/hunspell_dicts" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="hunspell-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get hunspell_pt_pt; # get hunspell_pt_pt source code
pig build dep hunspell_pt_pt; # install build dependencies
pig build pkg hunspell_pt_pt; # build extension rpm or deb
pig build ext hunspell_pt_pt; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install hunspell_pt_pt; # install by extension name, for the current active PG version
pig ext install hunspell_pt_pt; # install via package alias, for the active PG version
pig ext install hunspell_pt_pt -v 18;   # install for PG 18
pig ext install hunspell_pt_pt -v 17;   # install for PG 17
pig ext install hunspell_pt_pt -v 16;   # install for PG 16
pig ext install hunspell_pt_pt -v 15;   # install for PG 15
pig ext install hunspell_pt_pt -v 14;   # install for PG 14
pig ext install hunspell_pt_pt -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION hunspell_pt_pt;
```

