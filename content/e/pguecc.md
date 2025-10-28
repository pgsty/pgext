---
title: "pguecc"
linkTitle: "pguecc"
description: "uECC bindings for Postgres"
weight: 4460
categories: ["UTIL"]
width: full
---

uECC bindings for Postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4460** | {{< badge content="pguecc" link="https://github.com/ameensol/pg-ecdsa" >}} | {{< ext "pguecc" "pg_ecdsa" >}} | `1.0` | {{< category "UTIL" >}} | {{< license "BSD 2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "hashlib" >}} {{< ext "shacrypt" >}} {{< ext "cryptint" >}} {{< ext "pgcrypto" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pguecc" >}} | `1.0` | {{< bg "18" "pg_ecdsa_18*" "green" >}} {{< bg "17" "pg_ecdsa_17*" "green" >}} {{< bg "16" "pg_ecdsa_16*" "green" >}} {{< bg "15" "pg_ecdsa_15*" "green" >}} {{< bg "14" "pg_ecdsa_14*" "green" >}} {{< bg "13" "pg_ecdsa_13*" "green" >}} | `pg_ecdsa_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pguecc" >}} | `1.0` | {{< bg "18" "postgresql-18-pg-ecdsa" "green" >}} {{< bg "17" "postgresql-17-pg-ecdsa" "green" >}} {{< bg "16" "postgresql-16-pg-ecdsa" "green" >}} {{< bg "15" "postgresql-15-pg-ecdsa" "green" >}} {{< bg "14" "postgresql-14-pg-ecdsa" "green" >}} {{< bg "13" "postgresql-13-pg-ecdsa" "green" >}} | `postgresql-$v-pg-ecdsa` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pg_ecdsa_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "pg_ecdsa_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_ecdsa_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_ecdsa_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_ecdsa_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_ecdsa_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pg_ecdsa_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "pg_ecdsa_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_ecdsa_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_ecdsa_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_ecdsa_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_ecdsa_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pg_ecdsa_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "pg_ecdsa_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_ecdsa_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_ecdsa_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_ecdsa_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_ecdsa_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pg_ecdsa_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "pg_ecdsa_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_ecdsa_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_ecdsa_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_ecdsa_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_ecdsa_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    |      {{< bg "MISS" "pg_ecdsa_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ecdsa_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ecdsa_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ecdsa_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ecdsa_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ecdsa_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "pg_ecdsa_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ecdsa_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ecdsa_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ecdsa_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ecdsa_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ecdsa_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-ecdsa : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-ecdsa : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-ecdsa : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-ecdsa : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-ecdsa : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-ecdsa : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-ecdsa : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-ecdsa : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-ecdsa : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-ecdsa : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-ecdsa : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-ecdsa : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-ecdsa : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-ecdsa : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-ecdsa : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-ecdsa : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-ecdsa : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-ecdsa : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-ecdsa : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-ecdsa : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-ecdsa : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-ecdsa : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-ecdsa : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-ecdsa : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-ecdsa : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ecdsa_17` | 1.0 | `el8.x86_64` | pigsty | 26.9 KiB | [pg_ecdsa_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ecdsa_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_ecdsa_17` | 1.0 | `el8.aarch64` | pigsty | 26.3 KiB | [pg_ecdsa_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ecdsa_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_ecdsa_17` | 1.0 | `el9.x86_64` | pigsty | 25.5 KiB | [pg_ecdsa_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ecdsa_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_ecdsa_17` | 1.0 | `el9.aarch64` | pigsty | 24.8 KiB | [pg_ecdsa_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ecdsa_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pg-ecdsa` | 1.0 | `d12.x86_64` | pigsty | 67.8 KiB | [postgresql-17-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-17-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-ecdsa` | 1.0 | `d12.aarch64` | pigsty | 66.4 KiB | [postgresql-17-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-17-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-ecdsa` | 1.0 | `u22.x86_64` | pigsty | 69.4 KiB | [postgresql-17-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-17-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-ecdsa` | 1.0 | `u22.aarch64` | pigsty | 68.3 KiB | [postgresql-17-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-17-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-ecdsa` | 1.0 | `u24.x86_64` | pigsty | 66.4 KiB | [postgresql-17-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-17-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-ecdsa` | 1.0 | `u24.aarch64` | pigsty | 65.0 KiB | [postgresql-17-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-17-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ecdsa_16` | 1.0 | `el8.x86_64` | pigsty | 26.9 KiB | [pg_ecdsa_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ecdsa_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_ecdsa_16` | 1.0 | `el8.aarch64` | pigsty | 26.3 KiB | [pg_ecdsa_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ecdsa_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_ecdsa_16` | 1.0 | `el9.x86_64` | pigsty | 25.5 KiB | [pg_ecdsa_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ecdsa_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_ecdsa_16` | 1.0 | `el9.aarch64` | pigsty | 24.9 KiB | [pg_ecdsa_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ecdsa_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-ecdsa` | 1.0 | `d12.x86_64` | pigsty | 67.8 KiB | [postgresql-16-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-16-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-ecdsa` | 1.0 | `d12.aarch64` | pigsty | 66.4 KiB | [postgresql-16-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-16-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-ecdsa` | 1.0 | `u22.x86_64` | pigsty | 69.4 KiB | [postgresql-16-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-16-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-ecdsa` | 1.0 | `u22.aarch64` | pigsty | 68.3 KiB | [postgresql-16-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-16-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-ecdsa` | 1.0 | `u24.x86_64` | pigsty | 66.4 KiB | [postgresql-16-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-16-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-ecdsa` | 1.0 | `u24.aarch64` | pigsty | 65.0 KiB | [postgresql-16-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-16-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ecdsa_15` | 1.0 | `el8.x86_64` | pigsty | 27.2 KiB | [pg_ecdsa_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ecdsa_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_ecdsa_15` | 1.0 | `el8.aarch64` | pigsty | 26.6 KiB | [pg_ecdsa_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ecdsa_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_ecdsa_15` | 1.0 | `el9.x86_64` | pigsty | 27.9 KiB | [pg_ecdsa_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ecdsa_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_ecdsa_15` | 1.0 | `el9.aarch64` | pigsty | 27.3 KiB | [pg_ecdsa_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ecdsa_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-ecdsa` | 1.0 | `d12.x86_64` | pigsty | 68.2 KiB | [postgresql-15-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-15-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-ecdsa` | 1.0 | `d12.aarch64` | pigsty | 66.8 KiB | [postgresql-15-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-15-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-ecdsa` | 1.0 | `u22.x86_64` | pigsty | 71.4 KiB | [postgresql-15-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-15-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-ecdsa` | 1.0 | `u22.aarch64` | pigsty | 70.4 KiB | [postgresql-15-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-15-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-ecdsa` | 1.0 | `u24.x86_64` | pigsty | 68.5 KiB | [postgresql-15-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-15-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-ecdsa` | 1.0 | `u24.aarch64` | pigsty | 67.4 KiB | [postgresql-15-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-15-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ecdsa_14` | 1.0 | `el8.x86_64` | pigsty | 27.2 KiB | [pg_ecdsa_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ecdsa_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_ecdsa_14` | 1.0 | `el8.aarch64` | pigsty | 26.6 KiB | [pg_ecdsa_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ecdsa_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_ecdsa_14` | 1.0 | `el9.x86_64` | pigsty | 27.9 KiB | [pg_ecdsa_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ecdsa_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_ecdsa_14` | 1.0 | `el9.aarch64` | pigsty | 27.3 KiB | [pg_ecdsa_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ecdsa_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-ecdsa` | 1.0 | `d12.x86_64` | pigsty | 68.1 KiB | [postgresql-14-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-14-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-ecdsa` | 1.0 | `d12.aarch64` | pigsty | 66.8 KiB | [postgresql-14-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-14-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-ecdsa` | 1.0 | `u22.x86_64` | pigsty | 71.4 KiB | [postgresql-14-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-14-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-ecdsa` | 1.0 | `u22.aarch64` | pigsty | 70.4 KiB | [postgresql-14-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-14-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-ecdsa` | 1.0 | `u24.x86_64` | pigsty | 68.4 KiB | [postgresql-14-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-14-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-ecdsa` | 1.0 | `u24.aarch64` | pigsty | 67.4 KiB | [postgresql-14-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-14-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ecdsa_13` | 1.0 | `el8.x86_64` | pigsty | 27.1 KiB | [pg_ecdsa_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ecdsa_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_ecdsa_13` | 1.0 | `el8.aarch64` | pigsty | 26.6 KiB | [pg_ecdsa_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ecdsa_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_ecdsa_13` | 1.0 | `el9.x86_64` | pigsty | 27.8 KiB | [pg_ecdsa_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ecdsa_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_ecdsa_13` | 1.0 | `el9.aarch64` | pigsty | 27.3 KiB | [pg_ecdsa_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ecdsa_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-pg-ecdsa` | 1.0 | `d12.x86_64` | pigsty | 68.1 KiB | [postgresql-13-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-13-pg-ecdsa_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-ecdsa` | 1.0 | `d12.aarch64` | pigsty | 66.7 KiB | [postgresql-13-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ecdsa/postgresql-13-pg-ecdsa_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-ecdsa` | 1.0 | `u22.x86_64` | pigsty | 71.4 KiB | [postgresql-13-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-13-pg-ecdsa_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-ecdsa` | 1.0 | `u22.aarch64` | pigsty | 70.4 KiB | [postgresql-13-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ecdsa/postgresql-13-pg-ecdsa_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-ecdsa` | 1.0 | `u24.x86_64` | pigsty | 68.4 KiB | [postgresql-13-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-13-pg-ecdsa_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-ecdsa` | 1.0 | `u24.aarch64` | pigsty | 67.3 KiB | [postgresql-13-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ecdsa/postgresql-13-pg-ecdsa_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ameensol/pg-ecdsa" title="Repository" icon="github" subtitle="github.com/ameensol/pg-ecdsa" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg-ecdsa-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pguecc; # get pguecc source code
pig build dep pguecc; # install build dependencies
pig build pkg pguecc; # build extension rpm or deb
pig build ext pguecc; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pguecc; # install by extension name, for the current active PG version
pig ext install pg_ecdsa; # install via package alias, for the active PG version
pig ext install pguecc -v 18;   # install for PG 18
pig ext install pguecc -v 17;   # install for PG 17
pig ext install pguecc -v 16;   # install for PG 16
pig ext install pguecc -v 15;   # install for PG 15
pig ext install pguecc -v 14;   # install for PG 14
pig ext install pguecc -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pguecc;
```

