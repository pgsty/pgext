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
| **2100** | {{< badge content="pg_search" link="https://github.com/paradedb/paradedb/tree/dev/pg_search" >}} | {{< ext "pg_search" >}} | `0.23.1` | {{< category "FTS" >}} | {{< license "AGPL-3.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `paradedb` |
|   **See Also**    | {{< ext "pgroonga" >}} {{< ext "pgroonga_database" >}} {{< ext "pg_bestmatch" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_bigm" >}} {{< ext "zhparser" >}} {{< ext "pg_tokenizer" >}} {{< ext "pg_trgm" >}} |

> [!Note] bm25 am conflicts with pg_textsearch; PG15-16 require shared_preload_libraries while PG17-18 do not.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.23.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `pg_search` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.23.1` | {{< bg "18" "pg_search_18" "green" >}} {{< bg "17" "pg_search_17" "green" >}} {{< bg "16" "pg_search_16" "green" >}} {{< bg "15" "pg_search_15" "green" >}} {{< bg "14" "pg_search_14" "red" >}} | `pg_search_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.23.1` | {{< bg "18" "postgresql-18-pg-search" "green" >}} {{< bg "17" "postgresql-17-pg-search" "green" >}} {{< bg "16" "postgresql-16-pg-search" "green" >}} {{< bg "15" "postgresql-15-pg-search" "green" >}} {{< bg "14" "postgresql-14-pg-search" "red" >}} | `postgresql-$v-pg-search` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.23.0" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.0" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.0" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.0" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.10" "pg_search_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.23.1" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.10" "pg_search_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.23.1" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.10" "pg_search_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.23.1" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.10" "pg_search_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.23.1" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "pg_search_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_search_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "pg_search_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_search_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.5" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.5" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.23.1" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-search : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-search : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-search : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-search : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-search : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-search : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-search : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-search : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-search : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-search : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_18` | `0.23.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 61.6 MiB | [pg_search_18-0.23.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_18-0.23.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_search_18` | `0.23.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 76.1 MiB | [pg_search_18-0.23.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_18-0.23.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_search_18` | `0.23.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 62.0 MiB | [pg_search_18-0.23.1-1PARADEDB.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_18-0.23.1-1PARADEDB.el9.x86_64.rpm) |
| `pg_search_18` | `0.23.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 76.0 MiB | [pg_search_18-0.23.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_18-0.23.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_search_18` | `0.23.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.0 MiB | [pg_search_18-0.23.1-1PARADEDB.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_18-0.23.1-1PARADEDB.el10.x86_64.rpm) |
| `pg_search_18` | `0.22.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 61.0 MiB | [pg_search_18-0.22.6-1PARADEDB.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_18-0.22.6-1PARADEDB.el10.aarch64.rpm) |
| `postgresql-18-pg-search` | `0.23.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 60.5 MiB | [postgresql-18-pg-search_0.23.1_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-18-pg-search/postgresql-18-pg-search_0.23.1_amd64.deb) |
| `postgresql-18-pg-search` | `0.23.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 58.2 MiB | [postgresql-18-pg-search_0.23.1_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-18-pg-search/postgresql-18-pg-search_0.23.1_arm64.deb) |
| `postgresql-18-pg-search` | `0.23.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 60.5 MiB | [postgresql-18-pg-search_0.23.1_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-18-pg-search/postgresql-18-pg-search_0.23.1_amd64.deb) |
| `postgresql-18-pg-search` | `0.23.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 58.2 MiB | [postgresql-18-pg-search_0.23.1_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-18-pg-search/postgresql-18-pg-search_0.23.1_arm64.deb) |
| `postgresql-18-pg-search` | `0.23.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 60.5 MiB | [postgresql-18-pg-search_0.23.1_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-18-pg-search/postgresql-18-pg-search_0.23.1_amd64.deb) |
| `postgresql-18-pg-search` | `0.23.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 58.2 MiB | [postgresql-18-pg-search_0.23.1_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-18-pg-search/postgresql-18-pg-search_0.23.1_arm64.deb) |
| `postgresql-18-pg-search` | `0.23.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 60.5 MiB | [postgresql-18-pg-search_0.23.1_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-18-pg-search/postgresql-18-pg-search_0.23.1_amd64.deb) |
| `postgresql-18-pg-search` | `0.23.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 58.2 MiB | [postgresql-18-pg-search_0.23.1_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-18-pg-search/postgresql-18-pg-search_0.23.1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_17` | `0.23.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 61.6 MiB | [pg_search_17-0.23.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_17-0.23.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_search_17` | `0.23.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 76.1 MiB | [pg_search_17-0.23.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_17-0.23.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_search_17` | `0.23.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 62.0 MiB | [pg_search_17-0.23.1-1PARADEDB.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_17-0.23.1-1PARADEDB.el9.x86_64.rpm) |
| `pg_search_17` | `0.23.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 76.1 MiB | [pg_search_17-0.23.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_17-0.23.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_search_17` | `0.23.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.0 MiB | [pg_search_17-0.23.1-1PARADEDB.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_17-0.23.1-1PARADEDB.el10.x86_64.rpm) |
| `pg_search_17` | `0.22.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 61.1 MiB | [pg_search_17-0.22.6-1PARADEDB.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_17-0.22.6-1PARADEDB.el10.aarch64.rpm) |
| `postgresql-17-pg-search` | `0.23.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 60.5 MiB | [postgresql-17-pg-search_0.23.1_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-17-pg-search/postgresql-17-pg-search_0.23.1_amd64.deb) |
| `postgresql-17-pg-search` | `0.23.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 58.2 MiB | [postgresql-17-pg-search_0.23.1_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-17-pg-search/postgresql-17-pg-search_0.23.1_arm64.deb) |
| `postgresql-17-pg-search` | `0.23.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 60.5 MiB | [postgresql-17-pg-search_0.23.1_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-17-pg-search/postgresql-17-pg-search_0.23.1_amd64.deb) |
| `postgresql-17-pg-search` | `0.23.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 58.3 MiB | [postgresql-17-pg-search_0.23.1_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-17-pg-search/postgresql-17-pg-search_0.23.1_arm64.deb) |
| `postgresql-17-pg-search` | `0.23.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 60.5 MiB | [postgresql-17-pg-search_0.23.1_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-17-pg-search/postgresql-17-pg-search_0.23.1_amd64.deb) |
| `postgresql-17-pg-search` | `0.23.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 58.2 MiB | [postgresql-17-pg-search_0.23.1_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-17-pg-search/postgresql-17-pg-search_0.23.1_arm64.deb) |
| `postgresql-17-pg-search` | `0.23.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 60.5 MiB | [postgresql-17-pg-search_0.23.1_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-17-pg-search/postgresql-17-pg-search_0.23.1_amd64.deb) |
| `postgresql-17-pg-search` | `0.23.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 58.2 MiB | [postgresql-17-pg-search_0.23.1_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-17-pg-search/postgresql-17-pg-search_0.23.1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_16` | `0.23.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 61.6 MiB | [pg_search_16-0.23.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_16-0.23.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_search_16` | `0.23.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 76.1 MiB | [pg_search_16-0.23.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_16-0.23.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_search_16` | `0.23.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 62.0 MiB | [pg_search_16-0.23.1-1PARADEDB.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_16-0.23.1-1PARADEDB.el9.x86_64.rpm) |
| `pg_search_16` | `0.23.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 76.0 MiB | [pg_search_16-0.23.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_16-0.23.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_search_16` | `0.23.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.1 MiB | [pg_search_16-0.23.1-1PARADEDB.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_16-0.23.1-1PARADEDB.el10.x86_64.rpm) |
| `pg_search_16` | `0.22.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 61.1 MiB | [pg_search_16-0.22.6-1PARADEDB.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_16-0.22.6-1PARADEDB.el10.aarch64.rpm) |
| `postgresql-16-pg-search` | `0.23.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 60.5 MiB | [postgresql-16-pg-search_0.23.1_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-16-pg-search/postgresql-16-pg-search_0.23.1_amd64.deb) |
| `postgresql-16-pg-search` | `0.23.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 58.3 MiB | [postgresql-16-pg-search_0.23.1_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-16-pg-search/postgresql-16-pg-search_0.23.1_arm64.deb) |
| `postgresql-16-pg-search` | `0.23.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 60.5 MiB | [postgresql-16-pg-search_0.23.1_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-16-pg-search/postgresql-16-pg-search_0.23.1_amd64.deb) |
| `postgresql-16-pg-search` | `0.23.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 58.3 MiB | [postgresql-16-pg-search_0.23.1_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-16-pg-search/postgresql-16-pg-search_0.23.1_arm64.deb) |
| `postgresql-16-pg-search` | `0.23.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 60.5 MiB | [postgresql-16-pg-search_0.23.1_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-16-pg-search/postgresql-16-pg-search_0.23.1_amd64.deb) |
| `postgresql-16-pg-search` | `0.23.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 58.3 MiB | [postgresql-16-pg-search_0.23.1_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-16-pg-search/postgresql-16-pg-search_0.23.1_arm64.deb) |
| `postgresql-16-pg-search` | `0.23.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 60.5 MiB | [postgresql-16-pg-search_0.23.1_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-16-pg-search/postgresql-16-pg-search_0.23.1_amd64.deb) |
| `postgresql-16-pg-search` | `0.23.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 58.3 MiB | [postgresql-16-pg-search_0.23.1_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-16-pg-search/postgresql-16-pg-search_0.23.1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_15` | `0.23.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 61.5 MiB | [pg_search_15-0.23.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_15-0.23.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_search_15` | `0.23.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 76.1 MiB | [pg_search_15-0.23.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_15-0.23.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_search_15` | `0.23.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 62.0 MiB | [pg_search_15-0.23.1-1PARADEDB.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_15-0.23.1-1PARADEDB.el9.x86_64.rpm) |
| `pg_search_15` | `0.23.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 76.1 MiB | [pg_search_15-0.23.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_15-0.23.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_search_15` | `0.23.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.0 MiB | [pg_search_15-0.23.1-1PARADEDB.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_15-0.23.1-1PARADEDB.el10.x86_64.rpm) |
| `pg_search_15` | `0.23.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 60.8 MiB | [pg_search_15-0.23.1-1PARADEDB.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_15-0.23.1-1PARADEDB.el10.aarch64.rpm) |
| `postgresql-15-pg-search` | `0.23.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 60.5 MiB | [postgresql-15-pg-search_0.23.1_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-15-pg-search/postgresql-15-pg-search_0.23.1_amd64.deb) |
| `postgresql-15-pg-search` | `0.23.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 58.2 MiB | [postgresql-15-pg-search_0.23.1_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-15-pg-search/postgresql-15-pg-search_0.23.1_arm64.deb) |
| `postgresql-15-pg-search` | `0.23.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 60.5 MiB | [postgresql-15-pg-search_0.23.1_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-15-pg-search/postgresql-15-pg-search_0.23.1_amd64.deb) |
| `postgresql-15-pg-search` | `0.23.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 58.2 MiB | [postgresql-15-pg-search_0.23.1_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-15-pg-search/postgresql-15-pg-search_0.23.1_arm64.deb) |
| `postgresql-15-pg-search` | `0.23.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 60.4 MiB | [postgresql-15-pg-search_0.23.1_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-15-pg-search/postgresql-15-pg-search_0.23.1_amd64.deb) |
| `postgresql-15-pg-search` | `0.23.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 58.2 MiB | [postgresql-15-pg-search_0.23.1_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-15-pg-search/postgresql-15-pg-search_0.23.1_arm64.deb) |
| `postgresql-15-pg-search` | `0.23.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 60.5 MiB | [postgresql-15-pg-search_0.23.1_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-15-pg-search/postgresql-15-pg-search_0.23.1_amd64.deb) |
| `postgresql-15-pg-search` | `0.23.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 58.2 MiB | [postgresql-15-pg-search_0.23.1_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-15-pg-search/postgresql-15-pg-search_0.23.1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_14` | `0.20.10` | [el8.x86_64](/os/el8.x86_64) | pigsty | 46.3 MiB | [pg_search_14-0.20.10-1PARADEDB.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_14-0.20.10-1PARADEDB.el8.x86_64.rpm) |
| `pg_search_14` | `0.20.10` | [el8.aarch64](/os/el8.aarch64) | pigsty | 45.7 MiB | [pg_search_14-0.20.10-1PARADEDB.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_14-0.20.10-1PARADEDB.el8.aarch64.rpm) |
| `pg_search_14` | `0.20.10` | [el9.x86_64](/os/el9.x86_64) | pigsty | 46.1 MiB | [pg_search_14-0.20.10-1PARADEDB.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_14-0.20.10-1PARADEDB.el9.x86_64.rpm) |
| `pg_search_14` | `0.20.10` | [el9.aarch64](/os/el9.aarch64) | pigsty | 46.0 MiB | [pg_search_14-0.20.10-1PARADEDB.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_14-0.20.10-1PARADEDB.el9.aarch64.rpm) |
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
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_search-0.23.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_search;		# build deb
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

- Sources: [ParadeDB extension install docs](https://docs.paradedb.com/deploy/self-hosted/extension), [quickstart](https://docs.paradedb.com/documentation/getting-started/quickstart), [v0.23.0 release](https://github.com/paradedb/paradedb/releases/tag/v0.23.0), [pg_search README](https://github.com/paradedb/paradedb/blob/dev/pg_search/README.md)

`pg_search` is ParadeDB's BM25-based search extension for PostgreSQL. The upstream README says support starts at PostgreSQL 15, and the v0.23.0 self-hosted install docs still require preloading the library before `CREATE EXTENSION`.

### Enable And Create The Extension

```conf
shared_preload_libraries = 'pg_search'
```

```sql
CREATE EXTENSION pg_search;
```

The self-hosted extension docs for v0.23.0 describe prebuilt binaries for Postgres 15+.

### Create A BM25 Index

Quickstart examples use the `bm25` access method with a unique key field:

```sql
CREATE INDEX search_idx ON mock_items
USING bm25 (id, description, category, rating)
WITH (key_field = 'id');
```

The v0.23.0 release also notes newly tunable BM25 `k1` and `b` parameters per field.

### Query Operators And Helpers

The current quickstart uses these query operators:

- `|||`: match disjunction, equivalent to `term1 OR term2`.
- `&&&`: match conjunction, equivalent to `term1 AND term2`.

Examples:

```sql
SELECT description, rating
FROM mock_items
WHERE description ||| 'running shoes'
ORDER BY rating
LIMIT 5;

SELECT description, pdb.score(id)
FROM mock_items
WHERE description &&& 'running shoes'
ORDER BY score DESC
LIMIT 5;

SELECT description, pdb.snippet(description), pdb.score(id)
FROM mock_items
WHERE description ||| 'running shoes'
ORDER BY score DESC
LIMIT 5;
```

### Notes

The development README points users to the docs site for installation and usage instead of documenting SQL details inline. The quickstart is therefore the authoritative usage surface for current `pg_search` syntax, and it reflects the post-0.20 API rather than the older `@@@` examples still found in some secondary materials.
