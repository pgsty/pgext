---
title: "vchord_bm25"
linkTitle: "vchord_bm25"
description: "A postgresql extension for bm25 ranking algorithm"
weight: 2160
categories: ["FTS"]
languages: ["Rust"]
licenses: ["AGPL-3.0-only"]
repos: ["PIGSTY"]
page_width: full
---

[**vchord_bm25**](https://vchord.pgsty.com/docs/bm25/) : A postgresql extension for bm25 ranking algorithm


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2160** | {{< badge content="vchord_bm25" link="https://vchord.pgsty.com/docs/bm25/" >}} | {{< ext "vchord_bm25" >}} | `1.0.0` | {{< category "FTS" >}} | {{< license "AGPL-3.0-only" >}} | {{< language "Rust" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pg_textsearch" >}} {{< ext "pg_bestmatch" >}} {{< ext "pg_fts" >}} {{< ext "pgroonga" >}} {{< ext "pg_rrf" >}} {{< ext "psql_bm25s" >}} {{< ext "pgcontext" >}} {{< ext "vectorize" >}} |

> [!Note] PGSTY fork 1.0.0; native tsvector and new index format; preload required. In-place upgrade from 0.3.x is rejected.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `vchord_bm25` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "vchord_bm25_18" "green" >}} {{< bg "17" "vchord_bm25_17" "green" >}} {{< bg "16" "vchord_bm25_16" "green" >}} {{< bg "15" "vchord_bm25_15" "green" >}} {{< bg "14" "vchord_bm25_14" "green" >}} | `vchord_bm25_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-vchord-bm25" "green" >}} {{< bg "17" "postgresql-17-vchord-bm25" "green" >}} {{< bg "16" "postgresql-16-vchord-bm25" "green" >}} {{< bg "15" "postgresql-15-vchord-bm25" "green" >}} {{< bg "14" "postgresql-14-vchord-bm25" "green" >}} | `postgresql-$v-vchord-bm25` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_bm25_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.8 MiB | [vchord_bm25_18-1.0.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_18-1.0.0-1PGSTY.el8.x86_64.rpm) |
| `vchord_bm25_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [vchord_bm25_18-1.0.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_18-1.0.0-1PGSTY.el8.aarch64.rpm) |
| `vchord_bm25_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.8 MiB | [vchord_bm25_18-1.0.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_18-1.0.0-1PGSTY.el9.x86_64.rpm) |
| `vchord_bm25_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [vchord_bm25_18-1.0.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_18-1.0.0-1PGSTY.el9.aarch64.rpm) |
| `vchord_bm25_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.8 MiB | [vchord_bm25_18-1.0.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_18-1.0.0-1PGSTY.el10.x86_64.rpm) |
| `vchord_bm25_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [vchord_bm25_18-1.0.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_18-1.0.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-vchord-bm25` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.5 MiB | [postgresql-18-vchord-bm25_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-vchord-bm25` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-18-vchord-bm25_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-vchord-bm25` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.5 MiB | [postgresql-18-vchord-bm25_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-vchord-bm25` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-18-vchord-bm25_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-vchord-bm25` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.7 MiB | [postgresql-18-vchord-bm25_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-vchord-bm25` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.5 MiB | [postgresql-18-vchord-bm25_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-vchord-bm25` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.7 MiB | [postgresql-18-vchord-bm25_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-vchord-bm25` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.5 MiB | [postgresql-18-vchord-bm25_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-vchord-bm25` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.7 MiB | [postgresql-18-vchord-bm25_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-vchord-bm25` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.5 MiB | [postgresql-18-vchord-bm25_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.8 MiB | [vchord_bm25_17-1.0.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_17-1.0.0-1PGSTY.el8.x86_64.rpm) |
| `vchord_bm25_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [vchord_bm25_17-1.0.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_17-1.0.0-1PGSTY.el8.aarch64.rpm) |
| `vchord_bm25_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.8 MiB | [vchord_bm25_17-1.0.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_17-1.0.0-1PGSTY.el9.x86_64.rpm) |
| `vchord_bm25_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [vchord_bm25_17-1.0.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_17-1.0.0-1PGSTY.el9.aarch64.rpm) |
| `vchord_bm25_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.8 MiB | [vchord_bm25_17-1.0.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_17-1.0.0-1PGSTY.el10.x86_64.rpm) |
| `vchord_bm25_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [vchord_bm25_17-1.0.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_17-1.0.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-vchord-bm25` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.5 MiB | [postgresql-17-vchord-bm25_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-vchord-bm25` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-17-vchord-bm25_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-vchord-bm25` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.5 MiB | [postgresql-17-vchord-bm25_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-vchord-bm25` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-17-vchord-bm25_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-vchord-bm25` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.7 MiB | [postgresql-17-vchord-bm25_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-vchord-bm25` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.5 MiB | [postgresql-17-vchord-bm25_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-vchord-bm25` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.7 MiB | [postgresql-17-vchord-bm25_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-vchord-bm25` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.5 MiB | [postgresql-17-vchord-bm25_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-vchord-bm25` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.7 MiB | [postgresql-17-vchord-bm25_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-vchord-bm25` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.5 MiB | [postgresql-17-vchord-bm25_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.8 MiB | [vchord_bm25_16-1.0.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_16-1.0.0-1PGSTY.el8.x86_64.rpm) |
| `vchord_bm25_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [vchord_bm25_16-1.0.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_16-1.0.0-1PGSTY.el8.aarch64.rpm) |
| `vchord_bm25_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.8 MiB | [vchord_bm25_16-1.0.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_16-1.0.0-1PGSTY.el9.x86_64.rpm) |
| `vchord_bm25_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [vchord_bm25_16-1.0.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_16-1.0.0-1PGSTY.el9.aarch64.rpm) |
| `vchord_bm25_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.8 MiB | [vchord_bm25_16-1.0.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_16-1.0.0-1PGSTY.el10.x86_64.rpm) |
| `vchord_bm25_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [vchord_bm25_16-1.0.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_16-1.0.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-vchord-bm25` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.5 MiB | [postgresql-16-vchord-bm25_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-vchord-bm25` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-16-vchord-bm25_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-vchord-bm25` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.5 MiB | [postgresql-16-vchord-bm25_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-vchord-bm25` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-16-vchord-bm25_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-vchord-bm25` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.7 MiB | [postgresql-16-vchord-bm25_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-vchord-bm25` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.5 MiB | [postgresql-16-vchord-bm25_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-vchord-bm25` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.7 MiB | [postgresql-16-vchord-bm25_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-vchord-bm25` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.5 MiB | [postgresql-16-vchord-bm25_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-vchord-bm25` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.7 MiB | [postgresql-16-vchord-bm25_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-vchord-bm25` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.5 MiB | [postgresql-16-vchord-bm25_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.8 MiB | [vchord_bm25_15-1.0.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_15-1.0.0-1PGSTY.el8.x86_64.rpm) |
| `vchord_bm25_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [vchord_bm25_15-1.0.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_15-1.0.0-1PGSTY.el8.aarch64.rpm) |
| `vchord_bm25_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.8 MiB | [vchord_bm25_15-1.0.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_15-1.0.0-1PGSTY.el9.x86_64.rpm) |
| `vchord_bm25_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [vchord_bm25_15-1.0.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_15-1.0.0-1PGSTY.el9.aarch64.rpm) |
| `vchord_bm25_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.8 MiB | [vchord_bm25_15-1.0.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_15-1.0.0-1PGSTY.el10.x86_64.rpm) |
| `vchord_bm25_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [vchord_bm25_15-1.0.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_15-1.0.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-vchord-bm25` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.5 MiB | [postgresql-15-vchord-bm25_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-vchord-bm25` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-15-vchord-bm25_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-vchord-bm25` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.5 MiB | [postgresql-15-vchord-bm25_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-vchord-bm25` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-15-vchord-bm25_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-vchord-bm25` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.7 MiB | [postgresql-15-vchord-bm25_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-vchord-bm25` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.5 MiB | [postgresql-15-vchord-bm25_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-vchord-bm25` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.7 MiB | [postgresql-15-vchord-bm25_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-vchord-bm25` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.5 MiB | [postgresql-15-vchord-bm25_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-vchord-bm25` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.7 MiB | [postgresql-15-vchord-bm25_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-vchord-bm25` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.5 MiB | [postgresql-15-vchord-bm25_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.8 MiB | [vchord_bm25_14-1.0.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_14-1.0.0-1PGSTY.el8.x86_64.rpm) |
| `vchord_bm25_14` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [vchord_bm25_14-1.0.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_14-1.0.0-1PGSTY.el8.aarch64.rpm) |
| `vchord_bm25_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.8 MiB | [vchord_bm25_14-1.0.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_14-1.0.0-1PGSTY.el9.x86_64.rpm) |
| `vchord_bm25_14` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [vchord_bm25_14-1.0.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_14-1.0.0-1PGSTY.el9.aarch64.rpm) |
| `vchord_bm25_14` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.8 MiB | [vchord_bm25_14-1.0.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_14-1.0.0-1PGSTY.el10.x86_64.rpm) |
| `vchord_bm25_14` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [vchord_bm25_14-1.0.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_14-1.0.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-vchord-bm25` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.5 MiB | [postgresql-14-vchord-bm25_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-vchord-bm25` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.3 MiB | [postgresql-14-vchord-bm25_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-vchord-bm25` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.5 MiB | [postgresql-14-vchord-bm25_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-vchord-bm25` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.3 MiB | [postgresql-14-vchord-bm25_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-vchord-bm25` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.7 MiB | [postgresql-14-vchord-bm25_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-vchord-bm25` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.5 MiB | [postgresql-14-vchord-bm25_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-vchord-bm25` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.7 MiB | [postgresql-14-vchord-bm25_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-14-vchord-bm25` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.5 MiB | [postgresql-14-vchord-bm25_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-14-vchord-bm25` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.7 MiB | [postgresql-14-vchord-bm25_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-vchord-bm25` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.5 MiB | [postgresql-14-vchord-bm25_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://vchord.pgsty.com/docs/bm25/" title="Repository" icon="link" subtitle="vchord.pgsty.com/docs/bm25/" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="vchord_bm25-1.0.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg vchord_bm25;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install vchord_bm25;		# install via package name, for the active PG version

pig install vchord_bm25 -v 18;   # install for PG 18
pig install vchord_bm25 -v 17;   # install for PG 17
pig install vchord_bm25 -v 16;   # install for PG 16
pig install vchord_bm25 -v 15;   # install for PG 15
pig install vchord_bm25 -v 14;   # install for PG 14

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'vchord_bm25';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION vchord_bm25;
```

## Usage

Sources:

- [Installation](https://vchord.pgsty.com/docs/bm25/getting-started/installation/)
- [Native tsvector quick start](https://vchord.pgsty.com/docs/bm25/getting-started/quick-start/)
- [1.0.0 migration and permissions](https://vchord.pgsty.com/docs/bm25/reference/migration/)
- [SQL reference](https://vchord.pgsty.com/docs/bm25/reference/sql/)
- [GUCs and index options](https://vchord.pgsty.com/docs/bm25/reference/gucs/)
- [Index maintenance](https://vchord.pgsty.com/docs/bm25/usage/indexing/)
- [1.0.0 control file](https://github.com/pgsty/vchord_bm25/blob/1.0.0/vchord_bm25.control)

`vchord_bm25` 1.0.0 adds BM25 relevance ranking over native PostgreSQL `tsvector` values on PostgreSQL 14–18. Its `bm25` access method and `bm25_ops` operator class use a new index format. Standard text search needs no tokenizer extension; `pg_tokenizer` 1.0.0 is optional for custom analyzers and multilingual models. Installation requires superuser privileges and shared preloading.

### Create and Query an Index

Add the library to the existing preload list and restart PostgreSQL:

```ini
shared_preload_libraries = 'vchord_bm25'
```

```sql
CREATE EXTENSION vchord_bm25;
CREATE TABLE documents (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    passage text NOT NULL,
    terms tsvector GENERATED ALWAYS AS (to_tsvector('english', passage)) STORED
);
INSERT INTO documents (passage) VALUES
    ('PostgreSQL provides full text search.'),
    ('BM25 ranks PostgreSQL documents by relevance.'),
    ('A guide to database indexing.');
CREATE INDEX documents_bm25 ON documents USING bm25 (terms bm25_ops)
WITH (options = $$
k1 = 1.2
b = 0.75
$$, limit = 100);

SELECT id, passage,
       terms <&> to_bm25query(to_tsvector('english', 'PostgreSQL'), 'documents_bm25') AS rank
FROM documents
WHERE terms @@ plainto_tsquery('english', 'PostgreSQL')
ORDER BY rank LIMIT 10;
```

`to_bm25query(tsvector, regclass)` takes the query vector first and the target index second. `<&>` returns a negative BM25 score as `double precision`: ascending order puts greater relevance first. Negate the result only for display if a positive score is preferred. Ordering alone also includes zero-score nonmatches and NULLs; the example uses an explicit match predicate.

Keep positions in stored vectors: `strip(tsvector)` and `array_to_tsvector` cannot provide the term frequencies required by the index. Document weights are ignored, and native position limits can reduce frequency detail for long documents. Use the same analyzer for documents and queries. With custom tokenization, `analyze_to_tsvector` stores string lexemes and `tokenize_to_tsvector` stores model IDs as lexemes; both are `STABLE`, so maintain an ordinary column through the application or a trigger.

### Search Controls and Maintenance

| Setting or function | Meaning in 1.0.0 |
| --- | --- |
| `bm25.enable_scan` | Allow BM25 index plans; default true. |
| `bm25.limit`, index `limit` | Initial scan batch size, default 100; zero also selects 100. Scans continue for SQL limits, offsets, and filtering. |
| `bm25.prefilter`, index `prefilter` | Apply supported predicates before ranking candidates; default false. |
| `k1`, `b` | TOML scoring parameters; defaults 1.2 and 0.75, ranges 1.2–2.0 and 0–1. |
| `maintenance_min_rows`, `maintenance_scale_factor` | Automatic merge thresholds; defaults 1000 and 0.1. |
| `bm25_maintain(regclass)` | Request a merge without the automatic thresholds; returns whether a segment was published. |

Explicit GUC values, including false or the nominal default, override index options. Server, database, role, and startup settings also take precedence. Index defaults apply only when the GUC source is the built-in default.

```sql
SET "bm25.limit" = 200;
SET bm25.prefilter = true;
VACUUM documents;
SELECT bm25_maintain('documents_bm25');
```

`VACUUM` marks dead records and can merge segments; `ANALYZE` alone does not merge them. Manual maintenance requires table ownership or, on PostgreSQL 17+, `MAINTAIN` privilege, and does not replace dead-tuple detection. Deleted records stop appearing in results immediately but remain in corpus statistics until compaction. Long snapshots can delay page reuse and further merges. Standby readers encountering recycled pages can receive SQLSTATE `40001` and need to retry.

### Permissions and Migration

The extension is relocatable; applications need access to the selected installation schema and indexed data. Scoring requires `SELECT` on the indexed table, or column-level access for an ordinary non-partial single-column index. Expression and partial indexes require table-level access. Both scalar scoring and native scans reject queries when row-level security applies to the current role, because scores use global corpus statistics.

The index supports one indexed column and does not support unlogged tables. The access-method name is database-wide, so another extension defining `bm25` cannot coexist in that database.

Upgrades from 0.3.x are deliberately blocked. Keep the old database and library available, export source text, and create a fresh 1.0.0 installation with its own library directory. Recompute native vectors, rebuild indexes, and adapt query argument order, types, settings, and triggers before switching the application. Another database or port sharing the old library directory is insufficient isolation. Neither a binary index copy nor `REINDEX` on an old `bm25vector` column migrates the API.
