---
title: "pg_cardano"
linkTitle: "pg_cardano"
description: "A suite of Cardano-related tools"
weight: 2930
categories: ["FEAT"]
width: full
---

A suite of Cardano-related tools


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2930** | {{< badge content="pg_cardano" link="https://github.com/Fell-x27/pg_cardano" >}} | {{< ext "pg_cardano" >}} | `1.0.5` | {{< category "FEAT" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "age" >}} {{< ext "hll" >}} {{< ext "rum" >}} {{< ext "pg_graphql" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_hint_plan" >}} {{< ext "hypopg" >}} |

> [!Note] pgrx=0.14.1 (0.12)


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_cardano" >}} | `1.0.5` | {{< bg "18" "pg_cardano_18" "red" >}} {{< bg "17" "pg_cardano_17" "green" >}} {{< bg "16" "pg_cardano_16" "green" >}} {{< bg "15" "pg_cardano_15" "green" >}} {{< bg "14" "pg_cardano_14" "green" >}} | `pg_cardano_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_cardano" >}} | `1.0.5` | {{< bg "18" "postgresql-18-pg-cardano" "red" >}} {{< bg "17" "postgresql-17-pg-cardano" "green" >}} {{< bg "16" "postgresql-16-pg-cardano" "green" >}} {{< bg "15" "postgresql-15-pg-cardano" "green" >}} {{< bg "14" "postgresql-14-pg-cardano" "green" >}} | `postgresql-$v-pg-cardano` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pg_cardano_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.5" "pg_cardano_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_cardano_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_cardano_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_cardano_14 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pg_cardano_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.5" "pg_cardano_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_cardano_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_cardano_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_cardano_14 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pg_cardano_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.5" "pg_cardano_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_cardano_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_cardano_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_cardano_14 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pg_cardano_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.5" "pg_cardano_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_cardano_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_cardano_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pg_cardano_14 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-cardano : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.5" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-cardano : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.5" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-cardano : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.5" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-cardano : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.5" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-cardano : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.5" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-cardano : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.5" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cardano_17` | 1.0.5 | `el8.x86_64` | pigsty | 395.2 KiB | [pg_cardano_17-1.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cardano_17-1.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_cardano_17` | 1.0.5 | `el8.aarch64` | pigsty | 349.3 KiB | [pg_cardano_17-1.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cardano_17-1.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_cardano_17` | 1.0.5 | `el9.x86_64` | pigsty | 402.0 KiB | [pg_cardano_17-1.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cardano_17-1.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_cardano_17` | 1.0.5 | `el9.aarch64` | pigsty | 372.5 KiB | [pg_cardano_17-1.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cardano_17-1.0.5-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pg-cardano` | 1.0.5 | `d12.x86_64` | pigsty | 334.6 KiB | [postgresql-17-pg-cardano_1.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-cardano` | 1.0.5 | `d12.aarch64` | pigsty | 285.8 KiB | [postgresql-17-pg-cardano_1.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-cardano` | 1.0.5 | `u22.x86_64` | pigsty | 364.9 KiB | [postgresql-17-pg-cardano_1.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-cardano` | 1.0.5 | `u22.aarch64` | pigsty | 328.4 KiB | [postgresql-17-pg-cardano_1.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-cardano` | 1.0.5 | `u24.x86_64` | pigsty | 360.6 KiB | [postgresql-17-pg-cardano_1.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-cardano` | 1.0.5 | `u24.aarch64` | pigsty | 324.0 KiB | [postgresql-17-pg-cardano_1.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cardano_16` | 1.0.5 | `el8.x86_64` | pigsty | 395.2 KiB | [pg_cardano_16-1.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cardano_16-1.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_cardano_16` | 1.0.5 | `el8.aarch64` | pigsty | 348.9 KiB | [pg_cardano_16-1.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cardano_16-1.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_cardano_16` | 1.0.5 | `el9.x86_64` | pigsty | 402.0 KiB | [pg_cardano_16-1.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cardano_16-1.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_cardano_16` | 1.0.5 | `el9.aarch64` | pigsty | 372.2 KiB | [pg_cardano_16-1.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cardano_16-1.0.5-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-cardano` | 1.0.5 | `d12.x86_64` | pigsty | 334.9 KiB | [postgresql-16-pg-cardano_1.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-cardano` | 1.0.5 | `d12.aarch64` | pigsty | 285.8 KiB | [postgresql-16-pg-cardano_1.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-cardano` | 1.0.5 | `u22.x86_64` | pigsty | 365.0 KiB | [postgresql-16-pg-cardano_1.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-cardano` | 1.0.5 | `u22.aarch64` | pigsty | 328.3 KiB | [postgresql-16-pg-cardano_1.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-cardano` | 1.0.5 | `u24.x86_64` | pigsty | 360.7 KiB | [postgresql-16-pg-cardano_1.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-cardano` | 1.0.5 | `u24.aarch64` | pigsty | 323.8 KiB | [postgresql-16-pg-cardano_1.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cardano_15` | 1.0.5 | `el8.x86_64` | pigsty | 394.5 KiB | [pg_cardano_15-1.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cardano_15-1.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_cardano_15` | 1.0.5 | `el8.aarch64` | pigsty | 349.4 KiB | [pg_cardano_15-1.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cardano_15-1.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_cardano_15` | 1.0.5 | `el9.x86_64` | pigsty | 401.5 KiB | [pg_cardano_15-1.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cardano_15-1.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_cardano_15` | 1.0.5 | `el9.aarch64` | pigsty | 372.5 KiB | [pg_cardano_15-1.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cardano_15-1.0.5-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-cardano` | 1.0.5 | `d12.x86_64` | pigsty | 334.5 KiB | [postgresql-15-pg-cardano_1.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-cardano` | 1.0.5 | `d12.aarch64` | pigsty | 285.7 KiB | [postgresql-15-pg-cardano_1.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-cardano` | 1.0.5 | `u22.x86_64` | pigsty | 364.6 KiB | [postgresql-15-pg-cardano_1.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-cardano` | 1.0.5 | `u22.aarch64` | pigsty | 328.4 KiB | [postgresql-15-pg-cardano_1.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-cardano` | 1.0.5 | `u24.x86_64` | pigsty | 360.4 KiB | [postgresql-15-pg-cardano_1.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-cardano` | 1.0.5 | `u24.aarch64` | pigsty | 324.0 KiB | [postgresql-15-pg-cardano_1.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cardano_14` | 1.0.5 | `el8.x86_64` | pigsty | 395.1 KiB | [pg_cardano_14-1.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cardano_14-1.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_cardano_14` | 1.0.5 | `el8.aarch64` | pigsty | 349.4 KiB | [pg_cardano_14-1.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cardano_14-1.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_cardano_14` | 1.0.5 | `el9.x86_64` | pigsty | 402.6 KiB | [pg_cardano_14-1.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cardano_14-1.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_cardano_14` | 1.0.5 | `el9.aarch64` | pigsty | 372.6 KiB | [pg_cardano_14-1.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cardano_14-1.0.5-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-cardano` | 1.0.5 | `d12.x86_64` | pigsty | 334.9 KiB | [postgresql-14-pg-cardano_1.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-14-pg-cardano_1.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-cardano` | 1.0.5 | `d12.aarch64` | pigsty | 285.7 KiB | [postgresql-14-pg-cardano_1.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-14-pg-cardano_1.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-cardano` | 1.0.5 | `u22.x86_64` | pigsty | 365.1 KiB | [postgresql-14-pg-cardano_1.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-14-pg-cardano_1.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-cardano` | 1.0.5 | `u22.aarch64` | pigsty | 328.4 KiB | [postgresql-14-pg-cardano_1.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-14-pg-cardano_1.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-cardano` | 1.0.5 | `u24.x86_64` | pigsty | 360.5 KiB | [postgresql-14-pg-cardano_1.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-14-pg-cardano_1.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-cardano` | 1.0.5 | `u24.aarch64` | pigsty | 323.9 KiB | [postgresql-14-pg-cardano_1.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-14-pg-cardano_1.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Fell-x27/pg_cardano" title="Repository" icon="github" subtitle="github.com/Fell-x27/pg_cardano" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_cardano-1.0.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_cardano; # get pg_cardano source code
pig build dep pg_cardano; # install build dependencies
pig build pkg pg_cardano; # build extension rpm or deb
pig build ext pg_cardano; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_cardano; # install by extension name, for the current active PG version
pig ext install pg_cardano; # install via package alias, for the active PG version
pig ext install pg_cardano -v 17;   # install for PG 17
pig ext install pg_cardano -v 16;   # install for PG 16
pig ext install pg_cardano -v 15;   # install for PG 15
pig ext install pg_cardano -v 14;   # install for PG 14
pig ext install pg_cardano -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_cardano;
```

