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
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `vchord_bm25` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "vchord_bm25_18" "green" >}} {{< bg "17" "vchord_bm25_17" "green" >}} {{< bg "16" "vchord_bm25_16" "green" >}} {{< bg "15" "vchord_bm25_15" "green" >}} {{< bg "14" "vchord_bm25_14" "green" >}} | `vchord_bm25_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "postgresql-18-vchord-bm25" "green" >}} {{< bg "17" "postgresql-17-vchord-bm25" "green" >}} {{< bg "16" "postgresql-16-vchord-bm25" "green" >}} {{< bg "15" "postgresql-15-vchord-bm25" "green" >}} {{< bg "14" "postgresql-14-vchord-bm25" "green" >}} | `postgresql-$v-vchord-bm25` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "vchord_bm25_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-vchord-bm25 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-vchord-bm25 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-vchord-bm25 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-vchord-bm25 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-vchord-bm25 : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-vchord-bm25 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-vchord-bm25 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-vchord-bm25 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-vchord-bm25 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-vchord-bm25 : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_18` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 519.9 KiB | [vchord_bm25_18-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_18-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_18` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 403.2 KiB | [vchord_bm25_18-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_18-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_18` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 536.2 KiB | [vchord_bm25_18-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_18-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_18` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 433.6 KiB | [vchord_bm25_18-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_18-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_18` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 536.7 KiB | [vchord_bm25_18-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_18-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_18` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 433.2 KiB | [vchord_bm25_18-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_18-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 425.2 KiB | [postgresql-18-vchord-bm25_0.3.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 318.2 KiB | [postgresql-18-vchord-bm25_0.3.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 425.4 KiB | [postgresql-18-vchord-bm25_0.3.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 318.0 KiB | [postgresql-18-vchord-bm25_0.3.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 478.3 KiB | [postgresql-18-vchord-bm25_0.3.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 376.2 KiB | [postgresql-18-vchord-bm25_0.3.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 474.6 KiB | [postgresql-18-vchord-bm25_0.3.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 371.5 KiB | [postgresql-18-vchord-bm25_0.3.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_17` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 520.2 KiB | [vchord_bm25_17-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_17-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_17` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 403.2 KiB | [vchord_bm25_17-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_17-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_17` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 536.7 KiB | [vchord_bm25_17-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_17-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_17` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 433.7 KiB | [vchord_bm25_17-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_17-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_17` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 536.8 KiB | [vchord_bm25_17-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_17-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_17` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 433.5 KiB | [vchord_bm25_17-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_17-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 425.1 KiB | [postgresql-17-vchord-bm25_0.3.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 317.9 KiB | [postgresql-17-vchord-bm25_0.3.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 424.9 KiB | [postgresql-17-vchord-bm25_0.3.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 317.7 KiB | [postgresql-17-vchord-bm25_0.3.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 478.8 KiB | [postgresql-17-vchord-bm25_0.3.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 376.2 KiB | [postgresql-17-vchord-bm25_0.3.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 474.9 KiB | [postgresql-17-vchord-bm25_0.3.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 371.5 KiB | [postgresql-17-vchord-bm25_0.3.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_16` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 519.9 KiB | [vchord_bm25_16-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_16-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_16` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 403.3 KiB | [vchord_bm25_16-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_16-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_16` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 536.3 KiB | [vchord_bm25_16-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_16-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_16` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 433.3 KiB | [vchord_bm25_16-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_16-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_16` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 536.4 KiB | [vchord_bm25_16-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_16-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_16` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 433.3 KiB | [vchord_bm25_16-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_16-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 425.0 KiB | [postgresql-16-vchord-bm25_0.3.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 318.0 KiB | [postgresql-16-vchord-bm25_0.3.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 425.1 KiB | [postgresql-16-vchord-bm25_0.3.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 317.8 KiB | [postgresql-16-vchord-bm25_0.3.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 478.5 KiB | [postgresql-16-vchord-bm25_0.3.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 376.1 KiB | [postgresql-16-vchord-bm25_0.3.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 474.5 KiB | [postgresql-16-vchord-bm25_0.3.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 371.5 KiB | [postgresql-16-vchord-bm25_0.3.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_15` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 522.1 KiB | [vchord_bm25_15-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_15-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_15` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 404.9 KiB | [vchord_bm25_15-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_15-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_15` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 538.6 KiB | [vchord_bm25_15-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_15-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_15` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 435.4 KiB | [vchord_bm25_15-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_15-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_15` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 538.2 KiB | [vchord_bm25_15-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_15-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_15` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 434.8 KiB | [vchord_bm25_15-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_15-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 427.0 KiB | [postgresql-15-vchord-bm25_0.3.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 319.7 KiB | [postgresql-15-vchord-bm25_0.3.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 426.8 KiB | [postgresql-15-vchord-bm25_0.3.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 319.6 KiB | [postgresql-15-vchord-bm25_0.3.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 480.3 KiB | [postgresql-15-vchord-bm25_0.3.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 378.2 KiB | [postgresql-15-vchord-bm25_0.3.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 476.5 KiB | [postgresql-15-vchord-bm25_0.3.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 373.4 KiB | [postgresql-15-vchord-bm25_0.3.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_14` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 522.1 KiB | [vchord_bm25_14-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_14-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_14` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 405.2 KiB | [vchord_bm25_14-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_14-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_14` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 538.1 KiB | [vchord_bm25_14-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_14-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_14` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 435.4 KiB | [vchord_bm25_14-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_14-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_14` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 538.2 KiB | [vchord_bm25_14-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_14-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_14` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 435.0 KiB | [vchord_bm25_14-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_14-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 426.8 KiB | [postgresql-14-vchord-bm25_0.3.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 319.5 KiB | [postgresql-14-vchord-bm25_0.3.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 426.9 KiB | [postgresql-14-vchord-bm25_0.3.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 319.6 KiB | [postgresql-14-vchord-bm25_0.3.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 480.4 KiB | [postgresql-14-vchord-bm25_0.3.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 378.3 KiB | [postgresql-14-vchord-bm25_0.3.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 476.3 KiB | [postgresql-14-vchord-bm25_0.3.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 373.7 KiB | [postgresql-14-vchord-bm25_0.3.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-2PIGSTY~noble_arm64.deb) |

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

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'vchord_bm25';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION vchord_bm25;
```



## Usage

> [GitHub: tensorchord/VectorChord-bm25](https://github.com/tensorchord/VectorChord-bm25)

VectorChord-BM25 is a PostgreSQL extension for the BM25 ranking algorithm, implemented via Block-WeakAnd algorithms. It is designed to work together with [pg_tokenizer](https://github.com/tensorchord/pg_tokenizer.rs) for customized text tokenization.

## Architecture

The extension comprises three main components:

1. **Tokenizer**: Converts text into `bm25vector` (sparse vectors storing vocabulary IDs and term frequencies)
2. **bm25vector**: A custom data type for storing tokenized text
3. **bm25vector indexes**: Accelerate search and ranking operations

## Quick Start

```sql
-- Enable required extensions
CREATE EXTENSION IF NOT EXISTS pg_tokenizer CASCADE;
CREATE EXTENSION IF NOT EXISTS vchord_bm25 CASCADE;

-- Create a tokenizer (e.g., LLMLingua2 for English)
SELECT create_tokenizer('tokenizer1', $$
model = "llmlingua2"
$$);

-- Create a table with text content
CREATE TABLE documents (
  id SERIAL PRIMARY KEY,
  passage TEXT,
  embedding bm25vector
);

-- Tokenize text passages into bm25vectors
UPDATE documents SET embedding = tokenize(passage, 'tokenizer1');

-- Create a BM25 index
CREATE INDEX documents_embedding_bm25 ON documents USING bm25 (embedding bm25_ops);

-- Query with BM25 ranking
SELECT id, passage, embedding <&> to_bm25query('documents_embedding_bm25', tokenize('search query', 'tokenizer1')) AS score
FROM documents
ORDER BY score
LIMIT 10;
```

**Note**: BM25 scores in VectorChord-BM25 are negative, with more negative scores indicating greater relevance.

## The `<&>` Operator

The `<&>` operator computes the BM25 relevance score between a stored `bm25vector` and a query `bm25vector`. Queries must be wrapped in `to_bm25query()` which takes the index name and the tokenized query:

```sql
-- Basic search query
-- to_bm25query(index_name, tokenized_query)
SELECT id, passage, embedding <&> to_bm25query('documents_embedding_bm25', tokenize('database system', 'tokenizer1')) AS score
FROM documents
ORDER BY score
LIMIT 10;
```

## Language Support

VectorChord-BM25 supports multiple languages through different tokenizer configurations:

| Language | Approach | Model/Pre-tokenizer |
|----------|----------|---------------------|
| English | Pre-trained model | `model = "llmlingua2"` or `model = "bert_base_uncased"` |
| Chinese | Custom model with Jieba pre-tokenizer | `[pre_tokenizer.jieba]` |
| Japanese | Custom model with Lindera pre-tokenizer | Lindera with IPADIC dictionary |
| Custom | User-trained models via text analyzers | `create_custom_model_tokenizer_and_trigger()` |

### Chinese Text Search Example

Chinese text requires a custom model with a Jieba pre-tokenizer (not a pre-trained model):

```sql
-- Create a text analyzer with Jieba pre-tokenizer
SELECT create_text_analyzer('zh_text_analyzer', $$
[pre_tokenizer.jieba]
$$);

-- Create a custom model tokenizer that trains on your corpus
SELECT create_custom_model_tokenizer_and_trigger(
    tokenizer_name => 'zh_tokenizer',
    model_name => 'zh_model',
    text_analyzer_name => 'zh_text_analyzer',
    table_name => 'documents',
    source_column => 'passage',
    target_column => 'embedding'
);
```

### Custom Tokenizer Models

For domain-specific terminology, you can create text analyzers with stopwords, stemming, and other filters, then train custom models on your corpus using `create_custom_model_tokenizer_and_trigger()`.

## Comparison with Alternatives

| Feature | VectorChord-BM25 | PostgreSQL tsvector + ts_rank |
|---------|-------------------|-------------------------------|
| Ranking algorithm | BM25 | tf-idf variant |
| Custom tokenizers | Yes (via pg_tokenizer) | Limited to built-in configs |
| Index type | Dedicated BM25 index | GIN index |
| Native PostgreSQL | Yes (extension) | Built-in |
| Language support | Extensible via models | Via text search configs |
