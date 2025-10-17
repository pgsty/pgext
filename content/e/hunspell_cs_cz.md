---
title: "hunspell_cs_cz"
linkTitle: "hunspell_cs_cz"
description: "Czech Hunspell Dictionary"
weight: 2170
categories: ["Fts"]
width: full
---

Czech Hunspell Dictionary

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2170** | {{< badge content="hunspell_cs_cz" link="https://github.com/postgrespro/hunspell_dicts" >}} | {{< ext "hunspell_cs_cz" "hunspell_cs_cz" >}} | `1.0` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "Data" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="-----d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "hunspell_de_de" >}} {{< ext "hunspell_en_us" >}} {{< ext "hunspell_fr" >}} {{< ext "hunspell_nl_nl" >}} {{< ext "hunspell_ne_np" >}} {{< ext "hunspell_nn_no" >}} {{< ext "hunspell_pt_pt" >}} {{< ext "hunspell_ru_ru" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/hunspell_cs_cz" >}} | `1.0` | {{< badge content="18" color="red" alt="hunspell_cs_cz_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `hunspell_cs_cz_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/hunspell_cs_cz" >}} | `1.0` | {{< badge content="18" color="red" alt="postgresql-18-hunspell-cs-cz" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-hunspell-cs-cz` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "hunspell_cs_cz_18" >}}     | {{< pkg "hunspell_cs_cz_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/hunspell_cs_cz_17-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "hunspell_cs_cz_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/hunspell_cs_cz_16-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "hunspell_cs_cz_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/hunspell_cs_cz_15-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "hunspell_cs_cz_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/hunspell_cs_cz_14-1.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "hunspell_cs_cz_18" >}}     | {{< pkg "hunspell_cs_cz_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/hunspell_cs_cz_17-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "hunspell_cs_cz_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/hunspell_cs_cz_16-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "hunspell_cs_cz_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/hunspell_cs_cz_15-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "hunspell_cs_cz_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/hunspell_cs_cz_14-1.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "hunspell_cs_cz_18" >}}     | {{< pkg "hunspell_cs_cz_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/hunspell_cs_cz_17-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "hunspell_cs_cz_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/hunspell_cs_cz_16-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "hunspell_cs_cz_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/hunspell_cs_cz_15-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "hunspell_cs_cz_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/hunspell_cs_cz_14-1.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "hunspell_cs_cz_18" >}}     | {{< pkg "hunspell_cs_cz_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/hunspell_cs_cz_17-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "hunspell_cs_cz_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/hunspell_cs_cz_16-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "hunspell_cs_cz_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/hunspell_cs_cz_15-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "hunspell_cs_cz_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/hunspell_cs_cz_14-1.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-hunspell-cs-cz" >}}     | {{< pkg "postgresql-17-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-cs-cz/postgresql-17-hunspell-cs-cz_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-cs-cz/postgresql-16-hunspell-cs-cz_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-cs-cz/postgresql-15-hunspell-cs-cz_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-cs-cz/postgresql-14-hunspell-cs-cz_1.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-hunspell-cs-cz" >}}     | {{< pkg "postgresql-17-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-cs-cz/postgresql-17-hunspell-cs-cz_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-cs-cz/postgresql-16-hunspell-cs-cz_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-cs-cz/postgresql-15-hunspell-cs-cz_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-cs-cz/postgresql-14-hunspell-cs-cz_1.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-hunspell-cs-cz" >}}     | {{< pkg "postgresql-17-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-cs-cz/postgresql-17-hunspell-cs-cz_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-cs-cz/postgresql-16-hunspell-cs-cz_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-cs-cz/postgresql-15-hunspell-cs-cz_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-cs-cz/postgresql-14-hunspell-cs-cz_1.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-hunspell-cs-cz" >}}     | {{< pkg "postgresql-17-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-cs-cz/postgresql-17-hunspell-cs-cz_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-cs-cz/postgresql-16-hunspell-cs-cz_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-cs-cz/postgresql-15-hunspell-cs-cz_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-cs-cz/postgresql-14-hunspell-cs-cz_1.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-hunspell-cs-cz" >}}     | {{< pkg "postgresql-17-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-cs-cz/postgresql-17-hunspell-cs-cz_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-cs-cz/postgresql-16-hunspell-cs-cz_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-cs-cz/postgresql-15-hunspell-cs-cz_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-cs-cz/postgresql-14-hunspell-cs-cz_1.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-hunspell-cs-cz" >}}     | {{< pkg "postgresql-17-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-cs-cz/postgresql-17-hunspell-cs-cz_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-cs-cz/postgresql-16-hunspell-cs-cz_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-cs-cz/postgresql-15-hunspell-cs-cz_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-hunspell-cs-cz" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-cs-cz/postgresql-14-hunspell-cs-cz_1.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `hunspell_cs_cz_17` | 1.0 | `el8.x86_64` | pigsty | 596.3 KiB | [hunspell_cs_cz_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hunspell_cs_cz_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `hunspell_cs_cz_17` | 1.0 | `el8.aarch64` | pigsty | 596.3 KiB | [hunspell_cs_cz_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hunspell_cs_cz_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `hunspell_cs_cz_17` | 1.0 | `el9.aarch64` | pigsty | 528.8 KiB | [hunspell_cs_cz_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hunspell_cs_cz_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `hunspell_cs_cz_17` | 1.0 | `el9.x86_64` | pigsty | 529.2 KiB | [hunspell_cs_cz_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hunspell_cs_cz_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-hunspell-cs-cz` | 1.0 | `d12.x86_64` | pigsty | 503.5 KiB | [postgresql-17-hunspell-cs-cz_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-cs-cz/postgresql-17-hunspell-cs-cz_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-hunspell-cs-cz` | 1.0 | `d12.aarch64` | pigsty | 503.5 KiB | [postgresql-17-hunspell-cs-cz_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-cs-cz/postgresql-17-hunspell-cs-cz_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-hunspell-cs-cz` | 1.0 | `u22.x86_64` | pigsty | 522.0 KiB | [postgresql-17-hunspell-cs-cz_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-cs-cz/postgresql-17-hunspell-cs-cz_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-hunspell-cs-cz` | 1.0 | `u22.aarch64` | pigsty | 522.0 KiB | [postgresql-17-hunspell-cs-cz_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-cs-cz/postgresql-17-hunspell-cs-cz_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-hunspell-cs-cz` | 1.0 | `u24.x86_64` | pigsty | 521.9 KiB | [postgresql-17-hunspell-cs-cz_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-cs-cz/postgresql-17-hunspell-cs-cz_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-hunspell-cs-cz` | 1.0 | `u24.aarch64` | pigsty | 521.9 KiB | [postgresql-17-hunspell-cs-cz_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-cs-cz/postgresql-17-hunspell-cs-cz_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `hunspell_cs_cz_16` | 1.0 | `el8.x86_64` | pigsty | 596.3 KiB | [hunspell_cs_cz_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hunspell_cs_cz_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `hunspell_cs_cz_16` | 1.0 | `el8.aarch64` | pigsty | 596.3 KiB | [hunspell_cs_cz_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hunspell_cs_cz_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `hunspell_cs_cz_16` | 1.0 | `el9.x86_64` | pigsty | 529.2 KiB | [hunspell_cs_cz_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hunspell_cs_cz_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `hunspell_cs_cz_16` | 1.0 | `el9.aarch64` | pigsty | 528.8 KiB | [hunspell_cs_cz_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hunspell_cs_cz_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-hunspell-cs-cz` | 1.0 | `d12.x86_64` | pigsty | 503.1 KiB | [postgresql-16-hunspell-cs-cz_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-cs-cz/postgresql-16-hunspell-cs-cz_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-hunspell-cs-cz` | 1.0 | `d12.aarch64` | pigsty | 503.1 KiB | [postgresql-16-hunspell-cs-cz_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-cs-cz/postgresql-16-hunspell-cs-cz_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-hunspell-cs-cz` | 1.0 | `u22.aarch64` | pigsty | 521.8 KiB | [postgresql-16-hunspell-cs-cz_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-cs-cz/postgresql-16-hunspell-cs-cz_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-hunspell-cs-cz` | 1.0 | `u22.x86_64` | pigsty | 521.8 KiB | [postgresql-16-hunspell-cs-cz_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-cs-cz/postgresql-16-hunspell-cs-cz_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-hunspell-cs-cz` | 1.0 | `u24.x86_64` | pigsty | 520.7 KiB | [postgresql-16-hunspell-cs-cz_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-cs-cz/postgresql-16-hunspell-cs-cz_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-hunspell-cs-cz` | 1.0 | `u24.aarch64` | pigsty | 520.7 KiB | [postgresql-16-hunspell-cs-cz_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-cs-cz/postgresql-16-hunspell-cs-cz_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `hunspell_cs_cz_15` | 1.0 | `el8.x86_64` | pigsty | 596.3 KiB | [hunspell_cs_cz_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hunspell_cs_cz_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `hunspell_cs_cz_15` | 1.0 | `el8.aarch64` | pigsty | 596.3 KiB | [hunspell_cs_cz_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hunspell_cs_cz_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `hunspell_cs_cz_15` | 1.0 | `el9.x86_64` | pigsty | 529.2 KiB | [hunspell_cs_cz_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hunspell_cs_cz_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `hunspell_cs_cz_15` | 1.0 | `el9.aarch64` | pigsty | 528.7 KiB | [hunspell_cs_cz_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hunspell_cs_cz_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-hunspell-cs-cz` | 1.0 | `d12.aarch64` | pigsty | 503.4 KiB | [postgresql-15-hunspell-cs-cz_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-cs-cz/postgresql-15-hunspell-cs-cz_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-hunspell-cs-cz` | 1.0 | `d12.x86_64` | pigsty | 503.4 KiB | [postgresql-15-hunspell-cs-cz_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-cs-cz/postgresql-15-hunspell-cs-cz_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-hunspell-cs-cz` | 1.0 | `u22.aarch64` | pigsty | 521.5 KiB | [postgresql-15-hunspell-cs-cz_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-cs-cz/postgresql-15-hunspell-cs-cz_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-hunspell-cs-cz` | 1.0 | `u22.x86_64` | pigsty | 521.5 KiB | [postgresql-15-hunspell-cs-cz_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-cs-cz/postgresql-15-hunspell-cs-cz_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-hunspell-cs-cz` | 1.0 | `u24.x86_64` | pigsty | 521.6 KiB | [postgresql-15-hunspell-cs-cz_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-cs-cz/postgresql-15-hunspell-cs-cz_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-hunspell-cs-cz` | 1.0 | `u24.aarch64` | pigsty | 521.6 KiB | [postgresql-15-hunspell-cs-cz_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-cs-cz/postgresql-15-hunspell-cs-cz_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `hunspell_cs_cz_14` | 1.0 | `el8.x86_64` | pigsty | 596.3 KiB | [hunspell_cs_cz_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hunspell_cs_cz_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `hunspell_cs_cz_14` | 1.0 | `el8.aarch64` | pigsty | 596.3 KiB | [hunspell_cs_cz_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hunspell_cs_cz_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `hunspell_cs_cz_14` | 1.0 | `el9.x86_64` | pigsty | 528.9 KiB | [hunspell_cs_cz_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hunspell_cs_cz_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `hunspell_cs_cz_14` | 1.0 | `el9.aarch64` | pigsty | 529.0 KiB | [hunspell_cs_cz_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hunspell_cs_cz_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-hunspell-cs-cz` | 1.0 | `d12.x86_64` | pigsty | 503.4 KiB | [postgresql-14-hunspell-cs-cz_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-cs-cz/postgresql-14-hunspell-cs-cz_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-hunspell-cs-cz` | 1.0 | `d12.aarch64` | pigsty | 503.4 KiB | [postgresql-14-hunspell-cs-cz_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-cs-cz/postgresql-14-hunspell-cs-cz_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-hunspell-cs-cz` | 1.0 | `u22.x86_64` | pigsty | 521.5 KiB | [postgresql-14-hunspell-cs-cz_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-cs-cz/postgresql-14-hunspell-cs-cz_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-hunspell-cs-cz` | 1.0 | `u22.aarch64` | pigsty | 521.5 KiB | [postgresql-14-hunspell-cs-cz_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-cs-cz/postgresql-14-hunspell-cs-cz_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-hunspell-cs-cz` | 1.0 | `u24.x86_64` | pigsty | 521.2 KiB | [postgresql-14-hunspell-cs-cz_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-cs-cz/postgresql-14-hunspell-cs-cz_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-hunspell-cs-cz` | 1.0 | `u24.aarch64` | pigsty | 521.2 KiB | [postgresql-14-hunspell-cs-cz_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-cs-cz/postgresql-14-hunspell-cs-cz_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `hunspell_cs_cz_13` | 1.0 | `el8.aarch64` | pigsty | 596.3 KiB | [hunspell_cs_cz_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hunspell_cs_cz_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `hunspell_cs_cz_13` | 1.0 | `el8.x86_64` | pigsty | 596.3 KiB | [hunspell_cs_cz_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hunspell_cs_cz_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `hunspell_cs_cz_13` | 1.0 | `el9.aarch64` | pigsty | 528.8 KiB | [hunspell_cs_cz_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hunspell_cs_cz_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `hunspell_cs_cz_13` | 1.0 | `el9.x86_64` | pigsty | 528.8 KiB | [hunspell_cs_cz_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hunspell_cs_cz_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-hunspell-cs-cz` | 1.0 | `d12.aarch64` | pigsty | 503.5 KiB | [postgresql-13-hunspell-cs-cz_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-cs-cz/postgresql-13-hunspell-cs-cz_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-hunspell-cs-cz` | 1.0 | `d12.x86_64` | pigsty | 503.5 KiB | [postgresql-13-hunspell-cs-cz_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hunspell-cs-cz/postgresql-13-hunspell-cs-cz_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-hunspell-cs-cz` | 1.0 | `u22.aarch64` | pigsty | 521.7 KiB | [postgresql-13-hunspell-cs-cz_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-cs-cz/postgresql-13-hunspell-cs-cz_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-hunspell-cs-cz` | 1.0 | `u22.x86_64` | pigsty | 521.7 KiB | [postgresql-13-hunspell-cs-cz_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hunspell-cs-cz/postgresql-13-hunspell-cs-cz_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-hunspell-cs-cz` | 1.0 | `u24.aarch64` | pigsty | 521.3 KiB | [postgresql-13-hunspell-cs-cz_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-cs-cz/postgresql-13-hunspell-cs-cz_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-hunspell-cs-cz` | 1.0 | `u24.x86_64` | pigsty | 521.3 KiB | [postgresql-13-hunspell-cs-cz_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hunspell-cs-cz/postgresql-13-hunspell-cs-cz_1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/postgrespro/hunspell_dicts" title="Repository" icon="github" subtitle="github.com/postgrespro/hunspell_dicts" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="hunspell-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get hunspell_cs_cz; # get hunspell_cs_cz source code
pig build dep hunspell_cs_cz; # install build dependencies
pig build pkg hunspell_cs_cz; # build extension rpm or deb
pig build ext hunspell_cs_cz; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install hunspell_cs_cz; # install by extension name, for the current active PG version
pig ext install hunspell_cs_cz; # install via package alias, for the active PG version
pig ext install hunspell_cs_cz -v 18;   # install for PG 18
pig ext install hunspell_cs_cz -v 17;   # install for PG 17
pig ext install hunspell_cs_cz -v 16;   # install for PG 16
pig ext install hunspell_cs_cz -v 15;   # install for PG 15
pig ext install hunspell_cs_cz -v 14;   # install for PG 14
pig ext install hunspell_cs_cz -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION hunspell_cs_cz;
```

