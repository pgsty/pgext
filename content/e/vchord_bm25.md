---
title: "vchord_bm25"
linkTitle: "vchord_bm25"
description: "A postgresql extension for bm25 ranking algorithm"
weight: 2150
categories: ["FTS"]
languages: ["Rust"]
licenses: ["AGPL-3.0"]
repos: ["PIGSTY"]
page_width: full
---

[**vchord_bm25**](https://github.com/supervc-stack/VectorChord-bm25) : A postgresql extension for bm25 ranking algorithm


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2150** | {{< badge content="vchord_bm25" link="https://github.com/supervc-stack/VectorChord-bm25" >}} | {{< ext "vchord_bm25" >}} | `0.3.0` | {{< category "FTS" >}} | {{< license "AGPL-3.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `bm25_catalog` |
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pg_textsearch" >}} {{< ext "pg_bestmatch" >}} {{< ext "pg_fts" >}} {{< ext "pgroonga" >}} {{< ext "pg_rrf" >}} {{< ext "psql_bm25s" >}} {{< ext "pgcontext" >}} {{< ext "vectorize" >}} |

> [!Note] bm25 am conflicts with pg_textsearch and pg_search, build require clang upgrade.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `vchord_bm25` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "vchord_bm25_18" "green" >}} {{< bg "17" "vchord_bm25_17" "green" >}} {{< bg "16" "vchord_bm25_16" "green" >}} {{< bg "15" "vchord_bm25_15" "green" >}} {{< bg "14" "vchord_bm25_14" "green" >}} | `vchord_bm25_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "postgresql-18-vchord-bm25" "green" >}} {{< bg "17" "postgresql-17-vchord-bm25" "green" >}} {{< bg "16" "postgresql-16-vchord-bm25" "green" >}} {{< bg "15" "postgresql-15-vchord-bm25" "green" >}} {{< bg "14" "postgresql-14-vchord-bm25" "green" >}} | `postgresql-$v-vchord-bm25` | - |
{.packages}


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
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-vchord-bm25 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-vchord-bm25 : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_18` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [vchord_bm25_18-0.3.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_18-0.3.0-3PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_18` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1014.4 KiB | [vchord_bm25_18-0.3.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_18-0.3.0-3PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_18` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [vchord_bm25_18-0.3.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_18-0.3.0-3PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_18` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [vchord_bm25_18-0.3.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_18-0.3.0-3PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_18` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [vchord_bm25_18-0.3.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_18-0.3.0-3PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_18` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.0 MiB | [vchord_bm25_18-0.3.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_18-0.3.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 881.2 KiB | [postgresql-18-vchord-bm25_0.3.0-4PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-4PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 774.3 KiB | [postgresql-18-vchord-bm25_0.3.0-4PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-4PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 881.1 KiB | [postgresql-18-vchord-bm25_0.3.0-4PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-4PIGSTY~trixie_amd64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 774.9 KiB | [postgresql-18-vchord-bm25_0.3.0-4PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-4PIGSTY~trixie_arm64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 984.5 KiB | [postgresql-18-vchord-bm25_0.3.0-4PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-4PIGSTY~jammy_amd64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 920.7 KiB | [postgresql-18-vchord-bm25_0.3.0-4PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-4PIGSTY~jammy_arm64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 974.9 KiB | [postgresql-18-vchord-bm25_0.3.0-4PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-4PIGSTY~noble_amd64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 907.7 KiB | [postgresql-18-vchord-bm25_0.3.0-4PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-4PIGSTY~noble_arm64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 970.4 KiB | [postgresql-18-vchord-bm25_0.3.0-4PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-4PIGSTY~resolute_amd64.deb) |
| `postgresql-18-vchord-bm25` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 905.5 KiB | [postgresql-18-vchord-bm25_0.3.0-4PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-18-vchord-bm25_0.3.0-4PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_17` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [vchord_bm25_17-0.3.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_17-0.3.0-3PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_17` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1012.2 KiB | [vchord_bm25_17-0.3.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_17-0.3.0-3PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_17` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [vchord_bm25_17-0.3.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_17-0.3.0-3PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_17` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [vchord_bm25_17-0.3.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_17-0.3.0-3PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_17` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [vchord_bm25_17-0.3.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_17-0.3.0-3PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_17` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.0 MiB | [vchord_bm25_17-0.3.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_17-0.3.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 878.9 KiB | [postgresql-17-vchord-bm25_0.3.0-4PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-4PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 772.5 KiB | [postgresql-17-vchord-bm25_0.3.0-4PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-4PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 879.2 KiB | [postgresql-17-vchord-bm25_0.3.0-4PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-4PIGSTY~trixie_amd64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 772.9 KiB | [postgresql-17-vchord-bm25_0.3.0-4PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-4PIGSTY~trixie_arm64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 981.2 KiB | [postgresql-17-vchord-bm25_0.3.0-4PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-4PIGSTY~jammy_amd64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 917.0 KiB | [postgresql-17-vchord-bm25_0.3.0-4PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-4PIGSTY~jammy_arm64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 974.8 KiB | [postgresql-17-vchord-bm25_0.3.0-4PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-4PIGSTY~noble_amd64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 906.0 KiB | [postgresql-17-vchord-bm25_0.3.0-4PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-4PIGSTY~noble_arm64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 967.6 KiB | [postgresql-17-vchord-bm25_0.3.0-4PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-4PIGSTY~resolute_amd64.deb) |
| `postgresql-17-vchord-bm25` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 904.2 KiB | [postgresql-17-vchord-bm25_0.3.0-4PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-17-vchord-bm25_0.3.0-4PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_16` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [vchord_bm25_16-0.3.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_16-0.3.0-3PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_16` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1009.7 KiB | [vchord_bm25_16-0.3.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_16-0.3.0-3PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_16` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [vchord_bm25_16-0.3.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_16-0.3.0-3PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_16` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [vchord_bm25_16-0.3.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_16-0.3.0-3PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_16` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [vchord_bm25_16-0.3.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_16-0.3.0-3PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_16` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.0 MiB | [vchord_bm25_16-0.3.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_16-0.3.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 879.9 KiB | [postgresql-16-vchord-bm25_0.3.0-4PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-4PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 772.4 KiB | [postgresql-16-vchord-bm25_0.3.0-4PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-4PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 880.2 KiB | [postgresql-16-vchord-bm25_0.3.0-4PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-4PIGSTY~trixie_amd64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 772.1 KiB | [postgresql-16-vchord-bm25_0.3.0-4PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-4PIGSTY~trixie_arm64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 984.3 KiB | [postgresql-16-vchord-bm25_0.3.0-4PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-4PIGSTY~jammy_amd64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 916.1 KiB | [postgresql-16-vchord-bm25_0.3.0-4PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-4PIGSTY~jammy_arm64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 972.8 KiB | [postgresql-16-vchord-bm25_0.3.0-4PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-4PIGSTY~noble_amd64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 905.6 KiB | [postgresql-16-vchord-bm25_0.3.0-4PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-4PIGSTY~noble_arm64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 968.8 KiB | [postgresql-16-vchord-bm25_0.3.0-4PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-4PIGSTY~resolute_amd64.deb) |
| `postgresql-16-vchord-bm25` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 903.5 KiB | [postgresql-16-vchord-bm25_0.3.0-4PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-16-vchord-bm25_0.3.0-4PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_15` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [vchord_bm25_15-0.3.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_15-0.3.0-3PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_15` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1003.8 KiB | [vchord_bm25_15-0.3.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_15-0.3.0-3PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_15` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [vchord_bm25_15-0.3.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_15-0.3.0-3PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_15` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [vchord_bm25_15-0.3.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_15-0.3.0-3PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_15` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [vchord_bm25_15-0.3.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_15-0.3.0-3PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_15` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.0 MiB | [vchord_bm25_15-0.3.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_15-0.3.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 877.1 KiB | [postgresql-15-vchord-bm25_0.3.0-4PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-4PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 770.6 KiB | [postgresql-15-vchord-bm25_0.3.0-4PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-4PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 878.2 KiB | [postgresql-15-vchord-bm25_0.3.0-4PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-4PIGSTY~trixie_amd64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 770.8 KiB | [postgresql-15-vchord-bm25_0.3.0-4PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-4PIGSTY~trixie_arm64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 976.9 KiB | [postgresql-15-vchord-bm25_0.3.0-4PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-4PIGSTY~jammy_amd64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 914.5 KiB | [postgresql-15-vchord-bm25_0.3.0-4PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-4PIGSTY~jammy_arm64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 970.6 KiB | [postgresql-15-vchord-bm25_0.3.0-4PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-4PIGSTY~noble_amd64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 902.4 KiB | [postgresql-15-vchord-bm25_0.3.0-4PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-4PIGSTY~noble_arm64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 965.3 KiB | [postgresql-15-vchord-bm25_0.3.0-4PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-4PIGSTY~resolute_amd64.deb) |
| `postgresql-15-vchord-bm25` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 900.4 KiB | [postgresql-15-vchord-bm25_0.3.0-4PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-15-vchord-bm25_0.3.0-4PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_bm25_14` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [vchord_bm25_14-0.3.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_bm25_14-0.3.0-3PIGSTY.el8.x86_64.rpm) |
| `vchord_bm25_14` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1001.6 KiB | [vchord_bm25_14-0.3.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_bm25_14-0.3.0-3PIGSTY.el8.aarch64.rpm) |
| `vchord_bm25_14` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [vchord_bm25_14-0.3.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_bm25_14-0.3.0-3PIGSTY.el9.x86_64.rpm) |
| `vchord_bm25_14` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [vchord_bm25_14-0.3.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_bm25_14-0.3.0-3PIGSTY.el9.aarch64.rpm) |
| `vchord_bm25_14` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [vchord_bm25_14-0.3.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_bm25_14-0.3.0-3PIGSTY.el10.x86_64.rpm) |
| `vchord_bm25_14` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.0 MiB | [vchord_bm25_14-0.3.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_bm25_14-0.3.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 873.7 KiB | [postgresql-14-vchord-bm25_0.3.0-4PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-4PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 768.3 KiB | [postgresql-14-vchord-bm25_0.3.0-4PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-4PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 873.8 KiB | [postgresql-14-vchord-bm25_0.3.0-4PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-4PIGSTY~trixie_amd64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 769.1 KiB | [postgresql-14-vchord-bm25_0.3.0-4PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-4PIGSTY~trixie_arm64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 976.2 KiB | [postgresql-14-vchord-bm25_0.3.0-4PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-4PIGSTY~jammy_amd64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 912.7 KiB | [postgresql-14-vchord-bm25_0.3.0-4PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-4PIGSTY~jammy_arm64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 966.4 KiB | [postgresql-14-vchord-bm25_0.3.0-4PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-4PIGSTY~noble_amd64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 901.3 KiB | [postgresql-14-vchord-bm25_0.3.0-4PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-4PIGSTY~noble_arm64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 962.5 KiB | [postgresql-14-vchord-bm25_0.3.0-4PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-4PIGSTY~resolute_amd64.deb) |
| `postgresql-14-vchord-bm25` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 898.0 KiB | [postgresql-14-vchord-bm25_0.3.0-4PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord-bm25/postgresql-14-vchord-bm25_0.3.0-4PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supervc-stack/VectorChord-bm25" title="Repository" icon="github" subtitle="github.com/supervc-stack/VectorChord-bm25" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="VectorChord-bm25-0.3.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg vchord_bm25;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

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

Sources:

- [0.3.0 README](https://github.com/supervc-stack/VectorChord-bm25/blob/0.3.0/README.md)
- [Control file](https://github.com/supervc-stack/VectorChord-bm25/blob/0.3.0/vchord_bm25.control)
- [0.3.0 SQL objects](https://github.com/supervc-stack/VectorChord-bm25/blob/0.3.0/sql/install/vchord_bm25--0.3.0.sql)
- [Query settings](https://github.com/supervc-stack/VectorChord-bm25/blob/0.3.0/src/guc.rs)
- [0.3.0 migration](https://github.com/supervc-stack/VectorChord-bm25/blob/0.3.0/sql/vchord_bm25--0.2.2--0.3.0.sql)
- [0.3.0 release](https://github.com/supervc-stack/VectorChord-bm25/releases/tag/0.3.0)
- [Tokenizer installation](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/01-installation.md)
- [Tokenizer models](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/06-model.md)

`vchord_bm25` provides BM25 ranking with a sparse token-frequency type and the `bm25` index access method. Tokenization is supplied separately, commonly by pg_tokenizer. Extension objects live in the fixed `bm25_catalog` schema, and creation requires superuser privileges.

### Core Workflow

The example uses pg_tokenizer, which requires preloading and a restart. Preserve existing entries in the preload list:

```conf
shared_preload_libraries = 'pg_tokenizer'
```

```sql
CREATE EXTENSION pg_tokenizer;
CREATE EXTENSION vchord_bm25;
SET search_path = public, tokenizer_catalog, bm25_catalog;

SELECT create_tokenizer('english', $$
model = "bert_base_uncased"
$$);
CREATE TABLE documents (
    id bigserial PRIMARY KEY,
    passage text,
    embedding bm25vector
);
INSERT INTO documents(passage) VALUES ('PostgreSQL full text search');
UPDATE documents SET embedding = tokenize(passage, 'english')::bm25vector;
CREATE INDEX documents_bm25 ON documents USING bm25 (embedding bm25_ops);

SELECT id, passage,
       embedding <&> to_bm25query('documents_bm25',
           tokenize('PostgreSQL', 'english')::bm25vector) AS score
FROM documents
ORDER BY score
LIMIT 10;
```

The index supplies corpus statistics to `to_bm25query`; the score from `<&>` is negative, so ascending order returns greater relevance first. Use the same tokenizer/model for documents and queries. Update stored token vectors when source text changes, or use the tokenizer's maintenance-trigger helper. Changing the vocabulary requires retokenizing stored documents before rebuilding their index.

### Types, Functions, and Search Limits

- `bm25vector` stores token IDs and frequencies; the integer-array cast aggregates duplicate IDs and discards token order.
- `bm25query` binds the query vector to an index. `to_bm25query(regclass, bm25vector)` constructs it; `bm25_ops` is the index operator class.
- `bm25_catalog.bm25_limit` defaults to 100 and limits candidates returned by the index. Increase it for larger SQL limits or restrictive filters; changing SQL LIMIT alone does not increase this candidate budget.
- `bm25_catalog.enable_index` controls use of the index; `bm25_catalog.enable_prefilter` controls prefiltering. Both default to true.
- `bm25_catalog.segment_growing_max_page_size` defaults to 4096 pages before sealing a growing segment.

```sql
SET bm25_catalog.bm25_limit = 1000;
```

The access-method name is global: this extension cannot coexist in a database with another extension that creates the same bm25 access method, including pg_textsearch and the compatibility alias in pg_search. Sparse frequencies do not preserve positions for phrase matching. Chinese text can use a custom corpus model with a Jieba pre-tokenizer; Japanese Lindera support depends on the tokenizer build and dictionary configuration.

### Upgrade to 0.3.0

```sql
ALTER EXTENSION vchord_bm25 UPDATE TO '0.3.0';
```

The 0.2.2-to-0.3.0 migration adds `bm25_page_inspect(regclass, integer)`, returning diagnostic page text. The release changes sealed-segment page allocation for small tokens and does not document a mandatory index rebuild. Install matching extension files before updating database objects; replacing the preloaded tokenizer library also requires a restart. Keep tokenization and ranking upgrades compatible and check representative query results.
