---
title: "pgroonga"
linkTitle: "pgroonga"
description: "Use Groonga as index, fast full text search platform for all languages!"
weight: 2110
categories: ["FTS"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pgroonga**](https://github.com/pgroonga/pgroonga) : Use Groonga as index, fast full text search platform for all languages!


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2110** | {{< badge content="pgroonga" link="https://github.com/pgroonga/pgroonga" >}} | {{< ext "pgroonga" >}} | `4.0.9` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pg_textsearch" >}} {{< ext "pg_fts" >}} {{< ext "pg_bigm" >}} {{< ext "zhparser" >}} {{< ext "pg_tokenizer" >}} {{< ext "pg_cjk_parser" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_bestmatch" >}} {{< ext "pg_jieba" >}} {{< ext "dict_xsyn" >}} {{< ext "unaccent" >}} |
|    **Siblings**   | {{< ext "pgroonga_database" >}} |

> [!Note] require xxHash vendor repo to build


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.0.9` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgroonga` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.0.9` | {{< bg "18" "pgroonga_18" "green" >}} {{< bg "17" "pgroonga_17" "green" >}} {{< bg "16" "pgroonga_16" "green" >}} {{< bg "15" "pgroonga_15" "green" >}} {{< bg "14" "pgroonga_14" "green" >}} | `pgroonga_$v` | `groonga-libs` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.0.9` | {{< bg "18" "postgresql-18-pgroonga" "green" >}} {{< bg "17" "postgresql-17-pgroonga" "green" >}} {{< bg "16" "postgresql-16-pgroonga" "green" >}} {{< bg "15" "postgresql-15-pgroonga" "green" >}} {{< bg "14" "postgresql-14-pgroonga" "green" >}} | `postgresql-$v-pgroonga` | `libgroonga0` |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgroonga_18` | `4.0.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 242.0 KiB | [pgroonga_18-4.0.9-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_18-4.0.9-1PGSTY.el8.x86_64.rpm) |
| `pgroonga_18` | `4.0.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 228.8 KiB | [pgroonga_18-4.0.9-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_18-4.0.9-1PGSTY.el8.aarch64.rpm) |
| `pgroonga_18` | `4.0.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 246.2 KiB | [pgroonga_18-4.0.9-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_18-4.0.9-1PGSTY.el9.x86_64.rpm) |
| `pgroonga_18` | `4.0.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 238.3 KiB | [pgroonga_18-4.0.9-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_18-4.0.9-1PGSTY.el9.aarch64.rpm) |
| `pgroonga_18` | `4.0.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 248.8 KiB | [pgroonga_18-4.0.9-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgroonga_18-4.0.9-1PGSTY.el10.x86_64.rpm) |
| `pgroonga_18` | `4.0.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 239.5 KiB | [pgroonga_18-4.0.9-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgroonga_18-4.0.9-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgroonga` | `4.0.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 181.1 KiB | [postgresql-18-pgroonga_4.0.9-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-18-pgroonga_4.0.9-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgroonga` | `4.0.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 163.3 KiB | [postgresql-18-pgroonga_4.0.9-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-18-pgroonga_4.0.9-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgroonga` | `4.0.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 182.2 KiB | [postgresql-18-pgroonga_4.0.9-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-18-pgroonga_4.0.9-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pgroonga` | `4.0.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 163.8 KiB | [postgresql-18-pgroonga_4.0.9-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-18-pgroonga_4.0.9-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pgroonga` | `4.0.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 196.8 KiB | [postgresql-18-pgroonga_4.0.9-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-18-pgroonga_4.0.9-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pgroonga` | `4.0.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 190.1 KiB | [postgresql-18-pgroonga_4.0.9-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-18-pgroonga_4.0.9-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pgroonga` | `4.0.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 191.4 KiB | [postgresql-18-pgroonga_4.0.9-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-18-pgroonga_4.0.9-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pgroonga` | `4.0.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 185.2 KiB | [postgresql-18-pgroonga_4.0.9-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-18-pgroonga_4.0.9-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pgroonga` | `4.0.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 193.6 KiB | [postgresql-18-pgroonga_4.0.9-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgroonga/postgresql-18-pgroonga_4.0.9-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pgroonga` | `4.0.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 184.1 KiB | [postgresql-18-pgroonga_4.0.9-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgroonga/postgresql-18-pgroonga_4.0.9-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgroonga_17` | `4.0.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 242.2 KiB | [pgroonga_17-4.0.9-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_17-4.0.9-1PGSTY.el8.x86_64.rpm) |
| `pgroonga_17` | `4.0.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 228.8 KiB | [pgroonga_17-4.0.9-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_17-4.0.9-1PGSTY.el8.aarch64.rpm) |
| `pgroonga_17` | `4.0.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 246.0 KiB | [pgroonga_17-4.0.9-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_17-4.0.9-1PGSTY.el9.x86_64.rpm) |
| `pgroonga_17` | `4.0.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 238.1 KiB | [pgroonga_17-4.0.9-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_17-4.0.9-1PGSTY.el9.aarch64.rpm) |
| `pgroonga_17` | `4.0.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 248.9 KiB | [pgroonga_17-4.0.9-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgroonga_17-4.0.9-1PGSTY.el10.x86_64.rpm) |
| `pgroonga_17` | `4.0.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 239.3 KiB | [pgroonga_17-4.0.9-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgroonga_17-4.0.9-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgroonga` | `4.0.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 181.2 KiB | [postgresql-17-pgroonga_4.0.9-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.9-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgroonga` | `4.0.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 163.1 KiB | [postgresql-17-pgroonga_4.0.9-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.9-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgroonga` | `4.0.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 181.8 KiB | [postgresql-17-pgroonga_4.0.9-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.9-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pgroonga` | `4.0.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 163.7 KiB | [postgresql-17-pgroonga_4.0.9-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.9-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pgroonga` | `4.0.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 197.0 KiB | [postgresql-17-pgroonga_4.0.9-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.9-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pgroonga` | `4.0.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 190.0 KiB | [postgresql-17-pgroonga_4.0.9-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.9-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pgroonga` | `4.0.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 191.0 KiB | [postgresql-17-pgroonga_4.0.9-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.9-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pgroonga` | `4.0.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 185.0 KiB | [postgresql-17-pgroonga_4.0.9-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.9-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pgroonga` | `4.0.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 193.6 KiB | [postgresql-17-pgroonga_4.0.9-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.9-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pgroonga` | `4.0.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 184.0 KiB | [postgresql-17-pgroonga_4.0.9-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.9-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgroonga_16` | `4.0.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 239.3 KiB | [pgroonga_16-4.0.9-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_16-4.0.9-1PGSTY.el8.x86_64.rpm) |
| `pgroonga_16` | `4.0.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 226.6 KiB | [pgroonga_16-4.0.9-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_16-4.0.9-1PGSTY.el8.aarch64.rpm) |
| `pgroonga_16` | `4.0.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 243.7 KiB | [pgroonga_16-4.0.9-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_16-4.0.9-1PGSTY.el9.x86_64.rpm) |
| `pgroonga_16` | `4.0.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 236.1 KiB | [pgroonga_16-4.0.9-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_16-4.0.9-1PGSTY.el9.aarch64.rpm) |
| `pgroonga_16` | `4.0.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 246.6 KiB | [pgroonga_16-4.0.9-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgroonga_16-4.0.9-1PGSTY.el10.x86_64.rpm) |
| `pgroonga_16` | `4.0.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 237.2 KiB | [pgroonga_16-4.0.9-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgroonga_16-4.0.9-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgroonga` | `4.0.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 178.7 KiB | [postgresql-16-pgroonga_4.0.9-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.9-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgroonga` | `4.0.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 161.4 KiB | [postgresql-16-pgroonga_4.0.9-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.9-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgroonga` | `4.0.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 180.1 KiB | [postgresql-16-pgroonga_4.0.9-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.9-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pgroonga` | `4.0.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 161.6 KiB | [postgresql-16-pgroonga_4.0.9-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.9-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pgroonga` | `4.0.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 194.6 KiB | [postgresql-16-pgroonga_4.0.9-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.9-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pgroonga` | `4.0.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 187.5 KiB | [postgresql-16-pgroonga_4.0.9-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.9-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pgroonga` | `4.0.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 189.0 KiB | [postgresql-16-pgroonga_4.0.9-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.9-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pgroonga` | `4.0.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 182.6 KiB | [postgresql-16-pgroonga_4.0.9-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.9-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pgroonga` | `4.0.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 191.3 KiB | [postgresql-16-pgroonga_4.0.9-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.9-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pgroonga` | `4.0.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 182.4 KiB | [postgresql-16-pgroonga_4.0.9-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.9-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgroonga_15` | `4.0.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 239.0 KiB | [pgroonga_15-4.0.9-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_15-4.0.9-1PGSTY.el8.x86_64.rpm) |
| `pgroonga_15` | `4.0.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 226.1 KiB | [pgroonga_15-4.0.9-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_15-4.0.9-1PGSTY.el8.aarch64.rpm) |
| `pgroonga_15` | `4.0.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 242.8 KiB | [pgroonga_15-4.0.9-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_15-4.0.9-1PGSTY.el9.x86_64.rpm) |
| `pgroonga_15` | `4.0.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 235.4 KiB | [pgroonga_15-4.0.9-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_15-4.0.9-1PGSTY.el9.aarch64.rpm) |
| `pgroonga_15` | `4.0.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 246.4 KiB | [pgroonga_15-4.0.9-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgroonga_15-4.0.9-1PGSTY.el10.x86_64.rpm) |
| `pgroonga_15` | `4.0.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 236.8 KiB | [pgroonga_15-4.0.9-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgroonga_15-4.0.9-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgroonga` | `4.0.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 178.9 KiB | [postgresql-15-pgroonga_4.0.9-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.9-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgroonga` | `4.0.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 161.1 KiB | [postgresql-15-pgroonga_4.0.9-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.9-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgroonga` | `4.0.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 180.4 KiB | [postgresql-15-pgroonga_4.0.9-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.9-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pgroonga` | `4.0.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 161.8 KiB | [postgresql-15-pgroonga_4.0.9-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.9-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pgroonga` | `4.0.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 193.9 KiB | [postgresql-15-pgroonga_4.0.9-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.9-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pgroonga` | `4.0.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 187.2 KiB | [postgresql-15-pgroonga_4.0.9-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.9-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pgroonga` | `4.0.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 188.5 KiB | [postgresql-15-pgroonga_4.0.9-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.9-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pgroonga` | `4.0.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 182.1 KiB | [postgresql-15-pgroonga_4.0.9-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.9-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pgroonga` | `4.0.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 191.1 KiB | [postgresql-15-pgroonga_4.0.9-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.9-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pgroonga` | `4.0.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 181.7 KiB | [postgresql-15-pgroonga_4.0.9-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.9-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgroonga_14` | `4.0.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 221.0 KiB | [pgroonga_14-4.0.9-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_14-4.0.9-1PGSTY.el8.x86_64.rpm) |
| `pgroonga_14` | `4.0.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 210.6 KiB | [pgroonga_14-4.0.9-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_14-4.0.9-1PGSTY.el8.aarch64.rpm) |
| `pgroonga_14` | `4.0.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 225.4 KiB | [pgroonga_14-4.0.9-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_14-4.0.9-1PGSTY.el9.x86_64.rpm) |
| `pgroonga_14` | `4.0.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 219.3 KiB | [pgroonga_14-4.0.9-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_14-4.0.9-1PGSTY.el9.aarch64.rpm) |
| `pgroonga_14` | `4.0.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 228.9 KiB | [pgroonga_14-4.0.9-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgroonga_14-4.0.9-1PGSTY.el10.x86_64.rpm) |
| `pgroonga_14` | `4.0.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 220.4 KiB | [pgroonga_14-4.0.9-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgroonga_14-4.0.9-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgroonga` | `4.0.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 163.8 KiB | [postgresql-14-pgroonga_4.0.9-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.9-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgroonga` | `4.0.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 147.6 KiB | [postgresql-14-pgroonga_4.0.9-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.9-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgroonga` | `4.0.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 164.8 KiB | [postgresql-14-pgroonga_4.0.9-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.9-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pgroonga` | `4.0.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 148.9 KiB | [postgresql-14-pgroonga_4.0.9-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.9-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pgroonga` | `4.0.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 177.4 KiB | [postgresql-14-pgroonga_4.0.9-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.9-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pgroonga` | `4.0.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 171.3 KiB | [postgresql-14-pgroonga_4.0.9-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.9-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pgroonga` | `4.0.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 172.5 KiB | [postgresql-14-pgroonga_4.0.9-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.9-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pgroonga` | `4.0.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 167.0 KiB | [postgresql-14-pgroonga_4.0.9-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.9-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pgroonga` | `4.0.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 174.6 KiB | [postgresql-14-pgroonga_4.0.9-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.9-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pgroonga` | `4.0.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 166.6 KiB | [postgresql-14-pgroonga_4.0.9-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.9-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgroonga/pgroonga" title="Repository" icon="github" subtitle="github.com/pgroonga/pgroonga" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgroonga-4.0.9.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pgroonga;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pgroonga;		# install via package name, for the active PG version

pig install pgroonga -v 18;   # install for PG 18
pig install pgroonga -v 17;   # install for PG 17
pig install pgroonga -v 16;   # install for PG 16
pig install pgroonga -v 15;   # install for PG 15
pig install pgroonga -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgroonga;
```

## Usage

Sources:

- [Version 4.0.9 SQL](https://github.com/pgroonga/pgroonga/blob/4.0.9/data/pgroonga.sql)
- [Version 4.0.9 control](https://github.com/pgroonga/pgroonga/blob/4.0.9/pgroonga.control)
- [Official tutorial](https://pgroonga.github.io/tutorial/)
- [Version 4.0.9 release](https://github.com/pgroonga/pgroonga/releases/tag/4.0.9)
- [Upgrade guidance](https://pgroonga.github.io/upgrade/)

`pgroonga` 4.0.9 provides Groonga-backed indexes for multilingual full-text search. It installs the `pgroonga` access method and SQL operators; ordinary use does not require shared preload.

### Core Workflow

Install compatible PGroonga and Groonga libraries, then create the extension as an administrator:

```sql
CREATE EXTENSION pgroonga;
CREATE TABLE search_notes (id bigint PRIMARY KEY, body text);
CREATE INDEX search_notes_body_idx ON search_notes USING pgroonga (body);
INSERT INTO search_notes VALUES (1, 'PostgreSQL supports full text search');
SELECT id, body FROM search_notes WHERE body &@ 'PostgreSQL';
SELECT id, body, pgroonga_score(tableoid, ctid) AS score
FROM search_notes WHERE body &@~ 'PostgreSQL OR Groonga'
ORDER BY score DESC;
```

### Important Objects

- `&@` matches a keyword; `&@~` accepts Groonga query syntax. Supported LIKE/ILIKE searches can also use the index, with rechecks where required.
- `pgroonga_score(tableoid, ctid)` retrieves search scores. Confirm the intended index plan when using score-based ordering.
- `pgroonga_highlight_html()` and `pgroonga_query_extract_keywords()` produce highlighted search results; `pgroonga_snippet_html()` provides surrounding text.
- New in 4.0.9, `pgroonga_physical_table_names(partitioned_index, prefix)` returns a text array of Groonga command arguments identifying the physical tables behind partition indexes. This release also increments `pg_stat_user_indexes.idx_scan` for PGroonga scans.

### Maintenance and Privileges

The 4.0.9 control file declares neither trusted installation nor relocatability; do not assume ordinary users can install it or move it between schemas. Index creation and queries follow the relevant table privileges. Match the extension and Groonga libraries to the target PostgreSQL build and follow upstream upgrade guidance before replacing binaries.

PGroonga manages derived index files in addition to table data. Plan disk capacity and backup/recovery procedures accordingly. Use REINDEX for index repair where appropriate. The separate `pgroonga_database` module is a recovery tool for damaged internal Groonga databases and is not needed for normal searches. Do not disable sequential scans globally merely to force an index in production.

### Pigsty Runtime Compatibility

The current Pigsty EL8 and EL9 packages have a confirmed coexistence limit with PostGIS Raster on both x86_64 and aarch64, across PostgreSQL 14–18. Groonga uses Arrow 22, while the tested GDAL/Raster stacks use Arrow 8 on EL8 and Arrow 9 on EL9. Loading both stacks into one PostgreSQL backend can crash it; a successful SQL query does not prove a normal backend exit.

Changing load order is not a complete fix: automatic session preload still produced backend-exit crashes on EL9 aarch64. Avoid enabling both stacks in the same backend until a compatible dependency combination is verified; isolate their use when necessary. The tested EL10 and Debian/Ubuntu combinations did not reproduce this failure, which does not establish compatibility for arbitrary other dependency versions. This is a Pigsty package-stack boundary, not an upstream requirement to preload PGroonga for ordinary search.

MeCab tokenization additionally needs the matching Groonga tokenizer plugin and dictionary. Installing the PostgreSQL extension alone does not supply every optional tokenizer; verify the requested tokenizer before building an index that names it.
