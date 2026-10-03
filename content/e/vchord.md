---
title: "vchord"
linkTitle: "vchord"
description: "Vector database plugin for Postgres, written in Rust"
weight: 1810
categories: ["RAG"]
languages: ["Rust"]
licenses: ["AGPL-3.0"]
repos: ["PIGSTY"]
page_width: full
---

[**vchord**](https://github.com/supervc-stack/VectorChord) : Vector database plugin for Postgres, written in Rust


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1810** | {{< badge content="vchord" link="https://github.com/supervc-stack/VectorChord" >}} | {{< ext "vchord" >}} | `1.1.1` | {{< category "RAG" >}} | {{< license "AGPL-3.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "vector" >}} |
|   **See Also**    | {{< ext "vector" >}} {{< ext "vectorscale" >}} {{< ext "pgcontext" >}} {{< ext "vectorize" >}} {{< ext "pg_rrf" >}} {{< ext "pg_search" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_bestmatch" >}} {{< ext "pgml" >}} {{< ext "pg4ml" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `vchord` | `vector` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.1` | {{< bg "18" "vchord_18" "green" >}} {{< bg "17" "vchord_17" "green" >}} {{< bg "16" "vchord_16" "green" >}} {{< bg "15" "vchord_15" "green" >}} {{< bg "14" "vchord_14" "green" >}} | `vchord_$v` | `pgvector_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.1` | {{< bg "18" "postgresql-18-vchord" "green" >}} {{< bg "17" "postgresql-17-vchord" "green" >}} {{< bg "16" "postgresql-16-vchord" "green" >}} {{< bg "15" "postgresql-15-vchord" "green" >}} {{< bg "14" "postgresql-14-vchord" "green" >}} | `postgresql-$v-vchord` | `postgresql-$v-pgvector` |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "vchord_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "vchord_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "vchord_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "vchord_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "vchord_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "vchord_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "vchord_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-vchord : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-vchord : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-vchord : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-vchord : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-vchord : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-vchord : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-vchord : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-vchord : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-vchord : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-vchord : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_18` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.1 MiB | [vchord_18-1.1.1-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_18-1.1.1-3PIGSTY.el8.x86_64.rpm) |
| `vchord_18` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.7 MiB | [vchord_18-1.1.1-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_18-1.1.1-3PIGSTY.el8.aarch64.rpm) |
| `vchord_18` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.1 MiB | [vchord_18-1.1.1-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_18-1.1.1-3PIGSTY.el9.x86_64.rpm) |
| `vchord_18` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.9 MiB | [vchord_18-1.1.1-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_18-1.1.1-3PIGSTY.el9.aarch64.rpm) |
| `vchord_18` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.0 MiB | [vchord_18-1.1.1-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_18-1.1.1-3PIGSTY.el10.x86_64.rpm) |
| `vchord_18` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.8 MiB | [vchord_18-1.1.1-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_18-1.1.1-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-vchord` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.8 MiB | [postgresql-18-vchord_1.1.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-18-vchord_1.1.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-vchord` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-18-vchord_1.1.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-18-vchord_1.1.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-vchord` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.9 MiB | [postgresql-18-vchord_1.1.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-18-vchord_1.1.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-18-vchord` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.4 MiB | [postgresql-18-vchord_1.1.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-18-vchord_1.1.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-18-vchord` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.1 MiB | [postgresql-18-vchord_1.1.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-18-vchord_1.1.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-18-vchord` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.9 MiB | [postgresql-18-vchord_1.1.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-18-vchord_1.1.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-18-vchord` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.1 MiB | [postgresql-18-vchord_1.1.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-18-vchord_1.1.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-18-vchord` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.9 MiB | [postgresql-18-vchord_1.1.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-18-vchord_1.1.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-18-vchord` | `1.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.1 MiB | [postgresql-18-vchord_1.1.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-18-vchord_1.1.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-18-vchord` | `1.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.9 MiB | [postgresql-18-vchord_1.1.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-18-vchord_1.1.1-3PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_17` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.1 MiB | [vchord_17-1.1.1-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_17-1.1.1-3PIGSTY.el8.x86_64.rpm) |
| `vchord_17` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.7 MiB | [vchord_17-1.1.1-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_17-1.1.1-3PIGSTY.el8.aarch64.rpm) |
| `vchord_17` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.1 MiB | [vchord_17-1.1.1-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_17-1.1.1-3PIGSTY.el9.x86_64.rpm) |
| `vchord_17` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.9 MiB | [vchord_17-1.1.1-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_17-1.1.1-3PIGSTY.el9.aarch64.rpm) |
| `vchord_17` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.0 MiB | [vchord_17-1.1.1-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_17-1.1.1-3PIGSTY.el10.x86_64.rpm) |
| `vchord_17` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.8 MiB | [vchord_17-1.1.1-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_17-1.1.1-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-vchord` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.8 MiB | [postgresql-17-vchord_1.1.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-17-vchord_1.1.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-vchord` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-17-vchord_1.1.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-17-vchord_1.1.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-vchord` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.8 MiB | [postgresql-17-vchord_1.1.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-17-vchord_1.1.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-17-vchord` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.4 MiB | [postgresql-17-vchord_1.1.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-17-vchord_1.1.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-17-vchord` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.1 MiB | [postgresql-17-vchord_1.1.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-17-vchord_1.1.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-17-vchord` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.9 MiB | [postgresql-17-vchord_1.1.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-17-vchord_1.1.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-17-vchord` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.1 MiB | [postgresql-17-vchord_1.1.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-17-vchord_1.1.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-17-vchord` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.9 MiB | [postgresql-17-vchord_1.1.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-17-vchord_1.1.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-17-vchord` | `1.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.1 MiB | [postgresql-17-vchord_1.1.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-17-vchord_1.1.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-17-vchord` | `1.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.9 MiB | [postgresql-17-vchord_1.1.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-17-vchord_1.1.1-3PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_16` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.1 MiB | [vchord_16-1.1.1-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_16-1.1.1-3PIGSTY.el8.x86_64.rpm) |
| `vchord_16` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.6 MiB | [vchord_16-1.1.1-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_16-1.1.1-3PIGSTY.el8.aarch64.rpm) |
| `vchord_16` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.1 MiB | [vchord_16-1.1.1-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_16-1.1.1-3PIGSTY.el9.x86_64.rpm) |
| `vchord_16` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [vchord_16-1.1.1-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_16-1.1.1-3PIGSTY.el9.aarch64.rpm) |
| `vchord_16` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.0 MiB | [vchord_16-1.1.1-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_16-1.1.1-3PIGSTY.el10.x86_64.rpm) |
| `vchord_16` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.8 MiB | [vchord_16-1.1.1-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_16-1.1.1-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-vchord` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.8 MiB | [postgresql-16-vchord_1.1.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-16-vchord_1.1.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-vchord` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-16-vchord_1.1.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-16-vchord_1.1.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-vchord` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.8 MiB | [postgresql-16-vchord_1.1.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-16-vchord_1.1.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-16-vchord` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.4 MiB | [postgresql-16-vchord_1.1.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-16-vchord_1.1.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-16-vchord` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.1 MiB | [postgresql-16-vchord_1.1.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-16-vchord_1.1.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-16-vchord` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.8 MiB | [postgresql-16-vchord_1.1.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-16-vchord_1.1.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-16-vchord` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.1 MiB | [postgresql-16-vchord_1.1.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-16-vchord_1.1.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-16-vchord` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-16-vchord_1.1.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-16-vchord_1.1.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-16-vchord` | `1.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.0 MiB | [postgresql-16-vchord_1.1.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-16-vchord_1.1.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-16-vchord` | `1.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.8 MiB | [postgresql-16-vchord_1.1.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-16-vchord_1.1.1-3PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_15` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.1 MiB | [vchord_15-1.1.1-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_15-1.1.1-3PIGSTY.el8.x86_64.rpm) |
| `vchord_15` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.6 MiB | [vchord_15-1.1.1-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_15-1.1.1-3PIGSTY.el8.aarch64.rpm) |
| `vchord_15` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.1 MiB | [vchord_15-1.1.1-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_15-1.1.1-3PIGSTY.el9.x86_64.rpm) |
| `vchord_15` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [vchord_15-1.1.1-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_15-1.1.1-3PIGSTY.el9.aarch64.rpm) |
| `vchord_15` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.0 MiB | [vchord_15-1.1.1-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_15-1.1.1-3PIGSTY.el10.x86_64.rpm) |
| `vchord_15` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.8 MiB | [vchord_15-1.1.1-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_15-1.1.1-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-vchord` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.8 MiB | [postgresql-15-vchord_1.1.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-15-vchord_1.1.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-vchord` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-15-vchord_1.1.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-15-vchord_1.1.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-vchord` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.8 MiB | [postgresql-15-vchord_1.1.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-15-vchord_1.1.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-15-vchord` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.4 MiB | [postgresql-15-vchord_1.1.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-15-vchord_1.1.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-15-vchord` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.1 MiB | [postgresql-15-vchord_1.1.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-15-vchord_1.1.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-15-vchord` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.8 MiB | [postgresql-15-vchord_1.1.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-15-vchord_1.1.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-15-vchord` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.1 MiB | [postgresql-15-vchord_1.1.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-15-vchord_1.1.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-15-vchord` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-15-vchord_1.1.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-15-vchord_1.1.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-15-vchord` | `1.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.0 MiB | [postgresql-15-vchord_1.1.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-15-vchord_1.1.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-15-vchord` | `1.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.8 MiB | [postgresql-15-vchord_1.1.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-15-vchord_1.1.1-3PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_14` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.1 MiB | [vchord_14-1.1.1-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_14-1.1.1-3PIGSTY.el8.x86_64.rpm) |
| `vchord_14` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.6 MiB | [vchord_14-1.1.1-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_14-1.1.1-3PIGSTY.el8.aarch64.rpm) |
| `vchord_14` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.1 MiB | [vchord_14-1.1.1-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_14-1.1.1-3PIGSTY.el9.x86_64.rpm) |
| `vchord_14` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [vchord_14-1.1.1-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_14-1.1.1-3PIGSTY.el9.aarch64.rpm) |
| `vchord_14` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.0 MiB | [vchord_14-1.1.1-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_14-1.1.1-3PIGSTY.el10.x86_64.rpm) |
| `vchord_14` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.8 MiB | [vchord_14-1.1.1-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_14-1.1.1-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-vchord` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.8 MiB | [postgresql-14-vchord_1.1.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-14-vchord_1.1.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-vchord` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-14-vchord_1.1.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-14-vchord_1.1.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-vchord` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.8 MiB | [postgresql-14-vchord_1.1.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-14-vchord_1.1.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-14-vchord` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.4 MiB | [postgresql-14-vchord_1.1.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-14-vchord_1.1.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-14-vchord` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.1 MiB | [postgresql-14-vchord_1.1.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-14-vchord_1.1.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-14-vchord` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.8 MiB | [postgresql-14-vchord_1.1.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-14-vchord_1.1.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-14-vchord` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.0 MiB | [postgresql-14-vchord_1.1.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-14-vchord_1.1.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-14-vchord` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-14-vchord_1.1.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-14-vchord_1.1.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-14-vchord` | `1.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.0 MiB | [postgresql-14-vchord_1.1.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-14-vchord_1.1.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-14-vchord` | `1.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.8 MiB | [postgresql-14-vchord_1.1.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vchord/postgresql-14-vchord_1.1.1-3PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supervc-stack/VectorChord" title="Repository" icon="github" subtitle="github.com/supervc-stack/VectorChord" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="VectorChord-1.1.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg vchord;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install vchord;		# install via package name, for the active PG version

pig install vchord -v 18;   # install for PG 18
pig install vchord -v 17;   # install for PG 17
pig install vchord -v 16;   # install for PG 16
pig install vchord -v 15;   # install for PG 15
pig install vchord -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'vchord';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION vchord CASCADE; -- requires vector
```

## Usage

Sources:

- [1.1.1 README](https://github.com/supervc-stack/VectorChord/blob/1.1.1/README.md)
- [Control and dependency](https://github.com/supervc-stack/VectorChord/blob/1.1.1/vchord.control)
- [Preload requirement](https://github.com/supervc-stack/VectorChord/blob/1.1.1/src/lib.rs)
- [1.1.1 SQL objects](https://github.com/supervc-stack/VectorChord/blob/1.1.1/sql/install/vchord--1.1.1.sql)
- [Query settings](https://github.com/supervc-stack/VectorChord/blob/1.1.1/src/index/gucs.rs)
- [1.1.1 migration](https://github.com/supervc-stack/VectorChord/blob/1.1.1/sql/upgrade/vchord--1.1.0--1.1.1.sql)
- [1.1.1 release notes](https://github.com/supervc-stack/VectorChord/releases/tag/1.1.1)

`vchord` adds approximate vector indexes to PostgreSQL using pgvector's types. It provides the partition-based `vchordrq` and graph-based `vchordg` access methods. The extension requires `vector`, shared preloading, and superuser privileges to create.

### Create and Query an Index

Add the library to the existing preload list, preserving other entries, and restart PostgreSQL:

```conf
shared_preload_libraries = 'vchord'
```

```sql
CREATE EXTENSION vchord CASCADE;
CREATE TABLE items (id bigserial PRIMARY KEY, embedding vector(3));
INSERT INTO items(embedding) VALUES ('[1,2,3]'), ('[4,5,6]');
CREATE INDEX items_embedding_idx ON items
USING vchordrq (embedding vector_l2_ops);

SELECT id FROM items ORDER BY embedding <-> '[3,1,2]' LIMIT 5;
SELECT vchordrq_prewarm('items_embedding_idx'::regclass);
```

Use `vector_l2_ops` with `<->`, `vector_ip_ops` with `<#>`, and `vector_cosine_ops` with `<=>`. The inner-product operator returns a negative value for ascending index ordering. The same operator classes can be used with the graph access method; choose one index design for the workload:

```sql
CREATE INDEX items_embedding_graph_idx ON items
USING vchordg (embedding vector_l2_ops);
```

### Range Queries and Tuning

The extension supplies explicit sphere predicates for range search:

```sql
SELECT id FROM items
WHERE embedding <<->> sphere('[1,2,3]'::vector, 0.5);

SET vchordrq.probes = '100';
SET vchordrq.epsilon = 1.9;
SET vchordg.ef_search = 64;
```

`<<->>`, `<<#>>`, and `<<=>>` are sphere predicates for L2, inner product, and cosine metrics. Probe counts depend on the partition layout; tune them with representative data. The epsilon setting controls the reranking tradeoff. The graph search setting controls its candidate search breadth. Both index methods are approximate: check recall, filters, and query plans before choosing settings.

### Quantization in 1.1.1

`rabitq8` and `rabitq4` store quantized vectors. `quantize_to_rabitq8` and `quantize_to_rabitq4` accept `vector` or `halfvec`. Version 1.1.1 adds `dequantize_to_vector` and `dequantize_to_halfvec` overloads for both quantized types:

```sql
SELECT dequantize_to_vector(quantize_to_rabitq8('[1,2,3]'::vector));
SELECT dequantize_to_halfvec(quantize_to_rabitq4('[1,2,3]'::halfvec));
```

Quantization loses precision; dequantization returns an approximation. The release also replaces the quantization implementation. Install matching library and SQL files, restart for the preloaded library, then update each database:

```sql
ALTER EXTENSION vchord UPDATE TO '1.1.1';
```

The 1.1.0-to-1.1.1 script adds these four conversion overloads and declares no index-format migration. Earlier-version upgrade requirements depend on the starting version. Index construction and prewarming consume resources; schedule them for the dataset size. `vchordg_prewarm` is the corresponding graph-index helper. The control is relocatable, so qualify extension objects or include their installation schema in the search path when installed outside the usual schema.
