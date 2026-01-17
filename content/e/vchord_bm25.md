---
title: "vchord_bm25"
linkTitle: "vchord_bm25"
description: "A postgresql extension for bm25 ranking algorithm"
weight: 2150
categories: ["FTS"]
width: full
---

[**vchord_bm25**](https://github.com/tensorchord/VectorChord-bm25) : A postgresql extension for bm25 ranking algorithm


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2150** | {{< badge content="vchord_bm25" link="https://github.com/tensorchord/VectorChord-bm25" >}} | {{< ext "vchord_bm25" >}} | `0.3.0` | {{< category "FTS" >}} | {{< license "AGPL-3.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `bm25_catalog` |
|   **See Also**    | {{< ext "vector" >}} {{< ext "vchord" >}} {{< ext "pg_search" >}} {{< ext "pg_bestmatch" >}} {{< ext "vectorscale" >}} {{< ext "zhparser" >}} {{< ext "pg_tokenizer" >}} {{< ext "pgroonga" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `vchord_bm25` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "vchord_bm25_18" "green" >}} {{< bg "17" "vchord_bm25_17" "green" >}} {{< bg "16" "vchord_bm25_16" "green" >}} {{< bg "15" "vchord_bm25_15" "green" >}} {{< bg "14" "vchord_bm25_14" "green" >}} {{< bg "13" "vchord_bm25_13" "green" >}} | `vchord_bm25_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "postgresql-18-vchord-bm25" "green" >}} {{< bg "17" "postgresql-17-vchord-bm25" "green" >}} {{< bg "16" "postgresql-16-vchord-bm25" "green" >}} {{< bg "15" "postgresql-15-vchord-bm25" "green" >}} {{< bg "14" "postgresql-14-vchord-bm25" "green" >}} {{< bg "13" "postgresql-13-vchord-bm25" "green" >}} | `postgresql-$v-vchord-bm25` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-vchord-bm25 : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_18` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 523.6 KiB | [vchord_bm25_18-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_18-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_18` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 414.7 KiB | [vchord_bm25_18-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_18-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_18` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 539.9 KiB | [vchord_bm25_18-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_18-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_18` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 446.2 KiB | [vchord_bm25_18-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_18-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_18` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 540.0 KiB | [vchord_bm25_18-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_18-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_18` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 446.2 KiB | [vchord_bm25_18-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_18-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 428.6 KiB | [postgresql-18-vchord-bm25_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 327.8 KiB | [postgresql-18-vchord-bm25_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 428.5 KiB | [postgresql-18-vchord-bm25_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 327.8 KiB | [postgresql-18-vchord-bm25_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 480.9 KiB | [postgresql-18-vchord-bm25_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 387.6 KiB | [postgresql-18-vchord-bm25_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 477.8 KiB | [postgresql-18-vchord-bm25_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 383.3 KiB | [postgresql-18-vchord-bm25_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_17` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 523.9 KiB | [vchord_bm25_17-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_17-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_17` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 414.4 KiB | [vchord_bm25_17-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_17-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_17` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 539.9 KiB | [vchord_bm25_17-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_17-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_17` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 445.9 KiB | [vchord_bm25_17-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_17-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_17` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 540.0 KiB | [vchord_bm25_17-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_17-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_17` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 445.8 KiB | [vchord_bm25_17-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_17-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 428.2 KiB | [postgresql-17-vchord-bm25_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 327.7 KiB | [postgresql-17-vchord-bm25_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 428.5 KiB | [postgresql-17-vchord-bm25_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 327.7 KiB | [postgresql-17-vchord-bm25_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 480.8 KiB | [postgresql-17-vchord-bm25_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 387.6 KiB | [postgresql-17-vchord-bm25_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 477.2 KiB | [postgresql-17-vchord-bm25_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 383.4 KiB | [postgresql-17-vchord-bm25_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_16` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 523.7 KiB | [vchord_bm25_16-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_16-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_16` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 414.6 KiB | [vchord_bm25_16-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_16-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_16` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 539.7 KiB | [vchord_bm25_16-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_16-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_16` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 446.0 KiB | [vchord_bm25_16-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_16-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_16` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 539.8 KiB | [vchord_bm25_16-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_16-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_16` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 445.9 KiB | [vchord_bm25_16-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_16-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 428.5 KiB | [postgresql-16-vchord-bm25_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 327.7 KiB | [postgresql-16-vchord-bm25_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 428.1 KiB | [postgresql-16-vchord-bm25_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 327.7 KiB | [postgresql-16-vchord-bm25_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 480.5 KiB | [postgresql-16-vchord-bm25_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 387.3 KiB | [postgresql-16-vchord-bm25_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 477.1 KiB | [postgresql-16-vchord-bm25_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 382.8 KiB | [postgresql-16-vchord-bm25_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_15` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 525.9 KiB | [vchord_bm25_15-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_15-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_15` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 416.3 KiB | [vchord_bm25_15-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_15-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_15` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 541.3 KiB | [vchord_bm25_15-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_15-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_15` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 448.1 KiB | [vchord_bm25_15-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_15-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_15` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 541.5 KiB | [vchord_bm25_15-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_15-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_15` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 448.0 KiB | [vchord_bm25_15-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_15-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 430.4 KiB | [postgresql-15-vchord-bm25_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 329.3 KiB | [postgresql-15-vchord-bm25_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 430.4 KiB | [postgresql-15-vchord-bm25_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 329.3 KiB | [postgresql-15-vchord-bm25_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 482.4 KiB | [postgresql-15-vchord-bm25_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 388.9 KiB | [postgresql-15-vchord-bm25_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 479.3 KiB | [postgresql-15-vchord-bm25_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 384.6 KiB | [postgresql-15-vchord-bm25_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_14` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 526.0 KiB | [vchord_bm25_14-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_14-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_14` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 416.3 KiB | [vchord_bm25_14-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_14-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_14` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 542.0 KiB | [vchord_bm25_14-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_14-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_14` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 448.1 KiB | [vchord_bm25_14-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_14-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_14` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 542.1 KiB | [vchord_bm25_14-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_14-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_14` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 447.9 KiB | [vchord_bm25_14-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_14-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 430.4 KiB | [postgresql-14-vchord-bm25_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 329.7 KiB | [postgresql-14-vchord-bm25_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 430.5 KiB | [postgresql-14-vchord-bm25_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 329.5 KiB | [postgresql-14-vchord-bm25_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 482.2 KiB | [postgresql-14-vchord-bm25_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 389.1 KiB | [postgresql-14-vchord-bm25_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 479.3 KiB | [postgresql-14-vchord-bm25_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 384.5 KiB | [postgresql-14-vchord-bm25_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_13` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 526.1 KiB | [vchord_bm25_13-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_13-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_13` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 416.1 KiB | [vchord_bm25_13-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_13-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_13` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 542.2 KiB | [vchord_bm25_13-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_13-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_13` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 448.2 KiB | [vchord_bm25_13-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_13-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_13` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 542.3 KiB | [vchord_bm25_13-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_13-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_13` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 447.9 KiB | [vchord_bm25_13-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_13-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-vchord-bm25` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 430.6 KiB | [postgresql-13-vchord-bm25_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-13-vchord-bm25_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-vchord-bm25` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 329.4 KiB | [postgresql-13-vchord-bm25_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-13-vchord-bm25_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-vchord-bm25` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 430.4 KiB | [postgresql-13-vchord-bm25_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-13-vchord-bm25_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-vchord-bm25` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 329.1 KiB | [postgresql-13-vchord-bm25_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-13-vchord-bm25_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-vchord-bm25` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 482.7 KiB | [postgresql-13-vchord-bm25_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-13-vchord-bm25_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-vchord-bm25` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 389.1 KiB | [postgresql-13-vchord-bm25_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-13-vchord-bm25_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-vchord-bm25` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 479.5 KiB | [postgresql-13-vchord-bm25_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-13-vchord-bm25_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-vchord-bm25` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 384.2 KiB | [postgresql-13-vchord-bm25_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-13-vchord-bm25_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tensorchord/VectorChord-bm25" title="Repository" icon="github" subtitle="github.com/tensorchord/VectorChord-bm25" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="VectorChord-bm25-0.3.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg vchord_bm25;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install vchord_bm25;		# install via package name, for the active PG version

pig install vchord_bm25 -v 18;   # install for PG 18
pig install vchord_bm25 -v 17;   # install for PG 17
pig install vchord_bm25 -v 16;   # install for PG 16
pig install vchord_bm25 -v 15;   # install for PG 15
pig install vchord_bm25 -v 14;   # install for PG 14
pig install vchord_bm25 -v 13;   # install for PG 13

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'vchord_bm25';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION vchord_bm25;
```
