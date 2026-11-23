---
title: "decoder_raw"
linkTitle: "decoder_raw"
description: "Output plugin for logical replication in Raw SQL format"
weight: 9660
categories: ["ETL"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**decoder_raw**](https://github.com/michaelpq/pg_plugins/blob/main/decoder_raw/) : Output plugin for logical replication in Raw SQL format


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9660** | {{< badge content="decoder_raw" link="https://github.com/michaelpq/pg_plugins/blob/main/decoder_raw/" >}} | {{< ext "decoder_raw" >}} | `1.0` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s----" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "wal2json" >}} {{< ext "decoderbufs" >}} {{< ext "test_decoding" >}} {{< ext "pgoutput" >}} {{< ext "wal2mongo" >}} {{< ext "pgmqtt" >}} {{< ext "kafka_fdw" >}} {{< ext "pgq" >}} {{< ext "pg_protobuf" >}} {{< ext "pgproto" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `decoder_raw` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "19" "decoder_raw_19" "green" >}} {{< bg "18" "decoder_raw_18" "green" >}} {{< bg "17" "decoder_raw_17" "green" >}} {{< bg "16" "decoder_raw_16" "green" >}} {{< bg "15" "decoder_raw_15" "green" >}} | `decoder_raw_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "19" "postgresql-19-decoder-raw" "green" >}} {{< bg "18" "postgresql-18-decoder-raw" "green" >}} {{< bg "17" "postgresql-17-decoder-raw" "green" >}} {{< bg "16" "postgresql-16-decoder-raw" "green" >}} {{< bg "15" "postgresql-15-decoder-raw" "green" >}} | `postgresql-$v-decoder-raw` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "decoder_raw_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_15 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "decoder_raw_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_15 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_15 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_15 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_15 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "decoder_raw_15 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-19-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-decoder-raw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-19-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-decoder-raw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-19-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-decoder-raw : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-19-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-decoder-raw : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-19-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-decoder-raw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-19-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-decoder-raw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-19-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-decoder-raw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-19-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-decoder-raw : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-19-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-decoder-raw : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-19-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-decoder-raw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-decoder-raw : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `decoder_raw_19` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.5 KiB | [decoder_raw_19-1.0-2.git20260728.2271b0dPGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_19-1.0-2.git20260728.2271b0dPGSTY.el9.x86_64.rpm) |
| `decoder_raw_19` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.7 KiB | [decoder_raw_19-1.0-2.git20260728.2271b0dPGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_19-1.0-2.git20260728.2271b0dPGSTY.el9.aarch64.rpm) |
| `decoder_raw_19` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.8 KiB | [decoder_raw_19-1.0-2.git20260728.2271b0dPGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/decoder_raw_19-1.0-2.git20260728.2271b0dPGSTY.el10.x86_64.rpm) |
| `decoder_raw_19` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.6 KiB | [decoder_raw_19-1.0-2.git20260728.2271b0dPGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/decoder_raw_19-1.0-2.git20260728.2271b0dPGSTY.el10.aarch64.rpm) |
| `postgresql-19-decoder-raw` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.8 KiB | [postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_amd64.deb) |
| `postgresql-19-decoder-raw` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.6 KiB | [postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_arm64.deb) |
| `postgresql-19-decoder-raw` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.8 KiB | [postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/decoder-raw/postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_amd64.deb) |
| `postgresql-19-decoder-raw` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.6 KiB | [postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/decoder-raw/postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_arm64.deb) |
| `postgresql-19-decoder-raw` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 19.0 KiB | [postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_amd64.deb) |
| `postgresql-19-decoder-raw` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.6 KiB | [postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_arm64.deb) |
| `postgresql-19-decoder-raw` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.6 KiB | [postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_amd64.deb) |
| `postgresql-19-decoder-raw` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.4 KiB | [postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_arm64.deb) |
| `postgresql-19-decoder-raw` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.6 KiB | [postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/decoder-raw/postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_amd64.deb) |
| `postgresql-19-decoder-raw` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 18.7 KiB | [postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/decoder-raw/postgresql-19-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `decoder_raw_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.6 KiB | [decoder_raw_18-1.0-2.git20260728.2271b0dPGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_18-1.0-2.git20260728.2271b0dPGSTY.el8.x86_64.rpm) |
| `decoder_raw_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.7 KiB | [decoder_raw_18-1.0-2.git20260728.2271b0dPGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_18-1.0-2.git20260728.2271b0dPGSTY.el8.aarch64.rpm) |
| `decoder_raw_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.5 KiB | [decoder_raw_18-1.0-2.git20260728.2271b0dPGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_18-1.0-2.git20260728.2271b0dPGSTY.el9.x86_64.rpm) |
| `decoder_raw_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.7 KiB | [decoder_raw_18-1.0-2.git20260728.2271b0dPGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_18-1.0-2.git20260728.2271b0dPGSTY.el9.aarch64.rpm) |
| `decoder_raw_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.7 KiB | [decoder_raw_18-1.0-2.git20260728.2271b0dPGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/decoder_raw_18-1.0-2.git20260728.2271b0dPGSTY.el10.x86_64.rpm) |
| `decoder_raw_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.5 KiB | [decoder_raw_18-1.0-2.git20260728.2271b0dPGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/decoder_raw_18-1.0-2.git20260728.2271b0dPGSTY.el10.aarch64.rpm) |
| `postgresql-18-decoder-raw` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.8 KiB | [postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-decoder-raw` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.5 KiB | [postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-decoder-raw` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.8 KiB | [postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/decoder-raw/postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-decoder-raw` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.6 KiB | [postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/decoder-raw/postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-decoder-raw` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 19.0 KiB | [postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-decoder-raw` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.6 KiB | [postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-decoder-raw` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.6 KiB | [postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_amd64.deb) |
| `postgresql-18-decoder-raw` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.3 KiB | [postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_arm64.deb) |
| `postgresql-18-decoder-raw` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.7 KiB | [postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/decoder-raw/postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-decoder-raw` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 18.6 KiB | [postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/decoder-raw/postgresql-18-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `decoder_raw_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.4 KiB | [decoder_raw_17-1.0-2.git20260728.2271b0dPGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_17-1.0-2.git20260728.2271b0dPGSTY.el8.x86_64.rpm) |
| `decoder_raw_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.6 KiB | [decoder_raw_17-1.0-2.git20260728.2271b0dPGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_17-1.0-2.git20260728.2271b0dPGSTY.el8.aarch64.rpm) |
| `decoder_raw_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.5 KiB | [decoder_raw_17-1.0-2.git20260728.2271b0dPGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_17-1.0-2.git20260728.2271b0dPGSTY.el9.x86_64.rpm) |
| `decoder_raw_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.7 KiB | [decoder_raw_17-1.0-2.git20260728.2271b0dPGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_17-1.0-2.git20260728.2271b0dPGSTY.el9.aarch64.rpm) |
| `decoder_raw_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.6 KiB | [decoder_raw_17-1.0-2.git20260728.2271b0dPGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/decoder_raw_17-1.0-2.git20260728.2271b0dPGSTY.el10.x86_64.rpm) |
| `decoder_raw_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.5 KiB | [decoder_raw_17-1.0-2.git20260728.2271b0dPGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/decoder_raw_17-1.0-2.git20260728.2271b0dPGSTY.el10.aarch64.rpm) |
| `postgresql-17-decoder-raw` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.7 KiB | [postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-decoder-raw` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.5 KiB | [postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-decoder-raw` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.7 KiB | [postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-decoder-raw` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.5 KiB | [postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-decoder-raw` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.3 KiB | [postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-decoder-raw` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.9 KiB | [postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-decoder-raw` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.5 KiB | [postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_amd64.deb) |
| `postgresql-17-decoder-raw` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.3 KiB | [postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_arm64.deb) |
| `postgresql-17-decoder-raw` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.6 KiB | [postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-decoder-raw` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 18.7 KiB | [postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/decoder-raw/postgresql-17-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `decoder_raw_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.5 KiB | [decoder_raw_16-1.0-2.git20260728.2271b0dPGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_16-1.0-2.git20260728.2271b0dPGSTY.el8.x86_64.rpm) |
| `decoder_raw_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.7 KiB | [decoder_raw_16-1.0-2.git20260728.2271b0dPGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_16-1.0-2.git20260728.2271b0dPGSTY.el8.aarch64.rpm) |
| `decoder_raw_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.5 KiB | [decoder_raw_16-1.0-2.git20260728.2271b0dPGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_16-1.0-2.git20260728.2271b0dPGSTY.el9.x86_64.rpm) |
| `decoder_raw_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.8 KiB | [decoder_raw_16-1.0-2.git20260728.2271b0dPGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_16-1.0-2.git20260728.2271b0dPGSTY.el9.aarch64.rpm) |
| `decoder_raw_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.7 KiB | [decoder_raw_16-1.0-2.git20260728.2271b0dPGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/decoder_raw_16-1.0-2.git20260728.2271b0dPGSTY.el10.x86_64.rpm) |
| `decoder_raw_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.7 KiB | [decoder_raw_16-1.0-2.git20260728.2271b0dPGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/decoder_raw_16-1.0-2.git20260728.2271b0dPGSTY.el10.aarch64.rpm) |
| `postgresql-16-decoder-raw` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.8 KiB | [postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-decoder-raw` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.5 KiB | [postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-decoder-raw` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.8 KiB | [postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-decoder-raw` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.6 KiB | [postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-decoder-raw` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.4 KiB | [postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-decoder-raw` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.0 KiB | [postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-decoder-raw` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.6 KiB | [postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_amd64.deb) |
| `postgresql-16-decoder-raw` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.3 KiB | [postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_arm64.deb) |
| `postgresql-16-decoder-raw` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.7 KiB | [postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-decoder-raw` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 18.7 KiB | [postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/decoder-raw/postgresql-16-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `decoder_raw_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.0 KiB | [decoder_raw_15-1.0-2.git20260728.2271b0dPGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/decoder_raw_15-1.0-2.git20260728.2271b0dPGSTY.el8.x86_64.rpm) |
| `decoder_raw_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.1 KiB | [decoder_raw_15-1.0-2.git20260728.2271b0dPGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/decoder_raw_15-1.0-2.git20260728.2271b0dPGSTY.el8.aarch64.rpm) |
| `decoder_raw_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.0 KiB | [decoder_raw_15-1.0-2.git20260728.2271b0dPGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/decoder_raw_15-1.0-2.git20260728.2271b0dPGSTY.el9.x86_64.rpm) |
| `decoder_raw_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.2 KiB | [decoder_raw_15-1.0-2.git20260728.2271b0dPGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/decoder_raw_15-1.0-2.git20260728.2271b0dPGSTY.el9.aarch64.rpm) |
| `decoder_raw_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.3 KiB | [decoder_raw_15-1.0-2.git20260728.2271b0dPGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/decoder_raw_15-1.0-2.git20260728.2271b0dPGSTY.el10.x86_64.rpm) |
| `decoder_raw_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.3 KiB | [decoder_raw_15-1.0-2.git20260728.2271b0dPGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/decoder_raw_15-1.0-2.git20260728.2271b0dPGSTY.el10.aarch64.rpm) |
| `postgresql-15-decoder-raw` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.5 KiB | [postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-decoder-raw` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.2 KiB | [postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-decoder-raw` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.5 KiB | [postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-decoder-raw` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.3 KiB | [postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-decoder-raw` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.1 KiB | [postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-decoder-raw` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.7 KiB | [postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-decoder-raw` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.2 KiB | [postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_amd64.deb) |
| `postgresql-15-decoder-raw` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.1 KiB | [postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~noble_arm64.deb) |
| `postgresql-15-decoder-raw` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.2 KiB | [postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-decoder-raw` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 18.4 KiB | [postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/decoder-raw/postgresql-15-decoder-raw_1.0+git20260728.2271b0d-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/michaelpq/pg_plugins/blob/main/decoder_raw/" title="Repository" icon="github" subtitle="github.com/michaelpq/pg_plugins/blob/main/decoder_raw/" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="decoder_raw-1.0+git20260728.2271b0d.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg decoder_raw;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install decoder_raw;		# install via package name, for the active PG version

pig install decoder_raw -v 19;   # install for PG 19
pig install decoder_raw -v 18;   # install for PG 18
pig install decoder_raw -v 17;   # install for PG 17
pig install decoder_raw -v 16;   # install for PG 16
pig install decoder_raw -v 15;   # install for PG 15

```


This extension does not need `CREATE EXTENSION` DDL command






## Usage

> [decoder_raw: Output plugin for logical replication in Raw SQL format](https://github.com/michaelpq/pg_plugins/blob/main/decoder_raw/)

A logical decoding output plugin that converts WAL changes into raw SQL statements. Part of the pg_plugins collection by Michael Paquier.

### Configuration

In `postgresql.conf`:

```ini
wal_level = logical
max_replication_slots = 10
max_wal_senders = 10
```

### Using with SQL Functions

```sql
-- Create a logical replication slot
SELECT * FROM pg_create_logical_replication_slot('raw_slot', 'decoder_raw');

-- Perform DML operations
INSERT INTO my_table VALUES (1, 'hello');
UPDATE my_table SET val = 'world' WHERE id = 1;
DELETE FROM my_table WHERE id = 1;

-- Get changes as raw SQL
SELECT data FROM pg_logical_slot_get_changes('raw_slot', NULL, NULL);
-- Output:
-- INSERT INTO public.my_table (id, val) VALUES (1, 'hello');
-- UPDATE public.my_table SET val = 'world' WHERE id = 1;
-- DELETE FROM public.my_table WHERE id = 1;

-- Drop the slot
SELECT pg_drop_replication_slot('raw_slot');
```

### Using with pg_recvlogical

```bash
# Create slot
pg_recvlogical -d postgres --slot raw_slot --create-slot -P decoder_raw

# Stream changes as SQL statements
pg_recvlogical -d postgres --slot raw_slot --start -f -

# Drop slot
pg_recvlogical -d postgres --slot raw_slot --drop-slot
```

### Key Features

- Outputs changes as executable SQL statements (INSERT, UPDATE, DELETE)
- Useful for debugging logical decoding or replaying changes on another database
- Tables should have REPLICA IDENTITY set for proper UPDATE/DELETE output
- Lightweight plugin designed as a template for custom logical decoding plugins
