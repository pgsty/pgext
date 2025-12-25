---
title: "decoder_raw"
linkTitle: "decoder_raw"
description: "Output plugin for logical replication in Raw SQL format"
weight: 9660
categories: ["ETL"]
width: full
---

[**decoder_raw**](https://github.com/michaelpq/pg_plugins/blob/main/decoder_raw/) : Output plugin for logical replication in Raw SQL format


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9660** | {{< badge content="decoder_raw" link="https://github.com/michaelpq/pg_plugins/blob/main/decoder_raw/" >}} | {{< ext "decoder_raw" >}} | `1.0` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s----" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pglogical" >}} {{< ext "wal2json" >}} {{< ext "decoderbufs" >}} {{< ext "test_decoding" >}} {{< ext "pg_failover_slots" >}} {{< ext "pgactive" >}} {{< ext "wal2mongo" >}} {{< ext "pgoutput" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `decoder_raw` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "decoder_raw_18" "green" >}} {{< bg "17" "decoder_raw_17" "green" >}} {{< bg "16" "decoder_raw_16" "green" >}} {{< bg "15" "decoder_raw_15" "green" >}} {{< bg "14" "decoder_raw_14" "green" >}} {{< bg "13" "decoder_raw_13" "green" >}} | `decoder_raw_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-decoder-raw" "green" >}} {{< bg "17" "postgresql-17-decoder-raw" "green" >}} {{< bg "16" "postgresql-16-decoder-raw" "green" >}} {{< bg "15" "postgresql-15-decoder-raw" "green" >}} {{< bg "14" "postgresql-14-decoder-raw" "green" >}} {{< bg "13" "postgresql-13-decoder-raw" "green" >}} | `postgresql-$v-decoder-raw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "decoder_raw_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "decoder_raw_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "decoder_raw_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "decoder_raw_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "decoder_raw_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "decoder_raw_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "decoder_raw_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "decoder_raw_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-decoder-raw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-decoder-raw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-decoder-raw : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-decoder-raw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-decoder-raw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-decoder-raw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-decoder-raw : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-decoder-raw : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-decoder-raw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-decoder-raw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-decoder-raw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-decoder-raw : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-decoder-raw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-decoder-raw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-decoder-raw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-decoder-raw : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `decoder_raw_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.2 KiB | [decoder_raw_18-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_18-1.0-1PIGSTY.el8.x86_64.rpm) |
| `decoder_raw_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.2 KiB | [decoder_raw_18-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_18-1.0-1PIGSTY.el8.aarch64.rpm) |
| `decoder_raw_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.9 KiB | [decoder_raw_18-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_18-1.0-1PIGSTY.el9.x86_64.rpm) |
| `decoder_raw_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.7 KiB | [decoder_raw_18-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_18-1.0-1PIGSTY.el9.aarch64.rpm) |
| `decoder_raw_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.0 KiB | [decoder_raw_18-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/decoder_raw_18-1.0-1PIGSTY.el10.x86_64.rpm) |
| `decoder_raw_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.9 KiB | [decoder_raw_18-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/decoder_raw_18-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-decoder-raw` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.6 KiB | [postgresql-18-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-18-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-decoder-raw` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.5 KiB | [postgresql-18-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-18-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-decoder-raw` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.6 KiB | [postgresql-18-decoder-raw_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/decoder-raw/postgresql-18-decoder-raw_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-decoder-raw` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.6 KiB | [postgresql-18-decoder-raw_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/decoder-raw/postgresql-18-decoder-raw_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-decoder-raw` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.9 KiB | [postgresql-18-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-18-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-decoder-raw` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.5 KiB | [postgresql-18-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-18-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-decoder-raw` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.3 KiB | [postgresql-18-decoder-raw_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-18-decoder-raw_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-decoder-raw` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.3 KiB | [postgresql-18-decoder-raw_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-18-decoder-raw_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `decoder_raw_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.2 KiB | [decoder_raw_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `decoder_raw_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.3 KiB | [decoder_raw_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `decoder_raw_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.9 KiB | [decoder_raw_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `decoder_raw_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.7 KiB | [decoder_raw_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `decoder_raw_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.0 KiB | [decoder_raw_17-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/decoder_raw_17-1.0-1PIGSTY.el10.x86_64.rpm) |
| `decoder_raw_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.0 KiB | [decoder_raw_17-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/decoder_raw_17-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-decoder-raw` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.5 KiB | [postgresql-17-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-decoder-raw` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.5 KiB | [postgresql-17-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-decoder-raw` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.4 KiB | [postgresql-17-decoder-raw_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-decoder-raw` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.6 KiB | [postgresql-17-decoder-raw_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-decoder-raw` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.1 KiB | [postgresql-17-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-decoder-raw` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 19.7 KiB | [postgresql-17-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-decoder-raw` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.1 KiB | [postgresql-17-decoder-raw_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-decoder-raw` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.3 KiB | [postgresql-17-decoder-raw_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `decoder_raw_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.2 KiB | [decoder_raw_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `decoder_raw_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.3 KiB | [decoder_raw_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `decoder_raw_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.8 KiB | [decoder_raw_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `decoder_raw_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.7 KiB | [decoder_raw_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-decoder-raw` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.5 KiB | [postgresql-16-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-decoder-raw` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.5 KiB | [postgresql-16-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-decoder-raw` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.1 KiB | [postgresql-16-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-decoder-raw` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 19.7 KiB | [postgresql-16-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-decoder-raw` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.1 KiB | [postgresql-16-decoder-raw_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-decoder-raw` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.3 KiB | [postgresql-16-decoder-raw_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `decoder_raw_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.0 KiB | [decoder_raw_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `decoder_raw_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.1 KiB | [decoder_raw_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `decoder_raw_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.7 KiB | [decoder_raw_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `decoder_raw_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.6 KiB | [decoder_raw_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-decoder-raw` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.1 KiB | [postgresql-15-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-decoder-raw` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.2 KiB | [postgresql-15-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-decoder-raw` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 19.7 KiB | [postgresql-15-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-decoder-raw` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 19.3 KiB | [postgresql-15-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-decoder-raw` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.8 KiB | [postgresql-15-decoder-raw_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-decoder-raw` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.0 KiB | [postgresql-15-decoder-raw_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `decoder_raw_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.0 KiB | [decoder_raw_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `decoder_raw_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.1 KiB | [decoder_raw_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `decoder_raw_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.0 KiB | [decoder_raw_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `decoder_raw_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.9 KiB | [decoder_raw_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-decoder-raw` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.8 KiB | [postgresql-14-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-14-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-decoder-raw` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.6 KiB | [postgresql-14-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-14-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-decoder-raw` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.3 KiB | [postgresql-14-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-14-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-decoder-raw` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 19.9 KiB | [postgresql-14-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-14-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-decoder-raw` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.5 KiB | [postgresql-14-decoder-raw_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-14-decoder-raw_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-decoder-raw` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.4 KiB | [postgresql-14-decoder-raw_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-14-decoder-raw_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `decoder_raw_13` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.0 KiB | [decoder_raw_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `decoder_raw_13` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.1 KiB | [decoder_raw_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `decoder_raw_13` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.0 KiB | [decoder_raw_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `decoder_raw_13` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.9 KiB | [decoder_raw_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-decoder-raw` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.6 KiB | [postgresql-13-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-13-decoder-raw_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-decoder-raw` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.4 KiB | [postgresql-13-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-13-decoder-raw_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-decoder-raw` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.0 KiB | [postgresql-13-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-13-decoder-raw_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-decoder-raw` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 19.7 KiB | [postgresql-13-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-13-decoder-raw_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-decoder-raw` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.3 KiB | [postgresql-13-decoder-raw_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-13-decoder-raw_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-decoder-raw` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.2 KiB | [postgresql-13-decoder-raw_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-13-decoder-raw_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/michaelpq/pg_plugins/blob/main/decoder_raw/" title="Repository" icon="github" subtitle="github.com/michaelpq/pg_plugins/blob/main/decoder_raw/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="decoder_raw-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg decoder_raw;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install decoder_raw;		# install via package name, for the active PG version

pig install decoder_raw -v 18;   # install for PG 18
pig install decoder_raw -v 17;   # install for PG 17
pig install decoder_raw -v 16;   # install for PG 16
pig install decoder_raw -v 15;   # install for PG 15
pig install decoder_raw -v 14;   # install for PG 14
pig install decoder_raw -v 13;   # install for PG 13

```


This extension does not need `CREATE EXTENSION` DDL command


