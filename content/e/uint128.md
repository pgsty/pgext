---
title: "uint128"
linkTitle: "uint128"
description: "Native uint128 type"
weight: 3740
categories: ["TYPE"]
width: full
---

Native uint128 type


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3740** | {{< badge content="uint128" link="https://github.com/pg-uint/pg-uint128" >}} | {{< ext "uint128" "pg_uint128" >}} | `1.1.1` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |

> [!Note] breaks on el8 since 1.1 ,fix el8 build problem by adding __has_builtin marco


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/uint128" >}} | `1.1.1` | {{< bg "18" "pg_uint128_18*" "green" >}} {{< bg "17" "pg_uint128_17*" "green" >}} {{< bg "16" "pg_uint128_16*" "green" >}} {{< bg "15" "pg_uint128_15*" "green" >}} {{< bg "14" "pg_uint128_14*" "green" >}} {{< bg "13" "pg_uint128_13*" "green" >}} | `pg_uint128_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/uint128" >}} | `1.1.1` | {{< bg "18" "postgresql-18-pg-uint128" "green" >}} {{< bg "17" "postgresql-17-pg-uint128" "green" >}} {{< bg "16" "postgresql-16-pg-uint128" "green" >}} {{< bg "15" "postgresql-15-pg-uint128" "green" >}} {{< bg "14" "postgresql-14-pg-uint128" "green" >}} {{< bg "13" "postgresql-13-pg-uint128" "green" >}} | `postgresql-$v-pg-uint128` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pg_uint128_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "pg_uint128_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uint128_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uint128_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uint128_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uint128_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pg_uint128_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "pg_uint128_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uint128_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uint128_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uint128_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uint128_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pg_uint128_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.0" "pg_uint128_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uint128_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uint128_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uint128_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uint128_13 : AVAIL 2" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pg_uint128_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.0" "pg_uint128_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uint128_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uint128_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uint128_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uint128_13 : AVAIL 2" "green" >}} |
|    `el10.x86_64`    |      {{< bg "MISS" "pg_uint128_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_uint128_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_uint128_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_uint128_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_uint128_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_uint128_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "pg_uint128_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_uint128_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_uint128_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_uint128_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_uint128_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_uint128_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-uint128 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-13-pg-uint128 : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-uint128 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-13-pg-uint128 : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-uint128 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-uint128 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-uint128 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-uint128 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-uint128 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-uint128 : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-uint128 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-uint128 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-uint128 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-uint128 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-uint128 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-uint128 : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-uint128 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-13-pg-uint128 : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-uint128 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-13-pg-uint128 : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-uint128 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-13-pg-uint128 : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-uint128 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-13-pg-uint128 : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uint128_17` | 1.0.0 | `el8.x86_64` | pigsty | 88.3 KiB | [pg_uint128_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uint128_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uint128_17` | 1.0.0 | `el8.aarch64` | pigsty | 85.0 KiB | [pg_uint128_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uint128_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uint128_17` | 1.1.0 | `el9.x86_64` | pigsty | 135.8 KiB | [pg_uint128_17-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_17-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_17` | 1.0.0 | `el9.x86_64` | pigsty | 93.2 KiB | [pg_uint128_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_17` | 1.1.0 | `el9.aarch64` | pigsty | 127.3 KiB | [pg_uint128_17-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_17-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uint128_17` | 1.0.0 | `el9.aarch64` | pigsty | 89.8 KiB | [pg_uint128_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pg-uint128` | 1.1.0 | `d12.x86_64` | pigsty | 310.6 KiB | [postgresql-17-pg-uint128_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-uint128` | 1.1.0 | `d12.aarch64` | pigsty | 300.7 KiB | [postgresql-17-pg-uint128_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-uint128` | 1.1.0 | `u22.x86_64` | pigsty | 358.6 KiB | [postgresql-17-pg-uint128_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-uint128` | 1.1.0 | `u22.aarch64` | pigsty | 350.0 KiB | [postgresql-17-pg-uint128_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-uint128` | 1.1.0 | `u24.x86_64` | pigsty | 336.5 KiB | [postgresql-17-pg-uint128_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-uint128` | 1.1.0 | `u24.aarch64` | pigsty | 334.8 KiB | [postgresql-17-pg-uint128_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uint128_16` | 1.0.0 | `el8.x86_64` | pigsty | 88.4 KiB | [pg_uint128_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uint128_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uint128_16` | 1.0.0 | `el8.aarch64` | pigsty | 85.0 KiB | [pg_uint128_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uint128_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uint128_16` | 1.1.0 | `el9.x86_64` | pigsty | 135.8 KiB | [pg_uint128_16-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_16-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_16` | 1.0.0 | `el9.x86_64` | pigsty | 93.7 KiB | [pg_uint128_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_16` | 1.1.0 | `el9.aarch64` | pigsty | 129.0 KiB | [pg_uint128_16-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_16-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uint128_16` | 1.0.0 | `el9.aarch64` | pigsty | 89.8 KiB | [pg_uint128_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-uint128` | 1.1.0 | `d12.x86_64` | pigsty | 310.6 KiB | [postgresql-16-pg-uint128_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-uint128` | 1.1.0 | `d12.aarch64` | pigsty | 301.7 KiB | [postgresql-16-pg-uint128_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-uint128` | 1.1.0 | `u22.x86_64` | pigsty | 358.6 KiB | [postgresql-16-pg-uint128_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-uint128` | 1.1.0 | `u22.aarch64` | pigsty | 349.9 KiB | [postgresql-16-pg-uint128_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-uint128` | 1.1.0 | `u24.x86_64` | pigsty | 336.6 KiB | [postgresql-16-pg-uint128_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-uint128` | 1.1.0 | `u24.aarch64` | pigsty | 334.8 KiB | [postgresql-16-pg-uint128_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uint128_15` | 1.0.0 | `el8.x86_64` | pigsty | 88.8 KiB | [pg_uint128_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uint128_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uint128_15` | 1.0.0 | `el8.aarch64` | pigsty | 85.2 KiB | [pg_uint128_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uint128_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uint128_15` | 1.1.0 | `el9.x86_64` | pigsty | 136.7 KiB | [pg_uint128_15-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_15-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_15` | 1.0.0 | `el9.x86_64` | pigsty | 93.8 KiB | [pg_uint128_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_15` | 1.1.0 | `el9.aarch64` | pigsty | 128.4 KiB | [pg_uint128_15-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_15-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uint128_15` | 1.0.0 | `el9.aarch64` | pigsty | 90.4 KiB | [pg_uint128_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-uint128` | 1.1.0 | `d12.x86_64` | pigsty | 314.0 KiB | [postgresql-15-pg-uint128_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-uint128` | 1.1.0 | `d12.aarch64` | pigsty | 305.1 KiB | [postgresql-15-pg-uint128_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-uint128` | 1.1.0 | `u22.x86_64` | pigsty | 359.7 KiB | [postgresql-15-pg-uint128_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-uint128` | 1.1.0 | `u22.aarch64` | pigsty | 351.8 KiB | [postgresql-15-pg-uint128_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-uint128` | 1.1.0 | `u24.x86_64` | pigsty | 340.7 KiB | [postgresql-15-pg-uint128_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-uint128` | 1.1.0 | `u24.aarch64` | pigsty | 338.1 KiB | [postgresql-15-pg-uint128_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uint128_14` | 1.0.0 | `el8.x86_64` | pigsty | 88.8 KiB | [pg_uint128_14-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uint128_14-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uint128_14` | 1.0.0 | `el8.aarch64` | pigsty | 85.2 KiB | [pg_uint128_14-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uint128_14-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uint128_14` | 1.1.0 | `el9.x86_64` | pigsty | 136.8 KiB | [pg_uint128_14-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_14-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_14` | 1.0.0 | `el9.x86_64` | pigsty | 93.7 KiB | [pg_uint128_14-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_14-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_14` | 1.1.0 | `el9.aarch64` | pigsty | 128.3 KiB | [pg_uint128_14-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_14-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uint128_14` | 1.0.0 | `el9.aarch64` | pigsty | 91.2 KiB | [pg_uint128_14-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_14-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-uint128` | 1.1.0 | `d12.x86_64` | pigsty | 313.6 KiB | [postgresql-14-pg-uint128_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-14-pg-uint128_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-uint128` | 1.1.0 | `d12.aarch64` | pigsty | 304.3 KiB | [postgresql-14-pg-uint128_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-14-pg-uint128_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-uint128` | 1.1.0 | `u22.x86_64` | pigsty | 359.6 KiB | [postgresql-14-pg-uint128_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-14-pg-uint128_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-uint128` | 1.1.0 | `u22.aarch64` | pigsty | 351.9 KiB | [postgresql-14-pg-uint128_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-14-pg-uint128_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-uint128` | 1.1.0 | `u24.x86_64` | pigsty | 340.8 KiB | [postgresql-14-pg-uint128_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-14-pg-uint128_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-uint128` | 1.1.0 | `u24.aarch64` | pigsty | 338.1 KiB | [postgresql-14-pg-uint128_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-14-pg-uint128_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uint128_13` | 1.0.0 | `el8.x86_64` | pigsty | 87.2 KiB | [pg_uint128_13-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uint128_13-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uint128_13` | 1.0.0 | `el8.aarch64` | pigsty | 85.2 KiB | [pg_uint128_13-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uint128_13-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uint128_13` | 1.1.0 | `el9.x86_64` | pigsty | 136.6 KiB | [pg_uint128_13-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_13-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_13` | 1.0.0 | `el9.x86_64` | pigsty | 94.0 KiB | [pg_uint128_13-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_13-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_13` | 1.1.0 | `el9.aarch64` | pigsty | 128.4 KiB | [pg_uint128_13-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_13-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uint128_13` | 1.0.0 | `el9.aarch64` | pigsty | 90.5 KiB | [pg_uint128_13-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_13-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-pg-uint128` | 1.1.0 | `d12.x86_64` | pigsty | 314.3 KiB | [postgresql-13-pg-uint128_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-13-pg-uint128_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-uint128` | 1.1.0 | `d12.aarch64` | pigsty | 304.8 KiB | [postgresql-13-pg-uint128_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-13-pg-uint128_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-uint128` | 1.1.0 | `u22.x86_64` | pigsty | 359.1 KiB | [postgresql-13-pg-uint128_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-13-pg-uint128_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-uint128` | 1.1.0 | `u22.aarch64` | pigsty | 351.1 KiB | [postgresql-13-pg-uint128_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-13-pg-uint128_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-uint128` | 1.1.0 | `u24.x86_64` | pigsty | 340.7 KiB | [postgresql-13-pg-uint128_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-13-pg-uint128_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-uint128` | 1.1.0 | `u24.aarch64` | pigsty | 337.8 KiB | [postgresql-13-pg-uint128_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-13-pg-uint128_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pg-uint/pg-uint128" title="Repository" icon="github" subtitle="github.com/pg-uint/pg-uint128" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg-uint128-1.1.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get uint128; # get uint128 source code
pig build dep uint128; # install build dependencies
pig build pkg uint128; # build extension rpm or deb
pig build ext uint128; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install uint128; # install by extension name, for the current active PG version
pig ext install pg_uint128; # install via package alias, for the active PG version
pig ext install uint128 -v 18;   # install for PG 18
pig ext install uint128 -v 17;   # install for PG 17
pig ext install uint128 -v 16;   # install for PG 16
pig ext install uint128 -v 15;   # install for PG 15
pig ext install uint128 -v 14;   # install for PG 14
pig ext install uint128 -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION uint128;
```

