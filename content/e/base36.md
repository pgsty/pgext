---
title: "base36"
linkTitle: "base36"
description: "Integer Base36 types"
weight: 4800
categories: ["FUNC"]
width: full
---

Integer Base36 types


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4800** | {{< badge content="base36" link="https://github.com/adjust/pg-base36" >}} | {{< ext "base36" "pg_base36" >}} | `1.0.0` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "base62" >}} {{< ext "pg_base58" >}} {{< ext "pg_polyline" >}} {{< ext "uri" >}} {{< ext "pg_curl" >}} {{< ext "url_encode" >}} {{< ext "pg_rewrite" >}} {{< ext "sepgsql" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/base36" >}} | `1.0.0` | {{< bg "18" "pg_base36_18*" "green" >}} {{< bg "17" "pg_base36_17*" "green" >}} {{< bg "16" "pg_base36_16*" "green" >}} {{< bg "15" "pg_base36_15*" "green" >}} {{< bg "14" "pg_base36_14*" "green" >}} {{< bg "13" "pg_base36_13*" "green" >}} | `pg_base36_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/base36" >}} | `1.0.0` | {{< bg "18" "postgresql-18-base36" "green" >}} {{< bg "17" "postgresql-17-base36" "green" >}} {{< bg "16" "postgresql-16-base36" "green" >}} {{< bg "15" "postgresql-15-base36" "green" >}} {{< bg "14" "postgresql-14-base36" "green" >}} {{< bg "13" "postgresql-13-base36" "green" >}} | `postgresql-$v-base36` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 1.0.0" "pg_base36_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 1.0.0" "pg_base36_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 1.0.0" "pg_base36_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 1.0.0" "pg_base36_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    | {{< bg "PIGSTY 1.0.0" "pg_base36_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_13 : AVAIL 1" "green" >}} |
|    `el10.aarch64`    | {{< bg "PIGSTY 1.0.0" "pg_base36_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_base36_13 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-base36 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-base36 : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-base36 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-base36 : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-base36 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-base36 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-base36 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-base36 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-base36 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-base36 : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-base36 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-base36 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-base36 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-base36 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-base36 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-base36 : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-base36 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-base36 : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-base36 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-base36 : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-base36 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-base36 : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-base36 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-base36 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-base36 : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base36_18` | 1.0.0 | `el8.x86_64` | pigsty | 15.1 KiB | [pg_base36_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base36_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_base36_18` | 1.0.0 | `el8.aarch64` | pigsty | 15.1 KiB | [pg_base36_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base36_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_base36_18` | 1.0.0 | `el9.x86_64` | pigsty | 14.4 KiB | [pg_base36_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base36_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_base36_18` | 1.0.0 | `el9.aarch64` | pigsty | 14.4 KiB | [pg_base36_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base36_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_base36_18` | 1.0.0 | `el10.x86_64` | pigsty | 14.5 KiB | [pg_base36_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_base36_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_base36_18` | 1.0.0 | `el10.aarch64` | pigsty | 14.7 KiB | [pg_base36_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_base36_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base36_17` | 1.0.0 | `el8.x86_64` | pigsty | 15.1 KiB | [pg_base36_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base36_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_base36_17` | 1.0.0 | `el8.aarch64` | pigsty | 15.1 KiB | [pg_base36_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base36_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_base36_17` | 1.0.0 | `el9.x86_64` | pigsty | 14.4 KiB | [pg_base36_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base36_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_base36_17` | 1.0.0 | `el9.aarch64` | pigsty | 14.4 KiB | [pg_base36_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base36_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_base36_17` | 1.0.0 | `el10.x86_64` | pigsty | 14.6 KiB | [pg_base36_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_base36_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_base36_17` | 1.0.0 | `el10.aarch64` | pigsty | 14.7 KiB | [pg_base36_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_base36_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-base36` | 1.0.0 | `d12.x86_64` | pigsty | 15.2 KiB | [postgresql-17-base36_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base36/postgresql-17-base36_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-base36` | 1.0.0 | `d12.aarch64` | pigsty | 15.2 KiB | [postgresql-17-base36_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base36/postgresql-17-base36_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-base36` | 1.0.0 | `u22.x86_64` | pigsty | 16.0 KiB | [postgresql-17-base36_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base36/postgresql-17-base36_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-base36` | 1.0.0 | `u22.aarch64` | pigsty | 15.8 KiB | [postgresql-17-base36_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base36/postgresql-17-base36_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-base36` | 1.0.0 | `u24.x86_64` | pigsty | 15.7 KiB | [postgresql-17-base36_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base36/postgresql-17-base36_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-base36` | 1.0.0 | `u24.aarch64` | pigsty | 15.6 KiB | [postgresql-17-base36_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base36/postgresql-17-base36_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base36_16` | 1.0.0 | `el8.x86_64` | pigsty | 15.1 KiB | [pg_base36_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base36_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_base36_16` | 1.0.0 | `el8.aarch64` | pigsty | 15.1 KiB | [pg_base36_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base36_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_base36_16` | 1.0.0 | `el9.x86_64` | pigsty | 14.4 KiB | [pg_base36_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base36_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_base36_16` | 1.0.0 | `el9.aarch64` | pigsty | 14.4 KiB | [pg_base36_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base36_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_base36_16` | 1.0.0 | `el10.x86_64` | pigsty | 14.6 KiB | [pg_base36_16-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_base36_16-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_base36_16` | 1.0.0 | `el10.aarch64` | pigsty | 14.7 KiB | [pg_base36_16-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_base36_16-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-base36` | 1.0.0 | `d12.x86_64` | pigsty | 15.2 KiB | [postgresql-16-base36_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base36/postgresql-16-base36_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-base36` | 1.0.0 | `d12.aarch64` | pigsty | 15.2 KiB | [postgresql-16-base36_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base36/postgresql-16-base36_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-base36` | 1.0.0 | `u22.x86_64` | pigsty | 16.0 KiB | [postgresql-16-base36_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base36/postgresql-16-base36_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-base36` | 1.0.0 | `u22.aarch64` | pigsty | 15.8 KiB | [postgresql-16-base36_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base36/postgresql-16-base36_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-base36` | 1.0.0 | `u24.x86_64` | pigsty | 15.7 KiB | [postgresql-16-base36_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base36/postgresql-16-base36_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-base36` | 1.0.0 | `u24.aarch64` | pigsty | 15.6 KiB | [postgresql-16-base36_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base36/postgresql-16-base36_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base36_15` | 1.0.0 | `el8.x86_64` | pigsty | 15.1 KiB | [pg_base36_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base36_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_base36_15` | 1.0.0 | `el8.aarch64` | pigsty | 15.1 KiB | [pg_base36_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base36_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_base36_15` | 1.0.0 | `el9.x86_64` | pigsty | 14.4 KiB | [pg_base36_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base36_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_base36_15` | 1.0.0 | `el9.aarch64` | pigsty | 14.4 KiB | [pg_base36_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base36_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_base36_15` | 1.0.0 | `el10.x86_64` | pigsty | 14.5 KiB | [pg_base36_15-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_base36_15-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_base36_15` | 1.0.0 | `el10.aarch64` | pigsty | 14.7 KiB | [pg_base36_15-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_base36_15-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-base36` | 1.0.0 | `d12.x86_64` | pigsty | 15.1 KiB | [postgresql-15-base36_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base36/postgresql-15-base36_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-base36` | 1.0.0 | `d12.aarch64` | pigsty | 15.2 KiB | [postgresql-15-base36_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base36/postgresql-15-base36_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-base36` | 1.0.0 | `u22.x86_64` | pigsty | 16.0 KiB | [postgresql-15-base36_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base36/postgresql-15-base36_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-base36` | 1.0.0 | `u22.aarch64` | pigsty | 15.8 KiB | [postgresql-15-base36_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base36/postgresql-15-base36_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-base36` | 1.0.0 | `u24.x86_64` | pigsty | 15.7 KiB | [postgresql-15-base36_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base36/postgresql-15-base36_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-base36` | 1.0.0 | `u24.aarch64` | pigsty | 15.7 KiB | [postgresql-15-base36_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base36/postgresql-15-base36_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base36_14` | 1.0.0 | `el8.x86_64` | pigsty | 15.1 KiB | [pg_base36_14-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base36_14-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_base36_14` | 1.0.0 | `el8.aarch64` | pigsty | 15.1 KiB | [pg_base36_14-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base36_14-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_base36_14` | 1.0.0 | `el9.x86_64` | pigsty | 14.3 KiB | [pg_base36_14-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base36_14-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_base36_14` | 1.0.0 | `el9.aarch64` | pigsty | 14.4 KiB | [pg_base36_14-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base36_14-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_base36_14` | 1.0.0 | `el10.x86_64` | pigsty | 14.5 KiB | [pg_base36_14-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_base36_14-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_base36_14` | 1.0.0 | `el10.aarch64` | pigsty | 14.7 KiB | [pg_base36_14-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_base36_14-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-base36` | 1.0.0 | `d12.x86_64` | pigsty | 15.1 KiB | [postgresql-14-base36_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base36/postgresql-14-base36_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-base36` | 1.0.0 | `d12.aarch64` | pigsty | 15.2 KiB | [postgresql-14-base36_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base36/postgresql-14-base36_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-base36` | 1.0.0 | `u22.x86_64` | pigsty | 15.9 KiB | [postgresql-14-base36_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base36/postgresql-14-base36_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-base36` | 1.0.0 | `u22.aarch64` | pigsty | 15.8 KiB | [postgresql-14-base36_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base36/postgresql-14-base36_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-base36` | 1.0.0 | `u24.x86_64` | pigsty | 15.7 KiB | [postgresql-14-base36_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base36/postgresql-14-base36_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-base36` | 1.0.0 | `u24.aarch64` | pigsty | 15.6 KiB | [postgresql-14-base36_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base36/postgresql-14-base36_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base36_13` | 1.0.0 | `el8.x86_64` | pigsty | 15.0 KiB | [pg_base36_13-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base36_13-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_base36_13` | 1.0.0 | `el8.aarch64` | pigsty | 15.0 KiB | [pg_base36_13-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base36_13-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_base36_13` | 1.0.0 | `el9.x86_64` | pigsty | 14.3 KiB | [pg_base36_13-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base36_13-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_base36_13` | 1.0.0 | `el9.aarch64` | pigsty | 14.4 KiB | [pg_base36_13-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base36_13-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_base36_13` | 1.0.0 | `el10.x86_64` | pigsty | 14.5 KiB | [pg_base36_13-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_base36_13-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_base36_13` | 1.0.0 | `el10.aarch64` | pigsty | 14.7 KiB | [pg_base36_13-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_base36_13-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-base36` | 1.0.0 | `d12.x86_64` | pigsty | 14.8 KiB | [postgresql-13-base36_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base36/postgresql-13-base36_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-base36` | 1.0.0 | `d12.aarch64` | pigsty | 14.9 KiB | [postgresql-13-base36_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base36/postgresql-13-base36_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-base36` | 1.0.0 | `u22.x86_64` | pigsty | 15.7 KiB | [postgresql-13-base36_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base36/postgresql-13-base36_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-base36` | 1.0.0 | `u22.aarch64` | pigsty | 15.6 KiB | [postgresql-13-base36_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base36/postgresql-13-base36_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-base36` | 1.0.0 | `u24.x86_64` | pigsty | 15.4 KiB | [postgresql-13-base36_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base36/postgresql-13-base36_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-base36` | 1.0.0 | `u24.aarch64` | pigsty | 15.4 KiB | [postgresql-13-base36_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base36/postgresql-13-base36_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/adjust/pg-base36" title="Repository" icon="github" subtitle="github.com/adjust/pg-base36" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg-base36-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get base36; # get base36 source code
pig build dep base36; # install build dependencies
pig build pkg base36; # build extension rpm or deb
pig build ext base36; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install base36; # install by extension name, for the current active PG version
pig ext install pg_base36; # install via package alias, for the active PG version
pig ext install base36 -v 18;   # install for PG 18
pig ext install base36 -v 17;   # install for PG 17
pig ext install base36 -v 16;   # install for PG 16
pig ext install base36 -v 15;   # install for PG 15
pig ext install base36 -v 14;   # install for PG 14
pig ext install base36 -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION base36;
```

