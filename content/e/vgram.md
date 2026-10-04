---
title: "vgram"
linkTitle: "vgram"
description: "Variable-length gram GIN indexes and statistics for LIKE/ILIKE search"
weight: 2130
categories: ["FTS"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**vgram**](https://github.com/akorotkov/vgram) : Variable-length gram GIN indexes and statistics for LIKE/ILIKE search


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2130** | {{< badge content="vgram" link="https://github.com/akorotkov/vgram" >}} | {{< ext "vgram" >}} | `1.0` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_trgm" >}} {{< ext "pg_bigm" >}} {{< ext "biscuit" >}} {{< ext "pgroonga" >}} |

> [!Note] Collect corpus-specific q-gram statistics before creating the GIN index; upstream supports PostgreSQL 13+.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `vgram` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "vgram_18" "green" >}} {{< bg "17" "vgram_17" "green" >}} {{< bg "16" "vgram_16" "green" >}} {{< bg "15" "vgram_15" "green" >}} {{< bg "14" "vgram_14" "green" >}} | `vgram_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-vgram" "green" >}} {{< bg "17" "postgresql-17-vgram" "green" >}} {{< bg "16" "postgresql-16-vgram" "green" >}} {{< bg "15" "postgresql-15-vgram" "green" >}} {{< bg "14" "postgresql-14-vgram" "green" >}} | `postgresql-$v-vgram` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "vgram_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "vgram_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "vgram_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "vgram_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "vgram_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "vgram_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "vgram_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-vgram : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-vgram : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-vgram : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-vgram : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-vgram : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-vgram : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-vgram : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-vgram : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-vgram : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-vgram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-vgram : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vgram_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 49.5 KiB | [vgram_18-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vgram_18-1.0-1PGSTY.el8.x86_64.rpm) |
| `vgram_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.0 KiB | [vgram_18-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vgram_18-1.0-1PGSTY.el8.aarch64.rpm) |
| `vgram_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 50.1 KiB | [vgram_18-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vgram_18-1.0-1PGSTY.el9.x86_64.rpm) |
| `vgram_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 49.9 KiB | [vgram_18-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vgram_18-1.0-1PGSTY.el9.aarch64.rpm) |
| `vgram_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 50.5 KiB | [vgram_18-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vgram_18-1.0-1PGSTY.el10.x86_64.rpm) |
| `vgram_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 50.4 KiB | [vgram_18-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vgram_18-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-vgram` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 38.0 KiB | [postgresql-18-vgram_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vgram/postgresql-18-vgram_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-vgram` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 37.7 KiB | [postgresql-18-vgram_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vgram/postgresql-18-vgram_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-vgram` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 38.0 KiB | [postgresql-18-vgram_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vgram/postgresql-18-vgram_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-vgram` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 37.9 KiB | [postgresql-18-vgram_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vgram/postgresql-18-vgram_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-vgram` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 40.5 KiB | [postgresql-18-vgram_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vgram/postgresql-18-vgram_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-vgram` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 40.4 KiB | [postgresql-18-vgram_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vgram/postgresql-18-vgram_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-vgram` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 39.6 KiB | [postgresql-18-vgram_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vgram/postgresql-18-vgram_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-vgram` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 39.6 KiB | [postgresql-18-vgram_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vgram/postgresql-18-vgram_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-vgram` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 39.2 KiB | [postgresql-18-vgram_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vgram/postgresql-18-vgram_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-vgram` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 39.1 KiB | [postgresql-18-vgram_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vgram/postgresql-18-vgram_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vgram_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 49.3 KiB | [vgram_17-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vgram_17-1.0-1PGSTY.el8.x86_64.rpm) |
| `vgram_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.0 KiB | [vgram_17-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vgram_17-1.0-1PGSTY.el8.aarch64.rpm) |
| `vgram_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 49.9 KiB | [vgram_17-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vgram_17-1.0-1PGSTY.el9.x86_64.rpm) |
| `vgram_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 49.9 KiB | [vgram_17-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vgram_17-1.0-1PGSTY.el9.aarch64.rpm) |
| `vgram_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 50.3 KiB | [vgram_17-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vgram_17-1.0-1PGSTY.el10.x86_64.rpm) |
| `vgram_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 50.3 KiB | [vgram_17-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vgram_17-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-vgram` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 37.9 KiB | [postgresql-17-vgram_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vgram/postgresql-17-vgram_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-vgram` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 37.7 KiB | [postgresql-17-vgram_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vgram/postgresql-17-vgram_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-vgram` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 37.9 KiB | [postgresql-17-vgram_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vgram/postgresql-17-vgram_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-vgram` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 37.9 KiB | [postgresql-17-vgram_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vgram/postgresql-17-vgram_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-vgram` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 43.8 KiB | [postgresql-17-vgram_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vgram/postgresql-17-vgram_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-vgram` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 43.7 KiB | [postgresql-17-vgram_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vgram/postgresql-17-vgram_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-vgram` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 39.5 KiB | [postgresql-17-vgram_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vgram/postgresql-17-vgram_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-vgram` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 39.5 KiB | [postgresql-17-vgram_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vgram/postgresql-17-vgram_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-vgram` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 39.1 KiB | [postgresql-17-vgram_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vgram/postgresql-17-vgram_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-vgram` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 39.1 KiB | [postgresql-17-vgram_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vgram/postgresql-17-vgram_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vgram_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 49.4 KiB | [vgram_16-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vgram_16-1.0-1PGSTY.el8.x86_64.rpm) |
| `vgram_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.0 KiB | [vgram_16-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vgram_16-1.0-1PGSTY.el8.aarch64.rpm) |
| `vgram_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 50.0 KiB | [vgram_16-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vgram_16-1.0-1PGSTY.el9.x86_64.rpm) |
| `vgram_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 49.9 KiB | [vgram_16-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vgram_16-1.0-1PGSTY.el9.aarch64.rpm) |
| `vgram_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 50.3 KiB | [vgram_16-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vgram_16-1.0-1PGSTY.el10.x86_64.rpm) |
| `vgram_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 50.4 KiB | [vgram_16-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vgram_16-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-vgram` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 37.9 KiB | [postgresql-16-vgram_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vgram/postgresql-16-vgram_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-vgram` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 37.7 KiB | [postgresql-16-vgram_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vgram/postgresql-16-vgram_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-vgram` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 38.0 KiB | [postgresql-16-vgram_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vgram/postgresql-16-vgram_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-vgram` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 37.9 KiB | [postgresql-16-vgram_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vgram/postgresql-16-vgram_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-vgram` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 43.8 KiB | [postgresql-16-vgram_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vgram/postgresql-16-vgram_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-vgram` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 43.7 KiB | [postgresql-16-vgram_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vgram/postgresql-16-vgram_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-vgram` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 39.5 KiB | [postgresql-16-vgram_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vgram/postgresql-16-vgram_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-vgram` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 39.5 KiB | [postgresql-16-vgram_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vgram/postgresql-16-vgram_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-vgram` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 39.2 KiB | [postgresql-16-vgram_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vgram/postgresql-16-vgram_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-vgram` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 39.2 KiB | [postgresql-16-vgram_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vgram/postgresql-16-vgram_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vgram_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 49.6 KiB | [vgram_15-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vgram_15-1.0-1PGSTY.el8.x86_64.rpm) |
| `vgram_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.1 KiB | [vgram_15-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vgram_15-1.0-1PGSTY.el8.aarch64.rpm) |
| `vgram_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 50.3 KiB | [vgram_15-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vgram_15-1.0-1PGSTY.el9.x86_64.rpm) |
| `vgram_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 50.3 KiB | [vgram_15-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vgram_15-1.0-1PGSTY.el9.aarch64.rpm) |
| `vgram_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 50.8 KiB | [vgram_15-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vgram_15-1.0-1PGSTY.el10.x86_64.rpm) |
| `vgram_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 50.8 KiB | [vgram_15-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vgram_15-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-vgram` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 38.1 KiB | [postgresql-15-vgram_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vgram/postgresql-15-vgram_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-vgram` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 37.9 KiB | [postgresql-15-vgram_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vgram/postgresql-15-vgram_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-vgram` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 38.2 KiB | [postgresql-15-vgram_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vgram/postgresql-15-vgram_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-vgram` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 38.1 KiB | [postgresql-15-vgram_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vgram/postgresql-15-vgram_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-vgram` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 44.0 KiB | [postgresql-15-vgram_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vgram/postgresql-15-vgram_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-vgram` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 44.1 KiB | [postgresql-15-vgram_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vgram/postgresql-15-vgram_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-vgram` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 39.8 KiB | [postgresql-15-vgram_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vgram/postgresql-15-vgram_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-vgram` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 40.0 KiB | [postgresql-15-vgram_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vgram/postgresql-15-vgram_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-vgram` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 39.5 KiB | [postgresql-15-vgram_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vgram/postgresql-15-vgram_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-vgram` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 39.6 KiB | [postgresql-15-vgram_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vgram/postgresql-15-vgram_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vgram_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 49.6 KiB | [vgram_14-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vgram_14-1.0-1PGSTY.el8.x86_64.rpm) |
| `vgram_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.1 KiB | [vgram_14-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vgram_14-1.0-1PGSTY.el8.aarch64.rpm) |
| `vgram_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 50.3 KiB | [vgram_14-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vgram_14-1.0-1PGSTY.el9.x86_64.rpm) |
| `vgram_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 50.3 KiB | [vgram_14-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vgram_14-1.0-1PGSTY.el9.aarch64.rpm) |
| `vgram_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 50.7 KiB | [vgram_14-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vgram_14-1.0-1PGSTY.el10.x86_64.rpm) |
| `vgram_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 50.8 KiB | [vgram_14-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vgram_14-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-vgram` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 41.3 KiB | [postgresql-14-vgram_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vgram/postgresql-14-vgram_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-vgram` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 41.1 KiB | [postgresql-14-vgram_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vgram/postgresql-14-vgram_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-vgram` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 41.4 KiB | [postgresql-14-vgram_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vgram/postgresql-14-vgram_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-vgram` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 41.4 KiB | [postgresql-14-vgram_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vgram/postgresql-14-vgram_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-vgram` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 47.3 KiB | [postgresql-14-vgram_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vgram/postgresql-14-vgram_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-vgram` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 47.3 KiB | [postgresql-14-vgram_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vgram/postgresql-14-vgram_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-vgram` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 43.0 KiB | [postgresql-14-vgram_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vgram/postgresql-14-vgram_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-14-vgram` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 43.1 KiB | [postgresql-14-vgram_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vgram/postgresql-14-vgram_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-14-vgram` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 42.7 KiB | [postgresql-14-vgram_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vgram/postgresql-14-vgram_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-vgram` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 42.7 KiB | [postgresql-14-vgram_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vgram/postgresql-14-vgram_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/akorotkov/vgram" title="Repository" icon="github" subtitle="github.com/akorotkov/vgram" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="vgram-1.0-babae3775c9c.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg vgram;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install vgram;		# install via package name, for the active PG version

pig install vgram -v 18;   # install for PG 18
pig install vgram -v 17;   # install for PG 17
pig install vgram -v 16;   # install for PG 16
pig install vgram -v 15;   # install for PG 15
pig install vgram -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION vgram;
```

## Usage

Sources:

- [Official README](https://github.com/akorotkov/vgram/blob/babae3775c9c21bcf66f55c434642af3eeb0ce0d/README.md)
- [Official extension SQL](https://github.com/akorotkov/vgram/blob/babae3775c9c21bcf66f55c434642af3eeb0ce0d/vgram--1.0.sql)
- [Official GIN implementation](https://github.com/akorotkov/vgram/blob/babae3775c9c21bcf66f55c434642af3eeb0ce0d/vgram_gin.c)

`vgram` accelerates `LIKE` and `ILIKE` substring searches with GIN indexes built from variable-length q-grams. Unlike fixed trigrams, it learns frequent short grams from the target corpus and indexes longer or rarer grams, which can improve selectivity for a particular data distribution.

### Build a Corpus-Specific Index

Version `1.0` requires PostgreSQL 13 or later. In psql, collect frequent grams, save the result, and pass it to the `vgram_gin_ops` options:

```sql
CREATE EXTENSION vgram;

SELECT qgram_stat(title, 2, 4, 0.05) AS vgrams
FROM documents
\gset

CREATE INDEX documents_title_vgram_idx
ON documents USING gin (
  title vgram_gin_ops (minq=2, maxq=4, vgrams=:'vgrams')
);

SELECT *
FROM documents
WHERE title ILIKE '%indexing%';
```

`minq` and `maxq` bound extracted gram lengths. The threshold passed to `qgram_stat()` marks grams at or above that corpus frequency as too common to index. Use `get_vgrams()` to inspect extraction for a candidate value and learned gram array.

### Objects and Advanced Statistics

- `qgram_stat(text, int, int, float8)` aggregates frequent q-grams.
- `get_vgrams(text, int, int, text[])` shows the grams selected for text.
- `vgram_gin_ops` indexes ordinary `text` for `LIKE` and `ILIKE`.
- `vgram_text` is text-compatible but installs custom frequency statistics and selectivity estimation.
- `vgram_gin_ops2` is the GIN operator class for `vgram_text`.

### Maintenance Boundaries

The learned option array is embedded in the index definition and can be very large. It reflects the sample at build time; major language or distribution changes require new statistics, comparison of recall/plans, and usually `REINDEX` or a replacement index. GIN candidates are rechecked, but overly common or too-short patterns may still fall back to broad scans. Benchmark build time, index size, write amplification, pending-list behavior, VACUUM, collation, multibyte text, case folding, and planner estimates. Keep the training query and parameters reproducible so replicas and rebuilt environments use the same operator-class options.
