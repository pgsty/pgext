---
title: "bzip"
linkTitle: "bzip"
description: "Bzip compression and decompression"
weight: 4020
categories: ["UTIL"]
width: full
---

Bzip compression and decompression


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4020** | {{< badge content="bzip" link="https://github.com/steve-chavez/pg_bzip" >}} | {{< ext "bzip" "pg_bzip" >}} | `1.0.0` | {{< category "UTIL" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "gzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "pg_curl" >}} {{< ext "pgjq" >}} {{< ext "pgjwt" >}} {{< ext "pg_smtp_client" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/bzip" >}} | `1.0.0` | {{< bg "18" "pg_bzip_18*" "green" >}} {{< bg "17" "pg_bzip_17*" "green" >}} {{< bg "16" "pg_bzip_16*" "green" >}} {{< bg "15" "pg_bzip_15*" "green" >}} {{< bg "14" "pg_bzip_14*" "green" >}} {{< bg "13" "pg_bzip_13*" "green" >}} | `pg_bzip_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/bzip" >}} | `1.0.0` | {{< bg "18" "postgresql-18-bzip" "green" >}} {{< bg "17" "postgresql-17-bzip" "green" >}} {{< bg "16" "postgresql-16-bzip" "green" >}} {{< bg "15" "postgresql-15-bzip" "green" >}} {{< bg "14" "postgresql-14-bzip" "green" >}} {{< bg "13" "postgresql-13-bzip" "green" >}} | `postgresql-$v-bzip` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 1.0.0" "pg_bzip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 1.0.0" "pg_bzip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 1.0.0" "pg_bzip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 1.0.0" "pg_bzip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    | {{< bg "PIGSTY 1.0.0" "pg_bzip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_13 : AVAIL 1" "green" >}} |
|    `el10.aarch64`    | {{< bg "PIGSTY 1.0.0" "pg_bzip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_bzip_13 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-bzip : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-bzip : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-bzip : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-bzip : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-bzip : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-bzip : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-bzip : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-bzip : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-bzip : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-bzip : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-bzip : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-bzip : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-bzip : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-bzip : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-bzip : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-bzip : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-bzip : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-bzip : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-bzip : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-bzip : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-bzip : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-bzip : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-bzip : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-bzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-bzip : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bzip_18` | 1.0.0 | `el8.x86_64` | pigsty | 14.6 KiB | [pg_bzip_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bzip_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_bzip_18` | 1.0.0 | `el8.aarch64` | pigsty | 14.7 KiB | [pg_bzip_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bzip_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_bzip_18` | 1.0.0 | `el9.x86_64` | pigsty | 14.6 KiB | [pg_bzip_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bzip_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_bzip_18` | 1.0.0 | `el9.aarch64` | pigsty | 14.4 KiB | [pg_bzip_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bzip_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_bzip_18` | 1.0.0 | `el10.x86_64` | pigsty | 14.5 KiB | [pg_bzip_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bzip_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_bzip_18` | 1.0.0 | `el10.aarch64` | pigsty | 14.6 KiB | [pg_bzip_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bzip_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bzip_17` | 1.0.0 | `el8.x86_64` | pigsty | 14.6 KiB | [pg_bzip_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bzip_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_bzip_17` | 1.0.0 | `el8.aarch64` | pigsty | 14.7 KiB | [pg_bzip_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bzip_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_bzip_17` | 1.0.0 | `el9.x86_64` | pigsty | 14.6 KiB | [pg_bzip_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bzip_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_bzip_17` | 1.0.0 | `el9.aarch64` | pigsty | 14.4 KiB | [pg_bzip_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bzip_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_bzip_17` | 1.0.0 | `el10.x86_64` | pigsty | 14.5 KiB | [pg_bzip_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bzip_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_bzip_17` | 1.0.0 | `el10.aarch64` | pigsty | 14.6 KiB | [pg_bzip_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bzip_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-bzip` | 1.0.0 | `d12.x86_64` | pigsty | 13.6 KiB | [postgresql-17-bzip_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bzip/postgresql-17-bzip_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-bzip` | 1.0.0 | `d12.aarch64` | pigsty | 13.5 KiB | [postgresql-17-bzip_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bzip/postgresql-17-bzip_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-bzip` | 1.0.0 | `u22.x86_64` | pigsty | 14.7 KiB | [postgresql-17-bzip_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bzip/postgresql-17-bzip_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-bzip` | 1.0.0 | `u22.aarch64` | pigsty | 14.5 KiB | [postgresql-17-bzip_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bzip/postgresql-17-bzip_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-bzip` | 1.0.0 | `u24.x86_64` | pigsty | 14.2 KiB | [postgresql-17-bzip_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bzip/postgresql-17-bzip_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-bzip` | 1.0.0 | `u24.aarch64` | pigsty | 14.1 KiB | [postgresql-17-bzip_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bzip/postgresql-17-bzip_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bzip_16` | 1.0.0 | `el8.x86_64` | pigsty | 14.6 KiB | [pg_bzip_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bzip_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_bzip_16` | 1.0.0 | `el8.aarch64` | pigsty | 14.7 KiB | [pg_bzip_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bzip_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_bzip_16` | 1.0.0 | `el9.x86_64` | pigsty | 14.6 KiB | [pg_bzip_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bzip_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_bzip_16` | 1.0.0 | `el9.aarch64` | pigsty | 14.4 KiB | [pg_bzip_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bzip_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_bzip_16` | 1.0.0 | `el10.x86_64` | pigsty | 14.5 KiB | [pg_bzip_16-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bzip_16-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_bzip_16` | 1.0.0 | `el10.aarch64` | pigsty | 14.6 KiB | [pg_bzip_16-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bzip_16-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-bzip` | 1.0.0 | `d12.x86_64` | pigsty | 13.5 KiB | [postgresql-16-bzip_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bzip/postgresql-16-bzip_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-bzip` | 1.0.0 | `d12.aarch64` | pigsty | 13.5 KiB | [postgresql-16-bzip_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bzip/postgresql-16-bzip_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-bzip` | 1.0.0 | `u22.x86_64` | pigsty | 14.7 KiB | [postgresql-16-bzip_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bzip/postgresql-16-bzip_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-bzip` | 1.0.0 | `u22.aarch64` | pigsty | 14.5 KiB | [postgresql-16-bzip_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bzip/postgresql-16-bzip_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-bzip` | 1.0.0 | `u24.x86_64` | pigsty | 14.2 KiB | [postgresql-16-bzip_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bzip/postgresql-16-bzip_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-bzip` | 1.0.0 | `u24.aarch64` | pigsty | 14.1 KiB | [postgresql-16-bzip_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bzip/postgresql-16-bzip_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bzip_15` | 1.0.0 | `el8.x86_64` | pigsty | 14.6 KiB | [pg_bzip_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bzip_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_bzip_15` | 1.0.0 | `el8.aarch64` | pigsty | 14.7 KiB | [pg_bzip_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bzip_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_bzip_15` | 1.0.0 | `el9.x86_64` | pigsty | 14.6 KiB | [pg_bzip_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bzip_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_bzip_15` | 1.0.0 | `el9.aarch64` | pigsty | 14.4 KiB | [pg_bzip_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bzip_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_bzip_15` | 1.0.0 | `el10.x86_64` | pigsty | 14.5 KiB | [pg_bzip_15-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bzip_15-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_bzip_15` | 1.0.0 | `el10.aarch64` | pigsty | 14.6 KiB | [pg_bzip_15-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bzip_15-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-bzip` | 1.0.0 | `d12.x86_64` | pigsty | 13.6 KiB | [postgresql-15-bzip_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bzip/postgresql-15-bzip_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-bzip` | 1.0.0 | `d12.aarch64` | pigsty | 13.5 KiB | [postgresql-15-bzip_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bzip/postgresql-15-bzip_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-bzip` | 1.0.0 | `u22.x86_64` | pigsty | 14.7 KiB | [postgresql-15-bzip_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bzip/postgresql-15-bzip_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-bzip` | 1.0.0 | `u22.aarch64` | pigsty | 14.5 KiB | [postgresql-15-bzip_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bzip/postgresql-15-bzip_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-bzip` | 1.0.0 | `u24.x86_64` | pigsty | 14.2 KiB | [postgresql-15-bzip_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bzip/postgresql-15-bzip_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-bzip` | 1.0.0 | `u24.aarch64` | pigsty | 14.1 KiB | [postgresql-15-bzip_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bzip/postgresql-15-bzip_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bzip_14` | 1.0.0 | `el8.x86_64` | pigsty | 14.6 KiB | [pg_bzip_14-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bzip_14-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_bzip_14` | 1.0.0 | `el8.aarch64` | pigsty | 14.7 KiB | [pg_bzip_14-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bzip_14-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_bzip_14` | 1.0.0 | `el9.x86_64` | pigsty | 14.6 KiB | [pg_bzip_14-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bzip_14-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_bzip_14` | 1.0.0 | `el9.aarch64` | pigsty | 14.4 KiB | [pg_bzip_14-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bzip_14-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_bzip_14` | 1.0.0 | `el10.x86_64` | pigsty | 14.5 KiB | [pg_bzip_14-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bzip_14-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_bzip_14` | 1.0.0 | `el10.aarch64` | pigsty | 14.6 KiB | [pg_bzip_14-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bzip_14-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-bzip` | 1.0.0 | `d12.x86_64` | pigsty | 13.5 KiB | [postgresql-14-bzip_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bzip/postgresql-14-bzip_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-bzip` | 1.0.0 | `d12.aarch64` | pigsty | 13.5 KiB | [postgresql-14-bzip_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bzip/postgresql-14-bzip_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-bzip` | 1.0.0 | `u22.x86_64` | pigsty | 14.8 KiB | [postgresql-14-bzip_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bzip/postgresql-14-bzip_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-bzip` | 1.0.0 | `u22.aarch64` | pigsty | 14.6 KiB | [postgresql-14-bzip_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bzip/postgresql-14-bzip_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-bzip` | 1.0.0 | `u24.x86_64` | pigsty | 14.1 KiB | [postgresql-14-bzip_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bzip/postgresql-14-bzip_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-bzip` | 1.0.0 | `u24.aarch64` | pigsty | 14.0 KiB | [postgresql-14-bzip_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bzip/postgresql-14-bzip_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bzip_13` | 1.0.0 | `el8.x86_64` | pigsty | 14.5 KiB | [pg_bzip_13-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bzip_13-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_bzip_13` | 1.0.0 | `el8.aarch64` | pigsty | 14.7 KiB | [pg_bzip_13-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bzip_13-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_bzip_13` | 1.0.0 | `el9.x86_64` | pigsty | 14.6 KiB | [pg_bzip_13-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bzip_13-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_bzip_13` | 1.0.0 | `el9.aarch64` | pigsty | 14.4 KiB | [pg_bzip_13-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bzip_13-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_bzip_13` | 1.0.0 | `el10.x86_64` | pigsty | 14.5 KiB | [pg_bzip_13-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bzip_13-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_bzip_13` | 1.0.0 | `el10.aarch64` | pigsty | 14.6 KiB | [pg_bzip_13-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bzip_13-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-bzip` | 1.0.0 | `d12.x86_64` | pigsty | 13.4 KiB | [postgresql-13-bzip_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bzip/postgresql-13-bzip_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-bzip` | 1.0.0 | `d12.aarch64` | pigsty | 13.3 KiB | [postgresql-13-bzip_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bzip/postgresql-13-bzip_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-bzip` | 1.0.0 | `u22.x86_64` | pigsty | 14.6 KiB | [postgresql-13-bzip_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bzip/postgresql-13-bzip_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-bzip` | 1.0.0 | `u22.aarch64` | pigsty | 14.6 KiB | [postgresql-13-bzip_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bzip/postgresql-13-bzip_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-bzip` | 1.0.0 | `u24.x86_64` | pigsty | 14.0 KiB | [postgresql-13-bzip_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bzip/postgresql-13-bzip_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-bzip` | 1.0.0 | `u24.aarch64` | pigsty | 13.8 KiB | [postgresql-13-bzip_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bzip/postgresql-13-bzip_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/steve-chavez/pg_bzip" title="Repository" icon="github" subtitle="github.com/steve-chavez/pg_bzip" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_bzip-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get bzip; # get bzip source code
pig build dep bzip; # install build dependencies
pig build pkg bzip; # build extension rpm or deb
pig build ext bzip; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install bzip; # install by extension name, for the current active PG version
pig ext install pg_bzip; # install via package alias, for the active PG version
pig ext install bzip -v 18;   # install for PG 18
pig ext install bzip -v 17;   # install for PG 17
pig ext install bzip -v 16;   # install for PG 16
pig ext install bzip -v 15;   # install for PG 15
pig ext install bzip -v 14;   # install for PG 14
pig ext install bzip -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION bzip;
```

