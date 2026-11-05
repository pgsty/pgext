---
title: "vchord"
linkTitle: "vchord"
description: "Vector database plugin for Postgres, written in Rust"
weight: 1810
categories: ["RAG"]
languages: ["Rust"]
licenses: ["AGPL-3.0-only"]
repos: ["PIGSTY"]
page_width: full
---

[**vchord**](https://vchord.pgsty.com/docs/vchord/) : Vector database plugin for Postgres, written in Rust


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1810** | {{< badge content="vchord" link="https://vchord.pgsty.com/docs/vchord/" >}} | {{< ext "vchord" >}} | `1.2.0` | {{< category "RAG" >}} | {{< license "AGPL-3.0-only" >}} | {{< language "Rust" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "vector" >}} |
|   **See Also**    | {{< ext "vector" >}} {{< ext "vectorscale" >}} {{< ext "pgcontext" >}} {{< ext "vectorize" >}} {{< ext "pg_rrf" >}} {{< ext "pg_search" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_bestmatch" >}} {{< ext "pgml" >}} {{< ext "pg4ml" >}} |

> [!Note] PGSTY fork; requires vector and preload. Reissued 1.2.0 tag archive is byte-identical; 1.2.0-2PGSTY rebuilt and locally indexed across six RPM and ten DEB platforms, PG14-18. vchordg/halfvec scans can return distances out of order; explicitly rerank candidates when strict order is required. CDN sync pending.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `vchord` | `vector` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.0` | {{< bg "18" "vchord_18" "green" >}} {{< bg "17" "vchord_17" "green" >}} {{< bg "16" "vchord_16" "green" >}} {{< bg "15" "vchord_15" "green" >}} {{< bg "14" "vchord_14" "green" >}} | `vchord_$v` | `pgvector_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.0` | {{< bg "18" "postgresql-18-vchord" "green" >}} {{< bg "17" "postgresql-17-vchord" "green" >}} {{< bg "16" "postgresql-16-vchord" "green" >}} {{< bg "15" "postgresql-15-vchord" "green" >}} {{< bg "14" "postgresql-14-vchord" "green" >}} | `postgresql-$v-vchord` | `postgresql-$v-pgvector` |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "vchord_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "vchord_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "vchord_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "vchord_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "vchord_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "vchord_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "vchord_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-vchord : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-vchord : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-vchord : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-vchord : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-vchord : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-vchord : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-vchord : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-vchord : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-vchord : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-vchord : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_18` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.5 MiB | [vchord_18-1.2.0-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_18-1.2.0-2PGSTY.el8.x86_64.rpm) |
| `vchord_18` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.8 MiB | [vchord_18-1.2.0-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_18-1.2.0-2PGSTY.el8.aarch64.rpm) |
| `vchord_18` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.5 MiB | [vchord_18-1.2.0-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_18-1.2.0-2PGSTY.el9.x86_64.rpm) |
| `vchord_18` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.0 MiB | [vchord_18-1.2.0-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_18-1.2.0-2PGSTY.el9.aarch64.rpm) |
| `vchord_18` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.5 MiB | [vchord_18-1.2.0-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_18-1.2.0-2PGSTY.el10.x86_64.rpm) |
| `vchord_18` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.0 MiB | [vchord_18-1.2.0-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_18-1.2.0-2PGSTY.el10.aarch64.rpm) |
| `postgresql-18-vchord` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.0 MiB | [postgresql-18-vchord_1.2.0-2PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-18-vchord_1.2.0-2PGSTY~bookworm_amd64.deb) |
| `postgresql-18-vchord` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.3 MiB | [postgresql-18-vchord_1.2.0-2PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-18-vchord_1.2.0-2PGSTY~bookworm_arm64.deb) |
| `postgresql-18-vchord` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 MiB | [postgresql-18-vchord_1.2.0-2PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-18-vchord_1.2.0-2PGSTY~trixie_amd64.deb) |
| `postgresql-18-vchord` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.3 MiB | [postgresql-18-vchord_1.2.0-2PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-18-vchord_1.2.0-2PGSTY~trixie_arm64.deb) |
| `postgresql-18-vchord` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.3 MiB | [postgresql-18-vchord_1.2.0-2PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-18-vchord_1.2.0-2PGSTY~jammy_amd64.deb) |
| `postgresql-18-vchord` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.8 MiB | [postgresql-18-vchord_1.2.0-2PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-18-vchord_1.2.0-2PGSTY~jammy_arm64.deb) |
| `postgresql-18-vchord` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.3 MiB | [postgresql-18-vchord_1.2.0-2PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-18-vchord_1.2.0-2PGSTY~noble_amd64.deb) |
| `postgresql-18-vchord` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-18-vchord_1.2.0-2PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-18-vchord_1.2.0-2PGSTY~noble_arm64.deb) |
| `postgresql-18-vchord` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.2 MiB | [postgresql-18-vchord_1.2.0-2PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-18-vchord_1.2.0-2PGSTY~resolute_amd64.deb) |
| `postgresql-18-vchord` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.7 MiB | [postgresql-18-vchord_1.2.0-2PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-18-vchord_1.2.0-2PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_17` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.5 MiB | [vchord_17-1.2.0-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_17-1.2.0-2PGSTY.el8.x86_64.rpm) |
| `vchord_17` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.8 MiB | [vchord_17-1.2.0-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_17-1.2.0-2PGSTY.el8.aarch64.rpm) |
| `vchord_17` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.5 MiB | [vchord_17-1.2.0-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_17-1.2.0-2PGSTY.el9.x86_64.rpm) |
| `vchord_17` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.0 MiB | [vchord_17-1.2.0-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_17-1.2.0-2PGSTY.el9.aarch64.rpm) |
| `vchord_17` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.5 MiB | [vchord_17-1.2.0-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_17-1.2.0-2PGSTY.el10.x86_64.rpm) |
| `vchord_17` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.0 MiB | [vchord_17-1.2.0-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_17-1.2.0-2PGSTY.el10.aarch64.rpm) |
| `postgresql-17-vchord` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.0 MiB | [postgresql-17-vchord_1.2.0-2PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-17-vchord_1.2.0-2PGSTY~bookworm_amd64.deb) |
| `postgresql-17-vchord` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.3 MiB | [postgresql-17-vchord_1.2.0-2PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-17-vchord_1.2.0-2PGSTY~bookworm_arm64.deb) |
| `postgresql-17-vchord` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 MiB | [postgresql-17-vchord_1.2.0-2PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-17-vchord_1.2.0-2PGSTY~trixie_amd64.deb) |
| `postgresql-17-vchord` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.4 MiB | [postgresql-17-vchord_1.2.0-2PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-17-vchord_1.2.0-2PGSTY~trixie_arm64.deb) |
| `postgresql-17-vchord` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.3 MiB | [postgresql-17-vchord_1.2.0-2PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-17-vchord_1.2.0-2PGSTY~jammy_amd64.deb) |
| `postgresql-17-vchord` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.8 MiB | [postgresql-17-vchord_1.2.0-2PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-17-vchord_1.2.0-2PGSTY~jammy_arm64.deb) |
| `postgresql-17-vchord` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.3 MiB | [postgresql-17-vchord_1.2.0-2PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-17-vchord_1.2.0-2PGSTY~noble_amd64.deb) |
| `postgresql-17-vchord` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-17-vchord_1.2.0-2PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-17-vchord_1.2.0-2PGSTY~noble_arm64.deb) |
| `postgresql-17-vchord` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.2 MiB | [postgresql-17-vchord_1.2.0-2PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-17-vchord_1.2.0-2PGSTY~resolute_amd64.deb) |
| `postgresql-17-vchord` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.8 MiB | [postgresql-17-vchord_1.2.0-2PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-17-vchord_1.2.0-2PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_16` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.4 MiB | [vchord_16-1.2.0-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_16-1.2.0-2PGSTY.el8.x86_64.rpm) |
| `vchord_16` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.7 MiB | [vchord_16-1.2.0-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_16-1.2.0-2PGSTY.el8.aarch64.rpm) |
| `vchord_16` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.4 MiB | [vchord_16-1.2.0-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_16-1.2.0-2PGSTY.el9.x86_64.rpm) |
| `vchord_16` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.9 MiB | [vchord_16-1.2.0-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_16-1.2.0-2PGSTY.el9.aarch64.rpm) |
| `vchord_16` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.4 MiB | [vchord_16-1.2.0-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_16-1.2.0-2PGSTY.el10.x86_64.rpm) |
| `vchord_16` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.9 MiB | [vchord_16-1.2.0-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_16-1.2.0-2PGSTY.el10.aarch64.rpm) |
| `postgresql-16-vchord` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.0 MiB | [postgresql-16-vchord_1.2.0-2PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-16-vchord_1.2.0-2PGSTY~bookworm_amd64.deb) |
| `postgresql-16-vchord` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.3 MiB | [postgresql-16-vchord_1.2.0-2PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-16-vchord_1.2.0-2PGSTY~bookworm_arm64.deb) |
| `postgresql-16-vchord` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 MiB | [postgresql-16-vchord_1.2.0-2PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-16-vchord_1.2.0-2PGSTY~trixie_amd64.deb) |
| `postgresql-16-vchord` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.3 MiB | [postgresql-16-vchord_1.2.0-2PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-16-vchord_1.2.0-2PGSTY~trixie_arm64.deb) |
| `postgresql-16-vchord` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.2 MiB | [postgresql-16-vchord_1.2.0-2PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-16-vchord_1.2.0-2PGSTY~jammy_amd64.deb) |
| `postgresql-16-vchord` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.7 MiB | [postgresql-16-vchord_1.2.0-2PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-16-vchord_1.2.0-2PGSTY~jammy_arm64.deb) |
| `postgresql-16-vchord` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.2 MiB | [postgresql-16-vchord_1.2.0-2PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-16-vchord_1.2.0-2PGSTY~noble_amd64.deb) |
| `postgresql-16-vchord` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-16-vchord_1.2.0-2PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-16-vchord_1.2.0-2PGSTY~noble_arm64.deb) |
| `postgresql-16-vchord` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.2 MiB | [postgresql-16-vchord_1.2.0-2PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-16-vchord_1.2.0-2PGSTY~resolute_amd64.deb) |
| `postgresql-16-vchord` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.7 MiB | [postgresql-16-vchord_1.2.0-2PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-16-vchord_1.2.0-2PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_15` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.4 MiB | [vchord_15-1.2.0-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_15-1.2.0-2PGSTY.el8.x86_64.rpm) |
| `vchord_15` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.7 MiB | [vchord_15-1.2.0-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_15-1.2.0-2PGSTY.el8.aarch64.rpm) |
| `vchord_15` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.4 MiB | [vchord_15-1.2.0-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_15-1.2.0-2PGSTY.el9.x86_64.rpm) |
| `vchord_15` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.9 MiB | [vchord_15-1.2.0-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_15-1.2.0-2PGSTY.el9.aarch64.rpm) |
| `vchord_15` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.4 MiB | [vchord_15-1.2.0-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_15-1.2.0-2PGSTY.el10.x86_64.rpm) |
| `vchord_15` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.9 MiB | [vchord_15-1.2.0-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_15-1.2.0-2PGSTY.el10.aarch64.rpm) |
| `postgresql-15-vchord` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.0 MiB | [postgresql-15-vchord_1.2.0-2PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-15-vchord_1.2.0-2PGSTY~bookworm_amd64.deb) |
| `postgresql-15-vchord` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.3 MiB | [postgresql-15-vchord_1.2.0-2PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-15-vchord_1.2.0-2PGSTY~bookworm_arm64.deb) |
| `postgresql-15-vchord` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 MiB | [postgresql-15-vchord_1.2.0-2PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-15-vchord_1.2.0-2PGSTY~trixie_amd64.deb) |
| `postgresql-15-vchord` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.3 MiB | [postgresql-15-vchord_1.2.0-2PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-15-vchord_1.2.0-2PGSTY~trixie_arm64.deb) |
| `postgresql-15-vchord` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.2 MiB | [postgresql-15-vchord_1.2.0-2PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-15-vchord_1.2.0-2PGSTY~jammy_amd64.deb) |
| `postgresql-15-vchord` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.7 MiB | [postgresql-15-vchord_1.2.0-2PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-15-vchord_1.2.0-2PGSTY~jammy_arm64.deb) |
| `postgresql-15-vchord` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.2 MiB | [postgresql-15-vchord_1.2.0-2PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-15-vchord_1.2.0-2PGSTY~noble_amd64.deb) |
| `postgresql-15-vchord` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-15-vchord_1.2.0-2PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-15-vchord_1.2.0-2PGSTY~noble_arm64.deb) |
| `postgresql-15-vchord` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.2 MiB | [postgresql-15-vchord_1.2.0-2PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-15-vchord_1.2.0-2PGSTY~resolute_amd64.deb) |
| `postgresql-15-vchord` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.7 MiB | [postgresql-15-vchord_1.2.0-2PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-15-vchord_1.2.0-2PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_14` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.4 MiB | [vchord_14-1.2.0-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_14-1.2.0-2PGSTY.el8.x86_64.rpm) |
| `vchord_14` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.7 MiB | [vchord_14-1.2.0-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_14-1.2.0-2PGSTY.el8.aarch64.rpm) |
| `vchord_14` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.4 MiB | [vchord_14-1.2.0-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_14-1.2.0-2PGSTY.el9.x86_64.rpm) |
| `vchord_14` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.9 MiB | [vchord_14-1.2.0-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_14-1.2.0-2PGSTY.el9.aarch64.rpm) |
| `vchord_14` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.4 MiB | [vchord_14-1.2.0-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_14-1.2.0-2PGSTY.el10.x86_64.rpm) |
| `vchord_14` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.9 MiB | [vchord_14-1.2.0-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_14-1.2.0-2PGSTY.el10.aarch64.rpm) |
| `postgresql-14-vchord` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.0 MiB | [postgresql-14-vchord_1.2.0-2PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-14-vchord_1.2.0-2PGSTY~bookworm_amd64.deb) |
| `postgresql-14-vchord` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.3 MiB | [postgresql-14-vchord_1.2.0-2PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-14-vchord_1.2.0-2PGSTY~bookworm_arm64.deb) |
| `postgresql-14-vchord` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 MiB | [postgresql-14-vchord_1.2.0-2PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-14-vchord_1.2.0-2PGSTY~trixie_amd64.deb) |
| `postgresql-14-vchord` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.3 MiB | [postgresql-14-vchord_1.2.0-2PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-14-vchord_1.2.0-2PGSTY~trixie_arm64.deb) |
| `postgresql-14-vchord` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.2 MiB | [postgresql-14-vchord_1.2.0-2PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-14-vchord_1.2.0-2PGSTY~jammy_amd64.deb) |
| `postgresql-14-vchord` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.7 MiB | [postgresql-14-vchord_1.2.0-2PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-14-vchord_1.2.0-2PGSTY~jammy_arm64.deb) |
| `postgresql-14-vchord` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.2 MiB | [postgresql-14-vchord_1.2.0-2PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-14-vchord_1.2.0-2PGSTY~noble_amd64.deb) |
| `postgresql-14-vchord` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-14-vchord_1.2.0-2PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-14-vchord_1.2.0-2PGSTY~noble_arm64.deb) |
| `postgresql-14-vchord` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.2 MiB | [postgresql-14-vchord_1.2.0-2PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-14-vchord_1.2.0-2PGSTY~resolute_amd64.deb) |
| `postgresql-14-vchord` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.7 MiB | [postgresql-14-vchord_1.2.0-2PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-14-vchord_1.2.0-2PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://vchord.pgsty.com/docs/vchord/" title="Repository" icon="link" subtitle="vchord.pgsty.com/docs/vchord/" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="vchord-1.2.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg vchord;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install vchord;		# install via package name, for the active PG version

pig install vchord -v 18;   # install for PG 18
pig install vchord -v 17;   # install for PG 17
pig install vchord -v 16;   # install for PG 16
pig install vchord -v 15;   # install for PG 15
pig install vchord -v 14;   # install for PG 14

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'vchord';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION vchord CASCADE; -- requires vector
```

## Usage

Sources:

- [Installation and upgrades](https://vchord.pgsty.com/docs/vchord/getting-started/installation/)
- [Partitioned vector indexes](https://vchord.pgsty.com/docs/vchord/usage/indexing/)
- [Graph indexes and result ordering](https://vchord.pgsty.com/docs/vchord/usage/graph-index/)
- [Per-index search defaults](https://vchord.pgsty.com/docs/vchord/usage/fallback-parameters/)
- [Quantized vector types](https://vchord.pgsty.com/docs/vchord/usage/quantization-types/)
- [Recall evaluation](https://vchord.pgsty.com/docs/vchord/usage/measure-recall/)
- [1.2.0 release notes](https://vchord.pgsty.com/blog/release/vchord-1.2.0/)
- [1.2.0 control file](https://github.com/pgsty/vchord/blob/1.2.0/vchord.control)

`vchord` 1.2.0 provides approximate vector search on PostgreSQL 14–18. It uses pgvector's types and adds the partitioned `vchordrq` index and the disk graph `vchordg` index. It requires the `vector` extension, shared preloading, and superuser privileges to install. Graph indexing is documented as a preview feature.

### Enable and Search

Add the library to the existing preload list and restart PostgreSQL:

```ini
shared_preload_libraries = 'vchord'
```

```sql
CREATE EXTENSION vchord CASCADE;
CREATE TABLE items (id bigserial PRIMARY KEY, embedding vector(3));
INSERT INTO items (embedding) VALUES ('[1,2,3]'), ('[4,5,6]'), ('[3,1,2]');
CREATE INDEX items_embedding_idx ON items
USING vchordrq (embedding vector_l2_ops);

SELECT id, embedding <-> '[3,1,2]' AS distance
FROM items ORDER BY embedding <-> '[3,1,2]' LIMIT 5;
```

Small datasets can start with an unpartitioned index. For larger datasets, choose the partition hierarchy through TOML `build.internal.lists`, then set `vchordrq.probes` to the number of partitions to search at each level. More probes trade query work for recall; the probe list must match the hierarchy. The graph alternative uses the same column and distance operator classes:

```sql
CREATE INDEX items_embedding_graph_idx ON items
USING vchordg (embedding vector_l2_ops)
WITH (options = $$
m = 32
ef_construction = 64
$$, ef_search = 128);
```

Use the graph definition as an alternative to the partitioned index above. The recall example below assumes only the partitioned index exists; inspect the query plan when comparing index methods.

### Types, Operators, and Tuning

| Object | Purpose |
| --- | --- |
| `vector_l2_ops`, `<->` | Euclidean distance, ordered ascending. |
| `vector_ip_ops`, `<#>` | Negative inner product, ordered ascending. |
| `vector_cosine_ops`, `<=>` | Cosine distance, ordered ascending. |
| `halfvec` | Half-precision vectors, with corresponding `halfvec_*_ops` classes. |
| `rabitq4`, `rabitq8` | Quantized vector storage; conversion loses precision. |
| `sphere`, `<<->>`, `<<#>>`, `<<=>>` | Distance-radius predicates for L2, inner product, and cosine search. |
| `vchordrq.epsilon` | Reranking tradeoff; default 2.5 in 1.2.0, previously 1.9. |
| `vchordg.ef_search` | Graph search breadth; default 64. |
| `vchordrq_prewarm`, `vchordg_prewarm` | Load hot index pages; require table ownership or table-level `SELECT`. |

```sql
SELECT id FROM items
WHERE embedding <<->> sphere('[1,2,3]'::vector, 0.5);

SELECT dequantize_to_vector(quantize_to_rabitq8('[1,2,3]'::vector));
SELECT dequantize_to_halfvec(quantize_to_rabitq4('[1,2,3]'::halfvec));
```

Sphere radii use the distance operator's units. PostgreSQL rechecks graph sphere predicates, but approximate search can still miss matching rows. Dequantization cannot recover the original vector exactly.

Per-index `probes`, `epsilon`, and `ef_search` options supply search defaults. Explicit GUC settings, including database or server settings, override them. Resetting a session setting restores its configured default, which can still override the index.

### Ordering and Recall

Graph scans can discover closer rows after returning earlier candidates; nearly equal half-precision distances can also appear out of order. Increasing the search breadth reduces some ordering errors but does not guarantee strict ordering. Materialize a larger candidate set and sort it explicitly when ordering matters:

```sql
WITH candidates AS MATERIALIZED (
    SELECT id, embedding <-> '[3,1,2]' AS distance
    FROM items ORDER BY embedding <-> '[3,1,2]' LIMIT 100
)
SELECT * FROM candidates ORDER BY distance LIMIT 10;
```

This sorts the retrieved candidates; recall still depends on the candidate count and search parameters. For partitioned indexes, compare approximate results with an exact baseline:

```sql
SELECT vchordrq_evaluate_query_recall(
    query => $$SELECT ctid FROM items
                ORDER BY embedding <-> '[3,1,2]' LIMIT 5$$,
    exact_search => true
);
```

The query must return `ctid`; keep sequential scans enabled for the exact pass. Multi-vector MaxSim queries using `@#` require `exact_search => true` for recall evaluation.

### Maintenance and Upgrade

Keep normal vacuuming enabled: graph searches exclude deleted vertices after `VACUUM`, and index maintenance updates tuple/page statistics. `ANALYZE` alone does not rebuild or compact the index.

Install matching 1.2.0 library and SQL files, restart PostgreSQL to load the new library, then update each database as the extension owner or an administrator:

```sql
ALTER EXTENSION vchord UPDATE TO '1.2.0';
```

The release provides upgrade paths from 0.2.0 through 1.1.1 without reindexing. Recheck representative rankings after upgrading: 1.2.0 fixes MaxSim candidate estimates, graph sphere filters, deleted-vertex visibility, and quantization edge cases, and raises the default reranking budget. The extension is relocatable; include its chosen schema in the search path or qualify its objects.
