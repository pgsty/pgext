---
title: "hunspell_en_us"
linkTitle: "hunspell_en_us"
description: "en_US Hunspell Dictionary"
weight: 2172
categories: ["FTS"]
width: full
---

en_US Hunspell Dictionary


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2172** | {{< badge content="hunspell_en_us" link="https://github.com/postgrespro/hunspell_dicts" >}} | {{< ext "hunspell_en_us" >}} | `1.0` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "Data" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "hunspell_cs_cz" >}} {{< ext "hunspell_de_de" >}} {{< ext "hunspell_fr" >}} {{< ext "hunspell_nl_nl" >}} {{< ext "hunspell_ne_np" >}} {{< ext "hunspell_nn_no" >}} {{< ext "hunspell_pt_pt" >}} {{< ext "hunspell_ru_ru" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/hunspell_en_us" >}} | `1.0` | {{< bg "18" "hunspell_en_us_18" "red" >}} {{< bg "17" "hunspell_en_us_17" "green" >}} {{< bg "16" "hunspell_en_us_16" "green" >}} {{< bg "15" "hunspell_en_us_15" "green" >}} {{< bg "14" "hunspell_en_us_14" "green" >}} | `hunspell_en_us_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/hunspell_en_us" >}} | `1.0` | {{< bg "18" "postgresql-18-hunspell-en-us" "red" >}} {{< bg "17" "postgresql-17-hunspell-en-us" "green" >}} {{< bg "16" "postgresql-16-hunspell-en-us" "green" >}} {{< bg "15" "postgresql-15-hunspell-en-us" "green" >}} {{< bg "14" "postgresql-14-hunspell-en-us" "green" >}} | `postgresql-$v-hunspell-en-us` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "hunspell_en_us_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "hunspell_en_us_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_en_us_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_en_us_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_en_us_14 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "hunspell_en_us_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "hunspell_en_us_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_en_us_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_en_us_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_en_us_14 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "hunspell_en_us_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "hunspell_en_us_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_en_us_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_en_us_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_en_us_14 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "hunspell_en_us_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "hunspell_en_us_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_en_us_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_en_us_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "hunspell_en_us_14 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-hunspell-en-us : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-hunspell-en-us : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-hunspell-en-us : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-hunspell-en-us : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-hunspell-en-us : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-hunspell-en-us : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-hunspell-en-us : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-hunspell-en-us : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-hunspell-en-us : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-hunspell-en-us : AVAIL 1" "green" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-hunspell-en-us : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-hunspell-en-us : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-hunspell-en-us : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-hunspell-en-us : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-hunspell-en-us : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-hunspell-en-us : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-hunspell-en-us : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-hunspell-en-us : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-hunspell-en-us : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-hunspell-en-us : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-hunspell-en-us : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-hunspell-en-us : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-hunspell-en-us : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-hunspell-en-us : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-hunspell-en-us : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-hunspell-en-us : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-hunspell-en-us : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-hunspell-en-us : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-hunspell-en-us : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-hunspell-en-us : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hunspell_en_us_17` | 1.0 | `el8.x86_64` | pigsty | 186.7 KiB | [hunspell_en_us_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hunspell_en_us_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `hunspell_en_us_17` | 1.0 | `el8.aarch64` | pigsty | 186.7 KiB | [hunspell_en_us_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hunspell_en_us_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `hunspell_en_us_17` | 1.0 | `el9.x86_64` | pigsty | 174.4 KiB | [hunspell_en_us_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hunspell_en_us_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `hunspell_en_us_17` | 1.0 | `el9.aarch64` | pigsty | 174.3 KiB | [hunspell_en_us_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hunspell_en_us_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-hunspell-en-us` | 1.0 | `d12.x86_64` | pigsty | 157.7 KiB | [postgresql-17-hunspell-en-us_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-en-us/postgresql-17-hunspell-en-us_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-hunspell-en-us` | 1.0 | `d12.aarch64` | pigsty | 157.7 KiB | [postgresql-17-hunspell-en-us_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-en-us/postgresql-17-hunspell-en-us_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-hunspell-en-us` | 1.0 | `u22.x86_64` | pigsty | 166.9 KiB | [postgresql-17-hunspell-en-us_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-en-us/postgresql-17-hunspell-en-us_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-hunspell-en-us` | 1.0 | `u22.aarch64` | pigsty | 166.9 KiB | [postgresql-17-hunspell-en-us_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-en-us/postgresql-17-hunspell-en-us_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-hunspell-en-us` | 1.0 | `u24.x86_64` | pigsty | 166.5 KiB | [postgresql-17-hunspell-en-us_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-en-us/postgresql-17-hunspell-en-us_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-hunspell-en-us` | 1.0 | `u24.aarch64` | pigsty | 166.5 KiB | [postgresql-17-hunspell-en-us_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-en-us/postgresql-17-hunspell-en-us_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hunspell_en_us_16` | 1.0 | `el8.x86_64` | pigsty | 186.7 KiB | [hunspell_en_us_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hunspell_en_us_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `hunspell_en_us_16` | 1.0 | `el8.aarch64` | pigsty | 186.7 KiB | [hunspell_en_us_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hunspell_en_us_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `hunspell_en_us_16` | 1.0 | `el9.x86_64` | pigsty | 174.4 KiB | [hunspell_en_us_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hunspell_en_us_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `hunspell_en_us_16` | 1.0 | `el9.aarch64` | pigsty | 174.3 KiB | [hunspell_en_us_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hunspell_en_us_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-hunspell-en-us` | 1.0 | `d12.x86_64` | pigsty | 157.7 KiB | [postgresql-16-hunspell-en-us_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-en-us/postgresql-16-hunspell-en-us_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-hunspell-en-us` | 1.0 | `d12.aarch64` | pigsty | 157.7 KiB | [postgresql-16-hunspell-en-us_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-en-us/postgresql-16-hunspell-en-us_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-hunspell-en-us` | 1.0 | `u22.x86_64` | pigsty | 166.9 KiB | [postgresql-16-hunspell-en-us_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-en-us/postgresql-16-hunspell-en-us_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-hunspell-en-us` | 1.0 | `u22.aarch64` | pigsty | 166.9 KiB | [postgresql-16-hunspell-en-us_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-en-us/postgresql-16-hunspell-en-us_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-hunspell-en-us` | 1.0 | `u24.x86_64` | pigsty | 166.5 KiB | [postgresql-16-hunspell-en-us_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-en-us/postgresql-16-hunspell-en-us_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-hunspell-en-us` | 1.0 | `u24.aarch64` | pigsty | 166.5 KiB | [postgresql-16-hunspell-en-us_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-en-us/postgresql-16-hunspell-en-us_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hunspell_en_us_15` | 1.0 | `el8.x86_64` | pigsty | 186.7 KiB | [hunspell_en_us_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hunspell_en_us_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `hunspell_en_us_15` | 1.0 | `el8.aarch64` | pigsty | 186.7 KiB | [hunspell_en_us_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hunspell_en_us_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `hunspell_en_us_15` | 1.0 | `el9.x86_64` | pigsty | 174.4 KiB | [hunspell_en_us_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hunspell_en_us_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `hunspell_en_us_15` | 1.0 | `el9.aarch64` | pigsty | 174.3 KiB | [hunspell_en_us_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hunspell_en_us_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-hunspell-en-us` | 1.0 | `d12.x86_64` | pigsty | 157.6 KiB | [postgresql-15-hunspell-en-us_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-en-us/postgresql-15-hunspell-en-us_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-hunspell-en-us` | 1.0 | `d12.aarch64` | pigsty | 157.6 KiB | [postgresql-15-hunspell-en-us_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-en-us/postgresql-15-hunspell-en-us_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-hunspell-en-us` | 1.0 | `u22.x86_64` | pigsty | 167.1 KiB | [postgresql-15-hunspell-en-us_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-en-us/postgresql-15-hunspell-en-us_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-hunspell-en-us` | 1.0 | `u22.aarch64` | pigsty | 167.1 KiB | [postgresql-15-hunspell-en-us_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-en-us/postgresql-15-hunspell-en-us_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-hunspell-en-us` | 1.0 | `u24.x86_64` | pigsty | 166.5 KiB | [postgresql-15-hunspell-en-us_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-en-us/postgresql-15-hunspell-en-us_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-hunspell-en-us` | 1.0 | `u24.aarch64` | pigsty | 166.5 KiB | [postgresql-15-hunspell-en-us_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-en-us/postgresql-15-hunspell-en-us_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hunspell_en_us_14` | 1.0 | `el8.x86_64` | pigsty | 186.7 KiB | [hunspell_en_us_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hunspell_en_us_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `hunspell_en_us_14` | 1.0 | `el8.aarch64` | pigsty | 186.7 KiB | [hunspell_en_us_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hunspell_en_us_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `hunspell_en_us_14` | 1.0 | `el9.x86_64` | pigsty | 174.3 KiB | [hunspell_en_us_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hunspell_en_us_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `hunspell_en_us_14` | 1.0 | `el9.aarch64` | pigsty | 174.3 KiB | [hunspell_en_us_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hunspell_en_us_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-hunspell-en-us` | 1.0 | `d12.x86_64` | pigsty | 157.7 KiB | [postgresql-14-hunspell-en-us_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-en-us/postgresql-14-hunspell-en-us_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-hunspell-en-us` | 1.0 | `d12.aarch64` | pigsty | 157.7 KiB | [postgresql-14-hunspell-en-us_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-en-us/postgresql-14-hunspell-en-us_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-hunspell-en-us` | 1.0 | `u22.x86_64` | pigsty | 166.9 KiB | [postgresql-14-hunspell-en-us_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-en-us/postgresql-14-hunspell-en-us_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-hunspell-en-us` | 1.0 | `u22.aarch64` | pigsty | 166.9 KiB | [postgresql-14-hunspell-en-us_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-en-us/postgresql-14-hunspell-en-us_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-hunspell-en-us` | 1.0 | `u24.x86_64` | pigsty | 166.5 KiB | [postgresql-14-hunspell-en-us_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-en-us/postgresql-14-hunspell-en-us_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-hunspell-en-us` | 1.0 | `u24.aarch64` | pigsty | 166.5 KiB | [postgresql-14-hunspell-en-us_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-en-us/postgresql-14-hunspell-en-us_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/postgrespro/hunspell_dicts" title="Repository" icon="github" subtitle="github.com/postgrespro/hunspell_dicts" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="hunspell-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get hunspell_en_us; # get hunspell_en_us source code
pig build dep hunspell_en_us; # install build dependencies
pig build pkg hunspell_en_us; # build extension rpm or deb
pig build ext hunspell_en_us; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install hunspell_en_us; # install by extension name, for the current active PG version
pig ext install hunspell_en_us; # install via package alias, for the active PG version
pig ext install hunspell_en_us -v 18;   # install for PG 18
pig ext install hunspell_en_us -v 17;   # install for PG 17
pig ext install hunspell_en_us -v 16;   # install for PG 16
pig ext install hunspell_en_us -v 15;   # install for PG 15
pig ext install hunspell_en_us -v 14;   # install for PG 14
pig ext install hunspell_en_us -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION hunspell_en_us;
```

