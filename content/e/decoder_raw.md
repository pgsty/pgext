---
title: "decoder_raw"
linkTitle: "decoder_raw"
description: "Output plugin for logical replication in Raw SQL format"
weight: 9660
categories: ["Etl"]
width: full
---

Output plugin for logical replication in Raw SQL format

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9660** | {{< badge content="decoder_raw" link="https://github.com/michaelpq/pg_plugins/blob/main/decoder_raw/" >}} | {{< ext "decoder_raw" "decoder_raw" >}} | `1.0` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s----" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pglogical" >}} {{< ext "wal2json" >}} {{< ext "decoderbufs" >}} {{< ext "test_decoding" >}} {{< ext "pg_failover_slots" >}} {{< ext "pgactive" >}} {{< ext "wal2mongo" >}} {{< ext "pgoutput" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/decoder_raw" >}} | `1.0` | {{< badge content="18" color="red" alt="decoder_raw_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `decoder_raw_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/decoder_raw" >}} | `1.0` | {{< badge content="18" color="red" alt="postgresql-18-decoder-raw" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-decoder-raw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "decoder_raw_18" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_18-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "decoder_raw_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_17-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "decoder_raw_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_16-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "decoder_raw_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_15-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "decoder_raw_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_14-1.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "decoder_raw_18" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_18-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "decoder_raw_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_17-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "decoder_raw_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_16-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "decoder_raw_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_15-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "decoder_raw_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_14-1.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "decoder_raw_18" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_18-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "decoder_raw_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_17-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "decoder_raw_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_16-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "decoder_raw_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_15-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "decoder_raw_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_14-1.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "decoder_raw_18" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_18-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "decoder_raw_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_17-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "decoder_raw_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_16-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "decoder_raw_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_15-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "decoder_raw_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_14-1.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-decoder-raw" >}}     | {{< pkg "postgresql-17-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-14-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-decoder-raw" >}}     | {{< pkg "postgresql-17-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-14-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-decoder-raw" >}}     | {{< pkg "postgresql-17-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-14-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-decoder-raw" >}}     | {{< pkg "postgresql-17-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-14-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-decoder-raw" >}}     | {{< pkg "postgresql-17-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-14-decoder-raw_1.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-decoder-raw" >}}     | {{< pkg "postgresql-17-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-decoder-raw" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-14-decoder-raw_1.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `decoder_raw_18` | 1.0 | `el8.aarch64` | pigsty | 15.0 KiB | [decoder_raw_18-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_18-1.0-1PIGSTY.el8.aarch64.rpm) |
| `decoder_raw_18` | 1.0 | `el8.x86_64` | pigsty | 15.2 KiB | [decoder_raw_18-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_18-1.0-1PIGSTY.el8.x86_64.rpm) |
| `decoder_raw_18` | 1.0 | `el9.x86_64` | pigsty | 15.0 KiB | [decoder_raw_18-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_18-1.0-1PIGSTY.el9.x86_64.rpm) |
| `decoder_raw_18` | 1.0 | `el9.aarch64` | pigsty | 14.6 KiB | [decoder_raw_18-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_18-1.0-1PIGSTY.el9.aarch64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `decoder_raw_17` | 1.0 | `el8.aarch64` | pigsty | 15.0 KiB | [decoder_raw_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `decoder_raw_17` | 1.0 | `el8.x86_64` | pigsty | 15.2 KiB | [decoder_raw_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `decoder_raw_17` | 1.0 | `el9.aarch64` | pigsty | 14.7 KiB | [decoder_raw_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `decoder_raw_17` | 1.0 | `el9.x86_64` | pigsty | 15.0 KiB | [decoder_raw_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-decoder-raw` | 1.0 | `d12.aarch64` | pigsty | 19.3 KiB | [postgresql-17-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-decoder-raw` | 1.0 | `d12.x86_64` | pigsty | 19.2 KiB | [postgresql-17-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-decoder-raw` | 1.0 | `u22.aarch64` | pigsty | 19.7 KiB | [postgresql-17-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-decoder-raw` | 1.0 | `u22.x86_64` | pigsty | 20.1 KiB | [postgresql-17-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-decoder-raw` | 1.0 | `u24.x86_64` | pigsty | 17.5 KiB | [postgresql-17-decoder-raw_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-decoder-raw` | 1.0 | `u24.aarch64` | pigsty | 17.9 KiB | [postgresql-17-decoder-raw_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `decoder_raw_16` | 1.0 | `el8.aarch64` | pigsty | 15.1 KiB | [decoder_raw_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `decoder_raw_16` | 1.0 | `el8.x86_64` | pigsty | 15.2 KiB | [decoder_raw_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `decoder_raw_16` | 1.0 | `el9.x86_64` | pigsty | 15.0 KiB | [decoder_raw_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `decoder_raw_16` | 1.0 | `el9.aarch64` | pigsty | 14.8 KiB | [decoder_raw_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-decoder-raw` | 1.0 | `d12.x86_64` | pigsty | 19.2 KiB | [postgresql-16-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-decoder-raw` | 1.0 | `d12.aarch64` | pigsty | 19.3 KiB | [postgresql-16-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-decoder-raw` | 1.0 | `u22.aarch64` | pigsty | 19.7 KiB | [postgresql-16-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-decoder-raw` | 1.0 | `u22.x86_64` | pigsty | 20.1 KiB | [postgresql-16-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-decoder-raw` | 1.0 | `u24.aarch64` | pigsty | 17.9 KiB | [postgresql-16-decoder-raw_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-decoder-raw` | 1.0 | `u24.x86_64` | pigsty | 17.5 KiB | [postgresql-16-decoder-raw_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `decoder_raw_15` | 1.0 | `el8.x86_64` | pigsty | 15.0 KiB | [decoder_raw_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `decoder_raw_15` | 1.0 | `el8.aarch64` | pigsty | 14.9 KiB | [decoder_raw_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `decoder_raw_15` | 1.0 | `el9.x86_64` | pigsty | 14.9 KiB | [decoder_raw_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `decoder_raw_15` | 1.0 | `el9.aarch64` | pigsty | 14.6 KiB | [decoder_raw_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-decoder-raw` | 1.0 | `d12.x86_64` | pigsty | 18.8 KiB | [postgresql-15-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-decoder-raw` | 1.0 | `d12.aarch64` | pigsty | 18.9 KiB | [postgresql-15-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-decoder-raw` | 1.0 | `u22.x86_64` | pigsty | 19.7 KiB | [postgresql-15-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-decoder-raw` | 1.0 | `u22.aarch64` | pigsty | 19.3 KiB | [postgresql-15-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-decoder-raw` | 1.0 | `u24.aarch64` | pigsty | 17.5 KiB | [postgresql-15-decoder-raw_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-decoder-raw` | 1.0 | `u24.x86_64` | pigsty | 17.3 KiB | [postgresql-15-decoder-raw_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `decoder_raw_14` | 1.0 | `el8.x86_64` | pigsty | 15.0 KiB | [decoder_raw_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `decoder_raw_14` | 1.0 | `el8.aarch64` | pigsty | 14.9 KiB | [decoder_raw_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `decoder_raw_14` | 1.0 | `el9.aarch64` | pigsty | 14.9 KiB | [decoder_raw_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `decoder_raw_14` | 1.0 | `el9.x86_64` | pigsty | 15.2 KiB | [decoder_raw_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-14-decoder-raw` | 1.0 | `d12.x86_64` | pigsty | 19.5 KiB | [postgresql-14-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-14-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-decoder-raw` | 1.0 | `d12.aarch64` | pigsty | 19.4 KiB | [postgresql-14-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-14-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-decoder-raw` | 1.0 | `u22.x86_64` | pigsty | 20.3 KiB | [postgresql-14-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-14-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-decoder-raw` | 1.0 | `u22.aarch64` | pigsty | 19.8 KiB | [postgresql-14-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-14-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-decoder-raw` | 1.0 | `u24.aarch64` | pigsty | 18.2 KiB | [postgresql-14-decoder-raw_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-14-decoder-raw_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-decoder-raw` | 1.0 | `u24.x86_64` | pigsty | 18.1 KiB | [postgresql-14-decoder-raw_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-14-decoder-raw_1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `decoder_raw_13` | 1.0 | `el8.x86_64` | pigsty | 15.0 KiB | [decoder_raw_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `decoder_raw_13` | 1.0 | `el8.aarch64` | pigsty | 14.9 KiB | [decoder_raw_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `decoder_raw_13` | 1.0 | `el9.aarch64` | pigsty | 14.9 KiB | [decoder_raw_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `decoder_raw_13` | 1.0 | `el9.x86_64` | pigsty | 15.2 KiB | [decoder_raw_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-decoder-raw` | 1.0 | `d12.x86_64` | pigsty | 19.2 KiB | [postgresql-13-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-13-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-decoder-raw` | 1.0 | `d12.aarch64` | pigsty | 19.1 KiB | [postgresql-13-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-13-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-decoder-raw` | 1.0 | `u22.aarch64` | pigsty | 19.7 KiB | [postgresql-13-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-13-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-decoder-raw` | 1.0 | `u22.x86_64` | pigsty | 20.0 KiB | [postgresql-13-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-13-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-decoder-raw` | 1.0 | `u24.x86_64` | pigsty | 17.9 KiB | [postgresql-13-decoder-raw_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-13-decoder-raw_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-decoder-raw` | 1.0 | `u24.aarch64` | pigsty | 17.9 KiB | [postgresql-13-decoder-raw_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-13-decoder-raw_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/michaelpq/pg_plugins/blob/main/decoder_raw/" title="Repository" icon="github" subtitle="github.com/michaelpq/pg_plugins/blob/main/decoder_raw/" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="decoder_raw-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get decoder_raw; # get decoder_raw source code
pig build dep decoder_raw; # install build dependencies
pig build pkg decoder_raw; # build extension rpm or deb
pig build ext decoder_raw; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install decoder_raw; # install by extension name, for the current active PG version
pig ext install decoder_raw; # install via package alias, for the active PG version
pig ext install decoder_raw -v 18;   # install for PG 18
pig ext install decoder_raw -v 17;   # install for PG 17
pig ext install decoder_raw -v 16;   # install for PG 16
pig ext install decoder_raw -v 15;   # install for PG 15
pig ext install decoder_raw -v 14;   # install for PG 14
pig ext install decoder_raw -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION decoder_raw;
```

