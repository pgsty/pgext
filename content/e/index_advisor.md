---
title: "index_advisor"
linkTitle: "index_advisor"
description: "Query index advisor"
weight: 2840
categories: ["Feat"]
width: full
---

Query index advisor

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2840** | {{< badge content="index_advisor" link="https://github.com/supabase/index_advisor" >}} | {{< ext "index_advisor" "index_advisor" >}} | `0.2.0` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="-----d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "hypopg" >}} {{< ext "pg_qualstats" >}} {{< ext "powa" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_hint_plan" >}} {{< ext "auto_explain" >}} {{< ext "pg_profile" >}} {{< ext "pg_show_plans" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/index_advisor" >}} | `0.2.0` | {{< badge content="18" color="red" alt="index_advisor_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `index_advisor_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/index_advisor" >}} | `0.2.0` | {{< badge content="18" color="red" alt="postgresql-18-index-advisor" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-index-advisor` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "index_advisor_18" >}}     | {{< pkg "index_advisor_17" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/index_advisor_17-0.2.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "index_advisor_16" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/index_advisor_16-0.2.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "index_advisor_15" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/index_advisor_15-0.2.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "index_advisor_14" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/index_advisor_14-0.2.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "index_advisor_18" >}}     | {{< pkg "index_advisor_17" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/index_advisor_17-0.2.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "index_advisor_16" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/index_advisor_16-0.2.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "index_advisor_15" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/index_advisor_15-0.2.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "index_advisor_14" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/index_advisor_14-0.2.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "index_advisor_18" >}}     | {{< pkg "index_advisor_17" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/index_advisor_17-0.2.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "index_advisor_16" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/index_advisor_16-0.2.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "index_advisor_15" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/index_advisor_15-0.2.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "index_advisor_14" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/index_advisor_14-0.2.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "index_advisor_18" >}}     | {{< pkg "index_advisor_17" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/index_advisor_17-0.2.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "index_advisor_16" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/index_advisor_16-0.2.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "index_advisor_15" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/index_advisor_15-0.2.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "index_advisor_14" "0.2.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/index_advisor_14-0.2.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-index-advisor" >}}     | {{< pkg "postgresql-17-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-index-advisor" >}}     | {{< pkg "postgresql-17-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-index-advisor" >}}     | {{< pkg "postgresql-17-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-index-advisor" >}}     | {{< pkg "postgresql-17-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-index-advisor" >}}     | {{< pkg "postgresql-17-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-index-advisor" >}}     | {{< pkg "postgresql-17-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-index-advisor" "0.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `index_advisor_17` | 0.2.0 | `el8.x86_64` | pigsty | 10.2 KiB | [index_advisor_17-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/index_advisor_17-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `index_advisor_17` | 0.2.0 | `el8.aarch64` | pigsty | 10.2 KiB | [index_advisor_17-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/index_advisor_17-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `index_advisor_17` | 0.2.0 | `el9.aarch64` | pigsty | 10.1 KiB | [index_advisor_17-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/index_advisor_17-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `index_advisor_17` | 0.2.0 | `el9.x86_64` | pigsty | 10.1 KiB | [index_advisor_17-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/index_advisor_17-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-index-advisor` | 0.2.0 | `d12.x86_64` | pigsty | 5.6 KiB | [postgresql-17-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-index-advisor` | 0.2.0 | `d12.aarch64` | pigsty | 5.6 KiB | [postgresql-17-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-index-advisor` | 0.2.0 | `u22.x86_64` | pigsty | 5.6 KiB | [postgresql-17-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-index-advisor` | 0.2.0 | `u22.aarch64` | pigsty | 5.6 KiB | [postgresql-17-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-index-advisor` | 0.2.0 | `u24.x86_64` | pigsty | 5.6 KiB | [postgresql-17-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-index-advisor` | 0.2.0 | `u24.aarch64` | pigsty | 5.6 KiB | [postgresql-17-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `index_advisor_16` | 0.2.0 | `el8.x86_64` | pigsty | 10.2 KiB | [index_advisor_16-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/index_advisor_16-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `index_advisor_16` | 0.2.0 | `el8.aarch64` | pigsty | 10.2 KiB | [index_advisor_16-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/index_advisor_16-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `index_advisor_16` | 0.2.0 | `el9.x86_64` | pigsty | 10.1 KiB | [index_advisor_16-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/index_advisor_16-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `index_advisor_16` | 0.2.0 | `el9.aarch64` | pigsty | 10.1 KiB | [index_advisor_16-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/index_advisor_16-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-index-advisor` | 0.2.0 | `d12.x86_64` | pigsty | 5.6 KiB | [postgresql-16-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-index-advisor` | 0.2.0 | `d12.aarch64` | pigsty | 5.6 KiB | [postgresql-16-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-index-advisor` | 0.2.0 | `u22.aarch64` | pigsty | 5.6 KiB | [postgresql-16-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-index-advisor` | 0.2.0 | `u22.x86_64` | pigsty | 5.6 KiB | [postgresql-16-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-index-advisor` | 0.2.0 | `u24.x86_64` | pigsty | 5.6 KiB | [postgresql-16-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-index-advisor` | 0.2.0 | `u24.aarch64` | pigsty | 5.6 KiB | [postgresql-16-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `index_advisor_15` | 0.2.0 | `el8.x86_64` | pigsty | 10.2 KiB | [index_advisor_15-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/index_advisor_15-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `index_advisor_15` | 0.2.0 | `el8.aarch64` | pigsty | 10.2 KiB | [index_advisor_15-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/index_advisor_15-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `index_advisor_15` | 0.2.0 | `el9.x86_64` | pigsty | 10.1 KiB | [index_advisor_15-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/index_advisor_15-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `index_advisor_15` | 0.2.0 | `el9.aarch64` | pigsty | 10.1 KiB | [index_advisor_15-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/index_advisor_15-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-index-advisor` | 0.2.0 | `d12.aarch64` | pigsty | 5.6 KiB | [postgresql-15-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-index-advisor` | 0.2.0 | `d12.x86_64` | pigsty | 5.6 KiB | [postgresql-15-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-index-advisor` | 0.2.0 | `u22.aarch64` | pigsty | 5.6 KiB | [postgresql-15-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-index-advisor` | 0.2.0 | `u22.x86_64` | pigsty | 5.6 KiB | [postgresql-15-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-index-advisor` | 0.2.0 | `u24.x86_64` | pigsty | 5.6 KiB | [postgresql-15-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-index-advisor` | 0.2.0 | `u24.aarch64` | pigsty | 5.6 KiB | [postgresql-15-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `index_advisor_14` | 0.2.0 | `el8.x86_64` | pigsty | 10.2 KiB | [index_advisor_14-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/index_advisor_14-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `index_advisor_14` | 0.2.0 | `el8.aarch64` | pigsty | 10.2 KiB | [index_advisor_14-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/index_advisor_14-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `index_advisor_14` | 0.2.0 | `el9.x86_64` | pigsty | 10.1 KiB | [index_advisor_14-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/index_advisor_14-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `index_advisor_14` | 0.2.0 | `el9.aarch64` | pigsty | 10.1 KiB | [index_advisor_14-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/index_advisor_14-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-index-advisor` | 0.2.0 | `d12.x86_64` | pigsty | 5.6 KiB | [postgresql-14-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-index-advisor` | 0.2.0 | `d12.aarch64` | pigsty | 5.6 KiB | [postgresql-14-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-index-advisor` | 0.2.0 | `u22.x86_64` | pigsty | 5.6 KiB | [postgresql-14-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-index-advisor` | 0.2.0 | `u22.aarch64` | pigsty | 5.6 KiB | [postgresql-14-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-index-advisor` | 0.2.0 | `u24.x86_64` | pigsty | 5.6 KiB | [postgresql-14-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-index-advisor` | 0.2.0 | `u24.aarch64` | pigsty | 5.6 KiB | [postgresql-14-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `index_advisor_13` | 0.2.0 | `el8.aarch64` | pigsty | 10.2 KiB | [index_advisor_13-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/index_advisor_13-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `index_advisor_13` | 0.2.0 | `el8.x86_64` | pigsty | 10.2 KiB | [index_advisor_13-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/index_advisor_13-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `index_advisor_13` | 0.2.0 | `el9.aarch64` | pigsty | 10.1 KiB | [index_advisor_13-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/index_advisor_13-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `index_advisor_13` | 0.2.0 | `el9.x86_64` | pigsty | 10.1 KiB | [index_advisor_13-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/index_advisor_13-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-index-advisor` | 0.2.0 | `d12.aarch64` | pigsty | 5.6 KiB | [postgresql-13-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-13-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-index-advisor` | 0.2.0 | `d12.x86_64` | pigsty | 5.6 KiB | [postgresql-13-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-13-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-index-advisor` | 0.2.0 | `u22.aarch64` | pigsty | 5.6 KiB | [postgresql-13-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-13-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-index-advisor` | 0.2.0 | `u22.x86_64` | pigsty | 5.6 KiB | [postgresql-13-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-13-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-index-advisor` | 0.2.0 | `u24.aarch64` | pigsty | 5.6 KiB | [postgresql-13-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-13-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-index-advisor` | 0.2.0 | `u24.x86_64` | pigsty | 5.6 KiB | [postgresql-13-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-13-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supabase/index_advisor" title="Repository" icon="github" subtitle="github.com/supabase/index_advisor" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="index_advisor-0.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get index_advisor; # get index_advisor source code
pig build dep index_advisor; # install build dependencies
pig build pkg index_advisor; # build extension rpm or deb
pig build ext index_advisor; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install index_advisor; # install by extension name, for the current active PG version
pig ext install index_advisor; # install via package alias, for the active PG version
pig ext install index_advisor -v 18;   # install for PG 18
pig ext install index_advisor -v 17;   # install for PG 17
pig ext install index_advisor -v 16;   # install for PG 16
pig ext install index_advisor -v 15;   # install for PG 15
pig ext install index_advisor -v 14;   # install for PG 14
pig ext install index_advisor -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION index_advisor;
```

