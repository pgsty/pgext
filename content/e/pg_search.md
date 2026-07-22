---
title: "pg_search"
linkTitle: "pg_search"
description: "Full text search for PostgreSQL using BM25"
weight: 2100
categories: ["FTS"]
width: full
---

[**pg_search**](https://github.com/paradedb/paradedb/tree/dev/pg_search) : Full text search for PostgreSQL using BM25


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2100** | {{< badge content="pg_search" link="https://github.com/paradedb/paradedb/tree/dev/pg_search" >}} | {{< ext "pg_search" >}} | `0.24.3` | {{< category "FTS" >}} | {{< license "AGPL-3.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `paradedb` |
|   **See Also**    | {{< ext "pgroonga" >}} {{< ext "pgroonga_database" >}} {{< ext "pg_bestmatch" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_bigm" >}} {{< ext "zhparser" >}} {{< ext "pg_tokenizer" >}} {{< ext "pg_trgm" >}} |

> [!Note] bm25 access method conflicts with pg_textsearch and vchord_bm25; PIGSTY packaging uses pgrx 0.19.1 and the pinned builder Rust toolchain instead of upstream rolling stable.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.24.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `pg_search` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.24.3` | {{< bg "18" "pg_search_18" "green" >}} {{< bg "17" "pg_search_17" "green" >}} {{< bg "16" "pg_search_16" "green" >}} {{< bg "15" "pg_search_15" "green" >}} {{< bg "14" "pg_search_14" "red" >}} | `pg_search_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.24.3` | {{< bg "18" "postgresql-18-pg-search" "green" >}} {{< bg "17" "postgresql-17-pg-search" "green" >}} {{< bg "16" "postgresql-16-pg-search" "green" >}} {{< bg "15" "postgresql-15-pg-search" "green" >}} {{< bg "14" "postgresql-14-pg-search" "red" >}} | `postgresql-$v-pg-search` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_search_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_search_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_search_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_search_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_search_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "pg_search_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_search_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.5" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.5" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-15-pg-search : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-search : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.24.3" "postgresql-15-pg-search : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-search : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_18` | `0.24.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 70.9 MiB | [pg_search_18-0.24.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_18-0.24.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_search_18` | `0.24.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 68.2 MiB | [pg_search_18-0.24.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_18-0.24.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_search_18` | `0.24.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 69.2 MiB | [pg_search_18-0.24.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_18-0.24.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_search_18` | `0.24.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 68.1 MiB | [pg_search_18-0.24.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_18-0.24.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_search_18` | `0.24.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 69.1 MiB | [pg_search_18-0.24.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_18-0.24.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_search_18` | `0.24.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 68.0 MiB | [pg_search_18-0.24.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_18-0.24.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-search` | `0.24.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 65.2 MiB | [postgresql-18-pg-search_0.24.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-18-pg-search_0.24.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-search` | `0.24.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 62.3 MiB | [postgresql-18-pg-search_0.24.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-18-pg-search_0.24.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-search` | `0.24.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 65.2 MiB | [postgresql-18-pg-search_0.24.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-18-pg-search_0.24.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-search` | `0.24.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 62.3 MiB | [postgresql-18-pg-search_0.24.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-18-pg-search_0.24.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-search` | `0.24.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 67.3 MiB | [postgresql-18-pg-search_0.24.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-18-pg-search_0.24.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-search` | `0.24.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 65.7 MiB | [postgresql-18-pg-search_0.24.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-18-pg-search_0.24.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-search` | `0.24.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 67.2 MiB | [postgresql-18-pg-search_0.24.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-18-pg-search_0.24.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-search` | `0.24.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 65.5 MiB | [postgresql-18-pg-search_0.24.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-18-pg-search_0.24.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-search` | `0.24.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 67.0 MiB | [postgresql-18-pg-search_0.24.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-18-pg-search_0.24.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-search` | `0.24.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 65.4 MiB | [postgresql-18-pg-search_0.24.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-18-pg-search_0.24.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_17` | `0.24.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 71.0 MiB | [pg_search_17-0.24.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_17-0.24.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_search_17` | `0.24.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 68.2 MiB | [pg_search_17-0.24.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_17-0.24.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_search_17` | `0.24.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 69.2 MiB | [pg_search_17-0.24.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_17-0.24.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_search_17` | `0.24.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 68.1 MiB | [pg_search_17-0.24.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_17-0.24.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_search_17` | `0.24.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 69.2 MiB | [pg_search_17-0.24.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_17-0.24.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_search_17` | `0.24.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 68.0 MiB | [pg_search_17-0.24.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_17-0.24.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-search` | `0.24.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 65.2 MiB | [postgresql-17-pg-search_0.24.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-17-pg-search_0.24.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-search` | `0.24.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 62.3 MiB | [postgresql-17-pg-search_0.24.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-17-pg-search_0.24.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-search` | `0.24.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 65.2 MiB | [postgresql-17-pg-search_0.24.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-17-pg-search_0.24.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-search` | `0.24.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 62.3 MiB | [postgresql-17-pg-search_0.24.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-17-pg-search_0.24.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-search` | `0.24.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 67.3 MiB | [postgresql-17-pg-search_0.24.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-17-pg-search_0.24.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-search` | `0.24.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 65.6 MiB | [postgresql-17-pg-search_0.24.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-17-pg-search_0.24.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-search` | `0.24.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 67.2 MiB | [postgresql-17-pg-search_0.24.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-17-pg-search_0.24.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-search` | `0.24.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 65.5 MiB | [postgresql-17-pg-search_0.24.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-17-pg-search_0.24.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-search` | `0.24.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 67.1 MiB | [postgresql-17-pg-search_0.24.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-17-pg-search_0.24.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-search` | `0.24.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 65.4 MiB | [postgresql-17-pg-search_0.24.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-17-pg-search_0.24.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_16` | `0.24.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 71.0 MiB | [pg_search_16-0.24.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_16-0.24.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_search_16` | `0.24.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 68.2 MiB | [pg_search_16-0.24.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_16-0.24.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_search_16` | `0.24.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 69.3 MiB | [pg_search_16-0.24.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_16-0.24.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_search_16` | `0.24.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 68.1 MiB | [pg_search_16-0.24.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_16-0.24.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_search_16` | `0.24.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 69.1 MiB | [pg_search_16-0.24.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_16-0.24.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_search_16` | `0.24.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 68.1 MiB | [pg_search_16-0.24.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_16-0.24.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-search` | `0.24.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 65.2 MiB | [postgresql-16-pg-search_0.24.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-16-pg-search_0.24.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-search` | `0.24.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 62.3 MiB | [postgresql-16-pg-search_0.24.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-16-pg-search_0.24.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-search` | `0.24.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 65.2 MiB | [postgresql-16-pg-search_0.24.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-16-pg-search_0.24.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-search` | `0.24.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 62.3 MiB | [postgresql-16-pg-search_0.24.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-16-pg-search_0.24.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-search` | `0.24.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 67.3 MiB | [postgresql-16-pg-search_0.24.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-16-pg-search_0.24.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-search` | `0.24.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 65.7 MiB | [postgresql-16-pg-search_0.24.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-16-pg-search_0.24.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-search` | `0.24.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 67.2 MiB | [postgresql-16-pg-search_0.24.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-16-pg-search_0.24.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-search` | `0.24.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 65.5 MiB | [postgresql-16-pg-search_0.24.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-16-pg-search_0.24.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-search` | `0.24.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 67.1 MiB | [postgresql-16-pg-search_0.24.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-16-pg-search_0.24.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-search` | `0.24.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 65.4 MiB | [postgresql-16-pg-search_0.24.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-16-pg-search_0.24.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_15` | `0.24.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 70.9 MiB | [pg_search_15-0.24.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_15-0.24.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_search_15` | `0.24.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 68.1 MiB | [pg_search_15-0.24.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_15-0.24.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_search_15` | `0.24.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 69.2 MiB | [pg_search_15-0.24.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_15-0.24.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_search_15` | `0.24.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 68.1 MiB | [pg_search_15-0.24.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_15-0.24.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_search_15` | `0.24.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 69.1 MiB | [pg_search_15-0.24.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_15-0.24.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_search_15` | `0.24.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 68.0 MiB | [pg_search_15-0.24.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_15-0.24.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-search` | `0.24.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 65.2 MiB | [postgresql-15-pg-search_0.24.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-15-pg-search_0.24.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-search` | `0.24.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 62.3 MiB | [postgresql-15-pg-search_0.24.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-15-pg-search_0.24.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-search` | `0.24.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 65.2 MiB | [postgresql-15-pg-search_0.24.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-15-pg-search_0.24.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-search` | `0.24.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 62.3 MiB | [postgresql-15-pg-search_0.24.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-15-pg-search_0.24.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-search` | `0.24.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 67.3 MiB | [postgresql-15-pg-search_0.24.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-15-pg-search_0.24.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-search` | `0.24.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 65.7 MiB | [postgresql-15-pg-search_0.24.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-15-pg-search_0.24.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-search` | `0.24.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 67.1 MiB | [postgresql-15-pg-search_0.24.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-15-pg-search_0.24.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-search` | `0.24.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 65.5 MiB | [postgresql-15-pg-search_0.24.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-15-pg-search_0.24.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-search` | `0.24.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 67.0 MiB | [postgresql-15-pg-search_0.24.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-15-pg-search_0.24.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-search` | `0.24.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 65.4 MiB | [postgresql-15-pg-search_0.24.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-15-pg-search_0.24.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-14-pg-search` | `0.20.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.6 MiB | [postgresql-14-pg-search_0.20.7_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-14-pg-search/postgresql-14-pg-search_0.20.7_amd64.deb) |
| `postgresql-14-pg-search` | `0.20.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 45.0 MiB | [postgresql-14-pg-search_0.20.7_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-14-pg-search/postgresql-14-pg-search_0.20.7_arm64.deb) |
| `postgresql-14-pg-search` | `0.20.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.1 MiB | [postgresql-14-pg-search_0.20.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-14-pg-search_0.20.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-search` | `0.20.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 44.5 MiB | [postgresql-14-pg-search_0.20.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-14-pg-search_0.20.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-search` | `0.20.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 45.6 MiB | [postgresql-14-pg-search_0.20.7_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-14-pg-search/postgresql-14-pg-search_0.20.7_amd64.deb) |
| `postgresql-14-pg-search` | `0.20.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 45.0 MiB | [postgresql-14-pg-search_0.20.7_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-14-pg-search/postgresql-14-pg-search_0.20.7_arm64.deb) |
| `postgresql-14-pg-search` | `0.20.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 45.6 MiB | [postgresql-14-pg-search_0.20.7_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-14-pg-search/postgresql-14-pg-search_0.20.7_amd64.deb) |
| `postgresql-14-pg-search` | `0.20.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 45.1 MiB | [postgresql-14-pg-search_0.20.7_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-14-pg-search/postgresql-14-pg-search_0.20.7_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/paradedb/paradedb/tree/dev/pg_search" title="Repository" icon="github" subtitle="github.com/paradedb/paradedb/tree/dev/pg_search" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_search-0.24.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_search;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_search;		# install via package name, for the active PG version

pig install pg_search -v 18;   # install for PG 18
pig install pg_search -v 17;   # install for PG 17
pig install pg_search -v 16;   # install for PG 16
pig install pg_search -v 15;   # install for PG 15

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_search';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_search;
```

## Usage

Sources:

- [pg_search v0.24.3 README](https://github.com/paradedb/paradedb/blob/v0.24.3/pg_search/README.md)
- [pg_search v0.24.3 release notes](https://github.com/paradedb/paradedb/releases/tag/v0.24.3)
- [ParadeDB documentation](https://docs.paradedb.com/)

pg_search adds ParadeDB's BM25 access method and query operators to PostgreSQL for ranked full-text, structured, and hybrid search. Use it when search must remain transactional with PostgreSQL data and must support BM25 scoring, highlighting, filters, aggregates, and joins.

### Create the Extension

    CREATE EXTENSION pg_search;

Upstream v0.24.3 supports PostgreSQL 15 and later. Use a build packaged for the exact PostgreSQL major version. The extension participates in planner and executor paths, so test query plans and resource use before production upgrades.

### Build a BM25 Index

Create a BM25 index with a stable unique key field:

    CREATE INDEX products_search_idx
    ON products
    USING bm25 (
      id,
      title,
      description,
      category,
      rating
    )
    WITH (key_field = 'id');

Keep the key field unique and non-null. Index only fields used by search or filtering; every indexed field increases build time, disk use, and write amplification.

### Query, Rank, and Highlight

The @@@ operator matches a field or indexed row against a ParadeDB query:

    SELECT id,
           title,
           paradedb.score(id) AS score,
           paradedb.snippet(description) AS snippet
    FROM products
    WHERE description @@@ 'wireless keyboard'
      AND category = 'electronics'
    ORDER BY score DESC
    LIMIT 20;

Use field-qualified query strings or the paradedb query constructors when user input must be constrained. Do not concatenate untrusted input into query syntax without validation.

For boolean queries, paradedb.boolean() can combine must, should, and must_not clauses and can set minimum_should_match. The extension also exposes index_created_at() for inspecting index creation time.

### User-Facing Object Index

- bm25: index access method for text and structured fields.
- @@@: search-match operator used in WHERE clauses.
- paradedb.score(key): BM25 relevance score for a matching row.
- paradedb.snippet(field): highlighted excerpt for the current match.
- paradedb.parse(...), paradedb.term(...), paradedb.boolean(...): typed query constructors.
- paradedb.index_info(...): index metadata and field configuration.
- paradedb.index_created_at(...): index creation timestamp.

### Version 0.24.3 Operational Changes

The 0.24.x line enables more aggregate and join scan paths and adds time and timetz support. Version 0.24.3 also bounds sequential-scan buffering, caps index-build worker memory, checks available disk space earlier, fixes GROUP BY cardinality routing, and raises an error when Tantivy would truncate a value.

These safeguards reduce runaway resource use but do not eliminate capacity planning. Monitor temporary space, index size, build duration, and query memory. Re-run representative EXPLAIN ANALYZE plans after upgrading because planner behavior can change.

### Compatibility and Caveats

- pg_search uses its own BM25 index implementation. Do not assume an index created by another extension is interchangeable.
- Local catalog metadata reports a bm25 access-method conflict with pg_textsearch and vchord_bm25; avoid loading competing implementations in the same database unless their documentation explicitly supports coexistence.
- Search indexes must be maintained with the table and can materially increase update cost.
- Ranking is query- and corpus-dependent. Benchmark with production-like text and filters rather than treating example scores as portable.
