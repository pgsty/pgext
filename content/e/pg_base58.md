---
title: "pg_base58"
linkTitle: "pg_base58"
description: "Base58 Encoder/Decoder Extension for PostgreSQL"
weight: 4830
categories: ["FUNC"]
width: full
---

Base58 Encoder/Decoder Extension for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4830** | {{< badge content="pg_base58" link="https://github.com/Fell-x27/pg_base58" >}} | {{< ext "pg_base58" >}} | `0.0.1` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "url_encode" >}} {{< ext "pg_cardano" >}} {{< ext "base36" >}} {{< ext "base62" >}} {{< ext "pg_polyline" >}} {{< ext "uri" >}} {{< ext "pg_curl" >}} {{< ext "pg_rewrite" >}} |

> [!Note] pgrx=0.12.1


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_base58" >}} | `0.0.1` | {{< bg "18" "pg_base58_18" "red" >}} {{< bg "17" "pg_base58_17" "green" >}} {{< bg "16" "pg_base58_16" "green" >}} {{< bg "15" "pg_base58_15" "green" >}} {{< bg "14" "pg_base58_14" "green" >}} | `pg_base58_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_base58" >}} | `0.0.1` | {{< bg "18" "postgresql-18-pg-base58" "red" >}} {{< bg "17" "postgresql-17-pg-base58" "green" >}} {{< bg "16" "postgresql-16-pg-base58" "green" >}} {{< bg "15" "postgresql-15-pg-base58" "green" >}} {{< bg "14" "postgresql-14-pg-base58" "green" >}} | `postgresql-$v-pg-base58` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pg_base58_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_base58_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_14 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pg_base58_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_base58_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_14 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pg_base58_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_base58_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_14 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pg_base58_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_base58_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_14 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-base58 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-base58 : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-base58 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-base58 : AVAIL 1" "green" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-base58 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-base58 : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-base58 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-base58 : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-base58 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-base58 : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-base58 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-base58 : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base58_17` | 0.0.1 | `el8.x86_64` | pigsty | 197.0 KiB | [pg_base58_17-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base58_17-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_base58_17` | 0.0.1 | `el8.aarch64` | pigsty | 183.7 KiB | [pg_base58_17-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base58_17-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_base58_17` | 0.0.1 | `el9.x86_64` | pigsty | 201.1 KiB | [pg_base58_17-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base58_17-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_base58_17` | 0.0.1 | `el9.aarch64` | pigsty | 197.1 KiB | [pg_base58_17-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base58_17-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pg-base58` | 0.0.1 | `d12.x86_64` | pigsty | 159.2 KiB | [postgresql-17-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-base58` | 0.0.1 | `d12.aarch64` | pigsty | 141.9 KiB | [postgresql-17-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-base58` | 0.0.1 | `u22.x86_64` | pigsty | 173.6 KiB | [postgresql-17-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-base58` | 0.0.1 | `u22.aarch64` | pigsty | 164.6 KiB | [postgresql-17-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-base58` | 0.0.1 | `u24.x86_64` | pigsty | 172.3 KiB | [postgresql-17-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-base58` | 0.0.1 | `u24.aarch64` | pigsty | 163.6 KiB | [postgresql-17-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base58_16` | 0.0.1 | `el8.x86_64` | pigsty | 197.0 KiB | [pg_base58_16-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base58_16-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_base58_16` | 0.0.1 | `el8.aarch64` | pigsty | 183.7 KiB | [pg_base58_16-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base58_16-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_base58_16` | 0.0.1 | `el9.x86_64` | pigsty | 201.2 KiB | [pg_base58_16-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base58_16-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_base58_16` | 0.0.1 | `el9.aarch64` | pigsty | 197.3 KiB | [pg_base58_16-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base58_16-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-base58` | 0.0.1 | `d12.x86_64` | pigsty | 159.2 KiB | [postgresql-16-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-base58` | 0.0.1 | `d12.aarch64` | pigsty | 141.9 KiB | [postgresql-16-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-base58` | 0.0.1 | `u22.x86_64` | pigsty | 173.6 KiB | [postgresql-16-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-base58` | 0.0.1 | `u22.aarch64` | pigsty | 164.6 KiB | [postgresql-16-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-base58` | 0.0.1 | `u24.x86_64` | pigsty | 172.2 KiB | [postgresql-16-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-base58` | 0.0.1 | `u24.aarch64` | pigsty | 163.2 KiB | [postgresql-16-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base58_15` | 0.0.1 | `el8.x86_64` | pigsty | 197.0 KiB | [pg_base58_15-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base58_15-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_base58_15` | 0.0.1 | `el8.aarch64` | pigsty | 183.7 KiB | [pg_base58_15-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base58_15-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_base58_15` | 0.0.1 | `el9.x86_64` | pigsty | 201.1 KiB | [pg_base58_15-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base58_15-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_base58_15` | 0.0.1 | `el9.aarch64` | pigsty | 197.1 KiB | [pg_base58_15-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base58_15-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-base58` | 0.0.1 | `d12.x86_64` | pigsty | 159.2 KiB | [postgresql-15-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-base58` | 0.0.1 | `d12.aarch64` | pigsty | 141.9 KiB | [postgresql-15-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-base58` | 0.0.1 | `u22.x86_64` | pigsty | 173.6 KiB | [postgresql-15-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-base58` | 0.0.1 | `u22.aarch64` | pigsty | 164.6 KiB | [postgresql-15-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-base58` | 0.0.1 | `u24.x86_64` | pigsty | 172.2 KiB | [postgresql-15-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-base58` | 0.0.1 | `u24.aarch64` | pigsty | 163.2 KiB | [postgresql-15-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base58_14` | 0.0.1 | `el8.x86_64` | pigsty | 197.0 KiB | [pg_base58_14-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base58_14-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_base58_14` | 0.0.1 | `el8.aarch64` | pigsty | 183.7 KiB | [pg_base58_14-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base58_14-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_base58_14` | 0.0.1 | `el9.x86_64` | pigsty | 201.2 KiB | [pg_base58_14-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base58_14-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_base58_14` | 0.0.1 | `el9.aarch64` | pigsty | 197.1 KiB | [pg_base58_14-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base58_14-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-base58` | 0.0.1 | `d12.x86_64` | pigsty | 159.2 KiB | [postgresql-14-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-base58` | 0.0.1 | `d12.aarch64` | pigsty | 141.9 KiB | [postgresql-14-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-base58` | 0.0.1 | `u22.x86_64` | pigsty | 173.6 KiB | [postgresql-14-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-base58` | 0.0.1 | `u22.aarch64` | pigsty | 164.4 KiB | [postgresql-14-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-base58` | 0.0.1 | `u24.x86_64` | pigsty | 172.3 KiB | [postgresql-14-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-base58` | 0.0.1 | `u24.aarch64` | pigsty | 163.6 KiB | [postgresql-14-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Fell-x27/pg_base58" title="Repository" icon="github" subtitle="github.com/Fell-x27/pg_base58" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_base58-0.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_base58; # get pg_base58 source code
pig build dep pg_base58; # install build dependencies
pig build pkg pg_base58; # build extension rpm or deb
pig build ext pg_base58; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_base58; # install by extension name, for the current active PG version
pig ext install pg_base58; # install via package alias, for the active PG version
pig ext install pg_base58 -v 17;   # install for PG 17
pig ext install pg_base58 -v 16;   # install for PG 16
pig ext install pg_base58 -v 15;   # install for PG 15
pig ext install pg_base58 -v 14;   # install for PG 14
pig ext install pg_base58 -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_base58;
```

